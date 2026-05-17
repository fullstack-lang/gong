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

	__AllocatedResourceShape__00000000_ := (&models.AllocatedResourceShape{Name: `RQS SPL, AO out of the loop-ATS Unit-ATC`}).Stage(stage)

	__ControlFlow__00000000_ := (&models.ControlFlow{Name: `"" to "Transmits RQS"`}).Stage(stage)
	__ControlFlow__00000001_ := (&models.ControlFlow{Name: `"Receives SPL" to ""`}).Stage(stage)
	__ControlFlow__00000002_ := (&models.ControlFlow{Name: `"" to "Transmits RQS"`}).Stage(stage)
	__ControlFlow__00000003_ := (&models.ControlFlow{Name: `"Receives SPL" to ""`}).Stage(stage)
	__ControlFlow__00000004_ := (&models.ControlFlow{Name: `"" to "Submit Flight Plan"`}).Stage(stage)
	__ControlFlow__00000005_ := (&models.ControlFlow{Name: `"Submit Flight Plan" to "Manage Operational Changes"`}).Stage(stage)
	__ControlFlow__00000006_ := (&models.ControlFlow{Name: `"Validate and Accept" to "Distribute ATS Messages"`}).Stage(stage)
	__ControlFlow__00000007_ := (&models.ControlFlow{Name: `"Ingest Flight Data" to "Log Flight Movement"`}).Stage(stage)
	__ControlFlow__00000008_ := (&models.ControlFlow{Name: `"Log Flight Movement" to ""`}).Stage(stage)

	__Data__00000000_ := (&models.Data{Name: `RQS`}).Stage(stage)
	__Data__00000001_ := (&models.Data{Name: `SPL`}).Stage(stage)
	__Data__00000002_ := (&models.Data{Name: `FPL (incl. item 19)`}).Stage(stage)
	__Data__00000003_ := (&models.Data{Name: `CHG`}).Stage(stage)
	__Data__00000004_ := (&models.Data{Name: `CNL`}).Stage(stage)
	__Data__00000005_ := (&models.Data{Name: `DLA`}).Stage(stage)
	__Data__00000006_ := (&models.Data{Name: `DEP`}).Stage(stage)
	__Data__00000007_ := (&models.Data{Name: `ARR`}).Stage(stage)
	__Data__00000008_ := (&models.Data{Name: `FPL (w/o item 19)`}).Stage(stage)

	__DiagramProcess__00000000_ := (&models.DiagramProcess{Name: `RQS SPL, AO out of the loop`}).Stage(stage)
	__DiagramProcess__00000001_ := (&models.DiagramProcess{Name: `RQS SPL, AO in the loop`}).Stage(stage)
	__DiagramProcess__00000002_ := (&models.DiagramProcess{Name: `DiagramProcess`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `Aviation Messaging Operations`}).Stage(stage)

	__Participant__00000000_ := (&models.Participant{Name: `ATS Unit`}).Stage(stage)
	__Participant__00000001_ := (&models.Participant{Name: `EUROCONTROL IFPS`}).Stage(stage)
	__Participant__00000002_ := (&models.Participant{Name: `ATS Unit`}).Stage(stage)
	__Participant__00000003_ := (&models.Participant{Name: `EUROCONTROL IFPS`}).Stage(stage)
	__Participant__00000004_ := (&models.Participant{Name: `Aircraft Operator`}).Stage(stage)
	__Participant__00000005_ := (&models.Participant{Name: `Aircraft Operator`}).Stage(stage)
	__Participant__00000006_ := (&models.Participant{Name: `ATS Reporting Office`}).Stage(stage)
	__Participant__00000007_ := (&models.Participant{Name: `ATS Unit`}).Stage(stage)

	__ParticipantShape__00000000_ := (&models.ParticipantShape{Name: `ATS Unit-RQS SPL, AO out of the loop`}).Stage(stage)
	__ParticipantShape__00000001_ := (&models.ParticipantShape{Name: `EUROCONTROL IFPS-RQS SPL, AO out of the loop`}).Stage(stage)

	__Process__00000000_ := (&models.Process{Name: `RQS/SPL Exchange (AO out of the loop)`}).Stage(stage)
	__Process__00000001_ := (&models.Process{Name: `RQS/SPL Exchange`}).Stage(stage)
	__Process__00000002_ := (&models.Process{Name: `ARO Message Exchanges`}).Stage(stage)

	__ProcessShape__00000003_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)
	__ProcessShape__00000004_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)
	__ProcessShape__00000005_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)

	__Resource__00000000_ := (&models.Resource{Name: `ATC`}).Stage(stage)
	__Resource__00000001_ := (&models.Resource{Name: `IFPS`}).Stage(stage)
	__Resource__00000002_ := (&models.Resource{Name: `AO`}).Stage(stage)
	__Resource__00000003_ := (&models.Resource{Name: `ARO`}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `Transmits RQS`}).Stage(stage)
	__Task__00000001_ := (&models.Task{Name: `Validates and Transmit SPL`}).Stage(stage)
	__Task__00000002_ := (&models.Task{Name: `Receives SPL`}).Stage(stage)
	__Task__00000003_ := (&models.Task{Name: ``}).Stage(stage)
	__Task__00000004_ := (&models.Task{Name: ``}).Stage(stage)
	__Task__00000005_ := (&models.Task{Name: `Transmits RQS`}).Stage(stage)
	__Task__00000006_ := (&models.Task{Name: `Validates and Routes RQS`}).Stage(stage)
	__Task__00000007_ := (&models.Task{Name: `Compiles and Transmits SPL`}).Stage(stage)
	__Task__00000008_ := (&models.Task{Name: `Validates and Routes SPL`}).Stage(stage)
	__Task__00000009_ := (&models.Task{Name: `Receives SPL`}).Stage(stage)
	__Task__00000010_ := (&models.Task{Name: ``}).Stage(stage)
	__Task__00000011_ := (&models.Task{Name: ``}).Stage(stage)
	__Task__00000012_ := (&models.Task{Name: `Submit Flight Plan`}).Stage(stage)
	__Task__00000013_ := (&models.Task{Name: `Manage Operational Changes`}).Stage(stage)
	__Task__00000014_ := (&models.Task{Name: `Validate and Accept`}).Stage(stage)
	__Task__00000015_ := (&models.Task{Name: `Distribute ATS Messages`}).Stage(stage)
	__Task__00000016_ := (&models.Task{Name: `Ingest Flight Data`}).Stage(stage)
	__Task__00000017_ := (&models.Task{Name: `Log Flight Movement`}).Stage(stage)
	__Task__00000018_ := (&models.Task{Name: ``}).Stage(stage)
	__Task__00000019_ := (&models.Task{Name: ``}).Stage(stage)

	__TaskShape__00000000_ := (&models.TaskShape{Name: `Transmits RQS-RQS SPL, AO out of the loop`}).Stage(stage)
	__TaskShape__00000001_ := (&models.TaskShape{Name: `Receives SPL-RQS SPL, AO out of the loop`}).Stage(stage)
	__TaskShape__00000002_ := (&models.TaskShape{Name: `-RQS SPL, AO out of the loop`}).Stage(stage)
	__TaskShape__00000003_ := (&models.TaskShape{Name: `-RQS SPL, AO out of the loop`}).Stage(stage)

	// insertion point for initialization of values

	__AllocatedResourceShape__00000000_.Name = `RQS SPL, AO out of the loop-ATS Unit-ATC`

	__ControlFlow__00000000_.Name = `"" to "Transmits RQS"`
	__ControlFlow__00000000_.Description = ``
	__ControlFlow__00000000_.ComputedPrefix = ``

	__ControlFlow__00000001_.Name = `"Receives SPL" to ""`
	__ControlFlow__00000001_.Description = ``
	__ControlFlow__00000001_.ComputedPrefix = ``

	__ControlFlow__00000002_.Name = `"" to "Transmits RQS"`
	__ControlFlow__00000002_.Description = ``
	__ControlFlow__00000002_.ComputedPrefix = ``

	__ControlFlow__00000003_.Name = `"Receives SPL" to ""`
	__ControlFlow__00000003_.Description = ``
	__ControlFlow__00000003_.ComputedPrefix = ``

	__ControlFlow__00000004_.Name = `"" to "Submit Flight Plan"`
	__ControlFlow__00000004_.Description = ``
	__ControlFlow__00000004_.ComputedPrefix = ``

	__ControlFlow__00000005_.Name = `"Submit Flight Plan" to "Manage Operational Changes"`
	__ControlFlow__00000005_.Description = ``
	__ControlFlow__00000005_.ComputedPrefix = ``

	__ControlFlow__00000006_.Name = `"Validate and Accept" to "Distribute ATS Messages"`
	__ControlFlow__00000006_.Description = ``
	__ControlFlow__00000006_.ComputedPrefix = ``

	__ControlFlow__00000007_.Name = `"Ingest Flight Data" to "Log Flight Movement"`
	__ControlFlow__00000007_.Description = ``
	__ControlFlow__00000007_.ComputedPrefix = ``

	__ControlFlow__00000008_.Name = `"Log Flight Movement" to ""`
	__ControlFlow__00000008_.Description = ``
	__ControlFlow__00000008_.ComputedPrefix = ``

	__Data__00000000_.Name = `RQS`
	__Data__00000000_.Acronym = ``
	__Data__00000000_.Description = ``
	__Data__00000000_.ComputedPrefix = ``
	__Data__00000000_.SVG_Path = ``
	__Data__00000000_.InverseAppliedScaling = 0.000000

	__Data__00000001_.Name = `SPL`
	__Data__00000001_.Acronym = ``
	__Data__00000001_.Description = ``
	__Data__00000001_.ComputedPrefix = ``
	__Data__00000001_.SVG_Path = ``
	__Data__00000001_.InverseAppliedScaling = 0.000000

	__Data__00000002_.Name = `FPL (incl. item 19)`
	__Data__00000002_.Acronym = ``
	__Data__00000002_.Description = ``
	__Data__00000002_.ComputedPrefix = ``
	__Data__00000002_.SVG_Path = ``
	__Data__00000002_.InverseAppliedScaling = 0.000000

	__Data__00000003_.Name = `CHG`
	__Data__00000003_.Acronym = ``
	__Data__00000003_.Description = ``
	__Data__00000003_.ComputedPrefix = ``
	__Data__00000003_.SVG_Path = ``
	__Data__00000003_.InverseAppliedScaling = 0.000000

	__Data__00000004_.Name = `CNL`
	__Data__00000004_.Acronym = ``
	__Data__00000004_.Description = ``
	__Data__00000004_.ComputedPrefix = ``
	__Data__00000004_.SVG_Path = ``
	__Data__00000004_.InverseAppliedScaling = 0.000000

	__Data__00000005_.Name = `DLA`
	__Data__00000005_.Acronym = ``
	__Data__00000005_.Description = ``
	__Data__00000005_.ComputedPrefix = ``
	__Data__00000005_.SVG_Path = ``
	__Data__00000005_.InverseAppliedScaling = 0.000000

	__Data__00000006_.Name = `DEP`
	__Data__00000006_.Acronym = ``
	__Data__00000006_.Description = ``
	__Data__00000006_.ComputedPrefix = ``
	__Data__00000006_.SVG_Path = ``
	__Data__00000006_.InverseAppliedScaling = 0.000000

	__Data__00000007_.Name = `ARR`
	__Data__00000007_.Acronym = ``
	__Data__00000007_.Description = ``
	__Data__00000007_.ComputedPrefix = ``
	__Data__00000007_.SVG_Path = ``
	__Data__00000007_.InverseAppliedScaling = 0.000000

	__Data__00000008_.Name = `FPL (w/o item 19)`
	__Data__00000008_.Acronym = ``
	__Data__00000008_.Description = ``
	__Data__00000008_.ComputedPrefix = ``
	__Data__00000008_.SVG_Path = ``
	__Data__00000008_.InverseAppliedScaling = 0.000000

	__DiagramProcess__00000000_.Name = `RQS SPL, AO out of the loop`
	__DiagramProcess__00000000_.Description = ``
	__DiagramProcess__00000000_.ComputedPrefix = ``
	__DiagramProcess__00000000_.IsChecked = false
	__DiagramProcess__00000000_.IsEditable_ = false
	__DiagramProcess__00000000_.IsShowPrefix = false
	__DiagramProcess__00000000_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000000_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000000_.Width = 900.000000
	__DiagramProcess__00000000_.Height = 1350.000000
	__DiagramProcess__00000000_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000000_.IsParticipantsNodeExpanded = true
	__DiagramProcess__00000000_.IsExternalParticipantsNodeExpanded = false
	__DiagramProcess__00000000_.IsNotesNodeExpanded = false

	__DiagramProcess__00000001_.Name = `RQS SPL, AO in the loop`
	__DiagramProcess__00000001_.Description = ``
	__DiagramProcess__00000001_.ComputedPrefix = ``
	__DiagramProcess__00000001_.IsChecked = false
	__DiagramProcess__00000001_.IsEditable_ = false
	__DiagramProcess__00000001_.IsShowPrefix = false
	__DiagramProcess__00000001_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000001_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000001_.Width = 900.000000
	__DiagramProcess__00000001_.Height = 1350.000000
	__DiagramProcess__00000001_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000001_.IsParticipantsNodeExpanded = true
	__DiagramProcess__00000001_.IsExternalParticipantsNodeExpanded = false
	__DiagramProcess__00000001_.IsNotesNodeExpanded = false

	__DiagramProcess__00000002_.Name = `DiagramProcess`
	__DiagramProcess__00000002_.Description = ``
	__DiagramProcess__00000002_.ComputedPrefix = ``
	__DiagramProcess__00000002_.IsChecked = true
	__DiagramProcess__00000002_.IsEditable_ = false
	__DiagramProcess__00000002_.IsShowPrefix = false
	__DiagramProcess__00000002_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000002_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000002_.Width = 900.000000
	__DiagramProcess__00000002_.Height = 1350.000000
	__DiagramProcess__00000002_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000002_.IsParticipantsNodeExpanded = false
	__DiagramProcess__00000002_.IsExternalParticipantsNodeExpanded = false
	__DiagramProcess__00000002_.IsNotesNodeExpanded = false

	__Library__00000000_.Name = `Aviation Messaging Operations`
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

	__Participant__00000000_.Name = `ATS Unit`
	__Participant__00000000_.IsProcessResource = false
	__Participant__00000000_.Description = ``
	__Participant__00000000_.IsResourcesNodeExpanded = true
	__Participant__00000000_.IsProcessesNodeExpanded = false
	__Participant__00000000_.ComputedPrefix = ``
	__Participant__00000000_.IsTasksNodeExpanded = true
	__Participant__00000000_.IsControlFlowsNodeExpanded = false
	__Participant__00000000_.IsDataFlowsNodeExpanded = false

	__Participant__00000001_.Name = `EUROCONTROL IFPS`
	__Participant__00000001_.IsProcessResource = false
	__Participant__00000001_.Description = ``
	__Participant__00000001_.IsResourcesNodeExpanded = false
	__Participant__00000001_.IsProcessesNodeExpanded = false
	__Participant__00000001_.ComputedPrefix = ``
	__Participant__00000001_.IsTasksNodeExpanded = false
	__Participant__00000001_.IsControlFlowsNodeExpanded = false
	__Participant__00000001_.IsDataFlowsNodeExpanded = false

	__Participant__00000002_.Name = `ATS Unit`
	__Participant__00000002_.IsProcessResource = false
	__Participant__00000002_.Description = ``
	__Participant__00000002_.IsResourcesNodeExpanded = false
	__Participant__00000002_.IsProcessesNodeExpanded = false
	__Participant__00000002_.ComputedPrefix = ``
	__Participant__00000002_.IsTasksNodeExpanded = false
	__Participant__00000002_.IsControlFlowsNodeExpanded = false
	__Participant__00000002_.IsDataFlowsNodeExpanded = false

	__Participant__00000003_.Name = `EUROCONTROL IFPS`
	__Participant__00000003_.IsProcessResource = false
	__Participant__00000003_.Description = ``
	__Participant__00000003_.IsResourcesNodeExpanded = false
	__Participant__00000003_.IsProcessesNodeExpanded = false
	__Participant__00000003_.ComputedPrefix = ``
	__Participant__00000003_.IsTasksNodeExpanded = false
	__Participant__00000003_.IsControlFlowsNodeExpanded = false
	__Participant__00000003_.IsDataFlowsNodeExpanded = false

	__Participant__00000004_.Name = `Aircraft Operator`
	__Participant__00000004_.IsProcessResource = false
	__Participant__00000004_.Description = ``
	__Participant__00000004_.IsResourcesNodeExpanded = false
	__Participant__00000004_.IsProcessesNodeExpanded = false
	__Participant__00000004_.ComputedPrefix = ``
	__Participant__00000004_.IsTasksNodeExpanded = false
	__Participant__00000004_.IsControlFlowsNodeExpanded = false
	__Participant__00000004_.IsDataFlowsNodeExpanded = false

	__Participant__00000005_.Name = `Aircraft Operator`
	__Participant__00000005_.IsProcessResource = false
	__Participant__00000005_.Description = ``
	__Participant__00000005_.IsResourcesNodeExpanded = false
	__Participant__00000005_.IsProcessesNodeExpanded = false
	__Participant__00000005_.ComputedPrefix = ``
	__Participant__00000005_.IsTasksNodeExpanded = false
	__Participant__00000005_.IsControlFlowsNodeExpanded = false
	__Participant__00000005_.IsDataFlowsNodeExpanded = false

	__Participant__00000006_.Name = `ATS Reporting Office`
	__Participant__00000006_.IsProcessResource = false
	__Participant__00000006_.Description = ``
	__Participant__00000006_.IsResourcesNodeExpanded = false
	__Participant__00000006_.IsProcessesNodeExpanded = false
	__Participant__00000006_.ComputedPrefix = ``
	__Participant__00000006_.IsTasksNodeExpanded = false
	__Participant__00000006_.IsControlFlowsNodeExpanded = false
	__Participant__00000006_.IsDataFlowsNodeExpanded = false

	__Participant__00000007_.Name = `ATS Unit`
	__Participant__00000007_.IsProcessResource = false
	__Participant__00000007_.Description = ``
	__Participant__00000007_.IsResourcesNodeExpanded = false
	__Participant__00000007_.IsProcessesNodeExpanded = false
	__Participant__00000007_.ComputedPrefix = ``
	__Participant__00000007_.IsTasksNodeExpanded = false
	__Participant__00000007_.IsControlFlowsNodeExpanded = false
	__Participant__00000007_.IsDataFlowsNodeExpanded = false

	__ParticipantShape__00000000_.Name = `ATS Unit-RQS SPL, AO out of the loop`
	__ParticipantShape__00000000_.IsExpanded = false
	__ParticipantShape__00000000_.X = 113.987403
	__ParticipantShape__00000000_.Y = 146.428662
	__ParticipantShape__00000000_.Width = 250.000000
	__ParticipantShape__00000000_.Height = 70.000000
	__ParticipantShape__00000000_.IsHidden = false
	__ParticipantShape__00000000_.WidthWeight = 1.000000

	__ParticipantShape__00000001_.Name = `EUROCONTROL IFPS-RQS SPL, AO out of the loop`
	__ParticipantShape__00000001_.IsExpanded = false
	__ParticipantShape__00000001_.X = 153.930407
	__ParticipantShape__00000001_.Y = 113.052081
	__ParticipantShape__00000001_.Width = 250.000000
	__ParticipantShape__00000001_.Height = 70.000000
	__ParticipantShape__00000001_.IsHidden = false
	__ParticipantShape__00000001_.WidthWeight = 1.000000

	__Process__00000000_.Name = `RQS/SPL Exchange (AO out of the loop)`
	__Process__00000000_.Description = ``
	__Process__00000000_.ComputedPrefix = ``
	__Process__00000000_.SVG_Path = ``
	__Process__00000000_.InverseAppliedScaling = 0.000000
	__Process__00000000_.IsSubProcessNodeExpanded = false
	__Process__00000000_.IsDataFlowsNodeExpanded = false

	__Process__00000001_.Name = `RQS/SPL Exchange`
	__Process__00000001_.Description = ``
	__Process__00000001_.ComputedPrefix = ``
	__Process__00000001_.SVG_Path = ``
	__Process__00000001_.InverseAppliedScaling = 0.000000
	__Process__00000001_.IsSubProcessNodeExpanded = false
	__Process__00000001_.IsDataFlowsNodeExpanded = false

	__Process__00000002_.Name = `ARO Message Exchanges`
	__Process__00000002_.Description = ``
	__Process__00000002_.ComputedPrefix = ``
	__Process__00000002_.SVG_Path = ``
	__Process__00000002_.InverseAppliedScaling = 0.000000
	__Process__00000002_.IsSubProcessNodeExpanded = false
	__Process__00000002_.IsDataFlowsNodeExpanded = false

	__ProcessShape__00000003_.Name = `ProcessShape`
	__ProcessShape__00000003_.IsExpanded = false
	__ProcessShape__00000003_.X = 100.000000
	__ProcessShape__00000003_.Y = 50.000000
	__ProcessShape__00000003_.Width = 500.000000
	__ProcessShape__00000003_.Height = 1000.000000
	__ProcessShape__00000003_.IsHidden = false

	__ProcessShape__00000004_.Name = `ProcessShape`
	__ProcessShape__00000004_.IsExpanded = false
	__ProcessShape__00000004_.X = 100.000000
	__ProcessShape__00000004_.Y = 50.000000
	__ProcessShape__00000004_.Width = 500.000000
	__ProcessShape__00000004_.Height = 1000.000000
	__ProcessShape__00000004_.IsHidden = false

	__ProcessShape__00000005_.Name = `ProcessShape`
	__ProcessShape__00000005_.IsExpanded = false
	__ProcessShape__00000005_.X = 100.000000
	__ProcessShape__00000005_.Y = 50.000000
	__ProcessShape__00000005_.Width = 500.000000
	__ProcessShape__00000005_.Height = 1000.000000
	__ProcessShape__00000005_.IsHidden = false

	__Resource__00000000_.Name = `ATC`
	__Resource__00000000_.Acronym = ``
	__Resource__00000000_.Description = ``
	__Resource__00000000_.ComputedPrefix = ``
	__Resource__00000000_.SVG_Path = ``
	__Resource__00000000_.InverseAppliedScaling = 0.000000

	__Resource__00000001_.Name = `IFPS`
	__Resource__00000001_.Acronym = ``
	__Resource__00000001_.Description = ``
	__Resource__00000001_.ComputedPrefix = ``
	__Resource__00000001_.SVG_Path = ``
	__Resource__00000001_.InverseAppliedScaling = 0.000000

	__Resource__00000002_.Name = `AO`
	__Resource__00000002_.Acronym = ``
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

	__Task__00000000_.Name = `Transmits RQS`
	__Task__00000000_.Description = ``
	__Task__00000000_.ComputedPrefix = ``
	__Task__00000000_.IsStartTask = false
	__Task__00000000_.IsEndTask = false
	__Task__00000000_.IsTaskNameNotProcessName = false

	__Task__00000001_.Name = `Validates and Transmit SPL`
	__Task__00000001_.Description = ``
	__Task__00000001_.ComputedPrefix = ``
	__Task__00000001_.IsStartTask = false
	__Task__00000001_.IsEndTask = false
	__Task__00000001_.IsTaskNameNotProcessName = false

	__Task__00000002_.Name = `Receives SPL`
	__Task__00000002_.Description = ``
	__Task__00000002_.ComputedPrefix = ``
	__Task__00000002_.IsStartTask = false
	__Task__00000002_.IsEndTask = false
	__Task__00000002_.IsTaskNameNotProcessName = false

	__Task__00000003_.Name = ``
	__Task__00000003_.Description = ``
	__Task__00000003_.ComputedPrefix = ``
	__Task__00000003_.IsStartTask = false
	__Task__00000003_.IsEndTask = false
	__Task__00000003_.IsTaskNameNotProcessName = false

	__Task__00000004_.Name = ``
	__Task__00000004_.Description = ``
	__Task__00000004_.ComputedPrefix = ``
	__Task__00000004_.IsStartTask = false
	__Task__00000004_.IsEndTask = false
	__Task__00000004_.IsTaskNameNotProcessName = false

	__Task__00000005_.Name = `Transmits RQS`
	__Task__00000005_.Description = ``
	__Task__00000005_.ComputedPrefix = ``
	__Task__00000005_.IsStartTask = false
	__Task__00000005_.IsEndTask = false
	__Task__00000005_.IsTaskNameNotProcessName = false

	__Task__00000006_.Name = `Validates and Routes RQS`
	__Task__00000006_.Description = ``
	__Task__00000006_.ComputedPrefix = ``
	__Task__00000006_.IsStartTask = false
	__Task__00000006_.IsEndTask = false
	__Task__00000006_.IsTaskNameNotProcessName = false

	__Task__00000007_.Name = `Compiles and Transmits SPL`
	__Task__00000007_.Description = ``
	__Task__00000007_.ComputedPrefix = ``
	__Task__00000007_.IsStartTask = false
	__Task__00000007_.IsEndTask = false
	__Task__00000007_.IsTaskNameNotProcessName = false

	__Task__00000008_.Name = `Validates and Routes SPL`
	__Task__00000008_.Description = ``
	__Task__00000008_.ComputedPrefix = ``
	__Task__00000008_.IsStartTask = false
	__Task__00000008_.IsEndTask = false
	__Task__00000008_.IsTaskNameNotProcessName = false

	__Task__00000009_.Name = `Receives SPL`
	__Task__00000009_.Description = ``
	__Task__00000009_.ComputedPrefix = ``
	__Task__00000009_.IsStartTask = false
	__Task__00000009_.IsEndTask = false
	__Task__00000009_.IsTaskNameNotProcessName = false

	__Task__00000010_.Name = ``
	__Task__00000010_.Description = ``
	__Task__00000010_.ComputedPrefix = ``
	__Task__00000010_.IsStartTask = false
	__Task__00000010_.IsEndTask = false
	__Task__00000010_.IsTaskNameNotProcessName = false

	__Task__00000011_.Name = ``
	__Task__00000011_.Description = ``
	__Task__00000011_.ComputedPrefix = ``
	__Task__00000011_.IsStartTask = false
	__Task__00000011_.IsEndTask = false
	__Task__00000011_.IsTaskNameNotProcessName = false

	__Task__00000012_.Name = `Submit Flight Plan`
	__Task__00000012_.Description = ``
	__Task__00000012_.ComputedPrefix = ``
	__Task__00000012_.IsStartTask = false
	__Task__00000012_.IsEndTask = false
	__Task__00000012_.IsTaskNameNotProcessName = false

	__Task__00000013_.Name = `Manage Operational Changes`
	__Task__00000013_.Description = ``
	__Task__00000013_.ComputedPrefix = ``
	__Task__00000013_.IsStartTask = false
	__Task__00000013_.IsEndTask = false
	__Task__00000013_.IsTaskNameNotProcessName = false

	__Task__00000014_.Name = `Validate and Accept`
	__Task__00000014_.Description = ``
	__Task__00000014_.ComputedPrefix = ``
	__Task__00000014_.IsStartTask = false
	__Task__00000014_.IsEndTask = false
	__Task__00000014_.IsTaskNameNotProcessName = false

	__Task__00000015_.Name = `Distribute ATS Messages`
	__Task__00000015_.Description = ``
	__Task__00000015_.ComputedPrefix = ``
	__Task__00000015_.IsStartTask = false
	__Task__00000015_.IsEndTask = false
	__Task__00000015_.IsTaskNameNotProcessName = false

	__Task__00000016_.Name = `Ingest Flight Data`
	__Task__00000016_.Description = ``
	__Task__00000016_.ComputedPrefix = ``
	__Task__00000016_.IsStartTask = false
	__Task__00000016_.IsEndTask = false
	__Task__00000016_.IsTaskNameNotProcessName = false

	__Task__00000017_.Name = `Log Flight Movement`
	__Task__00000017_.Description = ``
	__Task__00000017_.ComputedPrefix = ``
	__Task__00000017_.IsStartTask = false
	__Task__00000017_.IsEndTask = false
	__Task__00000017_.IsTaskNameNotProcessName = false

	__Task__00000018_.Name = ``
	__Task__00000018_.Description = ``
	__Task__00000018_.ComputedPrefix = ``
	__Task__00000018_.IsStartTask = false
	__Task__00000018_.IsEndTask = false
	__Task__00000018_.IsTaskNameNotProcessName = false

	__Task__00000019_.Name = ``
	__Task__00000019_.Description = ``
	__Task__00000019_.ComputedPrefix = ``
	__Task__00000019_.IsStartTask = false
	__Task__00000019_.IsEndTask = false
	__Task__00000019_.IsTaskNameNotProcessName = false

	__TaskShape__00000000_.Name = `Transmits RQS-RQS SPL, AO out of the loop`
	__TaskShape__00000000_.IsExpanded = false
	__TaskShape__00000000_.X = 111.000000
	__TaskShape__00000000_.Y = 115.237571
	__TaskShape__00000000_.Width = 238.000000
	__TaskShape__00000000_.Height = 70.000000
	__TaskShape__00000000_.IsHidden = false

	__TaskShape__00000001_.Name = `Receives SPL-RQS SPL, AO out of the loop`
	__TaskShape__00000001_.IsExpanded = false
	__TaskShape__00000001_.X = 111.000000
	__TaskShape__00000001_.Y = 187.926708
	__TaskShape__00000001_.Width = 238.000000
	__TaskShape__00000001_.Height = 70.000000
	__TaskShape__00000001_.IsHidden = false

	__TaskShape__00000002_.Name = `-RQS SPL, AO out of the loop`
	__TaskShape__00000002_.IsExpanded = false
	__TaskShape__00000002_.X = 111.000000
	__TaskShape__00000002_.Y = 168.558203
	__TaskShape__00000002_.Width = 238.000000
	__TaskShape__00000002_.Height = 70.000000
	__TaskShape__00000002_.IsHidden = false

	__TaskShape__00000003_.Name = `-RQS SPL, AO out of the loop`
	__TaskShape__00000003_.IsExpanded = false
	__TaskShape__00000003_.X = 111.000000
	__TaskShape__00000003_.Y = 117.421095
	__TaskShape__00000003_.Width = 238.000000
	__TaskShape__00000003_.Height = 70.000000
	__TaskShape__00000003_.IsHidden = false

	// insertion point for setup of pointers
	__AllocatedResourceShape__00000000_.Participant = __Participant__00000000_
	__AllocatedResourceShape__00000000_.Resource = __Resource__00000000_
	__ControlFlow__00000000_.Start = __Task__00000003_
	__ControlFlow__00000000_.End = __Task__00000000_
	__ControlFlow__00000001_.Start = __Task__00000002_
	__ControlFlow__00000001_.End = __Task__00000004_
	__ControlFlow__00000002_.Start = __Task__00000010_
	__ControlFlow__00000002_.End = __Task__00000005_
	__ControlFlow__00000003_.Start = __Task__00000009_
	__ControlFlow__00000003_.End = __Task__00000011_
	__ControlFlow__00000004_.Start = __Task__00000018_
	__ControlFlow__00000004_.End = __Task__00000012_
	__ControlFlow__00000005_.Start = __Task__00000012_
	__ControlFlow__00000005_.End = __Task__00000013_
	__ControlFlow__00000006_.Start = __Task__00000014_
	__ControlFlow__00000006_.End = __Task__00000015_
	__ControlFlow__00000007_.Start = __Task__00000016_
	__ControlFlow__00000007_.End = __Task__00000017_
	__ControlFlow__00000008_.Start = __Task__00000017_
	__ControlFlow__00000008_.End = __Task__00000019_
	__DiagramProcess__00000000_.Process_Shapes = append(__DiagramProcess__00000000_.Process_Shapes, __ProcessShape__00000003_)
	__DiagramProcess__00000000_.Participant_Shapes = append(__DiagramProcess__00000000_.Participant_Shapes, __ParticipantShape__00000000_)
	__DiagramProcess__00000000_.Participant_Shapes = append(__DiagramProcess__00000000_.Participant_Shapes, __ParticipantShape__00000001_)
	__DiagramProcess__00000000_.ParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000000_.ParticipantWhoseNodeIsExpanded, __Participant__00000000_)
	__DiagramProcess__00000000_.Task_Shapes = append(__DiagramProcess__00000000_.Task_Shapes, __TaskShape__00000000_)
	__DiagramProcess__00000000_.Task_Shapes = append(__DiagramProcess__00000000_.Task_Shapes, __TaskShape__00000001_)
	__DiagramProcess__00000000_.Task_Shapes = append(__DiagramProcess__00000000_.Task_Shapes, __TaskShape__00000002_)
	__DiagramProcess__00000000_.Task_Shapes = append(__DiagramProcess__00000000_.Task_Shapes, __TaskShape__00000003_)
	__DiagramProcess__00000000_.AllocatedResourceShapes = append(__DiagramProcess__00000000_.AllocatedResourceShapes, __AllocatedResourceShape__00000000_)
	__DiagramProcess__00000001_.Process_Shapes = append(__DiagramProcess__00000001_.Process_Shapes, __ProcessShape__00000004_)
	__DiagramProcess__00000002_.Process_Shapes = append(__DiagramProcess__00000002_.Process_Shapes, __ProcessShape__00000005_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000000_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000001_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000002_)
	__Library__00000000_.ProcesssWhoseNodeIsExpanded = append(__Library__00000000_.ProcesssWhoseNodeIsExpanded, __Process__00000000_)
	__Library__00000000_.ProcesssWhoseNodeIsExpanded = append(__Library__00000000_.ProcesssWhoseNodeIsExpanded, __Process__00000001_)
	__Library__00000000_.ProcesssWhoseNodeIsExpanded = append(__Library__00000000_.ProcesssWhoseNodeIsExpanded, __Process__00000002_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000000_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000001_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000002_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000003_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000004_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000005_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000006_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000007_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000008_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000000_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000001_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000002_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000003_)
	__Participant__00000000_.Resources = append(__Participant__00000000_.Resources, __Resource__00000000_)
	__Participant__00000000_.Tasks = append(__Participant__00000000_.Tasks, __Task__00000000_)
	__Participant__00000000_.Tasks = append(__Participant__00000000_.Tasks, __Task__00000002_)
	__Participant__00000000_.Tasks = append(__Participant__00000000_.Tasks, __Task__00000003_)
	__Participant__00000000_.Tasks = append(__Participant__00000000_.Tasks, __Task__00000004_)
	__Participant__00000000_.ControlFlows = append(__Participant__00000000_.ControlFlows, __ControlFlow__00000000_)
	__Participant__00000000_.ControlFlows = append(__Participant__00000000_.ControlFlows, __ControlFlow__00000001_)
	__Participant__00000001_.Resources = append(__Participant__00000001_.Resources, __Resource__00000001_)
	__Participant__00000001_.Tasks = append(__Participant__00000001_.Tasks, __Task__00000001_)
	__Participant__00000002_.Resources = append(__Participant__00000002_.Resources, __Resource__00000000_)
	__Participant__00000002_.Tasks = append(__Participant__00000002_.Tasks, __Task__00000005_)
	__Participant__00000002_.Tasks = append(__Participant__00000002_.Tasks, __Task__00000009_)
	__Participant__00000002_.Tasks = append(__Participant__00000002_.Tasks, __Task__00000010_)
	__Participant__00000002_.Tasks = append(__Participant__00000002_.Tasks, __Task__00000011_)
	__Participant__00000002_.ControlFlows = append(__Participant__00000002_.ControlFlows, __ControlFlow__00000002_)
	__Participant__00000002_.ControlFlows = append(__Participant__00000002_.ControlFlows, __ControlFlow__00000003_)
	__Participant__00000003_.Resources = append(__Participant__00000003_.Resources, __Resource__00000001_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000006_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000008_)
	__Participant__00000004_.Resources = append(__Participant__00000004_.Resources, __Resource__00000002_)
	__Participant__00000004_.Tasks = append(__Participant__00000004_.Tasks, __Task__00000007_)
	__Participant__00000005_.Resources = append(__Participant__00000005_.Resources, __Resource__00000002_)
	__Participant__00000005_.Tasks = append(__Participant__00000005_.Tasks, __Task__00000018_)
	__Participant__00000005_.Tasks = append(__Participant__00000005_.Tasks, __Task__00000012_)
	__Participant__00000005_.Tasks = append(__Participant__00000005_.Tasks, __Task__00000013_)
	__Participant__00000005_.ControlFlows = append(__Participant__00000005_.ControlFlows, __ControlFlow__00000004_)
	__Participant__00000005_.ControlFlows = append(__Participant__00000005_.ControlFlows, __ControlFlow__00000005_)
	__Participant__00000006_.Resources = append(__Participant__00000006_.Resources, __Resource__00000003_)
	__Participant__00000006_.Tasks = append(__Participant__00000006_.Tasks, __Task__00000014_)
	__Participant__00000006_.Tasks = append(__Participant__00000006_.Tasks, __Task__00000015_)
	__Participant__00000006_.ControlFlows = append(__Participant__00000006_.ControlFlows, __ControlFlow__00000006_)
	__Participant__00000007_.Resources = append(__Participant__00000007_.Resources, __Resource__00000000_)
	__Participant__00000007_.Tasks = append(__Participant__00000007_.Tasks, __Task__00000016_)
	__Participant__00000007_.Tasks = append(__Participant__00000007_.Tasks, __Task__00000017_)
	__Participant__00000007_.Tasks = append(__Participant__00000007_.Tasks, __Task__00000019_)
	__Participant__00000007_.ControlFlows = append(__Participant__00000007_.ControlFlows, __ControlFlow__00000007_)
	__Participant__00000007_.ControlFlows = append(__Participant__00000007_.ControlFlows, __ControlFlow__00000008_)
	__ParticipantShape__00000000_.Participant = __Participant__00000000_
	__ParticipantShape__00000001_.Participant = __Participant__00000001_
	__Process__00000000_.DiagramProcesss = append(__Process__00000000_.DiagramProcesss, __DiagramProcess__00000000_)
	__Process__00000000_.Participants = append(__Process__00000000_.Participants, __Participant__00000000_)
	__Process__00000000_.Participants = append(__Process__00000000_.Participants, __Participant__00000001_)
	__Process__00000001_.DiagramProcesss = append(__Process__00000001_.DiagramProcesss, __DiagramProcess__00000001_)
	__Process__00000001_.Participants = append(__Process__00000001_.Participants, __Participant__00000002_)
	__Process__00000001_.Participants = append(__Process__00000001_.Participants, __Participant__00000003_)
	__Process__00000001_.Participants = append(__Process__00000001_.Participants, __Participant__00000004_)
	__Process__00000002_.DiagramProcesss = append(__Process__00000002_.DiagramProcesss, __DiagramProcess__00000002_)
	__Process__00000002_.Participants = append(__Process__00000002_.Participants, __Participant__00000005_)
	__Process__00000002_.Participants = append(__Process__00000002_.Participants, __Participant__00000006_)
	__Process__00000002_.Participants = append(__Process__00000002_.Participants, __Participant__00000007_)
	__ProcessShape__00000003_.Process = __Process__00000000_
	__ProcessShape__00000004_.Process = __Process__00000001_
	__ProcessShape__00000005_.Process = __Process__00000002_
	__Task__00000000_.Type = nil
	__Task__00000001_.Type = nil
	__Task__00000002_.Type = nil
	__Task__00000003_.Type = nil
	__Task__00000004_.Type = nil
	__Task__00000005_.Type = nil
	__Task__00000006_.Type = nil
	__Task__00000007_.Type = nil
	__Task__00000008_.Type = nil
	__Task__00000009_.Type = nil
	__Task__00000010_.Type = nil
	__Task__00000011_.Type = nil
	__Task__00000012_.Type = nil
	__Task__00000013_.Type = nil
	__Task__00000014_.Type = nil
	__Task__00000015_.Type = nil
	__Task__00000016_.Type = nil
	__Task__00000017_.Type = nil
	__Task__00000018_.Type = nil
	__Task__00000019_.Type = nil
	__TaskShape__00000000_.Task = __Task__00000000_
	__TaskShape__00000001_.Task = __Task__00000002_
	__TaskShape__00000002_.Task = __Task__00000003_
	__TaskShape__00000003_.Task = __Task__00000004_
}
