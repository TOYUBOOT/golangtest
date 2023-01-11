package gotest

type calculator struct {
	a int
	b int
}

func (self calculator) add() int {
	return self.a + self.b
}
