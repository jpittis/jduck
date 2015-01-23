package run

import (
	"log"
)

// Binary operator constants.
type BinType int

const (
	Add BinType = iota
	Sub
	Mul
	Div
	Mod

	GreatThan
	LessThan
	GreatThanEq
	LessThanEq
	EqEq
	NotEq

	And
	Or
)

// Unary operator constants.
type UnaType int

const (
	Not UnaType = iota
	Neg
)

type Exp interface {
	Eval(*Context) interface{}
}

type LitExp struct {
	Value interface{}
}

func (e LitExp) Eval(c *Context) interface{} {
	return e.Value
}

type BinExp struct {
	Op    BinType
	Left  Exp
	Right Exp
}

func (e BinExp) Eval(c *Context) interface{} {
	lhs := e.Left.Eval(c)
	rhs := e.Right.Eval(c)
	switch e.Op {
	case Add:
		switch lhs.(type) {
		case bool:
			log.Fatal("cannot add booleans")
		case int:
			return lhs.(int) + rhs.(int)
		case string:
			return lhs.(string) + rhs.(string)
		}
	case Sub:
		switch lhs.(type) {
		case bool:
			log.Fatal("cannot sub booleans")
		case int:
			return lhs.(int) - rhs.(int)
		case string:
			log.Fatal("cannot sub strings")
		}
	case Mul:
		switch lhs.(type) {
		case bool:
			log.Fatal("cannot mul booleans")
		case int:
			return lhs.(int) * rhs.(int)
		case string:
			log.Fatal("cannot mul strings")
		}
	case Div:
		switch lhs.(type) {
		case bool:
			log.Fatal("cannot div booleans")
		case int:
			return lhs.(int) / rhs.(int)
		case string:
			log.Fatal("cannot div strings")
		}
	case Mod:
		switch lhs.(type) {
		case bool:
			log.Fatal("cannot mod booleans")
		case int:
			return lhs.(int) % rhs.(int)
		case string:
			log.Fatal("cannot mod strings")
		}
	case GreatThan:
		switch lhs.(type) {
		case bool:
			log.Fatal("boolean not int")
		case int:
			return lhs.(int) > rhs.(int)
		case string:
			log.Fatal("string not int")
		}
	case LessThan:
		switch lhs.(type) {
		case bool:
			log.Fatal("boolean not int")
		case int:
			return lhs.(int) < rhs.(int)
		case string:
			log.Fatal("string not int")
		}
	case GreatThanEq:
		switch lhs.(type) {
		case bool:
			log.Fatal("boolean not int")
		case int:
			return lhs.(int) >= rhs.(int)
		case string:
			log.Fatal("string not int")
		}
	case LessThanEq:
		switch lhs.(type) {
		case bool:
			log.Fatal("boolean not int")
		case int:
			return lhs.(int) <= rhs.(int)
		case string:
			log.Fatal("string not int")
		}
	case EqEq:
		switch lhs.(type) {
		case bool:
			return lhs.(bool) == rhs.(bool)
		case int:
			return lhs.(int) == rhs.(int)
		case string:
			return lhs.(string) == rhs.(string)
		}
	case NotEq:
		switch lhs.(type) {
		case bool:
			return lhs.(bool) != rhs.(bool)
		case int:
			return lhs.(int) != rhs.(int)
		case string:
			return lhs.(string) != rhs.(string)
		}
	case And:
		switch lhs.(type) {
		case bool:
			return lhs.(bool) && rhs.(bool)
		case int:
			log.Fatal("int not boolean")
		case string:
			log.Fatal("string not boolean")
		}
	case Or:
		switch lhs.(type) {
		case bool:
			return lhs.(bool) || rhs.(bool)
		case int:
			log.Fatal("int not boolean")
		case string:
			log.Fatal("string not boolean")
		}
	}
	log.Fatal("unknown type")
	return nil
}

type UnaExp struct {
	Op    UnaType
	Right Exp
}

func (e UnaExp) Eval(c *Context) interface{} {
	rhs := e.Right.Eval(c)
	switch e.Op {
	case Not:
		switch rhs.(type) {
		case bool:
			return !rhs.(bool)
		case int:
			log.Fatal("int not boolean")
		case string:
			log.Fatal("string not boolean")
		}
	case Neg:
		switch rhs.(type) {
		case bool:
			log.Fatal("bool not int")
		case int:
			return rhs.(int) * -1
		case string:
			log.Fatal("string not int")
		}
	}
	log.Fatal("type not found")
	return nil
}

type VarExp struct {
	Name string
}

func (e VarExp) Eval(c *Context) interface{} {
	val, err := c.Get(e.Name)
	if err != nil {
		log.Fatal("variable not declared")
	}
	return val
}

/*type FuncExp struct {
	Name   string
	Params []exp
}

func (e FuncExp) Eval() interface{} {
	return nil
}*/

// BinType prints to these values.
func (b BinType) String() string {
	switch b {
	case Add:
		return "+"
	case Sub:
		return "-"
	case Mul:
		return "*"
	case Div:
		return "/"
	case Mod:
		return "%"
	case GreatThan:
		return ">"
	case LessThan:
		return "<"
	case GreatThanEq:
		return ">="
	case LessThanEq:
		return "<="
	case EqEq:
		return "=="
	case NotEq:
		return "!="
	case And:
		return "&&"
	case Or:
		return "||"
	default:
		return "error"
	}
}

// UnaType prints to these values.
func (u UnaType) String() string {
	switch u {
	case Not:
		return "!"
	case Neg:
		return "-"
	default:
		return "error"
	}
}
