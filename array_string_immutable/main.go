package main

import "fmt"

func main() {
	buf := []rune("Hello World")
	buf[5] = ','
	s := string(buf)
	fmt.Println(s) // "Hello,World"
}
