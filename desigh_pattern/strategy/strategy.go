package strategy

type Operator interface {
	Apply(a, b int) int
	GetName() string
}

type Operation struct {
	op Operator
}

func (oprn Operation) Operate(a, b int) int {
	return oprn.op.Apply(a, b)
}

type Add struct{}

func (op *Add) Apply(a, b int) int {
	return a + b

}
func (op *Add) GetName() string {
	return "+"
}

type Minus struct{}

func (op *Minus) Apply(a, b int) int {
	return a - b

}
func (op *Minus) GetName() string {
	return "-"
}

type Div struct{}

func (op *Div) Apply(a, b int) int {
	return a / b

}
func (op *Div) GetName() string {
	return "/"
}

type Multi struct{}

func (op *Multi) Apply(a, b int) int {
	return a * b

}
func (op *Multi) GetName() string {
	return "*"
}
