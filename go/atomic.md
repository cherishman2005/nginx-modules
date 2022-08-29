在Go语言中，原子包提供lower-level原子内存，这对实现同步算法很有帮助。 Go语言中的SwapUint32()函数用于将新值自动存储到* addr中，并返回先前的* addr值。此函数在原子包下定义。在这里，您需要导入“sync/atomic”软件包才能使用这些函数。

用法:
```
func SwapUint32(addr *uint32, new uint32) (old uint32)
```
在此，addr表示地址。而new是新的uint32值，而old是旧的uint32值。

注意：(* uint32)是指向uint32值的指针。但是，int32包含从0到4294967295的所有无符号32位整数的集合。

返回值：它将新的uint32值存储到* addr中，并返回先前的* addr值。
