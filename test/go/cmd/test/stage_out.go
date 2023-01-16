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

	"gongdoc_models.Classshape": &(gongdoc_models.Classshape{}),

	"gongdoc_models.Classshape.Name": (gongdoc_models.Classshape{}).Name,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["stage_out"] = stage_outInjection
// }

// stage_outInjection will stage objects of database "stage_out"
func stage_outInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Astruct
	__Astruct__000000_A1 := (&models.Astruct{Name: `A1`}).Stage()
	__Astruct__000001_A2 := (&models.Astruct{Name: `A2`}).Stage()
	__Astruct__000002_A3 := (&models.Astruct{Name: `A3`}).Stage()

	// Declarations of staged instances of AstructBstruct2Use
	__AstructBstruct2Use__000000_ := (&models.AstructBstruct2Use{Name: ``}).Stage()
	__AstructBstruct2Use__000001_ := (&models.AstructBstruct2Use{Name: ``}).Stage()

	// Declarations of staged instances of AstructBstructUse

	// Declarations of staged instances of Bstruct
	__Bstruct__000000_B1 := (&models.Bstruct{Name: `B1`}).Stage()
	__Bstruct__000001_B2 := (&models.Bstruct{Name: `B2`}).Stage()
	__Bstruct__000002_B3 := (&models.Bstruct{Name: `B3`}).Stage()

	// Declarations of staged instances of Dstruct
	__Dstruct__000000_D1 := (&models.Dstruct{Name: `D1`}).Stage()

	// Setup of values

	// Astruct values setup
	__Astruct__000000_A1.Name = `A1`
	__Astruct__000000_A1.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-11-14 03:08:21 +0000 +0000")
	__Astruct__000000_A1.Booleanfield = false
	__Astruct__000000_A1.Aenum = models.ENUM_VAL1
	__Astruct__000000_A1.Aenum_2 = models.ENUM_VAL2
	__Astruct__000000_A1.Benum = models.BENUM_VAL2
	__Astruct__000000_A1.CEnum = models.CENUM_VAL2
	__Astruct__000000_A1.CName = `CName1

	\n"""" fdfsqsfhqksf
Second Line`
	__Astruct__000000_A1.CFloatfield = 60.500000
	__Astruct__000000_A1.Floatfield = 0.000000
	__Astruct__000000_A1.Intfield = 3
	__Astruct__000000_A1.Anotherbooleanfield = false
	__Astruct__000000_A1.Duration1 = 79653000000000
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [gongdoc_models.Classshape]
	__Astruct__000000_A1.StructRef = `gongdoc_models.Classshape`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [gongdoc_models.Classshape.Name]
	__Astruct__000000_A1.FieldRef = `gongdoc_models.Classshape.Name`

	// Astruct values setup
	__Astruct__000001_A2.Name = `A2`
	__Astruct__000001_A2.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-01-18 01:09:21 +0000 +0000")
	__Astruct__000001_A2.Booleanfield = false
	__Astruct__000001_A2.Aenum = models.ENUM_VAL1
	__Astruct__000001_A2.Aenum_2 = models.ENUM_VAL2
	__Astruct__000001_A2.Benum = models.BENUM_VAL1
	__Astruct__000001_A2.CEnum = models.CENUM_VAL1
	__Astruct__000001_A2.CName = ``
	__Astruct__000001_A2.CFloatfield = 0.100000
	__Astruct__000001_A2.Floatfield = 0.000000
	__Astruct__000001_A2.Intfield = 0
	__Astruct__000001_A2.Anotherbooleanfield = false
	__Astruct__000001_A2.Duration1 = 0
	
	// comment added to overcome the problem with the comment map association

	//gong:ident []
	__Astruct__000001_A2.StructRef = ``
	
	// comment added to overcome the problem with the comment map association

	//gong:ident []
	__Astruct__000001_A2.FieldRef = ``

	// Astruct values setup
	__Astruct__000002_A3.Name = `A3`
	__Astruct__000002_A3.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-02-10 01:06:11.446 +0000 +0000")
	__Astruct__000002_A3.Booleanfield = true
	__Astruct__000002_A3.Benum = models.BENUM_VAL2
	__Astruct__000002_A3.CEnum = models.CENUM_VAL1
	__Astruct__000002_A3.CName = ``
	__Astruct__000002_A3.CFloatfield = 4.900000
	__Astruct__000002_A3.Floatfield = 2.000000
	__Astruct__000002_A3.Intfield = 0
	__Astruct__000002_A3.Anotherbooleanfield = false
	__Astruct__000002_A3.Duration1 = 0
	
	// comment added to overcome the problem with the comment map association

	//gong:ident []
	__Astruct__000002_A3.StructRef = ``
	
	// comment added to overcome the problem with the comment map association

	//gong:ident []
	__Astruct__000002_A3.FieldRef = ``

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
	__Astruct__000000_A1.Associationtob = __Bstruct__000001_B2
	__Astruct__000000_A1.Anarrayofb = append(__Astruct__000000_A1.Anarrayofb, __Bstruct__000002_B3)
	__Astruct__000000_A1.Anarrayofb = append(__Astruct__000000_A1.Anarrayofb, __Bstruct__000001_B2)
	__Astruct__000000_A1.Anarrayofb = append(__Astruct__000000_A1.Anarrayofb, __Bstruct__000000_B1)
	__Astruct__000000_A1.Anotherarrayofb = append(__Astruct__000000_A1.Anotherarrayofb, __Bstruct__000001_B2)
	__Astruct__000000_A1.Anotherarrayofb = append(__Astruct__000000_A1.Anotherarrayofb, __Bstruct__000002_B3)
	__Astruct__000000_A1.Anarrayofa = append(__Astruct__000000_A1.Anarrayofa, __Astruct__000000_A1)
	__Astruct__000000_A1.Anarrayofb2Use = append(__Astruct__000000_A1.Anarrayofb2Use, __AstructBstruct2Use__000000_)
	__Astruct__000001_A2.Anarrayofb2Use = append(__Astruct__000001_A2.Anarrayofb2Use, __AstructBstruct2Use__000001_)
	__AstructBstruct2Use__000000_.Bstrcut2 = __Bstruct__000000_B1
	__AstructBstruct2Use__000001_.Bstrcut2 = __Bstruct__000000_B1
}


