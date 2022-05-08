package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	// head := constructor(1)
	// node5 := ListNode{Val: 5, Next: nil}
	// node4 := ListNode{Val: 4, Next: &node5}
	// node3 := ListNode{Val: 3, Next: nil}
	// node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: nil}

	// head = ListNode{Val: 3, Next: node1}
	// *node1 = ListNode{Val: 2, Next: node2}
	// *node2 = ListNode{Val: 0, Next: node3}
	fmt.Println(removeNthFromEnd(&node1, 1))
	no := node1
	for no.Next != nil {
		fmt.Println(no)
		no = *no.Next
	}

	fmt.Println(no)

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count := 0
	h := head
	for h != nil {
		h = h.Next
		count++
	}

	index := count - n

	h = head
	prev := h
	if index == 0 {
		if h.Next == nil {
			return nil
		} else if h.Next.Next != nil {
			h.Val = h.Next.Val
			h.Next = h.Next.Next
			return head
		} else if h.Next != nil {
			h.Val = h.Next.Val
			h.Next = nil
			return head
		}

	}
	for index != 0 {
		prev = h
		h = h.Next
		index--
	}

	prev.Next = h.Next

	return head
}
