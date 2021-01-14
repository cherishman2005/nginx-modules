package main  
  
import (  
    "fmt"  
    "strings"  
)  


func TrimSpace1(str string) string {

	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除table
	str = strings.Replace(str, "\t", "", -1)  
	// 去除换行符
	str = strings.Replace(str, "\r\n", "", -1)  
	str = strings.Replace(str, "\n", "", -1)  

	return str
}

func TrimSpace(str string) string {
	// 去除空格
	str = strings.Trim(str, " ")
	// 去除table
	str = strings.Trim(str, "\t")
	// 去除换行符
	str = strings.Trim(str, "\r\n")
	str = strings.Trim(str, "\n")

	return str
}

func main() {  
    //str := "这里是 www\n.runoob\n.com"
    str := " \t\n Hello, \t Gophers \n\t\r\n"
    fmt.Println("-------- 原字符串 ----------")  
    fmt.Println(str)
    
    str = TrimSpace1(str)
    //str = strings.TrimSpace(str)

    fmt.Println("-------- 去除空格与换行后 ----------")  
    fmt.Println(str)  
}
