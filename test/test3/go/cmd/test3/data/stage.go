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

	__A__00000000_ := (&models.A{Name: `A09xxx`}).Stage(stage)
	__A__00000006_ := (&models.A{Name: `A6`}).Stage(stage)
	__A__00000007_ := (&models.A{Name: `fff`}).Stage(stage)

	__B__00000001_ := (&models.B{Name: `1___0009`}).Stage(stage)
	__B__00000002_ := (&models.B{Name: `2gfsdfggf`}).Stage(stage)
	__B__00000003_ := (&models.B{Name: `3`}).Stage(stage)
	__B__00000004_ := (&models.B{Name: `4`}).Stage(stage)
	__B__00000005_ := (&models.B{Name: `5`}).Stage(stage)
	__B__00000006_ := (&models.B{Name: `6`}).Stage(stage)
	__B__00000007_ := (&models.B{Name: `7`}).Stage(stage)
	__B__00000008_ := (&models.B{Name: `8`}).Stage(stage)
	__B__00000009_ := (&models.B{Name: `9`}).Stage(stage)
	__B__00000010_ := (&models.B{Name: `10`}).Stage(stage)
	__B__00000011_ := (&models.B{Name: `11`}).Stage(stage)
	__B__00000012_ := (&models.B{Name: `12`}).Stage(stage)
	__B__00000013_ := (&models.B{Name: `13`}).Stage(stage)
	__B__00000014_ := (&models.B{Name: `14`}).Stage(stage)
	__B__00000015_ := (&models.B{Name: `15`}).Stage(stage)
	__B__00000016_ := (&models.B{Name: `16`}).Stage(stage)
	__B__00000017_ := (&models.B{Name: `17`}).Stage(stage)
	__B__00000020_ := (&models.B{Name: `20`}).Stage(stage)
	__B__00000021_ := (&models.B{Name: `21`}).Stage(stage)
	__B__00000022_ := (&models.B{Name: `22`}).Stage(stage)
	__B__00000023_ := (&models.B{Name: `23`}).Stage(stage)
	__B__00000024_ := (&models.B{Name: `24`}).Stage(stage)
	__B__00000025_ := (&models.B{Name: `25`}).Stage(stage)
	__B__00000026_ := (&models.B{Name: `26`}).Stage(stage)
	__B__00000027_ := (&models.B{Name: `27`}).Stage(stage)
	__B__00000028_ := (&models.B{Name: `28`}).Stage(stage)
	__B__00000029_ := (&models.B{Name: `29`}).Stage(stage)
	__B__00000030_ := (&models.B{Name: `30`}).Stage(stage)
	__B__00000031_ := (&models.B{Name: `31`}).Stage(stage)
	__B__00000032_ := (&models.B{Name: `32`}).Stage(stage)
	__B__00000033_ := (&models.B{Name: `33`}).Stage(stage)
	__B__00000034_ := (&models.B{Name: `34`}).Stage(stage)
	__B__00000035_ := (&models.B{Name: `B35`}).Stage(stage)

	// insertion point for initialization of values

	__A__00000000_.Name = `A09xxx`

	__A__00000006_.Name = `A6`

	__A__00000007_.Name = `fff`

	__B__00000001_.Name = `1___0009`

	__B__00000002_.Name = `2gfsdfggf`

	__B__00000003_.Name = `3`

	__B__00000004_.Name = `4`

	__B__00000005_.Name = `5`

	__B__00000006_.Name = `6`

	__B__00000007_.Name = `7`

	__B__00000008_.Name = `8`

	__B__00000009_.Name = `9`

	__B__00000010_.Name = `10`

	__B__00000011_.Name = `11`

	__B__00000012_.Name = `12`

	__B__00000013_.Name = `13`

	__B__00000014_.Name = `14`

	__B__00000015_.Name = `15`

	__B__00000016_.Name = `16`

	__B__00000017_.Name = `17`

	__B__00000020_.Name = `20`

	__B__00000021_.Name = `21`

	__B__00000022_.Name = `22`

	__B__00000023_.Name = `23`

	__B__00000024_.Name = `24`

	__B__00000025_.Name = `25`

	__B__00000026_.Name = `26`

	__B__00000027_.Name = `27`

	__B__00000028_.Name = `28`

	__B__00000029_.Name = `29`

	__B__00000030_.Name = `30`

	__B__00000031_.Name = `31`

	__B__00000032_.Name = `32`

	__B__00000033_.Name = `33`

	__B__00000034_.Name = `34`

	__B__00000035_.Name = `B35`

	// insertion point for setup of pointers
	__A__00000000_.B = __B__00000011_
	__A__00000006_.B = __B__00000001_
	__A__00000006_.Bs = append(__A__00000006_.Bs, __B__00000007_)
	__A__00000006_.Bs = append(__A__00000006_.Bs, __B__00000008_)
	__A__00000006_.Bs = append(__A__00000006_.Bs, __B__00000005_)
	__A__00000007_.B = __B__00000012_
	__A__00000007_.Bs = append(__A__00000007_.Bs, __B__00000005_)
	__A__00000007_.Bs = append(__A__00000007_.Bs, __B__00000006_)
	__A__00000007_.Bs = append(__A__00000007_.Bs, __B__00000007_)
}
