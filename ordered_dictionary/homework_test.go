package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go


type OrderedMap struct {
	keys []int
	vals []int
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{
		keys: make([]int, 0),
		vals: make([]int, 0),
	}
}

func (m *OrderedMap) Insert(key, value int) {
	idx, ok := m.find(key)
	if ok {
		m.vals[idx] = value
		return
	}
	// вставка на позицию idx с сохранением порядка
	m.keys = append(m.keys, 0)
	m.vals = append(m.vals, 0)
	copy(m.keys[idx+1:], m.keys[idx:])
	copy(m.vals[idx+1:], m.vals[idx:])
	m.keys[idx] = key
	m.vals[idx] = value
}

func (m *OrderedMap) Erase(key int) {
	idx, ok := m.find(key)
	if ok {
		m.keys = append(m.keys[:idx], m.keys[idx+1:]...)
		m.vals = append(m.vals[:idx], m.vals[idx+1:]...)
	}
	return
}

func (m *OrderedMap) Contains(key int) bool {
	_, ok := m.find(key)
	return ok
}

func (m *OrderedMap) Size() int {
	return len(m.keys)
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	for i := range m.keys {
		action(m.keys[i], m.vals[i])
	}
}

// бинарный поиск: возвращает позицию для ключа и флаг найден/нет
func (m *OrderedMap) find(key int) (int, bool) {
	lo, hi := 0, len(m.keys)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if m.keys[mid] < key {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if lo < len(m.keys) && m.keys[lo] == key {
		return lo, true
	}
	return lo, false
} 

func TestCircularQueue(t *testing.T) {
	data := NewOrderedMap()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	data.Erase(14)
	data.Erase(2)

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}