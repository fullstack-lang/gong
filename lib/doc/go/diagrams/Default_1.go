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
var _ ref_models.Stage

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage

	__Classdiagram__000000_Default_1 := (&models.Classdiagram{}).Stage(stage)

	__Field__000000_IsInDrawMode := (&models.Field{}).Stage(stage)

	__GongStructShape__000000_Default_1_GongStructShape := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_1_Classdiagram := (&models.GongStructShape{}).Stage(stage)

	__Link__000000_GongStructShapes := (&models.Link{}).Stage(stage)

	__Position__000000_Pos_Default_1_GongStructShape := (&models.Position{}).Stage(stage)
	__Position__000001_Pos_Default_1_Classdiagram := (&models.Position{}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_1_in_middle_between_Default_1_Classdiagram_and_Default_1_GongStructShape := (&models.Vertice{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default_1.Name = `Default_1`
	__Classdiagram__000000_Default_1.IsInDrawMode = false

	__Field__000000_IsInDrawMode.Name = `IsInDrawMode`

	__Field__000000_IsInDrawMode.Identifier = `ref_models.Classdiagram.IsInDrawMode`
	__Field__000000_IsInDrawMode.FieldTypeAsString = ``
	__Field__000000_IsInDrawMode.Structname = `Classdiagram`
	__Field__000000_IsInDrawMode.Fieldtypename = `bool`

	__GongStructShape__000000_Default_1_GongStructShape.Name = `Default_1-GongStructShape`

	__GongStructShape__000000_Default_1_GongStructShape.Identifier = `ref_models.GongStructShape`

	__GongStructShape__000000_Default_1_GongStructShape.Width = 240.000000
	__GongStructShape__000000_Default_1_GongStructShape.Height = 63.000000
	__GongStructShape__000000_Default_1_GongStructShape.IsSelected = false

	__GongStructShape__000001_Default_1_Classdiagram.Name = `Default_1-Classdiagram`

	__GongStructShape__000001_Default_1_Classdiagram.Identifier = `ref_models.Classdiagram`

	__GongStructShape__000001_Default_1_Classdiagram.Width = 240.000000
	__GongStructShape__000001_Default_1_Classdiagram.Height = 78.000000
	__GongStructShape__000001_Default_1_Classdiagram.IsSelected = false

	__Link__000000_GongStructShapes.Name = `GongStructShapes`

	__Link__000000_GongStructShapes.Identifier = `ref_models.Classdiagram.GongStructShapes`

	__Link__000000_GongStructShapes.Fieldtypename = `ref_models.GongStructShape`
	__Link__000000_GongStructShapes.FieldOffsetX = 0.000000
	__Link__000000_GongStructShapes.FieldOffsetY = 0.000000
	__Link__000000_GongStructShapes.TargetMultiplicity = models.MANY
	__Link__000000_GongStructShapes.TargetMultiplicityOffsetX = 0.000000
	__Link__000000_GongStructShapes.TargetMultiplicityOffsetY = 0.000000
	__Link__000000_GongStructShapes.SourceMultiplicity = models.MANY
	__Link__000000_GongStructShapes.SourceMultiplicityOffsetX = 0.000000
	__Link__000000_GongStructShapes.SourceMultiplicityOffsetY = 0.000000
	__Link__000000_GongStructShapes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_GongStructShapes.StartRatio = 0.500000
	__Link__000000_GongStructShapes.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_GongStructShapes.EndRatio = 0.621680
	__Link__000000_GongStructShapes.CornerOffsetRatio = -3.555556

	__Position__000000_Pos_Default_1_GongStructShape.X = 609.000000
	__Position__000000_Pos_Default_1_GongStructShape.Y = 323.999992
	__Position__000000_Pos_Default_1_GongStructShape.Name = `Pos-Default_1-GongStructShape`

	__Position__000001_Pos_Default_1_Classdiagram.X = 200.000000
	__Position__000001_Pos_Default_1_Classdiagram.Y = 143.000000
	__Position__000001_Pos_Default_1_Classdiagram.Name = `Pos-Default_1-Classdiagram`

	__Vertice__000000_Verticle_in_class_diagram_Default_1_in_middle_between_Default_1_Classdiagram_and_Default_1_GongStructShape.X = 596.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_1_in_middle_between_Default_1_Classdiagram_and_Default_1_GongStructShape.Y = 244.499996
	__Vertice__000000_Verticle_in_class_diagram_Default_1_in_middle_between_Default_1_Classdiagram_and_Default_1_GongStructShape.Name = `Verticle in class diagram Default_1 in middle between Default_1-Classdiagram and Default_1-GongStructShape`

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default_1.GongStructShapes = append(__Classdiagram__000000_Default_1.GongStructShapes, __GongStructShape__000000_Default_1_GongStructShape)
	__Classdiagram__000000_Default_1.GongStructShapes = append(__Classdiagram__000000_Default_1.GongStructShapes, __GongStructShape__000001_Default_1_Classdiagram)
	// setup of Field instances pointers
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_1_GongStructShape.Position = __Position__000000_Pos_Default_1_GongStructShape
	__GongStructShape__000001_Default_1_Classdiagram.Position = __Position__000001_Pos_Default_1_Classdiagram
	__GongStructShape__000001_Default_1_Classdiagram.Fields = append(__GongStructShape__000001_Default_1_Classdiagram.Fields, __Field__000000_IsInDrawMode)
	__GongStructShape__000001_Default_1_Classdiagram.Links = append(__GongStructShape__000001_Default_1_Classdiagram.Links, __Link__000000_GongStructShapes)
	// setup of Link instances pointers
	__Link__000000_GongStructShapes.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_1_in_middle_between_Default_1_Classdiagram_and_Default_1_GongStructShape
	// setup of Position instances pointers
	// setup of Vertice instances pointers
}
