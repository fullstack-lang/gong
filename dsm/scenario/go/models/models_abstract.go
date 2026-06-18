package models

import (
	"log"
	"time"
)

// --- from actor_state.go ---
type ActorState struct {

	//gong:text
	Name string

	//gong:text gong:width 500 gong:height 500
	Description string

	// Probability Management
	IsWithProbaility bool
	Probability      ProbabilityEnum

	AbstractTypeFields
	LibraryAbstractFields
}

// --- from actor_state_transition.go ---
type ActorStateTransition struct {
	Name string

	StartState *ActorState
	EndState   *ActorState

	Justifications []*Parameter

	LibraryAbstractFields
	AbstractTypeFields
}

// --- from analysis.go ---
// Analysis is a first rank element
type Analysis struct {
	Name string

	//gong:text gong:width 500 gong:height 500
	Description string

	Scenarios               []*Scenario
	IsScenariosNodeExpanded bool

	GroupUse               []*GroupUse
	IsGroupUseNodeExpanded bool

	// GeoObjectUse provides link to geo objects
	GeoObjectUse               []*GeoObjectUse
	IsGeoObjectUseNodeExpanded bool

	// MapUse provides link to reference map for this analysis
	MapUse               []*MapObjectUse
	IsMapUseNodeExpanded bool

	LibraryAbstractFields
	AbstractTypeFields
}

// Sanitize remove orphaned actor state transitions
func Sanitize(stage *Stage) {
	for analysis := range *GetGongstructInstancesSet[Analysis](stage) {
		analysis.Sanitize()
	}
}

func (analysis *Analysis) Sanitize() {

	for _, scenario := range analysis.Scenarios {
		var newActorStateTransitions []*ActorStateTransition
		for _, actorStateTransition := range scenario.ActorStateTransitions {
			if actorStateTransition.EndState != nil &&
				actorStateTransition.StartState != nil {
				newActorStateTransitions = append(newActorStateTransitions, actorStateTransition)
			} else {
				log.Println(actorStateTransition.Name, "discarded")
			}
		}
		scenario.ActorStateTransitions = newActorStateTransitions
	}

}

// --- from document.go ---
type Document struct {
	Name string

	// GeoObjectUse provides link to geo objects
	GeoObjectUse []*GeoObjectUse
	LibraryAbstractFields
	AbstractTypeFields
}

type DocumentUse struct {
	Name     string
	Document *Document
}

// --- from evolution_direction.go ---
type EvolutionDirection struct {
	Name        string
	Description string

	AbstractTypeFields
	LibraryAbstractFields
}

// --- from geo_object.go ---
type GeoObject struct {
	Name string
	LibraryAbstractFields
	AbstractTypeFields
}

type GeoObjectUse struct {
	Name      string
	GeoObject *GeoObject
}

// --- from group.go ---
type Group struct {
	Name    string
	UserUse []*UserUse
	LibraryAbstractFields
	AbstractTypeFields
}

type GroupUse struct {
	Name  string
	Group *Group
}

// --- from library.go ---
type Library struct {
	Name string

	//gong:text width:300 height:300
	Description string

	LibraryAbstractFields
	AbstractTypeFields

	// There is one and only one root library per stage.
	IsRootLibrary bool

	objects []AbstractType

	Analyses               []*Analysis
	IsAnalysesNodeExpanded bool

	SubLibraries                    []*Library
	IsSubLibrariesNodeExpanded      bool
	SubLibrariesWhoseNodeIsExpanded []*Library

	NbPixPerCharacter float64 // stored at the root Library only

	//gong:width 600 gong:height 300
	LogoSVGFile string // the content of the logo file, used for the static site generation

	// temporary persistance of the library expand status.
	IsExpandedTmp bool
}

// --- from map_object.go ---
type MapObject struct {
	Name string
	LibraryAbstractFields
	AbstractTypeFields
}

type MapObjectUse struct {
	Name string
	Map  *MapObject
}

