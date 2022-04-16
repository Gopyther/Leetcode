package main

import "fmt"

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	l := len(s)
	words := []rune(s)
	for i := 0; i < l; i++ {
		j := i
		for j < l && words[j] != ' ' {
			j++
		}
		reverse(words, i, j-1)
		i = j
	}
	return string(words)
}

func reverse(words []rune, l, r int) {
	for l < r {
		words[l], words[r] = words[r], words[l]
		l++
		r--
	}
}
