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
	__A__00000001_ := (&models.A{Name: `A6`}).Stage(stage)
	__A__00000002_ := (&models.A{Name: `fff`}).Stage(stage)

	__B__00000000_ := (&models.B{Name: `1___0009`}).Stage(stage)
	__B__00000001_ := (&models.B{Name: `2gfsdfggf`}).Stage(stage)
	__B__00000002_ := (&models.B{Name: `6`}).Stage(stage)
	__B__00000003_ := (&models.B{Name: `7`}).Stage(stage)
	__B__00000004_ := (&models.B{Name: `8`}).Stage(stage)
	__B__00000005_ := (&models.B{Name: `9`}).Stage(stage)
	__B__00000006_ := (&models.B{Name: `10`}).Stage(stage)
	__B__00000007_ := (&models.B{Name: `11`}).Stage(stage)
	__B__00000008_ := (&models.B{Name: `12`}).Stage(stage)
	__B__00000009_ := (&models.B{Name: `13`}).Stage(stage)
	__B__00000010_ := (&models.B{Name: `14`}).Stage(stage)
	__B__00000011_ := (&models.B{Name: `15`}).Stage(stage)
	__B__00000012_ := (&models.B{Name: `16`}).Stage(stage)
	__B__00000013_ := (&models.B{Name: `17`}).Stage(stage)
	__B__00000014_ := (&models.B{Name: `20`}).Stage(stage)
	__B__00000015_ := (&models.B{Name: `21`}).Stage(stage)
	__B__00000016_ := (&models.B{Name: `22`}).Stage(stage)
	__B__00000017_ := (&models.B{Name: `24`}).Stage(stage)
	__B__00000018_ := (&models.B{Name: `25`}).Stage(stage)
	__B__00000019_ := (&models.B{Name: `26`}).Stage(stage)
	__B__00000020_ := (&models.B{Name: `27`}).Stage(stage)
	__B__00000021_ := (&models.B{Name: `29`}).Stage(stage)
	__B__00000022_ := (&models.B{Name: `30`}).Stage(stage)
	__B__00000023_ := (&models.B{Name: `31`}).Stage(stage)
	__B__00000024_ := (&models.B{Name: `32`}).Stage(stage)
	__B__00000025_ := (&models.B{Name: `33`}).Stage(stage)
	__B__00000026_ := (&models.B{Name: `34`}).Stage(stage)
	__B__00000027_ := (&models.B{Name: `B35`}).Stage(stage)

	// insertion point for initialization of values

	__A__00000000_.Name = `A09xxx`
	__A__00000000_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-01-01 00:00:00 +0000 UTC")
	__A__00000000_.FloatValue = 14.500000
	__A__00000000_.IntValue = 9
	__A__00000000_.Duration = -446400000000000
	__A__00000000_.EnumString = ""
	__A__00000000_.EnumInt = models.EnumTypeInt_Value2

	__A__00000001_.Name = `A6`
	__A__00000001_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__A__00000001_.FloatValue = 0.000000
	__A__00000001_.IntValue = 0
	__A__00000001_.Duration = 0
	__A__00000001_.EnumString = ""
	__A__00000001_.EnumInt = models.EnumTypeInt_Value2

	__A__00000002_.Name = `fff`
	__A__00000002_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__A__00000002_.FloatValue = 0.000000
	__A__00000002_.IntValue = 0
	__A__00000002_.Duration = 0
	__A__00000002_.EnumString = ""
	__A__00000002_.EnumInt = 0

	__B__00000000_.Name = `1___0009`

	__B__00000001_.Name = `2gfsdfggf`

	__B__00000002_.Name = `6`

	__B__00000003_.Name = `7`

	__B__00000004_.Name = `8`

	__B__00000005_.Name = `9`

	__B__00000006_.Name = `10`

	__B__00000007_.Name = `11`

	__B__00000008_.Name = `12`

	__B__00000009_.Name = `13`

	__B__00000010_.Name = `14`

	__B__00000011_.Name = `15`

	__B__00000012_.Name = `16`

	__B__00000013_.Name = `17`

	__B__00000014_.Name = `20`

	__B__00000015_.Name = `21`

	__B__00000016_.Name = `22`

	__B__00000017_.Name = `24`

	__B__00000018_.Name = `25`

	__B__00000019_.Name = `26`

	__B__00000020_.Name = `27`

	__B__00000021_.Name = `29`

	__B__00000022_.Name = `30`

	__B__00000023_.Name = `31`

	__B__00000024_.Name = `32`

	__B__00000025_.Name = `33`

	__B__00000026_.Name = `34`

	__B__00000027_.Name = `B35`

	// insertion point for setup of pointers
	__A__00000000_.B = __B__00000006_
	__A__00000001_.B = __B__00000000_
	__A__00000001_.Bs = append(__A__00000001_.Bs, __B__00000003_)
	__A__00000001_.Bs = append(__A__00000001_.Bs, __B__00000004_)
	__A__00000001_.Bs = append(__A__00000001_.Bs, __B__00000002_)
	__A__00000001_.Bs = append(__A__00000001_.Bs, __B__00000001_)
	__A__00000001_.Bs = append(__A__00000001_.Bs, __B__00000000_)
	__A__00000001_.Bs = append(__A__00000001_.Bs, __B__00000005_)
	__A__00000002_.B = __B__00000008_
	__A__00000002_.Bs = append(__A__00000002_.Bs, __B__00000002_)
	__A__00000002_.Bs = append(__A__00000002_.Bs, __B__00000004_)
}
