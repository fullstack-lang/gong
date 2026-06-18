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
	// Compute reverse map for named struct ActorState
	// insertion point per field

	// Compute reverse map for named struct ActorStateShape
	// insertion point per field

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

	// Compute reverse map for named struct ControlPointShape
	// insertion point per field

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

	// Compute reverse map for named struct DocumentUse
	// insertion point per field

	// Compute reverse map for named struct EvolutionDirection
	// insertion point per field

	// Compute reverse map for named struct EvolutionDirectionShape
	// insertion point per field

	// Compute reverse map for named struct Foo
	// insertion point per field

	// Compute reverse map for named struct GeoObject
	// insertion point per field

	// Compute reverse map for named struct GeoObjectUse
	// insertion point per field

	// Compute reverse map for named struct Group
	// insertion point per field
	stage.Group_UserUse_reverseMap = make(map[*UserUse]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _useruse := range group.UserUse {
			stage.Group_UserUse_reverseMap[_useruse] = group
		}
	}

	// Compute reverse map for named struct GroupUse
	// insertion point per field

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

	// Compute reverse map for named struct MapObject
	// insertion point per field

	// Compute reverse map for named struct MapObjectUse
	// insertion point per field

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

	// Compute reverse map for named struct ParameterCategoryUse
	// insertion point per field

	// Compute reverse map for named struct ParameterShape
	// insertion point per field

	// Compute reverse map for named struct ParametersAggregate
	// insertion point per field
	stage.ParametersAggregate_Parameters_reverseMap = make(map[*Parameter]*ParametersAggregate)
	for parametersaggregate := range stage.ParametersAggregates {
		_ = parametersaggregate
		for _, _parameter := range parametersaggregate.Parameters {
			stage.ParametersAggregate_Parameters_reverseMap[_parameter] = parametersaggregate
		}
	}

	// Compute reverse map for named struct ParametersAggregateShape
	// insertion point per field

	// Compute reverse map for named struct Position
	// insertion point per field

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

	// Compute reverse map for named struct User
	// insertion point per field

	// Compute reverse map for named struct UserUse
	// insertion point per field

	// Compute reverse map for named struct Workspace
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.ActorStates {
		res = append(res, instance)
	}

	for instance := range stage.ActorStateShapes {
		res = append(res, instance)
	}

	for instance := range stage.ActorStateTransitions {
		res = append(res, instance)
	}

	for instance := range stage.ActorStateTransitionShapes {
		res = append(res, instance)
	}

	for instance := range stage.Analysiss {
		res = append(res, instance)
	}

	for instance := range stage.ControlPointShapes {
		res = append(res, instance)
	}

	for instance := range stage.Diagrams {
		res = append(res, instance)
	}

	for instance := range stage.Documents {
		res = append(res, instance)
	}

	for instance := range stage.DocumentUses {
		res = append(res, instance)
	}

	for instance := range stage.EvolutionDirections {
		res = append(res, instance)
	}

	for instance := range stage.EvolutionDirectionShapes {
		res = append(res, instance)
	}

	for instance := range stage.Foos {
		res = append(res, instance)
	}

	for instance := range stage.GeoObjects {
		res = append(res, instance)
	}

	for instance := range stage.GeoObjectUses {
		res = append(res, instance)
	}

	for instance := range stage.Groups {
		res = append(res, instance)
	}

	for instance := range stage.GroupUses {
		res = append(res, instance)
	}

	for instance := range stage.Librarys {
		res = append(res, instance)
	}

	for instance := range stage.MapObjects {
		res = append(res, instance)
	}

	for instance := range stage.MapObjectUses {
		res = append(res, instance)
	}

	for instance := range stage.Parameters {
		res = append(res, instance)
	}

	for instance := range stage.ParameterCategorys {
		res = append(res, instance)
	}

	for instance := range stage.ParameterCategoryUses {
		res = append(res, instance)
	}

	for instance := range stage.ParameterShapes {
		res = append(res, instance)
	}

	for instance := range stage.ParametersAggregates {
		res = append(res, instance)
	}

	for instance := range stage.ParametersAggregateShapes {
		res = append(res, instance)
	}

	for instance := range stage.Positions {
		res = append(res, instance)
	}

	for instance := range stage.Repositorys {
		res = append(res, instance)
	}

	for instance := range stage.Scenarios {
		res = append(res, instance)
	}

	for instance := range stage.Users {
		res = append(res, instance)
	}

	for instance := range stage.UserUses {
		res = append(res, instance)
	}

	for instance := range stage.Workspaces {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (actorstate *ActorState) GongCopy() GongstructIF {
	newInstance := new(ActorState)
	actorstate.CopyBasicFields(newInstance)
	return newInstance
}

func (actorstateshape *ActorStateShape) GongCopy() GongstructIF {
	newInstance := new(ActorStateShape)
	actorstateshape.CopyBasicFields(newInstance)
	return newInstance
}

func (actorstatetransition *ActorStateTransition) GongCopy() GongstructIF {
	newInstance := new(ActorStateTransition)
	actorstatetransition.CopyBasicFields(newInstance)
	return newInstance
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongCopy() GongstructIF {
	newInstance := new(ActorStateTransitionShape)
	actorstatetransitionshape.CopyBasicFields(newInstance)
	return newInstance
}

func (analysis *Analysis) GongCopy() GongstructIF {
	newInstance := new(Analysis)
	analysis.CopyBasicFields(newInstance)
	return newInstance
}

func (controlpointshape *ControlPointShape) GongCopy() GongstructIF {
	newInstance := new(ControlPointShape)
	controlpointshape.CopyBasicFields(newInstance)
	return newInstance
}

func (diagram *Diagram) GongCopy() GongstructIF {
	newInstance := new(Diagram)
	diagram.CopyBasicFields(newInstance)
	return newInstance
}

func (document *Document) GongCopy() GongstructIF {
	newInstance := new(Document)
	document.CopyBasicFields(newInstance)
	return newInstance
}

func (documentuse *DocumentUse) GongCopy() GongstructIF {
	newInstance := new(DocumentUse)
	documentuse.CopyBasicFields(newInstance)
	return newInstance
}

func (evolutiondirection *EvolutionDirection) GongCopy() GongstructIF {
	newInstance := new(EvolutionDirection)
	evolutiondirection.CopyBasicFields(newInstance)
	return newInstance
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongCopy() GongstructIF {
	newInstance := new(EvolutionDirectionShape)
	evolutiondirectionshape.CopyBasicFields(newInstance)
	return newInstance
}

func (foo *Foo) GongCopy() GongstructIF {
	newInstance := new(Foo)
	foo.CopyBasicFields(newInstance)
	return newInstance
}

func (geoobject *GeoObject) GongCopy() GongstructIF {
	newInstance := new(GeoObject)
	geoobject.CopyBasicFields(newInstance)
	return newInstance
}

func (geoobjectuse *GeoObjectUse) GongCopy() GongstructIF {
	newInstance := new(GeoObjectUse)
	geoobjectuse.CopyBasicFields(newInstance)
	return newInstance
}

func (group *Group) GongCopy() GongstructIF {
	newInstance := new(Group)
	group.CopyBasicFields(newInstance)
	return newInstance
}

func (groupuse *GroupUse) GongCopy() GongstructIF {
	newInstance := new(GroupUse)
	groupuse.CopyBasicFields(newInstance)
	return newInstance
}

func (library *Library) GongCopy() GongstructIF {
	newInstance := new(Library)
	library.CopyBasicFields(newInstance)
	return newInstance
}

func (mapobject *MapObject) GongCopy() GongstructIF {
	newInstance := new(MapObject)
	mapobject.CopyBasicFields(newInstance)
	return newInstance
}

func (mapobjectuse *MapObjectUse) GongCopy() GongstructIF {
	newInstance := new(MapObjectUse)
	mapobjectuse.CopyBasicFields(newInstance)
	return newInstance
}

func (parameter *Parameter) GongCopy() GongstructIF {
	newInstance := new(Parameter)
	parameter.CopyBasicFields(newInstance)
	return newInstance
}

func (parametercategory *ParameterCategory) GongCopy() GongstructIF {
	newInstance := new(ParameterCategory)
	parametercategory.CopyBasicFields(newInstance)
	return newInstance
}

func (parametercategoryuse *ParameterCategoryUse) GongCopy() GongstructIF {
	newInstance := new(ParameterCategoryUse)
	parametercategoryuse.CopyBasicFields(newInstance)
	return newInstance
}

func (parametershape *ParameterShape) GongCopy() GongstructIF {
	newInstance := new(ParameterShape)
	parametershape.CopyBasicFields(newInstance)
	return newInstance
}

func (parametersaggregate *ParametersAggregate) GongCopy() GongstructIF {
	newInstance := new(ParametersAggregate)
	parametersaggregate.CopyBasicFields(newInstance)
	return newInstance
}

func (parametersaggregateshape *ParametersAggregateShape) GongCopy() GongstructIF {
	newInstance := new(ParametersAggregateShape)
	parametersaggregateshape.CopyBasicFields(newInstance)
	return newInstance
}

func (position *Position) GongCopy() GongstructIF {
	newInstance := new(Position)
	position.CopyBasicFields(newInstance)
	return newInstance
}

func (repository *Repository) GongCopy() GongstructIF {
	newInstance := new(Repository)
	repository.CopyBasicFields(newInstance)
	return newInstance
}

func (scenario *Scenario) GongCopy() GongstructIF {
	newInstance := new(Scenario)
	scenario.CopyBasicFields(newInstance)
	return newInstance
}

func (user *User) GongCopy() GongstructIF {
	newInstance := new(User)
	user.CopyBasicFields(newInstance)
	return newInstance
}

func (useruse *UserUse) GongCopy() GongstructIF {
	newInstance := new(UserUse)
	useruse.CopyBasicFields(newInstance)
	return newInstance
}

func (workspace *Workspace) GongCopy() GongstructIF {
	newInstance := new(Workspace)
	workspace.CopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (actorstate *ActorState) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(actorstate).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(actorstate), uint64(GetOrderPointerGongstruct(stage, actorstate)))
	return
}

func (actorstateshape *ActorStateShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(actorstateshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(actorstateshape), uint64(GetOrderPointerGongstruct(stage, actorstateshape)))
	return
}

func (actorstatetransition *ActorStateTransition) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(actorstatetransition).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(actorstatetransition), uint64(GetOrderPointerGongstruct(stage, actorstatetransition)))
	return
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(actorstatetransitionshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(actorstatetransitionshape), uint64(GetOrderPointerGongstruct(stage, actorstatetransitionshape)))
	return
}

func (analysis *Analysis) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(analysis).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(analysis), uint64(GetOrderPointerGongstruct(stage, analysis)))
	return
}

func (controlpointshape *ControlPointShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(controlpointshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(controlpointshape), uint64(GetOrderPointerGongstruct(stage, controlpointshape)))
	return
}

func (diagram *Diagram) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(diagram).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(diagram), uint64(GetOrderPointerGongstruct(stage, diagram)))
	return
}

func (document *Document) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(document).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(document), uint64(GetOrderPointerGongstruct(stage, document)))
	return
}

func (documentuse *DocumentUse) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(documentuse).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(documentuse), uint64(GetOrderPointerGongstruct(stage, documentuse)))
	return
}

func (evolutiondirection *EvolutionDirection) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(evolutiondirection).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(evolutiondirection), uint64(GetOrderPointerGongstruct(stage, evolutiondirection)))
	return
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(evolutiondirectionshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(evolutiondirectionshape), uint64(GetOrderPointerGongstruct(stage, evolutiondirectionshape)))
	return
}

func (foo *Foo) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(foo).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(foo), uint64(GetOrderPointerGongstruct(stage, foo)))
	return
}

func (geoobject *GeoObject) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(geoobject).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(geoobject), uint64(GetOrderPointerGongstruct(stage, geoobject)))
	return
}

func (geoobjectuse *GeoObjectUse) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(geoobjectuse).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(geoobjectuse), uint64(GetOrderPointerGongstruct(stage, geoobjectuse)))
	return
}

func (group *Group) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(group).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(group), uint64(GetOrderPointerGongstruct(stage, group)))
	return
}

func (groupuse *GroupUse) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(groupuse).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(groupuse), uint64(GetOrderPointerGongstruct(stage, groupuse)))
	return
}

func (library *Library) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(library).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(library), uint64(GetOrderPointerGongstruct(stage, library)))
	return
}

