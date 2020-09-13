package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func ArrayMatching(strArr []string) string {
	firstStr := strArr[0]
	secondStr := strArr[1]
	pattern, _ := regexp.Compile("[0-9]?[0-9]")

	// Find all the numbers
	firstArr := pattern.FindAllString(firstStr, -1)
	fmt.Println(firstArr)
	secondArr := pattern.FindAllString(secondStr, -1)
	fmt.Println(secondArr)

	// Initialise all the int arrays
	var firstIntArr []int
	var secondIntArr []int

	// Append all the ints
	for i := 0; i < len(firstArr); i++ {
		num, _ := strconv.Atoi(firstArr[i])
		firstIntArr = append(firstIntArr, num)
	}
	for i := 0; i < len(secondArr); i++ {
		num, _ := strconv.Atoi(secondArr[i])
		secondIntArr = append(secondIntArr, num)
	}
	fmt.Println(firstIntArr)
	fmt.Println(secondIntArr)

	// Initialise result array
	var resultArr []int
	var tempArrShort []int // use this to select shorter array
	var tempArrLong []int
	if len(firstIntArr) < len(secondIntArr) {
		tempArrShort = firstIntArr
		tempArrLong = secondIntArr
	} else {
		tempArrShort = secondIntArr
		tempArrLong = firstIntArr
	}

	for i := 0; i < len(tempArrShort); i++ {
		resultArr = append(resultArr, tempArrShort[i]+tempArrLong[i])
	}

	for i := len(tempArrShort); i < len(tempArrLong); i++ {
		resultArr = append(resultArr, tempArrLong[i])
	}
	resultString := ""
	for i := 0; i < len(resultArr); i++ {
		num := strconv.Itoa(resultArr[i])
		resultString += num + "-"
	}

	return resultString

}

func main() {
	// do not modify below here, readline is our function
	// that properly reads in the input for you
	var arr = []string{"[1,2,5,6]", "[5,2,8,11]"}
	fmt.Println(ArrayMatching(arr))
}
