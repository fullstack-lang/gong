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

	threejs_go "github.com/fullstack-lang/gong/lib/threejs/go"
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
	AmbiantLights                map[*AmbiantLight]struct{}
	AmbiantLights_instance       map[*AmbiantLight]*AmbiantLight
	AmbiantLights_mapString      map[string]*AmbiantLight
	AmbiantLightOrder            uint
	AmbiantLight_stagedOrder     map[*AmbiantLight]uint
	AmbiantLight_orderStaged     map[uint]*AmbiantLight
	AmbiantLights_reference      map[*AmbiantLight]*AmbiantLight
	AmbiantLights_referenceOrder map[*AmbiantLight]uint

	// insertion point for slice of pointers maps
	OnAfterAmbiantLightCreateCallback OnAfterCreateInterface[AmbiantLight]
	OnAfterAmbiantLightUpdateCallback OnAfterUpdateInterface[AmbiantLight]
	OnAfterAmbiantLightDeleteCallback OnAfterDeleteInterface[AmbiantLight]
	OnAfterAmbiantLightReadCallback   OnAfterReadInterface[AmbiantLight]

	BoxGeometrys                map[*BoxGeometry]struct{}
	BoxGeometrys_instance       map[*BoxGeometry]*BoxGeometry
	BoxGeometrys_mapString      map[string]*BoxGeometry
	BoxGeometryOrder            uint
	BoxGeometry_stagedOrder     map[*BoxGeometry]uint
	BoxGeometry_orderStaged     map[uint]*BoxGeometry
	BoxGeometrys_reference      map[*BoxGeometry]*BoxGeometry
	BoxGeometrys_referenceOrder map[*BoxGeometry]uint

	// insertion point for slice of pointers maps
	OnAfterBoxGeometryCreateCallback OnAfterCreateInterface[BoxGeometry]
	OnAfterBoxGeometryUpdateCallback OnAfterUpdateInterface[BoxGeometry]
	OnAfterBoxGeometryDeleteCallback OnAfterDeleteInterface[BoxGeometry]
	OnAfterBoxGeometryReadCallback   OnAfterReadInterface[BoxGeometry]

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

	OnAfterBufferGeometryCreateCallback OnAfterCreateInterface[BufferGeometry]
	OnAfterBufferGeometryUpdateCallback OnAfterUpdateInterface[BufferGeometry]
	OnAfterBufferGeometryDeleteCallback OnAfterDeleteInterface[BufferGeometry]
	OnAfterBufferGeometryReadCallback   OnAfterReadInterface[BufferGeometry]

	Cameras                map[*Camera]struct{}
	Cameras_instance       map[*Camera]*Camera
	Cameras_mapString      map[string]*Camera
	CameraOrder            uint
	Camera_stagedOrder     map[*Camera]uint
	Camera_orderStaged     map[uint]*Camera
	Cameras_reference      map[*Camera]*Camera
	Cameras_referenceOrder map[*Camera]uint

	// insertion point for slice of pointers maps
	OnAfterCameraCreateCallback OnAfterCreateInterface[Camera]
	OnAfterCameraUpdateCallback OnAfterUpdateInterface[Camera]
	OnAfterCameraDeleteCallback OnAfterDeleteInterface[Camera]
	OnAfterCameraReadCallback   OnAfterReadInterface[Camera]

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

	OnAfterCanvasCreateCallback OnAfterCreateInterface[Canvas]
	OnAfterCanvasUpdateCallback OnAfterUpdateInterface[Canvas]
	OnAfterCanvasDeleteCallback OnAfterDeleteInterface[Canvas]
	OnAfterCanvasReadCallback   OnAfterReadInterface[Canvas]

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

	OnAfterCurveCreateCallback OnAfterCreateInterface[Curve]
	OnAfterCurveUpdateCallback OnAfterUpdateInterface[Curve]
	OnAfterCurveDeleteCallback OnAfterDeleteInterface[Curve]
	OnAfterCurveReadCallback   OnAfterReadInterface[Curve]

	CylinderGeometrys                map[*CylinderGeometry]struct{}
	CylinderGeometrys_instance       map[*CylinderGeometry]*CylinderGeometry
	CylinderGeometrys_mapString      map[string]*CylinderGeometry
	CylinderGeometryOrder            uint
	CylinderGeometry_stagedOrder     map[*CylinderGeometry]uint
	CylinderGeometry_orderStaged     map[uint]*CylinderGeometry
	CylinderGeometrys_reference      map[*CylinderGeometry]*CylinderGeometry
	CylinderGeometrys_referenceOrder map[*CylinderGeometry]uint

	// insertion point for slice of pointers maps
	OnAfterCylinderGeometryCreateCallback OnAfterCreateInterface[CylinderGeometry]
	OnAfterCylinderGeometryUpdateCallback OnAfterUpdateInterface[CylinderGeometry]
	OnAfterCylinderGeometryDeleteCallback OnAfterDeleteInterface[CylinderGeometry]
	OnAfterCylinderGeometryReadCallback   OnAfterReadInterface[CylinderGeometry]

	DirectionalLights                map[*DirectionalLight]struct{}
	DirectionalLights_instance       map[*DirectionalLight]*DirectionalLight
	DirectionalLights_mapString      map[string]*DirectionalLight
	DirectionalLightOrder            uint
	DirectionalLight_stagedOrder     map[*DirectionalLight]uint
	DirectionalLight_orderStaged     map[uint]*DirectionalLight
	DirectionalLights_reference      map[*DirectionalLight]*DirectionalLight
	DirectionalLights_referenceOrder map[*DirectionalLight]uint

	// insertion point for slice of pointers maps
	OnAfterDirectionalLightCreateCallback OnAfterCreateInterface[DirectionalLight]
	OnAfterDirectionalLightUpdateCallback OnAfterUpdateInterface[DirectionalLight]
	OnAfterDirectionalLightDeleteCallback OnAfterDeleteInterface[DirectionalLight]
	OnAfterDirectionalLightReadCallback   OnAfterReadInterface[DirectionalLight]

	ExtrudeGeometrys                map[*ExtrudeGeometry]struct{}
	ExtrudeGeometrys_instance       map[*ExtrudeGeometry]*ExtrudeGeometry
	ExtrudeGeometrys_mapString      map[string]*ExtrudeGeometry
	ExtrudeGeometryOrder            uint
	ExtrudeGeometry_stagedOrder     map[*ExtrudeGeometry]uint
	ExtrudeGeometry_orderStaged     map[uint]*ExtrudeGeometry
	ExtrudeGeometrys_reference      map[*ExtrudeGeometry]*ExtrudeGeometry
	ExtrudeGeometrys_referenceOrder map[*ExtrudeGeometry]uint

	// insertion point for slice of pointers maps
	OnAfterExtrudeGeometryCreateCallback OnAfterCreateInterface[ExtrudeGeometry]
	OnAfterExtrudeGeometryUpdateCallback OnAfterUpdateInterface[ExtrudeGeometry]
	OnAfterExtrudeGeometryDeleteCallback OnAfterDeleteInterface[ExtrudeGeometry]
	OnAfterExtrudeGeometryReadCallback   OnAfterReadInterface[ExtrudeGeometry]

	Meshs                map[*Mesh]struct{}
	Meshs_instance       map[*Mesh]*Mesh
	Meshs_mapString      map[string]*Mesh
	MeshOrder            uint
	Mesh_stagedOrder     map[*Mesh]uint
	Mesh_orderStaged     map[uint]*Mesh
	Meshs_reference      map[*Mesh]*Mesh
	Meshs_referenceOrder map[*Mesh]uint

	// insertion point for slice of pointers maps
	OnAfterMeshCreateCallback OnAfterCreateInterface[Mesh]
	OnAfterMeshUpdateCallback OnAfterUpdateInterface[Mesh]
	OnAfterMeshDeleteCallback OnAfterDeleteInterface[Mesh]
	OnAfterMeshReadCallback   OnAfterReadInterface[Mesh]

	MeshMaterialBasics                map[*MeshMaterialBasic]struct{}
	MeshMaterialBasics_instance       map[*MeshMaterialBasic]*MeshMaterialBasic
	MeshMaterialBasics_mapString      map[string]*MeshMaterialBasic
	MeshMaterialBasicOrder            uint
	MeshMaterialBasic_stagedOrder     map[*MeshMaterialBasic]uint
	MeshMaterialBasic_orderStaged     map[uint]*MeshMaterialBasic
	MeshMaterialBasics_reference      map[*MeshMaterialBasic]*MeshMaterialBasic
	MeshMaterialBasics_referenceOrder map[*MeshMaterialBasic]uint

	// insertion point for slice of pointers maps
	OnAfterMeshMaterialBasicCreateCallback OnAfterCreateInterface[MeshMaterialBasic]
	OnAfterMeshMaterialBasicUpdateCallback OnAfterUpdateInterface[MeshMaterialBasic]
	OnAfterMeshMaterialBasicDeleteCallback OnAfterDeleteInterface[MeshMaterialBasic]
	OnAfterMeshMaterialBasicReadCallback   OnAfterReadInterface[MeshMaterialBasic]

	MeshPhysicalMaterials                map[*MeshPhysicalMaterial]struct{}
	MeshPhysicalMaterials_instance       map[*MeshPhysicalMaterial]*MeshPhysicalMaterial
	MeshPhysicalMaterials_mapString      map[string]*MeshPhysicalMaterial
	MeshPhysicalMaterialOrder            uint
	MeshPhysicalMaterial_stagedOrder     map[*MeshPhysicalMaterial]uint
	MeshPhysicalMaterial_orderStaged     map[uint]*MeshPhysicalMaterial
	MeshPhysicalMaterials_reference      map[*MeshPhysicalMaterial]*MeshPhysicalMaterial
	MeshPhysicalMaterials_referenceOrder map[*MeshPhysicalMaterial]uint

	// insertion point for slice of pointers maps
	OnAfterMeshPhysicalMaterialCreateCallback OnAfterCreateInterface[MeshPhysicalMaterial]
	OnAfterMeshPhysicalMaterialUpdateCallback OnAfterUpdateInterface[MeshPhysicalMaterial]
	OnAfterMeshPhysicalMaterialDeleteCallback OnAfterDeleteInterface[MeshPhysicalMaterial]
	OnAfterMeshPhysicalMaterialReadCallback   OnAfterReadInterface[MeshPhysicalMaterial]

	PlaneGeometrys                map[*PlaneGeometry]struct{}
	PlaneGeometrys_instance       map[*PlaneGeometry]*PlaneGeometry
	PlaneGeometrys_mapString      map[string]*PlaneGeometry
	PlaneGeometryOrder            uint
	PlaneGeometry_stagedOrder     map[*PlaneGeometry]uint
	PlaneGeometry_orderStaged     map[uint]*PlaneGeometry
	PlaneGeometrys_reference      map[*PlaneGeometry]*PlaneGeometry
	PlaneGeometrys_referenceOrder map[*PlaneGeometry]uint

	// insertion point for slice of pointers maps
	OnAfterPlaneGeometryCreateCallback OnAfterCreateInterface[PlaneGeometry]
	OnAfterPlaneGeometryUpdateCallback OnAfterUpdateInterface[PlaneGeometry]
	OnAfterPlaneGeometryDeleteCallback OnAfterDeleteInterface[PlaneGeometry]
	OnAfterPlaneGeometryReadCallback   OnAfterReadInterface[PlaneGeometry]

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

	OnAfterShapeCreateCallback OnAfterCreateInterface[Shape]
	OnAfterShapeUpdateCallback OnAfterUpdateInterface[Shape]
	OnAfterShapeDeleteCallback OnAfterDeleteInterface[Shape]
	OnAfterShapeReadCallback   OnAfterReadInterface[Shape]

	SphereGeometrys                map[*SphereGeometry]struct{}
	SphereGeometrys_instance       map[*SphereGeometry]*SphereGeometry
	SphereGeometrys_mapString      map[string]*SphereGeometry
	SphereGeometryOrder            uint
	SphereGeometry_stagedOrder     map[*SphereGeometry]uint
	SphereGeometry_orderStaged     map[uint]*SphereGeometry
	SphereGeometrys_reference      map[*SphereGeometry]*SphereGeometry
	SphereGeometrys_referenceOrder map[*SphereGeometry]uint

	// insertion point for slice of pointers maps
	OnAfterSphereGeometryCreateCallback OnAfterCreateInterface[SphereGeometry]
	OnAfterSphereGeometryUpdateCallback OnAfterUpdateInterface[SphereGeometry]
	OnAfterSphereGeometryDeleteCallback OnAfterDeleteInterface[SphereGeometry]
	OnAfterSphereGeometryReadCallback   OnAfterReadInterface[SphereGeometry]

	TorusGeometrys                map[*TorusGeometry]struct{}
	TorusGeometrys_instance       map[*TorusGeometry]*TorusGeometry
	TorusGeometrys_mapString      map[string]*TorusGeometry
	TorusGeometryOrder            uint
	TorusGeometry_stagedOrder     map[*TorusGeometry]uint
	TorusGeometry_orderStaged     map[uint]*TorusGeometry
	TorusGeometrys_reference      map[*TorusGeometry]*TorusGeometry
	TorusGeometrys_referenceOrder map[*TorusGeometry]uint

	// insertion point for slice of pointers maps
	OnAfterTorusGeometryCreateCallback OnAfterCreateInterface[TorusGeometry]
	OnAfterTorusGeometryUpdateCallback OnAfterUpdateInterface[TorusGeometry]
	OnAfterTorusGeometryDeleteCallback OnAfterDeleteInterface[TorusGeometry]
	OnAfterTorusGeometryReadCallback   OnAfterReadInterface[TorusGeometry]

	Triangles                map[*Triangle]struct{}
	Triangles_instance       map[*Triangle]*Triangle
	Triangles_mapString      map[string]*Triangle
	TriangleOrder            uint
	Triangle_stagedOrder     map[*Triangle]uint
	Triangle_orderStaged     map[uint]*Triangle
	Triangles_reference      map[*Triangle]*Triangle
	Triangles_referenceOrder map[*Triangle]uint

	// insertion point for slice of pointers maps
	OnAfterTriangleCreateCallback OnAfterCreateInterface[Triangle]
	OnAfterTriangleUpdateCallback OnAfterUpdateInterface[Triangle]
	OnAfterTriangleDeleteCallback OnAfterDeleteInterface[Triangle]
	OnAfterTriangleReadCallback   OnAfterReadInterface[Triangle]

	TubeGeometrys                map[*TubeGeometry]struct{}
	TubeGeometrys_instance       map[*TubeGeometry]*TubeGeometry
	TubeGeometrys_mapString      map[string]*TubeGeometry
	TubeGeometryOrder            uint
	TubeGeometry_stagedOrder     map[*TubeGeometry]uint
	TubeGeometry_orderStaged     map[uint]*TubeGeometry
	TubeGeometrys_reference      map[*TubeGeometry]*TubeGeometry
	TubeGeometrys_referenceOrder map[*TubeGeometry]uint

	// insertion point for slice of pointers maps
	OnAfterTubeGeometryCreateCallback OnAfterCreateInterface[TubeGeometry]
	OnAfterTubeGeometryUpdateCallback OnAfterUpdateInterface[TubeGeometry]
	OnAfterTubeGeometryDeleteCallback OnAfterDeleteInterface[TubeGeometry]
	OnAfterTubeGeometryReadCallback   OnAfterReadInterface[TubeGeometry]

	Vector2s                map[*Vector2]struct{}
	Vector2s_instance       map[*Vector2]*Vector2
	Vector2s_mapString      map[string]*Vector2
	Vector2Order            uint
	Vector2_stagedOrder     map[*Vector2]uint
	Vector2_orderStaged     map[uint]*Vector2
	Vector2s_reference      map[*Vector2]*Vector2
	Vector2s_referenceOrder map[*Vector2]uint

	// insertion point for slice of pointers maps
	OnAfterVector2CreateCallback OnAfterCreateInterface[Vector2]
	OnAfterVector2UpdateCallback OnAfterUpdateInterface[Vector2]
	OnAfterVector2DeleteCallback OnAfterDeleteInterface[Vector2]
	OnAfterVector2ReadCallback   OnAfterReadInterface[Vector2]

	Vector3s                map[*Vector3]struct{}
	Vector3s_instance       map[*Vector3]*Vector3
	Vector3s_mapString      map[string]*Vector3
	Vector3Order            uint
	Vector3_stagedOrder     map[*Vector3]uint
	Vector3_orderStaged     map[uint]*Vector3
	Vector3s_reference      map[*Vector3]*Vector3
	Vector3s_referenceOrder map[*Vector3]uint

	// insertion point for slice of pointers maps
	OnAfterVector3CreateCallback OnAfterCreateInterface[Vector3]
	OnAfterVector3UpdateCallback OnAfterUpdateInterface[Vector3]
	OnAfterVector3DeleteCallback OnAfterDeleteInterface[Vector3]
	OnAfterVector3ReadCallback   OnAfterReadInterface[Vector3]

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
	stage.AmbiantLights_reference = make(map[*AmbiantLight]*AmbiantLight)
	stage.AmbiantLights_instance = make(map[*AmbiantLight]*AmbiantLight)
	stage.AmbiantLights_referenceOrder = make(map[*AmbiantLight]uint)

	stage.BoxGeometrys_reference = make(map[*BoxGeometry]*BoxGeometry)
	stage.BoxGeometrys_instance = make(map[*BoxGeometry]*BoxGeometry)
	stage.BoxGeometrys_referenceOrder = make(map[*BoxGeometry]uint)

	stage.BufferGeometrys_reference = make(map[*BufferGeometry]*BufferGeometry)
	stage.BufferGeometrys_instance = make(map[*BufferGeometry]*BufferGeometry)
	stage.BufferGeometrys_referenceOrder = make(map[*BufferGeometry]uint)

	stage.Cameras_reference = make(map[*Camera]*Camera)
	stage.Cameras_instance = make(map[*Camera]*Camera)
	stage.Cameras_referenceOrder = make(map[*Camera]uint)

	stage.Canvass_reference = make(map[*Canvas]*Canvas)
	stage.Canvass_instance = make(map[*Canvas]*Canvas)
	stage.Canvass_referenceOrder = make(map[*Canvas]uint)

	stage.Curves_reference = make(map[*Curve]*Curve)
	stage.Curves_instance = make(map[*Curve]*Curve)
	stage.Curves_referenceOrder = make(map[*Curve]uint)

	stage.CylinderGeometrys_reference = make(map[*CylinderGeometry]*CylinderGeometry)
	stage.CylinderGeometrys_instance = make(map[*CylinderGeometry]*CylinderGeometry)
	stage.CylinderGeometrys_referenceOrder = make(map[*CylinderGeometry]uint)

	stage.DirectionalLights_reference = make(map[*DirectionalLight]*DirectionalLight)
	stage.DirectionalLights_instance = make(map[*DirectionalLight]*DirectionalLight)
	stage.DirectionalLights_referenceOrder = make(map[*DirectionalLight]uint)

	stage.ExtrudeGeometrys_reference = make(map[*ExtrudeGeometry]*ExtrudeGeometry)
	stage.ExtrudeGeometrys_instance = make(map[*ExtrudeGeometry]*ExtrudeGeometry)
	stage.ExtrudeGeometrys_referenceOrder = make(map[*ExtrudeGeometry]uint)

	stage.Meshs_reference = make(map[*Mesh]*Mesh)
	stage.Meshs_instance = make(map[*Mesh]*Mesh)
	stage.Meshs_referenceOrder = make(map[*Mesh]uint)

	stage.MeshMaterialBasics_reference = make(map[*MeshMaterialBasic]*MeshMaterialBasic)
	stage.MeshMaterialBasics_instance = make(map[*MeshMaterialBasic]*MeshMaterialBasic)
	stage.MeshMaterialBasics_referenceOrder = make(map[*MeshMaterialBasic]uint)

	stage.MeshPhysicalMaterials_reference = make(map[*MeshPhysicalMaterial]*MeshPhysicalMaterial)
	stage.MeshPhysicalMaterials_instance = make(map[*MeshPhysicalMaterial]*MeshPhysicalMaterial)
	stage.MeshPhysicalMaterials_referenceOrder = make(map[*MeshPhysicalMaterial]uint)

	stage.PlaneGeometrys_reference = make(map[*PlaneGeometry]*PlaneGeometry)
	stage.PlaneGeometrys_instance = make(map[*PlaneGeometry]*PlaneGeometry)
	stage.PlaneGeometrys_referenceOrder = make(map[*PlaneGeometry]uint)

	stage.Shapes_reference = make(map[*Shape]*Shape)
	stage.Shapes_instance = make(map[*Shape]*Shape)
	stage.Shapes_referenceOrder = make(map[*Shape]uint)

	stage.SphereGeometrys_reference = make(map[*SphereGeometry]*SphereGeometry)
	stage.SphereGeometrys_instance = make(map[*SphereGeometry]*SphereGeometry)
	stage.SphereGeometrys_referenceOrder = make(map[*SphereGeometry]uint)

	stage.TorusGeometrys_reference = make(map[*TorusGeometry]*TorusGeometry)
	stage.TorusGeometrys_instance = make(map[*TorusGeometry]*TorusGeometry)
	stage.TorusGeometrys_referenceOrder = make(map[*TorusGeometry]uint)

	stage.Triangles_reference = make(map[*Triangle]*Triangle)
	stage.Triangles_instance = make(map[*Triangle]*Triangle)
	stage.Triangles_referenceOrder = make(map[*Triangle]uint)

	stage.TubeGeometrys_reference = make(map[*TubeGeometry]*TubeGeometry)
	stage.TubeGeometrys_instance = make(map[*TubeGeometry]*TubeGeometry)
	stage.TubeGeometrys_referenceOrder = make(map[*TubeGeometry]uint)

	stage.Vector2s_reference = make(map[*Vector2]*Vector2)
	stage.Vector2s_instance = make(map[*Vector2]*Vector2)
	stage.Vector2s_referenceOrder = make(map[*Vector2]uint)

	stage.Vector3s_reference = make(map[*Vector3]*Vector3)
	stage.Vector3s_instance = make(map[*Vector3]*Vector3)
	stage.Vector3s_referenceOrder = make(map[*Vector3]uint)

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
	var maxAmbiantLightOrder uint
	var foundAmbiantLight bool
	for _, order := range stage.AmbiantLight_stagedOrder {
		if !foundAmbiantLight || order > maxAmbiantLightOrder {
			maxAmbiantLightOrder = order
			foundAmbiantLight = true
		}
	}
	if foundAmbiantLight {
		stage.AmbiantLightOrder = maxAmbiantLightOrder + 1
	} else {
		stage.AmbiantLightOrder = 0
	}

	var maxBoxGeometryOrder uint
	var foundBoxGeometry bool
	for _, order := range stage.BoxGeometry_stagedOrder {
		if !foundBoxGeometry || order > maxBoxGeometryOrder {
			maxBoxGeometryOrder = order
			foundBoxGeometry = true
		}
	}
	if foundBoxGeometry {
		stage.BoxGeometryOrder = maxBoxGeometryOrder + 1
	} else {
		stage.BoxGeometryOrder = 0
	}

	var maxBufferGeometryOrder uint
	var foundBufferGeometry bool
	for _, order := range stage.BufferGeometry_stagedOrder {
		if !foundBufferGeometry || order > maxBufferGeometryOrder {
			maxBufferGeometryOrder = order
			foundBufferGeometry = true
		}
	}
	if foundBufferGeometry {
		stage.BufferGeometryOrder = maxBufferGeometryOrder + 1
	} else {
		stage.BufferGeometryOrder = 0
	}

	var maxCameraOrder uint
	var foundCamera bool
	for _, order := range stage.Camera_stagedOrder {
		if !foundCamera || order > maxCameraOrder {
			maxCameraOrder = order
			foundCamera = true
		}
	}
	if foundCamera {
		stage.CameraOrder = maxCameraOrder + 1
	} else {
		stage.CameraOrder = 0
	}

	var maxCanvasOrder uint
	var foundCanvas bool
	for _, order := range stage.Canvas_stagedOrder {
		if !foundCanvas || order > maxCanvasOrder {
			maxCanvasOrder = order
			foundCanvas = true
		}
	}
	if foundCanvas {
		stage.CanvasOrder = maxCanvasOrder + 1
	} else {
		stage.CanvasOrder = 0
	}

	var maxCurveOrder uint
	var foundCurve bool
	for _, order := range stage.Curve_stagedOrder {
		if !foundCurve || order > maxCurveOrder {
			maxCurveOrder = order
			foundCurve = true
		}
	}
	if foundCurve {
		stage.CurveOrder = maxCurveOrder + 1
	} else {
		stage.CurveOrder = 0
	}

	var maxCylinderGeometryOrder uint
	var foundCylinderGeometry bool
	for _, order := range stage.CylinderGeometry_stagedOrder {
		if !foundCylinderGeometry || order > maxCylinderGeometryOrder {
			maxCylinderGeometryOrder = order
			foundCylinderGeometry = true
		}
	}
	if foundCylinderGeometry {
		stage.CylinderGeometryOrder = maxCylinderGeometryOrder + 1
	} else {
		stage.CylinderGeometryOrder = 0
	}

	var maxDirectionalLightOrder uint
	var foundDirectionalLight bool
	for _, order := range stage.DirectionalLight_stagedOrder {
		if !foundDirectionalLight || order > maxDirectionalLightOrder {
			maxDirectionalLightOrder = order
			foundDirectionalLight = true
		}
	}
	if foundDirectionalLight {
		stage.DirectionalLightOrder = maxDirectionalLightOrder + 1
	} else {
		stage.DirectionalLightOrder = 0
	}

	var maxExtrudeGeometryOrder uint
	var foundExtrudeGeometry bool
	for _, order := range stage.ExtrudeGeometry_stagedOrder {
		if !foundExtrudeGeometry || order > maxExtrudeGeometryOrder {
			maxExtrudeGeometryOrder = order
			foundExtrudeGeometry = true
		}
	}
	if foundExtrudeGeometry {
		stage.ExtrudeGeometryOrder = maxExtrudeGeometryOrder + 1
	} else {
		stage.ExtrudeGeometryOrder = 0
	}

	var maxMeshOrder uint
	var foundMesh bool
	for _, order := range stage.Mesh_stagedOrder {
		if !foundMesh || order > maxMeshOrder {
			maxMeshOrder = order
			foundMesh = true
		}
	}
	if foundMesh {
		stage.MeshOrder = maxMeshOrder + 1
	} else {
		stage.MeshOrder = 0
	}

	var maxMeshMaterialBasicOrder uint
	var foundMeshMaterialBasic bool
	for _, order := range stage.MeshMaterialBasic_stagedOrder {
		if !foundMeshMaterialBasic || order > maxMeshMaterialBasicOrder {
			maxMeshMaterialBasicOrder = order
			foundMeshMaterialBasic = true
		}
	}
	if foundMeshMaterialBasic {
		stage.MeshMaterialBasicOrder = maxMeshMaterialBasicOrder + 1
	} else {
		stage.MeshMaterialBasicOrder = 0
	}

	var maxMeshPhysicalMaterialOrder uint
	var foundMeshPhysicalMaterial bool
	for _, order := range stage.MeshPhysicalMaterial_stagedOrder {
		if !foundMeshPhysicalMaterial || order > maxMeshPhysicalMaterialOrder {
			maxMeshPhysicalMaterialOrder = order
			foundMeshPhysicalMaterial = true
		}
	}
	if foundMeshPhysicalMaterial {
		stage.MeshPhysicalMaterialOrder = maxMeshPhysicalMaterialOrder + 1
	} else {
		stage.MeshPhysicalMaterialOrder = 0
	}

	var maxPlaneGeometryOrder uint
	var foundPlaneGeometry bool
	for _, order := range stage.PlaneGeometry_stagedOrder {
		if !foundPlaneGeometry || order > maxPlaneGeometryOrder {
			maxPlaneGeometryOrder = order
			foundPlaneGeometry = true
		}
	}
	if foundPlaneGeometry {
		stage.PlaneGeometryOrder = maxPlaneGeometryOrder + 1
	} else {
		stage.PlaneGeometryOrder = 0
	}

	var maxShapeOrder uint
	var foundShape bool
	for _, order := range stage.Shape_stagedOrder {
		if !foundShape || order > maxShapeOrder {
			maxShapeOrder = order
			foundShape = true
		}
	}
	if foundShape {
		stage.ShapeOrder = maxShapeOrder + 1
	} else {
		stage.ShapeOrder = 0
	}

	var maxSphereGeometryOrder uint
	var foundSphereGeometry bool
	for _, order := range stage.SphereGeometry_stagedOrder {
		if !foundSphereGeometry || order > maxSphereGeometryOrder {
			maxSphereGeometryOrder = order
			foundSphereGeometry = true
		}
	}
	if foundSphereGeometry {
		stage.SphereGeometryOrder = maxSphereGeometryOrder + 1
	} else {
		stage.SphereGeometryOrder = 0
	}

	var maxTorusGeometryOrder uint
	var foundTorusGeometry bool
	for _, order := range stage.TorusGeometry_stagedOrder {
		if !foundTorusGeometry || order > maxTorusGeometryOrder {
			maxTorusGeometryOrder = order
			foundTorusGeometry = true
		}
	}
	if foundTorusGeometry {
		stage.TorusGeometryOrder = maxTorusGeometryOrder + 1
	} else {
		stage.TorusGeometryOrder = 0
	}

	var maxTriangleOrder uint
	var foundTriangle bool
	for _, order := range stage.Triangle_stagedOrder {
		if !foundTriangle || order > maxTriangleOrder {
			maxTriangleOrder = order
			foundTriangle = true
		}
	}
	if foundTriangle {
		stage.TriangleOrder = maxTriangleOrder + 1
	} else {
		stage.TriangleOrder = 0
	}

	var maxTubeGeometryOrder uint
	var foundTubeGeometry bool
	for _, order := range stage.TubeGeometry_stagedOrder {
		if !foundTubeGeometry || order > maxTubeGeometryOrder {
			maxTubeGeometryOrder = order
			foundTubeGeometry = true
		}
	}
	if foundTubeGeometry {
		stage.TubeGeometryOrder = maxTubeGeometryOrder + 1
	} else {
		stage.TubeGeometryOrder = 0
	}

	var maxVector2Order uint
	var foundVector2 bool
	for _, order := range stage.Vector2_stagedOrder {
		if !foundVector2 || order > maxVector2Order {
			maxVector2Order = order
			foundVector2 = true
		}
	}
	if foundVector2 {
		stage.Vector2Order = maxVector2Order + 1
	} else {
		stage.Vector2Order = 0
	}

	var maxVector3Order uint
	var foundVector3 bool
	for _, order := range stage.Vector3_stagedOrder {
		if !foundVector3 || order > maxVector3Order {
			maxVector3Order = order
			foundVector3 = true
		}
	}
	if foundVector3 {
		stage.Vector3Order = maxVector3Order + 1
	} else {
		stage.Vector3Order = 0
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
	case *AmbiantLight:
		tmp := GetStructInstancesByOrder(stage.AmbiantLights, stage.AmbiantLight_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *AmbiantLight implements.
			res = append(res, any(v).(T))
		}
		return res
	case *BoxGeometry:
		tmp := GetStructInstancesByOrder(stage.BoxGeometrys, stage.BoxGeometry_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *BoxGeometry implements.
			res = append(res, any(v).(T))
		}
		return res
	case *BufferGeometry:
		tmp := GetStructInstancesByOrder(stage.BufferGeometrys, stage.BufferGeometry_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *BufferGeometry implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Camera:
		tmp := GetStructInstancesByOrder(stage.Cameras, stage.Camera_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Camera implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Canvas:
		tmp := GetStructInstancesByOrder(stage.Canvass, stage.Canvas_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Canvas implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Curve:
		tmp := GetStructInstancesByOrder(stage.Curves, stage.Curve_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Curve implements.
			res = append(res, any(v).(T))
		}
		return res
	case *CylinderGeometry:
		tmp := GetStructInstancesByOrder(stage.CylinderGeometrys, stage.CylinderGeometry_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *CylinderGeometry implements.
			res = append(res, any(v).(T))
		}
		return res
	case *DirectionalLight:
		tmp := GetStructInstancesByOrder(stage.DirectionalLights, stage.DirectionalLight_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *DirectionalLight implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ExtrudeGeometry:
		tmp := GetStructInstancesByOrder(stage.ExtrudeGeometrys, stage.ExtrudeGeometry_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ExtrudeGeometry implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Mesh:
		tmp := GetStructInstancesByOrder(stage.Meshs, stage.Mesh_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Mesh implements.
			res = append(res, any(v).(T))
		}
		return res
	case *MeshMaterialBasic:
		tmp := GetStructInstancesByOrder(stage.MeshMaterialBasics, stage.MeshMaterialBasic_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *MeshMaterialBasic implements.
			res = append(res, any(v).(T))
		}
		return res
	case *MeshPhysicalMaterial:
		tmp := GetStructInstancesByOrder(stage.MeshPhysicalMaterials, stage.MeshPhysicalMaterial_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *MeshPhysicalMaterial implements.
			res = append(res, any(v).(T))
		}
		return res
	case *PlaneGeometry:
		tmp := GetStructInstancesByOrder(stage.PlaneGeometrys, stage.PlaneGeometry_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *PlaneGeometry implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Shape:
		tmp := GetStructInstancesByOrder(stage.Shapes, stage.Shape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Shape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *SphereGeometry:
		tmp := GetStructInstancesByOrder(stage.SphereGeometrys, stage.SphereGeometry_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SphereGeometry implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TorusGeometry:
		tmp := GetStructInstancesByOrder(stage.TorusGeometrys, stage.TorusGeometry_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TorusGeometry implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Triangle:
		tmp := GetStructInstancesByOrder(stage.Triangles, stage.Triangle_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Triangle implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TubeGeometry:
		tmp := GetStructInstancesByOrder(stage.TubeGeometrys, stage.TubeGeometry_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TubeGeometry implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Vector2:
		tmp := GetStructInstancesByOrder(stage.Vector2s, stage.Vector2_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Vector2 implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Vector3:
		tmp := GetStructInstancesByOrder(stage.Vector3s, stage.Vector3_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Vector3 implements.
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
	case "AmbiantLight":
		res = GetNamedStructInstances(stage.AmbiantLights, stage.AmbiantLight_stagedOrder)
	case "BoxGeometry":
		res = GetNamedStructInstances(stage.BoxGeometrys, stage.BoxGeometry_stagedOrder)
	case "BufferGeometry":
		res = GetNamedStructInstances(stage.BufferGeometrys, stage.BufferGeometry_stagedOrder)
	case "Camera":
		res = GetNamedStructInstances(stage.Cameras, stage.Camera_stagedOrder)
	case "Canvas":
		res = GetNamedStructInstances(stage.Canvass, stage.Canvas_stagedOrder)
	case "Curve":
		res = GetNamedStructInstances(stage.Curves, stage.Curve_stagedOrder)
	case "CylinderGeometry":
		res = GetNamedStructInstances(stage.CylinderGeometrys, stage.CylinderGeometry_stagedOrder)
	case "DirectionalLight":
		res = GetNamedStructInstances(stage.DirectionalLights, stage.DirectionalLight_stagedOrder)
	case "ExtrudeGeometry":
		res = GetNamedStructInstances(stage.ExtrudeGeometrys, stage.ExtrudeGeometry_stagedOrder)
	case "Mesh":
		res = GetNamedStructInstances(stage.Meshs, stage.Mesh_stagedOrder)
	case "MeshMaterialBasic":
		res = GetNamedStructInstances(stage.MeshMaterialBasics, stage.MeshMaterialBasic_stagedOrder)
	case "MeshPhysicalMaterial":
		res = GetNamedStructInstances(stage.MeshPhysicalMaterials, stage.MeshPhysicalMaterial_stagedOrder)
	case "PlaneGeometry":
		res = GetNamedStructInstances(stage.PlaneGeometrys, stage.PlaneGeometry_stagedOrder)
	case "Shape":
		res = GetNamedStructInstances(stage.Shapes, stage.Shape_stagedOrder)
	case "SphereGeometry":
		res = GetNamedStructInstances(stage.SphereGeometrys, stage.SphereGeometry_stagedOrder)
	case "TorusGeometry":
		res = GetNamedStructInstances(stage.TorusGeometrys, stage.TorusGeometry_stagedOrder)
	case "Triangle":
		res = GetNamedStructInstances(stage.Triangles, stage.Triangle_stagedOrder)
	case "TubeGeometry":
		res = GetNamedStructInstances(stage.TubeGeometrys, stage.TubeGeometry_stagedOrder)
	case "Vector2":
		res = GetNamedStructInstances(stage.Vector2s, stage.Vector2_stagedOrder)
	case "Vector3":
		res = GetNamedStructInstances(stage.Vector3s, stage.Vector3_stagedOrder)
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
	return "github.com/fullstack-lang/gong/lib/threejs/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return threejs_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return threejs_go.GoDiagramsDir
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
	CommitAmbiantLight(ambiantlight *AmbiantLight)
	CheckoutAmbiantLight(ambiantlight *AmbiantLight)
	CommitBoxGeometry(boxgeometry *BoxGeometry)
	CheckoutBoxGeometry(boxgeometry *BoxGeometry)
	CommitBufferGeometry(buffergeometry *BufferGeometry)
	CheckoutBufferGeometry(buffergeometry *BufferGeometry)
	CommitCamera(camera *Camera)
	CheckoutCamera(camera *Camera)
	CommitCanvas(canvas *Canvas)
	CheckoutCanvas(canvas *Canvas)
	CommitCurve(curve *Curve)
	CheckoutCurve(curve *Curve)
	CommitCylinderGeometry(cylindergeometry *CylinderGeometry)
	CheckoutCylinderGeometry(cylindergeometry *CylinderGeometry)
	CommitDirectionalLight(directionallight *DirectionalLight)
	CheckoutDirectionalLight(directionallight *DirectionalLight)
	CommitExtrudeGeometry(extrudegeometry *ExtrudeGeometry)
	CheckoutExtrudeGeometry(extrudegeometry *ExtrudeGeometry)
	CommitMesh(mesh *Mesh)
	CheckoutMesh(mesh *Mesh)
	CommitMeshMaterialBasic(meshmaterialbasic *MeshMaterialBasic)
	CheckoutMeshMaterialBasic(meshmaterialbasic *MeshMaterialBasic)
	CommitMeshPhysicalMaterial(meshphysicalmaterial *MeshPhysicalMaterial)
	CheckoutMeshPhysicalMaterial(meshphysicalmaterial *MeshPhysicalMaterial)
	CommitPlaneGeometry(planegeometry *PlaneGeometry)
	CheckoutPlaneGeometry(planegeometry *PlaneGeometry)
	CommitShape(shape *Shape)
	CheckoutShape(shape *Shape)
	CommitSphereGeometry(spheregeometry *SphereGeometry)
	CheckoutSphereGeometry(spheregeometry *SphereGeometry)
	CommitTorusGeometry(torusgeometry *TorusGeometry)
	CheckoutTorusGeometry(torusgeometry *TorusGeometry)
	CommitTriangle(triangle *Triangle)
	CheckoutTriangle(triangle *Triangle)
	CommitTubeGeometry(tubegeometry *TubeGeometry)
	CheckoutTubeGeometry(tubegeometry *TubeGeometry)
	CommitVector2(vector2 *Vector2)
	CheckoutVector2(vector2 *Vector2)
	CommitVector3(vector3 *Vector3)
	CheckoutVector3(vector3 *Vector3)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

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
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
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

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "AmbiantLight"},
			{name: "BoxGeometry"},
			{name: "BufferGeometry"},
			{name: "Camera"},
			{name: "Canvas"},
			{name: "Curve"},
			{name: "CylinderGeometry"},
			{name: "DirectionalLight"},
			{name: "ExtrudeGeometry"},
			{name: "Mesh"},
			{name: "MeshMaterialBasic"},
			{name: "MeshPhysicalMaterial"},
			{name: "PlaneGeometry"},
			{name: "Shape"},
			{name: "SphereGeometry"},
			{name: "TorusGeometry"},
			{name: "Triangle"},
			{name: "TubeGeometry"},
			{name: "Vector2"},
			{name: "Vector3"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *AmbiantLight:
		return stage.AmbiantLight_stagedOrder[instance]
	case *BoxGeometry:
		return stage.BoxGeometry_stagedOrder[instance]
	case *BufferGeometry:
		return stage.BufferGeometry_stagedOrder[instance]
	case *Camera:
		return stage.Camera_stagedOrder[instance]
	case *Canvas:
		return stage.Canvas_stagedOrder[instance]
	case *Curve:
		return stage.Curve_stagedOrder[instance]
	case *CylinderGeometry:
		return stage.CylinderGeometry_stagedOrder[instance]
	case *DirectionalLight:
		return stage.DirectionalLight_stagedOrder[instance]
	case *ExtrudeGeometry:
		return stage.ExtrudeGeometry_stagedOrder[instance]
	case *Mesh:
		return stage.Mesh_stagedOrder[instance]
	case *MeshMaterialBasic:
		return stage.MeshMaterialBasic_stagedOrder[instance]
	case *MeshPhysicalMaterial:
		return stage.MeshPhysicalMaterial_stagedOrder[instance]
	case *PlaneGeometry:
		return stage.PlaneGeometry_stagedOrder[instance]
	case *Shape:
		return stage.Shape_stagedOrder[instance]
	case *SphereGeometry:
		return stage.SphereGeometry_stagedOrder[instance]
	case *TorusGeometry:
		return stage.TorusGeometry_stagedOrder[instance]
	case *Triangle:
		return stage.Triangle_stagedOrder[instance]
	case *TubeGeometry:
		return stage.TubeGeometry_stagedOrder[instance]
	case *Vector2:
		return stage.Vector2_stagedOrder[instance]
	case *Vector3:
		return stage.Vector3_stagedOrder[instance]
	default:
		return 0 // should not happen
	}
}

func GongGetInstanceFromOrder[Type PointerToGongstruct](stage *Stage, order uint) (res Type) {
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

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *AmbiantLight:
		return stage.AmbiantLight_stagedOrder[instance]
	case *BoxGeometry:
		return stage.BoxGeometry_stagedOrder[instance]
	case *BufferGeometry:
		return stage.BufferGeometry_stagedOrder[instance]
	case *Camera:
		return stage.Camera_stagedOrder[instance]
	case *Canvas:
		return stage.Canvas_stagedOrder[instance]
	case *Curve:
		return stage.Curve_stagedOrder[instance]
	case *CylinderGeometry:
		return stage.CylinderGeometry_stagedOrder[instance]
	case *DirectionalLight:
		return stage.DirectionalLight_stagedOrder[instance]
	case *ExtrudeGeometry:
		return stage.ExtrudeGeometry_stagedOrder[instance]
	case *Mesh:
		return stage.Mesh_stagedOrder[instance]
	case *MeshMaterialBasic:
		return stage.MeshMaterialBasic_stagedOrder[instance]
	case *MeshPhysicalMaterial:
		return stage.MeshPhysicalMaterial_stagedOrder[instance]
	case *PlaneGeometry:
		return stage.PlaneGeometry_stagedOrder[instance]
	case *Shape:
		return stage.Shape_stagedOrder[instance]
	case *SphereGeometry:
		return stage.SphereGeometry_stagedOrder[instance]
	case *TorusGeometry:
		return stage.TorusGeometry_stagedOrder[instance]
	case *Triangle:
		return stage.Triangle_stagedOrder[instance]
	case *TubeGeometry:
		return stage.TubeGeometry_stagedOrder[instance]
	case *Vector2:
		return stage.Vector2_stagedOrder[instance]
	case *Vector3:
		return stage.Vector3_stagedOrder[instance]
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
	if _, ok := stage.AmbiantLights[ambiantlight]; !ok {
		stage.AmbiantLights[ambiantlight] = struct{}{}
		stage.AmbiantLight_stagedOrder[ambiantlight] = stage.AmbiantLightOrder
		stage.AmbiantLight_orderStaged[stage.AmbiantLightOrder] = ambiantlight
		stage.AmbiantLightOrder++
	}
	stage.AmbiantLights_mapString[ambiantlight.Name] = ambiantlight

	return ambiantlight
}

// StagePreserveOrder puts ambiantlight to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AmbiantLightOrder
// - update stage.AmbiantLightOrder accordingly
func (ambiantlight *AmbiantLight) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.AmbiantLights[ambiantlight]; !ok {
		stage.AmbiantLights[ambiantlight] = struct{}{}

		if order > stage.AmbiantLightOrder {
			stage.AmbiantLightOrder = order
		}
		stage.AmbiantLight_stagedOrder[ambiantlight] = order
		stage.AmbiantLight_orderStaged[order] = ambiantlight
		stage.AmbiantLightOrder++
	}
	stage.AmbiantLights_mapString[ambiantlight.Name] = ambiantlight
}

// Unstage removes ambiantlight off the model stage
func (ambiantlight *AmbiantLight) Unstage(stage *Stage) *AmbiantLight {
	delete(stage.AmbiantLights, ambiantlight)
	// issue1150
	// delete(stage.AmbiantLight_stagedOrder, ambiantlight)
	delete(stage.AmbiantLights_mapString, ambiantlight.Name)

	return ambiantlight
}

// UnstageVoid removes ambiantlight off the model stage
func (ambiantlight *AmbiantLight) UnstageVoid(stage *Stage) {
	delete(stage.AmbiantLights, ambiantlight)
	// issue1150
	// delete(stage.AmbiantLight_stagedOrder, ambiantlight)
	delete(stage.AmbiantLights_mapString, ambiantlight.Name)
}

// commit ambiantlight to the back repo (if it is already staged)
func (ambiantlight *AmbiantLight) Commit(stage *Stage) *AmbiantLight {
	if _, ok := stage.AmbiantLights[ambiantlight]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAmbiantLight(ambiantlight)
		}
	}
	return ambiantlight
}

func (ambiantlight *AmbiantLight) CommitVoid(stage *Stage) {
	ambiantlight.Commit(stage)
}

func (ambiantlight *AmbiantLight) StageVoid(stage *Stage) {
	ambiantlight.Stage(stage)
}

// Checkout ambiantlight to the back repo (if it is already staged)
func (ambiantlight *AmbiantLight) Checkout(stage *Stage) *AmbiantLight {
	if _, ok := stage.AmbiantLights[ambiantlight]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAmbiantLight(ambiantlight)
		}
	}
	return ambiantlight
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
	if _, ok := stage.BoxGeometrys[boxgeometry]; !ok {
		stage.BoxGeometrys[boxgeometry] = struct{}{}
		stage.BoxGeometry_stagedOrder[boxgeometry] = stage.BoxGeometryOrder
		stage.BoxGeometry_orderStaged[stage.BoxGeometryOrder] = boxgeometry
		stage.BoxGeometryOrder++
	}
	stage.BoxGeometrys_mapString[boxgeometry.Name] = boxgeometry

	return boxgeometry
}

// StagePreserveOrder puts boxgeometry to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BoxGeometryOrder
// - update stage.BoxGeometryOrder accordingly
func (boxgeometry *BoxGeometry) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.BoxGeometrys[boxgeometry]; !ok {
		stage.BoxGeometrys[boxgeometry] = struct{}{}

		if order > stage.BoxGeometryOrder {
			stage.BoxGeometryOrder = order
		}
		stage.BoxGeometry_stagedOrder[boxgeometry] = order
		stage.BoxGeometry_orderStaged[order] = boxgeometry
		stage.BoxGeometryOrder++
	}
	stage.BoxGeometrys_mapString[boxgeometry.Name] = boxgeometry
}

// Unstage removes boxgeometry off the model stage
func (boxgeometry *BoxGeometry) Unstage(stage *Stage) *BoxGeometry {
	delete(stage.BoxGeometrys, boxgeometry)
	// issue1150
	// delete(stage.BoxGeometry_stagedOrder, boxgeometry)
	delete(stage.BoxGeometrys_mapString, boxgeometry.Name)

	return boxgeometry
}

// UnstageVoid removes boxgeometry off the model stage
func (boxgeometry *BoxGeometry) UnstageVoid(stage *Stage) {
	delete(stage.BoxGeometrys, boxgeometry)
	// issue1150
	// delete(stage.BoxGeometry_stagedOrder, boxgeometry)
	delete(stage.BoxGeometrys_mapString, boxgeometry.Name)
}

// commit boxgeometry to the back repo (if it is already staged)
func (boxgeometry *BoxGeometry) Commit(stage *Stage) *BoxGeometry {
	if _, ok := stage.BoxGeometrys[boxgeometry]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBoxGeometry(boxgeometry)
		}
	}
	return boxgeometry
}

func (boxgeometry *BoxGeometry) CommitVoid(stage *Stage) {
	boxgeometry.Commit(stage)
}

func (boxgeometry *BoxGeometry) StageVoid(stage *Stage) {
	boxgeometry.Stage(stage)
}

// Checkout boxgeometry to the back repo (if it is already staged)
func (boxgeometry *BoxGeometry) Checkout(stage *Stage) *BoxGeometry {
	if _, ok := stage.BoxGeometrys[boxgeometry]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBoxGeometry(boxgeometry)
		}
	}
	return boxgeometry
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
	if _, ok := stage.BufferGeometrys[buffergeometry]; !ok {
		stage.BufferGeometrys[buffergeometry] = struct{}{}
		stage.BufferGeometry_stagedOrder[buffergeometry] = stage.BufferGeometryOrder
		stage.BufferGeometry_orderStaged[stage.BufferGeometryOrder] = buffergeometry
		stage.BufferGeometryOrder++
	}
	stage.BufferGeometrys_mapString[buffergeometry.Name] = buffergeometry

	return buffergeometry
}

// StagePreserveOrder puts buffergeometry to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BufferGeometryOrder
// - update stage.BufferGeometryOrder accordingly
func (buffergeometry *BufferGeometry) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.BufferGeometrys[buffergeometry]; !ok {
		stage.BufferGeometrys[buffergeometry] = struct{}{}

		if order > stage.BufferGeometryOrder {
			stage.BufferGeometryOrder = order
		}
		stage.BufferGeometry_stagedOrder[buffergeometry] = order
		stage.BufferGeometry_orderStaged[order] = buffergeometry
		stage.BufferGeometryOrder++
	}
	stage.BufferGeometrys_mapString[buffergeometry.Name] = buffergeometry
}

// Unstage removes buffergeometry off the model stage
func (buffergeometry *BufferGeometry) Unstage(stage *Stage) *BufferGeometry {
	delete(stage.BufferGeometrys, buffergeometry)
	// issue1150
	// delete(stage.BufferGeometry_stagedOrder, buffergeometry)
	delete(stage.BufferGeometrys_mapString, buffergeometry.Name)

	return buffergeometry
}

// UnstageVoid removes buffergeometry off the model stage
func (buffergeometry *BufferGeometry) UnstageVoid(stage *Stage) {
	delete(stage.BufferGeometrys, buffergeometry)
	// issue1150
	// delete(stage.BufferGeometry_stagedOrder, buffergeometry)
	delete(stage.BufferGeometrys_mapString, buffergeometry.Name)
}

// commit buffergeometry to the back repo (if it is already staged)
func (buffergeometry *BufferGeometry) Commit(stage *Stage) *BufferGeometry {
	if _, ok := stage.BufferGeometrys[buffergeometry]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBufferGeometry(buffergeometry)
		}
	}
	return buffergeometry
}

