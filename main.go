package main

import (
	"fmt"
	"github.com/jpittis/jduck/lex"
	"github.com/jpittis/jduck/parse"
	"log"
	"os"
)

func main() {
	f, err := os.Open("test.jdc")
	if err != nil {
		log.Fatal(err)
	}
	l := lex.New(f)
	fmt.Println(parse.Parse(st))
}
