package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	// head := constructor(1)
	node5 := ListNode{Val: 5, Next: nil}
	node4 := ListNode{Val: 4, Next: &node5}
	node3 := ListNode{Val: 3, Next: &node4}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}
	new := reverseList(&node1)
	for new.Next != nil {
		fmt.Println(new)
		new = new.Next
	}
	fmt.Println(new)
}

func reverseList(head *ListNode) *ListNode {
	end := head
	reverse := head
	if head == nil {
		return head
	}
	for end.Next != nil {
		end = end.Next
	}

	new := &ListNode{Val: end.Val, Next: nil}
	for reverse.Next != nil {
		buf := new.Next
		n := &ListNode{Val: reverse.Val, Next: buf}
		new.Next = n
		reverse = reverse.Next
	}

	return new
}
