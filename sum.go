package main

func Sum(numbers [5]int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}
