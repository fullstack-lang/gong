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
func (actorstate *ActorState) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(actorstate).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(actorstate), uint64(stage.GetOrder(actorstate)))
	return
}

func (actorstateshape *ActorStateShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(actorstateshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(actorstateshape), uint64(stage.GetOrder(actorstateshape)))
	return
}

func (actorstatetransition *ActorStateTransition) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(actorstatetransition).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(actorstatetransition), uint64(stage.GetOrder(actorstatetransition)))
	return
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(actorstatetransitionshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(actorstatetransitionshape), uint64(stage.GetOrder(actorstatetransitionshape)))
	return
}

func (analysis *Analysis) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(analysis).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(analysis), uint64(stage.GetOrder(analysis)))
	return
}

func (controlpointshape *ControlPointShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(controlpointshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(controlpointshape), uint64(stage.GetOrder(controlpointshape)))
	return
}

func (diagram *Diagram) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(diagram).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(diagram), uint64(stage.GetOrder(diagram)))
	return
}

func (document *Document) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(document).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(document), uint64(stage.GetOrder(document)))
	return
}

func (documentuse *DocumentUse) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(documentuse).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(documentuse), uint64(stage.GetOrder(documentuse)))
	return
}

func (evolutiondirection *EvolutionDirection) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(evolutiondirection).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(evolutiondirection), uint64(stage.GetOrder(evolutiondirection)))
	return
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(evolutiondirectionshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(evolutiondirectionshape), uint64(stage.GetOrder(evolutiondirectionshape)))
	return
}

func (foo *Foo) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(foo).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(foo), uint64(stage.GetOrder(foo)))
	return
}

func (geoobject *GeoObject) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(geoobject).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(geoobject), uint64(stage.GetOrder(geoobject)))
	return
}

func (geoobjectuse *GeoObjectUse) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(geoobjectuse).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(geoobjectuse), uint64(stage.GetOrder(geoobjectuse)))
	return
}

func (group *Group) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(group).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(group), uint64(stage.GetOrder(group)))
	return
}

func (groupuse *GroupUse) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(groupuse).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(groupuse), uint64(stage.GetOrder(groupuse)))
	return
}

func (library *Library) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(library).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(library), uint64(stage.GetOrder(library)))
	return
}

func (mapobject *MapObject) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(mapobject).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(mapobject), uint64(stage.GetOrder(mapobject)))
	return
}

func (mapobjectuse *MapObjectUse) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(mapobjectuse).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(mapobjectuse), uint64(stage.GetOrder(mapobjectuse)))
	return
}

func (parameter *Parameter) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(parameter).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(parameter), uint64(stage.GetOrder(parameter)))
	return
}

func (parametercategory *ParameterCategory) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(parametercategory).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(parametercategory), uint64(stage.GetOrder(parametercategory)))
	return
}

func (parametercategoryuse *ParameterCategoryUse) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(parametercategoryuse).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(parametercategoryuse), uint64(stage.GetOrder(parametercategoryuse)))
	return
}

func (parametershape *ParameterShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(parametershape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(parametershape), uint64(stage.GetOrder(parametershape)))
	return
}

func (parametersaggregate *ParametersAggregate) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(parametersaggregate).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(parametersaggregate), uint64(stage.GetOrder(parametersaggregate)))
	return
}

func (parametersaggregateshape *ParametersAggregateShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(parametersaggregateshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(parametersaggregateshape), uint64(stage.GetOrder(parametersaggregateshape)))
	return
}

func (position *Position) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(position).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(position), uint64(stage.GetOrder(position)))
	return
}

func (repository *Repository) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(repository).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(repository), uint64(stage.GetOrder(repository)))
	return
}

func (scenario *Scenario) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(scenario).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(scenario), uint64(stage.GetOrder(scenario)))
	return
}

func (user *User) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(user).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(user), uint64(stage.GetOrder(user)))
	return
}

func (useruse *UserUse) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(useruse).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(useruse), uint64(stage.GetOrder(useruse)))
	return
}

func (workspace *Workspace) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(workspace).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(workspace), uint64(stage.GetOrder(workspace)))
	return
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
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(actorstate.Name))
	return
}

func (actorstateshape *ActorStateShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", actorstateshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ActorStateShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(actorstateshape.Name))
	return
}

func (actorstatetransition *ActorStateTransition) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", actorstatetransition.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ActorStateTransition")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(actorstatetransition.Name))
	return
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", actorstatetransitionshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ActorStateTransitionShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(actorstatetransitionshape.Name))
	return
}

func (analysis *Analysis) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", analysis.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Analysis")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(analysis.Name))
	return
}

func (controlpointshape *ControlPointShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", controlpointshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ControlPointShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(controlpointshape.Name))
	return
}

func (diagram *Diagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Diagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(diagram.Name))
	return
}

func (document *Document) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", document.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Document")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(document.Name))
	return
}

func (documentuse *DocumentUse) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", documentuse.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DocumentUse")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(documentuse.Name))
	return
}

func (evolutiondirection *EvolutionDirection) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", evolutiondirection.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EvolutionDirection")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(evolutiondirection.Name))
	return
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", evolutiondirectionshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EvolutionDirectionShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(evolutiondirectionshape.Name))
	return
}

func (foo *Foo) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", foo.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Foo")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(foo.Name))
	return
}

func (geoobject *GeoObject) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", geoobject.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GeoObject")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(geoobject.Name))
	return
}

func (geoobjectuse *GeoObjectUse) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", geoobjectuse.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GeoObjectUse")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(geoobjectuse.Name))
	return
}

func (group *Group) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Group")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(group.Name))
	return
}

func (groupuse *GroupUse) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", groupuse.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GroupUse")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(groupuse.Name))
	return
}

func (library *Library) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Library")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(library.Name))
	return
}

func (mapobject *MapObject) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", mapobject.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MapObject")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(mapobject.Name))
	return
}

func (mapobjectuse *MapObjectUse) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", mapobjectuse.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MapObjectUse")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(mapobjectuse.Name))
	return
}

func (parameter *Parameter) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parameter.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Parameter")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(parameter.Name))
	return
}

func (parametercategory *ParameterCategory) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parametercategory.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParameterCategory")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(parametercategory.Name))
	return
}

func (parametercategoryuse *ParameterCategoryUse) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parametercategoryuse.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParameterCategoryUse")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(parametercategoryuse.Name))
	return
}

func (parametershape *ParameterShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parametershape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParameterShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(parametershape.Name))
	return
}

func (parametersaggregate *ParametersAggregate) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parametersaggregate.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParametersAggregate")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(parametersaggregate.Name))
	return
}

func (parametersaggregateshape *ParametersAggregateShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parametersaggregateshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParametersAggregateShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(parametersaggregateshape.Name))
	return
}

func (position *Position) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", position.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Position")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(position.Name))
	return
}

func (repository *Repository) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", repository.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Repository")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(repository.Name))
	return
}

func (scenario *Scenario) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", scenario.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Scenario")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(scenario.Name))
	return
}

func (user *User) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", user.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "User")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(user.Name))
	return
}

func (useruse *UserUse) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", useruse.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "UserUse")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(useruse.Name))
	return
}

func (workspace *Workspace) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", workspace.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Workspace")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(workspace.Name))
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

// end of template
