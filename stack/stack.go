// Package stack contains methods for creating and manipulating a stack.
package stack

import (
	"errors"
)

// Top level stack header.
type Stack struct {
	top  *entity // the top entity on the stack
	size int     // total number of entities in the stack
}

// New creates a new stack of size 0.
func New() *Stack {
	return &stack{top: nil, size: 0}
}

// Value in the stack. Stored as linked list.
type entity struct {
	value interface{} // interface value for storing anything
	next  *entity     // the entity below current
}

// Push adds a value to the top of the stack.
func (s *Stack) Push(value interface{}) {
	newTop := &entity{value: value, next: nil}

	if s.size != 0 {
		newTop.next = s.top
	}
	s.top = newTop
	s.size++
	return
}

// Pop removes a value from the top of the stack. Returns error if empty.
func (s *Stack) Pop() (interface{}, error) {
	if s.size == 0 {
		return nil, errors.New("stack is empty")
	}
	result := s.top.value
	s.top = s.top.next
	s.size--
	return result, nil
}

// Peek returns the top value of the stack. Returns error if empty.
func (s *Stack) Peek() (interface{}, error) {
	if s.size == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.top.value, nil
}

// Size returns the number of values in the stack.
func (s *Stack) Size() int {
	return s.size
}
