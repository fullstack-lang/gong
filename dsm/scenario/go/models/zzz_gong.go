// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
	"strings"
	"sync"
	"time"
)

// can be used for
//
//	days := __gong__abs(int(int(inferedInstance.ComputedDuration.Hours()) / 24))
func __gong__abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var (
	_ = __gong__abs
	_ = strings.Clone("")
)

const (
	GongProbeTreeSidebarSuffix           = ":sidebar of the probe"
	GongProbeNavigationTreeSidebarSuffix = ":sidebar of the probe, navigation"
	GongProbeTableSuffix                 = ":table of the probe"
	GongProbeNotificationTableSuffix     = ":notification table of the probe"
	GongProbeFormSuffix                  = ":form of the probe"
	GongProbeSplitSuffix                 = ":probe of the probe"
	GongProbeLoadSuffix                  = ":load of the probe"
)

type GongMarshallingMode string

const (
	// the whole stage is generated at each marshall. This is the default
	GongMarshallingNormal GongMarshallingMode = "GongMarshallingNormal"

	// only the last commit is append to the marshall file
	GongMarshallingAppendCommit GongMarshallingMode = "GongMarshallingAppendCommit"
)

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeNavigationTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeNavigationTreeSidebarSuffix
}

func (stage *Stage) GetProbeFormStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeFormSuffix
}

func (stage *Stage) GetProbeTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeTableSuffix
}

func (stage *Stage) GetProbeNotificationTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeNotificationTableSuffix
}

func (stage *Stage) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeSplitSuffix
}

func (stage *Stage) GetProbeLoadStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeLoadSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var (
	errUnkownEnum = errors.New("unkown enum")
	_             = errUnkownEnum
)

// needed to avoid when fmt package is not needed by generated code
var _ = fmt.Sprintf

// idem for math package when not need by generated code
var _ = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var (
	__member __void
	_        = __member
)

// MetaPackageImport represents a package import needed by a meta/diagram file
type MetaPackageImport struct {
	Alias string
	Path  string
}

// Stage enables storage of staged instances
type Stage struct {
	name string

	// isInDeltaMode is true when the stage is used to compute difference between
	// succesive commit
	isInDeltaMode bool

	// gongMarshallingMode set the marshalling mode
	gongMarshallingMode GongMarshallingMode
	// some stages have semantic rules that forbids them to be empty
	// like for git, the commit #0 (genesis commit) cannot be rolled back
	isWithGenesisCommit bool

	// insertion point for definition of arrays registering instances
	ActorStates                map[*ActorState]struct{}
	ActorStates_instance       map[*ActorState]*ActorState
	ActorStates_mapString      map[string]*ActorState
	ActorStateOrder            uint
	ActorState_stagedOrder     map[*ActorState]uint
	ActorState_orderStaged     map[uint]*ActorState
	ActorStates_reference      map[*ActorState]*ActorState
	ActorStates_referenceOrder map[*ActorState]uint

	// insertion point for slice of pointers maps
	OnAfterActorStateCreateCallback GongOnAfterCreateInterface[ActorState]
	OnAfterActorStateUpdateCallback GongOnAfterUpdateInterface[ActorState]
	OnAfterActorStateDeleteCallback GongOnAfterDeleteInterface[ActorState]

	ActorStateShapes                map[*ActorStateShape]struct{}
	ActorStateShapes_instance       map[*ActorStateShape]*ActorStateShape
	ActorStateShapes_mapString      map[string]*ActorStateShape
	ActorStateShapeOrder            uint
	ActorStateShape_stagedOrder     map[*ActorStateShape]uint
	ActorStateShape_orderStaged     map[uint]*ActorStateShape
	ActorStateShapes_reference      map[*ActorStateShape]*ActorStateShape
	ActorStateShapes_referenceOrder map[*ActorStateShape]uint

	// insertion point for slice of pointers maps
	OnAfterActorStateShapeCreateCallback GongOnAfterCreateInterface[ActorStateShape]
	OnAfterActorStateShapeUpdateCallback GongOnAfterUpdateInterface[ActorStateShape]
	OnAfterActorStateShapeDeleteCallback GongOnAfterDeleteInterface[ActorStateShape]

	ActorStateTransitions                map[*ActorStateTransition]struct{}
	ActorStateTransitions_instance       map[*ActorStateTransition]*ActorStateTransition
	ActorStateTransitions_mapString      map[string]*ActorStateTransition
	ActorStateTransitionOrder            uint
	ActorStateTransition_stagedOrder     map[*ActorStateTransition]uint
	ActorStateTransition_orderStaged     map[uint]*ActorStateTransition
	ActorStateTransitions_reference      map[*ActorStateTransition]*ActorStateTransition
	ActorStateTransitions_referenceOrder map[*ActorStateTransition]uint

	// insertion point for slice of pointers maps
	ActorStateTransition_Justifications_reverseMap map[*Parameter]*ActorStateTransition

	OnAfterActorStateTransitionCreateCallback GongOnAfterCreateInterface[ActorStateTransition]
	OnAfterActorStateTransitionUpdateCallback GongOnAfterUpdateInterface[ActorStateTransition]
	OnAfterActorStateTransitionDeleteCallback GongOnAfterDeleteInterface[ActorStateTransition]

	ActorStateTransitionShapes                map[*ActorStateTransitionShape]struct{}
	ActorStateTransitionShapes_instance       map[*ActorStateTransitionShape]*ActorStateTransitionShape
	ActorStateTransitionShapes_mapString      map[string]*ActorStateTransitionShape
	ActorStateTransitionShapeOrder            uint
	ActorStateTransitionShape_stagedOrder     map[*ActorStateTransitionShape]uint
	ActorStateTransitionShape_orderStaged     map[uint]*ActorStateTransitionShape
	ActorStateTransitionShapes_reference      map[*ActorStateTransitionShape]*ActorStateTransitionShape
	ActorStateTransitionShapes_referenceOrder map[*ActorStateTransitionShape]uint

	// insertion point for slice of pointers maps
	ActorStateTransitionShape_ControlPointShapes_reverseMap map[*ControlPointShape]*ActorStateTransitionShape

	OnAfterActorStateTransitionShapeCreateCallback GongOnAfterCreateInterface[ActorStateTransitionShape]
	OnAfterActorStateTransitionShapeUpdateCallback GongOnAfterUpdateInterface[ActorStateTransitionShape]
	OnAfterActorStateTransitionShapeDeleteCallback GongOnAfterDeleteInterface[ActorStateTransitionShape]

	Analysiss                map[*Analysis]struct{}
	Analysiss_instance       map[*Analysis]*Analysis
	Analysiss_mapString      map[string]*Analysis
	AnalysisOrder            uint
	Analysis_stagedOrder     map[*Analysis]uint
	Analysis_orderStaged     map[uint]*Analysis
	Analysiss_reference      map[*Analysis]*Analysis
	Analysiss_referenceOrder map[*Analysis]uint

	// insertion point for slice of pointers maps
	Analysis_Scenarios_reverseMap map[*Scenario]*Analysis

	Analysis_GroupUse_reverseMap map[*GroupUse]*Analysis

	Analysis_GeoObjectUse_reverseMap map[*GeoObjectUse]*Analysis

	Analysis_MapUse_reverseMap map[*MapObjectUse]*Analysis

	OnAfterAnalysisCreateCallback GongOnAfterCreateInterface[Analysis]
	OnAfterAnalysisUpdateCallback GongOnAfterUpdateInterface[Analysis]
	OnAfterAnalysisDeleteCallback GongOnAfterDeleteInterface[Analysis]

	ControlPointShapes                map[*ControlPointShape]struct{}
	ControlPointShapes_instance       map[*ControlPointShape]*ControlPointShape
	ControlPointShapes_mapString      map[string]*ControlPointShape
	ControlPointShapeOrder            uint
	ControlPointShape_stagedOrder     map[*ControlPointShape]uint
	ControlPointShape_orderStaged     map[uint]*ControlPointShape
	ControlPointShapes_reference      map[*ControlPointShape]*ControlPointShape
	ControlPointShapes_referenceOrder map[*ControlPointShape]uint

	// insertion point for slice of pointers maps
	OnAfterControlPointShapeCreateCallback GongOnAfterCreateInterface[ControlPointShape]
	OnAfterControlPointShapeUpdateCallback GongOnAfterUpdateInterface[ControlPointShape]
	OnAfterControlPointShapeDeleteCallback GongOnAfterDeleteInterface[ControlPointShape]

	Diagrams                map[*Diagram]struct{}
	Diagrams_instance       map[*Diagram]*Diagram
	Diagrams_mapString      map[string]*Diagram
	DiagramOrder            uint
	Diagram_stagedOrder     map[*Diagram]uint
	Diagram_orderStaged     map[uint]*Diagram
	Diagrams_reference      map[*Diagram]*Diagram
	Diagrams_referenceOrder map[*Diagram]uint

	// insertion point for slice of pointers maps
	Diagram_EvolutionDirectionShapes_reverseMap map[*EvolutionDirectionShape]*Diagram

	Diagram_EvolutionDirectionsWhoseNodeIsExpanded_reverseMap map[*EvolutionDirection]*Diagram

	Diagram_ActorStateShapes_reverseMap map[*ActorStateShape]*Diagram

	Diagram_ActorStatesWhoseNodeIsExpanded_reverseMap map[*ActorState]*Diagram

	Diagram_ParameterShapes_reverseMap map[*ParameterShape]*Diagram

	Diagram_ParametersWhoseNodeIsExpanded_reverseMap map[*Parameter]*Diagram

	Diagram_ScenarioParameterShapes_reverseMap map[*ParametersAggregateShape]*Diagram

	Diagram_ParametersAggregatesWhoseNodeIsExpanded_reverseMap map[*ParametersAggregate]*Diagram

	Diagram_ActorStateTransitionShapes_reverseMap map[*ActorStateTransitionShape]*Diagram

	Diagram_ActorStateTransitionsWhoseNodeIsExpanded_reverseMap map[*ActorStateTransition]*Diagram

	OnAfterDiagramCreateCallback GongOnAfterCreateInterface[Diagram]
	OnAfterDiagramUpdateCallback GongOnAfterUpdateInterface[Diagram]
	OnAfterDiagramDeleteCallback GongOnAfterDeleteInterface[Diagram]

	Documents                map[*Document]struct{}
	Documents_instance       map[*Document]*Document
	Documents_mapString      map[string]*Document
	DocumentOrder            uint
	Document_stagedOrder     map[*Document]uint
	Document_orderStaged     map[uint]*Document
	Documents_reference      map[*Document]*Document
	Documents_referenceOrder map[*Document]uint

	// insertion point for slice of pointers maps
	Document_GeoObjectUse_reverseMap map[*GeoObjectUse]*Document

	OnAfterDocumentCreateCallback GongOnAfterCreateInterface[Document]
	OnAfterDocumentUpdateCallback GongOnAfterUpdateInterface[Document]
	OnAfterDocumentDeleteCallback GongOnAfterDeleteInterface[Document]

	DocumentUses                map[*DocumentUse]struct{}
	DocumentUses_instance       map[*DocumentUse]*DocumentUse
	DocumentUses_mapString      map[string]*DocumentUse
	DocumentUseOrder            uint
	DocumentUse_stagedOrder     map[*DocumentUse]uint
	DocumentUse_orderStaged     map[uint]*DocumentUse
	DocumentUses_reference      map[*DocumentUse]*DocumentUse
	DocumentUses_referenceOrder map[*DocumentUse]uint

	// insertion point for slice of pointers maps
	OnAfterDocumentUseCreateCallback GongOnAfterCreateInterface[DocumentUse]
	OnAfterDocumentUseUpdateCallback GongOnAfterUpdateInterface[DocumentUse]
	OnAfterDocumentUseDeleteCallback GongOnAfterDeleteInterface[DocumentUse]

	EvolutionDirections                map[*EvolutionDirection]struct{}
	EvolutionDirections_instance       map[*EvolutionDirection]*EvolutionDirection
	EvolutionDirections_mapString      map[string]*EvolutionDirection
	EvolutionDirectionOrder            uint
	EvolutionDirection_stagedOrder     map[*EvolutionDirection]uint
	EvolutionDirection_orderStaged     map[uint]*EvolutionDirection
	EvolutionDirections_reference      map[*EvolutionDirection]*EvolutionDirection
	EvolutionDirections_referenceOrder map[*EvolutionDirection]uint

	// insertion point for slice of pointers maps
	OnAfterEvolutionDirectionCreateCallback GongOnAfterCreateInterface[EvolutionDirection]
	OnAfterEvolutionDirectionUpdateCallback GongOnAfterUpdateInterface[EvolutionDirection]
	OnAfterEvolutionDirectionDeleteCallback GongOnAfterDeleteInterface[EvolutionDirection]

	EvolutionDirectionShapes                map[*EvolutionDirectionShape]struct{}
	EvolutionDirectionShapes_instance       map[*EvolutionDirectionShape]*EvolutionDirectionShape
	EvolutionDirectionShapes_mapString      map[string]*EvolutionDirectionShape
	EvolutionDirectionShapeOrder            uint
	EvolutionDirectionShape_stagedOrder     map[*EvolutionDirectionShape]uint
	EvolutionDirectionShape_orderStaged     map[uint]*EvolutionDirectionShape
	EvolutionDirectionShapes_reference      map[*EvolutionDirectionShape]*EvolutionDirectionShape
	EvolutionDirectionShapes_referenceOrder map[*EvolutionDirectionShape]uint

	// insertion point for slice of pointers maps
	OnAfterEvolutionDirectionShapeCreateCallback GongOnAfterCreateInterface[EvolutionDirectionShape]
	OnAfterEvolutionDirectionShapeUpdateCallback GongOnAfterUpdateInterface[EvolutionDirectionShape]
	OnAfterEvolutionDirectionShapeDeleteCallback GongOnAfterDeleteInterface[EvolutionDirectionShape]

	Foos                map[*Foo]struct{}
	Foos_instance       map[*Foo]*Foo
	Foos_mapString      map[string]*Foo
	FooOrder            uint
	Foo_stagedOrder     map[*Foo]uint
	Foo_orderStaged     map[uint]*Foo
	Foos_reference      map[*Foo]*Foo
	Foos_referenceOrder map[*Foo]uint

	// insertion point for slice of pointers maps
	OnAfterFooCreateCallback GongOnAfterCreateInterface[Foo]
	OnAfterFooUpdateCallback GongOnAfterUpdateInterface[Foo]
	OnAfterFooDeleteCallback GongOnAfterDeleteInterface[Foo]

	GeoObjects                map[*GeoObject]struct{}
	GeoObjects_instance       map[*GeoObject]*GeoObject
	GeoObjects_mapString      map[string]*GeoObject
	GeoObjectOrder            uint
	GeoObject_stagedOrder     map[*GeoObject]uint
	GeoObject_orderStaged     map[uint]*GeoObject
	GeoObjects_reference      map[*GeoObject]*GeoObject
	GeoObjects_referenceOrder map[*GeoObject]uint

	// insertion point for slice of pointers maps
	OnAfterGeoObjectCreateCallback GongOnAfterCreateInterface[GeoObject]
	OnAfterGeoObjectUpdateCallback GongOnAfterUpdateInterface[GeoObject]
	OnAfterGeoObjectDeleteCallback GongOnAfterDeleteInterface[GeoObject]

	GeoObjectUses                map[*GeoObjectUse]struct{}
	GeoObjectUses_instance       map[*GeoObjectUse]*GeoObjectUse
	GeoObjectUses_mapString      map[string]*GeoObjectUse
	GeoObjectUseOrder            uint
	GeoObjectUse_stagedOrder     map[*GeoObjectUse]uint
	GeoObjectUse_orderStaged     map[uint]*GeoObjectUse
	GeoObjectUses_reference      map[*GeoObjectUse]*GeoObjectUse
	GeoObjectUses_referenceOrder map[*GeoObjectUse]uint

	// insertion point for slice of pointers maps
	OnAfterGeoObjectUseCreateCallback GongOnAfterCreateInterface[GeoObjectUse]
	OnAfterGeoObjectUseUpdateCallback GongOnAfterUpdateInterface[GeoObjectUse]
	OnAfterGeoObjectUseDeleteCallback GongOnAfterDeleteInterface[GeoObjectUse]

	Groups                map[*Group]struct{}
	Groups_instance       map[*Group]*Group
	Groups_mapString      map[string]*Group
	GroupOrder            uint
	Group_stagedOrder     map[*Group]uint
	Group_orderStaged     map[uint]*Group
	Groups_reference      map[*Group]*Group
	Groups_referenceOrder map[*Group]uint

	// insertion point for slice of pointers maps
	Group_UserUse_reverseMap map[*UserUse]*Group

	OnAfterGroupCreateCallback GongOnAfterCreateInterface[Group]
	OnAfterGroupUpdateCallback GongOnAfterUpdateInterface[Group]
	OnAfterGroupDeleteCallback GongOnAfterDeleteInterface[Group]

	GroupUses                map[*GroupUse]struct{}
	GroupUses_instance       map[*GroupUse]*GroupUse
	GroupUses_mapString      map[string]*GroupUse
	GroupUseOrder            uint
	GroupUse_stagedOrder     map[*GroupUse]uint
	GroupUse_orderStaged     map[uint]*GroupUse
	GroupUses_reference      map[*GroupUse]*GroupUse
	GroupUses_referenceOrder map[*GroupUse]uint

	// insertion point for slice of pointers maps
	OnAfterGroupUseCreateCallback GongOnAfterCreateInterface[GroupUse]
	OnAfterGroupUseUpdateCallback GongOnAfterUpdateInterface[GroupUse]
	OnAfterGroupUseDeleteCallback GongOnAfterDeleteInterface[GroupUse]

	Librarys                map[*Library]struct{}
	Librarys_instance       map[*Library]*Library
	Librarys_mapString      map[string]*Library
	LibraryOrder            uint
	Library_stagedOrder     map[*Library]uint
	Library_orderStaged     map[uint]*Library
	Librarys_reference      map[*Library]*Library
	Librarys_referenceOrder map[*Library]uint

	// insertion point for slice of pointers maps
	Library_Analyses_reverseMap map[*Analysis]*Library

	Library_SubLibraries_reverseMap map[*Library]*Library

	Library_SubLibrariesWhoseNodeIsExpanded_reverseMap map[*Library]*Library

	OnAfterLibraryCreateCallback GongOnAfterCreateInterface[Library]
	OnAfterLibraryUpdateCallback GongOnAfterUpdateInterface[Library]
	OnAfterLibraryDeleteCallback GongOnAfterDeleteInterface[Library]

	MapObjects                map[*MapObject]struct{}
	MapObjects_instance       map[*MapObject]*MapObject
	MapObjects_mapString      map[string]*MapObject
	MapObjectOrder            uint
	MapObject_stagedOrder     map[*MapObject]uint
	MapObject_orderStaged     map[uint]*MapObject
	MapObjects_reference      map[*MapObject]*MapObject
	MapObjects_referenceOrder map[*MapObject]uint

	// insertion point for slice of pointers maps
	OnAfterMapObjectCreateCallback GongOnAfterCreateInterface[MapObject]
	OnAfterMapObjectUpdateCallback GongOnAfterUpdateInterface[MapObject]
	OnAfterMapObjectDeleteCallback GongOnAfterDeleteInterface[MapObject]

	MapObjectUses                map[*MapObjectUse]struct{}
	MapObjectUses_instance       map[*MapObjectUse]*MapObjectUse
	MapObjectUses_mapString      map[string]*MapObjectUse
	MapObjectUseOrder            uint
	MapObjectUse_stagedOrder     map[*MapObjectUse]uint
	MapObjectUse_orderStaged     map[uint]*MapObjectUse
	MapObjectUses_reference      map[*MapObjectUse]*MapObjectUse
	MapObjectUses_referenceOrder map[*MapObjectUse]uint

	// insertion point for slice of pointers maps
	OnAfterMapObjectUseCreateCallback GongOnAfterCreateInterface[MapObjectUse]
	OnAfterMapObjectUseUpdateCallback GongOnAfterUpdateInterface[MapObjectUse]
	OnAfterMapObjectUseDeleteCallback GongOnAfterDeleteInterface[MapObjectUse]

	Parameters                map[*Parameter]struct{}
	Parameters_instance       map[*Parameter]*Parameter
	Parameters_mapString      map[string]*Parameter
	ParameterOrder            uint
	Parameter_stagedOrder     map[*Parameter]uint
	Parameter_orderStaged     map[uint]*Parameter
	Parameters_reference      map[*Parameter]*Parameter
	Parameters_referenceOrder map[*Parameter]uint

	// insertion point for slice of pointers maps
	Parameter_GroupUse_reverseMap map[*GroupUse]*Parameter

	Parameter_DocumentUse_reverseMap map[*DocumentUse]*Parameter

	Parameter_GeoObjectUse_reverseMap map[*GeoObjectUse]*Parameter

	OnAfterParameterCreateCallback GongOnAfterCreateInterface[Parameter]
	OnAfterParameterUpdateCallback GongOnAfterUpdateInterface[Parameter]
	OnAfterParameterDeleteCallback GongOnAfterDeleteInterface[Parameter]

	ParameterCategorys                map[*ParameterCategory]struct{}
	ParameterCategorys_instance       map[*ParameterCategory]*ParameterCategory
	ParameterCategorys_mapString      map[string]*ParameterCategory
	ParameterCategoryOrder            uint
	ParameterCategory_stagedOrder     map[*ParameterCategory]uint
	ParameterCategory_orderStaged     map[uint]*ParameterCategory
	ParameterCategorys_reference      map[*ParameterCategory]*ParameterCategory
	ParameterCategorys_referenceOrder map[*ParameterCategory]uint

	// insertion point for slice of pointers maps
	ParameterCategory_ParameterUse_reverseMap map[*ParameterShape]*ParameterCategory

	OnAfterParameterCategoryCreateCallback GongOnAfterCreateInterface[ParameterCategory]
	OnAfterParameterCategoryUpdateCallback GongOnAfterUpdateInterface[ParameterCategory]
	OnAfterParameterCategoryDeleteCallback GongOnAfterDeleteInterface[ParameterCategory]

	ParameterCategoryUses                map[*ParameterCategoryUse]struct{}
	ParameterCategoryUses_instance       map[*ParameterCategoryUse]*ParameterCategoryUse
	ParameterCategoryUses_mapString      map[string]*ParameterCategoryUse
	ParameterCategoryUseOrder            uint
	ParameterCategoryUse_stagedOrder     map[*ParameterCategoryUse]uint
	ParameterCategoryUse_orderStaged     map[uint]*ParameterCategoryUse
	ParameterCategoryUses_reference      map[*ParameterCategoryUse]*ParameterCategoryUse
	ParameterCategoryUses_referenceOrder map[*ParameterCategoryUse]uint

	// insertion point for slice of pointers maps
	OnAfterParameterCategoryUseCreateCallback GongOnAfterCreateInterface[ParameterCategoryUse]
	OnAfterParameterCategoryUseUpdateCallback GongOnAfterUpdateInterface[ParameterCategoryUse]
	OnAfterParameterCategoryUseDeleteCallback GongOnAfterDeleteInterface[ParameterCategoryUse]

	ParameterShapes                map[*ParameterShape]struct{}
	ParameterShapes_instance       map[*ParameterShape]*ParameterShape
	ParameterShapes_mapString      map[string]*ParameterShape
	ParameterShapeOrder            uint
	ParameterShape_stagedOrder     map[*ParameterShape]uint
	ParameterShape_orderStaged     map[uint]*ParameterShape
	ParameterShapes_reference      map[*ParameterShape]*ParameterShape
	ParameterShapes_referenceOrder map[*ParameterShape]uint

	// insertion point for slice of pointers maps
	OnAfterParameterShapeCreateCallback GongOnAfterCreateInterface[ParameterShape]
	OnAfterParameterShapeUpdateCallback GongOnAfterUpdateInterface[ParameterShape]
	OnAfterParameterShapeDeleteCallback GongOnAfterDeleteInterface[ParameterShape]

	ParametersAggregates                map[*ParametersAggregate]struct{}
	ParametersAggregates_instance       map[*ParametersAggregate]*ParametersAggregate
	ParametersAggregates_mapString      map[string]*ParametersAggregate
	ParametersAggregateOrder            uint
	ParametersAggregate_stagedOrder     map[*ParametersAggregate]uint
	ParametersAggregate_orderStaged     map[uint]*ParametersAggregate
	ParametersAggregates_reference      map[*ParametersAggregate]*ParametersAggregate
	ParametersAggregates_referenceOrder map[*ParametersAggregate]uint

	// insertion point for slice of pointers maps
	ParametersAggregate_Parameters_reverseMap map[*Parameter]*ParametersAggregate

	OnAfterParametersAggregateCreateCallback GongOnAfterCreateInterface[ParametersAggregate]
	OnAfterParametersAggregateUpdateCallback GongOnAfterUpdateInterface[ParametersAggregate]
	OnAfterParametersAggregateDeleteCallback GongOnAfterDeleteInterface[ParametersAggregate]

	ParametersAggregateShapes                map[*ParametersAggregateShape]struct{}
	ParametersAggregateShapes_instance       map[*ParametersAggregateShape]*ParametersAggregateShape
	ParametersAggregateShapes_mapString      map[string]*ParametersAggregateShape
	ParametersAggregateShapeOrder            uint
	ParametersAggregateShape_stagedOrder     map[*ParametersAggregateShape]uint
	ParametersAggregateShape_orderStaged     map[uint]*ParametersAggregateShape
	ParametersAggregateShapes_reference      map[*ParametersAggregateShape]*ParametersAggregateShape
	ParametersAggregateShapes_referenceOrder map[*ParametersAggregateShape]uint

	// insertion point for slice of pointers maps
	OnAfterParametersAggregateShapeCreateCallback GongOnAfterCreateInterface[ParametersAggregateShape]
	OnAfterParametersAggregateShapeUpdateCallback GongOnAfterUpdateInterface[ParametersAggregateShape]
	OnAfterParametersAggregateShapeDeleteCallback GongOnAfterDeleteInterface[ParametersAggregateShape]

	Positions                map[*Position]struct{}
	Positions_instance       map[*Position]*Position
	Positions_mapString      map[string]*Position
	PositionOrder            uint
	Position_stagedOrder     map[*Position]uint
	Position_orderStaged     map[uint]*Position
	Positions_reference      map[*Position]*Position
	Positions_referenceOrder map[*Position]uint

	// insertion point for slice of pointers maps
	OnAfterPositionCreateCallback GongOnAfterCreateInterface[Position]
	OnAfterPositionUpdateCallback GongOnAfterUpdateInterface[Position]
	OnAfterPositionDeleteCallback GongOnAfterDeleteInterface[Position]

	Repositorys                map[*Repository]struct{}
	Repositorys_instance       map[*Repository]*Repository
	Repositorys_mapString      map[string]*Repository
	RepositoryOrder            uint
	Repository_stagedOrder     map[*Repository]uint
	Repository_orderStaged     map[uint]*Repository
	Repositorys_reference      map[*Repository]*Repository
	Repositorys_referenceOrder map[*Repository]uint

	// insertion point for slice of pointers maps
	Repository_ParameterUse_reverseMap map[*ParameterShape]*Repository

	Repository_GroupUse_reverseMap map[*GroupUse]*Repository

	OnAfterRepositoryCreateCallback GongOnAfterCreateInterface[Repository]
	OnAfterRepositoryUpdateCallback GongOnAfterUpdateInterface[Repository]
	OnAfterRepositoryDeleteCallback GongOnAfterDeleteInterface[Repository]

	Scenarios                map[*Scenario]struct{}
	Scenarios_instance       map[*Scenario]*Scenario
	Scenarios_mapString      map[string]*Scenario
	ScenarioOrder            uint
	Scenario_stagedOrder     map[*Scenario]uint
	Scenario_orderStaged     map[uint]*Scenario
	Scenarios_reference      map[*Scenario]*Scenario
	Scenarios_referenceOrder map[*Scenario]uint

	// insertion point for slice of pointers maps
	Scenario_Diagrams_reverseMap map[*Diagram]*Scenario

	Scenario_ActorStates_reverseMap map[*ActorState]*Scenario

	Scenario_ActorStateTransitions_reverseMap map[*ActorStateTransition]*Scenario

	Scenario_EvolutionDirections_reverseMap map[*EvolutionDirection]*Scenario

	Scenario_Parameters_reverseMap map[*Parameter]*Scenario

	Scenario_ParametersAggretates_reverseMap map[*ParametersAggregate]*Scenario

	OnAfterScenarioCreateCallback GongOnAfterCreateInterface[Scenario]
	OnAfterScenarioUpdateCallback GongOnAfterUpdateInterface[Scenario]
	OnAfterScenarioDeleteCallback GongOnAfterDeleteInterface[Scenario]

	Users                map[*User]struct{}
	Users_instance       map[*User]*User
	Users_mapString      map[string]*User
	UserOrder            uint
	User_stagedOrder     map[*User]uint
	User_orderStaged     map[uint]*User
	Users_reference      map[*User]*User
	Users_referenceOrder map[*User]uint

	// insertion point for slice of pointers maps
	OnAfterUserCreateCallback GongOnAfterCreateInterface[User]
	OnAfterUserUpdateCallback GongOnAfterUpdateInterface[User]
	OnAfterUserDeleteCallback GongOnAfterDeleteInterface[User]

	UserUses                map[*UserUse]struct{}
	UserUses_instance       map[*UserUse]*UserUse
	UserUses_mapString      map[string]*UserUse
	UserUseOrder            uint
	UserUse_stagedOrder     map[*UserUse]uint
	UserUse_orderStaged     map[uint]*UserUse
	UserUses_reference      map[*UserUse]*UserUse
	UserUses_referenceOrder map[*UserUse]uint

	// insertion point for slice of pointers maps
	OnAfterUserUseCreateCallback GongOnAfterCreateInterface[UserUse]
	OnAfterUserUseUpdateCallback GongOnAfterUpdateInterface[UserUse]
	OnAfterUserUseDeleteCallback GongOnAfterDeleteInterface[UserUse]

	Workspaces                map[*Workspace]struct{}
	Workspaces_instance       map[*Workspace]*Workspace
	Workspaces_mapString      map[string]*Workspace
	WorkspaceOrder            uint
	Workspace_stagedOrder     map[*Workspace]uint
	Workspace_orderStaged     map[uint]*Workspace
	Workspaces_reference      map[*Workspace]*Workspace
	Workspaces_referenceOrder map[*Workspace]uint

	// insertion point for slice of pointers maps
	OnAfterWorkspaceCreateCallback GongOnAfterCreateInterface[Workspace]
	OnAfterWorkspaceUpdateCallback GongOnAfterUpdateInterface[Workspace]
	OnAfterWorkspaceDeleteCallback GongOnAfterDeleteInterface[Workspace]

	BackRepo GongBackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          GongOnInitCommitInterface
	OnInitCommitFromFrontCallback GongOnInitCommitInterface
	OnInitCommitFromBackCallback  GongOnInitCommitInterface

	// Private slices to hold the registered hooks
	beforeCommitHooks []func(stage *Stage)
	afterCommitHooks  []func(stage *Stage)

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string
	MetaPackageImports     []*MetaPackageImport

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here

	// store the stage order of each instance in order to
	// preserve this order when serializing them
	// insertion point for order fields declaration
	// end of insertion point

	// GongUnmarshallers is the registry of all model unmarshallers
	GongUnmarshallers map[string]GongModelUnmarshaller

	// probeIF is the interface to the probe that allows log
	// commit event to the probe
	probeIF GongProbeIF

	forwardCommits  []string
	backwardCommits []string

	// when navigating the commit history
	// navigationMode is set to Navigating
	navigationMode gongStageNavigationMode
	commitsBehind  int // the number of commits the stage is behind the front of the history

	isApplyingBackwardCommit bool
	isApplyingForwardCommit  bool
	isSquashing              bool

	modified bool

	lock sync.RWMutex
}

