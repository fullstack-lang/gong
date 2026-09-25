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
	GongBasicFields                map[*GongBasicField]struct{}
	GongBasicFields_instance       map[*GongBasicField]*GongBasicField
	GongBasicFields_mapString      map[string]*GongBasicField
	GongBasicFieldOrder            uint
	GongBasicField_stagedOrder     map[*GongBasicField]uint
	GongBasicField_orderStaged     map[uint]*GongBasicField
	GongBasicFields_reference      map[*GongBasicField]*GongBasicField
	GongBasicFields_referenceOrder map[*GongBasicField]uint

	// insertion point for slice of pointers maps
	OnAfterGongBasicFieldCreateCallback GongOnAfterCreateInterface[GongBasicField]
	OnAfterGongBasicFieldUpdateCallback GongOnAfterUpdateInterface[GongBasicField]
	OnAfterGongBasicFieldDeleteCallback GongOnAfterDeleteInterface[GongBasicField]

	GongEnums                map[*GongEnum]struct{}
	GongEnums_instance       map[*GongEnum]*GongEnum
	GongEnums_mapString      map[string]*GongEnum
	GongEnumOrder            uint
	GongEnum_stagedOrder     map[*GongEnum]uint
	GongEnum_orderStaged     map[uint]*GongEnum
	GongEnums_reference      map[*GongEnum]*GongEnum
	GongEnums_referenceOrder map[*GongEnum]uint

	// insertion point for slice of pointers maps
	GongEnum_GongEnumValues_reverseMap map[*GongEnumValue]*GongEnum

	OnAfterGongEnumCreateCallback GongOnAfterCreateInterface[GongEnum]
	OnAfterGongEnumUpdateCallback GongOnAfterUpdateInterface[GongEnum]
	OnAfterGongEnumDeleteCallback GongOnAfterDeleteInterface[GongEnum]

	GongEnumValues                map[*GongEnumValue]struct{}
	GongEnumValues_instance       map[*GongEnumValue]*GongEnumValue
	GongEnumValues_mapString      map[string]*GongEnumValue
	GongEnumValueOrder            uint
	GongEnumValue_stagedOrder     map[*GongEnumValue]uint
	GongEnumValue_orderStaged     map[uint]*GongEnumValue
	GongEnumValues_reference      map[*GongEnumValue]*GongEnumValue
	GongEnumValues_referenceOrder map[*GongEnumValue]uint

	// insertion point for slice of pointers maps
	OnAfterGongEnumValueCreateCallback GongOnAfterCreateInterface[GongEnumValue]
	OnAfterGongEnumValueUpdateCallback GongOnAfterUpdateInterface[GongEnumValue]
	OnAfterGongEnumValueDeleteCallback GongOnAfterDeleteInterface[GongEnumValue]

	GongLinks                map[*GongLink]struct{}
	GongLinks_instance       map[*GongLink]*GongLink
	GongLinks_mapString      map[string]*GongLink
	GongLinkOrder            uint
	GongLink_stagedOrder     map[*GongLink]uint
	GongLink_orderStaged     map[uint]*GongLink
	GongLinks_reference      map[*GongLink]*GongLink
	GongLinks_referenceOrder map[*GongLink]uint

	// insertion point for slice of pointers maps
	OnAfterGongLinkCreateCallback GongOnAfterCreateInterface[GongLink]
	OnAfterGongLinkUpdateCallback GongOnAfterUpdateInterface[GongLink]
	OnAfterGongLinkDeleteCallback GongOnAfterDeleteInterface[GongLink]

	GongNotes                map[*GongNote]struct{}
	GongNotes_instance       map[*GongNote]*GongNote
	GongNotes_mapString      map[string]*GongNote
	GongNoteOrder            uint
	GongNote_stagedOrder     map[*GongNote]uint
	GongNote_orderStaged     map[uint]*GongNote
	GongNotes_reference      map[*GongNote]*GongNote
	GongNotes_referenceOrder map[*GongNote]uint

	// insertion point for slice of pointers maps
	GongNote_Links_reverseMap map[*GongLink]*GongNote

	OnAfterGongNoteCreateCallback GongOnAfterCreateInterface[GongNote]
	OnAfterGongNoteUpdateCallback GongOnAfterUpdateInterface[GongNote]
	OnAfterGongNoteDeleteCallback GongOnAfterDeleteInterface[GongNote]

	GongStructs                map[*GongStruct]struct{}
	GongStructs_instance       map[*GongStruct]*GongStruct
	GongStructs_mapString      map[string]*GongStruct
	GongStructOrder            uint
	GongStruct_stagedOrder     map[*GongStruct]uint
	GongStruct_orderStaged     map[uint]*GongStruct
	GongStructs_reference      map[*GongStruct]*GongStruct
	GongStructs_referenceOrder map[*GongStruct]uint

	// insertion point for slice of pointers maps
	GongStruct_GongBasicFields_reverseMap map[*GongBasicField]*GongStruct

	GongStruct_GongTimeFields_reverseMap map[*GongTimeField]*GongStruct

	GongStruct_PointerToGongStructFields_reverseMap map[*PointerToGongStructField]*GongStruct

	GongStruct_SliceOfPointerToGongStructFields_reverseMap map[*SliceOfPointerToGongStructField]*GongStruct

	OnAfterGongStructCreateCallback GongOnAfterCreateInterface[GongStruct]
	OnAfterGongStructUpdateCallback GongOnAfterUpdateInterface[GongStruct]
	OnAfterGongStructDeleteCallback GongOnAfterDeleteInterface[GongStruct]

	GongTimeFields                map[*GongTimeField]struct{}
	GongTimeFields_instance       map[*GongTimeField]*GongTimeField
	GongTimeFields_mapString      map[string]*GongTimeField
	GongTimeFieldOrder            uint
	GongTimeField_stagedOrder     map[*GongTimeField]uint
	GongTimeField_orderStaged     map[uint]*GongTimeField
	GongTimeFields_reference      map[*GongTimeField]*GongTimeField
	GongTimeFields_referenceOrder map[*GongTimeField]uint

	// insertion point for slice of pointers maps
	OnAfterGongTimeFieldCreateCallback GongOnAfterCreateInterface[GongTimeField]
	OnAfterGongTimeFieldUpdateCallback GongOnAfterUpdateInterface[GongTimeField]
	OnAfterGongTimeFieldDeleteCallback GongOnAfterDeleteInterface[GongTimeField]

	MetaReferences                map[*MetaReference]struct{}
	MetaReferences_instance       map[*MetaReference]*MetaReference
	MetaReferences_mapString      map[string]*MetaReference
	MetaReferenceOrder            uint
	MetaReference_stagedOrder     map[*MetaReference]uint
	MetaReference_orderStaged     map[uint]*MetaReference
	MetaReferences_reference      map[*MetaReference]*MetaReference
	MetaReferences_referenceOrder map[*MetaReference]uint

	// insertion point for slice of pointers maps
	OnAfterMetaReferenceCreateCallback GongOnAfterCreateInterface[MetaReference]
	OnAfterMetaReferenceUpdateCallback GongOnAfterUpdateInterface[MetaReference]
	OnAfterMetaReferenceDeleteCallback GongOnAfterDeleteInterface[MetaReference]

	ModelPkgs                map[*ModelPkg]struct{}
	ModelPkgs_instance       map[*ModelPkg]*ModelPkg
	ModelPkgs_mapString      map[string]*ModelPkg
	ModelPkgOrder            uint
	ModelPkg_stagedOrder     map[*ModelPkg]uint
	ModelPkg_orderStaged     map[uint]*ModelPkg
	ModelPkgs_reference      map[*ModelPkg]*ModelPkg
	ModelPkgs_referenceOrder map[*ModelPkg]uint

	// insertion point for slice of pointers maps
	OnAfterModelPkgCreateCallback GongOnAfterCreateInterface[ModelPkg]
	OnAfterModelPkgUpdateCallback GongOnAfterUpdateInterface[ModelPkg]
	OnAfterModelPkgDeleteCallback GongOnAfterDeleteInterface[ModelPkg]

	PointerToGongStructFields                map[*PointerToGongStructField]struct{}
	PointerToGongStructFields_instance       map[*PointerToGongStructField]*PointerToGongStructField
	PointerToGongStructFields_mapString      map[string]*PointerToGongStructField
	PointerToGongStructFieldOrder            uint
	PointerToGongStructField_stagedOrder     map[*PointerToGongStructField]uint
	PointerToGongStructField_orderStaged     map[uint]*PointerToGongStructField
	PointerToGongStructFields_reference      map[*PointerToGongStructField]*PointerToGongStructField
	PointerToGongStructFields_referenceOrder map[*PointerToGongStructField]uint

	// insertion point for slice of pointers maps
	OnAfterPointerToGongStructFieldCreateCallback GongOnAfterCreateInterface[PointerToGongStructField]
	OnAfterPointerToGongStructFieldUpdateCallback GongOnAfterUpdateInterface[PointerToGongStructField]
	OnAfterPointerToGongStructFieldDeleteCallback GongOnAfterDeleteInterface[PointerToGongStructField]

	SliceOfPointerToGongStructFields                map[*SliceOfPointerToGongStructField]struct{}
	SliceOfPointerToGongStructFields_instance       map[*SliceOfPointerToGongStructField]*SliceOfPointerToGongStructField
	SliceOfPointerToGongStructFields_mapString      map[string]*SliceOfPointerToGongStructField
	SliceOfPointerToGongStructFieldOrder            uint
	SliceOfPointerToGongStructField_stagedOrder     map[*SliceOfPointerToGongStructField]uint
	SliceOfPointerToGongStructField_orderStaged     map[uint]*SliceOfPointerToGongStructField
	SliceOfPointerToGongStructFields_reference      map[*SliceOfPointerToGongStructField]*SliceOfPointerToGongStructField
	SliceOfPointerToGongStructFields_referenceOrder map[*SliceOfPointerToGongStructField]uint

	// insertion point for slice of pointers maps
	OnAfterSliceOfPointerToGongStructFieldCreateCallback GongOnAfterCreateInterface[SliceOfPointerToGongStructField]
	OnAfterSliceOfPointerToGongStructFieldUpdateCallback GongOnAfterUpdateInterface[SliceOfPointerToGongStructField]
	OnAfterSliceOfPointerToGongStructFieldDeleteCallback GongOnAfterDeleteInterface[SliceOfPointerToGongStructField]

	StageSetFields                map[*StageSetField]struct{}
	StageSetFields_instance       map[*StageSetField]*StageSetField
	StageSetFields_mapString      map[string]*StageSetField
	StageSetFieldOrder            uint
	StageSetField_stagedOrder     map[*StageSetField]uint
	StageSetField_orderStaged     map[uint]*StageSetField
	StageSetFields_reference      map[*StageSetField]*StageSetField
	StageSetFields_referenceOrder map[*StageSetField]uint

	// insertion point for slice of pointers maps
	OnAfterStageSetFieldCreateCallback GongOnAfterCreateInterface[StageSetField]
	OnAfterStageSetFieldUpdateCallback GongOnAfterUpdateInterface[StageSetField]
	OnAfterStageSetFieldDeleteCallback GongOnAfterDeleteInterface[StageSetField]

	StageSetModels                map[*StageSetModel]struct{}
	StageSetModels_instance       map[*StageSetModel]*StageSetModel
	StageSetModels_mapString      map[string]*StageSetModel
	StageSetModelOrder            uint
	StageSetModel_stagedOrder     map[*StageSetModel]uint
	StageSetModel_orderStaged     map[uint]*StageSetModel
	StageSetModels_reference      map[*StageSetModel]*StageSetModel
	StageSetModels_referenceOrder map[*StageSetModel]uint

	// insertion point for slice of pointers maps
	StageSetModel_Fields_reverseMap map[*StageSetField]*StageSetModel

	OnAfterStageSetModelCreateCallback GongOnAfterCreateInterface[StageSetModel]
	OnAfterStageSetModelUpdateCallback GongOnAfterUpdateInterface[StageSetModel]
	OnAfterStageSetModelDeleteCallback GongOnAfterDeleteInterface[StageSetModel]

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
	__gong__clearReferences(&stage.GongBasicFields_reference, &stage.GongBasicFields_instance, &stage.GongBasicFields_referenceOrder)

	__gong__clearReferences(&stage.GongEnums_reference, &stage.GongEnums_instance, &stage.GongEnums_referenceOrder)

	__gong__clearReferences(&stage.GongEnumValues_reference, &stage.GongEnumValues_instance, &stage.GongEnumValues_referenceOrder)

	__gong__clearReferences(&stage.GongLinks_reference, &stage.GongLinks_instance, &stage.GongLinks_referenceOrder)

	__gong__clearReferences(&stage.GongNotes_reference, &stage.GongNotes_instance, &stage.GongNotes_referenceOrder)

	__gong__clearReferences(&stage.GongStructs_reference, &stage.GongStructs_instance, &stage.GongStructs_referenceOrder)

	__gong__clearReferences(&stage.GongTimeFields_reference, &stage.GongTimeFields_instance, &stage.GongTimeFields_referenceOrder)

	__gong__clearReferences(&stage.MetaReferences_reference, &stage.MetaReferences_instance, &stage.MetaReferences_referenceOrder)

	__gong__clearReferences(&stage.ModelPkgs_reference, &stage.ModelPkgs_instance, &stage.ModelPkgs_referenceOrder)

	__gong__clearReferences(&stage.PointerToGongStructFields_reference, &stage.PointerToGongStructFields_instance, &stage.PointerToGongStructFields_referenceOrder)

	__gong__clearReferences(&stage.SliceOfPointerToGongStructFields_reference, &stage.SliceOfPointerToGongStructFields_instance, &stage.SliceOfPointerToGongStructFields_referenceOrder)

	__gong__clearReferences(&stage.StageSetFields_reference, &stage.StageSetFields_instance, &stage.StageSetFields_referenceOrder)

	__gong__clearReferences(&stage.StageSetModels_reference, &stage.StageSetModels_instance, &stage.StageSetModels_referenceOrder)

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
	stage.GongBasicFieldOrder = __gong__recomputeOrder(stage.GongBasicField_stagedOrder)

	stage.GongEnumOrder = __gong__recomputeOrder(stage.GongEnum_stagedOrder)

	stage.GongEnumValueOrder = __gong__recomputeOrder(stage.GongEnumValue_stagedOrder)

	stage.GongLinkOrder = __gong__recomputeOrder(stage.GongLink_stagedOrder)

	stage.GongNoteOrder = __gong__recomputeOrder(stage.GongNote_stagedOrder)

	stage.GongStructOrder = __gong__recomputeOrder(stage.GongStruct_stagedOrder)

	stage.GongTimeFieldOrder = __gong__recomputeOrder(stage.GongTimeField_stagedOrder)

	stage.MetaReferenceOrder = __gong__recomputeOrder(stage.MetaReference_stagedOrder)

	stage.ModelPkgOrder = __gong__recomputeOrder(stage.ModelPkg_stagedOrder)

	stage.PointerToGongStructFieldOrder = __gong__recomputeOrder(stage.PointerToGongStructField_stagedOrder)

	stage.SliceOfPointerToGongStructFieldOrder = __gong__recomputeOrder(stage.SliceOfPointerToGongStructField_stagedOrder)

	stage.StageSetFieldOrder = __gong__recomputeOrder(stage.StageSetField_stagedOrder)

	stage.StageSetModelOrder = __gong__recomputeOrder(stage.StageSetModel_stagedOrder)

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
	case *GongBasicField:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.GongBasicFields, stage.GongBasicField_stagedOrder))
	case *GongEnum:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.GongEnums, stage.GongEnum_stagedOrder))
	case *GongEnumValue:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.GongEnumValues, stage.GongEnumValue_stagedOrder))
	case *GongLink:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.GongLinks, stage.GongLink_stagedOrder))
	case *GongNote:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.GongNotes, stage.GongNote_stagedOrder))
	case *GongStruct:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.GongStructs, stage.GongStruct_stagedOrder))
	case *GongTimeField:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.GongTimeFields, stage.GongTimeField_stagedOrder))
	case *MetaReference:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.MetaReferences, stage.MetaReference_stagedOrder))
	case *ModelPkg:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ModelPkgs, stage.ModelPkg_stagedOrder))
	case *PointerToGongStructField:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.PointerToGongStructFields, stage.PointerToGongStructField_stagedOrder))
	case *SliceOfPointerToGongStructField:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.SliceOfPointerToGongStructFields, stage.SliceOfPointerToGongStructField_stagedOrder))
	case *StageSetField:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.StageSetFields, stage.StageSetField_stagedOrder))
	case *StageSetModel:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.StageSetModels, stage.StageSetModel_stagedOrder))

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
	return "github.com/fullstack-lang/gong/go/models"
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
		GongBasicFields:           make(map[*GongBasicField]struct{}),
		GongBasicFields_mapString: make(map[string]*GongBasicField),

		GongEnums:           make(map[*GongEnum]struct{}),
		GongEnums_mapString: make(map[string]*GongEnum),

		GongEnumValues:           make(map[*GongEnumValue]struct{}),
		GongEnumValues_mapString: make(map[string]*GongEnumValue),

		GongLinks:           make(map[*GongLink]struct{}),
		GongLinks_mapString: make(map[string]*GongLink),

		GongNotes:           make(map[*GongNote]struct{}),
		GongNotes_mapString: make(map[string]*GongNote),

		GongStructs:           make(map[*GongStruct]struct{}),
		GongStructs_mapString: make(map[string]*GongStruct),

		GongTimeFields:           make(map[*GongTimeField]struct{}),
		GongTimeFields_mapString: make(map[string]*GongTimeField),

		MetaReferences:           make(map[*MetaReference]struct{}),
		MetaReferences_mapString: make(map[string]*MetaReference),

		ModelPkgs:           make(map[*ModelPkg]struct{}),
		ModelPkgs_mapString: make(map[string]*ModelPkg),

		PointerToGongStructFields:           make(map[*PointerToGongStructField]struct{}),
		PointerToGongStructFields_mapString: make(map[string]*PointerToGongStructField),

		SliceOfPointerToGongStructFields:           make(map[*SliceOfPointerToGongStructField]struct{}),
		SliceOfPointerToGongStructFields_mapString: make(map[string]*SliceOfPointerToGongStructField),

		StageSetFields:           make(map[*StageSetField]struct{}),
		StageSetFields_mapString: make(map[string]*StageSetField),

		StageSetModels:           make(map[*StageSetModel]struct{}),
		StageSetModels_mapString: make(map[string]*StageSetModel),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		GongBasicField_stagedOrder: make(map[*GongBasicField]uint),
		GongBasicField_orderStaged: make(map[uint]*GongBasicField),
		GongBasicFields_reference:  make(map[*GongBasicField]*GongBasicField),

		GongEnum_stagedOrder: make(map[*GongEnum]uint),
		GongEnum_orderStaged: make(map[uint]*GongEnum),
		GongEnums_reference:  make(map[*GongEnum]*GongEnum),

		GongEnumValue_stagedOrder: make(map[*GongEnumValue]uint),
		GongEnumValue_orderStaged: make(map[uint]*GongEnumValue),
		GongEnumValues_reference:  make(map[*GongEnumValue]*GongEnumValue),

		GongLink_stagedOrder: make(map[*GongLink]uint),
		GongLink_orderStaged: make(map[uint]*GongLink),
		GongLinks_reference:  make(map[*GongLink]*GongLink),

		GongNote_stagedOrder: make(map[*GongNote]uint),
		GongNote_orderStaged: make(map[uint]*GongNote),
		GongNotes_reference:  make(map[*GongNote]*GongNote),

		GongStruct_stagedOrder: make(map[*GongStruct]uint),
		GongStruct_orderStaged: make(map[uint]*GongStruct),
		GongStructs_reference:  make(map[*GongStruct]*GongStruct),

		GongTimeField_stagedOrder: make(map[*GongTimeField]uint),
		GongTimeField_orderStaged: make(map[uint]*GongTimeField),
		GongTimeFields_reference:  make(map[*GongTimeField]*GongTimeField),

		MetaReference_stagedOrder: make(map[*MetaReference]uint),
		MetaReference_orderStaged: make(map[uint]*MetaReference),
		MetaReferences_reference:  make(map[*MetaReference]*MetaReference),

		ModelPkg_stagedOrder: make(map[*ModelPkg]uint),
		ModelPkg_orderStaged: make(map[uint]*ModelPkg),
		ModelPkgs_reference:  make(map[*ModelPkg]*ModelPkg),

		PointerToGongStructField_stagedOrder: make(map[*PointerToGongStructField]uint),
		PointerToGongStructField_orderStaged: make(map[uint]*PointerToGongStructField),
		PointerToGongStructFields_reference:  make(map[*PointerToGongStructField]*PointerToGongStructField),

		SliceOfPointerToGongStructField_stagedOrder: make(map[*SliceOfPointerToGongStructField]uint),
		SliceOfPointerToGongStructField_orderStaged: make(map[uint]*SliceOfPointerToGongStructField),
		SliceOfPointerToGongStructFields_reference:  make(map[*SliceOfPointerToGongStructField]*SliceOfPointerToGongStructField),

		StageSetField_stagedOrder: make(map[*StageSetField]uint),
		StageSetField_orderStaged: make(map[uint]*StageSetField),
		StageSetFields_reference:  make(map[*StageSetField]*StageSetField),

		StageSetModel_stagedOrder: make(map[*StageSetModel]uint),
		StageSetModel_orderStaged: make(map[uint]*StageSetModel),
		StageSetModels_reference:  make(map[*StageSetModel]*StageSetModel),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"GongBasicField": &GongBasicFieldUnmarshaller{},

			"GongEnum": &GongEnumUnmarshaller{},

			"GongEnumValue": &GongEnumValueUnmarshaller{},

			"GongLink": &GongLinkUnmarshaller{},

			"GongNote": &GongNoteUnmarshaller{},

			"GongStruct": &GongStructUnmarshaller{},

			"GongTimeField": &GongTimeFieldUnmarshaller{},

			"MetaReference": &MetaReferenceUnmarshaller{},

			"ModelPkg": &ModelPkgUnmarshaller{},

			"PointerToGongStructField": &PointerToGongStructFieldUnmarshaller{},

			"SliceOfPointerToGongStructField": &SliceOfPointerToGongStructFieldUnmarshaller{},

			"StageSetField": &StageSetFieldUnmarshaller{},

			"StageSetModel": &StageSetModelUnmarshaller{},

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
	case *GongBasicField:
		return any(stage.GongBasicField_orderStaged[order]).(Type)
	case *GongEnum:
		return any(stage.GongEnum_orderStaged[order]).(Type)
	case *GongEnumValue:
		return any(stage.GongEnumValue_orderStaged[order]).(Type)
	case *GongLink:
		return any(stage.GongLink_orderStaged[order]).(Type)
	case *GongNote:
		return any(stage.GongNote_orderStaged[order]).(Type)
	case *GongStruct:
		return any(stage.GongStruct_orderStaged[order]).(Type)
	case *GongTimeField:
		return any(stage.GongTimeField_orderStaged[order]).(Type)
	case *MetaReference:
		return any(stage.MetaReference_orderStaged[order]).(Type)
	case *ModelPkg:
		return any(stage.ModelPkg_orderStaged[order]).(Type)
	case *PointerToGongStructField:
		return any(stage.PointerToGongStructField_orderStaged[order]).(Type)
	case *SliceOfPointerToGongStructField:
		return any(stage.SliceOfPointerToGongStructField_orderStaged[order]).(Type)
	case *StageSetField:
		return any(stage.StageSetField_orderStaged[order]).(Type)
	case *StageSetModel:
		return any(stage.StageSetModel_orderStaged[order]).(Type)
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
	stage.Map_GongStructName_InstancesNb["GongBasicField"] = len(stage.GongBasicFields)
	stage.Map_GongStructName_InstancesNb["GongEnum"] = len(stage.GongEnums)
	stage.Map_GongStructName_InstancesNb["GongEnumValue"] = len(stage.GongEnumValues)
	stage.Map_GongStructName_InstancesNb["GongLink"] = len(stage.GongLinks)
	stage.Map_GongStructName_InstancesNb["GongNote"] = len(stage.GongNotes)
	stage.Map_GongStructName_InstancesNb["GongStruct"] = len(stage.GongStructs)
	stage.Map_GongStructName_InstancesNb["GongTimeField"] = len(stage.GongTimeFields)
	stage.Map_GongStructName_InstancesNb["MetaReference"] = len(stage.MetaReferences)
	stage.Map_GongStructName_InstancesNb["ModelPkg"] = len(stage.ModelPkgs)
	stage.Map_GongStructName_InstancesNb["PointerToGongStructField"] = len(stage.PointerToGongStructFields)
	stage.Map_GongStructName_InstancesNb["SliceOfPointerToGongStructField"] = len(stage.SliceOfPointerToGongStructFields)
	stage.Map_GongStructName_InstancesNb["StageSetField"] = len(stage.StageSetFields)
	stage.Map_GongStructName_InstancesNb["StageSetModel"] = len(stage.StageSetModels)
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
// Stage puts gongbasicfield to the model stage
func (gongbasicfield *GongBasicField) Stage(stage *Stage) *GongBasicField {
	__gong__stage(stage.GongBasicFields, stage.GongBasicField_stagedOrder, stage.GongBasicField_orderStaged, &stage.GongBasicFieldOrder, stage.GongBasicFields_mapString, gongbasicfield, gongbasicfield.Name)
	return gongbasicfield
}

