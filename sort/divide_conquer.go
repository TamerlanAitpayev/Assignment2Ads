package main

func divideAndConquerTrace(numbers []int) {
	_, _, _ = findMinMax(numbers, 0)
}

func findMinMax(arr []int, depth int) (int, int, int) {
	if len(arr) == 1 {
		return arr[0], arr[0], 0
	}

	if len(arr) == 2 {
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
	return minValue, maxValue, total
}
