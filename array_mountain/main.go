package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(validMountainArray(arr))

}

func validMountainArray(arr []int) bool {
	l := len(arr)
	if l < 3 {
		return false
	}
	var top int = l
	for i := 1; i < l; i++ {
		if arr[i] > arr[i-1] {
			continue
		} else if arr[i] == arr[i-1] {
			return false
		} else {
			top = i - 1
			break
		}
	}
	for i := top + 1; i < l; i++ {
		if arr[i] < arr[i-1] {
			continue
		} else {
			return false
		}
	}
}

// func validMountainArray(arr []int) bool {
// 	left := make(chan int)
// 	right := make(chan int)
// 	res1 := make(chan bool)
// 	res2 := make(chan bool)
// 	length := len(arr)
// 	if length <= 2 {
// 		return false
// 	}

// 	go leftside(arr, left, res1)
// 	go rightside(arr, right, res2)
// 	l := <-left
// 	r := <-right
// 	c1 := <-res1
// 	c2 := <-res2

// 	return checker(length, l, r, c1, c2)
// }

// func leftside(arr []int, ch chan int, res1 chan bool) {
// 	l := 0
// 	for i := 0; i < len(arr)-1; i++ {
// 		if arr[i] < arr[i+1] {
// 			l++
// 		} else if arr[i] == arr[i+1] {
// 			ch <- l
// 			res1 <- false
// 		} else {
// 			ch <- l
// 			break
// 		}
// 	}
// 	res1 <- true
// 	ch <- l

// }

// func rightside(arr []int, ch chan int, res2 chan bool) {
// 	r := len(arr) - 1
// 	for j := len(arr) - 1; j > 0; j-- {
// 		if arr[j] < arr[j-1] {
// 			r--
// 		} else if arr[j] == arr[j-1] {
// 			ch <- r
// 			res2 <- false
// 		} else {
// 			ch <- r
// 			break
// 		}
// 	}
// 	res2 <- true
// 	ch <- r
// }

// func checker(le, l, r int, res1, res2 bool) bool {
// 	if res1 == false || res2 == false {
// 		return false
// 	}
// 	if l == r && r != le-1 && r != 0 {
// 		return true
// 	}
// 	return false
// }

// func validMountainArray(arr []int) bool {
// 	if len(arr) <= 2 {
// 		return false
// 	}
// 	l := 0
// 	r := len(arr) - 1

// 	for i := 0; i < len(arr)-1; i++ {
// 		if arr[i] < arr[i+1] {
// 			l++
// 		} else if arr[i] == arr[i+1] {
// 			return false
// 		} else {
// 			break
// 		}
// 	}
// 	for j := len(arr) - 1; j > 0; j-- {
// 		if arr[j] < arr[j-1] {
// 			r--
// 		} else if arr[j] == arr[j-1] {
// 			return false
// 		} else {
// 			break
// 		}
// 	}

// 	if l == r && r != len(arr)-1 && r != 0 {
// 		return true
// 	}

// 	return false
// }

// func validMountainArray(arr []int) bool {
// 	if len(arr) <= 2 {
// 		return false
// 	}
// 	max := arr[0]
// 	checker := false
// 	bnum := -1

// 	for i := 0; i < len(arr); i++ {
// 		if max < arr[i] {
// 			max = arr[i]
// 			checker = true
// 		} else {
// 			checker = false
// 		}
// 		switch {
// 		case bnum == arr[i]:
// 			return false
// 		case checker == true:
// 			if i == len(arr)-1 {
// 				return false
// 			}
// 		case checker == false:
// 			if bnum < arr[i] {
// 				return false
// 			} else if i == 1 {
// 				return false
// 			}

// 		}
// 		bnum = arr[i]
// 	}
// 	return true
// }
