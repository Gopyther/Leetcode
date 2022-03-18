package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-3, -2, -1, 0}
	fmt.Println(thirdMax(nums))
}

func thirdMax(nums []int) int {
	sort.Ints(nums)
	max := make([]int, 0)

	i, last := len(nums)-1, 0
	for i >= 0 {
		if len(max) == 3 {
			break
		}
		if nums[i] != last {
			last = nums[i]
			max = append(max, last)
		}
		i--
	}

	if len(max) < 3 {
		return max[0]
	}
	return max[len(max)-1]
}

// func thirdMax(nums []int) int {
// 	first, second, third := -2147483649, -2147483649, -2147483649
// 	// rank := make([]int, 3)
// 	// third := rank[2]
// 	// second := rank[1]
// 	// first := rank[0]

// 	for _, v := range nums {
// 		if v > third && v != second && v != first {
// 			third = v
// 		}
// 		if third > second {
// 			second, third = third, second
// 		}
// 		if second > first {
// 			first, second = second, first
// 		}
// 	}

// 	if third != -2147483649 {
// 		return third
// 	}
// 	return first
// }

// func thirdMax(nums []int) int {
// 	r := len(nums)
// 	if r == 1 {
// 		return nums[0]
// 	} else if r == 2 {
// 		if nums[0] >= nums[1] {
// 			return nums[0]
// 		} else {
// 			return nums[1]
// 		}
// 	}
// 	quicksort(nums, 0, r-1)

// 	count := 1

// 	for i := r - 1; i > 0; i-- {
// 		if nums[i-1] < nums[i] {
// 			count++
// 		}

// 		if count == 3 {
// 			return nums[i-1]
// 		}

// 	}

// 	return nums[r-1]
// }

// func quicksort(arr []int, l, r int) {
// 	if l < r {
// 		p := partition(arr, l, r)

// 		quicksort(arr, l, p-1)
// 		quicksort(arr, p+1, r)
// 	}
// }

// func partition(arr []int, l, r int) int {
// 	i := l - 1
// 	for j := l; j <= r-1; j++ {
// 		if arr[j] <= arr[r] {
// 			i++
// 			arr[i], arr[j] = arr[j], arr[i]
// 		}

// 	}
// 	arr[i+1], arr[r] = arr[r], arr[i+1]
// 	return i + 1
// }
