// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *ActorState:
		ok = stage.IsStagedActorState(target)

	case *ActorStateShape:
		ok = stage.IsStagedActorStateShape(target)

	case *ActorStateTransition:
		ok = stage.IsStagedActorStateTransition(target)

	case *ActorStateTransitionShape:
		ok = stage.IsStagedActorStateTransitionShape(target)

	case *Analysis:
		ok = stage.IsStagedAnalysis(target)

	case *ControlPointShape:
		ok = stage.IsStagedControlPointShape(target)

	case *Diagram:
		ok = stage.IsStagedDiagram(target)

	case *Document:
		ok = stage.IsStagedDocument(target)

	case *DocumentUse:
		ok = stage.IsStagedDocumentUse(target)

	case *EvolutionDirection:
		ok = stage.IsStagedEvolutionDirection(target)

	case *EvolutionDirectionShape:
		ok = stage.IsStagedEvolutionDirectionShape(target)

	case *Foo:
		ok = stage.IsStagedFoo(target)

	case *GeoObject:
		ok = stage.IsStagedGeoObject(target)

	case *GeoObjectUse:
		ok = stage.IsStagedGeoObjectUse(target)

	case *Group:
		ok = stage.IsStagedGroup(target)

	case *GroupUse:
		ok = stage.IsStagedGroupUse(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *MapObject:
		ok = stage.IsStagedMapObject(target)

	case *MapObjectUse:
		ok = stage.IsStagedMapObjectUse(target)

	case *Parameter:
		ok = stage.IsStagedParameter(target)

	case *ParameterCategory:
		ok = stage.IsStagedParameterCategory(target)

	case *ParameterCategoryUse:
		ok = stage.IsStagedParameterCategoryUse(target)

	case *ParameterShape:
		ok = stage.IsStagedParameterShape(target)

	case *ParametersAggregate:
		ok = stage.IsStagedParametersAggregate(target)

	case *ParametersAggregateShape:
		ok = stage.IsStagedParametersAggregateShape(target)

	case *Position:
		ok = stage.IsStagedPosition(target)

	case *Repository:
		ok = stage.IsStagedRepository(target)

	case *Scenario:
		ok = stage.IsStagedScenario(target)

	case *User:
		ok = stage.IsStagedUser(target)

	case *UserUse:
		ok = stage.IsStagedUserUse(target)

	case *Workspace:
		ok = stage.IsStagedWorkspace(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *ActorState:
		ok = stage.IsStagedActorState(target)

	case *ActorStateShape:
		ok = stage.IsStagedActorStateShape(target)

	case *ActorStateTransition:
		ok = stage.IsStagedActorStateTransition(target)

	case *ActorStateTransitionShape:
		ok = stage.IsStagedActorStateTransitionShape(target)

	case *Analysis:
		ok = stage.IsStagedAnalysis(target)

	case *ControlPointShape:
		ok = stage.IsStagedControlPointShape(target)

	case *Diagram:
		ok = stage.IsStagedDiagram(target)

	case *Document:
		ok = stage.IsStagedDocument(target)

	case *DocumentUse:
		ok = stage.IsStagedDocumentUse(target)

	case *EvolutionDirection:
		ok = stage.IsStagedEvolutionDirection(target)

	case *EvolutionDirectionShape:
		ok = stage.IsStagedEvolutionDirectionShape(target)

	case *Foo:
		ok = stage.IsStagedFoo(target)

	case *GeoObject:
		ok = stage.IsStagedGeoObject(target)

	case *GeoObjectUse:
		ok = stage.IsStagedGeoObjectUse(target)

	case *Group:
		ok = stage.IsStagedGroup(target)

	case *GroupUse:
		ok = stage.IsStagedGroupUse(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *MapObject:
		ok = stage.IsStagedMapObject(target)

	case *MapObjectUse:
		ok = stage.IsStagedMapObjectUse(target)

	case *Parameter:
		ok = stage.IsStagedParameter(target)

	case *ParameterCategory:
		ok = stage.IsStagedParameterCategory(target)

	case *ParameterCategoryUse:
		ok = stage.IsStagedParameterCategoryUse(target)

	case *ParameterShape:
		ok = stage.IsStagedParameterShape(target)

	case *ParametersAggregate:
		ok = stage.IsStagedParametersAggregate(target)

	case *ParametersAggregateShape:
		ok = stage.IsStagedParametersAggregateShape(target)

	case *Position:
		ok = stage.IsStagedPosition(target)

	case *Repository:
		ok = stage.IsStagedRepository(target)

	case *Scenario:
		ok = stage.IsStagedScenario(target)

	case *User:
		ok = stage.IsStagedUser(target)

	case *UserUse:
		ok = stage.IsStagedUserUse(target)

	case *Workspace:
		ok = stage.IsStagedWorkspace(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedActorState(actorstate *ActorState) (ok bool) {

	_, ok = stage.ActorStates[actorstate]

	return
}

func (stage *Stage) IsStagedActorStateShape(actorstateshape *ActorStateShape) (ok bool) {

	_, ok = stage.ActorStateShapes[actorstateshape]

	return
}

func (stage *Stage) IsStagedActorStateTransition(actorstatetransition *ActorStateTransition) (ok bool) {

	_, ok = stage.ActorStateTransitions[actorstatetransition]

	return
}

func (stage *Stage) IsStagedActorStateTransitionShape(actorstatetransitionshape *ActorStateTransitionShape) (ok bool) {

	_, ok = stage.ActorStateTransitionShapes[actorstatetransitionshape]

	return
}

func (stage *Stage) IsStagedAnalysis(analysis *Analysis) (ok bool) {

	_, ok = stage.Analysiss[analysis]

	return
}

func (stage *Stage) IsStagedControlPointShape(controlpointshape *ControlPointShape) (ok bool) {

	_, ok = stage.ControlPointShapes[controlpointshape]

	return
}

func (stage *Stage) IsStagedDiagram(diagram *Diagram) (ok bool) {

	_, ok = stage.Diagrams[diagram]

	return
}

func (stage *Stage) IsStagedDocument(document *Document) (ok bool) {

	_, ok = stage.Documents[document]

	return
}

func (stage *Stage) IsStagedDocumentUse(documentuse *DocumentUse) (ok bool) {

	_, ok = stage.DocumentUses[documentuse]

	return
}

func (stage *Stage) IsStagedEvolutionDirection(evolutiondirection *EvolutionDirection) (ok bool) {

	_, ok = stage.EvolutionDirections[evolutiondirection]

	return
}

func (stage *Stage) IsStagedEvolutionDirectionShape(evolutiondirectionshape *EvolutionDirectionShape) (ok bool) {

	_, ok = stage.EvolutionDirectionShapes[evolutiondirectionshape]

	return
}

func (stage *Stage) IsStagedFoo(foo *Foo) (ok bool) {

	_, ok = stage.Foos[foo]

	return
}

func (stage *Stage) IsStagedGeoObject(geoobject *GeoObject) (ok bool) {

	_, ok = stage.GeoObjects[geoobject]

	return
}

func (stage *Stage) IsStagedGeoObjectUse(geoobjectuse *GeoObjectUse) (ok bool) {

	_, ok = stage.GeoObjectUses[geoobjectuse]

	return
}

func (stage *Stage) IsStagedGroup(group *Group) (ok bool) {

	_, ok = stage.Groups[group]

	return
}

func (stage *Stage) IsStagedGroupUse(groupuse *GroupUse) (ok bool) {

	_, ok = stage.GroupUses[groupuse]

	return
}

func (stage *Stage) IsStagedLibrary(library *Library) (ok bool) {

	_, ok = stage.Librarys[library]

	return
}

func (stage *Stage) IsStagedMapObject(mapobject *MapObject) (ok bool) {

	_, ok = stage.MapObjects[mapobject]

	return
}

func (stage *Stage) IsStagedMapObjectUse(mapobjectuse *MapObjectUse) (ok bool) {

	_, ok = stage.MapObjectUses[mapobjectuse]

	return
}

func (stage *Stage) IsStagedParameter(parameter *Parameter) (ok bool) {

	_, ok = stage.Parameters[parameter]

	return
}

func (stage *Stage) IsStagedParameterCategory(parametercategory *ParameterCategory) (ok bool) {

	_, ok = stage.ParameterCategorys[parametercategory]

	return
}

func (stage *Stage) IsStagedParameterCategoryUse(parametercategoryuse *ParameterCategoryUse) (ok bool) {

	_, ok = stage.ParameterCategoryUses[parametercategoryuse]

	return
}

func (stage *Stage) IsStagedParameterShape(parametershape *ParameterShape) (ok bool) {

	_, ok = stage.ParameterShapes[parametershape]

	return
}

func (stage *Stage) IsStagedParametersAggregate(parametersaggregate *ParametersAggregate) (ok bool) {

	_, ok = stage.ParametersAggregates[parametersaggregate]

	return
}

func (stage *Stage) IsStagedParametersAggregateShape(parametersaggregateshape *ParametersAggregateShape) (ok bool) {

	_, ok = stage.ParametersAggregateShapes[parametersaggregateshape]

	return
}

func (stage *Stage) IsStagedPosition(position *Position) (ok bool) {

	_, ok = stage.Positions[position]

	return
}

func (stage *Stage) IsStagedRepository(repository *Repository) (ok bool) {

	_, ok = stage.Repositorys[repository]

	return
}

func (stage *Stage) IsStagedScenario(scenario *Scenario) (ok bool) {

	_, ok = stage.Scenarios[scenario]

	return
}

func (stage *Stage) IsStagedUser(user *User) (ok bool) {

	_, ok = stage.Users[user]

	return
}

func (stage *Stage) IsStagedUserUse(useruse *UserUse) (ok bool) {

	_, ok = stage.UserUses[useruse]

	return
}

func (stage *Stage) IsStagedWorkspace(workspace *Workspace) (ok bool) {

	_, ok = stage.Workspaces[workspace]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *ActorState:
		stage.StageBranchActorState(target)

	case *ActorStateShape:
		stage.StageBranchActorStateShape(target)

	case *ActorStateTransition:
		stage.StageBranchActorStateTransition(target)

	case *ActorStateTransitionShape:
		stage.StageBranchActorStateTransitionShape(target)

	case *Analysis:
		stage.StageBranchAnalysis(target)

	case *ControlPointShape:
		stage.StageBranchControlPointShape(target)

	case *Diagram:
		stage.StageBranchDiagram(target)

	case *Document:
		stage.StageBranchDocument(target)

	case *DocumentUse:
		stage.StageBranchDocumentUse(target)

	case *EvolutionDirection:
		stage.StageBranchEvolutionDirection(target)

	case *EvolutionDirectionShape:
		stage.StageBranchEvolutionDirectionShape(target)

	case *Foo:
		stage.StageBranchFoo(target)

	case *GeoObject:
		stage.StageBranchGeoObject(target)

	case *GeoObjectUse:
		stage.StageBranchGeoObjectUse(target)

	case *Group:
		stage.StageBranchGroup(target)

	case *GroupUse:
		stage.StageBranchGroupUse(target)

	case *Library:
		stage.StageBranchLibrary(target)

	case *MapObject:
		stage.StageBranchMapObject(target)

	case *MapObjectUse:
		stage.StageBranchMapObjectUse(target)

	case *Parameter:
		stage.StageBranchParameter(target)

	case *ParameterCategory:
		stage.StageBranchParameterCategory(target)

	case *ParameterCategoryUse:
		stage.StageBranchParameterCategoryUse(target)

	case *ParameterShape:
		stage.StageBranchParameterShape(target)

	case *ParametersAggregate:
		stage.StageBranchParametersAggregate(target)

	case *ParametersAggregateShape:
		stage.StageBranchParametersAggregateShape(target)

	case *Position:
		stage.StageBranchPosition(target)

	case *Repository:
		stage.StageBranchRepository(target)

	case *Scenario:
		stage.StageBranchScenario(target)

	case *User:
		stage.StageBranchUser(target)

	case *UserUse:
		stage.StageBranchUserUse(target)

	case *Workspace:
		stage.StageBranchWorkspace(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchActorState(actorstate *ActorState) {

	// check if instance is already staged
	if IsStaged(stage, actorstate) {
		return
	}

	actorstate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchActorStateShape(actorstateshape *ActorStateShape) {

	// check if instance is already staged
	if IsStaged(stage, actorstateshape) {
		return
	}

	actorstateshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if actorstateshape.ActorState != nil {
		StageBranch(stage, actorstateshape.ActorState)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchActorStateTransition(actorstatetransition *ActorStateTransition) {

	// check if instance is already staged
	if IsStaged(stage, actorstatetransition) {
		return
	}

	actorstatetransition.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if actorstatetransition.StartState != nil {
		StageBranch(stage, actorstatetransition.StartState)
	}
	if actorstatetransition.EndState != nil {
		StageBranch(stage, actorstatetransition.EndState)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parameter := range actorstatetransition.Justifications {
		StageBranch(stage, _parameter)
	}

}

func (stage *Stage) StageBranchActorStateTransitionShape(actorstatetransitionshape *ActorStateTransitionShape) {

	// check if instance is already staged
	if IsStaged(stage, actorstatetransitionshape) {
		return
	}

	actorstatetransitionshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if actorstatetransitionshape.ActorStateTransition != nil {
		StageBranch(stage, actorstatetransitionshape.ActorStateTransition)
	}
	if actorstatetransitionshape.Start != nil {
		StageBranch(stage, actorstatetransitionshape.Start)
	}
	if actorstatetransitionshape.End != nil {
		StageBranch(stage, actorstatetransitionshape.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range actorstatetransitionshape.ControlPointShapes {
		StageBranch(stage, _controlpointshape)
	}

}

func (stage *Stage) StageBranchAnalysis(analysis *Analysis) {

	// check if instance is already staged
	if IsStaged(stage, analysis) {
		return
	}

	analysis.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _scenario := range analysis.Scenarios {
		StageBranch(stage, _scenario)
	}
	for _, _groupuse := range analysis.GroupUse {
		StageBranch(stage, _groupuse)
	}
	for _, _geoobjectuse := range analysis.GeoObjectUse {
		StageBranch(stage, _geoobjectuse)
	}
	for _, _mapobjectuse := range analysis.MapUse {
		StageBranch(stage, _mapobjectuse)
	}

}

func (stage *Stage) StageBranchControlPointShape(controlpointshape *ControlPointShape) {

	// check if instance is already staged
	if IsStaged(stage, controlpointshape) {
		return
	}

	controlpointshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchDiagram(diagram *Diagram) {

	// check if instance is already staged
	if IsStaged(stage, diagram) {
		return
	}

	diagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _evolutiondirectionshape := range diagram.EvolutionDirectionShapes {
		StageBranch(stage, _evolutiondirectionshape)
	}
	for _, _evolutiondirection := range diagram.EvolutionDirectionsWhoseNodeIsExpanded {
		StageBranch(stage, _evolutiondirection)
	}
	for _, _actorstateshape := range diagram.ActorStateShapes {
		StageBranch(stage, _actorstateshape)
	}
	for _, _actorstate := range diagram.ActorStatesWhoseNodeIsExpanded {
		StageBranch(stage, _actorstate)
	}
	for _, _parametershape := range diagram.ParameterShapes {
		StageBranch(stage, _parametershape)
	}
	for _, _parameter := range diagram.ParametersWhoseNodeIsExpanded {
		StageBranch(stage, _parameter)
	}
	for _, _parametersaggregateshape := range diagram.ScenarioParameterShapes {
		StageBranch(stage, _parametersaggregateshape)
	}
	for _, _parametersaggregate := range diagram.ParametersAggregatesWhoseNodeIsExpanded {
		StageBranch(stage, _parametersaggregate)
	}
	for _, _actorstatetransitionshape := range diagram.ActorStateTransitionShapes {
		StageBranch(stage, _actorstatetransitionshape)
	}
	for _, _actorstatetransition := range diagram.ActorStateTransitionsWhoseNodeIsExpanded {
		StageBranch(stage, _actorstatetransition)
	}

}

func (stage *Stage) StageBranchDocument(document *Document) {

	// check if instance is already staged
	if IsStaged(stage, document) {
		return
	}

	document.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _geoobjectuse := range document.GeoObjectUse {
		StageBranch(stage, _geoobjectuse)
	}

}

func (stage *Stage) StageBranchDocumentUse(documentuse *DocumentUse) {

	// check if instance is already staged
	if IsStaged(stage, documentuse) {
		return
	}

	documentuse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if documentuse.Document != nil {
		StageBranch(stage, documentuse.Document)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchEvolutionDirection(evolutiondirection *EvolutionDirection) {

	// check if instance is already staged
	if IsStaged(stage, evolutiondirection) {
		return
	}

	evolutiondirection.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchEvolutionDirectionShape(evolutiondirectionshape *EvolutionDirectionShape) {

	// check if instance is already staged
	if IsStaged(stage, evolutiondirectionshape) {
		return
	}

	evolutiondirectionshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if evolutiondirectionshape.EvolutionDirection != nil {
		StageBranch(stage, evolutiondirectionshape.EvolutionDirection)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchFoo(foo *Foo) {

	// check if instance is already staged
	if IsStaged(stage, foo) {
		return
	}

	foo.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGeoObject(geoobject *GeoObject) {

	// check if instance is already staged
	if IsStaged(stage, geoobject) {
		return
	}

	geoobject.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGeoObjectUse(geoobjectuse *GeoObjectUse) {

	// check if instance is already staged
	if IsStaged(stage, geoobjectuse) {
		return
	}

	geoobjectuse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if geoobjectuse.GeoObject != nil {
		StageBranch(stage, geoobjectuse.GeoObject)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGroup(group *Group) {

	// check if instance is already staged
	if IsStaged(stage, group) {
		return
	}

	group.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _useruse := range group.UserUse {
		StageBranch(stage, _useruse)
	}

}

func (stage *Stage) StageBranchGroupUse(groupuse *GroupUse) {

	// check if instance is already staged
	if IsStaged(stage, groupuse) {
		return
	}

	groupuse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if groupuse.Group != nil {
		StageBranch(stage, groupuse.Group)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchLibrary(library *Library) {

	// check if instance is already staged
	if IsStaged(stage, library) {
		return
	}

	library.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _analysis := range library.Analyses {
		StageBranch(stage, _analysis)
	}
	for _, _library := range library.SubLibraries {
		StageBranch(stage, _library)
	}
	for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
		StageBranch(stage, _library)
	}

}

func (stage *Stage) StageBranchMapObject(mapobject *MapObject) {

	// check if instance is already staged
	if IsStaged(stage, mapobject) {
		return
	}

	mapobject.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchMapObjectUse(mapobjectuse *MapObjectUse) {

	// check if instance is already staged
	if IsStaged(stage, mapobjectuse) {
		return
	}

	mapobjectuse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if mapobjectuse.Map != nil {
		StageBranch(stage, mapobjectuse.Map)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchParameter(parameter *Parameter) {

	// check if instance is already staged
	if IsStaged(stage, parameter) {
		return
	}

	parameter.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _groupuse := range parameter.GroupUse {
		StageBranch(stage, _groupuse)
	}
	for _, _documentuse := range parameter.DocumentUse {
		StageBranch(stage, _documentuse)
	}
	for _, _geoobjectuse := range parameter.GeoObjectUse {
		StageBranch(stage, _geoobjectuse)
	}

}

func (stage *Stage) StageBranchParameterCategory(parametercategory *ParameterCategory) {

	// check if instance is already staged
	if IsStaged(stage, parametercategory) {
		return
	}

	parametercategory.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parametershape := range parametercategory.ParameterUse {
		StageBranch(stage, _parametershape)
	}

}

func (stage *Stage) StageBranchParameterCategoryUse(parametercategoryuse *ParameterCategoryUse) {

	// check if instance is already staged
	if IsStaged(stage, parametercategoryuse) {
		return
	}

	parametercategoryuse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if parametercategoryuse.ParameterCategory != nil {
		StageBranch(stage, parametercategoryuse.ParameterCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchParameterShape(parametershape *ParameterShape) {

	// check if instance is already staged
	if IsStaged(stage, parametershape) {
		return
	}

	parametershape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if parametershape.Parameter != nil {
		StageBranch(stage, parametershape.Parameter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchParametersAggregate(parametersaggregate *ParametersAggregate) {

	// check if instance is already staged
	if IsStaged(stage, parametersaggregate) {
		return
	}

	parametersaggregate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parameter := range parametersaggregate.Parameters {
		StageBranch(stage, _parameter)
	}

}

func (stage *Stage) StageBranchParametersAggregateShape(parametersaggregateshape *ParametersAggregateShape) {

	// check if instance is already staged
	if IsStaged(stage, parametersaggregateshape) {
		return
	}

	parametersaggregateshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if parametersaggregateshape.ScenarioParameter != nil {
		StageBranch(stage, parametersaggregateshape.ScenarioParameter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPosition(position *Position) {

	// check if instance is already staged
	if IsStaged(stage, position) {
		return
	}

	position.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchRepository(repository *Repository) {

	// check if instance is already staged
	if IsStaged(stage, repository) {
		return
	}

	repository.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parametershape := range repository.ParameterUse {
		StageBranch(stage, _parametershape)
	}
	for _, _groupuse := range repository.GroupUse {
		StageBranch(stage, _groupuse)
	}

}

func (stage *Stage) StageBranchScenario(scenario *Scenario) {

	// check if instance is already staged
	if IsStaged(stage, scenario) {
		return
	}

	scenario.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _diagram := range scenario.Diagrams {
		StageBranch(stage, _diagram)
	}
	for _, _actorstate := range scenario.ActorStates {
		StageBranch(stage, _actorstate)
	}
	for _, _actorstatetransition := range scenario.ActorStateTransitions {
		StageBranch(stage, _actorstatetransition)
	}
	for _, _evolutiondirection := range scenario.EvolutionDirections {
		StageBranch(stage, _evolutiondirection)
	}
	for _, _parameter := range scenario.Parameters {
		StageBranch(stage, _parameter)
	}
	for _, _parametersaggregate := range scenario.ParametersAggretates {
		StageBranch(stage, _parametersaggregate)
	}

}

func (stage *Stage) StageBranchUser(user *User) {

	// check if instance is already staged
	if IsStaged(stage, user) {
		return
	}

	user.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchUserUse(useruse *UserUse) {

	// check if instance is already staged
	if IsStaged(stage, useruse) {
		return
	}

	useruse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if useruse.User != nil {
		StageBranch(stage, useruse.User)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchWorkspace(workspace *Workspace) {

	// check if instance is already staged
	if IsStaged(stage, workspace) {
		return
	}

	workspace.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if workspace.SelectedDiagram != nil {
		StageBranch(stage, workspace.SelectedDiagram)
	}
	if workspace.Default_EvolutionDirectionShape != nil {
		StageBranch(stage, workspace.Default_EvolutionDirectionShape)
	}
	if workspace.Default_ParameterShape != nil {
		StageBranch(stage, workspace.Default_ParameterShape)
	}
	if workspace.Default_ScenarioParameterShape != nil {
		StageBranch(stage, workspace.Default_ScenarioParameterShape)
	}
	if workspace.Default_ActorStateShape != nil {
		StageBranch(stage, workspace.Default_ActorStateShape)
	}
	if workspace.Default_ActorStateTransitionShape != nil {
		StageBranch(stage, workspace.Default_ActorStateTransitionShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *ActorState:
		toT := CopyBranchActorState(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ActorStateShape:
		toT := CopyBranchActorStateShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ActorStateTransition:
		toT := CopyBranchActorStateTransition(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ActorStateTransitionShape:
		toT := CopyBranchActorStateTransitionShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Analysis:
		toT := CopyBranchAnalysis(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ControlPointShape:
		toT := CopyBranchControlPointShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Diagram:
		toT := CopyBranchDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Document:
		toT := CopyBranchDocument(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DocumentUse:
		toT := CopyBranchDocumentUse(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EvolutionDirection:
		toT := CopyBranchEvolutionDirection(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EvolutionDirectionShape:
		toT := CopyBranchEvolutionDirectionShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Foo:
		toT := CopyBranchFoo(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GeoObject:
		toT := CopyBranchGeoObject(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GeoObjectUse:
		toT := CopyBranchGeoObjectUse(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Group:
		toT := CopyBranchGroup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GroupUse:
		toT := CopyBranchGroupUse(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := CopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MapObject:
		toT := CopyBranchMapObject(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MapObjectUse:
		toT := CopyBranchMapObjectUse(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Parameter:
		toT := CopyBranchParameter(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ParameterCategory:
		toT := CopyBranchParameterCategory(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ParameterCategoryUse:
		toT := CopyBranchParameterCategoryUse(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ParameterShape:
		toT := CopyBranchParameterShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ParametersAggregate:
		toT := CopyBranchParametersAggregate(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ParametersAggregateShape:
		toT := CopyBranchParametersAggregateShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Position:
		toT := CopyBranchPosition(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Repository:
		toT := CopyBranchRepository(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Scenario:
		toT := CopyBranchScenario(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *User:
		toT := CopyBranchUser(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *UserUse:
		toT := CopyBranchUserUse(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Workspace:
		toT := CopyBranchWorkspace(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchActorState(mapOrigCopy map[any]any, actorstateFrom *ActorState) (actorstateTo *ActorState) {

	// actorstateFrom has already been copied
	if _actorstateTo, ok := mapOrigCopy[actorstateFrom]; ok {
		actorstateTo = _actorstateTo.(*ActorState)
		return
	}

	actorstateTo = new(ActorState)
	mapOrigCopy[actorstateFrom] = actorstateTo
	actorstateFrom.CopyBasicFields(actorstateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchActorStateShape(mapOrigCopy map[any]any, actorstateshapeFrom *ActorStateShape) (actorstateshapeTo *ActorStateShape) {

	// actorstateshapeFrom has already been copied
	if _actorstateshapeTo, ok := mapOrigCopy[actorstateshapeFrom]; ok {
		actorstateshapeTo = _actorstateshapeTo.(*ActorStateShape)
		return
	}

	actorstateshapeTo = new(ActorStateShape)
	mapOrigCopy[actorstateshapeFrom] = actorstateshapeTo
	actorstateshapeFrom.CopyBasicFields(actorstateshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if actorstateshapeFrom.ActorState != nil {
		actorstateshapeTo.ActorState = actorstateshapeFrom.ActorState
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchActorStateTransition(mapOrigCopy map[any]any, actorstatetransitionFrom *ActorStateTransition) (actorstatetransitionTo *ActorStateTransition) {

	// actorstatetransitionFrom has already been copied
	if _actorstatetransitionTo, ok := mapOrigCopy[actorstatetransitionFrom]; ok {
		actorstatetransitionTo = _actorstatetransitionTo.(*ActorStateTransition)
		return
	}

	actorstatetransitionTo = new(ActorStateTransition)
	mapOrigCopy[actorstatetransitionFrom] = actorstatetransitionTo
	actorstatetransitionFrom.CopyBasicFields(actorstatetransitionTo)

	//insertion point for the staging of instances referenced by pointers
	if actorstatetransitionFrom.StartState != nil {
		actorstatetransitionTo.StartState = CopyBranchActorState(mapOrigCopy, actorstatetransitionFrom.StartState)
	}
	if actorstatetransitionFrom.EndState != nil {
		actorstatetransitionTo.EndState = CopyBranchActorState(mapOrigCopy, actorstatetransitionFrom.EndState)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parameter := range actorstatetransitionFrom.Justifications {
		actorstatetransitionTo.Justifications = append(actorstatetransitionTo.Justifications, CopyBranchParameter(mapOrigCopy, _parameter))
	}

	return
}

func CopyBranchActorStateTransitionShape(mapOrigCopy map[any]any, actorstatetransitionshapeFrom *ActorStateTransitionShape) (actorstatetransitionshapeTo *ActorStateTransitionShape) {

	// actorstatetransitionshapeFrom has already been copied
	if _actorstatetransitionshapeTo, ok := mapOrigCopy[actorstatetransitionshapeFrom]; ok {
		actorstatetransitionshapeTo = _actorstatetransitionshapeTo.(*ActorStateTransitionShape)
		return
	}

	actorstatetransitionshapeTo = new(ActorStateTransitionShape)
	mapOrigCopy[actorstatetransitionshapeFrom] = actorstatetransitionshapeTo
	actorstatetransitionshapeFrom.CopyBasicFields(actorstatetransitionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if actorstatetransitionshapeFrom.ActorStateTransition != nil {
		actorstatetransitionshapeTo.ActorStateTransition = actorstatetransitionshapeFrom.ActorStateTransition
	}
	if actorstatetransitionshapeFrom.Start != nil {
		actorstatetransitionshapeTo.Start = CopyBranchActorStateShape(mapOrigCopy, actorstatetransitionshapeFrom.Start)
	}
	if actorstatetransitionshapeFrom.End != nil {
		actorstatetransitionshapeTo.End = CopyBranchActorStateShape(mapOrigCopy, actorstatetransitionshapeFrom.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range actorstatetransitionshapeFrom.ControlPointShapes {
		actorstatetransitionshapeTo.ControlPointShapes = append(actorstatetransitionshapeTo.ControlPointShapes, CopyBranchControlPointShape(mapOrigCopy, _controlpointshape))
	}

	return
}

func CopyBranchAnalysis(mapOrigCopy map[any]any, analysisFrom *Analysis) (analysisTo *Analysis) {

	// analysisFrom has already been copied
	if _analysisTo, ok := mapOrigCopy[analysisFrom]; ok {
		analysisTo = _analysisTo.(*Analysis)
		return
	}

	analysisTo = new(Analysis)
	mapOrigCopy[analysisFrom] = analysisTo
	analysisFrom.CopyBasicFields(analysisTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _scenario := range analysisFrom.Scenarios {
		analysisTo.Scenarios = append(analysisTo.Scenarios, CopyBranchScenario(mapOrigCopy, _scenario))
	}
	for _, _groupuse := range analysisFrom.GroupUse {
		analysisTo.GroupUse = append(analysisTo.GroupUse, CopyBranchGroupUse(mapOrigCopy, _groupuse))
	}
	for _, _geoobjectuse := range analysisFrom.GeoObjectUse {
		analysisTo.GeoObjectUse = append(analysisTo.GeoObjectUse, CopyBranchGeoObjectUse(mapOrigCopy, _geoobjectuse))
	}
	for _, _mapobjectuse := range analysisFrom.MapUse {
		analysisTo.MapUse = append(analysisTo.MapUse, CopyBranchMapObjectUse(mapOrigCopy, _mapobjectuse))
	}

	return
}

func CopyBranchControlPointShape(mapOrigCopy map[any]any, controlpointshapeFrom *ControlPointShape) (controlpointshapeTo *ControlPointShape) {

	// controlpointshapeFrom has already been copied
	if _controlpointshapeTo, ok := mapOrigCopy[controlpointshapeFrom]; ok {
		controlpointshapeTo = _controlpointshapeTo.(*ControlPointShape)
		return
	}

	controlpointshapeTo = new(ControlPointShape)
	mapOrigCopy[controlpointshapeFrom] = controlpointshapeTo
	controlpointshapeFrom.CopyBasicFields(controlpointshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDiagram(mapOrigCopy map[any]any, diagramFrom *Diagram) (diagramTo *Diagram) {

	// diagramFrom has already been copied
	if _diagramTo, ok := mapOrigCopy[diagramFrom]; ok {
		diagramTo = _diagramTo.(*Diagram)
		return
	}

	diagramTo = new(Diagram)
	mapOrigCopy[diagramFrom] = diagramTo
	diagramFrom.CopyBasicFields(diagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _evolutiondirectionshape := range diagramFrom.EvolutionDirectionShapes {
		diagramTo.EvolutionDirectionShapes = append(diagramTo.EvolutionDirectionShapes, CopyBranchEvolutionDirectionShape(mapOrigCopy, _evolutiondirectionshape))
	}
	for _, _evolutiondirection := range diagramFrom.EvolutionDirectionsWhoseNodeIsExpanded {
		diagramTo.EvolutionDirectionsWhoseNodeIsExpanded = append(diagramTo.EvolutionDirectionsWhoseNodeIsExpanded, CopyBranchEvolutionDirection(mapOrigCopy, _evolutiondirection))
	}
	for _, _actorstateshape := range diagramFrom.ActorStateShapes {
		diagramTo.ActorStateShapes = append(diagramTo.ActorStateShapes, CopyBranchActorStateShape(mapOrigCopy, _actorstateshape))
	}
	for _, _actorstate := range diagramFrom.ActorStatesWhoseNodeIsExpanded {
		diagramTo.ActorStatesWhoseNodeIsExpanded = append(diagramTo.ActorStatesWhoseNodeIsExpanded, CopyBranchActorState(mapOrigCopy, _actorstate))
	}
	for _, _parametershape := range diagramFrom.ParameterShapes {
		diagramTo.ParameterShapes = append(diagramTo.ParameterShapes, CopyBranchParameterShape(mapOrigCopy, _parametershape))
	}
	for _, _parameter := range diagramFrom.ParametersWhoseNodeIsExpanded {
		diagramTo.ParametersWhoseNodeIsExpanded = append(diagramTo.ParametersWhoseNodeIsExpanded, CopyBranchParameter(mapOrigCopy, _parameter))
	}
	for _, _parametersaggregateshape := range diagramFrom.ScenarioParameterShapes {
		diagramTo.ScenarioParameterShapes = append(diagramTo.ScenarioParameterShapes, CopyBranchParametersAggregateShape(mapOrigCopy, _parametersaggregateshape))
	}
	for _, _parametersaggregate := range diagramFrom.ParametersAggregatesWhoseNodeIsExpanded {
		diagramTo.ParametersAggregatesWhoseNodeIsExpanded = append(diagramTo.ParametersAggregatesWhoseNodeIsExpanded, CopyBranchParametersAggregate(mapOrigCopy, _parametersaggregate))
	}
	for _, _actorstatetransitionshape := range diagramFrom.ActorStateTransitionShapes {
		diagramTo.ActorStateTransitionShapes = append(diagramTo.ActorStateTransitionShapes, CopyBranchActorStateTransitionShape(mapOrigCopy, _actorstatetransitionshape))
	}
	for _, _actorstatetransition := range diagramFrom.ActorStateTransitionsWhoseNodeIsExpanded {
		diagramTo.ActorStateTransitionsWhoseNodeIsExpanded = append(diagramTo.ActorStateTransitionsWhoseNodeIsExpanded, CopyBranchActorStateTransition(mapOrigCopy, _actorstatetransition))
	}

	return
}

func CopyBranchDocument(mapOrigCopy map[any]any, documentFrom *Document) (documentTo *Document) {

	// documentFrom has already been copied
	if _documentTo, ok := mapOrigCopy[documentFrom]; ok {
		documentTo = _documentTo.(*Document)
		return
	}

	documentTo = new(Document)
	mapOrigCopy[documentFrom] = documentTo
	documentFrom.CopyBasicFields(documentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _geoobjectuse := range documentFrom.GeoObjectUse {
		documentTo.GeoObjectUse = append(documentTo.GeoObjectUse, CopyBranchGeoObjectUse(mapOrigCopy, _geoobjectuse))
	}

	return
}

func CopyBranchDocumentUse(mapOrigCopy map[any]any, documentuseFrom *DocumentUse) (documentuseTo *DocumentUse) {

	// documentuseFrom has already been copied
	if _documentuseTo, ok := mapOrigCopy[documentuseFrom]; ok {
		documentuseTo = _documentuseTo.(*DocumentUse)
		return
	}

	documentuseTo = new(DocumentUse)
	mapOrigCopy[documentuseFrom] = documentuseTo
	documentuseFrom.CopyBasicFields(documentuseTo)

	//insertion point for the staging of instances referenced by pointers
	if documentuseFrom.Document != nil {
		documentuseTo.Document = CopyBranchDocument(mapOrigCopy, documentuseFrom.Document)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEvolutionDirection(mapOrigCopy map[any]any, evolutiondirectionFrom *EvolutionDirection) (evolutiondirectionTo *EvolutionDirection) {

	// evolutiondirectionFrom has already been copied
	if _evolutiondirectionTo, ok := mapOrigCopy[evolutiondirectionFrom]; ok {
		evolutiondirectionTo = _evolutiondirectionTo.(*EvolutionDirection)
		return
	}

	evolutiondirectionTo = new(EvolutionDirection)
	mapOrigCopy[evolutiondirectionFrom] = evolutiondirectionTo
	evolutiondirectionFrom.CopyBasicFields(evolutiondirectionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEvolutionDirectionShape(mapOrigCopy map[any]any, evolutiondirectionshapeFrom *EvolutionDirectionShape) (evolutiondirectionshapeTo *EvolutionDirectionShape) {

	// evolutiondirectionshapeFrom has already been copied
	if _evolutiondirectionshapeTo, ok := mapOrigCopy[evolutiondirectionshapeFrom]; ok {
		evolutiondirectionshapeTo = _evolutiondirectionshapeTo.(*EvolutionDirectionShape)
		return
	}

	evolutiondirectionshapeTo = new(EvolutionDirectionShape)
	mapOrigCopy[evolutiondirectionshapeFrom] = evolutiondirectionshapeTo
	evolutiondirectionshapeFrom.CopyBasicFields(evolutiondirectionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if evolutiondirectionshapeFrom.EvolutionDirection != nil {
		evolutiondirectionshapeTo.EvolutionDirection = evolutiondirectionshapeFrom.EvolutionDirection
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFoo(mapOrigCopy map[any]any, fooFrom *Foo) (fooTo *Foo) {

	// fooFrom has already been copied
	if _fooTo, ok := mapOrigCopy[fooFrom]; ok {
		fooTo = _fooTo.(*Foo)
		return
	}

	fooTo = new(Foo)
	mapOrigCopy[fooFrom] = fooTo
	fooFrom.CopyBasicFields(fooTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGeoObject(mapOrigCopy map[any]any, geoobjectFrom *GeoObject) (geoobjectTo *GeoObject) {

	// geoobjectFrom has already been copied
	if _geoobjectTo, ok := mapOrigCopy[geoobjectFrom]; ok {
		geoobjectTo = _geoobjectTo.(*GeoObject)
		return
	}

	geoobjectTo = new(GeoObject)
	mapOrigCopy[geoobjectFrom] = geoobjectTo
	geoobjectFrom.CopyBasicFields(geoobjectTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGeoObjectUse(mapOrigCopy map[any]any, geoobjectuseFrom *GeoObjectUse) (geoobjectuseTo *GeoObjectUse) {

	// geoobjectuseFrom has already been copied
	if _geoobjectuseTo, ok := mapOrigCopy[geoobjectuseFrom]; ok {
		geoobjectuseTo = _geoobjectuseTo.(*GeoObjectUse)
		return
	}

	geoobjectuseTo = new(GeoObjectUse)
	mapOrigCopy[geoobjectuseFrom] = geoobjectuseTo
	geoobjectuseFrom.CopyBasicFields(geoobjectuseTo)

	//insertion point for the staging of instances referenced by pointers
	if geoobjectuseFrom.GeoObject != nil {
		geoobjectuseTo.GeoObject = CopyBranchGeoObject(mapOrigCopy, geoobjectuseFrom.GeoObject)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGroup(mapOrigCopy map[any]any, groupFrom *Group) (groupTo *Group) {

	// groupFrom has already been copied
	if _groupTo, ok := mapOrigCopy[groupFrom]; ok {
		groupTo = _groupTo.(*Group)
		return
	}

	groupTo = new(Group)
	mapOrigCopy[groupFrom] = groupTo
	groupFrom.CopyBasicFields(groupTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _useruse := range groupFrom.UserUse {
		groupTo.UserUse = append(groupTo.UserUse, CopyBranchUserUse(mapOrigCopy, _useruse))
	}

	return
}

func CopyBranchGroupUse(mapOrigCopy map[any]any, groupuseFrom *GroupUse) (groupuseTo *GroupUse) {

	// groupuseFrom has already been copied
	if _groupuseTo, ok := mapOrigCopy[groupuseFrom]; ok {
		groupuseTo = _groupuseTo.(*GroupUse)
		return
	}

	groupuseTo = new(GroupUse)
	mapOrigCopy[groupuseFrom] = groupuseTo
	groupuseFrom.CopyBasicFields(groupuseTo)

	//insertion point for the staging of instances referenced by pointers
	if groupuseFrom.Group != nil {
		groupuseTo.Group = CopyBranchGroup(mapOrigCopy, groupuseFrom.Group)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchLibrary(mapOrigCopy map[any]any, libraryFrom *Library) (libraryTo *Library) {

	// libraryFrom has already been copied
	if _libraryTo, ok := mapOrigCopy[libraryFrom]; ok {
		libraryTo = _libraryTo.(*Library)
		return
	}

	libraryTo = new(Library)
	mapOrigCopy[libraryFrom] = libraryTo
	libraryFrom.CopyBasicFields(libraryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _analysis := range libraryFrom.Analyses {
		libraryTo.Analyses = append(libraryTo.Analyses, CopyBranchAnalysis(mapOrigCopy, _analysis))
	}
	for _, _library := range libraryFrom.SubLibraries {
		libraryTo.SubLibraries = append(libraryTo.SubLibraries, CopyBranchLibrary(mapOrigCopy, _library))
	}
	for _, _library := range libraryFrom.SubLibrariesWhoseNodeIsExpanded {
		libraryTo.SubLibrariesWhoseNodeIsExpanded = append(libraryTo.SubLibrariesWhoseNodeIsExpanded, CopyBranchLibrary(mapOrigCopy, _library))
	}

	return
}

func CopyBranchMapObject(mapOrigCopy map[any]any, mapobjectFrom *MapObject) (mapobjectTo *MapObject) {

	// mapobjectFrom has already been copied
	if _mapobjectTo, ok := mapOrigCopy[mapobjectFrom]; ok {
		mapobjectTo = _mapobjectTo.(*MapObject)
		return
	}

	mapobjectTo = new(MapObject)
	mapOrigCopy[mapobjectFrom] = mapobjectTo
	mapobjectFrom.CopyBasicFields(mapobjectTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMapObjectUse(mapOrigCopy map[any]any, mapobjectuseFrom *MapObjectUse) (mapobjectuseTo *MapObjectUse) {

	// mapobjectuseFrom has already been copied
	if _mapobjectuseTo, ok := mapOrigCopy[mapobjectuseFrom]; ok {
		mapobjectuseTo = _mapobjectuseTo.(*MapObjectUse)
		return
	}

	mapobjectuseTo = new(MapObjectUse)
	mapOrigCopy[mapobjectuseFrom] = mapobjectuseTo
	mapobjectuseFrom.CopyBasicFields(mapobjectuseTo)

	//insertion point for the staging of instances referenced by pointers
	if mapobjectuseFrom.Map != nil {
		mapobjectuseTo.Map = CopyBranchMapObject(mapOrigCopy, mapobjectuseFrom.Map)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchParameter(mapOrigCopy map[any]any, parameterFrom *Parameter) (parameterTo *Parameter) {

	// parameterFrom has already been copied
	if _parameterTo, ok := mapOrigCopy[parameterFrom]; ok {
		parameterTo = _parameterTo.(*Parameter)
		return
	}

	parameterTo = new(Parameter)
	mapOrigCopy[parameterFrom] = parameterTo
	parameterFrom.CopyBasicFields(parameterTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _groupuse := range parameterFrom.GroupUse {
		parameterTo.GroupUse = append(parameterTo.GroupUse, CopyBranchGroupUse(mapOrigCopy, _groupuse))
	}
	for _, _documentuse := range parameterFrom.DocumentUse {
		parameterTo.DocumentUse = append(parameterTo.DocumentUse, CopyBranchDocumentUse(mapOrigCopy, _documentuse))
	}
	for _, _geoobjectuse := range parameterFrom.GeoObjectUse {
		parameterTo.GeoObjectUse = append(parameterTo.GeoObjectUse, CopyBranchGeoObjectUse(mapOrigCopy, _geoobjectuse))
	}

	return
}

func CopyBranchParameterCategory(mapOrigCopy map[any]any, parametercategoryFrom *ParameterCategory) (parametercategoryTo *ParameterCategory) {

	// parametercategoryFrom has already been copied
	if _parametercategoryTo, ok := mapOrigCopy[parametercategoryFrom]; ok {
		parametercategoryTo = _parametercategoryTo.(*ParameterCategory)
		return
	}

	parametercategoryTo = new(ParameterCategory)
	mapOrigCopy[parametercategoryFrom] = parametercategoryTo
	parametercategoryFrom.CopyBasicFields(parametercategoryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parametershape := range parametercategoryFrom.ParameterUse {
		parametercategoryTo.ParameterUse = append(parametercategoryTo.ParameterUse, CopyBranchParameterShape(mapOrigCopy, _parametershape))
	}

	return
}

func CopyBranchParameterCategoryUse(mapOrigCopy map[any]any, parametercategoryuseFrom *ParameterCategoryUse) (parametercategoryuseTo *ParameterCategoryUse) {

	// parametercategoryuseFrom has already been copied
	if _parametercategoryuseTo, ok := mapOrigCopy[parametercategoryuseFrom]; ok {
		parametercategoryuseTo = _parametercategoryuseTo.(*ParameterCategoryUse)
		return
	}

	parametercategoryuseTo = new(ParameterCategoryUse)
	mapOrigCopy[parametercategoryuseFrom] = parametercategoryuseTo
	parametercategoryuseFrom.CopyBasicFields(parametercategoryuseTo)

	//insertion point for the staging of instances referenced by pointers
	if parametercategoryuseFrom.ParameterCategory != nil {
		parametercategoryuseTo.ParameterCategory = CopyBranchParameterCategory(mapOrigCopy, parametercategoryuseFrom.ParameterCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchParameterShape(mapOrigCopy map[any]any, parametershapeFrom *ParameterShape) (parametershapeTo *ParameterShape) {

	// parametershapeFrom has already been copied
	if _parametershapeTo, ok := mapOrigCopy[parametershapeFrom]; ok {
		parametershapeTo = _parametershapeTo.(*ParameterShape)
		return
	}

	parametershapeTo = new(ParameterShape)
	mapOrigCopy[parametershapeFrom] = parametershapeTo
	parametershapeFrom.CopyBasicFields(parametershapeTo)

	//insertion point for the staging of instances referenced by pointers
	if parametershapeFrom.Parameter != nil {
		parametershapeTo.Parameter = parametershapeFrom.Parameter
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchParametersAggregate(mapOrigCopy map[any]any, parametersaggregateFrom *ParametersAggregate) (parametersaggregateTo *ParametersAggregate) {

	// parametersaggregateFrom has already been copied
	if _parametersaggregateTo, ok := mapOrigCopy[parametersaggregateFrom]; ok {
		parametersaggregateTo = _parametersaggregateTo.(*ParametersAggregate)
		return
	}

	parametersaggregateTo = new(ParametersAggregate)
	mapOrigCopy[parametersaggregateFrom] = parametersaggregateTo
	parametersaggregateFrom.CopyBasicFields(parametersaggregateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parameter := range parametersaggregateFrom.Parameters {
		parametersaggregateTo.Parameters = append(parametersaggregateTo.Parameters, CopyBranchParameter(mapOrigCopy, _parameter))
	}

	return
}

func CopyBranchParametersAggregateShape(mapOrigCopy map[any]any, parametersaggregateshapeFrom *ParametersAggregateShape) (parametersaggregateshapeTo *ParametersAggregateShape) {

	// parametersaggregateshapeFrom has already been copied
	if _parametersaggregateshapeTo, ok := mapOrigCopy[parametersaggregateshapeFrom]; ok {
		parametersaggregateshapeTo = _parametersaggregateshapeTo.(*ParametersAggregateShape)
		return
	}

	parametersaggregateshapeTo = new(ParametersAggregateShape)
	mapOrigCopy[parametersaggregateshapeFrom] = parametersaggregateshapeTo
	parametersaggregateshapeFrom.CopyBasicFields(parametersaggregateshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if parametersaggregateshapeFrom.ScenarioParameter != nil {
		parametersaggregateshapeTo.ScenarioParameter = parametersaggregateshapeFrom.ScenarioParameter
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPosition(mapOrigCopy map[any]any, positionFrom *Position) (positionTo *Position) {

	// positionFrom has already been copied
	if _positionTo, ok := mapOrigCopy[positionFrom]; ok {
		positionTo = _positionTo.(*Position)
		return
	}

	positionTo = new(Position)
	mapOrigCopy[positionFrom] = positionTo
	positionFrom.CopyBasicFields(positionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRepository(mapOrigCopy map[any]any, repositoryFrom *Repository) (repositoryTo *Repository) {

	// repositoryFrom has already been copied
	if _repositoryTo, ok := mapOrigCopy[repositoryFrom]; ok {
		repositoryTo = _repositoryTo.(*Repository)
		return
	}

	repositoryTo = new(Repository)
	mapOrigCopy[repositoryFrom] = repositoryTo
	repositoryFrom.CopyBasicFields(repositoryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parametershape := range repositoryFrom.ParameterUse {
		repositoryTo.ParameterUse = append(repositoryTo.ParameterUse, CopyBranchParameterShape(mapOrigCopy, _parametershape))
	}
	for _, _groupuse := range repositoryFrom.GroupUse {
		repositoryTo.GroupUse = append(repositoryTo.GroupUse, CopyBranchGroupUse(mapOrigCopy, _groupuse))
	}

	return
}

func CopyBranchScenario(mapOrigCopy map[any]any, scenarioFrom *Scenario) (scenarioTo *Scenario) {

	// scenarioFrom has already been copied
	if _scenarioTo, ok := mapOrigCopy[scenarioFrom]; ok {
		scenarioTo = _scenarioTo.(*Scenario)
		return
	}

	scenarioTo = new(Scenario)
	mapOrigCopy[scenarioFrom] = scenarioTo
	scenarioFrom.CopyBasicFields(scenarioTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _diagram := range scenarioFrom.Diagrams {
		scenarioTo.Diagrams = append(scenarioTo.Diagrams, CopyBranchDiagram(mapOrigCopy, _diagram))
	}
	for _, _actorstate := range scenarioFrom.ActorStates {
		scenarioTo.ActorStates = append(scenarioTo.ActorStates, CopyBranchActorState(mapOrigCopy, _actorstate))
	}
	for _, _actorstatetransition := range scenarioFrom.ActorStateTransitions {
		scenarioTo.ActorStateTransitions = append(scenarioTo.ActorStateTransitions, CopyBranchActorStateTransition(mapOrigCopy, _actorstatetransition))
	}
	for _, _evolutiondirection := range scenarioFrom.EvolutionDirections {
		scenarioTo.EvolutionDirections = append(scenarioTo.EvolutionDirections, CopyBranchEvolutionDirection(mapOrigCopy, _evolutiondirection))
	}
	for _, _parameter := range scenarioFrom.Parameters {
		scenarioTo.Parameters = append(scenarioTo.Parameters, CopyBranchParameter(mapOrigCopy, _parameter))
	}
	for _, _parametersaggregate := range scenarioFrom.ParametersAggretates {
		scenarioTo.ParametersAggretates = append(scenarioTo.ParametersAggretates, CopyBranchParametersAggregate(mapOrigCopy, _parametersaggregate))
	}

	return
}

func CopyBranchUser(mapOrigCopy map[any]any, userFrom *User) (userTo *User) {

	// userFrom has already been copied
	if _userTo, ok := mapOrigCopy[userFrom]; ok {
		userTo = _userTo.(*User)
		return
	}

	userTo = new(User)
	mapOrigCopy[userFrom] = userTo
	userFrom.CopyBasicFields(userTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchUserUse(mapOrigCopy map[any]any, useruseFrom *UserUse) (useruseTo *UserUse) {

	// useruseFrom has already been copied
	if _useruseTo, ok := mapOrigCopy[useruseFrom]; ok {
		useruseTo = _useruseTo.(*UserUse)
		return
	}

	useruseTo = new(UserUse)
	mapOrigCopy[useruseFrom] = useruseTo
	useruseFrom.CopyBasicFields(useruseTo)

	//insertion point for the staging of instances referenced by pointers
	if useruseFrom.User != nil {
		useruseTo.User = CopyBranchUser(mapOrigCopy, useruseFrom.User)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchWorkspace(mapOrigCopy map[any]any, workspaceFrom *Workspace) (workspaceTo *Workspace) {

	// workspaceFrom has already been copied
	if _workspaceTo, ok := mapOrigCopy[workspaceFrom]; ok {
		workspaceTo = _workspaceTo.(*Workspace)
		return
	}

	workspaceTo = new(Workspace)
	mapOrigCopy[workspaceFrom] = workspaceTo
	workspaceFrom.CopyBasicFields(workspaceTo)

	//insertion point for the staging of instances referenced by pointers
	if workspaceFrom.SelectedDiagram != nil {
		workspaceTo.SelectedDiagram = CopyBranchDiagram(mapOrigCopy, workspaceFrom.SelectedDiagram)
	}
	if workspaceFrom.Default_EvolutionDirectionShape != nil {
		workspaceTo.Default_EvolutionDirectionShape = CopyBranchEvolutionDirectionShape(mapOrigCopy, workspaceFrom.Default_EvolutionDirectionShape)
	}
	if workspaceFrom.Default_ParameterShape != nil {
		workspaceTo.Default_ParameterShape = CopyBranchParameterShape(mapOrigCopy, workspaceFrom.Default_ParameterShape)
	}
	if workspaceFrom.Default_ScenarioParameterShape != nil {
		workspaceTo.Default_ScenarioParameterShape = CopyBranchParametersAggregateShape(mapOrigCopy, workspaceFrom.Default_ScenarioParameterShape)
	}
	if workspaceFrom.Default_ActorStateShape != nil {
		workspaceTo.Default_ActorStateShape = CopyBranchActorStateShape(mapOrigCopy, workspaceFrom.Default_ActorStateShape)
	}
	if workspaceFrom.Default_ActorStateTransitionShape != nil {
		workspaceTo.Default_ActorStateTransitionShape = CopyBranchActorStateTransitionShape(mapOrigCopy, workspaceFrom.Default_ActorStateTransitionShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *ActorState:
		stage.UnstageBranchActorState(target)

	case *ActorStateShape:
		stage.UnstageBranchActorStateShape(target)

	case *ActorStateTransition:
		stage.UnstageBranchActorStateTransition(target)

	case *ActorStateTransitionShape:
		stage.UnstageBranchActorStateTransitionShape(target)

	case *Analysis:
		stage.UnstageBranchAnalysis(target)

	case *ControlPointShape:
		stage.UnstageBranchControlPointShape(target)

	case *Diagram:
		stage.UnstageBranchDiagram(target)

	case *Document:
		stage.UnstageBranchDocument(target)

	case *DocumentUse:
		stage.UnstageBranchDocumentUse(target)

	case *EvolutionDirection:
		stage.UnstageBranchEvolutionDirection(target)

	case *EvolutionDirectionShape:
		stage.UnstageBranchEvolutionDirectionShape(target)

	case *Foo:
		stage.UnstageBranchFoo(target)

	case *GeoObject:
		stage.UnstageBranchGeoObject(target)

	case *GeoObjectUse:
		stage.UnstageBranchGeoObjectUse(target)

	case *Group:
		stage.UnstageBranchGroup(target)

	case *GroupUse:
		stage.UnstageBranchGroupUse(target)

	case *Library:
		stage.UnstageBranchLibrary(target)

	case *MapObject:
		stage.UnstageBranchMapObject(target)

	case *MapObjectUse:
		stage.UnstageBranchMapObjectUse(target)

	case *Parameter:
		stage.UnstageBranchParameter(target)

	case *ParameterCategory:
		stage.UnstageBranchParameterCategory(target)

	case *ParameterCategoryUse:
		stage.UnstageBranchParameterCategoryUse(target)

	case *ParameterShape:
		stage.UnstageBranchParameterShape(target)

	case *ParametersAggregate:
		stage.UnstageBranchParametersAggregate(target)

	case *ParametersAggregateShape:
		stage.UnstageBranchParametersAggregateShape(target)

	case *Position:
		stage.UnstageBranchPosition(target)

	case *Repository:
		stage.UnstageBranchRepository(target)

	case *Scenario:
		stage.UnstageBranchScenario(target)

	case *User:
		stage.UnstageBranchUser(target)

	case *UserUse:
		stage.UnstageBranchUserUse(target)

	case *Workspace:
		stage.UnstageBranchWorkspace(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchActorState(actorstate *ActorState) {

	// check if instance is already staged
	if !IsStaged(stage, actorstate) {
		return
	}

	actorstate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchActorStateShape(actorstateshape *ActorStateShape) {

	// check if instance is already staged
	if !IsStaged(stage, actorstateshape) {
		return
	}

	actorstateshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if actorstateshape.ActorState != nil {
		UnstageBranch(stage, actorstateshape.ActorState)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchActorStateTransition(actorstatetransition *ActorStateTransition) {

	// check if instance is already staged
	if !IsStaged(stage, actorstatetransition) {
		return
	}

	actorstatetransition.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if actorstatetransition.StartState != nil {
		UnstageBranch(stage, actorstatetransition.StartState)
	}
	if actorstatetransition.EndState != nil {
		UnstageBranch(stage, actorstatetransition.EndState)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parameter := range actorstatetransition.Justifications {
		UnstageBranch(stage, _parameter)
	}

}

func (stage *Stage) UnstageBranchActorStateTransitionShape(actorstatetransitionshape *ActorStateTransitionShape) {

	// check if instance is already staged
	if !IsStaged(stage, actorstatetransitionshape) {
		return
	}

	actorstatetransitionshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if actorstatetransitionshape.ActorStateTransition != nil {
		UnstageBranch(stage, actorstatetransitionshape.ActorStateTransition)
	}
	if actorstatetransitionshape.Start != nil {
		UnstageBranch(stage, actorstatetransitionshape.Start)
	}
	if actorstatetransitionshape.End != nil {
		UnstageBranch(stage, actorstatetransitionshape.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range actorstatetransitionshape.ControlPointShapes {
		UnstageBranch(stage, _controlpointshape)
	}

}

func (stage *Stage) UnstageBranchAnalysis(analysis *Analysis) {

	// check if instance is already staged
	if !IsStaged(stage, analysis) {
		return
	}

	analysis.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _scenario := range analysis.Scenarios {
		UnstageBranch(stage, _scenario)
	}
	for _, _groupuse := range analysis.GroupUse {
		UnstageBranch(stage, _groupuse)
	}
	for _, _geoobjectuse := range analysis.GeoObjectUse {
		UnstageBranch(stage, _geoobjectuse)
	}
	for _, _mapobjectuse := range analysis.MapUse {
		UnstageBranch(stage, _mapobjectuse)
	}

}

func (stage *Stage) UnstageBranchControlPointShape(controlpointshape *ControlPointShape) {

	// check if instance is already staged
	if !IsStaged(stage, controlpointshape) {
		return
	}

	controlpointshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchDiagram(diagram *Diagram) {

	// check if instance is already staged
	if !IsStaged(stage, diagram) {
		return
	}

	diagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _evolutiondirectionshape := range diagram.EvolutionDirectionShapes {
		UnstageBranch(stage, _evolutiondirectionshape)
	}
	for _, _evolutiondirection := range diagram.EvolutionDirectionsWhoseNodeIsExpanded {
		UnstageBranch(stage, _evolutiondirection)
	}
	for _, _actorstateshape := range diagram.ActorStateShapes {
		UnstageBranch(stage, _actorstateshape)
	}
	for _, _actorstate := range diagram.ActorStatesWhoseNodeIsExpanded {
		UnstageBranch(stage, _actorstate)
	}
	for _, _parametershape := range diagram.ParameterShapes {
		UnstageBranch(stage, _parametershape)
	}
	for _, _parameter := range diagram.ParametersWhoseNodeIsExpanded {
		UnstageBranch(stage, _parameter)
	}
	for _, _parametersaggregateshape := range diagram.ScenarioParameterShapes {
		UnstageBranch(stage, _parametersaggregateshape)
	}
	for _, _parametersaggregate := range diagram.ParametersAggregatesWhoseNodeIsExpanded {
		UnstageBranch(stage, _parametersaggregate)
	}
	for _, _actorstatetransitionshape := range diagram.ActorStateTransitionShapes {
		UnstageBranch(stage, _actorstatetransitionshape)
	}
	for _, _actorstatetransition := range diagram.ActorStateTransitionsWhoseNodeIsExpanded {
		UnstageBranch(stage, _actorstatetransition)
	}

}

func (stage *Stage) UnstageBranchDocument(document *Document) {

	// check if instance is already staged
	if !IsStaged(stage, document) {
		return
	}

	document.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _geoobjectuse := range document.GeoObjectUse {
		UnstageBranch(stage, _geoobjectuse)
	}

}

func (stage *Stage) UnstageBranchDocumentUse(documentuse *DocumentUse) {

	// check if instance is already staged
	if !IsStaged(stage, documentuse) {
		return
	}

	documentuse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if documentuse.Document != nil {
		UnstageBranch(stage, documentuse.Document)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchEvolutionDirection(evolutiondirection *EvolutionDirection) {

	// check if instance is already staged
	if !IsStaged(stage, evolutiondirection) {
		return
	}

	evolutiondirection.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchEvolutionDirectionShape(evolutiondirectionshape *EvolutionDirectionShape) {

	// check if instance is already staged
	if !IsStaged(stage, evolutiondirectionshape) {
		return
	}

	evolutiondirectionshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if evolutiondirectionshape.EvolutionDirection != nil {
		UnstageBranch(stage, evolutiondirectionshape.EvolutionDirection)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchFoo(foo *Foo) {

	// check if instance is already staged
	if !IsStaged(stage, foo) {
		return
	}

	foo.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGeoObject(geoobject *GeoObject) {

	// check if instance is already staged
	if !IsStaged(stage, geoobject) {
		return
	}

	geoobject.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGeoObjectUse(geoobjectuse *GeoObjectUse) {

	// check if instance is already staged
	if !IsStaged(stage, geoobjectuse) {
		return
	}

	geoobjectuse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if geoobjectuse.GeoObject != nil {
		UnstageBranch(stage, geoobjectuse.GeoObject)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGroup(group *Group) {

	// check if instance is already staged
	if !IsStaged(stage, group) {
		return
	}

	group.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _useruse := range group.UserUse {
		UnstageBranch(stage, _useruse)
	}

}

func (stage *Stage) UnstageBranchGroupUse(groupuse *GroupUse) {

	// check if instance is already staged
	if !IsStaged(stage, groupuse) {
		return
	}

	groupuse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if groupuse.Group != nil {
		UnstageBranch(stage, groupuse.Group)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchLibrary(library *Library) {

	// check if instance is already staged
	if !IsStaged(stage, library) {
		return
	}

	library.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _analysis := range library.Analyses {
		UnstageBranch(stage, _analysis)
	}
	for _, _library := range library.SubLibraries {
		UnstageBranch(stage, _library)
	}
	for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
		UnstageBranch(stage, _library)
	}

}

func (stage *Stage) UnstageBranchMapObject(mapobject *MapObject) {

	// check if instance is already staged
	if !IsStaged(stage, mapobject) {
		return
	}

	mapobject.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchMapObjectUse(mapobjectuse *MapObjectUse) {

	// check if instance is already staged
	if !IsStaged(stage, mapobjectuse) {
		return
	}

	mapobjectuse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if mapobjectuse.Map != nil {
		UnstageBranch(stage, mapobjectuse.Map)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchParameter(parameter *Parameter) {

	// check if instance is already staged
	if !IsStaged(stage, parameter) {
		return
	}

	parameter.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _groupuse := range parameter.GroupUse {
		UnstageBranch(stage, _groupuse)
	}
	for _, _documentuse := range parameter.DocumentUse {
		UnstageBranch(stage, _documentuse)
	}
	for _, _geoobjectuse := range parameter.GeoObjectUse {
		UnstageBranch(stage, _geoobjectuse)
	}

}

func (stage *Stage) UnstageBranchParameterCategory(parametercategory *ParameterCategory) {

	// check if instance is already staged
	if !IsStaged(stage, parametercategory) {
		return
	}

	parametercategory.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parametershape := range parametercategory.ParameterUse {
		UnstageBranch(stage, _parametershape)
	}

}

func (stage *Stage) UnstageBranchParameterCategoryUse(parametercategoryuse *ParameterCategoryUse) {

	// check if instance is already staged
	if !IsStaged(stage, parametercategoryuse) {
		return
	}

	parametercategoryuse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if parametercategoryuse.ParameterCategory != nil {
		UnstageBranch(stage, parametercategoryuse.ParameterCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchParameterShape(parametershape *ParameterShape) {

	// check if instance is already staged
	if !IsStaged(stage, parametershape) {
		return
	}

	parametershape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if parametershape.Parameter != nil {
		UnstageBranch(stage, parametershape.Parameter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchParametersAggregate(parametersaggregate *ParametersAggregate) {

	// check if instance is already staged
	if !IsStaged(stage, parametersaggregate) {
		return
	}

	parametersaggregate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parameter := range parametersaggregate.Parameters {
		UnstageBranch(stage, _parameter)
	}

}

func (stage *Stage) UnstageBranchParametersAggregateShape(parametersaggregateshape *ParametersAggregateShape) {

	// check if instance is already staged
	if !IsStaged(stage, parametersaggregateshape) {
		return
	}

	parametersaggregateshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if parametersaggregateshape.ScenarioParameter != nil {
		UnstageBranch(stage, parametersaggregateshape.ScenarioParameter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPosition(position *Position) {

	// check if instance is already staged
	if !IsStaged(stage, position) {
		return
	}

	position.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchRepository(repository *Repository) {

	// check if instance is already staged
	if !IsStaged(stage, repository) {
		return
	}

	repository.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parametershape := range repository.ParameterUse {
		UnstageBranch(stage, _parametershape)
	}
	for _, _groupuse := range repository.GroupUse {
		UnstageBranch(stage, _groupuse)
	}

}

func (stage *Stage) UnstageBranchScenario(scenario *Scenario) {

	// check if instance is already staged
	if !IsStaged(stage, scenario) {
		return
	}

	scenario.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _diagram := range scenario.Diagrams {
		UnstageBranch(stage, _diagram)
	}
	for _, _actorstate := range scenario.ActorStates {
		UnstageBranch(stage, _actorstate)
	}
	for _, _actorstatetransition := range scenario.ActorStateTransitions {
		UnstageBranch(stage, _actorstatetransition)
	}
	for _, _evolutiondirection := range scenario.EvolutionDirections {
		UnstageBranch(stage, _evolutiondirection)
	}
	for _, _parameter := range scenario.Parameters {
		UnstageBranch(stage, _parameter)
	}
	for _, _parametersaggregate := range scenario.ParametersAggretates {
		UnstageBranch(stage, _parametersaggregate)
	}

}

func (stage *Stage) UnstageBranchUser(user *User) {

	// check if instance is already staged
	if !IsStaged(stage, user) {
		return
	}

	user.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchUserUse(useruse *UserUse) {

	// check if instance is already staged
	if !IsStaged(stage, useruse) {
		return
	}

	useruse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if useruse.User != nil {
		UnstageBranch(stage, useruse.User)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchWorkspace(workspace *Workspace) {

	// check if instance is already staged
	if !IsStaged(stage, workspace) {
		return
	}

	workspace.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if workspace.SelectedDiagram != nil {
		UnstageBranch(stage, workspace.SelectedDiagram)
	}
	if workspace.Default_EvolutionDirectionShape != nil {
		UnstageBranch(stage, workspace.Default_EvolutionDirectionShape)
	}
	if workspace.Default_ParameterShape != nil {
		UnstageBranch(stage, workspace.Default_ParameterShape)
	}
	if workspace.Default_ScenarioParameterShape != nil {
		UnstageBranch(stage, workspace.Default_ScenarioParameterShape)
	}
	if workspace.Default_ActorStateShape != nil {
		UnstageBranch(stage, workspace.Default_ActorStateShape)
	}
	if workspace.Default_ActorStateTransitionShape != nil {
		UnstageBranch(stage, workspace.Default_ActorStateTransitionShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *ActorState) GongReconstructPointersFromReferences(stage *Stage, instance *ActorState) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ActorStateShape) GongReconstructPointersFromReferences(stage *Stage, instance *ActorStateShape) {
	// insertion point for pointers field
	if instance.ActorState != nil {
		reference.ActorState = stage.ActorStates_reference[instance.ActorState]
	}
	// insertion point for slice of pointers field
}

func (reference *ActorStateTransition) GongReconstructPointersFromReferences(stage *Stage, instance *ActorStateTransition) {
	// insertion point for pointers field
	if instance.StartState != nil {
		reference.StartState = stage.ActorStates_reference[instance.StartState]
	}
	if instance.EndState != nil {
		reference.EndState = stage.ActorStates_reference[instance.EndState]
	}
	// insertion point for slice of pointers field
	reference.Justifications = reference.Justifications[:0]
	for _, _b := range instance.Justifications {
		reference.Justifications = append(reference.Justifications, stage.Parameters_reference[_b])
	}
}

func (reference *ActorStateTransitionShape) GongReconstructPointersFromReferences(stage *Stage, instance *ActorStateTransitionShape) {
	// insertion point for pointers field
	if instance.ActorStateTransition != nil {
		reference.ActorStateTransition = stage.ActorStateTransitions_reference[instance.ActorStateTransition]
	}
	if instance.Start != nil {
		reference.Start = stage.ActorStateShapes_reference[instance.Start]
	}
	if instance.End != nil {
		reference.End = stage.ActorStateShapes_reference[instance.End]
	}
	// insertion point for slice of pointers field
	reference.ControlPointShapes = reference.ControlPointShapes[:0]
	for _, _b := range instance.ControlPointShapes {
		reference.ControlPointShapes = append(reference.ControlPointShapes, stage.ControlPointShapes_reference[_b])
	}
}

func (reference *Analysis) GongReconstructPointersFromReferences(stage *Stage, instance *Analysis) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Scenarios = reference.Scenarios[:0]
	for _, _b := range instance.Scenarios {
		reference.Scenarios = append(reference.Scenarios, stage.Scenarios_reference[_b])
	}
	reference.GroupUse = reference.GroupUse[:0]
	for _, _b := range instance.GroupUse {
		reference.GroupUse = append(reference.GroupUse, stage.GroupUses_reference[_b])
	}
	reference.GeoObjectUse = reference.GeoObjectUse[:0]
	for _, _b := range instance.GeoObjectUse {
		reference.GeoObjectUse = append(reference.GeoObjectUse, stage.GeoObjectUses_reference[_b])
	}
	reference.MapUse = reference.MapUse[:0]
	for _, _b := range instance.MapUse {
		reference.MapUse = append(reference.MapUse, stage.MapObjectUses_reference[_b])
	}
}

func (reference *ControlPointShape) GongReconstructPointersFromReferences(stage *Stage, instance *ControlPointShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Diagram) GongReconstructPointersFromReferences(stage *Stage, instance *Diagram) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.EvolutionDirectionShapes = reference.EvolutionDirectionShapes[:0]
	for _, _b := range instance.EvolutionDirectionShapes {
		reference.EvolutionDirectionShapes = append(reference.EvolutionDirectionShapes, stage.EvolutionDirectionShapes_reference[_b])
	}
	reference.EvolutionDirectionsWhoseNodeIsExpanded = reference.EvolutionDirectionsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.EvolutionDirectionsWhoseNodeIsExpanded {
		reference.EvolutionDirectionsWhoseNodeIsExpanded = append(reference.EvolutionDirectionsWhoseNodeIsExpanded, stage.EvolutionDirections_reference[_b])
	}
	reference.ActorStateShapes = reference.ActorStateShapes[:0]
	for _, _b := range instance.ActorStateShapes {
		reference.ActorStateShapes = append(reference.ActorStateShapes, stage.ActorStateShapes_reference[_b])
	}
	reference.ActorStatesWhoseNodeIsExpanded = reference.ActorStatesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ActorStatesWhoseNodeIsExpanded {
		reference.ActorStatesWhoseNodeIsExpanded = append(reference.ActorStatesWhoseNodeIsExpanded, stage.ActorStates_reference[_b])
	}
	reference.ParameterShapes = reference.ParameterShapes[:0]
	for _, _b := range instance.ParameterShapes {
		reference.ParameterShapes = append(reference.ParameterShapes, stage.ParameterShapes_reference[_b])
	}
	reference.ParametersWhoseNodeIsExpanded = reference.ParametersWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ParametersWhoseNodeIsExpanded {
		reference.ParametersWhoseNodeIsExpanded = append(reference.ParametersWhoseNodeIsExpanded, stage.Parameters_reference[_b])
	}
	reference.ScenarioParameterShapes = reference.ScenarioParameterShapes[:0]
	for _, _b := range instance.ScenarioParameterShapes {
		reference.ScenarioParameterShapes = append(reference.ScenarioParameterShapes, stage.ParametersAggregateShapes_reference[_b])
	}
	reference.ParametersAggregatesWhoseNodeIsExpanded = reference.ParametersAggregatesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ParametersAggregatesWhoseNodeIsExpanded {
		reference.ParametersAggregatesWhoseNodeIsExpanded = append(reference.ParametersAggregatesWhoseNodeIsExpanded, stage.ParametersAggregates_reference[_b])
	}
	reference.ActorStateTransitionShapes = reference.ActorStateTransitionShapes[:0]
	for _, _b := range instance.ActorStateTransitionShapes {
		reference.ActorStateTransitionShapes = append(reference.ActorStateTransitionShapes, stage.ActorStateTransitionShapes_reference[_b])
	}
	reference.ActorStateTransitionsWhoseNodeIsExpanded = reference.ActorStateTransitionsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ActorStateTransitionsWhoseNodeIsExpanded {
		reference.ActorStateTransitionsWhoseNodeIsExpanded = append(reference.ActorStateTransitionsWhoseNodeIsExpanded, stage.ActorStateTransitions_reference[_b])
	}
}

func (reference *Document) GongReconstructPointersFromReferences(stage *Stage, instance *Document) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.GeoObjectUse = reference.GeoObjectUse[:0]
	for _, _b := range instance.GeoObjectUse {
		reference.GeoObjectUse = append(reference.GeoObjectUse, stage.GeoObjectUses_reference[_b])
	}
}

func (reference *DocumentUse) GongReconstructPointersFromReferences(stage *Stage, instance *DocumentUse) {
	// insertion point for pointers field
	if instance.Document != nil {
		reference.Document = stage.Documents_reference[instance.Document]
	}
	// insertion point for slice of pointers field
}

func (reference *EvolutionDirection) GongReconstructPointersFromReferences(stage *Stage, instance *EvolutionDirection) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EvolutionDirectionShape) GongReconstructPointersFromReferences(stage *Stage, instance *EvolutionDirectionShape) {
	// insertion point for pointers field
	if instance.EvolutionDirection != nil {
		reference.EvolutionDirection = stage.EvolutionDirections_reference[instance.EvolutionDirection]
	}
	// insertion point for slice of pointers field
}

func (reference *Foo) GongReconstructPointersFromReferences(stage *Stage, instance *Foo) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GeoObject) GongReconstructPointersFromReferences(stage *Stage, instance *GeoObject) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GeoObjectUse) GongReconstructPointersFromReferences(stage *Stage, instance *GeoObjectUse) {
	// insertion point for pointers field
	if instance.GeoObject != nil {
		reference.GeoObject = stage.GeoObjects_reference[instance.GeoObject]
	}
	// insertion point for slice of pointers field
}

func (reference *Group) GongReconstructPointersFromReferences(stage *Stage, instance *Group) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.UserUse = reference.UserUse[:0]
	for _, _b := range instance.UserUse {
		reference.UserUse = append(reference.UserUse, stage.UserUses_reference[_b])
	}
}

func (reference *GroupUse) GongReconstructPointersFromReferences(stage *Stage, instance *GroupUse) {
	// insertion point for pointers field
	if instance.Group != nil {
		reference.Group = stage.Groups_reference[instance.Group]
	}
	// insertion point for slice of pointers field
}

func (reference *Library) GongReconstructPointersFromReferences(stage *Stage, instance *Library) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Analyses = reference.Analyses[:0]
	for _, _b := range instance.Analyses {
		reference.Analyses = append(reference.Analyses, stage.Analysiss_reference[_b])
	}
	reference.SubLibraries = reference.SubLibraries[:0]
	for _, _b := range instance.SubLibraries {
		reference.SubLibraries = append(reference.SubLibraries, stage.Librarys_reference[_b])
	}
	reference.SubLibrariesWhoseNodeIsExpanded = reference.SubLibrariesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.SubLibrariesWhoseNodeIsExpanded {
		reference.SubLibrariesWhoseNodeIsExpanded = append(reference.SubLibrariesWhoseNodeIsExpanded, stage.Librarys_reference[_b])
	}
}

func (reference *MapObject) GongReconstructPointersFromReferences(stage *Stage, instance *MapObject) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MapObjectUse) GongReconstructPointersFromReferences(stage *Stage, instance *MapObjectUse) {
	// insertion point for pointers field
	if instance.Map != nil {
		reference.Map = stage.MapObjects_reference[instance.Map]
	}
	// insertion point for slice of pointers field
}

func (reference *Parameter) GongReconstructPointersFromReferences(stage *Stage, instance *Parameter) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.GroupUse = reference.GroupUse[:0]
	for _, _b := range instance.GroupUse {
		reference.GroupUse = append(reference.GroupUse, stage.GroupUses_reference[_b])
	}
	reference.DocumentUse = reference.DocumentUse[:0]
	for _, _b := range instance.DocumentUse {
		reference.DocumentUse = append(reference.DocumentUse, stage.DocumentUses_reference[_b])
	}
	reference.GeoObjectUse = reference.GeoObjectUse[:0]
	for _, _b := range instance.GeoObjectUse {
		reference.GeoObjectUse = append(reference.GeoObjectUse, stage.GeoObjectUses_reference[_b])
	}
}

func (reference *ParameterCategory) GongReconstructPointersFromReferences(stage *Stage, instance *ParameterCategory) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.ParameterUse = reference.ParameterUse[:0]
	for _, _b := range instance.ParameterUse {
		reference.ParameterUse = append(reference.ParameterUse, stage.ParameterShapes_reference[_b])
	}
}

func (reference *ParameterCategoryUse) GongReconstructPointersFromReferences(stage *Stage, instance *ParameterCategoryUse) {
	// insertion point for pointers field
	if instance.ParameterCategory != nil {
		reference.ParameterCategory = stage.ParameterCategorys_reference[instance.ParameterCategory]
	}
	// insertion point for slice of pointers field
}

func (reference *ParameterShape) GongReconstructPointersFromReferences(stage *Stage, instance *ParameterShape) {
	// insertion point for pointers field
	if instance.Parameter != nil {
		reference.Parameter = stage.Parameters_reference[instance.Parameter]
	}
	// insertion point for slice of pointers field
}

func (reference *ParametersAggregate) GongReconstructPointersFromReferences(stage *Stage, instance *ParametersAggregate) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Parameters = reference.Parameters[:0]
	for _, _b := range instance.Parameters {
		reference.Parameters = append(reference.Parameters, stage.Parameters_reference[_b])
	}
}

func (reference *ParametersAggregateShape) GongReconstructPointersFromReferences(stage *Stage, instance *ParametersAggregateShape) {
	// insertion point for pointers field
	if instance.ScenarioParameter != nil {
		reference.ScenarioParameter = stage.ParametersAggregates_reference[instance.ScenarioParameter]
	}
	// insertion point for slice of pointers field
}

func (reference *Position) GongReconstructPointersFromReferences(stage *Stage, instance *Position) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Repository) GongReconstructPointersFromReferences(stage *Stage, instance *Repository) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.ParameterUse = reference.ParameterUse[:0]
	for _, _b := range instance.ParameterUse {
		reference.ParameterUse = append(reference.ParameterUse, stage.ParameterShapes_reference[_b])
	}
	reference.GroupUse = reference.GroupUse[:0]
	for _, _b := range instance.GroupUse {
		reference.GroupUse = append(reference.GroupUse, stage.GroupUses_reference[_b])
	}
}

func (reference *Scenario) GongReconstructPointersFromReferences(stage *Stage, instance *Scenario) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Diagrams = reference.Diagrams[:0]
	for _, _b := range instance.Diagrams {
		reference.Diagrams = append(reference.Diagrams, stage.Diagrams_reference[_b])
	}
	reference.ActorStates = reference.ActorStates[:0]
	for _, _b := range instance.ActorStates {
		reference.ActorStates = append(reference.ActorStates, stage.ActorStates_reference[_b])
	}
	reference.ActorStateTransitions = reference.ActorStateTransitions[:0]
	for _, _b := range instance.ActorStateTransitions {
		reference.ActorStateTransitions = append(reference.ActorStateTransitions, stage.ActorStateTransitions_reference[_b])
	}
	reference.EvolutionDirections = reference.EvolutionDirections[:0]
	for _, _b := range instance.EvolutionDirections {
		reference.EvolutionDirections = append(reference.EvolutionDirections, stage.EvolutionDirections_reference[_b])
	}
	reference.Parameters = reference.Parameters[:0]
	for _, _b := range instance.Parameters {
		reference.Parameters = append(reference.Parameters, stage.Parameters_reference[_b])
	}
	reference.ParametersAggretates = reference.ParametersAggretates[:0]
	for _, _b := range instance.ParametersAggretates {
		reference.ParametersAggretates = append(reference.ParametersAggretates, stage.ParametersAggregates_reference[_b])
	}
}

func (reference *User) GongReconstructPointersFromReferences(stage *Stage, instance *User) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *UserUse) GongReconstructPointersFromReferences(stage *Stage, instance *UserUse) {
	// insertion point for pointers field
	if instance.User != nil {
		reference.User = stage.Users_reference[instance.User]
	}
	// insertion point for slice of pointers field
}

func (reference *Workspace) GongReconstructPointersFromReferences(stage *Stage, instance *Workspace) {
	// insertion point for pointers field
	if instance.SelectedDiagram != nil {
		reference.SelectedDiagram = stage.Diagrams_reference[instance.SelectedDiagram]
	}
	if instance.Default_EvolutionDirectionShape != nil {
		reference.Default_EvolutionDirectionShape = stage.EvolutionDirectionShapes_reference[instance.Default_EvolutionDirectionShape]
	}
	if instance.Default_ParameterShape != nil {
		reference.Default_ParameterShape = stage.ParameterShapes_reference[instance.Default_ParameterShape]
	}
	if instance.Default_ScenarioParameterShape != nil {
		reference.Default_ScenarioParameterShape = stage.ParametersAggregateShapes_reference[instance.Default_ScenarioParameterShape]
	}
	if instance.Default_ActorStateShape != nil {
		reference.Default_ActorStateShape = stage.ActorStateShapes_reference[instance.Default_ActorStateShape]
	}
	if instance.Default_ActorStateTransitionShape != nil {
		reference.Default_ActorStateTransitionShape = stage.ActorStateTransitionShapes_reference[instance.Default_ActorStateTransitionShape]
	}
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *ActorState) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ActorStateShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ActorState; _reference != nil {
		reference.ActorState = nil
		if _instance, ok := stage.ActorStates_instance[_reference]; ok {
			reference.ActorState = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ActorStateTransition) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.StartState; _reference != nil {
		reference.StartState = nil
		if _instance, ok := stage.ActorStates_instance[_reference]; ok {
			reference.StartState = _instance
		}
	}
	if _reference := reference.EndState; _reference != nil {
		reference.EndState = nil
		if _instance, ok := stage.ActorStates_instance[_reference]; ok {
			reference.EndState = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Justifications []*Parameter
	for _, _reference := range reference.Justifications {
		if _instance, ok := stage.Parameters_instance[_reference]; ok {
			_Justifications = append(_Justifications, _instance)
		}
	}
	reference.Justifications = _Justifications
}

func (reference *ActorStateTransitionShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ActorStateTransition; _reference != nil {
		reference.ActorStateTransition = nil
		if _instance, ok := stage.ActorStateTransitions_instance[_reference]; ok {
			reference.ActorStateTransition = _instance
		}
	}
	if _reference := reference.Start; _reference != nil {
		reference.Start = nil
		if _instance, ok := stage.ActorStateShapes_instance[_reference]; ok {
			reference.Start = _instance
		}
	}
	if _reference := reference.End; _reference != nil {
		reference.End = nil
		if _instance, ok := stage.ActorStateShapes_instance[_reference]; ok {
			reference.End = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _ControlPointShapes []*ControlPointShape
	for _, _reference := range reference.ControlPointShapes {
		if _instance, ok := stage.ControlPointShapes_instance[_reference]; ok {
			_ControlPointShapes = append(_ControlPointShapes, _instance)
		}
	}
	reference.ControlPointShapes = _ControlPointShapes
}

func (reference *Analysis) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Scenarios []*Scenario
	for _, _reference := range reference.Scenarios {
		if _instance, ok := stage.Scenarios_instance[_reference]; ok {
			_Scenarios = append(_Scenarios, _instance)
		}
	}
	reference.Scenarios = _Scenarios
	var _GroupUse []*GroupUse
	for _, _reference := range reference.GroupUse {
		if _instance, ok := stage.GroupUses_instance[_reference]; ok {
			_GroupUse = append(_GroupUse, _instance)
		}
	}
	reference.GroupUse = _GroupUse
	var _GeoObjectUse []*GeoObjectUse
	for _, _reference := range reference.GeoObjectUse {
		if _instance, ok := stage.GeoObjectUses_instance[_reference]; ok {
			_GeoObjectUse = append(_GeoObjectUse, _instance)
		}
	}
	reference.GeoObjectUse = _GeoObjectUse
	var _MapUse []*MapObjectUse
	for _, _reference := range reference.MapUse {
		if _instance, ok := stage.MapObjectUses_instance[_reference]; ok {
			_MapUse = append(_MapUse, _instance)
		}
	}
	reference.MapUse = _MapUse
}

func (reference *ControlPointShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Diagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _EvolutionDirectionShapes []*EvolutionDirectionShape
	for _, _reference := range reference.EvolutionDirectionShapes {
		if _instance, ok := stage.EvolutionDirectionShapes_instance[_reference]; ok {
			_EvolutionDirectionShapes = append(_EvolutionDirectionShapes, _instance)
		}
	}
	reference.EvolutionDirectionShapes = _EvolutionDirectionShapes
	var _EvolutionDirectionsWhoseNodeIsExpanded []*EvolutionDirection
	for _, _reference := range reference.EvolutionDirectionsWhoseNodeIsExpanded {
		if _instance, ok := stage.EvolutionDirections_instance[_reference]; ok {
			_EvolutionDirectionsWhoseNodeIsExpanded = append(_EvolutionDirectionsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.EvolutionDirectionsWhoseNodeIsExpanded = _EvolutionDirectionsWhoseNodeIsExpanded
	var _ActorStateShapes []*ActorStateShape
	for _, _reference := range reference.ActorStateShapes {
		if _instance, ok := stage.ActorStateShapes_instance[_reference]; ok {
			_ActorStateShapes = append(_ActorStateShapes, _instance)
		}
	}
	reference.ActorStateShapes = _ActorStateShapes
	var _ActorStatesWhoseNodeIsExpanded []*ActorState
	for _, _reference := range reference.ActorStatesWhoseNodeIsExpanded {
		if _instance, ok := stage.ActorStates_instance[_reference]; ok {
			_ActorStatesWhoseNodeIsExpanded = append(_ActorStatesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ActorStatesWhoseNodeIsExpanded = _ActorStatesWhoseNodeIsExpanded
	var _ParameterShapes []*ParameterShape
	for _, _reference := range reference.ParameterShapes {
		if _instance, ok := stage.ParameterShapes_instance[_reference]; ok {
			_ParameterShapes = append(_ParameterShapes, _instance)
		}
	}
	reference.ParameterShapes = _ParameterShapes
	var _ParametersWhoseNodeIsExpanded []*Parameter
	for _, _reference := range reference.ParametersWhoseNodeIsExpanded {
		if _instance, ok := stage.Parameters_instance[_reference]; ok {
			_ParametersWhoseNodeIsExpanded = append(_ParametersWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ParametersWhoseNodeIsExpanded = _ParametersWhoseNodeIsExpanded
	var _ScenarioParameterShapes []*ParametersAggregateShape
	for _, _reference := range reference.ScenarioParameterShapes {
		if _instance, ok := stage.ParametersAggregateShapes_instance[_reference]; ok {
			_ScenarioParameterShapes = append(_ScenarioParameterShapes, _instance)
		}
	}
	reference.ScenarioParameterShapes = _ScenarioParameterShapes
	var _ParametersAggregatesWhoseNodeIsExpanded []*ParametersAggregate
	for _, _reference := range reference.ParametersAggregatesWhoseNodeIsExpanded {
		if _instance, ok := stage.ParametersAggregates_instance[_reference]; ok {
			_ParametersAggregatesWhoseNodeIsExpanded = append(_ParametersAggregatesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ParametersAggregatesWhoseNodeIsExpanded = _ParametersAggregatesWhoseNodeIsExpanded
	var _ActorStateTransitionShapes []*ActorStateTransitionShape
	for _, _reference := range reference.ActorStateTransitionShapes {
		if _instance, ok := stage.ActorStateTransitionShapes_instance[_reference]; ok {
			_ActorStateTransitionShapes = append(_ActorStateTransitionShapes, _instance)
		}
	}
	reference.ActorStateTransitionShapes = _ActorStateTransitionShapes
	var _ActorStateTransitionsWhoseNodeIsExpanded []*ActorStateTransition
	for _, _reference := range reference.ActorStateTransitionsWhoseNodeIsExpanded {
		if _instance, ok := stage.ActorStateTransitions_instance[_reference]; ok {
			_ActorStateTransitionsWhoseNodeIsExpanded = append(_ActorStateTransitionsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ActorStateTransitionsWhoseNodeIsExpanded = _ActorStateTransitionsWhoseNodeIsExpanded
}

func (reference *Document) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _GeoObjectUse []*GeoObjectUse
	for _, _reference := range reference.GeoObjectUse {
		if _instance, ok := stage.GeoObjectUses_instance[_reference]; ok {
			_GeoObjectUse = append(_GeoObjectUse, _instance)
		}
	}
	reference.GeoObjectUse = _GeoObjectUse
}

func (reference *DocumentUse) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Document; _reference != nil {
		reference.Document = nil
		if _instance, ok := stage.Documents_instance[_reference]; ok {
			reference.Document = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *EvolutionDirection) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EvolutionDirectionShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.EvolutionDirection; _reference != nil {
		reference.EvolutionDirection = nil
		if _instance, ok := stage.EvolutionDirections_instance[_reference]; ok {
			reference.EvolutionDirection = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Foo) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GeoObject) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GeoObjectUse) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.GeoObject; _reference != nil {
		reference.GeoObject = nil
		if _instance, ok := stage.GeoObjects_instance[_reference]; ok {
			reference.GeoObject = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Group) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _UserUse []*UserUse
	for _, _reference := range reference.UserUse {
		if _instance, ok := stage.UserUses_instance[_reference]; ok {
			_UserUse = append(_UserUse, _instance)
		}
	}
	reference.UserUse = _UserUse
}

func (reference *GroupUse) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Group; _reference != nil {
		reference.Group = nil
		if _instance, ok := stage.Groups_instance[_reference]; ok {
			reference.Group = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Library) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Analyses []*Analysis
	for _, _reference := range reference.Analyses {
		if _instance, ok := stage.Analysiss_instance[_reference]; ok {
			_Analyses = append(_Analyses, _instance)
		}
	}
	reference.Analyses = _Analyses
	var _SubLibraries []*Library
	for _, _reference := range reference.SubLibraries {
		if _instance, ok := stage.Librarys_instance[_reference]; ok {
			_SubLibraries = append(_SubLibraries, _instance)
		}
	}
	reference.SubLibraries = _SubLibraries
	var _SubLibrariesWhoseNodeIsExpanded []*Library
	for _, _reference := range reference.SubLibrariesWhoseNodeIsExpanded {
		if _instance, ok := stage.Librarys_instance[_reference]; ok {
			_SubLibrariesWhoseNodeIsExpanded = append(_SubLibrariesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.SubLibrariesWhoseNodeIsExpanded = _SubLibrariesWhoseNodeIsExpanded
}

func (reference *MapObject) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MapObjectUse) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Map; _reference != nil {
		reference.Map = nil
		if _instance, ok := stage.MapObjects_instance[_reference]; ok {
			reference.Map = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Parameter) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _GroupUse []*GroupUse
	for _, _reference := range reference.GroupUse {
		if _instance, ok := stage.GroupUses_instance[_reference]; ok {
			_GroupUse = append(_GroupUse, _instance)
		}
	}
	reference.GroupUse = _GroupUse
	var _DocumentUse []*DocumentUse
	for _, _reference := range reference.DocumentUse {
		if _instance, ok := stage.DocumentUses_instance[_reference]; ok {
			_DocumentUse = append(_DocumentUse, _instance)
		}
	}
	reference.DocumentUse = _DocumentUse
	var _GeoObjectUse []*GeoObjectUse
	for _, _reference := range reference.GeoObjectUse {
		if _instance, ok := stage.GeoObjectUses_instance[_reference]; ok {
			_GeoObjectUse = append(_GeoObjectUse, _instance)
		}
	}
	reference.GeoObjectUse = _GeoObjectUse
}

func (reference *ParameterCategory) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _ParameterUse []*ParameterShape
	for _, _reference := range reference.ParameterUse {
		if _instance, ok := stage.ParameterShapes_instance[_reference]; ok {
			_ParameterUse = append(_ParameterUse, _instance)
		}
	}
	reference.ParameterUse = _ParameterUse
}

func (reference *ParameterCategoryUse) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ParameterCategory; _reference != nil {
		reference.ParameterCategory = nil
		if _instance, ok := stage.ParameterCategorys_instance[_reference]; ok {
			reference.ParameterCategory = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ParameterShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Parameter; _reference != nil {
		reference.Parameter = nil
		if _instance, ok := stage.Parameters_instance[_reference]; ok {
			reference.Parameter = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ParametersAggregate) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Parameters []*Parameter
	for _, _reference := range reference.Parameters {
		if _instance, ok := stage.Parameters_instance[_reference]; ok {
			_Parameters = append(_Parameters, _instance)
		}
	}
	reference.Parameters = _Parameters
}

func (reference *ParametersAggregateShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ScenarioParameter; _reference != nil {
		reference.ScenarioParameter = nil
		if _instance, ok := stage.ParametersAggregates_instance[_reference]; ok {
			reference.ScenarioParameter = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Position) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Repository) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _ParameterUse []*ParameterShape
	for _, _reference := range reference.ParameterUse {
		if _instance, ok := stage.ParameterShapes_instance[_reference]; ok {
			_ParameterUse = append(_ParameterUse, _instance)
		}
	}
	reference.ParameterUse = _ParameterUse
	var _GroupUse []*GroupUse
	for _, _reference := range reference.GroupUse {
		if _instance, ok := stage.GroupUses_instance[_reference]; ok {
			_GroupUse = append(_GroupUse, _instance)
		}
	}
	reference.GroupUse = _GroupUse
}

func (reference *Scenario) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Diagrams []*Diagram
	for _, _reference := range reference.Diagrams {
		if _instance, ok := stage.Diagrams_instance[_reference]; ok {
			_Diagrams = append(_Diagrams, _instance)
		}
	}
	reference.Diagrams = _Diagrams
	var _ActorStates []*ActorState
	for _, _reference := range reference.ActorStates {
		if _instance, ok := stage.ActorStates_instance[_reference]; ok {
			_ActorStates = append(_ActorStates, _instance)
		}
	}
	reference.ActorStates = _ActorStates
	var _ActorStateTransitions []*ActorStateTransition
	for _, _reference := range reference.ActorStateTransitions {
		if _instance, ok := stage.ActorStateTransitions_instance[_reference]; ok {
			_ActorStateTransitions = append(_ActorStateTransitions, _instance)
		}
	}
	reference.ActorStateTransitions = _ActorStateTransitions
	var _EvolutionDirections []*EvolutionDirection
	for _, _reference := range reference.EvolutionDirections {
		if _instance, ok := stage.EvolutionDirections_instance[_reference]; ok {
			_EvolutionDirections = append(_EvolutionDirections, _instance)
		}
	}
	reference.EvolutionDirections = _EvolutionDirections
	var _Parameters []*Parameter
	for _, _reference := range reference.Parameters {
		if _instance, ok := stage.Parameters_instance[_reference]; ok {
			_Parameters = append(_Parameters, _instance)
		}
	}
	reference.Parameters = _Parameters
	var _ParametersAggretates []*ParametersAggregate
	for _, _reference := range reference.ParametersAggretates {
		if _instance, ok := stage.ParametersAggregates_instance[_reference]; ok {
			_ParametersAggretates = append(_ParametersAggretates, _instance)
		}
	}
	reference.ParametersAggretates = _ParametersAggretates
}

func (reference *User) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *UserUse) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.User; _reference != nil {
		reference.User = nil
		if _instance, ok := stage.Users_instance[_reference]; ok {
			reference.User = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Workspace) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.SelectedDiagram; _reference != nil {
		reference.SelectedDiagram = nil
		if _instance, ok := stage.Diagrams_instance[_reference]; ok {
			reference.SelectedDiagram = _instance
		}
	}
	if _reference := reference.Default_EvolutionDirectionShape; _reference != nil {
		reference.Default_EvolutionDirectionShape = nil
		if _instance, ok := stage.EvolutionDirectionShapes_instance[_reference]; ok {
			reference.Default_EvolutionDirectionShape = _instance
		}
	}
	if _reference := reference.Default_ParameterShape; _reference != nil {
		reference.Default_ParameterShape = nil
		if _instance, ok := stage.ParameterShapes_instance[_reference]; ok {
			reference.Default_ParameterShape = _instance
		}
	}
	if _reference := reference.Default_ScenarioParameterShape; _reference != nil {
		reference.Default_ScenarioParameterShape = nil
		if _instance, ok := stage.ParametersAggregateShapes_instance[_reference]; ok {
			reference.Default_ScenarioParameterShape = _instance
		}
	}
	if _reference := reference.Default_ActorStateShape; _reference != nil {
		reference.Default_ActorStateShape = nil
		if _instance, ok := stage.ActorStateShapes_instance[_reference]; ok {
			reference.Default_ActorStateShape = _instance
		}
	}
	if _reference := reference.Default_ActorStateTransitionShape; _reference != nil {
		reference.Default_ActorStateTransitionShape = nil
		if _instance, ok := stage.ActorStateTransitionShapes_instance[_reference]; ok {
			reference.Default_ActorStateTransitionShape = _instance
		}
	}
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (actorstate *ActorState) GongDiff(stage *Stage, actorstateOther *ActorState) (diffs []string) {
	// insertion point for field diffs
	if actorstate.Name != actorstateOther.Name {
		diffs = append(diffs, actorstate.GongMarshallField(stage, "Name"))
	}
	if actorstate.Description != actorstateOther.Description {
		diffs = append(diffs, actorstate.GongMarshallField(stage, "Description"))
	}
	if actorstate.IsWithProbaility != actorstateOther.IsWithProbaility {
		diffs = append(diffs, actorstate.GongMarshallField(stage, "IsWithProbaility"))
	}
	if actorstate.Probability != actorstateOther.Probability {
		diffs = append(diffs, actorstate.GongMarshallField(stage, "Probability"))
	}
	if actorstate.ComputedPrefix != actorstateOther.ComputedPrefix {
		diffs = append(diffs, actorstate.GongMarshallField(stage, "ComputedPrefix"))
	}
	if actorstate.IsExpanded != actorstateOther.IsExpanded {
		diffs = append(diffs, actorstate.GongMarshallField(stage, "IsExpanded"))
	}
	if actorstate.LayoutDirection != actorstateOther.LayoutDirection {
		diffs = append(diffs, actorstate.GongMarshallField(stage, "LayoutDirection"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (actorstateshape *ActorStateShape) GongDiff(stage *Stage, actorstateshapeOther *ActorStateShape) (diffs []string) {
	// insertion point for field diffs
	if actorstateshape.Name != actorstateshapeOther.Name {
		diffs = append(diffs, actorstateshape.GongMarshallField(stage, "Name"))
	}
	if (actorstateshape.ActorState == nil) != (actorstateshapeOther.ActorState == nil) {
		diffs = append(diffs, actorstateshape.GongMarshallField(stage, "ActorState"))
	} else if actorstateshape.ActorState != nil && actorstateshapeOther.ActorState != nil {
		if actorstateshape.ActorState != actorstateshapeOther.ActorState {
			diffs = append(diffs, actorstateshape.GongMarshallField(stage, "ActorState"))
		}
	}
	if actorstateshape.X != actorstateshapeOther.X {
		diffs = append(diffs, actorstateshape.GongMarshallField(stage, "X"))
	}
	if actorstateshape.Y != actorstateshapeOther.Y {
		diffs = append(diffs, actorstateshape.GongMarshallField(stage, "Y"))
	}
	if actorstateshape.Width != actorstateshapeOther.Width {
		diffs = append(diffs, actorstateshape.GongMarshallField(stage, "Width"))
	}
	if actorstateshape.Height != actorstateshapeOther.Height {
		diffs = append(diffs, actorstateshape.GongMarshallField(stage, "Height"))
	}
	if actorstateshape.IsHidden != actorstateshapeOther.IsHidden {
		diffs = append(diffs, actorstateshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (actorstatetransition *ActorStateTransition) GongDiff(stage *Stage, actorstatetransitionOther *ActorStateTransition) (diffs []string) {
	// insertion point for field diffs
	if actorstatetransition.Name != actorstatetransitionOther.Name {
		diffs = append(diffs, actorstatetransition.GongMarshallField(stage, "Name"))
	}
	if (actorstatetransition.StartState == nil) != (actorstatetransitionOther.StartState == nil) {
		diffs = append(diffs, actorstatetransition.GongMarshallField(stage, "StartState"))
	} else if actorstatetransition.StartState != nil && actorstatetransitionOther.StartState != nil {
		if actorstatetransition.StartState != actorstatetransitionOther.StartState {
			diffs = append(diffs, actorstatetransition.GongMarshallField(stage, "StartState"))
		}
	}
	if (actorstatetransition.EndState == nil) != (actorstatetransitionOther.EndState == nil) {
		diffs = append(diffs, actorstatetransition.GongMarshallField(stage, "EndState"))
	} else if actorstatetransition.EndState != nil && actorstatetransitionOther.EndState != nil {
		if actorstatetransition.EndState != actorstatetransitionOther.EndState {
			diffs = append(diffs, actorstatetransition.GongMarshallField(stage, "EndState"))
		}
	}
	JustificationsDifferent := false
	if len(actorstatetransition.Justifications) != len(actorstatetransitionOther.Justifications) {
		JustificationsDifferent = true
	} else {
		for i := range actorstatetransition.Justifications {
			if (actorstatetransition.Justifications[i] == nil) != (actorstatetransitionOther.Justifications[i] == nil) {
				JustificationsDifferent = true
				break
			} else if actorstatetransition.Justifications[i] != nil && actorstatetransitionOther.Justifications[i] != nil {
				// this is a pointer comparaison
				if actorstatetransition.Justifications[i] != actorstatetransitionOther.Justifications[i] {
					JustificationsDifferent = true
					break
				}
			}
		}
	}
	if JustificationsDifferent {
		ops := Diff(stage, actorstatetransition, actorstatetransitionOther, "Justifications", actorstatetransitionOther.Justifications, actorstatetransition.Justifications)
		diffs = append(diffs, ops)
	}
	if actorstatetransition.ComputedPrefix != actorstatetransitionOther.ComputedPrefix {
		diffs = append(diffs, actorstatetransition.GongMarshallField(stage, "ComputedPrefix"))
	}
	if actorstatetransition.IsExpanded != actorstatetransitionOther.IsExpanded {
		diffs = append(diffs, actorstatetransition.GongMarshallField(stage, "IsExpanded"))
	}
	if actorstatetransition.LayoutDirection != actorstatetransitionOther.LayoutDirection {
		diffs = append(diffs, actorstatetransition.GongMarshallField(stage, "LayoutDirection"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (actorstatetransitionshape *ActorStateTransitionShape) GongDiff(stage *Stage, actorstatetransitionshapeOther *ActorStateTransitionShape) (diffs []string) {
	// insertion point for field diffs
	if actorstatetransitionshape.Name != actorstatetransitionshapeOther.Name {
		diffs = append(diffs, actorstatetransitionshape.GongMarshallField(stage, "Name"))
	}
	if (actorstatetransitionshape.ActorStateTransition == nil) != (actorstatetransitionshapeOther.ActorStateTransition == nil) {
		diffs = append(diffs, actorstatetransitionshape.GongMarshallField(stage, "ActorStateTransition"))
	} else if actorstatetransitionshape.ActorStateTransition != nil && actorstatetransitionshapeOther.ActorStateTransition != nil {
		if actorstatetransitionshape.ActorStateTransition != actorstatetransitionshapeOther.ActorStateTransition {
			diffs = append(diffs, actorstatetransitionshape.GongMarshallField(stage, "ActorStateTransition"))
		}
	}
	if (actorstatetransitionshape.Start == nil) != (actorstatetransitionshapeOther.Start == nil) {
		diffs = append(diffs, actorstatetransitionshape.GongMarshallField(stage, "Start"))
	} else if actorstatetransitionshape.Start != nil && actorstatetransitionshapeOther.Start != nil {
		if actorstatetransitionshape.Start != actorstatetransitionshapeOther.Start {
			diffs = append(diffs, actorstatetransitionshape.GongMarshallField(stage, "Start"))
		}
	}
	if (actorstatetransitionshape.End == nil) != (actorstatetransitionshapeOther.End == nil) {
		diffs = append(diffs, actorstatetransitionshape.GongMarshallField(stage, "End"))
	} else if actorstatetransitionshape.End != nil && actorstatetransitionshapeOther.End != nil {
		if actorstatetransitionshape.End != actorstatetransitionshapeOther.End {
			diffs = append(diffs, actorstatetransitionshape.GongMarshallField(stage, "End"))
		}
	}
	if actorstatetransitionshape.X != actorstatetransitionshapeOther.X {
		diffs = append(diffs, actorstatetransitionshape.GongMarshallField(stage, "X"))
	}
	if actorstatetransitionshape.Y != actorstatetransitionshapeOther.Y {
		diffs = append(diffs, actorstatetransitionshape.GongMarshallField(stage, "Y"))
	}
	if actorstatetransitionshape.Width != actorstatetransitionshapeOther.Width {
		diffs = append(diffs, actorstatetransitionshape.GongMarshallField(stage, "Width"))
	}
	if actorstatetransitionshape.Height != actorstatetransitionshapeOther.Height {
		diffs = append(diffs, actorstatetransitionshape.GongMarshallField(stage, "Height"))
	}
	if actorstatetransitionshape.IsHidden != actorstatetransitionshapeOther.IsHidden {
		diffs = append(diffs, actorstatetransitionshape.GongMarshallField(stage, "IsHidden"))
	}
	ControlPointShapesDifferent := false
	if len(actorstatetransitionshape.ControlPointShapes) != len(actorstatetransitionshapeOther.ControlPointShapes) {
		ControlPointShapesDifferent = true
	} else {
		for i := range actorstatetransitionshape.ControlPointShapes {
			if (actorstatetransitionshape.ControlPointShapes[i] == nil) != (actorstatetransitionshapeOther.ControlPointShapes[i] == nil) {
				ControlPointShapesDifferent = true
				break
			} else if actorstatetransitionshape.ControlPointShapes[i] != nil && actorstatetransitionshapeOther.ControlPointShapes[i] != nil {
				// this is a pointer comparaison
				if actorstatetransitionshape.ControlPointShapes[i] != actorstatetransitionshapeOther.ControlPointShapes[i] {
					ControlPointShapesDifferent = true
					break
				}
			}
		}
	}
	if ControlPointShapesDifferent {
		ops := Diff(stage, actorstatetransitionshape, actorstatetransitionshapeOther, "ControlPointShapes", actorstatetransitionshapeOther.ControlPointShapes, actorstatetransitionshape.ControlPointShapes)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (analysis *Analysis) GongDiff(stage *Stage, analysisOther *Analysis) (diffs []string) {
	// insertion point for field diffs
	if analysis.Name != analysisOther.Name {
		diffs = append(diffs, analysis.GongMarshallField(stage, "Name"))
	}
	if analysis.Description != analysisOther.Description {
		diffs = append(diffs, analysis.GongMarshallField(stage, "Description"))
	}
	ScenariosDifferent := false
	if len(analysis.Scenarios) != len(analysisOther.Scenarios) {
		ScenariosDifferent = true
	} else {
		for i := range analysis.Scenarios {
			if (analysis.Scenarios[i] == nil) != (analysisOther.Scenarios[i] == nil) {
				ScenariosDifferent = true
				break
			} else if analysis.Scenarios[i] != nil && analysisOther.Scenarios[i] != nil {
				// this is a pointer comparaison
				if analysis.Scenarios[i] != analysisOther.Scenarios[i] {
					ScenariosDifferent = true
					break
				}
			}
		}
	}
	if ScenariosDifferent {
		ops := Diff(stage, analysis, analysisOther, "Scenarios", analysisOther.Scenarios, analysis.Scenarios)
		diffs = append(diffs, ops)
	}
	if analysis.IsScenariosNodeExpanded != analysisOther.IsScenariosNodeExpanded {
		diffs = append(diffs, analysis.GongMarshallField(stage, "IsScenariosNodeExpanded"))
	}
	GroupUseDifferent := false
	if len(analysis.GroupUse) != len(analysisOther.GroupUse) {
		GroupUseDifferent = true
	} else {
		for i := range analysis.GroupUse {
			if (analysis.GroupUse[i] == nil) != (analysisOther.GroupUse[i] == nil) {
				GroupUseDifferent = true
				break
			} else if analysis.GroupUse[i] != nil && analysisOther.GroupUse[i] != nil {
				// this is a pointer comparaison
				if analysis.GroupUse[i] != analysisOther.GroupUse[i] {
					GroupUseDifferent = true
					break
				}
			}
		}
	}
	if GroupUseDifferent {
		ops := Diff(stage, analysis, analysisOther, "GroupUse", analysisOther.GroupUse, analysis.GroupUse)
		diffs = append(diffs, ops)
	}
	if analysis.IsGroupUseNodeExpanded != analysisOther.IsGroupUseNodeExpanded {
		diffs = append(diffs, analysis.GongMarshallField(stage, "IsGroupUseNodeExpanded"))
	}
	GeoObjectUseDifferent := false
	if len(analysis.GeoObjectUse) != len(analysisOther.GeoObjectUse) {
		GeoObjectUseDifferent = true
	} else {
		for i := range analysis.GeoObjectUse {
			if (analysis.GeoObjectUse[i] == nil) != (analysisOther.GeoObjectUse[i] == nil) {
				GeoObjectUseDifferent = true
				break
			} else if analysis.GeoObjectUse[i] != nil && analysisOther.GeoObjectUse[i] != nil {
				// this is a pointer comparaison
				if analysis.GeoObjectUse[i] != analysisOther.GeoObjectUse[i] {
					GeoObjectUseDifferent = true
					break
				}
			}
		}
	}
	if GeoObjectUseDifferent {
		ops := Diff(stage, analysis, analysisOther, "GeoObjectUse", analysisOther.GeoObjectUse, analysis.GeoObjectUse)
		diffs = append(diffs, ops)
	}
	if analysis.IsGeoObjectUseNodeExpanded != analysisOther.IsGeoObjectUseNodeExpanded {
		diffs = append(diffs, analysis.GongMarshallField(stage, "IsGeoObjectUseNodeExpanded"))
	}
	MapUseDifferent := false
	if len(analysis.MapUse) != len(analysisOther.MapUse) {
		MapUseDifferent = true
	} else {
		for i := range analysis.MapUse {
			if (analysis.MapUse[i] == nil) != (analysisOther.MapUse[i] == nil) {
				MapUseDifferent = true
				break
			} else if analysis.MapUse[i] != nil && analysisOther.MapUse[i] != nil {
				// this is a pointer comparaison
				if analysis.MapUse[i] != analysisOther.MapUse[i] {
					MapUseDifferent = true
					break
				}
			}
		}
	}
	if MapUseDifferent {
		ops := Diff(stage, analysis, analysisOther, "MapUse", analysisOther.MapUse, analysis.MapUse)
		diffs = append(diffs, ops)
	}
	if analysis.IsMapUseNodeExpanded != analysisOther.IsMapUseNodeExpanded {
		diffs = append(diffs, analysis.GongMarshallField(stage, "IsMapUseNodeExpanded"))
	}
	if analysis.ComputedPrefix != analysisOther.ComputedPrefix {
		diffs = append(diffs, analysis.GongMarshallField(stage, "ComputedPrefix"))
	}
	if analysis.IsExpanded != analysisOther.IsExpanded {
		diffs = append(diffs, analysis.GongMarshallField(stage, "IsExpanded"))
	}
	if analysis.LayoutDirection != analysisOther.LayoutDirection {
		diffs = append(diffs, analysis.GongMarshallField(stage, "LayoutDirection"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (controlpointshape *ControlPointShape) GongDiff(stage *Stage, controlpointshapeOther *ControlPointShape) (diffs []string) {
	// insertion point for field diffs
	if controlpointshape.Name != controlpointshapeOther.Name {
		diffs = append(diffs, controlpointshape.GongMarshallField(stage, "Name"))
	}
	if controlpointshape.X_Relative != controlpointshapeOther.X_Relative {
		diffs = append(diffs, controlpointshape.GongMarshallField(stage, "X_Relative"))
	}
	if controlpointshape.Y_Relative != controlpointshapeOther.Y_Relative {
		diffs = append(diffs, controlpointshape.GongMarshallField(stage, "Y_Relative"))
	}
	if controlpointshape.IsStartShapeTheClosestShape != controlpointshapeOther.IsStartShapeTheClosestShape {
		diffs = append(diffs, controlpointshape.GongMarshallField(stage, "IsStartShapeTheClosestShape"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (diagram *Diagram) GongDiff(stage *Stage, diagramOther *Diagram) (diffs []string) {
	// insertion point for field diffs
	if diagram.Name != diagramOther.Name {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Name"))
	}
	if diagram.ComputedPrefix != diagramOther.ComputedPrefix {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ComputedPrefix"))
	}
	if diagram.IsExpanded != diagramOther.IsExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsExpanded"))
	}
	if diagram.LayoutDirection != diagramOther.LayoutDirection {
		diffs = append(diffs, diagram.GongMarshallField(stage, "LayoutDirection"))
	}
	if diagram.IsChecked != diagramOther.IsChecked {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsChecked"))
	}
	if diagram.IsShowPrefix != diagramOther.IsShowPrefix {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsShowPrefix"))
	}
	if diagram.Description != diagramOther.Description {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Description"))
	}
	EvolutionDirectionShapesDifferent := false
	if len(diagram.EvolutionDirectionShapes) != len(diagramOther.EvolutionDirectionShapes) {
		EvolutionDirectionShapesDifferent = true
	} else {
		for i := range diagram.EvolutionDirectionShapes {
			if (diagram.EvolutionDirectionShapes[i] == nil) != (diagramOther.EvolutionDirectionShapes[i] == nil) {
				EvolutionDirectionShapesDifferent = true
				break
			} else if diagram.EvolutionDirectionShapes[i] != nil && diagramOther.EvolutionDirectionShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.EvolutionDirectionShapes[i] != diagramOther.EvolutionDirectionShapes[i] {
					EvolutionDirectionShapesDifferent = true
					break
				}
			}
		}
	}
	if EvolutionDirectionShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "EvolutionDirectionShapes", diagramOther.EvolutionDirectionShapes, diagram.EvolutionDirectionShapes)
		diffs = append(diffs, ops)
	}
	EvolutionDirectionsWhoseNodeIsExpandedDifferent := false
	if len(diagram.EvolutionDirectionsWhoseNodeIsExpanded) != len(diagramOther.EvolutionDirectionsWhoseNodeIsExpanded) {
		EvolutionDirectionsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.EvolutionDirectionsWhoseNodeIsExpanded {
			if (diagram.EvolutionDirectionsWhoseNodeIsExpanded[i] == nil) != (diagramOther.EvolutionDirectionsWhoseNodeIsExpanded[i] == nil) {
				EvolutionDirectionsWhoseNodeIsExpandedDifferent = true
				break
			} else if diagram.EvolutionDirectionsWhoseNodeIsExpanded[i] != nil && diagramOther.EvolutionDirectionsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.EvolutionDirectionsWhoseNodeIsExpanded[i] != diagramOther.EvolutionDirectionsWhoseNodeIsExpanded[i] {
					EvolutionDirectionsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if EvolutionDirectionsWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagram, diagramOther, "EvolutionDirectionsWhoseNodeIsExpanded", diagramOther.EvolutionDirectionsWhoseNodeIsExpanded, diagram.EvolutionDirectionsWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if diagram.IsEvolutionDirectionsNodeExpanded != diagramOther.IsEvolutionDirectionsNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsEvolutionDirectionsNodeExpanded"))
	}
	ActorStateShapesDifferent := false
	if len(diagram.ActorStateShapes) != len(diagramOther.ActorStateShapes) {
		ActorStateShapesDifferent = true
	} else {
		for i := range diagram.ActorStateShapes {
			if (diagram.ActorStateShapes[i] == nil) != (diagramOther.ActorStateShapes[i] == nil) {
				ActorStateShapesDifferent = true
				break
			} else if diagram.ActorStateShapes[i] != nil && diagramOther.ActorStateShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.ActorStateShapes[i] != diagramOther.ActorStateShapes[i] {
					ActorStateShapesDifferent = true
					break
				}
			}
		}
	}
	if ActorStateShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "ActorStateShapes", diagramOther.ActorStateShapes, diagram.ActorStateShapes)
		diffs = append(diffs, ops)
	}
	ActorStatesWhoseNodeIsExpandedDifferent := false
	if len(diagram.ActorStatesWhoseNodeIsExpanded) != len(diagramOther.ActorStatesWhoseNodeIsExpanded) {
		ActorStatesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.ActorStatesWhoseNodeIsExpanded {
			if (diagram.ActorStatesWhoseNodeIsExpanded[i] == nil) != (diagramOther.ActorStatesWhoseNodeIsExpanded[i] == nil) {
				ActorStatesWhoseNodeIsExpandedDifferent = true
				break
			} else if diagram.ActorStatesWhoseNodeIsExpanded[i] != nil && diagramOther.ActorStatesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.ActorStatesWhoseNodeIsExpanded[i] != diagramOther.ActorStatesWhoseNodeIsExpanded[i] {
					ActorStatesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ActorStatesWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagram, diagramOther, "ActorStatesWhoseNodeIsExpanded", diagramOther.ActorStatesWhoseNodeIsExpanded, diagram.ActorStatesWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if diagram.IsActorStatesNodeExpanded != diagramOther.IsActorStatesNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsActorStatesNodeExpanded"))
	}
	ParameterShapesDifferent := false
	if len(diagram.ParameterShapes) != len(diagramOther.ParameterShapes) {
		ParameterShapesDifferent = true
	} else {
		for i := range diagram.ParameterShapes {
			if (diagram.ParameterShapes[i] == nil) != (diagramOther.ParameterShapes[i] == nil) {
				ParameterShapesDifferent = true
				break
			} else if diagram.ParameterShapes[i] != nil && diagramOther.ParameterShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.ParameterShapes[i] != diagramOther.ParameterShapes[i] {
					ParameterShapesDifferent = true
					break
				}
			}
		}
	}
	if ParameterShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "ParameterShapes", diagramOther.ParameterShapes, diagram.ParameterShapes)
		diffs = append(diffs, ops)
	}
	ParametersWhoseNodeIsExpandedDifferent := false
	if len(diagram.ParametersWhoseNodeIsExpanded) != len(diagramOther.ParametersWhoseNodeIsExpanded) {
		ParametersWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.ParametersWhoseNodeIsExpanded {
			if (diagram.ParametersWhoseNodeIsExpanded[i] == nil) != (diagramOther.ParametersWhoseNodeIsExpanded[i] == nil) {
				ParametersWhoseNodeIsExpandedDifferent = true
				break
			} else if diagram.ParametersWhoseNodeIsExpanded[i] != nil && diagramOther.ParametersWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.ParametersWhoseNodeIsExpanded[i] != diagramOther.ParametersWhoseNodeIsExpanded[i] {
					ParametersWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ParametersWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagram, diagramOther, "ParametersWhoseNodeIsExpanded", diagramOther.ParametersWhoseNodeIsExpanded, diagram.ParametersWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if diagram.IsParametersNodeExpanded != diagramOther.IsParametersNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsParametersNodeExpanded"))
	}
	ScenarioParameterShapesDifferent := false
	if len(diagram.ScenarioParameterShapes) != len(diagramOther.ScenarioParameterShapes) {
		ScenarioParameterShapesDifferent = true
	} else {
		for i := range diagram.ScenarioParameterShapes {
			if (diagram.ScenarioParameterShapes[i] == nil) != (diagramOther.ScenarioParameterShapes[i] == nil) {
				ScenarioParameterShapesDifferent = true
				break
			} else if diagram.ScenarioParameterShapes[i] != nil && diagramOther.ScenarioParameterShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.ScenarioParameterShapes[i] != diagramOther.ScenarioParameterShapes[i] {
					ScenarioParameterShapesDifferent = true
					break
				}
			}
		}
	}
	if ScenarioParameterShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "ScenarioParameterShapes", diagramOther.ScenarioParameterShapes, diagram.ScenarioParameterShapes)
		diffs = append(diffs, ops)
	}
	ParametersAggregatesWhoseNodeIsExpandedDifferent := false
	if len(diagram.ParametersAggregatesWhoseNodeIsExpanded) != len(diagramOther.ParametersAggregatesWhoseNodeIsExpanded) {
		ParametersAggregatesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.ParametersAggregatesWhoseNodeIsExpanded {
			if (diagram.ParametersAggregatesWhoseNodeIsExpanded[i] == nil) != (diagramOther.ParametersAggregatesWhoseNodeIsExpanded[i] == nil) {
				ParametersAggregatesWhoseNodeIsExpandedDifferent = true
				break
			} else if diagram.ParametersAggregatesWhoseNodeIsExpanded[i] != nil && diagramOther.ParametersAggregatesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.ParametersAggregatesWhoseNodeIsExpanded[i] != diagramOther.ParametersAggregatesWhoseNodeIsExpanded[i] {
					ParametersAggregatesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ParametersAggregatesWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagram, diagramOther, "ParametersAggregatesWhoseNodeIsExpanded", diagramOther.ParametersAggregatesWhoseNodeIsExpanded, diagram.ParametersAggregatesWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if diagram.IsParametersAggregatesNodeExpanded != diagramOther.IsParametersAggregatesNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsParametersAggregatesNodeExpanded"))
	}
	ActorStateTransitionShapesDifferent := false
	if len(diagram.ActorStateTransitionShapes) != len(diagramOther.ActorStateTransitionShapes) {
		ActorStateTransitionShapesDifferent = true
	} else {
		for i := range diagram.ActorStateTransitionShapes {
			if (diagram.ActorStateTransitionShapes[i] == nil) != (diagramOther.ActorStateTransitionShapes[i] == nil) {
				ActorStateTransitionShapesDifferent = true
				break
			} else if diagram.ActorStateTransitionShapes[i] != nil && diagramOther.ActorStateTransitionShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.ActorStateTransitionShapes[i] != diagramOther.ActorStateTransitionShapes[i] {
					ActorStateTransitionShapesDifferent = true
					break
				}
			}
		}
	}
	if ActorStateTransitionShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "ActorStateTransitionShapes", diagramOther.ActorStateTransitionShapes, diagram.ActorStateTransitionShapes)
		diffs = append(diffs, ops)
	}
	ActorStateTransitionsWhoseNodeIsExpandedDifferent := false
	if len(diagram.ActorStateTransitionsWhoseNodeIsExpanded) != len(diagramOther.ActorStateTransitionsWhoseNodeIsExpanded) {
		ActorStateTransitionsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.ActorStateTransitionsWhoseNodeIsExpanded {
			if (diagram.ActorStateTransitionsWhoseNodeIsExpanded[i] == nil) != (diagramOther.ActorStateTransitionsWhoseNodeIsExpanded[i] == nil) {
				ActorStateTransitionsWhoseNodeIsExpandedDifferent = true
				break
			} else if diagram.ActorStateTransitionsWhoseNodeIsExpanded[i] != nil && diagramOther.ActorStateTransitionsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.ActorStateTransitionsWhoseNodeIsExpanded[i] != diagramOther.ActorStateTransitionsWhoseNodeIsExpanded[i] {
					ActorStateTransitionsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ActorStateTransitionsWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagram, diagramOther, "ActorStateTransitionsWhoseNodeIsExpanded", diagramOther.ActorStateTransitionsWhoseNodeIsExpanded, diagram.ActorStateTransitionsWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if diagram.IsActorStateTransitionsNodeExpanded != diagramOther.IsActorStateTransitionsNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsActorStateTransitionsNodeExpanded"))
	}
	if diagram.AxisOrign_X != diagramOther.AxisOrign_X {
		diffs = append(diffs, diagram.GongMarshallField(stage, "AxisOrign_X"))
	}
	if diagram.AxisOrign_Y != diagramOther.AxisOrign_Y {
		diffs = append(diffs, diagram.GongMarshallField(stage, "AxisOrign_Y"))
	}
	if diagram.VerticalAxis_Top_Y != diagramOther.VerticalAxis_Top_Y {
		diffs = append(diffs, diagram.GongMarshallField(stage, "VerticalAxis_Top_Y"))
	}
	if diagram.VerticalAxis_Bottom_Y != diagramOther.VerticalAxis_Bottom_Y {
		diffs = append(diffs, diagram.GongMarshallField(stage, "VerticalAxis_Bottom_Y"))
	}
	if diagram.VerticalAxis_StrokeWidth != diagramOther.VerticalAxis_StrokeWidth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "VerticalAxis_StrokeWidth"))
	}
	if diagram.HorizontalAxis_Right_X != diagramOther.HorizontalAxis_Right_X {
		diffs = append(diffs, diagram.GongMarshallField(stage, "HorizontalAxis_Right_X"))
	}
	if diagram.Start != diagramOther.Start {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Start"))
	}
	if diagram.End != diagramOther.End {
		diffs = append(diffs, diagram.GongMarshallField(stage, "End"))
	}
	if diagram.NumberOfYearsBetweenTicks != diagramOther.NumberOfYearsBetweenTicks {
		diffs = append(diffs, diagram.GongMarshallField(stage, "NumberOfYearsBetweenTicks"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (document *Document) GongDiff(stage *Stage, documentOther *Document) (diffs []string) {
	// insertion point for field diffs
	if document.Name != documentOther.Name {
		diffs = append(diffs, document.GongMarshallField(stage, "Name"))
	}
	GeoObjectUseDifferent := false
	if len(document.GeoObjectUse) != len(documentOther.GeoObjectUse) {
		GeoObjectUseDifferent = true
	} else {
		for i := range document.GeoObjectUse {
			if (document.GeoObjectUse[i] == nil) != (documentOther.GeoObjectUse[i] == nil) {
				GeoObjectUseDifferent = true
				break
			} else if document.GeoObjectUse[i] != nil && documentOther.GeoObjectUse[i] != nil {
				// this is a pointer comparaison
				if document.GeoObjectUse[i] != documentOther.GeoObjectUse[i] {
					GeoObjectUseDifferent = true
					break
				}
			}
		}
	}
	if GeoObjectUseDifferent {
		ops := Diff(stage, document, documentOther, "GeoObjectUse", documentOther.GeoObjectUse, document.GeoObjectUse)
		diffs = append(diffs, ops)
	}
	if document.ComputedPrefix != documentOther.ComputedPrefix {
		diffs = append(diffs, document.GongMarshallField(stage, "ComputedPrefix"))
	}
	if document.IsExpanded != documentOther.IsExpanded {
		diffs = append(diffs, document.GongMarshallField(stage, "IsExpanded"))
	}
	if document.LayoutDirection != documentOther.LayoutDirection {
		diffs = append(diffs, document.GongMarshallField(stage, "LayoutDirection"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (documentuse *DocumentUse) GongDiff(stage *Stage, documentuseOther *DocumentUse) (diffs []string) {
	// insertion point for field diffs
	if documentuse.Name != documentuseOther.Name {
		diffs = append(diffs, documentuse.GongMarshallField(stage, "Name"))
	}
	if (documentuse.Document == nil) != (documentuseOther.Document == nil) {
		diffs = append(diffs, documentuse.GongMarshallField(stage, "Document"))
	} else if documentuse.Document != nil && documentuseOther.Document != nil {
		if documentuse.Document != documentuseOther.Document {
			diffs = append(diffs, documentuse.GongMarshallField(stage, "Document"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (evolutiondirection *EvolutionDirection) GongDiff(stage *Stage, evolutiondirectionOther *EvolutionDirection) (diffs []string) {
	// insertion point for field diffs
	if evolutiondirection.Name != evolutiondirectionOther.Name {
		diffs = append(diffs, evolutiondirection.GongMarshallField(stage, "Name"))
	}
	if evolutiondirection.Description != evolutiondirectionOther.Description {
		diffs = append(diffs, evolutiondirection.GongMarshallField(stage, "Description"))
	}
	if evolutiondirection.ComputedPrefix != evolutiondirectionOther.ComputedPrefix {
		diffs = append(diffs, evolutiondirection.GongMarshallField(stage, "ComputedPrefix"))
	}
	if evolutiondirection.IsExpanded != evolutiondirectionOther.IsExpanded {
		diffs = append(diffs, evolutiondirection.GongMarshallField(stage, "IsExpanded"))
	}
	if evolutiondirection.LayoutDirection != evolutiondirectionOther.LayoutDirection {
		diffs = append(diffs, evolutiondirection.GongMarshallField(stage, "LayoutDirection"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (evolutiondirectionshape *EvolutionDirectionShape) GongDiff(stage *Stage, evolutiondirectionshapeOther *EvolutionDirectionShape) (diffs []string) {
	// insertion point for field diffs
	if evolutiondirectionshape.Name != evolutiondirectionshapeOther.Name {
		diffs = append(diffs, evolutiondirectionshape.GongMarshallField(stage, "Name"))
	}
	if (evolutiondirectionshape.EvolutionDirection == nil) != (evolutiondirectionshapeOther.EvolutionDirection == nil) {
		diffs = append(diffs, evolutiondirectionshape.GongMarshallField(stage, "EvolutionDirection"))
	} else if evolutiondirectionshape.EvolutionDirection != nil && evolutiondirectionshapeOther.EvolutionDirection != nil {
		if evolutiondirectionshape.EvolutionDirection != evolutiondirectionshapeOther.EvolutionDirection {
			diffs = append(diffs, evolutiondirectionshape.GongMarshallField(stage, "EvolutionDirection"))
		}
	}
	if evolutiondirectionshape.X != evolutiondirectionshapeOther.X {
		diffs = append(diffs, evolutiondirectionshape.GongMarshallField(stage, "X"))
	}
	if evolutiondirectionshape.Y != evolutiondirectionshapeOther.Y {
		diffs = append(diffs, evolutiondirectionshape.GongMarshallField(stage, "Y"))
	}
	if evolutiondirectionshape.Width != evolutiondirectionshapeOther.Width {
		diffs = append(diffs, evolutiondirectionshape.GongMarshallField(stage, "Width"))
	}
	if evolutiondirectionshape.Height != evolutiondirectionshapeOther.Height {
		diffs = append(diffs, evolutiondirectionshape.GongMarshallField(stage, "Height"))
	}
	if evolutiondirectionshape.IsHidden != evolutiondirectionshapeOther.IsHidden {
		diffs = append(diffs, evolutiondirectionshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (foo *Foo) GongDiff(stage *Stage, fooOther *Foo) (diffs []string) {
	// insertion point for field diffs
	if foo.Name != fooOther.Name {
		diffs = append(diffs, foo.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (geoobject *GeoObject) GongDiff(stage *Stage, geoobjectOther *GeoObject) (diffs []string) {
	// insertion point for field diffs
	if geoobject.Name != geoobjectOther.Name {
		diffs = append(diffs, geoobject.GongMarshallField(stage, "Name"))
	}
	if geoobject.ComputedPrefix != geoobjectOther.ComputedPrefix {
		diffs = append(diffs, geoobject.GongMarshallField(stage, "ComputedPrefix"))
	}
	if geoobject.IsExpanded != geoobjectOther.IsExpanded {
		diffs = append(diffs, geoobject.GongMarshallField(stage, "IsExpanded"))
	}
	if geoobject.LayoutDirection != geoobjectOther.LayoutDirection {
		diffs = append(diffs, geoobject.GongMarshallField(stage, "LayoutDirection"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (geoobjectuse *GeoObjectUse) GongDiff(stage *Stage, geoobjectuseOther *GeoObjectUse) (diffs []string) {
	// insertion point for field diffs
	if geoobjectuse.Name != geoobjectuseOther.Name {
		diffs = append(diffs, geoobjectuse.GongMarshallField(stage, "Name"))
	}
	if (geoobjectuse.GeoObject == nil) != (geoobjectuseOther.GeoObject == nil) {
		diffs = append(diffs, geoobjectuse.GongMarshallField(stage, "GeoObject"))
	} else if geoobjectuse.GeoObject != nil && geoobjectuseOther.GeoObject != nil {
		if geoobjectuse.GeoObject != geoobjectuseOther.GeoObject {
			diffs = append(diffs, geoobjectuse.GongMarshallField(stage, "GeoObject"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (group *Group) GongDiff(stage *Stage, groupOther *Group) (diffs []string) {
	// insertion point for field diffs
	if group.Name != groupOther.Name {
		diffs = append(diffs, group.GongMarshallField(stage, "Name"))
	}
	UserUseDifferent := false
	if len(group.UserUse) != len(groupOther.UserUse) {
		UserUseDifferent = true
	} else {
		for i := range group.UserUse {
			if (group.UserUse[i] == nil) != (groupOther.UserUse[i] == nil) {
				UserUseDifferent = true
				break
			} else if group.UserUse[i] != nil && groupOther.UserUse[i] != nil {
				// this is a pointer comparaison
				if group.UserUse[i] != groupOther.UserUse[i] {
					UserUseDifferent = true
					break
				}
			}
		}
	}
	if UserUseDifferent {
		ops := Diff(stage, group, groupOther, "UserUse", groupOther.UserUse, group.UserUse)
		diffs = append(diffs, ops)
	}
	if group.ComputedPrefix != groupOther.ComputedPrefix {
		diffs = append(diffs, group.GongMarshallField(stage, "ComputedPrefix"))
	}
	if group.IsExpanded != groupOther.IsExpanded {
		diffs = append(diffs, group.GongMarshallField(stage, "IsExpanded"))
	}
	if group.LayoutDirection != groupOther.LayoutDirection {
		diffs = append(diffs, group.GongMarshallField(stage, "LayoutDirection"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (groupuse *GroupUse) GongDiff(stage *Stage, groupuseOther *GroupUse) (diffs []string) {
	// insertion point for field diffs
	if groupuse.Name != groupuseOther.Name {
		diffs = append(diffs, groupuse.GongMarshallField(stage, "Name"))
	}
	if (groupuse.Group == nil) != (groupuseOther.Group == nil) {
		diffs = append(diffs, groupuse.GongMarshallField(stage, "Group"))
	} else if groupuse.Group != nil && groupuseOther.Group != nil {
		if groupuse.Group != groupuseOther.Group {
			diffs = append(diffs, groupuse.GongMarshallField(stage, "Group"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (library *Library) GongDiff(stage *Stage, libraryOther *Library) (diffs []string) {
	// insertion point for field diffs
	if library.Name != libraryOther.Name {
		diffs = append(diffs, library.GongMarshallField(stage, "Name"))
	}
	if library.Description != libraryOther.Description {
		diffs = append(diffs, library.GongMarshallField(stage, "Description"))
	}
	if library.ComputedPrefix != libraryOther.ComputedPrefix {
		diffs = append(diffs, library.GongMarshallField(stage, "ComputedPrefix"))
	}
	if library.IsExpanded != libraryOther.IsExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsExpanded"))
	}
	if library.LayoutDirection != libraryOther.LayoutDirection {
		diffs = append(diffs, library.GongMarshallField(stage, "LayoutDirection"))
	}
	if library.IsRootLibrary != libraryOther.IsRootLibrary {
		diffs = append(diffs, library.GongMarshallField(stage, "IsRootLibrary"))
	}
	AnalysesDifferent := false
	if len(library.Analyses) != len(libraryOther.Analyses) {
		AnalysesDifferent = true
	} else {
		for i := range library.Analyses {
			if (library.Analyses[i] == nil) != (libraryOther.Analyses[i] == nil) {
				AnalysesDifferent = true
				break
			} else if library.Analyses[i] != nil && libraryOther.Analyses[i] != nil {
				// this is a pointer comparaison
				if library.Analyses[i] != libraryOther.Analyses[i] {
					AnalysesDifferent = true
					break
				}
			}
		}
	}
	if AnalysesDifferent {
		ops := Diff(stage, library, libraryOther, "Analyses", libraryOther.Analyses, library.Analyses)
		diffs = append(diffs, ops)
	}
	if library.IsAnalysesNodeExpanded != libraryOther.IsAnalysesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsAnalysesNodeExpanded"))
	}
	SubLibrariesDifferent := false
	if len(library.SubLibraries) != len(libraryOther.SubLibraries) {
		SubLibrariesDifferent = true
	} else {
		for i := range library.SubLibraries {
			if (library.SubLibraries[i] == nil) != (libraryOther.SubLibraries[i] == nil) {
				SubLibrariesDifferent = true
				break
			} else if library.SubLibraries[i] != nil && libraryOther.SubLibraries[i] != nil {
				// this is a pointer comparaison
				if library.SubLibraries[i] != libraryOther.SubLibraries[i] {
					SubLibrariesDifferent = true
					break
				}
			}
		}
	}
	if SubLibrariesDifferent {
		ops := Diff(stage, library, libraryOther, "SubLibraries", libraryOther.SubLibraries, library.SubLibraries)
		diffs = append(diffs, ops)
	}
	if library.IsSubLibrariesNodeExpanded != libraryOther.IsSubLibrariesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsSubLibrariesNodeExpanded"))
	}
	SubLibrariesWhoseNodeIsExpandedDifferent := false
	if len(library.SubLibrariesWhoseNodeIsExpanded) != len(libraryOther.SubLibrariesWhoseNodeIsExpanded) {
		SubLibrariesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.SubLibrariesWhoseNodeIsExpanded {
			if (library.SubLibrariesWhoseNodeIsExpanded[i] == nil) != (libraryOther.SubLibrariesWhoseNodeIsExpanded[i] == nil) {
				SubLibrariesWhoseNodeIsExpandedDifferent = true
				break
			} else if library.SubLibrariesWhoseNodeIsExpanded[i] != nil && libraryOther.SubLibrariesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.SubLibrariesWhoseNodeIsExpanded[i] != libraryOther.SubLibrariesWhoseNodeIsExpanded[i] {
					SubLibrariesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if SubLibrariesWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, library, libraryOther, "SubLibrariesWhoseNodeIsExpanded", libraryOther.SubLibrariesWhoseNodeIsExpanded, library.SubLibrariesWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if library.NbPixPerCharacter != libraryOther.NbPixPerCharacter {
		diffs = append(diffs, library.GongMarshallField(stage, "NbPixPerCharacter"))
	}
	if library.LogoSVGFile != libraryOther.LogoSVGFile {
		diffs = append(diffs, library.GongMarshallField(stage, "LogoSVGFile"))
	}
	if library.IsExpandedTmp != libraryOther.IsExpandedTmp {
		diffs = append(diffs, library.GongMarshallField(stage, "IsExpandedTmp"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (mapobject *MapObject) GongDiff(stage *Stage, mapobjectOther *MapObject) (diffs []string) {
	// insertion point for field diffs
	if mapobject.Name != mapobjectOther.Name {
		diffs = append(diffs, mapobject.GongMarshallField(stage, "Name"))
	}
	if mapobject.ComputedPrefix != mapobjectOther.ComputedPrefix {
		diffs = append(diffs, mapobject.GongMarshallField(stage, "ComputedPrefix"))
	}
	if mapobject.IsExpanded != mapobjectOther.IsExpanded {
		diffs = append(diffs, mapobject.GongMarshallField(stage, "IsExpanded"))
	}
	if mapobject.LayoutDirection != mapobjectOther.LayoutDirection {
		diffs = append(diffs, mapobject.GongMarshallField(stage, "LayoutDirection"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (mapobjectuse *MapObjectUse) GongDiff(stage *Stage, mapobjectuseOther *MapObjectUse) (diffs []string) {
	// insertion point for field diffs
	if mapobjectuse.Name != mapobjectuseOther.Name {
		diffs = append(diffs, mapobjectuse.GongMarshallField(stage, "Name"))
	}
	if (mapobjectuse.Map == nil) != (mapobjectuseOther.Map == nil) {
		diffs = append(diffs, mapobjectuse.GongMarshallField(stage, "Map"))
	} else if mapobjectuse.Map != nil && mapobjectuseOther.Map != nil {
		if mapobjectuse.Map != mapobjectuseOther.Map {
			diffs = append(diffs, mapobjectuse.GongMarshallField(stage, "Map"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (parameter *Parameter) GongDiff(stage *Stage, parameterOther *Parameter) (diffs []string) {
	// insertion point for field diffs
	if parameter.Name != parameterOther.Name {
		diffs = append(diffs, parameter.GongMarshallField(stage, "Name"))
	}
	if parameter.Description != parameterOther.Description {
		diffs = append(diffs, parameter.GongMarshallField(stage, "Description"))
	}
	if parameter.IsResponse != parameterOther.IsResponse {
		diffs = append(diffs, parameter.GongMarshallField(stage, "IsResponse"))
	}
	if parameter.Start != parameterOther.Start {
		diffs = append(diffs, parameter.GongMarshallField(stage, "Start"))
	}
	if parameter.End != parameterOther.End {
		diffs = append(diffs, parameter.GongMarshallField(stage, "End"))
	}
	if parameter.Force != parameterOther.Force {
		diffs = append(diffs, parameter.GongMarshallField(stage, "Force"))
	}
	GroupUseDifferent := false
	if len(parameter.GroupUse) != len(parameterOther.GroupUse) {
		GroupUseDifferent = true
	} else {
		for i := range parameter.GroupUse {
			if (parameter.GroupUse[i] == nil) != (parameterOther.GroupUse[i] == nil) {
				GroupUseDifferent = true
				break
			} else if parameter.GroupUse[i] != nil && parameterOther.GroupUse[i] != nil {
				// this is a pointer comparaison
				if parameter.GroupUse[i] != parameterOther.GroupUse[i] {
					GroupUseDifferent = true
					break
				}
			}
		}
	}
	if GroupUseDifferent {
		ops := Diff(stage, parameter, parameterOther, "GroupUse", parameterOther.GroupUse, parameter.GroupUse)
		diffs = append(diffs, ops)
	}
	DocumentUseDifferent := false
	if len(parameter.DocumentUse) != len(parameterOther.DocumentUse) {
		DocumentUseDifferent = true
	} else {
		for i := range parameter.DocumentUse {
			if (parameter.DocumentUse[i] == nil) != (parameterOther.DocumentUse[i] == nil) {
				DocumentUseDifferent = true
				break
			} else if parameter.DocumentUse[i] != nil && parameterOther.DocumentUse[i] != nil {
				// this is a pointer comparaison
				if parameter.DocumentUse[i] != parameterOther.DocumentUse[i] {
					DocumentUseDifferent = true
					break
				}
			}
		}
	}
	if DocumentUseDifferent {
		ops := Diff(stage, parameter, parameterOther, "DocumentUse", parameterOther.DocumentUse, parameter.DocumentUse)
		diffs = append(diffs, ops)
	}
	GeoObjectUseDifferent := false
	if len(parameter.GeoObjectUse) != len(parameterOther.GeoObjectUse) {
		GeoObjectUseDifferent = true
	} else {
		for i := range parameter.GeoObjectUse {
			if (parameter.GeoObjectUse[i] == nil) != (parameterOther.GeoObjectUse[i] == nil) {
				GeoObjectUseDifferent = true
				break
			} else if parameter.GeoObjectUse[i] != nil && parameterOther.GeoObjectUse[i] != nil {
				// this is a pointer comparaison
				if parameter.GeoObjectUse[i] != parameterOther.GeoObjectUse[i] {
					GeoObjectUseDifferent = true
					break
				}
			}
		}
	}
	if GeoObjectUseDifferent {
		ops := Diff(stage, parameter, parameterOther, "GeoObjectUse", parameterOther.GeoObjectUse, parameter.GeoObjectUse)
		diffs = append(diffs, ops)
	}
	if parameter.Tag != parameterOther.Tag {
		diffs = append(diffs, parameter.GongMarshallField(stage, "Tag"))
	}
	if parameter.ComputedPrefix != parameterOther.ComputedPrefix {
		diffs = append(diffs, parameter.GongMarshallField(stage, "ComputedPrefix"))
	}
	if parameter.IsExpanded != parameterOther.IsExpanded {
		diffs = append(diffs, parameter.GongMarshallField(stage, "IsExpanded"))
	}
	if parameter.LayoutDirection != parameterOther.LayoutDirection {
		diffs = append(diffs, parameter.GongMarshallField(stage, "LayoutDirection"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (parametercategory *ParameterCategory) GongDiff(stage *Stage, parametercategoryOther *ParameterCategory) (diffs []string) {
	// insertion point for field diffs
	if parametercategory.Name != parametercategoryOther.Name {
		diffs = append(diffs, parametercategory.GongMarshallField(stage, "Name"))
	}
	ParameterUseDifferent := false
	if len(parametercategory.ParameterUse) != len(parametercategoryOther.ParameterUse) {
		ParameterUseDifferent = true
	} else {
		for i := range parametercategory.ParameterUse {
			if (parametercategory.ParameterUse[i] == nil) != (parametercategoryOther.ParameterUse[i] == nil) {
				ParameterUseDifferent = true
				break
			} else if parametercategory.ParameterUse[i] != nil && parametercategoryOther.ParameterUse[i] != nil {
				// this is a pointer comparaison
				if parametercategory.ParameterUse[i] != parametercategoryOther.ParameterUse[i] {
					ParameterUseDifferent = true
					break
				}
			}
		}
	}
	if ParameterUseDifferent {
		ops := Diff(stage, parametercategory, parametercategoryOther, "ParameterUse", parametercategoryOther.ParameterUse, parametercategory.ParameterUse)
		diffs = append(diffs, ops)
	}
	if parametercategory.ComputedPrefix != parametercategoryOther.ComputedPrefix {
		diffs = append(diffs, parametercategory.GongMarshallField(stage, "ComputedPrefix"))
	}
	if parametercategory.IsExpanded != parametercategoryOther.IsExpanded {
		diffs = append(diffs, parametercategory.GongMarshallField(stage, "IsExpanded"))
	}
	if parametercategory.LayoutDirection != parametercategoryOther.LayoutDirection {
		diffs = append(diffs, parametercategory.GongMarshallField(stage, "LayoutDirection"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (parametercategoryuse *ParameterCategoryUse) GongDiff(stage *Stage, parametercategoryuseOther *ParameterCategoryUse) (diffs []string) {
	// insertion point for field diffs
	if parametercategoryuse.Name != parametercategoryuseOther.Name {
		diffs = append(diffs, parametercategoryuse.GongMarshallField(stage, "Name"))
	}
	if (parametercategoryuse.ParameterCategory == nil) != (parametercategoryuseOther.ParameterCategory == nil) {
		diffs = append(diffs, parametercategoryuse.GongMarshallField(stage, "ParameterCategory"))
	} else if parametercategoryuse.ParameterCategory != nil && parametercategoryuseOther.ParameterCategory != nil {
		if parametercategoryuse.ParameterCategory != parametercategoryuseOther.ParameterCategory {
			diffs = append(diffs, parametercategoryuse.GongMarshallField(stage, "ParameterCategory"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (parametershape *ParameterShape) GongDiff(stage *Stage, parametershapeOther *ParameterShape) (diffs []string) {
	// insertion point for field diffs
	if parametershape.Name != parametershapeOther.Name {
		diffs = append(diffs, parametershape.GongMarshallField(stage, "Name"))
	}
	if (parametershape.Parameter == nil) != (parametershapeOther.Parameter == nil) {
		diffs = append(diffs, parametershape.GongMarshallField(stage, "Parameter"))
	} else if parametershape.Parameter != nil && parametershapeOther.Parameter != nil {
		if parametershape.Parameter != parametershapeOther.Parameter {
			diffs = append(diffs, parametershape.GongMarshallField(stage, "Parameter"))
		}
	}
	if parametershape.Direction != parametershapeOther.Direction {
		diffs = append(diffs, parametershape.GongMarshallField(stage, "Direction"))
	}
	if parametershape.ShapeIsComputedFromModel != parametershapeOther.ShapeIsComputedFromModel {
		diffs = append(diffs, parametershape.GongMarshallField(stage, "ShapeIsComputedFromModel"))
	}
	if parametershape.X != parametershapeOther.X {
		diffs = append(diffs, parametershape.GongMarshallField(stage, "X"))
	}
	if parametershape.Y != parametershapeOther.Y {
		diffs = append(diffs, parametershape.GongMarshallField(stage, "Y"))
	}
	if parametershape.Width != parametershapeOther.Width {
		diffs = append(diffs, parametershape.GongMarshallField(stage, "Width"))
	}
	if parametershape.Height != parametershapeOther.Height {
		diffs = append(diffs, parametershape.GongMarshallField(stage, "Height"))
	}
	if parametershape.IsHidden != parametershapeOther.IsHidden {
		diffs = append(diffs, parametershape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (parametersaggregate *ParametersAggregate) GongDiff(stage *Stage, parametersaggregateOther *ParametersAggregate) (diffs []string) {
	// insertion point for field diffs
	if parametersaggregate.Name != parametersaggregateOther.Name {
		diffs = append(diffs, parametersaggregate.GongMarshallField(stage, "Name"))
	}
	if parametersaggregate.Tag != parametersaggregateOther.Tag {
		diffs = append(diffs, parametersaggregate.GongMarshallField(stage, "Tag"))
	}
	if parametersaggregate.Description != parametersaggregateOther.Description {
		diffs = append(diffs, parametersaggregate.GongMarshallField(stage, "Description"))
	}
	ParametersDifferent := false
	if len(parametersaggregate.Parameters) != len(parametersaggregateOther.Parameters) {
		ParametersDifferent = true
	} else {
		for i := range parametersaggregate.Parameters {
			if (parametersaggregate.Parameters[i] == nil) != (parametersaggregateOther.Parameters[i] == nil) {
				ParametersDifferent = true
				break
			} else if parametersaggregate.Parameters[i] != nil && parametersaggregateOther.Parameters[i] != nil {
				// this is a pointer comparaison
				if parametersaggregate.Parameters[i] != parametersaggregateOther.Parameters[i] {
					ParametersDifferent = true
					break
				}
			}
		}
	}
	if ParametersDifferent {
		ops := Diff(stage, parametersaggregate, parametersaggregateOther, "Parameters", parametersaggregateOther.Parameters, parametersaggregate.Parameters)
		diffs = append(diffs, ops)
	}
	if parametersaggregate.ComputedPrefix != parametersaggregateOther.ComputedPrefix {
		diffs = append(diffs, parametersaggregate.GongMarshallField(stage, "ComputedPrefix"))
	}
	if parametersaggregate.IsExpanded != parametersaggregateOther.IsExpanded {
		diffs = append(diffs, parametersaggregate.GongMarshallField(stage, "IsExpanded"))
	}
	if parametersaggregate.LayoutDirection != parametersaggregateOther.LayoutDirection {
		diffs = append(diffs, parametersaggregate.GongMarshallField(stage, "LayoutDirection"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (parametersaggregateshape *ParametersAggregateShape) GongDiff(stage *Stage, parametersaggregateshapeOther *ParametersAggregateShape) (diffs []string) {
	// insertion point for field diffs
	if parametersaggregateshape.Name != parametersaggregateshapeOther.Name {
		diffs = append(diffs, parametersaggregateshape.GongMarshallField(stage, "Name"))
	}
	if (parametersaggregateshape.ScenarioParameter == nil) != (parametersaggregateshapeOther.ScenarioParameter == nil) {
		diffs = append(diffs, parametersaggregateshape.GongMarshallField(stage, "ScenarioParameter"))
	} else if parametersaggregateshape.ScenarioParameter != nil && parametersaggregateshapeOther.ScenarioParameter != nil {
		if parametersaggregateshape.ScenarioParameter != parametersaggregateshapeOther.ScenarioParameter {
			diffs = append(diffs, parametersaggregateshape.GongMarshallField(stage, "ScenarioParameter"))
		}
	}
	if parametersaggregateshape.Direction != parametersaggregateshapeOther.Direction {
		diffs = append(diffs, parametersaggregateshape.GongMarshallField(stage, "Direction"))
	}
	if parametersaggregateshape.X != parametersaggregateshapeOther.X {
		diffs = append(diffs, parametersaggregateshape.GongMarshallField(stage, "X"))
	}
	if parametersaggregateshape.Y != parametersaggregateshapeOther.Y {
		diffs = append(diffs, parametersaggregateshape.GongMarshallField(stage, "Y"))
	}
	if parametersaggregateshape.Width != parametersaggregateshapeOther.Width {
		diffs = append(diffs, parametersaggregateshape.GongMarshallField(stage, "Width"))
	}
	if parametersaggregateshape.Height != parametersaggregateshapeOther.Height {
		diffs = append(diffs, parametersaggregateshape.GongMarshallField(stage, "Height"))
	}
	if parametersaggregateshape.IsHidden != parametersaggregateshapeOther.IsHidden {
		diffs = append(diffs, parametersaggregateshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (position *Position) GongDiff(stage *Stage, positionOther *Position) (diffs []string) {
	// insertion point for field diffs
	if position.Name != positionOther.Name {
		diffs = append(diffs, position.GongMarshallField(stage, "Name"))
	}
	if position.Date != positionOther.Date {
		diffs = append(diffs, position.GongMarshallField(stage, "Date"))
	}
	if position.Ordinate != positionOther.Ordinate {
		diffs = append(diffs, position.GongMarshallField(stage, "Ordinate"))
	}
	if position.ComputedPrefix != positionOther.ComputedPrefix {
		diffs = append(diffs, position.GongMarshallField(stage, "ComputedPrefix"))
	}
	if position.IsExpanded != positionOther.IsExpanded {
		diffs = append(diffs, position.GongMarshallField(stage, "IsExpanded"))
	}
	if position.LayoutDirection != positionOther.LayoutDirection {
		diffs = append(diffs, position.GongMarshallField(stage, "LayoutDirection"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (repository *Repository) GongDiff(stage *Stage, repositoryOther *Repository) (diffs []string) {
	// insertion point for field diffs
	if repository.Name != repositoryOther.Name {
		diffs = append(diffs, repository.GongMarshallField(stage, "Name"))
	}
	ParameterUseDifferent := false
	if len(repository.ParameterUse) != len(repositoryOther.ParameterUse) {
		ParameterUseDifferent = true
	} else {
		for i := range repository.ParameterUse {
			if (repository.ParameterUse[i] == nil) != (repositoryOther.ParameterUse[i] == nil) {
				ParameterUseDifferent = true
				break
			} else if repository.ParameterUse[i] != nil && repositoryOther.ParameterUse[i] != nil {
				// this is a pointer comparaison
				if repository.ParameterUse[i] != repositoryOther.ParameterUse[i] {
					ParameterUseDifferent = true
					break
				}
			}
		}
	}
	if ParameterUseDifferent {
		ops := Diff(stage, repository, repositoryOther, "ParameterUse", repositoryOther.ParameterUse, repository.ParameterUse)
		diffs = append(diffs, ops)
	}
	GroupUseDifferent := false
	if len(repository.GroupUse) != len(repositoryOther.GroupUse) {
		GroupUseDifferent = true
	} else {
		for i := range repository.GroupUse {
			if (repository.GroupUse[i] == nil) != (repositoryOther.GroupUse[i] == nil) {
				GroupUseDifferent = true
				break
			} else if repository.GroupUse[i] != nil && repositoryOther.GroupUse[i] != nil {
				// this is a pointer comparaison
				if repository.GroupUse[i] != repositoryOther.GroupUse[i] {
					GroupUseDifferent = true
					break
				}
			}
		}
	}
	if GroupUseDifferent {
		ops := Diff(stage, repository, repositoryOther, "GroupUse", repositoryOther.GroupUse, repository.GroupUse)
		diffs = append(diffs, ops)
	}
	if repository.ComputedPrefix != repositoryOther.ComputedPrefix {
		diffs = append(diffs, repository.GongMarshallField(stage, "ComputedPrefix"))
	}
	if repository.IsExpanded != repositoryOther.IsExpanded {
		diffs = append(diffs, repository.GongMarshallField(stage, "IsExpanded"))
	}
	if repository.LayoutDirection != repositoryOther.LayoutDirection {
		diffs = append(diffs, repository.GongMarshallField(stage, "LayoutDirection"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (scenario *Scenario) GongDiff(stage *Stage, scenarioOther *Scenario) (diffs []string) {
	// insertion point for field diffs
	if scenario.Name != scenarioOther.Name {
		diffs = append(diffs, scenario.GongMarshallField(stage, "Name"))
	}
	if scenario.Description != scenarioOther.Description {
		diffs = append(diffs, scenario.GongMarshallField(stage, "Description"))
	}
	DiagramsDifferent := false
	if len(scenario.Diagrams) != len(scenarioOther.Diagrams) {
		DiagramsDifferent = true
	} else {
		for i := range scenario.Diagrams {
			if (scenario.Diagrams[i] == nil) != (scenarioOther.Diagrams[i] == nil) {
				DiagramsDifferent = true
				break
			} else if scenario.Diagrams[i] != nil && scenarioOther.Diagrams[i] != nil {
				// this is a pointer comparaison
				if scenario.Diagrams[i] != scenarioOther.Diagrams[i] {
					DiagramsDifferent = true
					break
				}
			}
		}
	}
	if DiagramsDifferent {
		ops := Diff(stage, scenario, scenarioOther, "Diagrams", scenarioOther.Diagrams, scenario.Diagrams)
		diffs = append(diffs, ops)
	}
	if scenario.IsDiagramsNodeExpanded != scenarioOther.IsDiagramsNodeExpanded {
		diffs = append(diffs, scenario.GongMarshallField(stage, "IsDiagramsNodeExpanded"))
	}
	ActorStatesDifferent := false
	if len(scenario.ActorStates) != len(scenarioOther.ActorStates) {
		ActorStatesDifferent = true
	} else {
		for i := range scenario.ActorStates {
			if (scenario.ActorStates[i] == nil) != (scenarioOther.ActorStates[i] == nil) {
				ActorStatesDifferent = true
				break
			} else if scenario.ActorStates[i] != nil && scenarioOther.ActorStates[i] != nil {
				// this is a pointer comparaison
				if scenario.ActorStates[i] != scenarioOther.ActorStates[i] {
					ActorStatesDifferent = true
					break
				}
			}
		}
	}
	if ActorStatesDifferent {
		ops := Diff(stage, scenario, scenarioOther, "ActorStates", scenarioOther.ActorStates, scenario.ActorStates)
		diffs = append(diffs, ops)
	}
	if scenario.IsActorStatesNodeExpanded != scenarioOther.IsActorStatesNodeExpanded {
		diffs = append(diffs, scenario.GongMarshallField(stage, "IsActorStatesNodeExpanded"))
	}
	ActorStateTransitionsDifferent := false
	if len(scenario.ActorStateTransitions) != len(scenarioOther.ActorStateTransitions) {
		ActorStateTransitionsDifferent = true
	} else {
		for i := range scenario.ActorStateTransitions {
			if (scenario.ActorStateTransitions[i] == nil) != (scenarioOther.ActorStateTransitions[i] == nil) {
				ActorStateTransitionsDifferent = true
				break
			} else if scenario.ActorStateTransitions[i] != nil && scenarioOther.ActorStateTransitions[i] != nil {
				// this is a pointer comparaison
				if scenario.ActorStateTransitions[i] != scenarioOther.ActorStateTransitions[i] {
					ActorStateTransitionsDifferent = true
					break
				}
			}
		}
	}
	if ActorStateTransitionsDifferent {
		ops := Diff(stage, scenario, scenarioOther, "ActorStateTransitions", scenarioOther.ActorStateTransitions, scenario.ActorStateTransitions)
		diffs = append(diffs, ops)
	}
	if scenario.IsActorStateTransitionsNodeExpanded != scenarioOther.IsActorStateTransitionsNodeExpanded {
		diffs = append(diffs, scenario.GongMarshallField(stage, "IsActorStateTransitionsNodeExpanded"))
	}
	EvolutionDirectionsDifferent := false
	if len(scenario.EvolutionDirections) != len(scenarioOther.EvolutionDirections) {
		EvolutionDirectionsDifferent = true
	} else {
		for i := range scenario.EvolutionDirections {
			if (scenario.EvolutionDirections[i] == nil) != (scenarioOther.EvolutionDirections[i] == nil) {
				EvolutionDirectionsDifferent = true
				break
			} else if scenario.EvolutionDirections[i] != nil && scenarioOther.EvolutionDirections[i] != nil {
				// this is a pointer comparaison
				if scenario.EvolutionDirections[i] != scenarioOther.EvolutionDirections[i] {
					EvolutionDirectionsDifferent = true
					break
				}
			}
		}
	}
	if EvolutionDirectionsDifferent {
		ops := Diff(stage, scenario, scenarioOther, "EvolutionDirections", scenarioOther.EvolutionDirections, scenario.EvolutionDirections)
		diffs = append(diffs, ops)
	}
	if scenario.IsEvolutionDirectionsNodeExpanded != scenarioOther.IsEvolutionDirectionsNodeExpanded {
		diffs = append(diffs, scenario.GongMarshallField(stage, "IsEvolutionDirectionsNodeExpanded"))
	}
	ParametersDifferent := false
	if len(scenario.Parameters) != len(scenarioOther.Parameters) {
		ParametersDifferent = true
	} else {
		for i := range scenario.Parameters {
			if (scenario.Parameters[i] == nil) != (scenarioOther.Parameters[i] == nil) {
				ParametersDifferent = true
				break
			} else if scenario.Parameters[i] != nil && scenarioOther.Parameters[i] != nil {
				// this is a pointer comparaison
				if scenario.Parameters[i] != scenarioOther.Parameters[i] {
					ParametersDifferent = true
					break
				}
			}
		}
	}
	if ParametersDifferent {
		ops := Diff(stage, scenario, scenarioOther, "Parameters", scenarioOther.Parameters, scenario.Parameters)
		diffs = append(diffs, ops)
	}
	if scenario.IsParametersNodeExpanded != scenarioOther.IsParametersNodeExpanded {
		diffs = append(diffs, scenario.GongMarshallField(stage, "IsParametersNodeExpanded"))
	}
	ParametersAggretatesDifferent := false
	if len(scenario.ParametersAggretates) != len(scenarioOther.ParametersAggretates) {
		ParametersAggretatesDifferent = true
	} else {
		for i := range scenario.ParametersAggretates {
			if (scenario.ParametersAggretates[i] == nil) != (scenarioOther.ParametersAggretates[i] == nil) {
				ParametersAggretatesDifferent = true
				break
			} else if scenario.ParametersAggretates[i] != nil && scenarioOther.ParametersAggretates[i] != nil {
				// this is a pointer comparaison
				if scenario.ParametersAggretates[i] != scenarioOther.ParametersAggretates[i] {
					ParametersAggretatesDifferent = true
					break
				}
			}
		}
	}
	if ParametersAggretatesDifferent {
		ops := Diff(stage, scenario, scenarioOther, "ParametersAggretates", scenarioOther.ParametersAggretates, scenario.ParametersAggretates)
		diffs = append(diffs, ops)
	}
	if scenario.IsParametersAggretatesNodeExpanded != scenarioOther.IsParametersAggretatesNodeExpanded {
		diffs = append(diffs, scenario.GongMarshallField(stage, "IsParametersAggretatesNodeExpanded"))
	}
	if scenario.ComputedPrefix != scenarioOther.ComputedPrefix {
		diffs = append(diffs, scenario.GongMarshallField(stage, "ComputedPrefix"))
	}
	if scenario.IsExpanded != scenarioOther.IsExpanded {
		diffs = append(diffs, scenario.GongMarshallField(stage, "IsExpanded"))
	}
	if scenario.LayoutDirection != scenarioOther.LayoutDirection {
		diffs = append(diffs, scenario.GongMarshallField(stage, "LayoutDirection"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (user *User) GongDiff(stage *Stage, userOther *User) (diffs []string) {
	// insertion point for field diffs
	if user.Name != userOther.Name {
		diffs = append(diffs, user.GongMarshallField(stage, "Name"))
	}
	if user.ComputedPrefix != userOther.ComputedPrefix {
		diffs = append(diffs, user.GongMarshallField(stage, "ComputedPrefix"))
	}
	if user.IsExpanded != userOther.IsExpanded {
		diffs = append(diffs, user.GongMarshallField(stage, "IsExpanded"))
	}
	if user.LayoutDirection != userOther.LayoutDirection {
		diffs = append(diffs, user.GongMarshallField(stage, "LayoutDirection"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (useruse *UserUse) GongDiff(stage *Stage, useruseOther *UserUse) (diffs []string) {
	// insertion point for field diffs
	if useruse.Name != useruseOther.Name {
		diffs = append(diffs, useruse.GongMarshallField(stage, "Name"))
	}
	if (useruse.User == nil) != (useruseOther.User == nil) {
		diffs = append(diffs, useruse.GongMarshallField(stage, "User"))
	} else if useruse.User != nil && useruseOther.User != nil {
		if useruse.User != useruseOther.User {
			diffs = append(diffs, useruse.GongMarshallField(stage, "User"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (workspace *Workspace) GongDiff(stage *Stage, workspaceOther *Workspace) (diffs []string) {
	// insertion point for field diffs
	if workspace.Name != workspaceOther.Name {
		diffs = append(diffs, workspace.GongMarshallField(stage, "Name"))
	}
	if (workspace.SelectedDiagram == nil) != (workspaceOther.SelectedDiagram == nil) {
		diffs = append(diffs, workspace.GongMarshallField(stage, "SelectedDiagram"))
	} else if workspace.SelectedDiagram != nil && workspaceOther.SelectedDiagram != nil {
		if workspace.SelectedDiagram != workspaceOther.SelectedDiagram {
			diffs = append(diffs, workspace.GongMarshallField(stage, "SelectedDiagram"))
		}
	}
	if (workspace.Default_EvolutionDirectionShape == nil) != (workspaceOther.Default_EvolutionDirectionShape == nil) {
		diffs = append(diffs, workspace.GongMarshallField(stage, "Default_EvolutionDirectionShape"))
	} else if workspace.Default_EvolutionDirectionShape != nil && workspaceOther.Default_EvolutionDirectionShape != nil {
		if workspace.Default_EvolutionDirectionShape != workspaceOther.Default_EvolutionDirectionShape {
			diffs = append(diffs, workspace.GongMarshallField(stage, "Default_EvolutionDirectionShape"))
		}
	}
	if (workspace.Default_ParameterShape == nil) != (workspaceOther.Default_ParameterShape == nil) {
		diffs = append(diffs, workspace.GongMarshallField(stage, "Default_ParameterShape"))
	} else if workspace.Default_ParameterShape != nil && workspaceOther.Default_ParameterShape != nil {
		if workspace.Default_ParameterShape != workspaceOther.Default_ParameterShape {
			diffs = append(diffs, workspace.GongMarshallField(stage, "Default_ParameterShape"))
		}
	}
	if (workspace.Default_ScenarioParameterShape == nil) != (workspaceOther.Default_ScenarioParameterShape == nil) {
		diffs = append(diffs, workspace.GongMarshallField(stage, "Default_ScenarioParameterShape"))
	} else if workspace.Default_ScenarioParameterShape != nil && workspaceOther.Default_ScenarioParameterShape != nil {
		if workspace.Default_ScenarioParameterShape != workspaceOther.Default_ScenarioParameterShape {
			diffs = append(diffs, workspace.GongMarshallField(stage, "Default_ScenarioParameterShape"))
		}
	}
	if (workspace.Default_ActorStateShape == nil) != (workspaceOther.Default_ActorStateShape == nil) {
		diffs = append(diffs, workspace.GongMarshallField(stage, "Default_ActorStateShape"))
	} else if workspace.Default_ActorStateShape != nil && workspaceOther.Default_ActorStateShape != nil {
		if workspace.Default_ActorStateShape != workspaceOther.Default_ActorStateShape {
			diffs = append(diffs, workspace.GongMarshallField(stage, "Default_ActorStateShape"))
		}
	}
	if (workspace.Default_ActorStateTransitionShape == nil) != (workspaceOther.Default_ActorStateTransitionShape == nil) {
		diffs = append(diffs, workspace.GongMarshallField(stage, "Default_ActorStateTransitionShape"))
	} else if workspace.Default_ActorStateTransitionShape != nil && workspaceOther.Default_ActorStateTransitionShape != nil {
		if workspace.Default_ActorStateTransitionShape != workspaceOther.Default_ActorStateTransitionShape {
			diffs = append(diffs, workspace.GongMarshallField(stage, "Default_ActorStateTransitionShape"))
		}
	}
	if workspace.ComputedPrefix != workspaceOther.ComputedPrefix {
		diffs = append(diffs, workspace.GongMarshallField(stage, "ComputedPrefix"))
	}
	if workspace.IsExpanded != workspaceOther.IsExpanded {
		diffs = append(diffs, workspace.GongMarshallField(stage, "IsExpanded"))
	}
	if workspace.LayoutDirection != workspaceOther.LayoutDirection {
		diffs = append(diffs, workspace.GongMarshallField(stage, "LayoutDirection"))
	}

	return
}

// Diff returns the sequence of operations to transform oldSlice into newSlice.
// It requires type T to be comparable (e.g., pointers, ints, strings).
func Diff[T1, T2 PointerToGongstruct](stage *Stage, a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
	m, n := len(oldSlice), len(newSlice)

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if oldSlice[i] == newSlice[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if oldSlice[i-1] == newSlice[j-1] {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Create a temporary view of what's left after deletions for tracking matches
	var currentLCS []T2
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			currentLCS = append(currentLCS, oldSlice[k])
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k, targetVal := range newSlice {
		if lcsIdx < len(currentLCS) && currentLCS[lcsIdx] == targetVal {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
