package parse

import (
	"fmt"
	"github.com/jpittis/jduck/lex"
	"github.com/jpittis/jduck/run"
	"log"
)

func Parse(st *lex.Lexer) []run.Stmt {
	s := make([]run.Stmt, 0)
	for st.Peek().T != lex.EOF {
		switch st.Peek().T {
		case lex.Fun:
			st.Eat()
			if st.Peek().T != lex.Ident {
				log.Fatal("no ident after func keyword")
			}
			tok := st.Eat()
			s = append(s, parse_params(st, tok.Value.(string)))
		case lex.Let:
			st.Eat()
			tok := st.Eat()
			if st.Peek().T != lex.Eq {
				fmt.Println("equals not found after name")
			}
			st.Eat()
			s = append(s, parse_let(st, tok.Value.(string)))
		case lex.Ident:
			tok := st.Eat()
			if st.Peek().T != lex.Eq {
				fmt.Println("equals not found after name")
			}
			st.Eat()
			s = append(s, parse_ident(st, tok.Value.(string)))
		case lex.Print:
			st.Eat()
			s = append(s, parse_print(st))
		case lex.If:
			st.Eat()
			s = append(s, parse_if(st))
			/*	case lex.For:
				st.Eat()
				s = append(s, parse_for(st))*/
		case lex.While:
			st.Eat()
			s = append(s, parse_while(st))
		case lex.End:
			st.Eat()
			return s
		case lex.Else:
			return s
		default:
			log.Fatal("statement type not found")
		}
	}
	return s
}

func parse_if(st *lex.Lexer) run.Stmt {
	b := parse_exp(st)
	body := parse_body(st)
	var rest []run.Stmt
	if st.Peek().T == lex.Else {
		st.Eat()
		rest = parse_body(st)
	}
	return run.IfStmt{If: b, Then: body, Else: rest}
}

/*func parse_for(st *lex.Lexer) Stmt {

}*/

func parse_while(st *lex.Lexer) run.Stmt {
	b := parse_exp(st)
	body := parse_body(st)
	return run.WhileStmt{Bool: b, Body: body}
}

func parse_body(st *lex.Lexer) []run.Stmt {
	return Parse(st)
}

func parse_let(st *lex.Lexer, name string) run.Stmt {
	e := parse_exp(st)
	return run.Stmt(run.LetStmt{Name: name, Equals: e})
}

func parse_ident(st *lex.Lexer, name string) run.Stmt {
	e := parse_exp(st)
	return run.Stmt(run.VarStmt{Name: name, Equals: e})
}

func parse_print(st *lex.Lexer) run.Stmt {
	e := parse_exp(st)
	return run.Stmt(run.PrintStmt{Print: e})
}

func parse_exp(st *lex.Lexer) run.Exp {
	return parse_bb(st)
}

func parse_paren(st *lex.Lexer) run.Exp {
	left := parse_bb(st)
	switch st.Peek().T {
	case lex.RParen:
		st.Eat()
		return left
	default:
		fmt.Println("paren not closed")
		return nil
	}
}

func parse_bb(st *lex.Lexer) run.Exp {
	left := parse_ub(st)
	switch st.Peek().T {
	case lex.GreatThan:
		st.Eat()
		return run.Exp(run.BinExp{Op: run.GreatThan, Left: left, Right: parse_bb(st)})
	case lex.LessThan:
		st.Eat()
		return run.Exp(run.BinExp{Op: run.LessThan, Left: left, Right: parse_bb(st)})
	case lex.GreatThanEq:
		st.Eat()
		return run.Exp(run.BinExp{Op: run.GreatThanEq, Left: left, Right: parse_bb(st)})
	case lex.LessThanEq:
		st.Eat()
		return run.Exp(run.BinExp{Op: run.LessThanEq, Left: left, Right: parse_bb(st)})
	case lex.EqEq:
		st.Eat()
		return run.Exp(run.BinExp{Op: run.EqEq, Left: left, Right: parse_bb(st)})
	case lex.NotEq:
		st.Eat()
		return run.Exp(run.BinExp{Op: run.NotEq, Left: left, Right: parse_bb(st)})
	case lex.And:
		st.Eat()
		return run.Exp(run.BinExp{Op: run.And, Left: left, Right: parse_bb(st)})
	case lex.Or:
		st.Eat()
		return run.Exp(run.BinExp{Op: run.Or, Left: left, Right: parse_bb(st)})
	default:
		return left
	}
}

func parse_ub(st *lex.Lexer) run.Exp {
	left := parse_pm(st)
	switch st.Peek().T {
	case lex.Not:
		st.Eat()
		return run.Exp(run.UnaExp{Op: run.Not, Right: parse_ub(st)})
	default:
		return left
	}
}

func parse_pm(st *lex.Lexer) run.Exp {
	left := parse_md(st)
	switch st.Peek().T {
	case lex.Add:
		st.Eat()
		return run.Exp(run.BinExp{Op: run.Add, Left: left, Right: parse_pm(st)})
	case lex.Sub:
		st.Eat()
		return run.Exp(run.BinExp{Op: run.Sub, Left: left, Right: parse_pm(st)})
	default:
		return left
	}
}

func parse_md(st *lex.Lexer) run.Exp {
	left := parse_um(st)
	switch st.Peek().T {
	case lex.Mul:
		st.Eat()
		return run.Exp(run.BinExp{Op: run.Mul, Left: left, Right: parse_md(st)})
	case lex.Div:
		st.Eat()
		return run.Exp(run.BinExp{Op: run.Div, Left: left, Right: parse_md(st)})
	case lex.Mod:
		st.Eat()
		return run.Exp(run.BinExp{Op: run.Mod, Left: left, Right: parse_md(st)})
	default:
		return left
	}
}

func parse_um(st *lex.Lexer) run.Exp {
	left := parse_lit(st)
	switch st.Peek().T {
	case lex.Sub:
		st.Eat()
		return run.Exp(run.UnaExp{Op: run.Neg, Right: parse_um(st)})
	default:
		return left
	}
}

func parse_lit(st *lex.Lexer) run.Exp {
	switch st.Peek().T {
	case lex.String:
		tok := st.Eat()
		return run.Exp(run.LitExp{Value: tok.Value})
	case lex.Integer:
		tok := st.Eat()
		return run.Exp(run.LitExp{Value: tok.Value})
	case lex.Bool:
		tok := st.Eat()
		return run.Exp(run.LitExp{Value: tok.Value})
	case lex.Ident:
		tok := st.Eat()
		return run.Exp(run.VarExp{Name: tok.Value.(string)})
	case lex.LParen:
		st.Eat()
		return parse_paren(st)
	default:
		fmt.Println("unknown literal")
		return run.Exp(run.LitExp{Value: nil})
	}
}
