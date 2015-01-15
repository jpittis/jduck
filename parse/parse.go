package parse

import (
	"github.com/jpittis/jduck/lex"
)

func Parse(st *lex.Lexer) *exp {
	return parse_exp(st)
}

func parse_exp(st *lex.Lexer) *exp {
	if st.Peek().T == lex.EOL {
		return nil
	}
	return parse_pm(st)
}

func parse_pm(st *lex.Lexer) *exp {
	left := parse_lit(st)
	switch st.Peek().T {
	case lex.Add:
		st.Eat()
		return &exp{&BinExp{Op: Add, Left: left, Right: parse_lit(st)}}
	case lex.Sub:
		st.Eat()
		return &BinExp{Op: Sub, Left: left, Right: parse_lit(st)}
	default:
		return parse_exp(st)
	}
}

func parse_lit(st *lex.Lexer) *exp {
	switch st.Peek().T {
	case lex.String:
		st.Eat()
		return &LitExp{value: st.Peek().Value}
	case lex.Integer:
		st.Eat()
		return &LitExp{value: st.Peek().Value}
	default:
		return parse_exp(st)
	}
}
