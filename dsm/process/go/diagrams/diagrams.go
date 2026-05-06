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
	__GongStructShape__00000003_ := (&models.GongStructShape{Name: `Default-ControlFlow`}).Stage(stage)
	__GongStructShape__00000004_ := (&models.GongStructShape{Name: `Default-Library`}).Stage(stage)
	__GongStructShape__00000005_ := (&models.GongStructShape{Name: `Default-DataFlow`}).Stage(stage)
	__GongStructShape__00000006_ := (&models.GongStructShape{Name: `Default-Data`}).Stage(stage)

	__LinkShape__00000000_ := (&models.LinkShape{Name: `Participants`}).Stage(stage)
	__LinkShape__00000001_ := (&models.LinkShape{Name: `Tasks`}).Stage(stage)
	__LinkShape__00000002_ := (&models.LinkShape{Name: `Start`}).Stage(stage)
	__LinkShape__00000003_ := (&models.LinkShape{Name: `End`}).Stage(stage)
	__LinkShape__00000004_ := (&models.LinkShape{Name: `SubLibraries`}).Stage(stage)
	__LinkShape__00000005_ := (&models.LinkShape{Name: `RootProcesses`}).Stage(stage)
	__LinkShape__00000006_ := (&models.LinkShape{Name: `SubProcesses`}).Stage(stage)
	__LinkShape__00000007_ := (&models.LinkShape{Name: `Start`}).Stage(stage)
	__LinkShape__00000008_ := (&models.LinkShape{Name: `End`}).Stage(stage)
	__LinkShape__00000009_ := (&models.LinkShape{Name: `Datas`}).Stage(stage)
	__LinkShape__00000010_ := (&models.LinkShape{Name: `RootDatas`}).Stage(stage)
	__LinkShape__00000011_ := (&models.LinkShape{Name: `RootDataFlows`}).Stage(stage)

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
	__Classdiagram__00000000_.NodeGongStructNodeExpansion = `[true,false,true,true,true,false,true,true]`
	__Classdiagram__00000000_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000000_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000000_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000000_.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__00000000_.Name = `Diagram Package created the 2026-04-19T14:58:45Z`
	__DiagramPackage__00000000_.Path = ``
	__DiagramPackage__00000000_.GongModelPath = ``
	__DiagramPackage__00000000_.AbsolutePathToDiagramPackage = ``

	__GongStructShape__00000000_.Name = `Default-Process`
	__GongStructShape__00000000_.X = 716.000000
	__GongStructShape__00000000_.Y = 111.000000
	__GongStructShape__00000000_.IdentifierMeta = ref_models.Process{}
	__GongStructShape__00000000_.Width = 240.000000
	__GongStructShape__00000000_.Height = 63.000000
	__GongStructShape__00000000_.IsSelected = false

	__GongStructShape__00000001_.Name = `Default-Participant`
	__GongStructShape__00000001_.X = 720.000000
	__GongStructShape__00000001_.Y = 264.000000
	__GongStructShape__00000001_.IdentifierMeta = ref_models.Participant{}
	__GongStructShape__00000001_.Width = 240.000000
	__GongStructShape__00000001_.Height = 63.000000
	__GongStructShape__00000001_.IsSelected = false

	__GongStructShape__00000002_.Name = `Default-Task`
	__GongStructShape__00000002_.X = 718.000000
	__GongStructShape__00000002_.Y = 450.000000
	__GongStructShape__00000002_.IdentifierMeta = ref_models.Task{}
	__GongStructShape__00000002_.Width = 240.000000
	__GongStructShape__00000002_.Height = 106.000000
	__GongStructShape__00000002_.IsSelected = false

	__GongStructShape__00000003_.Name = `Default-ControlFlow`
	__GongStructShape__00000003_.X = 1121.000000
	__GongStructShape__00000003_.Y = 448.000000
	__GongStructShape__00000003_.IdentifierMeta = ref_models.ControlFlow{}
	__GongStructShape__00000003_.Width = 240.000000
	__GongStructShape__00000003_.Height = 113.000000
	__GongStructShape__00000003_.IsSelected = false

	__GongStructShape__00000004_.Name = `Default-Library`
	__GongStructShape__00000004_.X = 60.000000
	__GongStructShape__00000004_.Y = 107.000000
	__GongStructShape__00000004_.IdentifierMeta = ref_models.Library{}
	__GongStructShape__00000004_.Width = 469.000000
	__GongStructShape__00000004_.Height = 63.000000
	__GongStructShape__00000004_.IsSelected = false

	__GongStructShape__00000005_.Name = `Default-DataFlow`
	__GongStructShape__00000005_.X = 325.000000
	__GongStructShape__00000005_.Y = 451.000000
	__GongStructShape__00000005_.IdentifierMeta = ref_models.DataFlow{}
	__GongStructShape__00000005_.Width = 240.000000
	__GongStructShape__00000005_.Height = 106.000000
	__GongStructShape__00000005_.IsSelected = false

	__GongStructShape__00000006_.Name = `Default-Data`
	__GongStructShape__00000006_.X = 51.000000
	__GongStructShape__00000006_.Y = 464.000000
	__GongStructShape__00000006_.IdentifierMeta = ref_models.Data{}
	__GongStructShape__00000006_.Width = 151.000000
	__GongStructShape__00000006_.Height = 63.000000
	__GongStructShape__00000006_.IsSelected = false

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

	__LinkShape__00000002_.Name = `Start`
	__LinkShape__00000002_.IdentifierMeta = ref_models.ControlFlow{}.Start
	__LinkShape__00000002_.FieldTypeIdentifierMeta = ref_models.Task{}
	__LinkShape__00000002_.FieldOffsetX = 0.000000
	__LinkShape__00000002_.FieldOffsetY = 0.000000
	__LinkShape__00000002_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000002_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.SourceMultiplicity = models.MANY
	__LinkShape__00000002_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.X = 776.500000
	__LinkShape__00000002_.Y = 485.000000
	__LinkShape__00000002_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.StartRatio = 0.097529
	__LinkShape__00000002_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.EndRatio = 0.000000
	__LinkShape__00000002_.CornerOffsetRatio = -0.258594

	__LinkShape__00000003_.Name = `End`
	__LinkShape__00000003_.IdentifierMeta = ref_models.ControlFlow{}.End
	__LinkShape__00000003_.FieldTypeIdentifierMeta = ref_models.Task{}
	__LinkShape__00000003_.FieldOffsetX = 0.000000
	__LinkShape__00000003_.FieldOffsetY = 0.000000
	__LinkShape__00000003_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000003_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.SourceMultiplicity = models.MANY
	__LinkShape__00000003_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.X = 776.500000
	__LinkShape__00000003_.Y = 485.000000
	__LinkShape__00000003_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000003_.StartRatio = 0.840892
	__LinkShape__00000003_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000003_.EndRatio = 1.000000
	__LinkShape__00000003_.CornerOffsetRatio = -0.246094

	__LinkShape__00000004_.Name = `SubLibraries`
	__LinkShape__00000004_.IdentifierMeta = ref_models.Library{}.SubLibraries
	__LinkShape__00000004_.FieldTypeIdentifierMeta = ref_models.Library{}
	__LinkShape__00000004_.FieldOffsetX = 0.000000
	__LinkShape__00000004_.FieldOffsetY = 0.000000
	__LinkShape__00000004_.TargetMultiplicity = models.MANY
	__LinkShape__00000004_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.SourceMultiplicity = models.MANY
	__LinkShape__00000004_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.X = 425.000000
	__LinkShape__00000004_.Y = 59.500000
	__LinkShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000004_.StartRatio = 0.209505
	__LinkShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000004_.EndRatio = 0.851296
	__LinkShape__00000004_.CornerOffsetRatio = -0.968161

	__LinkShape__00000005_.Name = `RootProcesses`
	__LinkShape__00000005_.IdentifierMeta = ref_models.Library{}.RootProcesses
	__LinkShape__00000005_.FieldTypeIdentifierMeta = ref_models.Process{}
	__LinkShape__00000005_.FieldOffsetX = 0.000000
	__LinkShape__00000005_.FieldOffsetY = 0.000000
	__LinkShape__00000005_.TargetMultiplicity = models.MANY
	__LinkShape__00000005_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.SourceMultiplicity = models.MANY
	__LinkShape__00000005_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.X = 737.000000
	__LinkShape__00000005_.Y = 143.000000
	__LinkShape__00000005_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000005_.StartRatio = 0.571522
	__LinkShape__00000005_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000005_.EndRatio = 0.500000
	__LinkShape__00000005_.CornerOffsetRatio = 1.380000

	__LinkShape__00000006_.Name = `SubProcesses`
	__LinkShape__00000006_.IdentifierMeta = ref_models.Process{}.SubProcesses
	__LinkShape__00000006_.FieldTypeIdentifierMeta = ref_models.Process{}
	__LinkShape__00000006_.FieldOffsetX = 0.000000
	__LinkShape__00000006_.FieldOffsetY = 0.000000
	__LinkShape__00000006_.TargetMultiplicity = models.MANY
	__LinkShape__00000006_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.SourceMultiplicity = models.MANY
	__LinkShape__00000006_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.X = 1028.000000
	__LinkShape__00000006_.Y = 144.500000
	__LinkShape__00000006_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000006_.StartRatio = 0.259408
	__LinkShape__00000006_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000006_.EndRatio = 0.888628
	__LinkShape__00000006_.CornerOffsetRatio = -1.047619

	__LinkShape__00000007_.Name = `Start`
	__LinkShape__00000007_.IdentifierMeta = ref_models.DataFlow{}.StartTask
	__LinkShape__00000007_.FieldTypeIdentifierMeta = ref_models.Task{}
	__LinkShape__00000007_.FieldOffsetX = 0.000000
	__LinkShape__00000007_.FieldOffsetY = 0.000000
	__LinkShape__00000007_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000007_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000007_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000007_.SourceMultiplicity = models.MANY
	__LinkShape__00000007_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000007_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000007_.X = 845.500000
	__LinkShape__00000007_.Y = 506.000000
	__LinkShape__00000007_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000007_.StartRatio = 0.132272
	__LinkShape__00000007_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000007_.EndRatio = 0.170008
	__LinkShape__00000007_.CornerOffsetRatio = 1.380000

	__LinkShape__00000008_.Name = `End`
	__LinkShape__00000008_.IdentifierMeta = ref_models.DataFlow{}.EndTask
	__LinkShape__00000008_.FieldTypeIdentifierMeta = ref_models.Task{}
	__LinkShape__00000008_.FieldOffsetX = 0.000000
	__LinkShape__00000008_.FieldOffsetY = 0.000000
	__LinkShape__00000008_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000008_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000008_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000008_.SourceMultiplicity = models.MANY
	__LinkShape__00000008_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000008_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000008_.X = 845.500000
	__LinkShape__00000008_.Y = 506.000000
	__LinkShape__00000008_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000008_.StartRatio = 0.802083
	__LinkShape__00000008_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000008_.EndRatio = 0.745479
	__LinkShape__00000008_.CornerOffsetRatio = 1.380000

	__LinkShape__00000009_.Name = `Datas`
	__LinkShape__00000009_.IdentifierMeta = ref_models.DataFlow{}.Datas
	__LinkShape__00000009_.FieldTypeIdentifierMeta = ref_models.Data{}
	__LinkShape__00000009_.FieldOffsetX = 0.000000
	__LinkShape__00000009_.FieldOffsetY = 0.000000
	__LinkShape__00000009_.TargetMultiplicity = models.MANY
	__LinkShape__00000009_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000009_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000009_.SourceMultiplicity = models.MANY
	__LinkShape__00000009_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000009_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000009_.X = 529.000000
	__LinkShape__00000009_.Y = 513.500000
	__LinkShape__00000009_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000009_.StartRatio = 0.500000
	__LinkShape__00000009_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000009_.EndRatio = 0.738188
	__LinkShape__00000009_.CornerOffsetRatio = -0.211426

	__LinkShape__00000010_.Name = `RootDatas`
	__LinkShape__00000010_.IdentifierMeta = ref_models.Library{}.RootDatas
	__LinkShape__00000010_.FieldTypeIdentifierMeta = ref_models.Data{}
	__LinkShape__00000010_.FieldOffsetX = 0.000000
	__LinkShape__00000010_.FieldOffsetY = 0.000000
	__LinkShape__00000010_.TargetMultiplicity = models.MANY
	__LinkShape__00000010_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000010_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000010_.SourceMultiplicity = models.MANY
	__LinkShape__00000010_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000010_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000010_.X = 401.500000
	__LinkShape__00000010_.Y = 320.000000
	__LinkShape__00000010_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000010_.StartRatio = 0.222298
	__LinkShape__00000010_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000010_.EndRatio = 0.703694
	__LinkShape__00000010_.CornerOffsetRatio = 2.095331

	__LinkShape__00000011_.Name = `RootDataFlows`
	__LinkShape__00000011_.IdentifierMeta = ref_models.Library{}.RootDataFlows
	__LinkShape__00000011_.FieldTypeIdentifierMeta = ref_models.DataFlow{}
	__LinkShape__00000011_.FieldOffsetX = 0.000000
	__LinkShape__00000011_.FieldOffsetY = 0.000000
	__LinkShape__00000011_.TargetMultiplicity = models.MANY
	__LinkShape__00000011_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000011_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000011_.SourceMultiplicity = models.MANY
	__LinkShape__00000011_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000011_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000011_.X = 908.000000
	__LinkShape__00000011_.Y = 311.000000
	__LinkShape__00000011_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000011_.StartRatio = 0.842767
	__LinkShape__00000011_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000011_.EndRatio = 0.492741
	__LinkShape__00000011_.CornerOffsetRatio = 1.825490

	// insertion point for setup of pointers
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000000_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000001_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000002_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000003_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000004_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000005_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000006_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000000_)
	__DiagramPackage__00000000_.SelectedClassdiagram = __Classdiagram__00000000_
	__GongStructShape__00000000_.LinkShapes = append(__GongStructShape__00000000_.LinkShapes, __LinkShape__00000000_)
	__GongStructShape__00000000_.LinkShapes = append(__GongStructShape__00000000_.LinkShapes, __LinkShape__00000006_)
	__GongStructShape__00000001_.LinkShapes = append(__GongStructShape__00000001_.LinkShapes, __LinkShape__00000001_)
	__GongStructShape__00000003_.LinkShapes = append(__GongStructShape__00000003_.LinkShapes, __LinkShape__00000002_)
	__GongStructShape__00000003_.LinkShapes = append(__GongStructShape__00000003_.LinkShapes, __LinkShape__00000003_)
	__GongStructShape__00000004_.LinkShapes = append(__GongStructShape__00000004_.LinkShapes, __LinkShape__00000004_)
	__GongStructShape__00000004_.LinkShapes = append(__GongStructShape__00000004_.LinkShapes, __LinkShape__00000005_)
	__GongStructShape__00000004_.LinkShapes = append(__GongStructShape__00000004_.LinkShapes, __LinkShape__00000010_)
	__GongStructShape__00000004_.LinkShapes = append(__GongStructShape__00000004_.LinkShapes, __LinkShape__00000011_)
	__GongStructShape__00000005_.LinkShapes = append(__GongStructShape__00000005_.LinkShapes, __LinkShape__00000007_)
	__GongStructShape__00000005_.LinkShapes = append(__GongStructShape__00000005_.LinkShapes, __LinkShape__00000008_)
	__GongStructShape__00000005_.LinkShapes = append(__GongStructShape__00000005_.LinkShapes, __LinkShape__00000009_)
}
