// generated code - do not edit
package models

import "time"

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !IsStagedPointerToGongstruct(stage, *element) {
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
	modified = GongCleanPointer(stage, &actorstateshape.ActorState) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ActorStateTransition
func (actorstatetransition *ActorStateTransition) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &actorstatetransition.Justifications) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &actorstatetransition.StartState) || modified
	modified = GongCleanPointer(stage, &actorstatetransition.EndState) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ActorStateTransitionShape
func (actorstatetransitionshape *ActorStateTransitionShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &actorstatetransitionshape.ControlPointShapes) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &actorstatetransitionshape.ActorStateTransition) || modified
	modified = GongCleanPointer(stage, &actorstatetransitionshape.Start) || modified
	modified = GongCleanPointer(stage, &actorstatetransitionshape.End) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Analysis
func (analysis *Analysis) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &analysis.Scenarios) || modified
	modified = GongCleanSlice(stage, &analysis.GroupUse) || modified
	modified = GongCleanSlice(stage, &analysis.GeoObjectUse) || modified
	modified = GongCleanSlice(stage, &analysis.MapUse) || modified
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
	modified = GongCleanSlice(stage, &diagram.EvolutionDirectionShapes) || modified
	modified = GongCleanSlice(stage, &diagram.EvolutionDirectionsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagram.ActorStateShapes) || modified
	modified = GongCleanSlice(stage, &diagram.ActorStatesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagram.ParameterShapes) || modified
	modified = GongCleanSlice(stage, &diagram.ParametersWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagram.ScenarioParameterShapes) || modified
	modified = GongCleanSlice(stage, &diagram.ParametersAggregatesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagram.ActorStateTransitionShapes) || modified
	modified = GongCleanSlice(stage, &diagram.ActorStateTransitionsWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Document
func (document *Document) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &document.GeoObjectUse) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DocumentUse
func (documentuse *DocumentUse) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &documentuse.Document) || modified
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
	modified = GongCleanPointer(stage, &evolutiondirectionshape.EvolutionDirection) || modified
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
	modified = GongCleanPointer(stage, &geoobjectuse.GeoObject) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Group
func (group *Group) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &group.UserUse) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GroupUse
func (groupuse *GroupUse) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &groupuse.Group) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Library
func (library *Library) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &library.Analyses) || modified
	modified = GongCleanSlice(stage, &library.SubLibraries) || modified
	modified = GongCleanSlice(stage, &library.SubLibrariesWhoseNodeIsExpanded) || modified
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
	modified = GongCleanPointer(stage, &mapobjectuse.Map) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Parameter
func (parameter *Parameter) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &parameter.GroupUse) || modified
	modified = GongCleanSlice(stage, &parameter.DocumentUse) || modified
	modified = GongCleanSlice(stage, &parameter.GeoObjectUse) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ParameterCategory
func (parametercategory *ParameterCategory) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &parametercategory.ParameterUse) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ParameterCategoryUse
func (parametercategoryuse *ParameterCategoryUse) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &parametercategoryuse.ParameterCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ParameterShape
func (parametershape *ParameterShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &parametershape.Parameter) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ParametersAggregate
func (parametersaggregate *ParametersAggregate) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &parametersaggregate.Parameters) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ParametersAggregateShape
func (parametersaggregateshape *ParametersAggregateShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &parametersaggregateshape.ScenarioParameter) || modified
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
	modified = GongCleanSlice(stage, &repository.ParameterUse) || modified
	modified = GongCleanSlice(stage, &repository.GroupUse) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Scenario
func (scenario *Scenario) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &scenario.Diagrams) || modified
	modified = GongCleanSlice(stage, &scenario.ActorStates) || modified
	modified = GongCleanSlice(stage, &scenario.ActorStateTransitions) || modified
	modified = GongCleanSlice(stage, &scenario.EvolutionDirections) || modified
	modified = GongCleanSlice(stage, &scenario.Parameters) || modified
	modified = GongCleanSlice(stage, &scenario.ParametersAggretates) || modified
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
	modified = GongCleanPointer(stage, &useruse.User) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Workspace
func (workspace *Workspace) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &workspace.SelectedDiagram) || modified
	modified = GongCleanPointer(stage, &workspace.Default_EvolutionDirectionShape) || modified
	modified = GongCleanPointer(stage, &workspace.Default_ParameterShape) || modified
	modified = GongCleanPointer(stage, &workspace.Default_ScenarioParameterShape) || modified
	modified = GongCleanPointer(stage, &workspace.Default_ActorStateShape) || modified
	modified = GongCleanPointer(stage, &workspace.Default_ActorStateTransitionShape) || modified
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
