package main

import (
    "fmt"
    "hash/crc32"
)

type RedisOperator struct {
    data string
}

func NewRedisSentinelOperator(addrs []string, masterName string) *RedisOperator {
	return &RedisOperator{
		data: fmt.Sprintf("addrs %v, masterName %s", addrs, masterName),
	}
}



type RedisHelper struct {
	redisOperator []*RedisOperator
}

func NewRedisHelper(addrs []string, masterName []string) *RedisHelper {
    helper := &RedisHelper{}
	for i, v := range masterName {
		fmt.Printf("index:%d  value:%s\n", i, v)
        operator := NewRedisSentinelOperator(addrs, v)
        helper.redisOperator = append(helper.redisOperator, operator)
	}
    
	return helper
}

func (this *RedisHelper) GetRedisOperator(s string) *RedisOperator {
    key := Hash(s)
    i := key % len(this.redisOperator)
    fmt.Printf("key:%d  i:%s\n", key, i)
    return this.redisOperator[i]
}

func Hash(s string) int {
    v := int(crc32.ChecksumIEEE([]byte(s)))
    if v >= 0 {
            return v
    }
    if -v >= 0 {
            return -v
    }
    // v == MinInt
    return 0
}


var helper *RedisHelper

func main() {
    addrs := []string{"group1013.com:20075","group1013.com:20075","group1013.**.com:20075"}
    mastername := []string{"redis_001", "redis_002", "redis_003"}
    helper = NewRedisHelper(addrs, mastername)
    
    key := "123456abcd"
    r := helper.GetRedisOperator(key)
    
    fmt.Printf("result:%s\n", r)
}
