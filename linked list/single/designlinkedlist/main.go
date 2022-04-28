package main

import "fmt"

func main() {
	obj := Constructor()
	obj.AddAtHead(3)
	obj.AddAtHead(2)
	obj.AddAtHead(1)
	obj.AddAtTail(99)

	fmt.Println("New", obj.Get(0))
	fmt.Println("New", obj.Get(1))
	fmt.Println("New", obj.Get(2))
	fmt.Println("New", obj.Get(3))

	obj.AddAtIndex(3, 0)
	obj.AddAtIndex(0, 0)
	fmt.Println("New", obj.Get(0))
	fmt.Println("New", obj.Get(1))
	fmt.Println("New", obj.Get(2))
	fmt.Println("New", obj.Get(3))
	fmt.Println("New", obj.Get(4))
	obj.DeleteAtIndex(0)
	fmt.Println("New del", obj.Get(0))
	fmt.Println("New", obj.Get(1))
	fmt.Println("New", obj.Get(2))
	fmt.Println("New", obj.Get(3))
	fmt.Println("New", obj.Get(4))
	fmt.Println("New", obj.Get(5))
	obj.DeleteAtIndex(0)
	fmt.Println("New del", obj.Get(0))
	fmt.Println("New", obj.Get(1))
	fmt.Println("New", obj.Get(2))
	fmt.Println("New", obj.Get(3))
	fmt.Println("New", obj.Get(4))
	fmt.Println("New", obj.Get(5))
	// fmt.Println(obj.Get(1))
	// fmt.Println(obj.Get(2))

	// obj.DeleteAtIndex(0)

	// fmt.Println("New", obj.Get(0))
	// fmt.Println(obj.Get(1))
	// fmt.Println(obj.Get(2))

	// obj.AddAtHead(7)
	// obj.AddAtHead(2)
	// obj.AddAtHead(1)
	// obj.AddAtIndex(3, 0)
	// fmt.Println("New", obj.Get(0))
	// fmt.Println(obj.Get(1))
	// fmt.Println(obj.Get(2))
	// fmt.Println(obj.Get(3))
	// obj.DeleteAtIndex(2)
	// fmt.Println("New", obj.Get(0))
	// fmt.Println(obj.Get(1))
	// fmt.Println(obj.Get(2))
	// fmt.Println(obj.Get(3))
	// obj.AddAtHead(6)
	// fmt.Println("New", obj.Get(0))
	// fmt.Println(obj.Get(1))
	// fmt.Println(obj.Get(2))
	// fmt.Println(obj.Get(3))
	// fmt.Println(obj.Get(4))
	// obj.AddAtTail(4)
	// fmt.Println(obj.Get(4))
	// obj.DeleteAtIndex(0)
	// obj.AddAtIndex(0, 1)
	// obj.AddAtTail(33)
	// obj.AddAtTail(34)
	// obj.AddAtHead(7)

	// fmt.Println(obj.Get(0))
	// fmt.Println(obj.Get(1))
	// fmt.Println(obj.Get(2))
	// fmt.Println(obj.Get(3))
	// obj.AddAtIndex(1, 2)
	// fmt.Println(obj.Get(0))
	// fmt.Println(obj.Get(1))
	// fmt.Println(obj.Get(2))
	// fmt.Println(obj.Get(3))
	// fmt.Println(obj.Get(4))
	// obj.AddAtIndex(5, 2)
	// fmt.Println(obj.Get(5))
	// fmt.Println(obj.Get(0))
	// fmt.Println(obj.Get(1))
	// obj.DeleteAtIndex(0)
	// fmt.Println(obj.Get(0))
	// fmt.Println("check", obj.Get(1))
	// fmt.Println("check2", obj.Get(2))
	// obj.DeleteAtIndex(1)
	// fmt.Println("check1", obj.Get(1))

	// fmt.Println(obj.Get(0))
	// obj.AddAtHead(7)
	// obj.AddAtHead(2)
	// obj.AddAtHead(1)
	// obj.AddAtIndex(3, 0)
	// fmt.Println(obj.Get(3))
	// fmt.Println(obj.Get(0))
	// fmt.Println(obj.Get(1))
	// fmt.Println(obj.Get(2))
	// obj.DeleteAtIndex(2)
	// fmt.Println(obj.Get(0))
	// fmt.Println(obj.Get(1))
	// fmt.Println(obj.Get(2))
	// obj.AddAtHead(6)
	// obj.AddAtTail(4)
	// fmt.Println(obj.Get(4))
	// fmt.Println(obj.Get(1))
	// fmt.Println(obj.Get(2))
	// fmt.Println(obj.Get(3))
	// obj.AddAtTail(3)
	// fmt.Println(obj.Get(1))
	// obj.AddAtIndex(1, 2)
	// fmt.Println(obj.Get(1))
	// obj.DeleteAtIndex(1)
	// fmt.Println(obj.Get(1))
	// obj.AddAtHead(2)
	// obj.AddAtHead(1)
	// obj.DeleteAtIndex(2)
	// fmt.Println(obj.Get(1))
	///////////////////////////
	// obj.AddAtHead(84)
	// obj.AddAtTail(2)
	// obj.AddAtTail(39)
	// fmt.Println(obj.Get(3))
	// fmt.Println(obj.Get(1))
	// obj.AddAtTail(42)
	// obj.AddAtIndex(1, 80)
	// obj.AddAtHead(14)
	// obj.AddAtHead(1)
	// obj.AddAtTail(53)
	// obj.AddAtTail(98)
	// obj.AddAtTail(19)
	// obj.AddAtTail(12)
	// fmt.Println(obj.Get(2))
	// obj.AddAtHead(16)
	// obj.AddAtHead(33)
	// obj.AddAtIndex(4, 17)
	// obj.AddAtIndex(6, 8)
	// obj.AddAtHead(37)
	// obj.AddAtTail(43)
	// obj.DeleteAtIndex(11)
	// obj.AddAtHead(80)
	// obj.AddAtHead(31)
	// obj.AddAtIndex(13, 23)
	// obj.AddAtTail(17)
	// fmt.Println(obj.Get(4))
	// obj.AddAtIndex(10, 0)
	// obj.AddAtTail(21)
	// obj.AddAtHead(73)
	// obj.AddAtHead(22)
	// obj.AddAtIndex(24, 37)
	// obj.AddAtTail(14)
	// obj.AddAtHead(97)
	// obj.AddAtHead(8)
	// fmt.Println(obj.Get(6))
	// obj.DeleteAtIndex(17)
	// obj.AddAtTail(50)
	// obj.AddAtTail(28)
	// obj.AddAtHead(76)
	// obj.AddAtTail(79)
	// fmt.Println(obj.Get(18))
	// obj.DeleteAtIndex(30)
	// fmt.Println("호출")
	// obj.AddAtTail(5)

	// obj.AddAtHead(9)
	// obj.AddAtTail(83)
	// obj.DeleteAtIndex(3)
	// obj.AddAtTail(40)
	// obj.DeleteAtIndex(26)
	// obj.AddAtIndex(20, 90)
	// obj.DeleteAtIndex(30)
	// obj.AddAtTail(40)
	// obj.AddAtHead(56)
	// obj.AddAtIndex(15, 23)
	// obj.AddAtHead(51)
	// obj.AddAtHead(21)
	// fmt.Println(obj.Get(26))
	// obj.AddAtHead(83)
	// fmt.Println(obj.Get(30))
	// obj.AddAtHead(12)
	// obj.DeleteAtIndex(8)
	// fmt.Println(obj.Get(4))
	// obj.AddAtHead(20)
	// obj.AddAtTail(45)
	// fmt.Println(obj.Get(10))
	// obj.AddAtHead(56)
	// fmt.Println(obj.Get(18))
	// obj.AddAtTail(33)
	// fmt.Println(obj.Get(2))
	// obj.AddAtTail(70)
	// obj.AddAtHead(57)
	// obj.AddAtIndex(31, 24)
	// obj.AddAtIndex(16, 92)
	// obj.AddAtHead(40)
	// obj.AddAtHead(23)
	// obj.DeleteAtIndex(26)
	// fmt.Println(obj.Get(1))
	// obj.AddAtHead(92)
	// obj.AddAtIndex(3, 78)
	// obj.AddAtTail(42)
	// fmt.Println(obj.Get(18))
	// obj.AddAtIndex(39, 9)
	// fmt.Println(obj.Get(13))
	// obj.AddAtIndex(33, 17)
	// fmt.Println(obj.Get(51))
	// obj.DeleteAtIndex(1)

	// obj.DeleteAtIndex(2)

	// fmt.Println(obj)

	// fmt.Println(obj.Get(0))
	// fmt.Println(obj.Get(1))
	//
	// fmt.Println(obj.Get(1))
	// fmt.Println(obj.Get(2))
	// fmt.Println(obj.Get(4))
	// obj.DeleteAtIndex(0)
	// fmt.Println(obj.Get(0))
	//
	// fmt.Println(obj.Get(1))
	// obj.AddAtIndex(1, 2)

	// fmt.Println(obj.Get(1))
	// 	for obj.next != nil {
	// 		fmt.Println(obj)
	// 		obj = *obj.next
	// 	}
}

