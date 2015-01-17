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

}

type ForStmt struct {
	Init   exp
	Bool   exp
	Change exp
	Body   []Stmt
}

func (s ForStmt) Exec(data map[string]interface{}) {

}

type WhileStmt struct {
	Bool exp
	Body []Stmt
}

func (s WhileStmt) Exec(data map[string]interface{}) {

}

/*type FuncStmt struct {
	Body   []Stmt
}

func (s WhileStmt) Exec(data map[string]interface{}) {

}
*/
