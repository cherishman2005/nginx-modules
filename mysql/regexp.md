# regexp

## 查找索引

* func (*Regexp) FindIndex
```
func (re *Regexp) FindIndex(b []byte) (loc []int)
```

Find返回保管正则表达式re在b中的最左侧的一个匹配结果的起止位置的切片（显然len(loc)==2）。匹配结果可以通过起止位置对b做切片操作得到：b[loc[0]:loc[1]]。如果没有匹配到，会返回nil。
