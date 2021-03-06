package main

import (
	"time"

	"github.com/fullstack-lang/gong/test/go/models"
)

func init() {
	var __Dummy_time_variable time.Time
	_ = __Dummy_time_variable
	InjectionGateway["alternate"] = alternateInjection
}

// alternateInjection will stage objects of database "alternate"
func alternateInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Astruct
	__Astruct__000000_A1 := (&models.Astruct{Name: "A1"}).Stage()
	__Astruct__000001_A2 := (&models.Astruct{Name: "A2"}).Stage()
	__Astruct__000002_A3 := (&models.Astruct{Name: "A3"}).Stage()

	// Declarations of staged instances of AstructBstruct2Use
	__AstructBstruct2Use__000000_ := (&models.AstructBstruct2Use{Name: ""}).Stage()
	__AstructBstruct2Use__000001_ := (&models.AstructBstruct2Use{Name: ""}).Stage()

	// Declarations of staged instances of AstructBstructUse

	// Declarations of staged instances of Bstruct
	__Bstruct__000000_B1 := (&models.Bstruct{Name: "B1"}).Stage()
	__Bstruct__000001_B2 := (&models.Bstruct{Name: "B2"}).Stage()
	__Bstruct__000002_B3 := (&models.Bstruct{Name: "B3"}).Stage()

	// Declarations of staged instances of Dstruct
	__Dstruct__000000_D1 := (&models.Dstruct{Name: "D1"}).Stage()

	// Setup of values

	// Astruct A1 values setup
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

	// Astruct A2 values setup
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

	// Astruct A3 values setup
	__Astruct__000002_A3.Name = "A3"
	__Astruct__000002_A3.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-02-10 01:06:11.446 +0000 UTC")
	__Astruct__000002_A3.Booleanfield = true
	__Astruct__000002_A3.Aenum = "ENUM_VAL2"
	__Astruct__000002_A3.Aenum_2 = "ENUM_VAL1"
	__Astruct__000002_A3.Benum = "BENUM_VAL2"
	__Astruct__000002_A3.CName = ""
	__Astruct__000002_A3.CFloatfield = 4.900000
	__Astruct__000002_A3.Floatfield = 0.000000
	__Astruct__000002_A3.Intfield = 0
	__Astruct__000002_A3.Anotherbooleanfield = false
	__Astruct__000002_A3.Duration1 = 0

	// AstructBstruct2Use  values setup
	__AstructBstruct2Use__000000_.Name = ""

	// AstructBstruct2Use  values setup
	__AstructBstruct2Use__000001_.Name = ""

	// Bstruct B1 values setup
	__Bstruct__000000_B1.Name = "B1"
	__Bstruct__000000_B1.Floatfield = 0.000000
	__Bstruct__000000_B1.Intfield = 0

	// Bstruct B2 values setup
	__Bstruct__000001_B2.Name = "B2"
	__Bstruct__000001_B2.Floatfield = 0.000000
	__Bstruct__000001_B2.Intfield = 0

	// Bstruct B3 values setup
	__Bstruct__000002_B3.Name = "B3"
	__Bstruct__000002_B3.Floatfield = 0.000000
	__Bstruct__000002_B3.Intfield = 0

	// Dstruct D1 values setup
	__Dstruct__000000_D1.Name = "D1"

	// Setup of pointers
	__Astruct__000000_A1.Associationtob = __Bstruct__000001_B2
	__Astruct__000000_A1.Anarrayofb = append(__Astruct__000000_A1.Anarrayofb, __Bstruct__000002_B3)
	__Astruct__000000_A1.Anarrayofb = append(__Astruct__000000_A1.Anarrayofb, __Bstruct__000001_B2)
	__Astruct__000000_A1.Anarrayofb = append(__Astruct__000000_A1.Anarrayofb, __Bstruct__000000_B1)
	__Astruct__000000_A1.Anotherarrayofb = append(__Astruct__000000_A1.Anotherarrayofb, __Bstruct__000001_B2)
	__Astruct__000000_A1.Anotherarrayofb = append(__Astruct__000000_A1.Anotherarrayofb, __Bstruct__000002_B3)
	__Astruct__000000_A1.Anarrayofa = append(__Astruct__000000_A1.Anarrayofa, __Astruct__000000_A1)
	__Astruct__000000_A1.Anarrayofb2Use = append(__Astruct__000000_A1.Anarrayofb2Use, __AstructBstruct2Use__000001_)
	__Astruct__000001_A2.Anarrayofb2Use = append(__Astruct__000001_A2.Anarrayofb2Use, __AstructBstruct2Use__000000_)
	__AstructBstruct2Use__000000_.Bstrcut2 = __Bstruct__000000_B1
	__AstructBstruct2Use__000001_.Bstrcut2 = __Bstruct__000000_B1
}

