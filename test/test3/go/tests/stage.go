package tests

import (
	"time"

	"github.com/fullstack-lang/gong/test/test3/go/models"
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

	// insertion point for declaration of instances to stage

	__A__00000000_ := (&models.A{Name: `initial A0`}).Stage(stage)
	__A__00000001_ := (&models.A{Name: `initial A1`}).Stage(stage)
	__A__00000002_ := (&models.A{Name: `initial A2`}).Stage(stage)

	__B__00000000_ := (&models.B{Name: `initial B0`}).Stage(stage)
	// insertion point for initialization of values

	__A__00000000_.Name = `initial A0`
	__A__00000000_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__A__00000000_.FloatValue = 14.500000
	__A__00000000_.IntValue = 9
	__A__00000000_.Duration = -446400000000000
	__A__00000000_.EnumString = models.EnumTypeString_Value1
	__A__00000000_.EnumInt = models.EnumTypeInt_Value2

	__A__00000001_.Name = `initial A1`
	__A__00000001_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__A__00000001_.FloatValue = 0.000000
	__A__00000001_.IntValue = 0
	__A__00000001_.Duration = 0
	__A__00000001_.EnumString = models.EnumTypeString_Value1
	__A__00000001_.EnumInt = models.EnumTypeInt_Value2

	__A__00000002_.Name = `initial A2`
	__A__00000002_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__A__00000002_.FloatValue = 0.000000
	__A__00000002_.IntValue = 0
	__A__00000002_.Duration = 0
	__A__00000002_.EnumString = models.EnumTypeString_Value1
	__A__00000002_.EnumInt = models.EnumTypeInt_Value2

	__B__00000000_.Name = `initial B0`

	__A__00000001_.B = __B__00000000_
	__A__00000001_.Bs = append(__A__00000001_.Bs, __B__00000000_)

}
