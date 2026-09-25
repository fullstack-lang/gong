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
	Contents                map[*Content]struct{}
	Contents_instance       map[*Content]*Content
	Contents_mapString      map[string]*Content
	ContentOrder            uint
	Content_stagedOrder     map[*Content]uint
	Content_orderStaged     map[uint]*Content
	Contents_reference      map[*Content]*Content
	Contents_referenceOrder map[*Content]uint

	// insertion point for slice of pointers maps
	OnAfterContentCreateCallback GongOnAfterCreateInterface[Content]
	OnAfterContentUpdateCallback GongOnAfterUpdateInterface[Content]
	OnAfterContentDeleteCallback GongOnAfterDeleteInterface[Content]

	JpgImages                map[*JpgImage]struct{}
	JpgImages_instance       map[*JpgImage]*JpgImage
	JpgImages_mapString      map[string]*JpgImage
	JpgImageOrder            uint
	JpgImage_stagedOrder     map[*JpgImage]uint
	JpgImage_orderStaged     map[uint]*JpgImage
	JpgImages_reference      map[*JpgImage]*JpgImage
	JpgImages_referenceOrder map[*JpgImage]uint

	// insertion point for slice of pointers maps
	OnAfterJpgImageCreateCallback GongOnAfterCreateInterface[JpgImage]
	OnAfterJpgImageUpdateCallback GongOnAfterUpdateInterface[JpgImage]
	OnAfterJpgImageDeleteCallback GongOnAfterDeleteInterface[JpgImage]

	PngImages                map[*PngImage]struct{}
	PngImages_instance       map[*PngImage]*PngImage
	PngImages_mapString      map[string]*PngImage
	PngImageOrder            uint
	PngImage_stagedOrder     map[*PngImage]uint
	PngImage_orderStaged     map[uint]*PngImage
	PngImages_reference      map[*PngImage]*PngImage
	PngImages_referenceOrder map[*PngImage]uint

	// insertion point for slice of pointers maps
	OnAfterPngImageCreateCallback GongOnAfterCreateInterface[PngImage]
	OnAfterPngImageUpdateCallback GongOnAfterUpdateInterface[PngImage]
	OnAfterPngImageDeleteCallback GongOnAfterDeleteInterface[PngImage]

	SvgImages                map[*SvgImage]struct{}
	SvgImages_instance       map[*SvgImage]*SvgImage
	SvgImages_mapString      map[string]*SvgImage
	SvgImageOrder            uint
	SvgImage_stagedOrder     map[*SvgImage]uint
	SvgImage_orderStaged     map[uint]*SvgImage
	SvgImages_reference      map[*SvgImage]*SvgImage
	SvgImages_referenceOrder map[*SvgImage]uint

	// insertion point for slice of pointers maps
	OnAfterSvgImageCreateCallback GongOnAfterCreateInterface[SvgImage]
	OnAfterSvgImageUpdateCallback GongOnAfterUpdateInterface[SvgImage]
	OnAfterSvgImageDeleteCallback GongOnAfterDeleteInterface[SvgImage]

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
	__gong__clearReferences(&stage.Contents_reference, &stage.Contents_instance, &stage.Contents_referenceOrder)

	__gong__clearReferences(&stage.JpgImages_reference, &stage.JpgImages_instance, &stage.JpgImages_referenceOrder)

	__gong__clearReferences(&stage.PngImages_reference, &stage.PngImages_instance, &stage.PngImages_referenceOrder)

	__gong__clearReferences(&stage.SvgImages_reference, &stage.SvgImages_instance, &stage.SvgImages_referenceOrder)

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
	stage.ContentOrder = __gong__recomputeOrder(stage.Content_stagedOrder)

	stage.JpgImageOrder = __gong__recomputeOrder(stage.JpgImage_stagedOrder)

	stage.PngImageOrder = __gong__recomputeOrder(stage.PngImage_stagedOrder)

	stage.SvgImageOrder = __gong__recomputeOrder(stage.SvgImage_stagedOrder)

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
	case *Content:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Contents, stage.Content_stagedOrder))
	case *JpgImage:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.JpgImages, stage.JpgImage_stagedOrder))
	case *PngImage:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.PngImages, stage.PngImage_stagedOrder))
	case *SvgImage:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.SvgImages, stage.SvgImage_stagedOrder))

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
	return "github.com/fullstack-lang/gong/lib/markdown/go/models"
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
		Contents:           make(map[*Content]struct{}),
		Contents_mapString: make(map[string]*Content),

		JpgImages:           make(map[*JpgImage]struct{}),
		JpgImages_mapString: make(map[string]*JpgImage),

		PngImages:           make(map[*PngImage]struct{}),
		PngImages_mapString: make(map[string]*PngImage),

		SvgImages:           make(map[*SvgImage]struct{}),
		SvgImages_mapString: make(map[string]*SvgImage),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		Content_stagedOrder: make(map[*Content]uint),
		Content_orderStaged: make(map[uint]*Content),
		Contents_reference:  make(map[*Content]*Content),

		JpgImage_stagedOrder: make(map[*JpgImage]uint),
		JpgImage_orderStaged: make(map[uint]*JpgImage),
		JpgImages_reference:  make(map[*JpgImage]*JpgImage),

		PngImage_stagedOrder: make(map[*PngImage]uint),
		PngImage_orderStaged: make(map[uint]*PngImage),
		PngImages_reference:  make(map[*PngImage]*PngImage),

		SvgImage_stagedOrder: make(map[*SvgImage]uint),
		SvgImage_orderStaged: make(map[uint]*SvgImage),
		SvgImages_reference:  make(map[*SvgImage]*SvgImage),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"Content": &ContentUnmarshaller{},

			"JpgImage": &JpgImageUnmarshaller{},

			"PngImage": &PngImageUnmarshaller{},

			"SvgImage": &SvgImageUnmarshaller{},

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
	case *Content:
		return any(stage.Content_orderStaged[order]).(Type)
	case *JpgImage:
		return any(stage.JpgImage_orderStaged[order]).(Type)
	case *PngImage:
		return any(stage.PngImage_orderStaged[order]).(Type)
	case *SvgImage:
		return any(stage.SvgImage_orderStaged[order]).(Type)
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
	stage.Map_GongStructName_InstancesNb["Content"] = len(stage.Contents)
	stage.Map_GongStructName_InstancesNb["JpgImage"] = len(stage.JpgImages)
	stage.Map_GongStructName_InstancesNb["PngImage"] = len(stage.PngImages)
	stage.Map_GongStructName_InstancesNb["SvgImage"] = len(stage.SvgImages)
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
// Stage puts content to the model stage
func (content *Content) Stage(stage *Stage) *Content {
	__gong__stage(stage.Contents, stage.Content_stagedOrder, stage.Content_orderStaged, &stage.ContentOrder, stage.Contents_mapString, content, content.Name)
	return content
}

