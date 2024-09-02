package main

import "fmt"

func Sum(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

func main() {
	arr1 := []int{1, 3, 3}
	arr2 := []int{4, 5, 6}

	fmt.Println(SumAll(arr1, arr2))
}
