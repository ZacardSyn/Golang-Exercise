package main

import (
	"fmt"
	"strings"
)

func LetterCountl(str string) string {

	strSplit := strings.Split(str, " ")
	scoreArray := make([]int, len(strSplit)) // An array that stores values of iterations of most repeated letters
	for index, char := range strSplit {
		countTemp := 0
		count := 0
		//The most iterated letter for that one word
		for i := 0; i < len(char)-1; i++ { // to loop through all letter in a word
			for j := 0; j < len(char); j++ { // look through all letter to find a match
				//fmt.Println(char[i], char[j])
				if char[i] == char[j] {
					countTemp++
					//fmt.Println(countTemp)
				}
				if countTemp > count {
					count = countTemp
					//fmt.Println(count)
				}
			}
			countTemp = 0
		}
		scoreArray[index] = count
		fmt.Println(scoreArray)
	}
	maxIdx := 0
	for i := 0; i < len(scoreArray)-1; i++ {

		if scoreArray[i+1] > scoreArray[maxIdx] {
			maxIdx = i + 1
			//fmt.Println(scoreArray[maxIdx])
		}
	}
	if scoreArray[maxIdx] == 1 {
		return "-1"
	}
	return strSplit[maxIdx]
}

//[1 4], len = 2

func main() {
	// do not modify below here, readline is our function
	// that properly reads in the input for you

	fmt.Println(LetterCountl("Hello Apple Pie"))

}
