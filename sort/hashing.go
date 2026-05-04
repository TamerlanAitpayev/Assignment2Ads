package main

import "fmt"

func hashingTrace(numbers []int) {
	size := 7
	table := []int{-1, -1, -1, -1, -1, -1, -1}

	fmt.Println("B1. Hashing")
	for _, key := range numbers {
		start := key % size
		index := start

		fmt.Println("Insert", key, "h =", start)

		for table[index] != -1 {
			fmt.Println("Collision at index", index)
			index = (index + 1) % size
		}

		table[index] = key
		fmt.Println("Put at index", index)
	}

	fmt.Println("Final table:")
	for i := 0; i < size; i++ {
		fmt.Println(i, ":", table[i])
	}
}
