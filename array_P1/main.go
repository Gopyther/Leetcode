package main

import "fmt"

func main() {
	nums := []int{1, 0, 1, 1, 0, 1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	max := 0
	bnum := 1
	for index := range nums {
		if bnum == 1 {
			if nums[index] == 1 {
				count++
				bnum = 1
			} else {
				bnum = 0
			}
		} else {
			count = 0
			if nums[index] == 1 {
				count++
				bnum = 1
			}
		}
		if max < count {
			max = count
		}

	}
	return max
}

// if index+1 == len(nums) {
// 	break
// }
// if nums[index] == 1 && nums[index+1] == 1 {
// 	count++
// } else {
// 	if count > max {
// 		max = count
// 	}
// }
// }
// return count
