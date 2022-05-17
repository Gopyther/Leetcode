package main

import "fmt"

func main() {
	// node7 := ListNode{Val: 9, Next: nil}
	// node6 := ListNode{Val: 9, Next: &node7}
	// node5 := ListNode{Val: 9, Next: &node6}
	// node4 := ListNode{Val: 9, Next: &node5}
	// node3 := ListNode{Val: 9, Next: &node4}
	node2 := ListNode{Val: 7, Next: nil}
	node1 := ListNode{Val: 3, Next: &node2}

	// noded := ListNode{Val: 9, Next: nil}
	// nodec := ListNode{Val: 9, Next: &noded}
	nodeb := ListNode{Val: 2, Next: nil}
	nodea := ListNode{Val: 9, Next: &nodeb}

	// node3 := ListNode{Val: 9, Next: nil}
	// node2 := ListNode{Val: 9, Next: &node3}
	// node1 := ListNode{Val: 9, Next: &node2}

	// nodec := ListNode{Val: 9, Next: nil}
	// nodeb := ListNode{Val: 9, Next: &nodec}
	// nodea := ListNode{Val: 9, Next: &nodeb}

	result := addTwoNumbers(&node1, &nodea)

	for result.Next != nil {
		fmt.Println(result)
		result = result.Next
	}
	fmt.Println(result)
}

//ListNode is Definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	head := result
	round := 0

	for l1 != nil && l2 != nil {
		if l1.Val+l2.Val+round >= 10 {
			sum := l1.Val + l2.Val - 10 + round
			round = 1
			result.Next = &ListNode{Val: sum}
		} else {
			sum := l1.Val + l2.Val + round
			round = 0
			result.Next = &ListNode{Val: sum}
		}
		result = result.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum := l1.Val + round
		round = 0
		if sum >= 10 {
			sum = sum - 10
			round = 1

		}
		result.Next = &ListNode{Val: sum}
		result = result.Next
		l1 = l1.Next

	}

	for l2 != nil {
		sum := l2.Val + round
		round = 0
		if sum >= 10 {
			sum = sum - 10
			round = 1
		}
		result.Next = &ListNode{Val: sum}
		result = result.Next
		l2 = l2.Next
	}

	if result.Val >= 10 {
		result.Val -= 10
		round = 1
	}
	if round == 1 {
		result.Next = &ListNode{Val: 1}
	}

	return head.Next
}
