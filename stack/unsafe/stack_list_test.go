package unsafe

import (
	"testing"
)

func TestStackList(t *testing.T) {
	// 创建栈
	s := NewStackList()

	// 测试 IsEmpty 初始状态
	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty initially, but it was not.")
	}

	// 测试 Len 初始状态
	if got := s.Len(); got != 0 {
		t.Errorf("Expected stack length to be 0 initially, but got %d.", got)
	}

	// 测试 Push
	s.Push(10)
	s.Push("Hello")
	s.Push(3.14)

	// 测试 Len
	if got := s.Len(); got != 3 {
		t.Errorf("Expected stack length to be 3, but got %d.", got)
	}

	// 测试 Peek
	peekVal, err := s.Peek()
	if err != nil {
		t.Errorf("Unexpected error when peeking: %v", err)
	}
	if peekVal != 3.14 {
		t.Errorf("Expected top of the stack to be 3.14, but got %v.", peekVal)
	}

	// 测试 Pop
	val, err := s.Pop()
	if err != nil {
		t.Errorf("Unexpected error when popping: %v", err)
	}
	if val != 3.14 {
		t.Errorf("Expected popped value to be 3.14, but got %v.", val)
	}

	// 再次检查栈长度
	if got := s.Len(); got != 2 {
		t.Errorf("Expected stack length to be 2 after pop, but got %d.", got)
	}

	// 测试再次 Pop
	val, err = s.Pop()
	if err != nil {
		t.Errorf("Unexpected error when popping: %v", err)
	}
	if val != "Hello" {
		t.Errorf("Expected popped value to be 'Hello', but got %v.", val)
	}

	// 测试栈长度
	if got := s.Len(); got != 1 {
		t.Errorf("Expected stack length to be 1, but got %d.", got)
	}

	// 测试 Pop 直到栈空
	val, err = s.Pop()
	if err != nil {
		t.Errorf("Unexpected error when popping: %v", err)
	}
	if val != 10 {
		t.Errorf("Expected popped value to be 10, but got %v.", val)
	}

	// 测试栈空状态
	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty, but it was not.")
	}

	// 测试 Pop 在空栈时
	_, err = s.Pop()
	if err == nil {
		t.Errorf("Expected error when popping from empty stack, but got nil.")
	}

	// 测试 Peek 在空栈时
	_, err = s.Peek()
	if err == nil {
		t.Errorf("Expected error when peeking into empty stack, but got nil.")
	}
}
