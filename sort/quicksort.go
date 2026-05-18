package main

func quickSortTraceRec(arr []int, low, high, depth int) {
	if low >= high {
		return
	}

	pivotIndex := partition(arr, low, high)
	quickSortTraceRec(arr, low, pivotIndex-1, depth+1)
	quickSortTraceRec(arr, pivotIndex+1, high, depth+1)
}

func partition(arr []int, low, high int) int {
	pivot := arr[low]
	i := low + 1
	for j := low + 1; j <= high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[low], arr[i-1] = arr[i-1], arr[low]
	return i - 1
}
