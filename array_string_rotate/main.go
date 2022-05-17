package main

import "fmt"

func main() {
	nums := []int{1, 2}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	length := len(nums)
	temp := make([]int, 0)
	if length-k >= 1 {
		temp = append(temp, nums[length-k:]...)
		temp = append(temp, nums[:length-k]...)
		copy(nums, temp)
		return
	} else {
		count := 0
		buf := 0
		for count < k {
			buf = nums[length-1]
			for i := length - 1; i > 0; i-- {
				nums[i] = nums[i-1]
			}
			nums[0] = buf
			count++
		}
		return
	}
}

// func rotate(nums []int, k int) {
// 	count := 0
// 	temp := 0
// 	length := len(nums)
// 	for count < k {
// 		temp = nums[length-1]
// 		for i := length - 1; i > 0; i-- {
// 			nums[i] = nums[i-1]
// 		}
// 		nums[0] = temp
// 		count++
// 	}
// }
