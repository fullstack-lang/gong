// generated code - do not edit
package models

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (actorstate *ActorState) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ActorStates[actorstate]
	return ok
}

func (actorstateshape *ActorStateShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ActorStateShapes[actorstateshape]
	return ok
}

func (actorstatetransition *ActorStateTransition) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ActorStateTransitions[actorstatetransition]
	return ok
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ActorStateTransitionShapes[actorstatetransitionshape]
	return ok
}

func (analysis *Analysis) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Analysiss[analysis]
	return ok
}

func (controlpointshape *ControlPointShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ControlPointShapes[controlpointshape]
	return ok
}

func (diagram *Diagram) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Diagrams[diagram]
	return ok
}

func (document *Document) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Documents[document]
	return ok
}

func (documentuse *DocumentUse) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DocumentUses[documentuse]
	return ok
}

func (evolutiondirection *EvolutionDirection) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EvolutionDirections[evolutiondirection]
	return ok
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EvolutionDirectionShapes[evolutiondirectionshape]
	return ok
}

func (foo *Foo) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Foos[foo]
	return ok
}

func (geoobject *GeoObject) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GeoObjects[geoobject]
	return ok
}

func (geoobjectuse *GeoObjectUse) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GeoObjectUses[geoobjectuse]
	return ok
}

func (group *Group) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Groups[group]
	return ok
}

func (groupuse *GroupUse) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GroupUses[groupuse]
	return ok
}

func (library *Library) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Librarys[library]
	return ok
}

func (mapobject *MapObject) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MapObjects[mapobject]
	return ok
}

func (mapobjectuse *MapObjectUse) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MapObjectUses[mapobjectuse]
	return ok
}

func (parameter *Parameter) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Parameters[parameter]
	return ok
}

func (parametercategory *ParameterCategory) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ParameterCategorys[parametercategory]
	return ok
}

func (parametercategoryuse *ParameterCategoryUse) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ParameterCategoryUses[parametercategoryuse]
	return ok
}

func (parametershape *ParameterShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ParameterShapes[parametershape]
	return ok
}

func (parametersaggregate *ParametersAggregate) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ParametersAggregates[parametersaggregate]
	return ok
}

func (parametersaggregateshape *ParametersAggregateShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ParametersAggregateShapes[parametersaggregateshape]
	return ok
}

func (position *Position) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Positions[position]
	return ok
}

func (repository *Repository) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Repositorys[repository]
	return ok
}

func (scenario *Scenario) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Scenarios[scenario]
	return ok
}

func (user *User) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Users[user]
	return ok
}

func (useruse *UserUse) GongIsStaged(stage *Stage) bool {
	_, ok := stage.UserUses[useruse]
	return ok
}

