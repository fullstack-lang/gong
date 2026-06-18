// generated code - do not edit
package models

import (
	"cmp"
	"embed"
	"errors"
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
	"strings"
	"sync"
	"time"

	scenario_go "github.com/fullstack-lang/gong/dsm/scenario/go"
)

// can be used for
//
//	days := __Gong__Abs(int(int(inferedInstance.ComputedDuration.Hours()) / 24))
func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var (
	_ = __Gong__Abs
	_ = strings.Clone("")
)

const (
	ProbeTreeSidebarSuffix           = ":sidebar of the probe"
	ProbeNavigationTreeSidebarSuffix = ":sidebar of the probe, navigation"
	ProbeTableSuffix                 = ":table of the probe"
	ProbeNotificationTableSuffix     = ":notification table of the probe"
	ProbeFormSuffix                  = ":form of the probe"
	ProbeSplitSuffix                 = ":probe of the probe"
)

type GongMarshallingMode string

const (
	// the whole stage is generated at each marshall. This is the default
	GongMarshallingNormal GongMarshallingMode = "GongMarshallingNormal"

	// only the last commit is append to the marshall file
	GongMarshallingAppendCommit GongMarshallingMode = "GongMarshallingAppendCommit"
)

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeNavigationTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeNavigationTreeSidebarSuffix
}

func (stage *Stage) GetProbeFormStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeFormSuffix
}

func (stage *Stage) GetProbeTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTableSuffix
}

func (stage *Stage) GetProbeNotificationTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeNotificationTableSuffix
}

func (stage *Stage) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeSplitSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var (
	errUnkownEnum = errors.New("unkown enum")
	_             = errUnkownEnum
)

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

var _ = __dummy__fmt_variable

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

var _ = __dummy_math_variable

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var (
	__member __void
	_        = __member
)

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
	GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error
	GongGetGongstructName() string
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
	OnAfterActorStateCreateCallback OnAfterCreateInterface[ActorState]
	OnAfterActorStateUpdateCallback OnAfterUpdateInterface[ActorState]
	OnAfterActorStateDeleteCallback OnAfterDeleteInterface[ActorState]
	OnAfterActorStateReadCallback   OnAfterReadInterface[ActorState]

	ActorStateShapes                map[*ActorStateShape]struct{}
	ActorStateShapes_instance       map[*ActorStateShape]*ActorStateShape
	ActorStateShapes_mapString      map[string]*ActorStateShape
	ActorStateShapeOrder            uint
	ActorStateShape_stagedOrder     map[*ActorStateShape]uint
	ActorStateShape_orderStaged     map[uint]*ActorStateShape
	ActorStateShapes_reference      map[*ActorStateShape]*ActorStateShape
	ActorStateShapes_referenceOrder map[*ActorStateShape]uint

	// insertion point for slice of pointers maps
	OnAfterActorStateShapeCreateCallback OnAfterCreateInterface[ActorStateShape]
	OnAfterActorStateShapeUpdateCallback OnAfterUpdateInterface[ActorStateShape]
	OnAfterActorStateShapeDeleteCallback OnAfterDeleteInterface[ActorStateShape]
	OnAfterActorStateShapeReadCallback   OnAfterReadInterface[ActorStateShape]

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

	OnAfterActorStateTransitionCreateCallback OnAfterCreateInterface[ActorStateTransition]
	OnAfterActorStateTransitionUpdateCallback OnAfterUpdateInterface[ActorStateTransition]
	OnAfterActorStateTransitionDeleteCallback OnAfterDeleteInterface[ActorStateTransition]
	OnAfterActorStateTransitionReadCallback   OnAfterReadInterface[ActorStateTransition]

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

	OnAfterActorStateTransitionShapeCreateCallback OnAfterCreateInterface[ActorStateTransitionShape]
	OnAfterActorStateTransitionShapeUpdateCallback OnAfterUpdateInterface[ActorStateTransitionShape]
	OnAfterActorStateTransitionShapeDeleteCallback OnAfterDeleteInterface[ActorStateTransitionShape]
	OnAfterActorStateTransitionShapeReadCallback   OnAfterReadInterface[ActorStateTransitionShape]

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

	OnAfterAnalysisCreateCallback OnAfterCreateInterface[Analysis]
	OnAfterAnalysisUpdateCallback OnAfterUpdateInterface[Analysis]
	OnAfterAnalysisDeleteCallback OnAfterDeleteInterface[Analysis]
	OnAfterAnalysisReadCallback   OnAfterReadInterface[Analysis]

	ControlPointShapes                map[*ControlPointShape]struct{}
	ControlPointShapes_instance       map[*ControlPointShape]*ControlPointShape
	ControlPointShapes_mapString      map[string]*ControlPointShape
	ControlPointShapeOrder            uint
	ControlPointShape_stagedOrder     map[*ControlPointShape]uint
	ControlPointShape_orderStaged     map[uint]*ControlPointShape
	ControlPointShapes_reference      map[*ControlPointShape]*ControlPointShape
	ControlPointShapes_referenceOrder map[*ControlPointShape]uint

	// insertion point for slice of pointers maps
	OnAfterControlPointShapeCreateCallback OnAfterCreateInterface[ControlPointShape]
	OnAfterControlPointShapeUpdateCallback OnAfterUpdateInterface[ControlPointShape]
	OnAfterControlPointShapeDeleteCallback OnAfterDeleteInterface[ControlPointShape]
	OnAfterControlPointShapeReadCallback   OnAfterReadInterface[ControlPointShape]

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

	OnAfterDiagramCreateCallback OnAfterCreateInterface[Diagram]
	OnAfterDiagramUpdateCallback OnAfterUpdateInterface[Diagram]
	OnAfterDiagramDeleteCallback OnAfterDeleteInterface[Diagram]
	OnAfterDiagramReadCallback   OnAfterReadInterface[Diagram]

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

	OnAfterDocumentCreateCallback OnAfterCreateInterface[Document]
	OnAfterDocumentUpdateCallback OnAfterUpdateInterface[Document]
	OnAfterDocumentDeleteCallback OnAfterDeleteInterface[Document]
	OnAfterDocumentReadCallback   OnAfterReadInterface[Document]

	DocumentUses                map[*DocumentUse]struct{}
	DocumentUses_instance       map[*DocumentUse]*DocumentUse
	DocumentUses_mapString      map[string]*DocumentUse
	DocumentUseOrder            uint
	DocumentUse_stagedOrder     map[*DocumentUse]uint
	DocumentUse_orderStaged     map[uint]*DocumentUse
	DocumentUses_reference      map[*DocumentUse]*DocumentUse
	DocumentUses_referenceOrder map[*DocumentUse]uint

	// insertion point for slice of pointers maps
	OnAfterDocumentUseCreateCallback OnAfterCreateInterface[DocumentUse]
	OnAfterDocumentUseUpdateCallback OnAfterUpdateInterface[DocumentUse]
	OnAfterDocumentUseDeleteCallback OnAfterDeleteInterface[DocumentUse]
	OnAfterDocumentUseReadCallback   OnAfterReadInterface[DocumentUse]

	EvolutionDirections                map[*EvolutionDirection]struct{}
	EvolutionDirections_instance       map[*EvolutionDirection]*EvolutionDirection
	EvolutionDirections_mapString      map[string]*EvolutionDirection
	EvolutionDirectionOrder            uint
	EvolutionDirection_stagedOrder     map[*EvolutionDirection]uint
	EvolutionDirection_orderStaged     map[uint]*EvolutionDirection
	EvolutionDirections_reference      map[*EvolutionDirection]*EvolutionDirection
	EvolutionDirections_referenceOrder map[*EvolutionDirection]uint

	// insertion point for slice of pointers maps
	OnAfterEvolutionDirectionCreateCallback OnAfterCreateInterface[EvolutionDirection]
	OnAfterEvolutionDirectionUpdateCallback OnAfterUpdateInterface[EvolutionDirection]
	OnAfterEvolutionDirectionDeleteCallback OnAfterDeleteInterface[EvolutionDirection]
	OnAfterEvolutionDirectionReadCallback   OnAfterReadInterface[EvolutionDirection]

	EvolutionDirectionShapes                map[*EvolutionDirectionShape]struct{}
	EvolutionDirectionShapes_instance       map[*EvolutionDirectionShape]*EvolutionDirectionShape
	EvolutionDirectionShapes_mapString      map[string]*EvolutionDirectionShape
	EvolutionDirectionShapeOrder            uint
	EvolutionDirectionShape_stagedOrder     map[*EvolutionDirectionShape]uint
	EvolutionDirectionShape_orderStaged     map[uint]*EvolutionDirectionShape
	EvolutionDirectionShapes_reference      map[*EvolutionDirectionShape]*EvolutionDirectionShape
	EvolutionDirectionShapes_referenceOrder map[*EvolutionDirectionShape]uint

	// insertion point for slice of pointers maps
	OnAfterEvolutionDirectionShapeCreateCallback OnAfterCreateInterface[EvolutionDirectionShape]
	OnAfterEvolutionDirectionShapeUpdateCallback OnAfterUpdateInterface[EvolutionDirectionShape]
	OnAfterEvolutionDirectionShapeDeleteCallback OnAfterDeleteInterface[EvolutionDirectionShape]
	OnAfterEvolutionDirectionShapeReadCallback   OnAfterReadInterface[EvolutionDirectionShape]

	Foos                map[*Foo]struct{}
	Foos_instance       map[*Foo]*Foo
	Foos_mapString      map[string]*Foo
	FooOrder            uint
	Foo_stagedOrder     map[*Foo]uint
	Foo_orderStaged     map[uint]*Foo
	Foos_reference      map[*Foo]*Foo
	Foos_referenceOrder map[*Foo]uint

	// insertion point for slice of pointers maps
	OnAfterFooCreateCallback OnAfterCreateInterface[Foo]
	OnAfterFooUpdateCallback OnAfterUpdateInterface[Foo]
	OnAfterFooDeleteCallback OnAfterDeleteInterface[Foo]
	OnAfterFooReadCallback   OnAfterReadInterface[Foo]

	GeoObjects                map[*GeoObject]struct{}
	GeoObjects_instance       map[*GeoObject]*GeoObject
	GeoObjects_mapString      map[string]*GeoObject
	GeoObjectOrder            uint
	GeoObject_stagedOrder     map[*GeoObject]uint
	GeoObject_orderStaged     map[uint]*GeoObject
	GeoObjects_reference      map[*GeoObject]*GeoObject
	GeoObjects_referenceOrder map[*GeoObject]uint

	// insertion point for slice of pointers maps
	OnAfterGeoObjectCreateCallback OnAfterCreateInterface[GeoObject]
	OnAfterGeoObjectUpdateCallback OnAfterUpdateInterface[GeoObject]
	OnAfterGeoObjectDeleteCallback OnAfterDeleteInterface[GeoObject]
	OnAfterGeoObjectReadCallback   OnAfterReadInterface[GeoObject]

	GeoObjectUses                map[*GeoObjectUse]struct{}
	GeoObjectUses_instance       map[*GeoObjectUse]*GeoObjectUse
	GeoObjectUses_mapString      map[string]*GeoObjectUse
	GeoObjectUseOrder            uint
	GeoObjectUse_stagedOrder     map[*GeoObjectUse]uint
	GeoObjectUse_orderStaged     map[uint]*GeoObjectUse
	GeoObjectUses_reference      map[*GeoObjectUse]*GeoObjectUse
	GeoObjectUses_referenceOrder map[*GeoObjectUse]uint

	// insertion point for slice of pointers maps
	OnAfterGeoObjectUseCreateCallback OnAfterCreateInterface[GeoObjectUse]
	OnAfterGeoObjectUseUpdateCallback OnAfterUpdateInterface[GeoObjectUse]
	OnAfterGeoObjectUseDeleteCallback OnAfterDeleteInterface[GeoObjectUse]
	OnAfterGeoObjectUseReadCallback   OnAfterReadInterface[GeoObjectUse]

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

	OnAfterGroupCreateCallback OnAfterCreateInterface[Group]
	OnAfterGroupUpdateCallback OnAfterUpdateInterface[Group]
	OnAfterGroupDeleteCallback OnAfterDeleteInterface[Group]
	OnAfterGroupReadCallback   OnAfterReadInterface[Group]

	GroupUses                map[*GroupUse]struct{}
	GroupUses_instance       map[*GroupUse]*GroupUse
	GroupUses_mapString      map[string]*GroupUse
	GroupUseOrder            uint
	GroupUse_stagedOrder     map[*GroupUse]uint
	GroupUse_orderStaged     map[uint]*GroupUse
	GroupUses_reference      map[*GroupUse]*GroupUse
	GroupUses_referenceOrder map[*GroupUse]uint

	// insertion point for slice of pointers maps
	OnAfterGroupUseCreateCallback OnAfterCreateInterface[GroupUse]
	OnAfterGroupUseUpdateCallback OnAfterUpdateInterface[GroupUse]
	OnAfterGroupUseDeleteCallback OnAfterDeleteInterface[GroupUse]
	OnAfterGroupUseReadCallback   OnAfterReadInterface[GroupUse]

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

	OnAfterLibraryCreateCallback OnAfterCreateInterface[Library]
	OnAfterLibraryUpdateCallback OnAfterUpdateInterface[Library]
	OnAfterLibraryDeleteCallback OnAfterDeleteInterface[Library]
	OnAfterLibraryReadCallback   OnAfterReadInterface[Library]

	MapObjects                map[*MapObject]struct{}
	MapObjects_instance       map[*MapObject]*MapObject
	MapObjects_mapString      map[string]*MapObject
	MapObjectOrder            uint
	MapObject_stagedOrder     map[*MapObject]uint
	MapObject_orderStaged     map[uint]*MapObject
	MapObjects_reference      map[*MapObject]*MapObject
	MapObjects_referenceOrder map[*MapObject]uint

	// insertion point for slice of pointers maps
	OnAfterMapObjectCreateCallback OnAfterCreateInterface[MapObject]
	OnAfterMapObjectUpdateCallback OnAfterUpdateInterface[MapObject]
	OnAfterMapObjectDeleteCallback OnAfterDeleteInterface[MapObject]
	OnAfterMapObjectReadCallback   OnAfterReadInterface[MapObject]

	MapObjectUses                map[*MapObjectUse]struct{}
	MapObjectUses_instance       map[*MapObjectUse]*MapObjectUse
	MapObjectUses_mapString      map[string]*MapObjectUse
	MapObjectUseOrder            uint
	MapObjectUse_stagedOrder     map[*MapObjectUse]uint
	MapObjectUse_orderStaged     map[uint]*MapObjectUse
	MapObjectUses_reference      map[*MapObjectUse]*MapObjectUse
	MapObjectUses_referenceOrder map[*MapObjectUse]uint

	// insertion point for slice of pointers maps
	OnAfterMapObjectUseCreateCallback OnAfterCreateInterface[MapObjectUse]
	OnAfterMapObjectUseUpdateCallback OnAfterUpdateInterface[MapObjectUse]
	OnAfterMapObjectUseDeleteCallback OnAfterDeleteInterface[MapObjectUse]
	OnAfterMapObjectUseReadCallback   OnAfterReadInterface[MapObjectUse]

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

	OnAfterParameterCreateCallback OnAfterCreateInterface[Parameter]
	OnAfterParameterUpdateCallback OnAfterUpdateInterface[Parameter]
	OnAfterParameterDeleteCallback OnAfterDeleteInterface[Parameter]
	OnAfterParameterReadCallback   OnAfterReadInterface[Parameter]

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

	OnAfterParameterCategoryCreateCallback OnAfterCreateInterface[ParameterCategory]
	OnAfterParameterCategoryUpdateCallback OnAfterUpdateInterface[ParameterCategory]
	OnAfterParameterCategoryDeleteCallback OnAfterDeleteInterface[ParameterCategory]
	OnAfterParameterCategoryReadCallback   OnAfterReadInterface[ParameterCategory]

	ParameterCategoryUses                map[*ParameterCategoryUse]struct{}
	ParameterCategoryUses_instance       map[*ParameterCategoryUse]*ParameterCategoryUse
	ParameterCategoryUses_mapString      map[string]*ParameterCategoryUse
	ParameterCategoryUseOrder            uint
	ParameterCategoryUse_stagedOrder     map[*ParameterCategoryUse]uint
	ParameterCategoryUse_orderStaged     map[uint]*ParameterCategoryUse
	ParameterCategoryUses_reference      map[*ParameterCategoryUse]*ParameterCategoryUse
	ParameterCategoryUses_referenceOrder map[*ParameterCategoryUse]uint

	// insertion point for slice of pointers maps
	OnAfterParameterCategoryUseCreateCallback OnAfterCreateInterface[ParameterCategoryUse]
	OnAfterParameterCategoryUseUpdateCallback OnAfterUpdateInterface[ParameterCategoryUse]
	OnAfterParameterCategoryUseDeleteCallback OnAfterDeleteInterface[ParameterCategoryUse]
	OnAfterParameterCategoryUseReadCallback   OnAfterReadInterface[ParameterCategoryUse]

	ParameterShapes                map[*ParameterShape]struct{}
	ParameterShapes_instance       map[*ParameterShape]*ParameterShape
	ParameterShapes_mapString      map[string]*ParameterShape
	ParameterShapeOrder            uint
	ParameterShape_stagedOrder     map[*ParameterShape]uint
	ParameterShape_orderStaged     map[uint]*ParameterShape
	ParameterShapes_reference      map[*ParameterShape]*ParameterShape
	ParameterShapes_referenceOrder map[*ParameterShape]uint

	// insertion point for slice of pointers maps
	OnAfterParameterShapeCreateCallback OnAfterCreateInterface[ParameterShape]
	OnAfterParameterShapeUpdateCallback OnAfterUpdateInterface[ParameterShape]
	OnAfterParameterShapeDeleteCallback OnAfterDeleteInterface[ParameterShape]
	OnAfterParameterShapeReadCallback   OnAfterReadInterface[ParameterShape]

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

	OnAfterParametersAggregateCreateCallback OnAfterCreateInterface[ParametersAggregate]
	OnAfterParametersAggregateUpdateCallback OnAfterUpdateInterface[ParametersAggregate]
	OnAfterParametersAggregateDeleteCallback OnAfterDeleteInterface[ParametersAggregate]
	OnAfterParametersAggregateReadCallback   OnAfterReadInterface[ParametersAggregate]

	ParametersAggregateShapes                map[*ParametersAggregateShape]struct{}
	ParametersAggregateShapes_instance       map[*ParametersAggregateShape]*ParametersAggregateShape
	ParametersAggregateShapes_mapString      map[string]*ParametersAggregateShape
	ParametersAggregateShapeOrder            uint
	ParametersAggregateShape_stagedOrder     map[*ParametersAggregateShape]uint
	ParametersAggregateShape_orderStaged     map[uint]*ParametersAggregateShape
	ParametersAggregateShapes_reference      map[*ParametersAggregateShape]*ParametersAggregateShape
	ParametersAggregateShapes_referenceOrder map[*ParametersAggregateShape]uint

	// insertion point for slice of pointers maps
	OnAfterParametersAggregateShapeCreateCallback OnAfterCreateInterface[ParametersAggregateShape]
	OnAfterParametersAggregateShapeUpdateCallback OnAfterUpdateInterface[ParametersAggregateShape]
	OnAfterParametersAggregateShapeDeleteCallback OnAfterDeleteInterface[ParametersAggregateShape]
	OnAfterParametersAggregateShapeReadCallback   OnAfterReadInterface[ParametersAggregateShape]

	Positions                map[*Position]struct{}
	Positions_instance       map[*Position]*Position
	Positions_mapString      map[string]*Position
	PositionOrder            uint
	Position_stagedOrder     map[*Position]uint
	Position_orderStaged     map[uint]*Position
	Positions_reference      map[*Position]*Position
	Positions_referenceOrder map[*Position]uint

	// insertion point for slice of pointers maps
	OnAfterPositionCreateCallback OnAfterCreateInterface[Position]
	OnAfterPositionUpdateCallback OnAfterUpdateInterface[Position]
	OnAfterPositionDeleteCallback OnAfterDeleteInterface[Position]
	OnAfterPositionReadCallback   OnAfterReadInterface[Position]

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

	OnAfterRepositoryCreateCallback OnAfterCreateInterface[Repository]
	OnAfterRepositoryUpdateCallback OnAfterUpdateInterface[Repository]
	OnAfterRepositoryDeleteCallback OnAfterDeleteInterface[Repository]
	OnAfterRepositoryReadCallback   OnAfterReadInterface[Repository]

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

	OnAfterScenarioCreateCallback OnAfterCreateInterface[Scenario]
	OnAfterScenarioUpdateCallback OnAfterUpdateInterface[Scenario]
	OnAfterScenarioDeleteCallback OnAfterDeleteInterface[Scenario]
	OnAfterScenarioReadCallback   OnAfterReadInterface[Scenario]

	Users                map[*User]struct{}
	Users_instance       map[*User]*User
	Users_mapString      map[string]*User
	UserOrder            uint
	User_stagedOrder     map[*User]uint
	User_orderStaged     map[uint]*User
	Users_reference      map[*User]*User
	Users_referenceOrder map[*User]uint

	// insertion point for slice of pointers maps
	OnAfterUserCreateCallback OnAfterCreateInterface[User]
	OnAfterUserUpdateCallback OnAfterUpdateInterface[User]
	OnAfterUserDeleteCallback OnAfterDeleteInterface[User]
	OnAfterUserReadCallback   OnAfterReadInterface[User]

	UserUses                map[*UserUse]struct{}
	UserUses_instance       map[*UserUse]*UserUse
	UserUses_mapString      map[string]*UserUse
	UserUseOrder            uint
	UserUse_stagedOrder     map[*UserUse]uint
	UserUse_orderStaged     map[uint]*UserUse
	UserUses_reference      map[*UserUse]*UserUse
	UserUses_referenceOrder map[*UserUse]uint

	// insertion point for slice of pointers maps
	OnAfterUserUseCreateCallback OnAfterCreateInterface[UserUse]
	OnAfterUserUseUpdateCallback OnAfterUpdateInterface[UserUse]
	OnAfterUserUseDeleteCallback OnAfterDeleteInterface[UserUse]
	OnAfterUserUseReadCallback   OnAfterReadInterface[UserUse]

	Workspaces                map[*Workspace]struct{}
	Workspaces_instance       map[*Workspace]*Workspace
	Workspaces_mapString      map[string]*Workspace
	WorkspaceOrder            uint
	Workspace_stagedOrder     map[*Workspace]uint
	Workspace_orderStaged     map[uint]*Workspace
	Workspaces_reference      map[*Workspace]*Workspace
	Workspaces_referenceOrder map[*Workspace]uint

	// insertion point for slice of pointers maps
	OnAfterWorkspaceCreateCallback OnAfterCreateInterface[Workspace]
	OnAfterWorkspaceUpdateCallback OnAfterUpdateInterface[Workspace]
	OnAfterWorkspaceDeleteCallback OnAfterDeleteInterface[Workspace]
	OnAfterWorkspaceReadCallback   OnAfterReadInterface[Workspace]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// Private slices to hold the registered hooks
	beforeCommitHooks []func(stage *Stage)
	afterCommitHooks  []func(stage *Stage)

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here

	// store the stage order of each instance in order to
	// preserve this order when serializing them
	// insertion point for order fields declaration
	// end of insertion point

	NamedStructs []*NamedStruct

	// GongUnmarshallers is the registry of all model unmarshallers
	GongUnmarshallers map[string]ModelUnmarshaller

	// probeIF is the interface to the probe that allows log
	// commit event to the probe
	probeIF ProbeIF

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
	err := GongParseAstString(stage, commitToApply, true)
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
	err := GongParseAstString(stage, commitToApply, true)
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
	stage.ActorStates_reference = make(map[*ActorState]*ActorState)
	stage.ActorStates_instance = make(map[*ActorState]*ActorState)
	stage.ActorStates_referenceOrder = make(map[*ActorState]uint)

	stage.ActorStateShapes_reference = make(map[*ActorStateShape]*ActorStateShape)
	stage.ActorStateShapes_instance = make(map[*ActorStateShape]*ActorStateShape)
	stage.ActorStateShapes_referenceOrder = make(map[*ActorStateShape]uint)

	stage.ActorStateTransitions_reference = make(map[*ActorStateTransition]*ActorStateTransition)
	stage.ActorStateTransitions_instance = make(map[*ActorStateTransition]*ActorStateTransition)
	stage.ActorStateTransitions_referenceOrder = make(map[*ActorStateTransition]uint)

	stage.ActorStateTransitionShapes_reference = make(map[*ActorStateTransitionShape]*ActorStateTransitionShape)
	stage.ActorStateTransitionShapes_instance = make(map[*ActorStateTransitionShape]*ActorStateTransitionShape)
	stage.ActorStateTransitionShapes_referenceOrder = make(map[*ActorStateTransitionShape]uint)

	stage.Analysiss_reference = make(map[*Analysis]*Analysis)
	stage.Analysiss_instance = make(map[*Analysis]*Analysis)
	stage.Analysiss_referenceOrder = make(map[*Analysis]uint)

	stage.ControlPointShapes_reference = make(map[*ControlPointShape]*ControlPointShape)
	stage.ControlPointShapes_instance = make(map[*ControlPointShape]*ControlPointShape)
	stage.ControlPointShapes_referenceOrder = make(map[*ControlPointShape]uint)

	stage.Diagrams_reference = make(map[*Diagram]*Diagram)
	stage.Diagrams_instance = make(map[*Diagram]*Diagram)
	stage.Diagrams_referenceOrder = make(map[*Diagram]uint)

	stage.Documents_reference = make(map[*Document]*Document)
	stage.Documents_instance = make(map[*Document]*Document)
	stage.Documents_referenceOrder = make(map[*Document]uint)

	stage.DocumentUses_reference = make(map[*DocumentUse]*DocumentUse)
	stage.DocumentUses_instance = make(map[*DocumentUse]*DocumentUse)
	stage.DocumentUses_referenceOrder = make(map[*DocumentUse]uint)

	stage.EvolutionDirections_reference = make(map[*EvolutionDirection]*EvolutionDirection)
	stage.EvolutionDirections_instance = make(map[*EvolutionDirection]*EvolutionDirection)
	stage.EvolutionDirections_referenceOrder = make(map[*EvolutionDirection]uint)

	stage.EvolutionDirectionShapes_reference = make(map[*EvolutionDirectionShape]*EvolutionDirectionShape)
	stage.EvolutionDirectionShapes_instance = make(map[*EvolutionDirectionShape]*EvolutionDirectionShape)
	stage.EvolutionDirectionShapes_referenceOrder = make(map[*EvolutionDirectionShape]uint)

	stage.Foos_reference = make(map[*Foo]*Foo)
	stage.Foos_instance = make(map[*Foo]*Foo)
	stage.Foos_referenceOrder = make(map[*Foo]uint)

	stage.GeoObjects_reference = make(map[*GeoObject]*GeoObject)
	stage.GeoObjects_instance = make(map[*GeoObject]*GeoObject)
	stage.GeoObjects_referenceOrder = make(map[*GeoObject]uint)

	stage.GeoObjectUses_reference = make(map[*GeoObjectUse]*GeoObjectUse)
	stage.GeoObjectUses_instance = make(map[*GeoObjectUse]*GeoObjectUse)
	stage.GeoObjectUses_referenceOrder = make(map[*GeoObjectUse]uint)

	stage.Groups_reference = make(map[*Group]*Group)
	stage.Groups_instance = make(map[*Group]*Group)
	stage.Groups_referenceOrder = make(map[*Group]uint)

	stage.GroupUses_reference = make(map[*GroupUse]*GroupUse)
	stage.GroupUses_instance = make(map[*GroupUse]*GroupUse)
	stage.GroupUses_referenceOrder = make(map[*GroupUse]uint)

	stage.Librarys_reference = make(map[*Library]*Library)
	stage.Librarys_instance = make(map[*Library]*Library)
	stage.Librarys_referenceOrder = make(map[*Library]uint)

	stage.MapObjects_reference = make(map[*MapObject]*MapObject)
	stage.MapObjects_instance = make(map[*MapObject]*MapObject)
	stage.MapObjects_referenceOrder = make(map[*MapObject]uint)

	stage.MapObjectUses_reference = make(map[*MapObjectUse]*MapObjectUse)
	stage.MapObjectUses_instance = make(map[*MapObjectUse]*MapObjectUse)
	stage.MapObjectUses_referenceOrder = make(map[*MapObjectUse]uint)

	stage.Parameters_reference = make(map[*Parameter]*Parameter)
	stage.Parameters_instance = make(map[*Parameter]*Parameter)
	stage.Parameters_referenceOrder = make(map[*Parameter]uint)

	stage.ParameterCategorys_reference = make(map[*ParameterCategory]*ParameterCategory)
	stage.ParameterCategorys_instance = make(map[*ParameterCategory]*ParameterCategory)
	stage.ParameterCategorys_referenceOrder = make(map[*ParameterCategory]uint)

	stage.ParameterCategoryUses_reference = make(map[*ParameterCategoryUse]*ParameterCategoryUse)
	stage.ParameterCategoryUses_instance = make(map[*ParameterCategoryUse]*ParameterCategoryUse)
	stage.ParameterCategoryUses_referenceOrder = make(map[*ParameterCategoryUse]uint)

	stage.ParameterShapes_reference = make(map[*ParameterShape]*ParameterShape)
	stage.ParameterShapes_instance = make(map[*ParameterShape]*ParameterShape)
	stage.ParameterShapes_referenceOrder = make(map[*ParameterShape]uint)

	stage.ParametersAggregates_reference = make(map[*ParametersAggregate]*ParametersAggregate)
	stage.ParametersAggregates_instance = make(map[*ParametersAggregate]*ParametersAggregate)
	stage.ParametersAggregates_referenceOrder = make(map[*ParametersAggregate]uint)

	stage.ParametersAggregateShapes_reference = make(map[*ParametersAggregateShape]*ParametersAggregateShape)
	stage.ParametersAggregateShapes_instance = make(map[*ParametersAggregateShape]*ParametersAggregateShape)
	stage.ParametersAggregateShapes_referenceOrder = make(map[*ParametersAggregateShape]uint)

	stage.Positions_reference = make(map[*Position]*Position)
	stage.Positions_instance = make(map[*Position]*Position)
	stage.Positions_referenceOrder = make(map[*Position]uint)

	stage.Repositorys_reference = make(map[*Repository]*Repository)
	stage.Repositorys_instance = make(map[*Repository]*Repository)
	stage.Repositorys_referenceOrder = make(map[*Repository]uint)

	stage.Scenarios_reference = make(map[*Scenario]*Scenario)
	stage.Scenarios_instance = make(map[*Scenario]*Scenario)
	stage.Scenarios_referenceOrder = make(map[*Scenario]uint)

	stage.Users_reference = make(map[*User]*User)
	stage.Users_instance = make(map[*User]*User)
	stage.Users_referenceOrder = make(map[*User]uint)

	stage.UserUses_reference = make(map[*UserUse]*UserUse)
	stage.UserUses_instance = make(map[*UserUse]*UserUse)
	stage.UserUses_referenceOrder = make(map[*UserUse]uint)

	stage.Workspaces_reference = make(map[*Workspace]*Workspace)
	stage.Workspaces_instance = make(map[*Workspace]*Workspace)
	stage.Workspaces_referenceOrder = make(map[*Workspace]uint)

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
	var maxActorStateOrder uint
	var foundActorState bool
	for _, order := range stage.ActorState_stagedOrder {
		if !foundActorState || order > maxActorStateOrder {
			maxActorStateOrder = order
			foundActorState = true
		}
	}
	if foundActorState {
		stage.ActorStateOrder = maxActorStateOrder + 1
	} else {
		stage.ActorStateOrder = 0
	}

	var maxActorStateShapeOrder uint
	var foundActorStateShape bool
	for _, order := range stage.ActorStateShape_stagedOrder {
		if !foundActorStateShape || order > maxActorStateShapeOrder {
			maxActorStateShapeOrder = order
			foundActorStateShape = true
		}
	}
	if foundActorStateShape {
		stage.ActorStateShapeOrder = maxActorStateShapeOrder + 1
	} else {
		stage.ActorStateShapeOrder = 0
	}

	var maxActorStateTransitionOrder uint
	var foundActorStateTransition bool
	for _, order := range stage.ActorStateTransition_stagedOrder {
		if !foundActorStateTransition || order > maxActorStateTransitionOrder {
			maxActorStateTransitionOrder = order
			foundActorStateTransition = true
		}
	}
	if foundActorStateTransition {
		stage.ActorStateTransitionOrder = maxActorStateTransitionOrder + 1
	} else {
		stage.ActorStateTransitionOrder = 0
	}

	var maxActorStateTransitionShapeOrder uint
	var foundActorStateTransitionShape bool
	for _, order := range stage.ActorStateTransitionShape_stagedOrder {
		if !foundActorStateTransitionShape || order > maxActorStateTransitionShapeOrder {
			maxActorStateTransitionShapeOrder = order
			foundActorStateTransitionShape = true
		}
	}
	if foundActorStateTransitionShape {
		stage.ActorStateTransitionShapeOrder = maxActorStateTransitionShapeOrder + 1
	} else {
		stage.ActorStateTransitionShapeOrder = 0
	}

	var maxAnalysisOrder uint
	var foundAnalysis bool
	for _, order := range stage.Analysis_stagedOrder {
		if !foundAnalysis || order > maxAnalysisOrder {
			maxAnalysisOrder = order
			foundAnalysis = true
		}
	}
	if foundAnalysis {
		stage.AnalysisOrder = maxAnalysisOrder + 1
	} else {
		stage.AnalysisOrder = 0
	}

	var maxControlPointShapeOrder uint
	var foundControlPointShape bool
	for _, order := range stage.ControlPointShape_stagedOrder {
		if !foundControlPointShape || order > maxControlPointShapeOrder {
			maxControlPointShapeOrder = order
			foundControlPointShape = true
		}
	}
	if foundControlPointShape {
		stage.ControlPointShapeOrder = maxControlPointShapeOrder + 1
	} else {
		stage.ControlPointShapeOrder = 0
	}

	var maxDiagramOrder uint
	var foundDiagram bool
	for _, order := range stage.Diagram_stagedOrder {
		if !foundDiagram || order > maxDiagramOrder {
			maxDiagramOrder = order
			foundDiagram = true
		}
	}
	if foundDiagram {
		stage.DiagramOrder = maxDiagramOrder + 1
	} else {
		stage.DiagramOrder = 0
	}

	var maxDocumentOrder uint
	var foundDocument bool
	for _, order := range stage.Document_stagedOrder {
		if !foundDocument || order > maxDocumentOrder {
			maxDocumentOrder = order
			foundDocument = true
		}
	}
	if foundDocument {
		stage.DocumentOrder = maxDocumentOrder + 1
	} else {
		stage.DocumentOrder = 0
	}

	var maxDocumentUseOrder uint
	var foundDocumentUse bool
	for _, order := range stage.DocumentUse_stagedOrder {
		if !foundDocumentUse || order > maxDocumentUseOrder {
			maxDocumentUseOrder = order
			foundDocumentUse = true
		}
	}
	if foundDocumentUse {
		stage.DocumentUseOrder = maxDocumentUseOrder + 1
	} else {
		stage.DocumentUseOrder = 0
	}

	var maxEvolutionDirectionOrder uint
	var foundEvolutionDirection bool
	for _, order := range stage.EvolutionDirection_stagedOrder {
		if !foundEvolutionDirection || order > maxEvolutionDirectionOrder {
			maxEvolutionDirectionOrder = order
			foundEvolutionDirection = true
		}
	}
	if foundEvolutionDirection {
		stage.EvolutionDirectionOrder = maxEvolutionDirectionOrder + 1
	} else {
		stage.EvolutionDirectionOrder = 0
	}

	var maxEvolutionDirectionShapeOrder uint
	var foundEvolutionDirectionShape bool
	for _, order := range stage.EvolutionDirectionShape_stagedOrder {
		if !foundEvolutionDirectionShape || order > maxEvolutionDirectionShapeOrder {
			maxEvolutionDirectionShapeOrder = order
			foundEvolutionDirectionShape = true
		}
	}
	if foundEvolutionDirectionShape {
		stage.EvolutionDirectionShapeOrder = maxEvolutionDirectionShapeOrder + 1
	} else {
		stage.EvolutionDirectionShapeOrder = 0
	}

	var maxFooOrder uint
	var foundFoo bool
	for _, order := range stage.Foo_stagedOrder {
		if !foundFoo || order > maxFooOrder {
			maxFooOrder = order
			foundFoo = true
		}
	}
	if foundFoo {
		stage.FooOrder = maxFooOrder + 1
	} else {
		stage.FooOrder = 0
	}

	var maxGeoObjectOrder uint
	var foundGeoObject bool
	for _, order := range stage.GeoObject_stagedOrder {
		if !foundGeoObject || order > maxGeoObjectOrder {
			maxGeoObjectOrder = order
			foundGeoObject = true
		}
	}
	if foundGeoObject {
		stage.GeoObjectOrder = maxGeoObjectOrder + 1
	} else {
		stage.GeoObjectOrder = 0
	}

	var maxGeoObjectUseOrder uint
	var foundGeoObjectUse bool
	for _, order := range stage.GeoObjectUse_stagedOrder {
		if !foundGeoObjectUse || order > maxGeoObjectUseOrder {
			maxGeoObjectUseOrder = order
			foundGeoObjectUse = true
		}
	}
	if foundGeoObjectUse {
		stage.GeoObjectUseOrder = maxGeoObjectUseOrder + 1
	} else {
		stage.GeoObjectUseOrder = 0
	}

	var maxGroupOrder uint
	var foundGroup bool
	for _, order := range stage.Group_stagedOrder {
		if !foundGroup || order > maxGroupOrder {
			maxGroupOrder = order
			foundGroup = true
		}
	}
	if foundGroup {
		stage.GroupOrder = maxGroupOrder + 1
	} else {
		stage.GroupOrder = 0
	}

	var maxGroupUseOrder uint
	var foundGroupUse bool
	for _, order := range stage.GroupUse_stagedOrder {
		if !foundGroupUse || order > maxGroupUseOrder {
			maxGroupUseOrder = order
			foundGroupUse = true
		}
	}
	if foundGroupUse {
		stage.GroupUseOrder = maxGroupUseOrder + 1
	} else {
		stage.GroupUseOrder = 0
	}

	var maxLibraryOrder uint
	var foundLibrary bool
	for _, order := range stage.Library_stagedOrder {
		if !foundLibrary || order > maxLibraryOrder {
			maxLibraryOrder = order
			foundLibrary = true
		}
	}
	if foundLibrary {
		stage.LibraryOrder = maxLibraryOrder + 1
	} else {
		stage.LibraryOrder = 0
	}

	var maxMapObjectOrder uint
	var foundMapObject bool
	for _, order := range stage.MapObject_stagedOrder {
		if !foundMapObject || order > maxMapObjectOrder {
			maxMapObjectOrder = order
			foundMapObject = true
		}
	}
	if foundMapObject {
		stage.MapObjectOrder = maxMapObjectOrder + 1
	} else {
		stage.MapObjectOrder = 0
	}

	var maxMapObjectUseOrder uint
	var foundMapObjectUse bool
	for _, order := range stage.MapObjectUse_stagedOrder {
		if !foundMapObjectUse || order > maxMapObjectUseOrder {
			maxMapObjectUseOrder = order
			foundMapObjectUse = true
		}
	}
	if foundMapObjectUse {
		stage.MapObjectUseOrder = maxMapObjectUseOrder + 1
	} else {
		stage.MapObjectUseOrder = 0
	}

	var maxParameterOrder uint
	var foundParameter bool
	for _, order := range stage.Parameter_stagedOrder {
		if !foundParameter || order > maxParameterOrder {
			maxParameterOrder = order
			foundParameter = true
		}
	}
	if foundParameter {
		stage.ParameterOrder = maxParameterOrder + 1
	} else {
		stage.ParameterOrder = 0
	}

	var maxParameterCategoryOrder uint
	var foundParameterCategory bool
	for _, order := range stage.ParameterCategory_stagedOrder {
		if !foundParameterCategory || order > maxParameterCategoryOrder {
			maxParameterCategoryOrder = order
			foundParameterCategory = true
		}
	}
	if foundParameterCategory {
		stage.ParameterCategoryOrder = maxParameterCategoryOrder + 1
	} else {
		stage.ParameterCategoryOrder = 0
	}

	var maxParameterCategoryUseOrder uint
	var foundParameterCategoryUse bool
	for _, order := range stage.ParameterCategoryUse_stagedOrder {
		if !foundParameterCategoryUse || order > maxParameterCategoryUseOrder {
			maxParameterCategoryUseOrder = order
			foundParameterCategoryUse = true
		}
	}
	if foundParameterCategoryUse {
		stage.ParameterCategoryUseOrder = maxParameterCategoryUseOrder + 1
	} else {
		stage.ParameterCategoryUseOrder = 0
	}

	var maxParameterShapeOrder uint
	var foundParameterShape bool
	for _, order := range stage.ParameterShape_stagedOrder {
		if !foundParameterShape || order > maxParameterShapeOrder {
			maxParameterShapeOrder = order
			foundParameterShape = true
		}
	}
	if foundParameterShape {
		stage.ParameterShapeOrder = maxParameterShapeOrder + 1
	} else {
		stage.ParameterShapeOrder = 0
	}

	var maxParametersAggregateOrder uint
	var foundParametersAggregate bool
	for _, order := range stage.ParametersAggregate_stagedOrder {
		if !foundParametersAggregate || order > maxParametersAggregateOrder {
			maxParametersAggregateOrder = order
			foundParametersAggregate = true
		}
	}
	if foundParametersAggregate {
		stage.ParametersAggregateOrder = maxParametersAggregateOrder + 1
	} else {
		stage.ParametersAggregateOrder = 0
	}

	var maxParametersAggregateShapeOrder uint
	var foundParametersAggregateShape bool
	for _, order := range stage.ParametersAggregateShape_stagedOrder {
		if !foundParametersAggregateShape || order > maxParametersAggregateShapeOrder {
			maxParametersAggregateShapeOrder = order
			foundParametersAggregateShape = true
		}
	}
	if foundParametersAggregateShape {
		stage.ParametersAggregateShapeOrder = maxParametersAggregateShapeOrder + 1
	} else {
		stage.ParametersAggregateShapeOrder = 0
	}

	var maxPositionOrder uint
	var foundPosition bool
	for _, order := range stage.Position_stagedOrder {
		if !foundPosition || order > maxPositionOrder {
			maxPositionOrder = order
			foundPosition = true
		}
	}
	if foundPosition {
		stage.PositionOrder = maxPositionOrder + 1
	} else {
		stage.PositionOrder = 0
	}

	var maxRepositoryOrder uint
	var foundRepository bool
	for _, order := range stage.Repository_stagedOrder {
		if !foundRepository || order > maxRepositoryOrder {
			maxRepositoryOrder = order
			foundRepository = true
		}
	}
	if foundRepository {
		stage.RepositoryOrder = maxRepositoryOrder + 1
	} else {
		stage.RepositoryOrder = 0
	}

	var maxScenarioOrder uint
	var foundScenario bool
	for _, order := range stage.Scenario_stagedOrder {
		if !foundScenario || order > maxScenarioOrder {
			maxScenarioOrder = order
			foundScenario = true
		}
	}
	if foundScenario {
		stage.ScenarioOrder = maxScenarioOrder + 1
	} else {
		stage.ScenarioOrder = 0
	}

	var maxUserOrder uint
	var foundUser bool
	for _, order := range stage.User_stagedOrder {
		if !foundUser || order > maxUserOrder {
			maxUserOrder = order
			foundUser = true
		}
	}
	if foundUser {
		stage.UserOrder = maxUserOrder + 1
	} else {
		stage.UserOrder = 0
	}

	var maxUserUseOrder uint
	var foundUserUse bool
	for _, order := range stage.UserUse_stagedOrder {
		if !foundUserUse || order > maxUserUseOrder {
			maxUserUseOrder = order
			foundUserUse = true
		}
	}
	if foundUserUse {
		stage.UserUseOrder = maxUserUseOrder + 1
	} else {
		stage.UserUseOrder = 0
	}

	var maxWorkspaceOrder uint
	var foundWorkspace bool
	for _, order := range stage.Workspace_stagedOrder {
		if !foundWorkspace || order > maxWorkspaceOrder {
			maxWorkspaceOrder = order
			foundWorkspace = true
		}
	}
	if foundWorkspace {
		stage.WorkspaceOrder = maxWorkspaceOrder + 1
	} else {
		stage.WorkspaceOrder = 0
	}

	// end of insertion point for max order recomputation
}

