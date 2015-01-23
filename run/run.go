// Package run holds context for runtime and initial run methods.
package run

import (
	"errors"
	"github.com/jpittis/jduck/stack"
)

type Context struct {
	scope *stack.Stack
}

func NewContext() *Context {
	s := stack.New()
	global := make(map[string]interface{})
	s.Push(global)
	return &Context{scope: s}
}

// Push adds a new scope to the context.
func (c *Context) Push() {
	layer := make(map[string]interface{})
	c.scope.Push(layer)
}

// Pop removes last scope from context.
// Returns Error if size of stack is 1.
func (c *Context) Pop() error {
	if c.scope.Size() == 1 {
		return errors.New("cannot pop last global layer of stack")
	}
	_, err := c.scope.Pop()
	return err
}

// Let sets variable in current context.
func (c *Context) Let(key string, value interface{}) error {
	scope, err := c.scope.Peek()
	if err != nil {
		return err
	}
	_, prs := scope[key]
	if prs {
		return errors.New("variable already set")
	}
	scope[key] = value
	return nil
}

// Set sets already set variable in first context found.
func (c *Context) Set(key string, value interface{}) error {
	scope := c.scope.Entity()
	for scope != nil {
		_, prs := scope.Value[key]
		if prs {
			scope.Value[key] = value
			return nil
		}
		scope = scope.Next
	}
	return errors.New("variable not initialized")
}

// Get returns variable highest on the stack.
func (c *Context) Get(key string) (interface{}, error) {
	scope := c.scope.Entity()
	for scope != nil {
		value, prs := scope.Value[key]
		if prs {
			return value, nil
		}
		scope = scope.Next
	}
	return nil, errors.New("variable not set")

}

func Run(ast []Stmt) {
	c := NewContext()
	Run_all(ast, c)
}

func Run_all(s []Stmt, c *Context) {
	for _, s := range s {
		run_stmt(s, c)
	}
}

func run_stmt(s Stmt, c *Context) {
	s.Exec(c)
}
