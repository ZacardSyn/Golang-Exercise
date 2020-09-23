package main

import (
	"fmt"
)

func OffLineMinimum(arr []string) []string {
	var min []string
	for i, v := range arr {
		if v == "E" {
			arrTemp := arr[0:i]
			minTemp := arrTemp[0]
			idxTemp := 0
			arr[0] = "X"
			for j, k := range arrTemp {
				if k == "X" || k == "E" {
					continue
				}
				if minTemp > k {
					arrTemp[idxTemp] = minTemp
					minTemp = k
					idxTemp = j
					arr[j] = "X"
				}
			}
			fmt.Println(minTemp)
			min = append(min, minTemp)
		}
	}
	return min
}

//[1 4], len = 2

func main() {
	// do not modify below here, readline is our function
	// that properly reads in the input for you
	var arr = []string{"5", "4", "6", "E", "1", "7", "E", "E", "3", "2"}
	fmt.Println(OffLineMinimum(arr))

}
