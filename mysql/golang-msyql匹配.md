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


## FindInsertIgnorePosition

```
// FindInsertIgnorePosition 查找SQL语句中"insert ignore into"的起始和结束位置
func FindInsertIgnorePosition(sql string) (start, end int) {
	// 正则表达式匹配"insert ignore into"，忽略大小写
	re := regexp.MustCompile(`(?i)insert ignore into`)

	// 在SQL语句中查找匹配项
	match := re.FindStringIndex(sql)

	// 如果找到匹配项，返回起始和结束位置
	if match != nil {
		return match[0], match[1]
	}

	// 如果没有找到匹配项，返回-1作为起始和结束位置
	return -1, -1
}
```

## replaceInsertIgnore

```
func replaceInsertIgnore(input string) string {
	// 编译正则表达式，匹配"INSERT IGNORE INTO"，忽略大小写
	re, err := regexp.Compile(`(?i)\bINSERT\s+IGNORE\s+INTO\b`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return input
	}

	result := re.ReplaceAllStringFunc(input, func(match string) string {
		// 如果这是第一次替换，则返回新的子串
		// 否则，返回原始匹配项
		staticCount := 0
		if staticCount == 0 {
			staticCount++
			return "insert into"
		}
		return match
	})

	return result
}
```
