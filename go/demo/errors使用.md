# errors使用


## 错误类型的比较

代码中经常会出现err == nil 或者err == ErrNotExist之类的判断，对于error类型，由于其是interface类型，实际比较的是interface接口对象实体的地址。
也就是说，重复的new两个文本内容一样的error对象，这两个对象并不相等，因为比较的是这两个对象的地址。这是完全不同的两个对象

```
// 展示了error比较代码
if errors.New("hello error") == errors.New("hello error") { // false
}
errhello := errors.New("hello error")
if errhello == errhello { // true
}
```
