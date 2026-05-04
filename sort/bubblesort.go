package main

import "fmt"

func main() {
	arr := []int{17, 3, 24, 5, 10, 1, 31}
	target := 10

	fmt.Println("Dataset:", arr)
	fmt.Println("T1:", target)
	fmt.Println()

	bubbleSortTrace(arr)
	fmt.Println()
	quickSortTrace(arr)
	fmt.Println()
	mergeSortTrace(arr)
	fmt.Println()
	hashingTrace(arr)
	fmt.Println()
	divideAndConquerTrace(arr)
	fmt.Println()
	linearSearchTrace(arr, target)
}

func bubbleSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)

	n := len(result)
	for pass := 0; pass < n-1; pass++ {
		swapped := false
		for j := 0; j < n-pass-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return result
}

func bubbleSortTrace(numbers []int) {
	arr := make([]int, len(numbers))
	copy(arr, numbers)

	fmt.Println("A1. Bubble Sort")
	fmt.Println("Start:", arr)

	for pass := 1; pass < len(arr); pass++ {
		swapped := false

		for j := 0; j < len(arr)-pass; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		fmt.Println("Pass", pass, ":", arr)

		if !swapped {
			fmt.Println("Early Exit: array is already sorted after pass", pass)
			break
		}
	}

	fmt.Println("Result:", arr)
}
