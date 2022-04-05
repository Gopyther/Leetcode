package main

import "fmt"

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	shortest := 201
	short := ""
	for _, v := range strs {
		if len(v) < shortest {
			shortest = len(v)
			short = v
		}
	}
	for i := shortest; i >= 0; i-- {
		fmt.Println(short[0:i])

		for j := 0; j < len(strs); j++ {
			temp := strs[j]
			for k := len(temp); k >= 0; k-- {
				fmt.Println(temp[0:k])
			}
			if j+shortest > len(strs[j]) {
				break
			}
			if temp[j:j+shortest-i] == short[0:i] {
				fmt.Println(temp[j:j+shortest-i], "True")
			}
		}
	}

	return "ch"
}
