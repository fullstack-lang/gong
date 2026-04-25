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

	process_go "github.com/fullstack-lang/gong/dsm/process/go"
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
	DiagramProcesss                map[*DiagramProcess]struct{}
	DiagramProcesss_instance       map[*DiagramProcess]*DiagramProcess
	DiagramProcesss_mapString      map[string]*DiagramProcess
	DiagramProcessOrder            uint
	DiagramProcess_stagedOrder     map[*DiagramProcess]uint
	DiagramProcess_orderStaged     map[uint]*DiagramProcess
	DiagramProcesss_reference      map[*DiagramProcess]*DiagramProcess
	DiagramProcesss_referenceOrder map[*DiagramProcess]uint

	// insertion point for slice of pointers maps
	DiagramProcess_Process_Shapes_reverseMap map[*ProcessShape]*DiagramProcess

	DiagramProcess_ProcesssWhoseNodeIsExpanded_reverseMap map[*Process]*DiagramProcess

	DiagramProcess_ProcessComposition_Shapes_reverseMap map[*ProcessCompositionShape]*DiagramProcess

	OnAfterDiagramProcessCreateCallback OnAfterCreateInterface[DiagramProcess]
	OnAfterDiagramProcessUpdateCallback OnAfterUpdateInterface[DiagramProcess]
	OnAfterDiagramProcessDeleteCallback OnAfterDeleteInterface[DiagramProcess]
	OnAfterDiagramProcessReadCallback   OnAfterReadInterface[DiagramProcess]

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

	Library_RootProcesses_reverseMap map[*Process]*Library

	Library_ProcesssWhoseNodeIsExpanded_reverseMap map[*Process]*Library

	OnAfterLibraryCreateCallback OnAfterCreateInterface[Library]
	OnAfterLibraryUpdateCallback OnAfterUpdateInterface[Library]
	OnAfterLibraryDeleteCallback OnAfterDeleteInterface[Library]
	OnAfterLibraryReadCallback   OnAfterReadInterface[Library]

	Participants                map[*Participant]struct{}
	Participants_instance       map[*Participant]*Participant
	Participants_mapString      map[string]*Participant
	ParticipantOrder            uint
	Participant_stagedOrder     map[*Participant]uint
	Participant_orderStaged     map[uint]*Participant
	Participants_reference      map[*Participant]*Participant
	Participants_referenceOrder map[*Participant]uint

	// insertion point for slice of pointers maps
	OnAfterParticipantCreateCallback OnAfterCreateInterface[Participant]
	OnAfterParticipantUpdateCallback OnAfterUpdateInterface[Participant]
	OnAfterParticipantDeleteCallback OnAfterDeleteInterface[Participant]
	OnAfterParticipantReadCallback   OnAfterReadInterface[Participant]

	Processs                map[*Process]struct{}
	Processs_instance       map[*Process]*Process
	Processs_mapString      map[string]*Process
	ProcessOrder            uint
	Process_stagedOrder     map[*Process]uint
	Process_orderStaged     map[uint]*Process
	Processs_reference      map[*Process]*Process
	Processs_referenceOrder map[*Process]uint

	// insertion point for slice of pointers maps
	Process_SubProcesses_reverseMap map[*Process]*Process

	Process_DiagramProcesss_reverseMap map[*DiagramProcess]*Process

	OnAfterProcessCreateCallback OnAfterCreateInterface[Process]
	OnAfterProcessUpdateCallback OnAfterUpdateInterface[Process]
	OnAfterProcessDeleteCallback OnAfterDeleteInterface[Process]
	OnAfterProcessReadCallback   OnAfterReadInterface[Process]

	ProcessCompositionShapes                map[*ProcessCompositionShape]struct{}
	ProcessCompositionShapes_instance       map[*ProcessCompositionShape]*ProcessCompositionShape
	ProcessCompositionShapes_mapString      map[string]*ProcessCompositionShape
	ProcessCompositionShapeOrder            uint
	ProcessCompositionShape_stagedOrder     map[*ProcessCompositionShape]uint
	ProcessCompositionShape_orderStaged     map[uint]*ProcessCompositionShape
	ProcessCompositionShapes_reference      map[*ProcessCompositionShape]*ProcessCompositionShape
	ProcessCompositionShapes_referenceOrder map[*ProcessCompositionShape]uint

	// insertion point for slice of pointers maps
	OnAfterProcessCompositionShapeCreateCallback OnAfterCreateInterface[ProcessCompositionShape]
	OnAfterProcessCompositionShapeUpdateCallback OnAfterUpdateInterface[ProcessCompositionShape]
	OnAfterProcessCompositionShapeDeleteCallback OnAfterDeleteInterface[ProcessCompositionShape]
	OnAfterProcessCompositionShapeReadCallback   OnAfterReadInterface[ProcessCompositionShape]

	ProcessShapes                map[*ProcessShape]struct{}
	ProcessShapes_instance       map[*ProcessShape]*ProcessShape
	ProcessShapes_mapString      map[string]*ProcessShape
	ProcessShapeOrder            uint
	ProcessShape_stagedOrder     map[*ProcessShape]uint
	ProcessShape_orderStaged     map[uint]*ProcessShape
	ProcessShapes_reference      map[*ProcessShape]*ProcessShape
	ProcessShapes_referenceOrder map[*ProcessShape]uint

	// insertion point for slice of pointers maps
	OnAfterProcessShapeCreateCallback OnAfterCreateInterface[ProcessShape]
	OnAfterProcessShapeUpdateCallback OnAfterUpdateInterface[ProcessShape]
	OnAfterProcessShapeDeleteCallback OnAfterDeleteInterface[ProcessShape]
	OnAfterProcessShapeReadCallback   OnAfterReadInterface[ProcessShape]

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
	stage.DiagramProcesss_reference = make(map[*DiagramProcess]*DiagramProcess)
	stage.DiagramProcesss_instance = make(map[*DiagramProcess]*DiagramProcess)
	stage.DiagramProcesss_referenceOrder = make(map[*DiagramProcess]uint)

	stage.Librarys_reference = make(map[*Library]*Library)
	stage.Librarys_instance = make(map[*Library]*Library)
	stage.Librarys_referenceOrder = make(map[*Library]uint)

	stage.Participants_reference = make(map[*Participant]*Participant)
	stage.Participants_instance = make(map[*Participant]*Participant)
	stage.Participants_referenceOrder = make(map[*Participant]uint)

	stage.Processs_reference = make(map[*Process]*Process)
	stage.Processs_instance = make(map[*Process]*Process)
	stage.Processs_referenceOrder = make(map[*Process]uint)

	stage.ProcessCompositionShapes_reference = make(map[*ProcessCompositionShape]*ProcessCompositionShape)
	stage.ProcessCompositionShapes_instance = make(map[*ProcessCompositionShape]*ProcessCompositionShape)
	stage.ProcessCompositionShapes_referenceOrder = make(map[*ProcessCompositionShape]uint)

	stage.ProcessShapes_reference = make(map[*ProcessShape]*ProcessShape)
	stage.ProcessShapes_instance = make(map[*ProcessShape]*ProcessShape)
	stage.ProcessShapes_referenceOrder = make(map[*ProcessShape]uint)

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
	var maxDiagramProcessOrder uint
	var foundDiagramProcess bool
	for _, order := range stage.DiagramProcess_stagedOrder {
		if !foundDiagramProcess || order > maxDiagramProcessOrder {
			maxDiagramProcessOrder = order
			foundDiagramProcess = true
		}
	}
	if foundDiagramProcess {
		stage.DiagramProcessOrder = maxDiagramProcessOrder + 1
	} else {
		stage.DiagramProcessOrder = 0
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

	var maxParticipantOrder uint
	var foundParticipant bool
	for _, order := range stage.Participant_stagedOrder {
		if !foundParticipant || order > maxParticipantOrder {
			maxParticipantOrder = order
			foundParticipant = true
		}
	}
	if foundParticipant {
		stage.ParticipantOrder = maxParticipantOrder + 1
	} else {
		stage.ParticipantOrder = 0
	}

	var maxProcessOrder uint
	var foundProcess bool
	for _, order := range stage.Process_stagedOrder {
		if !foundProcess || order > maxProcessOrder {
			maxProcessOrder = order
			foundProcess = true
		}
	}
	if foundProcess {
		stage.ProcessOrder = maxProcessOrder + 1
	} else {
		stage.ProcessOrder = 0
	}

	var maxProcessCompositionShapeOrder uint
	var foundProcessCompositionShape bool
	for _, order := range stage.ProcessCompositionShape_stagedOrder {
		if !foundProcessCompositionShape || order > maxProcessCompositionShapeOrder {
			maxProcessCompositionShapeOrder = order
			foundProcessCompositionShape = true
		}
	}
	if foundProcessCompositionShape {
		stage.ProcessCompositionShapeOrder = maxProcessCompositionShapeOrder + 1
	} else {
		stage.ProcessCompositionShapeOrder = 0
	}

	var maxProcessShapeOrder uint
	var foundProcessShape bool
	for _, order := range stage.ProcessShape_stagedOrder {
		if !foundProcessShape || order > maxProcessShapeOrder {
			maxProcessShapeOrder = order
			foundProcessShape = true
		}
	}
	if foundProcessShape {
		stage.ProcessShapeOrder = maxProcessShapeOrder + 1
	} else {
		stage.ProcessShapeOrder = 0
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
	case *DiagramProcess:
		tmp := GetStructInstancesByOrder(stage.DiagramProcesss, stage.DiagramProcess_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *DiagramProcess implements.
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
	case *Participant:
		tmp := GetStructInstancesByOrder(stage.Participants, stage.Participant_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Participant implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Process:
		tmp := GetStructInstancesByOrder(stage.Processs, stage.Process_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Process implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ProcessCompositionShape:
		tmp := GetStructInstancesByOrder(stage.ProcessCompositionShapes, stage.ProcessCompositionShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ProcessCompositionShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ProcessShape:
		tmp := GetStructInstancesByOrder(stage.ProcessShapes, stage.ProcessShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ProcessShape implements.
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
	case "DiagramProcess":
		res = GetNamedStructInstances(stage.DiagramProcesss, stage.DiagramProcess_stagedOrder)
	case "Library":
		res = GetNamedStructInstances(stage.Librarys, stage.Library_stagedOrder)
	case "Participant":
		res = GetNamedStructInstances(stage.Participants, stage.Participant_stagedOrder)
	case "Process":
		res = GetNamedStructInstances(stage.Processs, stage.Process_stagedOrder)
	case "ProcessCompositionShape":
		res = GetNamedStructInstances(stage.ProcessCompositionShapes, stage.ProcessCompositionShape_stagedOrder)
	case "ProcessShape":
		res = GetNamedStructInstances(stage.ProcessShapes, stage.ProcessShape_stagedOrder)
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
	return "github.com/fullstack-lang/gong/dsm/process/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return process_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return process_go.GoDiagramsDir
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
	CommitDiagramProcess(diagramprocess *DiagramProcess)
	CheckoutDiagramProcess(diagramprocess *DiagramProcess)
	CommitLibrary(library *Library)
	CheckoutLibrary(library *Library)
	CommitParticipant(participant *Participant)
	CheckoutParticipant(participant *Participant)
	CommitProcess(process *Process)
	CheckoutProcess(process *Process)
	CommitProcessCompositionShape(processcompositionshape *ProcessCompositionShape)
	CheckoutProcessCompositionShape(processcompositionshape *ProcessCompositionShape)
	CommitProcessShape(processshape *ProcessShape)
	CheckoutProcessShape(processshape *ProcessShape)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		DiagramProcesss:           make(map[*DiagramProcess]struct{}),
		DiagramProcesss_mapString: make(map[string]*DiagramProcess),

		Librarys:           make(map[*Library]struct{}),
		Librarys_mapString: make(map[string]*Library),

		Participants:           make(map[*Participant]struct{}),
		Participants_mapString: make(map[string]*Participant),

		Processs:           make(map[*Process]struct{}),
		Processs_mapString: make(map[string]*Process),

		ProcessCompositionShapes:           make(map[*ProcessCompositionShape]struct{}),
		ProcessCompositionShapes_mapString: make(map[string]*ProcessCompositionShape),

		ProcessShapes:           make(map[*ProcessShape]struct{}),
		ProcessShapes_mapString: make(map[string]*ProcessShape),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		DiagramProcess_stagedOrder: make(map[*DiagramProcess]uint),
		DiagramProcess_orderStaged: make(map[uint]*DiagramProcess),
		DiagramProcesss_reference: make(map[*DiagramProcess]*DiagramProcess),

		Library_stagedOrder: make(map[*Library]uint),
		Library_orderStaged: make(map[uint]*Library),
		Librarys_reference: make(map[*Library]*Library),

		Participant_stagedOrder: make(map[*Participant]uint),
		Participant_orderStaged: make(map[uint]*Participant),
		Participants_reference: make(map[*Participant]*Participant),

		Process_stagedOrder: make(map[*Process]uint),
		Process_orderStaged: make(map[uint]*Process),
		Processs_reference: make(map[*Process]*Process),

		ProcessCompositionShape_stagedOrder: make(map[*ProcessCompositionShape]uint),
		ProcessCompositionShape_orderStaged: make(map[uint]*ProcessCompositionShape),
		ProcessCompositionShapes_reference: make(map[*ProcessCompositionShape]*ProcessCompositionShape),

		ProcessShape_stagedOrder: make(map[*ProcessShape]uint),
		ProcessShape_orderStaged: make(map[uint]*ProcessShape),
		ProcessShapes_reference: make(map[*ProcessShape]*ProcessShape),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"DiagramProcess": &DiagramProcessUnmarshaller{},

			"Library": &LibraryUnmarshaller{},

			"Participant": &ParticipantUnmarshaller{},

			"Process": &ProcessUnmarshaller{},

			"ProcessCompositionShape": &ProcessCompositionShapeUnmarshaller{},

			"ProcessShape": &ProcessShapeUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "DiagramProcess"},
			{name: "Library"},
			{name: "Participant"},
			{name: "Process"},
			{name: "ProcessCompositionShape"},
			{name: "ProcessShape"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *DiagramProcess:
		return stage.DiagramProcess_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *Participant:
		return stage.Participant_stagedOrder[instance]
	case *Process:
		return stage.Process_stagedOrder[instance]
	case *ProcessCompositionShape:
		return stage.ProcessCompositionShape_stagedOrder[instance]
	case *ProcessShape:
		return stage.ProcessShape_stagedOrder[instance]
	default:
		return 0 // should not happen
	}
}

func GongGetInstanceFromOrder[Type PointerToGongstruct](stage *Stage, order uint) (res Type) {
	var t Type
	switch any(t).(type) {
	// insertion point for order map initialisations
	case *DiagramProcess:
		return any(stage.DiagramProcess_orderStaged[order]).(Type)
	case *Library:
		return any(stage.Library_orderStaged[order]).(Type)
	case *Participant:
		return any(stage.Participant_orderStaged[order]).(Type)
	case *Process:
		return any(stage.Process_orderStaged[order]).(Type)
	case *ProcessCompositionShape:
		return any(stage.ProcessCompositionShape_orderStaged[order]).(Type)
	case *ProcessShape:
		return any(stage.ProcessShape_orderStaged[order]).(Type)
	default:
		return // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *DiagramProcess:
		return stage.DiagramProcess_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *Participant:
		return stage.Participant_stagedOrder[instance]
	case *Process:
		return stage.Process_stagedOrder[instance]
	case *ProcessCompositionShape:
		return stage.ProcessCompositionShape_stagedOrder[instance]
	case *ProcessShape:
		return stage.ProcessShape_stagedOrder[instance]
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
	stage.Map_GongStructName_InstancesNb["DiagramProcess"] = len(stage.DiagramProcesss)
	stage.Map_GongStructName_InstancesNb["Library"] = len(stage.Librarys)
	stage.Map_GongStructName_InstancesNb["Participant"] = len(stage.Participants)
	stage.Map_GongStructName_InstancesNb["Process"] = len(stage.Processs)
	stage.Map_GongStructName_InstancesNb["ProcessCompositionShape"] = len(stage.ProcessCompositionShapes)
	stage.Map_GongStructName_InstancesNb["ProcessShape"] = len(stage.ProcessShapes)
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
// Stage puts diagramprocess to the model stage
func (diagramprocess *DiagramProcess) Stage(stage *Stage) *DiagramProcess {
	if _, ok := stage.DiagramProcesss[diagramprocess]; !ok {
		stage.DiagramProcesss[diagramprocess] = struct{}{}
		stage.DiagramProcess_stagedOrder[diagramprocess] = stage.DiagramProcessOrder
		stage.DiagramProcess_orderStaged[stage.DiagramProcessOrder] = diagramprocess
		stage.DiagramProcessOrder++
	}
	stage.DiagramProcesss_mapString[diagramprocess.Name] = diagramprocess

	return diagramprocess
}

// StagePreserveOrder puts diagramprocess to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DiagramProcessOrder
// - update stage.DiagramProcessOrder accordingly
func (diagramprocess *DiagramProcess) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.DiagramProcesss[diagramprocess]; !ok {
		stage.DiagramProcesss[diagramprocess] = struct{}{}

		if order > stage.DiagramProcessOrder {
			stage.DiagramProcessOrder = order
		}
		stage.DiagramProcess_stagedOrder[diagramprocess] = order
		stage.DiagramProcess_orderStaged[order] = diagramprocess
		stage.DiagramProcessOrder++
	}
	stage.DiagramProcesss_mapString[diagramprocess.Name] = diagramprocess
}

// Unstage removes diagramprocess off the model stage
func (diagramprocess *DiagramProcess) Unstage(stage *Stage) *DiagramProcess {
	delete(stage.DiagramProcesss, diagramprocess)
	// issue1150
	// delete(stage.DiagramProcess_stagedOrder, diagramprocess)
	delete(stage.DiagramProcesss_mapString, diagramprocess.Name)

	return diagramprocess
}

// UnstageVoid removes diagramprocess off the model stage
func (diagramprocess *DiagramProcess) UnstageVoid(stage *Stage) {
	delete(stage.DiagramProcesss, diagramprocess)
	// issue1150
	// delete(stage.DiagramProcess_stagedOrder, diagramprocess)
	delete(stage.DiagramProcesss_mapString, diagramprocess.Name)
}

// commit diagramprocess to the back repo (if it is already staged)
func (diagramprocess *DiagramProcess) Commit(stage *Stage) *DiagramProcess {
	if _, ok := stage.DiagramProcesss[diagramprocess]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDiagramProcess(diagramprocess)
		}
	}
	return diagramprocess
}

func (diagramprocess *DiagramProcess) CommitVoid(stage *Stage) {
	diagramprocess.Commit(stage)
}

func (diagramprocess *DiagramProcess) StageVoid(stage *Stage) {
	diagramprocess.Stage(stage)
}

// Checkout diagramprocess to the back repo (if it is already staged)
func (diagramprocess *DiagramProcess) Checkout(stage *Stage) *DiagramProcess {
	if _, ok := stage.DiagramProcesss[diagramprocess]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDiagramProcess(diagramprocess)
		}
	}
	return diagramprocess
}

// for satisfaction of GongStruct interface
func (diagramprocess *DiagramProcess) GetName() (res string) {
	return diagramprocess.Name
}

// for satisfaction of GongStruct interface
func (diagramprocess *DiagramProcess) SetName(name string) {
	diagramprocess.Name = name
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

// Stage puts participant to the model stage
func (participant *Participant) Stage(stage *Stage) *Participant {
	if _, ok := stage.Participants[participant]; !ok {
		stage.Participants[participant] = struct{}{}
		stage.Participant_stagedOrder[participant] = stage.ParticipantOrder
		stage.Participant_orderStaged[stage.ParticipantOrder] = participant
		stage.ParticipantOrder++
	}
	stage.Participants_mapString[participant.Name] = participant

	return participant
}

// StagePreserveOrder puts participant to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ParticipantOrder
// - update stage.ParticipantOrder accordingly
func (participant *Participant) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Participants[participant]; !ok {
		stage.Participants[participant] = struct{}{}

		if order > stage.ParticipantOrder {
			stage.ParticipantOrder = order
		}
		stage.Participant_stagedOrder[participant] = order
		stage.Participant_orderStaged[order] = participant
		stage.ParticipantOrder++
	}
	stage.Participants_mapString[participant.Name] = participant
}

// Unstage removes participant off the model stage
func (participant *Participant) Unstage(stage *Stage) *Participant {
	delete(stage.Participants, participant)
	// issue1150
	// delete(stage.Participant_stagedOrder, participant)
	delete(stage.Participants_mapString, participant.Name)

	return participant
}

// UnstageVoid removes participant off the model stage
func (participant *Participant) UnstageVoid(stage *Stage) {
	delete(stage.Participants, participant)
	// issue1150
	// delete(stage.Participant_stagedOrder, participant)
	delete(stage.Participants_mapString, participant.Name)
}

// commit participant to the back repo (if it is already staged)
func (participant *Participant) Commit(stage *Stage) *Participant {
	if _, ok := stage.Participants[participant]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitParticipant(participant)
		}
	}
	return participant
}

