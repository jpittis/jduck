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
	default:
		return "error"
	}
}

type UnaType int

const (
	Not UnaType = iota
	Neg
)

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
	Op    UnaType
	Right exp
}

func (e UnaExp) Eval() interface{} {
	return nil
}

type VarExp struct {
	Name string
}

func (e VarExp) Eval() interface{} {
	return nil
}

type FuncExp struct {
	Name   string
	Params []exp
}

func (e FuncExp) Eval() interface{} {
	return nil
}
