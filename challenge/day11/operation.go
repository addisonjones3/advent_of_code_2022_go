package day11

type Operator string

const (
	mult Operator = "*"
	add  Operator = "+"
	sub  Operator = "-"
)

type Operation struct {
	Operator
	SelfOp bool
	Value  int
}

type TestOperation struct {
	Val int
}
