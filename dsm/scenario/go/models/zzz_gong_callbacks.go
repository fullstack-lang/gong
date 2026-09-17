// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (actorstate *ActorState) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterActorStateCreateCallback != nil {
		stage.OnAfterActorStateCreateCallback.OnAfterCreate(stage, actorstate)
	}
}

func (actorstate *ActorState) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterActorStateUpdateCallback != nil {
		var frontActorState *ActorState
		if front != nil {
			frontActorState, _ = front.(*ActorState)
		}
		stage.OnAfterActorStateUpdateCallback.OnAfterUpdate(stage, actorstate, frontActorState)
	}
}

func (actorstate *ActorState) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterActorStateDeleteCallback != nil {
		var frontActorState *ActorState
		if front != nil {
			frontActorState, _ = front.(*ActorState)
		}
		stage.OnAfterActorStateDeleteCallback.OnAfterDelete(stage, actorstate, frontActorState)
	}
}

func (actorstateshape *ActorStateShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterActorStateShapeCreateCallback != nil {
		stage.OnAfterActorStateShapeCreateCallback.OnAfterCreate(stage, actorstateshape)
	}
}

func (actorstateshape *ActorStateShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterActorStateShapeUpdateCallback != nil {
		var frontActorStateShape *ActorStateShape
		if front != nil {
			frontActorStateShape, _ = front.(*ActorStateShape)
		}
		stage.OnAfterActorStateShapeUpdateCallback.OnAfterUpdate(stage, actorstateshape, frontActorStateShape)
	}
}

func (actorstateshape *ActorStateShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterActorStateShapeDeleteCallback != nil {
		var frontActorStateShape *ActorStateShape
		if front != nil {
			frontActorStateShape, _ = front.(*ActorStateShape)
		}
		stage.OnAfterActorStateShapeDeleteCallback.OnAfterDelete(stage, actorstateshape, frontActorStateShape)
	}
}

func (actorstatetransition *ActorStateTransition) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterActorStateTransitionCreateCallback != nil {
		stage.OnAfterActorStateTransitionCreateCallback.OnAfterCreate(stage, actorstatetransition)
	}
}

func (actorstatetransition *ActorStateTransition) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterActorStateTransitionUpdateCallback != nil {
		var frontActorStateTransition *ActorStateTransition
		if front != nil {
			frontActorStateTransition, _ = front.(*ActorStateTransition)
		}
		stage.OnAfterActorStateTransitionUpdateCallback.OnAfterUpdate(stage, actorstatetransition, frontActorStateTransition)
	}
}

func (actorstatetransition *ActorStateTransition) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterActorStateTransitionDeleteCallback != nil {
		var frontActorStateTransition *ActorStateTransition
		if front != nil {
			frontActorStateTransition, _ = front.(*ActorStateTransition)
		}
		stage.OnAfterActorStateTransitionDeleteCallback.OnAfterDelete(stage, actorstatetransition, frontActorStateTransition)
	}
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterActorStateTransitionShapeCreateCallback != nil {
		stage.OnAfterActorStateTransitionShapeCreateCallback.OnAfterCreate(stage, actorstatetransitionshape)
	}
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterActorStateTransitionShapeUpdateCallback != nil {
		var frontActorStateTransitionShape *ActorStateTransitionShape
		if front != nil {
			frontActorStateTransitionShape, _ = front.(*ActorStateTransitionShape)
		}
		stage.OnAfterActorStateTransitionShapeUpdateCallback.OnAfterUpdate(stage, actorstatetransitionshape, frontActorStateTransitionShape)
	}
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterActorStateTransitionShapeDeleteCallback != nil {
		var frontActorStateTransitionShape *ActorStateTransitionShape
		if front != nil {
			frontActorStateTransitionShape, _ = front.(*ActorStateTransitionShape)
		}
		stage.OnAfterActorStateTransitionShapeDeleteCallback.OnAfterDelete(stage, actorstatetransitionshape, frontActorStateTransitionShape)
	}
}

