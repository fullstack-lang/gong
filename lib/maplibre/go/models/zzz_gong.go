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
	"time"

	gong_runtime "github.com/fullstack-lang/gong/pkg/runtime"
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

type GongMarshallingMode = gong_runtime.GongMarshallingMode

const (
	// the whole stage is generated at each marshall. This is the default
	GongMarshallingNormal = gong_runtime.GongMarshallingNormal

	// only the last commit is append to the marshall file
	GongMarshallingAppendCommit = gong_runtime.GongMarshallingAppendCommit
)

type gongStageNavigationMode = gong_runtime.StageNavigationMode

const (
	GongNavigationModeNormal     = gong_runtime.GongNavigationModeNormal
	GongNavigationModeNavigating = gong_runtime.GongNavigationModeNavigating
)

const (
	GongProbeTreeSidebarSuffix           = gong_runtime.GongProbeTreeSidebarSuffix
	GongProbeNavigationTreeSidebarSuffix = gong_runtime.GongProbeNavigationTreeSidebarSuffix
	GongProbeTableSuffix                 = gong_runtime.GongProbeTableSuffix
	GongProbeNotificationTableSuffix     = gong_runtime.GongProbeNotificationTableSuffix
	GongProbeFormSuffix                  = gong_runtime.GongProbeFormSuffix
	GongProbeSplitSuffix                 = gong_runtime.GongProbeSplitSuffix
	GongProbeLoadSuffix                  = gong_runtime.GongProbeLoadSuffix
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
	gong_runtime.StageCore

	// insertion point for definition of arrays registering instances
	Coordinates                map[*Coordinate]struct{}
	Coordinates_instance       map[*Coordinate]*Coordinate
	Coordinates_mapString      map[string]*Coordinate
	CoordinateOrder            uint
	Coordinate_stagedOrder     map[*Coordinate]uint
	Coordinate_orderStaged     map[uint]*Coordinate
	Coordinates_reference      map[*Coordinate]*Coordinate
	Coordinates_referenceOrder map[*Coordinate]uint

	// insertion point for slice of pointers maps
	OnAfterCoordinateCreateCallback GongOnAfterCreateInterface[Coordinate]
	OnAfterCoordinateUpdateCallback GongOnAfterUpdateInterface[Coordinate]
	OnAfterCoordinateDeleteCallback GongOnAfterDeleteInterface[Coordinate]

	Layers                map[*Layer]struct{}
	Layers_instance       map[*Layer]*Layer
	Layers_mapString      map[string]*Layer
	LayerOrder            uint
	Layer_stagedOrder     map[*Layer]uint
	Layer_orderStaged     map[uint]*Layer
	Layers_reference      map[*Layer]*Layer
	Layers_referenceOrder map[*Layer]uint

	// insertion point for slice of pointers maps
	Layer_Polygons_reverseMap map[*Polygon]*Layer

	Layer_Polylines_reverseMap map[*Polyline]*Layer

	Layer_Markers_reverseMap map[*Marker]*Layer

	OnAfterLayerCreateCallback GongOnAfterCreateInterface[Layer]
	OnAfterLayerUpdateCallback GongOnAfterUpdateInterface[Layer]
	OnAfterLayerDeleteCallback GongOnAfterDeleteInterface[Layer]

	MapViews                map[*MapView]struct{}
	MapViews_instance       map[*MapView]*MapView
	MapViews_mapString      map[string]*MapView
	MapViewOrder            uint
	MapView_stagedOrder     map[*MapView]uint
	MapView_orderStaged     map[uint]*MapView
	MapViews_reference      map[*MapView]*MapView
	MapViews_referenceOrder map[*MapView]uint

	// insertion point for slice of pointers maps
	MapView_Layers_reverseMap map[*Layer]*MapView

	OnAfterMapViewCreateCallback GongOnAfterCreateInterface[MapView]
	OnAfterMapViewUpdateCallback GongOnAfterUpdateInterface[MapView]
	OnAfterMapViewDeleteCallback GongOnAfterDeleteInterface[MapView]

	Markers                map[*Marker]struct{}
	Markers_instance       map[*Marker]*Marker
	Markers_mapString      map[string]*Marker
	MarkerOrder            uint
	Marker_stagedOrder     map[*Marker]uint
	Marker_orderStaged     map[uint]*Marker
	Markers_reference      map[*Marker]*Marker
	Markers_referenceOrder map[*Marker]uint

	// insertion point for slice of pointers maps
	OnAfterMarkerCreateCallback GongOnAfterCreateInterface[Marker]
	OnAfterMarkerUpdateCallback GongOnAfterUpdateInterface[Marker]
	OnAfterMarkerDeleteCallback GongOnAfterDeleteInterface[Marker]

	Polygons                map[*Polygon]struct{}
	Polygons_instance       map[*Polygon]*Polygon
	Polygons_mapString      map[string]*Polygon
	PolygonOrder            uint
	Polygon_stagedOrder     map[*Polygon]uint
	Polygon_orderStaged     map[uint]*Polygon
	Polygons_reference      map[*Polygon]*Polygon
	Polygons_referenceOrder map[*Polygon]uint

	// insertion point for slice of pointers maps
	Polygon_Coordinates_reverseMap map[*Coordinate]*Polygon

	OnAfterPolygonCreateCallback GongOnAfterCreateInterface[Polygon]
	OnAfterPolygonUpdateCallback GongOnAfterUpdateInterface[Polygon]
	OnAfterPolygonDeleteCallback GongOnAfterDeleteInterface[Polygon]

	Polylines                map[*Polyline]struct{}
	Polylines_instance       map[*Polyline]*Polyline
	Polylines_mapString      map[string]*Polyline
	PolylineOrder            uint
	Polyline_stagedOrder     map[*Polyline]uint
	Polyline_orderStaged     map[uint]*Polyline
	Polylines_reference      map[*Polyline]*Polyline
	Polylines_referenceOrder map[*Polyline]uint

	// insertion point for slice of pointers maps
	Polyline_Coordinates_reverseMap map[*Coordinate]*Polyline

	OnAfterPolylineCreateCallback GongOnAfterCreateInterface[Polyline]
	OnAfterPolylineUpdateCallback GongOnAfterUpdateInterface[Polyline]
	OnAfterPolylineDeleteCallback GongOnAfterDeleteInterface[Polyline]

	BackRepo GongBackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          GongOnInitCommitInterface
	OnInitCommitFromFrontCallback GongOnInitCommitInterface
	OnInitCommitFromBackCallback  GongOnInitCommitInterface

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
}

