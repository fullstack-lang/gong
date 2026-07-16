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
	ArcNormalVectorShapes                map[*ArcNormalVectorShape]struct{}
	ArcNormalVectorShapes_instance       map[*ArcNormalVectorShape]*ArcNormalVectorShape
	ArcNormalVectorShapes_mapString      map[string]*ArcNormalVectorShape
	ArcNormalVectorShapeOrder            uint
	ArcNormalVectorShape_stagedOrder     map[*ArcNormalVectorShape]uint
	ArcNormalVectorShape_orderStaged     map[uint]*ArcNormalVectorShape
	ArcNormalVectorShapes_reference      map[*ArcNormalVectorShape]*ArcNormalVectorShape
	ArcNormalVectorShapes_referenceOrder map[*ArcNormalVectorShape]uint

	// insertion point for slice of pointers maps
	OnAfterArcNormalVectorShapeCreateCallback OnAfterCreateInterface[ArcNormalVectorShape]
	OnAfterArcNormalVectorShapeUpdateCallback OnAfterUpdateInterface[ArcNormalVectorShape]
	OnAfterArcNormalVectorShapeDeleteCallback OnAfterDeleteInterface[ArcNormalVectorShape]
	OnAfterArcNormalVectorShapeReadCallback   OnAfterReadInterface[ArcNormalVectorShape]

	ArcNormalVectorShapeGrids                map[*ArcNormalVectorShapeGrid]struct{}
	ArcNormalVectorShapeGrids_instance       map[*ArcNormalVectorShapeGrid]*ArcNormalVectorShapeGrid
	ArcNormalVectorShapeGrids_mapString      map[string]*ArcNormalVectorShapeGrid
	ArcNormalVectorShapeGridOrder            uint
	ArcNormalVectorShapeGrid_stagedOrder     map[*ArcNormalVectorShapeGrid]uint
	ArcNormalVectorShapeGrid_orderStaged     map[uint]*ArcNormalVectorShapeGrid
	ArcNormalVectorShapeGrids_reference      map[*ArcNormalVectorShapeGrid]*ArcNormalVectorShapeGrid
	ArcNormalVectorShapeGrids_referenceOrder map[*ArcNormalVectorShapeGrid]uint

	// insertion point for slice of pointers maps
	ArcNormalVectorShapeGrid_ArcNormalVectorShapes_reverseMap map[*ArcNormalVectorShape]*ArcNormalVectorShapeGrid

	OnAfterArcNormalVectorShapeGridCreateCallback OnAfterCreateInterface[ArcNormalVectorShapeGrid]
	OnAfterArcNormalVectorShapeGridUpdateCallback OnAfterUpdateInterface[ArcNormalVectorShapeGrid]
	OnAfterArcNormalVectorShapeGridDeleteCallback OnAfterDeleteInterface[ArcNormalVectorShapeGrid]
	OnAfterArcNormalVectorShapeGridReadCallback   OnAfterReadInterface[ArcNormalVectorShapeGrid]

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

	BaseVectorShapes                map[*BaseVectorShape]struct{}
	BaseVectorShapes_instance       map[*BaseVectorShape]*BaseVectorShape
	BaseVectorShapes_mapString      map[string]*BaseVectorShape
	BaseVectorShapeOrder            uint
	BaseVectorShape_stagedOrder     map[*BaseVectorShape]uint
	BaseVectorShape_orderStaged     map[uint]*BaseVectorShape
	BaseVectorShapes_reference      map[*BaseVectorShape]*BaseVectorShape
	BaseVectorShapes_referenceOrder map[*BaseVectorShape]uint

	// insertion point for slice of pointers maps
	OnAfterBaseVectorShapeCreateCallback OnAfterCreateInterface[BaseVectorShape]
	OnAfterBaseVectorShapeUpdateCallback OnAfterUpdateInterface[BaseVectorShape]
	OnAfterBaseVectorShapeDeleteCallback OnAfterDeleteInterface[BaseVectorShape]
	OnAfterBaseVectorShapeReadCallback   OnAfterReadInterface[BaseVectorShape]

	BaseVectorShapeGrids                map[*BaseVectorShapeGrid]struct{}
	BaseVectorShapeGrids_instance       map[*BaseVectorShapeGrid]*BaseVectorShapeGrid
	BaseVectorShapeGrids_mapString      map[string]*BaseVectorShapeGrid
	BaseVectorShapeGridOrder            uint
	BaseVectorShapeGrid_stagedOrder     map[*BaseVectorShapeGrid]uint
	BaseVectorShapeGrid_orderStaged     map[uint]*BaseVectorShapeGrid
	BaseVectorShapeGrids_reference      map[*BaseVectorShapeGrid]*BaseVectorShapeGrid
	BaseVectorShapeGrids_referenceOrder map[*BaseVectorShapeGrid]uint

	// insertion point for slice of pointers maps
	BaseVectorShapeGrid_BaseVectorShapes_reverseMap map[*BaseVectorShape]*BaseVectorShapeGrid

	OnAfterBaseVectorShapeGridCreateCallback OnAfterCreateInterface[BaseVectorShapeGrid]
	OnAfterBaseVectorShapeGridUpdateCallback OnAfterUpdateInterface[BaseVectorShapeGrid]
	OnAfterBaseVectorShapeGridDeleteCallback OnAfterDeleteInterface[BaseVectorShapeGrid]
	OnAfterBaseVectorShapeGridReadCallback   OnAfterReadInterface[BaseVectorShapeGrid]

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

	GrowthCurve2Ds                map[*GrowthCurve2D]struct{}
	GrowthCurve2Ds_instance       map[*GrowthCurve2D]*GrowthCurve2D
	GrowthCurve2Ds_mapString      map[string]*GrowthCurve2D
	GrowthCurve2DOrder            uint
	GrowthCurve2D_stagedOrder     map[*GrowthCurve2D]uint
	GrowthCurve2D_orderStaged     map[uint]*GrowthCurve2D
	GrowthCurve2Ds_reference      map[*GrowthCurve2D]*GrowthCurve2D
	GrowthCurve2Ds_referenceOrder map[*GrowthCurve2D]uint

	// insertion point for slice of pointers maps
	OnAfterGrowthCurve2DCreateCallback OnAfterCreateInterface[GrowthCurve2D]
	OnAfterGrowthCurve2DUpdateCallback OnAfterUpdateInterface[GrowthCurve2D]
	OnAfterGrowthCurve2DDeleteCallback OnAfterDeleteInterface[GrowthCurve2D]
	OnAfterGrowthCurve2DReadCallback   OnAfterReadInterface[GrowthCurve2D]

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

	ShiftedLeftStackGrowthCurveEndArcShapes                map[*ShiftedLeftStackGrowthCurveEndArcShape]struct{}
	ShiftedLeftStackGrowthCurveEndArcShapes_instance       map[*ShiftedLeftStackGrowthCurveEndArcShape]*ShiftedLeftStackGrowthCurveEndArcShape
	ShiftedLeftStackGrowthCurveEndArcShapes_mapString      map[string]*ShiftedLeftStackGrowthCurveEndArcShape
	ShiftedLeftStackGrowthCurveEndArcShapeOrder            uint
	ShiftedLeftStackGrowthCurveEndArcShape_stagedOrder     map[*ShiftedLeftStackGrowthCurveEndArcShape]uint
	ShiftedLeftStackGrowthCurveEndArcShape_orderStaged     map[uint]*ShiftedLeftStackGrowthCurveEndArcShape
	ShiftedLeftStackGrowthCurveEndArcShapes_reference      map[*ShiftedLeftStackGrowthCurveEndArcShape]*ShiftedLeftStackGrowthCurveEndArcShape
	ShiftedLeftStackGrowthCurveEndArcShapes_referenceOrder map[*ShiftedLeftStackGrowthCurveEndArcShape]uint

	// insertion point for slice of pointers maps
	OnAfterShiftedLeftStackGrowthCurveEndArcShapeCreateCallback OnAfterCreateInterface[ShiftedLeftStackGrowthCurveEndArcShape]
	OnAfterShiftedLeftStackGrowthCurveEndArcShapeUpdateCallback OnAfterUpdateInterface[ShiftedLeftStackGrowthCurveEndArcShape]
	OnAfterShiftedLeftStackGrowthCurveEndArcShapeDeleteCallback OnAfterDeleteInterface[ShiftedLeftStackGrowthCurveEndArcShape]
	OnAfterShiftedLeftStackGrowthCurveEndArcShapeReadCallback   OnAfterReadInterface[ShiftedLeftStackGrowthCurveEndArcShape]

	ShiftedLeftStackGrowthCurveStartArcShapes                map[*ShiftedLeftStackGrowthCurveStartArcShape]struct{}
	ShiftedLeftStackGrowthCurveStartArcShapes_instance       map[*ShiftedLeftStackGrowthCurveStartArcShape]*ShiftedLeftStackGrowthCurveStartArcShape
	ShiftedLeftStackGrowthCurveStartArcShapes_mapString      map[string]*ShiftedLeftStackGrowthCurveStartArcShape
	ShiftedLeftStackGrowthCurveStartArcShapeOrder            uint
	ShiftedLeftStackGrowthCurveStartArcShape_stagedOrder     map[*ShiftedLeftStackGrowthCurveStartArcShape]uint
	ShiftedLeftStackGrowthCurveStartArcShape_orderStaged     map[uint]*ShiftedLeftStackGrowthCurveStartArcShape
	ShiftedLeftStackGrowthCurveStartArcShapes_reference      map[*ShiftedLeftStackGrowthCurveStartArcShape]*ShiftedLeftStackGrowthCurveStartArcShape
	ShiftedLeftStackGrowthCurveStartArcShapes_referenceOrder map[*ShiftedLeftStackGrowthCurveStartArcShape]uint

	// insertion point for slice of pointers maps
	OnAfterShiftedLeftStackGrowthCurveStartArcShapeCreateCallback OnAfterCreateInterface[ShiftedLeftStackGrowthCurveStartArcShape]
	OnAfterShiftedLeftStackGrowthCurveStartArcShapeUpdateCallback OnAfterUpdateInterface[ShiftedLeftStackGrowthCurveStartArcShape]
	OnAfterShiftedLeftStackGrowthCurveStartArcShapeDeleteCallback OnAfterDeleteInterface[ShiftedLeftStackGrowthCurveStartArcShape]
	OnAfterShiftedLeftStackGrowthCurveStartArcShapeReadCallback   OnAfterReadInterface[ShiftedLeftStackGrowthCurveStartArcShape]

	ShiftedLeftStackNormalVectors                map[*ShiftedLeftStackNormalVector]struct{}
	ShiftedLeftStackNormalVectors_instance       map[*ShiftedLeftStackNormalVector]*ShiftedLeftStackNormalVector
	ShiftedLeftStackNormalVectors_mapString      map[string]*ShiftedLeftStackNormalVector
	ShiftedLeftStackNormalVectorOrder            uint
	ShiftedLeftStackNormalVector_stagedOrder     map[*ShiftedLeftStackNormalVector]uint
	ShiftedLeftStackNormalVector_orderStaged     map[uint]*ShiftedLeftStackNormalVector
	ShiftedLeftStackNormalVectors_reference      map[*ShiftedLeftStackNormalVector]*ShiftedLeftStackNormalVector
	ShiftedLeftStackNormalVectors_referenceOrder map[*ShiftedLeftStackNormalVector]uint

	// insertion point for slice of pointers maps
	OnAfterShiftedLeftStackNormalVectorCreateCallback OnAfterCreateInterface[ShiftedLeftStackNormalVector]
	OnAfterShiftedLeftStackNormalVectorUpdateCallback OnAfterUpdateInterface[ShiftedLeftStackNormalVector]
	OnAfterShiftedLeftStackNormalVectorDeleteCallback OnAfterDeleteInterface[ShiftedLeftStackNormalVector]
	OnAfterShiftedLeftStackNormalVectorReadCallback   OnAfterReadInterface[ShiftedLeftStackNormalVector]

	ShiftedLeftStackOfGrowthCurves                map[*ShiftedLeftStackOfGrowthCurve]struct{}
	ShiftedLeftStackOfGrowthCurves_instance       map[*ShiftedLeftStackOfGrowthCurve]*ShiftedLeftStackOfGrowthCurve
	ShiftedLeftStackOfGrowthCurves_mapString      map[string]*ShiftedLeftStackOfGrowthCurve
	ShiftedLeftStackOfGrowthCurveOrder            uint
	ShiftedLeftStackOfGrowthCurve_stagedOrder     map[*ShiftedLeftStackOfGrowthCurve]uint
	ShiftedLeftStackOfGrowthCurve_orderStaged     map[uint]*ShiftedLeftStackOfGrowthCurve
	ShiftedLeftStackOfGrowthCurves_reference      map[*ShiftedLeftStackOfGrowthCurve]*ShiftedLeftStackOfGrowthCurve
	ShiftedLeftStackOfGrowthCurves_referenceOrder map[*ShiftedLeftStackOfGrowthCurve]uint

	// insertion point for slice of pointers maps
	ShiftedLeftStackOfGrowthCurve_ShiftedLeftStackGrowthCurveStartArcShapes_reverseMap map[*ShiftedLeftStackGrowthCurveStartArcShape]*ShiftedLeftStackOfGrowthCurve

	ShiftedLeftStackOfGrowthCurve_ShiftedLeftStackGrowthCurveEndArcShapes_reverseMap map[*ShiftedLeftStackGrowthCurveEndArcShape]*ShiftedLeftStackOfGrowthCurve

	OnAfterShiftedLeftStackOfGrowthCurveCreateCallback OnAfterCreateInterface[ShiftedLeftStackOfGrowthCurve]
	OnAfterShiftedLeftStackOfGrowthCurveUpdateCallback OnAfterUpdateInterface[ShiftedLeftStackOfGrowthCurve]
	OnAfterShiftedLeftStackOfGrowthCurveDeleteCallback OnAfterDeleteInterface[ShiftedLeftStackOfGrowthCurve]
	OnAfterShiftedLeftStackOfGrowthCurveReadCallback   OnAfterReadInterface[ShiftedLeftStackOfGrowthCurve]

	ShiftedLeftStackOfNormalVectors                map[*ShiftedLeftStackOfNormalVector]struct{}
	ShiftedLeftStackOfNormalVectors_instance       map[*ShiftedLeftStackOfNormalVector]*ShiftedLeftStackOfNormalVector
	ShiftedLeftStackOfNormalVectors_mapString      map[string]*ShiftedLeftStackOfNormalVector
	ShiftedLeftStackOfNormalVectorOrder            uint
	ShiftedLeftStackOfNormalVector_stagedOrder     map[*ShiftedLeftStackOfNormalVector]uint
	ShiftedLeftStackOfNormalVector_orderStaged     map[uint]*ShiftedLeftStackOfNormalVector
	ShiftedLeftStackOfNormalVectors_reference      map[*ShiftedLeftStackOfNormalVector]*ShiftedLeftStackOfNormalVector
	ShiftedLeftStackOfNormalVectors_referenceOrder map[*ShiftedLeftStackOfNormalVector]uint

	// insertion point for slice of pointers maps
	ShiftedLeftStackOfNormalVector_ShiftedLeftStackNormalVectors_reverseMap map[*ShiftedLeftStackNormalVector]*ShiftedLeftStackOfNormalVector

	OnAfterShiftedLeftStackOfNormalVectorCreateCallback OnAfterCreateInterface[ShiftedLeftStackOfNormalVector]
	OnAfterShiftedLeftStackOfNormalVectorUpdateCallback OnAfterUpdateInterface[ShiftedLeftStackOfNormalVector]
	OnAfterShiftedLeftStackOfNormalVectorDeleteCallback OnAfterDeleteInterface[ShiftedLeftStackOfNormalVector]
	OnAfterShiftedLeftStackOfNormalVectorReadCallback   OnAfterReadInterface[ShiftedLeftStackOfNormalVector]

	StackGrowthCurveEndArcShapes                map[*StackGrowthCurveEndArcShape]struct{}
	StackGrowthCurveEndArcShapes_instance       map[*StackGrowthCurveEndArcShape]*StackGrowthCurveEndArcShape
	StackGrowthCurveEndArcShapes_mapString      map[string]*StackGrowthCurveEndArcShape
	StackGrowthCurveEndArcShapeOrder            uint
	StackGrowthCurveEndArcShape_stagedOrder     map[*StackGrowthCurveEndArcShape]uint
	StackGrowthCurveEndArcShape_orderStaged     map[uint]*StackGrowthCurveEndArcShape
	StackGrowthCurveEndArcShapes_reference      map[*StackGrowthCurveEndArcShape]*StackGrowthCurveEndArcShape
	StackGrowthCurveEndArcShapes_referenceOrder map[*StackGrowthCurveEndArcShape]uint

	// insertion point for slice of pointers maps
	OnAfterStackGrowthCurveEndArcShapeCreateCallback OnAfterCreateInterface[StackGrowthCurveEndArcShape]
	OnAfterStackGrowthCurveEndArcShapeUpdateCallback OnAfterUpdateInterface[StackGrowthCurveEndArcShape]
	OnAfterStackGrowthCurveEndArcShapeDeleteCallback OnAfterDeleteInterface[StackGrowthCurveEndArcShape]
	OnAfterStackGrowthCurveEndArcShapeReadCallback   OnAfterReadInterface[StackGrowthCurveEndArcShape]

	StackGrowthCurveStartArcShapes                map[*StackGrowthCurveStartArcShape]struct{}
	StackGrowthCurveStartArcShapes_instance       map[*StackGrowthCurveStartArcShape]*StackGrowthCurveStartArcShape
	StackGrowthCurveStartArcShapes_mapString      map[string]*StackGrowthCurveStartArcShape
	StackGrowthCurveStartArcShapeOrder            uint
	StackGrowthCurveStartArcShape_stagedOrder     map[*StackGrowthCurveStartArcShape]uint
	StackGrowthCurveStartArcShape_orderStaged     map[uint]*StackGrowthCurveStartArcShape
	StackGrowthCurveStartArcShapes_reference      map[*StackGrowthCurveStartArcShape]*StackGrowthCurveStartArcShape
	StackGrowthCurveStartArcShapes_referenceOrder map[*StackGrowthCurveStartArcShape]uint

	// insertion point for slice of pointers maps
	OnAfterStackGrowthCurveStartArcShapeCreateCallback OnAfterCreateInterface[StackGrowthCurveStartArcShape]
	OnAfterStackGrowthCurveStartArcShapeUpdateCallback OnAfterUpdateInterface[StackGrowthCurveStartArcShape]
	OnAfterStackGrowthCurveStartArcShapeDeleteCallback OnAfterDeleteInterface[StackGrowthCurveStartArcShape]
	OnAfterStackGrowthCurveStartArcShapeReadCallback   OnAfterReadInterface[StackGrowthCurveStartArcShape]

	StackOfGrowthCurves                map[*StackOfGrowthCurve]struct{}
	StackOfGrowthCurves_instance       map[*StackOfGrowthCurve]*StackOfGrowthCurve
	StackOfGrowthCurves_mapString      map[string]*StackOfGrowthCurve
	StackOfGrowthCurveOrder            uint
	StackOfGrowthCurve_stagedOrder     map[*StackOfGrowthCurve]uint
	StackOfGrowthCurve_orderStaged     map[uint]*StackOfGrowthCurve
	StackOfGrowthCurves_reference      map[*StackOfGrowthCurve]*StackOfGrowthCurve
	StackOfGrowthCurves_referenceOrder map[*StackOfGrowthCurve]uint

	// insertion point for slice of pointers maps
	StackOfGrowthCurve_StackGrowthCurveStartArcShapes_reverseMap map[*StackGrowthCurveStartArcShape]*StackOfGrowthCurve

	StackOfGrowthCurve_StackGrowthCurveEndArcShapes_reverseMap map[*StackGrowthCurveEndArcShape]*StackOfGrowthCurve

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

	TopEndArcShapes                map[*TopEndArcShape]struct{}
	TopEndArcShapes_instance       map[*TopEndArcShape]*TopEndArcShape
	TopEndArcShapes_mapString      map[string]*TopEndArcShape
	TopEndArcShapeOrder            uint
	TopEndArcShape_stagedOrder     map[*TopEndArcShape]uint
	TopEndArcShape_orderStaged     map[uint]*TopEndArcShape
	TopEndArcShapes_reference      map[*TopEndArcShape]*TopEndArcShape
	TopEndArcShapes_referenceOrder map[*TopEndArcShape]uint

	// insertion point for slice of pointers maps
	OnAfterTopEndArcShapeCreateCallback OnAfterCreateInterface[TopEndArcShape]
	OnAfterTopEndArcShapeUpdateCallback OnAfterUpdateInterface[TopEndArcShape]
	OnAfterTopEndArcShapeDeleteCallback OnAfterDeleteInterface[TopEndArcShape]
	OnAfterTopEndArcShapeReadCallback   OnAfterReadInterface[TopEndArcShape]

	TopEndArcShapeGrids                map[*TopEndArcShapeGrid]struct{}
	TopEndArcShapeGrids_instance       map[*TopEndArcShapeGrid]*TopEndArcShapeGrid
	TopEndArcShapeGrids_mapString      map[string]*TopEndArcShapeGrid
	TopEndArcShapeGridOrder            uint
	TopEndArcShapeGrid_stagedOrder     map[*TopEndArcShapeGrid]uint
	TopEndArcShapeGrid_orderStaged     map[uint]*TopEndArcShapeGrid
	TopEndArcShapeGrids_reference      map[*TopEndArcShapeGrid]*TopEndArcShapeGrid
	TopEndArcShapeGrids_referenceOrder map[*TopEndArcShapeGrid]uint

	// insertion point for slice of pointers maps
	TopEndArcShapeGrid_TopEndArcShapes_reverseMap map[*TopEndArcShape]*TopEndArcShapeGrid

	OnAfterTopEndArcShapeGridCreateCallback OnAfterCreateInterface[TopEndArcShapeGrid]
	OnAfterTopEndArcShapeGridUpdateCallback OnAfterUpdateInterface[TopEndArcShapeGrid]
	OnAfterTopEndArcShapeGridDeleteCallback OnAfterDeleteInterface[TopEndArcShapeGrid]
	OnAfterTopEndArcShapeGridReadCallback   OnAfterReadInterface[TopEndArcShapeGrid]

	TopGrowthCurve2Ds                map[*TopGrowthCurve2D]struct{}
	TopGrowthCurve2Ds_instance       map[*TopGrowthCurve2D]*TopGrowthCurve2D
	TopGrowthCurve2Ds_mapString      map[string]*TopGrowthCurve2D
	TopGrowthCurve2DOrder            uint
	TopGrowthCurve2D_stagedOrder     map[*TopGrowthCurve2D]uint
	TopGrowthCurve2D_orderStaged     map[uint]*TopGrowthCurve2D
	TopGrowthCurve2Ds_reference      map[*TopGrowthCurve2D]*TopGrowthCurve2D
	TopGrowthCurve2Ds_referenceOrder map[*TopGrowthCurve2D]uint

	// insertion point for slice of pointers maps
	OnAfterTopGrowthCurve2DCreateCallback OnAfterCreateInterface[TopGrowthCurve2D]
	OnAfterTopGrowthCurve2DUpdateCallback OnAfterUpdateInterface[TopGrowthCurve2D]
	OnAfterTopGrowthCurve2DDeleteCallback OnAfterDeleteInterface[TopGrowthCurve2D]
	OnAfterTopGrowthCurve2DReadCallback   OnAfterReadInterface[TopGrowthCurve2D]

	TopStackGrowthCurveEndArcShapes                map[*TopStackGrowthCurveEndArcShape]struct{}
	TopStackGrowthCurveEndArcShapes_instance       map[*TopStackGrowthCurveEndArcShape]*TopStackGrowthCurveEndArcShape
	TopStackGrowthCurveEndArcShapes_mapString      map[string]*TopStackGrowthCurveEndArcShape
	TopStackGrowthCurveEndArcShapeOrder            uint
	TopStackGrowthCurveEndArcShape_stagedOrder     map[*TopStackGrowthCurveEndArcShape]uint
	TopStackGrowthCurveEndArcShape_orderStaged     map[uint]*TopStackGrowthCurveEndArcShape
	TopStackGrowthCurveEndArcShapes_reference      map[*TopStackGrowthCurveEndArcShape]*TopStackGrowthCurveEndArcShape
	TopStackGrowthCurveEndArcShapes_referenceOrder map[*TopStackGrowthCurveEndArcShape]uint

	// insertion point for slice of pointers maps
	OnAfterTopStackGrowthCurveEndArcShapeCreateCallback OnAfterCreateInterface[TopStackGrowthCurveEndArcShape]
	OnAfterTopStackGrowthCurveEndArcShapeUpdateCallback OnAfterUpdateInterface[TopStackGrowthCurveEndArcShape]
	OnAfterTopStackGrowthCurveEndArcShapeDeleteCallback OnAfterDeleteInterface[TopStackGrowthCurveEndArcShape]
	OnAfterTopStackGrowthCurveEndArcShapeReadCallback   OnAfterReadInterface[TopStackGrowthCurveEndArcShape]

	TopStackGrowthCurveStartArcShapes                map[*TopStackGrowthCurveStartArcShape]struct{}
	TopStackGrowthCurveStartArcShapes_instance       map[*TopStackGrowthCurveStartArcShape]*TopStackGrowthCurveStartArcShape
	TopStackGrowthCurveStartArcShapes_mapString      map[string]*TopStackGrowthCurveStartArcShape
	TopStackGrowthCurveStartArcShapeOrder            uint
	TopStackGrowthCurveStartArcShape_stagedOrder     map[*TopStackGrowthCurveStartArcShape]uint
	TopStackGrowthCurveStartArcShape_orderStaged     map[uint]*TopStackGrowthCurveStartArcShape
	TopStackGrowthCurveStartArcShapes_reference      map[*TopStackGrowthCurveStartArcShape]*TopStackGrowthCurveStartArcShape
	TopStackGrowthCurveStartArcShapes_referenceOrder map[*TopStackGrowthCurveStartArcShape]uint

	// insertion point for slice of pointers maps
	OnAfterTopStackGrowthCurveStartArcShapeCreateCallback OnAfterCreateInterface[TopStackGrowthCurveStartArcShape]
	OnAfterTopStackGrowthCurveStartArcShapeUpdateCallback OnAfterUpdateInterface[TopStackGrowthCurveStartArcShape]
	OnAfterTopStackGrowthCurveStartArcShapeDeleteCallback OnAfterDeleteInterface[TopStackGrowthCurveStartArcShape]
	OnAfterTopStackGrowthCurveStartArcShapeReadCallback   OnAfterReadInterface[TopStackGrowthCurveStartArcShape]

	TopStackOfGrowthCurves                map[*TopStackOfGrowthCurve]struct{}
	TopStackOfGrowthCurves_instance       map[*TopStackOfGrowthCurve]*TopStackOfGrowthCurve
	TopStackOfGrowthCurves_mapString      map[string]*TopStackOfGrowthCurve
	TopStackOfGrowthCurveOrder            uint
	TopStackOfGrowthCurve_stagedOrder     map[*TopStackOfGrowthCurve]uint
	TopStackOfGrowthCurve_orderStaged     map[uint]*TopStackOfGrowthCurve
	TopStackOfGrowthCurves_reference      map[*TopStackOfGrowthCurve]*TopStackOfGrowthCurve
	TopStackOfGrowthCurves_referenceOrder map[*TopStackOfGrowthCurve]uint

	// insertion point for slice of pointers maps
	TopStackOfGrowthCurve_TopStackGrowthCurveStartArcShapes_reverseMap map[*TopStackGrowthCurveStartArcShape]*TopStackOfGrowthCurve

	TopStackOfGrowthCurve_TopStackGrowthCurveEndArcShapes_reverseMap map[*TopStackGrowthCurveEndArcShape]*TopStackOfGrowthCurve

	OnAfterTopStackOfGrowthCurveCreateCallback OnAfterCreateInterface[TopStackOfGrowthCurve]
	OnAfterTopStackOfGrowthCurveUpdateCallback OnAfterUpdateInterface[TopStackOfGrowthCurve]
	OnAfterTopStackOfGrowthCurveDeleteCallback OnAfterDeleteInterface[TopStackOfGrowthCurve]
	OnAfterTopStackOfGrowthCurveReadCallback   OnAfterReadInterface[TopStackOfGrowthCurve]

	TopStartArcShapes                map[*TopStartArcShape]struct{}
	TopStartArcShapes_instance       map[*TopStartArcShape]*TopStartArcShape
	TopStartArcShapes_mapString      map[string]*TopStartArcShape
	TopStartArcShapeOrder            uint
	TopStartArcShape_stagedOrder     map[*TopStartArcShape]uint
	TopStartArcShape_orderStaged     map[uint]*TopStartArcShape
	TopStartArcShapes_reference      map[*TopStartArcShape]*TopStartArcShape
	TopStartArcShapes_referenceOrder map[*TopStartArcShape]uint

	// insertion point for slice of pointers maps
	OnAfterTopStartArcShapeCreateCallback OnAfterCreateInterface[TopStartArcShape]
	OnAfterTopStartArcShapeUpdateCallback OnAfterUpdateInterface[TopStartArcShape]
	OnAfterTopStartArcShapeDeleteCallback OnAfterDeleteInterface[TopStartArcShape]
	OnAfterTopStartArcShapeReadCallback   OnAfterReadInterface[TopStartArcShape]

	TopStartArcShapeGrids                map[*TopStartArcShapeGrid]struct{}
	TopStartArcShapeGrids_instance       map[*TopStartArcShapeGrid]*TopStartArcShapeGrid
	TopStartArcShapeGrids_mapString      map[string]*TopStartArcShapeGrid
	TopStartArcShapeGridOrder            uint
	TopStartArcShapeGrid_stagedOrder     map[*TopStartArcShapeGrid]uint
	TopStartArcShapeGrid_orderStaged     map[uint]*TopStartArcShapeGrid
	TopStartArcShapeGrids_reference      map[*TopStartArcShapeGrid]*TopStartArcShapeGrid
	TopStartArcShapeGrids_referenceOrder map[*TopStartArcShapeGrid]uint

	// insertion point for slice of pointers maps
	TopStartArcShapeGrid_TopStartArcShapes_reverseMap map[*TopStartArcShape]*TopStartArcShapeGrid

	OnAfterTopStartArcShapeGridCreateCallback OnAfterCreateInterface[TopStartArcShapeGrid]
	OnAfterTopStartArcShapeGridUpdateCallback OnAfterUpdateInterface[TopStartArcShapeGrid]
	OnAfterTopStartArcShapeGridDeleteCallback OnAfterDeleteInterface[TopStartArcShapeGrid]
	OnAfterTopStartArcShapeGridReadCallback   OnAfterReadInterface[TopStartArcShapeGrid]

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
	stage.ArcNormalVectorShapes_reference = make(map[*ArcNormalVectorShape]*ArcNormalVectorShape)
	stage.ArcNormalVectorShapes_instance = make(map[*ArcNormalVectorShape]*ArcNormalVectorShape)
	stage.ArcNormalVectorShapes_referenceOrder = make(map[*ArcNormalVectorShape]uint)

	stage.ArcNormalVectorShapeGrids_reference = make(map[*ArcNormalVectorShapeGrid]*ArcNormalVectorShapeGrid)
	stage.ArcNormalVectorShapeGrids_instance = make(map[*ArcNormalVectorShapeGrid]*ArcNormalVectorShapeGrid)
	stage.ArcNormalVectorShapeGrids_referenceOrder = make(map[*ArcNormalVectorShapeGrid]uint)

	stage.AxesShapes_reference = make(map[*AxesShape]*AxesShape)
	stage.AxesShapes_instance = make(map[*AxesShape]*AxesShape)
	stage.AxesShapes_referenceOrder = make(map[*AxesShape]uint)

	stage.BaseVectorShapes_reference = make(map[*BaseVectorShape]*BaseVectorShape)
	stage.BaseVectorShapes_instance = make(map[*BaseVectorShape]*BaseVectorShape)
	stage.BaseVectorShapes_referenceOrder = make(map[*BaseVectorShape]uint)

	stage.BaseVectorShapeGrids_reference = make(map[*BaseVectorShapeGrid]*BaseVectorShapeGrid)
	stage.BaseVectorShapeGrids_instance = make(map[*BaseVectorShapeGrid]*BaseVectorShapeGrid)
	stage.BaseVectorShapeGrids_referenceOrder = make(map[*BaseVectorShapeGrid]uint)

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

	stage.GrowthCurve2Ds_reference = make(map[*GrowthCurve2D]*GrowthCurve2D)
	stage.GrowthCurve2Ds_instance = make(map[*GrowthCurve2D]*GrowthCurve2D)
	stage.GrowthCurve2Ds_referenceOrder = make(map[*GrowthCurve2D]uint)

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

	stage.ShiftedLeftStackGrowthCurveEndArcShapes_reference = make(map[*ShiftedLeftStackGrowthCurveEndArcShape]*ShiftedLeftStackGrowthCurveEndArcShape)
	stage.ShiftedLeftStackGrowthCurveEndArcShapes_instance = make(map[*ShiftedLeftStackGrowthCurveEndArcShape]*ShiftedLeftStackGrowthCurveEndArcShape)
	stage.ShiftedLeftStackGrowthCurveEndArcShapes_referenceOrder = make(map[*ShiftedLeftStackGrowthCurveEndArcShape]uint)

	stage.ShiftedLeftStackGrowthCurveStartArcShapes_reference = make(map[*ShiftedLeftStackGrowthCurveStartArcShape]*ShiftedLeftStackGrowthCurveStartArcShape)
	stage.ShiftedLeftStackGrowthCurveStartArcShapes_instance = make(map[*ShiftedLeftStackGrowthCurveStartArcShape]*ShiftedLeftStackGrowthCurveStartArcShape)
	stage.ShiftedLeftStackGrowthCurveStartArcShapes_referenceOrder = make(map[*ShiftedLeftStackGrowthCurveStartArcShape]uint)

	stage.ShiftedLeftStackNormalVectors_reference = make(map[*ShiftedLeftStackNormalVector]*ShiftedLeftStackNormalVector)
	stage.ShiftedLeftStackNormalVectors_instance = make(map[*ShiftedLeftStackNormalVector]*ShiftedLeftStackNormalVector)
	stage.ShiftedLeftStackNormalVectors_referenceOrder = make(map[*ShiftedLeftStackNormalVector]uint)

	stage.ShiftedLeftStackOfGrowthCurves_reference = make(map[*ShiftedLeftStackOfGrowthCurve]*ShiftedLeftStackOfGrowthCurve)
	stage.ShiftedLeftStackOfGrowthCurves_instance = make(map[*ShiftedLeftStackOfGrowthCurve]*ShiftedLeftStackOfGrowthCurve)
	stage.ShiftedLeftStackOfGrowthCurves_referenceOrder = make(map[*ShiftedLeftStackOfGrowthCurve]uint)

	stage.ShiftedLeftStackOfNormalVectors_reference = make(map[*ShiftedLeftStackOfNormalVector]*ShiftedLeftStackOfNormalVector)
	stage.ShiftedLeftStackOfNormalVectors_instance = make(map[*ShiftedLeftStackOfNormalVector]*ShiftedLeftStackOfNormalVector)
	stage.ShiftedLeftStackOfNormalVectors_referenceOrder = make(map[*ShiftedLeftStackOfNormalVector]uint)

	stage.StackGrowthCurveEndArcShapes_reference = make(map[*StackGrowthCurveEndArcShape]*StackGrowthCurveEndArcShape)
	stage.StackGrowthCurveEndArcShapes_instance = make(map[*StackGrowthCurveEndArcShape]*StackGrowthCurveEndArcShape)
	stage.StackGrowthCurveEndArcShapes_referenceOrder = make(map[*StackGrowthCurveEndArcShape]uint)

	stage.StackGrowthCurveStartArcShapes_reference = make(map[*StackGrowthCurveStartArcShape]*StackGrowthCurveStartArcShape)
	stage.StackGrowthCurveStartArcShapes_instance = make(map[*StackGrowthCurveStartArcShape]*StackGrowthCurveStartArcShape)
	stage.StackGrowthCurveStartArcShapes_referenceOrder = make(map[*StackGrowthCurveStartArcShape]uint)

	stage.StackOfGrowthCurves_reference = make(map[*StackOfGrowthCurve]*StackOfGrowthCurve)
	stage.StackOfGrowthCurves_instance = make(map[*StackOfGrowthCurve]*StackOfGrowthCurve)
	stage.StackOfGrowthCurves_referenceOrder = make(map[*StackOfGrowthCurve]uint)

	stage.StartArcShapes_reference = make(map[*StartArcShape]*StartArcShape)
	stage.StartArcShapes_instance = make(map[*StartArcShape]*StartArcShape)
	stage.StartArcShapes_referenceOrder = make(map[*StartArcShape]uint)

	stage.StartArcShapeGrids_reference = make(map[*StartArcShapeGrid]*StartArcShapeGrid)
	stage.StartArcShapeGrids_instance = make(map[*StartArcShapeGrid]*StartArcShapeGrid)
	stage.StartArcShapeGrids_referenceOrder = make(map[*StartArcShapeGrid]uint)

	stage.TopEndArcShapes_reference = make(map[*TopEndArcShape]*TopEndArcShape)
	stage.TopEndArcShapes_instance = make(map[*TopEndArcShape]*TopEndArcShape)
	stage.TopEndArcShapes_referenceOrder = make(map[*TopEndArcShape]uint)

	stage.TopEndArcShapeGrids_reference = make(map[*TopEndArcShapeGrid]*TopEndArcShapeGrid)
	stage.TopEndArcShapeGrids_instance = make(map[*TopEndArcShapeGrid]*TopEndArcShapeGrid)
	stage.TopEndArcShapeGrids_referenceOrder = make(map[*TopEndArcShapeGrid]uint)

	stage.TopGrowthCurve2Ds_reference = make(map[*TopGrowthCurve2D]*TopGrowthCurve2D)
	stage.TopGrowthCurve2Ds_instance = make(map[*TopGrowthCurve2D]*TopGrowthCurve2D)
	stage.TopGrowthCurve2Ds_referenceOrder = make(map[*TopGrowthCurve2D]uint)

	stage.TopStackGrowthCurveEndArcShapes_reference = make(map[*TopStackGrowthCurveEndArcShape]*TopStackGrowthCurveEndArcShape)
	stage.TopStackGrowthCurveEndArcShapes_instance = make(map[*TopStackGrowthCurveEndArcShape]*TopStackGrowthCurveEndArcShape)
	stage.TopStackGrowthCurveEndArcShapes_referenceOrder = make(map[*TopStackGrowthCurveEndArcShape]uint)

	stage.TopStackGrowthCurveStartArcShapes_reference = make(map[*TopStackGrowthCurveStartArcShape]*TopStackGrowthCurveStartArcShape)
	stage.TopStackGrowthCurveStartArcShapes_instance = make(map[*TopStackGrowthCurveStartArcShape]*TopStackGrowthCurveStartArcShape)
	stage.TopStackGrowthCurveStartArcShapes_referenceOrder = make(map[*TopStackGrowthCurveStartArcShape]uint)

	stage.TopStackOfGrowthCurves_reference = make(map[*TopStackOfGrowthCurve]*TopStackOfGrowthCurve)
	stage.TopStackOfGrowthCurves_instance = make(map[*TopStackOfGrowthCurve]*TopStackOfGrowthCurve)
	stage.TopStackOfGrowthCurves_referenceOrder = make(map[*TopStackOfGrowthCurve]uint)

	stage.TopStartArcShapes_reference = make(map[*TopStartArcShape]*TopStartArcShape)
	stage.TopStartArcShapes_instance = make(map[*TopStartArcShape]*TopStartArcShape)
	stage.TopStartArcShapes_referenceOrder = make(map[*TopStartArcShape]uint)

	stage.TopStartArcShapeGrids_reference = make(map[*TopStartArcShapeGrid]*TopStartArcShapeGrid)
	stage.TopStartArcShapeGrids_instance = make(map[*TopStartArcShapeGrid]*TopStartArcShapeGrid)
	stage.TopStartArcShapeGrids_referenceOrder = make(map[*TopStartArcShapeGrid]uint)

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
	var maxArcNormalVectorShapeOrder uint
	var foundArcNormalVectorShape bool
	for _, order := range stage.ArcNormalVectorShape_stagedOrder {
		if !foundArcNormalVectorShape || order > maxArcNormalVectorShapeOrder {
			maxArcNormalVectorShapeOrder = order
			foundArcNormalVectorShape = true
		}
	}
	if foundArcNormalVectorShape {
		stage.ArcNormalVectorShapeOrder = maxArcNormalVectorShapeOrder + 1
	} else {
		stage.ArcNormalVectorShapeOrder = 0
	}

	var maxArcNormalVectorShapeGridOrder uint
	var foundArcNormalVectorShapeGrid bool
	for _, order := range stage.ArcNormalVectorShapeGrid_stagedOrder {
		if !foundArcNormalVectorShapeGrid || order > maxArcNormalVectorShapeGridOrder {
			maxArcNormalVectorShapeGridOrder = order
			foundArcNormalVectorShapeGrid = true
		}
	}
	if foundArcNormalVectorShapeGrid {
		stage.ArcNormalVectorShapeGridOrder = maxArcNormalVectorShapeGridOrder + 1
	} else {
		stage.ArcNormalVectorShapeGridOrder = 0
	}

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

	var maxBaseVectorShapeOrder uint
	var foundBaseVectorShape bool
	for _, order := range stage.BaseVectorShape_stagedOrder {
		if !foundBaseVectorShape || order > maxBaseVectorShapeOrder {
			maxBaseVectorShapeOrder = order
			foundBaseVectorShape = true
		}
	}
	if foundBaseVectorShape {
		stage.BaseVectorShapeOrder = maxBaseVectorShapeOrder + 1
	} else {
		stage.BaseVectorShapeOrder = 0
	}

	var maxBaseVectorShapeGridOrder uint
	var foundBaseVectorShapeGrid bool
	for _, order := range stage.BaseVectorShapeGrid_stagedOrder {
		if !foundBaseVectorShapeGrid || order > maxBaseVectorShapeGridOrder {
			maxBaseVectorShapeGridOrder = order
			foundBaseVectorShapeGrid = true
		}
	}
	if foundBaseVectorShapeGrid {
		stage.BaseVectorShapeGridOrder = maxBaseVectorShapeGridOrder + 1
	} else {
		stage.BaseVectorShapeGridOrder = 0
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

	var maxGrowthCurve2DOrder uint
	var foundGrowthCurve2D bool
	for _, order := range stage.GrowthCurve2D_stagedOrder {
		if !foundGrowthCurve2D || order > maxGrowthCurve2DOrder {
			maxGrowthCurve2DOrder = order
			foundGrowthCurve2D = true
		}
	}
	if foundGrowthCurve2D {
		stage.GrowthCurve2DOrder = maxGrowthCurve2DOrder + 1
	} else {
		stage.GrowthCurve2DOrder = 0
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

	var maxShiftedLeftStackGrowthCurveEndArcShapeOrder uint
	var foundShiftedLeftStackGrowthCurveEndArcShape bool
	for _, order := range stage.ShiftedLeftStackGrowthCurveEndArcShape_stagedOrder {
		if !foundShiftedLeftStackGrowthCurveEndArcShape || order > maxShiftedLeftStackGrowthCurveEndArcShapeOrder {
			maxShiftedLeftStackGrowthCurveEndArcShapeOrder = order
			foundShiftedLeftStackGrowthCurveEndArcShape = true
		}
	}
	if foundShiftedLeftStackGrowthCurveEndArcShape {
		stage.ShiftedLeftStackGrowthCurveEndArcShapeOrder = maxShiftedLeftStackGrowthCurveEndArcShapeOrder + 1
	} else {
		stage.ShiftedLeftStackGrowthCurveEndArcShapeOrder = 0
	}

	var maxShiftedLeftStackGrowthCurveStartArcShapeOrder uint
	var foundShiftedLeftStackGrowthCurveStartArcShape bool
	for _, order := range stage.ShiftedLeftStackGrowthCurveStartArcShape_stagedOrder {
		if !foundShiftedLeftStackGrowthCurveStartArcShape || order > maxShiftedLeftStackGrowthCurveStartArcShapeOrder {
			maxShiftedLeftStackGrowthCurveStartArcShapeOrder = order
			foundShiftedLeftStackGrowthCurveStartArcShape = true
		}
	}
	if foundShiftedLeftStackGrowthCurveStartArcShape {
		stage.ShiftedLeftStackGrowthCurveStartArcShapeOrder = maxShiftedLeftStackGrowthCurveStartArcShapeOrder + 1
	} else {
		stage.ShiftedLeftStackGrowthCurveStartArcShapeOrder = 0
	}

	var maxShiftedLeftStackNormalVectorOrder uint
	var foundShiftedLeftStackNormalVector bool
	for _, order := range stage.ShiftedLeftStackNormalVector_stagedOrder {
		if !foundShiftedLeftStackNormalVector || order > maxShiftedLeftStackNormalVectorOrder {
			maxShiftedLeftStackNormalVectorOrder = order
			foundShiftedLeftStackNormalVector = true
		}
	}
	if foundShiftedLeftStackNormalVector {
		stage.ShiftedLeftStackNormalVectorOrder = maxShiftedLeftStackNormalVectorOrder + 1
	} else {
		stage.ShiftedLeftStackNormalVectorOrder = 0
	}

	var maxShiftedLeftStackOfGrowthCurveOrder uint
	var foundShiftedLeftStackOfGrowthCurve bool
	for _, order := range stage.ShiftedLeftStackOfGrowthCurve_stagedOrder {
		if !foundShiftedLeftStackOfGrowthCurve || order > maxShiftedLeftStackOfGrowthCurveOrder {
			maxShiftedLeftStackOfGrowthCurveOrder = order
			foundShiftedLeftStackOfGrowthCurve = true
		}
	}
	if foundShiftedLeftStackOfGrowthCurve {
		stage.ShiftedLeftStackOfGrowthCurveOrder = maxShiftedLeftStackOfGrowthCurveOrder + 1
	} else {
		stage.ShiftedLeftStackOfGrowthCurveOrder = 0
	}

	var maxShiftedLeftStackOfNormalVectorOrder uint
	var foundShiftedLeftStackOfNormalVector bool
	for _, order := range stage.ShiftedLeftStackOfNormalVector_stagedOrder {
		if !foundShiftedLeftStackOfNormalVector || order > maxShiftedLeftStackOfNormalVectorOrder {
			maxShiftedLeftStackOfNormalVectorOrder = order
			foundShiftedLeftStackOfNormalVector = true
		}
	}
	if foundShiftedLeftStackOfNormalVector {
		stage.ShiftedLeftStackOfNormalVectorOrder = maxShiftedLeftStackOfNormalVectorOrder + 1
	} else {
		stage.ShiftedLeftStackOfNormalVectorOrder = 0
	}

	var maxStackGrowthCurveEndArcShapeOrder uint
	var foundStackGrowthCurveEndArcShape bool
	for _, order := range stage.StackGrowthCurveEndArcShape_stagedOrder {
		if !foundStackGrowthCurveEndArcShape || order > maxStackGrowthCurveEndArcShapeOrder {
			maxStackGrowthCurveEndArcShapeOrder = order
			foundStackGrowthCurveEndArcShape = true
		}
	}
	if foundStackGrowthCurveEndArcShape {
		stage.StackGrowthCurveEndArcShapeOrder = maxStackGrowthCurveEndArcShapeOrder + 1
	} else {
		stage.StackGrowthCurveEndArcShapeOrder = 0
	}

	var maxStackGrowthCurveStartArcShapeOrder uint
	var foundStackGrowthCurveStartArcShape bool
	for _, order := range stage.StackGrowthCurveStartArcShape_stagedOrder {
		if !foundStackGrowthCurveStartArcShape || order > maxStackGrowthCurveStartArcShapeOrder {
			maxStackGrowthCurveStartArcShapeOrder = order
			foundStackGrowthCurveStartArcShape = true
		}
	}
	if foundStackGrowthCurveStartArcShape {
		stage.StackGrowthCurveStartArcShapeOrder = maxStackGrowthCurveStartArcShapeOrder + 1
	} else {
		stage.StackGrowthCurveStartArcShapeOrder = 0
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

	var maxTopEndArcShapeOrder uint
	var foundTopEndArcShape bool
	for _, order := range stage.TopEndArcShape_stagedOrder {
		if !foundTopEndArcShape || order > maxTopEndArcShapeOrder {
			maxTopEndArcShapeOrder = order
			foundTopEndArcShape = true
		}
	}
	if foundTopEndArcShape {
		stage.TopEndArcShapeOrder = maxTopEndArcShapeOrder + 1
	} else {
		stage.TopEndArcShapeOrder = 0
	}

	var maxTopEndArcShapeGridOrder uint
	var foundTopEndArcShapeGrid bool
	for _, order := range stage.TopEndArcShapeGrid_stagedOrder {
		if !foundTopEndArcShapeGrid || order > maxTopEndArcShapeGridOrder {
			maxTopEndArcShapeGridOrder = order
			foundTopEndArcShapeGrid = true
		}
	}
	if foundTopEndArcShapeGrid {
		stage.TopEndArcShapeGridOrder = maxTopEndArcShapeGridOrder + 1
	} else {
		stage.TopEndArcShapeGridOrder = 0
	}

	var maxTopGrowthCurve2DOrder uint
	var foundTopGrowthCurve2D bool
	for _, order := range stage.TopGrowthCurve2D_stagedOrder {
		if !foundTopGrowthCurve2D || order > maxTopGrowthCurve2DOrder {
			maxTopGrowthCurve2DOrder = order
			foundTopGrowthCurve2D = true
		}
	}
	if foundTopGrowthCurve2D {
		stage.TopGrowthCurve2DOrder = maxTopGrowthCurve2DOrder + 1
	} else {
		stage.TopGrowthCurve2DOrder = 0
	}

	var maxTopStackGrowthCurveEndArcShapeOrder uint
	var foundTopStackGrowthCurveEndArcShape bool
	for _, order := range stage.TopStackGrowthCurveEndArcShape_stagedOrder {
		if !foundTopStackGrowthCurveEndArcShape || order > maxTopStackGrowthCurveEndArcShapeOrder {
			maxTopStackGrowthCurveEndArcShapeOrder = order
			foundTopStackGrowthCurveEndArcShape = true
		}
	}
	if foundTopStackGrowthCurveEndArcShape {
		stage.TopStackGrowthCurveEndArcShapeOrder = maxTopStackGrowthCurveEndArcShapeOrder + 1
	} else {
		stage.TopStackGrowthCurveEndArcShapeOrder = 0
	}

	var maxTopStackGrowthCurveStartArcShapeOrder uint
	var foundTopStackGrowthCurveStartArcShape bool
	for _, order := range stage.TopStackGrowthCurveStartArcShape_stagedOrder {
		if !foundTopStackGrowthCurveStartArcShape || order > maxTopStackGrowthCurveStartArcShapeOrder {
			maxTopStackGrowthCurveStartArcShapeOrder = order
			foundTopStackGrowthCurveStartArcShape = true
		}
	}
	if foundTopStackGrowthCurveStartArcShape {
		stage.TopStackGrowthCurveStartArcShapeOrder = maxTopStackGrowthCurveStartArcShapeOrder + 1
	} else {
		stage.TopStackGrowthCurveStartArcShapeOrder = 0
	}

	var maxTopStackOfGrowthCurveOrder uint
	var foundTopStackOfGrowthCurve bool
	for _, order := range stage.TopStackOfGrowthCurve_stagedOrder {
		if !foundTopStackOfGrowthCurve || order > maxTopStackOfGrowthCurveOrder {
			maxTopStackOfGrowthCurveOrder = order
			foundTopStackOfGrowthCurve = true
		}
	}
	if foundTopStackOfGrowthCurve {
		stage.TopStackOfGrowthCurveOrder = maxTopStackOfGrowthCurveOrder + 1
	} else {
		stage.TopStackOfGrowthCurveOrder = 0
	}

	var maxTopStartArcShapeOrder uint
	var foundTopStartArcShape bool
	for _, order := range stage.TopStartArcShape_stagedOrder {
		if !foundTopStartArcShape || order > maxTopStartArcShapeOrder {
			maxTopStartArcShapeOrder = order
			foundTopStartArcShape = true
		}
	}
	if foundTopStartArcShape {
		stage.TopStartArcShapeOrder = maxTopStartArcShapeOrder + 1
	} else {
		stage.TopStartArcShapeOrder = 0
	}

	var maxTopStartArcShapeGridOrder uint
	var foundTopStartArcShapeGrid bool
	for _, order := range stage.TopStartArcShapeGrid_stagedOrder {
		if !foundTopStartArcShapeGrid || order > maxTopStartArcShapeGridOrder {
			maxTopStartArcShapeGridOrder = order
			foundTopStartArcShapeGrid = true
		}
	}
	if foundTopStartArcShapeGrid {
		stage.TopStartArcShapeGridOrder = maxTopStartArcShapeGridOrder + 1
	} else {
		stage.TopStartArcShapeGridOrder = 0
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
	case *ArcNormalVectorShape:
		tmp := GetStructInstancesByOrder(stage.ArcNormalVectorShapes, stage.ArcNormalVectorShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ArcNormalVectorShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ArcNormalVectorShapeGrid:
		tmp := GetStructInstancesByOrder(stage.ArcNormalVectorShapeGrids, stage.ArcNormalVectorShapeGrid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ArcNormalVectorShapeGrid implements.
			res = append(res, any(v).(T))
		}
		return res
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
	case *BaseVectorShape:
		tmp := GetStructInstancesByOrder(stage.BaseVectorShapes, stage.BaseVectorShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *BaseVectorShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *BaseVectorShapeGrid:
		tmp := GetStructInstancesByOrder(stage.BaseVectorShapeGrids, stage.BaseVectorShapeGrid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *BaseVectorShapeGrid implements.
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
	case *GrowthCurve2D:
		tmp := GetStructInstancesByOrder(stage.GrowthCurve2Ds, stage.GrowthCurve2D_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GrowthCurve2D implements.
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
	case *ShiftedLeftStackGrowthCurveEndArcShape:
		tmp := GetStructInstancesByOrder(stage.ShiftedLeftStackGrowthCurveEndArcShapes, stage.ShiftedLeftStackGrowthCurveEndArcShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ShiftedLeftStackGrowthCurveEndArcShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ShiftedLeftStackGrowthCurveStartArcShape:
		tmp := GetStructInstancesByOrder(stage.ShiftedLeftStackGrowthCurveStartArcShapes, stage.ShiftedLeftStackGrowthCurveStartArcShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ShiftedLeftStackGrowthCurveStartArcShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ShiftedLeftStackNormalVector:
		tmp := GetStructInstancesByOrder(stage.ShiftedLeftStackNormalVectors, stage.ShiftedLeftStackNormalVector_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ShiftedLeftStackNormalVector implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ShiftedLeftStackOfGrowthCurve:
		tmp := GetStructInstancesByOrder(stage.ShiftedLeftStackOfGrowthCurves, stage.ShiftedLeftStackOfGrowthCurve_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ShiftedLeftStackOfGrowthCurve implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ShiftedLeftStackOfNormalVector:
		tmp := GetStructInstancesByOrder(stage.ShiftedLeftStackOfNormalVectors, stage.ShiftedLeftStackOfNormalVector_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ShiftedLeftStackOfNormalVector implements.
			res = append(res, any(v).(T))
		}
		return res
	case *StackGrowthCurveEndArcShape:
		tmp := GetStructInstancesByOrder(stage.StackGrowthCurveEndArcShapes, stage.StackGrowthCurveEndArcShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *StackGrowthCurveEndArcShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *StackGrowthCurveStartArcShape:
		tmp := GetStructInstancesByOrder(stage.StackGrowthCurveStartArcShapes, stage.StackGrowthCurveStartArcShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *StackGrowthCurveStartArcShape implements.
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
	case *TopEndArcShape:
		tmp := GetStructInstancesByOrder(stage.TopEndArcShapes, stage.TopEndArcShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TopEndArcShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TopEndArcShapeGrid:
		tmp := GetStructInstancesByOrder(stage.TopEndArcShapeGrids, stage.TopEndArcShapeGrid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TopEndArcShapeGrid implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TopGrowthCurve2D:
		tmp := GetStructInstancesByOrder(stage.TopGrowthCurve2Ds, stage.TopGrowthCurve2D_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TopGrowthCurve2D implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TopStackGrowthCurveEndArcShape:
		tmp := GetStructInstancesByOrder(stage.TopStackGrowthCurveEndArcShapes, stage.TopStackGrowthCurveEndArcShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TopStackGrowthCurveEndArcShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TopStackGrowthCurveStartArcShape:
		tmp := GetStructInstancesByOrder(stage.TopStackGrowthCurveStartArcShapes, stage.TopStackGrowthCurveStartArcShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TopStackGrowthCurveStartArcShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TopStackOfGrowthCurve:
		tmp := GetStructInstancesByOrder(stage.TopStackOfGrowthCurves, stage.TopStackOfGrowthCurve_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TopStackOfGrowthCurve implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TopStartArcShape:
		tmp := GetStructInstancesByOrder(stage.TopStartArcShapes, stage.TopStartArcShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TopStartArcShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TopStartArcShapeGrid:
		tmp := GetStructInstancesByOrder(stage.TopStartArcShapeGrids, stage.TopStartArcShapeGrid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TopStartArcShapeGrid implements.
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
	case "ArcNormalVectorShape":
		res = GetNamedStructInstances(stage.ArcNormalVectorShapes, stage.ArcNormalVectorShape_stagedOrder)
	case "ArcNormalVectorShapeGrid":
		res = GetNamedStructInstances(stage.ArcNormalVectorShapeGrids, stage.ArcNormalVectorShapeGrid_stagedOrder)
	case "AxesShape":
		res = GetNamedStructInstances(stage.AxesShapes, stage.AxesShape_stagedOrder)
	case "BaseVectorShape":
		res = GetNamedStructInstances(stage.BaseVectorShapes, stage.BaseVectorShape_stagedOrder)
	case "BaseVectorShapeGrid":
		res = GetNamedStructInstances(stage.BaseVectorShapeGrids, stage.BaseVectorShapeGrid_stagedOrder)
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
	case "GrowthCurve2D":
		res = GetNamedStructInstances(stage.GrowthCurve2Ds, stage.GrowthCurve2D_stagedOrder)
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
	case "ShiftedLeftStackGrowthCurveEndArcShape":
		res = GetNamedStructInstances(stage.ShiftedLeftStackGrowthCurveEndArcShapes, stage.ShiftedLeftStackGrowthCurveEndArcShape_stagedOrder)
	case "ShiftedLeftStackGrowthCurveStartArcShape":
		res = GetNamedStructInstances(stage.ShiftedLeftStackGrowthCurveStartArcShapes, stage.ShiftedLeftStackGrowthCurveStartArcShape_stagedOrder)
	case "ShiftedLeftStackNormalVector":
		res = GetNamedStructInstances(stage.ShiftedLeftStackNormalVectors, stage.ShiftedLeftStackNormalVector_stagedOrder)
	case "ShiftedLeftStackOfGrowthCurve":
		res = GetNamedStructInstances(stage.ShiftedLeftStackOfGrowthCurves, stage.ShiftedLeftStackOfGrowthCurve_stagedOrder)
	case "ShiftedLeftStackOfNormalVector":
		res = GetNamedStructInstances(stage.ShiftedLeftStackOfNormalVectors, stage.ShiftedLeftStackOfNormalVector_stagedOrder)
	case "StackGrowthCurveEndArcShape":
		res = GetNamedStructInstances(stage.StackGrowthCurveEndArcShapes, stage.StackGrowthCurveEndArcShape_stagedOrder)
	case "StackGrowthCurveStartArcShape":
		res = GetNamedStructInstances(stage.StackGrowthCurveStartArcShapes, stage.StackGrowthCurveStartArcShape_stagedOrder)
	case "StackOfGrowthCurve":
		res = GetNamedStructInstances(stage.StackOfGrowthCurves, stage.StackOfGrowthCurve_stagedOrder)
	case "StartArcShape":
		res = GetNamedStructInstances(stage.StartArcShapes, stage.StartArcShape_stagedOrder)
	case "StartArcShapeGrid":
		res = GetNamedStructInstances(stage.StartArcShapeGrids, stage.StartArcShapeGrid_stagedOrder)
	case "TopEndArcShape":
		res = GetNamedStructInstances(stage.TopEndArcShapes, stage.TopEndArcShape_stagedOrder)
	case "TopEndArcShapeGrid":
		res = GetNamedStructInstances(stage.TopEndArcShapeGrids, stage.TopEndArcShapeGrid_stagedOrder)
	case "TopGrowthCurve2D":
		res = GetNamedStructInstances(stage.TopGrowthCurve2Ds, stage.TopGrowthCurve2D_stagedOrder)
	case "TopStackGrowthCurveEndArcShape":
		res = GetNamedStructInstances(stage.TopStackGrowthCurveEndArcShapes, stage.TopStackGrowthCurveEndArcShape_stagedOrder)
	case "TopStackGrowthCurveStartArcShape":
		res = GetNamedStructInstances(stage.TopStackGrowthCurveStartArcShapes, stage.TopStackGrowthCurveStartArcShape_stagedOrder)
	case "TopStackOfGrowthCurve":
		res = GetNamedStructInstances(stage.TopStackOfGrowthCurves, stage.TopStackOfGrowthCurve_stagedOrder)
	case "TopStartArcShape":
		res = GetNamedStructInstances(stage.TopStartArcShapes, stage.TopStartArcShape_stagedOrder)
	case "TopStartArcShapeGrid":
		res = GetNamedStructInstances(stage.TopStartArcShapeGrids, stage.TopStartArcShapeGrid_stagedOrder)
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
	CommitArcNormalVectorShape(arcnormalvectorshape *ArcNormalVectorShape)
	CheckoutArcNormalVectorShape(arcnormalvectorshape *ArcNormalVectorShape)
	CommitArcNormalVectorShapeGrid(arcnormalvectorshapegrid *ArcNormalVectorShapeGrid)
	CheckoutArcNormalVectorShapeGrid(arcnormalvectorshapegrid *ArcNormalVectorShapeGrid)
	CommitAxesShape(axesshape *AxesShape)
	CheckoutAxesShape(axesshape *AxesShape)
	CommitBaseVectorShape(basevectorshape *BaseVectorShape)
	CheckoutBaseVectorShape(basevectorshape *BaseVectorShape)
	CommitBaseVectorShapeGrid(basevectorshapegrid *BaseVectorShapeGrid)
	CheckoutBaseVectorShapeGrid(basevectorshapegrid *BaseVectorShapeGrid)
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
	CommitGrowthCurve2D(growthcurve2d *GrowthCurve2D)
	CheckoutGrowthCurve2D(growthcurve2d *GrowthCurve2D)
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
	CommitShiftedLeftStackGrowthCurveEndArcShape(shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape)
	CheckoutShiftedLeftStackGrowthCurveEndArcShape(shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape)
	CommitShiftedLeftStackGrowthCurveStartArcShape(shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape)
	CheckoutShiftedLeftStackGrowthCurveStartArcShape(shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape)
	CommitShiftedLeftStackNormalVector(shiftedleftstacknormalvector *ShiftedLeftStackNormalVector)
	CheckoutShiftedLeftStackNormalVector(shiftedleftstacknormalvector *ShiftedLeftStackNormalVector)
	CommitShiftedLeftStackOfGrowthCurve(shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve)
	CheckoutShiftedLeftStackOfGrowthCurve(shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve)
	CommitShiftedLeftStackOfNormalVector(shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector)
	CheckoutShiftedLeftStackOfNormalVector(shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector)
	CommitStackGrowthCurveEndArcShape(stackgrowthcurveendarcshape *StackGrowthCurveEndArcShape)
	CheckoutStackGrowthCurveEndArcShape(stackgrowthcurveendarcshape *StackGrowthCurveEndArcShape)
	CommitStackGrowthCurveStartArcShape(stackgrowthcurvestartarcshape *StackGrowthCurveStartArcShape)
	CheckoutStackGrowthCurveStartArcShape(stackgrowthcurvestartarcshape *StackGrowthCurveStartArcShape)
	CommitStackOfGrowthCurve(stackofgrowthcurve *StackOfGrowthCurve)
	CheckoutStackOfGrowthCurve(stackofgrowthcurve *StackOfGrowthCurve)
	CommitStartArcShape(startarcshape *StartArcShape)
	CheckoutStartArcShape(startarcshape *StartArcShape)
	CommitStartArcShapeGrid(startarcshapegrid *StartArcShapeGrid)
	CheckoutStartArcShapeGrid(startarcshapegrid *StartArcShapeGrid)
	CommitTopEndArcShape(topendarcshape *TopEndArcShape)
	CheckoutTopEndArcShape(topendarcshape *TopEndArcShape)
	CommitTopEndArcShapeGrid(topendarcshapegrid *TopEndArcShapeGrid)
	CheckoutTopEndArcShapeGrid(topendarcshapegrid *TopEndArcShapeGrid)
	CommitTopGrowthCurve2D(topgrowthcurve2d *TopGrowthCurve2D)
	CheckoutTopGrowthCurve2D(topgrowthcurve2d *TopGrowthCurve2D)
	CommitTopStackGrowthCurveEndArcShape(topstackgrowthcurveendarcshape *TopStackGrowthCurveEndArcShape)
	CheckoutTopStackGrowthCurveEndArcShape(topstackgrowthcurveendarcshape *TopStackGrowthCurveEndArcShape)
	CommitTopStackGrowthCurveStartArcShape(topstackgrowthcurvestartarcshape *TopStackGrowthCurveStartArcShape)
	CheckoutTopStackGrowthCurveStartArcShape(topstackgrowthcurvestartarcshape *TopStackGrowthCurveStartArcShape)
	CommitTopStackOfGrowthCurve(topstackofgrowthcurve *TopStackOfGrowthCurve)
	CheckoutTopStackOfGrowthCurve(topstackofgrowthcurve *TopStackOfGrowthCurve)
	CommitTopStartArcShape(topstartarcshape *TopStartArcShape)
	CheckoutTopStartArcShape(topstartarcshape *TopStartArcShape)
	CommitTopStartArcShapeGrid(topstartarcshapegrid *TopStartArcShapeGrid)
	CheckoutTopStartArcShapeGrid(topstartarcshapegrid *TopStartArcShapeGrid)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		ArcNormalVectorShapes:           make(map[*ArcNormalVectorShape]struct{}),
		ArcNormalVectorShapes_mapString: make(map[string]*ArcNormalVectorShape),

		ArcNormalVectorShapeGrids:           make(map[*ArcNormalVectorShapeGrid]struct{}),
		ArcNormalVectorShapeGrids_mapString: make(map[string]*ArcNormalVectorShapeGrid),

		AxesShapes:           make(map[*AxesShape]struct{}),
		AxesShapes_mapString: make(map[string]*AxesShape),

		BaseVectorShapes:           make(map[*BaseVectorShape]struct{}),
		BaseVectorShapes_mapString: make(map[string]*BaseVectorShape),

		BaseVectorShapeGrids:           make(map[*BaseVectorShapeGrid]struct{}),
		BaseVectorShapeGrids_mapString: make(map[string]*BaseVectorShapeGrid),

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

		GrowthCurve2Ds:           make(map[*GrowthCurve2D]struct{}),
		GrowthCurve2Ds_mapString: make(map[string]*GrowthCurve2D),

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

		ShiftedLeftStackGrowthCurveEndArcShapes:           make(map[*ShiftedLeftStackGrowthCurveEndArcShape]struct{}),
		ShiftedLeftStackGrowthCurveEndArcShapes_mapString: make(map[string]*ShiftedLeftStackGrowthCurveEndArcShape),

		ShiftedLeftStackGrowthCurveStartArcShapes:           make(map[*ShiftedLeftStackGrowthCurveStartArcShape]struct{}),
		ShiftedLeftStackGrowthCurveStartArcShapes_mapString: make(map[string]*ShiftedLeftStackGrowthCurveStartArcShape),

		ShiftedLeftStackNormalVectors:           make(map[*ShiftedLeftStackNormalVector]struct{}),
		ShiftedLeftStackNormalVectors_mapString: make(map[string]*ShiftedLeftStackNormalVector),

		ShiftedLeftStackOfGrowthCurves:           make(map[*ShiftedLeftStackOfGrowthCurve]struct{}),
		ShiftedLeftStackOfGrowthCurves_mapString: make(map[string]*ShiftedLeftStackOfGrowthCurve),

		ShiftedLeftStackOfNormalVectors:           make(map[*ShiftedLeftStackOfNormalVector]struct{}),
		ShiftedLeftStackOfNormalVectors_mapString: make(map[string]*ShiftedLeftStackOfNormalVector),

		StackGrowthCurveEndArcShapes:           make(map[*StackGrowthCurveEndArcShape]struct{}),
		StackGrowthCurveEndArcShapes_mapString: make(map[string]*StackGrowthCurveEndArcShape),

		StackGrowthCurveStartArcShapes:           make(map[*StackGrowthCurveStartArcShape]struct{}),
		StackGrowthCurveStartArcShapes_mapString: make(map[string]*StackGrowthCurveStartArcShape),

		StackOfGrowthCurves:           make(map[*StackOfGrowthCurve]struct{}),
		StackOfGrowthCurves_mapString: make(map[string]*StackOfGrowthCurve),

		StartArcShapes:           make(map[*StartArcShape]struct{}),
		StartArcShapes_mapString: make(map[string]*StartArcShape),

		StartArcShapeGrids:           make(map[*StartArcShapeGrid]struct{}),
		StartArcShapeGrids_mapString: make(map[string]*StartArcShapeGrid),

		TopEndArcShapes:           make(map[*TopEndArcShape]struct{}),
		TopEndArcShapes_mapString: make(map[string]*TopEndArcShape),

		TopEndArcShapeGrids:           make(map[*TopEndArcShapeGrid]struct{}),
		TopEndArcShapeGrids_mapString: make(map[string]*TopEndArcShapeGrid),

		TopGrowthCurve2Ds:           make(map[*TopGrowthCurve2D]struct{}),
		TopGrowthCurve2Ds_mapString: make(map[string]*TopGrowthCurve2D),

		TopStackGrowthCurveEndArcShapes:           make(map[*TopStackGrowthCurveEndArcShape]struct{}),
		TopStackGrowthCurveEndArcShapes_mapString: make(map[string]*TopStackGrowthCurveEndArcShape),

		TopStackGrowthCurveStartArcShapes:           make(map[*TopStackGrowthCurveStartArcShape]struct{}),
		TopStackGrowthCurveStartArcShapes_mapString: make(map[string]*TopStackGrowthCurveStartArcShape),

		TopStackOfGrowthCurves:           make(map[*TopStackOfGrowthCurve]struct{}),
		TopStackOfGrowthCurves_mapString: make(map[string]*TopStackOfGrowthCurve),

		TopStartArcShapes:           make(map[*TopStartArcShape]struct{}),
		TopStartArcShapes_mapString: make(map[string]*TopStartArcShape),

		TopStartArcShapeGrids:           make(map[*TopStartArcShapeGrid]struct{}),
		TopStartArcShapeGrids_mapString: make(map[string]*TopStartArcShapeGrid),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		ArcNormalVectorShape_stagedOrder: make(map[*ArcNormalVectorShape]uint),
		ArcNormalVectorShape_orderStaged: make(map[uint]*ArcNormalVectorShape),
		ArcNormalVectorShapes_reference:  make(map[*ArcNormalVectorShape]*ArcNormalVectorShape),

		ArcNormalVectorShapeGrid_stagedOrder: make(map[*ArcNormalVectorShapeGrid]uint),
		ArcNormalVectorShapeGrid_orderStaged: make(map[uint]*ArcNormalVectorShapeGrid),
		ArcNormalVectorShapeGrids_reference:  make(map[*ArcNormalVectorShapeGrid]*ArcNormalVectorShapeGrid),

		AxesShape_stagedOrder: make(map[*AxesShape]uint),
		AxesShape_orderStaged: make(map[uint]*AxesShape),
		AxesShapes_reference:  make(map[*AxesShape]*AxesShape),

		BaseVectorShape_stagedOrder: make(map[*BaseVectorShape]uint),
		BaseVectorShape_orderStaged: make(map[uint]*BaseVectorShape),
		BaseVectorShapes_reference:  make(map[*BaseVectorShape]*BaseVectorShape),

		BaseVectorShapeGrid_stagedOrder: make(map[*BaseVectorShapeGrid]uint),
		BaseVectorShapeGrid_orderStaged: make(map[uint]*BaseVectorShapeGrid),
		BaseVectorShapeGrids_reference:  make(map[*BaseVectorShapeGrid]*BaseVectorShapeGrid),

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

		GrowthCurve2D_stagedOrder: make(map[*GrowthCurve2D]uint),
		GrowthCurve2D_orderStaged: make(map[uint]*GrowthCurve2D),
		GrowthCurve2Ds_reference:  make(map[*GrowthCurve2D]*GrowthCurve2D),

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

		ShiftedLeftStackGrowthCurveEndArcShape_stagedOrder: make(map[*ShiftedLeftStackGrowthCurveEndArcShape]uint),
		ShiftedLeftStackGrowthCurveEndArcShape_orderStaged: make(map[uint]*ShiftedLeftStackGrowthCurveEndArcShape),
		ShiftedLeftStackGrowthCurveEndArcShapes_reference:  make(map[*ShiftedLeftStackGrowthCurveEndArcShape]*ShiftedLeftStackGrowthCurveEndArcShape),

		ShiftedLeftStackGrowthCurveStartArcShape_stagedOrder: make(map[*ShiftedLeftStackGrowthCurveStartArcShape]uint),
		ShiftedLeftStackGrowthCurveStartArcShape_orderStaged: make(map[uint]*ShiftedLeftStackGrowthCurveStartArcShape),
		ShiftedLeftStackGrowthCurveStartArcShapes_reference:  make(map[*ShiftedLeftStackGrowthCurveStartArcShape]*ShiftedLeftStackGrowthCurveStartArcShape),

		ShiftedLeftStackNormalVector_stagedOrder: make(map[*ShiftedLeftStackNormalVector]uint),
		ShiftedLeftStackNormalVector_orderStaged: make(map[uint]*ShiftedLeftStackNormalVector),
		ShiftedLeftStackNormalVectors_reference:  make(map[*ShiftedLeftStackNormalVector]*ShiftedLeftStackNormalVector),

		ShiftedLeftStackOfGrowthCurve_stagedOrder: make(map[*ShiftedLeftStackOfGrowthCurve]uint),
		ShiftedLeftStackOfGrowthCurve_orderStaged: make(map[uint]*ShiftedLeftStackOfGrowthCurve),
		ShiftedLeftStackOfGrowthCurves_reference:  make(map[*ShiftedLeftStackOfGrowthCurve]*ShiftedLeftStackOfGrowthCurve),

		ShiftedLeftStackOfNormalVector_stagedOrder: make(map[*ShiftedLeftStackOfNormalVector]uint),
		ShiftedLeftStackOfNormalVector_orderStaged: make(map[uint]*ShiftedLeftStackOfNormalVector),
		ShiftedLeftStackOfNormalVectors_reference:  make(map[*ShiftedLeftStackOfNormalVector]*ShiftedLeftStackOfNormalVector),

		StackGrowthCurveEndArcShape_stagedOrder: make(map[*StackGrowthCurveEndArcShape]uint),
		StackGrowthCurveEndArcShape_orderStaged: make(map[uint]*StackGrowthCurveEndArcShape),
		StackGrowthCurveEndArcShapes_reference:  make(map[*StackGrowthCurveEndArcShape]*StackGrowthCurveEndArcShape),

		StackGrowthCurveStartArcShape_stagedOrder: make(map[*StackGrowthCurveStartArcShape]uint),
		StackGrowthCurveStartArcShape_orderStaged: make(map[uint]*StackGrowthCurveStartArcShape),
		StackGrowthCurveStartArcShapes_reference:  make(map[*StackGrowthCurveStartArcShape]*StackGrowthCurveStartArcShape),

		StackOfGrowthCurve_stagedOrder: make(map[*StackOfGrowthCurve]uint),
		StackOfGrowthCurve_orderStaged: make(map[uint]*StackOfGrowthCurve),
		StackOfGrowthCurves_reference:  make(map[*StackOfGrowthCurve]*StackOfGrowthCurve),

		StartArcShape_stagedOrder: make(map[*StartArcShape]uint),
		StartArcShape_orderStaged: make(map[uint]*StartArcShape),
		StartArcShapes_reference:  make(map[*StartArcShape]*StartArcShape),

		StartArcShapeGrid_stagedOrder: make(map[*StartArcShapeGrid]uint),
		StartArcShapeGrid_orderStaged: make(map[uint]*StartArcShapeGrid),
		StartArcShapeGrids_reference:  make(map[*StartArcShapeGrid]*StartArcShapeGrid),

		TopEndArcShape_stagedOrder: make(map[*TopEndArcShape]uint),
		TopEndArcShape_orderStaged: make(map[uint]*TopEndArcShape),
		TopEndArcShapes_reference:  make(map[*TopEndArcShape]*TopEndArcShape),

		TopEndArcShapeGrid_stagedOrder: make(map[*TopEndArcShapeGrid]uint),
		TopEndArcShapeGrid_orderStaged: make(map[uint]*TopEndArcShapeGrid),
		TopEndArcShapeGrids_reference:  make(map[*TopEndArcShapeGrid]*TopEndArcShapeGrid),

		TopGrowthCurve2D_stagedOrder: make(map[*TopGrowthCurve2D]uint),
		TopGrowthCurve2D_orderStaged: make(map[uint]*TopGrowthCurve2D),
		TopGrowthCurve2Ds_reference:  make(map[*TopGrowthCurve2D]*TopGrowthCurve2D),

		TopStackGrowthCurveEndArcShape_stagedOrder: make(map[*TopStackGrowthCurveEndArcShape]uint),
		TopStackGrowthCurveEndArcShape_orderStaged: make(map[uint]*TopStackGrowthCurveEndArcShape),
		TopStackGrowthCurveEndArcShapes_reference:  make(map[*TopStackGrowthCurveEndArcShape]*TopStackGrowthCurveEndArcShape),

		TopStackGrowthCurveStartArcShape_stagedOrder: make(map[*TopStackGrowthCurveStartArcShape]uint),
		TopStackGrowthCurveStartArcShape_orderStaged: make(map[uint]*TopStackGrowthCurveStartArcShape),
		TopStackGrowthCurveStartArcShapes_reference:  make(map[*TopStackGrowthCurveStartArcShape]*TopStackGrowthCurveStartArcShape),

		TopStackOfGrowthCurve_stagedOrder: make(map[*TopStackOfGrowthCurve]uint),
		TopStackOfGrowthCurve_orderStaged: make(map[uint]*TopStackOfGrowthCurve),
		TopStackOfGrowthCurves_reference:  make(map[*TopStackOfGrowthCurve]*TopStackOfGrowthCurve),

		TopStartArcShape_stagedOrder: make(map[*TopStartArcShape]uint),
		TopStartArcShape_orderStaged: make(map[uint]*TopStartArcShape),
		TopStartArcShapes_reference:  make(map[*TopStartArcShape]*TopStartArcShape),

		TopStartArcShapeGrid_stagedOrder: make(map[*TopStartArcShapeGrid]uint),
		TopStartArcShapeGrid_orderStaged: make(map[uint]*TopStartArcShapeGrid),
		TopStartArcShapeGrids_reference:  make(map[*TopStartArcShapeGrid]*TopStartArcShapeGrid),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"ArcNormalVectorShape": &ArcNormalVectorShapeUnmarshaller{},

			"ArcNormalVectorShapeGrid": &ArcNormalVectorShapeGridUnmarshaller{},

			"AxesShape": &AxesShapeUnmarshaller{},

			"BaseVectorShape": &BaseVectorShapeUnmarshaller{},

			"BaseVectorShapeGrid": &BaseVectorShapeGridUnmarshaller{},

			"CircleGridShape": &CircleGridShapeUnmarshaller{},

			"EndArcShape": &EndArcShapeUnmarshaller{},

			"EndArcShapeGrid": &EndArcShapeGridUnmarshaller{},

			"ExplanationTextShape": &ExplanationTextShapeUnmarshaller{},

			"GridPathShape": &GridPathShapeUnmarshaller{},

			"GrowthCurve2D": &GrowthCurve2DUnmarshaller{},

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

			"ShiftedLeftStackGrowthCurveEndArcShape": &ShiftedLeftStackGrowthCurveEndArcShapeUnmarshaller{},

			"ShiftedLeftStackGrowthCurveStartArcShape": &ShiftedLeftStackGrowthCurveStartArcShapeUnmarshaller{},

			"ShiftedLeftStackNormalVector": &ShiftedLeftStackNormalVectorUnmarshaller{},

			"ShiftedLeftStackOfGrowthCurve": &ShiftedLeftStackOfGrowthCurveUnmarshaller{},

			"ShiftedLeftStackOfNormalVector": &ShiftedLeftStackOfNormalVectorUnmarshaller{},

			"StackGrowthCurveEndArcShape": &StackGrowthCurveEndArcShapeUnmarshaller{},

			"StackGrowthCurveStartArcShape": &StackGrowthCurveStartArcShapeUnmarshaller{},

			"StackOfGrowthCurve": &StackOfGrowthCurveUnmarshaller{},

			"StartArcShape": &StartArcShapeUnmarshaller{},

			"StartArcShapeGrid": &StartArcShapeGridUnmarshaller{},

			"TopEndArcShape": &TopEndArcShapeUnmarshaller{},

			"TopEndArcShapeGrid": &TopEndArcShapeGridUnmarshaller{},

			"TopGrowthCurve2D": &TopGrowthCurve2DUnmarshaller{},

			"TopStackGrowthCurveEndArcShape": &TopStackGrowthCurveEndArcShapeUnmarshaller{},

			"TopStackGrowthCurveStartArcShape": &TopStackGrowthCurveStartArcShapeUnmarshaller{},

			"TopStackOfGrowthCurve": &TopStackOfGrowthCurveUnmarshaller{},

			"TopStartArcShape": &TopStartArcShapeUnmarshaller{},

			"TopStartArcShapeGrid": &TopStartArcShapeGridUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "ArcNormalVectorShape"},
			{name: "ArcNormalVectorShapeGrid"},
			{name: "AxesShape"},
			{name: "BaseVectorShape"},
			{name: "BaseVectorShapeGrid"},
			{name: "CircleGridShape"},
			{name: "EndArcShape"},
			{name: "EndArcShapeGrid"},
			{name: "ExplanationTextShape"},
			{name: "GridPathShape"},
			{name: "GrowthCurve2D"},
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
			{name: "ShiftedLeftStackGrowthCurveEndArcShape"},
			{name: "ShiftedLeftStackGrowthCurveStartArcShape"},
			{name: "ShiftedLeftStackNormalVector"},
			{name: "ShiftedLeftStackOfGrowthCurve"},
			{name: "ShiftedLeftStackOfNormalVector"},
			{name: "StackGrowthCurveEndArcShape"},
			{name: "StackGrowthCurveStartArcShape"},
			{name: "StackOfGrowthCurve"},
			{name: "StartArcShape"},
			{name: "StartArcShapeGrid"},
			{name: "TopEndArcShape"},
			{name: "TopEndArcShapeGrid"},
			{name: "TopGrowthCurve2D"},
			{name: "TopStackGrowthCurveEndArcShape"},
			{name: "TopStackGrowthCurveStartArcShape"},
			{name: "TopStackOfGrowthCurve"},
			{name: "TopStartArcShape"},
			{name: "TopStartArcShapeGrid"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *ArcNormalVectorShape:
		return stage.ArcNormalVectorShape_stagedOrder[instance]
	case *ArcNormalVectorShapeGrid:
		return stage.ArcNormalVectorShapeGrid_stagedOrder[instance]
	case *AxesShape:
		return stage.AxesShape_stagedOrder[instance]
	case *BaseVectorShape:
		return stage.BaseVectorShape_stagedOrder[instance]
	case *BaseVectorShapeGrid:
		return stage.BaseVectorShapeGrid_stagedOrder[instance]
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
	case *GrowthCurve2D:
		return stage.GrowthCurve2D_stagedOrder[instance]
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
	case *ShiftedLeftStackGrowthCurveEndArcShape:
		return stage.ShiftedLeftStackGrowthCurveEndArcShape_stagedOrder[instance]
	case *ShiftedLeftStackGrowthCurveStartArcShape:
		return stage.ShiftedLeftStackGrowthCurveStartArcShape_stagedOrder[instance]
	case *ShiftedLeftStackNormalVector:
		return stage.ShiftedLeftStackNormalVector_stagedOrder[instance]
	case *ShiftedLeftStackOfGrowthCurve:
		return stage.ShiftedLeftStackOfGrowthCurve_stagedOrder[instance]
	case *ShiftedLeftStackOfNormalVector:
		return stage.ShiftedLeftStackOfNormalVector_stagedOrder[instance]
	case *StackGrowthCurveEndArcShape:
		return stage.StackGrowthCurveEndArcShape_stagedOrder[instance]
	case *StackGrowthCurveStartArcShape:
		return stage.StackGrowthCurveStartArcShape_stagedOrder[instance]
	case *StackOfGrowthCurve:
		return stage.StackOfGrowthCurve_stagedOrder[instance]
	case *StartArcShape:
		return stage.StartArcShape_stagedOrder[instance]
	case *StartArcShapeGrid:
		return stage.StartArcShapeGrid_stagedOrder[instance]
	case *TopEndArcShape:
		return stage.TopEndArcShape_stagedOrder[instance]
	case *TopEndArcShapeGrid:
		return stage.TopEndArcShapeGrid_stagedOrder[instance]
	case *TopGrowthCurve2D:
		return stage.TopGrowthCurve2D_stagedOrder[instance]
	case *TopStackGrowthCurveEndArcShape:
		return stage.TopStackGrowthCurveEndArcShape_stagedOrder[instance]
	case *TopStackGrowthCurveStartArcShape:
		return stage.TopStackGrowthCurveStartArcShape_stagedOrder[instance]
	case *TopStackOfGrowthCurve:
		return stage.TopStackOfGrowthCurve_stagedOrder[instance]
	case *TopStartArcShape:
		return stage.TopStartArcShape_stagedOrder[instance]
	case *TopStartArcShapeGrid:
		return stage.TopStartArcShapeGrid_stagedOrder[instance]
	default:
		return 0 // should not happen
	}
}

func GongGetInstanceFromOrder[Type PointerToGongstruct](stage *Stage, order uint) (res Type) {
	var t Type
	switch any(t).(type) {
	// insertion point for order map initialisations
	case *ArcNormalVectorShape:
		return any(stage.ArcNormalVectorShape_orderStaged[order]).(Type)
	case *ArcNormalVectorShapeGrid:
		return any(stage.ArcNormalVectorShapeGrid_orderStaged[order]).(Type)
	case *AxesShape:
		return any(stage.AxesShape_orderStaged[order]).(Type)
	case *BaseVectorShape:
		return any(stage.BaseVectorShape_orderStaged[order]).(Type)
	case *BaseVectorShapeGrid:
		return any(stage.BaseVectorShapeGrid_orderStaged[order]).(Type)
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
	case *GrowthCurve2D:
		return any(stage.GrowthCurve2D_orderStaged[order]).(Type)
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
	case *ShiftedLeftStackGrowthCurveEndArcShape:
		return any(stage.ShiftedLeftStackGrowthCurveEndArcShape_orderStaged[order]).(Type)
	case *ShiftedLeftStackGrowthCurveStartArcShape:
		return any(stage.ShiftedLeftStackGrowthCurveStartArcShape_orderStaged[order]).(Type)
	case *ShiftedLeftStackNormalVector:
		return any(stage.ShiftedLeftStackNormalVector_orderStaged[order]).(Type)
	case *ShiftedLeftStackOfGrowthCurve:
		return any(stage.ShiftedLeftStackOfGrowthCurve_orderStaged[order]).(Type)
	case *ShiftedLeftStackOfNormalVector:
		return any(stage.ShiftedLeftStackOfNormalVector_orderStaged[order]).(Type)
	case *StackGrowthCurveEndArcShape:
		return any(stage.StackGrowthCurveEndArcShape_orderStaged[order]).(Type)
	case *StackGrowthCurveStartArcShape:
		return any(stage.StackGrowthCurveStartArcShape_orderStaged[order]).(Type)
	case *StackOfGrowthCurve:
		return any(stage.StackOfGrowthCurve_orderStaged[order]).(Type)
	case *StartArcShape:
		return any(stage.StartArcShape_orderStaged[order]).(Type)
	case *StartArcShapeGrid:
		return any(stage.StartArcShapeGrid_orderStaged[order]).(Type)
	case *TopEndArcShape:
		return any(stage.TopEndArcShape_orderStaged[order]).(Type)
	case *TopEndArcShapeGrid:
		return any(stage.TopEndArcShapeGrid_orderStaged[order]).(Type)
	case *TopGrowthCurve2D:
		return any(stage.TopGrowthCurve2D_orderStaged[order]).(Type)
	case *TopStackGrowthCurveEndArcShape:
		return any(stage.TopStackGrowthCurveEndArcShape_orderStaged[order]).(Type)
	case *TopStackGrowthCurveStartArcShape:
		return any(stage.TopStackGrowthCurveStartArcShape_orderStaged[order]).(Type)
	case *TopStackOfGrowthCurve:
		return any(stage.TopStackOfGrowthCurve_orderStaged[order]).(Type)
	case *TopStartArcShape:
		return any(stage.TopStartArcShape_orderStaged[order]).(Type)
	case *TopStartArcShapeGrid:
		return any(stage.TopStartArcShapeGrid_orderStaged[order]).(Type)
	default:
		return // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *ArcNormalVectorShape:
		return stage.ArcNormalVectorShape_stagedOrder[instance]
	case *ArcNormalVectorShapeGrid:
		return stage.ArcNormalVectorShapeGrid_stagedOrder[instance]
	case *AxesShape:
		return stage.AxesShape_stagedOrder[instance]
	case *BaseVectorShape:
		return stage.BaseVectorShape_stagedOrder[instance]
	case *BaseVectorShapeGrid:
		return stage.BaseVectorShapeGrid_stagedOrder[instance]
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
	case *GrowthCurve2D:
		return stage.GrowthCurve2D_stagedOrder[instance]
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
	case *ShiftedLeftStackGrowthCurveEndArcShape:
		return stage.ShiftedLeftStackGrowthCurveEndArcShape_stagedOrder[instance]
	case *ShiftedLeftStackGrowthCurveStartArcShape:
		return stage.ShiftedLeftStackGrowthCurveStartArcShape_stagedOrder[instance]
	case *ShiftedLeftStackNormalVector:
		return stage.ShiftedLeftStackNormalVector_stagedOrder[instance]
	case *ShiftedLeftStackOfGrowthCurve:
		return stage.ShiftedLeftStackOfGrowthCurve_stagedOrder[instance]
	case *ShiftedLeftStackOfNormalVector:
		return stage.ShiftedLeftStackOfNormalVector_stagedOrder[instance]
	case *StackGrowthCurveEndArcShape:
		return stage.StackGrowthCurveEndArcShape_stagedOrder[instance]
	case *StackGrowthCurveStartArcShape:
		return stage.StackGrowthCurveStartArcShape_stagedOrder[instance]
	case *StackOfGrowthCurve:
		return stage.StackOfGrowthCurve_stagedOrder[instance]
	case *StartArcShape:
		return stage.StartArcShape_stagedOrder[instance]
	case *StartArcShapeGrid:
		return stage.StartArcShapeGrid_stagedOrder[instance]
	case *TopEndArcShape:
		return stage.TopEndArcShape_stagedOrder[instance]
	case *TopEndArcShapeGrid:
		return stage.TopEndArcShapeGrid_stagedOrder[instance]
	case *TopGrowthCurve2D:
		return stage.TopGrowthCurve2D_stagedOrder[instance]
	case *TopStackGrowthCurveEndArcShape:
		return stage.TopStackGrowthCurveEndArcShape_stagedOrder[instance]
	case *TopStackGrowthCurveStartArcShape:
		return stage.TopStackGrowthCurveStartArcShape_stagedOrder[instance]
	case *TopStackOfGrowthCurve:
		return stage.TopStackOfGrowthCurve_stagedOrder[instance]
	case *TopStartArcShape:
		return stage.TopStartArcShape_stagedOrder[instance]
	case *TopStartArcShapeGrid:
		return stage.TopStartArcShapeGrid_stagedOrder[instance]
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
	stage.Map_GongStructName_InstancesNb["ArcNormalVectorShape"] = len(stage.ArcNormalVectorShapes)
	stage.Map_GongStructName_InstancesNb["ArcNormalVectorShapeGrid"] = len(stage.ArcNormalVectorShapeGrids)
	stage.Map_GongStructName_InstancesNb["AxesShape"] = len(stage.AxesShapes)
	stage.Map_GongStructName_InstancesNb["BaseVectorShape"] = len(stage.BaseVectorShapes)
	stage.Map_GongStructName_InstancesNb["BaseVectorShapeGrid"] = len(stage.BaseVectorShapeGrids)
	stage.Map_GongStructName_InstancesNb["CircleGridShape"] = len(stage.CircleGridShapes)
	stage.Map_GongStructName_InstancesNb["EndArcShape"] = len(stage.EndArcShapes)
	stage.Map_GongStructName_InstancesNb["EndArcShapeGrid"] = len(stage.EndArcShapeGrids)
	stage.Map_GongStructName_InstancesNb["ExplanationTextShape"] = len(stage.ExplanationTextShapes)
	stage.Map_GongStructName_InstancesNb["GridPathShape"] = len(stage.GridPathShapes)
	stage.Map_GongStructName_InstancesNb["GrowthCurve2D"] = len(stage.GrowthCurve2Ds)
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
	stage.Map_GongStructName_InstancesNb["ShiftedLeftStackGrowthCurveEndArcShape"] = len(stage.ShiftedLeftStackGrowthCurveEndArcShapes)
	stage.Map_GongStructName_InstancesNb["ShiftedLeftStackGrowthCurveStartArcShape"] = len(stage.ShiftedLeftStackGrowthCurveStartArcShapes)
	stage.Map_GongStructName_InstancesNb["ShiftedLeftStackNormalVector"] = len(stage.ShiftedLeftStackNormalVectors)
	stage.Map_GongStructName_InstancesNb["ShiftedLeftStackOfGrowthCurve"] = len(stage.ShiftedLeftStackOfGrowthCurves)
	stage.Map_GongStructName_InstancesNb["ShiftedLeftStackOfNormalVector"] = len(stage.ShiftedLeftStackOfNormalVectors)
	stage.Map_GongStructName_InstancesNb["StackGrowthCurveEndArcShape"] = len(stage.StackGrowthCurveEndArcShapes)
	stage.Map_GongStructName_InstancesNb["StackGrowthCurveStartArcShape"] = len(stage.StackGrowthCurveStartArcShapes)
	stage.Map_GongStructName_InstancesNb["StackOfGrowthCurve"] = len(stage.StackOfGrowthCurves)
	stage.Map_GongStructName_InstancesNb["StartArcShape"] = len(stage.StartArcShapes)
	stage.Map_GongStructName_InstancesNb["StartArcShapeGrid"] = len(stage.StartArcShapeGrids)
	stage.Map_GongStructName_InstancesNb["TopEndArcShape"] = len(stage.TopEndArcShapes)
	stage.Map_GongStructName_InstancesNb["TopEndArcShapeGrid"] = len(stage.TopEndArcShapeGrids)
	stage.Map_GongStructName_InstancesNb["TopGrowthCurve2D"] = len(stage.TopGrowthCurve2Ds)
	stage.Map_GongStructName_InstancesNb["TopStackGrowthCurveEndArcShape"] = len(stage.TopStackGrowthCurveEndArcShapes)
	stage.Map_GongStructName_InstancesNb["TopStackGrowthCurveStartArcShape"] = len(stage.TopStackGrowthCurveStartArcShapes)
	stage.Map_GongStructName_InstancesNb["TopStackOfGrowthCurve"] = len(stage.TopStackOfGrowthCurves)
	stage.Map_GongStructName_InstancesNb["TopStartArcShape"] = len(stage.TopStartArcShapes)
	stage.Map_GongStructName_InstancesNb["TopStartArcShapeGrid"] = len(stage.TopStartArcShapeGrids)
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
// Stage puts arcnormalvectorshape to the model stage
func (arcnormalvectorshape *ArcNormalVectorShape) Stage(stage *Stage) *ArcNormalVectorShape {
	if _, ok := stage.ArcNormalVectorShapes[arcnormalvectorshape]; !ok {
		stage.ArcNormalVectorShapes[arcnormalvectorshape] = struct{}{}
		stage.ArcNormalVectorShape_stagedOrder[arcnormalvectorshape] = stage.ArcNormalVectorShapeOrder
		stage.ArcNormalVectorShape_orderStaged[stage.ArcNormalVectorShapeOrder] = arcnormalvectorshape
		stage.ArcNormalVectorShapeOrder++
	}
	stage.ArcNormalVectorShapes_mapString[arcnormalvectorshape.Name] = arcnormalvectorshape

	return arcnormalvectorshape
}

// StagePreserveOrder puts arcnormalvectorshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ArcNormalVectorShapeOrder
// - update stage.ArcNormalVectorShapeOrder accordingly
func (arcnormalvectorshape *ArcNormalVectorShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ArcNormalVectorShapes[arcnormalvectorshape]; !ok {
		stage.ArcNormalVectorShapes[arcnormalvectorshape] = struct{}{}

		if order > stage.ArcNormalVectorShapeOrder {
			stage.ArcNormalVectorShapeOrder = order
		}
		stage.ArcNormalVectorShape_stagedOrder[arcnormalvectorshape] = order
		stage.ArcNormalVectorShape_orderStaged[order] = arcnormalvectorshape
		stage.ArcNormalVectorShapeOrder++
	}
	stage.ArcNormalVectorShapes_mapString[arcnormalvectorshape.Name] = arcnormalvectorshape
}

// Unstage removes arcnormalvectorshape off the model stage
func (arcnormalvectorshape *ArcNormalVectorShape) Unstage(stage *Stage) *ArcNormalVectorShape {
	delete(stage.ArcNormalVectorShapes, arcnormalvectorshape)
	// issue1150
	// delete(stage.ArcNormalVectorShape_stagedOrder, arcnormalvectorshape)
	delete(stage.ArcNormalVectorShapes_mapString, arcnormalvectorshape.Name)

	return arcnormalvectorshape
}

// UnstageVoid removes arcnormalvectorshape off the model stage
func (arcnormalvectorshape *ArcNormalVectorShape) UnstageVoid(stage *Stage) {
	delete(stage.ArcNormalVectorShapes, arcnormalvectorshape)
	// issue1150
	// delete(stage.ArcNormalVectorShape_stagedOrder, arcnormalvectorshape)
	delete(stage.ArcNormalVectorShapes_mapString, arcnormalvectorshape.Name)
}

// commit arcnormalvectorshape to the back repo (if it is already staged)
func (arcnormalvectorshape *ArcNormalVectorShape) Commit(stage *Stage) *ArcNormalVectorShape {
	if _, ok := stage.ArcNormalVectorShapes[arcnormalvectorshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitArcNormalVectorShape(arcnormalvectorshape)
		}
	}
	return arcnormalvectorshape
}

func (arcnormalvectorshape *ArcNormalVectorShape) CommitVoid(stage *Stage) {
	arcnormalvectorshape.Commit(stage)
}

func (arcnormalvectorshape *ArcNormalVectorShape) StageVoid(stage *Stage) {
	arcnormalvectorshape.Stage(stage)
}

// Checkout arcnormalvectorshape to the back repo (if it is already staged)
func (arcnormalvectorshape *ArcNormalVectorShape) Checkout(stage *Stage) *ArcNormalVectorShape {
	if _, ok := stage.ArcNormalVectorShapes[arcnormalvectorshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutArcNormalVectorShape(arcnormalvectorshape)
		}
	}
	return arcnormalvectorshape
}

// for satisfaction of GongStruct interface
func (arcnormalvectorshape *ArcNormalVectorShape) GetName() (res string) {
	return arcnormalvectorshape.Name
}

// for satisfaction of GongStruct interface
func (arcnormalvectorshape *ArcNormalVectorShape) SetName(name string) {
	arcnormalvectorshape.Name = name
}

// Stage puts arcnormalvectorshapegrid to the model stage
func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) Stage(stage *Stage) *ArcNormalVectorShapeGrid {
	if _, ok := stage.ArcNormalVectorShapeGrids[arcnormalvectorshapegrid]; !ok {
		stage.ArcNormalVectorShapeGrids[arcnormalvectorshapegrid] = struct{}{}
		stage.ArcNormalVectorShapeGrid_stagedOrder[arcnormalvectorshapegrid] = stage.ArcNormalVectorShapeGridOrder
		stage.ArcNormalVectorShapeGrid_orderStaged[stage.ArcNormalVectorShapeGridOrder] = arcnormalvectorshapegrid
		stage.ArcNormalVectorShapeGridOrder++
	}
	stage.ArcNormalVectorShapeGrids_mapString[arcnormalvectorshapegrid.Name] = arcnormalvectorshapegrid

	return arcnormalvectorshapegrid
}

// StagePreserveOrder puts arcnormalvectorshapegrid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ArcNormalVectorShapeGridOrder
// - update stage.ArcNormalVectorShapeGridOrder accordingly
func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ArcNormalVectorShapeGrids[arcnormalvectorshapegrid]; !ok {
		stage.ArcNormalVectorShapeGrids[arcnormalvectorshapegrid] = struct{}{}

		if order > stage.ArcNormalVectorShapeGridOrder {
			stage.ArcNormalVectorShapeGridOrder = order
		}
		stage.ArcNormalVectorShapeGrid_stagedOrder[arcnormalvectorshapegrid] = order
		stage.ArcNormalVectorShapeGrid_orderStaged[order] = arcnormalvectorshapegrid
		stage.ArcNormalVectorShapeGridOrder++
	}
	stage.ArcNormalVectorShapeGrids_mapString[arcnormalvectorshapegrid.Name] = arcnormalvectorshapegrid
}

// Unstage removes arcnormalvectorshapegrid off the model stage
func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) Unstage(stage *Stage) *ArcNormalVectorShapeGrid {
	delete(stage.ArcNormalVectorShapeGrids, arcnormalvectorshapegrid)
	// issue1150
	// delete(stage.ArcNormalVectorShapeGrid_stagedOrder, arcnormalvectorshapegrid)
	delete(stage.ArcNormalVectorShapeGrids_mapString, arcnormalvectorshapegrid.Name)

	return arcnormalvectorshapegrid
}

// UnstageVoid removes arcnormalvectorshapegrid off the model stage
func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) UnstageVoid(stage *Stage) {
	delete(stage.ArcNormalVectorShapeGrids, arcnormalvectorshapegrid)
	// issue1150
	// delete(stage.ArcNormalVectorShapeGrid_stagedOrder, arcnormalvectorshapegrid)
	delete(stage.ArcNormalVectorShapeGrids_mapString, arcnormalvectorshapegrid.Name)
}

// commit arcnormalvectorshapegrid to the back repo (if it is already staged)
func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) Commit(stage *Stage) *ArcNormalVectorShapeGrid {
	if _, ok := stage.ArcNormalVectorShapeGrids[arcnormalvectorshapegrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitArcNormalVectorShapeGrid(arcnormalvectorshapegrid)
		}
	}
	return arcnormalvectorshapegrid
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) CommitVoid(stage *Stage) {
	arcnormalvectorshapegrid.Commit(stage)
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) StageVoid(stage *Stage) {
	arcnormalvectorshapegrid.Stage(stage)
}

// Checkout arcnormalvectorshapegrid to the back repo (if it is already staged)
func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) Checkout(stage *Stage) *ArcNormalVectorShapeGrid {
	if _, ok := stage.ArcNormalVectorShapeGrids[arcnormalvectorshapegrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutArcNormalVectorShapeGrid(arcnormalvectorshapegrid)
		}
	}
	return arcnormalvectorshapegrid
}

// for satisfaction of GongStruct interface
func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GetName() (res string) {
	return arcnormalvectorshapegrid.Name
}

// for satisfaction of GongStruct interface
func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) SetName(name string) {
	arcnormalvectorshapegrid.Name = name
}

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

// Stage puts basevectorshape to the model stage
func (basevectorshape *BaseVectorShape) Stage(stage *Stage) *BaseVectorShape {
	if _, ok := stage.BaseVectorShapes[basevectorshape]; !ok {
		stage.BaseVectorShapes[basevectorshape] = struct{}{}
		stage.BaseVectorShape_stagedOrder[basevectorshape] = stage.BaseVectorShapeOrder
		stage.BaseVectorShape_orderStaged[stage.BaseVectorShapeOrder] = basevectorshape
		stage.BaseVectorShapeOrder++
	}
	stage.BaseVectorShapes_mapString[basevectorshape.Name] = basevectorshape

	return basevectorshape
}

// StagePreserveOrder puts basevectorshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BaseVectorShapeOrder
// - update stage.BaseVectorShapeOrder accordingly
func (basevectorshape *BaseVectorShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.BaseVectorShapes[basevectorshape]; !ok {
		stage.BaseVectorShapes[basevectorshape] = struct{}{}

		if order > stage.BaseVectorShapeOrder {
			stage.BaseVectorShapeOrder = order
		}
		stage.BaseVectorShape_stagedOrder[basevectorshape] = order
		stage.BaseVectorShape_orderStaged[order] = basevectorshape
		stage.BaseVectorShapeOrder++
	}
	stage.BaseVectorShapes_mapString[basevectorshape.Name] = basevectorshape
}

// Unstage removes basevectorshape off the model stage
func (basevectorshape *BaseVectorShape) Unstage(stage *Stage) *BaseVectorShape {
	delete(stage.BaseVectorShapes, basevectorshape)
	// issue1150
	// delete(stage.BaseVectorShape_stagedOrder, basevectorshape)
	delete(stage.BaseVectorShapes_mapString, basevectorshape.Name)

	return basevectorshape
}

// UnstageVoid removes basevectorshape off the model stage
func (basevectorshape *BaseVectorShape) UnstageVoid(stage *Stage) {
	delete(stage.BaseVectorShapes, basevectorshape)
	// issue1150
	// delete(stage.BaseVectorShape_stagedOrder, basevectorshape)
	delete(stage.BaseVectorShapes_mapString, basevectorshape.Name)
}

// commit basevectorshape to the back repo (if it is already staged)
func (basevectorshape *BaseVectorShape) Commit(stage *Stage) *BaseVectorShape {
	if _, ok := stage.BaseVectorShapes[basevectorshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBaseVectorShape(basevectorshape)
		}
	}
	return basevectorshape
}

func (basevectorshape *BaseVectorShape) CommitVoid(stage *Stage) {
	basevectorshape.Commit(stage)
}

func (basevectorshape *BaseVectorShape) StageVoid(stage *Stage) {
	basevectorshape.Stage(stage)
}

// Checkout basevectorshape to the back repo (if it is already staged)
func (basevectorshape *BaseVectorShape) Checkout(stage *Stage) *BaseVectorShape {
	if _, ok := stage.BaseVectorShapes[basevectorshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBaseVectorShape(basevectorshape)
		}
	}
	return basevectorshape
}

// for satisfaction of GongStruct interface
func (basevectorshape *BaseVectorShape) GetName() (res string) {
	return basevectorshape.Name
}

// for satisfaction of GongStruct interface
func (basevectorshape *BaseVectorShape) SetName(name string) {
	basevectorshape.Name = name
}

// Stage puts basevectorshapegrid to the model stage
func (basevectorshapegrid *BaseVectorShapeGrid) Stage(stage *Stage) *BaseVectorShapeGrid {
	if _, ok := stage.BaseVectorShapeGrids[basevectorshapegrid]; !ok {
		stage.BaseVectorShapeGrids[basevectorshapegrid] = struct{}{}
		stage.BaseVectorShapeGrid_stagedOrder[basevectorshapegrid] = stage.BaseVectorShapeGridOrder
		stage.BaseVectorShapeGrid_orderStaged[stage.BaseVectorShapeGridOrder] = basevectorshapegrid
		stage.BaseVectorShapeGridOrder++
	}
	stage.BaseVectorShapeGrids_mapString[basevectorshapegrid.Name] = basevectorshapegrid

	return basevectorshapegrid
}

// StagePreserveOrder puts basevectorshapegrid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BaseVectorShapeGridOrder
// - update stage.BaseVectorShapeGridOrder accordingly
func (basevectorshapegrid *BaseVectorShapeGrid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.BaseVectorShapeGrids[basevectorshapegrid]; !ok {
		stage.BaseVectorShapeGrids[basevectorshapegrid] = struct{}{}

		if order > stage.BaseVectorShapeGridOrder {
			stage.BaseVectorShapeGridOrder = order
		}
		stage.BaseVectorShapeGrid_stagedOrder[basevectorshapegrid] = order
		stage.BaseVectorShapeGrid_orderStaged[order] = basevectorshapegrid
		stage.BaseVectorShapeGridOrder++
	}
	stage.BaseVectorShapeGrids_mapString[basevectorshapegrid.Name] = basevectorshapegrid
}

// Unstage removes basevectorshapegrid off the model stage
func (basevectorshapegrid *BaseVectorShapeGrid) Unstage(stage *Stage) *BaseVectorShapeGrid {
	delete(stage.BaseVectorShapeGrids, basevectorshapegrid)
	// issue1150
	// delete(stage.BaseVectorShapeGrid_stagedOrder, basevectorshapegrid)
	delete(stage.BaseVectorShapeGrids_mapString, basevectorshapegrid.Name)

	return basevectorshapegrid
}

// UnstageVoid removes basevectorshapegrid off the model stage
func (basevectorshapegrid *BaseVectorShapeGrid) UnstageVoid(stage *Stage) {
	delete(stage.BaseVectorShapeGrids, basevectorshapegrid)
	// issue1150
	// delete(stage.BaseVectorShapeGrid_stagedOrder, basevectorshapegrid)
	delete(stage.BaseVectorShapeGrids_mapString, basevectorshapegrid.Name)
}

// commit basevectorshapegrid to the back repo (if it is already staged)
func (basevectorshapegrid *BaseVectorShapeGrid) Commit(stage *Stage) *BaseVectorShapeGrid {
	if _, ok := stage.BaseVectorShapeGrids[basevectorshapegrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBaseVectorShapeGrid(basevectorshapegrid)
		}
	}
	return basevectorshapegrid
}

func (basevectorshapegrid *BaseVectorShapeGrid) CommitVoid(stage *Stage) {
	basevectorshapegrid.Commit(stage)
}

func (basevectorshapegrid *BaseVectorShapeGrid) StageVoid(stage *Stage) {
	basevectorshapegrid.Stage(stage)
}

// Checkout basevectorshapegrid to the back repo (if it is already staged)
func (basevectorshapegrid *BaseVectorShapeGrid) Checkout(stage *Stage) *BaseVectorShapeGrid {
	if _, ok := stage.BaseVectorShapeGrids[basevectorshapegrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBaseVectorShapeGrid(basevectorshapegrid)
		}
	}
	return basevectorshapegrid
}

// for satisfaction of GongStruct interface
func (basevectorshapegrid *BaseVectorShapeGrid) GetName() (res string) {
	return basevectorshapegrid.Name
}

// for satisfaction of GongStruct interface
func (basevectorshapegrid *BaseVectorShapeGrid) SetName(name string) {
	basevectorshapegrid.Name = name
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

// Stage puts growthcurve2d to the model stage
func (growthcurve2d *GrowthCurve2D) Stage(stage *Stage) *GrowthCurve2D {
	if _, ok := stage.GrowthCurve2Ds[growthcurve2d]; !ok {
		stage.GrowthCurve2Ds[growthcurve2d] = struct{}{}
		stage.GrowthCurve2D_stagedOrder[growthcurve2d] = stage.GrowthCurve2DOrder
		stage.GrowthCurve2D_orderStaged[stage.GrowthCurve2DOrder] = growthcurve2d
		stage.GrowthCurve2DOrder++
	}
	stage.GrowthCurve2Ds_mapString[growthcurve2d.Name] = growthcurve2d

	return growthcurve2d
}

// StagePreserveOrder puts growthcurve2d to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GrowthCurve2DOrder
// - update stage.GrowthCurve2DOrder accordingly
func (growthcurve2d *GrowthCurve2D) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.GrowthCurve2Ds[growthcurve2d]; !ok {
		stage.GrowthCurve2Ds[growthcurve2d] = struct{}{}

		if order > stage.GrowthCurve2DOrder {
			stage.GrowthCurve2DOrder = order
		}
		stage.GrowthCurve2D_stagedOrder[growthcurve2d] = order
		stage.GrowthCurve2D_orderStaged[order] = growthcurve2d
		stage.GrowthCurve2DOrder++
	}
	stage.GrowthCurve2Ds_mapString[growthcurve2d.Name] = growthcurve2d
}

// Unstage removes growthcurve2d off the model stage
func (growthcurve2d *GrowthCurve2D) Unstage(stage *Stage) *GrowthCurve2D {
	delete(stage.GrowthCurve2Ds, growthcurve2d)
	// issue1150
	// delete(stage.GrowthCurve2D_stagedOrder, growthcurve2d)
	delete(stage.GrowthCurve2Ds_mapString, growthcurve2d.Name)

	return growthcurve2d
}

// UnstageVoid removes growthcurve2d off the model stage
func (growthcurve2d *GrowthCurve2D) UnstageVoid(stage *Stage) {
	delete(stage.GrowthCurve2Ds, growthcurve2d)
	// issue1150
	// delete(stage.GrowthCurve2D_stagedOrder, growthcurve2d)
	delete(stage.GrowthCurve2Ds_mapString, growthcurve2d.Name)
}

// commit growthcurve2d to the back repo (if it is already staged)
func (growthcurve2d *GrowthCurve2D) Commit(stage *Stage) *GrowthCurve2D {
	if _, ok := stage.GrowthCurve2Ds[growthcurve2d]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGrowthCurve2D(growthcurve2d)
		}
	}
	return growthcurve2d
}

func (growthcurve2d *GrowthCurve2D) CommitVoid(stage *Stage) {
	growthcurve2d.Commit(stage)
}

func (growthcurve2d *GrowthCurve2D) StageVoid(stage *Stage) {
	growthcurve2d.Stage(stage)
}

// Checkout growthcurve2d to the back repo (if it is already staged)
func (growthcurve2d *GrowthCurve2D) Checkout(stage *Stage) *GrowthCurve2D {
	if _, ok := stage.GrowthCurve2Ds[growthcurve2d]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGrowthCurve2D(growthcurve2d)
		}
	}
	return growthcurve2d
}

// for satisfaction of GongStruct interface
func (growthcurve2d *GrowthCurve2D) GetName() (res string) {
	return growthcurve2d.Name
}

// for satisfaction of GongStruct interface
func (growthcurve2d *GrowthCurve2D) SetName(name string) {
	growthcurve2d.Name = name
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

// Stage puts shiftedleftstackgrowthcurveendarcshape to the model stage
func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) Stage(stage *Stage) *ShiftedLeftStackGrowthCurveEndArcShape {
	if _, ok := stage.ShiftedLeftStackGrowthCurveEndArcShapes[shiftedleftstackgrowthcurveendarcshape]; !ok {
		stage.ShiftedLeftStackGrowthCurveEndArcShapes[shiftedleftstackgrowthcurveendarcshape] = struct{}{}
		stage.ShiftedLeftStackGrowthCurveEndArcShape_stagedOrder[shiftedleftstackgrowthcurveendarcshape] = stage.ShiftedLeftStackGrowthCurveEndArcShapeOrder
		stage.ShiftedLeftStackGrowthCurveEndArcShape_orderStaged[stage.ShiftedLeftStackGrowthCurveEndArcShapeOrder] = shiftedleftstackgrowthcurveendarcshape
		stage.ShiftedLeftStackGrowthCurveEndArcShapeOrder++
	}
	stage.ShiftedLeftStackGrowthCurveEndArcShapes_mapString[shiftedleftstackgrowthcurveendarcshape.Name] = shiftedleftstackgrowthcurveendarcshape

	return shiftedleftstackgrowthcurveendarcshape
}

// StagePreserveOrder puts shiftedleftstackgrowthcurveendarcshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ShiftedLeftStackGrowthCurveEndArcShapeOrder
// - update stage.ShiftedLeftStackGrowthCurveEndArcShapeOrder accordingly
func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ShiftedLeftStackGrowthCurveEndArcShapes[shiftedleftstackgrowthcurveendarcshape]; !ok {
		stage.ShiftedLeftStackGrowthCurveEndArcShapes[shiftedleftstackgrowthcurveendarcshape] = struct{}{}

		if order > stage.ShiftedLeftStackGrowthCurveEndArcShapeOrder {
			stage.ShiftedLeftStackGrowthCurveEndArcShapeOrder = order
		}
		stage.ShiftedLeftStackGrowthCurveEndArcShape_stagedOrder[shiftedleftstackgrowthcurveendarcshape] = order
		stage.ShiftedLeftStackGrowthCurveEndArcShape_orderStaged[order] = shiftedleftstackgrowthcurveendarcshape
		stage.ShiftedLeftStackGrowthCurveEndArcShapeOrder++
	}
	stage.ShiftedLeftStackGrowthCurveEndArcShapes_mapString[shiftedleftstackgrowthcurveendarcshape.Name] = shiftedleftstackgrowthcurveendarcshape
}

// Unstage removes shiftedleftstackgrowthcurveendarcshape off the model stage
func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) Unstage(stage *Stage) *ShiftedLeftStackGrowthCurveEndArcShape {
	delete(stage.ShiftedLeftStackGrowthCurveEndArcShapes, shiftedleftstackgrowthcurveendarcshape)
	// issue1150
	// delete(stage.ShiftedLeftStackGrowthCurveEndArcShape_stagedOrder, shiftedleftstackgrowthcurveendarcshape)
	delete(stage.ShiftedLeftStackGrowthCurveEndArcShapes_mapString, shiftedleftstackgrowthcurveendarcshape.Name)

	return shiftedleftstackgrowthcurveendarcshape
}

// UnstageVoid removes shiftedleftstackgrowthcurveendarcshape off the model stage
func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) UnstageVoid(stage *Stage) {
	delete(stage.ShiftedLeftStackGrowthCurveEndArcShapes, shiftedleftstackgrowthcurveendarcshape)
	// issue1150
	// delete(stage.ShiftedLeftStackGrowthCurveEndArcShape_stagedOrder, shiftedleftstackgrowthcurveendarcshape)
	delete(stage.ShiftedLeftStackGrowthCurveEndArcShapes_mapString, shiftedleftstackgrowthcurveendarcshape.Name)
}

