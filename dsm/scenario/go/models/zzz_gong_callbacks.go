// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *ActorState:
		if stage.OnAfterActorStateCreateCallback != nil {
			stage.OnAfterActorStateCreateCallback.OnAfterCreate(stage, target)
		}
	case *ActorStateShape:
		if stage.OnAfterActorStateShapeCreateCallback != nil {
			stage.OnAfterActorStateShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ActorStateTransition:
		if stage.OnAfterActorStateTransitionCreateCallback != nil {
			stage.OnAfterActorStateTransitionCreateCallback.OnAfterCreate(stage, target)
		}
	case *ActorStateTransitionShape:
		if stage.OnAfterActorStateTransitionShapeCreateCallback != nil {
			stage.OnAfterActorStateTransitionShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Analysis:
		if stage.OnAfterAnalysisCreateCallback != nil {
			stage.OnAfterAnalysisCreateCallback.OnAfterCreate(stage, target)
		}
	case *ControlPointShape:
		if stage.OnAfterControlPointShapeCreateCallback != nil {
			stage.OnAfterControlPointShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Diagram:
		if stage.OnAfterDiagramCreateCallback != nil {
			stage.OnAfterDiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *Document:
		if stage.OnAfterDocumentCreateCallback != nil {
			stage.OnAfterDocumentCreateCallback.OnAfterCreate(stage, target)
		}
	case *DocumentUse:
		if stage.OnAfterDocumentUseCreateCallback != nil {
			stage.OnAfterDocumentUseCreateCallback.OnAfterCreate(stage, target)
		}
	case *EvolutionDirection:
		if stage.OnAfterEvolutionDirectionCreateCallback != nil {
			stage.OnAfterEvolutionDirectionCreateCallback.OnAfterCreate(stage, target)
		}
	case *EvolutionDirectionShape:
		if stage.OnAfterEvolutionDirectionShapeCreateCallback != nil {
			stage.OnAfterEvolutionDirectionShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Foo:
		if stage.OnAfterFooCreateCallback != nil {
			stage.OnAfterFooCreateCallback.OnAfterCreate(stage, target)
		}
	case *GeoObject:
		if stage.OnAfterGeoObjectCreateCallback != nil {
			stage.OnAfterGeoObjectCreateCallback.OnAfterCreate(stage, target)
		}
	case *GeoObjectUse:
		if stage.OnAfterGeoObjectUseCreateCallback != nil {
			stage.OnAfterGeoObjectUseCreateCallback.OnAfterCreate(stage, target)
		}
	case *Group:
		if stage.OnAfterGroupCreateCallback != nil {
			stage.OnAfterGroupCreateCallback.OnAfterCreate(stage, target)
		}
	case *GroupUse:
		if stage.OnAfterGroupUseCreateCallback != nil {
			stage.OnAfterGroupUseCreateCallback.OnAfterCreate(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryCreateCallback != nil {
			stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, target)
		}
	case *MapObject:
		if stage.OnAfterMapObjectCreateCallback != nil {
			stage.OnAfterMapObjectCreateCallback.OnAfterCreate(stage, target)
		}
	case *MapObjectUse:
		if stage.OnAfterMapObjectUseCreateCallback != nil {
			stage.OnAfterMapObjectUseCreateCallback.OnAfterCreate(stage, target)
		}
	case *Parameter:
		if stage.OnAfterParameterCreateCallback != nil {
			stage.OnAfterParameterCreateCallback.OnAfterCreate(stage, target)
		}
	case *ParameterCategory:
		if stage.OnAfterParameterCategoryCreateCallback != nil {
			stage.OnAfterParameterCategoryCreateCallback.OnAfterCreate(stage, target)
		}
	case *ParameterCategoryUse:
		if stage.OnAfterParameterCategoryUseCreateCallback != nil {
			stage.OnAfterParameterCategoryUseCreateCallback.OnAfterCreate(stage, target)
		}
	case *ParameterShape:
		if stage.OnAfterParameterShapeCreateCallback != nil {
			stage.OnAfterParameterShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ParametersAggregate:
		if stage.OnAfterParametersAggregateCreateCallback != nil {
			stage.OnAfterParametersAggregateCreateCallback.OnAfterCreate(stage, target)
		}
	case *ParametersAggregateShape:
		if stage.OnAfterParametersAggregateShapeCreateCallback != nil {
			stage.OnAfterParametersAggregateShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Position:
		if stage.OnAfterPositionCreateCallback != nil {
			stage.OnAfterPositionCreateCallback.OnAfterCreate(stage, target)
		}
	case *Repository:
		if stage.OnAfterRepositoryCreateCallback != nil {
			stage.OnAfterRepositoryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Scenario:
		if stage.OnAfterScenarioCreateCallback != nil {
			stage.OnAfterScenarioCreateCallback.OnAfterCreate(stage, target)
		}
	case *User:
		if stage.OnAfterUserCreateCallback != nil {
			stage.OnAfterUserCreateCallback.OnAfterCreate(stage, target)
		}
	case *UserUse:
		if stage.OnAfterUserUseCreateCallback != nil {
			stage.OnAfterUserUseCreateCallback.OnAfterCreate(stage, target)
		}
	case *Workspace:
		if stage.OnAfterWorkspaceCreateCallback != nil {
			stage.OnAfterWorkspaceCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *ActorState:
		newTarget := any(new).(*ActorState)
		if stage.OnAfterActorStateUpdateCallback != nil {
			stage.OnAfterActorStateUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ActorStateShape:
		newTarget := any(new).(*ActorStateShape)
		if stage.OnAfterActorStateShapeUpdateCallback != nil {
			stage.OnAfterActorStateShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ActorStateTransition:
		newTarget := any(new).(*ActorStateTransition)
		if stage.OnAfterActorStateTransitionUpdateCallback != nil {
			stage.OnAfterActorStateTransitionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ActorStateTransitionShape:
		newTarget := any(new).(*ActorStateTransitionShape)
		if stage.OnAfterActorStateTransitionShapeUpdateCallback != nil {
			stage.OnAfterActorStateTransitionShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Analysis:
		newTarget := any(new).(*Analysis)
		if stage.OnAfterAnalysisUpdateCallback != nil {
			stage.OnAfterAnalysisUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ControlPointShape:
		newTarget := any(new).(*ControlPointShape)
		if stage.OnAfterControlPointShapeUpdateCallback != nil {
			stage.OnAfterControlPointShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Diagram:
		newTarget := any(new).(*Diagram)
		if stage.OnAfterDiagramUpdateCallback != nil {
			stage.OnAfterDiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Document:
		newTarget := any(new).(*Document)
		if stage.OnAfterDocumentUpdateCallback != nil {
			stage.OnAfterDocumentUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DocumentUse:
		newTarget := any(new).(*DocumentUse)
		if stage.OnAfterDocumentUseUpdateCallback != nil {
			stage.OnAfterDocumentUseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *EvolutionDirection:
		newTarget := any(new).(*EvolutionDirection)
		if stage.OnAfterEvolutionDirectionUpdateCallback != nil {
			stage.OnAfterEvolutionDirectionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *EvolutionDirectionShape:
		newTarget := any(new).(*EvolutionDirectionShape)
		if stage.OnAfterEvolutionDirectionShapeUpdateCallback != nil {
			stage.OnAfterEvolutionDirectionShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Foo:
		newTarget := any(new).(*Foo)
		if stage.OnAfterFooUpdateCallback != nil {
			stage.OnAfterFooUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GeoObject:
		newTarget := any(new).(*GeoObject)
		if stage.OnAfterGeoObjectUpdateCallback != nil {
			stage.OnAfterGeoObjectUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GeoObjectUse:
		newTarget := any(new).(*GeoObjectUse)
		if stage.OnAfterGeoObjectUseUpdateCallback != nil {
			stage.OnAfterGeoObjectUseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Group:
		newTarget := any(new).(*Group)
		if stage.OnAfterGroupUpdateCallback != nil {
			stage.OnAfterGroupUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GroupUse:
		newTarget := any(new).(*GroupUse)
		if stage.OnAfterGroupUseUpdateCallback != nil {
			stage.OnAfterGroupUseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Library:
		newTarget := any(new).(*Library)
		if stage.OnAfterLibraryUpdateCallback != nil {
			stage.OnAfterLibraryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MapObject:
		newTarget := any(new).(*MapObject)
		if stage.OnAfterMapObjectUpdateCallback != nil {
			stage.OnAfterMapObjectUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MapObjectUse:
		newTarget := any(new).(*MapObjectUse)
		if stage.OnAfterMapObjectUseUpdateCallback != nil {
			stage.OnAfterMapObjectUseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Parameter:
		newTarget := any(new).(*Parameter)
		if stage.OnAfterParameterUpdateCallback != nil {
			stage.OnAfterParameterUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ParameterCategory:
		newTarget := any(new).(*ParameterCategory)
		if stage.OnAfterParameterCategoryUpdateCallback != nil {
			stage.OnAfterParameterCategoryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ParameterCategoryUse:
		newTarget := any(new).(*ParameterCategoryUse)
		if stage.OnAfterParameterCategoryUseUpdateCallback != nil {
			stage.OnAfterParameterCategoryUseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ParameterShape:
		newTarget := any(new).(*ParameterShape)
		if stage.OnAfterParameterShapeUpdateCallback != nil {
			stage.OnAfterParameterShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ParametersAggregate:
		newTarget := any(new).(*ParametersAggregate)
		if stage.OnAfterParametersAggregateUpdateCallback != nil {
			stage.OnAfterParametersAggregateUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ParametersAggregateShape:
		newTarget := any(new).(*ParametersAggregateShape)
		if stage.OnAfterParametersAggregateShapeUpdateCallback != nil {
			stage.OnAfterParametersAggregateShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Position:
		newTarget := any(new).(*Position)
		if stage.OnAfterPositionUpdateCallback != nil {
			stage.OnAfterPositionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Repository:
		newTarget := any(new).(*Repository)
		if stage.OnAfterRepositoryUpdateCallback != nil {
			stage.OnAfterRepositoryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Scenario:
		newTarget := any(new).(*Scenario)
		if stage.OnAfterScenarioUpdateCallback != nil {
			stage.OnAfterScenarioUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *User:
		newTarget := any(new).(*User)
		if stage.OnAfterUserUpdateCallback != nil {
			stage.OnAfterUserUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *UserUse:
		newTarget := any(new).(*UserUse)
		if stage.OnAfterUserUseUpdateCallback != nil {
			stage.OnAfterUserUseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Workspace:
		newTarget := any(new).(*Workspace)
		if stage.OnAfterWorkspaceUpdateCallback != nil {
			stage.OnAfterWorkspaceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *ActorState:
		if stage.OnAfterActorStateDeleteCallback != nil {
			staged := any(staged).(*ActorState)
			stage.OnAfterActorStateDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ActorStateShape:
		if stage.OnAfterActorStateShapeDeleteCallback != nil {
			staged := any(staged).(*ActorStateShape)
			stage.OnAfterActorStateShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ActorStateTransition:
		if stage.OnAfterActorStateTransitionDeleteCallback != nil {
			staged := any(staged).(*ActorStateTransition)
			stage.OnAfterActorStateTransitionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ActorStateTransitionShape:
		if stage.OnAfterActorStateTransitionShapeDeleteCallback != nil {
			staged := any(staged).(*ActorStateTransitionShape)
			stage.OnAfterActorStateTransitionShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Analysis:
		if stage.OnAfterAnalysisDeleteCallback != nil {
			staged := any(staged).(*Analysis)
			stage.OnAfterAnalysisDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ControlPointShape:
		if stage.OnAfterControlPointShapeDeleteCallback != nil {
			staged := any(staged).(*ControlPointShape)
			stage.OnAfterControlPointShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Diagram:
		if stage.OnAfterDiagramDeleteCallback != nil {
			staged := any(staged).(*Diagram)
			stage.OnAfterDiagramDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Document:
		if stage.OnAfterDocumentDeleteCallback != nil {
			staged := any(staged).(*Document)
			stage.OnAfterDocumentDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DocumentUse:
		if stage.OnAfterDocumentUseDeleteCallback != nil {
			staged := any(staged).(*DocumentUse)
			stage.OnAfterDocumentUseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *EvolutionDirection:
		if stage.OnAfterEvolutionDirectionDeleteCallback != nil {
			staged := any(staged).(*EvolutionDirection)
			stage.OnAfterEvolutionDirectionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *EvolutionDirectionShape:
		if stage.OnAfterEvolutionDirectionShapeDeleteCallback != nil {
			staged := any(staged).(*EvolutionDirectionShape)
			stage.OnAfterEvolutionDirectionShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Foo:
		if stage.OnAfterFooDeleteCallback != nil {
			staged := any(staged).(*Foo)
			stage.OnAfterFooDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GeoObject:
		if stage.OnAfterGeoObjectDeleteCallback != nil {
			staged := any(staged).(*GeoObject)
			stage.OnAfterGeoObjectDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GeoObjectUse:
		if stage.OnAfterGeoObjectUseDeleteCallback != nil {
			staged := any(staged).(*GeoObjectUse)
			stage.OnAfterGeoObjectUseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Group:
		if stage.OnAfterGroupDeleteCallback != nil {
			staged := any(staged).(*Group)
			stage.OnAfterGroupDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GroupUse:
		if stage.OnAfterGroupUseDeleteCallback != nil {
			staged := any(staged).(*GroupUse)
			stage.OnAfterGroupUseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Library:
		if stage.OnAfterLibraryDeleteCallback != nil {
			staged := any(staged).(*Library)
			stage.OnAfterLibraryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MapObject:
		if stage.OnAfterMapObjectDeleteCallback != nil {
			staged := any(staged).(*MapObject)
			stage.OnAfterMapObjectDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MapObjectUse:
		if stage.OnAfterMapObjectUseDeleteCallback != nil {
			staged := any(staged).(*MapObjectUse)
			stage.OnAfterMapObjectUseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Parameter:
		if stage.OnAfterParameterDeleteCallback != nil {
			staged := any(staged).(*Parameter)
			stage.OnAfterParameterDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ParameterCategory:
		if stage.OnAfterParameterCategoryDeleteCallback != nil {
			staged := any(staged).(*ParameterCategory)
			stage.OnAfterParameterCategoryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ParameterCategoryUse:
		if stage.OnAfterParameterCategoryUseDeleteCallback != nil {
			staged := any(staged).(*ParameterCategoryUse)
			stage.OnAfterParameterCategoryUseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ParameterShape:
		if stage.OnAfterParameterShapeDeleteCallback != nil {
			staged := any(staged).(*ParameterShape)
			stage.OnAfterParameterShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ParametersAggregate:
		if stage.OnAfterParametersAggregateDeleteCallback != nil {
			staged := any(staged).(*ParametersAggregate)
			stage.OnAfterParametersAggregateDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ParametersAggregateShape:
		if stage.OnAfterParametersAggregateShapeDeleteCallback != nil {
			staged := any(staged).(*ParametersAggregateShape)
			stage.OnAfterParametersAggregateShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Position:
		if stage.OnAfterPositionDeleteCallback != nil {
			staged := any(staged).(*Position)
			stage.OnAfterPositionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Repository:
		if stage.OnAfterRepositoryDeleteCallback != nil {
			staged := any(staged).(*Repository)
			stage.OnAfterRepositoryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Scenario:
		if stage.OnAfterScenarioDeleteCallback != nil {
			staged := any(staged).(*Scenario)
			stage.OnAfterScenarioDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *User:
		if stage.OnAfterUserDeleteCallback != nil {
			staged := any(staged).(*User)
			stage.OnAfterUserDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *UserUse:
		if stage.OnAfterUserUseDeleteCallback != nil {
			staged := any(staged).(*UserUse)
			stage.OnAfterUserUseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Workspace:
		if stage.OnAfterWorkspaceDeleteCallback != nil {
			staged := any(staged).(*Workspace)
			stage.OnAfterWorkspaceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *ActorState:
		if stage.OnAfterActorStateReadCallback != nil {
			stage.OnAfterActorStateReadCallback.OnAfterRead(stage, target)
		}
	case *ActorStateShape:
		if stage.OnAfterActorStateShapeReadCallback != nil {
			stage.OnAfterActorStateShapeReadCallback.OnAfterRead(stage, target)
		}
	case *ActorStateTransition:
		if stage.OnAfterActorStateTransitionReadCallback != nil {
			stage.OnAfterActorStateTransitionReadCallback.OnAfterRead(stage, target)
		}
	case *ActorStateTransitionShape:
		if stage.OnAfterActorStateTransitionShapeReadCallback != nil {
			stage.OnAfterActorStateTransitionShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Analysis:
		if stage.OnAfterAnalysisReadCallback != nil {
			stage.OnAfterAnalysisReadCallback.OnAfterRead(stage, target)
		}
	case *ControlPointShape:
		if stage.OnAfterControlPointShapeReadCallback != nil {
			stage.OnAfterControlPointShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Diagram:
		if stage.OnAfterDiagramReadCallback != nil {
			stage.OnAfterDiagramReadCallback.OnAfterRead(stage, target)
		}
	case *Document:
		if stage.OnAfterDocumentReadCallback != nil {
			stage.OnAfterDocumentReadCallback.OnAfterRead(stage, target)
		}
	case *DocumentUse:
		if stage.OnAfterDocumentUseReadCallback != nil {
			stage.OnAfterDocumentUseReadCallback.OnAfterRead(stage, target)
		}
	case *EvolutionDirection:
		if stage.OnAfterEvolutionDirectionReadCallback != nil {
			stage.OnAfterEvolutionDirectionReadCallback.OnAfterRead(stage, target)
		}
	case *EvolutionDirectionShape:
		if stage.OnAfterEvolutionDirectionShapeReadCallback != nil {
			stage.OnAfterEvolutionDirectionShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Foo:
		if stage.OnAfterFooReadCallback != nil {
			stage.OnAfterFooReadCallback.OnAfterRead(stage, target)
		}
	case *GeoObject:
		if stage.OnAfterGeoObjectReadCallback != nil {
			stage.OnAfterGeoObjectReadCallback.OnAfterRead(stage, target)
		}
	case *GeoObjectUse:
		if stage.OnAfterGeoObjectUseReadCallback != nil {
			stage.OnAfterGeoObjectUseReadCallback.OnAfterRead(stage, target)
		}
	case *Group:
		if stage.OnAfterGroupReadCallback != nil {
			stage.OnAfterGroupReadCallback.OnAfterRead(stage, target)
		}
	case *GroupUse:
		if stage.OnAfterGroupUseReadCallback != nil {
			stage.OnAfterGroupUseReadCallback.OnAfterRead(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryReadCallback != nil {
			stage.OnAfterLibraryReadCallback.OnAfterRead(stage, target)
		}
	case *MapObject:
		if stage.OnAfterMapObjectReadCallback != nil {
			stage.OnAfterMapObjectReadCallback.OnAfterRead(stage, target)
		}
	case *MapObjectUse:
		if stage.OnAfterMapObjectUseReadCallback != nil {
			stage.OnAfterMapObjectUseReadCallback.OnAfterRead(stage, target)
		}
	case *Parameter:
		if stage.OnAfterParameterReadCallback != nil {
			stage.OnAfterParameterReadCallback.OnAfterRead(stage, target)
		}
	case *ParameterCategory:
		if stage.OnAfterParameterCategoryReadCallback != nil {
			stage.OnAfterParameterCategoryReadCallback.OnAfterRead(stage, target)
		}
	case *ParameterCategoryUse:
		if stage.OnAfterParameterCategoryUseReadCallback != nil {
			stage.OnAfterParameterCategoryUseReadCallback.OnAfterRead(stage, target)
		}
	case *ParameterShape:
		if stage.OnAfterParameterShapeReadCallback != nil {
			stage.OnAfterParameterShapeReadCallback.OnAfterRead(stage, target)
		}
	case *ParametersAggregate:
		if stage.OnAfterParametersAggregateReadCallback != nil {
			stage.OnAfterParametersAggregateReadCallback.OnAfterRead(stage, target)
		}
	case *ParametersAggregateShape:
		if stage.OnAfterParametersAggregateShapeReadCallback != nil {
			stage.OnAfterParametersAggregateShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Position:
		if stage.OnAfterPositionReadCallback != nil {
			stage.OnAfterPositionReadCallback.OnAfterRead(stage, target)
		}
	case *Repository:
		if stage.OnAfterRepositoryReadCallback != nil {
			stage.OnAfterRepositoryReadCallback.OnAfterRead(stage, target)
		}
	case *Scenario:
		if stage.OnAfterScenarioReadCallback != nil {
			stage.OnAfterScenarioReadCallback.OnAfterRead(stage, target)
		}
	case *User:
		if stage.OnAfterUserReadCallback != nil {
			stage.OnAfterUserReadCallback.OnAfterRead(stage, target)
		}
	case *UserUse:
		if stage.OnAfterUserUseReadCallback != nil {
			stage.OnAfterUserUseReadCallback.OnAfterRead(stage, target)
		}
	case *Workspace:
		if stage.OnAfterWorkspaceReadCallback != nil {
			stage.OnAfterWorkspaceReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *Stage, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *ActorState:
		stage.OnAfterActorStateUpdateCallback = any(callback).(OnAfterUpdateInterface[ActorState])
	case *ActorStateShape:
		stage.OnAfterActorStateShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ActorStateShape])
	case *ActorStateTransition:
		stage.OnAfterActorStateTransitionUpdateCallback = any(callback).(OnAfterUpdateInterface[ActorStateTransition])
	case *ActorStateTransitionShape:
		stage.OnAfterActorStateTransitionShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ActorStateTransitionShape])
	case *Analysis:
		stage.OnAfterAnalysisUpdateCallback = any(callback).(OnAfterUpdateInterface[Analysis])
	case *ControlPointShape:
		stage.OnAfterControlPointShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ControlPointShape])
	case *Diagram:
		stage.OnAfterDiagramUpdateCallback = any(callback).(OnAfterUpdateInterface[Diagram])
	case *Document:
		stage.OnAfterDocumentUpdateCallback = any(callback).(OnAfterUpdateInterface[Document])
	case *DocumentUse:
		stage.OnAfterDocumentUseUpdateCallback = any(callback).(OnAfterUpdateInterface[DocumentUse])
	case *EvolutionDirection:
		stage.OnAfterEvolutionDirectionUpdateCallback = any(callback).(OnAfterUpdateInterface[EvolutionDirection])
	case *EvolutionDirectionShape:
		stage.OnAfterEvolutionDirectionShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[EvolutionDirectionShape])
	case *Foo:
		stage.OnAfterFooUpdateCallback = any(callback).(OnAfterUpdateInterface[Foo])
	case *GeoObject:
		stage.OnAfterGeoObjectUpdateCallback = any(callback).(OnAfterUpdateInterface[GeoObject])
	case *GeoObjectUse:
		stage.OnAfterGeoObjectUseUpdateCallback = any(callback).(OnAfterUpdateInterface[GeoObjectUse])
	case *Group:
		stage.OnAfterGroupUpdateCallback = any(callback).(OnAfterUpdateInterface[Group])
	case *GroupUse:
		stage.OnAfterGroupUseUpdateCallback = any(callback).(OnAfterUpdateInterface[GroupUse])
	case *Library:
		stage.OnAfterLibraryUpdateCallback = any(callback).(OnAfterUpdateInterface[Library])
	case *MapObject:
		stage.OnAfterMapObjectUpdateCallback = any(callback).(OnAfterUpdateInterface[MapObject])
	case *MapObjectUse:
		stage.OnAfterMapObjectUseUpdateCallback = any(callback).(OnAfterUpdateInterface[MapObjectUse])
	case *Parameter:
		stage.OnAfterParameterUpdateCallback = any(callback).(OnAfterUpdateInterface[Parameter])
	case *ParameterCategory:
		stage.OnAfterParameterCategoryUpdateCallback = any(callback).(OnAfterUpdateInterface[ParameterCategory])
	case *ParameterCategoryUse:
		stage.OnAfterParameterCategoryUseUpdateCallback = any(callback).(OnAfterUpdateInterface[ParameterCategoryUse])
	case *ParameterShape:
		stage.OnAfterParameterShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ParameterShape])
	case *ParametersAggregate:
		stage.OnAfterParametersAggregateUpdateCallback = any(callback).(OnAfterUpdateInterface[ParametersAggregate])
	case *ParametersAggregateShape:
		stage.OnAfterParametersAggregateShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ParametersAggregateShape])
	case *Position:
		stage.OnAfterPositionUpdateCallback = any(callback).(OnAfterUpdateInterface[Position])
	case *Repository:
		stage.OnAfterRepositoryUpdateCallback = any(callback).(OnAfterUpdateInterface[Repository])
	case *Scenario:
		stage.OnAfterScenarioUpdateCallback = any(callback).(OnAfterUpdateInterface[Scenario])
	case *User:
		stage.OnAfterUserUpdateCallback = any(callback).(OnAfterUpdateInterface[User])
	case *UserUse:
		stage.OnAfterUserUseUpdateCallback = any(callback).(OnAfterUpdateInterface[UserUse])
	case *Workspace:
		stage.OnAfterWorkspaceUpdateCallback = any(callback).(OnAfterUpdateInterface[Workspace])
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *ActorState:
		stage.OnAfterActorStateCreateCallback = any(callback).(OnAfterCreateInterface[ActorState])
	case *ActorStateShape:
		stage.OnAfterActorStateShapeCreateCallback = any(callback).(OnAfterCreateInterface[ActorStateShape])
	case *ActorStateTransition:
		stage.OnAfterActorStateTransitionCreateCallback = any(callback).(OnAfterCreateInterface[ActorStateTransition])
	case *ActorStateTransitionShape:
		stage.OnAfterActorStateTransitionShapeCreateCallback = any(callback).(OnAfterCreateInterface[ActorStateTransitionShape])
	case *Analysis:
		stage.OnAfterAnalysisCreateCallback = any(callback).(OnAfterCreateInterface[Analysis])
	case *ControlPointShape:
		stage.OnAfterControlPointShapeCreateCallback = any(callback).(OnAfterCreateInterface[ControlPointShape])
	case *Diagram:
		stage.OnAfterDiagramCreateCallback = any(callback).(OnAfterCreateInterface[Diagram])
	case *Document:
		stage.OnAfterDocumentCreateCallback = any(callback).(OnAfterCreateInterface[Document])
	case *DocumentUse:
		stage.OnAfterDocumentUseCreateCallback = any(callback).(OnAfterCreateInterface[DocumentUse])
	case *EvolutionDirection:
		stage.OnAfterEvolutionDirectionCreateCallback = any(callback).(OnAfterCreateInterface[EvolutionDirection])
	case *EvolutionDirectionShape:
		stage.OnAfterEvolutionDirectionShapeCreateCallback = any(callback).(OnAfterCreateInterface[EvolutionDirectionShape])
	case *Foo:
		stage.OnAfterFooCreateCallback = any(callback).(OnAfterCreateInterface[Foo])
	case *GeoObject:
		stage.OnAfterGeoObjectCreateCallback = any(callback).(OnAfterCreateInterface[GeoObject])
	case *GeoObjectUse:
		stage.OnAfterGeoObjectUseCreateCallback = any(callback).(OnAfterCreateInterface[GeoObjectUse])
	case *Group:
		stage.OnAfterGroupCreateCallback = any(callback).(OnAfterCreateInterface[Group])
	case *GroupUse:
		stage.OnAfterGroupUseCreateCallback = any(callback).(OnAfterCreateInterface[GroupUse])
	case *Library:
		stage.OnAfterLibraryCreateCallback = any(callback).(OnAfterCreateInterface[Library])
	case *MapObject:
		stage.OnAfterMapObjectCreateCallback = any(callback).(OnAfterCreateInterface[MapObject])
	case *MapObjectUse:
		stage.OnAfterMapObjectUseCreateCallback = any(callback).(OnAfterCreateInterface[MapObjectUse])
	case *Parameter:
		stage.OnAfterParameterCreateCallback = any(callback).(OnAfterCreateInterface[Parameter])
	case *ParameterCategory:
		stage.OnAfterParameterCategoryCreateCallback = any(callback).(OnAfterCreateInterface[ParameterCategory])
	case *ParameterCategoryUse:
		stage.OnAfterParameterCategoryUseCreateCallback = any(callback).(OnAfterCreateInterface[ParameterCategoryUse])
	case *ParameterShape:
		stage.OnAfterParameterShapeCreateCallback = any(callback).(OnAfterCreateInterface[ParameterShape])
	case *ParametersAggregate:
		stage.OnAfterParametersAggregateCreateCallback = any(callback).(OnAfterCreateInterface[ParametersAggregate])
	case *ParametersAggregateShape:
		stage.OnAfterParametersAggregateShapeCreateCallback = any(callback).(OnAfterCreateInterface[ParametersAggregateShape])
	case *Position:
		stage.OnAfterPositionCreateCallback = any(callback).(OnAfterCreateInterface[Position])
	case *Repository:
		stage.OnAfterRepositoryCreateCallback = any(callback).(OnAfterCreateInterface[Repository])
	case *Scenario:
		stage.OnAfterScenarioCreateCallback = any(callback).(OnAfterCreateInterface[Scenario])
	case *User:
		stage.OnAfterUserCreateCallback = any(callback).(OnAfterCreateInterface[User])
	case *UserUse:
		stage.OnAfterUserUseCreateCallback = any(callback).(OnAfterCreateInterface[UserUse])
	case *Workspace:
		stage.OnAfterWorkspaceCreateCallback = any(callback).(OnAfterCreateInterface[Workspace])
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *ActorState:
		stage.OnAfterActorStateDeleteCallback = any(callback).(OnAfterDeleteInterface[ActorState])
	case *ActorStateShape:
		stage.OnAfterActorStateShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ActorStateShape])
	case *ActorStateTransition:
		stage.OnAfterActorStateTransitionDeleteCallback = any(callback).(OnAfterDeleteInterface[ActorStateTransition])
	case *ActorStateTransitionShape:
		stage.OnAfterActorStateTransitionShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ActorStateTransitionShape])
	case *Analysis:
		stage.OnAfterAnalysisDeleteCallback = any(callback).(OnAfterDeleteInterface[Analysis])
	case *ControlPointShape:
		stage.OnAfterControlPointShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ControlPointShape])
	case *Diagram:
		stage.OnAfterDiagramDeleteCallback = any(callback).(OnAfterDeleteInterface[Diagram])
	case *Document:
		stage.OnAfterDocumentDeleteCallback = any(callback).(OnAfterDeleteInterface[Document])
	case *DocumentUse:
		stage.OnAfterDocumentUseDeleteCallback = any(callback).(OnAfterDeleteInterface[DocumentUse])
	case *EvolutionDirection:
		stage.OnAfterEvolutionDirectionDeleteCallback = any(callback).(OnAfterDeleteInterface[EvolutionDirection])
	case *EvolutionDirectionShape:
		stage.OnAfterEvolutionDirectionShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[EvolutionDirectionShape])
	case *Foo:
		stage.OnAfterFooDeleteCallback = any(callback).(OnAfterDeleteInterface[Foo])
	case *GeoObject:
		stage.OnAfterGeoObjectDeleteCallback = any(callback).(OnAfterDeleteInterface[GeoObject])
	case *GeoObjectUse:
		stage.OnAfterGeoObjectUseDeleteCallback = any(callback).(OnAfterDeleteInterface[GeoObjectUse])
	case *Group:
		stage.OnAfterGroupDeleteCallback = any(callback).(OnAfterDeleteInterface[Group])
	case *GroupUse:
		stage.OnAfterGroupUseDeleteCallback = any(callback).(OnAfterDeleteInterface[GroupUse])
	case *Library:
		stage.OnAfterLibraryDeleteCallback = any(callback).(OnAfterDeleteInterface[Library])
	case *MapObject:
		stage.OnAfterMapObjectDeleteCallback = any(callback).(OnAfterDeleteInterface[MapObject])
	case *MapObjectUse:
		stage.OnAfterMapObjectUseDeleteCallback = any(callback).(OnAfterDeleteInterface[MapObjectUse])
	case *Parameter:
		stage.OnAfterParameterDeleteCallback = any(callback).(OnAfterDeleteInterface[Parameter])
	case *ParameterCategory:
		stage.OnAfterParameterCategoryDeleteCallback = any(callback).(OnAfterDeleteInterface[ParameterCategory])
	case *ParameterCategoryUse:
		stage.OnAfterParameterCategoryUseDeleteCallback = any(callback).(OnAfterDeleteInterface[ParameterCategoryUse])
	case *ParameterShape:
		stage.OnAfterParameterShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ParameterShape])
	case *ParametersAggregate:
		stage.OnAfterParametersAggregateDeleteCallback = any(callback).(OnAfterDeleteInterface[ParametersAggregate])
	case *ParametersAggregateShape:
		stage.OnAfterParametersAggregateShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ParametersAggregateShape])
	case *Position:
		stage.OnAfterPositionDeleteCallback = any(callback).(OnAfterDeleteInterface[Position])
	case *Repository:
		stage.OnAfterRepositoryDeleteCallback = any(callback).(OnAfterDeleteInterface[Repository])
	case *Scenario:
		stage.OnAfterScenarioDeleteCallback = any(callback).(OnAfterDeleteInterface[Scenario])
	case *User:
		stage.OnAfterUserDeleteCallback = any(callback).(OnAfterDeleteInterface[User])
	case *UserUse:
		stage.OnAfterUserUseDeleteCallback = any(callback).(OnAfterDeleteInterface[UserUse])
	case *Workspace:
		stage.OnAfterWorkspaceDeleteCallback = any(callback).(OnAfterDeleteInterface[Workspace])
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *ActorState:
		stage.OnAfterActorStateReadCallback = any(callback).(OnAfterReadInterface[ActorState])
	case *ActorStateShape:
		stage.OnAfterActorStateShapeReadCallback = any(callback).(OnAfterReadInterface[ActorStateShape])
	case *ActorStateTransition:
		stage.OnAfterActorStateTransitionReadCallback = any(callback).(OnAfterReadInterface[ActorStateTransition])
	case *ActorStateTransitionShape:
		stage.OnAfterActorStateTransitionShapeReadCallback = any(callback).(OnAfterReadInterface[ActorStateTransitionShape])
	case *Analysis:
		stage.OnAfterAnalysisReadCallback = any(callback).(OnAfterReadInterface[Analysis])
	case *ControlPointShape:
		stage.OnAfterControlPointShapeReadCallback = any(callback).(OnAfterReadInterface[ControlPointShape])
	case *Diagram:
		stage.OnAfterDiagramReadCallback = any(callback).(OnAfterReadInterface[Diagram])
	case *Document:
		stage.OnAfterDocumentReadCallback = any(callback).(OnAfterReadInterface[Document])
	case *DocumentUse:
		stage.OnAfterDocumentUseReadCallback = any(callback).(OnAfterReadInterface[DocumentUse])
	case *EvolutionDirection:
		stage.OnAfterEvolutionDirectionReadCallback = any(callback).(OnAfterReadInterface[EvolutionDirection])
	case *EvolutionDirectionShape:
		stage.OnAfterEvolutionDirectionShapeReadCallback = any(callback).(OnAfterReadInterface[EvolutionDirectionShape])
	case *Foo:
		stage.OnAfterFooReadCallback = any(callback).(OnAfterReadInterface[Foo])
	case *GeoObject:
		stage.OnAfterGeoObjectReadCallback = any(callback).(OnAfterReadInterface[GeoObject])
	case *GeoObjectUse:
		stage.OnAfterGeoObjectUseReadCallback = any(callback).(OnAfterReadInterface[GeoObjectUse])
	case *Group:
		stage.OnAfterGroupReadCallback = any(callback).(OnAfterReadInterface[Group])
	case *GroupUse:
		stage.OnAfterGroupUseReadCallback = any(callback).(OnAfterReadInterface[GroupUse])
	case *Library:
		stage.OnAfterLibraryReadCallback = any(callback).(OnAfterReadInterface[Library])
	case *MapObject:
		stage.OnAfterMapObjectReadCallback = any(callback).(OnAfterReadInterface[MapObject])
	case *MapObjectUse:
		stage.OnAfterMapObjectUseReadCallback = any(callback).(OnAfterReadInterface[MapObjectUse])
	case *Parameter:
		stage.OnAfterParameterReadCallback = any(callback).(OnAfterReadInterface[Parameter])
	case *ParameterCategory:
		stage.OnAfterParameterCategoryReadCallback = any(callback).(OnAfterReadInterface[ParameterCategory])
	case *ParameterCategoryUse:
		stage.OnAfterParameterCategoryUseReadCallback = any(callback).(OnAfterReadInterface[ParameterCategoryUse])
	case *ParameterShape:
		stage.OnAfterParameterShapeReadCallback = any(callback).(OnAfterReadInterface[ParameterShape])
	case *ParametersAggregate:
		stage.OnAfterParametersAggregateReadCallback = any(callback).(OnAfterReadInterface[ParametersAggregate])
	case *ParametersAggregateShape:
		stage.OnAfterParametersAggregateShapeReadCallback = any(callback).(OnAfterReadInterface[ParametersAggregateShape])
	case *Position:
		stage.OnAfterPositionReadCallback = any(callback).(OnAfterReadInterface[Position])
	case *Repository:
		stage.OnAfterRepositoryReadCallback = any(callback).(OnAfterReadInterface[Repository])
	case *Scenario:
		stage.OnAfterScenarioReadCallback = any(callback).(OnAfterReadInterface[Scenario])
	case *User:
		stage.OnAfterUserReadCallback = any(callback).(OnAfterReadInterface[User])
	case *UserUse:
		stage.OnAfterUserUseReadCallback = any(callback).(OnAfterReadInterface[UserUse])
	case *Workspace:
		stage.OnAfterWorkspaceReadCallback = any(callback).(OnAfterReadInterface[Workspace])
	}
}
