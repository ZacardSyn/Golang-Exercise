package main

import (
	"fmt"
)

func OtherProducts(arr []int) []int {
	var resultArr []int
	for i := 0; i < len(arr); i++ {
		result := 1
		for j := 0; j < len(arr); j++ {
			if j == i {
				continue
			}
			result *= arr[j]
		}
		resultArr = append(resultArr, result)

	}
	return resultArr
}

func main() {
	// do not modify below here, readline is our function
	// that properly reads in the input for you
	var arr = []int{8, 7, 25, 4, 8, 0}
	fmt.Println(OtherProducts(arr))
}
