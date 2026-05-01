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

	__ControlFlow__00000003_ := (&models.ControlFlow{Name: `Start to T2`}).Stage(stage)
	__ControlFlow__00000004_ := (&models.ControlFlow{Name: `T2 to End`}).Stage(stage)

	__ControlFlowShape__00000003_ := (&models.ControlFlowShape{Name: `Start to T2`}).Stage(stage)
	__ControlFlowShape__00000004_ := (&models.ControlFlowShape{Name: `T2 to End`}).Stage(stage)

	__DataFlow__00000000_ := (&models.DataFlow{Name: `T2 to PP2 T2`}).Stage(stage)
	__DataFlow__00000001_ := (&models.DataFlow{Name: `PP2 T2 to T2`}).Stage(stage)

	__DataFlowShape__00000001_ := (&models.DataFlowShape{Name: `T2 to PP2 T2`}).Stage(stage)
	__DataFlowShape__00000002_ := (&models.DataFlowShape{Name: `PP2 T2 to T2`}).Stage(stage)

	__DiagramProcess__00000005_ := (&models.DiagramProcess{Name: `DiagramProcess`}).Stage(stage)
	__DiagramProcess__00000006_ := (&models.DiagramProcess{Name: `DiagramProcess`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `Root`}).Stage(stage)

	__Participant__00000003_ := (&models.Participant{Name: `PP1`}).Stage(stage)
	__Participant__00000004_ := (&models.Participant{Name: `PP2`}).Stage(stage)
	__Participant__00000005_ := (&models.Participant{Name: `PP3`}).Stage(stage)
	__Participant__00000006_ := (&models.Participant{Name: `PP4`}).Stage(stage)

	__ParticipantShape__00000006_ := (&models.ParticipantShape{Name: `PP2-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000007_ := (&models.ParticipantShape{Name: `PP3-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000009_ := (&models.ParticipantShape{Name: `PP1-DiagramProcess`}).Stage(stage)

	__Process__00000005_ := (&models.Process{Name: `P1`}).Stage(stage)
	__Process__00000006_ := (&models.Process{Name: `P2`}).Stage(stage)

	__ProcessShape__00000005_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)
	__ProcessShape__00000006_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)

	__Task__00000002_ := (&models.Task{Name: `T2`}).Stage(stage)
	__Task__00000003_ := (&models.Task{Name: `PP2 T1`}).Stage(stage)
	__Task__00000004_ := (&models.Task{Name: `Start`}).Stage(stage)
	__Task__00000005_ := (&models.Task{Name: `End`}).Stage(stage)
	__Task__00000006_ := (&models.Task{Name: `T4`}).Stage(stage)
	__Task__00000007_ := (&models.Task{Name: `PP2 T2`}).Stage(stage)

	__TaskShape__00000002_ := (&models.TaskShape{Name: `T2-DiagramProcess`}).Stage(stage)
	__TaskShape__00000004_ := (&models.TaskShape{Name: `Start-DiagramProcess`}).Stage(stage)
	__TaskShape__00000005_ := (&models.TaskShape{Name: `PP2 T1-DiagramProcess`}).Stage(stage)
	__TaskShape__00000007_ := (&models.TaskShape{Name: `End-DiagramProcess`}).Stage(stage)
	__TaskShape__00000008_ := (&models.TaskShape{Name: `-DiagramProcess`}).Stage(stage)

	// insertion point for initialization of values

	__ControlFlow__00000003_.Name = `Start to T2`
	__ControlFlow__00000003_.ComputedPrefix = ``

	__ControlFlow__00000004_.Name = `T2 to End`
	__ControlFlow__00000004_.ComputedPrefix = ``

	__ControlFlowShape__00000003_.Name = `Start to T2`
	__ControlFlowShape__00000003_.StartRatio = 1.000000
	__ControlFlowShape__00000003_.EndRatio = 0.500000
	__ControlFlowShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000003_.CornerOffsetRatio = 6.316101
	__ControlFlowShape__00000003_.IsHidden = false

	__ControlFlowShape__00000004_.Name = `T2 to End`
	__ControlFlowShape__00000004_.StartRatio = 0.500000
	__ControlFlowShape__00000004_.EndRatio = 0.500000
	__ControlFlowShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000004_.CornerOffsetRatio = 1.500000
	__ControlFlowShape__00000004_.IsHidden = false

	__DataFlow__00000000_.Name = `T2 to PP2 T2`
	__DataFlow__00000000_.ComputedPrefix = ``

	__DataFlow__00000001_.Name = `PP2 T2 to T2`
	__DataFlow__00000001_.ComputedPrefix = ``

	__DataFlowShape__00000001_.Name = `T2 to PP2 T2`
	__DataFlowShape__00000001_.StartRatio = 0.688864
	__DataFlowShape__00000001_.EndRatio = 0.500000
	__DataFlowShape__00000001_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__DataFlowShape__00000001_.CornerOffsetRatio = 1.051558
	__DataFlowShape__00000001_.IsHidden = true

	__DataFlowShape__00000002_.Name = `PP2 T2 to T2`
	__DataFlowShape__00000002_.StartRatio = 0.500000
	__DataFlowShape__00000002_.EndRatio = 0.919979
	__DataFlowShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__DataFlowShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__DataFlowShape__00000002_.CornerOffsetRatio = -0.329531
	__DataFlowShape__00000002_.IsHidden = false

	__DiagramProcess__00000005_.Name = `DiagramProcess`
	__DiagramProcess__00000005_.ComputedPrefix = ``
	__DiagramProcess__00000005_.IsChecked = true
	__DiagramProcess__00000005_.IsEditable_ = true
	__DiagramProcess__00000005_.IsShowPrefix = false
	__DiagramProcess__00000005_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000005_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000005_.Width = 1521.000000
	__DiagramProcess__00000005_.Height = 1318.000000
	__DiagramProcess__00000005_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000005_.IsParticipantsNodeExpanded = false

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

	__Library__00000000_.Name = `Root`
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.IsExpandedTmp = true

	__Participant__00000003_.Name = `PP1`
	__Participant__00000003_.ComputedPrefix = ``
	__Participant__00000003_.IsTasksNodeExpanded = false
	__Participant__00000003_.IsControlFlowsNodeExpanded = false

	__Participant__00000004_.Name = `PP2`
	__Participant__00000004_.ComputedPrefix = ``
	__Participant__00000004_.IsTasksNodeExpanded = true
	__Participant__00000004_.IsControlFlowsNodeExpanded = false

	__Participant__00000005_.Name = `PP3`
	__Participant__00000005_.ComputedPrefix = ``
	__Participant__00000005_.IsTasksNodeExpanded = false
	__Participant__00000005_.IsControlFlowsNodeExpanded = false

	__Participant__00000006_.Name = `PP4`
	__Participant__00000006_.ComputedPrefix = ``
	__Participant__00000006_.IsTasksNodeExpanded = false
	__Participant__00000006_.IsControlFlowsNodeExpanded = false

	__ParticipantShape__00000006_.Name = `PP2-DiagramProcess`
	__ParticipantShape__00000006_.IsExpanded = false
	__ParticipantShape__00000006_.X = 461.000000
	__ParticipantShape__00000006_.Y = 68.000000
	__ParticipantShape__00000006_.Width = 375.000000
	__ParticipantShape__00000006_.Height = 940.000000
	__ParticipantShape__00000006_.IsHidden = false

	__ParticipantShape__00000007_.Name = `PP3-DiagramProcess`
	__ParticipantShape__00000007_.IsExpanded = false
	__ParticipantShape__00000007_.X = 673.250000
	__ParticipantShape__00000007_.Y = 68.000000
	__ParticipantShape__00000007_.Width = 195.750000
	__ParticipantShape__00000007_.Height = 940.000000
	__ParticipantShape__00000007_.IsHidden = false

	__ParticipantShape__00000009_.Name = `PP1-DiagramProcess`
	__ParticipantShape__00000009_.IsExpanded = false
	__ParticipantShape__00000009_.X = 225.000000
	__ParticipantShape__00000009_.Y = 68.000000
	__ParticipantShape__00000009_.Width = 328.666667
	__ParticipantShape__00000009_.Height = 940.000000
	__ParticipantShape__00000009_.IsHidden = false

	__Process__00000005_.Name = `P1`
	__Process__00000005_.ComputedPrefix = ``
	__Process__00000005_.IsSubProcessNodeExpanded = false
	__Process__00000005_.IsDataFlowsNodeExpanded = false

	__Process__00000006_.Name = `P2`
	__Process__00000006_.ComputedPrefix = ``
	__Process__00000006_.IsSubProcessNodeExpanded = false
	__Process__00000006_.IsDataFlowsNodeExpanded = false

	__ProcessShape__00000005_.Name = `ProcessShape`
	__ProcessShape__00000005_.IsExpanded = false
	__ProcessShape__00000005_.X = 215.000000
	__ProcessShape__00000005_.Y = 18.000000
	__ProcessShape__00000005_.Width = 1006.000000
	__ProcessShape__00000005_.Height = 1000.000000
	__ProcessShape__00000005_.IsHidden = false

	__ProcessShape__00000006_.Name = `ProcessShape`
	__ProcessShape__00000006_.IsExpanded = false
	__ProcessShape__00000006_.X = 100.000000
	__ProcessShape__00000006_.Y = 50.000000
	__ProcessShape__00000006_.Width = 500.000000
	__ProcessShape__00000006_.Height = 1000.000000
	__ProcessShape__00000006_.IsHidden = false

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

	__TaskShape__00000002_.Name = `T2-DiagramProcess`
	__TaskShape__00000002_.IsExpanded = false
	__TaskShape__00000002_.X = 314.000000
	__TaskShape__00000002_.Y = 376.786461
	__TaskShape__00000002_.Width = 152.000000
	__TaskShape__00000002_.Height = 70.000000
	__TaskShape__00000002_.IsHidden = false

	__TaskShape__00000004_.Name = `Start-DiagramProcess`
	__TaskShape__00000004_.IsExpanded = false
	__TaskShape__00000004_.X = 314.000000
	__TaskShape__00000004_.Y = 132.683844
	__TaskShape__00000004_.Width = 58.000000
	__TaskShape__00000004_.Height = 20.000000
	__TaskShape__00000004_.IsHidden = false

	__TaskShape__00000005_.Name = `PP2 T1-DiagramProcess`
	__TaskShape__00000005_.IsExpanded = false
	__TaskShape__00000005_.X = 572.500000
	__TaskShape__00000005_.Y = 179.449476
	__TaskShape__00000005_.Width = 75.000000
	__TaskShape__00000005_.Height = 20.000000
	__TaskShape__00000005_.IsHidden = false

	__TaskShape__00000007_.Name = `End-DiagramProcess`
	__TaskShape__00000007_.IsExpanded = false
	__TaskShape__00000007_.X = 314.000000
	__TaskShape__00000007_.Y = 706.136264
	__TaskShape__00000007_.Width = 79.000000
	__TaskShape__00000007_.Height = 36.000000
	__TaskShape__00000007_.IsHidden = false

	__TaskShape__00000008_.Name = `-DiagramProcess`
	__TaskShape__00000008_.IsExpanded = false
	__TaskShape__00000008_.X = 554.666667
	__TaskShape__00000008_.Y = 336.074115
	__TaskShape__00000008_.Width = 250.000000
	__TaskShape__00000008_.Height = 70.000000
	__TaskShape__00000008_.IsHidden = false

	// insertion point for setup of pointers
	__ControlFlow__00000003_.Start = __Task__00000004_
	__ControlFlow__00000003_.End = __Task__00000002_
	__ControlFlow__00000004_.Start = __Task__00000002_
	__ControlFlow__00000004_.End = __Task__00000005_
	__ControlFlowShape__00000003_.ControlFlow = __ControlFlow__00000003_
	__ControlFlowShape__00000004_.ControlFlow = __ControlFlow__00000004_
	__DataFlow__00000000_.Start = __Task__00000002_
	__DataFlow__00000000_.End = __Task__00000007_
	__DataFlow__00000001_.Start = __Task__00000007_
	__DataFlow__00000001_.End = __Task__00000002_
	__DataFlowShape__00000001_.DataFlow = __DataFlow__00000000_
	__DataFlowShape__00000002_.DataFlow = __DataFlow__00000001_
	__DiagramProcess__00000005_.Process_Shapes = append(__DiagramProcess__00000005_.Process_Shapes, __ProcessShape__00000005_)
	__DiagramProcess__00000005_.Participant_Shapes = append(__DiagramProcess__00000005_.Participant_Shapes, __ParticipantShape__00000009_)
	__DiagramProcess__00000005_.Participant_Shapes = append(__DiagramProcess__00000005_.Participant_Shapes, __ParticipantShape__00000006_)
	__DiagramProcess__00000005_.Participant_Shapes = append(__DiagramProcess__00000005_.Participant_Shapes, __ParticipantShape__00000007_)
	__DiagramProcess__00000005_.ParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000005_.ParticipantWhoseNodeIsExpanded, __Participant__00000006_)
	__DiagramProcess__00000005_.ParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000005_.ParticipantWhoseNodeIsExpanded, __Participant__00000004_)
	__DiagramProcess__00000005_.TasksWhoseNodeIsExpanded = append(__DiagramProcess__00000005_.TasksWhoseNodeIsExpanded, __Task__00000003_)
	__DiagramProcess__00000005_.TasksWhoseNodeIsExpanded = append(__DiagramProcess__00000005_.TasksWhoseNodeIsExpanded, __Task__00000006_)
	__DiagramProcess__00000005_.TasksWhoseNodeIsExpanded = append(__DiagramProcess__00000005_.TasksWhoseNodeIsExpanded, __Task__00000002_)
	__DiagramProcess__00000005_.TaskShapes = append(__DiagramProcess__00000005_.TaskShapes, __TaskShape__00000002_)
	__DiagramProcess__00000005_.TaskShapes = append(__DiagramProcess__00000005_.TaskShapes, __TaskShape__00000004_)
	__DiagramProcess__00000005_.TaskShapes = append(__DiagramProcess__00000005_.TaskShapes, __TaskShape__00000005_)
	__DiagramProcess__00000005_.TaskShapes = append(__DiagramProcess__00000005_.TaskShapes, __TaskShape__00000007_)
	__DiagramProcess__00000005_.TaskShapes = append(__DiagramProcess__00000005_.TaskShapes, __TaskShape__00000008_)
	__DiagramProcess__00000005_.ControlFlowShapes = append(__DiagramProcess__00000005_.ControlFlowShapes, __ControlFlowShape__00000003_)
	__DiagramProcess__00000005_.ControlFlowShapes = append(__DiagramProcess__00000005_.ControlFlowShapes, __ControlFlowShape__00000004_)
	__DiagramProcess__00000005_.DataFlowShapes = append(__DiagramProcess__00000005_.DataFlowShapes, __DataFlowShape__00000001_)
	__DiagramProcess__00000005_.DataFlowShapes = append(__DiagramProcess__00000005_.DataFlowShapes, __DataFlowShape__00000002_)
	__DiagramProcess__00000006_.Process_Shapes = append(__DiagramProcess__00000006_.Process_Shapes, __ProcessShape__00000006_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000005_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000006_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000000_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000001_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000002_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000004_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000005_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000006_)
	__Participant__00000003_.ControlFlows = append(__Participant__00000003_.ControlFlows, __ControlFlow__00000003_)
	__Participant__00000003_.ControlFlows = append(__Participant__00000003_.ControlFlows, __ControlFlow__00000004_)
	__Participant__00000003_.TaskWhoseOutControlFlowsNodeIsExpanded = append(__Participant__00000003_.TaskWhoseOutControlFlowsNodeIsExpanded, __Task__00000002_)
	__Participant__00000004_.Tasks = append(__Participant__00000004_.Tasks, __Task__00000003_)
	__Participant__00000004_.Tasks = append(__Participant__00000004_.Tasks, __Task__00000007_)
	__ParticipantShape__00000006_.Participant = __Participant__00000004_
	__ParticipantShape__00000007_.Participant = __Participant__00000005_
	__ParticipantShape__00000009_.Participant = __Participant__00000003_
	__Process__00000005_.DiagramProcesss = append(__Process__00000005_.DiagramProcesss, __DiagramProcess__00000005_)
	__Process__00000005_.DiagramProcessWhoseNodeIsExpanded = append(__Process__00000005_.DiagramProcessWhoseNodeIsExpanded, __DiagramProcess__00000005_)
	__Process__00000005_.Participants = append(__Process__00000005_.Participants, __Participant__00000003_)
	__Process__00000005_.Participants = append(__Process__00000005_.Participants, __Participant__00000004_)
	__Process__00000005_.Participants = append(__Process__00000005_.Participants, __Participant__00000005_)
	__Process__00000005_.Participants = append(__Process__00000005_.Participants, __Participant__00000006_)
	__Process__00000005_.ParticipantWhoseNodeIsExpanded = append(__Process__00000005_.ParticipantWhoseNodeIsExpanded, __Participant__00000003_)
	__Process__00000005_.DataFlows = append(__Process__00000005_.DataFlows, __DataFlow__00000000_)
	__Process__00000005_.DataFlows = append(__Process__00000005_.DataFlows, __DataFlow__00000001_)
	__Process__00000006_.DiagramProcesss = append(__Process__00000006_.DiagramProcesss, __DiagramProcess__00000006_)
	__ProcessShape__00000005_.Process = __Process__00000005_
	__ProcessShape__00000006_.Process = __Process__00000006_
	__TaskShape__00000002_.Task = __Task__00000002_
	__TaskShape__00000004_.Task = __Task__00000004_
	__TaskShape__00000005_.Task = __Task__00000003_
	__TaskShape__00000007_.Task = __Task__00000005_
	__TaskShape__00000008_.Task = __Task__00000007_
}
