# sql判断字段是否为空

sql语句条件查询时，有时会判断某个字段是否为空。

字段内容为空有两种情况

1.为null

2.为字符串的空''

语句如下：
```
select * from table where column is null or trim(column)=''
```
这样就可以排除字段内容为null、''的。

 

判断某个字段不为空
```
select * from table where trim(column) != ''
```

曾经尝试判断null：is not null但是不起作用，放弃。直接 trim(column) != '' 就能解决。