// StagePreserveOrder puts content to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ContentOrder
// - update stage.ContentOrder accordingly
func (content *Content) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Contents, stage.Content_stagedOrder, stage.Content_orderStaged, &stage.ContentOrder, stage.Contents_mapString, content, order, content.Name)
}

// Unstage removes content off the model stage
func (content *Content) Unstage(stage *Stage) *Content {
	__gong__unstage(stage.Contents, stage.Contents_mapString, content, content.Name)
	return content
}

// UnstageVoid removes content off the model stage
func (content *Content) UnstageVoid(stage *Stage) {
	content.Unstage(stage)
}

func (content *Content) StageVoid(stage *Stage) {
	content.Stage(stage)
}

// for satisfaction of GongStruct interface
func (content *Content) GetName() (res string) {
	return content.Name
}

// for satisfaction of GongStruct interface
func (content *Content) SetName(name string) {
	content.Name = name
}

// Stage puts jpgimage to the model stage
func (jpgimage *JpgImage) Stage(stage *Stage) *JpgImage {
	__gong__stage(stage.JpgImages, stage.JpgImage_stagedOrder, stage.JpgImage_orderStaged, &stage.JpgImageOrder, stage.JpgImages_mapString, jpgimage, jpgimage.Name)
	return jpgimage
}

