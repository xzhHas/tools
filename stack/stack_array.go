package stack

import (
	"errors"
)

// 定义一个泛型栈
type StackArray[T any] struct {
	elements []T
}

// NewStackArray 创建一个空的泛型数组栈
func NewStackArray[T any]() *StackArray[T] {
	return &StackArray[T]{}
}

// 推入元素
func (s *StackArray[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// 弹出元素
func (s *StackArray[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T // 返回该类型的零值
		return zero, errors.New("stack is empty")
	}
	top := s.elements[s.Len()-1]
	s.elements = s.elements[:s.Len()-1]
	return top, nil
}

// 查看栈顶元素
func (s *StackArray[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T // 返回该类型的零值
		return zero, errors.New("stack is empty")
	}
	return s.elements[s.Len()-1], nil
}

// 判断栈是否为空
func (s *StackArray[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// 获取栈的长度
func (s *StackArray[T]) Len() int {
	return len(s.elements)
}
