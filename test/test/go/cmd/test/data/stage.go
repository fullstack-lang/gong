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

	__Astruct__00000000_A0 := (&models.Astruct{}).Stage(stage)
	__Astruct__00000001_A1 := (&models.Astruct{}).Stage(stage)
	__Astruct__00000005_A5 := (&models.Astruct{}).Stage(stage)
	__Astruct__00000006_A6 := (&models.Astruct{}).Stage(stage)

	__AstructBstruct2Use__00000000_ := (&models.AstructBstruct2Use{}).Stage(stage)
	__AstructBstruct2Use__00000001_ := (&models.AstructBstruct2Use{}).Stage(stage)

	__Bstruct__00000000_B0 := (&models.Bstruct{}).Stage(stage)
	__Bstruct__00000001_B1 := (&models.Bstruct{}).Stage(stage)
	__Bstruct__00000003_B3 := (&models.Bstruct{}).Stage(stage)
	__Bstruct__00000004_B4 := (&models.Bstruct{}).Stage(stage)

	__Dstruct__00000000_D1 := (&models.Dstruct{}).Stage(stage)

	__Gstruct__00000000_ := (&models.Gstruct{}).Stage(stage)
	__Gstruct__00000001_g2 := (&models.Gstruct{}).Stage(stage)

	// Setup of values

	__Astruct__00000000_A0.Name = `A0`
	__Astruct__00000000_A0.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-11-14 03:08:21 +0000 UTC")
	__Astruct__00000000_A0.Date2, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__00000000_A0.Booleanfield = false
	__Astruct__00000000_A0.Aenum = models.ENUM_VAL1
	__Astruct__00000000_A0.Aenum_2 = models.ENUM_VAL2
	__Astruct__00000000_A0.Benum = models.BENUM_VAL2
	__Astruct__00000000_A0.CEnum = models.CENUM_VAL1
	__Astruct__00000000_A0.CName = `CName1

	\n"""" fdfsqjfhdqksfhqksf
Second Line`
	__Astruct__00000000_A0.CFloatfield = 60.500000
	__Astruct__00000000_A0.Floatfield = 0.000000
	__Astruct__00000000_A0.Intfield = 3
	__Astruct__00000000_A0.Anotherbooleanfield = false
	__Astruct__00000000_A0.Duration1 = -79653000000000
	__Astruct__00000000_A0.TextFieldBespokeSize = ``
	__Astruct__00000000_A0.TextArea = `dsqdqd
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

	__Astruct__00000001_A1.Name = `A1`
	__Astruct__00000001_A1.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-01-28 09:59:06 +0000 UTC")
	__Astruct__00000001_A1.Date2, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__00000001_A1.Booleanfield = false
	__Astruct__00000001_A1.CEnum = models.CENUM_VAL1
	__Astruct__00000001_A1.CName = ``
	__Astruct__00000001_A1.CFloatfield = 0.000000
	__Astruct__00000001_A1.Floatfield = 0.000000
	__Astruct__00000001_A1.Intfield = 0
	__Astruct__00000001_A1.Anotherbooleanfield = false
	__Astruct__00000001_A1.Duration1 = 796530000000000000
	__Astruct__00000001_A1.TextFieldBespokeSize = ``
	__Astruct__00000001_A1.TextArea = ``

	__Astruct__00000005_A5.Name = `A5`
	__Astruct__00000005_A5.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__00000005_A5.Date2, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__00000005_A5.Booleanfield = false
	__Astruct__00000005_A5.CEnum = models.CENUM_VAL1
	__Astruct__00000005_A5.CName = ``
	__Astruct__00000005_A5.CFloatfield = 0.000000
	__Astruct__00000005_A5.Floatfield = 0.000000
	__Astruct__00000005_A5.Intfield = 0
	__Astruct__00000005_A5.Anotherbooleanfield = false
	__Astruct__00000005_A5.Duration1 = 0
	__Astruct__00000005_A5.TextFieldBespokeSize = ``
	__Astruct__00000005_A5.TextArea = ``

	__Astruct__00000006_A6.Name = `A6`
	__Astruct__00000006_A6.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__00000006_A6.Date2, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__00000006_A6.Booleanfield = false
	__Astruct__00000006_A6.CEnum = models.CENUM_VAL1
	__Astruct__00000006_A6.CName = ``
	__Astruct__00000006_A6.CFloatfield = 0.000000
	__Astruct__00000006_A6.Floatfield = 0.000000
	__Astruct__00000006_A6.Intfield = 0
	__Astruct__00000006_A6.Anotherbooleanfield = false
	__Astruct__00000006_A6.Duration1 = 0
	__Astruct__00000006_A6.TextFieldBespokeSize = ``
	__Astruct__00000006_A6.TextArea = ``

	__AstructBstruct2Use__00000000_.Name = ``

	__AstructBstruct2Use__00000001_.Name = ``

	__Bstruct__00000000_B0.Name = `B0`
	__Bstruct__00000000_B0.Floatfield = 0.000000
	__Bstruct__00000000_B0.Floatfield2 = 0.000000
	__Bstruct__00000000_B0.Intfield = 0

	__Bstruct__00000001_B1.Name = `B1`
	__Bstruct__00000001_B1.Floatfield = 0.000000
	__Bstruct__00000001_B1.Floatfield2 = 0.000000
	__Bstruct__00000001_B1.Intfield = 0

	__Bstruct__00000003_B3.Name = `B3`
	__Bstruct__00000003_B3.Floatfield = 0.000000
	__Bstruct__00000003_B3.Floatfield2 = 0.000000
	__Bstruct__00000003_B3.Intfield = 0

	__Bstruct__00000004_B4.Name = `B4`
	__Bstruct__00000004_B4.Floatfield = 0.000000
	__Bstruct__00000004_B4.Floatfield2 = 0.000000
	__Bstruct__00000004_B4.Intfield = 0

	__Dstruct__00000000_D1.Name = `D1`

	__Gstruct__00000000_.Name = ``
	__Gstruct__00000000_.Floatfield = 0.000000
	__Gstruct__00000000_.Floatfield2 = 0.000000
	__Gstruct__00000000_.Intfield = 0

	__Gstruct__00000001_g2.Name = `g2`
	__Gstruct__00000001_g2.Floatfield = 0.000000
	__Gstruct__00000001_g2.Floatfield2 = 0.000000
	__Gstruct__00000001_g2.Intfield = 0

	// Setup of pointers
	// setup of Astruct instances pointers
	__Astruct__00000000_A0.Anarrayofb = append(__Astruct__00000000_A0.Anarrayofb, __Bstruct__00000000_B0)
	__Astruct__00000000_A0.Anarrayofb2Use = append(__Astruct__00000000_A0.Anarrayofb2Use, __AstructBstruct2Use__00000001_)
	__Astruct__00000001_A1.Anarrayofb = append(__Astruct__00000001_A1.Anarrayofb, __Bstruct__00000000_B0)
	__Astruct__00000001_A1.Anarrayofb = append(__Astruct__00000001_A1.Anarrayofb, __Bstruct__00000004_B4)
	__Astruct__00000001_A1.Anarrayofb = append(__Astruct__00000001_A1.Anarrayofb, __Bstruct__00000003_B3)
	// setup of AstructBstruct2Use instances pointers
	__AstructBstruct2Use__00000000_.Bstrcut2 = __Bstruct__00000000_B0
	__AstructBstruct2Use__00000001_.Bstrcut2 = __Bstruct__00000000_B0
	// setup of Bstruct instances pointers
	// setup of Dstruct instances pointers
	__Dstruct__00000000_D1.Anarrayofb = append(__Dstruct__00000000_D1.Anarrayofb, __Bstruct__00000001_B1)
	// setup of Gstruct instances pointers
}

