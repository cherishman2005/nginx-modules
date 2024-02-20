# golang http client如何上传和server如何接收文件


给了一个例子，利用mime/multipart来实现client如何上传一个文件到server，然后server如何接受这个文件。

## server.go代码
```
package main

import (
    "io"
    "os"
    "fmt"
    "io/ioutil"
    "net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    reader, err := r.MultipartReader()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    for {
        part, err := reader.NextPart()
        if err == io.EOF {
            break
        }

        fmt.Printf("FileName=[%s], FormName=[%s]\n", part.FileName(), part.FormName())
        if part.FileName() == "" {  // this is FormData
            data, _ := ioutil.ReadAll(part)
            fmt.Printf("FormData=[%s]\n", string(data))
        } else {    // This is FileData
            dst, _ := os.Create("./" + part.FileName() + ".upload")
            defer dst.Close()
            io.Copy(dst, part)
        }
    }
}

func main() {
    http.HandleFunc("/upload", uploadHandler)
    http.ListenAndServe(":8080", nil)
}
```

例子1：client 上传一个文件
```
package main

import (
    "io"
    "os"
    "log"
    "bytes"
    "io/ioutil"
    "net/http"
    "mime/multipart"
)

func main() {
    bodyBuffer := &bytes.Buffer{}
    bodyWriter := multipart.NewWriter(bodyBuffer)

    fileWriter, _ := bodyWriter.CreateFormFile("files", "file.txt")

    file, _ := os.Open("file.txt")
    defer file.Close()

    io.Copy(fileWriter, file)

    contentType := bodyWriter.FormDataContentType()
    bodyWriter.Close()

    resp, _ := http.Post("http://localhost:8080/upload", contentType, bodyBuffer)
    defer resp.Body.Close()

    resp_body, _ := ioutil.ReadAll(resp.Body)

    log.Println(resp.Status)
    log.Println(string(resp_body))
}
```

例子2：client上传多个文件
```
package main

import (
    "io"
    "os"
    "log"
    "bytes"
    "io/ioutil"
    "net/http"
    "mime/multipart"
)

func main() {
    bodyBuffer := &bytes.Buffer{}
    bodyWriter := multipart.NewWriter(bodyBuffer)

    // file1
    fileWriter1, _ := bodyWriter.CreateFormFile("files", "file1.txt")
    file1, _ := os.Open("file1.txt")
    defer file1.Close()
    io.Copy(fileWriter1, file1)

    // file2
    fileWriter2, _ := bodyWriter.CreateFormFile("files", "file2.txt")
    file2, _ := os.Open("file2.txt")
    defer file2.Close()
    io.Copy(fileWriter2, file2)

    contentType := bodyWriter.FormDataContentType()
    bodyWriter.Close()

    resp, _ := http.Post("http://localhost:8080/upload", contentType, bodyBuffer)
    defer resp.Body.Close()

    resp_body, _ := ioutil.ReadAll(resp.Body)

    log.Println(resp.Status)
    log.Println(string(resp_body))
}
```

例子3：上传其他Form数据
```
package main

import (
    "io"
    "os"
    "log"
    "bytes"
    "io/ioutil"
    "net/http"
    "mime/multipart"
)

func main() {
    bodyBuffer := &bytes.Buffer{}
    bodyWriter := multipart.NewWriter(bodyBuffer)

    // file1
    fileWriter1, _ := bodyWriter.CreateFormFile("files", "file1.txt")
    file1, _ := os.Open("file1.txt")
    defer file1.Close()
    io.Copy(fileWriter1, file1)

    // file2
    fileWriter2, _ := bodyWriter.CreateFormFile("files", "file2.txt")
    file2, _ := os.Open("file2.txt")
    defer file2.Close()
    io.Copy(fileWriter2, file2)

    // other form data
    extraParams := map[string]string{
        "title":       "My Document",
        "author":      "Matt Aimonetti",
        "description": "A document with all the Go programming language secrets",
    }
    for key, value := range extraParams {
        _ = bodyWriter.WriteField(key, value)
    }

    contentType := bodyWriter.FormDataContentType()
    bodyWriter.Close()

    resp, _ := http.Post("http://localhost:8080/upload", contentType, bodyBuffer)
    defer resp.Body.Close()

    resp_body, _ := ioutil.ReadAll(resp.Body)

    log.Println(resp.Status)
    log.Println(string(resp_body))
}
```

看server端的运行输出：
```
$ go build server.go && ./server
FileName=[file1.txt], FormName=[files]
FileName=[file2.txt], FormName=[files]
FileName=[], FormName=[description]
FormData=[A document with all the Go programming language secrets]
FileName=[], FormName=[title]
FormData=[My Document]
FileName=[], FormName=[author]
FormData=[Matt Aimonetti]

$ ls -l
total 25180
-rwxr-xr-x 1 ... 6179399 Jun 22 08:07 client
-rw-r--r-- 1 ...    1952 Jun 22 08:06 client.go
-rw-r--r-- 1 ...      15 Jun 22 07:11 file1.txt
-rw-r--r-- 1 ...      15 Jun 22 08:28 file1.txt.upload
-rw-r--r-- 1 ...      14 Jun 22 07:11 file2.txt
-rw-r--r-- 1 ...      14 Jun 22 08:28 file2.txt.upload
-rw-r--r-- 1 ...      15 Jun 22 07:56 file.txt
-rw-r--r-- 1 ...      15 Jun 22 08:10 file.txt.upload
-rwxr-xr-x 1 ... 6717437 Jun 22 07:59 server
```

-rw-r--r-- 1 ...    1580 Jun 22 07:58 server.go
