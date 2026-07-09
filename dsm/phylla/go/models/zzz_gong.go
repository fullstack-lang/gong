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

	phylla_go "github.com/fullstack-lang/gong/dsm/phylla/go"
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
	ProbeLoadSuffix                  = ":load of the probe"
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

func (stage *Stage) GetProbeLoadStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeLoadSuffix
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
	AxesShapes                map[*AxesShape]struct{}
	AxesShapes_instance       map[*AxesShape]*AxesShape
	AxesShapes_mapString      map[string]*AxesShape
	AxesShapeOrder            uint
	AxesShape_stagedOrder     map[*AxesShape]uint
	AxesShape_orderStaged     map[uint]*AxesShape
	AxesShapes_reference      map[*AxesShape]*AxesShape
	AxesShapes_referenceOrder map[*AxesShape]uint

	// insertion point for slice of pointers maps
	OnAfterAxesShapeCreateCallback OnAfterCreateInterface[AxesShape]
	OnAfterAxesShapeUpdateCallback OnAfterUpdateInterface[AxesShape]
	OnAfterAxesShapeDeleteCallback OnAfterDeleteInterface[AxesShape]
	OnAfterAxesShapeReadCallback   OnAfterReadInterface[AxesShape]

	CircleGridShapes                map[*CircleGridShape]struct{}
	CircleGridShapes_instance       map[*CircleGridShape]*CircleGridShape
	CircleGridShapes_mapString      map[string]*CircleGridShape
	CircleGridShapeOrder            uint
	CircleGridShape_stagedOrder     map[*CircleGridShape]uint
	CircleGridShape_orderStaged     map[uint]*CircleGridShape
	CircleGridShapes_reference      map[*CircleGridShape]*CircleGridShape
	CircleGridShapes_referenceOrder map[*CircleGridShape]uint

	// insertion point for slice of pointers maps
	OnAfterCircleGridShapeCreateCallback OnAfterCreateInterface[CircleGridShape]
	OnAfterCircleGridShapeUpdateCallback OnAfterUpdateInterface[CircleGridShape]
	OnAfterCircleGridShapeDeleteCallback OnAfterDeleteInterface[CircleGridShape]
	OnAfterCircleGridShapeReadCallback   OnAfterReadInterface[CircleGridShape]

	ExplanationTextShapes                map[*ExplanationTextShape]struct{}
	ExplanationTextShapes_instance       map[*ExplanationTextShape]*ExplanationTextShape
	ExplanationTextShapes_mapString      map[string]*ExplanationTextShape
	ExplanationTextShapeOrder            uint
	ExplanationTextShape_stagedOrder     map[*ExplanationTextShape]uint
	ExplanationTextShape_orderStaged     map[uint]*ExplanationTextShape
	ExplanationTextShapes_reference      map[*ExplanationTextShape]*ExplanationTextShape
	ExplanationTextShapes_referenceOrder map[*ExplanationTextShape]uint

	// insertion point for slice of pointers maps
	OnAfterExplanationTextShapeCreateCallback OnAfterCreateInterface[ExplanationTextShape]
	OnAfterExplanationTextShapeUpdateCallback OnAfterUpdateInterface[ExplanationTextShape]
	OnAfterExplanationTextShapeDeleteCallback OnAfterDeleteInterface[ExplanationTextShape]
	OnAfterExplanationTextShapeReadCallback   OnAfterReadInterface[ExplanationTextShape]

	GridPathShapes                map[*GridPathShape]struct{}
	GridPathShapes_instance       map[*GridPathShape]*GridPathShape
	GridPathShapes_mapString      map[string]*GridPathShape
	GridPathShapeOrder            uint
	GridPathShape_stagedOrder     map[*GridPathShape]uint
	GridPathShape_orderStaged     map[uint]*GridPathShape
	GridPathShapes_reference      map[*GridPathShape]*GridPathShape
	GridPathShapes_referenceOrder map[*GridPathShape]uint

	// insertion point for slice of pointers maps
	OnAfterGridPathShapeCreateCallback OnAfterCreateInterface[GridPathShape]
	OnAfterGridPathShapeUpdateCallback OnAfterUpdateInterface[GridPathShape]
	OnAfterGridPathShapeDeleteCallback OnAfterDeleteInterface[GridPathShape]
	OnAfterGridPathShapeReadCallback   OnAfterReadInterface[GridPathShape]

	GrowthCurveRhombusGridShapes                map[*GrowthCurveRhombusGridShape]struct{}
	GrowthCurveRhombusGridShapes_instance       map[*GrowthCurveRhombusGridShape]*GrowthCurveRhombusGridShape
	GrowthCurveRhombusGridShapes_mapString      map[string]*GrowthCurveRhombusGridShape
	GrowthCurveRhombusGridShapeOrder            uint
	GrowthCurveRhombusGridShape_stagedOrder     map[*GrowthCurveRhombusGridShape]uint
	GrowthCurveRhombusGridShape_orderStaged     map[uint]*GrowthCurveRhombusGridShape
	GrowthCurveRhombusGridShapes_reference      map[*GrowthCurveRhombusGridShape]*GrowthCurveRhombusGridShape
	GrowthCurveRhombusGridShapes_referenceOrder map[*GrowthCurveRhombusGridShape]uint

	// insertion point for slice of pointers maps
	GrowthCurveRhombusGridShape_GrowthCurveRhombusShapes_reverseMap map[*GrowthCurveRhombusShape]*GrowthCurveRhombusGridShape

	OnAfterGrowthCurveRhombusGridShapeCreateCallback OnAfterCreateInterface[GrowthCurveRhombusGridShape]
	OnAfterGrowthCurveRhombusGridShapeUpdateCallback OnAfterUpdateInterface[GrowthCurveRhombusGridShape]
	OnAfterGrowthCurveRhombusGridShapeDeleteCallback OnAfterDeleteInterface[GrowthCurveRhombusGridShape]
	OnAfterGrowthCurveRhombusGridShapeReadCallback   OnAfterReadInterface[GrowthCurveRhombusGridShape]

	GrowthCurveRhombusShapes                map[*GrowthCurveRhombusShape]struct{}
	GrowthCurveRhombusShapes_instance       map[*GrowthCurveRhombusShape]*GrowthCurveRhombusShape
	GrowthCurveRhombusShapes_mapString      map[string]*GrowthCurveRhombusShape
	GrowthCurveRhombusShapeOrder            uint
	GrowthCurveRhombusShape_stagedOrder     map[*GrowthCurveRhombusShape]uint
	GrowthCurveRhombusShape_orderStaged     map[uint]*GrowthCurveRhombusShape
	GrowthCurveRhombusShapes_reference      map[*GrowthCurveRhombusShape]*GrowthCurveRhombusShape
	GrowthCurveRhombusShapes_referenceOrder map[*GrowthCurveRhombusShape]uint

	// insertion point for slice of pointers maps
	OnAfterGrowthCurveRhombusShapeCreateCallback OnAfterCreateInterface[GrowthCurveRhombusShape]
	OnAfterGrowthCurveRhombusShapeUpdateCallback OnAfterUpdateInterface[GrowthCurveRhombusShape]
	OnAfterGrowthCurveRhombusShapeDeleteCallback OnAfterDeleteInterface[GrowthCurveRhombusShape]
	OnAfterGrowthCurveRhombusShapeReadCallback   OnAfterReadInterface[GrowthCurveRhombusShape]

	InitialRhombusGridShapes                map[*InitialRhombusGridShape]struct{}
	InitialRhombusGridShapes_instance       map[*InitialRhombusGridShape]*InitialRhombusGridShape
	InitialRhombusGridShapes_mapString      map[string]*InitialRhombusGridShape
	InitialRhombusGridShapeOrder            uint
	InitialRhombusGridShape_stagedOrder     map[*InitialRhombusGridShape]uint
	InitialRhombusGridShape_orderStaged     map[uint]*InitialRhombusGridShape
	InitialRhombusGridShapes_reference      map[*InitialRhombusGridShape]*InitialRhombusGridShape
	InitialRhombusGridShapes_referenceOrder map[*InitialRhombusGridShape]uint

	// insertion point for slice of pointers maps
	InitialRhombusGridShape_InitialRhombusShapes_reverseMap map[*InitialRhombusShape]*InitialRhombusGridShape

	OnAfterInitialRhombusGridShapeCreateCallback OnAfterCreateInterface[InitialRhombusGridShape]
	OnAfterInitialRhombusGridShapeUpdateCallback OnAfterUpdateInterface[InitialRhombusGridShape]
	OnAfterInitialRhombusGridShapeDeleteCallback OnAfterDeleteInterface[InitialRhombusGridShape]
	OnAfterInitialRhombusGridShapeReadCallback   OnAfterReadInterface[InitialRhombusGridShape]

	InitialRhombusShapes                map[*InitialRhombusShape]struct{}
	InitialRhombusShapes_instance       map[*InitialRhombusShape]*InitialRhombusShape
	InitialRhombusShapes_mapString      map[string]*InitialRhombusShape
	InitialRhombusShapeOrder            uint
	InitialRhombusShape_stagedOrder     map[*InitialRhombusShape]uint
	InitialRhombusShape_orderStaged     map[uint]*InitialRhombusShape
	InitialRhombusShapes_reference      map[*InitialRhombusShape]*InitialRhombusShape
	InitialRhombusShapes_referenceOrder map[*InitialRhombusShape]uint

	// insertion point for slice of pointers maps
	OnAfterInitialRhombusShapeCreateCallback OnAfterCreateInterface[InitialRhombusShape]
	OnAfterInitialRhombusShapeUpdateCallback OnAfterUpdateInterface[InitialRhombusShape]
	OnAfterInitialRhombusShapeDeleteCallback OnAfterDeleteInterface[InitialRhombusShape]
	OnAfterInitialRhombusShapeReadCallback   OnAfterReadInterface[InitialRhombusShape]

	Librarys                map[*Library]struct{}
	Librarys_instance       map[*Library]*Library
	Librarys_mapString      map[string]*Library
	LibraryOrder            uint
	Library_stagedOrder     map[*Library]uint
	Library_orderStaged     map[uint]*Library
	Librarys_reference      map[*Library]*Library
	Librarys_referenceOrder map[*Library]uint

	// insertion point for slice of pointers maps
	Library_SubLibraries_reverseMap map[*Library]*Library

	Library_Plants_reverseMap map[*Plant]*Library

	OnAfterLibraryCreateCallback OnAfterCreateInterface[Library]
	OnAfterLibraryUpdateCallback OnAfterUpdateInterface[Library]
	OnAfterLibraryDeleteCallback OnAfterDeleteInterface[Library]
	OnAfterLibraryReadCallback   OnAfterReadInterface[Library]

	NextCircleShapes                map[*NextCircleShape]struct{}
	NextCircleShapes_instance       map[*NextCircleShape]*NextCircleShape
	NextCircleShapes_mapString      map[string]*NextCircleShape
	NextCircleShapeOrder            uint
	NextCircleShape_stagedOrder     map[*NextCircleShape]uint
	NextCircleShape_orderStaged     map[uint]*NextCircleShape
	NextCircleShapes_reference      map[*NextCircleShape]*NextCircleShape
	NextCircleShapes_referenceOrder map[*NextCircleShape]uint

	// insertion point for slice of pointers maps
	OnAfterNextCircleShapeCreateCallback OnAfterCreateInterface[NextCircleShape]
	OnAfterNextCircleShapeUpdateCallback OnAfterUpdateInterface[NextCircleShape]
	OnAfterNextCircleShapeDeleteCallback OnAfterDeleteInterface[NextCircleShape]
	OnAfterNextCircleShapeReadCallback   OnAfterReadInterface[NextCircleShape]

	Plants                map[*Plant]struct{}
	Plants_instance       map[*Plant]*Plant
	Plants_mapString      map[string]*Plant
	PlantOrder            uint
	Plant_stagedOrder     map[*Plant]uint
	Plant_orderStaged     map[uint]*Plant
	Plants_reference      map[*Plant]*Plant
	Plants_referenceOrder map[*Plant]uint

	// insertion point for slice of pointers maps
	Plant_PlantDiagrams_reverseMap map[*PlantDiagram]*Plant

	OnAfterPlantCreateCallback OnAfterCreateInterface[Plant]
	OnAfterPlantUpdateCallback OnAfterUpdateInterface[Plant]
	OnAfterPlantDeleteCallback OnAfterDeleteInterface[Plant]
	OnAfterPlantReadCallback   OnAfterReadInterface[Plant]

	PlantCircumferenceShapes                map[*PlantCircumferenceShape]struct{}
	PlantCircumferenceShapes_instance       map[*PlantCircumferenceShape]*PlantCircumferenceShape
	PlantCircumferenceShapes_mapString      map[string]*PlantCircumferenceShape
	PlantCircumferenceShapeOrder            uint
	PlantCircumferenceShape_stagedOrder     map[*PlantCircumferenceShape]uint
	PlantCircumferenceShape_orderStaged     map[uint]*PlantCircumferenceShape
	PlantCircumferenceShapes_reference      map[*PlantCircumferenceShape]*PlantCircumferenceShape
	PlantCircumferenceShapes_referenceOrder map[*PlantCircumferenceShape]uint

	// insertion point for slice of pointers maps
	OnAfterPlantCircumferenceShapeCreateCallback OnAfterCreateInterface[PlantCircumferenceShape]
	OnAfterPlantCircumferenceShapeUpdateCallback OnAfterUpdateInterface[PlantCircumferenceShape]
	OnAfterPlantCircumferenceShapeDeleteCallback OnAfterDeleteInterface[PlantCircumferenceShape]
	OnAfterPlantCircumferenceShapeReadCallback   OnAfterReadInterface[PlantCircumferenceShape]

	PlantDiagrams                map[*PlantDiagram]struct{}
	PlantDiagrams_instance       map[*PlantDiagram]*PlantDiagram
	PlantDiagrams_mapString      map[string]*PlantDiagram
	PlantDiagramOrder            uint
	PlantDiagram_stagedOrder     map[*PlantDiagram]uint
	PlantDiagram_orderStaged     map[uint]*PlantDiagram
	PlantDiagrams_reference      map[*PlantDiagram]*PlantDiagram
	PlantDiagrams_referenceOrder map[*PlantDiagram]uint

	// insertion point for slice of pointers maps
	OnAfterPlantDiagramCreateCallback OnAfterCreateInterface[PlantDiagram]
	OnAfterPlantDiagramUpdateCallback OnAfterUpdateInterface[PlantDiagram]
	OnAfterPlantDiagramDeleteCallback OnAfterDeleteInterface[PlantDiagram]
	OnAfterPlantDiagramReadCallback   OnAfterReadInterface[PlantDiagram]

	RhombusGridShapes                map[*RhombusGridShape]struct{}
	RhombusGridShapes_instance       map[*RhombusGridShape]*RhombusGridShape
	RhombusGridShapes_mapString      map[string]*RhombusGridShape
	RhombusGridShapeOrder            uint
	RhombusGridShape_stagedOrder     map[*RhombusGridShape]uint
	RhombusGridShape_orderStaged     map[uint]*RhombusGridShape
	RhombusGridShapes_reference      map[*RhombusGridShape]*RhombusGridShape
	RhombusGridShapes_referenceOrder map[*RhombusGridShape]uint

	// insertion point for slice of pointers maps
	RhombusGridShape_RhombusShapes_reverseMap map[*RhombusShape]*RhombusGridShape

	OnAfterRhombusGridShapeCreateCallback OnAfterCreateInterface[RhombusGridShape]
	OnAfterRhombusGridShapeUpdateCallback OnAfterUpdateInterface[RhombusGridShape]
	OnAfterRhombusGridShapeDeleteCallback OnAfterDeleteInterface[RhombusGridShape]
	OnAfterRhombusGridShapeReadCallback   OnAfterReadInterface[RhombusGridShape]

	RhombusShapes                map[*RhombusShape]struct{}
	RhombusShapes_instance       map[*RhombusShape]*RhombusShape
	RhombusShapes_mapString      map[string]*RhombusShape
	RhombusShapeOrder            uint
	RhombusShape_stagedOrder     map[*RhombusShape]uint
	RhombusShape_orderStaged     map[uint]*RhombusShape
	RhombusShapes_reference      map[*RhombusShape]*RhombusShape
	RhombusShapes_referenceOrder map[*RhombusShape]uint

	// insertion point for slice of pointers maps
	OnAfterRhombusShapeCreateCallback OnAfterCreateInterface[RhombusShape]
	OnAfterRhombusShapeUpdateCallback OnAfterUpdateInterface[RhombusShape]
	OnAfterRhombusShapeDeleteCallback OnAfterDeleteInterface[RhombusShape]
	OnAfterRhombusShapeReadCallback   OnAfterReadInterface[RhombusShape]

	RotatedRhombusGridShapes                map[*RotatedRhombusGridShape]struct{}
	RotatedRhombusGridShapes_instance       map[*RotatedRhombusGridShape]*RotatedRhombusGridShape
	RotatedRhombusGridShapes_mapString      map[string]*RotatedRhombusGridShape
	RotatedRhombusGridShapeOrder            uint
	RotatedRhombusGridShape_stagedOrder     map[*RotatedRhombusGridShape]uint
	RotatedRhombusGridShape_orderStaged     map[uint]*RotatedRhombusGridShape
	RotatedRhombusGridShapes_reference      map[*RotatedRhombusGridShape]*RotatedRhombusGridShape
	RotatedRhombusGridShapes_referenceOrder map[*RotatedRhombusGridShape]uint

	// insertion point for slice of pointers maps
	RotatedRhombusGridShape_RotatedRhombusShapes_reverseMap map[*RotatedRhombusShape]*RotatedRhombusGridShape

	OnAfterRotatedRhombusGridShapeCreateCallback OnAfterCreateInterface[RotatedRhombusGridShape]
	OnAfterRotatedRhombusGridShapeUpdateCallback OnAfterUpdateInterface[RotatedRhombusGridShape]
	OnAfterRotatedRhombusGridShapeDeleteCallback OnAfterDeleteInterface[RotatedRhombusGridShape]
	OnAfterRotatedRhombusGridShapeReadCallback   OnAfterReadInterface[RotatedRhombusGridShape]

	RotatedRhombusShapes                map[*RotatedRhombusShape]struct{}
	RotatedRhombusShapes_instance       map[*RotatedRhombusShape]*RotatedRhombusShape
	RotatedRhombusShapes_mapString      map[string]*RotatedRhombusShape
	RotatedRhombusShapeOrder            uint
	RotatedRhombusShape_stagedOrder     map[*RotatedRhombusShape]uint
	RotatedRhombusShape_orderStaged     map[uint]*RotatedRhombusShape
	RotatedRhombusShapes_reference      map[*RotatedRhombusShape]*RotatedRhombusShape
	RotatedRhombusShapes_referenceOrder map[*RotatedRhombusShape]uint

	// insertion point for slice of pointers maps
	OnAfterRotatedRhombusShapeCreateCallback OnAfterCreateInterface[RotatedRhombusShape]
	OnAfterRotatedRhombusShapeUpdateCallback OnAfterUpdateInterface[RotatedRhombusShape]
	OnAfterRotatedRhombusShapeDeleteCallback OnAfterDeleteInterface[RotatedRhombusShape]
	OnAfterRotatedRhombusShapeReadCallback   OnAfterReadInterface[RotatedRhombusShape]

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
	stage.AxesShapes_reference = make(map[*AxesShape]*AxesShape)
	stage.AxesShapes_instance = make(map[*AxesShape]*AxesShape)
	stage.AxesShapes_referenceOrder = make(map[*AxesShape]uint)

	stage.CircleGridShapes_reference = make(map[*CircleGridShape]*CircleGridShape)
	stage.CircleGridShapes_instance = make(map[*CircleGridShape]*CircleGridShape)
	stage.CircleGridShapes_referenceOrder = make(map[*CircleGridShape]uint)

	stage.ExplanationTextShapes_reference = make(map[*ExplanationTextShape]*ExplanationTextShape)
	stage.ExplanationTextShapes_instance = make(map[*ExplanationTextShape]*ExplanationTextShape)
	stage.ExplanationTextShapes_referenceOrder = make(map[*ExplanationTextShape]uint)

	stage.GridPathShapes_reference = make(map[*GridPathShape]*GridPathShape)
	stage.GridPathShapes_instance = make(map[*GridPathShape]*GridPathShape)
	stage.GridPathShapes_referenceOrder = make(map[*GridPathShape]uint)

	stage.GrowthCurveRhombusGridShapes_reference = make(map[*GrowthCurveRhombusGridShape]*GrowthCurveRhombusGridShape)
	stage.GrowthCurveRhombusGridShapes_instance = make(map[*GrowthCurveRhombusGridShape]*GrowthCurveRhombusGridShape)
	stage.GrowthCurveRhombusGridShapes_referenceOrder = make(map[*GrowthCurveRhombusGridShape]uint)

	stage.GrowthCurveRhombusShapes_reference = make(map[*GrowthCurveRhombusShape]*GrowthCurveRhombusShape)
	stage.GrowthCurveRhombusShapes_instance = make(map[*GrowthCurveRhombusShape]*GrowthCurveRhombusShape)
	stage.GrowthCurveRhombusShapes_referenceOrder = make(map[*GrowthCurveRhombusShape]uint)

	stage.InitialRhombusGridShapes_reference = make(map[*InitialRhombusGridShape]*InitialRhombusGridShape)
	stage.InitialRhombusGridShapes_instance = make(map[*InitialRhombusGridShape]*InitialRhombusGridShape)
	stage.InitialRhombusGridShapes_referenceOrder = make(map[*InitialRhombusGridShape]uint)

	stage.InitialRhombusShapes_reference = make(map[*InitialRhombusShape]*InitialRhombusShape)
	stage.InitialRhombusShapes_instance = make(map[*InitialRhombusShape]*InitialRhombusShape)
	stage.InitialRhombusShapes_referenceOrder = make(map[*InitialRhombusShape]uint)

	stage.Librarys_reference = make(map[*Library]*Library)
	stage.Librarys_instance = make(map[*Library]*Library)
	stage.Librarys_referenceOrder = make(map[*Library]uint)

	stage.NextCircleShapes_reference = make(map[*NextCircleShape]*NextCircleShape)
	stage.NextCircleShapes_instance = make(map[*NextCircleShape]*NextCircleShape)
	stage.NextCircleShapes_referenceOrder = make(map[*NextCircleShape]uint)

	stage.Plants_reference = make(map[*Plant]*Plant)
	stage.Plants_instance = make(map[*Plant]*Plant)
	stage.Plants_referenceOrder = make(map[*Plant]uint)

	stage.PlantCircumferenceShapes_reference = make(map[*PlantCircumferenceShape]*PlantCircumferenceShape)
	stage.PlantCircumferenceShapes_instance = make(map[*PlantCircumferenceShape]*PlantCircumferenceShape)
	stage.PlantCircumferenceShapes_referenceOrder = make(map[*PlantCircumferenceShape]uint)

	stage.PlantDiagrams_reference = make(map[*PlantDiagram]*PlantDiagram)
	stage.PlantDiagrams_instance = make(map[*PlantDiagram]*PlantDiagram)
	stage.PlantDiagrams_referenceOrder = make(map[*PlantDiagram]uint)

	stage.RhombusGridShapes_reference = make(map[*RhombusGridShape]*RhombusGridShape)
	stage.RhombusGridShapes_instance = make(map[*RhombusGridShape]*RhombusGridShape)
	stage.RhombusGridShapes_referenceOrder = make(map[*RhombusGridShape]uint)

	stage.RhombusShapes_reference = make(map[*RhombusShape]*RhombusShape)
	stage.RhombusShapes_instance = make(map[*RhombusShape]*RhombusShape)
	stage.RhombusShapes_referenceOrder = make(map[*RhombusShape]uint)

	stage.RotatedRhombusGridShapes_reference = make(map[*RotatedRhombusGridShape]*RotatedRhombusGridShape)
	stage.RotatedRhombusGridShapes_instance = make(map[*RotatedRhombusGridShape]*RotatedRhombusGridShape)
	stage.RotatedRhombusGridShapes_referenceOrder = make(map[*RotatedRhombusGridShape]uint)

	stage.RotatedRhombusShapes_reference = make(map[*RotatedRhombusShape]*RotatedRhombusShape)
	stage.RotatedRhombusShapes_instance = make(map[*RotatedRhombusShape]*RotatedRhombusShape)
	stage.RotatedRhombusShapes_referenceOrder = make(map[*RotatedRhombusShape]uint)

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
	var maxAxesShapeOrder uint
	var foundAxesShape bool
	for _, order := range stage.AxesShape_stagedOrder {
		if !foundAxesShape || order > maxAxesShapeOrder {
			maxAxesShapeOrder = order
			foundAxesShape = true
		}
	}
	if foundAxesShape {
		stage.AxesShapeOrder = maxAxesShapeOrder + 1
	} else {
		stage.AxesShapeOrder = 0
	}

	var maxCircleGridShapeOrder uint
	var foundCircleGridShape bool
	for _, order := range stage.CircleGridShape_stagedOrder {
		if !foundCircleGridShape || order > maxCircleGridShapeOrder {
			maxCircleGridShapeOrder = order
			foundCircleGridShape = true
		}
	}
	if foundCircleGridShape {
		stage.CircleGridShapeOrder = maxCircleGridShapeOrder + 1
	} else {
		stage.CircleGridShapeOrder = 0
	}

	var maxExplanationTextShapeOrder uint
	var foundExplanationTextShape bool
	for _, order := range stage.ExplanationTextShape_stagedOrder {
		if !foundExplanationTextShape || order > maxExplanationTextShapeOrder {
			maxExplanationTextShapeOrder = order
			foundExplanationTextShape = true
		}
	}
	if foundExplanationTextShape {
		stage.ExplanationTextShapeOrder = maxExplanationTextShapeOrder + 1
	} else {
		stage.ExplanationTextShapeOrder = 0
	}

	var maxGridPathShapeOrder uint
	var foundGridPathShape bool
	for _, order := range stage.GridPathShape_stagedOrder {
		if !foundGridPathShape || order > maxGridPathShapeOrder {
			maxGridPathShapeOrder = order
			foundGridPathShape = true
		}
	}
	if foundGridPathShape {
		stage.GridPathShapeOrder = maxGridPathShapeOrder + 1
	} else {
		stage.GridPathShapeOrder = 0
	}

	var maxGrowthCurveRhombusGridShapeOrder uint
	var foundGrowthCurveRhombusGridShape bool
	for _, order := range stage.GrowthCurveRhombusGridShape_stagedOrder {
		if !foundGrowthCurveRhombusGridShape || order > maxGrowthCurveRhombusGridShapeOrder {
			maxGrowthCurveRhombusGridShapeOrder = order
			foundGrowthCurveRhombusGridShape = true
		}
	}
	if foundGrowthCurveRhombusGridShape {
		stage.GrowthCurveRhombusGridShapeOrder = maxGrowthCurveRhombusGridShapeOrder + 1
	} else {
		stage.GrowthCurveRhombusGridShapeOrder = 0
	}

	var maxGrowthCurveRhombusShapeOrder uint
	var foundGrowthCurveRhombusShape bool
	for _, order := range stage.GrowthCurveRhombusShape_stagedOrder {
		if !foundGrowthCurveRhombusShape || order > maxGrowthCurveRhombusShapeOrder {
			maxGrowthCurveRhombusShapeOrder = order
			foundGrowthCurveRhombusShape = true
		}
	}
	if foundGrowthCurveRhombusShape {
		stage.GrowthCurveRhombusShapeOrder = maxGrowthCurveRhombusShapeOrder + 1
	} else {
		stage.GrowthCurveRhombusShapeOrder = 0
	}

	var maxInitialRhombusGridShapeOrder uint
	var foundInitialRhombusGridShape bool
	for _, order := range stage.InitialRhombusGridShape_stagedOrder {
		if !foundInitialRhombusGridShape || order > maxInitialRhombusGridShapeOrder {
			maxInitialRhombusGridShapeOrder = order
			foundInitialRhombusGridShape = true
		}
	}
	if foundInitialRhombusGridShape {
		stage.InitialRhombusGridShapeOrder = maxInitialRhombusGridShapeOrder + 1
	} else {
		stage.InitialRhombusGridShapeOrder = 0
	}

	var maxInitialRhombusShapeOrder uint
	var foundInitialRhombusShape bool
	for _, order := range stage.InitialRhombusShape_stagedOrder {
		if !foundInitialRhombusShape || order > maxInitialRhombusShapeOrder {
			maxInitialRhombusShapeOrder = order
			foundInitialRhombusShape = true
		}
	}
	if foundInitialRhombusShape {
		stage.InitialRhombusShapeOrder = maxInitialRhombusShapeOrder + 1
	} else {
		stage.InitialRhombusShapeOrder = 0
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

	var maxNextCircleShapeOrder uint
	var foundNextCircleShape bool
	for _, order := range stage.NextCircleShape_stagedOrder {
		if !foundNextCircleShape || order > maxNextCircleShapeOrder {
			maxNextCircleShapeOrder = order
			foundNextCircleShape = true
		}
	}
	if foundNextCircleShape {
		stage.NextCircleShapeOrder = maxNextCircleShapeOrder + 1
	} else {
		stage.NextCircleShapeOrder = 0
	}

	var maxPlantOrder uint
	var foundPlant bool
	for _, order := range stage.Plant_stagedOrder {
		if !foundPlant || order > maxPlantOrder {
			maxPlantOrder = order
			foundPlant = true
		}
	}
	if foundPlant {
		stage.PlantOrder = maxPlantOrder + 1
	} else {
		stage.PlantOrder = 0
	}

	var maxPlantCircumferenceShapeOrder uint
	var foundPlantCircumferenceShape bool
	for _, order := range stage.PlantCircumferenceShape_stagedOrder {
		if !foundPlantCircumferenceShape || order > maxPlantCircumferenceShapeOrder {
			maxPlantCircumferenceShapeOrder = order
			foundPlantCircumferenceShape = true
		}
	}
	if foundPlantCircumferenceShape {
		stage.PlantCircumferenceShapeOrder = maxPlantCircumferenceShapeOrder + 1
	} else {
		stage.PlantCircumferenceShapeOrder = 0
	}

	var maxPlantDiagramOrder uint
	var foundPlantDiagram bool
	for _, order := range stage.PlantDiagram_stagedOrder {
		if !foundPlantDiagram || order > maxPlantDiagramOrder {
			maxPlantDiagramOrder = order
			foundPlantDiagram = true
		}
	}
	if foundPlantDiagram {
		stage.PlantDiagramOrder = maxPlantDiagramOrder + 1
	} else {
		stage.PlantDiagramOrder = 0
	}

	var maxRhombusGridShapeOrder uint
	var foundRhombusGridShape bool
	for _, order := range stage.RhombusGridShape_stagedOrder {
		if !foundRhombusGridShape || order > maxRhombusGridShapeOrder {
			maxRhombusGridShapeOrder = order
			foundRhombusGridShape = true
		}
	}
	if foundRhombusGridShape {
		stage.RhombusGridShapeOrder = maxRhombusGridShapeOrder + 1
	} else {
		stage.RhombusGridShapeOrder = 0
	}

	var maxRhombusShapeOrder uint
	var foundRhombusShape bool
	for _, order := range stage.RhombusShape_stagedOrder {
		if !foundRhombusShape || order > maxRhombusShapeOrder {
			maxRhombusShapeOrder = order
			foundRhombusShape = true
		}
	}
	if foundRhombusShape {
		stage.RhombusShapeOrder = maxRhombusShapeOrder + 1
	} else {
		stage.RhombusShapeOrder = 0
	}

	var maxRotatedRhombusGridShapeOrder uint
	var foundRotatedRhombusGridShape bool
	for _, order := range stage.RotatedRhombusGridShape_stagedOrder {
		if !foundRotatedRhombusGridShape || order > maxRotatedRhombusGridShapeOrder {
			maxRotatedRhombusGridShapeOrder = order
			foundRotatedRhombusGridShape = true
		}
	}
	if foundRotatedRhombusGridShape {
		stage.RotatedRhombusGridShapeOrder = maxRotatedRhombusGridShapeOrder + 1
	} else {
		stage.RotatedRhombusGridShapeOrder = 0
	}

	var maxRotatedRhombusShapeOrder uint
	var foundRotatedRhombusShape bool
	for _, order := range stage.RotatedRhombusShape_stagedOrder {
		if !foundRotatedRhombusShape || order > maxRotatedRhombusShapeOrder {
			maxRotatedRhombusShapeOrder = order
			foundRotatedRhombusShape = true
		}
	}
	if foundRotatedRhombusShape {
		stage.RotatedRhombusShapeOrder = maxRotatedRhombusShapeOrder + 1
	} else {
		stage.RotatedRhombusShapeOrder = 0
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
	case *AxesShape:
		tmp := GetStructInstancesByOrder(stage.AxesShapes, stage.AxesShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *AxesShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *CircleGridShape:
		tmp := GetStructInstancesByOrder(stage.CircleGridShapes, stage.CircleGridShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *CircleGridShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ExplanationTextShape:
		tmp := GetStructInstancesByOrder(stage.ExplanationTextShapes, stage.ExplanationTextShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ExplanationTextShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *GridPathShape:
		tmp := GetStructInstancesByOrder(stage.GridPathShapes, stage.GridPathShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GridPathShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *GrowthCurveRhombusGridShape:
		tmp := GetStructInstancesByOrder(stage.GrowthCurveRhombusGridShapes, stage.GrowthCurveRhombusGridShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GrowthCurveRhombusGridShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *GrowthCurveRhombusShape:
		tmp := GetStructInstancesByOrder(stage.GrowthCurveRhombusShapes, stage.GrowthCurveRhombusShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GrowthCurveRhombusShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *InitialRhombusGridShape:
		tmp := GetStructInstancesByOrder(stage.InitialRhombusGridShapes, stage.InitialRhombusGridShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *InitialRhombusGridShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *InitialRhombusShape:
		tmp := GetStructInstancesByOrder(stage.InitialRhombusShapes, stage.InitialRhombusShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *InitialRhombusShape implements.
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
	case *NextCircleShape:
		tmp := GetStructInstancesByOrder(stage.NextCircleShapes, stage.NextCircleShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *NextCircleShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Plant:
		tmp := GetStructInstancesByOrder(stage.Plants, stage.Plant_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Plant implements.
			res = append(res, any(v).(T))
		}
		return res
	case *PlantCircumferenceShape:
		tmp := GetStructInstancesByOrder(stage.PlantCircumferenceShapes, stage.PlantCircumferenceShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *PlantCircumferenceShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *PlantDiagram:
		tmp := GetStructInstancesByOrder(stage.PlantDiagrams, stage.PlantDiagram_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *PlantDiagram implements.
			res = append(res, any(v).(T))
		}
		return res
	case *RhombusGridShape:
		tmp := GetStructInstancesByOrder(stage.RhombusGridShapes, stage.RhombusGridShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *RhombusGridShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *RhombusShape:
		tmp := GetStructInstancesByOrder(stage.RhombusShapes, stage.RhombusShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *RhombusShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *RotatedRhombusGridShape:
		tmp := GetStructInstancesByOrder(stage.RotatedRhombusGridShapes, stage.RotatedRhombusGridShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *RotatedRhombusGridShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *RotatedRhombusShape:
		tmp := GetStructInstancesByOrder(stage.RotatedRhombusShapes, stage.RotatedRhombusShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *RotatedRhombusShape implements.
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
	case "AxesShape":
		res = GetNamedStructInstances(stage.AxesShapes, stage.AxesShape_stagedOrder)
	case "CircleGridShape":
		res = GetNamedStructInstances(stage.CircleGridShapes, stage.CircleGridShape_stagedOrder)
	case "ExplanationTextShape":
		res = GetNamedStructInstances(stage.ExplanationTextShapes, stage.ExplanationTextShape_stagedOrder)
	case "GridPathShape":
		res = GetNamedStructInstances(stage.GridPathShapes, stage.GridPathShape_stagedOrder)
	case "GrowthCurveRhombusGridShape":
		res = GetNamedStructInstances(stage.GrowthCurveRhombusGridShapes, stage.GrowthCurveRhombusGridShape_stagedOrder)
	case "GrowthCurveRhombusShape":
		res = GetNamedStructInstances(stage.GrowthCurveRhombusShapes, stage.GrowthCurveRhombusShape_stagedOrder)
	case "InitialRhombusGridShape":
		res = GetNamedStructInstances(stage.InitialRhombusGridShapes, stage.InitialRhombusGridShape_stagedOrder)
	case "InitialRhombusShape":
		res = GetNamedStructInstances(stage.InitialRhombusShapes, stage.InitialRhombusShape_stagedOrder)
	case "Library":
		res = GetNamedStructInstances(stage.Librarys, stage.Library_stagedOrder)
	case "NextCircleShape":
		res = GetNamedStructInstances(stage.NextCircleShapes, stage.NextCircleShape_stagedOrder)
	case "Plant":
		res = GetNamedStructInstances(stage.Plants, stage.Plant_stagedOrder)
	case "PlantCircumferenceShape":
		res = GetNamedStructInstances(stage.PlantCircumferenceShapes, stage.PlantCircumferenceShape_stagedOrder)
	case "PlantDiagram":
		res = GetNamedStructInstances(stage.PlantDiagrams, stage.PlantDiagram_stagedOrder)
	case "RhombusGridShape":
		res = GetNamedStructInstances(stage.RhombusGridShapes, stage.RhombusGridShape_stagedOrder)
	case "RhombusShape":
		res = GetNamedStructInstances(stage.RhombusShapes, stage.RhombusShape_stagedOrder)
	case "RotatedRhombusGridShape":
		res = GetNamedStructInstances(stage.RotatedRhombusGridShapes, stage.RotatedRhombusGridShape_stagedOrder)
	case "RotatedRhombusShape":
		res = GetNamedStructInstances(stage.RotatedRhombusShapes, stage.RotatedRhombusShape_stagedOrder)
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
	return "github.com/fullstack-lang/gong/dsm/phylla/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return phylla_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return phylla_go.GoDiagramsDir
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
	CommitAxesShape(axesshape *AxesShape)
	CheckoutAxesShape(axesshape *AxesShape)
	CommitCircleGridShape(circlegridshape *CircleGridShape)
	CheckoutCircleGridShape(circlegridshape *CircleGridShape)
	CommitExplanationTextShape(explanationtextshape *ExplanationTextShape)
	CheckoutExplanationTextShape(explanationtextshape *ExplanationTextShape)
	CommitGridPathShape(gridpathshape *GridPathShape)
	CheckoutGridPathShape(gridpathshape *GridPathShape)
	CommitGrowthCurveRhombusGridShape(growthcurverhombusgridshape *GrowthCurveRhombusGridShape)
	CheckoutGrowthCurveRhombusGridShape(growthcurverhombusgridshape *GrowthCurveRhombusGridShape)
	CommitGrowthCurveRhombusShape(growthcurverhombusshape *GrowthCurveRhombusShape)
	CheckoutGrowthCurveRhombusShape(growthcurverhombusshape *GrowthCurveRhombusShape)
	CommitInitialRhombusGridShape(initialrhombusgridshape *InitialRhombusGridShape)
	CheckoutInitialRhombusGridShape(initialrhombusgridshape *InitialRhombusGridShape)
	CommitInitialRhombusShape(initialrhombusshape *InitialRhombusShape)
	CheckoutInitialRhombusShape(initialrhombusshape *InitialRhombusShape)
	CommitLibrary(library *Library)
	CheckoutLibrary(library *Library)
	CommitNextCircleShape(nextcircleshape *NextCircleShape)
	CheckoutNextCircleShape(nextcircleshape *NextCircleShape)
	CommitPlant(plant *Plant)
	CheckoutPlant(plant *Plant)
	CommitPlantCircumferenceShape(plantcircumferenceshape *PlantCircumferenceShape)
	CheckoutPlantCircumferenceShape(plantcircumferenceshape *PlantCircumferenceShape)
	CommitPlantDiagram(plantdiagram *PlantDiagram)
	CheckoutPlantDiagram(plantdiagram *PlantDiagram)
	CommitRhombusGridShape(rhombusgridshape *RhombusGridShape)
	CheckoutRhombusGridShape(rhombusgridshape *RhombusGridShape)
	CommitRhombusShape(rhombusshape *RhombusShape)
	CheckoutRhombusShape(rhombusshape *RhombusShape)
	CommitRotatedRhombusGridShape(rotatedrhombusgridshape *RotatedRhombusGridShape)
	CheckoutRotatedRhombusGridShape(rotatedrhombusgridshape *RotatedRhombusGridShape)
	CommitRotatedRhombusShape(rotatedrhombusshape *RotatedRhombusShape)
	CheckoutRotatedRhombusShape(rotatedrhombusshape *RotatedRhombusShape)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		AxesShapes:           make(map[*AxesShape]struct{}),
		AxesShapes_mapString: make(map[string]*AxesShape),

		CircleGridShapes:           make(map[*CircleGridShape]struct{}),
		CircleGridShapes_mapString: make(map[string]*CircleGridShape),

		ExplanationTextShapes:           make(map[*ExplanationTextShape]struct{}),
		ExplanationTextShapes_mapString: make(map[string]*ExplanationTextShape),

		GridPathShapes:           make(map[*GridPathShape]struct{}),
		GridPathShapes_mapString: make(map[string]*GridPathShape),

		GrowthCurveRhombusGridShapes:           make(map[*GrowthCurveRhombusGridShape]struct{}),
		GrowthCurveRhombusGridShapes_mapString: make(map[string]*GrowthCurveRhombusGridShape),

		GrowthCurveRhombusShapes:           make(map[*GrowthCurveRhombusShape]struct{}),
		GrowthCurveRhombusShapes_mapString: make(map[string]*GrowthCurveRhombusShape),

		InitialRhombusGridShapes:           make(map[*InitialRhombusGridShape]struct{}),
		InitialRhombusGridShapes_mapString: make(map[string]*InitialRhombusGridShape),

		InitialRhombusShapes:           make(map[*InitialRhombusShape]struct{}),
		InitialRhombusShapes_mapString: make(map[string]*InitialRhombusShape),

		Librarys:           make(map[*Library]struct{}),
		Librarys_mapString: make(map[string]*Library),

		NextCircleShapes:           make(map[*NextCircleShape]struct{}),
		NextCircleShapes_mapString: make(map[string]*NextCircleShape),

		Plants:           make(map[*Plant]struct{}),
		Plants_mapString: make(map[string]*Plant),

		PlantCircumferenceShapes:           make(map[*PlantCircumferenceShape]struct{}),
		PlantCircumferenceShapes_mapString: make(map[string]*PlantCircumferenceShape),

		PlantDiagrams:           make(map[*PlantDiagram]struct{}),
		PlantDiagrams_mapString: make(map[string]*PlantDiagram),

		RhombusGridShapes:           make(map[*RhombusGridShape]struct{}),
		RhombusGridShapes_mapString: make(map[string]*RhombusGridShape),

		RhombusShapes:           make(map[*RhombusShape]struct{}),
		RhombusShapes_mapString: make(map[string]*RhombusShape),

		RotatedRhombusGridShapes:           make(map[*RotatedRhombusGridShape]struct{}),
		RotatedRhombusGridShapes_mapString: make(map[string]*RotatedRhombusGridShape),

		RotatedRhombusShapes:           make(map[*RotatedRhombusShape]struct{}),
		RotatedRhombusShapes_mapString: make(map[string]*RotatedRhombusShape),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		AxesShape_stagedOrder: make(map[*AxesShape]uint),
		AxesShape_orderStaged: make(map[uint]*AxesShape),
		AxesShapes_reference:  make(map[*AxesShape]*AxesShape),

		CircleGridShape_stagedOrder: make(map[*CircleGridShape]uint),
		CircleGridShape_orderStaged: make(map[uint]*CircleGridShape),
		CircleGridShapes_reference:  make(map[*CircleGridShape]*CircleGridShape),

		ExplanationTextShape_stagedOrder: make(map[*ExplanationTextShape]uint),
		ExplanationTextShape_orderStaged: make(map[uint]*ExplanationTextShape),
		ExplanationTextShapes_reference:  make(map[*ExplanationTextShape]*ExplanationTextShape),

		GridPathShape_stagedOrder: make(map[*GridPathShape]uint),
		GridPathShape_orderStaged: make(map[uint]*GridPathShape),
		GridPathShapes_reference:  make(map[*GridPathShape]*GridPathShape),

		GrowthCurveRhombusGridShape_stagedOrder: make(map[*GrowthCurveRhombusGridShape]uint),
		GrowthCurveRhombusGridShape_orderStaged: make(map[uint]*GrowthCurveRhombusGridShape),
		GrowthCurveRhombusGridShapes_reference:  make(map[*GrowthCurveRhombusGridShape]*GrowthCurveRhombusGridShape),

		GrowthCurveRhombusShape_stagedOrder: make(map[*GrowthCurveRhombusShape]uint),
		GrowthCurveRhombusShape_orderStaged: make(map[uint]*GrowthCurveRhombusShape),
		GrowthCurveRhombusShapes_reference:  make(map[*GrowthCurveRhombusShape]*GrowthCurveRhombusShape),

		InitialRhombusGridShape_stagedOrder: make(map[*InitialRhombusGridShape]uint),
		InitialRhombusGridShape_orderStaged: make(map[uint]*InitialRhombusGridShape),
		InitialRhombusGridShapes_reference:  make(map[*InitialRhombusGridShape]*InitialRhombusGridShape),

		InitialRhombusShape_stagedOrder: make(map[*InitialRhombusShape]uint),
		InitialRhombusShape_orderStaged: make(map[uint]*InitialRhombusShape),
		InitialRhombusShapes_reference:  make(map[*InitialRhombusShape]*InitialRhombusShape),

		Library_stagedOrder: make(map[*Library]uint),
		Library_orderStaged: make(map[uint]*Library),
		Librarys_reference:  make(map[*Library]*Library),

		NextCircleShape_stagedOrder: make(map[*NextCircleShape]uint),
		NextCircleShape_orderStaged: make(map[uint]*NextCircleShape),
		NextCircleShapes_reference:  make(map[*NextCircleShape]*NextCircleShape),

		Plant_stagedOrder: make(map[*Plant]uint),
		Plant_orderStaged: make(map[uint]*Plant),
		Plants_reference:  make(map[*Plant]*Plant),

		PlantCircumferenceShape_stagedOrder: make(map[*PlantCircumferenceShape]uint),
		PlantCircumferenceShape_orderStaged: make(map[uint]*PlantCircumferenceShape),
		PlantCircumferenceShapes_reference:  make(map[*PlantCircumferenceShape]*PlantCircumferenceShape),

		PlantDiagram_stagedOrder: make(map[*PlantDiagram]uint),
		PlantDiagram_orderStaged: make(map[uint]*PlantDiagram),
		PlantDiagrams_reference:  make(map[*PlantDiagram]*PlantDiagram),

		RhombusGridShape_stagedOrder: make(map[*RhombusGridShape]uint),
		RhombusGridShape_orderStaged: make(map[uint]*RhombusGridShape),
		RhombusGridShapes_reference:  make(map[*RhombusGridShape]*RhombusGridShape),

		RhombusShape_stagedOrder: make(map[*RhombusShape]uint),
		RhombusShape_orderStaged: make(map[uint]*RhombusShape),
		RhombusShapes_reference:  make(map[*RhombusShape]*RhombusShape),

		RotatedRhombusGridShape_stagedOrder: make(map[*RotatedRhombusGridShape]uint),
		RotatedRhombusGridShape_orderStaged: make(map[uint]*RotatedRhombusGridShape),
		RotatedRhombusGridShapes_reference:  make(map[*RotatedRhombusGridShape]*RotatedRhombusGridShape),

		RotatedRhombusShape_stagedOrder: make(map[*RotatedRhombusShape]uint),
		RotatedRhombusShape_orderStaged: make(map[uint]*RotatedRhombusShape),
		RotatedRhombusShapes_reference:  make(map[*RotatedRhombusShape]*RotatedRhombusShape),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"AxesShape": &AxesShapeUnmarshaller{},

			"CircleGridShape": &CircleGridShapeUnmarshaller{},

			"ExplanationTextShape": &ExplanationTextShapeUnmarshaller{},

			"GridPathShape": &GridPathShapeUnmarshaller{},

			"GrowthCurveRhombusGridShape": &GrowthCurveRhombusGridShapeUnmarshaller{},

			"GrowthCurveRhombusShape": &GrowthCurveRhombusShapeUnmarshaller{},

			"InitialRhombusGridShape": &InitialRhombusGridShapeUnmarshaller{},

			"InitialRhombusShape": &InitialRhombusShapeUnmarshaller{},

			"Library": &LibraryUnmarshaller{},

			"NextCircleShape": &NextCircleShapeUnmarshaller{},

			"Plant": &PlantUnmarshaller{},

			"PlantCircumferenceShape": &PlantCircumferenceShapeUnmarshaller{},

			"PlantDiagram": &PlantDiagramUnmarshaller{},

			"RhombusGridShape": &RhombusGridShapeUnmarshaller{},

			"RhombusShape": &RhombusShapeUnmarshaller{},

			"RotatedRhombusGridShape": &RotatedRhombusGridShapeUnmarshaller{},

			"RotatedRhombusShape": &RotatedRhombusShapeUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "AxesShape"},
			{name: "CircleGridShape"},
			{name: "ExplanationTextShape"},
			{name: "GridPathShape"},
			{name: "GrowthCurveRhombusGridShape"},
			{name: "GrowthCurveRhombusShape"},
			{name: "InitialRhombusGridShape"},
			{name: "InitialRhombusShape"},
			{name: "Library"},
			{name: "NextCircleShape"},
			{name: "Plant"},
			{name: "PlantCircumferenceShape"},
			{name: "PlantDiagram"},
			{name: "RhombusGridShape"},
			{name: "RhombusShape"},
			{name: "RotatedRhombusGridShape"},
			{name: "RotatedRhombusShape"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *AxesShape:
		return stage.AxesShape_stagedOrder[instance]
	case *CircleGridShape:
		return stage.CircleGridShape_stagedOrder[instance]
	case *ExplanationTextShape:
		return stage.ExplanationTextShape_stagedOrder[instance]
	case *GridPathShape:
		return stage.GridPathShape_stagedOrder[instance]
	case *GrowthCurveRhombusGridShape:
		return stage.GrowthCurveRhombusGridShape_stagedOrder[instance]
	case *GrowthCurveRhombusShape:
		return stage.GrowthCurveRhombusShape_stagedOrder[instance]
	case *InitialRhombusGridShape:
		return stage.InitialRhombusGridShape_stagedOrder[instance]
	case *InitialRhombusShape:
		return stage.InitialRhombusShape_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *NextCircleShape:
		return stage.NextCircleShape_stagedOrder[instance]
	case *Plant:
		return stage.Plant_stagedOrder[instance]
	case *PlantCircumferenceShape:
		return stage.PlantCircumferenceShape_stagedOrder[instance]
	case *PlantDiagram:
		return stage.PlantDiagram_stagedOrder[instance]
	case *RhombusGridShape:
		return stage.RhombusGridShape_stagedOrder[instance]
	case *RhombusShape:
		return stage.RhombusShape_stagedOrder[instance]
	case *RotatedRhombusGridShape:
		return stage.RotatedRhombusGridShape_stagedOrder[instance]
	case *RotatedRhombusShape:
		return stage.RotatedRhombusShape_stagedOrder[instance]
	default:
		return 0 // should not happen
	}
}

func GongGetInstanceFromOrder[Type PointerToGongstruct](stage *Stage, order uint) (res Type) {
	var t Type
	switch any(t).(type) {
	// insertion point for order map initialisations
	case *AxesShape:
		return any(stage.AxesShape_orderStaged[order]).(Type)
	case *CircleGridShape:
		return any(stage.CircleGridShape_orderStaged[order]).(Type)
	case *ExplanationTextShape:
		return any(stage.ExplanationTextShape_orderStaged[order]).(Type)
	case *GridPathShape:
		return any(stage.GridPathShape_orderStaged[order]).(Type)
	case *GrowthCurveRhombusGridShape:
		return any(stage.GrowthCurveRhombusGridShape_orderStaged[order]).(Type)
	case *GrowthCurveRhombusShape:
		return any(stage.GrowthCurveRhombusShape_orderStaged[order]).(Type)
	case *InitialRhombusGridShape:
		return any(stage.InitialRhombusGridShape_orderStaged[order]).(Type)
	case *InitialRhombusShape:
		return any(stage.InitialRhombusShape_orderStaged[order]).(Type)
	case *Library:
		return any(stage.Library_orderStaged[order]).(Type)
	case *NextCircleShape:
		return any(stage.NextCircleShape_orderStaged[order]).(Type)
	case *Plant:
		return any(stage.Plant_orderStaged[order]).(Type)
	case *PlantCircumferenceShape:
		return any(stage.PlantCircumferenceShape_orderStaged[order]).(Type)
	case *PlantDiagram:
		return any(stage.PlantDiagram_orderStaged[order]).(Type)
	case *RhombusGridShape:
		return any(stage.RhombusGridShape_orderStaged[order]).(Type)
	case *RhombusShape:
		return any(stage.RhombusShape_orderStaged[order]).(Type)
	case *RotatedRhombusGridShape:
		return any(stage.RotatedRhombusGridShape_orderStaged[order]).(Type)
	case *RotatedRhombusShape:
		return any(stage.RotatedRhombusShape_orderStaged[order]).(Type)
	default:
		return // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *AxesShape:
		return stage.AxesShape_stagedOrder[instance]
	case *CircleGridShape:
		return stage.CircleGridShape_stagedOrder[instance]
	case *ExplanationTextShape:
		return stage.ExplanationTextShape_stagedOrder[instance]
	case *GridPathShape:
		return stage.GridPathShape_stagedOrder[instance]
	case *GrowthCurveRhombusGridShape:
		return stage.GrowthCurveRhombusGridShape_stagedOrder[instance]
	case *GrowthCurveRhombusShape:
		return stage.GrowthCurveRhombusShape_stagedOrder[instance]
	case *InitialRhombusGridShape:
		return stage.InitialRhombusGridShape_stagedOrder[instance]
	case *InitialRhombusShape:
		return stage.InitialRhombusShape_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *NextCircleShape:
		return stage.NextCircleShape_stagedOrder[instance]
	case *Plant:
		return stage.Plant_stagedOrder[instance]
	case *PlantCircumferenceShape:
		return stage.PlantCircumferenceShape_stagedOrder[instance]
	case *PlantDiagram:
		return stage.PlantDiagram_stagedOrder[instance]
	case *RhombusGridShape:
		return stage.RhombusGridShape_stagedOrder[instance]
	case *RhombusShape:
		return stage.RhombusShape_stagedOrder[instance]
	case *RotatedRhombusGridShape:
		return stage.RotatedRhombusGridShape_stagedOrder[instance]
	case *RotatedRhombusShape:
		return stage.RotatedRhombusShape_stagedOrder[instance]
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
	stage.Map_GongStructName_InstancesNb["AxesShape"] = len(stage.AxesShapes)
	stage.Map_GongStructName_InstancesNb["CircleGridShape"] = len(stage.CircleGridShapes)
	stage.Map_GongStructName_InstancesNb["ExplanationTextShape"] = len(stage.ExplanationTextShapes)
	stage.Map_GongStructName_InstancesNb["GridPathShape"] = len(stage.GridPathShapes)
	stage.Map_GongStructName_InstancesNb["GrowthCurveRhombusGridShape"] = len(stage.GrowthCurveRhombusGridShapes)
	stage.Map_GongStructName_InstancesNb["GrowthCurveRhombusShape"] = len(stage.GrowthCurveRhombusShapes)
	stage.Map_GongStructName_InstancesNb["InitialRhombusGridShape"] = len(stage.InitialRhombusGridShapes)
	stage.Map_GongStructName_InstancesNb["InitialRhombusShape"] = len(stage.InitialRhombusShapes)
	stage.Map_GongStructName_InstancesNb["Library"] = len(stage.Librarys)
	stage.Map_GongStructName_InstancesNb["NextCircleShape"] = len(stage.NextCircleShapes)
	stage.Map_GongStructName_InstancesNb["Plant"] = len(stage.Plants)
	stage.Map_GongStructName_InstancesNb["PlantCircumferenceShape"] = len(stage.PlantCircumferenceShapes)
	stage.Map_GongStructName_InstancesNb["PlantDiagram"] = len(stage.PlantDiagrams)
	stage.Map_GongStructName_InstancesNb["RhombusGridShape"] = len(stage.RhombusGridShapes)
	stage.Map_GongStructName_InstancesNb["RhombusShape"] = len(stage.RhombusShapes)
	stage.Map_GongStructName_InstancesNb["RotatedRhombusGridShape"] = len(stage.RotatedRhombusGridShapes)
	stage.Map_GongStructName_InstancesNb["RotatedRhombusShape"] = len(stage.RotatedRhombusShapes)
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
// Stage puts axesshape to the model stage
func (axesshape *AxesShape) Stage(stage *Stage) *AxesShape {
	if _, ok := stage.AxesShapes[axesshape]; !ok {
		stage.AxesShapes[axesshape] = struct{}{}
		stage.AxesShape_stagedOrder[axesshape] = stage.AxesShapeOrder
		stage.AxesShape_orderStaged[stage.AxesShapeOrder] = axesshape
		stage.AxesShapeOrder++
	}
	stage.AxesShapes_mapString[axesshape.Name] = axesshape

	return axesshape
}

// StagePreserveOrder puts axesshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AxesShapeOrder
// - update stage.AxesShapeOrder accordingly
func (axesshape *AxesShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.AxesShapes[axesshape]; !ok {
		stage.AxesShapes[axesshape] = struct{}{}

		if order > stage.AxesShapeOrder {
			stage.AxesShapeOrder = order
		}
		stage.AxesShape_stagedOrder[axesshape] = order
		stage.AxesShape_orderStaged[order] = axesshape
		stage.AxesShapeOrder++
	}
	stage.AxesShapes_mapString[axesshape.Name] = axesshape
}

// Unstage removes axesshape off the model stage
func (axesshape *AxesShape) Unstage(stage *Stage) *AxesShape {
	delete(stage.AxesShapes, axesshape)
	// issue1150
	// delete(stage.AxesShape_stagedOrder, axesshape)
	delete(stage.AxesShapes_mapString, axesshape.Name)

	return axesshape
}

// UnstageVoid removes axesshape off the model stage
func (axesshape *AxesShape) UnstageVoid(stage *Stage) {
	delete(stage.AxesShapes, axesshape)
	// issue1150
	// delete(stage.AxesShape_stagedOrder, axesshape)
	delete(stage.AxesShapes_mapString, axesshape.Name)
}

// commit axesshape to the back repo (if it is already staged)
func (axesshape *AxesShape) Commit(stage *Stage) *AxesShape {
	if _, ok := stage.AxesShapes[axesshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAxesShape(axesshape)
		}
	}
	return axesshape
}

func (axesshape *AxesShape) CommitVoid(stage *Stage) {
	axesshape.Commit(stage)
}

func (axesshape *AxesShape) StageVoid(stage *Stage) {
	axesshape.Stage(stage)
}

// Checkout axesshape to the back repo (if it is already staged)
func (axesshape *AxesShape) Checkout(stage *Stage) *AxesShape {
	if _, ok := stage.AxesShapes[axesshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAxesShape(axesshape)
		}
	}
	return axesshape
}

// for satisfaction of GongStruct interface
func (axesshape *AxesShape) GetName() (res string) {
	return axesshape.Name
}

// for satisfaction of GongStruct interface
func (axesshape *AxesShape) SetName(name string) {
	axesshape.Name = name
}

// Stage puts circlegridshape to the model stage
func (circlegridshape *CircleGridShape) Stage(stage *Stage) *CircleGridShape {
	if _, ok := stage.CircleGridShapes[circlegridshape]; !ok {
		stage.CircleGridShapes[circlegridshape] = struct{}{}
		stage.CircleGridShape_stagedOrder[circlegridshape] = stage.CircleGridShapeOrder
		stage.CircleGridShape_orderStaged[stage.CircleGridShapeOrder] = circlegridshape
		stage.CircleGridShapeOrder++
	}
	stage.CircleGridShapes_mapString[circlegridshape.Name] = circlegridshape

	return circlegridshape
}

// StagePreserveOrder puts circlegridshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CircleGridShapeOrder
// - update stage.CircleGridShapeOrder accordingly
func (circlegridshape *CircleGridShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.CircleGridShapes[circlegridshape]; !ok {
		stage.CircleGridShapes[circlegridshape] = struct{}{}

		if order > stage.CircleGridShapeOrder {
			stage.CircleGridShapeOrder = order
		}
		stage.CircleGridShape_stagedOrder[circlegridshape] = order
		stage.CircleGridShape_orderStaged[order] = circlegridshape
		stage.CircleGridShapeOrder++
	}
	stage.CircleGridShapes_mapString[circlegridshape.Name] = circlegridshape
}

// Unstage removes circlegridshape off the model stage
func (circlegridshape *CircleGridShape) Unstage(stage *Stage) *CircleGridShape {
	delete(stage.CircleGridShapes, circlegridshape)
	// issue1150
	// delete(stage.CircleGridShape_stagedOrder, circlegridshape)
	delete(stage.CircleGridShapes_mapString, circlegridshape.Name)

	return circlegridshape
}

// UnstageVoid removes circlegridshape off the model stage
func (circlegridshape *CircleGridShape) UnstageVoid(stage *Stage) {
	delete(stage.CircleGridShapes, circlegridshape)
	// issue1150
	// delete(stage.CircleGridShape_stagedOrder, circlegridshape)
	delete(stage.CircleGridShapes_mapString, circlegridshape.Name)
}

// commit circlegridshape to the back repo (if it is already staged)
func (circlegridshape *CircleGridShape) Commit(stage *Stage) *CircleGridShape {
	if _, ok := stage.CircleGridShapes[circlegridshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCircleGridShape(circlegridshape)
		}
	}
	return circlegridshape
}

func (circlegridshape *CircleGridShape) CommitVoid(stage *Stage) {
	circlegridshape.Commit(stage)
}

func (circlegridshape *CircleGridShape) StageVoid(stage *Stage) {
	circlegridshape.Stage(stage)
}

// Checkout circlegridshape to the back repo (if it is already staged)
func (circlegridshape *CircleGridShape) Checkout(stage *Stage) *CircleGridShape {
	if _, ok := stage.CircleGridShapes[circlegridshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCircleGridShape(circlegridshape)
		}
	}
	return circlegridshape
}

// for satisfaction of GongStruct interface
func (circlegridshape *CircleGridShape) GetName() (res string) {
	return circlegridshape.Name
}

// for satisfaction of GongStruct interface
func (circlegridshape *CircleGridShape) SetName(name string) {
	circlegridshape.Name = name
}

// Stage puts explanationtextshape to the model stage
func (explanationtextshape *ExplanationTextShape) Stage(stage *Stage) *ExplanationTextShape {
	if _, ok := stage.ExplanationTextShapes[explanationtextshape]; !ok {
		stage.ExplanationTextShapes[explanationtextshape] = struct{}{}
		stage.ExplanationTextShape_stagedOrder[explanationtextshape] = stage.ExplanationTextShapeOrder
		stage.ExplanationTextShape_orderStaged[stage.ExplanationTextShapeOrder] = explanationtextshape
		stage.ExplanationTextShapeOrder++
	}
	stage.ExplanationTextShapes_mapString[explanationtextshape.Name] = explanationtextshape

	return explanationtextshape
}

// StagePreserveOrder puts explanationtextshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ExplanationTextShapeOrder
// - update stage.ExplanationTextShapeOrder accordingly
func (explanationtextshape *ExplanationTextShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ExplanationTextShapes[explanationtextshape]; !ok {
		stage.ExplanationTextShapes[explanationtextshape] = struct{}{}

		if order > stage.ExplanationTextShapeOrder {
			stage.ExplanationTextShapeOrder = order
		}
		stage.ExplanationTextShape_stagedOrder[explanationtextshape] = order
		stage.ExplanationTextShape_orderStaged[order] = explanationtextshape
		stage.ExplanationTextShapeOrder++
	}
	stage.ExplanationTextShapes_mapString[explanationtextshape.Name] = explanationtextshape
}

// Unstage removes explanationtextshape off the model stage
func (explanationtextshape *ExplanationTextShape) Unstage(stage *Stage) *ExplanationTextShape {
	delete(stage.ExplanationTextShapes, explanationtextshape)
	// issue1150
	// delete(stage.ExplanationTextShape_stagedOrder, explanationtextshape)
	delete(stage.ExplanationTextShapes_mapString, explanationtextshape.Name)

	return explanationtextshape
}

// UnstageVoid removes explanationtextshape off the model stage
func (explanationtextshape *ExplanationTextShape) UnstageVoid(stage *Stage) {
	delete(stage.ExplanationTextShapes, explanationtextshape)
	// issue1150
	// delete(stage.ExplanationTextShape_stagedOrder, explanationtextshape)
	delete(stage.ExplanationTextShapes_mapString, explanationtextshape.Name)
}

// commit explanationtextshape to the back repo (if it is already staged)
func (explanationtextshape *ExplanationTextShape) Commit(stage *Stage) *ExplanationTextShape {
	if _, ok := stage.ExplanationTextShapes[explanationtextshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitExplanationTextShape(explanationtextshape)
		}
	}
	return explanationtextshape
}

func (explanationtextshape *ExplanationTextShape) CommitVoid(stage *Stage) {
	explanationtextshape.Commit(stage)
}

func (explanationtextshape *ExplanationTextShape) StageVoid(stage *Stage) {
	explanationtextshape.Stage(stage)
}

// Checkout explanationtextshape to the back repo (if it is already staged)
func (explanationtextshape *ExplanationTextShape) Checkout(stage *Stage) *ExplanationTextShape {
	if _, ok := stage.ExplanationTextShapes[explanationtextshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutExplanationTextShape(explanationtextshape)
		}
	}
	return explanationtextshape
}

// for satisfaction of GongStruct interface
func (explanationtextshape *ExplanationTextShape) GetName() (res string) {
	return explanationtextshape.Name
}

// for satisfaction of GongStruct interface
func (explanationtextshape *ExplanationTextShape) SetName(name string) {
	explanationtextshape.Name = name
}

// Stage puts gridpathshape to the model stage
func (gridpathshape *GridPathShape) Stage(stage *Stage) *GridPathShape {
	if _, ok := stage.GridPathShapes[gridpathshape]; !ok {
		stage.GridPathShapes[gridpathshape] = struct{}{}
		stage.GridPathShape_stagedOrder[gridpathshape] = stage.GridPathShapeOrder
		stage.GridPathShape_orderStaged[stage.GridPathShapeOrder] = gridpathshape
		stage.GridPathShapeOrder++
	}
	stage.GridPathShapes_mapString[gridpathshape.Name] = gridpathshape

	return gridpathshape
}

// StagePreserveOrder puts gridpathshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GridPathShapeOrder
// - update stage.GridPathShapeOrder accordingly
func (gridpathshape *GridPathShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.GridPathShapes[gridpathshape]; !ok {
		stage.GridPathShapes[gridpathshape] = struct{}{}

		if order > stage.GridPathShapeOrder {
			stage.GridPathShapeOrder = order
		}
		stage.GridPathShape_stagedOrder[gridpathshape] = order
		stage.GridPathShape_orderStaged[order] = gridpathshape
		stage.GridPathShapeOrder++
	}
	stage.GridPathShapes_mapString[gridpathshape.Name] = gridpathshape
}

// Unstage removes gridpathshape off the model stage
func (gridpathshape *GridPathShape) Unstage(stage *Stage) *GridPathShape {
	delete(stage.GridPathShapes, gridpathshape)
	// issue1150
	// delete(stage.GridPathShape_stagedOrder, gridpathshape)
	delete(stage.GridPathShapes_mapString, gridpathshape.Name)

	return gridpathshape
}

// UnstageVoid removes gridpathshape off the model stage
func (gridpathshape *GridPathShape) UnstageVoid(stage *Stage) {
	delete(stage.GridPathShapes, gridpathshape)
	// issue1150
	// delete(stage.GridPathShape_stagedOrder, gridpathshape)
	delete(stage.GridPathShapes_mapString, gridpathshape.Name)
}

// commit gridpathshape to the back repo (if it is already staged)
func (gridpathshape *GridPathShape) Commit(stage *Stage) *GridPathShape {
	if _, ok := stage.GridPathShapes[gridpathshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGridPathShape(gridpathshape)
		}
	}
	return gridpathshape
}

func (gridpathshape *GridPathShape) CommitVoid(stage *Stage) {
	gridpathshape.Commit(stage)
}

func (gridpathshape *GridPathShape) StageVoid(stage *Stage) {
	gridpathshape.Stage(stage)
}

// Checkout gridpathshape to the back repo (if it is already staged)
func (gridpathshape *GridPathShape) Checkout(stage *Stage) *GridPathShape {
	if _, ok := stage.GridPathShapes[gridpathshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGridPathShape(gridpathshape)
		}
	}
	return gridpathshape
}

// for satisfaction of GongStruct interface
func (gridpathshape *GridPathShape) GetName() (res string) {
	return gridpathshape.Name
}

// for satisfaction of GongStruct interface
func (gridpathshape *GridPathShape) SetName(name string) {
	gridpathshape.Name = name
}

// Stage puts growthcurverhombusgridshape to the model stage
func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) Stage(stage *Stage) *GrowthCurveRhombusGridShape {
	if _, ok := stage.GrowthCurveRhombusGridShapes[growthcurverhombusgridshape]; !ok {
		stage.GrowthCurveRhombusGridShapes[growthcurverhombusgridshape] = struct{}{}
		stage.GrowthCurveRhombusGridShape_stagedOrder[growthcurverhombusgridshape] = stage.GrowthCurveRhombusGridShapeOrder
		stage.GrowthCurveRhombusGridShape_orderStaged[stage.GrowthCurveRhombusGridShapeOrder] = growthcurverhombusgridshape
		stage.GrowthCurveRhombusGridShapeOrder++
	}
	stage.GrowthCurveRhombusGridShapes_mapString[growthcurverhombusgridshape.Name] = growthcurverhombusgridshape

	return growthcurverhombusgridshape
}

// StagePreserveOrder puts growthcurverhombusgridshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GrowthCurveRhombusGridShapeOrder
// - update stage.GrowthCurveRhombusGridShapeOrder accordingly
func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.GrowthCurveRhombusGridShapes[growthcurverhombusgridshape]; !ok {
		stage.GrowthCurveRhombusGridShapes[growthcurverhombusgridshape] = struct{}{}

		if order > stage.GrowthCurveRhombusGridShapeOrder {
			stage.GrowthCurveRhombusGridShapeOrder = order
		}
		stage.GrowthCurveRhombusGridShape_stagedOrder[growthcurverhombusgridshape] = order
		stage.GrowthCurveRhombusGridShape_orderStaged[order] = growthcurverhombusgridshape
		stage.GrowthCurveRhombusGridShapeOrder++
	}
	stage.GrowthCurveRhombusGridShapes_mapString[growthcurverhombusgridshape.Name] = growthcurverhombusgridshape
}

// Unstage removes growthcurverhombusgridshape off the model stage
func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) Unstage(stage *Stage) *GrowthCurveRhombusGridShape {
	delete(stage.GrowthCurveRhombusGridShapes, growthcurverhombusgridshape)
	// issue1150
	// delete(stage.GrowthCurveRhombusGridShape_stagedOrder, growthcurverhombusgridshape)
	delete(stage.GrowthCurveRhombusGridShapes_mapString, growthcurverhombusgridshape.Name)

	return growthcurverhombusgridshape
}

// UnstageVoid removes growthcurverhombusgridshape off the model stage
func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) UnstageVoid(stage *Stage) {
	delete(stage.GrowthCurveRhombusGridShapes, growthcurverhombusgridshape)
	// issue1150
	// delete(stage.GrowthCurveRhombusGridShape_stagedOrder, growthcurverhombusgridshape)
	delete(stage.GrowthCurveRhombusGridShapes_mapString, growthcurverhombusgridshape.Name)
}

// commit growthcurverhombusgridshape to the back repo (if it is already staged)
func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) Commit(stage *Stage) *GrowthCurveRhombusGridShape {
	if _, ok := stage.GrowthCurveRhombusGridShapes[growthcurverhombusgridshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGrowthCurveRhombusGridShape(growthcurverhombusgridshape)
		}
	}
	return growthcurverhombusgridshape
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) CommitVoid(stage *Stage) {
	growthcurverhombusgridshape.Commit(stage)
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) StageVoid(stage *Stage) {
	growthcurverhombusgridshape.Stage(stage)
}

// Checkout growthcurverhombusgridshape to the back repo (if it is already staged)
func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) Checkout(stage *Stage) *GrowthCurveRhombusGridShape {
	if _, ok := stage.GrowthCurveRhombusGridShapes[growthcurverhombusgridshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGrowthCurveRhombusGridShape(growthcurverhombusgridshape)
		}
	}
	return growthcurverhombusgridshape
}

