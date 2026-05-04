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
	__ControlFlow__00000005_ := (&models.ControlFlow{Name: `PP2 T1 to PP2 T2`}).Stage(stage)

	__ControlFlowShape__00000000_ := (&models.ControlFlowShape{Name: `Start to T2`}).Stage(stage)
	__ControlFlowShape__00000001_ := (&models.ControlFlowShape{Name: `PP2 T1 to PP2 T2`}).Stage(stage)

	__Data__00000000_ := (&models.Data{Name: `D1`}).Stage(stage)
	__Data__00000001_ := (&models.Data{Name: `D2`}).Stage(stage)
	__Data__00000002_ := (&models.Data{Name: `L1.D1`}).Stage(stage)

	__DataFlow__00000000_ := (&models.DataFlow{Name: `T2 to PP2 T2`}).Stage(stage)
	__DataFlow__00000001_ := (&models.DataFlow{Name: `PP2 T2 to T2`}).Stage(stage)

	__DataFlowShape__00000001_ := (&models.DataFlowShape{Name: `PP2 T2 to T2`}).Stage(stage)

	__DiagramProcess__00000006_ := (&models.DiagramProcess{Name: `DiagramProcess`}).Stage(stage)
	__DiagramProcess__00000007_ := (&models.DiagramProcess{Name: `Diagram P1`}).Stage(stage)
	__DiagramProcess__00000008_ := (&models.DiagramProcess{Name: `DiagramProcess`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `Root`}).Stage(stage)
	__Library__00000001_ := (&models.Library{Name: `L1`}).Stage(stage)
	__Library__00000002_ := (&models.Library{Name: `L2`}).Stage(stage)

	__Participant__00000003_ := (&models.Participant{Name: `PP1`}).Stage(stage)
	__Participant__00000004_ := (&models.Participant{Name: `PP2`}).Stage(stage)
	__Participant__00000005_ := (&models.Participant{Name: `PP3`}).Stage(stage)
	__Participant__00000006_ := (&models.Participant{Name: `PP4`}).Stage(stage)

	__ParticipantShape__00000010_ := (&models.ParticipantShape{Name: `PP1-DP2`}).Stage(stage)
	__ParticipantShape__00000012_ := (&models.ParticipantShape{Name: `PP2-Diagram P1`}).Stage(stage)

	__Process__00000005_ := (&models.Process{Name: `P1`}).Stage(stage)
	__Process__00000006_ := (&models.Process{Name: `P2`}).Stage(stage)
	__Process__00000007_ := (&models.Process{Name: `L2.P1`}).Stage(stage)

	__ProcessShape__00000006_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)
	__ProcessShape__00000007_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)
	__ProcessShape__00000008_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)

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

	// insertion point for initialization of values

	__ControlFlow__00000003_.Name = `Start to T2`
	__ControlFlow__00000003_.ComputedPrefix = ``

	__ControlFlow__00000004_.Name = `T2 to End`
	__ControlFlow__00000004_.ComputedPrefix = ``

	__ControlFlow__00000005_.Name = `PP2 T1 to PP2 T2`
	__ControlFlow__00000005_.ComputedPrefix = ``

	__ControlFlowShape__00000000_.Name = `Start to T2`
	__ControlFlowShape__00000000_.StartRatio = 0.971920
	__ControlFlowShape__00000000_.EndRatio = 0.500000
	__ControlFlowShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000000_.CornerOffsetRatio = 4.178095
	__ControlFlowShape__00000000_.IsHidden = false

	__ControlFlowShape__00000001_.Name = `PP2 T1 to PP2 T2`
	__ControlFlowShape__00000001_.StartRatio = 0.500000
	__ControlFlowShape__00000001_.EndRatio = 0.500000
	__ControlFlowShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000001_.CornerOffsetRatio = 2.823598
	__ControlFlowShape__00000001_.IsHidden = false

	__Data__00000000_.Name = `D1`
	__Data__00000000_.ComputedPrefix = ``

	__Data__00000001_.Name = `D2`
	__Data__00000001_.ComputedPrefix = ``

	__Data__00000002_.Name = `L1.D1`
	__Data__00000002_.ComputedPrefix = ``

	__DataFlow__00000000_.Name = `T2 to PP2 T2`
	__DataFlow__00000000_.ComputedPrefix = ``

	__DataFlow__00000001_.Name = `PP2 T2 to T2`
	__DataFlow__00000001_.ComputedPrefix = ``

	__DataFlowShape__00000001_.Name = `PP2 T2 to T2`
	__DataFlowShape__00000001_.StartRatio = 0.500000
	__DataFlowShape__00000001_.EndRatio = 0.563197
	__DataFlowShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__DataFlowShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000001_.CornerOffsetRatio = -0.274000
	__DataFlowShape__00000001_.IsHidden = false

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

	__DiagramProcess__00000007_.Name = `Diagram P1`
	__DiagramProcess__00000007_.ComputedPrefix = ``
	__DiagramProcess__00000007_.IsChecked = true
	__DiagramProcess__00000007_.IsEditable_ = true
	__DiagramProcess__00000007_.IsShowPrefix = false
	__DiagramProcess__00000007_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000007_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000007_.Width = 1181.000000
	__DiagramProcess__00000007_.Height = 1350.000000
	__DiagramProcess__00000007_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000007_.IsParticipantsNodeExpanded = true

	__DiagramProcess__00000008_.Name = `DiagramProcess`
	__DiagramProcess__00000008_.ComputedPrefix = ``
	__DiagramProcess__00000008_.IsChecked = false
	__DiagramProcess__00000008_.IsEditable_ = false
	__DiagramProcess__00000008_.IsShowPrefix = false
	__DiagramProcess__00000008_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000008_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000008_.Width = 900.000000
	__DiagramProcess__00000008_.Height = 1350.000000
	__DiagramProcess__00000008_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000008_.IsParticipantsNodeExpanded = false

	__Library__00000000_.Name = `Root`
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsSubLibrariesNodeExpanded = false
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.IsProcessesNodeExpanded = true
	__Library__00000000_.IsDataFlowsNodeExpanded = true
	__Library__00000000_.IsDatasNodeExpanded = false
	__Library__00000000_.IsExpandedTmp = true

	__Library__00000001_.Name = `L1`
	__Library__00000001_.ComputedPrefix = ``
	__Library__00000001_.IsSubLibrariesNodeExpanded = false
	__Library__00000001_.NbPixPerCharacter = 0.000000
	__Library__00000001_.LogoSVGFile = ``
	__Library__00000001_.IsProcessesNodeExpanded = false
	__Library__00000001_.IsDataFlowsNodeExpanded = false
	__Library__00000001_.IsDatasNodeExpanded = true
	__Library__00000001_.IsExpandedTmp = false

	__Library__00000002_.Name = `L2`
	__Library__00000002_.ComputedPrefix = ``
	__Library__00000002_.IsSubLibrariesNodeExpanded = true
	__Library__00000002_.NbPixPerCharacter = 0.000000
	__Library__00000002_.LogoSVGFile = ``
	__Library__00000002_.IsProcessesNodeExpanded = false
	__Library__00000002_.IsDataFlowsNodeExpanded = false
	__Library__00000002_.IsDatasNodeExpanded = false
	__Library__00000002_.IsExpandedTmp = true

	__Participant__00000003_.Name = `PP1`
	__Participant__00000003_.ComputedPrefix = ``
	__Participant__00000003_.IsTasksNodeExpanded = false
	__Participant__00000003_.IsControlFlowsNodeExpanded = false
	__Participant__00000003_.IsDataFlowsNodeExpanded = false

	__Participant__00000004_.Name = `PP2`
	__Participant__00000004_.ComputedPrefix = ``
	__Participant__00000004_.IsTasksNodeExpanded = true
	__Participant__00000004_.IsControlFlowsNodeExpanded = true
	__Participant__00000004_.IsDataFlowsNodeExpanded = false

	__Participant__00000005_.Name = `PP3`
	__Participant__00000005_.ComputedPrefix = ``
	__Participant__00000005_.IsTasksNodeExpanded = false
	__Participant__00000005_.IsControlFlowsNodeExpanded = false
	__Participant__00000005_.IsDataFlowsNodeExpanded = false

	__Participant__00000006_.Name = `PP4`
	__Participant__00000006_.ComputedPrefix = ``
	__Participant__00000006_.IsTasksNodeExpanded = false
	__Participant__00000006_.IsControlFlowsNodeExpanded = false
	__Participant__00000006_.IsDataFlowsNodeExpanded = false

	__ParticipantShape__00000010_.Name = `PP1-DP2`
	__ParticipantShape__00000010_.IsExpanded = false
	__ParticipantShape__00000010_.X = 110.000000
	__ParticipantShape__00000010_.Y = 100.000000
	__ParticipantShape__00000010_.Width = 480.000000
	__ParticipantShape__00000010_.Height = 940.000000
	__ParticipantShape__00000010_.IsHidden = false

	__ParticipantShape__00000012_.Name = `PP2-Diagram P1`
	__ParticipantShape__00000012_.IsExpanded = false
	__ParticipantShape__00000012_.X = 490.500000
	__ParticipantShape__00000012_.Y = 100.000000
	__ParticipantShape__00000012_.Width = 380.500000
	__ParticipantShape__00000012_.Height = 940.000000
	__ParticipantShape__00000012_.IsHidden = false

	__Process__00000005_.Name = `P1`
	__Process__00000005_.ComputedPrefix = ``
	__Process__00000005_.IsSubProcessNodeExpanded = false
	__Process__00000005_.IsDataFlowsNodeExpanded = true

	__Process__00000006_.Name = `P2`
	__Process__00000006_.ComputedPrefix = ``
	__Process__00000006_.IsSubProcessNodeExpanded = false
	__Process__00000006_.IsDataFlowsNodeExpanded = false

	__Process__00000007_.Name = `L2.P1`
	__Process__00000007_.ComputedPrefix = ``
	__Process__00000007_.IsSubProcessNodeExpanded = false
	__Process__00000007_.IsDataFlowsNodeExpanded = false

	__ProcessShape__00000006_.Name = `ProcessShape`
	__ProcessShape__00000006_.IsExpanded = false
	__ProcessShape__00000006_.X = 100.000000
	__ProcessShape__00000006_.Y = 50.000000
	__ProcessShape__00000006_.Width = 500.000000
	__ProcessShape__00000006_.Height = 1000.000000
	__ProcessShape__00000006_.IsHidden = false

	__ProcessShape__00000007_.Name = `ProcessShape`
	__ProcessShape__00000007_.IsExpanded = false
	__ProcessShape__00000007_.X = 100.000000
	__ProcessShape__00000007_.Y = 50.000000
	__ProcessShape__00000007_.Width = 781.000000
	__ProcessShape__00000007_.Height = 1000.000000
	__ProcessShape__00000007_.IsHidden = false

	__ProcessShape__00000008_.Name = `ProcessShape`
	__ProcessShape__00000008_.IsExpanded = false
	__ProcessShape__00000008_.X = 100.000000
	__ProcessShape__00000008_.Y = 50.000000
	__ProcessShape__00000008_.Width = 500.000000
	__ProcessShape__00000008_.Height = 1000.000000
	__ProcessShape__00000008_.IsHidden = false

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
	__TaskShape__00000011_.X = 222.000000
	__TaskShape__00000011_.Y = 236.438108
	__TaskShape__00000011_.Width = 69.000000
	__TaskShape__00000011_.Height = 20.000000
	__TaskShape__00000011_.IsHidden = false

	__TaskShape__00000012_.Name = `T4-DP2`
	__TaskShape__00000012_.IsExpanded = false
	__TaskShape__00000012_.X = 111.000000
	__TaskShape__00000012_.Y = 563.089730
	__TaskShape__00000012_.Width = 250.000000
	__TaskShape__00000012_.Height = 70.000000
	__TaskShape__00000012_.IsHidden = false

	__TaskShape__00000013_.Name = `T2-DP2`
	__TaskShape__00000013_.IsExpanded = false
	__TaskShape__00000013_.X = 111.000000
	__TaskShape__00000013_.Y = 364.607464
	__TaskShape__00000013_.Width = 250.000000
	__TaskShape__00000013_.Height = 70.000000
	__TaskShape__00000013_.IsHidden = false

	__TaskShape__00000014_.Name = `PP2 T2-DP2`
	__TaskShape__00000014_.IsExpanded = false
	__TaskShape__00000014_.X = 529.500000
	__TaskShape__00000014_.Y = 280.080019
	__TaskShape__00000014_.Width = 250.000000
	__TaskShape__00000014_.Height = 70.000000
	__TaskShape__00000014_.IsHidden = false

	__TaskShape__00000015_.Name = `PP2 T1-Diagram P1`
	__TaskShape__00000015_.IsExpanded = false
	__TaskShape__00000015_.X = 595.500000
	__TaskShape__00000015_.Y = 179.528031
	__TaskShape__00000015_.Width = 75.000000
	__TaskShape__00000015_.Height = 20.000000
	__TaskShape__00000015_.IsHidden = false

	// insertion point for setup of pointers
	__ControlFlow__00000003_.Start = __Task__00000004_
	__ControlFlow__00000003_.End = __Task__00000002_
	__ControlFlow__00000004_.Start = __Task__00000002_
	__ControlFlow__00000004_.End = __Task__00000005_
	__ControlFlow__00000005_.Start = __Task__00000003_
	__ControlFlow__00000005_.End = __Task__00000007_
	__ControlFlowShape__00000000_.ControlFlow = __ControlFlow__00000003_
	__ControlFlowShape__00000001_.ControlFlow = __ControlFlow__00000005_
	__DataFlow__00000000_.Start = __Task__00000002_
	__DataFlow__00000000_.End = __Task__00000007_
	__DataFlow__00000001_.Start = __Task__00000007_
	__DataFlow__00000001_.End = __Task__00000002_
	__DataFlow__00000001_.Datas = append(__DataFlow__00000001_.Datas, __Data__00000000_)
	__DataFlow__00000001_.Datas = append(__DataFlow__00000001_.Datas, __Data__00000001_)
	__DataFlowShape__00000001_.DataFlow = __DataFlow__00000001_
	__DiagramProcess__00000006_.Process_Shapes = append(__DiagramProcess__00000006_.Process_Shapes, __ProcessShape__00000006_)
	__DiagramProcess__00000007_.Process_Shapes = append(__DiagramProcess__00000007_.Process_Shapes, __ProcessShape__00000007_)
	__DiagramProcess__00000007_.Participant_Shapes = append(__DiagramProcess__00000007_.Participant_Shapes, __ParticipantShape__00000010_)
	__DiagramProcess__00000007_.Participant_Shapes = append(__DiagramProcess__00000007_.Participant_Shapes, __ParticipantShape__00000012_)
	__DiagramProcess__00000007_.ParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000007_.ParticipantWhoseNodeIsExpanded, __Participant__00000004_)
	__DiagramProcess__00000007_.TasksWhoseNodeIsExpanded = append(__DiagramProcess__00000007_.TasksWhoseNodeIsExpanded, __Task__00000002_)
	__DiagramProcess__00000007_.TasksWhoseNodeIsExpanded = append(__DiagramProcess__00000007_.TasksWhoseNodeIsExpanded, __Task__00000004_)
	__DiagramProcess__00000007_.TasksWhoseNodeIsExpanded = append(__DiagramProcess__00000007_.TasksWhoseNodeIsExpanded, __Task__00000003_)
	__DiagramProcess__00000007_.TasksWhoseNodeIsExpanded = append(__DiagramProcess__00000007_.TasksWhoseNodeIsExpanded, __Task__00000007_)
	__DiagramProcess__00000007_.Task_Shapes = append(__DiagramProcess__00000007_.Task_Shapes, __TaskShape__00000011_)
	__DiagramProcess__00000007_.Task_Shapes = append(__DiagramProcess__00000007_.Task_Shapes, __TaskShape__00000012_)
	__DiagramProcess__00000007_.Task_Shapes = append(__DiagramProcess__00000007_.Task_Shapes, __TaskShape__00000013_)
	__DiagramProcess__00000007_.Task_Shapes = append(__DiagramProcess__00000007_.Task_Shapes, __TaskShape__00000014_)
	__DiagramProcess__00000007_.Task_Shapes = append(__DiagramProcess__00000007_.Task_Shapes, __TaskShape__00000015_)
	__DiagramProcess__00000007_.ControlFlow_Shapes = append(__DiagramProcess__00000007_.ControlFlow_Shapes, __ControlFlowShape__00000000_)
	__DiagramProcess__00000007_.ControlFlow_Shapes = append(__DiagramProcess__00000007_.ControlFlow_Shapes, __ControlFlowShape__00000001_)
	__DiagramProcess__00000007_.DataFlow_Shapes = append(__DiagramProcess__00000007_.DataFlow_Shapes, __DataFlowShape__00000001_)
	__DiagramProcess__00000008_.Process_Shapes = append(__DiagramProcess__00000008_.Process_Shapes, __ProcessShape__00000008_)
	__Library__00000000_.SubLibraries = append(__Library__00000000_.SubLibraries, __Library__00000001_)
	__Library__00000000_.SubLibraries = append(__Library__00000000_.SubLibraries, __Library__00000002_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000005_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000006_)
	__Library__00000000_.ProcesssWhoseNodeIsExpanded = append(__Library__00000000_.ProcesssWhoseNodeIsExpanded, __Process__00000005_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000000_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000001_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000000_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000001_)
	__Library__00000001_.RootDatas = append(__Library__00000001_.RootDatas, __Data__00000002_)
	__Library__00000002_.RootProcesses = append(__Library__00000002_.RootProcesses, __Process__00000007_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000002_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000004_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000005_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000006_)
	__Participant__00000003_.ControlFlows = append(__Participant__00000003_.ControlFlows, __ControlFlow__00000003_)
	__Participant__00000003_.ControlFlows = append(__Participant__00000003_.ControlFlows, __ControlFlow__00000004_)
	__Participant__00000003_.TaskWhoseOutControlFlowsNodeIsExpanded = append(__Participant__00000003_.TaskWhoseOutControlFlowsNodeIsExpanded, __Task__00000002_)
	__Participant__00000003_.TaskWhoseOutControlFlowsNodeIsExpanded = append(__Participant__00000003_.TaskWhoseOutControlFlowsNodeIsExpanded, __Task__00000004_)
	__Participant__00000003_.TaskWhoseInControlFlowsNodeIsExpanded = append(__Participant__00000003_.TaskWhoseInControlFlowsNodeIsExpanded, __Task__00000002_)
	__Participant__00000004_.Tasks = append(__Participant__00000004_.Tasks, __Task__00000003_)
	__Participant__00000004_.Tasks = append(__Participant__00000004_.Tasks, __Task__00000007_)
	__Participant__00000004_.ControlFlows = append(__Participant__00000004_.ControlFlows, __ControlFlow__00000005_)
	__Participant__00000004_.TaskWhoseInControlFlowsNodeIsExpanded = append(__Participant__00000004_.TaskWhoseInControlFlowsNodeIsExpanded, __Task__00000007_)
	__Participant__00000004_.TaskWhoseOutDataFlowsNodeIsExpanded = append(__Participant__00000004_.TaskWhoseOutDataFlowsNodeIsExpanded, __Task__00000007_)
	__Participant__00000004_.TaskWhoseInDataFlowsNodeIsExpanded = append(__Participant__00000004_.TaskWhoseInDataFlowsNodeIsExpanded, __Task__00000007_)
	__ParticipantShape__00000010_.Participant = __Participant__00000003_
	__ParticipantShape__00000012_.Participant = __Participant__00000004_
	__Process__00000005_.DiagramProcesss = append(__Process__00000005_.DiagramProcesss, __DiagramProcess__00000007_)
	__Process__00000005_.DiagramProcessWhoseNodeIsExpanded = append(__Process__00000005_.DiagramProcessWhoseNodeIsExpanded, __DiagramProcess__00000007_)
	__Process__00000005_.Participants = append(__Process__00000005_.Participants, __Participant__00000003_)
	__Process__00000005_.Participants = append(__Process__00000005_.Participants, __Participant__00000004_)
	__Process__00000005_.Participants = append(__Process__00000005_.Participants, __Participant__00000005_)
	__Process__00000005_.Participants = append(__Process__00000005_.Participants, __Participant__00000006_)
	__Process__00000005_.ParticipantWhoseNodeIsExpanded = append(__Process__00000005_.ParticipantWhoseNodeIsExpanded, __Participant__00000003_)
	__Process__00000005_.DataFlows = append(__Process__00000005_.DataFlows, __DataFlow__00000000_)
	__Process__00000005_.DataFlows = append(__Process__00000005_.DataFlows, __DataFlow__00000001_)
	__Process__00000006_.DiagramProcesss = append(__Process__00000006_.DiagramProcesss, __DiagramProcess__00000006_)
	__Process__00000007_.DiagramProcesss = append(__Process__00000007_.DiagramProcesss, __DiagramProcess__00000008_)
	__ProcessShape__00000006_.Process = __Process__00000006_
	__ProcessShape__00000007_.Process = __Process__00000005_
	__ProcessShape__00000008_.Process = __Process__00000007_
	__TaskShape__00000011_.Task = __Task__00000004_
	__TaskShape__00000012_.Task = __Task__00000006_
	__TaskShape__00000013_.Task = __Task__00000002_
	__TaskShape__00000014_.Task = __Task__00000007_
	__TaskShape__00000015_.Task = __Task__00000003_
}
