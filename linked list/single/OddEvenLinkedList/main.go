package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// head := constructor(1)

	node8 := ListNode{Val: 8, Next: nil}
	node7 := ListNode{Val: 7, Next: &node8}
	node6 := ListNode{Val: 6, Next: &node7}
	node5 := ListNode{Val: 5, Next: &node6}
	node4 := ListNode{Val: 4, Next: &node5}
	node3 := ListNode{Val: 3, Next: &node4}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}

	result := oddEvenList(&node1)
	for result.Next != nil {
		fmt.Println(result)
		result = result.Next
	}
	if result != nil {
		fmt.Println(result)
	} else {
		fmt.Println("nil")
	}

}
func oddEvenList(head *ListNode) *ListNode {
	var odd *ListNode
	var even *ListNode
	var e *ListNode

	if head != nil {
		odd = head
		even, e = head.Next, head.Next
	} else {
		return head
	}

	for odd != nil && even != nil && odd.Next != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = e
	return head
}
