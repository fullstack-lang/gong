package main

import (
	"time"

	"github.com/fullstack-lang/gong/test2/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_stage_out models.StageStruct
var ___dummy__Time_stage_out time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_stage_out map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["stage_out"] = stage_outInjection
// }

// stage_outInjection will stage objects of database "stage_out"
func stage_outInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Dummy
	__Dummy__000000_ := (&models.A{Name: ``}).Stage(stage)
	__Dummy__000001_ := (&models.A{Name: ``}).Stage(stage)
	__Dummy__000002_How_dummy_ := (&models.A{Name: `How dummy ?`}).Stage(stage)

	// Setup of values

	// Dummy values setup
	__Dummy__000000_.Name = ``

	// Dummy values setup
	__Dummy__000001_.Name = ``

	// Dummy values setup
	__Dummy__000002_How_dummy_.Name = `How dummy ?`

	// Setup of pointers
}