// commit shiftedleftstackgrowthcurveendarcshape to the back repo (if it is already staged)
func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) Commit(stage *Stage) *ShiftedLeftStackGrowthCurveEndArcShape {
	if _, ok := stage.ShiftedLeftStackGrowthCurveEndArcShapes[shiftedleftstackgrowthcurveendarcshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitShiftedLeftStackGrowthCurveEndArcShape(shiftedleftstackgrowthcurveendarcshape)
		}
	}
	return shiftedleftstackgrowthcurveendarcshape
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) CommitVoid(stage *Stage) {
	shiftedleftstackgrowthcurveendarcshape.Commit(stage)
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) StageVoid(stage *Stage) {
	shiftedleftstackgrowthcurveendarcshape.Stage(stage)
}

// Checkout shiftedleftstackgrowthcurveendarcshape to the back repo (if it is already staged)
func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) Checkout(stage *Stage) *ShiftedLeftStackGrowthCurveEndArcShape {
	if _, ok := stage.ShiftedLeftStackGrowthCurveEndArcShapes[shiftedleftstackgrowthcurveendarcshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutShiftedLeftStackGrowthCurveEndArcShape(shiftedleftstackgrowthcurveendarcshape)
		}
	}
	return shiftedleftstackgrowthcurveendarcshape
}

