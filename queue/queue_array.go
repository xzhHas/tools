package queue

import (
	"errors"
)

type QueueArray[T any] struct {
	elements []T
}

func NewQueueArray[T any]() *QueueArray[T] {
	return &QueueArray[T]{}
}

// Enqueue 方法：向队列尾部添加元素
func (q *QueueArray[T]) Enqueue(value T) {
	q.elements = append(q.elements, value)
}

// Dequeue 方法：从队列头部删除元素并返回
func (q *QueueArray[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}
	front := q.elements[0]
	q.elements = q.elements[1:]
	return front, nil
}

// Front 方法：返回队列头部的元素，但不删除
func (q *QueueArray[T]) Front() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}
	return q.elements[0], nil
}

// Back 方法：返回队列尾部的元素，但不删除
func (q *QueueArray[T]) Back() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}
	return q.elements[q.Len()-1], nil
}

// IsEmpty 方法：检查队列是否为空
func (q *QueueArray[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

// Len 方法：返回队列的大小
func (q *QueueArray[T]) Len() int {
	return len(q.elements)
}
