package main

import "fmt"

func main() {
	arr := []int{1, 5, 2, 0, 6, 8, 0, 6, 0}
	// arr := []int{1, 1, 1, 1, 1, 1, 1, 0, 1}
	duplicateZeros(arr)
	fmt.Println(arr)
}

// func duplicateZeros(arr []int) {
// 	for x := 0; x < len(arr)-1; x++ {
// 		if arr[x] == 0 {
// 			arr = shift(arr, x)
// 			arr[x+1] = 0
// 			x++
// 		}
// 	}
// }

// func shift(arr []int, x int) []int {
// 	for y := len(arr) - 2; y >= x+1; y-- {
// 		arr[y+1] = arr[y]
// 	}
// 	return arr
// }

// func duplicateZeros(arr []int) {
// 	temp := make([]int, len(arr))
// 	t := 0
// 	a := 0

// 	for t < len(arr) {
// 		if arr[a] != 0 {
// 			temp[t] = arr[a]
// 			t += 1
// 			a += 1
// 		} else {
// 			t += 2
// 			a += 1
// 		}
// 	}
// 	copy(arr, temp)
// }

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			for j := len(arr) - 1; j > i; j-- {
				arr[j] = arr[j-1]
			}
			i++
		}
	}
}
