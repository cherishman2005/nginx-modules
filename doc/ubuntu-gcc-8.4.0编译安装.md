# ubuntu gcc 8.4.0编译安装

一、基础环境

需要基础工具：如make,gcc,g++,pkg-config 等，才能进行后续操作。

命令：

sudo apt update && sudo  apt install build-essential
二、下载源码

到：gnu官网 上去找源码。

gcc-8.4.0

然后解压：tar -Jxvf gcc-8.4.0.tar.xz

# 三、编译

1.使用下面的命令下载依赖库：gmp、mpc、mpfr、isl。
```
cd gcc-8.4.0 && ./contrib/download_prerequisites
```
2.创建目录        
```
mkdir build && cd build
```

3.执行configure    
```
../configure --prefix=/usr/local/gcc-8.4.0
```
4.编译
```
make -j$(nproc)
```
5.安装
```
make install
```

# 四、多版本gcc管理

当前系统gcc/g++版本为7.5.0。



 在执行了make install 后，gcc-8.4.0被安装在/usr/local/gcc-8.4.0,执行一下编译出来的gcc,版本为8.4.0。



 将新版本gcc/g++和老版本gcc/g++都加入update-alternatives管理：

使用方法： update-alternatives --install <链接> <名称> <位置> <优先级>
```
sudo update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-7 1
sudo update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-7 1
sudo update-alternatives --install /usr/bin/gcc gcc /usr/local/gcc-8.4/bin/gcc 2
sudo update-alternatives --install /usr/bin/g++ g++ /usr/local/gcc-8.4/bin/g++ 2
```
使用  update-alternatives --config  <名称>查看一个程序的多版本信息并切换。

输入编号1，使用 /usr/bin/gcc-7作为gcc的实体程序。g++切换方法相同。

注意：gcc和g++因保持相同版本。
