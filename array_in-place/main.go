package main

import "fmt"

func main() {
	array := []int{9, -2, -9, 11, 56, -12, -3}

	fmt.Println(squareEven(array, len(array)))
}

func squareEven(array []int, length int) []int {

	// Check for edge cases.
	if array == nil {
		return nil
	}

	// Iterate through the original Array.
	for i := 0; i < length; i++ {

		// If the index is an even number, then we need to square the element
		// and replace the original value in the Array.
		if i%2 == 0 {
			array[i] *= array[i]
		}
		// Notice how we don't need to do *anything* for the odd indexes? :-)
	}
	// We just return the original array. Some problems on leetcode require you
	// to return it, and other's dont.
	return array
}
