package main

import (
   "log"
   "net"
   "net/http"
)

type Server struct {
   webHandler http.HandlerFunc
}

func (srv *Server) start(portStr string) error {
   lc, err := net.Listen("tcp", portStr)
   if err != nil {
      log.Printf("Listen err: %v", err)
      return err
   }
   defer lc.Close()

   http.HandleFunc("/", srv.webHandler)

   server := &http.Server{}
   err = server.Serve(lc)
   if err != nil {
      log.Printf("Serve err: %v", err)
      return err
   }

   return nil
}

func main() {
   srv := &Server{
      webHandler: func(w http.ResponseWriter, r *http.Request) {
         // 处理请求的逻辑
         w.Write([]byte("Hello, world!"))
      },
   }

   err := srv.start(":8080")
   if err != nil {
      log.Fatalf("Failed to start server: %v", err)
   }
}