func (workspace *Workspace) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Workspaces[workspace]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (actorstate *ActorState) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(actorstate) {
		return
	}

	actorstate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (actorstateshape *ActorStateShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(actorstateshape) {
		return
	}

	actorstateshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if actorstateshape.ActorState != nil {
		stage.StageBranch(actorstateshape.ActorState)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (actorstatetransition *ActorStateTransition) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(actorstatetransition) {
		return
	}

	actorstatetransition.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if actorstatetransition.StartState != nil {
		stage.StageBranch(actorstatetransition.StartState)
	}
	if actorstatetransition.EndState != nil {
		stage.StageBranch(actorstatetransition.EndState)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parameter := range actorstatetransition.Justifications {
		stage.StageBranch(_parameter)
	}

}

func (actorstatetransitionshape *ActorStateTransitionShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(actorstatetransitionshape) {
		return
	}

	actorstatetransitionshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if actorstatetransitionshape.ActorStateTransition != nil {
		stage.StageBranch(actorstatetransitionshape.ActorStateTransition)
	}
	if actorstatetransitionshape.Start != nil {
		stage.StageBranch(actorstatetransitionshape.Start)
	}
	if actorstatetransitionshape.End != nil {
		stage.StageBranch(actorstatetransitionshape.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range actorstatetransitionshape.ControlPointShapes {
		stage.StageBranch(_controlpointshape)
	}

}

func (analysis *Analysis) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(analysis) {
		return
	}

	analysis.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _scenario := range analysis.Scenarios {
		stage.StageBranch(_scenario)
	}
	for _, _groupuse := range analysis.GroupUse {
		stage.StageBranch(_groupuse)
	}
	for _, _geoobjectuse := range analysis.GeoObjectUse {
		stage.StageBranch(_geoobjectuse)
	}
	for _, _mapobjectuse := range analysis.MapUse {
		stage.StageBranch(_mapobjectuse)
	}

}

func (controlpointshape *ControlPointShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(controlpointshape) {
		return
	}

	controlpointshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (diagram *Diagram) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(diagram) {
		return
	}

	diagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _evolutiondirectionshape := range diagram.EvolutionDirectionShapes {
		stage.StageBranch(_evolutiondirectionshape)
	}
	for _, _evolutiondirection := range diagram.EvolutionDirectionsWhoseNodeIsExpanded {
		stage.StageBranch(_evolutiondirection)
	}
	for _, _actorstateshape := range diagram.ActorStateShapes {
		stage.StageBranch(_actorstateshape)
	}
	for _, _actorstate := range diagram.ActorStatesWhoseNodeIsExpanded {
		stage.StageBranch(_actorstate)
	}
	for _, _parametershape := range diagram.ParameterShapes {
		stage.StageBranch(_parametershape)
	}
	for _, _parameter := range diagram.ParametersWhoseNodeIsExpanded {
		stage.StageBranch(_parameter)
	}
	for _, _parametersaggregateshape := range diagram.ScenarioParameterShapes {
		stage.StageBranch(_parametersaggregateshape)
	}
	for _, _parametersaggregate := range diagram.ParametersAggregatesWhoseNodeIsExpanded {
		stage.StageBranch(_parametersaggregate)
	}
	for _, _actorstatetransitionshape := range diagram.ActorStateTransitionShapes {
		stage.StageBranch(_actorstatetransitionshape)
	}
	for _, _actorstatetransition := range diagram.ActorStateTransitionsWhoseNodeIsExpanded {
		stage.StageBranch(_actorstatetransition)
	}

}

func (document *Document) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(document) {
		return
	}

	document.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _geoobjectuse := range document.GeoObjectUse {
		stage.StageBranch(_geoobjectuse)
	}

}

func (documentuse *DocumentUse) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(documentuse) {
		return
	}

	documentuse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if documentuse.Document != nil {
		stage.StageBranch(documentuse.Document)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (evolutiondirection *EvolutionDirection) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(evolutiondirection) {
		return
	}

	evolutiondirection.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (evolutiondirectionshape *EvolutionDirectionShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(evolutiondirectionshape) {
		return
	}

	evolutiondirectionshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if evolutiondirectionshape.EvolutionDirection != nil {
		stage.StageBranch(evolutiondirectionshape.EvolutionDirection)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (foo *Foo) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(foo) {
		return
	}

	foo.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (geoobject *GeoObject) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(geoobject) {
		return
	}

	geoobject.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (geoobjectuse *GeoObjectUse) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(geoobjectuse) {
		return
	}

	geoobjectuse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if geoobjectuse.GeoObject != nil {
		stage.StageBranch(geoobjectuse.GeoObject)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (group *Group) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(group) {
		return
	}

	group.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _useruse := range group.UserUse {
		stage.StageBranch(_useruse)
	}

}

func (groupuse *GroupUse) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(groupuse) {
		return
	}

	groupuse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if groupuse.Group != nil {
		stage.StageBranch(groupuse.Group)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (library *Library) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(library) {
		return
	}

	library.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _analysis := range library.Analyses {
		stage.StageBranch(_analysis)
	}
	for _, _library := range library.SubLibraries {
		stage.StageBranch(_library)
	}
	for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
		stage.StageBranch(_library)
	}

}

func (mapobject *MapObject) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(mapobject) {
		return
	}

	mapobject.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (mapobjectuse *MapObjectUse) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(mapobjectuse) {
		return
	}

	mapobjectuse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if mapobjectuse.Map != nil {
		stage.StageBranch(mapobjectuse.Map)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (parameter *Parameter) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(parameter) {
		return
	}

	parameter.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _groupuse := range parameter.GroupUse {
		stage.StageBranch(_groupuse)
	}
	for _, _documentuse := range parameter.DocumentUse {
		stage.StageBranch(_documentuse)
	}
	for _, _geoobjectuse := range parameter.GeoObjectUse {
		stage.StageBranch(_geoobjectuse)
	}

}

func (parametercategory *ParameterCategory) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(parametercategory) {
		return
	}

	parametercategory.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parametershape := range parametercategory.ParameterUse {
		stage.StageBranch(_parametershape)
	}

}

func (parametercategoryuse *ParameterCategoryUse) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(parametercategoryuse) {
		return
	}

	parametercategoryuse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if parametercategoryuse.ParameterCategory != nil {
		stage.StageBranch(parametercategoryuse.ParameterCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (parametershape *ParameterShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(parametershape) {
		return
	}

	parametershape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if parametershape.Parameter != nil {
		stage.StageBranch(parametershape.Parameter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (parametersaggregate *ParametersAggregate) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(parametersaggregate) {
		return
	}

	parametersaggregate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parameter := range parametersaggregate.Parameters {
		stage.StageBranch(_parameter)
	}

}

func (parametersaggregateshape *ParametersAggregateShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(parametersaggregateshape) {
		return
	}

	parametersaggregateshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if parametersaggregateshape.ScenarioParameter != nil {
		stage.StageBranch(parametersaggregateshape.ScenarioParameter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (position *Position) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(position) {
		return
	}

	position.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (repository *Repository) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(repository) {
		return
	}

	repository.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parametershape := range repository.ParameterUse {
		stage.StageBranch(_parametershape)
	}
	for _, _groupuse := range repository.GroupUse {
		stage.StageBranch(_groupuse)
	}

}

func (scenario *Scenario) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(scenario) {
		return
	}

	scenario.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _diagram := range scenario.Diagrams {
		stage.StageBranch(_diagram)
	}
	for _, _actorstate := range scenario.ActorStates {
		stage.StageBranch(_actorstate)
	}
	for _, _actorstatetransition := range scenario.ActorStateTransitions {
		stage.StageBranch(_actorstatetransition)
	}
	for _, _evolutiondirection := range scenario.EvolutionDirections {
		stage.StageBranch(_evolutiondirection)
	}
	for _, _parameter := range scenario.Parameters {
		stage.StageBranch(_parameter)
	}
	for _, _parametersaggregate := range scenario.ParametersAggretates {
		stage.StageBranch(_parametersaggregate)
	}

}

func (user *User) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(user) {
		return
	}

	user.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (useruse *UserUse) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(useruse) {
		return
	}

	useruse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if useruse.User != nil {
		stage.StageBranch(useruse.User)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (workspace *Workspace) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(workspace) {
		return
	}

	workspace.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if workspace.SelectedDiagram != nil {
		stage.StageBranch(workspace.SelectedDiagram)
	}
	if workspace.Default_EvolutionDirectionShape != nil {
		stage.StageBranch(workspace.Default_EvolutionDirectionShape)
	}
	if workspace.Default_ParameterShape != nil {
		stage.StageBranch(workspace.Default_ParameterShape)
	}
	if workspace.Default_ScenarioParameterShape != nil {
		stage.StageBranch(workspace.Default_ScenarioParameterShape)
	}
	if workspace.Default_ActorStateShape != nil {
		stage.StageBranch(workspace.Default_ActorStateShape)
	}
	if workspace.Default_ActorStateTransitionShape != nil {
		stage.StageBranch(workspace.Default_ActorStateTransitionShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *ActorState:
		toT := GongCopyBranchActorState(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ActorStateShape:
		toT := GongCopyBranchActorStateShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ActorStateTransition:
		toT := GongCopyBranchActorStateTransition(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ActorStateTransitionShape:
		toT := GongCopyBranchActorStateTransitionShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Analysis:
		toT := GongCopyBranchAnalysis(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ControlPointShape:
		toT := GongCopyBranchControlPointShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Diagram:
		toT := GongCopyBranchDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Document:
		toT := GongCopyBranchDocument(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DocumentUse:
		toT := GongCopyBranchDocumentUse(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EvolutionDirection:
		toT := GongCopyBranchEvolutionDirection(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EvolutionDirectionShape:
		toT := GongCopyBranchEvolutionDirectionShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Foo:
		toT := GongCopyBranchFoo(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GeoObject:
		toT := GongCopyBranchGeoObject(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GeoObjectUse:
		toT := GongCopyBranchGeoObjectUse(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Group:
		toT := GongCopyBranchGroup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GroupUse:
		toT := GongCopyBranchGroupUse(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := GongCopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MapObject:
		toT := GongCopyBranchMapObject(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MapObjectUse:
		toT := GongCopyBranchMapObjectUse(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Parameter:
		toT := GongCopyBranchParameter(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ParameterCategory:
		toT := GongCopyBranchParameterCategory(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ParameterCategoryUse:
		toT := GongCopyBranchParameterCategoryUse(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ParameterShape:
		toT := GongCopyBranchParameterShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ParametersAggregate:
		toT := GongCopyBranchParametersAggregate(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ParametersAggregateShape:
		toT := GongCopyBranchParametersAggregateShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Position:
		toT := GongCopyBranchPosition(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Repository:
		toT := GongCopyBranchRepository(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Scenario:
		toT := GongCopyBranchScenario(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *User:
		toT := GongCopyBranchUser(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *UserUse:
		toT := GongCopyBranchUserUse(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Workspace:
		toT := GongCopyBranchWorkspace(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchActorState(mapOrigCopy map[any]any, actorstateFrom *ActorState) (actorstateTo *ActorState) {
	var alreadyCopied bool
	actorstateTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, actorstateFrom)
	if alreadyCopied {
		return
	}
	actorstateFrom.GongCopyBasicFields(actorstateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchActorStateShape(mapOrigCopy map[any]any, actorstateshapeFrom *ActorStateShape) (actorstateshapeTo *ActorStateShape) {
	var alreadyCopied bool
	actorstateshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, actorstateshapeFrom)
	if alreadyCopied {
		return
	}
	actorstateshapeFrom.GongCopyBasicFields(actorstateshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if actorstateshapeFrom.ActorState != nil {
		actorstateshapeTo.ActorState = actorstateshapeFrom.ActorState
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchActorStateTransition(mapOrigCopy map[any]any, actorstatetransitionFrom *ActorStateTransition) (actorstatetransitionTo *ActorStateTransition) {
	var alreadyCopied bool
	actorstatetransitionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, actorstatetransitionFrom)
	if alreadyCopied {
		return
	}
	actorstatetransitionFrom.GongCopyBasicFields(actorstatetransitionTo)

	//insertion point for the staging of instances referenced by pointers
	if actorstatetransitionFrom.StartState != nil {
		actorstatetransitionTo.StartState = GongCopyBranchActorState(mapOrigCopy, actorstatetransitionFrom.StartState)
	}
	if actorstatetransitionFrom.EndState != nil {
		actorstatetransitionTo.EndState = GongCopyBranchActorState(mapOrigCopy, actorstatetransitionFrom.EndState)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parameter := range actorstatetransitionFrom.Justifications {
		actorstatetransitionTo.Justifications = append(actorstatetransitionTo.Justifications, GongCopyBranchParameter(mapOrigCopy, _parameter))
	}

	return
}

func GongCopyBranchActorStateTransitionShape(mapOrigCopy map[any]any, actorstatetransitionshapeFrom *ActorStateTransitionShape) (actorstatetransitionshapeTo *ActorStateTransitionShape) {
	var alreadyCopied bool
	actorstatetransitionshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, actorstatetransitionshapeFrom)
	if alreadyCopied {
		return
	}
	actorstatetransitionshapeFrom.GongCopyBasicFields(actorstatetransitionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if actorstatetransitionshapeFrom.ActorStateTransition != nil {
		actorstatetransitionshapeTo.ActorStateTransition = actorstatetransitionshapeFrom.ActorStateTransition
	}
	if actorstatetransitionshapeFrom.Start != nil {
		actorstatetransitionshapeTo.Start = GongCopyBranchActorStateShape(mapOrigCopy, actorstatetransitionshapeFrom.Start)
	}
	if actorstatetransitionshapeFrom.End != nil {
		actorstatetransitionshapeTo.End = GongCopyBranchActorStateShape(mapOrigCopy, actorstatetransitionshapeFrom.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range actorstatetransitionshapeFrom.ControlPointShapes {
		actorstatetransitionshapeTo.ControlPointShapes = append(actorstatetransitionshapeTo.ControlPointShapes, GongCopyBranchControlPointShape(mapOrigCopy, _controlpointshape))
	}

	return
}

func GongCopyBranchAnalysis(mapOrigCopy map[any]any, analysisFrom *Analysis) (analysisTo *Analysis) {
	var alreadyCopied bool
	analysisTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, analysisFrom)
	if alreadyCopied {
		return
	}
	analysisFrom.GongCopyBasicFields(analysisTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _scenario := range analysisFrom.Scenarios {
		analysisTo.Scenarios = append(analysisTo.Scenarios, GongCopyBranchScenario(mapOrigCopy, _scenario))
	}
	for _, _groupuse := range analysisFrom.GroupUse {
		analysisTo.GroupUse = append(analysisTo.GroupUse, GongCopyBranchGroupUse(mapOrigCopy, _groupuse))
	}
	for _, _geoobjectuse := range analysisFrom.GeoObjectUse {
		analysisTo.GeoObjectUse = append(analysisTo.GeoObjectUse, GongCopyBranchGeoObjectUse(mapOrigCopy, _geoobjectuse))
	}
	for _, _mapobjectuse := range analysisFrom.MapUse {
		analysisTo.MapUse = append(analysisTo.MapUse, GongCopyBranchMapObjectUse(mapOrigCopy, _mapobjectuse))
	}

	return
}

func GongCopyBranchControlPointShape(mapOrigCopy map[any]any, controlpointshapeFrom *ControlPointShape) (controlpointshapeTo *ControlPointShape) {
	var alreadyCopied bool
	controlpointshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, controlpointshapeFrom)
	if alreadyCopied {
		return
	}
	controlpointshapeFrom.GongCopyBasicFields(controlpointshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDiagram(mapOrigCopy map[any]any, diagramFrom *Diagram) (diagramTo *Diagram) {
	var alreadyCopied bool
	diagramTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, diagramFrom)
	if alreadyCopied {
		return
	}
	diagramFrom.GongCopyBasicFields(diagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _evolutiondirectionshape := range diagramFrom.EvolutionDirectionShapes {
		diagramTo.EvolutionDirectionShapes = append(diagramTo.EvolutionDirectionShapes, GongCopyBranchEvolutionDirectionShape(mapOrigCopy, _evolutiondirectionshape))
	}
	for _, _evolutiondirection := range diagramFrom.EvolutionDirectionsWhoseNodeIsExpanded {
		diagramTo.EvolutionDirectionsWhoseNodeIsExpanded = append(diagramTo.EvolutionDirectionsWhoseNodeIsExpanded, GongCopyBranchEvolutionDirection(mapOrigCopy, _evolutiondirection))
	}
	for _, _actorstateshape := range diagramFrom.ActorStateShapes {
		diagramTo.ActorStateShapes = append(diagramTo.ActorStateShapes, GongCopyBranchActorStateShape(mapOrigCopy, _actorstateshape))
	}
	for _, _actorstate := range diagramFrom.ActorStatesWhoseNodeIsExpanded {
		diagramTo.ActorStatesWhoseNodeIsExpanded = append(diagramTo.ActorStatesWhoseNodeIsExpanded, GongCopyBranchActorState(mapOrigCopy, _actorstate))
	}
	for _, _parametershape := range diagramFrom.ParameterShapes {
		diagramTo.ParameterShapes = append(diagramTo.ParameterShapes, GongCopyBranchParameterShape(mapOrigCopy, _parametershape))
	}
	for _, _parameter := range diagramFrom.ParametersWhoseNodeIsExpanded {
		diagramTo.ParametersWhoseNodeIsExpanded = append(diagramTo.ParametersWhoseNodeIsExpanded, GongCopyBranchParameter(mapOrigCopy, _parameter))
	}
	for _, _parametersaggregateshape := range diagramFrom.ScenarioParameterShapes {
		diagramTo.ScenarioParameterShapes = append(diagramTo.ScenarioParameterShapes, GongCopyBranchParametersAggregateShape(mapOrigCopy, _parametersaggregateshape))
	}
	for _, _parametersaggregate := range diagramFrom.ParametersAggregatesWhoseNodeIsExpanded {
		diagramTo.ParametersAggregatesWhoseNodeIsExpanded = append(diagramTo.ParametersAggregatesWhoseNodeIsExpanded, GongCopyBranchParametersAggregate(mapOrigCopy, _parametersaggregate))
	}
	for _, _actorstatetransitionshape := range diagramFrom.ActorStateTransitionShapes {
		diagramTo.ActorStateTransitionShapes = append(diagramTo.ActorStateTransitionShapes, GongCopyBranchActorStateTransitionShape(mapOrigCopy, _actorstatetransitionshape))
	}
	for _, _actorstatetransition := range diagramFrom.ActorStateTransitionsWhoseNodeIsExpanded {
		diagramTo.ActorStateTransitionsWhoseNodeIsExpanded = append(diagramTo.ActorStateTransitionsWhoseNodeIsExpanded, GongCopyBranchActorStateTransition(mapOrigCopy, _actorstatetransition))
	}

	return
}

func GongCopyBranchDocument(mapOrigCopy map[any]any, documentFrom *Document) (documentTo *Document) {
	var alreadyCopied bool
	documentTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, documentFrom)
	if alreadyCopied {
		return
	}
	documentFrom.GongCopyBasicFields(documentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _geoobjectuse := range documentFrom.GeoObjectUse {
		documentTo.GeoObjectUse = append(documentTo.GeoObjectUse, GongCopyBranchGeoObjectUse(mapOrigCopy, _geoobjectuse))
	}

	return
}

func GongCopyBranchDocumentUse(mapOrigCopy map[any]any, documentuseFrom *DocumentUse) (documentuseTo *DocumentUse) {
	var alreadyCopied bool
	documentuseTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, documentuseFrom)
	if alreadyCopied {
		return
	}
	documentuseFrom.GongCopyBasicFields(documentuseTo)

	//insertion point for the staging of instances referenced by pointers
	if documentuseFrom.Document != nil {
		documentuseTo.Document = GongCopyBranchDocument(mapOrigCopy, documentuseFrom.Document)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEvolutionDirection(mapOrigCopy map[any]any, evolutiondirectionFrom *EvolutionDirection) (evolutiondirectionTo *EvolutionDirection) {
	var alreadyCopied bool
	evolutiondirectionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, evolutiondirectionFrom)
	if alreadyCopied {
		return
	}
	evolutiondirectionFrom.GongCopyBasicFields(evolutiondirectionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEvolutionDirectionShape(mapOrigCopy map[any]any, evolutiondirectionshapeFrom *EvolutionDirectionShape) (evolutiondirectionshapeTo *EvolutionDirectionShape) {
	var alreadyCopied bool
	evolutiondirectionshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, evolutiondirectionshapeFrom)
	if alreadyCopied {
		return
	}
	evolutiondirectionshapeFrom.GongCopyBasicFields(evolutiondirectionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if evolutiondirectionshapeFrom.EvolutionDirection != nil {
		evolutiondirectionshapeTo.EvolutionDirection = evolutiondirectionshapeFrom.EvolutionDirection
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFoo(mapOrigCopy map[any]any, fooFrom *Foo) (fooTo *Foo) {
	var alreadyCopied bool
	fooTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, fooFrom)
	if alreadyCopied {
		return
	}
	fooFrom.GongCopyBasicFields(fooTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGeoObject(mapOrigCopy map[any]any, geoobjectFrom *GeoObject) (geoobjectTo *GeoObject) {
	var alreadyCopied bool
	geoobjectTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, geoobjectFrom)
	if alreadyCopied {
		return
	}
	geoobjectFrom.GongCopyBasicFields(geoobjectTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGeoObjectUse(mapOrigCopy map[any]any, geoobjectuseFrom *GeoObjectUse) (geoobjectuseTo *GeoObjectUse) {
	var alreadyCopied bool
	geoobjectuseTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, geoobjectuseFrom)
	if alreadyCopied {
		return
	}
	geoobjectuseFrom.GongCopyBasicFields(geoobjectuseTo)

	//insertion point for the staging of instances referenced by pointers
	if geoobjectuseFrom.GeoObject != nil {
		geoobjectuseTo.GeoObject = GongCopyBranchGeoObject(mapOrigCopy, geoobjectuseFrom.GeoObject)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGroup(mapOrigCopy map[any]any, groupFrom *Group) (groupTo *Group) {
	var alreadyCopied bool
	groupTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, groupFrom)
	if alreadyCopied {
		return
	}
	groupFrom.GongCopyBasicFields(groupTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _useruse := range groupFrom.UserUse {
		groupTo.UserUse = append(groupTo.UserUse, GongCopyBranchUserUse(mapOrigCopy, _useruse))
	}

	return
}

func GongCopyBranchGroupUse(mapOrigCopy map[any]any, groupuseFrom *GroupUse) (groupuseTo *GroupUse) {
	var alreadyCopied bool
	groupuseTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, groupuseFrom)
	if alreadyCopied {
		return
	}
	groupuseFrom.GongCopyBasicFields(groupuseTo)

	//insertion point for the staging of instances referenced by pointers
	if groupuseFrom.Group != nil {
		groupuseTo.Group = GongCopyBranchGroup(mapOrigCopy, groupuseFrom.Group)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLibrary(mapOrigCopy map[any]any, libraryFrom *Library) (libraryTo *Library) {
	var alreadyCopied bool
	libraryTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, libraryFrom)
	if alreadyCopied {
		return
	}
	libraryFrom.GongCopyBasicFields(libraryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _analysis := range libraryFrom.Analyses {
		libraryTo.Analyses = append(libraryTo.Analyses, GongCopyBranchAnalysis(mapOrigCopy, _analysis))
	}
	for _, _library := range libraryFrom.SubLibraries {
		libraryTo.SubLibraries = append(libraryTo.SubLibraries, GongCopyBranchLibrary(mapOrigCopy, _library))
	}
	for _, _library := range libraryFrom.SubLibrariesWhoseNodeIsExpanded {
		libraryTo.SubLibrariesWhoseNodeIsExpanded = append(libraryTo.SubLibrariesWhoseNodeIsExpanded, GongCopyBranchLibrary(mapOrigCopy, _library))
	}

	return
}

func GongCopyBranchMapObject(mapOrigCopy map[any]any, mapobjectFrom *MapObject) (mapobjectTo *MapObject) {
	var alreadyCopied bool
	mapobjectTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, mapobjectFrom)
	if alreadyCopied {
		return
	}
	mapobjectFrom.GongCopyBasicFields(mapobjectTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMapObjectUse(mapOrigCopy map[any]any, mapobjectuseFrom *MapObjectUse) (mapobjectuseTo *MapObjectUse) {
	var alreadyCopied bool
	mapobjectuseTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, mapobjectuseFrom)
	if alreadyCopied {
		return
	}
	mapobjectuseFrom.GongCopyBasicFields(mapobjectuseTo)

	//insertion point for the staging of instances referenced by pointers
	if mapobjectuseFrom.Map != nil {
		mapobjectuseTo.Map = GongCopyBranchMapObject(mapOrigCopy, mapobjectuseFrom.Map)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchParameter(mapOrigCopy map[any]any, parameterFrom *Parameter) (parameterTo *Parameter) {
	var alreadyCopied bool
	parameterTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, parameterFrom)
	if alreadyCopied {
		return
	}
	parameterFrom.GongCopyBasicFields(parameterTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _groupuse := range parameterFrom.GroupUse {
		parameterTo.GroupUse = append(parameterTo.GroupUse, GongCopyBranchGroupUse(mapOrigCopy, _groupuse))
	}
	for _, _documentuse := range parameterFrom.DocumentUse {
		parameterTo.DocumentUse = append(parameterTo.DocumentUse, GongCopyBranchDocumentUse(mapOrigCopy, _documentuse))
	}
	for _, _geoobjectuse := range parameterFrom.GeoObjectUse {
		parameterTo.GeoObjectUse = append(parameterTo.GeoObjectUse, GongCopyBranchGeoObjectUse(mapOrigCopy, _geoobjectuse))
	}

	return
}

func GongCopyBranchParameterCategory(mapOrigCopy map[any]any, parametercategoryFrom *ParameterCategory) (parametercategoryTo *ParameterCategory) {
	var alreadyCopied bool
	parametercategoryTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, parametercategoryFrom)
	if alreadyCopied {
		return
	}
	parametercategoryFrom.GongCopyBasicFields(parametercategoryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parametershape := range parametercategoryFrom.ParameterUse {
		parametercategoryTo.ParameterUse = append(parametercategoryTo.ParameterUse, GongCopyBranchParameterShape(mapOrigCopy, _parametershape))
	}

	return
}

func GongCopyBranchParameterCategoryUse(mapOrigCopy map[any]any, parametercategoryuseFrom *ParameterCategoryUse) (parametercategoryuseTo *ParameterCategoryUse) {
	var alreadyCopied bool
	parametercategoryuseTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, parametercategoryuseFrom)
	if alreadyCopied {
		return
	}
	parametercategoryuseFrom.GongCopyBasicFields(parametercategoryuseTo)

	//insertion point for the staging of instances referenced by pointers
	if parametercategoryuseFrom.ParameterCategory != nil {
		parametercategoryuseTo.ParameterCategory = GongCopyBranchParameterCategory(mapOrigCopy, parametercategoryuseFrom.ParameterCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchParameterShape(mapOrigCopy map[any]any, parametershapeFrom *ParameterShape) (parametershapeTo *ParameterShape) {
	var alreadyCopied bool
	parametershapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, parametershapeFrom)
	if alreadyCopied {
		return
	}
	parametershapeFrom.GongCopyBasicFields(parametershapeTo)

	//insertion point for the staging of instances referenced by pointers
	if parametershapeFrom.Parameter != nil {
		parametershapeTo.Parameter = parametershapeFrom.Parameter
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchParametersAggregate(mapOrigCopy map[any]any, parametersaggregateFrom *ParametersAggregate) (parametersaggregateTo *ParametersAggregate) {
	var alreadyCopied bool
	parametersaggregateTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, parametersaggregateFrom)
	if alreadyCopied {
		return
	}
	parametersaggregateFrom.GongCopyBasicFields(parametersaggregateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parameter := range parametersaggregateFrom.Parameters {
		parametersaggregateTo.Parameters = append(parametersaggregateTo.Parameters, GongCopyBranchParameter(mapOrigCopy, _parameter))
	}

	return
}

func GongCopyBranchParametersAggregateShape(mapOrigCopy map[any]any, parametersaggregateshapeFrom *ParametersAggregateShape) (parametersaggregateshapeTo *ParametersAggregateShape) {
	var alreadyCopied bool
	parametersaggregateshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, parametersaggregateshapeFrom)
	if alreadyCopied {
		return
	}
	parametersaggregateshapeFrom.GongCopyBasicFields(parametersaggregateshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if parametersaggregateshapeFrom.ScenarioParameter != nil {
		parametersaggregateshapeTo.ScenarioParameter = parametersaggregateshapeFrom.ScenarioParameter
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPosition(mapOrigCopy map[any]any, positionFrom *Position) (positionTo *Position) {
	var alreadyCopied bool
	positionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, positionFrom)
	if alreadyCopied {
		return
	}
	positionFrom.GongCopyBasicFields(positionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRepository(mapOrigCopy map[any]any, repositoryFrom *Repository) (repositoryTo *Repository) {
	var alreadyCopied bool
	repositoryTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, repositoryFrom)
	if alreadyCopied {
		return
	}
	repositoryFrom.GongCopyBasicFields(repositoryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parametershape := range repositoryFrom.ParameterUse {
		repositoryTo.ParameterUse = append(repositoryTo.ParameterUse, GongCopyBranchParameterShape(mapOrigCopy, _parametershape))
	}
	for _, _groupuse := range repositoryFrom.GroupUse {
		repositoryTo.GroupUse = append(repositoryTo.GroupUse, GongCopyBranchGroupUse(mapOrigCopy, _groupuse))
	}

	return
}

func GongCopyBranchScenario(mapOrigCopy map[any]any, scenarioFrom *Scenario) (scenarioTo *Scenario) {
	var alreadyCopied bool
	scenarioTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, scenarioFrom)
	if alreadyCopied {
		return
	}
	scenarioFrom.GongCopyBasicFields(scenarioTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _diagram := range scenarioFrom.Diagrams {
		scenarioTo.Diagrams = append(scenarioTo.Diagrams, GongCopyBranchDiagram(mapOrigCopy, _diagram))
	}
	for _, _actorstate := range scenarioFrom.ActorStates {
		scenarioTo.ActorStates = append(scenarioTo.ActorStates, GongCopyBranchActorState(mapOrigCopy, _actorstate))
	}
	for _, _actorstatetransition := range scenarioFrom.ActorStateTransitions {
		scenarioTo.ActorStateTransitions = append(scenarioTo.ActorStateTransitions, GongCopyBranchActorStateTransition(mapOrigCopy, _actorstatetransition))
	}
	for _, _evolutiondirection := range scenarioFrom.EvolutionDirections {
		scenarioTo.EvolutionDirections = append(scenarioTo.EvolutionDirections, GongCopyBranchEvolutionDirection(mapOrigCopy, _evolutiondirection))
	}
	for _, _parameter := range scenarioFrom.Parameters {
		scenarioTo.Parameters = append(scenarioTo.Parameters, GongCopyBranchParameter(mapOrigCopy, _parameter))
	}
	for _, _parametersaggregate := range scenarioFrom.ParametersAggretates {
		scenarioTo.ParametersAggretates = append(scenarioTo.ParametersAggretates, GongCopyBranchParametersAggregate(mapOrigCopy, _parametersaggregate))
	}

	return
}

func GongCopyBranchUser(mapOrigCopy map[any]any, userFrom *User) (userTo *User) {
	var alreadyCopied bool
	userTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, userFrom)
	if alreadyCopied {
		return
	}
	userFrom.GongCopyBasicFields(userTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchUserUse(mapOrigCopy map[any]any, useruseFrom *UserUse) (useruseTo *UserUse) {
	var alreadyCopied bool
	useruseTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, useruseFrom)
	if alreadyCopied {
		return
	}
	useruseFrom.GongCopyBasicFields(useruseTo)

	//insertion point for the staging of instances referenced by pointers
	if useruseFrom.User != nil {
		useruseTo.User = GongCopyBranchUser(mapOrigCopy, useruseFrom.User)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchWorkspace(mapOrigCopy map[any]any, workspaceFrom *Workspace) (workspaceTo *Workspace) {
	var alreadyCopied bool
	workspaceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, workspaceFrom)
	if alreadyCopied {
		return
	}
	workspaceFrom.GongCopyBasicFields(workspaceTo)

	//insertion point for the staging of instances referenced by pointers
	if workspaceFrom.SelectedDiagram != nil {
		workspaceTo.SelectedDiagram = GongCopyBranchDiagram(mapOrigCopy, workspaceFrom.SelectedDiagram)
	}
	if workspaceFrom.Default_EvolutionDirectionShape != nil {
		workspaceTo.Default_EvolutionDirectionShape = GongCopyBranchEvolutionDirectionShape(mapOrigCopy, workspaceFrom.Default_EvolutionDirectionShape)
	}
	if workspaceFrom.Default_ParameterShape != nil {
		workspaceTo.Default_ParameterShape = GongCopyBranchParameterShape(mapOrigCopy, workspaceFrom.Default_ParameterShape)
	}
	if workspaceFrom.Default_ScenarioParameterShape != nil {
		workspaceTo.Default_ScenarioParameterShape = GongCopyBranchParametersAggregateShape(mapOrigCopy, workspaceFrom.Default_ScenarioParameterShape)
	}
	if workspaceFrom.Default_ActorStateShape != nil {
		workspaceTo.Default_ActorStateShape = GongCopyBranchActorStateShape(mapOrigCopy, workspaceFrom.Default_ActorStateShape)
	}
	if workspaceFrom.Default_ActorStateTransitionShape != nil {
		workspaceTo.Default_ActorStateTransitionShape = GongCopyBranchActorStateTransitionShape(mapOrigCopy, workspaceFrom.Default_ActorStateTransitionShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// insertion point for unstage branch per struct
func (actorstate *ActorState) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(actorstate) {
		return
	}

	actorstate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (actorstateshape *ActorStateShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(actorstateshape) {
		return
	}

	actorstateshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if actorstateshape.ActorState != nil {
		stage.UnstageBranch(actorstateshape.ActorState)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (actorstatetransition *ActorStateTransition) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(actorstatetransition) {
		return
	}

	actorstatetransition.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if actorstatetransition.StartState != nil {
		stage.UnstageBranch(actorstatetransition.StartState)
	}
	if actorstatetransition.EndState != nil {
		stage.UnstageBranch(actorstatetransition.EndState)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parameter := range actorstatetransition.Justifications {
		stage.UnstageBranch(_parameter)
	}

}

func (actorstatetransitionshape *ActorStateTransitionShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(actorstatetransitionshape) {
		return
	}

	actorstatetransitionshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if actorstatetransitionshape.ActorStateTransition != nil {
		stage.UnstageBranch(actorstatetransitionshape.ActorStateTransition)
	}
	if actorstatetransitionshape.Start != nil {
		stage.UnstageBranch(actorstatetransitionshape.Start)
	}
	if actorstatetransitionshape.End != nil {
		stage.UnstageBranch(actorstatetransitionshape.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range actorstatetransitionshape.ControlPointShapes {
		stage.UnstageBranch(_controlpointshape)
	}

}

func (analysis *Analysis) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(analysis) {
		return
	}

	analysis.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _scenario := range analysis.Scenarios {
		stage.UnstageBranch(_scenario)
	}
	for _, _groupuse := range analysis.GroupUse {
		stage.UnstageBranch(_groupuse)
	}
	for _, _geoobjectuse := range analysis.GeoObjectUse {
		stage.UnstageBranch(_geoobjectuse)
	}
	for _, _mapobjectuse := range analysis.MapUse {
		stage.UnstageBranch(_mapobjectuse)
	}

}

func (controlpointshape *ControlPointShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(controlpointshape) {
		return
	}

	controlpointshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (diagram *Diagram) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(diagram) {
		return
	}

	diagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _evolutiondirectionshape := range diagram.EvolutionDirectionShapes {
		stage.UnstageBranch(_evolutiondirectionshape)
	}
	for _, _evolutiondirection := range diagram.EvolutionDirectionsWhoseNodeIsExpanded {
		stage.UnstageBranch(_evolutiondirection)
	}
	for _, _actorstateshape := range diagram.ActorStateShapes {
		stage.UnstageBranch(_actorstateshape)
	}
	for _, _actorstate := range diagram.ActorStatesWhoseNodeIsExpanded {
		stage.UnstageBranch(_actorstate)
	}
	for _, _parametershape := range diagram.ParameterShapes {
		stage.UnstageBranch(_parametershape)
	}
	for _, _parameter := range diagram.ParametersWhoseNodeIsExpanded {
		stage.UnstageBranch(_parameter)
	}
	for _, _parametersaggregateshape := range diagram.ScenarioParameterShapes {
		stage.UnstageBranch(_parametersaggregateshape)
	}
	for _, _parametersaggregate := range diagram.ParametersAggregatesWhoseNodeIsExpanded {
		stage.UnstageBranch(_parametersaggregate)
	}
	for _, _actorstatetransitionshape := range diagram.ActorStateTransitionShapes {
		stage.UnstageBranch(_actorstatetransitionshape)
	}
	for _, _actorstatetransition := range diagram.ActorStateTransitionsWhoseNodeIsExpanded {
		stage.UnstageBranch(_actorstatetransition)
	}

}

func (document *Document) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(document) {
		return
	}

	document.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _geoobjectuse := range document.GeoObjectUse {
		stage.UnstageBranch(_geoobjectuse)
	}

}

func (documentuse *DocumentUse) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(documentuse) {
		return
	}

	documentuse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if documentuse.Document != nil {
		stage.UnstageBranch(documentuse.Document)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (evolutiondirection *EvolutionDirection) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(evolutiondirection) {
		return
	}

	evolutiondirection.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (evolutiondirectionshape *EvolutionDirectionShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(evolutiondirectionshape) {
		return
	}

	evolutiondirectionshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if evolutiondirectionshape.EvolutionDirection != nil {
		stage.UnstageBranch(evolutiondirectionshape.EvolutionDirection)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (foo *Foo) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(foo) {
		return
	}

	foo.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (geoobject *GeoObject) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(geoobject) {
		return
	}

	geoobject.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (geoobjectuse *GeoObjectUse) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(geoobjectuse) {
		return
	}

	geoobjectuse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if geoobjectuse.GeoObject != nil {
		stage.UnstageBranch(geoobjectuse.GeoObject)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (group *Group) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(group) {
		return
	}

	group.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _useruse := range group.UserUse {
		stage.UnstageBranch(_useruse)
	}

}

func (groupuse *GroupUse) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(groupuse) {
		return
	}

	groupuse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if groupuse.Group != nil {
		stage.UnstageBranch(groupuse.Group)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (library *Library) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(library) {
		return
	}

	library.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _analysis := range library.Analyses {
		stage.UnstageBranch(_analysis)
	}
	for _, _library := range library.SubLibraries {
		stage.UnstageBranch(_library)
	}
	for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
		stage.UnstageBranch(_library)
	}

}

func (mapobject *MapObject) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(mapobject) {
		return
	}

	mapobject.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (mapobjectuse *MapObjectUse) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(mapobjectuse) {
		return
	}

	mapobjectuse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if mapobjectuse.Map != nil {
		stage.UnstageBranch(mapobjectuse.Map)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (parameter *Parameter) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(parameter) {
		return
	}

	parameter.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _groupuse := range parameter.GroupUse {
		stage.UnstageBranch(_groupuse)
	}
	for _, _documentuse := range parameter.DocumentUse {
		stage.UnstageBranch(_documentuse)
	}
	for _, _geoobjectuse := range parameter.GeoObjectUse {
		stage.UnstageBranch(_geoobjectuse)
	}

}

func (parametercategory *ParameterCategory) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(parametercategory) {
		return
	}

	parametercategory.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parametershape := range parametercategory.ParameterUse {
		stage.UnstageBranch(_parametershape)
	}

}

func (parametercategoryuse *ParameterCategoryUse) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(parametercategoryuse) {
		return
	}

	parametercategoryuse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if parametercategoryuse.ParameterCategory != nil {
		stage.UnstageBranch(parametercategoryuse.ParameterCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (parametershape *ParameterShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(parametershape) {
		return
	}

	parametershape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if parametershape.Parameter != nil {
		stage.UnstageBranch(parametershape.Parameter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (parametersaggregate *ParametersAggregate) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(parametersaggregate) {
		return
	}

	parametersaggregate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parameter := range parametersaggregate.Parameters {
		stage.UnstageBranch(_parameter)
	}

}

func (parametersaggregateshape *ParametersAggregateShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(parametersaggregateshape) {
		return
	}

	parametersaggregateshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if parametersaggregateshape.ScenarioParameter != nil {
		stage.UnstageBranch(parametersaggregateshape.ScenarioParameter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (position *Position) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(position) {
		return
	}

	position.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (repository *Repository) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(repository) {
		return
	}

	repository.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _parametershape := range repository.ParameterUse {
		stage.UnstageBranch(_parametershape)
	}
	for _, _groupuse := range repository.GroupUse {
		stage.UnstageBranch(_groupuse)
	}

}

func (scenario *Scenario) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(scenario) {
		return
	}

	scenario.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _diagram := range scenario.Diagrams {
		stage.UnstageBranch(_diagram)
	}
	for _, _actorstate := range scenario.ActorStates {
		stage.UnstageBranch(_actorstate)
	}
	for _, _actorstatetransition := range scenario.ActorStateTransitions {
		stage.UnstageBranch(_actorstatetransition)
	}
	for _, _evolutiondirection := range scenario.EvolutionDirections {
		stage.UnstageBranch(_evolutiondirection)
	}
	for _, _parameter := range scenario.Parameters {
		stage.UnstageBranch(_parameter)
	}
	for _, _parametersaggregate := range scenario.ParametersAggretates {
		stage.UnstageBranch(_parametersaggregate)
	}

}

func (user *User) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(user) {
		return
	}

	user.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (useruse *UserUse) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(useruse) {
		return
	}

	useruse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if useruse.User != nil {
		stage.UnstageBranch(useruse.User)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (workspace *Workspace) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(workspace) {
		return
	}

	workspace.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if workspace.SelectedDiagram != nil {
		stage.UnstageBranch(workspace.SelectedDiagram)
	}
	if workspace.Default_EvolutionDirectionShape != nil {
		stage.UnstageBranch(workspace.Default_EvolutionDirectionShape)
	}
	if workspace.Default_ParameterShape != nil {
		stage.UnstageBranch(workspace.Default_ParameterShape)
	}
	if workspace.Default_ScenarioParameterShape != nil {
		stage.UnstageBranch(workspace.Default_ScenarioParameterShape)
	}
	if workspace.Default_ActorStateShape != nil {
		stage.UnstageBranch(workspace.Default_ActorStateShape)
	}
	if workspace.Default_ActorStateTransitionShape != nil {
		stage.UnstageBranch(workspace.Default_ActorStateTransitionShape)
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
	__gong__reconstructPointer(&reference.ActorState, stage.ActorStates_reference, instance.ActorState)
	// insertion point for slice of pointers field
}

func (reference *ActorStateTransition) GongReconstructPointersFromReferences(stage *Stage, instance *ActorStateTransition) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.StartState, stage.ActorStates_reference, instance.StartState)
	__gong__reconstructPointer(&reference.EndState, stage.ActorStates_reference, instance.EndState)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Justifications, stage.Parameters_reference, instance.Justifications)
}

func (reference *ActorStateTransitionShape) GongReconstructPointersFromReferences(stage *Stage, instance *ActorStateTransitionShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.ActorStateTransition, stage.ActorStateTransitions_reference, instance.ActorStateTransition)
	__gong__reconstructPointer(&reference.Start, stage.ActorStateShapes_reference, instance.Start)
	__gong__reconstructPointer(&reference.End, stage.ActorStateShapes_reference, instance.End)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ControlPointShapes, stage.ControlPointShapes_reference, instance.ControlPointShapes)
}

func (reference *Analysis) GongReconstructPointersFromReferences(stage *Stage, instance *Analysis) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Scenarios, stage.Scenarios_reference, instance.Scenarios)
	__gong__reconstructSliceOfPointersFromReferences(&reference.GroupUse, stage.GroupUses_reference, instance.GroupUse)
	__gong__reconstructSliceOfPointersFromReferences(&reference.GeoObjectUse, stage.GeoObjectUses_reference, instance.GeoObjectUse)
	__gong__reconstructSliceOfPointersFromReferences(&reference.MapUse, stage.MapObjectUses_reference, instance.MapUse)
}

func (reference *ControlPointShape) GongReconstructPointersFromReferences(stage *Stage, instance *ControlPointShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Diagram) GongReconstructPointersFromReferences(stage *Stage, instance *Diagram) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.EvolutionDirectionShapes, stage.EvolutionDirectionShapes_reference, instance.EvolutionDirectionShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.EvolutionDirectionsWhoseNodeIsExpanded, stage.EvolutionDirections_reference, instance.EvolutionDirectionsWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ActorStateShapes, stage.ActorStateShapes_reference, instance.ActorStateShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ActorStatesWhoseNodeIsExpanded, stage.ActorStates_reference, instance.ActorStatesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ParameterShapes, stage.ParameterShapes_reference, instance.ParameterShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ParametersWhoseNodeIsExpanded, stage.Parameters_reference, instance.ParametersWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ScenarioParameterShapes, stage.ParametersAggregateShapes_reference, instance.ScenarioParameterShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ParametersAggregatesWhoseNodeIsExpanded, stage.ParametersAggregates_reference, instance.ParametersAggregatesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ActorStateTransitionShapes, stage.ActorStateTransitionShapes_reference, instance.ActorStateTransitionShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ActorStateTransitionsWhoseNodeIsExpanded, stage.ActorStateTransitions_reference, instance.ActorStateTransitionsWhoseNodeIsExpanded)
}

func (reference *Document) GongReconstructPointersFromReferences(stage *Stage, instance *Document) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.GeoObjectUse, stage.GeoObjectUses_reference, instance.GeoObjectUse)
}

func (reference *DocumentUse) GongReconstructPointersFromReferences(stage *Stage, instance *DocumentUse) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Document, stage.Documents_reference, instance.Document)
	// insertion point for slice of pointers field
}

func (reference *EvolutionDirection) GongReconstructPointersFromReferences(stage *Stage, instance *EvolutionDirection) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EvolutionDirectionShape) GongReconstructPointersFromReferences(stage *Stage, instance *EvolutionDirectionShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.EvolutionDirection, stage.EvolutionDirections_reference, instance.EvolutionDirection)
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
	__gong__reconstructPointer(&reference.GeoObject, stage.GeoObjects_reference, instance.GeoObject)
	// insertion point for slice of pointers field
}

func (reference *Group) GongReconstructPointersFromReferences(stage *Stage, instance *Group) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.UserUse, stage.UserUses_reference, instance.UserUse)
}

func (reference *GroupUse) GongReconstructPointersFromReferences(stage *Stage, instance *GroupUse) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Group, stage.Groups_reference, instance.Group)
	// insertion point for slice of pointers field
}

func (reference *Library) GongReconstructPointersFromReferences(stage *Stage, instance *Library) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Analyses, stage.Analysiss_reference, instance.Analyses)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubLibraries, stage.Librarys_reference, instance.SubLibraries)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubLibrariesWhoseNodeIsExpanded, stage.Librarys_reference, instance.SubLibrariesWhoseNodeIsExpanded)
}

func (reference *MapObject) GongReconstructPointersFromReferences(stage *Stage, instance *MapObject) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MapObjectUse) GongReconstructPointersFromReferences(stage *Stage, instance *MapObjectUse) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Map, stage.MapObjects_reference, instance.Map)
	// insertion point for slice of pointers field
}

func (reference *Parameter) GongReconstructPointersFromReferences(stage *Stage, instance *Parameter) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.GroupUse, stage.GroupUses_reference, instance.GroupUse)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DocumentUse, stage.DocumentUses_reference, instance.DocumentUse)
	__gong__reconstructSliceOfPointersFromReferences(&reference.GeoObjectUse, stage.GeoObjectUses_reference, instance.GeoObjectUse)
}

func (reference *ParameterCategory) GongReconstructPointersFromReferences(stage *Stage, instance *ParameterCategory) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ParameterUse, stage.ParameterShapes_reference, instance.ParameterUse)
}

func (reference *ParameterCategoryUse) GongReconstructPointersFromReferences(stage *Stage, instance *ParameterCategoryUse) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.ParameterCategory, stage.ParameterCategorys_reference, instance.ParameterCategory)
	// insertion point for slice of pointers field
}

func (reference *ParameterShape) GongReconstructPointersFromReferences(stage *Stage, instance *ParameterShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Parameter, stage.Parameters_reference, instance.Parameter)
	// insertion point for slice of pointers field
}

func (reference *ParametersAggregate) GongReconstructPointersFromReferences(stage *Stage, instance *ParametersAggregate) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Parameters, stage.Parameters_reference, instance.Parameters)
}

func (reference *ParametersAggregateShape) GongReconstructPointersFromReferences(stage *Stage, instance *ParametersAggregateShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.ScenarioParameter, stage.ParametersAggregates_reference, instance.ScenarioParameter)
	// insertion point for slice of pointers field
}

func (reference *Position) GongReconstructPointersFromReferences(stage *Stage, instance *Position) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Repository) GongReconstructPointersFromReferences(stage *Stage, instance *Repository) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ParameterUse, stage.ParameterShapes_reference, instance.ParameterUse)
	__gong__reconstructSliceOfPointersFromReferences(&reference.GroupUse, stage.GroupUses_reference, instance.GroupUse)
}

func (reference *Scenario) GongReconstructPointersFromReferences(stage *Stage, instance *Scenario) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Diagrams, stage.Diagrams_reference, instance.Diagrams)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ActorStates, stage.ActorStates_reference, instance.ActorStates)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ActorStateTransitions, stage.ActorStateTransitions_reference, instance.ActorStateTransitions)
	__gong__reconstructSliceOfPointersFromReferences(&reference.EvolutionDirections, stage.EvolutionDirections_reference, instance.EvolutionDirections)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Parameters, stage.Parameters_reference, instance.Parameters)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ParametersAggretates, stage.ParametersAggregates_reference, instance.ParametersAggretates)
}

func (reference *User) GongReconstructPointersFromReferences(stage *Stage, instance *User) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *UserUse) GongReconstructPointersFromReferences(stage *Stage, instance *UserUse) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.User, stage.Users_reference, instance.User)
	// insertion point for slice of pointers field
}

func (reference *Workspace) GongReconstructPointersFromReferences(stage *Stage, instance *Workspace) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.SelectedDiagram, stage.Diagrams_reference, instance.SelectedDiagram)
	__gong__reconstructPointer(&reference.Default_EvolutionDirectionShape, stage.EvolutionDirectionShapes_reference, instance.Default_EvolutionDirectionShape)
	__gong__reconstructPointer(&reference.Default_ParameterShape, stage.ParameterShapes_reference, instance.Default_ParameterShape)
	__gong__reconstructPointer(&reference.Default_ScenarioParameterShape, stage.ParametersAggregateShapes_reference, instance.Default_ScenarioParameterShape)
	__gong__reconstructPointer(&reference.Default_ActorStateShape, stage.ActorStateShapes_reference, instance.Default_ActorStateShape)
	__gong__reconstructPointer(&reference.Default_ActorStateTransitionShape, stage.ActorStateTransitionShapes_reference, instance.Default_ActorStateTransitionShape)
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *ActorState) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ActorStateShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.ActorState, stage.ActorStates_instance)
	// insertion point for slice of pointers fields
}

func (reference *ActorStateTransition) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.StartState, stage.ActorStates_instance)
	__gong__reconstructPointerFromInstance(&reference.EndState, stage.ActorStates_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Justifications, stage.Parameters_instance)
}

func (reference *ActorStateTransitionShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.ActorStateTransition, stage.ActorStateTransitions_instance)
	__gong__reconstructPointerFromInstance(&reference.Start, stage.ActorStateShapes_instance)
	__gong__reconstructPointerFromInstance(&reference.End, stage.ActorStateShapes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ControlPointShapes, stage.ControlPointShapes_instance)
}

func (reference *Analysis) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Scenarios, stage.Scenarios_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.GroupUse, stage.GroupUses_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.GeoObjectUse, stage.GeoObjectUses_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.MapUse, stage.MapObjectUses_instance)
}

func (reference *ControlPointShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Diagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.EvolutionDirectionShapes, stage.EvolutionDirectionShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.EvolutionDirectionsWhoseNodeIsExpanded, stage.EvolutionDirections_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ActorStateShapes, stage.ActorStateShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ActorStatesWhoseNodeIsExpanded, stage.ActorStates_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ParameterShapes, stage.ParameterShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ParametersWhoseNodeIsExpanded, stage.Parameters_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ScenarioParameterShapes, stage.ParametersAggregateShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ParametersAggregatesWhoseNodeIsExpanded, stage.ParametersAggregates_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ActorStateTransitionShapes, stage.ActorStateTransitionShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ActorStateTransitionsWhoseNodeIsExpanded, stage.ActorStateTransitions_instance)
}

func (reference *Document) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.GeoObjectUse, stage.GeoObjectUses_instance)
}

func (reference *DocumentUse) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Document, stage.Documents_instance)
	// insertion point for slice of pointers fields
}

func (reference *EvolutionDirection) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EvolutionDirectionShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.EvolutionDirection, stage.EvolutionDirections_instance)
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
	__gong__reconstructPointerFromInstance(&reference.GeoObject, stage.GeoObjects_instance)
	// insertion point for slice of pointers fields
}

func (reference *Group) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.UserUse, stage.UserUses_instance)
}

func (reference *GroupUse) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Group, stage.Groups_instance)
	// insertion point for slice of pointers fields
}

