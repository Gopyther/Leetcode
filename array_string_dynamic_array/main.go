package main

import (
	"fmt"
	"sort"
)

// Java uses the main method, but Golang uses func main.
func main() {
	// 1. initialize
	var v0 = []int{}
	var v1 []int // v1 == null
	// 2. cast an array to a vector
	var a = [...]int{0, 1, 2, 3, 4} // Java cannot declare list with int.
	v1 = a[:]                       // Same as b := a[0:len(a)]

	// 3. make a copy
	var v2 = v1 // another reference to v1
	var v3 = []int{}
	copy(v3, v1) // make an actual copy of v1

	// 4. get length
	fmt.Println("The size of v1 is: ", len(v1))
	// 5. access element
	fmt.Println("The first element in v1 is: ", v1[0])
	// 6. iterate the vector
	fmt.Print("[Version 1] The contents of v1 are:")
	for i := 0; i < len(v1); i++ {
		fmt.Print(" ", v1[i])
	}
	fmt.Println()
	fmt.Print("[Version 2] The contents of v1 are:")
	for _, item := range v1 {
		fmt.Print(" ", item)
	}
	fmt.Println()
	// 7. modify element

	v2[0] = 5
	fmt.Println("The first element in v1 is:", v1[0])
	v3[0] = -1
	fmt.Println("The first element in v1 is:", v1[0])

	// 8. sort
	sort.Ints(v1)
	// 9. add new element at the end of the vector
	v1 = append(v1, -1)
	v1 = append(v1, 1, 6)

	// 9. delete the last element
	v1 = v1[:len(v1)-1]

	fmt.Println(v0)

}