type MyLinkedList struct {
	val  int
	next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{val: -1, next: nil}
}

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

func (this *MyLinkedList) AddAtHead(val int) {
	node := &MyLinkedList{val: val, next: this.next}
	this.next = node
}

func (this *MyLinkedList) AddAtTail(val int) {
	h := this
	for h.next != nil {
		h = h.next
	}
	h.next = &MyLinkedList{val: val, next: nil}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	prev, h := this, this.next
	for h != nil && index > 0 {
		prev, h = h, h.next
		index--
	}
	if index == 0 {
		prev.next = &MyLinkedList{val: val, next: h}
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
	}
}

// type MyLinkedList struct {
// 	next *MyLinkedList
// 	val  int
// }

// func Constructor() MyLinkedList {
// 	return MyLinkedList{val: -1}
// }

// func (this *MyLinkedList) Get(index int) int {
// 	node := this
// 	for i := 0; i < index; i++ {
// 		if node.next == nil {
// 			return -1
// 		}
// 		node = node.next
// 	}
// 	return node.val
// }

// func (this *MyLinkedList) AddAtHead(val int) {
// 	if this.val == -1 {
// 		this.val = val
// 		return
// 	}
// 	node := &MyLinkedList{val: this.val, next: this.next}
// 	*this = MyLinkedList{val: val, next: node}
// }

