# rsa公钥秘钥怎么生成

使用 OpenSSL 生成 RSA 密钥对（命令行方式）

1. 生成私钥：

在命令行中，可以使用 OpenSSL 工具来生成 RSA 私钥。以下是一个基本的命令：
```bash
openssl genrsa -out private_key.pem 2048
```
这个命令的含义是：使用 genrsa（generate RSA key）命令生成一个 RSA 私钥。-out 选项指定输出文件为 private_key.pem，最后的 2048 表示密钥的长度为 2048 位。生成的文件 private_key.pem 是以 PEM（Privacy - Enhanced Mail）格式存储的私钥文件。

2. 从私钥中提取公钥：

在有了私钥之后，可以通过私钥来生成公钥。命令如下：
```bash
openssl rsa -in private_key.pem -pubout -out public_key.pem
```
这里的 rsa 命令用于处理 RSA 密钥相关操作。-in 选项指定输入文件为之前生成的私钥文件 private_key.pem，-pubout 表示输出公钥，-out 选项指定公钥的输出文件为 public_key.pem。这样就得到了一个 PEM 格式的公钥文件。
