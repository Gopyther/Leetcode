package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	var s string = "Hello World"
// 	str := []rune(s)
// 	str[5] = ','
// 	fmt.Println(s)
// }

func main() {
	var n int = 10000
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString("hello")
	}
	var s string = sb.String()
	fmt.Println(s)
}
