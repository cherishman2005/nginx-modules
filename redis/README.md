# redis

## IO多路复用原理

![image](https://github.com/user-attachments/assets/f63cd8d2-640f-442f-9234-821de9e767ae)

## redis内存不足时的处理

* 惰性删除
  * 每隔100ms 通过LRU 抽取部分过期的数据进行释放。
  * 查询数据时，如果已经过期，便释放。
 
# 小结

* redis内存管理的思想 很多golang本地缓存 均可借鉴。