// for satisfaction of GongStruct interface
func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GetName() (res string) {
	return shiftedleftstackgrowthcurveendarcshape.Name
}

// for satisfaction of GongStruct interface
func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) SetName(name string) {
	shiftedleftstackgrowthcurveendarcshape.Name = name
}

// Stage puts shiftedleftstackgrowthcurvestartarcshape to the model stage
func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) Stage(stage *Stage) *ShiftedLeftStackGrowthCurveStartArcShape {
	if _, ok := stage.ShiftedLeftStackGrowthCurveStartArcShapes[shiftedleftstackgrowthcurvestartarcshape]; !ok {
		stage.ShiftedLeftStackGrowthCurveStartArcShapes[shiftedleftstackgrowthcurvestartarcshape] = struct{}{}
		stage.ShiftedLeftStackGrowthCurveStartArcShape_stagedOrder[shiftedleftstackgrowthcurvestartarcshape] = stage.ShiftedLeftStackGrowthCurveStartArcShapeOrder
		stage.ShiftedLeftStackGrowthCurveStartArcShape_orderStaged[stage.ShiftedLeftStackGrowthCurveStartArcShapeOrder] = shiftedleftstackgrowthcurvestartarcshape
		stage.ShiftedLeftStackGrowthCurveStartArcShapeOrder++
	}
	stage.ShiftedLeftStackGrowthCurveStartArcShapes_mapString[shiftedleftstackgrowthcurvestartarcshape.Name] = shiftedleftstackgrowthcurvestartarcshape

	return shiftedleftstackgrowthcurvestartarcshape
}

