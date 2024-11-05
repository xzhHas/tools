// stack_array.go
package stack

import (
	"errors"
)

type StackArray struct {
	elements []any
}

func NewStackArray() *StackArray {
	return &StackArray{}
}

// Push adds an element to the stack
func (s *StackArray) Push(value any) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element from the stack
func (s *StackArray) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	top := s.elements[s.Len()-1]
	s.elements = s.elements[:s.Len()-1]
	return top, nil
}

// Peek returns the top element of the stack without removing it
func (s *StackArray) Peek() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return s.elements[s.Len()-1], nil
}

// IsEmpty checks if the stack is empty
func (s *StackArray) IsEmpty() bool {
	return len(s.elements) == 0
}

// Len returns the number of elements in the stack
func (s *StackArray) Len() int {
	return len(s.elements)
}
