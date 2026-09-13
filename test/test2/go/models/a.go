package models

import (
	"github.com/fullstack-lang/gong/test/test2/go/models/x"
)

type A struct {
	Name string

	NumberField int

	B *B

	Bs []*B

	X *x.X

	// test the those field can be imported because we use the analyser now
	x.ToBeImported
}

func (a *A) Foo() {
	// stage := gong.NewStage(path)
}
