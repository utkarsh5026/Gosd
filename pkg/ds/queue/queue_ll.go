package queue

import (
	"errors"
	"fmt"
	"strings"
)

// Node represents a node in the queue.
type Node[T any] struct {
	Data T
	Next *Node[T]
}

// QueueLL represents a queue data structure implemented using a linked list.
type QueueLL[T any] struct {
	Front *Node[T] // Front points to the front node in the queue.
	Back  *Node[T] // Back points to the back node in the queue.
	Size  int      // Size holds the number of elements in the queue.
}

// Enquee adds a new element to the back of the queue.
func (q *QueueLL[T]) Enquee(data T) {
	node := &Node[T]{Data: data}

	if q.Back != nil {
		q.Back.Next = node
	}
	q.Back = node
	if q.Front == nil {
		q.Front = node
	}

	q.Size++
}

// Dequeue removes an element from the front of the queue and returns it.
// If the queue is empty, it returns an error.
func (q *QueueLL[T]) Dequeue() (T, error) {

	var zeroValue T
	if q.Front == nil {
		return zeroValue, errors.New("Queue is empty")
	}

	first := q.Front
	q.Front = q.Front.Next

	if q.Front == nil {
		q.Back = nil
	}
	q.Size--
	return first.Data, nil
}

// peek returns the data at the back of the queue without removing it.
// If the queue is empty, it returns an error.
func (q *QueueLL[T]) peek() (T, error) {
	var zeroValue T

	if q.IsEmpty() {
		return zeroValue, errors.New("Queue is Empty")
	}

	zeroValue = q.Back.Data
	return zeroValue, nil
}

// IsEmpty checks if the queue is empty.
func (q *QueueLL[T]) IsEmpty() bool {
	return q.Front == nil
}

// String returns a string representation of the queue.
func (q *QueueLL[T]) String() string {
	var builder strings.Builder

	builder.WriteString("[")

	current := q.Front

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
