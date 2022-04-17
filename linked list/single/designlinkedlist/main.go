package main

import "fmt"

func main() {
	obj := Constructor()
	obj.AddAtHead(1)
	fmt.Println(obj)

}

type MyLinkedList struct {
	val  int
	next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{0, nil}
}

func (this *MyLinkedList) Get(index int) int {
	for i := 0; i < index; i++ {
		this = this.next
	}
	return this.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := *this
	*this = MyLinkedList{val, &node}
}

func (this *MyLinkedList) AddAtTail(val int) {
	var tail *MyLinkedList
	tail = this
	for tail.next != nil {
		tail = tail.next
	}

	node := &MyLinkedList{val: val}
	tail.next = node

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

}

func (this *MyLinkedList) DeleteAtIndex(index int) {

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
