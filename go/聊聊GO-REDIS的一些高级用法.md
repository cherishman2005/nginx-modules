# 1. 前言

说到Golang的Redis库，用到最多的恐怕是
redigo 和 go-redis。其中 redigo 不支持对集群的访问。
本文想聊聊go-redis 2个高级用法node

# 2. 开启对Cluster中Slave Node的访问
在一个负载比较高的Redis Cluster中，若是容许对slave节点进行读操做将极大的提升集群的吞吐能力。git

开启对Slave 节点的访问，受如下3个参数的影响github
```
type ClusterOptions struct {
    // Enables read-only commands on slave nodes.
    ReadOnly bool
    // Allows routing read-only commands to the closest master or slave node.
    // It automatically enables ReadOnly.
    RouteByLatency bool
    // Allows routing read-only commands to the random master or slave node.
    // It automatically enables ReadOnly.
    RouteRandomly bool
    ... 
}
```
go-redis 选择节点的逻辑以下redis
```
func (c *ClusterClient) cmdSlotAndNode(cmd Cmder) (int, *clusterNode, error) {
    state, err := c.state.Get()
    if err != nil {
        return 0, nil, err
    }

    cmdInfo := c.cmdInfo(cmd.Name())
    slot := cmdSlot(cmd, cmdFirstKeyPos(cmd, cmdInfo))

    if c.opt.ReadOnly && cmdInfo != nil && cmdInfo.ReadOnly {
        if c.opt.RouteByLatency {
            node, err := state.slotClosestNode(slot)
            return slot, node, err
        }

        if c.opt.RouteRandomly {
            node := state.slotRandomNode(slot)
            return slot, node, nil
        }

        node, err := state.slotSlaveNode(slot)
        return slot, node, err
    }

    node, err := state.slotMasterNode(slot)
    return slot, node, err
}
```

若是ReadOnly = true，只选择Slave Node

若是ReadOnly = true 且 RouteByLatency = true 将从slot对应的Master Node 和 Slave Node选择，选择策略为: 选择PING 延迟最低的节点

若是ReadOnly = true 且 RouteRandomly = true 将从slot对应的Master Node 和 Slave Node选择，选择策略为:随机选择

# 3. 在集群模式下使用pipeline功能
Redis的pipeline功能的原理是 Client经过一次性将多条redis命令发往Redis Server，减小了每条命令分别传输的IO开销。同时减小了系统调用的次数，所以提高了总体的吞吐能力。算法

咱们在主-从模式的Redis中，pipeline功能应该用的不少，可是Cluster模式下，估计尚未几我的用过。
咱们知道 redis cluster 默认分配了 16384 个slot，当咱们set一个key 时，会用CRC16算法来取模获得所属的slot，而后将这个key 分到哈希槽区间的节点上，具体算法就是：CRC16(key) % 16384。若是咱们使用pipeline功能，一个批次中包含的多条命令，每条命令涉及的key可能属于不一样的slot并发

go-redis 为了解决这个问题, 分为3步
源码能够阅读 defaultProcessPipeline
1) 将计算command 所属的slot, 根据slot选择合适的Cluster Node
2）将同一个Cluster Node 的全部command，放在一个批次中发送（并发操做）
3）接收结果dom

注意：这里go-redis 为了处理简单，每一个command 只能涉及一个key, 不然你可能会收到以下错误code

err CROSSSLOT Keys in request don't hash to the same slot
也就是说go-redis不支持相似 MGET 命令的用法ip

一个简单的例子get
```
package main

import (
    "github.com/go-redis/redis"
    "fmt"
)

func main() {
    client := redis.NewClusterClient(&redis.ClusterOptions{
        Addrs: []string{"192.168.120.110:6380", "192.168.120.111:6380"},
        ReadOnly: true,
        RouteRandomly: true,
    })

    pipe := client.Pipeline()
    pipe.HGetAll("1371648200")
    pipe.HGetAll("1371648300")
    pipe.HGetAll("1371648400")
    cmders, err := pipe.Exec()
    if err != nil {
        fmt.Println("err", err)
    }
    for _, cmder := range cmders {
        cmd := cmder.(*redis.StringStringMapCmd)
        strMap, err := cmd.Result()
        if err != nil {
            fmt.Println("err", err)
        }
        fmt.Println("strMap", strMap)
    }
}
```
# 参考链接

- [https://www.shangmayuan.com/a/dcb403dcb4e94c22ab6aab5e.html](https://www.shangmayuan.com/a/dcb403dcb4e94c22ab6aab5e.html)
