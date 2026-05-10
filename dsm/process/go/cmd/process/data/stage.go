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

	__AllocatedResourceShape__00000010_ := (&models.AllocatedResourceShape{Name: `Diagram P1-EP1 Very Long Participant Name-R2`}).Stage(stage)
	__AllocatedResourceShape__00000012_ := (&models.AllocatedResourceShape{Name: `Diagram P1-EP1 Very Long Participant Name-R1 with a long name long long long`}).Stage(stage)
	__AllocatedResourceShape__00000014_ := (&models.AllocatedResourceShape{Name: `Diagram P1-PP1 Very lon participant name PP1 Very lon participant name-R2aaaaa`}).Stage(stage)
	__AllocatedResourceShape__00000015_ := (&models.AllocatedResourceShape{Name: `Diagram P1-PP1 Very lon participant name PP1 Very lon participant name-R1 with a long name long long long`}).Stage(stage)

	__ControlFlow__00000003_ := (&models.ControlFlow{Name: `Start to T2`}).Stage(stage)
	__ControlFlow__00000004_ := (&models.ControlFlow{Name: `T2 to End`}).Stage(stage)
	__ControlFlow__00000005_ := (&models.ControlFlow{Name: `PP2 T1 to PP2 T2`}).Stage(stage)
	__ControlFlow__00000006_ := (&models.ControlFlow{Name: `T4 to End`}).Stage(stage)

	__ControlFlowShape__00000000_ := (&models.ControlFlowShape{Name: `Start to T2`}).Stage(stage)
	__ControlFlowShape__00000001_ := (&models.ControlFlowShape{Name: `PP2 T1 to PP2 T2`}).Stage(stage)
	__ControlFlowShape__00000002_ := (&models.ControlFlowShape{Name: `T2 to End`}).Stage(stage)
	__ControlFlowShape__00000003_ := (&models.ControlFlowShape{Name: `T4 to End`}).Stage(stage)

	__Data__00000000_ := (&models.Data{Name: `D1`}).Stage(stage)
	__Data__00000001_ := (&models.Data{Name: `D2`}).Stage(stage)
	__Data__00000002_ := (&models.Data{Name: `L1.D1`}).Stage(stage)

	__DataFlow__00000000_ := (&models.DataFlow{Name: `T2 to PP2 T2`}).Stage(stage)
	__DataFlow__00000001_ := (&models.DataFlow{Name: `EP1 to T2`}).Stage(stage)

	__DataFlowShape__00000003_ := (&models.DataFlowShape{Name: `EP1 to T2`}).Stage(stage)
	__DataFlowShape__00000004_ := (&models.DataFlowShape{Name: `T2 to PP2 T2`}).Stage(stage)

	__DataShape__00000011_ := (&models.DataShape{Name: `T2 to PP2 T2-T2 to PP2 T2-Diagram P1`}).Stage(stage)
	__DataShape__00000012_ := (&models.DataShape{Name: `T2 to PP2 T2-T2 to PP2 T2-Diagram P1`}).Stage(stage)
	__DataShape__00000013_ := (&models.DataShape{Name: `PP2 T2 to T2 2-PP2 T2 to T2-Diagram P1`}).Stage(stage)
	__DataShape__00000014_ := (&models.DataShape{Name: `PP2 T2 to T2 2-PP2 T2 to T2-Diagram P1`}).Stage(stage)
	__DataShape__00000016_ := (&models.DataShape{Name: `PP2 T2 to T2-PP2 T2 to T2-Diagram P1`}).Stage(stage)
	__DataShape__00000017_ := (&models.DataShape{Name: `PP2 T2 to T2-PP2 T2 to T2-Diagram P1`}).Stage(stage)
	__DataShape__00000019_ := (&models.DataShape{Name: `T2 to PP2 T2-T2 to PP2 T2-Diagram P1`}).Stage(stage)
	__DataShape__00000022_ := (&models.DataShape{Name: `EP1 to T2-EP1 to T2-Diagram P1`}).Stage(stage)
	__DataShape__00000024_ := (&models.DataShape{Name: `EP1 to T2-EP1 to T2-Diagram P1`}).Stage(stage)

	__DiagramProcess__00000006_ := (&models.DiagramProcess{Name: `DiagramProcess`}).Stage(stage)
	__DiagramProcess__00000007_ := (&models.DiagramProcess{Name: `Diagram P1`}).Stage(stage)

	__ExternalParticipantShape__00000000_ := (&models.ExternalParticipantShape{Name: `EP1-Diagram P1`}).Stage(stage)
	__ExternalParticipantShape__00000001_ := (&models.ExternalParticipantShape{Name: `EP1-Diagram P1`}).Stage(stage)
	__ExternalParticipantShape__00000002_ := (&models.ExternalParticipantShape{Name: `EP1-Diagram P1`}).Stage(stage)
	__ExternalParticipantShape__00000003_ := (&models.ExternalParticipantShape{Name: `EP1-Diagram P1`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `Root`}).Stage(stage)

	__Participant__00000003_ := (&models.Participant{Name: `PP1 Very lon participant name PP1 Very lon participant name`}).Stage(stage)
	__Participant__00000004_ := (&models.Participant{Name: `PP2`}).Stage(stage)
	__Participant__00000005_ := (&models.Participant{Name: `PP3`}).Stage(stage)
	__Participant__00000006_ := (&models.Participant{Name: `PP4`}).Stage(stage)
	__Participant__00000007_ := (&models.Participant{Name: `EP1 Very Long Participant Name`}).Stage(stage)

	__ParticipantShape__00000019_ := (&models.ParticipantShape{Name: `PP3-Diagram P1`}).Stage(stage)
	__ParticipantShape__00000020_ := (&models.ParticipantShape{Name: `PP2-Diagram P1`}).Stage(stage)
	__ParticipantShape__00000022_ := (&models.ParticipantShape{Name: `PP1 Very lon participant name PP1 Very lon participant name-Diagram P1`}).Stage(stage)

	__Process__00000005_ := (&models.Process{Name: `P1`}).Stage(stage)
	__Process__00000006_ := (&models.Process{Name: `P2`}).Stage(stage)

	__ProcessShape__00000006_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)
	__ProcessShape__00000007_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)

	__Resource__00000000_ := (&models.Resource{Name: `R1 with a long name long long long`}).Stage(stage)
	__Resource__00000001_ := (&models.Resource{Name: `R2aaaaa`}).Stage(stage)

	__Task__00000002_ := (&models.Task{Name: `T2`}).Stage(stage)
	__Task__00000003_ := (&models.Task{Name: `PP2 T1`}).Stage(stage)
	__Task__00000004_ := (&models.Task{Name: `Start`}).Stage(stage)
	__Task__00000005_ := (&models.Task{Name: `End`}).Stage(stage)
	__Task__00000006_ := (&models.Task{Name: `T4`}).Stage(stage)
	__Task__00000007_ := (&models.Task{Name: `PP2 T2`}).Stage(stage)

	__TaskShape__00000011_ := (&models.TaskShape{Name: `Start-DP2`}).Stage(stage)
	__TaskShape__00000012_ := (&models.TaskShape{Name: `T4-DP2`}).Stage(stage)
	__TaskShape__00000013_ := (&models.TaskShape{Name: `T2-DP2`}).Stage(stage)
	__TaskShape__00000014_ := (&models.TaskShape{Name: `PP2 T2-DP2`}).Stage(stage)
	__TaskShape__00000015_ := (&models.TaskShape{Name: `PP2 T1-Diagram P1`}).Stage(stage)
	__TaskShape__00000016_ := (&models.TaskShape{Name: `End-Diagram P1`}).Stage(stage)

	// insertion point for initialization of values

	__AllocatedResourceShape__00000010_.Name = `Diagram P1-EP1 Very Long Participant Name-R2`

	__AllocatedResourceShape__00000012_.Name = `Diagram P1-EP1 Very Long Participant Name-R1 with a long name long long long`

	__AllocatedResourceShape__00000014_.Name = `Diagram P1-PP1 Very lon participant name PP1 Very lon participant name-R2aaaaa`

	__AllocatedResourceShape__00000015_.Name = `Diagram P1-PP1 Very lon participant name PP1 Very lon participant name-R1 with a long name long long long`

	__ControlFlow__00000003_.Name = `Start to T2`
	__ControlFlow__00000003_.ComputedPrefix = ``

	__ControlFlow__00000004_.Name = `T2 to End`
	__ControlFlow__00000004_.ComputedPrefix = ``

	__ControlFlow__00000005_.Name = `PP2 T1 to PP2 T2`
	__ControlFlow__00000005_.ComputedPrefix = ``

	__ControlFlow__00000006_.Name = `T4 to End`
	__ControlFlow__00000006_.ComputedPrefix = ``

	__ControlFlowShape__00000000_.Name = `Start to T2`
	__ControlFlowShape__00000000_.StartRatio = 0.971920
	__ControlFlowShape__00000000_.EndRatio = 0.500000
	__ControlFlowShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000000_.CornerOffsetRatio = 2.278388
	__ControlFlowShape__00000000_.IsHidden = false

	__ControlFlowShape__00000001_.Name = `PP2 T1 to PP2 T2`
	__ControlFlowShape__00000001_.StartRatio = 1.000000
	__ControlFlowShape__00000001_.EndRatio = 0.500000
	__ControlFlowShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000001_.CornerOffsetRatio = 2.823598
	__ControlFlowShape__00000001_.IsHidden = false

	__ControlFlowShape__00000002_.Name = `T2 to End`
	__ControlFlowShape__00000002_.StartRatio = 0.500000
	__ControlFlowShape__00000002_.EndRatio = 0.789929
	__ControlFlowShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000002_.CornerOffsetRatio = 1.100000
	__ControlFlowShape__00000002_.IsHidden = false

	__ControlFlowShape__00000003_.Name = `T4 to End`
	__ControlFlowShape__00000003_.StartRatio = 0.500000
	__ControlFlowShape__00000003_.EndRatio = 0.793912
	__ControlFlowShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000003_.CornerOffsetRatio = 1.500000
	__ControlFlowShape__00000003_.IsHidden = false

	__Data__00000000_.Name = `D1`
	__Data__00000000_.ComputedPrefix = ``

	__Data__00000001_.Name = `D2`
	__Data__00000001_.ComputedPrefix = ``

	__Data__00000002_.Name = `L1.D1`
	__Data__00000002_.ComputedPrefix = ``

	__DataFlow__00000000_.Name = `T2 to PP2 T2`
	__DataFlow__00000000_.ComputedPrefix = ``
	__DataFlow__00000000_.Type = models.DataFlow_Task2Task
	__DataFlow__00000000_.IsDatasNodeExpanded = false

	__DataFlow__00000001_.Name = `EP1 to T2`
	__DataFlow__00000001_.ComputedPrefix = ``
	__DataFlow__00000001_.Type = models.DataFlow_ExternalParticipant2Task
	__DataFlow__00000001_.IsDatasNodeExpanded = false

	__DataFlowShape__00000003_.Name = `EP1 to T2`
	__DataFlowShape__00000003_.StartRatio = 0.471068
	__DataFlowShape__00000003_.EndRatio = 0.562565
	__DataFlowShape__00000003_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000003_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000003_.CornerOffsetRatio = 678.193331
	__DataFlowShape__00000003_.IsHidden = false

	__DataFlowShape__00000004_.Name = `T2 to PP2 T2`
	__DataFlowShape__00000004_.StartRatio = 0.820898
	__DataFlowShape__00000004_.EndRatio = 0.500000
	__DataFlowShape__00000004_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__DataFlowShape__00000004_.CornerOffsetRatio = 1.008329
	__DataFlowShape__00000004_.IsHidden = false

	__DataShape__00000011_.Name = `T2 to PP2 T2-T2 to PP2 T2-Diagram P1`

	__DataShape__00000012_.Name = `T2 to PP2 T2-T2 to PP2 T2-Diagram P1`

	__DataShape__00000013_.Name = `PP2 T2 to T2 2-PP2 T2 to T2-Diagram P1`

	__DataShape__00000014_.Name = `PP2 T2 to T2 2-PP2 T2 to T2-Diagram P1`

	__DataShape__00000016_.Name = `PP2 T2 to T2-PP2 T2 to T2-Diagram P1`

	__DataShape__00000017_.Name = `PP2 T2 to T2-PP2 T2 to T2-Diagram P1`

	__DataShape__00000019_.Name = `T2 to PP2 T2-T2 to PP2 T2-Diagram P1`

	__DataShape__00000022_.Name = `EP1 to T2-EP1 to T2-Diagram P1`

	__DataShape__00000024_.Name = `EP1 to T2-EP1 to T2-Diagram P1`

	__DiagramProcess__00000006_.Name = `DiagramProcess`
	__DiagramProcess__00000006_.ComputedPrefix = ``
	__DiagramProcess__00000006_.IsChecked = false
	__DiagramProcess__00000006_.IsEditable_ = false
	__DiagramProcess__00000006_.IsShowPrefix = false
	__DiagramProcess__00000006_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000006_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000006_.Width = 900.000000
	__DiagramProcess__00000006_.Height = 1350.000000
	__DiagramProcess__00000006_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000006_.IsParticipantsNodeExpanded = false
	__DiagramProcess__00000006_.IsExternalParticipantsNodeExpanded = false

	__DiagramProcess__00000007_.Name = `Diagram P1`
	__DiagramProcess__00000007_.ComputedPrefix = ``
	__DiagramProcess__00000007_.IsChecked = true
	__DiagramProcess__00000007_.IsEditable_ = true
	__DiagramProcess__00000007_.IsShowPrefix = false
	__DiagramProcess__00000007_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000007_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000007_.Width = 1475.000000
	__DiagramProcess__00000007_.Height = 1513.000000
	__DiagramProcess__00000007_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000007_.IsParticipantsNodeExpanded = true
	__DiagramProcess__00000007_.IsExternalParticipantsNodeExpanded = true

	__ExternalParticipantShape__00000000_.Name = `EP1-Diagram P1`
	__ExternalParticipantShape__00000000_.IsExpanded = false
	__ExternalParticipantShape__00000000_.X = 116.914807
	__ExternalParticipantShape__00000000_.Y = 396.033942
	__ExternalParticipantShape__00000000_.Width = 10.000000
	__ExternalParticipantShape__00000000_.Height = 155.000000
	__ExternalParticipantShape__00000000_.IsHidden = false
	__ExternalParticipantShape__00000000_.TailHeigth = 0.000000

	__ExternalParticipantShape__00000001_.Name = `EP1-Diagram P1`
	__ExternalParticipantShape__00000001_.IsExpanded = false
	__ExternalParticipantShape__00000001_.X = 41.118302
	__ExternalParticipantShape__00000001_.Y = 286.216751
	__ExternalParticipantShape__00000001_.Width = 120.000000
	__ExternalParticipantShape__00000001_.Height = 70.000000
	__ExternalParticipantShape__00000001_.IsHidden = false
	__ExternalParticipantShape__00000001_.TailHeigth = 0.000000

	__ExternalParticipantShape__00000002_.Name = `EP1-Diagram P1`
	__ExternalParticipantShape__00000002_.IsExpanded = false
	__ExternalParticipantShape__00000002_.X = 28.331631
	__ExternalParticipantShape__00000002_.Y = 247.547418
	__ExternalParticipantShape__00000002_.Width = 130.000000
	__ExternalParticipantShape__00000002_.Height = 70.000000
	__ExternalParticipantShape__00000002_.IsHidden = false
	__ExternalParticipantShape__00000002_.TailHeigth = 0.000000

	__ExternalParticipantShape__00000003_.Name = `EP1-Diagram P1`
	__ExternalParticipantShape__00000003_.IsExpanded = false
	__ExternalParticipantShape__00000003_.X = 5.388870
	__ExternalParticipantShape__00000003_.Y = 293.256779
	__ExternalParticipantShape__00000003_.Width = 219.000000
	__ExternalParticipantShape__00000003_.Height = 70.000000
	__ExternalParticipantShape__00000003_.IsHidden = false
	__ExternalParticipantShape__00000003_.TailHeigth = 233.000000

	__Library__00000000_.Name = `Root`
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsSubLibrariesNodeExpanded = true
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.IsProcessesNodeExpanded = true
	__Library__00000000_.IsDataFlowsNodeExpanded = false
	__Library__00000000_.IsDatasNodeExpanded = false
	__Library__00000000_.IsResourcesNodeExpanded = true
	__Library__00000000_.IsExpandedTmp = true

	__Participant__00000003_.Name = `PP1 Very lon participant name PP1 Very lon participant name`
	__Participant__00000003_.IsResourcesNodeExpanded = true
	__Participant__00000003_.ComputedPrefix = ``
	__Participant__00000003_.IsTasksNodeExpanded = false
	__Participant__00000003_.IsControlFlowsNodeExpanded = false
	__Participant__00000003_.IsDataFlowsNodeExpanded = false

	__Participant__00000004_.Name = `PP2`
	__Participant__00000004_.IsResourcesNodeExpanded = false
	__Participant__00000004_.ComputedPrefix = ``
	__Participant__00000004_.IsTasksNodeExpanded = false
	__Participant__00000004_.IsControlFlowsNodeExpanded = false
	__Participant__00000004_.IsDataFlowsNodeExpanded = false

	__Participant__00000005_.Name = `PP3`
	__Participant__00000005_.IsResourcesNodeExpanded = false
	__Participant__00000005_.ComputedPrefix = ``
	__Participant__00000005_.IsTasksNodeExpanded = false
	__Participant__00000005_.IsControlFlowsNodeExpanded = false
	__Participant__00000005_.IsDataFlowsNodeExpanded = false

	__Participant__00000006_.Name = `PP4`
	__Participant__00000006_.IsResourcesNodeExpanded = false
	__Participant__00000006_.ComputedPrefix = ``
	__Participant__00000006_.IsTasksNodeExpanded = false
	__Participant__00000006_.IsControlFlowsNodeExpanded = false
	__Participant__00000006_.IsDataFlowsNodeExpanded = false

	__Participant__00000007_.Name = `EP1 Very Long Participant Name`
	__Participant__00000007_.IsResourcesNodeExpanded = true
	__Participant__00000007_.ComputedPrefix = ``
	__Participant__00000007_.IsTasksNodeExpanded = false
	__Participant__00000007_.IsControlFlowsNodeExpanded = false
	__Participant__00000007_.IsDataFlowsNodeExpanded = false

	__ParticipantShape__00000019_.Name = `PP3-Diagram P1`
	__ParticipantShape__00000019_.IsExpanded = false
	__ParticipantShape__00000019_.X = 107.650926
	__ParticipantShape__00000019_.Y = 50.648379
	__ParticipantShape__00000019_.Width = 250.000000
	__ParticipantShape__00000019_.Height = 70.000000
	__ParticipantShape__00000019_.IsHidden = false
	__ParticipantShape__00000019_.WidthWeight = 1.000000

	__ParticipantShape__00000020_.Name = `PP2-Diagram P1`
	__ParticipantShape__00000020_.IsExpanded = false
	__ParticipantShape__00000020_.X = 559.000000
	__ParticipantShape__00000020_.Y = 166.000000
	__ParticipantShape__00000020_.Width = 303.000000
	__ParticipantShape__00000020_.Height = 1037.000000
	__ParticipantShape__00000020_.IsHidden = false
	__ParticipantShape__00000020_.WidthWeight = 1.000000

	__ParticipantShape__00000022_.Name = `PP1 Very lon participant name PP1 Very lon participant name-Diagram P1`
	__ParticipantShape__00000022_.IsExpanded = false
	__ParticipantShape__00000022_.X = 192.968897
	__ParticipantShape__00000022_.Y = 162.495338
	__ParticipantShape__00000022_.Width = 250.000000
	__ParticipantShape__00000022_.Height = 70.000000
	__ParticipantShape__00000022_.IsHidden = false
	__ParticipantShape__00000022_.WidthWeight = 1.000000

	__Process__00000005_.Name = `P1`
	__Process__00000005_.ComputedPrefix = ``
	__Process__00000005_.IsSubProcessNodeExpanded = false
	__Process__00000005_.IsDataFlowsNodeExpanded = false

	__Process__00000006_.Name = `P2`
	__Process__00000006_.ComputedPrefix = ``
	__Process__00000006_.IsSubProcessNodeExpanded = false
	__Process__00000006_.IsDataFlowsNodeExpanded = false

	__ProcessShape__00000006_.Name = `ProcessShape`
	__ProcessShape__00000006_.IsExpanded = false
	__ProcessShape__00000006_.X = 100.000000
	__ProcessShape__00000006_.Y = 50.000000
	__ProcessShape__00000006_.Width = 500.000000
	__ProcessShape__00000006_.Height = 1000.000000
	__ProcessShape__00000006_.IsHidden = false

	__ProcessShape__00000007_.Name = `ProcessShape`
	__ProcessShape__00000007_.IsExpanded = false
	__ProcessShape__00000007_.X = 246.000000
	__ProcessShape__00000007_.Y = 56.000000
	__ProcessShape__00000007_.Width = 929.000000
	__ProcessShape__00000007_.Height = 1157.000000
	__ProcessShape__00000007_.IsHidden = false

	__Resource__00000000_.Name = `R1 with a long name long long long`
	__Resource__00000000_.ComputedPrefix = ``
	__Resource__00000000_.SVG_Path = `m280-40 112-564-72 28v136h-80v-188l202-86q14-6 29.5-7t29.5 4q14 5 26.5 14t20.5 23l40 64q26 42 70.5 69T760-520v80q-70 0-125-29t-94-74l-25 123 84 80v300h-80v-260l-84-64-72 324h-84Zm203.5-723.5Q460-787 460-820t23.5-56.5Q507-900 540-900t56.5 23.5Q620-853 620-820t-23.5 56.5Q573-740 540-740t-56.5-23.5Z`
	__Resource__00000000_.InverseAppliedScaling = 60.000000

	__Resource__00000001_.Name = `R2aaaaa`
	__Resource__00000001_.ComputedPrefix = ``
	__Resource__00000001_.SVG_Path = `M240-200v40q0 17-11.5 28.5T200-120h-40q-17 0-28.5-11.5T120-160v-320l84-240q6-18 21.5-29t34.5-11h440q19 0 34.5 11t21.5 29l84 240v320q0 17-11.5 28.5T800-120h-40q-17 0-28.5-11.5T720-160v-40H240Zm-8-360h496l-42-120H274l-42 120Zm-32 80v200-200Zm100 160q25 0 42.5-17.5T360-380q0-25-17.5-42.5T300-440q-25 0-42.5 17.5T240-380q0 25 17.5 42.5T300-320Zm360 0q25 0 42.5-17.5T720-380q0-25-17.5-42.5T660-440q-25 0-42.5 17.5T600-380q0 25 17.5 42.5T660-320Zm-460 40h560v-200H200v200Z`
	__Resource__00000001_.InverseAppliedScaling = 60.000000

	__Task__00000002_.Name = `T2`
	__Task__00000002_.ComputedPrefix = ``
	__Task__00000002_.IsStartTask = false
	__Task__00000002_.IsEndTask = false

	__Task__00000003_.Name = `PP2 T1`
	__Task__00000003_.ComputedPrefix = ``
	__Task__00000003_.IsStartTask = true
	__Task__00000003_.IsEndTask = false

	__Task__00000004_.Name = `Start`
	__Task__00000004_.ComputedPrefix = ``
	__Task__00000004_.IsStartTask = true
	__Task__00000004_.IsEndTask = false

	__Task__00000005_.Name = `End`
	__Task__00000005_.ComputedPrefix = ``
	__Task__00000005_.IsStartTask = false
	__Task__00000005_.IsEndTask = true

	__Task__00000006_.Name = `T4`
	__Task__00000006_.ComputedPrefix = ``
	__Task__00000006_.IsStartTask = false
	__Task__00000006_.IsEndTask = false

	__Task__00000007_.Name = `PP2 T2`
	__Task__00000007_.ComputedPrefix = ``
	__Task__00000007_.IsStartTask = false
	__Task__00000007_.IsEndTask = false

	__TaskShape__00000011_.Name = `Start-DP2`
	__TaskShape__00000011_.IsExpanded = false
	__TaskShape__00000011_.X = 358.000000
	__TaskShape__00000011_.Y = 281.000000
	__TaskShape__00000011_.Width = 54.000000
	__TaskShape__00000011_.Height = 20.000000
	__TaskShape__00000011_.IsHidden = false

	__TaskShape__00000012_.Name = `T4-DP2`
	__TaskShape__00000012_.IsExpanded = false
	__TaskShape__00000012_.X = 315.250000
	__TaskShape__00000012_.Y = 743.000000
	__TaskShape__00000012_.Width = 167.000000
	__TaskShape__00000012_.Height = 70.000000
	__TaskShape__00000012_.IsHidden = false

	__TaskShape__00000013_.Name = `T2-DP2`
	__TaskShape__00000013_.IsExpanded = false
	__TaskShape__00000013_.X = 304.000000
	__TaskShape__00000013_.Y = 406.000000
	__TaskShape__00000013_.Width = 174.000000
	__TaskShape__00000013_.Height = 240.000000
	__TaskShape__00000013_.IsHidden = false

	__TaskShape__00000014_.Name = `PP2 T2-DP2`
	__TaskShape__00000014_.IsExpanded = false
	__TaskShape__00000014_.X = 560.000000
	__TaskShape__00000014_.Y = 397.000000
	__TaskShape__00000014_.Width = 211.000000
	__TaskShape__00000014_.Height = 70.000000
	__TaskShape__00000014_.IsHidden = false

	__TaskShape__00000015_.Name = `PP2 T1-Diagram P1`
	__TaskShape__00000015_.IsExpanded = false
	__TaskShape__00000015_.X = 634.500000
	__TaskShape__00000015_.Y = 261.000000
	__TaskShape__00000015_.Width = 75.000000
	__TaskShape__00000015_.Height = 20.000000
	__TaskShape__00000015_.IsHidden = false

	__TaskShape__00000016_.Name = `End-Diagram P1`
	__TaskShape__00000016_.IsExpanded = false
	__TaskShape__00000016_.X = 317.000000
	__TaskShape__00000016_.Y = 966.000000
	__TaskShape__00000016_.Width = 89.000000
	__TaskShape__00000016_.Height = 36.000000
	__TaskShape__00000016_.IsHidden = false

	// insertion point for setup of pointers
	__AllocatedResourceShape__00000010_.Participant = __Participant__00000007_
	__AllocatedResourceShape__00000010_.Resource = __Resource__00000001_
	__AllocatedResourceShape__00000012_.Participant = __Participant__00000007_
	__AllocatedResourceShape__00000012_.Resource = __Resource__00000000_
	__AllocatedResourceShape__00000014_.Participant = __Participant__00000003_
	__AllocatedResourceShape__00000014_.Resource = __Resource__00000001_
	__AllocatedResourceShape__00000015_.Participant = __Participant__00000003_
	__AllocatedResourceShape__00000015_.Resource = __Resource__00000000_
	__ControlFlow__00000003_.Start = __Task__00000004_
	__ControlFlow__00000003_.End = __Task__00000002_
	__ControlFlow__00000004_.Start = __Task__00000002_
	__ControlFlow__00000004_.End = __Task__00000006_
	__ControlFlow__00000005_.Start = __Task__00000003_
	__ControlFlow__00000005_.End = __Task__00000007_
	__ControlFlow__00000006_.Start = __Task__00000006_
	__ControlFlow__00000006_.End = __Task__00000005_
	__ControlFlowShape__00000000_.ControlFlow = __ControlFlow__00000003_
	__ControlFlowShape__00000001_.ControlFlow = __ControlFlow__00000005_
	__ControlFlowShape__00000002_.ControlFlow = __ControlFlow__00000004_
	__ControlFlowShape__00000003_.ControlFlow = __ControlFlow__00000006_
	__DataFlow__00000000_.StartTask = __Task__00000002_
	__DataFlow__00000000_.EndTask = __Task__00000007_
	__DataFlow__00000000_.StartExternalParticipant = nil
	__DataFlow__00000000_.EndExternalParticipant = nil
	__DataFlow__00000000_.Datas = append(__DataFlow__00000000_.Datas, __Data__00000000_)
	__DataFlow__00000001_.StartTask = nil
	__DataFlow__00000001_.EndTask = __Task__00000002_
	__DataFlow__00000001_.StartExternalParticipant = __Participant__00000007_
	__DataFlow__00000001_.EndExternalParticipant = nil
	__DataFlow__00000001_.Datas = append(__DataFlow__00000001_.Datas, __Data__00000001_)
	__DataFlow__00000001_.Datas = append(__DataFlow__00000001_.Datas, __Data__00000000_)
	__DataFlowShape__00000003_.DataFlow = __DataFlow__00000001_
	__DataFlowShape__00000004_.DataFlow = __DataFlow__00000000_
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
	__DataShape__00000022_.DataFlow = __DataFlow__00000001_
	__DataShape__00000024_.Data = __Data__00000000_
	__DataShape__00000024_.DataFlow = __DataFlow__00000001_
	__DiagramProcess__00000006_.Process_Shapes = append(__DiagramProcess__00000006_.Process_Shapes, __ProcessShape__00000006_)
	__DiagramProcess__00000007_.Process_Shapes = append(__DiagramProcess__00000007_.Process_Shapes, __ProcessShape__00000007_)
	__DiagramProcess__00000007_.Participant_Shapes = append(__DiagramProcess__00000007_.Participant_Shapes, __ParticipantShape__00000022_)
	__DiagramProcess__00000007_.Participant_Shapes = append(__DiagramProcess__00000007_.Participant_Shapes, __ParticipantShape__00000020_)
	__DiagramProcess__00000007_.Participant_Shapes = append(__DiagramProcess__00000007_.Participant_Shapes, __ParticipantShape__00000019_)
	__DiagramProcess__00000007_.ParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000007_.ParticipantWhoseNodeIsExpanded, __Participant__00000003_)
	__DiagramProcess__00000007_.ExternalParticipant_Shapes = append(__DiagramProcess__00000007_.ExternalParticipant_Shapes, __ExternalParticipantShape__00000003_)
	__DiagramProcess__00000007_.ExternalParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000007_.ExternalParticipantWhoseNodeIsExpanded, __Participant__00000003_)
	__DiagramProcess__00000007_.ExternalParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000007_.ExternalParticipantWhoseNodeIsExpanded, __Participant__00000007_)
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
	__DiagramProcess__00000007_.DataFlow_Shapes = append(__DiagramProcess__00000007_.DataFlow_Shapes, __DataFlowShape__00000003_)
	__DiagramProcess__00000007_.DataFlow_Shapes = append(__DiagramProcess__00000007_.DataFlow_Shapes, __DataFlowShape__00000004_)
	__DiagramProcess__00000007_.Data_Shapes = append(__DiagramProcess__00000007_.Data_Shapes, __DataShape__00000011_)
	__DiagramProcess__00000007_.Data_Shapes = append(__DiagramProcess__00000007_.Data_Shapes, __DataShape__00000012_)
	__DiagramProcess__00000007_.Data_Shapes = append(__DiagramProcess__00000007_.Data_Shapes, __DataShape__00000013_)
	__DiagramProcess__00000007_.Data_Shapes = append(__DiagramProcess__00000007_.Data_Shapes, __DataShape__00000014_)
	__DiagramProcess__00000007_.Data_Shapes = append(__DiagramProcess__00000007_.Data_Shapes, __DataShape__00000016_)
	__DiagramProcess__00000007_.Data_Shapes = append(__DiagramProcess__00000007_.Data_Shapes, __DataShape__00000017_)
	__DiagramProcess__00000007_.Data_Shapes = append(__DiagramProcess__00000007_.Data_Shapes, __DataShape__00000019_)
	__DiagramProcess__00000007_.Data_Shapes = append(__DiagramProcess__00000007_.Data_Shapes, __DataShape__00000022_)
	__DiagramProcess__00000007_.Data_Shapes = append(__DiagramProcess__00000007_.Data_Shapes, __DataShape__00000024_)
	__DiagramProcess__00000007_.DataFlowsWhoseDataNodeIsExpanded = append(__DiagramProcess__00000007_.DataFlowsWhoseDataNodeIsExpanded, __DataFlow__00000000_)
	__DiagramProcess__00000007_.DataFlowsWhoseDataNodeIsExpanded = append(__DiagramProcess__00000007_.DataFlowsWhoseDataNodeIsExpanded, __DataFlow__00000001_)
	__DiagramProcess__00000007_.AllocatedResourceShapes = append(__DiagramProcess__00000007_.AllocatedResourceShapes, __AllocatedResourceShape__00000010_)
	__DiagramProcess__00000007_.AllocatedResourceShapes = append(__DiagramProcess__00000007_.AllocatedResourceShapes, __AllocatedResourceShape__00000012_)
	__DiagramProcess__00000007_.AllocatedResourceShapes = append(__DiagramProcess__00000007_.AllocatedResourceShapes, __AllocatedResourceShape__00000014_)
	__DiagramProcess__00000007_.AllocatedResourceShapes = append(__DiagramProcess__00000007_.AllocatedResourceShapes, __AllocatedResourceShape__00000015_)
	__ExternalParticipantShape__00000000_.Participant = __Participant__00000007_
	__ExternalParticipantShape__00000001_.Participant = __Participant__00000007_
	__ExternalParticipantShape__00000002_.Participant = __Participant__00000007_
	__ExternalParticipantShape__00000003_.Participant = __Participant__00000007_
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000005_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000006_)
	__Library__00000000_.ProcesssWhoseNodeIsExpanded = append(__Library__00000000_.ProcesssWhoseNodeIsExpanded, __Process__00000006_)
	__Library__00000000_.ProcesssWhoseNodeIsExpanded = append(__Library__00000000_.ProcesssWhoseNodeIsExpanded, __Process__00000005_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000000_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000001_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000000_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000001_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000000_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000001_)
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
	__ParticipantShape__00000019_.Participant = __Participant__00000005_
	__ParticipantShape__00000020_.Participant = __Participant__00000004_
	__ParticipantShape__00000022_.Participant = __Participant__00000003_
	__Process__00000005_.DiagramProcesss = append(__Process__00000005_.DiagramProcesss, __DiagramProcess__00000007_)
	__Process__00000005_.DiagramProcessWhoseNodeIsExpanded = append(__Process__00000005_.DiagramProcessWhoseNodeIsExpanded, __DiagramProcess__00000007_)
	__Process__00000005_.Participants = append(__Process__00000005_.Participants, __Participant__00000003_)
	__Process__00000005_.Participants = append(__Process__00000005_.Participants, __Participant__00000004_)
	__Process__00000005_.Participants = append(__Process__00000005_.Participants, __Participant__00000005_)
	__Process__00000005_.Participants = append(__Process__00000005_.Participants, __Participant__00000006_)
	__Process__00000005_.ParticipantWhoseNodeIsExpanded = append(__Process__00000005_.ParticipantWhoseNodeIsExpanded, __Participant__00000003_)
	__Process__00000005_.DataFlows = append(__Process__00000005_.DataFlows, __DataFlow__00000000_)
	__Process__00000005_.DataFlows = append(__Process__00000005_.DataFlows, __DataFlow__00000001_)
	__Process__00000005_.ExternalParticipants = append(__Process__00000005_.ExternalParticipants, __Participant__00000007_)
	__Process__00000006_.DiagramProcesss = append(__Process__00000006_.DiagramProcesss, __DiagramProcess__00000006_)
	__ProcessShape__00000006_.Process = __Process__00000006_
	__ProcessShape__00000007_.Process = __Process__00000005_
	__TaskShape__00000011_.Task = __Task__00000004_
	__TaskShape__00000012_.Task = __Task__00000006_
	__TaskShape__00000013_.Task = __Task__00000002_
	__TaskShape__00000014_.Task = __Task__00000007_
	__TaskShape__00000015_.Task = __Task__00000003_
	__TaskShape__00000016_.Task = __Task__00000005_
}
