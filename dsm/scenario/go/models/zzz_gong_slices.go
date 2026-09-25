// generated code - do not edit
package models

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct ActorStateTransition
	// insertion point per field
	stage.ActorStateTransition_Justifications_reverseMap = make(map[*Parameter]*ActorStateTransition)
	for actorstatetransition := range stage.ActorStateTransitions {
		_ = actorstatetransition
		for _, _parameter := range actorstatetransition.Justifications {
			stage.ActorStateTransition_Justifications_reverseMap[_parameter] = actorstatetransition
		}
	}

	// Compute reverse map for named struct ActorStateTransitionShape
	// insertion point per field
	stage.ActorStateTransitionShape_ControlPointShapes_reverseMap = make(map[*ControlPointShape]*ActorStateTransitionShape)
	for actorstatetransitionshape := range stage.ActorStateTransitionShapes {
		_ = actorstatetransitionshape
		for _, _controlpointshape := range actorstatetransitionshape.ControlPointShapes {
			stage.ActorStateTransitionShape_ControlPointShapes_reverseMap[_controlpointshape] = actorstatetransitionshape
		}
	}

	// Compute reverse map for named struct Analysis
	// insertion point per field
	stage.Analysis_Scenarios_reverseMap = make(map[*Scenario]*Analysis)
	for analysis := range stage.Analysiss {
		_ = analysis
		for _, _scenario := range analysis.Scenarios {
			stage.Analysis_Scenarios_reverseMap[_scenario] = analysis
		}
	}
	stage.Analysis_GroupUse_reverseMap = make(map[*GroupUse]*Analysis)
	for analysis := range stage.Analysiss {
		_ = analysis
		for _, _groupuse := range analysis.GroupUse {
			stage.Analysis_GroupUse_reverseMap[_groupuse] = analysis
		}
	}
	stage.Analysis_GeoObjectUse_reverseMap = make(map[*GeoObjectUse]*Analysis)
	for analysis := range stage.Analysiss {
		_ = analysis
		for _, _geoobjectuse := range analysis.GeoObjectUse {
			stage.Analysis_GeoObjectUse_reverseMap[_geoobjectuse] = analysis
		}
	}
	stage.Analysis_MapUse_reverseMap = make(map[*MapObjectUse]*Analysis)
	for analysis := range stage.Analysiss {
		_ = analysis
		for _, _mapobjectuse := range analysis.MapUse {
			stage.Analysis_MapUse_reverseMap[_mapobjectuse] = analysis
		}
	}

	// Compute reverse map for named struct Diagram
	// insertion point per field
	stage.Diagram_EvolutionDirectionShapes_reverseMap = make(map[*EvolutionDirectionShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _evolutiondirectionshape := range diagram.EvolutionDirectionShapes {
			stage.Diagram_EvolutionDirectionShapes_reverseMap[_evolutiondirectionshape] = diagram
		}
	}
	stage.Diagram_EvolutionDirectionsWhoseNodeIsExpanded_reverseMap = make(map[*EvolutionDirection]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _evolutiondirection := range diagram.EvolutionDirectionsWhoseNodeIsExpanded {
			stage.Diagram_EvolutionDirectionsWhoseNodeIsExpanded_reverseMap[_evolutiondirection] = diagram
		}
	}
	stage.Diagram_ActorStateShapes_reverseMap = make(map[*ActorStateShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _actorstateshape := range diagram.ActorStateShapes {
			stage.Diagram_ActorStateShapes_reverseMap[_actorstateshape] = diagram
		}
	}
	stage.Diagram_ActorStatesWhoseNodeIsExpanded_reverseMap = make(map[*ActorState]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _actorstate := range diagram.ActorStatesWhoseNodeIsExpanded {
			stage.Diagram_ActorStatesWhoseNodeIsExpanded_reverseMap[_actorstate] = diagram
		}
	}
	stage.Diagram_ParameterShapes_reverseMap = make(map[*ParameterShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _parametershape := range diagram.ParameterShapes {
			stage.Diagram_ParameterShapes_reverseMap[_parametershape] = diagram
		}
	}
	stage.Diagram_ParametersWhoseNodeIsExpanded_reverseMap = make(map[*Parameter]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _parameter := range diagram.ParametersWhoseNodeIsExpanded {
			stage.Diagram_ParametersWhoseNodeIsExpanded_reverseMap[_parameter] = diagram
		}
	}
	stage.Diagram_ScenarioParameterShapes_reverseMap = make(map[*ParametersAggregateShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _parametersaggregateshape := range diagram.ScenarioParameterShapes {
			stage.Diagram_ScenarioParameterShapes_reverseMap[_parametersaggregateshape] = diagram
		}
	}
	stage.Diagram_ParametersAggregatesWhoseNodeIsExpanded_reverseMap = make(map[*ParametersAggregate]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _parametersaggregate := range diagram.ParametersAggregatesWhoseNodeIsExpanded {
			stage.Diagram_ParametersAggregatesWhoseNodeIsExpanded_reverseMap[_parametersaggregate] = diagram
		}
	}
	stage.Diagram_ActorStateTransitionShapes_reverseMap = make(map[*ActorStateTransitionShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _actorstatetransitionshape := range diagram.ActorStateTransitionShapes {
			stage.Diagram_ActorStateTransitionShapes_reverseMap[_actorstatetransitionshape] = diagram
		}
	}
	stage.Diagram_ActorStateTransitionsWhoseNodeIsExpanded_reverseMap = make(map[*ActorStateTransition]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _actorstatetransition := range diagram.ActorStateTransitionsWhoseNodeIsExpanded {
			stage.Diagram_ActorStateTransitionsWhoseNodeIsExpanded_reverseMap[_actorstatetransition] = diagram
		}
	}

	// Compute reverse map for named struct Document
	// insertion point per field
	stage.Document_GeoObjectUse_reverseMap = make(map[*GeoObjectUse]*Document)
	for document := range stage.Documents {
		_ = document
		for _, _geoobjectuse := range document.GeoObjectUse {
			stage.Document_GeoObjectUse_reverseMap[_geoobjectuse] = document
		}
	}

	// Compute reverse map for named struct Group
	// insertion point per field
	stage.Group_UserUse_reverseMap = make(map[*UserUse]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _useruse := range group.UserUse {
			stage.Group_UserUse_reverseMap[_useruse] = group
		}
	}

	// Compute reverse map for named struct Library
	// insertion point per field
	stage.Library_Analyses_reverseMap = make(map[*Analysis]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _analysis := range library.Analyses {
			stage.Library_Analyses_reverseMap[_analysis] = library
		}
	}
	stage.Library_SubLibraries_reverseMap = make(map[*Library]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _library := range library.SubLibraries {
			stage.Library_SubLibraries_reverseMap[_library] = library
		}
	}
	stage.Library_SubLibrariesWhoseNodeIsExpanded_reverseMap = make(map[*Library]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
			stage.Library_SubLibrariesWhoseNodeIsExpanded_reverseMap[_library] = library
		}
	}

	// Compute reverse map for named struct Parameter
	// insertion point per field
	stage.Parameter_GroupUse_reverseMap = make(map[*GroupUse]*Parameter)
	for parameter := range stage.Parameters {
		_ = parameter
		for _, _groupuse := range parameter.GroupUse {
			stage.Parameter_GroupUse_reverseMap[_groupuse] = parameter
		}
	}
	stage.Parameter_DocumentUse_reverseMap = make(map[*DocumentUse]*Parameter)
	for parameter := range stage.Parameters {
		_ = parameter
		for _, _documentuse := range parameter.DocumentUse {
			stage.Parameter_DocumentUse_reverseMap[_documentuse] = parameter
		}
	}
	stage.Parameter_GeoObjectUse_reverseMap = make(map[*GeoObjectUse]*Parameter)
	for parameter := range stage.Parameters {
		_ = parameter
		for _, _geoobjectuse := range parameter.GeoObjectUse {
			stage.Parameter_GeoObjectUse_reverseMap[_geoobjectuse] = parameter
		}
	}

	// Compute reverse map for named struct ParameterCategory
	// insertion point per field
	stage.ParameterCategory_ParameterUse_reverseMap = make(map[*ParameterShape]*ParameterCategory)
	for parametercategory := range stage.ParameterCategorys {
		_ = parametercategory
		for _, _parametershape := range parametercategory.ParameterUse {
			stage.ParameterCategory_ParameterUse_reverseMap[_parametershape] = parametercategory
		}
	}

	// Compute reverse map for named struct ParametersAggregate
	// insertion point per field
	stage.ParametersAggregate_Parameters_reverseMap = make(map[*Parameter]*ParametersAggregate)
	for parametersaggregate := range stage.ParametersAggregates {
		_ = parametersaggregate
		for _, _parameter := range parametersaggregate.Parameters {
			stage.ParametersAggregate_Parameters_reverseMap[_parameter] = parametersaggregate
		}
	}

	// Compute reverse map for named struct Repository
	// insertion point per field
	stage.Repository_ParameterUse_reverseMap = make(map[*ParameterShape]*Repository)
	for repository := range stage.Repositorys {
		_ = repository
		for _, _parametershape := range repository.ParameterUse {
			stage.Repository_ParameterUse_reverseMap[_parametershape] = repository
		}
	}
	stage.Repository_GroupUse_reverseMap = make(map[*GroupUse]*Repository)
	for repository := range stage.Repositorys {
		_ = repository
		for _, _groupuse := range repository.GroupUse {
			stage.Repository_GroupUse_reverseMap[_groupuse] = repository
		}
	}

	// Compute reverse map for named struct Scenario
	// insertion point per field
	stage.Scenario_Diagrams_reverseMap = make(map[*Diagram]*Scenario)
	for scenario := range stage.Scenarios {
		_ = scenario
		for _, _diagram := range scenario.Diagrams {
			stage.Scenario_Diagrams_reverseMap[_diagram] = scenario
		}
	}
	stage.Scenario_ActorStates_reverseMap = make(map[*ActorState]*Scenario)
	for scenario := range stage.Scenarios {
		_ = scenario
		for _, _actorstate := range scenario.ActorStates {
			stage.Scenario_ActorStates_reverseMap[_actorstate] = scenario
		}
	}
	stage.Scenario_ActorStateTransitions_reverseMap = make(map[*ActorStateTransition]*Scenario)
	for scenario := range stage.Scenarios {
		_ = scenario
		for _, _actorstatetransition := range scenario.ActorStateTransitions {
			stage.Scenario_ActorStateTransitions_reverseMap[_actorstatetransition] = scenario
		}
	}
	stage.Scenario_EvolutionDirections_reverseMap = make(map[*EvolutionDirection]*Scenario)
	for scenario := range stage.Scenarios {
		_ = scenario
		for _, _evolutiondirection := range scenario.EvolutionDirections {
			stage.Scenario_EvolutionDirections_reverseMap[_evolutiondirection] = scenario
		}
	}
	stage.Scenario_Parameters_reverseMap = make(map[*Parameter]*Scenario)
	for scenario := range stage.Scenarios {
		_ = scenario
		for _, _parameter := range scenario.Parameters {
			stage.Scenario_Parameters_reverseMap[_parameter] = scenario
		}
	}
	stage.Scenario_ParametersAggretates_reverseMap = make(map[*ParametersAggregate]*Scenario)
	for scenario := range stage.Scenarios {
		_ = scenario
		for _, _parametersaggregate := range scenario.ParametersAggretates {
			stage.Scenario_ParametersAggretates_reverseMap[_parametersaggregate] = scenario
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.ActorStates)

	res = __gong__appendInstances(res, stage.ActorStateShapes)

	res = __gong__appendInstances(res, stage.ActorStateTransitions)

	res = __gong__appendInstances(res, stage.ActorStateTransitionShapes)

	res = __gong__appendInstances(res, stage.Analysiss)

	res = __gong__appendInstances(res, stage.ControlPointShapes)

	res = __gong__appendInstances(res, stage.Diagrams)

	res = __gong__appendInstances(res, stage.Documents)

	res = __gong__appendInstances(res, stage.DocumentUses)

	res = __gong__appendInstances(res, stage.EvolutionDirections)

	res = __gong__appendInstances(res, stage.EvolutionDirectionShapes)

	res = __gong__appendInstances(res, stage.Foos)

	res = __gong__appendInstances(res, stage.GeoObjects)

	res = __gong__appendInstances(res, stage.GeoObjectUses)

	res = __gong__appendInstances(res, stage.Groups)

	res = __gong__appendInstances(res, stage.GroupUses)

	res = __gong__appendInstances(res, stage.Librarys)

	res = __gong__appendInstances(res, stage.MapObjects)

	res = __gong__appendInstances(res, stage.MapObjectUses)

	res = __gong__appendInstances(res, stage.Parameters)

	res = __gong__appendInstances(res, stage.ParameterCategorys)

	res = __gong__appendInstances(res, stage.ParameterCategoryUses)

	res = __gong__appendInstances(res, stage.ParameterShapes)

	res = __gong__appendInstances(res, stage.ParametersAggregates)

	res = __gong__appendInstances(res, stage.ParametersAggregateShapes)

	res = __gong__appendInstances(res, stage.Positions)

	res = __gong__appendInstances(res, stage.Repositorys)

	res = __gong__appendInstances(res, stage.Scenarios)

	res = __gong__appendInstances(res, stage.Users)

	res = __gong__appendInstances(res, stage.UserUses)

	res = __gong__appendInstances(res, stage.Workspaces)

	return
}

// insertion point per named struct
func (actorstate *ActorState) GongCopy() GongstructIF {
	newInstance := new(ActorState)
	actorstate.GongCopyBasicFields(newInstance)
	return newInstance
}

func (actorstateshape *ActorStateShape) GongCopy() GongstructIF {
	newInstance := new(ActorStateShape)
	actorstateshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (actorstatetransition *ActorStateTransition) GongCopy() GongstructIF {
	newInstance := new(ActorStateTransition)
	actorstatetransition.GongCopyBasicFields(newInstance)
	return newInstance
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongCopy() GongstructIF {
	newInstance := new(ActorStateTransitionShape)
	actorstatetransitionshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (analysis *Analysis) GongCopy() GongstructIF {
	newInstance := new(Analysis)
	analysis.GongCopyBasicFields(newInstance)
	return newInstance
}

func (controlpointshape *ControlPointShape) GongCopy() GongstructIF {
	newInstance := new(ControlPointShape)
	controlpointshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (diagram *Diagram) GongCopy() GongstructIF {
	newInstance := new(Diagram)
	diagram.GongCopyBasicFields(newInstance)
	return newInstance
}

func (document *Document) GongCopy() GongstructIF {
	newInstance := new(Document)
	document.GongCopyBasicFields(newInstance)
	return newInstance
}

func (documentuse *DocumentUse) GongCopy() GongstructIF {
	newInstance := new(DocumentUse)
	documentuse.GongCopyBasicFields(newInstance)
	return newInstance
}

func (evolutiondirection *EvolutionDirection) GongCopy() GongstructIF {
	newInstance := new(EvolutionDirection)
	evolutiondirection.GongCopyBasicFields(newInstance)
	return newInstance
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongCopy() GongstructIF {
	newInstance := new(EvolutionDirectionShape)
	evolutiondirectionshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (foo *Foo) GongCopy() GongstructIF {
	newInstance := new(Foo)
	foo.GongCopyBasicFields(newInstance)
	return newInstance
}

func (geoobject *GeoObject) GongCopy() GongstructIF {
	newInstance := new(GeoObject)
	geoobject.GongCopyBasicFields(newInstance)
	return newInstance
}

func (geoobjectuse *GeoObjectUse) GongCopy() GongstructIF {
	newInstance := new(GeoObjectUse)
	geoobjectuse.GongCopyBasicFields(newInstance)
	return newInstance
}

func (group *Group) GongCopy() GongstructIF {
	newInstance := new(Group)
	group.GongCopyBasicFields(newInstance)
	return newInstance
}

func (groupuse *GroupUse) GongCopy() GongstructIF {
	newInstance := new(GroupUse)
	groupuse.GongCopyBasicFields(newInstance)
	return newInstance
}

func (library *Library) GongCopy() GongstructIF {
	newInstance := new(Library)
	library.GongCopyBasicFields(newInstance)
	return newInstance
}

func (mapobject *MapObject) GongCopy() GongstructIF {
	newInstance := new(MapObject)
	mapobject.GongCopyBasicFields(newInstance)
	return newInstance
}

func (mapobjectuse *MapObjectUse) GongCopy() GongstructIF {
	newInstance := new(MapObjectUse)
	mapobjectuse.GongCopyBasicFields(newInstance)
	return newInstance
}

func (parameter *Parameter) GongCopy() GongstructIF {
	newInstance := new(Parameter)
	parameter.GongCopyBasicFields(newInstance)
	return newInstance
}

func (parametercategory *ParameterCategory) GongCopy() GongstructIF {
	newInstance := new(ParameterCategory)
	parametercategory.GongCopyBasicFields(newInstance)
	return newInstance
}

func (parametercategoryuse *ParameterCategoryUse) GongCopy() GongstructIF {
	newInstance := new(ParameterCategoryUse)
	parametercategoryuse.GongCopyBasicFields(newInstance)
	return newInstance
}

func (parametershape *ParameterShape) GongCopy() GongstructIF {
	newInstance := new(ParameterShape)
	parametershape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (parametersaggregate *ParametersAggregate) GongCopy() GongstructIF {
	newInstance := new(ParametersAggregate)
	parametersaggregate.GongCopyBasicFields(newInstance)
	return newInstance
}

func (parametersaggregateshape *ParametersAggregateShape) GongCopy() GongstructIF {
	newInstance := new(ParametersAggregateShape)
	parametersaggregateshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (position *Position) GongCopy() GongstructIF {
	newInstance := new(Position)
	position.GongCopyBasicFields(newInstance)
	return newInstance
}

func (repository *Repository) GongCopy() GongstructIF {
	newInstance := new(Repository)
	repository.GongCopyBasicFields(newInstance)
	return newInstance
}

func (scenario *Scenario) GongCopy() GongstructIF {
	newInstance := new(Scenario)
	scenario.GongCopyBasicFields(newInstance)
	return newInstance
}

func (user *User) GongCopy() GongstructIF {
	newInstance := new(User)
	user.GongCopyBasicFields(newInstance)
	return newInstance
}

func (useruse *UserUse) GongCopy() GongstructIF {
	newInstance := new(UserUse)
	useruse.GongCopyBasicFields(newInstance)
	return newInstance
}

func (workspace *Workspace) GongCopy() GongstructIF {
	newInstance := new(Workspace)
	workspace.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (actorstate *ActorState) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, actorstate)
}

func (actorstateshape *ActorStateShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, actorstateshape)
}

func (actorstatetransition *ActorStateTransition) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, actorstatetransition)
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, actorstatetransitionshape)
}

func (analysis *Analysis) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, analysis)
}

func (controlpointshape *ControlPointShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, controlpointshape)
}

func (diagram *Diagram) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, diagram)
}

func (document *Document) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, document)
}

