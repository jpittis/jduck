package main

import (
	"fmt"
	"github.com/jpittis/jduck/lex"
	"log"
	"os"
)

func main() {
	f, err := os.Open("test.jdc")
	if err != nil {
		log.Fatal(err)
	}
	l := lex.New(f)
	t, err := l.Eat()
	for t.T != lex.EOF {
		t, err = l.Eat()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(t)
	}
}
