package main

import (
	"time"

	"github.com/fullstack-lang/gong/test/go/models"

	// injection point for ident package import declaration
	dummy "github.com/fullstack-lang/gong/test/go/dummy"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_mini_stage_out models.StageStruct
var ___dummy__Time_mini_stage_out time.Time

// Injection point for meta package dummy declaration
var ___dummy__dummy_mini_stage_out dummy.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_mini_stage_out map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"dummy.Dummy": &(dummy.Dummy{}),

	"dummy.Dummy.Name": (dummy.Dummy{}).Name,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["mini_stage_out"] = mini_stage_outInjection
// }

// mini_stage_outInjection will stage objects of database "mini_stage_out"
func mini_stage_outInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Astruct
	__Astruct__000000_Foo := (&models.Astruct{Name: `Foo`}).Stage()

	// Declarations of staged instances of AstructBstruct2Use

	// Declarations of staged instances of AstructBstructUse

	// Declarations of staged instances of Bstruct

	// Declarations of staged instances of Dstruct

	// Setup of values

	// Astruct values setup
	__Astruct__000000_Foo.Name = `Foo`
	__Astruct__000000_Foo.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 +0000")
	__Astruct__000000_Foo.Booleanfield = false
	__Astruct__000000_Foo.CEnum = models.CENUM_VAL1
	__Astruct__000000_Foo.CName = ``
	__Astruct__000000_Foo.CFloatfield = 0.000000
	__Astruct__000000_Foo.Floatfield = 0.000000
	__Astruct__000000_Foo.Intfield = 0
	__Astruct__000000_Foo.Anotherbooleanfield = false
	__Astruct__000000_Foo.Duration1 = 0
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [dummy.Dummy]
	__Astruct__000000_Foo.StructRef = `dummy.Dummy`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [dummy.Dummy.Name]
	__Astruct__000000_Foo.FieldRef = `dummy.Dummy.Name`

	// Setup of pointers
}