func (participant *Participant) CommitVoid(stage *Stage) {
	participant.Commit(stage)
}

func (participant *Participant) StageVoid(stage *Stage) {
	participant.Stage(stage)
}

// Checkout participant to the back repo (if it is already staged)
func (participant *Participant) Checkout(stage *Stage) *Participant {
	if _, ok := stage.Participants[participant]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutParticipant(participant)
		}
	}
	return participant
}

// for satisfaction of GongStruct interface
func (participant *Participant) GetName() (res string) {
	return participant.Name
}

// for satisfaction of GongStruct interface
func (participant *Participant) SetName(name string) {
	participant.Name = name
}

// Stage puts process to the model stage
func (process *Process) Stage(stage *Stage) *Process {
	if _, ok := stage.Processs[process]; !ok {
		stage.Processs[process] = struct{}{}
		stage.Process_stagedOrder[process] = stage.ProcessOrder
		stage.Process_orderStaged[stage.ProcessOrder] = process
		stage.ProcessOrder++
	}
	stage.Processs_mapString[process.Name] = process

	return process
}

// StagePreserveOrder puts process to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ProcessOrder
// - update stage.ProcessOrder accordingly
func (process *Process) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Processs[process]; !ok {
		stage.Processs[process] = struct{}{}

		if order > stage.ProcessOrder {
			stage.ProcessOrder = order
		}
		stage.Process_stagedOrder[process] = order
		stage.Process_orderStaged[order] = process
		stage.ProcessOrder++
	}
	stage.Processs_mapString[process.Name] = process
}

