package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	WhitelistUidMap map[uint64]bool
)

func init() {
	WhitelistUidMap = make(map[uint64]bool)
}

func main() {
	s := "123,456,789"
	
	TransferStringToMap(s, &WhitelistUidMap)
	
	fmt.Println(WhitelistUidMap)
}

func TransferStringToMap(s string, slice *map[uint64]bool) {
	elements := strings.Split(s, ",")
	
	for _, v := range elements {
		uid, err := strconv.ParseUint(v, 10, 64);
		if err != nil {
			continue
		}
		(*slice)[uid] = true
	}
}