# std::cerr 和 stderr

```
std::cerr << str;

std::fprintf(stderr, "%s", str.c_str());
std::fflush(stderr);
```

以上代码展示了在 C++ 中输出字符串到标准错误流 stderr 的两种方式。

std::cerr << str;：这是 C++ 标准库提供的一种输出方式，它会将字符串 str 输出到标准错误流 cerr 中。cerr 是一个异步流，可用于向屏幕或日志文件输出错误信息。这种方式比较简洁，但是不够灵活，无法像 fprintf 那样进行格式化输出。

std::fprintf(stderr, "%s", str.c_str()); 和 std::fflush(stderr);：这是使用标准 C 库函数 fprintf 将字符串 str 格式化输出到标准错误流 stderr 中，fflush 函数用于刷新 stderr 流。这种方式可以进行更加自由的格式化输出，可以按照指定的格式输出字符串。需要注意的是，c_str() 函数用于获取 C 风格的字符串指针。

根据以上描述，两种方式的选择取决于实际需求。如果只是简单地输出字符串，可以使用第一种方式，即 std::cerr << str;。如果需要格式化输出或更灵活的控制输出格式，则可以选择使用第二种方式，即 fprintf 函数。
