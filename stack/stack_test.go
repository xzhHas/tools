package stack

import (
	"testing"
)

func TestStackArray(t *testing.T) {
	// int 类型的栈
	intStack := NewStackArray[int]()

	// 向栈中推入元素
	intStack.Push(10)
	intStack.Push(20)

	// 弹出元素并验证
	value, err := intStack.Pop()
	if err != nil {
		t.Errorf("Pop failed: %v", err)
	} else if value != 20 {
		t.Errorf("Expected value 20, got %v", value)
	}

	// 查看栈顶元素并验证
	peekValue, err := intStack.Peek()
	if err != nil {
		t.Errorf("Peek failed: %v", err)
	} else if peekValue != 10 {
		t.Errorf("Expected peek value 10, got %v", peekValue)
	}

	// 创建一个 string 类型的栈
	stringStack := NewStackArray[string]()

	// 向栈中推入元素
	stringStack.Push("Hello")
	stringStack.Push("World")

	// 弹出元素并验证
	stringValue, err := stringStack.Pop()
	if err != nil {
		t.Errorf("Pop failed: %v", err)
	} else if stringValue != "World" {
		t.Errorf("Expected string value 'World', got %v", stringValue)
	}

	// 查看栈顶元素并验证
	peekStringValue, err := stringStack.Peek()
	if err != nil {
		t.Errorf("Peek failed: %v", err)
	} else if peekStringValue != "Hello" {
		t.Errorf("Expected peek string 'Hello', got %v", peekStringValue)
	}
}

func TestStackList(t *testing.T) {
	// 创建一个 int 类型的栈
	intStack := NewStackList[int]()

	// 向栈中推入元素
	intStack.Push(10)
	intStack.Push(20)

	// 弹出元素并验证
	value, err := intStack.Pop()
	if err != nil {
		t.Errorf("Pop failed: %v", err)
	} else if value != 20 {
		t.Errorf("Expected value 20, got %v", value)
	}

	// 查看栈顶元素并验证
	peekValue, err := intStack.Peek()
	if err != nil {
		t.Errorf("Peek failed: %v", err)
	} else if peekValue != 10 {
		t.Errorf("Expected peek value 10, got %v", peekValue)
	}

	// 创建一个 string 类型的栈
	stringStack := NewStackList[string]()

	// 向栈中推入元素
	stringStack.Push("Hello")
	stringStack.Push("World")

	// 弹出元素并验证
	stringValue, err := stringStack.Pop()
	if err != nil {
		t.Errorf("Pop failed: %v", err)
	} else if stringValue != "World" {
		t.Errorf("Expected string value 'World', got %v", stringValue)
	}

	// 查看栈顶元素并验证
	peekStringValue, err := stringStack.Peek()
	if err != nil {
		t.Errorf("Peek failed: %v", err)
	} else if peekStringValue != "Hello" {
		t.Errorf("Expected peek string 'Hello', got %v", peekStringValue)
	}
}
