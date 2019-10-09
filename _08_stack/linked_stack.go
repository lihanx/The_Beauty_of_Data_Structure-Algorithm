package _08_stack

import (
	"Lesson/_06_linked_list"
)

// 链式栈
type LinkedStack struct {
	Stack *_06_linked_list.DoubleLinkedList
	Length int
	MaxLength int
}

// 创建链式栈
func NewLinkedStack(maxLength int) *LinkedStack {
	stack := _06_linked_list.NewDoubleLinkedList()
	linkedStack := LinkedStack{
		Stack:     stack,
		Length:    0,
		MaxLength: maxLength,
	}
	return &linkedStack
}

// 出栈
func (this *LinkedStack) pop() *_06_linked_list.DNode {
	if this.Length == 0 {
		return nil
	}
	tail := this.Stack.Tail
	this.Stack.DeleteByNode(tail)
	this.Length--
	return tail
}

// 压栈
func (this *LinkedStack) push(value interface{}) int {
	if this.Length == this.MaxLength {
		return -1
	}
	this.Stack.Append(value)
	this.Length++
	return this.Length
}