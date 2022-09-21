# go-testing


调用以下命令:
```
go test -run=xxx -bench=. -benchtime="3s" -cpuprofile profile_cpu.out
```

该命令会跳过单元测试，执行所有benchmark,同时生成一个cpu性能描述文件.

这里有两个注意点:

-benchtime 可以控制benchmark的运行时间
b.ReportAllocs() ，在report中包含内存分配信息，例如结果是:
```
BenchmarkStringJoin1-4 300000 4351 ns/op 32 B/op 2 allocs/op
```
-4表示4个CPU线程执行；300000表示总共执行了30万次；4531ns/op，表示每次执行耗时4531纳秒；32B/op表示每次执行分配了32字节内存；2 allocs/op表示每次执行分配了2次对象。


## 压测结果

![image](https://user-images.githubusercontent.com/17688273/191460167-d7f8003b-95ee-406c-9488-d25dc1a3b5f0.png)

| map(set)    | 占用内存   |
| --------   | -----:  |
| map[int]int      | 88M   |
| map[int]bool      | 55M   |
| map[string]int      | 132M   |
| map[string]bool      | 99M   |
