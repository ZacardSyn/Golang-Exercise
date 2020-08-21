package main

import (
	"fmt"
)

func SecondGreatLow(arr []int) (int, int) {
	secondLow := 0
	secondGreat := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != arr[i+1] {
			secondLow = arr[i+1]
			break
		}
	}

	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] != arr[i-1] {
			secondGreat = arr[i-1]
			break
		}
	}
	return secondLow, secondGreat
}

//[1 4], len = 2

func main() {
	// do not modify below here, readline is our function
	// that properly reads in the input for you
	var arr = []int{2, 198, 45, 97, 65, 42}
	fmt.Println(SecondGreatLow(arr))

}
