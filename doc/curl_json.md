# curl json

```
curl 'http://localhost:18080/test-sinfo?action=getStreamNames&appid=15013&sid=2789874956' | json_pp
```

json输出：
```
{
   "stream_names" : {
      "xa_36569435_2789874956_0_0_0_1" : true,
      "xv_36569435_2789874956_0_10_0_1" : true,
      "xv_36569435_2789874956_0_0_0_1" : true,
      "xv_36569435_2789874956_0_100_0_1" : true,
      "xa_36569435_2789874956_0_10_0_1" : true,
      "v_36569435_2789874956_2314776959_0_0" : true,
      "xv_36569435_2789874956_0_1_0_1" : true,
      "a_36569435_2789874956_2314776959_0_0" : true
   },
   "appid" : 15013,
   "code" : 0,
   "sid" : 2789874956,
   "action" : "getStreamNames",
   "msg" : null
}
```

# 参考链接

- [https://mkyong.com/web/how-to-pretty-print-json-output-in-curl/](https://mkyong.com/web/how-to-pretty-print-json-output-in-curl/)