func (stage *Stage) SetDeltaMode(inDeltaMode bool) {
	stage.isInDeltaMode = inDeltaMode
}

func (stage *Stage) IsInDeltaMode() bool {
	return stage.isInDeltaMode
}

func (stage *Stage) SetProbeIF(probeIF ProbeIF) {
	stage.probeIF = probeIF
}

func (stage *Stage) GetProbeIF() ProbeIF {
	if stage.probeIF == nil {
		return nil
	}

	return stage.probeIF
}

// GetNamedStructs implements models.ProbebStage.
func (stage *Stage) GetNamedStructsNames() (res []string) {
	for _, namedStruct := range stage.NamedStructs {
		res = append(res, namedStruct.name)
	}

	return
}

func GetNamedStructInstances[T PointerToGongstruct](set map[T]struct{}, order map[T]uint) (res []string) {
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
			log.Fatalf("GetNamedStructInstances: pointer not found")
		}
		return i_order < j_order
	})

	for _, instance := range orderedSet {
		res = append(res, instance.GetName())
	}

	return
}

// GetStructInstancesByOrderAuto returns a slice of generic pointers to gongstructs
// ordered by their order in the stage.
func GetStructInstancesByOrderAuto[T PointerToGongstruct](stage *Stage) (res []T) {
	var t T
	switch any(t).(type) {
	// insertion point for case
	case *ActorState:
		tmp := GetStructInstancesByOrder(stage.ActorStates, stage.ActorState_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ActorState implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ActorStateShape:
		tmp := GetStructInstancesByOrder(stage.ActorStateShapes, stage.ActorStateShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ActorStateShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ActorStateTransition:
		tmp := GetStructInstancesByOrder(stage.ActorStateTransitions, stage.ActorStateTransition_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ActorStateTransition implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ActorStateTransitionShape:
		tmp := GetStructInstancesByOrder(stage.ActorStateTransitionShapes, stage.ActorStateTransitionShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ActorStateTransitionShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Analysis:
		tmp := GetStructInstancesByOrder(stage.Analysiss, stage.Analysis_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Analysis implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ControlPointShape:
		tmp := GetStructInstancesByOrder(stage.ControlPointShapes, stage.ControlPointShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ControlPointShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Diagram:
		tmp := GetStructInstancesByOrder(stage.Diagrams, stage.Diagram_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Diagram implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Document:
		tmp := GetStructInstancesByOrder(stage.Documents, stage.Document_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Document implements.
			res = append(res, any(v).(T))
		}
		return res
	case *DocumentUse:
		tmp := GetStructInstancesByOrder(stage.DocumentUses, stage.DocumentUse_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *DocumentUse implements.
			res = append(res, any(v).(T))
		}
		return res
	case *EvolutionDirection:
		tmp := GetStructInstancesByOrder(stage.EvolutionDirections, stage.EvolutionDirection_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *EvolutionDirection implements.
			res = append(res, any(v).(T))
		}
		return res
	case *EvolutionDirectionShape:
		tmp := GetStructInstancesByOrder(stage.EvolutionDirectionShapes, stage.EvolutionDirectionShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *EvolutionDirectionShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Foo:
		tmp := GetStructInstancesByOrder(stage.Foos, stage.Foo_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Foo implements.
			res = append(res, any(v).(T))
		}
		return res
	case *GeoObject:
		tmp := GetStructInstancesByOrder(stage.GeoObjects, stage.GeoObject_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GeoObject implements.
			res = append(res, any(v).(T))
		}
		return res
	case *GeoObjectUse:
		tmp := GetStructInstancesByOrder(stage.GeoObjectUses, stage.GeoObjectUse_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GeoObjectUse implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Group:
		tmp := GetStructInstancesByOrder(stage.Groups, stage.Group_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Group implements.
			res = append(res, any(v).(T))
		}
		return res
	case *GroupUse:
		tmp := GetStructInstancesByOrder(stage.GroupUses, stage.GroupUse_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GroupUse implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Library:
		tmp := GetStructInstancesByOrder(stage.Librarys, stage.Library_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Library implements.
			res = append(res, any(v).(T))
		}
		return res
	case *MapObject:
		tmp := GetStructInstancesByOrder(stage.MapObjects, stage.MapObject_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *MapObject implements.
			res = append(res, any(v).(T))
		}
		return res
	case *MapObjectUse:
		tmp := GetStructInstancesByOrder(stage.MapObjectUses, stage.MapObjectUse_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *MapObjectUse implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Parameter:
		tmp := GetStructInstancesByOrder(stage.Parameters, stage.Parameter_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Parameter implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ParameterCategory:
		tmp := GetStructInstancesByOrder(stage.ParameterCategorys, stage.ParameterCategory_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ParameterCategory implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ParameterCategoryUse:
		tmp := GetStructInstancesByOrder(stage.ParameterCategoryUses, stage.ParameterCategoryUse_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ParameterCategoryUse implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ParameterShape:
		tmp := GetStructInstancesByOrder(stage.ParameterShapes, stage.ParameterShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ParameterShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ParametersAggregate:
		tmp := GetStructInstancesByOrder(stage.ParametersAggregates, stage.ParametersAggregate_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ParametersAggregate implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ParametersAggregateShape:
		tmp := GetStructInstancesByOrder(stage.ParametersAggregateShapes, stage.ParametersAggregateShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ParametersAggregateShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Position:
		tmp := GetStructInstancesByOrder(stage.Positions, stage.Position_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Position implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Repository:
		tmp := GetStructInstancesByOrder(stage.Repositorys, stage.Repository_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Repository implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Scenario:
		tmp := GetStructInstancesByOrder(stage.Scenarios, stage.Scenario_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Scenario implements.
			res = append(res, any(v).(T))
		}
		return res
	case *User:
		tmp := GetStructInstancesByOrder(stage.Users, stage.User_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *User implements.
			res = append(res, any(v).(T))
		}
		return res
	case *UserUse:
		tmp := GetStructInstancesByOrder(stage.UserUses, stage.UserUse_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *UserUse implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Workspace:
		tmp := GetStructInstancesByOrder(stage.Workspaces, stage.Workspace_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Workspace implements.
			res = append(res, any(v).(T))
		}
		return res

	}
	return
}

func GetStructInstancesByOrder[T PointerToGongstruct](set map[T]struct{}, order map[T]uint) (res []T) {
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
			log.Fatalf("GetNamedStructInstances: pointer not found")
		}
		return i_order < j_order
	})

	res = append(res, orderedSet...)

	return
}

func (stage *Stage) GetNamedStructNamesByOrder(namedStructName string) (res []string) {
	switch namedStructName {
	// insertion point for case
	case "ActorState":
		res = GetNamedStructInstances(stage.ActorStates, stage.ActorState_stagedOrder)
	case "ActorStateShape":
		res = GetNamedStructInstances(stage.ActorStateShapes, stage.ActorStateShape_stagedOrder)
	case "ActorStateTransition":
		res = GetNamedStructInstances(stage.ActorStateTransitions, stage.ActorStateTransition_stagedOrder)
	case "ActorStateTransitionShape":
		res = GetNamedStructInstances(stage.ActorStateTransitionShapes, stage.ActorStateTransitionShape_stagedOrder)
	case "Analysis":
		res = GetNamedStructInstances(stage.Analysiss, stage.Analysis_stagedOrder)
	case "ControlPointShape":
		res = GetNamedStructInstances(stage.ControlPointShapes, stage.ControlPointShape_stagedOrder)
	case "Diagram":
		res = GetNamedStructInstances(stage.Diagrams, stage.Diagram_stagedOrder)
	case "Document":
		res = GetNamedStructInstances(stage.Documents, stage.Document_stagedOrder)
	case "DocumentUse":
		res = GetNamedStructInstances(stage.DocumentUses, stage.DocumentUse_stagedOrder)
	case "EvolutionDirection":
		res = GetNamedStructInstances(stage.EvolutionDirections, stage.EvolutionDirection_stagedOrder)
	case "EvolutionDirectionShape":
		res = GetNamedStructInstances(stage.EvolutionDirectionShapes, stage.EvolutionDirectionShape_stagedOrder)
	case "Foo":
		res = GetNamedStructInstances(stage.Foos, stage.Foo_stagedOrder)
	case "GeoObject":
		res = GetNamedStructInstances(stage.GeoObjects, stage.GeoObject_stagedOrder)
	case "GeoObjectUse":
		res = GetNamedStructInstances(stage.GeoObjectUses, stage.GeoObjectUse_stagedOrder)
	case "Group":
		res = GetNamedStructInstances(stage.Groups, stage.Group_stagedOrder)
	case "GroupUse":
		res = GetNamedStructInstances(stage.GroupUses, stage.GroupUse_stagedOrder)
	case "Library":
		res = GetNamedStructInstances(stage.Librarys, stage.Library_stagedOrder)
	case "MapObject":
		res = GetNamedStructInstances(stage.MapObjects, stage.MapObject_stagedOrder)
	case "MapObjectUse":
		res = GetNamedStructInstances(stage.MapObjectUses, stage.MapObjectUse_stagedOrder)
	case "Parameter":
		res = GetNamedStructInstances(stage.Parameters, stage.Parameter_stagedOrder)
	case "ParameterCategory":
		res = GetNamedStructInstances(stage.ParameterCategorys, stage.ParameterCategory_stagedOrder)
	case "ParameterCategoryUse":
		res = GetNamedStructInstances(stage.ParameterCategoryUses, stage.ParameterCategoryUse_stagedOrder)
	case "ParameterShape":
		res = GetNamedStructInstances(stage.ParameterShapes, stage.ParameterShape_stagedOrder)
	case "ParametersAggregate":
		res = GetNamedStructInstances(stage.ParametersAggregates, stage.ParametersAggregate_stagedOrder)
	case "ParametersAggregateShape":
		res = GetNamedStructInstances(stage.ParametersAggregateShapes, stage.ParametersAggregateShape_stagedOrder)
	case "Position":
		res = GetNamedStructInstances(stage.Positions, stage.Position_stagedOrder)
	case "Repository":
		res = GetNamedStructInstances(stage.Repositorys, stage.Repository_stagedOrder)
	case "Scenario":
		res = GetNamedStructInstances(stage.Scenarios, stage.Scenario_stagedOrder)
	case "User":
		res = GetNamedStructInstances(stage.Users, stage.User_stagedOrder)
	case "UserUse":
		res = GetNamedStructInstances(stage.UserUses, stage.UserUse_stagedOrder)
	case "Workspace":
		res = GetNamedStructInstances(stage.Workspaces, stage.Workspace_stagedOrder)
	}

	return
}

type NamedStruct struct {
	name string
}

func (namedStruct *NamedStruct) GetName() string {
	return namedStruct.name
}

func (stage *Stage) GetType() string {
	return "github.com/fullstack-lang/gong/dsm/scenario/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return scenario_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return scenario_go.GoDiagramsDir
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *Stage)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *Stage,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *Stage,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *Stage, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *Stage,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *Stage)
	Checkout(stage *Stage)
	Backup(stage *Stage, dirPath string)
	Restore(stage *Stage, dirPath string)
	BackupXL(stage *Stage, dirPath string)
	RestoreXL(stage *Stage, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitActorState(actorstate *ActorState)
	CheckoutActorState(actorstate *ActorState)
	CommitActorStateShape(actorstateshape *ActorStateShape)
	CheckoutActorStateShape(actorstateshape *ActorStateShape)
	CommitActorStateTransition(actorstatetransition *ActorStateTransition)
	CheckoutActorStateTransition(actorstatetransition *ActorStateTransition)
	CommitActorStateTransitionShape(actorstatetransitionshape *ActorStateTransitionShape)
	CheckoutActorStateTransitionShape(actorstatetransitionshape *ActorStateTransitionShape)
	CommitAnalysis(analysis *Analysis)
	CheckoutAnalysis(analysis *Analysis)
	CommitControlPointShape(controlpointshape *ControlPointShape)
	CheckoutControlPointShape(controlpointshape *ControlPointShape)
	CommitDiagram(diagram *Diagram)
	CheckoutDiagram(diagram *Diagram)
	CommitDocument(document *Document)
	CheckoutDocument(document *Document)
	CommitDocumentUse(documentuse *DocumentUse)
	CheckoutDocumentUse(documentuse *DocumentUse)
	CommitEvolutionDirection(evolutiondirection *EvolutionDirection)
	CheckoutEvolutionDirection(evolutiondirection *EvolutionDirection)
	CommitEvolutionDirectionShape(evolutiondirectionshape *EvolutionDirectionShape)
	CheckoutEvolutionDirectionShape(evolutiondirectionshape *EvolutionDirectionShape)
	CommitFoo(foo *Foo)
	CheckoutFoo(foo *Foo)
	CommitGeoObject(geoobject *GeoObject)
	CheckoutGeoObject(geoobject *GeoObject)
	CommitGeoObjectUse(geoobjectuse *GeoObjectUse)
	CheckoutGeoObjectUse(geoobjectuse *GeoObjectUse)
	CommitGroup(group *Group)
	CheckoutGroup(group *Group)
	CommitGroupUse(groupuse *GroupUse)
	CheckoutGroupUse(groupuse *GroupUse)
	CommitLibrary(library *Library)
	CheckoutLibrary(library *Library)
	CommitMapObject(mapobject *MapObject)
	CheckoutMapObject(mapobject *MapObject)
	CommitMapObjectUse(mapobjectuse *MapObjectUse)
	CheckoutMapObjectUse(mapobjectuse *MapObjectUse)
	CommitParameter(parameter *Parameter)
	CheckoutParameter(parameter *Parameter)
	CommitParameterCategory(parametercategory *ParameterCategory)
	CheckoutParameterCategory(parametercategory *ParameterCategory)
	CommitParameterCategoryUse(parametercategoryuse *ParameterCategoryUse)
	CheckoutParameterCategoryUse(parametercategoryuse *ParameterCategoryUse)
	CommitParameterShape(parametershape *ParameterShape)
	CheckoutParameterShape(parametershape *ParameterShape)
	CommitParametersAggregate(parametersaggregate *ParametersAggregate)
	CheckoutParametersAggregate(parametersaggregate *ParametersAggregate)
	CommitParametersAggregateShape(parametersaggregateshape *ParametersAggregateShape)
	CheckoutParametersAggregateShape(parametersaggregateshape *ParametersAggregateShape)
	CommitPosition(position *Position)
	CheckoutPosition(position *Position)
	CommitRepository(repository *Repository)
	CheckoutRepository(repository *Repository)
	CommitScenario(scenario *Scenario)
	CheckoutScenario(scenario *Scenario)
	CommitUser(user *User)
	CheckoutUser(user *User)
	CommitUserUse(useruse *UserUse)
	CheckoutUserUse(useruse *UserUse)
	CommitWorkspace(workspace *Workspace)
	CheckoutWorkspace(workspace *Workspace)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

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
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
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

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "ActorState"},
			{name: "ActorStateShape"},
			{name: "ActorStateTransition"},
			{name: "ActorStateTransitionShape"},
			{name: "Analysis"},
			{name: "ControlPointShape"},
			{name: "Diagram"},
			{name: "Document"},
			{name: "DocumentUse"},
			{name: "EvolutionDirection"},
			{name: "EvolutionDirectionShape"},
			{name: "Foo"},
			{name: "GeoObject"},
			{name: "GeoObjectUse"},
			{name: "Group"},
			{name: "GroupUse"},
			{name: "Library"},
			{name: "MapObject"},
			{name: "MapObjectUse"},
			{name: "Parameter"},
			{name: "ParameterCategory"},
			{name: "ParameterCategoryUse"},
			{name: "ParameterShape"},
			{name: "ParametersAggregate"},
			{name: "ParametersAggregateShape"},
			{name: "Position"},
			{name: "Repository"},
			{name: "Scenario"},
			{name: "User"},
			{name: "UserUse"},
			{name: "Workspace"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *ActorState:
		return stage.ActorState_stagedOrder[instance]
	case *ActorStateShape:
		return stage.ActorStateShape_stagedOrder[instance]
	case *ActorStateTransition:
		return stage.ActorStateTransition_stagedOrder[instance]
	case *ActorStateTransitionShape:
		return stage.ActorStateTransitionShape_stagedOrder[instance]
	case *Analysis:
		return stage.Analysis_stagedOrder[instance]
	case *ControlPointShape:
		return stage.ControlPointShape_stagedOrder[instance]
	case *Diagram:
		return stage.Diagram_stagedOrder[instance]
	case *Document:
		return stage.Document_stagedOrder[instance]
	case *DocumentUse:
		return stage.DocumentUse_stagedOrder[instance]
	case *EvolutionDirection:
		return stage.EvolutionDirection_stagedOrder[instance]
	case *EvolutionDirectionShape:
		return stage.EvolutionDirectionShape_stagedOrder[instance]
	case *Foo:
		return stage.Foo_stagedOrder[instance]
	case *GeoObject:
		return stage.GeoObject_stagedOrder[instance]
	case *GeoObjectUse:
		return stage.GeoObjectUse_stagedOrder[instance]
	case *Group:
		return stage.Group_stagedOrder[instance]
	case *GroupUse:
		return stage.GroupUse_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *MapObject:
		return stage.MapObject_stagedOrder[instance]
	case *MapObjectUse:
		return stage.MapObjectUse_stagedOrder[instance]
	case *Parameter:
		return stage.Parameter_stagedOrder[instance]
	case *ParameterCategory:
		return stage.ParameterCategory_stagedOrder[instance]
	case *ParameterCategoryUse:
		return stage.ParameterCategoryUse_stagedOrder[instance]
	case *ParameterShape:
		return stage.ParameterShape_stagedOrder[instance]
	case *ParametersAggregate:
		return stage.ParametersAggregate_stagedOrder[instance]
	case *ParametersAggregateShape:
		return stage.ParametersAggregateShape_stagedOrder[instance]
	case *Position:
		return stage.Position_stagedOrder[instance]
	case *Repository:
		return stage.Repository_stagedOrder[instance]
	case *Scenario:
		return stage.Scenario_stagedOrder[instance]
	case *User:
		return stage.User_stagedOrder[instance]
	case *UserUse:
		return stage.UserUse_stagedOrder[instance]
	case *Workspace:
		return stage.Workspace_stagedOrder[instance]
	default:
		return 0 // should not happen
	}
}

func GongGetInstanceFromOrder[Type PointerToGongstruct](stage *Stage, order uint) (res Type) {
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

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *ActorState:
		return stage.ActorState_stagedOrder[instance]
	case *ActorStateShape:
		return stage.ActorStateShape_stagedOrder[instance]
	case *ActorStateTransition:
		return stage.ActorStateTransition_stagedOrder[instance]
	case *ActorStateTransitionShape:
		return stage.ActorStateTransitionShape_stagedOrder[instance]
	case *Analysis:
		return stage.Analysis_stagedOrder[instance]
	case *ControlPointShape:
		return stage.ControlPointShape_stagedOrder[instance]
	case *Diagram:
		return stage.Diagram_stagedOrder[instance]
	case *Document:
		return stage.Document_stagedOrder[instance]
	case *DocumentUse:
		return stage.DocumentUse_stagedOrder[instance]
	case *EvolutionDirection:
		return stage.EvolutionDirection_stagedOrder[instance]
	case *EvolutionDirectionShape:
		return stage.EvolutionDirectionShape_stagedOrder[instance]
	case *Foo:
		return stage.Foo_stagedOrder[instance]
	case *GeoObject:
		return stage.GeoObject_stagedOrder[instance]
	case *GeoObjectUse:
		return stage.GeoObjectUse_stagedOrder[instance]
	case *Group:
		return stage.Group_stagedOrder[instance]
	case *GroupUse:
		return stage.GroupUse_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *MapObject:
		return stage.MapObject_stagedOrder[instance]
	case *MapObjectUse:
		return stage.MapObjectUse_stagedOrder[instance]
	case *Parameter:
		return stage.Parameter_stagedOrder[instance]
	case *ParameterCategory:
		return stage.ParameterCategory_stagedOrder[instance]
	case *ParameterCategoryUse:
		return stage.ParameterCategoryUse_stagedOrder[instance]
	case *ParameterShape:
		return stage.ParameterShape_stagedOrder[instance]
	case *ParametersAggregate:
		return stage.ParametersAggregate_stagedOrder[instance]
	case *ParametersAggregateShape:
		return stage.ParametersAggregateShape_stagedOrder[instance]
	case *Position:
		return stage.Position_stagedOrder[instance]
	case *Repository:
		return stage.Repository_stagedOrder[instance]
	case *Scenario:
		return stage.Scenario_stagedOrder[instance]
	case *User:
		return stage.User_stagedOrder[instance]
	case *UserUse:
		return stage.UserUse_stagedOrder[instance]
	case *Workspace:
		return stage.Workspace_stagedOrder[instance]
	default:
		return 0 // should not happen
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
	if _, ok := stage.ActorStates[actorstate]; !ok {
		stage.ActorStates[actorstate] = struct{}{}
		stage.ActorState_stagedOrder[actorstate] = stage.ActorStateOrder
		stage.ActorState_orderStaged[stage.ActorStateOrder] = actorstate
		stage.ActorStateOrder++
	}
	stage.ActorStates_mapString[actorstate.Name] = actorstate

	return actorstate
}

// StagePreserveOrder puts actorstate to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ActorStateOrder
// - update stage.ActorStateOrder accordingly
func (actorstate *ActorState) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ActorStates[actorstate]; !ok {
		stage.ActorStates[actorstate] = struct{}{}

		if order > stage.ActorStateOrder {
			stage.ActorStateOrder = order
		}
		stage.ActorState_stagedOrder[actorstate] = order
		stage.ActorState_orderStaged[order] = actorstate
		stage.ActorStateOrder++
	}
	stage.ActorStates_mapString[actorstate.Name] = actorstate
}

// Unstage removes actorstate off the model stage
func (actorstate *ActorState) Unstage(stage *Stage) *ActorState {
	delete(stage.ActorStates, actorstate)
	// issue1150
	// delete(stage.ActorState_stagedOrder, actorstate)
	delete(stage.ActorStates_mapString, actorstate.Name)

	return actorstate
}

// UnstageVoid removes actorstate off the model stage
func (actorstate *ActorState) UnstageVoid(stage *Stage) {
	delete(stage.ActorStates, actorstate)
	// issue1150
	// delete(stage.ActorState_stagedOrder, actorstate)
	delete(stage.ActorStates_mapString, actorstate.Name)
}

// commit actorstate to the back repo (if it is already staged)
func (actorstate *ActorState) Commit(stage *Stage) *ActorState {
	if _, ok := stage.ActorStates[actorstate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitActorState(actorstate)
		}
	}
	return actorstate
}

func (actorstate *ActorState) CommitVoid(stage *Stage) {
	actorstate.Commit(stage)
}

func (actorstate *ActorState) StageVoid(stage *Stage) {
	actorstate.Stage(stage)
}

// Checkout actorstate to the back repo (if it is already staged)
func (actorstate *ActorState) Checkout(stage *Stage) *ActorState {
	if _, ok := stage.ActorStates[actorstate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutActorState(actorstate)
		}
	}
	return actorstate
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
	if _, ok := stage.ActorStateShapes[actorstateshape]; !ok {
		stage.ActorStateShapes[actorstateshape] = struct{}{}
		stage.ActorStateShape_stagedOrder[actorstateshape] = stage.ActorStateShapeOrder
		stage.ActorStateShape_orderStaged[stage.ActorStateShapeOrder] = actorstateshape
		stage.ActorStateShapeOrder++
	}
	stage.ActorStateShapes_mapString[actorstateshape.Name] = actorstateshape

	return actorstateshape
}

// StagePreserveOrder puts actorstateshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ActorStateShapeOrder
// - update stage.ActorStateShapeOrder accordingly
func (actorstateshape *ActorStateShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ActorStateShapes[actorstateshape]; !ok {
		stage.ActorStateShapes[actorstateshape] = struct{}{}

		if order > stage.ActorStateShapeOrder {
			stage.ActorStateShapeOrder = order
		}
		stage.ActorStateShape_stagedOrder[actorstateshape] = order
		stage.ActorStateShape_orderStaged[order] = actorstateshape
		stage.ActorStateShapeOrder++
	}
	stage.ActorStateShapes_mapString[actorstateshape.Name] = actorstateshape
}

// Unstage removes actorstateshape off the model stage
func (actorstateshape *ActorStateShape) Unstage(stage *Stage) *ActorStateShape {
	delete(stage.ActorStateShapes, actorstateshape)
	// issue1150
	// delete(stage.ActorStateShape_stagedOrder, actorstateshape)
	delete(stage.ActorStateShapes_mapString, actorstateshape.Name)

	return actorstateshape
}

// UnstageVoid removes actorstateshape off the model stage
func (actorstateshape *ActorStateShape) UnstageVoid(stage *Stage) {
	delete(stage.ActorStateShapes, actorstateshape)
	// issue1150
	// delete(stage.ActorStateShape_stagedOrder, actorstateshape)
	delete(stage.ActorStateShapes_mapString, actorstateshape.Name)
}

// commit actorstateshape to the back repo (if it is already staged)
func (actorstateshape *ActorStateShape) Commit(stage *Stage) *ActorStateShape {
	if _, ok := stage.ActorStateShapes[actorstateshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitActorStateShape(actorstateshape)
		}
	}
	return actorstateshape
}

func (actorstateshape *ActorStateShape) CommitVoid(stage *Stage) {
	actorstateshape.Commit(stage)
}

func (actorstateshape *ActorStateShape) StageVoid(stage *Stage) {
	actorstateshape.Stage(stage)
}

// Checkout actorstateshape to the back repo (if it is already staged)
func (actorstateshape *ActorStateShape) Checkout(stage *Stage) *ActorStateShape {
	if _, ok := stage.ActorStateShapes[actorstateshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutActorStateShape(actorstateshape)
		}
	}
	return actorstateshape
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
	if _, ok := stage.ActorStateTransitions[actorstatetransition]; !ok {
		stage.ActorStateTransitions[actorstatetransition] = struct{}{}
		stage.ActorStateTransition_stagedOrder[actorstatetransition] = stage.ActorStateTransitionOrder
		stage.ActorStateTransition_orderStaged[stage.ActorStateTransitionOrder] = actorstatetransition
		stage.ActorStateTransitionOrder++
	}
	stage.ActorStateTransitions_mapString[actorstatetransition.Name] = actorstatetransition

	return actorstatetransition
}

// StagePreserveOrder puts actorstatetransition to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ActorStateTransitionOrder
// - update stage.ActorStateTransitionOrder accordingly
func (actorstatetransition *ActorStateTransition) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ActorStateTransitions[actorstatetransition]; !ok {
		stage.ActorStateTransitions[actorstatetransition] = struct{}{}

		if order > stage.ActorStateTransitionOrder {
			stage.ActorStateTransitionOrder = order
		}
		stage.ActorStateTransition_stagedOrder[actorstatetransition] = order
		stage.ActorStateTransition_orderStaged[order] = actorstatetransition
		stage.ActorStateTransitionOrder++
	}
	stage.ActorStateTransitions_mapString[actorstatetransition.Name] = actorstatetransition
}

// Unstage removes actorstatetransition off the model stage
func (actorstatetransition *ActorStateTransition) Unstage(stage *Stage) *ActorStateTransition {
	delete(stage.ActorStateTransitions, actorstatetransition)
	// issue1150
	// delete(stage.ActorStateTransition_stagedOrder, actorstatetransition)
	delete(stage.ActorStateTransitions_mapString, actorstatetransition.Name)

	return actorstatetransition
}

// UnstageVoid removes actorstatetransition off the model stage
func (actorstatetransition *ActorStateTransition) UnstageVoid(stage *Stage) {
	delete(stage.ActorStateTransitions, actorstatetransition)
	// issue1150
	// delete(stage.ActorStateTransition_stagedOrder, actorstatetransition)
	delete(stage.ActorStateTransitions_mapString, actorstatetransition.Name)
}

// commit actorstatetransition to the back repo (if it is already staged)
func (actorstatetransition *ActorStateTransition) Commit(stage *Stage) *ActorStateTransition {
	if _, ok := stage.ActorStateTransitions[actorstatetransition]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitActorStateTransition(actorstatetransition)
		}
	}
	return actorstatetransition
}

func (actorstatetransition *ActorStateTransition) CommitVoid(stage *Stage) {
	actorstatetransition.Commit(stage)
}

func (actorstatetransition *ActorStateTransition) StageVoid(stage *Stage) {
	actorstatetransition.Stage(stage)
}

// Checkout actorstatetransition to the back repo (if it is already staged)
func (actorstatetransition *ActorStateTransition) Checkout(stage *Stage) *ActorStateTransition {
	if _, ok := stage.ActorStateTransitions[actorstatetransition]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutActorStateTransition(actorstatetransition)
		}
	}
	return actorstatetransition
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
	if _, ok := stage.ActorStateTransitionShapes[actorstatetransitionshape]; !ok {
		stage.ActorStateTransitionShapes[actorstatetransitionshape] = struct{}{}
		stage.ActorStateTransitionShape_stagedOrder[actorstatetransitionshape] = stage.ActorStateTransitionShapeOrder
		stage.ActorStateTransitionShape_orderStaged[stage.ActorStateTransitionShapeOrder] = actorstatetransitionshape
		stage.ActorStateTransitionShapeOrder++
	}
	stage.ActorStateTransitionShapes_mapString[actorstatetransitionshape.Name] = actorstatetransitionshape

	return actorstatetransitionshape
}

// StagePreserveOrder puts actorstatetransitionshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ActorStateTransitionShapeOrder
// - update stage.ActorStateTransitionShapeOrder accordingly
func (actorstatetransitionshape *ActorStateTransitionShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ActorStateTransitionShapes[actorstatetransitionshape]; !ok {
		stage.ActorStateTransitionShapes[actorstatetransitionshape] = struct{}{}

		if order > stage.ActorStateTransitionShapeOrder {
			stage.ActorStateTransitionShapeOrder = order
		}
		stage.ActorStateTransitionShape_stagedOrder[actorstatetransitionshape] = order
		stage.ActorStateTransitionShape_orderStaged[order] = actorstatetransitionshape
		stage.ActorStateTransitionShapeOrder++
	}
	stage.ActorStateTransitionShapes_mapString[actorstatetransitionshape.Name] = actorstatetransitionshape
}

// Unstage removes actorstatetransitionshape off the model stage
func (actorstatetransitionshape *ActorStateTransitionShape) Unstage(stage *Stage) *ActorStateTransitionShape {
	delete(stage.ActorStateTransitionShapes, actorstatetransitionshape)
	// issue1150
	// delete(stage.ActorStateTransitionShape_stagedOrder, actorstatetransitionshape)
	delete(stage.ActorStateTransitionShapes_mapString, actorstatetransitionshape.Name)

	return actorstatetransitionshape
}

// UnstageVoid removes actorstatetransitionshape off the model stage
func (actorstatetransitionshape *ActorStateTransitionShape) UnstageVoid(stage *Stage) {
	delete(stage.ActorStateTransitionShapes, actorstatetransitionshape)
	// issue1150
	// delete(stage.ActorStateTransitionShape_stagedOrder, actorstatetransitionshape)
	delete(stage.ActorStateTransitionShapes_mapString, actorstatetransitionshape.Name)
}

// commit actorstatetransitionshape to the back repo (if it is already staged)
func (actorstatetransitionshape *ActorStateTransitionShape) Commit(stage *Stage) *ActorStateTransitionShape {
	if _, ok := stage.ActorStateTransitionShapes[actorstatetransitionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitActorStateTransitionShape(actorstatetransitionshape)
		}
	}
	return actorstatetransitionshape
}

func (actorstatetransitionshape *ActorStateTransitionShape) CommitVoid(stage *Stage) {
	actorstatetransitionshape.Commit(stage)
}

func (actorstatetransitionshape *ActorStateTransitionShape) StageVoid(stage *Stage) {
	actorstatetransitionshape.Stage(stage)
}

// Checkout actorstatetransitionshape to the back repo (if it is already staged)
func (actorstatetransitionshape *ActorStateTransitionShape) Checkout(stage *Stage) *ActorStateTransitionShape {
	if _, ok := stage.ActorStateTransitionShapes[actorstatetransitionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutActorStateTransitionShape(actorstatetransitionshape)
		}
	}
	return actorstatetransitionshape
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
	if _, ok := stage.Analysiss[analysis]; !ok {
		stage.Analysiss[analysis] = struct{}{}
		stage.Analysis_stagedOrder[analysis] = stage.AnalysisOrder
		stage.Analysis_orderStaged[stage.AnalysisOrder] = analysis
		stage.AnalysisOrder++
	}
	stage.Analysiss_mapString[analysis.Name] = analysis

	return analysis
}

// StagePreserveOrder puts analysis to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AnalysisOrder
// - update stage.AnalysisOrder accordingly
func (analysis *Analysis) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Analysiss[analysis]; !ok {
		stage.Analysiss[analysis] = struct{}{}

		if order > stage.AnalysisOrder {
			stage.AnalysisOrder = order
		}
		stage.Analysis_stagedOrder[analysis] = order
		stage.Analysis_orderStaged[order] = analysis
		stage.AnalysisOrder++
	}
	stage.Analysiss_mapString[analysis.Name] = analysis
}

// Unstage removes analysis off the model stage
func (analysis *Analysis) Unstage(stage *Stage) *Analysis {
	delete(stage.Analysiss, analysis)
	// issue1150
	// delete(stage.Analysis_stagedOrder, analysis)
	delete(stage.Analysiss_mapString, analysis.Name)

	return analysis
}

// UnstageVoid removes analysis off the model stage
func (analysis *Analysis) UnstageVoid(stage *Stage) {
	delete(stage.Analysiss, analysis)
	// issue1150
	// delete(stage.Analysis_stagedOrder, analysis)
	delete(stage.Analysiss_mapString, analysis.Name)
}

// commit analysis to the back repo (if it is already staged)
func (analysis *Analysis) Commit(stage *Stage) *Analysis {
	if _, ok := stage.Analysiss[analysis]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAnalysis(analysis)
		}
	}
	return analysis
}

