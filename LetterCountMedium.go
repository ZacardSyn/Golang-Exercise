package main

import (
	"fmt"
	"strings"
)

func LetterCount(str string) string {
	var largestArr []int
	strArr := strings.Split(str, " ")
	fmt.Println(strArr)
	fmt.Println(len(strArr))
	for i := 0; i < len(strArr); i++ {
		letterArr := strings.Split(strArr[i], "")
		var countArr []int
		counter := 1
		largest := 0
		fmt.Println(letterArr)
		for j := 0; j < len(letterArr); j++ {
			for k := j + 1; k < len(letterArr); k++ {
				if letterArr[j] == letterArr[k] {
					counter++
				}
			}
			fmt.Println(counter)
			countArr = append(countArr, counter)
			counter = 1

			for k := 0; k < len(countArr); k++ {
				if countArr[k] > largest {
					largest = countArr[k]
				}
			}
		}
		largestArr = append(largestArr, largest)

	}
	fmt.Println(largestArr)
	return "Done"
}

func main() {
	// do not modify below here, readline is our function
	// that properly reads in the input for you
	//array := []int{4, 6, 23, 10, 1, 3}
	fmt.Println(LetterCount("Today, is the greatest day ever!"))
}
