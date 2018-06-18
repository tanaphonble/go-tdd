package main

// Sum - Sum all numbers in an array
func Sum(numbers [5]int) (sum int) {
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return
}