// StagePreserveOrder puts gongbasicfield to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GongBasicFieldOrder
// - update stage.GongBasicFieldOrder accordingly
func (gongbasicfield *GongBasicField) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.GongBasicFields, stage.GongBasicField_stagedOrder, stage.GongBasicField_orderStaged, &stage.GongBasicFieldOrder, stage.GongBasicFields_mapString, gongbasicfield, order, gongbasicfield.Name)
}

// Unstage removes gongbasicfield off the model stage
func (gongbasicfield *GongBasicField) Unstage(stage *Stage) *GongBasicField {
	__gong__unstage(stage.GongBasicFields, stage.GongBasicFields_mapString, gongbasicfield, gongbasicfield.Name)
	return gongbasicfield
}

// UnstageVoid removes gongbasicfield off the model stage
func (gongbasicfield *GongBasicField) UnstageVoid(stage *Stage) {
	gongbasicfield.Unstage(stage)
}

func (gongbasicfield *GongBasicField) StageVoid(stage *Stage) {
	gongbasicfield.Stage(stage)
}

// for satisfaction of GongStruct interface
func (gongbasicfield *GongBasicField) GetName() (res string) {
	return gongbasicfield.Name
}

// for satisfaction of GongStruct interface
func (gongbasicfield *GongBasicField) SetName(name string) {
	gongbasicfield.Name = name
}

