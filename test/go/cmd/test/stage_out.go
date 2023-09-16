package main

import (
	"time"

	"github.com/fullstack-lang/gong/test/go/models"

	// injection point for ident package import declaration
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_stage_out models.StageStruct
var ___dummy__Time_stage_out time.Time

// Injection point for meta package dummy declaration
var ___dummy__gongdoc_models_stage_out gongdoc_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_stage_out map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"gongdoc_models.GongStructShape": &(gongdoc_models.GongStructShape{}),

	"gongdoc_models.GongStructShape.Name": (gongdoc_models.GongStructShape{}).Name,
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

	// Declarations of staged instances of Astruct

	// Declarations of staged instances of AstructBstruct2Use
	__AstructBstruct2Use__000000_ := (&models.AstructBstruct2Use{Name: ``}).Stage(stage)
	__AstructBstruct2Use__000001_ := (&models.AstructBstruct2Use{Name: ``}).Stage(stage)

	// Declarations of staged instances of AstructBstructUse

	// Declarations of staged instances of Bstruct
	__Bstruct__000000_B1 := (&models.Bstruct{Name: `B1`}).Stage(stage)
	__Bstruct__000001_B2 := (&models.Bstruct{Name: `B2`}).Stage(stage)
	__Bstruct__000002_B3 := (&models.Bstruct{Name: `B3`}).Stage(stage)

	// Declarations of staged instances of Dstruct
	__Dstruct__000000_D1 := (&models.Dstruct{Name: `D1`}).Stage(stage)

	// Setup of values

	// AstructBstruct2Use values setup
	__AstructBstruct2Use__000000_.Name = ``

	// AstructBstruct2Use values setup
	__AstructBstruct2Use__000001_.Name = ``

	// Bstruct values setup
	__Bstruct__000000_B1.Name = `B1`
	__Bstruct__000000_B1.Floatfield = 0.000000
	__Bstruct__000000_B1.Floatfield2 = 0.000000
	__Bstruct__000000_B1.Intfield = 0

	// Bstruct values setup
	__Bstruct__000001_B2.Name = `B2`
	__Bstruct__000001_B2.Floatfield = 0.000000
	__Bstruct__000001_B2.Floatfield2 = 0.000000
	__Bstruct__000001_B2.Intfield = 0

	// Bstruct values setup
	__Bstruct__000002_B3.Name = `B3`
	__Bstruct__000002_B3.Floatfield = 0.000000
	__Bstruct__000002_B3.Floatfield2 = 0.000000
	__Bstruct__000002_B3.Intfield = 0

	// Dstruct values setup
	__Dstruct__000000_D1.Name = `D1`

	// Setup of pointers
	__AstructBstruct2Use__000000_.Bstrcut2 = __Bstruct__000000_B1
	__AstructBstruct2Use__000001_.Bstrcut2 = __Bstruct__000000_B1
}


