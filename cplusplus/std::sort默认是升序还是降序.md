# std::sort默认是升序还是降序

std::sort 默认是按照升序进行排序的。这意味着元素将按照从小到大的顺序排列。

如果你想要按照降序进行排序，可以自定义排序方式，例如传入一个比较函数或 lambda 表达式，来告诉 std::sort 按照你希望的方式进行排序。例如，你可以使用以下方式使 std::sort 按照降序排序：

```Cpp
std::sort(vec.begin(), vec.end(), std::greater<int>());
```
或者使用 lambda 表达式：

```Cpp
std::sort(vec.begin(), vec.end(), [](int a, int b) {
    return a > b;
});
```
以上代码将告诉 std::sort 按照降序对容器中的元素进行排序。
