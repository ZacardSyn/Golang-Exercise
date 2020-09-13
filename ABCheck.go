package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ABCheck(str string) string {
	var idxArray [][]int
	var idx []int
	str = strings.ReplaceAll(str, " ", "")
	pattern, _ := regexp.Compile("[aA]")
	idxArray = pattern.FindAllStringIndex(str, -1)
	for i := 0; i < len(idxArray); i++ {
		idx = append(idx, idxArray[i][0])
	}

	for i := 0; i < len(idx); i++ {
		if string(str[idx[i]+3]) == "b" {
			return "true"
		}
	}
	return "false"
}

//[1 4], len = 2

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(ABCheck("iahvenoideawhatimdoing"))

}