// Unstage removes process off the model stage
func (process *Process) Unstage(stage *Stage) *Process {
	delete(stage.Processs, process)
	// issue1150
	// delete(stage.Process_stagedOrder, process)
	delete(stage.Processs_mapString, process.Name)

	return process
}

// UnstageVoid removes process off the model stage
func (process *Process) UnstageVoid(stage *Stage) {
	delete(stage.Processs, process)
	// issue1150
	// delete(stage.Process_stagedOrder, process)
	delete(stage.Processs_mapString, process.Name)
}

// commit process to the back repo (if it is already staged)
func (process *Process) Commit(stage *Stage) *Process {
	if _, ok := stage.Processs[process]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitProcess(process)
		}
	}
	return process
}

func (process *Process) CommitVoid(stage *Stage) {
	process.Commit(stage)
}

func (process *Process) StageVoid(stage *Stage) {
	process.Stage(stage)
}

// Checkout process to the back repo (if it is already staged)
func (process *Process) Checkout(stage *Stage) *Process {
	if _, ok := stage.Processs[process]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutProcess(process)
		}
	}
	return process
}

// for satisfaction of GongStruct interface
func (process *Process) GetName() (res string) {
	return process.Name
}

// for satisfaction of GongStruct interface
func (process *Process) SetName(name string) {
	process.Name = name
}

