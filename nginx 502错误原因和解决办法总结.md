# nginx 502错误原因和解决办法总结

## NGINX 502错误排查
NGINX 502 Bad Gateway错误是FastCGI有问题，造成NGINX 502错误的可能性比较多。将网上找到的一些和502 Bad Gateway错误有关的问题和排查方法列一下，先从FastCGI配置入手：

1. FastCGI进程是否已经启动
2. FastCGI worker进程数是否不够

运行 netstat -anpo | grep “php-cgi” | wc -l 判断是否接近FastCGI进程，接近配置文件中设置的数值，表明worker进程数设置太少

3. FastCGI执行时间过长
根据实际情况调高以下参数值
```
fastcgi_connect_timeout 300;
fastcgi_send_timeout 300;
fastcgi_read_timeout 300;
```

4. FastCGI Buffer不够
nginx和apache一样，有前端缓冲限制，可以调整缓冲参数
```
fastcgi_buffer_size 32k;
fastcgi_buffers 8 32k;
```

5. Proxy Buffer不够
如果你用了Proxying，调整
```
proxy_buffer_size   16k;
proxy_buffers    4 16k;
```
参见：http://www.server110.com

6. https转发配置错误
正确的配置方法
```
server_name www.mydomain.com;
location /myproj/repos {
set $fixed_destination $http_destination;
if ( $http_destination ~* ^https(.*)$ )
{
  set $fixed_destination http$1;
}
proxy_set_header Host $host;
proxy_set_header X-Real-IP $remote_addr;
proxy_set_header Destination $fixed_destination;
proxy_pass http://subversion_hosts;
}
```
当然，还要看你后端用的是哪种类型的FastCGI，我用过的有php-fpm，流量约为单台机器40万PV(动态页面), 现在基本上没有碰到502。

7. php脚本执行时间过长
将php-fpm.conf的<value name="request_terminate_timeout">0s</value>的0s改成一个时间

## Nginx 413错误的排查:修改上传文件大小限制
在上传时nginx返回了413错误，查看log文件，显示的错误信息是:”413 Request Entity Too Large”, 于是在网上找了下“nginx 413错误”发现需要做以下设置：
在nginx.conf增加 client_max_body_size的相关设置, 这个值默认是1m，可以增加到8m以增加提高文件大小限制；
如果运行的是php，那么还要检查php.ini，这个大小client_max_body_size要和php.ini中的如下值的最大值一致或者稍大，这样就不会因为提交数据大小不一致出现的错误。
```
post_max_size = 8M
upload_max_filesize = 2M
```

## Nginx 400错误排查：HTTP头/Cookie过大
今天有人汇报nginx的HTTP400错误，而且这个HTTP400错误并不是每次都会出现的，查了一下发现nginx400错误是由于request header过大，通常是由于cookie中写入了较长的字符串所引起的。
解决方法是不要在cookie里记录过多数据，如果实在需要的话可以考虑调整在nginx.conf中的client_header_buffer_size(默认1k)
若cookie太大，可能还需要调整large_client_header_buffers(默认4k)，该参数说明如下：
请求行如果超过buffer，就会报HTTP 414错误(URI Too Long)
nginx接受最长的HTTP头部大小必须比其中一个buffer大，否则就会报400的HTTP错误(Bad Request)。

Nginx 502 Bad Gateway的含义是请求的PHP-CGI已经执行，但是由于某种原因（一般是读取资源的问题）没有执行完毕而导致PHP-CGI进程终止。
Nginx 504 Gateway Time-out的含义是所请求的网关没有请求到，简单来说就是没有请求到可以执行的PHP-CGI。

