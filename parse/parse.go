package parse

import (
	"fmt"
	"github.com/jpittis/jduck/lex"
)

func Parse(st *lex.Lexer) exp {
	return parse_exp(st)
}

func parse_exp(st *lex.Lexer) exp {
	return parse_bb(st)
}

func parse_bb(st *lex.Lexer) exp {
	left := parse_ub(st)
	switch st.Peek().T {
	case lex.GreatThan:
		st.Eat()
		return exp(BinExp{Op: GreatThan, Left: left, Right: parse_bb(st)})
	case lex.LessThan:
		st.Eat()
		return exp(BinExp{Op: LessThan, Left: left, Right: parse_bb(st)})
	case lex.GreatThanEq:
		st.Eat()
		return exp(BinExp{Op: GreatThanEq, Left: left, Right: parse_bb(st)})
	case lex.LessThanEq:
		st.Eat()
		return exp(BinExp{Op: LessThanEq, Left: left, Right: parse_bb(st)})
	case lex.EqEq:
		st.Eat()
		return exp(BinExp{Op: EqEq, Left: left, Right: parse_bb(st)})
	case lex.NotEq:
		st.Eat()
		return exp(BinExp{Op: NotEq, Left: left, Right: parse_bb(st)})
	case lex.And:
		st.Eat()
		return exp(BinExp{Op: And, Left: left, Right: parse_bb(st)})
	case lex.Or:
		st.Eat()
		return exp(BinExp{Op: Or, Left: left, Right: parse_bb(st)})
	default:
		return left
	}
}

func parse_ub(st *lex.Lexer) exp {
	left := parse_pm(st)
	switch st.Peek().T {
	case lex.Not:
		st.Eat()
		return exp(UnaExp{Op: Not, Right: parse_ub(st)})
	default:
		return left
	}
}

func parse_pm(st *lex.Lexer) exp {
	left := parse_md(st)
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

func parse_md(st *lex.Lexer) exp {
	left := parse_um(st)
	switch st.Peek().T {
	case lex.Mul:
		st.Eat()
		return exp(BinExp{Op: Mul, Left: left, Right: parse_md(st)})
	case lex.Div:
		st.Eat()
		return exp(BinExp{Op: Div, Left: left, Right: parse_md(st)})
	case lex.Mod:
		st.Eat()
		return exp(BinExp{Op: Mod, Left: left, Right: parse_md(st)})
	default:
		return left
	}
}

func parse_um(st *lex.Lexer) exp {
	left := parse_lit(st)
	switch st.Peek().T {
	case lex.Sub:
		st.Eat()
		return exp(UnaExp{Op: Neg, Right: parse_um(st)})
	default:
		return left
	}
}

func parse_lit(st *lex.Lexer) exp {
	switch st.Peek().T {
	case lex.String:
		tok := st.Eat()
		return exp(LitExp{value: tok.Value})
	case lex.Integer:
		tok := st.Eat()
		return exp(LitExp{value: tok.Value})
	case lex.Bool:
		tok := st.Eat()
		return exp(LitExp{value: tok.Value})
	default:
		fmt.Println("Unknown Literal Reached")
		return exp(LitExp{value: nil})
	}
}
