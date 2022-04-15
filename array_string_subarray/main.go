package main

import (
	"fmt"
	"math"
)

func main() {
	target := 11
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(minSubArrayLen(target, nums))
}

func minSubArrayLen(target int, nums []int) int {
	lp, rp := -1, -1
	sum := 0
	size := math.MaxInt

	for rp < len(nums) || lp < len(nums) {
		if sum < target && rp != len(nums)-1 {
			rp++
			sum += nums[rp]
		} else if sum >= target && lp < rp {
			lp++
			sum -= nums[lp]
		} else if sum < target && rp == len(nums) {
			break
		} else {
			break
		}
		if sum >= target {
			if size > rp-lp {
				size = rp - lp
			}
		}

	}
	if size == math.MaxInt {
		return 0
	}
	return size
}