func (buffergeometry *BufferGeometry) CommitVoid(stage *Stage) {
	buffergeometry.Commit(stage)
}

func (buffergeometry *BufferGeometry) StageVoid(stage *Stage) {
	buffergeometry.Stage(stage)
}

// Checkout buffergeometry to the back repo (if it is already staged)
func (buffergeometry *BufferGeometry) Checkout(stage *Stage) *BufferGeometry {
	if _, ok := stage.BufferGeometrys[buffergeometry]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBufferGeometry(buffergeometry)
		}
	}
	return buffergeometry
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
	if _, ok := stage.Cameras[camera]; !ok {
		stage.Cameras[camera] = struct{}{}
		stage.Camera_stagedOrder[camera] = stage.CameraOrder
		stage.Camera_orderStaged[stage.CameraOrder] = camera
		stage.CameraOrder++
	}
	stage.Cameras_mapString[camera.Name] = camera

	return camera
}

// StagePreserveOrder puts camera to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CameraOrder
// - update stage.CameraOrder accordingly
func (camera *Camera) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Cameras[camera]; !ok {
		stage.Cameras[camera] = struct{}{}

		if order > stage.CameraOrder {
			stage.CameraOrder = order
		}
		stage.Camera_stagedOrder[camera] = order
		stage.Camera_orderStaged[order] = camera
		stage.CameraOrder++
	}
	stage.Cameras_mapString[camera.Name] = camera
}