func (documentuse *DocumentUse) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, documentuse)
}

func (evolutiondirection *EvolutionDirection) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, evolutiondirection)
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, evolutiondirectionshape)
}

func (foo *Foo) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, foo)
}

func (geoobject *GeoObject) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, geoobject)
}

func (geoobjectuse *GeoObjectUse) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, geoobjectuse)
}

func (group *Group) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, group)
}

func (groupuse *GroupUse) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, groupuse)
}

func (library *Library) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, library)
}

func (mapobject *MapObject) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, mapobject)
}

func (mapobjectuse *MapObjectUse) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, mapobjectuse)
}

func (parameter *Parameter) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, parameter)
}

func (parametercategory *ParameterCategory) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, parametercategory)
}

func (parametercategoryuse *ParameterCategoryUse) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, parametercategoryuse)
}

func (parametershape *ParameterShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, parametershape)
}

func (parametersaggregate *ParametersAggregate) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, parametersaggregate)
}

func (parametersaggregateshape *ParametersAggregateShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, parametersaggregateshape)
}

func (position *Position) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, position)
}

func (repository *Repository) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, repository)
}

func (scenario *Scenario) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, scenario)
}

func (user *User) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, user)
}

func (useruse *UserUse) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, useruse)
}

func (workspace *Workspace) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, workspace)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	computeCommitsForType(
		stage,
		stage.ActorStates,
		stage.ActorState_stagedOrder,
		stage.ActorStates_reference,
		&stage.ActorStates_referenceOrder,
		stage.ActorStates_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ActorStateShapes,
		stage.ActorStateShape_stagedOrder,
		stage.ActorStateShapes_reference,
		&stage.ActorStateShapes_referenceOrder,
		stage.ActorStateShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ActorStateTransitions,
		stage.ActorStateTransition_stagedOrder,
		stage.ActorStateTransitions_reference,
		&stage.ActorStateTransitions_referenceOrder,
		stage.ActorStateTransitions_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ActorStateTransitionShapes,
		stage.ActorStateTransitionShape_stagedOrder,
		stage.ActorStateTransitionShapes_reference,
		&stage.ActorStateTransitionShapes_referenceOrder,
		stage.ActorStateTransitionShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Analysiss,
		stage.Analysis_stagedOrder,
		stage.Analysiss_reference,
		&stage.Analysiss_referenceOrder,
		stage.Analysiss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ControlPointShapes,
		stage.ControlPointShape_stagedOrder,
		stage.ControlPointShapes_reference,
		&stage.ControlPointShapes_referenceOrder,
		stage.ControlPointShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Diagrams,
		stage.Diagram_stagedOrder,
		stage.Diagrams_reference,
		&stage.Diagrams_referenceOrder,
		stage.Diagrams_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Documents,
		stage.Document_stagedOrder,
		stage.Documents_reference,
		&stage.Documents_referenceOrder,
		stage.Documents_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DocumentUses,
		stage.DocumentUse_stagedOrder,
		stage.DocumentUses_reference,
		&stage.DocumentUses_referenceOrder,
		stage.DocumentUses_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.EvolutionDirections,
		stage.EvolutionDirection_stagedOrder,
		stage.EvolutionDirections_reference,
		&stage.EvolutionDirections_referenceOrder,
		stage.EvolutionDirections_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.EvolutionDirectionShapes,
		stage.EvolutionDirectionShape_stagedOrder,
		stage.EvolutionDirectionShapes_reference,
		&stage.EvolutionDirectionShapes_referenceOrder,
		stage.EvolutionDirectionShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Foos,
		stage.Foo_stagedOrder,
		stage.Foos_reference,
		&stage.Foos_referenceOrder,
		stage.Foos_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.GeoObjects,
		stage.GeoObject_stagedOrder,
		stage.GeoObjects_reference,
		&stage.GeoObjects_referenceOrder,
		stage.GeoObjects_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.GeoObjectUses,
		stage.GeoObjectUse_stagedOrder,
		stage.GeoObjectUses_reference,
		&stage.GeoObjectUses_referenceOrder,
		stage.GeoObjectUses_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Groups,
		stage.Group_stagedOrder,
		stage.Groups_reference,
		&stage.Groups_referenceOrder,
		stage.Groups_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.GroupUses,
		stage.GroupUse_stagedOrder,
		stage.GroupUses_reference,
		&stage.GroupUses_referenceOrder,
		stage.GroupUses_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Librarys,
		stage.Library_stagedOrder,
		stage.Librarys_reference,
		&stage.Librarys_referenceOrder,
		stage.Librarys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MapObjects,
		stage.MapObject_stagedOrder,
		stage.MapObjects_reference,
		&stage.MapObjects_referenceOrder,
		stage.MapObjects_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MapObjectUses,
		stage.MapObjectUse_stagedOrder,
		stage.MapObjectUses_reference,
		&stage.MapObjectUses_referenceOrder,
		stage.MapObjectUses_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Parameters,
		stage.Parameter_stagedOrder,
		stage.Parameters_reference,
		&stage.Parameters_referenceOrder,
		stage.Parameters_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ParameterCategorys,
		stage.ParameterCategory_stagedOrder,
		stage.ParameterCategorys_reference,
		&stage.ParameterCategorys_referenceOrder,
		stage.ParameterCategorys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ParameterCategoryUses,
		stage.ParameterCategoryUse_stagedOrder,
		stage.ParameterCategoryUses_reference,
		&stage.ParameterCategoryUses_referenceOrder,
		stage.ParameterCategoryUses_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ParameterShapes,
		stage.ParameterShape_stagedOrder,
		stage.ParameterShapes_reference,
		&stage.ParameterShapes_referenceOrder,
		stage.ParameterShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ParametersAggregates,
		stage.ParametersAggregate_stagedOrder,
		stage.ParametersAggregates_reference,
		&stage.ParametersAggregates_referenceOrder,
		stage.ParametersAggregates_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ParametersAggregateShapes,
		stage.ParametersAggregateShape_stagedOrder,
		stage.ParametersAggregateShapes_reference,
		&stage.ParametersAggregateShapes_referenceOrder,
		stage.ParametersAggregateShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Positions,
		stage.Position_stagedOrder,
		stage.Positions_reference,
		&stage.Positions_referenceOrder,
		stage.Positions_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Repositorys,
		stage.Repository_stagedOrder,
		stage.Repositorys_reference,
		&stage.Repositorys_referenceOrder,
		stage.Repositorys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Scenarios,
		stage.Scenario_stagedOrder,
		stage.Scenarios_reference,
		&stage.Scenarios_referenceOrder,
		stage.Scenarios_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Users,
		stage.User_stagedOrder,
		stage.Users_reference,
		&stage.Users_referenceOrder,
		stage.Users_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.UserUses,
		stage.UserUse_stagedOrder,
		stage.UserUses_reference,
		&stage.UserUses_referenceOrder,
		stage.UserUses_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Workspaces,
		stage.Workspace_stagedOrder,
		stage.Workspaces_reference,
		&stage.Workspaces_referenceOrder,
		stage.Workspaces_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)
		stage.modified = true
	} else {
		stage.modified = false
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.ActorStates, &stage.ActorStates_reference, &stage.ActorStates_referenceOrder, &stage.ActorStates_instance)

	__gong__computeReferencePass1(stage, stage.ActorStateShapes, &stage.ActorStateShapes_reference, &stage.ActorStateShapes_referenceOrder, &stage.ActorStateShapes_instance)

	__gong__computeReferencePass1(stage, stage.ActorStateTransitions, &stage.ActorStateTransitions_reference, &stage.ActorStateTransitions_referenceOrder, &stage.ActorStateTransitions_instance)

	__gong__computeReferencePass1(stage, stage.ActorStateTransitionShapes, &stage.ActorStateTransitionShapes_reference, &stage.ActorStateTransitionShapes_referenceOrder, &stage.ActorStateTransitionShapes_instance)

	__gong__computeReferencePass1(stage, stage.Analysiss, &stage.Analysiss_reference, &stage.Analysiss_referenceOrder, &stage.Analysiss_instance)

	__gong__computeReferencePass1(stage, stage.ControlPointShapes, &stage.ControlPointShapes_reference, &stage.ControlPointShapes_referenceOrder, &stage.ControlPointShapes_instance)

	__gong__computeReferencePass1(stage, stage.Diagrams, &stage.Diagrams_reference, &stage.Diagrams_referenceOrder, &stage.Diagrams_instance)

	__gong__computeReferencePass1(stage, stage.Documents, &stage.Documents_reference, &stage.Documents_referenceOrder, &stage.Documents_instance)

	__gong__computeReferencePass1(stage, stage.DocumentUses, &stage.DocumentUses_reference, &stage.DocumentUses_referenceOrder, &stage.DocumentUses_instance)

	__gong__computeReferencePass1(stage, stage.EvolutionDirections, &stage.EvolutionDirections_reference, &stage.EvolutionDirections_referenceOrder, &stage.EvolutionDirections_instance)

	__gong__computeReferencePass1(stage, stage.EvolutionDirectionShapes, &stage.EvolutionDirectionShapes_reference, &stage.EvolutionDirectionShapes_referenceOrder, &stage.EvolutionDirectionShapes_instance)

	__gong__computeReferencePass1(stage, stage.Foos, &stage.Foos_reference, &stage.Foos_referenceOrder, &stage.Foos_instance)

	__gong__computeReferencePass1(stage, stage.GeoObjects, &stage.GeoObjects_reference, &stage.GeoObjects_referenceOrder, &stage.GeoObjects_instance)

	__gong__computeReferencePass1(stage, stage.GeoObjectUses, &stage.GeoObjectUses_reference, &stage.GeoObjectUses_referenceOrder, &stage.GeoObjectUses_instance)

	__gong__computeReferencePass1(stage, stage.Groups, &stage.Groups_reference, &stage.Groups_referenceOrder, &stage.Groups_instance)

	__gong__computeReferencePass1(stage, stage.GroupUses, &stage.GroupUses_reference, &stage.GroupUses_referenceOrder, &stage.GroupUses_instance)

	__gong__computeReferencePass1(stage, stage.Librarys, &stage.Librarys_reference, &stage.Librarys_referenceOrder, &stage.Librarys_instance)

	__gong__computeReferencePass1(stage, stage.MapObjects, &stage.MapObjects_reference, &stage.MapObjects_referenceOrder, &stage.MapObjects_instance)

	__gong__computeReferencePass1(stage, stage.MapObjectUses, &stage.MapObjectUses_reference, &stage.MapObjectUses_referenceOrder, &stage.MapObjectUses_instance)

	__gong__computeReferencePass1(stage, stage.Parameters, &stage.Parameters_reference, &stage.Parameters_referenceOrder, &stage.Parameters_instance)

	__gong__computeReferencePass1(stage, stage.ParameterCategorys, &stage.ParameterCategorys_reference, &stage.ParameterCategorys_referenceOrder, &stage.ParameterCategorys_instance)

	__gong__computeReferencePass1(stage, stage.ParameterCategoryUses, &stage.ParameterCategoryUses_reference, &stage.ParameterCategoryUses_referenceOrder, &stage.ParameterCategoryUses_instance)

	__gong__computeReferencePass1(stage, stage.ParameterShapes, &stage.ParameterShapes_reference, &stage.ParameterShapes_referenceOrder, &stage.ParameterShapes_instance)

	__gong__computeReferencePass1(stage, stage.ParametersAggregates, &stage.ParametersAggregates_reference, &stage.ParametersAggregates_referenceOrder, &stage.ParametersAggregates_instance)

	__gong__computeReferencePass1(stage, stage.ParametersAggregateShapes, &stage.ParametersAggregateShapes_reference, &stage.ParametersAggregateShapes_referenceOrder, &stage.ParametersAggregateShapes_instance)

	__gong__computeReferencePass1(stage, stage.Positions, &stage.Positions_reference, &stage.Positions_referenceOrder, &stage.Positions_instance)

	__gong__computeReferencePass1(stage, stage.Repositorys, &stage.Repositorys_reference, &stage.Repositorys_referenceOrder, &stage.Repositorys_instance)

	__gong__computeReferencePass1(stage, stage.Scenarios, &stage.Scenarios_reference, &stage.Scenarios_referenceOrder, &stage.Scenarios_instance)

	__gong__computeReferencePass1(stage, stage.Users, &stage.Users_reference, &stage.Users_referenceOrder, &stage.Users_instance)

	__gong__computeReferencePass1(stage, stage.UserUses, &stage.UserUses_reference, &stage.UserUses_referenceOrder, &stage.UserUses_instance)

	__gong__computeReferencePass1(stage, stage.Workspaces, &stage.Workspaces_reference, &stage.Workspaces_referenceOrder, &stage.Workspaces_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.ActorStates, stage.ActorStates_reference, stage)

	__gong__computeReferencePass2(stage.ActorStateShapes, stage.ActorStateShapes_reference, stage)

	__gong__computeReferencePass2(stage.ActorStateTransitions, stage.ActorStateTransitions_reference, stage)

	__gong__computeReferencePass2(stage.ActorStateTransitionShapes, stage.ActorStateTransitionShapes_reference, stage)

	__gong__computeReferencePass2(stage.Analysiss, stage.Analysiss_reference, stage)

	__gong__computeReferencePass2(stage.ControlPointShapes, stage.ControlPointShapes_reference, stage)

	__gong__computeReferencePass2(stage.Diagrams, stage.Diagrams_reference, stage)

	__gong__computeReferencePass2(stage.Documents, stage.Documents_reference, stage)

	__gong__computeReferencePass2(stage.DocumentUses, stage.DocumentUses_reference, stage)

	__gong__computeReferencePass2(stage.EvolutionDirections, stage.EvolutionDirections_reference, stage)

	__gong__computeReferencePass2(stage.EvolutionDirectionShapes, stage.EvolutionDirectionShapes_reference, stage)

	__gong__computeReferencePass2(stage.Foos, stage.Foos_reference, stage)

	__gong__computeReferencePass2(stage.GeoObjects, stage.GeoObjects_reference, stage)

	__gong__computeReferencePass2(stage.GeoObjectUses, stage.GeoObjectUses_reference, stage)

	__gong__computeReferencePass2(stage.Groups, stage.Groups_reference, stage)

	__gong__computeReferencePass2(stage.GroupUses, stage.GroupUses_reference, stage)

	__gong__computeReferencePass2(stage.Librarys, stage.Librarys_reference, stage)

	__gong__computeReferencePass2(stage.MapObjects, stage.MapObjects_reference, stage)

	__gong__computeReferencePass2(stage.MapObjectUses, stage.MapObjectUses_reference, stage)

	__gong__computeReferencePass2(stage.Parameters, stage.Parameters_reference, stage)

	__gong__computeReferencePass2(stage.ParameterCategorys, stage.ParameterCategorys_reference, stage)

	__gong__computeReferencePass2(stage.ParameterCategoryUses, stage.ParameterCategoryUses_reference, stage)

	__gong__computeReferencePass2(stage.ParameterShapes, stage.ParameterShapes_reference, stage)

	__gong__computeReferencePass2(stage.ParametersAggregates, stage.ParametersAggregates_reference, stage)

	__gong__computeReferencePass2(stage.ParametersAggregateShapes, stage.ParametersAggregateShapes_reference, stage)

	__gong__computeReferencePass2(stage.Positions, stage.Positions_reference, stage)

	__gong__computeReferencePass2(stage.Repositorys, stage.Repositorys_reference, stage)

	__gong__computeReferencePass2(stage.Scenarios, stage.Scenarios_reference, stage)

	__gong__computeReferencePass2(stage.Users, stage.Users_reference, stage)

	__gong__computeReferencePass2(stage.UserUses, stage.UserUses_reference, stage)

	__gong__computeReferencePass2(stage.Workspaces, stage.Workspaces_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (actorstate *ActorState) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ActorState_stagedOrder, stage.ActorStates_referenceOrder, actorstate, "ActorState")
}

