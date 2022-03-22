package main

import (
	"fmt"
	"sort"
)

// Java uses the main method, but Golang uses func main.
func main() {
	// 1. Initialize
	var a0 [5]int
	var a1 = []int{1, 2, 3}
	// 2. Get Length
	fmt.Println("The size of a1 is:", len(a1))
	// 3. Access Element
	fmt.Println("The first element is: ", a1[0])
	// 4. Iterate all Elements
	fmt.Print("[Version 1] The contents of a1 are:")
	for i := 0; i < len(a1); i++ {
		fmt.Print(" ", a1[i])
	}
	fmt.Println(" ")
	fmt.Print("[Version 2] The contents of a1 are:")
	for _, v := range a1 {
		fmt.Print(" ", v)
	}
	fmt.Println(" ")
	// 5. Modify Element
	a1[0] = 4
	sort.Ints(a1)
	for _, v := range a1 {
		fmt.Print(" ", v)
	}
	fmt.Println(" ")
	// Go does not allow to use not used variable.
	for _, v := range a0 {
		fmt.Print(" ", v)
	}
}
