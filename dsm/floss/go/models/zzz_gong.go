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

	floss_go "github.com/fullstack-lang/gong/dsm/floss/go"
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
	DiagramFlosss                map[*DiagramFloss]struct{}
	DiagramFlosss_instance       map[*DiagramFloss]*DiagramFloss
	DiagramFlosss_mapString      map[string]*DiagramFloss
	DiagramFlossOrder            uint
	DiagramFloss_stagedOrder     map[*DiagramFloss]uint
	DiagramFloss_orderStaged     map[uint]*DiagramFloss
	DiagramFlosss_reference      map[*DiagramFloss]*DiagramFloss
	DiagramFlosss_referenceOrder map[*DiagramFloss]uint

	// insertion point for slice of pointers maps
	DiagramFloss_System_Shapes_reverseMap map[*SystemShape]*DiagramFloss

	DiagramFloss_SystemsWhoseNodeIsExpanded_reverseMap map[*System]*DiagramFloss

	OnAfterDiagramFlossCreateCallback OnAfterCreateInterface[DiagramFloss]
	OnAfterDiagramFlossUpdateCallback OnAfterUpdateInterface[DiagramFloss]
	OnAfterDiagramFlossDeleteCallback OnAfterDeleteInterface[DiagramFloss]
	OnAfterDiagramFlossReadCallback   OnAfterReadInterface[DiagramFloss]

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

	Library_SubLibrariesWhoseNodeIsExpanded_reverseMap map[*Library]*Library

	Library_RootSystemes_reverseMap map[*System]*Library

	Library_SystemsWhoseNodeIsExpanded_reverseMap map[*System]*Library

	OnAfterLibraryCreateCallback OnAfterCreateInterface[Library]
	OnAfterLibraryUpdateCallback OnAfterUpdateInterface[Library]
	OnAfterLibraryDeleteCallback OnAfterDeleteInterface[Library]
	OnAfterLibraryReadCallback   OnAfterReadInterface[Library]

	Systems                map[*System]struct{}
	Systems_instance       map[*System]*System
	Systems_mapString      map[string]*System
	SystemOrder            uint
	System_stagedOrder     map[*System]uint
	System_orderStaged     map[uint]*System
	Systems_reference      map[*System]*System
	Systems_referenceOrder map[*System]uint

	// insertion point for slice of pointers maps
	System_DiagramFlosses_reverseMap map[*DiagramFloss]*System

	System_DiagramFlossWhoseNodeIsExpanded_reverseMap map[*DiagramFloss]*System

	System_SubSystemes_reverseMap map[*System]*System

	OnAfterSystemCreateCallback OnAfterCreateInterface[System]
	OnAfterSystemUpdateCallback OnAfterUpdateInterface[System]
	OnAfterSystemDeleteCallback OnAfterDeleteInterface[System]
	OnAfterSystemReadCallback   OnAfterReadInterface[System]

	SystemShapes                map[*SystemShape]struct{}
	SystemShapes_instance       map[*SystemShape]*SystemShape
	SystemShapes_mapString      map[string]*SystemShape
	SystemShapeOrder            uint
	SystemShape_stagedOrder     map[*SystemShape]uint
	SystemShape_orderStaged     map[uint]*SystemShape
	SystemShapes_reference      map[*SystemShape]*SystemShape
	SystemShapes_referenceOrder map[*SystemShape]uint

	// insertion point for slice of pointers maps
	OnAfterSystemShapeCreateCallback OnAfterCreateInterface[SystemShape]
	OnAfterSystemShapeUpdateCallback OnAfterUpdateInterface[SystemShape]
	OnAfterSystemShapeDeleteCallback OnAfterDeleteInterface[SystemShape]
	OnAfterSystemShapeReadCallback   OnAfterReadInterface[SystemShape]

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
	stage.DiagramFlosss_reference = make(map[*DiagramFloss]*DiagramFloss)
	stage.DiagramFlosss_instance = make(map[*DiagramFloss]*DiagramFloss)
	stage.DiagramFlosss_referenceOrder = make(map[*DiagramFloss]uint)

	stage.Librarys_reference = make(map[*Library]*Library)
	stage.Librarys_instance = make(map[*Library]*Library)
	stage.Librarys_referenceOrder = make(map[*Library]uint)

	stage.Systems_reference = make(map[*System]*System)
	stage.Systems_instance = make(map[*System]*System)
	stage.Systems_referenceOrder = make(map[*System]uint)

	stage.SystemShapes_reference = make(map[*SystemShape]*SystemShape)
	stage.SystemShapes_instance = make(map[*SystemShape]*SystemShape)
	stage.SystemShapes_referenceOrder = make(map[*SystemShape]uint)

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
	var maxDiagramFlossOrder uint
	var foundDiagramFloss bool
	for _, order := range stage.DiagramFloss_stagedOrder {
		if !foundDiagramFloss || order > maxDiagramFlossOrder {
			maxDiagramFlossOrder = order
			foundDiagramFloss = true
		}
	}
	if foundDiagramFloss {
		stage.DiagramFlossOrder = maxDiagramFlossOrder + 1
	} else {
		stage.DiagramFlossOrder = 0
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

	var maxSystemOrder uint
	var foundSystem bool
	for _, order := range stage.System_stagedOrder {
		if !foundSystem || order > maxSystemOrder {
			maxSystemOrder = order
			foundSystem = true
		}
	}
	if foundSystem {
		stage.SystemOrder = maxSystemOrder + 1
	} else {
		stage.SystemOrder = 0
	}

	var maxSystemShapeOrder uint
	var foundSystemShape bool
	for _, order := range stage.SystemShape_stagedOrder {
		if !foundSystemShape || order > maxSystemShapeOrder {
			maxSystemShapeOrder = order
			foundSystemShape = true
		}
	}
	if foundSystemShape {
		stage.SystemShapeOrder = maxSystemShapeOrder + 1
	} else {
		stage.SystemShapeOrder = 0
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
	case *DiagramFloss:
		tmp := GetStructInstancesByOrder(stage.DiagramFlosss, stage.DiagramFloss_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *DiagramFloss implements.
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
	case *System:
		tmp := GetStructInstancesByOrder(stage.Systems, stage.System_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *System implements.
			res = append(res, any(v).(T))
		}
		return res
	case *SystemShape:
		tmp := GetStructInstancesByOrder(stage.SystemShapes, stage.SystemShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SystemShape implements.
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
	case "DiagramFloss":
		res = GetNamedStructInstances(stage.DiagramFlosss, stage.DiagramFloss_stagedOrder)
	case "Library":
		res = GetNamedStructInstances(stage.Librarys, stage.Library_stagedOrder)
	case "System":
		res = GetNamedStructInstances(stage.Systems, stage.System_stagedOrder)
	case "SystemShape":
		res = GetNamedStructInstances(stage.SystemShapes, stage.SystemShape_stagedOrder)
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
	return "github.com/fullstack-lang/gong/dsm/floss/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return floss_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return floss_go.GoDiagramsDir
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
	CommitDiagramFloss(diagramfloss *DiagramFloss)
	CheckoutDiagramFloss(diagramfloss *DiagramFloss)
	CommitLibrary(library *Library)
	CheckoutLibrary(library *Library)
	CommitSystem(system *System)
	CheckoutSystem(system *System)
	CommitSystemShape(systemshape *SystemShape)
	CheckoutSystemShape(systemshape *SystemShape)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		DiagramFlosss:           make(map[*DiagramFloss]struct{}),
		DiagramFlosss_mapString: make(map[string]*DiagramFloss),

		Librarys:           make(map[*Library]struct{}),
		Librarys_mapString: make(map[string]*Library),

		Systems:           make(map[*System]struct{}),
		Systems_mapString: make(map[string]*System),

		SystemShapes:           make(map[*SystemShape]struct{}),
		SystemShapes_mapString: make(map[string]*SystemShape),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		DiagramFloss_stagedOrder: make(map[*DiagramFloss]uint),
		DiagramFloss_orderStaged: make(map[uint]*DiagramFloss),
		DiagramFlosss_reference:  make(map[*DiagramFloss]*DiagramFloss),

		Library_stagedOrder: make(map[*Library]uint),
		Library_orderStaged: make(map[uint]*Library),
		Librarys_reference:  make(map[*Library]*Library),

		System_stagedOrder: make(map[*System]uint),
		System_orderStaged: make(map[uint]*System),
		Systems_reference:  make(map[*System]*System),

		SystemShape_stagedOrder: make(map[*SystemShape]uint),
		SystemShape_orderStaged: make(map[uint]*SystemShape),
		SystemShapes_reference:  make(map[*SystemShape]*SystemShape),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"DiagramFloss": &DiagramFlossUnmarshaller{},

			"Library": &LibraryUnmarshaller{},

			"System": &SystemUnmarshaller{},

			"SystemShape": &SystemShapeUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "DiagramFloss"},
			{name: "Library"},
			{name: "System"},
			{name: "SystemShape"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *DiagramFloss:
		return stage.DiagramFloss_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *System:
		return stage.System_stagedOrder[instance]
	case *SystemShape:
		return stage.SystemShape_stagedOrder[instance]
	default:
		return 0 // should not happen
	}
}

func GongGetInstanceFromOrder[Type PointerToGongstruct](stage *Stage, order uint) (res Type) {
	var t Type
	switch any(t).(type) {
	// insertion point for order map initialisations
	case *DiagramFloss:
		return any(stage.DiagramFloss_orderStaged[order]).(Type)
	case *Library:
		return any(stage.Library_orderStaged[order]).(Type)
	case *System:
		return any(stage.System_orderStaged[order]).(Type)
	case *SystemShape:
		return any(stage.SystemShape_orderStaged[order]).(Type)
	default:
		return // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *DiagramFloss:
		return stage.DiagramFloss_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *System:
		return stage.System_stagedOrder[instance]
	case *SystemShape:
		return stage.SystemShape_stagedOrder[instance]
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
	stage.Map_GongStructName_InstancesNb["DiagramFloss"] = len(stage.DiagramFlosss)
	stage.Map_GongStructName_InstancesNb["Library"] = len(stage.Librarys)
	stage.Map_GongStructName_InstancesNb["System"] = len(stage.Systems)
	stage.Map_GongStructName_InstancesNb["SystemShape"] = len(stage.SystemShapes)
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
// Stage puts diagramfloss to the model stage
func (diagramfloss *DiagramFloss) Stage(stage *Stage) *DiagramFloss {
	if _, ok := stage.DiagramFlosss[diagramfloss]; !ok {
		stage.DiagramFlosss[diagramfloss] = struct{}{}
		stage.DiagramFloss_stagedOrder[diagramfloss] = stage.DiagramFlossOrder
		stage.DiagramFloss_orderStaged[stage.DiagramFlossOrder] = diagramfloss
		stage.DiagramFlossOrder++
	}
	stage.DiagramFlosss_mapString[diagramfloss.Name] = diagramfloss

	return diagramfloss
}

// StagePreserveOrder puts diagramfloss to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DiagramFlossOrder
// - update stage.DiagramFlossOrder accordingly
func (diagramfloss *DiagramFloss) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.DiagramFlosss[diagramfloss]; !ok {
		stage.DiagramFlosss[diagramfloss] = struct{}{}

		if order > stage.DiagramFlossOrder {
			stage.DiagramFlossOrder = order
		}
		stage.DiagramFloss_stagedOrder[diagramfloss] = order
		stage.DiagramFloss_orderStaged[order] = diagramfloss
		stage.DiagramFlossOrder++
	}
	stage.DiagramFlosss_mapString[diagramfloss.Name] = diagramfloss
}

// Unstage removes diagramfloss off the model stage
func (diagramfloss *DiagramFloss) Unstage(stage *Stage) *DiagramFloss {
	delete(stage.DiagramFlosss, diagramfloss)
	// issue1150
	// delete(stage.DiagramFloss_stagedOrder, diagramfloss)
	delete(stage.DiagramFlosss_mapString, diagramfloss.Name)

	return diagramfloss
}

// UnstageVoid removes diagramfloss off the model stage
func (diagramfloss *DiagramFloss) UnstageVoid(stage *Stage) {
	delete(stage.DiagramFlosss, diagramfloss)
	// issue1150
	// delete(stage.DiagramFloss_stagedOrder, diagramfloss)
	delete(stage.DiagramFlosss_mapString, diagramfloss.Name)
}

// commit diagramfloss to the back repo (if it is already staged)
func (diagramfloss *DiagramFloss) Commit(stage *Stage) *DiagramFloss {
	if _, ok := stage.DiagramFlosss[diagramfloss]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDiagramFloss(diagramfloss)
		}
	}
	return diagramfloss
}

func (diagramfloss *DiagramFloss) CommitVoid(stage *Stage) {
	diagramfloss.Commit(stage)
}

func (diagramfloss *DiagramFloss) StageVoid(stage *Stage) {
	diagramfloss.Stage(stage)
}

// Checkout diagramfloss to the back repo (if it is already staged)
func (diagramfloss *DiagramFloss) Checkout(stage *Stage) *DiagramFloss {
	if _, ok := stage.DiagramFlosss[diagramfloss]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDiagramFloss(diagramfloss)
		}
	}
	return diagramfloss
}

