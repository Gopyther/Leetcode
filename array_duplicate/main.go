package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(nums)
	fmt.Println(k, nums)
}

func removeDuplicates(nums []int) int {
	idx := 0
	for i := 0; i < len(nums); i++ {
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
		nums[idx] = nums[i]
		idx++
	}
	return idx
}
