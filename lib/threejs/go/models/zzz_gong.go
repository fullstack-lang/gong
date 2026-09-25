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
	AmbiantLights                map[*AmbiantLight]struct{}
	AmbiantLights_instance       map[*AmbiantLight]*AmbiantLight
	AmbiantLights_mapString      map[string]*AmbiantLight
	AmbiantLightOrder            uint
	AmbiantLight_stagedOrder     map[*AmbiantLight]uint
	AmbiantLight_orderStaged     map[uint]*AmbiantLight
	AmbiantLights_reference      map[*AmbiantLight]*AmbiantLight
	AmbiantLights_referenceOrder map[*AmbiantLight]uint

	// insertion point for slice of pointers maps
	OnAfterAmbiantLightCreateCallback GongOnAfterCreateInterface[AmbiantLight]
	OnAfterAmbiantLightUpdateCallback GongOnAfterUpdateInterface[AmbiantLight]
	OnAfterAmbiantLightDeleteCallback GongOnAfterDeleteInterface[AmbiantLight]

	BoxGeometrys                map[*BoxGeometry]struct{}
	BoxGeometrys_instance       map[*BoxGeometry]*BoxGeometry
	BoxGeometrys_mapString      map[string]*BoxGeometry
	BoxGeometryOrder            uint
	BoxGeometry_stagedOrder     map[*BoxGeometry]uint
	BoxGeometry_orderStaged     map[uint]*BoxGeometry
	BoxGeometrys_reference      map[*BoxGeometry]*BoxGeometry
	BoxGeometrys_referenceOrder map[*BoxGeometry]uint

	// insertion point for slice of pointers maps
	OnAfterBoxGeometryCreateCallback GongOnAfterCreateInterface[BoxGeometry]
	OnAfterBoxGeometryUpdateCallback GongOnAfterUpdateInterface[BoxGeometry]
	OnAfterBoxGeometryDeleteCallback GongOnAfterDeleteInterface[BoxGeometry]

	BufferGeometrys                map[*BufferGeometry]struct{}
	BufferGeometrys_instance       map[*BufferGeometry]*BufferGeometry
	BufferGeometrys_mapString      map[string]*BufferGeometry
	BufferGeometryOrder            uint
	BufferGeometry_stagedOrder     map[*BufferGeometry]uint
	BufferGeometry_orderStaged     map[uint]*BufferGeometry
	BufferGeometrys_reference      map[*BufferGeometry]*BufferGeometry
	BufferGeometrys_referenceOrder map[*BufferGeometry]uint

	// insertion point for slice of pointers maps
	BufferGeometry_Vertices_reverseMap map[*Vector3]*BufferGeometry

	BufferGeometry_Faces_reverseMap map[*Triangle]*BufferGeometry

	OnAfterBufferGeometryCreateCallback GongOnAfterCreateInterface[BufferGeometry]
	OnAfterBufferGeometryUpdateCallback GongOnAfterUpdateInterface[BufferGeometry]
	OnAfterBufferGeometryDeleteCallback GongOnAfterDeleteInterface[BufferGeometry]

	Cameras                map[*Camera]struct{}
	Cameras_instance       map[*Camera]*Camera
	Cameras_mapString      map[string]*Camera
	CameraOrder            uint
	Camera_stagedOrder     map[*Camera]uint
	Camera_orderStaged     map[uint]*Camera
	Cameras_reference      map[*Camera]*Camera
	Cameras_referenceOrder map[*Camera]uint

	// insertion point for slice of pointers maps
	OnAfterCameraCreateCallback GongOnAfterCreateInterface[Camera]
	OnAfterCameraUpdateCallback GongOnAfterUpdateInterface[Camera]
	OnAfterCameraDeleteCallback GongOnAfterDeleteInterface[Camera]

	Canvass                map[*Canvas]struct{}
	Canvass_instance       map[*Canvas]*Canvas
	Canvass_mapString      map[string]*Canvas
	CanvasOrder            uint
	Canvas_stagedOrder     map[*Canvas]uint
	Canvas_orderStaged     map[uint]*Canvas
	Canvass_reference      map[*Canvas]*Canvas
	Canvass_referenceOrder map[*Canvas]uint

	// insertion point for slice of pointers maps
	Canvas_DirectionalLights_reverseMap map[*DirectionalLight]*Canvas

	Canvas_Meshs_reverseMap map[*Mesh]*Canvas

	OnAfterCanvasCreateCallback GongOnAfterCreateInterface[Canvas]
	OnAfterCanvasUpdateCallback GongOnAfterUpdateInterface[Canvas]
	OnAfterCanvasDeleteCallback GongOnAfterDeleteInterface[Canvas]

	Curves                map[*Curve]struct{}
	Curves_instance       map[*Curve]*Curve
	Curves_mapString      map[string]*Curve
	CurveOrder            uint
	Curve_stagedOrder     map[*Curve]uint
	Curve_orderStaged     map[uint]*Curve
	Curves_reference      map[*Curve]*Curve
	Curves_referenceOrder map[*Curve]uint

	// insertion point for slice of pointers maps
	Curve_Points_reverseMap map[*Vector3]*Curve

	OnAfterCurveCreateCallback GongOnAfterCreateInterface[Curve]
	OnAfterCurveUpdateCallback GongOnAfterUpdateInterface[Curve]
	OnAfterCurveDeleteCallback GongOnAfterDeleteInterface[Curve]

	CylinderGeometrys                map[*CylinderGeometry]struct{}
	CylinderGeometrys_instance       map[*CylinderGeometry]*CylinderGeometry
	CylinderGeometrys_mapString      map[string]*CylinderGeometry
	CylinderGeometryOrder            uint
	CylinderGeometry_stagedOrder     map[*CylinderGeometry]uint
	CylinderGeometry_orderStaged     map[uint]*CylinderGeometry
	CylinderGeometrys_reference      map[*CylinderGeometry]*CylinderGeometry
	CylinderGeometrys_referenceOrder map[*CylinderGeometry]uint

	// insertion point for slice of pointers maps
	OnAfterCylinderGeometryCreateCallback GongOnAfterCreateInterface[CylinderGeometry]
	OnAfterCylinderGeometryUpdateCallback GongOnAfterUpdateInterface[CylinderGeometry]
	OnAfterCylinderGeometryDeleteCallback GongOnAfterDeleteInterface[CylinderGeometry]

	DirectionalLights                map[*DirectionalLight]struct{}
	DirectionalLights_instance       map[*DirectionalLight]*DirectionalLight
	DirectionalLights_mapString      map[string]*DirectionalLight
	DirectionalLightOrder            uint
	DirectionalLight_stagedOrder     map[*DirectionalLight]uint
	DirectionalLight_orderStaged     map[uint]*DirectionalLight
	DirectionalLights_reference      map[*DirectionalLight]*DirectionalLight
	DirectionalLights_referenceOrder map[*DirectionalLight]uint

	// insertion point for slice of pointers maps
	OnAfterDirectionalLightCreateCallback GongOnAfterCreateInterface[DirectionalLight]
	OnAfterDirectionalLightUpdateCallback GongOnAfterUpdateInterface[DirectionalLight]
	OnAfterDirectionalLightDeleteCallback GongOnAfterDeleteInterface[DirectionalLight]

	ExtrudeGeometrys                map[*ExtrudeGeometry]struct{}
	ExtrudeGeometrys_instance       map[*ExtrudeGeometry]*ExtrudeGeometry
	ExtrudeGeometrys_mapString      map[string]*ExtrudeGeometry
	ExtrudeGeometryOrder            uint
	ExtrudeGeometry_stagedOrder     map[*ExtrudeGeometry]uint
	ExtrudeGeometry_orderStaged     map[uint]*ExtrudeGeometry
	ExtrudeGeometrys_reference      map[*ExtrudeGeometry]*ExtrudeGeometry
	ExtrudeGeometrys_referenceOrder map[*ExtrudeGeometry]uint

	// insertion point for slice of pointers maps
	OnAfterExtrudeGeometryCreateCallback GongOnAfterCreateInterface[ExtrudeGeometry]
	OnAfterExtrudeGeometryUpdateCallback GongOnAfterUpdateInterface[ExtrudeGeometry]
	OnAfterExtrudeGeometryDeleteCallback GongOnAfterDeleteInterface[ExtrudeGeometry]

	Meshs                map[*Mesh]struct{}
	Meshs_instance       map[*Mesh]*Mesh
	Meshs_mapString      map[string]*Mesh
	MeshOrder            uint
	Mesh_stagedOrder     map[*Mesh]uint
	Mesh_orderStaged     map[uint]*Mesh
	Meshs_reference      map[*Mesh]*Mesh
	Meshs_referenceOrder map[*Mesh]uint

	// insertion point for slice of pointers maps
	OnAfterMeshCreateCallback GongOnAfterCreateInterface[Mesh]
	OnAfterMeshUpdateCallback GongOnAfterUpdateInterface[Mesh]
	OnAfterMeshDeleteCallback GongOnAfterDeleteInterface[Mesh]

	MeshMaterialBasics                map[*MeshMaterialBasic]struct{}
	MeshMaterialBasics_instance       map[*MeshMaterialBasic]*MeshMaterialBasic
	MeshMaterialBasics_mapString      map[string]*MeshMaterialBasic
	MeshMaterialBasicOrder            uint
	MeshMaterialBasic_stagedOrder     map[*MeshMaterialBasic]uint
	MeshMaterialBasic_orderStaged     map[uint]*MeshMaterialBasic
	MeshMaterialBasics_reference      map[*MeshMaterialBasic]*MeshMaterialBasic
	MeshMaterialBasics_referenceOrder map[*MeshMaterialBasic]uint

	// insertion point for slice of pointers maps
	OnAfterMeshMaterialBasicCreateCallback GongOnAfterCreateInterface[MeshMaterialBasic]
	OnAfterMeshMaterialBasicUpdateCallback GongOnAfterUpdateInterface[MeshMaterialBasic]
	OnAfterMeshMaterialBasicDeleteCallback GongOnAfterDeleteInterface[MeshMaterialBasic]

	MeshPhysicalMaterials                map[*MeshPhysicalMaterial]struct{}
	MeshPhysicalMaterials_instance       map[*MeshPhysicalMaterial]*MeshPhysicalMaterial
	MeshPhysicalMaterials_mapString      map[string]*MeshPhysicalMaterial
	MeshPhysicalMaterialOrder            uint
	MeshPhysicalMaterial_stagedOrder     map[*MeshPhysicalMaterial]uint
	MeshPhysicalMaterial_orderStaged     map[uint]*MeshPhysicalMaterial
	MeshPhysicalMaterials_reference      map[*MeshPhysicalMaterial]*MeshPhysicalMaterial
	MeshPhysicalMaterials_referenceOrder map[*MeshPhysicalMaterial]uint

	// insertion point for slice of pointers maps
	OnAfterMeshPhysicalMaterialCreateCallback GongOnAfterCreateInterface[MeshPhysicalMaterial]
	OnAfterMeshPhysicalMaterialUpdateCallback GongOnAfterUpdateInterface[MeshPhysicalMaterial]
	OnAfterMeshPhysicalMaterialDeleteCallback GongOnAfterDeleteInterface[MeshPhysicalMaterial]

	PlaneGeometrys                map[*PlaneGeometry]struct{}
	PlaneGeometrys_instance       map[*PlaneGeometry]*PlaneGeometry
	PlaneGeometrys_mapString      map[string]*PlaneGeometry
	PlaneGeometryOrder            uint
	PlaneGeometry_stagedOrder     map[*PlaneGeometry]uint
	PlaneGeometry_orderStaged     map[uint]*PlaneGeometry
	PlaneGeometrys_reference      map[*PlaneGeometry]*PlaneGeometry
	PlaneGeometrys_referenceOrder map[*PlaneGeometry]uint

	// insertion point for slice of pointers maps
	OnAfterPlaneGeometryCreateCallback GongOnAfterCreateInterface[PlaneGeometry]
	OnAfterPlaneGeometryUpdateCallback GongOnAfterUpdateInterface[PlaneGeometry]
	OnAfterPlaneGeometryDeleteCallback GongOnAfterDeleteInterface[PlaneGeometry]

	Shapes                map[*Shape]struct{}
	Shapes_instance       map[*Shape]*Shape
	Shapes_mapString      map[string]*Shape
	ShapeOrder            uint
	Shape_stagedOrder     map[*Shape]uint
	Shape_orderStaged     map[uint]*Shape
	Shapes_reference      map[*Shape]*Shape
	Shapes_referenceOrder map[*Shape]uint

	// insertion point for slice of pointers maps
	Shape_Points_reverseMap map[*Vector2]*Shape

	OnAfterShapeCreateCallback GongOnAfterCreateInterface[Shape]
	OnAfterShapeUpdateCallback GongOnAfterUpdateInterface[Shape]
	OnAfterShapeDeleteCallback GongOnAfterDeleteInterface[Shape]

	SphereGeometrys                map[*SphereGeometry]struct{}
	SphereGeometrys_instance       map[*SphereGeometry]*SphereGeometry
	SphereGeometrys_mapString      map[string]*SphereGeometry
	SphereGeometryOrder            uint
	SphereGeometry_stagedOrder     map[*SphereGeometry]uint
	SphereGeometry_orderStaged     map[uint]*SphereGeometry
	SphereGeometrys_reference      map[*SphereGeometry]*SphereGeometry
	SphereGeometrys_referenceOrder map[*SphereGeometry]uint

	// insertion point for slice of pointers maps
	OnAfterSphereGeometryCreateCallback GongOnAfterCreateInterface[SphereGeometry]
	OnAfterSphereGeometryUpdateCallback GongOnAfterUpdateInterface[SphereGeometry]
	OnAfterSphereGeometryDeleteCallback GongOnAfterDeleteInterface[SphereGeometry]

	TorusGeometrys                map[*TorusGeometry]struct{}
	TorusGeometrys_instance       map[*TorusGeometry]*TorusGeometry
	TorusGeometrys_mapString      map[string]*TorusGeometry
	TorusGeometryOrder            uint
	TorusGeometry_stagedOrder     map[*TorusGeometry]uint
	TorusGeometry_orderStaged     map[uint]*TorusGeometry
	TorusGeometrys_reference      map[*TorusGeometry]*TorusGeometry
	TorusGeometrys_referenceOrder map[*TorusGeometry]uint

	// insertion point for slice of pointers maps
	OnAfterTorusGeometryCreateCallback GongOnAfterCreateInterface[TorusGeometry]
	OnAfterTorusGeometryUpdateCallback GongOnAfterUpdateInterface[TorusGeometry]
	OnAfterTorusGeometryDeleteCallback GongOnAfterDeleteInterface[TorusGeometry]

	Triangles                map[*Triangle]struct{}
	Triangles_instance       map[*Triangle]*Triangle
	Triangles_mapString      map[string]*Triangle
	TriangleOrder            uint
	Triangle_stagedOrder     map[*Triangle]uint
	Triangle_orderStaged     map[uint]*Triangle
	Triangles_reference      map[*Triangle]*Triangle
	Triangles_referenceOrder map[*Triangle]uint

	// insertion point for slice of pointers maps
	OnAfterTriangleCreateCallback GongOnAfterCreateInterface[Triangle]
	OnAfterTriangleUpdateCallback GongOnAfterUpdateInterface[Triangle]
	OnAfterTriangleDeleteCallback GongOnAfterDeleteInterface[Triangle]

	TubeGeometrys                map[*TubeGeometry]struct{}
	TubeGeometrys_instance       map[*TubeGeometry]*TubeGeometry
	TubeGeometrys_mapString      map[string]*TubeGeometry
	TubeGeometryOrder            uint
	TubeGeometry_stagedOrder     map[*TubeGeometry]uint
	TubeGeometry_orderStaged     map[uint]*TubeGeometry
	TubeGeometrys_reference      map[*TubeGeometry]*TubeGeometry
	TubeGeometrys_referenceOrder map[*TubeGeometry]uint

	// insertion point for slice of pointers maps
	OnAfterTubeGeometryCreateCallback GongOnAfterCreateInterface[TubeGeometry]
	OnAfterTubeGeometryUpdateCallback GongOnAfterUpdateInterface[TubeGeometry]
	OnAfterTubeGeometryDeleteCallback GongOnAfterDeleteInterface[TubeGeometry]

	Vector2s                map[*Vector2]struct{}
	Vector2s_instance       map[*Vector2]*Vector2
	Vector2s_mapString      map[string]*Vector2
	Vector2Order            uint
	Vector2_stagedOrder     map[*Vector2]uint
	Vector2_orderStaged     map[uint]*Vector2
	Vector2s_reference      map[*Vector2]*Vector2
	Vector2s_referenceOrder map[*Vector2]uint

	// insertion point for slice of pointers maps
	OnAfterVector2CreateCallback GongOnAfterCreateInterface[Vector2]
	OnAfterVector2UpdateCallback GongOnAfterUpdateInterface[Vector2]
	OnAfterVector2DeleteCallback GongOnAfterDeleteInterface[Vector2]

	Vector3s                map[*Vector3]struct{}
	Vector3s_instance       map[*Vector3]*Vector3
	Vector3s_mapString      map[string]*Vector3
	Vector3Order            uint
	Vector3_stagedOrder     map[*Vector3]uint
	Vector3_orderStaged     map[uint]*Vector3
	Vector3s_reference      map[*Vector3]*Vector3
	Vector3s_referenceOrder map[*Vector3]uint

	// insertion point for slice of pointers maps
	OnAfterVector3CreateCallback GongOnAfterCreateInterface[Vector3]
	OnAfterVector3UpdateCallback GongOnAfterUpdateInterface[Vector3]
	OnAfterVector3DeleteCallback GongOnAfterDeleteInterface[Vector3]

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
	__gong__clearReferences(&stage.AmbiantLights_reference, &stage.AmbiantLights_instance, &stage.AmbiantLights_referenceOrder)

	__gong__clearReferences(&stage.BoxGeometrys_reference, &stage.BoxGeometrys_instance, &stage.BoxGeometrys_referenceOrder)

	__gong__clearReferences(&stage.BufferGeometrys_reference, &stage.BufferGeometrys_instance, &stage.BufferGeometrys_referenceOrder)

	__gong__clearReferences(&stage.Cameras_reference, &stage.Cameras_instance, &stage.Cameras_referenceOrder)

	__gong__clearReferences(&stage.Canvass_reference, &stage.Canvass_instance, &stage.Canvass_referenceOrder)

	__gong__clearReferences(&stage.Curves_reference, &stage.Curves_instance, &stage.Curves_referenceOrder)

	__gong__clearReferences(&stage.CylinderGeometrys_reference, &stage.CylinderGeometrys_instance, &stage.CylinderGeometrys_referenceOrder)

	__gong__clearReferences(&stage.DirectionalLights_reference, &stage.DirectionalLights_instance, &stage.DirectionalLights_referenceOrder)

	__gong__clearReferences(&stage.ExtrudeGeometrys_reference, &stage.ExtrudeGeometrys_instance, &stage.ExtrudeGeometrys_referenceOrder)

	__gong__clearReferences(&stage.Meshs_reference, &stage.Meshs_instance, &stage.Meshs_referenceOrder)

	__gong__clearReferences(&stage.MeshMaterialBasics_reference, &stage.MeshMaterialBasics_instance, &stage.MeshMaterialBasics_referenceOrder)

	__gong__clearReferences(&stage.MeshPhysicalMaterials_reference, &stage.MeshPhysicalMaterials_instance, &stage.MeshPhysicalMaterials_referenceOrder)

	__gong__clearReferences(&stage.PlaneGeometrys_reference, &stage.PlaneGeometrys_instance, &stage.PlaneGeometrys_referenceOrder)

	__gong__clearReferences(&stage.Shapes_reference, &stage.Shapes_instance, &stage.Shapes_referenceOrder)

	__gong__clearReferences(&stage.SphereGeometrys_reference, &stage.SphereGeometrys_instance, &stage.SphereGeometrys_referenceOrder)

	__gong__clearReferences(&stage.TorusGeometrys_reference, &stage.TorusGeometrys_instance, &stage.TorusGeometrys_referenceOrder)

	__gong__clearReferences(&stage.Triangles_reference, &stage.Triangles_instance, &stage.Triangles_referenceOrder)

	__gong__clearReferences(&stage.TubeGeometrys_reference, &stage.TubeGeometrys_instance, &stage.TubeGeometrys_referenceOrder)

	__gong__clearReferences(&stage.Vector2s_reference, &stage.Vector2s_instance, &stage.Vector2s_referenceOrder)

	__gong__clearReferences(&stage.Vector3s_reference, &stage.Vector3s_instance, &stage.Vector3s_referenceOrder)

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
	stage.AmbiantLightOrder = __gong__recomputeOrder(stage.AmbiantLight_stagedOrder)

	stage.BoxGeometryOrder = __gong__recomputeOrder(stage.BoxGeometry_stagedOrder)

	stage.BufferGeometryOrder = __gong__recomputeOrder(stage.BufferGeometry_stagedOrder)

	stage.CameraOrder = __gong__recomputeOrder(stage.Camera_stagedOrder)

	stage.CanvasOrder = __gong__recomputeOrder(stage.Canvas_stagedOrder)

	stage.CurveOrder = __gong__recomputeOrder(stage.Curve_stagedOrder)

	stage.CylinderGeometryOrder = __gong__recomputeOrder(stage.CylinderGeometry_stagedOrder)

	stage.DirectionalLightOrder = __gong__recomputeOrder(stage.DirectionalLight_stagedOrder)

	stage.ExtrudeGeometryOrder = __gong__recomputeOrder(stage.ExtrudeGeometry_stagedOrder)

	stage.MeshOrder = __gong__recomputeOrder(stage.Mesh_stagedOrder)

	stage.MeshMaterialBasicOrder = __gong__recomputeOrder(stage.MeshMaterialBasic_stagedOrder)

	stage.MeshPhysicalMaterialOrder = __gong__recomputeOrder(stage.MeshPhysicalMaterial_stagedOrder)

	stage.PlaneGeometryOrder = __gong__recomputeOrder(stage.PlaneGeometry_stagedOrder)

	stage.ShapeOrder = __gong__recomputeOrder(stage.Shape_stagedOrder)

	stage.SphereGeometryOrder = __gong__recomputeOrder(stage.SphereGeometry_stagedOrder)

	stage.TorusGeometryOrder = __gong__recomputeOrder(stage.TorusGeometry_stagedOrder)

	stage.TriangleOrder = __gong__recomputeOrder(stage.Triangle_stagedOrder)

	stage.TubeGeometryOrder = __gong__recomputeOrder(stage.TubeGeometry_stagedOrder)

	stage.Vector2Order = __gong__recomputeOrder(stage.Vector2_stagedOrder)

	stage.Vector3Order = __gong__recomputeOrder(stage.Vector3_stagedOrder)

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
	case *AmbiantLight:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.AmbiantLights, stage.AmbiantLight_stagedOrder))
	case *BoxGeometry:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.BoxGeometrys, stage.BoxGeometry_stagedOrder))
	case *BufferGeometry:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.BufferGeometrys, stage.BufferGeometry_stagedOrder))
	case *Camera:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Cameras, stage.Camera_stagedOrder))
	case *Canvas:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Canvass, stage.Canvas_stagedOrder))
	case *Curve:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Curves, stage.Curve_stagedOrder))
	case *CylinderGeometry:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.CylinderGeometrys, stage.CylinderGeometry_stagedOrder))
	case *DirectionalLight:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.DirectionalLights, stage.DirectionalLight_stagedOrder))
	case *ExtrudeGeometry:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ExtrudeGeometrys, stage.ExtrudeGeometry_stagedOrder))
	case *Mesh:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Meshs, stage.Mesh_stagedOrder))
	case *MeshMaterialBasic:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.MeshMaterialBasics, stage.MeshMaterialBasic_stagedOrder))
	case *MeshPhysicalMaterial:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.MeshPhysicalMaterials, stage.MeshPhysicalMaterial_stagedOrder))
	case *PlaneGeometry:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.PlaneGeometrys, stage.PlaneGeometry_stagedOrder))
	case *Shape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Shapes, stage.Shape_stagedOrder))
	case *SphereGeometry:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.SphereGeometrys, stage.SphereGeometry_stagedOrder))
	case *TorusGeometry:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.TorusGeometrys, stage.TorusGeometry_stagedOrder))
	case *Triangle:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Triangles, stage.Triangle_stagedOrder))
	case *TubeGeometry:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.TubeGeometrys, stage.TubeGeometry_stagedOrder))
	case *Vector2:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Vector2s, stage.Vector2_stagedOrder))
	case *Vector3:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Vector3s, stage.Vector3_stagedOrder))

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
	return "github.com/fullstack-lang/gong/lib/threejs/go/models"
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
		AmbiantLights:           make(map[*AmbiantLight]struct{}),
		AmbiantLights_mapString: make(map[string]*AmbiantLight),

		BoxGeometrys:           make(map[*BoxGeometry]struct{}),
		BoxGeometrys_mapString: make(map[string]*BoxGeometry),

		BufferGeometrys:           make(map[*BufferGeometry]struct{}),
		BufferGeometrys_mapString: make(map[string]*BufferGeometry),

		Cameras:           make(map[*Camera]struct{}),
		Cameras_mapString: make(map[string]*Camera),

		Canvass:           make(map[*Canvas]struct{}),
		Canvass_mapString: make(map[string]*Canvas),

		Curves:           make(map[*Curve]struct{}),
		Curves_mapString: make(map[string]*Curve),

		CylinderGeometrys:           make(map[*CylinderGeometry]struct{}),
		CylinderGeometrys_mapString: make(map[string]*CylinderGeometry),

		DirectionalLights:           make(map[*DirectionalLight]struct{}),
		DirectionalLights_mapString: make(map[string]*DirectionalLight),

		ExtrudeGeometrys:           make(map[*ExtrudeGeometry]struct{}),
		ExtrudeGeometrys_mapString: make(map[string]*ExtrudeGeometry),

		Meshs:           make(map[*Mesh]struct{}),
		Meshs_mapString: make(map[string]*Mesh),

		MeshMaterialBasics:           make(map[*MeshMaterialBasic]struct{}),
		MeshMaterialBasics_mapString: make(map[string]*MeshMaterialBasic),

		MeshPhysicalMaterials:           make(map[*MeshPhysicalMaterial]struct{}),
		MeshPhysicalMaterials_mapString: make(map[string]*MeshPhysicalMaterial),

		PlaneGeometrys:           make(map[*PlaneGeometry]struct{}),
		PlaneGeometrys_mapString: make(map[string]*PlaneGeometry),

		Shapes:           make(map[*Shape]struct{}),
		Shapes_mapString: make(map[string]*Shape),

		SphereGeometrys:           make(map[*SphereGeometry]struct{}),
		SphereGeometrys_mapString: make(map[string]*SphereGeometry),

		TorusGeometrys:           make(map[*TorusGeometry]struct{}),
		TorusGeometrys_mapString: make(map[string]*TorusGeometry),

		Triangles:           make(map[*Triangle]struct{}),
		Triangles_mapString: make(map[string]*Triangle),

		TubeGeometrys:           make(map[*TubeGeometry]struct{}),
		TubeGeometrys_mapString: make(map[string]*TubeGeometry),

		Vector2s:           make(map[*Vector2]struct{}),
		Vector2s_mapString: make(map[string]*Vector2),

		Vector3s:           make(map[*Vector3]struct{}),
		Vector3s_mapString: make(map[string]*Vector3),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		AmbiantLight_stagedOrder: make(map[*AmbiantLight]uint),
		AmbiantLight_orderStaged: make(map[uint]*AmbiantLight),
		AmbiantLights_reference:  make(map[*AmbiantLight]*AmbiantLight),

		BoxGeometry_stagedOrder: make(map[*BoxGeometry]uint),
		BoxGeometry_orderStaged: make(map[uint]*BoxGeometry),
		BoxGeometrys_reference:  make(map[*BoxGeometry]*BoxGeometry),

		BufferGeometry_stagedOrder: make(map[*BufferGeometry]uint),
		BufferGeometry_orderStaged: make(map[uint]*BufferGeometry),
		BufferGeometrys_reference:  make(map[*BufferGeometry]*BufferGeometry),

		Camera_stagedOrder: make(map[*Camera]uint),
		Camera_orderStaged: make(map[uint]*Camera),
		Cameras_reference:  make(map[*Camera]*Camera),

		Canvas_stagedOrder: make(map[*Canvas]uint),
		Canvas_orderStaged: make(map[uint]*Canvas),
		Canvass_reference:  make(map[*Canvas]*Canvas),

		Curve_stagedOrder: make(map[*Curve]uint),
		Curve_orderStaged: make(map[uint]*Curve),
		Curves_reference:  make(map[*Curve]*Curve),

		CylinderGeometry_stagedOrder: make(map[*CylinderGeometry]uint),
		CylinderGeometry_orderStaged: make(map[uint]*CylinderGeometry),
		CylinderGeometrys_reference:  make(map[*CylinderGeometry]*CylinderGeometry),

		DirectionalLight_stagedOrder: make(map[*DirectionalLight]uint),
		DirectionalLight_orderStaged: make(map[uint]*DirectionalLight),
		DirectionalLights_reference:  make(map[*DirectionalLight]*DirectionalLight),

		ExtrudeGeometry_stagedOrder: make(map[*ExtrudeGeometry]uint),
		ExtrudeGeometry_orderStaged: make(map[uint]*ExtrudeGeometry),
		ExtrudeGeometrys_reference:  make(map[*ExtrudeGeometry]*ExtrudeGeometry),

		Mesh_stagedOrder: make(map[*Mesh]uint),
		Mesh_orderStaged: make(map[uint]*Mesh),
		Meshs_reference:  make(map[*Mesh]*Mesh),

		MeshMaterialBasic_stagedOrder: make(map[*MeshMaterialBasic]uint),
		MeshMaterialBasic_orderStaged: make(map[uint]*MeshMaterialBasic),
		MeshMaterialBasics_reference:  make(map[*MeshMaterialBasic]*MeshMaterialBasic),

		MeshPhysicalMaterial_stagedOrder: make(map[*MeshPhysicalMaterial]uint),
		MeshPhysicalMaterial_orderStaged: make(map[uint]*MeshPhysicalMaterial),
		MeshPhysicalMaterials_reference:  make(map[*MeshPhysicalMaterial]*MeshPhysicalMaterial),

		PlaneGeometry_stagedOrder: make(map[*PlaneGeometry]uint),
		PlaneGeometry_orderStaged: make(map[uint]*PlaneGeometry),
		PlaneGeometrys_reference:  make(map[*PlaneGeometry]*PlaneGeometry),

		Shape_stagedOrder: make(map[*Shape]uint),
		Shape_orderStaged: make(map[uint]*Shape),
		Shapes_reference:  make(map[*Shape]*Shape),

		SphereGeometry_stagedOrder: make(map[*SphereGeometry]uint),
		SphereGeometry_orderStaged: make(map[uint]*SphereGeometry),
		SphereGeometrys_reference:  make(map[*SphereGeometry]*SphereGeometry),

		TorusGeometry_stagedOrder: make(map[*TorusGeometry]uint),
		TorusGeometry_orderStaged: make(map[uint]*TorusGeometry),
		TorusGeometrys_reference:  make(map[*TorusGeometry]*TorusGeometry),

		Triangle_stagedOrder: make(map[*Triangle]uint),
		Triangle_orderStaged: make(map[uint]*Triangle),
		Triangles_reference:  make(map[*Triangle]*Triangle),

		TubeGeometry_stagedOrder: make(map[*TubeGeometry]uint),
		TubeGeometry_orderStaged: make(map[uint]*TubeGeometry),
		TubeGeometrys_reference:  make(map[*TubeGeometry]*TubeGeometry),

		Vector2_stagedOrder: make(map[*Vector2]uint),
		Vector2_orderStaged: make(map[uint]*Vector2),
		Vector2s_reference:  make(map[*Vector2]*Vector2),

		Vector3_stagedOrder: make(map[*Vector3]uint),
		Vector3_orderStaged: make(map[uint]*Vector3),
		Vector3s_reference:  make(map[*Vector3]*Vector3),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"AmbiantLight": &AmbiantLightUnmarshaller{},

			"BoxGeometry": &BoxGeometryUnmarshaller{},

			"BufferGeometry": &BufferGeometryUnmarshaller{},

			"Camera": &CameraUnmarshaller{},

			"Canvas": &CanvasUnmarshaller{},

			"Curve": &CurveUnmarshaller{},

			"CylinderGeometry": &CylinderGeometryUnmarshaller{},

			"DirectionalLight": &DirectionalLightUnmarshaller{},

			"ExtrudeGeometry": &ExtrudeGeometryUnmarshaller{},

			"Mesh": &MeshUnmarshaller{},

			"MeshMaterialBasic": &MeshMaterialBasicUnmarshaller{},

			"MeshPhysicalMaterial": &MeshPhysicalMaterialUnmarshaller{},

			"PlaneGeometry": &PlaneGeometryUnmarshaller{},

			"Shape": &ShapeUnmarshaller{},

			"SphereGeometry": &SphereGeometryUnmarshaller{},

			"TorusGeometry": &TorusGeometryUnmarshaller{},

			"Triangle": &TriangleUnmarshaller{},

			"TubeGeometry": &TubeGeometryUnmarshaller{},

			"Vector2": &Vector2Unmarshaller{},

			"Vector3": &Vector3Unmarshaller{},

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
	case *AmbiantLight:
		return any(stage.AmbiantLight_orderStaged[order]).(Type)
	case *BoxGeometry:
		return any(stage.BoxGeometry_orderStaged[order]).(Type)
	case *BufferGeometry:
		return any(stage.BufferGeometry_orderStaged[order]).(Type)
	case *Camera:
		return any(stage.Camera_orderStaged[order]).(Type)
	case *Canvas:
		return any(stage.Canvas_orderStaged[order]).(Type)
	case *Curve:
		return any(stage.Curve_orderStaged[order]).(Type)
	case *CylinderGeometry:
		return any(stage.CylinderGeometry_orderStaged[order]).(Type)
	case *DirectionalLight:
		return any(stage.DirectionalLight_orderStaged[order]).(Type)
	case *ExtrudeGeometry:
		return any(stage.ExtrudeGeometry_orderStaged[order]).(Type)
	case *Mesh:
		return any(stage.Mesh_orderStaged[order]).(Type)
	case *MeshMaterialBasic:
		return any(stage.MeshMaterialBasic_orderStaged[order]).(Type)
	case *MeshPhysicalMaterial:
		return any(stage.MeshPhysicalMaterial_orderStaged[order]).(Type)
	case *PlaneGeometry:
		return any(stage.PlaneGeometry_orderStaged[order]).(Type)
	case *Shape:
		return any(stage.Shape_orderStaged[order]).(Type)
	case *SphereGeometry:
		return any(stage.SphereGeometry_orderStaged[order]).(Type)
	case *TorusGeometry:
		return any(stage.TorusGeometry_orderStaged[order]).(Type)
	case *Triangle:
		return any(stage.Triangle_orderStaged[order]).(Type)
	case *TubeGeometry:
		return any(stage.TubeGeometry_orderStaged[order]).(Type)
	case *Vector2:
		return any(stage.Vector2_orderStaged[order]).(Type)
	case *Vector3:
		return any(stage.Vector3_orderStaged[order]).(Type)
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
	stage.Map_GongStructName_InstancesNb["AmbiantLight"] = len(stage.AmbiantLights)
	stage.Map_GongStructName_InstancesNb["BoxGeometry"] = len(stage.BoxGeometrys)
	stage.Map_GongStructName_InstancesNb["BufferGeometry"] = len(stage.BufferGeometrys)
	stage.Map_GongStructName_InstancesNb["Camera"] = len(stage.Cameras)
	stage.Map_GongStructName_InstancesNb["Canvas"] = len(stage.Canvass)
	stage.Map_GongStructName_InstancesNb["Curve"] = len(stage.Curves)
	stage.Map_GongStructName_InstancesNb["CylinderGeometry"] = len(stage.CylinderGeometrys)
	stage.Map_GongStructName_InstancesNb["DirectionalLight"] = len(stage.DirectionalLights)
	stage.Map_GongStructName_InstancesNb["ExtrudeGeometry"] = len(stage.ExtrudeGeometrys)
	stage.Map_GongStructName_InstancesNb["Mesh"] = len(stage.Meshs)
	stage.Map_GongStructName_InstancesNb["MeshMaterialBasic"] = len(stage.MeshMaterialBasics)
	stage.Map_GongStructName_InstancesNb["MeshPhysicalMaterial"] = len(stage.MeshPhysicalMaterials)
	stage.Map_GongStructName_InstancesNb["PlaneGeometry"] = len(stage.PlaneGeometrys)
	stage.Map_GongStructName_InstancesNb["Shape"] = len(stage.Shapes)
	stage.Map_GongStructName_InstancesNb["SphereGeometry"] = len(stage.SphereGeometrys)
	stage.Map_GongStructName_InstancesNb["TorusGeometry"] = len(stage.TorusGeometrys)
	stage.Map_GongStructName_InstancesNb["Triangle"] = len(stage.Triangles)
	stage.Map_GongStructName_InstancesNb["TubeGeometry"] = len(stage.TubeGeometrys)
	stage.Map_GongStructName_InstancesNb["Vector2"] = len(stage.Vector2s)
	stage.Map_GongStructName_InstancesNb["Vector3"] = len(stage.Vector3s)
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
// Stage puts ambiantlight to the model stage
func (ambiantlight *AmbiantLight) Stage(stage *Stage) *AmbiantLight {
	__gong__stage(stage.AmbiantLights, stage.AmbiantLight_stagedOrder, stage.AmbiantLight_orderStaged, &stage.AmbiantLightOrder, stage.AmbiantLights_mapString, ambiantlight, ambiantlight.Name)
	return ambiantlight
}

