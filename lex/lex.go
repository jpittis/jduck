// Package lex provides methods for lexing from a reader.
package lex

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// TokenType constant for each kind of token.
type TokenType int

const (
	String TokenType = iota
	Integer

	Ident

	EOF

	Eq

	Add
	Sub
	Mul
	Div
	Mod

	LParen
	RParen

	Bool

	EqEq
	LessThan
	GreatThan
	LessThanEq
	GreatThanEq

	And
	Or
)

// Token holds a type and a possible value.
type Token struct {
	T     TokenType
	Value interface{}
}

// Lexer is used to lex tokens from a reader.
type Lexer struct {
	reader *bufio.Reader
	token  *Token
}

func New(reader io.Reader) *Lexer {
	return &Lexer{reader: bufio.NewReader(reader)}
}

// Peek returns the last read token.
func (l *Lexer) Peek() (*Token, error) {
	if l.token == nil {
		t, err := l.lex()
		if err != nil {
			return nil, err
		}
		l.token = t
	}
	return l.token, nil
}

// Eat returns the last read token while reading another.
func (l *Lexer) Eat() (*Token, error) {
	if l.token == nil {
		t, err := l.lex()
		if err != nil {
			return nil, err
		}
		l.token = t
	}
	temp := l.token
	t, err := l.lex()
	if err != nil {
		return nil, err
	}
	l.token = t
	return temp, nil
}

func (l *Lexer) read() (rune, error) {
	r, _, err := l.reader.ReadRune()
	return r, err
}

func (l *Lexer) unread() {
	// error not handled because only ReadRune() used
	l.reader.UnreadRune()
}

func (l *Lexer) lex() (*Token, error) {
	r, err := l.read()
	if r == rune(0) {
		return &Token{T: EOF}, nil
	}
	if err != nil {
		return nil, err
	}
	switch {
	case r == '+':
		return &Token{T: Add}, nil
	case r == '-':
		return &Token{T: Sub}, nil
	case r == '*':
		return &Token{T: Mul}, nil
	case r == '/':
		return &Token{T: Div}, nil
	case r == '%':
		return &Token{T: Mod}, nil
	case r == '(':
		return &Token{T: LParen}, nil
	case r == ')':
		return &Token{T: RParen}, nil
	case r == '=':
		r, err = l.read()
		if err != nil {
			return nil, err
		}
		if r == '=' {
			return &Token{T: EqEq}, nil
		}
		l.unread()
		return &Token{T: Eq}, nil
	case r == '>':
		r, err = l.read()
		if err != nil {
			return nil, err
		}
		if r == '=' {
			return &Token{T: GreatThanEq}, nil
		}
		l.unread()
		return &Token{T: GreatThan}, nil
	case r == '<':
		r, err = l.read()
		if err != nil {
			return nil, err
		}
		if r == '=' {
			return &Token{T: LessThanEq}, nil
		}
		l.unread()
		return &Token{T: LessThan}, nil
	case r == '"':
		var runes string
		r, err = l.read()
		if err != nil {
			return nil, err
		}
		for r != '"' {
			if r == rune(0) {
				return nil, fmt.Errorf("end of file before end of string")
			}
			runes += string(r)
			r, err = l.read()
			if err != nil {
				return nil, err
			}
		}
		return &Token{T: String, Value: runes}, nil
	case isNumber(r):
		var runes string
		for isNumber(r) {
			runes += string(r)
			r, err = l.read()
			if err != nil {
				return nil, err
			}
		}
		l.unread()
		i, err := strconv.Atoi(runes)
		return &Token{T: Integer, Value: i}, err
	case isLetter(r):
		var runes string
		for isLetter(r) {
			runes += string(r)
			r, err = l.read()
			if err != nil {
				return nil, err
			}
		}
		l.unread()
		return lookup(runes), nil
	case isSpace(r):
		err = l.readSpace()
		if err != nil {
			return nil, err
		}
		return l.lex()
	default:
		return nil, fmt.Errorf("unknown rune '%s'", r)
	}
}

func (l *Lexer) readSpace() error {
	r, err := l.read()
	if r == rune(0) {
		l.unread()
		return nil
	}
	if err != nil {
		return err
	}
	for isSpace(r) {
		r, err = l.read()
		if err != nil {
			return err
		}
	}
	l.unread()
	return nil
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}

func lookup(ident string) *Token {
	switch ident {
	case "true":
		return &Token{T: Bool, Value: true}
	case "false":
		return &Token{T: Bool, Value: false}
	case "and":
		return &Token{T: And}
	case "or":
		return &Token{T: Or}
	default:
		return &Token{T: Ident, Value: ident}
	}
}
