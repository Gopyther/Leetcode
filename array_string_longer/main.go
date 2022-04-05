package main

import "fmt"

func main() {
	strs := []string{"power", "water", "teller"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	shortest := 201
	short := ""
	longest := ""
	buf := ""
	for _, v := range strs {
		if len(v) < shortest {
			shortest = len(v)
			short = v
		}
	}
	for i := shortest; i >= 0; i-- {
		length := i
		right := 0

		for _, v := range strs {
			temp := v
			for j := 0; j < len(temp); j++ {
				if j+length >= len(temp) {
					break
				}
				if temp[j:j+length] == short[0:i] {
					right++
					buf = short[0:i]
				}
			}
			if right == len(strs) && len(longest) < len(buf) {
				longest = buf
			}
		}
	}

	for i := 0; i < shortest; i++ {
		length := shortest - i
		left := 0
		for _, v := range strs {
			temp := v
			for j := 0; j < len(temp); j++ {
				if j+length >= len(temp) {
					break
				}
				if temp[j:j+length] == short[i:shortest] {
					left++
					buf = short[i:shortest]
				}
				if left == len(strs) && len(longest) < len(buf) {
					longest = buf
				}
			}
		}
	}

	return longest
}
