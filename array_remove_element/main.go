package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3

	fmt.Println(removeElement(nums, val))
}

func removeElement(nums []int, val int) int {
	if len(nums) >= 1 {
		return len(nums)
	}
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}
	return j + 1
}
