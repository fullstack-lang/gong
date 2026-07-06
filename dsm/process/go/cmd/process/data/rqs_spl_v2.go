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

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__AllocatedResourceShape__00000000_ := (&models.AllocatedResourceShape{Name: `DiagramProcess-ATS Unit-ATC`}).Stage(stage)
	__AllocatedResourceShape__00000001_ := (&models.AllocatedResourceShape{Name: `DiagramProcess-EUROCONTROL IFPS-IFPS`}).Stage(stage)
	__AllocatedResourceShape__00000002_ := (&models.AllocatedResourceShape{Name: `DiagramProcess-Aircraft Operator-AO`}).Stage(stage)

	__ControlFlow__00000000_ := (&models.ControlFlow{Name: `"" to "Transmits RQS"`}).Stage(stage)
	__ControlFlow__00000001_ := (&models.ControlFlow{Name: `"Receives SPL" to ""`}).Stage(stage)

	__ControlFlowShape__00000004_ := (&models.ControlFlowShape{Name: `"Receives SPL" to ""`}).Stage(stage)
	__ControlFlowShape__00000005_ := (&models.ControlFlowShape{Name: `"" to "Transmits RQS"`}).Stage(stage)

	__Data__00000000_ := (&models.Data{Name: `RQS`}).Stage(stage)
	__Data__00000001_ := (&models.Data{Name: `SPL`}).Stage(stage)

	__DataFlow__00000000_ := (&models.DataFlow{Name: `"Transmits RQS" to "Validates and Transmit SPL"`}).Stage(stage)
	__DataFlow__00000003_ := (&models.DataFlow{Name: `"Validates and Transmit SPL" to "Receives SPL"`}).Stage(stage)

	__DataFlowShape__00000004_ := (&models.DataFlowShape{Name: `"Transmits RQS" to "Validates and Transmit SPL"`}).Stage(stage)
	__DataFlowShape__00000005_ := (&models.DataFlowShape{Name: `"Validates and Transmit SPL" to "Receives SPL"`}).Stage(stage)

	__DataShape__00000001_ := (&models.DataShape{Name: `"Validates and Routes RQS" to "Compiles and Transmits SPL" - RQS`}).Stage(stage)
	__DataShape__00000002_ := (&models.DataShape{Name: `"Compiles and Transmits SPL" to "Validates and Routes SPL" - SPL`}).Stage(stage)
	__DataShape__00000003_ := (&models.DataShape{Name: `"Validates and Routes SPL" to "Receives SPL" - SPL`}).Stage(stage)
	__DataShape__00000004_ := (&models.DataShape{Name: `"Transmits RQS" to "Validates and Routes RQS"-"Transmits RQS" to "Validates and Routes RQS"-DiagramProcess`}).Stage(stage)

	__DiagramProcess__00000000_ := (&models.DiagramProcess{Name: `RQS SPL, AO out of the loop`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `IFPS Operations`}).Stage(stage)

	__Participant__00000000_ := (&models.Participant{Name: `ATS Unit`}).Stage(stage)
	__Participant__00000001_ := (&models.Participant{Name: `EUROCONTROL IFPS`}).Stage(stage)

	__ParticipantShape__00000000_ := (&models.ParticipantShape{Name: `ATS Unit-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000001_ := (&models.ParticipantShape{Name: `EUROCONTROL IFPS-DiagramProcess`}).Stage(stage)

	__Process__00000000_ := (&models.Process{Name: `RQS/SPL Exchange (AO out of the loop)`}).Stage(stage)

	__ProcessShape__00000000_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)

	__Resource__00000000_ := (&models.Resource{Name: `ATC`}).Stage(stage)
	__Resource__00000001_ := (&models.Resource{Name: `IFPS`}).Stage(stage)
	__Resource__00000002_ := (&models.Resource{Name: `AO`}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `Transmits RQS`}).Stage(stage)
	__Task__00000001_ := (&models.Task{Name: `Validates and Transmit SPL`}).Stage(stage)
	__Task__00000004_ := (&models.Task{Name: `Receives SPL`}).Stage(stage)
	__Task__00000005_ := (&models.Task{Name: ``}).Stage(stage)
	__Task__00000006_ := (&models.Task{Name: ``}).Stage(stage)

	__TaskShape__00000000_ := (&models.TaskShape{Name: `Transmits RQS-DiagramProcess`}).Stage(stage)
	__TaskShape__00000001_ := (&models.TaskShape{Name: `Validates and Routes RQS-DiagramProcess`}).Stage(stage)
	__TaskShape__00000004_ := (&models.TaskShape{Name: `Receives SPL-DiagramProcess`}).Stage(stage)
	__TaskShape__00000005_ := (&models.TaskShape{Name: `-DiagramProcess-Start`}).Stage(stage)
	__TaskShape__00000006_ := (&models.TaskShape{Name: `-DiagramProcess-End`}).Stage(stage)

	// insertion point for initialization of values

	__AllocatedResourceShape__00000000_.Name = `DiagramProcess-ATS Unit-ATC`

	__AllocatedResourceShape__00000001_.Name = `DiagramProcess-EUROCONTROL IFPS-IFPS`

	__AllocatedResourceShape__00000002_.Name = `DiagramProcess-Aircraft Operator-AO`

	__ControlFlow__00000000_.Name = `"" to "Transmits RQS"`
	__ControlFlow__00000000_.Description = ``
	__ControlFlow__00000000_.ComputedPrefix = ``

	__ControlFlow__00000001_.Name = `"Receives SPL" to ""`
	__ControlFlow__00000001_.Description = ``
	__ControlFlow__00000001_.ComputedPrefix = ``

	__ControlFlowShape__00000004_.Name = `"Receives SPL" to ""`
	__ControlFlowShape__00000004_.StartRatio = 0.500000
	__ControlFlowShape__00000004_.EndRatio = 0.500000
	__ControlFlowShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000004_.CornerOffsetRatio = 1.800098
	__ControlFlowShape__00000004_.IsHidden = false

	__ControlFlowShape__00000005_.Name = `"" to "Transmits RQS"`
	__ControlFlowShape__00000005_.StartRatio = 1.000000
	__ControlFlowShape__00000005_.EndRatio = 0.500000
	__ControlFlowShape__00000005_.StartOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000005_.EndOrientation = models.ORIENTATION_VERTICAL
	__ControlFlowShape__00000005_.CornerOffsetRatio = 2.900293
	__ControlFlowShape__00000005_.IsHidden = false

	__Data__00000000_.Name = `RQS`
	__Data__00000000_.Acronym = ``
	__Data__00000000_.Description = `Request Supplementary Flight Plan Message`
	__Data__00000000_.ComputedPrefix = ``
	__Data__00000000_.SVG_Path = ``
	__Data__00000000_.InverseAppliedScaling = 0.000000

	__Data__00000001_.Name = `SPL`
	__Data__00000001_.Acronym = ``
	__Data__00000001_.Description = `Supplementary Flight Plan Message (Item 19 data)`
	__Data__00000001_.ComputedPrefix = ``
	__Data__00000001_.SVG_Path = ``
	__Data__00000001_.InverseAppliedScaling = 0.000000

	__DataFlow__00000000_.Name = `"Transmits RQS" to "Validates and Transmit SPL"`
	__DataFlow__00000000_.Description = ``
	__DataFlow__00000000_.ComputedPrefix = ``
	__DataFlow__00000000_.Type = models.DataFlow_Task2Task
	__DataFlow__00000000_.IsDatasNodeExpanded = true

	__DataFlow__00000003_.Name = `"Validates and Transmit SPL" to "Receives SPL"`
	__DataFlow__00000003_.Description = ``
	__DataFlow__00000003_.ComputedPrefix = ``
	__DataFlow__00000003_.Type = models.DataFlow_Task2Task
	__DataFlow__00000003_.IsDatasNodeExpanded = false

	__DataFlowShape__00000004_.Name = `"Transmits RQS" to "Validates and Transmit SPL"`
	__DataFlowShape__00000004_.StartRatio = 0.500000
	__DataFlowShape__00000004_.EndRatio = 0.221653
	__DataFlowShape__00000004_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000004_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000004_.CornerOffsetRatio = 1.534805
	__DataFlowShape__00000004_.IsHidden = false

	__DataFlowShape__00000005_.Name = `"Validates and Transmit SPL" to "Receives SPL"`
	__DataFlowShape__00000005_.StartRatio = 0.713545
	__DataFlowShape__00000005_.EndRatio = 0.483431
	__DataFlowShape__00000005_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000005_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000005_.CornerOffsetRatio = -1.000195
	__DataFlowShape__00000005_.IsHidden = false

	__DataShape__00000001_.Name = `"Validates and Routes RQS" to "Compiles and Transmits SPL" - RQS`

	__DataShape__00000002_.Name = `"Compiles and Transmits SPL" to "Validates and Routes SPL" - SPL`

	__DataShape__00000003_.Name = `"Validates and Routes SPL" to "Receives SPL" - SPL`

	__DataShape__00000004_.Name = `"Transmits RQS" to "Validates and Routes RQS"-"Transmits RQS" to "Validates and Routes RQS"-DiagramProcess`

	__DiagramProcess__00000000_.Name = `RQS SPL, AO out of the loop`
	__DiagramProcess__00000000_.Description = ``
	__DiagramProcess__00000000_.ComputedPrefix = ``
	__DiagramProcess__00000000_.IsChecked = true
	__DiagramProcess__00000000_.IsEditable_ = true
	__DiagramProcess__00000000_.IsShowPrefix = false
	__DiagramProcess__00000000_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000000_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000000_.Width = 1225.000000
	__DiagramProcess__00000000_.Height = 1243.000000
	__DiagramProcess__00000000_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000000_.IsParticipantsNodeExpanded = true
	__DiagramProcess__00000000_.IsExternalParticipantsNodeExpanded = false
	__DiagramProcess__00000000_.IsNotesNodeExpanded = false

	__Library__00000000_.Name = `IFPS Operations`
	__Library__00000000_.Description = `Modeling of message exchanges handled by EUROCONTROL IFPS`
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsSubLibrariesNodeExpanded = false
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.IsProcessesNodeExpanded = false
	__Library__00000000_.IsDataFlowsNodeExpanded = false
	__Library__00000000_.IsDatasNodeExpanded = false
	__Library__00000000_.IsResourcesNodeExpanded = false
	__Library__00000000_.IsNotesNodeExpanded = false
	__Library__00000000_.IsExpandedTmp = true

	__Participant__00000000_.Name = `ATS Unit`
	__Participant__00000000_.IsProcessResource = false
	__Participant__00000000_.Description = `Air Traffic Services Unit / Rescue Coordination Center`
	__Participant__00000000_.IsResourcesNodeExpanded = false
	__Participant__00000000_.IsProcessesNodeExpanded = false
	__Participant__00000000_.ComputedPrefix = ``
	__Participant__00000000_.IsTasksNodeExpanded = false
	__Participant__00000000_.IsControlFlowsNodeExpanded = true
	__Participant__00000000_.IsDataFlowsNodeExpanded = false

	__Participant__00000001_.Name = `EUROCONTROL IFPS`
	__Participant__00000001_.IsProcessResource = false
	__Participant__00000001_.Description = `Initial Flight Plan Processing System`
	__Participant__00000001_.IsResourcesNodeExpanded = false
	__Participant__00000001_.IsProcessesNodeExpanded = false
	__Participant__00000001_.ComputedPrefix = ``
	__Participant__00000001_.IsTasksNodeExpanded = true
	__Participant__00000001_.IsControlFlowsNodeExpanded = false
	__Participant__00000001_.IsDataFlowsNodeExpanded = false

	__ParticipantShape__00000000_.Name = `ATS Unit-DiagramProcess`
	__ParticipantShape__00000000_.IsExpanded = false
	__ParticipantShape__00000000_.X = 100.000000
	__ParticipantShape__00000000_.Y = 100.000000
	__ParticipantShape__00000000_.Width = 250.000000
	__ParticipantShape__00000000_.Height = 70.000000
	__ParticipantShape__00000000_.IsHidden = false
	__ParticipantShape__00000000_.WidthWeight = 1.000000

	__ParticipantShape__00000001_.Name = `EUROCONTROL IFPS-DiagramProcess`
	__ParticipantShape__00000001_.IsExpanded = false
	__ParticipantShape__00000001_.X = 450.000000
	__ParticipantShape__00000001_.Y = 100.000000
	__ParticipantShape__00000001_.Width = 250.000000
	__ParticipantShape__00000001_.Height = 70.000000
	__ParticipantShape__00000001_.IsHidden = false
	__ParticipantShape__00000001_.WidthWeight = 1.000000

	__Process__00000000_.Name = `RQS/SPL Exchange (AO out of the loop)`
	__Process__00000000_.Description = `Process for requesting and delivering supplementary flight plan data (Item 19).`
	__Process__00000000_.ComputedPrefix = ``
	__Process__00000000_.SVG_Path = ``
	__Process__00000000_.InverseAppliedScaling = 0.000000
	__Process__00000000_.IsSubProcessNodeExpanded = false
	__Process__00000000_.IsDataFlowsNodeExpanded = false

	__ProcessShape__00000000_.Name = `ProcessShape`
	__ProcessShape__00000000_.IsExpanded = false
	__ProcessShape__00000000_.X = 47.000000
	__ProcessShape__00000000_.Y = 43.000000
	__ProcessShape__00000000_.Width = 878.000000
	__ProcessShape__00000000_.Height = 900.000000
	__ProcessShape__00000000_.IsHidden = false

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

	__Task__00000004_.Name = `Receives SPL`
	__Task__00000004_.Description = ``
	__Task__00000004_.ComputedPrefix = ``
	__Task__00000004_.IsStartTask = false
	__Task__00000004_.IsEndTask = false
	__Task__00000004_.IsTaskNameNotProcessName = false

	__Task__00000005_.Name = ``
	__Task__00000005_.Description = ``
	__Task__00000005_.ComputedPrefix = ``
	__Task__00000005_.IsStartTask = true
	__Task__00000005_.IsEndTask = false
	__Task__00000005_.IsTaskNameNotProcessName = false

	__Task__00000006_.Name = ``
	__Task__00000006_.Description = ``
	__Task__00000006_.ComputedPrefix = ``
	__Task__00000006_.IsStartTask = false
	__Task__00000006_.IsEndTask = true
	__Task__00000006_.IsTaskNameNotProcessName = false

	__TaskShape__00000000_.Name = `Transmits RQS-DiagramProcess`
	__TaskShape__00000000_.IsExpanded = false
	__TaskShape__00000000_.X = 119.000000
	__TaskShape__00000000_.Y = 304.000000
	__TaskShape__00000000_.Width = 200.000000
	__TaskShape__00000000_.Height = 60.000000
	__TaskShape__00000000_.IsHidden = false

	__TaskShape__00000001_.Name = `Validates and Routes RQS-DiagramProcess`
	__TaskShape__00000001_.IsExpanded = false
	__TaskShape__00000001_.X = 622.000000
	__TaskShape__00000001_.Y = 446.000000
	__TaskShape__00000001_.Width = 200.000000
	__TaskShape__00000001_.Height = 185.000000
	__TaskShape__00000001_.IsHidden = false

	__TaskShape__00000004_.Name = `Receives SPL-DiagramProcess`
	__TaskShape__00000004_.IsExpanded = false
	__TaskShape__00000004_.X = 120.000000
	__TaskShape__00000004_.Y = 652.000000
	__TaskShape__00000004_.Width = 200.000000
	__TaskShape__00000004_.Height = 60.000000
	__TaskShape__00000004_.IsHidden = false

	__TaskShape__00000005_.Name = `-DiagramProcess-Start`
	__TaskShape__00000005_.IsExpanded = false
	__TaskShape__00000005_.X = 174.000000
	__TaskShape__00000005_.Y = 204.000000
	__TaskShape__00000005_.Width = 50.000000
	__TaskShape__00000005_.Height = 20.000000
	__TaskShape__00000005_.IsHidden = false

	__TaskShape__00000006_.Name = `-DiagramProcess-End`
	__TaskShape__00000006_.IsExpanded = false
	__TaskShape__00000006_.X = 189.000000
	__TaskShape__00000006_.Y = 839.000000
	__TaskShape__00000006_.Width = 50.000000
	__TaskShape__00000006_.Height = 36.000000
	__TaskShape__00000006_.IsHidden = false

	// insertion point for setup of pointers
	__AllocatedResourceShape__00000000_.Participant = __Participant__00000000_
	__AllocatedResourceShape__00000000_.Resource = __Resource__00000000_
	__AllocatedResourceShape__00000001_.Participant = __Participant__00000001_
	__AllocatedResourceShape__00000001_.Resource = __Resource__00000001_
	__AllocatedResourceShape__00000002_.Participant = nil
	__AllocatedResourceShape__00000002_.Resource = __Resource__00000002_
	__ControlFlow__00000000_.Start = __Task__00000005_
	__ControlFlow__00000000_.End = __Task__00000000_
	__ControlFlow__00000001_.Start = __Task__00000004_
	__ControlFlow__00000001_.End = __Task__00000006_
	__ControlFlowShape__00000004_.ControlFlow = __ControlFlow__00000001_
	__ControlFlowShape__00000005_.ControlFlow = __ControlFlow__00000000_
	__DataFlow__00000000_.Datas = append(__DataFlow__00000000_.Datas, __Data__00000000_)
	__DataFlow__00000000_.StartTask = __Task__00000000_
	__DataFlow__00000000_.EndTask = __Task__00000001_
	__DataFlow__00000000_.StartExternalParticipant = nil
	__DataFlow__00000000_.EndExternalParticipant = nil
	__DataFlow__00000003_.Datas = append(__DataFlow__00000003_.Datas, __Data__00000001_)
	__DataFlow__00000003_.StartTask = __Task__00000001_
	__DataFlow__00000003_.EndTask = __Task__00000004_
	__DataFlow__00000003_.StartExternalParticipant = nil
	__DataFlow__00000003_.EndExternalParticipant = nil
	__DataFlowShape__00000004_.DataFlow = __DataFlow__00000000_
	__DataFlowShape__00000005_.DataFlow = __DataFlow__00000003_
	__DataShape__00000001_.Data = __Data__00000000_
	__DataShape__00000001_.DataFlow = nil
	__DataShape__00000002_.Data = __Data__00000001_
	__DataShape__00000002_.DataFlow = nil
	__DataShape__00000003_.Data = __Data__00000001_
	__DataShape__00000003_.DataFlow = __DataFlow__00000003_
	__DataShape__00000004_.Data = __Data__00000000_
	__DataShape__00000004_.DataFlow = __DataFlow__00000000_
	__DiagramProcess__00000000_.Process_Shapes = append(__DiagramProcess__00000000_.Process_Shapes, __ProcessShape__00000000_)
	__DiagramProcess__00000000_.Participant_Shapes = append(__DiagramProcess__00000000_.Participant_Shapes, __ParticipantShape__00000000_)
	__DiagramProcess__00000000_.Participant_Shapes = append(__DiagramProcess__00000000_.Participant_Shapes, __ParticipantShape__00000001_)
	__DiagramProcess__00000000_.TasksWhoseNodeIsExpanded = append(__DiagramProcess__00000000_.TasksWhoseNodeIsExpanded, __Task__00000001_)
	__DiagramProcess__00000000_.Task_Shapes = append(__DiagramProcess__00000000_.Task_Shapes, __TaskShape__00000000_)
	__DiagramProcess__00000000_.Task_Shapes = append(__DiagramProcess__00000000_.Task_Shapes, __TaskShape__00000001_)
	__DiagramProcess__00000000_.Task_Shapes = append(__DiagramProcess__00000000_.Task_Shapes, __TaskShape__00000004_)
	__DiagramProcess__00000000_.Task_Shapes = append(__DiagramProcess__00000000_.Task_Shapes, __TaskShape__00000005_)
	__DiagramProcess__00000000_.Task_Shapes = append(__DiagramProcess__00000000_.Task_Shapes, __TaskShape__00000006_)
	__DiagramProcess__00000000_.ControlFlow_Shapes = append(__DiagramProcess__00000000_.ControlFlow_Shapes, __ControlFlowShape__00000004_)
	__DiagramProcess__00000000_.ControlFlow_Shapes = append(__DiagramProcess__00000000_.ControlFlow_Shapes, __ControlFlowShape__00000005_)
	__DiagramProcess__00000000_.DataFlow_Shapes = append(__DiagramProcess__00000000_.DataFlow_Shapes, __DataFlowShape__00000004_)
	__DiagramProcess__00000000_.DataFlow_Shapes = append(__DiagramProcess__00000000_.DataFlow_Shapes, __DataFlowShape__00000005_)
	__DiagramProcess__00000000_.Data_Shapes = append(__DiagramProcess__00000000_.Data_Shapes, __DataShape__00000001_)
	__DiagramProcess__00000000_.Data_Shapes = append(__DiagramProcess__00000000_.Data_Shapes, __DataShape__00000002_)
	__DiagramProcess__00000000_.Data_Shapes = append(__DiagramProcess__00000000_.Data_Shapes, __DataShape__00000003_)
	__DiagramProcess__00000000_.Data_Shapes = append(__DiagramProcess__00000000_.Data_Shapes, __DataShape__00000004_)
	__DiagramProcess__00000000_.DataFlowsWhoseDataNodeIsExpanded = append(__DiagramProcess__00000000_.DataFlowsWhoseDataNodeIsExpanded, __DataFlow__00000000_)
	__DiagramProcess__00000000_.AllocatedResourceShapes = append(__DiagramProcess__00000000_.AllocatedResourceShapes, __AllocatedResourceShape__00000000_)
	__DiagramProcess__00000000_.AllocatedResourceShapes = append(__DiagramProcess__00000000_.AllocatedResourceShapes, __AllocatedResourceShape__00000001_)
	__DiagramProcess__00000000_.AllocatedResourceShapes = append(__DiagramProcess__00000000_.AllocatedResourceShapes, __AllocatedResourceShape__00000002_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000000_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000000_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000003_)
	__Library__00000000_.DataFlowsWhoseNodeIsExpanded = append(__Library__00000000_.DataFlowsWhoseNodeIsExpanded, __DataFlow__00000000_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000000_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000001_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000000_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000001_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000002_)
	__Participant__00000000_.Resources = append(__Participant__00000000_.Resources, __Resource__00000000_)
	__Participant__00000000_.Tasks = append(__Participant__00000000_.Tasks, __Task__00000000_)
	__Participant__00000000_.Tasks = append(__Participant__00000000_.Tasks, __Task__00000004_)
	__Participant__00000000_.Tasks = append(__Participant__00000000_.Tasks, __Task__00000005_)
	__Participant__00000000_.Tasks = append(__Participant__00000000_.Tasks, __Task__00000006_)
	__Participant__00000000_.ControlFlows = append(__Participant__00000000_.ControlFlows, __ControlFlow__00000000_)
	__Participant__00000000_.ControlFlows = append(__Participant__00000000_.ControlFlows, __ControlFlow__00000001_)
	__Participant__00000000_.TaskWhoseOutControlFlowsNodeIsExpanded = append(__Participant__00000000_.TaskWhoseOutControlFlowsNodeIsExpanded, __Task__00000005_)
	__Participant__00000000_.TaskWhoseOutControlFlowsNodeIsExpanded = append(__Participant__00000000_.TaskWhoseOutControlFlowsNodeIsExpanded, __Task__00000004_)
	__Participant__00000000_.TaskWhoseInControlFlowsNodeIsExpanded = append(__Participant__00000000_.TaskWhoseInControlFlowsNodeIsExpanded, __Task__00000006_)
	__Participant__00000000_.TaskWhoseInControlFlowsNodeIsExpanded = append(__Participant__00000000_.TaskWhoseInControlFlowsNodeIsExpanded, __Task__00000000_)
	__Participant__00000000_.TaskWhoseOutDataFlowsNodeIsExpanded = append(__Participant__00000000_.TaskWhoseOutDataFlowsNodeIsExpanded, __Task__00000000_)
	__Participant__00000001_.Resources = append(__Participant__00000001_.Resources, __Resource__00000001_)
	__Participant__00000001_.Tasks = append(__Participant__00000001_.Tasks, __Task__00000001_)
	__Participant__00000001_.TaskWhoseOutDataFlowsNodeIsExpanded = append(__Participant__00000001_.TaskWhoseOutDataFlowsNodeIsExpanded, __Task__00000001_)
	__Participant__00000001_.TaskWhoseInDataFlowsNodeIsExpanded = append(__Participant__00000001_.TaskWhoseInDataFlowsNodeIsExpanded, __Task__00000001_)
	__ParticipantShape__00000000_.Participant = __Participant__00000000_
	__ParticipantShape__00000001_.Participant = __Participant__00000001_
	__Process__00000000_.DiagramProcesss = append(__Process__00000000_.DiagramProcesss, __DiagramProcess__00000000_)
	__Process__00000000_.Participants = append(__Process__00000000_.Participants, __Participant__00000000_)
	__Process__00000000_.Participants = append(__Process__00000000_.Participants, __Participant__00000001_)
	__Process__00000000_.DataFlows = append(__Process__00000000_.DataFlows, __DataFlow__00000000_)
	__Process__00000000_.DataFlows = append(__Process__00000000_.DataFlows, __DataFlow__00000003_)
	__ProcessShape__00000000_.Process = __Process__00000000_
	__Task__00000000_.Type = nil
	__Task__00000001_.Type = nil
	__Task__00000004_.Type = nil
	__Task__00000005_.Type = nil
	__Task__00000006_.Type = nil
	__TaskShape__00000000_.Task = __Task__00000000_
	__TaskShape__00000001_.Task = __Task__00000001_
	__TaskShape__00000004_.Task = __Task__00000004_
	__TaskShape__00000005_.Task = __Task__00000005_
	__TaskShape__00000006_.Task = __Task__00000006_
}
