package main

import (
	"time"

	"github.com/fullstack-lang/gong/test/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_stageB models.StageStruct
var ___dummy__Time_stageB time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_stageB map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["stageB"] = stageBInjection
// }

// stageBInjection will stage objects of database "stageB"
func stageBInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Astruct
	__Astruct__000000_a1 := (&models.Astruct{Name: `a1`}).Stage(stage)

	// Declarations of staged instances of AstructBstruct2Use

	// Declarations of staged instances of AstructBstructUse

	// Declarations of staged instances of Bstruct

	// Declarations of staged instances of Dstruct

	// Setup of values

	// Astruct values setup
	__Astruct__000000_a1.Name = `a1`
	__Astruct__000000_a1.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2023-03-17 07:52:44.609 +0000 +0000")
	__Astruct__000000_a1.Booleanfield = true
	__Astruct__000000_a1.Aenum = models.ENUM_VAL2
	__Astruct__000000_a1.CEnum = models.CENUM_VAL1
	__Astruct__000000_a1.CName = ``
	__Astruct__000000_a1.CFloatfield = 45454.000000
	__Astruct__000000_a1.Floatfield = 0.000000
	__Astruct__000000_a1.Intfield = 0
	__Astruct__000000_a1.Anotherbooleanfield = false
	__Astruct__000000_a1.Duration1 = 0

	// comment added to overcome the problem with the comment map association

	//gong:ident [.]
	__Astruct__000000_a1.StructRef = `.`

	// comment added to overcome the problem with the comment map association

	//gong:ident [.]
	__Astruct__000000_a1.FieldRef = `.`

	// comment added to overcome the problem with the comment map association

	//gong:ident [.]
	__Astruct__000000_a1.EnumIntRef = `.`

	// comment added to overcome the problem with the comment map association

	//gong:ident [.]
	__Astruct__000000_a1.EnumStringRef = `.`

	// comment added to overcome the problem with the comment map association

	//gong:ident [.]
	__Astruct__000000_a1.EnumValue = `.`

	// comment added to overcome the problem with the comment map association

	//gong:ident [.]
	__Astruct__000000_a1.ConstIdentifierValue = `.`

	// Setup of pointers
}


