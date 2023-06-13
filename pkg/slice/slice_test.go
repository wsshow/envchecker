package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	sl := New[int]()
	sl.Push(3, 2, 4, 5, 1, 0)
	sl = sl.Filter(func(a int) bool { return a > 3 })
	assert.Equal(t, New[int](4, 5), sl)
}

func TestSort(t *testing.T) {
	sl := New[int]()
	sl.Push(3, 2, 4, 5, 1, 0)
	sl.Sort(func(a, b int) bool { return a < b })
	assert.Equal(t, New[int](0, 1, 2, 3, 4, 5), sl)
}

func TestFind(t *testing.T) {
	sl := New[int]()
	sl.Push(3, 2, 4, 5, 1, 0)
	v, _ := sl.Find(func(a int) bool { return a > 3 })
	assert.Equal(t, 4, v)
}

func TestPop(t *testing.T) {
	sl := New[int]()
	sl.Push(3, 2, 4, 5, 1, 0)
	assert.Equal(t, 0, sl.Pop())
	assert.Equal(t, 1, sl.Pop())
	assert.Equal(t, 5, sl.Pop())
}

func TestMap(t *testing.T) {
	sl := New[int]()
	sl.Push(3, 2, 4, 5, 1, 0)
	sl.Map(func(value int) int { return value * 2 })
	assert.Equal(t, New[int](6, 4, 8, 10, 2, 0), sl)
}

func TestReduce(t *testing.T) {
	sl := New[int]()
	sl.Push(3, 2, 4, 5, 1, 0)
	v := sl.Reduce(func(previousValue, currentValue int) int { return previousValue + currentValue }, 0)
	assert.Equal(t, 15, v)
}
