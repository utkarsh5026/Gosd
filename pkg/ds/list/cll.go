package list

import (
	"errors"
	"fmt"
)

type CircularLinkedList struct {
	Head *Node
}

func NewCircularLinkedList() *CircularLinkedList {
	return &CircularLinkedList{}
}

func (cll *CircularLinkedList) insert(data interface{}) {
	newNode := &Node{Data: data}

	if cll.Head == nil {
		cll.Head = newNode
		newNode.Next = cll.Head
	} else {
		current := cll.Head
		for current.Next != cll.Head {
			current = current.Next
		}

		current.Next = newNode
		newNode.Next = cll.Head
	}
}

func (cll *CircularLinkedList) delete(data interface{}) error {
	if cll.Head == nil {
		return errors.New("List is empty")
	}

	if cll.Head.Data == data && cll.Head.Next == cll.Head {
		cll.Head = nil
		return nil
	}

	current := cll.Head
	var prev *Node

	for {
		if current.Data == data {
			if prev != nil {
				prev.Next = current.Next
			} else {
				current = cll.Head

				for current.Next != cll.Head {
					current = current.Next
				}

				cll.Head = current.Next
				current.Next = cll.Head
			}

			return nil
		}

		prev = current
		current = current.Next

		if current == cll.Head {
			break
		}
	}

	return errors.New("Value not found")
}

func (cll *CircularLinkedList) display() {
	if cll.Head == nil {
		fmt.Println("List is empty")
		return
	}

	current := cll.Head
	for {
		fmt.Print(current.Data, " ")
		current = current.Next

		if current == cll.Head {
			break
		}
	}

	fmt.Println()
}
