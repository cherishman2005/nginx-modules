# 使用stl中的advance和distance方法来进行iterator的加减

以前在遍历vector的时候，经常使用iterator之间的加减来获得元素在容器里面的index。

今天在进行list的 iterator加减的时候，发现不能编译通过，后面想起list是非线性的容器，不能加减。

查了一下资料，发现stl有提供两个iterator加减的方法：advance 和 distance

advance是 将iterator移动，而distance是计算两个iterator直接的距离。

```  
 template<typename _InputIterator>

    inline typename iterator_traits<_InputIterator>::difference_type

    distance(_InputIterator __first, _InputIterator __last)

    {
      // concept requirements -- taken care of in __distance

      return std::__distance(__first, __last,

    std::__iterator_category(__first));

    }
```

第一个参数first,第二个参数last  

返回的是基于 first + n  = last 中的 n


也就是基于，last 是由first经过n可以到达的，注意n可能为负的。

我第一次使用这个就是将两个参数顺序弄反了，结果返回负的n，而我这个是对应一个数字的Index下标，结果导致找不到对应的index的元素。

例子代码如下：
```
std::vector<STableInfor>::iterator iter_begin = m_tableInfoList.begin();
std::vector<STableInfor>::iterator iter_end = m_tableInfoList.end();
int index = distance(iter_begin,iter_end);
```
