package main

import (
	"time"

	"github.com/fullstack-lang/gong/test/go/models"

	// injection point for ident package import declaration
	dummy "github.com/fullstack-lang/gong/test/go/dummy"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// Injection point for meta package dummy declaration
var _ dummy.StageStruct

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Astruct__000000_ := (&models.Astruct{Name: ``}).Stage(stage)

	// Setup of values

	__Astruct__000000_.Name = ``
	__Astruct__000000_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__000000_.Booleanfield = false
	__Astruct__000000_.CEnum = models.CENUM_VAL1
	__Astruct__000000_.CName = ``
	__Astruct__000000_.CFloatfield = 0.000000
	__Astruct__000000_.Floatfield = 0.000000
	__Astruct__000000_.Intfield = 0
	__Astruct__000000_.Anotherbooleanfield = false
	__Astruct__000000_.Duration1 = 0

	//gong:ident [dummy.Dummy] comment added to overcome the problem with the comment map association
	__Astruct__000000_.StructRef = `dummy.Dummy`

	//gong:ident [.] comment added to overcome the problem with the comment map association
	__Astruct__000000_.FieldRef = `.`

	//gong:ident [.] comment added to overcome the problem with the comment map association
	__Astruct__000000_.EnumIntRef = `.`

	//gong:ident [.] comment added to overcome the problem with the comment map association
	__Astruct__000000_.EnumStringRef = `.`

	//gong:ident [.] comment added to overcome the problem with the comment map association
	__Astruct__000000_.EnumValue = `.`

	//gong:ident [.] comment added to overcome the problem with the comment map association
	__Astruct__000000_.ConstIdentifierValue = `.`
	__Astruct__000000_.TextFieldBespokeSize = ``
	__Astruct__000000_.TextArea = ``

	// Setup of pointers
}
