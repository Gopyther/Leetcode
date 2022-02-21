package main

import "fmt"

func main() {
	// Declare an integer array of 6 elements
	var intArray [6]int
	var length int

	// Add 3 elements to the Array
	for i := 0; i < 3; i++ {
		intArray[length] = i
		length++
	}
	// Insert a new element at the end of the Array. Again,
	// it's important to ensure that there is enough space
	// in the array for inserting a new element.
	intArray[length] = 10
	length++

	// First, we will have to create space for a new element.
	// We do that by shifting each element one index to the right.
	// This will firstly move the element at index 3, then 2, then 1, then finally 0.
	// We need to go backwards to avoid overwriting any elements.
	for i := 3; i >= 0; i-- {
		intArray[i+1] = intArray[i]
	}
	// Now that we have created space for the new element,
	// we can insert it at the beginning.
	intArray[0] = 20
	// Say we want to insert the element at index 2.
	// First, we will have to create space for the new element.
	for i := 4; i >= 2; i-- {
		// Shift each element one position to the right.
		intArray[i+1] = intArray[i]
	}
	// Now that we have created space for the new element,
	// we can insert it at the required index.
	intArray[2] = 30

	for i := 0; i < len(intArray); i++ {
		fmt.Println("Index", i, "Contains", intArray[i])
	}

}
