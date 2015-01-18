package parse

import (
	"fmt"
)

type Stmt interface {
	Exec(map[string]interface{})
}

type VarStmt struct {
	Name   string
	Equals exp
}

func (s VarStmt) Exec(data map[string]interface{}) {
	data[s.Name] = s.Equals.Eval(data)
}

type PrintStmt struct {
	Print exp
}

func (s PrintStmt) Exec(data map[string]interface{}) {
	fmt.Println(s.Print.Eval(data))
}

type IfStmt struct {
	If   exp
	Then []Stmt
	Else IfStmt
}

func (s IfStmt) Exec(data map[string]interface{}) {
	b := If.Eval(data)
	if b {
		run.Run_all(s.Then, data)
	} else {
		s.IfStmt.Exec(data)
	}
}

type ForStmt struct {
	Init   Stmt
	Bool   exp
	Change Stmt
	Body   []Stmt
}

func (s ForStmt) Exec(data map[string]interface{}) {
	s.Init.Exec(data)
	b := s.Bool.Eval(data)
	for b {
		run.Run_all(s.body, data)
		s.Change.Exec(data)
		b = s.Bool.Eval(data)
	}
}

type WhileStmt struct {
	Bool exp
	Body []Stmt
}

func (s WhileStmt) Exec(data map[string]interface{}) {
	b := s.Bool.Eval(data)
	for b {
		run.Run_all(s.body, data)
		b = s.Bool.Eval(data)
	}
}

/*type FuncStmt struct {
	Body   []Stmt
}

func (s WhileStmt) Exec(data map[string]interface{}) {

}
*/
