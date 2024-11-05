// stack_array.go
package unsafe

import (
	"errors"
)

type StackArray struct {
	elements []any
}

func NewStackArray() *StackArray {
	return &StackArray{}
}

func (s *StackArray) Push(value any) {
	s.elements = append(s.elements, value)
}

func (s *StackArray) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	top := s.elements[s.Len()-1]
	s.elements = s.elements[:s.Len()-1]
	return top, nil
}

func (s *StackArray) Peek() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return s.elements[s.Len()-1], nil
}

func (s *StackArray) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *StackArray) Len() int {
	return len(s.elements)
}
