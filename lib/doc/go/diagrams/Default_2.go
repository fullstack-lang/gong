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

	__Classdiagram__000000_Default_2 := (&models.Classdiagram{}).Stage(stage)

	__GongEnumShape__000000_Default_2_OrientationType := (&models.GongEnumShape{}).Stage(stage)

	__GongStructShape__000000_Default_2_Classdiagram := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_2_Field := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_2_NoteShape := (&models.GongStructShape{}).Stage(stage)

	__Link__000000_NoteShapes := (&models.Link{}).Stage(stage)

	__NoteShape__000000_NoteOnGongdoc := (&models.NoteShape{}).Stage(stage)

	__NoteShapeLink__000000_NoteShape := (&models.NoteShapeLink{}).Stage(stage)
	__NoteShapeLink__000001_OrientationType := (&models.NoteShapeLink{}).Stage(stage)
	__NoteShapeLink__000002_NoteShapes := (&models.NoteShapeLink{}).Stage(stage)

	__Position__000000_Pos_Default_2_Classdiagram := (&models.Position{}).Stage(stage)
	__Position__000001_Pos_Default_2_Field := (&models.Position{}).Stage(stage)
	__Position__000002_Pos_Default_2_NoteShape := (&models.Position{}).Stage(stage)
	__Position__000003_Pos_Default_2_OrientationType := (&models.Position{}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_2_in_middle_between_Default_2_Classdiagram_and_Default_2_NoteShape := (&models.Vertice{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default_2.Name = `Default_2`
	__Classdiagram__000000_Default_2.IsInDrawMode = false

	__GongEnumShape__000000_Default_2_OrientationType.Name = `Default_2-OrientationType`

	__GongEnumShape__000000_Default_2_OrientationType.Identifier = `ref_models.OrientationType`
	__GongEnumShape__000000_Default_2_OrientationType.Width = 240.000000
	__GongEnumShape__000000_Default_2_OrientationType.Height = 63.000000

	__GongStructShape__000000_Default_2_Classdiagram.Name = `Default_2-Classdiagram`

	__GongStructShape__000000_Default_2_Classdiagram.Identifier = `ref_models.Classdiagram`
	__GongStructShape__000000_Default_2_Classdiagram.ShowNbInstances = false
	__GongStructShape__000000_Default_2_Classdiagram.NbInstances = 0
	__GongStructShape__000000_Default_2_Classdiagram.Width = 240.000000
	__GongStructShape__000000_Default_2_Classdiagram.Height = 63.000000
	__GongStructShape__000000_Default_2_Classdiagram.IsSelected = false

	__GongStructShape__000001_Default_2_Field.Name = `Default_2-Field`

	__GongStructShape__000001_Default_2_Field.Identifier = `ref_models.Field`
	__GongStructShape__000001_Default_2_Field.ShowNbInstances = false
	__GongStructShape__000001_Default_2_Field.NbInstances = 0
	__GongStructShape__000001_Default_2_Field.Width = 240.000000
	__GongStructShape__000001_Default_2_Field.Height = 63.000000
	__GongStructShape__000001_Default_2_Field.IsSelected = false

	__GongStructShape__000002_Default_2_NoteShape.Name = `Default_2-NoteShape`

	__GongStructShape__000002_Default_2_NoteShape.Identifier = `ref_models.NoteShape`
	__GongStructShape__000002_Default_2_NoteShape.ShowNbInstances = false
	__GongStructShape__000002_Default_2_NoteShape.NbInstances = 0
	__GongStructShape__000002_Default_2_NoteShape.Width = 240.000000
	__GongStructShape__000002_Default_2_NoteShape.Height = 63.000000
	__GongStructShape__000002_Default_2_NoteShape.IsSelected = false

	__Link__000000_NoteShapes.Name = `NoteShapes`

	__Link__000000_NoteShapes.Identifier = `ref_models.Classdiagram.NoteShapes`

	__Link__000000_NoteShapes.Fieldtypename = `ref_models.NoteShape`
	__Link__000000_NoteShapes.FieldOffsetX = -50.000000
	__Link__000000_NoteShapes.FieldOffsetY = -16.000000
	__Link__000000_NoteShapes.TargetMultiplicity = models.MANY
	__Link__000000_NoteShapes.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_NoteShapes.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_NoteShapes.SourceMultiplicity = models.MANY
	__Link__000000_NoteShapes.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_NoteShapes.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_NoteShapes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_NoteShapes.StartRatio = 0.500000
	__Link__000000_NoteShapes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_NoteShapes.EndRatio = 0.500000
	__Link__000000_NoteShapes.CornerOffsetRatio = 1.380000

	__NoteShape__000000_NoteOnGongdoc.Name = `NoteOnGongdoc`

	__NoteShape__000000_NoteOnGongdoc.Identifier = `ref_models.NoteOnGongdoc`
	__NoteShape__000000_NoteOnGongdoc.Body = `Note Example

This note can refers to [models.NoteShape]
or to [models.Classdiagram.NoteShapes]
or to [models.OrientationType]
`
	__NoteShape__000000_NoteOnGongdoc.BodyHTML = `<p>Note Example
<p>This note can refers to <a href="/models#NoteShape">models.NoteShape</a>
or to <a href="/models#Classdiagram.NoteShapes">models.Classdiagram.NoteShapes</a>
or to <a href="/models#OrientationType">models.OrientationType</a>
`
	__NoteShape__000000_NoteOnGongdoc.X = 478.000000
	__NoteShape__000000_NoteOnGongdoc.Y = 185.000000
	__NoteShape__000000_NoteOnGongdoc.Width = 357.000000
	__NoteShape__000000_NoteOnGongdoc.Height = 106.000000
	__NoteShape__000000_NoteOnGongdoc.Matched = false

	__NoteShapeLink__000000_NoteShape.Name = `NoteShape`

	__NoteShapeLink__000000_NoteShape.Identifier = `ref_models.NoteShape`
	__NoteShapeLink__000000_NoteShape.Type = models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE

	__NoteShapeLink__000001_OrientationType.Name = `OrientationType`

	__NoteShapeLink__000001_OrientationType.Identifier = `ref_models.OrientationType`
	__NoteShapeLink__000001_OrientationType.Type = models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE

	__NoteShapeLink__000002_NoteShapes.Name = `NoteShapes`

	__NoteShapeLink__000002_NoteShapes.Identifier = `ref_models.Classdiagram.NoteShapes`
	__NoteShapeLink__000002_NoteShapes.Type = models.NOTE_SHAPE_LINK_TO_GONG_FIELD

	__Position__000000_Pos_Default_2_Classdiagram.X = 78.000000
	__Position__000000_Pos_Default_2_Classdiagram.Y = 27.000000
	__Position__000000_Pos_Default_2_Classdiagram.Name = `Pos-Default_2-Classdiagram`

	__Position__000001_Pos_Default_2_Field.X = 85.000000
	__Position__000001_Pos_Default_2_Field.Y = 182.000000
	__Position__000001_Pos_Default_2_Field.Name = `Pos-Default_2-Field`

	__Position__000002_Pos_Default_2_NoteShape.X = 565.000000
	__Position__000002_Pos_Default_2_NoteShape.Y = 21.000000
	__Position__000002_Pos_Default_2_NoteShape.Name = `Pos-Default_2-NoteShape`

	__Position__000003_Pos_Default_2_OrientationType.X = 92.000000
	__Position__000003_Pos_Default_2_OrientationType.Y = 368.000000
	__Position__000003_Pos_Default_2_OrientationType.Name = `Pos-Default_2-OrientationType`

	__Vertice__000000_Verticle_in_class_diagram_Default_2_in_middle_between_Default_2_Classdiagram_and_Default_2_NoteShape.X = 681.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_2_in_middle_between_Default_2_Classdiagram_and_Default_2_NoteShape.Y = 55.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_2_in_middle_between_Default_2_Classdiagram_and_Default_2_NoteShape.Name = `Verticle in class diagram Default_2 in middle between Default_2-Classdiagram and Default_2-NoteShape`

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default_2.GongStructShapes = append(__Classdiagram__000000_Default_2.GongStructShapes, __GongStructShape__000000_Default_2_Classdiagram)
	__Classdiagram__000000_Default_2.GongStructShapes = append(__Classdiagram__000000_Default_2.GongStructShapes, __GongStructShape__000001_Default_2_Field)
	__Classdiagram__000000_Default_2.GongStructShapes = append(__Classdiagram__000000_Default_2.GongStructShapes, __GongStructShape__000002_Default_2_NoteShape)
	__Classdiagram__000000_Default_2.GongEnumShapes = append(__Classdiagram__000000_Default_2.GongEnumShapes, __GongEnumShape__000000_Default_2_OrientationType)
	__Classdiagram__000000_Default_2.NoteShapes = append(__Classdiagram__000000_Default_2.NoteShapes, __NoteShape__000000_NoteOnGongdoc)
	// setup of GongEnumShape instances pointers
	__GongEnumShape__000000_Default_2_OrientationType.Position = __Position__000003_Pos_Default_2_OrientationType
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_2_Classdiagram.Position = __Position__000000_Pos_Default_2_Classdiagram
	__GongStructShape__000000_Default_2_Classdiagram.Links = append(__GongStructShape__000000_Default_2_Classdiagram.Links, __Link__000000_NoteShapes)
	__GongStructShape__000001_Default_2_Field.Position = __Position__000001_Pos_Default_2_Field
	__GongStructShape__000002_Default_2_NoteShape.Position = __Position__000002_Pos_Default_2_NoteShape
	// setup of Link instances pointers
	__Link__000000_NoteShapes.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_2_in_middle_between_Default_2_Classdiagram_and_Default_2_NoteShape
	// setup of NoteShape instances pointers
	__NoteShape__000000_NoteOnGongdoc.NoteShapeLinks = append(__NoteShape__000000_NoteOnGongdoc.NoteShapeLinks, __NoteShapeLink__000000_NoteShape)
	__NoteShape__000000_NoteOnGongdoc.NoteShapeLinks = append(__NoteShape__000000_NoteOnGongdoc.NoteShapeLinks, __NoteShapeLink__000001_OrientationType)
	__NoteShape__000000_NoteOnGongdoc.NoteShapeLinks = append(__NoteShape__000000_NoteOnGongdoc.NoteShapeLinks, __NoteShapeLink__000002_NoteShapes)
	// setup of NoteShapeLink instances pointers
	// setup of Position instances pointers
	// setup of Vertice instances pointers
}
