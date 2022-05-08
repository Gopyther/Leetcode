package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node6 := ListNode{Val: 6, Next: nil}
	node5 := ListNode{Val: 5, Next: &node6}
	node4 := ListNode{Val: 4, Next: &node5}
	node3 := ListNode{Val: 3, Next: &node4}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}
	headA := ListNode{Val: 0, Next: &node1}
	nodeA := ListNode{Val: 13, Next: &node4}
	nodeB := ListNode{Val: 12, Next: &nodeA}
	headB := ListNode{Val: 11, Next: &nodeB}
	fmt.Println(getIntersectionNode(&headA, &headB))

	// this := headA
	// for this.Next != nil {
	// 	fmt.Println(this)
	// 	this = *this.Next
	// }
	// this = headB
	// for this.Next != nil {
	// 	fmt.Println(this)
	// 	this = *this.Next
	// }

}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	s := struct{}{}
	a := make(map[*ListNode]struct{})
	for {
		if headA != nil {
			if _, ok := a[headA]; ok {
				return headA
			} else {
				a[headA] = s
				headA = headA.Next
			}
		}
		if headB != nil {
			if _, ok := a[headB]; ok {
				return headB
			} else {
				a[headB] = s
				headB = headB.Next
			}
		}
		if headA == nil && headB == nil {
			break

		}
	}
	return nil
}

// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	s := struct{}{}
// 	a := make(map[*ListNode]struct{})
// 	for headA != nil {
// 		a[headA] = s
// 		headA = headA.Next
// 	}

// 	for headB != nil {
// 		if _, ok := a[headB]; ok {
// 			return headB
// 		}
// 		headB = headB.Next
// 	}
// 	return nil
// }
