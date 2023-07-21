# bufio没有超时机制

how to set timeout on bufio.ReadBytes()


* bufio没有超时机制，不好用。

```
bufio does not have this feature (you can read the docs here).

The more appropriate solution is to use the SetReadDeadline method of the net.Conn value.

    // A deadline is an absolute time after which I/O operations
    // fail instead of blocking. The deadline applies to all future
    // and pending I/O, not just the immediately following call to
    // Read or Write.
 ...
    // SetReadDeadline sets the deadline for future Read calls
    // and any currently-blocked Read call.
    // A zero value for t means Read will not time out.
For example:

conn.SetReadDeadline(time.Now().Add(time.Second))
```

# 参考链接

- [https://stackoverflow.com/questions/72419951/how-to-set-timeout-on-bufio-readbytes](https://stackoverflow.com/questions/72419951/how-to-set-timeout-on-bufio-readbytes)