// Unstage removes camera off the model stage
func (camera *Camera) Unstage(stage *Stage) *Camera {
	delete(stage.Cameras, camera)
	// issue1150
	// delete(stage.Camera_stagedOrder, camera)
	delete(stage.Cameras_mapString, camera.Name)

	return camera
}

// UnstageVoid removes camera off the model stage
func (camera *Camera) UnstageVoid(stage *Stage) {
	delete(stage.Cameras, camera)
	// issue1150
	// delete(stage.Camera_stagedOrder, camera)
	delete(stage.Cameras_mapString, camera.Name)
}

// commit camera to the back repo (if it is already staged)
func (camera *Camera) Commit(stage *Stage) *Camera {
	if _, ok := stage.Cameras[camera]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCamera(camera)
		}
	}
	return camera
}

func (camera *Camera) CommitVoid(stage *Stage) {
	camera.Commit(stage)
}

func (camera *Camera) StageVoid(stage *Stage) {
	camera.Stage(stage)
}

// Checkout camera to the back repo (if it is already staged)
func (camera *Camera) Checkout(stage *Stage) *Camera {
	if _, ok := stage.Cameras[camera]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCamera(camera)
		}
	}
	return camera
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
	if _, ok := stage.Canvass[canvas]; !ok {
		stage.Canvass[canvas] = struct{}{}
		stage.Canvas_stagedOrder[canvas] = stage.CanvasOrder
		stage.Canvas_orderStaged[stage.CanvasOrder] = canvas
		stage.CanvasOrder++
	}
	stage.Canvass_mapString[canvas.Name] = canvas

	return canvas
}

