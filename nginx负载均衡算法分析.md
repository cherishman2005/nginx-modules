# nginx负载均衡算法分析

## 轮询/加权轮询

（1）初始化：nginx-upstream采用轮询方式初始化；

（2）客户端请求：
  * 判定支持的最大连接数；
  * 判定超时失败次数；
  * 判定重试次数；

  * 根据权重计算出best节点；

## ip-hash

ip_hash算法：
```
for (i = 0; i < (ngx_uint_t) iphp->addrlen; i++) {
    hash = (hash * 113 + iphp->addr[i]) % 6271;
}
```

（1）初始化：nginx-upstream采用轮询方式初始化；

（2）客户端请求：
* 先根据客户端IP采用`ip_hash算法`计算hash值；
* 然后对total整体权重取模： hash mod(n)；

* 再hash mod(n) 选中节点；
  * 判定支持的最大连接数；
  * 判定超时失败次数；
  * 判定重试次数；

* 如果重试次数超过20次，则降级为轮询；

## hash

哈希是极高效单向的，不可逆。

（1）初始化：nginx-upstream采用轮询方式初始化；

（2）客户端请求：
* 根据hashKey计算hash；
* 然后对total整体权重取模；

* 再hash mod(n) 选中节点；
  * 判定支持的最大连接数；
  * 判定超时失败次数；
  * 判定重试次数；

* 如果重试次数超过20次，则降级为轮询；

## consistent-hash

（1）初始化：对upstream-servers的ip+port列表生成hash数组；并将每个物理ip+port映射成160个虚拟节点node；然后进行排序；

（2）客户端请求：
* 根据key值结算hash值，然后采用二分法查找到对应的upstream-server后端node（ip+port）；
  * 判定超时失败次数；
  * 判定支持的最大连接数；
  * 判定重试次数；
* 如果重试次数超过20次，则降级为轮询；


# 小结

* nginx的负载均衡算法 在异常处理、防止负载过高等方面，多维度进行了判定处理；—— 非常值得借鉴。

* 微服务系统、rpc的路由算法等常借鉴nginx负载均衡算法。

* 根据业务需求，可以定制化一些负载均衡算法；
  * 如果采用lua、njs开发效率较高；非常通用的算法可以尝试c module；

# FAQ

## hash table

nginx采用hash-table
```
ngx_crc32_init(hash);
ngx_crc32_update(&hash, host, host_len);
ngx_crc32_update(&hash, (u_char *) "", 1);
ngx_crc32_update(&hash, port, port_len);
ngx_crc32_final(hash);
```

## time33 哈希函数

time33 哈希函数是一个很流行的哈希算法被perl使用并出现在Berkeley DB中，这也是其中一个最好的已知针对string字符串的哈希函数因为其算法非常快而且分布（注：分布好代表不容易冲突）非常好。

```
function hash(text) {
    var hash = 5381, index = text.length;

    while (index) {
        hash = (hash * 33) ^ text.charCodeAt(--index);
    }

    return hash >>> 0;
}
```
