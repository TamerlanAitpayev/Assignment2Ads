package main

import (
	"fmt"
	"strings"
)

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		newArr := make([]int, len(arr))
		copy(newArr, arr)
		return newArr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func mergeSortTrace(dataset []int) {
	fmt.Println("A3. Merge Sort")
	fmt.Println("Divide:")
	printMergeTree(dataset, 0)

	left := mergeSort(dataset[:len(dataset)/2])
	right := mergeSort(dataset[len(dataset)/2:])
	fmt.Println("Left sorted:", left)
	fmt.Println("Right sorted:", right)
	fmt.Println("Merge:")

	i := 0
	j := 0
	result := []int{}

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			fmt.Println("take from left:", left[i], "->", result)
			i++
		} else {
			result = append(result, right[j])
			fmt.Println("take from right:", right[j], "->", result)
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		fmt.Println("take rest from left:", left[i], "->", result)
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		fmt.Println("take rest from right:", right[j], "->", result)
		j++
	}

	fmt.Println("Result:", result)
}

func printMergeTree(arr []int, depth int) {
	fmt.Println(strings.Repeat("  ", depth), arr)
	if len(arr) < 2 {
		return
	}

	mid := len(arr) / 2
	printMergeTree(arr[:mid], depth+1)
	printMergeTree(arr[mid:], depth+1)
}
