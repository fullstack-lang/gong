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
	// injection point for docLink to identifiers

	"ref_models.Classdiagram": &(ref_models.Classdiagram{}),

	"ref_models.Classdiagram.GongEnumShapes": (ref_models.Classdiagram{}).GongEnumShapes,

	"ref_models.Classdiagram.GongStructShapes": (ref_models.Classdiagram{}).GongStructShapes,

	"ref_models.Classdiagram.IsInDrawMode": (ref_models.Classdiagram{}).IsInDrawMode,

	"ref_models.Classdiagram.Name": (ref_models.Classdiagram{}).Name,

	"ref_models.Classdiagram.NoteShapes": (ref_models.Classdiagram{}).NoteShapes,

	"ref_models.DiagramPackage": &(ref_models.DiagramPackage{}),

	"ref_models.DiagramPackage.AbsolutePathToDiagramPackage": (ref_models.DiagramPackage{}).AbsolutePathToDiagramPackage,

	"ref_models.DiagramPackage.Classdiagrams": (ref_models.DiagramPackage{}).Classdiagrams,

	"ref_models.DiagramPackage.GongModelPath": (ref_models.DiagramPackage{}).GongModelPath,

	"ref_models.DiagramPackage.IsEditable": (ref_models.DiagramPackage{}).IsEditable,

	"ref_models.DiagramPackage.IsReloaded": (ref_models.DiagramPackage{}).IsReloaded,

	"ref_models.DiagramPackage.Name": (ref_models.DiagramPackage{}).Name,

	"ref_models.DiagramPackage.Path": (ref_models.DiagramPackage{}).Path,

	"ref_models.DiagramPackage.SelectedClassdiagram": (ref_models.DiagramPackage{}).SelectedClassdiagram,

	"ref_models.DiagramPackage.Umlscs": (ref_models.DiagramPackage{}).Umlscs,

	"ref_models.Field": &(ref_models.Field{}),

	"ref_models.Field.FieldTypeAsString": (ref_models.Field{}).FieldTypeAsString,

	"ref_models.Field.Fieldtypename": (ref_models.Field{}).Fieldtypename,

	"ref_models.Field.Identifier": (ref_models.Field{}).Identifier,

	"ref_models.Field.Name": (ref_models.Field{}).Name,

	"ref_models.Field.Structname": (ref_models.Field{}).Structname,

	"ref_models.GongEnumShape": &(ref_models.GongEnumShape{}),

	"ref_models.GongEnumShape.GongEnumValueEntrys": (ref_models.GongEnumShape{}).GongEnumValueEntrys,

	"ref_models.GongEnumShape.Height": (ref_models.GongEnumShape{}).Height,

	"ref_models.GongEnumShape.Identifier": (ref_models.GongEnumShape{}).Identifier,

	"ref_models.GongEnumShape.Name": (ref_models.GongEnumShape{}).Name,

	"ref_models.GongEnumShape.Position": (ref_models.GongEnumShape{}).Position,

	"ref_models.GongEnumShape.Width": (ref_models.GongEnumShape{}).Width,

	"ref_models.GongEnumShapeType": ref_models.GongEnumShapeType(0),

	"ref_models.GongEnumValueEntry": &(ref_models.GongEnumValueEntry{}),

	"ref_models.GongEnumValueEntry.Identifier": (ref_models.GongEnumValueEntry{}).Identifier,

	"ref_models.GongEnumValueEntry.Name": (ref_models.GongEnumValueEntry{}).Name,

	"ref_models.GongStructShape": &(ref_models.GongStructShape{}),

	"ref_models.GongStructShape.Fields": (ref_models.GongStructShape{}).Fields,

	"ref_models.GongStructShape.Height": (ref_models.GongStructShape{}).Height,

	"ref_models.GongStructShape.Identifier": (ref_models.GongStructShape{}).Identifier,

	"ref_models.GongStructShape.IsSelected": (ref_models.GongStructShape{}).IsSelected,

	"ref_models.GongStructShape.Links": (ref_models.GongStructShape{}).Links,

	"ref_models.GongStructShape.Name": (ref_models.GongStructShape{}).Name,

	"ref_models.GongStructShape.NbInstances": (ref_models.GongStructShape{}).NbInstances,

	"ref_models.GongStructShape.Position": (ref_models.GongStructShape{}).Position,

	"ref_models.GongStructShape.ShowNbInstances": (ref_models.GongStructShape{}).ShowNbInstances,

	"ref_models.GongStructShape.Width": (ref_models.GongStructShape{}).Width,

	"ref_models.GongdocStackName": ref_models.GongdocStackName,

	"ref_models.Int": ref_models.Int,

	"ref_models.Link": &(ref_models.Link{}),

	"ref_models.Link.CornerOffsetRatio": (ref_models.Link{}).CornerOffsetRatio,

	"ref_models.Link.EndOrientation": (ref_models.Link{}).EndOrientation,

	"ref_models.Link.EndRatio": (ref_models.Link{}).EndRatio,

	"ref_models.Link.FieldOffsetX": (ref_models.Link{}).FieldOffsetX,

	"ref_models.Link.FieldOffsetY": (ref_models.Link{}).FieldOffsetY,

	"ref_models.Link.Fieldtypename": (ref_models.Link{}).Fieldtypename,

	"ref_models.Link.Identifier": (ref_models.Link{}).Identifier,

	"ref_models.Link.Middlevertice": (ref_models.Link{}).Middlevertice,

	"ref_models.Link.Name": (ref_models.Link{}).Name,

	"ref_models.Link.SourceMultiplicity": (ref_models.Link{}).SourceMultiplicity,

	"ref_models.Link.SourceMultiplicityOffsetX": (ref_models.Link{}).SourceMultiplicityOffsetX,

	"ref_models.Link.SourceMultiplicityOffsetY": (ref_models.Link{}).SourceMultiplicityOffsetY,

	"ref_models.Link.StartOrientation": (ref_models.Link{}).StartOrientation,

	"ref_models.Link.StartRatio": (ref_models.Link{}).StartRatio,

	"ref_models.Link.TargetMultiplicity": (ref_models.Link{}).TargetMultiplicity,

	"ref_models.Link.TargetMultiplicityOffsetX": (ref_models.Link{}).TargetMultiplicityOffsetX,

	"ref_models.Link.TargetMultiplicityOffsetY": (ref_models.Link{}).TargetMultiplicityOffsetY,

	"ref_models.MANY": ref_models.MANY,

	"ref_models.MultiplicityType": ref_models.MultiplicityType(""),

	"ref_models.NOTE_SHAPE_LINK_TO_GONG_FIELD": ref_models.NOTE_SHAPE_LINK_TO_GONG_FIELD,

	"ref_models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE": ref_models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE,

	"ref_models.NoteOnGongdoc": ref_models.NoteOnGongdoc,

	"ref_models.NoteShape": &(ref_models.NoteShape{}),

	"ref_models.NoteShape.Body": (ref_models.NoteShape{}).Body,

	"ref_models.NoteShape.BodyHTML": (ref_models.NoteShape{}).BodyHTML,

	"ref_models.NoteShape.Height": (ref_models.NoteShape{}).Height,

	"ref_models.NoteShape.Identifier": (ref_models.NoteShape{}).Identifier,

	"ref_models.NoteShape.Matched": (ref_models.NoteShape{}).Matched,

	"ref_models.NoteShape.Name": (ref_models.NoteShape{}).Name,

	"ref_models.NoteShape.NoteShapeLinks": (ref_models.NoteShape{}).NoteShapeLinks,

	"ref_models.NoteShape.Width": (ref_models.NoteShape{}).Width,

	"ref_models.NoteShape.X": (ref_models.NoteShape{}).X,

	"ref_models.NoteShape.Y": (ref_models.NoteShape{}).Y,

	"ref_models.NoteShapeLink": &(ref_models.NoteShapeLink{}),

	"ref_models.NoteShapeLink.Identifier": (ref_models.NoteShapeLink{}).Identifier,

	"ref_models.NoteShapeLink.Name": (ref_models.NoteShapeLink{}).Name,

	"ref_models.NoteShapeLink.Type": (ref_models.NoteShapeLink{}).Type,

	"ref_models.NoteShapeLinkType": ref_models.NoteShapeLinkType(""),

	"ref_models.ONE": ref_models.ONE,

	"ref_models.ORIENTATION_HORIZONTAL": ref_models.ORIENTATION_HORIZONTAL,

	"ref_models.ORIENTATION_VERTICAL": ref_models.ORIENTATION_VERTICAL,

	"ref_models.OrientationType": ref_models.OrientationType(""),

	"ref_models.Position": &(ref_models.Position{}),

	"ref_models.Position.Name": (ref_models.Position{}).Name,

	"ref_models.Position.X": (ref_models.Position{}).X,

	"ref_models.Position.Y": (ref_models.Position{}).Y,

	"ref_models.StacksNames": ref_models.StacksNames(""),

	"ref_models.String": ref_models.String,

	"ref_models.SvgStackName": ref_models.SvgStackName,

	"ref_models.UmlState": &(ref_models.UmlState{}),

	"ref_models.UmlState.Name": (ref_models.UmlState{}).Name,

	"ref_models.UmlState.X": (ref_models.UmlState{}).X,

	"ref_models.UmlState.Y": (ref_models.UmlState{}).Y,

	"ref_models.Umlsc": &(ref_models.Umlsc{}),

	"ref_models.Umlsc.Activestate": (ref_models.Umlsc{}).Activestate,

	"ref_models.Umlsc.IsInDrawMode": (ref_models.Umlsc{}).IsInDrawMode,

	"ref_models.Umlsc.Name": (ref_models.Umlsc{}).Name,

	"ref_models.Umlsc.States": (ref_models.Umlsc{}).States,

	"ref_models.Vertice": &(ref_models.Vertice{}),

	"ref_models.Vertice.Name": (ref_models.Vertice{}).Name,

	"ref_models.Vertice.X": (ref_models.Vertice{}).X,

	"ref_models.Vertice.Y": (ref_models.Vertice{}).Y,

	"ref_models.ZERO_ONE": ref_models.ZERO_ONE,
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Classdiagram__000000_Default_2 := (&models.Classdiagram{Name: `Default_2`}).Stage(stage)

	__GongEnumShape__000000_Default_2_OrientationType := (&models.GongEnumShape{Name: `Default_2-OrientationType`}).Stage(stage)

	__GongStructShape__000000_Default_2_Classdiagram := (&models.GongStructShape{Name: `Default_2-Classdiagram`}).Stage(stage)
	__GongStructShape__000001_Default_2_Field := (&models.GongStructShape{Name: `Default_2-Field`}).Stage(stage)
	__GongStructShape__000002_Default_2_NoteShape := (&models.GongStructShape{Name: `Default_2-NoteShape`}).Stage(stage)

	__Link__000000_NoteShapes := (&models.Link{Name: `NoteShapes`}).Stage(stage)

	__NoteShape__000000_NoteOnGongdoc := (&models.NoteShape{Name: `NoteOnGongdoc`}).Stage(stage)

	__NoteShapeLink__000000_NoteShape := (&models.NoteShapeLink{Name: `NoteShape`}).Stage(stage)
	__NoteShapeLink__000001_NoteShapes := (&models.NoteShapeLink{Name: `NoteShapes`}).Stage(stage)
	__NoteShapeLink__000002_OrientationType := (&models.NoteShapeLink{Name: `OrientationType`}).Stage(stage)

	__Position__000000_Pos_Default_2_Classdiagram := (&models.Position{Name: `Pos-Default_2-Classdiagram`}).Stage(stage)
	__Position__000001_Pos_Default_2_Field := (&models.Position{Name: `Pos-Default_2-Field`}).Stage(stage)
	__Position__000002_Pos_Default_2_NoteShape := (&models.Position{Name: `Pos-Default_2-NoteShape`}).Stage(stage)
	__Position__000003_Pos_Default_2_OrientationType := (&models.Position{Name: `Pos-Default_2-OrientationType`}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_2_in_middle_between_Default_2_Classdiagram_and_Default_2_NoteShape := (&models.Vertice{Name: `Verticle in class diagram Default_2 in middle between Default_2-Classdiagram and Default_2-NoteShape`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default_2.Name = `Default_2`
	__Classdiagram__000000_Default_2.IsInDrawMode = false

	__GongEnumShape__000000_Default_2_OrientationType.Name = `Default_2-OrientationType`

	//gong:ident [ref_models.OrientationType] comment added to overcome the problem with the comment map association
	__GongEnumShape__000000_Default_2_OrientationType.Identifier = `ref_models.OrientationType`
	__GongEnumShape__000000_Default_2_OrientationType.Width = 240.000000
	__GongEnumShape__000000_Default_2_OrientationType.Height = 63.000000

	__GongStructShape__000000_Default_2_Classdiagram.Name = `Default_2-Classdiagram`

	//gong:ident [ref_models.Classdiagram] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_2_Classdiagram.Identifier = `ref_models.Classdiagram`
	__GongStructShape__000000_Default_2_Classdiagram.ShowNbInstances = false
	__GongStructShape__000000_Default_2_Classdiagram.NbInstances = 0
	__GongStructShape__000000_Default_2_Classdiagram.Width = 240.000000
	__GongStructShape__000000_Default_2_Classdiagram.Height = 63.000000
	__GongStructShape__000000_Default_2_Classdiagram.IsSelected = false

	__GongStructShape__000001_Default_2_Field.Name = `Default_2-Field`

	//gong:ident [ref_models.Field] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_2_Field.Identifier = `ref_models.Field`
	__GongStructShape__000001_Default_2_Field.ShowNbInstances = false
	__GongStructShape__000001_Default_2_Field.NbInstances = 0
	__GongStructShape__000001_Default_2_Field.Width = 240.000000
	__GongStructShape__000001_Default_2_Field.Height = 63.000000
	__GongStructShape__000001_Default_2_Field.IsSelected = false

	__GongStructShape__000002_Default_2_NoteShape.Name = `Default_2-NoteShape`

	//gong:ident [ref_models.NoteShape] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_2_NoteShape.Identifier = `ref_models.NoteShape`
	__GongStructShape__000002_Default_2_NoteShape.ShowNbInstances = false
	__GongStructShape__000002_Default_2_NoteShape.NbInstances = 0
	__GongStructShape__000002_Default_2_NoteShape.Width = 240.000000
	__GongStructShape__000002_Default_2_NoteShape.Height = 63.000000
	__GongStructShape__000002_Default_2_NoteShape.IsSelected = false

	__Link__000000_NoteShapes.Name = `NoteShapes`

	//gong:ident [ref_models.Classdiagram.NoteShapes] comment added to overcome the problem with the comment map association
	__Link__000000_NoteShapes.Identifier = `ref_models.Classdiagram.NoteShapes`

	//gong:ident [ref_models.NoteShape] comment added to overcome the problem with the comment map association
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

	//gong:ident [ref_models.NoteOnGongdoc] comment added to overcome the problem with the comment map association
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

	//gong:ident [ref_models.NoteShape] comment added to overcome the problem with the comment map association
	__NoteShapeLink__000000_NoteShape.Identifier = `ref_models.NoteShape`
	__NoteShapeLink__000000_NoteShape.Type = models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE

	__NoteShapeLink__000001_NoteShapes.Name = `NoteShapes`

	//gong:ident [ref_models.Classdiagram.NoteShapes] comment added to overcome the problem with the comment map association
	__NoteShapeLink__000001_NoteShapes.Identifier = `ref_models.Classdiagram.NoteShapes`
	__NoteShapeLink__000001_NoteShapes.Type = models.NOTE_SHAPE_LINK_TO_GONG_FIELD

	__NoteShapeLink__000002_OrientationType.Name = `OrientationType`

	//gong:ident [ref_models.OrientationType] comment added to overcome the problem with the comment map association
	__NoteShapeLink__000002_OrientationType.Identifier = `ref_models.OrientationType`
	__NoteShapeLink__000002_OrientationType.Type = models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE

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
	__Classdiagram__000000_Default_2.GongStructShapes = append(__Classdiagram__000000_Default_2.GongStructShapes, __GongStructShape__000000_Default_2_Classdiagram)
	__Classdiagram__000000_Default_2.GongStructShapes = append(__Classdiagram__000000_Default_2.GongStructShapes, __GongStructShape__000001_Default_2_Field)
	__Classdiagram__000000_Default_2.GongStructShapes = append(__Classdiagram__000000_Default_2.GongStructShapes, __GongStructShape__000002_Default_2_NoteShape)
	__Classdiagram__000000_Default_2.GongEnumShapes = append(__Classdiagram__000000_Default_2.GongEnumShapes, __GongEnumShape__000000_Default_2_OrientationType)
	__Classdiagram__000000_Default_2.NoteShapes = append(__Classdiagram__000000_Default_2.NoteShapes, __NoteShape__000000_NoteOnGongdoc)
	__GongEnumShape__000000_Default_2_OrientationType.Position = __Position__000003_Pos_Default_2_OrientationType
	__GongStructShape__000000_Default_2_Classdiagram.Position = __Position__000000_Pos_Default_2_Classdiagram
	__GongStructShape__000000_Default_2_Classdiagram.Links = append(__GongStructShape__000000_Default_2_Classdiagram.Links, __Link__000000_NoteShapes)
	__GongStructShape__000001_Default_2_Field.Position = __Position__000001_Pos_Default_2_Field
	__GongStructShape__000002_Default_2_NoteShape.Position = __Position__000002_Pos_Default_2_NoteShape
	__Link__000000_NoteShapes.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_2_in_middle_between_Default_2_Classdiagram_and_Default_2_NoteShape
	__NoteShape__000000_NoteOnGongdoc.NoteShapeLinks = append(__NoteShape__000000_NoteOnGongdoc.NoteShapeLinks, __NoteShapeLink__000000_NoteShape)
	__NoteShape__000000_NoteOnGongdoc.NoteShapeLinks = append(__NoteShape__000000_NoteOnGongdoc.NoteShapeLinks, __NoteShapeLink__000002_OrientationType)
	__NoteShape__000000_NoteOnGongdoc.NoteShapeLinks = append(__NoteShape__000000_NoteOnGongdoc.NoteShapeLinks, __NoteShapeLink__000001_NoteShapes)
}
