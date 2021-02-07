package main

import (
	"fmt"
	"strconv"
	"time"
	"encoding/json"
)

func main() {
	extends := map[string]interface{}{
		"bdResCode": "110001",
		"bdResMessage": "downloadFail",
		"uid": 123,
	}

	b, err := json.Marshal(extends)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}
	
	fmt.Println("b:", string(b))
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