func (reference *Library) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Analyses, stage.Analysiss_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubLibraries, stage.Librarys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubLibrariesWhoseNodeIsExpanded, stage.Librarys_instance)
}

func (reference *MapObject) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MapObjectUse) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Map, stage.MapObjects_instance)
	// insertion point for slice of pointers fields
}

func (reference *Parameter) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.GroupUse, stage.GroupUses_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DocumentUse, stage.DocumentUses_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.GeoObjectUse, stage.GeoObjectUses_instance)
}

func (reference *ParameterCategory) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ParameterUse, stage.ParameterShapes_instance)
}

func (reference *ParameterCategoryUse) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.ParameterCategory, stage.ParameterCategorys_instance)
	// insertion point for slice of pointers fields
}

func (reference *ParameterShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Parameter, stage.Parameters_instance)
	// insertion point for slice of pointers fields
}

func (reference *ParametersAggregate) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Parameters, stage.Parameters_instance)
}

func (reference *ParametersAggregateShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.ScenarioParameter, stage.ParametersAggregates_instance)
	// insertion point for slice of pointers fields
}

func (reference *Position) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Repository) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ParameterUse, stage.ParameterShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.GroupUse, stage.GroupUses_instance)
}

func (reference *Scenario) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Diagrams, stage.Diagrams_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ActorStates, stage.ActorStates_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ActorStateTransitions, stage.ActorStateTransitions_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.EvolutionDirections, stage.EvolutionDirections_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Parameters, stage.Parameters_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ParametersAggretates, stage.ParametersAggregates_instance)
}

