package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sameNaive([]int{1, 2, 3}, []int{5, 1, 9}))      // should return false
	fmt.Println(sameNaive([]int{1, 2, 3}, []int{4, 1, 9}))      // should return true
	fmt.Println(sameRefactored([]int{1, 2, 3}, []int{5, 1, 9})) // should return false
	fmt.Println(sameRefactored([]int{1, 2, 3}, []int{4, 1, 9})) // should return true
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

func sameRefactored(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	var (
		frequencyCounter1 = make(map[int]int)
		frequencyCounter2 = make(map[int]int)
	)

	for _, v := range arr1 {
		_, ok := frequencyCounter1[v]
		if ok {
			frequencyCounter1[v] = frequencyCounter1[v] + 1
		} else {
			frequencyCounter1[v] = 1
		}
	}

	for _, v := range arr2 {
		_, ok := frequencyCounter2[v]
		if ok {
			frequencyCounter2[v] = frequencyCounter2[v] + 1
		} else {
			frequencyCounter2[v] = 1
		}
	}

	for k, v := range frequencyCounter1 {
		counter, ok := frequencyCounter2[int(math.Pow(float64(k), 2))]
		if !ok || counter != v {
			return false
		}
	}

	return true
}
