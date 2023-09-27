package main

import (
  "log"
  "os"

  "github.com/fsnotify/fsnotify"
)

const (
    // 文件写入mode
    fileOpenMode = 0666

    // 文件Flag
    fileFlag = os.O_WRONLY | os.O_CREATE | os.O_APPEND
)

func main() {

    fw, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal("NewWatcher failed: ", err)
    }
    defer fw.Close()
  
    done := make(chan bool)
    go func() {
        defer close(done)

        for {
            select {
            case event, ok := <-fw.Events:
                if !ok {
                    return
                }
                
                log.Printf("%s %s %d\n", event.Name, event.Op, event.Op)
                if event.Op != fsnotify.Create && event.Op != fsnotify.Remove && event.Op != fsnotify.Rename {
                    // file still exists
                    log.Println("file still exists")
                }
            
            case err, ok := <-fw.Errors:
                if !ok {
                  return
                }
                log.Println("error:", err)
            }
        }
    }()


    logFilename := "1.txt"
    file, err := os.OpenFile(logFilename, fileFlag, fileOpenMode)
    if err != nil {
        // 重试
        file, err = os.OpenFile(logFilename, fileFlag, fileOpenMode)
        if err != nil {
            return
        }
    }


  err = fw.Add(file.Name())
  if err != nil {
    log.Fatal("Add failed:", err)
  }
  <-done
}
