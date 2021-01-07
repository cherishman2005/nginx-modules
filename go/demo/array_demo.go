package main

import(
	"fmt"
)


var (
	WhitelistUids []uint64
)

func search(slice *[]uint64, value uint64) int {
	for i, v := range *slice {
		if v == value {
			return i
		}
	}
	return -1
}

func delete(slice *[]uint64, index int) {
	if index > len(*slice) - 1 {
	   return  
	}
	*slice = append((*slice)[:index], (*slice)[index+1:]...)
}

func clear(slice *[]uint64) {
	*slice = []uint64{}
}


func main() {
	fmt.Println(search(&WhitelistUids, 123))
	WhitelistUids = append(WhitelistUids, 123)
	fmt.Println(search(&WhitelistUids, 123))
	fmt.Println(search(&WhitelistUids, 234))
	clear(&WhitelistUids)
	fmt.Println(WhitelistUids)
}
