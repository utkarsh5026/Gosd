package list

import (
	"fmt"
)

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func NewLiknedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) addAtLast(data int) {
	newNode := &Node{Data: data}

	if ll.Tail != nil {
		ll.Tail.Next = newNode
	} else {
		ll.Head = newNode
	}

	ll.Tail = newNode
	ll.Size++
}

func (ll *LinkedList) addAtBegining(data interface{}) {

	newNode := &Node{Data: data}
	newNode.Next = ll.Head

	if ll.Tail == nil {
		ll.Tail = newNode
	}
	ll.Size++
}

func (ll *LinkedList) display() {

	for node := ll.Head; node != nil; node = node.Next {
		fmt.Print(node.Data)

		if node.Next != nil {
			fmt.Print(" -> ")
		}
	}

	fmt.Println()
}

func (ll *LinkedList) size() int {
	return ll.Size
}

func (ll *LinkedList) find(data interface{}) *Node {

	for node := ll.Head; node != nil; node = node.Next {
		if node.Data == data {
			return node
		}
	}

	return nil
}

func (ll *LinkedList) insert(pos int, data interface{}) bool {

	if pos < 0 || pos > ll.Size {
		return false
	}

	newNode := &Node{Data: data}

	if ll.Head == nil || pos == 0 {
		newNode.Next = ll.Head
		ll.Head = newNode

		if ll.Tail == nil {
			ll.Tail = newNode
		}
	} else {

		preVNode := ll.Head

		for i := 1; i < pos; i++ {
			preVNode = preVNode.Next
		}

		newNode.Next = preVNode.Next
		preVNode.Next = newNode

		if newNode.Next != nil {
			ll.Tail = newNode
		}
	}

	ll.Size++
	return true
}

func (ll *LinkedList) remove(data interface{}) error {

	if ll.Head == nil {
		return fmt.Errorf("List is currently empty")
	}

	if ll.Head.Data == data {
		ll.Head = ll.Head.Next

		if ll.Size == 1 {
			ll.Tail = nil
		}
		ll.Size--
		return nil
	}

	current := ll.Head

	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next

			if current.Next == nil {
				ll.Tail = current
			}

			ll.Size--
			return nil
		}
		current = current.Next
	}

	return fmt.Errorf("Value not found")
}
