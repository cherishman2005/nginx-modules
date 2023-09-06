# go作用域

```
package main

import (
    "log"
	"strconv"
)

func getinfo() (appid, servanme, ip string) {
    return "123", "sacheck", "127.0.0.1"
}

func fn1() (err error) {
    var servername string
	defer func() {
	    log.Printf("servername:%s", servername)
	}()
	
    atype, err := strconv.ParseInt("", 10, 64)
    log.Printf("atype:%d, err:%v", atype, err)
	
	appid,servername,ip := getinfo()
	log.Printf("appid:%s, servername:%s, ip:%s", appid, servername, ip)
	
    return err
}

func main() {
    err := fn1()
    log.Print("main, err:", err)
}
```

* 运行结果
![D8B38C3EB0A99E4EDDDECD08332E8C5D](https://github.com/cherishman2005/nginx-modules/assets/17688273/098e49b5-9ce3-4e41-9ee7-96333d947473)