func (analysis *Analysis) CommitVoid(stage *Stage) {
	analysis.Commit(stage)
}

func (analysis *Analysis) StageVoid(stage *Stage) {
	analysis.Stage(stage)
}

// Checkout analysis to the back repo (if it is already staged)
func (analysis *Analysis) Checkout(stage *Stage) *Analysis {
	if _, ok := stage.Analysiss[analysis]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAnalysis(analysis)
		}
	}
	return analysis
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
	if _, ok := stage.ControlPointShapes[controlpointshape]; !ok {
		stage.ControlPointShapes[controlpointshape] = struct{}{}
		stage.ControlPointShape_stagedOrder[controlpointshape] = stage.ControlPointShapeOrder
		stage.ControlPointShape_orderStaged[stage.ControlPointShapeOrder] = controlpointshape
		stage.ControlPointShapeOrder++
	}
	stage.ControlPointShapes_mapString[controlpointshape.Name] = controlpointshape

	return controlpointshape
}

// StagePreserveOrder puts controlpointshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ControlPointShapeOrder
// - update stage.ControlPointShapeOrder accordingly
func (controlpointshape *ControlPointShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ControlPointShapes[controlpointshape]; !ok {
		stage.ControlPointShapes[controlpointshape] = struct{}{}

		if order > stage.ControlPointShapeOrder {
			stage.ControlPointShapeOrder = order
		}
		stage.ControlPointShape_stagedOrder[controlpointshape] = order
		stage.ControlPointShape_orderStaged[order] = controlpointshape
		stage.ControlPointShapeOrder++
	}
	stage.ControlPointShapes_mapString[controlpointshape.Name] = controlpointshape
}

