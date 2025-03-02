package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/doc/go/models"
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

	__Classdiagram__000000_Default_4 := (&models.Classdiagram{Name: `Default_4`}).Stage(stage)

	__GongEnumShape__000000_Default_4_MultiplicityType := (&models.GongEnumShape{Name: `Default_4-MultiplicityType`}).Stage(stage)

	__Position__000000_Pos_Default_4_MultiplicityType := (&models.Position{Name: `Pos-Default_4-MultiplicityType`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default_4.Name = `Default_4`
	__Classdiagram__000000_Default_4.IsInDrawMode = false

	__GongEnumShape__000000_Default_4_MultiplicityType.Name = `Default_4-MultiplicityType`

	//gong:ident [ref_models.MultiplicityType] comment added to overcome the problem with the comment map association
	__GongEnumShape__000000_Default_4_MultiplicityType.Identifier = `ref_models.MultiplicityType`
	__GongEnumShape__000000_Default_4_MultiplicityType.Width = 240.000000
	__GongEnumShape__000000_Default_4_MultiplicityType.Height = 63.000000

	__Position__000000_Pos_Default_4_MultiplicityType.X = 56.000000
	__Position__000000_Pos_Default_4_MultiplicityType.Y = 74.000000
	__Position__000000_Pos_Default_4_MultiplicityType.Name = `Pos-Default_4-MultiplicityType`

	// Setup of pointers
	__Classdiagram__000000_Default_4.GongEnumShapes = append(__Classdiagram__000000_Default_4.GongEnumShapes, __GongEnumShape__000000_Default_4_MultiplicityType)
	__GongEnumShape__000000_Default_4_MultiplicityType.Position = __Position__000000_Pos_Default_4_MultiplicityType
}
