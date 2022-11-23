# pprof


| 类型  | 含义 |
| ------------- | ------------- |
| inuse_space | amount of memory allocated and not released yet  |
| inuse_objects   | amount of objects allocated and not released yet |
| alloc_space | total amount of memory allocated (regardless of released)  |
| alloc_objects   | total amount of objects allocated (regardless of released) |

示例
```
go tool pprof --inuse_objects http://localhost:8080/debug/pprof/heap
```