// Unstage removes controlpointshape off the model stage
func (controlpointshape *ControlPointShape) Unstage(stage *Stage) *ControlPointShape {
	delete(stage.ControlPointShapes, controlpointshape)
	// issue1150
	// delete(stage.ControlPointShape_stagedOrder, controlpointshape)
	delete(stage.ControlPointShapes_mapString, controlpointshape.Name)

	return controlpointshape
}

// UnstageVoid removes controlpointshape off the model stage
func (controlpointshape *ControlPointShape) UnstageVoid(stage *Stage) {
	delete(stage.ControlPointShapes, controlpointshape)
	// issue1150
	// delete(stage.ControlPointShape_stagedOrder, controlpointshape)
	delete(stage.ControlPointShapes_mapString, controlpointshape.Name)
}

// commit controlpointshape to the back repo (if it is already staged)
func (controlpointshape *ControlPointShape) Commit(stage *Stage) *ControlPointShape {
	if _, ok := stage.ControlPointShapes[controlpointshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitControlPointShape(controlpointshape)
		}
	}
	return controlpointshape
}

func (controlpointshape *ControlPointShape) CommitVoid(stage *Stage) {
	controlpointshape.Commit(stage)
}

func (controlpointshape *ControlPointShape) StageVoid(stage *Stage) {
	controlpointshape.Stage(stage)
}

// Checkout controlpointshape to the back repo (if it is already staged)
func (controlpointshape *ControlPointShape) Checkout(stage *Stage) *ControlPointShape {
	if _, ok := stage.ControlPointShapes[controlpointshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutControlPointShape(controlpointshape)
		}
	}
	return controlpointshape
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
	if _, ok := stage.Diagrams[diagram]; !ok {
		stage.Diagrams[diagram] = struct{}{}
		stage.Diagram_stagedOrder[diagram] = stage.DiagramOrder
		stage.Diagram_orderStaged[stage.DiagramOrder] = diagram
		stage.DiagramOrder++
	}
	stage.Diagrams_mapString[diagram.Name] = diagram

	return diagram
}

// StagePreserveOrder puts diagram to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DiagramOrder
// - update stage.DiagramOrder accordingly
func (diagram *Diagram) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Diagrams[diagram]; !ok {
		stage.Diagrams[diagram] = struct{}{}

		if order > stage.DiagramOrder {
			stage.DiagramOrder = order
		}
		stage.Diagram_stagedOrder[diagram] = order
		stage.Diagram_orderStaged[order] = diagram
		stage.DiagramOrder++
	}
	stage.Diagrams_mapString[diagram.Name] = diagram
}

// Unstage removes diagram off the model stage
func (diagram *Diagram) Unstage(stage *Stage) *Diagram {
	delete(stage.Diagrams, diagram)
	// issue1150
	// delete(stage.Diagram_stagedOrder, diagram)
	delete(stage.Diagrams_mapString, diagram.Name)

	return diagram
}

// UnstageVoid removes diagram off the model stage
func (diagram *Diagram) UnstageVoid(stage *Stage) {
	delete(stage.Diagrams, diagram)
	// issue1150
	// delete(stage.Diagram_stagedOrder, diagram)
	delete(stage.Diagrams_mapString, diagram.Name)
}

// commit diagram to the back repo (if it is already staged)
func (diagram *Diagram) Commit(stage *Stage) *Diagram {
	if _, ok := stage.Diagrams[diagram]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDiagram(diagram)
		}
	}
	return diagram
}

func (diagram *Diagram) CommitVoid(stage *Stage) {
	diagram.Commit(stage)
}

func (diagram *Diagram) StageVoid(stage *Stage) {
	diagram.Stage(stage)
}

// Checkout diagram to the back repo (if it is already staged)
func (diagram *Diagram) Checkout(stage *Stage) *Diagram {
	if _, ok := stage.Diagrams[diagram]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDiagram(diagram)
		}
	}
	return diagram
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
	if _, ok := stage.Documents[document]; !ok {
		stage.Documents[document] = struct{}{}
		stage.Document_stagedOrder[document] = stage.DocumentOrder
		stage.Document_orderStaged[stage.DocumentOrder] = document
		stage.DocumentOrder++
	}
	stage.Documents_mapString[document.Name] = document

	return document
}

// StagePreserveOrder puts document to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DocumentOrder
// - update stage.DocumentOrder accordingly
func (document *Document) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Documents[document]; !ok {
		stage.Documents[document] = struct{}{}

		if order > stage.DocumentOrder {
			stage.DocumentOrder = order
		}
		stage.Document_stagedOrder[document] = order
		stage.Document_orderStaged[order] = document
		stage.DocumentOrder++
	}
	stage.Documents_mapString[document.Name] = document
}

// Unstage removes document off the model stage
func (document *Document) Unstage(stage *Stage) *Document {
	delete(stage.Documents, document)
	// issue1150
	// delete(stage.Document_stagedOrder, document)
	delete(stage.Documents_mapString, document.Name)

	return document
}

// UnstageVoid removes document off the model stage
func (document *Document) UnstageVoid(stage *Stage) {
	delete(stage.Documents, document)
	// issue1150
	// delete(stage.Document_stagedOrder, document)
	delete(stage.Documents_mapString, document.Name)
}

// commit document to the back repo (if it is already staged)
func (document *Document) Commit(stage *Stage) *Document {
	if _, ok := stage.Documents[document]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDocument(document)
		}
	}
	return document
}

func (document *Document) CommitVoid(stage *Stage) {
	document.Commit(stage)
}

func (document *Document) StageVoid(stage *Stage) {
	document.Stage(stage)
}

// Checkout document to the back repo (if it is already staged)
func (document *Document) Checkout(stage *Stage) *Document {
	if _, ok := stage.Documents[document]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDocument(document)
		}
	}
	return document
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
	if _, ok := stage.DocumentUses[documentuse]; !ok {
		stage.DocumentUses[documentuse] = struct{}{}
		stage.DocumentUse_stagedOrder[documentuse] = stage.DocumentUseOrder
		stage.DocumentUse_orderStaged[stage.DocumentUseOrder] = documentuse
		stage.DocumentUseOrder++
	}
	stage.DocumentUses_mapString[documentuse.Name] = documentuse

	return documentuse
}

// StagePreserveOrder puts documentuse to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DocumentUseOrder
// - update stage.DocumentUseOrder accordingly
func (documentuse *DocumentUse) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.DocumentUses[documentuse]; !ok {
		stage.DocumentUses[documentuse] = struct{}{}

		if order > stage.DocumentUseOrder {
			stage.DocumentUseOrder = order
		}
		stage.DocumentUse_stagedOrder[documentuse] = order
		stage.DocumentUse_orderStaged[order] = documentuse
		stage.DocumentUseOrder++
	}
	stage.DocumentUses_mapString[documentuse.Name] = documentuse
}

// Unstage removes documentuse off the model stage
func (documentuse *DocumentUse) Unstage(stage *Stage) *DocumentUse {
	delete(stage.DocumentUses, documentuse)
	// issue1150
	// delete(stage.DocumentUse_stagedOrder, documentuse)
	delete(stage.DocumentUses_mapString, documentuse.Name)

	return documentuse
}

// UnstageVoid removes documentuse off the model stage
func (documentuse *DocumentUse) UnstageVoid(stage *Stage) {
	delete(stage.DocumentUses, documentuse)
	// issue1150
	// delete(stage.DocumentUse_stagedOrder, documentuse)
	delete(stage.DocumentUses_mapString, documentuse.Name)
}

// commit documentuse to the back repo (if it is already staged)
func (documentuse *DocumentUse) Commit(stage *Stage) *DocumentUse {
	if _, ok := stage.DocumentUses[documentuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDocumentUse(documentuse)
		}
	}
	return documentuse
}

func (documentuse *DocumentUse) CommitVoid(stage *Stage) {
	documentuse.Commit(stage)
}

func (documentuse *DocumentUse) StageVoid(stage *Stage) {
	documentuse.Stage(stage)
}

// Checkout documentuse to the back repo (if it is already staged)
func (documentuse *DocumentUse) Checkout(stage *Stage) *DocumentUse {
	if _, ok := stage.DocumentUses[documentuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDocumentUse(documentuse)
		}
	}
	return documentuse
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
	if _, ok := stage.EvolutionDirections[evolutiondirection]; !ok {
		stage.EvolutionDirections[evolutiondirection] = struct{}{}
		stage.EvolutionDirection_stagedOrder[evolutiondirection] = stage.EvolutionDirectionOrder
		stage.EvolutionDirection_orderStaged[stage.EvolutionDirectionOrder] = evolutiondirection
		stage.EvolutionDirectionOrder++
	}
	stage.EvolutionDirections_mapString[evolutiondirection.Name] = evolutiondirection

	return evolutiondirection
}

// StagePreserveOrder puts evolutiondirection to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EvolutionDirectionOrder
// - update stage.EvolutionDirectionOrder accordingly
func (evolutiondirection *EvolutionDirection) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.EvolutionDirections[evolutiondirection]; !ok {
		stage.EvolutionDirections[evolutiondirection] = struct{}{}

		if order > stage.EvolutionDirectionOrder {
			stage.EvolutionDirectionOrder = order
		}
		stage.EvolutionDirection_stagedOrder[evolutiondirection] = order
		stage.EvolutionDirection_orderStaged[order] = evolutiondirection
		stage.EvolutionDirectionOrder++
	}
	stage.EvolutionDirections_mapString[evolutiondirection.Name] = evolutiondirection
}

// Unstage removes evolutiondirection off the model stage
func (evolutiondirection *EvolutionDirection) Unstage(stage *Stage) *EvolutionDirection {
	delete(stage.EvolutionDirections, evolutiondirection)
	// issue1150
	// delete(stage.EvolutionDirection_stagedOrder, evolutiondirection)
	delete(stage.EvolutionDirections_mapString, evolutiondirection.Name)

	return evolutiondirection
}

// UnstageVoid removes evolutiondirection off the model stage
func (evolutiondirection *EvolutionDirection) UnstageVoid(stage *Stage) {
	delete(stage.EvolutionDirections, evolutiondirection)
	// issue1150
	// delete(stage.EvolutionDirection_stagedOrder, evolutiondirection)
	delete(stage.EvolutionDirections_mapString, evolutiondirection.Name)
}

// commit evolutiondirection to the back repo (if it is already staged)
func (evolutiondirection *EvolutionDirection) Commit(stage *Stage) *EvolutionDirection {
	if _, ok := stage.EvolutionDirections[evolutiondirection]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitEvolutionDirection(evolutiondirection)
		}
	}
	return evolutiondirection
}

func (evolutiondirection *EvolutionDirection) CommitVoid(stage *Stage) {
	evolutiondirection.Commit(stage)
}

func (evolutiondirection *EvolutionDirection) StageVoid(stage *Stage) {
	evolutiondirection.Stage(stage)
}

// Checkout evolutiondirection to the back repo (if it is already staged)
func (evolutiondirection *EvolutionDirection) Checkout(stage *Stage) *EvolutionDirection {
	if _, ok := stage.EvolutionDirections[evolutiondirection]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutEvolutionDirection(evolutiondirection)
		}
	}
	return evolutiondirection
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
	if _, ok := stage.EvolutionDirectionShapes[evolutiondirectionshape]; !ok {
		stage.EvolutionDirectionShapes[evolutiondirectionshape] = struct{}{}
		stage.EvolutionDirectionShape_stagedOrder[evolutiondirectionshape] = stage.EvolutionDirectionShapeOrder
		stage.EvolutionDirectionShape_orderStaged[stage.EvolutionDirectionShapeOrder] = evolutiondirectionshape
		stage.EvolutionDirectionShapeOrder++
	}
	stage.EvolutionDirectionShapes_mapString[evolutiondirectionshape.Name] = evolutiondirectionshape

	return evolutiondirectionshape
}

// StagePreserveOrder puts evolutiondirectionshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EvolutionDirectionShapeOrder
// - update stage.EvolutionDirectionShapeOrder accordingly
func (evolutiondirectionshape *EvolutionDirectionShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.EvolutionDirectionShapes[evolutiondirectionshape]; !ok {
		stage.EvolutionDirectionShapes[evolutiondirectionshape] = struct{}{}

		if order > stage.EvolutionDirectionShapeOrder {
			stage.EvolutionDirectionShapeOrder = order
		}
		stage.EvolutionDirectionShape_stagedOrder[evolutiondirectionshape] = order
		stage.EvolutionDirectionShape_orderStaged[order] = evolutiondirectionshape
		stage.EvolutionDirectionShapeOrder++
	}
	stage.EvolutionDirectionShapes_mapString[evolutiondirectionshape.Name] = evolutiondirectionshape
}

// Unstage removes evolutiondirectionshape off the model stage
func (evolutiondirectionshape *EvolutionDirectionShape) Unstage(stage *Stage) *EvolutionDirectionShape {
	delete(stage.EvolutionDirectionShapes, evolutiondirectionshape)
	// issue1150
	// delete(stage.EvolutionDirectionShape_stagedOrder, evolutiondirectionshape)
	delete(stage.EvolutionDirectionShapes_mapString, evolutiondirectionshape.Name)

	return evolutiondirectionshape
}

// UnstageVoid removes evolutiondirectionshape off the model stage
func (evolutiondirectionshape *EvolutionDirectionShape) UnstageVoid(stage *Stage) {
	delete(stage.EvolutionDirectionShapes, evolutiondirectionshape)
	// issue1150
	// delete(stage.EvolutionDirectionShape_stagedOrder, evolutiondirectionshape)
	delete(stage.EvolutionDirectionShapes_mapString, evolutiondirectionshape.Name)
}

// commit evolutiondirectionshape to the back repo (if it is already staged)
func (evolutiondirectionshape *EvolutionDirectionShape) Commit(stage *Stage) *EvolutionDirectionShape {
	if _, ok := stage.EvolutionDirectionShapes[evolutiondirectionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitEvolutionDirectionShape(evolutiondirectionshape)
		}
	}
	return evolutiondirectionshape
}

func (evolutiondirectionshape *EvolutionDirectionShape) CommitVoid(stage *Stage) {
	evolutiondirectionshape.Commit(stage)
}

func (evolutiondirectionshape *EvolutionDirectionShape) StageVoid(stage *Stage) {
	evolutiondirectionshape.Stage(stage)
}

// Checkout evolutiondirectionshape to the back repo (if it is already staged)
func (evolutiondirectionshape *EvolutionDirectionShape) Checkout(stage *Stage) *EvolutionDirectionShape {
	if _, ok := stage.EvolutionDirectionShapes[evolutiondirectionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutEvolutionDirectionShape(evolutiondirectionshape)
		}
	}
	return evolutiondirectionshape
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
	if _, ok := stage.Foos[foo]; !ok {
		stage.Foos[foo] = struct{}{}
		stage.Foo_stagedOrder[foo] = stage.FooOrder
		stage.Foo_orderStaged[stage.FooOrder] = foo
		stage.FooOrder++
	}
	stage.Foos_mapString[foo.Name] = foo

	return foo
}

// StagePreserveOrder puts foo to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FooOrder
// - update stage.FooOrder accordingly
func (foo *Foo) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Foos[foo]; !ok {
		stage.Foos[foo] = struct{}{}

		if order > stage.FooOrder {
			stage.FooOrder = order
		}
		stage.Foo_stagedOrder[foo] = order
		stage.Foo_orderStaged[order] = foo
		stage.FooOrder++
	}
	stage.Foos_mapString[foo.Name] = foo
}

// Unstage removes foo off the model stage
func (foo *Foo) Unstage(stage *Stage) *Foo {
	delete(stage.Foos, foo)
	// issue1150
	// delete(stage.Foo_stagedOrder, foo)
	delete(stage.Foos_mapString, foo.Name)

	return foo
}

// UnstageVoid removes foo off the model stage
func (foo *Foo) UnstageVoid(stage *Stage) {
	delete(stage.Foos, foo)
	// issue1150
	// delete(stage.Foo_stagedOrder, foo)
	delete(stage.Foos_mapString, foo.Name)
}

// commit foo to the back repo (if it is already staged)
func (foo *Foo) Commit(stage *Stage) *Foo {
	if _, ok := stage.Foos[foo]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFoo(foo)
		}
	}
	return foo
}

func (foo *Foo) CommitVoid(stage *Stage) {
	foo.Commit(stage)
}

func (foo *Foo) StageVoid(stage *Stage) {
	foo.Stage(stage)
}

// Checkout foo to the back repo (if it is already staged)
func (foo *Foo) Checkout(stage *Stage) *Foo {
	if _, ok := stage.Foos[foo]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFoo(foo)
		}
	}
	return foo
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
	if _, ok := stage.GeoObjects[geoobject]; !ok {
		stage.GeoObjects[geoobject] = struct{}{}
		stage.GeoObject_stagedOrder[geoobject] = stage.GeoObjectOrder
		stage.GeoObject_orderStaged[stage.GeoObjectOrder] = geoobject
		stage.GeoObjectOrder++
	}
	stage.GeoObjects_mapString[geoobject.Name] = geoobject

	return geoobject
}

// StagePreserveOrder puts geoobject to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GeoObjectOrder
// - update stage.GeoObjectOrder accordingly
func (geoobject *GeoObject) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.GeoObjects[geoobject]; !ok {
		stage.GeoObjects[geoobject] = struct{}{}

		if order > stage.GeoObjectOrder {
			stage.GeoObjectOrder = order
		}
		stage.GeoObject_stagedOrder[geoobject] = order
		stage.GeoObject_orderStaged[order] = geoobject
		stage.GeoObjectOrder++
	}
	stage.GeoObjects_mapString[geoobject.Name] = geoobject
}

// Unstage removes geoobject off the model stage
func (geoobject *GeoObject) Unstage(stage *Stage) *GeoObject {
	delete(stage.GeoObjects, geoobject)
	// issue1150
	// delete(stage.GeoObject_stagedOrder, geoobject)
	delete(stage.GeoObjects_mapString, geoobject.Name)

	return geoobject
}

// UnstageVoid removes geoobject off the model stage
func (geoobject *GeoObject) UnstageVoid(stage *Stage) {
	delete(stage.GeoObjects, geoobject)
	// issue1150
	// delete(stage.GeoObject_stagedOrder, geoobject)
	delete(stage.GeoObjects_mapString, geoobject.Name)
}

// commit geoobject to the back repo (if it is already staged)
func (geoobject *GeoObject) Commit(stage *Stage) *GeoObject {
	if _, ok := stage.GeoObjects[geoobject]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGeoObject(geoobject)
		}
	}
	return geoobject
}

func (geoobject *GeoObject) CommitVoid(stage *Stage) {
	geoobject.Commit(stage)
}

func (geoobject *GeoObject) StageVoid(stage *Stage) {
	geoobject.Stage(stage)
}

// Checkout geoobject to the back repo (if it is already staged)
func (geoobject *GeoObject) Checkout(stage *Stage) *GeoObject {
	if _, ok := stage.GeoObjects[geoobject]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGeoObject(geoobject)
		}
	}
	return geoobject
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
	if _, ok := stage.GeoObjectUses[geoobjectuse]; !ok {
		stage.GeoObjectUses[geoobjectuse] = struct{}{}
		stage.GeoObjectUse_stagedOrder[geoobjectuse] = stage.GeoObjectUseOrder
		stage.GeoObjectUse_orderStaged[stage.GeoObjectUseOrder] = geoobjectuse
		stage.GeoObjectUseOrder++
	}
	stage.GeoObjectUses_mapString[geoobjectuse.Name] = geoobjectuse

	return geoobjectuse
}

// StagePreserveOrder puts geoobjectuse to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GeoObjectUseOrder
// - update stage.GeoObjectUseOrder accordingly
func (geoobjectuse *GeoObjectUse) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.GeoObjectUses[geoobjectuse]; !ok {
		stage.GeoObjectUses[geoobjectuse] = struct{}{}

		if order > stage.GeoObjectUseOrder {
			stage.GeoObjectUseOrder = order
		}
		stage.GeoObjectUse_stagedOrder[geoobjectuse] = order
		stage.GeoObjectUse_orderStaged[order] = geoobjectuse
		stage.GeoObjectUseOrder++
	}
	stage.GeoObjectUses_mapString[geoobjectuse.Name] = geoobjectuse
}

// Unstage removes geoobjectuse off the model stage
func (geoobjectuse *GeoObjectUse) Unstage(stage *Stage) *GeoObjectUse {
	delete(stage.GeoObjectUses, geoobjectuse)
	// issue1150
	// delete(stage.GeoObjectUse_stagedOrder, geoobjectuse)
	delete(stage.GeoObjectUses_mapString, geoobjectuse.Name)

	return geoobjectuse
}

// UnstageVoid removes geoobjectuse off the model stage
func (geoobjectuse *GeoObjectUse) UnstageVoid(stage *Stage) {
	delete(stage.GeoObjectUses, geoobjectuse)
	// issue1150
	// delete(stage.GeoObjectUse_stagedOrder, geoobjectuse)
	delete(stage.GeoObjectUses_mapString, geoobjectuse.Name)
}

// commit geoobjectuse to the back repo (if it is already staged)
func (geoobjectuse *GeoObjectUse) Commit(stage *Stage) *GeoObjectUse {
	if _, ok := stage.GeoObjectUses[geoobjectuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGeoObjectUse(geoobjectuse)
		}
	}
	return geoobjectuse
}

func (geoobjectuse *GeoObjectUse) CommitVoid(stage *Stage) {
	geoobjectuse.Commit(stage)
}

func (geoobjectuse *GeoObjectUse) StageVoid(stage *Stage) {
	geoobjectuse.Stage(stage)
}

// Checkout geoobjectuse to the back repo (if it is already staged)
func (geoobjectuse *GeoObjectUse) Checkout(stage *Stage) *GeoObjectUse {
	if _, ok := stage.GeoObjectUses[geoobjectuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGeoObjectUse(geoobjectuse)
		}
	}
	return geoobjectuse
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
	if _, ok := stage.Groups[group]; !ok {
		stage.Groups[group] = struct{}{}
		stage.Group_stagedOrder[group] = stage.GroupOrder
		stage.Group_orderStaged[stage.GroupOrder] = group
		stage.GroupOrder++
	}
	stage.Groups_mapString[group.Name] = group

	return group
}

// StagePreserveOrder puts group to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GroupOrder
// - update stage.GroupOrder accordingly
func (group *Group) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Groups[group]; !ok {
		stage.Groups[group] = struct{}{}

		if order > stage.GroupOrder {
			stage.GroupOrder = order
		}
		stage.Group_stagedOrder[group] = order
		stage.Group_orderStaged[order] = group
		stage.GroupOrder++
	}
	stage.Groups_mapString[group.Name] = group
}

// Unstage removes group off the model stage
func (group *Group) Unstage(stage *Stage) *Group {
	delete(stage.Groups, group)
	// issue1150
	// delete(stage.Group_stagedOrder, group)
	delete(stage.Groups_mapString, group.Name)

	return group
}

// UnstageVoid removes group off the model stage
func (group *Group) UnstageVoid(stage *Stage) {
	delete(stage.Groups, group)
	// issue1150
	// delete(stage.Group_stagedOrder, group)
	delete(stage.Groups_mapString, group.Name)
}

// commit group to the back repo (if it is already staged)
func (group *Group) Commit(stage *Stage) *Group {
	if _, ok := stage.Groups[group]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGroup(group)
		}
	}
	return group
}

func (group *Group) CommitVoid(stage *Stage) {
	group.Commit(stage)
}

func (group *Group) StageVoid(stage *Stage) {
	group.Stage(stage)
}

// Checkout group to the back repo (if it is already staged)
func (group *Group) Checkout(stage *Stage) *Group {
	if _, ok := stage.Groups[group]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGroup(group)
		}
	}
	return group
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
	if _, ok := stage.GroupUses[groupuse]; !ok {
		stage.GroupUses[groupuse] = struct{}{}
		stage.GroupUse_stagedOrder[groupuse] = stage.GroupUseOrder
		stage.GroupUse_orderStaged[stage.GroupUseOrder] = groupuse
		stage.GroupUseOrder++
	}
	stage.GroupUses_mapString[groupuse.Name] = groupuse

	return groupuse
}

// StagePreserveOrder puts groupuse to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GroupUseOrder
// - update stage.GroupUseOrder accordingly
func (groupuse *GroupUse) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.GroupUses[groupuse]; !ok {
		stage.GroupUses[groupuse] = struct{}{}

		if order > stage.GroupUseOrder {
			stage.GroupUseOrder = order
		}
		stage.GroupUse_stagedOrder[groupuse] = order
		stage.GroupUse_orderStaged[order] = groupuse
		stage.GroupUseOrder++
	}
	stage.GroupUses_mapString[groupuse.Name] = groupuse
}

// Unstage removes groupuse off the model stage
func (groupuse *GroupUse) Unstage(stage *Stage) *GroupUse {
	delete(stage.GroupUses, groupuse)
	// issue1150
	// delete(stage.GroupUse_stagedOrder, groupuse)
	delete(stage.GroupUses_mapString, groupuse.Name)

	return groupuse
}

// UnstageVoid removes groupuse off the model stage
func (groupuse *GroupUse) UnstageVoid(stage *Stage) {
	delete(stage.GroupUses, groupuse)
	// issue1150
	// delete(stage.GroupUse_stagedOrder, groupuse)
	delete(stage.GroupUses_mapString, groupuse.Name)
}

// commit groupuse to the back repo (if it is already staged)
func (groupuse *GroupUse) Commit(stage *Stage) *GroupUse {
	if _, ok := stage.GroupUses[groupuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGroupUse(groupuse)
		}
	}
	return groupuse
}

func (groupuse *GroupUse) CommitVoid(stage *Stage) {
	groupuse.Commit(stage)
}

func (groupuse *GroupUse) StageVoid(stage *Stage) {
	groupuse.Stage(stage)
}

// Checkout groupuse to the back repo (if it is already staged)
func (groupuse *GroupUse) Checkout(stage *Stage) *GroupUse {
	if _, ok := stage.GroupUses[groupuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGroupUse(groupuse)
		}
	}
	return groupuse
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
	if _, ok := stage.Librarys[library]; !ok {
		stage.Librarys[library] = struct{}{}
		stage.Library_stagedOrder[library] = stage.LibraryOrder
		stage.Library_orderStaged[stage.LibraryOrder] = library
		stage.LibraryOrder++
	}
	stage.Librarys_mapString[library.Name] = library

	return library
}

// StagePreserveOrder puts library to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LibraryOrder
// - update stage.LibraryOrder accordingly
func (library *Library) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Librarys[library]; !ok {
		stage.Librarys[library] = struct{}{}

		if order > stage.LibraryOrder {
			stage.LibraryOrder = order
		}
		stage.Library_stagedOrder[library] = order
		stage.Library_orderStaged[order] = library
		stage.LibraryOrder++
	}
	stage.Librarys_mapString[library.Name] = library
}

