package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getRemote(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		// 如果有错误返回错误内容
		return nil, err
	}
	// 使用完成后要关闭，不然会占用内存
	defer res.Body.Close()
	// 读取字节流
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return bytes, err
}

// 获得资源,从本地
// 这里以图片为例子，excel，world也是可以的
func getLocal(url string) ([]byte, error) {
	fp, err := os.OpenFile(url, os.O_CREATE|os.O_APPEND, 6) // 读写方式打开
	if err != nil {
		// 如果有错误返回错误内容
		return nil, err
	}
	defer fp.Close()
	bytes, err := ioutil.ReadAll(fp)
	if err != nil {
		return nil, err
	}
	return bytes, err
}

func main() {
	// 这里以图片为例子，excel，world也是可以的
	//url := "https://file.hytwfy.top/jxb.jpg"
	url := "https://yy-ai-train.bj.bcebos.com/dataset/zhangbiwu/20211227/encode_jpeg/sw/2314776959_0.jpeg"
	resByte, err := getRemote(url)
	if err != nil {
		log.Println(err)
	}
	// resByte 就是要上传到服务器的字节流
	log.Println("read image ok: size=", len(resByte))
}