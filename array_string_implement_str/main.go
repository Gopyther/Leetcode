package main

import (
	"fmt"
)

func main() {
	haystack, needle := "aaa", "aaa"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	needler := []rune(needle)
	haystacker := []rune(haystack)
	for i, v := range haystacker {
		if v == needler[0] {
			if len(needler) == 1 {
				return i
			}
			for j := 0; j < len(needler); j++ {
				if i+len(needler) > len(haystacker) {
					return -1
				}
				if haystacker[i+j] != needler[j] {
					break
				} else if j == len(needler)-1 {
					return i
				}
			}
		}

	}
	return -1
}

// func strStr(haystack string, needle string) int {
// 	return strings.Index(haystack, needle)
// }
