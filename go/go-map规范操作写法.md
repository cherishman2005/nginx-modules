# go-map规范操作写法

```
g_IfaceMap         map[string]*IfaceInfo
```

写法1：
```
	//if len(this.g_IfaceMap) == 0 || this.g_IfaceMap[ifName] == nil {
```

写法2：
```
	if _, ok := this.g_IfaceMap[ifName]; !ok {
```


写法2 是不是更规范？ 哪个性能好？

## 点评

写法2 规范，性能差别不大
