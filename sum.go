package main

// Sum - Sum all numbers in an array
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

// SumAll - SumAll numbers in arrays to new array
func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumber := len(numbersToSum)
	sums = make([]int, lengthOfNumber)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return
}
