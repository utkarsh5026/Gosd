package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []interface{}
}

func NewStack() *Stack {
	return &Stack{}
}

// IsEmpty returns true if the stack has no elements
func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

// // Push adds an element to the top of the stack
func (s *Stack) push(data interface{}) {
	s.items = append(s.items, data)
}

// Pop removes the top element of the stack and returns it
// If the stack is empty, it returns an error
func (s *Stack) pop() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, errors.New("Empty Stack")
	}

	lastIdx := len(s.items) - 1
	lastItem := s.items[lastIdx]
	s.items = s.items[:lastIdx]
	return lastItem, nil
}

// Peek returns the top element of the stack without removing it
// If the stack is empty, it returns an error
func (s *Stack) peek() (interface{}, error) {
	if s.isEmpty() {
		return nil, errors.New("Peeeking from an empty stack")
	}

	return s.items[len(s.items)-1], nil
}

// Size returns the number of elements in the stack
func (s *Stack) size() int {
	return len(s.items)
}

// Display prints all elements from the bottom to the top of the stack
func (s *Stack) display() {
	for _, item := range s.items {
		fmt.Println(item)
	}
}
