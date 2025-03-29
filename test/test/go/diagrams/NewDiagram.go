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
var _ ref_models.Stage

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage

	__Classdiagram__000000_NewDiagram := (&models.Classdiagram{}).Stage(stage)

	__GongStructShape__000000_NewDiagram_Bstruct := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_NewDiagram_Astruct := (&models.GongStructShape{}).Stage(stage)

	__Link__000000_Associationtob := (&models.Link{}).Stage(stage)
	__Link__000001_Anarrayofb := (&models.Link{}).Stage(stage)

	__Position__000000_Pos_NewDiagram_Bstruct := (&models.Position{}).Stage(stage)
	__Position__000001_Pos_NewDiagram_Astruct := (&models.Position{}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Astruct_and_NewDiagram_Bstruct := (&models.Vertice{}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Astruct_and_NewDiagram_Bstruct := (&models.Vertice{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = false

	__GongStructShape__000000_NewDiagram_Bstruct.Name = `NewDiagram-Bstruct`

	//gong:ident [ref_models.Bstruct] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_NewDiagram_Bstruct.Identifier = `ref_models.Bstruct`
	__GongStructShape__000000_NewDiagram_Bstruct.ShowNbInstances = false
	__GongStructShape__000000_NewDiagram_Bstruct.NbInstances = 3
	__GongStructShape__000000_NewDiagram_Bstruct.Width = 240.000000
	__GongStructShape__000000_NewDiagram_Bstruct.Height = 63.000000
	__GongStructShape__000000_NewDiagram_Bstruct.IsSelected = false

	__GongStructShape__000001_NewDiagram_Astruct.Name = `NewDiagram-Astruct`

	//gong:ident [ref_models.Astruct] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_NewDiagram_Astruct.Identifier = `ref_models.Astruct`
	__GongStructShape__000001_NewDiagram_Astruct.ShowNbInstances = false
	__GongStructShape__000001_NewDiagram_Astruct.NbInstances = 0
	__GongStructShape__000001_NewDiagram_Astruct.Width = 240.000000
	__GongStructShape__000001_NewDiagram_Astruct.Height = 63.000000
	__GongStructShape__000001_NewDiagram_Astruct.IsSelected = false

	__Link__000000_Associationtob.Name = `Associationtob`

	//gong:ident [ref_models.Astruct.Associationtob] comment added to overcome the problem with the comment map association
	__Link__000000_Associationtob.Identifier = `ref_models.Astruct.Associationtob`

	//gong:ident [ref_models.Bstruct] comment added to overcome the problem with the comment map association
	__Link__000000_Associationtob.Fieldtypename = `ref_models.Bstruct`
	__Link__000000_Associationtob.FieldOffsetX = 0.000000
	__Link__000000_Associationtob.FieldOffsetY = 0.000000
	__Link__000000_Associationtob.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_Associationtob.TargetMultiplicityOffsetX = 0.000000
	__Link__000000_Associationtob.TargetMultiplicityOffsetY = 0.000000
	__Link__000000_Associationtob.SourceMultiplicity = models.MANY
	__Link__000000_Associationtob.SourceMultiplicityOffsetX = 0.000000
	__Link__000000_Associationtob.SourceMultiplicityOffsetY = 0.000000
	__Link__000000_Associationtob.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Associationtob.StartRatio = 0.500000
	__Link__000000_Associationtob.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_Associationtob.EndRatio = 0.821712
	__Link__000000_Associationtob.CornerOffsetRatio = 3.176711

	__Link__000001_Anarrayofb.Name = `Anarrayofb`

	//gong:ident [ref_models.Astruct.Anarrayofb] comment added to overcome the problem with the comment map association
	__Link__000001_Anarrayofb.Identifier = `ref_models.Astruct.Anarrayofb`

	//gong:ident [ref_models.Bstruct] comment added to overcome the problem with the comment map association
	__Link__000001_Anarrayofb.Fieldtypename = `ref_models.Bstruct`
	__Link__000001_Anarrayofb.FieldOffsetX = 0.000000
	__Link__000001_Anarrayofb.FieldOffsetY = 0.000000
	__Link__000001_Anarrayofb.TargetMultiplicity = models.MANY
	__Link__000001_Anarrayofb.TargetMultiplicityOffsetX = 0.000000
	__Link__000001_Anarrayofb.TargetMultiplicityOffsetY = 0.000000
	__Link__000001_Anarrayofb.SourceMultiplicity = models.MANY
	__Link__000001_Anarrayofb.SourceMultiplicityOffsetX = 0.000000
	__Link__000001_Anarrayofb.SourceMultiplicityOffsetY = 0.000000
	__Link__000001_Anarrayofb.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000001_Anarrayofb.StartRatio = 0.296712
	__Link__000001_Anarrayofb.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Anarrayofb.EndRatio = 0.500000
	__Link__000001_Anarrayofb.CornerOffsetRatio = 1.002108

	__Position__000000_Pos_NewDiagram_Bstruct.X = 485.000000
	__Position__000000_Pos_NewDiagram_Bstruct.Y = 292.000000
	__Position__000000_Pos_NewDiagram_Bstruct.Name = `Pos-NewDiagram-Bstruct`

	__Position__000001_Pos_NewDiagram_Astruct.X = 70.000000
	__Position__000001_Pos_NewDiagram_Astruct.Y = 86.000000
	__Position__000001_Pos_NewDiagram_Astruct.Name = `Pos-NewDiagram-Astruct`

	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Astruct_and_NewDiagram_Bstruct.X = 637.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Astruct_and_NewDiagram_Bstruct.Y = 220.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Astruct_and_NewDiagram_Bstruct.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Astruct and NewDiagram-Bstruct`

	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Astruct_and_NewDiagram_Bstruct.X = 637.500000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Astruct_and_NewDiagram_Bstruct.Y = 220.500000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Astruct_and_NewDiagram_Bstruct.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Astruct and NewDiagram-Bstruct`

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000000_NewDiagram_Bstruct)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000001_NewDiagram_Astruct)
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_NewDiagram_Bstruct.Position = __Position__000000_Pos_NewDiagram_Bstruct
	__GongStructShape__000001_NewDiagram_Astruct.Position = __Position__000001_Pos_NewDiagram_Astruct
	__GongStructShape__000001_NewDiagram_Astruct.Links = append(__GongStructShape__000001_NewDiagram_Astruct.Links, __Link__000000_Associationtob)
	__GongStructShape__000001_NewDiagram_Astruct.Links = append(__GongStructShape__000001_NewDiagram_Astruct.Links, __Link__000001_Anarrayofb)
	// setup of Link instances pointers
	__Link__000000_Associationtob.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Astruct_and_NewDiagram_Bstruct
	__Link__000001_Anarrayofb.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Astruct_and_NewDiagram_Bstruct
	// setup of Position instances pointers
	// setup of Vertice instances pointers
}
