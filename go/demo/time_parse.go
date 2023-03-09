package main
import (
    "time"
    "fmt"
)

func main() {
	//t := time.Now().Unix()
	//d := time.Unix(t, 0).Format("2006-01-02 15:04:05")
	//tt, _ := time.Parse("2006-01-02 15:04:05", d)
	//fmt.Println("===========================")
	//fmt.Println("当前时间戳:", t)
	//fmt.Println("当前日期:", d)
	//fmt.Println("从日期得到时间戳:", tt.Unix())
	//fmt.Println("再次转化为日期:", time.Unix(tt.Unix(), 0).Format("2006-01-02 15:04:05"))
	//fmt.Println("===========================")
	
	
	expireTime := "2023-03-08 14:59:32"
	expire, _ := time.Parse("2006-01-02 15:04:05", expireTime)
	fmt.Println("expire:", expire.Unix())
	
	expire_ok, _ := time.ParseInLocation("2006-01-02 15:04:05", expireTime, time.Local)
	fmt.Println("expire_ok:", expire_ok.Unix())
	
}