// for satisfaction of GongStruct interface
func (diagramfloss *DiagramFloss) GetName() (res string) {
	return diagramfloss.Name
}

// for satisfaction of GongStruct interface
func (diagramfloss *DiagramFloss) SetName(name string) {
	diagramfloss.Name = name
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

// Stage puts system to the model stage
func (system *System) Stage(stage *Stage) *System {
	if _, ok := stage.Systems[system]; !ok {
		stage.Systems[system] = struct{}{}
		stage.System_stagedOrder[system] = stage.SystemOrder
		stage.System_orderStaged[stage.SystemOrder] = system
		stage.SystemOrder++
	}
	stage.Systems_mapString[system.Name] = system

	return system
}

// StagePreserveOrder puts system to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SystemOrder
// - update stage.SystemOrder accordingly
func (system *System) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Systems[system]; !ok {
		stage.Systems[system] = struct{}{}

		if order > stage.SystemOrder {
			stage.SystemOrder = order
		}
		stage.System_stagedOrder[system] = order
		stage.System_orderStaged[order] = system
		stage.SystemOrder++
	}
	stage.Systems_mapString[system.Name] = system
}

// Unstage removes system off the model stage
func (system *System) Unstage(stage *Stage) *System {
	delete(stage.Systems, system)
	// issue1150
	// delete(stage.System_stagedOrder, system)
	delete(stage.Systems_mapString, system.Name)

	return system
}

