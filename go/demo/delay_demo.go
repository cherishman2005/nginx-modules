package main

import(
	"fmt"
	"time"
	"strconv"
)

func metricDelayState(t string) {
	distFloat,err := strconv.ParseFloat(t, 64)
	if (err != nil) {
		fmt.Println("ParseFloat fail")
	}
	
	begin := int64(distFloat*1000)
	
	end := int64(time.Now().UnixNano()/1000000)

	usetime := end - begin

	fmt.Println(begin, end, usetime)
}

func main() {
	metricDelayState("1611299064.000")
}