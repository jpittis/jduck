package parse

type BinaryType int

const (
	Add BinaryType = iota
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
)

type UnaryType int

const (
	Not UnaryType = iota
	Min
	AddAdd
	SubSub
)

type exp interface {
	Eval() Literal
}

type LiteralExp struct {
	value interface{}
}

type BinaryExp struct {
	Operator BinaryType
	Left     exp
	Right    exp
}

type UnaryExp struct {
	Operator UnaryType
	Right    exp
}
