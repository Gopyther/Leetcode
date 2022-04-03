package main

import (
	"fmt"
)

func main() {
	var a, b string = "100", "110010"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	m := len(a)
	n := len(b)
	mn := max(m, n) + 1
	ans := make([]rune, mn)
	carry := 0

	for i := 1; i < mn; i++ {
		left, right := 0, 0
		if m-i >= 0 && a[m-i] == '1' {
			left = 1
		}
		if n-i >= 0 && b[n-i] == '1' {
			right = 1
		}
		remainder := (left + right + carry) % 2
		carry = (left + right + carry) / 2

		if remainder == 1 {
			ans[mn-i] = '1'
		} else {
			ans[mn-i] = '0'
		}
	}
	if carry == 1 {
		ans[0] = '1'
		return string(ans)
	}
	return string(ans[1:])

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func addBinary(a string, b string) string {
// 	aa := strings.Split(a, "")
// 	bb := strings.Split(b, "")
// 	ans := make([]string, 0)
// 	alen := len(aa) - 1
// 	blen := len(bb) - 1
// 	carry := 0

// 	for alen >= 0 && blen >= 0 {
// 		switch {
// 		case carry == 1:
// 			if aa[alen] == "0" && bb[blen] == "0" {
// 				ans = append([]string{"1"}, ans...)
// 				carry = 0
// 				alen--
// 				blen--
// 			} else if aa[alen] == "1" && bb[blen] == "0" || aa[alen] == "0" && bb[blen] == "1" {
// 				ans = append([]string{"0"}, ans...)
// 				carry = 1
// 				alen--
// 				blen--
// 			} else {
// 				ans = append([]string{"1"}, ans...)
// 				carry = 1
// 				alen--
// 				blen--
// 			}
// 		default:
// 			if aa[alen] == "0" && bb[blen] == "0" {
// 				ans = append([]string{"0"}, ans...)
// 				carry = 0
// 				alen--
// 				blen--
// 			} else if aa[alen] == "1" && bb[blen] == "0" || aa[alen] == "0" && bb[blen] == "1" {
// 				ans = append([]string{"1"}, ans...)
// 				carry = 0
// 				alen--
// 				blen--
// 			} else {
// 				ans = append([]string{"0"}, ans...)
// 				carry = 1
// 				alen--
// 				blen--
// 			}
// 		}
// 	}

// 	if alen >= 0 {
// 		for alen >= 0 {
// 			switch {
// 			case carry == 1:
// 				if aa[alen] == "1" {
// 					ans = append([]string{"0"}, ans...)
// 					carry = 1
// 					alen--
// 				} else if aa[alen] == "0" {
// 					ans = append([]string{"1"}, ans...)
// 					carry = 0
// 					alen--
// 				}
// 			default:
// 				for i := alen; i >= 0; i-- {
// 					ans = append([]string{aa[i]}, ans...)
// 					alen--
// 				}
// 			}
// 		}
// 	} else if blen >= 0 {
// 		for blen >= 0 {
// 			switch {
// 			case carry == 1:
// 				if bb[blen] == "1" {
// 					ans = append([]string{"0"}, ans...)
// 					carry = 1
// 					blen--
// 				} else if bb[blen] == "0" {
// 					ans = append([]string{"1"}, ans...)
// 					carry = 0
// 					blen--
// 				}
// 			default:
// 				for i := blen; i >= 0; i-- {
// 					ans = append([]string{bb[i]}, ans...)
// 					blen--
// 				}
// 			}
// 		}
// 	}
// 	if carry == 1 {
// 		ans = append([]string{"1"}, ans...)
// 	}

// 	return strings.Join(ans, "")

// }
