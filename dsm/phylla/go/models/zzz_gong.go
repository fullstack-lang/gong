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

	EndArcShapes                map[*EndArcShape]struct{}
	EndArcShapes_instance       map[*EndArcShape]*EndArcShape
	EndArcShapes_mapString      map[string]*EndArcShape
	EndArcShapeOrder            uint
	EndArcShape_stagedOrder     map[*EndArcShape]uint
	EndArcShape_orderStaged     map[uint]*EndArcShape
	EndArcShapes_reference      map[*EndArcShape]*EndArcShape
	EndArcShapes_referenceOrder map[*EndArcShape]uint

	// insertion point for slice of pointers maps
	OnAfterEndArcShapeCreateCallback OnAfterCreateInterface[EndArcShape]
	OnAfterEndArcShapeUpdateCallback OnAfterUpdateInterface[EndArcShape]
	OnAfterEndArcShapeDeleteCallback OnAfterDeleteInterface[EndArcShape]
	OnAfterEndArcShapeReadCallback   OnAfterReadInterface[EndArcShape]

	EndArcShapeGrids                map[*EndArcShapeGrid]struct{}
	EndArcShapeGrids_instance       map[*EndArcShapeGrid]*EndArcShapeGrid
	EndArcShapeGrids_mapString      map[string]*EndArcShapeGrid
	EndArcShapeGridOrder            uint
	EndArcShapeGrid_stagedOrder     map[*EndArcShapeGrid]uint
	EndArcShapeGrid_orderStaged     map[uint]*EndArcShapeGrid
	EndArcShapeGrids_reference      map[*EndArcShapeGrid]*EndArcShapeGrid
	EndArcShapeGrids_referenceOrder map[*EndArcShapeGrid]uint

	// insertion point for slice of pointers maps
	EndArcShapeGrid_EndArcShapes_reverseMap map[*EndArcShape]*EndArcShapeGrid

	OnAfterEndArcShapeGridCreateCallback OnAfterCreateInterface[EndArcShapeGrid]
	OnAfterEndArcShapeGridUpdateCallback OnAfterUpdateInterface[EndArcShapeGrid]
	OnAfterEndArcShapeGridDeleteCallback OnAfterDeleteInterface[EndArcShapeGrid]
	OnAfterEndArcShapeGridReadCallback   OnAfterReadInterface[EndArcShapeGrid]

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

	GrowthCurveBezierShapes                map[*GrowthCurveBezierShape]struct{}
	GrowthCurveBezierShapes_instance       map[*GrowthCurveBezierShape]*GrowthCurveBezierShape
	GrowthCurveBezierShapes_mapString      map[string]*GrowthCurveBezierShape
	GrowthCurveBezierShapeOrder            uint
	GrowthCurveBezierShape_stagedOrder     map[*GrowthCurveBezierShape]uint
	GrowthCurveBezierShape_orderStaged     map[uint]*GrowthCurveBezierShape
	GrowthCurveBezierShapes_reference      map[*GrowthCurveBezierShape]*GrowthCurveBezierShape
	GrowthCurveBezierShapes_referenceOrder map[*GrowthCurveBezierShape]uint

	// insertion point for slice of pointers maps
	OnAfterGrowthCurveBezierShapeCreateCallback OnAfterCreateInterface[GrowthCurveBezierShape]
	OnAfterGrowthCurveBezierShapeUpdateCallback OnAfterUpdateInterface[GrowthCurveBezierShape]
	OnAfterGrowthCurveBezierShapeDeleteCallback OnAfterDeleteInterface[GrowthCurveBezierShape]
	OnAfterGrowthCurveBezierShapeReadCallback   OnAfterReadInterface[GrowthCurveBezierShape]

	GrowthCurveBezierShapeGrids                map[*GrowthCurveBezierShapeGrid]struct{}
	GrowthCurveBezierShapeGrids_instance       map[*GrowthCurveBezierShapeGrid]*GrowthCurveBezierShapeGrid
	GrowthCurveBezierShapeGrids_mapString      map[string]*GrowthCurveBezierShapeGrid
	GrowthCurveBezierShapeGridOrder            uint
	GrowthCurveBezierShapeGrid_stagedOrder     map[*GrowthCurveBezierShapeGrid]uint
	GrowthCurveBezierShapeGrid_orderStaged     map[uint]*GrowthCurveBezierShapeGrid
	GrowthCurveBezierShapeGrids_reference      map[*GrowthCurveBezierShapeGrid]*GrowthCurveBezierShapeGrid
	GrowthCurveBezierShapeGrids_referenceOrder map[*GrowthCurveBezierShapeGrid]uint

	// insertion point for slice of pointers maps
	GrowthCurveBezierShapeGrid_GrowthCurveBezierShapes_reverseMap map[*GrowthCurveBezierShape]*GrowthCurveBezierShapeGrid

	OnAfterGrowthCurveBezierShapeGridCreateCallback OnAfterCreateInterface[GrowthCurveBezierShapeGrid]
	OnAfterGrowthCurveBezierShapeGridUpdateCallback OnAfterUpdateInterface[GrowthCurveBezierShapeGrid]
	OnAfterGrowthCurveBezierShapeGridDeleteCallback OnAfterDeleteInterface[GrowthCurveBezierShapeGrid]
	OnAfterGrowthCurveBezierShapeGridReadCallback   OnAfterReadInterface[GrowthCurveBezierShapeGrid]

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

	GrowthVectorShapes                map[*GrowthVectorShape]struct{}
	GrowthVectorShapes_instance       map[*GrowthVectorShape]*GrowthVectorShape
	GrowthVectorShapes_mapString      map[string]*GrowthVectorShape
	GrowthVectorShapeOrder            uint
	GrowthVectorShape_stagedOrder     map[*GrowthVectorShape]uint
	GrowthVectorShape_orderStaged     map[uint]*GrowthVectorShape
	GrowthVectorShapes_reference      map[*GrowthVectorShape]*GrowthVectorShape
	GrowthVectorShapes_referenceOrder map[*GrowthVectorShape]uint

	// insertion point for slice of pointers maps
	OnAfterGrowthVectorShapeCreateCallback OnAfterCreateInterface[GrowthVectorShape]
	OnAfterGrowthVectorShapeUpdateCallback OnAfterUpdateInterface[GrowthVectorShape]
	OnAfterGrowthVectorShapeDeleteCallback OnAfterDeleteInterface[GrowthVectorShape]
	OnAfterGrowthVectorShapeReadCallback   OnAfterReadInterface[GrowthVectorShape]

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

	PerpendicularVectors                map[*PerpendicularVector]struct{}
	PerpendicularVectors_instance       map[*PerpendicularVector]*PerpendicularVector
	PerpendicularVectors_mapString      map[string]*PerpendicularVector
	PerpendicularVectorOrder            uint
	PerpendicularVector_stagedOrder     map[*PerpendicularVector]uint
	PerpendicularVector_orderStaged     map[uint]*PerpendicularVector
	PerpendicularVectors_reference      map[*PerpendicularVector]*PerpendicularVector
	PerpendicularVectors_referenceOrder map[*PerpendicularVector]uint

	// insertion point for slice of pointers maps
	OnAfterPerpendicularVectorCreateCallback OnAfterCreateInterface[PerpendicularVector]
	OnAfterPerpendicularVectorUpdateCallback OnAfterUpdateInterface[PerpendicularVector]
	OnAfterPerpendicularVectorDeleteCallback OnAfterDeleteInterface[PerpendicularVector]
	OnAfterPerpendicularVectorReadCallback   OnAfterReadInterface[PerpendicularVector]

	PerpendicularVectorGrids                map[*PerpendicularVectorGrid]struct{}
	PerpendicularVectorGrids_instance       map[*PerpendicularVectorGrid]*PerpendicularVectorGrid
	PerpendicularVectorGrids_mapString      map[string]*PerpendicularVectorGrid
	PerpendicularVectorGridOrder            uint
	PerpendicularVectorGrid_stagedOrder     map[*PerpendicularVectorGrid]uint
	PerpendicularVectorGrid_orderStaged     map[uint]*PerpendicularVectorGrid
	PerpendicularVectorGrids_reference      map[*PerpendicularVectorGrid]*PerpendicularVectorGrid
	PerpendicularVectorGrids_referenceOrder map[*PerpendicularVectorGrid]uint

	// insertion point for slice of pointers maps
	PerpendicularVectorGrid_PerpendicularVectors_reverseMap map[*PerpendicularVector]*PerpendicularVectorGrid

	OnAfterPerpendicularVectorGridCreateCallback OnAfterCreateInterface[PerpendicularVectorGrid]
	OnAfterPerpendicularVectorGridUpdateCallback OnAfterUpdateInterface[PerpendicularVectorGrid]
	OnAfterPerpendicularVectorGridDeleteCallback OnAfterDeleteInterface[PerpendicularVectorGrid]
	OnAfterPerpendicularVectorGridReadCallback   OnAfterReadInterface[PerpendicularVectorGrid]

	PerpendicularVectorGridHalfways                map[*PerpendicularVectorGridHalfway]struct{}
	PerpendicularVectorGridHalfways_instance       map[*PerpendicularVectorGridHalfway]*PerpendicularVectorGridHalfway
	PerpendicularVectorGridHalfways_mapString      map[string]*PerpendicularVectorGridHalfway
	PerpendicularVectorGridHalfwayOrder            uint
	PerpendicularVectorGridHalfway_stagedOrder     map[*PerpendicularVectorGridHalfway]uint
	PerpendicularVectorGridHalfway_orderStaged     map[uint]*PerpendicularVectorGridHalfway
	PerpendicularVectorGridHalfways_reference      map[*PerpendicularVectorGridHalfway]*PerpendicularVectorGridHalfway
	PerpendicularVectorGridHalfways_referenceOrder map[*PerpendicularVectorGridHalfway]uint

	// insertion point for slice of pointers maps
	PerpendicularVectorGridHalfway_PerpendicularVectorHalfways_reverseMap map[*PerpendicularVectorHalfway]*PerpendicularVectorGridHalfway

	OnAfterPerpendicularVectorGridHalfwayCreateCallback OnAfterCreateInterface[PerpendicularVectorGridHalfway]
	OnAfterPerpendicularVectorGridHalfwayUpdateCallback OnAfterUpdateInterface[PerpendicularVectorGridHalfway]
	OnAfterPerpendicularVectorGridHalfwayDeleteCallback OnAfterDeleteInterface[PerpendicularVectorGridHalfway]
	OnAfterPerpendicularVectorGridHalfwayReadCallback   OnAfterReadInterface[PerpendicularVectorGridHalfway]

	PerpendicularVectorHalfways                map[*PerpendicularVectorHalfway]struct{}
	PerpendicularVectorHalfways_instance       map[*PerpendicularVectorHalfway]*PerpendicularVectorHalfway
	PerpendicularVectorHalfways_mapString      map[string]*PerpendicularVectorHalfway
	PerpendicularVectorHalfwayOrder            uint
	PerpendicularVectorHalfway_stagedOrder     map[*PerpendicularVectorHalfway]uint
	PerpendicularVectorHalfway_orderStaged     map[uint]*PerpendicularVectorHalfway
	PerpendicularVectorHalfways_reference      map[*PerpendicularVectorHalfway]*PerpendicularVectorHalfway
	PerpendicularVectorHalfways_referenceOrder map[*PerpendicularVectorHalfway]uint

	// insertion point for slice of pointers maps
	OnAfterPerpendicularVectorHalfwayCreateCallback OnAfterCreateInterface[PerpendicularVectorHalfway]
	OnAfterPerpendicularVectorHalfwayUpdateCallback OnAfterUpdateInterface[PerpendicularVectorHalfway]
	OnAfterPerpendicularVectorHalfwayDeleteCallback OnAfterDeleteInterface[PerpendicularVectorHalfway]
	OnAfterPerpendicularVectorHalfwayReadCallback   OnAfterReadInterface[PerpendicularVectorHalfway]

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

	Rendered3DShapes                map[*Rendered3DShape]struct{}
	Rendered3DShapes_instance       map[*Rendered3DShape]*Rendered3DShape
	Rendered3DShapes_mapString      map[string]*Rendered3DShape
	Rendered3DShapeOrder            uint
	Rendered3DShape_stagedOrder     map[*Rendered3DShape]uint
	Rendered3DShape_orderStaged     map[uint]*Rendered3DShape
	Rendered3DShapes_reference      map[*Rendered3DShape]*Rendered3DShape
	Rendered3DShapes_referenceOrder map[*Rendered3DShape]uint

	// insertion point for slice of pointers maps
	OnAfterRendered3DShapeCreateCallback OnAfterCreateInterface[Rendered3DShape]
	OnAfterRendered3DShapeUpdateCallback OnAfterUpdateInterface[Rendered3DShape]
	OnAfterRendered3DShapeDeleteCallback OnAfterDeleteInterface[Rendered3DShape]
	OnAfterRendered3DShapeReadCallback   OnAfterReadInterface[Rendered3DShape]

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

	StackGrowthCurveBezierShapes                map[*StackGrowthCurveBezierShape]struct{}
	StackGrowthCurveBezierShapes_instance       map[*StackGrowthCurveBezierShape]*StackGrowthCurveBezierShape
	StackGrowthCurveBezierShapes_mapString      map[string]*StackGrowthCurveBezierShape
	StackGrowthCurveBezierShapeOrder            uint
	StackGrowthCurveBezierShape_stagedOrder     map[*StackGrowthCurveBezierShape]uint
	StackGrowthCurveBezierShape_orderStaged     map[uint]*StackGrowthCurveBezierShape
	StackGrowthCurveBezierShapes_reference      map[*StackGrowthCurveBezierShape]*StackGrowthCurveBezierShape
	StackGrowthCurveBezierShapes_referenceOrder map[*StackGrowthCurveBezierShape]uint

	// insertion point for slice of pointers maps
	OnAfterStackGrowthCurveBezierShapeCreateCallback OnAfterCreateInterface[StackGrowthCurveBezierShape]
	OnAfterStackGrowthCurveBezierShapeUpdateCallback OnAfterUpdateInterface[StackGrowthCurveBezierShape]
	OnAfterStackGrowthCurveBezierShapeDeleteCallback OnAfterDeleteInterface[StackGrowthCurveBezierShape]
	OnAfterStackGrowthCurveBezierShapeReadCallback   OnAfterReadInterface[StackGrowthCurveBezierShape]

	StackOfGrowthCurves                map[*StackOfGrowthCurve]struct{}
	StackOfGrowthCurves_instance       map[*StackOfGrowthCurve]*StackOfGrowthCurve
	StackOfGrowthCurves_mapString      map[string]*StackOfGrowthCurve
	StackOfGrowthCurveOrder            uint
	StackOfGrowthCurve_stagedOrder     map[*StackOfGrowthCurve]uint
	StackOfGrowthCurve_orderStaged     map[uint]*StackOfGrowthCurve
	StackOfGrowthCurves_reference      map[*StackOfGrowthCurve]*StackOfGrowthCurve
	StackOfGrowthCurves_referenceOrder map[*StackOfGrowthCurve]uint

	// insertion point for slice of pointers maps
	StackOfGrowthCurve_StackGrowthCurveBezierShapes_reverseMap map[*StackGrowthCurveBezierShape]*StackOfGrowthCurve

	OnAfterStackOfGrowthCurveCreateCallback OnAfterCreateInterface[StackOfGrowthCurve]
	OnAfterStackOfGrowthCurveUpdateCallback OnAfterUpdateInterface[StackOfGrowthCurve]
	OnAfterStackOfGrowthCurveDeleteCallback OnAfterDeleteInterface[StackOfGrowthCurve]
	OnAfterStackOfGrowthCurveReadCallback   OnAfterReadInterface[StackOfGrowthCurve]

	StartArcShapes                map[*StartArcShape]struct{}
	StartArcShapes_instance       map[*StartArcShape]*StartArcShape
	StartArcShapes_mapString      map[string]*StartArcShape
	StartArcShapeOrder            uint
	StartArcShape_stagedOrder     map[*StartArcShape]uint
	StartArcShape_orderStaged     map[uint]*StartArcShape
	StartArcShapes_reference      map[*StartArcShape]*StartArcShape
	StartArcShapes_referenceOrder map[*StartArcShape]uint

	// insertion point for slice of pointers maps
	OnAfterStartArcShapeCreateCallback OnAfterCreateInterface[StartArcShape]
	OnAfterStartArcShapeUpdateCallback OnAfterUpdateInterface[StartArcShape]
	OnAfterStartArcShapeDeleteCallback OnAfterDeleteInterface[StartArcShape]
	OnAfterStartArcShapeReadCallback   OnAfterReadInterface[StartArcShape]

	StartArcShapeGrids                map[*StartArcShapeGrid]struct{}
	StartArcShapeGrids_instance       map[*StartArcShapeGrid]*StartArcShapeGrid
	StartArcShapeGrids_mapString      map[string]*StartArcShapeGrid
	StartArcShapeGridOrder            uint
	StartArcShapeGrid_stagedOrder     map[*StartArcShapeGrid]uint
	StartArcShapeGrid_orderStaged     map[uint]*StartArcShapeGrid
	StartArcShapeGrids_reference      map[*StartArcShapeGrid]*StartArcShapeGrid
	StartArcShapeGrids_referenceOrder map[*StartArcShapeGrid]uint

	// insertion point for slice of pointers maps
	StartArcShapeGrid_StartArcShapes_reverseMap map[*StartArcShape]*StartArcShapeGrid

	OnAfterStartArcShapeGridCreateCallback OnAfterCreateInterface[StartArcShapeGrid]
	OnAfterStartArcShapeGridUpdateCallback OnAfterUpdateInterface[StartArcShapeGrid]
	OnAfterStartArcShapeGridDeleteCallback OnAfterDeleteInterface[StartArcShapeGrid]
	OnAfterStartArcShapeGridReadCallback   OnAfterReadInterface[StartArcShapeGrid]

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

	stage.EndArcShapes_reference = make(map[*EndArcShape]*EndArcShape)
	stage.EndArcShapes_instance = make(map[*EndArcShape]*EndArcShape)
	stage.EndArcShapes_referenceOrder = make(map[*EndArcShape]uint)

	stage.EndArcShapeGrids_reference = make(map[*EndArcShapeGrid]*EndArcShapeGrid)
	stage.EndArcShapeGrids_instance = make(map[*EndArcShapeGrid]*EndArcShapeGrid)
	stage.EndArcShapeGrids_referenceOrder = make(map[*EndArcShapeGrid]uint)

	stage.ExplanationTextShapes_reference = make(map[*ExplanationTextShape]*ExplanationTextShape)
	stage.ExplanationTextShapes_instance = make(map[*ExplanationTextShape]*ExplanationTextShape)
	stage.ExplanationTextShapes_referenceOrder = make(map[*ExplanationTextShape]uint)

	stage.GridPathShapes_reference = make(map[*GridPathShape]*GridPathShape)
	stage.GridPathShapes_instance = make(map[*GridPathShape]*GridPathShape)
	stage.GridPathShapes_referenceOrder = make(map[*GridPathShape]uint)

	stage.GrowthCurveBezierShapes_reference = make(map[*GrowthCurveBezierShape]*GrowthCurveBezierShape)
	stage.GrowthCurveBezierShapes_instance = make(map[*GrowthCurveBezierShape]*GrowthCurveBezierShape)
	stage.GrowthCurveBezierShapes_referenceOrder = make(map[*GrowthCurveBezierShape]uint)

	stage.GrowthCurveBezierShapeGrids_reference = make(map[*GrowthCurveBezierShapeGrid]*GrowthCurveBezierShapeGrid)
	stage.GrowthCurveBezierShapeGrids_instance = make(map[*GrowthCurveBezierShapeGrid]*GrowthCurveBezierShapeGrid)
	stage.GrowthCurveBezierShapeGrids_referenceOrder = make(map[*GrowthCurveBezierShapeGrid]uint)

	stage.GrowthCurveRhombusGridShapes_reference = make(map[*GrowthCurveRhombusGridShape]*GrowthCurveRhombusGridShape)
	stage.GrowthCurveRhombusGridShapes_instance = make(map[*GrowthCurveRhombusGridShape]*GrowthCurveRhombusGridShape)
	stage.GrowthCurveRhombusGridShapes_referenceOrder = make(map[*GrowthCurveRhombusGridShape]uint)

	stage.GrowthCurveRhombusShapes_reference = make(map[*GrowthCurveRhombusShape]*GrowthCurveRhombusShape)
	stage.GrowthCurveRhombusShapes_instance = make(map[*GrowthCurveRhombusShape]*GrowthCurveRhombusShape)
	stage.GrowthCurveRhombusShapes_referenceOrder = make(map[*GrowthCurveRhombusShape]uint)

	stage.GrowthVectorShapes_reference = make(map[*GrowthVectorShape]*GrowthVectorShape)
	stage.GrowthVectorShapes_instance = make(map[*GrowthVectorShape]*GrowthVectorShape)
	stage.GrowthVectorShapes_referenceOrder = make(map[*GrowthVectorShape]uint)

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

	stage.PerpendicularVectors_reference = make(map[*PerpendicularVector]*PerpendicularVector)
	stage.PerpendicularVectors_instance = make(map[*PerpendicularVector]*PerpendicularVector)
	stage.PerpendicularVectors_referenceOrder = make(map[*PerpendicularVector]uint)

	stage.PerpendicularVectorGrids_reference = make(map[*PerpendicularVectorGrid]*PerpendicularVectorGrid)
	stage.PerpendicularVectorGrids_instance = make(map[*PerpendicularVectorGrid]*PerpendicularVectorGrid)
	stage.PerpendicularVectorGrids_referenceOrder = make(map[*PerpendicularVectorGrid]uint)

	stage.PerpendicularVectorGridHalfways_reference = make(map[*PerpendicularVectorGridHalfway]*PerpendicularVectorGridHalfway)
	stage.PerpendicularVectorGridHalfways_instance = make(map[*PerpendicularVectorGridHalfway]*PerpendicularVectorGridHalfway)
	stage.PerpendicularVectorGridHalfways_referenceOrder = make(map[*PerpendicularVectorGridHalfway]uint)

	stage.PerpendicularVectorHalfways_reference = make(map[*PerpendicularVectorHalfway]*PerpendicularVectorHalfway)
	stage.PerpendicularVectorHalfways_instance = make(map[*PerpendicularVectorHalfway]*PerpendicularVectorHalfway)
	stage.PerpendicularVectorHalfways_referenceOrder = make(map[*PerpendicularVectorHalfway]uint)

	stage.Plants_reference = make(map[*Plant]*Plant)
	stage.Plants_instance = make(map[*Plant]*Plant)
	stage.Plants_referenceOrder = make(map[*Plant]uint)

	stage.PlantCircumferenceShapes_reference = make(map[*PlantCircumferenceShape]*PlantCircumferenceShape)
	stage.PlantCircumferenceShapes_instance = make(map[*PlantCircumferenceShape]*PlantCircumferenceShape)
	stage.PlantCircumferenceShapes_referenceOrder = make(map[*PlantCircumferenceShape]uint)

	stage.PlantDiagrams_reference = make(map[*PlantDiagram]*PlantDiagram)
	stage.PlantDiagrams_instance = make(map[*PlantDiagram]*PlantDiagram)
	stage.PlantDiagrams_referenceOrder = make(map[*PlantDiagram]uint)

	stage.Rendered3DShapes_reference = make(map[*Rendered3DShape]*Rendered3DShape)
	stage.Rendered3DShapes_instance = make(map[*Rendered3DShape]*Rendered3DShape)
	stage.Rendered3DShapes_referenceOrder = make(map[*Rendered3DShape]uint)

	stage.RhombusShapes_reference = make(map[*RhombusShape]*RhombusShape)
	stage.RhombusShapes_instance = make(map[*RhombusShape]*RhombusShape)
	stage.RhombusShapes_referenceOrder = make(map[*RhombusShape]uint)

	stage.RotatedRhombusGridShapes_reference = make(map[*RotatedRhombusGridShape]*RotatedRhombusGridShape)
	stage.RotatedRhombusGridShapes_instance = make(map[*RotatedRhombusGridShape]*RotatedRhombusGridShape)
	stage.RotatedRhombusGridShapes_referenceOrder = make(map[*RotatedRhombusGridShape]uint)

	stage.RotatedRhombusShapes_reference = make(map[*RotatedRhombusShape]*RotatedRhombusShape)
	stage.RotatedRhombusShapes_instance = make(map[*RotatedRhombusShape]*RotatedRhombusShape)
	stage.RotatedRhombusShapes_referenceOrder = make(map[*RotatedRhombusShape]uint)

	stage.StackGrowthCurveBezierShapes_reference = make(map[*StackGrowthCurveBezierShape]*StackGrowthCurveBezierShape)
	stage.StackGrowthCurveBezierShapes_instance = make(map[*StackGrowthCurveBezierShape]*StackGrowthCurveBezierShape)
	stage.StackGrowthCurveBezierShapes_referenceOrder = make(map[*StackGrowthCurveBezierShape]uint)

	stage.StackOfGrowthCurves_reference = make(map[*StackOfGrowthCurve]*StackOfGrowthCurve)
	stage.StackOfGrowthCurves_instance = make(map[*StackOfGrowthCurve]*StackOfGrowthCurve)
	stage.StackOfGrowthCurves_referenceOrder = make(map[*StackOfGrowthCurve]uint)

	stage.StartArcShapes_reference = make(map[*StartArcShape]*StartArcShape)
	stage.StartArcShapes_instance = make(map[*StartArcShape]*StartArcShape)
	stage.StartArcShapes_referenceOrder = make(map[*StartArcShape]uint)

	stage.StartArcShapeGrids_reference = make(map[*StartArcShapeGrid]*StartArcShapeGrid)
	stage.StartArcShapeGrids_instance = make(map[*StartArcShapeGrid]*StartArcShapeGrid)
	stage.StartArcShapeGrids_referenceOrder = make(map[*StartArcShapeGrid]uint)

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

	var maxEndArcShapeOrder uint
	var foundEndArcShape bool
	for _, order := range stage.EndArcShape_stagedOrder {
		if !foundEndArcShape || order > maxEndArcShapeOrder {
			maxEndArcShapeOrder = order
			foundEndArcShape = true
		}
	}
	if foundEndArcShape {
		stage.EndArcShapeOrder = maxEndArcShapeOrder + 1
	} else {
		stage.EndArcShapeOrder = 0
	}

	var maxEndArcShapeGridOrder uint
	var foundEndArcShapeGrid bool
	for _, order := range stage.EndArcShapeGrid_stagedOrder {
		if !foundEndArcShapeGrid || order > maxEndArcShapeGridOrder {
			maxEndArcShapeGridOrder = order
			foundEndArcShapeGrid = true
		}
	}
	if foundEndArcShapeGrid {
		stage.EndArcShapeGridOrder = maxEndArcShapeGridOrder + 1
	} else {
		stage.EndArcShapeGridOrder = 0
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

	var maxGrowthCurveBezierShapeOrder uint
	var foundGrowthCurveBezierShape bool
	for _, order := range stage.GrowthCurveBezierShape_stagedOrder {
		if !foundGrowthCurveBezierShape || order > maxGrowthCurveBezierShapeOrder {
			maxGrowthCurveBezierShapeOrder = order
			foundGrowthCurveBezierShape = true
		}
	}
	if foundGrowthCurveBezierShape {
		stage.GrowthCurveBezierShapeOrder = maxGrowthCurveBezierShapeOrder + 1
	} else {
		stage.GrowthCurveBezierShapeOrder = 0
	}

	var maxGrowthCurveBezierShapeGridOrder uint
	var foundGrowthCurveBezierShapeGrid bool
	for _, order := range stage.GrowthCurveBezierShapeGrid_stagedOrder {
		if !foundGrowthCurveBezierShapeGrid || order > maxGrowthCurveBezierShapeGridOrder {
			maxGrowthCurveBezierShapeGridOrder = order
			foundGrowthCurveBezierShapeGrid = true
		}
	}
	if foundGrowthCurveBezierShapeGrid {
		stage.GrowthCurveBezierShapeGridOrder = maxGrowthCurveBezierShapeGridOrder + 1
	} else {
		stage.GrowthCurveBezierShapeGridOrder = 0
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

	var maxGrowthVectorShapeOrder uint
	var foundGrowthVectorShape bool
	for _, order := range stage.GrowthVectorShape_stagedOrder {
		if !foundGrowthVectorShape || order > maxGrowthVectorShapeOrder {
			maxGrowthVectorShapeOrder = order
			foundGrowthVectorShape = true
		}
	}
	if foundGrowthVectorShape {
		stage.GrowthVectorShapeOrder = maxGrowthVectorShapeOrder + 1
	} else {
		stage.GrowthVectorShapeOrder = 0
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

	var maxPerpendicularVectorOrder uint
	var foundPerpendicularVector bool
	for _, order := range stage.PerpendicularVector_stagedOrder {
		if !foundPerpendicularVector || order > maxPerpendicularVectorOrder {
			maxPerpendicularVectorOrder = order
			foundPerpendicularVector = true
		}
	}
	if foundPerpendicularVector {
		stage.PerpendicularVectorOrder = maxPerpendicularVectorOrder + 1
	} else {
		stage.PerpendicularVectorOrder = 0
	}

	var maxPerpendicularVectorGridOrder uint
	var foundPerpendicularVectorGrid bool
	for _, order := range stage.PerpendicularVectorGrid_stagedOrder {
		if !foundPerpendicularVectorGrid || order > maxPerpendicularVectorGridOrder {
			maxPerpendicularVectorGridOrder = order
			foundPerpendicularVectorGrid = true
		}
	}
	if foundPerpendicularVectorGrid {
		stage.PerpendicularVectorGridOrder = maxPerpendicularVectorGridOrder + 1
	} else {
		stage.PerpendicularVectorGridOrder = 0
	}

	var maxPerpendicularVectorGridHalfwayOrder uint
	var foundPerpendicularVectorGridHalfway bool
	for _, order := range stage.PerpendicularVectorGridHalfway_stagedOrder {
		if !foundPerpendicularVectorGridHalfway || order > maxPerpendicularVectorGridHalfwayOrder {
			maxPerpendicularVectorGridHalfwayOrder = order
			foundPerpendicularVectorGridHalfway = true
		}
	}
	if foundPerpendicularVectorGridHalfway {
		stage.PerpendicularVectorGridHalfwayOrder = maxPerpendicularVectorGridHalfwayOrder + 1
	} else {
		stage.PerpendicularVectorGridHalfwayOrder = 0
	}

	var maxPerpendicularVectorHalfwayOrder uint
	var foundPerpendicularVectorHalfway bool
	for _, order := range stage.PerpendicularVectorHalfway_stagedOrder {
		if !foundPerpendicularVectorHalfway || order > maxPerpendicularVectorHalfwayOrder {
			maxPerpendicularVectorHalfwayOrder = order
			foundPerpendicularVectorHalfway = true
		}
	}
	if foundPerpendicularVectorHalfway {
		stage.PerpendicularVectorHalfwayOrder = maxPerpendicularVectorHalfwayOrder + 1
	} else {
		stage.PerpendicularVectorHalfwayOrder = 0
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

	var maxRendered3DShapeOrder uint
	var foundRendered3DShape bool
	for _, order := range stage.Rendered3DShape_stagedOrder {
		if !foundRendered3DShape || order > maxRendered3DShapeOrder {
			maxRendered3DShapeOrder = order
			foundRendered3DShape = true
		}
	}
	if foundRendered3DShape {
		stage.Rendered3DShapeOrder = maxRendered3DShapeOrder + 1
	} else {
		stage.Rendered3DShapeOrder = 0
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

	var maxStackGrowthCurveBezierShapeOrder uint
	var foundStackGrowthCurveBezierShape bool
	for _, order := range stage.StackGrowthCurveBezierShape_stagedOrder {
		if !foundStackGrowthCurveBezierShape || order > maxStackGrowthCurveBezierShapeOrder {
			maxStackGrowthCurveBezierShapeOrder = order
			foundStackGrowthCurveBezierShape = true
		}
	}
	if foundStackGrowthCurveBezierShape {
		stage.StackGrowthCurveBezierShapeOrder = maxStackGrowthCurveBezierShapeOrder + 1
	} else {
		stage.StackGrowthCurveBezierShapeOrder = 0
	}

	var maxStackOfGrowthCurveOrder uint
	var foundStackOfGrowthCurve bool
	for _, order := range stage.StackOfGrowthCurve_stagedOrder {
		if !foundStackOfGrowthCurve || order > maxStackOfGrowthCurveOrder {
			maxStackOfGrowthCurveOrder = order
			foundStackOfGrowthCurve = true
		}
	}
	if foundStackOfGrowthCurve {
		stage.StackOfGrowthCurveOrder = maxStackOfGrowthCurveOrder + 1
	} else {
		stage.StackOfGrowthCurveOrder = 0
	}

	var maxStartArcShapeOrder uint
	var foundStartArcShape bool
	for _, order := range stage.StartArcShape_stagedOrder {
		if !foundStartArcShape || order > maxStartArcShapeOrder {
			maxStartArcShapeOrder = order
			foundStartArcShape = true
		}
	}
	if foundStartArcShape {
		stage.StartArcShapeOrder = maxStartArcShapeOrder + 1
	} else {
		stage.StartArcShapeOrder = 0
	}

	var maxStartArcShapeGridOrder uint
	var foundStartArcShapeGrid bool
	for _, order := range stage.StartArcShapeGrid_stagedOrder {
		if !foundStartArcShapeGrid || order > maxStartArcShapeGridOrder {
			maxStartArcShapeGridOrder = order
			foundStartArcShapeGrid = true
		}
	}
	if foundStartArcShapeGrid {
		stage.StartArcShapeGridOrder = maxStartArcShapeGridOrder + 1
	} else {
		stage.StartArcShapeGridOrder = 0
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
	case *EndArcShape:
		tmp := GetStructInstancesByOrder(stage.EndArcShapes, stage.EndArcShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *EndArcShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *EndArcShapeGrid:
		tmp := GetStructInstancesByOrder(stage.EndArcShapeGrids, stage.EndArcShapeGrid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *EndArcShapeGrid implements.
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
	case *GrowthCurveBezierShape:
		tmp := GetStructInstancesByOrder(stage.GrowthCurveBezierShapes, stage.GrowthCurveBezierShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GrowthCurveBezierShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *GrowthCurveBezierShapeGrid:
		tmp := GetStructInstancesByOrder(stage.GrowthCurveBezierShapeGrids, stage.GrowthCurveBezierShapeGrid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GrowthCurveBezierShapeGrid implements.
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
	case *GrowthVectorShape:
		tmp := GetStructInstancesByOrder(stage.GrowthVectorShapes, stage.GrowthVectorShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GrowthVectorShape implements.
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
	case *PerpendicularVector:
		tmp := GetStructInstancesByOrder(stage.PerpendicularVectors, stage.PerpendicularVector_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *PerpendicularVector implements.
			res = append(res, any(v).(T))
		}
		return res
	case *PerpendicularVectorGrid:
		tmp := GetStructInstancesByOrder(stage.PerpendicularVectorGrids, stage.PerpendicularVectorGrid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *PerpendicularVectorGrid implements.
			res = append(res, any(v).(T))
		}
		return res
	case *PerpendicularVectorGridHalfway:
		tmp := GetStructInstancesByOrder(stage.PerpendicularVectorGridHalfways, stage.PerpendicularVectorGridHalfway_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *PerpendicularVectorGridHalfway implements.
			res = append(res, any(v).(T))
		}
		return res
	case *PerpendicularVectorHalfway:
		tmp := GetStructInstancesByOrder(stage.PerpendicularVectorHalfways, stage.PerpendicularVectorHalfway_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *PerpendicularVectorHalfway implements.
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
	case *Rendered3DShape:
		tmp := GetStructInstancesByOrder(stage.Rendered3DShapes, stage.Rendered3DShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Rendered3DShape implements.
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
	case *StackGrowthCurveBezierShape:
		tmp := GetStructInstancesByOrder(stage.StackGrowthCurveBezierShapes, stage.StackGrowthCurveBezierShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *StackGrowthCurveBezierShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *StackOfGrowthCurve:
		tmp := GetStructInstancesByOrder(stage.StackOfGrowthCurves, stage.StackOfGrowthCurve_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *StackOfGrowthCurve implements.
			res = append(res, any(v).(T))
		}
		return res
	case *StartArcShape:
		tmp := GetStructInstancesByOrder(stage.StartArcShapes, stage.StartArcShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *StartArcShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *StartArcShapeGrid:
		tmp := GetStructInstancesByOrder(stage.StartArcShapeGrids, stage.StartArcShapeGrid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *StartArcShapeGrid implements.
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
	case "EndArcShape":
		res = GetNamedStructInstances(stage.EndArcShapes, stage.EndArcShape_stagedOrder)
	case "EndArcShapeGrid":
		res = GetNamedStructInstances(stage.EndArcShapeGrids, stage.EndArcShapeGrid_stagedOrder)
	case "ExplanationTextShape":
		res = GetNamedStructInstances(stage.ExplanationTextShapes, stage.ExplanationTextShape_stagedOrder)
	case "GridPathShape":
		res = GetNamedStructInstances(stage.GridPathShapes, stage.GridPathShape_stagedOrder)
	case "GrowthCurveBezierShape":
		res = GetNamedStructInstances(stage.GrowthCurveBezierShapes, stage.GrowthCurveBezierShape_stagedOrder)
	case "GrowthCurveBezierShapeGrid":
		res = GetNamedStructInstances(stage.GrowthCurveBezierShapeGrids, stage.GrowthCurveBezierShapeGrid_stagedOrder)
	case "GrowthCurveRhombusGridShape":
		res = GetNamedStructInstances(stage.GrowthCurveRhombusGridShapes, stage.GrowthCurveRhombusGridShape_stagedOrder)
	case "GrowthCurveRhombusShape":
		res = GetNamedStructInstances(stage.GrowthCurveRhombusShapes, stage.GrowthCurveRhombusShape_stagedOrder)
	case "GrowthVectorShape":
		res = GetNamedStructInstances(stage.GrowthVectorShapes, stage.GrowthVectorShape_stagedOrder)
	case "InitialRhombusGridShape":
		res = GetNamedStructInstances(stage.InitialRhombusGridShapes, stage.InitialRhombusGridShape_stagedOrder)
	case "InitialRhombusShape":
		res = GetNamedStructInstances(stage.InitialRhombusShapes, stage.InitialRhombusShape_stagedOrder)
	case "Library":
		res = GetNamedStructInstances(stage.Librarys, stage.Library_stagedOrder)
	case "NextCircleShape":
		res = GetNamedStructInstances(stage.NextCircleShapes, stage.NextCircleShape_stagedOrder)
	case "PerpendicularVector":
		res = GetNamedStructInstances(stage.PerpendicularVectors, stage.PerpendicularVector_stagedOrder)
	case "PerpendicularVectorGrid":
		res = GetNamedStructInstances(stage.PerpendicularVectorGrids, stage.PerpendicularVectorGrid_stagedOrder)
	case "PerpendicularVectorGridHalfway":
		res = GetNamedStructInstances(stage.PerpendicularVectorGridHalfways, stage.PerpendicularVectorGridHalfway_stagedOrder)
	case "PerpendicularVectorHalfway":
		res = GetNamedStructInstances(stage.PerpendicularVectorHalfways, stage.PerpendicularVectorHalfway_stagedOrder)
	case "Plant":
		res = GetNamedStructInstances(stage.Plants, stage.Plant_stagedOrder)
	case "PlantCircumferenceShape":
		res = GetNamedStructInstances(stage.PlantCircumferenceShapes, stage.PlantCircumferenceShape_stagedOrder)
	case "PlantDiagram":
		res = GetNamedStructInstances(stage.PlantDiagrams, stage.PlantDiagram_stagedOrder)
	case "Rendered3DShape":
		res = GetNamedStructInstances(stage.Rendered3DShapes, stage.Rendered3DShape_stagedOrder)
	case "RhombusShape":
		res = GetNamedStructInstances(stage.RhombusShapes, stage.RhombusShape_stagedOrder)
	case "RotatedRhombusGridShape":
		res = GetNamedStructInstances(stage.RotatedRhombusGridShapes, stage.RotatedRhombusGridShape_stagedOrder)
	case "RotatedRhombusShape":
		res = GetNamedStructInstances(stage.RotatedRhombusShapes, stage.RotatedRhombusShape_stagedOrder)
	case "StackGrowthCurveBezierShape":
		res = GetNamedStructInstances(stage.StackGrowthCurveBezierShapes, stage.StackGrowthCurveBezierShape_stagedOrder)
	case "StackOfGrowthCurve":
		res = GetNamedStructInstances(stage.StackOfGrowthCurves, stage.StackOfGrowthCurve_stagedOrder)
	case "StartArcShape":
		res = GetNamedStructInstances(stage.StartArcShapes, stage.StartArcShape_stagedOrder)
	case "StartArcShapeGrid":
		res = GetNamedStructInstances(stage.StartArcShapeGrids, stage.StartArcShapeGrid_stagedOrder)
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
	CommitEndArcShape(endarcshape *EndArcShape)
	CheckoutEndArcShape(endarcshape *EndArcShape)
	CommitEndArcShapeGrid(endarcshapegrid *EndArcShapeGrid)
	CheckoutEndArcShapeGrid(endarcshapegrid *EndArcShapeGrid)
	CommitExplanationTextShape(explanationtextshape *ExplanationTextShape)
	CheckoutExplanationTextShape(explanationtextshape *ExplanationTextShape)
	CommitGridPathShape(gridpathshape *GridPathShape)
	CheckoutGridPathShape(gridpathshape *GridPathShape)
	CommitGrowthCurveBezierShape(growthcurvebeziershape *GrowthCurveBezierShape)
	CheckoutGrowthCurveBezierShape(growthcurvebeziershape *GrowthCurveBezierShape)
	CommitGrowthCurveBezierShapeGrid(growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid)
	CheckoutGrowthCurveBezierShapeGrid(growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid)
	CommitGrowthCurveRhombusGridShape(growthcurverhombusgridshape *GrowthCurveRhombusGridShape)
	CheckoutGrowthCurveRhombusGridShape(growthcurverhombusgridshape *GrowthCurveRhombusGridShape)
	CommitGrowthCurveRhombusShape(growthcurverhombusshape *GrowthCurveRhombusShape)
	CheckoutGrowthCurveRhombusShape(growthcurverhombusshape *GrowthCurveRhombusShape)
	CommitGrowthVectorShape(growthvectorshape *GrowthVectorShape)
	CheckoutGrowthVectorShape(growthvectorshape *GrowthVectorShape)
	CommitInitialRhombusGridShape(initialrhombusgridshape *InitialRhombusGridShape)
	CheckoutInitialRhombusGridShape(initialrhombusgridshape *InitialRhombusGridShape)
	CommitInitialRhombusShape(initialrhombusshape *InitialRhombusShape)
	CheckoutInitialRhombusShape(initialrhombusshape *InitialRhombusShape)
	CommitLibrary(library *Library)
	CheckoutLibrary(library *Library)
	CommitNextCircleShape(nextcircleshape *NextCircleShape)
	CheckoutNextCircleShape(nextcircleshape *NextCircleShape)
	CommitPerpendicularVector(perpendicularvector *PerpendicularVector)
	CheckoutPerpendicularVector(perpendicularvector *PerpendicularVector)
	CommitPerpendicularVectorGrid(perpendicularvectorgrid *PerpendicularVectorGrid)
	CheckoutPerpendicularVectorGrid(perpendicularvectorgrid *PerpendicularVectorGrid)
	CommitPerpendicularVectorGridHalfway(perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway)
	CheckoutPerpendicularVectorGridHalfway(perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway)
	CommitPerpendicularVectorHalfway(perpendicularvectorhalfway *PerpendicularVectorHalfway)
	CheckoutPerpendicularVectorHalfway(perpendicularvectorhalfway *PerpendicularVectorHalfway)
	CommitPlant(plant *Plant)
	CheckoutPlant(plant *Plant)
	CommitPlantCircumferenceShape(plantcircumferenceshape *PlantCircumferenceShape)
	CheckoutPlantCircumferenceShape(plantcircumferenceshape *PlantCircumferenceShape)
	CommitPlantDiagram(plantdiagram *PlantDiagram)
	CheckoutPlantDiagram(plantdiagram *PlantDiagram)
	CommitRendered3DShape(rendered3dshape *Rendered3DShape)
	CheckoutRendered3DShape(rendered3dshape *Rendered3DShape)
	CommitRhombusShape(rhombusshape *RhombusShape)
	CheckoutRhombusShape(rhombusshape *RhombusShape)
	CommitRotatedRhombusGridShape(rotatedrhombusgridshape *RotatedRhombusGridShape)
	CheckoutRotatedRhombusGridShape(rotatedrhombusgridshape *RotatedRhombusGridShape)
	CommitRotatedRhombusShape(rotatedrhombusshape *RotatedRhombusShape)
	CheckoutRotatedRhombusShape(rotatedrhombusshape *RotatedRhombusShape)
	CommitStackGrowthCurveBezierShape(stackgrowthcurvebeziershape *StackGrowthCurveBezierShape)
	CheckoutStackGrowthCurveBezierShape(stackgrowthcurvebeziershape *StackGrowthCurveBezierShape)
	CommitStackOfGrowthCurve(stackofgrowthcurve *StackOfGrowthCurve)
	CheckoutStackOfGrowthCurve(stackofgrowthcurve *StackOfGrowthCurve)
	CommitStartArcShape(startarcshape *StartArcShape)
	CheckoutStartArcShape(startarcshape *StartArcShape)
	CommitStartArcShapeGrid(startarcshapegrid *StartArcShapeGrid)
	CheckoutStartArcShapeGrid(startarcshapegrid *StartArcShapeGrid)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		AxesShapes:           make(map[*AxesShape]struct{}),
		AxesShapes_mapString: make(map[string]*AxesShape),

		CircleGridShapes:           make(map[*CircleGridShape]struct{}),
		CircleGridShapes_mapString: make(map[string]*CircleGridShape),

		EndArcShapes:           make(map[*EndArcShape]struct{}),
		EndArcShapes_mapString: make(map[string]*EndArcShape),

		EndArcShapeGrids:           make(map[*EndArcShapeGrid]struct{}),
		EndArcShapeGrids_mapString: make(map[string]*EndArcShapeGrid),

		ExplanationTextShapes:           make(map[*ExplanationTextShape]struct{}),
		ExplanationTextShapes_mapString: make(map[string]*ExplanationTextShape),

		GridPathShapes:           make(map[*GridPathShape]struct{}),
		GridPathShapes_mapString: make(map[string]*GridPathShape),

		GrowthCurveBezierShapes:           make(map[*GrowthCurveBezierShape]struct{}),
		GrowthCurveBezierShapes_mapString: make(map[string]*GrowthCurveBezierShape),

		GrowthCurveBezierShapeGrids:           make(map[*GrowthCurveBezierShapeGrid]struct{}),
		GrowthCurveBezierShapeGrids_mapString: make(map[string]*GrowthCurveBezierShapeGrid),

		GrowthCurveRhombusGridShapes:           make(map[*GrowthCurveRhombusGridShape]struct{}),
		GrowthCurveRhombusGridShapes_mapString: make(map[string]*GrowthCurveRhombusGridShape),

		GrowthCurveRhombusShapes:           make(map[*GrowthCurveRhombusShape]struct{}),
		GrowthCurveRhombusShapes_mapString: make(map[string]*GrowthCurveRhombusShape),

		GrowthVectorShapes:           make(map[*GrowthVectorShape]struct{}),
		GrowthVectorShapes_mapString: make(map[string]*GrowthVectorShape),

		InitialRhombusGridShapes:           make(map[*InitialRhombusGridShape]struct{}),
		InitialRhombusGridShapes_mapString: make(map[string]*InitialRhombusGridShape),

		InitialRhombusShapes:           make(map[*InitialRhombusShape]struct{}),
		InitialRhombusShapes_mapString: make(map[string]*InitialRhombusShape),

		Librarys:           make(map[*Library]struct{}),
		Librarys_mapString: make(map[string]*Library),

		NextCircleShapes:           make(map[*NextCircleShape]struct{}),
		NextCircleShapes_mapString: make(map[string]*NextCircleShape),

		PerpendicularVectors:           make(map[*PerpendicularVector]struct{}),
		PerpendicularVectors_mapString: make(map[string]*PerpendicularVector),

		PerpendicularVectorGrids:           make(map[*PerpendicularVectorGrid]struct{}),
		PerpendicularVectorGrids_mapString: make(map[string]*PerpendicularVectorGrid),

		PerpendicularVectorGridHalfways:           make(map[*PerpendicularVectorGridHalfway]struct{}),
		PerpendicularVectorGridHalfways_mapString: make(map[string]*PerpendicularVectorGridHalfway),

		PerpendicularVectorHalfways:           make(map[*PerpendicularVectorHalfway]struct{}),
		PerpendicularVectorHalfways_mapString: make(map[string]*PerpendicularVectorHalfway),

		Plants:           make(map[*Plant]struct{}),
		Plants_mapString: make(map[string]*Plant),

		PlantCircumferenceShapes:           make(map[*PlantCircumferenceShape]struct{}),
		PlantCircumferenceShapes_mapString: make(map[string]*PlantCircumferenceShape),

		PlantDiagrams:           make(map[*PlantDiagram]struct{}),
		PlantDiagrams_mapString: make(map[string]*PlantDiagram),

		Rendered3DShapes:           make(map[*Rendered3DShape]struct{}),
		Rendered3DShapes_mapString: make(map[string]*Rendered3DShape),

		RhombusShapes:           make(map[*RhombusShape]struct{}),
		RhombusShapes_mapString: make(map[string]*RhombusShape),

		RotatedRhombusGridShapes:           make(map[*RotatedRhombusGridShape]struct{}),
		RotatedRhombusGridShapes_mapString: make(map[string]*RotatedRhombusGridShape),

		RotatedRhombusShapes:           make(map[*RotatedRhombusShape]struct{}),
		RotatedRhombusShapes_mapString: make(map[string]*RotatedRhombusShape),

		StackGrowthCurveBezierShapes:           make(map[*StackGrowthCurveBezierShape]struct{}),
		StackGrowthCurveBezierShapes_mapString: make(map[string]*StackGrowthCurveBezierShape),

		StackOfGrowthCurves:           make(map[*StackOfGrowthCurve]struct{}),
		StackOfGrowthCurves_mapString: make(map[string]*StackOfGrowthCurve),

		StartArcShapes:           make(map[*StartArcShape]struct{}),
		StartArcShapes_mapString: make(map[string]*StartArcShape),

		StartArcShapeGrids:           make(map[*StartArcShapeGrid]struct{}),
		StartArcShapeGrids_mapString: make(map[string]*StartArcShapeGrid),

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

		EndArcShape_stagedOrder: make(map[*EndArcShape]uint),
		EndArcShape_orderStaged: make(map[uint]*EndArcShape),
		EndArcShapes_reference:  make(map[*EndArcShape]*EndArcShape),

		EndArcShapeGrid_stagedOrder: make(map[*EndArcShapeGrid]uint),
		EndArcShapeGrid_orderStaged: make(map[uint]*EndArcShapeGrid),
		EndArcShapeGrids_reference:  make(map[*EndArcShapeGrid]*EndArcShapeGrid),

		ExplanationTextShape_stagedOrder: make(map[*ExplanationTextShape]uint),
		ExplanationTextShape_orderStaged: make(map[uint]*ExplanationTextShape),
		ExplanationTextShapes_reference:  make(map[*ExplanationTextShape]*ExplanationTextShape),

		GridPathShape_stagedOrder: make(map[*GridPathShape]uint),
		GridPathShape_orderStaged: make(map[uint]*GridPathShape),
		GridPathShapes_reference:  make(map[*GridPathShape]*GridPathShape),

		GrowthCurveBezierShape_stagedOrder: make(map[*GrowthCurveBezierShape]uint),
		GrowthCurveBezierShape_orderStaged: make(map[uint]*GrowthCurveBezierShape),
		GrowthCurveBezierShapes_reference:  make(map[*GrowthCurveBezierShape]*GrowthCurveBezierShape),

		GrowthCurveBezierShapeGrid_stagedOrder: make(map[*GrowthCurveBezierShapeGrid]uint),
		GrowthCurveBezierShapeGrid_orderStaged: make(map[uint]*GrowthCurveBezierShapeGrid),
		GrowthCurveBezierShapeGrids_reference:  make(map[*GrowthCurveBezierShapeGrid]*GrowthCurveBezierShapeGrid),

		GrowthCurveRhombusGridShape_stagedOrder: make(map[*GrowthCurveRhombusGridShape]uint),
		GrowthCurveRhombusGridShape_orderStaged: make(map[uint]*GrowthCurveRhombusGridShape),
		GrowthCurveRhombusGridShapes_reference:  make(map[*GrowthCurveRhombusGridShape]*GrowthCurveRhombusGridShape),

		GrowthCurveRhombusShape_stagedOrder: make(map[*GrowthCurveRhombusShape]uint),
		GrowthCurveRhombusShape_orderStaged: make(map[uint]*GrowthCurveRhombusShape),
		GrowthCurveRhombusShapes_reference:  make(map[*GrowthCurveRhombusShape]*GrowthCurveRhombusShape),

		GrowthVectorShape_stagedOrder: make(map[*GrowthVectorShape]uint),
		GrowthVectorShape_orderStaged: make(map[uint]*GrowthVectorShape),
		GrowthVectorShapes_reference:  make(map[*GrowthVectorShape]*GrowthVectorShape),

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

		PerpendicularVector_stagedOrder: make(map[*PerpendicularVector]uint),
		PerpendicularVector_orderStaged: make(map[uint]*PerpendicularVector),
		PerpendicularVectors_reference:  make(map[*PerpendicularVector]*PerpendicularVector),

		PerpendicularVectorGrid_stagedOrder: make(map[*PerpendicularVectorGrid]uint),
		PerpendicularVectorGrid_orderStaged: make(map[uint]*PerpendicularVectorGrid),
		PerpendicularVectorGrids_reference:  make(map[*PerpendicularVectorGrid]*PerpendicularVectorGrid),

		PerpendicularVectorGridHalfway_stagedOrder: make(map[*PerpendicularVectorGridHalfway]uint),
		PerpendicularVectorGridHalfway_orderStaged: make(map[uint]*PerpendicularVectorGridHalfway),
		PerpendicularVectorGridHalfways_reference:  make(map[*PerpendicularVectorGridHalfway]*PerpendicularVectorGridHalfway),

		PerpendicularVectorHalfway_stagedOrder: make(map[*PerpendicularVectorHalfway]uint),
		PerpendicularVectorHalfway_orderStaged: make(map[uint]*PerpendicularVectorHalfway),
		PerpendicularVectorHalfways_reference:  make(map[*PerpendicularVectorHalfway]*PerpendicularVectorHalfway),

		Plant_stagedOrder: make(map[*Plant]uint),
		Plant_orderStaged: make(map[uint]*Plant),
		Plants_reference:  make(map[*Plant]*Plant),

		PlantCircumferenceShape_stagedOrder: make(map[*PlantCircumferenceShape]uint),
		PlantCircumferenceShape_orderStaged: make(map[uint]*PlantCircumferenceShape),
		PlantCircumferenceShapes_reference:  make(map[*PlantCircumferenceShape]*PlantCircumferenceShape),

		PlantDiagram_stagedOrder: make(map[*PlantDiagram]uint),
		PlantDiagram_orderStaged: make(map[uint]*PlantDiagram),
		PlantDiagrams_reference:  make(map[*PlantDiagram]*PlantDiagram),

		Rendered3DShape_stagedOrder: make(map[*Rendered3DShape]uint),
		Rendered3DShape_orderStaged: make(map[uint]*Rendered3DShape),
		Rendered3DShapes_reference:  make(map[*Rendered3DShape]*Rendered3DShape),

		RhombusShape_stagedOrder: make(map[*RhombusShape]uint),
		RhombusShape_orderStaged: make(map[uint]*RhombusShape),
		RhombusShapes_reference:  make(map[*RhombusShape]*RhombusShape),

		RotatedRhombusGridShape_stagedOrder: make(map[*RotatedRhombusGridShape]uint),
		RotatedRhombusGridShape_orderStaged: make(map[uint]*RotatedRhombusGridShape),
		RotatedRhombusGridShapes_reference:  make(map[*RotatedRhombusGridShape]*RotatedRhombusGridShape),

		RotatedRhombusShape_stagedOrder: make(map[*RotatedRhombusShape]uint),
		RotatedRhombusShape_orderStaged: make(map[uint]*RotatedRhombusShape),
		RotatedRhombusShapes_reference:  make(map[*RotatedRhombusShape]*RotatedRhombusShape),

		StackGrowthCurveBezierShape_stagedOrder: make(map[*StackGrowthCurveBezierShape]uint),
		StackGrowthCurveBezierShape_orderStaged: make(map[uint]*StackGrowthCurveBezierShape),
		StackGrowthCurveBezierShapes_reference:  make(map[*StackGrowthCurveBezierShape]*StackGrowthCurveBezierShape),

		StackOfGrowthCurve_stagedOrder: make(map[*StackOfGrowthCurve]uint),
		StackOfGrowthCurve_orderStaged: make(map[uint]*StackOfGrowthCurve),
		StackOfGrowthCurves_reference:  make(map[*StackOfGrowthCurve]*StackOfGrowthCurve),

		StartArcShape_stagedOrder: make(map[*StartArcShape]uint),
		StartArcShape_orderStaged: make(map[uint]*StartArcShape),
		StartArcShapes_reference:  make(map[*StartArcShape]*StartArcShape),

		StartArcShapeGrid_stagedOrder: make(map[*StartArcShapeGrid]uint),
		StartArcShapeGrid_orderStaged: make(map[uint]*StartArcShapeGrid),
		StartArcShapeGrids_reference:  make(map[*StartArcShapeGrid]*StartArcShapeGrid),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"AxesShape": &AxesShapeUnmarshaller{},

			"CircleGridShape": &CircleGridShapeUnmarshaller{},

			"EndArcShape": &EndArcShapeUnmarshaller{},

			"EndArcShapeGrid": &EndArcShapeGridUnmarshaller{},

			"ExplanationTextShape": &ExplanationTextShapeUnmarshaller{},

			"GridPathShape": &GridPathShapeUnmarshaller{},

			"GrowthCurveBezierShape": &GrowthCurveBezierShapeUnmarshaller{},

			"GrowthCurveBezierShapeGrid": &GrowthCurveBezierShapeGridUnmarshaller{},

			"GrowthCurveRhombusGridShape": &GrowthCurveRhombusGridShapeUnmarshaller{},

			"GrowthCurveRhombusShape": &GrowthCurveRhombusShapeUnmarshaller{},

			"GrowthVectorShape": &GrowthVectorShapeUnmarshaller{},

			"InitialRhombusGridShape": &InitialRhombusGridShapeUnmarshaller{},

			"InitialRhombusShape": &InitialRhombusShapeUnmarshaller{},

			"Library": &LibraryUnmarshaller{},

			"NextCircleShape": &NextCircleShapeUnmarshaller{},

			"PerpendicularVector": &PerpendicularVectorUnmarshaller{},

			"PerpendicularVectorGrid": &PerpendicularVectorGridUnmarshaller{},

			"PerpendicularVectorGridHalfway": &PerpendicularVectorGridHalfwayUnmarshaller{},

			"PerpendicularVectorHalfway": &PerpendicularVectorHalfwayUnmarshaller{},

			"Plant": &PlantUnmarshaller{},

			"PlantCircumferenceShape": &PlantCircumferenceShapeUnmarshaller{},

			"PlantDiagram": &PlantDiagramUnmarshaller{},

			"Rendered3DShape": &Rendered3DShapeUnmarshaller{},

			"RhombusShape": &RhombusShapeUnmarshaller{},

			"RotatedRhombusGridShape": &RotatedRhombusGridShapeUnmarshaller{},

			"RotatedRhombusShape": &RotatedRhombusShapeUnmarshaller{},

			"StackGrowthCurveBezierShape": &StackGrowthCurveBezierShapeUnmarshaller{},

			"StackOfGrowthCurve": &StackOfGrowthCurveUnmarshaller{},

			"StartArcShape": &StartArcShapeUnmarshaller{},

			"StartArcShapeGrid": &StartArcShapeGridUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "AxesShape"},
			{name: "CircleGridShape"},
			{name: "EndArcShape"},
			{name: "EndArcShapeGrid"},
			{name: "ExplanationTextShape"},
			{name: "GridPathShape"},
			{name: "GrowthCurveBezierShape"},
			{name: "GrowthCurveBezierShapeGrid"},
			{name: "GrowthCurveRhombusGridShape"},
			{name: "GrowthCurveRhombusShape"},
			{name: "GrowthVectorShape"},
			{name: "InitialRhombusGridShape"},
			{name: "InitialRhombusShape"},
			{name: "Library"},
			{name: "NextCircleShape"},
			{name: "PerpendicularVector"},
			{name: "PerpendicularVectorGrid"},
			{name: "PerpendicularVectorGridHalfway"},
			{name: "PerpendicularVectorHalfway"},
			{name: "Plant"},
			{name: "PlantCircumferenceShape"},
			{name: "PlantDiagram"},
			{name: "Rendered3DShape"},
			{name: "RhombusShape"},
			{name: "RotatedRhombusGridShape"},
			{name: "RotatedRhombusShape"},
			{name: "StackGrowthCurveBezierShape"},
			{name: "StackOfGrowthCurve"},
			{name: "StartArcShape"},
			{name: "StartArcShapeGrid"},
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
	case *EndArcShape:
		return stage.EndArcShape_stagedOrder[instance]
	case *EndArcShapeGrid:
		return stage.EndArcShapeGrid_stagedOrder[instance]
	case *ExplanationTextShape:
		return stage.ExplanationTextShape_stagedOrder[instance]
	case *GridPathShape:
		return stage.GridPathShape_stagedOrder[instance]
	case *GrowthCurveBezierShape:
		return stage.GrowthCurveBezierShape_stagedOrder[instance]
	case *GrowthCurveBezierShapeGrid:
		return stage.GrowthCurveBezierShapeGrid_stagedOrder[instance]
	case *GrowthCurveRhombusGridShape:
		return stage.GrowthCurveRhombusGridShape_stagedOrder[instance]
	case *GrowthCurveRhombusShape:
		return stage.GrowthCurveRhombusShape_stagedOrder[instance]
	case *GrowthVectorShape:
		return stage.GrowthVectorShape_stagedOrder[instance]
	case *InitialRhombusGridShape:
		return stage.InitialRhombusGridShape_stagedOrder[instance]
	case *InitialRhombusShape:
		return stage.InitialRhombusShape_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *NextCircleShape:
		return stage.NextCircleShape_stagedOrder[instance]
	case *PerpendicularVector:
		return stage.PerpendicularVector_stagedOrder[instance]
	case *PerpendicularVectorGrid:
		return stage.PerpendicularVectorGrid_stagedOrder[instance]
	case *PerpendicularVectorGridHalfway:
		return stage.PerpendicularVectorGridHalfway_stagedOrder[instance]
	case *PerpendicularVectorHalfway:
		return stage.PerpendicularVectorHalfway_stagedOrder[instance]
	case *Plant:
		return stage.Plant_stagedOrder[instance]
	case *PlantCircumferenceShape:
		return stage.PlantCircumferenceShape_stagedOrder[instance]
	case *PlantDiagram:
		return stage.PlantDiagram_stagedOrder[instance]
	case *Rendered3DShape:
		return stage.Rendered3DShape_stagedOrder[instance]
	case *RhombusShape:
		return stage.RhombusShape_stagedOrder[instance]
	case *RotatedRhombusGridShape:
		return stage.RotatedRhombusGridShape_stagedOrder[instance]
	case *RotatedRhombusShape:
		return stage.RotatedRhombusShape_stagedOrder[instance]
	case *StackGrowthCurveBezierShape:
		return stage.StackGrowthCurveBezierShape_stagedOrder[instance]
	case *StackOfGrowthCurve:
		return stage.StackOfGrowthCurve_stagedOrder[instance]
	case *StartArcShape:
		return stage.StartArcShape_stagedOrder[instance]
	case *StartArcShapeGrid:
		return stage.StartArcShapeGrid_stagedOrder[instance]
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
	case *EndArcShape:
		return any(stage.EndArcShape_orderStaged[order]).(Type)
	case *EndArcShapeGrid:
		return any(stage.EndArcShapeGrid_orderStaged[order]).(Type)
	case *ExplanationTextShape:
		return any(stage.ExplanationTextShape_orderStaged[order]).(Type)
	case *GridPathShape:
		return any(stage.GridPathShape_orderStaged[order]).(Type)
	case *GrowthCurveBezierShape:
		return any(stage.GrowthCurveBezierShape_orderStaged[order]).(Type)
	case *GrowthCurveBezierShapeGrid:
		return any(stage.GrowthCurveBezierShapeGrid_orderStaged[order]).(Type)
	case *GrowthCurveRhombusGridShape:
		return any(stage.GrowthCurveRhombusGridShape_orderStaged[order]).(Type)
	case *GrowthCurveRhombusShape:
		return any(stage.GrowthCurveRhombusShape_orderStaged[order]).(Type)
	case *GrowthVectorShape:
		return any(stage.GrowthVectorShape_orderStaged[order]).(Type)
	case *InitialRhombusGridShape:
		return any(stage.InitialRhombusGridShape_orderStaged[order]).(Type)
	case *InitialRhombusShape:
		return any(stage.InitialRhombusShape_orderStaged[order]).(Type)
	case *Library:
		return any(stage.Library_orderStaged[order]).(Type)
	case *NextCircleShape:
		return any(stage.NextCircleShape_orderStaged[order]).(Type)
	case *PerpendicularVector:
		return any(stage.PerpendicularVector_orderStaged[order]).(Type)
	case *PerpendicularVectorGrid:
		return any(stage.PerpendicularVectorGrid_orderStaged[order]).(Type)
	case *PerpendicularVectorGridHalfway:
		return any(stage.PerpendicularVectorGridHalfway_orderStaged[order]).(Type)
	case *PerpendicularVectorHalfway:
		return any(stage.PerpendicularVectorHalfway_orderStaged[order]).(Type)
	case *Plant:
		return any(stage.Plant_orderStaged[order]).(Type)
	case *PlantCircumferenceShape:
		return any(stage.PlantCircumferenceShape_orderStaged[order]).(Type)
	case *PlantDiagram:
		return any(stage.PlantDiagram_orderStaged[order]).(Type)
	case *Rendered3DShape:
		return any(stage.Rendered3DShape_orderStaged[order]).(Type)
	case *RhombusShape:
		return any(stage.RhombusShape_orderStaged[order]).(Type)
	case *RotatedRhombusGridShape:
		return any(stage.RotatedRhombusGridShape_orderStaged[order]).(Type)
	case *RotatedRhombusShape:
		return any(stage.RotatedRhombusShape_orderStaged[order]).(Type)
	case *StackGrowthCurveBezierShape:
		return any(stage.StackGrowthCurveBezierShape_orderStaged[order]).(Type)
	case *StackOfGrowthCurve:
		return any(stage.StackOfGrowthCurve_orderStaged[order]).(Type)
	case *StartArcShape:
		return any(stage.StartArcShape_orderStaged[order]).(Type)
	case *StartArcShapeGrid:
		return any(stage.StartArcShapeGrid_orderStaged[order]).(Type)
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
	case *EndArcShape:
		return stage.EndArcShape_stagedOrder[instance]
	case *EndArcShapeGrid:
		return stage.EndArcShapeGrid_stagedOrder[instance]
	case *ExplanationTextShape:
		return stage.ExplanationTextShape_stagedOrder[instance]
	case *GridPathShape:
		return stage.GridPathShape_stagedOrder[instance]
	case *GrowthCurveBezierShape:
		return stage.GrowthCurveBezierShape_stagedOrder[instance]
	case *GrowthCurveBezierShapeGrid:
		return stage.GrowthCurveBezierShapeGrid_stagedOrder[instance]
	case *GrowthCurveRhombusGridShape:
		return stage.GrowthCurveRhombusGridShape_stagedOrder[instance]
	case *GrowthCurveRhombusShape:
		return stage.GrowthCurveRhombusShape_stagedOrder[instance]
	case *GrowthVectorShape:
		return stage.GrowthVectorShape_stagedOrder[instance]
	case *InitialRhombusGridShape:
		return stage.InitialRhombusGridShape_stagedOrder[instance]
	case *InitialRhombusShape:
		return stage.InitialRhombusShape_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *NextCircleShape:
		return stage.NextCircleShape_stagedOrder[instance]
	case *PerpendicularVector:
		return stage.PerpendicularVector_stagedOrder[instance]
	case *PerpendicularVectorGrid:
		return stage.PerpendicularVectorGrid_stagedOrder[instance]
	case *PerpendicularVectorGridHalfway:
		return stage.PerpendicularVectorGridHalfway_stagedOrder[instance]
	case *PerpendicularVectorHalfway:
		return stage.PerpendicularVectorHalfway_stagedOrder[instance]
	case *Plant:
		return stage.Plant_stagedOrder[instance]
	case *PlantCircumferenceShape:
		return stage.PlantCircumferenceShape_stagedOrder[instance]
	case *PlantDiagram:
		return stage.PlantDiagram_stagedOrder[instance]
	case *Rendered3DShape:
		return stage.Rendered3DShape_stagedOrder[instance]
	case *RhombusShape:
		return stage.RhombusShape_stagedOrder[instance]
	case *RotatedRhombusGridShape:
		return stage.RotatedRhombusGridShape_stagedOrder[instance]
	case *RotatedRhombusShape:
		return stage.RotatedRhombusShape_stagedOrder[instance]
	case *StackGrowthCurveBezierShape:
		return stage.StackGrowthCurveBezierShape_stagedOrder[instance]
	case *StackOfGrowthCurve:
		return stage.StackOfGrowthCurve_stagedOrder[instance]
	case *StartArcShape:
		return stage.StartArcShape_stagedOrder[instance]
	case *StartArcShapeGrid:
		return stage.StartArcShapeGrid_stagedOrder[instance]
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
	stage.Map_GongStructName_InstancesNb["EndArcShape"] = len(stage.EndArcShapes)
	stage.Map_GongStructName_InstancesNb["EndArcShapeGrid"] = len(stage.EndArcShapeGrids)
	stage.Map_GongStructName_InstancesNb["ExplanationTextShape"] = len(stage.ExplanationTextShapes)
	stage.Map_GongStructName_InstancesNb["GridPathShape"] = len(stage.GridPathShapes)
	stage.Map_GongStructName_InstancesNb["GrowthCurveBezierShape"] = len(stage.GrowthCurveBezierShapes)
	stage.Map_GongStructName_InstancesNb["GrowthCurveBezierShapeGrid"] = len(stage.GrowthCurveBezierShapeGrids)
	stage.Map_GongStructName_InstancesNb["GrowthCurveRhombusGridShape"] = len(stage.GrowthCurveRhombusGridShapes)
	stage.Map_GongStructName_InstancesNb["GrowthCurveRhombusShape"] = len(stage.GrowthCurveRhombusShapes)
	stage.Map_GongStructName_InstancesNb["GrowthVectorShape"] = len(stage.GrowthVectorShapes)
	stage.Map_GongStructName_InstancesNb["InitialRhombusGridShape"] = len(stage.InitialRhombusGridShapes)
	stage.Map_GongStructName_InstancesNb["InitialRhombusShape"] = len(stage.InitialRhombusShapes)
	stage.Map_GongStructName_InstancesNb["Library"] = len(stage.Librarys)
	stage.Map_GongStructName_InstancesNb["NextCircleShape"] = len(stage.NextCircleShapes)
	stage.Map_GongStructName_InstancesNb["PerpendicularVector"] = len(stage.PerpendicularVectors)
	stage.Map_GongStructName_InstancesNb["PerpendicularVectorGrid"] = len(stage.PerpendicularVectorGrids)
	stage.Map_GongStructName_InstancesNb["PerpendicularVectorGridHalfway"] = len(stage.PerpendicularVectorGridHalfways)
	stage.Map_GongStructName_InstancesNb["PerpendicularVectorHalfway"] = len(stage.PerpendicularVectorHalfways)
	stage.Map_GongStructName_InstancesNb["Plant"] = len(stage.Plants)
	stage.Map_GongStructName_InstancesNb["PlantCircumferenceShape"] = len(stage.PlantCircumferenceShapes)
	stage.Map_GongStructName_InstancesNb["PlantDiagram"] = len(stage.PlantDiagrams)
	stage.Map_GongStructName_InstancesNb["Rendered3DShape"] = len(stage.Rendered3DShapes)
	stage.Map_GongStructName_InstancesNb["RhombusShape"] = len(stage.RhombusShapes)
	stage.Map_GongStructName_InstancesNb["RotatedRhombusGridShape"] = len(stage.RotatedRhombusGridShapes)
	stage.Map_GongStructName_InstancesNb["RotatedRhombusShape"] = len(stage.RotatedRhombusShapes)
	stage.Map_GongStructName_InstancesNb["StackGrowthCurveBezierShape"] = len(stage.StackGrowthCurveBezierShapes)
	stage.Map_GongStructName_InstancesNb["StackOfGrowthCurve"] = len(stage.StackOfGrowthCurves)
	stage.Map_GongStructName_InstancesNb["StartArcShape"] = len(stage.StartArcShapes)
	stage.Map_GongStructName_InstancesNb["StartArcShapeGrid"] = len(stage.StartArcShapeGrids)
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

// Stage puts endarcshape to the model stage
func (endarcshape *EndArcShape) Stage(stage *Stage) *EndArcShape {
	if _, ok := stage.EndArcShapes[endarcshape]; !ok {
		stage.EndArcShapes[endarcshape] = struct{}{}
		stage.EndArcShape_stagedOrder[endarcshape] = stage.EndArcShapeOrder
		stage.EndArcShape_orderStaged[stage.EndArcShapeOrder] = endarcshape
		stage.EndArcShapeOrder++
	}
	stage.EndArcShapes_mapString[endarcshape.Name] = endarcshape

	return endarcshape
}

// StagePreserveOrder puts endarcshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EndArcShapeOrder
// - update stage.EndArcShapeOrder accordingly
func (endarcshape *EndArcShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.EndArcShapes[endarcshape]; !ok {
		stage.EndArcShapes[endarcshape] = struct{}{}

		if order > stage.EndArcShapeOrder {
			stage.EndArcShapeOrder = order
		}
		stage.EndArcShape_stagedOrder[endarcshape] = order
		stage.EndArcShape_orderStaged[order] = endarcshape
		stage.EndArcShapeOrder++
	}
	stage.EndArcShapes_mapString[endarcshape.Name] = endarcshape
}

// Unstage removes endarcshape off the model stage
func (endarcshape *EndArcShape) Unstage(stage *Stage) *EndArcShape {
	delete(stage.EndArcShapes, endarcshape)
	// issue1150
	// delete(stage.EndArcShape_stagedOrder, endarcshape)
	delete(stage.EndArcShapes_mapString, endarcshape.Name)

	return endarcshape
}

// UnstageVoid removes endarcshape off the model stage
func (endarcshape *EndArcShape) UnstageVoid(stage *Stage) {
	delete(stage.EndArcShapes, endarcshape)
	// issue1150
	// delete(stage.EndArcShape_stagedOrder, endarcshape)
	delete(stage.EndArcShapes_mapString, endarcshape.Name)
}

// commit endarcshape to the back repo (if it is already staged)
func (endarcshape *EndArcShape) Commit(stage *Stage) *EndArcShape {
	if _, ok := stage.EndArcShapes[endarcshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitEndArcShape(endarcshape)
		}
	}
	return endarcshape
}

func (endarcshape *EndArcShape) CommitVoid(stage *Stage) {
	endarcshape.Commit(stage)
}

func (endarcshape *EndArcShape) StageVoid(stage *Stage) {
	endarcshape.Stage(stage)
}

// Checkout endarcshape to the back repo (if it is already staged)
func (endarcshape *EndArcShape) Checkout(stage *Stage) *EndArcShape {
	if _, ok := stage.EndArcShapes[endarcshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutEndArcShape(endarcshape)
		}
	}
	return endarcshape
}

// for satisfaction of GongStruct interface
func (endarcshape *EndArcShape) GetName() (res string) {
	return endarcshape.Name
}

// for satisfaction of GongStruct interface
func (endarcshape *EndArcShape) SetName(name string) {
	endarcshape.Name = name
}

// Stage puts endarcshapegrid to the model stage
func (endarcshapegrid *EndArcShapeGrid) Stage(stage *Stage) *EndArcShapeGrid {
	if _, ok := stage.EndArcShapeGrids[endarcshapegrid]; !ok {
		stage.EndArcShapeGrids[endarcshapegrid] = struct{}{}
		stage.EndArcShapeGrid_stagedOrder[endarcshapegrid] = stage.EndArcShapeGridOrder
		stage.EndArcShapeGrid_orderStaged[stage.EndArcShapeGridOrder] = endarcshapegrid
		stage.EndArcShapeGridOrder++
	}
	stage.EndArcShapeGrids_mapString[endarcshapegrid.Name] = endarcshapegrid

	return endarcshapegrid
}

// StagePreserveOrder puts endarcshapegrid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EndArcShapeGridOrder
// - update stage.EndArcShapeGridOrder accordingly
func (endarcshapegrid *EndArcShapeGrid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.EndArcShapeGrids[endarcshapegrid]; !ok {
		stage.EndArcShapeGrids[endarcshapegrid] = struct{}{}

		if order > stage.EndArcShapeGridOrder {
			stage.EndArcShapeGridOrder = order
		}
		stage.EndArcShapeGrid_stagedOrder[endarcshapegrid] = order
		stage.EndArcShapeGrid_orderStaged[order] = endarcshapegrid
		stage.EndArcShapeGridOrder++
	}
	stage.EndArcShapeGrids_mapString[endarcshapegrid.Name] = endarcshapegrid
}

// Unstage removes endarcshapegrid off the model stage
func (endarcshapegrid *EndArcShapeGrid) Unstage(stage *Stage) *EndArcShapeGrid {
	delete(stage.EndArcShapeGrids, endarcshapegrid)
	// issue1150
	// delete(stage.EndArcShapeGrid_stagedOrder, endarcshapegrid)
	delete(stage.EndArcShapeGrids_mapString, endarcshapegrid.Name)

	return endarcshapegrid
}

// UnstageVoid removes endarcshapegrid off the model stage
func (endarcshapegrid *EndArcShapeGrid) UnstageVoid(stage *Stage) {
	delete(stage.EndArcShapeGrids, endarcshapegrid)
	// issue1150
	// delete(stage.EndArcShapeGrid_stagedOrder, endarcshapegrid)
	delete(stage.EndArcShapeGrids_mapString, endarcshapegrid.Name)
}

// commit endarcshapegrid to the back repo (if it is already staged)
func (endarcshapegrid *EndArcShapeGrid) Commit(stage *Stage) *EndArcShapeGrid {
	if _, ok := stage.EndArcShapeGrids[endarcshapegrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitEndArcShapeGrid(endarcshapegrid)
		}
	}
	return endarcshapegrid
}

func (endarcshapegrid *EndArcShapeGrid) CommitVoid(stage *Stage) {
	endarcshapegrid.Commit(stage)
}

func (endarcshapegrid *EndArcShapeGrid) StageVoid(stage *Stage) {
	endarcshapegrid.Stage(stage)
}

// Checkout endarcshapegrid to the back repo (if it is already staged)
func (endarcshapegrid *EndArcShapeGrid) Checkout(stage *Stage) *EndArcShapeGrid {
	if _, ok := stage.EndArcShapeGrids[endarcshapegrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutEndArcShapeGrid(endarcshapegrid)
		}
	}
	return endarcshapegrid
}

// for satisfaction of GongStruct interface
func (endarcshapegrid *EndArcShapeGrid) GetName() (res string) {
	return endarcshapegrid.Name
}

// for satisfaction of GongStruct interface
func (endarcshapegrid *EndArcShapeGrid) SetName(name string) {
	endarcshapegrid.Name = name
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

// Stage puts growthcurvebeziershape to the model stage
func (growthcurvebeziershape *GrowthCurveBezierShape) Stage(stage *Stage) *GrowthCurveBezierShape {
	if _, ok := stage.GrowthCurveBezierShapes[growthcurvebeziershape]; !ok {
		stage.GrowthCurveBezierShapes[growthcurvebeziershape] = struct{}{}
		stage.GrowthCurveBezierShape_stagedOrder[growthcurvebeziershape] = stage.GrowthCurveBezierShapeOrder
		stage.GrowthCurveBezierShape_orderStaged[stage.GrowthCurveBezierShapeOrder] = growthcurvebeziershape
		stage.GrowthCurveBezierShapeOrder++
	}
	stage.GrowthCurveBezierShapes_mapString[growthcurvebeziershape.Name] = growthcurvebeziershape

	return growthcurvebeziershape
}

// StagePreserveOrder puts growthcurvebeziershape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GrowthCurveBezierShapeOrder
// - update stage.GrowthCurveBezierShapeOrder accordingly
func (growthcurvebeziershape *GrowthCurveBezierShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.GrowthCurveBezierShapes[growthcurvebeziershape]; !ok {
		stage.GrowthCurveBezierShapes[growthcurvebeziershape] = struct{}{}

		if order > stage.GrowthCurveBezierShapeOrder {
			stage.GrowthCurveBezierShapeOrder = order
		}
		stage.GrowthCurveBezierShape_stagedOrder[growthcurvebeziershape] = order
		stage.GrowthCurveBezierShape_orderStaged[order] = growthcurvebeziershape
		stage.GrowthCurveBezierShapeOrder++
	}
	stage.GrowthCurveBezierShapes_mapString[growthcurvebeziershape.Name] = growthcurvebeziershape
}

// Unstage removes growthcurvebeziershape off the model stage
func (growthcurvebeziershape *GrowthCurveBezierShape) Unstage(stage *Stage) *GrowthCurveBezierShape {
	delete(stage.GrowthCurveBezierShapes, growthcurvebeziershape)
	// issue1150
	// delete(stage.GrowthCurveBezierShape_stagedOrder, growthcurvebeziershape)
	delete(stage.GrowthCurveBezierShapes_mapString, growthcurvebeziershape.Name)

	return growthcurvebeziershape
}

// UnstageVoid removes growthcurvebeziershape off the model stage
func (growthcurvebeziershape *GrowthCurveBezierShape) UnstageVoid(stage *Stage) {
	delete(stage.GrowthCurveBezierShapes, growthcurvebeziershape)
	// issue1150
	// delete(stage.GrowthCurveBezierShape_stagedOrder, growthcurvebeziershape)
	delete(stage.GrowthCurveBezierShapes_mapString, growthcurvebeziershape.Name)
}

// commit growthcurvebeziershape to the back repo (if it is already staged)
func (growthcurvebeziershape *GrowthCurveBezierShape) Commit(stage *Stage) *GrowthCurveBezierShape {
	if _, ok := stage.GrowthCurveBezierShapes[growthcurvebeziershape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGrowthCurveBezierShape(growthcurvebeziershape)
		}
	}
	return growthcurvebeziershape
}

func (growthcurvebeziershape *GrowthCurveBezierShape) CommitVoid(stage *Stage) {
	growthcurvebeziershape.Commit(stage)
}

func (growthcurvebeziershape *GrowthCurveBezierShape) StageVoid(stage *Stage) {
	growthcurvebeziershape.Stage(stage)
}

// Checkout growthcurvebeziershape to the back repo (if it is already staged)
func (growthcurvebeziershape *GrowthCurveBezierShape) Checkout(stage *Stage) *GrowthCurveBezierShape {
	if _, ok := stage.GrowthCurveBezierShapes[growthcurvebeziershape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGrowthCurveBezierShape(growthcurvebeziershape)
		}
	}
	return growthcurvebeziershape
}

// for satisfaction of GongStruct interface
func (growthcurvebeziershape *GrowthCurveBezierShape) GetName() (res string) {
	return growthcurvebeziershape.Name
}

// for satisfaction of GongStruct interface
func (growthcurvebeziershape *GrowthCurveBezierShape) SetName(name string) {
	growthcurvebeziershape.Name = name
}

// Stage puts growthcurvebeziershapegrid to the model stage
func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) Stage(stage *Stage) *GrowthCurveBezierShapeGrid {
	if _, ok := stage.GrowthCurveBezierShapeGrids[growthcurvebeziershapegrid]; !ok {
		stage.GrowthCurveBezierShapeGrids[growthcurvebeziershapegrid] = struct{}{}
		stage.GrowthCurveBezierShapeGrid_stagedOrder[growthcurvebeziershapegrid] = stage.GrowthCurveBezierShapeGridOrder
		stage.GrowthCurveBezierShapeGrid_orderStaged[stage.GrowthCurveBezierShapeGridOrder] = growthcurvebeziershapegrid
		stage.GrowthCurveBezierShapeGridOrder++
	}
	stage.GrowthCurveBezierShapeGrids_mapString[growthcurvebeziershapegrid.Name] = growthcurvebeziershapegrid

	return growthcurvebeziershapegrid
}

// StagePreserveOrder puts growthcurvebeziershapegrid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GrowthCurveBezierShapeGridOrder
// - update stage.GrowthCurveBezierShapeGridOrder accordingly
func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.GrowthCurveBezierShapeGrids[growthcurvebeziershapegrid]; !ok {
		stage.GrowthCurveBezierShapeGrids[growthcurvebeziershapegrid] = struct{}{}

		if order > stage.GrowthCurveBezierShapeGridOrder {
			stage.GrowthCurveBezierShapeGridOrder = order
		}
		stage.GrowthCurveBezierShapeGrid_stagedOrder[growthcurvebeziershapegrid] = order
		stage.GrowthCurveBezierShapeGrid_orderStaged[order] = growthcurvebeziershapegrid
		stage.GrowthCurveBezierShapeGridOrder++
	}
	stage.GrowthCurveBezierShapeGrids_mapString[growthcurvebeziershapegrid.Name] = growthcurvebeziershapegrid
}

// Unstage removes growthcurvebeziershapegrid off the model stage
func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) Unstage(stage *Stage) *GrowthCurveBezierShapeGrid {
	delete(stage.GrowthCurveBezierShapeGrids, growthcurvebeziershapegrid)
	// issue1150
	// delete(stage.GrowthCurveBezierShapeGrid_stagedOrder, growthcurvebeziershapegrid)
	delete(stage.GrowthCurveBezierShapeGrids_mapString, growthcurvebeziershapegrid.Name)

	return growthcurvebeziershapegrid
}

// UnstageVoid removes growthcurvebeziershapegrid off the model stage
func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) UnstageVoid(stage *Stage) {
	delete(stage.GrowthCurveBezierShapeGrids, growthcurvebeziershapegrid)
	// issue1150
	// delete(stage.GrowthCurveBezierShapeGrid_stagedOrder, growthcurvebeziershapegrid)
	delete(stage.GrowthCurveBezierShapeGrids_mapString, growthcurvebeziershapegrid.Name)
}

// commit growthcurvebeziershapegrid to the back repo (if it is already staged)
func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) Commit(stage *Stage) *GrowthCurveBezierShapeGrid {
	if _, ok := stage.GrowthCurveBezierShapeGrids[growthcurvebeziershapegrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGrowthCurveBezierShapeGrid(growthcurvebeziershapegrid)
		}
	}
	return growthcurvebeziershapegrid
}

func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) CommitVoid(stage *Stage) {
	growthcurvebeziershapegrid.Commit(stage)
}

func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) StageVoid(stage *Stage) {
	growthcurvebeziershapegrid.Stage(stage)
}

// Checkout growthcurvebeziershapegrid to the back repo (if it is already staged)
func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) Checkout(stage *Stage) *GrowthCurveBezierShapeGrid {
	if _, ok := stage.GrowthCurveBezierShapeGrids[growthcurvebeziershapegrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGrowthCurveBezierShapeGrid(growthcurvebeziershapegrid)
		}
	}
	return growthcurvebeziershapegrid
}

