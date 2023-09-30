package main

import (
    "log"
    "os"

    // couldn't find the go-fsnotify, this is what pops up on github
    "github.com/fsnotify/fsnotify"
)

func main() {
    monitorFile("./inlogs/test.log")
}

func monitorFile(filepath string) {

    // starting watcher
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()

    // monitor events
    go func() {
        for {
            select {
            case event := <-watcher.Events:
                switch event.Op {
                case fsnotify.Create:
                    log.Println("Created")

                case fsnotify.Write:
                    log.Println("Write")

                case fsnotify.Chmod:
                    log.Println("Chmod")

                case fsnotify.Remove, fsnotify.Rename:
                    log.Println("Moved or Deleted")

                    respawnFile(event.Name)

                    // add the file back to watcher, since it is removed from it
                    // when file is moved or deleted
                    log.Printf("add to watcher file:  %s\n", filepath)
                    // add appears to be concurrently safe so calling from multiple go routines is ok
                    err = watcher.Add(filepath)
                    if err != nil {
                        log.Fatal(err)
                    }

                    // there is  not need to break the loop
                    // we just continue waiting for events from the same watcher

                }
            case err := <-watcher.Errors:
                log.Println("Error:", err)
            }
        }
    }()

    // add file to the watcher first time
    log.Printf("add to watcher 1st: %s\n", filepath)
    err = watcher.Add(filepath)
    if err != nil {
        log.Fatal(err)
    }

    // to keep waiting forever, to prevent main exit
    // this is to replace the done channel
    select {}
}

func respawnFile(filepath string) {
    log.Printf("re creating file %s\n", filepath)

    // you just need the os.Create()
    respawned, err := os.Create(filepath)
    if err != nil {
        log.Fatalf("Err re-spawning file: %v", filepath)
    }
    defer respawned.Close()

    // there is no need to call monitorFile again, it never returns
    // the call to "go monitorFile(filepath)" was causing another go routine leak
}