// UnstageVoid removes system off the model stage
func (system *System) UnstageVoid(stage *Stage) {
	delete(stage.Systems, system)
	// issue1150
	// delete(stage.System_stagedOrder, system)
	delete(stage.Systems_mapString, system.Name)
}

// commit system to the back repo (if it is already staged)
func (system *System) Commit(stage *Stage) *System {
	if _, ok := stage.Systems[system]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSystem(system)
		}
	}
	return system
}

func (system *System) CommitVoid(stage *Stage) {
	system.Commit(stage)
}

func (system *System) StageVoid(stage *Stage) {
	system.Stage(stage)
}

// Checkout system to the back repo (if it is already staged)
func (system *System) Checkout(stage *Stage) *System {
	if _, ok := stage.Systems[system]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSystem(system)
		}
	}
	return system
}

// for satisfaction of GongStruct interface
func (system *System) GetName() (res string) {
	return system.Name
}

// for satisfaction of GongStruct interface
func (system *System) SetName(name string) {
	system.Name = name
}

// Stage puts systemshape to the model stage
func (systemshape *SystemShape) Stage(stage *Stage) *SystemShape {
	if _, ok := stage.SystemShapes[systemshape]; !ok {
		stage.SystemShapes[systemshape] = struct{}{}
		stage.SystemShape_stagedOrder[systemshape] = stage.SystemShapeOrder
		stage.SystemShape_orderStaged[stage.SystemShapeOrder] = systemshape
		stage.SystemShapeOrder++
	}
	stage.SystemShapes_mapString[systemshape.Name] = systemshape

	return systemshape
}