// Unstage removes library off the model stage
func (library *Library) Unstage(stage *Stage) *Library {
	delete(stage.Librarys, library)
	// issue1150
	// delete(stage.Library_stagedOrder, library)
	delete(stage.Librarys_mapString, library.Name)

	return library
}

// UnstageVoid removes library off the model stage
func (library *Library) UnstageVoid(stage *Stage) {
	delete(stage.Librarys, library)
	// issue1150
	// delete(stage.Library_stagedOrder, library)
	delete(stage.Librarys_mapString, library.Name)
}

// commit library to the back repo (if it is already staged)
func (library *Library) Commit(stage *Stage) *Library {
	if _, ok := stage.Librarys[library]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLibrary(library)
		}
	}
	return library
}

func (library *Library) CommitVoid(stage *Stage) {
	library.Commit(stage)
}

func (library *Library) StageVoid(stage *Stage) {
	library.Stage(stage)
}

// Checkout library to the back repo (if it is already staged)
func (library *Library) Checkout(stage *Stage) *Library {
	if _, ok := stage.Librarys[library]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLibrary(library)
		}
	}
	return library
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
	if _, ok := stage.MapObjects[mapobject]; !ok {
		stage.MapObjects[mapobject] = struct{}{}
		stage.MapObject_stagedOrder[mapobject] = stage.MapObjectOrder
		stage.MapObject_orderStaged[stage.MapObjectOrder] = mapobject
		stage.MapObjectOrder++
	}
	stage.MapObjects_mapString[mapobject.Name] = mapobject

	return mapobject
}

// StagePreserveOrder puts mapobject to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MapObjectOrder
// - update stage.MapObjectOrder accordingly
func (mapobject *MapObject) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.MapObjects[mapobject]; !ok {
		stage.MapObjects[mapobject] = struct{}{}

		if order > stage.MapObjectOrder {
			stage.MapObjectOrder = order
		}
		stage.MapObject_stagedOrder[mapobject] = order
		stage.MapObject_orderStaged[order] = mapobject
		stage.MapObjectOrder++
	}
	stage.MapObjects_mapString[mapobject.Name] = mapobject
}

// Unstage removes mapobject off the model stage
func (mapobject *MapObject) Unstage(stage *Stage) *MapObject {
	delete(stage.MapObjects, mapobject)
	// issue1150
	// delete(stage.MapObject_stagedOrder, mapobject)
	delete(stage.MapObjects_mapString, mapobject.Name)

	return mapobject
}

// UnstageVoid removes mapobject off the model stage
func (mapobject *MapObject) UnstageVoid(stage *Stage) {
	delete(stage.MapObjects, mapobject)
	// issue1150
	// delete(stage.MapObject_stagedOrder, mapobject)
	delete(stage.MapObjects_mapString, mapobject.Name)
}

// commit mapobject to the back repo (if it is already staged)
func (mapobject *MapObject) Commit(stage *Stage) *MapObject {
	if _, ok := stage.MapObjects[mapobject]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMapObject(mapobject)
		}
	}
	return mapobject
}

func (mapobject *MapObject) CommitVoid(stage *Stage) {
	mapobject.Commit(stage)
}

func (mapobject *MapObject) StageVoid(stage *Stage) {
	mapobject.Stage(stage)
}

// Checkout mapobject to the back repo (if it is already staged)
func (mapobject *MapObject) Checkout(stage *Stage) *MapObject {
	if _, ok := stage.MapObjects[mapobject]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMapObject(mapobject)
		}
	}
	return mapobject
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
	if _, ok := stage.MapObjectUses[mapobjectuse]; !ok {
		stage.MapObjectUses[mapobjectuse] = struct{}{}
		stage.MapObjectUse_stagedOrder[mapobjectuse] = stage.MapObjectUseOrder
		stage.MapObjectUse_orderStaged[stage.MapObjectUseOrder] = mapobjectuse
		stage.MapObjectUseOrder++
	}
	stage.MapObjectUses_mapString[mapobjectuse.Name] = mapobjectuse

	return mapobjectuse
}

// StagePreserveOrder puts mapobjectuse to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MapObjectUseOrder
// - update stage.MapObjectUseOrder accordingly
func (mapobjectuse *MapObjectUse) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.MapObjectUses[mapobjectuse]; !ok {
		stage.MapObjectUses[mapobjectuse] = struct{}{}

		if order > stage.MapObjectUseOrder {
			stage.MapObjectUseOrder = order
		}
		stage.MapObjectUse_stagedOrder[mapobjectuse] = order
		stage.MapObjectUse_orderStaged[order] = mapobjectuse
		stage.MapObjectUseOrder++
	}
	stage.MapObjectUses_mapString[mapobjectuse.Name] = mapobjectuse
}

// Unstage removes mapobjectuse off the model stage
func (mapobjectuse *MapObjectUse) Unstage(stage *Stage) *MapObjectUse {
	delete(stage.MapObjectUses, mapobjectuse)
	// issue1150
	// delete(stage.MapObjectUse_stagedOrder, mapobjectuse)
	delete(stage.MapObjectUses_mapString, mapobjectuse.Name)

	return mapobjectuse
}

// UnstageVoid removes mapobjectuse off the model stage
func (mapobjectuse *MapObjectUse) UnstageVoid(stage *Stage) {
	delete(stage.MapObjectUses, mapobjectuse)
	// issue1150
	// delete(stage.MapObjectUse_stagedOrder, mapobjectuse)
	delete(stage.MapObjectUses_mapString, mapobjectuse.Name)
}

// commit mapobjectuse to the back repo (if it is already staged)
func (mapobjectuse *MapObjectUse) Commit(stage *Stage) *MapObjectUse {
	if _, ok := stage.MapObjectUses[mapobjectuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMapObjectUse(mapobjectuse)
		}
	}
	return mapobjectuse
}

func (mapobjectuse *MapObjectUse) CommitVoid(stage *Stage) {
	mapobjectuse.Commit(stage)
}

func (mapobjectuse *MapObjectUse) StageVoid(stage *Stage) {
	mapobjectuse.Stage(stage)
}

// Checkout mapobjectuse to the back repo (if it is already staged)
func (mapobjectuse *MapObjectUse) Checkout(stage *Stage) *MapObjectUse {
	if _, ok := stage.MapObjectUses[mapobjectuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMapObjectUse(mapobjectuse)
		}
	}
	return mapobjectuse
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
	if _, ok := stage.Parameters[parameter]; !ok {
		stage.Parameters[parameter] = struct{}{}
		stage.Parameter_stagedOrder[parameter] = stage.ParameterOrder
		stage.Parameter_orderStaged[stage.ParameterOrder] = parameter
		stage.ParameterOrder++
	}
	stage.Parameters_mapString[parameter.Name] = parameter

	return parameter
}

// StagePreserveOrder puts parameter to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ParameterOrder
// - update stage.ParameterOrder accordingly
func (parameter *Parameter) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Parameters[parameter]; !ok {
		stage.Parameters[parameter] = struct{}{}

		if order > stage.ParameterOrder {
			stage.ParameterOrder = order
		}
		stage.Parameter_stagedOrder[parameter] = order
		stage.Parameter_orderStaged[order] = parameter
		stage.ParameterOrder++
	}
	stage.Parameters_mapString[parameter.Name] = parameter
}

// Unstage removes parameter off the model stage
func (parameter *Parameter) Unstage(stage *Stage) *Parameter {
	delete(stage.Parameters, parameter)
	// issue1150
	// delete(stage.Parameter_stagedOrder, parameter)
	delete(stage.Parameters_mapString, parameter.Name)

	return parameter
}

// UnstageVoid removes parameter off the model stage
func (parameter *Parameter) UnstageVoid(stage *Stage) {
	delete(stage.Parameters, parameter)
	// issue1150
	// delete(stage.Parameter_stagedOrder, parameter)
	delete(stage.Parameters_mapString, parameter.Name)
}

// commit parameter to the back repo (if it is already staged)
func (parameter *Parameter) Commit(stage *Stage) *Parameter {
	if _, ok := stage.Parameters[parameter]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitParameter(parameter)
		}
	}
	return parameter
}

func (parameter *Parameter) CommitVoid(stage *Stage) {
	parameter.Commit(stage)
}

func (parameter *Parameter) StageVoid(stage *Stage) {
	parameter.Stage(stage)
}

// Checkout parameter to the back repo (if it is already staged)
func (parameter *Parameter) Checkout(stage *Stage) *Parameter {
	if _, ok := stage.Parameters[parameter]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutParameter(parameter)
		}
	}
	return parameter
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
	if _, ok := stage.ParameterCategorys[parametercategory]; !ok {
		stage.ParameterCategorys[parametercategory] = struct{}{}
		stage.ParameterCategory_stagedOrder[parametercategory] = stage.ParameterCategoryOrder
		stage.ParameterCategory_orderStaged[stage.ParameterCategoryOrder] = parametercategory
		stage.ParameterCategoryOrder++
	}
	stage.ParameterCategorys_mapString[parametercategory.Name] = parametercategory

	return parametercategory
}

// StagePreserveOrder puts parametercategory to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ParameterCategoryOrder
// - update stage.ParameterCategoryOrder accordingly
func (parametercategory *ParameterCategory) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ParameterCategorys[parametercategory]; !ok {
		stage.ParameterCategorys[parametercategory] = struct{}{}

		if order > stage.ParameterCategoryOrder {
			stage.ParameterCategoryOrder = order
		}
		stage.ParameterCategory_stagedOrder[parametercategory] = order
		stage.ParameterCategory_orderStaged[order] = parametercategory
		stage.ParameterCategoryOrder++
	}
	stage.ParameterCategorys_mapString[parametercategory.Name] = parametercategory
}

// Unstage removes parametercategory off the model stage
func (parametercategory *ParameterCategory) Unstage(stage *Stage) *ParameterCategory {
	delete(stage.ParameterCategorys, parametercategory)
	// issue1150
	// delete(stage.ParameterCategory_stagedOrder, parametercategory)
	delete(stage.ParameterCategorys_mapString, parametercategory.Name)

	return parametercategory
}

// UnstageVoid removes parametercategory off the model stage
func (parametercategory *ParameterCategory) UnstageVoid(stage *Stage) {
	delete(stage.ParameterCategorys, parametercategory)
	// issue1150
	// delete(stage.ParameterCategory_stagedOrder, parametercategory)
	delete(stage.ParameterCategorys_mapString, parametercategory.Name)
}

// commit parametercategory to the back repo (if it is already staged)
func (parametercategory *ParameterCategory) Commit(stage *Stage) *ParameterCategory {
	if _, ok := stage.ParameterCategorys[parametercategory]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitParameterCategory(parametercategory)
		}
	}
	return parametercategory
}

func (parametercategory *ParameterCategory) CommitVoid(stage *Stage) {
	parametercategory.Commit(stage)
}

func (parametercategory *ParameterCategory) StageVoid(stage *Stage) {
	parametercategory.Stage(stage)
}

// Checkout parametercategory to the back repo (if it is already staged)
func (parametercategory *ParameterCategory) Checkout(stage *Stage) *ParameterCategory {
	if _, ok := stage.ParameterCategorys[parametercategory]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutParameterCategory(parametercategory)
		}
	}
	return parametercategory
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
	if _, ok := stage.ParameterCategoryUses[parametercategoryuse]; !ok {
		stage.ParameterCategoryUses[parametercategoryuse] = struct{}{}
		stage.ParameterCategoryUse_stagedOrder[parametercategoryuse] = stage.ParameterCategoryUseOrder
		stage.ParameterCategoryUse_orderStaged[stage.ParameterCategoryUseOrder] = parametercategoryuse
		stage.ParameterCategoryUseOrder++
	}
	stage.ParameterCategoryUses_mapString[parametercategoryuse.Name] = parametercategoryuse

	return parametercategoryuse
}

// StagePreserveOrder puts parametercategoryuse to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ParameterCategoryUseOrder
// - update stage.ParameterCategoryUseOrder accordingly
func (parametercategoryuse *ParameterCategoryUse) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ParameterCategoryUses[parametercategoryuse]; !ok {
		stage.ParameterCategoryUses[parametercategoryuse] = struct{}{}

		if order > stage.ParameterCategoryUseOrder {
			stage.ParameterCategoryUseOrder = order
		}
		stage.ParameterCategoryUse_stagedOrder[parametercategoryuse] = order
		stage.ParameterCategoryUse_orderStaged[order] = parametercategoryuse
		stage.ParameterCategoryUseOrder++
	}
	stage.ParameterCategoryUses_mapString[parametercategoryuse.Name] = parametercategoryuse
}

// Unstage removes parametercategoryuse off the model stage
func (parametercategoryuse *ParameterCategoryUse) Unstage(stage *Stage) *ParameterCategoryUse {
	delete(stage.ParameterCategoryUses, parametercategoryuse)
	// issue1150
	// delete(stage.ParameterCategoryUse_stagedOrder, parametercategoryuse)
	delete(stage.ParameterCategoryUses_mapString, parametercategoryuse.Name)

	return parametercategoryuse
}

// UnstageVoid removes parametercategoryuse off the model stage
func (parametercategoryuse *ParameterCategoryUse) UnstageVoid(stage *Stage) {
	delete(stage.ParameterCategoryUses, parametercategoryuse)
	// issue1150
	// delete(stage.ParameterCategoryUse_stagedOrder, parametercategoryuse)
	delete(stage.ParameterCategoryUses_mapString, parametercategoryuse.Name)
}

// commit parametercategoryuse to the back repo (if it is already staged)
func (parametercategoryuse *ParameterCategoryUse) Commit(stage *Stage) *ParameterCategoryUse {
	if _, ok := stage.ParameterCategoryUses[parametercategoryuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitParameterCategoryUse(parametercategoryuse)
		}
	}
	return parametercategoryuse
}

func (parametercategoryuse *ParameterCategoryUse) CommitVoid(stage *Stage) {
	parametercategoryuse.Commit(stage)
}

func (parametercategoryuse *ParameterCategoryUse) StageVoid(stage *Stage) {
	parametercategoryuse.Stage(stage)
}

// Checkout parametercategoryuse to the back repo (if it is already staged)
func (parametercategoryuse *ParameterCategoryUse) Checkout(stage *Stage) *ParameterCategoryUse {
	if _, ok := stage.ParameterCategoryUses[parametercategoryuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutParameterCategoryUse(parametercategoryuse)
		}
	}
	return parametercategoryuse
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
	if _, ok := stage.ParameterShapes[parametershape]; !ok {
		stage.ParameterShapes[parametershape] = struct{}{}
		stage.ParameterShape_stagedOrder[parametershape] = stage.ParameterShapeOrder
		stage.ParameterShape_orderStaged[stage.ParameterShapeOrder] = parametershape
		stage.ParameterShapeOrder++
	}
	stage.ParameterShapes_mapString[parametershape.Name] = parametershape

	return parametershape
}

// StagePreserveOrder puts parametershape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ParameterShapeOrder
// - update stage.ParameterShapeOrder accordingly
func (parametershape *ParameterShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ParameterShapes[parametershape]; !ok {
		stage.ParameterShapes[parametershape] = struct{}{}

		if order > stage.ParameterShapeOrder {
			stage.ParameterShapeOrder = order
		}
		stage.ParameterShape_stagedOrder[parametershape] = order
		stage.ParameterShape_orderStaged[order] = parametershape
		stage.ParameterShapeOrder++
	}
	stage.ParameterShapes_mapString[parametershape.Name] = parametershape
}

// Unstage removes parametershape off the model stage
func (parametershape *ParameterShape) Unstage(stage *Stage) *ParameterShape {
	delete(stage.ParameterShapes, parametershape)
	// issue1150
	// delete(stage.ParameterShape_stagedOrder, parametershape)
	delete(stage.ParameterShapes_mapString, parametershape.Name)

	return parametershape
}

// UnstageVoid removes parametershape off the model stage
func (parametershape *ParameterShape) UnstageVoid(stage *Stage) {
	delete(stage.ParameterShapes, parametershape)
	// issue1150
	// delete(stage.ParameterShape_stagedOrder, parametershape)
	delete(stage.ParameterShapes_mapString, parametershape.Name)
}

// commit parametershape to the back repo (if it is already staged)
func (parametershape *ParameterShape) Commit(stage *Stage) *ParameterShape {
	if _, ok := stage.ParameterShapes[parametershape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitParameterShape(parametershape)
		}
	}
	return parametershape
}

func (parametershape *ParameterShape) CommitVoid(stage *Stage) {
	parametershape.Commit(stage)
}

func (parametershape *ParameterShape) StageVoid(stage *Stage) {
	parametershape.Stage(stage)
}

// Checkout parametershape to the back repo (if it is already staged)
func (parametershape *ParameterShape) Checkout(stage *Stage) *ParameterShape {
	if _, ok := stage.ParameterShapes[parametershape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutParameterShape(parametershape)
		}
	}
	return parametershape
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
	if _, ok := stage.ParametersAggregates[parametersaggregate]; !ok {
		stage.ParametersAggregates[parametersaggregate] = struct{}{}
		stage.ParametersAggregate_stagedOrder[parametersaggregate] = stage.ParametersAggregateOrder
		stage.ParametersAggregate_orderStaged[stage.ParametersAggregateOrder] = parametersaggregate
		stage.ParametersAggregateOrder++
	}
	stage.ParametersAggregates_mapString[parametersaggregate.Name] = parametersaggregate

	return parametersaggregate
}

// StagePreserveOrder puts parametersaggregate to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ParametersAggregateOrder
// - update stage.ParametersAggregateOrder accordingly
func (parametersaggregate *ParametersAggregate) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ParametersAggregates[parametersaggregate]; !ok {
		stage.ParametersAggregates[parametersaggregate] = struct{}{}

		if order > stage.ParametersAggregateOrder {
			stage.ParametersAggregateOrder = order
		}
		stage.ParametersAggregate_stagedOrder[parametersaggregate] = order
		stage.ParametersAggregate_orderStaged[order] = parametersaggregate
		stage.ParametersAggregateOrder++
	}
	stage.ParametersAggregates_mapString[parametersaggregate.Name] = parametersaggregate
}

// Unstage removes parametersaggregate off the model stage
func (parametersaggregate *ParametersAggregate) Unstage(stage *Stage) *ParametersAggregate {
	delete(stage.ParametersAggregates, parametersaggregate)
	// issue1150
	// delete(stage.ParametersAggregate_stagedOrder, parametersaggregate)
	delete(stage.ParametersAggregates_mapString, parametersaggregate.Name)

	return parametersaggregate
}

// UnstageVoid removes parametersaggregate off the model stage
func (parametersaggregate *ParametersAggregate) UnstageVoid(stage *Stage) {
	delete(stage.ParametersAggregates, parametersaggregate)
	// issue1150
	// delete(stage.ParametersAggregate_stagedOrder, parametersaggregate)
	delete(stage.ParametersAggregates_mapString, parametersaggregate.Name)
}

// commit parametersaggregate to the back repo (if it is already staged)
func (parametersaggregate *ParametersAggregate) Commit(stage *Stage) *ParametersAggregate {
	if _, ok := stage.ParametersAggregates[parametersaggregate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitParametersAggregate(parametersaggregate)
		}
	}
	return parametersaggregate
}

func (parametersaggregate *ParametersAggregate) CommitVoid(stage *Stage) {
	parametersaggregate.Commit(stage)
}

func (parametersaggregate *ParametersAggregate) StageVoid(stage *Stage) {
	parametersaggregate.Stage(stage)
}

// Checkout parametersaggregate to the back repo (if it is already staged)
func (parametersaggregate *ParametersAggregate) Checkout(stage *Stage) *ParametersAggregate {
	if _, ok := stage.ParametersAggregates[parametersaggregate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutParametersAggregate(parametersaggregate)
		}
	}
	return parametersaggregate
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
	if _, ok := stage.ParametersAggregateShapes[parametersaggregateshape]; !ok {
		stage.ParametersAggregateShapes[parametersaggregateshape] = struct{}{}
		stage.ParametersAggregateShape_stagedOrder[parametersaggregateshape] = stage.ParametersAggregateShapeOrder
		stage.ParametersAggregateShape_orderStaged[stage.ParametersAggregateShapeOrder] = parametersaggregateshape
		stage.ParametersAggregateShapeOrder++
	}
	stage.ParametersAggregateShapes_mapString[parametersaggregateshape.Name] = parametersaggregateshape

	return parametersaggregateshape
}

// StagePreserveOrder puts parametersaggregateshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ParametersAggregateShapeOrder
// - update stage.ParametersAggregateShapeOrder accordingly
func (parametersaggregateshape *ParametersAggregateShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ParametersAggregateShapes[parametersaggregateshape]; !ok {
		stage.ParametersAggregateShapes[parametersaggregateshape] = struct{}{}

		if order > stage.ParametersAggregateShapeOrder {
			stage.ParametersAggregateShapeOrder = order
		}
		stage.ParametersAggregateShape_stagedOrder[parametersaggregateshape] = order
		stage.ParametersAggregateShape_orderStaged[order] = parametersaggregateshape
		stage.ParametersAggregateShapeOrder++
	}
	stage.ParametersAggregateShapes_mapString[parametersaggregateshape.Name] = parametersaggregateshape
}

// Unstage removes parametersaggregateshape off the model stage
func (parametersaggregateshape *ParametersAggregateShape) Unstage(stage *Stage) *ParametersAggregateShape {
	delete(stage.ParametersAggregateShapes, parametersaggregateshape)
	// issue1150
	// delete(stage.ParametersAggregateShape_stagedOrder, parametersaggregateshape)
	delete(stage.ParametersAggregateShapes_mapString, parametersaggregateshape.Name)

	return parametersaggregateshape
}

// UnstageVoid removes parametersaggregateshape off the model stage
func (parametersaggregateshape *ParametersAggregateShape) UnstageVoid(stage *Stage) {
	delete(stage.ParametersAggregateShapes, parametersaggregateshape)
	// issue1150
	// delete(stage.ParametersAggregateShape_stagedOrder, parametersaggregateshape)
	delete(stage.ParametersAggregateShapes_mapString, parametersaggregateshape.Name)
}

// commit parametersaggregateshape to the back repo (if it is already staged)
func (parametersaggregateshape *ParametersAggregateShape) Commit(stage *Stage) *ParametersAggregateShape {
	if _, ok := stage.ParametersAggregateShapes[parametersaggregateshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitParametersAggregateShape(parametersaggregateshape)
		}
	}
	return parametersaggregateshape
}

func (parametersaggregateshape *ParametersAggregateShape) CommitVoid(stage *Stage) {
	parametersaggregateshape.Commit(stage)
}

func (parametersaggregateshape *ParametersAggregateShape) StageVoid(stage *Stage) {
	parametersaggregateshape.Stage(stage)
}

// Checkout parametersaggregateshape to the back repo (if it is already staged)
func (parametersaggregateshape *ParametersAggregateShape) Checkout(stage *Stage) *ParametersAggregateShape {
	if _, ok := stage.ParametersAggregateShapes[parametersaggregateshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutParametersAggregateShape(parametersaggregateshape)
		}
	}
	return parametersaggregateshape
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
	if _, ok := stage.Positions[position]; !ok {
		stage.Positions[position] = struct{}{}
		stage.Position_stagedOrder[position] = stage.PositionOrder
		stage.Position_orderStaged[stage.PositionOrder] = position
		stage.PositionOrder++
	}
	stage.Positions_mapString[position.Name] = position

	return position
}

// StagePreserveOrder puts position to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PositionOrder
// - update stage.PositionOrder accordingly
func (position *Position) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Positions[position]; !ok {
		stage.Positions[position] = struct{}{}

		if order > stage.PositionOrder {
			stage.PositionOrder = order
		}
		stage.Position_stagedOrder[position] = order
		stage.Position_orderStaged[order] = position
		stage.PositionOrder++
	}
	stage.Positions_mapString[position.Name] = position
}

// Unstage removes position off the model stage
func (position *Position) Unstage(stage *Stage) *Position {
	delete(stage.Positions, position)
	// issue1150
	// delete(stage.Position_stagedOrder, position)
	delete(stage.Positions_mapString, position.Name)

	return position
}

// UnstageVoid removes position off the model stage
func (position *Position) UnstageVoid(stage *Stage) {
	delete(stage.Positions, position)
	// issue1150
	// delete(stage.Position_stagedOrder, position)
	delete(stage.Positions_mapString, position.Name)
}

// commit position to the back repo (if it is already staged)
func (position *Position) Commit(stage *Stage) *Position {
	if _, ok := stage.Positions[position]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPosition(position)
		}
	}
	return position
}

func (position *Position) CommitVoid(stage *Stage) {
	position.Commit(stage)
}

func (position *Position) StageVoid(stage *Stage) {
	position.Stage(stage)
}

// Checkout position to the back repo (if it is already staged)
func (position *Position) Checkout(stage *Stage) *Position {
	if _, ok := stage.Positions[position]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPosition(position)
		}
	}
	return position
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
	if _, ok := stage.Repositorys[repository]; !ok {
		stage.Repositorys[repository] = struct{}{}
		stage.Repository_stagedOrder[repository] = stage.RepositoryOrder
		stage.Repository_orderStaged[stage.RepositoryOrder] = repository
		stage.RepositoryOrder++
	}
	stage.Repositorys_mapString[repository.Name] = repository

	return repository
}

// StagePreserveOrder puts repository to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RepositoryOrder
// - update stage.RepositoryOrder accordingly
func (repository *Repository) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Repositorys[repository]; !ok {
		stage.Repositorys[repository] = struct{}{}

		if order > stage.RepositoryOrder {
			stage.RepositoryOrder = order
		}
		stage.Repository_stagedOrder[repository] = order
		stage.Repository_orderStaged[order] = repository
		stage.RepositoryOrder++
	}
	stage.Repositorys_mapString[repository.Name] = repository
}

// Unstage removes repository off the model stage
func (repository *Repository) Unstage(stage *Stage) *Repository {
	delete(stage.Repositorys, repository)
	// issue1150
	// delete(stage.Repository_stagedOrder, repository)
	delete(stage.Repositorys_mapString, repository.Name)

	return repository
}

// UnstageVoid removes repository off the model stage
func (repository *Repository) UnstageVoid(stage *Stage) {
	delete(stage.Repositorys, repository)
	// issue1150
	// delete(stage.Repository_stagedOrder, repository)
	delete(stage.Repositorys_mapString, repository.Name)
}

// commit repository to the back repo (if it is already staged)
func (repository *Repository) Commit(stage *Stage) *Repository {
	if _, ok := stage.Repositorys[repository]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRepository(repository)
		}
	}
	return repository
}

func (repository *Repository) CommitVoid(stage *Stage) {
	repository.Commit(stage)
}

func (repository *Repository) StageVoid(stage *Stage) {
	repository.Stage(stage)
}

