package lex

import (
	"bufio"
	"fmt"
)

type TokenType int

const (
	String TokenType = iota
	Integer

	Ident

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

type Token struct {
	T     TokenType
	Value interface{}
}

type Lexer struct {
	reader *bufio.Reader
	token  *Token
}

func (l *Lexer) Peek() *Token {
	return l.token
}

func (l *Lexer) Eat() *Token {
	temp := l.token
	l.token = l.lex()
	return temp
}

func (l *Lexer) read() (rune, error) { // TODO handle error
	r, _, err := l.reader.ReadRune()
	return r, err
}

func (l *Lexer) unread() error {
	return l.reader.UnreadRune()
}

func (l *Lexer) lex() (*Token, error) { // TODO handle error
	r := l.read()
	switch r {
	case '+':
		return &Token{T: Add}
	case '-':
		return &Token{T: Sub}
	case '*':
		return &Token{T: Mul}
	case '/':
		return &Token{T: Div}
	case '%':
		return &Token{T: Mod}
	case '(':
		return &Token{T: LParen}
	case ')':
		return &Token{T: RParen}
	case '=':
		r = l.read()
		if r == '=' {
			return &Token{T: EqEq}
		}
		l.unread()
		return &Token{T: Eq}
	case '>':
		r = l.read()
		if r == '=' {
			return &Token{T: GreatThanEq}
		}
		l.unread()
		return &Token{T: GreatThan}
	case '<':
		r = l.read()
		if r == '=' {
			return &Token{T: LessThanEq}
		}
		l.unread()
		return &Token{T: LessThan}
	case isNumber(r):
		var runes string
		for isNumber(r) {
			runes += r
			r = l.read()
		}
		l.unread()
		return strconv.Atoi(runes)
	case isLetter(r):
		var runes string
		for isLetter(r) {
			runes += r
			r = l.read()
		}
		l.unread()
		return lookup(runes)
	case isSpace(r):
		l.readSpace()
		return l.lex()
	default:
		return nil, fmt.Errorf("unknown rune '%s'", r)
	}
}

func (l *Lexer) readSpace() {

}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isSpace(r rune) bool {
	return ' ' || '\t' || '\n'
}

func lookup(ident string) *Token {
	return &Token{T: Ident, Value: ident}
}
