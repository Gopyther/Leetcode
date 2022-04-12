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
	lp, rp := 0, 0
	sum := 0
	size := math.MaxInt

	for rp < len(nums) || lp < len(nums) {
		if sum < target {

		}

	}
	return 0
}
