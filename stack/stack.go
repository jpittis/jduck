// Package stack contains methods for creating and manipulating a stack.
package stack

import (
	"errors"
)

// Top level stack header.
type Stack struct {
	top  *Entity // the top Entity on the stack
	size int     // total number of entities in the stack
}

// New creates a new stack of size 0.
func New() *Stack {
	return &Stack{top: nil, size: 0}
}

// Value in the stack. Stored as linked list.
type Entity struct {
	value map[string]interface{} // interface value for storing anything
	next  *Entity                // the Entity below current
}

// Entity returns a pointer to top Entity.
// Effectively turns stack into linked list.
func (s *Stack) Entity() *Entity {
	return s.top
}

// Push adds a value to the top of the stack.
func (s *Stack) Push(value map[string]interface{}) {
	newTop := &Entity{value: value, next: nil}

	if s.size != 0 {
		newTop.next = s.top
	}
	s.top = newTop
	s.size++
	return
}

// Pop removes a value from the top of the stack. Returns error if empty.
func (s *Stack) Pop() (map[string]interface{}, error) {
	if s.size == 0 {
		return nil, errors.New("stack is empty")
	}
	result := s.top.value
	s.top = s.top.next
	s.size--
	return result, nil
}

// Peek returns the top value of the stack. Returns error if empty.
func (s *Stack) Peek() (map[string]interface{}, error) {
	if s.size == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.top.value, nil
}

// Size returns the number of values in the stack.
func (s *Stack) Size() int {
	return s.size
}
