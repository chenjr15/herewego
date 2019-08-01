package method

type employee struct {
	salary float64
}

func (e *employee) giveRaise(percent float64) {
	e.salary *= 1.0 + percent
}
