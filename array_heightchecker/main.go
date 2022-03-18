package main

import (
	"fmt"
)

func main() {
	heights := []int{2, 1, 2, 1, 1, 2, 2, 1}
	fmt.Println(heightChecker(heights))
}

func heightChecker(heights []int) int {
	heightfreq := make([]int, 101)
	for _, v := range heights {
		heightfreq[v]++
	}

	check := 0
	result := 0

	for _, v := range heights {
		for heightfreq[check] == 0 {
			check++
		}

		if check != v {
			result++
		}
		heightfreq[check]--

	}

	return result
}

// func heightChecker(heights []int) int {
// 	if len(heights) <= 1 {
// 		return 0
// 	}

// 	temp := make([]int, len(heights))
// 	copy(temp, heights)
// 	// for _, v := range heights {
// 	// 	temp = append(temp, v)
// 	// }

// 	quickSort(temp, 0, len(temp)-1)
// 	count := 0
// 	for i := range temp {
// 		if temp[i] != heights[i] {
// 			count++
// 		}
// 	}
// 	return count
// }

func quickSort(arr []int, l, r int) {
	if l < r {
		p := partition(arr, l, r)
		quickSort(arr, l, p-1)
		quickSort(arr, p+1, r)
	}
}

func partition(arr []int, l, r int) int {
	pivot := arr[r] // pivot
	i := (l - 1)

	for j := l; j <= r-1; j++ {

		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}
