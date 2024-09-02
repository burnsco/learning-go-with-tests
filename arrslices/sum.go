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
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func main() {
	var slice []int

	slice = append(slice, 1, 2, 3)
	fmt.Println("Slice after first append:", slice, "Capacity:", cap(slice))

	slice = append(slice, 5, 6, 7)
	fmt.Println("Slice after secondappend:", slice, "Capacity:", cap(slice))

	arr1 := []int{1, 3, 3}
	arr2 := []int{4, 5, 6}

	fmt.Println(SumAll(arr1, arr2))
}
