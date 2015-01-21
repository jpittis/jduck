// Package run holds context for runtime and initial run methods.
package run

import (
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

}

// Pop removes last scope from context.
// Returns Error if size of stack is 1.
func (c *Context) Pop() error {

}

// Let sets variable in current context.
func (c *Context) Let() error {

}

// Set sets already set variable in first context found.
func (c *Context) Set() error {

}

// Get returns variable highest on the stack.
func (c *Context) Get() interface{} {

}

func Run(ast []Stmt) {
	data := make(map[string]interface{})
	Run_all(ast, data)
}

func Run_all(s []Stmt, data map[string]interface{}) {
	for _, s := range s {
		run_stmt(s, data)
	}
}

func run_stmt(s Stmt, data map[string]interface{}) {
	s.Exec(data)
}
