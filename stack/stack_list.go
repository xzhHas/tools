package stack

import (
	"container/list"
	"errors"
)

// 定义一个泛型栈，使用链表作为底层数据结构
type StackList[T any] struct {
	list *list.List
}

// NewStackList 创建一个空的链表栈
func NewStackList[T any]() *StackList[T] {
	return &StackList[T]{list: list.New()}
}

// 推入元素
func (s *StackList[T]) Push(value T) {
	s.list.PushBack(value)
}

// 弹出元素
func (s *StackList[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T // 返回该类型的零值
		return zero, errors.New("stack is empty")
	}
	element := s.list.Back()
	s.list.Remove(element)
	return element.Value.(T), nil // 类型断言为 T
}

// 查看栈顶元素
func (s *StackList[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T // 返回该类型的零值
		return zero, errors.New("stack is empty")
	}
	return s.list.Back().Value.(T), nil // 类型断言为 T
}

// 判断栈是否为空
func (s *StackList[T]) IsEmpty() bool {
	return s.list.Len() == 0
}

// 获取栈的长度
func (s *StackList[T]) Len() int {
	return s.list.Len()
}
