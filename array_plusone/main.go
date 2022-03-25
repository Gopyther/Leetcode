package main

import "fmt"

func main() {
	digits := []int{9,9,9}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	l := len(digits) + 1
	for _, d := range digits {
		if d < 9 {
			l--
			break
		}
	}
	sum := make([]int, l)
	c := 1
	for i, j := len(digits)-1, len(sum)-1; i >= 0; i-- {
		//fmt.Println(i, j, digits[i], sum[j],c)
		sum[j] = (digits[i] + c) % 10
		c = (digits[i] + c) / 10
		j--
	}
	if c > 0 {
		sum[0] = c
	}
	return sum
}

// func plusOne(digits []int) []int {
// 	length := len(digits)

// 	digits[length-1]++
// 	buffer := false

// 	for i := length - 1; i >= 0; i-- {
// 		if buffer {
// 			digits[i]++
// 			buffer = false
// 		}
// 		if digits[i] == 10 {
// 			digits[i] = 0
// 			buffer = true
// 		}
// 	}
// 	if buffer {
// 		temp := make([]int, 0)
// 		temp = append(temp, 1)
// 		temp = append(temp, digits...)
// 		return temp

// 	}

// 	return digits
// }

// func plusOne(digits []int) []int {
// 	length := len(digits)

// 	for i := length - 2; i >= 0; i-- {
// 		if digits[i+1] >= 9 {
// 			digits[i+1], digits[i] = 0, digits[i]+1
// 		} else {
// 			digits[i+1]++
// 			break
// 		}

// 	}
// 	fmt.Println(digits)

// 	if digits[0] == 10 {
// 		temp := make([]int, 1)
// 		temp[0], digits[0] = 1, 0
// 		temp = append(temp, digits...)
// 		return temp
// 	} else if digits[0] == 9 && length == 1 {
// 		temp := make([]int, 1)
// 		temp[0] = 1
// 		temp = append(temp, 0)
// 		return temp
// 	} else if length == 1 {
// 		digits[0]++
// 	}

// 	return digits
// }
