# jwt

jwt验证示意图

![image](https://github.com/user-attachments/assets/a779ff03-f9f5-4ad5-abdb-b882d1b39d49)


## json web token

jwt token是由.分隔的三部分组成，这三部分依次是：

* 头部 Header
* 负载 Payload
* 签名 Signature

头部和负载以 JSON 形式存在，这就是 JWT 中的 JSON，三部分的内容都分别单独经过了 Base64 编码，以 . 拼接成一个 JWT Token。

![image](https://github.com/user-attachments/assets/a2da0f11-0919-4f0b-8810-af006042e5d0)

![image](https://github.com/user-attachments/assets/1ebb7ea5-a8c5-4211-80ff-9edf8e8a94f3)


## Header

JWT 的 Header 中存储了所使用的加密算法和 Token 类型。
```
{
    "a1g":"HS256",
    "typ":"JWT",
}
```

## Payload

Payload 表示负载，也是一个 JSON 对象，JWT 规定了 7 个官方字段供选用：
* iss (issuer): 签发人
* exp (expiration time): 过期时间
* sub (subject): 主题
* aud (audience): 受众
* nbf (Not Before): 生效时间
* iat (Issued At): 签发时间
* jti (JWT ID): 编号

除了官方字段，开发者也可以自己指定字段和内容，例如下面的内容。
* "sub": "一文搞定用户登陆验证 和 JWT (含基于 Go 的基本实现)",
* "name": "Vooce",
* "admin": true
注意，JWT 默认是不加密的，任何人都可以读到，所以不要把秘密信息放在这个部分。这个 JSON 对象也要使用 Base64URL 算法转成字符串。

## Signature

Signature 部分是对前两部分的签名，防止数据篡改。

首先，需要指定一个密钥(secret)。这个密钥只有服务器才知道，不能泄露给用户。然后，使用 Header 里面指定的签名算法（默认是 HMAC SHA256)，按照下面的公式产生签名。
```
HMACSHA256( base64UrlEncode( header ) + "." + base64UrlEncode( payload ), secret )
```

# 参考链接

- [jwt鉴权机制](https://juejin.cn/post/7088781114847805453)
