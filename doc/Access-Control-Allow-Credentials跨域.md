# Access-Control-Allow-Credentials跨域


前后端分离的项目中肯定会碰到跨域的问题，究其原因还是为了安全。我在一个前端工程调试过程中发现，即使我后端已经允许了跨域，但是前端依然报一个跨域错误。
```
Access to XMLHttpRequest at 'http://localhost/api/admin/authorizations' from origin 'http://localhost:9528' has been blocked by CORS policy: Response to preflight request doesn't pass access control check: The value of the 'Access-Control-Allow-Credentials' header in the response is '' which must be 'true' when the request's credentials mode is 'include'. The credentials mode of requests initiated by the XMLHttpRequest is controlled by the withCredentials attribute.
```

尝试了很多网上的方法也都没有弄清原因在哪里。索性就仔细研究一下 Access-Control-Allow-Credentials 这个头的作用，果然药到病除。这个是服务端下发到客户端的 response 中头部字段，意义是允许客户端携带验证信息，例如 cookie 之类的。这样客户端在发起跨域请求的时候，不就可以携带允许的头，还可以携带验证信息的头，又由于客户端是请求框架是 axios，并且手残的设置了 withCredentials: true，意思是客户端想要携带验证信息头，但是我的服务端设置是 'supportsCredentials' => false, ，表示不允许携带信息头，好了，错误找到了。

我们的客户端和服务端交互的时候使用的是 token，通过 Authorization头发送到服务端，并没有使用到 cookie，所以客户端没有必要设置 withCredentials: true，一顿操作猛如虎。

# 参考文档

* mozilla的HTTP访问控制

* 介绍了跨域中其他的头信息