// for satisfaction of GongStruct interface
func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) GetName() (res string) {
	return growthcurvebeziershapegrid.Name
}

// for satisfaction of GongStruct interface
func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) SetName(name string) {
	growthcurvebeziershapegrid.Name = name
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

// Stage puts growthvectorshape to the model stage
func (growthvectorshape *GrowthVectorShape) Stage(stage *Stage) *GrowthVectorShape {
	if _, ok := stage.GrowthVectorShapes[growthvectorshape]; !ok {
		stage.GrowthVectorShapes[growthvectorshape] = struct{}{}
		stage.GrowthVectorShape_stagedOrder[growthvectorshape] = stage.GrowthVectorShapeOrder
		stage.GrowthVectorShape_orderStaged[stage.GrowthVectorShapeOrder] = growthvectorshape
		stage.GrowthVectorShapeOrder++
	}
	stage.GrowthVectorShapes_mapString[growthvectorshape.Name] = growthvectorshape

	return growthvectorshape
}

// StagePreserveOrder puts growthvectorshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GrowthVectorShapeOrder
// - update stage.GrowthVectorShapeOrder accordingly
func (growthvectorshape *GrowthVectorShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.GrowthVectorShapes[growthvectorshape]; !ok {
		stage.GrowthVectorShapes[growthvectorshape] = struct{}{}

		if order > stage.GrowthVectorShapeOrder {
			stage.GrowthVectorShapeOrder = order
		}
		stage.GrowthVectorShape_stagedOrder[growthvectorshape] = order
		stage.GrowthVectorShape_orderStaged[order] = growthvectorshape
		stage.GrowthVectorShapeOrder++
	}
	stage.GrowthVectorShapes_mapString[growthvectorshape.Name] = growthvectorshape
}

// Unstage removes growthvectorshape off the model stage
func (growthvectorshape *GrowthVectorShape) Unstage(stage *Stage) *GrowthVectorShape {
	delete(stage.GrowthVectorShapes, growthvectorshape)
	// issue1150
	// delete(stage.GrowthVectorShape_stagedOrder, growthvectorshape)
	delete(stage.GrowthVectorShapes_mapString, growthvectorshape.Name)

	return growthvectorshape
}

// UnstageVoid removes growthvectorshape off the model stage
func (growthvectorshape *GrowthVectorShape) UnstageVoid(stage *Stage) {
	delete(stage.GrowthVectorShapes, growthvectorshape)
	// issue1150
	// delete(stage.GrowthVectorShape_stagedOrder, growthvectorshape)
	delete(stage.GrowthVectorShapes_mapString, growthvectorshape.Name)
}

// commit growthvectorshape to the back repo (if it is already staged)
func (growthvectorshape *GrowthVectorShape) Commit(stage *Stage) *GrowthVectorShape {
	if _, ok := stage.GrowthVectorShapes[growthvectorshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGrowthVectorShape(growthvectorshape)
		}
	}
	return growthvectorshape
}