func (actorstateshape *ActorStateShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ActorStateShape_stagedOrder, stage.ActorStateShapes_referenceOrder, actorstateshape, "ActorStateShape")
}

func (actorstatetransition *ActorStateTransition) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ActorStateTransition_stagedOrder, stage.ActorStateTransitions_referenceOrder, actorstatetransition, "ActorStateTransition")
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ActorStateTransitionShape_stagedOrder, stage.ActorStateTransitionShapes_referenceOrder, actorstatetransitionshape, "ActorStateTransitionShape")
}

func (analysis *Analysis) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Analysis_stagedOrder, stage.Analysiss_referenceOrder, analysis, "Analysis")
}

func (controlpointshape *ControlPointShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ControlPointShape_stagedOrder, stage.ControlPointShapes_referenceOrder, controlpointshape, "ControlPointShape")
}

func (diagram *Diagram) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Diagram_stagedOrder, stage.Diagrams_referenceOrder, diagram, "Diagram")
}

func (document *Document) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Document_stagedOrder, stage.Documents_referenceOrder, document, "Document")
}

func (documentuse *DocumentUse) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DocumentUse_stagedOrder, stage.DocumentUses_referenceOrder, documentuse, "DocumentUse")
}

func (evolutiondirection *EvolutionDirection) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EvolutionDirection_stagedOrder, stage.EvolutionDirections_referenceOrder, evolutiondirection, "EvolutionDirection")
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EvolutionDirectionShape_stagedOrder, stage.EvolutionDirectionShapes_referenceOrder, evolutiondirectionshape, "EvolutionDirectionShape")
}