// StagePreserveOrder puts systemshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SystemShapeOrder
// - update stage.SystemShapeOrder accordingly
func (systemshape *SystemShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.SystemShapes[systemshape]; !ok {
		stage.SystemShapes[systemshape] = struct{}{}

		if order > stage.SystemShapeOrder {
			stage.SystemShapeOrder = order
		}
		stage.SystemShape_stagedOrder[systemshape] = order
		stage.SystemShape_orderStaged[order] = systemshape
		stage.SystemShapeOrder++
	}
	stage.SystemShapes_mapString[systemshape.Name] = systemshape
}

// Unstage removes systemshape off the model stage
func (systemshape *SystemShape) Unstage(stage *Stage) *SystemShape {
	delete(stage.SystemShapes, systemshape)
	// issue1150
	// delete(stage.SystemShape_stagedOrder, systemshape)
	delete(stage.SystemShapes_mapString, systemshape.Name)

	return systemshape
}

// UnstageVoid removes systemshape off the model stage
func (systemshape *SystemShape) UnstageVoid(stage *Stage) {
	delete(stage.SystemShapes, systemshape)
	// issue1150
	// delete(stage.SystemShape_stagedOrder, systemshape)
	delete(stage.SystemShapes_mapString, systemshape.Name)
}

// commit systemshape to the back repo (if it is already staged)
func (systemshape *SystemShape) Commit(stage *Stage) *SystemShape {
	if _, ok := stage.SystemShapes[systemshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSystemShape(systemshape)
		}
	}
	return systemshape
}

func (systemshape *SystemShape) CommitVoid(stage *Stage) {
	systemshape.Commit(stage)
}

func (systemshape *SystemShape) StageVoid(stage *Stage) {
	systemshape.Stage(stage)
}

// Checkout systemshape to the back repo (if it is already staged)
func (systemshape *SystemShape) Checkout(stage *Stage) *SystemShape {
	if _, ok := stage.SystemShapes[systemshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSystemShape(systemshape)
		}
	}
	return systemshape
}

// for satisfaction of GongStruct interface
func (systemshape *SystemShape) GetName() (res string) {
	return systemshape.Name
}

