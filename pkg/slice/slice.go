package slice

import "golang.org/x/exp/slices"

type Slice[T any] struct {
	data []T
}

func New[T any](values ...T) *Slice[T] {
	sl := &Slice[T]{
		data: make([]T, 0),
	}
	return sl.Push(values...)
}

func (sl *Slice[T]) Push(values ...T) *Slice[T] {
	sl.data = append(sl.data, values...)
	return sl
}

func (sl *Slice[T]) Pop() T {
	pos := sl.Length() - 1
	result := sl.data[pos]
	sl.data = sl.data[:pos]
	return result
}

func (sl *Slice[T]) Length() int {
	return len(sl.data)
}

func (sl *Slice[T]) Foreach(callbackfn func(value T)) *Slice[T] {
	for _, v := range sl.data {
		callbackfn(v)
	}
	return sl
}

func (sl *Slice[T]) Map(callbackfn func(value T) T) *Slice[T] {
	for i, v := range sl.data {
		sl.data[i] = callbackfn(v)
	}
	return sl
}

func (sl *Slice[T]) Find(predicate func(value T) bool) (result T, existed bool) {
	for _, v := range sl.data {
		if predicate(v) {
			return v, true
		}
	}
	return result, false
}

func (sl *Slice[T]) Filter(predicate func(value T) bool) *Slice[T] {
	result := &Slice[T]{
		data: make([]T, 0, len(sl.data)),
	}
	sl.Foreach(func(v T) {
		if predicate(v) {
			result.Push(v)
		}
	})
	return result
}

func (sl *Slice[T]) Reduce(callbackfn func(previousValue T, currentValue T) T, initialValue T) T {
	acc := initialValue
	for _, v := range sl.data {
		acc = callbackfn(acc, v)
	}
	return acc
}

func (sl *Slice[T]) Sort(compareFn func(a T, b T) bool) *Slice[T] {
	slices.SortFunc(sl.data, compareFn)
	return sl
}