func (foo *Foo) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Foo_stagedOrder, stage.Foos_referenceOrder, foo, "Foo")
}

func (geoobject *GeoObject) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GeoObject_stagedOrder, stage.GeoObjects_referenceOrder, geoobject, "GeoObject")
}

func (geoobjectuse *GeoObjectUse) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GeoObjectUse_stagedOrder, stage.GeoObjectUses_referenceOrder, geoobjectuse, "GeoObjectUse")
}

func (group *Group) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Group_stagedOrder, stage.Groups_referenceOrder, group, "Group")
}

func (groupuse *GroupUse) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GroupUse_stagedOrder, stage.GroupUses_referenceOrder, groupuse, "GroupUse")
}

func (library *Library) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Library_stagedOrder, stage.Librarys_referenceOrder, library, "Library")
}

func (mapobject *MapObject) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MapObject_stagedOrder, stage.MapObjects_referenceOrder, mapobject, "MapObject")
}

func (mapobjectuse *MapObjectUse) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MapObjectUse_stagedOrder, stage.MapObjectUses_referenceOrder, mapobjectuse, "MapObjectUse")
}

func (parameter *Parameter) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Parameter_stagedOrder, stage.Parameters_referenceOrder, parameter, "Parameter")
}

func (parametercategory *ParameterCategory) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ParameterCategory_stagedOrder, stage.ParameterCategorys_referenceOrder, parametercategory, "ParameterCategory")
}

