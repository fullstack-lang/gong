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

	__AllocatedResourceShape__00000000_ := (&models.AllocatedResourceShape{Name: `DiagramProcess-Originator-Aircraft Operator`}).Stage(stage)
	__AllocatedResourceShape__00000001_ := (&models.AllocatedResourceShape{Name: `DiagramProcess-Originator-Agency`}).Stage(stage)
	__AllocatedResourceShape__00000002_ := (&models.AllocatedResourceShape{Name: `DiagramProcess-Bureau des pistes (BIVC)-BIVC`}).Stage(stage)
	__AllocatedResourceShape__00000003_ := (&models.AllocatedResourceShape{Name: `DiagramProcess-IFPS-Integrated Initial Flight Plan Processing`}).Stage(stage)

	__DiagramProcess__00000000_ := (&models.DiagramProcess{Name: `DiagramProcess`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `IFPS User Manual`}).Stage(stage)

	__Participant__00000000_ := (&models.Participant{Name: `Originator`}).Stage(stage)
	__Participant__00000001_ := (&models.Participant{Name: `ARO`}).Stage(stage)
	__Participant__00000002_ := (&models.Participant{Name: `IFPS`}).Stage(stage)
	__Participant__00000003_ := (&models.Participant{Name: `ATCs`}).Stage(stage)

	__ParticipantShape__00000000_ := (&models.ParticipantShape{Name: `-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000001_ := (&models.ParticipantShape{Name: `ARO-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000002_ := (&models.ParticipantShape{Name: `IFPS-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000003_ := (&models.ParticipantShape{Name: `ATCs-DiagramProcess`}).Stage(stage)

	__Process__00000000_ := (&models.Process{Name: `Flight Data Message Flows`}).Stage(stage)

	__ProcessShape__00000000_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)

	__Resource__00000000_ := (&models.Resource{Name: `Aircraft Operator`}).Stage(stage)
	__Resource__00000001_ := (&models.Resource{Name: `Agency`}).Stage(stage)
	__Resource__00000002_ := (&models.Resource{Name: `Integrated Initial Flight Plan Processing`}).Stage(stage)
	__Resource__00000003_ := (&models.Resource{Name: `ARO`}).Stage(stage)
	__Resource__00000004_ := (&models.Resource{Name: `ATCU 1`}).Stage(stage)
	__Resource__00000005_ := (&models.Resource{Name: `ATCU 2`}).Stage(stage)
	__Resource__00000006_ := (&models.Resource{Name: `Enhanced Tactical Flow Management System`}).Stage(stage)
	__Resource__00000007_ := (&models.Resource{Name: `BIVC`}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `Submit PLN via ARO`}).Stage(stage)

	__TaskShape__00000000_ := (&models.TaskShape{Name: `-DiagramProcess`}).Stage(stage)

	// insertion point for initialization of values

	__AllocatedResourceShape__00000000_.Name = `DiagramProcess-Originator-Aircraft Operator`

	__AllocatedResourceShape__00000001_.Name = `DiagramProcess-Originator-Agency`

	__AllocatedResourceShape__00000002_.Name = `DiagramProcess-Bureau des pistes (BIVC)-BIVC`

	__AllocatedResourceShape__00000003_.Name = `DiagramProcess-IFPS-Integrated Initial Flight Plan Processing`

	__DiagramProcess__00000000_.Name = `DiagramProcess`
	__DiagramProcess__00000000_.Description = ``
	__DiagramProcess__00000000_.ComputedPrefix = ``
	__DiagramProcess__00000000_.IsChecked = true
	__DiagramProcess__00000000_.IsEditable_ = true
	__DiagramProcess__00000000_.IsShowPrefix = false
	__DiagramProcess__00000000_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000000_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000000_.Width = 1424.000000
	__DiagramProcess__00000000_.Height = 1350.000000
	__DiagramProcess__00000000_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000000_.IsParticipantsNodeExpanded = true
	__DiagramProcess__00000000_.IsExternalParticipantsNodeExpanded = false
	__DiagramProcess__00000000_.IsNotesNodeExpanded = false

	__Library__00000000_.Name = `IFPS User Manual`
	__Library__00000000_.Description = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsRootLibrary = true
	__Library__00000000_.IsSubLibrariesNodeExpanded = false
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.IsProcessesNodeExpanded = true
	__Library__00000000_.IsDataFlowsNodeExpanded = false
	__Library__00000000_.IsDatasNodeExpanded = false
	__Library__00000000_.IsResourcesNodeExpanded = true
	__Library__00000000_.IsNotesNodeExpanded = false
	__Library__00000000_.IsExpandedTmp = true

	__Participant__00000000_.Name = `Originator`
	__Participant__00000000_.IsProcessResource = false
	__Participant__00000000_.Description = ``
	__Participant__00000000_.IsResourcesNodeExpanded = false
	__Participant__00000000_.IsProcessesNodeExpanded = false
	__Participant__00000000_.ComputedPrefix = ``
	__Participant__00000000_.IsTasksNodeExpanded = true
	__Participant__00000000_.IsControlFlowsNodeExpanded = false
	__Participant__00000000_.IsDataFlowsNodeExpanded = false

	__Participant__00000001_.Name = `ARO`
	__Participant__00000001_.IsProcessResource = false
	__Participant__00000001_.Description = ``
	__Participant__00000001_.IsResourcesNodeExpanded = true
	__Participant__00000001_.IsProcessesNodeExpanded = false
	__Participant__00000001_.ComputedPrefix = ``
	__Participant__00000001_.IsTasksNodeExpanded = false
	__Participant__00000001_.IsControlFlowsNodeExpanded = false
	__Participant__00000001_.IsDataFlowsNodeExpanded = false

	__Participant__00000002_.Name = `IFPS`
	__Participant__00000002_.IsProcessResource = false
	__Participant__00000002_.Description = ``
	__Participant__00000002_.IsResourcesNodeExpanded = true
	__Participant__00000002_.IsProcessesNodeExpanded = false
	__Participant__00000002_.ComputedPrefix = ``
	__Participant__00000002_.IsTasksNodeExpanded = false
	__Participant__00000002_.IsControlFlowsNodeExpanded = false
	__Participant__00000002_.IsDataFlowsNodeExpanded = false

	__Participant__00000003_.Name = `ATCs`
	__Participant__00000003_.IsProcessResource = false
	__Participant__00000003_.Description = ``
	__Participant__00000003_.IsResourcesNodeExpanded = false
	__Participant__00000003_.IsProcessesNodeExpanded = false
	__Participant__00000003_.ComputedPrefix = ``
	__Participant__00000003_.IsTasksNodeExpanded = false
	__Participant__00000003_.IsControlFlowsNodeExpanded = false
	__Participant__00000003_.IsDataFlowsNodeExpanded = false

	__ParticipantShape__00000000_.Name = `-DiagramProcess`
	__ParticipantShape__00000000_.IsExpanded = false
	__ParticipantShape__00000000_.X = 170.882152
	__ParticipantShape__00000000_.Y = 150.684410
	__ParticipantShape__00000000_.Width = 250.000000
	__ParticipantShape__00000000_.Height = 70.000000
	__ParticipantShape__00000000_.IsHidden = false
	__ParticipantShape__00000000_.WidthWeight = 1.000000

	__ParticipantShape__00000001_.Name = `ARO-DiagramProcess`
	__ParticipantShape__00000001_.IsExpanded = false
	__ParticipantShape__00000001_.X = 174.345500
	__ParticipantShape__00000001_.Y = 129.819672
	__ParticipantShape__00000001_.Width = 250.000000
	__ParticipantShape__00000001_.Height = 70.000000
	__ParticipantShape__00000001_.IsHidden = false
	__ParticipantShape__00000001_.WidthWeight = 1.000000

	__ParticipantShape__00000002_.Name = `IFPS-DiagramProcess`
	__ParticipantShape__00000002_.IsExpanded = false
	__ParticipantShape__00000002_.X = 130.947851
	__ParticipantShape__00000002_.Y = 127.801982
	__ParticipantShape__00000002_.Width = 250.000000
	__ParticipantShape__00000002_.Height = 70.000000
	__ParticipantShape__00000002_.IsHidden = false
	__ParticipantShape__00000002_.WidthWeight = 1.000000

	__ParticipantShape__00000003_.Name = `ATCs-DiagramProcess`
	__ParticipantShape__00000003_.IsExpanded = false
	__ParticipantShape__00000003_.X = 105.352951
	__ParticipantShape__00000003_.Y = 115.067303
	__ParticipantShape__00000003_.Width = 250.000000
	__ParticipantShape__00000003_.Height = 70.000000
	__ParticipantShape__00000003_.IsHidden = false
	__ParticipantShape__00000003_.WidthWeight = 1.000000

	__Process__00000000_.Name = `Flight Data Message Flows`
	__Process__00000000_.Description = ``
	__Process__00000000_.ComputedPrefix = ``
	__Process__00000000_.SVG_Path = ``
	__Process__00000000_.InverseAppliedScaling = 0.000000
	__Process__00000000_.IsSubProcessNodeExpanded = false
	__Process__00000000_.IsDataFlowsNodeExpanded = false

	__ProcessShape__00000000_.Name = `ProcessShape`
	__ProcessShape__00000000_.IsExpanded = false
	__ProcessShape__00000000_.X = 100.000000
	__ProcessShape__00000000_.Y = 50.000000
	__ProcessShape__00000000_.Width = 1024.000000
	__ProcessShape__00000000_.Height = 1000.000000
	__ProcessShape__00000000_.IsHidden = false

	__Resource__00000000_.Name = `Aircraft Operator`
	__Resource__00000000_.Acronym = `AO`
	__Resource__00000000_.Description = ``
	__Resource__00000000_.ComputedPrefix = ``
	__Resource__00000000_.SVG_Path = ``
	__Resource__00000000_.InverseAppliedScaling = 0.000000

	__Resource__00000001_.Name = `Agency`
	__Resource__00000001_.Acronym = ``
	__Resource__00000001_.Description = ``
	__Resource__00000001_.ComputedPrefix = ``
	__Resource__00000001_.SVG_Path = ``
	__Resource__00000001_.InverseAppliedScaling = 0.000000

	__Resource__00000002_.Name = `Integrated Initial Flight Plan Processing`
	__Resource__00000002_.Acronym = `IFPS`
	__Resource__00000002_.Description = ``
	__Resource__00000002_.ComputedPrefix = ``
	__Resource__00000002_.SVG_Path = ``
	__Resource__00000002_.InverseAppliedScaling = 0.000000

	__Resource__00000003_.Name = `ARO`
	__Resource__00000003_.Acronym = ``
	__Resource__00000003_.Description = ``
	__Resource__00000003_.ComputedPrefix = ``
	__Resource__00000003_.SVG_Path = ``
	__Resource__00000003_.InverseAppliedScaling = 0.000000

	__Resource__00000004_.Name = `ATCU 1`
	__Resource__00000004_.Acronym = ``
	__Resource__00000004_.Description = ``
	__Resource__00000004_.ComputedPrefix = ``
	__Resource__00000004_.SVG_Path = ``
	__Resource__00000004_.InverseAppliedScaling = 0.000000

	__Resource__00000005_.Name = `ATCU 2`
	__Resource__00000005_.Acronym = ``
	__Resource__00000005_.Description = ``
	__Resource__00000005_.ComputedPrefix = ``
	__Resource__00000005_.SVG_Path = ``
	__Resource__00000005_.InverseAppliedScaling = 0.000000

	__Resource__00000006_.Name = `Enhanced Tactical Flow Management System`
	__Resource__00000006_.Acronym = `ETFMS`
	__Resource__00000006_.Description = ``
	__Resource__00000006_.ComputedPrefix = ``
	__Resource__00000006_.SVG_Path = ``
	__Resource__00000006_.InverseAppliedScaling = 0.000000

	__Resource__00000007_.Name = `BIVC`
	__Resource__00000007_.Acronym = ``
	__Resource__00000007_.Description = ``
	__Resource__00000007_.ComputedPrefix = ``
	__Resource__00000007_.SVG_Path = ``
	__Resource__00000007_.InverseAppliedScaling = 0.000000

	__Task__00000000_.Name = `Submit PLN via ARO`
	__Task__00000000_.Description = ``
	__Task__00000000_.ComputedPrefix = ``
	__Task__00000000_.IsStartTask = false
	__Task__00000000_.IsEndTask = false
	__Task__00000000_.IsTaskNameNotProcessName = false

	__TaskShape__00000000_.Name = `-DiagramProcess`
	__TaskShape__00000000_.IsExpanded = false
	__TaskShape__00000000_.X = 111.000000
	__TaskShape__00000000_.Y = 275.873809
	__TaskShape__00000000_.Width = 238.000000
	__TaskShape__00000000_.Height = 70.000000
	__TaskShape__00000000_.IsHidden = false

	// insertion point for setup of pointers
	__AllocatedResourceShape__00000000_.Participant = __Participant__00000000_
	__AllocatedResourceShape__00000000_.Resource = __Resource__00000000_
	__AllocatedResourceShape__00000001_.Participant = __Participant__00000000_
	__AllocatedResourceShape__00000001_.Resource = __Resource__00000001_
	__AllocatedResourceShape__00000002_.Participant = __Participant__00000001_
	__AllocatedResourceShape__00000002_.Resource = __Resource__00000007_
	__AllocatedResourceShape__00000003_.Participant = __Participant__00000002_
	__AllocatedResourceShape__00000003_.Resource = __Resource__00000002_
	__DiagramProcess__00000000_.Process_Shapes = append(__DiagramProcess__00000000_.Process_Shapes, __ProcessShape__00000000_)
	__DiagramProcess__00000000_.Participant_Shapes = append(__DiagramProcess__00000000_.Participant_Shapes, __ParticipantShape__00000000_)
	__DiagramProcess__00000000_.Participant_Shapes = append(__DiagramProcess__00000000_.Participant_Shapes, __ParticipantShape__00000001_)
	__DiagramProcess__00000000_.Participant_Shapes = append(__DiagramProcess__00000000_.Participant_Shapes, __ParticipantShape__00000002_)
	__DiagramProcess__00000000_.Participant_Shapes = append(__DiagramProcess__00000000_.Participant_Shapes, __ParticipantShape__00000003_)
	__DiagramProcess__00000000_.ParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000000_.ParticipantWhoseNodeIsExpanded, __Participant__00000000_)
	__DiagramProcess__00000000_.ParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000000_.ParticipantWhoseNodeIsExpanded, __Participant__00000002_)
	__DiagramProcess__00000000_.Task_Shapes = append(__DiagramProcess__00000000_.Task_Shapes, __TaskShape__00000000_)
	__DiagramProcess__00000000_.AllocatedResourceShapes = append(__DiagramProcess__00000000_.AllocatedResourceShapes, __AllocatedResourceShape__00000000_)
	__DiagramProcess__00000000_.AllocatedResourceShapes = append(__DiagramProcess__00000000_.AllocatedResourceShapes, __AllocatedResourceShape__00000001_)
	__DiagramProcess__00000000_.AllocatedResourceShapes = append(__DiagramProcess__00000000_.AllocatedResourceShapes, __AllocatedResourceShape__00000002_)
	__DiagramProcess__00000000_.AllocatedResourceShapes = append(__DiagramProcess__00000000_.AllocatedResourceShapes, __AllocatedResourceShape__00000003_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000000_)
	__Library__00000000_.ProcesssWhoseNodeIsExpanded = append(__Library__00000000_.ProcesssWhoseNodeIsExpanded, __Process__00000000_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000000_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000001_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000002_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000003_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000004_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000005_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000006_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000007_)
	__Participant__00000000_.Resources = append(__Participant__00000000_.Resources, __Resource__00000000_)
	__Participant__00000000_.Resources = append(__Participant__00000000_.Resources, __Resource__00000001_)
	__Participant__00000000_.Tasks = append(__Participant__00000000_.Tasks, __Task__00000000_)
	__Participant__00000001_.Resources = append(__Participant__00000001_.Resources, __Resource__00000007_)
	__Participant__00000002_.Resources = append(__Participant__00000002_.Resources, __Resource__00000002_)
	__ParticipantShape__00000000_.Participant = __Participant__00000000_
	__ParticipantShape__00000001_.Participant = __Participant__00000001_
	__ParticipantShape__00000002_.Participant = __Participant__00000002_
	__ParticipantShape__00000003_.Participant = __Participant__00000003_
	__Process__00000000_.DiagramProcesss = append(__Process__00000000_.DiagramProcesss, __DiagramProcess__00000000_)
	__Process__00000000_.DiagramProcessWhoseNodeIsExpanded = append(__Process__00000000_.DiagramProcessWhoseNodeIsExpanded, __DiagramProcess__00000000_)
	__Process__00000000_.Participants = append(__Process__00000000_.Participants, __Participant__00000000_)
	__Process__00000000_.Participants = append(__Process__00000000_.Participants, __Participant__00000001_)
	__Process__00000000_.Participants = append(__Process__00000000_.Participants, __Participant__00000002_)
	__Process__00000000_.Participants = append(__Process__00000000_.Participants, __Participant__00000003_)
	__ProcessShape__00000000_.Process = __Process__00000000_
	__Task__00000000_.Type = nil
	__TaskShape__00000000_.Task = __Task__00000000_
}
