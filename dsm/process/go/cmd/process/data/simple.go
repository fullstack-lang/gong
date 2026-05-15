package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/process/go/models"
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

	__Participant__00000000_ := (&models.Participant{Name: `Participant`}).Stage(stage)

	__ParticipantShape__00000000_ := (&models.ParticipantShape{Name: `Participant-DiagramProcess`}).Stage(stage)

	__Process__00000000_ := (&models.Process{Name: `P`}).Stage(stage)

	__ProcessShape__00000000_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `T1`}).Stage(stage)

	__TaskShape__00000000_ := (&models.TaskShape{Name: `-DiagramProcess`}).Stage(stage)

	// insertion point for initialization of values

	__DiagramProcess__00000000_.Name = `DiagramProcess`
	__DiagramProcess__00000000_.Description = ``
	__DiagramProcess__00000000_.ComputedPrefix = ``
	__DiagramProcess__00000000_.IsChecked = true
	__DiagramProcess__00000000_.IsEditable_ = true
	__DiagramProcess__00000000_.IsShowPrefix = false
	__DiagramProcess__00000000_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000000_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000000_.Width = 1114.000000
	__DiagramProcess__00000000_.Height = 1461.000046
	__DiagramProcess__00000000_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000000_.IsParticipantsNodeExpanded = true
	__DiagramProcess__00000000_.IsExternalParticipantsNodeExpanded = false
	__DiagramProcess__00000000_.IsNotesNodeExpanded = false

	__Library__00000000_.Name = ``
	__Library__00000000_.Description = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsSubLibrariesNodeExpanded = false
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.IsProcessesNodeExpanded = true
	__Library__00000000_.IsDataFlowsNodeExpanded = false
	__Library__00000000_.IsDatasNodeExpanded = false
	__Library__00000000_.IsResourcesNodeExpanded = false
	__Library__00000000_.IsNotesNodeExpanded = false
	__Library__00000000_.IsExpandedTmp = true

	__Participant__00000000_.Name = `Participant`
	__Participant__00000000_.IsProcessResource = false
	__Participant__00000000_.Description = ``
	__Participant__00000000_.IsResourcesNodeExpanded = false
	__Participant__00000000_.IsProcessesNodeExpanded = false
	__Participant__00000000_.ComputedPrefix = ``
	__Participant__00000000_.IsTasksNodeExpanded = true
	__Participant__00000000_.IsControlFlowsNodeExpanded = false
	__Participant__00000000_.IsDataFlowsNodeExpanded = false

	__ParticipantShape__00000000_.Name = `Participant-DiagramProcess`
	__ParticipantShape__00000000_.IsExpanded = false
	__ParticipantShape__00000000_.X = 376.000000
	__ParticipantShape__00000000_.Y = 312.000046
	__ParticipantShape__00000000_.Width = 250.000000
	__ParticipantShape__00000000_.Height = 70.000000
	__ParticipantShape__00000000_.IsHidden = false
	__ParticipantShape__00000000_.WidthWeight = 1.000000

	__Process__00000000_.Name = `P`
	__Process__00000000_.Description = ``
	__Process__00000000_.ComputedPrefix = ``
	__Process__00000000_.SVG_Path = ``
	__Process__00000000_.InverseAppliedScaling = 0.000000
	__Process__00000000_.IsSubProcessNodeExpanded = false
	__Process__00000000_.IsDataFlowsNodeExpanded = false

	__ProcessShape__00000000_.Name = `ProcessShape`
	__ProcessShape__00000000_.IsExpanded = false
	__ProcessShape__00000000_.X = 20.000000
	__ProcessShape__00000000_.Y = 73.000046
	__ProcessShape__00000000_.Width = 500.000000
	__ProcessShape__00000000_.Height = 1000.000000
	__ProcessShape__00000000_.IsHidden = false

	__Task__00000000_.Name = `T1`
	__Task__00000000_.Description = ``
	__Task__00000000_.ComputedPrefix = ``
	__Task__00000000_.IsStartTask = false
	__Task__00000000_.IsEndTask = false
	__Task__00000000_.IsTaskNameNotProcessName = false

	__TaskShape__00000000_.Name = `-DiagramProcess`
	__TaskShape__00000000_.IsExpanded = false
	__TaskShape__00000000_.X = 144.000000
	__TaskShape__00000000_.Y = 290.159636
	__TaskShape__00000000_.Width = 250.000000
	__TaskShape__00000000_.Height = 70.000000
	__TaskShape__00000000_.IsHidden = false

	// insertion point for setup of pointers
	__DiagramProcess__00000000_.Process_Shapes = append(__DiagramProcess__00000000_.Process_Shapes, __ProcessShape__00000000_)
	__DiagramProcess__00000000_.Participant_Shapes = append(__DiagramProcess__00000000_.Participant_Shapes, __ParticipantShape__00000000_)
	__DiagramProcess__00000000_.ParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000000_.ParticipantWhoseNodeIsExpanded, __Participant__00000000_)
	__DiagramProcess__00000000_.Task_Shapes = append(__DiagramProcess__00000000_.Task_Shapes, __TaskShape__00000000_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000000_)
	__Library__00000000_.ProcesssWhoseNodeIsExpanded = append(__Library__00000000_.ProcesssWhoseNodeIsExpanded, __Process__00000000_)
	__Participant__00000000_.Tasks = append(__Participant__00000000_.Tasks, __Task__00000000_)
	__ParticipantShape__00000000_.Participant = __Participant__00000000_
	__Process__00000000_.DiagramProcesss = append(__Process__00000000_.DiagramProcesss, __DiagramProcess__00000000_)
	__Process__00000000_.DiagramProcessWhoseNodeIsExpanded = append(__Process__00000000_.DiagramProcessWhoseNodeIsExpanded, __DiagramProcess__00000000_)
	__Process__00000000_.Participants = append(__Process__00000000_.Participants, __Participant__00000000_)
	__ProcessShape__00000000_.Process = __Process__00000000_
	__Task__00000000_.Type = nil
	__TaskShape__00000000_.Task = __Task__00000000_
}
