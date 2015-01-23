package run

import (
	"fmt"
	"log"
)

type Stmt interface {
	Exec(*Context)
}

type LetStmt struct {
	Name   string
	Equals Exp
}

func (s LetStmt) Exec(c *Context) {
	err := c.Let(s.Name, s.Equals.Eval(c))
	if err != nil {
		log.Fatal(err)
	}
}

type VarStmt struct {
	Name   string
	Equals Exp
}

func (s VarStmt) Exec(c *Context) {
	err := c.Set(s.Name, s.Equals.Eval(c))
	if err != nil {
		log.Fatal(err)
	}
}

type PrintStmt struct {
	Print Exp
}

func (s PrintStmt) Exec(c *Context) {
	fmt.Println(s.Print.Eval(c))
}

type IfStmt struct {
	If   Exp
	Then []Stmt
	Else []Stmt
}

func (s IfStmt) Exec(c *Context) {
	b := s.If.Eval(c)
	if b.(bool) {
		c.Push()
		Run_all(s.Then, c)
		c.Pop()
	} else if s.Else != nil {
		c.Push()
		Run_all(s.Else, c)
		c.Pop()
	}
}

/*type ForStmt struct {
	Init   Stmt
	Bool   Exp
	Change Stmt
	Body   []Stmt
}

func (s ForStmt) Exec(data map[string]interface{}) {
	s.Init.Exec(data)
	b := s.Bool.Eval(data)
	for b.(bool) {
		Run_all(s.Body, data)
		s.Change.Exec(data)
		b = s.Bool.Eval(data)
	}
}*/

type WhileStmt struct {
	Bool Exp
	Body []Stmt
}

func (s WhileStmt) Exec(c *Context) {
	b := s.Bool.Eval(c)
	for b.(bool) {
		c.Push()
		Run_all(s.Body, c)
		b = s.Bool.Eval(c)
		c.Pop()
	}
}

/*type FuncStmt struct {
	Body   []Stmt
}

func (s WhileStmt) Exec(data map[string]interface{}) {

}
*/
