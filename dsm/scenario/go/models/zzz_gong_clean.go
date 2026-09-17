// generated code - do not edit
package models

import "time"

// CleanSlice is the Stage method that removes unstaged elements from a slice of pointers.
func (stage *Stage) CleanSlice[T GongstructPtr](slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if stage.IsStaged(element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// CleanPointer is the Stage method that sets the pointer to nil if the referenced element is not staged.
func (stage *Stage) CleanPointer[T GongstructPtr](element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !stage.IsStaged(*element) {
		*element = zero
		modified = true
		return
	}
	return
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by ActorState
func (actorstate *ActorState) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ActorStateShape
func (actorstateshape *ActorStateShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&actorstateshape.ActorState) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ActorStateTransition
func (actorstatetransition *ActorStateTransition) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&actorstatetransition.Justifications) || modified
	// insertion point per field
	modified = stage.CleanPointer(&actorstatetransition.StartState) || modified
	modified = stage.CleanPointer(&actorstatetransition.EndState) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ActorStateTransitionShape
func (actorstatetransitionshape *ActorStateTransitionShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&actorstatetransitionshape.ControlPointShapes) || modified
	// insertion point per field
	modified = stage.CleanPointer(&actorstatetransitionshape.ActorStateTransition) || modified
	modified = stage.CleanPointer(&actorstatetransitionshape.Start) || modified
	modified = stage.CleanPointer(&actorstatetransitionshape.End) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Analysis
func (analysis *Analysis) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&analysis.Scenarios) || modified
	modified = stage.CleanSlice(&analysis.GroupUse) || modified
	modified = stage.CleanSlice(&analysis.GeoObjectUse) || modified
	modified = stage.CleanSlice(&analysis.MapUse) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ControlPointShape
func (controlpointshape *ControlPointShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Diagram
func (diagram *Diagram) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&diagram.EvolutionDirectionShapes) || modified
	modified = stage.CleanSlice(&diagram.EvolutionDirectionsWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.ActorStateShapes) || modified
	modified = stage.CleanSlice(&diagram.ActorStatesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.ParameterShapes) || modified
	modified = stage.CleanSlice(&diagram.ParametersWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.ScenarioParameterShapes) || modified
	modified = stage.CleanSlice(&diagram.ParametersAggregatesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.ActorStateTransitionShapes) || modified
	modified = stage.CleanSlice(&diagram.ActorStateTransitionsWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Document
func (document *Document) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&document.GeoObjectUse) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DocumentUse
func (documentuse *DocumentUse) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&documentuse.Document) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EvolutionDirection
func (evolutiondirection *EvolutionDirection) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by EvolutionDirectionShape
func (evolutiondirectionshape *EvolutionDirectionShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&evolutiondirectionshape.EvolutionDirection) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Foo
func (foo *Foo) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GeoObject
func (geoobject *GeoObject) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GeoObjectUse
func (geoobjectuse *GeoObjectUse) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&geoobjectuse.GeoObject) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Group
func (group *Group) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&group.UserUse) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GroupUse
func (groupuse *GroupUse) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&groupuse.Group) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Library
func (library *Library) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&library.Analyses) || modified
	modified = stage.CleanSlice(&library.SubLibraries) || modified
	modified = stage.CleanSlice(&library.SubLibrariesWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by MapObject
func (mapobject *MapObject) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by MapObjectUse
func (mapobjectuse *MapObjectUse) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&mapobjectuse.Map) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Parameter
func (parameter *Parameter) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&parameter.GroupUse) || modified
	modified = stage.CleanSlice(&parameter.DocumentUse) || modified
	modified = stage.CleanSlice(&parameter.GeoObjectUse) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ParameterCategory
func (parametercategory *ParameterCategory) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&parametercategory.ParameterUse) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ParameterCategoryUse
func (parametercategoryuse *ParameterCategoryUse) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&parametercategoryuse.ParameterCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ParameterShape
func (parametershape *ParameterShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&parametershape.Parameter) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ParametersAggregate
func (parametersaggregate *ParametersAggregate) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&parametersaggregate.Parameters) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ParametersAggregateShape
func (parametersaggregateshape *ParametersAggregateShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&parametersaggregateshape.ScenarioParameter) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Position
func (position *Position) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Repository
func (repository *Repository) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&repository.ParameterUse) || modified
	modified = stage.CleanSlice(&repository.GroupUse) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Scenario
func (scenario *Scenario) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&scenario.Diagrams) || modified
	modified = stage.CleanSlice(&scenario.ActorStates) || modified
	modified = stage.CleanSlice(&scenario.ActorStateTransitions) || modified
	modified = stage.CleanSlice(&scenario.EvolutionDirections) || modified
	modified = stage.CleanSlice(&scenario.Parameters) || modified
	modified = stage.CleanSlice(&scenario.ParametersAggretates) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by User
func (user *User) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by UserUse
func (useruse *UserUse) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&useruse.User) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Workspace
func (workspace *Workspace) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&workspace.SelectedDiagram) || modified
	modified = stage.CleanPointer(&workspace.Default_EvolutionDirectionShape) || modified
	modified = stage.CleanPointer(&workspace.Default_ParameterShape) || modified
	modified = stage.CleanPointer(&workspace.Default_ScenarioParameterShape) || modified
	modified = stage.CleanPointer(&workspace.Default_ActorStateShape) || modified
	modified = stage.CleanPointer(&workspace.Default_ActorStateTransitionShape) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
