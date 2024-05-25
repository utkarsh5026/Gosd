package list

import "fmt"

type DllNode struct {
	Data  interface{}
	Left  *DllNode
	Right *DllNode
}

type DoubleLinkedList struct {
	Head *DllNode
	Tail *DllNode
	Size int
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

func (dll *DoubleLinkedList) addAtEnd(data interface{}) {
	newNode := &DllNode{Data: data}

	if dll.Tail == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Left = dll.Tail
		dll.Tail.Right = newNode
		dll.Tail = newNode
	}

	dll.Size++
}

func (dll *DoubleLinkedList) addAtBegining(data interface{}) {

	newNode := &DllNode{Data: data}

	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {

		newNode.Right = dll.Head
		dll.Head.Left = newNode
		dll.Head = newNode
	}

	dll.Size++
}

func (dll *DoubleLinkedList) remove(data interface{}) error {
	current := dll.Head

	for current != nil {
		if current.Data == data {
			if current.Left != nil {
				current.Left.Right = current.Right
			} else {
				dll.Head = current.Right
			}

			if current.Right != nil {
				current.Right.Left = current.Left
			} else {
				dll.Tail = current.Left
			}

			dll.Size--
			return nil
		}

		current = current.Right
	}

	return fmt.Errorf("Value not Found")
}

func (dll *DoubleLinkedList) display() {

	for node := dll.Head; node != nil; node = node.Right {
		fmt.Print(node.Data)
		if node.Right != nil {
			fmt.Print(" < - >")
		}
	}

	fmt.Println()
}

func (dll *DoubleLinkedList) size() int {
	return dll.Size
}

func (dll *DoubleLinkedList) insert(pos int, data interface{}) bool {

	if pos < 0 || pos > dll.Size {
		return false
	}

	newNode := &DllNode{Data: data}

	if dll.Head == nil || pos == 0 {
		newNode.Right = dll.Head

		if dll.Head != nil {
			dll.Head.Left = newNode
		}
		dll.Head = newNode

		if dll.Tail == nil {
			dll.Tail = newNode
		}
	} else {

		current := dll.Head

		for i := 1; i < pos; i++ {
			current = current.Right
		}

		newNode.Right = current.Right
		newNode.Left = current

		if current.Right != nil {
			current.Right.Left = newNode
		} else {
			dll.Tail = newNode
		}

		current.Right = newNode
	}

	dll.Size++
	return true
}
