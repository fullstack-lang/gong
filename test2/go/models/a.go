package models

import (
	"github.com/fullstack-lang/gong/test2/go/models/y"
)

type A struct {
	Name string

	Y_A *y.Y_A
}

func (a *A) Foo() {
	// stage := gong.NewStage(path)
}