// Stage puts processcompositionshape to the model stage
func (processcompositionshape *ProcessCompositionShape) Stage(stage *Stage) *ProcessCompositionShape {
	if _, ok := stage.ProcessCompositionShapes[processcompositionshape]; !ok {
		stage.ProcessCompositionShapes[processcompositionshape] = struct{}{}
		stage.ProcessCompositionShape_stagedOrder[processcompositionshape] = stage.ProcessCompositionShapeOrder
		stage.ProcessCompositionShape_orderStaged[stage.ProcessCompositionShapeOrder] = processcompositionshape
		stage.ProcessCompositionShapeOrder++
	}
	stage.ProcessCompositionShapes_mapString[processcompositionshape.Name] = processcompositionshape

	return processcompositionshape
}

// StagePreserveOrder puts processcompositionshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ProcessCompositionShapeOrder
// - update stage.ProcessCompositionShapeOrder accordingly
func (processcompositionshape *ProcessCompositionShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ProcessCompositionShapes[processcompositionshape]; !ok {
		stage.ProcessCompositionShapes[processcompositionshape] = struct{}{}

		if order > stage.ProcessCompositionShapeOrder {
			stage.ProcessCompositionShapeOrder = order
		}
		stage.ProcessCompositionShape_stagedOrder[processcompositionshape] = order
		stage.ProcessCompositionShape_orderStaged[order] = processcompositionshape
		stage.ProcessCompositionShapeOrder++
	}
	stage.ProcessCompositionShapes_mapString[processcompositionshape.Name] = processcompositionshape
}

// Unstage removes processcompositionshape off the model stage
func (processcompositionshape *ProcessCompositionShape) Unstage(stage *Stage) *ProcessCompositionShape {
	delete(stage.ProcessCompositionShapes, processcompositionshape)
	// issue1150
	// delete(stage.ProcessCompositionShape_stagedOrder, processcompositionshape)
	delete(stage.ProcessCompositionShapes_mapString, processcompositionshape.Name)

	return processcompositionshape
}

// UnstageVoid removes processcompositionshape off the model stage
func (processcompositionshape *ProcessCompositionShape) UnstageVoid(stage *Stage) {
	delete(stage.ProcessCompositionShapes, processcompositionshape)
	// issue1150
	// delete(stage.ProcessCompositionShape_stagedOrder, processcompositionshape)
	delete(stage.ProcessCompositionShapes_mapString, processcompositionshape.Name)
}

