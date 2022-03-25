package main

import "fmt"

func main() {
	nums := []int{0, 0, 3, 2}
	fmt.Println(dominantIndex(nums))
}

func dominantIndex(nums []int) int {
	first := 0
	index := 0
	second := 0

	for i, v := range nums {
		if v > first {
			first, second, index = v, first, i
		} else if second < v {
			second = v
		}
	}
	if first < second*2 {
		return -1
	}
	return index
}
