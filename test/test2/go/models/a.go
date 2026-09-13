package models

import (
	"github.com/fullstack-lang/gong/test/test2/go/models/x"
)

type A struct {
	Name string

	NumberField int

	B *B

	Bs []*B

	X *x.A
}

func (a *A) Foo() {
	// stage := gong.NewStage(path)
}