// Checkout repository to the back repo (if it is already staged)
func (repository *Repository) Checkout(stage *Stage) *Repository {
	if _, ok := stage.Repositorys[repository]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRepository(repository)
		}
	}
	return repository
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
	if _, ok := stage.Scenarios[scenario]; !ok {
		stage.Scenarios[scenario] = struct{}{}
		stage.Scenario_stagedOrder[scenario] = stage.ScenarioOrder
		stage.Scenario_orderStaged[stage.ScenarioOrder] = scenario
		stage.ScenarioOrder++
	}
	stage.Scenarios_mapString[scenario.Name] = scenario

	return scenario
}

// StagePreserveOrder puts scenario to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ScenarioOrder
// - update stage.ScenarioOrder accordingly
func (scenario *Scenario) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Scenarios[scenario]; !ok {
		stage.Scenarios[scenario] = struct{}{}

		if order > stage.ScenarioOrder {
			stage.ScenarioOrder = order
		}
		stage.Scenario_stagedOrder[scenario] = order
		stage.Scenario_orderStaged[order] = scenario
		stage.ScenarioOrder++
	}
	stage.Scenarios_mapString[scenario.Name] = scenario
}

// Unstage removes scenario off the model stage
func (scenario *Scenario) Unstage(stage *Stage) *Scenario {
	delete(stage.Scenarios, scenario)
	// issue1150
	// delete(stage.Scenario_stagedOrder, scenario)
	delete(stage.Scenarios_mapString, scenario.Name)

	return scenario
}

// UnstageVoid removes scenario off the model stage
func (scenario *Scenario) UnstageVoid(stage *Stage) {
	delete(stage.Scenarios, scenario)
	// issue1150
	// delete(stage.Scenario_stagedOrder, scenario)
	delete(stage.Scenarios_mapString, scenario.Name)
}

// commit scenario to the back repo (if it is already staged)
func (scenario *Scenario) Commit(stage *Stage) *Scenario {
	if _, ok := stage.Scenarios[scenario]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitScenario(scenario)
		}
	}
	return scenario
}

func (scenario *Scenario) CommitVoid(stage *Stage) {
	scenario.Commit(stage)
}

func (scenario *Scenario) StageVoid(stage *Stage) {
	scenario.Stage(stage)
}

// Checkout scenario to the back repo (if it is already staged)
func (scenario *Scenario) Checkout(stage *Stage) *Scenario {
	if _, ok := stage.Scenarios[scenario]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutScenario(scenario)
		}
	}
	return scenario
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
	if _, ok := stage.Users[user]; !ok {
		stage.Users[user] = struct{}{}
		stage.User_stagedOrder[user] = stage.UserOrder
		stage.User_orderStaged[stage.UserOrder] = user
		stage.UserOrder++
	}
	stage.Users_mapString[user.Name] = user

	return user
}

// StagePreserveOrder puts user to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.UserOrder
// - update stage.UserOrder accordingly
func (user *User) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Users[user]; !ok {
		stage.Users[user] = struct{}{}

		if order > stage.UserOrder {
			stage.UserOrder = order
		}
		stage.User_stagedOrder[user] = order
		stage.User_orderStaged[order] = user
		stage.UserOrder++
	}
	stage.Users_mapString[user.Name] = user
}

// Unstage removes user off the model stage
func (user *User) Unstage(stage *Stage) *User {
	delete(stage.Users, user)
	// issue1150
	// delete(stage.User_stagedOrder, user)
	delete(stage.Users_mapString, user.Name)

	return user
}

// UnstageVoid removes user off the model stage
func (user *User) UnstageVoid(stage *Stage) {
	delete(stage.Users, user)
	// issue1150
	// delete(stage.User_stagedOrder, user)
	delete(stage.Users_mapString, user.Name)
}

// commit user to the back repo (if it is already staged)
func (user *User) Commit(stage *Stage) *User {
	if _, ok := stage.Users[user]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitUser(user)
		}
	}
	return user
}

func (user *User) CommitVoid(stage *Stage) {
	user.Commit(stage)
}

func (user *User) StageVoid(stage *Stage) {
	user.Stage(stage)
}

// Checkout user to the back repo (if it is already staged)
func (user *User) Checkout(stage *Stage) *User {
	if _, ok := stage.Users[user]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutUser(user)
		}
	}
	return user
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
	if _, ok := stage.UserUses[useruse]; !ok {
		stage.UserUses[useruse] = struct{}{}
		stage.UserUse_stagedOrder[useruse] = stage.UserUseOrder
		stage.UserUse_orderStaged[stage.UserUseOrder] = useruse
		stage.UserUseOrder++
	}
	stage.UserUses_mapString[useruse.Name] = useruse

	return useruse
}

// StagePreserveOrder puts useruse to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.UserUseOrder
// - update stage.UserUseOrder accordingly
func (useruse *UserUse) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.UserUses[useruse]; !ok {
		stage.UserUses[useruse] = struct{}{}

		if order > stage.UserUseOrder {
			stage.UserUseOrder = order
		}
		stage.UserUse_stagedOrder[useruse] = order
		stage.UserUse_orderStaged[order] = useruse
		stage.UserUseOrder++
	}
	stage.UserUses_mapString[useruse.Name] = useruse
}

// Unstage removes useruse off the model stage
func (useruse *UserUse) Unstage(stage *Stage) *UserUse {
	delete(stage.UserUses, useruse)
	// issue1150
	// delete(stage.UserUse_stagedOrder, useruse)
	delete(stage.UserUses_mapString, useruse.Name)

	return useruse
}

// UnstageVoid removes useruse off the model stage
func (useruse *UserUse) UnstageVoid(stage *Stage) {
	delete(stage.UserUses, useruse)
	// issue1150
	// delete(stage.UserUse_stagedOrder, useruse)
	delete(stage.UserUses_mapString, useruse.Name)
}

// commit useruse to the back repo (if it is already staged)
func (useruse *UserUse) Commit(stage *Stage) *UserUse {
	if _, ok := stage.UserUses[useruse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitUserUse(useruse)
		}
	}
	return useruse
}

func (useruse *UserUse) CommitVoid(stage *Stage) {
	useruse.Commit(stage)
}

func (useruse *UserUse) StageVoid(stage *Stage) {
	useruse.Stage(stage)
}

// Checkout useruse to the back repo (if it is already staged)
func (useruse *UserUse) Checkout(stage *Stage) *UserUse {
	if _, ok := stage.UserUses[useruse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutUserUse(useruse)
		}
	}
	return useruse
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
	if _, ok := stage.Workspaces[workspace]; !ok {
		stage.Workspaces[workspace] = struct{}{}
		stage.Workspace_stagedOrder[workspace] = stage.WorkspaceOrder
		stage.Workspace_orderStaged[stage.WorkspaceOrder] = workspace
		stage.WorkspaceOrder++
	}
	stage.Workspaces_mapString[workspace.Name] = workspace

	return workspace
}

// StagePreserveOrder puts workspace to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.WorkspaceOrder
// - update stage.WorkspaceOrder accordingly
func (workspace *Workspace) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Workspaces[workspace]; !ok {
		stage.Workspaces[workspace] = struct{}{}

		if order > stage.WorkspaceOrder {
			stage.WorkspaceOrder = order
		}
		stage.Workspace_stagedOrder[workspace] = order
		stage.Workspace_orderStaged[order] = workspace
		stage.WorkspaceOrder++
	}
	stage.Workspaces_mapString[workspace.Name] = workspace
}

// Unstage removes workspace off the model stage
func (workspace *Workspace) Unstage(stage *Stage) *Workspace {
	delete(stage.Workspaces, workspace)
	// issue1150
	// delete(stage.Workspace_stagedOrder, workspace)
	delete(stage.Workspaces_mapString, workspace.Name)

	return workspace
}

// UnstageVoid removes workspace off the model stage
func (workspace *Workspace) UnstageVoid(stage *Stage) {
	delete(stage.Workspaces, workspace)
	// issue1150
	// delete(stage.Workspace_stagedOrder, workspace)
	delete(stage.Workspaces_mapString, workspace.Name)
}

// commit workspace to the back repo (if it is already staged)
func (workspace *Workspace) Commit(stage *Stage) *Workspace {
	if _, ok := stage.Workspaces[workspace]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitWorkspace(workspace)
		}
	}
	return workspace
}

func (workspace *Workspace) CommitVoid(stage *Stage) {
	workspace.Commit(stage)
}

func (workspace *Workspace) StageVoid(stage *Stage) {
	workspace.Stage(stage)
}

// Checkout workspace to the back repo (if it is already staged)
func (workspace *Workspace) Checkout(stage *Stage) *Workspace {
	if _, ok := stage.Workspaces[workspace]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutWorkspace(workspace)
		}
	}
	return workspace
}

// for satisfaction of GongStruct interface
func (workspace *Workspace) GetName() (res string) {
	return workspace.Name
}

// for satisfaction of GongStruct interface
func (workspace *Workspace) SetName(name string) {
	workspace.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMActorState(ActorState *ActorState)
	CreateORMActorStateShape(ActorStateShape *ActorStateShape)
	CreateORMActorStateTransition(ActorStateTransition *ActorStateTransition)
	CreateORMActorStateTransitionShape(ActorStateTransitionShape *ActorStateTransitionShape)
	CreateORMAnalysis(Analysis *Analysis)
	CreateORMControlPointShape(ControlPointShape *ControlPointShape)
	CreateORMDiagram(Diagram *Diagram)
	CreateORMDocument(Document *Document)
	CreateORMDocumentUse(DocumentUse *DocumentUse)
	CreateORMEvolutionDirection(EvolutionDirection *EvolutionDirection)
	CreateORMEvolutionDirectionShape(EvolutionDirectionShape *EvolutionDirectionShape)
	CreateORMFoo(Foo *Foo)
	CreateORMGeoObject(GeoObject *GeoObject)
	CreateORMGeoObjectUse(GeoObjectUse *GeoObjectUse)
	CreateORMGroup(Group *Group)
	CreateORMGroupUse(GroupUse *GroupUse)
	CreateORMLibrary(Library *Library)
	CreateORMMapObject(MapObject *MapObject)
	CreateORMMapObjectUse(MapObjectUse *MapObjectUse)
	CreateORMParameter(Parameter *Parameter)
	CreateORMParameterCategory(ParameterCategory *ParameterCategory)
	CreateORMParameterCategoryUse(ParameterCategoryUse *ParameterCategoryUse)
	CreateORMParameterShape(ParameterShape *ParameterShape)
	CreateORMParametersAggregate(ParametersAggregate *ParametersAggregate)
	CreateORMParametersAggregateShape(ParametersAggregateShape *ParametersAggregateShape)
	CreateORMPosition(Position *Position)
	CreateORMRepository(Repository *Repository)
	CreateORMScenario(Scenario *Scenario)
	CreateORMUser(User *User)
	CreateORMUserUse(UserUse *UserUse)
	CreateORMWorkspace(Workspace *Workspace)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMActorState(ActorState *ActorState)
	DeleteORMActorStateShape(ActorStateShape *ActorStateShape)
	DeleteORMActorStateTransition(ActorStateTransition *ActorStateTransition)
	DeleteORMActorStateTransitionShape(ActorStateTransitionShape *ActorStateTransitionShape)
	DeleteORMAnalysis(Analysis *Analysis)
	DeleteORMControlPointShape(ControlPointShape *ControlPointShape)
	DeleteORMDiagram(Diagram *Diagram)
	DeleteORMDocument(Document *Document)
	DeleteORMDocumentUse(DocumentUse *DocumentUse)
	DeleteORMEvolutionDirection(EvolutionDirection *EvolutionDirection)
	DeleteORMEvolutionDirectionShape(EvolutionDirectionShape *EvolutionDirectionShape)
	DeleteORMFoo(Foo *Foo)
	DeleteORMGeoObject(GeoObject *GeoObject)
	DeleteORMGeoObjectUse(GeoObjectUse *GeoObjectUse)
	DeleteORMGroup(Group *Group)
	DeleteORMGroupUse(GroupUse *GroupUse)
	DeleteORMLibrary(Library *Library)
	DeleteORMMapObject(MapObject *MapObject)
	DeleteORMMapObjectUse(MapObjectUse *MapObjectUse)
	DeleteORMParameter(Parameter *Parameter)
	DeleteORMParameterCategory(ParameterCategory *ParameterCategory)
	DeleteORMParameterCategoryUse(ParameterCategoryUse *ParameterCategoryUse)
	DeleteORMParameterShape(ParameterShape *ParameterShape)
	DeleteORMParametersAggregate(ParametersAggregate *ParametersAggregate)
	DeleteORMParametersAggregateShape(ParametersAggregateShape *ParametersAggregateShape)
	DeleteORMPosition(Position *Position)
	DeleteORMRepository(Repository *Repository)
	DeleteORMScenario(Scenario *Scenario)
	DeleteORMUser(User *User)
	DeleteORMUserUse(UserUse *UserUse)
	DeleteORMWorkspace(Workspace *Workspace)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.ActorStates = make(map[*ActorState]struct{})
	stage.ActorStates_mapString = make(map[string]*ActorState)
	stage.ActorState_stagedOrder = make(map[*ActorState]uint)
	stage.ActorStateOrder = 0

	stage.ActorStateShapes = make(map[*ActorStateShape]struct{})
	stage.ActorStateShapes_mapString = make(map[string]*ActorStateShape)
	stage.ActorStateShape_stagedOrder = make(map[*ActorStateShape]uint)
	stage.ActorStateShapeOrder = 0

	stage.ActorStateTransitions = make(map[*ActorStateTransition]struct{})
	stage.ActorStateTransitions_mapString = make(map[string]*ActorStateTransition)
	stage.ActorStateTransition_stagedOrder = make(map[*ActorStateTransition]uint)
	stage.ActorStateTransitionOrder = 0

	stage.ActorStateTransitionShapes = make(map[*ActorStateTransitionShape]struct{})
	stage.ActorStateTransitionShapes_mapString = make(map[string]*ActorStateTransitionShape)
	stage.ActorStateTransitionShape_stagedOrder = make(map[*ActorStateTransitionShape]uint)
	stage.ActorStateTransitionShapeOrder = 0

	stage.Analysiss = make(map[*Analysis]struct{})
	stage.Analysiss_mapString = make(map[string]*Analysis)
	stage.Analysis_stagedOrder = make(map[*Analysis]uint)
	stage.AnalysisOrder = 0

	stage.ControlPointShapes = make(map[*ControlPointShape]struct{})
	stage.ControlPointShapes_mapString = make(map[string]*ControlPointShape)
	stage.ControlPointShape_stagedOrder = make(map[*ControlPointShape]uint)
	stage.ControlPointShapeOrder = 0

	stage.Diagrams = make(map[*Diagram]struct{})
	stage.Diagrams_mapString = make(map[string]*Diagram)
	stage.Diagram_stagedOrder = make(map[*Diagram]uint)
	stage.DiagramOrder = 0

	stage.Documents = make(map[*Document]struct{})
	stage.Documents_mapString = make(map[string]*Document)
	stage.Document_stagedOrder = make(map[*Document]uint)
	stage.DocumentOrder = 0

	stage.DocumentUses = make(map[*DocumentUse]struct{})
	stage.DocumentUses_mapString = make(map[string]*DocumentUse)
	stage.DocumentUse_stagedOrder = make(map[*DocumentUse]uint)
	stage.DocumentUseOrder = 0

	stage.EvolutionDirections = make(map[*EvolutionDirection]struct{})
	stage.EvolutionDirections_mapString = make(map[string]*EvolutionDirection)
	stage.EvolutionDirection_stagedOrder = make(map[*EvolutionDirection]uint)
	stage.EvolutionDirectionOrder = 0

	stage.EvolutionDirectionShapes = make(map[*EvolutionDirectionShape]struct{})
	stage.EvolutionDirectionShapes_mapString = make(map[string]*EvolutionDirectionShape)
	stage.EvolutionDirectionShape_stagedOrder = make(map[*EvolutionDirectionShape]uint)
	stage.EvolutionDirectionShapeOrder = 0

	stage.Foos = make(map[*Foo]struct{})
	stage.Foos_mapString = make(map[string]*Foo)
	stage.Foo_stagedOrder = make(map[*Foo]uint)
	stage.FooOrder = 0

	stage.GeoObjects = make(map[*GeoObject]struct{})
	stage.GeoObjects_mapString = make(map[string]*GeoObject)
	stage.GeoObject_stagedOrder = make(map[*GeoObject]uint)
	stage.GeoObjectOrder = 0

	stage.GeoObjectUses = make(map[*GeoObjectUse]struct{})
	stage.GeoObjectUses_mapString = make(map[string]*GeoObjectUse)
	stage.GeoObjectUse_stagedOrder = make(map[*GeoObjectUse]uint)
	stage.GeoObjectUseOrder = 0

	stage.Groups = make(map[*Group]struct{})
	stage.Groups_mapString = make(map[string]*Group)
	stage.Group_stagedOrder = make(map[*Group]uint)
	stage.GroupOrder = 0

	stage.GroupUses = make(map[*GroupUse]struct{})
	stage.GroupUses_mapString = make(map[string]*GroupUse)
	stage.GroupUse_stagedOrder = make(map[*GroupUse]uint)
	stage.GroupUseOrder = 0

	stage.Librarys = make(map[*Library]struct{})
	stage.Librarys_mapString = make(map[string]*Library)
	stage.Library_stagedOrder = make(map[*Library]uint)
	stage.LibraryOrder = 0

	stage.MapObjects = make(map[*MapObject]struct{})
	stage.MapObjects_mapString = make(map[string]*MapObject)
	stage.MapObject_stagedOrder = make(map[*MapObject]uint)
	stage.MapObjectOrder = 0

	stage.MapObjectUses = make(map[*MapObjectUse]struct{})
	stage.MapObjectUses_mapString = make(map[string]*MapObjectUse)
	stage.MapObjectUse_stagedOrder = make(map[*MapObjectUse]uint)
	stage.MapObjectUseOrder = 0

	stage.Parameters = make(map[*Parameter]struct{})
	stage.Parameters_mapString = make(map[string]*Parameter)
	stage.Parameter_stagedOrder = make(map[*Parameter]uint)
	stage.ParameterOrder = 0

	stage.ParameterCategorys = make(map[*ParameterCategory]struct{})
	stage.ParameterCategorys_mapString = make(map[string]*ParameterCategory)
	stage.ParameterCategory_stagedOrder = make(map[*ParameterCategory]uint)
	stage.ParameterCategoryOrder = 0

	stage.ParameterCategoryUses = make(map[*ParameterCategoryUse]struct{})
	stage.ParameterCategoryUses_mapString = make(map[string]*ParameterCategoryUse)
	stage.ParameterCategoryUse_stagedOrder = make(map[*ParameterCategoryUse]uint)
	stage.ParameterCategoryUseOrder = 0

	stage.ParameterShapes = make(map[*ParameterShape]struct{})
	stage.ParameterShapes_mapString = make(map[string]*ParameterShape)
	stage.ParameterShape_stagedOrder = make(map[*ParameterShape]uint)
	stage.ParameterShapeOrder = 0

	stage.ParametersAggregates = make(map[*ParametersAggregate]struct{})
	stage.ParametersAggregates_mapString = make(map[string]*ParametersAggregate)
	stage.ParametersAggregate_stagedOrder = make(map[*ParametersAggregate]uint)
	stage.ParametersAggregateOrder = 0

	stage.ParametersAggregateShapes = make(map[*ParametersAggregateShape]struct{})
	stage.ParametersAggregateShapes_mapString = make(map[string]*ParametersAggregateShape)
	stage.ParametersAggregateShape_stagedOrder = make(map[*ParametersAggregateShape]uint)
	stage.ParametersAggregateShapeOrder = 0

	stage.Positions = make(map[*Position]struct{})
	stage.Positions_mapString = make(map[string]*Position)
	stage.Position_stagedOrder = make(map[*Position]uint)
	stage.PositionOrder = 0

	stage.Repositorys = make(map[*Repository]struct{})
	stage.Repositorys_mapString = make(map[string]*Repository)
	stage.Repository_stagedOrder = make(map[*Repository]uint)
	stage.RepositoryOrder = 0

	stage.Scenarios = make(map[*Scenario]struct{})
	stage.Scenarios_mapString = make(map[string]*Scenario)
	stage.Scenario_stagedOrder = make(map[*Scenario]uint)
	stage.ScenarioOrder = 0

	stage.Users = make(map[*User]struct{})
	stage.Users_mapString = make(map[string]*User)
	stage.User_stagedOrder = make(map[*User]uint)
	stage.UserOrder = 0

	stage.UserUses = make(map[*UserUse]struct{})
	stage.UserUses_mapString = make(map[string]*UserUse)
	stage.UserUse_stagedOrder = make(map[*UserUse]uint)
	stage.UserUseOrder = 0

	stage.Workspaces = make(map[*Workspace]struct{})
	stage.Workspaces_mapString = make(map[string]*Workspace)
	stage.Workspace_stagedOrder = make(map[*Workspace]uint)
	stage.WorkspaceOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.ActorStates = nil
	stage.ActorStates_mapString = nil

	stage.ActorStateShapes = nil
	stage.ActorStateShapes_mapString = nil

	stage.ActorStateTransitions = nil
	stage.ActorStateTransitions_mapString = nil

	stage.ActorStateTransitionShapes = nil
	stage.ActorStateTransitionShapes_mapString = nil

	stage.Analysiss = nil
	stage.Analysiss_mapString = nil

	stage.ControlPointShapes = nil
	stage.ControlPointShapes_mapString = nil

	stage.Diagrams = nil
	stage.Diagrams_mapString = nil

	stage.Documents = nil
	stage.Documents_mapString = nil

	stage.DocumentUses = nil
	stage.DocumentUses_mapString = nil

	stage.EvolutionDirections = nil
	stage.EvolutionDirections_mapString = nil

	stage.EvolutionDirectionShapes = nil
	stage.EvolutionDirectionShapes_mapString = nil

	stage.Foos = nil
	stage.Foos_mapString = nil

	stage.GeoObjects = nil
	stage.GeoObjects_mapString = nil

	stage.GeoObjectUses = nil
	stage.GeoObjectUses_mapString = nil

	stage.Groups = nil
	stage.Groups_mapString = nil

	stage.GroupUses = nil
	stage.GroupUses_mapString = nil

	stage.Librarys = nil
	stage.Librarys_mapString = nil

	stage.MapObjects = nil
	stage.MapObjects_mapString = nil

	stage.MapObjectUses = nil
	stage.MapObjectUses_mapString = nil

	stage.Parameters = nil
	stage.Parameters_mapString = nil

	stage.ParameterCategorys = nil
	stage.ParameterCategorys_mapString = nil

	stage.ParameterCategoryUses = nil
	stage.ParameterCategoryUses_mapString = nil

	stage.ParameterShapes = nil
	stage.ParameterShapes_mapString = nil

	stage.ParametersAggregates = nil
	stage.ParametersAggregates_mapString = nil

	stage.ParametersAggregateShapes = nil
	stage.ParametersAggregateShapes_mapString = nil

	stage.Positions = nil
	stage.Positions_mapString = nil

	stage.Repositorys = nil
	stage.Repositorys_mapString = nil

	stage.Scenarios = nil
	stage.Scenarios_mapString = nil

	stage.Users = nil
	stage.Users_mapString = nil

	stage.UserUses = nil
	stage.UserUses_mapString = nil

	stage.Workspaces = nil
	stage.Workspaces_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for actorstate := range stage.ActorStates {
		actorstate.Unstage(stage)
	}

	for actorstateshape := range stage.ActorStateShapes {
		actorstateshape.Unstage(stage)
	}

	for actorstatetransition := range stage.ActorStateTransitions {
		actorstatetransition.Unstage(stage)
	}

	for actorstatetransitionshape := range stage.ActorStateTransitionShapes {
		actorstatetransitionshape.Unstage(stage)
	}

	for analysis := range stage.Analysiss {
		analysis.Unstage(stage)
	}

	for controlpointshape := range stage.ControlPointShapes {
		controlpointshape.Unstage(stage)
	}

	for diagram := range stage.Diagrams {
		diagram.Unstage(stage)
	}

	for document := range stage.Documents {
		document.Unstage(stage)
	}

	for documentuse := range stage.DocumentUses {
		documentuse.Unstage(stage)
	}

	for evolutiondirection := range stage.EvolutionDirections {
		evolutiondirection.Unstage(stage)
	}

	for evolutiondirectionshape := range stage.EvolutionDirectionShapes {
		evolutiondirectionshape.Unstage(stage)
	}

	for foo := range stage.Foos {
		foo.Unstage(stage)
	}

	for geoobject := range stage.GeoObjects {
		geoobject.Unstage(stage)
	}

	for geoobjectuse := range stage.GeoObjectUses {
		geoobjectuse.Unstage(stage)
	}

	for group := range stage.Groups {
		group.Unstage(stage)
	}

	for groupuse := range stage.GroupUses {
		groupuse.Unstage(stage)
	}

	for library := range stage.Librarys {
		library.Unstage(stage)
	}

	for mapobject := range stage.MapObjects {
		mapobject.Unstage(stage)
	}

	for mapobjectuse := range stage.MapObjectUses {
		mapobjectuse.Unstage(stage)
	}

	for parameter := range stage.Parameters {
		parameter.Unstage(stage)
	}

	for parametercategory := range stage.ParameterCategorys {
		parametercategory.Unstage(stage)
	}

	for parametercategoryuse := range stage.ParameterCategoryUses {
		parametercategoryuse.Unstage(stage)
	}

	for parametershape := range stage.ParameterShapes {
		parametershape.Unstage(stage)
	}

	for parametersaggregate := range stage.ParametersAggregates {
		parametersaggregate.Unstage(stage)
	}

	for parametersaggregateshape := range stage.ParametersAggregateShapes {
		parametersaggregateshape.Unstage(stage)
	}

	for position := range stage.Positions {
		position.Unstage(stage)
	}

	for repository := range stage.Repositorys {
		repository.Unstage(stage)
	}

	for scenario := range stage.Scenarios {
		scenario.Unstage(stage)
	}

	for user := range stage.Users {
		user.Unstage(stage)
	}

	for useruse := range stage.UserUses {
		useruse.Unstage(stage)
	}

	for workspace := range stage.Workspaces {
		workspace.Unstage(stage)
	}

	// end of insertion point for array nil
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface{}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type GongstructIF interface {
	GetName() string
	SetName(string)
	CommitVoid(*Stage)
	StageVoid(*Stage)
	UnstageVoid(stage *Stage)
	GongGetFieldHeaders() []GongFieldHeader
	GongClean(stage *Stage) (modified bool)
	GongGetFieldValue(fieldName string, stage *Stage) GongFieldValue
	GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error
	GongGetGongstructName() string
	GongGetOrder(stage *Stage) uint
	GongGetReferenceIdentifier(stage *Stage) string
	GongGetIdentifier(stage *Stage) string
	GongCopy() GongstructIF
	GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) string
	GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) GongstructIF
	GongGetUUID(stage *Stage) string
}
type PointerToGongstruct interface {
	GongstructIF
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]struct{}) (sortedSlice []T) {
	for key := range set {
		sortedSlice = append(sortedSlice, key)
	}
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *Stage) (sortedSlice []T) {
	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any
}

