package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string = "Hello World"
	// 1. concatenate
	s1 += "!"
	fmt.Println(s1)
	// 2. find
	fmt.Println("The position of first 'o' is: ", strings.Index(s1, "o"))
	fmt.Println("The position of last 'o' is: ", strings.LastIndex(s1, "o"))
	// 3. get substring
	fmt.Println(s1[6:11])
}