// StagePreserveOrder puts shiftedleftstackgrowthcurvestartarcshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ShiftedLeftStackGrowthCurveStartArcShapeOrder
// - update stage.ShiftedLeftStackGrowthCurveStartArcShapeOrder accordingly
func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ShiftedLeftStackGrowthCurveStartArcShapes[shiftedleftstackgrowthcurvestartarcshape]; !ok {
		stage.ShiftedLeftStackGrowthCurveStartArcShapes[shiftedleftstackgrowthcurvestartarcshape] = struct{}{}

		if order > stage.ShiftedLeftStackGrowthCurveStartArcShapeOrder {
			stage.ShiftedLeftStackGrowthCurveStartArcShapeOrder = order
		}
		stage.ShiftedLeftStackGrowthCurveStartArcShape_stagedOrder[shiftedleftstackgrowthcurvestartarcshape] = order
		stage.ShiftedLeftStackGrowthCurveStartArcShape_orderStaged[order] = shiftedleftstackgrowthcurvestartarcshape
		stage.ShiftedLeftStackGrowthCurveStartArcShapeOrder++
	}
	stage.ShiftedLeftStackGrowthCurveStartArcShapes_mapString[shiftedleftstackgrowthcurvestartarcshape.Name] = shiftedleftstackgrowthcurvestartarcshape
}

// Unstage removes shiftedleftstackgrowthcurvestartarcshape off the model stage
func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) Unstage(stage *Stage) *ShiftedLeftStackGrowthCurveStartArcShape {
	delete(stage.ShiftedLeftStackGrowthCurveStartArcShapes, shiftedleftstackgrowthcurvestartarcshape)
	// issue1150
	// delete(stage.ShiftedLeftStackGrowthCurveStartArcShape_stagedOrder, shiftedleftstackgrowthcurvestartarcshape)
	delete(stage.ShiftedLeftStackGrowthCurveStartArcShapes_mapString, shiftedleftstackgrowthcurvestartarcshape.Name)

	return shiftedleftstackgrowthcurvestartarcshape
}

// UnstageVoid removes shiftedleftstackgrowthcurvestartarcshape off the model stage
func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) UnstageVoid(stage *Stage) {
	delete(stage.ShiftedLeftStackGrowthCurveStartArcShapes, shiftedleftstackgrowthcurvestartarcshape)
	// issue1150
	// delete(stage.ShiftedLeftStackGrowthCurveStartArcShape_stagedOrder, shiftedleftstackgrowthcurvestartarcshape)
	delete(stage.ShiftedLeftStackGrowthCurveStartArcShapes_mapString, shiftedleftstackgrowthcurvestartarcshape.Name)
}

// commit shiftedleftstackgrowthcurvestartarcshape to the back repo (if it is already staged)
func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) Commit(stage *Stage) *ShiftedLeftStackGrowthCurveStartArcShape {
	if _, ok := stage.ShiftedLeftStackGrowthCurveStartArcShapes[shiftedleftstackgrowthcurvestartarcshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitShiftedLeftStackGrowthCurveStartArcShape(shiftedleftstackgrowthcurvestartarcshape)
		}
	}
	return shiftedleftstackgrowthcurvestartarcshape
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) CommitVoid(stage *Stage) {
	shiftedleftstackgrowthcurvestartarcshape.Commit(stage)
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) StageVoid(stage *Stage) {
	shiftedleftstackgrowthcurvestartarcshape.Stage(stage)
}

// Checkout shiftedleftstackgrowthcurvestartarcshape to the back repo (if it is already staged)
func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) Checkout(stage *Stage) *ShiftedLeftStackGrowthCurveStartArcShape {
	if _, ok := stage.ShiftedLeftStackGrowthCurveStartArcShapes[shiftedleftstackgrowthcurvestartarcshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutShiftedLeftStackGrowthCurveStartArcShape(shiftedleftstackgrowthcurvestartarcshape)
		}
	}
	return shiftedleftstackgrowthcurvestartarcshape
}

// for satisfaction of GongStruct interface
func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GetName() (res string) {
	return shiftedleftstackgrowthcurvestartarcshape.Name
}

// for satisfaction of GongStruct interface
func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) SetName(name string) {
	shiftedleftstackgrowthcurvestartarcshape.Name = name
}

// Stage puts shiftedleftstacknormalvector to the model stage
func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) Stage(stage *Stage) *ShiftedLeftStackNormalVector {
	if _, ok := stage.ShiftedLeftStackNormalVectors[shiftedleftstacknormalvector]; !ok {
		stage.ShiftedLeftStackNormalVectors[shiftedleftstacknormalvector] = struct{}{}
		stage.ShiftedLeftStackNormalVector_stagedOrder[shiftedleftstacknormalvector] = stage.ShiftedLeftStackNormalVectorOrder
		stage.ShiftedLeftStackNormalVector_orderStaged[stage.ShiftedLeftStackNormalVectorOrder] = shiftedleftstacknormalvector
		stage.ShiftedLeftStackNormalVectorOrder++
	}
	stage.ShiftedLeftStackNormalVectors_mapString[shiftedleftstacknormalvector.Name] = shiftedleftstacknormalvector

	return shiftedleftstacknormalvector
}

// StagePreserveOrder puts shiftedleftstacknormalvector to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ShiftedLeftStackNormalVectorOrder
// - update stage.ShiftedLeftStackNormalVectorOrder accordingly
func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ShiftedLeftStackNormalVectors[shiftedleftstacknormalvector]; !ok {
		stage.ShiftedLeftStackNormalVectors[shiftedleftstacknormalvector] = struct{}{}

		if order > stage.ShiftedLeftStackNormalVectorOrder {
			stage.ShiftedLeftStackNormalVectorOrder = order
		}
		stage.ShiftedLeftStackNormalVector_stagedOrder[shiftedleftstacknormalvector] = order
		stage.ShiftedLeftStackNormalVector_orderStaged[order] = shiftedleftstacknormalvector
		stage.ShiftedLeftStackNormalVectorOrder++
	}
	stage.ShiftedLeftStackNormalVectors_mapString[shiftedleftstacknormalvector.Name] = shiftedleftstacknormalvector
}

// Unstage removes shiftedleftstacknormalvector off the model stage
func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) Unstage(stage *Stage) *ShiftedLeftStackNormalVector {
	delete(stage.ShiftedLeftStackNormalVectors, shiftedleftstacknormalvector)
	// issue1150
	// delete(stage.ShiftedLeftStackNormalVector_stagedOrder, shiftedleftstacknormalvector)
	delete(stage.ShiftedLeftStackNormalVectors_mapString, shiftedleftstacknormalvector.Name)

	return shiftedleftstacknormalvector
}

// UnstageVoid removes shiftedleftstacknormalvector off the model stage
func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) UnstageVoid(stage *Stage) {
	delete(stage.ShiftedLeftStackNormalVectors, shiftedleftstacknormalvector)
	// issue1150
	// delete(stage.ShiftedLeftStackNormalVector_stagedOrder, shiftedleftstacknormalvector)
	delete(stage.ShiftedLeftStackNormalVectors_mapString, shiftedleftstacknormalvector.Name)
}

// commit shiftedleftstacknormalvector to the back repo (if it is already staged)
func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) Commit(stage *Stage) *ShiftedLeftStackNormalVector {
	if _, ok := stage.ShiftedLeftStackNormalVectors[shiftedleftstacknormalvector]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitShiftedLeftStackNormalVector(shiftedleftstacknormalvector)
		}
	}
	return shiftedleftstacknormalvector
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) CommitVoid(stage *Stage) {
	shiftedleftstacknormalvector.Commit(stage)
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) StageVoid(stage *Stage) {
	shiftedleftstacknormalvector.Stage(stage)
}

// Checkout shiftedleftstacknormalvector to the back repo (if it is already staged)
func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) Checkout(stage *Stage) *ShiftedLeftStackNormalVector {
	if _, ok := stage.ShiftedLeftStackNormalVectors[shiftedleftstacknormalvector]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutShiftedLeftStackNormalVector(shiftedleftstacknormalvector)
		}
	}
	return shiftedleftstacknormalvector
}

// for satisfaction of GongStruct interface
func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GetName() (res string) {
	return shiftedleftstacknormalvector.Name
}

// for satisfaction of GongStruct interface
func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) SetName(name string) {
	shiftedleftstacknormalvector.Name = name
}

// Stage puts shiftedleftstackofgrowthcurve to the model stage
func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) Stage(stage *Stage) *ShiftedLeftStackOfGrowthCurve {
	if _, ok := stage.ShiftedLeftStackOfGrowthCurves[shiftedleftstackofgrowthcurve]; !ok {
		stage.ShiftedLeftStackOfGrowthCurves[shiftedleftstackofgrowthcurve] = struct{}{}
		stage.ShiftedLeftStackOfGrowthCurve_stagedOrder[shiftedleftstackofgrowthcurve] = stage.ShiftedLeftStackOfGrowthCurveOrder
		stage.ShiftedLeftStackOfGrowthCurve_orderStaged[stage.ShiftedLeftStackOfGrowthCurveOrder] = shiftedleftstackofgrowthcurve
		stage.ShiftedLeftStackOfGrowthCurveOrder++
	}
	stage.ShiftedLeftStackOfGrowthCurves_mapString[shiftedleftstackofgrowthcurve.Name] = shiftedleftstackofgrowthcurve

	return shiftedleftstackofgrowthcurve
}

// StagePreserveOrder puts shiftedleftstackofgrowthcurve to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ShiftedLeftStackOfGrowthCurveOrder
// - update stage.ShiftedLeftStackOfGrowthCurveOrder accordingly
func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ShiftedLeftStackOfGrowthCurves[shiftedleftstackofgrowthcurve]; !ok {
		stage.ShiftedLeftStackOfGrowthCurves[shiftedleftstackofgrowthcurve] = struct{}{}

		if order > stage.ShiftedLeftStackOfGrowthCurveOrder {
			stage.ShiftedLeftStackOfGrowthCurveOrder = order
		}
		stage.ShiftedLeftStackOfGrowthCurve_stagedOrder[shiftedleftstackofgrowthcurve] = order
		stage.ShiftedLeftStackOfGrowthCurve_orderStaged[order] = shiftedleftstackofgrowthcurve
		stage.ShiftedLeftStackOfGrowthCurveOrder++
	}
	stage.ShiftedLeftStackOfGrowthCurves_mapString[shiftedleftstackofgrowthcurve.Name] = shiftedleftstackofgrowthcurve
}

// Unstage removes shiftedleftstackofgrowthcurve off the model stage
func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) Unstage(stage *Stage) *ShiftedLeftStackOfGrowthCurve {
	delete(stage.ShiftedLeftStackOfGrowthCurves, shiftedleftstackofgrowthcurve)
	// issue1150
	// delete(stage.ShiftedLeftStackOfGrowthCurve_stagedOrder, shiftedleftstackofgrowthcurve)
	delete(stage.ShiftedLeftStackOfGrowthCurves_mapString, shiftedleftstackofgrowthcurve.Name)

	return shiftedleftstackofgrowthcurve
}

// UnstageVoid removes shiftedleftstackofgrowthcurve off the model stage
func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) UnstageVoid(stage *Stage) {
	delete(stage.ShiftedLeftStackOfGrowthCurves, shiftedleftstackofgrowthcurve)
	// issue1150
	// delete(stage.ShiftedLeftStackOfGrowthCurve_stagedOrder, shiftedleftstackofgrowthcurve)
	delete(stage.ShiftedLeftStackOfGrowthCurves_mapString, shiftedleftstackofgrowthcurve.Name)
}

// commit shiftedleftstackofgrowthcurve to the back repo (if it is already staged)
func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) Commit(stage *Stage) *ShiftedLeftStackOfGrowthCurve {
	if _, ok := stage.ShiftedLeftStackOfGrowthCurves[shiftedleftstackofgrowthcurve]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitShiftedLeftStackOfGrowthCurve(shiftedleftstackofgrowthcurve)
		}
	}
	return shiftedleftstackofgrowthcurve
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) CommitVoid(stage *Stage) {
	shiftedleftstackofgrowthcurve.Commit(stage)
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) StageVoid(stage *Stage) {
	shiftedleftstackofgrowthcurve.Stage(stage)
}

// Checkout shiftedleftstackofgrowthcurve to the back repo (if it is already staged)
func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) Checkout(stage *Stage) *ShiftedLeftStackOfGrowthCurve {
	if _, ok := stage.ShiftedLeftStackOfGrowthCurves[shiftedleftstackofgrowthcurve]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutShiftedLeftStackOfGrowthCurve(shiftedleftstackofgrowthcurve)
		}
	}
	return shiftedleftstackofgrowthcurve
}

// for satisfaction of GongStruct interface
func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GetName() (res string) {
	return shiftedleftstackofgrowthcurve.Name
}

// for satisfaction of GongStruct interface
func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) SetName(name string) {
	shiftedleftstackofgrowthcurve.Name = name
}

// Stage puts shiftedleftstackofnormalvector to the model stage
func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) Stage(stage *Stage) *ShiftedLeftStackOfNormalVector {
	if _, ok := stage.ShiftedLeftStackOfNormalVectors[shiftedleftstackofnormalvector]; !ok {
		stage.ShiftedLeftStackOfNormalVectors[shiftedleftstackofnormalvector] = struct{}{}
		stage.ShiftedLeftStackOfNormalVector_stagedOrder[shiftedleftstackofnormalvector] = stage.ShiftedLeftStackOfNormalVectorOrder
		stage.ShiftedLeftStackOfNormalVector_orderStaged[stage.ShiftedLeftStackOfNormalVectorOrder] = shiftedleftstackofnormalvector
		stage.ShiftedLeftStackOfNormalVectorOrder++
	}
	stage.ShiftedLeftStackOfNormalVectors_mapString[shiftedleftstackofnormalvector.Name] = shiftedleftstackofnormalvector

	return shiftedleftstackofnormalvector
}

// StagePreserveOrder puts shiftedleftstackofnormalvector to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ShiftedLeftStackOfNormalVectorOrder
// - update stage.ShiftedLeftStackOfNormalVectorOrder accordingly
func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ShiftedLeftStackOfNormalVectors[shiftedleftstackofnormalvector]; !ok {
		stage.ShiftedLeftStackOfNormalVectors[shiftedleftstackofnormalvector] = struct{}{}

		if order > stage.ShiftedLeftStackOfNormalVectorOrder {
			stage.ShiftedLeftStackOfNormalVectorOrder = order
		}
		stage.ShiftedLeftStackOfNormalVector_stagedOrder[shiftedleftstackofnormalvector] = order
		stage.ShiftedLeftStackOfNormalVector_orderStaged[order] = shiftedleftstackofnormalvector
		stage.ShiftedLeftStackOfNormalVectorOrder++
	}
	stage.ShiftedLeftStackOfNormalVectors_mapString[shiftedleftstackofnormalvector.Name] = shiftedleftstackofnormalvector
}

// Unstage removes shiftedleftstackofnormalvector off the model stage
func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) Unstage(stage *Stage) *ShiftedLeftStackOfNormalVector {
	delete(stage.ShiftedLeftStackOfNormalVectors, shiftedleftstackofnormalvector)
	// issue1150
	// delete(stage.ShiftedLeftStackOfNormalVector_stagedOrder, shiftedleftstackofnormalvector)
	delete(stage.ShiftedLeftStackOfNormalVectors_mapString, shiftedleftstackofnormalvector.Name)

	return shiftedleftstackofnormalvector
}

// UnstageVoid removes shiftedleftstackofnormalvector off the model stage
func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) UnstageVoid(stage *Stage) {
	delete(stage.ShiftedLeftStackOfNormalVectors, shiftedleftstackofnormalvector)
	// issue1150
	// delete(stage.ShiftedLeftStackOfNormalVector_stagedOrder, shiftedleftstackofnormalvector)
	delete(stage.ShiftedLeftStackOfNormalVectors_mapString, shiftedleftstackofnormalvector.Name)
}

// commit shiftedleftstackofnormalvector to the back repo (if it is already staged)
func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) Commit(stage *Stage) *ShiftedLeftStackOfNormalVector {
	if _, ok := stage.ShiftedLeftStackOfNormalVectors[shiftedleftstackofnormalvector]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitShiftedLeftStackOfNormalVector(shiftedleftstackofnormalvector)
		}
	}
	return shiftedleftstackofnormalvector
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) CommitVoid(stage *Stage) {
	shiftedleftstackofnormalvector.Commit(stage)
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) StageVoid(stage *Stage) {
	shiftedleftstackofnormalvector.Stage(stage)
}

// Checkout shiftedleftstackofnormalvector to the back repo (if it is already staged)
func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) Checkout(stage *Stage) *ShiftedLeftStackOfNormalVector {
	if _, ok := stage.ShiftedLeftStackOfNormalVectors[shiftedleftstackofnormalvector]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutShiftedLeftStackOfNormalVector(shiftedleftstackofnormalvector)
		}
	}
	return shiftedleftstackofnormalvector
}

// for satisfaction of GongStruct interface
func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GetName() (res string) {
	return shiftedleftstackofnormalvector.Name
}

// for satisfaction of GongStruct interface
func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) SetName(name string) {
	shiftedleftstackofnormalvector.Name = name
}

// Stage puts stackgrowthcurveendarcshape to the model stage
func (stackgrowthcurveendarcshape *StackGrowthCurveEndArcShape) Stage(stage *Stage) *StackGrowthCurveEndArcShape {
	if _, ok := stage.StackGrowthCurveEndArcShapes[stackgrowthcurveendarcshape]; !ok {
		stage.StackGrowthCurveEndArcShapes[stackgrowthcurveendarcshape] = struct{}{}
		stage.StackGrowthCurveEndArcShape_stagedOrder[stackgrowthcurveendarcshape] = stage.StackGrowthCurveEndArcShapeOrder
		stage.StackGrowthCurveEndArcShape_orderStaged[stage.StackGrowthCurveEndArcShapeOrder] = stackgrowthcurveendarcshape
		stage.StackGrowthCurveEndArcShapeOrder++
	}
	stage.StackGrowthCurveEndArcShapes_mapString[stackgrowthcurveendarcshape.Name] = stackgrowthcurveendarcshape

	return stackgrowthcurveendarcshape
}

// StagePreserveOrder puts stackgrowthcurveendarcshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StackGrowthCurveEndArcShapeOrder
// - update stage.StackGrowthCurveEndArcShapeOrder accordingly
func (stackgrowthcurveendarcshape *StackGrowthCurveEndArcShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.StackGrowthCurveEndArcShapes[stackgrowthcurveendarcshape]; !ok {
		stage.StackGrowthCurveEndArcShapes[stackgrowthcurveendarcshape] = struct{}{}

		if order > stage.StackGrowthCurveEndArcShapeOrder {
			stage.StackGrowthCurveEndArcShapeOrder = order
		}
		stage.StackGrowthCurveEndArcShape_stagedOrder[stackgrowthcurveendarcshape] = order
		stage.StackGrowthCurveEndArcShape_orderStaged[order] = stackgrowthcurveendarcshape
		stage.StackGrowthCurveEndArcShapeOrder++
	}
	stage.StackGrowthCurveEndArcShapes_mapString[stackgrowthcurveendarcshape.Name] = stackgrowthcurveendarcshape
}

// Unstage removes stackgrowthcurveendarcshape off the model stage
func (stackgrowthcurveendarcshape *StackGrowthCurveEndArcShape) Unstage(stage *Stage) *StackGrowthCurveEndArcShape {
	delete(stage.StackGrowthCurveEndArcShapes, stackgrowthcurveendarcshape)
	// issue1150
	// delete(stage.StackGrowthCurveEndArcShape_stagedOrder, stackgrowthcurveendarcshape)
	delete(stage.StackGrowthCurveEndArcShapes_mapString, stackgrowthcurveendarcshape.Name)

	return stackgrowthcurveendarcshape
}

// UnstageVoid removes stackgrowthcurveendarcshape off the model stage
func (stackgrowthcurveendarcshape *StackGrowthCurveEndArcShape) UnstageVoid(stage *Stage) {
	delete(stage.StackGrowthCurveEndArcShapes, stackgrowthcurveendarcshape)
	// issue1150
	// delete(stage.StackGrowthCurveEndArcShape_stagedOrder, stackgrowthcurveendarcshape)
	delete(stage.StackGrowthCurveEndArcShapes_mapString, stackgrowthcurveendarcshape.Name)
}

// commit stackgrowthcurveendarcshape to the back repo (if it is already staged)
func (stackgrowthcurveendarcshape *StackGrowthCurveEndArcShape) Commit(stage *Stage) *StackGrowthCurveEndArcShape {
	if _, ok := stage.StackGrowthCurveEndArcShapes[stackgrowthcurveendarcshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitStackGrowthCurveEndArcShape(stackgrowthcurveendarcshape)
		}
	}
	return stackgrowthcurveendarcshape
}

func (stackgrowthcurveendarcshape *StackGrowthCurveEndArcShape) CommitVoid(stage *Stage) {
	stackgrowthcurveendarcshape.Commit(stage)
}

func (stackgrowthcurveendarcshape *StackGrowthCurveEndArcShape) StageVoid(stage *Stage) {
	stackgrowthcurveendarcshape.Stage(stage)
}

// Checkout stackgrowthcurveendarcshape to the back repo (if it is already staged)
func (stackgrowthcurveendarcshape *StackGrowthCurveEndArcShape) Checkout(stage *Stage) *StackGrowthCurveEndArcShape {
	if _, ok := stage.StackGrowthCurveEndArcShapes[stackgrowthcurveendarcshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutStackGrowthCurveEndArcShape(stackgrowthcurveendarcshape)
		}
	}
	return stackgrowthcurveendarcshape
}

// for satisfaction of GongStruct interface
func (stackgrowthcurveendarcshape *StackGrowthCurveEndArcShape) GetName() (res string) {
	return stackgrowthcurveendarcshape.Name
}

// for satisfaction of GongStruct interface
func (stackgrowthcurveendarcshape *StackGrowthCurveEndArcShape) SetName(name string) {
	stackgrowthcurveendarcshape.Name = name
}

// Stage puts stackgrowthcurvestartarcshape to the model stage
func (stackgrowthcurvestartarcshape *StackGrowthCurveStartArcShape) Stage(stage *Stage) *StackGrowthCurveStartArcShape {
	if _, ok := stage.StackGrowthCurveStartArcShapes[stackgrowthcurvestartarcshape]; !ok {
		stage.StackGrowthCurveStartArcShapes[stackgrowthcurvestartarcshape] = struct{}{}
		stage.StackGrowthCurveStartArcShape_stagedOrder[stackgrowthcurvestartarcshape] = stage.StackGrowthCurveStartArcShapeOrder
		stage.StackGrowthCurveStartArcShape_orderStaged[stage.StackGrowthCurveStartArcShapeOrder] = stackgrowthcurvestartarcshape
		stage.StackGrowthCurveStartArcShapeOrder++
	}
	stage.StackGrowthCurveStartArcShapes_mapString[stackgrowthcurvestartarcshape.Name] = stackgrowthcurvestartarcshape

	return stackgrowthcurvestartarcshape
}

// StagePreserveOrder puts stackgrowthcurvestartarcshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StackGrowthCurveStartArcShapeOrder
// - update stage.StackGrowthCurveStartArcShapeOrder accordingly
func (stackgrowthcurvestartarcshape *StackGrowthCurveStartArcShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.StackGrowthCurveStartArcShapes[stackgrowthcurvestartarcshape]; !ok {
		stage.StackGrowthCurveStartArcShapes[stackgrowthcurvestartarcshape] = struct{}{}

		if order > stage.StackGrowthCurveStartArcShapeOrder {
			stage.StackGrowthCurveStartArcShapeOrder = order
		}
		stage.StackGrowthCurveStartArcShape_stagedOrder[stackgrowthcurvestartarcshape] = order
		stage.StackGrowthCurveStartArcShape_orderStaged[order] = stackgrowthcurvestartarcshape
		stage.StackGrowthCurveStartArcShapeOrder++
	}
	stage.StackGrowthCurveStartArcShapes_mapString[stackgrowthcurvestartarcshape.Name] = stackgrowthcurvestartarcshape
}

// Unstage removes stackgrowthcurvestartarcshape off the model stage
func (stackgrowthcurvestartarcshape *StackGrowthCurveStartArcShape) Unstage(stage *Stage) *StackGrowthCurveStartArcShape {
	delete(stage.StackGrowthCurveStartArcShapes, stackgrowthcurvestartarcshape)
	// issue1150
	// delete(stage.StackGrowthCurveStartArcShape_stagedOrder, stackgrowthcurvestartarcshape)
	delete(stage.StackGrowthCurveStartArcShapes_mapString, stackgrowthcurvestartarcshape.Name)

	return stackgrowthcurvestartarcshape
}

// UnstageVoid removes stackgrowthcurvestartarcshape off the model stage
func (stackgrowthcurvestartarcshape *StackGrowthCurveStartArcShape) UnstageVoid(stage *Stage) {
	delete(stage.StackGrowthCurveStartArcShapes, stackgrowthcurvestartarcshape)
	// issue1150
	// delete(stage.StackGrowthCurveStartArcShape_stagedOrder, stackgrowthcurvestartarcshape)
	delete(stage.StackGrowthCurveStartArcShapes_mapString, stackgrowthcurvestartarcshape.Name)
}

// commit stackgrowthcurvestartarcshape to the back repo (if it is already staged)
func (stackgrowthcurvestartarcshape *StackGrowthCurveStartArcShape) Commit(stage *Stage) *StackGrowthCurveStartArcShape {
	if _, ok := stage.StackGrowthCurveStartArcShapes[stackgrowthcurvestartarcshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitStackGrowthCurveStartArcShape(stackgrowthcurvestartarcshape)
		}
	}
	return stackgrowthcurvestartarcshape
}

func (stackgrowthcurvestartarcshape *StackGrowthCurveStartArcShape) CommitVoid(stage *Stage) {
	stackgrowthcurvestartarcshape.Commit(stage)
}

func (stackgrowthcurvestartarcshape *StackGrowthCurveStartArcShape) StageVoid(stage *Stage) {
	stackgrowthcurvestartarcshape.Stage(stage)
}

// Checkout stackgrowthcurvestartarcshape to the back repo (if it is already staged)
func (stackgrowthcurvestartarcshape *StackGrowthCurveStartArcShape) Checkout(stage *Stage) *StackGrowthCurveStartArcShape {
	if _, ok := stage.StackGrowthCurveStartArcShapes[stackgrowthcurvestartarcshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutStackGrowthCurveStartArcShape(stackgrowthcurvestartarcshape)
		}
	}
	return stackgrowthcurvestartarcshape
}

// for satisfaction of GongStruct interface
func (stackgrowthcurvestartarcshape *StackGrowthCurveStartArcShape) GetName() (res string) {
	return stackgrowthcurvestartarcshape.Name
}

// for satisfaction of GongStruct interface
func (stackgrowthcurvestartarcshape *StackGrowthCurveStartArcShape) SetName(name string) {
	stackgrowthcurvestartarcshape.Name = name
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

// Stage puts topendarcshape to the model stage
func (topendarcshape *TopEndArcShape) Stage(stage *Stage) *TopEndArcShape {
	if _, ok := stage.TopEndArcShapes[topendarcshape]; !ok {
		stage.TopEndArcShapes[topendarcshape] = struct{}{}
		stage.TopEndArcShape_stagedOrder[topendarcshape] = stage.TopEndArcShapeOrder
		stage.TopEndArcShape_orderStaged[stage.TopEndArcShapeOrder] = topendarcshape
		stage.TopEndArcShapeOrder++
	}
	stage.TopEndArcShapes_mapString[topendarcshape.Name] = topendarcshape

	return topendarcshape
}

// StagePreserveOrder puts topendarcshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TopEndArcShapeOrder
// - update stage.TopEndArcShapeOrder accordingly
func (topendarcshape *TopEndArcShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TopEndArcShapes[topendarcshape]; !ok {
		stage.TopEndArcShapes[topendarcshape] = struct{}{}

		if order > stage.TopEndArcShapeOrder {
			stage.TopEndArcShapeOrder = order
		}
		stage.TopEndArcShape_stagedOrder[topendarcshape] = order
		stage.TopEndArcShape_orderStaged[order] = topendarcshape
		stage.TopEndArcShapeOrder++
	}
	stage.TopEndArcShapes_mapString[topendarcshape.Name] = topendarcshape
}

// Unstage removes topendarcshape off the model stage
func (topendarcshape *TopEndArcShape) Unstage(stage *Stage) *TopEndArcShape {
	delete(stage.TopEndArcShapes, topendarcshape)
	// issue1150
	// delete(stage.TopEndArcShape_stagedOrder, topendarcshape)
	delete(stage.TopEndArcShapes_mapString, topendarcshape.Name)

	return topendarcshape
}

// UnstageVoid removes topendarcshape off the model stage
func (topendarcshape *TopEndArcShape) UnstageVoid(stage *Stage) {
	delete(stage.TopEndArcShapes, topendarcshape)
	// issue1150
	// delete(stage.TopEndArcShape_stagedOrder, topendarcshape)
	delete(stage.TopEndArcShapes_mapString, topendarcshape.Name)
}

// commit topendarcshape to the back repo (if it is already staged)
func (topendarcshape *TopEndArcShape) Commit(stage *Stage) *TopEndArcShape {
	if _, ok := stage.TopEndArcShapes[topendarcshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTopEndArcShape(topendarcshape)
		}
	}
	return topendarcshape
}