解决这两个问题其实是需要综合思考的，一般来说Nginx 502 Bad Gateway和php-fpm.conf的设置有关，而Nginx 504 Gateway Time-out则是与nginx.conf的设置有关。
而正确的设置需要考虑服务器自身的性能和访客的数量等多重因素。
以我目前的服务器为例子CPU是奔四1.5G的，内存1GB，CENTOS的系统，访客大概是50人左右同时在线。
但是在线的人大都需要请求PHP-CGI进行大量的信息处理，因此我将nginx.conf设置为：
```
fastcgi_connect_timeout 300s;
fastcgi_send_timeout 300s;
fastcgi_read_timeout 300s;
fastcgi_buffer_size 128k;
fastcgi_buffers 8 128k;#8 128
fastcgi_busy_buffers_size 256k;
fastcgi_temp_file_write_size 256k;
fastcgi_intercept_errors on;
```
这里最主要的设置是前三条，即
```
fastcgi_connect_timeout 300s;
fastcgi_send_timeout 300s;
fastcgi_read_timeout 300s;
```
这里规定了PHP-CGI的连接、发送和读取的时间，300秒足够用了，因此我的服务器很少出现504 Gateway Time-out这个错误。最关键的是php-fpm.conf的设置，这个会直接导致502 Bad Gateway和504 Gateway Time-out。
下面我们来仔细分析一下php-fpm.conf几个重要的参数：
php-fpm.conf有两个至关重要的参数，一个是”max_children”,另一个是”request_terminate_timeout”
我的两个设置的值一个是”40″，一个是”900″，但是这个值不是通用的，而是需要自己计算的。
计算的方式如下：
如果你的服务器性能足够好，且宽带资源足够充足，PHP脚本没有系循环或BUG的话你可以直接将”request_terminate_timeout”设置成0s。0s的含义是让PHP-CGI一直执行下去而没有时间限制。而如果你做不到这一点，也就是说你的PHP-CGI可能出现某个BUG，或者你的宽带不够充足或者其他的原因导致你的PHP-CGI能够假死那么就建议你给”request_terminate_timeout”赋一个值，这个值可以根据你服务器的性能进行设定。一般来说性能越好你可以设置越高，20分钟-30分钟都可以。由于我的服务器PHP脚本需要长时间运行，有的可能会超过10分钟因此我设置了900秒，这样不会导致PHP-CGI死掉而出现502 Bad gateway这个错误。

而”max_children”这个值又是怎么计算出来的呢？这个值原则上是越大越好，php-cgi的进程多了就会处理的很快，排队的请求就会很少。设置”max_children”也需要根据服务器的性能进行设定，一般来说一台服务器正常情况下每一个php-cgi所耗费的内存在20M左右，因此我的”max_children”我设置成40个，20M*40=800M也就是说在峰值的时候所有PHP-CGI所耗内存在800M以内，低于我的有效内存1Gb。而如果我的”max_children”设置的较小，比如5-10个，那么php-cgi就会“很累”，处理速度也很慢，等待的时间也较长。如果长时间没有得到处理的请求就会出现504 Gateway Time-out这个错误，而正在处理的很累的那几个php-cgi如果遇到了问题就会出现502 Bad gateway这个错误。

nginx中配置php fastcgi组解决莫名其妙的502 Bad Gateway错误

一般nginx搭配php都采用这样的方式：
```
location ~ \.php$ {
proxy_pass        http://localhost:9000;
fastcgi_param   SCRIPT_FILENAME   /data/_hongdou$fastcgi_script_name;
include        fastcgi_params;
}
```

这个方式只能连接到一组spawn-fcgi开启的fastcgi，在服务器负载稍高时常常出现502 bad gateway错误。

起先怀疑这是php-cgi的进程开得太少，增加后仍然有反映时常有错，偶然间发现php-cgi会报出这样的错误：

zend_mm_heap corrupted

看来是php-cgi在执行某些代码时有问题，以致于该线程中止。

在服务器上可能还会看到php-cgi进程在不断变少，估计是出现错误的php-cgi的进程自动退出了。

php的问题总是不太容易能解决，所以在nginx方面想想办法，nginx的好处是它总是能爆出一些稀奇古怪的做法出来。