type GongStage = Stage

// RegisterBeforeCommit adds a hook that runs before the commit happens
func (s *Stage) RegisterBeforeCommit(hook func(stage *Stage)) {
	s.StageCore.RegisterBeforeCommitHook(func() { hook(s) })
}

// RegisterAfterCommit adds a hook that runs after the commit succeeds
func (s *Stage) RegisterAfterCommit(hook func(stage *Stage)) {
	s.StageCore.RegisterAfterCommitHook(func() { hook(s) })
}

// ApplyBackwardCommit applies the commit before the current one
func (stage *Stage) ApplyBackwardCommit() error {
	return stage.StageCore.ApplyBackwardCommit(stage.ParseAstString, stage.ComputeReferenceAndOrders)
}

// ApplyForwardCommit applies the commit after the current one
func (stage *Stage) ApplyForwardCommit() error {
	return stage.StageCore.ApplyForwardCommit(stage.ParseAstString, stage.ComputeReferenceAndOrders)
}

// ResetHard removes the more recent
// commitsBehind forward/backward Commits from the
// stage
func (stage *Stage) ResetHard() {
	stage.StageCore.ResetHardCore()

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	stage.StageCore.RunBeforeCommitHooks()

	// 2. Run all After Commit hooks
	stage.StageCore.RunAfterCommitHooks()
}

// Squash removes all commits and marshals the stage as a single commit
func (stage *Stage) Squash() {
	stage.StageCore.SquashCore()

	// insertion point for clear references
	__gong__clearReferences(&stage.Coordinates_reference, &stage.Coordinates_instance, &stage.Coordinates_referenceOrder)

	__gong__clearReferences(&stage.Layers_reference, &stage.Layers_instance, &stage.Layers_referenceOrder)

	__gong__clearReferences(&stage.MapViews_reference, &stage.MapViews_instance, &stage.MapViews_referenceOrder)

	__gong__clearReferences(&stage.Markers_reference, &stage.Markers_instance, &stage.Markers_referenceOrder)

	__gong__clearReferences(&stage.Polygons_reference, &stage.Polygons_instance, &stage.Polygons_referenceOrder)

	__gong__clearReferences(&stage.Polylines_reference, &stage.Polylines_instance, &stage.Polylines_referenceOrder)

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	stage.StageCore.RunBeforeCommitHooks()

	stage.StageCore.EndSquash()

	// 2. Run all After Commit hooks
	stage.StageCore.RunAfterCommitHooks()
}

