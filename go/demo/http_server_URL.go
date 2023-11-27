package main
  
import (
        "fmt"
        "net/http"
        //"strings"
)

func helloHandlers(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("URL.Path: %s\n", r.URL.Path)
        //remPartOfURL := r.URL.Path[len("/hello/"):] // get everything after the /hello/ part of the URL
        fmt.Fprintf(w, "Hello %s", r.URL.Path)
}

func main() {
        http.HandleFunc("/", helloHandlers)
        http.ListenAndServe("localhost:9999", nil)
}
