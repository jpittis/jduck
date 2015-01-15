package parse

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
)

type UnaryType int

const (
	Not UnaryType = iota
	Neg
	AddAdd
	SubSub
)

type exp interface {
	Eval() interface{}
}

type LitExp struct {
	value interface{}
}

type BinExp struct {
	Op    BinType
	Left  *exp
	Right *exp
}

type UnaExp struct {
	Operator UnaryType
	Right    *exp
}

type VarExp struct {
	Name string
}

type FuncExp struct {
	Name   string
	Params []*exp
}
