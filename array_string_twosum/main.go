package main

import "fmt"

func main() {
	numbers := []int{48, 49, 51}
	target := 100
	fmt.Println(twoSum(numbers, target))

}

func twoSum(numbers []int, target int) []int {
	posByNum := make(map[int]int, len(numbers)/2)
	fmt.Println(len(numbers) / 2)
	for i, num := range numbers {
		if pos, ok := posByNum[target-num]; ok {
			return []int{pos + 1, i + 1}
		}

		if num <= target/2 {
			posByNum[num] = i
		}
	}

	return []int{0, 1}
}

// func twoSum(numbers []int, target int) []int {
// 	length := len(numbers)
// 	lp, rp := 0, length-1

// 	for lp < rp {
// 		if numbers[lp]+numbers[rp] == target {
// 			break
// 		}
// 		if numbers[lp]+numbers[rp] > target {
// 			rp--
// 		}
// 		if numbers[lp]+numbers[rp] < target {
// 			lp++
// 		}

// 	}

// 	return []int{lp + 1, rp + 1}
// }
