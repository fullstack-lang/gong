package diagrams

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/dsm/structure/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

// _ point for meta package dummy declaration
var _ ref_models.Stage

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__Classdiagram__00000000_ := (&models.Classdiagram{Name: `01 - Abstract Syntax`}).Stage(stage)

	__DiagramPackage__00000000_ := (&models.DiagramPackage{Name: `Diagram Package created the 2026-06-20T12:07:50Z`}).Stage(stage)

	__GongStructShape__00000000_ := (&models.GongStructShape{Name: `Default-Note`}).Stage(stage)
	__GongStructShape__00000001_ := (&models.GongStructShape{Name: `Default-Library`}).Stage(stage)
	__GongStructShape__00000002_ := (&models.GongStructShape{Name: `Default-System`}).Stage(stage)
	__GongStructShape__00000003_ := (&models.GongStructShape{Name: `Default-Port`}).Stage(stage)
	__GongStructShape__00000004_ := (&models.GongStructShape{Name: `Default-Part`}).Stage(stage)

	__LinkShape__00000000_ := (&models.LinkShape{Name: `RootNotes`}).Stage(stage)
	__LinkShape__00000001_ := (&models.LinkShape{Name: `RootSystemes`}).Stage(stage)
	__LinkShape__00000002_ := (&models.LinkShape{Name: `Parts`}).Stage(stage)
	__LinkShape__00000003_ := (&models.LinkShape{Name: `SubSystemes`}).Stage(stage)
	__LinkShape__00000004_ := (&models.LinkShape{Name: `Ports`}).Stage(stage)

	// insertion point for initialization of values

	__Classdiagram__00000000_.Name = `01 - Abstract Syntax`
	__Classdiagram__00000000_.Description = ``
	__Classdiagram__00000000_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000000_.ShowNbInstances = false
	__Classdiagram__00000000_.ShowMultiplicity = false
	__Classdiagram__00000000_.ShowLinkNames = true
	__Classdiagram__00000000_.IsInRenameMode = false
	__Classdiagram__00000000_.IsExpanded = true
	__Classdiagram__00000000_.NodeGongStructsIsExpanded = true
	__Classdiagram__00000000_.NodeGongStructNodeExpansion = `[false,false,false,false,false,false,false,false,false,false,false,true,false,false,false,false,false,false,false,false,false]`
	__Classdiagram__00000000_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000000_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000000_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000000_.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__00000000_.Name = `Diagram Package created the 2026-06-20T12:07:50Z`
	__DiagramPackage__00000000_.Path = ``
	__DiagramPackage__00000000_.GongModelPath = ``
	__DiagramPackage__00000000_.AbsolutePathToDiagramPackage = ``

	__GongStructShape__00000000_.Name = `Default-Note`
	__GongStructShape__00000000_.X = 443.000000
	__GongStructShape__00000000_.Y = 256.000000
	__GongStructShape__00000000_.IdentifierMeta = ref_models.Note{}
	__GongStructShape__00000000_.Width = 240.000000
	__GongStructShape__00000000_.Height = 63.000000
	__GongStructShape__00000000_.IsSelected = false

	__GongStructShape__00000001_.Name = `Default-Library`
	__GongStructShape__00000001_.X = 37.000000
	__GongStructShape__00000001_.Y = 158.000000
	__GongStructShape__00000001_.IdentifierMeta = ref_models.Library{}
	__GongStructShape__00000001_.Width = 240.000000
	__GongStructShape__00000001_.Height = 63.000000
	__GongStructShape__00000001_.IsSelected = false

	__GongStructShape__00000002_.Name = `Default-System`
	__GongStructShape__00000002_.X = 439.000000
	__GongStructShape__00000002_.Y = 95.000000
	__GongStructShape__00000002_.IdentifierMeta = ref_models.System{}
	__GongStructShape__00000002_.Width = 240.000000
	__GongStructShape__00000002_.Height = 63.000000
	__GongStructShape__00000002_.IsSelected = false

	__GongStructShape__00000003_.Name = `Default-Port`
	__GongStructShape__00000003_.X = 842.000000
	__GongStructShape__00000003_.Y = 253.000000
	__GongStructShape__00000003_.IdentifierMeta = ref_models.Port{}
	__GongStructShape__00000003_.Width = 240.000000
	__GongStructShape__00000003_.Height = 63.000000
	__GongStructShape__00000003_.IsSelected = false

	__GongStructShape__00000004_.Name = `Default-Part`
	__GongStructShape__00000004_.X = 835.000000
	__GongStructShape__00000004_.Y = 92.000000
	__GongStructShape__00000004_.IdentifierMeta = ref_models.Part{}
	__GongStructShape__00000004_.Width = 240.000000
	__GongStructShape__00000004_.Height = 63.000000
	__GongStructShape__00000004_.IsSelected = false

	__LinkShape__00000000_.Name = `RootNotes`
	__LinkShape__00000000_.IdentifierMeta = ref_models.Library{}.RootNotes
	__LinkShape__00000000_.FieldTypeIdentifierMeta = ref_models.Note{}
	__LinkShape__00000000_.FieldOffsetX = 0.000000
	__LinkShape__00000000_.FieldOffsetY = 0.000000
	__LinkShape__00000000_.TargetMultiplicity = models.MANY
	__LinkShape__00000000_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.SourceMultiplicity = models.MANY
	__LinkShape__00000000_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.X = 619.000000
	__LinkShape__00000000_.Y = 193.500000
	__LinkShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.StartRatio = 0.591704
	__LinkShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.EndRatio = 0.500000
	__LinkShape__00000000_.CornerOffsetRatio = 1.188574

	__LinkShape__00000001_.Name = `RootSystemes`
	__LinkShape__00000001_.IdentifierMeta = ref_models.Library{}.RootSystemes
	__LinkShape__00000001_.FieldTypeIdentifierMeta = ref_models.System{}
	__LinkShape__00000001_.FieldOffsetX = 0.000000
	__LinkShape__00000001_.FieldOffsetY = 0.000000
	__LinkShape__00000001_.TargetMultiplicity = models.MANY
	__LinkShape__00000001_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.SourceMultiplicity = models.MANY
	__LinkShape__00000001_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.X = 598.000000
	__LinkShape__00000001_.Y = 158.000000
	__LinkShape__00000001_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.StartRatio = 0.575831
	__LinkShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.EndRatio = 0.500000
	__LinkShape__00000001_.CornerOffsetRatio = 1.184408

	__LinkShape__00000002_.Name = `Parts`
	__LinkShape__00000002_.IdentifierMeta = ref_models.System{}.Parts
	__LinkShape__00000002_.FieldTypeIdentifierMeta = ref_models.Part{}
	__LinkShape__00000002_.FieldOffsetX = 0.000000
	__LinkShape__00000002_.FieldOffsetY = 0.000000
	__LinkShape__00000002_.TargetMultiplicity = models.MANY
	__LinkShape__00000002_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.SourceMultiplicity = models.MANY
	__LinkShape__00000002_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.X = 997.000000
	__LinkShape__00000002_.Y = 125.000000
	__LinkShape__00000002_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.StartRatio = 0.500000
	__LinkShape__00000002_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.EndRatio = 0.500000
	__LinkShape__00000002_.CornerOffsetRatio = 1.380000

	__LinkShape__00000003_.Name = `SubSystemes`
	__LinkShape__00000003_.IdentifierMeta = ref_models.System{}.SubSystemes
	__LinkShape__00000003_.FieldTypeIdentifierMeta = ref_models.System{}
	__LinkShape__00000003_.FieldOffsetX = 0.000000
	__LinkShape__00000003_.FieldOffsetY = 0.000000
	__LinkShape__00000003_.TargetMultiplicity = models.MANY
	__LinkShape__00000003_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.SourceMultiplicity = models.MANY
	__LinkShape__00000003_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.X = 799.000000
	__LinkShape__00000003_.Y = 126.500000
	__LinkShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000003_.StartRatio = 0.092741
	__LinkShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000003_.EndRatio = 0.763574
	__LinkShape__00000003_.CornerOffsetRatio = 1.591704

	__LinkShape__00000004_.Name = `Ports`
	__LinkShape__00000004_.IdentifierMeta = ref_models.Note{}.Ports
	__LinkShape__00000004_.FieldTypeIdentifierMeta = ref_models.Port{}
	__LinkShape__00000004_.FieldOffsetX = 0.000000
	__LinkShape__00000004_.FieldOffsetY = 0.000000
	__LinkShape__00000004_.TargetMultiplicity = models.MANY
	__LinkShape__00000004_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.SourceMultiplicity = models.MANY
	__LinkShape__00000004_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.X = 994.500000
	__LinkShape__00000004_.Y = 386.000000
	__LinkShape__00000004_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000004_.StartRatio = 0.500000
	__LinkShape__00000004_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000004_.EndRatio = 0.500000
	__LinkShape__00000004_.CornerOffsetRatio = 1.380000

	// insertion point for setup of pointers
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000000_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000001_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000002_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000003_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000004_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000000_)
	__DiagramPackage__00000000_.SelectedClassdiagram = __Classdiagram__00000000_
	__GongStructShape__00000000_.LinkShapes = append(__GongStructShape__00000000_.LinkShapes, __LinkShape__00000004_)
	__GongStructShape__00000001_.LinkShapes = append(__GongStructShape__00000001_.LinkShapes, __LinkShape__00000000_)
	__GongStructShape__00000001_.LinkShapes = append(__GongStructShape__00000001_.LinkShapes, __LinkShape__00000001_)
	__GongStructShape__00000002_.LinkShapes = append(__GongStructShape__00000002_.LinkShapes, __LinkShape__00000002_)
	__GongStructShape__00000002_.LinkShapes = append(__GongStructShape__00000002_.LinkShapes, __LinkShape__00000003_)
}