// StagePreserveOrder puts canvas to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CanvasOrder
// - update stage.CanvasOrder accordingly
func (canvas *Canvas) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Canvass[canvas]; !ok {
		stage.Canvass[canvas] = struct{}{}

		if order > stage.CanvasOrder {
			stage.CanvasOrder = order
		}
		stage.Canvas_stagedOrder[canvas] = order
		stage.Canvas_orderStaged[order] = canvas
		stage.CanvasOrder++
	}
	stage.Canvass_mapString[canvas.Name] = canvas
}

// Unstage removes canvas off the model stage
func (canvas *Canvas) Unstage(stage *Stage) *Canvas {
	delete(stage.Canvass, canvas)
	// issue1150
	// delete(stage.Canvas_stagedOrder, canvas)
	delete(stage.Canvass_mapString, canvas.Name)

	return canvas
}

// UnstageVoid removes canvas off the model stage
func (canvas *Canvas) UnstageVoid(stage *Stage) {
	delete(stage.Canvass, canvas)
	// issue1150
	// delete(stage.Canvas_stagedOrder, canvas)
	delete(stage.Canvass_mapString, canvas.Name)
}

// commit canvas to the back repo (if it is already staged)
func (canvas *Canvas) Commit(stage *Stage) *Canvas {
	if _, ok := stage.Canvass[canvas]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCanvas(canvas)
		}
	}
	return canvas
}

func (canvas *Canvas) CommitVoid(stage *Stage) {
	canvas.Commit(stage)
}

func (canvas *Canvas) StageVoid(stage *Stage) {
	canvas.Stage(stage)
}

// Checkout canvas to the back repo (if it is already staged)
func (canvas *Canvas) Checkout(stage *Stage) *Canvas {
	if _, ok := stage.Canvass[canvas]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCanvas(canvas)
		}
	}
	return canvas
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
	if _, ok := stage.Curves[curve]; !ok {
		stage.Curves[curve] = struct{}{}
		stage.Curve_stagedOrder[curve] = stage.CurveOrder
		stage.Curve_orderStaged[stage.CurveOrder] = curve
		stage.CurveOrder++
	}
	stage.Curves_mapString[curve.Name] = curve

	return curve
}

// StagePreserveOrder puts curve to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CurveOrder
// - update stage.CurveOrder accordingly
func (curve *Curve) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Curves[curve]; !ok {
		stage.Curves[curve] = struct{}{}

		if order > stage.CurveOrder {
			stage.CurveOrder = order
		}
		stage.Curve_stagedOrder[curve] = order
		stage.Curve_orderStaged[order] = curve
		stage.CurveOrder++
	}
	stage.Curves_mapString[curve.Name] = curve
}

// Unstage removes curve off the model stage
func (curve *Curve) Unstage(stage *Stage) *Curve {
	delete(stage.Curves, curve)
	// issue1150
	// delete(stage.Curve_stagedOrder, curve)
	delete(stage.Curves_mapString, curve.Name)

	return curve
}

// UnstageVoid removes curve off the model stage
func (curve *Curve) UnstageVoid(stage *Stage) {
	delete(stage.Curves, curve)
	// issue1150
	// delete(stage.Curve_stagedOrder, curve)
	delete(stage.Curves_mapString, curve.Name)
}

// commit curve to the back repo (if it is already staged)
func (curve *Curve) Commit(stage *Stage) *Curve {
	if _, ok := stage.Curves[curve]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCurve(curve)
		}
	}
	return curve
}

func (curve *Curve) CommitVoid(stage *Stage) {
	curve.Commit(stage)
}

func (curve *Curve) StageVoid(stage *Stage) {
	curve.Stage(stage)
}

// Checkout curve to the back repo (if it is already staged)
func (curve *Curve) Checkout(stage *Stage) *Curve {
	if _, ok := stage.Curves[curve]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCurve(curve)
		}
	}
	return curve
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
	if _, ok := stage.CylinderGeometrys[cylindergeometry]; !ok {
		stage.CylinderGeometrys[cylindergeometry] = struct{}{}
		stage.CylinderGeometry_stagedOrder[cylindergeometry] = stage.CylinderGeometryOrder
		stage.CylinderGeometry_orderStaged[stage.CylinderGeometryOrder] = cylindergeometry
		stage.CylinderGeometryOrder++
	}
	stage.CylinderGeometrys_mapString[cylindergeometry.Name] = cylindergeometry

	return cylindergeometry
}

// StagePreserveOrder puts cylindergeometry to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CylinderGeometryOrder
// - update stage.CylinderGeometryOrder accordingly
func (cylindergeometry *CylinderGeometry) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.CylinderGeometrys[cylindergeometry]; !ok {
		stage.CylinderGeometrys[cylindergeometry] = struct{}{}

		if order > stage.CylinderGeometryOrder {
			stage.CylinderGeometryOrder = order
		}
		stage.CylinderGeometry_stagedOrder[cylindergeometry] = order
		stage.CylinderGeometry_orderStaged[order] = cylindergeometry
		stage.CylinderGeometryOrder++
	}
	stage.CylinderGeometrys_mapString[cylindergeometry.Name] = cylindergeometry
}

// Unstage removes cylindergeometry off the model stage
func (cylindergeometry *CylinderGeometry) Unstage(stage *Stage) *CylinderGeometry {
	delete(stage.CylinderGeometrys, cylindergeometry)
	// issue1150
	// delete(stage.CylinderGeometry_stagedOrder, cylindergeometry)
	delete(stage.CylinderGeometrys_mapString, cylindergeometry.Name)

	return cylindergeometry
}

// UnstageVoid removes cylindergeometry off the model stage
func (cylindergeometry *CylinderGeometry) UnstageVoid(stage *Stage) {
	delete(stage.CylinderGeometrys, cylindergeometry)
	// issue1150
	// delete(stage.CylinderGeometry_stagedOrder, cylindergeometry)
	delete(stage.CylinderGeometrys_mapString, cylindergeometry.Name)
}

// commit cylindergeometry to the back repo (if it is already staged)
func (cylindergeometry *CylinderGeometry) Commit(stage *Stage) *CylinderGeometry {
	if _, ok := stage.CylinderGeometrys[cylindergeometry]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCylinderGeometry(cylindergeometry)
		}
	}
	return cylindergeometry
}

func (cylindergeometry *CylinderGeometry) CommitVoid(stage *Stage) {
	cylindergeometry.Commit(stage)
}

func (cylindergeometry *CylinderGeometry) StageVoid(stage *Stage) {
	cylindergeometry.Stage(stage)
}

// Checkout cylindergeometry to the back repo (if it is already staged)
func (cylindergeometry *CylinderGeometry) Checkout(stage *Stage) *CylinderGeometry {
	if _, ok := stage.CylinderGeometrys[cylindergeometry]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCylinderGeometry(cylindergeometry)
		}
	}
	return cylindergeometry
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
	if _, ok := stage.DirectionalLights[directionallight]; !ok {
		stage.DirectionalLights[directionallight] = struct{}{}
		stage.DirectionalLight_stagedOrder[directionallight] = stage.DirectionalLightOrder
		stage.DirectionalLight_orderStaged[stage.DirectionalLightOrder] = directionallight
		stage.DirectionalLightOrder++
	}
	stage.DirectionalLights_mapString[directionallight.Name] = directionallight

	return directionallight
}

// StagePreserveOrder puts directionallight to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DirectionalLightOrder
// - update stage.DirectionalLightOrder accordingly
func (directionallight *DirectionalLight) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.DirectionalLights[directionallight]; !ok {
		stage.DirectionalLights[directionallight] = struct{}{}

		if order > stage.DirectionalLightOrder {
			stage.DirectionalLightOrder = order
		}
		stage.DirectionalLight_stagedOrder[directionallight] = order
		stage.DirectionalLight_orderStaged[order] = directionallight
		stage.DirectionalLightOrder++
	}
	stage.DirectionalLights_mapString[directionallight.Name] = directionallight
}

// Unstage removes directionallight off the model stage
func (directionallight *DirectionalLight) Unstage(stage *Stage) *DirectionalLight {
	delete(stage.DirectionalLights, directionallight)
	// issue1150
	// delete(stage.DirectionalLight_stagedOrder, directionallight)
	delete(stage.DirectionalLights_mapString, directionallight.Name)

	return directionallight
}

// UnstageVoid removes directionallight off the model stage
func (directionallight *DirectionalLight) UnstageVoid(stage *Stage) {
	delete(stage.DirectionalLights, directionallight)
	// issue1150
	// delete(stage.DirectionalLight_stagedOrder, directionallight)
	delete(stage.DirectionalLights_mapString, directionallight.Name)
}

// commit directionallight to the back repo (if it is already staged)
func (directionallight *DirectionalLight) Commit(stage *Stage) *DirectionalLight {
	if _, ok := stage.DirectionalLights[directionallight]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDirectionalLight(directionallight)
		}
	}
	return directionallight
}

func (directionallight *DirectionalLight) CommitVoid(stage *Stage) {
	directionallight.Commit(stage)
}

func (directionallight *DirectionalLight) StageVoid(stage *Stage) {
	directionallight.Stage(stage)
}

// Checkout directionallight to the back repo (if it is already staged)
func (directionallight *DirectionalLight) Checkout(stage *Stage) *DirectionalLight {
	if _, ok := stage.DirectionalLights[directionallight]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDirectionalLight(directionallight)
		}
	}
	return directionallight
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
	if _, ok := stage.ExtrudeGeometrys[extrudegeometry]; !ok {
		stage.ExtrudeGeometrys[extrudegeometry] = struct{}{}
		stage.ExtrudeGeometry_stagedOrder[extrudegeometry] = stage.ExtrudeGeometryOrder
		stage.ExtrudeGeometry_orderStaged[stage.ExtrudeGeometryOrder] = extrudegeometry
		stage.ExtrudeGeometryOrder++
	}
	stage.ExtrudeGeometrys_mapString[extrudegeometry.Name] = extrudegeometry

	return extrudegeometry
}

// StagePreserveOrder puts extrudegeometry to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ExtrudeGeometryOrder
// - update stage.ExtrudeGeometryOrder accordingly
func (extrudegeometry *ExtrudeGeometry) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ExtrudeGeometrys[extrudegeometry]; !ok {
		stage.ExtrudeGeometrys[extrudegeometry] = struct{}{}

		if order > stage.ExtrudeGeometryOrder {
			stage.ExtrudeGeometryOrder = order
		}
		stage.ExtrudeGeometry_stagedOrder[extrudegeometry] = order
		stage.ExtrudeGeometry_orderStaged[order] = extrudegeometry
		stage.ExtrudeGeometryOrder++
	}
	stage.ExtrudeGeometrys_mapString[extrudegeometry.Name] = extrudegeometry
}

// Unstage removes extrudegeometry off the model stage
func (extrudegeometry *ExtrudeGeometry) Unstage(stage *Stage) *ExtrudeGeometry {
	delete(stage.ExtrudeGeometrys, extrudegeometry)
	// issue1150
	// delete(stage.ExtrudeGeometry_stagedOrder, extrudegeometry)
	delete(stage.ExtrudeGeometrys_mapString, extrudegeometry.Name)

	return extrudegeometry
}

// UnstageVoid removes extrudegeometry off the model stage
func (extrudegeometry *ExtrudeGeometry) UnstageVoid(stage *Stage) {
	delete(stage.ExtrudeGeometrys, extrudegeometry)
	// issue1150
	// delete(stage.ExtrudeGeometry_stagedOrder, extrudegeometry)
	delete(stage.ExtrudeGeometrys_mapString, extrudegeometry.Name)
}

// commit extrudegeometry to the back repo (if it is already staged)
func (extrudegeometry *ExtrudeGeometry) Commit(stage *Stage) *ExtrudeGeometry {
	if _, ok := stage.ExtrudeGeometrys[extrudegeometry]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitExtrudeGeometry(extrudegeometry)
		}
	}
	return extrudegeometry
}

func (extrudegeometry *ExtrudeGeometry) CommitVoid(stage *Stage) {
	extrudegeometry.Commit(stage)
}

func (extrudegeometry *ExtrudeGeometry) StageVoid(stage *Stage) {
	extrudegeometry.Stage(stage)
}

// Checkout extrudegeometry to the back repo (if it is already staged)
func (extrudegeometry *ExtrudeGeometry) Checkout(stage *Stage) *ExtrudeGeometry {
	if _, ok := stage.ExtrudeGeometrys[extrudegeometry]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutExtrudeGeometry(extrudegeometry)
		}
	}
	return extrudegeometry
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
	if _, ok := stage.Meshs[mesh]; !ok {
		stage.Meshs[mesh] = struct{}{}
		stage.Mesh_stagedOrder[mesh] = stage.MeshOrder
		stage.Mesh_orderStaged[stage.MeshOrder] = mesh
		stage.MeshOrder++
	}
	stage.Meshs_mapString[mesh.Name] = mesh

	return mesh
}

