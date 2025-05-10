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

# pprof说明

Profile Descriptions:
```txt
allocs: A sampling of all past memory allocations
block: Stack traces that led to blocking on synchronization primitives
cmdline: The command line invocation of the current program
goroutine: Stack traces of all current goroutines
heap: A sampling of memory allocations of live objects. You can specify the gc GET parameter to run GC before taking the heap sample.
mutex: Stack traces of holders of contended mutexes
profile: CPU profile. You can specify the duration in the seconds GET parameter. After you get the profile file, use the go tool pprof command to investigate the profile.
threadcreate: Stack traces that led to the creation of new OS threads
trace: A trace of execution of the current program. You can specify the duration in the seconds GET parameter. After you get the trace file, use the go tool trace command to investigate the trace.
```