func (parametercategoryuse *ParameterCategoryUse) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ParameterCategoryUse_stagedOrder, stage.ParameterCategoryUses_referenceOrder, parametercategoryuse, "ParameterCategoryUse")
}

func (parametershape *ParameterShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ParameterShape_stagedOrder, stage.ParameterShapes_referenceOrder, parametershape, "ParameterShape")
}

func (parametersaggregate *ParametersAggregate) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ParametersAggregate_stagedOrder, stage.ParametersAggregates_referenceOrder, parametersaggregate, "ParametersAggregate")
}

func (parametersaggregateshape *ParametersAggregateShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ParametersAggregateShape_stagedOrder, stage.ParametersAggregateShapes_referenceOrder, parametersaggregateshape, "ParametersAggregateShape")
}

func (position *Position) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Position_stagedOrder, stage.Positions_referenceOrder, position, "Position")
}

func (repository *Repository) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Repository_stagedOrder, stage.Repositorys_referenceOrder, repository, "Repository")
}

func (scenario *Scenario) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Scenario_stagedOrder, stage.Scenarios_referenceOrder, scenario, "Scenario")
}

func (user *User) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.User_stagedOrder, stage.Users_referenceOrder, user, "User")
}

func (useruse *UserUse) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.UserUse_stagedOrder, stage.UserUses_referenceOrder, useruse, "UserUse")
}

