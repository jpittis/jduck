package parse

type stmt interface {
	Exec()
}

type VarStmt struct {
	name   string
	equals exp
}

func (s VarStmt) Exec() {
	return
}
