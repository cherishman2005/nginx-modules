# net

```
func (e *OpError) Temporary() bool {
	// Treat ECONNRESET and ECONNABORTED as temporary errors when
	// they come from calling accept. See issue 6163.
	if e.Op == "accept" && isConnError(e.Err) {
		return true
	}

	if ne, ok := e.Err.(*os.SyscallError); ok {
		t, ok := ne.Err.(temporary)
		return ok && t.Temporary()
	}
	t, ok := e.Err.(temporary)
	return ok && t.Temporary()
}
```

# 参考链接

- [https://cs.opensource.google/go/go/+/refs/tags/go1.19.3:src/net/net.go](https://cs.opensource.google/go/go/+/refs/tags/go1.19.3:src/net/net.go)
