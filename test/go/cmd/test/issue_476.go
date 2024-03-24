package main

import (
	"time"

	"github.com/fullstack-lang/gong/test/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_issue_476 models.StageStruct
var ___dummy__Time_issue_476 time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_issue_476 map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["issue_476"] = issue_476Injection
// }

// issue_476Injection will stage objects of database "issue_476"
func issue_476Injection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Astruct

	// Declarations of staged instances of AstructBstruct2Use

	// Declarations of staged instances of AstructBstructUse

	// Declarations of staged instances of Bstruct

	// Declarations of staged instances of Dstruct

	// Declarations of staged instances of Fstruct
	__Fstruct__000000_F := (&models.Fstruct{Name: `F`}).Stage(stage)

	// Setup of values

	// Fstruct values setup
	__Fstruct__000000_F.Name = `F`
	__Fstruct__000000_F.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")

	// Setup of pointers
}


