package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/statemachines/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time
var _ = slices.Index[[]int, int]

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__Activities__00000000_ := (&models.Activities{Name: `Show Muti Checkbox`}).Stage(stage)
	__Activities__00000001_ := (&models.Activities{Name: `Hide Multi Checkbox`}).Stage(stage)
	__Activities__00000002_ := (&models.Activities{Name: `Fill Empty Form`}).Stage(stage)

	__Architecture__00000000_ := (&models.Architecture{Name: `Architecture`}).Stage(stage)

	__Diagram__00000000_ := (&models.Diagram{Name: `New Diagram`}).Stage(stage)
	__Diagram__00000001_ := (&models.Diagram{Name: `New Diagram`}).Stage(stage)

	__State__00000000_ := (&models.State{Name: `Single Item Selectiond`}).Stage(stage)
	__State__00000001_ := (&models.State{Name: `Multi Item Selection`}).Stage(stage)

	__StateMachine__00000001_ := (&models.StateMachine{Name: `New StateMachine`}).Stage(stage)

	__StateShape__00000000_ := (&models.StateShape{Name: `Single Item Selectiond`}).Stage(stage)
	__StateShape__00000001_ := (&models.StateShape{Name: `Multi Item Selection`}).Stage(stage)

	__Transition__00000000_ := (&models.Transition{Name: `Click Multi Edit`}).Stage(stage)
	__Transition__00000001_ := (&models.Transition{Name: `Click Muti Edit Off`}).Stage(stage)

	__Transition_Shape__00000000_ := (&models.Transition_Shape{Name: `Single Item Selection to New State`}).Stage(stage)
	__Transition_Shape__00000001_ := (&models.Transition_Shape{Name: `Multi Item Selection to Single Item Selection`}).Stage(stage)

	// insertion point for initialization of values

	__Activities__00000000_.Name = `Show Muti Checkbox`
	__Activities__00000000_.Criticality = ""

	__Activities__00000001_.Name = `Hide Multi Checkbox`
	__Activities__00000001_.Criticality = ""

	__Activities__00000002_.Name = `Fill Empty Form`
	__Activities__00000002_.Criticality = ""

	__Architecture__00000000_.Name = `Architecture`
	__Architecture__00000000_.NbPixPerCharacter = 8.000000

	__Diagram__00000000_.Name = `New Diagram`
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsInRenameMode = false

	__Diagram__00000001_.Name = `New Diagram`
	__Diagram__00000001_.IsChecked = false
	__Diagram__00000001_.IsExpanded = false
	__Diagram__00000001_.IsEditable_ = true
	__Diagram__00000001_.IsInRenameMode = false

	__State__00000000_.Name = `Single Item Selectiond`
	__State__00000000_.IsDecisionNode = false
	__State__00000000_.IsFictious = false
	__State__00000000_.IsEndState = false
	__State__00000000_.IsInRenameMode = false

	__State__00000001_.Name = `Multi Item Selection`
	__State__00000001_.IsDecisionNode = false
	__State__00000001_.IsFictious = false
	__State__00000001_.IsEndState = false
	__State__00000001_.IsInRenameMode = false

	__StateMachine__00000001_.Name = `New StateMachine`
	__StateMachine__00000001_.IsNodeExpanded = true

	__StateShape__00000000_.Name = `Single Item Selectiond`
	__StateShape__00000000_.X = 89.000000
	__StateShape__00000000_.Y = 217.000000
	__StateShape__00000000_.Width = 200.000000
	__StateShape__00000000_.Height = 157.000000
	__StateShape__00000000_.IsHidden = false

	__StateShape__00000001_.Name = `Multi Item Selection`
	__StateShape__00000001_.X = 626.000000
	__StateShape__00000001_.Y = 438.000016
	__StateShape__00000001_.Width = 200.000000
	__StateShape__00000001_.Height = 161.000000
	__StateShape__00000001_.IsHidden = false

	__Transition__00000000_.Name = `Click Multi Edit`
	__Transition__00000000_.IsInRenameMode = false

	__Transition__00000001_.Name = `Click Muti Edit Off`
	__Transition__00000001_.IsInRenameMode = false

	__Transition_Shape__00000000_.Name = `Single Item Selection to New State`
	__Transition_Shape__00000000_.StartRatio = 0.235713
	__Transition_Shape__00000000_.EndRatio = 0.671415
	__Transition_Shape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000000_.CornerOffsetRatio = 1.401318
	__Transition_Shape__00000000_.IsHidden = false

	__Transition_Shape__00000001_.Name = `Multi Item Selection to Single Item Selection`
	__Transition_Shape__00000001_.StartRatio = 0.782652
	__Transition_Shape__00000001_.EndRatio = 0.486415
	__Transition_Shape__00000001_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000001_.CornerOffsetRatio = -0.378839
	__Transition_Shape__00000001_.IsHidden = false

	// insertion point for setup of pointers
	__Architecture__00000000_.StateMachines = append(__Architecture__00000000_.StateMachines, __StateMachine__00000001_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000000_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000001_)
	__Diagram__00000000_.StatesWhoseNodeIsExpanded = append(__Diagram__00000000_.StatesWhoseNodeIsExpanded, __State__00000000_)
	__Diagram__00000000_.StatesWhoseNodeIsExpanded = append(__Diagram__00000000_.StatesWhoseNodeIsExpanded, __State__00000001_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000000_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000001_)
	__State__00000000_.Parent = nil
	__State__00000000_.Diagrams = append(__State__00000000_.Diagrams, __Diagram__00000000_)
	__State__00000000_.Entry = nil
	__State__00000000_.Activities = append(__State__00000000_.Activities, __Activities__00000001_)
	__State__00000000_.Exit = nil
	__State__00000001_.Parent = nil
	__State__00000001_.Diagrams = append(__State__00000001_.Diagrams, __Diagram__00000000_)
	__State__00000001_.Entry = nil
	__State__00000001_.Activities = append(__State__00000001_.Activities, __Activities__00000000_)
	__State__00000001_.Activities = append(__State__00000001_.Activities, __Activities__00000002_)
	__State__00000001_.Exit = nil
	__StateMachine__00000001_.States = append(__StateMachine__00000001_.States, __State__00000000_)
	__StateMachine__00000001_.States = append(__StateMachine__00000001_.States, __State__00000001_)
	__StateMachine__00000001_.Diagrams = append(__StateMachine__00000001_.Diagrams, __Diagram__00000000_)
	__StateMachine__00000001_.Diagrams = append(__StateMachine__00000001_.Diagrams, __Diagram__00000001_)
	__StateMachine__00000001_.InitialState = nil
	__StateShape__00000000_.State = __State__00000000_
	__StateShape__00000001_.State = __State__00000001_
	__Transition__00000000_.Start = __State__00000000_
	__Transition__00000000_.End = __State__00000001_
	__Transition__00000000_.Guard = nil
	__Transition__00000000_.Diagrams = append(__Transition__00000000_.Diagrams, __Diagram__00000000_)
	__Transition__00000001_.Start = __State__00000001_
	__Transition__00000001_.End = __State__00000000_
	__Transition__00000001_.Guard = nil
	__Transition__00000001_.Diagrams = append(__Transition__00000001_.Diagrams, __Diagram__00000000_)
	__Transition_Shape__00000000_.Transition = __Transition__00000000_
	__Transition_Shape__00000001_.Transition = __Transition__00000001_
}
