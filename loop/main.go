package main

import "fmt"

func main() {
	squareNumbers := [10]int{}
	// Go through each of the Array indexes, from 0 to 9.
	for i := 0; i < 10; i++ {
		// We need to be careful with the 0-indexing. The next square number
		// is given by (i + 1) * (i + 1).
		// Calculate it and insert it into the Array at index i.
		square := (i + 1) * (i + 1)
		squareNumbers[i] = square
	}

	for i := 0; i < 10; i++ {
		// Access and print what's at the i'th index.
		fmt.Println(squareNumbers[i])
	}
	// Will print:
	// 1
	// 4
	// 9
	// 16
	// 25
	// 36
	// 49
	// 64
	// 81
	// 100

	for _, square := range squareNumbers {
		// Access and print what's at the i'th index.
		fmt.Println(square)
	}
}
