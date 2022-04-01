package main

import (
	"fmt"
	"strings"
)

func main() {
	// initialize
	var s1 string = "Hello World"
	fmt.Println("s1 is \"" + s1 + "\"")
	var s2 *string = &s1
	fmt.Println("s2 is another reference to s1")
	var s3 string = s1
	fmt.Println("s3 is a copy of s1")
	// compare using '=='
	fmt.Println("Compared by '==':")
	// true since string is immutable and s1 is binded to "Hello World"
	fmt.Println("s1 and Hello World:", (s1 == "Hello World"))
	// true since s1 and s2 is the reference of the same object
	fmt.Println("s1 and s2", (s1 == *s2))
	// true since s1 and s3 is the copy of the object.
	fmt.Println("s1 and s3", (s1 == s3))
	// Golang doesn't need to use function for checking strings.
	// compare returns when it is equal.
	fmt.Println("Compared by 'Compare':")
	fmt.Println("s1 and Hello World:", strings.Compare(s1, "Hello World"))
	fmt.Println("s1 and s2", strings.Compare(s1, *s2))
	fmt.Println("s1 and s3", strings.Compare(s1, s3))
}
