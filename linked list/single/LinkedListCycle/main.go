package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := constructor(1)
	// node1 := &ListNode{}
	// node2 := &ListNode{}
	// node3 := &ListNode{Val: -4, Next: node1}
	// head = ListNode{Val: 3, Next: node1}
	// *node1 = ListNode{Val: 2, Next: node2}
	// *node2 = ListNode{Val: 0, Next: node3}
	fmt.Println(hasCycle(&head))

}

func constructor(val int) ListNode {
	return ListNode{}
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != nil && fast != nil && fast.Next != nil && slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow == fast
}
