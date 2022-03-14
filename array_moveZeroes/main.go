package main

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}

// func moveZeroes(nums []int) {
// 	if len(nums) <= 1 {
// 		return
// 	}
// 	wp := 0
// 	for j, v := range nums {
// 		if j == wp && v != 0 {
// 			wp++
// 		} else if v != 0 {
// 			nums[wp], nums[j] = v, 0
// 			wp++
// 		}
// 	}
// 	fmt.Println(nums)
// }

// func moveZeroes(nums []int) {
// 	i := 0
// 	for _, num := range nums {
// 		if num != 0 {
// 			nums[i] = num
// 			i++
// 		}
// 	}

// 	for j := i; j < len(nums); j++ {
// 		nums[j] = 0
// 	}
// }

func moveZeroes(nums []int) {
	n := 0
	for i, elem := range nums {
		if elem != 0 {
			nums[i], nums[n] = 0, elem
			n++
		}
	}
}