func (reference *User) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *UserUse) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.User, stage.Users_instance)
	// insertion point for slice of pointers fields
}

func (reference *Workspace) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.SelectedDiagram, stage.Diagrams_instance)
	__gong__reconstructPointerFromInstance(&reference.Default_EvolutionDirectionShape, stage.EvolutionDirectionShapes_instance)
	__gong__reconstructPointerFromInstance(&reference.Default_ParameterShape, stage.ParameterShapes_instance)
	__gong__reconstructPointerFromInstance(&reference.Default_ScenarioParameterShape, stage.ParametersAggregateShapes_instance)
	__gong__reconstructPointerFromInstance(&reference.Default_ActorStateShape, stage.ActorStateShapes_instance)
	__gong__reconstructPointerFromInstance(&reference.Default_ActorStateTransitionShape, stage.ActorStateTransitionShapes_instance)
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

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (actorstateshape *ActorStateShape) GongDiff(stage *Stage, actorstateshapeOther *ActorStateShape) (diffs []string) {
	// insertion point for field diffs
	if actorstateshape.Name != actorstateshapeOther.Name {
		diffs = append(diffs, actorstateshape.GongMarshallField(stage, "Name"))
	}
	if actorstateshape.ActorState != actorstateshapeOther.ActorState {
		diffs = append(diffs, actorstateshape.GongMarshallField(stage, "ActorState"))
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
	if actorstatetransition.StartState != actorstatetransitionOther.StartState {
		diffs = append(diffs, actorstatetransition.GongMarshallField(stage, "StartState"))
	}
	if actorstatetransition.EndState != actorstatetransitionOther.EndState {
		diffs = append(diffs, actorstatetransition.GongMarshallField(stage, "EndState"))
	}
	if ops := __gong__diffSliceOfPointers(stage, actorstatetransition, "Justifications", actorstatetransitionOther.Justifications, actorstatetransition.Justifications); ops != "" {
		diffs = append(diffs, ops)
	}
	if actorstatetransition.ComputedPrefix != actorstatetransitionOther.ComputedPrefix {
		diffs = append(diffs, actorstatetransition.GongMarshallField(stage, "ComputedPrefix"))
	}
	if actorstatetransition.IsExpanded != actorstatetransitionOther.IsExpanded {
		diffs = append(diffs, actorstatetransition.GongMarshallField(stage, "IsExpanded"))
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
	if actorstatetransitionshape.ActorStateTransition != actorstatetransitionshapeOther.ActorStateTransition {
		diffs = append(diffs, actorstatetransitionshape.GongMarshallField(stage, "ActorStateTransition"))
	}
	if actorstatetransitionshape.Start != actorstatetransitionshapeOther.Start {
		diffs = append(diffs, actorstatetransitionshape.GongMarshallField(stage, "Start"))
	}
	if actorstatetransitionshape.End != actorstatetransitionshapeOther.End {
		diffs = append(diffs, actorstatetransitionshape.GongMarshallField(stage, "End"))
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
	if ops := __gong__diffSliceOfPointers(stage, actorstatetransitionshape, "ControlPointShapes", actorstatetransitionshapeOther.ControlPointShapes, actorstatetransitionshape.ControlPointShapes); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, analysis, "Scenarios", analysisOther.Scenarios, analysis.Scenarios); ops != "" {
		diffs = append(diffs, ops)
	}
	if analysis.IsScenariosNodeExpanded != analysisOther.IsScenariosNodeExpanded {
		diffs = append(diffs, analysis.GongMarshallField(stage, "IsScenariosNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, analysis, "GroupUse", analysisOther.GroupUse, analysis.GroupUse); ops != "" {
		diffs = append(diffs, ops)
	}
	if analysis.IsGroupUseNodeExpanded != analysisOther.IsGroupUseNodeExpanded {
		diffs = append(diffs, analysis.GongMarshallField(stage, "IsGroupUseNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, analysis, "GeoObjectUse", analysisOther.GeoObjectUse, analysis.GeoObjectUse); ops != "" {
		diffs = append(diffs, ops)
	}
	if analysis.IsGeoObjectUseNodeExpanded != analysisOther.IsGeoObjectUseNodeExpanded {
		diffs = append(diffs, analysis.GongMarshallField(stage, "IsGeoObjectUseNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, analysis, "MapUse", analysisOther.MapUse, analysis.MapUse); ops != "" {
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
	if diagram.IsChecked != diagramOther.IsChecked {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsChecked"))
	}
	if diagram.IsShowPrefix != diagramOther.IsShowPrefix {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsShowPrefix"))
	}
	if diagram.Description != diagramOther.Description {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "EvolutionDirectionShapes", diagramOther.EvolutionDirectionShapes, diagram.EvolutionDirectionShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "EvolutionDirectionsWhoseNodeIsExpanded", diagramOther.EvolutionDirectionsWhoseNodeIsExpanded, diagram.EvolutionDirectionsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagram.IsEvolutionDirectionsNodeExpanded != diagramOther.IsEvolutionDirectionsNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsEvolutionDirectionsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ActorStateShapes", diagramOther.ActorStateShapes, diagram.ActorStateShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ActorStatesWhoseNodeIsExpanded", diagramOther.ActorStatesWhoseNodeIsExpanded, diagram.ActorStatesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagram.IsActorStatesNodeExpanded != diagramOther.IsActorStatesNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsActorStatesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ParameterShapes", diagramOther.ParameterShapes, diagram.ParameterShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ParametersWhoseNodeIsExpanded", diagramOther.ParametersWhoseNodeIsExpanded, diagram.ParametersWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagram.IsParametersNodeExpanded != diagramOther.IsParametersNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsParametersNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ScenarioParameterShapes", diagramOther.ScenarioParameterShapes, diagram.ScenarioParameterShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ParametersAggregatesWhoseNodeIsExpanded", diagramOther.ParametersAggregatesWhoseNodeIsExpanded, diagram.ParametersAggregatesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagram.IsParametersAggregatesNodeExpanded != diagramOther.IsParametersAggregatesNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsParametersAggregatesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ActorStateTransitionShapes", diagramOther.ActorStateTransitionShapes, diagram.ActorStateTransitionShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ActorStateTransitionsWhoseNodeIsExpanded", diagramOther.ActorStateTransitionsWhoseNodeIsExpanded, diagram.ActorStateTransitionsWhoseNodeIsExpanded); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, document, "GeoObjectUse", documentOther.GeoObjectUse, document.GeoObjectUse); ops != "" {
		diffs = append(diffs, ops)
	}
	if document.ComputedPrefix != documentOther.ComputedPrefix {
		diffs = append(diffs, document.GongMarshallField(stage, "ComputedPrefix"))
	}
	if document.IsExpanded != documentOther.IsExpanded {
		diffs = append(diffs, document.GongMarshallField(stage, "IsExpanded"))
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
	if documentuse.Document != documentuseOther.Document {
		diffs = append(diffs, documentuse.GongMarshallField(stage, "Document"))
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

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (evolutiondirectionshape *EvolutionDirectionShape) GongDiff(stage *Stage, evolutiondirectionshapeOther *EvolutionDirectionShape) (diffs []string) {
	// insertion point for field diffs
	if evolutiondirectionshape.Name != evolutiondirectionshapeOther.Name {
		diffs = append(diffs, evolutiondirectionshape.GongMarshallField(stage, "Name"))
	}
	if evolutiondirectionshape.EvolutionDirection != evolutiondirectionshapeOther.EvolutionDirection {
		diffs = append(diffs, evolutiondirectionshape.GongMarshallField(stage, "EvolutionDirection"))
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

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (geoobjectuse *GeoObjectUse) GongDiff(stage *Stage, geoobjectuseOther *GeoObjectUse) (diffs []string) {
	// insertion point for field diffs
	if geoobjectuse.Name != geoobjectuseOther.Name {
		diffs = append(diffs, geoobjectuse.GongMarshallField(stage, "Name"))
	}
	if geoobjectuse.GeoObject != geoobjectuseOther.GeoObject {
		diffs = append(diffs, geoobjectuse.GongMarshallField(stage, "GeoObject"))
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
	if ops := __gong__diffSliceOfPointers(stage, group, "UserUse", groupOther.UserUse, group.UserUse); ops != "" {
		diffs = append(diffs, ops)
	}
	if group.ComputedPrefix != groupOther.ComputedPrefix {
		diffs = append(diffs, group.GongMarshallField(stage, "ComputedPrefix"))
	}
	if group.IsExpanded != groupOther.IsExpanded {
		diffs = append(diffs, group.GongMarshallField(stage, "IsExpanded"))
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
	if groupuse.Group != groupuseOther.Group {
		diffs = append(diffs, groupuse.GongMarshallField(stage, "Group"))
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
	if library.IsRootLibrary != libraryOther.IsRootLibrary {
		diffs = append(diffs, library.GongMarshallField(stage, "IsRootLibrary"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "Analyses", libraryOther.Analyses, library.Analyses); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsAnalysesNodeExpanded != libraryOther.IsAnalysesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsAnalysesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "SubLibraries", libraryOther.SubLibraries, library.SubLibraries); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsSubLibrariesNodeExpanded != libraryOther.IsSubLibrariesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsSubLibrariesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "SubLibrariesWhoseNodeIsExpanded", libraryOther.SubLibrariesWhoseNodeIsExpanded, library.SubLibrariesWhoseNodeIsExpanded); ops != "" {
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

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (mapobjectuse *MapObjectUse) GongDiff(stage *Stage, mapobjectuseOther *MapObjectUse) (diffs []string) {
	// insertion point for field diffs
	if mapobjectuse.Name != mapobjectuseOther.Name {
		diffs = append(diffs, mapobjectuse.GongMarshallField(stage, "Name"))
	}
	if mapobjectuse.Map != mapobjectuseOther.Map {
		diffs = append(diffs, mapobjectuse.GongMarshallField(stage, "Map"))
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
	if ops := __gong__diffSliceOfPointers(stage, parameter, "GroupUse", parameterOther.GroupUse, parameter.GroupUse); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, parameter, "DocumentUse", parameterOther.DocumentUse, parameter.DocumentUse); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, parameter, "GeoObjectUse", parameterOther.GeoObjectUse, parameter.GeoObjectUse); ops != "" {
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

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (parametercategory *ParameterCategory) GongDiff(stage *Stage, parametercategoryOther *ParameterCategory) (diffs []string) {
	// insertion point for field diffs
	if parametercategory.Name != parametercategoryOther.Name {
		diffs = append(diffs, parametercategory.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, parametercategory, "ParameterUse", parametercategoryOther.ParameterUse, parametercategory.ParameterUse); ops != "" {
		diffs = append(diffs, ops)
	}
	if parametercategory.ComputedPrefix != parametercategoryOther.ComputedPrefix {
		diffs = append(diffs, parametercategory.GongMarshallField(stage, "ComputedPrefix"))
	}
	if parametercategory.IsExpanded != parametercategoryOther.IsExpanded {
		diffs = append(diffs, parametercategory.GongMarshallField(stage, "IsExpanded"))
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
	if parametercategoryuse.ParameterCategory != parametercategoryuseOther.ParameterCategory {
		diffs = append(diffs, parametercategoryuse.GongMarshallField(stage, "ParameterCategory"))
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
	if parametershape.Parameter != parametershapeOther.Parameter {
		diffs = append(diffs, parametershape.GongMarshallField(stage, "Parameter"))
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
	if ops := __gong__diffSliceOfPointers(stage, parametersaggregate, "Parameters", parametersaggregateOther.Parameters, parametersaggregate.Parameters); ops != "" {
		diffs = append(diffs, ops)
	}
	if parametersaggregate.ComputedPrefix != parametersaggregateOther.ComputedPrefix {
		diffs = append(diffs, parametersaggregate.GongMarshallField(stage, "ComputedPrefix"))
	}
	if parametersaggregate.IsExpanded != parametersaggregateOther.IsExpanded {
		diffs = append(diffs, parametersaggregate.GongMarshallField(stage, "IsExpanded"))
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
	if parametersaggregateshape.ScenarioParameter != parametersaggregateshapeOther.ScenarioParameter {
		diffs = append(diffs, parametersaggregateshape.GongMarshallField(stage, "ScenarioParameter"))
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

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (repository *Repository) GongDiff(stage *Stage, repositoryOther *Repository) (diffs []string) {
	// insertion point for field diffs
	if repository.Name != repositoryOther.Name {
		diffs = append(diffs, repository.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, repository, "ParameterUse", repositoryOther.ParameterUse, repository.ParameterUse); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, repository, "GroupUse", repositoryOther.GroupUse, repository.GroupUse); ops != "" {
		diffs = append(diffs, ops)
	}
	if repository.ComputedPrefix != repositoryOther.ComputedPrefix {
		diffs = append(diffs, repository.GongMarshallField(stage, "ComputedPrefix"))
	}
	if repository.IsExpanded != repositoryOther.IsExpanded {
		diffs = append(diffs, repository.GongMarshallField(stage, "IsExpanded"))
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
	if ops := __gong__diffSliceOfPointers(stage, scenario, "Diagrams", scenarioOther.Diagrams, scenario.Diagrams); ops != "" {
		diffs = append(diffs, ops)
	}
	if scenario.IsDiagramsNodeExpanded != scenarioOther.IsDiagramsNodeExpanded {
		diffs = append(diffs, scenario.GongMarshallField(stage, "IsDiagramsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, scenario, "ActorStates", scenarioOther.ActorStates, scenario.ActorStates); ops != "" {
		diffs = append(diffs, ops)
	}
	if scenario.IsActorStatesNodeExpanded != scenarioOther.IsActorStatesNodeExpanded {
		diffs = append(diffs, scenario.GongMarshallField(stage, "IsActorStatesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, scenario, "ActorStateTransitions", scenarioOther.ActorStateTransitions, scenario.ActorStateTransitions); ops != "" {
		diffs = append(diffs, ops)
	}
	if scenario.IsActorStateTransitionsNodeExpanded != scenarioOther.IsActorStateTransitionsNodeExpanded {
		diffs = append(diffs, scenario.GongMarshallField(stage, "IsActorStateTransitionsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, scenario, "EvolutionDirections", scenarioOther.EvolutionDirections, scenario.EvolutionDirections); ops != "" {
		diffs = append(diffs, ops)
	}
	if scenario.IsEvolutionDirectionsNodeExpanded != scenarioOther.IsEvolutionDirectionsNodeExpanded {
		diffs = append(diffs, scenario.GongMarshallField(stage, "IsEvolutionDirectionsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, scenario, "Parameters", scenarioOther.Parameters, scenario.Parameters); ops != "" {
		diffs = append(diffs, ops)
	}
	if scenario.IsParametersNodeExpanded != scenarioOther.IsParametersNodeExpanded {
		diffs = append(diffs, scenario.GongMarshallField(stage, "IsParametersNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, scenario, "ParametersAggretates", scenarioOther.ParametersAggretates, scenario.ParametersAggretates); ops != "" {
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

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (useruse *UserUse) GongDiff(stage *Stage, useruseOther *UserUse) (diffs []string) {
	// insertion point for field diffs
	if useruse.Name != useruseOther.Name {
		diffs = append(diffs, useruse.GongMarshallField(stage, "Name"))
	}
	if useruse.User != useruseOther.User {
		diffs = append(diffs, useruse.GongMarshallField(stage, "User"))
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
	if workspace.SelectedDiagram != workspaceOther.SelectedDiagram {
		diffs = append(diffs, workspace.GongMarshallField(stage, "SelectedDiagram"))
	}
	if workspace.Default_EvolutionDirectionShape != workspaceOther.Default_EvolutionDirectionShape {
		diffs = append(diffs, workspace.GongMarshallField(stage, "Default_EvolutionDirectionShape"))
	}
	if workspace.Default_ParameterShape != workspaceOther.Default_ParameterShape {
		diffs = append(diffs, workspace.GongMarshallField(stage, "Default_ParameterShape"))
	}
	if workspace.Default_ScenarioParameterShape != workspaceOther.Default_ScenarioParameterShape {
		diffs = append(diffs, workspace.GongMarshallField(stage, "Default_ScenarioParameterShape"))
	}
	if workspace.Default_ActorStateShape != workspaceOther.Default_ActorStateShape {
		diffs = append(diffs, workspace.GongMarshallField(stage, "Default_ActorStateShape"))
	}
	if workspace.Default_ActorStateTransitionShape != workspaceOther.Default_ActorStateTransitionShape {
		diffs = append(diffs, workspace.GongMarshallField(stage, "Default_ActorStateTransitionShape"))
	}
	if workspace.ComputedPrefix != workspaceOther.ComputedPrefix {
		diffs = append(diffs, workspace.GongMarshallField(stage, "ComputedPrefix"))
	}
	if workspace.IsExpanded != workspaceOther.IsExpanded {
		diffs = append(diffs, workspace.GongMarshallField(stage, "IsExpanded"))
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
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

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
