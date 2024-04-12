# golang-msyql匹配

```
package main

import (
    "fmt"
    "regexp"
)

func main() {
    // 定义要匹配的字符串
    input := "INSERT IGNORE INTO table_name (column1, column2) VALUES (value1, value2);"

    // 编译正则表达式，设置IgnoreCase为true以实现不区分大小写的匹配
    re, err := regexp.Compile(`\bINSERT\s+IGNORE\s+INTO\b`, regexp.IgnoreCase)
    if err != nil {
        fmt.Println("Error compiling regex:", err)
        return
    }

    // 使用编译好的正则表达式进行匹配
    match := re.MatchString(input)
    if match {
        fmt.Println("Matched:", input)
    } else {
        fmt.Println("Not matched:", input)
    }
}
```
在上面的代码中，`\b是一个单词边界，确保我们匹配的是完整的单词而不是其他单词的一部分。\s+匹配一个或多个空白字符，这样可以适应不同的空格风格。IgnoreCase选项确保匹配时不区分大小写。`
