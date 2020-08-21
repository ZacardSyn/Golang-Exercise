package main

import (
	"fmt"
	"strconv"
	"strings"
)

func LargestPair(num int) string {
	numStr := strconv.Itoa(num)
	numStrArr := strings.Split(numStr, "")
	var resultStrArr []string

	// Loop through entire array
	for i := 0; i < len(numStrArr)-1; i++ {
		for j := i + 1; j < len(numStrArr); j++ {
			var strTemp = []string{string(numStr[i]), string(numStr[j])}
			num := strings.Join(strTemp, "")
			resultStrArr = append(resultStrArr, num)
		}
	}
	returnVal := resultStrArr[0]
	for i := 0; i < len(resultStrArr); i++ {
		if returnVal < resultStrArr[i] {
			returnVal = resultStrArr[i]
		}
	}
	return returnVal

}

func main() {
	// do not modify below here, readline is our function
	// that properly reads in the input for you

	fmt.Println(LargestPair(4759472))
}