// StagePreserveOrder puts mesh to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MeshOrder
// - update stage.MeshOrder accordingly
func (mesh *Mesh) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Meshs[mesh]; !ok {
		stage.Meshs[mesh] = struct{}{}

		if order > stage.MeshOrder {
			stage.MeshOrder = order
		}
		stage.Mesh_stagedOrder[mesh] = order
		stage.Mesh_orderStaged[order] = mesh
		stage.MeshOrder++
	}
	stage.Meshs_mapString[mesh.Name] = mesh
}

// Unstage removes mesh off the model stage
func (mesh *Mesh) Unstage(stage *Stage) *Mesh {
	delete(stage.Meshs, mesh)
	// issue1150
	// delete(stage.Mesh_stagedOrder, mesh)
	delete(stage.Meshs_mapString, mesh.Name)

	return mesh
}

// UnstageVoid removes mesh off the model stage
func (mesh *Mesh) UnstageVoid(stage *Stage) {
	delete(stage.Meshs, mesh)
	// issue1150
	// delete(stage.Mesh_stagedOrder, mesh)
	delete(stage.Meshs_mapString, mesh.Name)
}

// commit mesh to the back repo (if it is already staged)
func (mesh *Mesh) Commit(stage *Stage) *Mesh {
	if _, ok := stage.Meshs[mesh]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMesh(mesh)
		}
	}
	return mesh
}

func (mesh *Mesh) CommitVoid(stage *Stage) {
	mesh.Commit(stage)
}

func (mesh *Mesh) StageVoid(stage *Stage) {
	mesh.Stage(stage)
}

// Checkout mesh to the back repo (if it is already staged)
func (mesh *Mesh) Checkout(stage *Stage) *Mesh {
	if _, ok := stage.Meshs[mesh]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMesh(mesh)
		}
	}
	return mesh
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
	if _, ok := stage.MeshMaterialBasics[meshmaterialbasic]; !ok {
		stage.MeshMaterialBasics[meshmaterialbasic] = struct{}{}
		stage.MeshMaterialBasic_stagedOrder[meshmaterialbasic] = stage.MeshMaterialBasicOrder
		stage.MeshMaterialBasic_orderStaged[stage.MeshMaterialBasicOrder] = meshmaterialbasic
		stage.MeshMaterialBasicOrder++
	}
	stage.MeshMaterialBasics_mapString[meshmaterialbasic.Name] = meshmaterialbasic

	return meshmaterialbasic
}

// StagePreserveOrder puts meshmaterialbasic to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MeshMaterialBasicOrder
// - update stage.MeshMaterialBasicOrder accordingly
func (meshmaterialbasic *MeshMaterialBasic) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.MeshMaterialBasics[meshmaterialbasic]; !ok {
		stage.MeshMaterialBasics[meshmaterialbasic] = struct{}{}

		if order > stage.MeshMaterialBasicOrder {
			stage.MeshMaterialBasicOrder = order
		}
		stage.MeshMaterialBasic_stagedOrder[meshmaterialbasic] = order
		stage.MeshMaterialBasic_orderStaged[order] = meshmaterialbasic
		stage.MeshMaterialBasicOrder++
	}
	stage.MeshMaterialBasics_mapString[meshmaterialbasic.Name] = meshmaterialbasic
}

// Unstage removes meshmaterialbasic off the model stage
func (meshmaterialbasic *MeshMaterialBasic) Unstage(stage *Stage) *MeshMaterialBasic {
	delete(stage.MeshMaterialBasics, meshmaterialbasic)
	// issue1150
	// delete(stage.MeshMaterialBasic_stagedOrder, meshmaterialbasic)
	delete(stage.MeshMaterialBasics_mapString, meshmaterialbasic.Name)

	return meshmaterialbasic
}

// UnstageVoid removes meshmaterialbasic off the model stage
func (meshmaterialbasic *MeshMaterialBasic) UnstageVoid(stage *Stage) {
	delete(stage.MeshMaterialBasics, meshmaterialbasic)
	// issue1150
	// delete(stage.MeshMaterialBasic_stagedOrder, meshmaterialbasic)
	delete(stage.MeshMaterialBasics_mapString, meshmaterialbasic.Name)
}

// commit meshmaterialbasic to the back repo (if it is already staged)
func (meshmaterialbasic *MeshMaterialBasic) Commit(stage *Stage) *MeshMaterialBasic {
	if _, ok := stage.MeshMaterialBasics[meshmaterialbasic]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMeshMaterialBasic(meshmaterialbasic)
		}
	}
	return meshmaterialbasic
}

func (meshmaterialbasic *MeshMaterialBasic) CommitVoid(stage *Stage) {
	meshmaterialbasic.Commit(stage)
}

func (meshmaterialbasic *MeshMaterialBasic) StageVoid(stage *Stage) {
	meshmaterialbasic.Stage(stage)
}

// Checkout meshmaterialbasic to the back repo (if it is already staged)
func (meshmaterialbasic *MeshMaterialBasic) Checkout(stage *Stage) *MeshMaterialBasic {
	if _, ok := stage.MeshMaterialBasics[meshmaterialbasic]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMeshMaterialBasic(meshmaterialbasic)
		}
	}
	return meshmaterialbasic
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
	if _, ok := stage.MeshPhysicalMaterials[meshphysicalmaterial]; !ok {
		stage.MeshPhysicalMaterials[meshphysicalmaterial] = struct{}{}
		stage.MeshPhysicalMaterial_stagedOrder[meshphysicalmaterial] = stage.MeshPhysicalMaterialOrder
		stage.MeshPhysicalMaterial_orderStaged[stage.MeshPhysicalMaterialOrder] = meshphysicalmaterial
		stage.MeshPhysicalMaterialOrder++
	}
	stage.MeshPhysicalMaterials_mapString[meshphysicalmaterial.Name] = meshphysicalmaterial

	return meshphysicalmaterial
}

// StagePreserveOrder puts meshphysicalmaterial to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MeshPhysicalMaterialOrder
// - update stage.MeshPhysicalMaterialOrder accordingly
func (meshphysicalmaterial *MeshPhysicalMaterial) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.MeshPhysicalMaterials[meshphysicalmaterial]; !ok {
		stage.MeshPhysicalMaterials[meshphysicalmaterial] = struct{}{}

		if order > stage.MeshPhysicalMaterialOrder {
			stage.MeshPhysicalMaterialOrder = order
		}
		stage.MeshPhysicalMaterial_stagedOrder[meshphysicalmaterial] = order
		stage.MeshPhysicalMaterial_orderStaged[order] = meshphysicalmaterial
		stage.MeshPhysicalMaterialOrder++
	}
	stage.MeshPhysicalMaterials_mapString[meshphysicalmaterial.Name] = meshphysicalmaterial
}

// Unstage removes meshphysicalmaterial off the model stage
func (meshphysicalmaterial *MeshPhysicalMaterial) Unstage(stage *Stage) *MeshPhysicalMaterial {
	delete(stage.MeshPhysicalMaterials, meshphysicalmaterial)
	// issue1150
	// delete(stage.MeshPhysicalMaterial_stagedOrder, meshphysicalmaterial)
	delete(stage.MeshPhysicalMaterials_mapString, meshphysicalmaterial.Name)

	return meshphysicalmaterial
}

// UnstageVoid removes meshphysicalmaterial off the model stage
func (meshphysicalmaterial *MeshPhysicalMaterial) UnstageVoid(stage *Stage) {
	delete(stage.MeshPhysicalMaterials, meshphysicalmaterial)
	// issue1150
	// delete(stage.MeshPhysicalMaterial_stagedOrder, meshphysicalmaterial)
	delete(stage.MeshPhysicalMaterials_mapString, meshphysicalmaterial.Name)
}

// commit meshphysicalmaterial to the back repo (if it is already staged)
func (meshphysicalmaterial *MeshPhysicalMaterial) Commit(stage *Stage) *MeshPhysicalMaterial {
	if _, ok := stage.MeshPhysicalMaterials[meshphysicalmaterial]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMeshPhysicalMaterial(meshphysicalmaterial)
		}
	}
	return meshphysicalmaterial
}

func (meshphysicalmaterial *MeshPhysicalMaterial) CommitVoid(stage *Stage) {
	meshphysicalmaterial.Commit(stage)
}

func (meshphysicalmaterial *MeshPhysicalMaterial) StageVoid(stage *Stage) {
	meshphysicalmaterial.Stage(stage)
}

// Checkout meshphysicalmaterial to the back repo (if it is already staged)
func (meshphysicalmaterial *MeshPhysicalMaterial) Checkout(stage *Stage) *MeshPhysicalMaterial {
	if _, ok := stage.MeshPhysicalMaterials[meshphysicalmaterial]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMeshPhysicalMaterial(meshphysicalmaterial)
		}
	}
	return meshphysicalmaterial
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
	if _, ok := stage.PlaneGeometrys[planegeometry]; !ok {
		stage.PlaneGeometrys[planegeometry] = struct{}{}
		stage.PlaneGeometry_stagedOrder[planegeometry] = stage.PlaneGeometryOrder
		stage.PlaneGeometry_orderStaged[stage.PlaneGeometryOrder] = planegeometry
		stage.PlaneGeometryOrder++
	}
	stage.PlaneGeometrys_mapString[planegeometry.Name] = planegeometry

	return planegeometry
}

// StagePreserveOrder puts planegeometry to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PlaneGeometryOrder
// - update stage.PlaneGeometryOrder accordingly
func (planegeometry *PlaneGeometry) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.PlaneGeometrys[planegeometry]; !ok {
		stage.PlaneGeometrys[planegeometry] = struct{}{}

		if order > stage.PlaneGeometryOrder {
			stage.PlaneGeometryOrder = order
		}
		stage.PlaneGeometry_stagedOrder[planegeometry] = order
		stage.PlaneGeometry_orderStaged[order] = planegeometry
		stage.PlaneGeometryOrder++
	}
	stage.PlaneGeometrys_mapString[planegeometry.Name] = planegeometry
}

// Unstage removes planegeometry off the model stage
func (planegeometry *PlaneGeometry) Unstage(stage *Stage) *PlaneGeometry {
	delete(stage.PlaneGeometrys, planegeometry)
	// issue1150
	// delete(stage.PlaneGeometry_stagedOrder, planegeometry)
	delete(stage.PlaneGeometrys_mapString, planegeometry.Name)

	return planegeometry
}

// UnstageVoid removes planegeometry off the model stage
func (planegeometry *PlaneGeometry) UnstageVoid(stage *Stage) {
	delete(stage.PlaneGeometrys, planegeometry)
	// issue1150
	// delete(stage.PlaneGeometry_stagedOrder, planegeometry)
	delete(stage.PlaneGeometrys_mapString, planegeometry.Name)
}

// commit planegeometry to the back repo (if it is already staged)
func (planegeometry *PlaneGeometry) Commit(stage *Stage) *PlaneGeometry {
	if _, ok := stage.PlaneGeometrys[planegeometry]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPlaneGeometry(planegeometry)
		}
	}
	return planegeometry
}

func (planegeometry *PlaneGeometry) CommitVoid(stage *Stage) {
	planegeometry.Commit(stage)
}

func (planegeometry *PlaneGeometry) StageVoid(stage *Stage) {
	planegeometry.Stage(stage)
}

// Checkout planegeometry to the back repo (if it is already staged)
func (planegeometry *PlaneGeometry) Checkout(stage *Stage) *PlaneGeometry {
	if _, ok := stage.PlaneGeometrys[planegeometry]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPlaneGeometry(planegeometry)
		}
	}
	return planegeometry
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
	if _, ok := stage.Shapes[shape]; !ok {
		stage.Shapes[shape] = struct{}{}
		stage.Shape_stagedOrder[shape] = stage.ShapeOrder
		stage.Shape_orderStaged[stage.ShapeOrder] = shape
		stage.ShapeOrder++
	}
	stage.Shapes_mapString[shape.Name] = shape

	return shape
}

// StagePreserveOrder puts shape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ShapeOrder
// - update stage.ShapeOrder accordingly
func (shape *Shape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Shapes[shape]; !ok {
		stage.Shapes[shape] = struct{}{}

		if order > stage.ShapeOrder {
			stage.ShapeOrder = order
		}
		stage.Shape_stagedOrder[shape] = order
		stage.Shape_orderStaged[order] = shape
		stage.ShapeOrder++
	}
	stage.Shapes_mapString[shape.Name] = shape
}

// Unstage removes shape off the model stage
func (shape *Shape) Unstage(stage *Stage) *Shape {
	delete(stage.Shapes, shape)
	// issue1150
	// delete(stage.Shape_stagedOrder, shape)
	delete(stage.Shapes_mapString, shape.Name)

	return shape
}

// UnstageVoid removes shape off the model stage
func (shape *Shape) UnstageVoid(stage *Stage) {
	delete(stage.Shapes, shape)
	// issue1150
	// delete(stage.Shape_stagedOrder, shape)
	delete(stage.Shapes_mapString, shape.Name)
}

// commit shape to the back repo (if it is already staged)
func (shape *Shape) Commit(stage *Stage) *Shape {
	if _, ok := stage.Shapes[shape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitShape(shape)
		}
	}
	return shape
}

func (shape *Shape) CommitVoid(stage *Stage) {
	shape.Commit(stage)
}

func (shape *Shape) StageVoid(stage *Stage) {
	shape.Stage(stage)
}

// Checkout shape to the back repo (if it is already staged)
func (shape *Shape) Checkout(stage *Stage) *Shape {
	if _, ok := stage.Shapes[shape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutShape(shape)
		}
	}
	return shape
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
	if _, ok := stage.SphereGeometrys[spheregeometry]; !ok {
		stage.SphereGeometrys[spheregeometry] = struct{}{}
		stage.SphereGeometry_stagedOrder[spheregeometry] = stage.SphereGeometryOrder
		stage.SphereGeometry_orderStaged[stage.SphereGeometryOrder] = spheregeometry
		stage.SphereGeometryOrder++
	}
	stage.SphereGeometrys_mapString[spheregeometry.Name] = spheregeometry

	return spheregeometry
}

// StagePreserveOrder puts spheregeometry to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SphereGeometryOrder
// - update stage.SphereGeometryOrder accordingly
func (spheregeometry *SphereGeometry) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.SphereGeometrys[spheregeometry]; !ok {
		stage.SphereGeometrys[spheregeometry] = struct{}{}

		if order > stage.SphereGeometryOrder {
			stage.SphereGeometryOrder = order
		}
		stage.SphereGeometry_stagedOrder[spheregeometry] = order
		stage.SphereGeometry_orderStaged[order] = spheregeometry
		stage.SphereGeometryOrder++
	}
	stage.SphereGeometrys_mapString[spheregeometry.Name] = spheregeometry
}

// Unstage removes spheregeometry off the model stage
func (spheregeometry *SphereGeometry) Unstage(stage *Stage) *SphereGeometry {
	delete(stage.SphereGeometrys, spheregeometry)
	// issue1150
	// delete(stage.SphereGeometry_stagedOrder, spheregeometry)
	delete(stage.SphereGeometrys_mapString, spheregeometry.Name)

	return spheregeometry
}

// UnstageVoid removes spheregeometry off the model stage
func (spheregeometry *SphereGeometry) UnstageVoid(stage *Stage) {
	delete(stage.SphereGeometrys, spheregeometry)
	// issue1150
	// delete(stage.SphereGeometry_stagedOrder, spheregeometry)
	delete(stage.SphereGeometrys_mapString, spheregeometry.Name)
}

// commit spheregeometry to the back repo (if it is already staged)
func (spheregeometry *SphereGeometry) Commit(stage *Stage) *SphereGeometry {
	if _, ok := stage.SphereGeometrys[spheregeometry]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSphereGeometry(spheregeometry)
		}
	}
	return spheregeometry
}

func (spheregeometry *SphereGeometry) CommitVoid(stage *Stage) {
	spheregeometry.Commit(stage)
}

func (spheregeometry *SphereGeometry) StageVoid(stage *Stage) {
	spheregeometry.Stage(stage)
}

// Checkout spheregeometry to the back repo (if it is already staged)
func (spheregeometry *SphereGeometry) Checkout(stage *Stage) *SphereGeometry {
	if _, ok := stage.SphereGeometrys[spheregeometry]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSphereGeometry(spheregeometry)
		}
	}
	return spheregeometry
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
	if _, ok := stage.TorusGeometrys[torusgeometry]; !ok {
		stage.TorusGeometrys[torusgeometry] = struct{}{}
		stage.TorusGeometry_stagedOrder[torusgeometry] = stage.TorusGeometryOrder
		stage.TorusGeometry_orderStaged[stage.TorusGeometryOrder] = torusgeometry
		stage.TorusGeometryOrder++
	}
	stage.TorusGeometrys_mapString[torusgeometry.Name] = torusgeometry

	return torusgeometry
}

