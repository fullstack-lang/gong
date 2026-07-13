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

	BottomEndArcShapeV2s                map[*BottomEndArcShapeV2]struct{}
	BottomEndArcShapeV2s_instance       map[*BottomEndArcShapeV2]*BottomEndArcShapeV2
	BottomEndArcShapeV2s_mapString      map[string]*BottomEndArcShapeV2
	BottomEndArcShapeV2Order            uint
	BottomEndArcShapeV2_stagedOrder     map[*BottomEndArcShapeV2]uint
	BottomEndArcShapeV2_orderStaged     map[uint]*BottomEndArcShapeV2
	BottomEndArcShapeV2s_reference      map[*BottomEndArcShapeV2]*BottomEndArcShapeV2
	BottomEndArcShapeV2s_referenceOrder map[*BottomEndArcShapeV2]uint

	// insertion point for slice of pointers maps
	OnAfterBottomEndArcShapeV2CreateCallback OnAfterCreateInterface[BottomEndArcShapeV2]
	OnAfterBottomEndArcShapeV2UpdateCallback OnAfterUpdateInterface[BottomEndArcShapeV2]
	OnAfterBottomEndArcShapeV2DeleteCallback OnAfterDeleteInterface[BottomEndArcShapeV2]
	OnAfterBottomEndArcShapeV2ReadCallback   OnAfterReadInterface[BottomEndArcShapeV2]

	BottomEndArcShapeV2Grids                map[*BottomEndArcShapeV2Grid]struct{}
	BottomEndArcShapeV2Grids_instance       map[*BottomEndArcShapeV2Grid]*BottomEndArcShapeV2Grid
	BottomEndArcShapeV2Grids_mapString      map[string]*BottomEndArcShapeV2Grid
	BottomEndArcShapeV2GridOrder            uint
	BottomEndArcShapeV2Grid_stagedOrder     map[*BottomEndArcShapeV2Grid]uint
	BottomEndArcShapeV2Grid_orderStaged     map[uint]*BottomEndArcShapeV2Grid
	BottomEndArcShapeV2Grids_reference      map[*BottomEndArcShapeV2Grid]*BottomEndArcShapeV2Grid
	BottomEndArcShapeV2Grids_referenceOrder map[*BottomEndArcShapeV2Grid]uint

	// insertion point for slice of pointers maps
	BottomEndArcShapeV2Grid_BottomEndArcShapesV2_reverseMap map[*BottomEndArcShapeV2]*BottomEndArcShapeV2Grid

	OnAfterBottomEndArcShapeV2GridCreateCallback OnAfterCreateInterface[BottomEndArcShapeV2Grid]
	OnAfterBottomEndArcShapeV2GridUpdateCallback OnAfterUpdateInterface[BottomEndArcShapeV2Grid]
	OnAfterBottomEndArcShapeV2GridDeleteCallback OnAfterDeleteInterface[BottomEndArcShapeV2Grid]
	OnAfterBottomEndArcShapeV2GridReadCallback   OnAfterReadInterface[BottomEndArcShapeV2Grid]

	BottomStackGrowthCurveEndArcShapeV2s                map[*BottomStackGrowthCurveEndArcShapeV2]struct{}
	BottomStackGrowthCurveEndArcShapeV2s_instance       map[*BottomStackGrowthCurveEndArcShapeV2]*BottomStackGrowthCurveEndArcShapeV2
	BottomStackGrowthCurveEndArcShapeV2s_mapString      map[string]*BottomStackGrowthCurveEndArcShapeV2
	BottomStackGrowthCurveEndArcShapeV2Order            uint
	BottomStackGrowthCurveEndArcShapeV2_stagedOrder     map[*BottomStackGrowthCurveEndArcShapeV2]uint
	BottomStackGrowthCurveEndArcShapeV2_orderStaged     map[uint]*BottomStackGrowthCurveEndArcShapeV2
	BottomStackGrowthCurveEndArcShapeV2s_reference      map[*BottomStackGrowthCurveEndArcShapeV2]*BottomStackGrowthCurveEndArcShapeV2
	BottomStackGrowthCurveEndArcShapeV2s_referenceOrder map[*BottomStackGrowthCurveEndArcShapeV2]uint

	// insertion point for slice of pointers maps
	OnAfterBottomStackGrowthCurveEndArcShapeV2CreateCallback OnAfterCreateInterface[BottomStackGrowthCurveEndArcShapeV2]
	OnAfterBottomStackGrowthCurveEndArcShapeV2UpdateCallback OnAfterUpdateInterface[BottomStackGrowthCurveEndArcShapeV2]
	OnAfterBottomStackGrowthCurveEndArcShapeV2DeleteCallback OnAfterDeleteInterface[BottomStackGrowthCurveEndArcShapeV2]
	OnAfterBottomStackGrowthCurveEndArcShapeV2ReadCallback   OnAfterReadInterface[BottomStackGrowthCurveEndArcShapeV2]

	BottomStackGrowthCurveStartArcShapeV2s                map[*BottomStackGrowthCurveStartArcShapeV2]struct{}
	BottomStackGrowthCurveStartArcShapeV2s_instance       map[*BottomStackGrowthCurveStartArcShapeV2]*BottomStackGrowthCurveStartArcShapeV2
	BottomStackGrowthCurveStartArcShapeV2s_mapString      map[string]*BottomStackGrowthCurveStartArcShapeV2
	BottomStackGrowthCurveStartArcShapeV2Order            uint
	BottomStackGrowthCurveStartArcShapeV2_stagedOrder     map[*BottomStackGrowthCurveStartArcShapeV2]uint
	BottomStackGrowthCurveStartArcShapeV2_orderStaged     map[uint]*BottomStackGrowthCurveStartArcShapeV2
	BottomStackGrowthCurveStartArcShapeV2s_reference      map[*BottomStackGrowthCurveStartArcShapeV2]*BottomStackGrowthCurveStartArcShapeV2
	BottomStackGrowthCurveStartArcShapeV2s_referenceOrder map[*BottomStackGrowthCurveStartArcShapeV2]uint

	// insertion point for slice of pointers maps
	OnAfterBottomStackGrowthCurveStartArcShapeV2CreateCallback OnAfterCreateInterface[BottomStackGrowthCurveStartArcShapeV2]
	OnAfterBottomStackGrowthCurveStartArcShapeV2UpdateCallback OnAfterUpdateInterface[BottomStackGrowthCurveStartArcShapeV2]
	OnAfterBottomStackGrowthCurveStartArcShapeV2DeleteCallback OnAfterDeleteInterface[BottomStackGrowthCurveStartArcShapeV2]
	OnAfterBottomStackGrowthCurveStartArcShapeV2ReadCallback   OnAfterReadInterface[BottomStackGrowthCurveStartArcShapeV2]

	BottomStackOfGrowthCurveV2s                map[*BottomStackOfGrowthCurveV2]struct{}
	BottomStackOfGrowthCurveV2s_instance       map[*BottomStackOfGrowthCurveV2]*BottomStackOfGrowthCurveV2
	BottomStackOfGrowthCurveV2s_mapString      map[string]*BottomStackOfGrowthCurveV2
	BottomStackOfGrowthCurveV2Order            uint
	BottomStackOfGrowthCurveV2_stagedOrder     map[*BottomStackOfGrowthCurveV2]uint
	BottomStackOfGrowthCurveV2_orderStaged     map[uint]*BottomStackOfGrowthCurveV2
	BottomStackOfGrowthCurveV2s_reference      map[*BottomStackOfGrowthCurveV2]*BottomStackOfGrowthCurveV2
	BottomStackOfGrowthCurveV2s_referenceOrder map[*BottomStackOfGrowthCurveV2]uint

	// insertion point for slice of pointers maps
	BottomStackOfGrowthCurveV2_BottomStackGrowthCurveStartArcShapeV2s_reverseMap map[*BottomStackGrowthCurveStartArcShapeV2]*BottomStackOfGrowthCurveV2

	BottomStackOfGrowthCurveV2_BottomStackGrowthCurveEndArcShapeV2s_reverseMap map[*BottomStackGrowthCurveEndArcShapeV2]*BottomStackOfGrowthCurveV2

	OnAfterBottomStackOfGrowthCurveV2CreateCallback OnAfterCreateInterface[BottomStackOfGrowthCurveV2]
	OnAfterBottomStackOfGrowthCurveV2UpdateCallback OnAfterUpdateInterface[BottomStackOfGrowthCurveV2]
	OnAfterBottomStackOfGrowthCurveV2DeleteCallback OnAfterDeleteInterface[BottomStackOfGrowthCurveV2]
	OnAfterBottomStackOfGrowthCurveV2ReadCallback   OnAfterReadInterface[BottomStackOfGrowthCurveV2]

	BottomStartArcShapeV2s                map[*BottomStartArcShapeV2]struct{}
	BottomStartArcShapeV2s_instance       map[*BottomStartArcShapeV2]*BottomStartArcShapeV2
	BottomStartArcShapeV2s_mapString      map[string]*BottomStartArcShapeV2
	BottomStartArcShapeV2Order            uint
	BottomStartArcShapeV2_stagedOrder     map[*BottomStartArcShapeV2]uint
	BottomStartArcShapeV2_orderStaged     map[uint]*BottomStartArcShapeV2
	BottomStartArcShapeV2s_reference      map[*BottomStartArcShapeV2]*BottomStartArcShapeV2
	BottomStartArcShapeV2s_referenceOrder map[*BottomStartArcShapeV2]uint

	// insertion point for slice of pointers maps
	OnAfterBottomStartArcShapeV2CreateCallback OnAfterCreateInterface[BottomStartArcShapeV2]
	OnAfterBottomStartArcShapeV2UpdateCallback OnAfterUpdateInterface[BottomStartArcShapeV2]
	OnAfterBottomStartArcShapeV2DeleteCallback OnAfterDeleteInterface[BottomStartArcShapeV2]
	OnAfterBottomStartArcShapeV2ReadCallback   OnAfterReadInterface[BottomStartArcShapeV2]

	BottomStartArcShapeV2Grids                map[*BottomStartArcShapeV2Grid]struct{}
	BottomStartArcShapeV2Grids_instance       map[*BottomStartArcShapeV2Grid]*BottomStartArcShapeV2Grid
	BottomStartArcShapeV2Grids_mapString      map[string]*BottomStartArcShapeV2Grid
	BottomStartArcShapeV2GridOrder            uint
	BottomStartArcShapeV2Grid_stagedOrder     map[*BottomStartArcShapeV2Grid]uint
	BottomStartArcShapeV2Grid_orderStaged     map[uint]*BottomStartArcShapeV2Grid
	BottomStartArcShapeV2Grids_reference      map[*BottomStartArcShapeV2Grid]*BottomStartArcShapeV2Grid
	BottomStartArcShapeV2Grids_referenceOrder map[*BottomStartArcShapeV2Grid]uint

	// insertion point for slice of pointers maps
	BottomStartArcShapeV2Grid_BottomStartArcShapesV2_reverseMap map[*BottomStartArcShapeV2]*BottomStartArcShapeV2Grid

	OnAfterBottomStartArcShapeV2GridCreateCallback OnAfterCreateInterface[BottomStartArcShapeV2Grid]
	OnAfterBottomStartArcShapeV2GridUpdateCallback OnAfterUpdateInterface[BottomStartArcShapeV2Grid]
	OnAfterBottomStartArcShapeV2GridDeleteCallback OnAfterDeleteInterface[BottomStartArcShapeV2Grid]
	OnAfterBottomStartArcShapeV2GridReadCallback   OnAfterReadInterface[BottomStartArcShapeV2Grid]

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

	EndArcShapeV2s                map[*EndArcShapeV2]struct{}
	EndArcShapeV2s_instance       map[*EndArcShapeV2]*EndArcShapeV2
	EndArcShapeV2s_mapString      map[string]*EndArcShapeV2
	EndArcShapeV2Order            uint
	EndArcShapeV2_stagedOrder     map[*EndArcShapeV2]uint
	EndArcShapeV2_orderStaged     map[uint]*EndArcShapeV2
	EndArcShapeV2s_reference      map[*EndArcShapeV2]*EndArcShapeV2
	EndArcShapeV2s_referenceOrder map[*EndArcShapeV2]uint

	// insertion point for slice of pointers maps
	OnAfterEndArcShapeV2CreateCallback OnAfterCreateInterface[EndArcShapeV2]
	OnAfterEndArcShapeV2UpdateCallback OnAfterUpdateInterface[EndArcShapeV2]
	OnAfterEndArcShapeV2DeleteCallback OnAfterDeleteInterface[EndArcShapeV2]
	OnAfterEndArcShapeV2ReadCallback   OnAfterReadInterface[EndArcShapeV2]

	EndArcShapeV2Grids                map[*EndArcShapeV2Grid]struct{}
	EndArcShapeV2Grids_instance       map[*EndArcShapeV2Grid]*EndArcShapeV2Grid
	EndArcShapeV2Grids_mapString      map[string]*EndArcShapeV2Grid
	EndArcShapeV2GridOrder            uint
	EndArcShapeV2Grid_stagedOrder     map[*EndArcShapeV2Grid]uint
	EndArcShapeV2Grid_orderStaged     map[uint]*EndArcShapeV2Grid
	EndArcShapeV2Grids_reference      map[*EndArcShapeV2Grid]*EndArcShapeV2Grid
	EndArcShapeV2Grids_referenceOrder map[*EndArcShapeV2Grid]uint

	// insertion point for slice of pointers maps
	EndArcShapeV2Grid_EndArcShapesV2_reverseMap map[*EndArcShapeV2]*EndArcShapeV2Grid

	OnAfterEndArcShapeV2GridCreateCallback OnAfterCreateInterface[EndArcShapeV2Grid]
	OnAfterEndArcShapeV2GridUpdateCallback OnAfterUpdateInterface[EndArcShapeV2Grid]
	OnAfterEndArcShapeV2GridDeleteCallback OnAfterDeleteInterface[EndArcShapeV2Grid]
	OnAfterEndArcShapeV2GridReadCallback   OnAfterReadInterface[EndArcShapeV2Grid]

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

	StackGrowthCurveEndArcShapeV2s                map[*StackGrowthCurveEndArcShapeV2]struct{}
	StackGrowthCurveEndArcShapeV2s_instance       map[*StackGrowthCurveEndArcShapeV2]*StackGrowthCurveEndArcShapeV2
	StackGrowthCurveEndArcShapeV2s_mapString      map[string]*StackGrowthCurveEndArcShapeV2
	StackGrowthCurveEndArcShapeV2Order            uint
	StackGrowthCurveEndArcShapeV2_stagedOrder     map[*StackGrowthCurveEndArcShapeV2]uint
	StackGrowthCurveEndArcShapeV2_orderStaged     map[uint]*StackGrowthCurveEndArcShapeV2
	StackGrowthCurveEndArcShapeV2s_reference      map[*StackGrowthCurveEndArcShapeV2]*StackGrowthCurveEndArcShapeV2
	StackGrowthCurveEndArcShapeV2s_referenceOrder map[*StackGrowthCurveEndArcShapeV2]uint

	// insertion point for slice of pointers maps
	OnAfterStackGrowthCurveEndArcShapeV2CreateCallback OnAfterCreateInterface[StackGrowthCurveEndArcShapeV2]
	OnAfterStackGrowthCurveEndArcShapeV2UpdateCallback OnAfterUpdateInterface[StackGrowthCurveEndArcShapeV2]
	OnAfterStackGrowthCurveEndArcShapeV2DeleteCallback OnAfterDeleteInterface[StackGrowthCurveEndArcShapeV2]
	OnAfterStackGrowthCurveEndArcShapeV2ReadCallback   OnAfterReadInterface[StackGrowthCurveEndArcShapeV2]

	StackGrowthCurveStartArcShapeV2s                map[*StackGrowthCurveStartArcShapeV2]struct{}
	StackGrowthCurveStartArcShapeV2s_instance       map[*StackGrowthCurveStartArcShapeV2]*StackGrowthCurveStartArcShapeV2
	StackGrowthCurveStartArcShapeV2s_mapString      map[string]*StackGrowthCurveStartArcShapeV2
	StackGrowthCurveStartArcShapeV2Order            uint
	StackGrowthCurveStartArcShapeV2_stagedOrder     map[*StackGrowthCurveStartArcShapeV2]uint
	StackGrowthCurveStartArcShapeV2_orderStaged     map[uint]*StackGrowthCurveStartArcShapeV2
	StackGrowthCurveStartArcShapeV2s_reference      map[*StackGrowthCurveStartArcShapeV2]*StackGrowthCurveStartArcShapeV2
	StackGrowthCurveStartArcShapeV2s_referenceOrder map[*StackGrowthCurveStartArcShapeV2]uint

	// insertion point for slice of pointers maps
	OnAfterStackGrowthCurveStartArcShapeV2CreateCallback OnAfterCreateInterface[StackGrowthCurveStartArcShapeV2]
	OnAfterStackGrowthCurveStartArcShapeV2UpdateCallback OnAfterUpdateInterface[StackGrowthCurveStartArcShapeV2]
	OnAfterStackGrowthCurveStartArcShapeV2DeleteCallback OnAfterDeleteInterface[StackGrowthCurveStartArcShapeV2]
	OnAfterStackGrowthCurveStartArcShapeV2ReadCallback   OnAfterReadInterface[StackGrowthCurveStartArcShapeV2]

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

	StackOfGrowthCurveV2s                map[*StackOfGrowthCurveV2]struct{}
	StackOfGrowthCurveV2s_instance       map[*StackOfGrowthCurveV2]*StackOfGrowthCurveV2
	StackOfGrowthCurveV2s_mapString      map[string]*StackOfGrowthCurveV2
	StackOfGrowthCurveV2Order            uint
	StackOfGrowthCurveV2_stagedOrder     map[*StackOfGrowthCurveV2]uint
	StackOfGrowthCurveV2_orderStaged     map[uint]*StackOfGrowthCurveV2
	StackOfGrowthCurveV2s_reference      map[*StackOfGrowthCurveV2]*StackOfGrowthCurveV2
	StackOfGrowthCurveV2s_referenceOrder map[*StackOfGrowthCurveV2]uint

	// insertion point for slice of pointers maps
	StackOfGrowthCurveV2_StackGrowthCurveStartArcShapeV2s_reverseMap map[*StackGrowthCurveStartArcShapeV2]*StackOfGrowthCurveV2

	StackOfGrowthCurveV2_StackGrowthCurveEndArcShapeV2s_reverseMap map[*StackGrowthCurveEndArcShapeV2]*StackOfGrowthCurveV2

	OnAfterStackOfGrowthCurveV2CreateCallback OnAfterCreateInterface[StackOfGrowthCurveV2]
	OnAfterStackOfGrowthCurveV2UpdateCallback OnAfterUpdateInterface[StackOfGrowthCurveV2]
	OnAfterStackOfGrowthCurveV2DeleteCallback OnAfterDeleteInterface[StackOfGrowthCurveV2]
	OnAfterStackOfGrowthCurveV2ReadCallback   OnAfterReadInterface[StackOfGrowthCurveV2]

	StartArcShapeV2s                map[*StartArcShapeV2]struct{}
	StartArcShapeV2s_instance       map[*StartArcShapeV2]*StartArcShapeV2
	StartArcShapeV2s_mapString      map[string]*StartArcShapeV2
	StartArcShapeV2Order            uint
	StartArcShapeV2_stagedOrder     map[*StartArcShapeV2]uint
	StartArcShapeV2_orderStaged     map[uint]*StartArcShapeV2
	StartArcShapeV2s_reference      map[*StartArcShapeV2]*StartArcShapeV2
	StartArcShapeV2s_referenceOrder map[*StartArcShapeV2]uint

	// insertion point for slice of pointers maps
	OnAfterStartArcShapeV2CreateCallback OnAfterCreateInterface[StartArcShapeV2]
	OnAfterStartArcShapeV2UpdateCallback OnAfterUpdateInterface[StartArcShapeV2]
	OnAfterStartArcShapeV2DeleteCallback OnAfterDeleteInterface[StartArcShapeV2]
	OnAfterStartArcShapeV2ReadCallback   OnAfterReadInterface[StartArcShapeV2]

	StartArcShapeV2Grids                map[*StartArcShapeV2Grid]struct{}
	StartArcShapeV2Grids_instance       map[*StartArcShapeV2Grid]*StartArcShapeV2Grid
	StartArcShapeV2Grids_mapString      map[string]*StartArcShapeV2Grid
	StartArcShapeV2GridOrder            uint
	StartArcShapeV2Grid_stagedOrder     map[*StartArcShapeV2Grid]uint
	StartArcShapeV2Grid_orderStaged     map[uint]*StartArcShapeV2Grid
	StartArcShapeV2Grids_reference      map[*StartArcShapeV2Grid]*StartArcShapeV2Grid
	StartArcShapeV2Grids_referenceOrder map[*StartArcShapeV2Grid]uint

	// insertion point for slice of pointers maps
	StartArcShapeV2Grid_StartArcShapesV2_reverseMap map[*StartArcShapeV2]*StartArcShapeV2Grid

	OnAfterStartArcShapeV2GridCreateCallback OnAfterCreateInterface[StartArcShapeV2Grid]
	OnAfterStartArcShapeV2GridUpdateCallback OnAfterUpdateInterface[StartArcShapeV2Grid]
	OnAfterStartArcShapeV2GridDeleteCallback OnAfterDeleteInterface[StartArcShapeV2Grid]
	OnAfterStartArcShapeV2GridReadCallback   OnAfterReadInterface[StartArcShapeV2Grid]

	TopEndArcShapeV2s                map[*TopEndArcShapeV2]struct{}
	TopEndArcShapeV2s_instance       map[*TopEndArcShapeV2]*TopEndArcShapeV2
	TopEndArcShapeV2s_mapString      map[string]*TopEndArcShapeV2
	TopEndArcShapeV2Order            uint
	TopEndArcShapeV2_stagedOrder     map[*TopEndArcShapeV2]uint
	TopEndArcShapeV2_orderStaged     map[uint]*TopEndArcShapeV2
	TopEndArcShapeV2s_reference      map[*TopEndArcShapeV2]*TopEndArcShapeV2
	TopEndArcShapeV2s_referenceOrder map[*TopEndArcShapeV2]uint

	// insertion point for slice of pointers maps
	OnAfterTopEndArcShapeV2CreateCallback OnAfterCreateInterface[TopEndArcShapeV2]
	OnAfterTopEndArcShapeV2UpdateCallback OnAfterUpdateInterface[TopEndArcShapeV2]
	OnAfterTopEndArcShapeV2DeleteCallback OnAfterDeleteInterface[TopEndArcShapeV2]
	OnAfterTopEndArcShapeV2ReadCallback   OnAfterReadInterface[TopEndArcShapeV2]

	TopEndArcShapeV2Grids                map[*TopEndArcShapeV2Grid]struct{}
	TopEndArcShapeV2Grids_instance       map[*TopEndArcShapeV2Grid]*TopEndArcShapeV2Grid
	TopEndArcShapeV2Grids_mapString      map[string]*TopEndArcShapeV2Grid
	TopEndArcShapeV2GridOrder            uint
	TopEndArcShapeV2Grid_stagedOrder     map[*TopEndArcShapeV2Grid]uint
	TopEndArcShapeV2Grid_orderStaged     map[uint]*TopEndArcShapeV2Grid
	TopEndArcShapeV2Grids_reference      map[*TopEndArcShapeV2Grid]*TopEndArcShapeV2Grid
	TopEndArcShapeV2Grids_referenceOrder map[*TopEndArcShapeV2Grid]uint

	// insertion point for slice of pointers maps
	TopEndArcShapeV2Grid_TopEndArcShapesV2_reverseMap map[*TopEndArcShapeV2]*TopEndArcShapeV2Grid

	OnAfterTopEndArcShapeV2GridCreateCallback OnAfterCreateInterface[TopEndArcShapeV2Grid]
	OnAfterTopEndArcShapeV2GridUpdateCallback OnAfterUpdateInterface[TopEndArcShapeV2Grid]
	OnAfterTopEndArcShapeV2GridDeleteCallback OnAfterDeleteInterface[TopEndArcShapeV2Grid]
	OnAfterTopEndArcShapeV2GridReadCallback   OnAfterReadInterface[TopEndArcShapeV2Grid]

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

	TopStackGrowthCurveEndArcShapeV2s                map[*TopStackGrowthCurveEndArcShapeV2]struct{}
	TopStackGrowthCurveEndArcShapeV2s_instance       map[*TopStackGrowthCurveEndArcShapeV2]*TopStackGrowthCurveEndArcShapeV2
	TopStackGrowthCurveEndArcShapeV2s_mapString      map[string]*TopStackGrowthCurveEndArcShapeV2
	TopStackGrowthCurveEndArcShapeV2Order            uint
	TopStackGrowthCurveEndArcShapeV2_stagedOrder     map[*TopStackGrowthCurveEndArcShapeV2]uint
	TopStackGrowthCurveEndArcShapeV2_orderStaged     map[uint]*TopStackGrowthCurveEndArcShapeV2
	TopStackGrowthCurveEndArcShapeV2s_reference      map[*TopStackGrowthCurveEndArcShapeV2]*TopStackGrowthCurveEndArcShapeV2
	TopStackGrowthCurveEndArcShapeV2s_referenceOrder map[*TopStackGrowthCurveEndArcShapeV2]uint

	// insertion point for slice of pointers maps
	OnAfterTopStackGrowthCurveEndArcShapeV2CreateCallback OnAfterCreateInterface[TopStackGrowthCurveEndArcShapeV2]
	OnAfterTopStackGrowthCurveEndArcShapeV2UpdateCallback OnAfterUpdateInterface[TopStackGrowthCurveEndArcShapeV2]
	OnAfterTopStackGrowthCurveEndArcShapeV2DeleteCallback OnAfterDeleteInterface[TopStackGrowthCurveEndArcShapeV2]
	OnAfterTopStackGrowthCurveEndArcShapeV2ReadCallback   OnAfterReadInterface[TopStackGrowthCurveEndArcShapeV2]

	TopStackGrowthCurveStartArcShapeV2s                map[*TopStackGrowthCurveStartArcShapeV2]struct{}
	TopStackGrowthCurveStartArcShapeV2s_instance       map[*TopStackGrowthCurveStartArcShapeV2]*TopStackGrowthCurveStartArcShapeV2
	TopStackGrowthCurveStartArcShapeV2s_mapString      map[string]*TopStackGrowthCurveStartArcShapeV2
	TopStackGrowthCurveStartArcShapeV2Order            uint
	TopStackGrowthCurveStartArcShapeV2_stagedOrder     map[*TopStackGrowthCurveStartArcShapeV2]uint
	TopStackGrowthCurveStartArcShapeV2_orderStaged     map[uint]*TopStackGrowthCurveStartArcShapeV2
	TopStackGrowthCurveStartArcShapeV2s_reference      map[*TopStackGrowthCurveStartArcShapeV2]*TopStackGrowthCurveStartArcShapeV2
	TopStackGrowthCurveStartArcShapeV2s_referenceOrder map[*TopStackGrowthCurveStartArcShapeV2]uint

	// insertion point for slice of pointers maps
	OnAfterTopStackGrowthCurveStartArcShapeV2CreateCallback OnAfterCreateInterface[TopStackGrowthCurveStartArcShapeV2]
	OnAfterTopStackGrowthCurveStartArcShapeV2UpdateCallback OnAfterUpdateInterface[TopStackGrowthCurveStartArcShapeV2]
	OnAfterTopStackGrowthCurveStartArcShapeV2DeleteCallback OnAfterDeleteInterface[TopStackGrowthCurveStartArcShapeV2]
	OnAfterTopStackGrowthCurveStartArcShapeV2ReadCallback   OnAfterReadInterface[TopStackGrowthCurveStartArcShapeV2]

	TopStackOfGrowthCurveV2s                map[*TopStackOfGrowthCurveV2]struct{}
	TopStackOfGrowthCurveV2s_instance       map[*TopStackOfGrowthCurveV2]*TopStackOfGrowthCurveV2
	TopStackOfGrowthCurveV2s_mapString      map[string]*TopStackOfGrowthCurveV2
	TopStackOfGrowthCurveV2Order            uint
	TopStackOfGrowthCurveV2_stagedOrder     map[*TopStackOfGrowthCurveV2]uint
	TopStackOfGrowthCurveV2_orderStaged     map[uint]*TopStackOfGrowthCurveV2
	TopStackOfGrowthCurveV2s_reference      map[*TopStackOfGrowthCurveV2]*TopStackOfGrowthCurveV2
	TopStackOfGrowthCurveV2s_referenceOrder map[*TopStackOfGrowthCurveV2]uint

	// insertion point for slice of pointers maps
	TopStackOfGrowthCurveV2_TopStackGrowthCurveStartArcShapeV2s_reverseMap map[*TopStackGrowthCurveStartArcShapeV2]*TopStackOfGrowthCurveV2

	TopStackOfGrowthCurveV2_TopStackGrowthCurveEndArcShapeV2s_reverseMap map[*TopStackGrowthCurveEndArcShapeV2]*TopStackOfGrowthCurveV2

	OnAfterTopStackOfGrowthCurveV2CreateCallback OnAfterCreateInterface[TopStackOfGrowthCurveV2]
	OnAfterTopStackOfGrowthCurveV2UpdateCallback OnAfterUpdateInterface[TopStackOfGrowthCurveV2]
	OnAfterTopStackOfGrowthCurveV2DeleteCallback OnAfterDeleteInterface[TopStackOfGrowthCurveV2]
	OnAfterTopStackOfGrowthCurveV2ReadCallback   OnAfterReadInterface[TopStackOfGrowthCurveV2]

	TopStartArcShapeV2s                map[*TopStartArcShapeV2]struct{}
	TopStartArcShapeV2s_instance       map[*TopStartArcShapeV2]*TopStartArcShapeV2
	TopStartArcShapeV2s_mapString      map[string]*TopStartArcShapeV2
	TopStartArcShapeV2Order            uint
	TopStartArcShapeV2_stagedOrder     map[*TopStartArcShapeV2]uint
	TopStartArcShapeV2_orderStaged     map[uint]*TopStartArcShapeV2
	TopStartArcShapeV2s_reference      map[*TopStartArcShapeV2]*TopStartArcShapeV2
	TopStartArcShapeV2s_referenceOrder map[*TopStartArcShapeV2]uint

	// insertion point for slice of pointers maps
	OnAfterTopStartArcShapeV2CreateCallback OnAfterCreateInterface[TopStartArcShapeV2]
	OnAfterTopStartArcShapeV2UpdateCallback OnAfterUpdateInterface[TopStartArcShapeV2]
	OnAfterTopStartArcShapeV2DeleteCallback OnAfterDeleteInterface[TopStartArcShapeV2]
	OnAfterTopStartArcShapeV2ReadCallback   OnAfterReadInterface[TopStartArcShapeV2]

	TopStartArcShapeV2Grids                map[*TopStartArcShapeV2Grid]struct{}
	TopStartArcShapeV2Grids_instance       map[*TopStartArcShapeV2Grid]*TopStartArcShapeV2Grid
	TopStartArcShapeV2Grids_mapString      map[string]*TopStartArcShapeV2Grid
	TopStartArcShapeV2GridOrder            uint
	TopStartArcShapeV2Grid_stagedOrder     map[*TopStartArcShapeV2Grid]uint
	TopStartArcShapeV2Grid_orderStaged     map[uint]*TopStartArcShapeV2Grid
	TopStartArcShapeV2Grids_reference      map[*TopStartArcShapeV2Grid]*TopStartArcShapeV2Grid
	TopStartArcShapeV2Grids_referenceOrder map[*TopStartArcShapeV2Grid]uint

	// insertion point for slice of pointers maps
	TopStartArcShapeV2Grid_TopStartArcShapesV2_reverseMap map[*TopStartArcShapeV2]*TopStartArcShapeV2Grid

	OnAfterTopStartArcShapeV2GridCreateCallback OnAfterCreateInterface[TopStartArcShapeV2Grid]
	OnAfterTopStartArcShapeV2GridUpdateCallback OnAfterUpdateInterface[TopStartArcShapeV2Grid]
	OnAfterTopStartArcShapeV2GridDeleteCallback OnAfterDeleteInterface[TopStartArcShapeV2Grid]
	OnAfterTopStartArcShapeV2GridReadCallback   OnAfterReadInterface[TopStartArcShapeV2Grid]

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

	stage.BottomEndArcShapeV2s_reference = make(map[*BottomEndArcShapeV2]*BottomEndArcShapeV2)
	stage.BottomEndArcShapeV2s_instance = make(map[*BottomEndArcShapeV2]*BottomEndArcShapeV2)
	stage.BottomEndArcShapeV2s_referenceOrder = make(map[*BottomEndArcShapeV2]uint)

	stage.BottomEndArcShapeV2Grids_reference = make(map[*BottomEndArcShapeV2Grid]*BottomEndArcShapeV2Grid)
	stage.BottomEndArcShapeV2Grids_instance = make(map[*BottomEndArcShapeV2Grid]*BottomEndArcShapeV2Grid)
	stage.BottomEndArcShapeV2Grids_referenceOrder = make(map[*BottomEndArcShapeV2Grid]uint)

	stage.BottomStackGrowthCurveEndArcShapeV2s_reference = make(map[*BottomStackGrowthCurveEndArcShapeV2]*BottomStackGrowthCurveEndArcShapeV2)
	stage.BottomStackGrowthCurveEndArcShapeV2s_instance = make(map[*BottomStackGrowthCurveEndArcShapeV2]*BottomStackGrowthCurveEndArcShapeV2)
	stage.BottomStackGrowthCurveEndArcShapeV2s_referenceOrder = make(map[*BottomStackGrowthCurveEndArcShapeV2]uint)

	stage.BottomStackGrowthCurveStartArcShapeV2s_reference = make(map[*BottomStackGrowthCurveStartArcShapeV2]*BottomStackGrowthCurveStartArcShapeV2)
	stage.BottomStackGrowthCurveStartArcShapeV2s_instance = make(map[*BottomStackGrowthCurveStartArcShapeV2]*BottomStackGrowthCurveStartArcShapeV2)
	stage.BottomStackGrowthCurveStartArcShapeV2s_referenceOrder = make(map[*BottomStackGrowthCurveStartArcShapeV2]uint)

	stage.BottomStackOfGrowthCurveV2s_reference = make(map[*BottomStackOfGrowthCurveV2]*BottomStackOfGrowthCurveV2)
	stage.BottomStackOfGrowthCurveV2s_instance = make(map[*BottomStackOfGrowthCurveV2]*BottomStackOfGrowthCurveV2)
	stage.BottomStackOfGrowthCurveV2s_referenceOrder = make(map[*BottomStackOfGrowthCurveV2]uint)

	stage.BottomStartArcShapeV2s_reference = make(map[*BottomStartArcShapeV2]*BottomStartArcShapeV2)
	stage.BottomStartArcShapeV2s_instance = make(map[*BottomStartArcShapeV2]*BottomStartArcShapeV2)
	stage.BottomStartArcShapeV2s_referenceOrder = make(map[*BottomStartArcShapeV2]uint)

	stage.BottomStartArcShapeV2Grids_reference = make(map[*BottomStartArcShapeV2Grid]*BottomStartArcShapeV2Grid)
	stage.BottomStartArcShapeV2Grids_instance = make(map[*BottomStartArcShapeV2Grid]*BottomStartArcShapeV2Grid)
	stage.BottomStartArcShapeV2Grids_referenceOrder = make(map[*BottomStartArcShapeV2Grid]uint)

	stage.CircleGridShapes_reference = make(map[*CircleGridShape]*CircleGridShape)
	stage.CircleGridShapes_instance = make(map[*CircleGridShape]*CircleGridShape)
	stage.CircleGridShapes_referenceOrder = make(map[*CircleGridShape]uint)

	stage.EndArcShapeV2s_reference = make(map[*EndArcShapeV2]*EndArcShapeV2)
	stage.EndArcShapeV2s_instance = make(map[*EndArcShapeV2]*EndArcShapeV2)
	stage.EndArcShapeV2s_referenceOrder = make(map[*EndArcShapeV2]uint)

	stage.EndArcShapeV2Grids_reference = make(map[*EndArcShapeV2Grid]*EndArcShapeV2Grid)
	stage.EndArcShapeV2Grids_instance = make(map[*EndArcShapeV2Grid]*EndArcShapeV2Grid)
	stage.EndArcShapeV2Grids_referenceOrder = make(map[*EndArcShapeV2Grid]uint)

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

	stage.StackGrowthCurveBezierShapes_reference = make(map[*StackGrowthCurveBezierShape]*StackGrowthCurveBezierShape)
	stage.StackGrowthCurveBezierShapes_instance = make(map[*StackGrowthCurveBezierShape]*StackGrowthCurveBezierShape)
	stage.StackGrowthCurveBezierShapes_referenceOrder = make(map[*StackGrowthCurveBezierShape]uint)

	stage.StackGrowthCurveEndArcShapeV2s_reference = make(map[*StackGrowthCurveEndArcShapeV2]*StackGrowthCurveEndArcShapeV2)
	stage.StackGrowthCurveEndArcShapeV2s_instance = make(map[*StackGrowthCurveEndArcShapeV2]*StackGrowthCurveEndArcShapeV2)
	stage.StackGrowthCurveEndArcShapeV2s_referenceOrder = make(map[*StackGrowthCurveEndArcShapeV2]uint)

	stage.StackGrowthCurveStartArcShapeV2s_reference = make(map[*StackGrowthCurveStartArcShapeV2]*StackGrowthCurveStartArcShapeV2)
	stage.StackGrowthCurveStartArcShapeV2s_instance = make(map[*StackGrowthCurveStartArcShapeV2]*StackGrowthCurveStartArcShapeV2)
	stage.StackGrowthCurveStartArcShapeV2s_referenceOrder = make(map[*StackGrowthCurveStartArcShapeV2]uint)

	stage.StackOfGrowthCurves_reference = make(map[*StackOfGrowthCurve]*StackOfGrowthCurve)
	stage.StackOfGrowthCurves_instance = make(map[*StackOfGrowthCurve]*StackOfGrowthCurve)
	stage.StackOfGrowthCurves_referenceOrder = make(map[*StackOfGrowthCurve]uint)

	stage.StackOfGrowthCurveV2s_reference = make(map[*StackOfGrowthCurveV2]*StackOfGrowthCurveV2)
	stage.StackOfGrowthCurveV2s_instance = make(map[*StackOfGrowthCurveV2]*StackOfGrowthCurveV2)
	stage.StackOfGrowthCurveV2s_referenceOrder = make(map[*StackOfGrowthCurveV2]uint)

	stage.StartArcShapeV2s_reference = make(map[*StartArcShapeV2]*StartArcShapeV2)
	stage.StartArcShapeV2s_instance = make(map[*StartArcShapeV2]*StartArcShapeV2)
	stage.StartArcShapeV2s_referenceOrder = make(map[*StartArcShapeV2]uint)

	stage.StartArcShapeV2Grids_reference = make(map[*StartArcShapeV2Grid]*StartArcShapeV2Grid)
	stage.StartArcShapeV2Grids_instance = make(map[*StartArcShapeV2Grid]*StartArcShapeV2Grid)
	stage.StartArcShapeV2Grids_referenceOrder = make(map[*StartArcShapeV2Grid]uint)

	stage.TopEndArcShapeV2s_reference = make(map[*TopEndArcShapeV2]*TopEndArcShapeV2)
	stage.TopEndArcShapeV2s_instance = make(map[*TopEndArcShapeV2]*TopEndArcShapeV2)
	stage.TopEndArcShapeV2s_referenceOrder = make(map[*TopEndArcShapeV2]uint)

	stage.TopEndArcShapeV2Grids_reference = make(map[*TopEndArcShapeV2Grid]*TopEndArcShapeV2Grid)
	stage.TopEndArcShapeV2Grids_instance = make(map[*TopEndArcShapeV2Grid]*TopEndArcShapeV2Grid)
	stage.TopEndArcShapeV2Grids_referenceOrder = make(map[*TopEndArcShapeV2Grid]uint)

	stage.TopGrowthCurve2Ds_reference = make(map[*TopGrowthCurve2D]*TopGrowthCurve2D)
	stage.TopGrowthCurve2Ds_instance = make(map[*TopGrowthCurve2D]*TopGrowthCurve2D)
	stage.TopGrowthCurve2Ds_referenceOrder = make(map[*TopGrowthCurve2D]uint)

	stage.TopStackGrowthCurveEndArcShapeV2s_reference = make(map[*TopStackGrowthCurveEndArcShapeV2]*TopStackGrowthCurveEndArcShapeV2)
	stage.TopStackGrowthCurveEndArcShapeV2s_instance = make(map[*TopStackGrowthCurveEndArcShapeV2]*TopStackGrowthCurveEndArcShapeV2)
	stage.TopStackGrowthCurveEndArcShapeV2s_referenceOrder = make(map[*TopStackGrowthCurveEndArcShapeV2]uint)

	stage.TopStackGrowthCurveStartArcShapeV2s_reference = make(map[*TopStackGrowthCurveStartArcShapeV2]*TopStackGrowthCurveStartArcShapeV2)
	stage.TopStackGrowthCurveStartArcShapeV2s_instance = make(map[*TopStackGrowthCurveStartArcShapeV2]*TopStackGrowthCurveStartArcShapeV2)
	stage.TopStackGrowthCurveStartArcShapeV2s_referenceOrder = make(map[*TopStackGrowthCurveStartArcShapeV2]uint)

	stage.TopStackOfGrowthCurveV2s_reference = make(map[*TopStackOfGrowthCurveV2]*TopStackOfGrowthCurveV2)
	stage.TopStackOfGrowthCurveV2s_instance = make(map[*TopStackOfGrowthCurveV2]*TopStackOfGrowthCurveV2)
	stage.TopStackOfGrowthCurveV2s_referenceOrder = make(map[*TopStackOfGrowthCurveV2]uint)

	stage.TopStartArcShapeV2s_reference = make(map[*TopStartArcShapeV2]*TopStartArcShapeV2)
	stage.TopStartArcShapeV2s_instance = make(map[*TopStartArcShapeV2]*TopStartArcShapeV2)
	stage.TopStartArcShapeV2s_referenceOrder = make(map[*TopStartArcShapeV2]uint)

	stage.TopStartArcShapeV2Grids_reference = make(map[*TopStartArcShapeV2Grid]*TopStartArcShapeV2Grid)
	stage.TopStartArcShapeV2Grids_instance = make(map[*TopStartArcShapeV2Grid]*TopStartArcShapeV2Grid)
	stage.TopStartArcShapeV2Grids_referenceOrder = make(map[*TopStartArcShapeV2Grid]uint)

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

	var maxBottomEndArcShapeV2Order uint
	var foundBottomEndArcShapeV2 bool
	for _, order := range stage.BottomEndArcShapeV2_stagedOrder {
		if !foundBottomEndArcShapeV2 || order > maxBottomEndArcShapeV2Order {
			maxBottomEndArcShapeV2Order = order
			foundBottomEndArcShapeV2 = true
		}
	}
	if foundBottomEndArcShapeV2 {
		stage.BottomEndArcShapeV2Order = maxBottomEndArcShapeV2Order + 1
	} else {
		stage.BottomEndArcShapeV2Order = 0
	}

	var maxBottomEndArcShapeV2GridOrder uint
	var foundBottomEndArcShapeV2Grid bool
	for _, order := range stage.BottomEndArcShapeV2Grid_stagedOrder {
		if !foundBottomEndArcShapeV2Grid || order > maxBottomEndArcShapeV2GridOrder {
			maxBottomEndArcShapeV2GridOrder = order
			foundBottomEndArcShapeV2Grid = true
		}
	}
	if foundBottomEndArcShapeV2Grid {
		stage.BottomEndArcShapeV2GridOrder = maxBottomEndArcShapeV2GridOrder + 1
	} else {
		stage.BottomEndArcShapeV2GridOrder = 0
	}

	var maxBottomStackGrowthCurveEndArcShapeV2Order uint
	var foundBottomStackGrowthCurveEndArcShapeV2 bool
	for _, order := range stage.BottomStackGrowthCurveEndArcShapeV2_stagedOrder {
		if !foundBottomStackGrowthCurveEndArcShapeV2 || order > maxBottomStackGrowthCurveEndArcShapeV2Order {
			maxBottomStackGrowthCurveEndArcShapeV2Order = order
			foundBottomStackGrowthCurveEndArcShapeV2 = true
		}
	}
	if foundBottomStackGrowthCurveEndArcShapeV2 {
		stage.BottomStackGrowthCurveEndArcShapeV2Order = maxBottomStackGrowthCurveEndArcShapeV2Order + 1
	} else {
		stage.BottomStackGrowthCurveEndArcShapeV2Order = 0
	}

	var maxBottomStackGrowthCurveStartArcShapeV2Order uint
	var foundBottomStackGrowthCurveStartArcShapeV2 bool
	for _, order := range stage.BottomStackGrowthCurveStartArcShapeV2_stagedOrder {
		if !foundBottomStackGrowthCurveStartArcShapeV2 || order > maxBottomStackGrowthCurveStartArcShapeV2Order {
			maxBottomStackGrowthCurveStartArcShapeV2Order = order
			foundBottomStackGrowthCurveStartArcShapeV2 = true
		}
	}
	if foundBottomStackGrowthCurveStartArcShapeV2 {
		stage.BottomStackGrowthCurveStartArcShapeV2Order = maxBottomStackGrowthCurveStartArcShapeV2Order + 1
	} else {
		stage.BottomStackGrowthCurveStartArcShapeV2Order = 0
	}

	var maxBottomStackOfGrowthCurveV2Order uint
	var foundBottomStackOfGrowthCurveV2 bool
	for _, order := range stage.BottomStackOfGrowthCurveV2_stagedOrder {
		if !foundBottomStackOfGrowthCurveV2 || order > maxBottomStackOfGrowthCurveV2Order {
			maxBottomStackOfGrowthCurveV2Order = order
			foundBottomStackOfGrowthCurveV2 = true
		}
	}
	if foundBottomStackOfGrowthCurveV2 {
		stage.BottomStackOfGrowthCurveV2Order = maxBottomStackOfGrowthCurveV2Order + 1
	} else {
		stage.BottomStackOfGrowthCurveV2Order = 0
	}

	var maxBottomStartArcShapeV2Order uint
	var foundBottomStartArcShapeV2 bool
	for _, order := range stage.BottomStartArcShapeV2_stagedOrder {
		if !foundBottomStartArcShapeV2 || order > maxBottomStartArcShapeV2Order {
			maxBottomStartArcShapeV2Order = order
			foundBottomStartArcShapeV2 = true
		}
	}
	if foundBottomStartArcShapeV2 {
		stage.BottomStartArcShapeV2Order = maxBottomStartArcShapeV2Order + 1
	} else {
		stage.BottomStartArcShapeV2Order = 0
	}

	var maxBottomStartArcShapeV2GridOrder uint
	var foundBottomStartArcShapeV2Grid bool
	for _, order := range stage.BottomStartArcShapeV2Grid_stagedOrder {
		if !foundBottomStartArcShapeV2Grid || order > maxBottomStartArcShapeV2GridOrder {
			maxBottomStartArcShapeV2GridOrder = order
			foundBottomStartArcShapeV2Grid = true
		}
	}
	if foundBottomStartArcShapeV2Grid {
		stage.BottomStartArcShapeV2GridOrder = maxBottomStartArcShapeV2GridOrder + 1
	} else {
		stage.BottomStartArcShapeV2GridOrder = 0
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

	var maxEndArcShapeV2Order uint
	var foundEndArcShapeV2 bool
	for _, order := range stage.EndArcShapeV2_stagedOrder {
		if !foundEndArcShapeV2 || order > maxEndArcShapeV2Order {
			maxEndArcShapeV2Order = order
			foundEndArcShapeV2 = true
		}
	}
	if foundEndArcShapeV2 {
		stage.EndArcShapeV2Order = maxEndArcShapeV2Order + 1
	} else {
		stage.EndArcShapeV2Order = 0
	}

	var maxEndArcShapeV2GridOrder uint
	var foundEndArcShapeV2Grid bool
	for _, order := range stage.EndArcShapeV2Grid_stagedOrder {
		if !foundEndArcShapeV2Grid || order > maxEndArcShapeV2GridOrder {
			maxEndArcShapeV2GridOrder = order
			foundEndArcShapeV2Grid = true
		}
	}
	if foundEndArcShapeV2Grid {
		stage.EndArcShapeV2GridOrder = maxEndArcShapeV2GridOrder + 1
	} else {
		stage.EndArcShapeV2GridOrder = 0
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

	var maxStackGrowthCurveEndArcShapeV2Order uint
	var foundStackGrowthCurveEndArcShapeV2 bool
	for _, order := range stage.StackGrowthCurveEndArcShapeV2_stagedOrder {
		if !foundStackGrowthCurveEndArcShapeV2 || order > maxStackGrowthCurveEndArcShapeV2Order {
			maxStackGrowthCurveEndArcShapeV2Order = order
			foundStackGrowthCurveEndArcShapeV2 = true
		}
	}
	if foundStackGrowthCurveEndArcShapeV2 {
		stage.StackGrowthCurveEndArcShapeV2Order = maxStackGrowthCurveEndArcShapeV2Order + 1
	} else {
		stage.StackGrowthCurveEndArcShapeV2Order = 0
	}

	var maxStackGrowthCurveStartArcShapeV2Order uint
	var foundStackGrowthCurveStartArcShapeV2 bool
	for _, order := range stage.StackGrowthCurveStartArcShapeV2_stagedOrder {
		if !foundStackGrowthCurveStartArcShapeV2 || order > maxStackGrowthCurveStartArcShapeV2Order {
			maxStackGrowthCurveStartArcShapeV2Order = order
			foundStackGrowthCurveStartArcShapeV2 = true
		}
	}
	if foundStackGrowthCurveStartArcShapeV2 {
		stage.StackGrowthCurveStartArcShapeV2Order = maxStackGrowthCurveStartArcShapeV2Order + 1
	} else {
		stage.StackGrowthCurveStartArcShapeV2Order = 0
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

	var maxStackOfGrowthCurveV2Order uint
	var foundStackOfGrowthCurveV2 bool
	for _, order := range stage.StackOfGrowthCurveV2_stagedOrder {
		if !foundStackOfGrowthCurveV2 || order > maxStackOfGrowthCurveV2Order {
			maxStackOfGrowthCurveV2Order = order
			foundStackOfGrowthCurveV2 = true
		}
	}
	if foundStackOfGrowthCurveV2 {
		stage.StackOfGrowthCurveV2Order = maxStackOfGrowthCurveV2Order + 1
	} else {
		stage.StackOfGrowthCurveV2Order = 0
	}

	var maxStartArcShapeV2Order uint
	var foundStartArcShapeV2 bool
	for _, order := range stage.StartArcShapeV2_stagedOrder {
		if !foundStartArcShapeV2 || order > maxStartArcShapeV2Order {
			maxStartArcShapeV2Order = order
			foundStartArcShapeV2 = true
		}
	}
	if foundStartArcShapeV2 {
		stage.StartArcShapeV2Order = maxStartArcShapeV2Order + 1
	} else {
		stage.StartArcShapeV2Order = 0
	}

	var maxStartArcShapeV2GridOrder uint
	var foundStartArcShapeV2Grid bool
	for _, order := range stage.StartArcShapeV2Grid_stagedOrder {
		if !foundStartArcShapeV2Grid || order > maxStartArcShapeV2GridOrder {
			maxStartArcShapeV2GridOrder = order
			foundStartArcShapeV2Grid = true
		}
	}
	if foundStartArcShapeV2Grid {
		stage.StartArcShapeV2GridOrder = maxStartArcShapeV2GridOrder + 1
	} else {
		stage.StartArcShapeV2GridOrder = 0
	}

	var maxTopEndArcShapeV2Order uint
	var foundTopEndArcShapeV2 bool
	for _, order := range stage.TopEndArcShapeV2_stagedOrder {
		if !foundTopEndArcShapeV2 || order > maxTopEndArcShapeV2Order {
			maxTopEndArcShapeV2Order = order
			foundTopEndArcShapeV2 = true
		}
	}
	if foundTopEndArcShapeV2 {
		stage.TopEndArcShapeV2Order = maxTopEndArcShapeV2Order + 1
	} else {
		stage.TopEndArcShapeV2Order = 0
	}

	var maxTopEndArcShapeV2GridOrder uint
	var foundTopEndArcShapeV2Grid bool
	for _, order := range stage.TopEndArcShapeV2Grid_stagedOrder {
		if !foundTopEndArcShapeV2Grid || order > maxTopEndArcShapeV2GridOrder {
			maxTopEndArcShapeV2GridOrder = order
			foundTopEndArcShapeV2Grid = true
		}
	}
	if foundTopEndArcShapeV2Grid {
		stage.TopEndArcShapeV2GridOrder = maxTopEndArcShapeV2GridOrder + 1
	} else {
		stage.TopEndArcShapeV2GridOrder = 0
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

	var maxTopStackGrowthCurveEndArcShapeV2Order uint
	var foundTopStackGrowthCurveEndArcShapeV2 bool
	for _, order := range stage.TopStackGrowthCurveEndArcShapeV2_stagedOrder {
		if !foundTopStackGrowthCurveEndArcShapeV2 || order > maxTopStackGrowthCurveEndArcShapeV2Order {
			maxTopStackGrowthCurveEndArcShapeV2Order = order
			foundTopStackGrowthCurveEndArcShapeV2 = true
		}
	}
	if foundTopStackGrowthCurveEndArcShapeV2 {
		stage.TopStackGrowthCurveEndArcShapeV2Order = maxTopStackGrowthCurveEndArcShapeV2Order + 1
	} else {
		stage.TopStackGrowthCurveEndArcShapeV2Order = 0
	}

	var maxTopStackGrowthCurveStartArcShapeV2Order uint
	var foundTopStackGrowthCurveStartArcShapeV2 bool
	for _, order := range stage.TopStackGrowthCurveStartArcShapeV2_stagedOrder {
		if !foundTopStackGrowthCurveStartArcShapeV2 || order > maxTopStackGrowthCurveStartArcShapeV2Order {
			maxTopStackGrowthCurveStartArcShapeV2Order = order
			foundTopStackGrowthCurveStartArcShapeV2 = true
		}
	}
	if foundTopStackGrowthCurveStartArcShapeV2 {
		stage.TopStackGrowthCurveStartArcShapeV2Order = maxTopStackGrowthCurveStartArcShapeV2Order + 1
	} else {
		stage.TopStackGrowthCurveStartArcShapeV2Order = 0
	}

	var maxTopStackOfGrowthCurveV2Order uint
	var foundTopStackOfGrowthCurveV2 bool
	for _, order := range stage.TopStackOfGrowthCurveV2_stagedOrder {
		if !foundTopStackOfGrowthCurveV2 || order > maxTopStackOfGrowthCurveV2Order {
			maxTopStackOfGrowthCurveV2Order = order
			foundTopStackOfGrowthCurveV2 = true
		}
	}
	if foundTopStackOfGrowthCurveV2 {
		stage.TopStackOfGrowthCurveV2Order = maxTopStackOfGrowthCurveV2Order + 1
	} else {
		stage.TopStackOfGrowthCurveV2Order = 0
	}

	var maxTopStartArcShapeV2Order uint
	var foundTopStartArcShapeV2 bool
	for _, order := range stage.TopStartArcShapeV2_stagedOrder {
		if !foundTopStartArcShapeV2 || order > maxTopStartArcShapeV2Order {
			maxTopStartArcShapeV2Order = order
			foundTopStartArcShapeV2 = true
		}
	}
	if foundTopStartArcShapeV2 {
		stage.TopStartArcShapeV2Order = maxTopStartArcShapeV2Order + 1
	} else {
		stage.TopStartArcShapeV2Order = 0
	}

	var maxTopStartArcShapeV2GridOrder uint
	var foundTopStartArcShapeV2Grid bool
	for _, order := range stage.TopStartArcShapeV2Grid_stagedOrder {
		if !foundTopStartArcShapeV2Grid || order > maxTopStartArcShapeV2GridOrder {
			maxTopStartArcShapeV2GridOrder = order
			foundTopStartArcShapeV2Grid = true
		}
	}
	if foundTopStartArcShapeV2Grid {
		stage.TopStartArcShapeV2GridOrder = maxTopStartArcShapeV2GridOrder + 1
	} else {
		stage.TopStartArcShapeV2GridOrder = 0
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
	case *BottomEndArcShapeV2:
		tmp := GetStructInstancesByOrder(stage.BottomEndArcShapeV2s, stage.BottomEndArcShapeV2_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *BottomEndArcShapeV2 implements.
			res = append(res, any(v).(T))
		}
		return res
	case *BottomEndArcShapeV2Grid:
		tmp := GetStructInstancesByOrder(stage.BottomEndArcShapeV2Grids, stage.BottomEndArcShapeV2Grid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *BottomEndArcShapeV2Grid implements.
			res = append(res, any(v).(T))
		}
		return res
	case *BottomStackGrowthCurveEndArcShapeV2:
		tmp := GetStructInstancesByOrder(stage.BottomStackGrowthCurveEndArcShapeV2s, stage.BottomStackGrowthCurveEndArcShapeV2_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *BottomStackGrowthCurveEndArcShapeV2 implements.
			res = append(res, any(v).(T))
		}
		return res
	case *BottomStackGrowthCurveStartArcShapeV2:
		tmp := GetStructInstancesByOrder(stage.BottomStackGrowthCurveStartArcShapeV2s, stage.BottomStackGrowthCurveStartArcShapeV2_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *BottomStackGrowthCurveStartArcShapeV2 implements.
			res = append(res, any(v).(T))
		}
		return res
	case *BottomStackOfGrowthCurveV2:
		tmp := GetStructInstancesByOrder(stage.BottomStackOfGrowthCurveV2s, stage.BottomStackOfGrowthCurveV2_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *BottomStackOfGrowthCurveV2 implements.
			res = append(res, any(v).(T))
		}
		return res
	case *BottomStartArcShapeV2:
		tmp := GetStructInstancesByOrder(stage.BottomStartArcShapeV2s, stage.BottomStartArcShapeV2_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *BottomStartArcShapeV2 implements.
			res = append(res, any(v).(T))
		}
		return res
	case *BottomStartArcShapeV2Grid:
		tmp := GetStructInstancesByOrder(stage.BottomStartArcShapeV2Grids, stage.BottomStartArcShapeV2Grid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *BottomStartArcShapeV2Grid implements.
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
	case *EndArcShapeV2:
		tmp := GetStructInstancesByOrder(stage.EndArcShapeV2s, stage.EndArcShapeV2_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *EndArcShapeV2 implements.
			res = append(res, any(v).(T))
		}
		return res
	case *EndArcShapeV2Grid:
		tmp := GetStructInstancesByOrder(stage.EndArcShapeV2Grids, stage.EndArcShapeV2Grid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *EndArcShapeV2Grid implements.
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
	case *StackGrowthCurveEndArcShapeV2:
		tmp := GetStructInstancesByOrder(stage.StackGrowthCurveEndArcShapeV2s, stage.StackGrowthCurveEndArcShapeV2_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *StackGrowthCurveEndArcShapeV2 implements.
			res = append(res, any(v).(T))
		}
		return res
	case *StackGrowthCurveStartArcShapeV2:
		tmp := GetStructInstancesByOrder(stage.StackGrowthCurveStartArcShapeV2s, stage.StackGrowthCurveStartArcShapeV2_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *StackGrowthCurveStartArcShapeV2 implements.
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
	case *StackOfGrowthCurveV2:
		tmp := GetStructInstancesByOrder(stage.StackOfGrowthCurveV2s, stage.StackOfGrowthCurveV2_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *StackOfGrowthCurveV2 implements.
			res = append(res, any(v).(T))
		}
		return res
	case *StartArcShapeV2:
		tmp := GetStructInstancesByOrder(stage.StartArcShapeV2s, stage.StartArcShapeV2_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *StartArcShapeV2 implements.
			res = append(res, any(v).(T))
		}
		return res
	case *StartArcShapeV2Grid:
		tmp := GetStructInstancesByOrder(stage.StartArcShapeV2Grids, stage.StartArcShapeV2Grid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *StartArcShapeV2Grid implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TopEndArcShapeV2:
		tmp := GetStructInstancesByOrder(stage.TopEndArcShapeV2s, stage.TopEndArcShapeV2_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TopEndArcShapeV2 implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TopEndArcShapeV2Grid:
		tmp := GetStructInstancesByOrder(stage.TopEndArcShapeV2Grids, stage.TopEndArcShapeV2Grid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TopEndArcShapeV2Grid implements.
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
	case *TopStackGrowthCurveEndArcShapeV2:
		tmp := GetStructInstancesByOrder(stage.TopStackGrowthCurveEndArcShapeV2s, stage.TopStackGrowthCurveEndArcShapeV2_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TopStackGrowthCurveEndArcShapeV2 implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TopStackGrowthCurveStartArcShapeV2:
		tmp := GetStructInstancesByOrder(stage.TopStackGrowthCurveStartArcShapeV2s, stage.TopStackGrowthCurveStartArcShapeV2_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TopStackGrowthCurveStartArcShapeV2 implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TopStackOfGrowthCurveV2:
		tmp := GetStructInstancesByOrder(stage.TopStackOfGrowthCurveV2s, stage.TopStackOfGrowthCurveV2_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TopStackOfGrowthCurveV2 implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TopStartArcShapeV2:
		tmp := GetStructInstancesByOrder(stage.TopStartArcShapeV2s, stage.TopStartArcShapeV2_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TopStartArcShapeV2 implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TopStartArcShapeV2Grid:
		tmp := GetStructInstancesByOrder(stage.TopStartArcShapeV2Grids, stage.TopStartArcShapeV2Grid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TopStartArcShapeV2Grid implements.
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
	case "BottomEndArcShapeV2":
		res = GetNamedStructInstances(stage.BottomEndArcShapeV2s, stage.BottomEndArcShapeV2_stagedOrder)
	case "BottomEndArcShapeV2Grid":
		res = GetNamedStructInstances(stage.BottomEndArcShapeV2Grids, stage.BottomEndArcShapeV2Grid_stagedOrder)
	case "BottomStackGrowthCurveEndArcShapeV2":
		res = GetNamedStructInstances(stage.BottomStackGrowthCurveEndArcShapeV2s, stage.BottomStackGrowthCurveEndArcShapeV2_stagedOrder)
	case "BottomStackGrowthCurveStartArcShapeV2":
		res = GetNamedStructInstances(stage.BottomStackGrowthCurveStartArcShapeV2s, stage.BottomStackGrowthCurveStartArcShapeV2_stagedOrder)
	case "BottomStackOfGrowthCurveV2":
		res = GetNamedStructInstances(stage.BottomStackOfGrowthCurveV2s, stage.BottomStackOfGrowthCurveV2_stagedOrder)
	case "BottomStartArcShapeV2":
		res = GetNamedStructInstances(stage.BottomStartArcShapeV2s, stage.BottomStartArcShapeV2_stagedOrder)
	case "BottomStartArcShapeV2Grid":
		res = GetNamedStructInstances(stage.BottomStartArcShapeV2Grids, stage.BottomStartArcShapeV2Grid_stagedOrder)
	case "CircleGridShape":
		res = GetNamedStructInstances(stage.CircleGridShapes, stage.CircleGridShape_stagedOrder)
	case "EndArcShapeV2":
		res = GetNamedStructInstances(stage.EndArcShapeV2s, stage.EndArcShapeV2_stagedOrder)
	case "EndArcShapeV2Grid":
		res = GetNamedStructInstances(stage.EndArcShapeV2Grids, stage.EndArcShapeV2Grid_stagedOrder)
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
	case "StackGrowthCurveBezierShape":
		res = GetNamedStructInstances(stage.StackGrowthCurveBezierShapes, stage.StackGrowthCurveBezierShape_stagedOrder)
	case "StackGrowthCurveEndArcShapeV2":
		res = GetNamedStructInstances(stage.StackGrowthCurveEndArcShapeV2s, stage.StackGrowthCurveEndArcShapeV2_stagedOrder)
	case "StackGrowthCurveStartArcShapeV2":
		res = GetNamedStructInstances(stage.StackGrowthCurveStartArcShapeV2s, stage.StackGrowthCurveStartArcShapeV2_stagedOrder)
	case "StackOfGrowthCurve":
		res = GetNamedStructInstances(stage.StackOfGrowthCurves, stage.StackOfGrowthCurve_stagedOrder)
	case "StackOfGrowthCurveV2":
		res = GetNamedStructInstances(stage.StackOfGrowthCurveV2s, stage.StackOfGrowthCurveV2_stagedOrder)
	case "StartArcShapeV2":
		res = GetNamedStructInstances(stage.StartArcShapeV2s, stage.StartArcShapeV2_stagedOrder)
	case "StartArcShapeV2Grid":
		res = GetNamedStructInstances(stage.StartArcShapeV2Grids, stage.StartArcShapeV2Grid_stagedOrder)
	case "TopEndArcShapeV2":
		res = GetNamedStructInstances(stage.TopEndArcShapeV2s, stage.TopEndArcShapeV2_stagedOrder)
	case "TopEndArcShapeV2Grid":
		res = GetNamedStructInstances(stage.TopEndArcShapeV2Grids, stage.TopEndArcShapeV2Grid_stagedOrder)
	case "TopGrowthCurve2D":
		res = GetNamedStructInstances(stage.TopGrowthCurve2Ds, stage.TopGrowthCurve2D_stagedOrder)
	case "TopStackGrowthCurveEndArcShapeV2":
		res = GetNamedStructInstances(stage.TopStackGrowthCurveEndArcShapeV2s, stage.TopStackGrowthCurveEndArcShapeV2_stagedOrder)
	case "TopStackGrowthCurveStartArcShapeV2":
		res = GetNamedStructInstances(stage.TopStackGrowthCurveStartArcShapeV2s, stage.TopStackGrowthCurveStartArcShapeV2_stagedOrder)
	case "TopStackOfGrowthCurveV2":
		res = GetNamedStructInstances(stage.TopStackOfGrowthCurveV2s, stage.TopStackOfGrowthCurveV2_stagedOrder)
	case "TopStartArcShapeV2":
		res = GetNamedStructInstances(stage.TopStartArcShapeV2s, stage.TopStartArcShapeV2_stagedOrder)
	case "TopStartArcShapeV2Grid":
		res = GetNamedStructInstances(stage.TopStartArcShapeV2Grids, stage.TopStartArcShapeV2Grid_stagedOrder)
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
	CommitBottomEndArcShapeV2(bottomendarcshapev2 *BottomEndArcShapeV2)
	CheckoutBottomEndArcShapeV2(bottomendarcshapev2 *BottomEndArcShapeV2)
	CommitBottomEndArcShapeV2Grid(bottomendarcshapev2grid *BottomEndArcShapeV2Grid)
	CheckoutBottomEndArcShapeV2Grid(bottomendarcshapev2grid *BottomEndArcShapeV2Grid)
	CommitBottomStackGrowthCurveEndArcShapeV2(bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2)
	CheckoutBottomStackGrowthCurveEndArcShapeV2(bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2)
	CommitBottomStackGrowthCurveStartArcShapeV2(bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2)
	CheckoutBottomStackGrowthCurveStartArcShapeV2(bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2)
	CommitBottomStackOfGrowthCurveV2(bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2)
	CheckoutBottomStackOfGrowthCurveV2(bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2)
	CommitBottomStartArcShapeV2(bottomstartarcshapev2 *BottomStartArcShapeV2)
	CheckoutBottomStartArcShapeV2(bottomstartarcshapev2 *BottomStartArcShapeV2)
	CommitBottomStartArcShapeV2Grid(bottomstartarcshapev2grid *BottomStartArcShapeV2Grid)
	CheckoutBottomStartArcShapeV2Grid(bottomstartarcshapev2grid *BottomStartArcShapeV2Grid)
	CommitCircleGridShape(circlegridshape *CircleGridShape)
	CheckoutCircleGridShape(circlegridshape *CircleGridShape)
	CommitEndArcShapeV2(endarcshapev2 *EndArcShapeV2)
	CheckoutEndArcShapeV2(endarcshapev2 *EndArcShapeV2)
	CommitEndArcShapeV2Grid(endarcshapev2grid *EndArcShapeV2Grid)
	CheckoutEndArcShapeV2Grid(endarcshapev2grid *EndArcShapeV2Grid)
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
	CommitStackGrowthCurveBezierShape(stackgrowthcurvebeziershape *StackGrowthCurveBezierShape)
	CheckoutStackGrowthCurveBezierShape(stackgrowthcurvebeziershape *StackGrowthCurveBezierShape)
	CommitStackGrowthCurveEndArcShapeV2(stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2)
	CheckoutStackGrowthCurveEndArcShapeV2(stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2)
	CommitStackGrowthCurveStartArcShapeV2(stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2)
	CheckoutStackGrowthCurveStartArcShapeV2(stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2)
	CommitStackOfGrowthCurve(stackofgrowthcurve *StackOfGrowthCurve)
	CheckoutStackOfGrowthCurve(stackofgrowthcurve *StackOfGrowthCurve)
	CommitStackOfGrowthCurveV2(stackofgrowthcurvev2 *StackOfGrowthCurveV2)
	CheckoutStackOfGrowthCurveV2(stackofgrowthcurvev2 *StackOfGrowthCurveV2)
	CommitStartArcShapeV2(startarcshapev2 *StartArcShapeV2)
	CheckoutStartArcShapeV2(startarcshapev2 *StartArcShapeV2)
	CommitStartArcShapeV2Grid(startarcshapev2grid *StartArcShapeV2Grid)
	CheckoutStartArcShapeV2Grid(startarcshapev2grid *StartArcShapeV2Grid)
	CommitTopEndArcShapeV2(topendarcshapev2 *TopEndArcShapeV2)
	CheckoutTopEndArcShapeV2(topendarcshapev2 *TopEndArcShapeV2)
	CommitTopEndArcShapeV2Grid(topendarcshapev2grid *TopEndArcShapeV2Grid)
	CheckoutTopEndArcShapeV2Grid(topendarcshapev2grid *TopEndArcShapeV2Grid)
	CommitTopGrowthCurve2D(topgrowthcurve2d *TopGrowthCurve2D)
	CheckoutTopGrowthCurve2D(topgrowthcurve2d *TopGrowthCurve2D)
	CommitTopStackGrowthCurveEndArcShapeV2(topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2)
	CheckoutTopStackGrowthCurveEndArcShapeV2(topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2)
	CommitTopStackGrowthCurveStartArcShapeV2(topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2)
	CheckoutTopStackGrowthCurveStartArcShapeV2(topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2)
	CommitTopStackOfGrowthCurveV2(topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2)
	CheckoutTopStackOfGrowthCurveV2(topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2)
	CommitTopStartArcShapeV2(topstartarcshapev2 *TopStartArcShapeV2)
	CheckoutTopStartArcShapeV2(topstartarcshapev2 *TopStartArcShapeV2)
	CommitTopStartArcShapeV2Grid(topstartarcshapev2grid *TopStartArcShapeV2Grid)
	CheckoutTopStartArcShapeV2Grid(topstartarcshapev2grid *TopStartArcShapeV2Grid)
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

		BottomEndArcShapeV2s:           make(map[*BottomEndArcShapeV2]struct{}),
		BottomEndArcShapeV2s_mapString: make(map[string]*BottomEndArcShapeV2),

		BottomEndArcShapeV2Grids:           make(map[*BottomEndArcShapeV2Grid]struct{}),
		BottomEndArcShapeV2Grids_mapString: make(map[string]*BottomEndArcShapeV2Grid),

		BottomStackGrowthCurveEndArcShapeV2s:           make(map[*BottomStackGrowthCurveEndArcShapeV2]struct{}),
		BottomStackGrowthCurveEndArcShapeV2s_mapString: make(map[string]*BottomStackGrowthCurveEndArcShapeV2),

		BottomStackGrowthCurveStartArcShapeV2s:           make(map[*BottomStackGrowthCurveStartArcShapeV2]struct{}),
		BottomStackGrowthCurveStartArcShapeV2s_mapString: make(map[string]*BottomStackGrowthCurveStartArcShapeV2),

		BottomStackOfGrowthCurveV2s:           make(map[*BottomStackOfGrowthCurveV2]struct{}),
		BottomStackOfGrowthCurveV2s_mapString: make(map[string]*BottomStackOfGrowthCurveV2),

		BottomStartArcShapeV2s:           make(map[*BottomStartArcShapeV2]struct{}),
		BottomStartArcShapeV2s_mapString: make(map[string]*BottomStartArcShapeV2),

		BottomStartArcShapeV2Grids:           make(map[*BottomStartArcShapeV2Grid]struct{}),
		BottomStartArcShapeV2Grids_mapString: make(map[string]*BottomStartArcShapeV2Grid),

		CircleGridShapes:           make(map[*CircleGridShape]struct{}),
		CircleGridShapes_mapString: make(map[string]*CircleGridShape),

		EndArcShapeV2s:           make(map[*EndArcShapeV2]struct{}),
		EndArcShapeV2s_mapString: make(map[string]*EndArcShapeV2),

		EndArcShapeV2Grids:           make(map[*EndArcShapeV2Grid]struct{}),
		EndArcShapeV2Grids_mapString: make(map[string]*EndArcShapeV2Grid),

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

		StackGrowthCurveBezierShapes:           make(map[*StackGrowthCurveBezierShape]struct{}),
		StackGrowthCurveBezierShapes_mapString: make(map[string]*StackGrowthCurveBezierShape),

		StackGrowthCurveEndArcShapeV2s:           make(map[*StackGrowthCurveEndArcShapeV2]struct{}),
		StackGrowthCurveEndArcShapeV2s_mapString: make(map[string]*StackGrowthCurveEndArcShapeV2),

		StackGrowthCurveStartArcShapeV2s:           make(map[*StackGrowthCurveStartArcShapeV2]struct{}),
		StackGrowthCurveStartArcShapeV2s_mapString: make(map[string]*StackGrowthCurveStartArcShapeV2),

		StackOfGrowthCurves:           make(map[*StackOfGrowthCurve]struct{}),
		StackOfGrowthCurves_mapString: make(map[string]*StackOfGrowthCurve),

		StackOfGrowthCurveV2s:           make(map[*StackOfGrowthCurveV2]struct{}),
		StackOfGrowthCurveV2s_mapString: make(map[string]*StackOfGrowthCurveV2),

		StartArcShapeV2s:           make(map[*StartArcShapeV2]struct{}),
		StartArcShapeV2s_mapString: make(map[string]*StartArcShapeV2),

		StartArcShapeV2Grids:           make(map[*StartArcShapeV2Grid]struct{}),
		StartArcShapeV2Grids_mapString: make(map[string]*StartArcShapeV2Grid),

		TopEndArcShapeV2s:           make(map[*TopEndArcShapeV2]struct{}),
		TopEndArcShapeV2s_mapString: make(map[string]*TopEndArcShapeV2),

		TopEndArcShapeV2Grids:           make(map[*TopEndArcShapeV2Grid]struct{}),
		TopEndArcShapeV2Grids_mapString: make(map[string]*TopEndArcShapeV2Grid),

		TopGrowthCurve2Ds:           make(map[*TopGrowthCurve2D]struct{}),
		TopGrowthCurve2Ds_mapString: make(map[string]*TopGrowthCurve2D),

		TopStackGrowthCurveEndArcShapeV2s:           make(map[*TopStackGrowthCurveEndArcShapeV2]struct{}),
		TopStackGrowthCurveEndArcShapeV2s_mapString: make(map[string]*TopStackGrowthCurveEndArcShapeV2),

		TopStackGrowthCurveStartArcShapeV2s:           make(map[*TopStackGrowthCurveStartArcShapeV2]struct{}),
		TopStackGrowthCurveStartArcShapeV2s_mapString: make(map[string]*TopStackGrowthCurveStartArcShapeV2),

		TopStackOfGrowthCurveV2s:           make(map[*TopStackOfGrowthCurveV2]struct{}),
		TopStackOfGrowthCurveV2s_mapString: make(map[string]*TopStackOfGrowthCurveV2),

		TopStartArcShapeV2s:           make(map[*TopStartArcShapeV2]struct{}),
		TopStartArcShapeV2s_mapString: make(map[string]*TopStartArcShapeV2),

		TopStartArcShapeV2Grids:           make(map[*TopStartArcShapeV2Grid]struct{}),
		TopStartArcShapeV2Grids_mapString: make(map[string]*TopStartArcShapeV2Grid),

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

		BottomEndArcShapeV2_stagedOrder: make(map[*BottomEndArcShapeV2]uint),
		BottomEndArcShapeV2_orderStaged: make(map[uint]*BottomEndArcShapeV2),
		BottomEndArcShapeV2s_reference:  make(map[*BottomEndArcShapeV2]*BottomEndArcShapeV2),

		BottomEndArcShapeV2Grid_stagedOrder: make(map[*BottomEndArcShapeV2Grid]uint),
		BottomEndArcShapeV2Grid_orderStaged: make(map[uint]*BottomEndArcShapeV2Grid),
		BottomEndArcShapeV2Grids_reference:  make(map[*BottomEndArcShapeV2Grid]*BottomEndArcShapeV2Grid),

		BottomStackGrowthCurveEndArcShapeV2_stagedOrder: make(map[*BottomStackGrowthCurveEndArcShapeV2]uint),
		BottomStackGrowthCurveEndArcShapeV2_orderStaged: make(map[uint]*BottomStackGrowthCurveEndArcShapeV2),
		BottomStackGrowthCurveEndArcShapeV2s_reference:  make(map[*BottomStackGrowthCurveEndArcShapeV2]*BottomStackGrowthCurveEndArcShapeV2),

		BottomStackGrowthCurveStartArcShapeV2_stagedOrder: make(map[*BottomStackGrowthCurveStartArcShapeV2]uint),
		BottomStackGrowthCurveStartArcShapeV2_orderStaged: make(map[uint]*BottomStackGrowthCurveStartArcShapeV2),
		BottomStackGrowthCurveStartArcShapeV2s_reference:  make(map[*BottomStackGrowthCurveStartArcShapeV2]*BottomStackGrowthCurveStartArcShapeV2),

		BottomStackOfGrowthCurveV2_stagedOrder: make(map[*BottomStackOfGrowthCurveV2]uint),
		BottomStackOfGrowthCurveV2_orderStaged: make(map[uint]*BottomStackOfGrowthCurveV2),
		BottomStackOfGrowthCurveV2s_reference:  make(map[*BottomStackOfGrowthCurveV2]*BottomStackOfGrowthCurveV2),

		BottomStartArcShapeV2_stagedOrder: make(map[*BottomStartArcShapeV2]uint),
		BottomStartArcShapeV2_orderStaged: make(map[uint]*BottomStartArcShapeV2),
		BottomStartArcShapeV2s_reference:  make(map[*BottomStartArcShapeV2]*BottomStartArcShapeV2),

		BottomStartArcShapeV2Grid_stagedOrder: make(map[*BottomStartArcShapeV2Grid]uint),
		BottomStartArcShapeV2Grid_orderStaged: make(map[uint]*BottomStartArcShapeV2Grid),
		BottomStartArcShapeV2Grids_reference:  make(map[*BottomStartArcShapeV2Grid]*BottomStartArcShapeV2Grid),

		CircleGridShape_stagedOrder: make(map[*CircleGridShape]uint),
		CircleGridShape_orderStaged: make(map[uint]*CircleGridShape),
		CircleGridShapes_reference:  make(map[*CircleGridShape]*CircleGridShape),

		EndArcShapeV2_stagedOrder: make(map[*EndArcShapeV2]uint),
		EndArcShapeV2_orderStaged: make(map[uint]*EndArcShapeV2),
		EndArcShapeV2s_reference:  make(map[*EndArcShapeV2]*EndArcShapeV2),

		EndArcShapeV2Grid_stagedOrder: make(map[*EndArcShapeV2Grid]uint),
		EndArcShapeV2Grid_orderStaged: make(map[uint]*EndArcShapeV2Grid),
		EndArcShapeV2Grids_reference:  make(map[*EndArcShapeV2Grid]*EndArcShapeV2Grid),

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

		StackGrowthCurveBezierShape_stagedOrder: make(map[*StackGrowthCurveBezierShape]uint),
		StackGrowthCurveBezierShape_orderStaged: make(map[uint]*StackGrowthCurveBezierShape),
		StackGrowthCurveBezierShapes_reference:  make(map[*StackGrowthCurveBezierShape]*StackGrowthCurveBezierShape),

		StackGrowthCurveEndArcShapeV2_stagedOrder: make(map[*StackGrowthCurveEndArcShapeV2]uint),
		StackGrowthCurveEndArcShapeV2_orderStaged: make(map[uint]*StackGrowthCurveEndArcShapeV2),
		StackGrowthCurveEndArcShapeV2s_reference:  make(map[*StackGrowthCurveEndArcShapeV2]*StackGrowthCurveEndArcShapeV2),

		StackGrowthCurveStartArcShapeV2_stagedOrder: make(map[*StackGrowthCurveStartArcShapeV2]uint),
		StackGrowthCurveStartArcShapeV2_orderStaged: make(map[uint]*StackGrowthCurveStartArcShapeV2),
		StackGrowthCurveStartArcShapeV2s_reference:  make(map[*StackGrowthCurveStartArcShapeV2]*StackGrowthCurveStartArcShapeV2),

		StackOfGrowthCurve_stagedOrder: make(map[*StackOfGrowthCurve]uint),
		StackOfGrowthCurve_orderStaged: make(map[uint]*StackOfGrowthCurve),
		StackOfGrowthCurves_reference:  make(map[*StackOfGrowthCurve]*StackOfGrowthCurve),

		StackOfGrowthCurveV2_stagedOrder: make(map[*StackOfGrowthCurveV2]uint),
		StackOfGrowthCurveV2_orderStaged: make(map[uint]*StackOfGrowthCurveV2),
		StackOfGrowthCurveV2s_reference:  make(map[*StackOfGrowthCurveV2]*StackOfGrowthCurveV2),

		StartArcShapeV2_stagedOrder: make(map[*StartArcShapeV2]uint),
		StartArcShapeV2_orderStaged: make(map[uint]*StartArcShapeV2),
		StartArcShapeV2s_reference:  make(map[*StartArcShapeV2]*StartArcShapeV2),

		StartArcShapeV2Grid_stagedOrder: make(map[*StartArcShapeV2Grid]uint),
		StartArcShapeV2Grid_orderStaged: make(map[uint]*StartArcShapeV2Grid),
		StartArcShapeV2Grids_reference:  make(map[*StartArcShapeV2Grid]*StartArcShapeV2Grid),

		TopEndArcShapeV2_stagedOrder: make(map[*TopEndArcShapeV2]uint),
		TopEndArcShapeV2_orderStaged: make(map[uint]*TopEndArcShapeV2),
		TopEndArcShapeV2s_reference:  make(map[*TopEndArcShapeV2]*TopEndArcShapeV2),

		TopEndArcShapeV2Grid_stagedOrder: make(map[*TopEndArcShapeV2Grid]uint),
		TopEndArcShapeV2Grid_orderStaged: make(map[uint]*TopEndArcShapeV2Grid),
		TopEndArcShapeV2Grids_reference:  make(map[*TopEndArcShapeV2Grid]*TopEndArcShapeV2Grid),

		TopGrowthCurve2D_stagedOrder: make(map[*TopGrowthCurve2D]uint),
		TopGrowthCurve2D_orderStaged: make(map[uint]*TopGrowthCurve2D),
		TopGrowthCurve2Ds_reference:  make(map[*TopGrowthCurve2D]*TopGrowthCurve2D),

		TopStackGrowthCurveEndArcShapeV2_stagedOrder: make(map[*TopStackGrowthCurveEndArcShapeV2]uint),
		TopStackGrowthCurveEndArcShapeV2_orderStaged: make(map[uint]*TopStackGrowthCurveEndArcShapeV2),
		TopStackGrowthCurveEndArcShapeV2s_reference:  make(map[*TopStackGrowthCurveEndArcShapeV2]*TopStackGrowthCurveEndArcShapeV2),

		TopStackGrowthCurveStartArcShapeV2_stagedOrder: make(map[*TopStackGrowthCurveStartArcShapeV2]uint),
		TopStackGrowthCurveStartArcShapeV2_orderStaged: make(map[uint]*TopStackGrowthCurveStartArcShapeV2),
		TopStackGrowthCurveStartArcShapeV2s_reference:  make(map[*TopStackGrowthCurveStartArcShapeV2]*TopStackGrowthCurveStartArcShapeV2),

		TopStackOfGrowthCurveV2_stagedOrder: make(map[*TopStackOfGrowthCurveV2]uint),
		TopStackOfGrowthCurveV2_orderStaged: make(map[uint]*TopStackOfGrowthCurveV2),
		TopStackOfGrowthCurveV2s_reference:  make(map[*TopStackOfGrowthCurveV2]*TopStackOfGrowthCurveV2),

		TopStartArcShapeV2_stagedOrder: make(map[*TopStartArcShapeV2]uint),
		TopStartArcShapeV2_orderStaged: make(map[uint]*TopStartArcShapeV2),
		TopStartArcShapeV2s_reference:  make(map[*TopStartArcShapeV2]*TopStartArcShapeV2),

		TopStartArcShapeV2Grid_stagedOrder: make(map[*TopStartArcShapeV2Grid]uint),
		TopStartArcShapeV2Grid_orderStaged: make(map[uint]*TopStartArcShapeV2Grid),
		TopStartArcShapeV2Grids_reference:  make(map[*TopStartArcShapeV2Grid]*TopStartArcShapeV2Grid),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"ArcNormalVectorShape": &ArcNormalVectorShapeUnmarshaller{},

			"ArcNormalVectorShapeGrid": &ArcNormalVectorShapeGridUnmarshaller{},

			"AxesShape": &AxesShapeUnmarshaller{},

			"BaseVectorShape": &BaseVectorShapeUnmarshaller{},

			"BaseVectorShapeGrid": &BaseVectorShapeGridUnmarshaller{},

			"BottomEndArcShapeV2": &BottomEndArcShapeV2Unmarshaller{},

			"BottomEndArcShapeV2Grid": &BottomEndArcShapeV2GridUnmarshaller{},

			"BottomStackGrowthCurveEndArcShapeV2": &BottomStackGrowthCurveEndArcShapeV2Unmarshaller{},

			"BottomStackGrowthCurveStartArcShapeV2": &BottomStackGrowthCurveStartArcShapeV2Unmarshaller{},

			"BottomStackOfGrowthCurveV2": &BottomStackOfGrowthCurveV2Unmarshaller{},

			"BottomStartArcShapeV2": &BottomStartArcShapeV2Unmarshaller{},

			"BottomStartArcShapeV2Grid": &BottomStartArcShapeV2GridUnmarshaller{},

			"CircleGridShape": &CircleGridShapeUnmarshaller{},

			"EndArcShapeV2": &EndArcShapeV2Unmarshaller{},

			"EndArcShapeV2Grid": &EndArcShapeV2GridUnmarshaller{},

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

			"StackGrowthCurveBezierShape": &StackGrowthCurveBezierShapeUnmarshaller{},

			"StackGrowthCurveEndArcShapeV2": &StackGrowthCurveEndArcShapeV2Unmarshaller{},

			"StackGrowthCurveStartArcShapeV2": &StackGrowthCurveStartArcShapeV2Unmarshaller{},

			"StackOfGrowthCurve": &StackOfGrowthCurveUnmarshaller{},

			"StackOfGrowthCurveV2": &StackOfGrowthCurveV2Unmarshaller{},

			"StartArcShapeV2": &StartArcShapeV2Unmarshaller{},

			"StartArcShapeV2Grid": &StartArcShapeV2GridUnmarshaller{},

			"TopEndArcShapeV2": &TopEndArcShapeV2Unmarshaller{},

			"TopEndArcShapeV2Grid": &TopEndArcShapeV2GridUnmarshaller{},

			"TopGrowthCurve2D": &TopGrowthCurve2DUnmarshaller{},

			"TopStackGrowthCurveEndArcShapeV2": &TopStackGrowthCurveEndArcShapeV2Unmarshaller{},

			"TopStackGrowthCurveStartArcShapeV2": &TopStackGrowthCurveStartArcShapeV2Unmarshaller{},

			"TopStackOfGrowthCurveV2": &TopStackOfGrowthCurveV2Unmarshaller{},

			"TopStartArcShapeV2": &TopStartArcShapeV2Unmarshaller{},

			"TopStartArcShapeV2Grid": &TopStartArcShapeV2GridUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "ArcNormalVectorShape"},
			{name: "ArcNormalVectorShapeGrid"},
			{name: "AxesShape"},
			{name: "BaseVectorShape"},
			{name: "BaseVectorShapeGrid"},
			{name: "BottomEndArcShapeV2"},
			{name: "BottomEndArcShapeV2Grid"},
			{name: "BottomStackGrowthCurveEndArcShapeV2"},
			{name: "BottomStackGrowthCurveStartArcShapeV2"},
			{name: "BottomStackOfGrowthCurveV2"},
			{name: "BottomStartArcShapeV2"},
			{name: "BottomStartArcShapeV2Grid"},
			{name: "CircleGridShape"},
			{name: "EndArcShapeV2"},
			{name: "EndArcShapeV2Grid"},
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
			{name: "StackGrowthCurveBezierShape"},
			{name: "StackGrowthCurveEndArcShapeV2"},
			{name: "StackGrowthCurveStartArcShapeV2"},
			{name: "StackOfGrowthCurve"},
			{name: "StackOfGrowthCurveV2"},
			{name: "StartArcShapeV2"},
			{name: "StartArcShapeV2Grid"},
			{name: "TopEndArcShapeV2"},
			{name: "TopEndArcShapeV2Grid"},
			{name: "TopGrowthCurve2D"},
			{name: "TopStackGrowthCurveEndArcShapeV2"},
			{name: "TopStackGrowthCurveStartArcShapeV2"},
			{name: "TopStackOfGrowthCurveV2"},
			{name: "TopStartArcShapeV2"},
			{name: "TopStartArcShapeV2Grid"},
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
	case *BottomEndArcShapeV2:
		return stage.BottomEndArcShapeV2_stagedOrder[instance]
	case *BottomEndArcShapeV2Grid:
		return stage.BottomEndArcShapeV2Grid_stagedOrder[instance]
	case *BottomStackGrowthCurveEndArcShapeV2:
		return stage.BottomStackGrowthCurveEndArcShapeV2_stagedOrder[instance]
	case *BottomStackGrowthCurveStartArcShapeV2:
		return stage.BottomStackGrowthCurveStartArcShapeV2_stagedOrder[instance]
	case *BottomStackOfGrowthCurveV2:
		return stage.BottomStackOfGrowthCurveV2_stagedOrder[instance]
	case *BottomStartArcShapeV2:
		return stage.BottomStartArcShapeV2_stagedOrder[instance]
	case *BottomStartArcShapeV2Grid:
		return stage.BottomStartArcShapeV2Grid_stagedOrder[instance]
	case *CircleGridShape:
		return stage.CircleGridShape_stagedOrder[instance]
	case *EndArcShapeV2:
		return stage.EndArcShapeV2_stagedOrder[instance]
	case *EndArcShapeV2Grid:
		return stage.EndArcShapeV2Grid_stagedOrder[instance]
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
	case *StackGrowthCurveBezierShape:
		return stage.StackGrowthCurveBezierShape_stagedOrder[instance]
	case *StackGrowthCurveEndArcShapeV2:
		return stage.StackGrowthCurveEndArcShapeV2_stagedOrder[instance]
	case *StackGrowthCurveStartArcShapeV2:
		return stage.StackGrowthCurveStartArcShapeV2_stagedOrder[instance]
	case *StackOfGrowthCurve:
		return stage.StackOfGrowthCurve_stagedOrder[instance]
	case *StackOfGrowthCurveV2:
		return stage.StackOfGrowthCurveV2_stagedOrder[instance]
	case *StartArcShapeV2:
		return stage.StartArcShapeV2_stagedOrder[instance]
	case *StartArcShapeV2Grid:
		return stage.StartArcShapeV2Grid_stagedOrder[instance]
	case *TopEndArcShapeV2:
		return stage.TopEndArcShapeV2_stagedOrder[instance]
	case *TopEndArcShapeV2Grid:
		return stage.TopEndArcShapeV2Grid_stagedOrder[instance]
	case *TopGrowthCurve2D:
		return stage.TopGrowthCurve2D_stagedOrder[instance]
	case *TopStackGrowthCurveEndArcShapeV2:
		return stage.TopStackGrowthCurveEndArcShapeV2_stagedOrder[instance]
	case *TopStackGrowthCurveStartArcShapeV2:
		return stage.TopStackGrowthCurveStartArcShapeV2_stagedOrder[instance]
	case *TopStackOfGrowthCurveV2:
		return stage.TopStackOfGrowthCurveV2_stagedOrder[instance]
	case *TopStartArcShapeV2:
		return stage.TopStartArcShapeV2_stagedOrder[instance]
	case *TopStartArcShapeV2Grid:
		return stage.TopStartArcShapeV2Grid_stagedOrder[instance]
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
	case *BottomEndArcShapeV2:
		return any(stage.BottomEndArcShapeV2_orderStaged[order]).(Type)
	case *BottomEndArcShapeV2Grid:
		return any(stage.BottomEndArcShapeV2Grid_orderStaged[order]).(Type)
	case *BottomStackGrowthCurveEndArcShapeV2:
		return any(stage.BottomStackGrowthCurveEndArcShapeV2_orderStaged[order]).(Type)
	case *BottomStackGrowthCurveStartArcShapeV2:
		return any(stage.BottomStackGrowthCurveStartArcShapeV2_orderStaged[order]).(Type)
	case *BottomStackOfGrowthCurveV2:
		return any(stage.BottomStackOfGrowthCurveV2_orderStaged[order]).(Type)
	case *BottomStartArcShapeV2:
		return any(stage.BottomStartArcShapeV2_orderStaged[order]).(Type)
	case *BottomStartArcShapeV2Grid:
		return any(stage.BottomStartArcShapeV2Grid_orderStaged[order]).(Type)
	case *CircleGridShape:
		return any(stage.CircleGridShape_orderStaged[order]).(Type)
	case *EndArcShapeV2:
		return any(stage.EndArcShapeV2_orderStaged[order]).(Type)
	case *EndArcShapeV2Grid:
		return any(stage.EndArcShapeV2Grid_orderStaged[order]).(Type)
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
	case *StackGrowthCurveBezierShape:
		return any(stage.StackGrowthCurveBezierShape_orderStaged[order]).(Type)
	case *StackGrowthCurveEndArcShapeV2:
		return any(stage.StackGrowthCurveEndArcShapeV2_orderStaged[order]).(Type)
	case *StackGrowthCurveStartArcShapeV2:
		return any(stage.StackGrowthCurveStartArcShapeV2_orderStaged[order]).(Type)
	case *StackOfGrowthCurve:
		return any(stage.StackOfGrowthCurve_orderStaged[order]).(Type)
	case *StackOfGrowthCurveV2:
		return any(stage.StackOfGrowthCurveV2_orderStaged[order]).(Type)
	case *StartArcShapeV2:
		return any(stage.StartArcShapeV2_orderStaged[order]).(Type)
	case *StartArcShapeV2Grid:
		return any(stage.StartArcShapeV2Grid_orderStaged[order]).(Type)
	case *TopEndArcShapeV2:
		return any(stage.TopEndArcShapeV2_orderStaged[order]).(Type)
	case *TopEndArcShapeV2Grid:
		return any(stage.TopEndArcShapeV2Grid_orderStaged[order]).(Type)
	case *TopGrowthCurve2D:
		return any(stage.TopGrowthCurve2D_orderStaged[order]).(Type)
	case *TopStackGrowthCurveEndArcShapeV2:
		return any(stage.TopStackGrowthCurveEndArcShapeV2_orderStaged[order]).(Type)
	case *TopStackGrowthCurveStartArcShapeV2:
		return any(stage.TopStackGrowthCurveStartArcShapeV2_orderStaged[order]).(Type)
	case *TopStackOfGrowthCurveV2:
		return any(stage.TopStackOfGrowthCurveV2_orderStaged[order]).(Type)
	case *TopStartArcShapeV2:
		return any(stage.TopStartArcShapeV2_orderStaged[order]).(Type)
	case *TopStartArcShapeV2Grid:
		return any(stage.TopStartArcShapeV2Grid_orderStaged[order]).(Type)
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
	case *BottomEndArcShapeV2:
		return stage.BottomEndArcShapeV2_stagedOrder[instance]
	case *BottomEndArcShapeV2Grid:
		return stage.BottomEndArcShapeV2Grid_stagedOrder[instance]
	case *BottomStackGrowthCurveEndArcShapeV2:
		return stage.BottomStackGrowthCurveEndArcShapeV2_stagedOrder[instance]
	case *BottomStackGrowthCurveStartArcShapeV2:
		return stage.BottomStackGrowthCurveStartArcShapeV2_stagedOrder[instance]
	case *BottomStackOfGrowthCurveV2:
		return stage.BottomStackOfGrowthCurveV2_stagedOrder[instance]
	case *BottomStartArcShapeV2:
		return stage.BottomStartArcShapeV2_stagedOrder[instance]
	case *BottomStartArcShapeV2Grid:
		return stage.BottomStartArcShapeV2Grid_stagedOrder[instance]
	case *CircleGridShape:
		return stage.CircleGridShape_stagedOrder[instance]
	case *EndArcShapeV2:
		return stage.EndArcShapeV2_stagedOrder[instance]
	case *EndArcShapeV2Grid:
		return stage.EndArcShapeV2Grid_stagedOrder[instance]
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
	case *StackGrowthCurveBezierShape:
		return stage.StackGrowthCurveBezierShape_stagedOrder[instance]
	case *StackGrowthCurveEndArcShapeV2:
		return stage.StackGrowthCurveEndArcShapeV2_stagedOrder[instance]
	case *StackGrowthCurveStartArcShapeV2:
		return stage.StackGrowthCurveStartArcShapeV2_stagedOrder[instance]
	case *StackOfGrowthCurve:
		return stage.StackOfGrowthCurve_stagedOrder[instance]
	case *StackOfGrowthCurveV2:
		return stage.StackOfGrowthCurveV2_stagedOrder[instance]
	case *StartArcShapeV2:
		return stage.StartArcShapeV2_stagedOrder[instance]
	case *StartArcShapeV2Grid:
		return stage.StartArcShapeV2Grid_stagedOrder[instance]
	case *TopEndArcShapeV2:
		return stage.TopEndArcShapeV2_stagedOrder[instance]
	case *TopEndArcShapeV2Grid:
		return stage.TopEndArcShapeV2Grid_stagedOrder[instance]
	case *TopGrowthCurve2D:
		return stage.TopGrowthCurve2D_stagedOrder[instance]
	case *TopStackGrowthCurveEndArcShapeV2:
		return stage.TopStackGrowthCurveEndArcShapeV2_stagedOrder[instance]
	case *TopStackGrowthCurveStartArcShapeV2:
		return stage.TopStackGrowthCurveStartArcShapeV2_stagedOrder[instance]
	case *TopStackOfGrowthCurveV2:
		return stage.TopStackOfGrowthCurveV2_stagedOrder[instance]
	case *TopStartArcShapeV2:
		return stage.TopStartArcShapeV2_stagedOrder[instance]
	case *TopStartArcShapeV2Grid:
		return stage.TopStartArcShapeV2Grid_stagedOrder[instance]
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
	stage.Map_GongStructName_InstancesNb["BottomEndArcShapeV2"] = len(stage.BottomEndArcShapeV2s)
	stage.Map_GongStructName_InstancesNb["BottomEndArcShapeV2Grid"] = len(stage.BottomEndArcShapeV2Grids)
	stage.Map_GongStructName_InstancesNb["BottomStackGrowthCurveEndArcShapeV2"] = len(stage.BottomStackGrowthCurveEndArcShapeV2s)
	stage.Map_GongStructName_InstancesNb["BottomStackGrowthCurveStartArcShapeV2"] = len(stage.BottomStackGrowthCurveStartArcShapeV2s)
	stage.Map_GongStructName_InstancesNb["BottomStackOfGrowthCurveV2"] = len(stage.BottomStackOfGrowthCurveV2s)
	stage.Map_GongStructName_InstancesNb["BottomStartArcShapeV2"] = len(stage.BottomStartArcShapeV2s)
	stage.Map_GongStructName_InstancesNb["BottomStartArcShapeV2Grid"] = len(stage.BottomStartArcShapeV2Grids)
	stage.Map_GongStructName_InstancesNb["CircleGridShape"] = len(stage.CircleGridShapes)
	stage.Map_GongStructName_InstancesNb["EndArcShapeV2"] = len(stage.EndArcShapeV2s)
	stage.Map_GongStructName_InstancesNb["EndArcShapeV2Grid"] = len(stage.EndArcShapeV2Grids)
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
	stage.Map_GongStructName_InstancesNb["StackGrowthCurveBezierShape"] = len(stage.StackGrowthCurveBezierShapes)
	stage.Map_GongStructName_InstancesNb["StackGrowthCurveEndArcShapeV2"] = len(stage.StackGrowthCurveEndArcShapeV2s)
	stage.Map_GongStructName_InstancesNb["StackGrowthCurveStartArcShapeV2"] = len(stage.StackGrowthCurveStartArcShapeV2s)
	stage.Map_GongStructName_InstancesNb["StackOfGrowthCurve"] = len(stage.StackOfGrowthCurves)
	stage.Map_GongStructName_InstancesNb["StackOfGrowthCurveV2"] = len(stage.StackOfGrowthCurveV2s)
	stage.Map_GongStructName_InstancesNb["StartArcShapeV2"] = len(stage.StartArcShapeV2s)
	stage.Map_GongStructName_InstancesNb["StartArcShapeV2Grid"] = len(stage.StartArcShapeV2Grids)
	stage.Map_GongStructName_InstancesNb["TopEndArcShapeV2"] = len(stage.TopEndArcShapeV2s)
	stage.Map_GongStructName_InstancesNb["TopEndArcShapeV2Grid"] = len(stage.TopEndArcShapeV2Grids)
	stage.Map_GongStructName_InstancesNb["TopGrowthCurve2D"] = len(stage.TopGrowthCurve2Ds)
	stage.Map_GongStructName_InstancesNb["TopStackGrowthCurveEndArcShapeV2"] = len(stage.TopStackGrowthCurveEndArcShapeV2s)
	stage.Map_GongStructName_InstancesNb["TopStackGrowthCurveStartArcShapeV2"] = len(stage.TopStackGrowthCurveStartArcShapeV2s)
	stage.Map_GongStructName_InstancesNb["TopStackOfGrowthCurveV2"] = len(stage.TopStackOfGrowthCurveV2s)
	stage.Map_GongStructName_InstancesNb["TopStartArcShapeV2"] = len(stage.TopStartArcShapeV2s)
	stage.Map_GongStructName_InstancesNb["TopStartArcShapeV2Grid"] = len(stage.TopStartArcShapeV2Grids)
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

// Stage puts bottomendarcshapev2 to the model stage
func (bottomendarcshapev2 *BottomEndArcShapeV2) Stage(stage *Stage) *BottomEndArcShapeV2 {
	if _, ok := stage.BottomEndArcShapeV2s[bottomendarcshapev2]; !ok {
		stage.BottomEndArcShapeV2s[bottomendarcshapev2] = struct{}{}
		stage.BottomEndArcShapeV2_stagedOrder[bottomendarcshapev2] = stage.BottomEndArcShapeV2Order
		stage.BottomEndArcShapeV2_orderStaged[stage.BottomEndArcShapeV2Order] = bottomendarcshapev2
		stage.BottomEndArcShapeV2Order++
	}
	stage.BottomEndArcShapeV2s_mapString[bottomendarcshapev2.Name] = bottomendarcshapev2

	return bottomendarcshapev2
}

// StagePreserveOrder puts bottomendarcshapev2 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BottomEndArcShapeV2Order
// - update stage.BottomEndArcShapeV2Order accordingly
func (bottomendarcshapev2 *BottomEndArcShapeV2) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.BottomEndArcShapeV2s[bottomendarcshapev2]; !ok {
		stage.BottomEndArcShapeV2s[bottomendarcshapev2] = struct{}{}

		if order > stage.BottomEndArcShapeV2Order {
			stage.BottomEndArcShapeV2Order = order
		}
		stage.BottomEndArcShapeV2_stagedOrder[bottomendarcshapev2] = order
		stage.BottomEndArcShapeV2_orderStaged[order] = bottomendarcshapev2
		stage.BottomEndArcShapeV2Order++
	}
	stage.BottomEndArcShapeV2s_mapString[bottomendarcshapev2.Name] = bottomendarcshapev2
}

// Unstage removes bottomendarcshapev2 off the model stage
func (bottomendarcshapev2 *BottomEndArcShapeV2) Unstage(stage *Stage) *BottomEndArcShapeV2 {
	delete(stage.BottomEndArcShapeV2s, bottomendarcshapev2)
	// issue1150
	// delete(stage.BottomEndArcShapeV2_stagedOrder, bottomendarcshapev2)
	delete(stage.BottomEndArcShapeV2s_mapString, bottomendarcshapev2.Name)

	return bottomendarcshapev2
}

// UnstageVoid removes bottomendarcshapev2 off the model stage
func (bottomendarcshapev2 *BottomEndArcShapeV2) UnstageVoid(stage *Stage) {
	delete(stage.BottomEndArcShapeV2s, bottomendarcshapev2)
	// issue1150
	// delete(stage.BottomEndArcShapeV2_stagedOrder, bottomendarcshapev2)
	delete(stage.BottomEndArcShapeV2s_mapString, bottomendarcshapev2.Name)
}

// commit bottomendarcshapev2 to the back repo (if it is already staged)
func (bottomendarcshapev2 *BottomEndArcShapeV2) Commit(stage *Stage) *BottomEndArcShapeV2 {
	if _, ok := stage.BottomEndArcShapeV2s[bottomendarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBottomEndArcShapeV2(bottomendarcshapev2)
		}
	}
	return bottomendarcshapev2
}

func (bottomendarcshapev2 *BottomEndArcShapeV2) CommitVoid(stage *Stage) {
	bottomendarcshapev2.Commit(stage)
}

func (bottomendarcshapev2 *BottomEndArcShapeV2) StageVoid(stage *Stage) {
	bottomendarcshapev2.Stage(stage)
}

// Checkout bottomendarcshapev2 to the back repo (if it is already staged)
func (bottomendarcshapev2 *BottomEndArcShapeV2) Checkout(stage *Stage) *BottomEndArcShapeV2 {
	if _, ok := stage.BottomEndArcShapeV2s[bottomendarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBottomEndArcShapeV2(bottomendarcshapev2)
		}
	}
	return bottomendarcshapev2
}

// for satisfaction of GongStruct interface
func (bottomendarcshapev2 *BottomEndArcShapeV2) GetName() (res string) {
	return bottomendarcshapev2.Name
}

// for satisfaction of GongStruct interface
func (bottomendarcshapev2 *BottomEndArcShapeV2) SetName(name string) {
	bottomendarcshapev2.Name = name
}

// Stage puts bottomendarcshapev2grid to the model stage
func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) Stage(stage *Stage) *BottomEndArcShapeV2Grid {
	if _, ok := stage.BottomEndArcShapeV2Grids[bottomendarcshapev2grid]; !ok {
		stage.BottomEndArcShapeV2Grids[bottomendarcshapev2grid] = struct{}{}
		stage.BottomEndArcShapeV2Grid_stagedOrder[bottomendarcshapev2grid] = stage.BottomEndArcShapeV2GridOrder
		stage.BottomEndArcShapeV2Grid_orderStaged[stage.BottomEndArcShapeV2GridOrder] = bottomendarcshapev2grid
		stage.BottomEndArcShapeV2GridOrder++
	}
	stage.BottomEndArcShapeV2Grids_mapString[bottomendarcshapev2grid.Name] = bottomendarcshapev2grid

	return bottomendarcshapev2grid
}

// StagePreserveOrder puts bottomendarcshapev2grid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BottomEndArcShapeV2GridOrder
// - update stage.BottomEndArcShapeV2GridOrder accordingly
func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.BottomEndArcShapeV2Grids[bottomendarcshapev2grid]; !ok {
		stage.BottomEndArcShapeV2Grids[bottomendarcshapev2grid] = struct{}{}

		if order > stage.BottomEndArcShapeV2GridOrder {
			stage.BottomEndArcShapeV2GridOrder = order
		}
		stage.BottomEndArcShapeV2Grid_stagedOrder[bottomendarcshapev2grid] = order
		stage.BottomEndArcShapeV2Grid_orderStaged[order] = bottomendarcshapev2grid
		stage.BottomEndArcShapeV2GridOrder++
	}
	stage.BottomEndArcShapeV2Grids_mapString[bottomendarcshapev2grid.Name] = bottomendarcshapev2grid
}

// Unstage removes bottomendarcshapev2grid off the model stage
func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) Unstage(stage *Stage) *BottomEndArcShapeV2Grid {
	delete(stage.BottomEndArcShapeV2Grids, bottomendarcshapev2grid)
	// issue1150
	// delete(stage.BottomEndArcShapeV2Grid_stagedOrder, bottomendarcshapev2grid)
	delete(stage.BottomEndArcShapeV2Grids_mapString, bottomendarcshapev2grid.Name)

	return bottomendarcshapev2grid
}

// UnstageVoid removes bottomendarcshapev2grid off the model stage
func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) UnstageVoid(stage *Stage) {
	delete(stage.BottomEndArcShapeV2Grids, bottomendarcshapev2grid)
	// issue1150
	// delete(stage.BottomEndArcShapeV2Grid_stagedOrder, bottomendarcshapev2grid)
	delete(stage.BottomEndArcShapeV2Grids_mapString, bottomendarcshapev2grid.Name)
}

// commit bottomendarcshapev2grid to the back repo (if it is already staged)
func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) Commit(stage *Stage) *BottomEndArcShapeV2Grid {
	if _, ok := stage.BottomEndArcShapeV2Grids[bottomendarcshapev2grid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBottomEndArcShapeV2Grid(bottomendarcshapev2grid)
		}
	}
	return bottomendarcshapev2grid
}

func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) CommitVoid(stage *Stage) {
	bottomendarcshapev2grid.Commit(stage)
}

func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) StageVoid(stage *Stage) {
	bottomendarcshapev2grid.Stage(stage)
}

// Checkout bottomendarcshapev2grid to the back repo (if it is already staged)
func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) Checkout(stage *Stage) *BottomEndArcShapeV2Grid {
	if _, ok := stage.BottomEndArcShapeV2Grids[bottomendarcshapev2grid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBottomEndArcShapeV2Grid(bottomendarcshapev2grid)
		}
	}
	return bottomendarcshapev2grid
}

// for satisfaction of GongStruct interface
func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) GetName() (res string) {
	return bottomendarcshapev2grid.Name
}

// for satisfaction of GongStruct interface
func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) SetName(name string) {
	bottomendarcshapev2grid.Name = name
}

// Stage puts bottomstackgrowthcurveendarcshapev2 to the model stage
func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) Stage(stage *Stage) *BottomStackGrowthCurveEndArcShapeV2 {
	if _, ok := stage.BottomStackGrowthCurveEndArcShapeV2s[bottomstackgrowthcurveendarcshapev2]; !ok {
		stage.BottomStackGrowthCurveEndArcShapeV2s[bottomstackgrowthcurveendarcshapev2] = struct{}{}
		stage.BottomStackGrowthCurveEndArcShapeV2_stagedOrder[bottomstackgrowthcurveendarcshapev2] = stage.BottomStackGrowthCurveEndArcShapeV2Order
		stage.BottomStackGrowthCurveEndArcShapeV2_orderStaged[stage.BottomStackGrowthCurveEndArcShapeV2Order] = bottomstackgrowthcurveendarcshapev2
		stage.BottomStackGrowthCurveEndArcShapeV2Order++
	}
	stage.BottomStackGrowthCurveEndArcShapeV2s_mapString[bottomstackgrowthcurveendarcshapev2.Name] = bottomstackgrowthcurveendarcshapev2

	return bottomstackgrowthcurveendarcshapev2
}

// StagePreserveOrder puts bottomstackgrowthcurveendarcshapev2 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BottomStackGrowthCurveEndArcShapeV2Order
// - update stage.BottomStackGrowthCurveEndArcShapeV2Order accordingly
func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.BottomStackGrowthCurveEndArcShapeV2s[bottomstackgrowthcurveendarcshapev2]; !ok {
		stage.BottomStackGrowthCurveEndArcShapeV2s[bottomstackgrowthcurveendarcshapev2] = struct{}{}

		if order > stage.BottomStackGrowthCurveEndArcShapeV2Order {
			stage.BottomStackGrowthCurveEndArcShapeV2Order = order
		}
		stage.BottomStackGrowthCurveEndArcShapeV2_stagedOrder[bottomstackgrowthcurveendarcshapev2] = order
		stage.BottomStackGrowthCurveEndArcShapeV2_orderStaged[order] = bottomstackgrowthcurveendarcshapev2
		stage.BottomStackGrowthCurveEndArcShapeV2Order++
	}
	stage.BottomStackGrowthCurveEndArcShapeV2s_mapString[bottomstackgrowthcurveendarcshapev2.Name] = bottomstackgrowthcurveendarcshapev2
}

// Unstage removes bottomstackgrowthcurveendarcshapev2 off the model stage
func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) Unstage(stage *Stage) *BottomStackGrowthCurveEndArcShapeV2 {
	delete(stage.BottomStackGrowthCurveEndArcShapeV2s, bottomstackgrowthcurveendarcshapev2)
	// issue1150
	// delete(stage.BottomStackGrowthCurveEndArcShapeV2_stagedOrder, bottomstackgrowthcurveendarcshapev2)
	delete(stage.BottomStackGrowthCurveEndArcShapeV2s_mapString, bottomstackgrowthcurveendarcshapev2.Name)

	return bottomstackgrowthcurveendarcshapev2
}

// UnstageVoid removes bottomstackgrowthcurveendarcshapev2 off the model stage
func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) UnstageVoid(stage *Stage) {
	delete(stage.BottomStackGrowthCurveEndArcShapeV2s, bottomstackgrowthcurveendarcshapev2)
	// issue1150
	// delete(stage.BottomStackGrowthCurveEndArcShapeV2_stagedOrder, bottomstackgrowthcurveendarcshapev2)
	delete(stage.BottomStackGrowthCurveEndArcShapeV2s_mapString, bottomstackgrowthcurveendarcshapev2.Name)
}

// commit bottomstackgrowthcurveendarcshapev2 to the back repo (if it is already staged)
func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) Commit(stage *Stage) *BottomStackGrowthCurveEndArcShapeV2 {
	if _, ok := stage.BottomStackGrowthCurveEndArcShapeV2s[bottomstackgrowthcurveendarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBottomStackGrowthCurveEndArcShapeV2(bottomstackgrowthcurveendarcshapev2)
		}
	}
	return bottomstackgrowthcurveendarcshapev2
}

func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) CommitVoid(stage *Stage) {
	bottomstackgrowthcurveendarcshapev2.Commit(stage)
}

func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) StageVoid(stage *Stage) {
	bottomstackgrowthcurveendarcshapev2.Stage(stage)
}

// Checkout bottomstackgrowthcurveendarcshapev2 to the back repo (if it is already staged)
func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) Checkout(stage *Stage) *BottomStackGrowthCurveEndArcShapeV2 {
	if _, ok := stage.BottomStackGrowthCurveEndArcShapeV2s[bottomstackgrowthcurveendarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBottomStackGrowthCurveEndArcShapeV2(bottomstackgrowthcurveendarcshapev2)
		}
	}
	return bottomstackgrowthcurveendarcshapev2
}

// for satisfaction of GongStruct interface
func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) GetName() (res string) {
	return bottomstackgrowthcurveendarcshapev2.Name
}

// for satisfaction of GongStruct interface
func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) SetName(name string) {
	bottomstackgrowthcurveendarcshapev2.Name = name
}

// Stage puts bottomstackgrowthcurvestartarcshapev2 to the model stage
func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) Stage(stage *Stage) *BottomStackGrowthCurveStartArcShapeV2 {
	if _, ok := stage.BottomStackGrowthCurveStartArcShapeV2s[bottomstackgrowthcurvestartarcshapev2]; !ok {
		stage.BottomStackGrowthCurveStartArcShapeV2s[bottomstackgrowthcurvestartarcshapev2] = struct{}{}
		stage.BottomStackGrowthCurveStartArcShapeV2_stagedOrder[bottomstackgrowthcurvestartarcshapev2] = stage.BottomStackGrowthCurveStartArcShapeV2Order
		stage.BottomStackGrowthCurveStartArcShapeV2_orderStaged[stage.BottomStackGrowthCurveStartArcShapeV2Order] = bottomstackgrowthcurvestartarcshapev2
		stage.BottomStackGrowthCurveStartArcShapeV2Order++
	}
	stage.BottomStackGrowthCurveStartArcShapeV2s_mapString[bottomstackgrowthcurvestartarcshapev2.Name] = bottomstackgrowthcurvestartarcshapev2

	return bottomstackgrowthcurvestartarcshapev2
}

// StagePreserveOrder puts bottomstackgrowthcurvestartarcshapev2 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BottomStackGrowthCurveStartArcShapeV2Order
// - update stage.BottomStackGrowthCurveStartArcShapeV2Order accordingly
func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.BottomStackGrowthCurveStartArcShapeV2s[bottomstackgrowthcurvestartarcshapev2]; !ok {
		stage.BottomStackGrowthCurveStartArcShapeV2s[bottomstackgrowthcurvestartarcshapev2] = struct{}{}

		if order > stage.BottomStackGrowthCurveStartArcShapeV2Order {
			stage.BottomStackGrowthCurveStartArcShapeV2Order = order
		}
		stage.BottomStackGrowthCurveStartArcShapeV2_stagedOrder[bottomstackgrowthcurvestartarcshapev2] = order
		stage.BottomStackGrowthCurveStartArcShapeV2_orderStaged[order] = bottomstackgrowthcurvestartarcshapev2
		stage.BottomStackGrowthCurveStartArcShapeV2Order++
	}
	stage.BottomStackGrowthCurveStartArcShapeV2s_mapString[bottomstackgrowthcurvestartarcshapev2.Name] = bottomstackgrowthcurvestartarcshapev2
}

// Unstage removes bottomstackgrowthcurvestartarcshapev2 off the model stage
func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) Unstage(stage *Stage) *BottomStackGrowthCurveStartArcShapeV2 {
	delete(stage.BottomStackGrowthCurveStartArcShapeV2s, bottomstackgrowthcurvestartarcshapev2)
	// issue1150
	// delete(stage.BottomStackGrowthCurveStartArcShapeV2_stagedOrder, bottomstackgrowthcurvestartarcshapev2)
	delete(stage.BottomStackGrowthCurveStartArcShapeV2s_mapString, bottomstackgrowthcurvestartarcshapev2.Name)

	return bottomstackgrowthcurvestartarcshapev2
}

// UnstageVoid removes bottomstackgrowthcurvestartarcshapev2 off the model stage
func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) UnstageVoid(stage *Stage) {
	delete(stage.BottomStackGrowthCurveStartArcShapeV2s, bottomstackgrowthcurvestartarcshapev2)
	// issue1150
	// delete(stage.BottomStackGrowthCurveStartArcShapeV2_stagedOrder, bottomstackgrowthcurvestartarcshapev2)
	delete(stage.BottomStackGrowthCurveStartArcShapeV2s_mapString, bottomstackgrowthcurvestartarcshapev2.Name)
}

// commit bottomstackgrowthcurvestartarcshapev2 to the back repo (if it is already staged)
func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) Commit(stage *Stage) *BottomStackGrowthCurveStartArcShapeV2 {
	if _, ok := stage.BottomStackGrowthCurveStartArcShapeV2s[bottomstackgrowthcurvestartarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBottomStackGrowthCurveStartArcShapeV2(bottomstackgrowthcurvestartarcshapev2)
		}
	}
	return bottomstackgrowthcurvestartarcshapev2
}

func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) CommitVoid(stage *Stage) {
	bottomstackgrowthcurvestartarcshapev2.Commit(stage)
}

func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) StageVoid(stage *Stage) {
	bottomstackgrowthcurvestartarcshapev2.Stage(stage)
}

// Checkout bottomstackgrowthcurvestartarcshapev2 to the back repo (if it is already staged)
func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) Checkout(stage *Stage) *BottomStackGrowthCurveStartArcShapeV2 {
	if _, ok := stage.BottomStackGrowthCurveStartArcShapeV2s[bottomstackgrowthcurvestartarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBottomStackGrowthCurveStartArcShapeV2(bottomstackgrowthcurvestartarcshapev2)
		}
	}
	return bottomstackgrowthcurvestartarcshapev2
}

// for satisfaction of GongStruct interface
func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) GetName() (res string) {
	return bottomstackgrowthcurvestartarcshapev2.Name
}

// for satisfaction of GongStruct interface
func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) SetName(name string) {
	bottomstackgrowthcurvestartarcshapev2.Name = name
}

// Stage puts bottomstackofgrowthcurvev2 to the model stage
func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) Stage(stage *Stage) *BottomStackOfGrowthCurveV2 {
	if _, ok := stage.BottomStackOfGrowthCurveV2s[bottomstackofgrowthcurvev2]; !ok {
		stage.BottomStackOfGrowthCurveV2s[bottomstackofgrowthcurvev2] = struct{}{}
		stage.BottomStackOfGrowthCurveV2_stagedOrder[bottomstackofgrowthcurvev2] = stage.BottomStackOfGrowthCurveV2Order
		stage.BottomStackOfGrowthCurveV2_orderStaged[stage.BottomStackOfGrowthCurveV2Order] = bottomstackofgrowthcurvev2
		stage.BottomStackOfGrowthCurveV2Order++
	}
	stage.BottomStackOfGrowthCurveV2s_mapString[bottomstackofgrowthcurvev2.Name] = bottomstackofgrowthcurvev2

	return bottomstackofgrowthcurvev2
}

// StagePreserveOrder puts bottomstackofgrowthcurvev2 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BottomStackOfGrowthCurveV2Order
// - update stage.BottomStackOfGrowthCurveV2Order accordingly
func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.BottomStackOfGrowthCurveV2s[bottomstackofgrowthcurvev2]; !ok {
		stage.BottomStackOfGrowthCurveV2s[bottomstackofgrowthcurvev2] = struct{}{}

		if order > stage.BottomStackOfGrowthCurveV2Order {
			stage.BottomStackOfGrowthCurveV2Order = order
		}
		stage.BottomStackOfGrowthCurveV2_stagedOrder[bottomstackofgrowthcurvev2] = order
		stage.BottomStackOfGrowthCurveV2_orderStaged[order] = bottomstackofgrowthcurvev2
		stage.BottomStackOfGrowthCurveV2Order++
	}
	stage.BottomStackOfGrowthCurveV2s_mapString[bottomstackofgrowthcurvev2.Name] = bottomstackofgrowthcurvev2
}

// Unstage removes bottomstackofgrowthcurvev2 off the model stage
func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) Unstage(stage *Stage) *BottomStackOfGrowthCurveV2 {
	delete(stage.BottomStackOfGrowthCurveV2s, bottomstackofgrowthcurvev2)
	// issue1150
	// delete(stage.BottomStackOfGrowthCurveV2_stagedOrder, bottomstackofgrowthcurvev2)
	delete(stage.BottomStackOfGrowthCurveV2s_mapString, bottomstackofgrowthcurvev2.Name)

	return bottomstackofgrowthcurvev2
}

// UnstageVoid removes bottomstackofgrowthcurvev2 off the model stage
func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) UnstageVoid(stage *Stage) {
	delete(stage.BottomStackOfGrowthCurveV2s, bottomstackofgrowthcurvev2)
	// issue1150
	// delete(stage.BottomStackOfGrowthCurveV2_stagedOrder, bottomstackofgrowthcurvev2)
	delete(stage.BottomStackOfGrowthCurveV2s_mapString, bottomstackofgrowthcurvev2.Name)
}

// commit bottomstackofgrowthcurvev2 to the back repo (if it is already staged)
func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) Commit(stage *Stage) *BottomStackOfGrowthCurveV2 {
	if _, ok := stage.BottomStackOfGrowthCurveV2s[bottomstackofgrowthcurvev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBottomStackOfGrowthCurveV2(bottomstackofgrowthcurvev2)
		}
	}
	return bottomstackofgrowthcurvev2
}

func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) CommitVoid(stage *Stage) {
	bottomstackofgrowthcurvev2.Commit(stage)
}

func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) StageVoid(stage *Stage) {
	bottomstackofgrowthcurvev2.Stage(stage)
}

// Checkout bottomstackofgrowthcurvev2 to the back repo (if it is already staged)
func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) Checkout(stage *Stage) *BottomStackOfGrowthCurveV2 {
	if _, ok := stage.BottomStackOfGrowthCurveV2s[bottomstackofgrowthcurvev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBottomStackOfGrowthCurveV2(bottomstackofgrowthcurvev2)
		}
	}
	return bottomstackofgrowthcurvev2
}

// for satisfaction of GongStruct interface
func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) GetName() (res string) {
	return bottomstackofgrowthcurvev2.Name
}

// for satisfaction of GongStruct interface
func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) SetName(name string) {
	bottomstackofgrowthcurvev2.Name = name
}

// Stage puts bottomstartarcshapev2 to the model stage
func (bottomstartarcshapev2 *BottomStartArcShapeV2) Stage(stage *Stage) *BottomStartArcShapeV2 {
	if _, ok := stage.BottomStartArcShapeV2s[bottomstartarcshapev2]; !ok {
		stage.BottomStartArcShapeV2s[bottomstartarcshapev2] = struct{}{}
		stage.BottomStartArcShapeV2_stagedOrder[bottomstartarcshapev2] = stage.BottomStartArcShapeV2Order
		stage.BottomStartArcShapeV2_orderStaged[stage.BottomStartArcShapeV2Order] = bottomstartarcshapev2
		stage.BottomStartArcShapeV2Order++
	}
	stage.BottomStartArcShapeV2s_mapString[bottomstartarcshapev2.Name] = bottomstartarcshapev2

	return bottomstartarcshapev2
}

// StagePreserveOrder puts bottomstartarcshapev2 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BottomStartArcShapeV2Order
// - update stage.BottomStartArcShapeV2Order accordingly
func (bottomstartarcshapev2 *BottomStartArcShapeV2) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.BottomStartArcShapeV2s[bottomstartarcshapev2]; !ok {
		stage.BottomStartArcShapeV2s[bottomstartarcshapev2] = struct{}{}

		if order > stage.BottomStartArcShapeV2Order {
			stage.BottomStartArcShapeV2Order = order
		}
		stage.BottomStartArcShapeV2_stagedOrder[bottomstartarcshapev2] = order
		stage.BottomStartArcShapeV2_orderStaged[order] = bottomstartarcshapev2
		stage.BottomStartArcShapeV2Order++
	}
	stage.BottomStartArcShapeV2s_mapString[bottomstartarcshapev2.Name] = bottomstartarcshapev2
}

// Unstage removes bottomstartarcshapev2 off the model stage
func (bottomstartarcshapev2 *BottomStartArcShapeV2) Unstage(stage *Stage) *BottomStartArcShapeV2 {
	delete(stage.BottomStartArcShapeV2s, bottomstartarcshapev2)
	// issue1150
	// delete(stage.BottomStartArcShapeV2_stagedOrder, bottomstartarcshapev2)
	delete(stage.BottomStartArcShapeV2s_mapString, bottomstartarcshapev2.Name)

	return bottomstartarcshapev2
}

// UnstageVoid removes bottomstartarcshapev2 off the model stage
func (bottomstartarcshapev2 *BottomStartArcShapeV2) UnstageVoid(stage *Stage) {
	delete(stage.BottomStartArcShapeV2s, bottomstartarcshapev2)
	// issue1150
	// delete(stage.BottomStartArcShapeV2_stagedOrder, bottomstartarcshapev2)
	delete(stage.BottomStartArcShapeV2s_mapString, bottomstartarcshapev2.Name)
}

// commit bottomstartarcshapev2 to the back repo (if it is already staged)
func (bottomstartarcshapev2 *BottomStartArcShapeV2) Commit(stage *Stage) *BottomStartArcShapeV2 {
	if _, ok := stage.BottomStartArcShapeV2s[bottomstartarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBottomStartArcShapeV2(bottomstartarcshapev2)
		}
	}
	return bottomstartarcshapev2
}

func (bottomstartarcshapev2 *BottomStartArcShapeV2) CommitVoid(stage *Stage) {
	bottomstartarcshapev2.Commit(stage)
}

func (bottomstartarcshapev2 *BottomStartArcShapeV2) StageVoid(stage *Stage) {
	bottomstartarcshapev2.Stage(stage)
}

// Checkout bottomstartarcshapev2 to the back repo (if it is already staged)
func (bottomstartarcshapev2 *BottomStartArcShapeV2) Checkout(stage *Stage) *BottomStartArcShapeV2 {
	if _, ok := stage.BottomStartArcShapeV2s[bottomstartarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBottomStartArcShapeV2(bottomstartarcshapev2)
		}
	}
	return bottomstartarcshapev2
}

// for satisfaction of GongStruct interface
func (bottomstartarcshapev2 *BottomStartArcShapeV2) GetName() (res string) {
	return bottomstartarcshapev2.Name
}

// for satisfaction of GongStruct interface
func (bottomstartarcshapev2 *BottomStartArcShapeV2) SetName(name string) {
	bottomstartarcshapev2.Name = name
}

// Stage puts bottomstartarcshapev2grid to the model stage
func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) Stage(stage *Stage) *BottomStartArcShapeV2Grid {
	if _, ok := stage.BottomStartArcShapeV2Grids[bottomstartarcshapev2grid]; !ok {
		stage.BottomStartArcShapeV2Grids[bottomstartarcshapev2grid] = struct{}{}
		stage.BottomStartArcShapeV2Grid_stagedOrder[bottomstartarcshapev2grid] = stage.BottomStartArcShapeV2GridOrder
		stage.BottomStartArcShapeV2Grid_orderStaged[stage.BottomStartArcShapeV2GridOrder] = bottomstartarcshapev2grid
		stage.BottomStartArcShapeV2GridOrder++
	}
	stage.BottomStartArcShapeV2Grids_mapString[bottomstartarcshapev2grid.Name] = bottomstartarcshapev2grid

	return bottomstartarcshapev2grid
}

// StagePreserveOrder puts bottomstartarcshapev2grid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BottomStartArcShapeV2GridOrder
// - update stage.BottomStartArcShapeV2GridOrder accordingly
func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.BottomStartArcShapeV2Grids[bottomstartarcshapev2grid]; !ok {
		stage.BottomStartArcShapeV2Grids[bottomstartarcshapev2grid] = struct{}{}

		if order > stage.BottomStartArcShapeV2GridOrder {
			stage.BottomStartArcShapeV2GridOrder = order
		}
		stage.BottomStartArcShapeV2Grid_stagedOrder[bottomstartarcshapev2grid] = order
		stage.BottomStartArcShapeV2Grid_orderStaged[order] = bottomstartarcshapev2grid
		stage.BottomStartArcShapeV2GridOrder++
	}
	stage.BottomStartArcShapeV2Grids_mapString[bottomstartarcshapev2grid.Name] = bottomstartarcshapev2grid
}

// Unstage removes bottomstartarcshapev2grid off the model stage
func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) Unstage(stage *Stage) *BottomStartArcShapeV2Grid {
	delete(stage.BottomStartArcShapeV2Grids, bottomstartarcshapev2grid)
	// issue1150
	// delete(stage.BottomStartArcShapeV2Grid_stagedOrder, bottomstartarcshapev2grid)
	delete(stage.BottomStartArcShapeV2Grids_mapString, bottomstartarcshapev2grid.Name)

	return bottomstartarcshapev2grid
}

// UnstageVoid removes bottomstartarcshapev2grid off the model stage
func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) UnstageVoid(stage *Stage) {
	delete(stage.BottomStartArcShapeV2Grids, bottomstartarcshapev2grid)
	// issue1150
	// delete(stage.BottomStartArcShapeV2Grid_stagedOrder, bottomstartarcshapev2grid)
	delete(stage.BottomStartArcShapeV2Grids_mapString, bottomstartarcshapev2grid.Name)
}

// commit bottomstartarcshapev2grid to the back repo (if it is already staged)
func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) Commit(stage *Stage) *BottomStartArcShapeV2Grid {
	if _, ok := stage.BottomStartArcShapeV2Grids[bottomstartarcshapev2grid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBottomStartArcShapeV2Grid(bottomstartarcshapev2grid)
		}
	}
	return bottomstartarcshapev2grid
}

func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) CommitVoid(stage *Stage) {
	bottomstartarcshapev2grid.Commit(stage)
}

func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) StageVoid(stage *Stage) {
	bottomstartarcshapev2grid.Stage(stage)
}

// Checkout bottomstartarcshapev2grid to the back repo (if it is already staged)
func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) Checkout(stage *Stage) *BottomStartArcShapeV2Grid {
	if _, ok := stage.BottomStartArcShapeV2Grids[bottomstartarcshapev2grid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBottomStartArcShapeV2Grid(bottomstartarcshapev2grid)
		}
	}
	return bottomstartarcshapev2grid
}

// for satisfaction of GongStruct interface
func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) GetName() (res string) {
	return bottomstartarcshapev2grid.Name
}

// for satisfaction of GongStruct interface
func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) SetName(name string) {
	bottomstartarcshapev2grid.Name = name
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

// Stage puts endarcshapev2 to the model stage
func (endarcshapev2 *EndArcShapeV2) Stage(stage *Stage) *EndArcShapeV2 {
	if _, ok := stage.EndArcShapeV2s[endarcshapev2]; !ok {
		stage.EndArcShapeV2s[endarcshapev2] = struct{}{}
		stage.EndArcShapeV2_stagedOrder[endarcshapev2] = stage.EndArcShapeV2Order
		stage.EndArcShapeV2_orderStaged[stage.EndArcShapeV2Order] = endarcshapev2
		stage.EndArcShapeV2Order++
	}
	stage.EndArcShapeV2s_mapString[endarcshapev2.Name] = endarcshapev2

	return endarcshapev2
}

// StagePreserveOrder puts endarcshapev2 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EndArcShapeV2Order
// - update stage.EndArcShapeV2Order accordingly
func (endarcshapev2 *EndArcShapeV2) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.EndArcShapeV2s[endarcshapev2]; !ok {
		stage.EndArcShapeV2s[endarcshapev2] = struct{}{}

		if order > stage.EndArcShapeV2Order {
			stage.EndArcShapeV2Order = order
		}
		stage.EndArcShapeV2_stagedOrder[endarcshapev2] = order
		stage.EndArcShapeV2_orderStaged[order] = endarcshapev2
		stage.EndArcShapeV2Order++
	}
	stage.EndArcShapeV2s_mapString[endarcshapev2.Name] = endarcshapev2
}

// Unstage removes endarcshapev2 off the model stage
func (endarcshapev2 *EndArcShapeV2) Unstage(stage *Stage) *EndArcShapeV2 {
	delete(stage.EndArcShapeV2s, endarcshapev2)
	// issue1150
	// delete(stage.EndArcShapeV2_stagedOrder, endarcshapev2)
	delete(stage.EndArcShapeV2s_mapString, endarcshapev2.Name)

	return endarcshapev2
}

// UnstageVoid removes endarcshapev2 off the model stage
func (endarcshapev2 *EndArcShapeV2) UnstageVoid(stage *Stage) {
	delete(stage.EndArcShapeV2s, endarcshapev2)
	// issue1150
	// delete(stage.EndArcShapeV2_stagedOrder, endarcshapev2)
	delete(stage.EndArcShapeV2s_mapString, endarcshapev2.Name)
}

// commit endarcshapev2 to the back repo (if it is already staged)
func (endarcshapev2 *EndArcShapeV2) Commit(stage *Stage) *EndArcShapeV2 {
	if _, ok := stage.EndArcShapeV2s[endarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitEndArcShapeV2(endarcshapev2)
		}
	}
	return endarcshapev2
}

func (endarcshapev2 *EndArcShapeV2) CommitVoid(stage *Stage) {
	endarcshapev2.Commit(stage)
}

func (endarcshapev2 *EndArcShapeV2) StageVoid(stage *Stage) {
	endarcshapev2.Stage(stage)
}

// Checkout endarcshapev2 to the back repo (if it is already staged)
func (endarcshapev2 *EndArcShapeV2) Checkout(stage *Stage) *EndArcShapeV2 {
	if _, ok := stage.EndArcShapeV2s[endarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutEndArcShapeV2(endarcshapev2)
		}
	}
	return endarcshapev2
}

// for satisfaction of GongStruct interface
func (endarcshapev2 *EndArcShapeV2) GetName() (res string) {
	return endarcshapev2.Name
}

// for satisfaction of GongStruct interface
func (endarcshapev2 *EndArcShapeV2) SetName(name string) {
	endarcshapev2.Name = name
}

// Stage puts endarcshapev2grid to the model stage
func (endarcshapev2grid *EndArcShapeV2Grid) Stage(stage *Stage) *EndArcShapeV2Grid {
	if _, ok := stage.EndArcShapeV2Grids[endarcshapev2grid]; !ok {
		stage.EndArcShapeV2Grids[endarcshapev2grid] = struct{}{}
		stage.EndArcShapeV2Grid_stagedOrder[endarcshapev2grid] = stage.EndArcShapeV2GridOrder
		stage.EndArcShapeV2Grid_orderStaged[stage.EndArcShapeV2GridOrder] = endarcshapev2grid
		stage.EndArcShapeV2GridOrder++
	}
	stage.EndArcShapeV2Grids_mapString[endarcshapev2grid.Name] = endarcshapev2grid

	return endarcshapev2grid
}

// StagePreserveOrder puts endarcshapev2grid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EndArcShapeV2GridOrder
// - update stage.EndArcShapeV2GridOrder accordingly
func (endarcshapev2grid *EndArcShapeV2Grid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.EndArcShapeV2Grids[endarcshapev2grid]; !ok {
		stage.EndArcShapeV2Grids[endarcshapev2grid] = struct{}{}

		if order > stage.EndArcShapeV2GridOrder {
			stage.EndArcShapeV2GridOrder = order
		}
		stage.EndArcShapeV2Grid_stagedOrder[endarcshapev2grid] = order
		stage.EndArcShapeV2Grid_orderStaged[order] = endarcshapev2grid
		stage.EndArcShapeV2GridOrder++
	}
	stage.EndArcShapeV2Grids_mapString[endarcshapev2grid.Name] = endarcshapev2grid
}

// Unstage removes endarcshapev2grid off the model stage
func (endarcshapev2grid *EndArcShapeV2Grid) Unstage(stage *Stage) *EndArcShapeV2Grid {
	delete(stage.EndArcShapeV2Grids, endarcshapev2grid)
	// issue1150
	// delete(stage.EndArcShapeV2Grid_stagedOrder, endarcshapev2grid)
	delete(stage.EndArcShapeV2Grids_mapString, endarcshapev2grid.Name)

	return endarcshapev2grid
}

// UnstageVoid removes endarcshapev2grid off the model stage
func (endarcshapev2grid *EndArcShapeV2Grid) UnstageVoid(stage *Stage) {
	delete(stage.EndArcShapeV2Grids, endarcshapev2grid)
	// issue1150
	// delete(stage.EndArcShapeV2Grid_stagedOrder, endarcshapev2grid)
	delete(stage.EndArcShapeV2Grids_mapString, endarcshapev2grid.Name)
}

// commit endarcshapev2grid to the back repo (if it is already staged)
func (endarcshapev2grid *EndArcShapeV2Grid) Commit(stage *Stage) *EndArcShapeV2Grid {
	if _, ok := stage.EndArcShapeV2Grids[endarcshapev2grid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitEndArcShapeV2Grid(endarcshapev2grid)
		}
	}
	return endarcshapev2grid
}

func (endarcshapev2grid *EndArcShapeV2Grid) CommitVoid(stage *Stage) {
	endarcshapev2grid.Commit(stage)
}

func (endarcshapev2grid *EndArcShapeV2Grid) StageVoid(stage *Stage) {
	endarcshapev2grid.Stage(stage)
}

// Checkout endarcshapev2grid to the back repo (if it is already staged)
func (endarcshapev2grid *EndArcShapeV2Grid) Checkout(stage *Stage) *EndArcShapeV2Grid {
	if _, ok := stage.EndArcShapeV2Grids[endarcshapev2grid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutEndArcShapeV2Grid(endarcshapev2grid)
		}
	}
	return endarcshapev2grid
}

// for satisfaction of GongStruct interface
func (endarcshapev2grid *EndArcShapeV2Grid) GetName() (res string) {
	return endarcshapev2grid.Name
}

// for satisfaction of GongStruct interface
func (endarcshapev2grid *EndArcShapeV2Grid) SetName(name string) {
	endarcshapev2grid.Name = name
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

// Stage puts stackgrowthcurveendarcshapev2 to the model stage
func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) Stage(stage *Stage) *StackGrowthCurveEndArcShapeV2 {
	if _, ok := stage.StackGrowthCurveEndArcShapeV2s[stackgrowthcurveendarcshapev2]; !ok {
		stage.StackGrowthCurveEndArcShapeV2s[stackgrowthcurveendarcshapev2] = struct{}{}
		stage.StackGrowthCurveEndArcShapeV2_stagedOrder[stackgrowthcurveendarcshapev2] = stage.StackGrowthCurveEndArcShapeV2Order
		stage.StackGrowthCurveEndArcShapeV2_orderStaged[stage.StackGrowthCurveEndArcShapeV2Order] = stackgrowthcurveendarcshapev2
		stage.StackGrowthCurveEndArcShapeV2Order++
	}
	stage.StackGrowthCurveEndArcShapeV2s_mapString[stackgrowthcurveendarcshapev2.Name] = stackgrowthcurveendarcshapev2

	return stackgrowthcurveendarcshapev2
}

// StagePreserveOrder puts stackgrowthcurveendarcshapev2 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StackGrowthCurveEndArcShapeV2Order
// - update stage.StackGrowthCurveEndArcShapeV2Order accordingly
func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.StackGrowthCurveEndArcShapeV2s[stackgrowthcurveendarcshapev2]; !ok {
		stage.StackGrowthCurveEndArcShapeV2s[stackgrowthcurveendarcshapev2] = struct{}{}

		if order > stage.StackGrowthCurveEndArcShapeV2Order {
			stage.StackGrowthCurveEndArcShapeV2Order = order
		}
		stage.StackGrowthCurveEndArcShapeV2_stagedOrder[stackgrowthcurveendarcshapev2] = order
		stage.StackGrowthCurveEndArcShapeV2_orderStaged[order] = stackgrowthcurveendarcshapev2
		stage.StackGrowthCurveEndArcShapeV2Order++
	}
	stage.StackGrowthCurveEndArcShapeV2s_mapString[stackgrowthcurveendarcshapev2.Name] = stackgrowthcurveendarcshapev2
}

// Unstage removes stackgrowthcurveendarcshapev2 off the model stage
func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) Unstage(stage *Stage) *StackGrowthCurveEndArcShapeV2 {
	delete(stage.StackGrowthCurveEndArcShapeV2s, stackgrowthcurveendarcshapev2)
	// issue1150
	// delete(stage.StackGrowthCurveEndArcShapeV2_stagedOrder, stackgrowthcurveendarcshapev2)
	delete(stage.StackGrowthCurveEndArcShapeV2s_mapString, stackgrowthcurveendarcshapev2.Name)

	return stackgrowthcurveendarcshapev2
}

// UnstageVoid removes stackgrowthcurveendarcshapev2 off the model stage
func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) UnstageVoid(stage *Stage) {
	delete(stage.StackGrowthCurveEndArcShapeV2s, stackgrowthcurveendarcshapev2)
	// issue1150
	// delete(stage.StackGrowthCurveEndArcShapeV2_stagedOrder, stackgrowthcurveendarcshapev2)
	delete(stage.StackGrowthCurveEndArcShapeV2s_mapString, stackgrowthcurveendarcshapev2.Name)
}

// commit stackgrowthcurveendarcshapev2 to the back repo (if it is already staged)
func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) Commit(stage *Stage) *StackGrowthCurveEndArcShapeV2 {
	if _, ok := stage.StackGrowthCurveEndArcShapeV2s[stackgrowthcurveendarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitStackGrowthCurveEndArcShapeV2(stackgrowthcurveendarcshapev2)
		}
	}
	return stackgrowthcurveendarcshapev2
}

func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) CommitVoid(stage *Stage) {
	stackgrowthcurveendarcshapev2.Commit(stage)
}

func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) StageVoid(stage *Stage) {
	stackgrowthcurveendarcshapev2.Stage(stage)
}

// Checkout stackgrowthcurveendarcshapev2 to the back repo (if it is already staged)
func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) Checkout(stage *Stage) *StackGrowthCurveEndArcShapeV2 {
	if _, ok := stage.StackGrowthCurveEndArcShapeV2s[stackgrowthcurveendarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutStackGrowthCurveEndArcShapeV2(stackgrowthcurveendarcshapev2)
		}
	}
	return stackgrowthcurveendarcshapev2
}

// for satisfaction of GongStruct interface
func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) GetName() (res string) {
	return stackgrowthcurveendarcshapev2.Name
}

// for satisfaction of GongStruct interface
func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) SetName(name string) {
	stackgrowthcurveendarcshapev2.Name = name
}

// Stage puts stackgrowthcurvestartarcshapev2 to the model stage
func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) Stage(stage *Stage) *StackGrowthCurveStartArcShapeV2 {
	if _, ok := stage.StackGrowthCurveStartArcShapeV2s[stackgrowthcurvestartarcshapev2]; !ok {
		stage.StackGrowthCurveStartArcShapeV2s[stackgrowthcurvestartarcshapev2] = struct{}{}
		stage.StackGrowthCurveStartArcShapeV2_stagedOrder[stackgrowthcurvestartarcshapev2] = stage.StackGrowthCurveStartArcShapeV2Order
		stage.StackGrowthCurveStartArcShapeV2_orderStaged[stage.StackGrowthCurveStartArcShapeV2Order] = stackgrowthcurvestartarcshapev2
		stage.StackGrowthCurveStartArcShapeV2Order++
	}
	stage.StackGrowthCurveStartArcShapeV2s_mapString[stackgrowthcurvestartarcshapev2.Name] = stackgrowthcurvestartarcshapev2

	return stackgrowthcurvestartarcshapev2
}

// StagePreserveOrder puts stackgrowthcurvestartarcshapev2 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StackGrowthCurveStartArcShapeV2Order
// - update stage.StackGrowthCurveStartArcShapeV2Order accordingly
func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.StackGrowthCurveStartArcShapeV2s[stackgrowthcurvestartarcshapev2]; !ok {
		stage.StackGrowthCurveStartArcShapeV2s[stackgrowthcurvestartarcshapev2] = struct{}{}

		if order > stage.StackGrowthCurveStartArcShapeV2Order {
			stage.StackGrowthCurveStartArcShapeV2Order = order
		}
		stage.StackGrowthCurveStartArcShapeV2_stagedOrder[stackgrowthcurvestartarcshapev2] = order
		stage.StackGrowthCurveStartArcShapeV2_orderStaged[order] = stackgrowthcurvestartarcshapev2
		stage.StackGrowthCurveStartArcShapeV2Order++
	}
	stage.StackGrowthCurveStartArcShapeV2s_mapString[stackgrowthcurvestartarcshapev2.Name] = stackgrowthcurvestartarcshapev2
}

// Unstage removes stackgrowthcurvestartarcshapev2 off the model stage
func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) Unstage(stage *Stage) *StackGrowthCurveStartArcShapeV2 {
	delete(stage.StackGrowthCurveStartArcShapeV2s, stackgrowthcurvestartarcshapev2)
	// issue1150
	// delete(stage.StackGrowthCurveStartArcShapeV2_stagedOrder, stackgrowthcurvestartarcshapev2)
	delete(stage.StackGrowthCurveStartArcShapeV2s_mapString, stackgrowthcurvestartarcshapev2.Name)

	return stackgrowthcurvestartarcshapev2
}

// UnstageVoid removes stackgrowthcurvestartarcshapev2 off the model stage
func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) UnstageVoid(stage *Stage) {
	delete(stage.StackGrowthCurveStartArcShapeV2s, stackgrowthcurvestartarcshapev2)
	// issue1150
	// delete(stage.StackGrowthCurveStartArcShapeV2_stagedOrder, stackgrowthcurvestartarcshapev2)
	delete(stage.StackGrowthCurveStartArcShapeV2s_mapString, stackgrowthcurvestartarcshapev2.Name)
}

// commit stackgrowthcurvestartarcshapev2 to the back repo (if it is already staged)
func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) Commit(stage *Stage) *StackGrowthCurveStartArcShapeV2 {
	if _, ok := stage.StackGrowthCurveStartArcShapeV2s[stackgrowthcurvestartarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitStackGrowthCurveStartArcShapeV2(stackgrowthcurvestartarcshapev2)
		}
	}
	return stackgrowthcurvestartarcshapev2
}

func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) CommitVoid(stage *Stage) {
	stackgrowthcurvestartarcshapev2.Commit(stage)
}

func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) StageVoid(stage *Stage) {
	stackgrowthcurvestartarcshapev2.Stage(stage)
}

// Checkout stackgrowthcurvestartarcshapev2 to the back repo (if it is already staged)
func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) Checkout(stage *Stage) *StackGrowthCurveStartArcShapeV2 {
	if _, ok := stage.StackGrowthCurveStartArcShapeV2s[stackgrowthcurvestartarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutStackGrowthCurveStartArcShapeV2(stackgrowthcurvestartarcshapev2)
		}
	}
	return stackgrowthcurvestartarcshapev2
}

// for satisfaction of GongStruct interface
func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) GetName() (res string) {
	return stackgrowthcurvestartarcshapev2.Name
}

// for satisfaction of GongStruct interface
func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) SetName(name string) {
	stackgrowthcurvestartarcshapev2.Name = name
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

// Stage puts stackofgrowthcurvev2 to the model stage
func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) Stage(stage *Stage) *StackOfGrowthCurveV2 {
	if _, ok := stage.StackOfGrowthCurveV2s[stackofgrowthcurvev2]; !ok {
		stage.StackOfGrowthCurveV2s[stackofgrowthcurvev2] = struct{}{}
		stage.StackOfGrowthCurveV2_stagedOrder[stackofgrowthcurvev2] = stage.StackOfGrowthCurveV2Order
		stage.StackOfGrowthCurveV2_orderStaged[stage.StackOfGrowthCurveV2Order] = stackofgrowthcurvev2
		stage.StackOfGrowthCurveV2Order++
	}
	stage.StackOfGrowthCurveV2s_mapString[stackofgrowthcurvev2.Name] = stackofgrowthcurvev2

	return stackofgrowthcurvev2
}

// StagePreserveOrder puts stackofgrowthcurvev2 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StackOfGrowthCurveV2Order
// - update stage.StackOfGrowthCurveV2Order accordingly
func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.StackOfGrowthCurveV2s[stackofgrowthcurvev2]; !ok {
		stage.StackOfGrowthCurveV2s[stackofgrowthcurvev2] = struct{}{}

		if order > stage.StackOfGrowthCurveV2Order {
			stage.StackOfGrowthCurveV2Order = order
		}
		stage.StackOfGrowthCurveV2_stagedOrder[stackofgrowthcurvev2] = order
		stage.StackOfGrowthCurveV2_orderStaged[order] = stackofgrowthcurvev2
		stage.StackOfGrowthCurveV2Order++
	}
	stage.StackOfGrowthCurveV2s_mapString[stackofgrowthcurvev2.Name] = stackofgrowthcurvev2
}

// Unstage removes stackofgrowthcurvev2 off the model stage
func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) Unstage(stage *Stage) *StackOfGrowthCurveV2 {
	delete(stage.StackOfGrowthCurveV2s, stackofgrowthcurvev2)
	// issue1150
	// delete(stage.StackOfGrowthCurveV2_stagedOrder, stackofgrowthcurvev2)
	delete(stage.StackOfGrowthCurveV2s_mapString, stackofgrowthcurvev2.Name)

	return stackofgrowthcurvev2
}

// UnstageVoid removes stackofgrowthcurvev2 off the model stage
func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) UnstageVoid(stage *Stage) {
	delete(stage.StackOfGrowthCurveV2s, stackofgrowthcurvev2)
	// issue1150
	// delete(stage.StackOfGrowthCurveV2_stagedOrder, stackofgrowthcurvev2)
	delete(stage.StackOfGrowthCurveV2s_mapString, stackofgrowthcurvev2.Name)
}

// commit stackofgrowthcurvev2 to the back repo (if it is already staged)
func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) Commit(stage *Stage) *StackOfGrowthCurveV2 {
	if _, ok := stage.StackOfGrowthCurveV2s[stackofgrowthcurvev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitStackOfGrowthCurveV2(stackofgrowthcurvev2)
		}
	}
	return stackofgrowthcurvev2
}

func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) CommitVoid(stage *Stage) {
	stackofgrowthcurvev2.Commit(stage)
}

func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) StageVoid(stage *Stage) {
	stackofgrowthcurvev2.Stage(stage)
}

// Checkout stackofgrowthcurvev2 to the back repo (if it is already staged)
func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) Checkout(stage *Stage) *StackOfGrowthCurveV2 {
	if _, ok := stage.StackOfGrowthCurveV2s[stackofgrowthcurvev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutStackOfGrowthCurveV2(stackofgrowthcurvev2)
		}
	}
	return stackofgrowthcurvev2
}

// for satisfaction of GongStruct interface
func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) GetName() (res string) {
	return stackofgrowthcurvev2.Name
}

// for satisfaction of GongStruct interface
func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) SetName(name string) {
	stackofgrowthcurvev2.Name = name
}

// Stage puts startarcshapev2 to the model stage
func (startarcshapev2 *StartArcShapeV2) Stage(stage *Stage) *StartArcShapeV2 {
	if _, ok := stage.StartArcShapeV2s[startarcshapev2]; !ok {
		stage.StartArcShapeV2s[startarcshapev2] = struct{}{}
		stage.StartArcShapeV2_stagedOrder[startarcshapev2] = stage.StartArcShapeV2Order
		stage.StartArcShapeV2_orderStaged[stage.StartArcShapeV2Order] = startarcshapev2
		stage.StartArcShapeV2Order++
	}
	stage.StartArcShapeV2s_mapString[startarcshapev2.Name] = startarcshapev2

	return startarcshapev2
}

// StagePreserveOrder puts startarcshapev2 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StartArcShapeV2Order
// - update stage.StartArcShapeV2Order accordingly
func (startarcshapev2 *StartArcShapeV2) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.StartArcShapeV2s[startarcshapev2]; !ok {
		stage.StartArcShapeV2s[startarcshapev2] = struct{}{}

		if order > stage.StartArcShapeV2Order {
			stage.StartArcShapeV2Order = order
		}
		stage.StartArcShapeV2_stagedOrder[startarcshapev2] = order
		stage.StartArcShapeV2_orderStaged[order] = startarcshapev2
		stage.StartArcShapeV2Order++
	}
	stage.StartArcShapeV2s_mapString[startarcshapev2.Name] = startarcshapev2
}

// Unstage removes startarcshapev2 off the model stage
func (startarcshapev2 *StartArcShapeV2) Unstage(stage *Stage) *StartArcShapeV2 {
	delete(stage.StartArcShapeV2s, startarcshapev2)
	// issue1150
	// delete(stage.StartArcShapeV2_stagedOrder, startarcshapev2)
	delete(stage.StartArcShapeV2s_mapString, startarcshapev2.Name)

	return startarcshapev2
}

// UnstageVoid removes startarcshapev2 off the model stage
func (startarcshapev2 *StartArcShapeV2) UnstageVoid(stage *Stage) {
	delete(stage.StartArcShapeV2s, startarcshapev2)
	// issue1150
	// delete(stage.StartArcShapeV2_stagedOrder, startarcshapev2)
	delete(stage.StartArcShapeV2s_mapString, startarcshapev2.Name)
}

// commit startarcshapev2 to the back repo (if it is already staged)
func (startarcshapev2 *StartArcShapeV2) Commit(stage *Stage) *StartArcShapeV2 {
	if _, ok := stage.StartArcShapeV2s[startarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitStartArcShapeV2(startarcshapev2)
		}
	}
	return startarcshapev2
}

func (startarcshapev2 *StartArcShapeV2) CommitVoid(stage *Stage) {
	startarcshapev2.Commit(stage)
}

func (startarcshapev2 *StartArcShapeV2) StageVoid(stage *Stage) {
	startarcshapev2.Stage(stage)
}

// Checkout startarcshapev2 to the back repo (if it is already staged)
func (startarcshapev2 *StartArcShapeV2) Checkout(stage *Stage) *StartArcShapeV2 {
	if _, ok := stage.StartArcShapeV2s[startarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutStartArcShapeV2(startarcshapev2)
		}
	}
	return startarcshapev2
}

// for satisfaction of GongStruct interface
func (startarcshapev2 *StartArcShapeV2) GetName() (res string) {
	return startarcshapev2.Name
}

// for satisfaction of GongStruct interface
func (startarcshapev2 *StartArcShapeV2) SetName(name string) {
	startarcshapev2.Name = name
}

// Stage puts startarcshapev2grid to the model stage
func (startarcshapev2grid *StartArcShapeV2Grid) Stage(stage *Stage) *StartArcShapeV2Grid {
	if _, ok := stage.StartArcShapeV2Grids[startarcshapev2grid]; !ok {
		stage.StartArcShapeV2Grids[startarcshapev2grid] = struct{}{}
		stage.StartArcShapeV2Grid_stagedOrder[startarcshapev2grid] = stage.StartArcShapeV2GridOrder
		stage.StartArcShapeV2Grid_orderStaged[stage.StartArcShapeV2GridOrder] = startarcshapev2grid
		stage.StartArcShapeV2GridOrder++
	}
	stage.StartArcShapeV2Grids_mapString[startarcshapev2grid.Name] = startarcshapev2grid

	return startarcshapev2grid
}

// StagePreserveOrder puts startarcshapev2grid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StartArcShapeV2GridOrder
// - update stage.StartArcShapeV2GridOrder accordingly
func (startarcshapev2grid *StartArcShapeV2Grid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.StartArcShapeV2Grids[startarcshapev2grid]; !ok {
		stage.StartArcShapeV2Grids[startarcshapev2grid] = struct{}{}

		if order > stage.StartArcShapeV2GridOrder {
			stage.StartArcShapeV2GridOrder = order
		}
		stage.StartArcShapeV2Grid_stagedOrder[startarcshapev2grid] = order
		stage.StartArcShapeV2Grid_orderStaged[order] = startarcshapev2grid
		stage.StartArcShapeV2GridOrder++
	}
	stage.StartArcShapeV2Grids_mapString[startarcshapev2grid.Name] = startarcshapev2grid
}

// Unstage removes startarcshapev2grid off the model stage
func (startarcshapev2grid *StartArcShapeV2Grid) Unstage(stage *Stage) *StartArcShapeV2Grid {
	delete(stage.StartArcShapeV2Grids, startarcshapev2grid)
	// issue1150
	// delete(stage.StartArcShapeV2Grid_stagedOrder, startarcshapev2grid)
	delete(stage.StartArcShapeV2Grids_mapString, startarcshapev2grid.Name)

	return startarcshapev2grid
}

// UnstageVoid removes startarcshapev2grid off the model stage
func (startarcshapev2grid *StartArcShapeV2Grid) UnstageVoid(stage *Stage) {
	delete(stage.StartArcShapeV2Grids, startarcshapev2grid)
	// issue1150
	// delete(stage.StartArcShapeV2Grid_stagedOrder, startarcshapev2grid)
	delete(stage.StartArcShapeV2Grids_mapString, startarcshapev2grid.Name)
}

// commit startarcshapev2grid to the back repo (if it is already staged)
func (startarcshapev2grid *StartArcShapeV2Grid) Commit(stage *Stage) *StartArcShapeV2Grid {
	if _, ok := stage.StartArcShapeV2Grids[startarcshapev2grid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitStartArcShapeV2Grid(startarcshapev2grid)
		}
	}
	return startarcshapev2grid
}

func (startarcshapev2grid *StartArcShapeV2Grid) CommitVoid(stage *Stage) {
	startarcshapev2grid.Commit(stage)
}

func (startarcshapev2grid *StartArcShapeV2Grid) StageVoid(stage *Stage) {
	startarcshapev2grid.Stage(stage)
}

// Checkout startarcshapev2grid to the back repo (if it is already staged)
func (startarcshapev2grid *StartArcShapeV2Grid) Checkout(stage *Stage) *StartArcShapeV2Grid {
	if _, ok := stage.StartArcShapeV2Grids[startarcshapev2grid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutStartArcShapeV2Grid(startarcshapev2grid)
		}
	}
	return startarcshapev2grid
}

// for satisfaction of GongStruct interface
func (startarcshapev2grid *StartArcShapeV2Grid) GetName() (res string) {
	return startarcshapev2grid.Name
}

// for satisfaction of GongStruct interface
func (startarcshapev2grid *StartArcShapeV2Grid) SetName(name string) {
	startarcshapev2grid.Name = name
}

// Stage puts topendarcshapev2 to the model stage
func (topendarcshapev2 *TopEndArcShapeV2) Stage(stage *Stage) *TopEndArcShapeV2 {
	if _, ok := stage.TopEndArcShapeV2s[topendarcshapev2]; !ok {
		stage.TopEndArcShapeV2s[topendarcshapev2] = struct{}{}
		stage.TopEndArcShapeV2_stagedOrder[topendarcshapev2] = stage.TopEndArcShapeV2Order
		stage.TopEndArcShapeV2_orderStaged[stage.TopEndArcShapeV2Order] = topendarcshapev2
		stage.TopEndArcShapeV2Order++
	}
	stage.TopEndArcShapeV2s_mapString[topendarcshapev2.Name] = topendarcshapev2

	return topendarcshapev2
}

// StagePreserveOrder puts topendarcshapev2 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TopEndArcShapeV2Order
// - update stage.TopEndArcShapeV2Order accordingly
func (topendarcshapev2 *TopEndArcShapeV2) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TopEndArcShapeV2s[topendarcshapev2]; !ok {
		stage.TopEndArcShapeV2s[topendarcshapev2] = struct{}{}

		if order > stage.TopEndArcShapeV2Order {
			stage.TopEndArcShapeV2Order = order
		}
		stage.TopEndArcShapeV2_stagedOrder[topendarcshapev2] = order
		stage.TopEndArcShapeV2_orderStaged[order] = topendarcshapev2
		stage.TopEndArcShapeV2Order++
	}
	stage.TopEndArcShapeV2s_mapString[topendarcshapev2.Name] = topendarcshapev2
}

// Unstage removes topendarcshapev2 off the model stage
func (topendarcshapev2 *TopEndArcShapeV2) Unstage(stage *Stage) *TopEndArcShapeV2 {
	delete(stage.TopEndArcShapeV2s, topendarcshapev2)
	// issue1150
	// delete(stage.TopEndArcShapeV2_stagedOrder, topendarcshapev2)
	delete(stage.TopEndArcShapeV2s_mapString, topendarcshapev2.Name)

	return topendarcshapev2
}

// UnstageVoid removes topendarcshapev2 off the model stage
func (topendarcshapev2 *TopEndArcShapeV2) UnstageVoid(stage *Stage) {
	delete(stage.TopEndArcShapeV2s, topendarcshapev2)
	// issue1150
	// delete(stage.TopEndArcShapeV2_stagedOrder, topendarcshapev2)
	delete(stage.TopEndArcShapeV2s_mapString, topendarcshapev2.Name)
}

// commit topendarcshapev2 to the back repo (if it is already staged)
func (topendarcshapev2 *TopEndArcShapeV2) Commit(stage *Stage) *TopEndArcShapeV2 {
	if _, ok := stage.TopEndArcShapeV2s[topendarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTopEndArcShapeV2(topendarcshapev2)
		}
	}
	return topendarcshapev2
}

func (topendarcshapev2 *TopEndArcShapeV2) CommitVoid(stage *Stage) {
	topendarcshapev2.Commit(stage)
}

func (topendarcshapev2 *TopEndArcShapeV2) StageVoid(stage *Stage) {
	topendarcshapev2.Stage(stage)
}

// Checkout topendarcshapev2 to the back repo (if it is already staged)
func (topendarcshapev2 *TopEndArcShapeV2) Checkout(stage *Stage) *TopEndArcShapeV2 {
	if _, ok := stage.TopEndArcShapeV2s[topendarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTopEndArcShapeV2(topendarcshapev2)
		}
	}
	return topendarcshapev2
}

// for satisfaction of GongStruct interface
func (topendarcshapev2 *TopEndArcShapeV2) GetName() (res string) {
	return topendarcshapev2.Name
}

// for satisfaction of GongStruct interface
func (topendarcshapev2 *TopEndArcShapeV2) SetName(name string) {
	topendarcshapev2.Name = name
}

// Stage puts topendarcshapev2grid to the model stage
func (topendarcshapev2grid *TopEndArcShapeV2Grid) Stage(stage *Stage) *TopEndArcShapeV2Grid {
	if _, ok := stage.TopEndArcShapeV2Grids[topendarcshapev2grid]; !ok {
		stage.TopEndArcShapeV2Grids[topendarcshapev2grid] = struct{}{}
		stage.TopEndArcShapeV2Grid_stagedOrder[topendarcshapev2grid] = stage.TopEndArcShapeV2GridOrder
		stage.TopEndArcShapeV2Grid_orderStaged[stage.TopEndArcShapeV2GridOrder] = topendarcshapev2grid
		stage.TopEndArcShapeV2GridOrder++
	}
	stage.TopEndArcShapeV2Grids_mapString[topendarcshapev2grid.Name] = topendarcshapev2grid

	return topendarcshapev2grid
}

// StagePreserveOrder puts topendarcshapev2grid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TopEndArcShapeV2GridOrder
// - update stage.TopEndArcShapeV2GridOrder accordingly
func (topendarcshapev2grid *TopEndArcShapeV2Grid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TopEndArcShapeV2Grids[topendarcshapev2grid]; !ok {
		stage.TopEndArcShapeV2Grids[topendarcshapev2grid] = struct{}{}

		if order > stage.TopEndArcShapeV2GridOrder {
			stage.TopEndArcShapeV2GridOrder = order
		}
		stage.TopEndArcShapeV2Grid_stagedOrder[topendarcshapev2grid] = order
		stage.TopEndArcShapeV2Grid_orderStaged[order] = topendarcshapev2grid
		stage.TopEndArcShapeV2GridOrder++
	}
	stage.TopEndArcShapeV2Grids_mapString[topendarcshapev2grid.Name] = topendarcshapev2grid
}

// Unstage removes topendarcshapev2grid off the model stage
func (topendarcshapev2grid *TopEndArcShapeV2Grid) Unstage(stage *Stage) *TopEndArcShapeV2Grid {
	delete(stage.TopEndArcShapeV2Grids, topendarcshapev2grid)
	// issue1150
	// delete(stage.TopEndArcShapeV2Grid_stagedOrder, topendarcshapev2grid)
	delete(stage.TopEndArcShapeV2Grids_mapString, topendarcshapev2grid.Name)

	return topendarcshapev2grid
}

// UnstageVoid removes topendarcshapev2grid off the model stage
func (topendarcshapev2grid *TopEndArcShapeV2Grid) UnstageVoid(stage *Stage) {
	delete(stage.TopEndArcShapeV2Grids, topendarcshapev2grid)
	// issue1150
	// delete(stage.TopEndArcShapeV2Grid_stagedOrder, topendarcshapev2grid)
	delete(stage.TopEndArcShapeV2Grids_mapString, topendarcshapev2grid.Name)
}

// commit topendarcshapev2grid to the back repo (if it is already staged)
func (topendarcshapev2grid *TopEndArcShapeV2Grid) Commit(stage *Stage) *TopEndArcShapeV2Grid {
	if _, ok := stage.TopEndArcShapeV2Grids[topendarcshapev2grid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTopEndArcShapeV2Grid(topendarcshapev2grid)
		}
	}
	return topendarcshapev2grid
}

func (topendarcshapev2grid *TopEndArcShapeV2Grid) CommitVoid(stage *Stage) {
	topendarcshapev2grid.Commit(stage)
}

func (topendarcshapev2grid *TopEndArcShapeV2Grid) StageVoid(stage *Stage) {
	topendarcshapev2grid.Stage(stage)
}

// Checkout topendarcshapev2grid to the back repo (if it is already staged)
func (topendarcshapev2grid *TopEndArcShapeV2Grid) Checkout(stage *Stage) *TopEndArcShapeV2Grid {
	if _, ok := stage.TopEndArcShapeV2Grids[topendarcshapev2grid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTopEndArcShapeV2Grid(topendarcshapev2grid)
		}
	}
	return topendarcshapev2grid
}

// for satisfaction of GongStruct interface
func (topendarcshapev2grid *TopEndArcShapeV2Grid) GetName() (res string) {
	return topendarcshapev2grid.Name
}

// for satisfaction of GongStruct interface
func (topendarcshapev2grid *TopEndArcShapeV2Grid) SetName(name string) {
	topendarcshapev2grid.Name = name
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

// Stage puts topstackgrowthcurveendarcshapev2 to the model stage
func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) Stage(stage *Stage) *TopStackGrowthCurveEndArcShapeV2 {
	if _, ok := stage.TopStackGrowthCurveEndArcShapeV2s[topstackgrowthcurveendarcshapev2]; !ok {
		stage.TopStackGrowthCurveEndArcShapeV2s[topstackgrowthcurveendarcshapev2] = struct{}{}
		stage.TopStackGrowthCurveEndArcShapeV2_stagedOrder[topstackgrowthcurveendarcshapev2] = stage.TopStackGrowthCurveEndArcShapeV2Order
		stage.TopStackGrowthCurveEndArcShapeV2_orderStaged[stage.TopStackGrowthCurveEndArcShapeV2Order] = topstackgrowthcurveendarcshapev2
		stage.TopStackGrowthCurveEndArcShapeV2Order++
	}
	stage.TopStackGrowthCurveEndArcShapeV2s_mapString[topstackgrowthcurveendarcshapev2.Name] = topstackgrowthcurveendarcshapev2

	return topstackgrowthcurveendarcshapev2
}

// StagePreserveOrder puts topstackgrowthcurveendarcshapev2 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TopStackGrowthCurveEndArcShapeV2Order
// - update stage.TopStackGrowthCurveEndArcShapeV2Order accordingly
func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TopStackGrowthCurveEndArcShapeV2s[topstackgrowthcurveendarcshapev2]; !ok {
		stage.TopStackGrowthCurveEndArcShapeV2s[topstackgrowthcurveendarcshapev2] = struct{}{}

		if order > stage.TopStackGrowthCurveEndArcShapeV2Order {
			stage.TopStackGrowthCurveEndArcShapeV2Order = order
		}
		stage.TopStackGrowthCurveEndArcShapeV2_stagedOrder[topstackgrowthcurveendarcshapev2] = order
		stage.TopStackGrowthCurveEndArcShapeV2_orderStaged[order] = topstackgrowthcurveendarcshapev2
		stage.TopStackGrowthCurveEndArcShapeV2Order++
	}
	stage.TopStackGrowthCurveEndArcShapeV2s_mapString[topstackgrowthcurveendarcshapev2.Name] = topstackgrowthcurveendarcshapev2
}

// Unstage removes topstackgrowthcurveendarcshapev2 off the model stage
func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) Unstage(stage *Stage) *TopStackGrowthCurveEndArcShapeV2 {
	delete(stage.TopStackGrowthCurveEndArcShapeV2s, topstackgrowthcurveendarcshapev2)
	// issue1150
	// delete(stage.TopStackGrowthCurveEndArcShapeV2_stagedOrder, topstackgrowthcurveendarcshapev2)
	delete(stage.TopStackGrowthCurveEndArcShapeV2s_mapString, topstackgrowthcurveendarcshapev2.Name)

	return topstackgrowthcurveendarcshapev2
}

// UnstageVoid removes topstackgrowthcurveendarcshapev2 off the model stage
func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) UnstageVoid(stage *Stage) {
	delete(stage.TopStackGrowthCurveEndArcShapeV2s, topstackgrowthcurveendarcshapev2)
	// issue1150
	// delete(stage.TopStackGrowthCurveEndArcShapeV2_stagedOrder, topstackgrowthcurveendarcshapev2)
	delete(stage.TopStackGrowthCurveEndArcShapeV2s_mapString, topstackgrowthcurveendarcshapev2.Name)
}

// commit topstackgrowthcurveendarcshapev2 to the back repo (if it is already staged)
func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) Commit(stage *Stage) *TopStackGrowthCurveEndArcShapeV2 {
	if _, ok := stage.TopStackGrowthCurveEndArcShapeV2s[topstackgrowthcurveendarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTopStackGrowthCurveEndArcShapeV2(topstackgrowthcurveendarcshapev2)
		}
	}
	return topstackgrowthcurveendarcshapev2
}

func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) CommitVoid(stage *Stage) {
	topstackgrowthcurveendarcshapev2.Commit(stage)
}

func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) StageVoid(stage *Stage) {
	topstackgrowthcurveendarcshapev2.Stage(stage)
}

// Checkout topstackgrowthcurveendarcshapev2 to the back repo (if it is already staged)
func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) Checkout(stage *Stage) *TopStackGrowthCurveEndArcShapeV2 {
	if _, ok := stage.TopStackGrowthCurveEndArcShapeV2s[topstackgrowthcurveendarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTopStackGrowthCurveEndArcShapeV2(topstackgrowthcurveendarcshapev2)
		}
	}
	return topstackgrowthcurveendarcshapev2
}

// for satisfaction of GongStruct interface
func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) GetName() (res string) {
	return topstackgrowthcurveendarcshapev2.Name
}

// for satisfaction of GongStruct interface
func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) SetName(name string) {
	topstackgrowthcurveendarcshapev2.Name = name
}

// Stage puts topstackgrowthcurvestartarcshapev2 to the model stage
func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) Stage(stage *Stage) *TopStackGrowthCurveStartArcShapeV2 {
	if _, ok := stage.TopStackGrowthCurveStartArcShapeV2s[topstackgrowthcurvestartarcshapev2]; !ok {
		stage.TopStackGrowthCurveStartArcShapeV2s[topstackgrowthcurvestartarcshapev2] = struct{}{}
		stage.TopStackGrowthCurveStartArcShapeV2_stagedOrder[topstackgrowthcurvestartarcshapev2] = stage.TopStackGrowthCurveStartArcShapeV2Order
		stage.TopStackGrowthCurveStartArcShapeV2_orderStaged[stage.TopStackGrowthCurveStartArcShapeV2Order] = topstackgrowthcurvestartarcshapev2
		stage.TopStackGrowthCurveStartArcShapeV2Order++
	}
	stage.TopStackGrowthCurveStartArcShapeV2s_mapString[topstackgrowthcurvestartarcshapev2.Name] = topstackgrowthcurvestartarcshapev2

	return topstackgrowthcurvestartarcshapev2
}

// StagePreserveOrder puts topstackgrowthcurvestartarcshapev2 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TopStackGrowthCurveStartArcShapeV2Order
// - update stage.TopStackGrowthCurveStartArcShapeV2Order accordingly
func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TopStackGrowthCurveStartArcShapeV2s[topstackgrowthcurvestartarcshapev2]; !ok {
		stage.TopStackGrowthCurveStartArcShapeV2s[topstackgrowthcurvestartarcshapev2] = struct{}{}

		if order > stage.TopStackGrowthCurveStartArcShapeV2Order {
			stage.TopStackGrowthCurveStartArcShapeV2Order = order
		}
		stage.TopStackGrowthCurveStartArcShapeV2_stagedOrder[topstackgrowthcurvestartarcshapev2] = order
		stage.TopStackGrowthCurveStartArcShapeV2_orderStaged[order] = topstackgrowthcurvestartarcshapev2
		stage.TopStackGrowthCurveStartArcShapeV2Order++
	}
	stage.TopStackGrowthCurveStartArcShapeV2s_mapString[topstackgrowthcurvestartarcshapev2.Name] = topstackgrowthcurvestartarcshapev2
}

// Unstage removes topstackgrowthcurvestartarcshapev2 off the model stage
func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) Unstage(stage *Stage) *TopStackGrowthCurveStartArcShapeV2 {
	delete(stage.TopStackGrowthCurveStartArcShapeV2s, topstackgrowthcurvestartarcshapev2)
	// issue1150
	// delete(stage.TopStackGrowthCurveStartArcShapeV2_stagedOrder, topstackgrowthcurvestartarcshapev2)
	delete(stage.TopStackGrowthCurveStartArcShapeV2s_mapString, topstackgrowthcurvestartarcshapev2.Name)

	return topstackgrowthcurvestartarcshapev2
}

// UnstageVoid removes topstackgrowthcurvestartarcshapev2 off the model stage
func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) UnstageVoid(stage *Stage) {
	delete(stage.TopStackGrowthCurveStartArcShapeV2s, topstackgrowthcurvestartarcshapev2)
	// issue1150
	// delete(stage.TopStackGrowthCurveStartArcShapeV2_stagedOrder, topstackgrowthcurvestartarcshapev2)
	delete(stage.TopStackGrowthCurveStartArcShapeV2s_mapString, topstackgrowthcurvestartarcshapev2.Name)
}

// commit topstackgrowthcurvestartarcshapev2 to the back repo (if it is already staged)
func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) Commit(stage *Stage) *TopStackGrowthCurveStartArcShapeV2 {
	if _, ok := stage.TopStackGrowthCurveStartArcShapeV2s[topstackgrowthcurvestartarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTopStackGrowthCurveStartArcShapeV2(topstackgrowthcurvestartarcshapev2)
		}
	}
	return topstackgrowthcurvestartarcshapev2
}

func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) CommitVoid(stage *Stage) {
	topstackgrowthcurvestartarcshapev2.Commit(stage)
}

func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) StageVoid(stage *Stage) {
	topstackgrowthcurvestartarcshapev2.Stage(stage)
}

// Checkout topstackgrowthcurvestartarcshapev2 to the back repo (if it is already staged)
func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) Checkout(stage *Stage) *TopStackGrowthCurveStartArcShapeV2 {
	if _, ok := stage.TopStackGrowthCurveStartArcShapeV2s[topstackgrowthcurvestartarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTopStackGrowthCurveStartArcShapeV2(topstackgrowthcurvestartarcshapev2)
		}
	}
	return topstackgrowthcurvestartarcshapev2
}

// for satisfaction of GongStruct interface
func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) GetName() (res string) {
	return topstackgrowthcurvestartarcshapev2.Name
}

// for satisfaction of GongStruct interface
func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) SetName(name string) {
	topstackgrowthcurvestartarcshapev2.Name = name
}

// Stage puts topstackofgrowthcurvev2 to the model stage
func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) Stage(stage *Stage) *TopStackOfGrowthCurveV2 {
	if _, ok := stage.TopStackOfGrowthCurveV2s[topstackofgrowthcurvev2]; !ok {
		stage.TopStackOfGrowthCurveV2s[topstackofgrowthcurvev2] = struct{}{}
		stage.TopStackOfGrowthCurveV2_stagedOrder[topstackofgrowthcurvev2] = stage.TopStackOfGrowthCurveV2Order
		stage.TopStackOfGrowthCurveV2_orderStaged[stage.TopStackOfGrowthCurveV2Order] = topstackofgrowthcurvev2
		stage.TopStackOfGrowthCurveV2Order++
	}
	stage.TopStackOfGrowthCurveV2s_mapString[topstackofgrowthcurvev2.Name] = topstackofgrowthcurvev2

	return topstackofgrowthcurvev2
}

// StagePreserveOrder puts topstackofgrowthcurvev2 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TopStackOfGrowthCurveV2Order
// - update stage.TopStackOfGrowthCurveV2Order accordingly
func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TopStackOfGrowthCurveV2s[topstackofgrowthcurvev2]; !ok {
		stage.TopStackOfGrowthCurveV2s[topstackofgrowthcurvev2] = struct{}{}

		if order > stage.TopStackOfGrowthCurveV2Order {
			stage.TopStackOfGrowthCurveV2Order = order
		}
		stage.TopStackOfGrowthCurveV2_stagedOrder[topstackofgrowthcurvev2] = order
		stage.TopStackOfGrowthCurveV2_orderStaged[order] = topstackofgrowthcurvev2
		stage.TopStackOfGrowthCurveV2Order++
	}
	stage.TopStackOfGrowthCurveV2s_mapString[topstackofgrowthcurvev2.Name] = topstackofgrowthcurvev2
}

// Unstage removes topstackofgrowthcurvev2 off the model stage
func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) Unstage(stage *Stage) *TopStackOfGrowthCurveV2 {
	delete(stage.TopStackOfGrowthCurveV2s, topstackofgrowthcurvev2)
	// issue1150
	// delete(stage.TopStackOfGrowthCurveV2_stagedOrder, topstackofgrowthcurvev2)
	delete(stage.TopStackOfGrowthCurveV2s_mapString, topstackofgrowthcurvev2.Name)

	return topstackofgrowthcurvev2
}

// UnstageVoid removes topstackofgrowthcurvev2 off the model stage
func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) UnstageVoid(stage *Stage) {
	delete(stage.TopStackOfGrowthCurveV2s, topstackofgrowthcurvev2)
	// issue1150
	// delete(stage.TopStackOfGrowthCurveV2_stagedOrder, topstackofgrowthcurvev2)
	delete(stage.TopStackOfGrowthCurveV2s_mapString, topstackofgrowthcurvev2.Name)
}

// commit topstackofgrowthcurvev2 to the back repo (if it is already staged)
func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) Commit(stage *Stage) *TopStackOfGrowthCurveV2 {
	if _, ok := stage.TopStackOfGrowthCurveV2s[topstackofgrowthcurvev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTopStackOfGrowthCurveV2(topstackofgrowthcurvev2)
		}
	}
	return topstackofgrowthcurvev2
}

func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) CommitVoid(stage *Stage) {
	topstackofgrowthcurvev2.Commit(stage)
}

func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) StageVoid(stage *Stage) {
	topstackofgrowthcurvev2.Stage(stage)
}

// Checkout topstackofgrowthcurvev2 to the back repo (if it is already staged)
func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) Checkout(stage *Stage) *TopStackOfGrowthCurveV2 {
	if _, ok := stage.TopStackOfGrowthCurveV2s[topstackofgrowthcurvev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTopStackOfGrowthCurveV2(topstackofgrowthcurvev2)
		}
	}
	return topstackofgrowthcurvev2
}

// for satisfaction of GongStruct interface
func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) GetName() (res string) {
	return topstackofgrowthcurvev2.Name
}

// for satisfaction of GongStruct interface
func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) SetName(name string) {
	topstackofgrowthcurvev2.Name = name
}

// Stage puts topstartarcshapev2 to the model stage
func (topstartarcshapev2 *TopStartArcShapeV2) Stage(stage *Stage) *TopStartArcShapeV2 {
	if _, ok := stage.TopStartArcShapeV2s[topstartarcshapev2]; !ok {
		stage.TopStartArcShapeV2s[topstartarcshapev2] = struct{}{}
		stage.TopStartArcShapeV2_stagedOrder[topstartarcshapev2] = stage.TopStartArcShapeV2Order
		stage.TopStartArcShapeV2_orderStaged[stage.TopStartArcShapeV2Order] = topstartarcshapev2
		stage.TopStartArcShapeV2Order++
	}
	stage.TopStartArcShapeV2s_mapString[topstartarcshapev2.Name] = topstartarcshapev2

	return topstartarcshapev2
}

// StagePreserveOrder puts topstartarcshapev2 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TopStartArcShapeV2Order
// - update stage.TopStartArcShapeV2Order accordingly
func (topstartarcshapev2 *TopStartArcShapeV2) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TopStartArcShapeV2s[topstartarcshapev2]; !ok {
		stage.TopStartArcShapeV2s[topstartarcshapev2] = struct{}{}

		if order > stage.TopStartArcShapeV2Order {
			stage.TopStartArcShapeV2Order = order
		}
		stage.TopStartArcShapeV2_stagedOrder[topstartarcshapev2] = order
		stage.TopStartArcShapeV2_orderStaged[order] = topstartarcshapev2
		stage.TopStartArcShapeV2Order++
	}
	stage.TopStartArcShapeV2s_mapString[topstartarcshapev2.Name] = topstartarcshapev2
}

// Unstage removes topstartarcshapev2 off the model stage
func (topstartarcshapev2 *TopStartArcShapeV2) Unstage(stage *Stage) *TopStartArcShapeV2 {
	delete(stage.TopStartArcShapeV2s, topstartarcshapev2)
	// issue1150
	// delete(stage.TopStartArcShapeV2_stagedOrder, topstartarcshapev2)
	delete(stage.TopStartArcShapeV2s_mapString, topstartarcshapev2.Name)

	return topstartarcshapev2
}

// UnstageVoid removes topstartarcshapev2 off the model stage
func (topstartarcshapev2 *TopStartArcShapeV2) UnstageVoid(stage *Stage) {
	delete(stage.TopStartArcShapeV2s, topstartarcshapev2)
	// issue1150
	// delete(stage.TopStartArcShapeV2_stagedOrder, topstartarcshapev2)
	delete(stage.TopStartArcShapeV2s_mapString, topstartarcshapev2.Name)
}

// commit topstartarcshapev2 to the back repo (if it is already staged)
func (topstartarcshapev2 *TopStartArcShapeV2) Commit(stage *Stage) *TopStartArcShapeV2 {
	if _, ok := stage.TopStartArcShapeV2s[topstartarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTopStartArcShapeV2(topstartarcshapev2)
		}
	}
	return topstartarcshapev2
}

func (topstartarcshapev2 *TopStartArcShapeV2) CommitVoid(stage *Stage) {
	topstartarcshapev2.Commit(stage)
}

func (topstartarcshapev2 *TopStartArcShapeV2) StageVoid(stage *Stage) {
	topstartarcshapev2.Stage(stage)
}

// Checkout topstartarcshapev2 to the back repo (if it is already staged)
func (topstartarcshapev2 *TopStartArcShapeV2) Checkout(stage *Stage) *TopStartArcShapeV2 {
	if _, ok := stage.TopStartArcShapeV2s[topstartarcshapev2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTopStartArcShapeV2(topstartarcshapev2)
		}
	}
	return topstartarcshapev2
}

// for satisfaction of GongStruct interface
func (topstartarcshapev2 *TopStartArcShapeV2) GetName() (res string) {
	return topstartarcshapev2.Name
}

// for satisfaction of GongStruct interface
func (topstartarcshapev2 *TopStartArcShapeV2) SetName(name string) {
	topstartarcshapev2.Name = name
}

// Stage puts topstartarcshapev2grid to the model stage
func (topstartarcshapev2grid *TopStartArcShapeV2Grid) Stage(stage *Stage) *TopStartArcShapeV2Grid {
	if _, ok := stage.TopStartArcShapeV2Grids[topstartarcshapev2grid]; !ok {
		stage.TopStartArcShapeV2Grids[topstartarcshapev2grid] = struct{}{}
		stage.TopStartArcShapeV2Grid_stagedOrder[topstartarcshapev2grid] = stage.TopStartArcShapeV2GridOrder
		stage.TopStartArcShapeV2Grid_orderStaged[stage.TopStartArcShapeV2GridOrder] = topstartarcshapev2grid
		stage.TopStartArcShapeV2GridOrder++
	}
	stage.TopStartArcShapeV2Grids_mapString[topstartarcshapev2grid.Name] = topstartarcshapev2grid

	return topstartarcshapev2grid
}

// StagePreserveOrder puts topstartarcshapev2grid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TopStartArcShapeV2GridOrder
// - update stage.TopStartArcShapeV2GridOrder accordingly
func (topstartarcshapev2grid *TopStartArcShapeV2Grid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TopStartArcShapeV2Grids[topstartarcshapev2grid]; !ok {
		stage.TopStartArcShapeV2Grids[topstartarcshapev2grid] = struct{}{}

		if order > stage.TopStartArcShapeV2GridOrder {
			stage.TopStartArcShapeV2GridOrder = order
		}
		stage.TopStartArcShapeV2Grid_stagedOrder[topstartarcshapev2grid] = order
		stage.TopStartArcShapeV2Grid_orderStaged[order] = topstartarcshapev2grid
		stage.TopStartArcShapeV2GridOrder++
	}
	stage.TopStartArcShapeV2Grids_mapString[topstartarcshapev2grid.Name] = topstartarcshapev2grid
}

// Unstage removes topstartarcshapev2grid off the model stage
func (topstartarcshapev2grid *TopStartArcShapeV2Grid) Unstage(stage *Stage) *TopStartArcShapeV2Grid {
	delete(stage.TopStartArcShapeV2Grids, topstartarcshapev2grid)
	// issue1150
	// delete(stage.TopStartArcShapeV2Grid_stagedOrder, topstartarcshapev2grid)
	delete(stage.TopStartArcShapeV2Grids_mapString, topstartarcshapev2grid.Name)

	return topstartarcshapev2grid
}

// UnstageVoid removes topstartarcshapev2grid off the model stage
func (topstartarcshapev2grid *TopStartArcShapeV2Grid) UnstageVoid(stage *Stage) {
	delete(stage.TopStartArcShapeV2Grids, topstartarcshapev2grid)
	// issue1150
	// delete(stage.TopStartArcShapeV2Grid_stagedOrder, topstartarcshapev2grid)
	delete(stage.TopStartArcShapeV2Grids_mapString, topstartarcshapev2grid.Name)
}

// commit topstartarcshapev2grid to the back repo (if it is already staged)
func (topstartarcshapev2grid *TopStartArcShapeV2Grid) Commit(stage *Stage) *TopStartArcShapeV2Grid {
	if _, ok := stage.TopStartArcShapeV2Grids[topstartarcshapev2grid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTopStartArcShapeV2Grid(topstartarcshapev2grid)
		}
	}
	return topstartarcshapev2grid
}

func (topstartarcshapev2grid *TopStartArcShapeV2Grid) CommitVoid(stage *Stage) {
	topstartarcshapev2grid.Commit(stage)
}

func (topstartarcshapev2grid *TopStartArcShapeV2Grid) StageVoid(stage *Stage) {
	topstartarcshapev2grid.Stage(stage)
}

// Checkout topstartarcshapev2grid to the back repo (if it is already staged)
func (topstartarcshapev2grid *TopStartArcShapeV2Grid) Checkout(stage *Stage) *TopStartArcShapeV2Grid {
	if _, ok := stage.TopStartArcShapeV2Grids[topstartarcshapev2grid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTopStartArcShapeV2Grid(topstartarcshapev2grid)
		}
	}
	return topstartarcshapev2grid
}

// for satisfaction of GongStruct interface
func (topstartarcshapev2grid *TopStartArcShapeV2Grid) GetName() (res string) {
	return topstartarcshapev2grid.Name
}

// for satisfaction of GongStruct interface
func (topstartarcshapev2grid *TopStartArcShapeV2Grid) SetName(name string) {
	topstartarcshapev2grid.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMArcNormalVectorShape(ArcNormalVectorShape *ArcNormalVectorShape)
	CreateORMArcNormalVectorShapeGrid(ArcNormalVectorShapeGrid *ArcNormalVectorShapeGrid)
	CreateORMAxesShape(AxesShape *AxesShape)
	CreateORMBaseVectorShape(BaseVectorShape *BaseVectorShape)
	CreateORMBaseVectorShapeGrid(BaseVectorShapeGrid *BaseVectorShapeGrid)
	CreateORMBottomEndArcShapeV2(BottomEndArcShapeV2 *BottomEndArcShapeV2)
	CreateORMBottomEndArcShapeV2Grid(BottomEndArcShapeV2Grid *BottomEndArcShapeV2Grid)
	CreateORMBottomStackGrowthCurveEndArcShapeV2(BottomStackGrowthCurveEndArcShapeV2 *BottomStackGrowthCurveEndArcShapeV2)
	CreateORMBottomStackGrowthCurveStartArcShapeV2(BottomStackGrowthCurveStartArcShapeV2 *BottomStackGrowthCurveStartArcShapeV2)
	CreateORMBottomStackOfGrowthCurveV2(BottomStackOfGrowthCurveV2 *BottomStackOfGrowthCurveV2)
	CreateORMBottomStartArcShapeV2(BottomStartArcShapeV2 *BottomStartArcShapeV2)
	CreateORMBottomStartArcShapeV2Grid(BottomStartArcShapeV2Grid *BottomStartArcShapeV2Grid)
	CreateORMCircleGridShape(CircleGridShape *CircleGridShape)
	CreateORMEndArcShapeV2(EndArcShapeV2 *EndArcShapeV2)
	CreateORMEndArcShapeV2Grid(EndArcShapeV2Grid *EndArcShapeV2Grid)
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
	CreateORMStackGrowthCurveBezierShape(StackGrowthCurveBezierShape *StackGrowthCurveBezierShape)
	CreateORMStackGrowthCurveEndArcShapeV2(StackGrowthCurveEndArcShapeV2 *StackGrowthCurveEndArcShapeV2)
	CreateORMStackGrowthCurveStartArcShapeV2(StackGrowthCurveStartArcShapeV2 *StackGrowthCurveStartArcShapeV2)
	CreateORMStackOfGrowthCurve(StackOfGrowthCurve *StackOfGrowthCurve)
	CreateORMStackOfGrowthCurveV2(StackOfGrowthCurveV2 *StackOfGrowthCurveV2)
	CreateORMStartArcShapeV2(StartArcShapeV2 *StartArcShapeV2)
	CreateORMStartArcShapeV2Grid(StartArcShapeV2Grid *StartArcShapeV2Grid)
	CreateORMTopEndArcShapeV2(TopEndArcShapeV2 *TopEndArcShapeV2)
	CreateORMTopEndArcShapeV2Grid(TopEndArcShapeV2Grid *TopEndArcShapeV2Grid)
	CreateORMTopGrowthCurve2D(TopGrowthCurve2D *TopGrowthCurve2D)
	CreateORMTopStackGrowthCurveEndArcShapeV2(TopStackGrowthCurveEndArcShapeV2 *TopStackGrowthCurveEndArcShapeV2)
	CreateORMTopStackGrowthCurveStartArcShapeV2(TopStackGrowthCurveStartArcShapeV2 *TopStackGrowthCurveStartArcShapeV2)
	CreateORMTopStackOfGrowthCurveV2(TopStackOfGrowthCurveV2 *TopStackOfGrowthCurveV2)
	CreateORMTopStartArcShapeV2(TopStartArcShapeV2 *TopStartArcShapeV2)
	CreateORMTopStartArcShapeV2Grid(TopStartArcShapeV2Grid *TopStartArcShapeV2Grid)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMArcNormalVectorShape(ArcNormalVectorShape *ArcNormalVectorShape)
	DeleteORMArcNormalVectorShapeGrid(ArcNormalVectorShapeGrid *ArcNormalVectorShapeGrid)
	DeleteORMAxesShape(AxesShape *AxesShape)
	DeleteORMBaseVectorShape(BaseVectorShape *BaseVectorShape)
	DeleteORMBaseVectorShapeGrid(BaseVectorShapeGrid *BaseVectorShapeGrid)
	DeleteORMBottomEndArcShapeV2(BottomEndArcShapeV2 *BottomEndArcShapeV2)
	DeleteORMBottomEndArcShapeV2Grid(BottomEndArcShapeV2Grid *BottomEndArcShapeV2Grid)
	DeleteORMBottomStackGrowthCurveEndArcShapeV2(BottomStackGrowthCurveEndArcShapeV2 *BottomStackGrowthCurveEndArcShapeV2)
	DeleteORMBottomStackGrowthCurveStartArcShapeV2(BottomStackGrowthCurveStartArcShapeV2 *BottomStackGrowthCurveStartArcShapeV2)
	DeleteORMBottomStackOfGrowthCurveV2(BottomStackOfGrowthCurveV2 *BottomStackOfGrowthCurveV2)
	DeleteORMBottomStartArcShapeV2(BottomStartArcShapeV2 *BottomStartArcShapeV2)
	DeleteORMBottomStartArcShapeV2Grid(BottomStartArcShapeV2Grid *BottomStartArcShapeV2Grid)
	DeleteORMCircleGridShape(CircleGridShape *CircleGridShape)
	DeleteORMEndArcShapeV2(EndArcShapeV2 *EndArcShapeV2)
	DeleteORMEndArcShapeV2Grid(EndArcShapeV2Grid *EndArcShapeV2Grid)
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
	DeleteORMStackGrowthCurveBezierShape(StackGrowthCurveBezierShape *StackGrowthCurveBezierShape)
	DeleteORMStackGrowthCurveEndArcShapeV2(StackGrowthCurveEndArcShapeV2 *StackGrowthCurveEndArcShapeV2)
	DeleteORMStackGrowthCurveStartArcShapeV2(StackGrowthCurveStartArcShapeV2 *StackGrowthCurveStartArcShapeV2)
	DeleteORMStackOfGrowthCurve(StackOfGrowthCurve *StackOfGrowthCurve)
	DeleteORMStackOfGrowthCurveV2(StackOfGrowthCurveV2 *StackOfGrowthCurveV2)
	DeleteORMStartArcShapeV2(StartArcShapeV2 *StartArcShapeV2)
	DeleteORMStartArcShapeV2Grid(StartArcShapeV2Grid *StartArcShapeV2Grid)
	DeleteORMTopEndArcShapeV2(TopEndArcShapeV2 *TopEndArcShapeV2)
	DeleteORMTopEndArcShapeV2Grid(TopEndArcShapeV2Grid *TopEndArcShapeV2Grid)
	DeleteORMTopGrowthCurve2D(TopGrowthCurve2D *TopGrowthCurve2D)
	DeleteORMTopStackGrowthCurveEndArcShapeV2(TopStackGrowthCurveEndArcShapeV2 *TopStackGrowthCurveEndArcShapeV2)
	DeleteORMTopStackGrowthCurveStartArcShapeV2(TopStackGrowthCurveStartArcShapeV2 *TopStackGrowthCurveStartArcShapeV2)
	DeleteORMTopStackOfGrowthCurveV2(TopStackOfGrowthCurveV2 *TopStackOfGrowthCurveV2)
	DeleteORMTopStartArcShapeV2(TopStartArcShapeV2 *TopStartArcShapeV2)
	DeleteORMTopStartArcShapeV2Grid(TopStartArcShapeV2Grid *TopStartArcShapeV2Grid)
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

	stage.BottomEndArcShapeV2s = make(map[*BottomEndArcShapeV2]struct{})
	stage.BottomEndArcShapeV2s_mapString = make(map[string]*BottomEndArcShapeV2)
	stage.BottomEndArcShapeV2_stagedOrder = make(map[*BottomEndArcShapeV2]uint)
	stage.BottomEndArcShapeV2Order = 0

	stage.BottomEndArcShapeV2Grids = make(map[*BottomEndArcShapeV2Grid]struct{})
	stage.BottomEndArcShapeV2Grids_mapString = make(map[string]*BottomEndArcShapeV2Grid)
	stage.BottomEndArcShapeV2Grid_stagedOrder = make(map[*BottomEndArcShapeV2Grid]uint)
	stage.BottomEndArcShapeV2GridOrder = 0

	stage.BottomStackGrowthCurveEndArcShapeV2s = make(map[*BottomStackGrowthCurveEndArcShapeV2]struct{})
	stage.BottomStackGrowthCurveEndArcShapeV2s_mapString = make(map[string]*BottomStackGrowthCurveEndArcShapeV2)
	stage.BottomStackGrowthCurveEndArcShapeV2_stagedOrder = make(map[*BottomStackGrowthCurveEndArcShapeV2]uint)
	stage.BottomStackGrowthCurveEndArcShapeV2Order = 0

	stage.BottomStackGrowthCurveStartArcShapeV2s = make(map[*BottomStackGrowthCurveStartArcShapeV2]struct{})
	stage.BottomStackGrowthCurveStartArcShapeV2s_mapString = make(map[string]*BottomStackGrowthCurveStartArcShapeV2)
	stage.BottomStackGrowthCurveStartArcShapeV2_stagedOrder = make(map[*BottomStackGrowthCurveStartArcShapeV2]uint)
	stage.BottomStackGrowthCurveStartArcShapeV2Order = 0

	stage.BottomStackOfGrowthCurveV2s = make(map[*BottomStackOfGrowthCurveV2]struct{})
	stage.BottomStackOfGrowthCurveV2s_mapString = make(map[string]*BottomStackOfGrowthCurveV2)
	stage.BottomStackOfGrowthCurveV2_stagedOrder = make(map[*BottomStackOfGrowthCurveV2]uint)
	stage.BottomStackOfGrowthCurveV2Order = 0

	stage.BottomStartArcShapeV2s = make(map[*BottomStartArcShapeV2]struct{})
	stage.BottomStartArcShapeV2s_mapString = make(map[string]*BottomStartArcShapeV2)
	stage.BottomStartArcShapeV2_stagedOrder = make(map[*BottomStartArcShapeV2]uint)
	stage.BottomStartArcShapeV2Order = 0

	stage.BottomStartArcShapeV2Grids = make(map[*BottomStartArcShapeV2Grid]struct{})
	stage.BottomStartArcShapeV2Grids_mapString = make(map[string]*BottomStartArcShapeV2Grid)
	stage.BottomStartArcShapeV2Grid_stagedOrder = make(map[*BottomStartArcShapeV2Grid]uint)
	stage.BottomStartArcShapeV2GridOrder = 0

	stage.CircleGridShapes = make(map[*CircleGridShape]struct{})
	stage.CircleGridShapes_mapString = make(map[string]*CircleGridShape)
	stage.CircleGridShape_stagedOrder = make(map[*CircleGridShape]uint)
	stage.CircleGridShapeOrder = 0

	stage.EndArcShapeV2s = make(map[*EndArcShapeV2]struct{})
	stage.EndArcShapeV2s_mapString = make(map[string]*EndArcShapeV2)
	stage.EndArcShapeV2_stagedOrder = make(map[*EndArcShapeV2]uint)
	stage.EndArcShapeV2Order = 0

	stage.EndArcShapeV2Grids = make(map[*EndArcShapeV2Grid]struct{})
	stage.EndArcShapeV2Grids_mapString = make(map[string]*EndArcShapeV2Grid)
	stage.EndArcShapeV2Grid_stagedOrder = make(map[*EndArcShapeV2Grid]uint)
	stage.EndArcShapeV2GridOrder = 0

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

	stage.StackGrowthCurveBezierShapes = make(map[*StackGrowthCurveBezierShape]struct{})
	stage.StackGrowthCurveBezierShapes_mapString = make(map[string]*StackGrowthCurveBezierShape)
	stage.StackGrowthCurveBezierShape_stagedOrder = make(map[*StackGrowthCurveBezierShape]uint)
	stage.StackGrowthCurveBezierShapeOrder = 0

	stage.StackGrowthCurveEndArcShapeV2s = make(map[*StackGrowthCurveEndArcShapeV2]struct{})
	stage.StackGrowthCurveEndArcShapeV2s_mapString = make(map[string]*StackGrowthCurveEndArcShapeV2)
	stage.StackGrowthCurveEndArcShapeV2_stagedOrder = make(map[*StackGrowthCurveEndArcShapeV2]uint)
	stage.StackGrowthCurveEndArcShapeV2Order = 0

	stage.StackGrowthCurveStartArcShapeV2s = make(map[*StackGrowthCurveStartArcShapeV2]struct{})
	stage.StackGrowthCurveStartArcShapeV2s_mapString = make(map[string]*StackGrowthCurveStartArcShapeV2)
	stage.StackGrowthCurveStartArcShapeV2_stagedOrder = make(map[*StackGrowthCurveStartArcShapeV2]uint)
	stage.StackGrowthCurveStartArcShapeV2Order = 0

	stage.StackOfGrowthCurves = make(map[*StackOfGrowthCurve]struct{})
	stage.StackOfGrowthCurves_mapString = make(map[string]*StackOfGrowthCurve)
	stage.StackOfGrowthCurve_stagedOrder = make(map[*StackOfGrowthCurve]uint)
	stage.StackOfGrowthCurveOrder = 0

	stage.StackOfGrowthCurveV2s = make(map[*StackOfGrowthCurveV2]struct{})
	stage.StackOfGrowthCurveV2s_mapString = make(map[string]*StackOfGrowthCurveV2)
	stage.StackOfGrowthCurveV2_stagedOrder = make(map[*StackOfGrowthCurveV2]uint)
	stage.StackOfGrowthCurveV2Order = 0

	stage.StartArcShapeV2s = make(map[*StartArcShapeV2]struct{})
	stage.StartArcShapeV2s_mapString = make(map[string]*StartArcShapeV2)
	stage.StartArcShapeV2_stagedOrder = make(map[*StartArcShapeV2]uint)
	stage.StartArcShapeV2Order = 0

	stage.StartArcShapeV2Grids = make(map[*StartArcShapeV2Grid]struct{})
	stage.StartArcShapeV2Grids_mapString = make(map[string]*StartArcShapeV2Grid)
	stage.StartArcShapeV2Grid_stagedOrder = make(map[*StartArcShapeV2Grid]uint)
	stage.StartArcShapeV2GridOrder = 0

	stage.TopEndArcShapeV2s = make(map[*TopEndArcShapeV2]struct{})
	stage.TopEndArcShapeV2s_mapString = make(map[string]*TopEndArcShapeV2)
	stage.TopEndArcShapeV2_stagedOrder = make(map[*TopEndArcShapeV2]uint)
	stage.TopEndArcShapeV2Order = 0

	stage.TopEndArcShapeV2Grids = make(map[*TopEndArcShapeV2Grid]struct{})
	stage.TopEndArcShapeV2Grids_mapString = make(map[string]*TopEndArcShapeV2Grid)
	stage.TopEndArcShapeV2Grid_stagedOrder = make(map[*TopEndArcShapeV2Grid]uint)
	stage.TopEndArcShapeV2GridOrder = 0

	stage.TopGrowthCurve2Ds = make(map[*TopGrowthCurve2D]struct{})
	stage.TopGrowthCurve2Ds_mapString = make(map[string]*TopGrowthCurve2D)
	stage.TopGrowthCurve2D_stagedOrder = make(map[*TopGrowthCurve2D]uint)
	stage.TopGrowthCurve2DOrder = 0

	stage.TopStackGrowthCurveEndArcShapeV2s = make(map[*TopStackGrowthCurveEndArcShapeV2]struct{})
	stage.TopStackGrowthCurveEndArcShapeV2s_mapString = make(map[string]*TopStackGrowthCurveEndArcShapeV2)
	stage.TopStackGrowthCurveEndArcShapeV2_stagedOrder = make(map[*TopStackGrowthCurveEndArcShapeV2]uint)
	stage.TopStackGrowthCurveEndArcShapeV2Order = 0

	stage.TopStackGrowthCurveStartArcShapeV2s = make(map[*TopStackGrowthCurveStartArcShapeV2]struct{})
	stage.TopStackGrowthCurveStartArcShapeV2s_mapString = make(map[string]*TopStackGrowthCurveStartArcShapeV2)
	stage.TopStackGrowthCurveStartArcShapeV2_stagedOrder = make(map[*TopStackGrowthCurveStartArcShapeV2]uint)
	stage.TopStackGrowthCurveStartArcShapeV2Order = 0

	stage.TopStackOfGrowthCurveV2s = make(map[*TopStackOfGrowthCurveV2]struct{})
	stage.TopStackOfGrowthCurveV2s_mapString = make(map[string]*TopStackOfGrowthCurveV2)
	stage.TopStackOfGrowthCurveV2_stagedOrder = make(map[*TopStackOfGrowthCurveV2]uint)
	stage.TopStackOfGrowthCurveV2Order = 0

	stage.TopStartArcShapeV2s = make(map[*TopStartArcShapeV2]struct{})
	stage.TopStartArcShapeV2s_mapString = make(map[string]*TopStartArcShapeV2)
	stage.TopStartArcShapeV2_stagedOrder = make(map[*TopStartArcShapeV2]uint)
	stage.TopStartArcShapeV2Order = 0

	stage.TopStartArcShapeV2Grids = make(map[*TopStartArcShapeV2Grid]struct{})
	stage.TopStartArcShapeV2Grids_mapString = make(map[string]*TopStartArcShapeV2Grid)
	stage.TopStartArcShapeV2Grid_stagedOrder = make(map[*TopStartArcShapeV2Grid]uint)
	stage.TopStartArcShapeV2GridOrder = 0

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

	stage.BottomEndArcShapeV2s = nil
	stage.BottomEndArcShapeV2s_mapString = nil

	stage.BottomEndArcShapeV2Grids = nil
	stage.BottomEndArcShapeV2Grids_mapString = nil

	stage.BottomStackGrowthCurveEndArcShapeV2s = nil
	stage.BottomStackGrowthCurveEndArcShapeV2s_mapString = nil

	stage.BottomStackGrowthCurveStartArcShapeV2s = nil
	stage.BottomStackGrowthCurveStartArcShapeV2s_mapString = nil

	stage.BottomStackOfGrowthCurveV2s = nil
	stage.BottomStackOfGrowthCurveV2s_mapString = nil

	stage.BottomStartArcShapeV2s = nil
	stage.BottomStartArcShapeV2s_mapString = nil

	stage.BottomStartArcShapeV2Grids = nil
	stage.BottomStartArcShapeV2Grids_mapString = nil

	stage.CircleGridShapes = nil
	stage.CircleGridShapes_mapString = nil

	stage.EndArcShapeV2s = nil
	stage.EndArcShapeV2s_mapString = nil

	stage.EndArcShapeV2Grids = nil
	stage.EndArcShapeV2Grids_mapString = nil

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

	stage.StackGrowthCurveBezierShapes = nil
	stage.StackGrowthCurveBezierShapes_mapString = nil

	stage.StackGrowthCurveEndArcShapeV2s = nil
	stage.StackGrowthCurveEndArcShapeV2s_mapString = nil

	stage.StackGrowthCurveStartArcShapeV2s = nil
	stage.StackGrowthCurveStartArcShapeV2s_mapString = nil

	stage.StackOfGrowthCurves = nil
	stage.StackOfGrowthCurves_mapString = nil

	stage.StackOfGrowthCurveV2s = nil
	stage.StackOfGrowthCurveV2s_mapString = nil

	stage.StartArcShapeV2s = nil
	stage.StartArcShapeV2s_mapString = nil

	stage.StartArcShapeV2Grids = nil
	stage.StartArcShapeV2Grids_mapString = nil

	stage.TopEndArcShapeV2s = nil
	stage.TopEndArcShapeV2s_mapString = nil

	stage.TopEndArcShapeV2Grids = nil
	stage.TopEndArcShapeV2Grids_mapString = nil

	stage.TopGrowthCurve2Ds = nil
	stage.TopGrowthCurve2Ds_mapString = nil

	stage.TopStackGrowthCurveEndArcShapeV2s = nil
	stage.TopStackGrowthCurveEndArcShapeV2s_mapString = nil

	stage.TopStackGrowthCurveStartArcShapeV2s = nil
	stage.TopStackGrowthCurveStartArcShapeV2s_mapString = nil

	stage.TopStackOfGrowthCurveV2s = nil
	stage.TopStackOfGrowthCurveV2s_mapString = nil

	stage.TopStartArcShapeV2s = nil
	stage.TopStartArcShapeV2s_mapString = nil

	stage.TopStartArcShapeV2Grids = nil
	stage.TopStartArcShapeV2Grids_mapString = nil

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

	for bottomendarcshapev2 := range stage.BottomEndArcShapeV2s {
		bottomendarcshapev2.Unstage(stage)
	}

	for bottomendarcshapev2grid := range stage.BottomEndArcShapeV2Grids {
		bottomendarcshapev2grid.Unstage(stage)
	}

	for bottomstackgrowthcurveendarcshapev2 := range stage.BottomStackGrowthCurveEndArcShapeV2s {
		bottomstackgrowthcurveendarcshapev2.Unstage(stage)
	}

	for bottomstackgrowthcurvestartarcshapev2 := range stage.BottomStackGrowthCurveStartArcShapeV2s {
		bottomstackgrowthcurvestartarcshapev2.Unstage(stage)
	}

	for bottomstackofgrowthcurvev2 := range stage.BottomStackOfGrowthCurveV2s {
		bottomstackofgrowthcurvev2.Unstage(stage)
	}

	for bottomstartarcshapev2 := range stage.BottomStartArcShapeV2s {
		bottomstartarcshapev2.Unstage(stage)
	}

	for bottomstartarcshapev2grid := range stage.BottomStartArcShapeV2Grids {
		bottomstartarcshapev2grid.Unstage(stage)
	}

	for circlegridshape := range stage.CircleGridShapes {
		circlegridshape.Unstage(stage)
	}

	for endarcshapev2 := range stage.EndArcShapeV2s {
		endarcshapev2.Unstage(stage)
	}

	for endarcshapev2grid := range stage.EndArcShapeV2Grids {
		endarcshapev2grid.Unstage(stage)
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

	for stackgrowthcurvebeziershape := range stage.StackGrowthCurveBezierShapes {
		stackgrowthcurvebeziershape.Unstage(stage)
	}

	for stackgrowthcurveendarcshapev2 := range stage.StackGrowthCurveEndArcShapeV2s {
		stackgrowthcurveendarcshapev2.Unstage(stage)
	}

	for stackgrowthcurvestartarcshapev2 := range stage.StackGrowthCurveStartArcShapeV2s {
		stackgrowthcurvestartarcshapev2.Unstage(stage)
	}

	for stackofgrowthcurve := range stage.StackOfGrowthCurves {
		stackofgrowthcurve.Unstage(stage)
	}

	for stackofgrowthcurvev2 := range stage.StackOfGrowthCurveV2s {
		stackofgrowthcurvev2.Unstage(stage)
	}

	for startarcshapev2 := range stage.StartArcShapeV2s {
		startarcshapev2.Unstage(stage)
	}

	for startarcshapev2grid := range stage.StartArcShapeV2Grids {
		startarcshapev2grid.Unstage(stage)
	}

	for topendarcshapev2 := range stage.TopEndArcShapeV2s {
		topendarcshapev2.Unstage(stage)
	}

	for topendarcshapev2grid := range stage.TopEndArcShapeV2Grids {
		topendarcshapev2grid.Unstage(stage)
	}

	for topgrowthcurve2d := range stage.TopGrowthCurve2Ds {
		topgrowthcurve2d.Unstage(stage)
	}

	for topstackgrowthcurveendarcshapev2 := range stage.TopStackGrowthCurveEndArcShapeV2s {
		topstackgrowthcurveendarcshapev2.Unstage(stage)
	}

	for topstackgrowthcurvestartarcshapev2 := range stage.TopStackGrowthCurveStartArcShapeV2s {
		topstackgrowthcurvestartarcshapev2.Unstage(stage)
	}

	for topstackofgrowthcurvev2 := range stage.TopStackOfGrowthCurveV2s {
		topstackofgrowthcurvev2.Unstage(stage)
	}

	for topstartarcshapev2 := range stage.TopStartArcShapeV2s {
		topstartarcshapev2.Unstage(stage)
	}

	for topstartarcshapev2grid := range stage.TopStartArcShapeV2Grids {
		topstartarcshapev2grid.Unstage(stage)
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
	case map[*BottomEndArcShapeV2]any:
		return any(&stage.BottomEndArcShapeV2s).(*Type)
	case map[*BottomEndArcShapeV2Grid]any:
		return any(&stage.BottomEndArcShapeV2Grids).(*Type)
	case map[*BottomStackGrowthCurveEndArcShapeV2]any:
		return any(&stage.BottomStackGrowthCurveEndArcShapeV2s).(*Type)
	case map[*BottomStackGrowthCurveStartArcShapeV2]any:
		return any(&stage.BottomStackGrowthCurveStartArcShapeV2s).(*Type)
	case map[*BottomStackOfGrowthCurveV2]any:
		return any(&stage.BottomStackOfGrowthCurveV2s).(*Type)
	case map[*BottomStartArcShapeV2]any:
		return any(&stage.BottomStartArcShapeV2s).(*Type)
	case map[*BottomStartArcShapeV2Grid]any:
		return any(&stage.BottomStartArcShapeV2Grids).(*Type)
	case map[*CircleGridShape]any:
		return any(&stage.CircleGridShapes).(*Type)
	case map[*EndArcShapeV2]any:
		return any(&stage.EndArcShapeV2s).(*Type)
	case map[*EndArcShapeV2Grid]any:
		return any(&stage.EndArcShapeV2Grids).(*Type)
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
	case map[*StackGrowthCurveBezierShape]any:
		return any(&stage.StackGrowthCurveBezierShapes).(*Type)
	case map[*StackGrowthCurveEndArcShapeV2]any:
		return any(&stage.StackGrowthCurveEndArcShapeV2s).(*Type)
	case map[*StackGrowthCurveStartArcShapeV2]any:
		return any(&stage.StackGrowthCurveStartArcShapeV2s).(*Type)
	case map[*StackOfGrowthCurve]any:
		return any(&stage.StackOfGrowthCurves).(*Type)
	case map[*StackOfGrowthCurveV2]any:
		return any(&stage.StackOfGrowthCurveV2s).(*Type)
	case map[*StartArcShapeV2]any:
		return any(&stage.StartArcShapeV2s).(*Type)
	case map[*StartArcShapeV2Grid]any:
		return any(&stage.StartArcShapeV2Grids).(*Type)
	case map[*TopEndArcShapeV2]any:
		return any(&stage.TopEndArcShapeV2s).(*Type)
	case map[*TopEndArcShapeV2Grid]any:
		return any(&stage.TopEndArcShapeV2Grids).(*Type)
	case map[*TopGrowthCurve2D]any:
		return any(&stage.TopGrowthCurve2Ds).(*Type)
	case map[*TopStackGrowthCurveEndArcShapeV2]any:
		return any(&stage.TopStackGrowthCurveEndArcShapeV2s).(*Type)
	case map[*TopStackGrowthCurveStartArcShapeV2]any:
		return any(&stage.TopStackGrowthCurveStartArcShapeV2s).(*Type)
	case map[*TopStackOfGrowthCurveV2]any:
		return any(&stage.TopStackOfGrowthCurveV2s).(*Type)
	case map[*TopStartArcShapeV2]any:
		return any(&stage.TopStartArcShapeV2s).(*Type)
	case map[*TopStartArcShapeV2Grid]any:
		return any(&stage.TopStartArcShapeV2Grids).(*Type)
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
	case *BottomEndArcShapeV2:
		return any(stage.BottomEndArcShapeV2s_mapString).(map[string]Type)
	case *BottomEndArcShapeV2Grid:
		return any(stage.BottomEndArcShapeV2Grids_mapString).(map[string]Type)
	case *BottomStackGrowthCurveEndArcShapeV2:
		return any(stage.BottomStackGrowthCurveEndArcShapeV2s_mapString).(map[string]Type)
	case *BottomStackGrowthCurveStartArcShapeV2:
		return any(stage.BottomStackGrowthCurveStartArcShapeV2s_mapString).(map[string]Type)
	case *BottomStackOfGrowthCurveV2:
		return any(stage.BottomStackOfGrowthCurveV2s_mapString).(map[string]Type)
	case *BottomStartArcShapeV2:
		return any(stage.BottomStartArcShapeV2s_mapString).(map[string]Type)
	case *BottomStartArcShapeV2Grid:
		return any(stage.BottomStartArcShapeV2Grids_mapString).(map[string]Type)
	case *CircleGridShape:
		return any(stage.CircleGridShapes_mapString).(map[string]Type)
	case *EndArcShapeV2:
		return any(stage.EndArcShapeV2s_mapString).(map[string]Type)
	case *EndArcShapeV2Grid:
		return any(stage.EndArcShapeV2Grids_mapString).(map[string]Type)
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
	case *StackGrowthCurveBezierShape:
		return any(stage.StackGrowthCurveBezierShapes_mapString).(map[string]Type)
	case *StackGrowthCurveEndArcShapeV2:
		return any(stage.StackGrowthCurveEndArcShapeV2s_mapString).(map[string]Type)
	case *StackGrowthCurveStartArcShapeV2:
		return any(stage.StackGrowthCurveStartArcShapeV2s_mapString).(map[string]Type)
	case *StackOfGrowthCurve:
		return any(stage.StackOfGrowthCurves_mapString).(map[string]Type)
	case *StackOfGrowthCurveV2:
		return any(stage.StackOfGrowthCurveV2s_mapString).(map[string]Type)
	case *StartArcShapeV2:
		return any(stage.StartArcShapeV2s_mapString).(map[string]Type)
	case *StartArcShapeV2Grid:
		return any(stage.StartArcShapeV2Grids_mapString).(map[string]Type)
	case *TopEndArcShapeV2:
		return any(stage.TopEndArcShapeV2s_mapString).(map[string]Type)
	case *TopEndArcShapeV2Grid:
		return any(stage.TopEndArcShapeV2Grids_mapString).(map[string]Type)
	case *TopGrowthCurve2D:
		return any(stage.TopGrowthCurve2Ds_mapString).(map[string]Type)
	case *TopStackGrowthCurveEndArcShapeV2:
		return any(stage.TopStackGrowthCurveEndArcShapeV2s_mapString).(map[string]Type)
	case *TopStackGrowthCurveStartArcShapeV2:
		return any(stage.TopStackGrowthCurveStartArcShapeV2s_mapString).(map[string]Type)
	case *TopStackOfGrowthCurveV2:
		return any(stage.TopStackOfGrowthCurveV2s_mapString).(map[string]Type)
	case *TopStartArcShapeV2:
		return any(stage.TopStartArcShapeV2s_mapString).(map[string]Type)
	case *TopStartArcShapeV2Grid:
		return any(stage.TopStartArcShapeV2Grids_mapString).(map[string]Type)
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
	case BottomEndArcShapeV2:
		return any(&stage.BottomEndArcShapeV2s).(*map[*Type]struct{})
	case BottomEndArcShapeV2Grid:
		return any(&stage.BottomEndArcShapeV2Grids).(*map[*Type]struct{})
	case BottomStackGrowthCurveEndArcShapeV2:
		return any(&stage.BottomStackGrowthCurveEndArcShapeV2s).(*map[*Type]struct{})
	case BottomStackGrowthCurveStartArcShapeV2:
		return any(&stage.BottomStackGrowthCurveStartArcShapeV2s).(*map[*Type]struct{})
	case BottomStackOfGrowthCurveV2:
		return any(&stage.BottomStackOfGrowthCurveV2s).(*map[*Type]struct{})
	case BottomStartArcShapeV2:
		return any(&stage.BottomStartArcShapeV2s).(*map[*Type]struct{})
	case BottomStartArcShapeV2Grid:
		return any(&stage.BottomStartArcShapeV2Grids).(*map[*Type]struct{})
	case CircleGridShape:
		return any(&stage.CircleGridShapes).(*map[*Type]struct{})
	case EndArcShapeV2:
		return any(&stage.EndArcShapeV2s).(*map[*Type]struct{})
	case EndArcShapeV2Grid:
		return any(&stage.EndArcShapeV2Grids).(*map[*Type]struct{})
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
	case StackGrowthCurveBezierShape:
		return any(&stage.StackGrowthCurveBezierShapes).(*map[*Type]struct{})
	case StackGrowthCurveEndArcShapeV2:
		return any(&stage.StackGrowthCurveEndArcShapeV2s).(*map[*Type]struct{})
	case StackGrowthCurveStartArcShapeV2:
		return any(&stage.StackGrowthCurveStartArcShapeV2s).(*map[*Type]struct{})
	case StackOfGrowthCurve:
		return any(&stage.StackOfGrowthCurves).(*map[*Type]struct{})
	case StackOfGrowthCurveV2:
		return any(&stage.StackOfGrowthCurveV2s).(*map[*Type]struct{})
	case StartArcShapeV2:
		return any(&stage.StartArcShapeV2s).(*map[*Type]struct{})
	case StartArcShapeV2Grid:
		return any(&stage.StartArcShapeV2Grids).(*map[*Type]struct{})
	case TopEndArcShapeV2:
		return any(&stage.TopEndArcShapeV2s).(*map[*Type]struct{})
	case TopEndArcShapeV2Grid:
		return any(&stage.TopEndArcShapeV2Grids).(*map[*Type]struct{})
	case TopGrowthCurve2D:
		return any(&stage.TopGrowthCurve2Ds).(*map[*Type]struct{})
	case TopStackGrowthCurveEndArcShapeV2:
		return any(&stage.TopStackGrowthCurveEndArcShapeV2s).(*map[*Type]struct{})
	case TopStackGrowthCurveStartArcShapeV2:
		return any(&stage.TopStackGrowthCurveStartArcShapeV2s).(*map[*Type]struct{})
	case TopStackOfGrowthCurveV2:
		return any(&stage.TopStackOfGrowthCurveV2s).(*map[*Type]struct{})
	case TopStartArcShapeV2:
		return any(&stage.TopStartArcShapeV2s).(*map[*Type]struct{})
	case TopStartArcShapeV2Grid:
		return any(&stage.TopStartArcShapeV2Grids).(*map[*Type]struct{})
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
	case *BottomEndArcShapeV2:
		return any(&stage.BottomEndArcShapeV2s).(*map[Type]struct{})
	case *BottomEndArcShapeV2Grid:
		return any(&stage.BottomEndArcShapeV2Grids).(*map[Type]struct{})
	case *BottomStackGrowthCurveEndArcShapeV2:
		return any(&stage.BottomStackGrowthCurveEndArcShapeV2s).(*map[Type]struct{})
	case *BottomStackGrowthCurveStartArcShapeV2:
		return any(&stage.BottomStackGrowthCurveStartArcShapeV2s).(*map[Type]struct{})
	case *BottomStackOfGrowthCurveV2:
		return any(&stage.BottomStackOfGrowthCurveV2s).(*map[Type]struct{})
	case *BottomStartArcShapeV2:
		return any(&stage.BottomStartArcShapeV2s).(*map[Type]struct{})
	case *BottomStartArcShapeV2Grid:
		return any(&stage.BottomStartArcShapeV2Grids).(*map[Type]struct{})
	case *CircleGridShape:
		return any(&stage.CircleGridShapes).(*map[Type]struct{})
	case *EndArcShapeV2:
		return any(&stage.EndArcShapeV2s).(*map[Type]struct{})
	case *EndArcShapeV2Grid:
		return any(&stage.EndArcShapeV2Grids).(*map[Type]struct{})
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
	case *StackGrowthCurveBezierShape:
		return any(&stage.StackGrowthCurveBezierShapes).(*map[Type]struct{})
	case *StackGrowthCurveEndArcShapeV2:
		return any(&stage.StackGrowthCurveEndArcShapeV2s).(*map[Type]struct{})
	case *StackGrowthCurveStartArcShapeV2:
		return any(&stage.StackGrowthCurveStartArcShapeV2s).(*map[Type]struct{})
	case *StackOfGrowthCurve:
		return any(&stage.StackOfGrowthCurves).(*map[Type]struct{})
	case *StackOfGrowthCurveV2:
		return any(&stage.StackOfGrowthCurveV2s).(*map[Type]struct{})
	case *StartArcShapeV2:
		return any(&stage.StartArcShapeV2s).(*map[Type]struct{})
	case *StartArcShapeV2Grid:
		return any(&stage.StartArcShapeV2Grids).(*map[Type]struct{})
	case *TopEndArcShapeV2:
		return any(&stage.TopEndArcShapeV2s).(*map[Type]struct{})
	case *TopEndArcShapeV2Grid:
		return any(&stage.TopEndArcShapeV2Grids).(*map[Type]struct{})
	case *TopGrowthCurve2D:
		return any(&stage.TopGrowthCurve2Ds).(*map[Type]struct{})
	case *TopStackGrowthCurveEndArcShapeV2:
		return any(&stage.TopStackGrowthCurveEndArcShapeV2s).(*map[Type]struct{})
	case *TopStackGrowthCurveStartArcShapeV2:
		return any(&stage.TopStackGrowthCurveStartArcShapeV2s).(*map[Type]struct{})
	case *TopStackOfGrowthCurveV2:
		return any(&stage.TopStackOfGrowthCurveV2s).(*map[Type]struct{})
	case *TopStartArcShapeV2:
		return any(&stage.TopStartArcShapeV2s).(*map[Type]struct{})
	case *TopStartArcShapeV2Grid:
		return any(&stage.TopStartArcShapeV2Grids).(*map[Type]struct{})
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
	case BottomEndArcShapeV2:
		return any(&stage.BottomEndArcShapeV2s_mapString).(*map[string]*Type)
	case BottomEndArcShapeV2Grid:
		return any(&stage.BottomEndArcShapeV2Grids_mapString).(*map[string]*Type)
	case BottomStackGrowthCurveEndArcShapeV2:
		return any(&stage.BottomStackGrowthCurveEndArcShapeV2s_mapString).(*map[string]*Type)
	case BottomStackGrowthCurveStartArcShapeV2:
		return any(&stage.BottomStackGrowthCurveStartArcShapeV2s_mapString).(*map[string]*Type)
	case BottomStackOfGrowthCurveV2:
		return any(&stage.BottomStackOfGrowthCurveV2s_mapString).(*map[string]*Type)
	case BottomStartArcShapeV2:
		return any(&stage.BottomStartArcShapeV2s_mapString).(*map[string]*Type)
	case BottomStartArcShapeV2Grid:
		return any(&stage.BottomStartArcShapeV2Grids_mapString).(*map[string]*Type)
	case CircleGridShape:
		return any(&stage.CircleGridShapes_mapString).(*map[string]*Type)
	case EndArcShapeV2:
		return any(&stage.EndArcShapeV2s_mapString).(*map[string]*Type)
	case EndArcShapeV2Grid:
		return any(&stage.EndArcShapeV2Grids_mapString).(*map[string]*Type)
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
	case StackGrowthCurveBezierShape:
		return any(&stage.StackGrowthCurveBezierShapes_mapString).(*map[string]*Type)
	case StackGrowthCurveEndArcShapeV2:
		return any(&stage.StackGrowthCurveEndArcShapeV2s_mapString).(*map[string]*Type)
	case StackGrowthCurveStartArcShapeV2:
		return any(&stage.StackGrowthCurveStartArcShapeV2s_mapString).(*map[string]*Type)
	case StackOfGrowthCurve:
		return any(&stage.StackOfGrowthCurves_mapString).(*map[string]*Type)
	case StackOfGrowthCurveV2:
		return any(&stage.StackOfGrowthCurveV2s_mapString).(*map[string]*Type)
	case StartArcShapeV2:
		return any(&stage.StartArcShapeV2s_mapString).(*map[string]*Type)
	case StartArcShapeV2Grid:
		return any(&stage.StartArcShapeV2Grids_mapString).(*map[string]*Type)
	case TopEndArcShapeV2:
		return any(&stage.TopEndArcShapeV2s_mapString).(*map[string]*Type)
	case TopEndArcShapeV2Grid:
		return any(&stage.TopEndArcShapeV2Grids_mapString).(*map[string]*Type)
	case TopGrowthCurve2D:
		return any(&stage.TopGrowthCurve2Ds_mapString).(*map[string]*Type)
	case TopStackGrowthCurveEndArcShapeV2:
		return any(&stage.TopStackGrowthCurveEndArcShapeV2s_mapString).(*map[string]*Type)
	case TopStackGrowthCurveStartArcShapeV2:
		return any(&stage.TopStackGrowthCurveStartArcShapeV2s_mapString).(*map[string]*Type)
	case TopStackOfGrowthCurveV2:
		return any(&stage.TopStackOfGrowthCurveV2s_mapString).(*map[string]*Type)
	case TopStartArcShapeV2:
		return any(&stage.TopStartArcShapeV2s_mapString).(*map[string]*Type)
	case TopStartArcShapeV2Grid:
		return any(&stage.TopStartArcShapeV2Grids_mapString).(*map[string]*Type)
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
	case BottomEndArcShapeV2:
		return any(&BottomEndArcShapeV2{
			// Initialisation of associations
		}).(*Type)
	case BottomEndArcShapeV2Grid:
		return any(&BottomEndArcShapeV2Grid{
			// Initialisation of associations
			// field is initialized with an instance of BottomEndArcShapeV2 with the name of the field
			BottomEndArcShapesV2: []*BottomEndArcShapeV2{{Name: "BottomEndArcShapesV2"}},
		}).(*Type)
	case BottomStackGrowthCurveEndArcShapeV2:
		return any(&BottomStackGrowthCurveEndArcShapeV2{
			// Initialisation of associations
		}).(*Type)
	case BottomStackGrowthCurveStartArcShapeV2:
		return any(&BottomStackGrowthCurveStartArcShapeV2{
			// Initialisation of associations
		}).(*Type)
	case BottomStackOfGrowthCurveV2:
		return any(&BottomStackOfGrowthCurveV2{
			// Initialisation of associations
			// field is initialized with an instance of BottomStackGrowthCurveStartArcShapeV2 with the name of the field
			BottomStackGrowthCurveStartArcShapeV2s: []*BottomStackGrowthCurveStartArcShapeV2{{Name: "BottomStackGrowthCurveStartArcShapeV2s"}},
			// field is initialized with an instance of BottomStackGrowthCurveEndArcShapeV2 with the name of the field
			BottomStackGrowthCurveEndArcShapeV2s: []*BottomStackGrowthCurveEndArcShapeV2{{Name: "BottomStackGrowthCurveEndArcShapeV2s"}},
		}).(*Type)
	case BottomStartArcShapeV2:
		return any(&BottomStartArcShapeV2{
			// Initialisation of associations
		}).(*Type)
	case BottomStartArcShapeV2Grid:
		return any(&BottomStartArcShapeV2Grid{
			// Initialisation of associations
			// field is initialized with an instance of BottomStartArcShapeV2 with the name of the field
			BottomStartArcShapesV2: []*BottomStartArcShapeV2{{Name: "BottomStartArcShapesV2"}},
		}).(*Type)
	case CircleGridShape:
		return any(&CircleGridShape{
			// Initialisation of associations
		}).(*Type)
	case EndArcShapeV2:
		return any(&EndArcShapeV2{
			// Initialisation of associations
		}).(*Type)
	case EndArcShapeV2Grid:
		return any(&EndArcShapeV2Grid{
			// Initialisation of associations
			// field is initialized with an instance of EndArcShapeV2 with the name of the field
			EndArcShapesV2: []*EndArcShapeV2{{Name: "EndArcShapesV2"}},
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
			// field is initialized with an instance of StartArcShapeV2Grid with the name of the field
			StartArcShapeV2Grid: &StartArcShapeV2Grid{Name: "StartArcShapeV2Grid"},
			// field is initialized with an instance of EndArcShapeV2Grid with the name of the field
			EndArcShapeV2Grid: &EndArcShapeV2Grid{Name: "EndArcShapeV2Grid"},
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
			// field is initialized with an instance of StartArcShapeV2Grid with the name of the field
			StartArcShapeV2Grid: &StartArcShapeV2Grid{Name: "StartArcShapeV2Grid"},
			// field is initialized with an instance of TopStartArcShapeV2Grid with the name of the field
			TopStartArcShapeV2Grid: &TopStartArcShapeV2Grid{Name: "TopStartArcShapeV2Grid"},
			// field is initialized with an instance of EndArcShapeV2Grid with the name of the field
			EndArcShapeV2Grid: &EndArcShapeV2Grid{Name: "EndArcShapeV2Grid"},
			// field is initialized with an instance of TopEndArcShapeV2Grid with the name of the field
			TopEndArcShapeV2Grid: &TopEndArcShapeV2Grid{Name: "TopEndArcShapeV2Grid"},
			// field is initialized with an instance of BottomStartArcShapeV2Grid with the name of the field
			BottomStartArcShapeV2Grid: &BottomStartArcShapeV2Grid{Name: "BottomStartArcShapeV2Grid"},
			// field is initialized with an instance of BottomEndArcShapeV2Grid with the name of the field
			BottomEndArcShapeV2Grid: &BottomEndArcShapeV2Grid{Name: "BottomEndArcShapeV2Grid"},
			// field is initialized with an instance of GrowthCurveBezierShapeGrid with the name of the field
			GrowthCurveBezierShapeGrid: &GrowthCurveBezierShapeGrid{Name: "GrowthCurveBezierShapeGrid"},
			// field is initialized with an instance of StackOfGrowthCurve with the name of the field
			StackOfGrowthCurve: &StackOfGrowthCurve{Name: "StackOfGrowthCurve"},
			// field is initialized with an instance of StackOfGrowthCurveV2 with the name of the field
			StackOfGrowthCurveV2: &StackOfGrowthCurveV2{Name: "StackOfGrowthCurveV2"},
			// field is initialized with an instance of TopStackOfGrowthCurveV2 with the name of the field
			TopStackOfGrowthCurveV2: &TopStackOfGrowthCurveV2{Name: "TopStackOfGrowthCurveV2"},
			// field is initialized with an instance of BottomStackOfGrowthCurveV2 with the name of the field
			BottomStackOfGrowthCurveV2: &BottomStackOfGrowthCurveV2{Name: "BottomStackOfGrowthCurveV2"},
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
	case StackGrowthCurveBezierShape:
		return any(&StackGrowthCurveBezierShape{
			// Initialisation of associations
		}).(*Type)
	case StackGrowthCurveEndArcShapeV2:
		return any(&StackGrowthCurveEndArcShapeV2{
			// Initialisation of associations
		}).(*Type)
	case StackGrowthCurveStartArcShapeV2:
		return any(&StackGrowthCurveStartArcShapeV2{
			// Initialisation of associations
		}).(*Type)
	case StackOfGrowthCurve:
		return any(&StackOfGrowthCurve{
			// Initialisation of associations
			// field is initialized with an instance of StackGrowthCurveBezierShape with the name of the field
			StackGrowthCurveBezierShapes: []*StackGrowthCurveBezierShape{{Name: "StackGrowthCurveBezierShapes"}},
		}).(*Type)
	case StackOfGrowthCurveV2:
		return any(&StackOfGrowthCurveV2{
			// Initialisation of associations
			// field is initialized with an instance of StackGrowthCurveStartArcShapeV2 with the name of the field
			StackGrowthCurveStartArcShapeV2s: []*StackGrowthCurveStartArcShapeV2{{Name: "StackGrowthCurveStartArcShapeV2s"}},
			// field is initialized with an instance of StackGrowthCurveEndArcShapeV2 with the name of the field
			StackGrowthCurveEndArcShapeV2s: []*StackGrowthCurveEndArcShapeV2{{Name: "StackGrowthCurveEndArcShapeV2s"}},
		}).(*Type)
	case StartArcShapeV2:
		return any(&StartArcShapeV2{
			// Initialisation of associations
		}).(*Type)
	case StartArcShapeV2Grid:
		return any(&StartArcShapeV2Grid{
			// Initialisation of associations
			// field is initialized with an instance of StartArcShapeV2 with the name of the field
			StartArcShapesV2: []*StartArcShapeV2{{Name: "StartArcShapesV2"}},
		}).(*Type)
	case TopEndArcShapeV2:
		return any(&TopEndArcShapeV2{
			// Initialisation of associations
		}).(*Type)
	case TopEndArcShapeV2Grid:
		return any(&TopEndArcShapeV2Grid{
			// Initialisation of associations
			// field is initialized with an instance of TopEndArcShapeV2 with the name of the field
			TopEndArcShapesV2: []*TopEndArcShapeV2{{Name: "TopEndArcShapesV2"}},
		}).(*Type)
	case TopGrowthCurve2D:
		return any(&TopGrowthCurve2D{
			// Initialisation of associations
			// field is initialized with an instance of TopStartArcShapeV2Grid with the name of the field
			TopStartArcShapeV2Grid: &TopStartArcShapeV2Grid{Name: "TopStartArcShapeV2Grid"},
			// field is initialized with an instance of TopEndArcShapeV2Grid with the name of the field
			TopEndArcShapeV2Grid: &TopEndArcShapeV2Grid{Name: "TopEndArcShapeV2Grid"},
		}).(*Type)
	case TopStackGrowthCurveEndArcShapeV2:
		return any(&TopStackGrowthCurveEndArcShapeV2{
			// Initialisation of associations
		}).(*Type)
	case TopStackGrowthCurveStartArcShapeV2:
		return any(&TopStackGrowthCurveStartArcShapeV2{
			// Initialisation of associations
		}).(*Type)
	case TopStackOfGrowthCurveV2:
		return any(&TopStackOfGrowthCurveV2{
			// Initialisation of associations
			// field is initialized with an instance of TopStackGrowthCurveStartArcShapeV2 with the name of the field
			TopStackGrowthCurveStartArcShapeV2s: []*TopStackGrowthCurveStartArcShapeV2{{Name: "TopStackGrowthCurveStartArcShapeV2s"}},
			// field is initialized with an instance of TopStackGrowthCurveEndArcShapeV2 with the name of the field
			TopStackGrowthCurveEndArcShapeV2s: []*TopStackGrowthCurveEndArcShapeV2{{Name: "TopStackGrowthCurveEndArcShapeV2s"}},
		}).(*Type)
	case TopStartArcShapeV2:
		return any(&TopStartArcShapeV2{
			// Initialisation of associations
		}).(*Type)
	case TopStartArcShapeV2Grid:
		return any(&TopStartArcShapeV2Grid{
			// Initialisation of associations
			// field is initialized with an instance of TopStartArcShapeV2 with the name of the field
			TopStartArcShapesV2: []*TopStartArcShapeV2{{Name: "TopStartArcShapesV2"}},
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
	// reverse maps of direct associations of BottomEndArcShapeV2
	case BottomEndArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BottomEndArcShapeV2Grid
	case BottomEndArcShapeV2Grid:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BottomStackGrowthCurveEndArcShapeV2
	case BottomStackGrowthCurveEndArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BottomStackGrowthCurveStartArcShapeV2
	case BottomStackGrowthCurveStartArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BottomStackOfGrowthCurveV2
	case BottomStackOfGrowthCurveV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BottomStartArcShapeV2
	case BottomStartArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BottomStartArcShapeV2Grid
	case BottomStartArcShapeV2Grid:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CircleGridShape
	case CircleGridShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EndArcShapeV2
	case EndArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EndArcShapeV2Grid
	case EndArcShapeV2Grid:
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
		case "StartArcShapeV2Grid":
			res := make(map[*StartArcShapeV2Grid][]*GrowthCurve2D)
			for growthcurve2d := range stage.GrowthCurve2Ds {
				if growthcurve2d.StartArcShapeV2Grid != nil {
					startarcshapev2grid_ := growthcurve2d.StartArcShapeV2Grid
					var growthcurve2ds []*GrowthCurve2D
					_, ok := res[startarcshapev2grid_]
					if ok {
						growthcurve2ds = res[startarcshapev2grid_]
					} else {
						growthcurve2ds = make([]*GrowthCurve2D, 0)
					}
					growthcurve2ds = append(growthcurve2ds, growthcurve2d)
					res[startarcshapev2grid_] = growthcurve2ds
				}
			}
			return any(res).(map[*End][]*Start)
		case "EndArcShapeV2Grid":
			res := make(map[*EndArcShapeV2Grid][]*GrowthCurve2D)
			for growthcurve2d := range stage.GrowthCurve2Ds {
				if growthcurve2d.EndArcShapeV2Grid != nil {
					endarcshapev2grid_ := growthcurve2d.EndArcShapeV2Grid
					var growthcurve2ds []*GrowthCurve2D
					_, ok := res[endarcshapev2grid_]
					if ok {
						growthcurve2ds = res[endarcshapev2grid_]
					} else {
						growthcurve2ds = make([]*GrowthCurve2D, 0)
					}
					growthcurve2ds = append(growthcurve2ds, growthcurve2d)
					res[endarcshapev2grid_] = growthcurve2ds
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
		case "StartArcShapeV2Grid":
			res := make(map[*StartArcShapeV2Grid][]*Plant)
			for plant := range stage.Plants {
				if plant.StartArcShapeV2Grid != nil {
					startarcshapev2grid_ := plant.StartArcShapeV2Grid
					var plants []*Plant
					_, ok := res[startarcshapev2grid_]
					if ok {
						plants = res[startarcshapev2grid_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[startarcshapev2grid_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "TopStartArcShapeV2Grid":
			res := make(map[*TopStartArcShapeV2Grid][]*Plant)
			for plant := range stage.Plants {
				if plant.TopStartArcShapeV2Grid != nil {
					topstartarcshapev2grid_ := plant.TopStartArcShapeV2Grid
					var plants []*Plant
					_, ok := res[topstartarcshapev2grid_]
					if ok {
						plants = res[topstartarcshapev2grid_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[topstartarcshapev2grid_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "EndArcShapeV2Grid":
			res := make(map[*EndArcShapeV2Grid][]*Plant)
			for plant := range stage.Plants {
				if plant.EndArcShapeV2Grid != nil {
					endarcshapev2grid_ := plant.EndArcShapeV2Grid
					var plants []*Plant
					_, ok := res[endarcshapev2grid_]
					if ok {
						plants = res[endarcshapev2grid_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[endarcshapev2grid_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "TopEndArcShapeV2Grid":
			res := make(map[*TopEndArcShapeV2Grid][]*Plant)
			for plant := range stage.Plants {
				if plant.TopEndArcShapeV2Grid != nil {
					topendarcshapev2grid_ := plant.TopEndArcShapeV2Grid
					var plants []*Plant
					_, ok := res[topendarcshapev2grid_]
					if ok {
						plants = res[topendarcshapev2grid_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[topendarcshapev2grid_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "BottomStartArcShapeV2Grid":
			res := make(map[*BottomStartArcShapeV2Grid][]*Plant)
			for plant := range stage.Plants {
				if plant.BottomStartArcShapeV2Grid != nil {
					bottomstartarcshapev2grid_ := plant.BottomStartArcShapeV2Grid
					var plants []*Plant
					_, ok := res[bottomstartarcshapev2grid_]
					if ok {
						plants = res[bottomstartarcshapev2grid_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[bottomstartarcshapev2grid_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "BottomEndArcShapeV2Grid":
			res := make(map[*BottomEndArcShapeV2Grid][]*Plant)
			for plant := range stage.Plants {
				if plant.BottomEndArcShapeV2Grid != nil {
					bottomendarcshapev2grid_ := plant.BottomEndArcShapeV2Grid
					var plants []*Plant
					_, ok := res[bottomendarcshapev2grid_]
					if ok {
						plants = res[bottomendarcshapev2grid_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[bottomendarcshapev2grid_] = plants
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
		case "StackOfGrowthCurveV2":
			res := make(map[*StackOfGrowthCurveV2][]*Plant)
			for plant := range stage.Plants {
				if plant.StackOfGrowthCurveV2 != nil {
					stackofgrowthcurvev2_ := plant.StackOfGrowthCurveV2
					var plants []*Plant
					_, ok := res[stackofgrowthcurvev2_]
					if ok {
						plants = res[stackofgrowthcurvev2_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[stackofgrowthcurvev2_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "TopStackOfGrowthCurveV2":
			res := make(map[*TopStackOfGrowthCurveV2][]*Plant)
			for plant := range stage.Plants {
				if plant.TopStackOfGrowthCurveV2 != nil {
					topstackofgrowthcurvev2_ := plant.TopStackOfGrowthCurveV2
					var plants []*Plant
					_, ok := res[topstackofgrowthcurvev2_]
					if ok {
						plants = res[topstackofgrowthcurvev2_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[topstackofgrowthcurvev2_] = plants
				}
			}
			return any(res).(map[*End][]*Start)
		case "BottomStackOfGrowthCurveV2":
			res := make(map[*BottomStackOfGrowthCurveV2][]*Plant)
			for plant := range stage.Plants {
				if plant.BottomStackOfGrowthCurveV2 != nil {
					bottomstackofgrowthcurvev2_ := plant.BottomStackOfGrowthCurveV2
					var plants []*Plant
					_, ok := res[bottomstackofgrowthcurvev2_]
					if ok {
						plants = res[bottomstackofgrowthcurvev2_]
					} else {
						plants = make([]*Plant, 0)
					}
					plants = append(plants, plant)
					res[bottomstackofgrowthcurvev2_] = plants
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
	// reverse maps of direct associations of StackGrowthCurveBezierShape
	case StackGrowthCurveBezierShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StackGrowthCurveEndArcShapeV2
	case StackGrowthCurveEndArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StackGrowthCurveStartArcShapeV2
	case StackGrowthCurveStartArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StackOfGrowthCurve
	case StackOfGrowthCurve:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StackOfGrowthCurveV2
	case StackOfGrowthCurveV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StartArcShapeV2
	case StartArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StartArcShapeV2Grid
	case StartArcShapeV2Grid:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopEndArcShapeV2
	case TopEndArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopEndArcShapeV2Grid
	case TopEndArcShapeV2Grid:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopGrowthCurve2D
	case TopGrowthCurve2D:
		switch fieldname {
		// insertion point for per direct association field
		case "TopStartArcShapeV2Grid":
			res := make(map[*TopStartArcShapeV2Grid][]*TopGrowthCurve2D)
			for topgrowthcurve2d := range stage.TopGrowthCurve2Ds {
				if topgrowthcurve2d.TopStartArcShapeV2Grid != nil {
					topstartarcshapev2grid_ := topgrowthcurve2d.TopStartArcShapeV2Grid
					var topgrowthcurve2ds []*TopGrowthCurve2D
					_, ok := res[topstartarcshapev2grid_]
					if ok {
						topgrowthcurve2ds = res[topstartarcshapev2grid_]
					} else {
						topgrowthcurve2ds = make([]*TopGrowthCurve2D, 0)
					}
					topgrowthcurve2ds = append(topgrowthcurve2ds, topgrowthcurve2d)
					res[topstartarcshapev2grid_] = topgrowthcurve2ds
				}
			}
			return any(res).(map[*End][]*Start)
		case "TopEndArcShapeV2Grid":
			res := make(map[*TopEndArcShapeV2Grid][]*TopGrowthCurve2D)
			for topgrowthcurve2d := range stage.TopGrowthCurve2Ds {
				if topgrowthcurve2d.TopEndArcShapeV2Grid != nil {
					topendarcshapev2grid_ := topgrowthcurve2d.TopEndArcShapeV2Grid
					var topgrowthcurve2ds []*TopGrowthCurve2D
					_, ok := res[topendarcshapev2grid_]
					if ok {
						topgrowthcurve2ds = res[topendarcshapev2grid_]
					} else {
						topgrowthcurve2ds = make([]*TopGrowthCurve2D, 0)
					}
					topgrowthcurve2ds = append(topgrowthcurve2ds, topgrowthcurve2d)
					res[topendarcshapev2grid_] = topgrowthcurve2ds
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TopStackGrowthCurveEndArcShapeV2
	case TopStackGrowthCurveEndArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopStackGrowthCurveStartArcShapeV2
	case TopStackGrowthCurveStartArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopStackOfGrowthCurveV2
	case TopStackOfGrowthCurveV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopStartArcShapeV2
	case TopStartArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopStartArcShapeV2Grid
	case TopStartArcShapeV2Grid:
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
	// reverse maps of direct associations of BottomEndArcShapeV2
	case BottomEndArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BottomEndArcShapeV2Grid
	case BottomEndArcShapeV2Grid:
		switch fieldname {
		// insertion point for per direct association field
		case "BottomEndArcShapesV2":
			res := make(map[*BottomEndArcShapeV2][]*BottomEndArcShapeV2Grid)
			for bottomendarcshapev2grid := range stage.BottomEndArcShapeV2Grids {
				for _, bottomendarcshapev2_ := range bottomendarcshapev2grid.BottomEndArcShapesV2 {
					res[bottomendarcshapev2_] = append(res[bottomendarcshapev2_], bottomendarcshapev2grid)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of BottomStackGrowthCurveEndArcShapeV2
	case BottomStackGrowthCurveEndArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BottomStackGrowthCurveStartArcShapeV2
	case BottomStackGrowthCurveStartArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BottomStackOfGrowthCurveV2
	case BottomStackOfGrowthCurveV2:
		switch fieldname {
		// insertion point for per direct association field
		case "BottomStackGrowthCurveStartArcShapeV2s":
			res := make(map[*BottomStackGrowthCurveStartArcShapeV2][]*BottomStackOfGrowthCurveV2)
			for bottomstackofgrowthcurvev2 := range stage.BottomStackOfGrowthCurveV2s {
				for _, bottomstackgrowthcurvestartarcshapev2_ := range bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s {
					res[bottomstackgrowthcurvestartarcshapev2_] = append(res[bottomstackgrowthcurvestartarcshapev2_], bottomstackofgrowthcurvev2)
				}
			}
			return any(res).(map[*End][]*Start)
		case "BottomStackGrowthCurveEndArcShapeV2s":
			res := make(map[*BottomStackGrowthCurveEndArcShapeV2][]*BottomStackOfGrowthCurveV2)
			for bottomstackofgrowthcurvev2 := range stage.BottomStackOfGrowthCurveV2s {
				for _, bottomstackgrowthcurveendarcshapev2_ := range bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s {
					res[bottomstackgrowthcurveendarcshapev2_] = append(res[bottomstackgrowthcurveendarcshapev2_], bottomstackofgrowthcurvev2)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of BottomStartArcShapeV2
	case BottomStartArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BottomStartArcShapeV2Grid
	case BottomStartArcShapeV2Grid:
		switch fieldname {
		// insertion point for per direct association field
		case "BottomStartArcShapesV2":
			res := make(map[*BottomStartArcShapeV2][]*BottomStartArcShapeV2Grid)
			for bottomstartarcshapev2grid := range stage.BottomStartArcShapeV2Grids {
				for _, bottomstartarcshapev2_ := range bottomstartarcshapev2grid.BottomStartArcShapesV2 {
					res[bottomstartarcshapev2_] = append(res[bottomstartarcshapev2_], bottomstartarcshapev2grid)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of CircleGridShape
	case CircleGridShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EndArcShapeV2
	case EndArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EndArcShapeV2Grid
	case EndArcShapeV2Grid:
		switch fieldname {
		// insertion point for per direct association field
		case "EndArcShapesV2":
			res := make(map[*EndArcShapeV2][]*EndArcShapeV2Grid)
			for endarcshapev2grid := range stage.EndArcShapeV2Grids {
				for _, endarcshapev2_ := range endarcshapev2grid.EndArcShapesV2 {
					res[endarcshapev2_] = append(res[endarcshapev2_], endarcshapev2grid)
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
	// reverse maps of direct associations of StackGrowthCurveBezierShape
	case StackGrowthCurveBezierShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StackGrowthCurveEndArcShapeV2
	case StackGrowthCurveEndArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StackGrowthCurveStartArcShapeV2
	case StackGrowthCurveStartArcShapeV2:
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
	// reverse maps of direct associations of StackOfGrowthCurveV2
	case StackOfGrowthCurveV2:
		switch fieldname {
		// insertion point for per direct association field
		case "StackGrowthCurveStartArcShapeV2s":
			res := make(map[*StackGrowthCurveStartArcShapeV2][]*StackOfGrowthCurveV2)
			for stackofgrowthcurvev2 := range stage.StackOfGrowthCurveV2s {
				for _, stackgrowthcurvestartarcshapev2_ := range stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s {
					res[stackgrowthcurvestartarcshapev2_] = append(res[stackgrowthcurvestartarcshapev2_], stackofgrowthcurvev2)
				}
			}
			return any(res).(map[*End][]*Start)
		case "StackGrowthCurveEndArcShapeV2s":
			res := make(map[*StackGrowthCurveEndArcShapeV2][]*StackOfGrowthCurveV2)
			for stackofgrowthcurvev2 := range stage.StackOfGrowthCurveV2s {
				for _, stackgrowthcurveendarcshapev2_ := range stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s {
					res[stackgrowthcurveendarcshapev2_] = append(res[stackgrowthcurveendarcshapev2_], stackofgrowthcurvev2)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of StartArcShapeV2
	case StartArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StartArcShapeV2Grid
	case StartArcShapeV2Grid:
		switch fieldname {
		// insertion point for per direct association field
		case "StartArcShapesV2":
			res := make(map[*StartArcShapeV2][]*StartArcShapeV2Grid)
			for startarcshapev2grid := range stage.StartArcShapeV2Grids {
				for _, startarcshapev2_ := range startarcshapev2grid.StartArcShapesV2 {
					res[startarcshapev2_] = append(res[startarcshapev2_], startarcshapev2grid)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TopEndArcShapeV2
	case TopEndArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopEndArcShapeV2Grid
	case TopEndArcShapeV2Grid:
		switch fieldname {
		// insertion point for per direct association field
		case "TopEndArcShapesV2":
			res := make(map[*TopEndArcShapeV2][]*TopEndArcShapeV2Grid)
			for topendarcshapev2grid := range stage.TopEndArcShapeV2Grids {
				for _, topendarcshapev2_ := range topendarcshapev2grid.TopEndArcShapesV2 {
					res[topendarcshapev2_] = append(res[topendarcshapev2_], topendarcshapev2grid)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TopGrowthCurve2D
	case TopGrowthCurve2D:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopStackGrowthCurveEndArcShapeV2
	case TopStackGrowthCurveEndArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopStackGrowthCurveStartArcShapeV2
	case TopStackGrowthCurveStartArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopStackOfGrowthCurveV2
	case TopStackOfGrowthCurveV2:
		switch fieldname {
		// insertion point for per direct association field
		case "TopStackGrowthCurveStartArcShapeV2s":
			res := make(map[*TopStackGrowthCurveStartArcShapeV2][]*TopStackOfGrowthCurveV2)
			for topstackofgrowthcurvev2 := range stage.TopStackOfGrowthCurveV2s {
				for _, topstackgrowthcurvestartarcshapev2_ := range topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s {
					res[topstackgrowthcurvestartarcshapev2_] = append(res[topstackgrowthcurvestartarcshapev2_], topstackofgrowthcurvev2)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TopStackGrowthCurveEndArcShapeV2s":
			res := make(map[*TopStackGrowthCurveEndArcShapeV2][]*TopStackOfGrowthCurveV2)
			for topstackofgrowthcurvev2 := range stage.TopStackOfGrowthCurveV2s {
				for _, topstackgrowthcurveendarcshapev2_ := range topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s {
					res[topstackgrowthcurveendarcshapev2_] = append(res[topstackgrowthcurveendarcshapev2_], topstackofgrowthcurvev2)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TopStartArcShapeV2
	case TopStartArcShapeV2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TopStartArcShapeV2Grid
	case TopStartArcShapeV2Grid:
		switch fieldname {
		// insertion point for per direct association field
		case "TopStartArcShapesV2":
			res := make(map[*TopStartArcShapeV2][]*TopStartArcShapeV2Grid)
			for topstartarcshapev2grid := range stage.TopStartArcShapeV2Grids {
				for _, topstartarcshapev2_ := range topstartarcshapev2grid.TopStartArcShapesV2 {
					res[topstartarcshapev2_] = append(res[topstartarcshapev2_], topstartarcshapev2grid)
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
	case *BottomEndArcShapeV2:
		res = "BottomEndArcShapeV2"
	case *BottomEndArcShapeV2Grid:
		res = "BottomEndArcShapeV2Grid"
	case *BottomStackGrowthCurveEndArcShapeV2:
		res = "BottomStackGrowthCurveEndArcShapeV2"
	case *BottomStackGrowthCurveStartArcShapeV2:
		res = "BottomStackGrowthCurveStartArcShapeV2"
	case *BottomStackOfGrowthCurveV2:
		res = "BottomStackOfGrowthCurveV2"
	case *BottomStartArcShapeV2:
		res = "BottomStartArcShapeV2"
	case *BottomStartArcShapeV2Grid:
		res = "BottomStartArcShapeV2Grid"
	case *CircleGridShape:
		res = "CircleGridShape"
	case *EndArcShapeV2:
		res = "EndArcShapeV2"
	case *EndArcShapeV2Grid:
		res = "EndArcShapeV2Grid"
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
	case *StackGrowthCurveBezierShape:
		res = "StackGrowthCurveBezierShape"
	case *StackGrowthCurveEndArcShapeV2:
		res = "StackGrowthCurveEndArcShapeV2"
	case *StackGrowthCurveStartArcShapeV2:
		res = "StackGrowthCurveStartArcShapeV2"
	case *StackOfGrowthCurve:
		res = "StackOfGrowthCurve"
	case *StackOfGrowthCurveV2:
		res = "StackOfGrowthCurveV2"
	case *StartArcShapeV2:
		res = "StartArcShapeV2"
	case *StartArcShapeV2Grid:
		res = "StartArcShapeV2Grid"
	case *TopEndArcShapeV2:
		res = "TopEndArcShapeV2"
	case *TopEndArcShapeV2Grid:
		res = "TopEndArcShapeV2Grid"
	case *TopGrowthCurve2D:
		res = "TopGrowthCurve2D"
	case *TopStackGrowthCurveEndArcShapeV2:
		res = "TopStackGrowthCurveEndArcShapeV2"
	case *TopStackGrowthCurveStartArcShapeV2:
		res = "TopStackGrowthCurveStartArcShapeV2"
	case *TopStackOfGrowthCurveV2:
		res = "TopStackOfGrowthCurveV2"
	case *TopStartArcShapeV2:
		res = "TopStartArcShapeV2"
	case *TopStartArcShapeV2Grid:
		res = "TopStartArcShapeV2Grid"
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
	case *BottomEndArcShapeV2:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "BottomEndArcShapeV2Grid"
		rf.Fieldname = "BottomEndArcShapesV2"
		res = append(res, rf)
	case *BottomEndArcShapeV2Grid:
		var rf ReverseField
		_ = rf
	case *BottomStackGrowthCurveEndArcShapeV2:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "BottomStackOfGrowthCurveV2"
		rf.Fieldname = "BottomStackGrowthCurveEndArcShapeV2s"
		res = append(res, rf)
	case *BottomStackGrowthCurveStartArcShapeV2:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "BottomStackOfGrowthCurveV2"
		rf.Fieldname = "BottomStackGrowthCurveStartArcShapeV2s"
		res = append(res, rf)
	case *BottomStackOfGrowthCurveV2:
		var rf ReverseField
		_ = rf
	case *BottomStartArcShapeV2:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "BottomStartArcShapeV2Grid"
		rf.Fieldname = "BottomStartArcShapesV2"
		res = append(res, rf)
	case *BottomStartArcShapeV2Grid:
		var rf ReverseField
		_ = rf
	case *CircleGridShape:
		var rf ReverseField
		_ = rf
	case *EndArcShapeV2:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "EndArcShapeV2Grid"
		rf.Fieldname = "EndArcShapesV2"
		res = append(res, rf)
	case *EndArcShapeV2Grid:
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
	case *StackGrowthCurveBezierShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "StackOfGrowthCurve"
		rf.Fieldname = "StackGrowthCurveBezierShapes"
		res = append(res, rf)
	case *StackGrowthCurveEndArcShapeV2:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "StackOfGrowthCurveV2"
		rf.Fieldname = "StackGrowthCurveEndArcShapeV2s"
		res = append(res, rf)
	case *StackGrowthCurveStartArcShapeV2:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "StackOfGrowthCurveV2"
		rf.Fieldname = "StackGrowthCurveStartArcShapeV2s"
		res = append(res, rf)
	case *StackOfGrowthCurve:
		var rf ReverseField
		_ = rf
	case *StackOfGrowthCurveV2:
		var rf ReverseField
		_ = rf
	case *StartArcShapeV2:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "StartArcShapeV2Grid"
		rf.Fieldname = "StartArcShapesV2"
		res = append(res, rf)
	case *StartArcShapeV2Grid:
		var rf ReverseField
		_ = rf
	case *TopEndArcShapeV2:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "TopEndArcShapeV2Grid"
		rf.Fieldname = "TopEndArcShapesV2"
		res = append(res, rf)
	case *TopEndArcShapeV2Grid:
		var rf ReverseField
		_ = rf
	case *TopGrowthCurve2D:
		var rf ReverseField
		_ = rf
	case *TopStackGrowthCurveEndArcShapeV2:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "TopStackOfGrowthCurveV2"
		rf.Fieldname = "TopStackGrowthCurveEndArcShapeV2s"
		res = append(res, rf)
	case *TopStackGrowthCurveStartArcShapeV2:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "TopStackOfGrowthCurveV2"
		rf.Fieldname = "TopStackGrowthCurveStartArcShapeV2s"
		res = append(res, rf)
	case *TopStackOfGrowthCurveV2:
		var rf ReverseField
		_ = rf
	case *TopStartArcShapeV2:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "TopStartArcShapeV2Grid"
		rf.Fieldname = "TopStartArcShapesV2"
		res = append(res, rf)
	case *TopStartArcShapeV2Grid:
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

func (bottomendarcshapev2 *BottomEndArcShapeV2) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "BottomEndArcShapesV2",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "BottomEndArcShapeV2",
		},
	}
	return
}

func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "BottomStackGrowthCurveStartArcShapeV2s",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "BottomStackGrowthCurveStartArcShapeV2",
		},
		{
			Name:                 "BottomStackGrowthCurveEndArcShapeV2s",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "BottomStackGrowthCurveEndArcShapeV2",
		},
	}
	return
}

func (bottomstartarcshapev2 *BottomStartArcShapeV2) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "BottomStartArcShapesV2",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "BottomStartArcShapeV2",
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

func (endarcshapev2 *EndArcShapeV2) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (endarcshapev2grid *EndArcShapeV2Grid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "EndArcShapesV2",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "EndArcShapeV2",
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
			Name:                 "StartArcShapeV2Grid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "StartArcShapeV2Grid",
		},
		{
			Name:                 "EndArcShapeV2Grid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "EndArcShapeV2Grid",
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
			Name:               "Thickness",
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
			Name:                 "StartArcShapeV2Grid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "StartArcShapeV2Grid",
		},
		{
			Name:                 "TopStartArcShapeV2Grid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TopStartArcShapeV2Grid",
		},
		{
			Name:                 "EndArcShapeV2Grid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "EndArcShapeV2Grid",
		},
		{
			Name:                 "TopEndArcShapeV2Grid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TopEndArcShapeV2Grid",
		},
		{
			Name:                 "BottomStartArcShapeV2Grid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "BottomStartArcShapeV2Grid",
		},
		{
			Name:                 "BottomEndArcShapeV2Grid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "BottomEndArcShapeV2Grid",
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
			Name:                 "StackOfGrowthCurveV2",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "StackOfGrowthCurveV2",
		},
		{
			Name:                 "TopStackOfGrowthCurveV2",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TopStackOfGrowthCurveV2",
		},
		{
			Name:                 "BottomStackOfGrowthCurveV2",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "BottomStackOfGrowthCurveV2",
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
			Name:               "IsHiddenStartArcShapeV2Grid",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenTopStartArcShapeV2Grid",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenEndArcShapeGrid",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenEndArcShapeV2Grid",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenTopEndArcShapeV2Grid",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenBottomStartArcShapeV2Grid",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenBottomEndArcShapeV2Grid",
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
			Name:               "IsHiddenStackOfGrowthCurveV2",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenTopStackOfGrowthCurveV2",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsHiddenBottomStackOfGrowthCurveV2",
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

func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "StackGrowthCurveBezierShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "StackGrowthCurveBezierShape",
		},
	}
	return
}

func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "StackGrowthCurveStartArcShapeV2s",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "StackGrowthCurveStartArcShapeV2",
		},
		{
			Name:                 "StackGrowthCurveEndArcShapeV2s",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "StackGrowthCurveEndArcShapeV2",
		},
	}
	return
}

func (startarcshapev2 *StartArcShapeV2) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (startarcshapev2grid *StartArcShapeV2Grid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "StartArcShapesV2",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "StartArcShapeV2",
		},
	}
	return
}

func (topendarcshapev2 *TopEndArcShapeV2) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (topendarcshapev2grid *TopEndArcShapeV2Grid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "TopEndArcShapesV2",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TopEndArcShapeV2",
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
			Name:                 "TopStartArcShapeV2Grid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TopStartArcShapeV2Grid",
		},
		{
			Name:                 "TopEndArcShapeV2Grid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TopEndArcShapeV2Grid",
		},
	}
	return
}

func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "TopStackGrowthCurveStartArcShapeV2s",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TopStackGrowthCurveStartArcShapeV2",
		},
		{
			Name:                 "TopStackGrowthCurveEndArcShapeV2s",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TopStackGrowthCurveEndArcShapeV2",
		},
	}
	return
}

func (topstartarcshapev2 *TopStartArcShapeV2) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (topstartarcshapev2grid *TopStartArcShapeV2Grid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "TopStartArcShapesV2",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TopStartArcShapeV2",
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

func (bottomendarcshapev2 *BottomEndArcShapeV2) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = bottomendarcshapev2.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", bottomendarcshapev2.StartX)
		res.valueFloat = bottomendarcshapev2.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", bottomendarcshapev2.StartY)
		res.valueFloat = bottomendarcshapev2.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", bottomendarcshapev2.EndX)
		res.valueFloat = bottomendarcshapev2.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", bottomendarcshapev2.EndY)
		res.valueFloat = bottomendarcshapev2.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", bottomendarcshapev2.XAxisRotation)
		res.valueFloat = bottomendarcshapev2.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", bottomendarcshapev2.LargeArcFlag)
		res.valueBool = bottomendarcshapev2.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", bottomendarcshapev2.SweepFlag)
		res.valueBool = bottomendarcshapev2.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", bottomendarcshapev2.RadiusX)
		res.valueFloat = bottomendarcshapev2.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", bottomendarcshapev2.RadiusY)
		res.valueFloat = bottomendarcshapev2.RadiusY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = bottomendarcshapev2grid.Name
	case "BottomEndArcShapesV2":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range bottomendarcshapev2grid.BottomEndArcShapesV2 {
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

func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = bottomstackgrowthcurveendarcshapev2.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", bottomstackgrowthcurveendarcshapev2.StartX)
		res.valueFloat = bottomstackgrowthcurveendarcshapev2.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", bottomstackgrowthcurveendarcshapev2.StartY)
		res.valueFloat = bottomstackgrowthcurveendarcshapev2.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", bottomstackgrowthcurveendarcshapev2.EndX)
		res.valueFloat = bottomstackgrowthcurveendarcshapev2.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", bottomstackgrowthcurveendarcshapev2.EndY)
		res.valueFloat = bottomstackgrowthcurveendarcshapev2.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", bottomstackgrowthcurveendarcshapev2.XAxisRotation)
		res.valueFloat = bottomstackgrowthcurveendarcshapev2.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", bottomstackgrowthcurveendarcshapev2.LargeArcFlag)
		res.valueBool = bottomstackgrowthcurveendarcshapev2.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", bottomstackgrowthcurveendarcshapev2.SweepFlag)
		res.valueBool = bottomstackgrowthcurveendarcshapev2.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", bottomstackgrowthcurveendarcshapev2.RadiusX)
		res.valueFloat = bottomstackgrowthcurveendarcshapev2.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", bottomstackgrowthcurveendarcshapev2.RadiusY)
		res.valueFloat = bottomstackgrowthcurveendarcshapev2.RadiusY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = bottomstackgrowthcurvestartarcshapev2.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", bottomstackgrowthcurvestartarcshapev2.StartX)
		res.valueFloat = bottomstackgrowthcurvestartarcshapev2.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", bottomstackgrowthcurvestartarcshapev2.StartY)
		res.valueFloat = bottomstackgrowthcurvestartarcshapev2.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", bottomstackgrowthcurvestartarcshapev2.EndX)
		res.valueFloat = bottomstackgrowthcurvestartarcshapev2.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", bottomstackgrowthcurvestartarcshapev2.EndY)
		res.valueFloat = bottomstackgrowthcurvestartarcshapev2.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", bottomstackgrowthcurvestartarcshapev2.XAxisRotation)
		res.valueFloat = bottomstackgrowthcurvestartarcshapev2.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", bottomstackgrowthcurvestartarcshapev2.LargeArcFlag)
		res.valueBool = bottomstackgrowthcurvestartarcshapev2.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", bottomstackgrowthcurvestartarcshapev2.SweepFlag)
		res.valueBool = bottomstackgrowthcurvestartarcshapev2.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", bottomstackgrowthcurvestartarcshapev2.RadiusX)
		res.valueFloat = bottomstackgrowthcurvestartarcshapev2.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", bottomstackgrowthcurvestartarcshapev2.RadiusY)
		res.valueFloat = bottomstackgrowthcurvestartarcshapev2.RadiusY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = bottomstackofgrowthcurvev2.Name
	case "BottomStackGrowthCurveStartArcShapeV2s":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "BottomStackGrowthCurveEndArcShapeV2s":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s {
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

func (bottomstartarcshapev2 *BottomStartArcShapeV2) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = bottomstartarcshapev2.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", bottomstartarcshapev2.StartX)
		res.valueFloat = bottomstartarcshapev2.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", bottomstartarcshapev2.StartY)
		res.valueFloat = bottomstartarcshapev2.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", bottomstartarcshapev2.EndX)
		res.valueFloat = bottomstartarcshapev2.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", bottomstartarcshapev2.EndY)
		res.valueFloat = bottomstartarcshapev2.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", bottomstartarcshapev2.XAxisRotation)
		res.valueFloat = bottomstartarcshapev2.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", bottomstartarcshapev2.LargeArcFlag)
		res.valueBool = bottomstartarcshapev2.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", bottomstartarcshapev2.SweepFlag)
		res.valueBool = bottomstartarcshapev2.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", bottomstartarcshapev2.RadiusX)
		res.valueFloat = bottomstartarcshapev2.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", bottomstartarcshapev2.RadiusY)
		res.valueFloat = bottomstartarcshapev2.RadiusY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = bottomstartarcshapev2grid.Name
	case "BottomStartArcShapesV2":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range bottomstartarcshapev2grid.BottomStartArcShapesV2 {
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

func (endarcshapev2 *EndArcShapeV2) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = endarcshapev2.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", endarcshapev2.StartX)
		res.valueFloat = endarcshapev2.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", endarcshapev2.StartY)
		res.valueFloat = endarcshapev2.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", endarcshapev2.EndX)
		res.valueFloat = endarcshapev2.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", endarcshapev2.EndY)
		res.valueFloat = endarcshapev2.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", endarcshapev2.XAxisRotation)
		res.valueFloat = endarcshapev2.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", endarcshapev2.LargeArcFlag)
		res.valueBool = endarcshapev2.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", endarcshapev2.SweepFlag)
		res.valueBool = endarcshapev2.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", endarcshapev2.RadiusX)
		res.valueFloat = endarcshapev2.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", endarcshapev2.RadiusY)
		res.valueFloat = endarcshapev2.RadiusY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (endarcshapev2grid *EndArcShapeV2Grid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = endarcshapev2grid.Name
	case "EndArcShapesV2":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range endarcshapev2grid.EndArcShapesV2 {
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
	case "StartArcShapeV2Grid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if growthcurve2d.StartArcShapeV2Grid != nil {
			res.valueString = growthcurve2d.StartArcShapeV2Grid.Name
			res.ids = growthcurve2d.StartArcShapeV2Grid.GongGetUUID(stage)
		}
	case "EndArcShapeV2Grid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if growthcurve2d.EndArcShapeV2Grid != nil {
			res.valueString = growthcurve2d.EndArcShapeV2Grid.Name
			res.ids = growthcurve2d.EndArcShapeV2Grid.GongGetUUID(stage)
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
	case "Thickness":
		res.valueString = fmt.Sprintf("%f", plant.Thickness)
		res.valueFloat = plant.Thickness
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
	case "StartArcShapeV2Grid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.StartArcShapeV2Grid != nil {
			res.valueString = plant.StartArcShapeV2Grid.Name
			res.ids = plant.StartArcShapeV2Grid.GongGetUUID(stage)
		}
	case "TopStartArcShapeV2Grid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.TopStartArcShapeV2Grid != nil {
			res.valueString = plant.TopStartArcShapeV2Grid.Name
			res.ids = plant.TopStartArcShapeV2Grid.GongGetUUID(stage)
		}
	case "EndArcShapeV2Grid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.EndArcShapeV2Grid != nil {
			res.valueString = plant.EndArcShapeV2Grid.Name
			res.ids = plant.EndArcShapeV2Grid.GongGetUUID(stage)
		}
	case "TopEndArcShapeV2Grid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.TopEndArcShapeV2Grid != nil {
			res.valueString = plant.TopEndArcShapeV2Grid.Name
			res.ids = plant.TopEndArcShapeV2Grid.GongGetUUID(stage)
		}
	case "BottomStartArcShapeV2Grid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.BottomStartArcShapeV2Grid != nil {
			res.valueString = plant.BottomStartArcShapeV2Grid.Name
			res.ids = plant.BottomStartArcShapeV2Grid.GongGetUUID(stage)
		}
	case "BottomEndArcShapeV2Grid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.BottomEndArcShapeV2Grid != nil {
			res.valueString = plant.BottomEndArcShapeV2Grid.Name
			res.ids = plant.BottomEndArcShapeV2Grid.GongGetUUID(stage)
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
	case "StackOfGrowthCurveV2":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.StackOfGrowthCurveV2 != nil {
			res.valueString = plant.StackOfGrowthCurveV2.Name
			res.ids = plant.StackOfGrowthCurveV2.GongGetUUID(stage)
		}
	case "TopStackOfGrowthCurveV2":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.TopStackOfGrowthCurveV2 != nil {
			res.valueString = plant.TopStackOfGrowthCurveV2.Name
			res.ids = plant.TopStackOfGrowthCurveV2.GongGetUUID(stage)
		}
	case "BottomStackOfGrowthCurveV2":
		res.GongFieldValueType = GongFieldValueTypePointer
		if plant.BottomStackOfGrowthCurveV2 != nil {
			res.valueString = plant.BottomStackOfGrowthCurveV2.Name
			res.ids = plant.BottomStackOfGrowthCurveV2.GongGetUUID(stage)
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
	case "IsHiddenStartArcShapeV2Grid":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenStartArcShapeV2Grid)
		res.valueBool = plantdiagram.IsHiddenStartArcShapeV2Grid
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenTopStartArcShapeV2Grid":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenTopStartArcShapeV2Grid)
		res.valueBool = plantdiagram.IsHiddenTopStartArcShapeV2Grid
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenEndArcShapeGrid":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenEndArcShapeGrid)
		res.valueBool = plantdiagram.IsHiddenEndArcShapeGrid
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenEndArcShapeV2Grid":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenEndArcShapeV2Grid)
		res.valueBool = plantdiagram.IsHiddenEndArcShapeV2Grid
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenTopEndArcShapeV2Grid":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenTopEndArcShapeV2Grid)
		res.valueBool = plantdiagram.IsHiddenTopEndArcShapeV2Grid
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenBottomStartArcShapeV2Grid":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenBottomStartArcShapeV2Grid)
		res.valueBool = plantdiagram.IsHiddenBottomStartArcShapeV2Grid
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenBottomEndArcShapeV2Grid":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenBottomEndArcShapeV2Grid)
		res.valueBool = plantdiagram.IsHiddenBottomEndArcShapeV2Grid
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenGrowthCurveBezierShapeGrid":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenGrowthCurveBezierShapeGrid)
		res.valueBool = plantdiagram.IsHiddenGrowthCurveBezierShapeGrid
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenStackOfGrowthCurve":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenStackOfGrowthCurve)
		res.valueBool = plantdiagram.IsHiddenStackOfGrowthCurve
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenStackOfGrowthCurveV2":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenStackOfGrowthCurveV2)
		res.valueBool = plantdiagram.IsHiddenStackOfGrowthCurveV2
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenTopStackOfGrowthCurveV2":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenTopStackOfGrowthCurveV2)
		res.valueBool = plantdiagram.IsHiddenTopStackOfGrowthCurveV2
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsHiddenBottomStackOfGrowthCurveV2":
		res.valueString = fmt.Sprintf("%t", plantdiagram.IsHiddenBottomStackOfGrowthCurveV2)
		res.valueBool = plantdiagram.IsHiddenBottomStackOfGrowthCurveV2
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

func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = stackgrowthcurveendarcshapev2.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurveendarcshapev2.StartX)
		res.valueFloat = stackgrowthcurveendarcshapev2.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurveendarcshapev2.StartY)
		res.valueFloat = stackgrowthcurveendarcshapev2.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurveendarcshapev2.EndX)
		res.valueFloat = stackgrowthcurveendarcshapev2.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurveendarcshapev2.EndY)
		res.valueFloat = stackgrowthcurveendarcshapev2.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurveendarcshapev2.XAxisRotation)
		res.valueFloat = stackgrowthcurveendarcshapev2.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", stackgrowthcurveendarcshapev2.LargeArcFlag)
		res.valueBool = stackgrowthcurveendarcshapev2.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", stackgrowthcurveendarcshapev2.SweepFlag)
		res.valueBool = stackgrowthcurveendarcshapev2.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurveendarcshapev2.RadiusX)
		res.valueFloat = stackgrowthcurveendarcshapev2.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurveendarcshapev2.RadiusY)
		res.valueFloat = stackgrowthcurveendarcshapev2.RadiusY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = stackgrowthcurvestartarcshapev2.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvestartarcshapev2.StartX)
		res.valueFloat = stackgrowthcurvestartarcshapev2.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvestartarcshapev2.StartY)
		res.valueFloat = stackgrowthcurvestartarcshapev2.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvestartarcshapev2.EndX)
		res.valueFloat = stackgrowthcurvestartarcshapev2.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvestartarcshapev2.EndY)
		res.valueFloat = stackgrowthcurvestartarcshapev2.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvestartarcshapev2.XAxisRotation)
		res.valueFloat = stackgrowthcurvestartarcshapev2.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", stackgrowthcurvestartarcshapev2.LargeArcFlag)
		res.valueBool = stackgrowthcurvestartarcshapev2.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", stackgrowthcurvestartarcshapev2.SweepFlag)
		res.valueBool = stackgrowthcurvestartarcshapev2.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvestartarcshapev2.RadiusX)
		res.valueFloat = stackgrowthcurvestartarcshapev2.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", stackgrowthcurvestartarcshapev2.RadiusY)
		res.valueFloat = stackgrowthcurvestartarcshapev2.RadiusY
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

func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = stackofgrowthcurvev2.Name
	case "StackGrowthCurveStartArcShapeV2s":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "StackGrowthCurveEndArcShapeV2s":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s {
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

func (startarcshapev2 *StartArcShapeV2) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = startarcshapev2.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", startarcshapev2.StartX)
		res.valueFloat = startarcshapev2.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", startarcshapev2.StartY)
		res.valueFloat = startarcshapev2.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", startarcshapev2.EndX)
		res.valueFloat = startarcshapev2.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", startarcshapev2.EndY)
		res.valueFloat = startarcshapev2.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", startarcshapev2.XAxisRotation)
		res.valueFloat = startarcshapev2.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", startarcshapev2.LargeArcFlag)
		res.valueBool = startarcshapev2.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", startarcshapev2.SweepFlag)
		res.valueBool = startarcshapev2.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", startarcshapev2.RadiusX)
		res.valueFloat = startarcshapev2.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", startarcshapev2.RadiusY)
		res.valueFloat = startarcshapev2.RadiusY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (startarcshapev2grid *StartArcShapeV2Grid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = startarcshapev2grid.Name
	case "StartArcShapesV2":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range startarcshapev2grid.StartArcShapesV2 {
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

func (topendarcshapev2 *TopEndArcShapeV2) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = topendarcshapev2.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", topendarcshapev2.StartX)
		res.valueFloat = topendarcshapev2.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", topendarcshapev2.StartY)
		res.valueFloat = topendarcshapev2.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", topendarcshapev2.EndX)
		res.valueFloat = topendarcshapev2.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", topendarcshapev2.EndY)
		res.valueFloat = topendarcshapev2.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", topendarcshapev2.XAxisRotation)
		res.valueFloat = topendarcshapev2.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", topendarcshapev2.LargeArcFlag)
		res.valueBool = topendarcshapev2.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", topendarcshapev2.SweepFlag)
		res.valueBool = topendarcshapev2.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", topendarcshapev2.RadiusX)
		res.valueFloat = topendarcshapev2.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", topendarcshapev2.RadiusY)
		res.valueFloat = topendarcshapev2.RadiusY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (topendarcshapev2grid *TopEndArcShapeV2Grid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = topendarcshapev2grid.Name
	case "TopEndArcShapesV2":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range topendarcshapev2grid.TopEndArcShapesV2 {
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
	case "TopStartArcShapeV2Grid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if topgrowthcurve2d.TopStartArcShapeV2Grid != nil {
			res.valueString = topgrowthcurve2d.TopStartArcShapeV2Grid.Name
			res.ids = topgrowthcurve2d.TopStartArcShapeV2Grid.GongGetUUID(stage)
		}
	case "TopEndArcShapeV2Grid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if topgrowthcurve2d.TopEndArcShapeV2Grid != nil {
			res.valueString = topgrowthcurve2d.TopEndArcShapeV2Grid.Name
			res.ids = topgrowthcurve2d.TopEndArcShapeV2Grid.GongGetUUID(stage)
		}
	}
	return
}

func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = topstackgrowthcurveendarcshapev2.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurveendarcshapev2.StartX)
		res.valueFloat = topstackgrowthcurveendarcshapev2.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurveendarcshapev2.StartY)
		res.valueFloat = topstackgrowthcurveendarcshapev2.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurveendarcshapev2.EndX)
		res.valueFloat = topstackgrowthcurveendarcshapev2.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurveendarcshapev2.EndY)
		res.valueFloat = topstackgrowthcurveendarcshapev2.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurveendarcshapev2.XAxisRotation)
		res.valueFloat = topstackgrowthcurveendarcshapev2.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", topstackgrowthcurveendarcshapev2.LargeArcFlag)
		res.valueBool = topstackgrowthcurveendarcshapev2.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", topstackgrowthcurveendarcshapev2.SweepFlag)
		res.valueBool = topstackgrowthcurveendarcshapev2.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurveendarcshapev2.RadiusX)
		res.valueFloat = topstackgrowthcurveendarcshapev2.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurveendarcshapev2.RadiusY)
		res.valueFloat = topstackgrowthcurveendarcshapev2.RadiusY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = topstackgrowthcurvestartarcshapev2.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurvestartarcshapev2.StartX)
		res.valueFloat = topstackgrowthcurvestartarcshapev2.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurvestartarcshapev2.StartY)
		res.valueFloat = topstackgrowthcurvestartarcshapev2.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurvestartarcshapev2.EndX)
		res.valueFloat = topstackgrowthcurvestartarcshapev2.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurvestartarcshapev2.EndY)
		res.valueFloat = topstackgrowthcurvestartarcshapev2.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurvestartarcshapev2.XAxisRotation)
		res.valueFloat = topstackgrowthcurvestartarcshapev2.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", topstackgrowthcurvestartarcshapev2.LargeArcFlag)
		res.valueBool = topstackgrowthcurvestartarcshapev2.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", topstackgrowthcurvestartarcshapev2.SweepFlag)
		res.valueBool = topstackgrowthcurvestartarcshapev2.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurvestartarcshapev2.RadiusX)
		res.valueFloat = topstackgrowthcurvestartarcshapev2.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", topstackgrowthcurvestartarcshapev2.RadiusY)
		res.valueFloat = topstackgrowthcurvestartarcshapev2.RadiusY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = topstackofgrowthcurvev2.Name
	case "TopStackGrowthCurveStartArcShapeV2s":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "TopStackGrowthCurveEndArcShapeV2s":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s {
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

func (topstartarcshapev2 *TopStartArcShapeV2) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = topstartarcshapev2.Name
	case "StartX":
		res.valueString = fmt.Sprintf("%f", topstartarcshapev2.StartX)
		res.valueFloat = topstartarcshapev2.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", topstartarcshapev2.StartY)
		res.valueFloat = topstartarcshapev2.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", topstartarcshapev2.EndX)
		res.valueFloat = topstartarcshapev2.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", topstartarcshapev2.EndY)
		res.valueFloat = topstartarcshapev2.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XAxisRotation":
		res.valueString = fmt.Sprintf("%f", topstartarcshapev2.XAxisRotation)
		res.valueFloat = topstartarcshapev2.XAxisRotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LargeArcFlag":
		res.valueString = fmt.Sprintf("%t", topstartarcshapev2.LargeArcFlag)
		res.valueBool = topstartarcshapev2.LargeArcFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SweepFlag":
		res.valueString = fmt.Sprintf("%t", topstartarcshapev2.SweepFlag)
		res.valueBool = topstartarcshapev2.SweepFlag
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RadiusX":
		res.valueString = fmt.Sprintf("%f", topstartarcshapev2.RadiusX)
		res.valueFloat = topstartarcshapev2.RadiusX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusY":
		res.valueString = fmt.Sprintf("%f", topstartarcshapev2.RadiusY)
		res.valueFloat = topstartarcshapev2.RadiusY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (topstartarcshapev2grid *TopStartArcShapeV2Grid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = topstartarcshapev2grid.Name
	case "TopStartArcShapesV2":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range topstartarcshapev2grid.TopStartArcShapesV2 {
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

func (bottomendarcshapev2 *BottomEndArcShapeV2) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		bottomendarcshapev2.Name = value.GetValueString()
	case "StartX":
		bottomendarcshapev2.StartX = value.GetValueFloat()
	case "StartY":
		bottomendarcshapev2.StartY = value.GetValueFloat()
	case "EndX":
		bottomendarcshapev2.EndX = value.GetValueFloat()
	case "EndY":
		bottomendarcshapev2.EndY = value.GetValueFloat()
	case "XAxisRotation":
		bottomendarcshapev2.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		bottomendarcshapev2.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		bottomendarcshapev2.SweepFlag = value.GetValueBool()
	case "RadiusX":
		bottomendarcshapev2.RadiusX = value.GetValueFloat()
	case "RadiusY":
		bottomendarcshapev2.RadiusY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		bottomendarcshapev2grid.Name = value.GetValueString()
	case "BottomEndArcShapesV2":
		bottomendarcshapev2grid.BottomEndArcShapesV2 = make([]*BottomEndArcShapeV2, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.BottomEndArcShapeV2s {
					if stage.BottomEndArcShapeV2_stagedOrder[__instance__] == uint(id) {
						bottomendarcshapev2grid.BottomEndArcShapesV2 = append(bottomendarcshapev2grid.BottomEndArcShapesV2, __instance__)
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

func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		bottomstackgrowthcurveendarcshapev2.Name = value.GetValueString()
	case "StartX":
		bottomstackgrowthcurveendarcshapev2.StartX = value.GetValueFloat()
	case "StartY":
		bottomstackgrowthcurveendarcshapev2.StartY = value.GetValueFloat()
	case "EndX":
		bottomstackgrowthcurveendarcshapev2.EndX = value.GetValueFloat()
	case "EndY":
		bottomstackgrowthcurveendarcshapev2.EndY = value.GetValueFloat()
	case "XAxisRotation":
		bottomstackgrowthcurveendarcshapev2.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		bottomstackgrowthcurveendarcshapev2.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		bottomstackgrowthcurveendarcshapev2.SweepFlag = value.GetValueBool()
	case "RadiusX":
		bottomstackgrowthcurveendarcshapev2.RadiusX = value.GetValueFloat()
	case "RadiusY":
		bottomstackgrowthcurveendarcshapev2.RadiusY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		bottomstackgrowthcurvestartarcshapev2.Name = value.GetValueString()
	case "StartX":
		bottomstackgrowthcurvestartarcshapev2.StartX = value.GetValueFloat()
	case "StartY":
		bottomstackgrowthcurvestartarcshapev2.StartY = value.GetValueFloat()
	case "EndX":
		bottomstackgrowthcurvestartarcshapev2.EndX = value.GetValueFloat()
	case "EndY":
		bottomstackgrowthcurvestartarcshapev2.EndY = value.GetValueFloat()
	case "XAxisRotation":
		bottomstackgrowthcurvestartarcshapev2.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		bottomstackgrowthcurvestartarcshapev2.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		bottomstackgrowthcurvestartarcshapev2.SweepFlag = value.GetValueBool()
	case "RadiusX":
		bottomstackgrowthcurvestartarcshapev2.RadiusX = value.GetValueFloat()
	case "RadiusY":
		bottomstackgrowthcurvestartarcshapev2.RadiusY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		bottomstackofgrowthcurvev2.Name = value.GetValueString()
	case "BottomStackGrowthCurveStartArcShapeV2s":
		bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s = make([]*BottomStackGrowthCurveStartArcShapeV2, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.BottomStackGrowthCurveStartArcShapeV2s {
					if stage.BottomStackGrowthCurveStartArcShapeV2_stagedOrder[__instance__] == uint(id) {
						bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s = append(bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s, __instance__)
						break
					}
				}
			}
		}
	case "BottomStackGrowthCurveEndArcShapeV2s":
		bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s = make([]*BottomStackGrowthCurveEndArcShapeV2, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.BottomStackGrowthCurveEndArcShapeV2s {
					if stage.BottomStackGrowthCurveEndArcShapeV2_stagedOrder[__instance__] == uint(id) {
						bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s = append(bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s, __instance__)
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

func (bottomstartarcshapev2 *BottomStartArcShapeV2) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		bottomstartarcshapev2.Name = value.GetValueString()
	case "StartX":
		bottomstartarcshapev2.StartX = value.GetValueFloat()
	case "StartY":
		bottomstartarcshapev2.StartY = value.GetValueFloat()
	case "EndX":
		bottomstartarcshapev2.EndX = value.GetValueFloat()
	case "EndY":
		bottomstartarcshapev2.EndY = value.GetValueFloat()
	case "XAxisRotation":
		bottomstartarcshapev2.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		bottomstartarcshapev2.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		bottomstartarcshapev2.SweepFlag = value.GetValueBool()
	case "RadiusX":
		bottomstartarcshapev2.RadiusX = value.GetValueFloat()
	case "RadiusY":
		bottomstartarcshapev2.RadiusY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		bottomstartarcshapev2grid.Name = value.GetValueString()
	case "BottomStartArcShapesV2":
		bottomstartarcshapev2grid.BottomStartArcShapesV2 = make([]*BottomStartArcShapeV2, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.BottomStartArcShapeV2s {
					if stage.BottomStartArcShapeV2_stagedOrder[__instance__] == uint(id) {
						bottomstartarcshapev2grid.BottomStartArcShapesV2 = append(bottomstartarcshapev2grid.BottomStartArcShapesV2, __instance__)
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

func (endarcshapev2 *EndArcShapeV2) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		endarcshapev2.Name = value.GetValueString()
	case "StartX":
		endarcshapev2.StartX = value.GetValueFloat()
	case "StartY":
		endarcshapev2.StartY = value.GetValueFloat()
	case "EndX":
		endarcshapev2.EndX = value.GetValueFloat()
	case "EndY":
		endarcshapev2.EndY = value.GetValueFloat()
	case "XAxisRotation":
		endarcshapev2.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		endarcshapev2.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		endarcshapev2.SweepFlag = value.GetValueBool()
	case "RadiusX":
		endarcshapev2.RadiusX = value.GetValueFloat()
	case "RadiusY":
		endarcshapev2.RadiusY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (endarcshapev2grid *EndArcShapeV2Grid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		endarcshapev2grid.Name = value.GetValueString()
	case "EndArcShapesV2":
		endarcshapev2grid.EndArcShapesV2 = make([]*EndArcShapeV2, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.EndArcShapeV2s {
					if stage.EndArcShapeV2_stagedOrder[__instance__] == uint(id) {
						endarcshapev2grid.EndArcShapesV2 = append(endarcshapev2grid.EndArcShapesV2, __instance__)
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
	case "StartArcShapeV2Grid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			growthcurve2d.StartArcShapeV2Grid = nil
			for __instance__ := range stage.StartArcShapeV2Grids {
				if stage.StartArcShapeV2Grid_stagedOrder[__instance__] == uint(id) {
					growthcurve2d.StartArcShapeV2Grid = __instance__
					break
				}
			}
		}
	case "EndArcShapeV2Grid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			growthcurve2d.EndArcShapeV2Grid = nil
			for __instance__ := range stage.EndArcShapeV2Grids {
				if stage.EndArcShapeV2Grid_stagedOrder[__instance__] == uint(id) {
					growthcurve2d.EndArcShapeV2Grid = __instance__
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
	case "Thickness":
		plant.Thickness = value.GetValueFloat()
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
	case "StartArcShapeV2Grid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.StartArcShapeV2Grid = nil
			for __instance__ := range stage.StartArcShapeV2Grids {
				if stage.StartArcShapeV2Grid_stagedOrder[__instance__] == uint(id) {
					plant.StartArcShapeV2Grid = __instance__
					break
				}
			}
		}
	case "TopStartArcShapeV2Grid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.TopStartArcShapeV2Grid = nil
			for __instance__ := range stage.TopStartArcShapeV2Grids {
				if stage.TopStartArcShapeV2Grid_stagedOrder[__instance__] == uint(id) {
					plant.TopStartArcShapeV2Grid = __instance__
					break
				}
			}
		}
	case "EndArcShapeV2Grid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.EndArcShapeV2Grid = nil
			for __instance__ := range stage.EndArcShapeV2Grids {
				if stage.EndArcShapeV2Grid_stagedOrder[__instance__] == uint(id) {
					plant.EndArcShapeV2Grid = __instance__
					break
				}
			}
		}
	case "TopEndArcShapeV2Grid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.TopEndArcShapeV2Grid = nil
			for __instance__ := range stage.TopEndArcShapeV2Grids {
				if stage.TopEndArcShapeV2Grid_stagedOrder[__instance__] == uint(id) {
					plant.TopEndArcShapeV2Grid = __instance__
					break
				}
			}
		}
	case "BottomStartArcShapeV2Grid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.BottomStartArcShapeV2Grid = nil
			for __instance__ := range stage.BottomStartArcShapeV2Grids {
				if stage.BottomStartArcShapeV2Grid_stagedOrder[__instance__] == uint(id) {
					plant.BottomStartArcShapeV2Grid = __instance__
					break
				}
			}
		}
	case "BottomEndArcShapeV2Grid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.BottomEndArcShapeV2Grid = nil
			for __instance__ := range stage.BottomEndArcShapeV2Grids {
				if stage.BottomEndArcShapeV2Grid_stagedOrder[__instance__] == uint(id) {
					plant.BottomEndArcShapeV2Grid = __instance__
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
	case "StackOfGrowthCurveV2":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.StackOfGrowthCurveV2 = nil
			for __instance__ := range stage.StackOfGrowthCurveV2s {
				if stage.StackOfGrowthCurveV2_stagedOrder[__instance__] == uint(id) {
					plant.StackOfGrowthCurveV2 = __instance__
					break
				}
			}
		}
	case "TopStackOfGrowthCurveV2":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.TopStackOfGrowthCurveV2 = nil
			for __instance__ := range stage.TopStackOfGrowthCurveV2s {
				if stage.TopStackOfGrowthCurveV2_stagedOrder[__instance__] == uint(id) {
					plant.TopStackOfGrowthCurveV2 = __instance__
					break
				}
			}
		}
	case "BottomStackOfGrowthCurveV2":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			plant.BottomStackOfGrowthCurveV2 = nil
			for __instance__ := range stage.BottomStackOfGrowthCurveV2s {
				if stage.BottomStackOfGrowthCurveV2_stagedOrder[__instance__] == uint(id) {
					plant.BottomStackOfGrowthCurveV2 = __instance__
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
	case "IsHiddenStartArcShapeV2Grid":
		plantdiagram.IsHiddenStartArcShapeV2Grid = value.GetValueBool()
	case "IsHiddenTopStartArcShapeV2Grid":
		plantdiagram.IsHiddenTopStartArcShapeV2Grid = value.GetValueBool()
	case "IsHiddenEndArcShapeGrid":
		plantdiagram.IsHiddenEndArcShapeGrid = value.GetValueBool()
	case "IsHiddenEndArcShapeV2Grid":
		plantdiagram.IsHiddenEndArcShapeV2Grid = value.GetValueBool()
	case "IsHiddenTopEndArcShapeV2Grid":
		plantdiagram.IsHiddenTopEndArcShapeV2Grid = value.GetValueBool()
	case "IsHiddenBottomStartArcShapeV2Grid":
		plantdiagram.IsHiddenBottomStartArcShapeV2Grid = value.GetValueBool()
	case "IsHiddenBottomEndArcShapeV2Grid":
		plantdiagram.IsHiddenBottomEndArcShapeV2Grid = value.GetValueBool()
	case "IsHiddenGrowthCurveBezierShapeGrid":
		plantdiagram.IsHiddenGrowthCurveBezierShapeGrid = value.GetValueBool()
	case "IsHiddenStackOfGrowthCurve":
		plantdiagram.IsHiddenStackOfGrowthCurve = value.GetValueBool()
	case "IsHiddenStackOfGrowthCurveV2":
		plantdiagram.IsHiddenStackOfGrowthCurveV2 = value.GetValueBool()
	case "IsHiddenTopStackOfGrowthCurveV2":
		plantdiagram.IsHiddenTopStackOfGrowthCurveV2 = value.GetValueBool()
	case "IsHiddenBottomStackOfGrowthCurveV2":
		plantdiagram.IsHiddenBottomStackOfGrowthCurveV2 = value.GetValueBool()
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

func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		stackgrowthcurveendarcshapev2.Name = value.GetValueString()
	case "StartX":
		stackgrowthcurveendarcshapev2.StartX = value.GetValueFloat()
	case "StartY":
		stackgrowthcurveendarcshapev2.StartY = value.GetValueFloat()
	case "EndX":
		stackgrowthcurveendarcshapev2.EndX = value.GetValueFloat()
	case "EndY":
		stackgrowthcurveendarcshapev2.EndY = value.GetValueFloat()
	case "XAxisRotation":
		stackgrowthcurveendarcshapev2.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		stackgrowthcurveendarcshapev2.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		stackgrowthcurveendarcshapev2.SweepFlag = value.GetValueBool()
	case "RadiusX":
		stackgrowthcurveendarcshapev2.RadiusX = value.GetValueFloat()
	case "RadiusY":
		stackgrowthcurveendarcshapev2.RadiusY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		stackgrowthcurvestartarcshapev2.Name = value.GetValueString()
	case "StartX":
		stackgrowthcurvestartarcshapev2.StartX = value.GetValueFloat()
	case "StartY":
		stackgrowthcurvestartarcshapev2.StartY = value.GetValueFloat()
	case "EndX":
		stackgrowthcurvestartarcshapev2.EndX = value.GetValueFloat()
	case "EndY":
		stackgrowthcurvestartarcshapev2.EndY = value.GetValueFloat()
	case "XAxisRotation":
		stackgrowthcurvestartarcshapev2.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		stackgrowthcurvestartarcshapev2.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		stackgrowthcurvestartarcshapev2.SweepFlag = value.GetValueBool()
	case "RadiusX":
		stackgrowthcurvestartarcshapev2.RadiusX = value.GetValueFloat()
	case "RadiusY":
		stackgrowthcurvestartarcshapev2.RadiusY = value.GetValueFloat()
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

func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		stackofgrowthcurvev2.Name = value.GetValueString()
	case "StackGrowthCurveStartArcShapeV2s":
		stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s = make([]*StackGrowthCurveStartArcShapeV2, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.StackGrowthCurveStartArcShapeV2s {
					if stage.StackGrowthCurveStartArcShapeV2_stagedOrder[__instance__] == uint(id) {
						stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s = append(stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s, __instance__)
						break
					}
				}
			}
		}
	case "StackGrowthCurveEndArcShapeV2s":
		stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s = make([]*StackGrowthCurveEndArcShapeV2, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.StackGrowthCurveEndArcShapeV2s {
					if stage.StackGrowthCurveEndArcShapeV2_stagedOrder[__instance__] == uint(id) {
						stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s = append(stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s, __instance__)
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

func (startarcshapev2 *StartArcShapeV2) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		startarcshapev2.Name = value.GetValueString()
	case "StartX":
		startarcshapev2.StartX = value.GetValueFloat()
	case "StartY":
		startarcshapev2.StartY = value.GetValueFloat()
	case "EndX":
		startarcshapev2.EndX = value.GetValueFloat()
	case "EndY":
		startarcshapev2.EndY = value.GetValueFloat()
	case "XAxisRotation":
		startarcshapev2.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		startarcshapev2.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		startarcshapev2.SweepFlag = value.GetValueBool()
	case "RadiusX":
		startarcshapev2.RadiusX = value.GetValueFloat()
	case "RadiusY":
		startarcshapev2.RadiusY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (startarcshapev2grid *StartArcShapeV2Grid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		startarcshapev2grid.Name = value.GetValueString()
	case "StartArcShapesV2":
		startarcshapev2grid.StartArcShapesV2 = make([]*StartArcShapeV2, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.StartArcShapeV2s {
					if stage.StartArcShapeV2_stagedOrder[__instance__] == uint(id) {
						startarcshapev2grid.StartArcShapesV2 = append(startarcshapev2grid.StartArcShapesV2, __instance__)
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

func (topendarcshapev2 *TopEndArcShapeV2) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		topendarcshapev2.Name = value.GetValueString()
	case "StartX":
		topendarcshapev2.StartX = value.GetValueFloat()
	case "StartY":
		topendarcshapev2.StartY = value.GetValueFloat()
	case "EndX":
		topendarcshapev2.EndX = value.GetValueFloat()
	case "EndY":
		topendarcshapev2.EndY = value.GetValueFloat()
	case "XAxisRotation":
		topendarcshapev2.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		topendarcshapev2.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		topendarcshapev2.SweepFlag = value.GetValueBool()
	case "RadiusX":
		topendarcshapev2.RadiusX = value.GetValueFloat()
	case "RadiusY":
		topendarcshapev2.RadiusY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (topendarcshapev2grid *TopEndArcShapeV2Grid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		topendarcshapev2grid.Name = value.GetValueString()
	case "TopEndArcShapesV2":
		topendarcshapev2grid.TopEndArcShapesV2 = make([]*TopEndArcShapeV2, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TopEndArcShapeV2s {
					if stage.TopEndArcShapeV2_stagedOrder[__instance__] == uint(id) {
						topendarcshapev2grid.TopEndArcShapesV2 = append(topendarcshapev2grid.TopEndArcShapesV2, __instance__)
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
	case "TopStartArcShapeV2Grid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			topgrowthcurve2d.TopStartArcShapeV2Grid = nil
			for __instance__ := range stage.TopStartArcShapeV2Grids {
				if stage.TopStartArcShapeV2Grid_stagedOrder[__instance__] == uint(id) {
					topgrowthcurve2d.TopStartArcShapeV2Grid = __instance__
					break
				}
			}
		}
	case "TopEndArcShapeV2Grid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			topgrowthcurve2d.TopEndArcShapeV2Grid = nil
			for __instance__ := range stage.TopEndArcShapeV2Grids {
				if stage.TopEndArcShapeV2Grid_stagedOrder[__instance__] == uint(id) {
					topgrowthcurve2d.TopEndArcShapeV2Grid = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		topstackgrowthcurveendarcshapev2.Name = value.GetValueString()
	case "StartX":
		topstackgrowthcurveendarcshapev2.StartX = value.GetValueFloat()
	case "StartY":
		topstackgrowthcurveendarcshapev2.StartY = value.GetValueFloat()
	case "EndX":
		topstackgrowthcurveendarcshapev2.EndX = value.GetValueFloat()
	case "EndY":
		topstackgrowthcurveendarcshapev2.EndY = value.GetValueFloat()
	case "XAxisRotation":
		topstackgrowthcurveendarcshapev2.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		topstackgrowthcurveendarcshapev2.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		topstackgrowthcurveendarcshapev2.SweepFlag = value.GetValueBool()
	case "RadiusX":
		topstackgrowthcurveendarcshapev2.RadiusX = value.GetValueFloat()
	case "RadiusY":
		topstackgrowthcurveendarcshapev2.RadiusY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		topstackgrowthcurvestartarcshapev2.Name = value.GetValueString()
	case "StartX":
		topstackgrowthcurvestartarcshapev2.StartX = value.GetValueFloat()
	case "StartY":
		topstackgrowthcurvestartarcshapev2.StartY = value.GetValueFloat()
	case "EndX":
		topstackgrowthcurvestartarcshapev2.EndX = value.GetValueFloat()
	case "EndY":
		topstackgrowthcurvestartarcshapev2.EndY = value.GetValueFloat()
	case "XAxisRotation":
		topstackgrowthcurvestartarcshapev2.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		topstackgrowthcurvestartarcshapev2.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		topstackgrowthcurvestartarcshapev2.SweepFlag = value.GetValueBool()
	case "RadiusX":
		topstackgrowthcurvestartarcshapev2.RadiusX = value.GetValueFloat()
	case "RadiusY":
		topstackgrowthcurvestartarcshapev2.RadiusY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		topstackofgrowthcurvev2.Name = value.GetValueString()
	case "TopStackGrowthCurveStartArcShapeV2s":
		topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s = make([]*TopStackGrowthCurveStartArcShapeV2, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TopStackGrowthCurveStartArcShapeV2s {
					if stage.TopStackGrowthCurveStartArcShapeV2_stagedOrder[__instance__] == uint(id) {
						topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s = append(topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s, __instance__)
						break
					}
				}
			}
		}
	case "TopStackGrowthCurveEndArcShapeV2s":
		topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s = make([]*TopStackGrowthCurveEndArcShapeV2, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TopStackGrowthCurveEndArcShapeV2s {
					if stage.TopStackGrowthCurveEndArcShapeV2_stagedOrder[__instance__] == uint(id) {
						topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s = append(topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s, __instance__)
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

func (topstartarcshapev2 *TopStartArcShapeV2) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		topstartarcshapev2.Name = value.GetValueString()
	case "StartX":
		topstartarcshapev2.StartX = value.GetValueFloat()
	case "StartY":
		topstartarcshapev2.StartY = value.GetValueFloat()
	case "EndX":
		topstartarcshapev2.EndX = value.GetValueFloat()
	case "EndY":
		topstartarcshapev2.EndY = value.GetValueFloat()
	case "XAxisRotation":
		topstartarcshapev2.XAxisRotation = value.GetValueFloat()
	case "LargeArcFlag":
		topstartarcshapev2.LargeArcFlag = value.GetValueBool()
	case "SweepFlag":
		topstartarcshapev2.SweepFlag = value.GetValueBool()
	case "RadiusX":
		topstartarcshapev2.RadiusX = value.GetValueFloat()
	case "RadiusY":
		topstartarcshapev2.RadiusY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (topstartarcshapev2grid *TopStartArcShapeV2Grid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		topstartarcshapev2grid.Name = value.GetValueString()
	case "TopStartArcShapesV2":
		topstartarcshapev2grid.TopStartArcShapesV2 = make([]*TopStartArcShapeV2, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TopStartArcShapeV2s {
					if stage.TopStartArcShapeV2_stagedOrder[__instance__] == uint(id) {
						topstartarcshapev2grid.TopStartArcShapesV2 = append(topstartarcshapev2grid.TopStartArcShapesV2, __instance__)
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

func (bottomendarcshapev2 *BottomEndArcShapeV2) GongGetGongstructName() string {
	return "BottomEndArcShapeV2"
}

func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) GongGetGongstructName() string {
	return "BottomEndArcShapeV2Grid"
}

func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) GongGetGongstructName() string {
	return "BottomStackGrowthCurveEndArcShapeV2"
}

func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) GongGetGongstructName() string {
	return "BottomStackGrowthCurveStartArcShapeV2"
}

func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) GongGetGongstructName() string {
	return "BottomStackOfGrowthCurveV2"
}

func (bottomstartarcshapev2 *BottomStartArcShapeV2) GongGetGongstructName() string {
	return "BottomStartArcShapeV2"
}

func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) GongGetGongstructName() string {
	return "BottomStartArcShapeV2Grid"
}

func (circlegridshape *CircleGridShape) GongGetGongstructName() string {
	return "CircleGridShape"
}

func (endarcshapev2 *EndArcShapeV2) GongGetGongstructName() string {
	return "EndArcShapeV2"
}

func (endarcshapev2grid *EndArcShapeV2Grid) GongGetGongstructName() string {
	return "EndArcShapeV2Grid"
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

func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) GongGetGongstructName() string {
	return "StackGrowthCurveBezierShape"
}

func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) GongGetGongstructName() string {
	return "StackGrowthCurveEndArcShapeV2"
}

func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) GongGetGongstructName() string {
	return "StackGrowthCurveStartArcShapeV2"
}

func (stackofgrowthcurve *StackOfGrowthCurve) GongGetGongstructName() string {
	return "StackOfGrowthCurve"
}

func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) GongGetGongstructName() string {
	return "StackOfGrowthCurveV2"
}

func (startarcshapev2 *StartArcShapeV2) GongGetGongstructName() string {
	return "StartArcShapeV2"
}

func (startarcshapev2grid *StartArcShapeV2Grid) GongGetGongstructName() string {
	return "StartArcShapeV2Grid"
}

func (topendarcshapev2 *TopEndArcShapeV2) GongGetGongstructName() string {
	return "TopEndArcShapeV2"
}

func (topendarcshapev2grid *TopEndArcShapeV2Grid) GongGetGongstructName() string {
	return "TopEndArcShapeV2Grid"
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongGetGongstructName() string {
	return "TopGrowthCurve2D"
}

func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) GongGetGongstructName() string {
	return "TopStackGrowthCurveEndArcShapeV2"
}

func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) GongGetGongstructName() string {
	return "TopStackGrowthCurveStartArcShapeV2"
}

func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) GongGetGongstructName() string {
	return "TopStackOfGrowthCurveV2"
}

func (topstartarcshapev2 *TopStartArcShapeV2) GongGetGongstructName() string {
	return "TopStartArcShapeV2"
}

func (topstartarcshapev2grid *TopStartArcShapeV2Grid) GongGetGongstructName() string {
	return "TopStartArcShapeV2Grid"
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

	stage.BottomEndArcShapeV2s_mapString = make(map[string]*BottomEndArcShapeV2)
	for bottomendarcshapev2 := range stage.BottomEndArcShapeV2s {
		stage.BottomEndArcShapeV2s_mapString[bottomendarcshapev2.Name] = bottomendarcshapev2
	}

	stage.BottomEndArcShapeV2Grids_mapString = make(map[string]*BottomEndArcShapeV2Grid)
	for bottomendarcshapev2grid := range stage.BottomEndArcShapeV2Grids {
		stage.BottomEndArcShapeV2Grids_mapString[bottomendarcshapev2grid.Name] = bottomendarcshapev2grid
	}

	stage.BottomStackGrowthCurveEndArcShapeV2s_mapString = make(map[string]*BottomStackGrowthCurveEndArcShapeV2)
	for bottomstackgrowthcurveendarcshapev2 := range stage.BottomStackGrowthCurveEndArcShapeV2s {
		stage.BottomStackGrowthCurveEndArcShapeV2s_mapString[bottomstackgrowthcurveendarcshapev2.Name] = bottomstackgrowthcurveendarcshapev2
	}

	stage.BottomStackGrowthCurveStartArcShapeV2s_mapString = make(map[string]*BottomStackGrowthCurveStartArcShapeV2)
	for bottomstackgrowthcurvestartarcshapev2 := range stage.BottomStackGrowthCurveStartArcShapeV2s {
		stage.BottomStackGrowthCurveStartArcShapeV2s_mapString[bottomstackgrowthcurvestartarcshapev2.Name] = bottomstackgrowthcurvestartarcshapev2
	}

	stage.BottomStackOfGrowthCurveV2s_mapString = make(map[string]*BottomStackOfGrowthCurveV2)
	for bottomstackofgrowthcurvev2 := range stage.BottomStackOfGrowthCurveV2s {
		stage.BottomStackOfGrowthCurveV2s_mapString[bottomstackofgrowthcurvev2.Name] = bottomstackofgrowthcurvev2
	}

	stage.BottomStartArcShapeV2s_mapString = make(map[string]*BottomStartArcShapeV2)
	for bottomstartarcshapev2 := range stage.BottomStartArcShapeV2s {
		stage.BottomStartArcShapeV2s_mapString[bottomstartarcshapev2.Name] = bottomstartarcshapev2
	}

	stage.BottomStartArcShapeV2Grids_mapString = make(map[string]*BottomStartArcShapeV2Grid)
	for bottomstartarcshapev2grid := range stage.BottomStartArcShapeV2Grids {
		stage.BottomStartArcShapeV2Grids_mapString[bottomstartarcshapev2grid.Name] = bottomstartarcshapev2grid
	}

	stage.CircleGridShapes_mapString = make(map[string]*CircleGridShape)
	for circlegridshape := range stage.CircleGridShapes {
		stage.CircleGridShapes_mapString[circlegridshape.Name] = circlegridshape
	}

	stage.EndArcShapeV2s_mapString = make(map[string]*EndArcShapeV2)
	for endarcshapev2 := range stage.EndArcShapeV2s {
		stage.EndArcShapeV2s_mapString[endarcshapev2.Name] = endarcshapev2
	}

	stage.EndArcShapeV2Grids_mapString = make(map[string]*EndArcShapeV2Grid)
	for endarcshapev2grid := range stage.EndArcShapeV2Grids {
		stage.EndArcShapeV2Grids_mapString[endarcshapev2grid.Name] = endarcshapev2grid
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

	stage.StackGrowthCurveBezierShapes_mapString = make(map[string]*StackGrowthCurveBezierShape)
	for stackgrowthcurvebeziershape := range stage.StackGrowthCurveBezierShapes {
		stage.StackGrowthCurveBezierShapes_mapString[stackgrowthcurvebeziershape.Name] = stackgrowthcurvebeziershape
	}

	stage.StackGrowthCurveEndArcShapeV2s_mapString = make(map[string]*StackGrowthCurveEndArcShapeV2)
	for stackgrowthcurveendarcshapev2 := range stage.StackGrowthCurveEndArcShapeV2s {
		stage.StackGrowthCurveEndArcShapeV2s_mapString[stackgrowthcurveendarcshapev2.Name] = stackgrowthcurveendarcshapev2
	}

	stage.StackGrowthCurveStartArcShapeV2s_mapString = make(map[string]*StackGrowthCurveStartArcShapeV2)
	for stackgrowthcurvestartarcshapev2 := range stage.StackGrowthCurveStartArcShapeV2s {
		stage.StackGrowthCurveStartArcShapeV2s_mapString[stackgrowthcurvestartarcshapev2.Name] = stackgrowthcurvestartarcshapev2
	}

	stage.StackOfGrowthCurves_mapString = make(map[string]*StackOfGrowthCurve)
	for stackofgrowthcurve := range stage.StackOfGrowthCurves {
		stage.StackOfGrowthCurves_mapString[stackofgrowthcurve.Name] = stackofgrowthcurve
	}

	stage.StackOfGrowthCurveV2s_mapString = make(map[string]*StackOfGrowthCurveV2)
	for stackofgrowthcurvev2 := range stage.StackOfGrowthCurveV2s {
		stage.StackOfGrowthCurveV2s_mapString[stackofgrowthcurvev2.Name] = stackofgrowthcurvev2
	}

	stage.StartArcShapeV2s_mapString = make(map[string]*StartArcShapeV2)
	for startarcshapev2 := range stage.StartArcShapeV2s {
		stage.StartArcShapeV2s_mapString[startarcshapev2.Name] = startarcshapev2
	}

	stage.StartArcShapeV2Grids_mapString = make(map[string]*StartArcShapeV2Grid)
	for startarcshapev2grid := range stage.StartArcShapeV2Grids {
		stage.StartArcShapeV2Grids_mapString[startarcshapev2grid.Name] = startarcshapev2grid
	}

	stage.TopEndArcShapeV2s_mapString = make(map[string]*TopEndArcShapeV2)
	for topendarcshapev2 := range stage.TopEndArcShapeV2s {
		stage.TopEndArcShapeV2s_mapString[topendarcshapev2.Name] = topendarcshapev2
	}

	stage.TopEndArcShapeV2Grids_mapString = make(map[string]*TopEndArcShapeV2Grid)
	for topendarcshapev2grid := range stage.TopEndArcShapeV2Grids {
		stage.TopEndArcShapeV2Grids_mapString[topendarcshapev2grid.Name] = topendarcshapev2grid
	}

	stage.TopGrowthCurve2Ds_mapString = make(map[string]*TopGrowthCurve2D)
	for topgrowthcurve2d := range stage.TopGrowthCurve2Ds {
		stage.TopGrowthCurve2Ds_mapString[topgrowthcurve2d.Name] = topgrowthcurve2d
	}

	stage.TopStackGrowthCurveEndArcShapeV2s_mapString = make(map[string]*TopStackGrowthCurveEndArcShapeV2)
	for topstackgrowthcurveendarcshapev2 := range stage.TopStackGrowthCurveEndArcShapeV2s {
		stage.TopStackGrowthCurveEndArcShapeV2s_mapString[topstackgrowthcurveendarcshapev2.Name] = topstackgrowthcurveendarcshapev2
	}

	stage.TopStackGrowthCurveStartArcShapeV2s_mapString = make(map[string]*TopStackGrowthCurveStartArcShapeV2)
	for topstackgrowthcurvestartarcshapev2 := range stage.TopStackGrowthCurveStartArcShapeV2s {
		stage.TopStackGrowthCurveStartArcShapeV2s_mapString[topstackgrowthcurvestartarcshapev2.Name] = topstackgrowthcurvestartarcshapev2
	}

	stage.TopStackOfGrowthCurveV2s_mapString = make(map[string]*TopStackOfGrowthCurveV2)
	for topstackofgrowthcurvev2 := range stage.TopStackOfGrowthCurveV2s {
		stage.TopStackOfGrowthCurveV2s_mapString[topstackofgrowthcurvev2.Name] = topstackofgrowthcurvev2
	}

	stage.TopStartArcShapeV2s_mapString = make(map[string]*TopStartArcShapeV2)
	for topstartarcshapev2 := range stage.TopStartArcShapeV2s {
		stage.TopStartArcShapeV2s_mapString[topstartarcshapev2.Name] = topstartarcshapev2
	}

	stage.TopStartArcShapeV2Grids_mapString = make(map[string]*TopStartArcShapeV2Grid)
	for topstartarcshapev2grid := range stage.TopStartArcShapeV2Grids {
		stage.TopStartArcShapeV2Grids_mapString[topstartarcshapev2grid.Name] = topstartarcshapev2grid
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