// for satisfaction of GongStruct interface
func (systemshape *SystemShape) SetName(name string) {
	systemshape.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMDiagramFloss(DiagramFloss *DiagramFloss)
	CreateORMLibrary(Library *Library)
	CreateORMSystem(System *System)
	CreateORMSystemShape(SystemShape *SystemShape)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMDiagramFloss(DiagramFloss *DiagramFloss)
	DeleteORMLibrary(Library *Library)
	DeleteORMSystem(System *System)
	DeleteORMSystemShape(SystemShape *SystemShape)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.DiagramFlosss = make(map[*DiagramFloss]struct{})
	stage.DiagramFlosss_mapString = make(map[string]*DiagramFloss)
	stage.DiagramFloss_stagedOrder = make(map[*DiagramFloss]uint)
	stage.DiagramFlossOrder = 0

	stage.Librarys = make(map[*Library]struct{})
	stage.Librarys_mapString = make(map[string]*Library)
	stage.Library_stagedOrder = make(map[*Library]uint)
	stage.LibraryOrder = 0

	stage.Systems = make(map[*System]struct{})
	stage.Systems_mapString = make(map[string]*System)
	stage.System_stagedOrder = make(map[*System]uint)
	stage.SystemOrder = 0

	stage.SystemShapes = make(map[*SystemShape]struct{})
	stage.SystemShapes_mapString = make(map[string]*SystemShape)
	stage.SystemShape_stagedOrder = make(map[*SystemShape]uint)
	stage.SystemShapeOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.DiagramFlosss = nil
	stage.DiagramFlosss_mapString = nil

	stage.Librarys = nil
	stage.Librarys_mapString = nil

	stage.Systems = nil
	stage.Systems_mapString = nil

	stage.SystemShapes = nil
	stage.SystemShapes_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for diagramfloss := range stage.DiagramFlosss {
		diagramfloss.Unstage(stage)
	}

	for library := range stage.Librarys {
		library.Unstage(stage)
	}

	for system := range stage.Systems {
		system.Unstage(stage)
	}

	for systemshape := range stage.SystemShapes {
		systemshape.Unstage(stage)
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
	case map[*DiagramFloss]any:
		return any(&stage.DiagramFlosss).(*Type)
	case map[*Library]any:
		return any(&stage.Librarys).(*Type)
	case map[*System]any:
		return any(&stage.Systems).(*Type)
	case map[*SystemShape]any:
		return any(&stage.SystemShapes).(*Type)
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
	case *DiagramFloss:
		return any(stage.DiagramFlosss_mapString).(map[string]Type)
	case *Library:
		return any(stage.Librarys_mapString).(map[string]Type)
	case *System:
		return any(stage.Systems_mapString).(map[string]Type)
	case *SystemShape:
		return any(stage.SystemShapes_mapString).(map[string]Type)
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
	case DiagramFloss:
		return any(&stage.DiagramFlosss).(*map[*Type]struct{})
	case Library:
		return any(&stage.Librarys).(*map[*Type]struct{})
	case System:
		return any(&stage.Systems).(*map[*Type]struct{})
	case SystemShape:
		return any(&stage.SystemShapes).(*map[*Type]struct{})
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
	case *DiagramFloss:
		return any(&stage.DiagramFlosss).(*map[Type]struct{})
	case *Library:
		return any(&stage.Librarys).(*map[Type]struct{})
	case *System:
		return any(&stage.Systems).(*map[Type]struct{})
	case *SystemShape:
		return any(&stage.SystemShapes).(*map[Type]struct{})
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
	case DiagramFloss:
		return any(&stage.DiagramFlosss_mapString).(*map[string]*Type)
	case Library:
		return any(&stage.Librarys_mapString).(*map[string]*Type)
	case System:
		return any(&stage.Systems_mapString).(*map[string]*Type)
	case SystemShape:
		return any(&stage.SystemShapes_mapString).(*map[string]*Type)
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
	case DiagramFloss:
		return any(&DiagramFloss{
			// Initialisation of associations
			// field is initialized with an instance of SystemShape with the name of the field
			System_Shapes: []*SystemShape{{Name: "System_Shapes"}},
			// field is initialized with an instance of System with the name of the field
			SystemsWhoseNodeIsExpanded: []*System{{Name: "SystemsWhoseNodeIsExpanded"}},
		}).(*Type)
	case Library:
		return any(&Library{
			// Initialisation of associations
			// field is initialized with an instance of Library with the name of the field
			SubLibraries: []*Library{{Name: "SubLibraries"}},
			// field is initialized with an instance of Library with the name of the field
			SubLibrariesWhoseNodeIsExpanded: []*Library{{Name: "SubLibrariesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of System with the name of the field
			RootSystemes: []*System{{Name: "RootSystemes"}},
			// field is initialized with an instance of System with the name of the field
			SystemsWhoseNodeIsExpanded: []*System{{Name: "SystemsWhoseNodeIsExpanded"}},
		}).(*Type)
	case System:
		return any(&System{
			// Initialisation of associations
			// field is initialized with an instance of DiagramFloss with the name of the field
			DiagramFlosses: []*DiagramFloss{{Name: "DiagramFlosses"}},
			// field is initialized with an instance of DiagramFloss with the name of the field
			DiagramFlossWhoseNodeIsExpanded: []*DiagramFloss{{Name: "DiagramFlossWhoseNodeIsExpanded"}},
			// field is initialized with an instance of System with the name of the field
			SubSystemes: []*System{{Name: "SubSystemes"}},
		}).(*Type)
	case SystemShape:
		return any(&SystemShape{
			// Initialisation of associations
			// field is initialized with an instance of System with the name of the field
			System: &System{Name: "System"},
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
	// reverse maps of direct associations of DiagramFloss
	case DiagramFloss:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Library
	case Library:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of System
	case System:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SystemShape
	case SystemShape:
		switch fieldname {
		// insertion point for per direct association field
		case "System":
			res := make(map[*System][]*SystemShape)
			for systemshape := range stage.SystemShapes {
				if systemshape.System != nil {
					system_ := systemshape.System
					var systemshapes []*SystemShape
					_, ok := res[system_]
					if ok {
						systemshapes = res[system_]
					} else {
						systemshapes = make([]*SystemShape, 0)
					}
					systemshapes = append(systemshapes, systemshape)
					res[system_] = systemshapes
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
	// reverse maps of direct associations of DiagramFloss
	case DiagramFloss:
		switch fieldname {
		// insertion point for per direct association field
		case "System_Shapes":
			res := make(map[*SystemShape][]*DiagramFloss)
			for diagramfloss := range stage.DiagramFlosss {
				for _, systemshape_ := range diagramfloss.System_Shapes {
					res[systemshape_] = append(res[systemshape_], diagramfloss)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SystemsWhoseNodeIsExpanded":
			res := make(map[*System][]*DiagramFloss)
			for diagramfloss := range stage.DiagramFlosss {
				for _, system_ := range diagramfloss.SystemsWhoseNodeIsExpanded {
					res[system_] = append(res[system_], diagramfloss)
				}
			}
			return any(res).(map[*End][]*Start)
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
		case "SubLibrariesWhoseNodeIsExpanded":
			res := make(map[*Library][]*Library)
			for library := range stage.Librarys {
				for _, library_ := range library.SubLibrariesWhoseNodeIsExpanded {
					res[library_] = append(res[library_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootSystemes":
			res := make(map[*System][]*Library)
			for library := range stage.Librarys {
				for _, system_ := range library.RootSystemes {
					res[system_] = append(res[system_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SystemsWhoseNodeIsExpanded":
			res := make(map[*System][]*Library)
			for library := range stage.Librarys {
				for _, system_ := range library.SystemsWhoseNodeIsExpanded {
					res[system_] = append(res[system_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of System
	case System:
		switch fieldname {
		// insertion point for per direct association field
		case "DiagramFlosses":
			res := make(map[*DiagramFloss][]*System)
			for system := range stage.Systems {
				for _, diagramfloss_ := range system.DiagramFlosses {
					res[diagramfloss_] = append(res[diagramfloss_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DiagramFlossWhoseNodeIsExpanded":
			res := make(map[*DiagramFloss][]*System)
			for system := range stage.Systems {
				for _, diagramfloss_ := range system.DiagramFlossWhoseNodeIsExpanded {
					res[diagramfloss_] = append(res[diagramfloss_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SubSystemes":
			res := make(map[*System][]*System)
			for system := range stage.Systems {
				for _, system_ := range system.SubSystemes {
					res[system_] = append(res[system_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SystemShape
	case SystemShape:
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
	case *DiagramFloss:
		res = "DiagramFloss"
	case *Library:
		res = "Library"
	case *System:
		res = "System"
	case *SystemShape:
		res = "SystemShape"
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
	case *DiagramFloss:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "System"
		rf.Fieldname = "DiagramFlosses"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "DiagramFlossWhoseNodeIsExpanded"
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
	case *System:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramFloss"
		rf.Fieldname = "SystemsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootSystemes"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "SystemsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "SubSystemes"
		res = append(res, rf)
	case *SystemShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramFloss"
		rf.Fieldname = "System_Shapes"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
func (diagramfloss *DiagramFloss) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "IsChecked",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsEditable_",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsShowPrefix",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "DefaultBoxWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "DefaultBoxHeigth",
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
			Name:                 "System_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "SystemShape",
		},
		{
			Name:               "IsSystemsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SystemsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "System",
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
			Name:                 "RootSystemes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "System",
		},
		{
			Name:               "IsSystemesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SystemsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "System",
		},
		{
			Name:               "IsExpandedTmp",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (system *System) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "SVG_Path",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "InverseAppliedScaling",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "DiagramFlosses",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DiagramFloss",
		},
		{
			Name:                 "DiagramFlossWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DiagramFloss",
		},
		{
			Name:               "IsSubSystemNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SubSystemes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "System",
		},
	}
	return
}

func (systemshape *SystemShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "System",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "System",
		},
		{
			Name:               "IsExpanded",
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
func (diagramfloss *DiagramFloss) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = diagramfloss.Name
	case "Description":
		res.valueString = diagramfloss.Description
	case "ComputedPrefix":
		res.valueString = diagramfloss.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", diagramfloss.IsExpanded)
		res.valueBool = diagramfloss.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsChecked":
		res.valueString = fmt.Sprintf("%t", diagramfloss.IsChecked)
		res.valueBool = diagramfloss.IsChecked
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsEditable_":
		res.valueString = fmt.Sprintf("%t", diagramfloss.IsEditable_)
		res.valueBool = diagramfloss.IsEditable_
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsShowPrefix":
		res.valueString = fmt.Sprintf("%t", diagramfloss.IsShowPrefix)
		res.valueBool = diagramfloss.IsShowPrefix
		res.GongFieldValueType = GongFieldValueTypeBool
	case "DefaultBoxWidth":
		res.valueString = fmt.Sprintf("%f", diagramfloss.DefaultBoxWidth)
		res.valueFloat = diagramfloss.DefaultBoxWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DefaultBoxHeigth":
		res.valueString = fmt.Sprintf("%f", diagramfloss.DefaultBoxHeigth)
		res.valueFloat = diagramfloss.DefaultBoxHeigth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", diagramfloss.Width)
		res.valueFloat = diagramfloss.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", diagramfloss.Height)
		res.valueFloat = diagramfloss.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "System_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramfloss.System_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsSystemsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagramfloss.IsSystemsNodeExpanded)
		res.valueBool = diagramfloss.IsSystemsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SystemsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramfloss.SystemsWhoseNodeIsExpanded {
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
	case "RootSystemes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootSystemes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsSystemesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsSystemesNodeExpanded)
		res.valueBool = library.IsSystemesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SystemsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.SystemsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsExpandedTmp":
		res.valueString = fmt.Sprintf("%t", library.IsExpandedTmp)
		res.valueBool = library.IsExpandedTmp
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (system *System) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = system.Name
	case "Description":
		res.valueString = system.Description
	case "ComputedPrefix":
		res.valueString = system.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", system.IsExpanded)
		res.valueBool = system.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SVG_Path":
		res.valueString = system.SVG_Path
	case "InverseAppliedScaling":
		res.valueString = fmt.Sprintf("%f", system.InverseAppliedScaling)
		res.valueFloat = system.InverseAppliedScaling
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DiagramFlosses":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.DiagramFlosses {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DiagramFlossWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.DiagramFlossWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsSubSystemNodeExpanded":
		res.valueString = fmt.Sprintf("%t", system.IsSubSystemNodeExpanded)
		res.valueBool = system.IsSubSystemNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SubSystemes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.SubSystemes {
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

func (systemshape *SystemShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = systemshape.Name
	case "System":
		res.GongFieldValueType = GongFieldValueTypePointer
		if systemshape.System != nil {
			res.valueString = systemshape.System.Name
			res.ids = systemshape.System.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", systemshape.IsExpanded)
		res.valueBool = systemshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", systemshape.X)
		res.valueFloat = systemshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", systemshape.Y)
		res.valueFloat = systemshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", systemshape.Width)
		res.valueFloat = systemshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", systemshape.Height)
		res.valueFloat = systemshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", systemshape.IsHidden)
		res.valueBool = systemshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (diagramfloss *DiagramFloss) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		diagramfloss.Name = value.GetValueString()
	case "Description":
		diagramfloss.Description = value.GetValueString()
	case "ComputedPrefix":
		diagramfloss.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		diagramfloss.IsExpanded = value.GetValueBool()
	case "IsChecked":
		diagramfloss.IsChecked = value.GetValueBool()
	case "IsEditable_":
		diagramfloss.IsEditable_ = value.GetValueBool()
	case "IsShowPrefix":
		diagramfloss.IsShowPrefix = value.GetValueBool()
	case "DefaultBoxWidth":
		diagramfloss.DefaultBoxWidth = value.GetValueFloat()
	case "DefaultBoxHeigth":
		diagramfloss.DefaultBoxHeigth = value.GetValueFloat()
	case "Width":
		diagramfloss.Width = value.GetValueFloat()
	case "Height":
		diagramfloss.Height = value.GetValueFloat()
	case "System_Shapes":
		diagramfloss.System_Shapes = make([]*SystemShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.SystemShapes {
					if stage.SystemShape_stagedOrder[__instance__] == uint(id) {
						diagramfloss.System_Shapes = append(diagramfloss.System_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "IsSystemsNodeExpanded":
		diagramfloss.IsSystemsNodeExpanded = value.GetValueBool()
	case "SystemsWhoseNodeIsExpanded":
		diagramfloss.SystemsWhoseNodeIsExpanded = make([]*System, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Systems {
					if stage.System_stagedOrder[__instance__] == uint(id) {
						diagramfloss.SystemsWhoseNodeIsExpanded = append(diagramfloss.SystemsWhoseNodeIsExpanded, __instance__)
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
	case "IsRootLibrary":
		library.IsRootLibrary = value.GetValueBool()
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
	case "RootSystemes":
		library.RootSystemes = make([]*System, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Systems {
					if stage.System_stagedOrder[__instance__] == uint(id) {
						library.RootSystemes = append(library.RootSystemes, __instance__)
						break
					}
				}
			}
		}
	case "IsSystemesNodeExpanded":
		library.IsSystemesNodeExpanded = value.GetValueBool()
	case "SystemsWhoseNodeIsExpanded":
		library.SystemsWhoseNodeIsExpanded = make([]*System, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Systems {
					if stage.System_stagedOrder[__instance__] == uint(id) {
						library.SystemsWhoseNodeIsExpanded = append(library.SystemsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsExpandedTmp":
		library.IsExpandedTmp = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (system *System) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		system.Name = value.GetValueString()
	case "Description":
		system.Description = value.GetValueString()
	case "ComputedPrefix":
		system.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		system.IsExpanded = value.GetValueBool()
	case "SVG_Path":
		system.SVG_Path = value.GetValueString()
	case "InverseAppliedScaling":
		system.InverseAppliedScaling = value.GetValueFloat()
	case "DiagramFlosses":
		system.DiagramFlosses = make([]*DiagramFloss, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DiagramFlosss {
					if stage.DiagramFloss_stagedOrder[__instance__] == uint(id) {
						system.DiagramFlosses = append(system.DiagramFlosses, __instance__)
						break
					}
				}
			}
		}
	case "DiagramFlossWhoseNodeIsExpanded":
		system.DiagramFlossWhoseNodeIsExpanded = make([]*DiagramFloss, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DiagramFlosss {
					if stage.DiagramFloss_stagedOrder[__instance__] == uint(id) {
						system.DiagramFlossWhoseNodeIsExpanded = append(system.DiagramFlossWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsSubSystemNodeExpanded":
		system.IsSubSystemNodeExpanded = value.GetValueBool()
	case "SubSystemes":
		system.SubSystemes = make([]*System, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Systems {
					if stage.System_stagedOrder[__instance__] == uint(id) {
						system.SubSystemes = append(system.SubSystemes, __instance__)
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

func (systemshape *SystemShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		systemshape.Name = value.GetValueString()
	case "System":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			systemshape.System = nil
			for __instance__ := range stage.Systems {
				if stage.System_stagedOrder[__instance__] == uint(id) {
					systemshape.System = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		systemshape.IsExpanded = value.GetValueBool()
	case "X":
		systemshape.X = value.GetValueFloat()
	case "Y":
		systemshape.Y = value.GetValueFloat()
	case "Width":
		systemshape.Width = value.GetValueFloat()
	case "Height":
		systemshape.Height = value.GetValueFloat()
	case "IsHidden":
		systemshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (diagramfloss *DiagramFloss) GongGetGongstructName() string {
	return "DiagramFloss"
}

func (library *Library) GongGetGongstructName() string {
	return "Library"
}

func (system *System) GongGetGongstructName() string {
	return "System"
}

func (systemshape *SystemShape) GongGetGongstructName() string {
	return "SystemShape"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	stage.DiagramFlosss_mapString = make(map[string]*DiagramFloss)
	for diagramfloss := range stage.DiagramFlosss {
		stage.DiagramFlosss_mapString[diagramfloss.Name] = diagramfloss
	}

	stage.Librarys_mapString = make(map[string]*Library)
	for library := range stage.Librarys {
		stage.Librarys_mapString[library.Name] = library
	}

	stage.Systems_mapString = make(map[string]*System)
	for system := range stage.Systems {
		stage.Systems_mapString[system.Name] = system
	}

	stage.SystemShapes_mapString = make(map[string]*SystemShape)
	for systemshape := range stage.SystemShapes {
		stage.SystemShapes_mapString[systemshape.Name] = systemshape
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
