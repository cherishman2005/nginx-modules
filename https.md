# https

![https原理](/img/https.png)

1. 对称秘钥（随机数）通过非对称加密进行传输；

1. 请求与响应通过对称加密，然后再传输；—— 提供效率；

1. 非对称加密也可能被劫持？
   —— 通过数字证书传递公钥，防止被劫持；

   使用公有密匙加密的消息，只有对应的私有密匙才能解开。反过来，使用私有密匙加密的消息，只有公有密匙才能解开。这样客户端在发送消息前，先用服务器的公匙对消息进行加密，服务器收到后再用自己的私匙进行解密。

# FAQ

## http预防劫持
采用https（http + ssl）传输；
   
## DNS预防劫持

采用httpDNS预防劫持

