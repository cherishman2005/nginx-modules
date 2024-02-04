# Go检测端口是否被占用

Go检查服务器端口是否被占用，方法有很多，分享两种常用的方法。

推荐方式一，linux、windows均适用。

## 1、方式一

判断某个端口是否可用，最直接的办法，用一下就知道啦~

利用net.Listen在指定端口，新建一个tcp监听。

如果监听成功，就说明端口可用。（监听成功后，立马关闭监听，避免影响使用）

```
// PortCheck 检查端口是否可用，可用-true 不可用-false
func PortCheck(port int) bool {
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", strconv.Itoa(port)))

	if err != nil {
		return false
	}
	defer l.Close()
	return true
}
```

注意：defer l.Close()必须在err处理之后，否则可能panic。


## 2、方式二

该方法有一个缺点：在某些服务器上，可能没有lsof工具。
```
// portCheck 检查端口是否可用，可用-true 不可用-false
func portCheck(port int) bool {
	checkStatement := fmt.Sprintf("lsof -i:%d ", port)
	output, _ := exec.Command("sh", "-c", checkStatement).CombinedOutput()
	if len(output) > 0 {
		return true
	}
	return false
}
```