func (growthvectorshape *GrowthVectorShape) CommitVoid(stage *Stage) {
	growthvectorshape.Commit(stage)
}

func (growthvectorshape *GrowthVectorShape) StageVoid(stage *Stage) {
	growthvectorshape.Stage(stage)
}

// Checkout growthvectorshape to the back repo (if it is already staged)
func (growthvectorshape *GrowthVectorShape) Checkout(stage *Stage) *GrowthVectorShape {
	if _, ok := stage.GrowthVectorShapes[growthvectorshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGrowthVectorShape(growthvectorshape)
		}
	}
	return growthvectorshape
}

// for satisfaction of GongStruct interface
func (growthvectorshape *GrowthVectorShape) GetName() (res string) {
	return growthvectorshape.Name
}

// for satisfaction of GongStruct interface
func (growthvectorshape *GrowthVectorShape) SetName(name string) {
	growthvectorshape.Name = name
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

// Stage puts perpendicularvector to the model stage
func (perpendicularvector *PerpendicularVector) Stage(stage *Stage) *PerpendicularVector {
	if _, ok := stage.PerpendicularVectors[perpendicularvector]; !ok {
		stage.PerpendicularVectors[perpendicularvector] = struct{}{}
		stage.PerpendicularVector_stagedOrder[perpendicularvector] = stage.PerpendicularVectorOrder
		stage.PerpendicularVector_orderStaged[stage.PerpendicularVectorOrder] = perpendicularvector
		stage.PerpendicularVectorOrder++
	}
	stage.PerpendicularVectors_mapString[perpendicularvector.Name] = perpendicularvector

	return perpendicularvector
}

// StagePreserveOrder puts perpendicularvector to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PerpendicularVectorOrder
// - update stage.PerpendicularVectorOrder accordingly
func (perpendicularvector *PerpendicularVector) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.PerpendicularVectors[perpendicularvector]; !ok {
		stage.PerpendicularVectors[perpendicularvector] = struct{}{}

		if order > stage.PerpendicularVectorOrder {
			stage.PerpendicularVectorOrder = order
		}
		stage.PerpendicularVector_stagedOrder[perpendicularvector] = order
		stage.PerpendicularVector_orderStaged[order] = perpendicularvector
		stage.PerpendicularVectorOrder++
	}
	stage.PerpendicularVectors_mapString[perpendicularvector.Name] = perpendicularvector
}

