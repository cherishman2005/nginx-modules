# ebpf

## Debian下libbpf编译失败提示<asm/types.h>文件不存在解决方法

最近在编译libbpf时候遇到了几个问题

第一是我clang安装了clang-12，但是Makefile默认写的是clang，需要更改
其次就是<asm/types.h>提示不存在

报错如下
```txt
In file included from ../../headers/linux/bpf.h:11:
/usr/include/linux/types.h:5:10: fatal error: 'asm/types.h' file not found
#include <asm/types.h>
         ^~~~~~~~~~~~~
1 error generated.
```

## 解决方案

第一个问题很好解决，添加alternative即可
```shell
update-alternatives --install /usr/bin/llc llc /usr/bin/llc-12 100
update-alternatives --install /usr/bin/clang clang /usr/bin/clang-12 100
```

完美解决clang版本问题

第二个问题比较麻烦，我一开始看的CSDN，发现CSDN纯放屁，教的方法不如不教（他让我直接改代码，几百个文件改你老母？写教程之前经过点脑子好不好）

最后去翻了翻libbpf的mailing lists

最后发现官方给出的解决方案是：
```shell
apt-get install -y gcc-multilib
```
果然很复杂的问题解决方案往往很简单