// for satisfaction of GongStruct interface
func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GetName() (res string) {
	return growthcurverhombusgridshape.Name
}

// for satisfaction of GongStruct interface
func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) SetName(name string) {
	growthcurverhombusgridshape.Name = name
}

// Stage puts growthcurverhombusshape to the model stage
func (growthcurverhombusshape *GrowthCurveRhombusShape) Stage(stage *Stage) *GrowthCurveRhombusShape {
	if _, ok := stage.GrowthCurveRhombusShapes[growthcurverhombusshape]; !ok {
		stage.GrowthCurveRhombusShapes[growthcurverhombusshape] = struct{}{}
		stage.GrowthCurveRhombusShape_stagedOrder[growthcurverhombusshape] = stage.GrowthCurveRhombusShapeOrder
		stage.GrowthCurveRhombusShape_orderStaged[stage.GrowthCurveRhombusShapeOrder] = growthcurverhombusshape
		stage.GrowthCurveRhombusShapeOrder++
	}
	stage.GrowthCurveRhombusShapes_mapString[growthcurverhombusshape.Name] = growthcurverhombusshape

	return growthcurverhombusshape
}

// StagePreserveOrder puts growthcurverhombusshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GrowthCurveRhombusShapeOrder
// - update stage.GrowthCurveRhombusShapeOrder accordingly
func (growthcurverhombusshape *GrowthCurveRhombusShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.GrowthCurveRhombusShapes[growthcurverhombusshape]; !ok {
		stage.GrowthCurveRhombusShapes[growthcurverhombusshape] = struct{}{}

		if order > stage.GrowthCurveRhombusShapeOrder {
			stage.GrowthCurveRhombusShapeOrder = order
		}
		stage.GrowthCurveRhombusShape_stagedOrder[growthcurverhombusshape] = order
		stage.GrowthCurveRhombusShape_orderStaged[order] = growthcurverhombusshape
		stage.GrowthCurveRhombusShapeOrder++
	}
	stage.GrowthCurveRhombusShapes_mapString[growthcurverhombusshape.Name] = growthcurverhombusshape
}

