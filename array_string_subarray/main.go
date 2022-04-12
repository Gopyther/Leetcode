package main

import (
	"fmt"
	"math"
)

func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
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
		} else {
			break
		}
		if sum == target {
			fmt.Println(rp, lp)
			if size > 1+rp-lp {
				size = 1 + rp - lp
			}
		}

	}
	return size
}
