package diagrams

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/dsm/process/go/models"
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

	__Classdiagram__00000000_ := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	__DiagramPackage__00000000_ := (&models.DiagramPackage{Name: `Diagram Package created the 2026-04-19T14:58:45Z`}).Stage(stage)

	__GongStructShape__00000000_ := (&models.GongStructShape{Name: `Default-Process`}).Stage(stage)
	__GongStructShape__00000001_ := (&models.GongStructShape{Name: `Default-Participant`}).Stage(stage)
	__GongStructShape__00000002_ := (&models.GongStructShape{Name: `Default-Task`}).Stage(stage)

	__LinkShape__00000000_ := (&models.LinkShape{Name: `Participants`}).Stage(stage)
	__LinkShape__00000001_ := (&models.LinkShape{Name: `Tasks`}).Stage(stage)

	// insertion point for initialization of values

	__Classdiagram__00000000_.Name = `Default`
	__Classdiagram__00000000_.Description = ``
	__Classdiagram__00000000_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000000_.ShowNbInstances = true
	__Classdiagram__00000000_.ShowMultiplicity = true
	__Classdiagram__00000000_.ShowLinkNames = true
	__Classdiagram__00000000_.IsInRenameMode = false
	__Classdiagram__00000000_.IsExpanded = true
	__Classdiagram__00000000_.NodeGongStructsIsExpanded = true
	__Classdiagram__00000000_.NodeGongStructNodeExpansion = `[false,false,true,false,true]`
	__Classdiagram__00000000_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000000_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000000_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000000_.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__00000000_.Name = `Diagram Package created the 2026-04-19T14:58:45Z`
	__DiagramPackage__00000000_.Path = ``
	__DiagramPackage__00000000_.GongModelPath = ``
	__DiagramPackage__00000000_.AbsolutePathToDiagramPackage = ``

	__GongStructShape__00000000_.Name = `Default-Process`
	__GongStructShape__00000000_.X = 93.000000
	__GongStructShape__00000000_.Y = 102.000000
	__GongStructShape__00000000_.IdentifierMeta = ref_models.Process{}
	__GongStructShape__00000000_.Width = 240.000000
	__GongStructShape__00000000_.Height = 63.000000
	__GongStructShape__00000000_.IsSelected = false

	__GongStructShape__00000001_.Name = `Default-Participant`
	__GongStructShape__00000001_.X = 97.000000
	__GongStructShape__00000001_.Y = 255.000000
	__GongStructShape__00000001_.IdentifierMeta = ref_models.Participant{}
	__GongStructShape__00000001_.Width = 240.000000
	__GongStructShape__00000001_.Height = 63.000000
	__GongStructShape__00000001_.IsSelected = false

	__GongStructShape__00000002_.Name = `Default-Task`
	__GongStructShape__00000002_.X = 96.000000
	__GongStructShape__00000002_.Y = 441.000000
	__GongStructShape__00000002_.IdentifierMeta = ref_models.Task{}
	__GongStructShape__00000002_.Width = 240.000000
	__GongStructShape__00000002_.Height = 63.000000
	__GongStructShape__00000002_.IsSelected = false

	__LinkShape__00000000_.Name = `Participants`
	__LinkShape__00000000_.IdentifierMeta = ref_models.Process{}.Participants
	__LinkShape__00000000_.FieldTypeIdentifierMeta = ref_models.Participant{}
	__LinkShape__00000000_.FieldOffsetX = 0.000000
	__LinkShape__00000000_.FieldOffsetY = 0.000000
	__LinkShape__00000000_.TargetMultiplicity = models.MANY
	__LinkShape__00000000_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.SourceMultiplicity = models.MANY
	__LinkShape__00000000_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.X = 455.000000
	__LinkShape__00000000_.Y = 134.500000
	__LinkShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000000_.StartRatio = 0.266102
	__LinkShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000000_.EndRatio = 0.249436
	__LinkShape__00000000_.CornerOffsetRatio = 1.476356

	__LinkShape__00000001_.Name = `Tasks`
	__LinkShape__00000001_.IdentifierMeta = ref_models.Participant{}.Tasks
	__LinkShape__00000001_.FieldTypeIdentifierMeta = ref_models.Task{}
	__LinkShape__00000001_.FieldOffsetX = 0.000000
	__LinkShape__00000001_.FieldOffsetY = 0.000000
	__LinkShape__00000001_.TargetMultiplicity = models.MANY
	__LinkShape__00000001_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.SourceMultiplicity = models.MANY
	__LinkShape__00000001_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.X = 421.000000
	__LinkShape__00000001_.Y = 179.500000
	__LinkShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000001_.StartRatio = 0.241102
	__LinkShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000001_.EndRatio = 0.182769
	__LinkShape__00000001_.CornerOffsetRatio = 1.396991

	// insertion point for setup of pointers
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000000_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000001_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000002_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000000_)
	__DiagramPackage__00000000_.SelectedClassdiagram = __Classdiagram__00000000_
	__GongStructShape__00000000_.LinkShapes = append(__GongStructShape__00000000_.LinkShapes, __LinkShape__00000000_)
	__GongStructShape__00000001_.LinkShapes = append(__GongStructShape__00000001_.LinkShapes, __LinkShape__00000001_)
}
