MySQL Download URL
```
https://dev.mysql.com/get/Downloads/MySQL-5.5/mysql-5.5.56-linux-glibc2.5-x86_64.tar.gz
```
### Open the terminal and follow along:

- Uninstall any existing version of MySQL

```
sudo rm /var/lib/mysql/ -R
```
- Delete the MySQL profile
```
sudo rm /etc/mysql/ -R
```
- Automatically uninstall mysql
```
sudo apt-get autoremove mysql* --purge
sudo apt-get remove apparmor
```
- Download version 5.5.51 from MySQL site
```
wget https://dev.mysql.com/get/Downloads/MySQL-5.5/mysql-5.5.56-linux-glibc2.5-x86_64.tar.gz
```
- Add mysql user group
```
sudo groupadd mysql
```

- Add mysql (not the current user) to mysql user group
```
sudo useradd -g mysql mysql
```

- Extract it

```
sudo tar -xvf mysql-5.5.56-linux-glibc2.5-x86_64.tar.gz
```
- Move it to /usr/local

```
sudo mv mysql-5.5.56-linux-glibc2.5-x86_64 /usr/local/
```

- Create mysql folder in /usr/local by moving the untarred folder
```
cd /usr/local
sudo mv mysql-5.5.49-linux2.6-x86_64 mysql

```

- set MySql directory owner and user group

```
cd mysql
sudo chown -R mysql:mysql *
```
- Install the required lib package (works with 5.6 as well) 

```
sudo apt-get install libaio1
```
- Execute mysql installation script

```
sudo scripts/mysql_install_db --user=mysql
```

- Set mysql directory owner from outside the mysql directory

```
sudo chown -R root .
```
- Set data directory owner from inside mysql directory

```
sudo chown -R mysql data
```
- Copy the mysql configuration file

```
sudo cp support-files/my-medium.cnf /etc/my.cnf
```
- Start mysql
```
sudo bin/mysqld_safe --user=mysql &
sudo cp support-files/mysql.server /etc/init.d/mysql.server
```
- Set root user password

```
sudo bin/mysqladmin -u root password '[your new password]'
```

- Add mysql path to the system

```
sudo ln -s /usr/local/mysql/bin/mysql /usr/local/bin/mysql
```

- Reboot!

- Start mysql server

```
sudo /etc/init.d/mysql.server start
```
- Stop mysql server

```
sudo /etc/init.d/mysql.server stop
```

- Check status of mysql
```
sudo /etc/init.d/mysql.server status
```

- Enable myql on startup

```
sudo update-rc.d -f mysql.server defaults
```

*Disable mysql on startup (Optional)

```
sudo update-rc.d -f mysql.server remove
```

- REBOOT!

- Now login using below command, start mysql server if it's not running already 


```
mysql -u root -p
```



