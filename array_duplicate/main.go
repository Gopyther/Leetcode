package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(nums)
	fmt.Println(k, nums)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ch := make(chan int)
	go checker(nums, ch)
	k := <-ch
	return k
}

func checker(nums []int, ch chan int) {
	k := 0
	n := -101
	for _, v := range nums {
		if n != v {
			nums[k] = v
			k++
			n = v
		}
	}
	ch <- k
}
