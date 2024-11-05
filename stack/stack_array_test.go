package stack

import "testing"

// 测试数据 1 2 3 4 5

func TestStackArray(t *testing.T) {
	s := NewStackArray()
	if !s.IsEmpty() {
		t.Errorf("stack is empty")
	}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	if s.Len() != 5 {
		t.Errorf("stack length is not 5")
	}
	if v, err := s.Peek(); err != nil || v != 5 {
		t.Errorf("stack peek is not 5")
	}
	if v, err := s.Pop(); err != nil || v != 5 {
		t.Errorf("stack pop is not 5")
	}
	if v, err := s.Pop(); err != nil || v != 4 {
		t.Errorf("stack pop is not 4")
	}
	if v, err := s.Pop(); err != nil || v != 3 {
		t.Errorf("stack pop is not 3")
	}
}