func (analysis *Analysis) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAnalysisCreateCallback != nil {
		stage.OnAfterAnalysisCreateCallback.OnAfterCreate(stage, analysis)
	}
}

func (analysis *Analysis) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAnalysisUpdateCallback != nil {
		var frontAnalysis *Analysis
		if front != nil {
			frontAnalysis, _ = front.(*Analysis)
		}
		stage.OnAfterAnalysisUpdateCallback.OnAfterUpdate(stage, analysis, frontAnalysis)
	}
}

func (analysis *Analysis) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAnalysisDeleteCallback != nil {
		var frontAnalysis *Analysis
		if front != nil {
			frontAnalysis, _ = front.(*Analysis)
		}
		stage.OnAfterAnalysisDeleteCallback.OnAfterDelete(stage, analysis, frontAnalysis)
	}
}

func (controlpointshape *ControlPointShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterControlPointShapeCreateCallback != nil {
		stage.OnAfterControlPointShapeCreateCallback.OnAfterCreate(stage, controlpointshape)
	}
}

func (controlpointshape *ControlPointShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterControlPointShapeUpdateCallback != nil {
		var frontControlPointShape *ControlPointShape
		if front != nil {
			frontControlPointShape, _ = front.(*ControlPointShape)
		}
		stage.OnAfterControlPointShapeUpdateCallback.OnAfterUpdate(stage, controlpointshape, frontControlPointShape)
	}
}

func (controlpointshape *ControlPointShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterControlPointShapeDeleteCallback != nil {
		var frontControlPointShape *ControlPointShape
		if front != nil {
			frontControlPointShape, _ = front.(*ControlPointShape)
		}
		stage.OnAfterControlPointShapeDeleteCallback.OnAfterDelete(stage, controlpointshape, frontControlPointShape)
	}
}

func (diagram *Diagram) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDiagramCreateCallback != nil {
		stage.OnAfterDiagramCreateCallback.OnAfterCreate(stage, diagram)
	}
}

func (diagram *Diagram) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiagramUpdateCallback != nil {
		var frontDiagram *Diagram
		if front != nil {
			frontDiagram, _ = front.(*Diagram)
		}
		stage.OnAfterDiagramUpdateCallback.OnAfterUpdate(stage, diagram, frontDiagram)
	}
}

func (diagram *Diagram) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiagramDeleteCallback != nil {
		var frontDiagram *Diagram
		if front != nil {
			frontDiagram, _ = front.(*Diagram)
		}
		stage.OnAfterDiagramDeleteCallback.OnAfterDelete(stage, diagram, frontDiagram)
	}
}

func (document *Document) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDocumentCreateCallback != nil {
		stage.OnAfterDocumentCreateCallback.OnAfterCreate(stage, document)
	}
}

func (document *Document) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDocumentUpdateCallback != nil {
		var frontDocument *Document
		if front != nil {
			frontDocument, _ = front.(*Document)
		}
		stage.OnAfterDocumentUpdateCallback.OnAfterUpdate(stage, document, frontDocument)
	}
}

func (document *Document) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDocumentDeleteCallback != nil {
		var frontDocument *Document
		if front != nil {
			frontDocument, _ = front.(*Document)
		}
		stage.OnAfterDocumentDeleteCallback.OnAfterDelete(stage, document, frontDocument)
	}
}

func (documentuse *DocumentUse) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDocumentUseCreateCallback != nil {
		stage.OnAfterDocumentUseCreateCallback.OnAfterCreate(stage, documentuse)
	}
}

func (documentuse *DocumentUse) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDocumentUseUpdateCallback != nil {
		var frontDocumentUse *DocumentUse
		if front != nil {
			frontDocumentUse, _ = front.(*DocumentUse)
		}
		stage.OnAfterDocumentUseUpdateCallback.OnAfterUpdate(stage, documentuse, frontDocumentUse)
	}
}