// Unstage removes growthcurverhombusshape off the model stage
func (growthcurverhombusshape *GrowthCurveRhombusShape) Unstage(stage *Stage) *GrowthCurveRhombusShape {
	delete(stage.GrowthCurveRhombusShapes, growthcurverhombusshape)
	// issue1150
	// delete(stage.GrowthCurveRhombusShape_stagedOrder, growthcurverhombusshape)
	delete(stage.GrowthCurveRhombusShapes_mapString, growthcurverhombusshape.Name)

	return growthcurverhombusshape
}

// UnstageVoid removes growthcurverhombusshape off the model stage
func (growthcurverhombusshape *GrowthCurveRhombusShape) UnstageVoid(stage *Stage) {
	delete(stage.GrowthCurveRhombusShapes, growthcurverhombusshape)
	// issue1150
	// delete(stage.GrowthCurveRhombusShape_stagedOrder, growthcurverhombusshape)
	delete(stage.GrowthCurveRhombusShapes_mapString, growthcurverhombusshape.Name)
}

// commit growthcurverhombusshape to the back repo (if it is already staged)
func (growthcurverhombusshape *GrowthCurveRhombusShape) Commit(stage *Stage) *GrowthCurveRhombusShape {
	if _, ok := stage.GrowthCurveRhombusShapes[growthcurverhombusshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGrowthCurveRhombusShape(growthcurverhombusshape)
		}
	}
	return growthcurverhombusshape
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) CommitVoid(stage *Stage) {
	growthcurverhombusshape.Commit(stage)
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) StageVoid(stage *Stage) {
	growthcurverhombusshape.Stage(stage)
}

