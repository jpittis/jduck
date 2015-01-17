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
	Else []Stmt
}

func (s IfStmt) Exec(data map[string]interface{}) {
}

type ForStmt struct {
	Body []Stmt
}

func (s ForStmt) Exec(data map[string]interface{}) {
}
