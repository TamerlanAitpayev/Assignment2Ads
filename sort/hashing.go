package main

func hashingTrace(numbers []int) {
	size := 7
	table := []int{-1, -1, -1, -1, -1, -1, -1}

	for _, key := range numbers {
		start := key % size
		index := start

		for table[index] != -1 {
			index = (index + 1) % size
		}

		table[index] = key
	}

	_ = table
}