type GongStage = Stage

func (s *Stage) SetGongMarshallingMode(mode GongMarshallingMode) {
	s.gongMarshallingMode = mode
}

func (s *Stage) GetGongMarshallingMode() GongMarshallingMode {
	return s.gongMarshallingMode
}

func (s *Stage) SetIsWithGenesisCommit(isWithGenesisCommit bool) {
	s.isWithGenesisCommit = isWithGenesisCommit
}

func (s *Stage) GetIsWithGenesisCommit() bool {
	return s.isWithGenesisCommit
}

// RegisterBeforeCommit adds a hook that runs before the commit happens
func (s *Stage) RegisterBeforeCommit(hook func(stage *Stage)) {
	s.beforeCommitHooks = append(s.beforeCommitHooks, hook)
}

// RegisterAfterCommit adds a hook that runs after the commit succeeds
func (s *Stage) RegisterAfterCommit(hook func(stage *Stage)) {
	s.afterCommitHooks = append(s.afterCommitHooks, hook)
}

type gongStageNavigationMode string

const (
	GongNavigationModeNormal gongStageNavigationMode = "Normal"
	// when the mode is navigating, each commit backward and forward
	// it is possible to go apply the nbCommitsBackward forward commits
	GongNavigationModeNavigating gongStageNavigationMode = "Navigating"
)

// ApplyBackwardCommit applies the commit before the current one
func (stage *Stage) ApplyBackwardCommit() error {
	if len(stage.backwardCommits) == 0 {
		return errors.New("no backward commit to apply")
	}

	if stage.navigationMode == GongNavigationModeNormal && stage.commitsBehind != 0 {
		return errors.New("in navigation mode normal, cannot have commitsBehind != 0")
	}

	if stage.navigationMode == GongNavigationModeNormal {
		stage.navigationMode = GongNavigationModeNavigating
	}

	if stage.isWithGenesisCommit && stage.commitsBehind >= len(stage.backwardCommits)-1 {
		return errors.New("cannot rollback genesis commit")
	}

	if stage.commitsBehind >= len(stage.backwardCommits) {
		return errors.New("no more backward commit to apply")
	}

	commitToApply := stage.backwardCommits[len(stage.backwardCommits)-1-stage.commitsBehind]

	// umarshall the backward commit to the stage

	// the parsing of the commit will call the UX update
	// therefore, it is important to stage.commitsBehind before because it is used in the
	// UX
	stage.commitsBehind++
	stage.isApplyingBackwardCommit = true
	err := stage.ParseAstString(commitToApply, true)
	stage.isApplyingBackwardCommit = false
	if err != nil {
		log.Println("error during ApplyBackwardCommit: ", err)
		return err
	}

	stage.ComputeReferenceAndOrders()

	return nil
}

func (stage *Stage) GetForwardCommits() []string {
	return stage.forwardCommits
}

func (stage *Stage) GetBackwardCommits() []string {
	return stage.backwardCommits
}

func (stage *Stage) ApplyForwardCommit() error {
	if stage.navigationMode == GongNavigationModeNormal && stage.commitsBehind != 0 {
		return errors.New("in navigation mode normal, cannot have commitsBehind != 0")
	}

	if stage.commitsBehind == 0 {
		return errors.New("no more forward commit to apply")
	}

	if stage.navigationMode == GongNavigationModeNormal {
		stage.navigationMode = GongNavigationModeNavigating
	}

	commitToApply := stage.forwardCommits[len(stage.forwardCommits)-1-stage.commitsBehind+1]

	// the parsing of the commit will call the UX update
	// therefore, it is important to stage.commitsBehind before because it is used in the
	// UX
	stage.commitsBehind--
	stage.isApplyingForwardCommit = true
	err := stage.ParseAstString(commitToApply, true)
	stage.isApplyingForwardCommit = false
	if err != nil {
		log.Println("error during ApplyForwardCommit: ", err)
		return err
	}
	stage.ComputeReferenceAndOrders()

	return nil
}

func (stage *Stage) GetCommitsBehind() int {
	return stage.commitsBehind
}

func (stage *Stage) Lock() {
	stage.lock.Lock()
}

func (stage *Stage) Unlock() {
	stage.lock.Unlock()
}

func (stage *Stage) RLock() {
	stage.lock.RLock()
}

func (stage *Stage) RUnlock() {
	stage.lock.RUnlock()
}

// ResetHard removes the more recent
// commitsBehind forward/backward Commits from the
// stage
func (stage *Stage) ResetHard() {
	newCommitsLen := len(stage.forwardCommits) - stage.GetCommitsBehind()

	stage.forwardCommits = stage.forwardCommits[:newCommitsLen]
	stage.backwardCommits = stage.backwardCommits[:newCommitsLen]
	stage.commitsBehind = 0
	stage.navigationMode = GongNavigationModeNormal

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	for _, hook := range stage.beforeCommitHooks {
		hook(stage)
	}

	// 2. Run all After Commit hooks
	for _, hook := range stage.afterCommitHooks {
		hook(stage)
	}
}

// Squash removes all commits and marshals the stage as a single commit
func (stage *Stage) Squash() {
	stage.forwardCommits = stage.forwardCommits[:0]
	stage.backwardCommits = stage.backwardCommits[:0]
	stage.commitsBehind = 0
	stage.navigationMode = GongNavigationModeNormal

	stage.modified = true
	stage.isSquashing = true

	// insertion point for clear references
	__gong__clearReferences(&stage.ActorStates_reference, &stage.ActorStates_instance, &stage.ActorStates_referenceOrder)

	__gong__clearReferences(&stage.ActorStateShapes_reference, &stage.ActorStateShapes_instance, &stage.ActorStateShapes_referenceOrder)

	__gong__clearReferences(&stage.ActorStateTransitions_reference, &stage.ActorStateTransitions_instance, &stage.ActorStateTransitions_referenceOrder)

	__gong__clearReferences(&stage.ActorStateTransitionShapes_reference, &stage.ActorStateTransitionShapes_instance, &stage.ActorStateTransitionShapes_referenceOrder)

	__gong__clearReferences(&stage.Analysiss_reference, &stage.Analysiss_instance, &stage.Analysiss_referenceOrder)

	__gong__clearReferences(&stage.ControlPointShapes_reference, &stage.ControlPointShapes_instance, &stage.ControlPointShapes_referenceOrder)

	__gong__clearReferences(&stage.Diagrams_reference, &stage.Diagrams_instance, &stage.Diagrams_referenceOrder)

	__gong__clearReferences(&stage.Documents_reference, &stage.Documents_instance, &stage.Documents_referenceOrder)

	__gong__clearReferences(&stage.DocumentUses_reference, &stage.DocumentUses_instance, &stage.DocumentUses_referenceOrder)

	__gong__clearReferences(&stage.EvolutionDirections_reference, &stage.EvolutionDirections_instance, &stage.EvolutionDirections_referenceOrder)

	__gong__clearReferences(&stage.EvolutionDirectionShapes_reference, &stage.EvolutionDirectionShapes_instance, &stage.EvolutionDirectionShapes_referenceOrder)

	__gong__clearReferences(&stage.Foos_reference, &stage.Foos_instance, &stage.Foos_referenceOrder)

	__gong__clearReferences(&stage.GeoObjects_reference, &stage.GeoObjects_instance, &stage.GeoObjects_referenceOrder)

	__gong__clearReferences(&stage.GeoObjectUses_reference, &stage.GeoObjectUses_instance, &stage.GeoObjectUses_referenceOrder)

	__gong__clearReferences(&stage.Groups_reference, &stage.Groups_instance, &stage.Groups_referenceOrder)

	__gong__clearReferences(&stage.GroupUses_reference, &stage.GroupUses_instance, &stage.GroupUses_referenceOrder)

	__gong__clearReferences(&stage.Librarys_reference, &stage.Librarys_instance, &stage.Librarys_referenceOrder)

	__gong__clearReferences(&stage.MapObjects_reference, &stage.MapObjects_instance, &stage.MapObjects_referenceOrder)

	__gong__clearReferences(&stage.MapObjectUses_reference, &stage.MapObjectUses_instance, &stage.MapObjectUses_referenceOrder)

	__gong__clearReferences(&stage.Parameters_reference, &stage.Parameters_instance, &stage.Parameters_referenceOrder)

	__gong__clearReferences(&stage.ParameterCategorys_reference, &stage.ParameterCategorys_instance, &stage.ParameterCategorys_referenceOrder)

	__gong__clearReferences(&stage.ParameterCategoryUses_reference, &stage.ParameterCategoryUses_instance, &stage.ParameterCategoryUses_referenceOrder)

	__gong__clearReferences(&stage.ParameterShapes_reference, &stage.ParameterShapes_instance, &stage.ParameterShapes_referenceOrder)

	__gong__clearReferences(&stage.ParametersAggregates_reference, &stage.ParametersAggregates_instance, &stage.ParametersAggregates_referenceOrder)

	__gong__clearReferences(&stage.ParametersAggregateShapes_reference, &stage.ParametersAggregateShapes_instance, &stage.ParametersAggregateShapes_referenceOrder)

	__gong__clearReferences(&stage.Positions_reference, &stage.Positions_instance, &stage.Positions_referenceOrder)

	__gong__clearReferences(&stage.Repositorys_reference, &stage.Repositorys_instance, &stage.Repositorys_referenceOrder)

	__gong__clearReferences(&stage.Scenarios_reference, &stage.Scenarios_instance, &stage.Scenarios_referenceOrder)

	__gong__clearReferences(&stage.Users_reference, &stage.Users_instance, &stage.Users_referenceOrder)

	__gong__clearReferences(&stage.UserUses_reference, &stage.UserUses_instance, &stage.UserUses_referenceOrder)

	__gong__clearReferences(&stage.Workspaces_reference, &stage.Workspaces_instance, &stage.Workspaces_referenceOrder)

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	for _, hook := range stage.beforeCommitHooks {
		hook(stage)
	}

	// 2. Run all After Commit hooks
	for _, hook := range stage.afterCommitHooks {
		hook(stage)
	}

	stage.isSquashing = false
}

// recomputeOrders recomputes the next order for each struct
// this is necessary because the order might have been incremented
// during the commits that have been discarded
// insertion point for max order recomputation
func (stage *Stage) recomputeOrders() {
	// insertion point for max order recomputation
	stage.ActorStateOrder = __gong__recomputeOrder(stage.ActorState_stagedOrder)

	stage.ActorStateShapeOrder = __gong__recomputeOrder(stage.ActorStateShape_stagedOrder)

	stage.ActorStateTransitionOrder = __gong__recomputeOrder(stage.ActorStateTransition_stagedOrder)

	stage.ActorStateTransitionShapeOrder = __gong__recomputeOrder(stage.ActorStateTransitionShape_stagedOrder)

	stage.AnalysisOrder = __gong__recomputeOrder(stage.Analysis_stagedOrder)

	stage.ControlPointShapeOrder = __gong__recomputeOrder(stage.ControlPointShape_stagedOrder)

	stage.DiagramOrder = __gong__recomputeOrder(stage.Diagram_stagedOrder)

	stage.DocumentOrder = __gong__recomputeOrder(stage.Document_stagedOrder)

	stage.DocumentUseOrder = __gong__recomputeOrder(stage.DocumentUse_stagedOrder)

	stage.EvolutionDirectionOrder = __gong__recomputeOrder(stage.EvolutionDirection_stagedOrder)

	stage.EvolutionDirectionShapeOrder = __gong__recomputeOrder(stage.EvolutionDirectionShape_stagedOrder)

	stage.FooOrder = __gong__recomputeOrder(stage.Foo_stagedOrder)

	stage.GeoObjectOrder = __gong__recomputeOrder(stage.GeoObject_stagedOrder)

	stage.GeoObjectUseOrder = __gong__recomputeOrder(stage.GeoObjectUse_stagedOrder)

	stage.GroupOrder = __gong__recomputeOrder(stage.Group_stagedOrder)

	stage.GroupUseOrder = __gong__recomputeOrder(stage.GroupUse_stagedOrder)

	stage.LibraryOrder = __gong__recomputeOrder(stage.Library_stagedOrder)

	stage.MapObjectOrder = __gong__recomputeOrder(stage.MapObject_stagedOrder)

	stage.MapObjectUseOrder = __gong__recomputeOrder(stage.MapObjectUse_stagedOrder)

	stage.ParameterOrder = __gong__recomputeOrder(stage.Parameter_stagedOrder)

	stage.ParameterCategoryOrder = __gong__recomputeOrder(stage.ParameterCategory_stagedOrder)

	stage.ParameterCategoryUseOrder = __gong__recomputeOrder(stage.ParameterCategoryUse_stagedOrder)

	stage.ParameterShapeOrder = __gong__recomputeOrder(stage.ParameterShape_stagedOrder)

	stage.ParametersAggregateOrder = __gong__recomputeOrder(stage.ParametersAggregate_stagedOrder)

	stage.ParametersAggregateShapeOrder = __gong__recomputeOrder(stage.ParametersAggregateShape_stagedOrder)

	stage.PositionOrder = __gong__recomputeOrder(stage.Position_stagedOrder)

	stage.RepositoryOrder = __gong__recomputeOrder(stage.Repository_stagedOrder)

	stage.ScenarioOrder = __gong__recomputeOrder(stage.Scenario_stagedOrder)

	stage.UserOrder = __gong__recomputeOrder(stage.User_stagedOrder)

	stage.UserUseOrder = __gong__recomputeOrder(stage.UserUse_stagedOrder)

	stage.WorkspaceOrder = __gong__recomputeOrder(stage.Workspace_stagedOrder)

	// end of insertion point for max order recomputation
}

func (stage *Stage) SetDeltaMode(inDeltaMode bool) {
	stage.isInDeltaMode = inDeltaMode
}

func (stage *Stage) IsInDeltaMode() bool {
	return stage.isInDeltaMode
}

func (stage *Stage) SetProbeIF(probeIF GongProbeIF) {
	stage.probeIF = probeIF
}

func (stage *Stage) GetProbeIF() GongProbeIF {
	if stage.probeIF == nil {
		return nil
	}

	return stage.probeIF
}

// GetInstancesByOrder is the Stage method returning a slice of generic pointers to gongstructs
// ordered by their order in the stage.
func (stage *Stage) GetInstancesByOrder[T GongstructPtr]() (res []T) {
	var t T
	switch any(t).(type) {
	// insertion point for case
	case *ActorState:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ActorStates, stage.ActorState_stagedOrder))
	case *ActorStateShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ActorStateShapes, stage.ActorStateShape_stagedOrder))
	case *ActorStateTransition:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ActorStateTransitions, stage.ActorStateTransition_stagedOrder))
	case *ActorStateTransitionShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ActorStateTransitionShapes, stage.ActorStateTransitionShape_stagedOrder))
	case *Analysis:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Analysiss, stage.Analysis_stagedOrder))
	case *ControlPointShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ControlPointShapes, stage.ControlPointShape_stagedOrder))
	case *Diagram:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Diagrams, stage.Diagram_stagedOrder))
	case *Document:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Documents, stage.Document_stagedOrder))
	case *DocumentUse:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.DocumentUses, stage.DocumentUse_stagedOrder))
	case *EvolutionDirection:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.EvolutionDirections, stage.EvolutionDirection_stagedOrder))
	case *EvolutionDirectionShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.EvolutionDirectionShapes, stage.EvolutionDirectionShape_stagedOrder))
	case *Foo:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Foos, stage.Foo_stagedOrder))
	case *GeoObject:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.GeoObjects, stage.GeoObject_stagedOrder))
	case *GeoObjectUse:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.GeoObjectUses, stage.GeoObjectUse_stagedOrder))
	case *Group:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Groups, stage.Group_stagedOrder))
	case *GroupUse:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.GroupUses, stage.GroupUse_stagedOrder))
	case *Library:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Librarys, stage.Library_stagedOrder))
	case *MapObject:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.MapObjects, stage.MapObject_stagedOrder))
	case *MapObjectUse:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.MapObjectUses, stage.MapObjectUse_stagedOrder))
	case *Parameter:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Parameters, stage.Parameter_stagedOrder))
	case *ParameterCategory:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ParameterCategorys, stage.ParameterCategory_stagedOrder))
	case *ParameterCategoryUse:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ParameterCategoryUses, stage.ParameterCategoryUse_stagedOrder))
	case *ParameterShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ParameterShapes, stage.ParameterShape_stagedOrder))
	case *ParametersAggregate:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ParametersAggregates, stage.ParametersAggregate_stagedOrder))
	case *ParametersAggregateShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ParametersAggregateShapes, stage.ParametersAggregateShape_stagedOrder))
	case *Position:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Positions, stage.Position_stagedOrder))
	case *Repository:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Repositorys, stage.Repository_stagedOrder))
	case *Scenario:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Scenarios, stage.Scenario_stagedOrder))
	case *User:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Users, stage.User_stagedOrder))
	case *UserUse:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.UserUses, stage.UserUse_stagedOrder))
	case *Workspace:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Workspaces, stage.Workspace_stagedOrder))

	}
	return
}

func __gong__getStructInstancesByOrder[T GongstructPtr](set map[T]struct{}, order map[T]uint) (res []T) {
	orderedSet := []T{}
	for instance := range set {
		orderedSet = append(orderedSet, instance)
	}
	sort.Slice(orderedSet[:], func(i, j int) bool {
		instancei := orderedSet[i]
		instancej := orderedSet[j]
		i_order, oki := order[instancei]
		j_order, okj := order[instancej]
		if !oki || !okj {
			log.Fatalf("getStructInstancesByOrder: pointer not found")
		}
		return i_order < j_order
	})

	res = append(res, orderedSet...)

	return
}

func __gong__castSlice[T any, S any](s []S) []T {
	res := make([]T, len(s))
	for i, v := range s {
		res[i] = any(v).(T)
	}
	return res
}

func __gong__stage[T comparable](
	instances map[T]struct{},
	stagedOrder map[T]uint,
	orderStaged map[uint]T,
	order *uint,
	mapString map[string]T,
	instance T,
	name string,
) {
	if _, ok := instances[instance]; !ok {
		instances[instance] = struct{}{}
		stagedOrder[instance] = *order
		orderStaged[*order] = instance
		*order++
	}
	mapString[name] = instance
}

func __gong__stagePreserveOrder[T comparable](
	instances map[T]struct{},
	stagedOrder map[T]uint,
	orderStaged map[uint]T,
	currentOrder *uint,
	mapString map[string]T,
	instance T,
	order uint,
	name string,
) {
	if _, ok := instances[instance]; !ok {
		instances[instance] = struct{}{}
		if order > *currentOrder {
			*currentOrder = order
		}
		stagedOrder[instance] = order
		orderStaged[order] = instance
		*currentOrder++
	}
	mapString[name] = instance
}

func __gong__unstage[T comparable](
	instances map[T]struct{},
	mapString map[string]T,
	instance T,
	name string,
) {
	delete(instances, instance)
	delete(mapString, name)
}

func __gong__recomputeOrder[T comparable](stagedOrder map[T]uint) uint {
	var maxOrder uint
	var found bool
	for _, order := range stagedOrder {
		if !found || order > maxOrder {
			maxOrder = order
			found = true
		}
	}
	if found {
		return maxOrder + 1
	}
	return 0
}

func __gong__rebuildMapString[T interface {
	comparable
	GetName() string
}](staged map[T]struct{}, mapString *map[string]T) {
	*mapString = make(map[string]T, len(staged))
	for instance := range staged {
		(*mapString)[instance.GetName()] = instance
	}
}

func __gong__clearReferences[T comparable](ref *map[T]T, inst *map[T]T, refOrder *map[T]uint) {
	*ref = make(map[T]T)
	*inst = make(map[T]T)
	*refOrder = make(map[T]uint)
}

func __gong__resetStageType[T comparable](staged *map[T]struct{}, mapString *map[string]T, stagedOrder *map[T]uint, order *uint) {
	*staged = make(map[T]struct{})
	*mapString = make(map[string]T)
	*stagedOrder = make(map[T]uint)
	*order = 0
}

func (stage *Stage) GetType() string {
	return "github.com/fullstack-lang/gong/dsm/scenario/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type GongOnInitCommitInterface interface {
	BeforeCommit(stage *Stage)
}

type OnInitCommitInterface = GongOnInitCommitInterface

// GongOnAfterCreateInterface callback when an instance is updated from the front
type GongOnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *Stage,
		instance *Type)
}

