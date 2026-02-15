package main

import (
	"time"

	"github.com/fullstack-lang/gong/dsme/statemachines/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__Action__00000000_ := (&models.Action{Name: `Lock`}).Stage(stage)
	__Action__00000001_ := (&models.Action{Name: `Unlock`}).Stage(stage)
	__Action__00000002_ := (&models.Action{Name: `RLock`}).Stage(stage)
	__Action__00000003_ := (&models.Action{Name: `Runlock`}).Stage(stage)

	__Activities__00000000_ := (&models.Activities{Name: `Test`}).Stage(stage)

	__Architecture__00000000_ := (&models.Architecture{Name: `Gong UX loop Architecture`}).Stage(stage)

	__Diagram__00000000_ := (&models.Diagram{Name: `UX Loop Diagram 2`}).Stage(stage)

	__Guard__00000000_ := (&models.Guard{Name: `yes`}).Stage(stage)
	__Guard__00000001_ := (&models.Guard{Name: `no`}).Stage(stage)
	__Guard__00000002_ := (&models.Guard{Name: `a suppress action`}).Stage(stage)
	__Guard__00000003_ := (&models.Guard{Name: `not a suppress action`}).Stage(stage)

	__State__00000000_ := (&models.State{Name: `Probe Form`}).Stage(stage)
	__State__00000001_ := (&models.State{Name: `Probe Tree`}).Stage(stage)
	__State__00000002_ := (&models.State{Name: `Probe Table`}).Stage(stage)
	__State__00000003_ := (&models.State{Name: `Form - Update Stage Of Interest`}).Stage(stage)
	__State__00000004_ := (&models.State{Name: `Enforce Model Semantic`}).Stage(stage)
	__State__00000005_ := (&models.State{Name: `Update Probe Form`}).Stage(stage)
	__State__00000006_ := (&models.State{Name: `Update Probe Table`}).Stage(stage)
	__State__00000007_ := (&models.State{Name: `Update Probe Tree`}).Stage(stage)
	__State__00000008_ := (&models.State{Name: `Load Stage`}).Stage(stage)
	__State__00000009_ := (&models.State{Name: `Persist Stage`}).Stage(stage)
	__State__00000010_ := (&models.State{Name: `Update UX ?`}).Stage(stage)
	__State__00000011_ := (&models.State{Name: `Update UX ?`}).Stage(stage)
	__State__00000012_ := (&models.State{Name: `Clean Stage`}).Stage(stage)

	__StateMachine__00000000_ := (&models.StateMachine{Name: `UX Loop`}).Stage(stage)

	__StateShape__00000000_ := (&models.StateShape{Name: `Probe Form`}).Stage(stage)
	__StateShape__00000001_ := (&models.StateShape{Name: `Probe Tree`}).Stage(stage)
	__StateShape__00000002_ := (&models.StateShape{Name: `Probe Table`}).Stage(stage)
	__StateShape__00000003_ := (&models.StateShape{Name: `Form - Update Stage Of Interest`}).Stage(stage)
	__StateShape__00000004_ := (&models.StateShape{Name: `Enforce Model Semantic`}).Stage(stage)
	__StateShape__00000005_ := (&models.StateShape{Name: `Update Probe Form`}).Stage(stage)
	__StateShape__00000006_ := (&models.StateShape{Name: `Update Probe Table`}).Stage(stage)
	__StateShape__00000007_ := (&models.StateShape{Name: `Update Probe Tree`}).Stage(stage)
	__StateShape__00000008_ := (&models.StateShape{Name: `Load Stage`}).Stage(stage)
	__StateShape__00000009_ := (&models.StateShape{Name: `Persist Stage`}).Stage(stage)
	__StateShape__00000010_ := (&models.StateShape{Name: `Update UX ?`}).Stage(stage)
	__StateShape__00000011_ := (&models.StateShape{Name: `Update UX ?`}).Stage(stage)
	__StateShape__00000012_ := (&models.StateShape{Name: `Clean Stage`}).Stage(stage)

	__Transition__00000000_ := (&models.Transition{Name: `Submit`}).Stage(stage)
	__Transition__00000001_ := (&models.Transition{Name: ``}).Stage(stage)
	__Transition__00000002_ := (&models.Transition{Name: `/Model Updated`}).Stage(stage)
	__Transition__00000003_ := (&models.Transition{Name: ``}).Stage(stage)
	__Transition__00000004_ := (&models.Transition{Name: ``}).Stage(stage)
	__Transition__00000005_ := (&models.Transition{Name: `/Stage loaded`}).Stage(stage)
	__Transition__00000006_ := (&models.Transition{Name: `/Stage persisted`}).Stage(stage)
	__Transition__00000007_ := (&models.Transition{Name: `Yes`}).Stage(stage)
	__Transition__00000008_ := (&models.Transition{Name: `Yes`}).Stage(stage)
	__Transition__00000009_ := (&models.Transition{Name: `Form - Update Stage Of Interest to Clean Stage`}).Stage(stage)
	__Transition__00000010_ := (&models.Transition{Name: ``}).Stage(stage)
	__Transition__00000011_ := (&models.Transition{Name: `Stage Modified ? to Probe Form`}).Stage(stage)

	__Transition_Shape__00000000_ := (&models.Transition_Shape{Name: `UX 1 - Waiting for User Input to UX 1 - Update Stage`}).Stage(stage)
	__Transition_Shape__00000001_ := (&models.Transition_Shape{Name: `UX 1 - Update Stage to Enforce Model Semantic`}).Stage(stage)
	__Transition_Shape__00000002_ := (&models.Transition_Shape{Name: `Enforce Model Semantic to UX 1 - Update UX`}).Stage(stage)
	__Transition_Shape__00000003_ := (&models.Transition_Shape{Name: `Enforce Model Semantic to UX 2 - Update UX`}).Stage(stage)
	__Transition_Shape__00000004_ := (&models.Transition_Shape{Name: `Enforce Model Semantic to UX n - Update UX`}).Stage(stage)
	__Transition_Shape__00000005_ := (&models.Transition_Shape{Name: `Load Stage to Enforce Model Semantic`}).Stage(stage)
	__Transition_Shape__00000006_ := (&models.Transition_Shape{Name: `Persist Stage to UX 1 - Update UX`}).Stage(stage)
	__Transition_Shape__00000007_ := (&models.Transition_Shape{Name: `Stage Modified ? to Persist Stage`}).Stage(stage)
	__Transition_Shape__00000008_ := (&models.Transition_Shape{Name: `Update UX ? to UX 1 - Update UX`}).Stage(stage)
	__Transition_Shape__00000009_ := (&models.Transition_Shape{Name: `Form - Update Stage Of Interest to Clean Stage`}).Stage(stage)
	__Transition_Shape__00000010_ := (&models.Transition_Shape{Name: `Clean Stage to Enforce Model Semantic`}).Stage(stage)
	__Transition_Shape__00000011_ := (&models.Transition_Shape{Name: `Stage Modified ? to Probe Form`}).Stage(stage)

	// insertion point for initialization of values

	__Action__00000000_.Name = `Lock`
	__Action__00000000_.Criticality = models.CriticalityCritical

	__Action__00000001_.Name = `Unlock`
	__Action__00000001_.Criticality = models.CriticalityCritical

	__Action__00000002_.Name = `RLock`
	__Action__00000002_.Criticality = models.CriticalityDefault

	__Action__00000003_.Name = `Runlock`
	__Action__00000003_.Criticality = models.CriticalityDefault

	__Activities__00000000_.Name = `Test`
	__Activities__00000000_.Criticality = ""

	__Architecture__00000000_.Name = `Gong UX loop Architecture`
	__Architecture__00000000_.NbPixPerCharacter = 8.000000

	__Diagram__00000000_.Name = `UX Loop Diagram 2`
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsInRenameMode = false

	__Guard__00000000_.Name = `yes`

	__Guard__00000001_.Name = `no`

	__Guard__00000002_.Name = `a suppress action`

	__Guard__00000003_.Name = `not a suppress action`

	__State__00000000_.Name = `Probe Form`
	__State__00000000_.IsDecisionNode = false
	__State__00000000_.IsFictif = false
	__State__00000000_.IsEndState = false

	__State__00000001_.Name = `Probe Tree`
	__State__00000001_.IsDecisionNode = false
	__State__00000001_.IsFictif = false
	__State__00000001_.IsEndState = false

	__State__00000002_.Name = `Probe Table`
	__State__00000002_.IsDecisionNode = false
	__State__00000002_.IsFictif = false
	__State__00000002_.IsEndState = false

	__State__00000003_.Name = `Form - Update Stage Of Interest`
	__State__00000003_.IsDecisionNode = false
	__State__00000003_.IsFictif = false
	__State__00000003_.IsEndState = false

	__State__00000004_.Name = `Enforce Model Semantic`
	__State__00000004_.IsDecisionNode = false
	__State__00000004_.IsFictif = false
	__State__00000004_.IsEndState = false

	__State__00000005_.Name = `Update Probe Form`
	__State__00000005_.IsDecisionNode = false
	__State__00000005_.IsFictif = false
	__State__00000005_.IsEndState = false

	__State__00000006_.Name = `Update Probe Table`
	__State__00000006_.IsDecisionNode = false
	__State__00000006_.IsFictif = false
	__State__00000006_.IsEndState = false

	__State__00000007_.Name = `Update Probe Tree`
	__State__00000007_.IsDecisionNode = false
	__State__00000007_.IsFictif = false
	__State__00000007_.IsEndState = false

	__State__00000008_.Name = `Load Stage`
	__State__00000008_.IsDecisionNode = false
	__State__00000008_.IsFictif = false
	__State__00000008_.IsEndState = false

	__State__00000009_.Name = `Persist Stage`
	__State__00000009_.IsDecisionNode = false
	__State__00000009_.IsFictif = false
	__State__00000009_.IsEndState = false

	__State__00000010_.Name = `Update UX ?`
	__State__00000010_.IsDecisionNode = true
	__State__00000010_.IsFictif = true
	__State__00000010_.IsEndState = false

	__State__00000011_.Name = `Update UX ?`
	__State__00000011_.IsDecisionNode = true
	__State__00000011_.IsFictif = false
	__State__00000011_.IsEndState = false

	__State__00000012_.Name = `Clean Stage`
	__State__00000012_.IsDecisionNode = false
	__State__00000012_.IsFictif = false
	__State__00000012_.IsEndState = false

	__StateMachine__00000000_.Name = `UX Loop`
	__StateMachine__00000000_.IsNodeExpanded = true

	__StateShape__00000000_.Name = `Probe Form`
	__StateShape__00000000_.IsExpanded = false
	__StateShape__00000000_.X = 166.000000
	__StateShape__00000000_.Y = 165.000000
	__StateShape__00000000_.Width = 221.000000
	__StateShape__00000000_.Height = 80.000000

	__StateShape__00000001_.Name = `Probe Tree`
	__StateShape__00000001_.IsExpanded = false
	__StateShape__00000001_.X = 559.000000
	__StateShape__00000001_.Y = 242.000000
	__StateShape__00000001_.Width = 200.000000
	__StateShape__00000001_.Height = 80.000000

	__StateShape__00000002_.Name = `Probe Table`
	__StateShape__00000002_.IsExpanded = false
	__StateShape__00000002_.X = 782.000000
	__StateShape__00000002_.Y = 246.000000
	__StateShape__00000002_.Width = 200.000000
	__StateShape__00000002_.Height = 80.000000

	__StateShape__00000003_.Name = `Form - Update Stage Of Interest`
	__StateShape__00000003_.IsExpanded = false
	__StateShape__00000003_.X = 333.000000
	__StateShape__00000003_.Y = 402.999985
	__StateShape__00000003_.Width = 200.000000
	__StateShape__00000003_.Height = 125.000031

	__StateShape__00000004_.Name = `Enforce Model Semantic`
	__StateShape__00000004_.IsExpanded = false
	__StateShape__00000004_.X = 332.000000
	__StateShape__00000004_.Y = 629.000015
	__StateShape__00000004_.Width = 646.000000
	__StateShape__00000004_.Height = 121.000000

	__StateShape__00000005_.Name = `Update Probe Form`
	__StateShape__00000005_.IsExpanded = false
	__StateShape__00000005_.X = 329.000000
	__StateShape__00000005_.Y = 1005.999954
	__StateShape__00000005_.Width = 200.000000
	__StateShape__00000005_.Height = 80.000000

	__StateShape__00000006_.Name = `Update Probe Table`
	__StateShape__00000006_.IsExpanded = false
	__StateShape__00000006_.X = 552.000000
	__StateShape__00000006_.Y = 1003.999954
	__StateShape__00000006_.Width = 200.000000
	__StateShape__00000006_.Height = 80.000000

	__StateShape__00000007_.Name = `Update Probe Tree`
	__StateShape__00000007_.IsExpanded = false
	__StateShape__00000007_.X = 772.000000
	__StateShape__00000007_.Y = 1003.999954
	__StateShape__00000007_.Width = 200.000000
	__StateShape__00000007_.Height = 80.000000

	__StateShape__00000008_.Name = `Load Stage`
	__StateShape__00000008_.IsExpanded = false
	__StateShape__00000008_.X = 14.000000
	__StateShape__00000008_.Y = 549.999984
	__StateShape__00000008_.Width = 107.000000
	__StateShape__00000008_.Height = 20.000000

	__StateShape__00000009_.Name = `Persist Stage`
	__StateShape__00000009_.IsExpanded = false
	__StateShape__00000009_.X = 52.000000
	__StateShape__00000009_.Y = 888.999954
	__StateShape__00000009_.Width = 200.000000
	__StateShape__00000009_.Height = 80.000000

	__StateShape__00000010_.Name = `Update UX ?`
	__StateShape__00000010_.IsExpanded = false
	__StateShape__00000010_.X = 539.000000
	__StateShape__00000010_.Y = 826.999954
	__StateShape__00000010_.Width = 153.000000
	__StateShape__00000010_.Height = 50.000000

	__StateShape__00000011_.Name = `Update UX ?`
	__StateShape__00000011_.IsExpanded = false
	__StateShape__00000011_.X = 105.000000
	__StateShape__00000011_.Y = 1156.000015
	__StateShape__00000011_.Width = 120.000000
	__StateShape__00000011_.Height = 50.000000

	__StateShape__00000012_.Name = `Clean Stage`
	__StateShape__00000012_.IsExpanded = false
	__StateShape__00000012_.X = 740.000000
	__StateShape__00000012_.Y = 409.000000
	__StateShape__00000012_.Width = 200.000000
	__StateShape__00000012_.Height = 87.000000

	__Transition__00000000_.Name = `Submit`

	__Transition__00000001_.Name = ``

	__Transition__00000002_.Name = `/Model Updated`

	__Transition__00000003_.Name = ``

	__Transition__00000004_.Name = ``

	__Transition__00000005_.Name = `/Stage loaded`

	__Transition__00000006_.Name = `/Stage persisted`

	__Transition__00000007_.Name = `Yes`

	__Transition__00000008_.Name = `Yes`

	__Transition__00000009_.Name = `Form - Update Stage Of Interest to Clean Stage`

	__Transition__00000010_.Name = ``

	__Transition__00000011_.Name = `Stage Modified ? to Probe Form`

	__Transition_Shape__00000000_.Name = `UX 1 - Waiting for User Input to UX 1 - Update Stage`
	__Transition_Shape__00000000_.StartRatio = 0.815760
	__Transition_Shape__00000000_.EndRatio = 0.866415
	__Transition_Shape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000000_.CornerOffsetRatio = 2.175087

	__Transition_Shape__00000001_.Name = `UX 1 - Update Stage to Enforce Model Semantic`
	__Transition_Shape__00000001_.StartRatio = 0.540885
	__Transition_Shape__00000001_.EndRatio = 0.166852
	__Transition_Shape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000001_.CornerOffsetRatio = 1.413333

	__Transition_Shape__00000002_.Name = `Enforce Model Semantic to UX 1 - Update UX`
	__Transition_Shape__00000002_.StartRatio = 0.512053
	__Transition_Shape__00000002_.EndRatio = 0.831915
	__Transition_Shape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000002_.CornerOffsetRatio = 1.407713

	__Transition_Shape__00000003_.Name = `Enforce Model Semantic to UX 2 - Update UX`
	__Transition_Shape__00000003_.StartRatio = 0.821415
	__Transition_Shape__00000003_.EndRatio = 0.166415
	__Transition_Shape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000003_.CornerOffsetRatio = 1.625087

	__Transition_Shape__00000004_.Name = `Enforce Model Semantic to UX n - Update UX`
	__Transition_Shape__00000004_.StartRatio = 0.830939
	__Transition_Shape__00000004_.EndRatio = 0.801415
	__Transition_Shape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000004_.CornerOffsetRatio = 1.662587

	__Transition_Shape__00000005_.Name = `Load Stage to Enforce Model Semantic`
	__Transition_Shape__00000005_.StartRatio = 0.550349
	__Transition_Shape__00000005_.EndRatio = 0.298697
	__Transition_Shape__00000005_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000005_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000005_.CornerOffsetRatio = 1.320401

	__Transition_Shape__00000006_.Name = `Persist Stage to UX 1 - Update UX`
	__Transition_Shape__00000006_.StartRatio = 0.741415
	__Transition_Shape__00000006_.EndRatio = 0.760691
	__Transition_Shape__00000006_.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000006_.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000006_.CornerOffsetRatio = 1.737587

	__Transition_Shape__00000007_.Name = `Stage Modified ? to Persist Stage`
	__Transition_Shape__00000007_.StartRatio = 0.831915
	__Transition_Shape__00000007_.EndRatio = 0.383333
	__Transition_Shape__00000007_.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000007_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000007_.CornerOffsetRatio = 1.100139

	__Transition_Shape__00000008_.Name = `Update UX ? to UX 1 - Update UX`
	__Transition_Shape__00000008_.StartRatio = 0.475693
	__Transition_Shape__00000008_.EndRatio = 0.500000
	__Transition_Shape__00000008_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000008_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000008_.CornerOffsetRatio = 1.560691

	__Transition_Shape__00000009_.Name = `Form - Update Stage Of Interest to Clean Stage`
	__Transition_Shape__00000009_.StartRatio = 0.500000
	__Transition_Shape__00000009_.EndRatio = 0.500000
	__Transition_Shape__00000009_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000009_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000009_.CornerOffsetRatio = 1.200000

	__Transition_Shape__00000010_.Name = `Clean Stage to Enforce Model Semantic`
	__Transition_Shape__00000010_.StartRatio = 0.786415
	__Transition_Shape__00000010_.EndRatio = 0.887435
	__Transition_Shape__00000010_.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000010_.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000010_.CornerOffsetRatio = 1.950087

	__Transition_Shape__00000011_.Name = `Stage Modified ? to Probe Form`
	__Transition_Shape__00000011_.StartRatio = 0.500000
	__Transition_Shape__00000011_.EndRatio = 0.500000
	__Transition_Shape__00000011_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000011_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000011_.CornerOffsetRatio = 3.046909

	// insertion point for setup of pointers
	__Architecture__00000000_.StateMachines = append(__Architecture__00000000_.StateMachines, __StateMachine__00000000_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000011_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000000_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000001_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000002_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000003_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000004_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000005_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000006_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000007_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000008_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000009_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000010_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000012_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000000_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000001_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000002_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000003_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000004_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000005_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000006_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000007_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000008_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000009_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000010_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000011_)
	__State__00000000_.Parent = nil
	__State__00000000_.Diagrams = append(__State__00000000_.Diagrams, __Diagram__00000000_)
	__State__00000000_.Entry = nil
	__State__00000000_.Exit = nil
	__State__00000001_.Parent = nil
	__State__00000001_.Diagrams = append(__State__00000001_.Diagrams, __Diagram__00000000_)
	__State__00000001_.Entry = nil
	__State__00000001_.Exit = nil
	__State__00000002_.Parent = nil
	__State__00000002_.Diagrams = append(__State__00000002_.Diagrams, __Diagram__00000000_)
	__State__00000002_.Entry = nil
	__State__00000002_.Exit = nil
	__State__00000003_.Parent = nil
	__State__00000003_.Diagrams = append(__State__00000003_.Diagrams, __Diagram__00000000_)
	__State__00000003_.Entry = __Action__00000000_
	__State__00000003_.Exit = nil
	__State__00000004_.Parent = nil
	__State__00000004_.Diagrams = append(__State__00000004_.Diagrams, __Diagram__00000000_)
	__State__00000004_.Entry = nil
	__State__00000004_.Exit = __Action__00000001_
	__State__00000005_.Parent = nil
	__State__00000005_.Diagrams = append(__State__00000005_.Diagrams, __Diagram__00000000_)
	__State__00000005_.Entry = nil
	__State__00000005_.Exit = nil
	__State__00000006_.Parent = nil
	__State__00000006_.Diagrams = append(__State__00000006_.Diagrams, __Diagram__00000000_)
	__State__00000006_.Entry = nil
	__State__00000006_.Exit = nil
	__State__00000007_.Parent = nil
	__State__00000007_.Diagrams = append(__State__00000007_.Diagrams, __Diagram__00000000_)
	__State__00000007_.Entry = nil
	__State__00000007_.Exit = nil
	__State__00000008_.Parent = nil
	__State__00000008_.Diagrams = append(__State__00000008_.Diagrams, __Diagram__00000000_)
	__State__00000008_.Entry = nil
	__State__00000008_.Exit = nil
	__State__00000009_.Parent = nil
	__State__00000009_.Diagrams = append(__State__00000009_.Diagrams, __Diagram__00000000_)
	__State__00000009_.Entry = nil
	__State__00000009_.Exit = nil
	__State__00000010_.Parent = nil
	__State__00000010_.Diagrams = append(__State__00000010_.Diagrams, __Diagram__00000000_)
	__State__00000010_.Entry = nil
	__State__00000010_.Exit = nil
	__State__00000011_.Parent = nil
	__State__00000011_.Diagrams = append(__State__00000011_.Diagrams, __Diagram__00000000_)
	__State__00000011_.Entry = nil
	__State__00000011_.Exit = nil
	__State__00000012_.Parent = nil
	__State__00000012_.Diagrams = append(__State__00000012_.Diagrams, __Diagram__00000000_)
	__State__00000012_.Entry = nil
	__State__00000012_.Activities = append(__State__00000012_.Activities, __Activities__00000000_)
	__State__00000012_.Exit = nil
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000000_)
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000001_)
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000002_)
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000003_)
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000004_)
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000005_)
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000006_)
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000007_)
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000008_)
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000009_)
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000010_)
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000011_)
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000012_)
	__StateMachine__00000000_.Diagrams = append(__StateMachine__00000000_.Diagrams, __Diagram__00000000_)
	__StateMachine__00000000_.InitialState = __State__00000008_
	__StateShape__00000000_.State = __State__00000000_
	__StateShape__00000001_.State = __State__00000001_
	__StateShape__00000002_.State = __State__00000002_
	__StateShape__00000003_.State = __State__00000003_
	__StateShape__00000004_.State = __State__00000004_
	__StateShape__00000005_.State = __State__00000005_
	__StateShape__00000006_.State = __State__00000006_
	__StateShape__00000007_.State = __State__00000007_
	__StateShape__00000008_.State = __State__00000008_
	__StateShape__00000009_.State = __State__00000009_
	__StateShape__00000010_.State = __State__00000010_
	__StateShape__00000011_.State = __State__00000011_
	__StateShape__00000012_.State = __State__00000012_
	__Transition__00000000_.Start = __State__00000000_
	__Transition__00000000_.End = __State__00000003_
	__Transition__00000000_.Guard = nil
	__Transition__00000000_.Diagrams = append(__Transition__00000000_.Diagrams, __Diagram__00000000_)
	__Transition__00000001_.Start = __State__00000003_
	__Transition__00000001_.End = __State__00000004_
	__Transition__00000001_.Guard = __Guard__00000003_
	__Transition__00000001_.Diagrams = append(__Transition__00000001_.Diagrams, __Diagram__00000000_)
	__Transition__00000002_.Start = __State__00000004_
	__Transition__00000002_.End = __State__00000010_
	__Transition__00000002_.Guard = nil
	__Transition__00000002_.Diagrams = append(__Transition__00000002_.Diagrams, __Diagram__00000000_)
	__Transition__00000003_.Start = __State__00000005_
	__Transition__00000003_.End = __State__00000006_
	__Transition__00000003_.Guard = nil
	__Transition__00000003_.Diagrams = append(__Transition__00000003_.Diagrams, __Diagram__00000000_)
	__Transition__00000004_.Start = __State__00000006_
	__Transition__00000004_.End = __State__00000007_
	__Transition__00000004_.Guard = nil
	__Transition__00000004_.Diagrams = append(__Transition__00000004_.Diagrams, __Diagram__00000000_)
	__Transition__00000005_.Start = __State__00000008_
	__Transition__00000005_.End = __State__00000004_
	__Transition__00000005_.Guard = nil
	__Transition__00000005_.Diagrams = append(__Transition__00000005_.Diagrams, __Diagram__00000000_)
	__Transition__00000006_.Start = __State__00000009_
	__Transition__00000006_.End = __State__00000011_
	__Transition__00000006_.Guard = nil
	__Transition__00000006_.Diagrams = append(__Transition__00000006_.Diagrams, __Diagram__00000000_)
	__Transition__00000007_.Start = __State__00000010_
	__Transition__00000007_.End = __State__00000009_
	__Transition__00000007_.Guard = __Guard__00000000_
	__Transition__00000007_.Diagrams = append(__Transition__00000007_.Diagrams, __Diagram__00000000_)
	__Transition__00000008_.Start = __State__00000011_
	__Transition__00000008_.End = __State__00000005_
	__Transition__00000008_.Guard = __Guard__00000000_
	__Transition__00000008_.Diagrams = append(__Transition__00000008_.Diagrams, __Diagram__00000000_)
	__Transition__00000009_.Start = __State__00000003_
	__Transition__00000009_.End = __State__00000012_
	__Transition__00000009_.Guard = __Guard__00000002_
	__Transition__00000009_.Diagrams = append(__Transition__00000009_.Diagrams, __Diagram__00000000_)
	__Transition__00000010_.Start = __State__00000012_
	__Transition__00000010_.End = __State__00000004_
	__Transition__00000010_.Guard = nil
	__Transition__00000010_.Diagrams = append(__Transition__00000010_.Diagrams, __Diagram__00000000_)
	__Transition__00000011_.Start = __State__00000010_
	__Transition__00000011_.End = __State__00000000_
	__Transition__00000011_.Guard = __Guard__00000001_
	__Transition__00000011_.Diagrams = append(__Transition__00000011_.Diagrams, __Diagram__00000000_)
	__Transition_Shape__00000000_.Transition = __Transition__00000000_
	__Transition_Shape__00000001_.Transition = __Transition__00000001_
	__Transition_Shape__00000002_.Transition = __Transition__00000002_
	__Transition_Shape__00000003_.Transition = __Transition__00000003_
	__Transition_Shape__00000004_.Transition = __Transition__00000004_
	__Transition_Shape__00000005_.Transition = __Transition__00000005_
	__Transition_Shape__00000006_.Transition = __Transition__00000006_
	__Transition_Shape__00000007_.Transition = __Transition__00000007_
	__Transition_Shape__00000008_.Transition = __Transition__00000008_
	__Transition_Shape__00000009_.Transition = __Transition__00000009_
	__Transition_Shape__00000010_.Transition = __Transition__00000010_
	__Transition_Shape__00000011_.Transition = __Transition__00000011_
}
