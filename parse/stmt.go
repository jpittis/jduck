package parse

type stmt interface {
	Exec()
}

type VarStmt struct {
	Name   string
	Equals exp
}

func (s VarStmt) Exec() {
	return
}

type PrintStmt struct {
	Print exp
}

func (s PrintStmt) Exec() {
	return
}
