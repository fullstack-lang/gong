// generated code - do not edit
package music

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
	MusicAbstracts                map[*MusicAbstract]struct{}
	MusicAbstracts_instance       map[*MusicAbstract]*MusicAbstract
	MusicAbstracts_mapString      map[string]*MusicAbstract
	MusicAbstractOrder            uint
	MusicAbstract_stagedOrder     map[*MusicAbstract]uint
	MusicAbstract_orderStaged     map[uint]*MusicAbstract
	MusicAbstracts_reference      map[*MusicAbstract]*MusicAbstract
	MusicAbstracts_referenceOrder map[*MusicAbstract]uint

	// insertion point for slice of pointers maps
	OnAfterMusicAbstractCreateCallback GongOnAfterCreateInterface[MusicAbstract]
	OnAfterMusicAbstractUpdateCallback GongOnAfterUpdateInterface[MusicAbstract]
	OnAfterMusicAbstractDeleteCallback GongOnAfterDeleteInterface[MusicAbstract]
	OnAfterMusicAbstractReadCallback   GongOnAfterReadInterface[MusicAbstract]

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
	stage.MusicAbstracts_reference = make(map[*MusicAbstract]*MusicAbstract)
	stage.MusicAbstracts_instance = make(map[*MusicAbstract]*MusicAbstract)
	stage.MusicAbstracts_referenceOrder = make(map[*MusicAbstract]uint)

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
	var maxMusicAbstractOrder uint
	var foundMusicAbstract bool
	for _, order := range stage.MusicAbstract_stagedOrder {
		if !foundMusicAbstract || order > maxMusicAbstractOrder {
			maxMusicAbstractOrder = order
			foundMusicAbstract = true
		}
	}
	if foundMusicAbstract {
		stage.MusicAbstractOrder = maxMusicAbstractOrder + 1
	} else {
		stage.MusicAbstractOrder = 0
	}

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
	case *MusicAbstract:
		tmp := __gong__getStructInstancesByOrder(stage.MusicAbstracts, stage.MusicAbstract_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *MusicAbstract implements.
			res = append(res, any(v).(T))
		}
		return res

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

func (stage *Stage) GetType() string {
	return "github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/music"
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

// GongOnAfterReadInterface callback when an instance is updated from the front
type GongOnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *Stage,
		instance *Type)
}

type OnAfterReadInterface[Type Gongstruct] = GongOnAfterReadInterface[Type]

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
	// insertion point for Commit and Checkout signatures
	CommitMusicAbstract(musicabstract *MusicAbstract)
	CheckoutMusicAbstract(musicabstract *MusicAbstract)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

type BackRepoInterface = GongBackRepoInterface

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		MusicAbstracts:           make(map[*MusicAbstract]struct{}),
		MusicAbstracts_mapString: make(map[string]*MusicAbstract),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		MusicAbstract_stagedOrder: make(map[*MusicAbstract]uint),
		MusicAbstract_orderStaged: make(map[uint]*MusicAbstract),
		MusicAbstracts_reference:  make(map[*MusicAbstract]*MusicAbstract),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"MusicAbstract": &MusicAbstractUnmarshaller{},

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
	case *MusicAbstract:
		return any(stage.MusicAbstract_orderStaged[order]).(Type)
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
	stage.Map_GongStructName_InstancesNb["MusicAbstract"] = len(stage.MusicAbstracts)
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
// Stage puts musicabstract to the model stage
func (musicabstract *MusicAbstract) Stage(stage *Stage) *MusicAbstract {
	if _, ok := stage.MusicAbstracts[musicabstract]; !ok {
		stage.MusicAbstracts[musicabstract] = struct{}{}
		stage.MusicAbstract_stagedOrder[musicabstract] = stage.MusicAbstractOrder
		stage.MusicAbstract_orderStaged[stage.MusicAbstractOrder] = musicabstract
		stage.MusicAbstractOrder++
	}
	stage.MusicAbstracts_mapString[musicabstract.Name] = musicabstract

	return musicabstract
}

// StagePreserveOrder puts musicabstract to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MusicAbstractOrder
// - update stage.MusicAbstractOrder accordingly
func (musicabstract *MusicAbstract) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.MusicAbstracts[musicabstract]; !ok {
		stage.MusicAbstracts[musicabstract] = struct{}{}

		if order > stage.MusicAbstractOrder {
			stage.MusicAbstractOrder = order
		}
		stage.MusicAbstract_stagedOrder[musicabstract] = order
		stage.MusicAbstract_orderStaged[order] = musicabstract
		stage.MusicAbstractOrder++
	}
	stage.MusicAbstracts_mapString[musicabstract.Name] = musicabstract
}

