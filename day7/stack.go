package main

import (
	"fmt"
	"strings"
)

type Stack[T fmt.Stringer] []T

func (s Stack[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s Stack[T]) Push(v T) Stack[T] {
	return append(s, v)
}

func (s Stack[T]) Pop() (Stack[T], T) {
	if s.IsEmpty() {
		panic("Stack is empty!")
	}

	return s[:len(s)-1], s[len(s)-1]
}

func (s *Stack[T]) Peek() *T {
	if s.IsEmpty() {
		return nil
	}
	return &(*s)[len(*s)-1] // Return reference to last element in stack
}

func (s Stack[T]) String() string {
	result := []string{}

	for _, element := range s {
		result = append(result, element.String())
	}

	return strings.Join(result, "")
}
