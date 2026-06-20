package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/structure/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__DiagramProcess__00000000_ := (&models.DiagramProcess{Name: `DiagramProcess`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: ``}).Stage(stage)

	__Participant__00000000_ := (&models.Participant{Name: `P1`}).Stage(stage)
	__Participant__00000001_ := (&models.Participant{Name: `P2`}).Stage(stage)

	__ParticipantShape__00000000_ := (&models.ParticipantShape{Name: `P1-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000001_ := (&models.ParticipantShape{Name: `P2-DiagramProcess`}).Stage(stage)

	__Process__00000000_ := (&models.Process{Name: `System 1`}).Stage(stage)

	__ProcessShape__00000000_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)

	// insertion point for initialization of values

	__DiagramProcess__00000000_.Name = `DiagramProcess`
	__DiagramProcess__00000000_.Description = ``
	__DiagramProcess__00000000_.ComputedPrefix = ``
	__DiagramProcess__00000000_.IsExpanded = false
	__DiagramProcess__00000000_.LayoutDirection = models.Vertical
	__DiagramProcess__00000000_.IsChecked = true
	__DiagramProcess__00000000_.IsEditable_ = true
	__DiagramProcess__00000000_.IsShowPrefix = false
	__DiagramProcess__00000000_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000000_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000000_.Width = 1662.999939
	__DiagramProcess__00000000_.Height = 1350.000000
	__DiagramProcess__00000000_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000000_.IsParticipantsNodeExpanded = true
	__DiagramProcess__00000000_.IsExternalParticipantsNodeExpanded = false
	__DiagramProcess__00000000_.IsNotesNodeExpanded = false

	__Library__00000000_.Name = ``
	__Library__00000000_.Description = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsExpanded = true
	__Library__00000000_.LayoutDirection = models.Vertical
	__Library__00000000_.IsRootLibrary = true
	__Library__00000000_.IsSubLibrariesNodeExpanded = false
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.IsProcessesNodeExpanded = true
	__Library__00000000_.IsDataFlowsNodeExpanded = false
	__Library__00000000_.IsDatasNodeExpanded = false
	__Library__00000000_.IsResourcesNodeExpanded = false
	__Library__00000000_.IsNotesNodeExpanded = false
	__Library__00000000_.IsExpandedTmp = true

	__Participant__00000000_.Name = `P1`
	__Participant__00000000_.IsProcessResource = false
	__Participant__00000000_.Description = ``
	__Participant__00000000_.IsResourcesNodeExpanded = false
	__Participant__00000000_.IsProcessesNodeExpanded = false
	__Participant__00000000_.ComputedPrefix = ``
	__Participant__00000000_.IsExpanded = false
	__Participant__00000000_.LayoutDirection = models.Vertical
	__Participant__00000000_.IsTasksNodeExpanded = false
	__Participant__00000000_.IsControlFlowsNodeExpanded = false
	__Participant__00000000_.IsDataFlowsNodeExpanded = false

	__Participant__00000001_.Name = `P2`
	__Participant__00000001_.IsProcessResource = false
	__Participant__00000001_.Description = ``
	__Participant__00000001_.IsResourcesNodeExpanded = false
	__Participant__00000001_.IsProcessesNodeExpanded = false
	__Participant__00000001_.ComputedPrefix = ``
	__Participant__00000001_.IsExpanded = false
	__Participant__00000001_.LayoutDirection = models.Vertical
	__Participant__00000001_.IsTasksNodeExpanded = false
	__Participant__00000001_.IsControlFlowsNodeExpanded = false
	__Participant__00000001_.IsDataFlowsNodeExpanded = false

	__ParticipantShape__00000000_.Name = `P1-DiagramProcess`
	__ParticipantShape__00000000_.IsExpanded = false
	__ParticipantShape__00000000_.X = 191.494223
	__ParticipantShape__00000000_.Y = 116.217293
	__ParticipantShape__00000000_.Width = 250.000000
	__ParticipantShape__00000000_.Height = 70.000000
	__ParticipantShape__00000000_.IsHidden = false
	__ParticipantShape__00000000_.WidthWeight = 1.000000

	__ParticipantShape__00000001_.Name = `P2-DiagramProcess`
	__ParticipantShape__00000001_.IsExpanded = false
	__ParticipantShape__00000001_.X = 177.408274
	__ParticipantShape__00000001_.Y = 166.681744
	__ParticipantShape__00000001_.Width = 250.000000
	__ParticipantShape__00000001_.Height = 70.000000
	__ParticipantShape__00000001_.IsHidden = false
	__ParticipantShape__00000001_.WidthWeight = 1.000000

	__Process__00000000_.Name = `System 1`
	__Process__00000000_.Description = ``
	__Process__00000000_.ComputedPrefix = ``
	__Process__00000000_.IsExpanded = false
	__Process__00000000_.LayoutDirection = models.Vertical
	__Process__00000000_.SVG_Path = ``
	__Process__00000000_.InverseAppliedScaling = 0.000000
	__Process__00000000_.IsSubProcessNodeExpanded = false
	__Process__00000000_.IsDataFlowsNodeExpanded = false

	__ProcessShape__00000000_.Name = `ProcessShape`
	__ProcessShape__00000000_.IsExpanded = false
	__ProcessShape__00000000_.X = 100.000000
	__ProcessShape__00000000_.Y = 50.000000
	__ProcessShape__00000000_.Width = 1262.999939
	__ProcessShape__00000000_.Height = 1000.000000
	__ProcessShape__00000000_.IsHidden = false

	// insertion point for setup of pointers
	__DiagramProcess__00000000_.Process_Shapes = append(__DiagramProcess__00000000_.Process_Shapes, __ProcessShape__00000000_)
	__DiagramProcess__00000000_.Participant_Shapes = append(__DiagramProcess__00000000_.Participant_Shapes, __ParticipantShape__00000000_)
	__DiagramProcess__00000000_.Participant_Shapes = append(__DiagramProcess__00000000_.Participant_Shapes, __ParticipantShape__00000001_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000000_)
	__Library__00000000_.ProcesssWhoseNodeIsExpanded = append(__Library__00000000_.ProcesssWhoseNodeIsExpanded, __Process__00000000_)
	__ParticipantShape__00000000_.Participant = __Participant__00000000_
	__ParticipantShape__00000001_.Participant = __Participant__00000001_
	__Process__00000000_.DiagramProcesss = append(__Process__00000000_.DiagramProcesss, __DiagramProcess__00000000_)
	__Process__00000000_.DiagramProcessWhoseNodeIsExpanded = append(__Process__00000000_.DiagramProcessWhoseNodeIsExpanded, __DiagramProcess__00000000_)
	__Process__00000000_.Participants = append(__Process__00000000_.Participants, __Participant__00000000_)
	__Process__00000000_.Participants = append(__Process__00000000_.Participants, __Participant__00000001_)
	__ProcessShape__00000000_.Process = __Process__00000000_
}