type OnAfterCreateInterface[Type Gongstruct] = GongOnAfterCreateInterface[Type]

// GongOnAfterUpdateInterface callback when an instance is updated from the front
type GongOnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *Stage, old, new *Type)
}

type OnAfterUpdateInterface[Type Gongstruct] = GongOnAfterUpdateInterface[Type]

// GongOnAfterDeleteInterface callback when an instance is updated from the front
type GongOnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *Stage,
		staged, front *Type)
}

type OnAfterDeleteInterface[Type Gongstruct] = GongOnAfterDeleteInterface[Type]

type GongBackRepoInterface interface {
	Commit(stage *Stage)
	Checkout(stage *Stage)
	Backup(stage *Stage, dirPath string)
	Restore(stage *Stage, dirPath string)
	BackupXL(stage *Stage, dirPath string)
	RestoreXL(stage *Stage, dirPath string)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

type BackRepoInterface = GongBackRepoInterface

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		ActorStates:           make(map[*ActorState]struct{}),
		ActorStates_mapString: make(map[string]*ActorState),

		ActorStateShapes:           make(map[*ActorStateShape]struct{}),
		ActorStateShapes_mapString: make(map[string]*ActorStateShape),

		ActorStateTransitions:           make(map[*ActorStateTransition]struct{}),
		ActorStateTransitions_mapString: make(map[string]*ActorStateTransition),

		ActorStateTransitionShapes:           make(map[*ActorStateTransitionShape]struct{}),
		ActorStateTransitionShapes_mapString: make(map[string]*ActorStateTransitionShape),

		Analysiss:           make(map[*Analysis]struct{}),
		Analysiss_mapString: make(map[string]*Analysis),

		ControlPointShapes:           make(map[*ControlPointShape]struct{}),
		ControlPointShapes_mapString: make(map[string]*ControlPointShape),

		Diagrams:           make(map[*Diagram]struct{}),
		Diagrams_mapString: make(map[string]*Diagram),

		Documents:           make(map[*Document]struct{}),
		Documents_mapString: make(map[string]*Document),

		DocumentUses:           make(map[*DocumentUse]struct{}),
		DocumentUses_mapString: make(map[string]*DocumentUse),

		EvolutionDirections:           make(map[*EvolutionDirection]struct{}),
		EvolutionDirections_mapString: make(map[string]*EvolutionDirection),

		EvolutionDirectionShapes:           make(map[*EvolutionDirectionShape]struct{}),
		EvolutionDirectionShapes_mapString: make(map[string]*EvolutionDirectionShape),

		Foos:           make(map[*Foo]struct{}),
		Foos_mapString: make(map[string]*Foo),

		GeoObjects:           make(map[*GeoObject]struct{}),
		GeoObjects_mapString: make(map[string]*GeoObject),

		GeoObjectUses:           make(map[*GeoObjectUse]struct{}),
		GeoObjectUses_mapString: make(map[string]*GeoObjectUse),

		Groups:           make(map[*Group]struct{}),
		Groups_mapString: make(map[string]*Group),

		GroupUses:           make(map[*GroupUse]struct{}),
		GroupUses_mapString: make(map[string]*GroupUse),

		Librarys:           make(map[*Library]struct{}),
		Librarys_mapString: make(map[string]*Library),

		MapObjects:           make(map[*MapObject]struct{}),
		MapObjects_mapString: make(map[string]*MapObject),

		MapObjectUses:           make(map[*MapObjectUse]struct{}),
		MapObjectUses_mapString: make(map[string]*MapObjectUse),

		Parameters:           make(map[*Parameter]struct{}),
		Parameters_mapString: make(map[string]*Parameter),

		ParameterCategorys:           make(map[*ParameterCategory]struct{}),
		ParameterCategorys_mapString: make(map[string]*ParameterCategory),

		ParameterCategoryUses:           make(map[*ParameterCategoryUse]struct{}),
		ParameterCategoryUses_mapString: make(map[string]*ParameterCategoryUse),

		ParameterShapes:           make(map[*ParameterShape]struct{}),
		ParameterShapes_mapString: make(map[string]*ParameterShape),

		ParametersAggregates:           make(map[*ParametersAggregate]struct{}),
		ParametersAggregates_mapString: make(map[string]*ParametersAggregate),

		ParametersAggregateShapes:           make(map[*ParametersAggregateShape]struct{}),
		ParametersAggregateShapes_mapString: make(map[string]*ParametersAggregateShape),

		Positions:           make(map[*Position]struct{}),
		Positions_mapString: make(map[string]*Position),

		Repositorys:           make(map[*Repository]struct{}),
		Repositorys_mapString: make(map[string]*Repository),

		Scenarios:           make(map[*Scenario]struct{}),
		Scenarios_mapString: make(map[string]*Scenario),

		Users:           make(map[*User]struct{}),
		Users_mapString: make(map[string]*User),

		UserUses:           make(map[*UserUse]struct{}),
		UserUses_mapString: make(map[string]*UserUse),

		Workspaces:           make(map[*Workspace]struct{}),
		Workspaces_mapString: make(map[string]*Workspace),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		ActorState_stagedOrder: make(map[*ActorState]uint),
		ActorState_orderStaged: make(map[uint]*ActorState),
		ActorStates_reference:  make(map[*ActorState]*ActorState),

		ActorStateShape_stagedOrder: make(map[*ActorStateShape]uint),
		ActorStateShape_orderStaged: make(map[uint]*ActorStateShape),
		ActorStateShapes_reference:  make(map[*ActorStateShape]*ActorStateShape),

		ActorStateTransition_stagedOrder: make(map[*ActorStateTransition]uint),
		ActorStateTransition_orderStaged: make(map[uint]*ActorStateTransition),
		ActorStateTransitions_reference:  make(map[*ActorStateTransition]*ActorStateTransition),

		ActorStateTransitionShape_stagedOrder: make(map[*ActorStateTransitionShape]uint),
		ActorStateTransitionShape_orderStaged: make(map[uint]*ActorStateTransitionShape),
		ActorStateTransitionShapes_reference:  make(map[*ActorStateTransitionShape]*ActorStateTransitionShape),

		Analysis_stagedOrder: make(map[*Analysis]uint),
		Analysis_orderStaged: make(map[uint]*Analysis),
		Analysiss_reference:  make(map[*Analysis]*Analysis),

		ControlPointShape_stagedOrder: make(map[*ControlPointShape]uint),
		ControlPointShape_orderStaged: make(map[uint]*ControlPointShape),
		ControlPointShapes_reference:  make(map[*ControlPointShape]*ControlPointShape),

		Diagram_stagedOrder: make(map[*Diagram]uint),
		Diagram_orderStaged: make(map[uint]*Diagram),
		Diagrams_reference:  make(map[*Diagram]*Diagram),

		Document_stagedOrder: make(map[*Document]uint),
		Document_orderStaged: make(map[uint]*Document),
		Documents_reference:  make(map[*Document]*Document),

		DocumentUse_stagedOrder: make(map[*DocumentUse]uint),
		DocumentUse_orderStaged: make(map[uint]*DocumentUse),
		DocumentUses_reference:  make(map[*DocumentUse]*DocumentUse),

		EvolutionDirection_stagedOrder: make(map[*EvolutionDirection]uint),
		EvolutionDirection_orderStaged: make(map[uint]*EvolutionDirection),
		EvolutionDirections_reference:  make(map[*EvolutionDirection]*EvolutionDirection),

		EvolutionDirectionShape_stagedOrder: make(map[*EvolutionDirectionShape]uint),
		EvolutionDirectionShape_orderStaged: make(map[uint]*EvolutionDirectionShape),
		EvolutionDirectionShapes_reference:  make(map[*EvolutionDirectionShape]*EvolutionDirectionShape),

		Foo_stagedOrder: make(map[*Foo]uint),
		Foo_orderStaged: make(map[uint]*Foo),
		Foos_reference:  make(map[*Foo]*Foo),

		GeoObject_stagedOrder: make(map[*GeoObject]uint),
		GeoObject_orderStaged: make(map[uint]*GeoObject),
		GeoObjects_reference:  make(map[*GeoObject]*GeoObject),

		GeoObjectUse_stagedOrder: make(map[*GeoObjectUse]uint),
		GeoObjectUse_orderStaged: make(map[uint]*GeoObjectUse),
		GeoObjectUses_reference:  make(map[*GeoObjectUse]*GeoObjectUse),

		Group_stagedOrder: make(map[*Group]uint),
		Group_orderStaged: make(map[uint]*Group),
		Groups_reference:  make(map[*Group]*Group),

		GroupUse_stagedOrder: make(map[*GroupUse]uint),
		GroupUse_orderStaged: make(map[uint]*GroupUse),
		GroupUses_reference:  make(map[*GroupUse]*GroupUse),

		Library_stagedOrder: make(map[*Library]uint),
		Library_orderStaged: make(map[uint]*Library),
		Librarys_reference:  make(map[*Library]*Library),

		MapObject_stagedOrder: make(map[*MapObject]uint),
		MapObject_orderStaged: make(map[uint]*MapObject),
		MapObjects_reference:  make(map[*MapObject]*MapObject),

		MapObjectUse_stagedOrder: make(map[*MapObjectUse]uint),
		MapObjectUse_orderStaged: make(map[uint]*MapObjectUse),
		MapObjectUses_reference:  make(map[*MapObjectUse]*MapObjectUse),

		Parameter_stagedOrder: make(map[*Parameter]uint),
		Parameter_orderStaged: make(map[uint]*Parameter),
		Parameters_reference:  make(map[*Parameter]*Parameter),

		ParameterCategory_stagedOrder: make(map[*ParameterCategory]uint),
		ParameterCategory_orderStaged: make(map[uint]*ParameterCategory),
		ParameterCategorys_reference:  make(map[*ParameterCategory]*ParameterCategory),

		ParameterCategoryUse_stagedOrder: make(map[*ParameterCategoryUse]uint),
		ParameterCategoryUse_orderStaged: make(map[uint]*ParameterCategoryUse),
		ParameterCategoryUses_reference:  make(map[*ParameterCategoryUse]*ParameterCategoryUse),

		ParameterShape_stagedOrder: make(map[*ParameterShape]uint),
		ParameterShape_orderStaged: make(map[uint]*ParameterShape),
		ParameterShapes_reference:  make(map[*ParameterShape]*ParameterShape),

		ParametersAggregate_stagedOrder: make(map[*ParametersAggregate]uint),
		ParametersAggregate_orderStaged: make(map[uint]*ParametersAggregate),
		ParametersAggregates_reference:  make(map[*ParametersAggregate]*ParametersAggregate),

		ParametersAggregateShape_stagedOrder: make(map[*ParametersAggregateShape]uint),
		ParametersAggregateShape_orderStaged: make(map[uint]*ParametersAggregateShape),
		ParametersAggregateShapes_reference:  make(map[*ParametersAggregateShape]*ParametersAggregateShape),

		Position_stagedOrder: make(map[*Position]uint),
		Position_orderStaged: make(map[uint]*Position),
		Positions_reference:  make(map[*Position]*Position),

		Repository_stagedOrder: make(map[*Repository]uint),
		Repository_orderStaged: make(map[uint]*Repository),
		Repositorys_reference:  make(map[*Repository]*Repository),

		Scenario_stagedOrder: make(map[*Scenario]uint),
		Scenario_orderStaged: make(map[uint]*Scenario),
		Scenarios_reference:  make(map[*Scenario]*Scenario),

		User_stagedOrder: make(map[*User]uint),
		User_orderStaged: make(map[uint]*User),
		Users_reference:  make(map[*User]*User),

		UserUse_stagedOrder: make(map[*UserUse]uint),
		UserUse_orderStaged: make(map[uint]*UserUse),
		UserUses_reference:  make(map[*UserUse]*UserUse),

		Workspace_stagedOrder: make(map[*Workspace]uint),
		Workspace_orderStaged: make(map[uint]*Workspace),
		Workspaces_reference:  make(map[*Workspace]*Workspace),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"ActorState": &ActorStateUnmarshaller{},

			"ActorStateShape": &ActorStateShapeUnmarshaller{},

			"ActorStateTransition": &ActorStateTransitionUnmarshaller{},

			"ActorStateTransitionShape": &ActorStateTransitionShapeUnmarshaller{},

			"Analysis": &AnalysisUnmarshaller{},

			"ControlPointShape": &ControlPointShapeUnmarshaller{},

			"Diagram": &DiagramUnmarshaller{},

			"Document": &DocumentUnmarshaller{},

			"DocumentUse": &DocumentUseUnmarshaller{},

			"EvolutionDirection": &EvolutionDirectionUnmarshaller{},

			"EvolutionDirectionShape": &EvolutionDirectionShapeUnmarshaller{},

			"Foo": &FooUnmarshaller{},

			"GeoObject": &GeoObjectUnmarshaller{},

			"GeoObjectUse": &GeoObjectUseUnmarshaller{},

			"Group": &GroupUnmarshaller{},

			"GroupUse": &GroupUseUnmarshaller{},

			"Library": &LibraryUnmarshaller{},

			"MapObject": &MapObjectUnmarshaller{},

			"MapObjectUse": &MapObjectUseUnmarshaller{},

			"Parameter": &ParameterUnmarshaller{},

			"ParameterCategory": &ParameterCategoryUnmarshaller{},

			"ParameterCategoryUse": &ParameterCategoryUseUnmarshaller{},

			"ParameterShape": &ParameterShapeUnmarshaller{},

			"ParametersAggregate": &ParametersAggregateUnmarshaller{},

			"ParametersAggregateShape": &ParametersAggregateShapeUnmarshaller{},

			"Position": &PositionUnmarshaller{},

			"Repository": &RepositoryUnmarshaller{},

			"Scenario": &ScenarioUnmarshaller{},

			"User": &UserUnmarshaller{},

			"UserUse": &UserUseUnmarshaller{},

			"Workspace": &WorkspaceUnmarshaller{},

			// end of insertion point
		},

		navigationMode: GongNavigationModeNormal,
	}

	return
}

// GetOrder is the Stage method returning the order of a gongstruct instance.
func (stage *Stage) GetOrder(instance GongstructIF) uint {
	if instance != nil {
		return instance.GongGetOrder(stage)
	}
	return 0
}

// GetInstanceFromOrder is the Stage method returning a gongstruct instance from its order.
func (stage *Stage) GetInstanceFromOrder[Type GongstructPtr](order uint) (res Type) {
	var t Type
	switch any(t).(type) {
	// insertion point for order map initialisations
	case *ActorState:
		return any(stage.ActorState_orderStaged[order]).(Type)
	case *ActorStateShape:
		return any(stage.ActorStateShape_orderStaged[order]).(Type)
	case *ActorStateTransition:
		return any(stage.ActorStateTransition_orderStaged[order]).(Type)
	case *ActorStateTransitionShape:
		return any(stage.ActorStateTransitionShape_orderStaged[order]).(Type)
	case *Analysis:
		return any(stage.Analysis_orderStaged[order]).(Type)
	case *ControlPointShape:
		return any(stage.ControlPointShape_orderStaged[order]).(Type)
	case *Diagram:
		return any(stage.Diagram_orderStaged[order]).(Type)
	case *Document:
		return any(stage.Document_orderStaged[order]).(Type)
	case *DocumentUse:
		return any(stage.DocumentUse_orderStaged[order]).(Type)
	case *EvolutionDirection:
		return any(stage.EvolutionDirection_orderStaged[order]).(Type)
	case *EvolutionDirectionShape:
		return any(stage.EvolutionDirectionShape_orderStaged[order]).(Type)
	case *Foo:
		return any(stage.Foo_orderStaged[order]).(Type)
	case *GeoObject:
		return any(stage.GeoObject_orderStaged[order]).(Type)
	case *GeoObjectUse:
		return any(stage.GeoObjectUse_orderStaged[order]).(Type)
	case *Group:
		return any(stage.Group_orderStaged[order]).(Type)
	case *GroupUse:
		return any(stage.GroupUse_orderStaged[order]).(Type)
	case *Library:
		return any(stage.Library_orderStaged[order]).(Type)
	case *MapObject:
		return any(stage.MapObject_orderStaged[order]).(Type)
	case *MapObjectUse:
		return any(stage.MapObjectUse_orderStaged[order]).(Type)
	case *Parameter:
		return any(stage.Parameter_orderStaged[order]).(Type)
	case *ParameterCategory:
		return any(stage.ParameterCategory_orderStaged[order]).(Type)
	case *ParameterCategoryUse:
		return any(stage.ParameterCategoryUse_orderStaged[order]).(Type)
	case *ParameterShape:
		return any(stage.ParameterShape_orderStaged[order]).(Type)
	case *ParametersAggregate:
		return any(stage.ParametersAggregate_orderStaged[order]).(Type)
	case *ParametersAggregateShape:
		return any(stage.ParametersAggregateShape_orderStaged[order]).(Type)
	case *Position:
		return any(stage.Position_orderStaged[order]).(Type)
	case *Repository:
		return any(stage.Repository_orderStaged[order]).(Type)
	case *Scenario:
		return any(stage.Scenario_orderStaged[order]).(Type)
	case *User:
		return any(stage.User_orderStaged[order]).(Type)
	case *UserUse:
		return any(stage.UserUse_orderStaged[order]).(Type)
	case *Workspace:
		return any(stage.Workspace_orderStaged[order]).(Type)
	default:
		return // should not happen
	}
}

func (stage *Stage) GetName() string {
	return stage.name
}

func (stage *Stage) CommitWithSuspendedCallbacks() {
	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	tmp2 := stage.beforeCommitHooks
	stage.beforeCommitHooks = nil
	tmp3 := stage.afterCommitHooks
	stage.afterCommitHooks = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
	stage.beforeCommitHooks = tmp2
	stage.afterCommitHooks = tmp3
}

func (stage *Stage) Commit() {
	stage.ComputeReverseMaps()

	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	for _, hook := range stage.beforeCommitHooks {
		hook(stage)
	}

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}
	stage.ComputeInstancesNb()

	// if a commit is applied when in navigation mode
	// this will reset the commits behind and swith the
	// naviagation
	if stage.isInDeltaMode && stage.navigationMode == GongNavigationModeNavigating && stage.GetCommitsBehind() > 0 {
		stage.ResetHard()
	}

	if stage.IsInDeltaMode() {
		stage.ComputeForwardAndBackwardCommits()
		stage.ComputeReferenceAndOrders()
		if stage.probeIF != nil {
			stage.probeIF.RefreshNavigationTree()
		}
	}

	// 2. Run all After Commit hooks
	for _, hook := range stage.afterCommitHooks {
		hook(stage)
	}
}

func (stage *Stage) ComputeInstancesNb() {
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["ActorState"] = len(stage.ActorStates)
	stage.Map_GongStructName_InstancesNb["ActorStateShape"] = len(stage.ActorStateShapes)
	stage.Map_GongStructName_InstancesNb["ActorStateTransition"] = len(stage.ActorStateTransitions)
	stage.Map_GongStructName_InstancesNb["ActorStateTransitionShape"] = len(stage.ActorStateTransitionShapes)
	stage.Map_GongStructName_InstancesNb["Analysis"] = len(stage.Analysiss)
	stage.Map_GongStructName_InstancesNb["ControlPointShape"] = len(stage.ControlPointShapes)
	stage.Map_GongStructName_InstancesNb["Diagram"] = len(stage.Diagrams)
	stage.Map_GongStructName_InstancesNb["Document"] = len(stage.Documents)
	stage.Map_GongStructName_InstancesNb["DocumentUse"] = len(stage.DocumentUses)
	stage.Map_GongStructName_InstancesNb["EvolutionDirection"] = len(stage.EvolutionDirections)
	stage.Map_GongStructName_InstancesNb["EvolutionDirectionShape"] = len(stage.EvolutionDirectionShapes)
	stage.Map_GongStructName_InstancesNb["Foo"] = len(stage.Foos)
	stage.Map_GongStructName_InstancesNb["GeoObject"] = len(stage.GeoObjects)
	stage.Map_GongStructName_InstancesNb["GeoObjectUse"] = len(stage.GeoObjectUses)
	stage.Map_GongStructName_InstancesNb["Group"] = len(stage.Groups)
	stage.Map_GongStructName_InstancesNb["GroupUse"] = len(stage.GroupUses)
	stage.Map_GongStructName_InstancesNb["Library"] = len(stage.Librarys)
	stage.Map_GongStructName_InstancesNb["MapObject"] = len(stage.MapObjects)
	stage.Map_GongStructName_InstancesNb["MapObjectUse"] = len(stage.MapObjectUses)
	stage.Map_GongStructName_InstancesNb["Parameter"] = len(stage.Parameters)
	stage.Map_GongStructName_InstancesNb["ParameterCategory"] = len(stage.ParameterCategorys)
	stage.Map_GongStructName_InstancesNb["ParameterCategoryUse"] = len(stage.ParameterCategoryUses)
	stage.Map_GongStructName_InstancesNb["ParameterShape"] = len(stage.ParameterShapes)
	stage.Map_GongStructName_InstancesNb["ParametersAggregate"] = len(stage.ParametersAggregates)
	stage.Map_GongStructName_InstancesNb["ParametersAggregateShape"] = len(stage.ParametersAggregateShapes)
	stage.Map_GongStructName_InstancesNb["Position"] = len(stage.Positions)
	stage.Map_GongStructName_InstancesNb["Repository"] = len(stage.Repositorys)
	stage.Map_GongStructName_InstancesNb["Scenario"] = len(stage.Scenarios)
	stage.Map_GongStructName_InstancesNb["User"] = len(stage.Users)
	stage.Map_GongStructName_InstancesNb["UserUse"] = len(stage.UserUses)
	stage.Map_GongStructName_InstancesNb["Workspace"] = len(stage.Workspaces)
}

func (stage *Stage) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	stage.ComputeInstancesNb()
}

// backup generates backup files in the dirPath
func (stage *Stage) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *Stage) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *Stage) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *Stage) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts actorstate to the model stage
func (actorstate *ActorState) Stage(stage *Stage) *ActorState {
	__gong__stage(stage.ActorStates, stage.ActorState_stagedOrder, stage.ActorState_orderStaged, &stage.ActorStateOrder, stage.ActorStates_mapString, actorstate, actorstate.Name)
	return actorstate
}

// StagePreserveOrder puts actorstate to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ActorStateOrder
// - update stage.ActorStateOrder accordingly
func (actorstate *ActorState) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ActorStates, stage.ActorState_stagedOrder, stage.ActorState_orderStaged, &stage.ActorStateOrder, stage.ActorStates_mapString, actorstate, order, actorstate.Name)
}

// Unstage removes actorstate off the model stage
func (actorstate *ActorState) Unstage(stage *Stage) *ActorState {
	__gong__unstage(stage.ActorStates, stage.ActorStates_mapString, actorstate, actorstate.Name)
	return actorstate
}

// UnstageVoid removes actorstate off the model stage
func (actorstate *ActorState) UnstageVoid(stage *Stage) {
	actorstate.Unstage(stage)
}

func (actorstate *ActorState) StageVoid(stage *Stage) {
	actorstate.Stage(stage)
}

// for satisfaction of GongStruct interface
func (actorstate *ActorState) GetName() (res string) {
	return actorstate.Name
}

// for satisfaction of GongStruct interface
func (actorstate *ActorState) SetName(name string) {
	actorstate.Name = name
}

// Stage puts actorstateshape to the model stage
func (actorstateshape *ActorStateShape) Stage(stage *Stage) *ActorStateShape {
	__gong__stage(stage.ActorStateShapes, stage.ActorStateShape_stagedOrder, stage.ActorStateShape_orderStaged, &stage.ActorStateShapeOrder, stage.ActorStateShapes_mapString, actorstateshape, actorstateshape.Name)
	return actorstateshape
}

// StagePreserveOrder puts actorstateshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ActorStateShapeOrder
// - update stage.ActorStateShapeOrder accordingly
func (actorstateshape *ActorStateShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ActorStateShapes, stage.ActorStateShape_stagedOrder, stage.ActorStateShape_orderStaged, &stage.ActorStateShapeOrder, stage.ActorStateShapes_mapString, actorstateshape, order, actorstateshape.Name)
}

// Unstage removes actorstateshape off the model stage
func (actorstateshape *ActorStateShape) Unstage(stage *Stage) *ActorStateShape {
	__gong__unstage(stage.ActorStateShapes, stage.ActorStateShapes_mapString, actorstateshape, actorstateshape.Name)
	return actorstateshape
}

// UnstageVoid removes actorstateshape off the model stage
func (actorstateshape *ActorStateShape) UnstageVoid(stage *Stage) {
	actorstateshape.Unstage(stage)
}

func (actorstateshape *ActorStateShape) StageVoid(stage *Stage) {
	actorstateshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (actorstateshape *ActorStateShape) GetName() (res string) {
	return actorstateshape.Name
}

// for satisfaction of GongStruct interface
func (actorstateshape *ActorStateShape) SetName(name string) {
	actorstateshape.Name = name
}

// Stage puts actorstatetransition to the model stage
func (actorstatetransition *ActorStateTransition) Stage(stage *Stage) *ActorStateTransition {
	__gong__stage(stage.ActorStateTransitions, stage.ActorStateTransition_stagedOrder, stage.ActorStateTransition_orderStaged, &stage.ActorStateTransitionOrder, stage.ActorStateTransitions_mapString, actorstatetransition, actorstatetransition.Name)
	return actorstatetransition
}

// StagePreserveOrder puts actorstatetransition to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ActorStateTransitionOrder
// - update stage.ActorStateTransitionOrder accordingly
func (actorstatetransition *ActorStateTransition) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ActorStateTransitions, stage.ActorStateTransition_stagedOrder, stage.ActorStateTransition_orderStaged, &stage.ActorStateTransitionOrder, stage.ActorStateTransitions_mapString, actorstatetransition, order, actorstatetransition.Name)
}

// Unstage removes actorstatetransition off the model stage
func (actorstatetransition *ActorStateTransition) Unstage(stage *Stage) *ActorStateTransition {
	__gong__unstage(stage.ActorStateTransitions, stage.ActorStateTransitions_mapString, actorstatetransition, actorstatetransition.Name)
	return actorstatetransition
}