// recomputeOrders recomputes the next order for each struct
// this is necessary because the order might have been incremented
// during the commits that have been discarded
// insertion point for max order recomputation
func (stage *Stage) recomputeOrders() {
	// insertion point for max order recomputation
	stage.CoordinateOrder = __gong__recomputeOrder(stage.Coordinate_stagedOrder)

	stage.LayerOrder = __gong__recomputeOrder(stage.Layer_stagedOrder)

	stage.MapViewOrder = __gong__recomputeOrder(stage.MapView_stagedOrder)

	stage.MarkerOrder = __gong__recomputeOrder(stage.Marker_stagedOrder)

	stage.PolygonOrder = __gong__recomputeOrder(stage.Polygon_stagedOrder)

	stage.PolylineOrder = __gong__recomputeOrder(stage.Polyline_stagedOrder)

	// end of insertion point for max order recomputation
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

// insertion point for stage ops
func (*Coordinate) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Coordinates, stage.Coordinate_stagedOrder)
}

func (*Coordinate) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Coordinate_orderStaged[order]
}

func (*Coordinate) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Coordinates_mapString
}

func (*Coordinate) GongGetInstancesSet(stage *Stage) any {
	return &stage.Coordinates
}

func (*Coordinate) GongNewInstance() any {
	return new(Coordinate)
}

func (*Layer) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Layers, stage.Layer_stagedOrder)
}

func (*Layer) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Layer_orderStaged[order]
}

func (*Layer) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Layers_mapString
}

func (*Layer) GongGetInstancesSet(stage *Stage) any {
	return &stage.Layers
}

func (*Layer) GongNewInstance() any {
	return new(Layer)
}

func (*MapView) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MapViews, stage.MapView_stagedOrder)
}

func (*MapView) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MapView_orderStaged[order]
}

func (*MapView) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MapViews_mapString
}

func (*MapView) GongGetInstancesSet(stage *Stage) any {
	return &stage.MapViews
}

func (*MapView) GongNewInstance() any {
	return new(MapView)
}

func (*Marker) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Markers, stage.Marker_stagedOrder)
}

func (*Marker) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Marker_orderStaged[order]
}

func (*Marker) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Markers_mapString
}

func (*Marker) GongGetInstancesSet(stage *Stage) any {
	return &stage.Markers
}

func (*Marker) GongNewInstance() any {
	return new(Marker)
}

func (*Polygon) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Polygons, stage.Polygon_stagedOrder)
}

func (*Polygon) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Polygon_orderStaged[order]
}

func (*Polygon) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Polygons_mapString
}

func (*Polygon) GongGetInstancesSet(stage *Stage) any {
	return &stage.Polygons
}

func (*Polygon) GongNewInstance() any {
	return new(Polygon)
}

func (*Polyline) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Polylines, stage.Polyline_stagedOrder)
}

func (*Polyline) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Polyline_orderStaged[order]
}

func (*Polyline) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Polylines_mapString
}

func (*Polyline) GongGetInstancesSet(stage *Stage) any {
	return &stage.Polylines
}

func (*Polyline) GongNewInstance() any {
	return new(Polyline)
}


// GetInstancesByOrder is the Stage method returning a slice of generic pointers to gongstructs
// ordered by their order in the stage.
func (stage *Stage) GetInstancesByOrder[T GongstructPtr]() (res []T) {
	if stage == nil {
		return nil
	}
	var t T
	return t.GongGetInstancesByOrder(stage).([]T)
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
	return "github.com/fullstack-lang/gong/lib/maplibre/go/models"
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
		Coordinates:           make(map[*Coordinate]struct{}),
		Coordinates_mapString: make(map[string]*Coordinate),

		Layers:           make(map[*Layer]struct{}),
		Layers_mapString: make(map[string]*Layer),

		MapViews:           make(map[*MapView]struct{}),
		MapViews_mapString: make(map[string]*MapView),

		Markers:           make(map[*Marker]struct{}),
		Markers_mapString: make(map[string]*Marker),

		Polygons:           make(map[*Polygon]struct{}),
		Polygons_mapString: make(map[string]*Polygon),

		Polylines:           make(map[*Polyline]struct{}),
		Polylines_mapString: make(map[string]*Polyline),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		Coordinate_stagedOrder: make(map[*Coordinate]uint),
		Coordinate_orderStaged: make(map[uint]*Coordinate),
		Coordinates_reference:  make(map[*Coordinate]*Coordinate),

		Layer_stagedOrder: make(map[*Layer]uint),
		Layer_orderStaged: make(map[uint]*Layer),
		Layers_reference:  make(map[*Layer]*Layer),

		MapView_stagedOrder: make(map[*MapView]uint),
		MapView_orderStaged: make(map[uint]*MapView),
		MapViews_reference:  make(map[*MapView]*MapView),

		Marker_stagedOrder: make(map[*Marker]uint),
		Marker_orderStaged: make(map[uint]*Marker),
		Markers_reference:  make(map[*Marker]*Marker),

		Polygon_stagedOrder: make(map[*Polygon]uint),
		Polygon_orderStaged: make(map[uint]*Polygon),
		Polygons_reference:  make(map[*Polygon]*Polygon),

		Polyline_stagedOrder: make(map[*Polyline]uint),
		Polyline_orderStaged: make(map[uint]*Polyline),
		Polylines_reference:  make(map[*Polyline]*Polyline),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"Coordinate": &CoordinateUnmarshaller{},

			"Layer": &LayerUnmarshaller{},

			"MapView": &MapViewUnmarshaller{},

			"Marker": &MarkerUnmarshaller{},

			"Polygon": &PolygonUnmarshaller{},

			"Polyline": &PolylineUnmarshaller{},

			// end of insertion point
		},
	}
	stage.StageCore.SetName(name)
	stage.StageCore.SetNavigationMode(GongNavigationModeNormal)

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
	if stage == nil {
		return
	}
	var t Type
	val := t.GongGetInstanceFromOrder(stage, order)
	if val != nil {
		res = val.(Type)
	}
	return
}

