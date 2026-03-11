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

	// insertion point for declaration of instances to stage

	__AttributeShape__00000000_ := (&models.AttributeShape{Name: `X`}).Stage(stage)
	__AttributeShape__00000001_ := (&models.AttributeShape{Name: `Y`}).Stage(stage)
	__AttributeShape__00000002_ := (&models.AttributeShape{Name: `IdentifierMeta`}).Stage(stage)
	__AttributeShape__00000003_ := (&models.AttributeShape{Name: `Description`}).Stage(stage)
	__AttributeShape__00000004_ := (&models.AttributeShape{Name: `IsIncludedInStaticWebSite`}).Stage(stage)
	__AttributeShape__00000005_ := (&models.AttributeShape{Name: `X`}).Stage(stage)
	__AttributeShape__00000006_ := (&models.AttributeShape{Name: `Y`}).Stage(stage)
	__AttributeShape__00000007_ := (&models.AttributeShape{Name: `IdentifierMeta`}).Stage(stage)
	__AttributeShape__00000008_ := (&models.AttributeShape{Name: `Height`}).Stage(stage)
	__AttributeShape__00000009_ := (&models.AttributeShape{Name: `Width`}).Stage(stage)
	__AttributeShape__00000010_ := (&models.AttributeShape{Name: `IdentifierMeta`}).Stage(stage)
	__AttributeShape__00000011_ := (&models.AttributeShape{Name: `FieldTypeIdentifierMeta`}).Stage(stage)
	__AttributeShape__00000012_ := (&models.AttributeShape{Name: `FieldOffsetX`}).Stage(stage)
	__AttributeShape__00000013_ := (&models.AttributeShape{Name: `FieldOffsetY`}).Stage(stage)
	__AttributeShape__00000014_ := (&models.AttributeShape{Name: `TargetMultiplicity`}).Stage(stage)
	__AttributeShape__00000015_ := (&models.AttributeShape{Name: `TargetMultiplicityOffsetX`}).Stage(stage)
	__AttributeShape__00000016_ := (&models.AttributeShape{Name: `TargetMultiplicityOffsetY`}).Stage(stage)
	__AttributeShape__00000017_ := (&models.AttributeShape{Name: `SourceMultiplicity`}).Stage(stage)
	__AttributeShape__00000018_ := (&models.AttributeShape{Name: `SourceMultiplicityOffsetX`}).Stage(stage)
	__AttributeShape__00000019_ := (&models.AttributeShape{Name: `IdentifierMeta`}).Stage(stage)
	__AttributeShape__00000020_ := (&models.AttributeShape{Name: `FieldTypeAsString`}).Stage(stage)
	__AttributeShape__00000021_ := (&models.AttributeShape{Name: `Structname`}).Stage(stage)
	__AttributeShape__00000022_ := (&models.AttributeShape{Name: `Fieldtypename`}).Stage(stage)

	__Classdiagram__00000000_ := (&models.Classdiagram{Name: `Default`}).Stage(stage)
	__Classdiagram__00000001_ := (&models.Classdiagram{Name: `spoil diagram`}).Stage(stage)
	__Classdiagram__00000002_ := (&models.Classdiagram{Name: `NoteShape`}).Stage(stage)

	__DiagramPackage__00000000_ := (&models.DiagramPackage{Name: `Diagram Package created the 2025-05-04T22:53:27Z`}).Stage(stage)

	__GongEnumShape__00000000_ := (&models.GongEnumShape{Name: `Default-ButtonType`}).Stage(stage)

	__GongEnumValueShape__00000000_ := (&models.GongEnumValueShape{Name: `DUPLICATE`}).Stage(stage)
	__GongEnumValueShape__00000001_ := (&models.GongEnumValueShape{Name: `EDIT_CANCEL`}).Stage(stage)
	__GongEnumValueShape__00000002_ := (&models.GongEnumValueShape{Name: `EDIT`}).Stage(stage)

	__GongNoteShape__00000000_ := (&models.GongNoteShape{Name: `Default-NoteOnGongdoc`}).Stage(stage)

	__GongStructShape__00000000_ := (&models.GongStructShape{Name: `Default-AttributeShape`}).Stage(stage)
	__GongStructShape__00000001_ := (&models.GongStructShape{Name: `Default-Classdiagram`}).Stage(stage)
	__GongStructShape__00000002_ := (&models.GongStructShape{Name: `NoteShape-GongNoteShape`}).Stage(stage)
	__GongStructShape__00000003_ := (&models.GongStructShape{Name: `NoteShape-GongNoteLinkShape`}).Stage(stage)
	__GongStructShape__00000004_ := (&models.GongStructShape{Name: `Default-GongEnumShape`}).Stage(stage)
	__GongStructShape__00000006_ := (&models.GongStructShape{Name: `Default-GongStructShape`}).Stage(stage)
	__GongStructShape__00000007_ := (&models.GongStructShape{Name: `Default-LinkShape`}).Stage(stage)

	__LinkShape__00000000_ := (&models.LinkShape{Name: `GongNoteLinkShapes`}).Stage(stage)
	__LinkShape__00000003_ := (&models.LinkShape{Name: `GongStructShapes`}).Stage(stage)
	__LinkShape__00000004_ := (&models.LinkShape{Name: `AttributeShapes`}).Stage(stage)
	__LinkShape__00000005_ := (&models.LinkShape{Name: `LinkShapes`}).Stage(stage)
	__LinkShape__00000006_ := (&models.LinkShape{Name: `GongEnumShapes`}).Stage(stage)

	// insertion point for initialization of values

	__AttributeShape__00000000_.Name = `X`
	__AttributeShape__00000000_.IdentifierMeta = ref_models.GongEnumShape{}.X
	__AttributeShape__00000000_.FieldTypeAsString = ``
	__AttributeShape__00000000_.Structname = `GongEnumShape`
	__AttributeShape__00000000_.Fieldtypename = `float64`

	__AttributeShape__00000001_.Name = `Y`
	__AttributeShape__00000001_.IdentifierMeta = ref_models.GongEnumShape{}.Y
	__AttributeShape__00000001_.FieldTypeAsString = ``
	__AttributeShape__00000001_.Structname = `GongEnumShape`
	__AttributeShape__00000001_.Fieldtypename = `float64`

	__AttributeShape__00000002_.Name = `IdentifierMeta`
	__AttributeShape__00000002_.IdentifierMeta = ref_models.GongEnumShape{}.IdentifierMeta
	__AttributeShape__00000002_.FieldTypeAsString = ``
	__AttributeShape__00000002_.Structname = `GongEnumShape`
	__AttributeShape__00000002_.Fieldtypename = `any`

	__AttributeShape__00000003_.Name = `Description`
	__AttributeShape__00000003_.IdentifierMeta = ref_models.Classdiagram{}.Description
	__AttributeShape__00000003_.FieldTypeAsString = ``
	__AttributeShape__00000003_.Structname = `Classdiagram`
	__AttributeShape__00000003_.Fieldtypename = `string`

	__AttributeShape__00000004_.Name = `IsIncludedInStaticWebSite`
	__AttributeShape__00000004_.IdentifierMeta = ref_models.Classdiagram{}.IsIncludedInStaticWebSite
	__AttributeShape__00000004_.FieldTypeAsString = ``
	__AttributeShape__00000004_.Structname = `Classdiagram`
	__AttributeShape__00000004_.Fieldtypename = `bool`

	__AttributeShape__00000005_.Name = `X`
	__AttributeShape__00000005_.IdentifierMeta = ref_models.GongStructShape{}.X
	__AttributeShape__00000005_.FieldTypeAsString = ``
	__AttributeShape__00000005_.Structname = `GongStructShape`
	__AttributeShape__00000005_.Fieldtypename = `float64`

	__AttributeShape__00000006_.Name = `Y`
	__AttributeShape__00000006_.IdentifierMeta = ref_models.GongStructShape{}.Y
	__AttributeShape__00000006_.FieldTypeAsString = ``
	__AttributeShape__00000006_.Structname = `GongStructShape`
	__AttributeShape__00000006_.Fieldtypename = `float64`

	__AttributeShape__00000007_.Name = `IdentifierMeta`
	__AttributeShape__00000007_.IdentifierMeta = ref_models.GongStructShape{}.IdentifierMeta
	__AttributeShape__00000007_.FieldTypeAsString = ``
	__AttributeShape__00000007_.Structname = `GongStructShape`
	__AttributeShape__00000007_.Fieldtypename = `any`

	__AttributeShape__00000008_.Name = `Height`
	__AttributeShape__00000008_.IdentifierMeta = ref_models.GongStructShape{}.Height
	__AttributeShape__00000008_.FieldTypeAsString = ``
	__AttributeShape__00000008_.Structname = `GongStructShape`
	__AttributeShape__00000008_.Fieldtypename = `float64`

	__AttributeShape__00000009_.Name = `Width`
	__AttributeShape__00000009_.IdentifierMeta = ref_models.GongStructShape{}.Width
	__AttributeShape__00000009_.FieldTypeAsString = ``
	__AttributeShape__00000009_.Structname = `GongStructShape`
	__AttributeShape__00000009_.Fieldtypename = `float64`

	__AttributeShape__00000010_.Name = `IdentifierMeta`
	__AttributeShape__00000010_.IdentifierMeta = ref_models.LinkShape{}.IdentifierMeta
	__AttributeShape__00000010_.FieldTypeAsString = ``
	__AttributeShape__00000010_.Structname = `LinkShape`
	__AttributeShape__00000010_.Fieldtypename = `any`

	__AttributeShape__00000011_.Name = `FieldTypeIdentifierMeta`
	__AttributeShape__00000011_.IdentifierMeta = ref_models.LinkShape{}.FieldTypeIdentifierMeta
	__AttributeShape__00000011_.FieldTypeAsString = ``
	__AttributeShape__00000011_.Structname = `LinkShape`
	__AttributeShape__00000011_.Fieldtypename = `any`

	__AttributeShape__00000012_.Name = `FieldOffsetX`
	__AttributeShape__00000012_.IdentifierMeta = ref_models.LinkShape{}.FieldOffsetX
	__AttributeShape__00000012_.FieldTypeAsString = ``
	__AttributeShape__00000012_.Structname = `LinkShape`
	__AttributeShape__00000012_.Fieldtypename = `float64`

	__AttributeShape__00000013_.Name = `FieldOffsetY`
	__AttributeShape__00000013_.IdentifierMeta = ref_models.LinkShape{}.FieldOffsetY
	__AttributeShape__00000013_.FieldTypeAsString = ``
	__AttributeShape__00000013_.Structname = `LinkShape`
	__AttributeShape__00000013_.Fieldtypename = `float64`

	__AttributeShape__00000014_.Name = `TargetMultiplicity`
	__AttributeShape__00000014_.IdentifierMeta = ref_models.LinkShape{}.TargetMultiplicity
	__AttributeShape__00000014_.FieldTypeAsString = ``
	__AttributeShape__00000014_.Structname = `LinkShape`
	__AttributeShape__00000014_.Fieldtypename = `MultiplicityType`

	__AttributeShape__00000015_.Name = `TargetMultiplicityOffsetX`
	__AttributeShape__00000015_.IdentifierMeta = ref_models.LinkShape{}.TargetMultiplicityOffsetX
	__AttributeShape__00000015_.FieldTypeAsString = ``
	__AttributeShape__00000015_.Structname = `LinkShape`
	__AttributeShape__00000015_.Fieldtypename = `float64`

	__AttributeShape__00000016_.Name = `TargetMultiplicityOffsetY`
	__AttributeShape__00000016_.IdentifierMeta = ref_models.LinkShape{}.TargetMultiplicityOffsetY
	__AttributeShape__00000016_.FieldTypeAsString = ``
	__AttributeShape__00000016_.Structname = `LinkShape`
	__AttributeShape__00000016_.Fieldtypename = `float64`

	__AttributeShape__00000017_.Name = `SourceMultiplicity`
	__AttributeShape__00000017_.IdentifierMeta = ref_models.LinkShape{}.SourceMultiplicity
	__AttributeShape__00000017_.FieldTypeAsString = ``
	__AttributeShape__00000017_.Structname = `LinkShape`
	__AttributeShape__00000017_.Fieldtypename = `MultiplicityType`

	__AttributeShape__00000018_.Name = `SourceMultiplicityOffsetX`
	__AttributeShape__00000018_.IdentifierMeta = ref_models.LinkShape{}.SourceMultiplicityOffsetX
	__AttributeShape__00000018_.FieldTypeAsString = ``
	__AttributeShape__00000018_.Structname = `LinkShape`
	__AttributeShape__00000018_.Fieldtypename = `float64`

	__AttributeShape__00000019_.Name = `IdentifierMeta`
	__AttributeShape__00000019_.IdentifierMeta = ref_models.AttributeShape{}.IdentifierMeta
	__AttributeShape__00000019_.FieldTypeAsString = ``
	__AttributeShape__00000019_.Structname = `AttributeShape`
	__AttributeShape__00000019_.Fieldtypename = `any`

	__AttributeShape__00000020_.Name = `FieldTypeAsString`
	__AttributeShape__00000020_.IdentifierMeta = ref_models.AttributeShape{}.FieldTypeAsString
	__AttributeShape__00000020_.FieldTypeAsString = ``
	__AttributeShape__00000020_.Structname = `AttributeShape`
	__AttributeShape__00000020_.Fieldtypename = `string`

	__AttributeShape__00000021_.Name = `Structname`
	__AttributeShape__00000021_.IdentifierMeta = ref_models.AttributeShape{}.Structname
	__AttributeShape__00000021_.FieldTypeAsString = ``
	__AttributeShape__00000021_.Structname = `AttributeShape`
	__AttributeShape__00000021_.Fieldtypename = `string`

	__AttributeShape__00000022_.Name = `Fieldtypename`
	__AttributeShape__00000022_.IdentifierMeta = ref_models.AttributeShape{}.Fieldtypename
	__AttributeShape__00000022_.FieldTypeAsString = ``
	__AttributeShape__00000022_.Structname = `AttributeShape`
	__AttributeShape__00000022_.Fieldtypename = `string`

	__Classdiagram__00000000_.Name = `Default`
	__Classdiagram__00000000_.Description = `This diagram describes the model of the doc package. A DiagramPackage is composed of ClassDiagram and each ClassDiagram has shapes.`
	__Classdiagram__00000000_.IsIncludedInStaticWebSite = true
	__Classdiagram__00000000_.ShowNbInstances = true
	__Classdiagram__00000000_.ShowMultiplicity = true
	__Classdiagram__00000000_.ShowLinkNames = true
	__Classdiagram__00000000_.IsInRenameMode = false
	__Classdiagram__00000000_.IsExpanded = true
	__Classdiagram__00000000_.NodeGongStructsIsExpanded = false
	__Classdiagram__00000000_.NodeGongStructNodeExpansion = `[false,true,false,false,false,false,false,true,true]`
	__Classdiagram__00000000_.NodeGongEnumsIsExpanded = true
	__Classdiagram__00000000_.NodeGongEnumNodeExpansion = `[false]`
	__Classdiagram__00000000_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000000_.NodeGongNoteNodeExpansion = `[true]`

	__Classdiagram__00000001_.Name = `spoil diagram`
	__Classdiagram__00000001_.Description = `Spoil diagram`
	__Classdiagram__00000001_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000001_.ShowNbInstances = false
	__Classdiagram__00000001_.ShowMultiplicity = false
	__Classdiagram__00000001_.ShowLinkNames = false
	__Classdiagram__00000001_.IsInRenameMode = false
	__Classdiagram__00000001_.IsExpanded = false
	__Classdiagram__00000001_.NodeGongStructsIsExpanded = false
	__Classdiagram__00000001_.NodeGongStructNodeExpansion = ``
	__Classdiagram__00000001_.NodeGongEnumsIsExpanded = true
	__Classdiagram__00000001_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000001_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000001_.NodeGongNoteNodeExpansion = ``

	__Classdiagram__00000002_.Name = `NoteShape`
	__Classdiagram__00000002_.Description = ``
	__Classdiagram__00000002_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000002_.ShowNbInstances = true
	__Classdiagram__00000002_.ShowMultiplicity = true
	__Classdiagram__00000002_.ShowLinkNames = true
	__Classdiagram__00000002_.IsInRenameMode = false
	__Classdiagram__00000002_.IsExpanded = false
	__Classdiagram__00000002_.NodeGongStructsIsExpanded = true
	__Classdiagram__00000002_.NodeGongStructNodeExpansion = `[false,false,false,false,false,false,true]`
	__Classdiagram__00000002_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000002_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000002_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000002_.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__00000000_.Name = `Diagram Package created the 2025-05-04T22:53:27Z`
	__DiagramPackage__00000000_.Path = ``
	__DiagramPackage__00000000_.GongModelPath = ``
	__DiagramPackage__00000000_.AbsolutePathToDiagramPackage = ``

	__GongEnumShape__00000000_.Name = `Default-ButtonType`
	__GongEnumShape__00000000_.X = 1198.000000
	__GongEnumShape__00000000_.Y = 57.000000
	__GongEnumShape__00000000_.IdentifierMeta = new(ref_models.ButtonType)
	__GongEnumShape__00000000_.Width = 240.000000
	__GongEnumShape__00000000_.Height = 123.000000
	__GongEnumShape__00000000_.IsExpanded = false

	__GongEnumValueShape__00000000_.Name = `DUPLICATE`
	__GongEnumValueShape__00000000_.IdentifierMeta = ref_models.DUPLICATE

	__GongEnumValueShape__00000001_.Name = `EDIT_CANCEL`
	__GongEnumValueShape__00000001_.IdentifierMeta = ref_models.EDIT_CANCEL

	__GongEnumValueShape__00000002_.Name = `EDIT`
	__GongEnumValueShape__00000002_.IdentifierMeta = ref_models.EDIT

	__GongNoteShape__00000000_.Name = `Default-NoteOnGongdoc`
	__GongNoteShape__00000000_.Identifier = `ref_models.NoteOnGongdoc`
	__GongNoteShape__00000000_.Body = `Note Example

This note can refers to [models.GongNoteShape]
or to [models.Classdiagram.GongNoteShapes]
or to [models.OrientationType]
`
	__GongNoteShape__00000000_.BodyHTML = `<p>Note Example
<p>This note can refers to <a href="/models#GongNoteShape">models.GongNoteShape</a>
or to <a href="/models#Classdiagram.GongNoteShapes">models.Classdiagram.GongNoteShapes</a>
or to <a href="/models#OrientationType">models.OrientationType</a>
`
	__GongNoteShape__00000000_.X = 87.000000
	__GongNoteShape__00000000_.Y = 393.000000
	__GongNoteShape__00000000_.Width = 279.000000
	__GongNoteShape__00000000_.Height = 110.000000
	__GongNoteShape__00000000_.Matched = false
	__GongNoteShape__00000000_.IsExpanded = false

	__GongStructShape__00000000_.Name = `Default-AttributeShape`
	__GongStructShape__00000000_.X = 1167.000000
	__GongStructShape__00000000_.Y = 330.000000
	__GongStructShape__00000000_.IdentifierMeta = ref_models.AttributeShape{}
	__GongStructShape__00000000_.Width = 366.000000
	__GongStructShape__00000000_.Height = 143.000000
	__GongStructShape__00000000_.IsSelected = false

	__GongStructShape__00000001_.Name = `Default-Classdiagram`
	__GongStructShape__00000001_.X = 109.000000
	__GongStructShape__00000001_.Y = 85.000000
	__GongStructShape__00000001_.IdentifierMeta = ref_models.Classdiagram{}
	__GongStructShape__00000001_.Width = 278.000000
	__GongStructShape__00000001_.Height = 103.000000
	__GongStructShape__00000001_.IsSelected = false

	__GongStructShape__00000002_.Name = `NoteShape-GongNoteShape`
	__GongStructShape__00000002_.X = 89.000000
	__GongStructShape__00000002_.Y = 87.000000
	__GongStructShape__00000002_.IdentifierMeta = ref_models.GongNoteShape{}
	__GongStructShape__00000002_.Width = 240.000000
	__GongStructShape__00000002_.Height = 63.000000
	__GongStructShape__00000002_.IsSelected = false

	__GongStructShape__00000003_.Name = `NoteShape-GongNoteLinkShape`
	__GongStructShape__00000003_.X = 639.000000
	__GongStructShape__00000003_.Y = 88.000000
	__GongStructShape__00000003_.IdentifierMeta = ref_models.GongNoteLinkShape{}
	__GongStructShape__00000003_.Width = 240.000000
	__GongStructShape__00000003_.Height = 63.000000
	__GongStructShape__00000003_.IsSelected = false

	__GongStructShape__00000004_.Name = `Default-GongEnumShape`
	__GongStructShape__00000004_.X = 632.000000
	__GongStructShape__00000004_.Y = 83.000000
	__GongStructShape__00000004_.IdentifierMeta = ref_models.GongEnumShape{}
	__GongStructShape__00000004_.Width = 240.000000
	__GongStructShape__00000004_.Height = 123.000000
	__GongStructShape__00000004_.IsSelected = false

	__GongStructShape__00000006_.Name = `Default-GongStructShape`
	__GongStructShape__00000006_.X = 637.000000
	__GongStructShape__00000006_.Y = 304.000000
	__GongStructShape__00000006_.IdentifierMeta = ref_models.GongStructShape{}
	__GongStructShape__00000006_.Width = 240.000000
	__GongStructShape__00000006_.Height = 349.000000
	__GongStructShape__00000006_.IsSelected = false

	__GongStructShape__00000007_.Name = `Default-LinkShape`
	__GongStructShape__00000007_.X = 1169.000000
	__GongStructShape__00000007_.Y = 491.000000
	__GongStructShape__00000007_.IdentifierMeta = ref_models.LinkShape{}
	__GongStructShape__00000007_.Width = 370.000000
	__GongStructShape__00000007_.Height = 243.000000
	__GongStructShape__00000007_.IsSelected = false

	__LinkShape__00000000_.Name = `GongNoteLinkShapes`
	__LinkShape__00000000_.IdentifierMeta = ref_models.GongNoteShape{}.GongNoteLinkShapes
	__LinkShape__00000000_.FieldTypeIdentifierMeta = ref_models.GongNoteLinkShape{}
	__LinkShape__00000000_.FieldOffsetX = 0.000000
	__LinkShape__00000000_.FieldOffsetY = 0.000000
	__LinkShape__00000000_.TargetMultiplicity = models.MANY
	__LinkShape__00000000_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.SourceMultiplicity = models.MANY
	__LinkShape__00000000_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.X = 655.500000
	__LinkShape__00000000_.Y = 120.500000
	__LinkShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.StartRatio = 0.500000
	__LinkShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.EndRatio = 0.500000
	__LinkShape__00000000_.CornerOffsetRatio = 1.380000

	__LinkShape__00000003_.Name = `GongStructShapes`
	__LinkShape__00000003_.IdentifierMeta = ref_models.Classdiagram{}.GongStructShapes
	__LinkShape__00000003_.FieldTypeIdentifierMeta = ref_models.GongStructShape{}
	__LinkShape__00000003_.FieldOffsetX = 0.000000
	__LinkShape__00000003_.FieldOffsetY = 0.000000
	__LinkShape__00000003_.TargetMultiplicity = models.MANY
	__LinkShape__00000003_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.SourceMultiplicity = models.MANY
	__LinkShape__00000003_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.X = 790.000000
	__LinkShape__00000003_.Y = 246.000000
	__LinkShape__00000003_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000003_.StartRatio = 0.508590
	__LinkShape__00000003_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000003_.EndRatio = 0.500000
	__LinkShape__00000003_.CornerOffsetRatio = 1.212708

	__LinkShape__00000004_.Name = `AttributeShapes`
	__LinkShape__00000004_.IdentifierMeta = ref_models.GongStructShape{}.AttributeShapes
	__LinkShape__00000004_.FieldTypeIdentifierMeta = ref_models.AttributeShape{}
	__LinkShape__00000004_.FieldOffsetX = 0.000000
	__LinkShape__00000004_.FieldOffsetY = 0.000000
	__LinkShape__00000004_.TargetMultiplicity = models.MANY
	__LinkShape__00000004_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.SourceMultiplicity = models.MANY
	__LinkShape__00000004_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.X = 1158.000000
	__LinkShape__00000004_.Y = 423.500000
	__LinkShape__00000004_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000004_.StartRatio = 0.316289
	__LinkShape__00000004_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000004_.EndRatio = 0.500000
	__LinkShape__00000004_.CornerOffsetRatio = 1.380000

	__LinkShape__00000005_.Name = `LinkShapes`
	__LinkShape__00000005_.IdentifierMeta = ref_models.GongStructShape{}.LinkShapes
	__LinkShape__00000005_.FieldTypeIdentifierMeta = ref_models.LinkShape{}
	__LinkShape__00000005_.FieldOffsetX = 0.000000
	__LinkShape__00000005_.FieldOffsetY = 0.000000
	__LinkShape__00000005_.TargetMultiplicity = models.MANY
	__LinkShape__00000005_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.SourceMultiplicity = models.MANY
	__LinkShape__00000005_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.X = 1265.000000
	__LinkShape__00000005_.Y = 599.500000
	__LinkShape__00000005_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000005_.StartRatio = 0.791933
	__LinkShape__00000005_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000005_.EndRatio = 0.318456
	__LinkShape__00000005_.CornerOffsetRatio = 1.380000

	__LinkShape__00000006_.Name = `GongEnumShapes`
	__LinkShape__00000006_.IdentifierMeta = ref_models.Classdiagram{}.GongEnumShapes
	__LinkShape__00000006_.FieldTypeIdentifierMeta = ref_models.GongEnumShape{}
	__LinkShape__00000006_.FieldOffsetX = 0.000000
	__LinkShape__00000006_.FieldOffsetY = 0.000000
	__LinkShape__00000006_.TargetMultiplicity = models.MANY
	__LinkShape__00000006_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.SourceMultiplicity = models.MANY
	__LinkShape__00000006_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.X = 787.500000
	__LinkShape__00000006_.Y = 135.500000
	__LinkShape__00000006_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000006_.StartRatio = 0.500000
	__LinkShape__00000006_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000006_.EndRatio = 0.500000
	__LinkShape__00000006_.CornerOffsetRatio = 1.380000

	// insertion point for setup of pointers
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000000_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000001_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000004_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000006_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000007_)
	__Classdiagram__00000000_.GongEnumShapes = append(__Classdiagram__00000000_.GongEnumShapes, __GongEnumShape__00000000_)
	__Classdiagram__00000000_.GongNoteShapes = append(__Classdiagram__00000000_.GongNoteShapes, __GongNoteShape__00000000_)
	__Classdiagram__00000002_.GongStructShapes = append(__Classdiagram__00000002_.GongStructShapes, __GongStructShape__00000002_)
	__Classdiagram__00000002_.GongStructShapes = append(__Classdiagram__00000002_.GongStructShapes, __GongStructShape__00000003_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000000_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000001_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000002_)
	__DiagramPackage__00000000_.SelectedClassdiagram = __Classdiagram__00000000_
	__GongEnumShape__00000000_.GongEnumValueShapes = append(__GongEnumShape__00000000_.GongEnumValueShapes, __GongEnumValueShape__00000000_)
	__GongEnumShape__00000000_.GongEnumValueShapes = append(__GongEnumShape__00000000_.GongEnumValueShapes, __GongEnumValueShape__00000001_)
	__GongEnumShape__00000000_.GongEnumValueShapes = append(__GongEnumShape__00000000_.GongEnumValueShapes, __GongEnumValueShape__00000002_)
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000019_)
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000020_)
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000021_)
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000022_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000003_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000004_)
	__GongStructShape__00000001_.LinkShapes = append(__GongStructShape__00000001_.LinkShapes, __LinkShape__00000003_)
	__GongStructShape__00000001_.LinkShapes = append(__GongStructShape__00000001_.LinkShapes, __LinkShape__00000006_)
	__GongStructShape__00000002_.LinkShapes = append(__GongStructShape__00000002_.LinkShapes, __LinkShape__00000000_)
	__GongStructShape__00000004_.AttributeShapes = append(__GongStructShape__00000004_.AttributeShapes, __AttributeShape__00000000_)
	__GongStructShape__00000004_.AttributeShapes = append(__GongStructShape__00000004_.AttributeShapes, __AttributeShape__00000001_)
	__GongStructShape__00000004_.AttributeShapes = append(__GongStructShape__00000004_.AttributeShapes, __AttributeShape__00000002_)
	__GongStructShape__00000006_.AttributeShapes = append(__GongStructShape__00000006_.AttributeShapes, __AttributeShape__00000005_)
	__GongStructShape__00000006_.AttributeShapes = append(__GongStructShape__00000006_.AttributeShapes, __AttributeShape__00000006_)
	__GongStructShape__00000006_.AttributeShapes = append(__GongStructShape__00000006_.AttributeShapes, __AttributeShape__00000007_)
	__GongStructShape__00000006_.AttributeShapes = append(__GongStructShape__00000006_.AttributeShapes, __AttributeShape__00000009_)
	__GongStructShape__00000006_.AttributeShapes = append(__GongStructShape__00000006_.AttributeShapes, __AttributeShape__00000008_)
	__GongStructShape__00000006_.LinkShapes = append(__GongStructShape__00000006_.LinkShapes, __LinkShape__00000004_)
	__GongStructShape__00000006_.LinkShapes = append(__GongStructShape__00000006_.LinkShapes, __LinkShape__00000005_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000010_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000011_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000012_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000013_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000014_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000015_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000016_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000017_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000018_)
}
