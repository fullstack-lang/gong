package main

import (
	"time"

	"github.com/fullstack-lang/gong/test/test/go/models"
	// injection point for ident package import declaration
	dummy "github.com/fullstack-lang/gong/test/test/go/dummy"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration
var _ dummy.Stage

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage

	__Astruct__000000_A1 := (&models.Astruct{}).Stage(stage)

	__Bstruct__000000_B1 := (&models.Bstruct{}).Stage(stage)

	// Setup of values

	__Astruct__000000_A1.Name = `A1`
	__Astruct__000000_A1.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__000000_A1.Date2, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Astruct__000000_A1.Booleanfield = false
	__Astruct__000000_A1.CEnum = models.CENUM_VAL1
	__Astruct__000000_A1.CName = ``
	__Astruct__000000_A1.CFloatfield = 0.000000
	__Astruct__000000_A1.Floatfield = 0.000000
	__Astruct__000000_A1.Intfield = 0
	__Astruct__000000_A1.Anotherbooleanfield = false
	__Astruct__000000_A1.Duration1 = 0
	__Astruct__000000_A1.TextFieldBespokeSize = ``
	__Astruct__000000_A1.TextArea = ``

	__Bstruct__000000_B1.Name = `B1`
	__Bstruct__000000_B1.Floatfield = 0.000000
	__Bstruct__000000_B1.Floatfield2 = 0.000000
	__Bstruct__000000_B1.Intfield = 0

	// Setup of pointers
	// setup of Astruct instances pointers
	__Astruct__000000_A1.Anarrayofb = append(__Astruct__000000_A1.Anarrayofb, __Bstruct__000000_B1)
	// setup of Bstruct instances pointers
}
