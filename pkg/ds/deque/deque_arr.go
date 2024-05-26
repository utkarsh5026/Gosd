// Package deque provides a generic implementation of a double-ended queue (deque) using a circular array.
package deque

import (
	"errors"
	"fmt"
	"strings"
)

// DequeArr is a generic double-ended queue (deque) implemented using a circular array.
// It holds elements of any type T, and maintains indices for the front and back of the queue,
// as well as the current size and capacity of the underlying array.
type DequeArr[T any] struct {
	elements []T
	Front    int
	Back     int
	Size     int
	Capacity int
}

// NewDequeArr creates a new deque with the given initial capacity.
func NewDequeArr[T any](capacity int) *DequeArr[T] {
	return &DequeArr[T]{
		elements: make([]T, capacity),
		Front:    0,
		Back:     0,
		Size:     0,
		Capacity: capacity,
	}
}

// IsEmpty checks whether the deque is empty.
func (d *DequeArr[T]) IsEmpty() bool {
	return d.Size == 0
}

// Resize doubles the capacity of the deque's underlying array.
// If the current capacity is 0, it is set to 1.
// The elements are copied to the new array in their original order, and the Front and Back indices are updated.
func (d *DequeArr[T]) Resize() {
	newCapacity := d.Capacity * 2
	if newCapacity == 0 {
		newCapacity = 1
	}

	resized := make([]T, newCapacity)
	for i := 0; i < d.Size; i++ {
		idx := (d.Front + i) % d.Capacity
		resized[i] = d.elements[idx]
	}

	d.elements = resized
	d.Front = 0
	d.Back = d.Size
	d.Capacity = newCapacity
}

// AddFront adds an element to the front of the deque.
// If the deque is full, it is resized before the element is added.
func (d *DequeArr[T]) AddFront(data T) {
	if d.Size == d.Capacity {
		d.Resize()
	}

	d.Front = (d.Front - 1 + d.Capacity) % d.Capacity
	d.elements[d.Front] = data
	d.Size++
}

// AddBack adds an element to the back of the deque.
// If the deque is full, it is resized before the element is added.
func (d *DequeArr[T]) AddBack(data T) {
	if d.Size == d.Capacity {
		d.Resize()
	}

	d.elements[d.Back] = data
	d.Back = (d.Back + 1) % d.Capacity
	d.Size++
}

// PeekFront returns the element at the front of the deque without removing it.
// If the deque is empty, it returns an error.
func (d *DequeArr[T]) PeekFront() (T, error) {
	var zeroValue T
	if d.IsEmpty() {
		return zeroValue, errors.New("Empty Queue")
	}

	return d.elements[d.Front], nil
}

// PeekBack returns the element at the back of the deque without removing it.
// If the deque is empty, it returns an error.
func (d *DequeArr[T]) PeekBack() (T, error) {
	var zeroValue T

	if d.IsEmpty() {
		return zeroValue, errors.New("Queue is empty")
	}

	return d.elements[d.Back], nil
}

// RemoveBack removes and returns the element at the back of the deque.
// If the deque is empty, it returns an error.
func (d *DequeArr[T]) RemoveBack() (T, error) {
	var zeroValue T

	if d.IsEmpty() {
		return zeroValue, errors.New("Queue is empty")
	}
	back := d.elements[d.Back]
	d.Back = (d.Back - 1 + d.Capacity) % d.Capacity
	d.Size--
	return back, nil
}

// RemoveFront removes and returns the element at the front of the deque.
// If the deque is empty, it returns an error.
func (d *DequeArr[T]) RemoveFront() (T, error) {
	var zeroValue T

	if d.IsEmpty() {
		return zeroValue, errors.New("Queue is empty")
	}

	front := d.elements[d.Front]
	d.Front = (d.Front + 1) % d.Capacity
	d.Size--
	return front, nil
}

// String returns a string representation of the deque.
// The elements are listed in order from front to back, separated by commas and enclosed in square brackets.
func (d *DequeArr[T]) String() string {
	var builder strings.Builder
	builder.WriteString("[")

	for i := 0; i < d.Size; i++ {
		element := d.elements[(d.Front+i)%d.Capacity]
		builder.WriteString(fmt.Sprintf("%v", element))

		if i < d.Size-1 {
			builder.WriteString(", ")
		}
	}

	builder.WriteString("]")
	return builder.String()
}
