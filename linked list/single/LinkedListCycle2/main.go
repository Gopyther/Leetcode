package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// head := constructor(1)
	node3 := ListNode{Val: 3, Next: nil}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}
	node3.Next = &node2
	// head = ListNode{Val: 3, Next: node1}
	// *node1 = ListNode{Val: 2, Next: node2}
	// *node2 = ListNode{Val: 0, Next: node3}
	fmt.Println(detectCycle(&node1))

}

// func constructor(val int) ListNode {
// 	return ListNode{}
// }

func detectCycle(head *ListNode) *ListNode {
	var s struct{}
	if head == nil {
		return nil
	}
	slow := head
	a := make(map[*ListNode]struct{})
	for slow.Next != nil {
		if _, ok := a[slow]; !ok {
			a[slow] = s
		} else {
			return slow
		}
		slow = slow.Next
	}
	return nil
}

// func detectCycle(head *ListNode) *ListNode {
// 	slow, fast := head, head
// 	hasCycle := false
// 	for slow != nil && fast != nil && fast.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 		if slow == fast {
// 			hasCycle = true
// 			break
// 		}
// 	}

// 	if hasCycle {
// 		slow = head
// 	} else {
// 		return nil
// 	}
// 	for slow != fast {
// 		slow = slow.Next
// 		fast = fast.Next
// 	}
// 	return slow
// }
