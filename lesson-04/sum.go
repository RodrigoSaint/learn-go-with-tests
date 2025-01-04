package main

func Sum(numbers []int) int {
	acc := 0
	for _, number := range numbers {
		acc += number
	}

	return acc
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