在nginx的proxy中，规避莫名其妙错误的办法无非是proxy到一个upstream的服务器组中，然后配置 proxy_next_upstream，让nginx遇到某种错误码时，自动跳到下一个后端上。这样，应用服务器即使不稳定，但是在nginx后面就变成了稳定服务。想到nginx的fastcgi和proxy是一路东西，所以proxy能用的经验，移植到fastcgi也能跑得起来。

照着这个思路，用spawn-fcgi多开同样一组php进程，所不同的仅仅是端口：
```
spawn-fcgi -a 127.0.0.1 -p 9000 -u nobody -f php-cgi -C 100
spawn-fcgi -a 127.0.0.1 -p 9001 -u nobody -f php-cgi -C 100
```

然后把fastcgi的这段配置改成用upstream的方式：
```
upstream backend {
server 127.0.0.1:9000;
server 127.0.0.1:9001;
}

location ~ \.php$ {
  fastcgi_pass        backend;
  fastcgi_param   SCRIPT_FILENAME   /data/_hongdou$fastcgi_script_name;
  include        fastcgi_params;
}
```

检查配置结果正确，能跑起来；同时在服务器上netstat -n|grep 9000和grep 9001都有记录，证明连接无误；在前台查阅页面，一切运行正常。

这个配置是最简单的配置，既然能连接上upstream，那么很显然upstream的一些东西都可以拿来用，比如ip_hash、weight、max_fails等。

这样的配置在单机下不知能不能共享session，没有测试，如果有问题，可以加上ip_hash，或者配置php把session存进memcached中。

然后就是fastcgi_next_upstream的配置，nginx wiki中没有介绍到这个配置，查了一下，在nginx的CHANGES中有提到，而且出生年月是和proxy_next_upstream一样的。既然如此，那就照proxy_next_upstream一样配吧。一般按默认的值error timeout就可以工作，因为php出现502错误的异常是返回的500错误，所以我把fastcgi_next_upstream定为：
```
fastcgi_next_upstream error timeout invalid_header http_500;
```

通过这个配置，就可以基本杜绝任何时常性的500错误，出问题的几率会变小很多，如果客户反映仍然激烈，那么就多增加几组fastcgi进程。

以上配置能够杜绝由于php所引起的“莫名其妙”的时常性的502错误，同时可使nginx搭配php比从前方式更为强悍。假如nginx还是返回502错误，那这次就一定是出现服务器挂掉或其它严重问题的了。

# 小结

本篇分享主要是针对nginx-php的问题定位分析，针对其他nginx-upstream场景大多数适用。

# 参考链接

- [nginx 502错误原因和解决办法总结](https://mp.weixin.qq.com/s?__biz=MzA5Njg1OTI5Mg==&mid=2651026723&idx=1&sn=046a1422f3d4a303cc41600b04f9c312&chksm=8b5e70d0bc29f9c6dd300849d6c0ef0d339e2df092b3d7c32693e6cfc041a9b23ac6c3155aca&scene=126&sessionid=1606466547&key=38e610443a66d198d446f9f123c8c5cbfdf373603662ab47bd4b7891f35eb2e6e04d35bffcfc24d3218471dd86229af901098944c1ba0f6ab51c11df70bbed418ee1f0ceeeca5b623d7505a30047e2ed19047eb9ee5775ca86a836e495dc1f41fc1ffdfa1ab1fc38d74a22dd697b617327097cb8395c459b3e46731784267b9a&ascene=1&uin=ODkxOTg4MzIw&devicetype=Windows+10+x64&version=63000039&lang=zh_CN&exportkey=A%2BjfL4piixeJci%2FRElZHfZA%3D&pass_ticket=CspG9kOtDJxEt6ZQ37T888OzVCpRRx97aKf5nkB5EPKn1tO0uhbS8rvauJgY7mEK&wx_header=0)