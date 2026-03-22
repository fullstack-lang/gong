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

	__Architecture__00000000_ := (&models.Architecture{Name: `Architecture`}).Stage(stage)

	__Diagram__00000000_ := (&models.Diagram{Name: `New Diagram`}).Stage(stage)

	__State__00000000_ := (&models.State{Name: `Single Item Selection`}).Stage(stage)
	__State__00000001_ := (&models.State{Name: `New State`}).Stage(stage)

	__StateMachine__00000001_ := (&models.StateMachine{Name: `New StateMachine`}).Stage(stage)

	__StateShape__00000000_ := (&models.StateShape{Name: `Single Item Selection`}).Stage(stage)
	__StateShape__00000001_ := (&models.StateShape{Name: `New State`}).Stage(stage)

	// insertion point for initialization of values

	__Architecture__00000000_.Name = `Architecture`
	__Architecture__00000000_.NbPixPerCharacter = 8.000000

	__Diagram__00000000_.Name = `New Diagram`
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsInRenameMode = false

	__State__00000000_.Name = `Single Item Selection`
	__State__00000000_.IsDecisionNode = false
	__State__00000000_.IsFictif = false
	__State__00000000_.IsEndState = false

	__State__00000001_.Name = `New State`
	__State__00000001_.IsDecisionNode = false
	__State__00000001_.IsFictif = false
	__State__00000001_.IsEndState = false

	__StateMachine__00000001_.Name = `New StateMachine`
	__StateMachine__00000001_.IsNodeExpanded = true

	__StateShape__00000000_.Name = `Single Item Selection`
	__StateShape__00000000_.IsExpanded = false
	__StateShape__00000000_.X = 89.000000
	__StateShape__00000000_.Y = 217.000000
	__StateShape__00000000_.Width = 200.000000
	__StateShape__00000000_.Height = 157.000000

	__StateShape__00000001_.Name = `New State`
	__StateShape__00000001_.IsExpanded = false
	__StateShape__00000001_.X = 508.000000
	__StateShape__00000001_.Y = 215.999985
	__StateShape__00000001_.Width = 200.000000
	__StateShape__00000001_.Height = 161.000000

	// insertion point for setup of pointers
	__Architecture__00000000_.StateMachines = append(__Architecture__00000000_.StateMachines, __StateMachine__00000001_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000000_)
	__Diagram__00000000_.State_Shapes = append(__Diagram__00000000_.State_Shapes, __StateShape__00000001_)
	__State__00000000_.Parent = nil
	__State__00000000_.Diagrams = append(__State__00000000_.Diagrams, __Diagram__00000000_)
	__State__00000000_.Entry = nil
	__State__00000000_.Exit = nil
	__State__00000001_.Parent = nil
	__State__00000001_.Diagrams = append(__State__00000001_.Diagrams, __Diagram__00000000_)
	__State__00000001_.Entry = nil
	__State__00000001_.Exit = nil
	__StateMachine__00000001_.States = append(__StateMachine__00000001_.States, __State__00000000_)
	__StateMachine__00000001_.States = append(__StateMachine__00000001_.States, __State__00000001_)
	__StateMachine__00000001_.Diagrams = append(__StateMachine__00000001_.Diagrams, __Diagram__00000000_)
	__StateMachine__00000001_.InitialState = nil
	__StateShape__00000000_.State = __State__00000000_
	__StateShape__00000001_.State = __State__00000001_
}