// StagePreserveOrder puts torusgeometry to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TorusGeometryOrder
// - update stage.TorusGeometryOrder accordingly
func (torusgeometry *TorusGeometry) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TorusGeometrys[torusgeometry]; !ok {
		stage.TorusGeometrys[torusgeometry] = struct{}{}

		if order > stage.TorusGeometryOrder {
			stage.TorusGeometryOrder = order
		}
		stage.TorusGeometry_stagedOrder[torusgeometry] = order
		stage.TorusGeometry_orderStaged[order] = torusgeometry
		stage.TorusGeometryOrder++
	}
	stage.TorusGeometrys_mapString[torusgeometry.Name] = torusgeometry
}

// Unstage removes torusgeometry off the model stage
func (torusgeometry *TorusGeometry) Unstage(stage *Stage) *TorusGeometry {
	delete(stage.TorusGeometrys, torusgeometry)
	// issue1150
	// delete(stage.TorusGeometry_stagedOrder, torusgeometry)
	delete(stage.TorusGeometrys_mapString, torusgeometry.Name)

	return torusgeometry
}

// UnstageVoid removes torusgeometry off the model stage
func (torusgeometry *TorusGeometry) UnstageVoid(stage *Stage) {
	delete(stage.TorusGeometrys, torusgeometry)
	// issue1150
	// delete(stage.TorusGeometry_stagedOrder, torusgeometry)
	delete(stage.TorusGeometrys_mapString, torusgeometry.Name)
}

// commit torusgeometry to the back repo (if it is already staged)
func (torusgeometry *TorusGeometry) Commit(stage *Stage) *TorusGeometry {
	if _, ok := stage.TorusGeometrys[torusgeometry]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTorusGeometry(torusgeometry)
		}
	}
	return torusgeometry
}

func (torusgeometry *TorusGeometry) CommitVoid(stage *Stage) {
	torusgeometry.Commit(stage)
}

func (torusgeometry *TorusGeometry) StageVoid(stage *Stage) {
	torusgeometry.Stage(stage)
}

// Checkout torusgeometry to the back repo (if it is already staged)
func (torusgeometry *TorusGeometry) Checkout(stage *Stage) *TorusGeometry {
	if _, ok := stage.TorusGeometrys[torusgeometry]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTorusGeometry(torusgeometry)
		}
	}
	return torusgeometry
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
	if _, ok := stage.Triangles[triangle]; !ok {
		stage.Triangles[triangle] = struct{}{}
		stage.Triangle_stagedOrder[triangle] = stage.TriangleOrder
		stage.Triangle_orderStaged[stage.TriangleOrder] = triangle
		stage.TriangleOrder++
	}
	stage.Triangles_mapString[triangle.Name] = triangle

	return triangle
}

// StagePreserveOrder puts triangle to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TriangleOrder
// - update stage.TriangleOrder accordingly
func (triangle *Triangle) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Triangles[triangle]; !ok {
		stage.Triangles[triangle] = struct{}{}

		if order > stage.TriangleOrder {
			stage.TriangleOrder = order
		}
		stage.Triangle_stagedOrder[triangle] = order
		stage.Triangle_orderStaged[order] = triangle
		stage.TriangleOrder++
	}
	stage.Triangles_mapString[triangle.Name] = triangle
}

// Unstage removes triangle off the model stage
func (triangle *Triangle) Unstage(stage *Stage) *Triangle {
	delete(stage.Triangles, triangle)
	// issue1150
	// delete(stage.Triangle_stagedOrder, triangle)
	delete(stage.Triangles_mapString, triangle.Name)

	return triangle
}

// UnstageVoid removes triangle off the model stage
func (triangle *Triangle) UnstageVoid(stage *Stage) {
	delete(stage.Triangles, triangle)
	// issue1150
	// delete(stage.Triangle_stagedOrder, triangle)
	delete(stage.Triangles_mapString, triangle.Name)
}

// commit triangle to the back repo (if it is already staged)
func (triangle *Triangle) Commit(stage *Stage) *Triangle {
	if _, ok := stage.Triangles[triangle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTriangle(triangle)
		}
	}
	return triangle
}

func (triangle *Triangle) CommitVoid(stage *Stage) {
	triangle.Commit(stage)
}

func (triangle *Triangle) StageVoid(stage *Stage) {
	triangle.Stage(stage)
}

// Checkout triangle to the back repo (if it is already staged)
func (triangle *Triangle) Checkout(stage *Stage) *Triangle {
	if _, ok := stage.Triangles[triangle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTriangle(triangle)
		}
	}
	return triangle
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
	if _, ok := stage.TubeGeometrys[tubegeometry]; !ok {
		stage.TubeGeometrys[tubegeometry] = struct{}{}
		stage.TubeGeometry_stagedOrder[tubegeometry] = stage.TubeGeometryOrder
		stage.TubeGeometry_orderStaged[stage.TubeGeometryOrder] = tubegeometry
		stage.TubeGeometryOrder++
	}
	stage.TubeGeometrys_mapString[tubegeometry.Name] = tubegeometry

	return tubegeometry
}

// StagePreserveOrder puts tubegeometry to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TubeGeometryOrder
// - update stage.TubeGeometryOrder accordingly
func (tubegeometry *TubeGeometry) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TubeGeometrys[tubegeometry]; !ok {
		stage.TubeGeometrys[tubegeometry] = struct{}{}

		if order > stage.TubeGeometryOrder {
			stage.TubeGeometryOrder = order
		}
		stage.TubeGeometry_stagedOrder[tubegeometry] = order
		stage.TubeGeometry_orderStaged[order] = tubegeometry
		stage.TubeGeometryOrder++
	}
	stage.TubeGeometrys_mapString[tubegeometry.Name] = tubegeometry
}

// Unstage removes tubegeometry off the model stage
func (tubegeometry *TubeGeometry) Unstage(stage *Stage) *TubeGeometry {
	delete(stage.TubeGeometrys, tubegeometry)
	// issue1150
	// delete(stage.TubeGeometry_stagedOrder, tubegeometry)
	delete(stage.TubeGeometrys_mapString, tubegeometry.Name)

	return tubegeometry
}

// UnstageVoid removes tubegeometry off the model stage
func (tubegeometry *TubeGeometry) UnstageVoid(stage *Stage) {
	delete(stage.TubeGeometrys, tubegeometry)
	// issue1150
	// delete(stage.TubeGeometry_stagedOrder, tubegeometry)
	delete(stage.TubeGeometrys_mapString, tubegeometry.Name)
}

// commit tubegeometry to the back repo (if it is already staged)
func (tubegeometry *TubeGeometry) Commit(stage *Stage) *TubeGeometry {
	if _, ok := stage.TubeGeometrys[tubegeometry]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTubeGeometry(tubegeometry)
		}
	}
	return tubegeometry
}

func (tubegeometry *TubeGeometry) CommitVoid(stage *Stage) {
	tubegeometry.Commit(stage)
}

func (tubegeometry *TubeGeometry) StageVoid(stage *Stage) {
	tubegeometry.Stage(stage)
}

// Checkout tubegeometry to the back repo (if it is already staged)
func (tubegeometry *TubeGeometry) Checkout(stage *Stage) *TubeGeometry {
	if _, ok := stage.TubeGeometrys[tubegeometry]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTubeGeometry(tubegeometry)
		}
	}
	return tubegeometry
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
	if _, ok := stage.Vector2s[vector2]; !ok {
		stage.Vector2s[vector2] = struct{}{}
		stage.Vector2_stagedOrder[vector2] = stage.Vector2Order
		stage.Vector2_orderStaged[stage.Vector2Order] = vector2
		stage.Vector2Order++
	}
	stage.Vector2s_mapString[vector2.Name] = vector2

	return vector2
}

// StagePreserveOrder puts vector2 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.Vector2Order
// - update stage.Vector2Order accordingly
func (vector2 *Vector2) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Vector2s[vector2]; !ok {
		stage.Vector2s[vector2] = struct{}{}

		if order > stage.Vector2Order {
			stage.Vector2Order = order
		}
		stage.Vector2_stagedOrder[vector2] = order
		stage.Vector2_orderStaged[order] = vector2
		stage.Vector2Order++
	}
	stage.Vector2s_mapString[vector2.Name] = vector2
}

// Unstage removes vector2 off the model stage
func (vector2 *Vector2) Unstage(stage *Stage) *Vector2 {
	delete(stage.Vector2s, vector2)
	// issue1150
	// delete(stage.Vector2_stagedOrder, vector2)
	delete(stage.Vector2s_mapString, vector2.Name)

	return vector2
}

// UnstageVoid removes vector2 off the model stage
func (vector2 *Vector2) UnstageVoid(stage *Stage) {
	delete(stage.Vector2s, vector2)
	// issue1150
	// delete(stage.Vector2_stagedOrder, vector2)
	delete(stage.Vector2s_mapString, vector2.Name)
}

// commit vector2 to the back repo (if it is already staged)
func (vector2 *Vector2) Commit(stage *Stage) *Vector2 {
	if _, ok := stage.Vector2s[vector2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitVector2(vector2)
		}
	}
	return vector2
}

func (vector2 *Vector2) CommitVoid(stage *Stage) {
	vector2.Commit(stage)
}

func (vector2 *Vector2) StageVoid(stage *Stage) {
	vector2.Stage(stage)
}

// Checkout vector2 to the back repo (if it is already staged)
func (vector2 *Vector2) Checkout(stage *Stage) *Vector2 {
	if _, ok := stage.Vector2s[vector2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutVector2(vector2)
		}
	}
	return vector2
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
	if _, ok := stage.Vector3s[vector3]; !ok {
		stage.Vector3s[vector3] = struct{}{}
		stage.Vector3_stagedOrder[vector3] = stage.Vector3Order
		stage.Vector3_orderStaged[stage.Vector3Order] = vector3
		stage.Vector3Order++
	}
	stage.Vector3s_mapString[vector3.Name] = vector3

	return vector3
}

// StagePreserveOrder puts vector3 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.Vector3Order
// - update stage.Vector3Order accordingly
func (vector3 *Vector3) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Vector3s[vector3]; !ok {
		stage.Vector3s[vector3] = struct{}{}

		if order > stage.Vector3Order {
			stage.Vector3Order = order
		}
		stage.Vector3_stagedOrder[vector3] = order
		stage.Vector3_orderStaged[order] = vector3
		stage.Vector3Order++
	}
	stage.Vector3s_mapString[vector3.Name] = vector3
}

// Unstage removes vector3 off the model stage
func (vector3 *Vector3) Unstage(stage *Stage) *Vector3 {
	delete(stage.Vector3s, vector3)
	// issue1150
	// delete(stage.Vector3_stagedOrder, vector3)
	delete(stage.Vector3s_mapString, vector3.Name)

	return vector3
}

// UnstageVoid removes vector3 off the model stage
func (vector3 *Vector3) UnstageVoid(stage *Stage) {
	delete(stage.Vector3s, vector3)
	// issue1150
	// delete(stage.Vector3_stagedOrder, vector3)
	delete(stage.Vector3s_mapString, vector3.Name)
}

// commit vector3 to the back repo (if it is already staged)
func (vector3 *Vector3) Commit(stage *Stage) *Vector3 {
	if _, ok := stage.Vector3s[vector3]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitVector3(vector3)
		}
	}
	return vector3
}

func (vector3 *Vector3) CommitVoid(stage *Stage) {
	vector3.Commit(stage)
}

func (vector3 *Vector3) StageVoid(stage *Stage) {
	vector3.Stage(stage)
}

// Checkout vector3 to the back repo (if it is already staged)
func (vector3 *Vector3) Checkout(stage *Stage) *Vector3 {
	if _, ok := stage.Vector3s[vector3]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutVector3(vector3)
		}
	}
	return vector3
}

// for satisfaction of GongStruct interface
func (vector3 *Vector3) GetName() (res string) {
	return vector3.Name
}