// Checkout growthcurverhombusshape to the back repo (if it is already staged)
func (growthcurverhombusshape *GrowthCurveRhombusShape) Checkout(stage *Stage) *GrowthCurveRhombusShape {
	if _, ok := stage.GrowthCurveRhombusShapes[growthcurverhombusshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGrowthCurveRhombusShape(growthcurverhombusshape)
		}
	}
	return growthcurverhombusshape
}

// for satisfaction of GongStruct interface
func (growthcurverhombusshape *GrowthCurveRhombusShape) GetName() (res string) {
	return growthcurverhombusshape.Name
}

// for satisfaction of GongStruct interface
func (growthcurverhombusshape *GrowthCurveRhombusShape) SetName(name string) {
	growthcurverhombusshape.Name = name
}

// Stage puts initialrhombusgridshape to the model stage
func (initialrhombusgridshape *InitialRhombusGridShape) Stage(stage *Stage) *InitialRhombusGridShape {
	if _, ok := stage.InitialRhombusGridShapes[initialrhombusgridshape]; !ok {
		stage.InitialRhombusGridShapes[initialrhombusgridshape] = struct{}{}
		stage.InitialRhombusGridShape_stagedOrder[initialrhombusgridshape] = stage.InitialRhombusGridShapeOrder
		stage.InitialRhombusGridShape_orderStaged[stage.InitialRhombusGridShapeOrder] = initialrhombusgridshape
		stage.InitialRhombusGridShapeOrder++
	}
	stage.InitialRhombusGridShapes_mapString[initialrhombusgridshape.Name] = initialrhombusgridshape

	return initialrhombusgridshape
}

// StagePreserveOrder puts initialrhombusgridshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.InitialRhombusGridShapeOrder
// - update stage.InitialRhombusGridShapeOrder accordingly
func (initialrhombusgridshape *InitialRhombusGridShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.InitialRhombusGridShapes[initialrhombusgridshape]; !ok {
		stage.InitialRhombusGridShapes[initialrhombusgridshape] = struct{}{}

		if order > stage.InitialRhombusGridShapeOrder {
			stage.InitialRhombusGridShapeOrder = order
		}
		stage.InitialRhombusGridShape_stagedOrder[initialrhombusgridshape] = order
		stage.InitialRhombusGridShape_orderStaged[order] = initialrhombusgridshape
		stage.InitialRhombusGridShapeOrder++
	}
	stage.InitialRhombusGridShapes_mapString[initialrhombusgridshape.Name] = initialrhombusgridshape
}

// Unstage removes initialrhombusgridshape off the model stage
func (initialrhombusgridshape *InitialRhombusGridShape) Unstage(stage *Stage) *InitialRhombusGridShape {
	delete(stage.InitialRhombusGridShapes, initialrhombusgridshape)
	// issue1150
	// delete(stage.InitialRhombusGridShape_stagedOrder, initialrhombusgridshape)
	delete(stage.InitialRhombusGridShapes_mapString, initialrhombusgridshape.Name)

	return initialrhombusgridshape
}

// UnstageVoid removes initialrhombusgridshape off the model stage
func (initialrhombusgridshape *InitialRhombusGridShape) UnstageVoid(stage *Stage) {
	delete(stage.InitialRhombusGridShapes, initialrhombusgridshape)
	// issue1150
	// delete(stage.InitialRhombusGridShape_stagedOrder, initialrhombusgridshape)
	delete(stage.InitialRhombusGridShapes_mapString, initialrhombusgridshape.Name)
}

// commit initialrhombusgridshape to the back repo (if it is already staged)
func (initialrhombusgridshape *InitialRhombusGridShape) Commit(stage *Stage) *InitialRhombusGridShape {
	if _, ok := stage.InitialRhombusGridShapes[initialrhombusgridshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitInitialRhombusGridShape(initialrhombusgridshape)
		}
	}
	return initialrhombusgridshape
}

func (initialrhombusgridshape *InitialRhombusGridShape) CommitVoid(stage *Stage) {
	initialrhombusgridshape.Commit(stage)
}

func (initialrhombusgridshape *InitialRhombusGridShape) StageVoid(stage *Stage) {
	initialrhombusgridshape.Stage(stage)
}

// Checkout initialrhombusgridshape to the back repo (if it is already staged)
func (initialrhombusgridshape *InitialRhombusGridShape) Checkout(stage *Stage) *InitialRhombusGridShape {
	if _, ok := stage.InitialRhombusGridShapes[initialrhombusgridshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutInitialRhombusGridShape(initialrhombusgridshape)
		}
	}
	return initialrhombusgridshape
}

// for satisfaction of GongStruct interface
func (initialrhombusgridshape *InitialRhombusGridShape) GetName() (res string) {
	return initialrhombusgridshape.Name
}

// for satisfaction of GongStruct interface
func (initialrhombusgridshape *InitialRhombusGridShape) SetName(name string) {
	initialrhombusgridshape.Name = name
}

// Stage puts initialrhombusshape to the model stage
func (initialrhombusshape *InitialRhombusShape) Stage(stage *Stage) *InitialRhombusShape {
	if _, ok := stage.InitialRhombusShapes[initialrhombusshape]; !ok {
		stage.InitialRhombusShapes[initialrhombusshape] = struct{}{}
		stage.InitialRhombusShape_stagedOrder[initialrhombusshape] = stage.InitialRhombusShapeOrder
		stage.InitialRhombusShape_orderStaged[stage.InitialRhombusShapeOrder] = initialrhombusshape
		stage.InitialRhombusShapeOrder++
	}
	stage.InitialRhombusShapes_mapString[initialrhombusshape.Name] = initialrhombusshape

	return initialrhombusshape
}

// StagePreserveOrder puts initialrhombusshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.InitialRhombusShapeOrder
// - update stage.InitialRhombusShapeOrder accordingly
func (initialrhombusshape *InitialRhombusShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.InitialRhombusShapes[initialrhombusshape]; !ok {
		stage.InitialRhombusShapes[initialrhombusshape] = struct{}{}

		if order > stage.InitialRhombusShapeOrder {
			stage.InitialRhombusShapeOrder = order
		}
		stage.InitialRhombusShape_stagedOrder[initialrhombusshape] = order
		stage.InitialRhombusShape_orderStaged[order] = initialrhombusshape
		stage.InitialRhombusShapeOrder++
	}
	stage.InitialRhombusShapes_mapString[initialrhombusshape.Name] = initialrhombusshape
}

// Unstage removes initialrhombusshape off the model stage
func (initialrhombusshape *InitialRhombusShape) Unstage(stage *Stage) *InitialRhombusShape {
	delete(stage.InitialRhombusShapes, initialrhombusshape)
	// issue1150
	// delete(stage.InitialRhombusShape_stagedOrder, initialrhombusshape)
	delete(stage.InitialRhombusShapes_mapString, initialrhombusshape.Name)

	return initialrhombusshape
}

// UnstageVoid removes initialrhombusshape off the model stage
func (initialrhombusshape *InitialRhombusShape) UnstageVoid(stage *Stage) {
	delete(stage.InitialRhombusShapes, initialrhombusshape)
	// issue1150
	// delete(stage.InitialRhombusShape_stagedOrder, initialrhombusshape)
	delete(stage.InitialRhombusShapes_mapString, initialrhombusshape.Name)
}

// commit initialrhombusshape to the back repo (if it is already staged)
func (initialrhombusshape *InitialRhombusShape) Commit(stage *Stage) *InitialRhombusShape {
	if _, ok := stage.InitialRhombusShapes[initialrhombusshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitInitialRhombusShape(initialrhombusshape)
		}
	}
	return initialrhombusshape
}

func (initialrhombusshape *InitialRhombusShape) CommitVoid(stage *Stage) {
	initialrhombusshape.Commit(stage)
}

func (initialrhombusshape *InitialRhombusShape) StageVoid(stage *Stage) {
	initialrhombusshape.Stage(stage)
}

// Checkout initialrhombusshape to the back repo (if it is already staged)
func (initialrhombusshape *InitialRhombusShape) Checkout(stage *Stage) *InitialRhombusShape {
	if _, ok := stage.InitialRhombusShapes[initialrhombusshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutInitialRhombusShape(initialrhombusshape)
		}
	}
	return initialrhombusshape
}

// for satisfaction of GongStruct interface
func (initialrhombusshape *InitialRhombusShape) GetName() (res string) {
	return initialrhombusshape.Name
}

// for satisfaction of GongStruct interface
func (initialrhombusshape *InitialRhombusShape) SetName(name string) {
	initialrhombusshape.Name = name
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

// Stage puts nextcircleshape to the model stage
func (nextcircleshape *NextCircleShape) Stage(stage *Stage) *NextCircleShape {
	if _, ok := stage.NextCircleShapes[nextcircleshape]; !ok {
		stage.NextCircleShapes[nextcircleshape] = struct{}{}
		stage.NextCircleShape_stagedOrder[nextcircleshape] = stage.NextCircleShapeOrder
		stage.NextCircleShape_orderStaged[stage.NextCircleShapeOrder] = nextcircleshape
		stage.NextCircleShapeOrder++
	}
	stage.NextCircleShapes_mapString[nextcircleshape.Name] = nextcircleshape

	return nextcircleshape
}

// StagePreserveOrder puts nextcircleshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NextCircleShapeOrder
// - update stage.NextCircleShapeOrder accordingly
func (nextcircleshape *NextCircleShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.NextCircleShapes[nextcircleshape]; !ok {
		stage.NextCircleShapes[nextcircleshape] = struct{}{}

		if order > stage.NextCircleShapeOrder {
			stage.NextCircleShapeOrder = order
		}
		stage.NextCircleShape_stagedOrder[nextcircleshape] = order
		stage.NextCircleShape_orderStaged[order] = nextcircleshape
		stage.NextCircleShapeOrder++
	}
	stage.NextCircleShapes_mapString[nextcircleshape.Name] = nextcircleshape
}

// Unstage removes nextcircleshape off the model stage
func (nextcircleshape *NextCircleShape) Unstage(stage *Stage) *NextCircleShape {
	delete(stage.NextCircleShapes, nextcircleshape)
	// issue1150
	// delete(stage.NextCircleShape_stagedOrder, nextcircleshape)
	delete(stage.NextCircleShapes_mapString, nextcircleshape.Name)

	return nextcircleshape
}

// UnstageVoid removes nextcircleshape off the model stage
func (nextcircleshape *NextCircleShape) UnstageVoid(stage *Stage) {
	delete(stage.NextCircleShapes, nextcircleshape)
	// issue1150
	// delete(stage.NextCircleShape_stagedOrder, nextcircleshape)
	delete(stage.NextCircleShapes_mapString, nextcircleshape.Name)
}

// commit nextcircleshape to the back repo (if it is already staged)
func (nextcircleshape *NextCircleShape) Commit(stage *Stage) *NextCircleShape {
	if _, ok := stage.NextCircleShapes[nextcircleshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNextCircleShape(nextcircleshape)
		}
	}
	return nextcircleshape
}

func (nextcircleshape *NextCircleShape) CommitVoid(stage *Stage) {
	nextcircleshape.Commit(stage)
}

func (nextcircleshape *NextCircleShape) StageVoid(stage *Stage) {
	nextcircleshape.Stage(stage)
}

// Checkout nextcircleshape to the back repo (if it is already staged)
func (nextcircleshape *NextCircleShape) Checkout(stage *Stage) *NextCircleShape {
	if _, ok := stage.NextCircleShapes[nextcircleshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNextCircleShape(nextcircleshape)
		}
	}
	return nextcircleshape
}

// for satisfaction of GongStruct interface
func (nextcircleshape *NextCircleShape) GetName() (res string) {
	return nextcircleshape.Name
}

// for satisfaction of GongStruct interface
func (nextcircleshape *NextCircleShape) SetName(name string) {
	nextcircleshape.Name = name
}

// Stage puts plant to the model stage
func (plant *Plant) Stage(stage *Stage) *Plant {
	if _, ok := stage.Plants[plant]; !ok {
		stage.Plants[plant] = struct{}{}
		stage.Plant_stagedOrder[plant] = stage.PlantOrder
		stage.Plant_orderStaged[stage.PlantOrder] = plant
		stage.PlantOrder++
	}
	stage.Plants_mapString[plant.Name] = plant

	return plant
}

