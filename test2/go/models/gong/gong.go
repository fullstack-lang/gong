package gong

import (
	"github.com/fullstack-lang/gong/test2/go/models"
	"github.com/fullstack-lang/gong/test2/go/models/x"
	"github.com/fullstack-lang/gong/test2/go/models/y"
)

type StageStruct struct {
	Models
	X
	Y
}

type Models struct {
	models.StageStruct
}

type X struct {
	x.StageStruct
}

type Y struct {
	y.StageStruct
}

func Dummy() {
	a := new(models.A)
	_ = a

	a_c := new(x.X_A)

	stage := new(StageStruct)
	stage.Models.As[a] = true
	stage.X.X_As[a_c] = true

}