// for satisfaction of GongStruct interface
func (vector3 *Vector3) SetName(name string) {
	vector3.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMAmbiantLight(AmbiantLight *AmbiantLight)
	CreateORMBoxGeometry(BoxGeometry *BoxGeometry)
	CreateORMBufferGeometry(BufferGeometry *BufferGeometry)
	CreateORMCamera(Camera *Camera)
	CreateORMCanvas(Canvas *Canvas)
	CreateORMCurve(Curve *Curve)
	CreateORMCylinderGeometry(CylinderGeometry *CylinderGeometry)
	CreateORMDirectionalLight(DirectionalLight *DirectionalLight)
	CreateORMExtrudeGeometry(ExtrudeGeometry *ExtrudeGeometry)
	CreateORMMesh(Mesh *Mesh)
	CreateORMMeshMaterialBasic(MeshMaterialBasic *MeshMaterialBasic)
	CreateORMMeshPhysicalMaterial(MeshPhysicalMaterial *MeshPhysicalMaterial)
	CreateORMPlaneGeometry(PlaneGeometry *PlaneGeometry)
	CreateORMShape(Shape *Shape)
	CreateORMSphereGeometry(SphereGeometry *SphereGeometry)
	CreateORMTorusGeometry(TorusGeometry *TorusGeometry)
	CreateORMTriangle(Triangle *Triangle)
	CreateORMTubeGeometry(TubeGeometry *TubeGeometry)
	CreateORMVector2(Vector2 *Vector2)
	CreateORMVector3(Vector3 *Vector3)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAmbiantLight(AmbiantLight *AmbiantLight)
	DeleteORMBoxGeometry(BoxGeometry *BoxGeometry)
	DeleteORMBufferGeometry(BufferGeometry *BufferGeometry)
	DeleteORMCamera(Camera *Camera)
	DeleteORMCanvas(Canvas *Canvas)
	DeleteORMCurve(Curve *Curve)
	DeleteORMCylinderGeometry(CylinderGeometry *CylinderGeometry)
	DeleteORMDirectionalLight(DirectionalLight *DirectionalLight)
	DeleteORMExtrudeGeometry(ExtrudeGeometry *ExtrudeGeometry)
	DeleteORMMesh(Mesh *Mesh)
	DeleteORMMeshMaterialBasic(MeshMaterialBasic *MeshMaterialBasic)
	DeleteORMMeshPhysicalMaterial(MeshPhysicalMaterial *MeshPhysicalMaterial)
	DeleteORMPlaneGeometry(PlaneGeometry *PlaneGeometry)
	DeleteORMShape(Shape *Shape)
	DeleteORMSphereGeometry(SphereGeometry *SphereGeometry)
	DeleteORMTorusGeometry(TorusGeometry *TorusGeometry)
	DeleteORMTriangle(Triangle *Triangle)
	DeleteORMTubeGeometry(TubeGeometry *TubeGeometry)
	DeleteORMVector2(Vector2 *Vector2)
	DeleteORMVector3(Vector3 *Vector3)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.AmbiantLights = make(map[*AmbiantLight]struct{})
	stage.AmbiantLights_mapString = make(map[string]*AmbiantLight)
	stage.AmbiantLight_stagedOrder = make(map[*AmbiantLight]uint)
	stage.AmbiantLightOrder = 0

	stage.BoxGeometrys = make(map[*BoxGeometry]struct{})
	stage.BoxGeometrys_mapString = make(map[string]*BoxGeometry)
	stage.BoxGeometry_stagedOrder = make(map[*BoxGeometry]uint)
	stage.BoxGeometryOrder = 0

	stage.BufferGeometrys = make(map[*BufferGeometry]struct{})
	stage.BufferGeometrys_mapString = make(map[string]*BufferGeometry)
	stage.BufferGeometry_stagedOrder = make(map[*BufferGeometry]uint)
	stage.BufferGeometryOrder = 0

	stage.Cameras = make(map[*Camera]struct{})
	stage.Cameras_mapString = make(map[string]*Camera)
	stage.Camera_stagedOrder = make(map[*Camera]uint)
	stage.CameraOrder = 0

	stage.Canvass = make(map[*Canvas]struct{})
	stage.Canvass_mapString = make(map[string]*Canvas)
	stage.Canvas_stagedOrder = make(map[*Canvas]uint)
	stage.CanvasOrder = 0

	stage.Curves = make(map[*Curve]struct{})
	stage.Curves_mapString = make(map[string]*Curve)
	stage.Curve_stagedOrder = make(map[*Curve]uint)
	stage.CurveOrder = 0

	stage.CylinderGeometrys = make(map[*CylinderGeometry]struct{})
	stage.CylinderGeometrys_mapString = make(map[string]*CylinderGeometry)
	stage.CylinderGeometry_stagedOrder = make(map[*CylinderGeometry]uint)
	stage.CylinderGeometryOrder = 0

	stage.DirectionalLights = make(map[*DirectionalLight]struct{})
	stage.DirectionalLights_mapString = make(map[string]*DirectionalLight)
	stage.DirectionalLight_stagedOrder = make(map[*DirectionalLight]uint)
	stage.DirectionalLightOrder = 0

	stage.ExtrudeGeometrys = make(map[*ExtrudeGeometry]struct{})
	stage.ExtrudeGeometrys_mapString = make(map[string]*ExtrudeGeometry)
	stage.ExtrudeGeometry_stagedOrder = make(map[*ExtrudeGeometry]uint)
	stage.ExtrudeGeometryOrder = 0

	stage.Meshs = make(map[*Mesh]struct{})
	stage.Meshs_mapString = make(map[string]*Mesh)
	stage.Mesh_stagedOrder = make(map[*Mesh]uint)
	stage.MeshOrder = 0

	stage.MeshMaterialBasics = make(map[*MeshMaterialBasic]struct{})
	stage.MeshMaterialBasics_mapString = make(map[string]*MeshMaterialBasic)
	stage.MeshMaterialBasic_stagedOrder = make(map[*MeshMaterialBasic]uint)
	stage.MeshMaterialBasicOrder = 0

	stage.MeshPhysicalMaterials = make(map[*MeshPhysicalMaterial]struct{})
	stage.MeshPhysicalMaterials_mapString = make(map[string]*MeshPhysicalMaterial)
	stage.MeshPhysicalMaterial_stagedOrder = make(map[*MeshPhysicalMaterial]uint)
	stage.MeshPhysicalMaterialOrder = 0

	stage.PlaneGeometrys = make(map[*PlaneGeometry]struct{})
	stage.PlaneGeometrys_mapString = make(map[string]*PlaneGeometry)
	stage.PlaneGeometry_stagedOrder = make(map[*PlaneGeometry]uint)
	stage.PlaneGeometryOrder = 0

	stage.Shapes = make(map[*Shape]struct{})
	stage.Shapes_mapString = make(map[string]*Shape)
	stage.Shape_stagedOrder = make(map[*Shape]uint)
	stage.ShapeOrder = 0

	stage.SphereGeometrys = make(map[*SphereGeometry]struct{})
	stage.SphereGeometrys_mapString = make(map[string]*SphereGeometry)
	stage.SphereGeometry_stagedOrder = make(map[*SphereGeometry]uint)
	stage.SphereGeometryOrder = 0

	stage.TorusGeometrys = make(map[*TorusGeometry]struct{})
	stage.TorusGeometrys_mapString = make(map[string]*TorusGeometry)
	stage.TorusGeometry_stagedOrder = make(map[*TorusGeometry]uint)
	stage.TorusGeometryOrder = 0

	stage.Triangles = make(map[*Triangle]struct{})
	stage.Triangles_mapString = make(map[string]*Triangle)
	stage.Triangle_stagedOrder = make(map[*Triangle]uint)
	stage.TriangleOrder = 0

	stage.TubeGeometrys = make(map[*TubeGeometry]struct{})
	stage.TubeGeometrys_mapString = make(map[string]*TubeGeometry)
	stage.TubeGeometry_stagedOrder = make(map[*TubeGeometry]uint)
	stage.TubeGeometryOrder = 0

	stage.Vector2s = make(map[*Vector2]struct{})
	stage.Vector2s_mapString = make(map[string]*Vector2)
	stage.Vector2_stagedOrder = make(map[*Vector2]uint)
	stage.Vector2Order = 0

	stage.Vector3s = make(map[*Vector3]struct{})
	stage.Vector3s_mapString = make(map[string]*Vector3)
	stage.Vector3_stagedOrder = make(map[*Vector3]uint)
	stage.Vector3Order = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.AmbiantLights = nil
	stage.AmbiantLights_mapString = nil

	stage.BoxGeometrys = nil
	stage.BoxGeometrys_mapString = nil

	stage.BufferGeometrys = nil
	stage.BufferGeometrys_mapString = nil

	stage.Cameras = nil
	stage.Cameras_mapString = nil

	stage.Canvass = nil
	stage.Canvass_mapString = nil

	stage.Curves = nil
	stage.Curves_mapString = nil

	stage.CylinderGeometrys = nil
	stage.CylinderGeometrys_mapString = nil

	stage.DirectionalLights = nil
	stage.DirectionalLights_mapString = nil

	stage.ExtrudeGeometrys = nil
	stage.ExtrudeGeometrys_mapString = nil

	stage.Meshs = nil
	stage.Meshs_mapString = nil

	stage.MeshMaterialBasics = nil
	stage.MeshMaterialBasics_mapString = nil

	stage.MeshPhysicalMaterials = nil
	stage.MeshPhysicalMaterials_mapString = nil

	stage.PlaneGeometrys = nil
	stage.PlaneGeometrys_mapString = nil

	stage.Shapes = nil
	stage.Shapes_mapString = nil

	stage.SphereGeometrys = nil
	stage.SphereGeometrys_mapString = nil

	stage.TorusGeometrys = nil
	stage.TorusGeometrys_mapString = nil

	stage.Triangles = nil
	stage.Triangles_mapString = nil

	stage.TubeGeometrys = nil
	stage.TubeGeometrys_mapString = nil

	stage.Vector2s = nil
	stage.Vector2s_mapString = nil

	stage.Vector3s = nil
	stage.Vector3s_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for ambiantlight := range stage.AmbiantLights {
		ambiantlight.Unstage(stage)
	}

	for boxgeometry := range stage.BoxGeometrys {
		boxgeometry.Unstage(stage)
	}

	for buffergeometry := range stage.BufferGeometrys {
		buffergeometry.Unstage(stage)
	}

	for camera := range stage.Cameras {
		camera.Unstage(stage)
	}

	for canvas := range stage.Canvass {
		canvas.Unstage(stage)
	}

	for curve := range stage.Curves {
		curve.Unstage(stage)
	}

	for cylindergeometry := range stage.CylinderGeometrys {
		cylindergeometry.Unstage(stage)
	}

	for directionallight := range stage.DirectionalLights {
		directionallight.Unstage(stage)
	}

	for extrudegeometry := range stage.ExtrudeGeometrys {
		extrudegeometry.Unstage(stage)
	}

	for mesh := range stage.Meshs {
		mesh.Unstage(stage)
	}

	for meshmaterialbasic := range stage.MeshMaterialBasics {
		meshmaterialbasic.Unstage(stage)
	}

	for meshphysicalmaterial := range stage.MeshPhysicalMaterials {
		meshphysicalmaterial.Unstage(stage)
	}

	for planegeometry := range stage.PlaneGeometrys {
		planegeometry.Unstage(stage)
	}

	for shape := range stage.Shapes {
		shape.Unstage(stage)
	}

	for spheregeometry := range stage.SphereGeometrys {
		spheregeometry.Unstage(stage)
	}

	for torusgeometry := range stage.TorusGeometrys {
		torusgeometry.Unstage(stage)
	}

	for triangle := range stage.Triangles {
		triangle.Unstage(stage)
	}

	for tubegeometry := range stage.TubeGeometrys {
		tubegeometry.Unstage(stage)
	}

	for vector2 := range stage.Vector2s {
		vector2.Unstage(stage)
	}

	for vector3 := range stage.Vector3s {
		vector3.Unstage(stage)
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
	case map[*AmbiantLight]any:
		return any(&stage.AmbiantLights).(*Type)
	case map[*BoxGeometry]any:
		return any(&stage.BoxGeometrys).(*Type)
	case map[*BufferGeometry]any:
		return any(&stage.BufferGeometrys).(*Type)
	case map[*Camera]any:
		return any(&stage.Cameras).(*Type)
	case map[*Canvas]any:
		return any(&stage.Canvass).(*Type)
	case map[*Curve]any:
		return any(&stage.Curves).(*Type)
	case map[*CylinderGeometry]any:
		return any(&stage.CylinderGeometrys).(*Type)
	case map[*DirectionalLight]any:
		return any(&stage.DirectionalLights).(*Type)
	case map[*ExtrudeGeometry]any:
		return any(&stage.ExtrudeGeometrys).(*Type)
	case map[*Mesh]any:
		return any(&stage.Meshs).(*Type)
	case map[*MeshMaterialBasic]any:
		return any(&stage.MeshMaterialBasics).(*Type)
	case map[*MeshPhysicalMaterial]any:
		return any(&stage.MeshPhysicalMaterials).(*Type)
	case map[*PlaneGeometry]any:
		return any(&stage.PlaneGeometrys).(*Type)
	case map[*Shape]any:
		return any(&stage.Shapes).(*Type)
	case map[*SphereGeometry]any:
		return any(&stage.SphereGeometrys).(*Type)
	case map[*TorusGeometry]any:
		return any(&stage.TorusGeometrys).(*Type)
	case map[*Triangle]any:
		return any(&stage.Triangles).(*Type)
	case map[*TubeGeometry]any:
		return any(&stage.TubeGeometrys).(*Type)
	case map[*Vector2]any:
		return any(&stage.Vector2s).(*Type)
	case map[*Vector3]any:
		return any(&stage.Vector3s).(*Type)
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

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case AmbiantLight:
		return any(&stage.AmbiantLights).(*map[*Type]struct{})
	case BoxGeometry:
		return any(&stage.BoxGeometrys).(*map[*Type]struct{})
	case BufferGeometry:
		return any(&stage.BufferGeometrys).(*map[*Type]struct{})
	case Camera:
		return any(&stage.Cameras).(*map[*Type]struct{})
	case Canvas:
		return any(&stage.Canvass).(*map[*Type]struct{})
	case Curve:
		return any(&stage.Curves).(*map[*Type]struct{})
	case CylinderGeometry:
		return any(&stage.CylinderGeometrys).(*map[*Type]struct{})
	case DirectionalLight:
		return any(&stage.DirectionalLights).(*map[*Type]struct{})
	case ExtrudeGeometry:
		return any(&stage.ExtrudeGeometrys).(*map[*Type]struct{})
	case Mesh:
		return any(&stage.Meshs).(*map[*Type]struct{})
	case MeshMaterialBasic:
		return any(&stage.MeshMaterialBasics).(*map[*Type]struct{})
	case MeshPhysicalMaterial:
		return any(&stage.MeshPhysicalMaterials).(*map[*Type]struct{})
	case PlaneGeometry:
		return any(&stage.PlaneGeometrys).(*map[*Type]struct{})
	case Shape:
		return any(&stage.Shapes).(*map[*Type]struct{})
	case SphereGeometry:
		return any(&stage.SphereGeometrys).(*map[*Type]struct{})
	case TorusGeometry:
		return any(&stage.TorusGeometrys).(*map[*Type]struct{})
	case Triangle:
		return any(&stage.Triangles).(*map[*Type]struct{})
	case TubeGeometry:
		return any(&stage.TubeGeometrys).(*map[*Type]struct{})
	case Vector2:
		return any(&stage.Vector2s).(*map[*Type]struct{})
	case Vector3:
		return any(&stage.Vector3s).(*map[*Type]struct{})
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

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *Stage) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case AmbiantLight:
		return any(&stage.AmbiantLights_mapString).(*map[string]*Type)
	case BoxGeometry:
		return any(&stage.BoxGeometrys_mapString).(*map[string]*Type)
	case BufferGeometry:
		return any(&stage.BufferGeometrys_mapString).(*map[string]*Type)
	case Camera:
		return any(&stage.Cameras_mapString).(*map[string]*Type)
	case Canvas:
		return any(&stage.Canvass_mapString).(*map[string]*Type)
	case Curve:
		return any(&stage.Curves_mapString).(*map[string]*Type)
	case CylinderGeometry:
		return any(&stage.CylinderGeometrys_mapString).(*map[string]*Type)
	case DirectionalLight:
		return any(&stage.DirectionalLights_mapString).(*map[string]*Type)
	case ExtrudeGeometry:
		return any(&stage.ExtrudeGeometrys_mapString).(*map[string]*Type)
	case Mesh:
		return any(&stage.Meshs_mapString).(*map[string]*Type)
	case MeshMaterialBasic:
		return any(&stage.MeshMaterialBasics_mapString).(*map[string]*Type)
	case MeshPhysicalMaterial:
		return any(&stage.MeshPhysicalMaterials_mapString).(*map[string]*Type)
	case PlaneGeometry:
		return any(&stage.PlaneGeometrys_mapString).(*map[string]*Type)
	case Shape:
		return any(&stage.Shapes_mapString).(*map[string]*Type)
	case SphereGeometry:
		return any(&stage.SphereGeometrys_mapString).(*map[string]*Type)
	case TorusGeometry:
		return any(&stage.TorusGeometrys_mapString).(*map[string]*Type)
	case Triangle:
		return any(&stage.Triangles_mapString).(*map[string]*Type)
	case TubeGeometry:
		return any(&stage.TubeGeometrys_mapString).(*map[string]*Type)
	case Vector2:
		return any(&stage.Vector2s_mapString).(*map[string]*Type)
	case Vector3:
		return any(&stage.Vector3s_mapString).(*map[string]*Type)
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
	case AmbiantLight:
		return any(&AmbiantLight{
			// Initialisation of associations
		}).(*Type)
	case BoxGeometry:
		return any(&BoxGeometry{
			// Initialisation of associations
		}).(*Type)
	case BufferGeometry:
		return any(&BufferGeometry{
			// Initialisation of associations
			// field is initialized with an instance of Vector3 with the name of the field
			Vertices: []*Vector3{{Name: "Vertices"}},
			// field is initialized with an instance of Triangle with the name of the field
			Faces: []*Triangle{{Name: "Faces"}},
		}).(*Type)
	case Camera:
		return any(&Camera{
			// Initialisation of associations
		}).(*Type)
	case Canvas:
		return any(&Canvas{
			// Initialisation of associations
			// field is initialized with an instance of DirectionalLight with the name of the field
			DirectionalLights: []*DirectionalLight{{Name: "DirectionalLights"}},
			// field is initialized with an instance of AmbiantLight with the name of the field
			AmbiantLight: &AmbiantLight{Name: "AmbiantLight"},
			// field is initialized with an instance of Mesh with the name of the field
			Meshs: []*Mesh{{Name: "Meshs"}},
			// field is initialized with an instance of Camera with the name of the field
			Camera: &Camera{Name: "Camera"},
		}).(*Type)
	case Curve:
		return any(&Curve{
			// Initialisation of associations
			// field is initialized with an instance of Vector3 with the name of the field
			Points: []*Vector3{{Name: "Points"}},
		}).(*Type)
	case CylinderGeometry:
		return any(&CylinderGeometry{
			// Initialisation of associations
		}).(*Type)
	case DirectionalLight:
		return any(&DirectionalLight{
			// Initialisation of associations
		}).(*Type)
	case ExtrudeGeometry:
		return any(&ExtrudeGeometry{
			// Initialisation of associations
			// field is initialized with an instance of Shape with the name of the field
			Shape: &Shape{Name: "Shape"},
			// field is initialized with an instance of Curve with the name of the field
			ExtrudePath: &Curve{Name: "ExtrudePath"},
		}).(*Type)
	case Mesh:
		return any(&Mesh{
			// Initialisation of associations
			// field is initialized with an instance of MeshMaterialBasic with the name of the field
			MeshMaterialBasic: &MeshMaterialBasic{Name: "MeshMaterialBasic"},
			// field is initialized with an instance of MeshPhysicalMaterial with the name of the field
			MeshPhysicalMaterial: &MeshPhysicalMaterial{Name: "MeshPhysicalMaterial"},
			// field is initialized with an instance of CylinderGeometry with the name of the field
			CylinderGeometry: &CylinderGeometry{Name: "CylinderGeometry"},
			// field is initialized with an instance of BoxGeometry with the name of the field
			BoxGeometry: &BoxGeometry{Name: "BoxGeometry"},
			// field is initialized with an instance of SphereGeometry with the name of the field
			SphereGeometry: &SphereGeometry{Name: "SphereGeometry"},
			// field is initialized with an instance of TorusGeometry with the name of the field
			TorusGeometry: &TorusGeometry{Name: "TorusGeometry"},
			// field is initialized with an instance of PlaneGeometry with the name of the field
			PlaneGeometry: &PlaneGeometry{Name: "PlaneGeometry"},
			// field is initialized with an instance of TubeGeometry with the name of the field
			TubeGeometry: &TubeGeometry{Name: "TubeGeometry"},
			// field is initialized with an instance of ExtrudeGeometry with the name of the field
			ExtrudeGeometry: &ExtrudeGeometry{Name: "ExtrudeGeometry"},
			// field is initialized with an instance of BufferGeometry with the name of the field
			BufferGeometry: &BufferGeometry{Name: "BufferGeometry"},
		}).(*Type)
	case MeshMaterialBasic:
		return any(&MeshMaterialBasic{
			// Initialisation of associations
		}).(*Type)
	case MeshPhysicalMaterial:
		return any(&MeshPhysicalMaterial{
			// Initialisation of associations
		}).(*Type)
	case PlaneGeometry:
		return any(&PlaneGeometry{
			// Initialisation of associations
		}).(*Type)
	case Shape:
		return any(&Shape{
			// Initialisation of associations
			// field is initialized with an instance of Vector2 with the name of the field
			Points: []*Vector2{{Name: "Points"}},
		}).(*Type)
	case SphereGeometry:
		return any(&SphereGeometry{
			// Initialisation of associations
		}).(*Type)
	case TorusGeometry:
		return any(&TorusGeometry{
			// Initialisation of associations
		}).(*Type)
	case Triangle:
		return any(&Triangle{
			// Initialisation of associations
		}).(*Type)
	case TubeGeometry:
		return any(&TubeGeometry{
			// Initialisation of associations
			// field is initialized with an instance of Curve with the name of the field
			Path: &Curve{Name: "Path"},
		}).(*Type)
	case Vector2:
		return any(&Vector2{
			// Initialisation of associations
		}).(*Type)
	case Vector3:
		return any(&Vector3{
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

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type GongstructIF]() (res string) {
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

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type GongstructIF]() (res []ReverseField) {
	res = make([]ReverseField, 0)

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

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (ambiantlight *AmbiantLight) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		ambiantlight.Name = value.GetValueString()
	case "Intensity":
		ambiantlight.Intensity = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (boxgeometry *BoxGeometry) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		boxgeometry.Name = value.GetValueString()
	case "Width":
		boxgeometry.Width = value.GetValueFloat()
	case "Height":
		boxgeometry.Height = value.GetValueFloat()
	case "Depth":
		boxgeometry.Depth = value.GetValueFloat()
	case "WidthSegments":
		boxgeometry.WidthSegments = int(value.GetValueInt())
	case "HeightSegments":
		boxgeometry.HeightSegments = int(value.GetValueInt())
	case "DepthSegments":
		boxgeometry.DepthSegments = int(value.GetValueInt())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (buffergeometry *BufferGeometry) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		buffergeometry.Name = value.GetValueString()
	case "Vertices":
		buffergeometry.Vertices = make([]*Vector3, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Vector3s {
					if stage.Vector3_stagedOrder[__instance__] == uint(id) {
						buffergeometry.Vertices = append(buffergeometry.Vertices, __instance__)
						break
					}
				}
			}
		}
	case "Faces":
		buffergeometry.Faces = make([]*Triangle, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Triangles {
					if stage.Triangle_stagedOrder[__instance__] == uint(id) {
						buffergeometry.Faces = append(buffergeometry.Faces, __instance__)
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

func (camera *Camera) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		camera.Name = value.GetValueString()
	case "X":
		camera.X = value.GetValueFloat()
	case "Y":
		camera.Y = value.GetValueFloat()
	case "Z":
		camera.Z = value.GetValueFloat()
	case "TargetX":
		camera.TargetX = value.GetValueFloat()
	case "TargetY":
		camera.TargetY = value.GetValueFloat()
	case "TargetZ":
		camera.TargetZ = value.GetValueFloat()
	case "Fov":
		camera.Fov = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (canvas *Canvas) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		canvas.Name = value.GetValueString()
	case "DirectionalLights":
		canvas.DirectionalLights = make([]*DirectionalLight, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DirectionalLights {
					if stage.DirectionalLight_stagedOrder[__instance__] == uint(id) {
						canvas.DirectionalLights = append(canvas.DirectionalLights, __instance__)
						break
					}
				}
			}
		}
	case "AmbiantLight":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			canvas.AmbiantLight = nil
			for __instance__ := range stage.AmbiantLights {
				if stage.AmbiantLight_stagedOrder[__instance__] == uint(id) {
					canvas.AmbiantLight = __instance__
					break
				}
			}
		}
	case "Meshs":
		canvas.Meshs = make([]*Mesh, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Meshs {
					if stage.Mesh_stagedOrder[__instance__] == uint(id) {
						canvas.Meshs = append(canvas.Meshs, __instance__)
						break
					}
				}
			}
		}
	case "Camera":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			canvas.Camera = nil
			for __instance__ := range stage.Cameras {
				if stage.Camera_stagedOrder[__instance__] == uint(id) {
					canvas.Camera = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (curve *Curve) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		curve.Name = value.GetValueString()
	case "Points":
		curve.Points = make([]*Vector3, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Vector3s {
					if stage.Vector3_stagedOrder[__instance__] == uint(id) {
						curve.Points = append(curve.Points, __instance__)
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

func (cylindergeometry *CylinderGeometry) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		cylindergeometry.Name = value.GetValueString()
	case "RadiusTop":
		cylindergeometry.RadiusTop = value.GetValueFloat()
	case "RadiusBottom":
		cylindergeometry.RadiusBottom = value.GetValueFloat()
	case "Height":
		cylindergeometry.Height = value.GetValueFloat()
	case "RadialSegments":
		cylindergeometry.RadialSegments = int(value.GetValueInt())
	case "HeightSegments":
		cylindergeometry.HeightSegments = int(value.GetValueInt())
	case "OpenEnded":
		cylindergeometry.OpenEnded = value.GetValueBool()
	case "ThetaStart":
		cylindergeometry.ThetaStart = value.GetValueFloat()
	case "ThetaLength":
		cylindergeometry.ThetaLength = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (directionallight *DirectionalLight) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		directionallight.Name = value.GetValueString()
	case "X":
		directionallight.X = value.GetValueFloat()
	case "Y":
		directionallight.Y = value.GetValueFloat()
	case "Z":
		directionallight.Z = value.GetValueFloat()
	case "Intensity":
		directionallight.Intensity = value.GetValueFloat()
	case "IsWithCastShadow":
		directionallight.IsWithCastShadow = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (extrudegeometry *ExtrudeGeometry) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		extrudegeometry.Name = value.GetValueString()
	case "Shape":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			extrudegeometry.Shape = nil
			for __instance__ := range stage.Shapes {
				if stage.Shape_stagedOrder[__instance__] == uint(id) {
					extrudegeometry.Shape = __instance__
					break
				}
			}
		}
	case "ExtrudePath":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			extrudegeometry.ExtrudePath = nil
			for __instance__ := range stage.Curves {
				if stage.Curve_stagedOrder[__instance__] == uint(id) {
					extrudegeometry.ExtrudePath = __instance__
					break
				}
			}
		}
	case "Steps":
		extrudegeometry.Steps = int(value.GetValueInt())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (mesh *Mesh) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		mesh.Name = value.GetValueString()
	case "X":
		mesh.X = value.GetValueFloat()
	case "Y":
		mesh.Y = value.GetValueFloat()
	case "Z":
		mesh.Z = value.GetValueFloat()
	case "MeshMaterialBasic":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			mesh.MeshMaterialBasic = nil
			for __instance__ := range stage.MeshMaterialBasics {
				if stage.MeshMaterialBasic_stagedOrder[__instance__] == uint(id) {
					mesh.MeshMaterialBasic = __instance__
					break
				}
			}
		}
	case "MeshPhysicalMaterial":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			mesh.MeshPhysicalMaterial = nil
			for __instance__ := range stage.MeshPhysicalMaterials {
				if stage.MeshPhysicalMaterial_stagedOrder[__instance__] == uint(id) {
					mesh.MeshPhysicalMaterial = __instance__
					break
				}
			}
		}
	case "CylinderGeometry":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			mesh.CylinderGeometry = nil
			for __instance__ := range stage.CylinderGeometrys {
				if stage.CylinderGeometry_stagedOrder[__instance__] == uint(id) {
					mesh.CylinderGeometry = __instance__
					break
				}
			}
		}
	case "BoxGeometry":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			mesh.BoxGeometry = nil
			for __instance__ := range stage.BoxGeometrys {
				if stage.BoxGeometry_stagedOrder[__instance__] == uint(id) {
					mesh.BoxGeometry = __instance__
					break
				}
			}
		}
	case "SphereGeometry":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			mesh.SphereGeometry = nil
			for __instance__ := range stage.SphereGeometrys {
				if stage.SphereGeometry_stagedOrder[__instance__] == uint(id) {
					mesh.SphereGeometry = __instance__
					break
				}
			}
		}
	case "TorusGeometry":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			mesh.TorusGeometry = nil
			for __instance__ := range stage.TorusGeometrys {
				if stage.TorusGeometry_stagedOrder[__instance__] == uint(id) {
					mesh.TorusGeometry = __instance__
					break
				}
			}
		}
	case "PlaneGeometry":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			mesh.PlaneGeometry = nil
			for __instance__ := range stage.PlaneGeometrys {
				if stage.PlaneGeometry_stagedOrder[__instance__] == uint(id) {
					mesh.PlaneGeometry = __instance__
					break
				}
			}
		}
	case "TubeGeometry":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			mesh.TubeGeometry = nil
			for __instance__ := range stage.TubeGeometrys {
				if stage.TubeGeometry_stagedOrder[__instance__] == uint(id) {
					mesh.TubeGeometry = __instance__
					break
				}
			}
		}
	case "ExtrudeGeometry":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			mesh.ExtrudeGeometry = nil
			for __instance__ := range stage.ExtrudeGeometrys {
				if stage.ExtrudeGeometry_stagedOrder[__instance__] == uint(id) {
					mesh.ExtrudeGeometry = __instance__
					break
				}
			}
		}
	case "BufferGeometry":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			mesh.BufferGeometry = nil
			for __instance__ := range stage.BufferGeometrys {
				if stage.BufferGeometry_stagedOrder[__instance__] == uint(id) {
					mesh.BufferGeometry = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (meshmaterialbasic *MeshMaterialBasic) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		meshmaterialbasic.Name = value.GetValueString()
	case "Color":
		meshmaterialbasic.Color = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		meshphysicalmaterial.Name = value.GetValueString()
	case "Color":
		meshphysicalmaterial.Color = value.GetValueString()
	case "Wireframe":
		meshphysicalmaterial.Wireframe = value.GetValueBool()
	case "Opacity":
		meshphysicalmaterial.Opacity = value.GetValueFloat()
	case "Transparent":
		meshphysicalmaterial.Transparent = value.GetValueBool()
	case "Visible":
		meshphysicalmaterial.Visible = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (planegeometry *PlaneGeometry) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		planegeometry.Name = value.GetValueString()
	case "Width":
		planegeometry.Width = value.GetValueFloat()
	case "Height":
		planegeometry.Height = value.GetValueFloat()
	case "WidthSegments":
		planegeometry.WidthSegments = int(value.GetValueInt())
	case "HeightSegments":
		planegeometry.HeightSegments = int(value.GetValueInt())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (shape *Shape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		shape.Name = value.GetValueString()
	case "Points":
		shape.Points = make([]*Vector2, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Vector2s {
					if stage.Vector2_stagedOrder[__instance__] == uint(id) {
						shape.Points = append(shape.Points, __instance__)
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

func (spheregeometry *SphereGeometry) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		spheregeometry.Name = value.GetValueString()
	case "Radius":
		spheregeometry.Radius = value.GetValueFloat()
	case "WidthSegments":
		spheregeometry.WidthSegments = int(value.GetValueInt())
	case "HeightSegments":
		spheregeometry.HeightSegments = int(value.GetValueInt())
	case "PhiStart":
		spheregeometry.PhiStart = value.GetValueFloat()
	case "PhiLength":
		spheregeometry.PhiLength = value.GetValueFloat()
	case "ThetaStart":
		spheregeometry.ThetaStart = value.GetValueFloat()
	case "ThetaLength":
		spheregeometry.ThetaLength = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (torusgeometry *TorusGeometry) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		torusgeometry.Name = value.GetValueString()
	case "Radius":
		torusgeometry.Radius = value.GetValueFloat()
	case "Tube":
		torusgeometry.Tube = value.GetValueFloat()
	case "RadialSegments":
		torusgeometry.RadialSegments = int(value.GetValueInt())
	case "TubularSegments":
		torusgeometry.TubularSegments = int(value.GetValueInt())
	case "Arc":
		torusgeometry.Arc = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (triangle *Triangle) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		triangle.Name = value.GetValueString()
	case "V1":
		triangle.V1 = int(value.GetValueInt())
	case "V2":
		triangle.V2 = int(value.GetValueInt())
	case "V3":
		triangle.V3 = int(value.GetValueInt())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (tubegeometry *TubeGeometry) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		tubegeometry.Name = value.GetValueString()
	case "Path":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			tubegeometry.Path = nil
			for __instance__ := range stage.Curves {
				if stage.Curve_stagedOrder[__instance__] == uint(id) {
					tubegeometry.Path = __instance__
					break
				}
			}
		}
	case "TubularSegments":
		tubegeometry.TubularSegments = int(value.GetValueInt())
	case "Radius":
		tubegeometry.Radius = value.GetValueFloat()
	case "RadialSegments":
		tubegeometry.RadialSegments = int(value.GetValueInt())
	case "Closed":
		tubegeometry.Closed = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (vector2 *Vector2) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		vector2.Name = value.GetValueString()
	case "X":
		vector2.X = value.GetValueFloat()
	case "Y":
		vector2.Y = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (vector3 *Vector3) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		vector3.Name = value.GetValueString()
	case "X":
		vector3.X = value.GetValueFloat()
	case "Y":
		vector3.Y = value.GetValueFloat()
	case "Z":
		vector3.Z = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
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

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	stage.AmbiantLights_mapString = make(map[string]*AmbiantLight)
	for ambiantlight := range stage.AmbiantLights {
		stage.AmbiantLights_mapString[ambiantlight.Name] = ambiantlight
	}

	stage.BoxGeometrys_mapString = make(map[string]*BoxGeometry)
	for boxgeometry := range stage.BoxGeometrys {
		stage.BoxGeometrys_mapString[boxgeometry.Name] = boxgeometry
	}

	stage.BufferGeometrys_mapString = make(map[string]*BufferGeometry)
	for buffergeometry := range stage.BufferGeometrys {
		stage.BufferGeometrys_mapString[buffergeometry.Name] = buffergeometry
	}

	stage.Cameras_mapString = make(map[string]*Camera)
	for camera := range stage.Cameras {
		stage.Cameras_mapString[camera.Name] = camera
	}

	stage.Canvass_mapString = make(map[string]*Canvas)
	for canvas := range stage.Canvass {
		stage.Canvass_mapString[canvas.Name] = canvas
	}

	stage.Curves_mapString = make(map[string]*Curve)
	for curve := range stage.Curves {
		stage.Curves_mapString[curve.Name] = curve
	}

	stage.CylinderGeometrys_mapString = make(map[string]*CylinderGeometry)
	for cylindergeometry := range stage.CylinderGeometrys {
		stage.CylinderGeometrys_mapString[cylindergeometry.Name] = cylindergeometry
	}

	stage.DirectionalLights_mapString = make(map[string]*DirectionalLight)
	for directionallight := range stage.DirectionalLights {
		stage.DirectionalLights_mapString[directionallight.Name] = directionallight
	}

	stage.ExtrudeGeometrys_mapString = make(map[string]*ExtrudeGeometry)
	for extrudegeometry := range stage.ExtrudeGeometrys {
		stage.ExtrudeGeometrys_mapString[extrudegeometry.Name] = extrudegeometry
	}

	stage.Meshs_mapString = make(map[string]*Mesh)
	for mesh := range stage.Meshs {
		stage.Meshs_mapString[mesh.Name] = mesh
	}

	stage.MeshMaterialBasics_mapString = make(map[string]*MeshMaterialBasic)
	for meshmaterialbasic := range stage.MeshMaterialBasics {
		stage.MeshMaterialBasics_mapString[meshmaterialbasic.Name] = meshmaterialbasic
	}

	stage.MeshPhysicalMaterials_mapString = make(map[string]*MeshPhysicalMaterial)
	for meshphysicalmaterial := range stage.MeshPhysicalMaterials {
		stage.MeshPhysicalMaterials_mapString[meshphysicalmaterial.Name] = meshphysicalmaterial
	}

	stage.PlaneGeometrys_mapString = make(map[string]*PlaneGeometry)
	for planegeometry := range stage.PlaneGeometrys {
		stage.PlaneGeometrys_mapString[planegeometry.Name] = planegeometry
	}

	stage.Shapes_mapString = make(map[string]*Shape)
	for shape := range stage.Shapes {
		stage.Shapes_mapString[shape.Name] = shape
	}

	stage.SphereGeometrys_mapString = make(map[string]*SphereGeometry)
	for spheregeometry := range stage.SphereGeometrys {
		stage.SphereGeometrys_mapString[spheregeometry.Name] = spheregeometry
	}

	stage.TorusGeometrys_mapString = make(map[string]*TorusGeometry)
	for torusgeometry := range stage.TorusGeometrys {
		stage.TorusGeometrys_mapString[torusgeometry.Name] = torusgeometry
	}

	stage.Triangles_mapString = make(map[string]*Triangle)
	for triangle := range stage.Triangles {
		stage.Triangles_mapString[triangle.Name] = triangle
	}

	stage.TubeGeometrys_mapString = make(map[string]*TubeGeometry)
	for tubegeometry := range stage.TubeGeometrys {
		stage.TubeGeometrys_mapString[tubegeometry.Name] = tubegeometry
	}

	stage.Vector2s_mapString = make(map[string]*Vector2)
	for vector2 := range stage.Vector2s {
		stage.Vector2s_mapString[vector2.Name] = vector2
	}

	stage.Vector3s_mapString = make(map[string]*Vector3)
	for vector3 := range stage.Vector3s {
		stage.Vector3s_mapString[vector3.Name] = vector3
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
