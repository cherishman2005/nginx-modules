# 切片for循环删除元素


## 1.1 需求

a := []int{1, 2, 3, 4, 5}，slice 删除大于 3 的数字

### 2. 解决

#### 方法1

```
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		if a[i] > 3 {
			a = append(a[:i], a[i+1:]...)
			i--
		}
	}
	fmt.Println(a)
}
```

#### 方法2

```
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	j := 0

	for _, v := range a {
		if v <= 3 {
			a[j] = v
			j++
		}
	}
	a = a[:j]
	fmt.Println(a)
}
```

#### 方法3

```
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	j := 0
	q := make([]int, len(a))
	for _, v := range a {
		if v <= 3 {
			q[j] = v
			j++
		}
	}
	q = q[:j] // q is copy with numbers >= 0
	fmt.Println(q)
}
```
