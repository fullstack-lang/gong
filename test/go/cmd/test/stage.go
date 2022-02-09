package main

import (
	"time"

	"github.com/fullstack-lang/gong/test/go/models"
)

var __Dummy_time_variable time.Time

func Unmarshall(stage *models.StageStruct) {

	// map of identifiers

	// Declarations of staged instances of Astruct
	__Astruct__000000_A1 := new(models.Astruct).Stage()
	__Astruct__000001_A2 := new(models.Astruct).Stage()

	// Declarations of staged instances of AstructBstruct2Use
	__AstructBstruct2Use__000000_ := new(models.AstructBstruct2Use).Stage()
	__AstructBstruct2Use__000001_ := new(models.AstructBstruct2Use).Stage()

	// Declarations of staged instances of AstructBstructUse

	// Declarations of staged instances of Bstruct
	__Bstruct__000000_B1 := new(models.Bstruct).Stage()
	__Bstruct__000001_B2 := new(models.Bstruct).Stage()
	__Bstruct__000002_B3 := new(models.Bstruct).Stage()

	// Declarations of staged instances of Dstruct

	// initializers of values

	// Init Astruct values A1
	__Astruct__000000_A1.Name = "A1"
	__Astruct__000000_A1.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-11-14 03:08:21 +0000 UTC")
	__Astruct__000000_A1.Booleanfield = false
	__Astruct__000000_A1.Aenum = "ENUM_VAL2"
	__Astruct__000000_A1.Aenum_2 = ""
	__Astruct__000000_A1.Benum = ""
	__Astruct__000000_A1.CName = ""
	__Astruct__000000_A1.CFloatfield = 56.444000
	__Astruct__000000_A1.Floatfield = 0.000000
	__Astruct__000000_A1.Intfield = 3
	__Astruct__000000_A1.Anotherbooleanfield = false
	__Astruct__000000_A1.Duration1 = 79653000000000

	// Init Astruct values A2
	__Astruct__000001_A2.Name = "A2"
	__Astruct__000001_A2.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-01-18 01:09:21 +0000 UTC")
	__Astruct__000001_A2.Booleanfield = false
	__Astruct__000001_A2.Aenum = "ENUM_VAL2"
	__Astruct__000001_A2.Aenum_2 = ""
	__Astruct__000001_A2.Benum = ""
	__Astruct__000001_A2.CName = ""
	__Astruct__000001_A2.CFloatfield = 0.100000
	__Astruct__000001_A2.Floatfield = 0.000000
	__Astruct__000001_A2.Intfield = 0
	__Astruct__000001_A2.Anotherbooleanfield = false
	__Astruct__000001_A2.Duration1 = 0

	// Init AstructBstruct2Use values 
	__AstructBstruct2Use__000000_.Name = ""

	// Init AstructBstruct2Use values 
	__AstructBstruct2Use__000001_.Name = ""

	// Init Bstruct values B1
	__Bstruct__000000_B1.Name = "B1"
	__Bstruct__000000_B1.Floatfield = 0.000000
	__Bstruct__000000_B1.Intfield = 0

	// Init Bstruct values B2
	__Bstruct__000001_B2.Name = "B2"
	__Bstruct__000001_B2.Floatfield = 0.000000
	__Bstruct__000001_B2.Intfield = 0

	// Init Bstruct values B3
	__Bstruct__000002_B3.Name = "B3"
	__Bstruct__000002_B3.Floatfield = 0.000000
	__Bstruct__000002_B3.Intfield = 0

	// initializers of pointers
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

	return
}


