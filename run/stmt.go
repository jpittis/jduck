package run

import (
	"fmt"
)

type Stmt interface {
	Exec(map[string]interface{})
}

type VarStmt struct {
	Name   string
	Equals Exp
}

func (s VarStmt) Exec(data map[string]interface{}) {
	data[s.Name] = s.Equals.Eval(data)
}

type PrintStmt struct {
	Print Exp
}

func (s PrintStmt) Exec(data map[string]interface{}) {
	fmt.Println(s.Print.Eval(data))
}

type IfStmt struct {
	If   Exp
	Then []Stmt
	Else Stmt
}

func (s IfStmt) Exec(data map[string]interface{}) {
	b := s.If.Eval(data)
	if b.(bool) {
		Run_all(s.Then, data)
	} else {
		s.Else.Exec(data)
	}
}

type ForStmt struct {
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
}

type WhileStmt struct {
	Bool Exp
	Body []Stmt
}

func (s WhileStmt) Exec(data map[string]interface{}) {
	b := s.Bool.Eval(data)
	for b.(bool) {
		Run_all(s.Body, data)
		b = s.Bool.Eval(data)
	}
}

/*type FuncStmt struct {
	Body   []Stmt
}

func (s WhileStmt) Exec(data map[string]interface{}) {

}
*/
