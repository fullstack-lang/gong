package main

import (
	"time"

	"github.com/fullstack-lang/gong/test/go/models"
)

func init() {
	var __Dummy_time_variable time.Time
	_ = __Dummy_time_variable
	InjectionGateway["stage"] = stageInjection
}

// stageInjection will stage objects of database "stage"
func stageInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Astruct
	__Astruct__000000_A1 := (&models.Astruct{Name: "A1"}).Stage()
	__Bstruct__000000_B1 := (&models.Bstruct{Name: "B1"}).Stage()

	__Astruct__000000_A1.Name = `A1`
	__Astruct__000000_A1.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-11-14 03:08:21 +0000 +0000")
	__Astruct__000000_A1.Booleanfield = false
	__Astruct__000000_A1.Aenum = models.ENUM_VAL1
	__Astruct__000000_A1.Aenum_2 = models.ENUM_VAL2
	__Astruct__000000_A1.Benum = models.BENUM_VAL2
	__Astruct__000000_A1.CEnum = models.CENUM_VAL2
	__Astruct__000000_A1.CName = `CName1
	//
	// Second Line`

	__Astruct__000000_A1.Anarrayofb = append(__Astruct__000000_A1.Anarrayofb, __Bstruct__000000_B1)

}
