# golang优雅且安全的并发写slice或map

并发写 slice和map是不安全的
验证map&slice并发不安全

怎么能高并发写入？
slice&map并发不安全，都是在写的时候发生的，那么就要保证同一时间只有一个gorutine来写这个slice或者map。

有两种方式实现：

## 加锁方式

n个goroutine都有可能执行写入操作，保证同一时间只能有一个在执行写操作。 加锁操作简单，适用于性能要求低和逻辑不复杂的场景。
```
package main

import (
	"fmt"
	"sync"
)

func main() {
	slc := []int{}

	n := 10000
	var wg sync.WaitGroup
	var lock sync.Mutex
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(a int) {
			lock.Lock()
			slc = append(slc, a)
			lock.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println("done len:", len(slc))
}
```

## Active Object方式

只有1个goroutine在执行写操作。避免多个goroutine竞争锁。 适合业务场景复杂，性能要求高的场景。
```
package main

import (
	"fmt"
	"sync"
)

// active object对象
type Service struct {
	channel chan int `desc:"即将加入到数据slice的数据"`
	data    []int    `desc:"数据slice"`
}

// 新建一个size大小缓存的active object对象
func NewService(size int, done func()) *Service {
	s := &Service{
		channel: make(chan int, size),
		data:    make([]int, 0),
	}

	go func() {
		s.schedule()
		done()
	}()
	return s
}

// 把管道中的数据append到slice中
func (s *Service) schedule() {
	for v := range s.channel {
		s.data = append(s.data, v)
	}
}

// 增加一个值
func (s *Service) Add(v int) {
	s.channel <- v
}

// 管道使用完关闭
func (s *Service) Close() {
	close(s.channel)
}

// 返回slice
func (s *Service) Slice() []int {
	return s.data
}

func main() {

	// 1. 新建一个active object, 并增加结束信号
	c := make(chan struct{})
	s := NewService(100, func() { c <- struct{}{} })

	// 2. 起n个goroutine不断执行增加操作
	n := 10000
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(a int) {
			s.Add(a)
			wg.Done()
		}(i)
	}
	wg.Wait()
	s.Close()

	<-c

	// 3. 校验所有结果是否都被添加上
	fmt.Println("done len:", len(s.Slice()))
}
```