// StagePreserveOrder puts jpgimage to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.JpgImageOrder
// - update stage.JpgImageOrder accordingly
func (jpgimage *JpgImage) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.JpgImages, stage.JpgImage_stagedOrder, stage.JpgImage_orderStaged, &stage.JpgImageOrder, stage.JpgImages_mapString, jpgimage, order, jpgimage.Name)
}

// Unstage removes jpgimage off the model stage
func (jpgimage *JpgImage) Unstage(stage *Stage) *JpgImage {
	__gong__unstage(stage.JpgImages, stage.JpgImages_mapString, jpgimage, jpgimage.Name)
	return jpgimage
}

// UnstageVoid removes jpgimage off the model stage
func (jpgimage *JpgImage) UnstageVoid(stage *Stage) {
	jpgimage.Unstage(stage)
}

func (jpgimage *JpgImage) StageVoid(stage *Stage) {
	jpgimage.Stage(stage)
}

// for satisfaction of GongStruct interface
func (jpgimage *JpgImage) GetName() (res string) {
	return jpgimage.Name
}

// for satisfaction of GongStruct interface
func (jpgimage *JpgImage) SetName(name string) {
	jpgimage.Name = name
}

// Stage puts pngimage to the model stage
func (pngimage *PngImage) Stage(stage *Stage) *PngImage {
	__gong__stage(stage.PngImages, stage.PngImage_stagedOrder, stage.PngImage_orderStaged, &stage.PngImageOrder, stage.PngImages_mapString, pngimage, pngimage.Name)
	return pngimage
}

// StagePreserveOrder puts pngimage to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PngImageOrder
// - update stage.PngImageOrder accordingly
func (pngimage *PngImage) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.PngImages, stage.PngImage_stagedOrder, stage.PngImage_orderStaged, &stage.PngImageOrder, stage.PngImages_mapString, pngimage, order, pngimage.Name)
}

// Unstage removes pngimage off the model stage
func (pngimage *PngImage) Unstage(stage *Stage) *PngImage {
	__gong__unstage(stage.PngImages, stage.PngImages_mapString, pngimage, pngimage.Name)
	return pngimage
}

// UnstageVoid removes pngimage off the model stage
func (pngimage *PngImage) UnstageVoid(stage *Stage) {
	pngimage.Unstage(stage)
}

func (pngimage *PngImage) StageVoid(stage *Stage) {
	pngimage.Stage(stage)
}

// for satisfaction of GongStruct interface
func (pngimage *PngImage) GetName() (res string) {
	return pngimage.Name
}

// for satisfaction of GongStruct interface
func (pngimage *PngImage) SetName(name string) {
	pngimage.Name = name
}

// Stage puts svgimage to the model stage
func (svgimage *SvgImage) Stage(stage *Stage) *SvgImage {
	__gong__stage(stage.SvgImages, stage.SvgImage_stagedOrder, stage.SvgImage_orderStaged, &stage.SvgImageOrder, stage.SvgImages_mapString, svgimage, svgimage.Name)
	return svgimage
}

// StagePreserveOrder puts svgimage to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SvgImageOrder
// - update stage.SvgImageOrder accordingly
func (svgimage *SvgImage) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.SvgImages, stage.SvgImage_stagedOrder, stage.SvgImage_orderStaged, &stage.SvgImageOrder, stage.SvgImages_mapString, svgimage, order, svgimage.Name)
}

// Unstage removes svgimage off the model stage
func (svgimage *SvgImage) Unstage(stage *Stage) *SvgImage {
	__gong__unstage(stage.SvgImages, stage.SvgImages_mapString, svgimage, svgimage.Name)
	return svgimage
}

// UnstageVoid removes svgimage off the model stage
func (svgimage *SvgImage) UnstageVoid(stage *Stage) {
	svgimage.Unstage(stage)
}

func (svgimage *SvgImage) StageVoid(stage *Stage) {
	svgimage.Stage(stage)
}

