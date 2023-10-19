package main

import (
	"time"

	"github.com/fullstack-lang/gong/test2/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_stage2 models.StageStruct
var ___dummy__Time_stage2 time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_stage2 map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["stage2"] = stage2Injection
// }

// stage2Injection will stage objects of database "stage2"
func stage2Injection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of A
	__A__000000_A1 := (&models.A{Name: `A1`}).Stage(stage)

	// Declarations of staged instances of B
	__B__000000_B1 := (&models.B{Name: `B1`}).Stage(stage)
	__B__000001_B2 := (&models.B{Name: `B2`}).Stage(stage)
	__B__000002_B3 := (&models.B{Name: `B3`}).Stage(stage)

	// Setup of values

	// A values setup
	__A__000000_A1.Name = `A1`
	__A__000000_A1.NumberField = 40

	// B values setup
	__B__000000_B1.Name = `B1`

	// B values setup
	__B__000001_B2.Name = `B2`

	// B values setup
	__B__000002_B3.Name = `B3`

	// Setup of pointers
	__A__000000_A1.B = __B__000001_B2
	__A__000000_A1.Bs = append(__A__000000_A1.Bs, __B__000000_B1)
	__A__000000_A1.Bs = append(__A__000000_A1.Bs, __B__000001_B2)
	__A__000000_A1.Bs = append(__A__000000_A1.Bs, __B__000002_B3)
}