func (workspace *Workspace) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Workspace_stagedOrder, stage.Workspaces_referenceOrder, workspace, "Workspace")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (actorstate *ActorState) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(actorstate, actorstate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (actorstate *ActorState) GongGetReferenceIdentifier(stage *Stage) string {
	return actorstate.GongGetIdentifier(stage)
}

func (actorstateshape *ActorStateShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(actorstateshape, actorstateshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (actorstateshape *ActorStateShape) GongGetReferenceIdentifier(stage *Stage) string {
	return actorstateshape.GongGetIdentifier(stage)
}

func (actorstatetransition *ActorStateTransition) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(actorstatetransition, actorstatetransition.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (actorstatetransition *ActorStateTransition) GongGetReferenceIdentifier(stage *Stage) string {
	return actorstatetransition.GongGetIdentifier(stage)
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(actorstatetransitionshape, actorstatetransitionshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (actorstatetransitionshape *ActorStateTransitionShape) GongGetReferenceIdentifier(stage *Stage) string {
	return actorstatetransitionshape.GongGetIdentifier(stage)
}

func (analysis *Analysis) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(analysis, analysis.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (analysis *Analysis) GongGetReferenceIdentifier(stage *Stage) string {
	return analysis.GongGetIdentifier(stage)
}

func (controlpointshape *ControlPointShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(controlpointshape, controlpointshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (controlpointshape *ControlPointShape) GongGetReferenceIdentifier(stage *Stage) string {
	return controlpointshape.GongGetIdentifier(stage)
}

func (diagram *Diagram) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(diagram, diagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagram *Diagram) GongGetReferenceIdentifier(stage *Stage) string {
	return diagram.GongGetIdentifier(stage)
}

func (document *Document) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(document, document.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (document *Document) GongGetReferenceIdentifier(stage *Stage) string {
	return document.GongGetIdentifier(stage)
}

func (documentuse *DocumentUse) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(documentuse, documentuse.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (documentuse *DocumentUse) GongGetReferenceIdentifier(stage *Stage) string {
	return documentuse.GongGetIdentifier(stage)
}

func (evolutiondirection *EvolutionDirection) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(evolutiondirection, evolutiondirection.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (evolutiondirection *EvolutionDirection) GongGetReferenceIdentifier(stage *Stage) string {
	return evolutiondirection.GongGetIdentifier(stage)
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(evolutiondirectionshape, evolutiondirectionshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (evolutiondirectionshape *EvolutionDirectionShape) GongGetReferenceIdentifier(stage *Stage) string {
	return evolutiondirectionshape.GongGetIdentifier(stage)
}

func (foo *Foo) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(foo, foo.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (foo *Foo) GongGetReferenceIdentifier(stage *Stage) string {
	return foo.GongGetIdentifier(stage)
}

func (geoobject *GeoObject) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(geoobject, geoobject.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (geoobject *GeoObject) GongGetReferenceIdentifier(stage *Stage) string {
	return geoobject.GongGetIdentifier(stage)
}

func (geoobjectuse *GeoObjectUse) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(geoobjectuse, geoobjectuse.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (geoobjectuse *GeoObjectUse) GongGetReferenceIdentifier(stage *Stage) string {
	return geoobjectuse.GongGetIdentifier(stage)
}

func (group *Group) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(group, group.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (group *Group) GongGetReferenceIdentifier(stage *Stage) string {
	return group.GongGetIdentifier(stage)
}

func (groupuse *GroupUse) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(groupuse, groupuse.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (groupuse *GroupUse) GongGetReferenceIdentifier(stage *Stage) string {
	return groupuse.GongGetIdentifier(stage)
}

func (library *Library) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(library, library.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (library *Library) GongGetReferenceIdentifier(stage *Stage) string {
	return library.GongGetIdentifier(stage)
}

func (mapobject *MapObject) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(mapobject, mapobject.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (mapobject *MapObject) GongGetReferenceIdentifier(stage *Stage) string {
	return mapobject.GongGetIdentifier(stage)
}

func (mapobjectuse *MapObjectUse) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(mapobjectuse, mapobjectuse.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (mapobjectuse *MapObjectUse) GongGetReferenceIdentifier(stage *Stage) string {
	return mapobjectuse.GongGetIdentifier(stage)
}

func (parameter *Parameter) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(parameter, parameter.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (parameter *Parameter) GongGetReferenceIdentifier(stage *Stage) string {
	return parameter.GongGetIdentifier(stage)
}

func (parametercategory *ParameterCategory) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(parametercategory, parametercategory.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (parametercategory *ParameterCategory) GongGetReferenceIdentifier(stage *Stage) string {
	return parametercategory.GongGetIdentifier(stage)
}

func (parametercategoryuse *ParameterCategoryUse) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(parametercategoryuse, parametercategoryuse.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (parametercategoryuse *ParameterCategoryUse) GongGetReferenceIdentifier(stage *Stage) string {
	return parametercategoryuse.GongGetIdentifier(stage)
}

func (parametershape *ParameterShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(parametershape, parametershape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (parametershape *ParameterShape) GongGetReferenceIdentifier(stage *Stage) string {
	return parametershape.GongGetIdentifier(stage)
}

func (parametersaggregate *ParametersAggregate) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(parametersaggregate, parametersaggregate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (parametersaggregate *ParametersAggregate) GongGetReferenceIdentifier(stage *Stage) string {
	return parametersaggregate.GongGetIdentifier(stage)
}

func (parametersaggregateshape *ParametersAggregateShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(parametersaggregateshape, parametersaggregateshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (parametersaggregateshape *ParametersAggregateShape) GongGetReferenceIdentifier(stage *Stage) string {
	return parametersaggregateshape.GongGetIdentifier(stage)
}

func (position *Position) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(position, position.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (position *Position) GongGetReferenceIdentifier(stage *Stage) string {
	return position.GongGetIdentifier(stage)
}

func (repository *Repository) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(repository, repository.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (repository *Repository) GongGetReferenceIdentifier(stage *Stage) string {
	return repository.GongGetIdentifier(stage)
}

func (scenario *Scenario) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(scenario, scenario.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (scenario *Scenario) GongGetReferenceIdentifier(stage *Stage) string {
	return scenario.GongGetIdentifier(stage)
}

func (user *User) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(user, user.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (user *User) GongGetReferenceIdentifier(stage *Stage) string {
	return user.GongGetIdentifier(stage)
}

func (useruse *UserUse) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(useruse, useruse.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (useruse *UserUse) GongGetReferenceIdentifier(stage *Stage) string {
	return useruse.GongGetIdentifier(stage)
}

func (workspace *Workspace) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(workspace, workspace.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (workspace *Workspace) GongGetReferenceIdentifier(stage *Stage) string {
	return workspace.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (actorstate *ActorState) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(actorstate.GongGetIdentifier(stage), "ActorState", actorstate.Name)
}

func (actorstateshape *ActorStateShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(actorstateshape.GongGetIdentifier(stage), "ActorStateShape", actorstateshape.Name)
}

func (actorstatetransition *ActorStateTransition) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(actorstatetransition.GongGetIdentifier(stage), "ActorStateTransition", actorstatetransition.Name)
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(actorstatetransitionshape.GongGetIdentifier(stage), "ActorStateTransitionShape", actorstatetransitionshape.Name)
}

func (analysis *Analysis) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(analysis.GongGetIdentifier(stage), "Analysis", analysis.Name)
}

func (controlpointshape *ControlPointShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(controlpointshape.GongGetIdentifier(stage), "ControlPointShape", controlpointshape.Name)
}

func (diagram *Diagram) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(diagram.GongGetIdentifier(stage), "Diagram", diagram.Name)
}

func (document *Document) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(document.GongGetIdentifier(stage), "Document", document.Name)
}

func (documentuse *DocumentUse) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(documentuse.GongGetIdentifier(stage), "DocumentUse", documentuse.Name)
}

func (evolutiondirection *EvolutionDirection) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(evolutiondirection.GongGetIdentifier(stage), "EvolutionDirection", evolutiondirection.Name)
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(evolutiondirectionshape.GongGetIdentifier(stage), "EvolutionDirectionShape", evolutiondirectionshape.Name)
}

func (foo *Foo) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(foo.GongGetIdentifier(stage), "Foo", foo.Name)
}

func (geoobject *GeoObject) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(geoobject.GongGetIdentifier(stage), "GeoObject", geoobject.Name)
}

func (geoobjectuse *GeoObjectUse) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(geoobjectuse.GongGetIdentifier(stage), "GeoObjectUse", geoobjectuse.Name)
}

func (group *Group) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(group.GongGetIdentifier(stage), "Group", group.Name)
}

func (groupuse *GroupUse) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(groupuse.GongGetIdentifier(stage), "GroupUse", groupuse.Name)
}

func (library *Library) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(library.GongGetIdentifier(stage), "Library", library.Name)
}

func (mapobject *MapObject) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(mapobject.GongGetIdentifier(stage), "MapObject", mapobject.Name)
}

func (mapobjectuse *MapObjectUse) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(mapobjectuse.GongGetIdentifier(stage), "MapObjectUse", mapobjectuse.Name)
}

func (parameter *Parameter) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(parameter.GongGetIdentifier(stage), "Parameter", parameter.Name)
}

func (parametercategory *ParameterCategory) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(parametercategory.GongGetIdentifier(stage), "ParameterCategory", parametercategory.Name)
}

func (parametercategoryuse *ParameterCategoryUse) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(parametercategoryuse.GongGetIdentifier(stage), "ParameterCategoryUse", parametercategoryuse.Name)
}

func (parametershape *ParameterShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(parametershape.GongGetIdentifier(stage), "ParameterShape", parametershape.Name)
}

func (parametersaggregate *ParametersAggregate) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(parametersaggregate.GongGetIdentifier(stage), "ParametersAggregate", parametersaggregate.Name)
}

func (parametersaggregateshape *ParametersAggregateShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(parametersaggregateshape.GongGetIdentifier(stage), "ParametersAggregateShape", parametersaggregateshape.Name)
}

func (position *Position) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(position.GongGetIdentifier(stage), "Position", position.Name)
}

func (repository *Repository) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(repository.GongGetIdentifier(stage), "Repository", repository.Name)
}

func (scenario *Scenario) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(scenario.GongGetIdentifier(stage), "Scenario", scenario.Name)
}

func (user *User) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(user.GongGetIdentifier(stage), "User", user.Name)
}

func (useruse *UserUse) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(useruse.GongGetIdentifier(stage), "UserUse", useruse.Name)
}

func (workspace *Workspace) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(workspace.GongGetIdentifier(stage), "Workspace", workspace.Name)
}

// insertion point for unstaging
func (actorstate *ActorState) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(actorstate.GongGetReferenceIdentifier(stage))
}

func (actorstateshape *ActorStateShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(actorstateshape.GongGetReferenceIdentifier(stage))
}

func (actorstatetransition *ActorStateTransition) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(actorstatetransition.GongGetReferenceIdentifier(stage))
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(actorstatetransitionshape.GongGetReferenceIdentifier(stage))
}

func (analysis *Analysis) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(analysis.GongGetReferenceIdentifier(stage))
}

func (controlpointshape *ControlPointShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(controlpointshape.GongGetReferenceIdentifier(stage))
}

func (diagram *Diagram) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(diagram.GongGetReferenceIdentifier(stage))
}

func (document *Document) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(document.GongGetReferenceIdentifier(stage))
}

func (documentuse *DocumentUse) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(documentuse.GongGetReferenceIdentifier(stage))
}

func (evolutiondirection *EvolutionDirection) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(evolutiondirection.GongGetReferenceIdentifier(stage))
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(evolutiondirectionshape.GongGetReferenceIdentifier(stage))
}

func (foo *Foo) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(foo.GongGetReferenceIdentifier(stage))
}

func (geoobject *GeoObject) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(geoobject.GongGetReferenceIdentifier(stage))
}

func (geoobjectuse *GeoObjectUse) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(geoobjectuse.GongGetReferenceIdentifier(stage))
}

func (group *Group) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(group.GongGetReferenceIdentifier(stage))
}

func (groupuse *GroupUse) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(groupuse.GongGetReferenceIdentifier(stage))
}