// StagePreserveOrder puts ambiantlight to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AmbiantLightOrder
// - update stage.AmbiantLightOrder accordingly
func (ambiantlight *AmbiantLight) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AmbiantLights, stage.AmbiantLight_stagedOrder, stage.AmbiantLight_orderStaged, &stage.AmbiantLightOrder, stage.AmbiantLights_mapString, ambiantlight, order, ambiantlight.Name)
}

// Unstage removes ambiantlight off the model stage
func (ambiantlight *AmbiantLight) Unstage(stage *Stage) *AmbiantLight {
	__gong__unstage(stage.AmbiantLights, stage.AmbiantLights_mapString, ambiantlight, ambiantlight.Name)
	return ambiantlight
}

// UnstageVoid removes ambiantlight off the model stage
func (ambiantlight *AmbiantLight) UnstageVoid(stage *Stage) {
	ambiantlight.Unstage(stage)
}

func (ambiantlight *AmbiantLight) StageVoid(stage *Stage) {
	ambiantlight.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ambiantlight *AmbiantLight) GetName() (res string) {
	return ambiantlight.Name
}

// for satisfaction of GongStruct interface
func (ambiantlight *AmbiantLight) SetName(name string) {
	ambiantlight.Name = name
}

// Stage puts boxgeometry to the model stage
func (boxgeometry *BoxGeometry) Stage(stage *Stage) *BoxGeometry {
	__gong__stage(stage.BoxGeometrys, stage.BoxGeometry_stagedOrder, stage.BoxGeometry_orderStaged, &stage.BoxGeometryOrder, stage.BoxGeometrys_mapString, boxgeometry, boxgeometry.Name)
	return boxgeometry
}

