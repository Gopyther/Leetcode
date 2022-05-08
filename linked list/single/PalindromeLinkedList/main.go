package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// head := constructor(1)

	// node8 := ListNode{Val: 8, Next: nil}
	// node7 := ListNode{Val: 7, Next: &node8}
	// node6 := ListNode{Val: 6, Next: &node7}
	// node5 := ListNode{Val: 5, Next: &node6}
	// node6 := ListNode{Val: 1, Next: nil}
	// node5 := ListNode{Val: 2, Next: &node6}
	// node4 := ListNode{Val: 2, Next: &node5}
	node3 := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 0, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}

	result := isPalindrome(&node1)
	fmt.Println(result)
	// for result.Next != nil {
	// 	fmt.Println(result)
	// 	result = result.Next
	// }
	// if result != nil {
	// 	fmt.Println(result)
	// } else {
	// 	fmt.Println("nil")
	// }

}

func isPalindrome(head *ListNode) bool {
	count := 0

	node := head
	palindrome := []int{}

	for node != nil {
		palindrome = append(palindrome, node.Val)
		count++
		node = node.Next
	}
	if count == 1 {
		return true
	}

	mid := count / 2
	lp, rp := 0, 0
	if count%2 != 0 {
		lp, rp = mid-1, mid+1
	} else {
		lp, rp = mid-1, mid
	}
	for lp >= 0 && rp < count {
		if palindrome[lp] != palindrome[rp] {
			return false
		}
		lp--
		rp++

	}
	return true
}