func (mapobject *MapObject) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(mapobject).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(mapobject), uint64(GetOrderPointerGongstruct(stage, mapobject)))
	return
}

func (mapobjectuse *MapObjectUse) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(mapobjectuse).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(mapobjectuse), uint64(GetOrderPointerGongstruct(stage, mapobjectuse)))
	return
}

func (parameter *Parameter) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(parameter).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(parameter), uint64(GetOrderPointerGongstruct(stage, parameter)))
	return
}

func (parametercategory *ParameterCategory) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(parametercategory).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(parametercategory), uint64(GetOrderPointerGongstruct(stage, parametercategory)))
	return
}

func (parametercategoryuse *ParameterCategoryUse) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(parametercategoryuse).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(parametercategoryuse), uint64(GetOrderPointerGongstruct(stage, parametercategoryuse)))
	return
}

func (parametershape *ParameterShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(parametershape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(parametershape), uint64(GetOrderPointerGongstruct(stage, parametershape)))
	return
}

func (parametersaggregate *ParametersAggregate) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(parametersaggregate).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(parametersaggregate), uint64(GetOrderPointerGongstruct(stage, parametersaggregate)))
	return
}

func (parametersaggregateshape *ParametersAggregateShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(parametersaggregateshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(parametersaggregateshape), uint64(GetOrderPointerGongstruct(stage, parametersaggregateshape)))
	return
}

func (position *Position) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(position).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(position), uint64(GetOrderPointerGongstruct(stage, position)))
	return
}

func (repository *Repository) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(repository).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(repository), uint64(GetOrderPointerGongstruct(stage, repository)))
	return
}

func (scenario *Scenario) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(scenario).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(scenario), uint64(GetOrderPointerGongstruct(stage, scenario)))
	return
}

func (user *User) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(user).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(user), uint64(GetOrderPointerGongstruct(stage, user)))
	return
}

func (useruse *UserUse) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(useruse).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(useruse), uint64(GetOrderPointerGongstruct(stage, useruse)))
	return
}

