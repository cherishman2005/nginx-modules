# mysql

## LIMIT 1
* 在某些情况下，如果明知道查询结果只有一个，SQL语句中使用LIMIT 1会提高查询效率。

* 索引就不需要加上LIMIT 1。

  如果是根据主键查询一条记录也不需要LIMIT 1，主键也是索引。


## Error 1047: Prepare unsupported!

golang-mysql编写时发现的bug

db.Query("select ?", 1); interpolateparams=true only works if the parameter follows the query in that same call.

- [https://github.com/go-sql-driver/mysql/issues/360](https://github.com/go-sql-driver/mysql/issues/360)
