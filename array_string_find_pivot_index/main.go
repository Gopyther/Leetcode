package main

import "fmt"

func main() {
	nums := []int{-1, -1, 0, 1, 1, 0}
	fmt.Println(pivotIndex(nums))
}

func pivotIndex(nums []int) int {
	length := len(nums)
	right := make([]int, length+1)
	sum := 0

	for i := length - 1; i >= 1; i-- {
		sum += nums[i]
		right[i] = sum
	}
	sum = 0
	for j := 0; j < length; j++ {
		if sum == right[j+1] {
			return j
		}
		sum += nums[j]
	}
	return -1
}
