package main

import (
	"fmt"
)

type Stack interface {
	Pop() (interface{}, bool)
	Push(e interface{}) bool
}

type ArrayStack struct {
	arr    []interface{}
	length int
}

func NewArrayStack(n int) *ArrayStack {
	return &ArrayStack{
		arr: make([]interface{}, n),
	}
}

func (s *ArrayStack) Pop() (interface{}, bool) {
	if s.length == 0 {
		return nil, false
	}
	s.length--
	return s.arr[s.length], true
}

func (s *ArrayStack) Push(e interface{}) bool {
	if s.length == len(s.arr) {
		return false
	}
	s.arr[s.length] = e
	s.length++
	return true
}

func main() {
	stack := NewArrayStack(3)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	for true {
		e, ok := stack.Pop()
		if !ok {
			break
		}
		fmt.Println(e)
	}
}