// StagePreserveOrder puts plant to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PlantOrder
// - update stage.PlantOrder accordingly
func (plant *Plant) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Plants[plant]; !ok {
		stage.Plants[plant] = struct{}{}

		if order > stage.PlantOrder {
			stage.PlantOrder = order
		}
		stage.Plant_stagedOrder[plant] = order
		stage.Plant_orderStaged[order] = plant
		stage.PlantOrder++
	}
	stage.Plants_mapString[plant.Name] = plant
}

// Unstage removes plant off the model stage
func (plant *Plant) Unstage(stage *Stage) *Plant {
	delete(stage.Plants, plant)
	// issue1150
	// delete(stage.Plant_stagedOrder, plant)
	delete(stage.Plants_mapString, plant.Name)

	return plant
}

// UnstageVoid removes plant off the model stage
func (plant *Plant) UnstageVoid(stage *Stage) {
	delete(stage.Plants, plant)
	// issue1150
	// delete(stage.Plant_stagedOrder, plant)
	delete(stage.Plants_mapString, plant.Name)
}

// commit plant to the back repo (if it is already staged)
func (plant *Plant) Commit(stage *Stage) *Plant {
	if _, ok := stage.Plants[plant]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPlant(plant)
		}
	}
	return plant
}

func (plant *Plant) CommitVoid(stage *Stage) {
	plant.Commit(stage)
}

func (plant *Plant) StageVoid(stage *Stage) {
	plant.Stage(stage)
}

// Checkout plant to the back repo (if it is already staged)
func (plant *Plant) Checkout(stage *Stage) *Plant {
	if _, ok := stage.Plants[plant]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPlant(plant)
		}
	}
	return plant
}

// for satisfaction of GongStruct interface
func (plant *Plant) GetName() (res string) {
	return plant.Name
}

// for satisfaction of GongStruct interface
func (plant *Plant) SetName(name string) {
	plant.Name = name
}

// Stage puts plantcircumferenceshape to the model stage
func (plantcircumferenceshape *PlantCircumferenceShape) Stage(stage *Stage) *PlantCircumferenceShape {
	if _, ok := stage.PlantCircumferenceShapes[plantcircumferenceshape]; !ok {
		stage.PlantCircumferenceShapes[plantcircumferenceshape] = struct{}{}
		stage.PlantCircumferenceShape_stagedOrder[plantcircumferenceshape] = stage.PlantCircumferenceShapeOrder
		stage.PlantCircumferenceShape_orderStaged[stage.PlantCircumferenceShapeOrder] = plantcircumferenceshape
		stage.PlantCircumferenceShapeOrder++
	}
	stage.PlantCircumferenceShapes_mapString[plantcircumferenceshape.Name] = plantcircumferenceshape

	return plantcircumferenceshape
}

// StagePreserveOrder puts plantcircumferenceshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PlantCircumferenceShapeOrder
// - update stage.PlantCircumferenceShapeOrder accordingly
func (plantcircumferenceshape *PlantCircumferenceShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.PlantCircumferenceShapes[plantcircumferenceshape]; !ok {
		stage.PlantCircumferenceShapes[plantcircumferenceshape] = struct{}{}

		if order > stage.PlantCircumferenceShapeOrder {
			stage.PlantCircumferenceShapeOrder = order
		}
		stage.PlantCircumferenceShape_stagedOrder[plantcircumferenceshape] = order
		stage.PlantCircumferenceShape_orderStaged[order] = plantcircumferenceshape
		stage.PlantCircumferenceShapeOrder++
	}
	stage.PlantCircumferenceShapes_mapString[plantcircumferenceshape.Name] = plantcircumferenceshape
}

// Unstage removes plantcircumferenceshape off the model stage
func (plantcircumferenceshape *PlantCircumferenceShape) Unstage(stage *Stage) *PlantCircumferenceShape {
	delete(stage.PlantCircumferenceShapes, plantcircumferenceshape)
	// issue1150
	// delete(stage.PlantCircumferenceShape_stagedOrder, plantcircumferenceshape)
	delete(stage.PlantCircumferenceShapes_mapString, plantcircumferenceshape.Name)

	return plantcircumferenceshape
}

// UnstageVoid removes plantcircumferenceshape off the model stage
func (plantcircumferenceshape *PlantCircumferenceShape) UnstageVoid(stage *Stage) {
	delete(stage.PlantCircumferenceShapes, plantcircumferenceshape)
	// issue1150
	// delete(stage.PlantCircumferenceShape_stagedOrder, plantcircumferenceshape)
	delete(stage.PlantCircumferenceShapes_mapString, plantcircumferenceshape.Name)
}

// commit plantcircumferenceshape to the back repo (if it is already staged)
func (plantcircumferenceshape *PlantCircumferenceShape) Commit(stage *Stage) *PlantCircumferenceShape {
	if _, ok := stage.PlantCircumferenceShapes[plantcircumferenceshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPlantCircumferenceShape(plantcircumferenceshape)
		}
	}
	return plantcircumferenceshape
}

func (plantcircumferenceshape *PlantCircumferenceShape) CommitVoid(stage *Stage) {
	plantcircumferenceshape.Commit(stage)
}

func (plantcircumferenceshape *PlantCircumferenceShape) StageVoid(stage *Stage) {
	plantcircumferenceshape.Stage(stage)
}

// Checkout plantcircumferenceshape to the back repo (if it is already staged)
func (plantcircumferenceshape *PlantCircumferenceShape) Checkout(stage *Stage) *PlantCircumferenceShape {
	if _, ok := stage.PlantCircumferenceShapes[plantcircumferenceshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPlantCircumferenceShape(plantcircumferenceshape)
		}
	}
	return plantcircumferenceshape
}

// for satisfaction of GongStruct interface
func (plantcircumferenceshape *PlantCircumferenceShape) GetName() (res string) {
	return plantcircumferenceshape.Name
}

// for satisfaction of GongStruct interface
func (plantcircumferenceshape *PlantCircumferenceShape) SetName(name string) {
	plantcircumferenceshape.Name = name
}

// Stage puts plantdiagram to the model stage
func (plantdiagram *PlantDiagram) Stage(stage *Stage) *PlantDiagram {
	if _, ok := stage.PlantDiagrams[plantdiagram]; !ok {
		stage.PlantDiagrams[plantdiagram] = struct{}{}
		stage.PlantDiagram_stagedOrder[plantdiagram] = stage.PlantDiagramOrder
		stage.PlantDiagram_orderStaged[stage.PlantDiagramOrder] = plantdiagram
		stage.PlantDiagramOrder++
	}
	stage.PlantDiagrams_mapString[plantdiagram.Name] = plantdiagram

	return plantdiagram
}

// StagePreserveOrder puts plantdiagram to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PlantDiagramOrder
// - update stage.PlantDiagramOrder accordingly
func (plantdiagram *PlantDiagram) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.PlantDiagrams[plantdiagram]; !ok {
		stage.PlantDiagrams[plantdiagram] = struct{}{}

		if order > stage.PlantDiagramOrder {
			stage.PlantDiagramOrder = order
		}
		stage.PlantDiagram_stagedOrder[plantdiagram] = order
		stage.PlantDiagram_orderStaged[order] = plantdiagram
		stage.PlantDiagramOrder++
	}
	stage.PlantDiagrams_mapString[plantdiagram.Name] = plantdiagram
}

// Unstage removes plantdiagram off the model stage
func (plantdiagram *PlantDiagram) Unstage(stage *Stage) *PlantDiagram {
	delete(stage.PlantDiagrams, plantdiagram)
	// issue1150
	// delete(stage.PlantDiagram_stagedOrder, plantdiagram)
	delete(stage.PlantDiagrams_mapString, plantdiagram.Name)

	return plantdiagram
}

// UnstageVoid removes plantdiagram off the model stage
func (plantdiagram *PlantDiagram) UnstageVoid(stage *Stage) {
	delete(stage.PlantDiagrams, plantdiagram)
	// issue1150
	// delete(stage.PlantDiagram_stagedOrder, plantdiagram)
	delete(stage.PlantDiagrams_mapString, plantdiagram.Name)
}

// commit plantdiagram to the back repo (if it is already staged)
func (plantdiagram *PlantDiagram) Commit(stage *Stage) *PlantDiagram {
	if _, ok := stage.PlantDiagrams[plantdiagram]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPlantDiagram(plantdiagram)
		}
	}
	return plantdiagram
}

func (plantdiagram *PlantDiagram) CommitVoid(stage *Stage) {
	plantdiagram.Commit(stage)
}

func (plantdiagram *PlantDiagram) StageVoid(stage *Stage) {
	plantdiagram.Stage(stage)
}

// Checkout plantdiagram to the back repo (if it is already staged)
func (plantdiagram *PlantDiagram) Checkout(stage *Stage) *PlantDiagram {
	if _, ok := stage.PlantDiagrams[plantdiagram]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPlantDiagram(plantdiagram)
		}
	}
	return plantdiagram
}

// for satisfaction of GongStruct interface
func (plantdiagram *PlantDiagram) GetName() (res string) {
	return plantdiagram.Name
}

// for satisfaction of GongStruct interface
func (plantdiagram *PlantDiagram) SetName(name string) {
	plantdiagram.Name = name
}

// Stage puts rhombusgridshape to the model stage
func (rhombusgridshape *RhombusGridShape) Stage(stage *Stage) *RhombusGridShape {
	if _, ok := stage.RhombusGridShapes[rhombusgridshape]; !ok {
		stage.RhombusGridShapes[rhombusgridshape] = struct{}{}
		stage.RhombusGridShape_stagedOrder[rhombusgridshape] = stage.RhombusGridShapeOrder
		stage.RhombusGridShape_orderStaged[stage.RhombusGridShapeOrder] = rhombusgridshape
		stage.RhombusGridShapeOrder++
	}
	stage.RhombusGridShapes_mapString[rhombusgridshape.Name] = rhombusgridshape

	return rhombusgridshape
}

// StagePreserveOrder puts rhombusgridshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RhombusGridShapeOrder
// - update stage.RhombusGridShapeOrder accordingly
func (rhombusgridshape *RhombusGridShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.RhombusGridShapes[rhombusgridshape]; !ok {
		stage.RhombusGridShapes[rhombusgridshape] = struct{}{}

		if order > stage.RhombusGridShapeOrder {
			stage.RhombusGridShapeOrder = order
		}
		stage.RhombusGridShape_stagedOrder[rhombusgridshape] = order
		stage.RhombusGridShape_orderStaged[order] = rhombusgridshape
		stage.RhombusGridShapeOrder++
	}
	stage.RhombusGridShapes_mapString[rhombusgridshape.Name] = rhombusgridshape
}

// Unstage removes rhombusgridshape off the model stage
func (rhombusgridshape *RhombusGridShape) Unstage(stage *Stage) *RhombusGridShape {
	delete(stage.RhombusGridShapes, rhombusgridshape)
	// issue1150
	// delete(stage.RhombusGridShape_stagedOrder, rhombusgridshape)
	delete(stage.RhombusGridShapes_mapString, rhombusgridshape.Name)

	return rhombusgridshape
}

// UnstageVoid removes rhombusgridshape off the model stage
func (rhombusgridshape *RhombusGridShape) UnstageVoid(stage *Stage) {
	delete(stage.RhombusGridShapes, rhombusgridshape)
	// issue1150
	// delete(stage.RhombusGridShape_stagedOrder, rhombusgridshape)
	delete(stage.RhombusGridShapes_mapString, rhombusgridshape.Name)
}

// commit rhombusgridshape to the back repo (if it is already staged)
func (rhombusgridshape *RhombusGridShape) Commit(stage *Stage) *RhombusGridShape {
	if _, ok := stage.RhombusGridShapes[rhombusgridshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRhombusGridShape(rhombusgridshape)
		}
	}
	return rhombusgridshape
}

func (rhombusgridshape *RhombusGridShape) CommitVoid(stage *Stage) {
	rhombusgridshape.Commit(stage)
}

func (rhombusgridshape *RhombusGridShape) StageVoid(stage *Stage) {
	rhombusgridshape.Stage(stage)
}

// Checkout rhombusgridshape to the back repo (if it is already staged)
func (rhombusgridshape *RhombusGridShape) Checkout(stage *Stage) *RhombusGridShape {
	if _, ok := stage.RhombusGridShapes[rhombusgridshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRhombusGridShape(rhombusgridshape)
		}
	}
	return rhombusgridshape
}

// for satisfaction of GongStruct interface
func (rhombusgridshape *RhombusGridShape) GetName() (res string) {
	return rhombusgridshape.Name
}

// for satisfaction of GongStruct interface
func (rhombusgridshape *RhombusGridShape) SetName(name string) {
	rhombusgridshape.Name = name
}

// Stage puts rhombusshape to the model stage
func (rhombusshape *RhombusShape) Stage(stage *Stage) *RhombusShape {
	if _, ok := stage.RhombusShapes[rhombusshape]; !ok {
		stage.RhombusShapes[rhombusshape] = struct{}{}
		stage.RhombusShape_stagedOrder[rhombusshape] = stage.RhombusShapeOrder
		stage.RhombusShape_orderStaged[stage.RhombusShapeOrder] = rhombusshape
		stage.RhombusShapeOrder++
	}
	stage.RhombusShapes_mapString[rhombusshape.Name] = rhombusshape

	return rhombusshape
}

// StagePreserveOrder puts rhombusshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RhombusShapeOrder
// - update stage.RhombusShapeOrder accordingly
func (rhombusshape *RhombusShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.RhombusShapes[rhombusshape]; !ok {
		stage.RhombusShapes[rhombusshape] = struct{}{}

		if order > stage.RhombusShapeOrder {
			stage.RhombusShapeOrder = order
		}
		stage.RhombusShape_stagedOrder[rhombusshape] = order
		stage.RhombusShape_orderStaged[order] = rhombusshape
		stage.RhombusShapeOrder++
	}
	stage.RhombusShapes_mapString[rhombusshape.Name] = rhombusshape
}

// Unstage removes rhombusshape off the model stage
func (rhombusshape *RhombusShape) Unstage(stage *Stage) *RhombusShape {
	delete(stage.RhombusShapes, rhombusshape)
	// issue1150
	// delete(stage.RhombusShape_stagedOrder, rhombusshape)
	delete(stage.RhombusShapes_mapString, rhombusshape.Name)

	return rhombusshape
}

// UnstageVoid removes rhombusshape off the model stage
func (rhombusshape *RhombusShape) UnstageVoid(stage *Stage) {
	delete(stage.RhombusShapes, rhombusshape)
	// issue1150
	// delete(stage.RhombusShape_stagedOrder, rhombusshape)
	delete(stage.RhombusShapes_mapString, rhombusshape.Name)
}

// commit rhombusshape to the back repo (if it is already staged)
func (rhombusshape *RhombusShape) Commit(stage *Stage) *RhombusShape {
	if _, ok := stage.RhombusShapes[rhombusshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRhombusShape(rhombusshape)
		}
	}
	return rhombusshape
}

func (rhombusshape *RhombusShape) CommitVoid(stage *Stage) {
	rhombusshape.Commit(stage)
}

func (rhombusshape *RhombusShape) StageVoid(stage *Stage) {
	rhombusshape.Stage(stage)
}

// Checkout rhombusshape to the back repo (if it is already staged)
func (rhombusshape *RhombusShape) Checkout(stage *Stage) *RhombusShape {
	if _, ok := stage.RhombusShapes[rhombusshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRhombusShape(rhombusshape)
		}
	}
	return rhombusshape
}

// for satisfaction of GongStruct interface
func (rhombusshape *RhombusShape) GetName() (res string) {
	return rhombusshape.Name
}

// for satisfaction of GongStruct interface
func (rhombusshape *RhombusShape) SetName(name string) {
	rhombusshape.Name = name
}

// Stage puts rotatedrhombusgridshape to the model stage
func (rotatedrhombusgridshape *RotatedRhombusGridShape) Stage(stage *Stage) *RotatedRhombusGridShape {
	if _, ok := stage.RotatedRhombusGridShapes[rotatedrhombusgridshape]; !ok {
		stage.RotatedRhombusGridShapes[rotatedrhombusgridshape] = struct{}{}
		stage.RotatedRhombusGridShape_stagedOrder[rotatedrhombusgridshape] = stage.RotatedRhombusGridShapeOrder
		stage.RotatedRhombusGridShape_orderStaged[stage.RotatedRhombusGridShapeOrder] = rotatedrhombusgridshape
		stage.RotatedRhombusGridShapeOrder++
	}
	stage.RotatedRhombusGridShapes_mapString[rotatedrhombusgridshape.Name] = rotatedrhombusgridshape

	return rotatedrhombusgridshape
}

// StagePreserveOrder puts rotatedrhombusgridshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RotatedRhombusGridShapeOrder
// - update stage.RotatedRhombusGridShapeOrder accordingly
func (rotatedrhombusgridshape *RotatedRhombusGridShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.RotatedRhombusGridShapes[rotatedrhombusgridshape]; !ok {
		stage.RotatedRhombusGridShapes[rotatedrhombusgridshape] = struct{}{}

		if order > stage.RotatedRhombusGridShapeOrder {
			stage.RotatedRhombusGridShapeOrder = order
		}
		stage.RotatedRhombusGridShape_stagedOrder[rotatedrhombusgridshape] = order
		stage.RotatedRhombusGridShape_orderStaged[order] = rotatedrhombusgridshape
		stage.RotatedRhombusGridShapeOrder++
	}
	stage.RotatedRhombusGridShapes_mapString[rotatedrhombusgridshape.Name] = rotatedrhombusgridshape
}

// Unstage removes rotatedrhombusgridshape off the model stage
func (rotatedrhombusgridshape *RotatedRhombusGridShape) Unstage(stage *Stage) *RotatedRhombusGridShape {
	delete(stage.RotatedRhombusGridShapes, rotatedrhombusgridshape)
	// issue1150
	// delete(stage.RotatedRhombusGridShape_stagedOrder, rotatedrhombusgridshape)
	delete(stage.RotatedRhombusGridShapes_mapString, rotatedrhombusgridshape.Name)

	return rotatedrhombusgridshape
}

// UnstageVoid removes rotatedrhombusgridshape off the model stage
func (rotatedrhombusgridshape *RotatedRhombusGridShape) UnstageVoid(stage *Stage) {
	delete(stage.RotatedRhombusGridShapes, rotatedrhombusgridshape)
	// issue1150
	// delete(stage.RotatedRhombusGridShape_stagedOrder, rotatedrhombusgridshape)
	delete(stage.RotatedRhombusGridShapes_mapString, rotatedrhombusgridshape.Name)
}

// commit rotatedrhombusgridshape to the back repo (if it is already staged)
func (rotatedrhombusgridshape *RotatedRhombusGridShape) Commit(stage *Stage) *RotatedRhombusGridShape {
	if _, ok := stage.RotatedRhombusGridShapes[rotatedrhombusgridshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRotatedRhombusGridShape(rotatedrhombusgridshape)
		}
	}
	return rotatedrhombusgridshape
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) CommitVoid(stage *Stage) {
	rotatedrhombusgridshape.Commit(stage)
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) StageVoid(stage *Stage) {
	rotatedrhombusgridshape.Stage(stage)
}

// Checkout rotatedrhombusgridshape to the back repo (if it is already staged)
func (rotatedrhombusgridshape *RotatedRhombusGridShape) Checkout(stage *Stage) *RotatedRhombusGridShape {
	if _, ok := stage.RotatedRhombusGridShapes[rotatedrhombusgridshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRotatedRhombusGridShape(rotatedrhombusgridshape)
		}
	}
	return rotatedrhombusgridshape
}

// for satisfaction of GongStruct interface
func (rotatedrhombusgridshape *RotatedRhombusGridShape) GetName() (res string) {
	return rotatedrhombusgridshape.Name
}

// for satisfaction of GongStruct interface
func (rotatedrhombusgridshape *RotatedRhombusGridShape) SetName(name string) {
	rotatedrhombusgridshape.Name = name
}

// Stage puts rotatedrhombusshape to the model stage
func (rotatedrhombusshape *RotatedRhombusShape) Stage(stage *Stage) *RotatedRhombusShape {
	if _, ok := stage.RotatedRhombusShapes[rotatedrhombusshape]; !ok {
		stage.RotatedRhombusShapes[rotatedrhombusshape] = struct{}{}
		stage.RotatedRhombusShape_stagedOrder[rotatedrhombusshape] = stage.RotatedRhombusShapeOrder
		stage.RotatedRhombusShape_orderStaged[stage.RotatedRhombusShapeOrder] = rotatedrhombusshape
		stage.RotatedRhombusShapeOrder++
	}
	stage.RotatedRhombusShapes_mapString[rotatedrhombusshape.Name] = rotatedrhombusshape

	return rotatedrhombusshape
}

// StagePreserveOrder puts rotatedrhombusshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RotatedRhombusShapeOrder
// - update stage.RotatedRhombusShapeOrder accordingly
func (rotatedrhombusshape *RotatedRhombusShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.RotatedRhombusShapes[rotatedrhombusshape]; !ok {
		stage.RotatedRhombusShapes[rotatedrhombusshape] = struct{}{}

		if order > stage.RotatedRhombusShapeOrder {
			stage.RotatedRhombusShapeOrder = order
		}
		stage.RotatedRhombusShape_stagedOrder[rotatedrhombusshape] = order
		stage.RotatedRhombusShape_orderStaged[order] = rotatedrhombusshape
		stage.RotatedRhombusShapeOrder++
	}
	stage.RotatedRhombusShapes_mapString[rotatedrhombusshape.Name] = rotatedrhombusshape
}

// Unstage removes rotatedrhombusshape off the model stage
func (rotatedrhombusshape *RotatedRhombusShape) Unstage(stage *Stage) *RotatedRhombusShape {
	delete(stage.RotatedRhombusShapes, rotatedrhombusshape)
	// issue1150
	// delete(stage.RotatedRhombusShape_stagedOrder, rotatedrhombusshape)
	delete(stage.RotatedRhombusShapes_mapString, rotatedrhombusshape.Name)

	return rotatedrhombusshape
}

// UnstageVoid removes rotatedrhombusshape off the model stage
func (rotatedrhombusshape *RotatedRhombusShape) UnstageVoid(stage *Stage) {
	delete(stage.RotatedRhombusShapes, rotatedrhombusshape)
	// issue1150
	// delete(stage.RotatedRhombusShape_stagedOrder, rotatedrhombusshape)
	delete(stage.RotatedRhombusShapes_mapString, rotatedrhombusshape.Name)
}

