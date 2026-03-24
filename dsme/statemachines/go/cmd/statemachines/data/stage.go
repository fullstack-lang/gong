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

	__Architecture__00000000_ := (&models.Architecture{Name: `Traffic Lights of the world`}).Stage(stage)

	__Diagram__00000000_ := (&models.Diagram{Name: `Traffic Light FR Diagram`}).Stage(stage)
	__Diagram__00000001_ := (&models.Diagram{Name: `Traffic Light UK Diagram`}).Stage(stage)

	__Message__00000003_ := (&models.Message{Name: `03:37:16 Off UK to Red UK -> Repair Report`}).Stage(stage)
	__Message__00000004_ := (&models.Message{Name: `03:48:43 Off UK to Red UK -> Repair Report`}).Stage(stage)

	__MessageType__00000000_ := (&models.MessageType{Name: `Repair Report`}).Stage(stage)

	__Object__00000001_ := (&models.Object{Name: `01/MI DOF/ 2025-12-08 DEP/ 102802`}).Stage(stage)
	__Object__00000003_ := (&models.Object{Name: `02/MI DOF/ 2026-03-25 DEP/ 033701`}).Stage(stage)

	__Role__00000000_ := (&models.Role{Name: `Technician`}).Stage(stage)
	__Role__00000001_ := (&models.Role{Name: `Timer`}).Stage(stage)

	__State__00000000_ := (&models.State{Name: `Initial State FR`}).Stage(stage)
	__State__00000001_ := (&models.State{Name: `Red FR`}).Stage(stage)
	__State__00000002_ := (&models.State{Name: `Green FR`}).Stage(stage)
	__State__00000003_ := (&models.State{Name: `Yellow FR`}).Stage(stage)
	__State__00000004_ := (&models.State{Name: `On FR`}).Stage(stage)
	__State__00000005_ := (&models.State{Name: `Off FR`}).Stage(stage)
	__State__00000006_ := (&models.State{Name: `Initial State UK`}).Stage(stage)
	__State__00000007_ := (&models.State{Name: `Red UK`}).Stage(stage)
	__State__00000008_ := (&models.State{Name: `Yellow UK`}).Stage(stage)
	__State__00000009_ := (&models.State{Name: `Green UK`}).Stage(stage)
	__State__00000010_ := (&models.State{Name: `On UK`}).Stage(stage)
	__State__00000011_ := (&models.State{Name: `Off UK`}).Stage(stage)

	__StateMachine__00000000_ := (&models.StateMachine{Name: `Traffic Lights FR`}).Stage(stage)
	__StateMachine__00000001_ := (&models.StateMachine{Name: `Traffic Light UK`}).Stage(stage)

	__StateShape__00000000_ := (&models.StateShape{Name: `Red FR`}).Stage(stage)
	__StateShape__00000002_ := (&models.StateShape{Name: `Green FR`}).Stage(stage)
	__StateShape__00000003_ := (&models.StateShape{Name: `Yellow FR`}).Stage(stage)
	__StateShape__00000004_ := (&models.StateShape{Name: `On FR`}).Stage(stage)
	__StateShape__00000005_ := (&models.StateShape{Name: `Off FR`}).Stage(stage)
	__StateShape__00000006_ := (&models.StateShape{Name: `Red UK`}).Stage(stage)
	__StateShape__00000007_ := (&models.StateShape{Name: `Initial State UK`}).Stage(stage)
	__StateShape__00000008_ := (&models.StateShape{Name: `Yellow UK`}).Stage(stage)
	__StateShape__00000009_ := (&models.StateShape{Name: `Green UK`}).Stage(stage)
	__StateShape__00000010_ := (&models.StateShape{Name: `On UK`}).Stage(stage)
	__StateShape__00000011_ := (&models.StateShape{Name: `Off UK`}).Stage(stage)
	__StateShape__00000012_ := (&models.StateShape{Name: `Initial State FR`}).Stage(stage)

	__Transition__00000000_ := (&models.Transition{Name: `Initial State FR to Red FR`}).Stage(stage)
	__Transition__00000001_ := (&models.Transition{Name: `Red FR to Green FR`}).Stage(stage)
	__Transition__00000002_ := (&models.Transition{Name: `Green FR to Yellow FR`}).Stage(stage)
	__Transition__00000003_ := (&models.Transition{Name: `Yellow FR to Red FR`}).Stage(stage)
	__Transition__00000004_ := (&models.Transition{Name: `On to Off`}).Stage(stage)
	__Transition__00000005_ := (&models.Transition{Name: `Off to On`}).Stage(stage)
	__Transition__00000006_ := (&models.Transition{Name: `Initial State UK to New State`}).Stage(stage)
	__Transition__00000007_ := (&models.Transition{Name: `Red UK to Yellow UK`}).Stage(stage)
	__Transition__00000008_ := (&models.Transition{Name: `Yellow UK to Green UK`}).Stage(stage)
	__Transition__00000009_ := (&models.Transition{Name: `On UK to Off UK`}).Stage(stage)
	__Transition__00000010_ := (&models.Transition{Name: `Off UK to Red UK`}).Stage(stage)
	__Transition__00000011_ := (&models.Transition{Name: `Green UK to Red UK`}).Stage(stage)

	__Transition_Shape__00000000_ := (&models.Transition_Shape{Name: `Initial State FR to Red FR`}).Stage(stage)
	__Transition_Shape__00000001_ := (&models.Transition_Shape{Name: `Red FR to Green FR`}).Stage(stage)
	__Transition_Shape__00000002_ := (&models.Transition_Shape{Name: `Green FR to Yellow FR`}).Stage(stage)
	__Transition_Shape__00000003_ := (&models.Transition_Shape{Name: `Yellow FR to Red FR`}).Stage(stage)
	__Transition_Shape__00000004_ := (&models.Transition_Shape{Name: `New State to New State`}).Stage(stage)
	__Transition_Shape__00000005_ := (&models.Transition_Shape{Name: `Off to On`}).Stage(stage)
	__Transition_Shape__00000006_ := (&models.Transition_Shape{Name: `Initial State UK to New State`}).Stage(stage)
	__Transition_Shape__00000007_ := (&models.Transition_Shape{Name: `Red UK to Yellow UK`}).Stage(stage)
	__Transition_Shape__00000008_ := (&models.Transition_Shape{Name: `Yellow UK to Green UK`}).Stage(stage)
	__Transition_Shape__00000009_ := (&models.Transition_Shape{Name: `On UK to Off UK`}).Stage(stage)
	__Transition_Shape__00000010_ := (&models.Transition_Shape{Name: `Off UK to On UK`}).Stage(stage)
	__Transition_Shape__00000011_ := (&models.Transition_Shape{Name: `Green UK to Red UK`}).Stage(stage)

	// insertion point for initialization of values

	__Architecture__00000000_.Name = `Traffic Lights of the world`
	__Architecture__00000000_.NbPixPerCharacter = 8.000000

	__Diagram__00000000_.Name = `Traffic Light FR Diagram`
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsInRenameMode = false

	__Diagram__00000001_.Name = `Traffic Light UK Diagram`
	__Diagram__00000001_.IsChecked = false
	__Diagram__00000001_.IsExpanded = true
	__Diagram__00000001_.IsEditable_ = true
	__Diagram__00000001_.IsInRenameMode = false

	__Message__00000003_.Name = `03:37:16 Off UK to Red UK -> Repair Report`
	__Message__00000003_.IsSelected = false

	__Message__00000004_.Name = `03:48:43 Off UK to Red UK -> Repair Report`
	__Message__00000004_.IsSelected = true

	__MessageType__00000000_.Name = `Repair Report`
	__MessageType__00000000_.Description = ``

	__Object__00000001_.Name = `01/MI DOF/ 2025-12-08 DEP/ 102802`
	__Object__00000001_.IsSelected = false
	__Object__00000001_.Rank = 0
	__Object__00000001_.DOF, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")

	__Object__00000003_.Name = `02/MI DOF/ 2026-03-25 DEP/ 033701`
	__Object__00000003_.IsSelected = true
	__Object__00000003_.Rank = 0
	__Object__00000003_.DOF, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")

	__Role__00000000_.Name = `Technician`
	__Role__00000000_.Acronym = `Tech`

	__Role__00000001_.Name = `Timer`
	__Role__00000001_.Acronym = `Timer`

	__State__00000000_.Name = `Initial State FR`
	__State__00000000_.IsDecisionNode = false
	__State__00000000_.IsFictif = false
	__State__00000000_.IsEndState = false
	__State__00000000_.IsInRenameMode = false

	__State__00000001_.Name = `Red FR`
	__State__00000001_.IsDecisionNode = false
	__State__00000001_.IsFictif = false
	__State__00000001_.IsEndState = false
	__State__00000001_.IsInRenameMode = false

	__State__00000002_.Name = `Green FR`
	__State__00000002_.IsDecisionNode = false
	__State__00000002_.IsFictif = false
	__State__00000002_.IsEndState = false
	__State__00000002_.IsInRenameMode = false

	__State__00000003_.Name = `Yellow FR`
	__State__00000003_.IsDecisionNode = false
	__State__00000003_.IsFictif = false
	__State__00000003_.IsEndState = false
	__State__00000003_.IsInRenameMode = false

	__State__00000004_.Name = `On FR`
	__State__00000004_.IsDecisionNode = false
	__State__00000004_.IsFictif = false
	__State__00000004_.IsEndState = false
	__State__00000004_.IsInRenameMode = false

	__State__00000005_.Name = `Off FR`
	__State__00000005_.IsDecisionNode = false
	__State__00000005_.IsFictif = false
	__State__00000005_.IsEndState = true
	__State__00000005_.IsInRenameMode = false

	__State__00000006_.Name = `Initial State UK`
	__State__00000006_.IsDecisionNode = false
	__State__00000006_.IsFictif = false
	__State__00000006_.IsEndState = false
	__State__00000006_.IsInRenameMode = false

	__State__00000007_.Name = `Red UK`
	__State__00000007_.IsDecisionNode = false
	__State__00000007_.IsFictif = false
	__State__00000007_.IsEndState = false
	__State__00000007_.IsInRenameMode = false

	__State__00000008_.Name = `Yellow UK`
	__State__00000008_.IsDecisionNode = false
	__State__00000008_.IsFictif = false
	__State__00000008_.IsEndState = false
	__State__00000008_.IsInRenameMode = false

	__State__00000009_.Name = `Green UK`
	__State__00000009_.IsDecisionNode = false
	__State__00000009_.IsFictif = false
	__State__00000009_.IsEndState = false
	__State__00000009_.IsInRenameMode = false

	__State__00000010_.Name = `On UK`
	__State__00000010_.IsDecisionNode = false
	__State__00000010_.IsFictif = false
	__State__00000010_.IsEndState = false
	__State__00000010_.IsInRenameMode = false

	__State__00000011_.Name = `Off UK`
	__State__00000011_.IsDecisionNode = false
	__State__00000011_.IsFictif = false
	__State__00000011_.IsEndState = true
	__State__00000011_.IsInRenameMode = false

	__StateMachine__00000000_.Name = `Traffic Lights FR`
	__StateMachine__00000000_.IsNodeExpanded = true

	__StateMachine__00000001_.Name = `Traffic Light UK`
	__StateMachine__00000001_.IsNodeExpanded = true

	__StateShape__00000000_.Name = `Red FR`
	__StateShape__00000000_.X = 578.000000
	__StateShape__00000000_.Y = 204.000000
	__StateShape__00000000_.Width = 200.000000
	__StateShape__00000000_.Height = 80.000000
	__StateShape__00000000_.IsHidden = false

	__StateShape__00000002_.Name = `Green FR`
	__StateShape__00000002_.X = 556.000000
	__StateShape__00000002_.Y = 356.000000
	__StateShape__00000002_.Width = 200.000000
	__StateShape__00000002_.Height = 80.000000
	__StateShape__00000002_.IsHidden = false

	__StateShape__00000003_.Name = `Yellow FR`
	__StateShape__00000003_.X = 566.000000
	__StateShape__00000003_.Y = 524.000000
	__StateShape__00000003_.Width = 200.000000
	__StateShape__00000003_.Height = 80.000000
	__StateShape__00000003_.IsHidden = false

	__StateShape__00000004_.Name = `On FR`
	__StateShape__00000004_.X = 351.000000
	__StateShape__00000004_.Y = 141.714286
	__StateShape__00000004_.Width = 678.000000
	__StateShape__00000004_.Height = 492.285714
	__StateShape__00000004_.IsHidden = false

	__StateShape__00000005_.Name = `Off FR`
	__StateShape__00000005_.X = 1120.000000
	__StateShape__00000005_.Y = 379.875000
	__StateShape__00000005_.Width = 91.000000
	__StateShape__00000005_.Height = 36.000000
	__StateShape__00000005_.IsHidden = false

	__StateShape__00000006_.Name = `Red UK`
	__StateShape__00000006_.X = 344.000000
	__StateShape__00000006_.Y = 253.000000
	__StateShape__00000006_.Width = 200.000000
	__StateShape__00000006_.Height = 80.000000
	__StateShape__00000006_.IsHidden = false

	__StateShape__00000007_.Name = `Initial State UK`
	__StateShape__00000007_.X = 324.000000
	__StateShape__00000007_.Y = 91.000000
	__StateShape__00000007_.Width = 121.000000
	__StateShape__00000007_.Height = 20.000000
	__StateShape__00000007_.IsHidden = false

	__StateShape__00000008_.Name = `Yellow UK`
	__StateShape__00000008_.X = 340.000000
	__StateShape__00000008_.Y = 405.000000
	__StateShape__00000008_.Width = 200.000000
	__StateShape__00000008_.Height = 80.000000
	__StateShape__00000008_.IsHidden = false

	__StateShape__00000009_.Name = `Green UK`
	__StateShape__00000009_.X = 344.000000
	__StateShape__00000009_.Y = 576.000000
	__StateShape__00000009_.Width = 200.000000
	__StateShape__00000009_.Height = 80.000000
	__StateShape__00000009_.IsHidden = false

	__StateShape__00000010_.Name = `On UK`
	__StateShape__00000010_.X = 197.000000
	__StateShape__00000010_.Y = 209.000000
	__StateShape__00000010_.Width = 632.000000
	__StateShape__00000010_.Height = 479.000000
	__StateShape__00000010_.IsHidden = false

	__StateShape__00000011_.Name = `Off UK`
	__StateShape__00000011_.X = 998.000000
	__StateShape__00000011_.Y = 388.000000
	__StateShape__00000011_.Width = 94.000000
	__StateShape__00000011_.Height = 36.000000
	__StateShape__00000011_.IsHidden = false

	__StateShape__00000012_.Name = `Initial State FR`
	__StateShape__00000012_.X = 100.000000
	__StateShape__00000012_.Y = 100.000000
	__StateShape__00000012_.Width = 200.000000
	__StateShape__00000012_.Height = 80.000000
	__StateShape__00000012_.IsHidden = false

	__Transition__00000000_.Name = `Initial State FR to Red FR`
	__Transition__00000000_.IsInRenameMode = false

	__Transition__00000001_.Name = `Red FR to Green FR`
	__Transition__00000001_.IsInRenameMode = false

	__Transition__00000002_.Name = `Green FR to Yellow FR`
	__Transition__00000002_.IsInRenameMode = false

	__Transition__00000003_.Name = `Yellow FR to Red FR`
	__Transition__00000003_.IsInRenameMode = false

	__Transition__00000004_.Name = `On to Off`
	__Transition__00000004_.IsInRenameMode = false

	__Transition__00000005_.Name = `Off to On`
	__Transition__00000005_.IsInRenameMode = false

	__Transition__00000006_.Name = `Initial State UK to New State`
	__Transition__00000006_.IsInRenameMode = false

	__Transition__00000007_.Name = `Red UK to Yellow UK`
	__Transition__00000007_.IsInRenameMode = false

	__Transition__00000008_.Name = `Yellow UK to Green UK`
	__Transition__00000008_.IsInRenameMode = false

	__Transition__00000009_.Name = `On UK to Off UK`
	__Transition__00000009_.IsInRenameMode = false

	__Transition__00000010_.Name = `Off UK to Red UK`
	__Transition__00000010_.IsInRenameMode = false

	__Transition__00000011_.Name = `Green UK to Red UK`
	__Transition__00000011_.IsInRenameMode = false

	__Transition_Shape__00000000_.Name = `Initial State FR to Red FR`
	__Transition_Shape__00000000_.StartRatio = 1.000000
	__Transition_Shape__00000000_.EndRatio = 0.100073
	__Transition_Shape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000000_.CornerOffsetRatio = 1.100293
	__Transition_Shape__00000000_.IsHidden = false

	__Transition_Shape__00000001_.Name = `Red FR to Green FR`
	__Transition_Shape__00000001_.StartRatio = 0.598789
	__Transition_Shape__00000001_.EndRatio = 0.608789
	__Transition_Shape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000001_.CornerOffsetRatio = 1.425073
	__Transition_Shape__00000001_.IsHidden = false

	__Transition_Shape__00000002_.Name = `Green FR to Yellow FR`
	__Transition_Shape__00000002_.StartRatio = 0.617656
	__Transition_Shape__00000002_.EndRatio = 0.623789
	__Transition_Shape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000002_.CornerOffsetRatio = 1.687573
	__Transition_Shape__00000002_.IsHidden = false

	__Transition_Shape__00000003_.Name = `Yellow FR to Red FR`
	__Transition_Shape__00000003_.StartRatio = 0.500000
	__Transition_Shape__00000003_.EndRatio = 0.700073
	__Transition_Shape__00000003_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000003_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000003_.CornerOffsetRatio = 2.103789
	__Transition_Shape__00000003_.IsHidden = false

	__Transition_Shape__00000004_.Name = `New State to New State`
	__Transition_Shape__00000004_.StartRatio = 0.884949
	__Transition_Shape__00000004_.EndRatio = 0.503635
	__Transition_Shape__00000004_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000004_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000004_.CornerOffsetRatio = 1.436221
	__Transition_Shape__00000004_.IsHidden = false

	__Transition_Shape__00000005_.Name = `Off to On`
	__Transition_Shape__00000005_.StartRatio = 0.799536
	__Transition_Shape__00000005_.EndRatio = 0.225073
	__Transition_Shape__00000005_.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000005_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000005_.CornerOffsetRatio = -8.888889
	__Transition_Shape__00000005_.IsHidden = false

	__Transition_Shape__00000006_.Name = `Initial State UK to New State`
	__Transition_Shape__00000006_.StartRatio = 1.000000
	__Transition_Shape__00000006_.EndRatio = 0.523789
	__Transition_Shape__00000006_.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000006_.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000006_.CornerOffsetRatio = 4.350293
	__Transition_Shape__00000006_.IsHidden = false

	__Transition_Shape__00000007_.Name = `Red UK to Yellow UK`
	__Transition_Shape__00000007_.StartRatio = 0.458789
	__Transition_Shape__00000007_.EndRatio = 0.478789
	__Transition_Shape__00000007_.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000007_.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000007_.CornerOffsetRatio = 1.625073
	__Transition_Shape__00000007_.IsHidden = false

	__Transition_Shape__00000008_.Name = `Yellow UK to Green UK`
	__Transition_Shape__00000008_.StartRatio = 0.478789
	__Transition_Shape__00000008_.EndRatio = 0.363789
	__Transition_Shape__00000008_.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000008_.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000008_.CornerOffsetRatio = 1.537573
	__Transition_Shape__00000008_.IsHidden = false

	__Transition_Shape__00000009_.Name = `On UK to Off UK`
	__Transition_Shape__00000009_.StartRatio = 0.722350
	__Transition_Shape__00000009_.EndRatio = 0.816572
	__Transition_Shape__00000009_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000009_.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000009_.CornerOffsetRatio = 0.488530
	__Transition_Shape__00000009_.IsHidden = false

	__Transition_Shape__00000010_.Name = `Off UK to On UK`
	__Transition_Shape__00000010_.StartRatio = 0.805934
	__Transition_Shape__00000010_.EndRatio = 0.187573
	__Transition_Shape__00000010_.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000010_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000010_.CornerOffsetRatio = 1.000163
	__Transition_Shape__00000010_.IsHidden = false

	__Transition_Shape__00000011_.Name = `Green UK to Red UK`
	__Transition_Shape__00000011_.StartRatio = 0.500000
	__Transition_Shape__00000011_.EndRatio = 0.675073
	__Transition_Shape__00000011_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000011_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000011_.CornerOffsetRatio = 2.206415
	__Transition_Shape__00000011_.IsHidden = false

	// insertion point for setup of pointers
	__Architecture__00000000_.StateMachines = append(__Architecture__00000000_.StateMachines, __StateMachine__00000000_)
	__Architecture__00000000_.StateMachines = append(__Architecture__00000000_.StateMachines, __StateMachine__00000001_)
	__Architecture__00000000_.Roles = append(__Architecture__00000000_.Roles, __Role__00000000_)
	__Architecture__00000000_.Roles = append(__Architecture__00000000_.Roles, __Role__00000001_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000004_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000000_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000002_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000003_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000005_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000012_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000000_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000001_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000002_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000003_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000004_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000005_)
	__Diagram__00000001_.State_Shapes = append(__Diagram__00000001_.State_Shapes, __StateShape__00000010_)
	__Diagram__00000001_.State_Shapes = append(__Diagram__00000001_.State_Shapes, __StateShape__00000006_)
	__Diagram__00000001_.State_Shapes = append(__Diagram__00000001_.State_Shapes, __StateShape__00000007_)
	__Diagram__00000001_.State_Shapes = append(__Diagram__00000001_.State_Shapes, __StateShape__00000008_)
	__Diagram__00000001_.State_Shapes = append(__Diagram__00000001_.State_Shapes, __StateShape__00000009_)
	__Diagram__00000001_.State_Shapes = append(__Diagram__00000001_.State_Shapes, __StateShape__00000011_)
	__Diagram__00000001_.Transition_Shapes = append(__Diagram__00000001_.Transition_Shapes, __Transition_Shape__00000006_)
	__Diagram__00000001_.Transition_Shapes = append(__Diagram__00000001_.Transition_Shapes, __Transition_Shape__00000007_)
	__Diagram__00000001_.Transition_Shapes = append(__Diagram__00000001_.Transition_Shapes, __Transition_Shape__00000008_)
	__Diagram__00000001_.Transition_Shapes = append(__Diagram__00000001_.Transition_Shapes, __Transition_Shape__00000009_)
	__Diagram__00000001_.Transition_Shapes = append(__Diagram__00000001_.Transition_Shapes, __Transition_Shape__00000010_)
	__Diagram__00000001_.Transition_Shapes = append(__Diagram__00000001_.Transition_Shapes, __Transition_Shape__00000011_)
	__Message__00000003_.MessageType = __MessageType__00000000_
	__Message__00000003_.OriginTransition = __Transition__00000010_
	__Message__00000004_.MessageType = __MessageType__00000000_
	__Message__00000004_.OriginTransition = __Transition__00000010_
	__Object__00000001_.State = __State__00000001_
	__Object__00000003_.State = __State__00000007_
	__Object__00000003_.Messages = append(__Object__00000003_.Messages, __Message__00000003_)
	__Object__00000003_.Messages = append(__Object__00000003_.Messages, __Message__00000004_)
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
	__State__00000003_.Entry = nil
	__State__00000003_.Exit = nil
	__State__00000004_.Parent = nil
	__State__00000004_.SubStates = append(__State__00000004_.SubStates, __State__00000001_)
	__State__00000004_.SubStates = append(__State__00000004_.SubStates, __State__00000002_)
	__State__00000004_.SubStates = append(__State__00000004_.SubStates, __State__00000003_)
	__State__00000004_.Diagrams = append(__State__00000004_.Diagrams, __Diagram__00000000_)
	__State__00000004_.Entry = nil
	__State__00000004_.Exit = nil
	__State__00000005_.Parent = nil
	__State__00000005_.Diagrams = append(__State__00000005_.Diagrams, __Diagram__00000000_)
	__State__00000005_.Entry = nil
	__State__00000005_.Exit = nil
	__State__00000006_.Parent = nil
	__State__00000006_.Diagrams = append(__State__00000006_.Diagrams, __Diagram__00000001_)
	__State__00000006_.Entry = nil
	__State__00000006_.Exit = nil
	__State__00000007_.Parent = nil
	__State__00000007_.Diagrams = append(__State__00000007_.Diagrams, __Diagram__00000001_)
	__State__00000007_.Entry = nil
	__State__00000007_.Exit = nil
	__State__00000008_.Parent = nil
	__State__00000008_.Diagrams = append(__State__00000008_.Diagrams, __Diagram__00000001_)
	__State__00000008_.Entry = nil
	__State__00000008_.Exit = nil
	__State__00000009_.Parent = nil
	__State__00000009_.Diagrams = append(__State__00000009_.Diagrams, __Diagram__00000001_)
	__State__00000009_.Entry = nil
	__State__00000009_.Exit = nil
	__State__00000010_.Parent = nil
	__State__00000010_.SubStates = append(__State__00000010_.SubStates, __State__00000008_)
	__State__00000010_.SubStates = append(__State__00000010_.SubStates, __State__00000009_)
	__State__00000010_.SubStates = append(__State__00000010_.SubStates, __State__00000007_)
	__State__00000010_.Diagrams = append(__State__00000010_.Diagrams, __Diagram__00000001_)
	__State__00000010_.Entry = nil
	__State__00000010_.Exit = nil
	__State__00000011_.Parent = nil
	__State__00000011_.Diagrams = append(__State__00000011_.Diagrams, __Diagram__00000001_)
	__State__00000011_.Entry = nil
	__State__00000011_.Exit = nil
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000000_)
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000001_)
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000002_)
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000003_)
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000004_)
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000005_)
	__StateMachine__00000000_.Diagrams = append(__StateMachine__00000000_.Diagrams, __Diagram__00000000_)
	__StateMachine__00000000_.InitialState = __State__00000000_
	__StateMachine__00000001_.States = append(__StateMachine__00000001_.States, __State__00000006_)
	__StateMachine__00000001_.States = append(__StateMachine__00000001_.States, __State__00000007_)
	__StateMachine__00000001_.States = append(__StateMachine__00000001_.States, __State__00000008_)
	__StateMachine__00000001_.States = append(__StateMachine__00000001_.States, __State__00000009_)
	__StateMachine__00000001_.States = append(__StateMachine__00000001_.States, __State__00000010_)
	__StateMachine__00000001_.States = append(__StateMachine__00000001_.States, __State__00000011_)
	__StateMachine__00000001_.Diagrams = append(__StateMachine__00000001_.Diagrams, __Diagram__00000001_)
	__StateMachine__00000001_.InitialState = __State__00000006_
	__StateShape__00000000_.State = __State__00000001_
	__StateShape__00000002_.State = __State__00000002_
	__StateShape__00000003_.State = __State__00000003_
	__StateShape__00000004_.State = __State__00000004_
	__StateShape__00000005_.State = __State__00000005_
	__StateShape__00000006_.State = __State__00000007_
	__StateShape__00000007_.State = __State__00000006_
	__StateShape__00000008_.State = __State__00000008_
	__StateShape__00000009_.State = __State__00000009_
	__StateShape__00000010_.State = __State__00000010_
	__StateShape__00000011_.State = __State__00000011_
	__StateShape__00000012_.State = __State__00000000_
	__Transition__00000000_.Start = __State__00000000_
	__Transition__00000000_.End = __State__00000001_
	__Transition__00000000_.RolesWithPermissions = append(__Transition__00000000_.RolesWithPermissions, __Role__00000000_)
	__Transition__00000000_.Guard = nil
	__Transition__00000000_.Diagrams = append(__Transition__00000000_.Diagrams, __Diagram__00000000_)
	__Transition__00000001_.Start = __State__00000001_
	__Transition__00000001_.End = __State__00000002_
	__Transition__00000001_.RolesWithPermissions = append(__Transition__00000001_.RolesWithPermissions, __Role__00000001_)
	__Transition__00000001_.Guard = nil
	__Transition__00000001_.Diagrams = append(__Transition__00000001_.Diagrams, __Diagram__00000000_)
	__Transition__00000002_.Start = __State__00000002_
	__Transition__00000002_.End = __State__00000003_
	__Transition__00000002_.RolesWithPermissions = append(__Transition__00000002_.RolesWithPermissions, __Role__00000001_)
	__Transition__00000002_.Guard = nil
	__Transition__00000002_.Diagrams = append(__Transition__00000002_.Diagrams, __Diagram__00000000_)
	__Transition__00000003_.Start = __State__00000003_
	__Transition__00000003_.End = __State__00000001_
	__Transition__00000003_.RolesWithPermissions = append(__Transition__00000003_.RolesWithPermissions, __Role__00000001_)
	__Transition__00000003_.Guard = nil
	__Transition__00000003_.Diagrams = append(__Transition__00000003_.Diagrams, __Diagram__00000000_)
	__Transition__00000004_.Start = __State__00000004_
	__Transition__00000004_.End = __State__00000005_
	__Transition__00000004_.RolesWithPermissions = append(__Transition__00000004_.RolesWithPermissions, __Role__00000000_)
	__Transition__00000004_.Guard = nil
	__Transition__00000004_.Diagrams = append(__Transition__00000004_.Diagrams, __Diagram__00000000_)
	__Transition__00000005_.Start = __State__00000005_
	__Transition__00000005_.End = __State__00000001_
	__Transition__00000005_.RolesWithPermissions = append(__Transition__00000005_.RolesWithPermissions, __Role__00000000_)
	__Transition__00000005_.Guard = nil
	__Transition__00000005_.Diagrams = append(__Transition__00000005_.Diagrams, __Diagram__00000000_)
	__Transition__00000006_.Start = __State__00000006_
	__Transition__00000006_.End = __State__00000007_
	__Transition__00000006_.RolesWithPermissions = append(__Transition__00000006_.RolesWithPermissions, __Role__00000000_)
	__Transition__00000006_.Guard = nil
	__Transition__00000006_.Diagrams = append(__Transition__00000006_.Diagrams, __Diagram__00000001_)
	__Transition__00000007_.Start = __State__00000007_
	__Transition__00000007_.End = __State__00000008_
	__Transition__00000007_.RolesWithPermissions = append(__Transition__00000007_.RolesWithPermissions, __Role__00000001_)
	__Transition__00000007_.Guard = nil
	__Transition__00000007_.Diagrams = append(__Transition__00000007_.Diagrams, __Diagram__00000001_)
	__Transition__00000008_.Start = __State__00000008_
	__Transition__00000008_.End = __State__00000009_
	__Transition__00000008_.RolesWithPermissions = append(__Transition__00000008_.RolesWithPermissions, __Role__00000001_)
	__Transition__00000008_.Guard = nil
	__Transition__00000008_.Diagrams = append(__Transition__00000008_.Diagrams, __Diagram__00000001_)
	__Transition__00000009_.Start = __State__00000010_
	__Transition__00000009_.End = __State__00000011_
	__Transition__00000009_.RolesWithPermissions = append(__Transition__00000009_.RolesWithPermissions, __Role__00000000_)
	__Transition__00000009_.Guard = nil
	__Transition__00000009_.Diagrams = append(__Transition__00000009_.Diagrams, __Diagram__00000001_)
	__Transition__00000010_.Start = __State__00000011_
	__Transition__00000010_.End = __State__00000007_
	__Transition__00000010_.RolesWithPermissions = append(__Transition__00000010_.RolesWithPermissions, __Role__00000000_)
	__Transition__00000010_.GeneratedMessages = append(__Transition__00000010_.GeneratedMessages, __MessageType__00000000_)
	__Transition__00000010_.Guard = nil
	__Transition__00000010_.Diagrams = append(__Transition__00000010_.Diagrams, __Diagram__00000001_)
	__Transition__00000011_.Start = __State__00000009_
	__Transition__00000011_.End = __State__00000007_
	__Transition__00000011_.RolesWithPermissions = append(__Transition__00000011_.RolesWithPermissions, __Role__00000001_)
	__Transition__00000011_.Guard = nil
	__Transition__00000011_.Diagrams = append(__Transition__00000011_.Diagrams, __Diagram__00000001_)
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
