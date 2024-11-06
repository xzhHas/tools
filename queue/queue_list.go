package queue

import (
	"container/list"
	"errors"
)

type QueueList[T any] struct {
	elements *list.List // 使用 container/list 来实现双向链表
}

// NewQueueList 创建一个新的链表队列
func NewQueueList[T any]() *QueueList[T] {
	return &QueueList[T]{elements: list.New()}
}

// Enqueue 向队列尾部添加元素
func (q *QueueList[T]) Enqueue(value T) {
	q.elements.PushBack(value)
}

// Dequeue 从队列头部删除元素并返回
func (q *QueueList[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}
	front := q.elements.Front()
	q.elements.Remove(front)
	return front.Value.(T), nil
}

// Front 返回队列头部的元素，但不删除
func (q *QueueList[T]) Front() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}
	return q.elements.Front().Value.(T), nil
}

// Back 返回队列尾部的元素，但不删除
func (q *QueueList[T]) Back() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}
	return q.elements.Back().Value.(T), nil
}

// IsEmpty 检查队列是否为空
func (q *QueueList[T]) IsEmpty() bool {
	return q.elements.Len() == 0
}

// Len 返回队列的大小
func (q *QueueList[T]) Len() int {
	return q.elements.Len()
}
