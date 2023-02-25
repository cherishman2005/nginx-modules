# nodejs解析yaml

* 1、可以看到，解析后的 theJson 变量，就是整个配置文件的 json，可以直接使用其中的字段。
* 2、使用 yamljs 解析时，参数的值可以为 null 或 NULL。这点与 yaml-cpp 库不一样。
* 3、如果字段不存在时，得到的结果为 undefined，并不会出现段错误。这点与 yaml-cpp 库也不一样。

# 参考链接

- [https://blog.csdn.net/subfate/article/details/111994745](https://blog.csdn.net/subfate/article/details/111994745)

