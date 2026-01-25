package tests

import (
	"slices"
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

	// Forward commits:

	__A__00000000_ := (&models.A{Name: `A1`}).Stage(stage)
	__A__00000001_ := (&models.A{Name: `A0`}).Stage(stage)
	__B__00000000_ := (&models.B{Name: `B0`}).Stage(stage)
	__B__00000001_ := (&models.B{Name: `B1`}).Stage(stage)
	__A__00000000_.Name = `A1`
	__A__00000000_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__A__00000000_.FloatValue = 14.500000
	__A__00000000_.IntValue = 100
	__A__00000000_.Duration = -446400000
	__A__00000000_.EnumString = models.EnumTypeString_Value1
	__A__00000000_.EnumInt = models.EnumTypeInt_Value2
	__A__00000000_.B = nil
	__A__00000001_.Name = `A0`
	__A__00000001_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__A__00000001_.FloatValue = 14.500000
	__A__00000001_.IntValue = 100
	__A__00000001_.Duration = -446400000
	__A__00000001_.EnumString = models.EnumTypeString_Value1
	__A__00000001_.EnumInt = models.EnumTypeInt_Value2
	__A__00000001_.B = __B__00000000_
	__B__00000000_.Name = `B0`
	__B__00000001_.Name = `B1`
	stage.Commit()
	// A1
	__A__00000000_.IntValue = 150
	// A0
	__A__00000001_.IntValue = 200
	__A__00000001_.B = nil
	__A__00000001_.Bs = slices.Insert(__A__00000001_.Bs, 0, __B__00000000_)
	__A__00000001_.Bs = slices.Insert(__A__00000001_.Bs, 1, __B__00000001_)
	stage.Commit()
	// A0
	__A__00000001_.B = __B__00000001_
	stage.Commit()

}
