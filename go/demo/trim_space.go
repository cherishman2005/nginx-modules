package main  
  
import (
	"fmt"
	"strings"
	"regexp"
)  

// 删除字符串中的多余空格，有多个空格时，仅保留一个空格
func DeleteExtraSpace(s string) string {
		s1 := strings.Replace(s, "\t", " ", -1)      //替换tab为空格
		regstr := "\\s{2,}"                          //两个及两个以上空格的正则表达式
		reg, _ := regexp.Compile(regstr)             //编译正则表达式
		s2 := make([]byte, len(s1))                  //定义字符数组切片
		copy(s2, s1)                                 //将字符串复制到切片
		spc_index := reg.FindStringIndex(string(s2)) //在字符串中搜索
		for len(spc_index) > 0 {                     //找到适配项
				s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...) //删除多余空格
				spc_index = reg.FindStringIndex(string(s2))            //继续在字符串中搜索
		}
		return string(s2)
}

func TrimSpace(str string) string {
		// trim left-right end
		str = strings.TrimSpace(str)
		//str = strings.Trim(str, "\n")
		str = strings.Replace(str, "\n", "", -1)        

		str = DeleteExtraSpace(str)

		return str
}

/*
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
*/

func main() {  
	//str := "这里是 www\n.runoob\n.com"
	str := " \t\n Hello, \t aaa\nGophers \n\t\r\n"
	fmt.Println("-------- 原字符串 ----------")  
	fmt.Println(str)
	
	str = TrimSpace(str)
	//str = strings.TrimSpace(str)

	fmt.Println("-------- 去除空格与换行后 ----------")  
	fmt.Println(str)  
}
