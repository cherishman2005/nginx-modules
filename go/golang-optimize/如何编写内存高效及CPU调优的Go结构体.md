# 如何编写内存高效及CPU调优的Go结构体

结构体是包含多个字段的集合类型，用于将数据组合为记录。这样可以将与同一实体相关联的数据利落地封装到一个轻量的类型定义中，然后通过对该结构体类型定义方法来实现不同的行为。

本文会尝试从内存利用和CPU周期的角度讲解如何高效编写struct。

我们来看下面这一结构体，这是我们一个奇怪用例所定义的terraform资源类型：
```
type TerraformResource struct {
  Cloud                string                       // 16字节
  Name                 string                       // 16字节
  HaveDSL              bool                         //  1字节
  PluginVersion        string                       // 16字节
  IsVersionControlled  bool                         //  1字节
  TerraformVersion     string                       // 16字节
  ModuleVersionMajor   int32                        //  4字节
}
```

使用如下代码来了解TerraformResource结构体需要分配多少内存：
```
package main
import (
	"fmt"
	"unsafe"
)

type TerraformResource struct {
	Cloud               string // 16字节
	Name                string // 16字节
	HaveDSL             bool   //  1字节
	PluginVersion       string // 16字节
	IsVersionControlled bool   //  1字节
	TerraformVersion    string // 16字节
	ModuleVersionMajor  int32  //  4字节
}

func main() {
  var d TerraformResource
	d.Cloud = "aws"
	d.Name = "ec2"
	d.HaveDSL = true
	d.PluginVersion = "3.64"
	d.TerraformVersion = "1.1"
	d.ModuleVersionMajor = 1
	d.IsVersionControlled = true
	fmt.Println("==============================================================\n")
	fmt.Printf("结构体使用的总内存:d %T => [%d]\n", d, unsafe.Sizeof(d))
	fmt.Println("==============================================================\n")
	fmt.Printf("结构体中的Cloud字段:d.Cloud %T => [%d]\n", d.Cloud, unsafe.Sizeof(d.Cloud))
	fmt.Printf("结构体中的Name字段:d.Name %T => [%d]\n", d.Name, unsafe.Sizeof(d.Name))
	fmt.Printf("结构体中的HaveDSL字段:d.HaveDSL %T => [%d]\n", d.HaveDSL, unsafe.Sizeof(d.HaveDSL))
	fmt.Printf("结构体中的PluginVersion字段:d.PluginVersion %T => [%d]\n", d.PluginVersion, unsafe.Sizeof(d.PluginVersion))
	fmt.Printf("结构体中的ModuleVersionMajor字段:d.IsVersionControlled %T => [%d]\n", d.IsVersionControlled, unsafe.Sizeof(d.IsVersionControlled))
	fmt.Printf("结构体中的TerraformVersion字段:d.TerraformVersion %T => [%d]\n", d.TerraformVersion, unsafe.Sizeof(d.TerraformVersion))
	fmt.Printf("结构体中的ModuleVersionMajor字段:d.ModuleVersionMajor %T => [%d]\n", d.ModuleVersionMajor, unsafe.Sizeof(d.ModuleVersionMajor))
}
```
输出结果
```
$ go run golang-struct-memory-allocation.go 
==============================================================
结构体使用的总内存:d main.TerraformResource => [88]
==============================================================
结构体中的Cloud字段:d.Cloud string => [16]
结构体中的Name字段:d.Name string => [16]
结构体中的HaveDSL字段:d.HaveDSL bool => [1]
结构体中的PluginVersion字段:d.PluginVersion string => [16]
结构体中的ModuleVersionMajor字段:d.IsVersionControlled bool => [1]
结构体中的TerraformVersion字段:d.TerraformVersion string => [16]
结构体中的ModuleVersionMajor字段:d.ModuleVersionMajor int32 => [4]
```

因此结构体TerraformResource所需分配的总内存是88字节。TerraformResource类型内存分配如下图所示：