func (stage *Stage) CommitWithSuspendedCallbacks() {
	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	tmp2 := stage.StageCore.GetBeforeCommitHooks()
	stage.StageCore.ClearBeforeCommitHooks()
	tmp3 := stage.StageCore.GetAfterCommitHooks()
	stage.StageCore.ClearAfterCommitHooks()
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
	stage.StageCore.SetBeforeCommitHooks(tmp2)
	stage.StageCore.SetAfterCommitHooks(tmp3)
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
	stage.StageCore.RunBeforeCommitHooks()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}
	stage.ComputeInstancesNb()

	// if a commit is applied when in navigation mode
	// this will reset the commits behind and swith the
	// naviagation
	if stage.IsInDeltaMode() && stage.GetNavigationMode() == GongNavigationModeNavigating && stage.GetCommitsBehind() > 0 {
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
	stage.StageCore.RunAfterCommitHooks()
}

func (stage *Stage) ComputeInstancesNb() {
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Coordinate"] = len(stage.Coordinates)
	stage.Map_GongStructName_InstancesNb["Layer"] = len(stage.Layers)
	stage.Map_GongStructName_InstancesNb["MapView"] = len(stage.MapViews)
	stage.Map_GongStructName_InstancesNb["Marker"] = len(stage.Markers)
	stage.Map_GongStructName_InstancesNb["Polygon"] = len(stage.Polygons)
	stage.Map_GongStructName_InstancesNb["Polyline"] = len(stage.Polylines)
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
// Stage puts coordinate to the model stage
func (coordinate *Coordinate) Stage(stage *Stage) *Coordinate {
	__gong__stage(stage.Coordinates, stage.Coordinate_stagedOrder, stage.Coordinate_orderStaged, &stage.CoordinateOrder, stage.Coordinates_mapString, coordinate, coordinate.Name)
	return coordinate
}

// StagePreserveOrder puts coordinate to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CoordinateOrder
// - update stage.CoordinateOrder accordingly
func (coordinate *Coordinate) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Coordinates, stage.Coordinate_stagedOrder, stage.Coordinate_orderStaged, &stage.CoordinateOrder, stage.Coordinates_mapString, coordinate, order, coordinate.Name)
}

// Unstage removes coordinate off the model stage
func (coordinate *Coordinate) Unstage(stage *Stage) *Coordinate {
	__gong__unstage(stage.Coordinates, stage.Coordinates_mapString, coordinate, coordinate.Name)
	return coordinate
}

// UnstageVoid removes coordinate off the model stage
func (coordinate *Coordinate) UnstageVoid(stage *Stage) {
	coordinate.Unstage(stage)
}

func (coordinate *Coordinate) StageVoid(stage *Stage) {
	coordinate.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coordinate *Coordinate) GetName() (res string) {
	return coordinate.Name
}

// for satisfaction of GongStruct interface
func (coordinate *Coordinate) SetName(name string) {
	coordinate.Name = name
}

// Stage puts layer to the model stage
func (layer *Layer) Stage(stage *Stage) *Layer {
	__gong__stage(stage.Layers, stage.Layer_stagedOrder, stage.Layer_orderStaged, &stage.LayerOrder, stage.Layers_mapString, layer, layer.Name)
	return layer
}

// StagePreserveOrder puts layer to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LayerOrder
// - update stage.LayerOrder accordingly
func (layer *Layer) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Layers, stage.Layer_stagedOrder, stage.Layer_orderStaged, &stage.LayerOrder, stage.Layers_mapString, layer, order, layer.Name)
}

// Unstage removes layer off the model stage
func (layer *Layer) Unstage(stage *Stage) *Layer {
	__gong__unstage(stage.Layers, stage.Layers_mapString, layer, layer.Name)
	return layer
}