func (documentuse *DocumentUse) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDocumentUseDeleteCallback != nil {
		var frontDocumentUse *DocumentUse
		if front != nil {
			frontDocumentUse, _ = front.(*DocumentUse)
		}
		stage.OnAfterDocumentUseDeleteCallback.OnAfterDelete(stage, documentuse, frontDocumentUse)
	}
}

func (evolutiondirection *EvolutionDirection) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEvolutionDirectionCreateCallback != nil {
		stage.OnAfterEvolutionDirectionCreateCallback.OnAfterCreate(stage, evolutiondirection)
	}
}

func (evolutiondirection *EvolutionDirection) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEvolutionDirectionUpdateCallback != nil {
		var frontEvolutionDirection *EvolutionDirection
		if front != nil {
			frontEvolutionDirection, _ = front.(*EvolutionDirection)
		}
		stage.OnAfterEvolutionDirectionUpdateCallback.OnAfterUpdate(stage, evolutiondirection, frontEvolutionDirection)
	}
}

func (evolutiondirection *EvolutionDirection) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEvolutionDirectionDeleteCallback != nil {
		var frontEvolutionDirection *EvolutionDirection
		if front != nil {
			frontEvolutionDirection, _ = front.(*EvolutionDirection)
		}
		stage.OnAfterEvolutionDirectionDeleteCallback.OnAfterDelete(stage, evolutiondirection, frontEvolutionDirection)
	}
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEvolutionDirectionShapeCreateCallback != nil {
		stage.OnAfterEvolutionDirectionShapeCreateCallback.OnAfterCreate(stage, evolutiondirectionshape)
	}
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEvolutionDirectionShapeUpdateCallback != nil {
		var frontEvolutionDirectionShape *EvolutionDirectionShape
		if front != nil {
			frontEvolutionDirectionShape, _ = front.(*EvolutionDirectionShape)
		}
		stage.OnAfterEvolutionDirectionShapeUpdateCallback.OnAfterUpdate(stage, evolutiondirectionshape, frontEvolutionDirectionShape)
	}
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEvolutionDirectionShapeDeleteCallback != nil {
		var frontEvolutionDirectionShape *EvolutionDirectionShape
		if front != nil {
			frontEvolutionDirectionShape, _ = front.(*EvolutionDirectionShape)
		}
		stage.OnAfterEvolutionDirectionShapeDeleteCallback.OnAfterDelete(stage, evolutiondirectionshape, frontEvolutionDirectionShape)
	}
}

func (foo *Foo) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFooCreateCallback != nil {
		stage.OnAfterFooCreateCallback.OnAfterCreate(stage, foo)
	}
}

func (foo *Foo) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFooUpdateCallback != nil {
		var frontFoo *Foo
		if front != nil {
			frontFoo, _ = front.(*Foo)
		}
		stage.OnAfterFooUpdateCallback.OnAfterUpdate(stage, foo, frontFoo)
	}
}

func (foo *Foo) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFooDeleteCallback != nil {
		var frontFoo *Foo
		if front != nil {
			frontFoo, _ = front.(*Foo)
		}
		stage.OnAfterFooDeleteCallback.OnAfterDelete(stage, foo, frontFoo)
	}
}

func (geoobject *GeoObject) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeoObjectCreateCallback != nil {
		stage.OnAfterGeoObjectCreateCallback.OnAfterCreate(stage, geoobject)
	}
}

func (geoobject *GeoObject) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeoObjectUpdateCallback != nil {
		var frontGeoObject *GeoObject
		if front != nil {
			frontGeoObject, _ = front.(*GeoObject)
		}
		stage.OnAfterGeoObjectUpdateCallback.OnAfterUpdate(stage, geoobject, frontGeoObject)
	}
}

func (geoobject *GeoObject) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeoObjectDeleteCallback != nil {
		var frontGeoObject *GeoObject
		if front != nil {
			frontGeoObject, _ = front.(*GeoObject)
		}
		stage.OnAfterGeoObjectDeleteCallback.OnAfterDelete(stage, geoobject, frontGeoObject)
	}
}