// --- from parameter.go ---
type Parameter struct {
	Name string

	// Description provides a longer account than the name of the parameter
	// This should allow the paramter to be understood
	// It also provides a reference to any input documents that
	// generated the parameter
	//gong:text gong:width 250 gong:height 500
	Description string

	IsResponse bool

	// Start and End are used when Paramter is used in a simulation
	Start, End time.Time
	Force      float64

	// GroupUse provides the need to know capability
	GroupUse []*GroupUse

	// DocumentUse enables link to a document
	DocumentUse []*DocumentUse

	// GeoObjectUse provides link to geo objects
	GeoObjectUse []*GeoObjectUse

	// Tag is used for labelling a given version of a parameter
	Tag string

	AbstractTypeFields
	LibraryAbstractFields
}

// --- from parameter_category.go ---
type ParameterCategory struct {
	Name         string
	ParameterUse []*ParameterShape
	LibraryAbstractFields
	AbstractTypeFields
}

type ParameterCategoryUse struct {
	Name              string
	ParameterCategory *ParameterCategory
}

// --- from parameters_aggregate.go ---
type ParametersAggregate struct {
	Name string

	// Tag is used for labelling a given version of a parameter
	Tag string

	// Description provides a longer account than the name of the parameter
	// This should allow the paramter to be understood
	// It also provides a reference to any input documents that
	// generated the parameter
	//gong:text gong:width 250 gong:height 500
	Description string

	// Parameters is the list of parameters this scenario
	// paramters aggregation
	Parameters []*Parameter

	AbstractTypeFields
	LibraryAbstractFields
}

// --- from position.go ---
type Position struct {
	Name string

	Date time.Time

	// between -1 and 1
	Ordinate float64
	LibraryAbstractFields
	AbstractTypeFields
}

// --- from scenario.go ---
type Scenario struct {
	//gong:text, magic code to have the form editor have a text area instead of an input
	Name string

	//gong:text, magic code to have the form editor have a text area instead of an input
	Description string

	Diagrams               []*Diagram
	IsDiagramsNodeExpanded bool

	ActorStates               []*ActorState
	IsActorStatesNodeExpanded bool

	ActorStateTransitions               []*ActorStateTransition
	IsActorStateTransitionsNodeExpanded bool

	EvolutionDirections               []*EvolutionDirection
	IsEvolutionDirectionsNodeExpanded bool

	Parameters               []*Parameter
	IsParametersNodeExpanded bool

	ParametersAggretates               []*ParametersAggregate
	IsParametersAggretatesNodeExpanded bool

	LibraryAbstractFields
	AbstractTypeFields
}

type I_RefreshDiagram interface {
	RefreshDiagram(diagram *Diagram)
}

// --- from workspace.go ---
// Workspace is the singloton
type Workspace struct {
	Name            string
	SelectedDiagram *Diagram

	// Default values for creation of shapes
	Default_EvolutionDirectionShape   *EvolutionDirectionShape
	Default_ParameterShape            *ParameterShape
	Default_ScenarioParameterShape    *ParametersAggregateShape
	Default_ActorStateShape           *ActorStateShape
	Default_ActorStateTransitionShape *ActorStateTransitionShape
	LibraryAbstractFields
	AbstractTypeFields
}

func GetWorkspace(stage *Stage) (workspace *Workspace) {

	workspaces := *GetGongstructInstancesSet[Workspace](stage)

	for _workspace := range workspaces {
		workspace = _workspace
	}

	return
}

// GetCurrentScenario parse all diagram and if one diagram
// matches the
func (workspace *Workspace) GetCurrentScenario(stage *Stage) (scenario *Scenario) {

	fieldName := GetAssociationName[Scenario]().Diagrams[0].Name
	map_Diagram_Scenario := GetSliceOfPointersReverseMap[Scenario, Diagram](fieldName, stage)

	scenario = map_Diagram_Scenario[workspace.SelectedDiagram][0]

	return
}

// --- from user.go ---
type User struct {
	Name string
	LibraryAbstractFields
	AbstractTypeFields
}

type UserUse struct {
	Name string
	User *User
}
