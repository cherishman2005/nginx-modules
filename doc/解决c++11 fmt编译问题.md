# 解决unable to find string literal operator 'operator""fmt' with 'const char [15]', 编译问题

```
#define show_log(tag, fmt, arg...)   printf("[%s][%s:%d]: "fmt"\n", tag, __func__, __LINE__, ##arg)
```
C可以编译通过，而C++编译出标题错误。

说是C++11要求，当字符串跟变量连接的时候，必须fmt前后增加一个空格才行。
```
#define show_log(tag, fmt, arg...)   printf("[%s][%s:%d]: "  fmt  "\n", tag, __func__, __LINE__, ##arg)
```
