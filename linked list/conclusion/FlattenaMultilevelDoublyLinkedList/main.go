package main

import "fmt"

func main() {
	node12 := Node{Val: 12, Prev: nil, Next: nil}
	node11 := Node{Val: 11, Prev: nil, Next: &node12, Child: nil}
	node10 := Node{Val: 10, Prev: nil, Next: nil, Child: nil}
	node9 := Node{Val: 9, Prev: nil, Next: &node10, Child: nil}
	node8 := Node{Val: 8, Prev: nil, Next: &node9, Child: &node11}
	node7 := Node{Val: 7, Prev: nil, Next: &node8, Child: nil}
	node6 := Node{Val: 6, Prev: nil, Next: nil, Child: nil}
	node5 := Node{Val: 5, Prev: nil, Next: &node6, Child: nil}
	node4 := Node{Val: 4, Prev: nil, Next: &node5, Child: nil}
	node3 := Node{Val: 3, Prev: nil, Next: &node4, Child: &node7}
	node2 := Node{Val: 2, Prev: nil, Next: &node3, Child: nil}
	node1 := Node{Val: 1, Prev: nil, Next: &node2, Child: nil}
	node2.Prev = &node1
	node3.Prev = &node2
	node4.Prev = &node3
	node5.Prev = &node4
	node6.Prev = &node5
	node8.Prev = &node7
	node9.Prev = &node8
	node10.Prev = &node9
	node12.Prev = &node11

	result := flatten(&node1)

	for result.Next != nil {
		fmt.Println(result)
		result = result.Next
	}
	fmt.Println(result)
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}
	head := root
	end := &Node{}

	for root.Next != nil {
		if root.Child != nil {
			end = flatten(root.Child)
			root.Next.Prev = end
			end.Next = root.Next
			root.Next = root.Child
			root.Next.Prev = root
			root.Child = nil
		}
		root = root.Next
	}

	if root.Child != nil {
		root.Next = root.Child
		root.Next.Prev = root
		root.Child = nil
		return head
	}
	if head.Val != 1 {
		return root
	}
	return head

}
