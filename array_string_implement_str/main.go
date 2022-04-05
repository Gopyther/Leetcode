package main

import (
	"fmt"
)

func main() {
	haystack, needle := "harrypotter", "pot"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	h := make(map[string]struct{})
	h[needle] = struct{}{}
	fmt.Println(h)

	index := 0

	for index < len(haystack) {
		if haystack[index] == needle[0] {
			if len(haystack) < index+len(needle) {
				return -1
			}
			tempS := haystack[index:(index + len(needle))]
			if _, ok := h[tempS]; ok {
				return index
			}
		}

		index++
	}

	return -1
}

// func strStr(haystack string, needle string) int {
// 	if len(haystack) < len(needle) {
// 		return -1
// 	}
// 	needler := []rune(needle)
// 	haystacker := []rune(haystack)
// 	for i, v := range haystacker {
// 		if v == needler[0] {
// 			if len(needler) == 1 {
// 				return i
// 			}
// 			for j := 0; j < len(needler); j++ {
// 				if i+len(needler) > len(haystacker) {
// 					return -1
// 				}
// 				if haystacker[i+j] != needler[j] {
// 					break
// 				} else if j == len(needler)-1 {
// 					return i
// 				}
// 			}
// 		}

// 	}
// 	return -1
// }

// func strStr(haystack string, needle string) int {
// 	return strings.Index(haystack, needle)
// }
