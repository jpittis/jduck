package main

import (
	"github.com/jpittis/jduck/lex"
	"github.com/jpittis/jduck/parse"
	"github.com/jpittis/jduck/run"
	"log"
	"os"
)

func main() {
	f, err := os.Open("test.jdc")
	if err != nil {
		log.Fatal(err)
	}
	st := lex.New(f)
	ast := parse.Parse(st)
	run.Run(ast)
}
