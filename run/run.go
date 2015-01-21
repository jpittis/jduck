package run

type Context struct {
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