// commit processcompositionshape to the back repo (if it is already staged)
func (processcompositionshape *ProcessCompositionShape) Commit(stage *Stage) *ProcessCompositionShape {
	if _, ok := stage.ProcessCompositionShapes[processcompositionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitProcessCompositionShape(processcompositionshape)
		}
	}
	return processcompositionshape
}

func (processcompositionshape *ProcessCompositionShape) CommitVoid(stage *Stage) {
	processcompositionshape.Commit(stage)
}

func (processcompositionshape *ProcessCompositionShape) StageVoid(stage *Stage) {
	processcompositionshape.Stage(stage)
}

// Checkout processcompositionshape to the back repo (if it is already staged)
func (processcompositionshape *ProcessCompositionShape) Checkout(stage *Stage) *ProcessCompositionShape {
	if _, ok := stage.ProcessCompositionShapes[processcompositionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutProcessCompositionShape(processcompositionshape)
		}
	}
	return processcompositionshape
}

// for satisfaction of GongStruct interface
func (processcompositionshape *ProcessCompositionShape) GetName() (res string) {
	return processcompositionshape.Name
}

// for satisfaction of GongStruct interface
func (processcompositionshape *ProcessCompositionShape) SetName(name string) {
	processcompositionshape.Name = name
}

// Stage puts processshape to the model stage
func (processshape *ProcessShape) Stage(stage *Stage) *ProcessShape {
	if _, ok := stage.ProcessShapes[processshape]; !ok {
		stage.ProcessShapes[processshape] = struct{}{}
		stage.ProcessShape_stagedOrder[processshape] = stage.ProcessShapeOrder
		stage.ProcessShape_orderStaged[stage.ProcessShapeOrder] = processshape
		stage.ProcessShapeOrder++
	}
	stage.ProcessShapes_mapString[processshape.Name] = processshape

	return processshape
}

// StagePreserveOrder puts processshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ProcessShapeOrder
// - update stage.ProcessShapeOrder accordingly
func (processshape *ProcessShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ProcessShapes[processshape]; !ok {
		stage.ProcessShapes[processshape] = struct{}{}

		if order > stage.ProcessShapeOrder {
			stage.ProcessShapeOrder = order
		}
		stage.ProcessShape_stagedOrder[processshape] = order
		stage.ProcessShape_orderStaged[order] = processshape
		stage.ProcessShapeOrder++
	}
	stage.ProcessShapes_mapString[processshape.Name] = processshape
}

// Unstage removes processshape off the model stage
func (processshape *ProcessShape) Unstage(stage *Stage) *ProcessShape {
	delete(stage.ProcessShapes, processshape)
	// issue1150
	// delete(stage.ProcessShape_stagedOrder, processshape)
	delete(stage.ProcessShapes_mapString, processshape.Name)

	return processshape
}

// UnstageVoid removes processshape off the model stage
func (processshape *ProcessShape) UnstageVoid(stage *Stage) {
	delete(stage.ProcessShapes, processshape)
	// issue1150
	// delete(stage.ProcessShape_stagedOrder, processshape)
	delete(stage.ProcessShapes_mapString, processshape.Name)
}

// commit processshape to the back repo (if it is already staged)
func (processshape *ProcessShape) Commit(stage *Stage) *ProcessShape {
	if _, ok := stage.ProcessShapes[processshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitProcessShape(processshape)
		}
	}
	return processshape
}

func (processshape *ProcessShape) CommitVoid(stage *Stage) {
	processshape.Commit(stage)
}

func (processshape *ProcessShape) StageVoid(stage *Stage) {
	processshape.Stage(stage)
}

// Checkout processshape to the back repo (if it is already staged)
func (processshape *ProcessShape) Checkout(stage *Stage) *ProcessShape {
	if _, ok := stage.ProcessShapes[processshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutProcessShape(processshape)
		}
	}
	return processshape
}

// for satisfaction of GongStruct interface
func (processshape *ProcessShape) GetName() (res string) {
	return processshape.Name
}

