package util

import "fmt"

type Stack[T comparable] struct {
	values []T
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

func (v *Stack[T]) Push(t T) {
	v.values = append(v.values, t)
}

func (v *Stack[T]) Pop() T {
	t := v.values[len(v.values)-1]
	v.values = v.values[:len(v.values)-1]
	return t
}

func (v *Stack[T]) Value() T {
	if len(v.values) == 0 {
		var t T
		return t
	}
	return v.values[len(v.values)-1]
}

func (v *Stack[T]) Size() int {
	return len(v.values)
}

func (v *Stack[T]) String() string {
	return fmt.Sprintf("%v", v.values)
}
