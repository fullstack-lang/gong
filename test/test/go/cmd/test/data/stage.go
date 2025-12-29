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

	__Astruct__00000000_ := (&models.Astruct{Name: `A0xxdd`}).Stage(stage)
	__Astruct__00000001_ := (&models.Astruct{Name: `A1`}).Stage(stage)
	__Astruct__00000005_ := (&models.Astruct{Name: `A5`}).Stage(stage)
	__Astruct__00000006_ := (&models.Astruct{Name: `A6`}).Stage(stage)

	__AstructBstruct2Use__00000000_ := (&models.AstructBstruct2Use{Name: ``}).Stage(stage)
	__AstructBstruct2Use__00000001_ := (&models.AstructBstruct2Use{Name: ``}).Stage(stage)

	__Bstruct__00000000_ := (&models.Bstruct{Name: `B0`}).Stage(stage)
	__Bstruct__00000001_ := (&models.Bstruct{Name: `B1`}).Stage(stage)
	__Bstruct__00000003_ := (&models.Bstruct{Name: `B3`}).Stage(stage)
	__Bstruct__00000004_ := (&models.Bstruct{Name: `B4`}).Stage(stage)

	__Dstruct__00000000_ := (&models.Dstruct{Name: `D1`}).Stage(stage)

	__Gstruct__00000000_ := (&models.Gstruct{Name: ``}).Stage(stage)
	__Gstruct__00000001_ := (&models.Gstruct{Name: `g2`}).Stage(stage)

	// Setup of values

	__Astruct__00000000_.Name = `A0xxdd`
	__Astruct__00000000_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2042-11-14 03:08:21 +0000 UTC")
	__Astruct__00000000_.Date2, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__00000000_.Booleanfield = false
	__Astruct__00000000_.Aenum = models.ENUM_VAL2
	__Astruct__00000000_.Aenum_2 = models.ENUM_VAL2
	__Astruct__00000000_.Benum = models.BENUM_VAL2
	__Astruct__00000000_.CEnum = models.CENUM_VAL1
	__Astruct__00000000_.CName = `CName1

	\n"""" fdfsqjfhdqksfhqksf
Second Line`
	__Astruct__00000000_.CFloatfield = 60.500000
	__Astruct__00000000_.Floatfield = 0.000000
	__Astruct__00000000_.Intfield = 3
	__Astruct__00000000_.Anotherbooleanfield = false
	__Astruct__00000000_.Duration1 = -79653000000000
	__Astruct__00000000_.TextFieldBespokeSize = ``
	__Astruct__00000000_.TextArea = `dsqdqd
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

	__Astruct__00000001_.Name = `A1`
	__Astruct__00000001_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-01-28 09:59:06 +0000 UTC")
	__Astruct__00000001_.Date2, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__00000001_.Booleanfield = false
	__Astruct__00000001_.CEnum = models.CENUM_VAL1
	__Astruct__00000001_.CName = ``
	__Astruct__00000001_.CFloatfield = 0.000000
	__Astruct__00000001_.Floatfield = 0.000000
	__Astruct__00000001_.Intfield = 0
	__Astruct__00000001_.Anotherbooleanfield = false
	__Astruct__00000001_.Duration1 = 796530000000000000
	__Astruct__00000001_.TextFieldBespokeSize = ``
	__Astruct__00000001_.TextArea = ``

	__Astruct__00000005_.Name = `A5`
	__Astruct__00000005_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__00000005_.Date2, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__00000005_.Booleanfield = false
	__Astruct__00000005_.CEnum = models.CENUM_VAL1
	__Astruct__00000005_.CName = ``
	__Astruct__00000005_.CFloatfield = 0.000000
	__Astruct__00000005_.Floatfield = 0.000000
	__Astruct__00000005_.Intfield = 0
	__Astruct__00000005_.Anotherbooleanfield = false
	__Astruct__00000005_.Duration1 = 0
	__Astruct__00000005_.TextFieldBespokeSize = ``
	__Astruct__00000005_.TextArea = ``

	__Astruct__00000006_.Name = `A6`
	__Astruct__00000006_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__00000006_.Date2, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__00000006_.Booleanfield = false
	__Astruct__00000006_.CEnum = models.CENUM_VAL1
	__Astruct__00000006_.CName = ``
	__Astruct__00000006_.CFloatfield = 0.000000
	__Astruct__00000006_.Floatfield = 0.000000
	__Astruct__00000006_.Intfield = 0
	__Astruct__00000006_.Anotherbooleanfield = false
	__Astruct__00000006_.Duration1 = 0
	__Astruct__00000006_.TextFieldBespokeSize = ``
	__Astruct__00000006_.TextArea = ``

	__AstructBstruct2Use__00000000_.Name = ``

	__AstructBstruct2Use__00000001_.Name = ``

	__Bstruct__00000000_.Name = `B0`
	__Bstruct__00000000_.Floatfield = 0.000000
	__Bstruct__00000000_.Floatfield2 = 0.000000
	__Bstruct__00000000_.Intfield = 0

	__Bstruct__00000001_.Name = `B1`
	__Bstruct__00000001_.Floatfield = 0.000000
	__Bstruct__00000001_.Floatfield2 = 0.000000
	__Bstruct__00000001_.Intfield = 0

	__Bstruct__00000003_.Name = `B3`
	__Bstruct__00000003_.Floatfield = 0.000000
	__Bstruct__00000003_.Floatfield2 = 0.000000
	__Bstruct__00000003_.Intfield = 0

	__Bstruct__00000004_.Name = `B4`
	__Bstruct__00000004_.Floatfield = 0.000000
	__Bstruct__00000004_.Floatfield2 = 0.000000
	__Bstruct__00000004_.Intfield = 0

	__Dstruct__00000000_.Name = `D1`

	__Gstruct__00000000_.Name = ``
	__Gstruct__00000000_.Floatfield = 0.000000
	__Gstruct__00000000_.Floatfield2 = 0.000000
	__Gstruct__00000000_.Intfield = 0

	__Gstruct__00000001_.Name = `g2`
	__Gstruct__00000001_.Floatfield = 0.000000
	__Gstruct__00000001_.Floatfield2 = 0.000000
	__Gstruct__00000001_.Intfield = 0

	// Setup of pointers
	// setup of Astruct instances pointers
	__Astruct__00000000_.Anarrayofb = append(__Astruct__00000000_.Anarrayofb, __Bstruct__00000000_)
	__Astruct__00000000_.Anarrayofb2Use = append(__Astruct__00000000_.Anarrayofb2Use, __AstructBstruct2Use__00000001_)
	__Astruct__00000001_.Anarrayofb = append(__Astruct__00000001_.Anarrayofb, __Bstruct__00000000_)
	__Astruct__00000001_.Anarrayofb = append(__Astruct__00000001_.Anarrayofb, __Bstruct__00000004_)
	__Astruct__00000001_.Anarrayofb = append(__Astruct__00000001_.Anarrayofb, __Bstruct__00000003_)
	// setup of AstructBstruct2Use instances pointers
	__AstructBstruct2Use__00000000_.Bstrcut2 = __Bstruct__00000000_
	__AstructBstruct2Use__00000001_.Bstrcut2 = __Bstruct__00000000_
	// setup of Bstruct instances pointers
	// setup of Dstruct instances pointers
	__Dstruct__00000000_.Anarrayofb = append(__Dstruct__00000000_.Anarrayofb, __Bstruct__00000001_)
	// setup of Gstruct instances pointers
}

