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
                   
