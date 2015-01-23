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
	Value map[string]interface{} // interface Value for storing anything
	Next  *Entity                // the Entity below current
}

// Entity returns a pointer to top Entity.
// Effectively turns stack into linked list.
func (s *Stack) Entity() *Entity {
	return s.top
}

// Push adds a Value to the top of the stack.
func (s *Stack) Push(Value map[string]interface{}) {
	newTop := &Entity{Value: Value, Next: nil}

	if s.size != 0 {
		newTop.Next = s.top
	}
	s.top = newTop
	s.size++
	return
}

// Pop removes a Value from the top of the stack. Returns error if empty.
func (s *Stack) Pop() (map[string]interface{}, error) {
	if s.size == 0 {
		return nil, errors.New("stack is empty")
	}
	result := s.top.Value
	s.top = s.top.Next
	s.size--
	return result, nil
}

// Peek returns the top Value of the stack. Returns error if empty.
func (s *Stack) Peek() (map[string]interface{}, error) {
	if s.size == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.top.Value, nil
}

// Size returns the number of Values in the stack.
func (s *Stack) Size() int {
	return s.size
}
