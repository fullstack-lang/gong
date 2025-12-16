package main

import (
	"time"

	"github.com/fullstack-lang/gong/test/test/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage

	__Astruct__000000_A0 := (&models.Astruct{}).Stage(stage)
	__Astruct__000001_A1 := (&models.Astruct{}).Stage(stage)
	__Astruct__000002_A2 := (&models.Astruct{}).Stage(stage)
	__Astruct__000003_A4 := (&models.Astruct{}).Stage(stage)
	__Astruct__000004_A8 := (&models.Astruct{}).Stage(stage)

	__AstructBstruct2Use__000000_ := (&models.AstructBstruct2Use{}).Stage(stage)
	__AstructBstruct2Use__000001_ := (&models.AstructBstruct2Use{}).Stage(stage)

	__Bstruct__000000_B0 := (&models.Bstruct{}).Stage(stage)
	__Bstruct__000001_B1 := (&models.Bstruct{}).Stage(stage)
	__Bstruct__000002_B2 := (&models.Bstruct{}).Stage(stage)

	__Dstruct__000000_D1 := (&models.Dstruct{}).Stage(stage)

	__Gstruct__000000_ := (&models.Gstruct{}).Stage(stage)
	__Gstruct__000001_g2 := (&models.Gstruct{}).Stage(stage)

	// Setup of values

	__Astruct__000000_A0.Name = `A0`
	__Astruct__000000_A0.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-11-14 03:08:21 +0000 UTC")
	__Astruct__000000_A0.Date2, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__000000_A0.Booleanfield = false
	__Astruct__000000_A0.Aenum = models.ENUM_VAL1
	__Astruct__000000_A0.Aenum_2 = models.ENUM_VAL2
	__Astruct__000000_A0.Benum = models.BENUM_VAL2
	__Astruct__000000_A0.CEnum = models.CENUM_VAL1
	__Astruct__000000_A0.CName = `CName1

	\n"""" fdfsqjfhdqksfhqksf
Second Line`
	__Astruct__000000_A0.CFloatfield = 60.500000
	__Astruct__000000_A0.Floatfield = 0.000000
	__Astruct__000000_A0.Intfield = 3
	__Astruct__000000_A0.Anotherbooleanfield = false
	__Astruct__000000_A0.Duration1 = -79653000000000
	__Astruct__000000_A0.TextFieldBespokeSize = ``
	__Astruct__000000_A0.TextArea = `dsqdqd
dqsdqsd

dsq
dq
dsq
SQD
sqd
sqd
sqd
dsq
sqd
sqd
`

	__Astruct__000001_A1.Name = `A1`
	__Astruct__000001_A1.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-01-28 09:59:06 +0000 UTC")
	__Astruct__000001_A1.Date2, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__000001_A1.Booleanfield = false
	__Astruct__000001_A1.CEnum = models.CENUM_VAL1
	__Astruct__000001_A1.CName = ``
	__Astruct__000001_A1.CFloatfield = 0.000000
	__Astruct__000001_A1.Floatfield = 0.000000
	__Astruct__000001_A1.Intfield = 0
	__Astruct__000001_A1.Anotherbooleanfield = false
	__Astruct__000001_A1.Duration1 = 796530000000000000
	__Astruct__000001_A1.TextFieldBespokeSize = ``
	__Astruct__000001_A1.TextArea = ``

	__Astruct__000002_A2.Name = `A2`
	__Astruct__000002_A2.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-01-28 11:16:35 +0000 UTC")
	__Astruct__000002_A2.Date2, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__000002_A2.Booleanfield = true
	__Astruct__000002_A2.Benum = models.BENUM_VAL2
	__Astruct__000002_A2.CEnum = models.CENUM_VAL1
	__Astruct__000002_A2.CName = ``
	__Astruct__000002_A2.CFloatfield = 4.900000
	__Astruct__000002_A2.Floatfield = 2.000000
	__Astruct__000002_A2.Intfield = 0
	__Astruct__000002_A2.Anotherbooleanfield = false
	__Astruct__000002_A2.Duration1 = -79653000000000
	__Astruct__000002_A2.TextFieldBespokeSize = ``
	__Astruct__000002_A2.TextArea = ``

	__Astruct__000003_A4.Name = `A4`
	__Astruct__000003_A4.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__000003_A4.Date2, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__000003_A4.Booleanfield = false
	__Astruct__000003_A4.CEnum = models.CENUM_VAL1
	__Astruct__000003_A4.CName = ``
	__Astruct__000003_A4.CFloatfield = 0.000000
	__Astruct__000003_A4.Floatfield = 0.000000
	__Astruct__000003_A4.Intfield = 0
	__Astruct__000003_A4.Anotherbooleanfield = false
	__Astruct__000003_A4.Duration1 = 0
	__Astruct__000003_A4.TextFieldBespokeSize = ``
	__Astruct__000003_A4.TextArea = ``

	__Astruct__000004_A8.Name = `A8`
	__Astruct__000004_A8.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__000004_A8.Date2, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__000004_A8.Booleanfield = false
	__Astruct__000004_A8.CEnum = models.CENUM_VAL1
	__Astruct__000004_A8.CName = ``
	__Astruct__000004_A8.CFloatfield = 0.000000
	__Astruct__000004_A8.Floatfield = 0.000000
	__Astruct__000004_A8.Intfield = 0
	__Astruct__000004_A8.Anotherbooleanfield = false
	__Astruct__000004_A8.Duration1 = 0
	__Astruct__000004_A8.TextFieldBespokeSize = ``
	__Astruct__000004_A8.TextArea = ``

	__AstructBstruct2Use__000000_.Name = ``

	__AstructBstruct2Use__000001_.Name = ``

	__Bstruct__000000_B0.Name = `B0`
	__Bstruct__000000_B0.Floatfield = 0.000000
	__Bstruct__000000_B0.Floatfield2 = 0.000000
	__Bstruct__000000_B0.Intfield = 0

	__Bstruct__000001_B1.Name = `B1`
	__Bstruct__000001_B1.Floatfield = 0.000000
	__Bstruct__000001_B1.Floatfield2 = 0.000000
	__Bstruct__000001_B1.Intfield = 0

	__Bstruct__000002_B2.Name = `B2`
	__Bstruct__000002_B2.Floatfield = 0.000000
	__Bstruct__000002_B2.Floatfield2 = 0.000000
	__Bstruct__000002_B2.Intfield = 0

	__Dstruct__000000_D1.Name = `D1`

	__Gstruct__000000_.Name = ``
	__Gstruct__000000_.Floatfield = 0.000000
	__Gstruct__000000_.Floatfield2 = 0.000000
	__Gstruct__000000_.Intfield = 0

	__Gstruct__000001_g2.Name = `g2`
	__Gstruct__000001_g2.Floatfield = 0.000000
	__Gstruct__000001_g2.Floatfield2 = 0.000000
	__Gstruct__000001_g2.Intfield = 0

	// Setup of pointers
	// setup of Astruct instances pointers
	__Astruct__000000_A0.Anarrayofb = append(__Astruct__000000_A0.Anarrayofb, __Bstruct__000002_B2)
	__Astruct__000000_A0.Anarrayofb = append(__Astruct__000000_A0.Anarrayofb, __Bstruct__000001_B1)
	__Astruct__000000_A0.Anarrayofb = append(__Astruct__000000_A0.Anarrayofb, __Bstruct__000000_B0)
	__Astruct__000000_A0.Anarrayofb2Use = append(__Astruct__000000_A0.Anarrayofb2Use, __AstructBstruct2Use__000001_)
	// setup of AstructBstruct2Use instances pointers
	__AstructBstruct2Use__000000_.Bstrcut2 = __Bstruct__000000_B0
	__AstructBstruct2Use__000001_.Bstrcut2 = __Bstruct__000000_B0
	// setup of Bstruct instances pointers
	// setup of Dstruct instances pointers
	__Dstruct__000000_D1.Anarrayofb = append(__Dstruct__000000_D1.Anarrayofb, __Bstruct__000001_B1)
	__Dstruct__000000_D1.Anarrayofb = append(__Dstruct__000000_D1.Anarrayofb, __Bstruct__000002_B2)
	// setup of Gstruct instances pointers
}

