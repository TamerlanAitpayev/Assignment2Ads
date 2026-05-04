package main

import "fmt"

func linearSearchTrace(numbers []int, target int) {
	fmt.Println("B3. Linear Search")

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Compare", numbers[i], "with", target)

		if numbers[i] == target {
			fmt.Println("Found at index:", i)
			fmt.Println("Comparisons:", i+1)
			fmt.Println("Average-case time complexity: O(n)")
			return
		}
	}

	fmt.Println("Not found")
	fmt.Println("Comparisons:", len(numbers))
	fmt.Println("Average-case time complexity: O(n)")
}
