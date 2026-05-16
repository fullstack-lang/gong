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

	__AllocatedProcessShape__00000000_ := (&models.AllocatedProcessShape{Name: `Diagram P1-EP1 Very Long Participant Name-P1`}).Stage(stage)
	__AllocatedProcessShape__00000001_ := (&models.AllocatedProcessShape{Name: `Diagram P1-EP1 Very Long Participant Name-P2`}).Stage(stage)
	__AllocatedProcessShape__00000002_ := (&models.AllocatedProcessShape{Name: `Diagram P1-EP2-P1`}).Stage(stage)

	__AllocatedResourceShape__00000010_ := (&models.AllocatedResourceShape{Name: `Diagram P1-EP1 Very Long Participant Name-R2`}).Stage(stage)
	__AllocatedResourceShape__00000012_ := (&models.AllocatedResourceShape{Name: `Diagram P1-EP1 Very Long Participant Name-R1 with a long name long long long`}).Stage(stage)
	__AllocatedResourceShape__00000014_ := (&models.AllocatedResourceShape{Name: `Diagram P1-PP1 Very lon participant name PP1 Very lon participant name-R2aaaaa`}).Stage(stage)
	__AllocatedResourceShape__00000015_ := (&models.AllocatedResourceShape{Name: `Diagram P1-PP1 Very lon participant name PP1 Very lon participant name-R1 with a long name long long long`}).Stage(stage)
	__AllocatedResourceShape__00000016_ := (&models.AllocatedResourceShape{Name: `D2-P2.P1-R2aaaaa`}).Stage(stage)
	__AllocatedResourceShape__00000017_ := (&models.AllocatedResourceShape{Name: `Diagram P1-EP2-R2aaaaa`}).Stage(stage)

	__ControlFlow__00000003_ := (&models.ControlFlow{Name: `"Start" to "P1"`}).Stage(stage)
	__ControlFlow__00000004_ := (&models.ControlFlow{Name: `"P1" to "T4"`}).Stage(stage)
	__ControlFlow__00000005_ := (&models.ControlFlow{Name: `"PP2 T1" to "PP2 T2"`}).Stage(stage)
	__ControlFlow__00000006_ := (&models.ControlFlow{Name: `"T4" to "End"`}).Stage(stage)
	__ControlFlow__00000007_ := (&models.ControlFlow{Name: `"Start" to "T1"`}).Stage(stage)
	__ControlFlow__00000008_ := (&models.ControlFlow{Name: `"T1" to "End"`}).Stage(stage)
	__ControlFlow__00000009_ := (&models.ControlFlow{Name: `"P2.P2.Start" to "P2.P2.T1"`}).Stage(stage)

	__ControlFlowShape__00000000_ := (&models.ControlFlowShape{Name: `"Start" to "P1"`}).Stage(stage)
	__ControlFlowShape__00000001_ := (&models.ControlFlowShape{Name: `"PP2 T1" to "PP2 T2"`}).Stage(stage)
	__ControlFlowShape__00000002_ := (&models.ControlFlowShape{Name: `"P1" to "T4"`}).Stage(stage)
	__ControlFlowShape__00000003_ := (&models.ControlFlowShape{Name: `"T4" to "End"`}).Stage(stage)
	__ControlFlowShape__00000004_ := (&models.ControlFlowShape{Name: `"Start" to "T1"`}).Stage(stage)
	__ControlFlowShape__00000005_ := (&models.ControlFlowShape{Name: `"T1" to "End"`}).Stage(stage)
	__ControlFlowShape__00000006_ := (&models.ControlFlowShape{Name: `"P2.P2.Start" to "P2.P2.T1"`}).Stage(stage)

	__Data__00000000_ := (&models.Data{Name: `D1`}).Stage(stage)
	__Data__00000001_ := (&models.Data{Name: `D2`}).Stage(stage)
	__Data__00000002_ := (&models.Data{Name: `L1.D1`}).Stage(stage)
	__Data__00000003_ := (&models.Data{Name: `D3D3D3D3D3`}).Stage(stage)

	__DataFlow__00000000_ := (&models.DataFlow{Name: `"P1" to "PP2 T2"`}).Stage(stage)
	__DataFlow__00000002_ := (&models.DataFlow{Name: `"T1" to "P2.P2.T1"`}).Stage(stage)
	__DataFlow__00000003_ := (&models.DataFlow{Name: `EP1 Very Long Participant Name to P1`}).Stage(stage)
	__DataFlow__00000005_ := (&models.DataFlow{Name: `EP2 to PP2 T2`}).Stage(stage)

	__DataFlowShape__00000004_ := (&models.DataFlowShape{Name: `"P1" to "PP2 T2"`}).Stage(stage)
	__DataFlowShape__00000005_ := (&models.DataFlowShape{Name: `"T1" to "P2.P2.T1"`}).Stage(stage)
	__DataFlowShape__00000006_ := (&models.DataFlowShape{Name: `EP1 Very Long Participant Name to P1`}).Stage(stage)
	__DataFlowShape__00000008_ := (&models.DataFlowShape{Name: `EP2 to PP2 T2`}).Stage(stage)

	__DataShape__00000011_ := (&models.DataShape{Name: `T2 to PP2 T2-T2 to PP2 T2-Diagram P1`}).Stage(stage)
	__DataShape__00000012_ := (&models.DataShape{Name: `T2 to PP2 T2-T2 to PP2 T2-Diagram P1`}).Stage(stage)
	__DataShape__00000013_ := (&models.DataShape{Name: `PP2 T2 to T2 2-PP2 T2 to T2-Diagram P1`}).Stage(stage)
	__DataShape__00000014_ := (&models.DataShape{Name: `PP2 T2 to T2 2-PP2 T2 to T2-Diagram P1`}).Stage(stage)
	__DataShape__00000016_ := (&models.DataShape{Name: `PP2 T2 to T2-PP2 T2 to T2-Diagram P1`}).Stage(stage)
	__DataShape__00000017_ := (&models.DataShape{Name: `PP2 T2 to T2-PP2 T2 to T2-Diagram P1`}).Stage(stage)
	__DataShape__00000019_ := (&models.DataShape{Name: `T2 to PP2 T2-T2 to PP2 T2-Diagram P1`}).Stage(stage)
	__DataShape__00000022_ := (&models.DataShape{Name: `EP1 to T2-EP1 to T2-Diagram P1`}).Stage(stage)
	__DataShape__00000024_ := (&models.DataShape{Name: `EP1 to T2-EP1 to T2-Diagram P1`}).Stage(stage)
	__DataShape__00000025_ := (&models.DataShape{Name: `T1 to P2.P2.T1-T1 to P2.P2.T1-DiagramProcess`}).Stage(stage)
	__DataShape__00000026_ := (&models.DataShape{Name: `EP1 Very Long Participant Name to P1-EP1 Very Long Participant Name to P1-Diagram P1`}).Stage(stage)
	__DataShape__00000027_ := (&models.DataShape{Name: `EP1 Very Long Participant Name to P1-EP1 Very Long Participant Name to P1-Diagram P1`}).Stage(stage)

	__DiagramProcess__00000006_ := (&models.DiagramProcess{Name: `DiagramProcess`}).Stage(stage)
	__DiagramProcess__00000007_ := (&models.DiagramProcess{Name: `Diagram P1`}).Stage(stage)
	__DiagramProcess__00000009_ := (&models.DiagramProcess{Name: `D2`}).Stage(stage)

	__ExternalParticipantShape__00000003_ := (&models.ExternalParticipantShape{Name: `EP1-Diagram P1`}).Stage(stage)
	__ExternalParticipantShape__00000005_ := (&models.ExternalParticipantShape{Name: `EP2-Diagram P1`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `Root`}).Stage(stage)

	__Note__00000000_ := (&models.Note{Name: `Note on the diagram`}).Stage(stage)

	__NoteShape__00000000_ := (&models.NoteShape{Name: `Note on the diagram shape`}).Stage(stage)

	__NoteTaskShape__00000000_ := (&models.NoteTaskShape{Name: `P1 shape`}).Stage(stage)

	__Participant__00000003_ := (&models.Participant{Name: `PP1 Very lon participant name PP1 Very lon participant name`}).Stage(stage)
	__Participant__00000004_ := (&models.Participant{Name: `PP2`}).Stage(stage)
	__Participant__00000005_ := (&models.Participant{Name: `PP3`}).Stage(stage)
	__Participant__00000006_ := (&models.Participant{Name: `PP4`}).Stage(stage)
	__Participant__00000007_ := (&models.Participant{Name: `EP1 Very Long Participant Name`}).Stage(stage)
	__Participant__00000008_ := (&models.Participant{Name: `P2.P1`}).Stage(stage)
	__Participant__00000009_ := (&models.Participant{Name: `P2.P2`}).Stage(stage)
	__Participant__00000011_ := (&models.Participant{Name: `EP2`}).Stage(stage)

	__ParticipantShape__00000019_ := (&models.ParticipantShape{Name: `PP3-Diagram P1`}).Stage(stage)
	__ParticipantShape__00000020_ := (&models.ParticipantShape{Name: `PP2-Diagram P1`}).Stage(stage)
	__ParticipantShape__00000022_ := (&models.ParticipantShape{Name: `PP1 Very lon participant name PP1 Very lon participant name-Diagram P1`}).Stage(stage)
	__ParticipantShape__00000023_ := (&models.ParticipantShape{Name: `-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000024_ := (&models.ParticipantShape{Name: `P2.P2-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000025_ := (&models.ParticipantShape{Name: `P2.P1-D2`}).Stage(stage)

	__Process__00000005_ := (&models.Process{Name: `P1`}).Stage(stage)
	__Process__00000006_ := (&models.Process{Name: `P2`}).Stage(stage)

	__ProcessShape__00000006_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)
	__ProcessShape__00000007_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)
	__ProcessShape__00000009_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)

	__Resource__00000000_ := (&models.Resource{Name: `R1 with a long name long long long`}).Stage(stage)
	__Resource__00000001_ := (&models.Resource{Name: `R2aaaaa`}).Stage(stage)

	__Task__00000002_ := (&models.Task{Name: `P1`}).Stage(stage)
	__Task__00000003_ := (&models.Task{Name: `PP2 T1`}).Stage(stage)
	__Task__00000004_ := (&models.Task{Name: `Start`}).Stage(stage)
	__Task__00000005_ := (&models.Task{Name: `End`}).Stage(stage)
	__Task__00000006_ := (&models.Task{Name: `T4`}).Stage(stage)
	__Task__00000007_ := (&models.Task{Name: `PP2 T2`}).Stage(stage)
	__Task__00000008_ := (&models.Task{Name: `Start`}).Stage(stage)
	__Task__00000009_ := (&models.Task{Name: `T1`}).Stage(stage)
	__Task__00000010_ := (&models.Task{Name: `End`}).Stage(stage)
	__Task__00000011_ := (&models.Task{Name: `P2.P2.Start`}).Stage(stage)
	__Task__00000012_ := (&models.Task{Name: `P2.P2.T1`}).Stage(stage)

	__TaskShape__00000011_ := (&models.TaskShape{Name: `Start-DP2`}).Stage(stage)
	__TaskShape__00000012_ := (&models.TaskShape{Name: `T4-DP2`}).Stage(stage)
	__TaskShape__00000013_ := (&models.TaskShape{Name: `T2-DP2`}).Stage(stage)
	__TaskShape__00000014_ := (&models.TaskShape{Name: `PP2 T2-DP2`}).Stage(stage)
	__TaskShape__00000015_ := (&models.TaskShape{Name: `PP2 T1-Diagram P1`}).Stage(stage)
	__TaskShape__00000016_ := (&models.TaskShape{Name: `End-Diagram P1`}).Stage(stage)
	__TaskShape__00000019_ := (&models.TaskShape{Name: `-DiagramProcess`}).Stage(stage)
	__TaskShape__00000020_ := (&models.TaskShape{Name: `Start-DiagramProcess`}).Stage(stage)
	__TaskShape__00000021_ := (&models.TaskShape{Name: `-DiagramProcess`}).Stage(stage)
	__TaskShape__00000022_ := (&models.TaskShape{Name: `-DiagramProcess`}).Stage(stage)
	__TaskShape__00000023_ := (&models.TaskShape{Name: `-DiagramProcess`}).Stage(stage)

	// insertion point for initialization of values

	__AllocatedProcessShape__00000000_.Name = `Diagram P1-EP1 Very Long Participant Name-P1`

	__AllocatedProcessShape__00000001_.Name = `Diagram P1-EP1 Very Long Participant Name-P2`

	__AllocatedProcessShape__00000002_.Name = `Diagram P1-EP2-P1`

	__AllocatedResourceShape__00000010_.Name = `Diagram P1-EP1 Very Long Participant Name-R2`

	__AllocatedResourceShape__00000012_.Name = `Diagram P1-EP1 Very Long Participant Name-R1 with a long name long long long`

	__AllocatedResourceShape__00000014_.Name = `Diagram P1-PP1 Very lon participant name PP1 Very lon participant name-R2aaaaa`

	__AllocatedResourceShape__00000015_.Name = `Diagram P1-PP1 Very lon participant name PP1 Very lon participant name-R1 with a long name long long long`

	__AllocatedResourceShape__00000016_.Name = `D2-P2.P1-R2aaaaa`

	__AllocatedResourceShape__00000017_.Name = `Diagram P1-EP2-R2aaaaa`

	__ControlFlow__00000003_.Name = `"Start" to "P1"`
	__ControlFlow__00000003_.Description = ``
	__ControlFlow__00000003_.ComputedPrefix = ``

	__ControlFlow__00000004_.Name = `"P1" to "T4"`
	__ControlFlow__00000004_.Description = ``
	__ControlFlow__00000004_.ComputedPrefix = ``

	__ControlFlow__00000005_.Name = `"PP2 T1" to "PP2 T2"`
	__ControlFlow__00000005_.Description = ``
	__ControlFlow__00000005_.ComputedPrefix = ``

	__ControlFlow__00000006_.Name = `"T4" to "End"`
	__ControlFlow__00000006_.Description = ``
	__ControlFlow__00000006_.ComputedPrefix = ``

	__ControlFlow__00000007_.Name = `"Start" to "T1"`
	__ControlFlow__00000007_.Description = ``
	__ControlFlow__00000007_.ComputedPrefix = ``

	__ControlFlow__00000008_.Name = `"T1" to "End"`
	__ControlFlow__00000008_.Description = ``
	__ControlFlow__00000008_.ComputedPrefix = ``

	__ControlFlow__00000009_.Name = `"P2.P2.Start" to "P2.P2.T1"`
	__ControlFlow__00000009_.Description = ``
	__ControlFlow__00000009_.ComputedPrefix = ``

	__ControlFlowShape__00000000_.Name = `"Start" to "P1"`
	__ControlFlowShape__00000000_.StartRatio = 0.971920
	__ControlFlowShape__00000000_.EndRatio = 0.500000
	__ControlFlowShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000000_.CornerOffsetRatio = 2.278388
	__ControlFlowShape__00000000_.IsHidden = false

	__ControlFlowShape__00000001_.Name = `"PP2 T1" to "PP2 T2"`
	__ControlFlowShape__00000001_.StartRatio = 1.000000
	__ControlFlowShape__00000001_.EndRatio = 0.500000
	__ControlFlowShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000001_.CornerOffsetRatio = 2.823598
	__ControlFlowShape__00000001_.IsHidden = false

	__ControlFlowShape__00000002_.Name = `"P1" to "T4"`
	__ControlFlowShape__00000002_.StartRatio = 0.500000
	__ControlFlowShape__00000002_.EndRatio = 0.789929
	__ControlFlowShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000002_.CornerOffsetRatio = 1.100000
	__ControlFlowShape__00000002_.IsHidden = false

	__ControlFlowShape__00000003_.Name = `"T4" to "End"`
	__ControlFlowShape__00000003_.StartRatio = 0.500000
	__ControlFlowShape__00000003_.EndRatio = 0.793912
	__ControlFlowShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000003_.CornerOffsetRatio = 1.500000
	__ControlFlowShape__00000003_.IsHidden = false

	__ControlFlowShape__00000004_.Name = `"Start" to "T1"`
	__ControlFlowShape__00000004_.StartRatio = 0.500000
	__ControlFlowShape__00000004_.EndRatio = 0.500000
	__ControlFlowShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000004_.CornerOffsetRatio = 1.500000
	__ControlFlowShape__00000004_.IsHidden = false

	__ControlFlowShape__00000005_.Name = `"T1" to "End"`
	__ControlFlowShape__00000005_.StartRatio = 0.500000
	__ControlFlowShape__00000005_.EndRatio = 0.500000
	__ControlFlowShape__00000005_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000005_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000005_.CornerOffsetRatio = 1.500000
	__ControlFlowShape__00000005_.IsHidden = false

	__ControlFlowShape__00000006_.Name = `"P2.P2.Start" to "P2.P2.T1"`
	__ControlFlowShape__00000006_.StartRatio = 0.500000
	__ControlFlowShape__00000006_.EndRatio = 0.500000
	__ControlFlowShape__00000006_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000006_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000006_.CornerOffsetRatio = 1.500000
	__ControlFlowShape__00000006_.IsHidden = false

	__Data__00000000_.Name = `D1`
	__Data__00000000_.Description = ``
	__Data__00000000_.ComputedPrefix = ``
	__Data__00000000_.SVG_Path = ``
	__Data__00000000_.InverseAppliedScaling = 60.000000

	__Data__00000001_.Name = `D2`
	__Data__00000001_.Description = ``
	__Data__00000001_.ComputedPrefix = ``
	__Data__00000001_.SVG_Path = `m280-40 112-564-72 28v136h-80v-188l202-86q14-6 29.5-7t29.5 4q14 5 26.5 14t20.5 23l40 64q26 42 70.5 69T760-520v80q-70 0-125-29t-94-74l-25 123 84 80v300h-80v-260l-84-64-72 324h-84Zm203.5-723.5Q460-787 460-820t23.5-56.5Q507-900 540-900t56.5 23.5Q620-853 620-820t-23.5 56.5Q573-740 540-740t-56.5-23.5Z`
	__Data__00000001_.InverseAppliedScaling = 60.000000

	__Data__00000002_.Name = `L1.D1`
	__Data__00000002_.Description = ``
	__Data__00000002_.ComputedPrefix = ``
	__Data__00000002_.SVG_Path = ``
	__Data__00000002_.InverseAppliedScaling = 0.000000

	__Data__00000003_.Name = `D3D3D3D3D3`
	__Data__00000003_.Description = ``
	__Data__00000003_.ComputedPrefix = ``
	__Data__00000003_.SVG_Path = ``
	__Data__00000003_.InverseAppliedScaling = 1.000000

	__DataFlow__00000000_.Name = `"P1" to "PP2 T2"`
	__DataFlow__00000000_.Description = ``
	__DataFlow__00000000_.ComputedPrefix = ``
	__DataFlow__00000000_.Type = models.DataFlow_Task2Task
	__DataFlow__00000000_.IsDatasNodeExpanded = false

	__DataFlow__00000002_.Name = `"T1" to "P2.P2.T1"`
	__DataFlow__00000002_.Description = ``
	__DataFlow__00000002_.ComputedPrefix = ``
	__DataFlow__00000002_.Type = models.DataFlow_Task2Task
	__DataFlow__00000002_.IsDatasNodeExpanded = false

	__DataFlow__00000003_.Name = `EP1 Very Long Participant Name to P1`
	__DataFlow__00000003_.Description = ``
	__DataFlow__00000003_.ComputedPrefix = ``
	__DataFlow__00000003_.Type = models.DataFlow_ExternalParticipant2Task
	__DataFlow__00000003_.IsDatasNodeExpanded = false

	__DataFlow__00000005_.Name = `EP2 to PP2 T2`
	__DataFlow__00000005_.Description = ``
	__DataFlow__00000005_.ComputedPrefix = ``
	__DataFlow__00000005_.Type = models.DataFlow_ExternalParticipant2Task
	__DataFlow__00000005_.IsDatasNodeExpanded = false

	__DataFlowShape__00000004_.Name = `"P1" to "PP2 T2"`
	__DataFlowShape__00000004_.StartRatio = 0.741696
	__DataFlowShape__00000004_.EndRatio = 0.163649
	__DataFlowShape__00000004_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__DataFlowShape__00000004_.CornerOffsetRatio = 1.008329
	__DataFlowShape__00000004_.IsHidden = false

	__DataFlowShape__00000005_.Name = `"T1" to "P2.P2.T1"`
	__DataFlowShape__00000005_.StartRatio = 1.000000
	__DataFlowShape__00000005_.EndRatio = 0.302071
	__DataFlowShape__00000005_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000005_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000005_.CornerOffsetRatio = 4.495882
	__DataFlowShape__00000005_.IsHidden = false

	__DataFlowShape__00000006_.Name = `EP1 Very Long Participant Name to P1`
	__DataFlowShape__00000006_.StartRatio = 0.584678
	__DataFlowShape__00000006_.EndRatio = 0.545862
	__DataFlowShape__00000006_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000006_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000006_.CornerOffsetRatio = 24.620834
	__DataFlowShape__00000006_.IsHidden = false

	__DataFlowShape__00000008_.Name = `EP2 to PP2 T2`
	__DataFlowShape__00000008_.StartRatio = 0.500000
	__DataFlowShape__00000008_.EndRatio = 0.800099
	__DataFlowShape__00000008_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000008_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000008_.CornerOffsetRatio = -8.704439
	__DataFlowShape__00000008_.IsHidden = false

	__DataShape__00000011_.Name = `T2 to PP2 T2-T2 to PP2 T2-Diagram P1`

	__DataShape__00000012_.Name = `T2 to PP2 T2-T2 to PP2 T2-Diagram P1`

	__DataShape__00000013_.Name = `PP2 T2 to T2 2-PP2 T2 to T2-Diagram P1`

	__DataShape__00000014_.Name = `PP2 T2 to T2 2-PP2 T2 to T2-Diagram P1`

	__DataShape__00000016_.Name = `PP2 T2 to T2-PP2 T2 to T2-Diagram P1`

	__DataShape__00000017_.Name = `PP2 T2 to T2-PP2 T2 to T2-Diagram P1`

	__DataShape__00000019_.Name = `T2 to PP2 T2-T2 to PP2 T2-Diagram P1`

	__DataShape__00000022_.Name = `EP1 to T2-EP1 to T2-Diagram P1`

	__DataShape__00000024_.Name = `EP1 to T2-EP1 to T2-Diagram P1`

	__DataShape__00000025_.Name = `T1 to P2.P2.T1-T1 to P2.P2.T1-DiagramProcess`

	__DataShape__00000026_.Name = `EP1 Very Long Participant Name to P1-EP1 Very Long Participant Name to P1-Diagram P1`

	__DataShape__00000027_.Name = `EP1 Very Long Participant Name to P1-EP1 Very Long Participant Name to P1-Diagram P1`

	__DiagramProcess__00000006_.Name = `DiagramProcess`
	__DiagramProcess__00000006_.Description = ``
	__DiagramProcess__00000006_.ComputedPrefix = ``
	__DiagramProcess__00000006_.IsChecked = true
	__DiagramProcess__00000006_.IsEditable_ = true
	__DiagramProcess__00000006_.IsShowPrefix = false
	__DiagramProcess__00000006_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000006_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000006_.Width = 1253.000000
	__DiagramProcess__00000006_.Height = 1349.000000
	__DiagramProcess__00000006_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000006_.IsParticipantsNodeExpanded = false
	__DiagramProcess__00000006_.IsExternalParticipantsNodeExpanded = false
	__DiagramProcess__00000006_.IsNotesNodeExpanded = false

	__DiagramProcess__00000007_.Name = `Diagram P1`
	__DiagramProcess__00000007_.Description = ``
	__DiagramProcess__00000007_.ComputedPrefix = ``
	__DiagramProcess__00000007_.IsChecked = false
	__DiagramProcess__00000007_.IsEditable_ = true
	__DiagramProcess__00000007_.IsShowPrefix = false
	__DiagramProcess__00000007_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000007_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000007_.Width = 1574.015236
	__DiagramProcess__00000007_.Height = 1522.000000
	__DiagramProcess__00000007_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000007_.IsParticipantsNodeExpanded = false
	__DiagramProcess__00000007_.IsExternalParticipantsNodeExpanded = true
	__DiagramProcess__00000007_.IsNotesNodeExpanded = true

	__DiagramProcess__00000009_.Name = `D2`
	__DiagramProcess__00000009_.Description = ``
	__DiagramProcess__00000009_.ComputedPrefix = ``
	__DiagramProcess__00000009_.IsChecked = false
	__DiagramProcess__00000009_.IsEditable_ = true
	__DiagramProcess__00000009_.IsShowPrefix = false
	__DiagramProcess__00000009_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000009_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000009_.Width = 900.000000
	__DiagramProcess__00000009_.Height = 1350.000000
	__DiagramProcess__00000009_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000009_.IsParticipantsNodeExpanded = true
	__DiagramProcess__00000009_.IsExternalParticipantsNodeExpanded = false
	__DiagramProcess__00000009_.IsNotesNodeExpanded = false

	__ExternalParticipantShape__00000003_.Name = `EP1-Diagram P1`
	__ExternalParticipantShape__00000003_.IsExpanded = false
	__ExternalParticipantShape__00000003_.X = 42.388870
	__ExternalParticipantShape__00000003_.Y = 344.256840
	__ExternalParticipantShape__00000003_.Width = 211.000000
	__ExternalParticipantShape__00000003_.Height = 70.000000
	__ExternalParticipantShape__00000003_.IsHidden = false
	__ExternalParticipantShape__00000003_.TailHeigth = 233.000000

	__ExternalParticipantShape__00000005_.Name = `EP2-Diagram P1`
	__ExternalParticipantShape__00000005_.IsExpanded = false
	__ExternalParticipantShape__00000005_.X = 1121.015236
	__ExternalParticipantShape__00000005_.Y = 199.292561
	__ExternalParticipantShape__00000005_.Width = 153.000000
	__ExternalParticipantShape__00000005_.Height = 55.000000
	__ExternalParticipantShape__00000005_.IsHidden = false
	__ExternalParticipantShape__00000005_.TailHeigth = 100.000000

	__Library__00000000_.Name = `Root`
	__Library__00000000_.Description = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsSubLibrariesNodeExpanded = true
	__Library__00000000_.NbPixPerCharacter = 11.000000
	__Library__00000000_.LogoSVGFile = `<?xml version="1.0"?> <svg width="508.204" height="141.732" xmlns="http://www.w3.org/2000/svg"> <path fill="#DB362D" d="M91.991,104.699c1.576,5.961,4.119,8.266,8.613,8.266c4.659,0,7.102-2.799,7.102-8.266V3.2h29.184v101.499 c0,14.307-1.856,20.506-9.11,27.762c-5.228,5.229-14.871,9.271-27.047,9.271c-9.837,0-19.25-3.256-25.253-9.27 c-5.263-5.273-8.154-10.689-12.672-27.764L44.9,37.033c-1.577-5.961-4.119-8.265-8.613-8.265c-4.66,0-7.103,2.798-7.103,8.265 v101.5H0v-101.5C0,22.727,1.857,16.527,9.111,9.271C14.337,4.044,23.981,0,36.158,0c9.837,0,19.25,3.257,25.253,9.27 c5.263,5.273,8.154,10.689,12.672,27.764L91.991,104.699z"/>  <path fill="#DB362D" d="M478.038,138.533L444.334,33.096c-0.372-1.164-0.723-2.152-1.263-2.811 c-0.926-1.127-2.207-1.719-3.931-1.719c-1.723,0-3.004,0.592-3.931,1.719c-0.539,0.658-0.891,1.646-1.262,2.811l-33.703,105.437 h-30.167l36.815-115.177c1.918-6,4.66-11.094,8.139-14.488C421.002,3.047,428.038,0,439.141,0s18.14,3.047,24.109,8.867 c3.479,3.395,6.221,8.488,8.14,14.488l36.814,115.177H478.038z"/>  <path fill="#DB362D" d="M328.878,138.533c19.12,0,28.446-4.062,35.814-11.389c8.153-8.105,12.053-16.973,12.053-30.213 c0-11.699-4.283-22.535-10.804-29.019c-8.526-8.479-19.116-11.151-36.384-11.151L305.37,56.76c-9.242,0-12.925-1.117-15.839-3.98 c-2.001-1.964-2.939-4.885-2.939-8.328c0-3.559,0.857-7.074,3.303-9.475c2.171-2.131,5.13-3.109,10.816-3.109h69.903V3.2H306.05 c-19.12,0-28.445,4.063-35.814,11.389c-8.152,8.105-12.053,16.972-12.053,30.212c0,11.701,4.283,22.536,10.804,29.019 c8.527,8.479,19.116,11.152,36.384,11.152l24.188,0.002c9.242,0,12.925,1.115,15.839,3.979c2.001,1.965,2.939,4.885,2.939,8.328 c0,3.559-0.857,7.074-3.302,9.475c-2.172,2.131-5.131,3.109-10.817,3.109h-72.094l-27.651-86.509 c-1.918-6-4.66-11.094-8.139-14.488C220.363,3.047,213.327,0,202.224,0s-18.14,3.047-24.108,8.867 c-3.48,3.395-6.221,8.488-8.139,14.488l-36.815,115.177h30.166l33.704-105.437c0.372-1.164,0.723-2.152,1.263-2.811 c0.926-1.127,2.208-1.719,3.931-1.719s3.004,0.592,3.931,1.719c0.54,0.658,0.891,1.646,1.262,2.811l33.704,105.437H328.878z"/> </svg>`
	__Library__00000000_.IsProcessesNodeExpanded = true
	__Library__00000000_.IsDataFlowsNodeExpanded = false
	__Library__00000000_.IsDatasNodeExpanded = false
	__Library__00000000_.IsResourcesNodeExpanded = false
	__Library__00000000_.IsNotesNodeExpanded = true
	__Library__00000000_.IsExpandedTmp = true

	__Note__00000000_.Name = `Note on the diagram`
	__Note__00000000_.ComputedPrefix = ``
	__Note__00000000_.IsTasksNodeExpanded = true

	__NoteShape__00000000_.Name = `Note on the diagram shape`
	__NoteShape__00000000_.X = 3.261263
	__NoteShape__00000000_.Y = 136.057359
	__NoteShape__00000000_.Width = 173.000000
	__NoteShape__00000000_.Height = 70.000000
	__NoteShape__00000000_.IsHidden = false

	__NoteTaskShape__00000000_.Name = `P1 shape`
	__NoteTaskShape__00000000_.StartRatio = 0.000000
	__NoteTaskShape__00000000_.EndRatio = 0.000000
	__NoteTaskShape__00000000_.StartOrientation = ""
	__NoteTaskShape__00000000_.EndOrientation = ""
	__NoteTaskShape__00000000_.CornerOffsetRatio = 0.000000
	__NoteTaskShape__00000000_.IsHidden = false

	__Participant__00000003_.Name = `PP1 Very lon participant name PP1 Very lon participant name`
	__Participant__00000003_.IsProcessResource = false
	__Participant__00000003_.Description = ``
	__Participant__00000003_.IsResourcesNodeExpanded = false
	__Participant__00000003_.IsProcessesNodeExpanded = false
	__Participant__00000003_.ComputedPrefix = ``
	__Participant__00000003_.IsTasksNodeExpanded = false
	__Participant__00000003_.IsControlFlowsNodeExpanded = false
	__Participant__00000003_.IsDataFlowsNodeExpanded = false

	__Participant__00000004_.Name = `PP2`
	__Participant__00000004_.IsProcessResource = false
	__Participant__00000004_.Description = ``
	__Participant__00000004_.IsResourcesNodeExpanded = false
	__Participant__00000004_.IsProcessesNodeExpanded = false
	__Participant__00000004_.ComputedPrefix = ``
	__Participant__00000004_.IsTasksNodeExpanded = false
	__Participant__00000004_.IsControlFlowsNodeExpanded = false
	__Participant__00000004_.IsDataFlowsNodeExpanded = false

	__Participant__00000005_.Name = `PP3`
	__Participant__00000005_.IsProcessResource = false
	__Participant__00000005_.Description = ``
	__Participant__00000005_.IsResourcesNodeExpanded = false
	__Participant__00000005_.IsProcessesNodeExpanded = false
	__Participant__00000005_.ComputedPrefix = ``
	__Participant__00000005_.IsTasksNodeExpanded = false
	__Participant__00000005_.IsControlFlowsNodeExpanded = false
	__Participant__00000005_.IsDataFlowsNodeExpanded = false

	__Participant__00000006_.Name = `PP4`
	__Participant__00000006_.IsProcessResource = false
	__Participant__00000006_.Description = ``
	__Participant__00000006_.IsResourcesNodeExpanded = false
	__Participant__00000006_.IsProcessesNodeExpanded = false
	__Participant__00000006_.ComputedPrefix = ``
	__Participant__00000006_.IsTasksNodeExpanded = false
	__Participant__00000006_.IsControlFlowsNodeExpanded = false
	__Participant__00000006_.IsDataFlowsNodeExpanded = false

	__Participant__00000007_.Name = `EP1 Very Long Participant Name`
	__Participant__00000007_.IsProcessResource = false
	__Participant__00000007_.Description = ``
	__Participant__00000007_.IsResourcesNodeExpanded = true
	__Participant__00000007_.IsProcessesNodeExpanded = true
	__Participant__00000007_.ComputedPrefix = ``
	__Participant__00000007_.IsTasksNodeExpanded = false
	__Participant__00000007_.IsControlFlowsNodeExpanded = false
	__Participant__00000007_.IsDataFlowsNodeExpanded = false

	__Participant__00000008_.Name = `P2.P1`
	__Participant__00000008_.IsProcessResource = false
	__Participant__00000008_.Description = ``
	__Participant__00000008_.IsResourcesNodeExpanded = true
	__Participant__00000008_.IsProcessesNodeExpanded = false
	__Participant__00000008_.ComputedPrefix = ``
	__Participant__00000008_.IsTasksNodeExpanded = true
	__Participant__00000008_.IsControlFlowsNodeExpanded = false
	__Participant__00000008_.IsDataFlowsNodeExpanded = false

	__Participant__00000009_.Name = `P2.P2`
	__Participant__00000009_.IsProcessResource = false
	__Participant__00000009_.Description = ``
	__Participant__00000009_.IsResourcesNodeExpanded = false
	__Participant__00000009_.IsProcessesNodeExpanded = false
	__Participant__00000009_.ComputedPrefix = ``
	__Participant__00000009_.IsTasksNodeExpanded = true
	__Participant__00000009_.IsControlFlowsNodeExpanded = false
	__Participant__00000009_.IsDataFlowsNodeExpanded = false

	__Participant__00000011_.Name = `EP2`
	__Participant__00000011_.IsProcessResource = false
	__Participant__00000011_.Description = ``
	__Participant__00000011_.IsResourcesNodeExpanded = true
	__Participant__00000011_.IsProcessesNodeExpanded = true
	__Participant__00000011_.ComputedPrefix = ``
	__Participant__00000011_.IsTasksNodeExpanded = false
	__Participant__00000011_.IsControlFlowsNodeExpanded = false
	__Participant__00000011_.IsDataFlowsNodeExpanded = false

	__ParticipantShape__00000019_.Name = `PP3-Diagram P1`
	__ParticipantShape__00000019_.IsExpanded = false
	__ParticipantShape__00000019_.X = 156.650926
	__ParticipantShape__00000019_.Y = 69.648379
	__ParticipantShape__00000019_.Width = 250.000000
	__ParticipantShape__00000019_.Height = 70.000000
	__ParticipantShape__00000019_.IsHidden = false
	__ParticipantShape__00000019_.WidthWeight = 1.000000

	__ParticipantShape__00000020_.Name = `PP2-Diagram P1`
	__ParticipantShape__00000020_.IsExpanded = false
	__ParticipantShape__00000020_.X = 608.000000
	__ParticipantShape__00000020_.Y = 185.000000
	__ParticipantShape__00000020_.Width = 303.000000
	__ParticipantShape__00000020_.Height = 1037.000000
	__ParticipantShape__00000020_.IsHidden = false
	__ParticipantShape__00000020_.WidthWeight = 1.341690

	__ParticipantShape__00000022_.Name = `PP1 Very lon participant name PP1 Very lon participant name-Diagram P1`
	__ParticipantShape__00000022_.IsExpanded = false
	__ParticipantShape__00000022_.X = 241.968897
	__ParticipantShape__00000022_.Y = 181.495338
	__ParticipantShape__00000022_.Width = 250.000000
	__ParticipantShape__00000022_.Height = 70.000000
	__ParticipantShape__00000022_.IsHidden = false
	__ParticipantShape__00000022_.WidthWeight = 1.647398

	__ParticipantShape__00000023_.Name = `-DiagramProcess`
	__ParticipantShape__00000023_.IsExpanded = false
	__ParticipantShape__00000023_.X = 152.749068
	__ParticipantShape__00000023_.Y = 117.020557
	__ParticipantShape__00000023_.Width = 250.000000
	__ParticipantShape__00000023_.Height = 70.000000
	__ParticipantShape__00000023_.IsHidden = false
	__ParticipantShape__00000023_.WidthWeight = 1.584615

	__ParticipantShape__00000024_.Name = `P2.P2-DiagramProcess`
	__ParticipantShape__00000024_.IsExpanded = false
	__ParticipantShape__00000024_.X = 102.892796
	__ParticipantShape__00000024_.Y = 144.667849
	__ParticipantShape__00000024_.Width = 250.000000
	__ParticipantShape__00000024_.Height = 70.000000
	__ParticipantShape__00000024_.IsHidden = false
	__ParticipantShape__00000024_.WidthWeight = 1.000000

	__ParticipantShape__00000025_.Name = `P2.P1-D2`
	__ParticipantShape__00000025_.IsExpanded = false
	__ParticipantShape__00000025_.X = 125.725413
	__ParticipantShape__00000025_.Y = 134.043835
	__ParticipantShape__00000025_.Width = 250.000000
	__ParticipantShape__00000025_.Height = 70.000000
	__ParticipantShape__00000025_.IsHidden = false
	__ParticipantShape__00000025_.WidthWeight = 1.000000

	__Process__00000005_.Name = `P1`
	__Process__00000005_.Description = ``
	__Process__00000005_.ComputedPrefix = ``
	__Process__00000005_.SVG_Path = `M296-270q-42 35-87.5 32T129-269q-34-28-46.5-73.5T99-436l75-124q-25-22-39.5-53T120-680q0-66 47-113t113-47q66 0 113 47t47 113q0 66-47 113t-113 47q-9 0-18-1t-17-3l-77 130q-11 18-7 35.5t17 28.5q13 11 31 12.5t35-12.5l420-361q42-35 88-31.5t80 31.5q34 28 46 73.5T861-524l-75 124q25 22 39.5 53t14.5 67q0 66-47 113t-113 47q-66 0-113-47t-47-113q0-66 47-113t113-47q9 0 17.5 1t16.5 3l78-130q11-18 7-35.5T782-630q-13-11-31-12.5T716-630L296-270Zm40.5-353.5Q360-647 360-680t-23.5-56.5Q313-760 280-760t-56.5 23.5Q200-713 200-680t23.5 56.5Q247-600 280-600t56.5-23.5Zm400 400Q760-247 760-280t-23.5-56.5Q713-360 680-360t-56.5 23.5Q600-313 600-280t23.5 56.5Q647-200 680-200t56.5-23.5ZM280-680Zm400 400Z`
	__Process__00000005_.InverseAppliedScaling = 60.000000
	__Process__00000005_.IsSubProcessNodeExpanded = false
	__Process__00000005_.IsDataFlowsNodeExpanded = true

	__Process__00000006_.Name = `P2`
	__Process__00000006_.Description = ``
	__Process__00000006_.ComputedPrefix = ``
	__Process__00000006_.SVG_Path = `M296-270q-42 35-87.5 32T129-269q-34-28-46.5-73.5T99-436l75-124q-25-22-39.5-53T120-680q0-66 47-113t113-47q66 0 113 47t47 113q0 66-47 113t-113 47q-9 0-18-1t-17-3l-77 130q-11 18-7 35.5t17 28.5q13 11 31 12.5t35-12.5l420-361q42-35 88-31.5t80 31.5q34 28 46 73.5T861-524l-75 124q25 22 39.5 53t14.5 67q0 66-47 113t-113 47q-66 0-113-47t-47-113q0-66 47-113t113-47q9 0 17.5 1t16.5 3l78-130q11-18 7-35.5T782-630q-13-11-31-12.5T716-630L296-270Zm40.5-353.5Q360-647 360-680t-23.5-56.5Q313-760 280-760t-56.5 23.5Q200-713 200-680t23.5 56.5Q247-600 280-600t56.5-23.5Zm400 400Q760-247 760-280t-23.5-56.5Q713-360 680-360t-56.5 23.5Q600-313 600-280t23.5 56.5Q647-200 680-200t56.5-23.5ZM280-680Zm400 400Z`
	__Process__00000006_.InverseAppliedScaling = 60.000000
	__Process__00000006_.IsSubProcessNodeExpanded = false
	__Process__00000006_.IsDataFlowsNodeExpanded = true

	__ProcessShape__00000006_.Name = `ProcessShape`
	__ProcessShape__00000006_.IsExpanded = false
	__ProcessShape__00000006_.X = 93.000000
	__ProcessShape__00000006_.Y = 49.000000
	__ProcessShape__00000006_.Width = 860.000000
	__ProcessShape__00000006_.Height = 1000.000000
	__ProcessShape__00000006_.IsHidden = false

	__ProcessShape__00000007_.Name = `ProcessShape`
	__ProcessShape__00000007_.IsExpanded = false
	__ProcessShape__00000007_.X = 332.015236
	__ProcessShape__00000007_.Y = 180.000030
	__ProcessShape__00000007_.Width = 761.000000
	__ProcessShape__00000007_.Height = 845.999939
	__ProcessShape__00000007_.IsHidden = false

	__ProcessShape__00000009_.Name = `ProcessShape`
	__ProcessShape__00000009_.IsExpanded = false
	__ProcessShape__00000009_.X = 100.000000
	__ProcessShape__00000009_.Y = 50.000000
	__ProcessShape__00000009_.Width = 500.000000
	__ProcessShape__00000009_.Height = 1000.000000
	__ProcessShape__00000009_.IsHidden = false

	__Resource__00000000_.Name = `R1 with a long name long long long`
	__Resource__00000000_.Description = ``
	__Resource__00000000_.ComputedPrefix = ``
	__Resource__00000000_.SVG_Path = `m280-40 112-564-72 28v136h-80v-188l202-86q14-6 29.5-7t29.5 4q14 5 26.5 14t20.5 23l40 64q26 42 70.5 69T760-520v80q-70 0-125-29t-94-74l-25 123 84 80v300h-80v-260l-84-64-72 324h-84Zm203.5-723.5Q460-787 460-820t23.5-56.5Q507-900 540-900t56.5 23.5Q620-853 620-820t-23.5 56.5Q573-740 540-740t-56.5-23.5Z`
	__Resource__00000000_.InverseAppliedScaling = 60.000000

	__Resource__00000001_.Name = `R2aaaaa`
	__Resource__00000001_.Description = ``
	__Resource__00000001_.ComputedPrefix = ``
	__Resource__00000001_.SVG_Path = `M240-200v40q0 17-11.5 28.5T200-120h-40q-17 0-28.5-11.5T120-160v-320l84-240q6-18 21.5-29t34.5-11h440q19 0 34.5 11t21.5 29l84 240v320q0 17-11.5 28.5T800-120h-40q-17 0-28.5-11.5T720-160v-40H240Zm-8-360h496l-42-120H274l-42 120Zm-32 80v200-200Zm100 160q25 0 42.5-17.5T360-380q0-25-17.5-42.5T300-440q-25 0-42.5 17.5T240-380q0 25 17.5 42.5T300-320Zm360 0q25 0 42.5-17.5T720-380q0-25-17.5-42.5T660-440q-25 0-42.5 17.5T600-380q0 25 17.5 42.5T660-320Zm-460 40h560v-200H200v200Z`
	__Resource__00000001_.InverseAppliedScaling = 60.000000

	__Task__00000002_.Name = `P1`
	__Task__00000002_.Description = ``
	__Task__00000002_.ComputedPrefix = ``
	__Task__00000002_.IsStartTask = false
	__Task__00000002_.IsEndTask = false
	__Task__00000002_.IsTaskNameNotProcessName = false

	__Task__00000003_.Name = `PP2 T1`
	__Task__00000003_.Description = ``
	__Task__00000003_.ComputedPrefix = ``
	__Task__00000003_.IsStartTask = true
	__Task__00000003_.IsEndTask = false
	__Task__00000003_.IsTaskNameNotProcessName = false

	__Task__00000004_.Name = `Start`
	__Task__00000004_.Description = ``
	__Task__00000004_.ComputedPrefix = ``
	__Task__00000004_.IsStartTask = true
	__Task__00000004_.IsEndTask = false
	__Task__00000004_.IsTaskNameNotProcessName = false

	__Task__00000005_.Name = `End`
	__Task__00000005_.Description = ``
	__Task__00000005_.ComputedPrefix = ``
	__Task__00000005_.IsStartTask = false
	__Task__00000005_.IsEndTask = true
	__Task__00000005_.IsTaskNameNotProcessName = false

	__Task__00000006_.Name = `T4`
	__Task__00000006_.Description = ``
	__Task__00000006_.ComputedPrefix = ``
	__Task__00000006_.IsStartTask = false
	__Task__00000006_.IsEndTask = false
	__Task__00000006_.IsTaskNameNotProcessName = false

	__Task__00000007_.Name = `PP2 T2`
	__Task__00000007_.Description = ``
	__Task__00000007_.ComputedPrefix = ``
	__Task__00000007_.IsStartTask = false
	__Task__00000007_.IsEndTask = false
	__Task__00000007_.IsTaskNameNotProcessName = false

	__Task__00000008_.Name = `Start`
	__Task__00000008_.Description = ``
	__Task__00000008_.ComputedPrefix = ``
	__Task__00000008_.IsStartTask = true
	__Task__00000008_.IsEndTask = false
	__Task__00000008_.IsTaskNameNotProcessName = false

	__Task__00000009_.Name = `T1`
	__Task__00000009_.Description = ``
	__Task__00000009_.ComputedPrefix = ``
	__Task__00000009_.IsStartTask = false
	__Task__00000009_.IsEndTask = false
	__Task__00000009_.IsTaskNameNotProcessName = false

	__Task__00000010_.Name = `End`
	__Task__00000010_.Description = ``
	__Task__00000010_.ComputedPrefix = ``
	__Task__00000010_.IsStartTask = false
	__Task__00000010_.IsEndTask = true
	__Task__00000010_.IsTaskNameNotProcessName = false

	__Task__00000011_.Name = `P2.P2.Start`
	__Task__00000011_.Description = ``
	__Task__00000011_.ComputedPrefix = ``
	__Task__00000011_.IsStartTask = true
	__Task__00000011_.IsEndTask = false
	__Task__00000011_.IsTaskNameNotProcessName = false

	__Task__00000012_.Name = `P2.P2.T1`
	__Task__00000012_.Description = ``
	__Task__00000012_.ComputedPrefix = ``
	__Task__00000012_.IsStartTask = false
	__Task__00000012_.IsEndTask = false
	__Task__00000012_.IsTaskNameNotProcessName = false

	__TaskShape__00000011_.Name = `Start-DP2`
	__TaskShape__00000011_.IsExpanded = false
	__TaskShape__00000011_.X = 522.000000
	__TaskShape__00000011_.Y = 449.000076
	__TaskShape__00000011_.Width = 54.000000
	__TaskShape__00000011_.Height = 20.000000
	__TaskShape__00000011_.IsHidden = false

	__TaskShape__00000012_.Name = `T4-DP2`
	__TaskShape__00000012_.IsExpanded = false
	__TaskShape__00000012_.X = 480.030525
	__TaskShape__00000012_.Y = 825.000076
	__TaskShape__00000012_.Width = 167.000000
	__TaskShape__00000012_.Height = 70.000000
	__TaskShape__00000012_.IsHidden = false

	__TaskShape__00000013_.Name = `T2-DP2`
	__TaskShape__00000013_.IsExpanded = false
	__TaskShape__00000013_.X = 473.030525
	__TaskShape__00000013_.Y = 538.000076
	__TaskShape__00000013_.Width = 174.000000
	__TaskShape__00000013_.Height = 240.000000
	__TaskShape__00000013_.IsHidden = false

	__TaskShape__00000014_.Name = `PP2 T2-DP2`
	__TaskShape__00000014_.IsExpanded = false
	__TaskShape__00000014_.X = 731.536446
	__TaskShape__00000014_.Y = 596.000015
	__TaskShape__00000014_.Width = 164.722046
	__TaskShape__00000014_.Height = 70.000000
	__TaskShape__00000014_.IsHidden = false

	__TaskShape__00000015_.Name = `PP2 T1-Diagram P1`
	__TaskShape__00000015_.IsExpanded = false
	__TaskShape__00000015_.X = 821.258492
	__TaskShape__00000015_.Y = 414.000030
	__TaskShape__00000015_.Width = 75.000000
	__TaskShape__00000015_.Height = 20.000000
	__TaskShape__00000015_.IsHidden = false

	__TaskShape__00000016_.Name = `End-Diagram P1`
	__TaskShape__00000016_.IsExpanded = false
	__TaskShape__00000016_.X = 552.000000
	__TaskShape__00000016_.Y = 978.999969
	__TaskShape__00000016_.Width = 89.000000
	__TaskShape__00000016_.Height = 36.000000
	__TaskShape__00000016_.IsHidden = false

	__TaskShape__00000019_.Name = `-DiagramProcess`
	__TaskShape__00000019_.IsExpanded = false
	__TaskShape__00000019_.X = 164.000000
	__TaskShape__00000019_.Y = 332.672456
	__TaskShape__00000019_.Width = 129.000000
	__TaskShape__00000019_.Height = 70.000000
	__TaskShape__00000019_.IsHidden = false

	__TaskShape__00000020_.Name = `Start-DiagramProcess`
	__TaskShape__00000020_.IsExpanded = false
	__TaskShape__00000020_.X = 187.000000
	__TaskShape__00000020_.Y = 225.222502
	__TaskShape__00000020_.Width = 112.000000
	__TaskShape__00000020_.Height = 20.000000
	__TaskShape__00000020_.IsHidden = false

	__TaskShape__00000021_.Name = `-DiagramProcess`
	__TaskShape__00000021_.IsExpanded = false
	__TaskShape__00000021_.X = 153.000000
	__TaskShape__00000021_.Y = 506.000015
	__TaskShape__00000021_.Width = 126.000000
	__TaskShape__00000021_.Height = 70.000000
	__TaskShape__00000021_.IsHidden = false

	__TaskShape__00000022_.Name = `-DiagramProcess`
	__TaskShape__00000022_.IsExpanded = false
	__TaskShape__00000022_.X = 720.000000
	__TaskShape__00000022_.Y = 216.452871
	__TaskShape__00000022_.Width = 116.000000
	__TaskShape__00000022_.Height = 20.000000
	__TaskShape__00000022_.IsHidden = false

	__TaskShape__00000023_.Name = `-DiagramProcess`
	__TaskShape__00000023_.IsExpanded = false
	__TaskShape__00000023_.X = 738.000000
	__TaskShape__00000023_.Y = 269.238277
	__TaskShape__00000023_.Width = 127.000000
	__TaskShape__00000023_.Height = 178.000000
	__TaskShape__00000023_.IsHidden = false

	// insertion point for setup of pointers
	__AllocatedProcessShape__00000000_.Participant = __Participant__00000007_
	__AllocatedProcessShape__00000000_.Process = __Process__00000005_
	__AllocatedProcessShape__00000001_.Participant = __Participant__00000007_
	__AllocatedProcessShape__00000001_.Process = __Process__00000006_
	__AllocatedProcessShape__00000002_.Participant = __Participant__00000011_
	__AllocatedProcessShape__00000002_.Process = __Process__00000005_
	__AllocatedResourceShape__00000010_.Participant = __Participant__00000007_
	__AllocatedResourceShape__00000010_.Resource = __Resource__00000001_
	__AllocatedResourceShape__00000012_.Participant = __Participant__00000007_
	__AllocatedResourceShape__00000012_.Resource = __Resource__00000000_
	__AllocatedResourceShape__00000014_.Participant = __Participant__00000003_
	__AllocatedResourceShape__00000014_.Resource = __Resource__00000001_
	__AllocatedResourceShape__00000015_.Participant = __Participant__00000003_
	__AllocatedResourceShape__00000015_.Resource = __Resource__00000000_
	__AllocatedResourceShape__00000016_.Participant = __Participant__00000008_
	__AllocatedResourceShape__00000016_.Resource = __Resource__00000001_
	__AllocatedResourceShape__00000017_.Participant = __Participant__00000011_
	__AllocatedResourceShape__00000017_.Resource = __Resource__00000001_
	__ControlFlow__00000003_.Start = __Task__00000004_
	__ControlFlow__00000003_.End = __Task__00000002_
	__ControlFlow__00000004_.Start = __Task__00000002_
	__ControlFlow__00000004_.End = __Task__00000006_
	__ControlFlow__00000005_.Start = __Task__00000003_
	__ControlFlow__00000005_.End = __Task__00000007_
	__ControlFlow__00000006_.Start = __Task__00000006_
	__ControlFlow__00000006_.End = __Task__00000005_
	__ControlFlow__00000007_.Start = __Task__00000008_
	__ControlFlow__00000007_.End = __Task__00000009_
	__ControlFlow__00000008_.Start = __Task__00000009_
	__ControlFlow__00000008_.End = __Task__00000010_
	__ControlFlow__00000009_.Start = __Task__00000011_
	__ControlFlow__00000009_.End = __Task__00000012_
	__ControlFlowShape__00000000_.ControlFlow = __ControlFlow__00000003_
	__ControlFlowShape__00000001_.ControlFlow = __ControlFlow__00000005_
	__ControlFlowShape__00000002_.ControlFlow = __ControlFlow__00000004_
	__ControlFlowShape__00000003_.ControlFlow = __ControlFlow__00000006_
	__ControlFlowShape__00000004_.ControlFlow = __ControlFlow__00000007_
	__ControlFlowShape__00000005_.ControlFlow = __ControlFlow__00000008_
	__ControlFlowShape__00000006_.ControlFlow = __ControlFlow__00000009_
	__DataFlow__00000000_.Datas = append(__DataFlow__00000000_.Datas, __Data__00000000_)
	__DataFlow__00000000_.StartTask = __Task__00000002_
	__DataFlow__00000000_.EndTask = __Task__00000007_
	__DataFlow__00000000_.StartExternalParticipant = nil
	__DataFlow__00000000_.EndExternalParticipant = nil
	__DataFlow__00000002_.Datas = append(__DataFlow__00000002_.Datas, __Data__00000003_)
	__DataFlow__00000002_.StartTask = __Task__00000009_
	__DataFlow__00000002_.EndTask = __Task__00000012_
	__DataFlow__00000002_.StartExternalParticipant = nil
	__DataFlow__00000002_.EndExternalParticipant = nil
	__DataFlow__00000003_.Datas = append(__DataFlow__00000003_.Datas, __Data__00000001_)
	__DataFlow__00000003_.Datas = append(__DataFlow__00000003_.Datas, __Data__00000000_)
	__DataFlow__00000003_.StartTask = nil
	__DataFlow__00000003_.EndTask = __Task__00000002_
	__DataFlow__00000003_.StartExternalParticipant = __Participant__00000007_
	__DataFlow__00000003_.EndExternalParticipant = nil
	__DataFlow__00000005_.StartTask = nil
	__DataFlow__00000005_.EndTask = __Task__00000007_
	__DataFlow__00000005_.StartExternalParticipant = __Participant__00000011_
	__DataFlow__00000005_.EndExternalParticipant = nil
	__DataFlowShape__00000004_.DataFlow = __DataFlow__00000000_
	__DataFlowShape__00000005_.DataFlow = __DataFlow__00000002_
	__DataFlowShape__00000006_.DataFlow = __DataFlow__00000003_
	__DataFlowShape__00000008_.DataFlow = __DataFlow__00000005_
	__DataShape__00000011_.Data = __Data__00000000_
	__DataShape__00000011_.DataFlow = nil
	__DataShape__00000012_.Data = __Data__00000001_
	__DataShape__00000012_.DataFlow = nil
	__DataShape__00000013_.Data = __Data__00000001_
	__DataShape__00000013_.DataFlow = nil
	__DataShape__00000014_.Data = __Data__00000000_
	__DataShape__00000014_.DataFlow = nil
	__DataShape__00000016_.Data = __Data__00000001_
	__DataShape__00000016_.DataFlow = nil
	__DataShape__00000017_.Data = __Data__00000000_
	__DataShape__00000017_.DataFlow = nil
	__DataShape__00000019_.Data = __Data__00000000_
	__DataShape__00000019_.DataFlow = __DataFlow__00000000_
	__DataShape__00000022_.Data = __Data__00000001_
	__DataShape__00000022_.DataFlow = nil
	__DataShape__00000024_.Data = __Data__00000000_
	__DataShape__00000024_.DataFlow = nil
	__DataShape__00000025_.Data = __Data__00000003_
	__DataShape__00000025_.DataFlow = __DataFlow__00000002_
	__DataShape__00000026_.Data = __Data__00000001_
	__DataShape__00000026_.DataFlow = __DataFlow__00000003_
	__DataShape__00000027_.Data = __Data__00000000_
	__DataShape__00000027_.DataFlow = __DataFlow__00000003_
	__DiagramProcess__00000006_.Process_Shapes = append(__DiagramProcess__00000006_.Process_Shapes, __ProcessShape__00000006_)
	__DiagramProcess__00000006_.Participant_Shapes = append(__DiagramProcess__00000006_.Participant_Shapes, __ParticipantShape__00000023_)
	__DiagramProcess__00000006_.Participant_Shapes = append(__DiagramProcess__00000006_.Participant_Shapes, __ParticipantShape__00000024_)
	__DiagramProcess__00000006_.ParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000006_.ParticipantWhoseNodeIsExpanded, __Participant__00000008_)
	__DiagramProcess__00000006_.ParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000006_.ParticipantWhoseNodeIsExpanded, __Participant__00000009_)
	__DiagramProcess__00000006_.Task_Shapes = append(__DiagramProcess__00000006_.Task_Shapes, __TaskShape__00000019_)
	__DiagramProcess__00000006_.Task_Shapes = append(__DiagramProcess__00000006_.Task_Shapes, __TaskShape__00000020_)
	__DiagramProcess__00000006_.Task_Shapes = append(__DiagramProcess__00000006_.Task_Shapes, __TaskShape__00000021_)
	__DiagramProcess__00000006_.Task_Shapes = append(__DiagramProcess__00000006_.Task_Shapes, __TaskShape__00000022_)
	__DiagramProcess__00000006_.Task_Shapes = append(__DiagramProcess__00000006_.Task_Shapes, __TaskShape__00000023_)
	__DiagramProcess__00000006_.ControlFlow_Shapes = append(__DiagramProcess__00000006_.ControlFlow_Shapes, __ControlFlowShape__00000004_)
	__DiagramProcess__00000006_.ControlFlow_Shapes = append(__DiagramProcess__00000006_.ControlFlow_Shapes, __ControlFlowShape__00000005_)
	__DiagramProcess__00000006_.ControlFlow_Shapes = append(__DiagramProcess__00000006_.ControlFlow_Shapes, __ControlFlowShape__00000006_)
	__DiagramProcess__00000006_.DataFlowsWhoseNodeIsExpanded = append(__DiagramProcess__00000006_.DataFlowsWhoseNodeIsExpanded, __DataFlow__00000002_)
	__DiagramProcess__00000006_.DataFlow_Shapes = append(__DiagramProcess__00000006_.DataFlow_Shapes, __DataFlowShape__00000005_)
	__DiagramProcess__00000006_.Data_Shapes = append(__DiagramProcess__00000006_.Data_Shapes, __DataShape__00000025_)
	__DiagramProcess__00000007_.Process_Shapes = append(__DiagramProcess__00000007_.Process_Shapes, __ProcessShape__00000007_)
	__DiagramProcess__00000007_.Participant_Shapes = append(__DiagramProcess__00000007_.Participant_Shapes, __ParticipantShape__00000022_)
	__DiagramProcess__00000007_.Participant_Shapes = append(__DiagramProcess__00000007_.Participant_Shapes, __ParticipantShape__00000020_)
	__DiagramProcess__00000007_.Participant_Shapes = append(__DiagramProcess__00000007_.Participant_Shapes, __ParticipantShape__00000019_)
	__DiagramProcess__00000007_.ParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000007_.ParticipantWhoseNodeIsExpanded, __Participant__00000003_)
	__DiagramProcess__00000007_.ExternalParticipant_Shapes = append(__DiagramProcess__00000007_.ExternalParticipant_Shapes, __ExternalParticipantShape__00000003_)
	__DiagramProcess__00000007_.ExternalParticipant_Shapes = append(__DiagramProcess__00000007_.ExternalParticipant_Shapes, __ExternalParticipantShape__00000005_)
	__DiagramProcess__00000007_.ExternalParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000007_.ExternalParticipantWhoseNodeIsExpanded, __Participant__00000003_)
	__DiagramProcess__00000007_.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded = append(__DiagramProcess__00000007_.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded, __Participant__00000007_)
	__DiagramProcess__00000007_.TasksWhoseNodeIsExpanded = append(__DiagramProcess__00000007_.TasksWhoseNodeIsExpanded, __Task__00000002_)
	__DiagramProcess__00000007_.Task_Shapes = append(__DiagramProcess__00000007_.Task_Shapes, __TaskShape__00000011_)
	__DiagramProcess__00000007_.Task_Shapes = append(__DiagramProcess__00000007_.Task_Shapes, __TaskShape__00000012_)
	__DiagramProcess__00000007_.Task_Shapes = append(__DiagramProcess__00000007_.Task_Shapes, __TaskShape__00000013_)
	__DiagramProcess__00000007_.Task_Shapes = append(__DiagramProcess__00000007_.Task_Shapes, __TaskShape__00000014_)
	__DiagramProcess__00000007_.Task_Shapes = append(__DiagramProcess__00000007_.Task_Shapes, __TaskShape__00000015_)
	__DiagramProcess__00000007_.Task_Shapes = append(__DiagramProcess__00000007_.Task_Shapes, __TaskShape__00000016_)
	__DiagramProcess__00000007_.ControlFlow_Shapes = append(__DiagramProcess__00000007_.ControlFlow_Shapes, __ControlFlowShape__00000000_)
	__DiagramProcess__00000007_.ControlFlow_Shapes = append(__DiagramProcess__00000007_.ControlFlow_Shapes, __ControlFlowShape__00000001_)
	__DiagramProcess__00000007_.ControlFlow_Shapes = append(__DiagramProcess__00000007_.ControlFlow_Shapes, __ControlFlowShape__00000002_)
	__DiagramProcess__00000007_.ControlFlow_Shapes = append(__DiagramProcess__00000007_.ControlFlow_Shapes, __ControlFlowShape__00000003_)
	__DiagramProcess__00000007_.DataFlowsWhoseNodeIsExpanded = append(__DiagramProcess__00000007_.DataFlowsWhoseNodeIsExpanded, __DataFlow__00000000_)
	__DiagramProcess__00000007_.DataFlow_Shapes = append(__DiagramProcess__00000007_.DataFlow_Shapes, __DataFlowShape__00000004_)
	__DiagramProcess__00000007_.DataFlow_Shapes = append(__DiagramProcess__00000007_.DataFlow_Shapes, __DataFlowShape__00000006_)
	__DiagramProcess__00000007_.DataFlow_Shapes = append(__DiagramProcess__00000007_.DataFlow_Shapes, __DataFlowShape__00000008_)
	__DiagramProcess__00000007_.Data_Shapes = append(__DiagramProcess__00000007_.Data_Shapes, __DataShape__00000011_)
	__DiagramProcess__00000007_.Data_Shapes = append(__DiagramProcess__00000007_.Data_Shapes, __DataShape__00000012_)
	__DiagramProcess__00000007_.Data_Shapes = append(__DiagramProcess__00000007_.Data_Shapes, __DataShape__00000013_)
	__DiagramProcess__00000007_.Data_Shapes = append(__DiagramProcess__00000007_.Data_Shapes, __DataShape__00000014_)
	__DiagramProcess__00000007_.Data_Shapes = append(__DiagramProcess__00000007_.Data_Shapes, __DataShape__00000016_)
	__DiagramProcess__00000007_.Data_Shapes = append(__DiagramProcess__00000007_.Data_Shapes, __DataShape__00000017_)
	__DiagramProcess__00000007_.Data_Shapes = append(__DiagramProcess__00000007_.Data_Shapes, __DataShape__00000019_)
	__DiagramProcess__00000007_.Data_Shapes = append(__DiagramProcess__00000007_.Data_Shapes, __DataShape__00000022_)
	__DiagramProcess__00000007_.Data_Shapes = append(__DiagramProcess__00000007_.Data_Shapes, __DataShape__00000024_)
	__DiagramProcess__00000007_.Data_Shapes = append(__DiagramProcess__00000007_.Data_Shapes, __DataShape__00000026_)
	__DiagramProcess__00000007_.Data_Shapes = append(__DiagramProcess__00000007_.Data_Shapes, __DataShape__00000027_)
	__DiagramProcess__00000007_.DataFlowsWhoseDataNodeIsExpanded = append(__DiagramProcess__00000007_.DataFlowsWhoseDataNodeIsExpanded, __DataFlow__00000000_)
	__DiagramProcess__00000007_.DataFlowsWhoseDataNodeIsExpanded = append(__DiagramProcess__00000007_.DataFlowsWhoseDataNodeIsExpanded, __DataFlow__00000003_)
	__DiagramProcess__00000007_.AllocatedResourceShapes = append(__DiagramProcess__00000007_.AllocatedResourceShapes, __AllocatedResourceShape__00000010_)
	__DiagramProcess__00000007_.AllocatedResourceShapes = append(__DiagramProcess__00000007_.AllocatedResourceShapes, __AllocatedResourceShape__00000012_)
	__DiagramProcess__00000007_.AllocatedResourceShapes = append(__DiagramProcess__00000007_.AllocatedResourceShapes, __AllocatedResourceShape__00000014_)
	__DiagramProcess__00000007_.AllocatedResourceShapes = append(__DiagramProcess__00000007_.AllocatedResourceShapes, __AllocatedResourceShape__00000015_)
	__DiagramProcess__00000007_.AllocatedResourceShapes = append(__DiagramProcess__00000007_.AllocatedResourceShapes, __AllocatedResourceShape__00000017_)
	__DiagramProcess__00000007_.AllocatedProcessShapes = append(__DiagramProcess__00000007_.AllocatedProcessShapes, __AllocatedProcessShape__00000000_)
	__DiagramProcess__00000007_.AllocatedProcessShapes = append(__DiagramProcess__00000007_.AllocatedProcessShapes, __AllocatedProcessShape__00000001_)
	__DiagramProcess__00000007_.AllocatedProcessShapes = append(__DiagramProcess__00000007_.AllocatedProcessShapes, __AllocatedProcessShape__00000002_)
	__DiagramProcess__00000007_.Note_Shapes = append(__DiagramProcess__00000007_.Note_Shapes, __NoteShape__00000000_)
	__DiagramProcess__00000007_.NotesWhoseNodeIsExpanded = append(__DiagramProcess__00000007_.NotesWhoseNodeIsExpanded, __Note__00000000_)
	__DiagramProcess__00000007_.NoteTaskShapes = append(__DiagramProcess__00000007_.NoteTaskShapes, __NoteTaskShape__00000000_)
	__DiagramProcess__00000009_.Process_Shapes = append(__DiagramProcess__00000009_.Process_Shapes, __ProcessShape__00000009_)
	__DiagramProcess__00000009_.Participant_Shapes = append(__DiagramProcess__00000009_.Participant_Shapes, __ParticipantShape__00000025_)
	__DiagramProcess__00000009_.ParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000009_.ParticipantWhoseNodeIsExpanded, __Participant__00000008_)
	__DiagramProcess__00000009_.AllocatedResourceShapes = append(__DiagramProcess__00000009_.AllocatedResourceShapes, __AllocatedResourceShape__00000016_)
	__ExternalParticipantShape__00000003_.Participant = __Participant__00000007_
	__ExternalParticipantShape__00000005_.Participant = __Participant__00000011_
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000005_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000006_)
	__Library__00000000_.ProcesssWhoseNodeIsExpanded = append(__Library__00000000_.ProcesssWhoseNodeIsExpanded, __Process__00000006_)
	__Library__00000000_.ProcesssWhoseNodeIsExpanded = append(__Library__00000000_.ProcesssWhoseNodeIsExpanded, __Process__00000005_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000000_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000002_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000003_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000005_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000000_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000001_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000003_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000000_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000001_)
	__Library__00000000_.RootNotes = append(__Library__00000000_.RootNotes, __Note__00000000_)
	__Library__00000000_.NotesWhoseNodeIsExpanded = append(__Library__00000000_.NotesWhoseNodeIsExpanded, __Note__00000000_)
	__Note__00000000_.Tasks = append(__Note__00000000_.Tasks, __Task__00000002_)
	__NoteShape__00000000_.Note = __Note__00000000_
	__NoteTaskShape__00000000_.Note = __Note__00000000_
	__NoteTaskShape__00000000_.Task = __Task__00000002_
	__Participant__00000003_.Resources = append(__Participant__00000003_.Resources, __Resource__00000000_)
	__Participant__00000003_.Resources = append(__Participant__00000003_.Resources, __Resource__00000001_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000002_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000004_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000005_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000006_)
	__Participant__00000003_.ControlFlows = append(__Participant__00000003_.ControlFlows, __ControlFlow__00000003_)
	__Participant__00000003_.ControlFlows = append(__Participant__00000003_.ControlFlows, __ControlFlow__00000004_)
	__Participant__00000003_.ControlFlows = append(__Participant__00000003_.ControlFlows, __ControlFlow__00000006_)
	__Participant__00000003_.TaskWhoseOutControlFlowsNodeIsExpanded = append(__Participant__00000003_.TaskWhoseOutControlFlowsNodeIsExpanded, __Task__00000004_)
	__Participant__00000003_.TaskWhoseOutDataFlowsNodeIsExpanded = append(__Participant__00000003_.TaskWhoseOutDataFlowsNodeIsExpanded, __Task__00000002_)
	__Participant__00000003_.TaskWhoseInDataFlowsNodeIsExpanded = append(__Participant__00000003_.TaskWhoseInDataFlowsNodeIsExpanded, __Task__00000002_)
	__Participant__00000004_.Tasks = append(__Participant__00000004_.Tasks, __Task__00000003_)
	__Participant__00000004_.Tasks = append(__Participant__00000004_.Tasks, __Task__00000007_)
	__Participant__00000004_.ControlFlows = append(__Participant__00000004_.ControlFlows, __ControlFlow__00000005_)
	__Participant__00000004_.TaskWhoseOutControlFlowsNodeIsExpanded = append(__Participant__00000004_.TaskWhoseOutControlFlowsNodeIsExpanded, __Task__00000003_)
	__Participant__00000004_.TaskWhoseInDataFlowsNodeIsExpanded = append(__Participant__00000004_.TaskWhoseInDataFlowsNodeIsExpanded, __Task__00000007_)
	__Participant__00000007_.Resources = append(__Participant__00000007_.Resources, __Resource__00000000_)
	__Participant__00000007_.Resources = append(__Participant__00000007_.Resources, __Resource__00000001_)
	__Participant__00000007_.Processes = append(__Participant__00000007_.Processes, __Process__00000005_)
	__Participant__00000007_.Processes = append(__Participant__00000007_.Processes, __Process__00000006_)
	__Participant__00000008_.Resources = append(__Participant__00000008_.Resources, __Resource__00000001_)
	__Participant__00000008_.Tasks = append(__Participant__00000008_.Tasks, __Task__00000008_)
	__Participant__00000008_.Tasks = append(__Participant__00000008_.Tasks, __Task__00000009_)
	__Participant__00000008_.Tasks = append(__Participant__00000008_.Tasks, __Task__00000010_)
	__Participant__00000008_.ControlFlows = append(__Participant__00000008_.ControlFlows, __ControlFlow__00000007_)
	__Participant__00000008_.ControlFlows = append(__Participant__00000008_.ControlFlows, __ControlFlow__00000008_)
	__Participant__00000009_.Tasks = append(__Participant__00000009_.Tasks, __Task__00000011_)
	__Participant__00000009_.Tasks = append(__Participant__00000009_.Tasks, __Task__00000012_)
	__Participant__00000009_.ControlFlows = append(__Participant__00000009_.ControlFlows, __ControlFlow__00000009_)
	__Participant__00000011_.Resources = append(__Participant__00000011_.Resources, __Resource__00000001_)
	__Participant__00000011_.Processes = append(__Participant__00000011_.Processes, __Process__00000005_)
	__ParticipantShape__00000019_.Participant = __Participant__00000005_
	__ParticipantShape__00000020_.Participant = __Participant__00000004_
	__ParticipantShape__00000022_.Participant = __Participant__00000003_
	__ParticipantShape__00000023_.Participant = __Participant__00000008_
	__ParticipantShape__00000024_.Participant = __Participant__00000009_
	__ParticipantShape__00000025_.Participant = __Participant__00000008_
	__Process__00000005_.DiagramProcesss = append(__Process__00000005_.DiagramProcesss, __DiagramProcess__00000007_)
	__Process__00000005_.DiagramProcessWhoseNodeIsExpanded = append(__Process__00000005_.DiagramProcessWhoseNodeIsExpanded, __DiagramProcess__00000007_)
	__Process__00000005_.Participants = append(__Process__00000005_.Participants, __Participant__00000003_)
	__Process__00000005_.Participants = append(__Process__00000005_.Participants, __Participant__00000004_)
	__Process__00000005_.Participants = append(__Process__00000005_.Participants, __Participant__00000005_)
	__Process__00000005_.Participants = append(__Process__00000005_.Participants, __Participant__00000006_)
	__Process__00000005_.ParticipantWhoseNodeIsExpanded = append(__Process__00000005_.ParticipantWhoseNodeIsExpanded, __Participant__00000003_)
	__Process__00000005_.DataFlows = append(__Process__00000005_.DataFlows, __DataFlow__00000000_)
	__Process__00000005_.DataFlows = append(__Process__00000005_.DataFlows, __DataFlow__00000003_)
	__Process__00000005_.DataFlows = append(__Process__00000005_.DataFlows, __DataFlow__00000005_)
	__Process__00000005_.ExternalParticipants = append(__Process__00000005_.ExternalParticipants, __Participant__00000007_)
	__Process__00000005_.ExternalParticipants = append(__Process__00000005_.ExternalParticipants, __Participant__00000011_)
	__Process__00000006_.DiagramProcesss = append(__Process__00000006_.DiagramProcesss, __DiagramProcess__00000006_)
	__Process__00000006_.DiagramProcesss = append(__Process__00000006_.DiagramProcesss, __DiagramProcess__00000009_)
	__Process__00000006_.DiagramProcessWhoseNodeIsExpanded = append(__Process__00000006_.DiagramProcessWhoseNodeIsExpanded, __DiagramProcess__00000006_)
	__Process__00000006_.DiagramProcessWhoseNodeIsExpanded = append(__Process__00000006_.DiagramProcessWhoseNodeIsExpanded, __DiagramProcess__00000009_)
	__Process__00000006_.Participants = append(__Process__00000006_.Participants, __Participant__00000008_)
	__Process__00000006_.Participants = append(__Process__00000006_.Participants, __Participant__00000009_)
	__Process__00000006_.DataFlows = append(__Process__00000006_.DataFlows, __DataFlow__00000002_)
	__ProcessShape__00000006_.Process = __Process__00000006_
	__ProcessShape__00000007_.Process = __Process__00000005_
	__ProcessShape__00000009_.Process = __Process__00000006_
	__Task__00000002_.Type = __Process__00000005_
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
	__TaskShape__00000011_.Task = __Task__00000004_
	__TaskShape__00000012_.Task = __Task__00000006_
	__TaskShape__00000013_.Task = __Task__00000002_
	__TaskShape__00000014_.Task = __Task__00000007_
	__TaskShape__00000015_.Task = __Task__00000003_
	__TaskShape__00000016_.Task = __Task__00000005_
	__TaskShape__00000019_.Task = __Task__00000009_
	__TaskShape__00000020_.Task = __Task__00000008_
	__TaskShape__00000021_.Task = __Task__00000010_
	__TaskShape__00000022_.Task = __Task__00000011_
	__TaskShape__00000023_.Task = __Task__00000012_
}