// UnstageVoid removes layer off the model stage
func (layer *Layer) UnstageVoid(stage *Stage) {
	layer.Unstage(stage)
}

func (layer *Layer) StageVoid(stage *Stage) {
	layer.Stage(stage)
}

// for satisfaction of GongStruct interface
func (layer *Layer) GetName() (res string) {
	return layer.Name
}

// for satisfaction of GongStruct interface
func (layer *Layer) SetName(name string) {
	layer.Name = name
}

// Stage puts mapview to the model stage
func (mapview *MapView) Stage(stage *Stage) *MapView {
	__gong__stage(stage.MapViews, stage.MapView_stagedOrder, stage.MapView_orderStaged, &stage.MapViewOrder, stage.MapViews_mapString, mapview, mapview.Name)
	return mapview
}

// StagePreserveOrder puts mapview to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MapViewOrder
// - update stage.MapViewOrder accordingly
func (mapview *MapView) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MapViews, stage.MapView_stagedOrder, stage.MapView_orderStaged, &stage.MapViewOrder, stage.MapViews_mapString, mapview, order, mapview.Name)
}

// Unstage removes mapview off the model stage
func (mapview *MapView) Unstage(stage *Stage) *MapView {
	__gong__unstage(stage.MapViews, stage.MapViews_mapString, mapview, mapview.Name)
	return mapview
}

// UnstageVoid removes mapview off the model stage
func (mapview *MapView) UnstageVoid(stage *Stage) {
	mapview.Unstage(stage)
}

func (mapview *MapView) StageVoid(stage *Stage) {
	mapview.Stage(stage)
}

// for satisfaction of GongStruct interface
func (mapview *MapView) GetName() (res string) {
	return mapview.Name
}

// for satisfaction of GongStruct interface
func (mapview *MapView) SetName(name string) {
	mapview.Name = name
}

// Stage puts marker to the model stage
func (marker *Marker) Stage(stage *Stage) *Marker {
	__gong__stage(stage.Markers, stage.Marker_stagedOrder, stage.Marker_orderStaged, &stage.MarkerOrder, stage.Markers_mapString, marker, marker.Name)
	return marker
}

// StagePreserveOrder puts marker to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MarkerOrder
// - update stage.MarkerOrder accordingly
func (marker *Marker) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Markers, stage.Marker_stagedOrder, stage.Marker_orderStaged, &stage.MarkerOrder, stage.Markers_mapString, marker, order, marker.Name)
}

// Unstage removes marker off the model stage
func (marker *Marker) Unstage(stage *Stage) *Marker {
	__gong__unstage(stage.Markers, stage.Markers_mapString, marker, marker.Name)
	return marker
}

// UnstageVoid removes marker off the model stage
func (marker *Marker) UnstageVoid(stage *Stage) {
	marker.Unstage(stage)
}

func (marker *Marker) StageVoid(stage *Stage) {
	marker.Stage(stage)
}

// for satisfaction of GongStruct interface
func (marker *Marker) GetName() (res string) {
	return marker.Name
}

// for satisfaction of GongStruct interface
func (marker *Marker) SetName(name string) {
	marker.Name = name
}

// Stage puts polygon to the model stage
func (polygon *Polygon) Stage(stage *Stage) *Polygon {
	__gong__stage(stage.Polygons, stage.Polygon_stagedOrder, stage.Polygon_orderStaged, &stage.PolygonOrder, stage.Polygons_mapString, polygon, polygon.Name)
	return polygon
}

// StagePreserveOrder puts polygon to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PolygonOrder
// - update stage.PolygonOrder accordingly
func (polygon *Polygon) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Polygons, stage.Polygon_stagedOrder, stage.Polygon_orderStaged, &stage.PolygonOrder, stage.Polygons_mapString, polygon, order, polygon.Name)
}

// Unstage removes polygon off the model stage
func (polygon *Polygon) Unstage(stage *Stage) *Polygon {
	__gong__unstage(stage.Polygons, stage.Polygons_mapString, polygon, polygon.Name)
	return polygon
}

// UnstageVoid removes polygon off the model stage
func (polygon *Polygon) UnstageVoid(stage *Stage) {
	polygon.Unstage(stage)
}

func (polygon *Polygon) StageVoid(stage *Stage) {
	polygon.Stage(stage)
}

// for satisfaction of GongStruct interface
func (polygon *Polygon) GetName() (res string) {
	return polygon.Name
}

// for satisfaction of GongStruct interface
func (polygon *Polygon) SetName(name string) {
	polygon.Name = name
}

// Stage puts polyline to the model stage
func (polyline *Polyline) Stage(stage *Stage) *Polyline {
	__gong__stage(stage.Polylines, stage.Polyline_stagedOrder, stage.Polyline_orderStaged, &stage.PolylineOrder, stage.Polylines_mapString, polyline, polyline.Name)
	return polyline
}

