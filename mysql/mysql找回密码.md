# mysql找回密码

1. 放开mysqld.cnf最后的`skip-grant-tables`
```
/etc/mysql/mysql.conf.d# vi mysqld.cnf
```

2. 重启mysql
```
sudo service mysql restart
```

3. 更新密码
```
# mysql
use mysql;
update user set authentication_string=PASSWORD("new-password") where user='root';
```

4. 还原mysqld.cnf配置，重启mysql
```
sudo service mysql restart
```
