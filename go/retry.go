package main

import (
    //"fmt"
    "log"
    //"net/http"
    //"io/ioutil"
    //"encoding/base64"
    //"encoding/json"
    //"go.uber.org/zap"
)

func calcRetryInterval(retry uint32) uint32 {
  var interval uint32 = 60

  const (
    cnt uint32 = 10
    incr uint32 = 10
    maxTryInterval uint32 = 60
  )

  if (retry >= 0 && retry < cnt) {
    interval = 1
  } else if (retry >= cnt && retry < (cnt + incr)) {
    interval = 2
  } else if (retry >= (cnt + incr) && retry < (cnt + 2*incr)) {
    interval = 4
  } else {
    interval = maxTryInterval
  }
  //log.Printf("interval=%+v\n", interval)

  return interval
}


func main() {
  for i := 0; i < 80; i++ {
    interval := calcRetryInterval(uint32(i))
    log.Printf("interval=%+v\n", interval)
  }
}