// StagePreserveOrder puts polyline to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PolylineOrder
// - update stage.PolylineOrder accordingly
func (polyline *Polyline) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Polylines, stage.Polyline_stagedOrder, stage.Polyline_orderStaged, &stage.PolylineOrder, stage.Polylines_mapString, polyline, order, polyline.Name)
}

// Unstage removes polyline off the model stage
func (polyline *Polyline) Unstage(stage *Stage) *Polyline {
	__gong__unstage(stage.Polylines, stage.Polylines_mapString, polyline, polyline.Name)
	return polyline
}

// UnstageVoid removes polyline off the model stage
func (polyline *Polyline) UnstageVoid(stage *Stage) {
	polyline.Unstage(stage)
}

func (polyline *Polyline) StageVoid(stage *Stage) {
	polyline.Stage(stage)
}

// for satisfaction of GongStruct interface
func (polyline *Polyline) GetName() (res string) {
	return polyline.Name
}

// for satisfaction of GongStruct interface
func (polyline *Polyline) SetName(name string) {
	polyline.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.Coordinates, &stage.Coordinates_mapString, &stage.Coordinate_stagedOrder, &stage.CoordinateOrder)

	__gong__resetStageType(&stage.Layers, &stage.Layers_mapString, &stage.Layer_stagedOrder, &stage.LayerOrder)

	__gong__resetStageType(&stage.MapViews, &stage.MapViews_mapString, &stage.MapView_stagedOrder, &stage.MapViewOrder)

	__gong__resetStageType(&stage.Markers, &stage.Markers_mapString, &stage.Marker_stagedOrder, &stage.MarkerOrder)

	__gong__resetStageType(&stage.Polygons, &stage.Polygons_mapString, &stage.Polygon_stagedOrder, &stage.PolygonOrder)

	__gong__resetStageType(&stage.Polylines, &stage.Polylines_mapString, &stage.Polyline_stagedOrder, &stage.PolylineOrder)

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
type Gongstruct interface {
	GongGetAssociationName() any
}

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

	GongGetInstancesByOrder(stage *Stage) any
	GongGetInstanceFromOrder(stage *Stage, order uint) any
	GongGetInstancesMapByName(stage *Stage) any
	GongGetInstancesSet(stage *Stage) any
	GongNewInstance() any
	GongGetReverseFields() []GongReverseField
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
	if stage == nil {
		return nil
	}
	var t Type
	if any(t) == nil {
		return nil
	}
	return t.GongGetInstancesMapByName(stage).(map[string]Type)
}

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
	if stage == nil {
		return nil
	}
	var t Type
	return t.GongGetInstancesSet(stage).(*map[Type]struct{})
}

// insertion point for instance with special fields
func (Coordinate) GongGetAssociationName() any {
	return &Coordinate{
	}
}

func (Layer) GongGetAssociationName() any {
	return &Layer{
			Polygons: []*Polygon{{Name: "Polygons"}},
			Polylines: []*Polyline{{Name: "Polylines"}},
			Markers: []*Marker{{Name: "Markers"}},
	}
}

func (MapView) GongGetAssociationName() any {
	return &MapView{
			Layers: []*Layer{{Name: "Layers"}},
	}
}

func (Marker) GongGetAssociationName() any {
	return &Marker{
	}
}

func (Polygon) GongGetAssociationName() any {
	return &Polygon{
			Coordinates: []*Coordinate{{Name: "Coordinates"}},
	}
}

func (Polyline) GongGetAssociationName() any {
	return &Polyline{
			Coordinates: []*Coordinate{{Name: "Coordinates"}},
	}
}


// GongGetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GongGetAssociationName[Type Gongstruct]() *Type {
	var t Type
	return t.GongGetAssociationName().(*Type)
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
	// reverse maps of direct associations of Coordinate
	case Coordinate:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Layer
	case Layer:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MapView
	case MapView:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Marker
	case Marker:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Polygon
	case Polygon:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Polyline
	case Polyline:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap is the Stage method for backtrack navigation of slice-of-pointers associations.
