# goland json.Marshal导致&变成\u0026

背景：goland后台使用json.Marshal转换时，会将<,>,&转化为unicode编码，导致入库时&变成\u0026。

原因： json.marshal默认escapeHtml为true，会将<、>、&等字符转义。

## 解决方案1

```
import (
	"bytes"
	"encoding/json"
	"fmt"
)
 
type MarshalTest struct {
	Url string `json:"url"`
}
 
//序列化
func marshal_inner(data interface{}) ([]byte, error) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	if err := jsonEncoder.Encode(data); err != nil {
		return nil, err
	}
 
	return bf.Bytes(), nil
}
 
func main() {
	t := &MarshalTest{
		Url: "http://www.baidu.com?seq=213&uuid=1",
	}
	val, err := marshal_inner(t)
	if err != nil {
		fmt.Println("marshal_inner failed.err:", err)
		return
	}
	fmt.Println("marshal_inner val:", string(val))
}
```
上面的解决方案能使转换正确，但是会在string(bf.Bytes())或者bf.String()时，默认在字符串结尾加上\n，也不能达到想要的效果。

## 方案2

```
content = strings.ReplaceAll(content, "\\u003c", "<") // 必须是两个斜杆，不能改成一个斜杆
content = strings.ReplaceAll(content, "\\u003e", ">")
content = strings.ReplaceAll(content, "\\u0026", "&")
```
直接使用替换，来替换不符合的字符，简单有效。