// Stage puts gongenum to the model stage
func (gongenum *GongEnum) Stage(stage *Stage) *GongEnum {
	__gong__stage(stage.GongEnums, stage.GongEnum_stagedOrder, stage.GongEnum_orderStaged, &stage.GongEnumOrder, stage.GongEnums_mapString, gongenum, gongenum.Name)
	return gongenum
}

// StagePreserveOrder puts gongenum to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GongEnumOrder
// - update stage.GongEnumOrder accordingly
func (gongenum *GongEnum) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.GongEnums, stage.GongEnum_stagedOrder, stage.GongEnum_orderStaged, &stage.GongEnumOrder, stage.GongEnums_mapString, gongenum, order, gongenum.Name)
}

// Unstage removes gongenum off the model stage
func (gongenum *GongEnum) Unstage(stage *Stage) *GongEnum {
	__gong__unstage(stage.GongEnums, stage.GongEnums_mapString, gongenum, gongenum.Name)
	return gongenum
}

// UnstageVoid removes gongenum off the model stage
func (gongenum *GongEnum) UnstageVoid(stage *Stage) {
	gongenum.Unstage(stage)
}

func (gongenum *GongEnum) StageVoid(stage *Stage) {
	gongenum.Stage(stage)
}

// for satisfaction of GongStruct interface
func (gongenum *GongEnum) GetName() (res string) {
	return gongenum.Name
}

// for satisfaction of GongStruct interface
func (gongenum *GongEnum) SetName(name string) {
	gongenum.Name = name
}

// Stage puts gongenumvalue to the model stage
func (gongenumvalue *GongEnumValue) Stage(stage *Stage) *GongEnumValue {
	__gong__stage(stage.GongEnumValues, stage.GongEnumValue_stagedOrder, stage.GongEnumValue_orderStaged, &stage.GongEnumValueOrder, stage.GongEnumValues_mapString, gongenumvalue, gongenumvalue.Name)
	return gongenumvalue
}

// StagePreserveOrder puts gongenumvalue to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GongEnumValueOrder
// - update stage.GongEnumValueOrder accordingly
func (gongenumvalue *GongEnumValue) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.GongEnumValues, stage.GongEnumValue_stagedOrder, stage.GongEnumValue_orderStaged, &stage.GongEnumValueOrder, stage.GongEnumValues_mapString, gongenumvalue, order, gongenumvalue.Name)
}

