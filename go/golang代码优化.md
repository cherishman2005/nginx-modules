# golang代码优化


“Go 语言中的好代码与差代码”（https://medium.com/@teivah/good-code-vs-bad-code-in-golang-84cb3c5da49d）的文章，作者一步步地向我们介绍了一个实际业务用例的重构。

利用 Go 语言的特性将“差代码”转换成“好代码”，即更加符合惯例和更易读的代码。

为了正确地确定性能的优先级，最有价值的策略是找到瓶颈，然后集中精力改善。可以使用分析工具来做！例如 Pprof（https://blog.golang.org/profiling-go-programs） 和 Trace（https://making.pusher.com/go-tool-trace/）：


## 小结（tips）

Here are a few disclaimers and take-aways:

* Performance can be improved at many levels of abstraction, using different techniques, and the gains are multiplicative.
* Tune the high-level abstractions first: data structures, algorithms, proper decoupling. Tune the low-level abstractions later: I/O, batching, concurrency, stdlib usage, memory management.
* Big-O analysis is fundamental but often it’s not the relevant tool to make a given program run faster.
* Benchmarking is hard. Use profiling and benchmarks to discover bottlenecks and get insight about your code. Keep in mind that the benchmark results are not the “real” latencies experienced by the end-users in production, and take the numbers with a grain of salt.
* Fortunately, the tooling (Bench, Pprof, Trace, Race detector, Cover) makes performance exploration approachable and exciting.
* Writing good, relevant tests is non-trivial. But they are extremely precious to help “stay on track”, i.e. to refactor while preserving the original correctness and semantics.
* Take a moment to ask yourself how fast will be “fast enough”. Don’t waste time over-optimizing a one-shot script. Consider that optimization comes with a cost: engineering time, complexity, bugs, technical debt.
* Think twice before obscuring the code.
* Algorithms in Ω(n²) and above are usually expensive.
* Complexity in O(n) or O(n log n), or below, is usually fine.
* The hidden factors are not negligible! For example, all of the improvements in the article were achieved by lowering those factors, and not by changing the complexity class of the algorithm.
* I/O is often a bottleneck: network requests, DB queries, filesystem.
* Regular expressions tend to be a more expensive solution than really needed.
* Memory allocation is more expensive than computations.
* An object in the stack is cheaper than an object in the heap.
* Slices are useful as an alternative to costly reallocation.
* Strings are efficient for read-only usage (including reslicing), but for any other manipulations, []byte are more efficient.
* Memory locality is important (CPU-cache-friendliness).
* Concurrency and parallelism are useful, but tricky to get right.
* When digging deeper and lower-level, there’s a “glass floor” that you don’t really want to break through in go. If you crave asm instructions, intrinsics, SIMD… maybe you should consider go for prototyping, and then switch to a lower-level language to take full advantage of the hardware and of every nanosecond!


# 参考链接

- [Go 代码重构：23 倍性能提升](https://blog.csdn.net/csdnnews/article/details/81009432)

- [https://medium.com/@val_deleplace/go-code-refactoring-the-23x-performance-hunt-156746b522f7](https://medium.com/@val_deleplace/go-code-refactoring-the-23x-performance-hunt-156746b522f7)
