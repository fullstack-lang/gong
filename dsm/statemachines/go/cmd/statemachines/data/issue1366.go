package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/statemachines/go/models"
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

	__Architecture__00000000_ := (&models.Architecture{Name: `Architecture`}).Stage(stage)

	__Diagram__00000000_ := (&models.Diagram{Name: `New Diagram`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `Issue 1366, statemachines: problem with text layout`}).Stage(stage)

	__State__00000000_ := (&models.State{Name: `Start State, with a long name to show that there is a layout issue`}).Stage(stage)
	__State__00000001_ := (&models.State{Name: `Decision node, with a long name to show that there is a layout issue`}).Stage(stage)
	__State__00000002_ := (&models.State{Name: `A standard state, Start State, with a long name to show that there is no layout issue with the kind of state
`}).Stage(stage)
	__State__00000003_ := (&models.State{Name: `End state, Start State, with a long name to show that there is a layout issue`}).Stage(stage)
	__State__00000004_ := (&models.State{Name: `A standard state, Start State, with a long name to show that there is no layout issue with the kind of state

`}).Stage(stage)

	__StateMachine__00000000_ := (&models.StateMachine{Name: `Dummy`}).Stage(stage)

	__StateShape__00000000_ := (&models.StateShape{Name: `Start State, with a long name to show that there is a layout issue`}).Stage(stage)
	__StateShape__00000001_ := (&models.StateShape{Name: `Decision node, with a long name to show that there is a layout issue`}).Stage(stage)
	__StateShape__00000002_ := (&models.StateShape{Name: `A standard state, Start State, with a long name to show that there is no layout issue with the kind of state
`}).Stage(stage)
	__StateShape__00000003_ := (&models.StateShape{Name: `End state, Start State, with a long name to show that there is a layout issue`}).Stage(stage)
	__StateShape__00000004_ := (&models.StateShape{Name: `A standard state, Start State, with a long name to show that there is no layout issue with the kind of state

`}).Stage(stage)

	__Transition__00000000_ := (&models.Transition{Name: ``}).Stage(stage)
	__Transition__00000003_ := (&models.Transition{Name: ``}).Stage(stage)
	__Transition__00000004_ := (&models.Transition{Name: ``}).Stage(stage)

	__Transition_Shape__00000000_ := (&models.Transition_Shape{Name: `Start State, with a long name to show that there is a layout issue to Decision node, with a long name to show that there is a layout issue`}).Stage(stage)
	__Transition_Shape__00000003_ := (&models.Transition_Shape{Name: `Decision node, with a long name to show that there is a layout issue to End state, Start State, with a long name to show that there is a layout issue`}).Stage(stage)
	__Transition_Shape__00000004_ := (&models.Transition_Shape{Name: `A standard state, Start State, with a long name to show that there is no layout issue with the kind of state
 to A standard state, Start State, with a long name to show that there is no layout issue with the kind of state

`}).Stage(stage)

	// insertion point for initialization of values

	__Architecture__00000000_.Name = `Architecture`
	__Architecture__00000000_.NbPixPerCharacter = 8.000000

	__Diagram__00000000_.Name = `New Diagram`
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsStatesNodeExpanded = true
	__Diagram__00000000_.IsNotesNodeExpanded = false

	__Library__00000000_.Name = `Issue 1366, statemachines: problem with text layout`
	__Library__00000000_.NbPixPerCharacter = 0.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsExpanded = false
	__Library__00000000_.LayoutDirection = models.Vertical
	__Library__00000000_.IsRootLibrary = true
	__Library__00000000_.IsStateMachinesNodeExpanded = true
	__Library__00000000_.IsNotesNodeExpanded = false
	__Library__00000000_.IsSubLibrariesNodeExpanded = false
	__Library__00000000_.IsExpandedTmp = true

	__State__00000000_.Name = `Start State, with a long name to show that there is a layout issue`
	__State__00000000_.IsDecisionNode = false
	__State__00000000_.IsFictious = false
	__State__00000000_.IsEndState = false

	__State__00000001_.Name = `Decision node, with a long name to show that there is a layout issue`
	__State__00000001_.IsDecisionNode = true
	__State__00000001_.IsFictious = false
	__State__00000001_.IsEndState = false

	__State__00000002_.Name = `A standard state, Start State, with a long name to show that there is no layout issue with the kind of state
`
	__State__00000002_.IsDecisionNode = false
	__State__00000002_.IsFictious = false
	__State__00000002_.IsEndState = false

	__State__00000003_.Name = `End state, Start State, with a long name to show that there is a layout issue`
	__State__00000003_.IsDecisionNode = false
	__State__00000003_.IsFictious = false
	__State__00000003_.IsEndState = true

	__State__00000004_.Name = `A standard state, Start State, with a long name to show that there is no layout issue with the kind of state

`
	__State__00000004_.IsDecisionNode = false
	__State__00000004_.IsFictious = false
	__State__00000004_.IsEndState = false

	__StateMachine__00000000_.Name = `Dummy`
	__StateMachine__00000000_.ComputedPrefix = ``
	__StateMachine__00000000_.IsExpanded = false
	__StateMachine__00000000_.LayoutDirection = models.Vertical

	__StateShape__00000000_.Name = `Start State, with a long name to show that there is a layout issue`
	__StateShape__00000000_.X = 640.000000
	__StateShape__00000000_.Y = 123.000000
	__StateShape__00000000_.Width = 187.000000
	__StateShape__00000000_.Height = 20.000000
	__StateShape__00000000_.IsHidden = false

	__StateShape__00000001_.Name = `Decision node, with a long name to show that there is a layout issue`
	__StateShape__00000001_.X = 408.000000
	__StateShape__00000001_.Y = 283.000000
	__StateShape__00000001_.Width = 200.000000
	__StateShape__00000001_.Height = 50.000000
	__StateShape__00000001_.IsHidden = false

	__StateShape__00000002_.Name = `A standard state, Start State, with a long name to show that there is no layout issue with the kind of state
`
	__StateShape__00000002_.X = 207.000000
	__StateShape__00000002_.Y = 546.000000
	__StateShape__00000002_.Width = 200.000000
	__StateShape__00000002_.Height = 118.000000
	__StateShape__00000002_.IsHidden = false

	__StateShape__00000003_.Name = `End state, Start State, with a long name to show that there is a layout issue`
	__StateShape__00000003_.X = 303.000000
	__StateShape__00000003_.Y = 437.000000
	__StateShape__00000003_.Width = 200.000000
	__StateShape__00000003_.Height = 36.000000
	__StateShape__00000003_.IsHidden = false

	__StateShape__00000004_.Name = `A standard state, Start State, with a long name to show that there is no layout issue with the kind of state

`
	__StateShape__00000004_.X = 631.000000
	__StateShape__00000004_.Y = 529.000000
	__StateShape__00000004_.Width = 200.000000
	__StateShape__00000004_.Height = 122.000000
	__StateShape__00000004_.IsHidden = false

	__Transition__00000000_.Name = ``

	__Transition__00000003_.Name = ``

	__Transition__00000004_.Name = ``

	__Transition_Shape__00000000_.Name = `Start State, with a long name to show that there is a layout issue to Decision node, with a long name to show that there is a layout issue`
	__Transition_Shape__00000000_.StartRatio = 0.500000
	__Transition_Shape__00000000_.EndRatio = 0.867510
	__Transition_Shape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__00000000_.CornerOffsetRatio = 6.550488
	__Transition_Shape__00000000_.IsHidden = false

	__Transition_Shape__00000003_.Name = `Decision node, with a long name to show that there is a layout issue to End state, Start State, with a long name to show that there is a layout issue`
	__Transition_Shape__00000003_.StartRatio = 0.420195
	__Transition_Shape__00000003_.EndRatio = 0.500000
	__Transition_Shape__00000003_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000003_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000003_.CornerOffsetRatio = -0.947490
	__Transition_Shape__00000003_.IsHidden = false

	__Transition_Shape__00000004_.Name = `A standard state, Start State, with a long name to show that there is no layout issue with the kind of state
 to A standard state, Start State, with a long name to show that there is no layout issue with the kind of state

`
	__Transition_Shape__00000004_.StartRatio = 0.500000
	__Transition_Shape__00000004_.EndRatio = 0.647621
	__Transition_Shape__00000004_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000004_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__00000004_.CornerOffsetRatio = 1.200000
	__Transition_Shape__00000004_.IsHidden = false

	// insertion point for setup of pointers
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000000_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000001_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000002_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000003_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000004_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000000_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000003_)
	__Diagram__00000000_.Transition_Shapes = append(__Diagram__00000000_.Transition_Shapes, __Transition_Shape__00000004_)
	__Library__00000000_.RootStateMachines = append(__Library__00000000_.RootStateMachines, __StateMachine__00000000_)
	__Library__00000000_.StateMachinesWhoseNodeIsExpanded = append(__Library__00000000_.StateMachinesWhoseNodeIsExpanded, __StateMachine__00000000_)
	__State__00000000_.Parent = nil
	__State__00000000_.Entry = nil
	__State__00000000_.Exit = nil
	__State__00000000_.Diagrams = append(__State__00000000_.Diagrams, __Diagram__00000000_)
	__State__00000001_.Parent = nil
	__State__00000001_.Entry = nil
	__State__00000001_.Exit = nil
	__State__00000001_.Diagrams = append(__State__00000001_.Diagrams, __Diagram__00000000_)
	__State__00000002_.Parent = nil
	__State__00000002_.Entry = nil
	__State__00000002_.Exit = nil
	__State__00000002_.Diagrams = append(__State__00000002_.Diagrams, __Diagram__00000000_)
	__State__00000003_.Parent = nil
	__State__00000003_.Entry = nil
	__State__00000003_.Exit = nil
	__State__00000003_.Diagrams = append(__State__00000003_.Diagrams, __Diagram__00000000_)
	__State__00000004_.Parent = nil
	__State__00000004_.Entry = nil
	__State__00000004_.Exit = nil
	__State__00000004_.Diagrams = append(__State__00000004_.Diagrams, __Diagram__00000000_)
	__StateMachine__00000000_.InitialState = __State__00000000_
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000000_)
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000001_)
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000002_)
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000003_)
	__StateMachine__00000000_.States = append(__StateMachine__00000000_.States, __State__00000004_)
	__StateMachine__00000000_.Diagrams = append(__StateMachine__00000000_.Diagrams, __Diagram__00000000_)
	__StateShape__00000000_.State = __State__00000000_
	__StateShape__00000001_.State = __State__00000001_
	__StateShape__00000002_.State = __State__00000002_
	__StateShape__00000003_.State = __State__00000003_
	__StateShape__00000004_.State = __State__00000004_
	__Transition__00000000_.Start = __State__00000000_
	__Transition__00000000_.End = __State__00000001_
	__Transition__00000000_.Guard = nil
	__Transition__00000000_.Diagrams = append(__Transition__00000000_.Diagrams, __Diagram__00000000_)
	__Transition__00000003_.Start = __State__00000001_
	__Transition__00000003_.End = __State__00000003_
	__Transition__00000003_.Guard = nil
	__Transition__00000003_.Diagrams = append(__Transition__00000003_.Diagrams, __Diagram__00000000_)
	__Transition__00000004_.Start = __State__00000002_
	__Transition__00000004_.End = __State__00000004_
	__Transition__00000004_.Guard = nil
	__Transition_Shape__00000000_.Transition = __Transition__00000000_
	__Transition_Shape__00000003_.Transition = __Transition__00000003_
	__Transition_Shape__00000004_.Transition = __Transition__00000004_
}