// Unstage removes gongenumvalue off the model stage
func (gongenumvalue *GongEnumValue) Unstage(stage *Stage) *GongEnumValue {
	__gong__unstage(stage.GongEnumValues, stage.GongEnumValues_mapString, gongenumvalue, gongenumvalue.Name)
	return gongenumvalue
}

// UnstageVoid removes gongenumvalue off the model stage
func (gongenumvalue *GongEnumValue) UnstageVoid(stage *Stage) {
	gongenumvalue.Unstage(stage)
}

func (gongenumvalue *GongEnumValue) StageVoid(stage *Stage) {
	gongenumvalue.Stage(stage)
}

// for satisfaction of GongStruct interface
func (gongenumvalue *GongEnumValue) GetName() (res string) {
	return gongenumvalue.Name
}

// for satisfaction of GongStruct interface
func (gongenumvalue *GongEnumValue) SetName(name string) {
	gongenumvalue.Name = name
}

// Stage puts gonglink to the model stage
func (gonglink *GongLink) Stage(stage *Stage) *GongLink {
	__gong__stage(stage.GongLinks, stage.GongLink_stagedOrder, stage.GongLink_orderStaged, &stage.GongLinkOrder, stage.GongLinks_mapString, gonglink, gonglink.Name)
	return gonglink
}

// StagePreserveOrder puts gonglink to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GongLinkOrder
// - update stage.GongLinkOrder accordingly
func (gonglink *GongLink) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.GongLinks, stage.GongLink_stagedOrder, stage.GongLink_orderStaged, &stage.GongLinkOrder, stage.GongLinks_mapString, gonglink, order, gonglink.Name)
}

// Unstage removes gonglink off the model stage
func (gonglink *GongLink) Unstage(stage *Stage) *GongLink {
	__gong__unstage(stage.GongLinks, stage.GongLinks_mapString, gonglink, gonglink.Name)
	return gonglink
}

// UnstageVoid removes gonglink off the model stage
func (gonglink *GongLink) UnstageVoid(stage *Stage) {
	gonglink.Unstage(stage)
}

func (gonglink *GongLink) StageVoid(stage *Stage) {
	gonglink.Stage(stage)
}

// for satisfaction of GongStruct interface
func (gonglink *GongLink) GetName() (res string) {
	return gonglink.Name
}

// for satisfaction of GongStruct interface
func (gonglink *GongLink) SetName(name string) {
	gonglink.Name = name
}

// Stage puts gongnote to the model stage
func (gongnote *GongNote) Stage(stage *Stage) *GongNote {
	__gong__stage(stage.GongNotes, stage.GongNote_stagedOrder, stage.GongNote_orderStaged, &stage.GongNoteOrder, stage.GongNotes_mapString, gongnote, gongnote.Name)
	return gongnote
}

// StagePreserveOrder puts gongnote to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GongNoteOrder
// - update stage.GongNoteOrder accordingly
func (gongnote *GongNote) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.GongNotes, stage.GongNote_stagedOrder, stage.GongNote_orderStaged, &stage.GongNoteOrder, stage.GongNotes_mapString, gongnote, order, gongnote.Name)
}

// Unstage removes gongnote off the model stage
func (gongnote *GongNote) Unstage(stage *Stage) *GongNote {
	__gong__unstage(stage.GongNotes, stage.GongNotes_mapString, gongnote, gongnote.Name)
	return gongnote
}

// UnstageVoid removes gongnote off the model stage
func (gongnote *GongNote) UnstageVoid(stage *Stage) {
	gongnote.Unstage(stage)
}

func (gongnote *GongNote) StageVoid(stage *Stage) {
	gongnote.Stage(stage)
}

// for satisfaction of GongStruct interface
func (gongnote *GongNote) GetName() (res string) {
	return gongnote.Name
}

// for satisfaction of GongStruct interface
func (gongnote *GongNote) SetName(name string) {
	gongnote.Name = name
}

// Stage puts gongstruct to the model stage
func (gongstruct *GongStruct) Stage(stage *Stage) *GongStruct {
	__gong__stage(stage.GongStructs, stage.GongStruct_stagedOrder, stage.GongStruct_orderStaged, &stage.GongStructOrder, stage.GongStructs_mapString, gongstruct, gongstruct.Name)
	return gongstruct
}

// StagePreserveOrder puts gongstruct to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GongStructOrder
// - update stage.GongStructOrder accordingly
func (gongstruct *GongStruct) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.GongStructs, stage.GongStruct_stagedOrder, stage.GongStruct_orderStaged, &stage.GongStructOrder, stage.GongStructs_mapString, gongstruct, order, gongstruct.Name)
}

// Unstage removes gongstruct off the model stage
func (gongstruct *GongStruct) Unstage(stage *Stage) *GongStruct {
	__gong__unstage(stage.GongStructs, stage.GongStructs_mapString, gongstruct, gongstruct.Name)
	return gongstruct
}

// UnstageVoid removes gongstruct off the model stage
func (gongstruct *GongStruct) UnstageVoid(stage *Stage) {
	gongstruct.Unstage(stage)
}

func (gongstruct *GongStruct) StageVoid(stage *Stage) {
	gongstruct.Stage(stage)
}

// for satisfaction of GongStruct interface
func (gongstruct *GongStruct) GetName() (res string) {
	return gongstruct.Name
}

// for satisfaction of GongStruct interface
func (gongstruct *GongStruct) SetName(name string) {
	gongstruct.Name = name
}

// Stage puts gongtimefield to the model stage
func (gongtimefield *GongTimeField) Stage(stage *Stage) *GongTimeField {
	__gong__stage(stage.GongTimeFields, stage.GongTimeField_stagedOrder, stage.GongTimeField_orderStaged, &stage.GongTimeFieldOrder, stage.GongTimeFields_mapString, gongtimefield, gongtimefield.Name)
	return gongtimefield
}

// StagePreserveOrder puts gongtimefield to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GongTimeFieldOrder
// - update stage.GongTimeFieldOrder accordingly
func (gongtimefield *GongTimeField) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.GongTimeFields, stage.GongTimeField_stagedOrder, stage.GongTimeField_orderStaged, &stage.GongTimeFieldOrder, stage.GongTimeFields_mapString, gongtimefield, order, gongtimefield.Name)
}

// Unstage removes gongtimefield off the model stage
func (gongtimefield *GongTimeField) Unstage(stage *Stage) *GongTimeField {
	__gong__unstage(stage.GongTimeFields, stage.GongTimeFields_mapString, gongtimefield, gongtimefield.Name)
	return gongtimefield
}

// UnstageVoid removes gongtimefield off the model stage
func (gongtimefield *GongTimeField) UnstageVoid(stage *Stage) {
	gongtimefield.Unstage(stage)
}

func (gongtimefield *GongTimeField) StageVoid(stage *Stage) {
	gongtimefield.Stage(stage)
}

// for satisfaction of GongStruct interface
func (gongtimefield *GongTimeField) GetName() (res string) {
	return gongtimefield.Name
}

// for satisfaction of GongStruct interface
func (gongtimefield *GongTimeField) SetName(name string) {
	gongtimefield.Name = name
}

// Stage puts metareference to the model stage
func (metareference *MetaReference) Stage(stage *Stage) *MetaReference {
	__gong__stage(stage.MetaReferences, stage.MetaReference_stagedOrder, stage.MetaReference_orderStaged, &stage.MetaReferenceOrder, stage.MetaReferences_mapString, metareference, metareference.Name)
	return metareference
}

// StagePreserveOrder puts metareference to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MetaReferenceOrder
// - update stage.MetaReferenceOrder accordingly
func (metareference *MetaReference) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MetaReferences, stage.MetaReference_stagedOrder, stage.MetaReference_orderStaged, &stage.MetaReferenceOrder, stage.MetaReferences_mapString, metareference, order, metareference.Name)
}

// Unstage removes metareference off the model stage
func (metareference *MetaReference) Unstage(stage *Stage) *MetaReference {
	__gong__unstage(stage.MetaReferences, stage.MetaReferences_mapString, metareference, metareference.Name)
	return metareference
}

// UnstageVoid removes metareference off the model stage
func (metareference *MetaReference) UnstageVoid(stage *Stage) {
	metareference.Unstage(stage)
}

func (metareference *MetaReference) StageVoid(stage *Stage) {
	metareference.Stage(stage)
}

// for satisfaction of GongStruct interface
func (metareference *MetaReference) GetName() (res string) {
	return metareference.Name
}

// for satisfaction of GongStruct interface
func (metareference *MetaReference) SetName(name string) {
	metareference.Name = name
}

// Stage puts modelpkg to the model stage
func (modelpkg *ModelPkg) Stage(stage *Stage) *ModelPkg {
	__gong__stage(stage.ModelPkgs, stage.ModelPkg_stagedOrder, stage.ModelPkg_orderStaged, &stage.ModelPkgOrder, stage.ModelPkgs_mapString, modelpkg, modelpkg.Name)
	return modelpkg
}

// StagePreserveOrder puts modelpkg to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ModelPkgOrder
// - update stage.ModelPkgOrder accordingly
func (modelpkg *ModelPkg) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ModelPkgs, stage.ModelPkg_stagedOrder, stage.ModelPkg_orderStaged, &stage.ModelPkgOrder, stage.ModelPkgs_mapString, modelpkg, order, modelpkg.Name)
}

// Unstage removes modelpkg off the model stage
func (modelpkg *ModelPkg) Unstage(stage *Stage) *ModelPkg {
	__gong__unstage(stage.ModelPkgs, stage.ModelPkgs_mapString, modelpkg, modelpkg.Name)
	return modelpkg
}

// UnstageVoid removes modelpkg off the model stage
func (modelpkg *ModelPkg) UnstageVoid(stage *Stage) {
	modelpkg.Unstage(stage)
}

