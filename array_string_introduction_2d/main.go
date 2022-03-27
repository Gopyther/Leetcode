package main

import "fmt"

func printArrayA(a [2][5]int) { // When array is implemented, function also needs to get size of array.
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ { // array cannot have nil value
			fmt.Print(a[i][j], " ")
		}
		fmt.Println()
	}
}

func printArrayB(a [][]int) { // When array is implemented, function also needs to get size of array.
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	for i := 0; i < len(a); i++ {
		for j := 0; a[i] != nil && j < len(a[i]); j++ { // array cannot have nil value
			fmt.Print(a[i][j], " ")
		}
		fmt.Println()
	}
}
func main() {
	fmt.Println("Example I:")
	a := [2][5]int{}
	printArrayA(a)
	fmt.Println("Example II:")
	b := make([][]int, 2)
	printArrayB(b)
	fmt.Println("Example III:")
	b[0] = make([]int, 3)
	b[1] = make([]int, 5)
	printArrayB(b)
}