// StagePreserveOrder puts boxgeometry to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BoxGeometryOrder
// - update stage.BoxGeometryOrder accordingly
func (boxgeometry *BoxGeometry) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.BoxGeometrys, stage.BoxGeometry_stagedOrder, stage.BoxGeometry_orderStaged, &stage.BoxGeometryOrder, stage.BoxGeometrys_mapString, boxgeometry, order, boxgeometry.Name)
}

// Unstage removes boxgeometry off the model stage
func (boxgeometry *BoxGeometry) Unstage(stage *Stage) *BoxGeometry {
	__gong__unstage(stage.BoxGeometrys, stage.BoxGeometrys_mapString, boxgeometry, boxgeometry.Name)
	return boxgeometry
}

// UnstageVoid removes boxgeometry off the model stage
func (boxgeometry *BoxGeometry) UnstageVoid(stage *Stage) {
	boxgeometry.Unstage(stage)
}

func (boxgeometry *BoxGeometry) StageVoid(stage *Stage) {
	boxgeometry.Stage(stage)
}

// for satisfaction of GongStruct interface
func (boxgeometry *BoxGeometry) GetName() (res string) {
	return boxgeometry.Name
}

// for satisfaction of GongStruct interface
func (boxgeometry *BoxGeometry) SetName(name string) {
	boxgeometry.Name = name
}

// Stage puts buffergeometry to the model stage
func (buffergeometry *BufferGeometry) Stage(stage *Stage) *BufferGeometry {
	__gong__stage(stage.BufferGeometrys, stage.BufferGeometry_stagedOrder, stage.BufferGeometry_orderStaged, &stage.BufferGeometryOrder, stage.BufferGeometrys_mapString, buffergeometry, buffergeometry.Name)
	return buffergeometry
}

// StagePreserveOrder puts buffergeometry to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BufferGeometryOrder
// - update stage.BufferGeometryOrder accordingly
func (buffergeometry *BufferGeometry) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.BufferGeometrys, stage.BufferGeometry_stagedOrder, stage.BufferGeometry_orderStaged, &stage.BufferGeometryOrder, stage.BufferGeometrys_mapString, buffergeometry, order, buffergeometry.Name)
}

// Unstage removes buffergeometry off the model stage
func (buffergeometry *BufferGeometry) Unstage(stage *Stage) *BufferGeometry {
	__gong__unstage(stage.BufferGeometrys, stage.BufferGeometrys_mapString, buffergeometry, buffergeometry.Name)
	return buffergeometry
}

// UnstageVoid removes buffergeometry off the model stage
func (buffergeometry *BufferGeometry) UnstageVoid(stage *Stage) {
	buffergeometry.Unstage(stage)
}

func (buffergeometry *BufferGeometry) StageVoid(stage *Stage) {
	buffergeometry.Stage(stage)
}

// for satisfaction of GongStruct interface
func (buffergeometry *BufferGeometry) GetName() (res string) {
	return buffergeometry.Name
}

// for satisfaction of GongStruct interface
func (buffergeometry *BufferGeometry) SetName(name string) {
	buffergeometry.Name = name
}

// Stage puts camera to the model stage
func (camera *Camera) Stage(stage *Stage) *Camera {
	__gong__stage(stage.Cameras, stage.Camera_stagedOrder, stage.Camera_orderStaged, &stage.CameraOrder, stage.Cameras_mapString, camera, camera.Name)
	return camera
}

// StagePreserveOrder puts camera to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CameraOrder
// - update stage.CameraOrder accordingly
func (camera *Camera) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Cameras, stage.Camera_stagedOrder, stage.Camera_orderStaged, &stage.CameraOrder, stage.Cameras_mapString, camera, order, camera.Name)
}

// Unstage removes camera off the model stage
func (camera *Camera) Unstage(stage *Stage) *Camera {
	__gong__unstage(stage.Cameras, stage.Cameras_mapString, camera, camera.Name)
	return camera
}

// UnstageVoid removes camera off the model stage
func (camera *Camera) UnstageVoid(stage *Stage) {
	camera.Unstage(stage)
}

func (camera *Camera) StageVoid(stage *Stage) {
	camera.Stage(stage)
}

// for satisfaction of GongStruct interface
func (camera *Camera) GetName() (res string) {
	return camera.Name
}

// for satisfaction of GongStruct interface
func (camera *Camera) SetName(name string) {
	camera.Name = name
}

// Stage puts canvas to the model stage
func (canvas *Canvas) Stage(stage *Stage) *Canvas {
	__gong__stage(stage.Canvass, stage.Canvas_stagedOrder, stage.Canvas_orderStaged, &stage.CanvasOrder, stage.Canvass_mapString, canvas, canvas.Name)
	return canvas
}