func (topendarcshape *TopEndArcShape) CommitVoid(stage *Stage) {
	topendarcshape.Commit(stage)
}

func (topendarcshape *TopEndArcShape) StageVoid(stage *Stage) {
	topendarcshape.Stage(stage)
}

// Checkout topendarcshape to the back repo (if it is already staged)
func (topendarcshape *TopEndArcShape) Checkout(stage *Stage) *TopEndArcShape {
	if _, ok := stage.TopEndArcShapes[topendarcshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTopEndArcShape(topendarcshape)
		}
	}
	return topendarcshape
}

// for satisfaction of GongStruct interface
func (topendarcshape *TopEndArcShape) GetName() (res string) {
	return topendarcshape.Name
}

// for satisfaction of GongStruct interface
func (topendarcshape *TopEndArcShape) SetName(name string) {
	topendarcshape.Name = name
}

// Stage puts topendarcshapegrid to the model stage
func (topendarcshapegrid *TopEndArcShapeGrid) Stage(stage *Stage) *TopEndArcShapeGrid {
	if _, ok := stage.TopEndArcShapeGrids[topendarcshapegrid]; !ok {
		stage.TopEndArcShapeGrids[topendarcshapegrid] = struct{}{}
		stage.TopEndArcShapeGrid_stagedOrder[topendarcshapegrid] = stage.TopEndArcShapeGridOrder
		stage.TopEndArcShapeGrid_orderStaged[stage.TopEndArcShapeGridOrder] = topendarcshapegrid
		stage.TopEndArcShapeGridOrder++
	}
	stage.TopEndArcShapeGrids_mapString[topendarcshapegrid.Name] = topendarcshapegrid

	return topendarcshapegrid
}

// StagePreserveOrder puts topendarcshapegrid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TopEndArcShapeGridOrder
// - update stage.TopEndArcShapeGridOrder accordingly
func (topendarcshapegrid *TopEndArcShapeGrid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TopEndArcShapeGrids[topendarcshapegrid]; !ok {
		stage.TopEndArcShapeGrids[topendarcshapegrid] = struct{}{}

		if order > stage.TopEndArcShapeGridOrder {
			stage.TopEndArcShapeGridOrder = order
		}
		stage.TopEndArcShapeGrid_stagedOrder[topendarcshapegrid] = order
		stage.TopEndArcShapeGrid_orderStaged[order] = topendarcshapegrid
		stage.TopEndArcShapeGridOrder++
	}
	stage.TopEndArcShapeGrids_mapString[topendarcshapegrid.Name] = topendarcshapegrid
}

// Unstage removes topendarcshapegrid off the model stage
func (topendarcshapegrid *TopEndArcShapeGrid) Unstage(stage *Stage) *TopEndArcShapeGrid {
	delete(stage.TopEndArcShapeGrids, topendarcshapegrid)
	// issue1150
	// delete(stage.TopEndArcShapeGrid_stagedOrder, topendarcshapegrid)
	delete(stage.TopEndArcShapeGrids_mapString, topendarcshapegrid.Name)

	return topendarcshapegrid
}

// UnstageVoid removes topendarcshapegrid off the model stage
func (topendarcshapegrid *TopEndArcShapeGrid) UnstageVoid(stage *Stage) {
	delete(stage.TopEndArcShapeGrids, topendarcshapegrid)
	// issue1150
	// delete(stage.TopEndArcShapeGrid_stagedOrder, topendarcshapegrid)
	delete(stage.TopEndArcShapeGrids_mapString, topendarcshapegrid.Name)
}

// commit topendarcshapegrid to the back repo (if it is already staged)
func (topendarcshapegrid *TopEndArcShapeGrid) Commit(stage *Stage) *TopEndArcShapeGrid {
	if _, ok := stage.TopEndArcShapeGrids[topendarcshapegrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTopEndArcShapeGrid(topendarcshapegrid)
		}
	}
	return topendarcshapegrid
}

func (topendarcshapegrid *TopEndArcShapeGrid) CommitVoid(stage *Stage) {
	topendarcshapegrid.Commit(stage)
}

func (topendarcshapegrid *TopEndArcShapeGrid) StageVoid(stage *Stage) {
	topendarcshapegrid.Stage(stage)
}

// Checkout topendarcshapegrid to the back repo (if it is already staged)
func (topendarcshapegrid *TopEndArcShapeGrid) Checkout(stage *Stage) *TopEndArcShapeGrid {
	if _, ok := stage.TopEndArcShapeGrids[topendarcshapegrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTopEndArcShapeGrid(topendarcshapegrid)
		}
	}
	return topendarcshapegrid
}

// for satisfaction of GongStruct interface
func (topendarcshapegrid *TopEndArcShapeGrid) GetName() (res string) {
	return topendarcshapegrid.Name
}

// for satisfaction of GongStruct interface
func (topendarcshapegrid *TopEndArcShapeGrid) SetName(name string) {
	topendarcshapegrid.Name = name
}

// Stage puts topgrowthcurve2d to the model stage
func (topgrowthcurve2d *TopGrowthCurve2D) Stage(stage *Stage) *TopGrowthCurve2D {
	if _, ok := stage.TopGrowthCurve2Ds[topgrowthcurve2d]; !ok {
		stage.TopGrowthCurve2Ds[topgrowthcurve2d] = struct{}{}
		stage.TopGrowthCurve2D_stagedOrder[topgrowthcurve2d] = stage.TopGrowthCurve2DOrder
		stage.TopGrowthCurve2D_orderStaged[stage.TopGrowthCurve2DOrder] = topgrowthcurve2d
		stage.TopGrowthCurve2DOrder++
	}
	stage.TopGrowthCurve2Ds_mapString[topgrowthcurve2d.Name] = topgrowthcurve2d

	return topgrowthcurve2d
}

// StagePreserveOrder puts topgrowthcurve2d to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TopGrowthCurve2DOrder
// - update stage.TopGrowthCurve2DOrder accordingly
func (topgrowthcurve2d *TopGrowthCurve2D) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TopGrowthCurve2Ds[topgrowthcurve2d]; !ok {
		stage.TopGrowthCurve2Ds[topgrowthcurve2d] = struct{}{}

		if order > stage.TopGrowthCurve2DOrder {
			stage.TopGrowthCurve2DOrder = order
		}
		stage.TopGrowthCurve2D_stagedOrder[topgrowthcurve2d] = order
		stage.TopGrowthCurve2D_orderStaged[order] = topgrowthcurve2d
		stage.TopGrowthCurve2DOrder++
	}
	stage.TopGrowthCurve2Ds_mapString[topgrowthcurve2d.Name] = topgrowthcurve2d
}

// Unstage removes topgrowthcurve2d off the model stage
func (topgrowthcurve2d *TopGrowthCurve2D) Unstage(stage *Stage) *TopGrowthCurve2D {
	delete(stage.TopGrowthCurve2Ds, topgrowthcurve2d)
	// issue1150
	// delete(stage.TopGrowthCurve2D_stagedOrder, topgrowthcurve2d)
	delete(stage.TopGrowthCurve2Ds_mapString, topgrowthcurve2d.Name)

	return topgrowthcurve2d
}

// UnstageVoid removes topgrowthcurve2d off the model stage
func (topgrowthcurve2d *TopGrowthCurve2D) UnstageVoid(stage *Stage) {
	delete(stage.TopGrowthCurve2Ds, topgrowthcurve2d)
	// issue1150
	// delete(stage.TopGrowthCurve2D_stagedOrder, topgrowthcurve2d)
	delete(stage.TopGrowthCurve2Ds_mapString, topgrowthcurve2d.Name)
}

// commit topgrowthcurve2d to the back repo (if it is already staged)
func (topgrowthcurve2d *TopGrowthCurve2D) Commit(stage *Stage) *TopGrowthCurve2D {
	if _, ok := stage.TopGrowthCurve2Ds[topgrowthcurve2d]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTopGrowthCurve2D(topgrowthcurve2d)
		}
	}
	return topgrowthcurve2d
}

func (topgrowthcurve2d *TopGrowthCurve2D) CommitVoid(stage *Stage) {
	topgrowthcurve2d.Commit(stage)
}

func (topgrowthcurve2d *TopGrowthCurve2D) StageVoid(stage *Stage) {
	topgrowthcurve2d.Stage(stage)
}

// Checkout topgrowthcurve2d to the back repo (if it is already staged)
func (topgrowthcurve2d *TopGrowthCurve2D) Checkout(stage *Stage) *TopGrowthCurve2D {
	if _, ok := stage.TopGrowthCurve2Ds[topgrowthcurve2d]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTopGrowthCurve2D(topgrowthcurve2d)
		}
	}
	return topgrowthcurve2d
}

// for satisfaction of GongStruct interface
func (topgrowthcurve2d *TopGrowthCurve2D) GetName() (res string) {
	return topgrowthcurve2d.Name
}

// for satisfaction of GongStruct interface
func (topgrowthcurve2d *TopGrowthCurve2D) SetName(name string) {
	topgrowthcurve2d.Name = name
}

// Stage puts topstackgrowthcurveendarcshape to the model stage
func (topstackgrowthcurveendarcshape *TopStackGrowthCurveEndArcShape) Stage(stage *Stage) *TopStackGrowthCurveEndArcShape {
	if _, ok := stage.TopStackGrowthCurveEndArcShapes[topstackgrowthcurveendarcshape]; !ok {
		stage.TopStackGrowthCurveEndArcShapes[topstackgrowthcurveendarcshape] = struct{}{}
		stage.TopStackGrowthCurveEndArcShape_stagedOrder[topstackgrowthcurveendarcshape] = stage.TopStackGrowthCurveEndArcShapeOrder
		stage.TopStackGrowthCurveEndArcShape_orderStaged[stage.TopStackGrowthCurveEndArcShapeOrder] = topstackgrowthcurveendarcshape
		stage.TopStackGrowthCurveEndArcShapeOrder++
	}
	stage.TopStackGrowthCurveEndArcShapes_mapString[topstackgrowthcurveendarcshape.Name] = topstackgrowthcurveendarcshape

	return topstackgrowthcurveendarcshape
}

// StagePreserveOrder puts topstackgrowthcurveendarcshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TopStackGrowthCurveEndArcShapeOrder
// - update stage.TopStackGrowthCurveEndArcShapeOrder accordingly
func (topstackgrowthcurveendarcshape *TopStackGrowthCurveEndArcShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TopStackGrowthCurveEndArcShapes[topstackgrowthcurveendarcshape]; !ok {
		stage.TopStackGrowthCurveEndArcShapes[topstackgrowthcurveendarcshape] = struct{}{}

		if order > stage.TopStackGrowthCurveEndArcShapeOrder {
			stage.TopStackGrowthCurveEndArcShapeOrder = order
		}
		stage.TopStackGrowthCurveEndArcShape_stagedOrder[topstackgrowthcurveendarcshape] = order
		stage.TopStackGrowthCurveEndArcShape_orderStaged[order] = topstackgrowthcurveendarcshape
		stage.TopStackGrowthCurveEndArcShapeOrder++
	}
	stage.TopStackGrowthCurveEndArcShapes_mapString[topstackgrowthcurveendarcshape.Name] = topstackgrowthcurveendarcshape
}

// Unstage removes topstackgrowthcurveendarcshape off the model stage
func (topstackgrowthcurveendarcshape *TopStackGrowthCurveEndArcShape) Unstage(stage *Stage) *TopStackGrowthCurveEndArcShape {
	delete(stage.TopStackGrowthCurveEndArcShapes, topstackgrowthcurveendarcshape)
	// issue1150
	// delete(stage.TopStackGrowthCurveEndArcShape_stagedOrder, topstackgrowthcurveendarcshape)
	delete(stage.TopStackGrowthCurveEndArcShapes_mapString, topstackgrowthcurveendarcshape.Name)

	return topstackgrowthcurveendarcshape
}

// UnstageVoid removes topstackgrowthcurveendarcshape off the model stage
func (topstackgrowthcurveendarcshape *TopStackGrowthCurveEndArcShape) UnstageVoid(stage *Stage) {
	delete(stage.TopStackGrowthCurveEndArcShapes, topstackgrowthcurveendarcshape)
	// issue1150
	// delete(stage.TopStackGrowthCurveEndArcShape_stagedOrder, topstackgrowthcurveendarcshape)
	delete(stage.TopStackGrowthCurveEndArcShapes_mapString, topstackgrowthcurveendarcshape.Name)
}

// commit topstackgrowthcurveendarcshape to the back repo (if it is already staged)
func (topstackgrowthcurveendarcshape *TopStackGrowthCurveEndArcShape) Commit(stage *Stage) *TopStackGrowthCurveEndArcShape {
	if _, ok := stage.TopStackGrowthCurveEndArcShapes[topstackgrowthcurveendarcshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTopStackGrowthCurveEndArcShape(topstackgrowthcurveendarcshape)
		}
	}
	return topstackgrowthcurveendarcshape
}

func (topstackgrowthcurveendarcshape *TopStackGrowthCurveEndArcShape) CommitVoid(stage *Stage) {
	topstackgrowthcurveendarcshape.Commit(stage)
}

func (topstackgrowthcurveendarcshape *TopStackGrowthCurveEndArcShape) StageVoid(stage *Stage) {
	topstackgrowthcurveendarcshape.Stage(stage)
}

// Checkout topstackgrowthcurveendarcshape to the back repo (if it is already staged)
func (topstackgrowthcurveendarcshape *TopStackGrowthCurveEndArcShape) Checkout(stage *Stage) *TopStackGrowthCurveEndArcShape {
	if _, ok := stage.TopStackGrowthCurveEndArcShapes[topstackgrowthcurveendarcshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTopStackGrowthCurveEndArcShape(topstackgrowthcurveendarcshape)
		}
	}
	return topstackgrowthcurveendarcshape
}

// for satisfaction of GongStruct interface
func (topstackgrowthcurveendarcshape *TopStackGrowthCurveEndArcShape) GetName() (res string) {
	return topstackgrowthcurveendarcshape.Name
}

// for satisfaction of GongStruct interface
func (topstackgrowthcurveendarcshape *TopStackGrowthCurveEndArcShape) SetName(name string) {
	topstackgrowthcurveendarcshape.Name = name
}

// Stage puts topstackgrowthcurvestartarcshape to the model stage
func (topstackgrowthcurvestartarcshape *TopStackGrowthCurveStartArcShape) Stage(stage *Stage) *TopStackGrowthCurveStartArcShape {
	if _, ok := stage.TopStackGrowthCurveStartArcShapes[topstackgrowthcurvestartarcshape]; !ok {
		stage.TopStackGrowthCurveStartArcShapes[topstackgrowthcurvestartarcshape] = struct{}{}
		stage.TopStackGrowthCurveStartArcShape_stagedOrder[topstackgrowthcurvestartarcshape] = stage.TopStackGrowthCurveStartArcShapeOrder
		stage.TopStackGrowthCurveStartArcShape_orderStaged[stage.TopStackGrowthCurveStartArcShapeOrder] = topstackgrowthcurvestartarcshape
		stage.TopStackGrowthCurveStartArcShapeOrder++
	}
	stage.TopStackGrowthCurveStartArcShapes_mapString[topstackgrowthcurvestartarcshape.Name] = topstackgrowthcurvestartarcshape

	return topstackgrowthcurvestartarcshape
}

// StagePreserveOrder puts topstackgrowthcurvestartarcshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TopStackGrowthCurveStartArcShapeOrder
// - update stage.TopStackGrowthCurveStartArcShapeOrder accordingly
func (topstackgrowthcurvestartarcshape *TopStackGrowthCurveStartArcShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TopStackGrowthCurveStartArcShapes[topstackgrowthcurvestartarcshape]; !ok {
		stage.TopStackGrowthCurveStartArcShapes[topstackgrowthcurvestartarcshape] = struct{}{}

		if order > stage.TopStackGrowthCurveStartArcShapeOrder {
			stage.TopStackGrowthCurveStartArcShapeOrder = order
		}
		stage.TopStackGrowthCurveStartArcShape_stagedOrder[topstackgrowthcurvestartarcshape] = order
		stage.TopStackGrowthCurveStartArcShape_orderStaged[order] = topstackgrowthcurvestartarcshape
		stage.TopStackGrowthCurveStartArcShapeOrder++
	}
	stage.TopStackGrowthCurveStartArcShapes_mapString[topstackgrowthcurvestartarcshape.Name] = topstackgrowthcurvestartarcshape
}

// Unstage removes topstackgrowthcurvestartarcshape off the model stage
func (topstackgrowthcurvestartarcshape *TopStackGrowthCurveStartArcShape) Unstage(stage *Stage) *TopStackGrowthCurveStartArcShape {
	delete(stage.TopStackGrowthCurveStartArcShapes, topstackgrowthcurvestartarcshape)
	// issue1150
	// delete(stage.TopStackGrowthCurveStartArcShape_stagedOrder, topstackgrowthcurvestartarcshape)
	delete(stage.TopStackGrowthCurveStartArcShapes_mapString, topstackgrowthcurvestartarcshape.Name)

	return topstackgrowthcurvestartarcshape
}

// UnstageVoid removes topstackgrowthcurvestartarcshape off the model stage
func (topstackgrowthcurvestartarcshape *TopStackGrowthCurveStartArcShape) UnstageVoid(stage *Stage) {
	delete(stage.TopStackGrowthCurveStartArcShapes, topstackgrowthcurvestartarcshape)
	// issue1150
	// delete(stage.TopStackGrowthCurveStartArcShape_stagedOrder, topstackgrowthcurvestartarcshape)
	delete(stage.TopStackGrowthCurveStartArcShapes_mapString, topstackgrowthcurvestartarcshape.Name)
}

// commit topstackgrowthcurvestartarcshape to the back repo (if it is already staged)
func (topstackgrowthcurvestartarcshape *TopStackGrowthCurveStartArcShape) Commit(stage *Stage) *TopStackGrowthCurveStartArcShape {
	if _, ok := stage.TopStackGrowthCurveStartArcShapes[topstackgrowthcurvestartarcshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTopStackGrowthCurveStartArcShape(topstackgrowthcurvestartarcshape)
		}
	}
	return topstackgrowthcurvestartarcshape
}

func (topstackgrowthcurvestartarcshape *TopStackGrowthCurveStartArcShape) CommitVoid(stage *Stage) {
	topstackgrowthcurvestartarcshape.Commit(stage)
}

func (topstackgrowthcurvestartarcshape *TopStackGrowthCurveStartArcShape) StageVoid(stage *Stage) {
	topstackgrowthcurvestartarcshape.Stage(stage)
}

// Checkout topstackgrowthcurvestartarcshape to the back repo (if it is already staged)
func (topstackgrowthcurvestartarcshape *TopStackGrowthCurveStartArcShape) Checkout(stage *Stage) *TopStackGrowthCurveStartArcShape {
	if _, ok := stage.TopStackGrowthCurveStartArcShapes[topstackgrowthcurvestartarcshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTopStackGrowthCurveStartArcShape(topstackgrowthcurvestartarcshape)
		}
	}
	return topstackgrowthcurvestartarcshape
}

// for satisfaction of GongStruct interface
func (topstackgrowthcurvestartarcshape *TopStackGrowthCurveStartArcShape) GetName() (res string) {
	return topstackgrowthcurvestartarcshape.Name
}

// for satisfaction of GongStruct interface
func (topstackgrowthcurvestartarcshape *TopStackGrowthCurveStartArcShape) SetName(name string) {
	topstackgrowthcurvestartarcshape.Name = name
}

// Stage puts topstackofgrowthcurve to the model stage
func (topstackofgrowthcurve *TopStackOfGrowthCurve) Stage(stage *Stage) *TopStackOfGrowthCurve {
	if _, ok := stage.TopStackOfGrowthCurves[topstackofgrowthcurve]; !ok {
		stage.TopStackOfGrowthCurves[topstackofgrowthcurve] = struct{}{}
		stage.TopStackOfGrowthCurve_stagedOrder[topstackofgrowthcurve] = stage.TopStackOfGrowthCurveOrder
		stage.TopStackOfGrowthCurve_orderStaged[stage.TopStackOfGrowthCurveOrder] = topstackofgrowthcurve
		stage.TopStackOfGrowthCurveOrder++
	}
	stage.TopStackOfGrowthCurves_mapString[topstackofgrowthcurve.Name] = topstackofgrowthcurve

	return topstackofgrowthcurve
}

// StagePreserveOrder puts topstackofgrowthcurve to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TopStackOfGrowthCurveOrder
// - update stage.TopStackOfGrowthCurveOrder accordingly
func (topstackofgrowthcurve *TopStackOfGrowthCurve) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TopStackOfGrowthCurves[topstackofgrowthcurve]; !ok {
		stage.TopStackOfGrowthCurves[topstackofgrowthcurve] = struct{}{}

		if order > stage.TopStackOfGrowthCurveOrder {
			stage.TopStackOfGrowthCurveOrder = order
		}
		stage.TopStackOfGrowthCurve_stagedOrder[topstackofgrowthcurve] = order
		stage.TopStackOfGrowthCurve_orderStaged[order] = topstackofgrowthcurve
		stage.TopStackOfGrowthCurveOrder++
	}
	stage.TopStackOfGrowthCurves_mapString[topstackofgrowthcurve.Name] = topstackofgrowthcurve
}

// Unstage removes topstackofgrowthcurve off the model stage
func (topstackofgrowthcurve *TopStackOfGrowthCurve) Unstage(stage *Stage) *TopStackOfGrowthCurve {
	delete(stage.TopStackOfGrowthCurves, topstackofgrowthcurve)
	// issue1150
	// delete(stage.TopStackOfGrowthCurve_stagedOrder, topstackofgrowthcurve)
	delete(stage.TopStackOfGrowthCurves_mapString, topstackofgrowthcurve.Name)

	return topstackofgrowthcurve
}

// UnstageVoid removes topstackofgrowthcurve off the model stage
func (topstackofgrowthcurve *TopStackOfGrowthCurve) UnstageVoid(stage *Stage) {
	delete(stage.TopStackOfGrowthCurves, topstackofgrowthcurve)
	// issue1150
	// delete(stage.TopStackOfGrowthCurve_stagedOrder, topstackofgrowthcurve)
	delete(stage.TopStackOfGrowthCurves_mapString, topstackofgrowthcurve.Name)
}

// commit topstackofgrowthcurve to the back repo (if it is already staged)
func (topstackofgrowthcurve *TopStackOfGrowthCurve) Commit(stage *Stage) *TopStackOfGrowthCurve {
	if _, ok := stage.TopStackOfGrowthCurves[topstackofgrowthcurve]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTopStackOfGrowthCurve(topstackofgrowthcurve)
		}
	}
	return topstackofgrowthcurve
}

func (topstackofgrowthcurve *TopStackOfGrowthCurve) CommitVoid(stage *Stage) {
	topstackofgrowthcurve.Commit(stage)
}

func (topstackofgrowthcurve *TopStackOfGrowthCurve) StageVoid(stage *Stage) {
	topstackofgrowthcurve.Stage(stage)
}

// Checkout topstackofgrowthcurve to the back repo (if it is already staged)
func (topstackofgrowthcurve *TopStackOfGrowthCurve) Checkout(stage *Stage) *TopStackOfGrowthCurve {
	if _, ok := stage.TopStackOfGrowthCurves[topstackofgrowthcurve]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTopStackOfGrowthCurve(topstackofgrowthcurve)
		}
	}
	return topstackofgrowthcurve
}

// for satisfaction of GongStruct interface
func (topstackofgrowthcurve *TopStackOfGrowthCurve) GetName() (res string) {
	return topstackofgrowthcurve.Name
}

// for satisfaction of GongStruct interface
func (topstackofgrowthcurve *TopStackOfGrowthCurve) SetName(name string) {
	topstackofgrowthcurve.Name = name
}

// Stage puts topstartarcshape to the model stage
func (topstartarcshape *TopStartArcShape) Stage(stage *Stage) *TopStartArcShape {
	if _, ok := stage.TopStartArcShapes[topstartarcshape]; !ok {
		stage.TopStartArcShapes[topstartarcshape] = struct{}{}
		stage.TopStartArcShape_stagedOrder[topstartarcshape] = stage.TopStartArcShapeOrder
		stage.TopStartArcShape_orderStaged[stage.TopStartArcShapeOrder] = topstartarcshape
		stage.TopStartArcShapeOrder++
	}
	stage.TopStartArcShapes_mapString[topstartarcshape.Name] = topstartarcshape

	return topstartarcshape
}

// StagePreserveOrder puts topstartarcshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TopStartArcShapeOrder
// - update stage.TopStartArcShapeOrder accordingly
func (topstartarcshape *TopStartArcShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TopStartArcShapes[topstartarcshape]; !ok {
		stage.TopStartArcShapes[topstartarcshape] = struct{}{}

		if order > stage.TopStartArcShapeOrder {
			stage.TopStartArcShapeOrder = order
		}
		stage.TopStartArcShape_stagedOrder[topstartarcshape] = order
		stage.TopStartArcShape_orderStaged[order] = topstartarcshape
		stage.TopStartArcShapeOrder++
	}
	stage.TopStartArcShapes_mapString[topstartarcshape.Name] = topstartarcshape
}

// Unstage removes topstartarcshape off the model stage
func (topstartarcshape *TopStartArcShape) Unstage(stage *Stage) *TopStartArcShape {
	delete(stage.TopStartArcShapes, topstartarcshape)
	// issue1150
	// delete(stage.TopStartArcShape_stagedOrder, topstartarcshape)
	delete(stage.TopStartArcShapes_mapString, topstartarcshape.Name)

	return topstartarcshape
}

// UnstageVoid removes topstartarcshape off the model stage
func (topstartarcshape *TopStartArcShape) UnstageVoid(stage *Stage) {
	delete(stage.TopStartArcShapes, topstartarcshape)
	// issue1150
	// delete(stage.TopStartArcShape_stagedOrder, topstartarcshape)
	delete(stage.TopStartArcShapes_mapString, topstartarcshape.Name)
}

// commit topstartarcshape to the back repo (if it is already staged)
func (topstartarcshape *TopStartArcShape) Commit(stage *Stage) *TopStartArcShape {
	if _, ok := stage.TopStartArcShapes[topstartarcshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTopStartArcShape(topstartarcshape)
		}
	}
	return topstartarcshape
}

func (topstartarcshape *TopStartArcShape) CommitVoid(stage *Stage) {
	topstartarcshape.Commit(stage)
}

func (topstartarcshape *TopStartArcShape) StageVoid(stage *Stage) {
	topstartarcshape.Stage(stage)
}

// Checkout topstartarcshape to the back repo (if it is already staged)
func (topstartarcshape *TopStartArcShape) Checkout(stage *Stage) *TopStartArcShape {
	if _, ok := stage.TopStartArcShapes[topstartarcshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTopStartArcShape(topstartarcshape)
		}
	}
	return topstartarcshape
}

// for satisfaction of GongStruct interface
func (topstartarcshape *TopStartArcShape) GetName() (res string) {
	return topstartarcshape.Name
}

// for satisfaction of GongStruct interface
func (topstartarcshape *TopStartArcShape) SetName(name string) {
	topstartarcshape.Name = name
}

// Stage puts topstartarcshapegrid to the model stage
func (topstartarcshapegrid *TopStartArcShapeGrid) Stage(stage *Stage) *TopStartArcShapeGrid {
	if _, ok := stage.TopStartArcShapeGrids[topstartarcshapegrid]; !ok {
		stage.TopStartArcShapeGrids[topstartarcshapegrid] = struct{}{}
		stage.TopStartArcShapeGrid_stagedOrder[topstartarcshapegrid] = stage.TopStartArcShapeGridOrder
		stage.TopStartArcShapeGrid_orderStaged[stage.TopStartArcShapeGridOrder] = topstartarcshapegrid
		stage.TopStartArcShapeGridOrder++
	}
	stage.TopStartArcShapeGrids_mapString[topstartarcshapegrid.Name] = topstartarcshapegrid

	return topstartarcshapegrid
}

// StagePreserveOrder puts topstartarcshapegrid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TopStartArcShapeGridOrder
// - update stage.TopStartArcShapeGridOrder accordingly
func (topstartarcshapegrid *TopStartArcShapeGrid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TopStartArcShapeGrids[topstartarcshapegrid]; !ok {
		stage.TopStartArcShapeGrids[topstartarcshapegrid] = struct{}{}

		if order > stage.TopStartArcShapeGridOrder {
			stage.TopStartArcShapeGridOrder = order
		}
		stage.TopStartArcShapeGrid_stagedOrder[topstartarcshapegrid] = order
		stage.TopStartArcShapeGrid_orderStaged[order] = topstartarcshapegrid
		stage.TopStartArcShapeGridOrder++
	}
	stage.TopStartArcShapeGrids_mapString[topstartarcshapegrid.Name] = topstartarcshapegrid
}

// Unstage removes topstartarcshapegrid off the model stage
func (topstartarcshapegrid *TopStartArcShapeGrid) Unstage(stage *Stage) *TopStartArcShapeGrid {
	delete(stage.TopStartArcShapeGrids, topstartarcshapegrid)
	// issue1150
	// delete(stage.TopStartArcShapeGrid_stagedOrder, topstartarcshapegrid)
	delete(stage.TopStartArcShapeGrids_mapString, topstartarcshapegrid.Name)

	return topstartarcshapegrid
}

// UnstageVoid removes topstartarcshapegrid off the model stage
func (topstartarcshapegrid *TopStartArcShapeGrid) UnstageVoid(stage *Stage) {
	delete(stage.TopStartArcShapeGrids, topstartarcshapegrid)
	// issue1150
	// delete(stage.TopStartArcShapeGrid_stagedOrder, topstartarcshapegrid)
	delete(stage.TopStartArcShapeGrids_mapString, topstartarcshapegrid.Name)
}

// commit topstartarcshapegrid to the back repo (if it is already staged)
func (topstartarcshapegrid *TopStartArcShapeGrid) Commit(stage *Stage) *TopStartArcShapeGrid {
	if _, ok := stage.TopStartArcShapeGrids[topstartarcshapegrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTopStartArcShapeGrid(topstartarcshapegrid)
		}
	}
	return topstartarcshapegrid
}

func (topstartarcshapegrid *TopStartArcShapeGrid) CommitVoid(stage *Stage) {
	topstartarcshapegrid.Commit(stage)
}

func (topstartarcshapegrid *TopStartArcShapeGrid) StageVoid(stage *Stage) {
	topstartarcshapegrid.Stage(stage)
}

// Checkout topstartarcshapegrid to the back repo (if it is already staged)
func (topstartarcshapegrid *TopStartArcShapeGrid) Checkout(stage *Stage) *TopStartArcShapeGrid {
	if _, ok := stage.TopStartArcShapeGrids[topstartarcshapegrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTopStartArcShapeGrid(topstartarcshapegrid)
		}
	}
	return topstartarcshapegrid
}

// for satisfaction of GongStruct interface
func (topstartarcshapegrid *TopStartArcShapeGrid) GetName() (res string) {
	return topstartarcshapegrid.Name
}

