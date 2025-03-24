package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/test/test/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration
var _ ref_models.StageStruct

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__GongStructShape__000000_Default_Astruct := (&models.GongStructShape{}).Stage(stage)

	__Position__000000_Pos_Default_Astruct := (&models.Position{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__GongStructShape__000000_Default_Astruct.Name = `Default-Astruct`

	//gong:ident [ref_models.Astruct] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Astruct.Identifier = `ref_models.Astruct`
	__GongStructShape__000000_Default_Astruct.ShowNbInstances = false
	__GongStructShape__000000_Default_Astruct.NbInstances = 0
	__GongStructShape__000000_Default_Astruct.Width = 240.000000
	__GongStructShape__000000_Default_Astruct.Height = 63.000000
	__GongStructShape__000000_Default_Astruct.IsSelected = false

	__Position__000000_Pos_Default_Astruct.X = 407.000000
	__Position__000000_Pos_Default_Astruct.Y = 118.000000
	__Position__000000_Pos_Default_Astruct.Name = `Pos-Default-Astruct`

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Astruct)
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Astruct.Position = __Position__000000_Pos_Default_Astruct
	// setup of Position instances pointers
}
