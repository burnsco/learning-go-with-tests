package main

import "fmt"

func Sum(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func main() {
	exampleArray := []int{1, 3, 4, 5, 1, 2, 3}
	fmt.Println("Array elements: ", exampleArray)
	fmt.Println("Sum of elements: ", Sum(exampleArray))
}
