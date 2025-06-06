package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
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

	// Declaration of instances to stage

	__AttributeShape__000000_FieldTypeAsString := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000001_Fieldtypename := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000002_TargetMultiplicityOffsetX := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000003_TargetMultiplicityOffsetY := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)
	__Classdiagram__000001_spoil_diagram := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z := (&models.DiagramPackage{}).Stage(stage)

	__GongNoteShape__000000_Default_NoteOnGongdoc := (&models.GongNoteShape{}).Stage(stage)

	__GongStructShape__000000_Default_AttributeShape := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_DiagramPackage := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_GongStructShape := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000003_Default_Classdiagram := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000004_Default_1_AttributeShape := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000005_Default_1_Classdiagram := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000006_Default_1_DiagramPackage := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000007_Default_1_GongEnumShape := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000008_Default_1_GongEnumValueShape := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000009_Default_1_GongNoteLinkShape := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000010_Default_1_GongNoteShape := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000011_Default_1_GongStructShape := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000012_Default_1_LinkShape := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_Classdiagrams := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000001_GongStructShapes := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000002_AttributeShapes := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000003_GongStructShapes := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000004_GongEnumShapes := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_FieldTypeAsString.Name = `FieldTypeAsString`

	__AttributeShape__000000_FieldTypeAsString.Identifier = `ref_models.AttributeShape.FieldTypeAsString`
	__AttributeShape__000000_FieldTypeAsString.FieldTypeAsString = ``
	__AttributeShape__000000_FieldTypeAsString.Structname = `AttributeShape`
	__AttributeShape__000000_FieldTypeAsString.Fieldtypename = `string`

	__AttributeShape__000001_Fieldtypename.Name = `Fieldtypename`

	__AttributeShape__000001_Fieldtypename.Identifier = `ref_models.AttributeShape.Fieldtypename`
	__AttributeShape__000001_Fieldtypename.FieldTypeAsString = ``
	__AttributeShape__000001_Fieldtypename.Structname = `AttributeShape`
	__AttributeShape__000001_Fieldtypename.Fieldtypename = `string`

	__AttributeShape__000002_TargetMultiplicityOffsetX.Name = `TargetMultiplicityOffsetX`

	__AttributeShape__000002_TargetMultiplicityOffsetX.Identifier = `ref_models.LinkShape.TargetMultiplicityOffsetX`
	__AttributeShape__000002_TargetMultiplicityOffsetX.FieldTypeAsString = ``
	__AttributeShape__000002_TargetMultiplicityOffsetX.Structname = `LinkShape`
	__AttributeShape__000002_TargetMultiplicityOffsetX.Fieldtypename = `float64`

	__AttributeShape__000003_TargetMultiplicityOffsetY.Name = `TargetMultiplicityOffsetY`

	__AttributeShape__000003_TargetMultiplicityOffsetY.Identifier = `ref_models.LinkShape.TargetMultiplicityOffsetY`
	__AttributeShape__000003_TargetMultiplicityOffsetY.FieldTypeAsString = ``
	__AttributeShape__000003_TargetMultiplicityOffsetY.Structname = `LinkShape`
	__AttributeShape__000003_TargetMultiplicityOffsetY.Fieldtypename = `float64`

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = `This diagram describes the model of the doc2 package. A DiagramPackage is composed of ClassDiagram and each ClassDiagram has shapes.`
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = true
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = `[false,true,false,true,false,false,false,true]`
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = true
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = `[true]`

	__Classdiagram__000001_spoil_diagram.Name = `spoil diagram`
	__Classdiagram__000001_spoil_diagram.Description = `Spoil diagram`
	__Classdiagram__000001_spoil_diagram.IsIncludedInStaticWebSite = false
	__Classdiagram__000001_spoil_diagram.IsInRenameMode = false
	__Classdiagram__000001_spoil_diagram.IsExpanded = true
	__Classdiagram__000001_spoil_diagram.NodeGongStructsIsExpanded = true
	__Classdiagram__000001_spoil_diagram.NodeGongStructNodeExpansion = ``
	__Classdiagram__000001_spoil_diagram.NodeGongEnumsIsExpanded = true
	__Classdiagram__000001_spoil_diagram.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000001_spoil_diagram.NodeGongNotesIsExpanded = false
	__Classdiagram__000001_spoil_diagram.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.Name = `Diagram Package created the 2025-05-04T22:53:27Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.AbsolutePathToDiagramPackage = ``

	__GongNoteShape__000000_Default_NoteOnGongdoc.Name = `Default-NoteOnGongdoc`

	__GongNoteShape__000000_Default_NoteOnGongdoc.Identifier = `ref_models.NoteOnGongdoc`
	__GongNoteShape__000000_Default_NoteOnGongdoc.Body = `Note Example

