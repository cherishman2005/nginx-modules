# yaml命名不能有下划线

在 YAML 中，变量名是不允许使用下划线（_）作为命名的，因为 YAML 规范中规定了一些特殊字符的用法。

在 YAML 中，可以使用连字符（-）或者驼峰命名法（Camel Case）来表示变量名，例如：

```Yaml
my-variable-name: value
```

或

```Yaml
myVariableName: value
```
这两种风格都是 YAML 中常见的变量命名风格。

因此，在编写 YAML 配置文件时，应该遵循 YAML 的命名规范，尽量不使用下划线作为变量名的一部分。