// commit rotatedrhombusshape to the back repo (if it is already staged)
func (rotatedrhombusshape *RotatedRhombusShape) Commit(stage *Stage) *RotatedRhombusShape {
	if _, ok := stage.RotatedRhombusShapes[rotatedrhombusshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRotatedRhombusShape(rotatedrhombusshape)
		}
	}
	return rotatedrhombusshape
}

func (rotatedrhombusshape *RotatedRhombusShape) CommitVoid(stage *Stage) {
	rotatedrhombusshape.Commit(stage)
}

func (rotatedrhombusshape *RotatedRhombusShape) StageVoid(stage *Stage) {
	rotatedrhombusshape.Stage(stage)
}

// Checkout rotatedrhombusshape to the back repo (if it is already staged)
func (rotatedrhombusshape *RotatedRhombusShape) Checkout(stage *Stage) *RotatedRhombusShape {
	if _, ok := stage.RotatedRhombusShapes[rotatedrhombusshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRotatedRhombusShape(rotatedrhombusshape)
		}
	}
	return rotatedrhombusshape
}

// for satisfaction of GongStruct interface
func (rotatedrhombusshape *RotatedRhombusShape) GetName() (res string) {
	return rotatedrhombusshape.Name
}

// for satisfaction of GongStruct interface
func (rotatedrhombusshape *RotatedRhombusShape) SetName(name string) {
	rotatedrhombusshape.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMAxesShape(AxesShape *AxesShape)
	CreateORMCircleGridShape(CircleGridShape *CircleGridShape)
	CreateORMExplanationTextShape(ExplanationTextShape *ExplanationTextShape)
	CreateORMGridPathShape(GridPathShape *GridPathShape)
	CreateORMGrowthCurveRhombusGridShape(GrowthCurveRhombusGridShape *GrowthCurveRhombusGridShape)
	CreateORMGrowthCurveRhombusShape(GrowthCurveRhombusShape *GrowthCurveRhombusShape)
	CreateORMInitialRhombusGridShape(InitialRhombusGridShape *InitialRhombusGridShape)
	CreateORMInitialRhombusShape(InitialRhombusShape *InitialRhombusShape)
	CreateORMLibrary(Library *Library)
	CreateORMNextCircleShape(NextCircleShape *NextCircleShape)
	CreateORMPlant(Plant *Plant)
	CreateORMPlantCircumferenceShape(PlantCircumferenceShape *PlantCircumferenceShape)
	CreateORMPlantDiagram(PlantDiagram *PlantDiagram)
	CreateORMRhombusGridShape(RhombusGridShape *RhombusGridShape)
	CreateORMRhombusShape(RhombusShape *RhombusShape)
	CreateORMRotatedRhombusGridShape(RotatedRhombusGridShape *RotatedRhombusGridShape)
	CreateORMRotatedRhombusShape(RotatedRhombusShape *RotatedRhombusShape)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAxesShape(AxesShape *AxesShape)
	DeleteORMCircleGridShape(CircleGridShape *CircleGridShape)
	DeleteORMExplanationTextShape(ExplanationTextShape *ExplanationTextShape)
	DeleteORMGridPathShape(GridPathShape *GridPathShape)
	DeleteORMGrowthCurveRhombusGridShape(GrowthCurveRhombusGridShape *GrowthCurveRhombusGridShape)
	DeleteORMGrowthCurveRhombusShape(GrowthCurveRhombusShape *GrowthCurveRhombusShape)
	DeleteORMInitialRhombusGridShape(InitialRhombusGridShape *InitialRhombusGridShape)
	DeleteORMInitialRhombusShape(InitialRhombusShape *InitialRhombusShape)
	DeleteORMLibrary(Library *Library)
	DeleteORMNextCircleShape(NextCircleShape *NextCircleShape)
	DeleteORMPlant(Plant *Plant)
	DeleteORMPlantCircumferenceShape(PlantCircumferenceShape *PlantCircumferenceShape)
	DeleteORMPlantDiagram(PlantDiagram *PlantDiagram)
	DeleteORMRhombusGridShape(RhombusGridShape *RhombusGridShape)
	DeleteORMRhombusShape(RhombusShape *RhombusShape)
	DeleteORMRotatedRhombusGridShape(RotatedRhombusGridShape *RotatedRhombusGridShape)
	DeleteORMRotatedRhombusShape(RotatedRhombusShape *RotatedRhombusShape)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.AxesShapes = make(map[*AxesShape]struct{})
	stage.AxesShapes_mapString = make(map[string]*AxesShape)
	stage.AxesShape_stagedOrder = make(map[*AxesShape]uint)
	stage.AxesShapeOrder = 0

	stage.CircleGridShapes = make(map[*CircleGridShape]struct{})
	stage.CircleGridShapes_mapString = make(map[string]*CircleGridShape)
	stage.CircleGridShape_stagedOrder = make(map[*CircleGridShape]uint)
	stage.CircleGridShapeOrder = 0

	stage.ExplanationTextShapes = make(map[*ExplanationTextShape]struct{})
	stage.ExplanationTextShapes_mapString = make(map[string]*ExplanationTextShape)
	stage.ExplanationTextShape_stagedOrder = make(map[*ExplanationTextShape]uint)
	stage.ExplanationTextShapeOrder = 0

	stage.GridPathShapes = make(map[*GridPathShape]struct{})
	stage.GridPathShapes_mapString = make(map[string]*GridPathShape)
	stage.GridPathShape_stagedOrder = make(map[*GridPathShape]uint)
	stage.GridPathShapeOrder = 0

	stage.GrowthCurveRhombusGridShapes = make(map[*GrowthCurveRhombusGridShape]struct{})
	stage.GrowthCurveRhombusGridShapes_mapString = make(map[string]*GrowthCurveRhombusGridShape)
	stage.GrowthCurveRhombusGridShape_stagedOrder = make(map[*GrowthCurveRhombusGridShape]uint)
	stage.GrowthCurveRhombusGridShapeOrder = 0

	stage.GrowthCurveRhombusShapes = make(map[*GrowthCurveRhombusShape]struct{})
	stage.GrowthCurveRhombusShapes_mapString = make(map[string]*GrowthCurveRhombusShape)
	stage.GrowthCurveRhombusShape_stagedOrder = make(map[*GrowthCurveRhombusShape]uint)
	stage.GrowthCurveRhombusShapeOrder = 0

	stage.InitialRhombusGridShapes = make(map[*InitialRhombusGridShape]struct{})
	stage.InitialRhombusGridShapes_mapString = make(map[string]*InitialRhombusGridShape)
	stage.InitialRhombusGridShape_stagedOrder = make(map[*InitialRhombusGridShape]uint)
	stage.InitialRhombusGridShapeOrder = 0

	stage.InitialRhombusShapes = make(map[*InitialRhombusShape]struct{})
	stage.InitialRhombusShapes_mapString = make(map[string]*InitialRhombusShape)
	stage.InitialRhombusShape_stagedOrder = make(map[*InitialRhombusShape]uint)
	stage.InitialRhombusShapeOrder = 0

	stage.Librarys = make(map[*Library]struct{})
	stage.Librarys_mapString = make(map[string]*Library)
	stage.Library_stagedOrder = make(map[*Library]uint)
	stage.LibraryOrder = 0

	stage.NextCircleShapes = make(map[*NextCircleShape]struct{})
	stage.NextCircleShapes_mapString = make(map[string]*NextCircleShape)
	stage.NextCircleShape_stagedOrder = make(map[*NextCircleShape]uint)
	stage.NextCircleShapeOrder = 0

	stage.Plants = make(map[*Plant]struct{})
	stage.Plants_mapString = make(map[string]*Plant)
	stage.Plant_stagedOrder = make(map[*Plant]uint)
	stage.PlantOrder = 0

	stage.PlantCircumferenceShapes = make(map[*PlantCircumferenceShape]struct{})
	stage.PlantCircumferenceShapes_mapString = make(map[string]*PlantCircumferenceShape)
	stage.PlantCircumferenceShape_stagedOrder = make(map[*PlantCircumferenceShape]uint)
	stage.PlantCircumferenceShapeOrder = 0

	stage.PlantDiagrams = make(map[*PlantDiagram]struct{})
	stage.PlantDiagrams_mapString = make(map[string]*PlantDiagram)
	stage.PlantDiagram_stagedOrder = make(map[*PlantDiagram]uint)
	stage.PlantDiagramOrder = 0

	stage.RhombusGridShapes = make(map[*RhombusGridShape]struct{})
	stage.RhombusGridShapes_mapString = make(map[string]*RhombusGridShape)
	stage.RhombusGridShape_stagedOrder = make(map[*RhombusGridShape]uint)
	stage.RhombusGridShapeOrder = 0

	stage.RhombusShapes = make(map[*RhombusShape]struct{})
	stage.RhombusShapes_mapString = make(map[string]*RhombusShape)
	stage.RhombusShape_stagedOrder = make(map[*RhombusShape]uint)
	stage.RhombusShapeOrder = 0

	stage.RotatedRhombusGridShapes = make(map[*RotatedRhombusGridShape]struct{})
	stage.RotatedRhombusGridShapes_mapString = make(map[string]*RotatedRhombusGridShape)
	stage.RotatedRhombusGridShape_stagedOrder = make(map[*RotatedRhombusGridShape]uint)
	stage.RotatedRhombusGridShapeOrder = 0

	stage.RotatedRhombusShapes = make(map[*RotatedRhombusShape]struct{})
	stage.RotatedRhombusShapes_mapString = make(map[string]*RotatedRhombusShape)
	stage.RotatedRhombusShape_stagedOrder = make(map[*RotatedRhombusShape]uint)
	stage.RotatedRhombusShapeOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.AxesShapes = nil
	stage.AxesShapes_mapString = nil

	stage.CircleGridShapes = nil
	stage.CircleGridShapes_mapString = nil

	stage.ExplanationTextShapes = nil
	stage.ExplanationTextShapes_mapString = nil

	stage.GridPathShapes = nil
	stage.GridPathShapes_mapString = nil

	stage.GrowthCurveRhombusGridShapes = nil
	stage.GrowthCurveRhombusGridShapes_mapString = nil

	stage.GrowthCurveRhombusShapes = nil
	stage.GrowthCurveRhombusShapes_mapString = nil

	stage.InitialRhombusGridShapes = nil
	stage.InitialRhombusGridShapes_mapString = nil

	stage.InitialRhombusShapes = nil
	stage.InitialRhombusShapes_mapString = nil

	stage.Librarys = nil
	stage.Librarys_mapString = nil

	stage.NextCircleShapes = nil
	stage.NextCircleShapes_mapString = nil

	stage.Plants = nil
	stage.Plants_mapString = nil

	stage.PlantCircumferenceShapes = nil
	stage.PlantCircumferenceShapes_mapString = nil

	stage.PlantDiagrams = nil
	stage.PlantDiagrams_mapString = nil

	stage.RhombusGridShapes = nil
	stage.RhombusGridShapes_mapString = nil

	stage.RhombusShapes = nil
	stage.RhombusShapes_mapString = nil

	stage.RotatedRhombusGridShapes = nil
	stage.RotatedRhombusGridShapes_mapString = nil

	stage.RotatedRhombusShapes = nil
	stage.RotatedRhombusShapes_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for axesshape := range stage.AxesShapes {
		axesshape.Unstage(stage)
	}

	for circlegridshape := range stage.CircleGridShapes {
		circlegridshape.Unstage(stage)
	}

	for explanationtextshape := range stage.ExplanationTextShapes {
		explanationtextshape.Unstage(stage)
	}

	for gridpathshape := range stage.GridPathShapes {
		gridpathshape.Unstage(stage)
	}

	for growthcurverhombusgridshape := range stage.GrowthCurveRhombusGridShapes {
		growthcurverhombusgridshape.Unstage(stage)
	}

	for growthcurverhombusshape := range stage.GrowthCurveRhombusShapes {
		growthcurverhombusshape.Unstage(stage)
	}

	for initialrhombusgridshape := range stage.InitialRhombusGridShapes {
		initialrhombusgridshape.Unstage(stage)
	}

	for initialrhombusshape := range stage.InitialRhombusShapes {
		initialrhombusshape.Unstage(stage)
	}

	for library := range stage.Librarys {
		library.Unstage(stage)
	}

	for nextcircleshape := range stage.NextCircleShapes {
		nextcircleshape.Unstage(stage)
	}

	for plant := range stage.Plants {
		plant.Unstage(stage)
	}

	for plantcircumferenceshape := range stage.PlantCircumferenceShapes {
		plantcircumferenceshape.Unstage(stage)
	}

	for plantdiagram := range stage.PlantDiagrams {
		plantdiagram.Unstage(stage)
	}

	for rhombusgridshape := range stage.RhombusGridShapes {
		rhombusgridshape.Unstage(stage)
	}

	for rhombusshape := range stage.RhombusShapes {
		rhombusshape.Unstage(stage)
	}

	for rotatedrhombusgridshape := range stage.RotatedRhombusGridShapes {
		rotatedrhombusgridshape.Unstage(stage)
	}

	for rotatedrhombusshape := range stage.RotatedRhombusShapes {
		rotatedrhombusshape.Unstage(stage)
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
	case map[*AxesShape]any:
		return any(&stage.AxesShapes).(*Type)
	case map[*CircleGridShape]any:
		return any(&stage.CircleGridShapes).(*Type)
	case map[*ExplanationTextShape]any:
		return any(&stage.ExplanationTextShapes).(*Type)
	case map[*GridPathShape]any:
		return any(&stage.GridPathShapes).(*Type)
	case map[*GrowthCurveRhombusGridShape]any:
		return any(&stage.GrowthCurveRhombusGridShapes).(*Type)
	case map[*GrowthCurveRhombusShape]any:
		return any(&stage.GrowthCurveRhombusShapes).(*Type)
	case map[*InitialRhombusGridShape]any:
		return any(&stage.InitialRhombusGridShapes).(*Type)
	case map[*InitialRhombusShape]any:
		return any(&stage.InitialRhombusShapes).(*Type)
	case map[*Library]any:
		return any(&stage.Librarys).(*Type)
	case map[*NextCircleShape]any:
		return any(&stage.NextCircleShapes).(*Type)
	case map[*Plant]any:
		return any(&stage.Plants).(*Type)
	case map[*PlantCircumferenceShape]any:
		return any(&stage.PlantCircumferenceShapes).(*Type)
	case map[*PlantDiagram]any:
		return any(&stage.PlantDiagrams).(*Type)
	case map[*RhombusGridShape]any:
		return any(&stage.RhombusGridShapes).(*Type)
	case map[*RhombusShape]any:
		return any(&stage.RhombusShapes).(*Type)
	case map[*RotatedRhombusGridShape]any:
		return any(&stage.RotatedRhombusGridShapes).(*Type)
	case map[*RotatedRhombusShape]any:
		return any(&stage.RotatedRhombusShapes).(*Type)
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
	case *AxesShape:
		return any(stage.AxesShapes_mapString).(map[string]Type)
	case *CircleGridShape:
		return any(stage.CircleGridShapes_mapString).(map[string]Type)
	case *ExplanationTextShape:
		return any(stage.ExplanationTextShapes_mapString).(map[string]Type)
	case *GridPathShape:
		return any(stage.GridPathShapes_mapString).(map[string]Type)
	case *GrowthCurveRhombusGridShape:
		return any(stage.GrowthCurveRhombusGridShapes_mapString).(map[string]Type)
	case *GrowthCurveRhombusShape:
		return any(stage.GrowthCurveRhombusShapes_mapString).(map[string]Type)
	case *InitialRhombusGridShape:
		return any(stage.InitialRhombusGridShapes_mapString).(map[string]Type)
	case *InitialRhombusShape:
		return any(stage.InitialRhombusShapes_mapString).(map[string]Type)
	case *Library:
		return any(stage.Librarys_mapString).(map[string]Type)
	case *NextCircleShape:
		return any(stage.NextCircleShapes_mapString).(map[string]Type)
	case *Plant:
		return any(stage.Plants_mapString).(map[string]Type)
	case *PlantCircumferenceShape:
		return any(stage.PlantCircumferenceShapes_mapString).(map[string]Type)
	case *PlantDiagram:
		return any(stage.PlantDiagrams_mapString).(map[string]Type)
	case *RhombusGridShape:
		return any(stage.RhombusGridShapes_mapString).(map[string]Type)
	case *RhombusShape:
		return any(stage.RhombusShapes_mapString).(map[string]Type)
	case *RotatedRhombusGridShape:
		return any(stage.RotatedRhombusGridShapes_mapString).(map[string]Type)
	case *RotatedRhombusShape:
		return any(stage.RotatedRhombusShapes_mapString).(map[string]Type)
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
	case AxesShape:
		return any(&stage.AxesShapes).(*map[*Type]struct{})
	case CircleGridShape:
		return any(&stage.CircleGridShapes).(*map[*Type]struct{})
	case ExplanationTextShape:
		return any(&stage.ExplanationTextShapes).(*map[*Type]struct{})
	case GridPathShape:
		return any(&stage.GridPathShapes).(*map[*Type]struct{})
	case GrowthCurveRhombusGridShape:
		return any(&stage.GrowthCurveRhombusGridShapes).(*map[*Type]struct{})
	case GrowthCurveRhombusShape:
		return any(&stage.GrowthCurveRhombusShapes).(*map[*Type]struct{})
	case InitialRhombusGridShape:
		return any(&stage.InitialRhombusGridShapes).(*map[*Type]struct{})
	case InitialRhombusShape:
		return any(&stage.InitialRhombusShapes).(*map[*Type]struct{})
	case Library:
		return any(&stage.Librarys).(*map[*Type]struct{})
	case NextCircleShape:
		return any(&stage.NextCircleShapes).(*map[*Type]struct{})
	case Plant:
		return any(&stage.Plants).(*map[*Type]struct{})
	case PlantCircumferenceShape:
		return any(&stage.PlantCircumferenceShapes).(*map[*Type]struct{})
	case PlantDiagram:
		return any(&stage.PlantDiagrams).(*map[*Type]struct{})
	case RhombusGridShape:
		return any(&stage.RhombusGridShapes).(*map[*Type]struct{})
	case RhombusShape:
		return any(&stage.RhombusShapes).(*map[*Type]struct{})
	case RotatedRhombusGridShape:
		return any(&stage.RotatedRhombusGridShapes).(*map[*Type]struct{})
	case RotatedRhombusShape:
		return any(&stage.RotatedRhombusShapes).(*map[*Type]struct{})
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
	case *AxesShape:
		return any(&stage.AxesShapes).(*map[Type]struct{})
	case *CircleGridShape:
		return any(&stage.CircleGridShapes).(*map[Type]struct{})
	case *ExplanationTextShape:
		return any(&stage.ExplanationTextShapes).(*map[Type]struct{})
	case *GridPathShape:
		return any(&stage.GridPathShapes).(*map[Type]struct{})
	case *GrowthCurveRhombusGridShape:
		return any(&stage.GrowthCurveRhombusGridShapes).(*map[Type]struct{})
	case *GrowthCurveRhombusShape:
		return any(&stage.GrowthCurveRhombusShapes).(*map[Type]struct{})
	case *InitialRhombusGridShape:
		return any(&stage.InitialRhombusGridShapes).(*map[Type]struct{})
	case *InitialRhombusShape:
		return any(&stage.InitialRhombusShapes).(*map[Type]struct{})
	case *Library:
		return any(&stage.Librarys).(*map[Type]struct{})
	case *NextCircleShape:
		return any(&stage.NextCircleShapes).(*map[Type]struct{})
	case *Plant:
		return any(&stage.Plants).(*map[Type]struct{})
	case *PlantCircumferenceShape:
		return any(&stage.PlantCircumferenceShapes).(*map[Type]struct{})
	case *PlantDiagram:
		return any(&stage.PlantDiagrams).(*map[Type]struct{})
	case *RhombusGridShape:
		return any(&stage.RhombusGridShapes).(*map[Type]struct{})
	case *RhombusShape:
		return any(&stage.RhombusShapes).(*map[Type]struct{})
	case *RotatedRhombusGridShape:
		return any(&stage.RotatedRhombusGridShapes).(*map[Type]struct{})
	case *RotatedRhombusShape:
		return any(&stage.RotatedRhombusShapes).(*map[Type]struct{})
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
	case AxesShape:
		return any(&stage.AxesShapes_mapString).(*map[string]*Type)
	case CircleGridShape:
		return any(&stage.CircleGridShapes_mapString).(*map[string]*Type)
	case ExplanationTextShape:
		return any(&stage.ExplanationTextShapes_mapString).(*map[string]*Type)
	case GridPathShape:
		return any(&stage.GridPathShapes_mapString).(*map[string]*Type)
	case GrowthCurveRhombusGridShape:
		return any(&stage.GrowthCurveRhombusGridShapes_mapString).(*map[string]*Type)
	case GrowthCurveRhombusShape:
		return any(&stage.GrowthCurveRhombusShapes_mapString).(*map[string]*Type)
	case InitialRhombusGridShape:
		return any(&stage.InitialRhombusGridShapes_mapString).(*map[string]*Type)
	case InitialRhombusShape:
		return any(&stage.InitialRhombusShapes_mapString).(*map[string]*Type)
	case Library:
		return any(&stage.Librarys_mapString).(*map[string]*Type)
	case NextCircleShape:
		return any(&stage.NextCircleShapes_mapString).(*map[string]*Type)
	case Plant:
		return any(&stage.Plants_mapString).(*map[string]*Type)
	case PlantCircumferenceShape:
		return any(&stage.PlantCircumferenceShapes_mapString).(*map[string]*Type)
	case PlantDiagram:
		return any(&stage.PlantDiagrams_mapString).(*map[string]*Type)
	case RhombusGridShape:
		return any(&stage.RhombusGridShapes_mapString).(*map[string]*Type)
	case RhombusShape:
		return any(&stage.RhombusShapes_mapString).(*map[string]*Type)
	case RotatedRhombusGridShape:
		return any(&stage.RotatedRhombusGridShapes_mapString).(*map[string]*Type)
	case RotatedRhombusShape:
		return any(&stage.RotatedRhombusShapes_mapString).(*map[string]*Type)
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
	case AxesShape:
		return any(&AxesShape{
			// Initialisation of associations
		}).(*Type)
	case CircleGridShape:
		return any(&CircleGridShape{
			// Initialisation of associations
		}).(*Type)
	case ExplanationTextShape:
		return any(&ExplanationTextShape{
			// Initialisation of associations
		}).(*Type)
	case GridPathShape:
		return any(&GridPathShape{
			// Initialisation of associations
		}).(*Type)
	case GrowthCurveRhombusGridShape:
		return any(&GrowthCurveRhombusGridShape{
			// Initialisation of associations
			// field is initialized with an instance of GrowthCurveRhombusShape with the name of the field
			GrowthCurveRhombusShapes: []*GrowthCurveRhombusShape{{Name: "GrowthCurveRhombusShapes"}},
		}).(*Type)
	case GrowthCurveRhombusShape:
		return any(&GrowthCurveRhombusShape{
			// Initialisation of associations
		}).(*Type)
	case InitialRhombusGridShape:
		return any(&InitialRhombusGridShape{
			// Initialisation of associations
			// field is initialized with an instance of InitialRhombusShape with the name of the field
			InitialRhombusShapes: []*InitialRhombusShape{{Name: "InitialRhombusShapes"}},
		}).(*Type)
	case InitialRhombusShape:
		return any(&InitialRhombusShape{
			// Initialisation of associations
		}).(*Type)
	case Library:
		return any(&Library{
			// Initialisation of associations
			// field is initialized with an instance of Library with the name of the field
			SubLibraries: []*Library{{Name: "SubLibraries"}},
			// field is initialized with an instance of Plant with the name of the field
			Plants: []*Plant{{Name: "Plants"}},
		}).(*Type)
	case NextCircleShape:
		return any(&NextCircleShape{
			// Initialisation of associations
		}).(*Type)
	case Plant:
		return any(&Plant{
			// Initialisation of associations
			// field is initialized with an instance of PlantDiagram with the name of the field
			PlantDiagrams: []*PlantDiagram{{Name: "PlantDiagrams"}},
			// field is initialized with an instance of AxesShape with the name of the field
			AxesShape: &AxesShape{Name: "AxesShape"},
			// field is initialized with an instance of RhombusShape with the name of the field
			ReferenceRhombus: &RhombusShape{Name: "ReferenceRhombus"},
			// field is initialized with an instance of PlantCircumferenceShape with the name of the field
			PlantCircumferenceShape: &PlantCircumferenceShape{Name: "PlantCircumferenceShape"},
			// field is initialized with an instance of GridPathShape with the name of the field
			GridPathShape: &GridPathShape{Name: "GridPathShape"},
			// field is initialized with an instance of InitialRhombusGridShape with the name of the field
			InitialRhombusGridShape: &InitialRhombusGridShape{Name: "InitialRhombusGridShape"},
			// field is initialized with an instance of ExplanationTextShape with the name of the field
			ExplanationTextShape: &ExplanationTextShape{Name: "ExplanationTextShape"},
			// field is initialized with an instance of RhombusShape with the name of the field
			RotatedReferenceRhombus: &RhombusShape{Name: "RotatedReferenceRhombus"},
			// field is initialized with an instance of PlantCircumferenceShape with the name of the field
			RotatedPlantCircumferenceShape: &PlantCircumferenceShape{Name: "RotatedPlantCircumferenceShape"},
			// field is initialized with an instance of GridPathShape with the name of the field
			RotatedGridPathShape: &GridPathShape{Name: "RotatedGridPathShape"},
			// field is initialized with an instance of RotatedRhombusGridShape with the name of the field
			RotatedRhombusGridShape2: &RotatedRhombusGridShape{Name: "RotatedRhombusGridShape2"},
			// field is initialized with an instance of GrowthCurveRhombusGridShape with the name of the field
			GrowthCurveRhombusGridShape: &GrowthCurveRhombusGridShape{Name: "GrowthCurveRhombusGridShape"},
		}).(*Type)
	case PlantCircumferenceShape:
		return any(&PlantCircumferenceShape{
			// Initialisation of associations
		}).(*Type)
	case PlantDiagram:
		return any(&PlantDiagram{
			// Initialisation of associations
		}).(*Type)
	case RhombusGridShape:
		return any(&RhombusGridShape{
			// Initialisation of associations
			// field is initialized with an instance of RhombusShape with the name of the field
			RhombusShapes: []*RhombusShape{{Name: "RhombusShapes"}},
		}).(*Type)
	case RhombusShape:
		return any(&RhombusShape{
			// Initialisation of associations
		}).(*Type)
	case RotatedRhombusGridShape:
		return any(&RotatedRhombusGridShape{
			// Initialisation of associations
			// field is initialized with an instance of RotatedRhombusShape with the name of the field
			RotatedRhombusShapes: []*RotatedRhombusShape{{Name: "RotatedRhombusShapes"}},
		}).(*Type)
	case RotatedRhombusShape:
		return any(&RotatedRhombusShape{
			// Initialisation of associations
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
	// reverse maps of direct associations of AxesShape
	case AxesShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CircleGridShape
	case CircleGridShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ExplanationTextShape
	case ExplanationTextShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GridPathShape
	case GridPathShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GrowthCurveRhombusGridShape
	case GrowthCurveRhombusGridShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GrowthCurveRhombusShape
	case GrowthCurveRhombusShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of InitialRhombusGridShape
	case InitialRhombusGridShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of InitialRhombusShape
	case InitialRhombusShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Library
	case Library:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NextCircleShape
	case NextCircleShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Plant
	case Plant:
		switch fieldname {
		// insertion point for per direct association field
		case "AxesShape":
			res := make(map[*AxesShape][]*Plant)
			for plant := range stage.Plants {
				if plant.AxesShape != nil {
					axesshape_ := plant.AxesShape
					var plants []*Plant
					_, ok := res[axesshape_]
					if ok {
						plants = res[axesshape_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[axesshape_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "ReferenceRhombus":
			res := make(map[*RhombusShape][]*Plant)
			for plant := range stage.Plants {
				if plant.ReferenceRhombus != nil {
					rhombusshape_ := plant.ReferenceRhombus
					var plants []*Plant
					_, ok := res[rhombusshape_]
					if ok {
						plants = res[rhombusshape_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[rhombusshape_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "PlantCircumferenceShape":
			res := make(map[*PlantCircumferenceShape][]*Plant)
			for plant := range stage.Plants {
				if plant.PlantCircumferenceShape != nil {
					plantcircumferenceshape_ := plant.PlantCircumferenceShape
					var plants []*Plant
					_, ok := res[plantcircumferenceshape_]
					if ok {
						plants = res[plantcircumferenceshape_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[plantcircumferenceshape_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "GridPathShape":
			res := make(map[*GridPathShape][]*Plant)
			for plant := range stage.Plants {
				if plant.GridPathShape != nil {
					gridpathshape_ := plant.GridPathShape
					var plants []*Plant
					_, ok := res[gridpathshape_]
					if ok {
						plants = res[gridpathshape_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[gridpathshape_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "InitialRhombusGridShape":
			res := make(map[*InitialRhombusGridShape][]*Plant)
			for plant := range stage.Plants {
				if plant.InitialRhombusGridShape != nil {
					initialrhombusgridshape_ := plant.InitialRhombusGridShape
					var plants []*Plant
					_, ok := res[initialrhombusgridshape_]
					if ok {
						plants = res[initialrhombusgridshape_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[initialrhombusgridshape_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "ExplanationTextShape":
			res := make(map[*ExplanationTextShape][]*Plant)
			for plant := range stage.Plants {
				if plant.ExplanationTextShape != nil {
					explanationtextshape_ := plant.ExplanationTextShape
					var plants []*Plant
					_, ok := res[explanationtextshape_]
					if ok {
						plants = res[explanationtextshape_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[explanationtextshape_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "RotatedReferenceRhombus":
			res := make(map[*RhombusShape][]*Plant)
			for plant := range stage.Plants {
				if plant.RotatedReferenceRhombus != nil {
					rhombusshape_ := plant.RotatedReferenceRhombus
					var plants []*Plant
					_, ok := res[rhombusshape_]
					if ok {
						plants = res[rhombusshape_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[rhombusshape_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "RotatedPlantCircumferenceShape":
			res := make(map[*PlantCircumferenceShape][]*Plant)
			for plant := range stage.Plants {
				if plant.RotatedPlantCircumferenceShape != nil {
					plantcircumferenceshape_ := plant.RotatedPlantCircumferenceShape
					var plants []*Plant
					_, ok := res[plantcircumferenceshape_]
					if ok {
						plants = res[plantcircumferenceshape_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[plantcircumferenceshape_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "RotatedGridPathShape":
			res := make(map[*GridPathShape][]*Plant)
			for plant := range stage.Plants {
				if plant.RotatedGridPathShape != nil {
					gridpathshape_ := plant.RotatedGridPathShape
					var plants []*Plant
					_, ok := res[gridpathshape_]
					if ok {
						plants = res[gridpathshape_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[gridpathshape_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "RotatedRhombusGridShape2":
			res := make(map[*RotatedRhombusGridShape][]*Plant)
			for plant := range stage.Plants {
				if plant.RotatedRhombusGridShape2 != nil {
					rotatedrhombusgridshape_ := plant.RotatedRhombusGridShape2
					var plants []*Plant
					_, ok := res[rotatedrhombusgridshape_]
					if ok {
						plants = res[rotatedrhombusgridshape_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[rotatedrhombusgridshape_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "GrowthCurveRhombusGridShape":
			res := make(map[*GrowthCurveRhombusGridShape][]*Plant)
			for plant := range stage.Plants {
				if plant.GrowthCurveRhombusGridShape != nil {
					growthcurverhombusgridshape_ := plant.GrowthCurveRhombusGridShape
					var plants []*Plant
					_, ok := res[growthcurverhombusgridshape_]
					if ok {
						plants = res[growthcurverhombusgridshape_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[growthcurverhombusgridshape_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of PlantCircumferenceShape
	case PlantCircumferenceShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PlantDiagram
	case PlantDiagram:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RhombusGridShape
	case RhombusGridShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RhombusShape
	case RhombusShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RotatedRhombusGridShape
	case RotatedRhombusGridShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RotatedRhombusShape
	case RotatedRhombusShape:
		switch fieldname {
		// insertion point for per direct association field
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
	// reverse maps of direct associations of AxesShape
	case AxesShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CircleGridShape
	case CircleGridShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ExplanationTextShape
	case ExplanationTextShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GridPathShape
	case GridPathShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GrowthCurveRhombusGridShape
	case GrowthCurveRhombusGridShape:
		switch fieldname {
		// insertion point for per direct association field
		case "GrowthCurveRhombusShapes":
			res := make(map[*GrowthCurveRhombusShape][]*GrowthCurveRhombusGridShape)
			for growthcurverhombusgridshape := range stage.GrowthCurveRhombusGridShapes {
				for _, growthcurverhombusshape_ := range growthcurverhombusgridshape.GrowthCurveRhombusShapes {
					res[growthcurverhombusshape_] = append(res[growthcurverhombusshape_], growthcurverhombusgridshape)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GrowthCurveRhombusShape
	case GrowthCurveRhombusShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of InitialRhombusGridShape
	case InitialRhombusGridShape:
		switch fieldname {
		// insertion point for per direct association field
		case "InitialRhombusShapes":
			res := make(map[*InitialRhombusShape][]*InitialRhombusGridShape)
			for initialrhombusgridshape := range stage.InitialRhombusGridShapes {
				for _, initialrhombusshape_ := range initialrhombusgridshape.InitialRhombusShapes {
					res[initialrhombusshape_] = append(res[initialrhombusshape_], initialrhombusgridshape)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of InitialRhombusShape
	case InitialRhombusShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Library
	case Library:
		switch fieldname {
		// insertion point for per direct association field
		case "SubLibraries":
			res := make(map[*Library][]*Library)
			for library := range stage.Librarys {
				for _, library_ := range library.SubLibraries {
					res[library_] = append(res[library_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Plants":
			res := make(map[*Plant][]*Library)
			for library := range stage.Librarys {
				for _, plant_ := range library.Plants {
					res[plant_] = append(res[plant_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of NextCircleShape
	case NextCircleShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Plant
	case Plant:
		switch fieldname {
		// insertion point for per direct association field
		case "PlantDiagrams":
			res := make(map[*PlantDiagram][]*Plant)
			for plant := range stage.Plants {
				for _, plantdiagram_ := range plant.PlantDiagrams {
					res[plantdiagram_] = append(res[plantdiagram_], plant)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of PlantCircumferenceShape
	case PlantCircumferenceShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PlantDiagram
	case PlantDiagram:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RhombusGridShape
	case RhombusGridShape:
		switch fieldname {
		// insertion point for per direct association field
		case "RhombusShapes":
			res := make(map[*RhombusShape][]*RhombusGridShape)
			for rhombusgridshape := range stage.RhombusGridShapes {
				for _, rhombusshape_ := range rhombusgridshape.RhombusShapes {
					res[rhombusshape_] = append(res[rhombusshape_], rhombusgridshape)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of RhombusShape
	case RhombusShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RotatedRhombusGridShape
	case RotatedRhombusGridShape:
		switch fieldname {
		// insertion point for per direct association field
		case "RotatedRhombusShapes":
			res := make(map[*RotatedRhombusShape][]*RotatedRhombusGridShape)
			for rotatedrhombusgridshape := range stage.RotatedRhombusGridShapes {
				for _, rotatedrhombusshape_ := range rotatedrhombusgridshape.RotatedRhombusShapes {
					res[rotatedrhombusshape_] = append(res[rotatedrhombusshape_], rotatedrhombusgridshape)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of RotatedRhombusShape
	case RotatedRhombusShape:
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
	case *AxesShape:
		res = "AxesShape"
	case *CircleGridShape:
		res = "CircleGridShape"
	case *ExplanationTextShape:
		res = "ExplanationTextShape"
	case *GridPathShape:
		res = "GridPathShape"
	case *GrowthCurveRhombusGridShape:
		res = "GrowthCurveRhombusGridShape"
	case *GrowthCurveRhombusShape:
		res = "GrowthCurveRhombusShape"
	case *InitialRhombusGridShape:
		res = "InitialRhombusGridShape"
	case *InitialRhombusShape:
		res = "InitialRhombusShape"
	case *Library:
		res = "Library"
	case *NextCircleShape:
		res = "NextCircleShape"
	case *Plant:
		res = "Plant"
	case *PlantCircumferenceShape:
		res = "PlantCircumferenceShape"
	case *PlantDiagram:
		res = "PlantDiagram"
	case *RhombusGridShape:
		res = "RhombusGridShape"
	case *RhombusShape:
		res = "RhombusShape"
	case *RotatedRhombusGridShape:
		res = "RotatedRhombusGridShape"
	case *RotatedRhombusShape:
		res = "RotatedRhombusShape"
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
	case *AxesShape:
		var rf ReverseField
		_ = rf
	case *CircleGridShape:
		var rf ReverseField
		_ = rf
	case *ExplanationTextShape:
		var rf ReverseField
		_ = rf
	case *GridPathShape:
		var rf ReverseField
		_ = rf
	case *GrowthCurveRhombusGridShape:
		var rf ReverseField
		_ = rf
	case *GrowthCurveRhombusShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GrowthCurveRhombusGridShape"
		rf.Fieldname = "GrowthCurveRhombusShapes"
		res = append(res, rf)
	case *InitialRhombusGridShape:
		var rf ReverseField
		_ = rf
	case *InitialRhombusShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "InitialRhombusGridShape"
		rf.Fieldname = "InitialRhombusShapes"
		res = append(res, rf)
	case *Library:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Library"
		rf.Fieldname = "SubLibraries"
		res = append(res, rf)
	case *NextCircleShape:
		var rf ReverseField
		_ = rf
	case *Plant:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Library"
		rf.Fieldname = "Plants"
		res = append(res, rf)
	case *PlantCircumferenceShape:
		var rf ReverseField
		_ = rf
	case *PlantDiagram:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Plant"
		rf.Fieldname = "PlantDiagrams"
		res = append(res, rf)
	case *RhombusGridShape:
		var rf ReverseField
		_ = rf
	case *RhombusShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "RhombusGridShape"
		rf.Fieldname = "RhombusShapes"
		res = append(res, rf)
	case *RotatedRhombusGridShape:
		var rf ReverseField
		_ = rf
	case *RotatedRhombusShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "RotatedRhombusGridShape"
		rf.Fieldname = "RotatedRhombusShapes"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
func (axesshape *AxesShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "LengthX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "LengthY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsWithHiddenHandle",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (circlegridshape *CircleGridShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (explanationtextshape *ExplanationTextShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (gridpathshape *GridPathShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "GrowthCurveRhombusShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GrowthCurveRhombusShape",
		},
	}
	return
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "InitialRhombusShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "InitialRhombusShape",
		},
	}
	return
}

func (initialrhombusshape *InitialRhombusShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeFloat,
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
			Name:                 "SubLibraries",
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
			Name:                 "Plants",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Plant",
		},
	}
	return
}

func (nextcircleshape *NextCircleShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (plant *Plant) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "N",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "M",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "RhombusInsideAngle",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "RhombusSideLength",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ShiftToNearestCircle",
			GongFieldValueType: GongFieldValueTypeInt,
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
			Name:               "IsSelected",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsPlantDiagramsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "PlantDiagrams",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "PlantDiagram",
		},
		{
			Name:                 "AxesShape",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AxesShape",
		},
		{
			Name:                 "ReferenceRhombus",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "RhombusShape",
		},
		{
			Name:                 "PlantCircumferenceShape",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PlantCircumferenceShape",
		},
		{
			Name:                 "GridPathShape",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "GridPathShape",
		},
		{
			Name:                 "InitialRhombusGridShape",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "InitialRhombusGridShape",
		},
		{
			Name:                 "ExplanationTextShape",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ExplanationTextShape",
		},
		{
			Name:                 "RotatedReferenceRhombus",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "RhombusShape",
		},
		{
			Name:                 "RotatedPlantCircumferenceShape",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PlantCircumferenceShape",
		},
		{
			Name:                 "RotatedGridPathShape",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "GridPathShape",
		},
		{
			Name:                 "RotatedRhombusGridShape2",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "RotatedRhombusGridShape",
		},
		{
			Name:                 "GrowthCurveRhombusGridShape",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "GrowthCurveRhombusGridShape",
		},
	}
	return
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "AngleDegree",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Length",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (plantdiagram *PlantDiagram) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "OriginX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "OriginY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHiddenAxesShape",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenReferenceRhombus",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenPlantCircumferenceShape",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenGridPathShape",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenRhombusGridShape",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenExplanationTextShape",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenRotatedReferenceRhombus",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenRotatedPlantCircumferenceShape",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenRotatedGridPathShape",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenRotatedRhombusGridShape",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenGrowthPathRhombusGridShape",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsChecked",
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

func (rhombusgridshape *RhombusGridShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "RhombusShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "RhombusShape",
		},
	}
	return
}

func (rhombusshape *RhombusShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "RotatedRhombusShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "RotatedRhombusShape",
		},
	}
	return
}

func (rotatedrhombusshape *RotatedRhombusShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeFloat,
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
func (axesshape *AxesShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = axesshape.Name
	case "LengthX":
		res.valueString = fmt.Sprintf("%f", axesshape.LengthX)
		res.valueFloat = axesshape.LengthX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LengthY":
		res.valueString = fmt.Sprintf("%f", axesshape.LengthY)
		res.valueFloat = axesshape.LengthY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsWithHiddenHandle":
		res.valueString = fmt.Sprintf("%t", axesshape.IsWithHiddenHandle)
		res.valueBool = axesshape.IsWithHiddenHandle
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (circlegridshape *CircleGridShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = circlegridshape.Name
	}
	return
}

func (explanationtextshape *ExplanationTextShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = explanationtextshape.Name
	}
	return
}

func (gridpathshape *GridPathShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gridpathshape.Name
	}
	return
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = growthcurverhombusgridshape.Name
	case "GrowthCurveRhombusShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range growthcurverhombusgridshape.GrowthCurveRhombusShapes {
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

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = growthcurverhombusshape.Name
	case "X":
		res.valueString = fmt.Sprintf("%f", growthcurverhombusshape.X)
		res.valueFloat = growthcurverhombusshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", growthcurverhombusshape.Y)
		res.valueFloat = growthcurverhombusshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = initialrhombusgridshape.Name
	case "InitialRhombusShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range initialrhombusgridshape.InitialRhombusShapes {
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

func (initialrhombusshape *InitialRhombusShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = initialrhombusshape.Name
	case "X":
		res.valueString = fmt.Sprintf("%f", initialrhombusshape.X)
		res.valueFloat = initialrhombusshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", initialrhombusshape.Y)
		res.valueFloat = initialrhombusshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (library *Library) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = library.Name
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
	case "NbPixPerCharacter":
		res.valueString = fmt.Sprintf("%f", library.NbPixPerCharacter)
		res.valueFloat = library.NbPixPerCharacter
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LogoSVGFile":
		res.valueString = library.LogoSVGFile
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
	case "Plants":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.Plants {
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

func (nextcircleshape *NextCircleShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = nextcircleshape.Name
	}
	return
}

func (plant *Plant) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = plant.Name
	case "N":
		res.valueString = fmt.Sprintf("%d", plant.N)
		res.valueInt = plant.N
		res.GongFieldValueType = GongFieldValueTypeInt
	case "M":
		res.valueString = fmt.Sprintf("%d", plant.M)
		res.valueInt = plant.M
		res.GongFieldValueType = GongFieldValueTypeInt
	case "RhombusInsideAngle":
		res.valueString = fmt.Sprintf("%f", plant.RhombusInsideAngle)
		res.valueFloat = plant.RhombusInsideAngle
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RhombusSideLength":
		res.valueString = fmt.Sprintf("%f", plant.RhombusSideLength)
		res.valueFloat = plant.RhombusSideLength
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ShiftToNearestCircle":
		res.valueString = fmt.Sprintf("%d", plant.ShiftToNearestCircle)
		res.valueInt = plant.ShiftToNearestCircle
		res.GongFieldValueType = GongFieldValueTypeInt
	case "ComputedPrefix":
		res.valueString = plant.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", plant.IsExpanded)
		res.valueBool = plant.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsSelected":
		res.valueString = fmt.Sprintf("%t", plant.IsSelected)
		res.valueBool = plant.IsSelected
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsPlantDiagramsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", plant.IsPlantDiagramsNodeExpanded)
		res.valueBool = plant.IsPlantDiagramsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "PlantDiagrams":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range plant.PlantDiagrams {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "AxesShape":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.AxesShape != nil {
			res.valueString = plant.AxesShape.Name
			res.ids = plant.AxesShape.GongGetUUID(stage)
		}
	case "ReferenceRhombus":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.ReferenceRhombus != nil {
			res.valueString = plant.ReferenceRhombus.Name
			res.ids = plant.ReferenceRhombus.GongGetUUID(stage)
		}
	case "PlantCircumferenceShape":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.PlantCircumferenceShape != nil {
			res.valueString = plant.PlantCircumferenceShape.Name
			res.ids = plant.PlantCircumferenceShape.GongGetUUID(stage)
		}
	case "GridPathShape":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.GridPathShape != nil {
			res.valueString = plant.GridPathShape.Name
			res.ids = plant.GridPathShape.GongGetUUID(stage)
		}
	case "InitialRhombusGridShape":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.InitialRhombusGridShape != nil {
			res.valueString = plant.InitialRhombusGridShape.Name
			res.ids = plant.InitialRhombusGridShape.GongGetUUID(stage)
		}
	case "ExplanationTextShape":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.ExplanationTextShape != nil {
			res.valueString = plant.ExplanationTextShape.Name
			res.ids = plant.ExplanationTextShape.GongGetUUID(stage)
		}
	case "RotatedReferenceRhombus":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.RotatedReferenceRhombus != nil {
			res.valueString = plant.RotatedReferenceRhombus.Name
			res.ids = plant.RotatedReferenceRhombus.GongGetUUID(stage)
		}
	case "RotatedPlantCircumferenceShape":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.RotatedPlantCircumferenceShape != nil {
			res.valueString = plant.RotatedPlantCircumferenceShape.Name
			res.ids = plant.RotatedPlantCircumferenceShape.GongGetUUID(stage)
		}
	case "RotatedGridPathShape":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.RotatedGridPathShape != nil {
			res.valueString = plant.RotatedGridPathShape.Name
			res.ids = plant.RotatedGridPathShape.GongGetUUID(stage)
		}
	case "RotatedRhombusGridShape2":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.RotatedRhombusGridShape2 != nil {
			res.valueString = plant.RotatedRhombusGridShape2.Name
			res.ids = plant.RotatedRhombusGridShape2.GongGetUUID(stage)
		}
	case "GrowthCurveRhombusGridShape":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.GrowthCurveRhombusGridShape != nil {
			res.valueString = plant.GrowthCurveRhombusGridShape.Name
			res.ids = plant.GrowthCurveRhombusGridShape.GongGetUUID(stage)
		}
	}
	return
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = plantcircumferenceshape.Name
	case "AngleDegree":
		res.valueString = fmt.Sprintf("%f", plantcircumferenceshape.AngleDegree)
		res.valueFloat = plantcircumferenceshape.AngleDegree
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Length":
		res.valueString = fmt.Sprintf("%f", plantcircumferenceshape.Length)
		res.valueFloat = plantcircumferenceshape.Length
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (plantdiagram *PlantDiagram) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = plantdiagram.Name
	case "OriginX":
		res.valueString = fmt.Sprintf("%f", plantdiagram.OriginX)
		res.valueFloat = plantdiagram.OriginX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "OriginY":
		res.valueString = fmt.Sprintf("%f", plantdiagram.OriginY)
		res.valueFloat = plantdiagram.OriginY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHiddenAxesShape":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenAxesShape)
		res.valueBool = plantdiagram.IsHiddenAxesShape
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenReferenceRhombus":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenReferenceRhombus)
		res.valueBool = plantdiagram.IsHiddenReferenceRhombus
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenPlantCircumferenceShape":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenPlantCircumferenceShape)
		res.valueBool = plantdiagram.IsHiddenPlantCircumferenceShape
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenGridPathShape":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenGridPathShape)
		res.valueBool = plantdiagram.IsHiddenGridPathShape
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenRhombusGridShape":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenRhombusGridShape)
		res.valueBool = plantdiagram.IsHiddenRhombusGridShape
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenExplanationTextShape":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenExplanationTextShape)
		res.valueBool = plantdiagram.IsHiddenExplanationTextShape
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenRotatedReferenceRhombus":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenRotatedReferenceRhombus)
		res.valueBool = plantdiagram.IsHiddenRotatedReferenceRhombus
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenRotatedPlantCircumferenceShape":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenRotatedPlantCircumferenceShape)
		res.valueBool = plantdiagram.IsHiddenRotatedPlantCircumferenceShape
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenRotatedGridPathShape":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenRotatedGridPathShape)
		res.valueBool = plantdiagram.IsHiddenRotatedGridPathShape
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenRotatedRhombusGridShape":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenRotatedRhombusGridShape)
		res.valueBool = plantdiagram.IsHiddenRotatedRhombusGridShape
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenGrowthPathRhombusGridShape":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenGrowthPathRhombusGridShape)
		res.valueBool = plantdiagram.IsHiddenGrowthPathRhombusGridShape
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsChecked":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsChecked)
		res.valueBool = plantdiagram.IsChecked
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ComputedPrefix":
		res.valueString = plantdiagram.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsExpanded)
		res.valueBool = plantdiagram.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (rhombusgridshape *RhombusGridShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = rhombusgridshape.Name
	case "RhombusShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range rhombusgridshape.RhombusShapes {
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

func (rhombusshape *RhombusShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = rhombusshape.Name
	case "X":
		res.valueString = fmt.Sprintf("%f", rhombusshape.X)
		res.valueFloat = rhombusshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", rhombusshape.Y)
		res.valueFloat = rhombusshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = rotatedrhombusgridshape.Name
	case "RotatedRhombusShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range rotatedrhombusgridshape.RotatedRhombusShapes {
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

func (rotatedrhombusshape *RotatedRhombusShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = rotatedrhombusshape.Name
	case "X":
		res.valueString = fmt.Sprintf("%f", rotatedrhombusshape.X)
		res.valueFloat = rotatedrhombusshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", rotatedrhombusshape.Y)
		res.valueFloat = rotatedrhombusshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (axesshape *AxesShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		axesshape.Name = value.GetValueString()
	case "LengthX":
		axesshape.LengthX = value.GetValueFloat()
	case "LengthY":
		axesshape.LengthY = value.GetValueFloat()
	case "IsWithHiddenHandle":
		axesshape.IsWithHiddenHandle = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (circlegridshape *CircleGridShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		circlegridshape.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (explanationtextshape *ExplanationTextShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		explanationtextshape.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (gridpathshape *GridPathShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		gridpathshape.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		growthcurverhombusgridshape.Name = value.GetValueString()
	case "GrowthCurveRhombusShapes":
		growthcurverhombusgridshape.GrowthCurveRhombusShapes = make([]*GrowthCurveRhombusShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.GrowthCurveRhombusShapes {
					if stage.GrowthCurveRhombusShape_stagedOrder[__instance__] == uint(id) {
						growthcurverhombusgridshape.GrowthCurveRhombusShapes = append(growthcurverhombusgridshape.GrowthCurveRhombusShapes, __instance__)
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

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		growthcurverhombusshape.Name = value.GetValueString()
	case "X":
		growthcurverhombusshape.X = value.GetValueFloat()
	case "Y":
		growthcurverhombusshape.Y = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		initialrhombusgridshape.Name = value.GetValueString()
	case "InitialRhombusShapes":
		initialrhombusgridshape.InitialRhombusShapes = make([]*InitialRhombusShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.InitialRhombusShapes {
					if stage.InitialRhombusShape_stagedOrder[__instance__] == uint(id) {
						initialrhombusgridshape.InitialRhombusShapes = append(initialrhombusgridshape.InitialRhombusShapes, __instance__)
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

func (initialrhombusshape *InitialRhombusShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		initialrhombusshape.Name = value.GetValueString()
	case "X":
		initialrhombusshape.X = value.GetValueFloat()
	case "Y":
		initialrhombusshape.Y = value.GetValueFloat()
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
	case "NbPixPerCharacter":
		library.NbPixPerCharacter = value.GetValueFloat()
	case "LogoSVGFile":
		library.LogoSVGFile = value.GetValueString()
	case "ComputedPrefix":
		library.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		library.IsExpanded = value.GetValueBool()
	case "IsRootLibrary":
		library.IsRootLibrary = value.GetValueBool()
	case "Plants":
		library.Plants = make([]*Plant, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Plants {
					if stage.Plant_stagedOrder[__instance__] == uint(id) {
						library.Plants = append(library.Plants, __instance__)
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

func (nextcircleshape *NextCircleShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		nextcircleshape.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (plant *Plant) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		plant.Name = value.GetValueString()
	case "N":
		plant.N = int(value.GetValueInt())
	case "M":
		plant.M = int(value.GetValueInt())
	case "RhombusInsideAngle":
		plant.RhombusInsideAngle = value.GetValueFloat()
	case "RhombusSideLength":
		plant.RhombusSideLength = value.GetValueFloat()
	case "ShiftToNearestCircle":
		plant.ShiftToNearestCircle = int(value.GetValueInt())
	case "ComputedPrefix":
		plant.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		plant.IsExpanded = value.GetValueBool()
	case "IsSelected":
		plant.IsSelected = value.GetValueBool()
	case "IsPlantDiagramsNodeExpanded":
		plant.IsPlantDiagramsNodeExpanded = value.GetValueBool()
	case "PlantDiagrams":
		plant.PlantDiagrams = make([]*PlantDiagram, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.PlantDiagrams {
					if stage.PlantDiagram_stagedOrder[__instance__] == uint(id) {
						plant.PlantDiagrams = append(plant.PlantDiagrams, __instance__)
						break
					}
				}
			}
		}
	case "AxesShape":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.AxesShape = nil
			for __instance__ := range stage.AxesShapes {
				if stage.AxesShape_stagedOrder[__instance__] == uint(id) {
					plant.AxesShape = __instance__
					break
				}
			}
		}
	case "ReferenceRhombus":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.ReferenceRhombus = nil
			for __instance__ := range stage.RhombusShapes {
				if stage.RhombusShape_stagedOrder[__instance__] == uint(id) {
					plant.ReferenceRhombus = __instance__
					break
				}
			}
		}
	case "PlantCircumferenceShape":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.PlantCircumferenceShape = nil
			for __instance__ := range stage.PlantCircumferenceShapes {
				if stage.PlantCircumferenceShape_stagedOrder[__instance__] == uint(id) {
					plant.PlantCircumferenceShape = __instance__
					break
				}
			}
		}
	case "GridPathShape":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.GridPathShape = nil
			for __instance__ := range stage.GridPathShapes {
				if stage.GridPathShape_stagedOrder[__instance__] == uint(id) {
					plant.GridPathShape = __instance__
					break
				}
			}
		}
	case "InitialRhombusGridShape":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.InitialRhombusGridShape = nil
			for __instance__ := range stage.InitialRhombusGridShapes {
				if stage.InitialRhombusGridShape_stagedOrder[__instance__] == uint(id) {
					plant.InitialRhombusGridShape = __instance__
					break
				}
			}
		}
	case "ExplanationTextShape":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.ExplanationTextShape = nil
			for __instance__ := range stage.ExplanationTextShapes {
				if stage.ExplanationTextShape_stagedOrder[__instance__] == uint(id) {
					plant.ExplanationTextShape = __instance__
					break
				}
			}
		}
	case "RotatedReferenceRhombus":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.RotatedReferenceRhombus = nil
			for __instance__ := range stage.RhombusShapes {
				if stage.RhombusShape_stagedOrder[__instance__] == uint(id) {
					plant.RotatedReferenceRhombus = __instance__
					break
				}
			}
		}
	case "RotatedPlantCircumferenceShape":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.RotatedPlantCircumferenceShape = nil
			for __instance__ := range stage.PlantCircumferenceShapes {
				if stage.PlantCircumferenceShape_stagedOrder[__instance__] == uint(id) {
					plant.RotatedPlantCircumferenceShape = __instance__
					break
				}
			}
		}
	case "RotatedGridPathShape":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.RotatedGridPathShape = nil
			for __instance__ := range stage.GridPathShapes {
				if stage.GridPathShape_stagedOrder[__instance__] == uint(id) {
					plant.RotatedGridPathShape = __instance__
					break
				}
			}
		}
	case "RotatedRhombusGridShape2":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.RotatedRhombusGridShape2 = nil
			for __instance__ := range stage.RotatedRhombusGridShapes {
				if stage.RotatedRhombusGridShape_stagedOrder[__instance__] == uint(id) {
					plant.RotatedRhombusGridShape2 = __instance__
					break
				}
			}
		}
	case "GrowthCurveRhombusGridShape":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.GrowthCurveRhombusGridShape = nil
			for __instance__ := range stage.GrowthCurveRhombusGridShapes {
				if stage.GrowthCurveRhombusGridShape_stagedOrder[__instance__] == uint(id) {
					plant.GrowthCurveRhombusGridShape = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		plantcircumferenceshape.Name = value.GetValueString()
	case "AngleDegree":
		plantcircumferenceshape.AngleDegree = value.GetValueFloat()
	case "Length":
		plantcircumferenceshape.Length = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (plantdiagram *PlantDiagram) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		plantdiagram.Name = value.GetValueString()
	case "OriginX":
		plantdiagram.OriginX = value.GetValueFloat()
	case "OriginY":
		plantdiagram.OriginY = value.GetValueFloat()
	case "IsHiddenAxesShape":
		plantdiagram.IsHiddenAxesShape = value.GetValueBool()
	case "IsHiddenReferenceRhombus":
		plantdiagram.IsHiddenReferenceRhombus = value.GetValueBool()
	case "IsHiddenPlantCircumferenceShape":
		plantdiagram.IsHiddenPlantCircumferenceShape = value.GetValueBool()
	case "IsHiddenGridPathShape":
		plantdiagram.IsHiddenGridPathShape = value.GetValueBool()
	case "IsHiddenRhombusGridShape":
		plantdiagram.IsHiddenRhombusGridShape = value.GetValueBool()
	case "IsHiddenExplanationTextShape":
		plantdiagram.IsHiddenExplanationTextShape = value.GetValueBool()
	case "IsHiddenRotatedReferenceRhombus":
		plantdiagram.IsHiddenRotatedReferenceRhombus = value.GetValueBool()
	case "IsHiddenRotatedPlantCircumferenceShape":
		plantdiagram.IsHiddenRotatedPlantCircumferenceShape = value.GetValueBool()
	case "IsHiddenRotatedGridPathShape":
		plantdiagram.IsHiddenRotatedGridPathShape = value.GetValueBool()
	case "IsHiddenRotatedRhombusGridShape":
		plantdiagram.IsHiddenRotatedRhombusGridShape = value.GetValueBool()
	case "IsHiddenGrowthPathRhombusGridShape":
		plantdiagram.IsHiddenGrowthPathRhombusGridShape = value.GetValueBool()
	case "IsChecked":
		plantdiagram.IsChecked = value.GetValueBool()
	case "ComputedPrefix":
		plantdiagram.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		plantdiagram.IsExpanded = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (rhombusgridshape *RhombusGridShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		rhombusgridshape.Name = value.GetValueString()
	case "RhombusShapes":
		rhombusgridshape.RhombusShapes = make([]*RhombusShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.RhombusShapes {
					if stage.RhombusShape_stagedOrder[__instance__] == uint(id) {
						rhombusgridshape.RhombusShapes = append(rhombusgridshape.RhombusShapes, __instance__)
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

func (rhombusshape *RhombusShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		rhombusshape.Name = value.GetValueString()
	case "X":
		rhombusshape.X = value.GetValueFloat()
	case "Y":
		rhombusshape.Y = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		rotatedrhombusgridshape.Name = value.GetValueString()
	case "RotatedRhombusShapes":
		rotatedrhombusgridshape.RotatedRhombusShapes = make([]*RotatedRhombusShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.RotatedRhombusShapes {
					if stage.RotatedRhombusShape_stagedOrder[__instance__] == uint(id) {
						rotatedrhombusgridshape.RotatedRhombusShapes = append(rotatedrhombusgridshape.RotatedRhombusShapes, __instance__)
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

func (rotatedrhombusshape *RotatedRhombusShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		rotatedrhombusshape.Name = value.GetValueString()
	case "X":
		rotatedrhombusshape.X = value.GetValueFloat()
	case "Y":
		rotatedrhombusshape.Y = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (axesshape *AxesShape) GongGetGongstructName() string {
	return "AxesShape"
}

func (circlegridshape *CircleGridShape) GongGetGongstructName() string {
	return "CircleGridShape"
}

func (explanationtextshape *ExplanationTextShape) GongGetGongstructName() string {
	return "ExplanationTextShape"
}

func (gridpathshape *GridPathShape) GongGetGongstructName() string {
	return "GridPathShape"
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongGetGongstructName() string {
	return "GrowthCurveRhombusGridShape"
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongGetGongstructName() string {
	return "GrowthCurveRhombusShape"
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongGetGongstructName() string {
	return "InitialRhombusGridShape"
}

func (initialrhombusshape *InitialRhombusShape) GongGetGongstructName() string {
	return "InitialRhombusShape"
}

func (library *Library) GongGetGongstructName() string {
	return "Library"
}

func (nextcircleshape *NextCircleShape) GongGetGongstructName() string {
	return "NextCircleShape"
}

func (plant *Plant) GongGetGongstructName() string {
	return "Plant"
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongGetGongstructName() string {
	return "PlantCircumferenceShape"
}

func (plantdiagram *PlantDiagram) GongGetGongstructName() string {
	return "PlantDiagram"
}

func (rhombusgridshape *RhombusGridShape) GongGetGongstructName() string {
	return "RhombusGridShape"
}

func (rhombusshape *RhombusShape) GongGetGongstructName() string {
	return "RhombusShape"
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongGetGongstructName() string {
	return "RotatedRhombusGridShape"
}

func (rotatedrhombusshape *RotatedRhombusShape) GongGetGongstructName() string {
	return "RotatedRhombusShape"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	stage.AxesShapes_mapString = make(map[string]*AxesShape)
	for axesshape := range stage.AxesShapes {
		stage.AxesShapes_mapString[axesshape.Name] = axesshape
	}

	stage.CircleGridShapes_mapString = make(map[string]*CircleGridShape)
	for circlegridshape := range stage.CircleGridShapes {
		stage.CircleGridShapes_mapString[circlegridshape.Name] = circlegridshape
	}

	stage.ExplanationTextShapes_mapString = make(map[string]*ExplanationTextShape)
	for explanationtextshape := range stage.ExplanationTextShapes {
		stage.ExplanationTextShapes_mapString[explanationtextshape.Name] = explanationtextshape
	}

	stage.GridPathShapes_mapString = make(map[string]*GridPathShape)
	for gridpathshape := range stage.GridPathShapes {
		stage.GridPathShapes_mapString[gridpathshape.Name] = gridpathshape
	}

	stage.GrowthCurveRhombusGridShapes_mapString = make(map[string]*GrowthCurveRhombusGridShape)
	for growthcurverhombusgridshape := range stage.GrowthCurveRhombusGridShapes {
		stage.GrowthCurveRhombusGridShapes_mapString[growthcurverhombusgridshape.Name] = growthcurverhombusgridshape
	}

	stage.GrowthCurveRhombusShapes_mapString = make(map[string]*GrowthCurveRhombusShape)
	for growthcurverhombusshape := range stage.GrowthCurveRhombusShapes {
		stage.GrowthCurveRhombusShapes_mapString[growthcurverhombusshape.Name] = growthcurverhombusshape
	}

	stage.InitialRhombusGridShapes_mapString = make(map[string]*InitialRhombusGridShape)
	for initialrhombusgridshape := range stage.InitialRhombusGridShapes {
		stage.InitialRhombusGridShapes_mapString[initialrhombusgridshape.Name] = initialrhombusgridshape
	}

	stage.InitialRhombusShapes_mapString = make(map[string]*InitialRhombusShape)
	for initialrhombusshape := range stage.InitialRhombusShapes {
		stage.InitialRhombusShapes_mapString[initialrhombusshape.Name] = initialrhombusshape
	}

	stage.Librarys_mapString = make(map[string]*Library)
	for library := range stage.Librarys {
		stage.Librarys_mapString[library.Name] = library
	}

	stage.NextCircleShapes_mapString = make(map[string]*NextCircleShape)
	for nextcircleshape := range stage.NextCircleShapes {
		stage.NextCircleShapes_mapString[nextcircleshape.Name] = nextcircleshape
	}

	stage.Plants_mapString = make(map[string]*Plant)
	for plant := range stage.Plants {
		stage.Plants_mapString[plant.Name] = plant
	}

	stage.PlantCircumferenceShapes_mapString = make(map[string]*PlantCircumferenceShape)
	for plantcircumferenceshape := range stage.PlantCircumferenceShapes {
		stage.PlantCircumferenceShapes_mapString[plantcircumferenceshape.Name] = plantcircumferenceshape
	}

	stage.PlantDiagrams_mapString = make(map[string]*PlantDiagram)
	for plantdiagram := range stage.PlantDiagrams {
		stage.PlantDiagrams_mapString[plantdiagram.Name] = plantdiagram
	}

	stage.RhombusGridShapes_mapString = make(map[string]*RhombusGridShape)
	for rhombusgridshape := range stage.RhombusGridShapes {
		stage.RhombusGridShapes_mapString[rhombusgridshape.Name] = rhombusgridshape
	}

	stage.RhombusShapes_mapString = make(map[string]*RhombusShape)
	for rhombusshape := range stage.RhombusShapes {
		stage.RhombusShapes_mapString[rhombusshape.Name] = rhombusshape
	}

	stage.RotatedRhombusGridShapes_mapString = make(map[string]*RotatedRhombusGridShape)
	for rotatedrhombusgridshape := range stage.RotatedRhombusGridShapes {
		stage.RotatedRhombusGridShapes_mapString[rotatedrhombusgridshape.Name] = rotatedrhombusgridshape
	}

	stage.RotatedRhombusShapes_mapString = make(map[string]*RotatedRhombusShape)
	for rotatedrhombusshape := range stage.RotatedRhombusShapes {
		stage.RotatedRhombusShapes_mapString[rotatedrhombusshape.Name] = rotatedrhombusshape
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
