package main

import "fmt"

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, 0, n)
	for _, v := range nums {
		result = append(result, v*v)
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

// func sortedSquares(nums []int) []int {
// 	for i := range nums {
// 		nums[i] = square(nums[i])
// 	}
// 	nums = bubble(nums)
// 	return nums
// }

// func square(n int) int {
// 	n = n * n
// 	return n
// }

// func bubble(nums []int) []int {
// 	for x := range nums {
// 		for y := 0; y < len(nums)-x; y++ {
// 			if y+1 > len(nums)-1 {
// 				break
// 			}
// 			if nums[y] > nums[y+1] {
// 				buf := nums[y+1]
// 				nums[y+1] = nums[y]
// 				nums[y] = buf
// 			}
// 		}
// 	}
// 	return nums
// }
