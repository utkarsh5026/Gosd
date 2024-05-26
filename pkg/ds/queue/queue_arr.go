// Package queue provides a generic implementation of a queue using an array (slice) in Go.
package queue

import (
	"errors"
	"fmt"
	"strings"
)

// QueueArr represents a queue using an array (slice).
// It supports basic queue operations such as Enqueue, Dequeue, Peek, IsEmpty, and Length.
type QueueArr[T any] struct {
	elements []T // elements stores the queue items.
	Size     int // Size is the current number of elements in the queue.
	Front    int // Front is the index of the front element.
	Back     int // Back is the index of the back element.
}

// NewQueue creates and returns a new QueueArr.
func NewQueue[T any]() *QueueArr[T] {
	return &QueueArr[T]{
		elements: make([]T, 10),
		Front:    0,
		Back:     0,
		Size:     0,
	}
}

// IsEmpty checks if the queue is empty.
// It returns true if the queue is empty, otherwise false.
func (q *QueueArr[T]) IsEmpty() bool {
	return q.Size == 0
}

// Resize increases the capacity of the queue.
func (q *QueueArr[T]) Resize() {
	newCapacity := len(q.elements) * 2

	if newCapacity == 0 {
		newCapacity = 1
	}

	resized := make([]T, newCapacity)

	for i := 0; i < q.Size; i++ {
		resized[i] = q.elements[(q.Front+i)%len(q.elements)]
	}

	q.elements = resized
	q.Front = 0
	q.Back = q.Size
}

// Enqueue adds a new element to the back of the queue.
func (q *QueueArr[T]) Enqueue(data T) {
	if q.Size == len(q.elements) {
		q.Resize()
	}

	q.elements[q.Back] = data
	q.Back = (q.Back + 1) % len(q.elements)
	q.Size++
}

// Dequeue removes and returns the front element of the queue.
// It returns an error if the queue is empty.
func (q *QueueArr[T]) Dequeue() (T, error) {
	var zeroValue T

	if q.IsEmpty() {
		return zeroValue, errors.New("Queue is empty")
	}

	first := q.elements[q.Front]
	q.Front = (q.Front + 1) % len(q.elements)

	q.Size--
	return first, nil
}

// Length returns the number of elements in the queue.
func (q *QueueArr[T]) Length() int {
	return q.Size
}

// Peek returns the front element of the queue without removing it.
// It returns an error if the queue is empty.
func (q *QueueArr[T]) Peek() (T, error) {
	var zeroValue T

	if q.IsEmpty() {
		return zeroValue, errors.New("Queue is empty")
	}

	first := q.elements[q.Front]
	return first, nil
}

// String returns a string representation of the queue.
func (q *QueueArr[T]) String() string {
	var builder strings.Builder
	builder.WriteString("[")

	for i := 0; i < q.Size; i++ {
		element := q.elements[(q.Front+i)%len(q.elements)]
		builder.WriteString(fmt.Sprintf("%v", element))

		if i < q.Size-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("]")
	return builder.String()
}
