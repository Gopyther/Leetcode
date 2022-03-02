package main

import "fmt"

func main() {
	// Declare a new array of 6 elements
	array := make([]int, 6)

	// Variable to keep track of the length of the array
	var length int = 0

	// Iterate through the 6 indexes of the Array.
	for i := 0; i < 6; i++ {
		// Add a new element and increment the length as well
		array[length] = i
		length++
	}

	// Print out the results of searching for 4 and 30.
	fmt.Println("Does the array contain the element 4? - ", linearSearch(array, length, 4))
	fmt.Println("Does the array contain the element 30? - ", linearSearch(array, length, 30))

	// Does the array contain the element 4? - true
	// Does the array contain the element 30? - false
}

func linearSearch(array []int, length, element int) bool {
	// Check for edge cases. Is the array null or empty?
	// If it is, then we return false because the element we're
	// searching for couldn't possibly be in it.
	if array == nil || length == 0 {
		return false
	}

	// Carry out the linear search by checking each element,
	// starting from the first one.
	for i := 0; i < length; i++ {
		// We found the element at index i.
		// So we return true to say it exists.
		if array[i] == element {
			return true
		}
	}

	// We didn't fine the element in the array
	return false
}