// func (this *MyLinkedList) AddAtTail(val int) {
// 	node := this
// 	if this.val == -1 {
// 		this.val = val
// 		return
// 	}

// 	for node.next != nil {
// 		node = node.next
// 	}
// 	tail := &MyLinkedList{val: val}
// 	node.next = tail
// }

// func (this *MyLinkedList) AddAtIndex(index int, val int) {
// 	node := this
// 	if this.val == -1 {
// 		this.val = val
// 		return
// 	}
// 	for i := 0; i < index; i++ {
// 		if node.next == nil && index-1 == i {
// 			tail := &MyLinkedList{val: val}
// 			node.next = tail
// 		} else if node.next == nil {
// 			return
// 		}
// 		node = node.next
// 	}
// 	new := &MyLinkedList{val: node.val, next: node.next}
// 	*node = MyLinkedList{val: val, next: new}
// }

// func (this *MyLinkedList) DeleteAtIndex(index int) {
// 	prev, h := &MyLinkedList{}, this
// 	for i := 0; i < index; i++ {
// 		prev = h
// 		h = h.next
// 	}

// 	if index == 0 && h.next != nil {
// 		*this = MyLinkedList{val: h.next.val, next: h.next.next}
// 	} else if h != nil {
// 		prev.next = h.next
// 	}

// }

/*
type MyLinkedList struct {
	next *MyLinkedList
	val  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{val: -1}
}

func (this *MyLinkedList) Get(index int) int {
	node := this
	for i := 0; i < index; i++ {
		if node.next != nil {
			node = node.next
		} else {
			return -1
		}

	}
	return node.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := *this
	*this = MyLinkedList{val: val, next: &node}
}

func (this *MyLinkedList) AddAtTail(val int) {
	for this.next != nil {
		this = this.next
	}
	new := MyLinkedList{val: this.val}
	*this = MyLinkedList{val: val, next: &new}

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 && this.val != -1 {
		this.AddAtHead(val)
		return
	} else if index == 1 && this.val == -1 {
		return
	}
	node := this
	for i := 0; i < index; i++ {
		// if reflect.DeepEqual(*node, MyLinkedList{}) {
		// 	return
		// }
		if node.next == nil && index-1 == i {
			this.AddAtTail(val)
			return
		} else if node.next == nil {
			return
		}

		node = node.next
	}

	new := *node
	new = MyLinkedList{node.next, node.val}
	*node = MyLinkedList{&new, val}

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	node := this
	if index == 0 && node.next == nil {
		*this = MyLinkedList{val: -1}
		return
	}
	for i := 0; i < index; i++ {
		if node.next != nil {
			node = node.next
		} else {
			return
		}

	}
	if node.next == nil {
		*node = MyLinkedList{val: -1}
		return
	}
	if node.next.next != nil {
		*node = MyLinkedList{node.next.next, node.next.val}
	} else {
		*node = MyLinkedList{val: node.next.val}
	}

}

*/
/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