// for satisfaction of GongStruct interface
func (svgimage *SvgImage) GetName() (res string) {
	return svgimage.Name
}

// for satisfaction of GongStruct interface
func (svgimage *SvgImage) SetName(name string) {
	svgimage.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.Contents, &stage.Contents_mapString, &stage.Content_stagedOrder, &stage.ContentOrder)

	__gong__resetStageType(&stage.JpgImages, &stage.JpgImages_mapString, &stage.JpgImage_stagedOrder, &stage.JpgImageOrder)

	__gong__resetStageType(&stage.PngImages, &stage.PngImages_mapString, &stage.PngImage_stagedOrder, &stage.PngImageOrder)

	__gong__resetStageType(&stage.SvgImages, &stage.SvgImages_mapString, &stage.SvgImage_stagedOrder, &stage.SvgImageOrder)

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
	case *Content:
		return any(stage.Contents_mapString).(map[string]Type)
	case *JpgImage:
		return any(stage.JpgImages_mapString).(map[string]Type)
	case *PngImage:
		return any(stage.PngImages_mapString).(map[string]Type)
	case *SvgImage:
		return any(stage.SvgImages_mapString).(map[string]Type)
	default:
		return nil
	}
}

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Content:
		return any(&stage.Contents).(*map[Type]struct{})
	case *JpgImage:
		return any(&stage.JpgImages).(*map[Type]struct{})
	case *PngImage:
		return any(&stage.PngImages).(*map[Type]struct{})
	case *SvgImage:
		return any(&stage.SvgImages).(*map[Type]struct{})
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
	// reverse maps of direct associations of Content
	case Content:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of JpgImage
	case JpgImage:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PngImage
	case PngImage:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SvgImage
	case SvgImage:
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
	// reverse maps of direct associations of Content
	case Content:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of JpgImage
	case JpgImage:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PngImage
	case PngImage:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SvgImage
	case SvgImage:
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
	case *Content:
		res = any(new(Content)).(Type)
	case *JpgImage:
		res = any(new(JpgImage)).(Type)
	case *PngImage:
		res = any(new(PngImage)).(Type)
	case *SvgImage:
		res = any(new(SvgImage)).(Type)
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
	case *Content:
		res = "Content"
	case *JpgImage:
		res = "JpgImage"
	case *PngImage:
		res = "PngImage"
	case *SvgImage:
		res = "SvgImage"
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
	case *Content:
		var rf ReverseField
		_ = rf
	case *JpgImage:
		var rf ReverseField
		_ = rf
	case *PngImage:
		var rf ReverseField
		_ = rf
	case *SvgImage:
		var rf ReverseField
		_ = rf
	}
	return
}

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
}

// insertion point for get fields header method
func (content *Content) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Content",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (jpgimage *JpgImage) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Base64Content",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (pngimage *PngImage) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Base64Content",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (svgimage *SvgImage) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Content",
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
func (content *Content) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = content.Name
	case "Content":
		res.valueString = content.Content
	}
	return
}

func (jpgimage *JpgImage) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = jpgimage.Name
	case "Base64Content":
		res.valueString = jpgimage.Base64Content
	}
	return
}

func (pngimage *PngImage) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = pngimage.Name
	case "Base64Content":
		res.valueString = pngimage.Base64Content
	}
	return
}

func (svgimage *SvgImage) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = svgimage.Name
	case "Content":
		res.valueString = svgimage.Content
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
func (content *Content) GongGetGongstructName() string {
	return "Content"
}

func (jpgimage *JpgImage) GongGetGongstructName() string {
	return "JpgImage"
}

func (pngimage *PngImage) GongGetGongstructName() string {
	return "PngImage"
}

func (svgimage *SvgImage) GongGetGongstructName() string {
	return "SvgImage"
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
	__gong__rebuildMapString(stage.Contents, &stage.Contents_mapString)

	__gong__rebuildMapString(stage.JpgImages, &stage.JpgImages_mapString)

	__gong__rebuildMapString(stage.PngImages, &stage.PngImages_mapString)

	__gong__rebuildMapString(stage.SvgImages, &stage.SvgImages_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
