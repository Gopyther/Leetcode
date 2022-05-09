package main

import "fmt"

func main() {
	obj := Constructor()
	obj.AddAtIndex(0, 1)
	// obj.AddAtHead(3)
	// obj.AddAtHead(2)
	// obj.AddAtHead(1)
	// obj.AddAtTail(4)
	fmt.Println(obj.Get(0))
}

//MyLinkedList is struct for Linked List
type MyLinkedList struct {
	val        int
	prev, next *MyLinkedList
}

//Constructor is constructing Linked List
func Constructor() MyLinkedList {
	return MyLinkedList{val: -1, prev: nil, next: nil}
}

//Get is for used get value of linked list
func (this *MyLinkedList) Get(index int) int {
	h := this
	for h != nil && index >= 0 {
		h = h.next
		index--
	}

	if h != nil {
		return h.val
	}

	return -1
}

//AddAtHead is adding linked list to head
func (this *MyLinkedList) AddAtHead(val int) {
	node := this
	new := &MyLinkedList{val: val, prev: node, next: node.next}
	if node.next != nil {
		node.next.prev = new
	}
	node.next = new

}

func (this *MyLinkedList) AddAtTail(val int) {
	node := this
	for node.next != nil {
		node = node.next
	}
	node.next = &MyLinkedList{val: val, prev: node}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	prev, h := this, this.next
	for h != nil && index > 0 {
		prev, h = h, h.next
		index--
	}
	if index == 0 {
		prev.next = &MyLinkedList{val: val, prev: prev, next: h}
	}

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	prev, h := this, this.next
	for h != nil && index > 0 {
		prev, h = h, h.next
		index--
	}

	if index == 0 && h != nil {
		prev.next = h.next
		prev = h.prev
	}
}
