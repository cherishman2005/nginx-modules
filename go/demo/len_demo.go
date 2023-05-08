// http://mengqi.info/html/2015/201506012300-using-golang-to-count-the-number-of-characters.html

package main
 
import (
  "log"
  "unicode/utf8"
)
 

const  CONTENT_THRESHOLD int = 70

func main() {
    content := "رمز التحقق: 625391، صالح لمدة 10 دقائق، لا تكشفه لأي شخص."
    count := utf8.RuneCountInString(content)
    level := count/CONTENT_THRESHOLD
    
    size := len(content)
    log.Printf("size:%d, count:%d, level:%d", size, count, level)
}
