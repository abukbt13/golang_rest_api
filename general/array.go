package main

import (
	"fmt"
)

func main() {
	var arr [5]float32
	arr[0] = 85.89
	arr[1] = 45
	arr[2] = 87
	arr[3] = 95
	arr[4] = 29

	// fmt.Println(arr[3])
	var total float32 = 0
	for i := 0; i < len(arr); i++ {
		total += arr[i]

	}
	fmt.Println(float32(total) / float32(len(arr)))
}
