package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 1, 7, 11}
	fmt.Println(checkIfExist(arr))

}

// func checkIfExist(arr []int) bool {
// 	e := len(arr)
// 	for l := 0; l < e; l++ {
// 		for r := e - 1; r > l; r-- {
// 			switch {
// 			case arr[l]*2 == arr[r] || arr[r]*2 == arr[l]:
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
func checkIfExist(arr []int) bool {
	container := make(map[int]struct{})

	for i := 0; i < len(arr); i++ {
		if _, ok := container[arr[i]*2]; ok {
			return true
		} else if _, ok := container[arr[i]/2]; ok && arr[i]%2 == 0 {
			return true
		} else {
			container[arr[i]] = j{}
		}
	}
	return false
}

type j struct{}

// func checkIfExist(arr []int) bool {

// 	ch := make(chan bool)
// 	go checker(arr, ch)

// 	return <-ch

// }

// func checker(arr []int, ch chan bool) {
// 	e := len(arr)
// 	for l := 0; l < e; l++ {
// 		for r := e - 1; r > l; r-- {
// 			switch {
// 			case arr[l] == 0 && arr[r] == 0:
// 				ch <- true
// 			case arr[l] == 0 && arr[r] != 0 || arr[r] == 0 && arr[l] != 0:
// 				break
// 			case arr[l]%arr[r] == 0 || arr[r]%arr[l] == 0:
// 				if arr[l]/arr[r] == 2 || arr[r]/arr[l] == 2 {
// 					ch <- true
// 				}
// 			}
// 		}
// 	}
// 	ch <- false

// }