// Unstage removes musicabstract off the model stage
func (musicabstract *MusicAbstract) Unstage(stage *Stage) *MusicAbstract {
	delete(stage.MusicAbstracts, musicabstract)
	// issue1150
	// delete(stage.MusicAbstract_stagedOrder, musicabstract)
	delete(stage.MusicAbstracts_mapString, musicabstract.Name)

	return musicabstract
}

// UnstageVoid removes musicabstract off the model stage
func (musicabstract *MusicAbstract) UnstageVoid(stage *Stage) {
	delete(stage.MusicAbstracts, musicabstract)
	// issue1150
	// delete(stage.MusicAbstract_stagedOrder, musicabstract)
	delete(stage.MusicAbstracts_mapString, musicabstract.Name)
}

// commit musicabstract to the back repo (if it is already staged)
func (musicabstract *MusicAbstract) Commit(stage *Stage) *MusicAbstract {
	if _, ok := stage.MusicAbstracts[musicabstract]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMusicAbstract(musicabstract)
		}
	}
	return musicabstract
}

func (musicabstract *MusicAbstract) StageVoid(stage *Stage) {
	musicabstract.Stage(stage)
}

// Checkout musicabstract to the back repo (if it is already staged)
func (musicabstract *MusicAbstract) Checkout(stage *Stage) *MusicAbstract {
	if _, ok := stage.MusicAbstracts[musicabstract]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMusicAbstract(musicabstract)
		}
	}
	return musicabstract
}

// for satisfaction of GongStruct interface
func (musicabstract *MusicAbstract) GetName() (res string) {
	return musicabstract.Name
}

// for satisfaction of GongStruct interface
func (musicabstract *MusicAbstract) SetName(name string) {
	musicabstract.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.MusicAbstracts = make(map[*MusicAbstract]struct{})
	stage.MusicAbstracts_mapString = make(map[string]*MusicAbstract)
	stage.MusicAbstract_stagedOrder = make(map[*MusicAbstract]uint)
	stage.MusicAbstractOrder = 0

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
	GongClean(stage *Stage) (modified bool)
	GongGetFieldValue(fieldName string, stage *Stage) GongFieldValue
	GongGetGongstructName() string
	GongGetOrder(stage *Stage) uint
	GongGetReferenceIdentifier(stage *Stage) string
	GongGetIdentifier(stage *Stage) string
	GongCopy() GongstructIF
	GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) string
	GongGetReverseFieldOwner(stage *Stage, reverseField *GongReverseField) GongstructIF
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
	case *MusicAbstract:
		return any(stage.MusicAbstracts_mapString).(map[string]Type)
	default:
		return nil
	}
}

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *MusicAbstract:
		return any(&stage.MusicAbstracts).(*map[Type]struct{})
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
	// reverse maps of direct associations of MusicAbstract
	case MusicAbstract:
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
	// reverse maps of direct associations of MusicAbstract
	case MusicAbstract:
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
	case *MusicAbstract:
		res = any(new(MusicAbstract)).(Type)
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
	case *MusicAbstract:
		res = "MusicAbstract"
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
	case *MusicAbstract:
		var rf ReverseField
		_ = rf
	}
	return
}

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
}

