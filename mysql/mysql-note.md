# mysql

## LIMIT 1
* 在某些情况下，如果明知道查询结果只有一个，SQL语句中使用LIMIT 1会提高查询效率。

* 索引就不需要加上LIMIT 1。

  如果是根据主键查询一条记录也不需要LIMIT 1，主键也是索引。