// Unstage removes perpendicularvector off the model stage
func (perpendicularvector *PerpendicularVector) Unstage(stage *Stage) *PerpendicularVector {
	delete(stage.PerpendicularVectors, perpendicularvector)
	// issue1150
	// delete(stage.PerpendicularVector_stagedOrder, perpendicularvector)
	delete(stage.PerpendicularVectors_mapString, perpendicularvector.Name)

	return perpendicularvector
}

// UnstageVoid removes perpendicularvector off the model stage
func (perpendicularvector *PerpendicularVector) UnstageVoid(stage *Stage) {
	delete(stage.PerpendicularVectors, perpendicularvector)
	// issue1150
	// delete(stage.PerpendicularVector_stagedOrder, perpendicularvector)
	delete(stage.PerpendicularVectors_mapString, perpendicularvector.Name)
}

// commit perpendicularvector to the back repo (if it is already staged)
func (perpendicularvector *PerpendicularVector) Commit(stage *Stage) *PerpendicularVector {
	if _, ok := stage.PerpendicularVectors[perpendicularvector]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPerpendicularVector(perpendicularvector)
		}
	}
	return perpendicularvector
}

func (perpendicularvector *PerpendicularVector) CommitVoid(stage *Stage) {
	perpendicularvector.Commit(stage)
}

func (perpendicularvector *PerpendicularVector) StageVoid(stage *Stage) {
	perpendicularvector.Stage(stage)
}

// Checkout perpendicularvector to the back repo (if it is already staged)
func (perpendicularvector *PerpendicularVector) Checkout(stage *Stage) *PerpendicularVector {
	if _, ok := stage.PerpendicularVectors[perpendicularvector]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPerpendicularVector(perpendicularvector)
		}
	}
	return perpendicularvector
}

// for satisfaction of GongStruct interface
func (perpendicularvector *PerpendicularVector) GetName() (res string) {
	return perpendicularvector.Name
}

// for satisfaction of GongStruct interface
func (perpendicularvector *PerpendicularVector) SetName(name string) {
	perpendicularvector.Name = name
}

// Stage puts perpendicularvectorgrid to the model stage
func (perpendicularvectorgrid *PerpendicularVectorGrid) Stage(stage *Stage) *PerpendicularVectorGrid {
	if _, ok := stage.PerpendicularVectorGrids[perpendicularvectorgrid]; !ok {
		stage.PerpendicularVectorGrids[perpendicularvectorgrid] = struct{}{}
		stage.PerpendicularVectorGrid_stagedOrder[perpendicularvectorgrid] = stage.PerpendicularVectorGridOrder
		stage.PerpendicularVectorGrid_orderStaged[stage.PerpendicularVectorGridOrder] = perpendicularvectorgrid
		stage.PerpendicularVectorGridOrder++
	}
	stage.PerpendicularVectorGrids_mapString[perpendicularvectorgrid.Name] = perpendicularvectorgrid

	return perpendicularvectorgrid
}

// StagePreserveOrder puts perpendicularvectorgrid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PerpendicularVectorGridOrder
// - update stage.PerpendicularVectorGridOrder accordingly
func (perpendicularvectorgrid *PerpendicularVectorGrid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.PerpendicularVectorGrids[perpendicularvectorgrid]; !ok {
		stage.PerpendicularVectorGrids[perpendicularvectorgrid] = struct{}{}

		if order > stage.PerpendicularVectorGridOrder {
			stage.PerpendicularVectorGridOrder = order
		}
		stage.PerpendicularVectorGrid_stagedOrder[perpendicularvectorgrid] = order
		stage.PerpendicularVectorGrid_orderStaged[order] = perpendicularvectorgrid
		stage.PerpendicularVectorGridOrder++
	}
	stage.PerpendicularVectorGrids_mapString[perpendicularvectorgrid.Name] = perpendicularvectorgrid
}

// Unstage removes perpendicularvectorgrid off the model stage
func (perpendicularvectorgrid *PerpendicularVectorGrid) Unstage(stage *Stage) *PerpendicularVectorGrid {
	delete(stage.PerpendicularVectorGrids, perpendicularvectorgrid)
	// issue1150
	// delete(stage.PerpendicularVectorGrid_stagedOrder, perpendicularvectorgrid)
	delete(stage.PerpendicularVectorGrids_mapString, perpendicularvectorgrid.Name)

	return perpendicularvectorgrid
}

// UnstageVoid removes perpendicularvectorgrid off the model stage
func (perpendicularvectorgrid *PerpendicularVectorGrid) UnstageVoid(stage *Stage) {
	delete(stage.PerpendicularVectorGrids, perpendicularvectorgrid)
	// issue1150
	// delete(stage.PerpendicularVectorGrid_stagedOrder, perpendicularvectorgrid)
	delete(stage.PerpendicularVectorGrids_mapString, perpendicularvectorgrid.Name)
}

// commit perpendicularvectorgrid to the back repo (if it is already staged)
func (perpendicularvectorgrid *PerpendicularVectorGrid) Commit(stage *Stage) *PerpendicularVectorGrid {
	if _, ok := stage.PerpendicularVectorGrids[perpendicularvectorgrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPerpendicularVectorGrid(perpendicularvectorgrid)
		}
	}
	return perpendicularvectorgrid
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) CommitVoid(stage *Stage) {
	perpendicularvectorgrid.Commit(stage)
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) StageVoid(stage *Stage) {
	perpendicularvectorgrid.Stage(stage)
}

// Checkout perpendicularvectorgrid to the back repo (if it is already staged)
func (perpendicularvectorgrid *PerpendicularVectorGrid) Checkout(stage *Stage) *PerpendicularVectorGrid {
	if _, ok := stage.PerpendicularVectorGrids[perpendicularvectorgrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPerpendicularVectorGrid(perpendicularvectorgrid)
		}
	}
	return perpendicularvectorgrid
}

// for satisfaction of GongStruct interface
func (perpendicularvectorgrid *PerpendicularVectorGrid) GetName() (res string) {
	return perpendicularvectorgrid.Name
}

// for satisfaction of GongStruct interface
func (perpendicularvectorgrid *PerpendicularVectorGrid) SetName(name string) {
	perpendicularvectorgrid.Name = name
}

// Stage puts perpendicularvectorgridhalfway to the model stage
func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) Stage(stage *Stage) *PerpendicularVectorGridHalfway {
	if _, ok := stage.PerpendicularVectorGridHalfways[perpendicularvectorgridhalfway]; !ok {
		stage.PerpendicularVectorGridHalfways[perpendicularvectorgridhalfway] = struct{}{}
		stage.PerpendicularVectorGridHalfway_stagedOrder[perpendicularvectorgridhalfway] = stage.PerpendicularVectorGridHalfwayOrder
		stage.PerpendicularVectorGridHalfway_orderStaged[stage.PerpendicularVectorGridHalfwayOrder] = perpendicularvectorgridhalfway
		stage.PerpendicularVectorGridHalfwayOrder++
	}
	stage.PerpendicularVectorGridHalfways_mapString[perpendicularvectorgridhalfway.Name] = perpendicularvectorgridhalfway

	return perpendicularvectorgridhalfway
}

// StagePreserveOrder puts perpendicularvectorgridhalfway to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PerpendicularVectorGridHalfwayOrder
// - update stage.PerpendicularVectorGridHalfwayOrder accordingly
func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.PerpendicularVectorGridHalfways[perpendicularvectorgridhalfway]; !ok {
		stage.PerpendicularVectorGridHalfways[perpendicularvectorgridhalfway] = struct{}{}

		if order > stage.PerpendicularVectorGridHalfwayOrder {
			stage.PerpendicularVectorGridHalfwayOrder = order
		}
		stage.PerpendicularVectorGridHalfway_stagedOrder[perpendicularvectorgridhalfway] = order
		stage.PerpendicularVectorGridHalfway_orderStaged[order] = perpendicularvectorgridhalfway
		stage.PerpendicularVectorGridHalfwayOrder++
	}
	stage.PerpendicularVectorGridHalfways_mapString[perpendicularvectorgridhalfway.Name] = perpendicularvectorgridhalfway
}

// Unstage removes perpendicularvectorgridhalfway off the model stage
func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) Unstage(stage *Stage) *PerpendicularVectorGridHalfway {
	delete(stage.PerpendicularVectorGridHalfways, perpendicularvectorgridhalfway)
	// issue1150
	// delete(stage.PerpendicularVectorGridHalfway_stagedOrder, perpendicularvectorgridhalfway)
	delete(stage.PerpendicularVectorGridHalfways_mapString, perpendicularvectorgridhalfway.Name)

	return perpendicularvectorgridhalfway
}

// UnstageVoid removes perpendicularvectorgridhalfway off the model stage
func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) UnstageVoid(stage *Stage) {
	delete(stage.PerpendicularVectorGridHalfways, perpendicularvectorgridhalfway)
	// issue1150
	// delete(stage.PerpendicularVectorGridHalfway_stagedOrder, perpendicularvectorgridhalfway)
	delete(stage.PerpendicularVectorGridHalfways_mapString, perpendicularvectorgridhalfway.Name)
}

// commit perpendicularvectorgridhalfway to the back repo (if it is already staged)
func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) Commit(stage *Stage) *PerpendicularVectorGridHalfway {
	if _, ok := stage.PerpendicularVectorGridHalfways[perpendicularvectorgridhalfway]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPerpendicularVectorGridHalfway(perpendicularvectorgridhalfway)
		}
	}
	return perpendicularvectorgridhalfway
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) CommitVoid(stage *Stage) {
	perpendicularvectorgridhalfway.Commit(stage)
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) StageVoid(stage *Stage) {
	perpendicularvectorgridhalfway.Stage(stage)
}

// Checkout perpendicularvectorgridhalfway to the back repo (if it is already staged)
func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) Checkout(stage *Stage) *PerpendicularVectorGridHalfway {
	if _, ok := stage.PerpendicularVectorGridHalfways[perpendicularvectorgridhalfway]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPerpendicularVectorGridHalfway(perpendicularvectorgridhalfway)
		}
	}
	return perpendicularvectorgridhalfway
}

// for satisfaction of GongStruct interface
func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GetName() (res string) {
	return perpendicularvectorgridhalfway.Name
}

// for satisfaction of GongStruct interface
func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) SetName(name string) {
	perpendicularvectorgridhalfway.Name = name
}