// UnstageVoid removes actorstatetransition off the model stage
func (actorstatetransition *ActorStateTransition) UnstageVoid(stage *Stage) {
	actorstatetransition.Unstage(stage)
}

func (actorstatetransition *ActorStateTransition) StageVoid(stage *Stage) {
	actorstatetransition.Stage(stage)
}

// for satisfaction of GongStruct interface
func (actorstatetransition *ActorStateTransition) GetName() (res string) {
	return actorstatetransition.Name
}

// for satisfaction of GongStruct interface
func (actorstatetransition *ActorStateTransition) SetName(name string) {
	actorstatetransition.Name = name
}

// Stage puts actorstatetransitionshape to the model stage
func (actorstatetransitionshape *ActorStateTransitionShape) Stage(stage *Stage) *ActorStateTransitionShape {
	__gong__stage(stage.ActorStateTransitionShapes, stage.ActorStateTransitionShape_stagedOrder, stage.ActorStateTransitionShape_orderStaged, &stage.ActorStateTransitionShapeOrder, stage.ActorStateTransitionShapes_mapString, actorstatetransitionshape, actorstatetransitionshape.Name)
	return actorstatetransitionshape
}

// StagePreserveOrder puts actorstatetransitionshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ActorStateTransitionShapeOrder
// - update stage.ActorStateTransitionShapeOrder accordingly
func (actorstatetransitionshape *ActorStateTransitionShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ActorStateTransitionShapes, stage.ActorStateTransitionShape_stagedOrder, stage.ActorStateTransitionShape_orderStaged, &stage.ActorStateTransitionShapeOrder, stage.ActorStateTransitionShapes_mapString, actorstatetransitionshape, order, actorstatetransitionshape.Name)
}

// Unstage removes actorstatetransitionshape off the model stage
func (actorstatetransitionshape *ActorStateTransitionShape) Unstage(stage *Stage) *ActorStateTransitionShape {
	__gong__unstage(stage.ActorStateTransitionShapes, stage.ActorStateTransitionShapes_mapString, actorstatetransitionshape, actorstatetransitionshape.Name)
	return actorstatetransitionshape
}

// UnstageVoid removes actorstatetransitionshape off the model stage
func (actorstatetransitionshape *ActorStateTransitionShape) UnstageVoid(stage *Stage) {
	actorstatetransitionshape.Unstage(stage)
}

func (actorstatetransitionshape *ActorStateTransitionShape) StageVoid(stage *Stage) {
	actorstatetransitionshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (actorstatetransitionshape *ActorStateTransitionShape) GetName() (res string) {
	return actorstatetransitionshape.Name
}

// for satisfaction of GongStruct interface
func (actorstatetransitionshape *ActorStateTransitionShape) SetName(name string) {
	actorstatetransitionshape.Name = name
}

// Stage puts analysis to the model stage
func (analysis *Analysis) Stage(stage *Stage) *Analysis {
	__gong__stage(stage.Analysiss, stage.Analysis_stagedOrder, stage.Analysis_orderStaged, &stage.AnalysisOrder, stage.Analysiss_mapString, analysis, analysis.Name)
	return analysis
}

// StagePreserveOrder puts analysis to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AnalysisOrder
// - update stage.AnalysisOrder accordingly
func (analysis *Analysis) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Analysiss, stage.Analysis_stagedOrder, stage.Analysis_orderStaged, &stage.AnalysisOrder, stage.Analysiss_mapString, analysis, order, analysis.Name)
}

// Unstage removes analysis off the model stage
func (analysis *Analysis) Unstage(stage *Stage) *Analysis {
	__gong__unstage(stage.Analysiss, stage.Analysiss_mapString, analysis, analysis.Name)
	return analysis
}

// UnstageVoid removes analysis off the model stage
func (analysis *Analysis) UnstageVoid(stage *Stage) {
	analysis.Unstage(stage)
}

func (analysis *Analysis) StageVoid(stage *Stage) {
	analysis.Stage(stage)
}

// for satisfaction of GongStruct interface
func (analysis *Analysis) GetName() (res string) {
	return analysis.Name
}

// for satisfaction of GongStruct interface
func (analysis *Analysis) SetName(name string) {
	analysis.Name = name
}

// Stage puts controlpointshape to the model stage
func (controlpointshape *ControlPointShape) Stage(stage *Stage) *ControlPointShape {
	__gong__stage(stage.ControlPointShapes, stage.ControlPointShape_stagedOrder, stage.ControlPointShape_orderStaged, &stage.ControlPointShapeOrder, stage.ControlPointShapes_mapString, controlpointshape, controlpointshape.Name)
	return controlpointshape
}

// StagePreserveOrder puts controlpointshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ControlPointShapeOrder
// - update stage.ControlPointShapeOrder accordingly
func (controlpointshape *ControlPointShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ControlPointShapes, stage.ControlPointShape_stagedOrder, stage.ControlPointShape_orderStaged, &stage.ControlPointShapeOrder, stage.ControlPointShapes_mapString, controlpointshape, order, controlpointshape.Name)
}

// Unstage removes controlpointshape off the model stage
func (controlpointshape *ControlPointShape) Unstage(stage *Stage) *ControlPointShape {
	__gong__unstage(stage.ControlPointShapes, stage.ControlPointShapes_mapString, controlpointshape, controlpointshape.Name)
	return controlpointshape
}

// UnstageVoid removes controlpointshape off the model stage
func (controlpointshape *ControlPointShape) UnstageVoid(stage *Stage) {
	controlpointshape.Unstage(stage)
}

func (controlpointshape *ControlPointShape) StageVoid(stage *Stage) {
	controlpointshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (controlpointshape *ControlPointShape) GetName() (res string) {
	return controlpointshape.Name
}

// for satisfaction of GongStruct interface
func (controlpointshape *ControlPointShape) SetName(name string) {
	controlpointshape.Name = name
}

// Stage puts diagram to the model stage
func (diagram *Diagram) Stage(stage *Stage) *Diagram {
	__gong__stage(stage.Diagrams, stage.Diagram_stagedOrder, stage.Diagram_orderStaged, &stage.DiagramOrder, stage.Diagrams_mapString, diagram, diagram.Name)
	return diagram
}

// StagePreserveOrder puts diagram to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DiagramOrder
// - update stage.DiagramOrder accordingly
func (diagram *Diagram) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Diagrams, stage.Diagram_stagedOrder, stage.Diagram_orderStaged, &stage.DiagramOrder, stage.Diagrams_mapString, diagram, order, diagram.Name)
}

// Unstage removes diagram off the model stage
func (diagram *Diagram) Unstage(stage *Stage) *Diagram {
	__gong__unstage(stage.Diagrams, stage.Diagrams_mapString, diagram, diagram.Name)
	return diagram
}

// UnstageVoid removes diagram off the model stage
func (diagram *Diagram) UnstageVoid(stage *Stage) {
	diagram.Unstage(stage)
}

func (diagram *Diagram) StageVoid(stage *Stage) {
	diagram.Stage(stage)
}

// for satisfaction of GongStruct interface
func (diagram *Diagram) GetName() (res string) {
	return diagram.Name
}

// for satisfaction of GongStruct interface
func (diagram *Diagram) SetName(name string) {
	diagram.Name = name
}

// Stage puts document to the model stage
func (document *Document) Stage(stage *Stage) *Document {
	__gong__stage(stage.Documents, stage.Document_stagedOrder, stage.Document_orderStaged, &stage.DocumentOrder, stage.Documents_mapString, document, document.Name)
	return document
}

// StagePreserveOrder puts document to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DocumentOrder
// - update stage.DocumentOrder accordingly
func (document *Document) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Documents, stage.Document_stagedOrder, stage.Document_orderStaged, &stage.DocumentOrder, stage.Documents_mapString, document, order, document.Name)
}

// Unstage removes document off the model stage
func (document *Document) Unstage(stage *Stage) *Document {
	__gong__unstage(stage.Documents, stage.Documents_mapString, document, document.Name)
	return document
}

// UnstageVoid removes document off the model stage
func (document *Document) UnstageVoid(stage *Stage) {
	document.Unstage(stage)
}

func (document *Document) StageVoid(stage *Stage) {
	document.Stage(stage)
}

// for satisfaction of GongStruct interface
func (document *Document) GetName() (res string) {
	return document.Name
}

// for satisfaction of GongStruct interface
func (document *Document) SetName(name string) {
	document.Name = name
}

// Stage puts documentuse to the model stage
func (documentuse *DocumentUse) Stage(stage *Stage) *DocumentUse {
	__gong__stage(stage.DocumentUses, stage.DocumentUse_stagedOrder, stage.DocumentUse_orderStaged, &stage.DocumentUseOrder, stage.DocumentUses_mapString, documentuse, documentuse.Name)
	return documentuse
}

// StagePreserveOrder puts documentuse to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DocumentUseOrder
// - update stage.DocumentUseOrder accordingly
func (documentuse *DocumentUse) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DocumentUses, stage.DocumentUse_stagedOrder, stage.DocumentUse_orderStaged, &stage.DocumentUseOrder, stage.DocumentUses_mapString, documentuse, order, documentuse.Name)
}

// Unstage removes documentuse off the model stage
func (documentuse *DocumentUse) Unstage(stage *Stage) *DocumentUse {
	__gong__unstage(stage.DocumentUses, stage.DocumentUses_mapString, documentuse, documentuse.Name)
	return documentuse
}

// UnstageVoid removes documentuse off the model stage
func (documentuse *DocumentUse) UnstageVoid(stage *Stage) {
	documentuse.Unstage(stage)
}

func (documentuse *DocumentUse) StageVoid(stage *Stage) {
	documentuse.Stage(stage)
}

// for satisfaction of GongStruct interface
func (documentuse *DocumentUse) GetName() (res string) {
	return documentuse.Name
}

// for satisfaction of GongStruct interface
func (documentuse *DocumentUse) SetName(name string) {
	documentuse.Name = name
}

// Stage puts evolutiondirection to the model stage
func (evolutiondirection *EvolutionDirection) Stage(stage *Stage) *EvolutionDirection {
	__gong__stage(stage.EvolutionDirections, stage.EvolutionDirection_stagedOrder, stage.EvolutionDirection_orderStaged, &stage.EvolutionDirectionOrder, stage.EvolutionDirections_mapString, evolutiondirection, evolutiondirection.Name)
	return evolutiondirection
}

// StagePreserveOrder puts evolutiondirection to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EvolutionDirectionOrder
// - update stage.EvolutionDirectionOrder accordingly
func (evolutiondirection *EvolutionDirection) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.EvolutionDirections, stage.EvolutionDirection_stagedOrder, stage.EvolutionDirection_orderStaged, &stage.EvolutionDirectionOrder, stage.EvolutionDirections_mapString, evolutiondirection, order, evolutiondirection.Name)
}

// Unstage removes evolutiondirection off the model stage
func (evolutiondirection *EvolutionDirection) Unstage(stage *Stage) *EvolutionDirection {
	__gong__unstage(stage.EvolutionDirections, stage.EvolutionDirections_mapString, evolutiondirection, evolutiondirection.Name)
	return evolutiondirection
}

// UnstageVoid removes evolutiondirection off the model stage
func (evolutiondirection *EvolutionDirection) UnstageVoid(stage *Stage) {
	evolutiondirection.Unstage(stage)
}

func (evolutiondirection *EvolutionDirection) StageVoid(stage *Stage) {
	evolutiondirection.Stage(stage)
}

// for satisfaction of GongStruct interface
func (evolutiondirection *EvolutionDirection) GetName() (res string) {
	return evolutiondirection.Name
}

// for satisfaction of GongStruct interface
func (evolutiondirection *EvolutionDirection) SetName(name string) {
	evolutiondirection.Name = name
}

// Stage puts evolutiondirectionshape to the model stage
func (evolutiondirectionshape *EvolutionDirectionShape) Stage(stage *Stage) *EvolutionDirectionShape {
	__gong__stage(stage.EvolutionDirectionShapes, stage.EvolutionDirectionShape_stagedOrder, stage.EvolutionDirectionShape_orderStaged, &stage.EvolutionDirectionShapeOrder, stage.EvolutionDirectionShapes_mapString, evolutiondirectionshape, evolutiondirectionshape.Name)
	return evolutiondirectionshape
}

// StagePreserveOrder puts evolutiondirectionshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EvolutionDirectionShapeOrder
// - update stage.EvolutionDirectionShapeOrder accordingly
func (evolutiondirectionshape *EvolutionDirectionShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.EvolutionDirectionShapes, stage.EvolutionDirectionShape_stagedOrder, stage.EvolutionDirectionShape_orderStaged, &stage.EvolutionDirectionShapeOrder, stage.EvolutionDirectionShapes_mapString, evolutiondirectionshape, order, evolutiondirectionshape.Name)
}

// Unstage removes evolutiondirectionshape off the model stage
func (evolutiondirectionshape *EvolutionDirectionShape) Unstage(stage *Stage) *EvolutionDirectionShape {
	__gong__unstage(stage.EvolutionDirectionShapes, stage.EvolutionDirectionShapes_mapString, evolutiondirectionshape, evolutiondirectionshape.Name)
	return evolutiondirectionshape
}

// UnstageVoid removes evolutiondirectionshape off the model stage
func (evolutiondirectionshape *EvolutionDirectionShape) UnstageVoid(stage *Stage) {
	evolutiondirectionshape.Unstage(stage)
}

func (evolutiondirectionshape *EvolutionDirectionShape) StageVoid(stage *Stage) {
	evolutiondirectionshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (evolutiondirectionshape *EvolutionDirectionShape) GetName() (res string) {
	return evolutiondirectionshape.Name
}

// for satisfaction of GongStruct interface
func (evolutiondirectionshape *EvolutionDirectionShape) SetName(name string) {
	evolutiondirectionshape.Name = name
}

// Stage puts foo to the model stage
func (foo *Foo) Stage(stage *Stage) *Foo {
	__gong__stage(stage.Foos, stage.Foo_stagedOrder, stage.Foo_orderStaged, &stage.FooOrder, stage.Foos_mapString, foo, foo.Name)
	return foo
}

// StagePreserveOrder puts foo to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FooOrder
// - update stage.FooOrder accordingly
func (foo *Foo) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Foos, stage.Foo_stagedOrder, stage.Foo_orderStaged, &stage.FooOrder, stage.Foos_mapString, foo, order, foo.Name)
}

// Unstage removes foo off the model stage
func (foo *Foo) Unstage(stage *Stage) *Foo {
	__gong__unstage(stage.Foos, stage.Foos_mapString, foo, foo.Name)
	return foo
}

// UnstageVoid removes foo off the model stage
func (foo *Foo) UnstageVoid(stage *Stage) {
	foo.Unstage(stage)
}

func (foo *Foo) StageVoid(stage *Stage) {
	foo.Stage(stage)
}

// for satisfaction of GongStruct interface
func (foo *Foo) GetName() (res string) {
	return foo.Name
}

// for satisfaction of GongStruct interface
func (foo *Foo) SetName(name string) {
	foo.Name = name
}

// Stage puts geoobject to the model stage
func (geoobject *GeoObject) Stage(stage *Stage) *GeoObject {
	__gong__stage(stage.GeoObjects, stage.GeoObject_stagedOrder, stage.GeoObject_orderStaged, &stage.GeoObjectOrder, stage.GeoObjects_mapString, geoobject, geoobject.Name)
	return geoobject
}

// StagePreserveOrder puts geoobject to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GeoObjectOrder
// - update stage.GeoObjectOrder accordingly
func (geoobject *GeoObject) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.GeoObjects, stage.GeoObject_stagedOrder, stage.GeoObject_orderStaged, &stage.GeoObjectOrder, stage.GeoObjects_mapString, geoobject, order, geoobject.Name)
}

// Unstage removes geoobject off the model stage
func (geoobject *GeoObject) Unstage(stage *Stage) *GeoObject {
	__gong__unstage(stage.GeoObjects, stage.GeoObjects_mapString, geoobject, geoobject.Name)
	return geoobject
}

// UnstageVoid removes geoobject off the model stage
func (geoobject *GeoObject) UnstageVoid(stage *Stage) {
	geoobject.Unstage(stage)
}

func (geoobject *GeoObject) StageVoid(stage *Stage) {
	geoobject.Stage(stage)
}

// for satisfaction of GongStruct interface
func (geoobject *GeoObject) GetName() (res string) {
	return geoobject.Name
}

// for satisfaction of GongStruct interface
func (geoobject *GeoObject) SetName(name string) {
	geoobject.Name = name
}

// Stage puts geoobjectuse to the model stage
func (geoobjectuse *GeoObjectUse) Stage(stage *Stage) *GeoObjectUse {
	__gong__stage(stage.GeoObjectUses, stage.GeoObjectUse_stagedOrder, stage.GeoObjectUse_orderStaged, &stage.GeoObjectUseOrder, stage.GeoObjectUses_mapString, geoobjectuse, geoobjectuse.Name)
	return geoobjectuse
}

// StagePreserveOrder puts geoobjectuse to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GeoObjectUseOrder
// - update stage.GeoObjectUseOrder accordingly
func (geoobjectuse *GeoObjectUse) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.GeoObjectUses, stage.GeoObjectUse_stagedOrder, stage.GeoObjectUse_orderStaged, &stage.GeoObjectUseOrder, stage.GeoObjectUses_mapString, geoobjectuse, order, geoobjectuse.Name)
}

// Unstage removes geoobjectuse off the model stage
func (geoobjectuse *GeoObjectUse) Unstage(stage *Stage) *GeoObjectUse {
	__gong__unstage(stage.GeoObjectUses, stage.GeoObjectUses_mapString, geoobjectuse, geoobjectuse.Name)
	return geoobjectuse
}

// UnstageVoid removes geoobjectuse off the model stage
func (geoobjectuse *GeoObjectUse) UnstageVoid(stage *Stage) {
	geoobjectuse.Unstage(stage)
}

func (geoobjectuse *GeoObjectUse) StageVoid(stage *Stage) {
	geoobjectuse.Stage(stage)
}

// for satisfaction of GongStruct interface
func (geoobjectuse *GeoObjectUse) GetName() (res string) {
	return geoobjectuse.Name
}

// for satisfaction of GongStruct interface
func (geoobjectuse *GeoObjectUse) SetName(name string) {
	geoobjectuse.Name = name
}

// Stage puts group to the model stage
func (group *Group) Stage(stage *Stage) *Group {
	__gong__stage(stage.Groups, stage.Group_stagedOrder, stage.Group_orderStaged, &stage.GroupOrder, stage.Groups_mapString, group, group.Name)
	return group
}

// StagePreserveOrder puts group to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GroupOrder
// - update stage.GroupOrder accordingly
func (group *Group) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Groups, stage.Group_stagedOrder, stage.Group_orderStaged, &stage.GroupOrder, stage.Groups_mapString, group, order, group.Name)
}

// Unstage removes group off the model stage
func (group *Group) Unstage(stage *Stage) *Group {
	__gong__unstage(stage.Groups, stage.Groups_mapString, group, group.Name)
	return group
}

// UnstageVoid removes group off the model stage
func (group *Group) UnstageVoid(stage *Stage) {
	group.Unstage(stage)
}

func (group *Group) StageVoid(stage *Stage) {
	group.Stage(stage)
}

// for satisfaction of GongStruct interface
func (group *Group) GetName() (res string) {
	return group.Name
}

// for satisfaction of GongStruct interface
func (group *Group) SetName(name string) {
	group.Name = name
}

// Stage puts groupuse to the model stage
func (groupuse *GroupUse) Stage(stage *Stage) *GroupUse {
	__gong__stage(stage.GroupUses, stage.GroupUse_stagedOrder, stage.GroupUse_orderStaged, &stage.GroupUseOrder, stage.GroupUses_mapString, groupuse, groupuse.Name)
	return groupuse
}

// StagePreserveOrder puts groupuse to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GroupUseOrder
// - update stage.GroupUseOrder accordingly
func (groupuse *GroupUse) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.GroupUses, stage.GroupUse_stagedOrder, stage.GroupUse_orderStaged, &stage.GroupUseOrder, stage.GroupUses_mapString, groupuse, order, groupuse.Name)
}

// Unstage removes groupuse off the model stage
func (groupuse *GroupUse) Unstage(stage *Stage) *GroupUse {
	__gong__unstage(stage.GroupUses, stage.GroupUses_mapString, groupuse, groupuse.Name)
	return groupuse
}

// UnstageVoid removes groupuse off the model stage
func (groupuse *GroupUse) UnstageVoid(stage *Stage) {
	groupuse.Unstage(stage)
}

func (groupuse *GroupUse) StageVoid(stage *Stage) {
	groupuse.Stage(stage)
}

// for satisfaction of GongStruct interface
func (groupuse *GroupUse) GetName() (res string) {
	return groupuse.Name
}

// for satisfaction of GongStruct interface
func (groupuse *GroupUse) SetName(name string) {
	groupuse.Name = name
}

// Stage puts library to the model stage
func (library *Library) Stage(stage *Stage) *Library {
	__gong__stage(stage.Librarys, stage.Library_stagedOrder, stage.Library_orderStaged, &stage.LibraryOrder, stage.Librarys_mapString, library, library.Name)
	return library
}

// StagePreserveOrder puts library to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LibraryOrder
// - update stage.LibraryOrder accordingly
func (library *Library) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Librarys, stage.Library_stagedOrder, stage.Library_orderStaged, &stage.LibraryOrder, stage.Librarys_mapString, library, order, library.Name)
}

// Unstage removes library off the model stage
func (library *Library) Unstage(stage *Stage) *Library {
	__gong__unstage(stage.Librarys, stage.Librarys_mapString, library, library.Name)
	return library
}

// UnstageVoid removes library off the model stage
func (library *Library) UnstageVoid(stage *Stage) {
	library.Unstage(stage)
}

func (library *Library) StageVoid(stage *Stage) {
	library.Stage(stage)
}

// for satisfaction of GongStruct interface
func (library *Library) GetName() (res string) {
	return library.Name
}

// for satisfaction of GongStruct interface
func (library *Library) SetName(name string) {
	library.Name = name
}

// Stage puts mapobject to the model stage
func (mapobject *MapObject) Stage(stage *Stage) *MapObject {
	__gong__stage(stage.MapObjects, stage.MapObject_stagedOrder, stage.MapObject_orderStaged, &stage.MapObjectOrder, stage.MapObjects_mapString, mapobject, mapobject.Name)
	return mapobject
}

// StagePreserveOrder puts mapobject to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MapObjectOrder
// - update stage.MapObjectOrder accordingly
func (mapobject *MapObject) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MapObjects, stage.MapObject_stagedOrder, stage.MapObject_orderStaged, &stage.MapObjectOrder, stage.MapObjects_mapString, mapobject, order, mapobject.Name)
}

// Unstage removes mapobject off the model stage
func (mapobject *MapObject) Unstage(stage *Stage) *MapObject {
	__gong__unstage(stage.MapObjects, stage.MapObjects_mapString, mapobject, mapobject.Name)
	return mapobject
}

// UnstageVoid removes mapobject off the model stage
func (mapobject *MapObject) UnstageVoid(stage *Stage) {
	mapobject.Unstage(stage)
}

func (mapobject *MapObject) StageVoid(stage *Stage) {
	mapobject.Stage(stage)
}

// for satisfaction of GongStruct interface
func (mapobject *MapObject) GetName() (res string) {
	return mapobject.Name
}

// for satisfaction of GongStruct interface
func (mapobject *MapObject) SetName(name string) {
	mapobject.Name = name
}

// Stage puts mapobjectuse to the model stage
func (mapobjectuse *MapObjectUse) Stage(stage *Stage) *MapObjectUse {
	__gong__stage(stage.MapObjectUses, stage.MapObjectUse_stagedOrder, stage.MapObjectUse_orderStaged, &stage.MapObjectUseOrder, stage.MapObjectUses_mapString, mapobjectuse, mapobjectuse.Name)
	return mapobjectuse
}

// StagePreserveOrder puts mapobjectuse to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MapObjectUseOrder
// - update stage.MapObjectUseOrder accordingly
func (mapobjectuse *MapObjectUse) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MapObjectUses, stage.MapObjectUse_stagedOrder, stage.MapObjectUse_orderStaged, &stage.MapObjectUseOrder, stage.MapObjectUses_mapString, mapobjectuse, order, mapobjectuse.Name)
}

// Unstage removes mapobjectuse off the model stage
func (mapobjectuse *MapObjectUse) Unstage(stage *Stage) *MapObjectUse {
	__gong__unstage(stage.MapObjectUses, stage.MapObjectUses_mapString, mapobjectuse, mapobjectuse.Name)
	return mapobjectuse
}

// UnstageVoid removes mapobjectuse off the model stage
func (mapobjectuse *MapObjectUse) UnstageVoid(stage *Stage) {
	mapobjectuse.Unstage(stage)
}

func (mapobjectuse *MapObjectUse) StageVoid(stage *Stage) {
	mapobjectuse.Stage(stage)
}

// for satisfaction of GongStruct interface
func (mapobjectuse *MapObjectUse) GetName() (res string) {
	return mapobjectuse.Name
}

// for satisfaction of GongStruct interface
func (mapobjectuse *MapObjectUse) SetName(name string) {
	mapobjectuse.Name = name
}

// Stage puts parameter to the model stage
func (parameter *Parameter) Stage(stage *Stage) *Parameter {
	__gong__stage(stage.Parameters, stage.Parameter_stagedOrder, stage.Parameter_orderStaged, &stage.ParameterOrder, stage.Parameters_mapString, parameter, parameter.Name)
	return parameter
}

// StagePreserveOrder puts parameter to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ParameterOrder
// - update stage.ParameterOrder accordingly
func (parameter *Parameter) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Parameters, stage.Parameter_stagedOrder, stage.Parameter_orderStaged, &stage.ParameterOrder, stage.Parameters_mapString, parameter, order, parameter.Name)
}

// Unstage removes parameter off the model stage
func (parameter *Parameter) Unstage(stage *Stage) *Parameter {
	__gong__unstage(stage.Parameters, stage.Parameters_mapString, parameter, parameter.Name)
	return parameter
}

// UnstageVoid removes parameter off the model stage
func (parameter *Parameter) UnstageVoid(stage *Stage) {
	parameter.Unstage(stage)
}

