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

	__Classdiagram__000000_Default_3 := (&models.Classdiagram{Name: `Default_3`}).Stage(stage)

	__Field__000000_IsInDrawMode := (&models.Field{Name: `IsInDrawMode`}).Stage(stage)
	__Field__000001_Name := (&models.Field{Name: `Name`}).Stage(stage)

	__GongStructShape__000000_Default_3_Classdiagram := (&models.GongStructShape{Name: `Default_3-Classdiagram`}).Stage(stage)
	__GongStructShape__000001_Default_3_DiagramPackage := (&models.GongStructShape{Name: `Default_3-DiagramPackage`}).Stage(stage)

	__Link__000000_Classdiagrams := (&models.Link{Name: `Classdiagrams`}).Stage(stage)

	__Position__000000_Pos_Default_3_Classdiagram := (&models.Position{Name: `Pos-Default_3-Classdiagram`}).Stage(stage)
	__Position__000001_Pos_Default_3_DiagramPackage := (&models.Position{Name: `Pos-Default_3-DiagramPackage`}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_3_in_middle_between_Default_3_DiagramPackage_and_Default_3_Classdiagram := (&models.Vertice{Name: `Verticle in class diagram Default_3 in middle between Default_3-DiagramPackage and Default_3-Classdiagram`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default_3.Name = `Default_3`
	__Classdiagram__000000_Default_3.IsInDrawMode = false

	__Field__000000_IsInDrawMode.Name = `IsInDrawMode`

	//gong:ident [ref_models.Classdiagram.IsInDrawMode] comment added to overcome the problem with the comment map association
	__Field__000000_IsInDrawMode.Identifier = `ref_models.Classdiagram.IsInDrawMode`
	__Field__000000_IsInDrawMode.FieldTypeAsString = ``
	__Field__000000_IsInDrawMode.Structname = `Classdiagram`
	__Field__000000_IsInDrawMode.Fieldtypename = `bool`

	__Field__000001_Name.Name = `Name`

	//gong:ident [ref_models.Classdiagram.Name] comment added to overcome the problem with the comment map association
	__Field__000001_Name.Identifier = `ref_models.Classdiagram.Name`
	__Field__000001_Name.FieldTypeAsString = ``
	__Field__000001_Name.Structname = `Classdiagram`
	__Field__000001_Name.Fieldtypename = `string`

	__GongStructShape__000000_Default_3_Classdiagram.Name = `Default_3-Classdiagram`

	//gong:ident [ref_models.Classdiagram] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_3_Classdiagram.Identifier = `ref_models.Classdiagram`
	__GongStructShape__000000_Default_3_Classdiagram.ShowNbInstances = false
	__GongStructShape__000000_Default_3_Classdiagram.NbInstances = 0
	__GongStructShape__000000_Default_3_Classdiagram.Width = 240.000000
	__GongStructShape__000000_Default_3_Classdiagram.Height = 93.000000
	__GongStructShape__000000_Default_3_Classdiagram.IsSelected = false

	__GongStructShape__000001_Default_3_DiagramPackage.Name = `Default_3-DiagramPackage`

	//gong:ident [ref_models.DiagramPackage] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_3_DiagramPackage.Identifier = `ref_models.DiagramPackage`
	__GongStructShape__000001_Default_3_DiagramPackage.ShowNbInstances = false
	__GongStructShape__000001_Default_3_DiagramPackage.NbInstances = 0
	__GongStructShape__000001_Default_3_DiagramPackage.Width = 240.000000
	__GongStructShape__000001_Default_3_DiagramPackage.Height = 63.000000
	__GongStructShape__000001_Default_3_DiagramPackage.IsSelected = false

	__Link__000000_Classdiagrams.Name = `Classdiagrams`

	//gong:ident [ref_models.DiagramPackage.Classdiagrams] comment added to overcome the problem with the comment map association
	__Link__000000_Classdiagrams.Identifier = `ref_models.DiagramPackage.Classdiagrams`

	//gong:ident [ref_models.Classdiagram] comment added to overcome the problem with the comment map association
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
	__Classdiagram__000000_Default_3.GongStructShapes = append(__Classdiagram__000000_Default_3.GongStructShapes, __GongStructShape__000000_Default_3_Classdiagram)
	__Classdiagram__000000_Default_3.GongStructShapes = append(__Classdiagram__000000_Default_3.GongStructShapes, __GongStructShape__000001_Default_3_DiagramPackage)
	__GongStructShape__000000_Default_3_Classdiagram.Position = __Position__000000_Pos_Default_3_Classdiagram
	__GongStructShape__000000_Default_3_Classdiagram.Fields = append(__GongStructShape__000000_Default_3_Classdiagram.Fields, __Field__000001_Name)
	__GongStructShape__000000_Default_3_Classdiagram.Fields = append(__GongStructShape__000000_Default_3_Classdiagram.Fields, __Field__000000_IsInDrawMode)
	__GongStructShape__000001_Default_3_DiagramPackage.Position = __Position__000001_Pos_Default_3_DiagramPackage
	__GongStructShape__000001_Default_3_DiagramPackage.Links = append(__GongStructShape__000001_Default_3_DiagramPackage.Links, __Link__000000_Classdiagrams)
	__Link__000000_Classdiagrams.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_3_in_middle_between_Default_3_DiagramPackage_and_Default_3_Classdiagram
}
