package main

import (
	
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v queue_circular_test.go

type CircularQueue struct {
	values []int  	// фиксированный буфер; 0 — пустая ячейка
	start   int 	// индекс начала (элемент для Pop/Front)
	next    int   	// индекс следующей позиции для Push
	count   int   	// текущее число элементов
}

// создать очередь с определенным размером буфера
func NewCircularQueue(size int) CircularQueue {
	// if size <= 0 {
	// 	size = 0
	// }
	return CircularQueue{
		values: make([]int, size),
		start:  0,
		next:  	0,
		count: 	0,
	}
}

// проверить заполнена ли очередь
func (q *CircularQueue) Full() bool {
	return q.count == len(q.values)
}

// проверить пустая ли очередь
func (q *CircularQueue) Empty() bool {
	return q.count == 0
}

// добавить значение в конец очереди (false, если очередь заполнена)
func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}
	q.values[q.next] = value
	q.next = (q.next + 1) % len(q.values)
	q.count++
	return true
}

// удалить значение из начала очереди (false, если очередь пустая)
func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}
	q.values[q.start] = 0
	q.start = (q.start + 1) % len(q.values)
	q.count--
	return true
}

// получить значение из начала очереди (-1, если очередь пустая)
func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.start]
}

// получить значение из конца очереди (-1, если очередь пустая)
func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}
	idx := (q.next - 1 + len(q.values)) % len(q.values)
	return q.values[idx]
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}