// insertion point for get fields header method
func (musicabstract *MusicAbstract) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsChecked",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "PitchHeight",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "NbOfBeatsInTheme",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "BeatsPerSecond",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "FirstVoiceShiftX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "FirstVoiceShiftY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "PitchDifference",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Level",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ActualBeatsTemporalShift",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "IsMinor",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ThemeBinaryEncoding",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "BezierControlLengthRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "NbPitchLines",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NbBeatLines",
			GongFieldValueType: GongFieldValueTypeInt,
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
			Name:               "ScoreScale",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ShowFirstVoice",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ShowFirstVoiceShiftRight",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ShowSecondVoice",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ShowSecondVoiceShiftRight",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ShowFirstVoiceNotes",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ShowFirstVoiceNotesShiftRight",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ShowSecondVoiceNotes",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ShowSecondVoiceNotesShiftRight",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsComposerNodeExpanded",
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
func (musicabstract *MusicAbstract) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = musicabstract.Name
	case "IsChecked":
		res.valueString = fmt.Sprintf("%t", musicabstract.IsChecked)
		res.valueBool = musicabstract.IsChecked
		res.GongFieldValueType = GongFieldValueTypeBool
	case "PitchHeight":
		res.valueString = fmt.Sprintf("%f", musicabstract.PitchHeight)
		res.valueFloat = musicabstract.PitchHeight
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "NbOfBeatsInTheme":
		res.valueString = fmt.Sprintf("%d", musicabstract.NbOfBeatsInTheme)
		res.valueInt = musicabstract.NbOfBeatsInTheme
		res.GongFieldValueType = GongFieldValueTypeInt
	case "BeatsPerSecond":
		res.valueString = fmt.Sprintf("%f", musicabstract.BeatsPerSecond)
		res.valueFloat = musicabstract.BeatsPerSecond
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "FirstVoiceShiftX":
		res.valueString = fmt.Sprintf("%f", musicabstract.FirstVoiceShiftX)
		res.valueFloat = musicabstract.FirstVoiceShiftX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "FirstVoiceShiftY":
		res.valueString = fmt.Sprintf("%f", musicabstract.FirstVoiceShiftY)
		res.valueFloat = musicabstract.FirstVoiceShiftY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "PitchDifference":
		res.valueString = fmt.Sprintf("%d", musicabstract.PitchDifference)
		res.valueInt = musicabstract.PitchDifference
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Level":
		res.valueString = fmt.Sprintf("%f", musicabstract.Level)
		res.valueFloat = musicabstract.Level
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ActualBeatsTemporalShift":
		res.valueString = fmt.Sprintf("%d", musicabstract.ActualBeatsTemporalShift)
		res.valueInt = musicabstract.ActualBeatsTemporalShift
		res.GongFieldValueType = GongFieldValueTypeInt
	case "IsMinor":
		res.valueString = fmt.Sprintf("%t", musicabstract.IsMinor)
		res.valueBool = musicabstract.IsMinor
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ThemeBinaryEncoding":
		res.valueString = fmt.Sprintf("%d", musicabstract.ThemeBinaryEncoding)
		res.valueInt = musicabstract.ThemeBinaryEncoding
		res.GongFieldValueType = GongFieldValueTypeInt
	case "BezierControlLengthRatio":
		res.valueString = fmt.Sprintf("%f", musicabstract.BezierControlLengthRatio)
		res.valueFloat = musicabstract.BezierControlLengthRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "NbPitchLines":
		res.valueString = fmt.Sprintf("%d", musicabstract.NbPitchLines)
		res.valueInt = musicabstract.NbPitchLines
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NbBeatLines":
		res.valueString = fmt.Sprintf("%d", musicabstract.NbBeatLines)
		res.valueInt = musicabstract.NbBeatLines
		res.GongFieldValueType = GongFieldValueTypeInt
	case "OriginX":
		res.valueString = fmt.Sprintf("%f", musicabstract.OriginX)
		res.valueFloat = musicabstract.OriginX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "OriginY":
		res.valueString = fmt.Sprintf("%f", musicabstract.OriginY)
		res.valueFloat = musicabstract.OriginY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ScoreScale":
		res.valueString = fmt.Sprintf("%f", musicabstract.ScoreScale)
		res.valueFloat = musicabstract.ScoreScale
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ShowFirstVoice":
		res.valueString = fmt.Sprintf("%t", musicabstract.ShowFirstVoice)
		res.valueBool = musicabstract.ShowFirstVoice
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShowFirstVoiceShiftRight":
		res.valueString = fmt.Sprintf("%t", musicabstract.ShowFirstVoiceShiftRight)
		res.valueBool = musicabstract.ShowFirstVoiceShiftRight
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShowSecondVoice":
		res.valueString = fmt.Sprintf("%t", musicabstract.ShowSecondVoice)
		res.valueBool = musicabstract.ShowSecondVoice
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShowSecondVoiceShiftRight":
		res.valueString = fmt.Sprintf("%t", musicabstract.ShowSecondVoiceShiftRight)
		res.valueBool = musicabstract.ShowSecondVoiceShiftRight
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShowFirstVoiceNotes":
		res.valueString = fmt.Sprintf("%t", musicabstract.ShowFirstVoiceNotes)
		res.valueBool = musicabstract.ShowFirstVoiceNotes
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShowFirstVoiceNotesShiftRight":
		res.valueString = fmt.Sprintf("%t", musicabstract.ShowFirstVoiceNotesShiftRight)
		res.valueBool = musicabstract.ShowFirstVoiceNotesShiftRight
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShowSecondVoiceNotes":
		res.valueString = fmt.Sprintf("%t", musicabstract.ShowSecondVoiceNotes)
		res.valueBool = musicabstract.ShowSecondVoiceNotes
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShowSecondVoiceNotesShiftRight":
		res.valueString = fmt.Sprintf("%t", musicabstract.ShowSecondVoiceNotesShiftRight)
		res.valueBool = musicabstract.ShowSecondVoiceNotesShiftRight
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsComposerNodeExpanded":
		res.valueString = fmt.Sprintf("%t", musicabstract.IsComposerNodeExpanded)
		res.valueBool = musicabstract.IsComposerNodeExpanded
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
func (musicabstract *MusicAbstract) GongGetGongstructName() string {
	return "MusicAbstract"
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
	stage.MusicAbstracts_mapString = make(map[string]*MusicAbstract)
	for musicabstract := range stage.MusicAbstracts {
		stage.MusicAbstracts_mapString[musicabstract.Name] = musicabstract
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
