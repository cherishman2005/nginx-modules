# 报错：[Error] ‘memset’ was not declared in this scope问题解决

## memset

memset是C 库函数， void *memset(void *str, int c, size_t n) 复制字符 c（一个无符号字符）到参数 str 所指向的字符串的前 n 个字符。

所以在调用时需要注意

如果是C语言编译的话，直接调用string库

如果是C++编译，需要调用的是cstring。

报错：Error] 'memset' was not declared in this scope是因为我在C++环境下使用了memset却没有导入cstring，所以解决办法很简答，加上

```
#include <cstring>
```
