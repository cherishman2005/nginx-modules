package main

import (
	"log"
	"net/http"
	"net/url"
	"time"
	"encoding/json"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    Addr string `json:"addr,omitempty"`
}

func userHandler(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		urlStr := r.URL.RawQuery
		if len(urlStr) > 0 {
			m, _ := url.ParseQuery(urlStr)
			log.Println(m)
			if len(m["name"]) > 0 {
				log.Println(m["name"][0])
		
			}
		}

		p1 := Person {
			Name: "awu",
			Age:  18,
		}

		data, err := json.Marshal(p1)
		if err != nil {
			panic(err)
		}

		w.Write([]byte(data))
	}
	return http.HandlerFunc(fn)
}

func main() {
	mux := http.NewServeMux()

	th := userHandler(time.RFC1123)
	mux.Handle("/user", th)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}