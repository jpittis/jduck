package run

import (
	"github.com/jpittis/jduck/parse"
)

type state struct {
}

func Run(ast []parse.Stmt) {
	data := make(map[string]interface{})
	for _, s := range ast {
		run_stmt(s, data)
	}
}

func run_stmt(s parse.Stmt, data map[string]interface{}) {
	s.Exec(data)
}
