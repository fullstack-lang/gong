package main

import (
	"time"

	"github.com/fullstack-lang/gong/test/go/models"

	// injection point for ident package import declaration
	dummy "github.com/fullstack-lang/gong/test/go/dummy"
)

// Injection point for meta package dummy declaration
var ___dummy__gongdoc_Stage2 dummy.StageStruct

var __Dummy_time_variable time.Time

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier map[string]any = map[string]any{
	// injection point for docLink to identifiers
	"dummy.Dummy":      &(dummy.Dummy2{}),
	"dummy.Dummy.Name": (dummy.Dummy2{}).Name,
}

// stageInjection will stage objects of database "stage"
func stageInjection2() {

	// Declarations of staged instances of Astruct
	__Astruct__000000_A1 := (&models.Astruct{Name: `A1`}).Stage()

	__Astruct__000000_A1.Name = "Foo"

	//gong:ident [dummy.Dummy]
	__Astruct__000000_A1.StructRef = ""

	// comment added to overcome the problem with the comment map association

	//gong:ident [dummy.Dummy.Name]
	__Astruct__000000_A1.FieldRef = ""

}
