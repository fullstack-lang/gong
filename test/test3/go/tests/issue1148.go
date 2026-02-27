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

	__A__00000000_ := (&models.A{Name: `A0`}).Stage(stage)
	__B__00000000_ := (&models.B{Name: `B0`}).Stage(stage)

	__A__00000000_.Name = `A0`
	__B__00000000_.Name = `B0`

	__A__00000000_.Bs = append(__A__00000000_.Bs, __B__00000000_)

	stage.Commit()

	__A__00000000_.IntValue = 42

	stage.Commit()
}