func (parameter *Parameter) StageVoid(stage *Stage) {
	parameter.Stage(stage)
}

// for satisfaction of GongStruct interface
func (parameter *Parameter) GetName() (res string) {
	return parameter.Name
}

// for satisfaction of GongStruct interface
func (parameter *Parameter) SetName(name string) {
	parameter.Name = name
}

// Stage puts parametercategory to the model stage
func (parametercategory *ParameterCategory) Stage(stage *Stage) *ParameterCategory {
	__gong__stage(stage.ParameterCategorys, stage.ParameterCategory_stagedOrder, stage.ParameterCategory_orderStaged, &stage.ParameterCategoryOrder, stage.ParameterCategorys_mapString, parametercategory, parametercategory.Name)
	return parametercategory
}

// StagePreserveOrder puts parametercategory to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ParameterCategoryOrder
// - update stage.ParameterCategoryOrder accordingly
func (parametercategory *ParameterCategory) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ParameterCategorys, stage.ParameterCategory_stagedOrder, stage.ParameterCategory_orderStaged, &stage.ParameterCategoryOrder, stage.ParameterCategorys_mapString, parametercategory, order, parametercategory.Name)
}

// Unstage removes parametercategory off the model stage
func (parametercategory *ParameterCategory) Unstage(stage *Stage) *ParameterCategory {
	__gong__unstage(stage.ParameterCategorys, stage.ParameterCategorys_mapString, parametercategory, parametercategory.Name)
	return parametercategory
}

// UnstageVoid removes parametercategory off the model stage
func (parametercategory *ParameterCategory) UnstageVoid(stage *Stage) {
	parametercategory.Unstage(stage)
}

func (parametercategory *ParameterCategory) StageVoid(stage *Stage) {
	parametercategory.Stage(stage)
}

// for satisfaction of GongStruct interface
func (parametercategory *ParameterCategory) GetName() (res string) {
	return parametercategory.Name
}

// for satisfaction of GongStruct interface
func (parametercategory *ParameterCategory) SetName(name string) {
	parametercategory.Name = name
}

// Stage puts parametercategoryuse to the model stage
func (parametercategoryuse *ParameterCategoryUse) Stage(stage *Stage) *ParameterCategoryUse {
	__gong__stage(stage.ParameterCategoryUses, stage.ParameterCategoryUse_stagedOrder, stage.ParameterCategoryUse_orderStaged, &stage.ParameterCategoryUseOrder, stage.ParameterCategoryUses_mapString, parametercategoryuse, parametercategoryuse.Name)
	return parametercategoryuse
}

// StagePreserveOrder puts parametercategoryuse to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ParameterCategoryUseOrder
// - update stage.ParameterCategoryUseOrder accordingly
func (parametercategoryuse *ParameterCategoryUse) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ParameterCategoryUses, stage.ParameterCategoryUse_stagedOrder, stage.ParameterCategoryUse_orderStaged, &stage.ParameterCategoryUseOrder, stage.ParameterCategoryUses_mapString, parametercategoryuse, order, parametercategoryuse.Name)
}

// Unstage removes parametercategoryuse off the model stage
func (parametercategoryuse *ParameterCategoryUse) Unstage(stage *Stage) *ParameterCategoryUse {
	__gong__unstage(stage.ParameterCategoryUses, stage.ParameterCategoryUses_mapString, parametercategoryuse, parametercategoryuse.Name)
	return parametercategoryuse
}

// UnstageVoid removes parametercategoryuse off the model stage
func (parametercategoryuse *ParameterCategoryUse) UnstageVoid(stage *Stage) {
	parametercategoryuse.Unstage(stage)
}

func (parametercategoryuse *ParameterCategoryUse) StageVoid(stage *Stage) {
	parametercategoryuse.Stage(stage)
}

// for satisfaction of GongStruct interface
func (parametercategoryuse *ParameterCategoryUse) GetName() (res string) {
	return parametercategoryuse.Name
}

// for satisfaction of GongStruct interface
func (parametercategoryuse *ParameterCategoryUse) SetName(name string) {
	parametercategoryuse.Name = name
}

// Stage puts parametershape to the model stage
func (parametershape *ParameterShape) Stage(stage *Stage) *ParameterShape {
	__gong__stage(stage.ParameterShapes, stage.ParameterShape_stagedOrder, stage.ParameterShape_orderStaged, &stage.ParameterShapeOrder, stage.ParameterShapes_mapString, parametershape, parametershape.Name)
	return parametershape
}

// StagePreserveOrder puts parametershape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ParameterShapeOrder
// - update stage.ParameterShapeOrder accordingly
func (parametershape *ParameterShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ParameterShapes, stage.ParameterShape_stagedOrder, stage.ParameterShape_orderStaged, &stage.ParameterShapeOrder, stage.ParameterShapes_mapString, parametershape, order, parametershape.Name)
}

// Unstage removes parametershape off the model stage
func (parametershape *ParameterShape) Unstage(stage *Stage) *ParameterShape {
	__gong__unstage(stage.ParameterShapes, stage.ParameterShapes_mapString, parametershape, parametershape.Name)
	return parametershape
}

// UnstageVoid removes parametershape off the model stage
func (parametershape *ParameterShape) UnstageVoid(stage *Stage) {
	parametershape.Unstage(stage)
}

func (parametershape *ParameterShape) StageVoid(stage *Stage) {
	parametershape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (parametershape *ParameterShape) GetName() (res string) {
	return parametershape.Name
}

// for satisfaction of GongStruct interface
func (parametershape *ParameterShape) SetName(name string) {
	parametershape.Name = name
}

// Stage puts parametersaggregate to the model stage
func (parametersaggregate *ParametersAggregate) Stage(stage *Stage) *ParametersAggregate {
	__gong__stage(stage.ParametersAggregates, stage.ParametersAggregate_stagedOrder, stage.ParametersAggregate_orderStaged, &stage.ParametersAggregateOrder, stage.ParametersAggregates_mapString, parametersaggregate, parametersaggregate.Name)
	return parametersaggregate
}

// StagePreserveOrder puts parametersaggregate to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ParametersAggregateOrder
// - update stage.ParametersAggregateOrder accordingly
func (parametersaggregate *ParametersAggregate) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ParametersAggregates, stage.ParametersAggregate_stagedOrder, stage.ParametersAggregate_orderStaged, &stage.ParametersAggregateOrder, stage.ParametersAggregates_mapString, parametersaggregate, order, parametersaggregate.Name)
}

// Unstage removes parametersaggregate off the model stage
func (parametersaggregate *ParametersAggregate) Unstage(stage *Stage) *ParametersAggregate {
	__gong__unstage(stage.ParametersAggregates, stage.ParametersAggregates_mapString, parametersaggregate, parametersaggregate.Name)
	return parametersaggregate
}

// UnstageVoid removes parametersaggregate off the model stage
func (parametersaggregate *ParametersAggregate) UnstageVoid(stage *Stage) {
	parametersaggregate.Unstage(stage)
}

func (parametersaggregate *ParametersAggregate) StageVoid(stage *Stage) {
	parametersaggregate.Stage(stage)
}

// for satisfaction of GongStruct interface
func (parametersaggregate *ParametersAggregate) GetName() (res string) {
	return parametersaggregate.Name
}

// for satisfaction of GongStruct interface
func (parametersaggregate *ParametersAggregate) SetName(name string) {
	parametersaggregate.Name = name
}

// Stage puts parametersaggregateshape to the model stage
func (parametersaggregateshape *ParametersAggregateShape) Stage(stage *Stage) *ParametersAggregateShape {
	__gong__stage(stage.ParametersAggregateShapes, stage.ParametersAggregateShape_stagedOrder, stage.ParametersAggregateShape_orderStaged, &stage.ParametersAggregateShapeOrder, stage.ParametersAggregateShapes_mapString, parametersaggregateshape, parametersaggregateshape.Name)
	return parametersaggregateshape
}

// StagePreserveOrder puts parametersaggregateshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ParametersAggregateShapeOrder
// - update stage.ParametersAggregateShapeOrder accordingly
func (parametersaggregateshape *ParametersAggregateShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ParametersAggregateShapes, stage.ParametersAggregateShape_stagedOrder, stage.ParametersAggregateShape_orderStaged, &stage.ParametersAggregateShapeOrder, stage.ParametersAggregateShapes_mapString, parametersaggregateshape, order, parametersaggregateshape.Name)
}

// Unstage removes parametersaggregateshape off the model stage
func (parametersaggregateshape *ParametersAggregateShape) Unstage(stage *Stage) *ParametersAggregateShape {
	__gong__unstage(stage.ParametersAggregateShapes, stage.ParametersAggregateShapes_mapString, parametersaggregateshape, parametersaggregateshape.Name)
	return parametersaggregateshape
}

// UnstageVoid removes parametersaggregateshape off the model stage
func (parametersaggregateshape *ParametersAggregateShape) UnstageVoid(stage *Stage) {
	parametersaggregateshape.Unstage(stage)
}

func (parametersaggregateshape *ParametersAggregateShape) StageVoid(stage *Stage) {
	parametersaggregateshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (parametersaggregateshape *ParametersAggregateShape) GetName() (res string) {
	return parametersaggregateshape.Name
}

// for satisfaction of GongStruct interface
func (parametersaggregateshape *ParametersAggregateShape) SetName(name string) {
	parametersaggregateshape.Name = name
}

// Stage puts position to the model stage
func (position *Position) Stage(stage *Stage) *Position {
	__gong__stage(stage.Positions, stage.Position_stagedOrder, stage.Position_orderStaged, &stage.PositionOrder, stage.Positions_mapString, position, position.Name)
	return position
}

// StagePreserveOrder puts position to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PositionOrder
// - update stage.PositionOrder accordingly
func (position *Position) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Positions, stage.Position_stagedOrder, stage.Position_orderStaged, &stage.PositionOrder, stage.Positions_mapString, position, order, position.Name)
}

// Unstage removes position off the model stage
func (position *Position) Unstage(stage *Stage) *Position {
	__gong__unstage(stage.Positions, stage.Positions_mapString, position, position.Name)
	return position
}

// UnstageVoid removes position off the model stage
func (position *Position) UnstageVoid(stage *Stage) {
	position.Unstage(stage)
}

func (position *Position) StageVoid(stage *Stage) {
	position.Stage(stage)
}

// for satisfaction of GongStruct interface
func (position *Position) GetName() (res string) {
	return position.Name
}

// for satisfaction of GongStruct interface
func (position *Position) SetName(name string) {
	position.Name = name
}

// Stage puts repository to the model stage
func (repository *Repository) Stage(stage *Stage) *Repository {
	__gong__stage(stage.Repositorys, stage.Repository_stagedOrder, stage.Repository_orderStaged, &stage.RepositoryOrder, stage.Repositorys_mapString, repository, repository.Name)
	return repository
}

// StagePreserveOrder puts repository to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RepositoryOrder
// - update stage.RepositoryOrder accordingly
func (repository *Repository) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Repositorys, stage.Repository_stagedOrder, stage.Repository_orderStaged, &stage.RepositoryOrder, stage.Repositorys_mapString, repository, order, repository.Name)
}

// Unstage removes repository off the model stage
func (repository *Repository) Unstage(stage *Stage) *Repository {
	__gong__unstage(stage.Repositorys, stage.Repositorys_mapString, repository, repository.Name)
	return repository
}

// UnstageVoid removes repository off the model stage
func (repository *Repository) UnstageVoid(stage *Stage) {
	repository.Unstage(stage)
}

func (repository *Repository) StageVoid(stage *Stage) {
	repository.Stage(stage)
}

// for satisfaction of GongStruct interface
func (repository *Repository) GetName() (res string) {
	return repository.Name
}

// for satisfaction of GongStruct interface
func (repository *Repository) SetName(name string) {
	repository.Name = name
}

// Stage puts scenario to the model stage
func (scenario *Scenario) Stage(stage *Stage) *Scenario {
	__gong__stage(stage.Scenarios, stage.Scenario_stagedOrder, stage.Scenario_orderStaged, &stage.ScenarioOrder, stage.Scenarios_mapString, scenario, scenario.Name)
	return scenario
}

// StagePreserveOrder puts scenario to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ScenarioOrder
// - update stage.ScenarioOrder accordingly
func (scenario *Scenario) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Scenarios, stage.Scenario_stagedOrder, stage.Scenario_orderStaged, &stage.ScenarioOrder, stage.Scenarios_mapString, scenario, order, scenario.Name)
}

// Unstage removes scenario off the model stage
func (scenario *Scenario) Unstage(stage *Stage) *Scenario {
	__gong__unstage(stage.Scenarios, stage.Scenarios_mapString, scenario, scenario.Name)
	return scenario
}

// UnstageVoid removes scenario off the model stage
func (scenario *Scenario) UnstageVoid(stage *Stage) {
	scenario.Unstage(stage)
}

func (scenario *Scenario) StageVoid(stage *Stage) {
	scenario.Stage(stage)
}

// for satisfaction of GongStruct interface
func (scenario *Scenario) GetName() (res string) {
	return scenario.Name
}

// for satisfaction of GongStruct interface
func (scenario *Scenario) SetName(name string) {
	scenario.Name = name
}

// Stage puts user to the model stage
func (user *User) Stage(stage *Stage) *User {
	__gong__stage(stage.Users, stage.User_stagedOrder, stage.User_orderStaged, &stage.UserOrder, stage.Users_mapString, user, user.Name)
	return user
}

// StagePreserveOrder puts user to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.UserOrder
// - update stage.UserOrder accordingly
func (user *User) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Users, stage.User_stagedOrder, stage.User_orderStaged, &stage.UserOrder, stage.Users_mapString, user, order, user.Name)
}

// Unstage removes user off the model stage
func (user *User) Unstage(stage *Stage) *User {
	__gong__unstage(stage.Users, stage.Users_mapString, user, user.Name)
	return user
}

// UnstageVoid removes user off the model stage
func (user *User) UnstageVoid(stage *Stage) {
	user.Unstage(stage)
}

func (user *User) StageVoid(stage *Stage) {
	user.Stage(stage)
}

// for satisfaction of GongStruct interface
func (user *User) GetName() (res string) {
	return user.Name
}

// for satisfaction of GongStruct interface
func (user *User) SetName(name string) {
	user.Name = name
}

// Stage puts useruse to the model stage
func (useruse *UserUse) Stage(stage *Stage) *UserUse {
	__gong__stage(stage.UserUses, stage.UserUse_stagedOrder, stage.UserUse_orderStaged, &stage.UserUseOrder, stage.UserUses_mapString, useruse, useruse.Name)
	return useruse
}

// StagePreserveOrder puts useruse to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.UserUseOrder
// - update stage.UserUseOrder accordingly
func (useruse *UserUse) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.UserUses, stage.UserUse_stagedOrder, stage.UserUse_orderStaged, &stage.UserUseOrder, stage.UserUses_mapString, useruse, order, useruse.Name)
}

// Unstage removes useruse off the model stage
func (useruse *UserUse) Unstage(stage *Stage) *UserUse {
	__gong__unstage(stage.UserUses, stage.UserUses_mapString, useruse, useruse.Name)
	return useruse
}

// UnstageVoid removes useruse off the model stage
func (useruse *UserUse) UnstageVoid(stage *Stage) {
	useruse.Unstage(stage)
}

func (useruse *UserUse) StageVoid(stage *Stage) {
	useruse.Stage(stage)
}

// for satisfaction of GongStruct interface
func (useruse *UserUse) GetName() (res string) {
	return useruse.Name
}

// for satisfaction of GongStruct interface
func (useruse *UserUse) SetName(name string) {
	useruse.Name = name
}

// Stage puts workspace to the model stage
func (workspace *Workspace) Stage(stage *Stage) *Workspace {
	__gong__stage(stage.Workspaces, stage.Workspace_stagedOrder, stage.Workspace_orderStaged, &stage.WorkspaceOrder, stage.Workspaces_mapString, workspace, workspace.Name)
	return workspace
}

// StagePreserveOrder puts workspace to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.WorkspaceOrder
// - update stage.WorkspaceOrder accordingly
func (workspace *Workspace) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Workspaces, stage.Workspace_stagedOrder, stage.Workspace_orderStaged, &stage.WorkspaceOrder, stage.Workspaces_mapString, workspace, order, workspace.Name)
}

// Unstage removes workspace off the model stage
func (workspace *Workspace) Unstage(stage *Stage) *Workspace {
	__gong__unstage(stage.Workspaces, stage.Workspaces_mapString, workspace, workspace.Name)
	return workspace
}

// UnstageVoid removes workspace off the model stage
func (workspace *Workspace) UnstageVoid(stage *Stage) {
	workspace.Unstage(stage)
}

func (workspace *Workspace) StageVoid(stage *Stage) {
	workspace.Stage(stage)
}

// for satisfaction of GongStruct interface
func (workspace *Workspace) GetName() (res string) {
	return workspace.Name
}

// for satisfaction of GongStruct interface
func (workspace *Workspace) SetName(name string) {
	workspace.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.ActorStates, &stage.ActorStates_mapString, &stage.ActorState_stagedOrder, &stage.ActorStateOrder)

	__gong__resetStageType(&stage.ActorStateShapes, &stage.ActorStateShapes_mapString, &stage.ActorStateShape_stagedOrder, &stage.ActorStateShapeOrder)

	__gong__resetStageType(&stage.ActorStateTransitions, &stage.ActorStateTransitions_mapString, &stage.ActorStateTransition_stagedOrder, &stage.ActorStateTransitionOrder)

	__gong__resetStageType(&stage.ActorStateTransitionShapes, &stage.ActorStateTransitionShapes_mapString, &stage.ActorStateTransitionShape_stagedOrder, &stage.ActorStateTransitionShapeOrder)

	__gong__resetStageType(&stage.Analysiss, &stage.Analysiss_mapString, &stage.Analysis_stagedOrder, &stage.AnalysisOrder)

	__gong__resetStageType(&stage.ControlPointShapes, &stage.ControlPointShapes_mapString, &stage.ControlPointShape_stagedOrder, &stage.ControlPointShapeOrder)

	__gong__resetStageType(&stage.Diagrams, &stage.Diagrams_mapString, &stage.Diagram_stagedOrder, &stage.DiagramOrder)

	__gong__resetStageType(&stage.Documents, &stage.Documents_mapString, &stage.Document_stagedOrder, &stage.DocumentOrder)

	__gong__resetStageType(&stage.DocumentUses, &stage.DocumentUses_mapString, &stage.DocumentUse_stagedOrder, &stage.DocumentUseOrder)

	__gong__resetStageType(&stage.EvolutionDirections, &stage.EvolutionDirections_mapString, &stage.EvolutionDirection_stagedOrder, &stage.EvolutionDirectionOrder)

	__gong__resetStageType(&stage.EvolutionDirectionShapes, &stage.EvolutionDirectionShapes_mapString, &stage.EvolutionDirectionShape_stagedOrder, &stage.EvolutionDirectionShapeOrder)

	__gong__resetStageType(&stage.Foos, &stage.Foos_mapString, &stage.Foo_stagedOrder, &stage.FooOrder)

	__gong__resetStageType(&stage.GeoObjects, &stage.GeoObjects_mapString, &stage.GeoObject_stagedOrder, &stage.GeoObjectOrder)

	__gong__resetStageType(&stage.GeoObjectUses, &stage.GeoObjectUses_mapString, &stage.GeoObjectUse_stagedOrder, &stage.GeoObjectUseOrder)

	__gong__resetStageType(&stage.Groups, &stage.Groups_mapString, &stage.Group_stagedOrder, &stage.GroupOrder)

	__gong__resetStageType(&stage.GroupUses, &stage.GroupUses_mapString, &stage.GroupUse_stagedOrder, &stage.GroupUseOrder)

	__gong__resetStageType(&stage.Librarys, &stage.Librarys_mapString, &stage.Library_stagedOrder, &stage.LibraryOrder)

	__gong__resetStageType(&stage.MapObjects, &stage.MapObjects_mapString, &stage.MapObject_stagedOrder, &stage.MapObjectOrder)

	__gong__resetStageType(&stage.MapObjectUses, &stage.MapObjectUses_mapString, &stage.MapObjectUse_stagedOrder, &stage.MapObjectUseOrder)

	__gong__resetStageType(&stage.Parameters, &stage.Parameters_mapString, &stage.Parameter_stagedOrder, &stage.ParameterOrder)

	__gong__resetStageType(&stage.ParameterCategorys, &stage.ParameterCategorys_mapString, &stage.ParameterCategory_stagedOrder, &stage.ParameterCategoryOrder)

	__gong__resetStageType(&stage.ParameterCategoryUses, &stage.ParameterCategoryUses_mapString, &stage.ParameterCategoryUse_stagedOrder, &stage.ParameterCategoryUseOrder)

	__gong__resetStageType(&stage.ParameterShapes, &stage.ParameterShapes_mapString, &stage.ParameterShape_stagedOrder, &stage.ParameterShapeOrder)

	__gong__resetStageType(&stage.ParametersAggregates, &stage.ParametersAggregates_mapString, &stage.ParametersAggregate_stagedOrder, &stage.ParametersAggregateOrder)

	__gong__resetStageType(&stage.ParametersAggregateShapes, &stage.ParametersAggregateShapes_mapString, &stage.ParametersAggregateShape_stagedOrder, &stage.ParametersAggregateShapeOrder)

	__gong__resetStageType(&stage.Positions, &stage.Positions_mapString, &stage.Position_stagedOrder, &stage.PositionOrder)

	__gong__resetStageType(&stage.Repositorys, &stage.Repositorys_mapString, &stage.Repository_stagedOrder, &stage.RepositoryOrder)

	__gong__resetStageType(&stage.Scenarios, &stage.Scenarios_mapString, &stage.Scenario_stagedOrder, &stage.ScenarioOrder)

	__gong__resetStageType(&stage.Users, &stage.Users_mapString, &stage.User_stagedOrder, &stage.UserOrder)

	__gong__resetStageType(&stage.UserUses, &stage.UserUses_mapString, &stage.UserUse_stagedOrder, &stage.UserUseOrder)

	__gong__resetStageType(&stage.Workspaces, &stage.Workspaces_mapString, &stage.Workspace_stagedOrder, &stage.WorkspaceOrder)

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct any

type GongstructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

type GongtructBasicField = GongstructBasicField

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type GongstructIF interface {
	GetName() string
	SetName(string)
	StageVoid(*Stage)
	UnstageVoid(stage *Stage)
	GongGetFieldHeaders() []GongFieldHeader
	GongGetFieldValue(fieldName string, stage *Stage) GongFieldValue
	GongGetGongstructName() string
	GongGetOrder(stage *Stage) uint
	GongGetReferenceIdentifier(stage *Stage) string
	GongGetIdentifier(stage *Stage) string
	GongCopy() GongstructIF
	GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) string
	GongGetUUID(stage *Stage) string
	GongAfterCreateFromFront(stage *Stage)
	GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF)
	GongAfterDeleteFromFront(stage *Stage, front GongstructIF)
	GongIsStaged(stage *Stage) bool
	GongStageBranch(stage *Stage)
	GongUnstageBranch(stage *Stage)
}
type GongstructPtr interface {
	GongstructIF
	comparable
}

type PointerToGongstruct = GongstructPtr

