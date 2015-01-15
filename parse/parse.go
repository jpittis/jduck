package parse

import (
	"fmt"
	"github.com/jpittis/jduck/lex"
)

func Parse(st *lex.Lexer) exp {
	return parse_exp(st)
}

func parse_exp(st *lex.Lexer) exp {
	return parse_pm(st)
}

func parse_pm(st *lex.Lexer) exp {
	left := parse_lit(st)
	fmt.Printf("%+v, (pm)\n", st.Peek())
	switch st.Peek().T {
	case lex.Add:
		st.Eat()
		return exp(BinExp{Op: Add, Left: left, Right: parse_pm(st)})
	case lex.Sub:
		st.Eat()
		return exp(BinExp{Op: Sub, Left: left, Right: parse_pm(st)})
	default:
		return left
	}
}

func parse_lit(st *lex.Lexer) exp {
	fmt.Printf("%+v, (lit)\n", st.Peek())
	switch st.Peek().T {
	case lex.String:
		st.Eat()
		return exp(LitExp{value: st.Peek().Value})
	case lex.Integer:
		st.Eat()
		fmt.Println("RETURNING!")
		return exp(LitExp{value: st.Peek().Value})
	default:
		fmt.Println("BAM!")
		return exp(LitExp{value: nil})
	}
}
