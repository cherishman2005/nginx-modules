package main

import (
  "os"
  "fmt"
  "github.com/fsnotify/fsnotify"
  "path/filepath"
)

const (
    // 文件写入mode
    fileOpenMode = 0666

    // 文件Flag
    fileFlag = os.O_WRONLY | os.O_CREATE | os.O_APPEND
)

// This log writer sends output to a file
type TimeFileLogWriter struct {
	// The opened file
	filename     string
	baseFilename string // abs path
	file         *os.File
    fw           *fsnotify.Watcher

}

func main() {
    w := &TimeFileLogWriter{
    	filename: "bfe.log",
	}

	fw, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("new file status watcher err")
		return
	}
	w.fw = fw
    
    dir := filepath.Dir(w.filename)
	w.fw.Add(dir)
    fmt.Printf("dir:%s, filename:%s\n", dir, w.filename)

	for {
		select {
		case event, ok := <-w.fw.Events:
			if !ok {
				fmt.Println("fw.Events error")
				return
			}
            
            nameAbs, _ := filepath.Abs(event.Name)
            filenameAbs, _ := filepath.Abs(w.filename)
            fmt.Printf("event.Name Abs:%s, filename Abs:%s\n", nameAbs, filenameAbs)
            if nameAbs  != filenameAbs {
                fmt.Printf("not need to process, event: %s %s\n", event.Name, event.Op)
                continue
            }
			fmt.Printf("event: %s %s\n", event.Name, event.Op)
            if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
            	fmt.Println("file still exists")
				continue
            }

			//if event.Op != fsnotify.Create && event.Op != fsnotify.Remove && event.Op != fsnotify.Rename {
			//	// file still exists
			//	//fmt.Println("file still exists")
			//	continue
			//}
            
			fmt.Printf("event: %s %s\n", event.Name, event.Op)

			if w.file != nil {
				w.file.Close()
			}
			//if w.fw != nil {
			//	w.fw.Close()
			//}

			// 写入log文件
			file, err := w.openFileNoCache()
			if err != nil {
				//panic(err)
				continue
			}
			w.file = file

			fmt.Printf("create filename:%s\n", w.filename)

		case err, ok := <-w.fw.Errors:
			if !ok {
				fmt.Println("fw.Errors")
				return
			}
			fmt.Println("err:", err)
		}
	}
}

// 打开日志文件(不缓存句柄)
func (w *TimeFileLogWriter) openFileNoCache() (*os.File, error) {
	file, err := os.OpenFile(w.filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		// 重试
		file, err = os.OpenFile(w.filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return file, err
		}
	}

	return file, nil
}