func (workspace *Workspace) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(workspace).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(workspace), uint64(GetOrderPointerGongstruct(stage, workspace)))
	return
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
	var actorstates_newInstances []*ActorState
	var actorstates_deletedInstances []*ActorState

	// parse all staged instances and check if they have a reference
	for actorstate := range stage.ActorStates {
		if ref, ok := stage.ActorStates_reference[actorstate]; !ok {
			actorstates_newInstances = append(actorstates_newInstances, actorstate)
			newInstancesSlice = append(newInstancesSlice, actorstate.GongMarshallIdentifier(stage))
			if stage.ActorStates_referenceOrder == nil {
				stage.ActorStates_referenceOrder = make(map[*ActorState]uint)
			}
			stage.ActorStates_referenceOrder[actorstate] = stage.ActorState_stagedOrder[actorstate]
			newInstancesReverseSlice = append(newInstancesReverseSlice, actorstate.GongMarshallUnstaging(stage))
			// delete(stage.ActorStates_referenceOrder, actorstate)
			fieldInitializers, pointersInitializations := actorstate.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ActorState_stagedOrder[ref] = stage.ActorState_stagedOrder[actorstate]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := actorstate.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, actorstate)
			// delete(stage.ActorState_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if actorstate.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", actorstate.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ActorStates_reference {
		instance := stage.ActorStates_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ActorStates[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			actorstates_deletedInstances = append(actorstates_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(actorstates_newInstances)
	lenDeletedInstances += len(actorstates_deletedInstances)
	var actorstateshapes_newInstances []*ActorStateShape
	var actorstateshapes_deletedInstances []*ActorStateShape

	// parse all staged instances and check if they have a reference
	for actorstateshape := range stage.ActorStateShapes {
		if ref, ok := stage.ActorStateShapes_reference[actorstateshape]; !ok {
			actorstateshapes_newInstances = append(actorstateshapes_newInstances, actorstateshape)
			newInstancesSlice = append(newInstancesSlice, actorstateshape.GongMarshallIdentifier(stage))
			if stage.ActorStateShapes_referenceOrder == nil {
				stage.ActorStateShapes_referenceOrder = make(map[*ActorStateShape]uint)
			}
			stage.ActorStateShapes_referenceOrder[actorstateshape] = stage.ActorStateShape_stagedOrder[actorstateshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, actorstateshape.GongMarshallUnstaging(stage))
			// delete(stage.ActorStateShapes_referenceOrder, actorstateshape)
			fieldInitializers, pointersInitializations := actorstateshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ActorStateShape_stagedOrder[ref] = stage.ActorStateShape_stagedOrder[actorstateshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := actorstateshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, actorstateshape)
			// delete(stage.ActorStateShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if actorstateshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", actorstateshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ActorStateShapes_reference {
		instance := stage.ActorStateShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ActorStateShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			actorstateshapes_deletedInstances = append(actorstateshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(actorstateshapes_newInstances)
	lenDeletedInstances += len(actorstateshapes_deletedInstances)
	var actorstatetransitions_newInstances []*ActorStateTransition
	var actorstatetransitions_deletedInstances []*ActorStateTransition

	// parse all staged instances and check if they have a reference
	for actorstatetransition := range stage.ActorStateTransitions {
		if ref, ok := stage.ActorStateTransitions_reference[actorstatetransition]; !ok {
			actorstatetransitions_newInstances = append(actorstatetransitions_newInstances, actorstatetransition)
			newInstancesSlice = append(newInstancesSlice, actorstatetransition.GongMarshallIdentifier(stage))
			if stage.ActorStateTransitions_referenceOrder == nil {
				stage.ActorStateTransitions_referenceOrder = make(map[*ActorStateTransition]uint)
			}
			stage.ActorStateTransitions_referenceOrder[actorstatetransition] = stage.ActorStateTransition_stagedOrder[actorstatetransition]
			newInstancesReverseSlice = append(newInstancesReverseSlice, actorstatetransition.GongMarshallUnstaging(stage))
			// delete(stage.ActorStateTransitions_referenceOrder, actorstatetransition)
			fieldInitializers, pointersInitializations := actorstatetransition.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ActorStateTransition_stagedOrder[ref] = stage.ActorStateTransition_stagedOrder[actorstatetransition]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := actorstatetransition.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, actorstatetransition)
			// delete(stage.ActorStateTransition_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if actorstatetransition.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", actorstatetransition.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ActorStateTransitions_reference {
		instance := stage.ActorStateTransitions_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ActorStateTransitions[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			actorstatetransitions_deletedInstances = append(actorstatetransitions_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(actorstatetransitions_newInstances)
	lenDeletedInstances += len(actorstatetransitions_deletedInstances)
	var actorstatetransitionshapes_newInstances []*ActorStateTransitionShape
	var actorstatetransitionshapes_deletedInstances []*ActorStateTransitionShape

	// parse all staged instances and check if they have a reference
	for actorstatetransitionshape := range stage.ActorStateTransitionShapes {
		if ref, ok := stage.ActorStateTransitionShapes_reference[actorstatetransitionshape]; !ok {
			actorstatetransitionshapes_newInstances = append(actorstatetransitionshapes_newInstances, actorstatetransitionshape)
			newInstancesSlice = append(newInstancesSlice, actorstatetransitionshape.GongMarshallIdentifier(stage))
			if stage.ActorStateTransitionShapes_referenceOrder == nil {
				stage.ActorStateTransitionShapes_referenceOrder = make(map[*ActorStateTransitionShape]uint)
			}
			stage.ActorStateTransitionShapes_referenceOrder[actorstatetransitionshape] = stage.ActorStateTransitionShape_stagedOrder[actorstatetransitionshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, actorstatetransitionshape.GongMarshallUnstaging(stage))
			// delete(stage.ActorStateTransitionShapes_referenceOrder, actorstatetransitionshape)
			fieldInitializers, pointersInitializations := actorstatetransitionshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ActorStateTransitionShape_stagedOrder[ref] = stage.ActorStateTransitionShape_stagedOrder[actorstatetransitionshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := actorstatetransitionshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, actorstatetransitionshape)
			// delete(stage.ActorStateTransitionShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if actorstatetransitionshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", actorstatetransitionshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ActorStateTransitionShapes_reference {
		instance := stage.ActorStateTransitionShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ActorStateTransitionShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			actorstatetransitionshapes_deletedInstances = append(actorstatetransitionshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(actorstatetransitionshapes_newInstances)
	lenDeletedInstances += len(actorstatetransitionshapes_deletedInstances)
	var analysiss_newInstances []*Analysis
	var analysiss_deletedInstances []*Analysis

	// parse all staged instances and check if they have a reference
	for analysis := range stage.Analysiss {
		if ref, ok := stage.Analysiss_reference[analysis]; !ok {
			analysiss_newInstances = append(analysiss_newInstances, analysis)
			newInstancesSlice = append(newInstancesSlice, analysis.GongMarshallIdentifier(stage))
			if stage.Analysiss_referenceOrder == nil {
				stage.Analysiss_referenceOrder = make(map[*Analysis]uint)
			}
			stage.Analysiss_referenceOrder[analysis] = stage.Analysis_stagedOrder[analysis]
			newInstancesReverseSlice = append(newInstancesReverseSlice, analysis.GongMarshallUnstaging(stage))
			// delete(stage.Analysiss_referenceOrder, analysis)
			fieldInitializers, pointersInitializations := analysis.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Analysis_stagedOrder[ref] = stage.Analysis_stagedOrder[analysis]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := analysis.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, analysis)
			// delete(stage.Analysis_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if analysis.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", analysis.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Analysiss_reference {
		instance := stage.Analysiss_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Analysiss[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			analysiss_deletedInstances = append(analysiss_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(analysiss_newInstances)
	lenDeletedInstances += len(analysiss_deletedInstances)
	var controlpointshapes_newInstances []*ControlPointShape
	var controlpointshapes_deletedInstances []*ControlPointShape

	// parse all staged instances and check if they have a reference
	for controlpointshape := range stage.ControlPointShapes {
		if ref, ok := stage.ControlPointShapes_reference[controlpointshape]; !ok {
			controlpointshapes_newInstances = append(controlpointshapes_newInstances, controlpointshape)
			newInstancesSlice = append(newInstancesSlice, controlpointshape.GongMarshallIdentifier(stage))
			if stage.ControlPointShapes_referenceOrder == nil {
				stage.ControlPointShapes_referenceOrder = make(map[*ControlPointShape]uint)
			}
			stage.ControlPointShapes_referenceOrder[controlpointshape] = stage.ControlPointShape_stagedOrder[controlpointshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, controlpointshape.GongMarshallUnstaging(stage))
			// delete(stage.ControlPointShapes_referenceOrder, controlpointshape)
			fieldInitializers, pointersInitializations := controlpointshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ControlPointShape_stagedOrder[ref] = stage.ControlPointShape_stagedOrder[controlpointshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := controlpointshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, controlpointshape)
			// delete(stage.ControlPointShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if controlpointshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", controlpointshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ControlPointShapes_reference {
		instance := stage.ControlPointShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ControlPointShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			controlpointshapes_deletedInstances = append(controlpointshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(controlpointshapes_newInstances)
	lenDeletedInstances += len(controlpointshapes_deletedInstances)
	var diagrams_newInstances []*Diagram
	var diagrams_deletedInstances []*Diagram

	// parse all staged instances and check if they have a reference
	for diagram := range stage.Diagrams {
		if ref, ok := stage.Diagrams_reference[diagram]; !ok {
			diagrams_newInstances = append(diagrams_newInstances, diagram)
			newInstancesSlice = append(newInstancesSlice, diagram.GongMarshallIdentifier(stage))
			if stage.Diagrams_referenceOrder == nil {
				stage.Diagrams_referenceOrder = make(map[*Diagram]uint)
			}
			stage.Diagrams_referenceOrder[diagram] = stage.Diagram_stagedOrder[diagram]
			newInstancesReverseSlice = append(newInstancesReverseSlice, diagram.GongMarshallUnstaging(stage))
			// delete(stage.Diagrams_referenceOrder, diagram)
			fieldInitializers, pointersInitializations := diagram.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Diagram_stagedOrder[ref] = stage.Diagram_stagedOrder[diagram]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := diagram.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, diagram)
			// delete(stage.Diagram_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if diagram.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", diagram.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Diagrams_reference {
		instance := stage.Diagrams_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Diagrams[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			diagrams_deletedInstances = append(diagrams_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(diagrams_newInstances)
	lenDeletedInstances += len(diagrams_deletedInstances)
	var documents_newInstances []*Document
	var documents_deletedInstances []*Document

	// parse all staged instances and check if they have a reference
	for document := range stage.Documents {
		if ref, ok := stage.Documents_reference[document]; !ok {
			documents_newInstances = append(documents_newInstances, document)
			newInstancesSlice = append(newInstancesSlice, document.GongMarshallIdentifier(stage))
			if stage.Documents_referenceOrder == nil {
				stage.Documents_referenceOrder = make(map[*Document]uint)
			}
			stage.Documents_referenceOrder[document] = stage.Document_stagedOrder[document]
			newInstancesReverseSlice = append(newInstancesReverseSlice, document.GongMarshallUnstaging(stage))
			// delete(stage.Documents_referenceOrder, document)
			fieldInitializers, pointersInitializations := document.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Document_stagedOrder[ref] = stage.Document_stagedOrder[document]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := document.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, document)
			// delete(stage.Document_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if document.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", document.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Documents_reference {
		instance := stage.Documents_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Documents[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			documents_deletedInstances = append(documents_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(documents_newInstances)
	lenDeletedInstances += len(documents_deletedInstances)
	var documentuses_newInstances []*DocumentUse
	var documentuses_deletedInstances []*DocumentUse

	// parse all staged instances and check if they have a reference
	for documentuse := range stage.DocumentUses {
		if ref, ok := stage.DocumentUses_reference[documentuse]; !ok {
			documentuses_newInstances = append(documentuses_newInstances, documentuse)
			newInstancesSlice = append(newInstancesSlice, documentuse.GongMarshallIdentifier(stage))
			if stage.DocumentUses_referenceOrder == nil {
				stage.DocumentUses_referenceOrder = make(map[*DocumentUse]uint)
			}
			stage.DocumentUses_referenceOrder[documentuse] = stage.DocumentUse_stagedOrder[documentuse]
			newInstancesReverseSlice = append(newInstancesReverseSlice, documentuse.GongMarshallUnstaging(stage))
			// delete(stage.DocumentUses_referenceOrder, documentuse)
			fieldInitializers, pointersInitializations := documentuse.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DocumentUse_stagedOrder[ref] = stage.DocumentUse_stagedOrder[documentuse]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := documentuse.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, documentuse)
			// delete(stage.DocumentUse_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if documentuse.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", documentuse.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.DocumentUses_reference {
		instance := stage.DocumentUses_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.DocumentUses[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			documentuses_deletedInstances = append(documentuses_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(documentuses_newInstances)
	lenDeletedInstances += len(documentuses_deletedInstances)
	var evolutiondirections_newInstances []*EvolutionDirection
	var evolutiondirections_deletedInstances []*EvolutionDirection

	// parse all staged instances and check if they have a reference
	for evolutiondirection := range stage.EvolutionDirections {
		if ref, ok := stage.EvolutionDirections_reference[evolutiondirection]; !ok {
			evolutiondirections_newInstances = append(evolutiondirections_newInstances, evolutiondirection)
			newInstancesSlice = append(newInstancesSlice, evolutiondirection.GongMarshallIdentifier(stage))
			if stage.EvolutionDirections_referenceOrder == nil {
				stage.EvolutionDirections_referenceOrder = make(map[*EvolutionDirection]uint)
			}
			stage.EvolutionDirections_referenceOrder[evolutiondirection] = stage.EvolutionDirection_stagedOrder[evolutiondirection]
			newInstancesReverseSlice = append(newInstancesReverseSlice, evolutiondirection.GongMarshallUnstaging(stage))
			// delete(stage.EvolutionDirections_referenceOrder, evolutiondirection)
			fieldInitializers, pointersInitializations := evolutiondirection.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.EvolutionDirection_stagedOrder[ref] = stage.EvolutionDirection_stagedOrder[evolutiondirection]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := evolutiondirection.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, evolutiondirection)
			// delete(stage.EvolutionDirection_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if evolutiondirection.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", evolutiondirection.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.EvolutionDirections_reference {
		instance := stage.EvolutionDirections_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.EvolutionDirections[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			evolutiondirections_deletedInstances = append(evolutiondirections_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(evolutiondirections_newInstances)
	lenDeletedInstances += len(evolutiondirections_deletedInstances)
	var evolutiondirectionshapes_newInstances []*EvolutionDirectionShape
	var evolutiondirectionshapes_deletedInstances []*EvolutionDirectionShape

	// parse all staged instances and check if they have a reference
	for evolutiondirectionshape := range stage.EvolutionDirectionShapes {
		if ref, ok := stage.EvolutionDirectionShapes_reference[evolutiondirectionshape]; !ok {
			evolutiondirectionshapes_newInstances = append(evolutiondirectionshapes_newInstances, evolutiondirectionshape)
			newInstancesSlice = append(newInstancesSlice, evolutiondirectionshape.GongMarshallIdentifier(stage))
			if stage.EvolutionDirectionShapes_referenceOrder == nil {
				stage.EvolutionDirectionShapes_referenceOrder = make(map[*EvolutionDirectionShape]uint)
			}
			stage.EvolutionDirectionShapes_referenceOrder[evolutiondirectionshape] = stage.EvolutionDirectionShape_stagedOrder[evolutiondirectionshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, evolutiondirectionshape.GongMarshallUnstaging(stage))
			// delete(stage.EvolutionDirectionShapes_referenceOrder, evolutiondirectionshape)
			fieldInitializers, pointersInitializations := evolutiondirectionshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.EvolutionDirectionShape_stagedOrder[ref] = stage.EvolutionDirectionShape_stagedOrder[evolutiondirectionshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := evolutiondirectionshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, evolutiondirectionshape)
			// delete(stage.EvolutionDirectionShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if evolutiondirectionshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", evolutiondirectionshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.EvolutionDirectionShapes_reference {
		instance := stage.EvolutionDirectionShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.EvolutionDirectionShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			evolutiondirectionshapes_deletedInstances = append(evolutiondirectionshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(evolutiondirectionshapes_newInstances)
	lenDeletedInstances += len(evolutiondirectionshapes_deletedInstances)
	var foos_newInstances []*Foo
	var foos_deletedInstances []*Foo

	// parse all staged instances and check if they have a reference
	for foo := range stage.Foos {
		if ref, ok := stage.Foos_reference[foo]; !ok {
			foos_newInstances = append(foos_newInstances, foo)
			newInstancesSlice = append(newInstancesSlice, foo.GongMarshallIdentifier(stage))
			if stage.Foos_referenceOrder == nil {
				stage.Foos_referenceOrder = make(map[*Foo]uint)
			}
			stage.Foos_referenceOrder[foo] = stage.Foo_stagedOrder[foo]
			newInstancesReverseSlice = append(newInstancesReverseSlice, foo.GongMarshallUnstaging(stage))
			// delete(stage.Foos_referenceOrder, foo)
			fieldInitializers, pointersInitializations := foo.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Foo_stagedOrder[ref] = stage.Foo_stagedOrder[foo]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := foo.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, foo)
			// delete(stage.Foo_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if foo.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", foo.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Foos_reference {
		instance := stage.Foos_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Foos[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			foos_deletedInstances = append(foos_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(foos_newInstances)
	lenDeletedInstances += len(foos_deletedInstances)
	var geoobjects_newInstances []*GeoObject
	var geoobjects_deletedInstances []*GeoObject

	// parse all staged instances and check if they have a reference
	for geoobject := range stage.GeoObjects {
		if ref, ok := stage.GeoObjects_reference[geoobject]; !ok {
			geoobjects_newInstances = append(geoobjects_newInstances, geoobject)
			newInstancesSlice = append(newInstancesSlice, geoobject.GongMarshallIdentifier(stage))
			if stage.GeoObjects_referenceOrder == nil {
				stage.GeoObjects_referenceOrder = make(map[*GeoObject]uint)
			}
			stage.GeoObjects_referenceOrder[geoobject] = stage.GeoObject_stagedOrder[geoobject]
			newInstancesReverseSlice = append(newInstancesReverseSlice, geoobject.GongMarshallUnstaging(stage))
			// delete(stage.GeoObjects_referenceOrder, geoobject)
			fieldInitializers, pointersInitializations := geoobject.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GeoObject_stagedOrder[ref] = stage.GeoObject_stagedOrder[geoobject]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := geoobject.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, geoobject)
			// delete(stage.GeoObject_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if geoobject.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", geoobject.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.GeoObjects_reference {
		instance := stage.GeoObjects_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.GeoObjects[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			geoobjects_deletedInstances = append(geoobjects_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(geoobjects_newInstances)
	lenDeletedInstances += len(geoobjects_deletedInstances)
	var geoobjectuses_newInstances []*GeoObjectUse
	var geoobjectuses_deletedInstances []*GeoObjectUse

	// parse all staged instances and check if they have a reference
	for geoobjectuse := range stage.GeoObjectUses {
		if ref, ok := stage.GeoObjectUses_reference[geoobjectuse]; !ok {
			geoobjectuses_newInstances = append(geoobjectuses_newInstances, geoobjectuse)
			newInstancesSlice = append(newInstancesSlice, geoobjectuse.GongMarshallIdentifier(stage))
			if stage.GeoObjectUses_referenceOrder == nil {
				stage.GeoObjectUses_referenceOrder = make(map[*GeoObjectUse]uint)
			}
			stage.GeoObjectUses_referenceOrder[geoobjectuse] = stage.GeoObjectUse_stagedOrder[geoobjectuse]
			newInstancesReverseSlice = append(newInstancesReverseSlice, geoobjectuse.GongMarshallUnstaging(stage))
			// delete(stage.GeoObjectUses_referenceOrder, geoobjectuse)
			fieldInitializers, pointersInitializations := geoobjectuse.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GeoObjectUse_stagedOrder[ref] = stage.GeoObjectUse_stagedOrder[geoobjectuse]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := geoobjectuse.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, geoobjectuse)
			// delete(stage.GeoObjectUse_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if geoobjectuse.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", geoobjectuse.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.GeoObjectUses_reference {
		instance := stage.GeoObjectUses_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.GeoObjectUses[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			geoobjectuses_deletedInstances = append(geoobjectuses_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(geoobjectuses_newInstances)
	lenDeletedInstances += len(geoobjectuses_deletedInstances)
	var groups_newInstances []*Group
	var groups_deletedInstances []*Group

	// parse all staged instances and check if they have a reference
	for group := range stage.Groups {
		if ref, ok := stage.Groups_reference[group]; !ok {
			groups_newInstances = append(groups_newInstances, group)
			newInstancesSlice = append(newInstancesSlice, group.GongMarshallIdentifier(stage))
			if stage.Groups_referenceOrder == nil {
				stage.Groups_referenceOrder = make(map[*Group]uint)
			}
			stage.Groups_referenceOrder[group] = stage.Group_stagedOrder[group]
			newInstancesReverseSlice = append(newInstancesReverseSlice, group.GongMarshallUnstaging(stage))
			// delete(stage.Groups_referenceOrder, group)
			fieldInitializers, pointersInitializations := group.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Group_stagedOrder[ref] = stage.Group_stagedOrder[group]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := group.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, group)
			// delete(stage.Group_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if group.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", group.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Groups_reference {
		instance := stage.Groups_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Groups[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			groups_deletedInstances = append(groups_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(groups_newInstances)
	lenDeletedInstances += len(groups_deletedInstances)
	var groupuses_newInstances []*GroupUse
	var groupuses_deletedInstances []*GroupUse

	// parse all staged instances and check if they have a reference
	for groupuse := range stage.GroupUses {
		if ref, ok := stage.GroupUses_reference[groupuse]; !ok {
			groupuses_newInstances = append(groupuses_newInstances, groupuse)
			newInstancesSlice = append(newInstancesSlice, groupuse.GongMarshallIdentifier(stage))
			if stage.GroupUses_referenceOrder == nil {
				stage.GroupUses_referenceOrder = make(map[*GroupUse]uint)
			}
			stage.GroupUses_referenceOrder[groupuse] = stage.GroupUse_stagedOrder[groupuse]
			newInstancesReverseSlice = append(newInstancesReverseSlice, groupuse.GongMarshallUnstaging(stage))
			// delete(stage.GroupUses_referenceOrder, groupuse)
			fieldInitializers, pointersInitializations := groupuse.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GroupUse_stagedOrder[ref] = stage.GroupUse_stagedOrder[groupuse]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := groupuse.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, groupuse)
			// delete(stage.GroupUse_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if groupuse.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", groupuse.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.GroupUses_reference {
		instance := stage.GroupUses_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.GroupUses[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			groupuses_deletedInstances = append(groupuses_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(groupuses_newInstances)
	lenDeletedInstances += len(groupuses_deletedInstances)
	var librarys_newInstances []*Library
	var librarys_deletedInstances []*Library

	// parse all staged instances and check if they have a reference
	for library := range stage.Librarys {
		if ref, ok := stage.Librarys_reference[library]; !ok {
			librarys_newInstances = append(librarys_newInstances, library)
			newInstancesSlice = append(newInstancesSlice, library.GongMarshallIdentifier(stage))
			if stage.Librarys_referenceOrder == nil {
				stage.Librarys_referenceOrder = make(map[*Library]uint)
			}
			stage.Librarys_referenceOrder[library] = stage.Library_stagedOrder[library]
			newInstancesReverseSlice = append(newInstancesReverseSlice, library.GongMarshallUnstaging(stage))
			// delete(stage.Librarys_referenceOrder, library)
			fieldInitializers, pointersInitializations := library.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Library_stagedOrder[ref] = stage.Library_stagedOrder[library]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := library.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, library)
			// delete(stage.Library_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if library.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", library.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Librarys_reference {
		instance := stage.Librarys_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Librarys[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			librarys_deletedInstances = append(librarys_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(librarys_newInstances)
	lenDeletedInstances += len(librarys_deletedInstances)
	var mapobjects_newInstances []*MapObject
	var mapobjects_deletedInstances []*MapObject

	// parse all staged instances and check if they have a reference
	for mapobject := range stage.MapObjects {
		if ref, ok := stage.MapObjects_reference[mapobject]; !ok {
			mapobjects_newInstances = append(mapobjects_newInstances, mapobject)
			newInstancesSlice = append(newInstancesSlice, mapobject.GongMarshallIdentifier(stage))
			if stage.MapObjects_referenceOrder == nil {
				stage.MapObjects_referenceOrder = make(map[*MapObject]uint)
			}
			stage.MapObjects_referenceOrder[mapobject] = stage.MapObject_stagedOrder[mapobject]
			newInstancesReverseSlice = append(newInstancesReverseSlice, mapobject.GongMarshallUnstaging(stage))
			// delete(stage.MapObjects_referenceOrder, mapobject)
			fieldInitializers, pointersInitializations := mapobject.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MapObject_stagedOrder[ref] = stage.MapObject_stagedOrder[mapobject]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := mapobject.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, mapobject)
			// delete(stage.MapObject_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if mapobject.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", mapobject.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.MapObjects_reference {
		instance := stage.MapObjects_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.MapObjects[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			mapobjects_deletedInstances = append(mapobjects_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(mapobjects_newInstances)
	lenDeletedInstances += len(mapobjects_deletedInstances)
	var mapobjectuses_newInstances []*MapObjectUse
	var mapobjectuses_deletedInstances []*MapObjectUse

	// parse all staged instances and check if they have a reference
	for mapobjectuse := range stage.MapObjectUses {
		if ref, ok := stage.MapObjectUses_reference[mapobjectuse]; !ok {
			mapobjectuses_newInstances = append(mapobjectuses_newInstances, mapobjectuse)
			newInstancesSlice = append(newInstancesSlice, mapobjectuse.GongMarshallIdentifier(stage))
			if stage.MapObjectUses_referenceOrder == nil {
				stage.MapObjectUses_referenceOrder = make(map[*MapObjectUse]uint)
			}
			stage.MapObjectUses_referenceOrder[mapobjectuse] = stage.MapObjectUse_stagedOrder[mapobjectuse]
			newInstancesReverseSlice = append(newInstancesReverseSlice, mapobjectuse.GongMarshallUnstaging(stage))
			// delete(stage.MapObjectUses_referenceOrder, mapobjectuse)
			fieldInitializers, pointersInitializations := mapobjectuse.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MapObjectUse_stagedOrder[ref] = stage.MapObjectUse_stagedOrder[mapobjectuse]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := mapobjectuse.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, mapobjectuse)
			// delete(stage.MapObjectUse_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if mapobjectuse.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", mapobjectuse.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.MapObjectUses_reference {
		instance := stage.MapObjectUses_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.MapObjectUses[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			mapobjectuses_deletedInstances = append(mapobjectuses_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(mapobjectuses_newInstances)
	lenDeletedInstances += len(mapobjectuses_deletedInstances)
	var parameters_newInstances []*Parameter
	var parameters_deletedInstances []*Parameter

	// parse all staged instances and check if they have a reference
	for parameter := range stage.Parameters {
		if ref, ok := stage.Parameters_reference[parameter]; !ok {
			parameters_newInstances = append(parameters_newInstances, parameter)
			newInstancesSlice = append(newInstancesSlice, parameter.GongMarshallIdentifier(stage))
			if stage.Parameters_referenceOrder == nil {
				stage.Parameters_referenceOrder = make(map[*Parameter]uint)
			}
			stage.Parameters_referenceOrder[parameter] = stage.Parameter_stagedOrder[parameter]
			newInstancesReverseSlice = append(newInstancesReverseSlice, parameter.GongMarshallUnstaging(stage))
			// delete(stage.Parameters_referenceOrder, parameter)
			fieldInitializers, pointersInitializations := parameter.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Parameter_stagedOrder[ref] = stage.Parameter_stagedOrder[parameter]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := parameter.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, parameter)
			// delete(stage.Parameter_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if parameter.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", parameter.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Parameters_reference {
		instance := stage.Parameters_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Parameters[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			parameters_deletedInstances = append(parameters_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(parameters_newInstances)
	lenDeletedInstances += len(parameters_deletedInstances)
	var parametercategorys_newInstances []*ParameterCategory
	var parametercategorys_deletedInstances []*ParameterCategory

	// parse all staged instances and check if they have a reference
	for parametercategory := range stage.ParameterCategorys {
		if ref, ok := stage.ParameterCategorys_reference[parametercategory]; !ok {
			parametercategorys_newInstances = append(parametercategorys_newInstances, parametercategory)
			newInstancesSlice = append(newInstancesSlice, parametercategory.GongMarshallIdentifier(stage))
			if stage.ParameterCategorys_referenceOrder == nil {
				stage.ParameterCategorys_referenceOrder = make(map[*ParameterCategory]uint)
			}
			stage.ParameterCategorys_referenceOrder[parametercategory] = stage.ParameterCategory_stagedOrder[parametercategory]
			newInstancesReverseSlice = append(newInstancesReverseSlice, parametercategory.GongMarshallUnstaging(stage))
			// delete(stage.ParameterCategorys_referenceOrder, parametercategory)
			fieldInitializers, pointersInitializations := parametercategory.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ParameterCategory_stagedOrder[ref] = stage.ParameterCategory_stagedOrder[parametercategory]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := parametercategory.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, parametercategory)
			// delete(stage.ParameterCategory_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if parametercategory.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", parametercategory.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ParameterCategorys_reference {
		instance := stage.ParameterCategorys_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ParameterCategorys[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			parametercategorys_deletedInstances = append(parametercategorys_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(parametercategorys_newInstances)
	lenDeletedInstances += len(parametercategorys_deletedInstances)
	var parametercategoryuses_newInstances []*ParameterCategoryUse
	var parametercategoryuses_deletedInstances []*ParameterCategoryUse

	// parse all staged instances and check if they have a reference
	for parametercategoryuse := range stage.ParameterCategoryUses {
		if ref, ok := stage.ParameterCategoryUses_reference[parametercategoryuse]; !ok {
			parametercategoryuses_newInstances = append(parametercategoryuses_newInstances, parametercategoryuse)
			newInstancesSlice = append(newInstancesSlice, parametercategoryuse.GongMarshallIdentifier(stage))
			if stage.ParameterCategoryUses_referenceOrder == nil {
				stage.ParameterCategoryUses_referenceOrder = make(map[*ParameterCategoryUse]uint)
			}
			stage.ParameterCategoryUses_referenceOrder[parametercategoryuse] = stage.ParameterCategoryUse_stagedOrder[parametercategoryuse]
			newInstancesReverseSlice = append(newInstancesReverseSlice, parametercategoryuse.GongMarshallUnstaging(stage))
			// delete(stage.ParameterCategoryUses_referenceOrder, parametercategoryuse)
			fieldInitializers, pointersInitializations := parametercategoryuse.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ParameterCategoryUse_stagedOrder[ref] = stage.ParameterCategoryUse_stagedOrder[parametercategoryuse]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := parametercategoryuse.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, parametercategoryuse)
			// delete(stage.ParameterCategoryUse_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if parametercategoryuse.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", parametercategoryuse.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ParameterCategoryUses_reference {
		instance := stage.ParameterCategoryUses_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ParameterCategoryUses[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			parametercategoryuses_deletedInstances = append(parametercategoryuses_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(parametercategoryuses_newInstances)
	lenDeletedInstances += len(parametercategoryuses_deletedInstances)
	var parametershapes_newInstances []*ParameterShape
	var parametershapes_deletedInstances []*ParameterShape

	// parse all staged instances and check if they have a reference
	for parametershape := range stage.ParameterShapes {
		if ref, ok := stage.ParameterShapes_reference[parametershape]; !ok {
			parametershapes_newInstances = append(parametershapes_newInstances, parametershape)
			newInstancesSlice = append(newInstancesSlice, parametershape.GongMarshallIdentifier(stage))
			if stage.ParameterShapes_referenceOrder == nil {
				stage.ParameterShapes_referenceOrder = make(map[*ParameterShape]uint)
			}
			stage.ParameterShapes_referenceOrder[parametershape] = stage.ParameterShape_stagedOrder[parametershape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, parametershape.GongMarshallUnstaging(stage))
			// delete(stage.ParameterShapes_referenceOrder, parametershape)
			fieldInitializers, pointersInitializations := parametershape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ParameterShape_stagedOrder[ref] = stage.ParameterShape_stagedOrder[parametershape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := parametershape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, parametershape)
			// delete(stage.ParameterShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if parametershape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", parametershape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ParameterShapes_reference {
		instance := stage.ParameterShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ParameterShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			parametershapes_deletedInstances = append(parametershapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(parametershapes_newInstances)
	lenDeletedInstances += len(parametershapes_deletedInstances)
	var parametersaggregates_newInstances []*ParametersAggregate
	var parametersaggregates_deletedInstances []*ParametersAggregate

	// parse all staged instances and check if they have a reference
	for parametersaggregate := range stage.ParametersAggregates {
		if ref, ok := stage.ParametersAggregates_reference[parametersaggregate]; !ok {
			parametersaggregates_newInstances = append(parametersaggregates_newInstances, parametersaggregate)
			newInstancesSlice = append(newInstancesSlice, parametersaggregate.GongMarshallIdentifier(stage))
			if stage.ParametersAggregates_referenceOrder == nil {
				stage.ParametersAggregates_referenceOrder = make(map[*ParametersAggregate]uint)
			}
			stage.ParametersAggregates_referenceOrder[parametersaggregate] = stage.ParametersAggregate_stagedOrder[parametersaggregate]
			newInstancesReverseSlice = append(newInstancesReverseSlice, parametersaggregate.GongMarshallUnstaging(stage))
			// delete(stage.ParametersAggregates_referenceOrder, parametersaggregate)
			fieldInitializers, pointersInitializations := parametersaggregate.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ParametersAggregate_stagedOrder[ref] = stage.ParametersAggregate_stagedOrder[parametersaggregate]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := parametersaggregate.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, parametersaggregate)
			// delete(stage.ParametersAggregate_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if parametersaggregate.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", parametersaggregate.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ParametersAggregates_reference {
		instance := stage.ParametersAggregates_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ParametersAggregates[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			parametersaggregates_deletedInstances = append(parametersaggregates_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(parametersaggregates_newInstances)
	lenDeletedInstances += len(parametersaggregates_deletedInstances)
	var parametersaggregateshapes_newInstances []*ParametersAggregateShape
	var parametersaggregateshapes_deletedInstances []*ParametersAggregateShape

	// parse all staged instances and check if they have a reference
	for parametersaggregateshape := range stage.ParametersAggregateShapes {
		if ref, ok := stage.ParametersAggregateShapes_reference[parametersaggregateshape]; !ok {
			parametersaggregateshapes_newInstances = append(parametersaggregateshapes_newInstances, parametersaggregateshape)
			newInstancesSlice = append(newInstancesSlice, parametersaggregateshape.GongMarshallIdentifier(stage))
			if stage.ParametersAggregateShapes_referenceOrder == nil {
				stage.ParametersAggregateShapes_referenceOrder = make(map[*ParametersAggregateShape]uint)
			}
			stage.ParametersAggregateShapes_referenceOrder[parametersaggregateshape] = stage.ParametersAggregateShape_stagedOrder[parametersaggregateshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, parametersaggregateshape.GongMarshallUnstaging(stage))
			// delete(stage.ParametersAggregateShapes_referenceOrder, parametersaggregateshape)
			fieldInitializers, pointersInitializations := parametersaggregateshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ParametersAggregateShape_stagedOrder[ref] = stage.ParametersAggregateShape_stagedOrder[parametersaggregateshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := parametersaggregateshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, parametersaggregateshape)
			// delete(stage.ParametersAggregateShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if parametersaggregateshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", parametersaggregateshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ParametersAggregateShapes_reference {
		instance := stage.ParametersAggregateShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ParametersAggregateShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			parametersaggregateshapes_deletedInstances = append(parametersaggregateshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(parametersaggregateshapes_newInstances)
	lenDeletedInstances += len(parametersaggregateshapes_deletedInstances)
	var positions_newInstances []*Position
	var positions_deletedInstances []*Position

	// parse all staged instances and check if they have a reference
	for position := range stage.Positions {
		if ref, ok := stage.Positions_reference[position]; !ok {
			positions_newInstances = append(positions_newInstances, position)
			newInstancesSlice = append(newInstancesSlice, position.GongMarshallIdentifier(stage))
			if stage.Positions_referenceOrder == nil {
				stage.Positions_referenceOrder = make(map[*Position]uint)
			}
			stage.Positions_referenceOrder[position] = stage.Position_stagedOrder[position]
			newInstancesReverseSlice = append(newInstancesReverseSlice, position.GongMarshallUnstaging(stage))
			// delete(stage.Positions_referenceOrder, position)
			fieldInitializers, pointersInitializations := position.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Position_stagedOrder[ref] = stage.Position_stagedOrder[position]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := position.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, position)
			// delete(stage.Position_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if position.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", position.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Positions_reference {
		instance := stage.Positions_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Positions[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			positions_deletedInstances = append(positions_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(positions_newInstances)
	lenDeletedInstances += len(positions_deletedInstances)
	var repositorys_newInstances []*Repository
	var repositorys_deletedInstances []*Repository

	// parse all staged instances and check if they have a reference
	for repository := range stage.Repositorys {
		if ref, ok := stage.Repositorys_reference[repository]; !ok {
			repositorys_newInstances = append(repositorys_newInstances, repository)
			newInstancesSlice = append(newInstancesSlice, repository.GongMarshallIdentifier(stage))
			if stage.Repositorys_referenceOrder == nil {
				stage.Repositorys_referenceOrder = make(map[*Repository]uint)
			}
			stage.Repositorys_referenceOrder[repository] = stage.Repository_stagedOrder[repository]
			newInstancesReverseSlice = append(newInstancesReverseSlice, repository.GongMarshallUnstaging(stage))
			// delete(stage.Repositorys_referenceOrder, repository)
			fieldInitializers, pointersInitializations := repository.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Repository_stagedOrder[ref] = stage.Repository_stagedOrder[repository]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := repository.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, repository)
			// delete(stage.Repository_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if repository.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", repository.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Repositorys_reference {
		instance := stage.Repositorys_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Repositorys[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			repositorys_deletedInstances = append(repositorys_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(repositorys_newInstances)
	lenDeletedInstances += len(repositorys_deletedInstances)
	var scenarios_newInstances []*Scenario
	var scenarios_deletedInstances []*Scenario

	// parse all staged instances and check if they have a reference
	for scenario := range stage.Scenarios {
		if ref, ok := stage.Scenarios_reference[scenario]; !ok {
			scenarios_newInstances = append(scenarios_newInstances, scenario)
			newInstancesSlice = append(newInstancesSlice, scenario.GongMarshallIdentifier(stage))
			if stage.Scenarios_referenceOrder == nil {
				stage.Scenarios_referenceOrder = make(map[*Scenario]uint)
			}
			stage.Scenarios_referenceOrder[scenario] = stage.Scenario_stagedOrder[scenario]
			newInstancesReverseSlice = append(newInstancesReverseSlice, scenario.GongMarshallUnstaging(stage))
			// delete(stage.Scenarios_referenceOrder, scenario)
			fieldInitializers, pointersInitializations := scenario.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Scenario_stagedOrder[ref] = stage.Scenario_stagedOrder[scenario]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := scenario.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, scenario)
			// delete(stage.Scenario_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if scenario.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", scenario.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Scenarios_reference {
		instance := stage.Scenarios_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Scenarios[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			scenarios_deletedInstances = append(scenarios_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(scenarios_newInstances)
	lenDeletedInstances += len(scenarios_deletedInstances)
	var users_newInstances []*User
	var users_deletedInstances []*User

	// parse all staged instances and check if they have a reference
	for user := range stage.Users {
		if ref, ok := stage.Users_reference[user]; !ok {
			users_newInstances = append(users_newInstances, user)
			newInstancesSlice = append(newInstancesSlice, user.GongMarshallIdentifier(stage))
			if stage.Users_referenceOrder == nil {
				stage.Users_referenceOrder = make(map[*User]uint)
			}
			stage.Users_referenceOrder[user] = stage.User_stagedOrder[user]
			newInstancesReverseSlice = append(newInstancesReverseSlice, user.GongMarshallUnstaging(stage))
			// delete(stage.Users_referenceOrder, user)
			fieldInitializers, pointersInitializations := user.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.User_stagedOrder[ref] = stage.User_stagedOrder[user]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := user.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, user)
			// delete(stage.User_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if user.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", user.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Users_reference {
		instance := stage.Users_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Users[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			users_deletedInstances = append(users_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(users_newInstances)
	lenDeletedInstances += len(users_deletedInstances)
	var useruses_newInstances []*UserUse
	var useruses_deletedInstances []*UserUse

	// parse all staged instances and check if they have a reference
	for useruse := range stage.UserUses {
		if ref, ok := stage.UserUses_reference[useruse]; !ok {
			useruses_newInstances = append(useruses_newInstances, useruse)
			newInstancesSlice = append(newInstancesSlice, useruse.GongMarshallIdentifier(stage))
			if stage.UserUses_referenceOrder == nil {
				stage.UserUses_referenceOrder = make(map[*UserUse]uint)
			}
			stage.UserUses_referenceOrder[useruse] = stage.UserUse_stagedOrder[useruse]
			newInstancesReverseSlice = append(newInstancesReverseSlice, useruse.GongMarshallUnstaging(stage))
			// delete(stage.UserUses_referenceOrder, useruse)
			fieldInitializers, pointersInitializations := useruse.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.UserUse_stagedOrder[ref] = stage.UserUse_stagedOrder[useruse]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := useruse.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, useruse)
			// delete(stage.UserUse_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if useruse.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", useruse.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.UserUses_reference {
		instance := stage.UserUses_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.UserUses[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			useruses_deletedInstances = append(useruses_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(useruses_newInstances)
	lenDeletedInstances += len(useruses_deletedInstances)
	var workspaces_newInstances []*Workspace
	var workspaces_deletedInstances []*Workspace

	// parse all staged instances and check if they have a reference
	for workspace := range stage.Workspaces {
		if ref, ok := stage.Workspaces_reference[workspace]; !ok {
			workspaces_newInstances = append(workspaces_newInstances, workspace)
			newInstancesSlice = append(newInstancesSlice, workspace.GongMarshallIdentifier(stage))
			if stage.Workspaces_referenceOrder == nil {
				stage.Workspaces_referenceOrder = make(map[*Workspace]uint)
			}
			stage.Workspaces_referenceOrder[workspace] = stage.Workspace_stagedOrder[workspace]
			newInstancesReverseSlice = append(newInstancesReverseSlice, workspace.GongMarshallUnstaging(stage))
			// delete(stage.Workspaces_referenceOrder, workspace)
			fieldInitializers, pointersInitializations := workspace.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Workspace_stagedOrder[ref] = stage.Workspace_stagedOrder[workspace]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := workspace.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, workspace)
			// delete(stage.Workspace_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if workspace.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", workspace.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Workspaces_reference {
		instance := stage.Workspaces_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Workspaces[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			workspaces_deletedInstances = append(workspaces_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(workspaces_newInstances)
	lenDeletedInstances += len(workspaces_deletedInstances)

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
	stage.ActorStates_reference = make(map[*ActorState]*ActorState)
	stage.ActorStates_referenceOrder = make(map[*ActorState]uint) // diff Unstage needs the reference order
	stage.ActorStates_instance = make(map[*ActorState]*ActorState)
	for instance := range stage.ActorStates {
		_copy := instance.GongCopy().(*ActorState)
		stage.ActorStates_reference[instance] = _copy
		stage.ActorStates_instance[_copy] = instance
		stage.ActorStates_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ActorStateShapes_reference = make(map[*ActorStateShape]*ActorStateShape)
	stage.ActorStateShapes_referenceOrder = make(map[*ActorStateShape]uint) // diff Unstage needs the reference order
	stage.ActorStateShapes_instance = make(map[*ActorStateShape]*ActorStateShape)
	for instance := range stage.ActorStateShapes {
		_copy := instance.GongCopy().(*ActorStateShape)
		stage.ActorStateShapes_reference[instance] = _copy
		stage.ActorStateShapes_instance[_copy] = instance
		stage.ActorStateShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ActorStateTransitions_reference = make(map[*ActorStateTransition]*ActorStateTransition)
	stage.ActorStateTransitions_referenceOrder = make(map[*ActorStateTransition]uint) // diff Unstage needs the reference order
	stage.ActorStateTransitions_instance = make(map[*ActorStateTransition]*ActorStateTransition)
	for instance := range stage.ActorStateTransitions {
		_copy := instance.GongCopy().(*ActorStateTransition)
		stage.ActorStateTransitions_reference[instance] = _copy
		stage.ActorStateTransitions_instance[_copy] = instance
		stage.ActorStateTransitions_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ActorStateTransitionShapes_reference = make(map[*ActorStateTransitionShape]*ActorStateTransitionShape)
	stage.ActorStateTransitionShapes_referenceOrder = make(map[*ActorStateTransitionShape]uint) // diff Unstage needs the reference order
	stage.ActorStateTransitionShapes_instance = make(map[*ActorStateTransitionShape]*ActorStateTransitionShape)
	for instance := range stage.ActorStateTransitionShapes {
		_copy := instance.GongCopy().(*ActorStateTransitionShape)
		stage.ActorStateTransitionShapes_reference[instance] = _copy
		stage.ActorStateTransitionShapes_instance[_copy] = instance
		stage.ActorStateTransitionShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Analysiss_reference = make(map[*Analysis]*Analysis)
	stage.Analysiss_referenceOrder = make(map[*Analysis]uint) // diff Unstage needs the reference order
	stage.Analysiss_instance = make(map[*Analysis]*Analysis)
	for instance := range stage.Analysiss {
		_copy := instance.GongCopy().(*Analysis)
		stage.Analysiss_reference[instance] = _copy
		stage.Analysiss_instance[_copy] = instance
		stage.Analysiss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ControlPointShapes_reference = make(map[*ControlPointShape]*ControlPointShape)
	stage.ControlPointShapes_referenceOrder = make(map[*ControlPointShape]uint) // diff Unstage needs the reference order
	stage.ControlPointShapes_instance = make(map[*ControlPointShape]*ControlPointShape)
	for instance := range stage.ControlPointShapes {
		_copy := instance.GongCopy().(*ControlPointShape)
		stage.ControlPointShapes_reference[instance] = _copy
		stage.ControlPointShapes_instance[_copy] = instance
		stage.ControlPointShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Diagrams_reference = make(map[*Diagram]*Diagram)
	stage.Diagrams_referenceOrder = make(map[*Diagram]uint) // diff Unstage needs the reference order
	stage.Diagrams_instance = make(map[*Diagram]*Diagram)
	for instance := range stage.Diagrams {
		_copy := instance.GongCopy().(*Diagram)
		stage.Diagrams_reference[instance] = _copy
		stage.Diagrams_instance[_copy] = instance
		stage.Diagrams_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Documents_reference = make(map[*Document]*Document)
	stage.Documents_referenceOrder = make(map[*Document]uint) // diff Unstage needs the reference order
	stage.Documents_instance = make(map[*Document]*Document)
	for instance := range stage.Documents {
		_copy := instance.GongCopy().(*Document)
		stage.Documents_reference[instance] = _copy
		stage.Documents_instance[_copy] = instance
		stage.Documents_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DocumentUses_reference = make(map[*DocumentUse]*DocumentUse)
	stage.DocumentUses_referenceOrder = make(map[*DocumentUse]uint) // diff Unstage needs the reference order
	stage.DocumentUses_instance = make(map[*DocumentUse]*DocumentUse)
	for instance := range stage.DocumentUses {
		_copy := instance.GongCopy().(*DocumentUse)
		stage.DocumentUses_reference[instance] = _copy
		stage.DocumentUses_instance[_copy] = instance
		stage.DocumentUses_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.EvolutionDirections_reference = make(map[*EvolutionDirection]*EvolutionDirection)
	stage.EvolutionDirections_referenceOrder = make(map[*EvolutionDirection]uint) // diff Unstage needs the reference order
	stage.EvolutionDirections_instance = make(map[*EvolutionDirection]*EvolutionDirection)
	for instance := range stage.EvolutionDirections {
		_copy := instance.GongCopy().(*EvolutionDirection)
		stage.EvolutionDirections_reference[instance] = _copy
		stage.EvolutionDirections_instance[_copy] = instance
		stage.EvolutionDirections_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.EvolutionDirectionShapes_reference = make(map[*EvolutionDirectionShape]*EvolutionDirectionShape)
	stage.EvolutionDirectionShapes_referenceOrder = make(map[*EvolutionDirectionShape]uint) // diff Unstage needs the reference order
	stage.EvolutionDirectionShapes_instance = make(map[*EvolutionDirectionShape]*EvolutionDirectionShape)
	for instance := range stage.EvolutionDirectionShapes {
		_copy := instance.GongCopy().(*EvolutionDirectionShape)
		stage.EvolutionDirectionShapes_reference[instance] = _copy
		stage.EvolutionDirectionShapes_instance[_copy] = instance
		stage.EvolutionDirectionShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Foos_reference = make(map[*Foo]*Foo)
	stage.Foos_referenceOrder = make(map[*Foo]uint) // diff Unstage needs the reference order
	stage.Foos_instance = make(map[*Foo]*Foo)
	for instance := range stage.Foos {
		_copy := instance.GongCopy().(*Foo)
		stage.Foos_reference[instance] = _copy
		stage.Foos_instance[_copy] = instance
		stage.Foos_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GeoObjects_reference = make(map[*GeoObject]*GeoObject)
	stage.GeoObjects_referenceOrder = make(map[*GeoObject]uint) // diff Unstage needs the reference order
	stage.GeoObjects_instance = make(map[*GeoObject]*GeoObject)
	for instance := range stage.GeoObjects {
		_copy := instance.GongCopy().(*GeoObject)
		stage.GeoObjects_reference[instance] = _copy
		stage.GeoObjects_instance[_copy] = instance
		stage.GeoObjects_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GeoObjectUses_reference = make(map[*GeoObjectUse]*GeoObjectUse)
	stage.GeoObjectUses_referenceOrder = make(map[*GeoObjectUse]uint) // diff Unstage needs the reference order
	stage.GeoObjectUses_instance = make(map[*GeoObjectUse]*GeoObjectUse)
	for instance := range stage.GeoObjectUses {
		_copy := instance.GongCopy().(*GeoObjectUse)
		stage.GeoObjectUses_reference[instance] = _copy
		stage.GeoObjectUses_instance[_copy] = instance
		stage.GeoObjectUses_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Groups_reference = make(map[*Group]*Group)
	stage.Groups_referenceOrder = make(map[*Group]uint) // diff Unstage needs the reference order
	stage.Groups_instance = make(map[*Group]*Group)
	for instance := range stage.Groups {
		_copy := instance.GongCopy().(*Group)
		stage.Groups_reference[instance] = _copy
		stage.Groups_instance[_copy] = instance
		stage.Groups_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GroupUses_reference = make(map[*GroupUse]*GroupUse)
	stage.GroupUses_referenceOrder = make(map[*GroupUse]uint) // diff Unstage needs the reference order
	stage.GroupUses_instance = make(map[*GroupUse]*GroupUse)
	for instance := range stage.GroupUses {
		_copy := instance.GongCopy().(*GroupUse)
		stage.GroupUses_reference[instance] = _copy
		stage.GroupUses_instance[_copy] = instance
		stage.GroupUses_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Librarys_reference = make(map[*Library]*Library)
	stage.Librarys_referenceOrder = make(map[*Library]uint) // diff Unstage needs the reference order
	stage.Librarys_instance = make(map[*Library]*Library)
	for instance := range stage.Librarys {
		_copy := instance.GongCopy().(*Library)
		stage.Librarys_reference[instance] = _copy
		stage.Librarys_instance[_copy] = instance
		stage.Librarys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.MapObjects_reference = make(map[*MapObject]*MapObject)
	stage.MapObjects_referenceOrder = make(map[*MapObject]uint) // diff Unstage needs the reference order
	stage.MapObjects_instance = make(map[*MapObject]*MapObject)
	for instance := range stage.MapObjects {
		_copy := instance.GongCopy().(*MapObject)
		stage.MapObjects_reference[instance] = _copy
		stage.MapObjects_instance[_copy] = instance
		stage.MapObjects_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.MapObjectUses_reference = make(map[*MapObjectUse]*MapObjectUse)
	stage.MapObjectUses_referenceOrder = make(map[*MapObjectUse]uint) // diff Unstage needs the reference order
	stage.MapObjectUses_instance = make(map[*MapObjectUse]*MapObjectUse)
	for instance := range stage.MapObjectUses {
		_copy := instance.GongCopy().(*MapObjectUse)
		stage.MapObjectUses_reference[instance] = _copy
		stage.MapObjectUses_instance[_copy] = instance
		stage.MapObjectUses_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Parameters_reference = make(map[*Parameter]*Parameter)
	stage.Parameters_referenceOrder = make(map[*Parameter]uint) // diff Unstage needs the reference order
	stage.Parameters_instance = make(map[*Parameter]*Parameter)
	for instance := range stage.Parameters {
		_copy := instance.GongCopy().(*Parameter)
		stage.Parameters_reference[instance] = _copy
		stage.Parameters_instance[_copy] = instance
		stage.Parameters_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ParameterCategorys_reference = make(map[*ParameterCategory]*ParameterCategory)
	stage.ParameterCategorys_referenceOrder = make(map[*ParameterCategory]uint) // diff Unstage needs the reference order
	stage.ParameterCategorys_instance = make(map[*ParameterCategory]*ParameterCategory)
	for instance := range stage.ParameterCategorys {
		_copy := instance.GongCopy().(*ParameterCategory)
		stage.ParameterCategorys_reference[instance] = _copy
		stage.ParameterCategorys_instance[_copy] = instance
		stage.ParameterCategorys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ParameterCategoryUses_reference = make(map[*ParameterCategoryUse]*ParameterCategoryUse)
	stage.ParameterCategoryUses_referenceOrder = make(map[*ParameterCategoryUse]uint) // diff Unstage needs the reference order
	stage.ParameterCategoryUses_instance = make(map[*ParameterCategoryUse]*ParameterCategoryUse)
	for instance := range stage.ParameterCategoryUses {
		_copy := instance.GongCopy().(*ParameterCategoryUse)
		stage.ParameterCategoryUses_reference[instance] = _copy
		stage.ParameterCategoryUses_instance[_copy] = instance
		stage.ParameterCategoryUses_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ParameterShapes_reference = make(map[*ParameterShape]*ParameterShape)
	stage.ParameterShapes_referenceOrder = make(map[*ParameterShape]uint) // diff Unstage needs the reference order
	stage.ParameterShapes_instance = make(map[*ParameterShape]*ParameterShape)
	for instance := range stage.ParameterShapes {
		_copy := instance.GongCopy().(*ParameterShape)
		stage.ParameterShapes_reference[instance] = _copy
		stage.ParameterShapes_instance[_copy] = instance
		stage.ParameterShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ParametersAggregates_reference = make(map[*ParametersAggregate]*ParametersAggregate)
	stage.ParametersAggregates_referenceOrder = make(map[*ParametersAggregate]uint) // diff Unstage needs the reference order
	stage.ParametersAggregates_instance = make(map[*ParametersAggregate]*ParametersAggregate)
	for instance := range stage.ParametersAggregates {
		_copy := instance.GongCopy().(*ParametersAggregate)
		stage.ParametersAggregates_reference[instance] = _copy
		stage.ParametersAggregates_instance[_copy] = instance
		stage.ParametersAggregates_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ParametersAggregateShapes_reference = make(map[*ParametersAggregateShape]*ParametersAggregateShape)
	stage.ParametersAggregateShapes_referenceOrder = make(map[*ParametersAggregateShape]uint) // diff Unstage needs the reference order
	stage.ParametersAggregateShapes_instance = make(map[*ParametersAggregateShape]*ParametersAggregateShape)
	for instance := range stage.ParametersAggregateShapes {
		_copy := instance.GongCopy().(*ParametersAggregateShape)
		stage.ParametersAggregateShapes_reference[instance] = _copy
		stage.ParametersAggregateShapes_instance[_copy] = instance
		stage.ParametersAggregateShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Positions_reference = make(map[*Position]*Position)
	stage.Positions_referenceOrder = make(map[*Position]uint) // diff Unstage needs the reference order
	stage.Positions_instance = make(map[*Position]*Position)
	for instance := range stage.Positions {
		_copy := instance.GongCopy().(*Position)
		stage.Positions_reference[instance] = _copy
		stage.Positions_instance[_copy] = instance
		stage.Positions_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Repositorys_reference = make(map[*Repository]*Repository)
	stage.Repositorys_referenceOrder = make(map[*Repository]uint) // diff Unstage needs the reference order
	stage.Repositorys_instance = make(map[*Repository]*Repository)
	for instance := range stage.Repositorys {
		_copy := instance.GongCopy().(*Repository)
		stage.Repositorys_reference[instance] = _copy
		stage.Repositorys_instance[_copy] = instance
		stage.Repositorys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Scenarios_reference = make(map[*Scenario]*Scenario)
	stage.Scenarios_referenceOrder = make(map[*Scenario]uint) // diff Unstage needs the reference order
	stage.Scenarios_instance = make(map[*Scenario]*Scenario)
	for instance := range stage.Scenarios {
		_copy := instance.GongCopy().(*Scenario)
		stage.Scenarios_reference[instance] = _copy
		stage.Scenarios_instance[_copy] = instance
		stage.Scenarios_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Users_reference = make(map[*User]*User)
	stage.Users_referenceOrder = make(map[*User]uint) // diff Unstage needs the reference order
	stage.Users_instance = make(map[*User]*User)
	for instance := range stage.Users {
		_copy := instance.GongCopy().(*User)
		stage.Users_reference[instance] = _copy
		stage.Users_instance[_copy] = instance
		stage.Users_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.UserUses_reference = make(map[*UserUse]*UserUse)
	stage.UserUses_referenceOrder = make(map[*UserUse]uint) // diff Unstage needs the reference order
	stage.UserUses_instance = make(map[*UserUse]*UserUse)
	for instance := range stage.UserUses {
		_copy := instance.GongCopy().(*UserUse)
		stage.UserUses_reference[instance] = _copy
		stage.UserUses_instance[_copy] = instance
		stage.UserUses_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Workspaces_reference = make(map[*Workspace]*Workspace)
	stage.Workspaces_referenceOrder = make(map[*Workspace]uint) // diff Unstage needs the reference order
	stage.Workspaces_instance = make(map[*Workspace]*Workspace)
	for instance := range stage.Workspaces {
		_copy := instance.GongCopy().(*Workspace)
		stage.Workspaces_reference[instance] = _copy
		stage.Workspaces_instance[_copy] = instance
		stage.Workspaces_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.ActorStates {
		reference := stage.ActorStates_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ActorStateShapes {
		reference := stage.ActorStateShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ActorStateTransitions {
		reference := stage.ActorStateTransitions_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ActorStateTransitionShapes {
		reference := stage.ActorStateTransitionShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Analysiss {
		reference := stage.Analysiss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ControlPointShapes {
		reference := stage.ControlPointShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Diagrams {
		reference := stage.Diagrams_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Documents {
		reference := stage.Documents_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DocumentUses {
		reference := stage.DocumentUses_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.EvolutionDirections {
		reference := stage.EvolutionDirections_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.EvolutionDirectionShapes {
		reference := stage.EvolutionDirectionShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Foos {
		reference := stage.Foos_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GeoObjects {
		reference := stage.GeoObjects_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GeoObjectUses {
		reference := stage.GeoObjectUses_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Groups {
		reference := stage.Groups_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GroupUses {
		reference := stage.GroupUses_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Librarys {
		reference := stage.Librarys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.MapObjects {
		reference := stage.MapObjects_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.MapObjectUses {
		reference := stage.MapObjectUses_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Parameters {
		reference := stage.Parameters_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ParameterCategorys {
		reference := stage.ParameterCategorys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ParameterCategoryUses {
		reference := stage.ParameterCategoryUses_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ParameterShapes {
		reference := stage.ParameterShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ParametersAggregates {
		reference := stage.ParametersAggregates_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ParametersAggregateShapes {
		reference := stage.ParametersAggregateShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Positions {
		reference := stage.Positions_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Repositorys {
		reference := stage.Repositorys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Scenarios {
		reference := stage.Scenarios_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Users {
		reference := stage.Users_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.UserUses {
		reference := stage.UserUses_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Workspaces {
		reference := stage.Workspaces_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

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
	if order, ok := stage.ActorState_stagedOrder[actorstate]; ok {
		return order
	}
	if order, ok := stage.ActorStates_referenceOrder[actorstate]; ok {
		return order
	} else {
		log.Printf("instance %p of type ActorState was not staged and does not have a reference order", actorstate)
		return 0
	}
}

func (actorstateshape *ActorStateShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ActorStateShape_stagedOrder[actorstateshape]; ok {
		return order
	}
	if order, ok := stage.ActorStateShapes_referenceOrder[actorstateshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ActorStateShape was not staged and does not have a reference order", actorstateshape)
		return 0
	}
}

func (actorstatetransition *ActorStateTransition) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ActorStateTransition_stagedOrder[actorstatetransition]; ok {
		return order
	}
	if order, ok := stage.ActorStateTransitions_referenceOrder[actorstatetransition]; ok {
		return order
	} else {
		log.Printf("instance %p of type ActorStateTransition was not staged and does not have a reference order", actorstatetransition)
		return 0
	}
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ActorStateTransitionShape_stagedOrder[actorstatetransitionshape]; ok {
		return order
	}
	if order, ok := stage.ActorStateTransitionShapes_referenceOrder[actorstatetransitionshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ActorStateTransitionShape was not staged and does not have a reference order", actorstatetransitionshape)
		return 0
	}
}

func (analysis *Analysis) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Analysis_stagedOrder[analysis]; ok {
		return order
	}
	if order, ok := stage.Analysiss_referenceOrder[analysis]; ok {
		return order
	} else {
		log.Printf("instance %p of type Analysis was not staged and does not have a reference order", analysis)
		return 0
	}
}

func (controlpointshape *ControlPointShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ControlPointShape_stagedOrder[controlpointshape]; ok {
		return order
	}
	if order, ok := stage.ControlPointShapes_referenceOrder[controlpointshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ControlPointShape was not staged and does not have a reference order", controlpointshape)
		return 0
	}
}

func (diagram *Diagram) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Diagram_stagedOrder[diagram]; ok {
		return order
	}
	if order, ok := stage.Diagrams_referenceOrder[diagram]; ok {
		return order
	} else {
		log.Printf("instance %p of type Diagram was not staged and does not have a reference order", diagram)
		return 0
	}
}

func (document *Document) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Document_stagedOrder[document]; ok {
		return order
	}
	if order, ok := stage.Documents_referenceOrder[document]; ok {
		return order
	} else {
		log.Printf("instance %p of type Document was not staged and does not have a reference order", document)
		return 0
	}
}

func (documentuse *DocumentUse) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DocumentUse_stagedOrder[documentuse]; ok {
		return order
	}
	if order, ok := stage.DocumentUses_referenceOrder[documentuse]; ok {
		return order
	} else {
		log.Printf("instance %p of type DocumentUse was not staged and does not have a reference order", documentuse)
		return 0
	}
}

func (evolutiondirection *EvolutionDirection) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EvolutionDirection_stagedOrder[evolutiondirection]; ok {
		return order
	}
	if order, ok := stage.EvolutionDirections_referenceOrder[evolutiondirection]; ok {
		return order
	} else {
		log.Printf("instance %p of type EvolutionDirection was not staged and does not have a reference order", evolutiondirection)
		return 0
	}
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EvolutionDirectionShape_stagedOrder[evolutiondirectionshape]; ok {
		return order
	}
	if order, ok := stage.EvolutionDirectionShapes_referenceOrder[evolutiondirectionshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type EvolutionDirectionShape was not staged and does not have a reference order", evolutiondirectionshape)
		return 0
	}
}

func (foo *Foo) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Foo_stagedOrder[foo]; ok {
		return order
	}
	if order, ok := stage.Foos_referenceOrder[foo]; ok {
		return order
	} else {
		log.Printf("instance %p of type Foo was not staged and does not have a reference order", foo)
		return 0
	}
}

func (geoobject *GeoObject) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GeoObject_stagedOrder[geoobject]; ok {
		return order
	}
	if order, ok := stage.GeoObjects_referenceOrder[geoobject]; ok {
		return order
	} else {
		log.Printf("instance %p of type GeoObject was not staged and does not have a reference order", geoobject)
		return 0
	}
}

func (geoobjectuse *GeoObjectUse) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GeoObjectUse_stagedOrder[geoobjectuse]; ok {
		return order
	}
	if order, ok := stage.GeoObjectUses_referenceOrder[geoobjectuse]; ok {
		return order
	} else {
		log.Printf("instance %p of type GeoObjectUse was not staged and does not have a reference order", geoobjectuse)
		return 0
	}
}

func (group *Group) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Group_stagedOrder[group]; ok {
		return order
	}
	if order, ok := stage.Groups_referenceOrder[group]; ok {
		return order
	} else {
		log.Printf("instance %p of type Group was not staged and does not have a reference order", group)
		return 0
	}
}

func (groupuse *GroupUse) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GroupUse_stagedOrder[groupuse]; ok {
		return order
	}
	if order, ok := stage.GroupUses_referenceOrder[groupuse]; ok {
		return order
	} else {
		log.Printf("instance %p of type GroupUse was not staged and does not have a reference order", groupuse)
		return 0
	}
}

func (library *Library) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Library_stagedOrder[library]; ok {
		return order
	}
	if order, ok := stage.Librarys_referenceOrder[library]; ok {
		return order
	} else {
		log.Printf("instance %p of type Library was not staged and does not have a reference order", library)
		return 0
	}
}

func (mapobject *MapObject) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MapObject_stagedOrder[mapobject]; ok {
		return order
	}
	if order, ok := stage.MapObjects_referenceOrder[mapobject]; ok {
		return order
	} else {
		log.Printf("instance %p of type MapObject was not staged and does not have a reference order", mapobject)
		return 0
	}
}

func (mapobjectuse *MapObjectUse) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MapObjectUse_stagedOrder[mapobjectuse]; ok {
		return order
	}
	if order, ok := stage.MapObjectUses_referenceOrder[mapobjectuse]; ok {
		return order
	} else {
		log.Printf("instance %p of type MapObjectUse was not staged and does not have a reference order", mapobjectuse)
		return 0
	}
}

func (parameter *Parameter) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Parameter_stagedOrder[parameter]; ok {
		return order
	}
	if order, ok := stage.Parameters_referenceOrder[parameter]; ok {
		return order
	} else {
		log.Printf("instance %p of type Parameter was not staged and does not have a reference order", parameter)
		return 0
	}
}

func (parametercategory *ParameterCategory) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ParameterCategory_stagedOrder[parametercategory]; ok {
		return order
	}
	if order, ok := stage.ParameterCategorys_referenceOrder[parametercategory]; ok {
		return order
	} else {
		log.Printf("instance %p of type ParameterCategory was not staged and does not have a reference order", parametercategory)
		return 0
	}
}

func (parametercategoryuse *ParameterCategoryUse) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ParameterCategoryUse_stagedOrder[parametercategoryuse]; ok {
		return order
	}
	if order, ok := stage.ParameterCategoryUses_referenceOrder[parametercategoryuse]; ok {
		return order
	} else {
		log.Printf("instance %p of type ParameterCategoryUse was not staged and does not have a reference order", parametercategoryuse)
		return 0
	}
}

func (parametershape *ParameterShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ParameterShape_stagedOrder[parametershape]; ok {
		return order
	}
	if order, ok := stage.ParameterShapes_referenceOrder[parametershape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ParameterShape was not staged and does not have a reference order", parametershape)
		return 0
	}
}

func (parametersaggregate *ParametersAggregate) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ParametersAggregate_stagedOrder[parametersaggregate]; ok {
		return order
	}
	if order, ok := stage.ParametersAggregates_referenceOrder[parametersaggregate]; ok {
		return order
	} else {
		log.Printf("instance %p of type ParametersAggregate was not staged and does not have a reference order", parametersaggregate)
		return 0
	}
}

func (parametersaggregateshape *ParametersAggregateShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ParametersAggregateShape_stagedOrder[parametersaggregateshape]; ok {
		return order
	}
	if order, ok := stage.ParametersAggregateShapes_referenceOrder[parametersaggregateshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ParametersAggregateShape was not staged and does not have a reference order", parametersaggregateshape)
		return 0
	}
}

func (position *Position) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Position_stagedOrder[position]; ok {
		return order
	}
	if order, ok := stage.Positions_referenceOrder[position]; ok {
		return order
	} else {
		log.Printf("instance %p of type Position was not staged and does not have a reference order", position)
		return 0
	}
}

func (repository *Repository) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Repository_stagedOrder[repository]; ok {
		return order
	}
	if order, ok := stage.Repositorys_referenceOrder[repository]; ok {
		return order
	} else {
		log.Printf("instance %p of type Repository was not staged and does not have a reference order", repository)
		return 0
	}
}

func (scenario *Scenario) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Scenario_stagedOrder[scenario]; ok {
		return order
	}
	if order, ok := stage.Scenarios_referenceOrder[scenario]; ok {
		return order
	} else {
		log.Printf("instance %p of type Scenario was not staged and does not have a reference order", scenario)
		return 0
	}
}

func (user *User) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.User_stagedOrder[user]; ok {
		return order
	}
	if order, ok := stage.Users_referenceOrder[user]; ok {
		return order
	} else {
		log.Printf("instance %p of type User was not staged and does not have a reference order", user)
		return 0
	}
}

func (useruse *UserUse) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.UserUse_stagedOrder[useruse]; ok {
		return order
	}
	if order, ok := stage.UserUses_referenceOrder[useruse]; ok {
		return order
	} else {
		log.Printf("instance %p of type UserUse was not staged and does not have a reference order", useruse)
		return 0
	}
}

func (workspace *Workspace) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Workspace_stagedOrder[workspace]; ok {
		return order
	}
	if order, ok := stage.Workspaces_referenceOrder[workspace]; ok {
		return order
	} else {
		log.Printf("instance %p of type Workspace was not staged and does not have a reference order", workspace)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (actorstate *ActorState) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", actorstate.GongGetGongstructName(), actorstate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (actorstate *ActorState) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", actorstate.GongGetGongstructName(), actorstate.GongGetOrder(stage))
}

func (actorstateshape *ActorStateShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", actorstateshape.GongGetGongstructName(), actorstateshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (actorstateshape *ActorStateShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", actorstateshape.GongGetGongstructName(), actorstateshape.GongGetOrder(stage))
}

func (actorstatetransition *ActorStateTransition) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", actorstatetransition.GongGetGongstructName(), actorstatetransition.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (actorstatetransition *ActorStateTransition) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", actorstatetransition.GongGetGongstructName(), actorstatetransition.GongGetOrder(stage))
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", actorstatetransitionshape.GongGetGongstructName(), actorstatetransitionshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (actorstatetransitionshape *ActorStateTransitionShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", actorstatetransitionshape.GongGetGongstructName(), actorstatetransitionshape.GongGetOrder(stage))
}

func (analysis *Analysis) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", analysis.GongGetGongstructName(), analysis.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (analysis *Analysis) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", analysis.GongGetGongstructName(), analysis.GongGetOrder(stage))
}

func (controlpointshape *ControlPointShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", controlpointshape.GongGetGongstructName(), controlpointshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (controlpointshape *ControlPointShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", controlpointshape.GongGetGongstructName(), controlpointshape.GongGetOrder(stage))
}

func (diagram *Diagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagram.GongGetGongstructName(), diagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagram *Diagram) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagram.GongGetGongstructName(), diagram.GongGetOrder(stage))
}

func (document *Document) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", document.GongGetGongstructName(), document.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (document *Document) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", document.GongGetGongstructName(), document.GongGetOrder(stage))
}

func (documentuse *DocumentUse) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", documentuse.GongGetGongstructName(), documentuse.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (documentuse *DocumentUse) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", documentuse.GongGetGongstructName(), documentuse.GongGetOrder(stage))
}

func (evolutiondirection *EvolutionDirection) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", evolutiondirection.GongGetGongstructName(), evolutiondirection.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (evolutiondirection *EvolutionDirection) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", evolutiondirection.GongGetGongstructName(), evolutiondirection.GongGetOrder(stage))
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", evolutiondirectionshape.GongGetGongstructName(), evolutiondirectionshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (evolutiondirectionshape *EvolutionDirectionShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", evolutiondirectionshape.GongGetGongstructName(), evolutiondirectionshape.GongGetOrder(stage))
}

func (foo *Foo) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", foo.GongGetGongstructName(), foo.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (foo *Foo) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", foo.GongGetGongstructName(), foo.GongGetOrder(stage))
}

func (geoobject *GeoObject) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", geoobject.GongGetGongstructName(), geoobject.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (geoobject *GeoObject) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", geoobject.GongGetGongstructName(), geoobject.GongGetOrder(stage))
}

func (geoobjectuse *GeoObjectUse) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", geoobjectuse.GongGetGongstructName(), geoobjectuse.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (geoobjectuse *GeoObjectUse) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", geoobjectuse.GongGetGongstructName(), geoobjectuse.GongGetOrder(stage))
}

func (group *Group) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group.GongGetGongstructName(), group.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (group *Group) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group.GongGetGongstructName(), group.GongGetOrder(stage))
}

func (groupuse *GroupUse) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", groupuse.GongGetGongstructName(), groupuse.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (groupuse *GroupUse) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", groupuse.GongGetGongstructName(), groupuse.GongGetOrder(stage))
}

func (library *Library) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (library *Library) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

func (mapobject *MapObject) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", mapobject.GongGetGongstructName(), mapobject.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (mapobject *MapObject) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", mapobject.GongGetGongstructName(), mapobject.GongGetOrder(stage))
}

func (mapobjectuse *MapObjectUse) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", mapobjectuse.GongGetGongstructName(), mapobjectuse.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (mapobjectuse *MapObjectUse) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", mapobjectuse.GongGetGongstructName(), mapobjectuse.GongGetOrder(stage))
}

func (parameter *Parameter) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", parameter.GongGetGongstructName(), parameter.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (parameter *Parameter) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", parameter.GongGetGongstructName(), parameter.GongGetOrder(stage))
}

func (parametercategory *ParameterCategory) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", parametercategory.GongGetGongstructName(), parametercategory.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (parametercategory *ParameterCategory) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", parametercategory.GongGetGongstructName(), parametercategory.GongGetOrder(stage))
}

func (parametercategoryuse *ParameterCategoryUse) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", parametercategoryuse.GongGetGongstructName(), parametercategoryuse.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (parametercategoryuse *ParameterCategoryUse) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", parametercategoryuse.GongGetGongstructName(), parametercategoryuse.GongGetOrder(stage))
}

func (parametershape *ParameterShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", parametershape.GongGetGongstructName(), parametershape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (parametershape *ParameterShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", parametershape.GongGetGongstructName(), parametershape.GongGetOrder(stage))
}

func (parametersaggregate *ParametersAggregate) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", parametersaggregate.GongGetGongstructName(), parametersaggregate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (parametersaggregate *ParametersAggregate) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", parametersaggregate.GongGetGongstructName(), parametersaggregate.GongGetOrder(stage))
}

func (parametersaggregateshape *ParametersAggregateShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", parametersaggregateshape.GongGetGongstructName(), parametersaggregateshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (parametersaggregateshape *ParametersAggregateShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", parametersaggregateshape.GongGetGongstructName(), parametersaggregateshape.GongGetOrder(stage))
}

func (position *Position) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", position.GongGetGongstructName(), position.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (position *Position) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", position.GongGetGongstructName(), position.GongGetOrder(stage))
}

func (repository *Repository) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", repository.GongGetGongstructName(), repository.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (repository *Repository) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", repository.GongGetGongstructName(), repository.GongGetOrder(stage))
}

func (scenario *Scenario) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", scenario.GongGetGongstructName(), scenario.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (scenario *Scenario) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", scenario.GongGetGongstructName(), scenario.GongGetOrder(stage))
}

func (user *User) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", user.GongGetGongstructName(), user.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (user *User) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", user.GongGetGongstructName(), user.GongGetOrder(stage))
}

func (useruse *UserUse) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", useruse.GongGetGongstructName(), useruse.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (useruse *UserUse) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", useruse.GongGetGongstructName(), useruse.GongGetOrder(stage))
}

func (workspace *Workspace) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", workspace.GongGetGongstructName(), workspace.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (workspace *Workspace) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", workspace.GongGetGongstructName(), workspace.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (actorstate *ActorState) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", actorstate.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ActorState")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(actorstate.Name))
	return
}

func (actorstateshape *ActorStateShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", actorstateshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ActorStateShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(actorstateshape.Name))
	return
}

func (actorstatetransition *ActorStateTransition) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", actorstatetransition.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ActorStateTransition")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(actorstatetransition.Name))
	return
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", actorstatetransitionshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ActorStateTransitionShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(actorstatetransitionshape.Name))
	return
}

func (analysis *Analysis) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", analysis.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Analysis")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(analysis.Name))
	return
}

func (controlpointshape *ControlPointShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", controlpointshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ControlPointShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(controlpointshape.Name))
	return
}

func (diagram *Diagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Diagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.Name))
	return
}

func (document *Document) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", document.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Document")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(document.Name))
	return
}

func (documentuse *DocumentUse) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", documentuse.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DocumentUse")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(documentuse.Name))
	return
}

func (evolutiondirection *EvolutionDirection) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", evolutiondirection.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EvolutionDirection")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(evolutiondirection.Name))
	return
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", evolutiondirectionshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EvolutionDirectionShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(evolutiondirectionshape.Name))
	return
}

func (foo *Foo) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", foo.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Foo")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(foo.Name))
	return
}

func (geoobject *GeoObject) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", geoobject.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GeoObject")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(geoobject.Name))
	return
}

func (geoobjectuse *GeoObjectUse) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", geoobjectuse.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GeoObjectUse")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(geoobjectuse.Name))
	return
}

func (group *Group) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Group")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(group.Name))
	return
}

func (groupuse *GroupUse) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", groupuse.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GroupUse")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(groupuse.Name))
	return
}

func (library *Library) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Library")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(library.Name))
	return
}

func (mapobject *MapObject) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", mapobject.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MapObject")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(mapobject.Name))
	return
}

func (mapobjectuse *MapObjectUse) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", mapobjectuse.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MapObjectUse")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(mapobjectuse.Name))
	return
}

func (parameter *Parameter) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parameter.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Parameter")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(parameter.Name))
	return
}

func (parametercategory *ParameterCategory) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parametercategory.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParameterCategory")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(parametercategory.Name))
	return
}

func (parametercategoryuse *ParameterCategoryUse) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parametercategoryuse.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParameterCategoryUse")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(parametercategoryuse.Name))
	return
}

func (parametershape *ParameterShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parametershape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParameterShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(parametershape.Name))
	return
}

func (parametersaggregate *ParametersAggregate) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parametersaggregate.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParametersAggregate")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(parametersaggregate.Name))
	return
}

func (parametersaggregateshape *ParametersAggregateShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parametersaggregateshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParametersAggregateShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(parametersaggregateshape.Name))
	return
}

func (position *Position) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", position.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Position")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(position.Name))
	return
}

func (repository *Repository) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", repository.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Repository")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(repository.Name))
	return
}

func (scenario *Scenario) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", scenario.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Scenario")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(scenario.Name))
	return
}

func (user *User) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", user.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "User")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(user.Name))
	return
}

func (useruse *UserUse) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", useruse.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "UserUse")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(useruse.Name))
	return
}

func (workspace *Workspace) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", workspace.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Workspace")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(workspace.Name))
	return
}

// insertion point for unstaging
func (actorstate *ActorState) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", actorstate.GongGetReferenceIdentifier(stage))
	return
}

func (actorstateshape *ActorStateShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", actorstateshape.GongGetReferenceIdentifier(stage))
	return
}

func (actorstatetransition *ActorStateTransition) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", actorstatetransition.GongGetReferenceIdentifier(stage))
	return
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", actorstatetransitionshape.GongGetReferenceIdentifier(stage))
	return
}

func (analysis *Analysis) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", analysis.GongGetReferenceIdentifier(stage))
	return
}

func (controlpointshape *ControlPointShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", controlpointshape.GongGetReferenceIdentifier(stage))
	return
}

func (diagram *Diagram) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagram.GongGetReferenceIdentifier(stage))
	return
}

func (document *Document) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", document.GongGetReferenceIdentifier(stage))
	return
}

func (documentuse *DocumentUse) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", documentuse.GongGetReferenceIdentifier(stage))
	return
}

func (evolutiondirection *EvolutionDirection) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", evolutiondirection.GongGetReferenceIdentifier(stage))
	return
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", evolutiondirectionshape.GongGetReferenceIdentifier(stage))
	return
}

func (foo *Foo) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", foo.GongGetReferenceIdentifier(stage))
	return
}

func (geoobject *GeoObject) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", geoobject.GongGetReferenceIdentifier(stage))
	return
}

func (geoobjectuse *GeoObjectUse) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", geoobjectuse.GongGetReferenceIdentifier(stage))
	return
}

func (group *Group) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group.GongGetReferenceIdentifier(stage))
	return
}

func (groupuse *GroupUse) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", groupuse.GongGetReferenceIdentifier(stage))
	return
}

func (library *Library) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetReferenceIdentifier(stage))
	return
}

func (mapobject *MapObject) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", mapobject.GongGetReferenceIdentifier(stage))
	return
}

func (mapobjectuse *MapObjectUse) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", mapobjectuse.GongGetReferenceIdentifier(stage))
	return
}

func (parameter *Parameter) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parameter.GongGetReferenceIdentifier(stage))
	return
}

func (parametercategory *ParameterCategory) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parametercategory.GongGetReferenceIdentifier(stage))
	return
}

func (parametercategoryuse *ParameterCategoryUse) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parametercategoryuse.GongGetReferenceIdentifier(stage))
	return
}

func (parametershape *ParameterShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parametershape.GongGetReferenceIdentifier(stage))
	return
}

func (parametersaggregate *ParametersAggregate) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parametersaggregate.GongGetReferenceIdentifier(stage))
	return
}

func (parametersaggregateshape *ParametersAggregateShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parametersaggregateshape.GongGetReferenceIdentifier(stage))
	return
}

func (position *Position) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", position.GongGetReferenceIdentifier(stage))
	return
}

func (repository *Repository) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", repository.GongGetReferenceIdentifier(stage))
	return
}

func (scenario *Scenario) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", scenario.GongGetReferenceIdentifier(stage))
	return
}

func (user *User) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", user.GongGetReferenceIdentifier(stage))
	return
}

func (useruse *UserUse) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", useruse.GongGetReferenceIdentifier(stage))
	return
}

func (workspace *Workspace) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", workspace.GongGetReferenceIdentifier(stage))
	return
}

func IntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += IntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
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

// end of template
