package main
import (
    //"fmt"
    "log"
    "strings"
)

func main() {
    remark := "test offline"

    if strings.HasPrefix(remark,"offline") {
        log.Printf("\"%s\"  HasPrefix", remark)
    } else {
        log.Printf("\"%s\"  not HasPrefix", remark)
    }

    if strings.HasSuffix(remark,"offline") {
        log.Printf("\"%s\"  HasSuffix", remark)
    } else {
        log.Printf("\"%s\"  not HasSuffix", remark)
    }
}