func (modelpkg *ModelPkg) StageVoid(stage *Stage) {
	modelpkg.Stage(stage)
}

// for satisfaction of GongStruct interface
func (modelpkg *ModelPkg) GetName() (res string) {
	return modelpkg.Name
}

// for satisfaction of GongStruct interface
func (modelpkg *ModelPkg) SetName(name string) {
	modelpkg.Name = name
}

// Stage puts pointertogongstructfield to the model stage
func (pointertogongstructfield *PointerToGongStructField) Stage(stage *Stage) *PointerToGongStructField {
	__gong__stage(stage.PointerToGongStructFields, stage.PointerToGongStructField_stagedOrder, stage.PointerToGongStructField_orderStaged, &stage.PointerToGongStructFieldOrder, stage.PointerToGongStructFields_mapString, pointertogongstructfield, pointertogongstructfield.Name)
	return pointertogongstructfield
}

// StagePreserveOrder puts pointertogongstructfield to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PointerToGongStructFieldOrder
// - update stage.PointerToGongStructFieldOrder accordingly
func (pointertogongstructfield *PointerToGongStructField) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.PointerToGongStructFields, stage.PointerToGongStructField_stagedOrder, stage.PointerToGongStructField_orderStaged, &stage.PointerToGongStructFieldOrder, stage.PointerToGongStructFields_mapString, pointertogongstructfield, order, pointertogongstructfield.Name)
}

// Unstage removes pointertogongstructfield off the model stage
func (pointertogongstructfield *PointerToGongStructField) Unstage(stage *Stage) *PointerToGongStructField {
	__gong__unstage(stage.PointerToGongStructFields, stage.PointerToGongStructFields_mapString, pointertogongstructfield, pointertogongstructfield.Name)
	return pointertogongstructfield
}

// UnstageVoid removes pointertogongstructfield off the model stage
func (pointertogongstructfield *PointerToGongStructField) UnstageVoid(stage *Stage) {
	pointertogongstructfield.Unstage(stage)
}

func (pointertogongstructfield *PointerToGongStructField) StageVoid(stage *Stage) {
	pointertogongstructfield.Stage(stage)
}

// for satisfaction of GongStruct interface
func (pointertogongstructfield *PointerToGongStructField) GetName() (res string) {
	return pointertogongstructfield.Name
}

// for satisfaction of GongStruct interface
func (pointertogongstructfield *PointerToGongStructField) SetName(name string) {
	pointertogongstructfield.Name = name
}

// Stage puts sliceofpointertogongstructfield to the model stage
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) Stage(stage *Stage) *SliceOfPointerToGongStructField {
	__gong__stage(stage.SliceOfPointerToGongStructFields, stage.SliceOfPointerToGongStructField_stagedOrder, stage.SliceOfPointerToGongStructField_orderStaged, &stage.SliceOfPointerToGongStructFieldOrder, stage.SliceOfPointerToGongStructFields_mapString, sliceofpointertogongstructfield, sliceofpointertogongstructfield.Name)
	return sliceofpointertogongstructfield
}

// StagePreserveOrder puts sliceofpointertogongstructfield to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SliceOfPointerToGongStructFieldOrder
// - update stage.SliceOfPointerToGongStructFieldOrder accordingly
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.SliceOfPointerToGongStructFields, stage.SliceOfPointerToGongStructField_stagedOrder, stage.SliceOfPointerToGongStructField_orderStaged, &stage.SliceOfPointerToGongStructFieldOrder, stage.SliceOfPointerToGongStructFields_mapString, sliceofpointertogongstructfield, order, sliceofpointertogongstructfield.Name)
}

// Unstage removes sliceofpointertogongstructfield off the model stage
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) Unstage(stage *Stage) *SliceOfPointerToGongStructField {
	__gong__unstage(stage.SliceOfPointerToGongStructFields, stage.SliceOfPointerToGongStructFields_mapString, sliceofpointertogongstructfield, sliceofpointertogongstructfield.Name)
	return sliceofpointertogongstructfield
}

// UnstageVoid removes sliceofpointertogongstructfield off the model stage
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) UnstageVoid(stage *Stage) {
	sliceofpointertogongstructfield.Unstage(stage)
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) StageVoid(stage *Stage) {
	sliceofpointertogongstructfield.Stage(stage)
}

// for satisfaction of GongStruct interface
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GetName() (res string) {
	return sliceofpointertogongstructfield.Name
}

// for satisfaction of GongStruct interface
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) SetName(name string) {
	sliceofpointertogongstructfield.Name = name
}

// Stage puts stagesetfield to the model stage
func (stagesetfield *StageSetField) Stage(stage *Stage) *StageSetField {
	__gong__stage(stage.StageSetFields, stage.StageSetField_stagedOrder, stage.StageSetField_orderStaged, &stage.StageSetFieldOrder, stage.StageSetFields_mapString, stagesetfield, stagesetfield.Name)
	return stagesetfield
}

// StagePreserveOrder puts stagesetfield to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StageSetFieldOrder
// - update stage.StageSetFieldOrder accordingly
func (stagesetfield *StageSetField) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.StageSetFields, stage.StageSetField_stagedOrder, stage.StageSetField_orderStaged, &stage.StageSetFieldOrder, stage.StageSetFields_mapString, stagesetfield, order, stagesetfield.Name)
}

// Unstage removes stagesetfield off the model stage
func (stagesetfield *StageSetField) Unstage(stage *Stage) *StageSetField {
	__gong__unstage(stage.StageSetFields, stage.StageSetFields_mapString, stagesetfield, stagesetfield.Name)
	return stagesetfield
}

// UnstageVoid removes stagesetfield off the model stage
func (stagesetfield *StageSetField) UnstageVoid(stage *Stage) {
	stagesetfield.Unstage(stage)
}

func (stagesetfield *StageSetField) StageVoid(stage *Stage) {
	stagesetfield.Stage(stage)
}

// for satisfaction of GongStruct interface
func (stagesetfield *StageSetField) GetName() (res string) {
	return stagesetfield.Name
}

// for satisfaction of GongStruct interface
func (stagesetfield *StageSetField) SetName(name string) {
	stagesetfield.Name = name
}

// Stage puts stagesetmodel to the model stage
func (stagesetmodel *StageSetModel) Stage(stage *Stage) *StageSetModel {
	__gong__stage(stage.StageSetModels, stage.StageSetModel_stagedOrder, stage.StageSetModel_orderStaged, &stage.StageSetModelOrder, stage.StageSetModels_mapString, stagesetmodel, stagesetmodel.Name)
	return stagesetmodel
}

// StagePreserveOrder puts stagesetmodel to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StageSetModelOrder
// - update stage.StageSetModelOrder accordingly
func (stagesetmodel *StageSetModel) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.StageSetModels, stage.StageSetModel_stagedOrder, stage.StageSetModel_orderStaged, &stage.StageSetModelOrder, stage.StageSetModels_mapString, stagesetmodel, order, stagesetmodel.Name)
}

// Unstage removes stagesetmodel off the model stage
func (stagesetmodel *StageSetModel) Unstage(stage *Stage) *StageSetModel {
	__gong__unstage(stage.StageSetModels, stage.StageSetModels_mapString, stagesetmodel, stagesetmodel.Name)
	return stagesetmodel
}

// UnstageVoid removes stagesetmodel off the model stage
func (stagesetmodel *StageSetModel) UnstageVoid(stage *Stage) {
	stagesetmodel.Unstage(stage)
}

func (stagesetmodel *StageSetModel) StageVoid(stage *Stage) {
	stagesetmodel.Stage(stage)
}

// for satisfaction of GongStruct interface
func (stagesetmodel *StageSetModel) GetName() (res string) {
	return stagesetmodel.Name
}

