package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	extends := map[string]string{
		"bdResCode": "110001",
		"bdResMessage": "downloadFail",
	}
	code := getLogoAuditResCode(extends)
	fmt.Println(code)
	
	fmt.Println(GetContext(123))
}


func GetContext(id uint64) string {
	return fmt.Sprintf("active.%v.%v", id, time.Now().Unix())
}

func getLogoAuditResCode(extends map[string]string) int {
	if extends == nil {
		return 0
	}
	if _, ok := extends["bdResCode"]; ok {
		code, _ := strconv.ParseInt(extends["bdResCode"], 10, 64);
		return int(code)
	}

	return 0
}