func (geoobjectuse *GeoObjectUse) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeoObjectUseCreateCallback != nil {
		stage.OnAfterGeoObjectUseCreateCallback.OnAfterCreate(stage, geoobjectuse)
	}
}

func (geoobjectuse *GeoObjectUse) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeoObjectUseUpdateCallback != nil {
		var frontGeoObjectUse *GeoObjectUse
		if front != nil {
			frontGeoObjectUse, _ = front.(*GeoObjectUse)
		}
		stage.OnAfterGeoObjectUseUpdateCallback.OnAfterUpdate(stage, geoobjectuse, frontGeoObjectUse)
	}
}

func (geoobjectuse *GeoObjectUse) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeoObjectUseDeleteCallback != nil {
		var frontGeoObjectUse *GeoObjectUse
		if front != nil {
			frontGeoObjectUse, _ = front.(*GeoObjectUse)
		}
		stage.OnAfterGeoObjectUseDeleteCallback.OnAfterDelete(stage, geoobjectuse, frontGeoObjectUse)
	}
}

func (group *Group) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGroupCreateCallback != nil {
		stage.OnAfterGroupCreateCallback.OnAfterCreate(stage, group)
	}
}

func (group *Group) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroupUpdateCallback != nil {
		var frontGroup *Group
		if front != nil {
			frontGroup, _ = front.(*Group)
		}
		stage.OnAfterGroupUpdateCallback.OnAfterUpdate(stage, group, frontGroup)
	}
}

func (group *Group) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroupDeleteCallback != nil {
		var frontGroup *Group
		if front != nil {
			frontGroup, _ = front.(*Group)
		}
		stage.OnAfterGroupDeleteCallback.OnAfterDelete(stage, group, frontGroup)
	}
}

func (groupuse *GroupUse) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGroupUseCreateCallback != nil {
		stage.OnAfterGroupUseCreateCallback.OnAfterCreate(stage, groupuse)
	}
}

func (groupuse *GroupUse) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroupUseUpdateCallback != nil {
		var frontGroupUse *GroupUse
		if front != nil {
			frontGroupUse, _ = front.(*GroupUse)
		}
		stage.OnAfterGroupUseUpdateCallback.OnAfterUpdate(stage, groupuse, frontGroupUse)
	}
}

func (groupuse *GroupUse) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroupUseDeleteCallback != nil {
		var frontGroupUse *GroupUse
		if front != nil {
			frontGroupUse, _ = front.(*GroupUse)
		}
		stage.OnAfterGroupUseDeleteCallback.OnAfterDelete(stage, groupuse, frontGroupUse)
	}
}

func (library *Library) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLibraryCreateCallback != nil {
		stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, library)
	}
}

func (library *Library) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLibraryUpdateCallback != nil {
		var frontLibrary *Library
		if front != nil {
			frontLibrary, _ = front.(*Library)
		}
		stage.OnAfterLibraryUpdateCallback.OnAfterUpdate(stage, library, frontLibrary)
	}
}

func (library *Library) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLibraryDeleteCallback != nil {
		var frontLibrary *Library
		if front != nil {
			frontLibrary, _ = front.(*Library)
		}
		stage.OnAfterLibraryDeleteCallback.OnAfterDelete(stage, library, frontLibrary)
	}
}

func (mapobject *MapObject) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMapObjectCreateCallback != nil {
		stage.OnAfterMapObjectCreateCallback.OnAfterCreate(stage, mapobject)
	}
}

func (mapobject *MapObject) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMapObjectUpdateCallback != nil {
		var frontMapObject *MapObject
		if front != nil {
			frontMapObject, _ = front.(*MapObject)
		}
		stage.OnAfterMapObjectUpdateCallback.OnAfterUpdate(stage, mapobject, frontMapObject)
	}
}

