// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

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

// AfterCreateFromFront is a backward-compatible package-level forwarder.
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {
	stage.AfterCreateFromFront(instance)
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront[Type Gongstruct](old, new *Type) {

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

// OnAfterUpdateFromFront is a backward-compatible package-level forwarder.
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {
	stage.OnAfterUpdateFromFront(old, new)
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront[Type Gongstruct](staged, front *Type) {

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

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