// for satisfaction of GongStruct interface
func (processshape *ProcessShape) SetName(name string) {
	processshape.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMDiagramProcess(DiagramProcess *DiagramProcess)
	CreateORMLibrary(Library *Library)
	CreateORMParticipant(Participant *Participant)
	CreateORMProcess(Process *Process)
	CreateORMProcessCompositionShape(ProcessCompositionShape *ProcessCompositionShape)
	CreateORMProcessShape(ProcessShape *ProcessShape)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMDiagramProcess(DiagramProcess *DiagramProcess)
	DeleteORMLibrary(Library *Library)
	DeleteORMParticipant(Participant *Participant)
	DeleteORMProcess(Process *Process)
	DeleteORMProcessCompositionShape(ProcessCompositionShape *ProcessCompositionShape)
	DeleteORMProcessShape(ProcessShape *ProcessShape)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.DiagramProcesss = make(map[*DiagramProcess]struct{})
	stage.DiagramProcesss_mapString = make(map[string]*DiagramProcess)
	stage.DiagramProcess_stagedOrder = make(map[*DiagramProcess]uint)
	stage.DiagramProcessOrder = 0

	stage.Librarys = make(map[*Library]struct{})
	stage.Librarys_mapString = make(map[string]*Library)
	stage.Library_stagedOrder = make(map[*Library]uint)
	stage.LibraryOrder = 0

	stage.Participants = make(map[*Participant]struct{})
	stage.Participants_mapString = make(map[string]*Participant)
	stage.Participant_stagedOrder = make(map[*Participant]uint)
	stage.ParticipantOrder = 0

	stage.Processs = make(map[*Process]struct{})
	stage.Processs_mapString = make(map[string]*Process)
	stage.Process_stagedOrder = make(map[*Process]uint)
	stage.ProcessOrder = 0

	stage.ProcessCompositionShapes = make(map[*ProcessCompositionShape]struct{})
	stage.ProcessCompositionShapes_mapString = make(map[string]*ProcessCompositionShape)
	stage.ProcessCompositionShape_stagedOrder = make(map[*ProcessCompositionShape]uint)
	stage.ProcessCompositionShapeOrder = 0

	stage.ProcessShapes = make(map[*ProcessShape]struct{})
	stage.ProcessShapes_mapString = make(map[string]*ProcessShape)
	stage.ProcessShape_stagedOrder = make(map[*ProcessShape]uint)
	stage.ProcessShapeOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.DiagramProcesss = nil
	stage.DiagramProcesss_mapString = nil

	stage.Librarys = nil
	stage.Librarys_mapString = nil

	stage.Participants = nil
	stage.Participants_mapString = nil

	stage.Processs = nil
	stage.Processs_mapString = nil

	stage.ProcessCompositionShapes = nil
	stage.ProcessCompositionShapes_mapString = nil

	stage.ProcessShapes = nil
	stage.ProcessShapes_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for diagramprocess := range stage.DiagramProcesss {
		diagramprocess.Unstage(stage)
	}

	for library := range stage.Librarys {
		library.Unstage(stage)
	}

	for participant := range stage.Participants {
		participant.Unstage(stage)
	}

	for process := range stage.Processs {
		process.Unstage(stage)
	}

	for processcompositionshape := range stage.ProcessCompositionShapes {
		processcompositionshape.Unstage(stage)
	}

	for processshape := range stage.ProcessShapes {
		processshape.Unstage(stage)
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
	case map[*DiagramProcess]any:
		return any(&stage.DiagramProcesss).(*Type)
	case map[*Library]any:
		return any(&stage.Librarys).(*Type)
	case map[*Participant]any:
		return any(&stage.Participants).(*Type)
	case map[*Process]any:
		return any(&stage.Processs).(*Type)
	case map[*ProcessCompositionShape]any:
		return any(&stage.ProcessCompositionShapes).(*Type)
	case map[*ProcessShape]any:
		return any(&stage.ProcessShapes).(*Type)
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
	case *DiagramProcess:
		return any(stage.DiagramProcesss_mapString).(map[string]Type)
	case *Library:
		return any(stage.Librarys_mapString).(map[string]Type)
	case *Participant:
		return any(stage.Participants_mapString).(map[string]Type)
	case *Process:
		return any(stage.Processs_mapString).(map[string]Type)
	case *ProcessCompositionShape:
		return any(stage.ProcessCompositionShapes_mapString).(map[string]Type)
	case *ProcessShape:
		return any(stage.ProcessShapes_mapString).(map[string]Type)
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
	case DiagramProcess:
		return any(&stage.DiagramProcesss).(*map[*Type]struct{})
	case Library:
		return any(&stage.Librarys).(*map[*Type]struct{})
	case Participant:
		return any(&stage.Participants).(*map[*Type]struct{})
	case Process:
		return any(&stage.Processs).(*map[*Type]struct{})
	case ProcessCompositionShape:
		return any(&stage.ProcessCompositionShapes).(*map[*Type]struct{})
	case ProcessShape:
		return any(&stage.ProcessShapes).(*map[*Type]struct{})
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
	case *DiagramProcess:
		return any(&stage.DiagramProcesss).(*map[Type]struct{})
	case *Library:
		return any(&stage.Librarys).(*map[Type]struct{})
	case *Participant:
		return any(&stage.Participants).(*map[Type]struct{})
	case *Process:
		return any(&stage.Processs).(*map[Type]struct{})
	case *ProcessCompositionShape:
		return any(&stage.ProcessCompositionShapes).(*map[Type]struct{})
	case *ProcessShape:
		return any(&stage.ProcessShapes).(*map[Type]struct{})
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
	case DiagramProcess:
		return any(&stage.DiagramProcesss_mapString).(*map[string]*Type)
	case Library:
		return any(&stage.Librarys_mapString).(*map[string]*Type)
	case Participant:
		return any(&stage.Participants_mapString).(*map[string]*Type)
	case Process:
		return any(&stage.Processs_mapString).(*map[string]*Type)
	case ProcessCompositionShape:
		return any(&stage.ProcessCompositionShapes_mapString).(*map[string]*Type)
	case ProcessShape:
		return any(&stage.ProcessShapes_mapString).(*map[string]*Type)
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
	case DiagramProcess:
		return any(&DiagramProcess{
			// Initialisation of associations
			// field is initialized with an instance of ProcessShape with the name of the field
			Process_Shapes: []*ProcessShape{{Name: "Process_Shapes"}},
			// field is initialized with an instance of Process with the name of the field
			ProcesssWhoseNodeIsExpanded: []*Process{{Name: "ProcesssWhoseNodeIsExpanded"}},
			// field is initialized with an instance of ProcessCompositionShape with the name of the field
			ProcessComposition_Shapes: []*ProcessCompositionShape{{Name: "ProcessComposition_Shapes"}},
		}).(*Type)
	case Library:
		return any(&Library{
			// Initialisation of associations
			// field is initialized with an instance of Library with the name of the field
			SubLibraries: []*Library{{Name: "SubLibraries"}},
			// field is initialized with an instance of Process with the name of the field
			RootProcesses: []*Process{{Name: "RootProcesses"}},
			// field is initialized with an instance of Process with the name of the field
			ProcesssWhoseNodeIsExpanded: []*Process{{Name: "ProcesssWhoseNodeIsExpanded"}},
		}).(*Type)
	case Participant:
		return any(&Participant{
			// Initialisation of associations
		}).(*Type)
	case Process:
		return any(&Process{
			// Initialisation of associations
			// field is initialized with an instance of Process with the name of the field
			SubProcesses: []*Process{{Name: "SubProcesses"}},
			// field is initialized with an instance of DiagramProcess with the name of the field
			DiagramProcesss: []*DiagramProcess{{Name: "DiagramProcesss"}},
		}).(*Type)
	case ProcessCompositionShape:
		return any(&ProcessCompositionShape{
			// Initialisation of associations
			// field is initialized with an instance of Process with the name of the field
			Process: &Process{Name: "Process"},
		}).(*Type)
	case ProcessShape:
		return any(&ProcessShape{
			// Initialisation of associations
			// field is initialized with an instance of Process with the name of the field
			Process: &Process{Name: "Process"},
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
	// reverse maps of direct associations of DiagramProcess
	case DiagramProcess:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Library
	case Library:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Participant
	case Participant:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Process
	case Process:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ProcessCompositionShape
	case ProcessCompositionShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Process":
			res := make(map[*Process][]*ProcessCompositionShape)
			for processcompositionshape := range stage.ProcessCompositionShapes {
				if processcompositionshape.Process != nil {
					process_ := processcompositionshape.Process
					var processcompositionshapes []*ProcessCompositionShape
					_, ok := res[process_]
					if ok {
						processcompositionshapes = res[process_]
					} else {
						processcompositionshapes = make([]*ProcessCompositionShape, 0)
					}
					processcompositionshapes = append(processcompositionshapes, processcompositionshape)
					res[process_] = processcompositionshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ProcessShape
	case ProcessShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Process":
			res := make(map[*Process][]*ProcessShape)
			for processshape := range stage.ProcessShapes {
				if processshape.Process != nil {
					process_ := processshape.Process
					var processshapes []*ProcessShape
					_, ok := res[process_]
					if ok {
						processshapes = res[process_]
					} else {
						processshapes = make([]*ProcessShape, 0)
					}
					processshapes = append(processshapes, processshape)
					res[process_] = processshapes
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
	// reverse maps of direct associations of DiagramProcess
	case DiagramProcess:
		switch fieldname {
		// insertion point for per direct association field
		case "Process_Shapes":
			res := make(map[*ProcessShape][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, processshape_ := range diagramprocess.Process_Shapes {
					res[processshape_] = append(res[processshape_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ProcesssWhoseNodeIsExpanded":
			res := make(map[*Process][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, process_ := range diagramprocess.ProcesssWhoseNodeIsExpanded {
					res[process_] = append(res[process_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ProcessComposition_Shapes":
			res := make(map[*ProcessCompositionShape][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, processcompositionshape_ := range diagramprocess.ProcessComposition_Shapes {
					res[processcompositionshape_] = append(res[processcompositionshape_], diagramprocess)
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
		case "RootProcesses":
			res := make(map[*Process][]*Library)
			for library := range stage.Librarys {
				for _, process_ := range library.RootProcesses {
					res[process_] = append(res[process_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ProcesssWhoseNodeIsExpanded":
			res := make(map[*Process][]*Library)
			for library := range stage.Librarys {
				for _, process_ := range library.ProcesssWhoseNodeIsExpanded {
					res[process_] = append(res[process_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Participant
	case Participant:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Process
	case Process:
		switch fieldname {
		// insertion point for per direct association field
		case "SubProcesses":
			res := make(map[*Process][]*Process)
			for process := range stage.Processs {
				for _, process_ := range process.SubProcesses {
					res[process_] = append(res[process_], process)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DiagramProcesss":
			res := make(map[*DiagramProcess][]*Process)
			for process := range stage.Processs {
				for _, diagramprocess_ := range process.DiagramProcesss {
					res[diagramprocess_] = append(res[diagramprocess_], process)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ProcessCompositionShape
	case ProcessCompositionShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ProcessShape
	case ProcessShape:
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
	case *DiagramProcess:
		res = "DiagramProcess"
	case *Library:
		res = "Library"
	case *Participant:
		res = "Participant"
	case *Process:
		res = "Process"
	case *ProcessCompositionShape:
		res = "ProcessCompositionShape"
	case *ProcessShape:
		res = "ProcessShape"
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
	case *DiagramProcess:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Process"
		rf.Fieldname = "DiagramProcesss"
		res = append(res, rf)
	case *Library:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Library"
		rf.Fieldname = "SubLibraries"
		res = append(res, rf)
	case *Participant:
		var rf ReverseField
		_ = rf
	case *Process:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "ProcesssWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootProcesses"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "ProcesssWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Process"
		rf.Fieldname = "SubProcesses"
		res = append(res, rf)
	case *ProcessCompositionShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "ProcessComposition_Shapes"
		res = append(res, rf)
	case *ProcessShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "Process_Shapes"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
func (diagramprocess *DiagramProcess) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
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
			Name:                 "Process_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ProcessShape",
		},
		{
			Name:                 "ProcesssWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Process",
		},
		{
			Name:               "IsProcesssNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ProcessComposition_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ProcessCompositionShape",
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
			Name:               "ComputedPrefix",
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
			Name:                 "RootProcesses",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Process",
		},
		{
			Name:                 "ProcesssWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Process",
		},
	}
	return
}

func (participant *Participant) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (process *Process) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "SubProcesses",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Process",
		},
		{
			Name:                 "DiagramProcesss",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DiagramProcess",
		},
		{
			Name:               "IsSubProcessNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (processcompositionshape *ProcessCompositionShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Process",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Process",
		},
		{
			Name:               "StartRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "StartOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:                 "EndOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:               "CornerOffsetRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (processshape *ProcessShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Process",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Process",
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
func (diagramprocess *DiagramProcess) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = diagramprocess.Name
	case "ComputedPrefix":
		res.valueString = diagramprocess.ComputedPrefix
	case "IsChecked":
		res.valueString = fmt.Sprintf("%t", diagramprocess.IsChecked)
		res.valueBool = diagramprocess.IsChecked
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsEditable_":
		res.valueString = fmt.Sprintf("%t", diagramprocess.IsEditable_)
		res.valueBool = diagramprocess.IsEditable_
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsShowPrefix":
		res.valueString = fmt.Sprintf("%t", diagramprocess.IsShowPrefix)
		res.valueBool = diagramprocess.IsShowPrefix
		res.GongFieldValueType = GongFieldValueTypeBool
	case "DefaultBoxWidth":
		res.valueString = fmt.Sprintf("%f", diagramprocess.DefaultBoxWidth)
		res.valueFloat = diagramprocess.DefaultBoxWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DefaultBoxHeigth":
		res.valueString = fmt.Sprintf("%f", diagramprocess.DefaultBoxHeigth)
		res.valueFloat = diagramprocess.DefaultBoxHeigth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", diagramprocess.Width)
		res.valueFloat = diagramprocess.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", diagramprocess.Height)
		res.valueFloat = diagramprocess.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Process_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.Process_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ProcesssWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.ProcesssWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsProcesssNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagramprocess.IsProcesssNodeExpanded)
		res.valueBool = diagramprocess.IsProcesssNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ProcessComposition_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.ProcessComposition_Shapes {
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
	case "ComputedPrefix":
		res.valueString = library.ComputedPrefix
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
	case "RootProcesses":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootProcesses {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ProcesssWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.ProcesssWhoseNodeIsExpanded {
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

func (participant *Participant) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = participant.Name
	case "ComputedPrefix":
		res.valueString = participant.ComputedPrefix
	}
	return
}

func (process *Process) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = process.Name
	case "ComputedPrefix":
		res.valueString = process.ComputedPrefix
	case "SubProcesses":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range process.SubProcesses {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DiagramProcesss":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range process.DiagramProcesss {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsSubProcessNodeExpanded":
		res.valueString = fmt.Sprintf("%t", process.IsSubProcessNodeExpanded)
		res.valueBool = process.IsSubProcessNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (processcompositionshape *ProcessCompositionShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = processcompositionshape.Name
	case "Process":
		res.GongFieldValueType = GongFieldValueTypePointer
		if processcompositionshape.Process != nil {
			res.valueString = processcompositionshape.Process.Name
			res.ids = processcompositionshape.Process.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", processcompositionshape.StartRatio)
		res.valueFloat = processcompositionshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", processcompositionshape.EndRatio)
		res.valueFloat = processcompositionshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := processcompositionshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := processcompositionshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", processcompositionshape.CornerOffsetRatio)
		res.valueFloat = processcompositionshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", processcompositionshape.IsHidden)
		res.valueBool = processcompositionshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (processshape *ProcessShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = processshape.Name
	case "Process":
		res.GongFieldValueType = GongFieldValueTypePointer
		if processshape.Process != nil {
			res.valueString = processshape.Process.Name
			res.ids = processshape.Process.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", processshape.IsExpanded)
		res.valueBool = processshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", processshape.X)
		res.valueFloat = processshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", processshape.Y)
		res.valueFloat = processshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", processshape.Width)
		res.valueFloat = processshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", processshape.Height)
		res.valueFloat = processshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", processshape.IsHidden)
		res.valueBool = processshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (diagramprocess *DiagramProcess) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		diagramprocess.Name = value.GetValueString()
	case "ComputedPrefix":
		diagramprocess.ComputedPrefix = value.GetValueString()
	case "IsChecked":
		diagramprocess.IsChecked = value.GetValueBool()
	case "IsEditable_":
		diagramprocess.IsEditable_ = value.GetValueBool()
	case "IsShowPrefix":
		diagramprocess.IsShowPrefix = value.GetValueBool()
	case "DefaultBoxWidth":
		diagramprocess.DefaultBoxWidth = value.GetValueFloat()
	case "DefaultBoxHeigth":
		diagramprocess.DefaultBoxHeigth = value.GetValueFloat()
	case "Width":
		diagramprocess.Width = value.GetValueFloat()
	case "Height":
		diagramprocess.Height = value.GetValueFloat()
	case "Process_Shapes":
		diagramprocess.Process_Shapes = make([]*ProcessShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ProcessShapes {
					if stage.ProcessShape_stagedOrder[__instance__] == uint(id) {
						diagramprocess.Process_Shapes = append(diagramprocess.Process_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "ProcesssWhoseNodeIsExpanded":
		diagramprocess.ProcesssWhoseNodeIsExpanded = make([]*Process, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Processs {
					if stage.Process_stagedOrder[__instance__] == uint(id) {
						diagramprocess.ProcesssWhoseNodeIsExpanded = append(diagramprocess.ProcesssWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsProcesssNodeExpanded":
		diagramprocess.IsProcesssNodeExpanded = value.GetValueBool()
	case "ProcessComposition_Shapes":
		diagramprocess.ProcessComposition_Shapes = make([]*ProcessCompositionShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ProcessCompositionShapes {
					if stage.ProcessCompositionShape_stagedOrder[__instance__] == uint(id) {
						diagramprocess.ProcessComposition_Shapes = append(diagramprocess.ProcessComposition_Shapes, __instance__)
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
	case "ComputedPrefix":
		library.ComputedPrefix = value.GetValueString()
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
	case "RootProcesses":
		library.RootProcesses = make([]*Process, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Processs {
					if stage.Process_stagedOrder[__instance__] == uint(id) {
						library.RootProcesses = append(library.RootProcesses, __instance__)
						break
					}
				}
			}
		}
	case "ProcesssWhoseNodeIsExpanded":
		library.ProcesssWhoseNodeIsExpanded = make([]*Process, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Processs {
					if stage.Process_stagedOrder[__instance__] == uint(id) {
						library.ProcesssWhoseNodeIsExpanded = append(library.ProcesssWhoseNodeIsExpanded, __instance__)
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

func (participant *Participant) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		participant.Name = value.GetValueString()
	case "ComputedPrefix":
		participant.ComputedPrefix = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (process *Process) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		process.Name = value.GetValueString()
	case "ComputedPrefix":
		process.ComputedPrefix = value.GetValueString()
	case "SubProcesses":
		process.SubProcesses = make([]*Process, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Processs {
					if stage.Process_stagedOrder[__instance__] == uint(id) {
						process.SubProcesses = append(process.SubProcesses, __instance__)
						break
					}
				}
			}
		}
	case "DiagramProcesss":
		process.DiagramProcesss = make([]*DiagramProcess, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DiagramProcesss {
					if stage.DiagramProcess_stagedOrder[__instance__] == uint(id) {
						process.DiagramProcesss = append(process.DiagramProcesss, __instance__)
						break
					}
				}
			}
		}
	case "IsSubProcessNodeExpanded":
		process.IsSubProcessNodeExpanded = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (processcompositionshape *ProcessCompositionShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		processcompositionshape.Name = value.GetValueString()
	case "Process":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			processcompositionshape.Process = nil
			for __instance__ := range stage.Processs {
				if stage.Process_stagedOrder[__instance__] == uint(id) {
					processcompositionshape.Process = __instance__
					break
				}
			}
		}
	case "StartRatio":
		processcompositionshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		processcompositionshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		processcompositionshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		processcompositionshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		processcompositionshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		processcompositionshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (processshape *ProcessShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		processshape.Name = value.GetValueString()
	case "Process":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			processshape.Process = nil
			for __instance__ := range stage.Processs {
				if stage.Process_stagedOrder[__instance__] == uint(id) {
					processshape.Process = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		processshape.IsExpanded = value.GetValueBool()
	case "X":
		processshape.X = value.GetValueFloat()
	case "Y":
		processshape.Y = value.GetValueFloat()
	case "Width":
		processshape.Width = value.GetValueFloat()
	case "Height":
		processshape.Height = value.GetValueFloat()
	case "IsHidden":
		processshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (diagramprocess *DiagramProcess) GongGetGongstructName() string {
	return "DiagramProcess"
}

func (library *Library) GongGetGongstructName() string {
	return "Library"
}

func (participant *Participant) GongGetGongstructName() string {
	return "Participant"
}

func (process *Process) GongGetGongstructName() string {
	return "Process"
}

func (processcompositionshape *ProcessCompositionShape) GongGetGongstructName() string {
	return "ProcessCompositionShape"
}

func (processshape *ProcessShape) GongGetGongstructName() string {
	return "ProcessShape"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	stage.DiagramProcesss_mapString = make(map[string]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		stage.DiagramProcesss_mapString[diagramprocess.Name] = diagramprocess
	}

	stage.Librarys_mapString = make(map[string]*Library)
	for library := range stage.Librarys {
		stage.Librarys_mapString[library.Name] = library
	}

	stage.Participants_mapString = make(map[string]*Participant)
	for participant := range stage.Participants {
		stage.Participants_mapString[participant.Name] = participant
	}

	stage.Processs_mapString = make(map[string]*Process)
	for process := range stage.Processs {
		stage.Processs_mapString[process.Name] = process
	}

	stage.ProcessCompositionShapes_mapString = make(map[string]*ProcessCompositionShape)
	for processcompositionshape := range stage.ProcessCompositionShapes {
		stage.ProcessCompositionShapes_mapString[processcompositionshape.Name] = processcompositionshape
	}

	stage.ProcessShapes_mapString = make(map[string]*ProcessShape)
	for processshape := range stage.ProcessShapes {
		stage.ProcessShapes_mapString[processshape.Name] = processshape
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
