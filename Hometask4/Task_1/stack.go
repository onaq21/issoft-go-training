package main

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	head *Node[T]
}

type Node[T any] struct {
	data T
	next *Node[T]
}

func (stack *Stack[T]) Push(data T) {
	stack.head = &Node[T]{data, stack.head}
}

func (stack *Stack[T]) Pop() (T, error) {
	var data T
	if stack.IsEmpty() {
		return data, errors.New("Stack is empty")
	}
	data = stack.head.data
	stack.head = stack.head.next
	return data, nil
}

func(stack *Stack[T]) IsEmpty() bool {
	return stack.head == nil
}

func main() {
	stack := &Stack[string]{}
	
	fmt.Println(stack.IsEmpty())

	stack.Push("first")
	stack.Push("second")
	fmt.Println(stack.Pop())

	fmt.Println(stack.IsEmpty())
}