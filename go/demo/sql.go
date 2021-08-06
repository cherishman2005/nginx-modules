package main

import (
	"fmt"
	"strings"
)


func main() {
    query := "select appid, stream_name, type, mix, rate, json from stream_infos where appid=15013 and type=2 and mix=1 and uid=2314776959 and sid=36569435"
    
    strTrim := strings.TrimSpace(query)
	strSplit := strings.SplitN(strTrim, " ", 2)
	if len(strSplit) <= 0 {
		return 
	}
    
    fmt.Printf("MysqlDBQuery strSplit:%v \n", strSplit)
    
	key := strSplit[0]
	fmt.Printf("MysqlDBQuery key:%v \n", key)
}