![image](https://user-images.githubusercontent.com/17688273/203674120-e0dfd4d7-e9b9-4a33-bd72-c12553cc6206.png)



为什么是88字节呢？16 +16 + 1 + 16 + 1+ 16 + 4 = 70 bytes，多出来的18字节是从哪来的？

涉及到结构体的内存分配时，总是会分配连续、字节对齐的内存，字段按所定义的顺序进行内存分配和存储。这里的字节对齐表示连续的内存块按平台的字大小进行偏移排列。

![image](https://user-images.githubusercontent.com/17688273/203674165-568fc530-5099-4bc2-b728-9fa96d4a51e2.png)


可以很清楚地看到TerraformResource.HaveDSL、TerraformResource.isVersionControlled和TerraformResource.ModuleVersionMajor分别仅占用1字节、1字节和4字节。剩余的空间使用空白字节进行填充。

所以重新计算一下：

数据占用字节 = 16字节 + 16字节 + 1字节 + 16字节 + 1字节 + 16字节 + 4字节 = 70字节

空白字节 = 7字节 + 7字节 + 4字节 = 18字节

总字节数 = 数据占用字节 + 空白字节 = 70字节 + 18字节 = 88字节

那如何修复这个问题呢？通过恰当地的数据结构对齐，我们可以这样来定义结构体：
```
type TerraformResource struct {
	Cloud               string // 16字节
	Name                string // 16字节
	PluginVersion       string // 16字节
	TerraformVersion    string // 16字节
	ModuleVersionMajor  int32  //  4字节
	HaveDSL             bool   //  1字节
	IsVersionControlled bool   //  1字节
}
```

使用优化后的结构体来运行同一段代码：

输出结果
```
$ go run golang-struct-memory-allocation.go 
==============================================================
结构体使用的总内存:d main.TerraformResource => [72]
==============================================================
结构体中的Cloud字段:d.Cloud string => [16]
结构体中的Name字段:d.Name string => [16]
结构体中的HaveDSL字段:d.HaveDSL bool => [1]
结构体中的PluginVersion字段:d.PluginVersion string => [16]
结构体中的ModuleVersionMajor字段:d.IsVersionControlled bool => [1]
结构体中的TerraformVersion字段:d.TerraformVersion string => [16]
结构体中的ModuleVersionMajor字段:d.ModuleVersionMajor int32 => [4]
现在TerraformResource类型总的内存占用是72字节。我们来看下在内存中是如何排列的：
```


![image](https://user-images.githubusercontent.com/17688273/203674310-bde71435-a5b5-44eb-bb17-7536d0cac778.png)


仅仅是通过对结构体元素进行了一轮数据结构对齐我们就将所占用的内存由88字节降到了72字节，真是太棒了！！！

我们再来算一下

数据占用字节 = 16字节 + 16字节 + 16字节 + 16字节 +4字节 + 1 byte + 1字节 = 70字节

空白字节 = 2字节

总字节数 = 数据占用字节 + 空白字节 = 70字节 + 2字节 = 72字节

通过恰当的数据结构对齐不仅优化了内存占用，还优化了CPU读取周期，怎么做到的呢？

CPU以字为单位从内存中进行读取，一个字在32位系统中占用4字节、64位系统中占用8字节。我们声明的第一个结构体类型TerraformResourceCPU需要读取11个字才能读完：

![image](https://user-images.githubusercontent.com/17688273/203674471-77b9a6d1-8635-4b35-a970-d62036de88e5.png)



但对优化后的结构体只需要读取9个字:

![image](https://user-images.githubusercontent.com/17688273/203674519-c4864940-302b-4d14-b251-c12191d5b2ce.png)


通过恰当地对结构体进行数据结构排序我们可以让内存分配和CPU 读取都变得高效。

这只是一个小例子，试想一个带有20或30个不同类型字段的大型结构体。有计划的数据结构对齐一定是有回报的

希望这篇文章可以让读者更加了解结构体的内部原理、内存分配和所需的CPU读取周期。但愿有所帮助！！

内容整理自：https://medium.com/towardsdev/golang-writing-memory-efficient-and-cpu-optimized-go-structs-62fcef4dbfd0

# 参考链接

- [如何编写内存高效及CPU调优的Go结构体](https://mp.weixin.qq.com/s/LYDYiO8osawOMvpSCESQtg)
