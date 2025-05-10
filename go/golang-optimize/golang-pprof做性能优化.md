# golang-pprof做性能优化

分析
* CPU
* heap内存
* goroutine
* trace

主要是通过cpu，内存快照对比分析，定位到要优化的模块。

## CPU

## heap

## goroutine

http://localhost:9090/debug/pprof/goroutine?debug=1
![image](https://github.com/user-attachments/assets/9f8f6056-6f34-4329-a6f7-80c7a06b4716)

## goroutine调用的详细堆栈信息

http://localhost:9090/debug/pprof/goroutine?debug=2

![image](https://github.com/user-attachments/assets/0742eca1-08d0-4147-8b4e-0263d14c345a)
