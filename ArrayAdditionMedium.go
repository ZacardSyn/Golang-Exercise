package main

import (
	"fmt"
)

func ArrayAdditionMedium(arr []int) string {
	largest := arr[0]
	largestIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
			largestIndex = i
		}
	}
	copy(arr[largestIndex:], arr[largestIndex+1:])
	arr = arr[:len(arr)-1]
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
	for i := 0; i < len(arr)-1; i++ {
		sum := arr[i]
		for j := i + 1; j < len(arr); j++ {
			sum += arr[j]
			fmt.Println(sum)
			if sum == largest {
				return "true"
			}
		}
	}
	return "Done"
}

func main() {
	// do not modify below here, readline is our function
	// that properly reads in the input for you
	array := []int{4, 6, 23, 10, 1, 3}
	fmt.Println(ArrayAddition(array))
}
