package main

import (
    "fmt"
    "github.com/go-redis/redis"
    "strconv"
    "time"
)

func main() {
    client := redis.NewClient(&redis.Options{
        Addr:     "127.0.0.1:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    for i := 0; i < 106; i++ {
        client.Set("ip_"+strconv.Itoa(i), "value"+strconv.Itoa(i), 120*time.Second)
    }

    var cursor uint64
    var n int
    for {
        var keys []string
        var err error
        //*扫描所有key，每次20条
        keys, cursor, err = client.Scan(cursor, "ip_*", 20).Result()
        if err != nil {
            panic(err)
        }
        n += len(keys)

        fmt.Printf("\nfound %d keys\n", n)
        var value string
        for _, key := range keys {
            value, err = client.Get(key).Result()
            fmt.Printf("%v %v\n", key, value)
        }
        if cursor == 0 {
            break
        }
    }
    fmt.Printf("n=%d\n", n)
}