// for satisfaction of GongStruct interface
func (stagesetmodel *StageSetModel) SetName(name string) {
	stagesetmodel.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.GongBasicFields, &stage.GongBasicFields_mapString, &stage.GongBasicField_stagedOrder, &stage.GongBasicFieldOrder)

	__gong__resetStageType(&stage.GongEnums, &stage.GongEnums_mapString, &stage.GongEnum_stagedOrder, &stage.GongEnumOrder)

	__gong__resetStageType(&stage.GongEnumValues, &stage.GongEnumValues_mapString, &stage.GongEnumValue_stagedOrder, &stage.GongEnumValueOrder)

	__gong__resetStageType(&stage.GongLinks, &stage.GongLinks_mapString, &stage.GongLink_stagedOrder, &stage.GongLinkOrder)

	__gong__resetStageType(&stage.GongNotes, &stage.GongNotes_mapString, &stage.GongNote_stagedOrder, &stage.GongNoteOrder)

	__gong__resetStageType(&stage.GongStructs, &stage.GongStructs_mapString, &stage.GongStruct_stagedOrder, &stage.GongStructOrder)

	__gong__resetStageType(&stage.GongTimeFields, &stage.GongTimeFields_mapString, &stage.GongTimeField_stagedOrder, &stage.GongTimeFieldOrder)

	__gong__resetStageType(&stage.MetaReferences, &stage.MetaReferences_mapString, &stage.MetaReference_stagedOrder, &stage.MetaReferenceOrder)

	__gong__resetStageType(&stage.ModelPkgs, &stage.ModelPkgs_mapString, &stage.ModelPkg_stagedOrder, &stage.ModelPkgOrder)

	__gong__resetStageType(&stage.PointerToGongStructFields, &stage.PointerToGongStructFields_mapString, &stage.PointerToGongStructField_stagedOrder, &stage.PointerToGongStructFieldOrder)

	__gong__resetStageType(&stage.SliceOfPointerToGongStructFields, &stage.SliceOfPointerToGongStructFields_mapString, &stage.SliceOfPointerToGongStructField_stagedOrder, &stage.SliceOfPointerToGongStructFieldOrder)

	__gong__resetStageType(&stage.StageSetFields, &stage.StageSetFields_mapString, &stage.StageSetField_stagedOrder, &stage.StageSetFieldOrder)

	__gong__resetStageType(&stage.StageSetModels, &stage.StageSetModels_mapString, &stage.StageSetModel_stagedOrder, &stage.StageSetModelOrder)

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
	case *GongBasicField:
		return any(stage.GongBasicFields_mapString).(map[string]Type)
	case *GongEnum:
		return any(stage.GongEnums_mapString).(map[string]Type)
	case *GongEnumValue:
		return any(stage.GongEnumValues_mapString).(map[string]Type)
	case *GongLink:
		return any(stage.GongLinks_mapString).(map[string]Type)
	case *GongNote:
		return any(stage.GongNotes_mapString).(map[string]Type)
	case *GongStruct:
		return any(stage.GongStructs_mapString).(map[string]Type)
	case *GongTimeField:
		return any(stage.GongTimeFields_mapString).(map[string]Type)
	case *MetaReference:
		return any(stage.MetaReferences_mapString).(map[string]Type)
	case *ModelPkg:
		return any(stage.ModelPkgs_mapString).(map[string]Type)
	case *PointerToGongStructField:
		return any(stage.PointerToGongStructFields_mapString).(map[string]Type)
	case *SliceOfPointerToGongStructField:
		return any(stage.SliceOfPointerToGongStructFields_mapString).(map[string]Type)
	case *StageSetField:
		return any(stage.StageSetFields_mapString).(map[string]Type)
	case *StageSetModel:
		return any(stage.StageSetModels_mapString).(map[string]Type)
	default:
		return nil
	}
}

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *GongBasicField:
		return any(&stage.GongBasicFields).(*map[Type]struct{})
	case *GongEnum:
		return any(&stage.GongEnums).(*map[Type]struct{})
	case *GongEnumValue:
		return any(&stage.GongEnumValues).(*map[Type]struct{})
	case *GongLink:
		return any(&stage.GongLinks).(*map[Type]struct{})
	case *GongNote:
		return any(&stage.GongNotes).(*map[Type]struct{})
	case *GongStruct:
		return any(&stage.GongStructs).(*map[Type]struct{})
	case *GongTimeField:
		return any(&stage.GongTimeFields).(*map[Type]struct{})
	case *MetaReference:
		return any(&stage.MetaReferences).(*map[Type]struct{})
	case *ModelPkg:
		return any(&stage.ModelPkgs).(*map[Type]struct{})
	case *PointerToGongStructField:
		return any(&stage.PointerToGongStructFields).(*map[Type]struct{})
	case *SliceOfPointerToGongStructField:
		return any(&stage.SliceOfPointerToGongStructFields).(*map[Type]struct{})
	case *StageSetField:
		return any(&stage.StageSetFields).(*map[Type]struct{})
	case *StageSetModel:
		return any(&stage.StageSetModels).(*map[Type]struct{})
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
	case GongBasicField:
		return any(&GongBasicField{
			GongEnum: &GongEnum{Name: "GongEnum"},
		}).(*Type)
	case GongEnum:
		return any(&GongEnum{
			GongEnumValues: []*GongEnumValue{{Name: "GongEnumValues"}},
		}).(*Type)
	case GongNote:
		return any(&GongNote{
			Links: []*GongLink{{Name: "Links"}},
		}).(*Type)
	case GongStruct:
		return any(&GongStruct{
			GongBasicFields: []*GongBasicField{{Name: "GongBasicFields"}},
			GongTimeFields: []*GongTimeField{{Name: "GongTimeFields"}},
			PointerToGongStructFields: []*PointerToGongStructField{{Name: "PointerToGongStructFields"}},
			SliceOfPointerToGongStructFields: []*SliceOfPointerToGongStructField{{Name: "SliceOfPointerToGongStructFields"}},
		}).(*Type)
	case PointerToGongStructField:
		return any(&PointerToGongStructField{
			GongStruct: &GongStruct{Name: "GongStruct"},
		}).(*Type)
	case SliceOfPointerToGongStructField:
		return any(&SliceOfPointerToGongStructField{
			GongStruct: &GongStruct{Name: "GongStruct"},
		}).(*Type)
	case StageSetModel:
		return any(&StageSetModel{
			Fields: []*StageSetField{{Name: "Fields"}},
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
	// reverse maps of direct associations of GongBasicField
	case GongBasicField:
		switch fieldname {
		// insertion point for per direct association field
		case "GongEnum":
			res := make(map[*GongEnum][]*GongBasicField)
			for gongbasicfield := range stage.GongBasicFields {
				if gongbasicfield.GongEnum != nil {
					gongenum_ := gongbasicfield.GongEnum
					var gongbasicfields []*GongBasicField
					_, ok := res[gongenum_]
					if ok {
						gongbasicfields = res[gongenum_]
					} else {
						gongbasicfields = make([]*GongBasicField, 0)
					}
					gongbasicfields = append(gongbasicfields, gongbasicfield)
					res[gongenum_] = gongbasicfields
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GongEnum
	case GongEnum:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongEnumValue
	case GongEnumValue:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongLink
	case GongLink:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongNote
	case GongNote:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongStruct
	case GongStruct:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongTimeField
	case GongTimeField:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MetaReference
	case MetaReference:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ModelPkg
	case ModelPkg:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PointerToGongStructField
	case PointerToGongStructField:
		switch fieldname {
		// insertion point for per direct association field
		case "GongStruct":
			res := make(map[*GongStruct][]*PointerToGongStructField)
			for pointertogongstructfield := range stage.PointerToGongStructFields {
				if pointertogongstructfield.GongStruct != nil {
					gongstruct_ := pointertogongstructfield.GongStruct
					var pointertogongstructfields []*PointerToGongStructField
					_, ok := res[gongstruct_]
					if ok {
						pointertogongstructfields = res[gongstruct_]
					} else {
						pointertogongstructfields = make([]*PointerToGongStructField, 0)
					}
					pointertogongstructfields = append(pointertogongstructfields, pointertogongstructfield)
					res[gongstruct_] = pointertogongstructfields
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SliceOfPointerToGongStructField
	case SliceOfPointerToGongStructField:
		switch fieldname {
		// insertion point for per direct association field
		case "GongStruct":
			res := make(map[*GongStruct][]*SliceOfPointerToGongStructField)
			for sliceofpointertogongstructfield := range stage.SliceOfPointerToGongStructFields {
				if sliceofpointertogongstructfield.GongStruct != nil {
					gongstruct_ := sliceofpointertogongstructfield.GongStruct
					var sliceofpointertogongstructfields []*SliceOfPointerToGongStructField
					_, ok := res[gongstruct_]
					if ok {
						sliceofpointertogongstructfields = res[gongstruct_]
					} else {
						sliceofpointertogongstructfields = make([]*SliceOfPointerToGongStructField, 0)
					}
					sliceofpointertogongstructfields = append(sliceofpointertogongstructfields, sliceofpointertogongstructfield)
					res[gongstruct_] = sliceofpointertogongstructfields
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of StageSetField
	case StageSetField:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StageSetModel
	case StageSetModel:
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
	// reverse maps of direct associations of GongBasicField
	case GongBasicField:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongEnum
	case GongEnum:
		switch fieldname {
		// insertion point for per direct association field
		case "GongEnumValues":
			res := make(map[*GongEnumValue][]*GongEnum)
			for gongenum := range stage.GongEnums {
				for _, gongenumvalue_ := range gongenum.GongEnumValues {
					res[gongenumvalue_] = append(res[gongenumvalue_], gongenum)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GongEnumValue
	case GongEnumValue:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongLink
	case GongLink:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongNote
	case GongNote:
		switch fieldname {
		// insertion point for per direct association field
		case "Links":
			res := make(map[*GongLink][]*GongNote)
			for gongnote := range stage.GongNotes {
				for _, gonglink_ := range gongnote.Links {
					res[gonglink_] = append(res[gonglink_], gongnote)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GongStruct
	case GongStruct:
		switch fieldname {
		// insertion point for per direct association field
		case "GongBasicFields":
			res := make(map[*GongBasicField][]*GongStruct)
			for gongstruct := range stage.GongStructs {
				for _, gongbasicfield_ := range gongstruct.GongBasicFields {
					res[gongbasicfield_] = append(res[gongbasicfield_], gongstruct)
				}
			}
			return any(res).(map[*End][]*Start)
		case "GongTimeFields":
			res := make(map[*GongTimeField][]*GongStruct)
			for gongstruct := range stage.GongStructs {
				for _, gongtimefield_ := range gongstruct.GongTimeFields {
					res[gongtimefield_] = append(res[gongtimefield_], gongstruct)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PointerToGongStructFields":
			res := make(map[*PointerToGongStructField][]*GongStruct)
			for gongstruct := range stage.GongStructs {
				for _, pointertogongstructfield_ := range gongstruct.PointerToGongStructFields {
					res[pointertogongstructfield_] = append(res[pointertogongstructfield_], gongstruct)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SliceOfPointerToGongStructFields":
			res := make(map[*SliceOfPointerToGongStructField][]*GongStruct)
			for gongstruct := range stage.GongStructs {
				for _, sliceofpointertogongstructfield_ := range gongstruct.SliceOfPointerToGongStructFields {
					res[sliceofpointertogongstructfield_] = append(res[sliceofpointertogongstructfield_], gongstruct)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GongTimeField
	case GongTimeField:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MetaReference
	case MetaReference:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ModelPkg
	case ModelPkg:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PointerToGongStructField
	case PointerToGongStructField:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SliceOfPointerToGongStructField
	case SliceOfPointerToGongStructField:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StageSetField
	case StageSetField:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StageSetModel
	case StageSetModel:
		switch fieldname {
		// insertion point for per direct association field
		case "Fields":
			res := make(map[*StageSetField][]*StageSetModel)
			for stagesetmodel := range stage.StageSetModels {
				for _, stagesetfield_ := range stagesetmodel.Fields {
					res[stagesetfield_] = append(res[stagesetfield_], stagesetmodel)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	}
	return nil
}

// GongNewInstance creates a new instance of the Gongstruct
func GongNewInstance[Type GongstructPtr]() (res Type) {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic new instance
	case *GongBasicField:
		res = any(new(GongBasicField)).(Type)
	case *GongEnum:
		res = any(new(GongEnum)).(Type)
	case *GongEnumValue:
		res = any(new(GongEnumValue)).(Type)
	case *GongLink:
		res = any(new(GongLink)).(Type)
	case *GongNote:
		res = any(new(GongNote)).(Type)
	case *GongStruct:
		res = any(new(GongStruct)).(Type)
	case *GongTimeField:
		res = any(new(GongTimeField)).(Type)
	case *MetaReference:
		res = any(new(MetaReference)).(Type)
	case *ModelPkg:
		res = any(new(ModelPkg)).(Type)
	case *PointerToGongStructField:
		res = any(new(PointerToGongStructField)).(Type)
	case *SliceOfPointerToGongStructField:
		res = any(new(SliceOfPointerToGongStructField)).(Type)
	case *StageSetField:
		res = any(new(StageSetField)).(Type)
	case *StageSetModel:
		res = any(new(StageSetModel)).(Type)
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
	case *GongBasicField:
		res = "GongBasicField"
	case *GongEnum:
		res = "GongEnum"
	case *GongEnumValue:
		res = "GongEnumValue"
	case *GongLink:
		res = "GongLink"
	case *GongNote:
		res = "GongNote"
	case *GongStruct:
		res = "GongStruct"
	case *GongTimeField:
		res = "GongTimeField"
	case *MetaReference:
		res = "MetaReference"
	case *ModelPkg:
		res = "ModelPkg"
	case *PointerToGongStructField:
		res = "PointerToGongStructField"
	case *SliceOfPointerToGongStructField:
		res = "SliceOfPointerToGongStructField"
	case *StageSetField:
		res = "StageSetField"
	case *StageSetModel:
		res = "StageSetModel"
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
	case *GongBasicField:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongStruct"
		rf.Fieldname = "GongBasicFields"
		res = append(res, rf)
	case *GongEnum:
		var rf ReverseField
		_ = rf
	case *GongEnumValue:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongEnum"
		rf.Fieldname = "GongEnumValues"
		res = append(res, rf)
	case *GongLink:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongNote"
		rf.Fieldname = "Links"
		res = append(res, rf)
	case *GongNote:
		var rf ReverseField
		_ = rf
	case *GongStruct:
		var rf ReverseField
		_ = rf
	case *GongTimeField:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongStruct"
		rf.Fieldname = "GongTimeFields"
		res = append(res, rf)
	case *MetaReference:
		var rf ReverseField
		_ = rf
	case *ModelPkg:
		var rf ReverseField
		_ = rf
	case *PointerToGongStructField:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongStruct"
		rf.Fieldname = "PointerToGongStructFields"
		res = append(res, rf)
	case *SliceOfPointerToGongStructField:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongStruct"
		rf.Fieldname = "SliceOfPointerToGongStructFields"
		res = append(res, rf)
	case *StageSetField:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "StageSetModel"
		rf.Fieldname = "Fields"
		res = append(res, rf)
	case *StageSetModel:
		var rf ReverseField
		_ = rf
	}
	return
}

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
}

// insertion point for get fields header method
func (gongbasicfield *GongBasicField) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BasicKindName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "GongEnum",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "GongEnum",
		},
		{
			Name:               "DeclaredType",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "CompositeStructName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsAccordionStart",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "AccordionName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsAccordionEnd",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "Index",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "IsTextArea",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsBespokeWidth",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "BespokeWidth",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "IsBespokeHeight",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "BespokeHeight",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (gongenum *GongEnum) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Type",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "GongEnumType",
		},
		{
			Name:                 "GongEnumValues",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GongEnumValue",
		},
	}
	return
}

func (gongenumvalue *GongEnumValue) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (gonglink *GongLink) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Recv",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ImportPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (gongnote *GongNote) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Body",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BodyHTML",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Links",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GongLink",
		},
	}
	return
}

func (gongstruct *GongStruct) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "GongBasicFields",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GongBasicField",
		},
		{
			Name:                 "GongTimeFields",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GongTimeField",
		},
		{
			Name:                 "PointerToGongStructFields",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "PointerToGongStructField",
		},
		{
			Name:                 "SliceOfPointerToGongStructFields",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "SliceOfPointerToGongStructField",
		},
		{
			Name:               "HasOnAfterUpdateSignature",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsIgnoredForFront",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsOmittedForMarshalling",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (gongtimefield *GongTimeField) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Index",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "CompositeStructName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsAccordionStart",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "AccordionName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsAccordionEnd",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "BespokeTimeFormat",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "TimeFormOnly",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (metareference *MetaReference) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (modelpkg *ModelPkg) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "PkgGoName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "PkgPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "PathToGoSubDirectory",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "OrmPkgGenPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DbOrmPkgGenPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DbLiteOrmPkgGenPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DbPkgGenPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ControllersPkgGenPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FullstackPkgGenPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StackPkgGenPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Level1StackPkgGenPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StaticPkgGenPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ProbePkgGenPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NgWorkspacePath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NgWorkspaceName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NgDataLibrarySourceCodeDirectory",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NgSpecificLibrarySourceCodeDirectory",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MaterialLibDatamodelTargetPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (pointertogongstructfield *PointerToGongStructField) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "GongStruct",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "GongStruct",
		},
		{
			Name:               "Index",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "CompositeStructName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsAccordionStart",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "AccordionName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsAccordionEnd",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsType",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "GongStruct",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "GongStruct",
		},
		{
			Name:               "Index",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "CompositeStructName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsAccordionStart",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "AccordionName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsAccordionEnd",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (stagesetfield *StageSetField) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "PackageName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "PackagePath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsLocal",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ImportAlias",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (stagesetmodel *StageSetModel) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Fields",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "StageSetField",
		},
		{
			Name:               "IsManual",
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
func (gongbasicfield *GongBasicField) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gongbasicfield.Name
	case "BasicKindName":
		res.valueString = gongbasicfield.BasicKindName
	case "GongEnum":
		res.GongFieldValueType = GongFieldValueTypePointer
		if gongbasicfield.GongEnum != nil {
			res.valueString = gongbasicfield.GongEnum.Name
			res.ids = gongbasicfield.GongEnum.GongGetUUID(stage)
		}
	case "DeclaredType":
		res.valueString = gongbasicfield.DeclaredType
	case "CompositeStructName":
		res.valueString = gongbasicfield.CompositeStructName
	case "IsAccordionStart":
		res.valueString = fmt.Sprintf("%t", gongbasicfield.IsAccordionStart)
		res.valueBool = gongbasicfield.IsAccordionStart
		res.GongFieldValueType = GongFieldValueTypeBool
	case "AccordionName":
		res.valueString = gongbasicfield.AccordionName
	case "IsAccordionEnd":
		res.valueString = fmt.Sprintf("%t", gongbasicfield.IsAccordionEnd)
		res.valueBool = gongbasicfield.IsAccordionEnd
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Index":
		res.valueString = fmt.Sprintf("%d", gongbasicfield.Index)
		res.valueInt = gongbasicfield.Index
		res.GongFieldValueType = GongFieldValueTypeInt
	case "IsTextArea":
		res.valueString = fmt.Sprintf("%t", gongbasicfield.IsTextArea)
		res.valueBool = gongbasicfield.IsTextArea
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsBespokeWidth":
		res.valueString = fmt.Sprintf("%t", gongbasicfield.IsBespokeWidth)
		res.valueBool = gongbasicfield.IsBespokeWidth
		res.GongFieldValueType = GongFieldValueTypeBool
	case "BespokeWidth":
		res.valueString = fmt.Sprintf("%d", gongbasicfield.BespokeWidth)
		res.valueInt = gongbasicfield.BespokeWidth
		res.GongFieldValueType = GongFieldValueTypeInt
	case "IsBespokeHeight":
		res.valueString = fmt.Sprintf("%t", gongbasicfield.IsBespokeHeight)
		res.valueBool = gongbasicfield.IsBespokeHeight
		res.GongFieldValueType = GongFieldValueTypeBool
	case "BespokeHeight":
		res.valueString = fmt.Sprintf("%d", gongbasicfield.BespokeHeight)
		res.valueInt = gongbasicfield.BespokeHeight
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (gongenum *GongEnum) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gongenum.Name
	case "Type":
		enum := gongenum.Type
		res.valueString = enum.ToCodeString()
	case "GongEnumValues":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gongenum.GongEnumValues {
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

func (gongenumvalue *GongEnumValue) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gongenumvalue.Name
	case "Value":
		res.valueString = gongenumvalue.Value
	}
	return
}

func (gonglink *GongLink) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gonglink.Name
	case "Recv":
		res.valueString = gonglink.Recv
	case "ImportPath":
		res.valueString = gonglink.ImportPath
	}
	return
}

func (gongnote *GongNote) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gongnote.Name
	case "Body":
		res.valueString = gongnote.Body
	case "BodyHTML":
		res.valueString = gongnote.BodyHTML
	case "Links":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gongnote.Links {
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

func (gongstruct *GongStruct) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gongstruct.Name
	case "GongBasicFields":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gongstruct.GongBasicFields {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "GongTimeFields":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gongstruct.GongTimeFields {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PointerToGongStructFields":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gongstruct.PointerToGongStructFields {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SliceOfPointerToGongStructFields":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gongstruct.SliceOfPointerToGongStructFields {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "HasOnAfterUpdateSignature":
		res.valueString = fmt.Sprintf("%t", gongstruct.HasOnAfterUpdateSignature)
		res.valueBool = gongstruct.HasOnAfterUpdateSignature
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsIgnoredForFront":
		res.valueString = fmt.Sprintf("%t", gongstruct.IsIgnoredForFront)
		res.valueBool = gongstruct.IsIgnoredForFront
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsOmittedForMarshalling":
		res.valueString = fmt.Sprintf("%t", gongstruct.IsOmittedForMarshalling)
		res.valueBool = gongstruct.IsOmittedForMarshalling
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (gongtimefield *GongTimeField) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gongtimefield.Name
	case "Index":
		res.valueString = fmt.Sprintf("%d", gongtimefield.Index)
		res.valueInt = gongtimefield.Index
		res.GongFieldValueType = GongFieldValueTypeInt
	case "CompositeStructName":
		res.valueString = gongtimefield.CompositeStructName
	case "IsAccordionStart":
		res.valueString = fmt.Sprintf("%t", gongtimefield.IsAccordionStart)
		res.valueBool = gongtimefield.IsAccordionStart
		res.GongFieldValueType = GongFieldValueTypeBool
	case "AccordionName":
		res.valueString = gongtimefield.AccordionName
	case "IsAccordionEnd":
		res.valueString = fmt.Sprintf("%t", gongtimefield.IsAccordionEnd)
		res.valueBool = gongtimefield.IsAccordionEnd
		res.GongFieldValueType = GongFieldValueTypeBool
	case "BespokeTimeFormat":
		res.valueString = gongtimefield.BespokeTimeFormat
	case "TimeFormOnly":
		res.valueString = fmt.Sprintf("%t", gongtimefield.TimeFormOnly)
		res.valueBool = gongtimefield.TimeFormOnly
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (metareference *MetaReference) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = metareference.Name
	}
	return
}

func (modelpkg *ModelPkg) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = modelpkg.Name
	case "PkgGoName":
		res.valueString = modelpkg.PkgGoName
	case "PkgPath":
		res.valueString = modelpkg.PkgPath
	case "PathToGoSubDirectory":
		res.valueString = modelpkg.PathToGoSubDirectory
	case "OrmPkgGenPath":
		res.valueString = modelpkg.OrmPkgGenPath
	case "DbOrmPkgGenPath":
		res.valueString = modelpkg.DbOrmPkgGenPath
	case "DbLiteOrmPkgGenPath":
		res.valueString = modelpkg.DbLiteOrmPkgGenPath
	case "DbPkgGenPath":
		res.valueString = modelpkg.DbPkgGenPath
	case "ControllersPkgGenPath":
		res.valueString = modelpkg.ControllersPkgGenPath
	case "FullstackPkgGenPath":
		res.valueString = modelpkg.FullstackPkgGenPath
	case "StackPkgGenPath":
		res.valueString = modelpkg.StackPkgGenPath
	case "Level1StackPkgGenPath":
		res.valueString = modelpkg.Level1StackPkgGenPath
	case "StaticPkgGenPath":
		res.valueString = modelpkg.StaticPkgGenPath
	case "ProbePkgGenPath":
		res.valueString = modelpkg.ProbePkgGenPath
	case "NgWorkspacePath":
		res.valueString = modelpkg.NgWorkspacePath
	case "NgWorkspaceName":
		res.valueString = modelpkg.NgWorkspaceName
	case "NgDataLibrarySourceCodeDirectory":
		res.valueString = modelpkg.NgDataLibrarySourceCodeDirectory
	case "NgSpecificLibrarySourceCodeDirectory":
		res.valueString = modelpkg.NgSpecificLibrarySourceCodeDirectory
	case "MaterialLibDatamodelTargetPath":
		res.valueString = modelpkg.MaterialLibDatamodelTargetPath
	}
	return
}

func (pointertogongstructfield *PointerToGongStructField) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = pointertogongstructfield.Name
	case "GongStruct":
		res.GongFieldValueType = GongFieldValueTypePointer
		if pointertogongstructfield.GongStruct != nil {
			res.valueString = pointertogongstructfield.GongStruct.Name
			res.ids = pointertogongstructfield.GongStruct.GongGetUUID(stage)
		}
	case "Index":
		res.valueString = fmt.Sprintf("%d", pointertogongstructfield.Index)
		res.valueInt = pointertogongstructfield.Index
		res.GongFieldValueType = GongFieldValueTypeInt
	case "CompositeStructName":
		res.valueString = pointertogongstructfield.CompositeStructName
	case "IsAccordionStart":
		res.valueString = fmt.Sprintf("%t", pointertogongstructfield.IsAccordionStart)
		res.valueBool = pointertogongstructfield.IsAccordionStart
		res.GongFieldValueType = GongFieldValueTypeBool
	case "AccordionName":
		res.valueString = pointertogongstructfield.AccordionName
	case "IsAccordionEnd":
		res.valueString = fmt.Sprintf("%t", pointertogongstructfield.IsAccordionEnd)
		res.valueBool = pointertogongstructfield.IsAccordionEnd
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsType":
		res.valueString = fmt.Sprintf("%t", pointertogongstructfield.IsType)
		res.valueBool = pointertogongstructfield.IsType
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = sliceofpointertogongstructfield.Name
	case "GongStruct":
		res.GongFieldValueType = GongFieldValueTypePointer
		if sliceofpointertogongstructfield.GongStruct != nil {
			res.valueString = sliceofpointertogongstructfield.GongStruct.Name
			res.ids = sliceofpointertogongstructfield.GongStruct.GongGetUUID(stage)
		}
	case "Index":
		res.valueString = fmt.Sprintf("%d", sliceofpointertogongstructfield.Index)
		res.valueInt = sliceofpointertogongstructfield.Index
		res.GongFieldValueType = GongFieldValueTypeInt
	case "CompositeStructName":
		res.valueString = sliceofpointertogongstructfield.CompositeStructName
	case "IsAccordionStart":
		res.valueString = fmt.Sprintf("%t", sliceofpointertogongstructfield.IsAccordionStart)
		res.valueBool = sliceofpointertogongstructfield.IsAccordionStart
		res.GongFieldValueType = GongFieldValueTypeBool
	case "AccordionName":
		res.valueString = sliceofpointertogongstructfield.AccordionName
	case "IsAccordionEnd":
		res.valueString = fmt.Sprintf("%t", sliceofpointertogongstructfield.IsAccordionEnd)
		res.valueBool = sliceofpointertogongstructfield.IsAccordionEnd
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (stagesetfield *StageSetField) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = stagesetfield.Name
	case "PackageName":
		res.valueString = stagesetfield.PackageName
	case "PackagePath":
		res.valueString = stagesetfield.PackagePath
	case "IsLocal":
		res.valueString = fmt.Sprintf("%t", stagesetfield.IsLocal)
		res.valueBool = stagesetfield.IsLocal
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ImportAlias":
		res.valueString = stagesetfield.ImportAlias
	}
	return
}

func (stagesetmodel *StageSetModel) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = stagesetmodel.Name
	case "Fields":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range stagesetmodel.Fields {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsManual":
		res.valueString = fmt.Sprintf("%t", stagesetmodel.IsManual)
		res.valueBool = stagesetmodel.IsManual
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
func (gongbasicfield *GongBasicField) GongGetGongstructName() string {
	return "GongBasicField"
}

func (gongenum *GongEnum) GongGetGongstructName() string {
	return "GongEnum"
}

func (gongenumvalue *GongEnumValue) GongGetGongstructName() string {
	return "GongEnumValue"
}

func (gonglink *GongLink) GongGetGongstructName() string {
	return "GongLink"
}

func (gongnote *GongNote) GongGetGongstructName() string {
	return "GongNote"
}

func (gongstruct *GongStruct) GongGetGongstructName() string {
	return "GongStruct"
}

func (gongtimefield *GongTimeField) GongGetGongstructName() string {
	return "GongTimeField"
}

func (metareference *MetaReference) GongGetGongstructName() string {
	return "MetaReference"
}

func (modelpkg *ModelPkg) GongGetGongstructName() string {
	return "ModelPkg"
}

func (pointertogongstructfield *PointerToGongStructField) GongGetGongstructName() string {
	return "PointerToGongStructField"
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongGetGongstructName() string {
	return "SliceOfPointerToGongStructField"
}

func (stagesetfield *StageSetField) GongGetGongstructName() string {
	return "StageSetField"
}

func (stagesetmodel *StageSetModel) GongGetGongstructName() string {
	return "StageSetModel"
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
	__gong__rebuildMapString(stage.GongBasicFields, &stage.GongBasicFields_mapString)

	__gong__rebuildMapString(stage.GongEnums, &stage.GongEnums_mapString)

	__gong__rebuildMapString(stage.GongEnumValues, &stage.GongEnumValues_mapString)

	__gong__rebuildMapString(stage.GongLinks, &stage.GongLinks_mapString)

	__gong__rebuildMapString(stage.GongNotes, &stage.GongNotes_mapString)

	__gong__rebuildMapString(stage.GongStructs, &stage.GongStructs_mapString)

	__gong__rebuildMapString(stage.GongTimeFields, &stage.GongTimeFields_mapString)

	__gong__rebuildMapString(stage.MetaReferences, &stage.MetaReferences_mapString)

	__gong__rebuildMapString(stage.ModelPkgs, &stage.ModelPkgs_mapString)

	__gong__rebuildMapString(stage.PointerToGongStructFields, &stage.PointerToGongStructFields_mapString)

	__gong__rebuildMapString(stage.SliceOfPointerToGongStructFields, &stage.SliceOfPointerToGongStructFields_mapString)

	__gong__rebuildMapString(stage.StageSetFields, &stage.StageSetFields_mapString)

	__gong__rebuildMapString(stage.StageSetModels, &stage.StageSetModels_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
