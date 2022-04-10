package main

import (
	"fmt"
)

func main() {
	nums := []int{6, 2, 6, 5, 1, 2}
	fmt.Println(arrayPairSum(nums))
}

func arrayPairSum(nums []int) int {
	arr := make([]int, 20001)
	for _, v := range nums {
		arr[v+10000]++
	}

	fmt.Println(arr)
	var sum int

	return sum
}

// func arrayPairSum(nums []int) int {
// 	length := len(nums)
// 	sum := 0
// 	maxi := 0
// 	counter := 0
// 	for length > 0 {
// 		max := math.MinInt
// 		for i := 0; i < length; i++ {
// 			if max <= nums[i] {
// 				max = nums[i]
// 				maxi = i
// 			}
// 		}

// 		nums = append(nums[:maxi], nums[maxi+1:]...)
// 		counter++
// 		length--
// 		if counter%2 == 0 {
// 			sum += max
// 		}

// 	}
// 	return sum
// }
