package gotest

type Calculator struct {
	Ia int
	Ib int
}

func (self Calculator) Add() int {
	return self.Ia + self.Ib
}