func (mapobject *MapObject) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMapObjectDeleteCallback != nil {
		var frontMapObject *MapObject
		if front != nil {
			frontMapObject, _ = front.(*MapObject)
		}
		stage.OnAfterMapObjectDeleteCallback.OnAfterDelete(stage, mapobject, frontMapObject)
	}
}

func (mapobjectuse *MapObjectUse) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMapObjectUseCreateCallback != nil {
		stage.OnAfterMapObjectUseCreateCallback.OnAfterCreate(stage, mapobjectuse)
	}
}

func (mapobjectuse *MapObjectUse) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMapObjectUseUpdateCallback != nil {
		var frontMapObjectUse *MapObjectUse
		if front != nil {
			frontMapObjectUse, _ = front.(*MapObjectUse)
		}
		stage.OnAfterMapObjectUseUpdateCallback.OnAfterUpdate(stage, mapobjectuse, frontMapObjectUse)
	}
}

func (mapobjectuse *MapObjectUse) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMapObjectUseDeleteCallback != nil {
		var frontMapObjectUse *MapObjectUse
		if front != nil {
			frontMapObjectUse, _ = front.(*MapObjectUse)
		}
		stage.OnAfterMapObjectUseDeleteCallback.OnAfterDelete(stage, mapobjectuse, frontMapObjectUse)
	}
}

func (parameter *Parameter) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterParameterCreateCallback != nil {
		stage.OnAfterParameterCreateCallback.OnAfterCreate(stage, parameter)
	}
}

func (parameter *Parameter) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParameterUpdateCallback != nil {
		var frontParameter *Parameter
		if front != nil {
			frontParameter, _ = front.(*Parameter)
		}
		stage.OnAfterParameterUpdateCallback.OnAfterUpdate(stage, parameter, frontParameter)
	}
}

func (parameter *Parameter) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParameterDeleteCallback != nil {
		var frontParameter *Parameter
		if front != nil {
			frontParameter, _ = front.(*Parameter)
		}
		stage.OnAfterParameterDeleteCallback.OnAfterDelete(stage, parameter, frontParameter)
	}
}

func (parametercategory *ParameterCategory) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterParameterCategoryCreateCallback != nil {
		stage.OnAfterParameterCategoryCreateCallback.OnAfterCreate(stage, parametercategory)
	}
}

func (parametercategory *ParameterCategory) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParameterCategoryUpdateCallback != nil {
		var frontParameterCategory *ParameterCategory
		if front != nil {
			frontParameterCategory, _ = front.(*ParameterCategory)
		}
		stage.OnAfterParameterCategoryUpdateCallback.OnAfterUpdate(stage, parametercategory, frontParameterCategory)
	}
}

func (parametercategory *ParameterCategory) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParameterCategoryDeleteCallback != nil {
		var frontParameterCategory *ParameterCategory
		if front != nil {
			frontParameterCategory, _ = front.(*ParameterCategory)
		}
		stage.OnAfterParameterCategoryDeleteCallback.OnAfterDelete(stage, parametercategory, frontParameterCategory)
	}
}

func (parametercategoryuse *ParameterCategoryUse) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterParameterCategoryUseCreateCallback != nil {
		stage.OnAfterParameterCategoryUseCreateCallback.OnAfterCreate(stage, parametercategoryuse)
	}
}

func (parametercategoryuse *ParameterCategoryUse) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParameterCategoryUseUpdateCallback != nil {
		var frontParameterCategoryUse *ParameterCategoryUse
		if front != nil {
			frontParameterCategoryUse, _ = front.(*ParameterCategoryUse)
		}
		stage.OnAfterParameterCategoryUseUpdateCallback.OnAfterUpdate(stage, parametercategoryuse, frontParameterCategoryUse)
	}
}

func (parametercategoryuse *ParameterCategoryUse) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParameterCategoryUseDeleteCallback != nil {
		var frontParameterCategoryUse *ParameterCategoryUse
		if front != nil {
			frontParameterCategoryUse, _ = front.(*ParameterCategoryUse)
		}
		stage.OnAfterParameterCategoryUseDeleteCallback.OnAfterDelete(stage, parametercategoryuse, frontParameterCategoryUse)
	}
}

