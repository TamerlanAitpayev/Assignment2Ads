package main

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


func main() {
	arr := []int{17, 3, 24, 5, 10, 1, 31}
	target := 10

	bubbleSortTrace(arr)
	quickSortTrace(arr)
	mergeSortTrace(arr)
	hashingTrace(arr)
	divideAndConquerTrace(arr)
	linearSearchTrace(arr, target)
}
