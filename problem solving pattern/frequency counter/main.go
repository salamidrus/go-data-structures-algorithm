package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sameNaive([]int{1, 2, 3}, []int{4, 1, 9}))
}

func sameNaive(arr1 []int, arr2 []int) bool {
	// compare length, should be same
	if len(arr1) != len(arr2) {
		return false
	}

	// loop over the first array
	// check the squared value in second array
	// remove the value on the array
	for _, num1 := range arr1 {
		correctIndex := checkValueByIndex(arr2, int(math.Pow(float64(num1), 2)))
		if correctIndex == -1 {
			return false
		}
		remove(arr2, correctIndex)
	}

	return true
}

func remove(slice []int, idx int) []int {
	return append(slice[:idx], slice[idx+1:]...)
}

func checkValueByIndex(slice []int, val int) int {
	for idx, v := range slice {
		if val == v {
			return idx
		}
	}

	return -1
}
