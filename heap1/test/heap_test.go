package heap

import (
	"container/heap"
	"fmt"
	"github.com/xzhHas/tools/heap1"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	// 创建一个新的大根堆
	hp := heap1.NewMaxHeap()

	// 使用 heap1.Init 初始化堆
	heap.Init(hp)
	heap.Push(hp, 3)
	heap.Push(hp, 1)
	heap.Push(hp, 2)
	fmt.Println(hp.Len())
	fmt.Println(heap.Pop(hp))
	fmt.Println(hp.Len())
}

func TestMinHeap(t *testing.T) {
	// 创建一个新的小根堆
	heap2 := heap1.NewMinHeap()

	// 使用 heap1.Init 初始化堆
	heap.Init(heap2)

	// 向堆中插入元素
	heap.Push(heap2, 1)
	heap.Push(heap2, 2132)
	heap.Push(heap2, 3)

	// 弹出堆顶元素并打印
	fmt.Println("MinHeap - Pop:", heap.Pop(heap2)) // 应该输出 1
	fmt.Println("MinHeap - Pop:", heap.Pop(heap2)) // 应该输出 3
	fmt.Println("MinHeap - Pop:", heap.Pop(heap2)) // 应该输出 2132
}
