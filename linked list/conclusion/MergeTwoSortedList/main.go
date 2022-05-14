package main

import "fmt"

func main() {
	node3 := ListNode{Val: 4, Next: nil}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}

	nodec := ListNode{Val: 4, Next: nil}
	nodeb := ListNode{Val: 3, Next: &nodec}
	nodea := ListNode{Val: 1, Next: &nodeb}

	result := mergeTwoLists(&node1, &nodea)

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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	merge := &ListNode{}
	head := merge

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			merge.Next = list1
			merge = merge.Next
			list1 = list1.Next
		} else {
			merge.Next = list2
			merge = merge.Next
			list2 = list2.Next
		}
	}

	if list1 != nil {
		merge.Next = list1
	} else {
		merge.Next = list2
	}

	return head.Next
}
