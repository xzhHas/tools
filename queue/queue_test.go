package queue

import (
	"testing"
)

// 测试数组队列（QueueArray）
func TestQueueArray(t *testing.T) {
	// 创建一个 int 类型的队列
	intQueue := NewQueueArray[int]()

	// 向队列中添加元素
	intQueue.Enqueue(10)
	intQueue.Enqueue(20)

	// 测试队列大小
	if intQueue.Len() != 2 {
		t.Errorf("Expected queue size 2, got %d", intQueue.Len())
	}

	// 从队列中取出元素并验证
	value, err := intQueue.Dequeue()
	if err != nil {
		t.Errorf("Dequeue failed: %v", err)
	} else if value != 10 {
		t.Errorf("Expected value 10, got %v", value)
	}

	// 查看队列头部元素并验证
	peekValue, err := intQueue.Front()
	if err != nil {
		t.Errorf("Front failed: %v", err)
	} else if peekValue != 20 {
		t.Errorf("Expected front value 20, got %v", peekValue)
	}

	// 查看队列尾部元素并验证
	peekBackValue, err := intQueue.Back()
	if err != nil {
		t.Errorf("Back failed: %v", err)
	} else if peekBackValue != 20 {
		t.Errorf("Expected back value 20, got %v", peekBackValue)
	}

	// 测试队列是否为空
	if intQueue.IsEmpty() != false {
		t.Errorf("Expected queue to not be empty")
	}

	// 弹出剩余元素
	_, err = intQueue.Dequeue()
	if err != nil {
		t.Errorf("Dequeue failed: %v", err)
	}

	// 测试队列是否为空
	if intQueue.IsEmpty() != true {
		t.Errorf("Expected queue to be empty")
	}

	// 测试弹出空队列的错误
	_, err = intQueue.Dequeue()
	if err == nil {
		t.Errorf("Expected error, got nil")
	} else if err.Error() != "queue is empty" {
		t.Errorf("Expected 'queue is empty' error, got %v", err)
	}
}

// 测试链表队列（QueueList）
func TestQueueList(t *testing.T) {
	// 创建一个 int 类型的队列
	intQueue := NewQueueList[int]()

	// 向队列中添加元素
	intQueue.Enqueue(10)
	intQueue.Enqueue(20)

	// 测试队列大小
	if intQueue.Len() != 2 {
		t.Errorf("Expected queue size 2, got %d", intQueue.Len())
	}

	// 从队列中取出元素并验证
	value, err := intQueue.Dequeue()
	if err != nil {
		t.Errorf("Dequeue failed: %v", err)
	} else if value != 10 {
		t.Errorf("Expected value 10, got %v", value)
	}

	// 查看队列头部元素并验证
	peekValue, err := intQueue.Front()
	if err != nil {
		t.Errorf("Front failed: %v", err)
	} else if peekValue != 20 {
		t.Errorf("Expected front value 20, got %v", peekValue)
	}

	// 查看队列尾部元素并验证
	peekBackValue, err := intQueue.Back()
	if err != nil {
		t.Errorf("Back failed: %v", err)
	} else if peekBackValue != 20 {
		t.Errorf("Expected back value 20, got %v", peekBackValue)
	}

	// 测试队列是否为空
	if intQueue.IsEmpty() != false {
		t.Errorf("Expected queue to not be empty")
	}

	// 弹出剩余元素
	_, err = intQueue.Dequeue()
	if err != nil {
		t.Errorf("Dequeue failed: %v", err)
	}

	// 测试队列是否为空
	if intQueue.IsEmpty() != true {
		t.Errorf("Expected queue to be empty")
	}

	// 测试弹出空队列的错误
	_, err = intQueue.Dequeue()
	if err == nil {
		t.Errorf("Expected error, got nil")
	} else if err.Error() != "queue is empty" {
		t.Errorf("Expected 'queue is empty' error, got %v", err)
	}
}
