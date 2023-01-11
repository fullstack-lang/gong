package main

import (
	"time"

	"github.com/fullstack-lang/gong/test/go/models"

	// injection point for ident package import declaration
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

// Injection point for meta package dummy declaration
var ___dummy__gongdoc_Stage2 gongdoc_models.StageStruct

var __Dummy_time_variable time.Time

// stageInjection will stage objects of database "stage"
func stageInjection2() {

	// Declarations of staged instances of Astruct
	__Astruct__000000_A1 := (&models.Astruct{Name: `A1`}).Stage()

	__Astruct__000000_A1.Name = "Foo"

	//gong:ident [gongdoc_models.Classshape]
	__Astruct__000000_A1.Ref = ""

}
