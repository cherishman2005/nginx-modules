# C++ map遍历删除元素

​ 今天看到一个patch fix从std::map中遍历删除元素导致crash问题，突然意识到自己对如何正确地从map等C++容器中删除元素也没有很牢固清醒的认知。重新梳理了下这块的正确做法，记录在此，以备后忘。

基础知识
​ C++的容器按存储方式分为两类：

以数组形式存储的顺序容器，如：vector，deque
以不连续节点形式存储的容易，如：list, set, map
在使用erase方法遍历删除元素时，需要注意一些问题，否则就会踩坑。

对容器进行增删元素操作，可能会使迭代器失效。如果一个元素已经被删除，则其对应的迭代器会失效，不应该再被使用；否则会导致程序无定义的行为，基本上就会挂了。

## 正确的遍历删除方法
对于遍历删除map、list、set可以使用下面2种正确方法：

1. 使用删除元素之前的迭代器定义下一个元素，建议使用的方式
```
for(auto it=mymap.begin(); it!=mymap.end();) {
    if (it->first == target) {
        mymap.erase(it++); //here is the key
    } else {
        it++;
    }
}
```
2. 使用erase()返回下一个元素的迭代器
```
for(auto it=mymap.begin(); it!=mymap.end();) {
    if (it->first == target) {
        it = mymap.erase(it);
    } else {
        it++;
    }
}
```
注意：在对 vector、deque遍历删除元素时，可以通过erase的返回值来获取下一个元素的位置，也就是上面的第2种方法；但不能使用上面的第1种方法来遍历删除。

## 错误的遍历删除方法
把经常会踩坑的错误的写法贴在下面，作为警示！

下面的写法是错误的！下面的写法是错误的！下面的写法是错误的！
```
for(auto it=mymap.begin(); it!=mymap.end(); it++) {
    if (it->first == target) {
        mymap.erase(it); //这里的写法是错误的，错误的，错误的！！！
        //it对应的元素已经被删除，it迭代器失效，在for循环中执行it++会导致未定义行为
    }
}
```
下面的写法对vector是错误的！下面的写法对vector是错误的！下面的写法对vector是错误的！
```
for(auto it=myvector.begin(); it!=myvector.end();) {
    if (*it == target) {
        myvector.erase(it++); //对vector不能工作
    } else {
        it++;
    }
}
```
