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
	__Astruct__000000__A1_ := (&models.Astruct{Name: "A1"}).Stage()

	// Declarations of staged instances of AstructBstruct2Use

	// Declarations of staged instances of AstructBstructUse

	// Declarations of staged instances of Bstruct

	// Declarations of staged instances of Dstruct

	// Setup of values

	// Astruct "A1" values setup
	__Astruct__000000__A1_.Name = `"A1"`
	__Astruct__000000__A1_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__000000__A1_.Booleanfield = false
	__Astruct__000000__A1_.CEnum = models.CENUM_VAL1
	__Astruct__000000__A1_.CName = ``
	__Astruct__000000__A1_.CFloatfield = 0.000000
	__Astruct__000000__A1_.Floatfield = 0.000000
	__Astruct__000000__A1_.Intfield = 0
	__Astruct__000000__A1_.Anotherbooleanfield = false
	__Astruct__000000__A1_.Duration1 = 0

	// Setup of pointers
}