func GongCompareGongstructByName[T GongstructPtr](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func GongSortGongstructSetByName[T GongstructPtr](set map[T]struct{}) (sortedSlice []T) {
	for key := range set {
		sortedSlice = append(sortedSlice, key)
	}
	slices.SortFunc(sortedSlice, GongCompareGongstructByName)

	return
}

// GetInstancesSorted is the Stage method returning sorted instances of a gongstruct.
func (stage *Stage) GetInstancesSorted[T GongstructPtr]() (sortedSlice []T) {
	set := stage.GetInstancesSet[T]()
	sortedSlice = GongSortGongstructSetByName(*set)

	return
}

// GetInstancesMapByName is the Stage method returning a map of staged instances by their name.
func (stage *Stage) GetInstancesMapByName[Type GongstructIF]() map[string]Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *ActorState:
		return any(stage.ActorStates_mapString).(map[string]Type)
	case *ActorStateShape:
		return any(stage.ActorStateShapes_mapString).(map[string]Type)
	case *ActorStateTransition:
		return any(stage.ActorStateTransitions_mapString).(map[string]Type)
	case *ActorStateTransitionShape:
		return any(stage.ActorStateTransitionShapes_mapString).(map[string]Type)
	case *Analysis:
		return any(stage.Analysiss_mapString).(map[string]Type)
	case *ControlPointShape:
		return any(stage.ControlPointShapes_mapString).(map[string]Type)
	case *Diagram:
		return any(stage.Diagrams_mapString).(map[string]Type)
	case *Document:
		return any(stage.Documents_mapString).(map[string]Type)
	case *DocumentUse:
		return any(stage.DocumentUses_mapString).(map[string]Type)
	case *EvolutionDirection:
		return any(stage.EvolutionDirections_mapString).(map[string]Type)
	case *EvolutionDirectionShape:
		return any(stage.EvolutionDirectionShapes_mapString).(map[string]Type)
	case *Foo:
		return any(stage.Foos_mapString).(map[string]Type)
	case *GeoObject:
		return any(stage.GeoObjects_mapString).(map[string]Type)
	case *GeoObjectUse:
		return any(stage.GeoObjectUses_mapString).(map[string]Type)
	case *Group:
		return any(stage.Groups_mapString).(map[string]Type)
	case *GroupUse:
		return any(stage.GroupUses_mapString).(map[string]Type)
	case *Library:
		return any(stage.Librarys_mapString).(map[string]Type)
	case *MapObject:
		return any(stage.MapObjects_mapString).(map[string]Type)
	case *MapObjectUse:
		return any(stage.MapObjectUses_mapString).(map[string]Type)
	case *Parameter:
		return any(stage.Parameters_mapString).(map[string]Type)
	case *ParameterCategory:
		return any(stage.ParameterCategorys_mapString).(map[string]Type)
	case *ParameterCategoryUse:
		return any(stage.ParameterCategoryUses_mapString).(map[string]Type)
	case *ParameterShape:
		return any(stage.ParameterShapes_mapString).(map[string]Type)
	case *ParametersAggregate:
		return any(stage.ParametersAggregates_mapString).(map[string]Type)
	case *ParametersAggregateShape:
		return any(stage.ParametersAggregateShapes_mapString).(map[string]Type)
	case *Position:
		return any(stage.Positions_mapString).(map[string]Type)
	case *Repository:
		return any(stage.Repositorys_mapString).(map[string]Type)
	case *Scenario:
		return any(stage.Scenarios_mapString).(map[string]Type)
	case *User:
		return any(stage.Users_mapString).(map[string]Type)
	case *UserUse:
		return any(stage.UserUses_mapString).(map[string]Type)
	case *Workspace:
		return any(stage.Workspaces_mapString).(map[string]Type)
	default:
		return nil
	}
}

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *ActorState:
		return any(&stage.ActorStates).(*map[Type]struct{})
	case *ActorStateShape:
		return any(&stage.ActorStateShapes).(*map[Type]struct{})
	case *ActorStateTransition:
		return any(&stage.ActorStateTransitions).(*map[Type]struct{})
	case *ActorStateTransitionShape:
		return any(&stage.ActorStateTransitionShapes).(*map[Type]struct{})
	case *Analysis:
		return any(&stage.Analysiss).(*map[Type]struct{})
	case *ControlPointShape:
		return any(&stage.ControlPointShapes).(*map[Type]struct{})
	case *Diagram:
		return any(&stage.Diagrams).(*map[Type]struct{})
	case *Document:
		return any(&stage.Documents).(*map[Type]struct{})
	case *DocumentUse:
		return any(&stage.DocumentUses).(*map[Type]struct{})
	case *EvolutionDirection:
		return any(&stage.EvolutionDirections).(*map[Type]struct{})
	case *EvolutionDirectionShape:
		return any(&stage.EvolutionDirectionShapes).(*map[Type]struct{})
	case *Foo:
		return any(&stage.Foos).(*map[Type]struct{})
	case *GeoObject:
		return any(&stage.GeoObjects).(*map[Type]struct{})
	case *GeoObjectUse:
		return any(&stage.GeoObjectUses).(*map[Type]struct{})
	case *Group:
		return any(&stage.Groups).(*map[Type]struct{})
	case *GroupUse:
		return any(&stage.GroupUses).(*map[Type]struct{})
	case *Library:
		return any(&stage.Librarys).(*map[Type]struct{})
	case *MapObject:
		return any(&stage.MapObjects).(*map[Type]struct{})
	case *MapObjectUse:
		return any(&stage.MapObjectUses).(*map[Type]struct{})
	case *Parameter:
		return any(&stage.Parameters).(*map[Type]struct{})
	case *ParameterCategory:
		return any(&stage.ParameterCategorys).(*map[Type]struct{})
	case *ParameterCategoryUse:
		return any(&stage.ParameterCategoryUses).(*map[Type]struct{})
	case *ParameterShape:
		return any(&stage.ParameterShapes).(*map[Type]struct{})
	case *ParametersAggregate:
		return any(&stage.ParametersAggregates).(*map[Type]struct{})
	case *ParametersAggregateShape:
		return any(&stage.ParametersAggregateShapes).(*map[Type]struct{})
	case *Position:
		return any(&stage.Positions).(*map[Type]struct{})
	case *Repository:
		return any(&stage.Repositorys).(*map[Type]struct{})
	case *Scenario:
		return any(&stage.Scenarios).(*map[Type]struct{})
	case *User:
		return any(&stage.Users).(*map[Type]struct{})
	case *UserUse:
		return any(&stage.UserUses).(*map[Type]struct{})
	case *Workspace:
		return any(&stage.Workspaces).(*map[Type]struct{})
	default:
		return nil
	}
}

// GongGetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GongGetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case ActorStateShape:
		return any(&ActorStateShape{
			ActorState: &ActorState{Name: "ActorState"},
		}).(*Type)
	case ActorStateTransition:
		return any(&ActorStateTransition{
			StartState: &ActorState{Name: "StartState"},
			EndState: &ActorState{Name: "EndState"},
			Justifications: []*Parameter{{Name: "Justifications"}},
		}).(*Type)
	case ActorStateTransitionShape:
		return any(&ActorStateTransitionShape{
			ActorStateTransition: &ActorStateTransition{Name: "ActorStateTransition"},
			Start: &ActorStateShape{Name: "Start"},
			End: &ActorStateShape{Name: "End"},
			ControlPointShapes: []*ControlPointShape{{Name: "ControlPointShapes"}},
		}).(*Type)
	case Analysis:
		return any(&Analysis{
			Scenarios: []*Scenario{{Name: "Scenarios"}},
			GroupUse: []*GroupUse{{Name: "GroupUse"}},
			GeoObjectUse: []*GeoObjectUse{{Name: "GeoObjectUse"}},
			MapUse: []*MapObjectUse{{Name: "MapUse"}},
		}).(*Type)
	case Diagram:
		return any(&Diagram{
			EvolutionDirectionShapes: []*EvolutionDirectionShape{{Name: "EvolutionDirectionShapes"}},
			EvolutionDirectionsWhoseNodeIsExpanded: []*EvolutionDirection{{Name: "EvolutionDirectionsWhoseNodeIsExpanded"}},
			ActorStateShapes: []*ActorStateShape{{Name: "ActorStateShapes"}},
			ActorStatesWhoseNodeIsExpanded: []*ActorState{{Name: "ActorStatesWhoseNodeIsExpanded"}},
			ParameterShapes: []*ParameterShape{{Name: "ParameterShapes"}},
			ParametersWhoseNodeIsExpanded: []*Parameter{{Name: "ParametersWhoseNodeIsExpanded"}},
			ScenarioParameterShapes: []*ParametersAggregateShape{{Name: "ScenarioParameterShapes"}},
			ParametersAggregatesWhoseNodeIsExpanded: []*ParametersAggregate{{Name: "ParametersAggregatesWhoseNodeIsExpanded"}},
			ActorStateTransitionShapes: []*ActorStateTransitionShape{{Name: "ActorStateTransitionShapes"}},
			ActorStateTransitionsWhoseNodeIsExpanded: []*ActorStateTransition{{Name: "ActorStateTransitionsWhoseNodeIsExpanded"}},
		}).(*Type)
	case Document:
		return any(&Document{
			GeoObjectUse: []*GeoObjectUse{{Name: "GeoObjectUse"}},
		}).(*Type)
	case DocumentUse:
		return any(&DocumentUse{
			Document: &Document{Name: "Document"},
		}).(*Type)
	case EvolutionDirectionShape:
		return any(&EvolutionDirectionShape{
			EvolutionDirection: &EvolutionDirection{Name: "EvolutionDirection"},
		}).(*Type)
	case GeoObjectUse:
		return any(&GeoObjectUse{
			GeoObject: &GeoObject{Name: "GeoObject"},
		}).(*Type)
	case Group:
		return any(&Group{
			UserUse: []*UserUse{{Name: "UserUse"}},
		}).(*Type)
	case GroupUse:
		return any(&GroupUse{
			Group: &Group{Name: "Group"},
		}).(*Type)
	case Library:
		return any(&Library{
			Analyses: []*Analysis{{Name: "Analyses"}},
			SubLibraries: []*Library{{Name: "SubLibraries"}},
			SubLibrariesWhoseNodeIsExpanded: []*Library{{Name: "SubLibrariesWhoseNodeIsExpanded"}},
		}).(*Type)
	case MapObjectUse:
		return any(&MapObjectUse{
			Map: &MapObject{Name: "Map"},
		}).(*Type)
	case Parameter:
		return any(&Parameter{
			GroupUse: []*GroupUse{{Name: "GroupUse"}},
			DocumentUse: []*DocumentUse{{Name: "DocumentUse"}},
			GeoObjectUse: []*GeoObjectUse{{Name: "GeoObjectUse"}},
		}).(*Type)
	case ParameterCategory:
		return any(&ParameterCategory{
			ParameterUse: []*ParameterShape{{Name: "ParameterUse"}},
		}).(*Type)
	case ParameterCategoryUse:
		return any(&ParameterCategoryUse{
			ParameterCategory: &ParameterCategory{Name: "ParameterCategory"},
		}).(*Type)
	case ParameterShape:
		return any(&ParameterShape{
			Parameter: &Parameter{Name: "Parameter"},
		}).(*Type)
	case ParametersAggregate:
		return any(&ParametersAggregate{
			Parameters: []*Parameter{{Name: "Parameters"}},
		}).(*Type)
	case ParametersAggregateShape:
		return any(&ParametersAggregateShape{
			ScenarioParameter: &ParametersAggregate{Name: "ScenarioParameter"},
		}).(*Type)
	case Repository:
		return any(&Repository{
			ParameterUse: []*ParameterShape{{Name: "ParameterUse"}},
			GroupUse: []*GroupUse{{Name: "GroupUse"}},
		}).(*Type)
	case Scenario:
		return any(&Scenario{
			Diagrams: []*Diagram{{Name: "Diagrams"}},
			ActorStates: []*ActorState{{Name: "ActorStates"}},
			ActorStateTransitions: []*ActorStateTransition{{Name: "ActorStateTransitions"}},
			EvolutionDirections: []*EvolutionDirection{{Name: "EvolutionDirections"}},
			Parameters: []*Parameter{{Name: "Parameters"}},
			ParametersAggretates: []*ParametersAggregate{{Name: "ParametersAggretates"}},
		}).(*Type)
	case UserUse:
		return any(&UserUse{
			User: &User{Name: "User"},
		}).(*Type)
	case Workspace:
		return any(&Workspace{
			SelectedDiagram: &Diagram{Name: "SelectedDiagram"},
			Default_EvolutionDirectionShape: &EvolutionDirectionShape{Name: "Default_EvolutionDirectionShape"},
			Default_ParameterShape: &ParameterShape{Name: "Default_ParameterShape"},
			Default_ScenarioParameterShape: &ParametersAggregateShape{Name: "Default_ScenarioParameterShape"},
			Default_ActorStateShape: &ActorStateShape{Name: "Default_ActorStateShape"},
			Default_ActorStateTransitionShape: &ActorStateTransitionShape{Name: "Default_ActorStateTransitionShape"},
		}).(*Type)
	default:
		return &ret
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
// GetPointerReverseMap is the Stage method for backtrack navigation of pointer associations.
func (stage *Stage) GetPointerReverseMap[Start, End Gongstruct](fieldname string) map[*End][]*Start {
	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of ActorState
	case ActorState:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ActorStateShape
	case ActorStateShape:
		switch fieldname {
		// insertion point for per direct association field
		case "ActorState":
			res := make(map[*ActorState][]*ActorStateShape)
			for actorstateshape := range stage.ActorStateShapes {
				if actorstateshape.ActorState != nil {
					actorstate_ := actorstateshape.ActorState
					var actorstateshapes []*ActorStateShape
					_, ok := res[actorstate_]
					if ok {
						actorstateshapes = res[actorstate_]
					} else {
						actorstateshapes = make([]*ActorStateShape, 0)
					}
					actorstateshapes = append(actorstateshapes, actorstateshape)
					res[actorstate_] = actorstateshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ActorStateTransition
	case ActorStateTransition:
		switch fieldname {
		// insertion point for per direct association field
		case "StartState":
			res := make(map[*ActorState][]*ActorStateTransition)
			for actorstatetransition := range stage.ActorStateTransitions {
				if actorstatetransition.StartState != nil {
					actorstate_ := actorstatetransition.StartState
					var actorstatetransitions []*ActorStateTransition
					_, ok := res[actorstate_]
					if ok {
						actorstatetransitions = res[actorstate_]
					} else {
						actorstatetransitions = make([]*ActorStateTransition, 0)
					}
					actorstatetransitions = append(actorstatetransitions, actorstatetransition)
					res[actorstate_] = actorstatetransitions
				}
			}
			return any(res).(map[*End][]*Start)
		case "EndState":
			res := make(map[*ActorState][]*ActorStateTransition)
			for actorstatetransition := range stage.ActorStateTransitions {
				if actorstatetransition.EndState != nil {
					actorstate_ := actorstatetransition.EndState
					var actorstatetransitions []*ActorStateTransition
					_, ok := res[actorstate_]
					if ok {
						actorstatetransitions = res[actorstate_]
					} else {
						actorstatetransitions = make([]*ActorStateTransition, 0)
					}
					actorstatetransitions = append(actorstatetransitions, actorstatetransition)
					res[actorstate_] = actorstatetransitions
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ActorStateTransitionShape
	case ActorStateTransitionShape:
		switch fieldname {
		// insertion point for per direct association field
		case "ActorStateTransition":
			res := make(map[*ActorStateTransition][]*ActorStateTransitionShape)
			for actorstatetransitionshape := range stage.ActorStateTransitionShapes {
				if actorstatetransitionshape.ActorStateTransition != nil {
					actorstatetransition_ := actorstatetransitionshape.ActorStateTransition
					var actorstatetransitionshapes []*ActorStateTransitionShape
					_, ok := res[actorstatetransition_]
					if ok {
						actorstatetransitionshapes = res[actorstatetransition_]
					} else {
						actorstatetransitionshapes = make([]*ActorStateTransitionShape, 0)
					}
					actorstatetransitionshapes = append(actorstatetransitionshapes, actorstatetransitionshape)
					res[actorstatetransition_] = actorstatetransitionshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Start":
			res := make(map[*ActorStateShape][]*ActorStateTransitionShape)
			for actorstatetransitionshape := range stage.ActorStateTransitionShapes {
				if actorstatetransitionshape.Start != nil {
					actorstateshape_ := actorstatetransitionshape.Start
					var actorstatetransitionshapes []*ActorStateTransitionShape
					_, ok := res[actorstateshape_]
					if ok {
						actorstatetransitionshapes = res[actorstateshape_]
					} else {
						actorstatetransitionshapes = make([]*ActorStateTransitionShape, 0)
					}
					actorstatetransitionshapes = append(actorstatetransitionshapes, actorstatetransitionshape)
					res[actorstateshape_] = actorstatetransitionshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "End":
			res := make(map[*ActorStateShape][]*ActorStateTransitionShape)
			for actorstatetransitionshape := range stage.ActorStateTransitionShapes {
				if actorstatetransitionshape.End != nil {
					actorstateshape_ := actorstatetransitionshape.End
					var actorstatetransitionshapes []*ActorStateTransitionShape
					_, ok := res[actorstateshape_]
					if ok {
						actorstatetransitionshapes = res[actorstateshape_]
					} else {
						actorstatetransitionshapes = make([]*ActorStateTransitionShape, 0)
					}
					actorstatetransitionshapes = append(actorstatetransitionshapes, actorstatetransitionshape)
					res[actorstateshape_] = actorstatetransitionshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Analysis
	case Analysis:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ControlPointShape
	case ControlPointShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Diagram
	case Diagram:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Document
	case Document:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DocumentUse
	case DocumentUse:
		switch fieldname {
		// insertion point for per direct association field
		case "Document":
			res := make(map[*Document][]*DocumentUse)
			for documentuse := range stage.DocumentUses {
				if documentuse.Document != nil {
					document_ := documentuse.Document
					var documentuses []*DocumentUse
					_, ok := res[document_]
					if ok {
						documentuses = res[document_]
					} else {
						documentuses = make([]*DocumentUse, 0)
					}
					documentuses = append(documentuses, documentuse)
					res[document_] = documentuses
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of EvolutionDirection
	case EvolutionDirection:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EvolutionDirectionShape
	case EvolutionDirectionShape:
		switch fieldname {
		// insertion point for per direct association field
		case "EvolutionDirection":
			res := make(map[*EvolutionDirection][]*EvolutionDirectionShape)
			for evolutiondirectionshape := range stage.EvolutionDirectionShapes {
				if evolutiondirectionshape.EvolutionDirection != nil {
					evolutiondirection_ := evolutiondirectionshape.EvolutionDirection
					var evolutiondirectionshapes []*EvolutionDirectionShape
					_, ok := res[evolutiondirection_]
					if ok {
						evolutiondirectionshapes = res[evolutiondirection_]
					} else {
						evolutiondirectionshapes = make([]*EvolutionDirectionShape, 0)
					}
					evolutiondirectionshapes = append(evolutiondirectionshapes, evolutiondirectionshape)
					res[evolutiondirection_] = evolutiondirectionshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Foo
	case Foo:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GeoObject
	case GeoObject:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GeoObjectUse
	case GeoObjectUse:
		switch fieldname {
		// insertion point for per direct association field
		case "GeoObject":
			res := make(map[*GeoObject][]*GeoObjectUse)
			for geoobjectuse := range stage.GeoObjectUses {
				if geoobjectuse.GeoObject != nil {
					geoobject_ := geoobjectuse.GeoObject
					var geoobjectuses []*GeoObjectUse
					_, ok := res[geoobject_]
					if ok {
						geoobjectuses = res[geoobject_]
					} else {
						geoobjectuses = make([]*GeoObjectUse, 0)
					}
					geoobjectuses = append(geoobjectuses, geoobjectuse)
					res[geoobject_] = geoobjectuses
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Group
	case Group:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GroupUse
	case GroupUse:
		switch fieldname {
		// insertion point for per direct association field
		case "Group":
			res := make(map[*Group][]*GroupUse)
			for groupuse := range stage.GroupUses {
				if groupuse.Group != nil {
					group_ := groupuse.Group
					var groupuses []*GroupUse
					_, ok := res[group_]
					if ok {
						groupuses = res[group_]
					} else {
						groupuses = make([]*GroupUse, 0)
					}
					groupuses = append(groupuses, groupuse)
					res[group_] = groupuses
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Library
	case Library:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MapObject
	case MapObject:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MapObjectUse
	case MapObjectUse:
		switch fieldname {
		// insertion point for per direct association field
		case "Map":
			res := make(map[*MapObject][]*MapObjectUse)
			for mapobjectuse := range stage.MapObjectUses {
				if mapobjectuse.Map != nil {
					mapobject_ := mapobjectuse.Map
					var mapobjectuses []*MapObjectUse
					_, ok := res[mapobject_]
					if ok {
						mapobjectuses = res[mapobject_]
					} else {
						mapobjectuses = make([]*MapObjectUse, 0)
					}
					mapobjectuses = append(mapobjectuses, mapobjectuse)
					res[mapobject_] = mapobjectuses
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Parameter
	case Parameter:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ParameterCategory
	case ParameterCategory:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ParameterCategoryUse
	case ParameterCategoryUse:
		switch fieldname {
		// insertion point for per direct association field
		case "ParameterCategory":
			res := make(map[*ParameterCategory][]*ParameterCategoryUse)
			for parametercategoryuse := range stage.ParameterCategoryUses {
				if parametercategoryuse.ParameterCategory != nil {
					parametercategory_ := parametercategoryuse.ParameterCategory
					var parametercategoryuses []*ParameterCategoryUse
					_, ok := res[parametercategory_]
					if ok {
						parametercategoryuses = res[parametercategory_]
					} else {
						parametercategoryuses = make([]*ParameterCategoryUse, 0)
					}
					parametercategoryuses = append(parametercategoryuses, parametercategoryuse)
					res[parametercategory_] = parametercategoryuses
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ParameterShape
	case ParameterShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Parameter":
			res := make(map[*Parameter][]*ParameterShape)
			for parametershape := range stage.ParameterShapes {
				if parametershape.Parameter != nil {
					parameter_ := parametershape.Parameter
					var parametershapes []*ParameterShape
					_, ok := res[parameter_]
					if ok {
						parametershapes = res[parameter_]
					} else {
						parametershapes = make([]*ParameterShape, 0)
					}
					parametershapes = append(parametershapes, parametershape)
					res[parameter_] = parametershapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ParametersAggregate
	case ParametersAggregate:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ParametersAggregateShape
	case ParametersAggregateShape:
		switch fieldname {
		// insertion point for per direct association field
		case "ScenarioParameter":
			res := make(map[*ParametersAggregate][]*ParametersAggregateShape)
			for parametersaggregateshape := range stage.ParametersAggregateShapes {
				if parametersaggregateshape.ScenarioParameter != nil {
					parametersaggregate_ := parametersaggregateshape.ScenarioParameter
					var parametersaggregateshapes []*ParametersAggregateShape
					_, ok := res[parametersaggregate_]
					if ok {
						parametersaggregateshapes = res[parametersaggregate_]
					} else {
						parametersaggregateshapes = make([]*ParametersAggregateShape, 0)
					}
					parametersaggregateshapes = append(parametersaggregateshapes, parametersaggregateshape)
					res[parametersaggregate_] = parametersaggregateshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Position
	case Position:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Repository
	case Repository:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Scenario
	case Scenario:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of User
	case User:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of UserUse
	case UserUse:
		switch fieldname {
		// insertion point for per direct association field
		case "User":
			res := make(map[*User][]*UserUse)
			for useruse := range stage.UserUses {
				if useruse.User != nil {
					user_ := useruse.User
					var useruses []*UserUse
					_, ok := res[user_]
					if ok {
						useruses = res[user_]
					} else {
						useruses = make([]*UserUse, 0)
					}
					useruses = append(useruses, useruse)
					res[user_] = useruses
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Workspace
	case Workspace:
		switch fieldname {
		// insertion point for per direct association field
		case "SelectedDiagram":
			res := make(map[*Diagram][]*Workspace)
			for workspace := range stage.Workspaces {
				if workspace.SelectedDiagram != nil {
					diagram_ := workspace.SelectedDiagram
					var workspaces []*Workspace
					_, ok := res[diagram_]
					if ok {
						workspaces = res[diagram_]
					} else {
						workspaces = make([]*Workspace, 0)
					}
					workspaces = append(workspaces, workspace)
					res[diagram_] = workspaces
				}
			}
			return any(res).(map[*End][]*Start)
		case "Default_EvolutionDirectionShape":
			res := make(map[*EvolutionDirectionShape][]*Workspace)
			for workspace := range stage.Workspaces {
				if workspace.Default_EvolutionDirectionShape != nil {
					evolutiondirectionshape_ := workspace.Default_EvolutionDirectionShape
					var workspaces []*Workspace
					_, ok := res[evolutiondirectionshape_]
					if ok {
						workspaces = res[evolutiondirectionshape_]
					} else {
						workspaces = make([]*Workspace, 0)
					}
					workspaces = append(workspaces, workspace)
					res[evolutiondirectionshape_] = workspaces
				}
			}
			return any(res).(map[*End][]*Start)
		case "Default_ParameterShape":
			res := make(map[*ParameterShape][]*Workspace)
			for workspace := range stage.Workspaces {
				if workspace.Default_ParameterShape != nil {
					parametershape_ := workspace.Default_ParameterShape
					var workspaces []*Workspace
					_, ok := res[parametershape_]
					if ok {
						workspaces = res[parametershape_]
					} else {
						workspaces = make([]*Workspace, 0)
					}
					workspaces = append(workspaces, workspace)
					res[parametershape_] = workspaces
				}
			}
			return any(res).(map[*End][]*Start)
		case "Default_ScenarioParameterShape":
			res := make(map[*ParametersAggregateShape][]*Workspace)
			for workspace := range stage.Workspaces {
				if workspace.Default_ScenarioParameterShape != nil {
					parametersaggregateshape_ := workspace.Default_ScenarioParameterShape
					var workspaces []*Workspace
					_, ok := res[parametersaggregateshape_]
					if ok {
						workspaces = res[parametersaggregateshape_]
					} else {
						workspaces = make([]*Workspace, 0)
					}
					workspaces = append(workspaces, workspace)
					res[parametersaggregateshape_] = workspaces
				}
			}
			return any(res).(map[*End][]*Start)
		case "Default_ActorStateShape":
			res := make(map[*ActorStateShape][]*Workspace)
			for workspace := range stage.Workspaces {
				if workspace.Default_ActorStateShape != nil {
					actorstateshape_ := workspace.Default_ActorStateShape
					var workspaces []*Workspace
					_, ok := res[actorstateshape_]
					if ok {
						workspaces = res[actorstateshape_]
					} else {
						workspaces = make([]*Workspace, 0)
					}
					workspaces = append(workspaces, workspace)
					res[actorstateshape_] = workspaces
				}
			}
			return any(res).(map[*End][]*Start)
		case "Default_ActorStateTransitionShape":
			res := make(map[*ActorStateTransitionShape][]*Workspace)
			for workspace := range stage.Workspaces {
				if workspace.Default_ActorStateTransitionShape != nil {
					actorstatetransitionshape_ := workspace.Default_ActorStateTransitionShape
					var workspaces []*Workspace
					_, ok := res[actorstatetransitionshape_]
					if ok {
						workspaces = res[actorstatetransitionshape_]
					} else {
						workspaces = make([]*Workspace, 0)
					}
					workspaces = append(workspaces, workspace)
					res[actorstatetransitionshape_] = workspaces
				}
			}
			return any(res).(map[*End][]*Start)
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap is the Stage method for backtrack navigation of slice-of-pointers associations.
func (stage *Stage) GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string) map[*End][]*Start {
	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of ActorState
	case ActorState:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ActorStateShape
	case ActorStateShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ActorStateTransition
	case ActorStateTransition:
		switch fieldname {
		// insertion point for per direct association field
		case "Justifications":
			res := make(map[*Parameter][]*ActorStateTransition)
			for actorstatetransition := range stage.ActorStateTransitions {
				for _, parameter_ := range actorstatetransition.Justifications {
					res[parameter_] = append(res[parameter_], actorstatetransition)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ActorStateTransitionShape
	case ActorStateTransitionShape:
		switch fieldname {
		// insertion point for per direct association field
		case "ControlPointShapes":
			res := make(map[*ControlPointShape][]*ActorStateTransitionShape)
			for actorstatetransitionshape := range stage.ActorStateTransitionShapes {
				for _, controlpointshape_ := range actorstatetransitionshape.ControlPointShapes {
					res[controlpointshape_] = append(res[controlpointshape_], actorstatetransitionshape)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Analysis
	case Analysis:
		switch fieldname {
		// insertion point for per direct association field
		case "Scenarios":
			res := make(map[*Scenario][]*Analysis)
			for analysis := range stage.Analysiss {
				for _, scenario_ := range analysis.Scenarios {
					res[scenario_] = append(res[scenario_], analysis)
				}
			}
			return any(res).(map[*End][]*Start)
		case "GroupUse":
			res := make(map[*GroupUse][]*Analysis)
			for analysis := range stage.Analysiss {
				for _, groupuse_ := range analysis.GroupUse {
					res[groupuse_] = append(res[groupuse_], analysis)
				}
			}
			return any(res).(map[*End][]*Start)
		case "GeoObjectUse":
			res := make(map[*GeoObjectUse][]*Analysis)
			for analysis := range stage.Analysiss {
				for _, geoobjectuse_ := range analysis.GeoObjectUse {
					res[geoobjectuse_] = append(res[geoobjectuse_], analysis)
				}
			}
			return any(res).(map[*End][]*Start)
		case "MapUse":
			res := make(map[*MapObjectUse][]*Analysis)
			for analysis := range stage.Analysiss {
				for _, mapobjectuse_ := range analysis.MapUse {
					res[mapobjectuse_] = append(res[mapobjectuse_], analysis)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ControlPointShape
	case ControlPointShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Diagram
	case Diagram:
		switch fieldname {
		// insertion point for per direct association field
		case "EvolutionDirectionShapes":
			res := make(map[*EvolutionDirectionShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, evolutiondirectionshape_ := range diagram.EvolutionDirectionShapes {
					res[evolutiondirectionshape_] = append(res[evolutiondirectionshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "EvolutionDirectionsWhoseNodeIsExpanded":
			res := make(map[*EvolutionDirection][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, evolutiondirection_ := range diagram.EvolutionDirectionsWhoseNodeIsExpanded {
					res[evolutiondirection_] = append(res[evolutiondirection_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ActorStateShapes":
			res := make(map[*ActorStateShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, actorstateshape_ := range diagram.ActorStateShapes {
					res[actorstateshape_] = append(res[actorstateshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ActorStatesWhoseNodeIsExpanded":
			res := make(map[*ActorState][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, actorstate_ := range diagram.ActorStatesWhoseNodeIsExpanded {
					res[actorstate_] = append(res[actorstate_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ParameterShapes":
			res := make(map[*ParameterShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, parametershape_ := range diagram.ParameterShapes {
					res[parametershape_] = append(res[parametershape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ParametersWhoseNodeIsExpanded":
			res := make(map[*Parameter][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, parameter_ := range diagram.ParametersWhoseNodeIsExpanded {
					res[parameter_] = append(res[parameter_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ScenarioParameterShapes":
			res := make(map[*ParametersAggregateShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, parametersaggregateshape_ := range diagram.ScenarioParameterShapes {
					res[parametersaggregateshape_] = append(res[parametersaggregateshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ParametersAggregatesWhoseNodeIsExpanded":
			res := make(map[*ParametersAggregate][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, parametersaggregate_ := range diagram.ParametersAggregatesWhoseNodeIsExpanded {
					res[parametersaggregate_] = append(res[parametersaggregate_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ActorStateTransitionShapes":
			res := make(map[*ActorStateTransitionShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, actorstatetransitionshape_ := range diagram.ActorStateTransitionShapes {
					res[actorstatetransitionshape_] = append(res[actorstatetransitionshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ActorStateTransitionsWhoseNodeIsExpanded":
			res := make(map[*ActorStateTransition][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, actorstatetransition_ := range diagram.ActorStateTransitionsWhoseNodeIsExpanded {
					res[actorstatetransition_] = append(res[actorstatetransition_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Document
	case Document:
		switch fieldname {
		// insertion point for per direct association field
		case "GeoObjectUse":
			res := make(map[*GeoObjectUse][]*Document)
			for document := range stage.Documents {
				for _, geoobjectuse_ := range document.GeoObjectUse {
					res[geoobjectuse_] = append(res[geoobjectuse_], document)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DocumentUse
	case DocumentUse:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EvolutionDirection
	case EvolutionDirection:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EvolutionDirectionShape
	case EvolutionDirectionShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Foo
	case Foo:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GeoObject
	case GeoObject:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GeoObjectUse
	case GeoObjectUse:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Group
	case Group:
		switch fieldname {
		// insertion point for per direct association field
		case "UserUse":
			res := make(map[*UserUse][]*Group)
			for group := range stage.Groups {
				for _, useruse_ := range group.UserUse {
					res[useruse_] = append(res[useruse_], group)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GroupUse
	case GroupUse:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Library
	case Library:
		switch fieldname {
		// insertion point for per direct association field
		case "Analyses":
			res := make(map[*Analysis][]*Library)
			for library := range stage.Librarys {
				for _, analysis_ := range library.Analyses {
					res[analysis_] = append(res[analysis_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SubLibraries":
			res := make(map[*Library][]*Library)
			for library := range stage.Librarys {
				for _, library_ := range library.SubLibraries {
					res[library_] = append(res[library_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SubLibrariesWhoseNodeIsExpanded":
			res := make(map[*Library][]*Library)
			for library := range stage.Librarys {
				for _, library_ := range library.SubLibrariesWhoseNodeIsExpanded {
					res[library_] = append(res[library_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MapObject
	case MapObject:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MapObjectUse
	case MapObjectUse:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Parameter
	case Parameter:
		switch fieldname {
		// insertion point for per direct association field
		case "GroupUse":
			res := make(map[*GroupUse][]*Parameter)
			for parameter := range stage.Parameters {
				for _, groupuse_ := range parameter.GroupUse {
					res[groupuse_] = append(res[groupuse_], parameter)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DocumentUse":
			res := make(map[*DocumentUse][]*Parameter)
			for parameter := range stage.Parameters {
				for _, documentuse_ := range parameter.DocumentUse {
					res[documentuse_] = append(res[documentuse_], parameter)
				}
			}
			return any(res).(map[*End][]*Start)
		case "GeoObjectUse":
			res := make(map[*GeoObjectUse][]*Parameter)
			for parameter := range stage.Parameters {
				for _, geoobjectuse_ := range parameter.GeoObjectUse {
					res[geoobjectuse_] = append(res[geoobjectuse_], parameter)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ParameterCategory
	case ParameterCategory:
		switch fieldname {
		// insertion point for per direct association field
		case "ParameterUse":
			res := make(map[*ParameterShape][]*ParameterCategory)
			for parametercategory := range stage.ParameterCategorys {
				for _, parametershape_ := range parametercategory.ParameterUse {
					res[parametershape_] = append(res[parametershape_], parametercategory)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ParameterCategoryUse
	case ParameterCategoryUse:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ParameterShape
	case ParameterShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ParametersAggregate
	case ParametersAggregate:
		switch fieldname {
		// insertion point for per direct association field
		case "Parameters":
			res := make(map[*Parameter][]*ParametersAggregate)
			for parametersaggregate := range stage.ParametersAggregates {
				for _, parameter_ := range parametersaggregate.Parameters {
					res[parameter_] = append(res[parameter_], parametersaggregate)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ParametersAggregateShape
	case ParametersAggregateShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Position
	case Position:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Repository
	case Repository:
		switch fieldname {
		// insertion point for per direct association field
		case "ParameterUse":
			res := make(map[*ParameterShape][]*Repository)
			for repository := range stage.Repositorys {
				for _, parametershape_ := range repository.ParameterUse {
					res[parametershape_] = append(res[parametershape_], repository)
				}
			}
			return any(res).(map[*End][]*Start)
		case "GroupUse":
			res := make(map[*GroupUse][]*Repository)
			for repository := range stage.Repositorys {
				for _, groupuse_ := range repository.GroupUse {
					res[groupuse_] = append(res[groupuse_], repository)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Scenario
	case Scenario:
		switch fieldname {
		// insertion point for per direct association field
		case "Diagrams":
			res := make(map[*Diagram][]*Scenario)
			for scenario := range stage.Scenarios {
				for _, diagram_ := range scenario.Diagrams {
					res[diagram_] = append(res[diagram_], scenario)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ActorStates":
			res := make(map[*ActorState][]*Scenario)
			for scenario := range stage.Scenarios {
				for _, actorstate_ := range scenario.ActorStates {
					res[actorstate_] = append(res[actorstate_], scenario)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ActorStateTransitions":
			res := make(map[*ActorStateTransition][]*Scenario)
			for scenario := range stage.Scenarios {
				for _, actorstatetransition_ := range scenario.ActorStateTransitions {
					res[actorstatetransition_] = append(res[actorstatetransition_], scenario)
				}
			}
			return any(res).(map[*End][]*Start)
		case "EvolutionDirections":
			res := make(map[*EvolutionDirection][]*Scenario)
			for scenario := range stage.Scenarios {
				for _, evolutiondirection_ := range scenario.EvolutionDirections {
					res[evolutiondirection_] = append(res[evolutiondirection_], scenario)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Parameters":
			res := make(map[*Parameter][]*Scenario)
			for scenario := range stage.Scenarios {
				for _, parameter_ := range scenario.Parameters {
					res[parameter_] = append(res[parameter_], scenario)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ParametersAggretates":
			res := make(map[*ParametersAggregate][]*Scenario)
			for scenario := range stage.Scenarios {
				for _, parametersaggregate_ := range scenario.ParametersAggretates {
					res[parametersaggregate_] = append(res[parametersaggregate_], scenario)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of User
	case User:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of UserUse
	case UserUse:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Workspace
	case Workspace:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GongNewInstance creates a new instance of the Gongstruct
func GongNewInstance[Type GongstructPtr]() (res Type) {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic new instance
	case *ActorState:
		res = any(new(ActorState)).(Type)
	case *ActorStateShape:
		res = any(new(ActorStateShape)).(Type)
	case *ActorStateTransition:
		res = any(new(ActorStateTransition)).(Type)
	case *ActorStateTransitionShape:
		res = any(new(ActorStateTransitionShape)).(Type)
	case *Analysis:
		res = any(new(Analysis)).(Type)
	case *ControlPointShape:
		res = any(new(ControlPointShape)).(Type)
	case *Diagram:
		res = any(new(Diagram)).(Type)
	case *Document:
		res = any(new(Document)).(Type)
	case *DocumentUse:
		res = any(new(DocumentUse)).(Type)
	case *EvolutionDirection:
		res = any(new(EvolutionDirection)).(Type)
	case *EvolutionDirectionShape:
		res = any(new(EvolutionDirectionShape)).(Type)
	case *Foo:
		res = any(new(Foo)).(Type)
	case *GeoObject:
		res = any(new(GeoObject)).(Type)
	case *GeoObjectUse:
		res = any(new(GeoObjectUse)).(Type)
	case *Group:
		res = any(new(Group)).(Type)
	case *GroupUse:
		res = any(new(GroupUse)).(Type)
	case *Library:
		res = any(new(Library)).(Type)
	case *MapObject:
		res = any(new(MapObject)).(Type)
	case *MapObjectUse:
		res = any(new(MapObjectUse)).(Type)
	case *Parameter:
		res = any(new(Parameter)).(Type)
	case *ParameterCategory:
		res = any(new(ParameterCategory)).(Type)
	case *ParameterCategoryUse:
		res = any(new(ParameterCategoryUse)).(Type)
	case *ParameterShape:
		res = any(new(ParameterShape)).(Type)
	case *ParametersAggregate:
		res = any(new(ParametersAggregate)).(Type)
	case *ParametersAggregateShape:
		res = any(new(ParametersAggregateShape)).(Type)
	case *Position:
		res = any(new(Position)).(Type)
	case *Repository:
		res = any(new(Repository)).(Type)
	case *Scenario:
		res = any(new(Scenario)).(Type)
	case *User:
		res = any(new(User)).(Type)
	case *UserUse:
		res = any(new(UserUse)).(Type)
	case *Workspace:
		res = any(new(Workspace)).(Type)
	}
	return res
}

func NewInstance[Type GongstructPtr]() (res Type) {
	return GongNewInstance[Type]()
}

func (stage *Stage) GongNewInstance[Type GongstructPtr]() (res Type) {
	res = GongNewInstance[Type]()
	var zero Type
	if res != zero {
		res.StageVoid(stage)
	}
	return res
}

func (stage *Stage) NewInstance[Type GongstructPtr]() (res Type) {
	return stage.GongNewInstance[Type]()
}

// GongGetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GongGetPointerToGongstructName[Type GongstructIF]() (res string) {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *ActorState:
		res = "ActorState"
	case *ActorStateShape:
		res = "ActorStateShape"
	case *ActorStateTransition:
		res = "ActorStateTransition"
	case *ActorStateTransitionShape:
		res = "ActorStateTransitionShape"
	case *Analysis:
		res = "Analysis"
	case *ControlPointShape:
		res = "ControlPointShape"
	case *Diagram:
		res = "Diagram"
	case *Document:
		res = "Document"
	case *DocumentUse:
		res = "DocumentUse"
	case *EvolutionDirection:
		res = "EvolutionDirection"
	case *EvolutionDirectionShape:
		res = "EvolutionDirectionShape"
	case *Foo:
		res = "Foo"
	case *GeoObject:
		res = "GeoObject"
	case *GeoObjectUse:
		res = "GeoObjectUse"
	case *Group:
		res = "Group"
	case *GroupUse:
		res = "GroupUse"
	case *Library:
		res = "Library"
	case *MapObject:
		res = "MapObject"
	case *MapObjectUse:
		res = "MapObjectUse"
	case *Parameter:
		res = "Parameter"
	case *ParameterCategory:
		res = "ParameterCategory"
	case *ParameterCategoryUse:
		res = "ParameterCategoryUse"
	case *ParameterShape:
		res = "ParameterShape"
	case *ParametersAggregate:
		res = "ParametersAggregate"
	case *ParametersAggregateShape:
		res = "ParametersAggregateShape"
	case *Position:
		res = "Position"
	case *Repository:
		res = "Repository"
	case *Scenario:
		res = "Scenario"
	case *User:
		res = "User"
	case *UserUse:
		res = "UserUse"
	case *Workspace:
		res = "Workspace"
	}
	return res
}

func GetPointerToGongstructName[Type GongstructIF]() (res string) {
	return GongGetPointerToGongstructName[Type]()
}

type GongReverseField struct {
	GongstructName string
	Fieldname      string
}

type ReverseField = GongReverseField

func GongGetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	res = make([]GongReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case *ActorState:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ActorStatesWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Scenario"
		rf.Fieldname = "ActorStates"
		res = append(res, rf)
	case *ActorStateShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ActorStateShapes"
		res = append(res, rf)
	case *ActorStateTransition:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ActorStateTransitionsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Scenario"
		rf.Fieldname = "ActorStateTransitions"
		res = append(res, rf)
	case *ActorStateTransitionShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ActorStateTransitionShapes"
		res = append(res, rf)
	case *Analysis:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Library"
		rf.Fieldname = "Analyses"
		res = append(res, rf)
	case *ControlPointShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "ActorStateTransitionShape"
		rf.Fieldname = "ControlPointShapes"
		res = append(res, rf)
	case *Diagram:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Scenario"
		rf.Fieldname = "Diagrams"
		res = append(res, rf)
	case *Document:
		var rf ReverseField
		_ = rf
	case *DocumentUse:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Parameter"
		rf.Fieldname = "DocumentUse"
		res = append(res, rf)
	case *EvolutionDirection:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "EvolutionDirectionsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Scenario"
		rf.Fieldname = "EvolutionDirections"
		res = append(res, rf)
	case *EvolutionDirectionShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "EvolutionDirectionShapes"
		res = append(res, rf)
	case *Foo:
		var rf ReverseField
		_ = rf
	case *GeoObject:
		var rf ReverseField
		_ = rf
	case *GeoObjectUse:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Analysis"
		rf.Fieldname = "GeoObjectUse"
		res = append(res, rf)
		rf.GongstructName = "Document"
		rf.Fieldname = "GeoObjectUse"
		res = append(res, rf)
		rf.GongstructName = "Parameter"
		rf.Fieldname = "GeoObjectUse"
		res = append(res, rf)
	case *Group:
		var rf ReverseField
		_ = rf
	case *GroupUse:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Analysis"
		rf.Fieldname = "GroupUse"
		res = append(res, rf)
		rf.GongstructName = "Parameter"
		rf.Fieldname = "GroupUse"
		res = append(res, rf)
		rf.GongstructName = "Repository"
		rf.Fieldname = "GroupUse"
		res = append(res, rf)
	case *Library:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Library"
		rf.Fieldname = "SubLibraries"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "SubLibrariesWhoseNodeIsExpanded"
		res = append(res, rf)
	case *MapObject:
		var rf ReverseField
		_ = rf
	case *MapObjectUse:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Analysis"
		rf.Fieldname = "MapUse"
		res = append(res, rf)
	case *Parameter:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "ActorStateTransition"
		rf.Fieldname = "Justifications"
		res = append(res, rf)
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ParametersWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "ParametersAggregate"
		rf.Fieldname = "Parameters"
		res = append(res, rf)
		rf.GongstructName = "Scenario"
		rf.Fieldname = "Parameters"
		res = append(res, rf)
	case *ParameterCategory:
		var rf ReverseField
		_ = rf
	case *ParameterCategoryUse:
		var rf ReverseField
		_ = rf
	case *ParameterShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ParameterShapes"
		res = append(res, rf)
		rf.GongstructName = "ParameterCategory"
		rf.Fieldname = "ParameterUse"
		res = append(res, rf)
		rf.GongstructName = "Repository"
		rf.Fieldname = "ParameterUse"
		res = append(res, rf)
	case *ParametersAggregate:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ParametersAggregatesWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Scenario"
		rf.Fieldname = "ParametersAggretates"
		res = append(res, rf)
	case *ParametersAggregateShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ScenarioParameterShapes"
		res = append(res, rf)
	case *Position:
		var rf ReverseField
		_ = rf
	case *Repository:
		var rf ReverseField
		_ = rf
	case *Scenario:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Analysis"
		rf.Fieldname = "Scenarios"
		res = append(res, rf)
	case *User:
		var rf ReverseField
		_ = rf
	case *UserUse:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Group"
		rf.Fieldname = "UserUse"
		res = append(res, rf)
	case *Workspace:
		var rf ReverseField
		_ = rf
	}
	return
}

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
}

// insertion point for get fields header method
func (actorstate *ActorState) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsWithProbaility",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Probability",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "ProbabilityEnum",
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (actorstateshape *ActorStateShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "ActorState",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ActorState",
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (actorstatetransition *ActorStateTransition) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "StartState",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ActorState",
		},
		{
			Name:                 "EndState",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ActorState",
		},
		{
			Name:                 "Justifications",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Parameter",
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "ActorStateTransition",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ActorStateTransition",
		},
		{
			Name:                 "Start",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ActorStateShape",
		},
		{
			Name:                 "End",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ActorStateShape",
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ControlPointShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlPointShape",
		},
	}
	return
}

func (analysis *Analysis) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Scenarios",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Scenario",
		},
		{
			Name:               "IsScenariosNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "GroupUse",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GroupUse",
		},
		{
			Name:               "IsGroupUseNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "GeoObjectUse",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GeoObjectUse",
		},
		{
			Name:               "IsGeoObjectUseNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "MapUse",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MapObjectUse",
		},
		{
			Name:               "IsMapUseNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (controlpointshape *ControlPointShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "X_Relative",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y_Relative",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsStartShapeTheClosestShape",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (diagram *Diagram) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsChecked",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsShowPrefix",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "EvolutionDirectionShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "EvolutionDirectionShape",
		},
		{
			Name:                 "EvolutionDirectionsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "EvolutionDirection",
		},
		{
			Name:               "IsEvolutionDirectionsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ActorStateShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ActorStateShape",
		},
		{
			Name:                 "ActorStatesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ActorState",
		},
		{
			Name:               "IsActorStatesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ParameterShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ParameterShape",
		},
		{
			Name:                 "ParametersWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Parameter",
		},
		{
			Name:               "IsParametersNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ScenarioParameterShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ParametersAggregateShape",
		},
		{
			Name:                 "ParametersAggregatesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ParametersAggregate",
		},
		{
			Name:               "IsParametersAggregatesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ActorStateTransitionShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ActorStateTransitionShape",
		},
		{
			Name:                 "ActorStateTransitionsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ActorStateTransition",
		},
		{
			Name:               "IsActorStateTransitionsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "AxisOrign_X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "AxisOrign_Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "VerticalAxis_Top_Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "VerticalAxis_Bottom_Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "VerticalAxis_StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "HorizontalAxis_Right_X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Start",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "End",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "NumberOfYearsBetweenTicks",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (document *Document) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "GeoObjectUse",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GeoObjectUse",
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (documentuse *DocumentUse) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Document",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Document",
		},
	}
	return
}

func (evolutiondirection *EvolutionDirection) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "EvolutionDirection",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "EvolutionDirection",
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (foo *Foo) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (geoobject *GeoObject) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (geoobjectuse *GeoObjectUse) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "GeoObject",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "GeoObject",
		},
	}
	return
}

func (group *Group) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "UserUse",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "UserUse",
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (groupuse *GroupUse) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Group",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Group",
		},
	}
	return
}

func (library *Library) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsRootLibrary",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Analyses",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Analysis",
		},
		{
			Name:               "IsAnalysesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SubLibraries",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Library",
		},
		{
			Name:               "IsSubLibrariesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SubLibrariesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Library",
		},
		{
			Name:               "NbPixPerCharacter",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "LogoSVGFile",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpandedTmp",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (mapobject *MapObject) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (mapobjectuse *MapObjectUse) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Map",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MapObject",
		},
	}
	return
}

func (parameter *Parameter) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsResponse",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "Start",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "End",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "Force",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "GroupUse",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GroupUse",
		},
		{
			Name:                 "DocumentUse",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DocumentUse",
		},
		{
			Name:                 "GeoObjectUse",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GeoObjectUse",
		},
		{
			Name:               "Tag",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (parametercategory *ParameterCategory) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "ParameterUse",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ParameterShape",
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (parametercategoryuse *ParameterCategoryUse) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "ParameterCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ParameterCategory",
		},
	}
	return
}

func (parametershape *ParameterShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Parameter",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Parameter",
		},
		{
			Name:                 "Direction",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "DirectionType",
		},
		{
			Name:               "ShapeIsComputedFromModel",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (parametersaggregate *ParametersAggregate) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Tag",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Parameters",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Parameter",
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (parametersaggregateshape *ParametersAggregateShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "ScenarioParameter",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ParametersAggregate",
		},
		{
			Name:                 "Direction",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "DirectionType",
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (position *Position) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Date",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "Ordinate",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (repository *Repository) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "ParameterUse",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ParameterShape",
		},
		{
			Name:                 "GroupUse",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GroupUse",
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (scenario *Scenario) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Diagrams",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Diagram",
		},
		{
			Name:               "IsDiagramsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ActorStates",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ActorState",
		},
		{
			Name:               "IsActorStatesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ActorStateTransitions",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ActorStateTransition",
		},
		{
			Name:               "IsActorStateTransitionsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "EvolutionDirections",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "EvolutionDirection",
		},
		{
			Name:               "IsEvolutionDirectionsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Parameters",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Parameter",
		},
		{
			Name:               "IsParametersNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ParametersAggretates",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ParametersAggregate",
		},
		{
			Name:               "IsParametersAggretatesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (user *User) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (useruse *UserUse) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "User",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "User",
		},
	}
	return
}

func (workspace *Workspace) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "SelectedDiagram",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Diagram",
		},
		{
			Name:                 "Default_EvolutionDirectionShape",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "EvolutionDirectionShape",
		},
		{
			Name:                 "Default_ParameterShape",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ParameterShape",
		},
		{
			Name:                 "Default_ScenarioParameterShape",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ParametersAggregateShape",
		},
		{
			Name:                 "Default_ActorStateShape",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ActorStateShape",
		},
		{
			Name:                 "Default_ActorStateTransitionShape",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ActorStateTransitionShape",
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

// GongGetFieldsFromPointer return the array of the fields
func GongGetFieldsFromPointer[Type GongstructPtr]() (res []GongFieldHeader) {
	var ret Type
	return ret.GongGetFieldHeaders()
}

func GetFieldsFromPointer[Type GongstructPtr]() (res []GongFieldHeader) {
	return GongGetFieldsFromPointer[Type]()
}

type GongFieldValueType string

const (
	GongFieldValueTypeInt             GongFieldValueType = "GongFieldValueTypeInt"
	GongFieldValueTypeIntDuration     GongFieldValueType = "GongFieldValueTypeIntDuration"
	GongFieldValueTypeFloat           GongFieldValueType = "GongFieldValueTypeFloat"
	GongFieldValueTypeBool            GongFieldValueType = "GongFieldValueTypeBool"
	GongFieldValueTypeString          GongFieldValueType = "GongFieldValueTypeString"
	GongFieldValueTypeDate            GongFieldValueType = "GongFieldValueTypeDate"
	GongFieldValueTypeBasicKind       GongFieldValueType = "GongFieldValueTypeBasicKind"
	GongFieldValueTypePointer         GongFieldValueType = "GongFieldValueTypePointer"
	GongFieldValueTypeSliceOfPointers GongFieldValueType = "GongFieldValueTypeSliceOfPointers"
)

type GongFieldValue struct {
	GongFieldValueType
	valueString string
	valueInt    int
	valueFloat  float64
	valueBool   bool

	// in case of a pointer, the ID of the pointed element
	// in case of a slice of pointers, the IDs, separated by semi columbs
	ids string
}

type GongFieldHeader struct {
	Name string
	GongFieldValueType
	TargetGongstructName string
}

func (gongValueField *GongFieldValue) GetValueString() string {
	return gongValueField.valueString
}

func (gongValueField *GongFieldValue) GetValueInt() int {
	return gongValueField.valueInt
}

func (gongValueField *GongFieldValue) GetValueFloat() float64 {
	return gongValueField.valueFloat
}

func (gongValueField *GongFieldValue) GetValueBool() bool {
	return gongValueField.valueBool
}

// insertion point for generic get gongstruct field value
func (actorstate *ActorState) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = actorstate.Name
	case "Description":
		res.valueString = actorstate.Description
	case "IsWithProbaility":
		res.valueString = fmt.Sprintf("%t", actorstate.IsWithProbaility)
		res.valueBool = actorstate.IsWithProbaility
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Probability":
		enum := actorstate.Probability
		res.valueString = enum.ToCodeString()
	case "ComputedPrefix":
		res.valueString = actorstate.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", actorstate.IsExpanded)
		res.valueBool = actorstate.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (actorstateshape *ActorStateShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = actorstateshape.Name
	case "ActorState":
		res.GongFieldValueType = GongFieldValueTypePointer
		if actorstateshape.ActorState != nil {
			res.valueString = actorstateshape.ActorState.Name
			res.ids = actorstateshape.ActorState.GongGetUUID(stage)
		}
	case "X":
		res.valueString = fmt.Sprintf("%f", actorstateshape.X)
		res.valueFloat = actorstateshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", actorstateshape.Y)
		res.valueFloat = actorstateshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", actorstateshape.Width)
		res.valueFloat = actorstateshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", actorstateshape.Height)
		res.valueFloat = actorstateshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", actorstateshape.IsHidden)
		res.valueBool = actorstateshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (actorstatetransition *ActorStateTransition) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = actorstatetransition.Name
	case "StartState":
		res.GongFieldValueType = GongFieldValueTypePointer
		if actorstatetransition.StartState != nil {
			res.valueString = actorstatetransition.StartState.Name
			res.ids = actorstatetransition.StartState.GongGetUUID(stage)
		}
	case "EndState":
		res.GongFieldValueType = GongFieldValueTypePointer
		if actorstatetransition.EndState != nil {
			res.valueString = actorstatetransition.EndState.Name
			res.ids = actorstatetransition.EndState.GongGetUUID(stage)
		}
	case "Justifications":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range actorstatetransition.Justifications {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ComputedPrefix":
		res.valueString = actorstatetransition.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", actorstatetransition.IsExpanded)
		res.valueBool = actorstatetransition.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = actorstatetransitionshape.Name
	case "ActorStateTransition":
		res.GongFieldValueType = GongFieldValueTypePointer
		if actorstatetransitionshape.ActorStateTransition != nil {
			res.valueString = actorstatetransitionshape.ActorStateTransition.Name
			res.ids = actorstatetransitionshape.ActorStateTransition.GongGetUUID(stage)
		}
	case "Start":
		res.GongFieldValueType = GongFieldValueTypePointer
		if actorstatetransitionshape.Start != nil {
			res.valueString = actorstatetransitionshape.Start.Name
			res.ids = actorstatetransitionshape.Start.GongGetUUID(stage)
		}
	case "End":
		res.GongFieldValueType = GongFieldValueTypePointer
		if actorstatetransitionshape.End != nil {
			res.valueString = actorstatetransitionshape.End.Name
			res.ids = actorstatetransitionshape.End.GongGetUUID(stage)
		}
	case "X":
		res.valueString = fmt.Sprintf("%f", actorstatetransitionshape.X)
		res.valueFloat = actorstatetransitionshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", actorstatetransitionshape.Y)
		res.valueFloat = actorstatetransitionshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", actorstatetransitionshape.Width)
		res.valueFloat = actorstatetransitionshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", actorstatetransitionshape.Height)
		res.valueFloat = actorstatetransitionshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", actorstatetransitionshape.IsHidden)
		res.valueBool = actorstatetransitionshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ControlPointShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range actorstatetransitionshape.ControlPointShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (analysis *Analysis) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = analysis.Name
	case "Description":
		res.valueString = analysis.Description
	case "Scenarios":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range analysis.Scenarios {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsScenariosNodeExpanded":
		res.valueString = fmt.Sprintf("%t", analysis.IsScenariosNodeExpanded)
		res.valueBool = analysis.IsScenariosNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "GroupUse":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range analysis.GroupUse {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsGroupUseNodeExpanded":
		res.valueString = fmt.Sprintf("%t", analysis.IsGroupUseNodeExpanded)
		res.valueBool = analysis.IsGroupUseNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "GeoObjectUse":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range analysis.GeoObjectUse {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsGeoObjectUseNodeExpanded":
		res.valueString = fmt.Sprintf("%t", analysis.IsGeoObjectUseNodeExpanded)
		res.valueBool = analysis.IsGeoObjectUseNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "MapUse":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range analysis.MapUse {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsMapUseNodeExpanded":
		res.valueString = fmt.Sprintf("%t", analysis.IsMapUseNodeExpanded)
		res.valueBool = analysis.IsMapUseNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ComputedPrefix":
		res.valueString = analysis.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", analysis.IsExpanded)
		res.valueBool = analysis.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (controlpointshape *ControlPointShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = controlpointshape.Name
	case "X_Relative":
		res.valueString = fmt.Sprintf("%f", controlpointshape.X_Relative)
		res.valueFloat = controlpointshape.X_Relative
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y_Relative":
		res.valueString = fmt.Sprintf("%f", controlpointshape.Y_Relative)
		res.valueFloat = controlpointshape.Y_Relative
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsStartShapeTheClosestShape":
		res.valueString = fmt.Sprintf("%t", controlpointshape.IsStartShapeTheClosestShape)
		res.valueBool = controlpointshape.IsStartShapeTheClosestShape
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (diagram *Diagram) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = diagram.Name
	case "ComputedPrefix":
		res.valueString = diagram.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsExpanded)
		res.valueBool = diagram.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsChecked":
		res.valueString = fmt.Sprintf("%t", diagram.IsChecked)
		res.valueBool = diagram.IsChecked
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsShowPrefix":
		res.valueString = fmt.Sprintf("%t", diagram.IsShowPrefix)
		res.valueBool = diagram.IsShowPrefix
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Description":
		res.valueString = diagram.Description
	case "EvolutionDirectionShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.EvolutionDirectionShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "EvolutionDirectionsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.EvolutionDirectionsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsEvolutionDirectionsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsEvolutionDirectionsNodeExpanded)
		res.valueBool = diagram.IsEvolutionDirectionsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ActorStateShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ActorStateShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ActorStatesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ActorStatesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsActorStatesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsActorStatesNodeExpanded)
		res.valueBool = diagram.IsActorStatesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ParameterShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ParameterShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ParametersWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ParametersWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsParametersNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsParametersNodeExpanded)
		res.valueBool = diagram.IsParametersNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ScenarioParameterShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ScenarioParameterShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ParametersAggregatesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ParametersAggregatesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsParametersAggregatesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsParametersAggregatesNodeExpanded)
		res.valueBool = diagram.IsParametersAggregatesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ActorStateTransitionShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ActorStateTransitionShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ActorStateTransitionsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ActorStateTransitionsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsActorStateTransitionsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsActorStateTransitionsNodeExpanded)
		res.valueBool = diagram.IsActorStateTransitionsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "AxisOrign_X":
		res.valueString = fmt.Sprintf("%f", diagram.AxisOrign_X)
		res.valueFloat = diagram.AxisOrign_X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "AxisOrign_Y":
		res.valueString = fmt.Sprintf("%f", diagram.AxisOrign_Y)
		res.valueFloat = diagram.AxisOrign_Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "VerticalAxis_Top_Y":
		res.valueString = fmt.Sprintf("%f", diagram.VerticalAxis_Top_Y)
		res.valueFloat = diagram.VerticalAxis_Top_Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "VerticalAxis_Bottom_Y":
		res.valueString = fmt.Sprintf("%f", diagram.VerticalAxis_Bottom_Y)
		res.valueFloat = diagram.VerticalAxis_Bottom_Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "VerticalAxis_StrokeWidth":
		res.valueString = fmt.Sprintf("%f", diagram.VerticalAxis_StrokeWidth)
		res.valueFloat = diagram.VerticalAxis_StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "HorizontalAxis_Right_X":
		res.valueString = fmt.Sprintf("%f", diagram.HorizontalAxis_Right_X)
		res.valueFloat = diagram.HorizontalAxis_Right_X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Start":
		res.valueString = diagram.Start.String()
	case "End":
		res.valueString = diagram.End.String()
	case "NumberOfYearsBetweenTicks":
		res.valueString = fmt.Sprintf("%d", diagram.NumberOfYearsBetweenTicks)
		res.valueInt = diagram.NumberOfYearsBetweenTicks
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (document *Document) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = document.Name
	case "GeoObjectUse":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range document.GeoObjectUse {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ComputedPrefix":
		res.valueString = document.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", document.IsExpanded)
		res.valueBool = document.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (documentuse *DocumentUse) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = documentuse.Name
	case "Document":
		res.GongFieldValueType = GongFieldValueTypePointer
		if documentuse.Document != nil {
			res.valueString = documentuse.Document.Name
			res.ids = documentuse.Document.GongGetUUID(stage)
		}
	}
	return
}

func (evolutiondirection *EvolutionDirection) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = evolutiondirection.Name
	case "Description":
		res.valueString = evolutiondirection.Description
	case "ComputedPrefix":
		res.valueString = evolutiondirection.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", evolutiondirection.IsExpanded)
		res.valueBool = evolutiondirection.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = evolutiondirectionshape.Name
	case "EvolutionDirection":
		res.GongFieldValueType = GongFieldValueTypePointer
		if evolutiondirectionshape.EvolutionDirection != nil {
			res.valueString = evolutiondirectionshape.EvolutionDirection.Name
			res.ids = evolutiondirectionshape.EvolutionDirection.GongGetUUID(stage)
		}
	case "X":
		res.valueString = fmt.Sprintf("%f", evolutiondirectionshape.X)
		res.valueFloat = evolutiondirectionshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", evolutiondirectionshape.Y)
		res.valueFloat = evolutiondirectionshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", evolutiondirectionshape.Width)
		res.valueFloat = evolutiondirectionshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", evolutiondirectionshape.Height)
		res.valueFloat = evolutiondirectionshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", evolutiondirectionshape.IsHidden)
		res.valueBool = evolutiondirectionshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (foo *Foo) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = foo.Name
	}
	return
}

func (geoobject *GeoObject) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = geoobject.Name
	case "ComputedPrefix":
		res.valueString = geoobject.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", geoobject.IsExpanded)
		res.valueBool = geoobject.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (geoobjectuse *GeoObjectUse) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = geoobjectuse.Name
	case "GeoObject":
		res.GongFieldValueType = GongFieldValueTypePointer
		if geoobjectuse.GeoObject != nil {
			res.valueString = geoobjectuse.GeoObject.Name
			res.ids = geoobjectuse.GeoObject.GongGetUUID(stage)
		}
	}
	return
}

func (group *Group) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = group.Name
	case "UserUse":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range group.UserUse {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ComputedPrefix":
		res.valueString = group.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", group.IsExpanded)
		res.valueBool = group.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (groupuse *GroupUse) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = groupuse.Name
	case "Group":
		res.GongFieldValueType = GongFieldValueTypePointer
		if groupuse.Group != nil {
			res.valueString = groupuse.Group.Name
			res.ids = groupuse.Group.GongGetUUID(stage)
		}
	}
	return
}

func (library *Library) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = library.Name
	case "Description":
		res.valueString = library.Description
	case "ComputedPrefix":
		res.valueString = library.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsExpanded)
		res.valueBool = library.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsRootLibrary":
		res.valueString = fmt.Sprintf("%t", library.IsRootLibrary)
		res.valueBool = library.IsRootLibrary
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Analyses":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.Analyses {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsAnalysesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsAnalysesNodeExpanded)
		res.valueBool = library.IsAnalysesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SubLibraries":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.SubLibraries {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsSubLibrariesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsSubLibrariesNodeExpanded)
		res.valueBool = library.IsSubLibrariesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SubLibrariesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.SubLibrariesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "NbPixPerCharacter":
		res.valueString = fmt.Sprintf("%f", library.NbPixPerCharacter)
		res.valueFloat = library.NbPixPerCharacter
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LogoSVGFile":
		res.valueString = library.LogoSVGFile
	case "IsExpandedTmp":
		res.valueString = fmt.Sprintf("%t", library.IsExpandedTmp)
		res.valueBool = library.IsExpandedTmp
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (mapobject *MapObject) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = mapobject.Name
	case "ComputedPrefix":
		res.valueString = mapobject.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", mapobject.IsExpanded)
		res.valueBool = mapobject.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (mapobjectuse *MapObjectUse) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = mapobjectuse.Name
	case "Map":
		res.GongFieldValueType = GongFieldValueTypePointer
		if mapobjectuse.Map != nil {
			res.valueString = mapobjectuse.Map.Name
			res.ids = mapobjectuse.Map.GongGetUUID(stage)
		}
	}
	return
}

func (parameter *Parameter) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = parameter.Name
	case "Description":
		res.valueString = parameter.Description
	case "IsResponse":
		res.valueString = fmt.Sprintf("%t", parameter.IsResponse)
		res.valueBool = parameter.IsResponse
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Start":
		res.valueString = parameter.Start.String()
	case "End":
		res.valueString = parameter.End.String()
	case "Force":
		res.valueString = fmt.Sprintf("%f", parameter.Force)
		res.valueFloat = parameter.Force
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "GroupUse":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range parameter.GroupUse {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DocumentUse":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range parameter.DocumentUse {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "GeoObjectUse":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range parameter.GeoObjectUse {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Tag":
		res.valueString = parameter.Tag
	case "ComputedPrefix":
		res.valueString = parameter.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", parameter.IsExpanded)
		res.valueBool = parameter.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (parametercategory *ParameterCategory) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = parametercategory.Name
	case "ParameterUse":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range parametercategory.ParameterUse {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ComputedPrefix":
		res.valueString = parametercategory.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", parametercategory.IsExpanded)
		res.valueBool = parametercategory.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (parametercategoryuse *ParameterCategoryUse) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = parametercategoryuse.Name
	case "ParameterCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parametercategoryuse.ParameterCategory != nil {
			res.valueString = parametercategoryuse.ParameterCategory.Name
			res.ids = parametercategoryuse.ParameterCategory.GongGetUUID(stage)
		}
	}
	return
}

func (parametershape *ParameterShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = parametershape.Name
	case "Parameter":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parametershape.Parameter != nil {
			res.valueString = parametershape.Parameter.Name
			res.ids = parametershape.Parameter.GongGetUUID(stage)
		}
	case "Direction":
		enum := parametershape.Direction
		res.valueString = enum.ToCodeString()
	case "ShapeIsComputedFromModel":
		res.valueString = fmt.Sprintf("%t", parametershape.ShapeIsComputedFromModel)
		res.valueBool = parametershape.ShapeIsComputedFromModel
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", parametershape.X)
		res.valueFloat = parametershape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", parametershape.Y)
		res.valueFloat = parametershape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", parametershape.Width)
		res.valueFloat = parametershape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", parametershape.Height)
		res.valueFloat = parametershape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", parametershape.IsHidden)
		res.valueBool = parametershape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (parametersaggregate *ParametersAggregate) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = parametersaggregate.Name
	case "Tag":
		res.valueString = parametersaggregate.Tag
	case "Description":
		res.valueString = parametersaggregate.Description
	case "Parameters":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range parametersaggregate.Parameters {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ComputedPrefix":
		res.valueString = parametersaggregate.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", parametersaggregate.IsExpanded)
		res.valueBool = parametersaggregate.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (parametersaggregateshape *ParametersAggregateShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = parametersaggregateshape.Name
	case "ScenarioParameter":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parametersaggregateshape.ScenarioParameter != nil {
			res.valueString = parametersaggregateshape.ScenarioParameter.Name
			res.ids = parametersaggregateshape.ScenarioParameter.GongGetUUID(stage)
		}
	case "Direction":
		enum := parametersaggregateshape.Direction
		res.valueString = enum.ToCodeString()
	case "X":
		res.valueString = fmt.Sprintf("%f", parametersaggregateshape.X)
		res.valueFloat = parametersaggregateshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", parametersaggregateshape.Y)
		res.valueFloat = parametersaggregateshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", parametersaggregateshape.Width)
		res.valueFloat = parametersaggregateshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", parametersaggregateshape.Height)
		res.valueFloat = parametersaggregateshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", parametersaggregateshape.IsHidden)
		res.valueBool = parametersaggregateshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (position *Position) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = position.Name
	case "Date":
		res.valueString = position.Date.String()
	case "Ordinate":
		res.valueString = fmt.Sprintf("%f", position.Ordinate)
		res.valueFloat = position.Ordinate
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ComputedPrefix":
		res.valueString = position.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", position.IsExpanded)
		res.valueBool = position.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (repository *Repository) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = repository.Name
	case "ParameterUse":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range repository.ParameterUse {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "GroupUse":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range repository.GroupUse {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ComputedPrefix":
		res.valueString = repository.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", repository.IsExpanded)
		res.valueBool = repository.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (scenario *Scenario) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = scenario.Name
	case "Description":
		res.valueString = scenario.Description
	case "Diagrams":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range scenario.Diagrams {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsDiagramsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", scenario.IsDiagramsNodeExpanded)
		res.valueBool = scenario.IsDiagramsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ActorStates":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range scenario.ActorStates {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsActorStatesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", scenario.IsActorStatesNodeExpanded)
		res.valueBool = scenario.IsActorStatesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ActorStateTransitions":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range scenario.ActorStateTransitions {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsActorStateTransitionsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", scenario.IsActorStateTransitionsNodeExpanded)
		res.valueBool = scenario.IsActorStateTransitionsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "EvolutionDirections":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range scenario.EvolutionDirections {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsEvolutionDirectionsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", scenario.IsEvolutionDirectionsNodeExpanded)
		res.valueBool = scenario.IsEvolutionDirectionsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Parameters":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range scenario.Parameters {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsParametersNodeExpanded":
		res.valueString = fmt.Sprintf("%t", scenario.IsParametersNodeExpanded)
		res.valueBool = scenario.IsParametersNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ParametersAggretates":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range scenario.ParametersAggretates {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsParametersAggretatesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", scenario.IsParametersAggretatesNodeExpanded)
		res.valueBool = scenario.IsParametersAggretatesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ComputedPrefix":
		res.valueString = scenario.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", scenario.IsExpanded)
		res.valueBool = scenario.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (user *User) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = user.Name
	case "ComputedPrefix":
		res.valueString = user.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", user.IsExpanded)
		res.valueBool = user.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (useruse *UserUse) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = useruse.Name
	case "User":
		res.GongFieldValueType = GongFieldValueTypePointer
		if useruse.User != nil {
			res.valueString = useruse.User.Name
			res.ids = useruse.User.GongGetUUID(stage)
		}
	}
	return
}

func (workspace *Workspace) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = workspace.Name
	case "SelectedDiagram":
		res.GongFieldValueType = GongFieldValueTypePointer
		if workspace.SelectedDiagram != nil {
			res.valueString = workspace.SelectedDiagram.Name
			res.ids = workspace.SelectedDiagram.GongGetUUID(stage)
		}
	case "Default_EvolutionDirectionShape":
		res.GongFieldValueType = GongFieldValueTypePointer
		if workspace.Default_EvolutionDirectionShape != nil {
			res.valueString = workspace.Default_EvolutionDirectionShape.Name
			res.ids = workspace.Default_EvolutionDirectionShape.GongGetUUID(stage)
		}
	case "Default_ParameterShape":
		res.GongFieldValueType = GongFieldValueTypePointer
		if workspace.Default_ParameterShape != nil {
			res.valueString = workspace.Default_ParameterShape.Name
			res.ids = workspace.Default_ParameterShape.GongGetUUID(stage)
		}
	case "Default_ScenarioParameterShape":
		res.GongFieldValueType = GongFieldValueTypePointer
		if workspace.Default_ScenarioParameterShape != nil {
			res.valueString = workspace.Default_ScenarioParameterShape.Name
			res.ids = workspace.Default_ScenarioParameterShape.GongGetUUID(stage)
		}
	case "Default_ActorStateShape":
		res.GongFieldValueType = GongFieldValueTypePointer
		if workspace.Default_ActorStateShape != nil {
			res.valueString = workspace.Default_ActorStateShape.Name
			res.ids = workspace.Default_ActorStateShape.GongGetUUID(stage)
		}
	case "Default_ActorStateTransitionShape":
		res.GongFieldValueType = GongFieldValueTypePointer
		if workspace.Default_ActorStateTransitionShape != nil {
			res.valueString = workspace.Default_ActorStateTransitionShape.Name
			res.ids = workspace.Default_ActorStateTransitionShape.GongGetUUID(stage)
		}
	case "ComputedPrefix":
		res.valueString = workspace.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", workspace.IsExpanded)
		res.valueBool = workspace.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (stage *Stage) GetFieldStringValueFromPointer(instance GongstructIF, fieldName string) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	return stage.GetFieldStringValueFromPointer(instance, fieldName)
}

// insertion point for generic get gongstruct name
func (actorstate *ActorState) GongGetGongstructName() string {
	return "ActorState"
}

func (actorstateshape *ActorStateShape) GongGetGongstructName() string {
	return "ActorStateShape"
}

func (actorstatetransition *ActorStateTransition) GongGetGongstructName() string {
	return "ActorStateTransition"
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongGetGongstructName() string {
	return "ActorStateTransitionShape"
}

func (analysis *Analysis) GongGetGongstructName() string {
	return "Analysis"
}

func (controlpointshape *ControlPointShape) GongGetGongstructName() string {
	return "ControlPointShape"
}

func (diagram *Diagram) GongGetGongstructName() string {
	return "Diagram"
}

func (document *Document) GongGetGongstructName() string {
	return "Document"
}

func (documentuse *DocumentUse) GongGetGongstructName() string {
	return "DocumentUse"
}

func (evolutiondirection *EvolutionDirection) GongGetGongstructName() string {
	return "EvolutionDirection"
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongGetGongstructName() string {
	return "EvolutionDirectionShape"
}

func (foo *Foo) GongGetGongstructName() string {
	return "Foo"
}

func (geoobject *GeoObject) GongGetGongstructName() string {
	return "GeoObject"
}

func (geoobjectuse *GeoObjectUse) GongGetGongstructName() string {
	return "GeoObjectUse"
}

func (group *Group) GongGetGongstructName() string {
	return "Group"
}

func (groupuse *GroupUse) GongGetGongstructName() string {
	return "GroupUse"
}

func (library *Library) GongGetGongstructName() string {
	return "Library"
}

func (mapobject *MapObject) GongGetGongstructName() string {
	return "MapObject"
}

func (mapobjectuse *MapObjectUse) GongGetGongstructName() string {
	return "MapObjectUse"
}

func (parameter *Parameter) GongGetGongstructName() string {
	return "Parameter"
}

func (parametercategory *ParameterCategory) GongGetGongstructName() string {
	return "ParameterCategory"
}

func (parametercategoryuse *ParameterCategoryUse) GongGetGongstructName() string {
	return "ParameterCategoryUse"
}

func (parametershape *ParameterShape) GongGetGongstructName() string {
	return "ParameterShape"
}

func (parametersaggregate *ParametersAggregate) GongGetGongstructName() string {
	return "ParametersAggregate"
}

func (parametersaggregateshape *ParametersAggregateShape) GongGetGongstructName() string {
	return "ParametersAggregateShape"
}

func (position *Position) GongGetGongstructName() string {
	return "Position"
}

func (repository *Repository) GongGetGongstructName() string {
	return "Repository"
}

func (scenario *Scenario) GongGetGongstructName() string {
	return "Scenario"
}

func (user *User) GongGetGongstructName() string {
	return "User"
}

func (useruse *UserUse) GongGetGongstructName() string {
	return "UserUse"
}

func (workspace *Workspace) GongGetGongstructName() string {
	return "Workspace"
}

func GongGetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	return GongGetGongstructNameFromPointer(instance)
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	__gong__rebuildMapString(stage.ActorStates, &stage.ActorStates_mapString)

	__gong__rebuildMapString(stage.ActorStateShapes, &stage.ActorStateShapes_mapString)

	__gong__rebuildMapString(stage.ActorStateTransitions, &stage.ActorStateTransitions_mapString)

	__gong__rebuildMapString(stage.ActorStateTransitionShapes, &stage.ActorStateTransitionShapes_mapString)

	__gong__rebuildMapString(stage.Analysiss, &stage.Analysiss_mapString)

	__gong__rebuildMapString(stage.ControlPointShapes, &stage.ControlPointShapes_mapString)

	__gong__rebuildMapString(stage.Diagrams, &stage.Diagrams_mapString)

	__gong__rebuildMapString(stage.Documents, &stage.Documents_mapString)

	__gong__rebuildMapString(stage.DocumentUses, &stage.DocumentUses_mapString)

	__gong__rebuildMapString(stage.EvolutionDirections, &stage.EvolutionDirections_mapString)

	__gong__rebuildMapString(stage.EvolutionDirectionShapes, &stage.EvolutionDirectionShapes_mapString)

	__gong__rebuildMapString(stage.Foos, &stage.Foos_mapString)

	__gong__rebuildMapString(stage.GeoObjects, &stage.GeoObjects_mapString)

	__gong__rebuildMapString(stage.GeoObjectUses, &stage.GeoObjectUses_mapString)

	__gong__rebuildMapString(stage.Groups, &stage.Groups_mapString)

	__gong__rebuildMapString(stage.GroupUses, &stage.GroupUses_mapString)

	__gong__rebuildMapString(stage.Librarys, &stage.Librarys_mapString)

	__gong__rebuildMapString(stage.MapObjects, &stage.MapObjects_mapString)

	__gong__rebuildMapString(stage.MapObjectUses, &stage.MapObjectUses_mapString)

	__gong__rebuildMapString(stage.Parameters, &stage.Parameters_mapString)

	__gong__rebuildMapString(stage.ParameterCategorys, &stage.ParameterCategorys_mapString)

	__gong__rebuildMapString(stage.ParameterCategoryUses, &stage.ParameterCategoryUses_mapString)

	__gong__rebuildMapString(stage.ParameterShapes, &stage.ParameterShapes_mapString)

	__gong__rebuildMapString(stage.ParametersAggregates, &stage.ParametersAggregates_mapString)

	__gong__rebuildMapString(stage.ParametersAggregateShapes, &stage.ParametersAggregateShapes_mapString)

	__gong__rebuildMapString(stage.Positions, &stage.Positions_mapString)

	__gong__rebuildMapString(stage.Repositorys, &stage.Repositorys_mapString)

	__gong__rebuildMapString(stage.Scenarios, &stage.Scenarios_mapString)

	__gong__rebuildMapString(stage.Users, &stage.Users_mapString)

	__gong__rebuildMapString(stage.UserUses, &stage.UserUses_mapString)

	__gong__rebuildMapString(stage.Workspaces, &stage.Workspaces_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