func (library *Library) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(library.GongGetReferenceIdentifier(stage))
}

func (mapobject *MapObject) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(mapobject.GongGetReferenceIdentifier(stage))
}

func (mapobjectuse *MapObjectUse) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(mapobjectuse.GongGetReferenceIdentifier(stage))
}

func (parameter *Parameter) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(parameter.GongGetReferenceIdentifier(stage))
}

func (parametercategory *ParameterCategory) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(parametercategory.GongGetReferenceIdentifier(stage))
}

func (parametercategoryuse *ParameterCategoryUse) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(parametercategoryuse.GongGetReferenceIdentifier(stage))
}

func (parametershape *ParameterShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(parametershape.GongGetReferenceIdentifier(stage))
}

func (parametersaggregate *ParametersAggregate) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(parametersaggregate.GongGetReferenceIdentifier(stage))
}

func (parametersaggregateshape *ParametersAggregateShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(parametersaggregateshape.GongGetReferenceIdentifier(stage))
}

func (position *Position) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(position.GongGetReferenceIdentifier(stage))
}

func (repository *Repository) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(repository.GongGetReferenceIdentifier(stage))
}

func (scenario *Scenario) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(scenario.GongGetReferenceIdentifier(stage))
}

func (user *User) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(user.GongGetReferenceIdentifier(stage))
}

func (useruse *UserUse) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(useruse.GongGetReferenceIdentifier(stage))
}

func (workspace *Workspace) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(workspace.GongGetReferenceIdentifier(stage))
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