// for satisfaction of GongStruct interface
func (topstartarcshapegrid *TopStartArcShapeGrid) SetName(name string) {
	topstartarcshapegrid.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMArcNormalVectorShape(ArcNormalVectorShape *ArcNormalVectorShape)
	CreateORMArcNormalVectorShapeGrid(ArcNormalVectorShapeGrid *ArcNormalVectorShapeGrid)
	CreateORMAxesShape(AxesShape *AxesShape)
	CreateORMBaseVectorShape(BaseVectorShape *BaseVectorShape)
	CreateORMBaseVectorShapeGrid(BaseVectorShapeGrid *BaseVectorShapeGrid)
	CreateORMCircleGridShape(CircleGridShape *CircleGridShape)
	CreateORMEndArcShape(EndArcShape *EndArcShape)
	CreateORMEndArcShapeGrid(EndArcShapeGrid *EndArcShapeGrid)
	CreateORMExplanationTextShape(ExplanationTextShape *ExplanationTextShape)
	CreateORMGridPathShape(GridPathShape *GridPathShape)
	CreateORMGrowthCurve2D(GrowthCurve2D *GrowthCurve2D)
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
	CreateORMShiftedLeftStackGrowthCurveEndArcShape(ShiftedLeftStackGrowthCurveEndArcShape *ShiftedLeftStackGrowthCurveEndArcShape)
	CreateORMShiftedLeftStackGrowthCurveStartArcShape(ShiftedLeftStackGrowthCurveStartArcShape *ShiftedLeftStackGrowthCurveStartArcShape)
	CreateORMShiftedLeftStackNormalVector(ShiftedLeftStackNormalVector *ShiftedLeftStackNormalVector)
	CreateORMShiftedLeftStackOfGrowthCurve(ShiftedLeftStackOfGrowthCurve *ShiftedLeftStackOfGrowthCurve)
	CreateORMShiftedLeftStackOfNormalVector(ShiftedLeftStackOfNormalVector *ShiftedLeftStackOfNormalVector)
	CreateORMStackGrowthCurveEndArcShape(StackGrowthCurveEndArcShape *StackGrowthCurveEndArcShape)
	CreateORMStackGrowthCurveStartArcShape(StackGrowthCurveStartArcShape *StackGrowthCurveStartArcShape)
	CreateORMStackOfGrowthCurve(StackOfGrowthCurve *StackOfGrowthCurve)
	CreateORMStartArcShape(StartArcShape *StartArcShape)
	CreateORMStartArcShapeGrid(StartArcShapeGrid *StartArcShapeGrid)
	CreateORMTopEndArcShape(TopEndArcShape *TopEndArcShape)
	CreateORMTopEndArcShapeGrid(TopEndArcShapeGrid *TopEndArcShapeGrid)
	CreateORMTopGrowthCurve2D(TopGrowthCurve2D *TopGrowthCurve2D)
	CreateORMTopStackGrowthCurveEndArcShape(TopStackGrowthCurveEndArcShape *TopStackGrowthCurveEndArcShape)
	CreateORMTopStackGrowthCurveStartArcShape(TopStackGrowthCurveStartArcShape *TopStackGrowthCurveStartArcShape)
	CreateORMTopStackOfGrowthCurve(TopStackOfGrowthCurve *TopStackOfGrowthCurve)
	CreateORMTopStartArcShape(TopStartArcShape *TopStartArcShape)
	CreateORMTopStartArcShapeGrid(TopStartArcShapeGrid *TopStartArcShapeGrid)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMArcNormalVectorShape(ArcNormalVectorShape *ArcNormalVectorShape)
	DeleteORMArcNormalVectorShapeGrid(ArcNormalVectorShapeGrid *ArcNormalVectorShapeGrid)
	DeleteORMAxesShape(AxesShape *AxesShape)
	DeleteORMBaseVectorShape(BaseVectorShape *BaseVectorShape)
	DeleteORMBaseVectorShapeGrid(BaseVectorShapeGrid *BaseVectorShapeGrid)
	DeleteORMCircleGridShape(CircleGridShape *CircleGridShape)
	DeleteORMEndArcShape(EndArcShape *EndArcShape)
	DeleteORMEndArcShapeGrid(EndArcShapeGrid *EndArcShapeGrid)
	DeleteORMExplanationTextShape(ExplanationTextShape *ExplanationTextShape)
	DeleteORMGridPathShape(GridPathShape *GridPathShape)
	DeleteORMGrowthCurve2D(GrowthCurve2D *GrowthCurve2D)
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
	DeleteORMShiftedLeftStackGrowthCurveEndArcShape(ShiftedLeftStackGrowthCurveEndArcShape *ShiftedLeftStackGrowthCurveEndArcShape)
	DeleteORMShiftedLeftStackGrowthCurveStartArcShape(ShiftedLeftStackGrowthCurveStartArcShape *ShiftedLeftStackGrowthCurveStartArcShape)
	DeleteORMShiftedLeftStackNormalVector(ShiftedLeftStackNormalVector *ShiftedLeftStackNormalVector)
	DeleteORMShiftedLeftStackOfGrowthCurve(ShiftedLeftStackOfGrowthCurve *ShiftedLeftStackOfGrowthCurve)
	DeleteORMShiftedLeftStackOfNormalVector(ShiftedLeftStackOfNormalVector *ShiftedLeftStackOfNormalVector)
	DeleteORMStackGrowthCurveEndArcShape(StackGrowthCurveEndArcShape *StackGrowthCurveEndArcShape)
	DeleteORMStackGrowthCurveStartArcShape(StackGrowthCurveStartArcShape *StackGrowthCurveStartArcShape)
	DeleteORMStackOfGrowthCurve(StackOfGrowthCurve *StackOfGrowthCurve)
	DeleteORMStartArcShape(StartArcShape *StartArcShape)
	DeleteORMStartArcShapeGrid(StartArcShapeGrid *StartArcShapeGrid)
	DeleteORMTopEndArcShape(TopEndArcShape *TopEndArcShape)
	DeleteORMTopEndArcShapeGrid(TopEndArcShapeGrid *TopEndArcShapeGrid)
	DeleteORMTopGrowthCurve2D(TopGrowthCurve2D *TopGrowthCurve2D)
	DeleteORMTopStackGrowthCurveEndArcShape(TopStackGrowthCurveEndArcShape *TopStackGrowthCurveEndArcShape)
	DeleteORMTopStackGrowthCurveStartArcShape(TopStackGrowthCurveStartArcShape *TopStackGrowthCurveStartArcShape)
	DeleteORMTopStackOfGrowthCurve(TopStackOfGrowthCurve *TopStackOfGrowthCurve)
	DeleteORMTopStartArcShape(TopStartArcShape *TopStartArcShape)
	DeleteORMTopStartArcShapeGrid(TopStartArcShapeGrid *TopStartArcShapeGrid)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.ArcNormalVectorShapes = make(map[*ArcNormalVectorShape]struct{})
	stage.ArcNormalVectorShapes_mapString = make(map[string]*ArcNormalVectorShape)
	stage.ArcNormalVectorShape_stagedOrder = make(map[*ArcNormalVectorShape]uint)
	stage.ArcNormalVectorShapeOrder = 0

	stage.ArcNormalVectorShapeGrids = make(map[*ArcNormalVectorShapeGrid]struct{})
	stage.ArcNormalVectorShapeGrids_mapString = make(map[string]*ArcNormalVectorShapeGrid)
	stage.ArcNormalVectorShapeGrid_stagedOrder = make(map[*ArcNormalVectorShapeGrid]uint)
	stage.ArcNormalVectorShapeGridOrder = 0

	stage.AxesShapes = make(map[*AxesShape]struct{})
	stage.AxesShapes_mapString = make(map[string]*AxesShape)
	stage.AxesShape_stagedOrder = make(map[*AxesShape]uint)
	stage.AxesShapeOrder = 0

	stage.BaseVectorShapes = make(map[*BaseVectorShape]struct{})
	stage.BaseVectorShapes_mapString = make(map[string]*BaseVectorShape)
	stage.BaseVectorShape_stagedOrder = make(map[*BaseVectorShape]uint)
	stage.BaseVectorShapeOrder = 0

	stage.BaseVectorShapeGrids = make(map[*BaseVectorShapeGrid]struct{})
	stage.BaseVectorShapeGrids_mapString = make(map[string]*BaseVectorShapeGrid)
	stage.BaseVectorShapeGrid_stagedOrder = make(map[*BaseVectorShapeGrid]uint)
	stage.BaseVectorShapeGridOrder = 0

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

	stage.GrowthCurve2Ds = make(map[*GrowthCurve2D]struct{})
	stage.GrowthCurve2Ds_mapString = make(map[string]*GrowthCurve2D)
	stage.GrowthCurve2D_stagedOrder = make(map[*GrowthCurve2D]uint)
	stage.GrowthCurve2DOrder = 0

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

	stage.ShiftedLeftStackGrowthCurveEndArcShapes = make(map[*ShiftedLeftStackGrowthCurveEndArcShape]struct{})
	stage.ShiftedLeftStackGrowthCurveEndArcShapes_mapString = make(map[string]*ShiftedLeftStackGrowthCurveEndArcShape)
	stage.ShiftedLeftStackGrowthCurveEndArcShape_stagedOrder = make(map[*ShiftedLeftStackGrowthCurveEndArcShape]uint)
	stage.ShiftedLeftStackGrowthCurveEndArcShapeOrder = 0

	stage.ShiftedLeftStackGrowthCurveStartArcShapes = make(map[*ShiftedLeftStackGrowthCurveStartArcShape]struct{})
	stage.ShiftedLeftStackGrowthCurveStartArcShapes_mapString = make(map[string]*ShiftedLeftStackGrowthCurveStartArcShape)
	stage.ShiftedLeftStackGrowthCurveStartArcShape_stagedOrder = make(map[*ShiftedLeftStackGrowthCurveStartArcShape]uint)
	stage.ShiftedLeftStackGrowthCurveStartArcShapeOrder = 0

	stage.ShiftedLeftStackNormalVectors = make(map[*ShiftedLeftStackNormalVector]struct{})
	stage.ShiftedLeftStackNormalVectors_mapString = make(map[string]*ShiftedLeftStackNormalVector)
	stage.ShiftedLeftStackNormalVector_stagedOrder = make(map[*ShiftedLeftStackNormalVector]uint)
	stage.ShiftedLeftStackNormalVectorOrder = 0

	stage.ShiftedLeftStackOfGrowthCurves = make(map[*ShiftedLeftStackOfGrowthCurve]struct{})
	stage.ShiftedLeftStackOfGrowthCurves_mapString = make(map[string]*ShiftedLeftStackOfGrowthCurve)
	stage.ShiftedLeftStackOfGrowthCurve_stagedOrder = make(map[*ShiftedLeftStackOfGrowthCurve]uint)
	stage.ShiftedLeftStackOfGrowthCurveOrder = 0

	stage.ShiftedLeftStackOfNormalVectors = make(map[*ShiftedLeftStackOfNormalVector]struct{})
	stage.ShiftedLeftStackOfNormalVectors_mapString = make(map[string]*ShiftedLeftStackOfNormalVector)
	stage.ShiftedLeftStackOfNormalVector_stagedOrder = make(map[*ShiftedLeftStackOfNormalVector]uint)
	stage.ShiftedLeftStackOfNormalVectorOrder = 0

	stage.StackGrowthCurveEndArcShapes = make(map[*StackGrowthCurveEndArcShape]struct{})
	stage.StackGrowthCurveEndArcShapes_mapString = make(map[string]*StackGrowthCurveEndArcShape)
	stage.StackGrowthCurveEndArcShape_stagedOrder = make(map[*StackGrowthCurveEndArcShape]uint)
	stage.StackGrowthCurveEndArcShapeOrder = 0

	stage.StackGrowthCurveStartArcShapes = make(map[*StackGrowthCurveStartArcShape]struct{})
	stage.StackGrowthCurveStartArcShapes_mapString = make(map[string]*StackGrowthCurveStartArcShape)
	stage.StackGrowthCurveStartArcShape_stagedOrder = make(map[*StackGrowthCurveStartArcShape]uint)
	stage.StackGrowthCurveStartArcShapeOrder = 0

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

	stage.TopEndArcShapes = make(map[*TopEndArcShape]struct{})
	stage.TopEndArcShapes_mapString = make(map[string]*TopEndArcShape)
	stage.TopEndArcShape_stagedOrder = make(map[*TopEndArcShape]uint)
	stage.TopEndArcShapeOrder = 0

	stage.TopEndArcShapeGrids = make(map[*TopEndArcShapeGrid]struct{})
	stage.TopEndArcShapeGrids_mapString = make(map[string]*TopEndArcShapeGrid)
	stage.TopEndArcShapeGrid_stagedOrder = make(map[*TopEndArcShapeGrid]uint)
	stage.TopEndArcShapeGridOrder = 0

	stage.TopGrowthCurve2Ds = make(map[*TopGrowthCurve2D]struct{})
	stage.TopGrowthCurve2Ds_mapString = make(map[string]*TopGrowthCurve2D)
	stage.TopGrowthCurve2D_stagedOrder = make(map[*TopGrowthCurve2D]uint)
	stage.TopGrowthCurve2DOrder = 0

	stage.TopStackGrowthCurveEndArcShapes = make(map[*TopStackGrowthCurveEndArcShape]struct{})
	stage.TopStackGrowthCurveEndArcShapes_mapString = make(map[string]*TopStackGrowthCurveEndArcShape)
	stage.TopStackGrowthCurveEndArcShape_stagedOrder = make(map[*TopStackGrowthCurveEndArcShape]uint)
	stage.TopStackGrowthCurveEndArcShapeOrder = 0

	stage.TopStackGrowthCurveStartArcShapes = make(map[*TopStackGrowthCurveStartArcShape]struct{})
	stage.TopStackGrowthCurveStartArcShapes_mapString = make(map[string]*TopStackGrowthCurveStartArcShape)
	stage.TopStackGrowthCurveStartArcShape_stagedOrder = make(map[*TopStackGrowthCurveStartArcShape]uint)
	stage.TopStackGrowthCurveStartArcShapeOrder = 0

	stage.TopStackOfGrowthCurves = make(map[*TopStackOfGrowthCurve]struct{})
	stage.TopStackOfGrowthCurves_mapString = make(map[string]*TopStackOfGrowthCurve)
	stage.TopStackOfGrowthCurve_stagedOrder = make(map[*TopStackOfGrowthCurve]uint)
	stage.TopStackOfGrowthCurveOrder = 0

	stage.TopStartArcShapes = make(map[*TopStartArcShape]struct{})
	stage.TopStartArcShapes_mapString = make(map[string]*TopStartArcShape)
	stage.TopStartArcShape_stagedOrder = make(map[*TopStartArcShape]uint)
	stage.TopStartArcShapeOrder = 0

	stage.TopStartArcShapeGrids = make(map[*TopStartArcShapeGrid]struct{})
	stage.TopStartArcShapeGrids_mapString = make(map[string]*TopStartArcShapeGrid)
	stage.TopStartArcShapeGrid_stagedOrder = make(map[*TopStartArcShapeGrid]uint)
	stage.TopStartArcShapeGridOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.ArcNormalVectorShapes = nil
	stage.ArcNormalVectorShapes_mapString = nil

	stage.ArcNormalVectorShapeGrids = nil
	stage.ArcNormalVectorShapeGrids_mapString = nil

	stage.AxesShapes = nil
	stage.AxesShapes_mapString = nil

	stage.BaseVectorShapes = nil
	stage.BaseVectorShapes_mapString = nil

	stage.BaseVectorShapeGrids = nil
	stage.BaseVectorShapeGrids_mapString = nil

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

	stage.GrowthCurve2Ds = nil
	stage.GrowthCurve2Ds_mapString = nil

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

	stage.ShiftedLeftStackGrowthCurveEndArcShapes = nil
	stage.ShiftedLeftStackGrowthCurveEndArcShapes_mapString = nil

	stage.ShiftedLeftStackGrowthCurveStartArcShapes = nil
	stage.ShiftedLeftStackGrowthCurveStartArcShapes_mapString = nil

	stage.ShiftedLeftStackNormalVectors = nil
	stage.ShiftedLeftStackNormalVectors_mapString = nil

	stage.ShiftedLeftStackOfGrowthCurves = nil
	stage.ShiftedLeftStackOfGrowthCurves_mapString = nil

	stage.ShiftedLeftStackOfNormalVectors = nil
	stage.ShiftedLeftStackOfNormalVectors_mapString = nil

	stage.StackGrowthCurveEndArcShapes = nil
	stage.StackGrowthCurveEndArcShapes_mapString = nil

	stage.StackGrowthCurveStartArcShapes = nil
	stage.StackGrowthCurveStartArcShapes_mapString = nil

	stage.StackOfGrowthCurves = nil
	stage.StackOfGrowthCurves_mapString = nil

	stage.StartArcShapes = nil
	stage.StartArcShapes_mapString = nil

	stage.StartArcShapeGrids = nil
	stage.StartArcShapeGrids_mapString = nil

	stage.TopEndArcShapes = nil
	stage.TopEndArcShapes_mapString = nil

	stage.TopEndArcShapeGrids = nil
	stage.TopEndArcShapeGrids_mapString = nil

	stage.TopGrowthCurve2Ds = nil
	stage.TopGrowthCurve2Ds_mapString = nil

	stage.TopStackGrowthCurveEndArcShapes = nil
	stage.TopStackGrowthCurveEndArcShapes_mapString = nil

	stage.TopStackGrowthCurveStartArcShapes = nil
	stage.TopStackGrowthCurveStartArcShapes_mapString = nil

	stage.TopStackOfGrowthCurves = nil
	stage.TopStackOfGrowthCurves_mapString = nil

	stage.TopStartArcShapes = nil
	stage.TopStartArcShapes_mapString = nil

	stage.TopStartArcShapeGrids = nil
	stage.TopStartArcShapeGrids_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for arcnormalvectorshape := range stage.ArcNormalVectorShapes {
		arcnormalvectorshape.Unstage(stage)
	}

	for arcnormalvectorshapegrid := range stage.ArcNormalVectorShapeGrids {
		arcnormalvectorshapegrid.Unstage(stage)
	}

	for axesshape := range stage.AxesShapes {
		axesshape.Unstage(stage)
	}

	for basevectorshape := range stage.BaseVectorShapes {
		basevectorshape.Unstage(stage)
	}

	for basevectorshapegrid := range stage.BaseVectorShapeGrids {
		basevectorshapegrid.Unstage(stage)
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

	for growthcurve2d := range stage.GrowthCurve2Ds {
		growthcurve2d.Unstage(stage)
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

	for shiftedleftstackgrowthcurveendarcshape := range stage.ShiftedLeftStackGrowthCurveEndArcShapes {
		shiftedleftstackgrowthcurveendarcshape.Unstage(stage)
	}

	for shiftedleftstackgrowthcurvestartarcshape := range stage.ShiftedLeftStackGrowthCurveStartArcShapes {
		shiftedleftstackgrowthcurvestartarcshape.Unstage(stage)
	}

	for shiftedleftstacknormalvector := range stage.ShiftedLeftStackNormalVectors {
		shiftedleftstacknormalvector.Unstage(stage)
	}

	for shiftedleftstackofgrowthcurve := range stage.ShiftedLeftStackOfGrowthCurves {
		shiftedleftstackofgrowthcurve.Unstage(stage)
	}

	for shiftedleftstackofnormalvector := range stage.ShiftedLeftStackOfNormalVectors {
		shiftedleftstackofnormalvector.Unstage(stage)
	}

	for stackgrowthcurveendarcshape := range stage.StackGrowthCurveEndArcShapes {
		stackgrowthcurveendarcshape.Unstage(stage)
	}

	for stackgrowthcurvestartarcshape := range stage.StackGrowthCurveStartArcShapes {
		stackgrowthcurvestartarcshape.Unstage(stage)
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

	for topendarcshape := range stage.TopEndArcShapes {
		topendarcshape.Unstage(stage)
	}

	for topendarcshapegrid := range stage.TopEndArcShapeGrids {
		topendarcshapegrid.Unstage(stage)
	}

	for topgrowthcurve2d := range stage.TopGrowthCurve2Ds {
		topgrowthcurve2d.Unstage(stage)
	}

	for topstackgrowthcurveendarcshape := range stage.TopStackGrowthCurveEndArcShapes {
		topstackgrowthcurveendarcshape.Unstage(stage)
	}

	for topstackgrowthcurvestartarcshape := range stage.TopStackGrowthCurveStartArcShapes {
		topstackgrowthcurvestartarcshape.Unstage(stage)
	}

	for topstackofgrowthcurve := range stage.TopStackOfGrowthCurves {
		topstackofgrowthcurve.Unstage(stage)
	}

	for topstartarcshape := range stage.TopStartArcShapes {
		topstartarcshape.Unstage(stage)
	}

	for topstartarcshapegrid := range stage.TopStartArcShapeGrids {
		topstartarcshapegrid.Unstage(stage)
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
	case map[*ArcNormalVectorShape]any:
		return any(&stage.ArcNormalVectorShapes).(*Type)
	case map[*ArcNormalVectorShapeGrid]any:
		return any(&stage.ArcNormalVectorShapeGrids).(*Type)
	case map[*AxesShape]any:
		return any(&stage.AxesShapes).(*Type)
	case map[*BaseVectorShape]any:
		return any(&stage.BaseVectorShapes).(*Type)
	case map[*BaseVectorShapeGrid]any:
		return any(&stage.BaseVectorShapeGrids).(*Type)
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
	case map[*GrowthCurve2D]any:
		return any(&stage.GrowthCurve2Ds).(*Type)
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
	case map[*ShiftedLeftStackGrowthCurveEndArcShape]any:
		return any(&stage.ShiftedLeftStackGrowthCurveEndArcShapes).(*Type)
	case map[*ShiftedLeftStackGrowthCurveStartArcShape]any:
		return any(&stage.ShiftedLeftStackGrowthCurveStartArcShapes).(*Type)
	case map[*ShiftedLeftStackNormalVector]any:
		return any(&stage.ShiftedLeftStackNormalVectors).(*Type)
	case map[*ShiftedLeftStackOfGrowthCurve]any:
		return any(&stage.ShiftedLeftStackOfGrowthCurves).(*Type)
	case map[*ShiftedLeftStackOfNormalVector]any:
		return any(&stage.ShiftedLeftStackOfNormalVectors).(*Type)
	case map[*StackGrowthCurveEndArcShape]any:
		return any(&stage.StackGrowthCurveEndArcShapes).(*Type)
	case map[*StackGrowthCurveStartArcShape]any:
		return any(&stage.StackGrowthCurveStartArcShapes).(*Type)
	case map[*StackOfGrowthCurve]any:
		return any(&stage.StackOfGrowthCurves).(*Type)
	case map[*StartArcShape]any:
		return any(&stage.StartArcShapes).(*Type)
	case map[*StartArcShapeGrid]any:
		return any(&stage.StartArcShapeGrids).(*Type)
	case map[*TopEndArcShape]any:
		return any(&stage.TopEndArcShapes).(*Type)
	case map[*TopEndArcShapeGrid]any:
		return any(&stage.TopEndArcShapeGrids).(*Type)
	case map[*TopGrowthCurve2D]any:
		return any(&stage.TopGrowthCurve2Ds).(*Type)
	case map[*TopStackGrowthCurveEndArcShape]any:
		return any(&stage.TopStackGrowthCurveEndArcShapes).(*Type)
	case map[*TopStackGrowthCurveStartArcShape]any:
		return any(&stage.TopStackGrowthCurveStartArcShapes).(*Type)
	case map[*TopStackOfGrowthCurve]any:
		return any(&stage.TopStackOfGrowthCurves).(*Type)
	case map[*TopStartArcShape]any:
		return any(&stage.TopStartArcShapes).(*Type)
	case map[*TopStartArcShapeGrid]any:
		return any(&stage.TopStartArcShapeGrids).(*Type)
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
	case *ArcNormalVectorShape:
		return any(stage.ArcNormalVectorShapes_mapString).(map[string]Type)
	case *ArcNormalVectorShapeGrid:
		return any(stage.ArcNormalVectorShapeGrids_mapString).(map[string]Type)
	case *AxesShape:
		return any(stage.AxesShapes_mapString).(map[string]Type)
	case *BaseVectorShape:
		return any(stage.BaseVectorShapes_mapString).(map[string]Type)
	case *BaseVectorShapeGrid:
		return any(stage.BaseVectorShapeGrids_mapString).(map[string]Type)
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
	case *GrowthCurve2D:
		return any(stage.GrowthCurve2Ds_mapString).(map[string]Type)
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
	case *ShiftedLeftStackGrowthCurveEndArcShape:
		return any(stage.ShiftedLeftStackGrowthCurveEndArcShapes_mapString).(map[string]Type)
	case *ShiftedLeftStackGrowthCurveStartArcShape:
		return any(stage.ShiftedLeftStackGrowthCurveStartArcShapes_mapString).(map[string]Type)
	case *ShiftedLeftStackNormalVector:
		return any(stage.ShiftedLeftStackNormalVectors_mapString).(map[string]Type)
	case *ShiftedLeftStackOfGrowthCurve:
		return any(stage.ShiftedLeftStackOfGrowthCurves_mapString).(map[string]Type)
	case *ShiftedLeftStackOfNormalVector:
		return any(stage.ShiftedLeftStackOfNormalVectors_mapString).(map[string]Type)
	case *StackGrowthCurveEndArcShape:
		return any(stage.StackGrowthCurveEndArcShapes_mapString).(map[string]Type)
	case *StackGrowthCurveStartArcShape:
		return any(stage.StackGrowthCurveStartArcShapes_mapString).(map[string]Type)
	case *StackOfGrowthCurve:
		return any(stage.StackOfGrowthCurves_mapString).(map[string]Type)
	case *StartArcShape:
		return any(stage.StartArcShapes_mapString).(map[string]Type)
	case *StartArcShapeGrid:
		return any(stage.StartArcShapeGrids_mapString).(map[string]Type)
	case *TopEndArcShape:
		return any(stage.TopEndArcShapes_mapString).(map[string]Type)
	case *TopEndArcShapeGrid:
		return any(stage.TopEndArcShapeGrids_mapString).(map[string]Type)
	case *TopGrowthCurve2D:
		return any(stage.TopGrowthCurve2Ds_mapString).(map[string]Type)
	case *TopStackGrowthCurveEndArcShape:
		return any(stage.TopStackGrowthCurveEndArcShapes_mapString).(map[string]Type)
	case *TopStackGrowthCurveStartArcShape:
		return any(stage.TopStackGrowthCurveStartArcShapes_mapString).(map[string]Type)
	case *TopStackOfGrowthCurve:
		return any(stage.TopStackOfGrowthCurves_mapString).(map[string]Type)
	case *TopStartArcShape:
		return any(stage.TopStartArcShapes_mapString).(map[string]Type)
	case *TopStartArcShapeGrid:
		return any(stage.TopStartArcShapeGrids_mapString).(map[string]Type)
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
	case ArcNormalVectorShape:
		return any(&stage.ArcNormalVectorShapes).(*map[*Type]struct{})
	case ArcNormalVectorShapeGrid:
		return any(&stage.ArcNormalVectorShapeGrids).(*map[*Type]struct{})
	case AxesShape:
		return any(&stage.AxesShapes).(*map[*Type]struct{})
	case BaseVectorShape:
		return any(&stage.BaseVectorShapes).(*map[*Type]struct{})
	case BaseVectorShapeGrid:
		return any(&stage.BaseVectorShapeGrids).(*map[*Type]struct{})
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
	case GrowthCurve2D:
		return any(&stage.GrowthCurve2Ds).(*map[*Type]struct{})
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
	case ShiftedLeftStackGrowthCurveEndArcShape:
		return any(&stage.ShiftedLeftStackGrowthCurveEndArcShapes).(*map[*Type]struct{})
	case ShiftedLeftStackGrowthCurveStartArcShape:
		return any(&stage.ShiftedLeftStackGrowthCurveStartArcShapes).(*map[*Type]struct{})
	case ShiftedLeftStackNormalVector:
		return any(&stage.ShiftedLeftStackNormalVectors).(*map[*Type]struct{})
	case ShiftedLeftStackOfGrowthCurve:
		return any(&stage.ShiftedLeftStackOfGrowthCurves).(*map[*Type]struct{})
	case ShiftedLeftStackOfNormalVector:
		return any(&stage.ShiftedLeftStackOfNormalVectors).(*map[*Type]struct{})
	case StackGrowthCurveEndArcShape:
		return any(&stage.StackGrowthCurveEndArcShapes).(*map[*Type]struct{})
	case StackGrowthCurveStartArcShape:
		return any(&stage.StackGrowthCurveStartArcShapes).(*map[*Type]struct{})
	case StackOfGrowthCurve:
		return any(&stage.StackOfGrowthCurves).(*map[*Type]struct{})
	case StartArcShape:
		return any(&stage.StartArcShapes).(*map[*Type]struct{})
	case StartArcShapeGrid:
		return any(&stage.StartArcShapeGrids).(*map[*Type]struct{})
	case TopEndArcShape:
		return any(&stage.TopEndArcShapes).(*map[*Type]struct{})
	case TopEndArcShapeGrid:
		return any(&stage.TopEndArcShapeGrids).(*map[*Type]struct{})
	case TopGrowthCurve2D:
		return any(&stage.TopGrowthCurve2Ds).(*map[*Type]struct{})
	case TopStackGrowthCurveEndArcShape:
		return any(&stage.TopStackGrowthCurveEndArcShapes).(*map[*Type]struct{})
	case TopStackGrowthCurveStartArcShape:
		return any(&stage.TopStackGrowthCurveStartArcShapes).(*map[*Type]struct{})
	case TopStackOfGrowthCurve:
		return any(&stage.TopStackOfGrowthCurves).(*map[*Type]struct{})
	case TopStartArcShape:
		return any(&stage.TopStartArcShapes).(*map[*Type]struct{})
	case TopStartArcShapeGrid:
		return any(&stage.TopStartArcShapeGrids).(*map[*Type]struct{})
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
	case *ArcNormalVectorShape:
		return any(&stage.ArcNormalVectorShapes).(*map[Type]struct{})
	case *ArcNormalVectorShapeGrid:
		return any(&stage.ArcNormalVectorShapeGrids).(*map[Type]struct{})
	case *AxesShape:
		return any(&stage.AxesShapes).(*map[Type]struct{})
	case *BaseVectorShape:
		return any(&stage.BaseVectorShapes).(*map[Type]struct{})
	case *BaseVectorShapeGrid:
		return any(&stage.BaseVectorShapeGrids).(*map[Type]struct{})
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
	case *GrowthCurve2D:
		return any(&stage.GrowthCurve2Ds).(*map[Type]struct{})
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
	case *ShiftedLeftStackGrowthCurveEndArcShape:
		return any(&stage.ShiftedLeftStackGrowthCurveEndArcShapes).(*map[Type]struct{})
	case *ShiftedLeftStackGrowthCurveStartArcShape:
		return any(&stage.ShiftedLeftStackGrowthCurveStartArcShapes).(*map[Type]struct{})
	case *ShiftedLeftStackNormalVector:
		return any(&stage.ShiftedLeftStackNormalVectors).(*map[Type]struct{})
	case *ShiftedLeftStackOfGrowthCurve:
		return any(&stage.ShiftedLeftStackOfGrowthCurves).(*map[Type]struct{})
	case *ShiftedLeftStackOfNormalVector:
		return any(&stage.ShiftedLeftStackOfNormalVectors).(*map[Type]struct{})
	case *StackGrowthCurveEndArcShape:
		return any(&stage.StackGrowthCurveEndArcShapes).(*map[Type]struct{})
	case *StackGrowthCurveStartArcShape:
		return any(&stage.StackGrowthCurveStartArcShapes).(*map[Type]struct{})
	case *StackOfGrowthCurve:
		return any(&stage.StackOfGrowthCurves).(*map[Type]struct{})
	case *StartArcShape:
		return any(&stage.StartArcShapes).(*map[Type]struct{})
	case *StartArcShapeGrid:
		return any(&stage.StartArcShapeGrids).(*map[Type]struct{})
	case *TopEndArcShape:
		return any(&stage.TopEndArcShapes).(*map[Type]struct{})
	case *TopEndArcShapeGrid:
		return any(&stage.TopEndArcShapeGrids).(*map[Type]struct{})
	case *TopGrowthCurve2D:
		return any(&stage.TopGrowthCurve2Ds).(*map[Type]struct{})
	case *TopStackGrowthCurveEndArcShape:
		return any(&stage.TopStackGrowthCurveEndArcShapes).(*map[Type]struct{})
	case *TopStackGrowthCurveStartArcShape:
		return any(&stage.TopStackGrowthCurveStartArcShapes).(*map[Type]struct{})
	case *TopStackOfGrowthCurve:
		return any(&stage.TopStackOfGrowthCurves).(*map[Type]struct{})
	case *TopStartArcShape:
		return any(&stage.TopStartArcShapes).(*map[Type]struct{})
	case *TopStartArcShapeGrid:
		return any(&stage.TopStartArcShapeGrids).(*map[Type]struct{})
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
	case ArcNormalVectorShape:
		return any(&stage.ArcNormalVectorShapes_mapString).(*map[string]*Type)
	case ArcNormalVectorShapeGrid:
		return any(&stage.ArcNormalVectorShapeGrids_mapString).(*map[string]*Type)
	case AxesShape:
		return any(&stage.AxesShapes_mapString).(*map[string]*Type)
	case BaseVectorShape:
		return any(&stage.BaseVectorShapes_mapString).(*map[string]*Type)
	case BaseVectorShapeGrid:
		return any(&stage.BaseVectorShapeGrids_mapString).(*map[string]*Type)
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
	case GrowthCurve2D:
		return any(&stage.GrowthCurve2Ds_mapString).(*map[string]*Type)
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
	case ShiftedLeftStackGrowthCurveEndArcShape:
		return any(&stage.ShiftedLeftStackGrowthCurveEndArcShapes_mapString).(*map[string]*Type)
	case ShiftedLeftStackGrowthCurveStartArcShape:
		return any(&stage.ShiftedLeftStackGrowthCurveStartArcShapes_mapString).(*map[string]*Type)
	case ShiftedLeftStackNormalVector:
		return any(&stage.ShiftedLeftStackNormalVectors_mapString).(*map[string]*Type)
	case ShiftedLeftStackOfGrowthCurve:
		return any(&stage.ShiftedLeftStackOfGrowthCurves_mapString).(*map[string]*Type)
	case ShiftedLeftStackOfNormalVector:
		return any(&stage.ShiftedLeftStackOfNormalVectors_mapString).(*map[string]*Type)
	case StackGrowthCurveEndArcShape:
		return any(&stage.StackGrowthCurveEndArcShapes_mapString).(*map[string]*Type)
	case StackGrowthCurveStartArcShape:
		return any(&stage.StackGrowthCurveStartArcShapes_mapString).(*map[string]*Type)
	case StackOfGrowthCurve:
		return any(&stage.StackOfGrowthCurves_mapString).(*map[string]*Type)
	case StartArcShape:
		return any(&stage.StartArcShapes_mapString).(*map[string]*Type)
	case StartArcShapeGrid:
		return any(&stage.StartArcShapeGrids_mapString).(*map[string]*Type)
	case TopEndArcShape:
		return any(&stage.TopEndArcShapes_mapString).(*map[string]*Type)
	case TopEndArcShapeGrid:
		return any(&stage.TopEndArcShapeGrids_mapString).(*map[string]*Type)
	case TopGrowthCurve2D:
		return any(&stage.TopGrowthCurve2Ds_mapString).(*map[string]*Type)
	case TopStackGrowthCurveEndArcShape:
		return any(&stage.TopStackGrowthCurveEndArcShapes_mapString).(*map[string]*Type)
	case TopStackGrowthCurveStartArcShape:
		return any(&stage.TopStackGrowthCurveStartArcShapes_mapString).(*map[string]*Type)
	case TopStackOfGrowthCurve:
		return any(&stage.TopStackOfGrowthCurves_mapString).(*map[string]*Type)
	case TopStartArcShape:
		return any(&stage.TopStartArcShapes_mapString).(*map[string]*Type)
	case TopStartArcShapeGrid:
		return any(&stage.TopStartArcShapeGrids_mapString).(*map[string]*Type)
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
	case ArcNormalVectorShape:
		return any(&ArcNormalVectorShape{
			// Initialisation of associations
		}).(*Type)
	case ArcNormalVectorShapeGrid:
		return any(&ArcNormalVectorShapeGrid{
			// Initialisation of associations
			// field is initialized with an instance of ArcNormalVectorShape with the name of the field
			ArcNormalVectorShapes: []*ArcNormalVectorShape{{Name: "ArcNormalVectorShapes"}},
		}).(*Type)
	case AxesShape:
		return any(&AxesShape{
			// Initialisation of associations
		}).(*Type)
	case BaseVectorShape:
		return any(&BaseVectorShape{
			// Initialisation of associations
		}).(*Type)
	case BaseVectorShapeGrid:
		return any(&BaseVectorShapeGrid{
			// Initialisation of associations
			// field is initialized with an instance of BaseVectorShape with the name of the field
			BaseVectorShapes: []*BaseVectorShape{{Name: "BaseVectorShapes"}},
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
	case GrowthCurve2D:
		return any(&GrowthCurve2D{
			// Initialisation of associations
			// field is initialized with an instance of StartArcShapeGrid with the name of the field
			StartArcShapeGrid: &StartArcShapeGrid{Name: "StartArcShapeGrid"},
			// field is initialized with an instance of EndArcShapeGrid with the name of the field
			EndArcShapeGrid: &EndArcShapeGrid{Name: "EndArcShapeGrid"},
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
			// field is initialized with an instance of BaseVectorShapeGrid with the name of the field
			BaseVectorShapeGrid: &BaseVectorShapeGrid{Name: "BaseVectorShapeGrid"},
			// field is initialized with an instance of ArcNormalVectorShapeGrid with the name of the field
			ArcNormalVectorShapeGrid: &ArcNormalVectorShapeGrid{Name: "ArcNormalVectorShapeGrid"},
			// field is initialized with an instance of StartArcShapeGrid with the name of the field
			StartArcShapeGrid: &StartArcShapeGrid{Name: "StartArcShapeGrid"},
			// field is initialized with an instance of TopStartArcShapeGrid with the name of the field
			TopStartArcShapeGrid: &TopStartArcShapeGrid{Name: "TopStartArcShapeGrid"},
			// field is initialized with an instance of EndArcShapeGrid with the name of the field
			EndArcShapeGrid: &EndArcShapeGrid{Name: "EndArcShapeGrid"},
			// field is initialized with an instance of TopEndArcShapeGrid with the name of the field
			TopEndArcShapeGrid: &TopEndArcShapeGrid{Name: "TopEndArcShapeGrid"},
			// field is initialized with an instance of GrowthCurveBezierShapeGrid with the name of the field
			GrowthCurveBezierShapeGrid: &GrowthCurveBezierShapeGrid{Name: "GrowthCurveBezierShapeGrid"},
			// field is initialized with an instance of StackOfGrowthCurve with the name of the field
			StackOfGrowthCurve: &StackOfGrowthCurve{Name: "StackOfGrowthCurve"},
			// field is initialized with an instance of TopStackOfGrowthCurve with the name of the field
			TopStackOfGrowthCurve: &TopStackOfGrowthCurve{Name: "TopStackOfGrowthCurve"},
			// field is initialized with an instance of ShiftedLeftStackOfGrowthCurve with the name of the field
			ShiftedLeftStackOfGrowthCurve: &ShiftedLeftStackOfGrowthCurve{Name: "ShiftedLeftStackOfGrowthCurve"},
			// field is initialized with an instance of ShiftedLeftStackOfNormalVector with the name of the field
			ShiftedLeftStackOfNormalVector: &ShiftedLeftStackOfNormalVector{Name: "ShiftedLeftStackOfNormalVector"},
			// field is initialized with an instance of GrowthCurve2D with the name of the field
			GrowthCurve2D: &GrowthCurve2D{Name: "GrowthCurve2D"},
			// field is initialized with an instance of TopGrowthCurve2D with the name of the field
			TopGrowthCurve2D: &TopGrowthCurve2D{Name: "TopGrowthCurve2D"},
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
	case ShiftedLeftStackGrowthCurveEndArcShape:
		return any(&ShiftedLeftStackGrowthCurveEndArcShape{
			// Initialisation of associations
		}).(*Type)
	case ShiftedLeftStackGrowthCurveStartArcShape:
		return any(&ShiftedLeftStackGrowthCurveStartArcShape{
			// Initialisation of associations
		}).(*Type)
	case ShiftedLeftStackNormalVector:
		return any(&ShiftedLeftStackNormalVector{
			// Initialisation of associations
		}).(*Type)
	case ShiftedLeftStackOfGrowthCurve:
		return any(&ShiftedLeftStackOfGrowthCurve{
			// Initialisation of associations
			// field is initialized with an instance of ShiftedLeftStackGrowthCurveStartArcShape with the name of the field
			ShiftedLeftStackGrowthCurveStartArcShapes: []*ShiftedLeftStackGrowthCurveStartArcShape{{Name: "ShiftedLeftStackGrowthCurveStartArcShapes"}},
			// field is initialized with an instance of ShiftedLeftStackGrowthCurveEndArcShape with the name of the field
			ShiftedLeftStackGrowthCurveEndArcShapes: []*ShiftedLeftStackGrowthCurveEndArcShape{{Name: "ShiftedLeftStackGrowthCurveEndArcShapes"}},
		}).(*Type)
	case ShiftedLeftStackOfNormalVector:
		return any(&ShiftedLeftStackOfNormalVector{
			// Initialisation of associations
			// field is initialized with an instance of ShiftedLeftStackNormalVector with the name of the field
			ShiftedLeftStackNormalVectors: []*ShiftedLeftStackNormalVector{{Name: "ShiftedLeftStackNormalVectors"}},
		}).(*Type)
	case StackGrowthCurveEndArcShape:
		return any(&StackGrowthCurveEndArcShape{
			// Initialisation of associations
		}).(*Type)
	case StackGrowthCurveStartArcShape:
		return any(&StackGrowthCurveStartArcShape{
			// Initialisation of associations
		}).(*Type)
	case StackOfGrowthCurve:
		return any(&StackOfGrowthCurve{
			// Initialisation of associations
			// field is initialized with an instance of StackGrowthCurveStartArcShape with the name of the field
			StackGrowthCurveStartArcShapes: []*StackGrowthCurveStartArcShape{{Name: "StackGrowthCurveStartArcShapes"}},
			// field is initialized with an instance of StackGrowthCurveEndArcShape with the name of the field
			StackGrowthCurveEndArcShapes: []*StackGrowthCurveEndArcShape{{Name: "StackGrowthCurveEndArcShapes"}},
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
	case TopEndArcShape:
		return any(&TopEndArcShape{
			// Initialisation of associations
		}).(*Type)
	case TopEndArcShapeGrid:
		return any(&TopEndArcShapeGrid{
			// Initialisation of associations
			// field is initialized with an instance of TopEndArcShape with the name of the field
			TopEndArcShapes: []*TopEndArcShape{{Name: "TopEndArcShapes"}},
		}).(*Type)
	case TopGrowthCurve2D:
		return any(&TopGrowthCurve2D{
			// Initialisation of associations
			// field is initialized with an instance of TopStartArcShapeGrid with the name of the field
			TopStartArcShapeGrid: &TopStartArcShapeGrid{Name: "TopStartArcShapeGrid"},
			// field is initialized with an instance of TopEndArcShapeGrid with the name of the field
			TopEndArcShapeGrid: &TopEndArcShapeGrid{Name: "TopEndArcShapeGrid"},
		}).(*Type)
	case TopStackGrowthCurveEndArcShape:
		return any(&TopStackGrowthCurveEndArcShape{
			// Initialisation of associations
		}).(*Type)
	case TopStackGrowthCurveStartArcShape:
		return any(&TopStackGrowthCurveStartArcShape{
			// Initialisation of associations
		}).(*Type)
	case TopStackOfGrowthCurve:
		return any(&TopStackOfGrowthCurve{
			// Initialisation of associations
			// field is initialized with an instance of TopStackGrowthCurveStartArcShape with the name of the field
			TopStackGrowthCurveStartArcShapes: []*TopStackGrowthCurveStartArcShape{{Name: "TopStackGrowthCurveStartArcShapes"}},
			// field is initialized with an instance of TopStackGrowthCurveEndArcShape with the name of the field
			TopStackGrowthCurveEndArcShapes: []*TopStackGrowthCurveEndArcShape{{Name: "TopStackGrowthCurveEndArcShapes"}},
		}).(*Type)
	case TopStartArcShape:
		return any(&TopStartArcShape{
			// Initialisation of associations
		}).(*Type)
	case TopStartArcShapeGrid:
		return any(&TopStartArcShapeGrid{
			// Initialisation of associations
			// field is initialized with an instance of TopStartArcShape with the name of the field
			TopStartArcShapes: []*TopStartArcShape{{Name: "TopStartArcShapes"}},
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
	// reverse maps of direct associations of ArcNormalVectorShape
	case ArcNormalVectorShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ArcNormalVectorShapeGrid
	case ArcNormalVectorShapeGrid:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AxesShape
	case AxesShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BaseVectorShape
	case BaseVectorShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BaseVectorShapeGrid
	case BaseVectorShapeGrid:
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
	// reverse maps of direct associations of GrowthCurve2D
	case GrowthCurve2D:
		switch fieldname {
		// insertion point for per direct association field
		case "StartArcShapeGrid":
			res := make(map[*StartArcShapeGrid][]*GrowthCurve2D)
			for growthcurve2d := range stage.GrowthCurve2Ds {
				if growthcurve2d.StartArcShapeGrid != nil {
					startarcshapegrid_ := growthcurve2d.StartArcShapeGrid
					var growthcurve2ds []*GrowthCurve2D
					_, ok := res[startarcshapegrid_]
					if ok {
						growthcurve2ds = res[startarcshapegrid_]
					} else {
						growthcurve2ds = make([]*GrowthCurve2D, 0)
					}
					growthcurve2ds = append(growthcurve2ds, growthcurve2d)
					res[startarcshapegrid_] = growthcurve2ds
				}
			}
			return any(res).(map[*End][]*Start)
		case "EndArcShapeGrid":
			res := make(map[*EndArcShapeGrid][]*GrowthCurve2D)
			for growthcurve2d := range stage.GrowthCurve2Ds {
				if growthcurve2d.EndArcShapeGrid != nil {
					endarcshapegrid_ := growthcurve2d.EndArcShapeGrid
					var growthcurve2ds []*GrowthCurve2D
					_, ok := res[endarcshapegrid_]
					if ok {
						growthcurve2ds = res[endarcshapegrid_]
					} else {
						growthcurve2ds = make([]*GrowthCurve2D, 0)
					}
					growthcurve2ds = append(growthcurve2ds, growthcurve2d)
					res[endarcshapegrid_] = growthcurve2ds
				}
			}
			return any(res).(map[*End][]*Start)
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
		case "BaseVectorShapeGrid":
			res := make(map[*BaseVectorShapeGrid][]*Plant)
			for plant := range stage.Plants {
				if plant.BaseVectorShapeGrid != nil {
					basevectorshapegrid_ := plant.BaseVectorShapeGrid
					var plants []*Plant
					_, ok := res[basevectorshapegrid_]
					if ok {
						plants = res[basevectorshapegrid_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[basevectorshapegrid_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "ArcNormalVectorShapeGrid":
			res := make(map[*ArcNormalVectorShapeGrid][]*Plant)
			for plant := range stage.Plants {
				if plant.ArcNormalVectorShapeGrid != nil {
					arcnormalvectorshapegrid_ := plant.ArcNormalVectorShapeGrid
					var plants []*Plant
					_, ok := res[arcnormalvectorshapegrid_]
					if ok {
						plants = res[arcnormalvectorshapegrid_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[arcnormalvectorshapegrid_] = plants
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
		case "TopStartArcShapeGrid":
			res := make(map[*TopStartArcShapeGrid][]*Plant)
			for plant := range stage.Plants {
				if plant.TopStartArcShapeGrid != nil {
					topstartarcshapegrid_ := plant.TopStartArcShapeGrid
					var plants []*Plant
					_, ok := res[topstartarcshapegrid_]
					if ok {
						plants = res[topstartarcshapegrid_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[topstartarcshapegrid_] = plants
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
		case "TopEndArcShapeGrid":
			res := make(map[*TopEndArcShapeGrid][]*Plant)
			for plant := range stage.Plants {
				if plant.TopEndArcShapeGrid != nil {
					topendarcshapegrid_ := plant.TopEndArcShapeGrid
					var plants []*Plant
					_, ok := res[topendarcshapegrid_]
					if ok {
						plants = res[topendarcshapegrid_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[topendarcshapegrid_] = plants
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
		case "TopStackOfGrowthCurve":
			res := make(map[*TopStackOfGrowthCurve][]*Plant)
			for plant := range stage.Plants {
				if plant.TopStackOfGrowthCurve != nil {
					topstackofgrowthcurve_ := plant.TopStackOfGrowthCurve
					var plants []*Plant
					_, ok := res[topstackofgrowthcurve_]
					if ok {
						plants = res[topstackofgrowthcurve_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[topstackofgrowthcurve_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "ShiftedLeftStackOfGrowthCurve":
			res := make(map[*ShiftedLeftStackOfGrowthCurve][]*Plant)
			for plant := range stage.Plants {
				if plant.ShiftedLeftStackOfGrowthCurve != nil {
					shiftedleftstackofgrowthcurve_ := plant.ShiftedLeftStackOfGrowthCurve
					var plants []*Plant
					_, ok := res[shiftedleftstackofgrowthcurve_]
					if ok {
						plants = res[shiftedleftstackofgrowthcurve_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[shiftedleftstackofgrowthcurve_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "ShiftedLeftStackOfNormalVector":
			res := make(map[*ShiftedLeftStackOfNormalVector][]*Plant)
			for plant := range stage.Plants {
				if plant.ShiftedLeftStackOfNormalVector != nil {
					shiftedleftstackofnormalvector_ := plant.ShiftedLeftStackOfNormalVector
					var plants []*Plant
					_, ok := res[shiftedleftstackofnormalvector_]
					if ok {
						plants = res[shiftedleftstackofnormalvector_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[shiftedleftstackofnormalvector_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "GrowthCurve2D":
			res := make(map[*GrowthCurve2D][]*Plant)
			for plant := range stage.Plants {
				if plant.GrowthCurve2D != nil {
					growthcurve2d_ := plant.GrowthCurve2D
					var plants []*Plant
					_, ok := res[growthcurve2d_]
					if ok {
						plants = res[growthcurve2d_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[growthcurve2d_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "TopGrowthCurve2D":
			res := make(map[*TopGrowthCurve2D][]*Plant)
			for plant := range stage.Plants {
				if plant.TopGrowthCurve2D != nil {
					topgrowthcurve2d_ := plant.TopGrowthCurve2D
					var plants []*Plant
					_, ok := res[topgrowthcurve2d_]
					if ok {
						plants = res[topgrowthcurve2d_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[topgrowthcurve2d_] = plants
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
	// reverse maps of direct associations of ShiftedLeftStackGrowthCurveEndArcShape
	case ShiftedLeftStackGrowthCurveEndArcShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ShiftedLeftStackGrowthCurveStartArcShape
	case ShiftedLeftStackGrowthCurveStartArcShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ShiftedLeftStackNormalVector
	case ShiftedLeftStackNormalVector:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ShiftedLeftStackOfGrowthCurve
	case ShiftedLeftStackOfGrowthCurve:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ShiftedLeftStackOfNormalVector
	case ShiftedLeftStackOfNormalVector:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StackGrowthCurveEndArcShape
	case StackGrowthCurveEndArcShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StackGrowthCurveStartArcShape
	case StackGrowthCurveStartArcShape:
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
	// reverse maps of direct associations of TopEndArcShape
	case TopEndArcShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopEndArcShapeGrid
	case TopEndArcShapeGrid:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopGrowthCurve2D
	case TopGrowthCurve2D:
		switch fieldname {
		// insertion point for per direct association field
		case "TopStartArcShapeGrid":
			res := make(map[*TopStartArcShapeGrid][]*TopGrowthCurve2D)
			for topgrowthcurve2d := range stage.TopGrowthCurve2Ds {
				if topgrowthcurve2d.TopStartArcShapeGrid != nil {
					topstartarcshapegrid_ := topgrowthcurve2d.TopStartArcShapeGrid
					var topgrowthcurve2ds []*TopGrowthCurve2D
					_, ok := res[topstartarcshapegrid_]
					if ok {
						topgrowthcurve2ds = res[topstartarcshapegrid_]
					} else {
						topgrowthcurve2ds = make([]*TopGrowthCurve2D, 0)
					}
					topgrowthcurve2ds = append(topgrowthcurve2ds, topgrowthcurve2d)
					res[topstartarcshapegrid_] = topgrowthcurve2ds
				}
			}
			return any(res).(map[*End][]*Start)
		case "TopEndArcShapeGrid":
			res := make(map[*TopEndArcShapeGrid][]*TopGrowthCurve2D)
			for topgrowthcurve2d := range stage.TopGrowthCurve2Ds {
				if topgrowthcurve2d.TopEndArcShapeGrid != nil {
					topendarcshapegrid_ := topgrowthcurve2d.TopEndArcShapeGrid
					var topgrowthcurve2ds []*TopGrowthCurve2D
					_, ok := res[topendarcshapegrid_]
					if ok {
						topgrowthcurve2ds = res[topendarcshapegrid_]
					} else {
						topgrowthcurve2ds = make([]*TopGrowthCurve2D, 0)
					}
					topgrowthcurve2ds = append(topgrowthcurve2ds, topgrowthcurve2d)
					res[topendarcshapegrid_] = topgrowthcurve2ds
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TopStackGrowthCurveEndArcShape
	case TopStackGrowthCurveEndArcShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopStackGrowthCurveStartArcShape
	case TopStackGrowthCurveStartArcShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopStackOfGrowthCurve
	case TopStackOfGrowthCurve:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopStartArcShape
	case TopStartArcShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopStartArcShapeGrid
	case TopStartArcShapeGrid:
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
	// reverse maps of direct associations of ArcNormalVectorShape
	case ArcNormalVectorShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ArcNormalVectorShapeGrid
	case ArcNormalVectorShapeGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "ArcNormalVectorShapes":
			res := make(map[*ArcNormalVectorShape][]*ArcNormalVectorShapeGrid)
			for arcnormalvectorshapegrid := range stage.ArcNormalVectorShapeGrids {
				for _, arcnormalvectorshape_ := range arcnormalvectorshapegrid.ArcNormalVectorShapes {
					res[arcnormalvectorshape_] = append(res[arcnormalvectorshape_], arcnormalvectorshapegrid)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AxesShape
	case AxesShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BaseVectorShape
	case BaseVectorShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BaseVectorShapeGrid
	case BaseVectorShapeGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "BaseVectorShapes":
			res := make(map[*BaseVectorShape][]*BaseVectorShapeGrid)
			for basevectorshapegrid := range stage.BaseVectorShapeGrids {
				for _, basevectorshape_ := range basevectorshapegrid.BaseVectorShapes {
					res[basevectorshape_] = append(res[basevectorshape_], basevectorshapegrid)
				}
			}
			return any(res).(map[*End][]*Start)
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
	// reverse maps of direct associations of GrowthCurve2D
	case GrowthCurve2D:
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
	// reverse maps of direct associations of ShiftedLeftStackGrowthCurveEndArcShape
	case ShiftedLeftStackGrowthCurveEndArcShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ShiftedLeftStackGrowthCurveStartArcShape
	case ShiftedLeftStackGrowthCurveStartArcShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ShiftedLeftStackNormalVector
	case ShiftedLeftStackNormalVector:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ShiftedLeftStackOfGrowthCurve
	case ShiftedLeftStackOfGrowthCurve:
		switch fieldname {
		// insertion point for per direct association field
		case "ShiftedLeftStackGrowthCurveStartArcShapes":
			res := make(map[*ShiftedLeftStackGrowthCurveStartArcShape][]*ShiftedLeftStackOfGrowthCurve)
			for shiftedleftstackofgrowthcurve := range stage.ShiftedLeftStackOfGrowthCurves {
				for _, shiftedleftstackgrowthcurvestartarcshape_ := range shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes {
					res[shiftedleftstackgrowthcurvestartarcshape_] = append(res[shiftedleftstackgrowthcurvestartarcshape_], shiftedleftstackofgrowthcurve)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ShiftedLeftStackGrowthCurveEndArcShapes":
			res := make(map[*ShiftedLeftStackGrowthCurveEndArcShape][]*ShiftedLeftStackOfGrowthCurve)
			for shiftedleftstackofgrowthcurve := range stage.ShiftedLeftStackOfGrowthCurves {
				for _, shiftedleftstackgrowthcurveendarcshape_ := range shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes {
					res[shiftedleftstackgrowthcurveendarcshape_] = append(res[shiftedleftstackgrowthcurveendarcshape_], shiftedleftstackofgrowthcurve)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ShiftedLeftStackOfNormalVector
	case ShiftedLeftStackOfNormalVector:
		switch fieldname {
		// insertion point for per direct association field
		case "ShiftedLeftStackNormalVectors":
			res := make(map[*ShiftedLeftStackNormalVector][]*ShiftedLeftStackOfNormalVector)
			for shiftedleftstackofnormalvector := range stage.ShiftedLeftStackOfNormalVectors {
				for _, shiftedleftstacknormalvector_ := range shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors {
					res[shiftedleftstacknormalvector_] = append(res[shiftedleftstacknormalvector_], shiftedleftstackofnormalvector)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of StackGrowthCurveEndArcShape
	case StackGrowthCurveEndArcShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StackGrowthCurveStartArcShape
	case StackGrowthCurveStartArcShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StackOfGrowthCurve
	case StackOfGrowthCurve:
		switch fieldname {
		// insertion point for per direct association field
		case "StackGrowthCurveStartArcShapes":
			res := make(map[*StackGrowthCurveStartArcShape][]*StackOfGrowthCurve)
			for stackofgrowthcurve := range stage.StackOfGrowthCurves {
				for _, stackgrowthcurvestartarcshape_ := range stackofgrowthcurve.StackGrowthCurveStartArcShapes {
					res[stackgrowthcurvestartarcshape_] = append(res[stackgrowthcurvestartarcshape_], stackofgrowthcurve)
				}
			}
			return any(res).(map[*End][]*Start)
		case "StackGrowthCurveEndArcShapes":
			res := make(map[*StackGrowthCurveEndArcShape][]*StackOfGrowthCurve)
			for stackofgrowthcurve := range stage.StackOfGrowthCurves {
				for _, stackgrowthcurveendarcshape_ := range stackofgrowthcurve.StackGrowthCurveEndArcShapes {
					res[stackgrowthcurveendarcshape_] = append(res[stackgrowthcurveendarcshape_], stackofgrowthcurve)
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
	// reverse maps of direct associations of TopEndArcShape
	case TopEndArcShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopEndArcShapeGrid
	case TopEndArcShapeGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "TopEndArcShapes":
			res := make(map[*TopEndArcShape][]*TopEndArcShapeGrid)
			for topendarcshapegrid := range stage.TopEndArcShapeGrids {
				for _, topendarcshape_ := range topendarcshapegrid.TopEndArcShapes {
					res[topendarcshape_] = append(res[topendarcshape_], topendarcshapegrid)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TopGrowthCurve2D
	case TopGrowthCurve2D:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopStackGrowthCurveEndArcShape
	case TopStackGrowthCurveEndArcShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopStackGrowthCurveStartArcShape
	case TopStackGrowthCurveStartArcShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopStackOfGrowthCurve
	case TopStackOfGrowthCurve:
		switch fieldname {
		// insertion point for per direct association field
		case "TopStackGrowthCurveStartArcShapes":
			res := make(map[*TopStackGrowthCurveStartArcShape][]*TopStackOfGrowthCurve)
			for topstackofgrowthcurve := range stage.TopStackOfGrowthCurves {
				for _, topstackgrowthcurvestartarcshape_ := range topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes {
					res[topstackgrowthcurvestartarcshape_] = append(res[topstackgrowthcurvestartarcshape_], topstackofgrowthcurve)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TopStackGrowthCurveEndArcShapes":
			res := make(map[*TopStackGrowthCurveEndArcShape][]*TopStackOfGrowthCurve)
			for topstackofgrowthcurve := range stage.TopStackOfGrowthCurves {
				for _, topstackgrowthcurveendarcshape_ := range topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes {
					res[topstackgrowthcurveendarcshape_] = append(res[topstackgrowthcurveendarcshape_], topstackofgrowthcurve)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TopStartArcShape
	case TopStartArcShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopStartArcShapeGrid
	case TopStartArcShapeGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "TopStartArcShapes":
			res := make(map[*TopStartArcShape][]*TopStartArcShapeGrid)
			for topstartarcshapegrid := range stage.TopStartArcShapeGrids {
				for _, topstartarcshape_ := range topstartarcshapegrid.TopStartArcShapes {
					res[topstartarcshape_] = append(res[topstartarcshape_], topstartarcshapegrid)
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
	case *ArcNormalVectorShape:
		res = "ArcNormalVectorShape"
	case *ArcNormalVectorShapeGrid:
		res = "ArcNormalVectorShapeGrid"
	case *AxesShape:
		res = "AxesShape"
	case *BaseVectorShape:
		res = "BaseVectorShape"
	case *BaseVectorShapeGrid:
		res = "BaseVectorShapeGrid"
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
	case *GrowthCurve2D:
		res = "GrowthCurve2D"
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
	case *ShiftedLeftStackGrowthCurveEndArcShape:
		res = "ShiftedLeftStackGrowthCurveEndArcShape"
	case *ShiftedLeftStackGrowthCurveStartArcShape:
		res = "ShiftedLeftStackGrowthCurveStartArcShape"
	case *ShiftedLeftStackNormalVector:
		res = "ShiftedLeftStackNormalVector"
	case *ShiftedLeftStackOfGrowthCurve:
		res = "ShiftedLeftStackOfGrowthCurve"
	case *ShiftedLeftStackOfNormalVector:
		res = "ShiftedLeftStackOfNormalVector"
	case *StackGrowthCurveEndArcShape:
		res = "StackGrowthCurveEndArcShape"
	case *StackGrowthCurveStartArcShape:
		res = "StackGrowthCurveStartArcShape"
	case *StackOfGrowthCurve:
		res = "StackOfGrowthCurve"
	case *StartArcShape:
		res = "StartArcShape"
	case *StartArcShapeGrid:
		res = "StartArcShapeGrid"
	case *TopEndArcShape:
		res = "TopEndArcShape"
	case *TopEndArcShapeGrid:
		res = "TopEndArcShapeGrid"
	case *TopGrowthCurve2D:
		res = "TopGrowthCurve2D"
	case *TopStackGrowthCurveEndArcShape:
		res = "TopStackGrowthCurveEndArcShape"
	case *TopStackGrowthCurveStartArcShape:
		res = "TopStackGrowthCurveStartArcShape"
	case *TopStackOfGrowthCurve:
		res = "TopStackOfGrowthCurve"
	case *TopStartArcShape:
		res = "TopStartArcShape"
	case *TopStartArcShapeGrid:
		res = "TopStartArcShapeGrid"
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
	case *ArcNormalVectorShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "ArcNormalVectorShapeGrid"
		rf.Fieldname = "ArcNormalVectorShapes"
		res = append(res, rf)
	case *ArcNormalVectorShapeGrid:
		var rf ReverseField
		_ = rf
	case *AxesShape:
		var rf ReverseField
		_ = rf
	case *BaseVectorShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "BaseVectorShapeGrid"
		rf.Fieldname = "BaseVectorShapes"
		res = append(res, rf)
	case *BaseVectorShapeGrid:
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
	case *GrowthCurve2D:
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
	case *ShiftedLeftStackGrowthCurveEndArcShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "ShiftedLeftStackOfGrowthCurve"
		rf.Fieldname = "ShiftedLeftStackGrowthCurveEndArcShapes"
		res = append(res, rf)
	case *ShiftedLeftStackGrowthCurveStartArcShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "ShiftedLeftStackOfGrowthCurve"
		rf.Fieldname = "ShiftedLeftStackGrowthCurveStartArcShapes"
		res = append(res, rf)
	case *ShiftedLeftStackNormalVector:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "ShiftedLeftStackOfNormalVector"
		rf.Fieldname = "ShiftedLeftStackNormalVectors"
		res = append(res, rf)
	case *ShiftedLeftStackOfGrowthCurve:
		var rf ReverseField
		_ = rf
	case *ShiftedLeftStackOfNormalVector:
		var rf ReverseField
		_ = rf
	case *StackGrowthCurveEndArcShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "StackOfGrowthCurve"
		rf.Fieldname = "StackGrowthCurveEndArcShapes"
		res = append(res, rf)
	case *StackGrowthCurveStartArcShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "StackOfGrowthCurve"
		rf.Fieldname = "StackGrowthCurveStartArcShapes"
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
	case *TopEndArcShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "TopEndArcShapeGrid"
		rf.Fieldname = "TopEndArcShapes"
		res = append(res, rf)
	case *TopEndArcShapeGrid:
		var rf ReverseField
		_ = rf
	case *TopGrowthCurve2D:
		var rf ReverseField
		_ = rf
	case *TopStackGrowthCurveEndArcShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "TopStackOfGrowthCurve"
		rf.Fieldname = "TopStackGrowthCurveEndArcShapes"
		res = append(res, rf)
	case *TopStackGrowthCurveStartArcShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "TopStackOfGrowthCurve"
		rf.Fieldname = "TopStackGrowthCurveStartArcShapes"
		res = append(res, rf)
	case *TopStackOfGrowthCurve:
		var rf ReverseField
		_ = rf
	case *TopStartArcShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "TopStartArcShapeGrid"
		rf.Fieldname = "TopStartArcShapes"
		res = append(res, rf)
	case *TopStartArcShapeGrid:
		var rf ReverseField
		_ = rf
	}
	return
}

// insertion point for get fields header method
func (arcnormalvectorshape *ArcNormalVectorShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "ArcNormalVectorShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ArcNormalVectorShape",
		},
	}
	return
}

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

func (basevectorshape *BaseVectorShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (basevectorshapegrid *BaseVectorShapeGrid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "BaseVectorShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "BaseVectorShape",
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

func (growthcurve2d *GrowthCurve2D) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
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
			Name:               "RelativeVerticalThickness",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "RadialThickness",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "RhombusSideLength",
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
			Name:                 "BaseVectorShapeGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "BaseVectorShapeGrid",
		},
		{
			Name:                 "ArcNormalVectorShapeGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ArcNormalVectorShapeGrid",
		},
		{
			Name:                 "StartArcShapeGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "StartArcShapeGrid",
		},
		{
			Name:                 "TopStartArcShapeGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TopStartArcShapeGrid",
		},
		{
			Name:                 "EndArcShapeGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "EndArcShapeGrid",
		},
		{
			Name:                 "TopEndArcShapeGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TopEndArcShapeGrid",
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
		{
			Name:                 "TopStackOfGrowthCurve",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TopStackOfGrowthCurve",
		},
		{
			Name:                 "ShiftedLeftStackOfGrowthCurve",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShiftedLeftStackOfGrowthCurve",
		},
		{
			Name:                 "ShiftedLeftStackOfNormalVector",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShiftedLeftStackOfNormalVector",
		},
		{
			Name:                 "GrowthCurve2D",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "GrowthCurve2D",
		},
		{
			Name:                 "TopGrowthCurve2D",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TopGrowthCurve2D",
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
			Name:               "IsRhombusNodesExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
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
			Name:               "IsHiddenBaseVectorShapeGrid",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenArcNormalVectorShapeGrid",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenStartArcShapeGrid",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenTopStartArcShapeGrid",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenEndArcShapeGrid",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenTopEndArcShapeGrid",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenBottomStartArcShapeGrid",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenBottomEndArcShapeGrid",
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
			Name:               "IsHiddenTopStackOfGrowthCurve",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenBottomStackOfGrowthCurve",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenShiftedLeftStackOfGrowthCurve",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenShiftedLeftStackOfNormalVector",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenGrowthCurve2D",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenTopGrowthCurve2D",
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

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "ShiftedLeftStackGrowthCurveStartArcShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ShiftedLeftStackGrowthCurveStartArcShape",
		},
		{
			Name:                 "ShiftedLeftStackGrowthCurveEndArcShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ShiftedLeftStackGrowthCurveEndArcShape",
		},
	}
	return
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "ShiftedLeftStackNormalVectors",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ShiftedLeftStackNormalVector",
		},
	}
	return
}

func (stackgrowthcurveendarcshape *StackGrowthCurveEndArcShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (stackgrowthcurvestartarcshape *StackGrowthCurveStartArcShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (stackofgrowthcurve *StackOfGrowthCurve) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "StackGrowthCurveStartArcShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "StackGrowthCurveStartArcShape",
		},
		{
			Name:                 "StackGrowthCurveEndArcShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "StackGrowthCurveEndArcShape",
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

func (topendarcshape *TopEndArcShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (topendarcshapegrid *TopEndArcShapeGrid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "TopEndArcShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TopEndArcShape",
		},
	}
	return
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "TopStartArcShapeGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TopStartArcShapeGrid",
		},
		{
			Name:                 "TopEndArcShapeGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TopEndArcShapeGrid",
		},
	}
	return
}

func (topstackgrowthcurveendarcshape *TopStackGrowthCurveEndArcShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (topstackgrowthcurvestartarcshape *TopStackGrowthCurveStartArcShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (topstackofgrowthcurve *TopStackOfGrowthCurve) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "TopStackGrowthCurveStartArcShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TopStackGrowthCurveStartArcShape",
		},
		{
			Name:                 "TopStackGrowthCurveEndArcShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TopStackGrowthCurveEndArcShape",
		},
	}
	return
}

func (topstartarcshape *TopStartArcShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (topstartarcshapegrid *TopStartArcShapeGrid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "TopStartArcShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TopStartArcShape",
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
func (arcnormalvectorshape *ArcNormalVectorShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = arcnormalvectorshape.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", arcnormalvectorshape.StartX)
		res.valueFloat = arcnormalvectorshape.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", arcnormalvectorshape.StartY)
		res.valueFloat = arcnormalvectorshape.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", arcnormalvectorshape.EndX)
		res.valueFloat = arcnormalvectorshape.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", arcnormalvectorshape.EndY)
		res.valueFloat = arcnormalvectorshape.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = arcnormalvectorshapegrid.Name
	case "ArcNormalVectorShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range arcnormalvectorshapegrid.ArcNormalVectorShapes {
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

func (basevectorshape *BaseVectorShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = basevectorshape.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", basevectorshape.StartX)
		res.valueFloat = basevectorshape.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", basevectorshape.StartY)
		res.valueFloat = basevectorshape.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", basevectorshape.EndX)
		res.valueFloat = basevectorshape.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", basevectorshape.EndY)
		res.valueFloat = basevectorshape.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = basevectorshapegrid.Name
	case "BaseVectorShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range basevectorshapegrid.BaseVectorShapes {
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

func (growthcurve2d *GrowthCurve2D) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = growthcurve2d.Name
	case "StartArcShapeGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if growthcurve2d.StartArcShapeGrid != nil {
			res.valueString = growthcurve2d.StartArcShapeGrid.Name
			res.ids = growthcurve2d.StartArcShapeGrid.GongGetUUID(stage)
		}
	case "EndArcShapeGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if growthcurve2d.EndArcShapeGrid != nil {
			res.valueString = growthcurve2d.EndArcShapeGrid.Name
			res.ids = growthcurve2d.EndArcShapeGrid.GongGetUUID(stage)
		}
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
	case "RelativeVerticalThickness":
		res.valueString = fmt.Sprintf("%f", plant.RelativeVerticalThickness)
		res.valueFloat = plant.RelativeVerticalThickness
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadialThickness":
		res.valueString = fmt.Sprintf("%f", plant.RadialThickness)
		res.valueFloat = plant.RadialThickness
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RhombusSideLength":
		res.valueString = fmt.Sprintf("%f", plant.RhombusSideLength)
		res.valueFloat = plant.RhombusSideLength
		res.GongFieldValueType = GongFieldValueTypeFloat
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
	case "BaseVectorShapeGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.BaseVectorShapeGrid != nil {
			res.valueString = plant.BaseVectorShapeGrid.Name
			res.ids = plant.BaseVectorShapeGrid.GongGetUUID(stage)
		}
	case "ArcNormalVectorShapeGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.ArcNormalVectorShapeGrid != nil {
			res.valueString = plant.ArcNormalVectorShapeGrid.Name
			res.ids = plant.ArcNormalVectorShapeGrid.GongGetUUID(stage)
		}
	case "StartArcShapeGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.StartArcShapeGrid != nil {
			res.valueString = plant.StartArcShapeGrid.Name
			res.ids = plant.StartArcShapeGrid.GongGetUUID(stage)
		}
	case "TopStartArcShapeGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.TopStartArcShapeGrid != nil {
			res.valueString = plant.TopStartArcShapeGrid.Name
			res.ids = plant.TopStartArcShapeGrid.GongGetUUID(stage)
		}
	case "EndArcShapeGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.EndArcShapeGrid != nil {
			res.valueString = plant.EndArcShapeGrid.Name
			res.ids = plant.EndArcShapeGrid.GongGetUUID(stage)
		}
	case "TopEndArcShapeGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.TopEndArcShapeGrid != nil {
			res.valueString = plant.TopEndArcShapeGrid.Name
			res.ids = plant.TopEndArcShapeGrid.GongGetUUID(stage)
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
	case "TopStackOfGrowthCurve":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.TopStackOfGrowthCurve != nil {
			res.valueString = plant.TopStackOfGrowthCurve.Name
			res.ids = plant.TopStackOfGrowthCurve.GongGetUUID(stage)
		}
	case "ShiftedLeftStackOfGrowthCurve":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.ShiftedLeftStackOfGrowthCurve != nil {
			res.valueString = plant.ShiftedLeftStackOfGrowthCurve.Name
			res.ids = plant.ShiftedLeftStackOfGrowthCurve.GongGetUUID(stage)
		}
	case "ShiftedLeftStackOfNormalVector":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.ShiftedLeftStackOfNormalVector != nil {
			res.valueString = plant.ShiftedLeftStackOfNormalVector.Name
			res.ids = plant.ShiftedLeftStackOfNormalVector.GongGetUUID(stage)
		}
	case "GrowthCurve2D":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.GrowthCurve2D != nil {
			res.valueString = plant.GrowthCurve2D.Name
			res.ids = plant.GrowthCurve2D.GongGetUUID(stage)
		}
	case "TopGrowthCurve2D":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.TopGrowthCurve2D != nil {
			res.valueString = plant.TopGrowthCurve2D.Name
			res.ids = plant.TopGrowthCurve2D.GongGetUUID(stage)
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
	case "IsRhombusNodesExpanded":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsRhombusNodesExpanded)
		res.valueBool = plantdiagram.IsRhombusNodesExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "IsHiddenBaseVectorShapeGrid":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenBaseVectorShapeGrid)
		res.valueBool = plantdiagram.IsHiddenBaseVectorShapeGrid
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenArcNormalVectorShapeGrid":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenArcNormalVectorShapeGrid)
		res.valueBool = plantdiagram.IsHiddenArcNormalVectorShapeGrid
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenStartArcShapeGrid":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenStartArcShapeGrid)
		res.valueBool = plantdiagram.IsHiddenStartArcShapeGrid
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenTopStartArcShapeGrid":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenTopStartArcShapeGrid)
		res.valueBool = plantdiagram.IsHiddenTopStartArcShapeGrid
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenEndArcShapeGrid":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenEndArcShapeGrid)
		res.valueBool = plantdiagram.IsHiddenEndArcShapeGrid
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenTopEndArcShapeGrid":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenTopEndArcShapeGrid)
		res.valueBool = plantdiagram.IsHiddenTopEndArcShapeGrid
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenBottomStartArcShapeGrid":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenBottomStartArcShapeGrid)
		res.valueBool = plantdiagram.IsHiddenBottomStartArcShapeGrid
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenBottomEndArcShapeGrid":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenBottomEndArcShapeGrid)
		res.valueBool = plantdiagram.IsHiddenBottomEndArcShapeGrid
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenGrowthCurveBezierShapeGrid":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenGrowthCurveBezierShapeGrid)
		res.valueBool = plantdiagram.IsHiddenGrowthCurveBezierShapeGrid
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenStackOfGrowthCurve":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenStackOfGrowthCurve)
		res.valueBool = plantdiagram.IsHiddenStackOfGrowthCurve
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenTopStackOfGrowthCurve":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenTopStackOfGrowthCurve)
		res.valueBool = plantdiagram.IsHiddenTopStackOfGrowthCurve
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenBottomStackOfGrowthCurve":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenBottomStackOfGrowthCurve)
		res.valueBool = plantdiagram.IsHiddenBottomStackOfGrowthCurve
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenShiftedLeftStackOfGrowthCurve":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenShiftedLeftStackOfGrowthCurve)
		res.valueBool = plantdiagram.IsHiddenShiftedLeftStackOfGrowthCurve
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenShiftedLeftStackOfNormalVector":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenShiftedLeftStackOfNormalVector)
		res.valueBool = plantdiagram.IsHiddenShiftedLeftStackOfNormalVector
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenGrowthCurve2D":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenGrowthCurve2D)
		res.valueBool = plantdiagram.IsHiddenGrowthCurve2D
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenTopGrowthCurve2D":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenTopGrowthCurve2D)
		res.valueBool = plantdiagram.IsHiddenTopGrowthCurve2D
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

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = shiftedleftstackgrowthcurveendarcshape.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", shiftedleftstackgrowthcurveendarcshape.StartX)
		res.valueFloat = shiftedleftstackgrowthcurveendarcshape.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", shiftedleftstackgrowthcurveendarcshape.StartY)
		res.valueFloat = shiftedleftstackgrowthcurveendarcshape.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", shiftedleftstackgrowthcurveendarcshape.EndX)
		res.valueFloat = shiftedleftstackgrowthcurveendarcshape.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", shiftedleftstackgrowthcurveendarcshape.EndY)
		res.valueFloat = shiftedleftstackgrowthcurveendarcshape.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", shiftedleftstackgrowthcurveendarcshape.XAxisRotation)
		res.valueFloat = shiftedleftstackgrowthcurveendarcshape.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", shiftedleftstackgrowthcurveendarcshape.LargeArcFlag)
		res.valueBool = shiftedleftstackgrowthcurveendarcshape.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", shiftedleftstackgrowthcurveendarcshape.SweepFlag)
		res.valueBool = shiftedleftstackgrowthcurveendarcshape.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", shiftedleftstackgrowthcurveendarcshape.RadiusX)
		res.valueFloat = shiftedleftstackgrowthcurveendarcshape.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", shiftedleftstackgrowthcurveendarcshape.RadiusY)
		res.valueFloat = shiftedleftstackgrowthcurveendarcshape.RadiusY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = shiftedleftstackgrowthcurvestartarcshape.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", shiftedleftstackgrowthcurvestartarcshape.StartX)
		res.valueFloat = shiftedleftstackgrowthcurvestartarcshape.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", shiftedleftstackgrowthcurvestartarcshape.StartY)
		res.valueFloat = shiftedleftstackgrowthcurvestartarcshape.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", shiftedleftstackgrowthcurvestartarcshape.EndX)
		res.valueFloat = shiftedleftstackgrowthcurvestartarcshape.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", shiftedleftstackgrowthcurvestartarcshape.EndY)
		res.valueFloat = shiftedleftstackgrowthcurvestartarcshape.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", shiftedleftstackgrowthcurvestartarcshape.XAxisRotation)
		res.valueFloat = shiftedleftstackgrowthcurvestartarcshape.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", shiftedleftstackgrowthcurvestartarcshape.LargeArcFlag)
		res.valueBool = shiftedleftstackgrowthcurvestartarcshape.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", shiftedleftstackgrowthcurvestartarcshape.SweepFlag)
		res.valueBool = shiftedleftstackgrowthcurvestartarcshape.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", shiftedleftstackgrowthcurvestartarcshape.RadiusX)
		res.valueFloat = shiftedleftstackgrowthcurvestartarcshape.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", shiftedleftstackgrowthcurvestartarcshape.RadiusY)
		res.valueFloat = shiftedleftstackgrowthcurvestartarcshape.RadiusY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = shiftedleftstacknormalvector.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", shiftedleftstacknormalvector.StartX)
		res.valueFloat = shiftedleftstacknormalvector.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", shiftedleftstacknormalvector.StartY)
		res.valueFloat = shiftedleftstacknormalvector.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", shiftedleftstacknormalvector.EndX)
		res.valueFloat = shiftedleftstacknormalvector.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", shiftedleftstacknormalvector.EndY)
		res.valueFloat = shiftedleftstacknormalvector.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = shiftedleftstackofgrowthcurve.Name
	case "ShiftedLeftStackGrowthCurveStartArcShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ShiftedLeftStackGrowthCurveEndArcShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes {
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

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = shiftedleftstackofnormalvector.Name
	case "ShiftedLeftStackNormalVectors":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors {
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

func (stackgrowthcurveendarcshape *StackGrowthCurveEndArcShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = stackgrowthcurveendarcshape.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurveendarcshape.StartX)
		res.valueFloat = stackgrowthcurveendarcshape.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurveendarcshape.StartY)
		res.valueFloat = stackgrowthcurveendarcshape.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurveendarcshape.EndX)
		res.valueFloat = stackgrowthcurveendarcshape.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurveendarcshape.EndY)
		res.valueFloat = stackgrowthcurveendarcshape.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurveendarcshape.XAxisRotation)
		res.valueFloat = stackgrowthcurveendarcshape.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", stackgrowthcurveendarcshape.LargeArcFlag)
		res.valueBool = stackgrowthcurveendarcshape.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", stackgrowthcurveendarcshape.SweepFlag)
		res.valueBool = stackgrowthcurveendarcshape.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurveendarcshape.RadiusX)
		res.valueFloat = stackgrowthcurveendarcshape.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurveendarcshape.RadiusY)
		res.valueFloat = stackgrowthcurveendarcshape.RadiusY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (stackgrowthcurvestartarcshape *StackGrowthCurveStartArcShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = stackgrowthcurvestartarcshape.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvestartarcshape.StartX)
		res.valueFloat = stackgrowthcurvestartarcshape.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvestartarcshape.StartY)
		res.valueFloat = stackgrowthcurvestartarcshape.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvestartarcshape.EndX)
		res.valueFloat = stackgrowthcurvestartarcshape.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvestartarcshape.EndY)
		res.valueFloat = stackgrowthcurvestartarcshape.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvestartarcshape.XAxisRotation)
		res.valueFloat = stackgrowthcurvestartarcshape.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", stackgrowthcurvestartarcshape.LargeArcFlag)
		res.valueBool = stackgrowthcurvestartarcshape.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", stackgrowthcurvestartarcshape.SweepFlag)
		res.valueBool = stackgrowthcurvestartarcshape.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvestartarcshape.RadiusX)
		res.valueFloat = stackgrowthcurvestartarcshape.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvestartarcshape.RadiusY)
		res.valueFloat = stackgrowthcurvestartarcshape.RadiusY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (stackofgrowthcurve *StackOfGrowthCurve) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = stackofgrowthcurve.Name
	case "StackGrowthCurveStartArcShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range stackofgrowthcurve.StackGrowthCurveStartArcShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "StackGrowthCurveEndArcShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range stackofgrowthcurve.StackGrowthCurveEndArcShapes {
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

func (topendarcshape *TopEndArcShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = topendarcshape.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", topendarcshape.StartX)
		res.valueFloat = topendarcshape.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", topendarcshape.StartY)
		res.valueFloat = topendarcshape.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", topendarcshape.EndX)
		res.valueFloat = topendarcshape.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", topendarcshape.EndY)
		res.valueFloat = topendarcshape.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", topendarcshape.XAxisRotation)
		res.valueFloat = topendarcshape.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", topendarcshape.LargeArcFlag)
		res.valueBool = topendarcshape.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", topendarcshape.SweepFlag)
		res.valueBool = topendarcshape.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", topendarcshape.RadiusX)
		res.valueFloat = topendarcshape.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", topendarcshape.RadiusY)
		res.valueFloat = topendarcshape.RadiusY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = topendarcshapegrid.Name
	case "TopEndArcShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range topendarcshapegrid.TopEndArcShapes {
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

func (topgrowthcurve2d *TopGrowthCurve2D) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = topgrowthcurve2d.Name
	case "TopStartArcShapeGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if topgrowthcurve2d.TopStartArcShapeGrid != nil {
			res.valueString = topgrowthcurve2d.TopStartArcShapeGrid.Name
			res.ids = topgrowthcurve2d.TopStartArcShapeGrid.GongGetUUID(stage)
		}
	case "TopEndArcShapeGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if topgrowthcurve2d.TopEndArcShapeGrid != nil {
			res.valueString = topgrowthcurve2d.TopEndArcShapeGrid.Name
			res.ids = topgrowthcurve2d.TopEndArcShapeGrid.GongGetUUID(stage)
		}
	}
	return
}

func (topstackgrowthcurveendarcshape *TopStackGrowthCurveEndArcShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = topstackgrowthcurveendarcshape.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurveendarcshape.StartX)
		res.valueFloat = topstackgrowthcurveendarcshape.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurveendarcshape.StartY)
		res.valueFloat = topstackgrowthcurveendarcshape.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurveendarcshape.EndX)
		res.valueFloat = topstackgrowthcurveendarcshape.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurveendarcshape.EndY)
		res.valueFloat = topstackgrowthcurveendarcshape.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurveendarcshape.XAxisRotation)
		res.valueFloat = topstackgrowthcurveendarcshape.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", topstackgrowthcurveendarcshape.LargeArcFlag)
		res.valueBool = topstackgrowthcurveendarcshape.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", topstackgrowthcurveendarcshape.SweepFlag)
		res.valueBool = topstackgrowthcurveendarcshape.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurveendarcshape.RadiusX)
		res.valueFloat = topstackgrowthcurveendarcshape.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurveendarcshape.RadiusY)
		res.valueFloat = topstackgrowthcurveendarcshape.RadiusY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (topstackgrowthcurvestartarcshape *TopStackGrowthCurveStartArcShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = topstackgrowthcurvestartarcshape.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurvestartarcshape.StartX)
		res.valueFloat = topstackgrowthcurvestartarcshape.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurvestartarcshape.StartY)
		res.valueFloat = topstackgrowthcurvestartarcshape.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurvestartarcshape.EndX)
		res.valueFloat = topstackgrowthcurvestartarcshape.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurvestartarcshape.EndY)
		res.valueFloat = topstackgrowthcurvestartarcshape.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurvestartarcshape.XAxisRotation)
		res.valueFloat = topstackgrowthcurvestartarcshape.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", topstackgrowthcurvestartarcshape.LargeArcFlag)
		res.valueBool = topstackgrowthcurvestartarcshape.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", topstackgrowthcurvestartarcshape.SweepFlag)
		res.valueBool = topstackgrowthcurvestartarcshape.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurvestartarcshape.RadiusX)
		res.valueFloat = topstackgrowthcurvestartarcshape.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurvestartarcshape.RadiusY)
		res.valueFloat = topstackgrowthcurvestartarcshape.RadiusY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (topstackofgrowthcurve *TopStackOfGrowthCurve) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = topstackofgrowthcurve.Name
	case "TopStackGrowthCurveStartArcShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "TopStackGrowthCurveEndArcShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes {
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

func (topstartarcshape *TopStartArcShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = topstartarcshape.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", topstartarcshape.StartX)
		res.valueFloat = topstartarcshape.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", topstartarcshape.StartY)
		res.valueFloat = topstartarcshape.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", topstartarcshape.EndX)
		res.valueFloat = topstartarcshape.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", topstartarcshape.EndY)
		res.valueFloat = topstartarcshape.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", topstartarcshape.XAxisRotation)
		res.valueFloat = topstartarcshape.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", topstartarcshape.LargeArcFlag)
		res.valueBool = topstartarcshape.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", topstartarcshape.SweepFlag)
		res.valueBool = topstartarcshape.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", topstartarcshape.RadiusX)
		res.valueFloat = topstartarcshape.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", topstartarcshape.RadiusY)
		res.valueFloat = topstartarcshape.RadiusY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = topstartarcshapegrid.Name
	case "TopStartArcShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range topstartarcshapegrid.TopStartArcShapes {
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
func (arcnormalvectorshape *ArcNormalVectorShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		arcnormalvectorshape.Name = value.GetValueString()
	case "StartX":
		arcnormalvectorshape.StartX = value.GetValueFloat()
	case "StartY":
		arcnormalvectorshape.StartY = value.GetValueFloat()
	case "EndX":
		arcnormalvectorshape.EndX = value.GetValueFloat()
	case "EndY":
		arcnormalvectorshape.EndY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		arcnormalvectorshapegrid.Name = value.GetValueString()
	case "ArcNormalVectorShapes":
		arcnormalvectorshapegrid.ArcNormalVectorShapes = make([]*ArcNormalVectorShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ArcNormalVectorShapes {
					if stage.ArcNormalVectorShape_stagedOrder[__instance__] == uint(id) {
						arcnormalvectorshapegrid.ArcNormalVectorShapes = append(arcnormalvectorshapegrid.ArcNormalVectorShapes, __instance__)
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

func (basevectorshape *BaseVectorShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		basevectorshape.Name = value.GetValueString()
	case "StartX":
		basevectorshape.StartX = value.GetValueFloat()
	case "StartY":
		basevectorshape.StartY = value.GetValueFloat()
	case "EndX":
		basevectorshape.EndX = value.GetValueFloat()
	case "EndY":
		basevectorshape.EndY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		basevectorshapegrid.Name = value.GetValueString()
	case "BaseVectorShapes":
		basevectorshapegrid.BaseVectorShapes = make([]*BaseVectorShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.BaseVectorShapes {
					if stage.BaseVectorShape_stagedOrder[__instance__] == uint(id) {
						basevectorshapegrid.BaseVectorShapes = append(basevectorshapegrid.BaseVectorShapes, __instance__)
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

func (growthcurve2d *GrowthCurve2D) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		growthcurve2d.Name = value.GetValueString()
	case "StartArcShapeGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			growthcurve2d.StartArcShapeGrid = nil
			for __instance__ := range stage.StartArcShapeGrids {
				if stage.StartArcShapeGrid_stagedOrder[__instance__] == uint(id) {
					growthcurve2d.StartArcShapeGrid = __instance__
					break
				}
			}
		}
	case "EndArcShapeGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			growthcurve2d.EndArcShapeGrid = nil
			for __instance__ := range stage.EndArcShapeGrids {
				if stage.EndArcShapeGrid_stagedOrder[__instance__] == uint(id) {
					growthcurve2d.EndArcShapeGrid = __instance__
					break
				}
			}
		}
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
	case "RelativeVerticalThickness":
		plant.RelativeVerticalThickness = value.GetValueFloat()
	case "RadialThickness":
		plant.RadialThickness = value.GetValueFloat()
	case "RhombusSideLength":
		plant.RhombusSideLength = value.GetValueFloat()
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
	case "BaseVectorShapeGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.BaseVectorShapeGrid = nil
			for __instance__ := range stage.BaseVectorShapeGrids {
				if stage.BaseVectorShapeGrid_stagedOrder[__instance__] == uint(id) {
					plant.BaseVectorShapeGrid = __instance__
					break
				}
			}
		}
	case "ArcNormalVectorShapeGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.ArcNormalVectorShapeGrid = nil
			for __instance__ := range stage.ArcNormalVectorShapeGrids {
				if stage.ArcNormalVectorShapeGrid_stagedOrder[__instance__] == uint(id) {
					plant.ArcNormalVectorShapeGrid = __instance__
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
	case "TopStartArcShapeGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.TopStartArcShapeGrid = nil
			for __instance__ := range stage.TopStartArcShapeGrids {
				if stage.TopStartArcShapeGrid_stagedOrder[__instance__] == uint(id) {
					plant.TopStartArcShapeGrid = __instance__
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
	case "TopEndArcShapeGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.TopEndArcShapeGrid = nil
			for __instance__ := range stage.TopEndArcShapeGrids {
				if stage.TopEndArcShapeGrid_stagedOrder[__instance__] == uint(id) {
					plant.TopEndArcShapeGrid = __instance__
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
	case "TopStackOfGrowthCurve":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.TopStackOfGrowthCurve = nil
			for __instance__ := range stage.TopStackOfGrowthCurves {
				if stage.TopStackOfGrowthCurve_stagedOrder[__instance__] == uint(id) {
					plant.TopStackOfGrowthCurve = __instance__
					break
				}
			}
		}
	case "ShiftedLeftStackOfGrowthCurve":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.ShiftedLeftStackOfGrowthCurve = nil
			for __instance__ := range stage.ShiftedLeftStackOfGrowthCurves {
				if stage.ShiftedLeftStackOfGrowthCurve_stagedOrder[__instance__] == uint(id) {
					plant.ShiftedLeftStackOfGrowthCurve = __instance__
					break
				}
			}
		}
	case "ShiftedLeftStackOfNormalVector":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.ShiftedLeftStackOfNormalVector = nil
			for __instance__ := range stage.ShiftedLeftStackOfNormalVectors {
				if stage.ShiftedLeftStackOfNormalVector_stagedOrder[__instance__] == uint(id) {
					plant.ShiftedLeftStackOfNormalVector = __instance__
					break
				}
			}
		}
	case "GrowthCurve2D":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.GrowthCurve2D = nil
			for __instance__ := range stage.GrowthCurve2Ds {
				if stage.GrowthCurve2D_stagedOrder[__instance__] == uint(id) {
					plant.GrowthCurve2D = __instance__
					break
				}
			}
		}
	case "TopGrowthCurve2D":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.TopGrowthCurve2D = nil
			for __instance__ := range stage.TopGrowthCurve2Ds {
				if stage.TopGrowthCurve2D_stagedOrder[__instance__] == uint(id) {
					plant.TopGrowthCurve2D = __instance__
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
	case "IsRhombusNodesExpanded":
		plantdiagram.IsRhombusNodesExpanded = value.GetValueBool()
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
	case "IsHiddenBaseVectorShapeGrid":
		plantdiagram.IsHiddenBaseVectorShapeGrid = value.GetValueBool()
	case "IsHiddenArcNormalVectorShapeGrid":
		plantdiagram.IsHiddenArcNormalVectorShapeGrid = value.GetValueBool()
	case "IsHiddenStartArcShapeGrid":
		plantdiagram.IsHiddenStartArcShapeGrid = value.GetValueBool()
	case "IsHiddenTopStartArcShapeGrid":
		plantdiagram.IsHiddenTopStartArcShapeGrid = value.GetValueBool()
	case "IsHiddenEndArcShapeGrid":
		plantdiagram.IsHiddenEndArcShapeGrid = value.GetValueBool()
	case "IsHiddenTopEndArcShapeGrid":
		plantdiagram.IsHiddenTopEndArcShapeGrid = value.GetValueBool()
	case "IsHiddenBottomStartArcShapeGrid":
		plantdiagram.IsHiddenBottomStartArcShapeGrid = value.GetValueBool()
	case "IsHiddenBottomEndArcShapeGrid":
		plantdiagram.IsHiddenBottomEndArcShapeGrid = value.GetValueBool()
	case "IsHiddenGrowthCurveBezierShapeGrid":
		plantdiagram.IsHiddenGrowthCurveBezierShapeGrid = value.GetValueBool()
	case "IsHiddenStackOfGrowthCurve":
		plantdiagram.IsHiddenStackOfGrowthCurve = value.GetValueBool()
	case "IsHiddenTopStackOfGrowthCurve":
		plantdiagram.IsHiddenTopStackOfGrowthCurve = value.GetValueBool()
	case "IsHiddenBottomStackOfGrowthCurve":
		plantdiagram.IsHiddenBottomStackOfGrowthCurve = value.GetValueBool()
	case "IsHiddenShiftedLeftStackOfGrowthCurve":
		plantdiagram.IsHiddenShiftedLeftStackOfGrowthCurve = value.GetValueBool()
	case "IsHiddenShiftedLeftStackOfNormalVector":
		plantdiagram.IsHiddenShiftedLeftStackOfNormalVector = value.GetValueBool()
	case "IsHiddenGrowthCurve2D":
		plantdiagram.IsHiddenGrowthCurve2D = value.GetValueBool()
	case "IsHiddenTopGrowthCurve2D":
		plantdiagram.IsHiddenTopGrowthCurve2D = value.GetValueBool()
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

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		shiftedleftstackgrowthcurveendarcshape.Name = value.GetValueString()
	case "StartX":
		shiftedleftstackgrowthcurveendarcshape.StartX = value.GetValueFloat()
	case "StartY":
		shiftedleftstackgrowthcurveendarcshape.StartY = value.GetValueFloat()
	case "EndX":
		shiftedleftstackgrowthcurveendarcshape.EndX = value.GetValueFloat()
	case "EndY":
		shiftedleftstackgrowthcurveendarcshape.EndY = value.GetValueFloat()
	case "XAxisRotation":
		shiftedleftstackgrowthcurveendarcshape.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		shiftedleftstackgrowthcurveendarcshape.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		shiftedleftstackgrowthcurveendarcshape.SweepFlag = value.GetValueBool()
	case "RadiusX":
		shiftedleftstackgrowthcurveendarcshape.RadiusX = value.GetValueFloat()
	case "RadiusY":
		shiftedleftstackgrowthcurveendarcshape.RadiusY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		shiftedleftstackgrowthcurvestartarcshape.Name = value.GetValueString()
	case "StartX":
		shiftedleftstackgrowthcurvestartarcshape.StartX = value.GetValueFloat()
	case "StartY":
		shiftedleftstackgrowthcurvestartarcshape.StartY = value.GetValueFloat()
	case "EndX":
		shiftedleftstackgrowthcurvestartarcshape.EndX = value.GetValueFloat()
	case "EndY":
		shiftedleftstackgrowthcurvestartarcshape.EndY = value.GetValueFloat()
	case "XAxisRotation":
		shiftedleftstackgrowthcurvestartarcshape.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		shiftedleftstackgrowthcurvestartarcshape.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		shiftedleftstackgrowthcurvestartarcshape.SweepFlag = value.GetValueBool()
	case "RadiusX":
		shiftedleftstackgrowthcurvestartarcshape.RadiusX = value.GetValueFloat()
	case "RadiusY":
		shiftedleftstackgrowthcurvestartarcshape.RadiusY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		shiftedleftstacknormalvector.Name = value.GetValueString()
	case "StartX":
		shiftedleftstacknormalvector.StartX = value.GetValueFloat()
	case "StartY":
		shiftedleftstacknormalvector.StartY = value.GetValueFloat()
	case "EndX":
		shiftedleftstacknormalvector.EndX = value.GetValueFloat()
	case "EndY":
		shiftedleftstacknormalvector.EndY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		shiftedleftstackofgrowthcurve.Name = value.GetValueString()
	case "ShiftedLeftStackGrowthCurveStartArcShapes":
		shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes = make([]*ShiftedLeftStackGrowthCurveStartArcShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ShiftedLeftStackGrowthCurveStartArcShapes {
					if stage.ShiftedLeftStackGrowthCurveStartArcShape_stagedOrder[__instance__] == uint(id) {
						shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes = append(shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes, __instance__)
						break
					}
				}
			}
		}
	case "ShiftedLeftStackGrowthCurveEndArcShapes":
		shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes = make([]*ShiftedLeftStackGrowthCurveEndArcShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ShiftedLeftStackGrowthCurveEndArcShapes {
					if stage.ShiftedLeftStackGrowthCurveEndArcShape_stagedOrder[__instance__] == uint(id) {
						shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes = append(shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes, __instance__)
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

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		shiftedleftstackofnormalvector.Name = value.GetValueString()
	case "ShiftedLeftStackNormalVectors":
		shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors = make([]*ShiftedLeftStackNormalVector, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ShiftedLeftStackNormalVectors {
					if stage.ShiftedLeftStackNormalVector_stagedOrder[__instance__] == uint(id) {
						shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors = append(shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors, __instance__)
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

func (stackgrowthcurveendarcshape *StackGrowthCurveEndArcShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		stackgrowthcurveendarcshape.Name = value.GetValueString()
	case "StartX":
		stackgrowthcurveendarcshape.StartX = value.GetValueFloat()
	case "StartY":
		stackgrowthcurveendarcshape.StartY = value.GetValueFloat()
	case "EndX":
		stackgrowthcurveendarcshape.EndX = value.GetValueFloat()
	case "EndY":
		stackgrowthcurveendarcshape.EndY = value.GetValueFloat()
	case "XAxisRotation":
		stackgrowthcurveendarcshape.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		stackgrowthcurveendarcshape.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		stackgrowthcurveendarcshape.SweepFlag = value.GetValueBool()
	case "RadiusX":
		stackgrowthcurveendarcshape.RadiusX = value.GetValueFloat()
	case "RadiusY":
		stackgrowthcurveendarcshape.RadiusY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (stackgrowthcurvestartarcshape *StackGrowthCurveStartArcShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		stackgrowthcurvestartarcshape.Name = value.GetValueString()
	case "StartX":
		stackgrowthcurvestartarcshape.StartX = value.GetValueFloat()
	case "StartY":
		stackgrowthcurvestartarcshape.StartY = value.GetValueFloat()
	case "EndX":
		stackgrowthcurvestartarcshape.EndX = value.GetValueFloat()
	case "EndY":
		stackgrowthcurvestartarcshape.EndY = value.GetValueFloat()
	case "XAxisRotation":
		stackgrowthcurvestartarcshape.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		stackgrowthcurvestartarcshape.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		stackgrowthcurvestartarcshape.SweepFlag = value.GetValueBool()
	case "RadiusX":
		stackgrowthcurvestartarcshape.RadiusX = value.GetValueFloat()
	case "RadiusY":
		stackgrowthcurvestartarcshape.RadiusY = value.GetValueFloat()
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
	case "StackGrowthCurveStartArcShapes":
		stackofgrowthcurve.StackGrowthCurveStartArcShapes = make([]*StackGrowthCurveStartArcShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.StackGrowthCurveStartArcShapes {
					if stage.StackGrowthCurveStartArcShape_stagedOrder[__instance__] == uint(id) {
						stackofgrowthcurve.StackGrowthCurveStartArcShapes = append(stackofgrowthcurve.StackGrowthCurveStartArcShapes, __instance__)
						break
					}
				}
			}
		}
	case "StackGrowthCurveEndArcShapes":
		stackofgrowthcurve.StackGrowthCurveEndArcShapes = make([]*StackGrowthCurveEndArcShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.StackGrowthCurveEndArcShapes {
					if stage.StackGrowthCurveEndArcShape_stagedOrder[__instance__] == uint(id) {
						stackofgrowthcurve.StackGrowthCurveEndArcShapes = append(stackofgrowthcurve.StackGrowthCurveEndArcShapes, __instance__)
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

func (topendarcshape *TopEndArcShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		topendarcshape.Name = value.GetValueString()
	case "StartX":
		topendarcshape.StartX = value.GetValueFloat()
	case "StartY":
		topendarcshape.StartY = value.GetValueFloat()
	case "EndX":
		topendarcshape.EndX = value.GetValueFloat()
	case "EndY":
		topendarcshape.EndY = value.GetValueFloat()
	case "XAxisRotation":
		topendarcshape.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		topendarcshape.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		topendarcshape.SweepFlag = value.GetValueBool()
	case "RadiusX":
		topendarcshape.RadiusX = value.GetValueFloat()
	case "RadiusY":
		topendarcshape.RadiusY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		topendarcshapegrid.Name = value.GetValueString()
	case "TopEndArcShapes":
		topendarcshapegrid.TopEndArcShapes = make([]*TopEndArcShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TopEndArcShapes {
					if stage.TopEndArcShape_stagedOrder[__instance__] == uint(id) {
						topendarcshapegrid.TopEndArcShapes = append(topendarcshapegrid.TopEndArcShapes, __instance__)
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

func (topgrowthcurve2d *TopGrowthCurve2D) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		topgrowthcurve2d.Name = value.GetValueString()
	case "TopStartArcShapeGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			topgrowthcurve2d.TopStartArcShapeGrid = nil
			for __instance__ := range stage.TopStartArcShapeGrids {
				if stage.TopStartArcShapeGrid_stagedOrder[__instance__] == uint(id) {
					topgrowthcurve2d.TopStartArcShapeGrid = __instance__
					break
				}
			}
		}
	case "TopEndArcShapeGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			topgrowthcurve2d.TopEndArcShapeGrid = nil
			for __instance__ := range stage.TopEndArcShapeGrids {
				if stage.TopEndArcShapeGrid_stagedOrder[__instance__] == uint(id) {
					topgrowthcurve2d.TopEndArcShapeGrid = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (topstackgrowthcurveendarcshape *TopStackGrowthCurveEndArcShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		topstackgrowthcurveendarcshape.Name = value.GetValueString()
	case "StartX":
		topstackgrowthcurveendarcshape.StartX = value.GetValueFloat()
	case "StartY":
		topstackgrowthcurveendarcshape.StartY = value.GetValueFloat()
	case "EndX":
		topstackgrowthcurveendarcshape.EndX = value.GetValueFloat()
	case "EndY":
		topstackgrowthcurveendarcshape.EndY = value.GetValueFloat()
	case "XAxisRotation":
		topstackgrowthcurveendarcshape.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		topstackgrowthcurveendarcshape.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		topstackgrowthcurveendarcshape.SweepFlag = value.GetValueBool()
	case "RadiusX":
		topstackgrowthcurveendarcshape.RadiusX = value.GetValueFloat()
	case "RadiusY":
		topstackgrowthcurveendarcshape.RadiusY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (topstackgrowthcurvestartarcshape *TopStackGrowthCurveStartArcShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		topstackgrowthcurvestartarcshape.Name = value.GetValueString()
	case "StartX":
		topstackgrowthcurvestartarcshape.StartX = value.GetValueFloat()
	case "StartY":
		topstackgrowthcurvestartarcshape.StartY = value.GetValueFloat()
	case "EndX":
		topstackgrowthcurvestartarcshape.EndX = value.GetValueFloat()
	case "EndY":
		topstackgrowthcurvestartarcshape.EndY = value.GetValueFloat()
	case "XAxisRotation":
		topstackgrowthcurvestartarcshape.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		topstackgrowthcurvestartarcshape.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		topstackgrowthcurvestartarcshape.SweepFlag = value.GetValueBool()
	case "RadiusX":
		topstackgrowthcurvestartarcshape.RadiusX = value.GetValueFloat()
	case "RadiusY":
		topstackgrowthcurvestartarcshape.RadiusY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (topstackofgrowthcurve *TopStackOfGrowthCurve) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		topstackofgrowthcurve.Name = value.GetValueString()
	case "TopStackGrowthCurveStartArcShapes":
		topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes = make([]*TopStackGrowthCurveStartArcShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TopStackGrowthCurveStartArcShapes {
					if stage.TopStackGrowthCurveStartArcShape_stagedOrder[__instance__] == uint(id) {
						topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes = append(topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes, __instance__)
						break
					}
				}
			}
		}
	case "TopStackGrowthCurveEndArcShapes":
		topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes = make([]*TopStackGrowthCurveEndArcShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TopStackGrowthCurveEndArcShapes {
					if stage.TopStackGrowthCurveEndArcShape_stagedOrder[__instance__] == uint(id) {
						topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes = append(topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes, __instance__)
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

func (topstartarcshape *TopStartArcShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		topstartarcshape.Name = value.GetValueString()
	case "StartX":
		topstartarcshape.StartX = value.GetValueFloat()
	case "StartY":
		topstartarcshape.StartY = value.GetValueFloat()
	case "EndX":
		topstartarcshape.EndX = value.GetValueFloat()
	case "EndY":
		topstartarcshape.EndY = value.GetValueFloat()
	case "XAxisRotation":
		topstartarcshape.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		topstartarcshape.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		topstartarcshape.SweepFlag = value.GetValueBool()
	case "RadiusX":
		topstartarcshape.RadiusX = value.GetValueFloat()
	case "RadiusY":
		topstartarcshape.RadiusY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		topstartarcshapegrid.Name = value.GetValueString()
	case "TopStartArcShapes":
		topstartarcshapegrid.TopStartArcShapes = make([]*TopStartArcShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TopStartArcShapes {
					if stage.TopStartArcShape_stagedOrder[__instance__] == uint(id) {
						topstartarcshapegrid.TopStartArcShapes = append(topstartarcshapegrid.TopStartArcShapes, __instance__)
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
func (arcnormalvectorshape *ArcNormalVectorShape) GongGetGongstructName() string {
	return "ArcNormalVectorShape"
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongGetGongstructName() string {
	return "ArcNormalVectorShapeGrid"
}

func (axesshape *AxesShape) GongGetGongstructName() string {
	return "AxesShape"
}

func (basevectorshape *BaseVectorShape) GongGetGongstructName() string {
	return "BaseVectorShape"
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongGetGongstructName() string {
	return "BaseVectorShapeGrid"
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

func (growthcurve2d *GrowthCurve2D) GongGetGongstructName() string {
	return "GrowthCurve2D"
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

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongGetGongstructName() string {
	return "ShiftedLeftStackGrowthCurveEndArcShape"
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongGetGongstructName() string {
	return "ShiftedLeftStackGrowthCurveStartArcShape"
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongGetGongstructName() string {
	return "ShiftedLeftStackNormalVector"
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongGetGongstructName() string {
	return "ShiftedLeftStackOfGrowthCurve"
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongGetGongstructName() string {
	return "ShiftedLeftStackOfNormalVector"
}

func (stackgrowthcurveendarcshape *StackGrowthCurveEndArcShape) GongGetGongstructName() string {
	return "StackGrowthCurveEndArcShape"
}

func (stackgrowthcurvestartarcshape *StackGrowthCurveStartArcShape) GongGetGongstructName() string {
	return "StackGrowthCurveStartArcShape"
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

func (topendarcshape *TopEndArcShape) GongGetGongstructName() string {
	return "TopEndArcShape"
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongGetGongstructName() string {
	return "TopEndArcShapeGrid"
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongGetGongstructName() string {
	return "TopGrowthCurve2D"
}

func (topstackgrowthcurveendarcshape *TopStackGrowthCurveEndArcShape) GongGetGongstructName() string {
	return "TopStackGrowthCurveEndArcShape"
}

func (topstackgrowthcurvestartarcshape *TopStackGrowthCurveStartArcShape) GongGetGongstructName() string {
	return "TopStackGrowthCurveStartArcShape"
}

func (topstackofgrowthcurve *TopStackOfGrowthCurve) GongGetGongstructName() string {
	return "TopStackOfGrowthCurve"
}

func (topstartarcshape *TopStartArcShape) GongGetGongstructName() string {
	return "TopStartArcShape"
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongGetGongstructName() string {
	return "TopStartArcShapeGrid"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	stage.ArcNormalVectorShapes_mapString = make(map[string]*ArcNormalVectorShape)
	for arcnormalvectorshape := range stage.ArcNormalVectorShapes {
		stage.ArcNormalVectorShapes_mapString[arcnormalvectorshape.Name] = arcnormalvectorshape
	}

	stage.ArcNormalVectorShapeGrids_mapString = make(map[string]*ArcNormalVectorShapeGrid)
	for arcnormalvectorshapegrid := range stage.ArcNormalVectorShapeGrids {
		stage.ArcNormalVectorShapeGrids_mapString[arcnormalvectorshapegrid.Name] = arcnormalvectorshapegrid
	}

	stage.AxesShapes_mapString = make(map[string]*AxesShape)
	for axesshape := range stage.AxesShapes {
		stage.AxesShapes_mapString[axesshape.Name] = axesshape
	}

	stage.BaseVectorShapes_mapString = make(map[string]*BaseVectorShape)
	for basevectorshape := range stage.BaseVectorShapes {
		stage.BaseVectorShapes_mapString[basevectorshape.Name] = basevectorshape
	}

	stage.BaseVectorShapeGrids_mapString = make(map[string]*BaseVectorShapeGrid)
	for basevectorshapegrid := range stage.BaseVectorShapeGrids {
		stage.BaseVectorShapeGrids_mapString[basevectorshapegrid.Name] = basevectorshapegrid
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

	stage.GrowthCurve2Ds_mapString = make(map[string]*GrowthCurve2D)
	for growthcurve2d := range stage.GrowthCurve2Ds {
		stage.GrowthCurve2Ds_mapString[growthcurve2d.Name] = growthcurve2d
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

	stage.ShiftedLeftStackGrowthCurveEndArcShapes_mapString = make(map[string]*ShiftedLeftStackGrowthCurveEndArcShape)
	for shiftedleftstackgrowthcurveendarcshape := range stage.ShiftedLeftStackGrowthCurveEndArcShapes {
		stage.ShiftedLeftStackGrowthCurveEndArcShapes_mapString[shiftedleftstackgrowthcurveendarcshape.Name] = shiftedleftstackgrowthcurveendarcshape
	}

	stage.ShiftedLeftStackGrowthCurveStartArcShapes_mapString = make(map[string]*ShiftedLeftStackGrowthCurveStartArcShape)
	for shiftedleftstackgrowthcurvestartarcshape := range stage.ShiftedLeftStackGrowthCurveStartArcShapes {
		stage.ShiftedLeftStackGrowthCurveStartArcShapes_mapString[shiftedleftstackgrowthcurvestartarcshape.Name] = shiftedleftstackgrowthcurvestartarcshape
	}

	stage.ShiftedLeftStackNormalVectors_mapString = make(map[string]*ShiftedLeftStackNormalVector)
	for shiftedleftstacknormalvector := range stage.ShiftedLeftStackNormalVectors {
		stage.ShiftedLeftStackNormalVectors_mapString[shiftedleftstacknormalvector.Name] = shiftedleftstacknormalvector
	}

	stage.ShiftedLeftStackOfGrowthCurves_mapString = make(map[string]*ShiftedLeftStackOfGrowthCurve)
	for shiftedleftstackofgrowthcurve := range stage.ShiftedLeftStackOfGrowthCurves {
		stage.ShiftedLeftStackOfGrowthCurves_mapString[shiftedleftstackofgrowthcurve.Name] = shiftedleftstackofgrowthcurve
	}

	stage.ShiftedLeftStackOfNormalVectors_mapString = make(map[string]*ShiftedLeftStackOfNormalVector)
	for shiftedleftstackofnormalvector := range stage.ShiftedLeftStackOfNormalVectors {
		stage.ShiftedLeftStackOfNormalVectors_mapString[shiftedleftstackofnormalvector.Name] = shiftedleftstackofnormalvector
	}

	stage.StackGrowthCurveEndArcShapes_mapString = make(map[string]*StackGrowthCurveEndArcShape)
	for stackgrowthcurveendarcshape := range stage.StackGrowthCurveEndArcShapes {
		stage.StackGrowthCurveEndArcShapes_mapString[stackgrowthcurveendarcshape.Name] = stackgrowthcurveendarcshape
	}

	stage.StackGrowthCurveStartArcShapes_mapString = make(map[string]*StackGrowthCurveStartArcShape)
	for stackgrowthcurvestartarcshape := range stage.StackGrowthCurveStartArcShapes {
		stage.StackGrowthCurveStartArcShapes_mapString[stackgrowthcurvestartarcshape.Name] = stackgrowthcurvestartarcshape
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

	stage.TopEndArcShapes_mapString = make(map[string]*TopEndArcShape)
	for topendarcshape := range stage.TopEndArcShapes {
		stage.TopEndArcShapes_mapString[topendarcshape.Name] = topendarcshape
	}

	stage.TopEndArcShapeGrids_mapString = make(map[string]*TopEndArcShapeGrid)
	for topendarcshapegrid := range stage.TopEndArcShapeGrids {
		stage.TopEndArcShapeGrids_mapString[topendarcshapegrid.Name] = topendarcshapegrid
	}

	stage.TopGrowthCurve2Ds_mapString = make(map[string]*TopGrowthCurve2D)
	for topgrowthcurve2d := range stage.TopGrowthCurve2Ds {
		stage.TopGrowthCurve2Ds_mapString[topgrowthcurve2d.Name] = topgrowthcurve2d
	}

	stage.TopStackGrowthCurveEndArcShapes_mapString = make(map[string]*TopStackGrowthCurveEndArcShape)
	for topstackgrowthcurveendarcshape := range stage.TopStackGrowthCurveEndArcShapes {
		stage.TopStackGrowthCurveEndArcShapes_mapString[topstackgrowthcurveendarcshape.Name] = topstackgrowthcurveendarcshape
	}

	stage.TopStackGrowthCurveStartArcShapes_mapString = make(map[string]*TopStackGrowthCurveStartArcShape)
	for topstackgrowthcurvestartarcshape := range stage.TopStackGrowthCurveStartArcShapes {
		stage.TopStackGrowthCurveStartArcShapes_mapString[topstackgrowthcurvestartarcshape.Name] = topstackgrowthcurvestartarcshape
	}

	stage.TopStackOfGrowthCurves_mapString = make(map[string]*TopStackOfGrowthCurve)
	for topstackofgrowthcurve := range stage.TopStackOfGrowthCurves {
		stage.TopStackOfGrowthCurves_mapString[topstackofgrowthcurve.Name] = topstackofgrowthcurve
	}

	stage.TopStartArcShapes_mapString = make(map[string]*TopStartArcShape)
	for topstartarcshape := range stage.TopStartArcShapes {
		stage.TopStartArcShapes_mapString[topstartarcshape.Name] = topstartarcshape
	}

	stage.TopStartArcShapeGrids_mapString = make(map[string]*TopStartArcShapeGrid)
	for topstartarcshapegrid := range stage.TopStartArcShapeGrids {
		stage.TopStartArcShapeGrids_mapString[topstartarcshapegrid.Name] = topstartarcshapegrid
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
