// Package lex provides methods for lexing jduck files.
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
	Bool

	If
	Else
	While
	For
	Func
	End

	Print

	Ident

	Eq

	Add
	Sub
	Mul
	Div
	Mod

	LParen
	RParen

	LessThan
	GreatThan
	LessThanEq
	GreatThanEq
	EqEq
	NotEq
	And
	Or

	Not

	EOL // do I even need EOL?
	EOF
)

// Token holds a type and a possible value.
type Token struct {
	T     TokenType   // type of token
	Value interface{} // value of type
}

// Lexer is used to lex tokens from a reader.
type Lexer struct {
	reader *bufio.Reader // reader for jduck file
	token  *Token        // token buffer
}

// New creates a new lexer on given file.
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
	temp, err := l.Peek()
	if err != nil {
		return nil, err
	}
	t, err := l.lex()
	if err != nil {
		return nil, err
	}
	l.token = t
	return temp, nil
}

// Reads one rune from the file.
func (l *Lexer) read() (rune, error) {
	r, _, err := l.reader.ReadRune()
	return r, err
}

func (l *Lexer) unread() {
	l.reader.UnreadRune() // not handled because only ReadRune() used
}

func (l *Lexer) lex() (*Token, error) {
	r, err := l.read()
	if err == io.EOF {
		return &Token{T: EOF}, nil // return end of file token
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
	case r == ';':
		return &Token{T: EOL}, nil
	case r == '!':
		r, err = l.read()
		if err != nil {
			return nil, err
		}
		if r == '=' {
			return &Token{T: NotEq}, nil
		}
		l.unread()
		return &Token{T: Not}, nil
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
		r, err := l.read()
		for r != '"' {
			if err != nil {
				return nil, err
			}
			runes += string(r)
			r, err = l.read()
		}
		return &Token{T: String, Value: runes}, nil
	case isNumber(r):
		var runes string
		r, err := l.read()
		for isNumber(r) {
			if err != nil {
				return nil, err
			}
			runes += string(r)
			r, err = l.read()
		}
		l.unread()
		i, err := strconv.Atoi(runes)
		return &Token{T: Integer, Value: i}, err
	case isLetter(r):
		var runes string
		r, err := l.read()
		for isLetter(r) {
			if err != nil {
				return nil, err
			}
			runes += string(r)
			r, err = l.read()
		}
		l.unread()
		return lookup(runes), nil // return either ident or keyword token
	case isSpace(r):
		err = l.readSpace() // remove space
		if err == io.EOF {
			return &Token{T: EOF}, nil
		}
		if err != nil {
			return nil, err
		}
		return l.lex() // continue lexing
	default:
		return nil, fmt.Errorf("unknown rune '%v'", r)
	}
}

// Reads all space up to next none space rune.
func (l *Lexer) readSpace() error {
	r, err := l.read()
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

// Returns true if number.
func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

// Returns true if letter.
func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

// Returns true if space, tab or newline.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}

// Returns keyword token or default ident token.
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
	case "if":
		return &Token{T: If}
	case "else":
		return &Token{T: Else}
	case "for":
		return &Token{T: For}
	case "while":
		return &Token{T: While}
	case "end":
		return &Token{T: End}
	case "print":
		return &Token{T: Print}
	case "func":
		return &Token{T: Func}
	default:
		return &Token{T: Ident, Value: ident}
	}
}
