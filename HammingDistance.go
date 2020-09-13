package main

import (
	"fmt"
)

func HammingDistance(strArr []string) int {
	firstString := strArr[0]
	secondString := strArr[1]
	distance := 0
	for i := 0; i < len(firstString); i++ {
		if firstString[i] != secondString[i] {
			distance++
		}
	}
	return distance
}

//[1 4], len = 2

func main() {
	// do not modify below here, readline is our function
	// that properly reads in the input for you
	var arr = []string{"dac", "dap"}
	fmt.Println(HammingDistance(arr))

}
