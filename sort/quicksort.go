package main

import (
	"fmt"
	"strings"
)

func quickSortTrace(numbers []int) {
	arr := make([]int, len(numbers))
	copy(arr, numbers)

	fmt.Println("A2. Quick Sort")
	fmt.Println("Start:", arr)

	quickSortTraceRec(arr, 0, len(arr)-1, 0)

	fmt.Println("Result:", arr)
}

func quickSortTraceRec(arr []int, low, high, depth int) {
	if low >= high {
		return
	}

	indent := strings.Repeat("  ", depth)
	pivotIndex := partition(arr, low, high, indent)
	quickSortTraceRec(arr, low, pivotIndex-1, depth+1)
	quickSortTraceRec(arr, pivotIndex+1, high, depth+1)
}

func partition(arr []int, low, high int, indent string) int {
	pivot := arr[low]
	fmt.Printf("%sPartition %v, pivot=%d\n", indent, arr[low:high+1], pivot)

	i := low + 1
	for j := low + 1; j <= high; j++ {
		if arr[j] < pivot {
			if i != j {
				fmt.Printf("%s swap %d and %d -> %v\n", indent, arr[j], arr[i], arr[low:high+1])
			}
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	fmt.Printf("%s swap pivot %d with %d -> ", indent, arr[low], arr[i-1])
	arr[low], arr[i-1] = arr[i-1], arr[low]
	fmt.Printf("%v\n", arr[low:high+1])

	return i - 1
}
