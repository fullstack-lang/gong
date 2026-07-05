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

	__AllocatedResourceShape__00000000_ := (&models.AllocatedResourceShape{Name: `DiagramProcess-Aircraft Operator-AO`}).Stage(stage)
	__AllocatedResourceShape__00000001_ := (&models.AllocatedResourceShape{Name: `DiagramProcess-ATS Reporting Office-ARO`}).Stage(stage)
	__AllocatedResourceShape__00000002_ := (&models.AllocatedResourceShape{Name: `DiagramProcess-ATS Unit-ATC`}).Stage(stage)

	__ControlFlow__00000000_ := (&models.ControlFlow{Name: `"Submit Flight Plan" to "Manage Operational Changes"`}).Stage(stage)
	__ControlFlow__00000001_ := (&models.ControlFlow{Name: `"Validate and Accept" to "Distribute ATS Messages"`}).Stage(stage)
	__ControlFlow__00000002_ := (&models.ControlFlow{Name: `"Ingest Flight Data" to "Log Flight Movement"`}).Stage(stage)
	__ControlFlow__00000003_ := (&models.ControlFlow{Name: `"" to "Submit Flight Plan"`}).Stage(stage)
	__ControlFlow__00000004_ := (&models.ControlFlow{Name: `"Log Flight Movement" to "End"`}).Stage(stage)

	__ControlFlowShape__00000000_ := (&models.ControlFlowShape{Name: `"Submit Flight Plan" to "Manage Operational Changes"`}).Stage(stage)
	__ControlFlowShape__00000001_ := (&models.ControlFlowShape{Name: `"Validate and Accept" to "Distribute ATS Messages"`}).Stage(stage)
	__ControlFlowShape__00000002_ := (&models.ControlFlowShape{Name: `"Ingest Flight Data" to "Log Flight Movement"`}).Stage(stage)
	__ControlFlowShape__00000003_ := (&models.ControlFlowShape{Name: `"" to "Submit Flight Plan"`}).Stage(stage)
	__ControlFlowShape__00000004_ := (&models.ControlFlowShape{Name: `"Log Flight Movement" to "End"`}).Stage(stage)

	__Data__00000000_ := (&models.Data{Name: `FPL (incl. item 19)`}).Stage(stage)
	__Data__00000001_ := (&models.Data{Name: `CHG`}).Stage(stage)
	__Data__00000002_ := (&models.Data{Name: `CNL`}).Stage(stage)
	__Data__00000003_ := (&models.Data{Name: `DLA`}).Stage(stage)
	__Data__00000004_ := (&models.Data{Name: `DEP`}).Stage(stage)
	__Data__00000005_ := (&models.Data{Name: `ARR`}).Stage(stage)
	__Data__00000006_ := (&models.Data{Name: `FPL (w/o item 19)`}).Stage(stage)
	__Data__00000007_ := (&models.Data{Name: `RQS`}).Stage(stage)
	__Data__00000008_ := (&models.Data{Name: `SPL`}).Stage(stage)

	__DataFlow__00000000_ := (&models.DataFlow{Name: `"Submit Flight Plan" to "Validate and Accept"`}).Stage(stage)
	__DataFlow__00000001_ := (&models.DataFlow{Name: `"Manage Operational Changes" to "Validate and Accept"`}).Stage(stage)
	__DataFlow__00000002_ := (&models.DataFlow{Name: `"Distribute ATS Messages" to "Ingest Flight Data"`}).Stage(stage)
	__DataFlow__00000003_ := (&models.DataFlow{Name: `"Log Flight Movement" to "Validate and Accept"`}).Stage(stage)
	__DataFlow__00000004_ := (&models.DataFlow{Name: `"Log Flight Movement" to "Validate and Accept"`}).Stage(stage)

	__DataFlowShape__00000000_ := (&models.DataFlowShape{Name: `"Submit Flight Plan" to "Validate and Accept"`}).Stage(stage)
	__DataFlowShape__00000001_ := (&models.DataFlowShape{Name: `"Manage Operational Changes" to "Validate and Accept"`}).Stage(stage)
	__DataFlowShape__00000002_ := (&models.DataFlowShape{Name: `"Distribute ATS Messages" to "Ingest Flight Data"`}).Stage(stage)
	__DataFlowShape__00000003_ := (&models.DataFlowShape{Name: `"Log Flight Movement" to "Validate and Accept"`}).Stage(stage)
	__DataFlowShape__00000004_ := (&models.DataFlowShape{Name: `"Log Flight Movement" to "Validate and Accept"`}).Stage(stage)

	__DataShape__00000000_ := (&models.DataShape{Name: `"Submit Flight Plan" to "Validate and Accept"-"Submit Flight Plan" to "Validate and Accept"-DiagramProcess`}).Stage(stage)
	__DataShape__00000001_ := (&models.DataShape{Name: `"Manage Operational Changes" to "Validate and Accept"-"Manage Operational Changes" to "Validate and Accept"-DiagramProcess`}).Stage(stage)
	__DataShape__00000002_ := (&models.DataShape{Name: `"Distribute ATS Messages" to "Ingest Flight Data"-"Distribute ATS Messages" to "Ingest Flight Data"-DiagramProcess`}).Stage(stage)
	__DataShape__00000003_ := (&models.DataShape{Name: `"Log Flight Movement" to "Validate and Accept" - DEP-"Log Flight Movement" to "Validate and Accept" - DEP-DiagramProcess`}).Stage(stage)
	__DataShape__00000004_ := (&models.DataShape{Name: `"Log Flight Movement" to "Validate and Accept" - ARR-"Log Flight Movement" to "Validate and Accept" - ARR-DiagramProcess`}).Stage(stage)
	__DataShape__00000005_ := (&models.DataShape{Name: `"Manage Operational Changes" to "Validate and Accept"-"Manage Operational Changes" to "Validate and Accept"-DiagramProcess`}).Stage(stage)
	__DataShape__00000006_ := (&models.DataShape{Name: `"Manage Operational Changes" to "Validate and Accept"-"Manage Operational Changes" to "Validate and Accept"-DiagramProcess`}).Stage(stage)
	__DataShape__00000007_ := (&models.DataShape{Name: `"Distribute ATS Messages" to "Ingest Flight Data"-"Distribute ATS Messages" to "Ingest Flight Data"-DiagramProcess`}).Stage(stage)
	__DataShape__00000008_ := (&models.DataShape{Name: `"Distribute ATS Messages" to "Ingest Flight Data"-"Distribute ATS Messages" to "Ingest Flight Data"-DiagramProcess`}).Stage(stage)
	__DataShape__00000009_ := (&models.DataShape{Name: `"Distribute ATS Messages" to "Ingest Flight Data"-"Distribute ATS Messages" to "Ingest Flight Data"-DiagramProcess`}).Stage(stage)
	__DataShape__00000010_ := (&models.DataShape{Name: `"Distribute ATS Messages" to "Ingest Flight Data"-"Distribute ATS Messages" to "Ingest Flight Data"-DiagramProcess`}).Stage(stage)

	__DiagramProcess__00000000_ := (&models.DiagramProcess{Name: `DiagramProcess`}).Stage(stage)

	__ExternalParticipantShape__00000000_ := (&models.ExternalParticipantShape{Name: `E1-DiagramProcess`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `ICAO 4444 Chapter 11`}).Stage(stage)

	__Participant__00000000_ := (&models.Participant{Name: `Flight Plan Originator`}).Stage(stage)
	__Participant__00000001_ := (&models.Participant{Name: `Air Traffic Flow Manager`}).Stage(stage)
	__Participant__00000002_ := (&models.Participant{Name: `ATC Unit`}).Stage(stage)
	__Participant__00000003_ := (&models.Participant{Name: `NMUI`}).Stage(stage)

	__ParticipantShape__00000000_ := (&models.ParticipantShape{Name: `Aircraft Operator-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000001_ := (&models.ParticipantShape{Name: `ATS Reporting Office-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000002_ := (&models.ParticipantShape{Name: `ATS Unit-DiagramProcess`}).Stage(stage)

	__Process__00000000_ := (&models.Process{Name: `Flight Messages Exchanges`}).Stage(stage)

	__ProcessShape__00000000_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)

	__Resource__00000001_ := (&models.Resource{Name: `IFPS`}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `Submit Flight Plan`}).Stage(stage)
	__Task__00000001_ := (&models.Task{Name: `Manage Operational Changes`}).Stage(stage)
	__Task__00000002_ := (&models.Task{Name: `Validate and Accept`}).Stage(stage)
	__Task__00000003_ := (&models.Task{Name: `Distribute ATS Messages`}).Stage(stage)
	__Task__00000004_ := (&models.Task{Name: `Ingest Flight Data`}).Stage(stage)
	__Task__00000005_ := (&models.Task{Name: `Log Flight Movement`}).Stage(stage)
	__Task__00000006_ := (&models.Task{Name: ``}).Stage(stage)
	__Task__00000007_ := (&models.Task{Name: `End`}).Stage(stage)

	__TaskShape__00000000_ := (&models.TaskShape{Name: `Submit Flight Plan-DiagramProcess`}).Stage(stage)
	__TaskShape__00000001_ := (&models.TaskShape{Name: `Manage Operational Changes-DiagramProcess`}).Stage(stage)
	__TaskShape__00000002_ := (&models.TaskShape{Name: `Validate and Accept-DiagramProcess`}).Stage(stage)
	__TaskShape__00000003_ := (&models.TaskShape{Name: `Distribute ATS Messages-DiagramProcess`}).Stage(stage)
	__TaskShape__00000004_ := (&models.TaskShape{Name: `Ingest Flight Data-DiagramProcess`}).Stage(stage)
	__TaskShape__00000005_ := (&models.TaskShape{Name: `Log Flight Movement-DiagramProcess`}).Stage(stage)
	__TaskShape__00000006_ := (&models.TaskShape{Name: `-DiagramProcess`}).Stage(stage)
	__TaskShape__00000007_ := (&models.TaskShape{Name: `-DiagramProcess`}).Stage(stage)

	// insertion point for initialization of values

	__AllocatedResourceShape__00000000_.Name = `DiagramProcess-Aircraft Operator-AO`

	__AllocatedResourceShape__00000001_.Name = `DiagramProcess-ATS Reporting Office-ARO`

	__AllocatedResourceShape__00000002_.Name = `DiagramProcess-ATS Unit-ATC`

	__ControlFlow__00000000_.Name = `"Submit Flight Plan" to "Manage Operational Changes"`
	__ControlFlow__00000000_.Description = ``
	__ControlFlow__00000000_.ComputedPrefix = ``
	__ControlFlow__00000000_.IsExpanded = false

	__ControlFlow__00000001_.Name = `"Validate and Accept" to "Distribute ATS Messages"`
	__ControlFlow__00000001_.Description = ``
	__ControlFlow__00000001_.ComputedPrefix = ``
	__ControlFlow__00000001_.IsExpanded = false

	__ControlFlow__00000002_.Name = `"Ingest Flight Data" to "Log Flight Movement"`
	__ControlFlow__00000002_.Description = ``
	__ControlFlow__00000002_.ComputedPrefix = ``
	__ControlFlow__00000002_.IsExpanded = false

	__ControlFlow__00000003_.Name = `"" to "Submit Flight Plan"`
	__ControlFlow__00000003_.Description = ``
	__ControlFlow__00000003_.ComputedPrefix = ``
	__ControlFlow__00000003_.IsExpanded = false

	__ControlFlow__00000004_.Name = `"Log Flight Movement" to "End"`
	__ControlFlow__00000004_.Description = ``
	__ControlFlow__00000004_.ComputedPrefix = ``
	__ControlFlow__00000004_.IsExpanded = false

	__ControlFlowShape__00000000_.Name = `"Submit Flight Plan" to "Manage Operational Changes"`
	__ControlFlowShape__00000000_.StartRatio = 0.449246
	__ControlFlowShape__00000000_.EndRatio = 0.419955
	__ControlFlowShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000000_.CornerOffsetRatio = 1.973017
	__ControlFlowShape__00000000_.IsHidden = false

	__ControlFlowShape__00000001_.Name = `"Validate and Accept" to "Distribute ATS Messages"`
	__ControlFlowShape__00000001_.StartRatio = 0.500000
	__ControlFlowShape__00000001_.EndRatio = 0.500000
	__ControlFlowShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000001_.CornerOffsetRatio = 1.148260
	__ControlFlowShape__00000001_.IsHidden = false

	__ControlFlowShape__00000002_.Name = `"Ingest Flight Data" to "Log Flight Movement"`
	__ControlFlowShape__00000002_.StartRatio = 0.441208
	__ControlFlowShape__00000002_.EndRatio = 0.533504
	__ControlFlowShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000002_.CornerOffsetRatio = 1.500000
	__ControlFlowShape__00000002_.IsHidden = false

	__ControlFlowShape__00000003_.Name = `"" to "Submit Flight Plan"`
	__ControlFlowShape__00000003_.StartRatio = 1.000000
	__ControlFlowShape__00000003_.EndRatio = 0.500000
	__ControlFlowShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000003_.CornerOffsetRatio = 1.837715
	__ControlFlowShape__00000003_.IsHidden = false

	__ControlFlowShape__00000004_.Name = `"Log Flight Movement" to "End"`
	__ControlFlowShape__00000004_.StartRatio = 0.552734
	__ControlFlowShape__00000004_.EndRatio = 0.848395
	__ControlFlowShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000004_.CornerOffsetRatio = 1.500000
	__ControlFlowShape__00000004_.IsHidden = false

	__Data__00000000_.Name = `FPL (incl. item 19)`
	__Data__00000000_.Acronym = ``
	__Data__00000000_.Description = `Filed Flight Plan Message`
	__Data__00000000_.ComputedPrefix = ``
	__Data__00000000_.IsExpanded = false
	__Data__00000000_.SVG_Path = ``
	__Data__00000000_.InverseAppliedScaling = 0.000000

	__Data__00000001_.Name = `CHG`
	__Data__00000001_.Acronym = ``
	__Data__00000001_.Description = `Modification Message`
	__Data__00000001_.ComputedPrefix = ``
	__Data__00000001_.IsExpanded = false
	__Data__00000001_.SVG_Path = ``
	__Data__00000001_.InverseAppliedScaling = 0.000000

	__Data__00000002_.Name = `CNL`
	__Data__00000002_.Acronym = ``
	__Data__00000002_.Description = `Flight Plan Cancellation Message`
	__Data__00000002_.ComputedPrefix = ``
	__Data__00000002_.IsExpanded = false
	__Data__00000002_.SVG_Path = ``
	__Data__00000002_.InverseAppliedScaling = 0.000000

	__Data__00000003_.Name = `DLA`
	__Data__00000003_.Acronym = ``
	__Data__00000003_.Description = `Delay Message`
	__Data__00000003_.ComputedPrefix = ``
	__Data__00000003_.IsExpanded = false
	__Data__00000003_.SVG_Path = ``
	__Data__00000003_.InverseAppliedScaling = 0.000000

	__Data__00000004_.Name = `DEP`
	__Data__00000004_.Acronym = ``
	__Data__00000004_.Description = `Departure Message`
	__Data__00000004_.ComputedPrefix = ``
	__Data__00000004_.IsExpanded = false
	__Data__00000004_.SVG_Path = ``
	__Data__00000004_.InverseAppliedScaling = 0.000000

	__Data__00000005_.Name = `ARR`
	__Data__00000005_.Acronym = ``
	__Data__00000005_.Description = `Arrival Message`
	__Data__00000005_.ComputedPrefix = ``
	__Data__00000005_.IsExpanded = false
	__Data__00000005_.SVG_Path = ``
	__Data__00000005_.InverseAppliedScaling = 0.000000

	__Data__00000006_.Name = `FPL (w/o item 19)`
	__Data__00000006_.Acronym = ``
	__Data__00000006_.Description = ``
	__Data__00000006_.ComputedPrefix = ``
	__Data__00000006_.IsExpanded = false
	__Data__00000006_.SVG_Path = ``
	__Data__00000006_.InverseAppliedScaling = 0.000000

	__Data__00000007_.Name = `RQS`
	__Data__00000007_.Acronym = ``
	__Data__00000007_.Description = ``
	__Data__00000007_.ComputedPrefix = ``
	__Data__00000007_.IsExpanded = false
	__Data__00000007_.SVG_Path = ``
	__Data__00000007_.InverseAppliedScaling = 0.000000

	__Data__00000008_.Name = `SPL`
	__Data__00000008_.Acronym = ``
	__Data__00000008_.Description = ``
	__Data__00000008_.ComputedPrefix = ``
	__Data__00000008_.IsExpanded = false
	__Data__00000008_.SVG_Path = ``
	__Data__00000008_.InverseAppliedScaling = 0.000000

	__DataFlow__00000000_.Name = `"Submit Flight Plan" to "Validate and Accept"`
	__DataFlow__00000000_.Description = ``
	__DataFlow__00000000_.ComputedPrefix = ``
	__DataFlow__00000000_.IsExpanded = false
	__DataFlow__00000000_.Type = models.DataFlow_Task2Task
	__DataFlow__00000000_.IsDatasNodeExpanded = false

	__DataFlow__00000001_.Name = `"Manage Operational Changes" to "Validate and Accept"`
	__DataFlow__00000001_.Description = ``
	__DataFlow__00000001_.ComputedPrefix = ``
	__DataFlow__00000001_.IsExpanded = false
	__DataFlow__00000001_.Type = models.DataFlow_Task2Task
	__DataFlow__00000001_.IsDatasNodeExpanded = false

	__DataFlow__00000002_.Name = `"Distribute ATS Messages" to "Ingest Flight Data"`
	__DataFlow__00000002_.Description = ``
	__DataFlow__00000002_.ComputedPrefix = ``
	__DataFlow__00000002_.IsExpanded = false
	__DataFlow__00000002_.Type = models.DataFlow_Task2Task
	__DataFlow__00000002_.IsDatasNodeExpanded = false

	__DataFlow__00000003_.Name = `"Log Flight Movement" to "Validate and Accept"`
	__DataFlow__00000003_.Description = ``
	__DataFlow__00000003_.ComputedPrefix = ``
	__DataFlow__00000003_.IsExpanded = false
	__DataFlow__00000003_.Type = models.DataFlow_Task2Task
	__DataFlow__00000003_.IsDatasNodeExpanded = false

	__DataFlow__00000004_.Name = `"Log Flight Movement" to "Validate and Accept"`
	__DataFlow__00000004_.Description = ``
	__DataFlow__00000004_.ComputedPrefix = ``
	__DataFlow__00000004_.IsExpanded = false
	__DataFlow__00000004_.Type = models.DataFlow_Task2Task
	__DataFlow__00000004_.IsDatasNodeExpanded = false

	__DataFlowShape__00000000_.Name = `"Submit Flight Plan" to "Validate and Accept"`
	__DataFlowShape__00000000_.StartRatio = 0.500000
	__DataFlowShape__00000000_.EndRatio = 0.204267
	__DataFlowShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000000_.CornerOffsetRatio = 1.023416
	__DataFlowShape__00000000_.IsHidden = false

	__DataFlowShape__00000001_.Name = `"Manage Operational Changes" to "Validate and Accept"`
	__DataFlowShape__00000001_.StartRatio = 0.500000
	__DataFlowShape__00000001_.EndRatio = 0.569109
	__DataFlowShape__00000001_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000001_.CornerOffsetRatio = 1.109519
	__DataFlowShape__00000001_.IsHidden = false

	__DataFlowShape__00000002_.Name = `"Distribute ATS Messages" to "Ingest Flight Data"`
	__DataFlowShape__00000002_.StartRatio = 0.500000
	__DataFlowShape__00000002_.EndRatio = 0.600197
	__DataFlowShape__00000002_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000002_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000002_.CornerOffsetRatio = 1.028348
	__DataFlowShape__00000002_.IsHidden = false

	__DataFlowShape__00000003_.Name = `"Log Flight Movement" to "Validate and Accept"`
	__DataFlowShape__00000003_.StartRatio = 0.731541
	__DataFlowShape__00000003_.EndRatio = 0.723550
	__DataFlowShape__00000003_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000003_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000003_.CornerOffsetRatio = -2.634866
	__DataFlowShape__00000003_.IsHidden = false

	__DataFlowShape__00000004_.Name = `"Log Flight Movement" to "Validate and Accept"`
	__DataFlowShape__00000004_.StartRatio = 0.298208
	__DataFlowShape__00000004_.EndRatio = 0.928183
	__DataFlowShape__00000004_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000004_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000004_.CornerOffsetRatio = -2.468199
	__DataFlowShape__00000004_.IsHidden = false

	__DataShape__00000000_.Name = `"Submit Flight Plan" to "Validate and Accept"-"Submit Flight Plan" to "Validate and Accept"-DiagramProcess`

	__DataShape__00000001_.Name = `"Manage Operational Changes" to "Validate and Accept"-"Manage Operational Changes" to "Validate and Accept"-DiagramProcess`

	__DataShape__00000002_.Name = `"Distribute ATS Messages" to "Ingest Flight Data"-"Distribute ATS Messages" to "Ingest Flight Data"-DiagramProcess`

	__DataShape__00000003_.Name = `"Log Flight Movement" to "Validate and Accept" - DEP-"Log Flight Movement" to "Validate and Accept" - DEP-DiagramProcess`

	__DataShape__00000004_.Name = `"Log Flight Movement" to "Validate and Accept" - ARR-"Log Flight Movement" to "Validate and Accept" - ARR-DiagramProcess`

	__DataShape__00000005_.Name = `"Manage Operational Changes" to "Validate and Accept"-"Manage Operational Changes" to "Validate and Accept"-DiagramProcess`

	__DataShape__00000006_.Name = `"Manage Operational Changes" to "Validate and Accept"-"Manage Operational Changes" to "Validate and Accept"-DiagramProcess`

	__DataShape__00000007_.Name = `"Distribute ATS Messages" to "Ingest Flight Data"-"Distribute ATS Messages" to "Ingest Flight Data"-DiagramProcess`

	__DataShape__00000008_.Name = `"Distribute ATS Messages" to "Ingest Flight Data"-"Distribute ATS Messages" to "Ingest Flight Data"-DiagramProcess`

	__DataShape__00000009_.Name = `"Distribute ATS Messages" to "Ingest Flight Data"-"Distribute ATS Messages" to "Ingest Flight Data"-DiagramProcess`

	__DataShape__00000010_.Name = `"Distribute ATS Messages" to "Ingest Flight Data"-"Distribute ATS Messages" to "Ingest Flight Data"-DiagramProcess`

	__DiagramProcess__00000000_.Name = `DiagramProcess`
	__DiagramProcess__00000000_.Description = ``
	__DiagramProcess__00000000_.ComputedPrefix = ``
	__DiagramProcess__00000000_.IsExpanded = false
	__DiagramProcess__00000000_.IsChecked = true
	__DiagramProcess__00000000_.IsEditable_ = true
	__DiagramProcess__00000000_.IsShowPrefix = false
	__DiagramProcess__00000000_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000000_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000000_.Width = 1716.444502
	__DiagramProcess__00000000_.Height = 1465.000000
	__DiagramProcess__00000000_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000000_.IsParticipantsNodeExpanded = true
	__DiagramProcess__00000000_.IsExternalParticipantsNodeExpanded = true
	__DiagramProcess__00000000_.IsNotesNodeExpanded = false

	__ExternalParticipantShape__00000000_.Name = `E1-DiagramProcess`
	__ExternalParticipantShape__00000000_.IsExpanded = false
	__ExternalParticipantShape__00000000_.X = 1166.444502
	__ExternalParticipantShape__00000000_.Y = 73.671133
	__ExternalParticipantShape__00000000_.Width = 250.000000
	__ExternalParticipantShape__00000000_.Height = 70.000000
	__ExternalParticipantShape__00000000_.IsHidden = false
	__ExternalParticipantShape__00000000_.TailHeigth = 921.999924

	__Library__00000000_.Name = `ICAO 4444 Chapter 11`
	__Library__00000000_.Description = `Modeling of message exchanges handled by the ATS Reporting Office (ARO)`
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsExpanded = false
	__Library__00000000_.IsRootLibrary = true
	__Library__00000000_.IsSubLibrariesNodeExpanded = false
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.IsProcessesNodeExpanded = true
	__Library__00000000_.IsDataFlowsNodeExpanded = true
	__Library__00000000_.IsDatasNodeExpanded = false
	__Library__00000000_.IsResourcesNodeExpanded = false
	__Library__00000000_.IsNotesNodeExpanded = false
	__Library__00000000_.IsExpandedTmp = true

	__Participant__00000000_.Name = `Flight Plan Originator`
	__Participant__00000000_.IsProcessResource = false
	__Participant__00000000_.Description = `Originator of initial flight plan data and operational modifications`
	__Participant__00000000_.IsResourcesNodeExpanded = true
	__Participant__00000000_.IsProcessesNodeExpanded = false
	__Participant__00000000_.ComputedPrefix = ``
	__Participant__00000000_.IsExpanded = false
	__Participant__00000000_.IsTasksNodeExpanded = true
	__Participant__00000000_.IsControlFlowsNodeExpanded = false
	__Participant__00000000_.IsDataFlowsNodeExpanded = false

	__Participant__00000001_.Name = `Air Traffic Flow Manager`
	__Participant__00000001_.IsProcessResource = false
	__Participant__00000001_.Description = `Receives, validates, accepts, and distributes flight plan-related messages`
	__Participant__00000001_.IsResourcesNodeExpanded = false
	__Participant__00000001_.IsProcessesNodeExpanded = false
	__Participant__00000001_.ComputedPrefix = ``
	__Participant__00000001_.IsExpanded = false
	__Participant__00000001_.IsTasksNodeExpanded = true
	__Participant__00000001_.IsControlFlowsNodeExpanded = false
	__Participant__00000001_.IsDataFlowsNodeExpanded = false

	__Participant__00000002_.Name = `ATC Unit`
	__Participant__00000002_.IsProcessResource = false
	__Participant__00000002_.Description = `Recipient of distributed movement and control data used for traffic separation`
	__Participant__00000002_.IsResourcesNodeExpanded = false
	__Participant__00000002_.IsProcessesNodeExpanded = false
	__Participant__00000002_.ComputedPrefix = ``
	__Participant__00000002_.IsExpanded = false
	__Participant__00000002_.IsTasksNodeExpanded = true
	__Participant__00000002_.IsControlFlowsNodeExpanded = false
	__Participant__00000002_.IsDataFlowsNodeExpanded = false

	__Participant__00000003_.Name = `NMUI`
	__Participant__00000003_.IsProcessResource = false
	__Participant__00000003_.Description = ``
	__Participant__00000003_.IsResourcesNodeExpanded = false
	__Participant__00000003_.IsProcessesNodeExpanded = false
	__Participant__00000003_.ComputedPrefix = ``
	__Participant__00000003_.IsExpanded = false
	__Participant__00000003_.IsTasksNodeExpanded = false
	__Participant__00000003_.IsControlFlowsNodeExpanded = false
	__Participant__00000003_.IsDataFlowsNodeExpanded = false

	__ParticipantShape__00000000_.Name = `Aircraft Operator-DiagramProcess`
	__ParticipantShape__00000000_.IsExpanded = false
	__ParticipantShape__00000000_.X = 153.916317
	__ParticipantShape__00000000_.Y = 178.327215
	__ParticipantShape__00000000_.Width = 250.000000
	__ParticipantShape__00000000_.Height = 70.000000
	__ParticipantShape__00000000_.IsHidden = false
	__ParticipantShape__00000000_.WidthWeight = 0.892644

	__ParticipantShape__00000001_.Name = `ATS Reporting Office-DiagramProcess`
	__ParticipantShape__00000001_.IsExpanded = false
	__ParticipantShape__00000001_.X = 108.634825
	__ParticipantShape__00000001_.Y = 100.269741
	__ParticipantShape__00000001_.Width = 250.000000
	__ParticipantShape__00000001_.Height = 70.000000
	__ParticipantShape__00000001_.IsHidden = false
	__ParticipantShape__00000001_.WidthWeight = 1.000000

	__ParticipantShape__00000002_.Name = `ATS Unit-DiagramProcess`
	__ParticipantShape__00000002_.IsExpanded = false
	__ParticipantShape__00000002_.X = 143.845785
	__ParticipantShape__00000002_.Y = 168.047525
	__ParticipantShape__00000002_.Width = 250.000000
	__ParticipantShape__00000002_.Height = 70.000000
	__ParticipantShape__00000002_.IsHidden = false
	__ParticipantShape__00000002_.WidthWeight = 1.000000

	__Process__00000000_.Name = `Flight Messages Exchanges`
	__Process__00000000_.Description = `Process governing the submission, validation, modification, and distribution of ATS messages via the ARO.`
	__Process__00000000_.ComputedPrefix = ``
	__Process__00000000_.IsExpanded = false
	__Process__00000000_.SVG_Path = ``
	__Process__00000000_.InverseAppliedScaling = 0.000000
	__Process__00000000_.IsSubProcessNodeExpanded = false
	__Process__00000000_.IsDataFlowsNodeExpanded = true

	__ProcessShape__00000000_.Name = `ProcessShape`
	__ProcessShape__00000000_.IsExpanded = false
	__ProcessShape__00000000_.X = 86.000000
	__ProcessShape__00000000_.Y = 39.000000
	__ProcessShape__00000000_.Width = 1020.000000
	__ProcessShape__00000000_.Height = 1126.000000
	__ProcessShape__00000000_.IsHidden = false

	__Resource__00000001_.Name = `IFPS`
	__Resource__00000001_.Acronym = ``
	__Resource__00000001_.Description = `Euroncontrol's Initial Flight Plan Processing`
	__Resource__00000001_.ComputedPrefix = ``
	__Resource__00000001_.IsExpanded = false
	__Resource__00000001_.LayoutDirection = models.Vertical
	__Resource__00000001_.SVG_Path = ``
	__Resource__00000001_.InverseAppliedScaling = 0.000000

	__Task__00000000_.Name = `Submit Flight Plan`
	__Task__00000000_.Description = `Originates and files the initial field flight plan parameters before departure.`
	__Task__00000000_.ComputedPrefix = ``
	__Task__00000000_.IsExpanded = false
	__Task__00000000_.LayoutDirection = models.Vertical
	__Task__00000000_.IsStartTask = false
	__Task__00000000_.IsEndTask = false
	__Task__00000000_.IsTaskNameNotProcessName = true

	__Task__00000001_.Name = `Manage Operational Changes`
	__Task__00000001_.Description = `Monitors operational status and decides if a flight needs to be delayed, modified, or aborted prior to departure.`
	__Task__00000001_.ComputedPrefix = ``
	__Task__00000001_.IsExpanded = false
	__Task__00000001_.LayoutDirection = models.Vertical
	__Task__00000001_.IsStartTask = false
	__Task__00000001_.IsEndTask = false
	__Task__00000001_.IsTaskNameNotProcessName = true

	__Task__00000002_.Name = `Validate and Accept`
	__Task__00000002_.Description = `Receives incoming submissions and checks them for syntax, formatting correctness, and logical consistency according to ICAO standards.`
	__Task__00000002_.ComputedPrefix = ``
	__Task__00000002_.IsExpanded = false
	__Task__00000002_.LayoutDirection = models.Vertical
	__Task__00000002_.IsStartTask = false
	__Task__00000002_.IsEndTask = false
	__Task__00000002_.IsTaskNameNotProcessName = true

	__Task__00000003_.Name = `Distribute ATS Messages`
	__Task__00000003_.Description = `Determines the appropriate addressing routing and distributes accepted flight plans or updates to downstream units.`
	__Task__00000003_.ComputedPrefix = ``
	__Task__00000003_.IsExpanded = false
	__Task__00000003_.LayoutDirection = models.Vertical
	__Task__00000003_.IsStartTask = false
	__Task__00000003_.IsEndTask = false
	__Task__00000003_.IsTaskNameNotProcessName = true

	__Task__00000004_.Name = `Ingest Flight Data`
	__Task__00000004_.Description = `Receives distributed data and loads it into the local Flight Data Processing System (FDPS) for air traffic controllers.`
	__Task__00000004_.ComputedPrefix = ``
	__Task__00000004_.IsExpanded = false
	__Task__00000004_.LayoutDirection = models.Vertical
	__Task__00000004_.IsStartTask = false
	__Task__00000004_.IsEndTask = false
	__Task__00000004_.IsTaskNameNotProcessName = true

	__Task__00000005_.Name = `Log Flight Movement`
	__Task__00000005_.Description = `Tracks physical execution milestones (actual departure and arrival times) and generates status updates.`
	__Task__00000005_.ComputedPrefix = ``
	__Task__00000005_.IsExpanded = false
	__Task__00000005_.LayoutDirection = models.Vertical
	__Task__00000005_.IsStartTask = false
	__Task__00000005_.IsEndTask = false
	__Task__00000005_.IsTaskNameNotProcessName = true

	__Task__00000006_.Name = ``
	__Task__00000006_.Description = ``
	__Task__00000006_.ComputedPrefix = ``
	__Task__00000006_.IsExpanded = false
	__Task__00000006_.LayoutDirection = models.Vertical
	__Task__00000006_.IsStartTask = true
	__Task__00000006_.IsEndTask = false
	__Task__00000006_.IsTaskNameNotProcessName = false

	__Task__00000007_.Name = `End`
	__Task__00000007_.Description = ``
	__Task__00000007_.ComputedPrefix = ``
	__Task__00000007_.IsExpanded = false
	__Task__00000007_.LayoutDirection = models.Vertical
	__Task__00000007_.IsStartTask = false
	__Task__00000007_.IsEndTask = true
	__Task__00000007_.IsTaskNameNotProcessName = false

	__TaskShape__00000000_.Name = `Submit Flight Plan-DiagramProcess`
	__TaskShape__00000000_.IsExpanded = false
	__TaskShape__00000000_.X = 121.000000
	__TaskShape__00000000_.Y = 300.795175
	__TaskShape__00000000_.Width = 240.333333
	__TaskShape__00000000_.Height = 48.000000
	__TaskShape__00000000_.IsHidden = false

	__TaskShape__00000001_.Name = `Manage Operational Changes-DiagramProcess`
	__TaskShape__00000001_.IsExpanded = false
	__TaskShape__00000001_.X = 122.000000
	__TaskShape__00000001_.Y = 400.907212
	__TaskShape__00000001_.Width = 240.333333
	__TaskShape__00000001_.Height = 70.000000
	__TaskShape__00000001_.IsHidden = false

	__TaskShape__00000002_.Name = `Validate and Accept-DiagramProcess`
	__TaskShape__00000002_.IsExpanded = false
	__TaskShape__00000002_.X = 492.000000
	__TaskShape__00000002_.Y = 283.100648
	__TaskShape__00000002_.Width = 189.333333
	__TaskShape__00000002_.Height = 259.000000
	__TaskShape__00000002_.IsHidden = false

	__TaskShape__00000003_.Name = `Distribute ATS Messages-DiagramProcess`
	__TaskShape__00000003_.IsExpanded = false
	__TaskShape__00000003_.X = 492.000000
	__TaskShape__00000003_.Y = 602.893285
	__TaskShape__00000003_.Width = 175.000000
	__TaskShape__00000003_.Height = 70.000000
	__TaskShape__00000003_.IsHidden = false

	__TaskShape__00000004_.Name = `Ingest Flight Data-DiagramProcess`
	__TaskShape__00000004_.IsExpanded = false
	__TaskShape__00000004_.X = 845.666667
	__TaskShape__00000004_.Y = 577.382953
	__TaskShape__00000004_.Width = 160.000000
	__TaskShape__00000004_.Height = 116.000000
	__TaskShape__00000004_.IsHidden = false

	__TaskShape__00000005_.Name = `Log Flight Movement-DiagramProcess`
	__TaskShape__00000005_.IsExpanded = false
	__TaskShape__00000005_.X = 840.000000
	__TaskShape__00000005_.Y = 789.607545
	__TaskShape__00000005_.Width = 156.000000
	__TaskShape__00000005_.Height = 60.000000
	__TaskShape__00000005_.IsHidden = false

	__TaskShape__00000006_.Name = `-DiagramProcess`
	__TaskShape__00000006_.IsExpanded = false
	__TaskShape__00000006_.X = 149.000000
	__TaskShape__00000006_.Y = 233.251563
	__TaskShape__00000006_.Width = 83.000000
	__TaskShape__00000006_.Height = 20.000000
	__TaskShape__00000006_.IsHidden = false

	__TaskShape__00000007_.Name = `-DiagramProcess`
	__TaskShape__00000007_.IsExpanded = false
	__TaskShape__00000007_.X = 845.666667
	__TaskShape__00000007_.Y = 950.033596
	__TaskShape__00000007_.Width = 97.000000
	__TaskShape__00000007_.Height = 36.000000
	__TaskShape__00000007_.IsHidden = false

	// insertion point for setup of pointers
	__AllocatedResourceShape__00000000_.Participant = __Participant__00000000_
	__AllocatedResourceShape__00000000_.Resource = nil
	__AllocatedResourceShape__00000001_.Participant = __Participant__00000001_
	__AllocatedResourceShape__00000001_.Resource = __Resource__00000001_
	__AllocatedResourceShape__00000002_.Participant = __Participant__00000002_
	__AllocatedResourceShape__00000002_.Resource = nil
	__ControlFlow__00000000_.Start = __Task__00000000_
	__ControlFlow__00000000_.End = __Task__00000001_
	__ControlFlow__00000001_.Start = __Task__00000002_
	__ControlFlow__00000001_.End = __Task__00000003_
	__ControlFlow__00000002_.Start = __Task__00000004_
	__ControlFlow__00000002_.End = __Task__00000005_
	__ControlFlow__00000003_.Start = __Task__00000006_
	__ControlFlow__00000003_.End = __Task__00000000_
	__ControlFlow__00000004_.Start = __Task__00000005_
	__ControlFlow__00000004_.End = __Task__00000007_
	__ControlFlowShape__00000000_.ControlFlow = __ControlFlow__00000000_
	__ControlFlowShape__00000001_.ControlFlow = __ControlFlow__00000001_
	__ControlFlowShape__00000002_.ControlFlow = __ControlFlow__00000002_
	__ControlFlowShape__00000003_.ControlFlow = __ControlFlow__00000003_
	__ControlFlowShape__00000004_.ControlFlow = __ControlFlow__00000004_
	__DataFlow__00000000_.Datas = append(__DataFlow__00000000_.Datas, __Data__00000000_)
	__DataFlow__00000000_.StartTask = __Task__00000000_
	__DataFlow__00000000_.EndTask = __Task__00000002_
	__DataFlow__00000000_.StartExternalParticipant = nil
	__DataFlow__00000000_.EndExternalParticipant = nil
	__DataFlow__00000001_.Datas = append(__DataFlow__00000001_.Datas, __Data__00000001_)
	__DataFlow__00000001_.Datas = append(__DataFlow__00000001_.Datas, __Data__00000003_)
	__DataFlow__00000001_.Datas = append(__DataFlow__00000001_.Datas, __Data__00000002_)
	__DataFlow__00000001_.StartTask = __Task__00000001_
	__DataFlow__00000001_.EndTask = __Task__00000002_
	__DataFlow__00000001_.StartExternalParticipant = nil
	__DataFlow__00000001_.EndExternalParticipant = nil
	__DataFlow__00000002_.Datas = append(__DataFlow__00000002_.Datas, __Data__00000001_)
	__DataFlow__00000002_.Datas = append(__DataFlow__00000002_.Datas, __Data__00000003_)
	__DataFlow__00000002_.Datas = append(__DataFlow__00000002_.Datas, __Data__00000002_)
	__DataFlow__00000002_.Datas = append(__DataFlow__00000002_.Datas, __Data__00000006_)
	__DataFlow__00000002_.StartTask = __Task__00000003_
	__DataFlow__00000002_.EndTask = __Task__00000004_
	__DataFlow__00000002_.StartExternalParticipant = nil
	__DataFlow__00000002_.EndExternalParticipant = nil
	__DataFlow__00000003_.Datas = append(__DataFlow__00000003_.Datas, __Data__00000004_)
	__DataFlow__00000003_.StartTask = __Task__00000005_
	__DataFlow__00000003_.EndTask = __Task__00000002_
	__DataFlow__00000003_.StartExternalParticipant = nil
	__DataFlow__00000003_.EndExternalParticipant = nil
	__DataFlow__00000004_.Datas = append(__DataFlow__00000004_.Datas, __Data__00000005_)
	__DataFlow__00000004_.StartTask = __Task__00000005_
	__DataFlow__00000004_.EndTask = __Task__00000002_
	__DataFlow__00000004_.StartExternalParticipant = nil
	__DataFlow__00000004_.EndExternalParticipant = nil
	__DataFlowShape__00000000_.DataFlow = __DataFlow__00000000_
	__DataFlowShape__00000001_.DataFlow = __DataFlow__00000001_
	__DataFlowShape__00000002_.DataFlow = __DataFlow__00000002_
	__DataFlowShape__00000003_.DataFlow = __DataFlow__00000003_
	__DataFlowShape__00000004_.DataFlow = __DataFlow__00000004_
	__DataShape__00000000_.Data = __Data__00000000_
	__DataShape__00000000_.DataFlow = __DataFlow__00000000_
	__DataShape__00000001_.Data = __Data__00000001_
	__DataShape__00000001_.DataFlow = __DataFlow__00000001_
	__DataShape__00000002_.Data = __Data__00000000_
	__DataShape__00000002_.DataFlow = __DataFlow__00000002_
	__DataShape__00000003_.Data = __Data__00000004_
	__DataShape__00000003_.DataFlow = __DataFlow__00000003_
	__DataShape__00000004_.Data = __Data__00000005_
	__DataShape__00000004_.DataFlow = __DataFlow__00000004_
	__DataShape__00000005_.Data = __Data__00000003_
	__DataShape__00000005_.DataFlow = __DataFlow__00000001_
	__DataShape__00000006_.Data = __Data__00000002_
	__DataShape__00000006_.DataFlow = __DataFlow__00000001_
	__DataShape__00000007_.Data = __Data__00000001_
	__DataShape__00000007_.DataFlow = __DataFlow__00000002_
	__DataShape__00000008_.Data = __Data__00000003_
	__DataShape__00000008_.DataFlow = __DataFlow__00000002_
	__DataShape__00000009_.Data = __Data__00000002_
	__DataShape__00000009_.DataFlow = __DataFlow__00000002_
	__DataShape__00000010_.Data = __Data__00000006_
	__DataShape__00000010_.DataFlow = __DataFlow__00000002_
	__DiagramProcess__00000000_.Process_Shapes = append(__DiagramProcess__00000000_.Process_Shapes, __ProcessShape__00000000_)
	__DiagramProcess__00000000_.Participant_Shapes = append(__DiagramProcess__00000000_.Participant_Shapes, __ParticipantShape__00000000_)
	__DiagramProcess__00000000_.Participant_Shapes = append(__DiagramProcess__00000000_.Participant_Shapes, __ParticipantShape__00000001_)
	__DiagramProcess__00000000_.Participant_Shapes = append(__DiagramProcess__00000000_.Participant_Shapes, __ParticipantShape__00000002_)
	__DiagramProcess__00000000_.ParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000000_.ParticipantWhoseNodeIsExpanded, __Participant__00000002_)
	__DiagramProcess__00000000_.ExternalParticipant_Shapes = append(__DiagramProcess__00000000_.ExternalParticipant_Shapes, __ExternalParticipantShape__00000000_)
	__DiagramProcess__00000000_.TasksWhoseNodeIsExpanded = append(__DiagramProcess__00000000_.TasksWhoseNodeIsExpanded, __Task__00000003_)
	__DiagramProcess__00000000_.TasksWhoseNodeIsExpanded = append(__DiagramProcess__00000000_.TasksWhoseNodeIsExpanded, __Task__00000007_)
	__DiagramProcess__00000000_.Task_Shapes = append(__DiagramProcess__00000000_.Task_Shapes, __TaskShape__00000000_)
	__DiagramProcess__00000000_.Task_Shapes = append(__DiagramProcess__00000000_.Task_Shapes, __TaskShape__00000001_)
	__DiagramProcess__00000000_.Task_Shapes = append(__DiagramProcess__00000000_.Task_Shapes, __TaskShape__00000002_)
	__DiagramProcess__00000000_.Task_Shapes = append(__DiagramProcess__00000000_.Task_Shapes, __TaskShape__00000003_)
	__DiagramProcess__00000000_.Task_Shapes = append(__DiagramProcess__00000000_.Task_Shapes, __TaskShape__00000004_)
	__DiagramProcess__00000000_.Task_Shapes = append(__DiagramProcess__00000000_.Task_Shapes, __TaskShape__00000005_)
	__DiagramProcess__00000000_.Task_Shapes = append(__DiagramProcess__00000000_.Task_Shapes, __TaskShape__00000006_)
	__DiagramProcess__00000000_.Task_Shapes = append(__DiagramProcess__00000000_.Task_Shapes, __TaskShape__00000007_)
	__DiagramProcess__00000000_.ControlFlow_Shapes = append(__DiagramProcess__00000000_.ControlFlow_Shapes, __ControlFlowShape__00000000_)
	__DiagramProcess__00000000_.ControlFlow_Shapes = append(__DiagramProcess__00000000_.ControlFlow_Shapes, __ControlFlowShape__00000001_)
	__DiagramProcess__00000000_.ControlFlow_Shapes = append(__DiagramProcess__00000000_.ControlFlow_Shapes, __ControlFlowShape__00000002_)
	__DiagramProcess__00000000_.ControlFlow_Shapes = append(__DiagramProcess__00000000_.ControlFlow_Shapes, __ControlFlowShape__00000003_)
	__DiagramProcess__00000000_.ControlFlow_Shapes = append(__DiagramProcess__00000000_.ControlFlow_Shapes, __ControlFlowShape__00000004_)
	__DiagramProcess__00000000_.DataFlow_Shapes = append(__DiagramProcess__00000000_.DataFlow_Shapes, __DataFlowShape__00000000_)
	__DiagramProcess__00000000_.DataFlow_Shapes = append(__DiagramProcess__00000000_.DataFlow_Shapes, __DataFlowShape__00000001_)
	__DiagramProcess__00000000_.DataFlow_Shapes = append(__DiagramProcess__00000000_.DataFlow_Shapes, __DataFlowShape__00000002_)
	__DiagramProcess__00000000_.DataFlow_Shapes = append(__DiagramProcess__00000000_.DataFlow_Shapes, __DataFlowShape__00000003_)
	__DiagramProcess__00000000_.DataFlow_Shapes = append(__DiagramProcess__00000000_.DataFlow_Shapes, __DataFlowShape__00000004_)
	__DiagramProcess__00000000_.Data_Shapes = append(__DiagramProcess__00000000_.Data_Shapes, __DataShape__00000000_)
	__DiagramProcess__00000000_.Data_Shapes = append(__DiagramProcess__00000000_.Data_Shapes, __DataShape__00000001_)
	__DiagramProcess__00000000_.Data_Shapes = append(__DiagramProcess__00000000_.Data_Shapes, __DataShape__00000002_)
	__DiagramProcess__00000000_.Data_Shapes = append(__DiagramProcess__00000000_.Data_Shapes, __DataShape__00000003_)
	__DiagramProcess__00000000_.Data_Shapes = append(__DiagramProcess__00000000_.Data_Shapes, __DataShape__00000004_)
	__DiagramProcess__00000000_.Data_Shapes = append(__DiagramProcess__00000000_.Data_Shapes, __DataShape__00000005_)
	__DiagramProcess__00000000_.Data_Shapes = append(__DiagramProcess__00000000_.Data_Shapes, __DataShape__00000006_)
	__DiagramProcess__00000000_.Data_Shapes = append(__DiagramProcess__00000000_.Data_Shapes, __DataShape__00000007_)
	__DiagramProcess__00000000_.Data_Shapes = append(__DiagramProcess__00000000_.Data_Shapes, __DataShape__00000008_)
	__DiagramProcess__00000000_.Data_Shapes = append(__DiagramProcess__00000000_.Data_Shapes, __DataShape__00000009_)
	__DiagramProcess__00000000_.Data_Shapes = append(__DiagramProcess__00000000_.Data_Shapes, __DataShape__00000010_)
	__DiagramProcess__00000000_.DataFlowsWhoseDataNodeIsExpanded = append(__DiagramProcess__00000000_.DataFlowsWhoseDataNodeIsExpanded, __DataFlow__00000000_)
	__DiagramProcess__00000000_.DataFlowsWhoseDataNodeIsExpanded = append(__DiagramProcess__00000000_.DataFlowsWhoseDataNodeIsExpanded, __DataFlow__00000001_)
	__DiagramProcess__00000000_.DataFlowsWhoseDataNodeIsExpanded = append(__DiagramProcess__00000000_.DataFlowsWhoseDataNodeIsExpanded, __DataFlow__00000002_)
	__DiagramProcess__00000000_.DataFlowsWhoseDataNodeIsExpanded = append(__DiagramProcess__00000000_.DataFlowsWhoseDataNodeIsExpanded, __DataFlow__00000003_)
	__DiagramProcess__00000000_.DataFlowsWhoseDataNodeIsExpanded = append(__DiagramProcess__00000000_.DataFlowsWhoseDataNodeIsExpanded, __DataFlow__00000004_)
	__DiagramProcess__00000000_.AllocatedResourceShapes = append(__DiagramProcess__00000000_.AllocatedResourceShapes, __AllocatedResourceShape__00000000_)
	__DiagramProcess__00000000_.AllocatedResourceShapes = append(__DiagramProcess__00000000_.AllocatedResourceShapes, __AllocatedResourceShape__00000001_)
	__DiagramProcess__00000000_.AllocatedResourceShapes = append(__DiagramProcess__00000000_.AllocatedResourceShapes, __AllocatedResourceShape__00000002_)
	__ExternalParticipantShape__00000000_.Participant = __Participant__00000003_
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000000_)
	__Library__00000000_.ProcesssWhoseNodeIsExpanded = append(__Library__00000000_.ProcesssWhoseNodeIsExpanded, __Process__00000000_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000000_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000001_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000002_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000003_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000004_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000000_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000006_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000001_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000002_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000003_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000004_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000005_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000007_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000008_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000001_)
	__Participant__00000000_.Tasks = append(__Participant__00000000_.Tasks, __Task__00000000_)
	__Participant__00000000_.Tasks = append(__Participant__00000000_.Tasks, __Task__00000001_)
	__Participant__00000000_.Tasks = append(__Participant__00000000_.Tasks, __Task__00000006_)
	__Participant__00000000_.ControlFlows = append(__Participant__00000000_.ControlFlows, __ControlFlow__00000000_)
	__Participant__00000000_.ControlFlows = append(__Participant__00000000_.ControlFlows, __ControlFlow__00000003_)
	__Participant__00000000_.TaskWhoseOutDataFlowsNodeIsExpanded = append(__Participant__00000000_.TaskWhoseOutDataFlowsNodeIsExpanded, __Task__00000000_)
	__Participant__00000000_.TaskWhoseOutDataFlowsNodeIsExpanded = append(__Participant__00000000_.TaskWhoseOutDataFlowsNodeIsExpanded, __Task__00000001_)
	__Participant__00000001_.Resources = append(__Participant__00000001_.Resources, __Resource__00000001_)
	__Participant__00000001_.Tasks = append(__Participant__00000001_.Tasks, __Task__00000002_)
	__Participant__00000001_.Tasks = append(__Participant__00000001_.Tasks, __Task__00000003_)
	__Participant__00000001_.ControlFlows = append(__Participant__00000001_.ControlFlows, __ControlFlow__00000001_)
	__Participant__00000001_.TaskWhoseOutDataFlowsNodeIsExpanded = append(__Participant__00000001_.TaskWhoseOutDataFlowsNodeIsExpanded, __Task__00000003_)
	__Participant__00000002_.Tasks = append(__Participant__00000002_.Tasks, __Task__00000004_)
	__Participant__00000002_.Tasks = append(__Participant__00000002_.Tasks, __Task__00000005_)
	__Participant__00000002_.Tasks = append(__Participant__00000002_.Tasks, __Task__00000007_)
	__Participant__00000002_.ControlFlows = append(__Participant__00000002_.ControlFlows, __ControlFlow__00000002_)
	__Participant__00000002_.ControlFlows = append(__Participant__00000002_.ControlFlows, __ControlFlow__00000004_)
	__Participant__00000002_.TaskWhoseOutDataFlowsNodeIsExpanded = append(__Participant__00000002_.TaskWhoseOutDataFlowsNodeIsExpanded, __Task__00000005_)
	__ParticipantShape__00000000_.Participant = __Participant__00000000_
	__ParticipantShape__00000001_.Participant = __Participant__00000001_
	__ParticipantShape__00000002_.Participant = __Participant__00000002_
	__Process__00000000_.DiagramProcesss = append(__Process__00000000_.DiagramProcesss, __DiagramProcess__00000000_)
	__Process__00000000_.DiagramProcessWhoseNodeIsExpanded = append(__Process__00000000_.DiagramProcessWhoseNodeIsExpanded, __DiagramProcess__00000000_)
	__Process__00000000_.Participants = append(__Process__00000000_.Participants, __Participant__00000000_)
	__Process__00000000_.Participants = append(__Process__00000000_.Participants, __Participant__00000001_)
	__Process__00000000_.Participants = append(__Process__00000000_.Participants, __Participant__00000002_)
	__Process__00000000_.DataFlows = append(__Process__00000000_.DataFlows, __DataFlow__00000000_)
	__Process__00000000_.DataFlows = append(__Process__00000000_.DataFlows, __DataFlow__00000001_)
	__Process__00000000_.DataFlows = append(__Process__00000000_.DataFlows, __DataFlow__00000002_)
	__Process__00000000_.DataFlows = append(__Process__00000000_.DataFlows, __DataFlow__00000003_)
	__Process__00000000_.DataFlows = append(__Process__00000000_.DataFlows, __DataFlow__00000004_)
	__Process__00000000_.ExternalParticipants = append(__Process__00000000_.ExternalParticipants, __Participant__00000003_)
	__ProcessShape__00000000_.Process = __Process__00000000_
	__Task__00000000_.Type = nil
	__Task__00000001_.Type = nil
	__Task__00000002_.Type = nil
	__Task__00000003_.Type = nil
	__Task__00000004_.Type = nil
	__Task__00000005_.Type = nil
	__Task__00000006_.Type = nil
	__Task__00000007_.Type = nil
	__TaskShape__00000000_.Task = __Task__00000000_
	__TaskShape__00000001_.Task = __Task__00000001_
	__TaskShape__00000002_.Task = __Task__00000002_
	__TaskShape__00000003_.Task = __Task__00000003_
	__TaskShape__00000004_.Task = __Task__00000004_
	__TaskShape__00000005_.Task = __Task__00000005_
	__TaskShape__00000006_.Task = __Task__00000006_
	__TaskShape__00000007_.Task = __Task__00000007_
}
