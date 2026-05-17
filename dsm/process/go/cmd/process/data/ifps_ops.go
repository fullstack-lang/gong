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
	__AllocatedResourceShape__00000003_ := (&models.AllocatedResourceShape{Name: `DiagramProcess-ATS Unit-ATC`}).Stage(stage)
	__AllocatedResourceShape__00000004_ := (&models.AllocatedResourceShape{Name: `DiagramProcess-EUROCONTROL IFPS-IFPS`}).Stage(stage)
	__AllocatedResourceShape__00000005_ := (&models.AllocatedResourceShape{Name: `DiagramProcess-Aircraft Operator-AO`}).Stage(stage)
	__AllocatedResourceShape__00000006_ := (&models.AllocatedResourceShape{Name: `DiagramProcess-ATS Unit-ATC`}).Stage(stage)
	__AllocatedResourceShape__00000007_ := (&models.AllocatedResourceShape{Name: `DiagramProcess-EUROCONTROL IFPS-IFPS`}).Stage(stage)
	__AllocatedResourceShape__00000008_ := (&models.AllocatedResourceShape{Name: `DiagramProcess-Aircraft Operator-AO`}).Stage(stage)

	__ControlFlow__00000000_ := (&models.ControlFlow{Name: `"Submit Flight Plan" to "Manage Operational Changes"`}).Stage(stage)
	__ControlFlow__00000001_ := (&models.ControlFlow{Name: `"Validate and Accept" to "Distribute ATS Messages"`}).Stage(stage)
	__ControlFlow__00000002_ := (&models.ControlFlow{Name: `"Ingest Flight Data" to "Log Flight Movement"`}).Stage(stage)
	__ControlFlow__00000003_ := (&models.ControlFlow{Name: `"" to "Submit Flight Plan"`}).Stage(stage)
	__ControlFlow__00000004_ := (&models.ControlFlow{Name: `"Log Flight Movement" to ""`}).Stage(stage)
	__ControlFlow__00000005_ := (&models.ControlFlow{Name: `"" to "Transmits RQS"`}).Stage(stage)
	__ControlFlow__00000006_ := (&models.ControlFlow{Name: `"Receives SPL" to ""`}).Stage(stage)
	__ControlFlow__00000007_ := (&models.ControlFlow{Name: `"" to "Transmits RQS"`}).Stage(stage)
	__ControlFlow__00000008_ := (&models.ControlFlow{Name: `"Receives SPL" to ""`}).Stage(stage)

	__ControlFlowShape__00000000_ := (&models.ControlFlowShape{Name: `"Submit Flight Plan" to "Manage Operational Changes"`}).Stage(stage)
	__ControlFlowShape__00000001_ := (&models.ControlFlowShape{Name: `"Validate and Accept" to "Distribute ATS Messages"`}).Stage(stage)
	__ControlFlowShape__00000002_ := (&models.ControlFlowShape{Name: `"Ingest Flight Data" to "Log Flight Movement"`}).Stage(stage)
	__ControlFlowShape__00000003_ := (&models.ControlFlowShape{Name: `"" to "Submit Flight Plan"`}).Stage(stage)
	__ControlFlowShape__00000004_ := (&models.ControlFlowShape{Name: `"Log Flight Movement" to ""`}).Stage(stage)
	__ControlFlowShape__00000005_ := (&models.ControlFlowShape{Name: `"Receives SPL" to ""`}).Stage(stage)
	__ControlFlowShape__00000006_ := (&models.ControlFlowShape{Name: `"" to "Transmits RQS"`}).Stage(stage)
	__ControlFlowShape__00000007_ := (&models.ControlFlowShape{Name: `"Receives SPL" to ""`}).Stage(stage)
	__ControlFlowShape__00000008_ := (&models.ControlFlowShape{Name: `"" to "Transmits RQS"`}).Stage(stage)

	__Data__00000000_ := (&models.Data{Name: `FPL (incl. item 19)`}).Stage(stage)
	__Data__00000001_ := (&models.Data{Name: `CHG`}).Stage(stage)
	__Data__00000002_ := (&models.Data{Name: `CNL`}).Stage(stage)
	__Data__00000003_ := (&models.Data{Name: `DLA`}).Stage(stage)
	__Data__00000004_ := (&models.Data{Name: `DEP`}).Stage(stage)
	__Data__00000005_ := (&models.Data{Name: `ARR`}).Stage(stage)
	__Data__00000006_ := (&models.Data{Name: `FPL (w/o item 19)`}).Stage(stage)
	__Data__00000007_ := (&models.Data{Name: `RQS`}).Stage(stage)
	__Data__00000008_ := (&models.Data{Name: `SPL`}).Stage(stage)
	__Data__00000009_ := (&models.Data{Name: `RQS`}).Stage(stage)
	__Data__00000010_ := (&models.Data{Name: `SPL`}).Stage(stage)
	__Data__00000011_ := (&models.Data{Name: `RQS`}).Stage(stage)
	__Data__00000012_ := (&models.Data{Name: `SPL`}).Stage(stage)

	__DataFlow__00000000_ := (&models.DataFlow{Name: `"Submit Flight Plan" to "Validate and Accept"`}).Stage(stage)
	__DataFlow__00000001_ := (&models.DataFlow{Name: `"Manage Operational Changes" to "Validate and Accept"`}).Stage(stage)
	__DataFlow__00000002_ := (&models.DataFlow{Name: `"Distribute ATS Messages" to "Ingest Flight Data"`}).Stage(stage)
	__DataFlow__00000003_ := (&models.DataFlow{Name: `"Log Flight Movement" to "Validate and Accept"`}).Stage(stage)
	__DataFlow__00000004_ := (&models.DataFlow{Name: `"Log Flight Movement" to "Validate and Accept"`}).Stage(stage)
	__DataFlow__00000005_ := (&models.DataFlow{Name: `"Transmits RQS" to "Validates and Routes RQS"`}).Stage(stage)
	__DataFlow__00000006_ := (&models.DataFlow{Name: `"Validates and Routes RQS" to "Compiles and Transmits SPL"`}).Stage(stage)
	__DataFlow__00000007_ := (&models.DataFlow{Name: `"Compiles and Transmits SPL" to "Validates and Routes SPL"`}).Stage(stage)
	__DataFlow__00000008_ := (&models.DataFlow{Name: `"Validates and Routes SPL" to "Receives SPL"`}).Stage(stage)
	__DataFlow__00000009_ := (&models.DataFlow{Name: `"Transmits RQS" to "Validates and Transmit SPL"`}).Stage(stage)
	__DataFlow__00000010_ := (&models.DataFlow{Name: `"Validates and Transmit SPL" to "Receives SPL"`}).Stage(stage)

	__DataFlowShape__00000000_ := (&models.DataFlowShape{Name: `"Submit Flight Plan" to "Validate and Accept"`}).Stage(stage)
	__DataFlowShape__00000001_ := (&models.DataFlowShape{Name: `"Manage Operational Changes" to "Validate and Accept"`}).Stage(stage)
	__DataFlowShape__00000002_ := (&models.DataFlowShape{Name: `"Distribute ATS Messages" to "Ingest Flight Data"`}).Stage(stage)
	__DataFlowShape__00000003_ := (&models.DataFlowShape{Name: `"Log Flight Movement" to "Validate and Accept"`}).Stage(stage)
	__DataFlowShape__00000004_ := (&models.DataFlowShape{Name: `"Log Flight Movement" to "Validate and Accept"`}).Stage(stage)
	__DataFlowShape__00000005_ := (&models.DataFlowShape{Name: `"Transmits RQS" to "Validates and Routes RQS"`}).Stage(stage)
	__DataFlowShape__00000006_ := (&models.DataFlowShape{Name: `"Validates and Routes SPL" to "Receives SPL"`}).Stage(stage)
	__DataFlowShape__00000007_ := (&models.DataFlowShape{Name: `"Validates and Routes RQS" to "Compiles and Transmits SPL"`}).Stage(stage)
	__DataFlowShape__00000008_ := (&models.DataFlowShape{Name: `"Compiles and Transmits SPL" to "Validates and Routes SPL"`}).Stage(stage)
	__DataFlowShape__00000009_ := (&models.DataFlowShape{Name: `"Transmits RQS" to "Validates and Transmit SPL"`}).Stage(stage)
	__DataFlowShape__00000010_ := (&models.DataFlowShape{Name: `"Validates and Transmit SPL" to "Receives SPL"`}).Stage(stage)

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
	__DataShape__00000011_ := (&models.DataShape{Name: `"Validates and Routes RQS" to "Compiles and Transmits SPL" - RQS`}).Stage(stage)
	__DataShape__00000012_ := (&models.DataShape{Name: `"Compiles and Transmits SPL" to "Validates and Routes SPL" - SPL`}).Stage(stage)
	__DataShape__00000013_ := (&models.DataShape{Name: `"Validates and Routes SPL" to "Receives SPL" - SPL`}).Stage(stage)
	__DataShape__00000014_ := (&models.DataShape{Name: `"Transmits RQS" to "Validates and Routes RQS"-"Transmits RQS" to "Validates and Routes RQS"-DiagramProcess`}).Stage(stage)
	__DataShape__00000015_ := (&models.DataShape{Name: `"Validates and Routes RQS" to "Compiles and Transmits SPL" - RQS`}).Stage(stage)
	__DataShape__00000016_ := (&models.DataShape{Name: `"Compiles and Transmits SPL" to "Validates and Routes SPL" - SPL`}).Stage(stage)
	__DataShape__00000017_ := (&models.DataShape{Name: `"Validates and Routes SPL" to "Receives SPL" - SPL`}).Stage(stage)
	__DataShape__00000018_ := (&models.DataShape{Name: `"Transmits RQS" to "Validates and Routes RQS"-"Transmits RQS" to "Validates and Routes RQS"-DiagramProcess`}).Stage(stage)

	__DiagramProcess__00000000_ := (&models.DiagramProcess{Name: `DiagramProcess`}).Stage(stage)
	__DiagramProcess__00000001_ := (&models.DiagramProcess{Name: `RQS SPL, AO in the loop`}).Stage(stage)
	__DiagramProcess__00000002_ := (&models.DiagramProcess{Name: `RQS SPL, AO out of the loop`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `ICAO 4444 Chapter 11`}).Stage(stage)

	__Participant__00000000_ := (&models.Participant{Name: `Aircraft Operator`}).Stage(stage)
	__Participant__00000001_ := (&models.Participant{Name: `ATS Reporting Office`}).Stage(stage)
	__Participant__00000002_ := (&models.Participant{Name: `ATS Unit`}).Stage(stage)
	__Participant__00000003_ := (&models.Participant{Name: `ATS Unit`}).Stage(stage)
	__Participant__00000004_ := (&models.Participant{Name: `EUROCONTROL IFPS`}).Stage(stage)
	__Participant__00000005_ := (&models.Participant{Name: `Aircraft Operator`}).Stage(stage)
	__Participant__00000006_ := (&models.Participant{Name: `ATS Unit`}).Stage(stage)
	__Participant__00000007_ := (&models.Participant{Name: `EUROCONTROL IFPS`}).Stage(stage)

	__ParticipantShape__00000000_ := (&models.ParticipantShape{Name: `Aircraft Operator-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000001_ := (&models.ParticipantShape{Name: `ATS Reporting Office-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000002_ := (&models.ParticipantShape{Name: `ATS Unit-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000003_ := (&models.ParticipantShape{Name: `ATS Unit-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000004_ := (&models.ParticipantShape{Name: `EUROCONTROL IFPS-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000005_ := (&models.ParticipantShape{Name: `Aircraft Operator-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000006_ := (&models.ParticipantShape{Name: `ATS Unit-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000007_ := (&models.ParticipantShape{Name: `EUROCONTROL IFPS-DiagramProcess`}).Stage(stage)

	__Process__00000000_ := (&models.Process{Name: `ARO Message Exchanges`}).Stage(stage)
	__Process__00000001_ := (&models.Process{Name: `RQS/SPL Exchange`}).Stage(stage)
	__Process__00000002_ := (&models.Process{Name: `RQS/SPL Exchange (AO out of the loop)`}).Stage(stage)

	__ProcessShape__00000000_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)
	__ProcessShape__00000001_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)
	__ProcessShape__00000002_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)

	__Resource__00000000_ := (&models.Resource{Name: `AO`}).Stage(stage)
	__Resource__00000001_ := (&models.Resource{Name: `ARO`}).Stage(stage)
	__Resource__00000002_ := (&models.Resource{Name: `ATC`}).Stage(stage)
	__Resource__00000003_ := (&models.Resource{Name: `ATC`}).Stage(stage)
	__Resource__00000004_ := (&models.Resource{Name: `IFPS`}).Stage(stage)
	__Resource__00000005_ := (&models.Resource{Name: `AO`}).Stage(stage)
	__Resource__00000006_ := (&models.Resource{Name: `ATC`}).Stage(stage)
	__Resource__00000007_ := (&models.Resource{Name: `IFPS`}).Stage(stage)
	__Resource__00000008_ := (&models.Resource{Name: `AO`}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `Submit Flight Plan`}).Stage(stage)
	__Task__00000001_ := (&models.Task{Name: `Manage Operational Changes`}).Stage(stage)
	__Task__00000002_ := (&models.Task{Name: `Validate and Accept`}).Stage(stage)
	__Task__00000003_ := (&models.Task{Name: `Distribute ATS Messages`}).Stage(stage)
	__Task__00000004_ := (&models.Task{Name: `Ingest Flight Data`}).Stage(stage)
	__Task__00000005_ := (&models.Task{Name: `Log Flight Movement`}).Stage(stage)
	__Task__00000006_ := (&models.Task{Name: ``}).Stage(stage)
	__Task__00000007_ := (&models.Task{Name: ``}).Stage(stage)
	__Task__00000008_ := (&models.Task{Name: `Transmits RQS`}).Stage(stage)
	__Task__00000009_ := (&models.Task{Name: `Validates and Routes RQS`}).Stage(stage)
	__Task__00000010_ := (&models.Task{Name: `Compiles and Transmits SPL`}).Stage(stage)
	__Task__00000011_ := (&models.Task{Name: `Validates and Routes SPL`}).Stage(stage)
	__Task__00000012_ := (&models.Task{Name: `Receives SPL`}).Stage(stage)
	__Task__00000013_ := (&models.Task{Name: ``}).Stage(stage)
	__Task__00000014_ := (&models.Task{Name: ``}).Stage(stage)
	__Task__00000015_ := (&models.Task{Name: `Transmits RQS`}).Stage(stage)
	__Task__00000016_ := (&models.Task{Name: `Validates and Transmit SPL`}).Stage(stage)
	__Task__00000017_ := (&models.Task{Name: `Receives SPL`}).Stage(stage)
	__Task__00000018_ := (&models.Task{Name: ``}).Stage(stage)
	__Task__00000019_ := (&models.Task{Name: ``}).Stage(stage)

	__TaskShape__00000000_ := (&models.TaskShape{Name: `Submit Flight Plan-DiagramProcess`}).Stage(stage)
	__TaskShape__00000001_ := (&models.TaskShape{Name: `Manage Operational Changes-DiagramProcess`}).Stage(stage)
	__TaskShape__00000002_ := (&models.TaskShape{Name: `Validate and Accept-DiagramProcess`}).Stage(stage)
	__TaskShape__00000003_ := (&models.TaskShape{Name: `Distribute ATS Messages-DiagramProcess`}).Stage(stage)
	__TaskShape__00000004_ := (&models.TaskShape{Name: `Ingest Flight Data-DiagramProcess`}).Stage(stage)
	__TaskShape__00000005_ := (&models.TaskShape{Name: `Log Flight Movement-DiagramProcess`}).Stage(stage)
	__TaskShape__00000006_ := (&models.TaskShape{Name: `-DiagramProcess`}).Stage(stage)
	__TaskShape__00000007_ := (&models.TaskShape{Name: `-DiagramProcess`}).Stage(stage)
	__TaskShape__00000008_ := (&models.TaskShape{Name: `Transmits RQS-DiagramProcess`}).Stage(stage)
	__TaskShape__00000009_ := (&models.TaskShape{Name: `Validates and Routes RQS-DiagramProcess`}).Stage(stage)
	__TaskShape__00000010_ := (&models.TaskShape{Name: `Compiles and Transmits SPL-DiagramProcess`}).Stage(stage)
	__TaskShape__00000011_ := (&models.TaskShape{Name: `Validates and Routes SPL-DiagramProcess`}).Stage(stage)
	__TaskShape__00000012_ := (&models.TaskShape{Name: `Receives SPL-DiagramProcess`}).Stage(stage)
	__TaskShape__00000013_ := (&models.TaskShape{Name: `-DiagramProcess-Start`}).Stage(stage)
	__TaskShape__00000014_ := (&models.TaskShape{Name: `-DiagramProcess-End`}).Stage(stage)
	__TaskShape__00000015_ := (&models.TaskShape{Name: `Transmits RQS-DiagramProcess`}).Stage(stage)
	__TaskShape__00000016_ := (&models.TaskShape{Name: `Validates and Routes RQS-DiagramProcess`}).Stage(stage)
	__TaskShape__00000017_ := (&models.TaskShape{Name: `Receives SPL-DiagramProcess`}).Stage(stage)
	__TaskShape__00000018_ := (&models.TaskShape{Name: `-DiagramProcess-Start`}).Stage(stage)
	__TaskShape__00000019_ := (&models.TaskShape{Name: `-DiagramProcess-End`}).Stage(stage)

	// insertion point for initialization of values

	__AllocatedResourceShape__00000000_.Name = `DiagramProcess-Aircraft Operator-AO`

	__AllocatedResourceShape__00000001_.Name = `DiagramProcess-ATS Reporting Office-ARO`

	__AllocatedResourceShape__00000002_.Name = `DiagramProcess-ATS Unit-ATC`

	__AllocatedResourceShape__00000003_.Name = `DiagramProcess-ATS Unit-ATC`

	__AllocatedResourceShape__00000004_.Name = `DiagramProcess-EUROCONTROL IFPS-IFPS`

	__AllocatedResourceShape__00000005_.Name = `DiagramProcess-Aircraft Operator-AO`

	__AllocatedResourceShape__00000006_.Name = `DiagramProcess-ATS Unit-ATC`

	__AllocatedResourceShape__00000007_.Name = `DiagramProcess-EUROCONTROL IFPS-IFPS`

	__AllocatedResourceShape__00000008_.Name = `DiagramProcess-Aircraft Operator-AO`

	__ControlFlow__00000000_.Name = `"Submit Flight Plan" to "Manage Operational Changes"`
	__ControlFlow__00000000_.Description = ``
	__ControlFlow__00000000_.ComputedPrefix = ``

	__ControlFlow__00000001_.Name = `"Validate and Accept" to "Distribute ATS Messages"`
	__ControlFlow__00000001_.Description = ``
	__ControlFlow__00000001_.ComputedPrefix = ``

	__ControlFlow__00000002_.Name = `"Ingest Flight Data" to "Log Flight Movement"`
	__ControlFlow__00000002_.Description = ``
	__ControlFlow__00000002_.ComputedPrefix = ``

	__ControlFlow__00000003_.Name = `"" to "Submit Flight Plan"`
	__ControlFlow__00000003_.Description = ``
	__ControlFlow__00000003_.ComputedPrefix = ``

	__ControlFlow__00000004_.Name = `"Log Flight Movement" to ""`
	__ControlFlow__00000004_.Description = ``
	__ControlFlow__00000004_.ComputedPrefix = ``

	__ControlFlow__00000005_.Name = `"" to "Transmits RQS"`
	__ControlFlow__00000005_.Description = ``
	__ControlFlow__00000005_.ComputedPrefix = ``

	__ControlFlow__00000006_.Name = `"Receives SPL" to ""`
	__ControlFlow__00000006_.Description = ``
	__ControlFlow__00000006_.ComputedPrefix = ``

	__ControlFlow__00000007_.Name = `"" to "Transmits RQS"`
	__ControlFlow__00000007_.Description = ``
	__ControlFlow__00000007_.ComputedPrefix = ``

	__ControlFlow__00000008_.Name = `"Receives SPL" to ""`
	__ControlFlow__00000008_.Description = ``
	__ControlFlow__00000008_.ComputedPrefix = ``

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

	__ControlFlowShape__00000004_.Name = `"Log Flight Movement" to ""`
	__ControlFlowShape__00000004_.StartRatio = 0.552734
	__ControlFlowShape__00000004_.EndRatio = 0.848395
	__ControlFlowShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000004_.CornerOffsetRatio = 1.500000
	__ControlFlowShape__00000004_.IsHidden = false

	__ControlFlowShape__00000005_.Name = `"Receives SPL" to ""`
	__ControlFlowShape__00000005_.StartRatio = 0.500000
	__ControlFlowShape__00000005_.EndRatio = 0.500000
	__ControlFlowShape__00000005_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000005_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000005_.CornerOffsetRatio = 1.800098
	__ControlFlowShape__00000005_.IsHidden = false

	__ControlFlowShape__00000006_.Name = `"" to "Transmits RQS"`
	__ControlFlowShape__00000006_.StartRatio = 1.000000
	__ControlFlowShape__00000006_.EndRatio = 0.500000
	__ControlFlowShape__00000006_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000006_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000006_.CornerOffsetRatio = 2.900293
	__ControlFlowShape__00000006_.IsHidden = false

	__ControlFlowShape__00000007_.Name = `"Receives SPL" to ""`
	__ControlFlowShape__00000007_.StartRatio = 0.500000
	__ControlFlowShape__00000007_.EndRatio = 0.500000
	__ControlFlowShape__00000007_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000007_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000007_.CornerOffsetRatio = 1.800098
	__ControlFlowShape__00000007_.IsHidden = false

	__ControlFlowShape__00000008_.Name = `"" to "Transmits RQS"`
	__ControlFlowShape__00000008_.StartRatio = 1.000000
	__ControlFlowShape__00000008_.EndRatio = 0.500000
	__ControlFlowShape__00000008_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000008_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000008_.CornerOffsetRatio = 2.900293
	__ControlFlowShape__00000008_.IsHidden = false

	__Data__00000000_.Name = `FPL (incl. item 19)`
	__Data__00000000_.Acronym = ``
	__Data__00000000_.Description = `Filed Flight Plan Message`
	__Data__00000000_.ComputedPrefix = ``
	__Data__00000000_.SVG_Path = ``
	__Data__00000000_.InverseAppliedScaling = 0.000000

	__Data__00000001_.Name = `CHG`
	__Data__00000001_.Acronym = ``
	__Data__00000001_.Description = `Modification Message`
	__Data__00000001_.ComputedPrefix = ``
	__Data__00000001_.SVG_Path = ``
	__Data__00000001_.InverseAppliedScaling = 0.000000

	__Data__00000002_.Name = `CNL`
	__Data__00000002_.Acronym = ``
	__Data__00000002_.Description = `Flight Plan Cancellation Message`
	__Data__00000002_.ComputedPrefix = ``
	__Data__00000002_.SVG_Path = ``
	__Data__00000002_.InverseAppliedScaling = 0.000000

	__Data__00000003_.Name = `DLA`
	__Data__00000003_.Acronym = ``
	__Data__00000003_.Description = `Delay Message`
	__Data__00000003_.ComputedPrefix = ``
	__Data__00000003_.SVG_Path = ``
	__Data__00000003_.InverseAppliedScaling = 0.000000

	__Data__00000004_.Name = `DEP`
	__Data__00000004_.Acronym = ``
	__Data__00000004_.Description = `Departure Message`
	__Data__00000004_.ComputedPrefix = ``
	__Data__00000004_.SVG_Path = ``
	__Data__00000004_.InverseAppliedScaling = 0.000000

	__Data__00000005_.Name = `ARR`
	__Data__00000005_.Acronym = ``
	__Data__00000005_.Description = `Arrival Message`
	__Data__00000005_.ComputedPrefix = ``
	__Data__00000005_.SVG_Path = ``
	__Data__00000005_.InverseAppliedScaling = 0.000000

	__Data__00000006_.Name = `FPL (w/o item 19)`
	__Data__00000006_.Acronym = ``
	__Data__00000006_.Description = ``
	__Data__00000006_.ComputedPrefix = ``
	__Data__00000006_.SVG_Path = ``
	__Data__00000006_.InverseAppliedScaling = 0.000000

	__Data__00000007_.Name = `RQS`
	__Data__00000007_.Acronym = ``
	__Data__00000007_.Description = ``
	__Data__00000007_.ComputedPrefix = ``
	__Data__00000007_.SVG_Path = ``
	__Data__00000007_.InverseAppliedScaling = 0.000000

	__Data__00000008_.Name = `SPL`
	__Data__00000008_.Acronym = ``
	__Data__00000008_.Description = ``
	__Data__00000008_.ComputedPrefix = ``
	__Data__00000008_.SVG_Path = ``
	__Data__00000008_.InverseAppliedScaling = 0.000000

	__Data__00000009_.Name = `RQS`
	__Data__00000009_.Acronym = ``
	__Data__00000009_.Description = `Request Supplementary Flight Plan Message`
	__Data__00000009_.ComputedPrefix = ``
	__Data__00000009_.SVG_Path = ``
	__Data__00000009_.InverseAppliedScaling = 0.000000

	__Data__00000010_.Name = `SPL`
	__Data__00000010_.Acronym = ``
	__Data__00000010_.Description = `Supplementary Flight Plan Message (Item 19 data)`
	__Data__00000010_.ComputedPrefix = ``
	__Data__00000010_.SVG_Path = ``
	__Data__00000010_.InverseAppliedScaling = 0.000000

	__Data__00000011_.Name = `RQS`
	__Data__00000011_.Acronym = ``
	__Data__00000011_.Description = `Request Supplementary Flight Plan Message`
	__Data__00000011_.ComputedPrefix = ``
	__Data__00000011_.SVG_Path = ``
	__Data__00000011_.InverseAppliedScaling = 0.000000

	__Data__00000012_.Name = `SPL`
	__Data__00000012_.Acronym = ``
	__Data__00000012_.Description = `Supplementary Flight Plan Message (Item 19 data)`
	__Data__00000012_.ComputedPrefix = ``
	__Data__00000012_.SVG_Path = ``
	__Data__00000012_.InverseAppliedScaling = 0.000000

	__DataFlow__00000000_.Name = `"Submit Flight Plan" to "Validate and Accept"`
	__DataFlow__00000000_.Description = ``
	__DataFlow__00000000_.ComputedPrefix = ``
	__DataFlow__00000000_.Type = models.DataFlow_Task2Task
	__DataFlow__00000000_.IsDatasNodeExpanded = false

	__DataFlow__00000001_.Name = `"Manage Operational Changes" to "Validate and Accept"`
	__DataFlow__00000001_.Description = ``
	__DataFlow__00000001_.ComputedPrefix = ``
	__DataFlow__00000001_.Type = models.DataFlow_Task2Task
	__DataFlow__00000001_.IsDatasNodeExpanded = false

	__DataFlow__00000002_.Name = `"Distribute ATS Messages" to "Ingest Flight Data"`
	__DataFlow__00000002_.Description = ``
	__DataFlow__00000002_.ComputedPrefix = ``
	__DataFlow__00000002_.Type = models.DataFlow_Task2Task
	__DataFlow__00000002_.IsDatasNodeExpanded = false

	__DataFlow__00000003_.Name = `"Log Flight Movement" to "Validate and Accept"`
	__DataFlow__00000003_.Description = ``
	__DataFlow__00000003_.ComputedPrefix = ``
	__DataFlow__00000003_.Type = models.DataFlow_Task2Task
	__DataFlow__00000003_.IsDatasNodeExpanded = false

	__DataFlow__00000004_.Name = `"Log Flight Movement" to "Validate and Accept"`
	__DataFlow__00000004_.Description = ``
	__DataFlow__00000004_.ComputedPrefix = ``
	__DataFlow__00000004_.Type = models.DataFlow_Task2Task
	__DataFlow__00000004_.IsDatasNodeExpanded = false

	__DataFlow__00000005_.Name = `"Transmits RQS" to "Validates and Routes RQS"`
	__DataFlow__00000005_.Description = ``
	__DataFlow__00000005_.ComputedPrefix = ``
	__DataFlow__00000005_.Type = models.DataFlow_Task2Task
	__DataFlow__00000005_.IsDatasNodeExpanded = true

	__DataFlow__00000006_.Name = `"Validates and Routes RQS" to "Compiles and Transmits SPL"`
	__DataFlow__00000006_.Description = ``
	__DataFlow__00000006_.ComputedPrefix = ``
	__DataFlow__00000006_.Type = models.DataFlow_Task2Task
	__DataFlow__00000006_.IsDatasNodeExpanded = false

	__DataFlow__00000007_.Name = `"Compiles and Transmits SPL" to "Validates and Routes SPL"`
	__DataFlow__00000007_.Description = ``
	__DataFlow__00000007_.ComputedPrefix = ``
	__DataFlow__00000007_.Type = models.DataFlow_Task2Task
	__DataFlow__00000007_.IsDatasNodeExpanded = false

	__DataFlow__00000008_.Name = `"Validates and Routes SPL" to "Receives SPL"`
	__DataFlow__00000008_.Description = ``
	__DataFlow__00000008_.ComputedPrefix = ``
	__DataFlow__00000008_.Type = models.DataFlow_Task2Task
	__DataFlow__00000008_.IsDatasNodeExpanded = false

	__DataFlow__00000009_.Name = `"Transmits RQS" to "Validates and Transmit SPL"`
	__DataFlow__00000009_.Description = ``
	__DataFlow__00000009_.ComputedPrefix = ``
	__DataFlow__00000009_.Type = models.DataFlow_Task2Task
	__DataFlow__00000009_.IsDatasNodeExpanded = true

	__DataFlow__00000010_.Name = `"Validates and Transmit SPL" to "Receives SPL"`
	__DataFlow__00000010_.Description = ``
	__DataFlow__00000010_.ComputedPrefix = ``
	__DataFlow__00000010_.Type = models.DataFlow_Task2Task
	__DataFlow__00000010_.IsDatasNodeExpanded = false

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

	__DataFlowShape__00000005_.Name = `"Transmits RQS" to "Validates and Routes RQS"`
	__DataFlowShape__00000005_.StartRatio = 0.500000
	__DataFlowShape__00000005_.EndRatio = 0.500000
	__DataFlowShape__00000005_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000005_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000005_.CornerOffsetRatio = 1.204805
	__DataFlowShape__00000005_.IsHidden = false

	__DataFlowShape__00000006_.Name = `"Validates and Routes SPL" to "Receives SPL"`
	__DataFlowShape__00000006_.StartRatio = 0.500000
	__DataFlowShape__00000006_.EndRatio = 0.483431
	__DataFlowShape__00000006_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000006_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000006_.CornerOffsetRatio = -0.420195
	__DataFlowShape__00000006_.IsHidden = false

	__DataFlowShape__00000007_.Name = `"Validates and Routes RQS" to "Compiles and Transmits SPL"`
	__DataFlowShape__00000007_.StartRatio = 0.500000
	__DataFlowShape__00000007_.EndRatio = 0.262453
	__DataFlowShape__00000007_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000007_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000007_.CornerOffsetRatio = 1.100000
	__DataFlowShape__00000007_.IsHidden = false

	__DataFlowShape__00000008_.Name = `"Compiles and Transmits SPL" to "Validates and Routes SPL"`
	__DataFlowShape__00000008_.StartRatio = 0.907843
	__DataFlowShape__00000008_.EndRatio = 0.500000
	__DataFlowShape__00000008_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000008_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000008_.CornerOffsetRatio = -0.585195
	__DataFlowShape__00000008_.IsHidden = false

	__DataFlowShape__00000009_.Name = `"Transmits RQS" to "Validates and Transmit SPL"`
	__DataFlowShape__00000009_.StartRatio = 0.500000
	__DataFlowShape__00000009_.EndRatio = 0.221653
	__DataFlowShape__00000009_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000009_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000009_.CornerOffsetRatio = 1.534805
	__DataFlowShape__00000009_.IsHidden = false

	__DataFlowShape__00000010_.Name = `"Validates and Transmit SPL" to "Receives SPL"`
	__DataFlowShape__00000010_.StartRatio = 0.713545
	__DataFlowShape__00000010_.EndRatio = 0.483431
	__DataFlowShape__00000010_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000010_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000010_.CornerOffsetRatio = -1.000195
	__DataFlowShape__00000010_.IsHidden = false

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

	__DataShape__00000011_.Name = `"Validates and Routes RQS" to "Compiles and Transmits SPL" - RQS`

	__DataShape__00000012_.Name = `"Compiles and Transmits SPL" to "Validates and Routes SPL" - SPL`

	__DataShape__00000013_.Name = `"Validates and Routes SPL" to "Receives SPL" - SPL`

	__DataShape__00000014_.Name = `"Transmits RQS" to "Validates and Routes RQS"-"Transmits RQS" to "Validates and Routes RQS"-DiagramProcess`

	__DataShape__00000015_.Name = `"Validates and Routes RQS" to "Compiles and Transmits SPL" - RQS`

	__DataShape__00000016_.Name = `"Compiles and Transmits SPL" to "Validates and Routes SPL" - SPL`

	__DataShape__00000017_.Name = `"Validates and Routes SPL" to "Receives SPL" - SPL`

	__DataShape__00000018_.Name = `"Transmits RQS" to "Validates and Routes RQS"-"Transmits RQS" to "Validates and Routes RQS"-DiagramProcess`

	__DiagramProcess__00000000_.Name = `DiagramProcess`
	__DiagramProcess__00000000_.Description = ``
	__DiagramProcess__00000000_.ComputedPrefix = ``
	__DiagramProcess__00000000_.IsChecked = false
	__DiagramProcess__00000000_.IsEditable_ = true
	__DiagramProcess__00000000_.IsShowPrefix = false
	__DiagramProcess__00000000_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000000_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000000_.Width = 1373.000000
	__DiagramProcess__00000000_.Height = 1442.000000
	__DiagramProcess__00000000_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000000_.IsParticipantsNodeExpanded = true
	__DiagramProcess__00000000_.IsExternalParticipantsNodeExpanded = false
	__DiagramProcess__00000000_.IsNotesNodeExpanded = false

	__DiagramProcess__00000001_.Name = `RQS SPL, AO in the loop`
	__DiagramProcess__00000001_.Description = ``
	__DiagramProcess__00000001_.ComputedPrefix = ``
	__DiagramProcess__00000001_.IsChecked = false
	__DiagramProcess__00000001_.IsEditable_ = true
	__DiagramProcess__00000001_.IsShowPrefix = false
	__DiagramProcess__00000001_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000001_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000001_.Width = 1387.000000
	__DiagramProcess__00000001_.Height = 1254.000000
	__DiagramProcess__00000001_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000001_.IsParticipantsNodeExpanded = true
	__DiagramProcess__00000001_.IsExternalParticipantsNodeExpanded = false
	__DiagramProcess__00000001_.IsNotesNodeExpanded = false

	__DiagramProcess__00000002_.Name = `RQS SPL, AO out of the loop`
	__DiagramProcess__00000002_.Description = ``
	__DiagramProcess__00000002_.ComputedPrefix = ``
	__DiagramProcess__00000002_.IsChecked = true
	__DiagramProcess__00000002_.IsEditable_ = true
	__DiagramProcess__00000002_.IsShowPrefix = false
	__DiagramProcess__00000002_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000002_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000002_.Width = 1225.000000
	__DiagramProcess__00000002_.Height = 1243.000000
	__DiagramProcess__00000002_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000002_.IsParticipantsNodeExpanded = true
	__DiagramProcess__00000002_.IsExternalParticipantsNodeExpanded = false
	__DiagramProcess__00000002_.IsNotesNodeExpanded = false

	__Library__00000000_.Name = `ICAO 4444 Chapter 11`
	__Library__00000000_.Description = `Modeling of message exchanges handled by the ATS Reporting Office (ARO)`
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsSubLibrariesNodeExpanded = false
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.IsProcessesNodeExpanded = true
	__Library__00000000_.IsDataFlowsNodeExpanded = false
	__Library__00000000_.IsDatasNodeExpanded = true
	__Library__00000000_.IsResourcesNodeExpanded = true
	__Library__00000000_.IsNotesNodeExpanded = false
	__Library__00000000_.IsExpandedTmp = true

	__Participant__00000000_.Name = `Aircraft Operator`
	__Participant__00000000_.IsProcessResource = false
	__Participant__00000000_.Description = `Originator of initial flight plan data and operational modifications`
	__Participant__00000000_.IsResourcesNodeExpanded = true
	__Participant__00000000_.IsProcessesNodeExpanded = false
	__Participant__00000000_.ComputedPrefix = ``
	__Participant__00000000_.IsTasksNodeExpanded = true
	__Participant__00000000_.IsControlFlowsNodeExpanded = false
	__Participant__00000000_.IsDataFlowsNodeExpanded = false

	__Participant__00000001_.Name = `ATS Reporting Office`
	__Participant__00000001_.IsProcessResource = false
	__Participant__00000001_.Description = `Receives, validates, accepts, and distributes flight plan-related messages`
	__Participant__00000001_.IsResourcesNodeExpanded = false
	__Participant__00000001_.IsProcessesNodeExpanded = false
	__Participant__00000001_.ComputedPrefix = ``
	__Participant__00000001_.IsTasksNodeExpanded = true
	__Participant__00000001_.IsControlFlowsNodeExpanded = false
	__Participant__00000001_.IsDataFlowsNodeExpanded = false

	__Participant__00000002_.Name = `ATS Unit`
	__Participant__00000002_.IsProcessResource = false
	__Participant__00000002_.Description = `Recipient of distributed movement and control data used for traffic separation`
	__Participant__00000002_.IsResourcesNodeExpanded = false
	__Participant__00000002_.IsProcessesNodeExpanded = false
	__Participant__00000002_.ComputedPrefix = ``
	__Participant__00000002_.IsTasksNodeExpanded = true
	__Participant__00000002_.IsControlFlowsNodeExpanded = false
	__Participant__00000002_.IsDataFlowsNodeExpanded = false

	__Participant__00000003_.Name = `ATS Unit`
	__Participant__00000003_.IsProcessResource = false
	__Participant__00000003_.Description = `Air Traffic Services Unit / Rescue Coordination Center`
	__Participant__00000003_.IsResourcesNodeExpanded = false
	__Participant__00000003_.IsProcessesNodeExpanded = false
	__Participant__00000003_.ComputedPrefix = ``
	__Participant__00000003_.IsTasksNodeExpanded = false
	__Participant__00000003_.IsControlFlowsNodeExpanded = true
	__Participant__00000003_.IsDataFlowsNodeExpanded = false

	__Participant__00000004_.Name = `EUROCONTROL IFPS`
	__Participant__00000004_.IsProcessResource = false
	__Participant__00000004_.Description = `Initial Flight Plan Processing System`
	__Participant__00000004_.IsResourcesNodeExpanded = false
	__Participant__00000004_.IsProcessesNodeExpanded = false
	__Participant__00000004_.ComputedPrefix = ``
	__Participant__00000004_.IsTasksNodeExpanded = true
	__Participant__00000004_.IsControlFlowsNodeExpanded = false
	__Participant__00000004_.IsDataFlowsNodeExpanded = false

	__Participant__00000005_.Name = `Aircraft Operator`
	__Participant__00000005_.IsProcessResource = false
	__Participant__00000005_.Description = `Aircraft Operator or automated dispatch`
	__Participant__00000005_.IsResourcesNodeExpanded = false
	__Participant__00000005_.IsProcessesNodeExpanded = false
	__Participant__00000005_.ComputedPrefix = ``
	__Participant__00000005_.IsTasksNodeExpanded = false
	__Participant__00000005_.IsControlFlowsNodeExpanded = false
	__Participant__00000005_.IsDataFlowsNodeExpanded = false

	__Participant__00000006_.Name = `ATS Unit`
	__Participant__00000006_.IsProcessResource = false
	__Participant__00000006_.Description = `Air Traffic Services Unit / Rescue Coordination Center`
	__Participant__00000006_.IsResourcesNodeExpanded = false
	__Participant__00000006_.IsProcessesNodeExpanded = false
	__Participant__00000006_.ComputedPrefix = ``
	__Participant__00000006_.IsTasksNodeExpanded = false
	__Participant__00000006_.IsControlFlowsNodeExpanded = true
	__Participant__00000006_.IsDataFlowsNodeExpanded = false

	__Participant__00000007_.Name = `EUROCONTROL IFPS`
	__Participant__00000007_.IsProcessResource = false
	__Participant__00000007_.Description = `Initial Flight Plan Processing System`
	__Participant__00000007_.IsResourcesNodeExpanded = false
	__Participant__00000007_.IsProcessesNodeExpanded = false
	__Participant__00000007_.ComputedPrefix = ``
	__Participant__00000007_.IsTasksNodeExpanded = true
	__Participant__00000007_.IsControlFlowsNodeExpanded = false
	__Participant__00000007_.IsDataFlowsNodeExpanded = false

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

	__ParticipantShape__00000003_.Name = `ATS Unit-DiagramProcess`
	__ParticipantShape__00000003_.IsExpanded = false
	__ParticipantShape__00000003_.X = 100.000000
	__ParticipantShape__00000003_.Y = 100.000000
	__ParticipantShape__00000003_.Width = 250.000000
	__ParticipantShape__00000003_.Height = 70.000000
	__ParticipantShape__00000003_.IsHidden = false
	__ParticipantShape__00000003_.WidthWeight = 1.000000

	__ParticipantShape__00000004_.Name = `EUROCONTROL IFPS-DiagramProcess`
	__ParticipantShape__00000004_.IsExpanded = false
	__ParticipantShape__00000004_.X = 450.000000
	__ParticipantShape__00000004_.Y = 100.000000
	__ParticipantShape__00000004_.Width = 250.000000
	__ParticipantShape__00000004_.Height = 70.000000
	__ParticipantShape__00000004_.IsHidden = false
	__ParticipantShape__00000004_.WidthWeight = 1.000000

	__ParticipantShape__00000005_.Name = `Aircraft Operator-DiagramProcess`
	__ParticipantShape__00000005_.IsExpanded = false
	__ParticipantShape__00000005_.X = 800.000000
	__ParticipantShape__00000005_.Y = 100.000000
	__ParticipantShape__00000005_.Width = 250.000000
	__ParticipantShape__00000005_.Height = 70.000000
	__ParticipantShape__00000005_.IsHidden = false
	__ParticipantShape__00000005_.WidthWeight = 1.000000

	__ParticipantShape__00000006_.Name = `ATS Unit-DiagramProcess`
	__ParticipantShape__00000006_.IsExpanded = false
	__ParticipantShape__00000006_.X = 100.000000
	__ParticipantShape__00000006_.Y = 100.000000
	__ParticipantShape__00000006_.Width = 250.000000
	__ParticipantShape__00000006_.Height = 70.000000
	__ParticipantShape__00000006_.IsHidden = false
	__ParticipantShape__00000006_.WidthWeight = 1.000000

	__ParticipantShape__00000007_.Name = `EUROCONTROL IFPS-DiagramProcess`
	__ParticipantShape__00000007_.IsExpanded = false
	__ParticipantShape__00000007_.X = 450.000000
	__ParticipantShape__00000007_.Y = 100.000000
	__ParticipantShape__00000007_.Width = 250.000000
	__ParticipantShape__00000007_.Height = 70.000000
	__ParticipantShape__00000007_.IsHidden = false
	__ParticipantShape__00000007_.WidthWeight = 1.000000

	__Process__00000000_.Name = `ARO Message Exchanges`
	__Process__00000000_.Description = `Process governing the submission, validation, modification, and distribution of ATS messages via the ARO.`
	__Process__00000000_.ComputedPrefix = ``
	__Process__00000000_.SVG_Path = ``
	__Process__00000000_.InverseAppliedScaling = 0.000000
	__Process__00000000_.IsSubProcessNodeExpanded = false
	__Process__00000000_.IsDataFlowsNodeExpanded = false

	__Process__00000001_.Name = `RQS/SPL Exchange`
	__Process__00000001_.Description = `Process for requesting and delivering supplementary flight plan data (Item 19).`
	__Process__00000001_.ComputedPrefix = ``
	__Process__00000001_.SVG_Path = ``
	__Process__00000001_.InverseAppliedScaling = 0.000000
	__Process__00000001_.IsSubProcessNodeExpanded = false
	__Process__00000001_.IsDataFlowsNodeExpanded = false

	__Process__00000002_.Name = `RQS/SPL Exchange (AO out of the loop)`
	__Process__00000002_.Description = `Process for requesting and delivering supplementary flight plan data (Item 19).`
	__Process__00000002_.ComputedPrefix = ``
	__Process__00000002_.SVG_Path = ``
	__Process__00000002_.InverseAppliedScaling = 0.000000
	__Process__00000002_.IsSubProcessNodeExpanded = false
	__Process__00000002_.IsDataFlowsNodeExpanded = false

	__ProcessShape__00000000_.Name = `ProcessShape`
	__ProcessShape__00000000_.IsExpanded = false
	__ProcessShape__00000000_.X = 53.000000
	__ProcessShape__00000000_.Y = 16.000000
	__ProcessShape__00000000_.Width = 1020.000000
	__ProcessShape__00000000_.Height = 1126.000000
	__ProcessShape__00000000_.IsHidden = false

	__ProcessShape__00000001_.Name = `ProcessShape`
	__ProcessShape__00000001_.IsExpanded = false
	__ProcessShape__00000001_.X = 37.000000
	__ProcessShape__00000001_.Y = 54.000000
	__ProcessShape__00000001_.Width = 1050.000000
	__ProcessShape__00000001_.Height = 900.000000
	__ProcessShape__00000001_.IsHidden = false

	__ProcessShape__00000002_.Name = `ProcessShape`
	__ProcessShape__00000002_.IsExpanded = false
	__ProcessShape__00000002_.X = 47.000000
	__ProcessShape__00000002_.Y = 43.000000
	__ProcessShape__00000002_.Width = 878.000000
	__ProcessShape__00000002_.Height = 900.000000
	__ProcessShape__00000002_.IsHidden = false

	__Resource__00000000_.Name = `AO`
	__Resource__00000000_.Acronym = ``
	__Resource__00000000_.Description = `Aircraft Operator / Flight Crew`
	__Resource__00000000_.ComputedPrefix = ``
	__Resource__00000000_.SVG_Path = ``
	__Resource__00000000_.InverseAppliedScaling = 0.000000

	__Resource__00000001_.Name = `ARO`
	__Resource__00000001_.Acronym = ``
	__Resource__00000001_.Description = `ATS Reporting Office`
	__Resource__00000001_.ComputedPrefix = ``
	__Resource__00000001_.SVG_Path = ``
	__Resource__00000001_.InverseAppliedScaling = 0.000000

	__Resource__00000002_.Name = `ATC`
	__Resource__00000002_.Acronym = ``
	__Resource__00000002_.Description = `Air Traffic Control Unit / ACC / TWR`
	__Resource__00000002_.ComputedPrefix = ``
	__Resource__00000002_.SVG_Path = ``
	__Resource__00000002_.InverseAppliedScaling = 0.000000

	__Resource__00000003_.Name = `ATC`
	__Resource__00000003_.Acronym = ``
	__Resource__00000003_.Description = ``
	__Resource__00000003_.ComputedPrefix = ``
	__Resource__00000003_.SVG_Path = ``
	__Resource__00000003_.InverseAppliedScaling = 0.000000

	__Resource__00000004_.Name = `IFPS`
	__Resource__00000004_.Acronym = ``
	__Resource__00000004_.Description = ``
	__Resource__00000004_.ComputedPrefix = ``
	__Resource__00000004_.SVG_Path = ``
	__Resource__00000004_.InverseAppliedScaling = 0.000000

	__Resource__00000005_.Name = `AO`
	__Resource__00000005_.Acronym = ``
	__Resource__00000005_.Description = ``
	__Resource__00000005_.ComputedPrefix = ``
	__Resource__00000005_.SVG_Path = ``
	__Resource__00000005_.InverseAppliedScaling = 0.000000

	__Resource__00000006_.Name = `ATC`
	__Resource__00000006_.Acronym = ``
	__Resource__00000006_.Description = ``
	__Resource__00000006_.ComputedPrefix = ``
	__Resource__00000006_.SVG_Path = ``
	__Resource__00000006_.InverseAppliedScaling = 0.000000

	__Resource__00000007_.Name = `IFPS`
	__Resource__00000007_.Acronym = ``
	__Resource__00000007_.Description = ``
	__Resource__00000007_.ComputedPrefix = ``
	__Resource__00000007_.SVG_Path = ``
	__Resource__00000007_.InverseAppliedScaling = 0.000000

	__Resource__00000008_.Name = `AO`
	__Resource__00000008_.Acronym = ``
	__Resource__00000008_.Description = ``
	__Resource__00000008_.ComputedPrefix = ``
	__Resource__00000008_.SVG_Path = ``
	__Resource__00000008_.InverseAppliedScaling = 0.000000

	__Task__00000000_.Name = `Submit Flight Plan`
	__Task__00000000_.Description = `Originates and files the initial field flight plan parameters before departure.`
	__Task__00000000_.ComputedPrefix = ``
	__Task__00000000_.IsStartTask = false
	__Task__00000000_.IsEndTask = false
	__Task__00000000_.IsTaskNameNotProcessName = true

	__Task__00000001_.Name = `Manage Operational Changes`
	__Task__00000001_.Description = `Monitors operational status and decides if a flight needs to be delayed, modified, or aborted prior to departure.`
	__Task__00000001_.ComputedPrefix = ``
	__Task__00000001_.IsStartTask = false
	__Task__00000001_.IsEndTask = false
	__Task__00000001_.IsTaskNameNotProcessName = true

	__Task__00000002_.Name = `Validate and Accept`
	__Task__00000002_.Description = `Receives incoming submissions and checks them for syntax, formatting correctness, and logical consistency according to ICAO standards.`
	__Task__00000002_.ComputedPrefix = ``
	__Task__00000002_.IsStartTask = false
	__Task__00000002_.IsEndTask = false
	__Task__00000002_.IsTaskNameNotProcessName = true

	__Task__00000003_.Name = `Distribute ATS Messages`
	__Task__00000003_.Description = `Determines the appropriate addressing routing and distributes accepted flight plans or updates to downstream units.`
	__Task__00000003_.ComputedPrefix = ``
	__Task__00000003_.IsStartTask = false
	__Task__00000003_.IsEndTask = false
	__Task__00000003_.IsTaskNameNotProcessName = true

	__Task__00000004_.Name = `Ingest Flight Data`
	__Task__00000004_.Description = `Receives distributed data and loads it into the local Flight Data Processing System (FDPS) for air traffic controllers.`
	__Task__00000004_.ComputedPrefix = ``
	__Task__00000004_.IsStartTask = false
	__Task__00000004_.IsEndTask = false
	__Task__00000004_.IsTaskNameNotProcessName = true

	__Task__00000005_.Name = `Log Flight Movement`
	__Task__00000005_.Description = `Tracks physical execution milestones (actual departure and arrival times) and generates status updates.`
	__Task__00000005_.ComputedPrefix = ``
	__Task__00000005_.IsStartTask = false
	__Task__00000005_.IsEndTask = false
	__Task__00000005_.IsTaskNameNotProcessName = true

	__Task__00000006_.Name = ``
	__Task__00000006_.Description = ``
	__Task__00000006_.ComputedPrefix = ``
	__Task__00000006_.IsStartTask = true
	__Task__00000006_.IsEndTask = false
	__Task__00000006_.IsTaskNameNotProcessName = false

	__Task__00000007_.Name = ``
	__Task__00000007_.Description = ``
	__Task__00000007_.ComputedPrefix = ``
	__Task__00000007_.IsStartTask = false
	__Task__00000007_.IsEndTask = true
	__Task__00000007_.IsTaskNameNotProcessName = false

	__Task__00000008_.Name = `Transmits RQS`
	__Task__00000008_.Description = ``
	__Task__00000008_.ComputedPrefix = ``
	__Task__00000008_.IsStartTask = false
	__Task__00000008_.IsEndTask = false
	__Task__00000008_.IsTaskNameNotProcessName = false

	__Task__00000009_.Name = `Validates and Routes RQS`
	__Task__00000009_.Description = ``
	__Task__00000009_.ComputedPrefix = ``
	__Task__00000009_.IsStartTask = false
	__Task__00000009_.IsEndTask = false
	__Task__00000009_.IsTaskNameNotProcessName = false

	__Task__00000010_.Name = `Compiles and Transmits SPL`
	__Task__00000010_.Description = ``
	__Task__00000010_.ComputedPrefix = ``
	__Task__00000010_.IsStartTask = false
	__Task__00000010_.IsEndTask = false
	__Task__00000010_.IsTaskNameNotProcessName = false

	__Task__00000011_.Name = `Validates and Routes SPL`
	__Task__00000011_.Description = ``
	__Task__00000011_.ComputedPrefix = ``
	__Task__00000011_.IsStartTask = false
	__Task__00000011_.IsEndTask = false
	__Task__00000011_.IsTaskNameNotProcessName = false

	__Task__00000012_.Name = `Receives SPL`
	__Task__00000012_.Description = ``
	__Task__00000012_.ComputedPrefix = ``
	__Task__00000012_.IsStartTask = false
	__Task__00000012_.IsEndTask = false
	__Task__00000012_.IsTaskNameNotProcessName = false

	__Task__00000013_.Name = ``
	__Task__00000013_.Description = ``
	__Task__00000013_.ComputedPrefix = ``
	__Task__00000013_.IsStartTask = true
	__Task__00000013_.IsEndTask = false
	__Task__00000013_.IsTaskNameNotProcessName = false

	__Task__00000014_.Name = ``
	__Task__00000014_.Description = ``
	__Task__00000014_.ComputedPrefix = ``
	__Task__00000014_.IsStartTask = false
	__Task__00000014_.IsEndTask = true
	__Task__00000014_.IsTaskNameNotProcessName = false

	__Task__00000015_.Name = `Transmits RQS`
	__Task__00000015_.Description = ``
	__Task__00000015_.ComputedPrefix = ``
	__Task__00000015_.IsStartTask = false
	__Task__00000015_.IsEndTask = false
	__Task__00000015_.IsTaskNameNotProcessName = false

	__Task__00000016_.Name = `Validates and Transmit SPL`
	__Task__00000016_.Description = ``
	__Task__00000016_.ComputedPrefix = ``
	__Task__00000016_.IsStartTask = false
	__Task__00000016_.IsEndTask = false
	__Task__00000016_.IsTaskNameNotProcessName = false

	__Task__00000017_.Name = `Receives SPL`
	__Task__00000017_.Description = ``
	__Task__00000017_.ComputedPrefix = ``
	__Task__00000017_.IsStartTask = false
	__Task__00000017_.IsEndTask = false
	__Task__00000017_.IsTaskNameNotProcessName = false

	__Task__00000018_.Name = ``
	__Task__00000018_.Description = ``
	__Task__00000018_.ComputedPrefix = ``
	__Task__00000018_.IsStartTask = true
	__Task__00000018_.IsEndTask = false
	__Task__00000018_.IsTaskNameNotProcessName = false

	__Task__00000019_.Name = ``
	__Task__00000019_.Description = ``
	__Task__00000019_.ComputedPrefix = ``
	__Task__00000019_.IsStartTask = false
	__Task__00000019_.IsEndTask = true
	__Task__00000019_.IsTaskNameNotProcessName = false

	__TaskShape__00000000_.Name = `Submit Flight Plan-DiagramProcess`
	__TaskShape__00000000_.IsExpanded = false
	__TaskShape__00000000_.X = 86.000000
	__TaskShape__00000000_.Y = 277.795175
	__TaskShape__00000000_.Width = 240.333333
	__TaskShape__00000000_.Height = 48.000000
	__TaskShape__00000000_.IsHidden = false

	__TaskShape__00000001_.Name = `Manage Operational Changes-DiagramProcess`
	__TaskShape__00000001_.IsExpanded = false
	__TaskShape__00000001_.X = 87.000000
	__TaskShape__00000001_.Y = 377.907212
	__TaskShape__00000001_.Width = 240.333333
	__TaskShape__00000001_.Height = 70.000000
	__TaskShape__00000001_.IsHidden = false

	__TaskShape__00000002_.Name = `Validate and Accept-DiagramProcess`
	__TaskShape__00000002_.IsExpanded = false
	__TaskShape__00000002_.X = 461.000000
	__TaskShape__00000002_.Y = 256.100648
	__TaskShape__00000002_.Width = 189.333333
	__TaskShape__00000002_.Height = 259.000000
	__TaskShape__00000002_.IsHidden = false

	__TaskShape__00000003_.Name = `Distribute ATS Messages-DiagramProcess`
	__TaskShape__00000003_.IsExpanded = false
	__TaskShape__00000003_.X = 461.000000
	__TaskShape__00000003_.Y = 579.893285
	__TaskShape__00000003_.Width = 175.000000
	__TaskShape__00000003_.Height = 70.000000
	__TaskShape__00000003_.IsHidden = false

	__TaskShape__00000004_.Name = `Ingest Flight Data-DiagramProcess`
	__TaskShape__00000004_.IsExpanded = false
	__TaskShape__00000004_.X = 814.666667
	__TaskShape__00000004_.Y = 554.382953
	__TaskShape__00000004_.Width = 160.000000
	__TaskShape__00000004_.Height = 116.000000
	__TaskShape__00000004_.IsHidden = false

	__TaskShape__00000005_.Name = `Log Flight Movement-DiagramProcess`
	__TaskShape__00000005_.IsExpanded = false
	__TaskShape__00000005_.X = 809.000000
	__TaskShape__00000005_.Y = 766.607545
	__TaskShape__00000005_.Width = 156.000000
	__TaskShape__00000005_.Height = 60.000000
	__TaskShape__00000005_.IsHidden = false

	__TaskShape__00000006_.Name = `-DiagramProcess`
	__TaskShape__00000006_.IsExpanded = false
	__TaskShape__00000006_.X = 114.000000
	__TaskShape__00000006_.Y = 210.251563
	__TaskShape__00000006_.Width = 83.000000
	__TaskShape__00000006_.Height = 20.000000
	__TaskShape__00000006_.IsHidden = false

	__TaskShape__00000007_.Name = `-DiagramProcess`
	__TaskShape__00000007_.IsExpanded = false
	__TaskShape__00000007_.X = 933.666667
	__TaskShape__00000007_.Y = 1061.033535
	__TaskShape__00000007_.Width = 97.000000
	__TaskShape__00000007_.Height = 36.000000
	__TaskShape__00000007_.IsHidden = false

	__TaskShape__00000008_.Name = `Transmits RQS-DiagramProcess`
	__TaskShape__00000008_.IsExpanded = false
	__TaskShape__00000008_.X = 112.000000
	__TaskShape__00000008_.Y = 314.000000
	__TaskShape__00000008_.Width = 200.000000
	__TaskShape__00000008_.Height = 60.000000
	__TaskShape__00000008_.IsHidden = false

	__TaskShape__00000009_.Name = `Validates and Routes RQS-DiagramProcess`
	__TaskShape__00000009_.IsExpanded = false
	__TaskShape__00000009_.X = 462.000000
	__TaskShape__00000009_.Y = 414.000000
	__TaskShape__00000009_.Width = 200.000000
	__TaskShape__00000009_.Height = 60.000000
	__TaskShape__00000009_.IsHidden = false

	__TaskShape__00000010_.Name = `Compiles and Transmits SPL-DiagramProcess`
	__TaskShape__00000010_.IsExpanded = false
	__TaskShape__00000010_.X = 814.000000
	__TaskShape__00000010_.Y = 451.000000
	__TaskShape__00000010_.Width = 200.000000
	__TaskShape__00000010_.Height = 141.000000
	__TaskShape__00000010_.IsHidden = false

	__TaskShape__00000011_.Name = `Validates and Routes SPL-DiagramProcess`
	__TaskShape__00000011_.IsExpanded = false
	__TaskShape__00000011_.X = 462.000000
	__TaskShape__00000011_.Y = 614.000000
	__TaskShape__00000011_.Width = 200.000000
	__TaskShape__00000011_.Height = 60.000000
	__TaskShape__00000011_.IsHidden = false

	__TaskShape__00000012_.Name = `Receives SPL-DiagramProcess`
	__TaskShape__00000012_.IsExpanded = false
	__TaskShape__00000012_.X = 113.000000
	__TaskShape__00000012_.Y = 662.000000
	__TaskShape__00000012_.Width = 200.000000
	__TaskShape__00000012_.Height = 60.000000
	__TaskShape__00000012_.IsHidden = false

	__TaskShape__00000013_.Name = `-DiagramProcess-Start`
	__TaskShape__00000013_.IsExpanded = false
	__TaskShape__00000013_.X = 167.000000
	__TaskShape__00000013_.Y = 214.000000
	__TaskShape__00000013_.Width = 50.000000
	__TaskShape__00000013_.Height = 20.000000
	__TaskShape__00000013_.IsHidden = false

	__TaskShape__00000014_.Name = `-DiagramProcess-End`
	__TaskShape__00000014_.IsExpanded = false
	__TaskShape__00000014_.X = 182.000000
	__TaskShape__00000014_.Y = 849.000000
	__TaskShape__00000014_.Width = 50.000000
	__TaskShape__00000014_.Height = 36.000000
	__TaskShape__00000014_.IsHidden = false

	__TaskShape__00000015_.Name = `Transmits RQS-DiagramProcess`
	__TaskShape__00000015_.IsExpanded = false
	__TaskShape__00000015_.X = 119.000000
	__TaskShape__00000015_.Y = 304.000000
	__TaskShape__00000015_.Width = 200.000000
	__TaskShape__00000015_.Height = 60.000000
	__TaskShape__00000015_.IsHidden = false

	__TaskShape__00000016_.Name = `Validates and Routes RQS-DiagramProcess`
	__TaskShape__00000016_.IsExpanded = false
	__TaskShape__00000016_.X = 622.000000
	__TaskShape__00000016_.Y = 446.000000
	__TaskShape__00000016_.Width = 200.000000
	__TaskShape__00000016_.Height = 185.000000
	__TaskShape__00000016_.IsHidden = false

	__TaskShape__00000017_.Name = `Receives SPL-DiagramProcess`
	__TaskShape__00000017_.IsExpanded = false
	__TaskShape__00000017_.X = 120.000000
	__TaskShape__00000017_.Y = 652.000000
	__TaskShape__00000017_.Width = 200.000000
	__TaskShape__00000017_.Height = 60.000000
	__TaskShape__00000017_.IsHidden = false

	__TaskShape__00000018_.Name = `-DiagramProcess-Start`
	__TaskShape__00000018_.IsExpanded = false
	__TaskShape__00000018_.X = 174.000000
	__TaskShape__00000018_.Y = 204.000000
	__TaskShape__00000018_.Width = 50.000000
	__TaskShape__00000018_.Height = 20.000000
	__TaskShape__00000018_.IsHidden = false

	__TaskShape__00000019_.Name = `-DiagramProcess-End`
	__TaskShape__00000019_.IsExpanded = false
	__TaskShape__00000019_.X = 189.000000
	__TaskShape__00000019_.Y = 839.000000
	__TaskShape__00000019_.Width = 50.000000
	__TaskShape__00000019_.Height = 36.000000
	__TaskShape__00000019_.IsHidden = false

	// insertion point for setup of pointers
	__AllocatedResourceShape__00000000_.Participant = __Participant__00000000_
	__AllocatedResourceShape__00000000_.Resource = __Resource__00000000_
	__AllocatedResourceShape__00000001_.Participant = __Participant__00000001_
	__AllocatedResourceShape__00000001_.Resource = __Resource__00000001_
	__AllocatedResourceShape__00000002_.Participant = __Participant__00000002_
	__AllocatedResourceShape__00000002_.Resource = __Resource__00000002_
	__AllocatedResourceShape__00000003_.Participant = __Participant__00000003_
	__AllocatedResourceShape__00000003_.Resource = __Resource__00000003_
	__AllocatedResourceShape__00000004_.Participant = __Participant__00000004_
	__AllocatedResourceShape__00000004_.Resource = __Resource__00000004_
	__AllocatedResourceShape__00000005_.Participant = __Participant__00000005_
	__AllocatedResourceShape__00000005_.Resource = __Resource__00000005_
	__AllocatedResourceShape__00000006_.Participant = __Participant__00000006_
	__AllocatedResourceShape__00000006_.Resource = __Resource__00000006_
	__AllocatedResourceShape__00000007_.Participant = __Participant__00000007_
	__AllocatedResourceShape__00000007_.Resource = __Resource__00000007_
	__AllocatedResourceShape__00000008_.Participant = nil
	__AllocatedResourceShape__00000008_.Resource = __Resource__00000008_
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
	__ControlFlow__00000005_.Start = __Task__00000013_
	__ControlFlow__00000005_.End = __Task__00000008_
	__ControlFlow__00000006_.Start = __Task__00000012_
	__ControlFlow__00000006_.End = __Task__00000014_
	__ControlFlow__00000007_.Start = __Task__00000018_
	__ControlFlow__00000007_.End = __Task__00000015_
	__ControlFlow__00000008_.Start = __Task__00000017_
	__ControlFlow__00000008_.End = __Task__00000019_
	__ControlFlowShape__00000000_.ControlFlow = __ControlFlow__00000000_
	__ControlFlowShape__00000001_.ControlFlow = __ControlFlow__00000001_
	__ControlFlowShape__00000002_.ControlFlow = __ControlFlow__00000002_
	__ControlFlowShape__00000003_.ControlFlow = __ControlFlow__00000003_
	__ControlFlowShape__00000004_.ControlFlow = __ControlFlow__00000004_
	__ControlFlowShape__00000005_.ControlFlow = __ControlFlow__00000006_
	__ControlFlowShape__00000006_.ControlFlow = __ControlFlow__00000005_
	__ControlFlowShape__00000007_.ControlFlow = __ControlFlow__00000008_
	__ControlFlowShape__00000008_.ControlFlow = __ControlFlow__00000007_
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
	__DataFlow__00000005_.Datas = append(__DataFlow__00000005_.Datas, __Data__00000009_)
	__DataFlow__00000005_.StartTask = __Task__00000008_
	__DataFlow__00000005_.EndTask = __Task__00000009_
	__DataFlow__00000005_.StartExternalParticipant = nil
	__DataFlow__00000005_.EndExternalParticipant = nil
	__DataFlow__00000006_.Datas = append(__DataFlow__00000006_.Datas, __Data__00000009_)
	__DataFlow__00000006_.StartTask = __Task__00000009_
	__DataFlow__00000006_.EndTask = __Task__00000010_
	__DataFlow__00000006_.StartExternalParticipant = nil
	__DataFlow__00000006_.EndExternalParticipant = nil
	__DataFlow__00000007_.Datas = append(__DataFlow__00000007_.Datas, __Data__00000010_)
	__DataFlow__00000007_.StartTask = __Task__00000010_
	__DataFlow__00000007_.EndTask = __Task__00000011_
	__DataFlow__00000007_.StartExternalParticipant = nil
	__DataFlow__00000007_.EndExternalParticipant = nil
	__DataFlow__00000008_.Datas = append(__DataFlow__00000008_.Datas, __Data__00000010_)
	__DataFlow__00000008_.StartTask = __Task__00000011_
	__DataFlow__00000008_.EndTask = __Task__00000012_
	__DataFlow__00000008_.StartExternalParticipant = nil
	__DataFlow__00000008_.EndExternalParticipant = nil
	__DataFlow__00000009_.Datas = append(__DataFlow__00000009_.Datas, __Data__00000011_)
	__DataFlow__00000009_.StartTask = __Task__00000015_
	__DataFlow__00000009_.EndTask = __Task__00000016_
	__DataFlow__00000009_.StartExternalParticipant = nil
	__DataFlow__00000009_.EndExternalParticipant = nil
	__DataFlow__00000010_.Datas = append(__DataFlow__00000010_.Datas, __Data__00000012_)
	__DataFlow__00000010_.StartTask = __Task__00000016_
	__DataFlow__00000010_.EndTask = __Task__00000017_
	__DataFlow__00000010_.StartExternalParticipant = nil
	__DataFlow__00000010_.EndExternalParticipant = nil
	__DataFlowShape__00000000_.DataFlow = __DataFlow__00000000_
	__DataFlowShape__00000001_.DataFlow = __DataFlow__00000001_
	__DataFlowShape__00000002_.DataFlow = __DataFlow__00000002_
	__DataFlowShape__00000003_.DataFlow = __DataFlow__00000003_
	__DataFlowShape__00000004_.DataFlow = __DataFlow__00000004_
	__DataFlowShape__00000005_.DataFlow = __DataFlow__00000005_
	__DataFlowShape__00000006_.DataFlow = __DataFlow__00000008_
	__DataFlowShape__00000007_.DataFlow = __DataFlow__00000006_
	__DataFlowShape__00000008_.DataFlow = __DataFlow__00000007_
	__DataFlowShape__00000009_.DataFlow = __DataFlow__00000009_
	__DataFlowShape__00000010_.DataFlow = __DataFlow__00000010_
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
	__DataShape__00000011_.Data = __Data__00000009_
	__DataShape__00000011_.DataFlow = __DataFlow__00000006_
	__DataShape__00000012_.Data = __Data__00000010_
	__DataShape__00000012_.DataFlow = __DataFlow__00000007_
	__DataShape__00000013_.Data = __Data__00000010_
	__DataShape__00000013_.DataFlow = __DataFlow__00000008_
	__DataShape__00000014_.Data = __Data__00000009_
	__DataShape__00000014_.DataFlow = __DataFlow__00000005_
	__DataShape__00000015_.Data = __Data__00000011_
	__DataShape__00000015_.DataFlow = nil
	__DataShape__00000016_.Data = __Data__00000012_
	__DataShape__00000016_.DataFlow = nil
	__DataShape__00000017_.Data = __Data__00000012_
	__DataShape__00000017_.DataFlow = __DataFlow__00000010_
	__DataShape__00000018_.Data = __Data__00000011_
	__DataShape__00000018_.DataFlow = __DataFlow__00000009_
	__DiagramProcess__00000000_.Process_Shapes = append(__DiagramProcess__00000000_.Process_Shapes, __ProcessShape__00000000_)
	__DiagramProcess__00000000_.Participant_Shapes = append(__DiagramProcess__00000000_.Participant_Shapes, __ParticipantShape__00000000_)
	__DiagramProcess__00000000_.Participant_Shapes = append(__DiagramProcess__00000000_.Participant_Shapes, __ParticipantShape__00000001_)
	__DiagramProcess__00000000_.Participant_Shapes = append(__DiagramProcess__00000000_.Participant_Shapes, __ParticipantShape__00000002_)
	__DiagramProcess__00000000_.ParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000000_.ParticipantWhoseNodeIsExpanded, __Participant__00000002_)
	__DiagramProcess__00000000_.TasksWhoseNodeIsExpanded = append(__DiagramProcess__00000000_.TasksWhoseNodeIsExpanded, __Task__00000003_)
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
	__DiagramProcess__00000001_.Process_Shapes = append(__DiagramProcess__00000001_.Process_Shapes, __ProcessShape__00000001_)
	__DiagramProcess__00000001_.Participant_Shapes = append(__DiagramProcess__00000001_.Participant_Shapes, __ParticipantShape__00000003_)
	__DiagramProcess__00000001_.Participant_Shapes = append(__DiagramProcess__00000001_.Participant_Shapes, __ParticipantShape__00000004_)
	__DiagramProcess__00000001_.Participant_Shapes = append(__DiagramProcess__00000001_.Participant_Shapes, __ParticipantShape__00000005_)
	__DiagramProcess__00000001_.ParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000001_.ParticipantWhoseNodeIsExpanded, __Participant__00000003_)
	__DiagramProcess__00000001_.ParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000001_.ParticipantWhoseNodeIsExpanded, __Participant__00000004_)
	__DiagramProcess__00000001_.TasksWhoseNodeIsExpanded = append(__DiagramProcess__00000001_.TasksWhoseNodeIsExpanded, __Task__00000009_)
	__DiagramProcess__00000001_.TasksWhoseNodeIsExpanded = append(__DiagramProcess__00000001_.TasksWhoseNodeIsExpanded, __Task__00000011_)
	__DiagramProcess__00000001_.Task_Shapes = append(__DiagramProcess__00000001_.Task_Shapes, __TaskShape__00000008_)
	__DiagramProcess__00000001_.Task_Shapes = append(__DiagramProcess__00000001_.Task_Shapes, __TaskShape__00000009_)
	__DiagramProcess__00000001_.Task_Shapes = append(__DiagramProcess__00000001_.Task_Shapes, __TaskShape__00000010_)
	__DiagramProcess__00000001_.Task_Shapes = append(__DiagramProcess__00000001_.Task_Shapes, __TaskShape__00000011_)
	__DiagramProcess__00000001_.Task_Shapes = append(__DiagramProcess__00000001_.Task_Shapes, __TaskShape__00000012_)
	__DiagramProcess__00000001_.Task_Shapes = append(__DiagramProcess__00000001_.Task_Shapes, __TaskShape__00000013_)
	__DiagramProcess__00000001_.Task_Shapes = append(__DiagramProcess__00000001_.Task_Shapes, __TaskShape__00000014_)
	__DiagramProcess__00000001_.ControlFlow_Shapes = append(__DiagramProcess__00000001_.ControlFlow_Shapes, __ControlFlowShape__00000005_)
	__DiagramProcess__00000001_.ControlFlow_Shapes = append(__DiagramProcess__00000001_.ControlFlow_Shapes, __ControlFlowShape__00000006_)
	__DiagramProcess__00000001_.DataFlow_Shapes = append(__DiagramProcess__00000001_.DataFlow_Shapes, __DataFlowShape__00000005_)
	__DiagramProcess__00000001_.DataFlow_Shapes = append(__DiagramProcess__00000001_.DataFlow_Shapes, __DataFlowShape__00000006_)
	__DiagramProcess__00000001_.DataFlow_Shapes = append(__DiagramProcess__00000001_.DataFlow_Shapes, __DataFlowShape__00000007_)
	__DiagramProcess__00000001_.DataFlow_Shapes = append(__DiagramProcess__00000001_.DataFlow_Shapes, __DataFlowShape__00000008_)
	__DiagramProcess__00000001_.Data_Shapes = append(__DiagramProcess__00000001_.Data_Shapes, __DataShape__00000011_)
	__DiagramProcess__00000001_.Data_Shapes = append(__DiagramProcess__00000001_.Data_Shapes, __DataShape__00000012_)
	__DiagramProcess__00000001_.Data_Shapes = append(__DiagramProcess__00000001_.Data_Shapes, __DataShape__00000013_)
	__DiagramProcess__00000001_.Data_Shapes = append(__DiagramProcess__00000001_.Data_Shapes, __DataShape__00000014_)
	__DiagramProcess__00000001_.DataFlowsWhoseDataNodeIsExpanded = append(__DiagramProcess__00000001_.DataFlowsWhoseDataNodeIsExpanded, __DataFlow__00000005_)
	__DiagramProcess__00000001_.AllocatedResourceShapes = append(__DiagramProcess__00000001_.AllocatedResourceShapes, __AllocatedResourceShape__00000003_)
	__DiagramProcess__00000001_.AllocatedResourceShapes = append(__DiagramProcess__00000001_.AllocatedResourceShapes, __AllocatedResourceShape__00000004_)
	__DiagramProcess__00000001_.AllocatedResourceShapes = append(__DiagramProcess__00000001_.AllocatedResourceShapes, __AllocatedResourceShape__00000005_)
	__DiagramProcess__00000002_.Process_Shapes = append(__DiagramProcess__00000002_.Process_Shapes, __ProcessShape__00000002_)
	__DiagramProcess__00000002_.Participant_Shapes = append(__DiagramProcess__00000002_.Participant_Shapes, __ParticipantShape__00000006_)
	__DiagramProcess__00000002_.Participant_Shapes = append(__DiagramProcess__00000002_.Participant_Shapes, __ParticipantShape__00000007_)
	__DiagramProcess__00000002_.TasksWhoseNodeIsExpanded = append(__DiagramProcess__00000002_.TasksWhoseNodeIsExpanded, __Task__00000016_)
	__DiagramProcess__00000002_.Task_Shapes = append(__DiagramProcess__00000002_.Task_Shapes, __TaskShape__00000015_)
	__DiagramProcess__00000002_.Task_Shapes = append(__DiagramProcess__00000002_.Task_Shapes, __TaskShape__00000016_)
	__DiagramProcess__00000002_.Task_Shapes = append(__DiagramProcess__00000002_.Task_Shapes, __TaskShape__00000017_)
	__DiagramProcess__00000002_.Task_Shapes = append(__DiagramProcess__00000002_.Task_Shapes, __TaskShape__00000018_)
	__DiagramProcess__00000002_.Task_Shapes = append(__DiagramProcess__00000002_.Task_Shapes, __TaskShape__00000019_)
	__DiagramProcess__00000002_.ControlFlow_Shapes = append(__DiagramProcess__00000002_.ControlFlow_Shapes, __ControlFlowShape__00000007_)
	__DiagramProcess__00000002_.ControlFlow_Shapes = append(__DiagramProcess__00000002_.ControlFlow_Shapes, __ControlFlowShape__00000008_)
	__DiagramProcess__00000002_.DataFlow_Shapes = append(__DiagramProcess__00000002_.DataFlow_Shapes, __DataFlowShape__00000009_)
	__DiagramProcess__00000002_.DataFlow_Shapes = append(__DiagramProcess__00000002_.DataFlow_Shapes, __DataFlowShape__00000010_)
	__DiagramProcess__00000002_.Data_Shapes = append(__DiagramProcess__00000002_.Data_Shapes, __DataShape__00000015_)
	__DiagramProcess__00000002_.Data_Shapes = append(__DiagramProcess__00000002_.Data_Shapes, __DataShape__00000016_)
	__DiagramProcess__00000002_.Data_Shapes = append(__DiagramProcess__00000002_.Data_Shapes, __DataShape__00000017_)
	__DiagramProcess__00000002_.Data_Shapes = append(__DiagramProcess__00000002_.Data_Shapes, __DataShape__00000018_)
	__DiagramProcess__00000002_.DataFlowsWhoseDataNodeIsExpanded = append(__DiagramProcess__00000002_.DataFlowsWhoseDataNodeIsExpanded, __DataFlow__00000009_)
	__DiagramProcess__00000002_.AllocatedResourceShapes = append(__DiagramProcess__00000002_.AllocatedResourceShapes, __AllocatedResourceShape__00000006_)
	__DiagramProcess__00000002_.AllocatedResourceShapes = append(__DiagramProcess__00000002_.AllocatedResourceShapes, __AllocatedResourceShape__00000007_)
	__DiagramProcess__00000002_.AllocatedResourceShapes = append(__DiagramProcess__00000002_.AllocatedResourceShapes, __AllocatedResourceShape__00000008_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000000_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000001_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000002_)
	__Library__00000000_.ProcesssWhoseNodeIsExpanded = append(__Library__00000000_.ProcesssWhoseNodeIsExpanded, __Process__00000000_)
	__Library__00000000_.ProcesssWhoseNodeIsExpanded = append(__Library__00000000_.ProcesssWhoseNodeIsExpanded, __Process__00000001_)
	__Library__00000000_.ProcesssWhoseNodeIsExpanded = append(__Library__00000000_.ProcesssWhoseNodeIsExpanded, __Process__00000002_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000000_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000001_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000002_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000003_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000004_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000005_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000006_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000007_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000008_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000009_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000010_)
	__Library__00000000_.DataFlowsWhoseNodeIsExpanded = append(__Library__00000000_.DataFlowsWhoseNodeIsExpanded, __DataFlow__00000005_)
	__Library__00000000_.DataFlowsWhoseNodeIsExpanded = append(__Library__00000000_.DataFlowsWhoseNodeIsExpanded, __DataFlow__00000009_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000000_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000006_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000001_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000002_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000003_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000004_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000005_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000007_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000008_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000009_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000010_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000011_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000012_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000000_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000001_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000002_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000003_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000004_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000005_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000006_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000007_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000008_)
	__Participant__00000000_.Resources = append(__Participant__00000000_.Resources, __Resource__00000000_)
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
	__Participant__00000002_.Resources = append(__Participant__00000002_.Resources, __Resource__00000002_)
	__Participant__00000002_.Tasks = append(__Participant__00000002_.Tasks, __Task__00000004_)
	__Participant__00000002_.Tasks = append(__Participant__00000002_.Tasks, __Task__00000005_)
	__Participant__00000002_.Tasks = append(__Participant__00000002_.Tasks, __Task__00000007_)
	__Participant__00000002_.ControlFlows = append(__Participant__00000002_.ControlFlows, __ControlFlow__00000002_)
	__Participant__00000002_.ControlFlows = append(__Participant__00000002_.ControlFlows, __ControlFlow__00000004_)
	__Participant__00000002_.TaskWhoseOutDataFlowsNodeIsExpanded = append(__Participant__00000002_.TaskWhoseOutDataFlowsNodeIsExpanded, __Task__00000005_)
	__Participant__00000003_.Resources = append(__Participant__00000003_.Resources, __Resource__00000003_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000008_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000012_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000013_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000014_)
	__Participant__00000003_.ControlFlows = append(__Participant__00000003_.ControlFlows, __ControlFlow__00000005_)
	__Participant__00000003_.ControlFlows = append(__Participant__00000003_.ControlFlows, __ControlFlow__00000006_)
	__Participant__00000003_.TaskWhoseOutControlFlowsNodeIsExpanded = append(__Participant__00000003_.TaskWhoseOutControlFlowsNodeIsExpanded, __Task__00000013_)
	__Participant__00000003_.TaskWhoseOutControlFlowsNodeIsExpanded = append(__Participant__00000003_.TaskWhoseOutControlFlowsNodeIsExpanded, __Task__00000012_)
	__Participant__00000003_.TaskWhoseInControlFlowsNodeIsExpanded = append(__Participant__00000003_.TaskWhoseInControlFlowsNodeIsExpanded, __Task__00000014_)
	__Participant__00000003_.TaskWhoseInControlFlowsNodeIsExpanded = append(__Participant__00000003_.TaskWhoseInControlFlowsNodeIsExpanded, __Task__00000008_)
	__Participant__00000003_.TaskWhoseOutDataFlowsNodeIsExpanded = append(__Participant__00000003_.TaskWhoseOutDataFlowsNodeIsExpanded, __Task__00000008_)
	__Participant__00000004_.Resources = append(__Participant__00000004_.Resources, __Resource__00000004_)
	__Participant__00000004_.Tasks = append(__Participant__00000004_.Tasks, __Task__00000009_)
	__Participant__00000004_.Tasks = append(__Participant__00000004_.Tasks, __Task__00000011_)
	__Participant__00000004_.TaskWhoseOutDataFlowsNodeIsExpanded = append(__Participant__00000004_.TaskWhoseOutDataFlowsNodeIsExpanded, __Task__00000009_)
	__Participant__00000004_.TaskWhoseInDataFlowsNodeIsExpanded = append(__Participant__00000004_.TaskWhoseInDataFlowsNodeIsExpanded, __Task__00000009_)
	__Participant__00000004_.TaskWhoseInDataFlowsNodeIsExpanded = append(__Participant__00000004_.TaskWhoseInDataFlowsNodeIsExpanded, __Task__00000011_)
	__Participant__00000005_.Resources = append(__Participant__00000005_.Resources, __Resource__00000005_)
	__Participant__00000005_.Tasks = append(__Participant__00000005_.Tasks, __Task__00000010_)
	__Participant__00000006_.Resources = append(__Participant__00000006_.Resources, __Resource__00000006_)
	__Participant__00000006_.Tasks = append(__Participant__00000006_.Tasks, __Task__00000015_)
	__Participant__00000006_.Tasks = append(__Participant__00000006_.Tasks, __Task__00000017_)
	__Participant__00000006_.Tasks = append(__Participant__00000006_.Tasks, __Task__00000018_)
	__Participant__00000006_.Tasks = append(__Participant__00000006_.Tasks, __Task__00000019_)
	__Participant__00000006_.ControlFlows = append(__Participant__00000006_.ControlFlows, __ControlFlow__00000007_)
	__Participant__00000006_.ControlFlows = append(__Participant__00000006_.ControlFlows, __ControlFlow__00000008_)
	__Participant__00000006_.TaskWhoseOutControlFlowsNodeIsExpanded = append(__Participant__00000006_.TaskWhoseOutControlFlowsNodeIsExpanded, __Task__00000018_)
	__Participant__00000006_.TaskWhoseOutControlFlowsNodeIsExpanded = append(__Participant__00000006_.TaskWhoseOutControlFlowsNodeIsExpanded, __Task__00000017_)
	__Participant__00000006_.TaskWhoseInControlFlowsNodeIsExpanded = append(__Participant__00000006_.TaskWhoseInControlFlowsNodeIsExpanded, __Task__00000019_)
	__Participant__00000006_.TaskWhoseInControlFlowsNodeIsExpanded = append(__Participant__00000006_.TaskWhoseInControlFlowsNodeIsExpanded, __Task__00000015_)
	__Participant__00000006_.TaskWhoseOutDataFlowsNodeIsExpanded = append(__Participant__00000006_.TaskWhoseOutDataFlowsNodeIsExpanded, __Task__00000015_)
	__Participant__00000007_.Resources = append(__Participant__00000007_.Resources, __Resource__00000007_)
	__Participant__00000007_.Tasks = append(__Participant__00000007_.Tasks, __Task__00000016_)
	__Participant__00000007_.TaskWhoseOutDataFlowsNodeIsExpanded = append(__Participant__00000007_.TaskWhoseOutDataFlowsNodeIsExpanded, __Task__00000016_)
	__Participant__00000007_.TaskWhoseInDataFlowsNodeIsExpanded = append(__Participant__00000007_.TaskWhoseInDataFlowsNodeIsExpanded, __Task__00000016_)
	__ParticipantShape__00000000_.Participant = __Participant__00000000_
	__ParticipantShape__00000001_.Participant = __Participant__00000001_
	__ParticipantShape__00000002_.Participant = __Participant__00000002_
	__ParticipantShape__00000003_.Participant = __Participant__00000003_
	__ParticipantShape__00000004_.Participant = __Participant__00000004_
	__ParticipantShape__00000005_.Participant = __Participant__00000005_
	__ParticipantShape__00000006_.Participant = __Participant__00000006_
	__ParticipantShape__00000007_.Participant = __Participant__00000007_
	__Process__00000000_.DiagramProcesss = append(__Process__00000000_.DiagramProcesss, __DiagramProcess__00000000_)
	__Process__00000000_.Participants = append(__Process__00000000_.Participants, __Participant__00000000_)
	__Process__00000000_.Participants = append(__Process__00000000_.Participants, __Participant__00000001_)
	__Process__00000000_.Participants = append(__Process__00000000_.Participants, __Participant__00000002_)
	__Process__00000000_.DataFlows = append(__Process__00000000_.DataFlows, __DataFlow__00000000_)
	__Process__00000000_.DataFlows = append(__Process__00000000_.DataFlows, __DataFlow__00000001_)
	__Process__00000000_.DataFlows = append(__Process__00000000_.DataFlows, __DataFlow__00000002_)
	__Process__00000000_.DataFlows = append(__Process__00000000_.DataFlows, __DataFlow__00000003_)
	__Process__00000000_.DataFlows = append(__Process__00000000_.DataFlows, __DataFlow__00000004_)
	__Process__00000001_.DiagramProcesss = append(__Process__00000001_.DiagramProcesss, __DiagramProcess__00000001_)
	__Process__00000001_.Participants = append(__Process__00000001_.Participants, __Participant__00000003_)
	__Process__00000001_.Participants = append(__Process__00000001_.Participants, __Participant__00000004_)
	__Process__00000001_.Participants = append(__Process__00000001_.Participants, __Participant__00000005_)
	__Process__00000001_.DataFlows = append(__Process__00000001_.DataFlows, __DataFlow__00000005_)
	__Process__00000001_.DataFlows = append(__Process__00000001_.DataFlows, __DataFlow__00000006_)
	__Process__00000001_.DataFlows = append(__Process__00000001_.DataFlows, __DataFlow__00000007_)
	__Process__00000001_.DataFlows = append(__Process__00000001_.DataFlows, __DataFlow__00000008_)
	__Process__00000002_.DiagramProcesss = append(__Process__00000002_.DiagramProcesss, __DiagramProcess__00000002_)
	__Process__00000002_.Participants = append(__Process__00000002_.Participants, __Participant__00000006_)
	__Process__00000002_.Participants = append(__Process__00000002_.Participants, __Participant__00000007_)
	__Process__00000002_.DataFlows = append(__Process__00000002_.DataFlows, __DataFlow__00000009_)
	__Process__00000002_.DataFlows = append(__Process__00000002_.DataFlows, __DataFlow__00000010_)
	__ProcessShape__00000000_.Process = __Process__00000000_
	__ProcessShape__00000001_.Process = __Process__00000001_
	__ProcessShape__00000002_.Process = __Process__00000002_
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
	__TaskShape__00000001_.Task = __Task__00000001_
	__TaskShape__00000002_.Task = __Task__00000002_
	__TaskShape__00000003_.Task = __Task__00000003_
	__TaskShape__00000004_.Task = __Task__00000004_
	__TaskShape__00000005_.Task = __Task__00000005_
	__TaskShape__00000006_.Task = __Task__00000006_
	__TaskShape__00000007_.Task = __Task__00000007_
	__TaskShape__00000008_.Task = __Task__00000008_
	__TaskShape__00000009_.Task = __Task__00000009_
	__TaskShape__00000010_.Task = __Task__00000010_
	__TaskShape__00000011_.Task = __Task__00000011_
	__TaskShape__00000012_.Task = __Task__00000012_
	__TaskShape__00000013_.Task = __Task__00000013_
	__TaskShape__00000014_.Task = __Task__00000014_
	__TaskShape__00000015_.Task = __Task__00000015_
	__TaskShape__00000016_.Task = __Task__00000016_
	__TaskShape__00000017_.Task = __Task__00000017_
	__TaskShape__00000018_.Task = __Task__00000018_
	__TaskShape__00000019_.Task = __Task__00000019_
}
