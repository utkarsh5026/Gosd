// Package stack provides a Stack data structure implementation using a linked list.
package stack

import (
	"errors"
	"fmt"
	"strings"
)

// Node represents a node in the stack with data of any type and a pointer to the next node.
type Node struct {
	Data interface{}
	Next *Node
}

// StackLL represents a stack data structure. It has a pointer to the top node and a size.
type StackLL struct {
	Top  *Node
	Size int
}

// isEmpty checks if the stack is empty. It returns true if the stack is empty, false otherwise.
func (s *StackLL) isEmpty() bool {
	return s.Top == nil
}

// push adds a new node with the given data to the top of the stack.
func (s *StackLL) push(data interface{}) {
	node := &Node{Data: data}
	node.Next = s.Top
	s.Top = node
	s.Size++
}

// pop removes the top node from the stack and returns it. If the stack is empty, it returns an error.
func (s *StackLL) pop() (*Node, error) {
	if s.isEmpty() {
		return nil, errors.New("Empty Stack")
	}
	lastNode := s.Top
	s.Top = s.Top.Next
	s.Size--
	return lastNode, nil
}

// peek returns the top node from the stack without removing it. If the stack is empty, it returns an error.
func (s *StackLL) peek() (*Node, error) {
	if s.isEmpty() {
		return nil, errors.New("Empty Stack")
	}

	return s.Top, nil
}

// size returns the number of nodes in the stack.
func (s *StackLL) size() int {
	return s.Size
}

// String returns a string representation of the stack. It starts with "Stack[" and ends with "]".
// The data of each node is separated by ", ".
func (s *StackLL) String() string {
	var builder strings.Builder
	builder.WriteString("Stack[")
	current := s.Top

	for current != nil {
		builder.WriteString(fmt.Sprintf("%d", current.Data))

		if current.Next != nil {
			builder.WriteString(", ")
		}
		current = current.Next
	}

	builder.WriteString("]")
	return builder.String()
}
