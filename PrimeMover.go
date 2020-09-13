package main

import (
	"fmt"
)

func PrimeMover(num int) int {
	counter := 1
	digit := 2
	for 1 == 1 {
		digit++
		if PrimeTime(int16(digit)) == true {
			counter++
			fmt.Println(counter, digit)
		}
		if counter == num {
			break
		}
	}
	return digit
}

func PrimeTime(num int16) bool {
	midPoint := num / 2
	for i := int16(2); i <= midPoint; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// do not modify below here, readline is our function
	// that properly reads in the input for you

	fmt.Println(PrimeMover(10))
}
