package main

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

	__A__00000000_ := (&models.A{Name: `A0`}).Stage(stage)
	__A__00000001_ := (&models.A{Name: `A1`}).Stage(stage)

	__B__00000000_ := (&models.B{Name: `B0`}).Stage(stage)
	__B__00000001_ := (&models.B{Name: `B1`}).Stage(stage)
	__B__00000002_ := (&models.B{Name: `B2`}).Stage(stage)

	// insertion point for initialization of values

	__A__00000000_.Name = `A0`

	__A__00000001_.Name = `A1`

	__B__00000000_.Name = `B0`

	__B__00000001_.Name = `B1`

	__B__00000002_.Name = `B2`

	// insertion point for setup of pointers
	__A__00000000_.B = __B__00000000_
	__A__00000000_.Bs = append(__A__00000000_.Bs, __B__00000001_)
	__A__00000000_.Bs = append(__A__00000000_.Bs, __B__00000002_)
	__A__00000000_.Bs = append(__A__00000000_.Bs, __B__00000000_)
	__A__00000001_.B = __B__00000002_
	__A__00000001_.Bs = append(__A__00000001_.Bs, __B__00000002_)
	__A__00000001_.Bs = append(__A__00000001_.Bs, __B__00000001_)
}
