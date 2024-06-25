package main

import (
	"time"

	"github.com/fullstack-lang/gong/test/go/models"

	// injection point for ident package import declaration
	dummy "github.com/fullstack-lang/gong/test/go/dummy"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ models.StageStruct
var _ time.Time

// Injection point for meta package dummy declaration
var _ dummy.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
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
// 	InjectionGateway["mini_stage_out"] = mini_stage_outInjection
// }

// mini_stage_outInjection will stage objects of database "mini_stage_out"
var _ = func(stage *models.StageStruct) {

	// Declaration of instances to stage
	__Astruct__000000_Foo := (&models.Astruct{Name: `Foo`}).Stage(stage)
	__Bstruct__000000_B1 := (&models.Bstruct{Name: `B1`}).Stage(stage)

	// Setup of values
	__Astruct__000000_Foo.Name = `Foo`
	__Astruct__000000_Foo.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__000000_Foo.Booleanfield = false
	__Astruct__000000_Foo.CEnum = models.CENUM_VAL1
	__Astruct__000000_Foo.CName = ``
	__Astruct__000000_Foo.CFloatfield = 0.000000
	__Astruct__000000_Foo.Floatfield = 0.000000
	__Astruct__000000_Foo.Intfield = 0
	__Astruct__000000_Foo.Anotherbooleanfield = false
	__Astruct__000000_Foo.Duration1 = 0

	//gong:ident [dummy.Dummy2] comment added to overcome the problem with the comment map association
	__Astruct__000000_Foo.StructRef = `dummy.Dummy2`

	//gong:ident [dummy.Dummy2.Name] comment added to overcome the problem with the comment map association
	__Astruct__000000_Foo.FieldRef = `dummy.Dummy2.Name`

	//gong:ident [dummy.DummyTypeInt] comment added to overcome the problem with the comment map association
	__Astruct__000000_Foo.EnumIntRef = `dummy.DummyTypeInt`

	//gong:ident [dummy.DummyTypeString] comment added to overcome the problem with the comment map association
	__Astruct__000000_Foo.EnumStringRef = `dummy.DummyTypeString`

	//gong:ident [dummy.A] comment added to overcome the problem with the comment map association
	__Astruct__000000_Foo.EnumValue = `dummy.A`

	//gong:ident [...] comment added to overcome the problem with the comment map association
	__Astruct__000000_Foo.ConstIdentifierValue = `...`
	__Astruct__000000_Foo.TextFieldBespokeSize = ``
	__Astruct__000000_Foo.TextArea = ``
	__Bstruct__000000_B1.Name = `B1`
	__Bstruct__000000_B1.Floatfield = 0.000000
	__Bstruct__000000_B1.Floatfield2 = 0.000000
	__Bstruct__000000_B1.Intfield = 0

	// Setup of pointers
	__Astruct__000000_Foo.Anarrayofb = append(__Astruct__000000_Foo.Anarrayofb, __Bstruct__000000_B1)
}