// StagePreserveOrder puts canvas to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CanvasOrder
// - update stage.CanvasOrder accordingly
func (canvas *Canvas) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Canvass, stage.Canvas_stagedOrder, stage.Canvas_orderStaged, &stage.CanvasOrder, stage.Canvass_mapString, canvas, order, canvas.Name)
}

// Unstage removes canvas off the model stage
func (canvas *Canvas) Unstage(stage *Stage) *Canvas {
	__gong__unstage(stage.Canvass, stage.Canvass_mapString, canvas, canvas.Name)
	return canvas
}

// UnstageVoid removes canvas off the model stage
func (canvas *Canvas) UnstageVoid(stage *Stage) {
	canvas.Unstage(stage)
}

func (canvas *Canvas) StageVoid(stage *Stage) {
	canvas.Stage(stage)
}

// for satisfaction of GongStruct interface
func (canvas *Canvas) GetName() (res string) {
	return canvas.Name
}

// for satisfaction of GongStruct interface
func (canvas *Canvas) SetName(name string) {
	canvas.Name = name
}

// Stage puts curve to the model stage
func (curve *Curve) Stage(stage *Stage) *Curve {
	__gong__stage(stage.Curves, stage.Curve_stagedOrder, stage.Curve_orderStaged, &stage.CurveOrder, stage.Curves_mapString, curve, curve.Name)
	return curve
}

// StagePreserveOrder puts curve to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CurveOrder
// - update stage.CurveOrder accordingly
func (curve *Curve) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Curves, stage.Curve_stagedOrder, stage.Curve_orderStaged, &stage.CurveOrder, stage.Curves_mapString, curve, order, curve.Name)
}

// Unstage removes curve off the model stage
func (curve *Curve) Unstage(stage *Stage) *Curve {
	__gong__unstage(stage.Curves, stage.Curves_mapString, curve, curve.Name)
	return curve
}

// UnstageVoid removes curve off the model stage
func (curve *Curve) UnstageVoid(stage *Stage) {
	curve.Unstage(stage)
}

func (curve *Curve) StageVoid(stage *Stage) {
	curve.Stage(stage)
}

// for satisfaction of GongStruct interface
func (curve *Curve) GetName() (res string) {
	return curve.Name
}

// for satisfaction of GongStruct interface
func (curve *Curve) SetName(name string) {
	curve.Name = name
}

// Stage puts cylindergeometry to the model stage
func (cylindergeometry *CylinderGeometry) Stage(stage *Stage) *CylinderGeometry {
	__gong__stage(stage.CylinderGeometrys, stage.CylinderGeometry_stagedOrder, stage.CylinderGeometry_orderStaged, &stage.CylinderGeometryOrder, stage.CylinderGeometrys_mapString, cylindergeometry, cylindergeometry.Name)
	return cylindergeometry
}

// StagePreserveOrder puts cylindergeometry to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CylinderGeometryOrder
// - update stage.CylinderGeometryOrder accordingly
func (cylindergeometry *CylinderGeometry) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CylinderGeometrys, stage.CylinderGeometry_stagedOrder, stage.CylinderGeometry_orderStaged, &stage.CylinderGeometryOrder, stage.CylinderGeometrys_mapString, cylindergeometry, order, cylindergeometry.Name)
}

// Unstage removes cylindergeometry off the model stage
func (cylindergeometry *CylinderGeometry) Unstage(stage *Stage) *CylinderGeometry {
	__gong__unstage(stage.CylinderGeometrys, stage.CylinderGeometrys_mapString, cylindergeometry, cylindergeometry.Name)
	return cylindergeometry
}

// UnstageVoid removes cylindergeometry off the model stage
func (cylindergeometry *CylinderGeometry) UnstageVoid(stage *Stage) {
	cylindergeometry.Unstage(stage)
}

func (cylindergeometry *CylinderGeometry) StageVoid(stage *Stage) {
	cylindergeometry.Stage(stage)
}

// for satisfaction of GongStruct interface
func (cylindergeometry *CylinderGeometry) GetName() (res string) {
	return cylindergeometry.Name
}

// for satisfaction of GongStruct interface
func (cylindergeometry *CylinderGeometry) SetName(name string) {
	cylindergeometry.Name = name
}

// Stage puts directionallight to the model stage
func (directionallight *DirectionalLight) Stage(stage *Stage) *DirectionalLight {
	__gong__stage(stage.DirectionalLights, stage.DirectionalLight_stagedOrder, stage.DirectionalLight_orderStaged, &stage.DirectionalLightOrder, stage.DirectionalLights_mapString, directionallight, directionallight.Name)
	return directionallight
}

// StagePreserveOrder puts directionallight to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DirectionalLightOrder
// - update stage.DirectionalLightOrder accordingly
func (directionallight *DirectionalLight) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DirectionalLights, stage.DirectionalLight_stagedOrder, stage.DirectionalLight_orderStaged, &stage.DirectionalLightOrder, stage.DirectionalLights_mapString, directionallight, order, directionallight.Name)
}

// Unstage removes directionallight off the model stage
func (directionallight *DirectionalLight) Unstage(stage *Stage) *DirectionalLight {
	__gong__unstage(stage.DirectionalLights, stage.DirectionalLights_mapString, directionallight, directionallight.Name)
	return directionallight
}

// UnstageVoid removes directionallight off the model stage
func (directionallight *DirectionalLight) UnstageVoid(stage *Stage) {
	directionallight.Unstage(stage)
}

func (directionallight *DirectionalLight) StageVoid(stage *Stage) {
	directionallight.Stage(stage)
}

// for satisfaction of GongStruct interface
func (directionallight *DirectionalLight) GetName() (res string) {
	return directionallight.Name
}

// for satisfaction of GongStruct interface
func (directionallight *DirectionalLight) SetName(name string) {
	directionallight.Name = name
}

// Stage puts extrudegeometry to the model stage
func (extrudegeometry *ExtrudeGeometry) Stage(stage *Stage) *ExtrudeGeometry {
	__gong__stage(stage.ExtrudeGeometrys, stage.ExtrudeGeometry_stagedOrder, stage.ExtrudeGeometry_orderStaged, &stage.ExtrudeGeometryOrder, stage.ExtrudeGeometrys_mapString, extrudegeometry, extrudegeometry.Name)
	return extrudegeometry
}

// StagePreserveOrder puts extrudegeometry to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ExtrudeGeometryOrder
// - update stage.ExtrudeGeometryOrder accordingly
func (extrudegeometry *ExtrudeGeometry) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ExtrudeGeometrys, stage.ExtrudeGeometry_stagedOrder, stage.ExtrudeGeometry_orderStaged, &stage.ExtrudeGeometryOrder, stage.ExtrudeGeometrys_mapString, extrudegeometry, order, extrudegeometry.Name)
}

// Unstage removes extrudegeometry off the model stage
func (extrudegeometry *ExtrudeGeometry) Unstage(stage *Stage) *ExtrudeGeometry {
	__gong__unstage(stage.ExtrudeGeometrys, stage.ExtrudeGeometrys_mapString, extrudegeometry, extrudegeometry.Name)
	return extrudegeometry
}

// UnstageVoid removes extrudegeometry off the model stage
func (extrudegeometry *ExtrudeGeometry) UnstageVoid(stage *Stage) {
	extrudegeometry.Unstage(stage)
}

func (extrudegeometry *ExtrudeGeometry) StageVoid(stage *Stage) {
	extrudegeometry.Stage(stage)
}

// for satisfaction of GongStruct interface
func (extrudegeometry *ExtrudeGeometry) GetName() (res string) {
	return extrudegeometry.Name
}

// for satisfaction of GongStruct interface
func (extrudegeometry *ExtrudeGeometry) SetName(name string) {
	extrudegeometry.Name = name
}

// Stage puts mesh to the model stage
func (mesh *Mesh) Stage(stage *Stage) *Mesh {
	__gong__stage(stage.Meshs, stage.Mesh_stagedOrder, stage.Mesh_orderStaged, &stage.MeshOrder, stage.Meshs_mapString, mesh, mesh.Name)
	return mesh
}

// StagePreserveOrder puts mesh to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MeshOrder
// - update stage.MeshOrder accordingly
func (mesh *Mesh) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Meshs, stage.Mesh_stagedOrder, stage.Mesh_orderStaged, &stage.MeshOrder, stage.Meshs_mapString, mesh, order, mesh.Name)
}

// Unstage removes mesh off the model stage
func (mesh *Mesh) Unstage(stage *Stage) *Mesh {
	__gong__unstage(stage.Meshs, stage.Meshs_mapString, mesh, mesh.Name)
	return mesh
}

// UnstageVoid removes mesh off the model stage
func (mesh *Mesh) UnstageVoid(stage *Stage) {
	mesh.Unstage(stage)
}

func (mesh *Mesh) StageVoid(stage *Stage) {
	mesh.Stage(stage)
}

// for satisfaction of GongStruct interface
func (mesh *Mesh) GetName() (res string) {
	return mesh.Name
}

// for satisfaction of GongStruct interface
func (mesh *Mesh) SetName(name string) {
	mesh.Name = name
}

// Stage puts meshmaterialbasic to the model stage
func (meshmaterialbasic *MeshMaterialBasic) Stage(stage *Stage) *MeshMaterialBasic {
	__gong__stage(stage.MeshMaterialBasics, stage.MeshMaterialBasic_stagedOrder, stage.MeshMaterialBasic_orderStaged, &stage.MeshMaterialBasicOrder, stage.MeshMaterialBasics_mapString, meshmaterialbasic, meshmaterialbasic.Name)
	return meshmaterialbasic
}

// StagePreserveOrder puts meshmaterialbasic to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MeshMaterialBasicOrder
// - update stage.MeshMaterialBasicOrder accordingly
func (meshmaterialbasic *MeshMaterialBasic) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MeshMaterialBasics, stage.MeshMaterialBasic_stagedOrder, stage.MeshMaterialBasic_orderStaged, &stage.MeshMaterialBasicOrder, stage.MeshMaterialBasics_mapString, meshmaterialbasic, order, meshmaterialbasic.Name)
}

// Unstage removes meshmaterialbasic off the model stage
func (meshmaterialbasic *MeshMaterialBasic) Unstage(stage *Stage) *MeshMaterialBasic {
	__gong__unstage(stage.MeshMaterialBasics, stage.MeshMaterialBasics_mapString, meshmaterialbasic, meshmaterialbasic.Name)
	return meshmaterialbasic
}

// UnstageVoid removes meshmaterialbasic off the model stage
func (meshmaterialbasic *MeshMaterialBasic) UnstageVoid(stage *Stage) {
	meshmaterialbasic.Unstage(stage)
}

func (meshmaterialbasic *MeshMaterialBasic) StageVoid(stage *Stage) {
	meshmaterialbasic.Stage(stage)
}

// for satisfaction of GongStruct interface
func (meshmaterialbasic *MeshMaterialBasic) GetName() (res string) {
	return meshmaterialbasic.Name
}

// for satisfaction of GongStruct interface
func (meshmaterialbasic *MeshMaterialBasic) SetName(name string) {
	meshmaterialbasic.Name = name
}

// Stage puts meshphysicalmaterial to the model stage
func (meshphysicalmaterial *MeshPhysicalMaterial) Stage(stage *Stage) *MeshPhysicalMaterial {
	__gong__stage(stage.MeshPhysicalMaterials, stage.MeshPhysicalMaterial_stagedOrder, stage.MeshPhysicalMaterial_orderStaged, &stage.MeshPhysicalMaterialOrder, stage.MeshPhysicalMaterials_mapString, meshphysicalmaterial, meshphysicalmaterial.Name)
	return meshphysicalmaterial
}

// StagePreserveOrder puts meshphysicalmaterial to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MeshPhysicalMaterialOrder
// - update stage.MeshPhysicalMaterialOrder accordingly
func (meshphysicalmaterial *MeshPhysicalMaterial) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MeshPhysicalMaterials, stage.MeshPhysicalMaterial_stagedOrder, stage.MeshPhysicalMaterial_orderStaged, &stage.MeshPhysicalMaterialOrder, stage.MeshPhysicalMaterials_mapString, meshphysicalmaterial, order, meshphysicalmaterial.Name)
}

// Unstage removes meshphysicalmaterial off the model stage
func (meshphysicalmaterial *MeshPhysicalMaterial) Unstage(stage *Stage) *MeshPhysicalMaterial {
	__gong__unstage(stage.MeshPhysicalMaterials, stage.MeshPhysicalMaterials_mapString, meshphysicalmaterial, meshphysicalmaterial.Name)
	return meshphysicalmaterial
}

// UnstageVoid removes meshphysicalmaterial off the model stage
func (meshphysicalmaterial *MeshPhysicalMaterial) UnstageVoid(stage *Stage) {
	meshphysicalmaterial.Unstage(stage)
}

func (meshphysicalmaterial *MeshPhysicalMaterial) StageVoid(stage *Stage) {
	meshphysicalmaterial.Stage(stage)
}

// for satisfaction of GongStruct interface
func (meshphysicalmaterial *MeshPhysicalMaterial) GetName() (res string) {
	return meshphysicalmaterial.Name
}

// for satisfaction of GongStruct interface
func (meshphysicalmaterial *MeshPhysicalMaterial) SetName(name string) {
	meshphysicalmaterial.Name = name
}

// Stage puts planegeometry to the model stage
func (planegeometry *PlaneGeometry) Stage(stage *Stage) *PlaneGeometry {
	__gong__stage(stage.PlaneGeometrys, stage.PlaneGeometry_stagedOrder, stage.PlaneGeometry_orderStaged, &stage.PlaneGeometryOrder, stage.PlaneGeometrys_mapString, planegeometry, planegeometry.Name)
	return planegeometry
}

// StagePreserveOrder puts planegeometry to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PlaneGeometryOrder
// - update stage.PlaneGeometryOrder accordingly
func (planegeometry *PlaneGeometry) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.PlaneGeometrys, stage.PlaneGeometry_stagedOrder, stage.PlaneGeometry_orderStaged, &stage.PlaneGeometryOrder, stage.PlaneGeometrys_mapString, planegeometry, order, planegeometry.Name)
}

// Unstage removes planegeometry off the model stage
func (planegeometry *PlaneGeometry) Unstage(stage *Stage) *PlaneGeometry {
	__gong__unstage(stage.PlaneGeometrys, stage.PlaneGeometrys_mapString, planegeometry, planegeometry.Name)
	return planegeometry
}

// UnstageVoid removes planegeometry off the model stage
func (planegeometry *PlaneGeometry) UnstageVoid(stage *Stage) {
	planegeometry.Unstage(stage)
}

func (planegeometry *PlaneGeometry) StageVoid(stage *Stage) {
	planegeometry.Stage(stage)
}

// for satisfaction of GongStruct interface
func (planegeometry *PlaneGeometry) GetName() (res string) {
	return planegeometry.Name
}

// for satisfaction of GongStruct interface
func (planegeometry *PlaneGeometry) SetName(name string) {
	planegeometry.Name = name
}

