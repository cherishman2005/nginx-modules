# Mysql开启binlog日志

binlog日志文件只对 增删改有记录，查询操作是没有记录的

二进制日志文件，MySql8.0默认已经开启，低版本的MySql需要通过配置文件开启，并配置MySql日志格式，windows系统：myini，Linux系统：my.cnf

## 1. 查看是否开启binlog
```
show variables like 'log_%';
```

```
mysql> show variables like 'log_%';
+----------------------------------------+--------------------------+
| Variable_name                          | Value                    |
+----------------------------------------+--------------------------+
| log_bin                                | OFF                      |
| log_bin_basename                       |                          |
| log_bin_index                          |                          |
| log_bin_trust_function_creators        | OFF                      |
| log_bin_use_v1_row_events              | OFF                      |
| log_builtin_as_identified_by_password  | OFF                      |
| log_error                              | /var/log/mysql/error.log |
| log_error_verbosity                    | 3                        |
| log_output                             | FILE                     |
| log_queries_not_using_indexes          | OFF                      |
| log_slave_updates                      | OFF                      |
| log_slow_admin_statements              | OFF                      |
| log_slow_slave_statements              | OFF                      |
| log_statements_unsafe_for_binlog       | ON                       |
| log_syslog                             | OFF                      |
| log_syslog_facility                    | daemon                   |
| log_syslog_include_pid                 | ON                       |
| log_syslog_tag                         |                          |
| log_throttle_queries_not_using_indexes | 0                        |
| log_timestamps                         | UTC                      |
| log_warnings                           | 2                        |
+----------------------------------------+--------------------------+
```
ON开启状态，OFF关闭状态
 
## 2. 修改my.cnf配置文件  vi /etc/my.cnf

```
server-id              = 1
log_bin                        = /var/log/mysql/mysql-bin.log
expire_logs_days        = 10
max_binlog_size   = 100M
```

![image](https://github.com/cherishman2005/nginx-modules/assets/17688273/8975fd36-b445-4577-9147-ce6a92c53081)


## 3. 重启mysql服务

重启服务
```
systemctl restart mysqld
```

## 4. 再次查看binlog开启状态

```
show variables like 'log_%';
```

![image](https://github.com/cherishman2005/nginx-modules/assets/17688273/a82ca3ec-29d9-4d7c-aa33-a17f1562885d)


# FAQ

## The replication sender thread cannot start in AUTO_POSITION mode: this server has GTID_MODE = OFF instead of ON.
* 暂时忽略 -- mysql开启binlog的方法已经找到

# 参考链接

- [Mysql开启binlog日志](https://www.cnblogs.com/sportsky/p/16357479.html)

