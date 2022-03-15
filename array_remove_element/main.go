package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3

	fmt.Println(removeElement(nums, val))
}

func removeElement(nums []int, val int) int {
	i := 0
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}
	return i
}
