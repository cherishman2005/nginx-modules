# mysql

## mysql导出到本地的指令
```
echo 'select ' |mysql > txt
```
然后将txt贴到excel。


## 数据库表删除操作

一定要谨慎，特别是敲命令时，尽量把其他写完了，最后写delete。

## mysql开启binlog日志的方法

- [Mysql开启binlog日志](./Mysql开启binlog日志.md)

* 用途： 订阅mysql binlog，向其他IDC做同步备份，或者做向其他数据库（如es，kafka等）做同步。

# FAQ

## SQL优化

建议按优先级依次介绍慢 SQL 定位、索引优化、表结构设计和 SQL 优化等内容。架构层面的优化，如读写分离和分库分表、数据冷热分离 应作为最后的手段，除非在特定场景下有明显的性能瓶颈，否则不应轻易使用，因其引入的复杂性会带来额外的维护成本。
