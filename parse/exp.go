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

func (e LitExp) Eval() interface{} {
	return nil
}

type BinExp struct {
	Op    BinType
	Left  exp
	Right exp
}

func (e BinExp) Eval() interface{} {
	return nil
}

type UnaExp struct {
	Operator UnaryType
	Right    exp
}

func (e *UnaExp) Eval() interface{} {
	return nil
}

type VarExp struct {
	Name string
}

func (e *VarExp) Eval() interface{} {
	return nil
}

type FuncExp struct {
	Name   string
	Params []exp
}

func (e *FuncExp) Eval() interface{} {
	return nil
}
