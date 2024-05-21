package main

import (
	"time"

	"github.com/fullstack-lang/gong/test/go/models"

	// injection point for ident package import declaration
	dummy "github.com/fullstack-lang/gong/test/go/dummy"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_mini_stage_in models.StageStruct
var ___dummy__Time_mini_stage_in time.Time

// Injection point for meta package dummy declaration
var ___dummy__dummy_mini_stage_in dummy.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_mini_stage_in map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"dummy.A": dummy.A,

	"dummy.Dummy": &(dummy.Dummy{}),

	"dummy.Dummy.Name": (dummy.Dummy{}).Name,

	"dummy.DummyTypeInt": dummy.DummyTypeInt(0),

	"dummy.DummyTypeString": dummy.DummyTypeString(""),
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["mini_stage_in"] = mini_stage_inInjection
// }

// mini_stage_inInjection will stage objects of database "mini_stage_in"
func mini_stage_inInjection(stage *models.StageStruct) {

	__Bstruct__000000_B1 := (&models.Bstruct{Name: `B1`}).Stage(stage)

	__Astruct__000000_Foo := (&models.Astruct{
		Name: `Foo`,
		Date: func() time.Time {
			t, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
			return t
		}(),
		Booleanfield: false,
		CEnum:        models.CENUM_VAL1,
		Cstruct: models.Cstruct{
			CName:       ``,
			CFloatfield: 0.000000,
			Bstruct:     __Bstruct__000000_B1,
		},
		Floatfield:          0.000000,
		Intfield:            0,
		Anotherbooleanfield: false,
		Duration1:           0,
		Associationtob:      __Bstruct__000000_B1,
	}).Stage(stage)

	__Astruct__000000_Foo.Name = `Foo`

	// comment added to overcome the problem with the comment map association

	//gong:ident [dummy.Dummy2]
	__Astruct__000000_Foo.StructRef = `dummy.Dummy2`

	// comment added to overcome the problem with the comment map association

	//gong:ident [dummy.Dummy2.Name]
	__Astruct__000000_Foo.FieldRef = `dummy.Dummy2.Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [dummy.DummyTypeInt]
	__Astruct__000000_Foo.EnumIntRef = `dummy.DummyTypeInt`

	// comment added to overcome the problem with the comment map association

	//gong:ident [dummy.DummyTypeString]
	__Astruct__000000_Foo.EnumStringRef = `dummy.DummyTypeString`

	// comment added to overcome the problem with the comment map association

	//gong:ident [dummy.A]
	__Astruct__000000_Foo.EnumValue = `dummy.A`

	// comment added to overcome the problem with the comment map association

	//gong:ident [..]
	__Astruct__000000_Foo.ConstIdentifierValue = `..`

	__Bstruct__000000_B1.Name = `B1`
	__Bstruct__000000_B1.Floatfield = 0.000000
	__Bstruct__000000_B1.Floatfield2 = 0.000000
	__Bstruct__000000_B1.Intfield = 0

	// Setup of pointers
	__Astruct__000000_Foo.Bstruct = __Bstruct__000000_B1
	__Astruct__000000_Foo.Anarrayofb = append(__Astruct__000000_Foo.Anarrayofb, __Bstruct__000000_B1)
}
