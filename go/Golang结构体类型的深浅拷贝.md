# Golang结构体类型的深浅拷贝

在开发中会经常的把一个变量复制给另一个变量，有时候为了不让对象直接相互干扰，我们需要深度赋值对象

那么这个过程，可能涉及到深浅拷贝

## 1、浅拷贝

浅拷贝是指对地址的拷贝

浅拷贝的是数据地址，只复制指向的对象的指针，此时新对象和老对象指向的内存地址是一样的，新对象值修改时老对象也会变化，释放内存地址时，同时释放内存地址

引用类型的都是浅拷贝：slice、map、function

浅拷贝的特点：

拷贝的时候仅仅拷贝地址，地址指向的都是同一个值

在a中修改，则b中也跟着变化

内存销毁是一致的

## 2、深拷贝

深拷贝是指将地址指向的值进行拷贝

深拷贝的是数据本身，创造一个一样的新对象，新创建的对象与原对象不共享内存，新创建的对象在内存中开辟一个新的内存地址，新对象值修改时不会影响原对象值。既然内存地址不同，释放内存地址时，可分别释放

值类似的都是深拷贝：int、float、bool、array、struct

深拷贝的特点：

复制的时候会新创建一个对象
指向完全不同的内存地址
修改是互不影响的
通过指针求值，将值拷贝实现，修改拷贝的值不影响原来的值
```
type Author struct {
	Name string
	Aage int
}

type Title struct {
	Main string
	Sub  string
}

type Book2 struct {
	Author *Author
	Title  *Title
}

func (b *Book2) GetName() string {
	return b.Author.GetName() + "book"
}

func TestMain8(t *testing.T) {
	b1 := Book2{
		Author: &Author{
			Name: "old author",
		},
		Title: &Title{},
	}

	b2 := &Book2{
		Author: &Author{},
		Title: &Title{},
	}
	*b2.Author = *b1.Author
	*b2.Title = *b1.Title

	b2.Author.Name = "new author"
	fmt.Println(b1.Author.Name)  // old author
	fmt.Println(b2.Author.Name)  // new author
}
```

## 3、结构体的深拷贝

默认情况下，结构体类型中的字段是值类型，拷贝时都是深拷贝
```
type Per struct {
	Name     string
	Age      int
	HouseIds [2]int
}

func main()  {
	p1 := Per{
		Name:     "ssgeek",
		Age:      24,
		HouseIds: [2]int{22, 33},
	}
	p2 := p1
	fmt.Printf("%v %p \n", p1, &p1)  // {ssgeek 24 [22 33]} 0xc000180030
	fmt.Printf("%v %p \n", p2, &p2)  // {ssgeek 24 [22 33]} 0xc000180060
	p2.Age = 19
	p2.Name = "likui"
	p2.HouseIds[1] = 44
	fmt.Printf("%v %p \n", p1, &p1)  // {ssgeek 24 [22 33]} 0xc000098180
	fmt.Printf("%v %p \n", p2, &p2)  // {likui 19 [22 44]} 0xc0000981b0
}
```

## 4、结构体的浅拷贝

使用指针进行浅拷贝，浅拷贝中，可以看到p1和p2的内存地址是相同的，修改其中一个对象的属性时，另一个也会产生变化
```
package main

import "fmt"

type Per struct {
	Name     string
	Age      int
	HouseIds [2]int
}

func main()  {
	p1 := Per{
		Name:     "ssgeek",
		Age:      24,
		HouseIds: [2]int{22, 33},
	}
	p2 := &p1
	fmt.Printf("%v %p \n", p1, &p1)  // {ssgeek 24 [22 33]} 0xc000076180
	fmt.Printf("%v %p \n", p2, p2)  // &{ssgeek 24 [22 33]} 0xc000076180
	p2.Age = 19
	p2.Name = "likui"
	p2.HouseIds[1] = 44
	fmt.Printf("%v %p \n", p1, &p1)  // {likui 19 [22 44]} 0xc000076180
	fmt.Printf("%v %p \n", p2, p2)  // &{likui 19 [22 44]} 0xc000076180
}
```

## 5、结构体值类型的浅拷贝

使用new函数实现值类型的浅拷贝

值类型的默认是深拷贝，想要实现值类型的浅拷贝，一般是两种方法

使用指针
使用new函数（new函数返回的是指针）
```
package main

import "fmt"

type Per struct {
	Name     string
	Age      int
	HouseIds [2]int
}

func main()  {
	p1 := new(Per)
	p1.HouseIds = [2]int{22, 33}
	p1.Name = "songjiang"
	p1.Age = 20

	p2 := p1
	fmt.Printf("%v %p \n", p1, p1)  // &{songjiang 20 [22 33]} 0xc000076180
	fmt.Printf("%v %p \n", p2, p2)  // &{songjiang 20 [22 33]} 0xc000076180
	p2.Age = 19
	p2.Name = "likui"
	p2.HouseIds[1] = 44
	fmt.Printf("%v %p \n", p1, p1)  // &{likui 19 [22 44]} 0xc000076180
	fmt.Printf("%v %p \n", p2, p2)  // &{likui 19 [22 44]} 0xc000076180
}
```

## 6、结构体引用类型的浅拷贝

结构体默认是深拷贝，但如果结构体中包含map、slice等这些引用类型，默认也还是浅拷贝

