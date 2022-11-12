package out

import (
	"time"

	"github.com/fullstack-lang/gong/test/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage models.StageStruct

func init() {
	var __Dummy_time_variable time.Time
	_ = __Dummy_time_variable
	InjectionGateway["out"] = outInjection
}

// outInjection will stage objects of database "out"
func outInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Astruct
	__Astruct__000000_A1 := (&models.Astruct{Name: `A1`}).Stage()

	// Declarations of staged instances of AstructBstruct2Use

	// Declarations of staged instances of AstructBstructUse

	// Declarations of staged instances of Bstruct
	__Bstruct__000000_B1 := (&models.Bstruct{Name: `B1`}).Stage()

	// Declarations of staged instances of Dstruct

	// Setup of values

	// Astruct A1 values setup
	__Astruct__000000_A1.Name = `A1`
	__Astruct__000000_A1.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-11-14 03:08:21 +0000 +0000")
	__Astruct__000000_A1.Booleanfield = true
	__Astruct__000000_A1.Aenum = models.
	__Astruct__000000_A1.Aenum_2 = models.ENUM_VAL2
	__Astruct__000000_A1.Benum = models.BENUM_VAL2
	__Astruct__000000_A1.CEnum = models.CENUM_VAL1
	__Astruct__000000_A1.CName = `CName1
	//
	// Second Line`
	__Astruct__000000_A1.CFloatfield = 60.500000
	__Astruct__000000_A1.Floatfield = 0.000000
	__Astruct__000000_A1.Intfield = 3
	__Astruct__000000_A1.Anotherbooleanfield = false
	__Astruct__000000_A1.Duration1 = 79653000000000

	// Bstruct B1 values setup
	__Bstruct__000000_B1.Name = `B1`
	__Bstruct__000000_B1.Floatfield = 0.000000
	__Bstruct__000000_B1.Floatfield2 = 0.000000
	__Bstruct__000000_B1.Intfield = 0

	// Setup of pointers
}