This note can refers to [models.GongNoteShape]
or to [models.Classdiagram.GongNoteShapes]
or to [models.OrientationType]
`
	__GongNoteShape__000000_Default_NoteOnGongdoc.BodyHTML = `<p>Note Example
<p>This note can refers to <a href="/models#GongNoteShape">models.GongNoteShape</a>
or to <a href="/models#Classdiagram.GongNoteShapes">models.Classdiagram.GongNoteShapes</a>
or to <a href="/models#OrientationType">models.OrientationType</a>
`
	__GongNoteShape__000000_Default_NoteOnGongdoc.X = 61.000000
	__GongNoteShape__000000_Default_NoteOnGongdoc.Y = 196.000000
	__GongNoteShape__000000_Default_NoteOnGongdoc.Width = 370.000000
	__GongNoteShape__000000_Default_NoteOnGongdoc.Height = 112.000000
	__GongNoteShape__000000_Default_NoteOnGongdoc.Matched = false
	__GongNoteShape__000000_Default_NoteOnGongdoc.IsExpanded = false

	__GongStructShape__000000_Default_AttributeShape.Name = `Default-AttributeShape`
	__GongStructShape__000000_Default_AttributeShape.X = 984.000000
	__GongStructShape__000000_Default_AttributeShape.Y = 230.000000

	__GongStructShape__000000_Default_AttributeShape.ShowNbInstances = false
	__GongStructShape__000000_Default_AttributeShape.NbInstances = 0
	__GongStructShape__000000_Default_AttributeShape.Width = 240.000000
	__GongStructShape__000000_Default_AttributeShape.Height = 63.000000
	__GongStructShape__000000_Default_AttributeShape.IsSelected = false

	__GongStructShape__000001_Default_DiagramPackage.Name = `Default-DiagramPackage`
	__GongStructShape__000001_Default_DiagramPackage.X = 58.000000
	__GongStructShape__000001_Default_DiagramPackage.Y = 36.000000

	__GongStructShape__000001_Default_DiagramPackage.ShowNbInstances = false
	__GongStructShape__000001_Default_DiagramPackage.NbInstances = 0
	__GongStructShape__000001_Default_DiagramPackage.Width = 240.000000
	__GongStructShape__000001_Default_DiagramPackage.Height = 63.000000
	__GongStructShape__000001_Default_DiagramPackage.IsSelected = false

	__GongStructShape__000002_Default_GongStructShape.Name = `Default-GongStructShape`
	__GongStructShape__000002_Default_GongStructShape.X = 985.000000
	__GongStructShape__000002_Default_GongStructShape.Y = 32.000000

	__GongStructShape__000002_Default_GongStructShape.ShowNbInstances = false
	__GongStructShape__000002_Default_GongStructShape.NbInstances = 0
	__GongStructShape__000002_Default_GongStructShape.Width = 240.000000
	__GongStructShape__000002_Default_GongStructShape.Height = 63.000000
	__GongStructShape__000002_Default_GongStructShape.IsSelected = false

	__GongStructShape__000003_Default_Classdiagram.Name = `Default-Classdiagram`
	__GongStructShape__000003_Default_Classdiagram.X = 545.000000
	__GongStructShape__000003_Default_Classdiagram.Y = 94.000000

	__GongStructShape__000003_Default_Classdiagram.ShowNbInstances = false
	__GongStructShape__000003_Default_Classdiagram.NbInstances = 0
	__GongStructShape__000003_Default_Classdiagram.Width = 240.000000
	__GongStructShape__000003_Default_Classdiagram.Height = 63.000000
	__GongStructShape__000003_Default_Classdiagram.IsSelected = false

	__GongStructShape__000004_Default_1_AttributeShape.Name = `Default_1-AttributeShape`
	__GongStructShape__000004_Default_1_AttributeShape.X = 25.000000
	__GongStructShape__000004_Default_1_AttributeShape.Y = 46.000000

	__GongStructShape__000004_Default_1_AttributeShape.ShowNbInstances = false
	__GongStructShape__000004_Default_1_AttributeShape.NbInstances = 0
	__GongStructShape__000004_Default_1_AttributeShape.Width = 240.000000
	__GongStructShape__000004_Default_1_AttributeShape.Height = 103.000000
	__GongStructShape__000004_Default_1_AttributeShape.IsSelected = false

	__GongStructShape__000005_Default_1_Classdiagram.Name = `Default_1-Classdiagram`
	__GongStructShape__000005_Default_1_Classdiagram.X = 401.000000
	__GongStructShape__000005_Default_1_Classdiagram.Y = 306.000000

	__GongStructShape__000005_Default_1_Classdiagram.ShowNbInstances = false
	__GongStructShape__000005_Default_1_Classdiagram.NbInstances = 0
	__GongStructShape__000005_Default_1_Classdiagram.Width = 240.000000
	__GongStructShape__000005_Default_1_Classdiagram.Height = 63.000000
	__GongStructShape__000005_Default_1_Classdiagram.IsSelected = false

	__GongStructShape__000006_Default_1_DiagramPackage.Name = `Default_1-DiagramPackage`
	__GongStructShape__000006_Default_1_DiagramPackage.X = 95.000000
	__GongStructShape__000006_Default_1_DiagramPackage.Y = 101.000000

	__GongStructShape__000006_Default_1_DiagramPackage.ShowNbInstances = false
	__GongStructShape__000006_Default_1_DiagramPackage.NbInstances = 0
	__GongStructShape__000006_Default_1_DiagramPackage.Width = 240.000000
	__GongStructShape__000006_Default_1_DiagramPackage.Height = 63.000000
	__GongStructShape__000006_Default_1_DiagramPackage.IsSelected = false

	__GongStructShape__000007_Default_1_GongEnumShape.Name = `Default_1-GongEnumShape`
	__GongStructShape__000007_Default_1_GongEnumShape.X = 416.000000
	__GongStructShape__000007_Default_1_GongEnumShape.Y = 105.000000

	__GongStructShape__000007_Default_1_GongEnumShape.ShowNbInstances = false
	__GongStructShape__000007_Default_1_GongEnumShape.NbInstances = 0
	__GongStructShape__000007_Default_1_GongEnumShape.Width = 240.000000
	__GongStructShape__000007_Default_1_GongEnumShape.Height = 63.000000
	__GongStructShape__000007_Default_1_GongEnumShape.IsSelected = false

	__GongStructShape__000008_Default_1_GongEnumValueShape.Name = `Default_1-GongEnumValueShape`
	__GongStructShape__000008_Default_1_GongEnumValueShape.X = 22.000000
	__GongStructShape__000008_Default_1_GongEnumValueShape.Y = 100.000000

	__GongStructShape__000008_Default_1_GongEnumValueShape.ShowNbInstances = false
	__GongStructShape__000008_Default_1_GongEnumValueShape.NbInstances = 0
	__GongStructShape__000008_Default_1_GongEnumValueShape.Width = 240.000000
	__GongStructShape__000008_Default_1_GongEnumValueShape.Height = 63.000000
	__GongStructShape__000008_Default_1_GongEnumValueShape.IsSelected = false

	__GongStructShape__000009_Default_1_GongNoteLinkShape.Name = `Default_1-GongNoteLinkShape`
	__GongStructShape__000009_Default_1_GongNoteLinkShape.X = 32.000000
	__GongStructShape__000009_Default_1_GongNoteLinkShape.Y = 100.000000

	__GongStructShape__000009_Default_1_GongNoteLinkShape.ShowNbInstances = false
	__GongStructShape__000009_Default_1_GongNoteLinkShape.NbInstances = 0
	__GongStructShape__000009_Default_1_GongNoteLinkShape.Width = 240.000000
	__GongStructShape__000009_Default_1_GongNoteLinkShape.Height = 63.000000
	__GongStructShape__000009_Default_1_GongNoteLinkShape.IsSelected = false

	__GongStructShape__000010_Default_1_GongNoteShape.Name = `Default_1-GongNoteShape`
	__GongStructShape__000010_Default_1_GongNoteShape.X = 33.000000
	__GongStructShape__000010_Default_1_GongNoteShape.Y = 105.000000

	__GongStructShape__000010_Default_1_GongNoteShape.ShowNbInstances = false
	__GongStructShape__000010_Default_1_GongNoteShape.NbInstances = 0
	__GongStructShape__000010_Default_1_GongNoteShape.Width = 240.000000
	__GongStructShape__000010_Default_1_GongNoteShape.Height = 63.000000
	__GongStructShape__000010_Default_1_GongNoteShape.IsSelected = false

	__GongStructShape__000011_Default_1_GongStructShape.Name = `Default_1-GongStructShape`
	__GongStructShape__000011_Default_1_GongStructShape.X = 335.000000
	__GongStructShape__000011_Default_1_GongStructShape.Y = 542.000000

	__GongStructShape__000011_Default_1_GongStructShape.ShowNbInstances = false
	__GongStructShape__000011_Default_1_GongStructShape.NbInstances = 0
	__GongStructShape__000011_Default_1_GongStructShape.Width = 240.000000
	__GongStructShape__000011_Default_1_GongStructShape.Height = 63.000000
	__GongStructShape__000011_Default_1_GongStructShape.IsSelected = false

	__GongStructShape__000012_Default_1_LinkShape.Name = `Default_1-LinkShape`
	__GongStructShape__000012_Default_1_LinkShape.X = 91.000000
	__GongStructShape__000012_Default_1_LinkShape.Y = 77.000000

	__GongStructShape__000012_Default_1_LinkShape.ShowNbInstances = false
	__GongStructShape__000012_Default_1_LinkShape.NbInstances = 0
	__GongStructShape__000012_Default_1_LinkShape.Width = 240.000000
	__GongStructShape__000012_Default_1_LinkShape.Height = 103.000000
	__GongStructShape__000012_Default_1_LinkShape.IsSelected = false

	__LinkShape__000000_Classdiagrams.Name = `Classdiagrams`

	__LinkShape__000000_Classdiagrams.Identifier = `ref_models.DiagramPackage.Classdiagrams`

	__LinkShape__000000_Classdiagrams.FieldOffsetX = 0.000000
	__LinkShape__000000_Classdiagrams.FieldOffsetY = 0.000000
	__LinkShape__000000_Classdiagrams.TargetMultiplicity = models.MANY
	__LinkShape__000000_Classdiagrams.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Classdiagrams.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Classdiagrams.SourceMultiplicity = models.MANY
	__LinkShape__000000_Classdiagrams.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Classdiagrams.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Classdiagrams.X = 441.500000
	__LinkShape__000000_Classdiagrams.Y = 162.000000
	__LinkShape__000000_Classdiagrams.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Classdiagrams.StartRatio = 0.500000
	__LinkShape__000000_Classdiagrams.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Classdiagrams.EndRatio = 0.460379
	__LinkShape__000000_Classdiagrams.CornerOffsetRatio = 1.380000

	__LinkShape__000001_GongStructShapes.Name = `GongStructShapes`

	__LinkShape__000001_GongStructShapes.FieldOffsetX = 0.000000
	__LinkShape__000001_GongStructShapes.FieldOffsetY = 0.000000
	__LinkShape__000001_GongStructShapes.TargetMultiplicity = models.MANY
	__LinkShape__000001_GongStructShapes.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000001_GongStructShapes.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000001_GongStructShapes.SourceMultiplicity = models.MANY
	__LinkShape__000001_GongStructShapes.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000001_GongStructShapes.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000001_GongStructShapes.X = 1140.500000
	__LinkShape__000001_GongStructShapes.Y = 197.000000
	__LinkShape__000001_GongStructShapes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_GongStructShapes.StartRatio = 0.500000
	__LinkShape__000001_GongStructShapes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_GongStructShapes.EndRatio = 0.412760
	__LinkShape__000001_GongStructShapes.CornerOffsetRatio = 1.380000

	__LinkShape__000002_AttributeShapes.Name = `AttributeShapes`

	__LinkShape__000002_AttributeShapes.Identifier = `ref_models.GongStructShape.AttributeShapes`

	__LinkShape__000002_AttributeShapes.FieldOffsetX = 0.000000
	__LinkShape__000002_AttributeShapes.FieldOffsetY = 0.000000
	__LinkShape__000002_AttributeShapes.TargetMultiplicity = models.MANY
	__LinkShape__000002_AttributeShapes.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000002_AttributeShapes.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000002_AttributeShapes.SourceMultiplicity = models.MANY
	__LinkShape__000002_AttributeShapes.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000002_AttributeShapes.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000002_AttributeShapes.X = 1575.500000
	__LinkShape__000002_AttributeShapes.Y = 191.500000
	__LinkShape__000002_AttributeShapes.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000002_AttributeShapes.StartRatio = 0.780241
	__LinkShape__000002_AttributeShapes.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000002_AttributeShapes.EndRatio = 0.734408
	__LinkShape__000002_AttributeShapes.CornerOffsetRatio = 2.055711

	__LinkShape__000003_GongStructShapes.Name = `GongStructShapes`

	__LinkShape__000003_GongStructShapes.FieldOffsetX = 0.000000
	__LinkShape__000003_GongStructShapes.FieldOffsetY = 0.000000
	__LinkShape__000003_GongStructShapes.TargetMultiplicity = models.MANY
	__LinkShape__000003_GongStructShapes.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000003_GongStructShapes.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000003_GongStructShapes.SourceMultiplicity = models.MANY
	__LinkShape__000003_GongStructShapes.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000003_GongStructShapes.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000003_GongStructShapes.X = 394.500000
	__LinkShape__000003_GongStructShapes.Y = 173.000000
	__LinkShape__000003_GongStructShapes.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000003_GongStructShapes.StartRatio = 0.680241
	__LinkShape__000003_GongStructShapes.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000003_GongStructShapes.EndRatio = 0.955241
	__LinkShape__000003_GongStructShapes.CornerOffsetRatio = 1.690631

	__LinkShape__000004_GongEnumShapes.Name = `GongEnumShapes`

	__LinkShape__000004_GongEnumShapes.Identifier = `ref_models.Classdiagram.GongEnumShapes`

	__LinkShape__000004_GongEnumShapes.FieldOffsetX = 0.000000
	__LinkShape__000004_GongEnumShapes.FieldOffsetY = 0.000000
	__LinkShape__000004_GongEnumShapes.TargetMultiplicity = models.MANY
	__LinkShape__000004_GongEnumShapes.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000004_GongEnumShapes.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000004_GongEnumShapes.SourceMultiplicity = models.MANY
	__LinkShape__000004_GongEnumShapes.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000004_GongEnumShapes.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000004_GongEnumShapes.X = 440.500000
	__LinkShape__000004_GongEnumShapes.Y = 187.000000
	__LinkShape__000004_GongEnumShapes.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000004_GongEnumShapes.StartRatio = 0.646908
	__LinkShape__000004_GongEnumShapes.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000004_GongEnumShapes.EndRatio = 0.717741
	__LinkShape__000004_GongEnumShapes.CornerOffsetRatio = -0.960162

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_AttributeShape)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_DiagramPackage)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_GongStructShape)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_Classdiagram)
	__Classdiagram__000000_Default.GongNoteShapes = append(__Classdiagram__000000_Default.GongNoteShapes, __GongNoteShape__000000_Default_NoteOnGongdoc)
	__Classdiagram__000001_spoil_diagram.GongStructShapes = append(__Classdiagram__000001_spoil_diagram.GongStructShapes, __GongStructShape__000004_Default_1_AttributeShape)
	__Classdiagram__000001_spoil_diagram.GongStructShapes = append(__Classdiagram__000001_spoil_diagram.GongStructShapes, __GongStructShape__000005_Default_1_Classdiagram)
	__Classdiagram__000001_spoil_diagram.GongStructShapes = append(__Classdiagram__000001_spoil_diagram.GongStructShapes, __GongStructShape__000006_Default_1_DiagramPackage)
	__Classdiagram__000001_spoil_diagram.GongStructShapes = append(__Classdiagram__000001_spoil_diagram.GongStructShapes, __GongStructShape__000007_Default_1_GongEnumShape)
	__Classdiagram__000001_spoil_diagram.GongStructShapes = append(__Classdiagram__000001_spoil_diagram.GongStructShapes, __GongStructShape__000008_Default_1_GongEnumValueShape)
	__Classdiagram__000001_spoil_diagram.GongStructShapes = append(__Classdiagram__000001_spoil_diagram.GongStructShapes, __GongStructShape__000009_Default_1_GongNoteLinkShape)
	__Classdiagram__000001_spoil_diagram.GongStructShapes = append(__Classdiagram__000001_spoil_diagram.GongStructShapes, __GongStructShape__000010_Default_1_GongNoteShape)
	__Classdiagram__000001_spoil_diagram.GongStructShapes = append(__Classdiagram__000001_spoil_diagram.GongStructShapes, __GongStructShape__000011_Default_1_GongStructShape)
	__Classdiagram__000001_spoil_diagram.GongStructShapes = append(__Classdiagram__000001_spoil_diagram.GongStructShapes, __GongStructShape__000012_Default_1_LinkShape)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.Classdiagrams, __Classdiagram__000001_spoil_diagram)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongNoteShape instances pointers
	// setup of GongStructShape instances pointers
	__GongStructShape__000001_Default_DiagramPackage.LinkShapes = append(__GongStructShape__000001_Default_DiagramPackage.LinkShapes, __LinkShape__000000_Classdiagrams)
	__GongStructShape__000002_Default_GongStructShape.LinkShapes = append(__GongStructShape__000002_Default_GongStructShape.LinkShapes, __LinkShape__000002_AttributeShapes)
	__GongStructShape__000003_Default_Classdiagram.LinkShapes = append(__GongStructShape__000003_Default_Classdiagram.LinkShapes, __LinkShape__000001_GongStructShapes)
	__GongStructShape__000004_Default_1_AttributeShape.AttributeShapes = append(__GongStructShape__000004_Default_1_AttributeShape.AttributeShapes, __AttributeShape__000000_FieldTypeAsString)
	__GongStructShape__000004_Default_1_AttributeShape.AttributeShapes = append(__GongStructShape__000004_Default_1_AttributeShape.AttributeShapes, __AttributeShape__000001_Fieldtypename)
	__GongStructShape__000005_Default_1_Classdiagram.LinkShapes = append(__GongStructShape__000005_Default_1_Classdiagram.LinkShapes, __LinkShape__000003_GongStructShapes)
	__GongStructShape__000005_Default_1_Classdiagram.LinkShapes = append(__GongStructShape__000005_Default_1_Classdiagram.LinkShapes, __LinkShape__000004_GongEnumShapes)
	__GongStructShape__000012_Default_1_LinkShape.AttributeShapes = append(__GongStructShape__000012_Default_1_LinkShape.AttributeShapes, __AttributeShape__000002_TargetMultiplicityOffsetX)
	__GongStructShape__000012_Default_1_LinkShape.AttributeShapes = append(__GongStructShape__000012_Default_1_LinkShape.AttributeShapes, __AttributeShape__000003_TargetMultiplicityOffsetY)
	// setup of LinkShape instances pointers
}