func (stage *Stage) GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string) map[*End][]*Start {
	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Coordinate
	case Coordinate:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Layer
	case Layer:
		switch fieldname {
		// insertion point for per direct association field
		case "Polygons":
			res := make(map[*Polygon][]*Layer)
			for layer := range stage.Layers {
				for _, polygon_ := range layer.Polygons {
					res[polygon_] = append(res[polygon_], layer)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Polylines":
			res := make(map[*Polyline][]*Layer)
			for layer := range stage.Layers {
				for _, polyline_ := range layer.Polylines {
					res[polyline_] = append(res[polyline_], layer)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Markers":
			res := make(map[*Marker][]*Layer)
			for layer := range stage.Layers {
				for _, marker_ := range layer.Markers {
					res[marker_] = append(res[marker_], layer)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MapView
	case MapView:
		switch fieldname {
		// insertion point for per direct association field
		case "Layers":
			res := make(map[*Layer][]*MapView)
			for mapview := range stage.MapViews {
				for _, layer_ := range mapview.Layers {
					res[layer_] = append(res[layer_], mapview)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Marker
	case Marker:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Polygon
	case Polygon:
		switch fieldname {
		// insertion point for per direct association field
		case "Coordinates":
			res := make(map[*Coordinate][]*Polygon)
			for polygon := range stage.Polygons {
				for _, coordinate_ := range polygon.Coordinates {
					res[coordinate_] = append(res[coordinate_], polygon)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Polyline
	case Polyline:
		switch fieldname {
		// insertion point for per direct association field
		case "Coordinates":
			res := make(map[*Coordinate][]*Polyline)
			for polyline := range stage.Polylines {
				for _, coordinate_ := range polyline.Coordinates {
					res[coordinate_] = append(res[coordinate_], polyline)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	}
	return nil
}

// GongNewInstance creates a new instance of the Gongstruct
func GongNewInstance[Type GongstructPtr]() (res Type) {
	var t Type
	return t.GongNewInstance().(Type)
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
	var t Type
	if any(t) == nil {
		return ""
	}
	return t.GongGetGongstructName()
}

func GetPointerToGongstructName[Type GongstructIF]() (res string) {
	return GongGetPointerToGongstructName[Type]()
}

type GongReverseField struct {
	GongstructName string
	Fieldname      string
}

type ReverseField = GongReverseField

// insertion point for generic get reverse fields
func (*Coordinate) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Polygon",
			Fieldname: "Coordinates",
		},
		{
			GongstructName: "Polyline",
			Fieldname: "Coordinates",
		},
	}
}

func (*Layer) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "MapView",
			Fieldname: "Layers",
		},
	}
}

func (*MapView) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Marker) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Layer",
			Fieldname: "Markers",
		},
	}
}

func (*Polygon) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Layer",
			Fieldname: "Polygons",
		},
	}
}

func (*Polyline) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Layer",
			Fieldname: "Polylines",
		},
	}
}


func GongGetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	var t Type
	if any(t) == nil {
		return nil
	}
	return t.GongGetReverseFields()
}

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
}

// insertion point for get fields header method
func (coordinate *Coordinate) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Longitude",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Latitude",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Altitude",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (layer *Layer) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsVisible",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Polygons",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Polygon",
		},
		{
			Name:                 "Polylines",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Polyline",
		},
		{
			Name:                 "Markers",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Marker",
		},
	}
	return
}

func (mapview *MapView) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StyleURL",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Longitude",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Latitude",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Zoom",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Pitch",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Bearing",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "Layers",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Layer",
		},
	}
	return
}