// Stage puts perpendicularvectorhalfway to the model stage
func (perpendicularvectorhalfway *PerpendicularVectorHalfway) Stage(stage *Stage) *PerpendicularVectorHalfway {
	if _, ok := stage.PerpendicularVectorHalfways[perpendicularvectorhalfway]; !ok {
		stage.PerpendicularVectorHalfways[perpendicularvectorhalfway] = struct{}{}
		stage.PerpendicularVectorHalfway_stagedOrder[perpendicularvectorhalfway] = stage.PerpendicularVectorHalfwayOrder
		stage.PerpendicularVectorHalfway_orderStaged[stage.PerpendicularVectorHalfwayOrder] = perpendicularvectorhalfway
		stage.PerpendicularVectorHalfwayOrder++
	}
	stage.PerpendicularVectorHalfways_mapString[perpendicularvectorhalfway.Name] = perpendicularvectorhalfway

	return perpendicularvectorhalfway
}

// StagePreserveOrder puts perpendicularvectorhalfway to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PerpendicularVectorHalfwayOrder
// - update stage.PerpendicularVectorHalfwayOrder accordingly
func (perpendicularvectorhalfway *PerpendicularVectorHalfway) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.PerpendicularVectorHalfways[perpendicularvectorhalfway]; !ok {
		stage.PerpendicularVectorHalfways[perpendicularvectorhalfway] = struct{}{}

		if order > stage.PerpendicularVectorHalfwayOrder {
			stage.PerpendicularVectorHalfwayOrder = order
		}
		stage.PerpendicularVectorHalfway_stagedOrder[perpendicularvectorhalfway] = order
		stage.PerpendicularVectorHalfway_orderStaged[order] = perpendicularvectorhalfway
		stage.PerpendicularVectorHalfwayOrder++
	}
	stage.PerpendicularVectorHalfways_mapString[perpendicularvectorhalfway.Name] = perpendicularvectorhalfway
}

// Unstage removes perpendicularvectorhalfway off the model stage
func (perpendicularvectorhalfway *PerpendicularVectorHalfway) Unstage(stage *Stage) *PerpendicularVectorHalfway {
	delete(stage.PerpendicularVectorHalfways, perpendicularvectorhalfway)
	// issue1150
	// delete(stage.PerpendicularVectorHalfway_stagedOrder, perpendicularvectorhalfway)
	delete(stage.PerpendicularVectorHalfways_mapString, perpendicularvectorhalfway.Name)

	return perpendicularvectorhalfway
}

// UnstageVoid removes perpendicularvectorhalfway off the model stage
func (perpendicularvectorhalfway *PerpendicularVectorHalfway) UnstageVoid(stage *Stage) {
	delete(stage.PerpendicularVectorHalfways, perpendicularvectorhalfway)
	// issue1150
	// delete(stage.PerpendicularVectorHalfway_stagedOrder, perpendicularvectorhalfway)
	delete(stage.PerpendicularVectorHalfways_mapString, perpendicularvectorhalfway.Name)
}

// commit perpendicularvectorhalfway to the back repo (if it is already staged)
func (perpendicularvectorhalfway *PerpendicularVectorHalfway) Commit(stage *Stage) *PerpendicularVectorHalfway {
	if _, ok := stage.PerpendicularVectorHalfways[perpendicularvectorhalfway]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPerpendicularVectorHalfway(perpendicularvectorhalfway)
		}
	}
	return perpendicularvectorhalfway
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) CommitVoid(stage *Stage) {
	perpendicularvectorhalfway.Commit(stage)
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) StageVoid(stage *Stage) {
	perpendicularvectorhalfway.Stage(stage)
}

// Checkout perpendicularvectorhalfway to the back repo (if it is already staged)
func (perpendicularvectorhalfway *PerpendicularVectorHalfway) Checkout(stage *Stage) *PerpendicularVectorHalfway {
	if _, ok := stage.PerpendicularVectorHalfways[perpendicularvectorhalfway]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPerpendicularVectorHalfway(perpendicularvectorhalfway)
		}
	}
	return perpendicularvectorhalfway
}

// for satisfaction of GongStruct interface
func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GetName() (res string) {
	return perpendicularvectorhalfway.Name
}

// for satisfaction of GongStruct interface
func (perpendicularvectorhalfway *PerpendicularVectorHalfway) SetName(name string) {
	perpendicularvectorhalfway.Name = name
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

// Stage puts rendered3dshape to the model stage
func (rendered3dshape *Rendered3DShape) Stage(stage *Stage) *Rendered3DShape {
	if _, ok := stage.Rendered3DShapes[rendered3dshape]; !ok {
		stage.Rendered3DShapes[rendered3dshape] = struct{}{}
		stage.Rendered3DShape_stagedOrder[rendered3dshape] = stage.Rendered3DShapeOrder
		stage.Rendered3DShape_orderStaged[stage.Rendered3DShapeOrder] = rendered3dshape
		stage.Rendered3DShapeOrder++
	}
	stage.Rendered3DShapes_mapString[rendered3dshape.Name] = rendered3dshape

	return rendered3dshape
}

// StagePreserveOrder puts rendered3dshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.Rendered3DShapeOrder
// - update stage.Rendered3DShapeOrder accordingly
func (rendered3dshape *Rendered3DShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Rendered3DShapes[rendered3dshape]; !ok {
		stage.Rendered3DShapes[rendered3dshape] = struct{}{}

		if order > stage.Rendered3DShapeOrder {
			stage.Rendered3DShapeOrder = order
		}
		stage.Rendered3DShape_stagedOrder[rendered3dshape] = order
		stage.Rendered3DShape_orderStaged[order] = rendered3dshape
		stage.Rendered3DShapeOrder++
	}
	stage.Rendered3DShapes_mapString[rendered3dshape.Name] = rendered3dshape
}

// Unstage removes rendered3dshape off the model stage
func (rendered3dshape *Rendered3DShape) Unstage(stage *Stage) *Rendered3DShape {
	delete(stage.Rendered3DShapes, rendered3dshape)
	// issue1150
	// delete(stage.Rendered3DShape_stagedOrder, rendered3dshape)
	delete(stage.Rendered3DShapes_mapString, rendered3dshape.Name)

	return rendered3dshape
}

// UnstageVoid removes rendered3dshape off the model stage
func (rendered3dshape *Rendered3DShape) UnstageVoid(stage *Stage) {
	delete(stage.Rendered3DShapes, rendered3dshape)
	// issue1150
	// delete(stage.Rendered3DShape_stagedOrder, rendered3dshape)
	delete(stage.Rendered3DShapes_mapString, rendered3dshape.Name)
}

// commit rendered3dshape to the back repo (if it is already staged)
func (rendered3dshape *Rendered3DShape) Commit(stage *Stage) *Rendered3DShape {
	if _, ok := stage.Rendered3DShapes[rendered3dshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRendered3DShape(rendered3dshape)
		}
	}
	return rendered3dshape
}

func (rendered3dshape *Rendered3DShape) CommitVoid(stage *Stage) {
	rendered3dshape.Commit(stage)
}

func (rendered3dshape *Rendered3DShape) StageVoid(stage *Stage) {
	rendered3dshape.Stage(stage)
}

// Checkout rendered3dshape to the back repo (if it is already staged)
func (rendered3dshape *Rendered3DShape) Checkout(stage *Stage) *Rendered3DShape {
	if _, ok := stage.Rendered3DShapes[rendered3dshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRendered3DShape(rendered3dshape)
		}
	}
	return rendered3dshape
}

// for satisfaction of GongStruct interface
func (rendered3dshape *Rendered3DShape) GetName() (res string) {
	return rendered3dshape.Name
}

// for satisfaction of GongStruct interface
func (rendered3dshape *Rendered3DShape) SetName(name string) {
	rendered3dshape.Name = name
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

// Stage puts stackgrowthcurvebeziershape to the model stage
func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) Stage(stage *Stage) *StackGrowthCurveBezierShape {
	if _, ok := stage.StackGrowthCurveBezierShapes[stackgrowthcurvebeziershape]; !ok {
		stage.StackGrowthCurveBezierShapes[stackgrowthcurvebeziershape] = struct{}{}
		stage.StackGrowthCurveBezierShape_stagedOrder[stackgrowthcurvebeziershape] = stage.StackGrowthCurveBezierShapeOrder
		stage.StackGrowthCurveBezierShape_orderStaged[stage.StackGrowthCurveBezierShapeOrder] = stackgrowthcurvebeziershape
		stage.StackGrowthCurveBezierShapeOrder++
	}
	stage.StackGrowthCurveBezierShapes_mapString[stackgrowthcurvebeziershape.Name] = stackgrowthcurvebeziershape

	return stackgrowthcurvebeziershape
}

// StagePreserveOrder puts stackgrowthcurvebeziershape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StackGrowthCurveBezierShapeOrder
// - update stage.StackGrowthCurveBezierShapeOrder accordingly
func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.StackGrowthCurveBezierShapes[stackgrowthcurvebeziershape]; !ok {
		stage.StackGrowthCurveBezierShapes[stackgrowthcurvebeziershape] = struct{}{}

		if order > stage.StackGrowthCurveBezierShapeOrder {
			stage.StackGrowthCurveBezierShapeOrder = order
		}
		stage.StackGrowthCurveBezierShape_stagedOrder[stackgrowthcurvebeziershape] = order
		stage.StackGrowthCurveBezierShape_orderStaged[order] = stackgrowthcurvebeziershape
		stage.StackGrowthCurveBezierShapeOrder++
	}
	stage.StackGrowthCurveBezierShapes_mapString[stackgrowthcurvebeziershape.Name] = stackgrowthcurvebeziershape
}

// Unstage removes stackgrowthcurvebeziershape off the model stage
func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) Unstage(stage *Stage) *StackGrowthCurveBezierShape {
	delete(stage.StackGrowthCurveBezierShapes, stackgrowthcurvebeziershape)
	// issue1150
	// delete(stage.StackGrowthCurveBezierShape_stagedOrder, stackgrowthcurvebeziershape)
	delete(stage.StackGrowthCurveBezierShapes_mapString, stackgrowthcurvebeziershape.Name)

	return stackgrowthcurvebeziershape
}

// UnstageVoid removes stackgrowthcurvebeziershape off the model stage
func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) UnstageVoid(stage *Stage) {
	delete(stage.StackGrowthCurveBezierShapes, stackgrowthcurvebeziershape)
	// issue1150
	// delete(stage.StackGrowthCurveBezierShape_stagedOrder, stackgrowthcurvebeziershape)
	delete(stage.StackGrowthCurveBezierShapes_mapString, stackgrowthcurvebeziershape.Name)
}

// commit stackgrowthcurvebeziershape to the back repo (if it is already staged)
func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) Commit(stage *Stage) *StackGrowthCurveBezierShape {
	if _, ok := stage.StackGrowthCurveBezierShapes[stackgrowthcurvebeziershape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitStackGrowthCurveBezierShape(stackgrowthcurvebeziershape)
		}
	}
	return stackgrowthcurvebeziershape
}

func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) CommitVoid(stage *Stage) {
	stackgrowthcurvebeziershape.Commit(stage)
}

func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) StageVoid(stage *Stage) {
	stackgrowthcurvebeziershape.Stage(stage)
}

// Checkout stackgrowthcurvebeziershape to the back repo (if it is already staged)
func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) Checkout(stage *Stage) *StackGrowthCurveBezierShape {
	if _, ok := stage.StackGrowthCurveBezierShapes[stackgrowthcurvebeziershape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutStackGrowthCurveBezierShape(stackgrowthcurvebeziershape)
		}
	}
	return stackgrowthcurvebeziershape
}

// for satisfaction of GongStruct interface
func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) GetName() (res string) {
	return stackgrowthcurvebeziershape.Name
}

// for satisfaction of GongStruct interface
func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) SetName(name string) {
	stackgrowthcurvebeziershape.Name = name
}

// Stage puts stackofgrowthcurve to the model stage
func (stackofgrowthcurve *StackOfGrowthCurve) Stage(stage *Stage) *StackOfGrowthCurve {
	if _, ok := stage.StackOfGrowthCurves[stackofgrowthcurve]; !ok {
		stage.StackOfGrowthCurves[stackofgrowthcurve] = struct{}{}
		stage.StackOfGrowthCurve_stagedOrder[stackofgrowthcurve] = stage.StackOfGrowthCurveOrder
		stage.StackOfGrowthCurve_orderStaged[stage.StackOfGrowthCurveOrder] = stackofgrowthcurve
		stage.StackOfGrowthCurveOrder++
	}
	stage.StackOfGrowthCurves_mapString[stackofgrowthcurve.Name] = stackofgrowthcurve

	return stackofgrowthcurve
}

// StagePreserveOrder puts stackofgrowthcurve to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StackOfGrowthCurveOrder
// - update stage.StackOfGrowthCurveOrder accordingly
func (stackofgrowthcurve *StackOfGrowthCurve) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.StackOfGrowthCurves[stackofgrowthcurve]; !ok {
		stage.StackOfGrowthCurves[stackofgrowthcurve] = struct{}{}

		if order > stage.StackOfGrowthCurveOrder {
			stage.StackOfGrowthCurveOrder = order
		}
		stage.StackOfGrowthCurve_stagedOrder[stackofgrowthcurve] = order
		stage.StackOfGrowthCurve_orderStaged[order] = stackofgrowthcurve
		stage.StackOfGrowthCurveOrder++
	}
	stage.StackOfGrowthCurves_mapString[stackofgrowthcurve.Name] = stackofgrowthcurve
}

// Unstage removes stackofgrowthcurve off the model stage
func (stackofgrowthcurve *StackOfGrowthCurve) Unstage(stage *Stage) *StackOfGrowthCurve {
	delete(stage.StackOfGrowthCurves, stackofgrowthcurve)
	// issue1150
	// delete(stage.StackOfGrowthCurve_stagedOrder, stackofgrowthcurve)
	delete(stage.StackOfGrowthCurves_mapString, stackofgrowthcurve.Name)

	return stackofgrowthcurve
}

// UnstageVoid removes stackofgrowthcurve off the model stage
func (stackofgrowthcurve *StackOfGrowthCurve) UnstageVoid(stage *Stage) {
	delete(stage.StackOfGrowthCurves, stackofgrowthcurve)
	// issue1150
	// delete(stage.StackOfGrowthCurve_stagedOrder, stackofgrowthcurve)
	delete(stage.StackOfGrowthCurves_mapString, stackofgrowthcurve.Name)
}

// commit stackofgrowthcurve to the back repo (if it is already staged)
func (stackofgrowthcurve *StackOfGrowthCurve) Commit(stage *Stage) *StackOfGrowthCurve {
	if _, ok := stage.StackOfGrowthCurves[stackofgrowthcurve]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitStackOfGrowthCurve(stackofgrowthcurve)
		}
	}
	return stackofgrowthcurve
}

func (stackofgrowthcurve *StackOfGrowthCurve) CommitVoid(stage *Stage) {
	stackofgrowthcurve.Commit(stage)
}

func (stackofgrowthcurve *StackOfGrowthCurve) StageVoid(stage *Stage) {
	stackofgrowthcurve.Stage(stage)
}

// Checkout stackofgrowthcurve to the back repo (if it is already staged)
func (stackofgrowthcurve *StackOfGrowthCurve) Checkout(stage *Stage) *StackOfGrowthCurve {
	if _, ok := stage.StackOfGrowthCurves[stackofgrowthcurve]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutStackOfGrowthCurve(stackofgrowthcurve)
		}
	}
	return stackofgrowthcurve
}

// for satisfaction of GongStruct interface
func (stackofgrowthcurve *StackOfGrowthCurve) GetName() (res string) {
	return stackofgrowthcurve.Name
}

// for satisfaction of GongStruct interface
func (stackofgrowthcurve *StackOfGrowthCurve) SetName(name string) {
	stackofgrowthcurve.Name = name
}

// Stage puts startarcshape to the model stage
func (startarcshape *StartArcShape) Stage(stage *Stage) *StartArcShape {
	if _, ok := stage.StartArcShapes[startarcshape]; !ok {
		stage.StartArcShapes[startarcshape] = struct{}{}
		stage.StartArcShape_stagedOrder[startarcshape] = stage.StartArcShapeOrder
		stage.StartArcShape_orderStaged[stage.StartArcShapeOrder] = startarcshape
		stage.StartArcShapeOrder++
	}
	stage.StartArcShapes_mapString[startarcshape.Name] = startarcshape

	return startarcshape
}

// StagePreserveOrder puts startarcshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StartArcShapeOrder
// - update stage.StartArcShapeOrder accordingly
func (startarcshape *StartArcShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.StartArcShapes[startarcshape]; !ok {
		stage.StartArcShapes[startarcshape] = struct{}{}

		if order > stage.StartArcShapeOrder {
			stage.StartArcShapeOrder = order
		}
		stage.StartArcShape_stagedOrder[startarcshape] = order
		stage.StartArcShape_orderStaged[order] = startarcshape
		stage.StartArcShapeOrder++
	}
	stage.StartArcShapes_mapString[startarcshape.Name] = startarcshape
}

// Unstage removes startarcshape off the model stage
func (startarcshape *StartArcShape) Unstage(stage *Stage) *StartArcShape {
	delete(stage.StartArcShapes, startarcshape)
	// issue1150
	// delete(stage.StartArcShape_stagedOrder, startarcshape)
	delete(stage.StartArcShapes_mapString, startarcshape.Name)

	return startarcshape
}

// UnstageVoid removes startarcshape off the model stage
func (startarcshape *StartArcShape) UnstageVoid(stage *Stage) {
	delete(stage.StartArcShapes, startarcshape)
	// issue1150
	// delete(stage.StartArcShape_stagedOrder, startarcshape)
	delete(stage.StartArcShapes_mapString, startarcshape.Name)
}

// commit startarcshape to the back repo (if it is already staged)
func (startarcshape *StartArcShape) Commit(stage *Stage) *StartArcShape {
	if _, ok := stage.StartArcShapes[startarcshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitStartArcShape(startarcshape)
		}
	}
	return startarcshape
}

func (startarcshape *StartArcShape) CommitVoid(stage *Stage) {
	startarcshape.Commit(stage)
}

func (startarcshape *StartArcShape) StageVoid(stage *Stage) {
	startarcshape.Stage(stage)
}

// Checkout startarcshape to the back repo (if it is already staged)
func (startarcshape *StartArcShape) Checkout(stage *Stage) *StartArcShape {
	if _, ok := stage.StartArcShapes[startarcshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutStartArcShape(startarcshape)
		}
	}
	return startarcshape
}

// for satisfaction of GongStruct interface
func (startarcshape *StartArcShape) GetName() (res string) {
	return startarcshape.Name
}

// for satisfaction of GongStruct interface
func (startarcshape *StartArcShape) SetName(name string) {
	startarcshape.Name = name
}

// Stage puts startarcshapegrid to the model stage
func (startarcshapegrid *StartArcShapeGrid) Stage(stage *Stage) *StartArcShapeGrid {
	if _, ok := stage.StartArcShapeGrids[startarcshapegrid]; !ok {
		stage.StartArcShapeGrids[startarcshapegrid] = struct{}{}
		stage.StartArcShapeGrid_stagedOrder[startarcshapegrid] = stage.StartArcShapeGridOrder
		stage.StartArcShapeGrid_orderStaged[stage.StartArcShapeGridOrder] = startarcshapegrid
		stage.StartArcShapeGridOrder++
	}
	stage.StartArcShapeGrids_mapString[startarcshapegrid.Name] = startarcshapegrid

	return startarcshapegrid
}

// StagePreserveOrder puts startarcshapegrid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StartArcShapeGridOrder
// - update stage.StartArcShapeGridOrder accordingly
func (startarcshapegrid *StartArcShapeGrid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.StartArcShapeGrids[startarcshapegrid]; !ok {
		stage.StartArcShapeGrids[startarcshapegrid] = struct{}{}

		if order > stage.StartArcShapeGridOrder {
			stage.StartArcShapeGridOrder = order
		}
		stage.StartArcShapeGrid_stagedOrder[startarcshapegrid] = order
		stage.StartArcShapeGrid_orderStaged[order] = startarcshapegrid
		stage.StartArcShapeGridOrder++
	}
	stage.StartArcShapeGrids_mapString[startarcshapegrid.Name] = startarcshapegrid
}

// Unstage removes startarcshapegrid off the model stage
func (startarcshapegrid *StartArcShapeGrid) Unstage(stage *Stage) *StartArcShapeGrid {
	delete(stage.StartArcShapeGrids, startarcshapegrid)
	// issue1150
	// delete(stage.StartArcShapeGrid_stagedOrder, startarcshapegrid)
	delete(stage.StartArcShapeGrids_mapString, startarcshapegrid.Name)

	return startarcshapegrid
}

// UnstageVoid removes startarcshapegrid off the model stage
func (startarcshapegrid *StartArcShapeGrid) UnstageVoid(stage *Stage) {
	delete(stage.StartArcShapeGrids, startarcshapegrid)
	// issue1150
	// delete(stage.StartArcShapeGrid_stagedOrder, startarcshapegrid)
	delete(stage.StartArcShapeGrids_mapString, startarcshapegrid.Name)
}

// commit startarcshapegrid to the back repo (if it is already staged)
func (startarcshapegrid *StartArcShapeGrid) Commit(stage *Stage) *StartArcShapeGrid {
	if _, ok := stage.StartArcShapeGrids[startarcshapegrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitStartArcShapeGrid(startarcshapegrid)
		}
	}
	return startarcshapegrid
}

func (startarcshapegrid *StartArcShapeGrid) CommitVoid(stage *Stage) {
	startarcshapegrid.Commit(stage)
}

func (startarcshapegrid *StartArcShapeGrid) StageVoid(stage *Stage) {
	startarcshapegrid.Stage(stage)
}

// Checkout startarcshapegrid to the back repo (if it is already staged)
func (startarcshapegrid *StartArcShapeGrid) Checkout(stage *Stage) *StartArcShapeGrid {
	if _, ok := stage.StartArcShapeGrids[startarcshapegrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutStartArcShapeGrid(startarcshapegrid)
		}
	}
	return startarcshapegrid
}

// for satisfaction of GongStruct interface
func (startarcshapegrid *StartArcShapeGrid) GetName() (res string) {
	return startarcshapegrid.Name
}

