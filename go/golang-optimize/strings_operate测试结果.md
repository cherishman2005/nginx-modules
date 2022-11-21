# strings_operate测试结果

```
 go test -run=xxx -bench=. -benchtime="3s"  -count=5 -cpuprofile profile_cpu.out



goos: linux
goarch: amd64
pkg: hello
Benchmark_AppendWithAdd                        1        3538478236 ns/op        30393257232 B/op          100012 allocs/op
Benchmark_AppendWithAdd                        1        3628653777 ns/op        30393254528 B/op          100005 allocs/op
Benchmark_AppendWithAdd                        1        3558503466 ns/op        30393254336 B/op          100001 allocs/op
Benchmark_AppendWithAdd                        1        3633243426 ns/op        30393237952 B/op           99999 allocs/op
Benchmark_AppendWithAdd                        1        3368539698 ns/op        30393238240 B/op          100001 allocs/op
Benchmark_AppendWithSprintf                    1        7520981331 ns/op        60449155696 B/op          381521 allocs/op
Benchmark_AppendWithSprintf                    1        7447245243 ns/op        60449166032 B/op          381793 allocs/op
Benchmark_AppendWithSprintf                    1        7434540530 ns/op        60449198512 B/op          382239 allocs/op
Benchmark_AppendWithSprintf                    1        7668287532 ns/op        60449212624 B/op          382510 allocs/op
Benchmark_AppendWithSprintf                    1        7581239513 ns/op        60449186312 B/op          382362 allocs/op
Benchmark_AppendWithBytesBuffer             3529           1068675 ns/op         2912690 B/op         16 allocs/op
Benchmark_AppendWithBytesBuffer             3300           1059010 ns/op         2912688 B/op         16 allocs/op
Benchmark_AppendWithBytesBuffer             3397           1007932 ns/op         2912692 B/op         16 allocs/op
Benchmark_AppendWithBytesBuffer             3442           1023023 ns/op         2912688 B/op         16 allocs/op
Benchmark_AppendWithBytesBuffer             3542           1004720 ns/op         2912699 B/op         16 allocs/op
Benchmark_AppendWithStringBuilder           6133            571913 ns/op         2930682 B/op         31 allocs/op
Benchmark_AppendWithStringBuilder           6242            539275 ns/op         2930680 B/op         31 allocs/op
Benchmark_AppendWithStringBuilder           6570            562604 ns/op         2930681 B/op         31 allocs/op
Benchmark_AppendWithStringBuilder           6099            547193 ns/op         2930680 B/op         31 allocs/op
Benchmark_AppendWithStringBuilder           6440            547894 ns/op         2930682 B/op         31 allocs/op
PASS
ok      hello   91.569s
```
