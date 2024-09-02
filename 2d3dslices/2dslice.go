package main

import "fmt"

func main() {
	// Initialize a 3d slice
	threeDSlice := [][][]int{
		{
			{1, 2},
			{3, 4},
			{5, 6},
		},
		{
			{7, 8},
			{9, 0},
		},
	}

	fmt.Println("2d Slice:", threeDSlice[1][1])

}