// for satisfaction of GongStruct interface
func (startarcshapegrid *StartArcShapeGrid) SetName(name string) {
	startarcshapegrid.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMAxesShape(AxesShape *AxesShape)
	CreateORMCircleGridShape(CircleGridShape *CircleGridShape)
	CreateORMEndArcShape(EndArcShape *EndArcShape)
	CreateORMEndArcShapeGrid(EndArcShapeGrid *EndArcShapeGrid)
	CreateORMExplanationTextShape(ExplanationTextShape *ExplanationTextShape)
	CreateORMGridPathShape(GridPathShape *GridPathShape)
	CreateORMGrowthCurveBezierShape(GrowthCurveBezierShape *GrowthCurveBezierShape)
	CreateORMGrowthCurveBezierShapeGrid(GrowthCurveBezierShapeGrid *GrowthCurveBezierShapeGrid)
	CreateORMGrowthCurveRhombusGridShape(GrowthCurveRhombusGridShape *GrowthCurveRhombusGridShape)
	CreateORMGrowthCurveRhombusShape(GrowthCurveRhombusShape *GrowthCurveRhombusShape)
	CreateORMGrowthVectorShape(GrowthVectorShape *GrowthVectorShape)
	CreateORMInitialRhombusGridShape(InitialRhombusGridShape *InitialRhombusGridShape)
	CreateORMInitialRhombusShape(InitialRhombusShape *InitialRhombusShape)
	CreateORMLibrary(Library *Library)
	CreateORMNextCircleShape(NextCircleShape *NextCircleShape)
	CreateORMPerpendicularVector(PerpendicularVector *PerpendicularVector)
	CreateORMPerpendicularVectorGrid(PerpendicularVectorGrid *PerpendicularVectorGrid)
	CreateORMPerpendicularVectorGridHalfway(PerpendicularVectorGridHalfway *PerpendicularVectorGridHalfway)
	CreateORMPerpendicularVectorHalfway(PerpendicularVectorHalfway *PerpendicularVectorHalfway)
	CreateORMPlant(Plant *Plant)
	CreateORMPlantCircumferenceShape(PlantCircumferenceShape *PlantCircumferenceShape)
	CreateORMPlantDiagram(PlantDiagram *PlantDiagram)
	CreateORMRendered3DShape(Rendered3DShape *Rendered3DShape)
	CreateORMRhombusShape(RhombusShape *RhombusShape)
	CreateORMRotatedRhombusGridShape(RotatedRhombusGridShape *RotatedRhombusGridShape)
	CreateORMRotatedRhombusShape(RotatedRhombusShape *RotatedRhombusShape)
	CreateORMStackGrowthCurveBezierShape(StackGrowthCurveBezierShape *StackGrowthCurveBezierShape)
	CreateORMStackOfGrowthCurve(StackOfGrowthCurve *StackOfGrowthCurve)
	CreateORMStartArcShape(StartArcShape *StartArcShape)
	CreateORMStartArcShapeGrid(StartArcShapeGrid *StartArcShapeGrid)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAxesShape(AxesShape *AxesShape)
	DeleteORMCircleGridShape(CircleGridShape *CircleGridShape)
	DeleteORMEndArcShape(EndArcShape *EndArcShape)
	DeleteORMEndArcShapeGrid(EndArcShapeGrid *EndArcShapeGrid)
	DeleteORMExplanationTextShape(ExplanationTextShape *ExplanationTextShape)
	DeleteORMGridPathShape(GridPathShape *GridPathShape)
	DeleteORMGrowthCurveBezierShape(GrowthCurveBezierShape *GrowthCurveBezierShape)
	DeleteORMGrowthCurveBezierShapeGrid(GrowthCurveBezierShapeGrid *GrowthCurveBezierShapeGrid)
	DeleteORMGrowthCurveRhombusGridShape(GrowthCurveRhombusGridShape *GrowthCurveRhombusGridShape)
	DeleteORMGrowthCurveRhombusShape(GrowthCurveRhombusShape *GrowthCurveRhombusShape)
	DeleteORMGrowthVectorShape(GrowthVectorShape *GrowthVectorShape)
	DeleteORMInitialRhombusGridShape(InitialRhombusGridShape *InitialRhombusGridShape)
	DeleteORMInitialRhombusShape(InitialRhombusShape *InitialRhombusShape)
	DeleteORMLibrary(Library *Library)
	DeleteORMNextCircleShape(NextCircleShape *NextCircleShape)
	DeleteORMPerpendicularVector(PerpendicularVector *PerpendicularVector)
	DeleteORMPerpendicularVectorGrid(PerpendicularVectorGrid *PerpendicularVectorGrid)
	DeleteORMPerpendicularVectorGridHalfway(PerpendicularVectorGridHalfway *PerpendicularVectorGridHalfway)
	DeleteORMPerpendicularVectorHalfway(PerpendicularVectorHalfway *PerpendicularVectorHalfway)
	DeleteORMPlant(Plant *Plant)
	DeleteORMPlantCircumferenceShape(PlantCircumferenceShape *PlantCircumferenceShape)
	DeleteORMPlantDiagram(PlantDiagram *PlantDiagram)
	DeleteORMRendered3DShape(Rendered3DShape *Rendered3DShape)
	DeleteORMRhombusShape(RhombusShape *RhombusShape)
	DeleteORMRotatedRhombusGridShape(RotatedRhombusGridShape *RotatedRhombusGridShape)
	DeleteORMRotatedRhombusShape(RotatedRhombusShape *RotatedRhombusShape)
	DeleteORMStackGrowthCurveBezierShape(StackGrowthCurveBezierShape *StackGrowthCurveBezierShape)
	DeleteORMStackOfGrowthCurve(StackOfGrowthCurve *StackOfGrowthCurve)
	DeleteORMStartArcShape(StartArcShape *StartArcShape)
	DeleteORMStartArcShapeGrid(StartArcShapeGrid *StartArcShapeGrid)
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

	stage.EndArcShapes = make(map[*EndArcShape]struct{})
	stage.EndArcShapes_mapString = make(map[string]*EndArcShape)
	stage.EndArcShape_stagedOrder = make(map[*EndArcShape]uint)
	stage.EndArcShapeOrder = 0

	stage.EndArcShapeGrids = make(map[*EndArcShapeGrid]struct{})
	stage.EndArcShapeGrids_mapString = make(map[string]*EndArcShapeGrid)
	stage.EndArcShapeGrid_stagedOrder = make(map[*EndArcShapeGrid]uint)
	stage.EndArcShapeGridOrder = 0

	stage.ExplanationTextShapes = make(map[*ExplanationTextShape]struct{})
	stage.ExplanationTextShapes_mapString = make(map[string]*ExplanationTextShape)
	stage.ExplanationTextShape_stagedOrder = make(map[*ExplanationTextShape]uint)
	stage.ExplanationTextShapeOrder = 0

	stage.GridPathShapes = make(map[*GridPathShape]struct{})
	stage.GridPathShapes_mapString = make(map[string]*GridPathShape)
	stage.GridPathShape_stagedOrder = make(map[*GridPathShape]uint)
	stage.GridPathShapeOrder = 0

	stage.GrowthCurveBezierShapes = make(map[*GrowthCurveBezierShape]struct{})
	stage.GrowthCurveBezierShapes_mapString = make(map[string]*GrowthCurveBezierShape)
	stage.GrowthCurveBezierShape_stagedOrder = make(map[*GrowthCurveBezierShape]uint)
	stage.GrowthCurveBezierShapeOrder = 0

	stage.GrowthCurveBezierShapeGrids = make(map[*GrowthCurveBezierShapeGrid]struct{})
	stage.GrowthCurveBezierShapeGrids_mapString = make(map[string]*GrowthCurveBezierShapeGrid)
	stage.GrowthCurveBezierShapeGrid_stagedOrder = make(map[*GrowthCurveBezierShapeGrid]uint)
	stage.GrowthCurveBezierShapeGridOrder = 0

	stage.GrowthCurveRhombusGridShapes = make(map[*GrowthCurveRhombusGridShape]struct{})
	stage.GrowthCurveRhombusGridShapes_mapString = make(map[string]*GrowthCurveRhombusGridShape)
	stage.GrowthCurveRhombusGridShape_stagedOrder = make(map[*GrowthCurveRhombusGridShape]uint)
	stage.GrowthCurveRhombusGridShapeOrder = 0

	stage.GrowthCurveRhombusShapes = make(map[*GrowthCurveRhombusShape]struct{})
	stage.GrowthCurveRhombusShapes_mapString = make(map[string]*GrowthCurveRhombusShape)
	stage.GrowthCurveRhombusShape_stagedOrder = make(map[*GrowthCurveRhombusShape]uint)
	stage.GrowthCurveRhombusShapeOrder = 0

	stage.GrowthVectorShapes = make(map[*GrowthVectorShape]struct{})
	stage.GrowthVectorShapes_mapString = make(map[string]*GrowthVectorShape)
	stage.GrowthVectorShape_stagedOrder = make(map[*GrowthVectorShape]uint)
	stage.GrowthVectorShapeOrder = 0

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

	stage.PerpendicularVectors = make(map[*PerpendicularVector]struct{})
	stage.PerpendicularVectors_mapString = make(map[string]*PerpendicularVector)
	stage.PerpendicularVector_stagedOrder = make(map[*PerpendicularVector]uint)
	stage.PerpendicularVectorOrder = 0

	stage.PerpendicularVectorGrids = make(map[*PerpendicularVectorGrid]struct{})
	stage.PerpendicularVectorGrids_mapString = make(map[string]*PerpendicularVectorGrid)
	stage.PerpendicularVectorGrid_stagedOrder = make(map[*PerpendicularVectorGrid]uint)
	stage.PerpendicularVectorGridOrder = 0

	stage.PerpendicularVectorGridHalfways = make(map[*PerpendicularVectorGridHalfway]struct{})
	stage.PerpendicularVectorGridHalfways_mapString = make(map[string]*PerpendicularVectorGridHalfway)
	stage.PerpendicularVectorGridHalfway_stagedOrder = make(map[*PerpendicularVectorGridHalfway]uint)
	stage.PerpendicularVectorGridHalfwayOrder = 0

	stage.PerpendicularVectorHalfways = make(map[*PerpendicularVectorHalfway]struct{})
	stage.PerpendicularVectorHalfways_mapString = make(map[string]*PerpendicularVectorHalfway)
	stage.PerpendicularVectorHalfway_stagedOrder = make(map[*PerpendicularVectorHalfway]uint)
	stage.PerpendicularVectorHalfwayOrder = 0

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

	stage.Rendered3DShapes = make(map[*Rendered3DShape]struct{})
	stage.Rendered3DShapes_mapString = make(map[string]*Rendered3DShape)
	stage.Rendered3DShape_stagedOrder = make(map[*Rendered3DShape]uint)
	stage.Rendered3DShapeOrder = 0

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

	stage.StackGrowthCurveBezierShapes = make(map[*StackGrowthCurveBezierShape]struct{})
	stage.StackGrowthCurveBezierShapes_mapString = make(map[string]*StackGrowthCurveBezierShape)
	stage.StackGrowthCurveBezierShape_stagedOrder = make(map[*StackGrowthCurveBezierShape]uint)
	stage.StackGrowthCurveBezierShapeOrder = 0

	stage.StackOfGrowthCurves = make(map[*StackOfGrowthCurve]struct{})
	stage.StackOfGrowthCurves_mapString = make(map[string]*StackOfGrowthCurve)
	stage.StackOfGrowthCurve_stagedOrder = make(map[*StackOfGrowthCurve]uint)
	stage.StackOfGrowthCurveOrder = 0

	stage.StartArcShapes = make(map[*StartArcShape]struct{})
	stage.StartArcShapes_mapString = make(map[string]*StartArcShape)
	stage.StartArcShape_stagedOrder = make(map[*StartArcShape]uint)
	stage.StartArcShapeOrder = 0

	stage.StartArcShapeGrids = make(map[*StartArcShapeGrid]struct{})
	stage.StartArcShapeGrids_mapString = make(map[string]*StartArcShapeGrid)
	stage.StartArcShapeGrid_stagedOrder = make(map[*StartArcShapeGrid]uint)
	stage.StartArcShapeGridOrder = 0

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

	stage.EndArcShapes = nil
	stage.EndArcShapes_mapString = nil

	stage.EndArcShapeGrids = nil
	stage.EndArcShapeGrids_mapString = nil

	stage.ExplanationTextShapes = nil
	stage.ExplanationTextShapes_mapString = nil

	stage.GridPathShapes = nil
	stage.GridPathShapes_mapString = nil

	stage.GrowthCurveBezierShapes = nil
	stage.GrowthCurveBezierShapes_mapString = nil

	stage.GrowthCurveBezierShapeGrids = nil
	stage.GrowthCurveBezierShapeGrids_mapString = nil

	stage.GrowthCurveRhombusGridShapes = nil
	stage.GrowthCurveRhombusGridShapes_mapString = nil

	stage.GrowthCurveRhombusShapes = nil
	stage.GrowthCurveRhombusShapes_mapString = nil

	stage.GrowthVectorShapes = nil
	stage.GrowthVectorShapes_mapString = nil

	stage.InitialRhombusGridShapes = nil
	stage.InitialRhombusGridShapes_mapString = nil

	stage.InitialRhombusShapes = nil
	stage.InitialRhombusShapes_mapString = nil

	stage.Librarys = nil
	stage.Librarys_mapString = nil

	stage.NextCircleShapes = nil
	stage.NextCircleShapes_mapString = nil

	stage.PerpendicularVectors = nil
	stage.PerpendicularVectors_mapString = nil

	stage.PerpendicularVectorGrids = nil
	stage.PerpendicularVectorGrids_mapString = nil

	stage.PerpendicularVectorGridHalfways = nil
	stage.PerpendicularVectorGridHalfways_mapString = nil

	stage.PerpendicularVectorHalfways = nil
	stage.PerpendicularVectorHalfways_mapString = nil

	stage.Plants = nil
	stage.Plants_mapString = nil

	stage.PlantCircumferenceShapes = nil
	stage.PlantCircumferenceShapes_mapString = nil

	stage.PlantDiagrams = nil
	stage.PlantDiagrams_mapString = nil

	stage.Rendered3DShapes = nil
	stage.Rendered3DShapes_mapString = nil

	stage.RhombusShapes = nil
	stage.RhombusShapes_mapString = nil

	stage.RotatedRhombusGridShapes = nil
	stage.RotatedRhombusGridShapes_mapString = nil

	stage.RotatedRhombusShapes = nil
	stage.RotatedRhombusShapes_mapString = nil

	stage.StackGrowthCurveBezierShapes = nil
	stage.StackGrowthCurveBezierShapes_mapString = nil

	stage.StackOfGrowthCurves = nil
	stage.StackOfGrowthCurves_mapString = nil

	stage.StartArcShapes = nil
	stage.StartArcShapes_mapString = nil

	stage.StartArcShapeGrids = nil
	stage.StartArcShapeGrids_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for axesshape := range stage.AxesShapes {
		axesshape.Unstage(stage)
	}

	for circlegridshape := range stage.CircleGridShapes {
		circlegridshape.Unstage(stage)
	}

	for endarcshape := range stage.EndArcShapes {
		endarcshape.Unstage(stage)
	}

	for endarcshapegrid := range stage.EndArcShapeGrids {
		endarcshapegrid.Unstage(stage)
	}

	for explanationtextshape := range stage.ExplanationTextShapes {
		explanationtextshape.Unstage(stage)
	}

	for gridpathshape := range stage.GridPathShapes {
		gridpathshape.Unstage(stage)
	}

	for growthcurvebeziershape := range stage.GrowthCurveBezierShapes {
		growthcurvebeziershape.Unstage(stage)
	}

	for growthcurvebeziershapegrid := range stage.GrowthCurveBezierShapeGrids {
		growthcurvebeziershapegrid.Unstage(stage)
	}

	for growthcurverhombusgridshape := range stage.GrowthCurveRhombusGridShapes {
		growthcurverhombusgridshape.Unstage(stage)
	}

	for growthcurverhombusshape := range stage.GrowthCurveRhombusShapes {
		growthcurverhombusshape.Unstage(stage)
	}

	for growthvectorshape := range stage.GrowthVectorShapes {
		growthvectorshape.Unstage(stage)
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

	for perpendicularvector := range stage.PerpendicularVectors {
		perpendicularvector.Unstage(stage)
	}

	for perpendicularvectorgrid := range stage.PerpendicularVectorGrids {
		perpendicularvectorgrid.Unstage(stage)
	}

	for perpendicularvectorgridhalfway := range stage.PerpendicularVectorGridHalfways {
		perpendicularvectorgridhalfway.Unstage(stage)
	}

	for perpendicularvectorhalfway := range stage.PerpendicularVectorHalfways {
		perpendicularvectorhalfway.Unstage(stage)
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

	for rendered3dshape := range stage.Rendered3DShapes {
		rendered3dshape.Unstage(stage)
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

	for stackgrowthcurvebeziershape := range stage.StackGrowthCurveBezierShapes {
		stackgrowthcurvebeziershape.Unstage(stage)
	}

	for stackofgrowthcurve := range stage.StackOfGrowthCurves {
		stackofgrowthcurve.Unstage(stage)
	}

	for startarcshape := range stage.StartArcShapes {
		startarcshape.Unstage(stage)
	}

	for startarcshapegrid := range stage.StartArcShapeGrids {
		startarcshapegrid.Unstage(stage)
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
	case map[*EndArcShape]any:
		return any(&stage.EndArcShapes).(*Type)
	case map[*EndArcShapeGrid]any:
		return any(&stage.EndArcShapeGrids).(*Type)
	case map[*ExplanationTextShape]any:
		return any(&stage.ExplanationTextShapes).(*Type)
	case map[*GridPathShape]any:
		return any(&stage.GridPathShapes).(*Type)
	case map[*GrowthCurveBezierShape]any:
		return any(&stage.GrowthCurveBezierShapes).(*Type)
	case map[*GrowthCurveBezierShapeGrid]any:
		return any(&stage.GrowthCurveBezierShapeGrids).(*Type)
	case map[*GrowthCurveRhombusGridShape]any:
		return any(&stage.GrowthCurveRhombusGridShapes).(*Type)
	case map[*GrowthCurveRhombusShape]any:
		return any(&stage.GrowthCurveRhombusShapes).(*Type)
	case map[*GrowthVectorShape]any:
		return any(&stage.GrowthVectorShapes).(*Type)
	case map[*InitialRhombusGridShape]any:
		return any(&stage.InitialRhombusGridShapes).(*Type)
	case map[*InitialRhombusShape]any:
		return any(&stage.InitialRhombusShapes).(*Type)
	case map[*Library]any:
		return any(&stage.Librarys).(*Type)
	case map[*NextCircleShape]any:
		return any(&stage.NextCircleShapes).(*Type)
	case map[*PerpendicularVector]any:
		return any(&stage.PerpendicularVectors).(*Type)
	case map[*PerpendicularVectorGrid]any:
		return any(&stage.PerpendicularVectorGrids).(*Type)
	case map[*PerpendicularVectorGridHalfway]any:
		return any(&stage.PerpendicularVectorGridHalfways).(*Type)
	case map[*PerpendicularVectorHalfway]any:
		return any(&stage.PerpendicularVectorHalfways).(*Type)
	case map[*Plant]any:
		return any(&stage.Plants).(*Type)
	case map[*PlantCircumferenceShape]any:
		return any(&stage.PlantCircumferenceShapes).(*Type)
	case map[*PlantDiagram]any:
		return any(&stage.PlantDiagrams).(*Type)
	case map[*Rendered3DShape]any:
		return any(&stage.Rendered3DShapes).(*Type)
	case map[*RhombusShape]any:
		return any(&stage.RhombusShapes).(*Type)
	case map[*RotatedRhombusGridShape]any:
		return any(&stage.RotatedRhombusGridShapes).(*Type)
	case map[*RotatedRhombusShape]any:
		return any(&stage.RotatedRhombusShapes).(*Type)
	case map[*StackGrowthCurveBezierShape]any:
		return any(&stage.StackGrowthCurveBezierShapes).(*Type)
	case map[*StackOfGrowthCurve]any:
		return any(&stage.StackOfGrowthCurves).(*Type)
	case map[*StartArcShape]any:
		return any(&stage.StartArcShapes).(*Type)
	case map[*StartArcShapeGrid]any:
		return any(&stage.StartArcShapeGrids).(*Type)
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
	case *EndArcShape:
		return any(stage.EndArcShapes_mapString).(map[string]Type)
	case *EndArcShapeGrid:
		return any(stage.EndArcShapeGrids_mapString).(map[string]Type)
	case *ExplanationTextShape:
		return any(stage.ExplanationTextShapes_mapString).(map[string]Type)
	case *GridPathShape:
		return any(stage.GridPathShapes_mapString).(map[string]Type)
	case *GrowthCurveBezierShape:
		return any(stage.GrowthCurveBezierShapes_mapString).(map[string]Type)
	case *GrowthCurveBezierShapeGrid:
		return any(stage.GrowthCurveBezierShapeGrids_mapString).(map[string]Type)
	case *GrowthCurveRhombusGridShape:
		return any(stage.GrowthCurveRhombusGridShapes_mapString).(map[string]Type)
	case *GrowthCurveRhombusShape:
		return any(stage.GrowthCurveRhombusShapes_mapString).(map[string]Type)
	case *GrowthVectorShape:
		return any(stage.GrowthVectorShapes_mapString).(map[string]Type)
	case *InitialRhombusGridShape:
		return any(stage.InitialRhombusGridShapes_mapString).(map[string]Type)
	case *InitialRhombusShape:
		return any(stage.InitialRhombusShapes_mapString).(map[string]Type)
	case *Library:
		return any(stage.Librarys_mapString).(map[string]Type)
	case *NextCircleShape:
		return any(stage.NextCircleShapes_mapString).(map[string]Type)
	case *PerpendicularVector:
		return any(stage.PerpendicularVectors_mapString).(map[string]Type)
	case *PerpendicularVectorGrid:
		return any(stage.PerpendicularVectorGrids_mapString).(map[string]Type)
	case *PerpendicularVectorGridHalfway:
		return any(stage.PerpendicularVectorGridHalfways_mapString).(map[string]Type)
	case *PerpendicularVectorHalfway:
		return any(stage.PerpendicularVectorHalfways_mapString).(map[string]Type)
	case *Plant:
		return any(stage.Plants_mapString).(map[string]Type)
	case *PlantCircumferenceShape:
		return any(stage.PlantCircumferenceShapes_mapString).(map[string]Type)
	case *PlantDiagram:
		return any(stage.PlantDiagrams_mapString).(map[string]Type)
	case *Rendered3DShape:
		return any(stage.Rendered3DShapes_mapString).(map[string]Type)
	case *RhombusShape:
		return any(stage.RhombusShapes_mapString).(map[string]Type)
	case *RotatedRhombusGridShape:
		return any(stage.RotatedRhombusGridShapes_mapString).(map[string]Type)
	case *RotatedRhombusShape:
		return any(stage.RotatedRhombusShapes_mapString).(map[string]Type)
	case *StackGrowthCurveBezierShape:
		return any(stage.StackGrowthCurveBezierShapes_mapString).(map[string]Type)
	case *StackOfGrowthCurve:
		return any(stage.StackOfGrowthCurves_mapString).(map[string]Type)
	case *StartArcShape:
		return any(stage.StartArcShapes_mapString).(map[string]Type)
	case *StartArcShapeGrid:
		return any(stage.StartArcShapeGrids_mapString).(map[string]Type)
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
	case EndArcShape:
		return any(&stage.EndArcShapes).(*map[*Type]struct{})
	case EndArcShapeGrid:
		return any(&stage.EndArcShapeGrids).(*map[*Type]struct{})
	case ExplanationTextShape:
		return any(&stage.ExplanationTextShapes).(*map[*Type]struct{})
	case GridPathShape:
		return any(&stage.GridPathShapes).(*map[*Type]struct{})
	case GrowthCurveBezierShape:
		return any(&stage.GrowthCurveBezierShapes).(*map[*Type]struct{})
	case GrowthCurveBezierShapeGrid:
		return any(&stage.GrowthCurveBezierShapeGrids).(*map[*Type]struct{})
	case GrowthCurveRhombusGridShape:
		return any(&stage.GrowthCurveRhombusGridShapes).(*map[*Type]struct{})
	case GrowthCurveRhombusShape:
		return any(&stage.GrowthCurveRhombusShapes).(*map[*Type]struct{})
	case GrowthVectorShape:
		return any(&stage.GrowthVectorShapes).(*map[*Type]struct{})
	case InitialRhombusGridShape:
		return any(&stage.InitialRhombusGridShapes).(*map[*Type]struct{})
	case InitialRhombusShape:
		return any(&stage.InitialRhombusShapes).(*map[*Type]struct{})
	case Library:
		return any(&stage.Librarys).(*map[*Type]struct{})
	case NextCircleShape:
		return any(&stage.NextCircleShapes).(*map[*Type]struct{})
	case PerpendicularVector:
		return any(&stage.PerpendicularVectors).(*map[*Type]struct{})
	case PerpendicularVectorGrid:
		return any(&stage.PerpendicularVectorGrids).(*map[*Type]struct{})
	case PerpendicularVectorGridHalfway:
		return any(&stage.PerpendicularVectorGridHalfways).(*map[*Type]struct{})
	case PerpendicularVectorHalfway:
		return any(&stage.PerpendicularVectorHalfways).(*map[*Type]struct{})
	case Plant:
		return any(&stage.Plants).(*map[*Type]struct{})
	case PlantCircumferenceShape:
		return any(&stage.PlantCircumferenceShapes).(*map[*Type]struct{})
	case PlantDiagram:
		return any(&stage.PlantDiagrams).(*map[*Type]struct{})
	case Rendered3DShape:
		return any(&stage.Rendered3DShapes).(*map[*Type]struct{})
	case RhombusShape:
		return any(&stage.RhombusShapes).(*map[*Type]struct{})
	case RotatedRhombusGridShape:
		return any(&stage.RotatedRhombusGridShapes).(*map[*Type]struct{})
	case RotatedRhombusShape:
		return any(&stage.RotatedRhombusShapes).(*map[*Type]struct{})
	case StackGrowthCurveBezierShape:
		return any(&stage.StackGrowthCurveBezierShapes).(*map[*Type]struct{})
	case StackOfGrowthCurve:
		return any(&stage.StackOfGrowthCurves).(*map[*Type]struct{})
	case StartArcShape:
		return any(&stage.StartArcShapes).(*map[*Type]struct{})
	case StartArcShapeGrid:
		return any(&stage.StartArcShapeGrids).(*map[*Type]struct{})
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
	case *EndArcShape:
		return any(&stage.EndArcShapes).(*map[Type]struct{})
	case *EndArcShapeGrid:
		return any(&stage.EndArcShapeGrids).(*map[Type]struct{})
	case *ExplanationTextShape:
		return any(&stage.ExplanationTextShapes).(*map[Type]struct{})
	case *GridPathShape:
		return any(&stage.GridPathShapes).(*map[Type]struct{})
	case *GrowthCurveBezierShape:
		return any(&stage.GrowthCurveBezierShapes).(*map[Type]struct{})
	case *GrowthCurveBezierShapeGrid:
		return any(&stage.GrowthCurveBezierShapeGrids).(*map[Type]struct{})
	case *GrowthCurveRhombusGridShape:
		return any(&stage.GrowthCurveRhombusGridShapes).(*map[Type]struct{})
	case *GrowthCurveRhombusShape:
		return any(&stage.GrowthCurveRhombusShapes).(*map[Type]struct{})
	case *GrowthVectorShape:
		return any(&stage.GrowthVectorShapes).(*map[Type]struct{})
	case *InitialRhombusGridShape:
		return any(&stage.InitialRhombusGridShapes).(*map[Type]struct{})
	case *InitialRhombusShape:
		return any(&stage.InitialRhombusShapes).(*map[Type]struct{})
	case *Library:
		return any(&stage.Librarys).(*map[Type]struct{})
	case *NextCircleShape:
		return any(&stage.NextCircleShapes).(*map[Type]struct{})
	case *PerpendicularVector:
		return any(&stage.PerpendicularVectors).(*map[Type]struct{})
	case *PerpendicularVectorGrid:
		return any(&stage.PerpendicularVectorGrids).(*map[Type]struct{})
	case *PerpendicularVectorGridHalfway:
		return any(&stage.PerpendicularVectorGridHalfways).(*map[Type]struct{})
	case *PerpendicularVectorHalfway:
		return any(&stage.PerpendicularVectorHalfways).(*map[Type]struct{})
	case *Plant:
		return any(&stage.Plants).(*map[Type]struct{})
	case *PlantCircumferenceShape:
		return any(&stage.PlantCircumferenceShapes).(*map[Type]struct{})
	case *PlantDiagram:
		return any(&stage.PlantDiagrams).(*map[Type]struct{})
	case *Rendered3DShape:
		return any(&stage.Rendered3DShapes).(*map[Type]struct{})
	case *RhombusShape:
		return any(&stage.RhombusShapes).(*map[Type]struct{})
	case *RotatedRhombusGridShape:
		return any(&stage.RotatedRhombusGridShapes).(*map[Type]struct{})
	case *RotatedRhombusShape:
		return any(&stage.RotatedRhombusShapes).(*map[Type]struct{})
	case *StackGrowthCurveBezierShape:
		return any(&stage.StackGrowthCurveBezierShapes).(*map[Type]struct{})
	case *StackOfGrowthCurve:
		return any(&stage.StackOfGrowthCurves).(*map[Type]struct{})
	case *StartArcShape:
		return any(&stage.StartArcShapes).(*map[Type]struct{})
	case *StartArcShapeGrid:
		return any(&stage.StartArcShapeGrids).(*map[Type]struct{})
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
	case EndArcShape:
		return any(&stage.EndArcShapes_mapString).(*map[string]*Type)
	case EndArcShapeGrid:
		return any(&stage.EndArcShapeGrids_mapString).(*map[string]*Type)
	case ExplanationTextShape:
		return any(&stage.ExplanationTextShapes_mapString).(*map[string]*Type)
	case GridPathShape:
		return any(&stage.GridPathShapes_mapString).(*map[string]*Type)
	case GrowthCurveBezierShape:
		return any(&stage.GrowthCurveBezierShapes_mapString).(*map[string]*Type)
	case GrowthCurveBezierShapeGrid:
		return any(&stage.GrowthCurveBezierShapeGrids_mapString).(*map[string]*Type)
	case GrowthCurveRhombusGridShape:
		return any(&stage.GrowthCurveRhombusGridShapes_mapString).(*map[string]*Type)
	case GrowthCurveRhombusShape:
		return any(&stage.GrowthCurveRhombusShapes_mapString).(*map[string]*Type)
	case GrowthVectorShape:
		return any(&stage.GrowthVectorShapes_mapString).(*map[string]*Type)
	case InitialRhombusGridShape:
		return any(&stage.InitialRhombusGridShapes_mapString).(*map[string]*Type)
	case InitialRhombusShape:
		return any(&stage.InitialRhombusShapes_mapString).(*map[string]*Type)
	case Library:
		return any(&stage.Librarys_mapString).(*map[string]*Type)
	case NextCircleShape:
		return any(&stage.NextCircleShapes_mapString).(*map[string]*Type)
	case PerpendicularVector:
		return any(&stage.PerpendicularVectors_mapString).(*map[string]*Type)
	case PerpendicularVectorGrid:
		return any(&stage.PerpendicularVectorGrids_mapString).(*map[string]*Type)
	case PerpendicularVectorGridHalfway:
		return any(&stage.PerpendicularVectorGridHalfways_mapString).(*map[string]*Type)
	case PerpendicularVectorHalfway:
		return any(&stage.PerpendicularVectorHalfways_mapString).(*map[string]*Type)
	case Plant:
		return any(&stage.Plants_mapString).(*map[string]*Type)
	case PlantCircumferenceShape:
		return any(&stage.PlantCircumferenceShapes_mapString).(*map[string]*Type)
	case PlantDiagram:
		return any(&stage.PlantDiagrams_mapString).(*map[string]*Type)
	case Rendered3DShape:
		return any(&stage.Rendered3DShapes_mapString).(*map[string]*Type)
	case RhombusShape:
		return any(&stage.RhombusShapes_mapString).(*map[string]*Type)
	case RotatedRhombusGridShape:
		return any(&stage.RotatedRhombusGridShapes_mapString).(*map[string]*Type)
	case RotatedRhombusShape:
		return any(&stage.RotatedRhombusShapes_mapString).(*map[string]*Type)
	case StackGrowthCurveBezierShape:
		return any(&stage.StackGrowthCurveBezierShapes_mapString).(*map[string]*Type)
	case StackOfGrowthCurve:
		return any(&stage.StackOfGrowthCurves_mapString).(*map[string]*Type)
	case StartArcShape:
		return any(&stage.StartArcShapes_mapString).(*map[string]*Type)
	case StartArcShapeGrid:
		return any(&stage.StartArcShapeGrids_mapString).(*map[string]*Type)
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
	case EndArcShape:
		return any(&EndArcShape{
			// Initialisation of associations
		}).(*Type)
	case EndArcShapeGrid:
		return any(&EndArcShapeGrid{
			// Initialisation of associations
			// field is initialized with an instance of EndArcShape with the name of the field
			EndArcShapes: []*EndArcShape{{Name: "EndArcShapes"}},
		}).(*Type)
	case ExplanationTextShape:
		return any(&ExplanationTextShape{
			// Initialisation of associations
		}).(*Type)
	case GridPathShape:
		return any(&GridPathShape{
			// Initialisation of associations
		}).(*Type)
	case GrowthCurveBezierShape:
		return any(&GrowthCurveBezierShape{
			// Initialisation of associations
		}).(*Type)
	case GrowthCurveBezierShapeGrid:
		return any(&GrowthCurveBezierShapeGrid{
			// Initialisation of associations
			// field is initialized with an instance of GrowthCurveBezierShape with the name of the field
			GrowthCurveBezierShapes: []*GrowthCurveBezierShape{{Name: "GrowthCurveBezierShapes"}},
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
	case GrowthVectorShape:
		return any(&GrowthVectorShape{
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
	case PerpendicularVector:
		return any(&PerpendicularVector{
			// Initialisation of associations
		}).(*Type)
	case PerpendicularVectorGrid:
		return any(&PerpendicularVectorGrid{
			// Initialisation of associations
			// field is initialized with an instance of PerpendicularVector with the name of the field
			PerpendicularVectors: []*PerpendicularVector{{Name: "PerpendicularVectors"}},
		}).(*Type)
	case PerpendicularVectorGridHalfway:
		return any(&PerpendicularVectorGridHalfway{
			// Initialisation of associations
			// field is initialized with an instance of PerpendicularVectorHalfway with the name of the field
			PerpendicularVectorHalfways: []*PerpendicularVectorHalfway{{Name: "PerpendicularVectorHalfways"}},
		}).(*Type)
	case PerpendicularVectorHalfway:
		return any(&PerpendicularVectorHalfway{
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
			// field is initialized with an instance of GrowthVectorShape with the name of the field
			GrowthVectorShape: &GrowthVectorShape{Name: "GrowthVectorShape"},
			// field is initialized with an instance of PerpendicularVectorGrid with the name of the field
			PerpendicularVectorGrid: &PerpendicularVectorGrid{Name: "PerpendicularVectorGrid"},
			// field is initialized with an instance of PerpendicularVectorGridHalfway with the name of the field
			PerpendicularVectorGridHalfway: &PerpendicularVectorGridHalfway{Name: "PerpendicularVectorGridHalfway"},
			// field is initialized with an instance of StartArcShapeGrid with the name of the field
			StartArcShapeGrid: &StartArcShapeGrid{Name: "StartArcShapeGrid"},
			// field is initialized with an instance of EndArcShapeGrid with the name of the field
			EndArcShapeGrid: &EndArcShapeGrid{Name: "EndArcShapeGrid"},
			// field is initialized with an instance of GrowthCurveBezierShapeGrid with the name of the field
			GrowthCurveBezierShapeGrid: &GrowthCurveBezierShapeGrid{Name: "GrowthCurveBezierShapeGrid"},
			// field is initialized with an instance of StackOfGrowthCurve with the name of the field
			StackOfGrowthCurve: &StackOfGrowthCurve{Name: "StackOfGrowthCurve"},
		}).(*Type)
	case PlantCircumferenceShape:
		return any(&PlantCircumferenceShape{
			// Initialisation of associations
		}).(*Type)
	case PlantDiagram:
		return any(&PlantDiagram{
			// Initialisation of associations
			// field is initialized with an instance of Rendered3DShape with the name of the field
			Rendered3DShape: &Rendered3DShape{Name: "Rendered3DShape"},
		}).(*Type)
	case Rendered3DShape:
		return any(&Rendered3DShape{
			// Initialisation of associations
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
	case StackGrowthCurveBezierShape:
		return any(&StackGrowthCurveBezierShape{
			// Initialisation of associations
		}).(*Type)
	case StackOfGrowthCurve:
		return any(&StackOfGrowthCurve{
			// Initialisation of associations
			// field is initialized with an instance of StackGrowthCurveBezierShape with the name of the field
			StackGrowthCurveBezierShapes: []*StackGrowthCurveBezierShape{{Name: "StackGrowthCurveBezierShapes"}},
		}).(*Type)
	case StartArcShape:
		return any(&StartArcShape{
			// Initialisation of associations
		}).(*Type)
	case StartArcShapeGrid:
		return any(&StartArcShapeGrid{
			// Initialisation of associations
			// field is initialized with an instance of StartArcShape with the name of the field
			StartArcShapes: []*StartArcShape{{Name: "StartArcShapes"}},
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
	// reverse maps of direct associations of EndArcShape
	case EndArcShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EndArcShapeGrid
	case EndArcShapeGrid:
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
	// reverse maps of direct associations of GrowthCurveBezierShape
	case GrowthCurveBezierShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GrowthCurveBezierShapeGrid
	case GrowthCurveBezierShapeGrid:
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
	// reverse maps of direct associations of GrowthVectorShape
	case GrowthVectorShape:
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
	// reverse maps of direct associations of PerpendicularVector
	case PerpendicularVector:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PerpendicularVectorGrid
	case PerpendicularVectorGrid:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PerpendicularVectorGridHalfway
	case PerpendicularVectorGridHalfway:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PerpendicularVectorHalfway
	case PerpendicularVectorHalfway:
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
		case "GrowthVectorShape":
			res := make(map[*GrowthVectorShape][]*Plant)
			for plant := range stage.Plants {
				if plant.GrowthVectorShape != nil {
					growthvectorshape_ := plant.GrowthVectorShape
					var plants []*Plant
					_, ok := res[growthvectorshape_]
					if ok {
						plants = res[growthvectorshape_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[growthvectorshape_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "PerpendicularVectorGrid":
			res := make(map[*PerpendicularVectorGrid][]*Plant)
			for plant := range stage.Plants {
				if plant.PerpendicularVectorGrid != nil {
					perpendicularvectorgrid_ := plant.PerpendicularVectorGrid
					var plants []*Plant
					_, ok := res[perpendicularvectorgrid_]
					if ok {
						plants = res[perpendicularvectorgrid_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[perpendicularvectorgrid_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "PerpendicularVectorGridHalfway":
			res := make(map[*PerpendicularVectorGridHalfway][]*Plant)
			for plant := range stage.Plants {
				if plant.PerpendicularVectorGridHalfway != nil {
					perpendicularvectorgridhalfway_ := plant.PerpendicularVectorGridHalfway
					var plants []*Plant
					_, ok := res[perpendicularvectorgridhalfway_]
					if ok {
						plants = res[perpendicularvectorgridhalfway_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[perpendicularvectorgridhalfway_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "StartArcShapeGrid":
			res := make(map[*StartArcShapeGrid][]*Plant)
			for plant := range stage.Plants {
				if plant.StartArcShapeGrid != nil {
					startarcshapegrid_ := plant.StartArcShapeGrid
					var plants []*Plant
					_, ok := res[startarcshapegrid_]
					if ok {
						plants = res[startarcshapegrid_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[startarcshapegrid_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "EndArcShapeGrid":
			res := make(map[*EndArcShapeGrid][]*Plant)
			for plant := range stage.Plants {
				if plant.EndArcShapeGrid != nil {
					endarcshapegrid_ := plant.EndArcShapeGrid
					var plants []*Plant
					_, ok := res[endarcshapegrid_]
					if ok {
						plants = res[endarcshapegrid_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[endarcshapegrid_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "GrowthCurveBezierShapeGrid":
			res := make(map[*GrowthCurveBezierShapeGrid][]*Plant)
			for plant := range stage.Plants {
				if plant.GrowthCurveBezierShapeGrid != nil {
					growthcurvebeziershapegrid_ := plant.GrowthCurveBezierShapeGrid
					var plants []*Plant
					_, ok := res[growthcurvebeziershapegrid_]
					if ok {
						plants = res[growthcurvebeziershapegrid_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[growthcurvebeziershapegrid_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "StackOfGrowthCurve":
			res := make(map[*StackOfGrowthCurve][]*Plant)
			for plant := range stage.Plants {
				if plant.StackOfGrowthCurve != nil {
					stackofgrowthcurve_ := plant.StackOfGrowthCurve
					var plants []*Plant
					_, ok := res[stackofgrowthcurve_]
					if ok {
						plants = res[stackofgrowthcurve_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[stackofgrowthcurve_] = plants
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
		case "Rendered3DShape":
			res := make(map[*Rendered3DShape][]*PlantDiagram)
			for plantdiagram := range stage.PlantDiagrams {
				if plantdiagram.Rendered3DShape != nil {
					rendered3dshape_ := plantdiagram.Rendered3DShape
					var plantdiagrams []*PlantDiagram
					_, ok := res[rendered3dshape_]
					if ok {
						plantdiagrams = res[rendered3dshape_]
					} else {
						plantdiagrams = make([]*PlantDiagram, 0)
					}
					plantdiagrams = append(plantdiagrams, plantdiagram)
					res[rendered3dshape_] = plantdiagrams
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Rendered3DShape
	case Rendered3DShape:
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
	// reverse maps of direct associations of StackGrowthCurveBezierShape
	case StackGrowthCurveBezierShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StackOfGrowthCurve
	case StackOfGrowthCurve:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StartArcShape
	case StartArcShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StartArcShapeGrid
	case StartArcShapeGrid:
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
	// reverse maps of direct associations of EndArcShape
	case EndArcShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EndArcShapeGrid
	case EndArcShapeGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "EndArcShapes":
			res := make(map[*EndArcShape][]*EndArcShapeGrid)
			for endarcshapegrid := range stage.EndArcShapeGrids {
				for _, endarcshape_ := range endarcshapegrid.EndArcShapes {
					res[endarcshape_] = append(res[endarcshape_], endarcshapegrid)
				}
			}
			return any(res).(map[*End][]*Start)
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
	// reverse maps of direct associations of GrowthCurveBezierShape
	case GrowthCurveBezierShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GrowthCurveBezierShapeGrid
	case GrowthCurveBezierShapeGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "GrowthCurveBezierShapes":
			res := make(map[*GrowthCurveBezierShape][]*GrowthCurveBezierShapeGrid)
			for growthcurvebeziershapegrid := range stage.GrowthCurveBezierShapeGrids {
				for _, growthcurvebeziershape_ := range growthcurvebeziershapegrid.GrowthCurveBezierShapes {
					res[growthcurvebeziershape_] = append(res[growthcurvebeziershape_], growthcurvebeziershapegrid)
				}
			}
			return any(res).(map[*End][]*Start)
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
	// reverse maps of direct associations of GrowthVectorShape
	case GrowthVectorShape:
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
	// reverse maps of direct associations of PerpendicularVector
	case PerpendicularVector:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PerpendicularVectorGrid
	case PerpendicularVectorGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "PerpendicularVectors":
			res := make(map[*PerpendicularVector][]*PerpendicularVectorGrid)
			for perpendicularvectorgrid := range stage.PerpendicularVectorGrids {
				for _, perpendicularvector_ := range perpendicularvectorgrid.PerpendicularVectors {
					res[perpendicularvector_] = append(res[perpendicularvector_], perpendicularvectorgrid)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of PerpendicularVectorGridHalfway
	case PerpendicularVectorGridHalfway:
		switch fieldname {
		// insertion point for per direct association field
		case "PerpendicularVectorHalfways":
			res := make(map[*PerpendicularVectorHalfway][]*PerpendicularVectorGridHalfway)
			for perpendicularvectorgridhalfway := range stage.PerpendicularVectorGridHalfways {
				for _, perpendicularvectorhalfway_ := range perpendicularvectorgridhalfway.PerpendicularVectorHalfways {
					res[perpendicularvectorhalfway_] = append(res[perpendicularvectorhalfway_], perpendicularvectorgridhalfway)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of PerpendicularVectorHalfway
	case PerpendicularVectorHalfway:
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
	// reverse maps of direct associations of Rendered3DShape
	case Rendered3DShape:
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
	// reverse maps of direct associations of StackGrowthCurveBezierShape
	case StackGrowthCurveBezierShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StackOfGrowthCurve
	case StackOfGrowthCurve:
		switch fieldname {
		// insertion point for per direct association field
		case "StackGrowthCurveBezierShapes":
			res := make(map[*StackGrowthCurveBezierShape][]*StackOfGrowthCurve)
			for stackofgrowthcurve := range stage.StackOfGrowthCurves {
				for _, stackgrowthcurvebeziershape_ := range stackofgrowthcurve.StackGrowthCurveBezierShapes {
					res[stackgrowthcurvebeziershape_] = append(res[stackgrowthcurvebeziershape_], stackofgrowthcurve)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of StartArcShape
	case StartArcShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StartArcShapeGrid
	case StartArcShapeGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "StartArcShapes":
			res := make(map[*StartArcShape][]*StartArcShapeGrid)
			for startarcshapegrid := range stage.StartArcShapeGrids {
				for _, startarcshape_ := range startarcshapegrid.StartArcShapes {
					res[startarcshape_] = append(res[startarcshape_], startarcshapegrid)
				}
			}
			return any(res).(map[*End][]*Start)
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
	case *EndArcShape:
		res = "EndArcShape"
	case *EndArcShapeGrid:
		res = "EndArcShapeGrid"
	case *ExplanationTextShape:
		res = "ExplanationTextShape"
	case *GridPathShape:
		res = "GridPathShape"
	case *GrowthCurveBezierShape:
		res = "GrowthCurveBezierShape"
	case *GrowthCurveBezierShapeGrid:
		res = "GrowthCurveBezierShapeGrid"
	case *GrowthCurveRhombusGridShape:
		res = "GrowthCurveRhombusGridShape"
	case *GrowthCurveRhombusShape:
		res = "GrowthCurveRhombusShape"
	case *GrowthVectorShape:
		res = "GrowthVectorShape"
	case *InitialRhombusGridShape:
		res = "InitialRhombusGridShape"
	case *InitialRhombusShape:
		res = "InitialRhombusShape"
	case *Library:
		res = "Library"
	case *NextCircleShape:
		res = "NextCircleShape"
	case *PerpendicularVector:
		res = "PerpendicularVector"
	case *PerpendicularVectorGrid:
		res = "PerpendicularVectorGrid"
	case *PerpendicularVectorGridHalfway:
		res = "PerpendicularVectorGridHalfway"
	case *PerpendicularVectorHalfway:
		res = "PerpendicularVectorHalfway"
	case *Plant:
		res = "Plant"
	case *PlantCircumferenceShape:
		res = "PlantCircumferenceShape"
	case *PlantDiagram:
		res = "PlantDiagram"
	case *Rendered3DShape:
		res = "Rendered3DShape"
	case *RhombusShape:
		res = "RhombusShape"
	case *RotatedRhombusGridShape:
		res = "RotatedRhombusGridShape"
	case *RotatedRhombusShape:
		res = "RotatedRhombusShape"
	case *StackGrowthCurveBezierShape:
		res = "StackGrowthCurveBezierShape"
	case *StackOfGrowthCurve:
		res = "StackOfGrowthCurve"
	case *StartArcShape:
		res = "StartArcShape"
	case *StartArcShapeGrid:
		res = "StartArcShapeGrid"
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
	case *EndArcShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "EndArcShapeGrid"
		rf.Fieldname = "EndArcShapes"
		res = append(res, rf)
	case *EndArcShapeGrid:
		var rf ReverseField
		_ = rf
	case *ExplanationTextShape:
		var rf ReverseField
		_ = rf
	case *GridPathShape:
		var rf ReverseField
		_ = rf
	case *GrowthCurveBezierShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GrowthCurveBezierShapeGrid"
		rf.Fieldname = "GrowthCurveBezierShapes"
		res = append(res, rf)
	case *GrowthCurveBezierShapeGrid:
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
	case *GrowthVectorShape:
		var rf ReverseField
		_ = rf
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
	case *PerpendicularVector:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "PerpendicularVectorGrid"
		rf.Fieldname = "PerpendicularVectors"
		res = append(res, rf)
	case *PerpendicularVectorGrid:
		var rf ReverseField
		_ = rf
	case *PerpendicularVectorGridHalfway:
		var rf ReverseField
		_ = rf
	case *PerpendicularVectorHalfway:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "PerpendicularVectorGridHalfway"
		rf.Fieldname = "PerpendicularVectorHalfways"
		res = append(res, rf)
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
	case *Rendered3DShape:
		var rf ReverseField
		_ = rf
	case *RhombusShape:
		var rf ReverseField
		_ = rf
	case *RotatedRhombusGridShape:
		var rf ReverseField
		_ = rf
	case *RotatedRhombusShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "RotatedRhombusGridShape"
		rf.Fieldname = "RotatedRhombusShapes"
		res = append(res, rf)
	case *StackGrowthCurveBezierShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "StackOfGrowthCurve"
		rf.Fieldname = "StackGrowthCurveBezierShapes"
		res = append(res, rf)
	case *StackOfGrowthCurve:
		var rf ReverseField
		_ = rf
	case *StartArcShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "StartArcShapeGrid"
		rf.Fieldname = "StartArcShapes"
		res = append(res, rf)
	case *StartArcShapeGrid:
		var rf ReverseField
		_ = rf
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

func (endarcshape *EndArcShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StartX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StartY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "XAxisRotation",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "LargeArcFlag",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "SweepFlag",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "RadiusX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "RadiusY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (endarcshapegrid *EndArcShapeGrid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "EndArcShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "EndArcShape",
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

func (growthcurvebeziershape *GrowthCurveBezierShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StartX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StartY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ControlPointStartX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ControlPointStartY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ControlPointEndX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ControlPointEndY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "GrowthCurveBezierShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GrowthCurveBezierShape",
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

func (growthvectorshape *GrowthVectorShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (perpendicularvector *PerpendicularVector) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StartX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StartY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "PerpendicularVectors",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "PerpendicularVector",
		},
	}
	return
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "PerpendicularVectorHalfways",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "PerpendicularVectorHalfway",
		},
	}
	return
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StartX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StartY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndY",
			GongFieldValueType: GongFieldValueTypeFloat,
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
			Name:               "StackHeight",
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
		{
			Name:                 "GrowthVectorShape",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "GrowthVectorShape",
		},
		{
			Name:                 "PerpendicularVectorGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PerpendicularVectorGrid",
		},
		{
			Name:                 "PerpendicularVectorGridHalfway",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PerpendicularVectorGridHalfway",
		},
		{
			Name:                 "StartArcShapeGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "StartArcShapeGrid",
		},
		{
			Name:                 "EndArcShapeGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "EndArcShapeGrid",
		},
		{
			Name:                 "GrowthCurveBezierShapeGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "GrowthCurveBezierShapeGrid",
		},
		{
			Name:                 "StackOfGrowthCurve",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "StackOfGrowthCurve",
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
			Name:               "IsHiddenGrowthVectorShape",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenPerpendicularVectorGrid",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenPerpendicularVectorGridHalfway",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenStartArcShapeGrid",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenEndArcShapeGrid",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenGrowthCurveBezierShapeGrid",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenStackOfGrowthCurve",
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
		{
			Name:                 "Rendered3DShape",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Rendered3DShape",
		},
	}
	return
}

func (rendered3dshape *Rendered3DShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ViewX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ViewY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ViewZ",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "TargetX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "TargetY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "TargetZ",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Fov",
			GongFieldValueType: GongFieldValueTypeFloat,
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

func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StartX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StartY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ControlPointStartX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ControlPointStartY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ControlPointEndX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ControlPointEndY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (stackofgrowthcurve *StackOfGrowthCurve) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "StackGrowthCurveBezierShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "StackGrowthCurveBezierShape",
		},
	}
	return
}

func (startarcshape *StartArcShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StartX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StartY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "XAxisRotation",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "LargeArcFlag",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "SweepFlag",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "RadiusX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "RadiusY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (startarcshapegrid *StartArcShapeGrid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "StartArcShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "StartArcShape",
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

func (endarcshape *EndArcShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = endarcshape.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", endarcshape.StartX)
		res.valueFloat = endarcshape.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", endarcshape.StartY)
		res.valueFloat = endarcshape.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", endarcshape.EndX)
		res.valueFloat = endarcshape.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", endarcshape.EndY)
		res.valueFloat = endarcshape.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", endarcshape.XAxisRotation)
		res.valueFloat = endarcshape.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", endarcshape.LargeArcFlag)
		res.valueBool = endarcshape.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", endarcshape.SweepFlag)
		res.valueBool = endarcshape.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", endarcshape.RadiusX)
		res.valueFloat = endarcshape.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", endarcshape.RadiusY)
		res.valueFloat = endarcshape.RadiusY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (endarcshapegrid *EndArcShapeGrid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = endarcshapegrid.Name
	case "EndArcShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range endarcshapegrid.EndArcShapes {
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

func (growthcurvebeziershape *GrowthCurveBezierShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = growthcurvebeziershape.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", growthcurvebeziershape.StartX)
		res.valueFloat = growthcurvebeziershape.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", growthcurvebeziershape.StartY)
		res.valueFloat = growthcurvebeziershape.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ControlPointStartX":
		res.valueString = fmt.Sprintf("%f", growthcurvebeziershape.ControlPointStartX)
		res.valueFloat = growthcurvebeziershape.ControlPointStartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ControlPointStartY":
		res.valueString = fmt.Sprintf("%f", growthcurvebeziershape.ControlPointStartY)
		res.valueFloat = growthcurvebeziershape.ControlPointStartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", growthcurvebeziershape.EndX)
		res.valueFloat = growthcurvebeziershape.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", growthcurvebeziershape.EndY)
		res.valueFloat = growthcurvebeziershape.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ControlPointEndX":
		res.valueString = fmt.Sprintf("%f", growthcurvebeziershape.ControlPointEndX)
		res.valueFloat = growthcurvebeziershape.ControlPointEndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ControlPointEndY":
		res.valueString = fmt.Sprintf("%f", growthcurvebeziershape.ControlPointEndY)
		res.valueFloat = growthcurvebeziershape.ControlPointEndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = growthcurvebeziershapegrid.Name
	case "GrowthCurveBezierShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range growthcurvebeziershapegrid.GrowthCurveBezierShapes {
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

func (growthvectorshape *GrowthVectorShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = growthvectorshape.Name
	case "X":
		res.valueString = fmt.Sprintf("%f", growthvectorshape.X)
		res.valueFloat = growthvectorshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", growthvectorshape.Y)
		res.valueFloat = growthvectorshape.Y
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

func (perpendicularvector *PerpendicularVector) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = perpendicularvector.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", perpendicularvector.StartX)
		res.valueFloat = perpendicularvector.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", perpendicularvector.StartY)
		res.valueFloat = perpendicularvector.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", perpendicularvector.EndX)
		res.valueFloat = perpendicularvector.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", perpendicularvector.EndY)
		res.valueFloat = perpendicularvector.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = perpendicularvectorgrid.Name
	case "PerpendicularVectors":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range perpendicularvectorgrid.PerpendicularVectors {
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

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = perpendicularvectorgridhalfway.Name
	case "PerpendicularVectorHalfways":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range perpendicularvectorgridhalfway.PerpendicularVectorHalfways {
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

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = perpendicularvectorhalfway.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", perpendicularvectorhalfway.StartX)
		res.valueFloat = perpendicularvectorhalfway.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", perpendicularvectorhalfway.StartY)
		res.valueFloat = perpendicularvectorhalfway.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", perpendicularvectorhalfway.EndX)
		res.valueFloat = perpendicularvectorhalfway.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", perpendicularvectorhalfway.EndY)
		res.valueFloat = perpendicularvectorhalfway.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
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
	case "StackHeight":
		res.valueString = fmt.Sprintf("%d", plant.StackHeight)
		res.valueInt = plant.StackHeight
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
	case "GrowthVectorShape":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.GrowthVectorShape != nil {
			res.valueString = plant.GrowthVectorShape.Name
			res.ids = plant.GrowthVectorShape.GongGetUUID(stage)
		}
	case "PerpendicularVectorGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.PerpendicularVectorGrid != nil {
			res.valueString = plant.PerpendicularVectorGrid.Name
			res.ids = plant.PerpendicularVectorGrid.GongGetUUID(stage)
		}
	case "PerpendicularVectorGridHalfway":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.PerpendicularVectorGridHalfway != nil {
			res.valueString = plant.PerpendicularVectorGridHalfway.Name
			res.ids = plant.PerpendicularVectorGridHalfway.GongGetUUID(stage)
		}
	case "StartArcShapeGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.StartArcShapeGrid != nil {
			res.valueString = plant.StartArcShapeGrid.Name
			res.ids = plant.StartArcShapeGrid.GongGetUUID(stage)
		}
	case "EndArcShapeGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.EndArcShapeGrid != nil {
			res.valueString = plant.EndArcShapeGrid.Name
			res.ids = plant.EndArcShapeGrid.GongGetUUID(stage)
		}
	case "GrowthCurveBezierShapeGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.GrowthCurveBezierShapeGrid != nil {
			res.valueString = plant.GrowthCurveBezierShapeGrid.Name
			res.ids = plant.GrowthCurveBezierShapeGrid.GongGetUUID(stage)
		}
	case "StackOfGrowthCurve":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.StackOfGrowthCurve != nil {
			res.valueString = plant.StackOfGrowthCurve.Name
			res.ids = plant.StackOfGrowthCurve.GongGetUUID(stage)
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
	case "IsHiddenGrowthVectorShape":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenGrowthVectorShape)
		res.valueBool = plantdiagram.IsHiddenGrowthVectorShape
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenPerpendicularVectorGrid":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenPerpendicularVectorGrid)
		res.valueBool = plantdiagram.IsHiddenPerpendicularVectorGrid
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenPerpendicularVectorGridHalfway":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenPerpendicularVectorGridHalfway)
		res.valueBool = plantdiagram.IsHiddenPerpendicularVectorGridHalfway
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenStartArcShapeGrid":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenStartArcShapeGrid)
		res.valueBool = plantdiagram.IsHiddenStartArcShapeGrid
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenEndArcShapeGrid":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenEndArcShapeGrid)
		res.valueBool = plantdiagram.IsHiddenEndArcShapeGrid
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenGrowthCurveBezierShapeGrid":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenGrowthCurveBezierShapeGrid)
		res.valueBool = plantdiagram.IsHiddenGrowthCurveBezierShapeGrid
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenStackOfGrowthCurve":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenStackOfGrowthCurve)
		res.valueBool = plantdiagram.IsHiddenStackOfGrowthCurve
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
	case "Rendered3DShape":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plantdiagram.Rendered3DShape != nil {
			res.valueString = plantdiagram.Rendered3DShape.Name
			res.ids = plantdiagram.Rendered3DShape.GongGetUUID(stage)
		}
	}
	return
}

func (rendered3dshape *Rendered3DShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = rendered3dshape.Name
	case "ViewX":
		res.valueString = fmt.Sprintf("%f", rendered3dshape.ViewX)
		res.valueFloat = rendered3dshape.ViewX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ViewY":
		res.valueString = fmt.Sprintf("%f", rendered3dshape.ViewY)
		res.valueFloat = rendered3dshape.ViewY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ViewZ":
		res.valueString = fmt.Sprintf("%f", rendered3dshape.ViewZ)
		res.valueFloat = rendered3dshape.ViewZ
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "TargetX":
		res.valueString = fmt.Sprintf("%f", rendered3dshape.TargetX)
		res.valueFloat = rendered3dshape.TargetX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "TargetY":
		res.valueString = fmt.Sprintf("%f", rendered3dshape.TargetY)
		res.valueFloat = rendered3dshape.TargetY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "TargetZ":
		res.valueString = fmt.Sprintf("%f", rendered3dshape.TargetZ)
		res.valueFloat = rendered3dshape.TargetZ
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Fov":
		res.valueString = fmt.Sprintf("%f", rendered3dshape.Fov)
		res.valueFloat = rendered3dshape.Fov
		res.GongFieldValueType = GongFieldValueTypeFloat
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

func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = stackgrowthcurvebeziershape.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvebeziershape.StartX)
		res.valueFloat = stackgrowthcurvebeziershape.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvebeziershape.StartY)
		res.valueFloat = stackgrowthcurvebeziershape.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ControlPointStartX":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvebeziershape.ControlPointStartX)
		res.valueFloat = stackgrowthcurvebeziershape.ControlPointStartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ControlPointStartY":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvebeziershape.ControlPointStartY)
		res.valueFloat = stackgrowthcurvebeziershape.ControlPointStartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvebeziershape.EndX)
		res.valueFloat = stackgrowthcurvebeziershape.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvebeziershape.EndY)
		res.valueFloat = stackgrowthcurvebeziershape.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ControlPointEndX":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvebeziershape.ControlPointEndX)
		res.valueFloat = stackgrowthcurvebeziershape.ControlPointEndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ControlPointEndY":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvebeziershape.ControlPointEndY)
		res.valueFloat = stackgrowthcurvebeziershape.ControlPointEndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (stackofgrowthcurve *StackOfGrowthCurve) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = stackofgrowthcurve.Name
	case "StackGrowthCurveBezierShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range stackofgrowthcurve.StackGrowthCurveBezierShapes {
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

func (startarcshape *StartArcShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = startarcshape.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", startarcshape.StartX)
		res.valueFloat = startarcshape.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", startarcshape.StartY)
		res.valueFloat = startarcshape.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", startarcshape.EndX)
		res.valueFloat = startarcshape.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", startarcshape.EndY)
		res.valueFloat = startarcshape.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", startarcshape.XAxisRotation)
		res.valueFloat = startarcshape.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", startarcshape.LargeArcFlag)
		res.valueBool = startarcshape.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", startarcshape.SweepFlag)
		res.valueBool = startarcshape.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", startarcshape.RadiusX)
		res.valueFloat = startarcshape.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", startarcshape.RadiusY)
		res.valueFloat = startarcshape.RadiusY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (startarcshapegrid *StartArcShapeGrid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = startarcshapegrid.Name
	case "StartArcShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range startarcshapegrid.StartArcShapes {
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

func (endarcshape *EndArcShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		endarcshape.Name = value.GetValueString()
	case "StartX":
		endarcshape.StartX = value.GetValueFloat()
	case "StartY":
		endarcshape.StartY = value.GetValueFloat()
	case "EndX":
		endarcshape.EndX = value.GetValueFloat()
	case "EndY":
		endarcshape.EndY = value.GetValueFloat()
	case "XAxisRotation":
		endarcshape.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		endarcshape.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		endarcshape.SweepFlag = value.GetValueBool()
	case "RadiusX":
		endarcshape.RadiusX = value.GetValueFloat()
	case "RadiusY":
		endarcshape.RadiusY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (endarcshapegrid *EndArcShapeGrid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		endarcshapegrid.Name = value.GetValueString()
	case "EndArcShapes":
		endarcshapegrid.EndArcShapes = make([]*EndArcShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.EndArcShapes {
					if stage.EndArcShape_stagedOrder[__instance__] == uint(id) {
						endarcshapegrid.EndArcShapes = append(endarcshapegrid.EndArcShapes, __instance__)
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

func (growthcurvebeziershape *GrowthCurveBezierShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		growthcurvebeziershape.Name = value.GetValueString()
	case "StartX":
		growthcurvebeziershape.StartX = value.GetValueFloat()
	case "StartY":
		growthcurvebeziershape.StartY = value.GetValueFloat()
	case "ControlPointStartX":
		growthcurvebeziershape.ControlPointStartX = value.GetValueFloat()
	case "ControlPointStartY":
		growthcurvebeziershape.ControlPointStartY = value.GetValueFloat()
	case "EndX":
		growthcurvebeziershape.EndX = value.GetValueFloat()
	case "EndY":
		growthcurvebeziershape.EndY = value.GetValueFloat()
	case "ControlPointEndX":
		growthcurvebeziershape.ControlPointEndX = value.GetValueFloat()
	case "ControlPointEndY":
		growthcurvebeziershape.ControlPointEndY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		growthcurvebeziershapegrid.Name = value.GetValueString()
	case "GrowthCurveBezierShapes":
		growthcurvebeziershapegrid.GrowthCurveBezierShapes = make([]*GrowthCurveBezierShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.GrowthCurveBezierShapes {
					if stage.GrowthCurveBezierShape_stagedOrder[__instance__] == uint(id) {
						growthcurvebeziershapegrid.GrowthCurveBezierShapes = append(growthcurvebeziershapegrid.GrowthCurveBezierShapes, __instance__)
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

func (growthvectorshape *GrowthVectorShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		growthvectorshape.Name = value.GetValueString()
	case "X":
		growthvectorshape.X = value.GetValueFloat()
	case "Y":
		growthvectorshape.Y = value.GetValueFloat()
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

func (perpendicularvector *PerpendicularVector) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		perpendicularvector.Name = value.GetValueString()
	case "StartX":
		perpendicularvector.StartX = value.GetValueFloat()
	case "StartY":
		perpendicularvector.StartY = value.GetValueFloat()
	case "EndX":
		perpendicularvector.EndX = value.GetValueFloat()
	case "EndY":
		perpendicularvector.EndY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		perpendicularvectorgrid.Name = value.GetValueString()
	case "PerpendicularVectors":
		perpendicularvectorgrid.PerpendicularVectors = make([]*PerpendicularVector, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.PerpendicularVectors {
					if stage.PerpendicularVector_stagedOrder[__instance__] == uint(id) {
						perpendicularvectorgrid.PerpendicularVectors = append(perpendicularvectorgrid.PerpendicularVectors, __instance__)
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

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		perpendicularvectorgridhalfway.Name = value.GetValueString()
	case "PerpendicularVectorHalfways":
		perpendicularvectorgridhalfway.PerpendicularVectorHalfways = make([]*PerpendicularVectorHalfway, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.PerpendicularVectorHalfways {
					if stage.PerpendicularVectorHalfway_stagedOrder[__instance__] == uint(id) {
						perpendicularvectorgridhalfway.PerpendicularVectorHalfways = append(perpendicularvectorgridhalfway.PerpendicularVectorHalfways, __instance__)
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

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		perpendicularvectorhalfway.Name = value.GetValueString()
	case "StartX":
		perpendicularvectorhalfway.StartX = value.GetValueFloat()
	case "StartY":
		perpendicularvectorhalfway.StartY = value.GetValueFloat()
	case "EndX":
		perpendicularvectorhalfway.EndX = value.GetValueFloat()
	case "EndY":
		perpendicularvectorhalfway.EndY = value.GetValueFloat()
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
	case "StackHeight":
		plant.StackHeight = int(value.GetValueInt())
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
	case "GrowthVectorShape":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.GrowthVectorShape = nil
			for __instance__ := range stage.GrowthVectorShapes {
				if stage.GrowthVectorShape_stagedOrder[__instance__] == uint(id) {
					plant.GrowthVectorShape = __instance__
					break
				}
			}
		}
	case "PerpendicularVectorGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.PerpendicularVectorGrid = nil
			for __instance__ := range stage.PerpendicularVectorGrids {
				if stage.PerpendicularVectorGrid_stagedOrder[__instance__] == uint(id) {
					plant.PerpendicularVectorGrid = __instance__
					break
				}
			}
		}
	case "PerpendicularVectorGridHalfway":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.PerpendicularVectorGridHalfway = nil
			for __instance__ := range stage.PerpendicularVectorGridHalfways {
				if stage.PerpendicularVectorGridHalfway_stagedOrder[__instance__] == uint(id) {
					plant.PerpendicularVectorGridHalfway = __instance__
					break
				}
			}
		}
	case "StartArcShapeGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.StartArcShapeGrid = nil
			for __instance__ := range stage.StartArcShapeGrids {
				if stage.StartArcShapeGrid_stagedOrder[__instance__] == uint(id) {
					plant.StartArcShapeGrid = __instance__
					break
				}
			}
		}
	case "EndArcShapeGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.EndArcShapeGrid = nil
			for __instance__ := range stage.EndArcShapeGrids {
				if stage.EndArcShapeGrid_stagedOrder[__instance__] == uint(id) {
					plant.EndArcShapeGrid = __instance__
					break
				}
			}
		}
	case "GrowthCurveBezierShapeGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.GrowthCurveBezierShapeGrid = nil
			for __instance__ := range stage.GrowthCurveBezierShapeGrids {
				if stage.GrowthCurveBezierShapeGrid_stagedOrder[__instance__] == uint(id) {
					plant.GrowthCurveBezierShapeGrid = __instance__
					break
				}
			}
		}
	case "StackOfGrowthCurve":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.StackOfGrowthCurve = nil
			for __instance__ := range stage.StackOfGrowthCurves {
				if stage.StackOfGrowthCurve_stagedOrder[__instance__] == uint(id) {
					plant.StackOfGrowthCurve = __instance__
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
	case "IsHiddenGrowthVectorShape":
		plantdiagram.IsHiddenGrowthVectorShape = value.GetValueBool()
	case "IsHiddenPerpendicularVectorGrid":
		plantdiagram.IsHiddenPerpendicularVectorGrid = value.GetValueBool()
	case "IsHiddenPerpendicularVectorGridHalfway":
		plantdiagram.IsHiddenPerpendicularVectorGridHalfway = value.GetValueBool()
	case "IsHiddenStartArcShapeGrid":
		plantdiagram.IsHiddenStartArcShapeGrid = value.GetValueBool()
	case "IsHiddenEndArcShapeGrid":
		plantdiagram.IsHiddenEndArcShapeGrid = value.GetValueBool()
	case "IsHiddenGrowthCurveBezierShapeGrid":
		plantdiagram.IsHiddenGrowthCurveBezierShapeGrid = value.GetValueBool()
	case "IsHiddenStackOfGrowthCurve":
		plantdiagram.IsHiddenStackOfGrowthCurve = value.GetValueBool()
	case "IsChecked":
		plantdiagram.IsChecked = value.GetValueBool()
	case "ComputedPrefix":
		plantdiagram.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		plantdiagram.IsExpanded = value.GetValueBool()
	case "Rendered3DShape":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plantdiagram.Rendered3DShape = nil
			for __instance__ := range stage.Rendered3DShapes {
				if stage.Rendered3DShape_stagedOrder[__instance__] == uint(id) {
					plantdiagram.Rendered3DShape = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (rendered3dshape *Rendered3DShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		rendered3dshape.Name = value.GetValueString()
	case "ViewX":
		rendered3dshape.ViewX = value.GetValueFloat()
	case "ViewY":
		rendered3dshape.ViewY = value.GetValueFloat()
	case "ViewZ":
		rendered3dshape.ViewZ = value.GetValueFloat()
	case "TargetX":
		rendered3dshape.TargetX = value.GetValueFloat()
	case "TargetY":
		rendered3dshape.TargetY = value.GetValueFloat()
	case "TargetZ":
		rendered3dshape.TargetZ = value.GetValueFloat()
	case "Fov":
		rendered3dshape.Fov = value.GetValueFloat()
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

func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		stackgrowthcurvebeziershape.Name = value.GetValueString()
	case "StartX":
		stackgrowthcurvebeziershape.StartX = value.GetValueFloat()
	case "StartY":
		stackgrowthcurvebeziershape.StartY = value.GetValueFloat()
	case "ControlPointStartX":
		stackgrowthcurvebeziershape.ControlPointStartX = value.GetValueFloat()
	case "ControlPointStartY":
		stackgrowthcurvebeziershape.ControlPointStartY = value.GetValueFloat()
	case "EndX":
		stackgrowthcurvebeziershape.EndX = value.GetValueFloat()
	case "EndY":
		stackgrowthcurvebeziershape.EndY = value.GetValueFloat()
	case "ControlPointEndX":
		stackgrowthcurvebeziershape.ControlPointEndX = value.GetValueFloat()
	case "ControlPointEndY":
		stackgrowthcurvebeziershape.ControlPointEndY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (stackofgrowthcurve *StackOfGrowthCurve) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		stackofgrowthcurve.Name = value.GetValueString()
	case "StackGrowthCurveBezierShapes":
		stackofgrowthcurve.StackGrowthCurveBezierShapes = make([]*StackGrowthCurveBezierShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.StackGrowthCurveBezierShapes {
					if stage.StackGrowthCurveBezierShape_stagedOrder[__instance__] == uint(id) {
						stackofgrowthcurve.StackGrowthCurveBezierShapes = append(stackofgrowthcurve.StackGrowthCurveBezierShapes, __instance__)
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

func (startarcshape *StartArcShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		startarcshape.Name = value.GetValueString()
	case "StartX":
		startarcshape.StartX = value.GetValueFloat()
	case "StartY":
		startarcshape.StartY = value.GetValueFloat()
	case "EndX":
		startarcshape.EndX = value.GetValueFloat()
	case "EndY":
		startarcshape.EndY = value.GetValueFloat()
	case "XAxisRotation":
		startarcshape.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		startarcshape.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		startarcshape.SweepFlag = value.GetValueBool()
	case "RadiusX":
		startarcshape.RadiusX = value.GetValueFloat()
	case "RadiusY":
		startarcshape.RadiusY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (startarcshapegrid *StartArcShapeGrid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		startarcshapegrid.Name = value.GetValueString()
	case "StartArcShapes":
		startarcshapegrid.StartArcShapes = make([]*StartArcShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.StartArcShapes {
					if stage.StartArcShape_stagedOrder[__instance__] == uint(id) {
						startarcshapegrid.StartArcShapes = append(startarcshapegrid.StartArcShapes, __instance__)
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

func (endarcshape *EndArcShape) GongGetGongstructName() string {
	return "EndArcShape"
}

func (endarcshapegrid *EndArcShapeGrid) GongGetGongstructName() string {
	return "EndArcShapeGrid"
}

func (explanationtextshape *ExplanationTextShape) GongGetGongstructName() string {
	return "ExplanationTextShape"
}

func (gridpathshape *GridPathShape) GongGetGongstructName() string {
	return "GridPathShape"
}

func (growthcurvebeziershape *GrowthCurveBezierShape) GongGetGongstructName() string {
	return "GrowthCurveBezierShape"
}

func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) GongGetGongstructName() string {
	return "GrowthCurveBezierShapeGrid"
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongGetGongstructName() string {
	return "GrowthCurveRhombusGridShape"
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongGetGongstructName() string {
	return "GrowthCurveRhombusShape"
}

func (growthvectorshape *GrowthVectorShape) GongGetGongstructName() string {
	return "GrowthVectorShape"
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

func (perpendicularvector *PerpendicularVector) GongGetGongstructName() string {
	return "PerpendicularVector"
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongGetGongstructName() string {
	return "PerpendicularVectorGrid"
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongGetGongstructName() string {
	return "PerpendicularVectorGridHalfway"
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongGetGongstructName() string {
	return "PerpendicularVectorHalfway"
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

func (rendered3dshape *Rendered3DShape) GongGetGongstructName() string {
	return "Rendered3DShape"
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

func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) GongGetGongstructName() string {
	return "StackGrowthCurveBezierShape"
}

func (stackofgrowthcurve *StackOfGrowthCurve) GongGetGongstructName() string {
	return "StackOfGrowthCurve"
}

func (startarcshape *StartArcShape) GongGetGongstructName() string {
	return "StartArcShape"
}

func (startarcshapegrid *StartArcShapeGrid) GongGetGongstructName() string {
	return "StartArcShapeGrid"
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

	stage.EndArcShapes_mapString = make(map[string]*EndArcShape)
	for endarcshape := range stage.EndArcShapes {
		stage.EndArcShapes_mapString[endarcshape.Name] = endarcshape
	}

	stage.EndArcShapeGrids_mapString = make(map[string]*EndArcShapeGrid)
	for endarcshapegrid := range stage.EndArcShapeGrids {
		stage.EndArcShapeGrids_mapString[endarcshapegrid.Name] = endarcshapegrid
	}

	stage.ExplanationTextShapes_mapString = make(map[string]*ExplanationTextShape)
	for explanationtextshape := range stage.ExplanationTextShapes {
		stage.ExplanationTextShapes_mapString[explanationtextshape.Name] = explanationtextshape
	}

	stage.GridPathShapes_mapString = make(map[string]*GridPathShape)
	for gridpathshape := range stage.GridPathShapes {
		stage.GridPathShapes_mapString[gridpathshape.Name] = gridpathshape
	}

	stage.GrowthCurveBezierShapes_mapString = make(map[string]*GrowthCurveBezierShape)
	for growthcurvebeziershape := range stage.GrowthCurveBezierShapes {
		stage.GrowthCurveBezierShapes_mapString[growthcurvebeziershape.Name] = growthcurvebeziershape
	}

	stage.GrowthCurveBezierShapeGrids_mapString = make(map[string]*GrowthCurveBezierShapeGrid)
	for growthcurvebeziershapegrid := range stage.GrowthCurveBezierShapeGrids {
		stage.GrowthCurveBezierShapeGrids_mapString[growthcurvebeziershapegrid.Name] = growthcurvebeziershapegrid
	}

	stage.GrowthCurveRhombusGridShapes_mapString = make(map[string]*GrowthCurveRhombusGridShape)
	for growthcurverhombusgridshape := range stage.GrowthCurveRhombusGridShapes {
		stage.GrowthCurveRhombusGridShapes_mapString[growthcurverhombusgridshape.Name] = growthcurverhombusgridshape
	}

	stage.GrowthCurveRhombusShapes_mapString = make(map[string]*GrowthCurveRhombusShape)
	for growthcurverhombusshape := range stage.GrowthCurveRhombusShapes {
		stage.GrowthCurveRhombusShapes_mapString[growthcurverhombusshape.Name] = growthcurverhombusshape
	}

	stage.GrowthVectorShapes_mapString = make(map[string]*GrowthVectorShape)
	for growthvectorshape := range stage.GrowthVectorShapes {
		stage.GrowthVectorShapes_mapString[growthvectorshape.Name] = growthvectorshape
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

	stage.PerpendicularVectors_mapString = make(map[string]*PerpendicularVector)
	for perpendicularvector := range stage.PerpendicularVectors {
		stage.PerpendicularVectors_mapString[perpendicularvector.Name] = perpendicularvector
	}

	stage.PerpendicularVectorGrids_mapString = make(map[string]*PerpendicularVectorGrid)
	for perpendicularvectorgrid := range stage.PerpendicularVectorGrids {
		stage.PerpendicularVectorGrids_mapString[perpendicularvectorgrid.Name] = perpendicularvectorgrid
	}

	stage.PerpendicularVectorGridHalfways_mapString = make(map[string]*PerpendicularVectorGridHalfway)
	for perpendicularvectorgridhalfway := range stage.PerpendicularVectorGridHalfways {
		stage.PerpendicularVectorGridHalfways_mapString[perpendicularvectorgridhalfway.Name] = perpendicularvectorgridhalfway
	}

	stage.PerpendicularVectorHalfways_mapString = make(map[string]*PerpendicularVectorHalfway)
	for perpendicularvectorhalfway := range stage.PerpendicularVectorHalfways {
		stage.PerpendicularVectorHalfways_mapString[perpendicularvectorhalfway.Name] = perpendicularvectorhalfway
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

	stage.Rendered3DShapes_mapString = make(map[string]*Rendered3DShape)
	for rendered3dshape := range stage.Rendered3DShapes {
		stage.Rendered3DShapes_mapString[rendered3dshape.Name] = rendered3dshape
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

	stage.StackGrowthCurveBezierShapes_mapString = make(map[string]*StackGrowthCurveBezierShape)
	for stackgrowthcurvebeziershape := range stage.StackGrowthCurveBezierShapes {
		stage.StackGrowthCurveBezierShapes_mapString[stackgrowthcurvebeziershape.Name] = stackgrowthcurvebeziershape
	}

	stage.StackOfGrowthCurves_mapString = make(map[string]*StackOfGrowthCurve)
	for stackofgrowthcurve := range stage.StackOfGrowthCurves {
		stage.StackOfGrowthCurves_mapString[stackofgrowthcurve.Name] = stackofgrowthcurve
	}

	stage.StartArcShapes_mapString = make(map[string]*StartArcShape)
	for startarcshape := range stage.StartArcShapes {
		stage.StartArcShapes_mapString[startarcshape.Name] = startarcshape
	}

	stage.StartArcShapeGrids_mapString = make(map[string]*StartArcShapeGrid)
	for startarcshapegrid := range stage.StartArcShapeGrids {
		stage.StartArcShapeGrids_mapString[startarcshapegrid.Name] = startarcshapegrid
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
