package main

import (
	"fmt"
	"strings"
)

func divideAndConquerTrace(numbers []int) {
	fmt.Println("B2. Divide and Conquer")
	minValue, maxValue, comparisons := findMinMax(numbers, 0)

	fmt.Println("Min:", minValue)
	fmt.Println("Max:", maxValue)
	fmt.Println("Range:", maxValue-minValue)
	fmt.Println("D&C comparisons:", comparisons)
	fmt.Println("Naive comparisons:", 2*(len(numbers)-1))
	fmt.Println("D&C is better because it compares smaller parts first.")
}

func findMinMax(arr []int, depth int) (int, int, int) {
	if len(arr) == 1 {
		fmt.Println(strings.Repeat("  ", depth), "base:", arr)
		return arr[0], arr[0], 0
	}

	if len(arr) == 2 {
		fmt.Println(strings.Repeat("  ", depth), "base:", arr)
		if arr[0] < arr[1] {
			return arr[0], arr[1], 1
		}
		return arr[1], arr[0], 1
	}

	mid := len(arr) / 2
	leftMin, leftMax, leftCount := findMinMax(arr[:mid], depth+1)
	rightMin, rightMax, rightCount := findMinMax(arr[mid:], depth+1)

	minValue := leftMin
	if rightMin < minValue {
		minValue = rightMin
	}

	maxValue := leftMax
	if rightMax > maxValue {
		maxValue = rightMax
	}

	total := leftCount + rightCount + 2

	fmt.Println(strings.Repeat("  ", depth), "combine:", arr, "min =", minValue, "max =", maxValue)
	return minValue, maxValue, total
}
