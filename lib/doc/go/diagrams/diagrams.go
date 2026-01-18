package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
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

	// insertion point for declaration of instances to stage

	__Classdiagram__00000000_ := (&models.Classdiagram{Name: `Default`}).Stage(stage)
	__Classdiagram__00000001_ := (&models.Classdiagram{Name: `spoil diagram`}).Stage(stage)
	__Classdiagram__00000002_ := (&models.Classdiagram{Name: `NoteShape`}).Stage(stage)

	__DiagramPackage__00000000_ := (&models.DiagramPackage{Name: `Diagram Package created the 2025-05-04T22:53:27Z`}).Stage(stage)

	__GongNoteShape__00000000_ := (&models.GongNoteShape{Name: `Default-NoteOnGongdoc`}).Stage(stage)

	__GongStructShape__00000000_ := (&models.GongStructShape{Name: `Default-AttributeShape`}).Stage(stage)
	__GongStructShape__00000001_ := (&models.GongStructShape{Name: `Default-Classdiagram`}).Stage(stage)
	__GongStructShape__00000002_ := (&models.GongStructShape{Name: `NoteShape-GongNoteShape`}).Stage(stage)
	__GongStructShape__00000003_ := (&models.GongStructShape{Name: `NoteShape-GongNoteLinkShape`}).Stage(stage)
	__GongStructShape__00000004_ := (&models.GongStructShape{Name: `Default-GongEnumShape`}).Stage(stage)

	__LinkShape__00000000_ := (&models.LinkShape{Name: `GongNoteLinkShapes`}).Stage(stage)
	__LinkShape__00000001_ := (&models.LinkShape{Name: `GongEnumShapes`}).Stage(stage)

	// insertion point for initialization of values

	__Classdiagram__00000000_.Name = `Default`
	__Classdiagram__00000000_.Description = `This diagram describes the model of the doc package. A DiagramPackage is composed of ClassDiagram and each ClassDiagram has shapes.`
	__Classdiagram__00000000_.IsIncludedInStaticWebSite = true
	__Classdiagram__00000000_.ShowNbInstances = true
	__Classdiagram__00000000_.ShowMultiplicity = true
	__Classdiagram__00000000_.ShowLinkNames = true
	__Classdiagram__00000000_.IsInRenameMode = false
	__Classdiagram__00000000_.IsExpanded = true
	__Classdiagram__00000000_.NodeGongStructsIsExpanded = false
	__Classdiagram__00000000_.NodeGongStructNodeExpansion = `[false,false,false,true,false,false,false,true]`
	__Classdiagram__00000000_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000000_.NodeGongEnumNodeExpansion = `[true]`
	__Classdiagram__00000000_.NodeGongNotesIsExpanded = true
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
	__GongNoteShape__00000000_.X = 907.000000
	__GongNoteShape__00000000_.Y = 157.000000
	__GongNoteShape__00000000_.Width = 416.000000
	__GongNoteShape__00000000_.Height = 93.000000
	__GongNoteShape__00000000_.Matched = false
	__GongNoteShape__00000000_.IsExpanded = false

	__GongStructShape__00000000_.Name = `Default-AttributeShape`
	__GongStructShape__00000000_.X = 577.000000
	__GongStructShape__00000000_.Y = 142.000000
	__GongStructShape__00000000_.IdentifierMeta = ref_models.AttributeShape{}
	__GongStructShape__00000000_.Width = 240.000000
	__GongStructShape__00000000_.Height = 63.000000
	__GongStructShape__00000000_.IsSelected = false

	__GongStructShape__00000001_.Name = `Default-Classdiagram`
	__GongStructShape__00000001_.X = 147.000000
	__GongStructShape__00000001_.Y = 85.000000
	__GongStructShape__00000001_.IdentifierMeta = ref_models.Classdiagram{}
	__GongStructShape__00000001_.Width = 240.000000
	__GongStructShape__00000001_.Height = 63.000000
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
	__GongStructShape__00000004_.X = 127.000000
	__GongStructShape__00000004_.Y = 247.000000
	__GongStructShape__00000004_.IdentifierMeta = ref_models.GongEnumShape{}
	__GongStructShape__00000004_.Width = 240.000000
	__GongStructShape__00000004_.Height = 63.000000
	__GongStructShape__00000004_.IsSelected = false

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

	__LinkShape__00000001_.Name = `GongEnumShapes`
	__LinkShape__00000001_.IdentifierMeta = ref_models.Classdiagram{}.GongEnumShapes
	__LinkShape__00000001_.FieldTypeIdentifierMeta = ref_models.GongEnumShape{}
	__LinkShape__00000001_.FieldOffsetX = 0.000000
	__LinkShape__00000001_.FieldOffsetY = 0.000000
	__LinkShape__00000001_.TargetMultiplicity = models.MANY
	__LinkShape__00000001_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.SourceMultiplicity = models.MANY
	__LinkShape__00000001_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.X = 497.000000
	__LinkShape__00000001_.Y = 197.500000
	__LinkShape__00000001_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.StartRatio = 0.500000
	__LinkShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.EndRatio = 0.500000
	__LinkShape__00000001_.CornerOffsetRatio = 1.640658

	// insertion point for setup of pointers
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000000_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000001_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000004_)
	__Classdiagram__00000000_.GongNoteShapes = append(__Classdiagram__00000000_.GongNoteShapes, __GongNoteShape__00000000_)
	__Classdiagram__00000002_.GongStructShapes = append(__Classdiagram__00000002_.GongStructShapes, __GongStructShape__00000002_)
	__Classdiagram__00000002_.GongStructShapes = append(__Classdiagram__00000002_.GongStructShapes, __GongStructShape__00000003_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000000_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000001_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000002_)
	__DiagramPackage__00000000_.SelectedClassdiagram = __Classdiagram__00000000_
	__GongStructShape__00000001_.LinkShapes = append(__GongStructShape__00000001_.LinkShapes, __LinkShape__00000001_)
	__GongStructShape__00000002_.LinkShapes = append(__GongStructShape__00000002_.LinkShapes, __LinkShape__00000000_)
}
