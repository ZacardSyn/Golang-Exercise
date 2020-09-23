package main

import (
	"fmt"
	"strings"
)

func LetterCapitalise(str string) string {
	stringSplit := strings.Split(str, " ")
	//upperStringSplit := ""
	for i := 0; i < len(stringSplit); i++ {
		stringSplit[i] = strings.Replace(stringSplit[i], string(stringSplit[i][0]), strings.ToUpper(string(stringSplit[i][0])), 1)
	}
	str = strings.Join(stringSplit, " ")
	return str
}

//[1 4], len = 2

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(LetterCapitalise("hello world"))

}
