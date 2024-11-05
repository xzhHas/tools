// stack_list.go
package stack

import (
	"container/list"
	"errors"
)

type StackList struct {
	list *list.List // 使用 container/list 来实现双向链表
}

func NewStackList() *StackList {
	return &StackList{list: list.New()}
}

func (s *StackList) Push(value any) {
	s.list.PushBack(value)
}

func (s *StackList) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	element := s.list.Back()
	s.list.Remove(element)
	return element.Value, nil
}

func (s *StackList) Peek() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return s.list.Back().Value, nil
}

func (s *StackList) IsEmpty() bool {
	return s.list.Len() == 0
}

func (s *StackList) Len() int {
	return s.list.Len()
}
