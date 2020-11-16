

# FAQ

## SQL 注入
在写 SQL 语句的时间尽量不要使用 SQL 拼装，因为很容易被 SQL注入，从而引发安全问题，如果数据和 SQL 语句需要分离，那么请使用 占位符 的方式。

```
connection.query("select * from users where id = ? and name = ?", [1, 'jmjc'], (err, result)=>{}) // 这种方式 mysql 模块内部会调用 escape 方法，过滤掉一些非法的操作

/*
  当前我们也可以自己使用 escape 方法
*/
connection.query('select * from users where id = ' + connection.escape(userId), (err, result) => {})

/*
 或者 format 方法
*/
let sql = "select * from ?? where ?? = ?"
const inserts = ['users', 'id', 1]
sql = mysql.format(sql, inserts) // select * from users where id = 1
```

# 参考链接

- [https://www.jmjc.tech/less/113](https://www.jmjc.tech/less/113)