map是引用类型，引用类型浅拷贝是默认的情况
```
package main

import "fmt"

type Per struct {
	Name     string
	Age      int
	HouseIds [2]int  // 数组，指定了长度
	CarIds   []int  // 切片，没指定长度
	Labels   map[string]string
}

func main() {
	//p1 := new(Per)
	p1 := Per{
		Name:     "ssgeek",
		Age:      24,
		HouseIds: [2]int{22, 33},  // 数组，指定了长度，深拷贝
		CarIds: []int{911, 718},  // 切片，引用类型，浅拷贝
		Labels:   map[string]string{"k1": "v1", "k2": "v2"},
		// 上述三个都是值类型，深拷贝，这个map是引用类型，浅拷贝
	}
	p2 := p1
	fmt.Printf("%v %p \n", p1, &p1) // {ssgeek 18 [22 33] map[k1:v1 k2:v2]} 0xc000076180
	fmt.Printf("%v %p \n", p2, &p2) // {ssgeek 18 [22 33] map[k1:v1 k2:v2]} 0xc0000761e0
	p2.Age = 19
	p2.Name = "likui"
	p2.HouseIds[1] = 44
	p2.CarIds[0] = 119
	p2.Labels["k1"] = "m1"
	fmt.Printf("%v %p \n", p1, &p1) // {ssgeek 24 [22 33] map[k1:m1 k2:v2]} 0xc000076180
	fmt.Printf("%v %p \n", p2, &p2) // {likui 19 [22 44] map[k1:m1 k2:v2]} 0xc0000761e0
}
```

## 7、结构体引用类型的深拷贝

结构体中含有引用类型的字段，那么这个字段就是浅拷贝，但是往往希望的是深拷贝，解决方案如下

### 方法一：挨个把可导致浅拷贝的引用类型字段自行赋值

赋值后，修改值就相互不影响了
```
package main

import "fmt"

type Per struct {
	Name     string
	Age      int
	HouseIds [2]int  // 数组，指定了长度
	CarIds   []int  // 切片，没指定长度
	Labels   map[string]string
}

func main() {
	//p1 := new(Per)
	p1 := Per{
		Name:     "ssgeek",
		Age:      24,
		HouseIds: [2]int{22, 33},  // 数组，指定了长度，深拷贝
		CarIds: []int{911, 718},  // 切片，引用类型，浅拷贝
		Labels:   map[string]string{"k1": "v1", "k2": "v2"},
		// 上述三个都是值类型，深拷贝，这个map是引用类型，浅拷贝
	}
	p2 := p1
	// 切片赋值到新的切片
	tmpCarIds := make([]int, 0)
	for _, c := range p1.CarIds{
		tmpCarIds = append(tmpCarIds, c)
	}
	// map赋值到新的map
	tmpLabels := make(map[string]string)
	for k, v := range p1.Labels{
		tmpLabels[k] = v
	}
	p2.CarIds = tmpCarIds
	p2.Labels = tmpLabels
	fmt.Printf("%v %p \n", p1, &p1) // {ssgeek 24 [22 33] [911 718] map[k1:v1 k2:v2]} 0xc00006c050
	fmt.Printf("%v %p \n", p2, &p2) // {ssgeek 24 [22 33] [911 718] map[k1:v1 k2:v2]} 0xc00006c0a0
	p1.Age = 19
	p1.Name = "likui"
	p1.HouseIds[1] = 44
	p1.CarIds[0] = 119
	p1.Labels["k1"] = "m1"
	fmt.Printf("%v %p \n", p1, &p1) // {likui 19 [22 44] [119 718] map[k1:m1 k2:v2]} 0xc00006c050
	fmt.Printf("%v %p \n", p2, &p2) // {ssgeek 24 [22 33] [911 718] map[k1:v1 k2:v2]} 0xc00006c0a0
}
```

### 方法二：使用json或反射

简单来说：json将引用类型的数据进行dump，dump后就和原来的引用类型没有关系了

package main

import (
	"encoding/json"
	"fmt"
)

type Per struct {
	Name     string
	Age      int
	HouseIds [2]int // 数组，指定了长度
	CarIds   []int  // 切片，没指定长度
	Labels   map[string]string
}

func main() {
	p1 := Per{
		Name:     "ssgeek",
		Age:      24,
		HouseIds: [2]int{22, 33},  // 数组，指定了长度，深拷贝
		CarIds:   []int{911, 718}, // 切片，引用类型，浅拷贝
		Labels:   map[string]string{"k1": "v1", "k2": "v2"},
		// 上述三个都是值类型，深拷贝，这个map是引用类型，浅拷贝
	}
	data, _ := json.Marshal(p1)
	var p2 Per
	json.Unmarshal(data, &p2)
	fmt.Printf("%v %p \n", p1, &p1) // {ssgeek 24 [22 33] [911 718] map[k1:v1 k2:v2]} 0xc00006c050
	fmt.Printf("%v %p \n", p2, &p2) // {ssgeek 24 [22 33] [911 718] map[k1:v1 k2:v2]} 0xc00006c140
	p1.Age = 19
	p1.Name = "likui"
	p1.HouseIds[1] = 44
	p1.CarIds[0] = 119
	p1.Labels["k1"] = "m1"
	fmt.Printf("%v %p \n", p1, &p1) // {likui 19 [22 44] [119 718] map[k1:m1 k2:v2]} 0xc00006c050
	fmt.Printf("%v %p \n", p2, &p2) // {ssgeek 24 [22 33] [911 718] map[k1:v1 k2:v2]} 0xc00006c140
}

### 方法三：使用其他三方库（这里还没深入）
