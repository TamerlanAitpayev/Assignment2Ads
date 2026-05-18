package main

func linearSearchTrace(numbers []int, target int) {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == target {
			return
		}
	}
}
