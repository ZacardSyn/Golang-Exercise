package main

import (
	"fmt"
	"regexp"
)

func VowelCount(str string) int {
	pattern, _ := regexp.Compile("[aAeEiIoOuU]")
	var count [][]int
	count = pattern.FindAllStringIndex(str, -1)
	return len(count)
}

//[1 4], len = 2

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(VowelCount("iahvenoideawhatimdoing"))

}