type GongstructMapString interface {
	map[any]any
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *Stage) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*ActorState]any:
		return any(&stage.ActorStates).(*Type)
	case map[*ActorStateShape]any:
		return any(&stage.ActorStateShapes).(*Type)
	case map[*ActorStateTransition]any:
		return any(&stage.ActorStateTransitions).(*Type)
	case map[*ActorStateTransitionShape]any:
		return any(&stage.ActorStateTransitionShapes).(*Type)
	case map[*Analysis]any:
		return any(&stage.Analysiss).(*Type)
	case map[*ControlPointShape]any:
		return any(&stage.ControlPointShapes).(*Type)
	case map[*Diagram]any:
		return any(&stage.Diagrams).(*Type)
	case map[*Document]any:
		return any(&stage.Documents).(*Type)
	case map[*DocumentUse]any:
		return any(&stage.DocumentUses).(*Type)
	case map[*EvolutionDirection]any:
		return any(&stage.EvolutionDirections).(*Type)
	case map[*EvolutionDirectionShape]any:
		return any(&stage.EvolutionDirectionShapes).(*Type)
	case map[*Foo]any:
		return any(&stage.Foos).(*Type)
	case map[*GeoObject]any:
		return any(&stage.GeoObjects).(*Type)
	case map[*GeoObjectUse]any:
		return any(&stage.GeoObjectUses).(*Type)
	case map[*Group]any:
		return any(&stage.Groups).(*Type)
	case map[*GroupUse]any:
		return any(&stage.GroupUses).(*Type)
	case map[*Library]any:
		return any(&stage.Librarys).(*Type)
	case map[*MapObject]any:
		return any(&stage.MapObjects).(*Type)
	case map[*MapObjectUse]any:
		return any(&stage.MapObjectUses).(*Type)
	case map[*Parameter]any:
		return any(&stage.Parameters).(*Type)
	case map[*ParameterCategory]any:
		return any(&stage.ParameterCategorys).(*Type)
	case map[*ParameterCategoryUse]any:
		return any(&stage.ParameterCategoryUses).(*Type)
	case map[*ParameterShape]any:
		return any(&stage.ParameterShapes).(*Type)
	case map[*ParametersAggregate]any:
		return any(&stage.ParametersAggregates).(*Type)
	case map[*ParametersAggregateShape]any:
		return any(&stage.ParametersAggregateShapes).(*Type)
	case map[*Position]any:
		return any(&stage.Positions).(*Type)
	case map[*Repository]any:
		return any(&stage.Repositorys).(*Type)
	case map[*Scenario]any:
		return any(&stage.Scenarios).(*Type)
	case map[*User]any:
		return any(&stage.Users).(*Type)
	case map[*UserUse]any:
		return any(&stage.UserUses).(*Type)
	case map[*Workspace]any:
		return any(&stage.Workspaces).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged Gonstruct instance by their name