// Stage puts shape to the model stage
func (shape *Shape) Stage(stage *Stage) *Shape {
	__gong__stage(stage.Shapes, stage.Shape_stagedOrder, stage.Shape_orderStaged, &stage.ShapeOrder, stage.Shapes_mapString, shape, shape.Name)
	return shape
}

// StagePreserveOrder puts shape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ShapeOrder
// - update stage.ShapeOrder accordingly
func (shape *Shape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Shapes, stage.Shape_stagedOrder, stage.Shape_orderStaged, &stage.ShapeOrder, stage.Shapes_mapString, shape, order, shape.Name)
}

// Unstage removes shape off the model stage
func (shape *Shape) Unstage(stage *Stage) *Shape {
	__gong__unstage(stage.Shapes, stage.Shapes_mapString, shape, shape.Name)
	return shape
}

// UnstageVoid removes shape off the model stage
func (shape *Shape) UnstageVoid(stage *Stage) {
	shape.Unstage(stage)
}

func (shape *Shape) StageVoid(stage *Stage) {
	shape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (shape *Shape) GetName() (res string) {
	return shape.Name
}

// for satisfaction of GongStruct interface
func (shape *Shape) SetName(name string) {
	shape.Name = name
}

// Stage puts spheregeometry to the model stage
func (spheregeometry *SphereGeometry) Stage(stage *Stage) *SphereGeometry {
	__gong__stage(stage.SphereGeometrys, stage.SphereGeometry_stagedOrder, stage.SphereGeometry_orderStaged, &stage.SphereGeometryOrder, stage.SphereGeometrys_mapString, spheregeometry, spheregeometry.Name)
	return spheregeometry
}

// StagePreserveOrder puts spheregeometry to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SphereGeometryOrder
// - update stage.SphereGeometryOrder accordingly
func (spheregeometry *SphereGeometry) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.SphereGeometrys, stage.SphereGeometry_stagedOrder, stage.SphereGeometry_orderStaged, &stage.SphereGeometryOrder, stage.SphereGeometrys_mapString, spheregeometry, order, spheregeometry.Name)
}

// Unstage removes spheregeometry off the model stage
func (spheregeometry *SphereGeometry) Unstage(stage *Stage) *SphereGeometry {
	__gong__unstage(stage.SphereGeometrys, stage.SphereGeometrys_mapString, spheregeometry, spheregeometry.Name)
	return spheregeometry
}

// UnstageVoid removes spheregeometry off the model stage
func (spheregeometry *SphereGeometry) UnstageVoid(stage *Stage) {
	spheregeometry.Unstage(stage)
}

func (spheregeometry *SphereGeometry) StageVoid(stage *Stage) {
	spheregeometry.Stage(stage)
}

// for satisfaction of GongStruct interface
func (spheregeometry *SphereGeometry) GetName() (res string) {
	return spheregeometry.Name
}

// for satisfaction of GongStruct interface
func (spheregeometry *SphereGeometry) SetName(name string) {
	spheregeometry.Name = name
}

// Stage puts torusgeometry to the model stage
func (torusgeometry *TorusGeometry) Stage(stage *Stage) *TorusGeometry {
	__gong__stage(stage.TorusGeometrys, stage.TorusGeometry_stagedOrder, stage.TorusGeometry_orderStaged, &stage.TorusGeometryOrder, stage.TorusGeometrys_mapString, torusgeometry, torusgeometry.Name)
	return torusgeometry
}

// StagePreserveOrder puts torusgeometry to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TorusGeometryOrder
// - update stage.TorusGeometryOrder accordingly
func (torusgeometry *TorusGeometry) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TorusGeometrys, stage.TorusGeometry_stagedOrder, stage.TorusGeometry_orderStaged, &stage.TorusGeometryOrder, stage.TorusGeometrys_mapString, torusgeometry, order, torusgeometry.Name)
}

// Unstage removes torusgeometry off the model stage
func (torusgeometry *TorusGeometry) Unstage(stage *Stage) *TorusGeometry {
	__gong__unstage(stage.TorusGeometrys, stage.TorusGeometrys_mapString, torusgeometry, torusgeometry.Name)
	return torusgeometry
}

// UnstageVoid removes torusgeometry off the model stage
func (torusgeometry *TorusGeometry) UnstageVoid(stage *Stage) {
	torusgeometry.Unstage(stage)
}

func (torusgeometry *TorusGeometry) StageVoid(stage *Stage) {
	torusgeometry.Stage(stage)
}

// for satisfaction of GongStruct interface
func (torusgeometry *TorusGeometry) GetName() (res string) {
	return torusgeometry.Name
}

// for satisfaction of GongStruct interface
func (torusgeometry *TorusGeometry) SetName(name string) {
	torusgeometry.Name = name
}

// Stage puts triangle to the model stage
func (triangle *Triangle) Stage(stage *Stage) *Triangle {
	__gong__stage(stage.Triangles, stage.Triangle_stagedOrder, stage.Triangle_orderStaged, &stage.TriangleOrder, stage.Triangles_mapString, triangle, triangle.Name)
	return triangle
}

// StagePreserveOrder puts triangle to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TriangleOrder
// - update stage.TriangleOrder accordingly
func (triangle *Triangle) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Triangles, stage.Triangle_stagedOrder, stage.Triangle_orderStaged, &stage.TriangleOrder, stage.Triangles_mapString, triangle, order, triangle.Name)
}

// Unstage removes triangle off the model stage
func (triangle *Triangle) Unstage(stage *Stage) *Triangle {
	__gong__unstage(stage.Triangles, stage.Triangles_mapString, triangle, triangle.Name)
	return triangle
}

// UnstageVoid removes triangle off the model stage
func (triangle *Triangle) UnstageVoid(stage *Stage) {
	triangle.Unstage(stage)
}

func (triangle *Triangle) StageVoid(stage *Stage) {
	triangle.Stage(stage)
}

// for satisfaction of GongStruct interface
func (triangle *Triangle) GetName() (res string) {
	return triangle.Name
}

// for satisfaction of GongStruct interface
func (triangle *Triangle) SetName(name string) {
	triangle.Name = name
}

// Stage puts tubegeometry to the model stage
func (tubegeometry *TubeGeometry) Stage(stage *Stage) *TubeGeometry {
	__gong__stage(stage.TubeGeometrys, stage.TubeGeometry_stagedOrder, stage.TubeGeometry_orderStaged, &stage.TubeGeometryOrder, stage.TubeGeometrys_mapString, tubegeometry, tubegeometry.Name)
	return tubegeometry
}

// StagePreserveOrder puts tubegeometry to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TubeGeometryOrder
// - update stage.TubeGeometryOrder accordingly
func (tubegeometry *TubeGeometry) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TubeGeometrys, stage.TubeGeometry_stagedOrder, stage.TubeGeometry_orderStaged, &stage.TubeGeometryOrder, stage.TubeGeometrys_mapString, tubegeometry, order, tubegeometry.Name)
}

// Unstage removes tubegeometry off the model stage
func (tubegeometry *TubeGeometry) Unstage(stage *Stage) *TubeGeometry {
	__gong__unstage(stage.TubeGeometrys, stage.TubeGeometrys_mapString, tubegeometry, tubegeometry.Name)
	return tubegeometry
}

// UnstageVoid removes tubegeometry off the model stage
func (tubegeometry *TubeGeometry) UnstageVoid(stage *Stage) {
	tubegeometry.Unstage(stage)
}

func (tubegeometry *TubeGeometry) StageVoid(stage *Stage) {
	tubegeometry.Stage(stage)
}

// for satisfaction of GongStruct interface
func (tubegeometry *TubeGeometry) GetName() (res string) {
	return tubegeometry.Name
}

// for satisfaction of GongStruct interface
func (tubegeometry *TubeGeometry) SetName(name string) {
	tubegeometry.Name = name
}

// Stage puts vector2 to the model stage
func (vector2 *Vector2) Stage(stage *Stage) *Vector2 {
	__gong__stage(stage.Vector2s, stage.Vector2_stagedOrder, stage.Vector2_orderStaged, &stage.Vector2Order, stage.Vector2s_mapString, vector2, vector2.Name)
	return vector2
}

// StagePreserveOrder puts vector2 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.Vector2Order
// - update stage.Vector2Order accordingly
func (vector2 *Vector2) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Vector2s, stage.Vector2_stagedOrder, stage.Vector2_orderStaged, &stage.Vector2Order, stage.Vector2s_mapString, vector2, order, vector2.Name)
}

// Unstage removes vector2 off the model stage
func (vector2 *Vector2) Unstage(stage *Stage) *Vector2 {
	__gong__unstage(stage.Vector2s, stage.Vector2s_mapString, vector2, vector2.Name)
	return vector2
}

// UnstageVoid removes vector2 off the model stage
func (vector2 *Vector2) UnstageVoid(stage *Stage) {
	vector2.Unstage(stage)
}

func (vector2 *Vector2) StageVoid(stage *Stage) {
	vector2.Stage(stage)
}

// for satisfaction of GongStruct interface
func (vector2 *Vector2) GetName() (res string) {
	return vector2.Name
}

// for satisfaction of GongStruct interface
func (vector2 *Vector2) SetName(name string) {
	vector2.Name = name
}

// Stage puts vector3 to the model stage
func (vector3 *Vector3) Stage(stage *Stage) *Vector3 {
	__gong__stage(stage.Vector3s, stage.Vector3_stagedOrder, stage.Vector3_orderStaged, &stage.Vector3Order, stage.Vector3s_mapString, vector3, vector3.Name)
	return vector3
}

// StagePreserveOrder puts vector3 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.Vector3Order
// - update stage.Vector3Order accordingly
func (vector3 *Vector3) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Vector3s, stage.Vector3_stagedOrder, stage.Vector3_orderStaged, &stage.Vector3Order, stage.Vector3s_mapString, vector3, order, vector3.Name)
}

// Unstage removes vector3 off the model stage
func (vector3 *Vector3) Unstage(stage *Stage) *Vector3 {
	__gong__unstage(stage.Vector3s, stage.Vector3s_mapString, vector3, vector3.Name)
	return vector3
}

// UnstageVoid removes vector3 off the model stage
func (vector3 *Vector3) UnstageVoid(stage *Stage) {
	vector3.Unstage(stage)
}

func (vector3 *Vector3) StageVoid(stage *Stage) {
	vector3.Stage(stage)
}

// for satisfaction of GongStruct interface
func (vector3 *Vector3) GetName() (res string) {
	return vector3.Name
}

