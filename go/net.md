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


## net.SplitHostPort

```
func SplitHostPort 
func SplitHostPort(hostport string) (host, port string, err error)
```
SplitHostPort splits a network address of the form "host:port", "host%zone:port", "[host]:port" or "[host%zone]:port" into host or host%zone and port.

A literal IPv6 address in hostport must be enclosed in square brackets, as in "[::1]:80", "[::1%lo0]:80".

See func Dial for a description of the hostport parameter, and host and port results.

# 参考链接

- [https://cs.opensource.google/go/go/+/refs/tags/go1.19.3:src/net/net.go](https://cs.opensource.google/go/go/+/refs/tags/go1.19.3:src/net/net.go)
