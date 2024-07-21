package main

import (
	"time"

	"github.com/fullstack-lang/gong/test/go/models"

	// injection point for ident package import declaration
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// Injection point for meta package dummy declaration
var _ gongdoc_models.StageStruct

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"gongdoc_models.GongStructShape": &(gongdoc_models.GongStructShape{}),

	"gongdoc_models.GongStructShape.Name": (gongdoc_models.GongStructShape{}).Name,
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Astruct__000000_A1 := (&models.Astruct{Name: `A1`}).Stage(stage)
	__Astruct__000001_A2 := (&models.Astruct{Name: `A2`}).Stage(stage)
	__Astruct__000002_A3 := (&models.Astruct{Name: `A3`}).Stage(stage)

	__AstructBstruct2Use__000000_ := (&models.AstructBstruct2Use{Name: ``}).Stage(stage)
	__AstructBstruct2Use__000001_ := (&models.AstructBstruct2Use{Name: ``}).Stage(stage)

	__Bstruct__000000_B1 := (&models.Bstruct{Name: `B1`}).Stage(stage)
	__Bstruct__000001_B2_ := (&models.Bstruct{Name: `B2 *`}).Stage(stage)
	__Bstruct__000002_B3 := (&models.Bstruct{Name: `B3`}).Stage(stage)

	__Dstruct__000000_D1 := (&models.Dstruct{Name: `D1`}).Stage(stage)

	// Setup of values

	__Astruct__000000_A1.Name = `A1`
	__Astruct__000000_A1.AnonymousStructField1.TheName1 = `A`
	__Astruct__000000_A1.AnonymousStructField2.TheName1 = ``
	__Astruct__000000_A1.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-11-14 03:08:21 +0000 UTC")
	__Astruct__000000_A1.Booleanfield = false
	__Astruct__000000_A1.Aenum = models.ENUM_VAL1
	__Astruct__000000_A1.Aenum_2 = models.ENUM_VAL2
	__Astruct__000000_A1.Benum = models.BENUM_VAL2
	__Astruct__000000_A1.CEnum = models.CENUM_VAL2
	__Astruct__000000_A1.CName = `CName1

	\n"""" fdfsqjfhdqksfhqksf
Second Line`
	__Astruct__000000_A1.CFloatfield = 60.500000
	__Astruct__000000_A1.Floatfield = 0.000000
	__Astruct__000000_A1.Intfield = 3
	__Astruct__000000_A1.Anotherbooleanfield = false
	__Astruct__000000_A1.Duration1 = -79653000000000

	//gong:ident [gongdoc_models.GongStructShape] comment added to overcome the problem with the comment map association
	__Astruct__000000_A1.StructRef = `gongdoc_models.GongStructShape`

	//gong:ident [gongdoc_models.GongStructShape.Name] comment added to overcome the problem with the comment map association
	__Astruct__000000_A1.FieldRef = `gongdoc_models.GongStructShape.Name`

	//gong:ident [........] comment added to overcome the problem with the comment map association
	__Astruct__000000_A1.EnumIntRef = `........`

	//gong:ident [........] comment added to overcome the problem with the comment map association
	__Astruct__000000_A1.EnumStringRef = `........`

	//gong:ident [........] comment added to overcome the problem with the comment map association
	__Astruct__000000_A1.EnumValue = `........`

	//gong:ident [........] comment added to overcome the problem with the comment map association
	__Astruct__000000_A1.ConstIdentifierValue = `........`
	__Astruct__000000_A1.TextFieldBespokeSize = ``
	__Astruct__000000_A1.TextArea = ``

	__Astruct__000001_A2.Name = `A2`
	__Astruct__000001_A2.AnonymousStructField1.TheName1 = ``
	__Astruct__000001_A2.AnonymousStructField2.TheName1 = ``
	__Astruct__000001_A2.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-01-28 09:59:06.744 +0000 UTC")
	__Astruct__000001_A2.Booleanfield = false
	__Astruct__000001_A2.CEnum = models.CENUM_VAL1
	__Astruct__000001_A2.CName = ``
	__Astruct__000001_A2.CFloatfield = 0.000000
	__Astruct__000001_A2.Floatfield = 0.000000
	__Astruct__000001_A2.Intfield = 0
	__Astruct__000001_A2.Anotherbooleanfield = false
	__Astruct__000001_A2.Duration1 = 796530000000000000

	//gong:ident [...] comment added to overcome the problem with the comment map association
	__Astruct__000001_A2.StructRef = `...`

	//gong:ident [...] comment added to overcome the problem with the comment map association
	__Astruct__000001_A2.FieldRef = `...`

	//gong:ident [...] comment added to overcome the problem with the comment map association
	__Astruct__000001_A2.EnumIntRef = `...`

	//gong:ident [...] comment added to overcome the problem with the comment map association
	__Astruct__000001_A2.EnumStringRef = `...`

	//gong:ident [...] comment added to overcome the problem with the comment map association
	__Astruct__000001_A2.EnumValue = `...`

	//gong:ident [...] comment added to overcome the problem with the comment map association
	__Astruct__000001_A2.ConstIdentifierValue = `...`
	__Astruct__000001_A2.TextFieldBespokeSize = ``
	__Astruct__000001_A2.TextArea = ``

	__Astruct__000002_A3.Name = `A3`
	__Astruct__000002_A3.AnonymousStructField1.TheName1 = ``
	__Astruct__000002_A3.AnonymousStructField2.TheName1 = ``
	__Astruct__000002_A3.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-01-28 11:16:35.978 +0000 UTC")
	__Astruct__000002_A3.Booleanfield = true
	__Astruct__000002_A3.Benum = models.BENUM_VAL2
	__Astruct__000002_A3.CEnum = models.CENUM_VAL1
	__Astruct__000002_A3.CName = ``
	__Astruct__000002_A3.CFloatfield = 4.900000
	__Astruct__000002_A3.Floatfield = 2.000000
	__Astruct__000002_A3.Intfield = 0
	__Astruct__000002_A3.Anotherbooleanfield = false
	__Astruct__000002_A3.Duration1 = -79653000000000

	//gong:ident [........] comment added to overcome the problem with the comment map association
	__Astruct__000002_A3.StructRef = `........`

	//gong:ident [........] comment added to overcome the problem with the comment map association
	__Astruct__000002_A3.FieldRef = `........`

	//gong:ident [........] comment added to overcome the problem with the comment map association
	__Astruct__000002_A3.EnumIntRef = `........`

	//gong:ident [........] comment added to overcome the problem with the comment map association
	__Astruct__000002_A3.EnumStringRef = `........`

	//gong:ident [........] comment added to overcome the problem with the comment map association
	__Astruct__000002_A3.EnumValue = `........`

	//gong:ident [........] comment added to overcome the problem with the comment map association
	__Astruct__000002_A3.ConstIdentifierValue = `........`
	__Astruct__000002_A3.TextFieldBespokeSize = ``
	__Astruct__000002_A3.TextArea = ``

	__AstructBstruct2Use__000000_.Name = ``

	__AstructBstruct2Use__000001_.Name = ``

	__Bstruct__000000_B1.Name = `B1`
	__Bstruct__000000_B1.Floatfield = 0.000000
	__Bstruct__000000_B1.Floatfield2 = 0.000000
	__Bstruct__000000_B1.Intfield = 0

	__Bstruct__000001_B2_.Name = `B2 *`
	__Bstruct__000001_B2_.Floatfield = 0.000000
	__Bstruct__000001_B2_.Floatfield2 = 0.000000
	__Bstruct__000001_B2_.Intfield = 0

	__Bstruct__000002_B3.Name = `B3`
	__Bstruct__000002_B3.Floatfield = 0.000000
	__Bstruct__000002_B3.Floatfield2 = 0.000000
	__Bstruct__000002_B3.Intfield = 0

	__Dstruct__000000_D1.Name = `D1`

	// Setup of pointers
	__Astruct__000000_A1.Associationtob = __Bstruct__000001_B2_
	__Astruct__000000_A1.Anarrayofb = append(__Astruct__000000_A1.Anarrayofb, __Bstruct__000002_B3)
	__Astruct__000000_A1.Anarrayofb = append(__Astruct__000000_A1.Anarrayofb, __Bstruct__000000_B1)
	__Astruct__000000_A1.Anarrayofa = append(__Astruct__000000_A1.Anarrayofa, __Astruct__000000_A1)
	__Astruct__000000_A1.Anotherarrayofb = append(__Astruct__000000_A1.Anotherarrayofb, __Bstruct__000001_B2_)
	__Astruct__000000_A1.Anotherarrayofb = append(__Astruct__000000_A1.Anotherarrayofb, __Bstruct__000002_B3)
	__Astruct__000000_A1.Anarrayofb2Use = append(__Astruct__000000_A1.Anarrayofb2Use, __AstructBstruct2Use__000001_)
	__AstructBstruct2Use__000000_.Bstrcut2 = __Bstruct__000000_B1
	__AstructBstruct2Use__000001_.Bstrcut2 = __Bstruct__000000_B1
	__Dstruct__000000_D1.Anarrayofb = append(__Dstruct__000000_D1.Anarrayofb, __Bstruct__000001_B2_)
	__Dstruct__000000_D1.Anarrayofb = append(__Dstruct__000000_D1.Anarrayofb, __Bstruct__000002_B3)
}
