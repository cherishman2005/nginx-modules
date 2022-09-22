# Go如何优雅的解决方法重载的问题


## 1.问题

C++支持方法重载，GO不支持：
```
func Handler()  {
	//...	
}

func Handler(timeOut time.Duration) {
 //.. 
}

func Handler(timeOut time.Duration, retry int) {
	//..
}
```

如果要处理超时和重试等，或者有些可以是默认值。怎么办？

写三个方法吗，然后方法名字不同，三个方法名字，这种做法正确吗？显然是不合理的。如何解决呢。这就是go的缺点，缺点就是不支持方法的重载。

## 2.解决问题

其实解决的思想很简单，我们知道有可变参数，那么到底传入什么样的参数。

难道是这样子？
```
func Handler(op ...interface{}) {

}
```
太不合理了吧，如何解析呢。

能不能这样子？
```
type Op struct {
	TimeOut time.Duration
	Retry   int
}
func Handler(op *Op) {

}
```

但是人家不想传入参数的情况 怎么解决？换种思维，其实很多框架都使用了这种方式。
```
type Option func(*Options)
type Options struct {
	TimeOut     time.Duration
	RetryMaxNum int
}
func loadOp(option ...Option) *Options {
	options := new(Options)
	for _, e := range option {
		e(options)
	}
	return options
}


// 最终这样子
func Handler(option ...Option) {
	op := loadOp(option ...)
}
```
然后我们可以愉快的使用了。
```
func main() {

	Handler()
	
  Handler(func(options *Options) {
		options.TimeOut = time.Millisecond
	})
	
  Handler(func(options *Options) {
		options.RetryMaxNum = 1
	})
  
	Handler(func(options *Options) {
		options.RetryMaxNum = 1
	}, func(options *Options) {
		options.TimeOut = time.Millisecond
	})
}
```
是不是问题就解决了，却是这就是编程设计的魅力。

# 总结

Go语言追求精简，同时很多面向对象的特性都没有，开发中缺失面向对象是不可能的，所以只有不断的挖掘新的思想才可以。