// Can be usefull if names are unique
func GongGetMap[Type GongstructIF](stage *Stage) map[string]Type {
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

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case ActorState:
		return any(&stage.ActorStates).(*map[*Type]struct{})
	case ActorStateShape:
		return any(&stage.ActorStateShapes).(*map[*Type]struct{})
	case ActorStateTransition:
		return any(&stage.ActorStateTransitions).(*map[*Type]struct{})
	case ActorStateTransitionShape:
		return any(&stage.ActorStateTransitionShapes).(*map[*Type]struct{})
	case Analysis:
		return any(&stage.Analysiss).(*map[*Type]struct{})
	case ControlPointShape:
		return any(&stage.ControlPointShapes).(*map[*Type]struct{})
	case Diagram:
		return any(&stage.Diagrams).(*map[*Type]struct{})
	case Document:
		return any(&stage.Documents).(*map[*Type]struct{})
	case DocumentUse:
		return any(&stage.DocumentUses).(*map[*Type]struct{})
	case EvolutionDirection:
		return any(&stage.EvolutionDirections).(*map[*Type]struct{})
	case EvolutionDirectionShape:
		return any(&stage.EvolutionDirectionShapes).(*map[*Type]struct{})
	case Foo:
		return any(&stage.Foos).(*map[*Type]struct{})
	case GeoObject:
		return any(&stage.GeoObjects).(*map[*Type]struct{})
	case GeoObjectUse:
		return any(&stage.GeoObjectUses).(*map[*Type]struct{})
	case Group:
		return any(&stage.Groups).(*map[*Type]struct{})
	case GroupUse:
		return any(&stage.GroupUses).(*map[*Type]struct{})
	case Library:
		return any(&stage.Librarys).(*map[*Type]struct{})
	case MapObject:
		return any(&stage.MapObjects).(*map[*Type]struct{})
	case MapObjectUse:
		return any(&stage.MapObjectUses).(*map[*Type]struct{})
	case Parameter:
		return any(&stage.Parameters).(*map[*Type]struct{})
	case ParameterCategory:
		return any(&stage.ParameterCategorys).(*map[*Type]struct{})
	case ParameterCategoryUse:
		return any(&stage.ParameterCategoryUses).(*map[*Type]struct{})
	case ParameterShape:
		return any(&stage.ParameterShapes).(*map[*Type]struct{})
	case ParametersAggregate:
		return any(&stage.ParametersAggregates).(*map[*Type]struct{})
	case ParametersAggregateShape:
		return any(&stage.ParametersAggregateShapes).(*map[*Type]struct{})
	case Position:
		return any(&stage.Positions).(*map[*Type]struct{})
	case Repository:
		return any(&stage.Repositorys).(*map[*Type]struct{})
	case Scenario:
		return any(&stage.Scenarios).(*map[*Type]struct{})
	case User:
		return any(&stage.Users).(*map[*Type]struct{})
	case UserUse:
		return any(&stage.UserUses).(*map[*Type]struct{})
	case Workspace:
		return any(&stage.Workspaces).(*map[*Type]struct{})
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *Stage) *map[Type]struct{} {
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

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *Stage) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case ActorState:
		return any(&stage.ActorStates_mapString).(*map[string]*Type)
	case ActorStateShape:
		return any(&stage.ActorStateShapes_mapString).(*map[string]*Type)
	case ActorStateTransition:
		return any(&stage.ActorStateTransitions_mapString).(*map[string]*Type)
	case ActorStateTransitionShape:
		return any(&stage.ActorStateTransitionShapes_mapString).(*map[string]*Type)
	case Analysis:
		return any(&stage.Analysiss_mapString).(*map[string]*Type)
	case ControlPointShape:
		return any(&stage.ControlPointShapes_mapString).(*map[string]*Type)
	case Diagram:
		return any(&stage.Diagrams_mapString).(*map[string]*Type)
	case Document:
		return any(&stage.Documents_mapString).(*map[string]*Type)
	case DocumentUse:
		return any(&stage.DocumentUses_mapString).(*map[string]*Type)
	case EvolutionDirection:
		return any(&stage.EvolutionDirections_mapString).(*map[string]*Type)
	case EvolutionDirectionShape:
		return any(&stage.EvolutionDirectionShapes_mapString).(*map[string]*Type)
	case Foo:
		return any(&stage.Foos_mapString).(*map[string]*Type)
	case GeoObject:
		return any(&stage.GeoObjects_mapString).(*map[string]*Type)
	case GeoObjectUse:
		return any(&stage.GeoObjectUses_mapString).(*map[string]*Type)
	case Group:
		return any(&stage.Groups_mapString).(*map[string]*Type)
	case GroupUse:
		return any(&stage.GroupUses_mapString).(*map[string]*Type)
	case Library:
		return any(&stage.Librarys_mapString).(*map[string]*Type)
	case MapObject:
		return any(&stage.MapObjects_mapString).(*map[string]*Type)
	case MapObjectUse:
		return any(&stage.MapObjectUses_mapString).(*map[string]*Type)
	case Parameter:
		return any(&stage.Parameters_mapString).(*map[string]*Type)
	case ParameterCategory:
		return any(&stage.ParameterCategorys_mapString).(*map[string]*Type)
	case ParameterCategoryUse:
		return any(&stage.ParameterCategoryUses_mapString).(*map[string]*Type)
	case ParameterShape:
		return any(&stage.ParameterShapes_mapString).(*map[string]*Type)
	case ParametersAggregate:
		return any(&stage.ParametersAggregates_mapString).(*map[string]*Type)
	case ParametersAggregateShape:
		return any(&stage.ParametersAggregateShapes_mapString).(*map[string]*Type)
	case Position:
		return any(&stage.Positions_mapString).(*map[string]*Type)
	case Repository:
		return any(&stage.Repositorys_mapString).(*map[string]*Type)
	case Scenario:
		return any(&stage.Scenarios_mapString).(*map[string]*Type)
	case User:
		return any(&stage.Users_mapString).(*map[string]*Type)
	case UserUse:
		return any(&stage.UserUses_mapString).(*map[string]*Type)
	case Workspace:
		return any(&stage.Workspaces_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case ActorState:
		return any(&ActorState{
			// Initialisation of associations
		}).(*Type)
	case ActorStateShape:
		return any(&ActorStateShape{
			// Initialisation of associations
			// field is initialized with an instance of ActorState with the name of the field
			ActorState: &ActorState{Name: "ActorState"},
		}).(*Type)
	case ActorStateTransition:
		return any(&ActorStateTransition{
			// Initialisation of associations
			// field is initialized with an instance of ActorState with the name of the field
			StartState: &ActorState{Name: "StartState"},
			// field is initialized with an instance of ActorState with the name of the field
			EndState: &ActorState{Name: "EndState"},
			// field is initialized with an instance of Parameter with the name of the field
			Justifications: []*Parameter{{Name: "Justifications"}},
		}).(*Type)
	case ActorStateTransitionShape:
		return any(&ActorStateTransitionShape{
			// Initialisation of associations
			// field is initialized with an instance of ActorStateTransition with the name of the field
			ActorStateTransition: &ActorStateTransition{Name: "ActorStateTransition"},
			// field is initialized with an instance of ActorStateShape with the name of the field
			Start: &ActorStateShape{Name: "Start"},
			// field is initialized with an instance of ActorStateShape with the name of the field
			End: &ActorStateShape{Name: "End"},
			// field is initialized with an instance of ControlPointShape with the name of the field
			ControlPointShapes: []*ControlPointShape{{Name: "ControlPointShapes"}},
		}).(*Type)
	case Analysis:
		return any(&Analysis{
			// Initialisation of associations
			// field is initialized with an instance of Scenario with the name of the field
			Scenarios: []*Scenario{{Name: "Scenarios"}},
			// field is initialized with an instance of GroupUse with the name of the field
			GroupUse: []*GroupUse{{Name: "GroupUse"}},
			// field is initialized with an instance of GeoObjectUse with the name of the field
			GeoObjectUse: []*GeoObjectUse{{Name: "GeoObjectUse"}},
			// field is initialized with an instance of MapObjectUse with the name of the field
			MapUse: []*MapObjectUse{{Name: "MapUse"}},
		}).(*Type)
	case ControlPointShape:
		return any(&ControlPointShape{
			// Initialisation of associations
		}).(*Type)
	case Diagram:
		return any(&Diagram{
			// Initialisation of associations
			// field is initialized with an instance of EvolutionDirectionShape with the name of the field
			EvolutionDirectionShapes: []*EvolutionDirectionShape{{Name: "EvolutionDirectionShapes"}},
			// field is initialized with an instance of EvolutionDirection with the name of the field
			EvolutionDirectionsWhoseNodeIsExpanded: []*EvolutionDirection{{Name: "EvolutionDirectionsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of ActorStateShape with the name of the field
			ActorStateShapes: []*ActorStateShape{{Name: "ActorStateShapes"}},
			// field is initialized with an instance of ActorState with the name of the field
			ActorStatesWhoseNodeIsExpanded: []*ActorState{{Name: "ActorStatesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of ParameterShape with the name of the field
			ParameterShapes: []*ParameterShape{{Name: "ParameterShapes"}},
			// field is initialized with an instance of Parameter with the name of the field
			ParametersWhoseNodeIsExpanded: []*Parameter{{Name: "ParametersWhoseNodeIsExpanded"}},
			// field is initialized with an instance of ParametersAggregateShape with the name of the field
			ScenarioParameterShapes: []*ParametersAggregateShape{{Name: "ScenarioParameterShapes"}},
			// field is initialized with an instance of ParametersAggregate with the name of the field
			ParametersAggregatesWhoseNodeIsExpanded: []*ParametersAggregate{{Name: "ParametersAggregatesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of ActorStateTransitionShape with the name of the field
			ActorStateTransitionShapes: []*ActorStateTransitionShape{{Name: "ActorStateTransitionShapes"}},
			// field is initialized with an instance of ActorStateTransition with the name of the field
			ActorStateTransitionsWhoseNodeIsExpanded: []*ActorStateTransition{{Name: "ActorStateTransitionsWhoseNodeIsExpanded"}},
		}).(*Type)
	case Document:
		return any(&Document{
			// Initialisation of associations
			// field is initialized with an instance of GeoObjectUse with the name of the field
			GeoObjectUse: []*GeoObjectUse{{Name: "GeoObjectUse"}},
		}).(*Type)
	case DocumentUse:
		return any(&DocumentUse{
			// Initialisation of associations
			// field is initialized with an instance of Document with the name of the field
			Document: &Document{Name: "Document"},
		}).(*Type)
	case EvolutionDirection:
		return any(&EvolutionDirection{
			// Initialisation of associations
		}).(*Type)
	case EvolutionDirectionShape:
		return any(&EvolutionDirectionShape{
			// Initialisation of associations
			// field is initialized with an instance of EvolutionDirection with the name of the field
			EvolutionDirection: &EvolutionDirection{Name: "EvolutionDirection"},
		}).(*Type)
	case Foo:
		return any(&Foo{
			// Initialisation of associations
		}).(*Type)
	case GeoObject:
		return any(&GeoObject{
			// Initialisation of associations
		}).(*Type)
	case GeoObjectUse:
		return any(&GeoObjectUse{
			// Initialisation of associations
			// field is initialized with an instance of GeoObject with the name of the field
			GeoObject: &GeoObject{Name: "GeoObject"},
		}).(*Type)
	case Group:
		return any(&Group{
			// Initialisation of associations
			// field is initialized with an instance of UserUse with the name of the field
			UserUse: []*UserUse{{Name: "UserUse"}},
		}).(*Type)
	case GroupUse:
		return any(&GroupUse{
			// Initialisation of associations
			// field is initialized with an instance of Group with the name of the field
			Group: &Group{Name: "Group"},
		}).(*Type)
	case Library:
		return any(&Library{
			// Initialisation of associations
			// field is initialized with an instance of Analysis with the name of the field
			Analyses: []*Analysis{{Name: "Analyses"}},
			// field is initialized with an instance of Library with the name of the field
			SubLibraries: []*Library{{Name: "SubLibraries"}},
			// field is initialized with an instance of Library with the name of the field
			SubLibrariesWhoseNodeIsExpanded: []*Library{{Name: "SubLibrariesWhoseNodeIsExpanded"}},
		}).(*Type)
	case MapObject:
		return any(&MapObject{
			// Initialisation of associations
		}).(*Type)
	case MapObjectUse:
		return any(&MapObjectUse{
			// Initialisation of associations
			// field is initialized with an instance of MapObject with the name of the field
			Map: &MapObject{Name: "Map"},
		}).(*Type)
	case Parameter:
		return any(&Parameter{
			// Initialisation of associations
			// field is initialized with an instance of GroupUse with the name of the field
			GroupUse: []*GroupUse{{Name: "GroupUse"}},
			// field is initialized with an instance of DocumentUse with the name of the field
			DocumentUse: []*DocumentUse{{Name: "DocumentUse"}},
			// field is initialized with an instance of GeoObjectUse with the name of the field
			GeoObjectUse: []*GeoObjectUse{{Name: "GeoObjectUse"}},
		}).(*Type)
	case ParameterCategory:
		return any(&ParameterCategory{
			// Initialisation of associations
			// field is initialized with an instance of ParameterShape with the name of the field
			ParameterUse: []*ParameterShape{{Name: "ParameterUse"}},
		}).(*Type)
	case ParameterCategoryUse:
		return any(&ParameterCategoryUse{
			// Initialisation of associations
			// field is initialized with an instance of ParameterCategory with the name of the field
			ParameterCategory: &ParameterCategory{Name: "ParameterCategory"},
		}).(*Type)
	case ParameterShape:
		return any(&ParameterShape{
			// Initialisation of associations
			// field is initialized with an instance of Parameter with the name of the field
			Parameter: &Parameter{Name: "Parameter"},
		}).(*Type)
	case ParametersAggregate:
		return any(&ParametersAggregate{
			// Initialisation of associations
			// field is initialized with an instance of Parameter with the name of the field
			Parameters: []*Parameter{{Name: "Parameters"}},
		}).(*Type)
	case ParametersAggregateShape:
		return any(&ParametersAggregateShape{
			// Initialisation of associations
			// field is initialized with an instance of ParametersAggregate with the name of the field
			ScenarioParameter: &ParametersAggregate{Name: "ScenarioParameter"},
		}).(*Type)
	case Position:
		return any(&Position{
			// Initialisation of associations
		}).(*Type)
	case Repository:
		return any(&Repository{
			// Initialisation of associations
			// field is initialized with an instance of ParameterShape with the name of the field
			ParameterUse: []*ParameterShape{{Name: "ParameterUse"}},
			// field is initialized with an instance of GroupUse with the name of the field
			GroupUse: []*GroupUse{{Name: "GroupUse"}},
		}).(*Type)
	case Scenario:
		return any(&Scenario{
			// Initialisation of associations
			// field is initialized with an instance of Diagram with the name of the field
			Diagrams: []*Diagram{{Name: "Diagrams"}},
			// field is initialized with an instance of ActorState with the name of the field
			ActorStates: []*ActorState{{Name: "ActorStates"}},
			// field is initialized with an instance of ActorStateTransition with the name of the field
			ActorStateTransitions: []*ActorStateTransition{{Name: "ActorStateTransitions"}},
			// field is initialized with an instance of EvolutionDirection with the name of the field
			EvolutionDirections: []*EvolutionDirection{{Name: "EvolutionDirections"}},
			// field is initialized with an instance of Parameter with the name of the field
			Parameters: []*Parameter{{Name: "Parameters"}},
			// field is initialized with an instance of ParametersAggregate with the name of the field
			ParametersAggretates: []*ParametersAggregate{{Name: "ParametersAggretates"}},
		}).(*Type)
	case User:
		return any(&User{
			// Initialisation of associations
		}).(*Type)
	case UserUse:
		return any(&UserUse{
			// Initialisation of associations
			// field is initialized with an instance of User with the name of the field
			User: &User{Name: "User"},
		}).(*Type)
	case Workspace:
		return any(&Workspace{
			// Initialisation of associations
			// field is initialized with an instance of Diagram with the name of the field
			SelectedDiagram: &Diagram{Name: "SelectedDiagram"},
			// field is initialized with an instance of EvolutionDirectionShape with the name of the field
			Default_EvolutionDirectionShape: &EvolutionDirectionShape{Name: "Default_EvolutionDirectionShape"},
			// field is initialized with an instance of ParameterShape with the name of the field
			Default_ParameterShape: &ParameterShape{Name: "Default_ParameterShape"},
			// field is initialized with an instance of ParametersAggregateShape with the name of the field
			Default_ScenarioParameterShape: &ParametersAggregateShape{Name: "Default_ScenarioParameterShape"},
			// field is initialized with an instance of ActorStateShape with the name of the field
			Default_ActorStateShape: &ActorStateShape{Name: "Default_ActorStateShape"},
			// field is initialized with an instance of ActorStateTransitionShape with the name of the field
			Default_ActorStateTransitionShape: &ActorStateTransitionShape{Name: "Default_ActorStateTransitionShape"},
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End][]*Start {
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

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End][]*Start {
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

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type GongstructIF]() (res string) {
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

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type GongstructIF]() (res []ReverseField) {
	res = make([]ReverseField, 0)

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
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []GongFieldHeader) {
	var ret Type
	return ret.GongGetFieldHeaders()
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
	case "LayoutDirection":
		enum := actorstate.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := actorstatetransition.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := analysis.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := diagram.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := document.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := evolutiondirection.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := geoobject.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := group.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := library.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := mapobject.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := parameter.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := parametercategory.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := parametersaggregate.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := position.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := repository.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := scenario.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := user.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := workspace.LayoutDirection
		res.valueString = enum.ToCodeString()
	}
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (actorstate *ActorState) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		actorstate.Name = value.GetValueString()
	case "Description":
		actorstate.Description = value.GetValueString()
	case "IsWithProbaility":
		actorstate.IsWithProbaility = value.GetValueBool()
	case "Probability":
		actorstate.Probability.FromCodeString(value.GetValueString())
	case "ComputedPrefix":
		actorstate.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		actorstate.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		actorstate.LayoutDirection.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (actorstateshape *ActorStateShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		actorstateshape.Name = value.GetValueString()
	case "ActorState":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			actorstateshape.ActorState = nil
			for __instance__ := range stage.ActorStates {
				if stage.ActorState_stagedOrder[__instance__] == uint(id) {
					actorstateshape.ActorState = __instance__
					break
				}
			}
		}
	case "X":
		actorstateshape.X = value.GetValueFloat()
	case "Y":
		actorstateshape.Y = value.GetValueFloat()
	case "Width":
		actorstateshape.Width = value.GetValueFloat()
	case "Height":
		actorstateshape.Height = value.GetValueFloat()
	case "IsHidden":
		actorstateshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (actorstatetransition *ActorStateTransition) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		actorstatetransition.Name = value.GetValueString()
	case "StartState":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			actorstatetransition.StartState = nil
			for __instance__ := range stage.ActorStates {
				if stage.ActorState_stagedOrder[__instance__] == uint(id) {
					actorstatetransition.StartState = __instance__
					break
				}
			}
		}
	case "EndState":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			actorstatetransition.EndState = nil
			for __instance__ := range stage.ActorStates {
				if stage.ActorState_stagedOrder[__instance__] == uint(id) {
					actorstatetransition.EndState = __instance__
					break
				}
			}
		}
	case "Justifications":
		actorstatetransition.Justifications = make([]*Parameter, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Parameters {
					if stage.Parameter_stagedOrder[__instance__] == uint(id) {
						actorstatetransition.Justifications = append(actorstatetransition.Justifications, __instance__)
						break
					}
				}
			}
		}
	case "ComputedPrefix":
		actorstatetransition.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		actorstatetransition.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		actorstatetransition.LayoutDirection.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		actorstatetransitionshape.Name = value.GetValueString()
	case "ActorStateTransition":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			actorstatetransitionshape.ActorStateTransition = nil
			for __instance__ := range stage.ActorStateTransitions {
				if stage.ActorStateTransition_stagedOrder[__instance__] == uint(id) {
					actorstatetransitionshape.ActorStateTransition = __instance__
					break
				}
			}
		}
	case "Start":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			actorstatetransitionshape.Start = nil
			for __instance__ := range stage.ActorStateShapes {
				if stage.ActorStateShape_stagedOrder[__instance__] == uint(id) {
					actorstatetransitionshape.Start = __instance__
					break
				}
			}
		}
	case "End":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			actorstatetransitionshape.End = nil
			for __instance__ := range stage.ActorStateShapes {
				if stage.ActorStateShape_stagedOrder[__instance__] == uint(id) {
					actorstatetransitionshape.End = __instance__
					break
				}
			}
		}
	case "X":
		actorstatetransitionshape.X = value.GetValueFloat()
	case "Y":
		actorstatetransitionshape.Y = value.GetValueFloat()
	case "Width":
		actorstatetransitionshape.Width = value.GetValueFloat()
	case "Height":
		actorstatetransitionshape.Height = value.GetValueFloat()
	case "IsHidden":
		actorstatetransitionshape.IsHidden = value.GetValueBool()
	case "ControlPointShapes":
		actorstatetransitionshape.ControlPointShapes = make([]*ControlPointShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ControlPointShapes {
					if stage.ControlPointShape_stagedOrder[__instance__] == uint(id) {
						actorstatetransitionshape.ControlPointShapes = append(actorstatetransitionshape.ControlPointShapes, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (analysis *Analysis) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		analysis.Name = value.GetValueString()
	case "Description":
		analysis.Description = value.GetValueString()
	case "Scenarios":
		analysis.Scenarios = make([]*Scenario, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Scenarios {
					if stage.Scenario_stagedOrder[__instance__] == uint(id) {
						analysis.Scenarios = append(analysis.Scenarios, __instance__)
						break
					}
				}
			}
		}
	case "IsScenariosNodeExpanded":
		analysis.IsScenariosNodeExpanded = value.GetValueBool()
	case "GroupUse":
		analysis.GroupUse = make([]*GroupUse, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.GroupUses {
					if stage.GroupUse_stagedOrder[__instance__] == uint(id) {
						analysis.GroupUse = append(analysis.GroupUse, __instance__)
						break
					}
				}
			}
		}
	case "IsGroupUseNodeExpanded":
		analysis.IsGroupUseNodeExpanded = value.GetValueBool()
	case "GeoObjectUse":
		analysis.GeoObjectUse = make([]*GeoObjectUse, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.GeoObjectUses {
					if stage.GeoObjectUse_stagedOrder[__instance__] == uint(id) {
						analysis.GeoObjectUse = append(analysis.GeoObjectUse, __instance__)
						break
					}
				}
			}
		}
	case "IsGeoObjectUseNodeExpanded":
		analysis.IsGeoObjectUseNodeExpanded = value.GetValueBool()
	case "MapUse":
		analysis.MapUse = make([]*MapObjectUse, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.MapObjectUses {
					if stage.MapObjectUse_stagedOrder[__instance__] == uint(id) {
						analysis.MapUse = append(analysis.MapUse, __instance__)
						break
					}
				}
			}
		}
	case "IsMapUseNodeExpanded":
		analysis.IsMapUseNodeExpanded = value.GetValueBool()
	case "ComputedPrefix":
		analysis.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		analysis.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		analysis.LayoutDirection.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (controlpointshape *ControlPointShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		controlpointshape.Name = value.GetValueString()
	case "X_Relative":
		controlpointshape.X_Relative = value.GetValueFloat()
	case "Y_Relative":
		controlpointshape.Y_Relative = value.GetValueFloat()
	case "IsStartShapeTheClosestShape":
		controlpointshape.IsStartShapeTheClosestShape = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (diagram *Diagram) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		diagram.Name = value.GetValueString()
	case "ComputedPrefix":
		diagram.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		diagram.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		diagram.LayoutDirection.FromCodeString(value.GetValueString())
	case "IsChecked":
		diagram.IsChecked = value.GetValueBool()
	case "IsShowPrefix":
		diagram.IsShowPrefix = value.GetValueBool()
	case "Description":
		diagram.Description = value.GetValueString()
	case "EvolutionDirectionShapes":
		diagram.EvolutionDirectionShapes = make([]*EvolutionDirectionShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.EvolutionDirectionShapes {
					if stage.EvolutionDirectionShape_stagedOrder[__instance__] == uint(id) {
						diagram.EvolutionDirectionShapes = append(diagram.EvolutionDirectionShapes, __instance__)
						break
					}
				}
			}
		}
	case "EvolutionDirectionsWhoseNodeIsExpanded":
		diagram.EvolutionDirectionsWhoseNodeIsExpanded = make([]*EvolutionDirection, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.EvolutionDirections {
					if stage.EvolutionDirection_stagedOrder[__instance__] == uint(id) {
						diagram.EvolutionDirectionsWhoseNodeIsExpanded = append(diagram.EvolutionDirectionsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsEvolutionDirectionsNodeExpanded":
		diagram.IsEvolutionDirectionsNodeExpanded = value.GetValueBool()
	case "ActorStateShapes":
		diagram.ActorStateShapes = make([]*ActorStateShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ActorStateShapes {
					if stage.ActorStateShape_stagedOrder[__instance__] == uint(id) {
						diagram.ActorStateShapes = append(diagram.ActorStateShapes, __instance__)
						break
					}
				}
			}
		}
	case "ActorStatesWhoseNodeIsExpanded":
		diagram.ActorStatesWhoseNodeIsExpanded = make([]*ActorState, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ActorStates {
					if stage.ActorState_stagedOrder[__instance__] == uint(id) {
						diagram.ActorStatesWhoseNodeIsExpanded = append(diagram.ActorStatesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsActorStatesNodeExpanded":
		diagram.IsActorStatesNodeExpanded = value.GetValueBool()
	case "ParameterShapes":
		diagram.ParameterShapes = make([]*ParameterShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ParameterShapes {
					if stage.ParameterShape_stagedOrder[__instance__] == uint(id) {
						diagram.ParameterShapes = append(diagram.ParameterShapes, __instance__)
						break
					}
				}
			}
		}
	case "ParametersWhoseNodeIsExpanded":
		diagram.ParametersWhoseNodeIsExpanded = make([]*Parameter, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Parameters {
					if stage.Parameter_stagedOrder[__instance__] == uint(id) {
						diagram.ParametersWhoseNodeIsExpanded = append(diagram.ParametersWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsParametersNodeExpanded":
		diagram.IsParametersNodeExpanded = value.GetValueBool()
	case "ScenarioParameterShapes":
		diagram.ScenarioParameterShapes = make([]*ParametersAggregateShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ParametersAggregateShapes {
					if stage.ParametersAggregateShape_stagedOrder[__instance__] == uint(id) {
						diagram.ScenarioParameterShapes = append(diagram.ScenarioParameterShapes, __instance__)
						break
					}
				}
			}
		}
	case "ParametersAggregatesWhoseNodeIsExpanded":
		diagram.ParametersAggregatesWhoseNodeIsExpanded = make([]*ParametersAggregate, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ParametersAggregates {
					if stage.ParametersAggregate_stagedOrder[__instance__] == uint(id) {
						diagram.ParametersAggregatesWhoseNodeIsExpanded = append(diagram.ParametersAggregatesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsParametersAggregatesNodeExpanded":
		diagram.IsParametersAggregatesNodeExpanded = value.GetValueBool()
	case "ActorStateTransitionShapes":
		diagram.ActorStateTransitionShapes = make([]*ActorStateTransitionShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ActorStateTransitionShapes {
					if stage.ActorStateTransitionShape_stagedOrder[__instance__] == uint(id) {
						diagram.ActorStateTransitionShapes = append(diagram.ActorStateTransitionShapes, __instance__)
						break
					}
				}
			}
		}
	case "ActorStateTransitionsWhoseNodeIsExpanded":
		diagram.ActorStateTransitionsWhoseNodeIsExpanded = make([]*ActorStateTransition, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ActorStateTransitions {
					if stage.ActorStateTransition_stagedOrder[__instance__] == uint(id) {
						diagram.ActorStateTransitionsWhoseNodeIsExpanded = append(diagram.ActorStateTransitionsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsActorStateTransitionsNodeExpanded":
		diagram.IsActorStateTransitionsNodeExpanded = value.GetValueBool()
	case "AxisOrign_X":
		diagram.AxisOrign_X = value.GetValueFloat()
	case "AxisOrign_Y":
		diagram.AxisOrign_Y = value.GetValueFloat()
	case "VerticalAxis_Top_Y":
		diagram.VerticalAxis_Top_Y = value.GetValueFloat()
	case "VerticalAxis_Bottom_Y":
		diagram.VerticalAxis_Bottom_Y = value.GetValueFloat()
	case "VerticalAxis_StrokeWidth":
		diagram.VerticalAxis_StrokeWidth = value.GetValueFloat()
	case "HorizontalAxis_Right_X":
		diagram.HorizontalAxis_Right_X = value.GetValueFloat()
	case "NumberOfYearsBetweenTicks":
		diagram.NumberOfYearsBetweenTicks = int(value.GetValueInt())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (document *Document) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		document.Name = value.GetValueString()
	case "GeoObjectUse":
		document.GeoObjectUse = make([]*GeoObjectUse, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.GeoObjectUses {
					if stage.GeoObjectUse_stagedOrder[__instance__] == uint(id) {
						document.GeoObjectUse = append(document.GeoObjectUse, __instance__)
						break
					}
				}
			}
		}
	case "ComputedPrefix":
		document.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		document.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		document.LayoutDirection.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (documentuse *DocumentUse) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		documentuse.Name = value.GetValueString()
	case "Document":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			documentuse.Document = nil
			for __instance__ := range stage.Documents {
				if stage.Document_stagedOrder[__instance__] == uint(id) {
					documentuse.Document = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (evolutiondirection *EvolutionDirection) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		evolutiondirection.Name = value.GetValueString()
	case "Description":
		evolutiondirection.Description = value.GetValueString()
	case "ComputedPrefix":
		evolutiondirection.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		evolutiondirection.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		evolutiondirection.LayoutDirection.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		evolutiondirectionshape.Name = value.GetValueString()
	case "EvolutionDirection":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			evolutiondirectionshape.EvolutionDirection = nil
			for __instance__ := range stage.EvolutionDirections {
				if stage.EvolutionDirection_stagedOrder[__instance__] == uint(id) {
					evolutiondirectionshape.EvolutionDirection = __instance__
					break
				}
			}
		}
	case "X":
		evolutiondirectionshape.X = value.GetValueFloat()
	case "Y":
		evolutiondirectionshape.Y = value.GetValueFloat()
	case "Width":
		evolutiondirectionshape.Width = value.GetValueFloat()
	case "Height":
		evolutiondirectionshape.Height = value.GetValueFloat()
	case "IsHidden":
		evolutiondirectionshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (foo *Foo) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		foo.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (geoobject *GeoObject) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		geoobject.Name = value.GetValueString()
	case "ComputedPrefix":
		geoobject.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		geoobject.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		geoobject.LayoutDirection.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (geoobjectuse *GeoObjectUse) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		geoobjectuse.Name = value.GetValueString()
	case "GeoObject":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			geoobjectuse.GeoObject = nil
			for __instance__ := range stage.GeoObjects {
				if stage.GeoObject_stagedOrder[__instance__] == uint(id) {
					geoobjectuse.GeoObject = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (group *Group) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		group.Name = value.GetValueString()
	case "UserUse":
		group.UserUse = make([]*UserUse, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.UserUses {
					if stage.UserUse_stagedOrder[__instance__] == uint(id) {
						group.UserUse = append(group.UserUse, __instance__)
						break
					}
				}
			}
		}
	case "ComputedPrefix":
		group.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		group.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		group.LayoutDirection.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (groupuse *GroupUse) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		groupuse.Name = value.GetValueString()
	case "Group":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			groupuse.Group = nil
			for __instance__ := range stage.Groups {
				if stage.Group_stagedOrder[__instance__] == uint(id) {
					groupuse.Group = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (library *Library) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		library.Name = value.GetValueString()
	case "Description":
		library.Description = value.GetValueString()
	case "ComputedPrefix":
		library.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		library.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		library.LayoutDirection.FromCodeString(value.GetValueString())
	case "IsRootLibrary":
		library.IsRootLibrary = value.GetValueBool()
	case "Analyses":
		library.Analyses = make([]*Analysis, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Analysiss {
					if stage.Analysis_stagedOrder[__instance__] == uint(id) {
						library.Analyses = append(library.Analyses, __instance__)
						break
					}
				}
			}
		}
	case "IsAnalysesNodeExpanded":
		library.IsAnalysesNodeExpanded = value.GetValueBool()
	case "SubLibraries":
		library.SubLibraries = make([]*Library, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Librarys {
					if stage.Library_stagedOrder[__instance__] == uint(id) {
						library.SubLibraries = append(library.SubLibraries, __instance__)
						break
					}
				}
			}
		}
	case "IsSubLibrariesNodeExpanded":
		library.IsSubLibrariesNodeExpanded = value.GetValueBool()
	case "SubLibrariesWhoseNodeIsExpanded":
		library.SubLibrariesWhoseNodeIsExpanded = make([]*Library, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Librarys {
					if stage.Library_stagedOrder[__instance__] == uint(id) {
						library.SubLibrariesWhoseNodeIsExpanded = append(library.SubLibrariesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "NbPixPerCharacter":
		library.NbPixPerCharacter = value.GetValueFloat()
	case "LogoSVGFile":
		library.LogoSVGFile = value.GetValueString()
	case "IsExpandedTmp":
		library.IsExpandedTmp = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (mapobject *MapObject) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		mapobject.Name = value.GetValueString()
	case "ComputedPrefix":
		mapobject.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		mapobject.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		mapobject.LayoutDirection.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (mapobjectuse *MapObjectUse) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		mapobjectuse.Name = value.GetValueString()
	case "Map":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			mapobjectuse.Map = nil
			for __instance__ := range stage.MapObjects {
				if stage.MapObject_stagedOrder[__instance__] == uint(id) {
					mapobjectuse.Map = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (parameter *Parameter) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		parameter.Name = value.GetValueString()
	case "Description":
		parameter.Description = value.GetValueString()
	case "IsResponse":
		parameter.IsResponse = value.GetValueBool()
	case "Force":
		parameter.Force = value.GetValueFloat()
	case "GroupUse":
		parameter.GroupUse = make([]*GroupUse, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.GroupUses {
					if stage.GroupUse_stagedOrder[__instance__] == uint(id) {
						parameter.GroupUse = append(parameter.GroupUse, __instance__)
						break
					}
				}
			}
		}
	case "DocumentUse":
		parameter.DocumentUse = make([]*DocumentUse, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DocumentUses {
					if stage.DocumentUse_stagedOrder[__instance__] == uint(id) {
						parameter.DocumentUse = append(parameter.DocumentUse, __instance__)
						break
					}
				}
			}
		}
	case "GeoObjectUse":
		parameter.GeoObjectUse = make([]*GeoObjectUse, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.GeoObjectUses {
					if stage.GeoObjectUse_stagedOrder[__instance__] == uint(id) {
						parameter.GeoObjectUse = append(parameter.GeoObjectUse, __instance__)
						break
					}
				}
			}
		}
	case "Tag":
		parameter.Tag = value.GetValueString()
	case "ComputedPrefix":
		parameter.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		parameter.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		parameter.LayoutDirection.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (parametercategory *ParameterCategory) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		parametercategory.Name = value.GetValueString()
	case "ParameterUse":
		parametercategory.ParameterUse = make([]*ParameterShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ParameterShapes {
					if stage.ParameterShape_stagedOrder[__instance__] == uint(id) {
						parametercategory.ParameterUse = append(parametercategory.ParameterUse, __instance__)
						break
					}
				}
			}
		}
	case "ComputedPrefix":
		parametercategory.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		parametercategory.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		parametercategory.LayoutDirection.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (parametercategoryuse *ParameterCategoryUse) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		parametercategoryuse.Name = value.GetValueString()
	case "ParameterCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parametercategoryuse.ParameterCategory = nil
			for __instance__ := range stage.ParameterCategorys {
				if stage.ParameterCategory_stagedOrder[__instance__] == uint(id) {
					parametercategoryuse.ParameterCategory = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (parametershape *ParameterShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		parametershape.Name = value.GetValueString()
	case "Parameter":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parametershape.Parameter = nil
			for __instance__ := range stage.Parameters {
				if stage.Parameter_stagedOrder[__instance__] == uint(id) {
					parametershape.Parameter = __instance__
					break
				}
			}
		}
	case "Direction":
		parametershape.Direction.FromCodeString(value.GetValueString())
	case "ShapeIsComputedFromModel":
		parametershape.ShapeIsComputedFromModel = value.GetValueBool()
	case "X":
		parametershape.X = value.GetValueFloat()
	case "Y":
		parametershape.Y = value.GetValueFloat()
	case "Width":
		parametershape.Width = value.GetValueFloat()
	case "Height":
		parametershape.Height = value.GetValueFloat()
	case "IsHidden":
		parametershape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (parametersaggregate *ParametersAggregate) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		parametersaggregate.Name = value.GetValueString()
	case "Tag":
		parametersaggregate.Tag = value.GetValueString()
	case "Description":
		parametersaggregate.Description = value.GetValueString()
	case "Parameters":
		parametersaggregate.Parameters = make([]*Parameter, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Parameters {
					if stage.Parameter_stagedOrder[__instance__] == uint(id) {
						parametersaggregate.Parameters = append(parametersaggregate.Parameters, __instance__)
						break
					}
				}
			}
		}
	case "ComputedPrefix":
		parametersaggregate.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		parametersaggregate.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		parametersaggregate.LayoutDirection.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (parametersaggregateshape *ParametersAggregateShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		parametersaggregateshape.Name = value.GetValueString()
	case "ScenarioParameter":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parametersaggregateshape.ScenarioParameter = nil
			for __instance__ := range stage.ParametersAggregates {
				if stage.ParametersAggregate_stagedOrder[__instance__] == uint(id) {
					parametersaggregateshape.ScenarioParameter = __instance__
					break
				}
			}
		}
	case "Direction":
		parametersaggregateshape.Direction.FromCodeString(value.GetValueString())
	case "X":
		parametersaggregateshape.X = value.GetValueFloat()
	case "Y":
		parametersaggregateshape.Y = value.GetValueFloat()
	case "Width":
		parametersaggregateshape.Width = value.GetValueFloat()
	case "Height":
		parametersaggregateshape.Height = value.GetValueFloat()
	case "IsHidden":
		parametersaggregateshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (position *Position) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		position.Name = value.GetValueString()
	case "Ordinate":
		position.Ordinate = value.GetValueFloat()
	case "ComputedPrefix":
		position.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		position.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		position.LayoutDirection.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (repository *Repository) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		repository.Name = value.GetValueString()
	case "ParameterUse":
		repository.ParameterUse = make([]*ParameterShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ParameterShapes {
					if stage.ParameterShape_stagedOrder[__instance__] == uint(id) {
						repository.ParameterUse = append(repository.ParameterUse, __instance__)
						break
					}
				}
			}
		}
	case "GroupUse":
		repository.GroupUse = make([]*GroupUse, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.GroupUses {
					if stage.GroupUse_stagedOrder[__instance__] == uint(id) {
						repository.GroupUse = append(repository.GroupUse, __instance__)
						break
					}
				}
			}
		}
	case "ComputedPrefix":
		repository.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		repository.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		repository.LayoutDirection.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (scenario *Scenario) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		scenario.Name = value.GetValueString()
	case "Description":
		scenario.Description = value.GetValueString()
	case "Diagrams":
		scenario.Diagrams = make([]*Diagram, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Diagrams {
					if stage.Diagram_stagedOrder[__instance__] == uint(id) {
						scenario.Diagrams = append(scenario.Diagrams, __instance__)
						break
					}
				}
			}
		}
	case "IsDiagramsNodeExpanded":
		scenario.IsDiagramsNodeExpanded = value.GetValueBool()
	case "ActorStates":
		scenario.ActorStates = make([]*ActorState, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ActorStates {
					if stage.ActorState_stagedOrder[__instance__] == uint(id) {
						scenario.ActorStates = append(scenario.ActorStates, __instance__)
						break
					}
				}
			}
		}
	case "IsActorStatesNodeExpanded":
		scenario.IsActorStatesNodeExpanded = value.GetValueBool()
	case "ActorStateTransitions":
		scenario.ActorStateTransitions = make([]*ActorStateTransition, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ActorStateTransitions {
					if stage.ActorStateTransition_stagedOrder[__instance__] == uint(id) {
						scenario.ActorStateTransitions = append(scenario.ActorStateTransitions, __instance__)
						break
					}
				}
			}
		}
	case "IsActorStateTransitionsNodeExpanded":
		scenario.IsActorStateTransitionsNodeExpanded = value.GetValueBool()
	case "EvolutionDirections":
		scenario.EvolutionDirections = make([]*EvolutionDirection, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.EvolutionDirections {
					if stage.EvolutionDirection_stagedOrder[__instance__] == uint(id) {
						scenario.EvolutionDirections = append(scenario.EvolutionDirections, __instance__)
						break
					}
				}
			}
		}
	case "IsEvolutionDirectionsNodeExpanded":
		scenario.IsEvolutionDirectionsNodeExpanded = value.GetValueBool()
	case "Parameters":
		scenario.Parameters = make([]*Parameter, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Parameters {
					if stage.Parameter_stagedOrder[__instance__] == uint(id) {
						scenario.Parameters = append(scenario.Parameters, __instance__)
						break
					}
				}
			}
		}
	case "IsParametersNodeExpanded":
		scenario.IsParametersNodeExpanded = value.GetValueBool()
	case "ParametersAggretates":
		scenario.ParametersAggretates = make([]*ParametersAggregate, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ParametersAggregates {
					if stage.ParametersAggregate_stagedOrder[__instance__] == uint(id) {
						scenario.ParametersAggretates = append(scenario.ParametersAggretates, __instance__)
						break
					}
				}
			}
		}
	case "IsParametersAggretatesNodeExpanded":
		scenario.IsParametersAggretatesNodeExpanded = value.GetValueBool()
	case "ComputedPrefix":
		scenario.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		scenario.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		scenario.LayoutDirection.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (user *User) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		user.Name = value.GetValueString()
	case "ComputedPrefix":
		user.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		user.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		user.LayoutDirection.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (useruse *UserUse) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		useruse.Name = value.GetValueString()
	case "User":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			useruse.User = nil
			for __instance__ := range stage.Users {
				if stage.User_stagedOrder[__instance__] == uint(id) {
					useruse.User = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (workspace *Workspace) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		workspace.Name = value.GetValueString()
	case "SelectedDiagram":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			workspace.SelectedDiagram = nil
			for __instance__ := range stage.Diagrams {
				if stage.Diagram_stagedOrder[__instance__] == uint(id) {
					workspace.SelectedDiagram = __instance__
					break
				}
			}
		}
	case "Default_EvolutionDirectionShape":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			workspace.Default_EvolutionDirectionShape = nil
			for __instance__ := range stage.EvolutionDirectionShapes {
				if stage.EvolutionDirectionShape_stagedOrder[__instance__] == uint(id) {
					workspace.Default_EvolutionDirectionShape = __instance__
					break
				}
			}
		}
	case "Default_ParameterShape":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			workspace.Default_ParameterShape = nil
			for __instance__ := range stage.ParameterShapes {
				if stage.ParameterShape_stagedOrder[__instance__] == uint(id) {
					workspace.Default_ParameterShape = __instance__
					break
				}
			}
		}
	case "Default_ScenarioParameterShape":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			workspace.Default_ScenarioParameterShape = nil
			for __instance__ := range stage.ParametersAggregateShapes {
				if stage.ParametersAggregateShape_stagedOrder[__instance__] == uint(id) {
					workspace.Default_ScenarioParameterShape = __instance__
					break
				}
			}
		}
	case "Default_ActorStateShape":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			workspace.Default_ActorStateShape = nil
			for __instance__ := range stage.ActorStateShapes {
				if stage.ActorStateShape_stagedOrder[__instance__] == uint(id) {
					workspace.Default_ActorStateShape = __instance__
					break
				}
			}
		}
	case "Default_ActorStateTransitionShape":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			workspace.Default_ActorStateTransitionShape = nil
			for __instance__ := range stage.ActorStateTransitionShapes {
				if stage.ActorStateTransitionShape_stagedOrder[__instance__] == uint(id) {
					workspace.Default_ActorStateTransitionShape = __instance__
					break
				}
			}
		}
	case "ComputedPrefix":
		workspace.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		workspace.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		workspace.LayoutDirection.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
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

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	stage.ActorStates_mapString = make(map[string]*ActorState)
	for actorstate := range stage.ActorStates {
		stage.ActorStates_mapString[actorstate.Name] = actorstate
	}

	stage.ActorStateShapes_mapString = make(map[string]*ActorStateShape)
	for actorstateshape := range stage.ActorStateShapes {
		stage.ActorStateShapes_mapString[actorstateshape.Name] = actorstateshape
	}

	stage.ActorStateTransitions_mapString = make(map[string]*ActorStateTransition)
	for actorstatetransition := range stage.ActorStateTransitions {
		stage.ActorStateTransitions_mapString[actorstatetransition.Name] = actorstatetransition
	}

	stage.ActorStateTransitionShapes_mapString = make(map[string]*ActorStateTransitionShape)
	for actorstatetransitionshape := range stage.ActorStateTransitionShapes {
		stage.ActorStateTransitionShapes_mapString[actorstatetransitionshape.Name] = actorstatetransitionshape
	}

	stage.Analysiss_mapString = make(map[string]*Analysis)
	for analysis := range stage.Analysiss {
		stage.Analysiss_mapString[analysis.Name] = analysis
	}

	stage.ControlPointShapes_mapString = make(map[string]*ControlPointShape)
	for controlpointshape := range stage.ControlPointShapes {
		stage.ControlPointShapes_mapString[controlpointshape.Name] = controlpointshape
	}

	stage.Diagrams_mapString = make(map[string]*Diagram)
	for diagram := range stage.Diagrams {
		stage.Diagrams_mapString[diagram.Name] = diagram
	}

	stage.Documents_mapString = make(map[string]*Document)
	for document := range stage.Documents {
		stage.Documents_mapString[document.Name] = document
	}

	stage.DocumentUses_mapString = make(map[string]*DocumentUse)
	for documentuse := range stage.DocumentUses {
		stage.DocumentUses_mapString[documentuse.Name] = documentuse
	}

	stage.EvolutionDirections_mapString = make(map[string]*EvolutionDirection)
	for evolutiondirection := range stage.EvolutionDirections {
		stage.EvolutionDirections_mapString[evolutiondirection.Name] = evolutiondirection
	}

	stage.EvolutionDirectionShapes_mapString = make(map[string]*EvolutionDirectionShape)
	for evolutiondirectionshape := range stage.EvolutionDirectionShapes {
		stage.EvolutionDirectionShapes_mapString[evolutiondirectionshape.Name] = evolutiondirectionshape
	}

	stage.Foos_mapString = make(map[string]*Foo)
	for foo := range stage.Foos {
		stage.Foos_mapString[foo.Name] = foo
	}

	stage.GeoObjects_mapString = make(map[string]*GeoObject)
	for geoobject := range stage.GeoObjects {
		stage.GeoObjects_mapString[geoobject.Name] = geoobject
	}

	stage.GeoObjectUses_mapString = make(map[string]*GeoObjectUse)
	for geoobjectuse := range stage.GeoObjectUses {
		stage.GeoObjectUses_mapString[geoobjectuse.Name] = geoobjectuse
	}

	stage.Groups_mapString = make(map[string]*Group)
	for group := range stage.Groups {
		stage.Groups_mapString[group.Name] = group
	}

	stage.GroupUses_mapString = make(map[string]*GroupUse)
	for groupuse := range stage.GroupUses {
		stage.GroupUses_mapString[groupuse.Name] = groupuse
	}

	stage.Librarys_mapString = make(map[string]*Library)
	for library := range stage.Librarys {
		stage.Librarys_mapString[library.Name] = library
	}

	stage.MapObjects_mapString = make(map[string]*MapObject)
	for mapobject := range stage.MapObjects {
		stage.MapObjects_mapString[mapobject.Name] = mapobject
	}

	stage.MapObjectUses_mapString = make(map[string]*MapObjectUse)
	for mapobjectuse := range stage.MapObjectUses {
		stage.MapObjectUses_mapString[mapobjectuse.Name] = mapobjectuse
	}

	stage.Parameters_mapString = make(map[string]*Parameter)
	for parameter := range stage.Parameters {
		stage.Parameters_mapString[parameter.Name] = parameter
	}

	stage.ParameterCategorys_mapString = make(map[string]*ParameterCategory)
	for parametercategory := range stage.ParameterCategorys {
		stage.ParameterCategorys_mapString[parametercategory.Name] = parametercategory
	}

	stage.ParameterCategoryUses_mapString = make(map[string]*ParameterCategoryUse)
	for parametercategoryuse := range stage.ParameterCategoryUses {
		stage.ParameterCategoryUses_mapString[parametercategoryuse.Name] = parametercategoryuse
	}

	stage.ParameterShapes_mapString = make(map[string]*ParameterShape)
	for parametershape := range stage.ParameterShapes {
		stage.ParameterShapes_mapString[parametershape.Name] = parametershape
	}

	stage.ParametersAggregates_mapString = make(map[string]*ParametersAggregate)
	for parametersaggregate := range stage.ParametersAggregates {
		stage.ParametersAggregates_mapString[parametersaggregate.Name] = parametersaggregate
	}

	stage.ParametersAggregateShapes_mapString = make(map[string]*ParametersAggregateShape)
	for parametersaggregateshape := range stage.ParametersAggregateShapes {
		stage.ParametersAggregateShapes_mapString[parametersaggregateshape.Name] = parametersaggregateshape
	}

	stage.Positions_mapString = make(map[string]*Position)
	for position := range stage.Positions {
		stage.Positions_mapString[position.Name] = position
	}

	stage.Repositorys_mapString = make(map[string]*Repository)
	for repository := range stage.Repositorys {
		stage.Repositorys_mapString[repository.Name] = repository
	}

	stage.Scenarios_mapString = make(map[string]*Scenario)
	for scenario := range stage.Scenarios {
		stage.Scenarios_mapString[scenario.Name] = scenario
	}

	stage.Users_mapString = make(map[string]*User)
	for user := range stage.Users {
		stage.Users_mapString[user.Name] = user
	}

	stage.UserUses_mapString = make(map[string]*UserUse)
	for useruse := range stage.UserUses {
		stage.UserUses_mapString[useruse.Name] = useruse
	}

	stage.Workspaces_mapString = make(map[string]*Workspace)
	for workspace := range stage.Workspaces {
		stage.Workspaces_mapString[workspace.Name] = workspace
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
