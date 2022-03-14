package main

import "fmt"

func main() {
	nums := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(nums))
}

func sortArrayByParity(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	lp := 0
	rp := len(nums) - 1
	for {
		if lp >= rp {
			break
		}
		if nums[lp]%2 == 0 {
			lp++
		} else {
			nums[lp], nums[rp] = nums[rp], nums[lp]
		}
		if nums[rp]%2 != 0 {
			rp--
		} else {
			nums[lp], nums[rp] = nums[rp], nums[lp]
		}
	}
	return nums
}
