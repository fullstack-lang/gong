package main

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

	__A__00000000_ := (&models.A{Name: `A0`}).Stage(stage)
	__A__00000000_.Name = `A0`
	__A__00000000_.B = nil
	// 2026-01-11T15:24:52.449614+01:00
	stage.Commit()
	__A__00000001_ := (&models.A{Name: `A1`}).Stage(stage)
	__A__00000001_.Name = `A1`
	__A__00000001_.B = nil
	// 2026-01-11T15:24:58.952296+01:00
	stage.Commit()
	__B__00000000_ := (&models.B{Name: `B0`}).Stage(stage)
	// A0
	__A__00000000_.Bs = slices.Insert(__A__00000000_.Bs, 0, __B__00000000_)
	__B__00000000_.Name = `B0`
	// 2026-01-11T15:25:19.952731+01:00
	stage.Commit()
	__B__00000001_ := (&models.B{Name: `B1`}).Stage(stage)
	__B__00000001_.Name = `B1`
	// 2026-01-11T15:25:32.582034+01:00
	stage.Commit()
	__B__00000002_ := (&models.B{Name: `B2`}).Stage(stage)
	__B__00000002_.Name = `B2`
	// 2026-01-11T15:26:08.581156+01:00
	stage.Commit()
	// A0
	__A__00000000_.Bs = slices.Insert(__A__00000000_.Bs, 1, __B__00000001_)
	__A__00000000_.Bs = slices.Insert(__A__00000000_.Bs, 2, __B__00000002_)
	// 2026-01-11T15:26:18.401153+01:00
	stage.Commit()
	// A1
	__A__00000001_.B = __B__00000002_
	__A__00000001_.Bs = slices.Insert(__A__00000001_.Bs, 0, __B__00000001_)
	__A__00000001_.Bs = slices.Insert(__A__00000001_.Bs, 1, __B__00000002_)
	// 2026-01-11T15:26:34.932119+01:00
	stage.Commit()
	// A1
	__A__00000001_.Bs = slices.Delete(__A__00000001_.Bs, 0, 1)
	__A__00000001_.Bs = slices.Insert(__A__00000001_.Bs, 1, __B__00000001_)
	// 2026-01-11T15:26:41.41876+01:00
	stage.Commit()
	// A0
	__A__00000000_.B = __B__00000000_
	__A__00000000_.Bs = slices.Delete(__A__00000000_.Bs, 0, 1)
	__A__00000000_.Bs = slices.Insert(__A__00000000_.Bs, 2, __B__00000000_)
	// 2026-01-11T15:27:18.750774+01:00
	stage.Commit()
}