func (parametershape *ParameterShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterParameterShapeCreateCallback != nil {
		stage.OnAfterParameterShapeCreateCallback.OnAfterCreate(stage, parametershape)
	}
}

func (parametershape *ParameterShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParameterShapeUpdateCallback != nil {
		var frontParameterShape *ParameterShape
		if front != nil {
			frontParameterShape, _ = front.(*ParameterShape)
		}
		stage.OnAfterParameterShapeUpdateCallback.OnAfterUpdate(stage, parametershape, frontParameterShape)
	}
}

func (parametershape *ParameterShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParameterShapeDeleteCallback != nil {
		var frontParameterShape *ParameterShape
		if front != nil {
			frontParameterShape, _ = front.(*ParameterShape)
		}
		stage.OnAfterParameterShapeDeleteCallback.OnAfterDelete(stage, parametershape, frontParameterShape)
	}
}

func (parametersaggregate *ParametersAggregate) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterParametersAggregateCreateCallback != nil {
		stage.OnAfterParametersAggregateCreateCallback.OnAfterCreate(stage, parametersaggregate)
	}
}

func (parametersaggregate *ParametersAggregate) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParametersAggregateUpdateCallback != nil {
		var frontParametersAggregate *ParametersAggregate
		if front != nil {
			frontParametersAggregate, _ = front.(*ParametersAggregate)
		}
		stage.OnAfterParametersAggregateUpdateCallback.OnAfterUpdate(stage, parametersaggregate, frontParametersAggregate)
	}
}

func (parametersaggregate *ParametersAggregate) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParametersAggregateDeleteCallback != nil {
		var frontParametersAggregate *ParametersAggregate
		if front != nil {
			frontParametersAggregate, _ = front.(*ParametersAggregate)
		}
		stage.OnAfterParametersAggregateDeleteCallback.OnAfterDelete(stage, parametersaggregate, frontParametersAggregate)
	}
}

func (parametersaggregateshape *ParametersAggregateShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterParametersAggregateShapeCreateCallback != nil {
		stage.OnAfterParametersAggregateShapeCreateCallback.OnAfterCreate(stage, parametersaggregateshape)
	}
}

func (parametersaggregateshape *ParametersAggregateShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParametersAggregateShapeUpdateCallback != nil {
		var frontParametersAggregateShape *ParametersAggregateShape
		if front != nil {
			frontParametersAggregateShape, _ = front.(*ParametersAggregateShape)
		}
		stage.OnAfterParametersAggregateShapeUpdateCallback.OnAfterUpdate(stage, parametersaggregateshape, frontParametersAggregateShape)
	}
}

func (parametersaggregateshape *ParametersAggregateShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParametersAggregateShapeDeleteCallback != nil {
		var frontParametersAggregateShape *ParametersAggregateShape
		if front != nil {
			frontParametersAggregateShape, _ = front.(*ParametersAggregateShape)
		}
		stage.OnAfterParametersAggregateShapeDeleteCallback.OnAfterDelete(stage, parametersaggregateshape, frontParametersAggregateShape)
	}
}

func (position *Position) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPositionCreateCallback != nil {
		stage.OnAfterPositionCreateCallback.OnAfterCreate(stage, position)
	}
}

func (position *Position) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPositionUpdateCallback != nil {
		var frontPosition *Position
		if front != nil {
			frontPosition, _ = front.(*Position)
		}
		stage.OnAfterPositionUpdateCallback.OnAfterUpdate(stage, position, frontPosition)
	}
}

func (position *Position) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPositionDeleteCallback != nil {
		var frontPosition *Position
		if front != nil {
			frontPosition, _ = front.(*Position)
		}
		stage.OnAfterPositionDeleteCallback.OnAfterDelete(stage, position, frontPosition)
	}
}