func (marker *Marker) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Label",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "PopupHTML",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Longitude",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Latitude",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Altitude",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "SVG",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Scale",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Rotation",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "RotationAlignment",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "MarkerPitchAlignmentEnum",
		},
		{
			Name:                 "PitchAlignment",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "MarkerPitchAlignmentEnum",
		},
		{
			Name:                 "Anchor",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "MarkerAnchorEnum",
		},
		{
			Name:               "IsDraggable",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (polygon *Polygon) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Label",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "PopupHTML",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Coordinates",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Coordinate",
		},
		{
			Name:               "FillColor",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeColor",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsExtruded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "BaseAltitude",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ExtrusionHeight",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (polyline *Polyline) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Label",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "PopupHTML",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Coordinates",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Coordinate",
		},
		{
			Name:               "StrokeColor",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
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
func (coordinate *Coordinate) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coordinate.Name
	case "Longitude":
		res.valueString = fmt.Sprintf("%f", coordinate.Longitude)
		res.valueFloat = coordinate.Longitude
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Latitude":
		res.valueString = fmt.Sprintf("%f", coordinate.Latitude)
		res.valueFloat = coordinate.Latitude
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Altitude":
		res.valueString = fmt.Sprintf("%f", coordinate.Altitude)
		res.valueFloat = coordinate.Altitude
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (layer *Layer) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = layer.Name
	case "IsVisible":
		res.valueString = fmt.Sprintf("%t", layer.IsVisible)
		res.valueBool = layer.IsVisible
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Polygons":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.Polygons {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Polylines":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.Polylines {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Markers":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.Markers {
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

func (mapview *MapView) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = mapview.Name
	case "StyleURL":
		res.valueString = mapview.StyleURL
	case "Longitude":
		res.valueString = fmt.Sprintf("%f", mapview.Longitude)
		res.valueFloat = mapview.Longitude
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Latitude":
		res.valueString = fmt.Sprintf("%f", mapview.Latitude)
		res.valueFloat = mapview.Latitude
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Zoom":
		res.valueString = fmt.Sprintf("%f", mapview.Zoom)
		res.valueFloat = mapview.Zoom
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Pitch":
		res.valueString = fmt.Sprintf("%f", mapview.Pitch)
		res.valueFloat = mapview.Pitch
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Bearing":
		res.valueString = fmt.Sprintf("%f", mapview.Bearing)
		res.valueFloat = mapview.Bearing
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Layers":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range mapview.Layers {
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

func (marker *Marker) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = marker.Name
	case "Label":
		res.valueString = marker.Label
	case "PopupHTML":
		res.valueString = marker.PopupHTML
	case "Longitude":
		res.valueString = fmt.Sprintf("%f", marker.Longitude)
		res.valueFloat = marker.Longitude
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Latitude":
		res.valueString = fmt.Sprintf("%f", marker.Latitude)
		res.valueFloat = marker.Latitude
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Altitude":
		res.valueString = fmt.Sprintf("%f", marker.Altitude)
		res.valueFloat = marker.Altitude
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "SVG":
		res.valueString = marker.SVG
	case "Color":
		res.valueString = marker.Color
	case "Scale":
		res.valueString = fmt.Sprintf("%f", marker.Scale)
		res.valueFloat = marker.Scale
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Rotation":
		res.valueString = fmt.Sprintf("%f", marker.Rotation)
		res.valueFloat = marker.Rotation
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RotationAlignment":
		enum := marker.RotationAlignment
		res.valueString = enum.ToCodeString()
	case "PitchAlignment":
		enum := marker.PitchAlignment
		res.valueString = enum.ToCodeString()
	case "Anchor":
		enum := marker.Anchor
		res.valueString = enum.ToCodeString()
	case "IsDraggable":
		res.valueString = fmt.Sprintf("%t", marker.IsDraggable)
		res.valueBool = marker.IsDraggable
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (polygon *Polygon) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = polygon.Name
	case "Label":
		res.valueString = polygon.Label
	case "PopupHTML":
		res.valueString = polygon.PopupHTML
	case "Coordinates":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range polygon.Coordinates {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "FillColor":
		res.valueString = polygon.FillColor
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", polygon.FillOpacity)
		res.valueFloat = polygon.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeColor":
		res.valueString = polygon.StrokeColor
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", polygon.StrokeOpacity)
		res.valueFloat = polygon.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", polygon.StrokeWidth)
		res.valueFloat = polygon.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsExtruded":
		res.valueString = fmt.Sprintf("%t", polygon.IsExtruded)
		res.valueBool = polygon.IsExtruded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "BaseAltitude":
		res.valueString = fmt.Sprintf("%f", polygon.BaseAltitude)
		res.valueFloat = polygon.BaseAltitude
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ExtrusionHeight":
		res.valueString = fmt.Sprintf("%f", polygon.ExtrusionHeight)
		res.valueFloat = polygon.ExtrusionHeight
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (polyline *Polyline) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = polyline.Name
	case "Label":
		res.valueString = polyline.Label
	case "PopupHTML":
		res.valueString = polyline.PopupHTML
	case "Coordinates":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range polyline.Coordinates {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "StrokeColor":
		res.valueString = polyline.StrokeColor
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", polyline.StrokeOpacity)
		res.valueFloat = polyline.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", polyline.StrokeWidth)
		res.valueFloat = polyline.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = polyline.StrokeDashArray
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
func (coordinate *Coordinate) GongGetGongstructName() string {
	return "Coordinate"
}

func (layer *Layer) GongGetGongstructName() string {
	return "Layer"
}

func (mapview *MapView) GongGetGongstructName() string {
	return "MapView"
}

func (marker *Marker) GongGetGongstructName() string {
	return "Marker"
}

func (polygon *Polygon) GongGetGongstructName() string {
	return "Polygon"
}

func (polyline *Polyline) GongGetGongstructName() string {
	return "Polyline"
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
	__gong__rebuildMapString(stage.Coordinates, &stage.Coordinates_mapString)

	__gong__rebuildMapString(stage.Layers, &stage.Layers_mapString)

	__gong__rebuildMapString(stage.MapViews, &stage.MapViews_mapString)

	__gong__rebuildMapString(stage.Markers, &stage.Markers_mapString)

	__gong__rebuildMapString(stage.Polygons, &stage.Polygons_mapString)

	__gong__rebuildMapString(stage.Polylines, &stage.Polylines_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
