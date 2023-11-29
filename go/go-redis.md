# go-redis

## redis.Nil

redis.Nil 是一种特殊的错误，严格意义上来说它并不是错误，而是代表一种状态，例如你使用 Get 命令获取 key 的值，当 key 不存在时，返回 redis.Nil。在其他比如 BLPOP 、 ZSCORE 也有类似的响应，你需要区分错误：
```
val, err := rdb.Get(ctx, "key").Result()
switch {
case err == redis.Nil:
	fmt.Println("key不存在")
case err != nil:
	fmt.Println("错误", err)
case val == "":
	fmt.Println("值是空字符串")
}
```

# 参考文档

- [redis-nil](https://redis.uptrace.dev/zh/guide/go-redis.html#redis-nil)