// for satisfaction of GongStruct interface
func (vector3 *Vector3) SetName(name string) {
	vector3.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.AmbiantLights, &stage.AmbiantLights_mapString, &stage.AmbiantLight_stagedOrder, &stage.AmbiantLightOrder)

	__gong__resetStageType(&stage.BoxGeometrys, &stage.BoxGeometrys_mapString, &stage.BoxGeometry_stagedOrder, &stage.BoxGeometryOrder)

	__gong__resetStageType(&stage.BufferGeometrys, &stage.BufferGeometrys_mapString, &stage.BufferGeometry_stagedOrder, &stage.BufferGeometryOrder)

	__gong__resetStageType(&stage.Cameras, &stage.Cameras_mapString, &stage.Camera_stagedOrder, &stage.CameraOrder)

	__gong__resetStageType(&stage.Canvass, &stage.Canvass_mapString, &stage.Canvas_stagedOrder, &stage.CanvasOrder)

	__gong__resetStageType(&stage.Curves, &stage.Curves_mapString, &stage.Curve_stagedOrder, &stage.CurveOrder)

	__gong__resetStageType(&stage.CylinderGeometrys, &stage.CylinderGeometrys_mapString, &stage.CylinderGeometry_stagedOrder, &stage.CylinderGeometryOrder)

	__gong__resetStageType(&stage.DirectionalLights, &stage.DirectionalLights_mapString, &stage.DirectionalLight_stagedOrder, &stage.DirectionalLightOrder)

	__gong__resetStageType(&stage.ExtrudeGeometrys, &stage.ExtrudeGeometrys_mapString, &stage.ExtrudeGeometry_stagedOrder, &stage.ExtrudeGeometryOrder)

	__gong__resetStageType(&stage.Meshs, &stage.Meshs_mapString, &stage.Mesh_stagedOrder, &stage.MeshOrder)

	__gong__resetStageType(&stage.MeshMaterialBasics, &stage.MeshMaterialBasics_mapString, &stage.MeshMaterialBasic_stagedOrder, &stage.MeshMaterialBasicOrder)

	__gong__resetStageType(&stage.MeshPhysicalMaterials, &stage.MeshPhysicalMaterials_mapString, &stage.MeshPhysicalMaterial_stagedOrder, &stage.MeshPhysicalMaterialOrder)

	__gong__resetStageType(&stage.PlaneGeometrys, &stage.PlaneGeometrys_mapString, &stage.PlaneGeometry_stagedOrder, &stage.PlaneGeometryOrder)

	__gong__resetStageType(&stage.Shapes, &stage.Shapes_mapString, &stage.Shape_stagedOrder, &stage.ShapeOrder)

	__gong__resetStageType(&stage.SphereGeometrys, &stage.SphereGeometrys_mapString, &stage.SphereGeometry_stagedOrder, &stage.SphereGeometryOrder)

	__gong__resetStageType(&stage.TorusGeometrys, &stage.TorusGeometrys_mapString, &stage.TorusGeometry_stagedOrder, &stage.TorusGeometryOrder)

	__gong__resetStageType(&stage.Triangles, &stage.Triangles_mapString, &stage.Triangle_stagedOrder, &stage.TriangleOrder)

	__gong__resetStageType(&stage.TubeGeometrys, &stage.TubeGeometrys_mapString, &stage.TubeGeometry_stagedOrder, &stage.TubeGeometryOrder)

	__gong__resetStageType(&stage.Vector2s, &stage.Vector2s_mapString, &stage.Vector2_stagedOrder, &stage.Vector2Order)

	__gong__resetStageType(&stage.Vector3s, &stage.Vector3s_mapString, &stage.Vector3_stagedOrder, &stage.Vector3Order)

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
	case *AmbiantLight:
		return any(stage.AmbiantLights_mapString).(map[string]Type)
	case *BoxGeometry:
		return any(stage.BoxGeometrys_mapString).(map[string]Type)
	case *BufferGeometry:
		return any(stage.BufferGeometrys_mapString).(map[string]Type)
	case *Camera:
		return any(stage.Cameras_mapString).(map[string]Type)
	case *Canvas:
		return any(stage.Canvass_mapString).(map[string]Type)
	case *Curve:
		return any(stage.Curves_mapString).(map[string]Type)
	case *CylinderGeometry:
		return any(stage.CylinderGeometrys_mapString).(map[string]Type)
	case *DirectionalLight:
		return any(stage.DirectionalLights_mapString).(map[string]Type)
	case *ExtrudeGeometry:
		return any(stage.ExtrudeGeometrys_mapString).(map[string]Type)
	case *Mesh:
		return any(stage.Meshs_mapString).(map[string]Type)
	case *MeshMaterialBasic:
		return any(stage.MeshMaterialBasics_mapString).(map[string]Type)
	case *MeshPhysicalMaterial:
		return any(stage.MeshPhysicalMaterials_mapString).(map[string]Type)
	case *PlaneGeometry:
		return any(stage.PlaneGeometrys_mapString).(map[string]Type)
	case *Shape:
		return any(stage.Shapes_mapString).(map[string]Type)
	case *SphereGeometry:
		return any(stage.SphereGeometrys_mapString).(map[string]Type)
	case *TorusGeometry:
		return any(stage.TorusGeometrys_mapString).(map[string]Type)
	case *Triangle:
		return any(stage.Triangles_mapString).(map[string]Type)
	case *TubeGeometry:
		return any(stage.TubeGeometrys_mapString).(map[string]Type)
	case *Vector2:
		return any(stage.Vector2s_mapString).(map[string]Type)
	case *Vector3:
		return any(stage.Vector3s_mapString).(map[string]Type)
	default:
		return nil
	}
}

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *AmbiantLight:
		return any(&stage.AmbiantLights).(*map[Type]struct{})
	case *BoxGeometry:
		return any(&stage.BoxGeometrys).(*map[Type]struct{})
	case *BufferGeometry:
		return any(&stage.BufferGeometrys).(*map[Type]struct{})
	case *Camera:
		return any(&stage.Cameras).(*map[Type]struct{})
	case *Canvas:
		return any(&stage.Canvass).(*map[Type]struct{})
	case *Curve:
		return any(&stage.Curves).(*map[Type]struct{})
	case *CylinderGeometry:
		return any(&stage.CylinderGeometrys).(*map[Type]struct{})
	case *DirectionalLight:
		return any(&stage.DirectionalLights).(*map[Type]struct{})
	case *ExtrudeGeometry:
		return any(&stage.ExtrudeGeometrys).(*map[Type]struct{})
	case *Mesh:
		return any(&stage.Meshs).(*map[Type]struct{})
	case *MeshMaterialBasic:
		return any(&stage.MeshMaterialBasics).(*map[Type]struct{})
	case *MeshPhysicalMaterial:
		return any(&stage.MeshPhysicalMaterials).(*map[Type]struct{})
	case *PlaneGeometry:
		return any(&stage.PlaneGeometrys).(*map[Type]struct{})
	case *Shape:
		return any(&stage.Shapes).(*map[Type]struct{})
	case *SphereGeometry:
		return any(&stage.SphereGeometrys).(*map[Type]struct{})
	case *TorusGeometry:
		return any(&stage.TorusGeometrys).(*map[Type]struct{})
	case *Triangle:
		return any(&stage.Triangles).(*map[Type]struct{})
	case *TubeGeometry:
		return any(&stage.TubeGeometrys).(*map[Type]struct{})
	case *Vector2:
		return any(&stage.Vector2s).(*map[Type]struct{})
	case *Vector3:
		return any(&stage.Vector3s).(*map[Type]struct{})
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
	case BufferGeometry:
		return any(&BufferGeometry{
			Vertices: []*Vector3{{Name: "Vertices"}},
			Faces: []*Triangle{{Name: "Faces"}},
		}).(*Type)
	case Canvas:
		return any(&Canvas{
			DirectionalLights: []*DirectionalLight{{Name: "DirectionalLights"}},
			AmbiantLight: &AmbiantLight{Name: "AmbiantLight"},
			Meshs: []*Mesh{{Name: "Meshs"}},
			Camera: &Camera{Name: "Camera"},
		}).(*Type)
	case Curve:
		return any(&Curve{
			Points: []*Vector3{{Name: "Points"}},
		}).(*Type)
	case ExtrudeGeometry:
		return any(&ExtrudeGeometry{
			Shape: &Shape{Name: "Shape"},
			ExtrudePath: &Curve{Name: "ExtrudePath"},
		}).(*Type)
	case Mesh:
		return any(&Mesh{
			MeshMaterialBasic: &MeshMaterialBasic{Name: "MeshMaterialBasic"},
			MeshPhysicalMaterial: &MeshPhysicalMaterial{Name: "MeshPhysicalMaterial"},
			CylinderGeometry: &CylinderGeometry{Name: "CylinderGeometry"},
			BoxGeometry: &BoxGeometry{Name: "BoxGeometry"},
			SphereGeometry: &SphereGeometry{Name: "SphereGeometry"},
			TorusGeometry: &TorusGeometry{Name: "TorusGeometry"},
			PlaneGeometry: &PlaneGeometry{Name: "PlaneGeometry"},
			TubeGeometry: &TubeGeometry{Name: "TubeGeometry"},
			ExtrudeGeometry: &ExtrudeGeometry{Name: "ExtrudeGeometry"},
			BufferGeometry: &BufferGeometry{Name: "BufferGeometry"},
		}).(*Type)
	case Shape:
		return any(&Shape{
			Points: []*Vector2{{Name: "Points"}},
		}).(*Type)
	case TubeGeometry:
		return any(&TubeGeometry{
			Path: &Curve{Name: "Path"},
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
	// reverse maps of direct associations of AmbiantLight
	case AmbiantLight:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BoxGeometry
	case BoxGeometry:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BufferGeometry
	case BufferGeometry:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Camera
	case Camera:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Canvas
	case Canvas:
		switch fieldname {
		// insertion point for per direct association field
		case "AmbiantLight":
			res := make(map[*AmbiantLight][]*Canvas)
			for canvas := range stage.Canvass {
				if canvas.AmbiantLight != nil {
					ambiantlight_ := canvas.AmbiantLight
					var canvass []*Canvas
					_, ok := res[ambiantlight_]
					if ok {
						canvass = res[ambiantlight_]
					} else {
						canvass = make([]*Canvas, 0)
					}
					canvass = append(canvass, canvas)
					res[ambiantlight_] = canvass
				}
			}
			return any(res).(map[*End][]*Start)
		case "Camera":
			res := make(map[*Camera][]*Canvas)
			for canvas := range stage.Canvass {
				if canvas.Camera != nil {
					camera_ := canvas.Camera
					var canvass []*Canvas
					_, ok := res[camera_]
					if ok {
						canvass = res[camera_]
					} else {
						canvass = make([]*Canvas, 0)
					}
					canvass = append(canvass, canvas)
					res[camera_] = canvass
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Curve
	case Curve:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CylinderGeometry
	case CylinderGeometry:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DirectionalLight
	case DirectionalLight:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ExtrudeGeometry
	case ExtrudeGeometry:
		switch fieldname {
		// insertion point for per direct association field
		case "Shape":
			res := make(map[*Shape][]*ExtrudeGeometry)
			for extrudegeometry := range stage.ExtrudeGeometrys {
				if extrudegeometry.Shape != nil {
					shape_ := extrudegeometry.Shape
					var extrudegeometrys []*ExtrudeGeometry
					_, ok := res[shape_]
					if ok {
						extrudegeometrys = res[shape_]
					} else {
						extrudegeometrys = make([]*ExtrudeGeometry, 0)
					}
					extrudegeometrys = append(extrudegeometrys, extrudegeometry)
					res[shape_] = extrudegeometrys
				}
			}
			return any(res).(map[*End][]*Start)
		case "ExtrudePath":
			res := make(map[*Curve][]*ExtrudeGeometry)
			for extrudegeometry := range stage.ExtrudeGeometrys {
				if extrudegeometry.ExtrudePath != nil {
					curve_ := extrudegeometry.ExtrudePath
					var extrudegeometrys []*ExtrudeGeometry
					_, ok := res[curve_]
					if ok {
						extrudegeometrys = res[curve_]
					} else {
						extrudegeometrys = make([]*ExtrudeGeometry, 0)
					}
					extrudegeometrys = append(extrudegeometrys, extrudegeometry)
					res[curve_] = extrudegeometrys
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Mesh
	case Mesh:
		switch fieldname {
		// insertion point for per direct association field
		case "MeshMaterialBasic":
			res := make(map[*MeshMaterialBasic][]*Mesh)
			for mesh := range stage.Meshs {
				if mesh.MeshMaterialBasic != nil {
					meshmaterialbasic_ := mesh.MeshMaterialBasic
					var meshs []*Mesh
					_, ok := res[meshmaterialbasic_]
					if ok {
						meshs = res[meshmaterialbasic_]
					} else {
						meshs = make([]*Mesh, 0)
					}
					meshs = append(meshs, mesh)
					res[meshmaterialbasic_] = meshs
				}
			}
			return any(res).(map[*End][]*Start)
		case "MeshPhysicalMaterial":
			res := make(map[*MeshPhysicalMaterial][]*Mesh)
			for mesh := range stage.Meshs {
				if mesh.MeshPhysicalMaterial != nil {
					meshphysicalmaterial_ := mesh.MeshPhysicalMaterial
					var meshs []*Mesh
					_, ok := res[meshphysicalmaterial_]
					if ok {
						meshs = res[meshphysicalmaterial_]
					} else {
						meshs = make([]*Mesh, 0)
					}
					meshs = append(meshs, mesh)
					res[meshphysicalmaterial_] = meshs
				}
			}
			return any(res).(map[*End][]*Start)
		case "CylinderGeometry":
			res := make(map[*CylinderGeometry][]*Mesh)
			for mesh := range stage.Meshs {
				if mesh.CylinderGeometry != nil {
					cylindergeometry_ := mesh.CylinderGeometry
					var meshs []*Mesh
					_, ok := res[cylindergeometry_]
					if ok {
						meshs = res[cylindergeometry_]
					} else {
						meshs = make([]*Mesh, 0)
					}
					meshs = append(meshs, mesh)
					res[cylindergeometry_] = meshs
				}
			}
			return any(res).(map[*End][]*Start)
		case "BoxGeometry":
			res := make(map[*BoxGeometry][]*Mesh)
			for mesh := range stage.Meshs {
				if mesh.BoxGeometry != nil {
					boxgeometry_ := mesh.BoxGeometry
					var meshs []*Mesh
					_, ok := res[boxgeometry_]
					if ok {
						meshs = res[boxgeometry_]
					} else {
						meshs = make([]*Mesh, 0)
					}
					meshs = append(meshs, mesh)
					res[boxgeometry_] = meshs
				}
			}
			return any(res).(map[*End][]*Start)
		case "SphereGeometry":
			res := make(map[*SphereGeometry][]*Mesh)
			for mesh := range stage.Meshs {
				if mesh.SphereGeometry != nil {
					spheregeometry_ := mesh.SphereGeometry
					var meshs []*Mesh
					_, ok := res[spheregeometry_]
					if ok {
						meshs = res[spheregeometry_]
					} else {
						meshs = make([]*Mesh, 0)
					}
					meshs = append(meshs, mesh)
					res[spheregeometry_] = meshs
				}
			}
			return any(res).(map[*End][]*Start)
		case "TorusGeometry":
			res := make(map[*TorusGeometry][]*Mesh)
			for mesh := range stage.Meshs {
				if mesh.TorusGeometry != nil {
					torusgeometry_ := mesh.TorusGeometry
					var meshs []*Mesh
					_, ok := res[torusgeometry_]
					if ok {
						meshs = res[torusgeometry_]
					} else {
						meshs = make([]*Mesh, 0)
					}
					meshs = append(meshs, mesh)
					res[torusgeometry_] = meshs
				}
			}
			return any(res).(map[*End][]*Start)
		case "PlaneGeometry":
			res := make(map[*PlaneGeometry][]*Mesh)
			for mesh := range stage.Meshs {
				if mesh.PlaneGeometry != nil {
					planegeometry_ := mesh.PlaneGeometry
					var meshs []*Mesh
					_, ok := res[planegeometry_]
					if ok {
						meshs = res[planegeometry_]
					} else {
						meshs = make([]*Mesh, 0)
					}
					meshs = append(meshs, mesh)
					res[planegeometry_] = meshs
				}
			}
			return any(res).(map[*End][]*Start)
		case "TubeGeometry":
			res := make(map[*TubeGeometry][]*Mesh)
			for mesh := range stage.Meshs {
				if mesh.TubeGeometry != nil {
					tubegeometry_ := mesh.TubeGeometry
					var meshs []*Mesh
					_, ok := res[tubegeometry_]
					if ok {
						meshs = res[tubegeometry_]
					} else {
						meshs = make([]*Mesh, 0)
					}
					meshs = append(meshs, mesh)
					res[tubegeometry_] = meshs
				}
			}
			return any(res).(map[*End][]*Start)
		case "ExtrudeGeometry":
			res := make(map[*ExtrudeGeometry][]*Mesh)
			for mesh := range stage.Meshs {
				if mesh.ExtrudeGeometry != nil {
					extrudegeometry_ := mesh.ExtrudeGeometry
					var meshs []*Mesh
					_, ok := res[extrudegeometry_]
					if ok {
						meshs = res[extrudegeometry_]
					} else {
						meshs = make([]*Mesh, 0)
					}
					meshs = append(meshs, mesh)
					res[extrudegeometry_] = meshs
				}
			}
			return any(res).(map[*End][]*Start)
		case "BufferGeometry":
			res := make(map[*BufferGeometry][]*Mesh)
			for mesh := range stage.Meshs {
				if mesh.BufferGeometry != nil {
					buffergeometry_ := mesh.BufferGeometry
					var meshs []*Mesh
					_, ok := res[buffergeometry_]
					if ok {
						meshs = res[buffergeometry_]
					} else {
						meshs = make([]*Mesh, 0)
					}
					meshs = append(meshs, mesh)
					res[buffergeometry_] = meshs
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MeshMaterialBasic
	case MeshMaterialBasic:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MeshPhysicalMaterial
	case MeshPhysicalMaterial:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PlaneGeometry
	case PlaneGeometry:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Shape
	case Shape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SphereGeometry
	case SphereGeometry:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TorusGeometry
	case TorusGeometry:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Triangle
	case Triangle:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TubeGeometry
	case TubeGeometry:
		switch fieldname {
		// insertion point for per direct association field
		case "Path":
			res := make(map[*Curve][]*TubeGeometry)
			for tubegeometry := range stage.TubeGeometrys {
				if tubegeometry.Path != nil {
					curve_ := tubegeometry.Path
					var tubegeometrys []*TubeGeometry
					_, ok := res[curve_]
					if ok {
						tubegeometrys = res[curve_]
					} else {
						tubegeometrys = make([]*TubeGeometry, 0)
					}
					tubegeometrys = append(tubegeometrys, tubegeometry)
					res[curve_] = tubegeometrys
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Vector2
	case Vector2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Vector3
	case Vector3:
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
	// reverse maps of direct associations of AmbiantLight
	case AmbiantLight:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BoxGeometry
	case BoxGeometry:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BufferGeometry
	case BufferGeometry:
		switch fieldname {
		// insertion point for per direct association field
		case "Vertices":
			res := make(map[*Vector3][]*BufferGeometry)
			for buffergeometry := range stage.BufferGeometrys {
				for _, vector3_ := range buffergeometry.Vertices {
					res[vector3_] = append(res[vector3_], buffergeometry)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Faces":
			res := make(map[*Triangle][]*BufferGeometry)
			for buffergeometry := range stage.BufferGeometrys {
				for _, triangle_ := range buffergeometry.Faces {
					res[triangle_] = append(res[triangle_], buffergeometry)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Camera
	case Camera:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Canvas
	case Canvas:
		switch fieldname {
		// insertion point for per direct association field
		case "DirectionalLights":
			res := make(map[*DirectionalLight][]*Canvas)
			for canvas := range stage.Canvass {
				for _, directionallight_ := range canvas.DirectionalLights {
					res[directionallight_] = append(res[directionallight_], canvas)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Meshs":
			res := make(map[*Mesh][]*Canvas)
			for canvas := range stage.Canvass {
				for _, mesh_ := range canvas.Meshs {
					res[mesh_] = append(res[mesh_], canvas)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Curve
	case Curve:
		switch fieldname {
		// insertion point for per direct association field
		case "Points":
			res := make(map[*Vector3][]*Curve)
			for curve := range stage.Curves {
				for _, vector3_ := range curve.Points {
					res[vector3_] = append(res[vector3_], curve)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of CylinderGeometry
	case CylinderGeometry:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DirectionalLight
	case DirectionalLight:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ExtrudeGeometry
	case ExtrudeGeometry:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Mesh
	case Mesh:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MeshMaterialBasic
	case MeshMaterialBasic:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MeshPhysicalMaterial
	case MeshPhysicalMaterial:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PlaneGeometry
	case PlaneGeometry:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Shape
	case Shape:
		switch fieldname {
		// insertion point for per direct association field
		case "Points":
			res := make(map[*Vector2][]*Shape)
			for shape := range stage.Shapes {
				for _, vector2_ := range shape.Points {
					res[vector2_] = append(res[vector2_], shape)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SphereGeometry
	case SphereGeometry:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TorusGeometry
	case TorusGeometry:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Triangle
	case Triangle:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TubeGeometry
	case TubeGeometry:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Vector2
	case Vector2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Vector3
	case Vector3:
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
	case *AmbiantLight:
		res = any(new(AmbiantLight)).(Type)
	case *BoxGeometry:
		res = any(new(BoxGeometry)).(Type)
	case *BufferGeometry:
		res = any(new(BufferGeometry)).(Type)
	case *Camera:
		res = any(new(Camera)).(Type)
	case *Canvas:
		res = any(new(Canvas)).(Type)
	case *Curve:
		res = any(new(Curve)).(Type)
	case *CylinderGeometry:
		res = any(new(CylinderGeometry)).(Type)
	case *DirectionalLight:
		res = any(new(DirectionalLight)).(Type)
	case *ExtrudeGeometry:
		res = any(new(ExtrudeGeometry)).(Type)
	case *Mesh:
		res = any(new(Mesh)).(Type)
	case *MeshMaterialBasic:
		res = any(new(MeshMaterialBasic)).(Type)
	case *MeshPhysicalMaterial:
		res = any(new(MeshPhysicalMaterial)).(Type)
	case *PlaneGeometry:
		res = any(new(PlaneGeometry)).(Type)
	case *Shape:
		res = any(new(Shape)).(Type)
	case *SphereGeometry:
		res = any(new(SphereGeometry)).(Type)
	case *TorusGeometry:
		res = any(new(TorusGeometry)).(Type)
	case *Triangle:
		res = any(new(Triangle)).(Type)
	case *TubeGeometry:
		res = any(new(TubeGeometry)).(Type)
	case *Vector2:
		res = any(new(Vector2)).(Type)
	case *Vector3:
		res = any(new(Vector3)).(Type)
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
	case *AmbiantLight:
		res = "AmbiantLight"
	case *BoxGeometry:
		res = "BoxGeometry"
	case *BufferGeometry:
		res = "BufferGeometry"
	case *Camera:
		res = "Camera"
	case *Canvas:
		res = "Canvas"
	case *Curve:
		res = "Curve"
	case *CylinderGeometry:
		res = "CylinderGeometry"
	case *DirectionalLight:
		res = "DirectionalLight"
	case *ExtrudeGeometry:
		res = "ExtrudeGeometry"
	case *Mesh:
		res = "Mesh"
	case *MeshMaterialBasic:
		res = "MeshMaterialBasic"
	case *MeshPhysicalMaterial:
		res = "MeshPhysicalMaterial"
	case *PlaneGeometry:
		res = "PlaneGeometry"
	case *Shape:
		res = "Shape"
	case *SphereGeometry:
		res = "SphereGeometry"
	case *TorusGeometry:
		res = "TorusGeometry"
	case *Triangle:
		res = "Triangle"
	case *TubeGeometry:
		res = "TubeGeometry"
	case *Vector2:
		res = "Vector2"
	case *Vector3:
		res = "Vector3"
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
	case *AmbiantLight:
		var rf ReverseField
		_ = rf
	case *BoxGeometry:
		var rf ReverseField
		_ = rf
	case *BufferGeometry:
		var rf ReverseField
		_ = rf
	case *Camera:
		var rf ReverseField
		_ = rf
	case *Canvas:
		var rf ReverseField
		_ = rf
	case *Curve:
		var rf ReverseField
		_ = rf
	case *CylinderGeometry:
		var rf ReverseField
		_ = rf
	case *DirectionalLight:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Canvas"
		rf.Fieldname = "DirectionalLights"
		res = append(res, rf)
	case *ExtrudeGeometry:
		var rf ReverseField
		_ = rf
	case *Mesh:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Canvas"
		rf.Fieldname = "Meshs"
		res = append(res, rf)
	case *MeshMaterialBasic:
		var rf ReverseField
		_ = rf
	case *MeshPhysicalMaterial:
		var rf ReverseField
		_ = rf
	case *PlaneGeometry:
		var rf ReverseField
		_ = rf
	case *Shape:
		var rf ReverseField
		_ = rf
	case *SphereGeometry:
		var rf ReverseField
		_ = rf
	case *TorusGeometry:
		var rf ReverseField
		_ = rf
	case *Triangle:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "BufferGeometry"
		rf.Fieldname = "Faces"
		res = append(res, rf)
	case *TubeGeometry:
		var rf ReverseField
		_ = rf
	case *Vector2:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Shape"
		rf.Fieldname = "Points"
		res = append(res, rf)
	case *Vector3:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "BufferGeometry"
		rf.Fieldname = "Vertices"
		res = append(res, rf)
		rf.GongstructName = "Curve"
		rf.Fieldname = "Points"
		res = append(res, rf)
	}
	return
}

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
}

// insertion point for get fields header method
func (ambiantlight *AmbiantLight) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Intensity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (boxgeometry *BoxGeometry) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
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
			Name:               "Depth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "WidthSegments",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "HeightSegments",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "DepthSegments",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (buffergeometry *BufferGeometry) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Vertices",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Vector3",
		},
		{
			Name:                 "Faces",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Triangle",
		},
	}
	return
}

func (camera *Camera) GongGetFieldHeaders() (res []GongFieldHeader) {
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
		{
			Name:               "Z",
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

func (canvas *Canvas) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "DirectionalLights",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DirectionalLight",
		},
		{
			Name:                 "AmbiantLight",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AmbiantLight",
		},
		{
			Name:                 "Meshs",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Mesh",
		},
		{
			Name:                 "Camera",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Camera",
		},
		{
			Name:               "IsWithLastRenderingUpdate",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "LastRendering",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "Frame64BitsEncoded",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (curve *Curve) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Points",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Vector3",
		},
	}
	return
}

func (cylindergeometry *CylinderGeometry) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RadiusTop",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "RadiusBottom",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "RadialSegments",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "HeightSegments",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "OpenEnded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ThetaStart",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ThetaLength",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (directionallight *DirectionalLight) GongGetFieldHeaders() (res []GongFieldHeader) {
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
		{
			Name:               "Z",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Intensity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsWithCastShadow",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (extrudegeometry *ExtrudeGeometry) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Shape",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Shape",
		},
		{
			Name:                 "ExtrudePath",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Curve",
		},
		{
			Name:               "Steps",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (mesh *Mesh) GongGetFieldHeaders() (res []GongFieldHeader) {
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
		{
			Name:               "Z",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "MeshMaterialBasic",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MeshMaterialBasic",
		},
		{
			Name:                 "MeshPhysicalMaterial",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MeshPhysicalMaterial",
		},
		{
			Name:                 "CylinderGeometry",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CylinderGeometry",
		},
		{
			Name:                 "BoxGeometry",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "BoxGeometry",
		},
		{
			Name:                 "SphereGeometry",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SphereGeometry",
		},
		{
			Name:                 "TorusGeometry",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TorusGeometry",
		},
		{
			Name:                 "PlaneGeometry",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PlaneGeometry",
		},
		{
			Name:                 "TubeGeometry",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TubeGeometry",
		},
		{
			Name:                 "ExtrudeGeometry",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ExtrudeGeometry",
		},
		{
			Name:                 "BufferGeometry",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "BufferGeometry",
		},
	}
	return
}

func (meshmaterialbasic *MeshMaterialBasic) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Wireframe",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "Opacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Transparent",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "Visible",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (planegeometry *PlaneGeometry) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
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
			Name:               "WidthSegments",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "HeightSegments",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (shape *Shape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Points",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Vector2",
		},
	}
	return
}

func (spheregeometry *SphereGeometry) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Radius",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "WidthSegments",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "HeightSegments",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "PhiStart",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "PhiLength",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ThetaStart",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ThetaLength",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (torusgeometry *TorusGeometry) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Radius",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Tube",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "RadialSegments",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "TubularSegments",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Arc",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (triangle *Triangle) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "V1",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "V2",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "V3",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (tubegeometry *TubeGeometry) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Path",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Curve",
		},
		{
			Name:               "TubularSegments",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Radius",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "RadialSegments",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Closed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (vector2 *Vector2) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (vector3 *Vector3) GongGetFieldHeaders() (res []GongFieldHeader) {
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
		{
			Name:               "Z",
			GongFieldValueType: GongFieldValueTypeFloat,
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
func (ambiantlight *AmbiantLight) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ambiantlight.Name
	case "Intensity":
		res.valueString = fmt.Sprintf("%f", ambiantlight.Intensity)
		res.valueFloat = ambiantlight.Intensity
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (boxgeometry *BoxGeometry) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = boxgeometry.Name
	case "Width":
		res.valueString = fmt.Sprintf("%f", boxgeometry.Width)
		res.valueFloat = boxgeometry.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", boxgeometry.Height)
		res.valueFloat = boxgeometry.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Depth":
		res.valueString = fmt.Sprintf("%f", boxgeometry.Depth)
		res.valueFloat = boxgeometry.Depth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "WidthSegments":
		res.valueString = fmt.Sprintf("%d", boxgeometry.WidthSegments)
		res.valueInt = boxgeometry.WidthSegments
		res.GongFieldValueType = GongFieldValueTypeInt
	case "HeightSegments":
		res.valueString = fmt.Sprintf("%d", boxgeometry.HeightSegments)
		res.valueInt = boxgeometry.HeightSegments
		res.GongFieldValueType = GongFieldValueTypeInt
	case "DepthSegments":
		res.valueString = fmt.Sprintf("%d", boxgeometry.DepthSegments)
		res.valueInt = boxgeometry.DepthSegments
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (buffergeometry *BufferGeometry) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = buffergeometry.Name
	case "Vertices":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range buffergeometry.Vertices {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Faces":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range buffergeometry.Faces {
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

func (camera *Camera) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = camera.Name
	case "X":
		res.valueString = fmt.Sprintf("%f", camera.X)
		res.valueFloat = camera.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", camera.Y)
		res.valueFloat = camera.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Z":
		res.valueString = fmt.Sprintf("%f", camera.Z)
		res.valueFloat = camera.Z
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "TargetX":
		res.valueString = fmt.Sprintf("%f", camera.TargetX)
		res.valueFloat = camera.TargetX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "TargetY":
		res.valueString = fmt.Sprintf("%f", camera.TargetY)
		res.valueFloat = camera.TargetY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "TargetZ":
		res.valueString = fmt.Sprintf("%f", camera.TargetZ)
		res.valueFloat = camera.TargetZ
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Fov":
		res.valueString = fmt.Sprintf("%f", camera.Fov)
		res.valueFloat = camera.Fov
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (canvas *Canvas) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = canvas.Name
	case "DirectionalLights":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range canvas.DirectionalLights {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "AmbiantLight":
		res.GongFieldValueType = GongFieldValueTypePointer
		if canvas.AmbiantLight != nil {
			res.valueString = canvas.AmbiantLight.Name
			res.ids = canvas.AmbiantLight.GongGetUUID(stage)
		}
	case "Meshs":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range canvas.Meshs {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Camera":
		res.GongFieldValueType = GongFieldValueTypePointer
		if canvas.Camera != nil {
			res.valueString = canvas.Camera.Name
			res.ids = canvas.Camera.GongGetUUID(stage)
		}
	case "IsWithLastRenderingUpdate":
		res.valueString = fmt.Sprintf("%t", canvas.IsWithLastRenderingUpdate)
		res.valueBool = canvas.IsWithLastRenderingUpdate
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LastRendering":
		res.valueString = canvas.LastRendering.String()
	case "Frame64BitsEncoded":
		res.valueString = canvas.Frame64BitsEncoded
	}
	return
}

func (curve *Curve) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = curve.Name
	case "Points":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range curve.Points {
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

func (cylindergeometry *CylinderGeometry) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = cylindergeometry.Name
	case "RadiusTop":
		res.valueString = fmt.Sprintf("%f", cylindergeometry.RadiusTop)
		res.valueFloat = cylindergeometry.RadiusTop
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadiusBottom":
		res.valueString = fmt.Sprintf("%f", cylindergeometry.RadiusBottom)
		res.valueFloat = cylindergeometry.RadiusBottom
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", cylindergeometry.Height)
		res.valueFloat = cylindergeometry.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadialSegments":
		res.valueString = fmt.Sprintf("%d", cylindergeometry.RadialSegments)
		res.valueInt = cylindergeometry.RadialSegments
		res.GongFieldValueType = GongFieldValueTypeInt
	case "HeightSegments":
		res.valueString = fmt.Sprintf("%d", cylindergeometry.HeightSegments)
		res.valueInt = cylindergeometry.HeightSegments
		res.GongFieldValueType = GongFieldValueTypeInt
	case "OpenEnded":
		res.valueString = fmt.Sprintf("%t", cylindergeometry.OpenEnded)
		res.valueBool = cylindergeometry.OpenEnded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ThetaStart":
		res.valueString = fmt.Sprintf("%f", cylindergeometry.ThetaStart)
		res.valueFloat = cylindergeometry.ThetaStart
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ThetaLength":
		res.valueString = fmt.Sprintf("%f", cylindergeometry.ThetaLength)
		res.valueFloat = cylindergeometry.ThetaLength
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (directionallight *DirectionalLight) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = directionallight.Name
	case "X":
		res.valueString = fmt.Sprintf("%f", directionallight.X)
		res.valueFloat = directionallight.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", directionallight.Y)
		res.valueFloat = directionallight.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Z":
		res.valueString = fmt.Sprintf("%f", directionallight.Z)
		res.valueFloat = directionallight.Z
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Intensity":
		res.valueString = fmt.Sprintf("%f", directionallight.Intensity)
		res.valueFloat = directionallight.Intensity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsWithCastShadow":
		res.valueString = fmt.Sprintf("%t", directionallight.IsWithCastShadow)
		res.valueBool = directionallight.IsWithCastShadow
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (extrudegeometry *ExtrudeGeometry) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = extrudegeometry.Name
	case "Shape":
		res.GongFieldValueType = GongFieldValueTypePointer
		if extrudegeometry.Shape != nil {
			res.valueString = extrudegeometry.Shape.Name
			res.ids = extrudegeometry.Shape.GongGetUUID(stage)
		}
	case "ExtrudePath":
		res.GongFieldValueType = GongFieldValueTypePointer
		if extrudegeometry.ExtrudePath != nil {
			res.valueString = extrudegeometry.ExtrudePath.Name
			res.ids = extrudegeometry.ExtrudePath.GongGetUUID(stage)
		}
	case "Steps":
		res.valueString = fmt.Sprintf("%d", extrudegeometry.Steps)
		res.valueInt = extrudegeometry.Steps
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (mesh *Mesh) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = mesh.Name
	case "X":
		res.valueString = fmt.Sprintf("%f", mesh.X)
		res.valueFloat = mesh.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", mesh.Y)
		res.valueFloat = mesh.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Z":
		res.valueString = fmt.Sprintf("%f", mesh.Z)
		res.valueFloat = mesh.Z
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "MeshMaterialBasic":
		res.GongFieldValueType = GongFieldValueTypePointer
		if mesh.MeshMaterialBasic != nil {
			res.valueString = mesh.MeshMaterialBasic.Name
			res.ids = mesh.MeshMaterialBasic.GongGetUUID(stage)
		}
	case "MeshPhysicalMaterial":
		res.GongFieldValueType = GongFieldValueTypePointer
		if mesh.MeshPhysicalMaterial != nil {
			res.valueString = mesh.MeshPhysicalMaterial.Name
			res.ids = mesh.MeshPhysicalMaterial.GongGetUUID(stage)
		}
	case "CylinderGeometry":
		res.GongFieldValueType = GongFieldValueTypePointer
		if mesh.CylinderGeometry != nil {
			res.valueString = mesh.CylinderGeometry.Name
			res.ids = mesh.CylinderGeometry.GongGetUUID(stage)
		}
	case "BoxGeometry":
		res.GongFieldValueType = GongFieldValueTypePointer
		if mesh.BoxGeometry != nil {
			res.valueString = mesh.BoxGeometry.Name
			res.ids = mesh.BoxGeometry.GongGetUUID(stage)
		}
	case "SphereGeometry":
		res.GongFieldValueType = GongFieldValueTypePointer
		if mesh.SphereGeometry != nil {
			res.valueString = mesh.SphereGeometry.Name
			res.ids = mesh.SphereGeometry.GongGetUUID(stage)
		}
	case "TorusGeometry":
		res.GongFieldValueType = GongFieldValueTypePointer
		if mesh.TorusGeometry != nil {
			res.valueString = mesh.TorusGeometry.Name
			res.ids = mesh.TorusGeometry.GongGetUUID(stage)
		}
	case "PlaneGeometry":
		res.GongFieldValueType = GongFieldValueTypePointer
		if mesh.PlaneGeometry != nil {
			res.valueString = mesh.PlaneGeometry.Name
			res.ids = mesh.PlaneGeometry.GongGetUUID(stage)
		}
	case "TubeGeometry":
		res.GongFieldValueType = GongFieldValueTypePointer
		if mesh.TubeGeometry != nil {
			res.valueString = mesh.TubeGeometry.Name
			res.ids = mesh.TubeGeometry.GongGetUUID(stage)
		}
	case "ExtrudeGeometry":
		res.GongFieldValueType = GongFieldValueTypePointer
		if mesh.ExtrudeGeometry != nil {
			res.valueString = mesh.ExtrudeGeometry.Name
			res.ids = mesh.ExtrudeGeometry.GongGetUUID(stage)
		}
	case "BufferGeometry":
		res.GongFieldValueType = GongFieldValueTypePointer
		if mesh.BufferGeometry != nil {
			res.valueString = mesh.BufferGeometry.Name
			res.ids = mesh.BufferGeometry.GongGetUUID(stage)
		}
	}
	return
}

func (meshmaterialbasic *MeshMaterialBasic) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = meshmaterialbasic.Name
	case "Color":
		res.valueString = meshmaterialbasic.Color
	}
	return
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = meshphysicalmaterial.Name
	case "Color":
		res.valueString = meshphysicalmaterial.Color
	case "Wireframe":
		res.valueString = fmt.Sprintf("%t", meshphysicalmaterial.Wireframe)
		res.valueBool = meshphysicalmaterial.Wireframe
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Opacity":
		res.valueString = fmt.Sprintf("%f", meshphysicalmaterial.Opacity)
		res.valueFloat = meshphysicalmaterial.Opacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Transparent":
		res.valueString = fmt.Sprintf("%t", meshphysicalmaterial.Transparent)
		res.valueBool = meshphysicalmaterial.Transparent
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Visible":
		res.valueString = fmt.Sprintf("%t", meshphysicalmaterial.Visible)
		res.valueBool = meshphysicalmaterial.Visible
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (planegeometry *PlaneGeometry) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = planegeometry.Name
	case "Width":
		res.valueString = fmt.Sprintf("%f", planegeometry.Width)
		res.valueFloat = planegeometry.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", planegeometry.Height)
		res.valueFloat = planegeometry.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "WidthSegments":
		res.valueString = fmt.Sprintf("%d", planegeometry.WidthSegments)
		res.valueInt = planegeometry.WidthSegments
		res.GongFieldValueType = GongFieldValueTypeInt
	case "HeightSegments":
		res.valueString = fmt.Sprintf("%d", planegeometry.HeightSegments)
		res.valueInt = planegeometry.HeightSegments
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (shape *Shape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = shape.Name
	case "Points":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range shape.Points {
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

func (spheregeometry *SphereGeometry) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = spheregeometry.Name
	case "Radius":
		res.valueString = fmt.Sprintf("%f", spheregeometry.Radius)
		res.valueFloat = spheregeometry.Radius
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "WidthSegments":
		res.valueString = fmt.Sprintf("%d", spheregeometry.WidthSegments)
		res.valueInt = spheregeometry.WidthSegments
		res.GongFieldValueType = GongFieldValueTypeInt
	case "HeightSegments":
		res.valueString = fmt.Sprintf("%d", spheregeometry.HeightSegments)
		res.valueInt = spheregeometry.HeightSegments
		res.GongFieldValueType = GongFieldValueTypeInt
	case "PhiStart":
		res.valueString = fmt.Sprintf("%f", spheregeometry.PhiStart)
		res.valueFloat = spheregeometry.PhiStart
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "PhiLength":
		res.valueString = fmt.Sprintf("%f", spheregeometry.PhiLength)
		res.valueFloat = spheregeometry.PhiLength
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ThetaStart":
		res.valueString = fmt.Sprintf("%f", spheregeometry.ThetaStart)
		res.valueFloat = spheregeometry.ThetaStart
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ThetaLength":
		res.valueString = fmt.Sprintf("%f", spheregeometry.ThetaLength)
		res.valueFloat = spheregeometry.ThetaLength
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (torusgeometry *TorusGeometry) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = torusgeometry.Name
	case "Radius":
		res.valueString = fmt.Sprintf("%f", torusgeometry.Radius)
		res.valueFloat = torusgeometry.Radius
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Tube":
		res.valueString = fmt.Sprintf("%f", torusgeometry.Tube)
		res.valueFloat = torusgeometry.Tube
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadialSegments":
		res.valueString = fmt.Sprintf("%d", torusgeometry.RadialSegments)
		res.valueInt = torusgeometry.RadialSegments
		res.GongFieldValueType = GongFieldValueTypeInt
	case "TubularSegments":
		res.valueString = fmt.Sprintf("%d", torusgeometry.TubularSegments)
		res.valueInt = torusgeometry.TubularSegments
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Arc":
		res.valueString = fmt.Sprintf("%f", torusgeometry.Arc)
		res.valueFloat = torusgeometry.Arc
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (triangle *Triangle) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = triangle.Name
	case "V1":
		res.valueString = fmt.Sprintf("%d", triangle.V1)
		res.valueInt = triangle.V1
		res.GongFieldValueType = GongFieldValueTypeInt
	case "V2":
		res.valueString = fmt.Sprintf("%d", triangle.V2)
		res.valueInt = triangle.V2
		res.GongFieldValueType = GongFieldValueTypeInt
	case "V3":
		res.valueString = fmt.Sprintf("%d", triangle.V3)
		res.valueInt = triangle.V3
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (tubegeometry *TubeGeometry) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = tubegeometry.Name
	case "Path":
		res.GongFieldValueType = GongFieldValueTypePointer
		if tubegeometry.Path != nil {
			res.valueString = tubegeometry.Path.Name
			res.ids = tubegeometry.Path.GongGetUUID(stage)
		}
	case "TubularSegments":
		res.valueString = fmt.Sprintf("%d", tubegeometry.TubularSegments)
		res.valueInt = tubegeometry.TubularSegments
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Radius":
		res.valueString = fmt.Sprintf("%f", tubegeometry.Radius)
		res.valueFloat = tubegeometry.Radius
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RadialSegments":
		res.valueString = fmt.Sprintf("%d", tubegeometry.RadialSegments)
		res.valueInt = tubegeometry.RadialSegments
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Closed":
		res.valueString = fmt.Sprintf("%t", tubegeometry.Closed)
		res.valueBool = tubegeometry.Closed
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (vector2 *Vector2) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = vector2.Name
	case "X":
		res.valueString = fmt.Sprintf("%f", vector2.X)
		res.valueFloat = vector2.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", vector2.Y)
		res.valueFloat = vector2.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (vector3 *Vector3) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = vector3.Name
	case "X":
		res.valueString = fmt.Sprintf("%f", vector3.X)
		res.valueFloat = vector3.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", vector3.Y)
		res.valueFloat = vector3.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Z":
		res.valueString = fmt.Sprintf("%f", vector3.Z)
		res.valueFloat = vector3.Z
		res.GongFieldValueType = GongFieldValueTypeFloat
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
func (ambiantlight *AmbiantLight) GongGetGongstructName() string {
	return "AmbiantLight"
}

func (boxgeometry *BoxGeometry) GongGetGongstructName() string {
	return "BoxGeometry"
}

func (buffergeometry *BufferGeometry) GongGetGongstructName() string {
	return "BufferGeometry"
}

func (camera *Camera) GongGetGongstructName() string {
	return "Camera"
}

func (canvas *Canvas) GongGetGongstructName() string {
	return "Canvas"
}

func (curve *Curve) GongGetGongstructName() string {
	return "Curve"
}

func (cylindergeometry *CylinderGeometry) GongGetGongstructName() string {
	return "CylinderGeometry"
}

func (directionallight *DirectionalLight) GongGetGongstructName() string {
	return "DirectionalLight"
}

func (extrudegeometry *ExtrudeGeometry) GongGetGongstructName() string {
	return "ExtrudeGeometry"
}

func (mesh *Mesh) GongGetGongstructName() string {
	return "Mesh"
}

func (meshmaterialbasic *MeshMaterialBasic) GongGetGongstructName() string {
	return "MeshMaterialBasic"
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongGetGongstructName() string {
	return "MeshPhysicalMaterial"
}

func (planegeometry *PlaneGeometry) GongGetGongstructName() string {
	return "PlaneGeometry"
}

func (shape *Shape) GongGetGongstructName() string {
	return "Shape"
}

func (spheregeometry *SphereGeometry) GongGetGongstructName() string {
	return "SphereGeometry"
}

func (torusgeometry *TorusGeometry) GongGetGongstructName() string {
	return "TorusGeometry"
}

func (triangle *Triangle) GongGetGongstructName() string {
	return "Triangle"
}

func (tubegeometry *TubeGeometry) GongGetGongstructName() string {
	return "TubeGeometry"
}

func (vector2 *Vector2) GongGetGongstructName() string {
	return "Vector2"
}

func (vector3 *Vector3) GongGetGongstructName() string {
	return "Vector3"
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
	__gong__rebuildMapString(stage.AmbiantLights, &stage.AmbiantLights_mapString)

	__gong__rebuildMapString(stage.BoxGeometrys, &stage.BoxGeometrys_mapString)

	__gong__rebuildMapString(stage.BufferGeometrys, &stage.BufferGeometrys_mapString)

	__gong__rebuildMapString(stage.Cameras, &stage.Cameras_mapString)

	__gong__rebuildMapString(stage.Canvass, &stage.Canvass_mapString)

	__gong__rebuildMapString(stage.Curves, &stage.Curves_mapString)

	__gong__rebuildMapString(stage.CylinderGeometrys, &stage.CylinderGeometrys_mapString)

	__gong__rebuildMapString(stage.DirectionalLights, &stage.DirectionalLights_mapString)

	__gong__rebuildMapString(stage.ExtrudeGeometrys, &stage.ExtrudeGeometrys_mapString)

	__gong__rebuildMapString(stage.Meshs, &stage.Meshs_mapString)

	__gong__rebuildMapString(stage.MeshMaterialBasics, &stage.MeshMaterialBasics_mapString)

	__gong__rebuildMapString(stage.MeshPhysicalMaterials, &stage.MeshPhysicalMaterials_mapString)

	__gong__rebuildMapString(stage.PlaneGeometrys, &stage.PlaneGeometrys_mapString)

	__gong__rebuildMapString(stage.Shapes, &stage.Shapes_mapString)

	__gong__rebuildMapString(stage.SphereGeometrys, &stage.SphereGeometrys_mapString)

	__gong__rebuildMapString(stage.TorusGeometrys, &stage.TorusGeometrys_mapString)

	__gong__rebuildMapString(stage.Triangles, &stage.Triangles_mapString)

	__gong__rebuildMapString(stage.TubeGeometrys, &stage.TubeGeometrys_mapString)

	__gong__rebuildMapString(stage.Vector2s, &stage.Vector2s_mapString)

	__gong__rebuildMapString(stage.Vector3s, &stage.Vector3s_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
