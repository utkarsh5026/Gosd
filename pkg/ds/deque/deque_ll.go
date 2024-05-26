package deque

import (
	"errors"
	"fmt"
	"strings"
)

type Node[T any] struct {
	Data T
	Prev *Node[T]
	Next *Node[T]
}

type DequeLL[T any] struct {
	Front *Node[T]
	Back  *Node[T]
	Size  int
}

func NewDeque[T any]() *DequeLL[T] {
	return &DequeLL[T]{}
}

func (d *DequeLL[T]) IsEmpty() bool {
	return d.Size == 0
}

func (d *DequeLL[T]) Length() int {
	return d.Size
}

func (d *DequeLL[T]) AddFront(data T) {
	node := &Node[T]{Data: data}
	node.Next = d.Front

	if d.IsEmpty() {
		d.Back = node
	} else {
		d.Front.Prev = node
	}

	d.Front = node
	d.Size++
}

func (d *DequeLL[T]) AddBack(data T) {
	node := &Node[T]{Data: data}

	if d.IsEmpty() {
		d.Front = node
		d.Back = node
	} else {
		d.Back.Next = node
		node.Prev = d.Back
		d.Back = node
	}

	d.Size++
}

func (d *DequeLL[T]) PeekFront() (T, error) {
	var zeroValue T
	if d.IsEmpty() {
		return zeroValue, errors.New("Queue is empty")
	}

	return d.Front.Data, nil
}

func (d *DequeLL[T]) PeekLast() (T, error) {
	var zeroValue T

	if d.IsEmpty() {
		return zeroValue, errors.New("Queue is empty")
	}

	return d.Back.Data, nil
}

func (d *DequeLL[T]) RemoveFront() (T, error) {

	var zeroValue T

	if d.IsEmpty() {
		return zeroValue, nil
	}

	frontNode := d.Front
	d.Front = d.Front.Next

	if d.Front == nil {
		d.Back = nil
	} else {
		d.Front.Prev = nil
	}

	frontNode.Next = nil
	d.Size--
	return frontNode.Data, nil
}

func (d *DequeLL[T]) RemoveBack() (T, error) {

	var zeroValue T
	if d.IsEmpty() {
		return zeroValue, nil
	}

	backNode := d.Back
	d.Back = d.Back.Prev

	if d.Back == nil {
		d.Front = nil
	} else {
		d.Back.Next = nil
	}

	backNode.Prev = nil
	d.Size--
	return backNode.Data, nil
}

func (d *DequeLL[T]) String() string {
	var builder strings.Builder

	builder.WriteString("[")
	current := d.Front

	for current != nil {
		builder.WriteString(fmt.Sprintf("%v", current.Data))

		if current.Next != nil {
			builder.WriteString(", ")
		}

		current = current.Next
	}

	builder.WriteString("]")
	return builder.String()
}
