package main

import (
	"fmt"
	"sync"
)

/*
sync.Map主要是由 互斥锁Mutex 和 map组合而成；并且做到读写时部分逻辑绕过Mutex互斥锁
*/

func main() {
	var m sync.Map

	// 存储键值对
	m.Store("apple", 1)
	m.Store("banana", 2)
	m.Store("cherry", 3)

	// 加载键值对
	if value, ok := m.Load("apple"); ok {
		fmt.Println("Value of apple:", value)
	}

	// 加载或存储
	actual, loaded := m.LoadOrStore("apple", 4)
	fmt.Println("Actual value of apple:", actual)
	fmt.Println("Was apple already loaded?", loaded)

	// 删除键值对
	m.Delete("banana")

	// 遍历键值对
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
		return true
	})
}
