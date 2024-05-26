package queue

import (
	"errors"
	"fmt"
	"strings"
)

type QueueArr[T any] struct {
	elements []T
	Size     int
	Front    int
	Back     int
}

func NewQueue[T any]() *QueueArr[T] {
	return &QueueArr[T]{
		elements: make([]T, 10),
		Front:    0,
		Back:     0,
		Size:     0,
	}
}

func (q *QueueArr[T]) IsEmpty() bool {
	return q.Size == 0
}

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

func (q *QueueArr[T]) Enqueue(data T) {
	if q.Size == len(q.elements) {
		q.Resize()
	}

	q.elements[q.Back] = data
	q.Back = (q.Back + 1) % len(q.elements)
	q.Size++
}

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

func (q *QueueArr[T]) Length() int {
	return q.Size
}

func (q *QueueArr[T]) Peek() (T, error) {
	var zeroValue T

	if q.IsEmpty() {
		return zeroValue, nil
	}

	first := q.elements[q.Front]
	return first, nil
}

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
