package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "a good   example"
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	words := make([]string, 0)
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] == ' ' {
			continue
		}
		j := i
		for j < l && s[j] != ' ' {
			j++
		}
		words = append(words, s[i:j])
		i = j

	}
	return reverse(words)
}

func reverse(words []string) string {
	for lp, rp := 0, len(words)-1; lp < rp; lp, rp = lp+1, rp-1 {
		words[lp], words[rp] = words[rp], words[lp]
	}
	return strings.Join(words, " ")
}

// func reverseWords(s string) string {
// 	words := strings.Fields(s)

// 	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
// 		words[i], words[j] = words[j], words[i]
// 	}

// 	return strings.Join(words, " ")
// }

// func reverseWords(s string) string {
// 	buf := []rune(s)
// 	buf = frontspace(buf)
// 	buf = backspace(buf)
// 	buf = removespace(buf)

// 	l := len(buf)
// 	reverse(buf, 0, l)
// 	i := 0
// 	for j := 0; j < l; j++ {
// 		if string(buf[j]) == " " {
// 			reverse(buf, i, j)
// 			i = j + 1
// 			fmt.Println(buf, i, j)
// 		}
// 		if j == l-1 {
// 			reverse(buf, i, l)
// 		}

// 	}
// 	return string(buf)

// }

// func reverse(buf []rune, start, end int) {
// 	for start < end {
// 		end--
// 		buf[start], buf[end] = buf[end], buf[start]
// 		start++
// 	}

// }

// func frontspace(buf []rune) []rune {
// 	l := len(buf)
// 	temp := make([]rune, 0)
// 	for i := 0; i < l; i++ {
// 		if string(buf[i]) != " " {
// 			buf = append(buf[:0], buf[i:]...)
// 			break
// 		}
// 	}
// 	temp = buf
// 	return temp
// }

// func backspace(buf []rune) []rune {
// 	l := len(buf)
// 	temp := make([]rune, 0)
// 	for i := l - 1; i >= 0; i-- {
// 		if string(buf[i]) != " " {
// 			buf = buf[0 : i+1]
// 			break
// 		}
// 	}
// 	temp = buf
// 	return temp
// }

// func removespace(buf []rune) []rune {
// 	q := 0
// 	temp := make([]rune, 0)
// 	for {
// 		if q > len(buf)-1 {
// 			temp = buf
// 			return temp
// 		}
// 		if string(buf[q]) == " " && string(buf[q+1]) == " " {
// 			buf = append(buf[:q], buf[q+1:]...)
// 			continue
// 		}
// 		q++

// 	}
// }

// 착각해서 단어를 바꾸는 게 아니라 앞으로 보내는 것만 생각함
// func reverseWords(s string) string {
// 	buf := []rune(s)
// 	buf = frontspace(buf)
// 	buf = backspace(buf)
// 	buf = removespace(buf)

// 	l := len(buf)
// 	end := -1
// 	for i := l - 1; i >= 0; i-- {
// 		if string(buf[i]) == " " {
// 			end = i + 1
// 			break
// 		}
// 	}
// 	reverse(buf, 0, l)
// 	reverse(buf, 0, end)
// 	reverse(buf, end+1, l)
// 	return string(buf)

// }

// func reverse(buf []rune, start, end int) {
// 	for start < end {
// 		end--
// 		buf[start], buf[end] = buf[end], buf[start]
// 		start++
// 	}

// }

// func frontspace(buf []rune) []rune {
// 	l := len(buf)
// 	temp := make([]rune, 0)
// 	for i := 0; i < l; i++ {
// 		if string(buf[i]) != " " {
// 			buf = append(buf[:0], buf[i:]...)
// 			break
// 		}
// 	}
// 	temp = buf
// 	return temp
// }

// func backspace(buf []rune) []rune {
// 	l := len(buf)
// 	temp := make([]rune, 0)
// 	for i := l - 1; i >= 0; i-- {
// 		if string(buf[i]) != " " {
// 			buf = buf[0 : i+1]
// 			break
// 		}
// 	}
// 	temp = buf
// 	return temp
// }

// func removespace(buf []rune) []rune {
// 	q := 0
// 	temp := make([]rune, 0)
// 	for {
// 		if q > len(buf)-1 {
// 			temp = buf
// 			return temp
// 		}
// 		if string(buf[q]) == " " && string(buf[q+1]) == " " {
// 			buf = append(buf[:q], buf[q+1:]...)
// 			continue
// 		}
// 		q++

// 	}
// }