func (repository *Repository) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRepositoryCreateCallback != nil {
		stage.OnAfterRepositoryCreateCallback.OnAfterCreate(stage, repository)
	}
}

func (repository *Repository) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRepositoryUpdateCallback != nil {
		var frontRepository *Repository
		if front != nil {
			frontRepository, _ = front.(*Repository)
		}
		stage.OnAfterRepositoryUpdateCallback.OnAfterUpdate(stage, repository, frontRepository)
	}
}

func (repository *Repository) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRepositoryDeleteCallback != nil {
		var frontRepository *Repository
		if front != nil {
			frontRepository, _ = front.(*Repository)
		}
		stage.OnAfterRepositoryDeleteCallback.OnAfterDelete(stage, repository, frontRepository)
	}
}

func (scenario *Scenario) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterScenarioCreateCallback != nil {
		stage.OnAfterScenarioCreateCallback.OnAfterCreate(stage, scenario)
	}
}

func (scenario *Scenario) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScenarioUpdateCallback != nil {
		var frontScenario *Scenario
		if front != nil {
			frontScenario, _ = front.(*Scenario)
		}
		stage.OnAfterScenarioUpdateCallback.OnAfterUpdate(stage, scenario, frontScenario)
	}
}

func (scenario *Scenario) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScenarioDeleteCallback != nil {
		var frontScenario *Scenario
		if front != nil {
			frontScenario, _ = front.(*Scenario)
		}
		stage.OnAfterScenarioDeleteCallback.OnAfterDelete(stage, scenario, frontScenario)
	}
}

func (user *User) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUserCreateCallback != nil {
		stage.OnAfterUserCreateCallback.OnAfterCreate(stage, user)
	}
}

func (user *User) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUserUpdateCallback != nil {
		var frontUser *User
		if front != nil {
			frontUser, _ = front.(*User)
		}
		stage.OnAfterUserUpdateCallback.OnAfterUpdate(stage, user, frontUser)
	}
}

func (user *User) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUserDeleteCallback != nil {
		var frontUser *User
		if front != nil {
			frontUser, _ = front.(*User)
		}
		stage.OnAfterUserDeleteCallback.OnAfterDelete(stage, user, frontUser)
	}
}

func (useruse *UserUse) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUserUseCreateCallback != nil {
		stage.OnAfterUserUseCreateCallback.OnAfterCreate(stage, useruse)
	}
}

func (useruse *UserUse) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUserUseUpdateCallback != nil {
		var frontUserUse *UserUse
		if front != nil {
			frontUserUse, _ = front.(*UserUse)
		}
		stage.OnAfterUserUseUpdateCallback.OnAfterUpdate(stage, useruse, frontUserUse)
	}
}

func (useruse *UserUse) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUserUseDeleteCallback != nil {
		var frontUserUse *UserUse
		if front != nil {
			frontUserUse, _ = front.(*UserUse)
		}
		stage.OnAfterUserUseDeleteCallback.OnAfterDelete(stage, useruse, frontUserUse)
	}
}

func (workspace *Workspace) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterWorkspaceCreateCallback != nil {
		stage.OnAfterWorkspaceCreateCallback.OnAfterCreate(stage, workspace)
	}
}

func (workspace *Workspace) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWorkspaceUpdateCallback != nil {
		var frontWorkspace *Workspace
		if front != nil {
			frontWorkspace, _ = front.(*Workspace)
		}
		stage.OnAfterWorkspaceUpdateCallback.OnAfterUpdate(stage, workspace, frontWorkspace)
	}
}

func (workspace *Workspace) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWorkspaceDeleteCallback != nil {
		var frontWorkspace *Workspace
		if front != nil {
			frontWorkspace, _ = front.(*Workspace)
		}
		stage.OnAfterWorkspaceDeleteCallback.OnAfterDelete(stage, workspace, frontWorkspace)
	}
}

