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

	__Classdiagram__000000_Default_3 := (&models.Classdiagram{}).Stage(stage)

	__Field__000000_Name := (&models.Field{}).Stage(stage)
	__Field__000001_IsInDrawMode := (&models.Field{}).Stage(stage)

	__GongStructShape__000000_Default_3_Classdiagram := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_3_DiagramPackage := (&models.GongStructShape{}).Stage(stage)

	__Link__000000_Classdiagrams := (&models.Link{}).Stage(stage)

	__Position__000000_Pos_Default_3_Classdiagram := (&models.Position{}).Stage(stage)
	__Position__000001_Pos_Default_3_DiagramPackage := (&models.Position{}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_3_in_middle_between_Default_3_DiagramPackage_and_Default_3_Classdiagram := (&models.Vertice{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default_3.Name = `Default_3`
	__Classdiagram__000000_Default_3.IsInDrawMode = false

	__Field__000000_Name.Name = `Name`

	__Field__000000_Name.Identifier = `ref_models.Classdiagram.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `Classdiagram`
	__Field__000000_Name.Fieldtypename = `string`

	__Field__000001_IsInDrawMode.Name = `IsInDrawMode`

	__Field__000001_IsInDrawMode.Identifier = `ref_models.Classdiagram.IsInDrawMode`
	__Field__000001_IsInDrawMode.FieldTypeAsString = ``
	__Field__000001_IsInDrawMode.Structname = `Classdiagram`
	__Field__000001_IsInDrawMode.Fieldtypename = `bool`

	__GongStructShape__000000_Default_3_Classdiagram.Name = `Default_3-Classdiagram`

	__GongStructShape__000000_Default_3_Classdiagram.Identifier = `ref_models.Classdiagram`
	__GongStructShape__000000_Default_3_Classdiagram.ShowNbInstances = false
	__GongStructShape__000000_Default_3_Classdiagram.NbInstances = 0
	__GongStructShape__000000_Default_3_Classdiagram.Width = 240.000000
	__GongStructShape__000000_Default_3_Classdiagram.Height = 93.000000
	__GongStructShape__000000_Default_3_Classdiagram.IsSelected = false

	__GongStructShape__000001_Default_3_DiagramPackage.Name = `Default_3-DiagramPackage`

	__GongStructShape__000001_Default_3_DiagramPackage.Identifier = `ref_models.DiagramPackage`
	__GongStructShape__000001_Default_3_DiagramPackage.ShowNbInstances = false
	__GongStructShape__000001_Default_3_DiagramPackage.NbInstances = 0
	__GongStructShape__000001_Default_3_DiagramPackage.Width = 240.000000
	__GongStructShape__000001_Default_3_DiagramPackage.Height = 63.000000
	__GongStructShape__000001_Default_3_DiagramPackage.IsSelected = false

	__Link__000000_Classdiagrams.Name = `Classdiagrams`

	__Link__000000_Classdiagrams.Identifier = `ref_models.DiagramPackage.Classdiagrams`

	__Link__000000_Classdiagrams.Fieldtypename = `ref_models.Classdiagram`
	__Link__000000_Classdiagrams.FieldOffsetX = -50.000000
	__Link__000000_Classdiagrams.FieldOffsetY = -16.000000
	__Link__000000_Classdiagrams.TargetMultiplicity = models.MANY
	__Link__000000_Classdiagrams.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_Classdiagrams.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_Classdiagrams.SourceMultiplicity = models.MANY
	__Link__000000_Classdiagrams.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_Classdiagrams.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_Classdiagrams.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Classdiagrams.StartRatio = 0.500000
	__Link__000000_Classdiagrams.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Classdiagrams.EndRatio = 0.500000
	__Link__000000_Classdiagrams.CornerOffsetRatio = 1.380000

	__Position__000000_Pos_Default_3_Classdiagram.X = 78.000000
	__Position__000000_Pos_Default_3_Classdiagram.Y = 41.000000
	__Position__000000_Pos_Default_3_Classdiagram.Name = `Pos-Default_3-Classdiagram`

	__Position__000001_Pos_Default_3_DiagramPackage.X = 199.000000
	__Position__000001_Pos_Default_3_DiagramPackage.Y = 261.000000
	__Position__000001_Pos_Default_3_DiagramPackage.Name = `Pos-Default_3-DiagramPackage`

	__Vertice__000000_Verticle_in_class_diagram_Default_3_in_middle_between_Default_3_DiagramPackage_and_Default_3_Classdiagram.X = 451.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_3_in_middle_between_Default_3_DiagramPackage_and_Default_3_Classdiagram.Y = 98.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_3_in_middle_between_Default_3_DiagramPackage_and_Default_3_Classdiagram.Name = `Verticle in class diagram Default_3 in middle between Default_3-DiagramPackage and Default_3-Classdiagram`

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default_3.GongStructShapes = append(__Classdiagram__000000_Default_3.GongStructShapes, __GongStructShape__000000_Default_3_Classdiagram)
	__Classdiagram__000000_Default_3.GongStructShapes = append(__Classdiagram__000000_Default_3.GongStructShapes, __GongStructShape__000001_Default_3_DiagramPackage)
	// setup of Field instances pointers
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_3_Classdiagram.Position = __Position__000000_Pos_Default_3_Classdiagram
	__GongStructShape__000000_Default_3_Classdiagram.Fields = append(__GongStructShape__000000_Default_3_Classdiagram.Fields, __Field__000000_Name)
	__GongStructShape__000000_Default_3_Classdiagram.Fields = append(__GongStructShape__000000_Default_3_Classdiagram.Fields, __Field__000001_IsInDrawMode)
	__GongStructShape__000001_Default_3_DiagramPackage.Position = __Position__000001_Pos_Default_3_DiagramPackage
	__GongStructShape__000001_Default_3_DiagramPackage.Links = append(__GongStructShape__000001_Default_3_DiagramPackage.Links, __Link__000000_Classdiagrams)
	// setup of Link instances pointers
	__Link__000000_Classdiagrams.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_3_in_middle_between_Default_3_DiagramPackage_and_Default_3_Classdiagram
	// setup of Position instances pointers
	// setup of Vertice instances pointers
}
