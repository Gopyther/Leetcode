package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// head := constructor(1)
	node7 := ListNode{Val: 7, Next: nil}
	node6 := ListNode{Val: 7, Next: &node7}
	node5 := ListNode{Val: 7, Next: &node6}
	node4 := ListNode{Val: 7, Next: &node5}
	node3 := ListNode{Val: 7, Next: &node4}
	node2 := ListNode{Val: 7, Next: &node3}
	node1 := ListNode{Val: 7, Next: &node2}
	// node3.Next = &node2
	// head = ListNode{Val: 3, Next: node1}
	// *node1 = ListNode{Val: 2, Next: node2}
	// *node2 = ListNode{Val: 0, Next: node3}
	result := removeElements(&node1, 7)
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

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	for head.Val == val && head.Next != nil {
		head = head.Next
	}

	if head.Val == val {
		return nil
	}

	prev, node := head, head.Next

	for node != nil && prev != nil {
		if node.Val == val {
			prev.Next, node = node.Next, node.Next
		} else {
			prev, node = node, node.Next
		}
	}

	return head
}
