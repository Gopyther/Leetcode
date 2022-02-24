package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	count := 3

	// nums1 := []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	// nums2 := []int{1, 2, 2}
	// count := 6

	merge(nums1, count, nums2, len(nums2))
	fmt.Println(nums1)
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, 0)
	i := 0
	j := 0
	for i+j < m+n {
		if j >= n {
			temp = append(temp, nums1[i])
			i++
		} else if i >= m {
			temp = append(temp, nums2[j])
			j++
		} else if nums1[i] > nums2[j] {
			temp = append(temp, nums2[j])
			j++
		} else {
			temp = append(temp, nums1[i])
			i++
		}
	}
	if n == 0 {
		return
	}
	copy(nums1, temp)
}

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	temp := make([]int, 0)
// 	i := 0
// 	j := 0
// 	for m+n > i+j {
// 		if j >= n {
// 			temp = append(temp, nums1[i])
// 			i++
// 		} else if nums1[i] > nums2[j] || nums1[i] == 0 {
// 			temp = append(temp, nums2[j])
// 			j++
// 		} else {
// 			temp = append(temp, nums1[i])
// 			i++
// 		}
// 	}
// 	if n == 0 {
// 		return
// 	} else {
// 		for k := range nums1 {
// 			nums1[k] = temp[k]
// 		}
// 	}
// }

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	temp := make([]int, 0)
// 	i := 0
// 	j := 0
// 	for j < n {
// 		if nums1[i] > nums2[j] || nums1[i] == 0 {
// 			temp = append(temp, nums2[j])
// 			j++
// 		} else {
// 			temp = append(temp, nums1[i])
// 			i++
// 		}
// 	}
// 	if n == 0 {
// 		return
// 	} else {
// 		for k := i + 1; k < m+n; k++ {
// 			nums1[k] = temp[k]
// 		}
// 	}
// }

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	temp := make([]int, 0)
// 	if m > n {
// 		i := 0
// 		j := 0
// 		for i < m {
// 			if j == n {
// 				break
// 			}
// 			switch {
// 			case nums1[i] == 0 && nums2[j] != 0:
// 				temp = append(temp, nums2[j])
// 				i++
// 				j++
// 			case nums1[i] == 0:
// 				i++
// 			case nums2[j] == 0 && nums1[i] != 0:
// 				temp = append(temp, nums1[i])
// 				i++
// 				j++
// 			case nums2[j] == 0:
// 				j++
// 			case nums1[i] <= nums2[j]:
// 				temp = append(temp, nums1[i])
// 				i++
// 			case nums1[i] > nums2[j]:
// 				temp = append(temp, nums2[j])
// 				j++
// 			}
// 		}
// 	} else {
// 		i := 0
// 		j := 0
// 		for j < n {
// 			if i == m {
// 				break
// 			}
// 			switch {
// 			case nums1[i] == 0 && nums2[j] != 0:
// 				temp = append(temp, nums2[j])
// 				i++
// 				j++
// 			case nums1[i] == 0:
// 				i++
// 			case nums2[j] == 0 && nums1[i] != 0:
// 				temp = append(temp, nums1[i])
// 				i++
// 				j++
// 			case nums2[j] == 0:
// 				j++
// 			case nums1[i] <= nums2[j]:
// 				temp = append(temp, nums1[i])
// 				i++
// 			case nums1[i] > nums2[j]:
// 				temp = append(temp, nums2[j])
// 				j++
// 			}
// 		}

// 	}
// 	for k := range temp {
// 		nums1[k] = temp[k]
// 	}

// }

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	count := 0
// 	temp := append(nums1, nums2...)
// 	temp = quicksort(temp, 0, m+n-1)
// 	for i := range temp {
// 		if temp[i] == 0 {
// 			count++
// 		}
// 	}
// 	temp = temp[count:]

// 	for j := range nums1 {
// 		nums1[j] = temp[j] //self assignment for sure
// 	}

// }
// func quicksort(arr []int, low, high int) []int {
// 	if low < high {
// 		arr, p := partition(arr, low, high)
// 		arr = quicksort(arr, low, p-1)
// 		arr = quicksort(arr, p+1, high)
// 	}
// 	return arr
// }

// func partition(arr []int, low, high int) ([]int, int) {
// 	pivot := arr[high]
// 	i := low
// 	for j := low; j < high; j++ {
// 		if arr[j] < pivot {
// 			arr[i], arr[j] = arr[j], arr[i]
// 			i++
// 		}
// 	}
// 	arr[i], arr[high] = arr[high], arr[i]
// 	return arr, i
// }
