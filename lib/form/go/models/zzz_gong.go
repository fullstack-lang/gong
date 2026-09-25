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
	CheckBoxs                map[*CheckBox]struct{}
	CheckBoxs_instance       map[*CheckBox]*CheckBox
	CheckBoxs_mapString      map[string]*CheckBox
	CheckBoxOrder            uint
	CheckBox_stagedOrder     map[*CheckBox]uint
	CheckBox_orderStaged     map[uint]*CheckBox
	CheckBoxs_reference      map[*CheckBox]*CheckBox
	CheckBoxs_referenceOrder map[*CheckBox]uint

	// insertion point for slice of pointers maps
	OnAfterCheckBoxCreateCallback GongOnAfterCreateInterface[CheckBox]
	OnAfterCheckBoxUpdateCallback GongOnAfterUpdateInterface[CheckBox]
	OnAfterCheckBoxDeleteCallback GongOnAfterDeleteInterface[CheckBox]

	FormDivs                map[*FormDiv]struct{}
	FormDivs_instance       map[*FormDiv]*FormDiv
	FormDivs_mapString      map[string]*FormDiv
	FormDivOrder            uint
	FormDiv_stagedOrder     map[*FormDiv]uint
	FormDiv_orderStaged     map[uint]*FormDiv
	FormDivs_reference      map[*FormDiv]*FormDiv
	FormDivs_referenceOrder map[*FormDiv]uint

	// insertion point for slice of pointers maps
	FormDiv_FormFields_reverseMap map[*FormField]*FormDiv

	FormDiv_CheckBoxs_reverseMap map[*CheckBox]*FormDiv

	OnAfterFormDivCreateCallback GongOnAfterCreateInterface[FormDiv]
	OnAfterFormDivUpdateCallback GongOnAfterUpdateInterface[FormDiv]
	OnAfterFormDivDeleteCallback GongOnAfterDeleteInterface[FormDiv]

	FormEditAssocButtons                map[*FormEditAssocButton]struct{}
	FormEditAssocButtons_instance       map[*FormEditAssocButton]*FormEditAssocButton
	FormEditAssocButtons_mapString      map[string]*FormEditAssocButton
	FormEditAssocButtonOrder            uint
	FormEditAssocButton_stagedOrder     map[*FormEditAssocButton]uint
	FormEditAssocButton_orderStaged     map[uint]*FormEditAssocButton
	FormEditAssocButtons_reference      map[*FormEditAssocButton]*FormEditAssocButton
	FormEditAssocButtons_referenceOrder map[*FormEditAssocButton]uint

	// insertion point for slice of pointers maps
	OnAfterFormEditAssocButtonCreateCallback GongOnAfterCreateInterface[FormEditAssocButton]
	OnAfterFormEditAssocButtonUpdateCallback GongOnAfterUpdateInterface[FormEditAssocButton]
	OnAfterFormEditAssocButtonDeleteCallback GongOnAfterDeleteInterface[FormEditAssocButton]

	FormFields                map[*FormField]struct{}
	FormFields_instance       map[*FormField]*FormField
	FormFields_mapString      map[string]*FormField
	FormFieldOrder            uint
	FormField_stagedOrder     map[*FormField]uint
	FormField_orderStaged     map[uint]*FormField
	FormFields_reference      map[*FormField]*FormField
	FormFields_referenceOrder map[*FormField]uint

	// insertion point for slice of pointers maps
	OnAfterFormFieldCreateCallback GongOnAfterCreateInterface[FormField]
	OnAfterFormFieldUpdateCallback GongOnAfterUpdateInterface[FormField]
	OnAfterFormFieldDeleteCallback GongOnAfterDeleteInterface[FormField]

	FormFieldDates                map[*FormFieldDate]struct{}
	FormFieldDates_instance       map[*FormFieldDate]*FormFieldDate
	FormFieldDates_mapString      map[string]*FormFieldDate
	FormFieldDateOrder            uint
	FormFieldDate_stagedOrder     map[*FormFieldDate]uint
	FormFieldDate_orderStaged     map[uint]*FormFieldDate
	FormFieldDates_reference      map[*FormFieldDate]*FormFieldDate
	FormFieldDates_referenceOrder map[*FormFieldDate]uint

	// insertion point for slice of pointers maps
	OnAfterFormFieldDateCreateCallback GongOnAfterCreateInterface[FormFieldDate]
	OnAfterFormFieldDateUpdateCallback GongOnAfterUpdateInterface[FormFieldDate]
	OnAfterFormFieldDateDeleteCallback GongOnAfterDeleteInterface[FormFieldDate]

	FormFieldDateTimes                map[*FormFieldDateTime]struct{}
	FormFieldDateTimes_instance       map[*FormFieldDateTime]*FormFieldDateTime
	FormFieldDateTimes_mapString      map[string]*FormFieldDateTime
	FormFieldDateTimeOrder            uint
	FormFieldDateTime_stagedOrder     map[*FormFieldDateTime]uint
	FormFieldDateTime_orderStaged     map[uint]*FormFieldDateTime
	FormFieldDateTimes_reference      map[*FormFieldDateTime]*FormFieldDateTime
	FormFieldDateTimes_referenceOrder map[*FormFieldDateTime]uint

	// insertion point for slice of pointers maps
	OnAfterFormFieldDateTimeCreateCallback GongOnAfterCreateInterface[FormFieldDateTime]
	OnAfterFormFieldDateTimeUpdateCallback GongOnAfterUpdateInterface[FormFieldDateTime]
	OnAfterFormFieldDateTimeDeleteCallback GongOnAfterDeleteInterface[FormFieldDateTime]

	FormFieldFloat64s                map[*FormFieldFloat64]struct{}
	FormFieldFloat64s_instance       map[*FormFieldFloat64]*FormFieldFloat64
	FormFieldFloat64s_mapString      map[string]*FormFieldFloat64
	FormFieldFloat64Order            uint
	FormFieldFloat64_stagedOrder     map[*FormFieldFloat64]uint
	FormFieldFloat64_orderStaged     map[uint]*FormFieldFloat64
	FormFieldFloat64s_reference      map[*FormFieldFloat64]*FormFieldFloat64
	FormFieldFloat64s_referenceOrder map[*FormFieldFloat64]uint

	// insertion point for slice of pointers maps
	OnAfterFormFieldFloat64CreateCallback GongOnAfterCreateInterface[FormFieldFloat64]
	OnAfterFormFieldFloat64UpdateCallback GongOnAfterUpdateInterface[FormFieldFloat64]
	OnAfterFormFieldFloat64DeleteCallback GongOnAfterDeleteInterface[FormFieldFloat64]

	FormFieldInts                map[*FormFieldInt]struct{}
	FormFieldInts_instance       map[*FormFieldInt]*FormFieldInt
	FormFieldInts_mapString      map[string]*FormFieldInt
	FormFieldIntOrder            uint
	FormFieldInt_stagedOrder     map[*FormFieldInt]uint
	FormFieldInt_orderStaged     map[uint]*FormFieldInt
	FormFieldInts_reference      map[*FormFieldInt]*FormFieldInt
	FormFieldInts_referenceOrder map[*FormFieldInt]uint

	// insertion point for slice of pointers maps
	OnAfterFormFieldIntCreateCallback GongOnAfterCreateInterface[FormFieldInt]
	OnAfterFormFieldIntUpdateCallback GongOnAfterUpdateInterface[FormFieldInt]
	OnAfterFormFieldIntDeleteCallback GongOnAfterDeleteInterface[FormFieldInt]

	FormFieldSelects                map[*FormFieldSelect]struct{}
	FormFieldSelects_instance       map[*FormFieldSelect]*FormFieldSelect
	FormFieldSelects_mapString      map[string]*FormFieldSelect
	FormFieldSelectOrder            uint
	FormFieldSelect_stagedOrder     map[*FormFieldSelect]uint
	FormFieldSelect_orderStaged     map[uint]*FormFieldSelect
	FormFieldSelects_reference      map[*FormFieldSelect]*FormFieldSelect
	FormFieldSelects_referenceOrder map[*FormFieldSelect]uint

	// insertion point for slice of pointers maps
	FormFieldSelect_Options_reverseMap map[*Option]*FormFieldSelect

	OnAfterFormFieldSelectCreateCallback GongOnAfterCreateInterface[FormFieldSelect]
	OnAfterFormFieldSelectUpdateCallback GongOnAfterUpdateInterface[FormFieldSelect]
	OnAfterFormFieldSelectDeleteCallback GongOnAfterDeleteInterface[FormFieldSelect]

	FormFieldStrings                map[*FormFieldString]struct{}
	FormFieldStrings_instance       map[*FormFieldString]*FormFieldString
	FormFieldStrings_mapString      map[string]*FormFieldString
	FormFieldStringOrder            uint
	FormFieldString_stagedOrder     map[*FormFieldString]uint
	FormFieldString_orderStaged     map[uint]*FormFieldString
	FormFieldStrings_reference      map[*FormFieldString]*FormFieldString
	FormFieldStrings_referenceOrder map[*FormFieldString]uint

	// insertion point for slice of pointers maps
	OnAfterFormFieldStringCreateCallback GongOnAfterCreateInterface[FormFieldString]
	OnAfterFormFieldStringUpdateCallback GongOnAfterUpdateInterface[FormFieldString]
	OnAfterFormFieldStringDeleteCallback GongOnAfterDeleteInterface[FormFieldString]

	FormFieldTimes                map[*FormFieldTime]struct{}
	FormFieldTimes_instance       map[*FormFieldTime]*FormFieldTime
	FormFieldTimes_mapString      map[string]*FormFieldTime
	FormFieldTimeOrder            uint
	FormFieldTime_stagedOrder     map[*FormFieldTime]uint
	FormFieldTime_orderStaged     map[uint]*FormFieldTime
	FormFieldTimes_reference      map[*FormFieldTime]*FormFieldTime
	FormFieldTimes_referenceOrder map[*FormFieldTime]uint

	// insertion point for slice of pointers maps
	OnAfterFormFieldTimeCreateCallback GongOnAfterCreateInterface[FormFieldTime]
	OnAfterFormFieldTimeUpdateCallback GongOnAfterUpdateInterface[FormFieldTime]
	OnAfterFormFieldTimeDeleteCallback GongOnAfterDeleteInterface[FormFieldTime]

	FormGroups                map[*FormGroup]struct{}
	FormGroups_instance       map[*FormGroup]*FormGroup
	FormGroups_mapString      map[string]*FormGroup
	FormGroupOrder            uint
	FormGroup_stagedOrder     map[*FormGroup]uint
	FormGroup_orderStaged     map[uint]*FormGroup
	FormGroups_reference      map[*FormGroup]*FormGroup
	FormGroups_referenceOrder map[*FormGroup]uint

	// insertion point for slice of pointers maps
	FormGroup_FormDivs_reverseMap map[*FormDiv]*FormGroup

	OnAfterFormGroupCreateCallback GongOnAfterCreateInterface[FormGroup]
	OnAfterFormGroupUpdateCallback GongOnAfterUpdateInterface[FormGroup]
	OnAfterFormGroupDeleteCallback GongOnAfterDeleteInterface[FormGroup]

	FormSortAssocButtons                map[*FormSortAssocButton]struct{}
	FormSortAssocButtons_instance       map[*FormSortAssocButton]*FormSortAssocButton
	FormSortAssocButtons_mapString      map[string]*FormSortAssocButton
	FormSortAssocButtonOrder            uint
	FormSortAssocButton_stagedOrder     map[*FormSortAssocButton]uint
	FormSortAssocButton_orderStaged     map[uint]*FormSortAssocButton
	FormSortAssocButtons_reference      map[*FormSortAssocButton]*FormSortAssocButton
	FormSortAssocButtons_referenceOrder map[*FormSortAssocButton]uint

	// insertion point for slice of pointers maps
	OnAfterFormSortAssocButtonCreateCallback GongOnAfterCreateInterface[FormSortAssocButton]
	OnAfterFormSortAssocButtonUpdateCallback GongOnAfterUpdateInterface[FormSortAssocButton]
	OnAfterFormSortAssocButtonDeleteCallback GongOnAfterDeleteInterface[FormSortAssocButton]

	Options                map[*Option]struct{}
	Options_instance       map[*Option]*Option
	Options_mapString      map[string]*Option
	OptionOrder            uint
	Option_stagedOrder     map[*Option]uint
	Option_orderStaged     map[uint]*Option
	Options_reference      map[*Option]*Option
	Options_referenceOrder map[*Option]uint

	// insertion point for slice of pointers maps
	OnAfterOptionCreateCallback GongOnAfterCreateInterface[Option]
	OnAfterOptionUpdateCallback GongOnAfterUpdateInterface[Option]
	OnAfterOptionDeleteCallback GongOnAfterDeleteInterface[Option]

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
	__gong__clearReferences(&stage.CheckBoxs_reference, &stage.CheckBoxs_instance, &stage.CheckBoxs_referenceOrder)

	__gong__clearReferences(&stage.FormDivs_reference, &stage.FormDivs_instance, &stage.FormDivs_referenceOrder)

	__gong__clearReferences(&stage.FormEditAssocButtons_reference, &stage.FormEditAssocButtons_instance, &stage.FormEditAssocButtons_referenceOrder)

	__gong__clearReferences(&stage.FormFields_reference, &stage.FormFields_instance, &stage.FormFields_referenceOrder)

	__gong__clearReferences(&stage.FormFieldDates_reference, &stage.FormFieldDates_instance, &stage.FormFieldDates_referenceOrder)

	__gong__clearReferences(&stage.FormFieldDateTimes_reference, &stage.FormFieldDateTimes_instance, &stage.FormFieldDateTimes_referenceOrder)

	__gong__clearReferences(&stage.FormFieldFloat64s_reference, &stage.FormFieldFloat64s_instance, &stage.FormFieldFloat64s_referenceOrder)

	__gong__clearReferences(&stage.FormFieldInts_reference, &stage.FormFieldInts_instance, &stage.FormFieldInts_referenceOrder)

	__gong__clearReferences(&stage.FormFieldSelects_reference, &stage.FormFieldSelects_instance, &stage.FormFieldSelects_referenceOrder)

	__gong__clearReferences(&stage.FormFieldStrings_reference, &stage.FormFieldStrings_instance, &stage.FormFieldStrings_referenceOrder)

	__gong__clearReferences(&stage.FormFieldTimes_reference, &stage.FormFieldTimes_instance, &stage.FormFieldTimes_referenceOrder)

	__gong__clearReferences(&stage.FormGroups_reference, &stage.FormGroups_instance, &stage.FormGroups_referenceOrder)

	__gong__clearReferences(&stage.FormSortAssocButtons_reference, &stage.FormSortAssocButtons_instance, &stage.FormSortAssocButtons_referenceOrder)

	__gong__clearReferences(&stage.Options_reference, &stage.Options_instance, &stage.Options_referenceOrder)

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
	stage.CheckBoxOrder = __gong__recomputeOrder(stage.CheckBox_stagedOrder)

	stage.FormDivOrder = __gong__recomputeOrder(stage.FormDiv_stagedOrder)

	stage.FormEditAssocButtonOrder = __gong__recomputeOrder(stage.FormEditAssocButton_stagedOrder)

	stage.FormFieldOrder = __gong__recomputeOrder(stage.FormField_stagedOrder)

	stage.FormFieldDateOrder = __gong__recomputeOrder(stage.FormFieldDate_stagedOrder)

	stage.FormFieldDateTimeOrder = __gong__recomputeOrder(stage.FormFieldDateTime_stagedOrder)

	stage.FormFieldFloat64Order = __gong__recomputeOrder(stage.FormFieldFloat64_stagedOrder)

	stage.FormFieldIntOrder = __gong__recomputeOrder(stage.FormFieldInt_stagedOrder)

	stage.FormFieldSelectOrder = __gong__recomputeOrder(stage.FormFieldSelect_stagedOrder)

	stage.FormFieldStringOrder = __gong__recomputeOrder(stage.FormFieldString_stagedOrder)

	stage.FormFieldTimeOrder = __gong__recomputeOrder(stage.FormFieldTime_stagedOrder)

	stage.FormGroupOrder = __gong__recomputeOrder(stage.FormGroup_stagedOrder)

	stage.FormSortAssocButtonOrder = __gong__recomputeOrder(stage.FormSortAssocButton_stagedOrder)

	stage.OptionOrder = __gong__recomputeOrder(stage.Option_stagedOrder)

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
	case *CheckBox:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.CheckBoxs, stage.CheckBox_stagedOrder))
	case *FormDiv:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.FormDivs, stage.FormDiv_stagedOrder))
	case *FormEditAssocButton:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.FormEditAssocButtons, stage.FormEditAssocButton_stagedOrder))
	case *FormField:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.FormFields, stage.FormField_stagedOrder))
	case *FormFieldDate:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.FormFieldDates, stage.FormFieldDate_stagedOrder))
	case *FormFieldDateTime:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.FormFieldDateTimes, stage.FormFieldDateTime_stagedOrder))
	case *FormFieldFloat64:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.FormFieldFloat64s, stage.FormFieldFloat64_stagedOrder))
	case *FormFieldInt:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.FormFieldInts, stage.FormFieldInt_stagedOrder))
	case *FormFieldSelect:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.FormFieldSelects, stage.FormFieldSelect_stagedOrder))
	case *FormFieldString:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.FormFieldStrings, stage.FormFieldString_stagedOrder))
	case *FormFieldTime:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.FormFieldTimes, stage.FormFieldTime_stagedOrder))
	case *FormGroup:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.FormGroups, stage.FormGroup_stagedOrder))
	case *FormSortAssocButton:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.FormSortAssocButtons, stage.FormSortAssocButton_stagedOrder))
	case *Option:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Options, stage.Option_stagedOrder))

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
	return "github.com/fullstack-lang/gong/lib/form/go/models"
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
		CheckBoxs:           make(map[*CheckBox]struct{}),
		CheckBoxs_mapString: make(map[string]*CheckBox),

		FormDivs:           make(map[*FormDiv]struct{}),
		FormDivs_mapString: make(map[string]*FormDiv),

		FormEditAssocButtons:           make(map[*FormEditAssocButton]struct{}),
		FormEditAssocButtons_mapString: make(map[string]*FormEditAssocButton),

		FormFields:           make(map[*FormField]struct{}),
		FormFields_mapString: make(map[string]*FormField),

		FormFieldDates:           make(map[*FormFieldDate]struct{}),
		FormFieldDates_mapString: make(map[string]*FormFieldDate),

		FormFieldDateTimes:           make(map[*FormFieldDateTime]struct{}),
		FormFieldDateTimes_mapString: make(map[string]*FormFieldDateTime),

		FormFieldFloat64s:           make(map[*FormFieldFloat64]struct{}),
		FormFieldFloat64s_mapString: make(map[string]*FormFieldFloat64),

		FormFieldInts:           make(map[*FormFieldInt]struct{}),
		FormFieldInts_mapString: make(map[string]*FormFieldInt),

		FormFieldSelects:           make(map[*FormFieldSelect]struct{}),
		FormFieldSelects_mapString: make(map[string]*FormFieldSelect),

		FormFieldStrings:           make(map[*FormFieldString]struct{}),
		FormFieldStrings_mapString: make(map[string]*FormFieldString),

		FormFieldTimes:           make(map[*FormFieldTime]struct{}),
		FormFieldTimes_mapString: make(map[string]*FormFieldTime),

		FormGroups:           make(map[*FormGroup]struct{}),
		FormGroups_mapString: make(map[string]*FormGroup),

		FormSortAssocButtons:           make(map[*FormSortAssocButton]struct{}),
		FormSortAssocButtons_mapString: make(map[string]*FormSortAssocButton),

		Options:           make(map[*Option]struct{}),
		Options_mapString: make(map[string]*Option),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		CheckBox_stagedOrder: make(map[*CheckBox]uint),
		CheckBox_orderStaged: make(map[uint]*CheckBox),
		CheckBoxs_reference:  make(map[*CheckBox]*CheckBox),

		FormDiv_stagedOrder: make(map[*FormDiv]uint),
		FormDiv_orderStaged: make(map[uint]*FormDiv),
		FormDivs_reference:  make(map[*FormDiv]*FormDiv),

		FormEditAssocButton_stagedOrder: make(map[*FormEditAssocButton]uint),
		FormEditAssocButton_orderStaged: make(map[uint]*FormEditAssocButton),
		FormEditAssocButtons_reference:  make(map[*FormEditAssocButton]*FormEditAssocButton),

		FormField_stagedOrder: make(map[*FormField]uint),
		FormField_orderStaged: make(map[uint]*FormField),
		FormFields_reference:  make(map[*FormField]*FormField),

		FormFieldDate_stagedOrder: make(map[*FormFieldDate]uint),
		FormFieldDate_orderStaged: make(map[uint]*FormFieldDate),
		FormFieldDates_reference:  make(map[*FormFieldDate]*FormFieldDate),

		FormFieldDateTime_stagedOrder: make(map[*FormFieldDateTime]uint),
		FormFieldDateTime_orderStaged: make(map[uint]*FormFieldDateTime),
		FormFieldDateTimes_reference:  make(map[*FormFieldDateTime]*FormFieldDateTime),

		FormFieldFloat64_stagedOrder: make(map[*FormFieldFloat64]uint),
		FormFieldFloat64_orderStaged: make(map[uint]*FormFieldFloat64),
		FormFieldFloat64s_reference:  make(map[*FormFieldFloat64]*FormFieldFloat64),

		FormFieldInt_stagedOrder: make(map[*FormFieldInt]uint),
		FormFieldInt_orderStaged: make(map[uint]*FormFieldInt),
		FormFieldInts_reference:  make(map[*FormFieldInt]*FormFieldInt),

		FormFieldSelect_stagedOrder: make(map[*FormFieldSelect]uint),
		FormFieldSelect_orderStaged: make(map[uint]*FormFieldSelect),
		FormFieldSelects_reference:  make(map[*FormFieldSelect]*FormFieldSelect),

		FormFieldString_stagedOrder: make(map[*FormFieldString]uint),
		FormFieldString_orderStaged: make(map[uint]*FormFieldString),
		FormFieldStrings_reference:  make(map[*FormFieldString]*FormFieldString),

		FormFieldTime_stagedOrder: make(map[*FormFieldTime]uint),
		FormFieldTime_orderStaged: make(map[uint]*FormFieldTime),
		FormFieldTimes_reference:  make(map[*FormFieldTime]*FormFieldTime),

		FormGroup_stagedOrder: make(map[*FormGroup]uint),
		FormGroup_orderStaged: make(map[uint]*FormGroup),
		FormGroups_reference:  make(map[*FormGroup]*FormGroup),

		FormSortAssocButton_stagedOrder: make(map[*FormSortAssocButton]uint),
		FormSortAssocButton_orderStaged: make(map[uint]*FormSortAssocButton),
		FormSortAssocButtons_reference:  make(map[*FormSortAssocButton]*FormSortAssocButton),

		Option_stagedOrder: make(map[*Option]uint),
		Option_orderStaged: make(map[uint]*Option),
		Options_reference:  make(map[*Option]*Option),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"CheckBox": &CheckBoxUnmarshaller{},

			"FormDiv": &FormDivUnmarshaller{},

			"FormEditAssocButton": &FormEditAssocButtonUnmarshaller{},

			"FormField": &FormFieldUnmarshaller{},

			"FormFieldDate": &FormFieldDateUnmarshaller{},

			"FormFieldDateTime": &FormFieldDateTimeUnmarshaller{},

			"FormFieldFloat64": &FormFieldFloat64Unmarshaller{},

			"FormFieldInt": &FormFieldIntUnmarshaller{},

			"FormFieldSelect": &FormFieldSelectUnmarshaller{},

			"FormFieldString": &FormFieldStringUnmarshaller{},

			"FormFieldTime": &FormFieldTimeUnmarshaller{},

			"FormGroup": &FormGroupUnmarshaller{},

			"FormSortAssocButton": &FormSortAssocButtonUnmarshaller{},

			"Option": &OptionUnmarshaller{},

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
	case *CheckBox:
		return any(stage.CheckBox_orderStaged[order]).(Type)
	case *FormDiv:
		return any(stage.FormDiv_orderStaged[order]).(Type)
	case *FormEditAssocButton:
		return any(stage.FormEditAssocButton_orderStaged[order]).(Type)
	case *FormField:
		return any(stage.FormField_orderStaged[order]).(Type)
	case *FormFieldDate:
		return any(stage.FormFieldDate_orderStaged[order]).(Type)
	case *FormFieldDateTime:
		return any(stage.FormFieldDateTime_orderStaged[order]).(Type)
	case *FormFieldFloat64:
		return any(stage.FormFieldFloat64_orderStaged[order]).(Type)
	case *FormFieldInt:
		return any(stage.FormFieldInt_orderStaged[order]).(Type)
	case *FormFieldSelect:
		return any(stage.FormFieldSelect_orderStaged[order]).(Type)
	case *FormFieldString:
		return any(stage.FormFieldString_orderStaged[order]).(Type)
	case *FormFieldTime:
		return any(stage.FormFieldTime_orderStaged[order]).(Type)
	case *FormGroup:
		return any(stage.FormGroup_orderStaged[order]).(Type)
	case *FormSortAssocButton:
		return any(stage.FormSortAssocButton_orderStaged[order]).(Type)
	case *Option:
		return any(stage.Option_orderStaged[order]).(Type)
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
	stage.Map_GongStructName_InstancesNb["CheckBox"] = len(stage.CheckBoxs)
	stage.Map_GongStructName_InstancesNb["FormDiv"] = len(stage.FormDivs)
	stage.Map_GongStructName_InstancesNb["FormEditAssocButton"] = len(stage.FormEditAssocButtons)
	stage.Map_GongStructName_InstancesNb["FormField"] = len(stage.FormFields)
	stage.Map_GongStructName_InstancesNb["FormFieldDate"] = len(stage.FormFieldDates)
	stage.Map_GongStructName_InstancesNb["FormFieldDateTime"] = len(stage.FormFieldDateTimes)
	stage.Map_GongStructName_InstancesNb["FormFieldFloat64"] = len(stage.FormFieldFloat64s)
	stage.Map_GongStructName_InstancesNb["FormFieldInt"] = len(stage.FormFieldInts)
	stage.Map_GongStructName_InstancesNb["FormFieldSelect"] = len(stage.FormFieldSelects)
	stage.Map_GongStructName_InstancesNb["FormFieldString"] = len(stage.FormFieldStrings)
	stage.Map_GongStructName_InstancesNb["FormFieldTime"] = len(stage.FormFieldTimes)
	stage.Map_GongStructName_InstancesNb["FormGroup"] = len(stage.FormGroups)
	stage.Map_GongStructName_InstancesNb["FormSortAssocButton"] = len(stage.FormSortAssocButtons)
	stage.Map_GongStructName_InstancesNb["Option"] = len(stage.Options)
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
// Stage puts checkbox to the model stage
func (checkbox *CheckBox) Stage(stage *Stage) *CheckBox {
	__gong__stage(stage.CheckBoxs, stage.CheckBox_stagedOrder, stage.CheckBox_orderStaged, &stage.CheckBoxOrder, stage.CheckBoxs_mapString, checkbox, checkbox.Name)
	return checkbox
}

// StagePreserveOrder puts checkbox to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CheckBoxOrder
// - update stage.CheckBoxOrder accordingly
func (checkbox *CheckBox) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CheckBoxs, stage.CheckBox_stagedOrder, stage.CheckBox_orderStaged, &stage.CheckBoxOrder, stage.CheckBoxs_mapString, checkbox, order, checkbox.Name)
}

// Unstage removes checkbox off the model stage
func (checkbox *CheckBox) Unstage(stage *Stage) *CheckBox {
	__gong__unstage(stage.CheckBoxs, stage.CheckBoxs_mapString, checkbox, checkbox.Name)
	return checkbox
}

// UnstageVoid removes checkbox off the model stage
func (checkbox *CheckBox) UnstageVoid(stage *Stage) {
	checkbox.Unstage(stage)
}

func (checkbox *CheckBox) StageVoid(stage *Stage) {
	checkbox.Stage(stage)
}

// for satisfaction of GongStruct interface
func (checkbox *CheckBox) GetName() (res string) {
	return checkbox.Name
}

// for satisfaction of GongStruct interface
func (checkbox *CheckBox) SetName(name string) {
	checkbox.Name = name
}

// Stage puts formdiv to the model stage
func (formdiv *FormDiv) Stage(stage *Stage) *FormDiv {
	__gong__stage(stage.FormDivs, stage.FormDiv_stagedOrder, stage.FormDiv_orderStaged, &stage.FormDivOrder, stage.FormDivs_mapString, formdiv, formdiv.Name)
	return formdiv
}

// StagePreserveOrder puts formdiv to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormDivOrder
// - update stage.FormDivOrder accordingly
func (formdiv *FormDiv) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FormDivs, stage.FormDiv_stagedOrder, stage.FormDiv_orderStaged, &stage.FormDivOrder, stage.FormDivs_mapString, formdiv, order, formdiv.Name)
}

// Unstage removes formdiv off the model stage
func (formdiv *FormDiv) Unstage(stage *Stage) *FormDiv {
	__gong__unstage(stage.FormDivs, stage.FormDivs_mapString, formdiv, formdiv.Name)
	return formdiv
}

// UnstageVoid removes formdiv off the model stage
func (formdiv *FormDiv) UnstageVoid(stage *Stage) {
	formdiv.Unstage(stage)
}

func (formdiv *FormDiv) StageVoid(stage *Stage) {
	formdiv.Stage(stage)
}

// for satisfaction of GongStruct interface
func (formdiv *FormDiv) GetName() (res string) {
	return formdiv.Name
}

// for satisfaction of GongStruct interface
func (formdiv *FormDiv) SetName(name string) {
	formdiv.Name = name
}

// Stage puts formeditassocbutton to the model stage
func (formeditassocbutton *FormEditAssocButton) Stage(stage *Stage) *FormEditAssocButton {
	__gong__stage(stage.FormEditAssocButtons, stage.FormEditAssocButton_stagedOrder, stage.FormEditAssocButton_orderStaged, &stage.FormEditAssocButtonOrder, stage.FormEditAssocButtons_mapString, formeditassocbutton, formeditassocbutton.Name)
	return formeditassocbutton
}

// StagePreserveOrder puts formeditassocbutton to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormEditAssocButtonOrder
// - update stage.FormEditAssocButtonOrder accordingly
func (formeditassocbutton *FormEditAssocButton) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FormEditAssocButtons, stage.FormEditAssocButton_stagedOrder, stage.FormEditAssocButton_orderStaged, &stage.FormEditAssocButtonOrder, stage.FormEditAssocButtons_mapString, formeditassocbutton, order, formeditassocbutton.Name)
}

// Unstage removes formeditassocbutton off the model stage
func (formeditassocbutton *FormEditAssocButton) Unstage(stage *Stage) *FormEditAssocButton {
	__gong__unstage(stage.FormEditAssocButtons, stage.FormEditAssocButtons_mapString, formeditassocbutton, formeditassocbutton.Name)
	return formeditassocbutton
}

// UnstageVoid removes formeditassocbutton off the model stage
func (formeditassocbutton *FormEditAssocButton) UnstageVoid(stage *Stage) {
	formeditassocbutton.Unstage(stage)
}

func (formeditassocbutton *FormEditAssocButton) StageVoid(stage *Stage) {
	formeditassocbutton.Stage(stage)
}

// for satisfaction of GongStruct interface
func (formeditassocbutton *FormEditAssocButton) GetName() (res string) {
	return formeditassocbutton.Name
}

// for satisfaction of GongStruct interface
func (formeditassocbutton *FormEditAssocButton) SetName(name string) {
	formeditassocbutton.Name = name
}

// Stage puts formfield to the model stage
func (formfield *FormField) Stage(stage *Stage) *FormField {
	__gong__stage(stage.FormFields, stage.FormField_stagedOrder, stage.FormField_orderStaged, &stage.FormFieldOrder, stage.FormFields_mapString, formfield, formfield.Name)
	return formfield
}

// StagePreserveOrder puts formfield to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormFieldOrder
// - update stage.FormFieldOrder accordingly
func (formfield *FormField) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FormFields, stage.FormField_stagedOrder, stage.FormField_orderStaged, &stage.FormFieldOrder, stage.FormFields_mapString, formfield, order, formfield.Name)
}

// Unstage removes formfield off the model stage
func (formfield *FormField) Unstage(stage *Stage) *FormField {
	__gong__unstage(stage.FormFields, stage.FormFields_mapString, formfield, formfield.Name)
	return formfield
}

// UnstageVoid removes formfield off the model stage
func (formfield *FormField) UnstageVoid(stage *Stage) {
	formfield.Unstage(stage)
}

func (formfield *FormField) StageVoid(stage *Stage) {
	formfield.Stage(stage)
}

// for satisfaction of GongStruct interface
func (formfield *FormField) GetName() (res string) {
	return formfield.Name
}

// for satisfaction of GongStruct interface
func (formfield *FormField) SetName(name string) {
	formfield.Name = name
}

// Stage puts formfielddate to the model stage
func (formfielddate *FormFieldDate) Stage(stage *Stage) *FormFieldDate {
	__gong__stage(stage.FormFieldDates, stage.FormFieldDate_stagedOrder, stage.FormFieldDate_orderStaged, &stage.FormFieldDateOrder, stage.FormFieldDates_mapString, formfielddate, formfielddate.Name)
	return formfielddate
}

// StagePreserveOrder puts formfielddate to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormFieldDateOrder
// - update stage.FormFieldDateOrder accordingly
func (formfielddate *FormFieldDate) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FormFieldDates, stage.FormFieldDate_stagedOrder, stage.FormFieldDate_orderStaged, &stage.FormFieldDateOrder, stage.FormFieldDates_mapString, formfielddate, order, formfielddate.Name)
}

// Unstage removes formfielddate off the model stage
func (formfielddate *FormFieldDate) Unstage(stage *Stage) *FormFieldDate {
	__gong__unstage(stage.FormFieldDates, stage.FormFieldDates_mapString, formfielddate, formfielddate.Name)
	return formfielddate
}

// UnstageVoid removes formfielddate off the model stage
func (formfielddate *FormFieldDate) UnstageVoid(stage *Stage) {
	formfielddate.Unstage(stage)
}

func (formfielddate *FormFieldDate) StageVoid(stage *Stage) {
	formfielddate.Stage(stage)
}

// for satisfaction of GongStruct interface
func (formfielddate *FormFieldDate) GetName() (res string) {
	return formfielddate.Name
}

// for satisfaction of GongStruct interface
func (formfielddate *FormFieldDate) SetName(name string) {
	formfielddate.Name = name
}

// Stage puts formfielddatetime to the model stage
func (formfielddatetime *FormFieldDateTime) Stage(stage *Stage) *FormFieldDateTime {
	__gong__stage(stage.FormFieldDateTimes, stage.FormFieldDateTime_stagedOrder, stage.FormFieldDateTime_orderStaged, &stage.FormFieldDateTimeOrder, stage.FormFieldDateTimes_mapString, formfielddatetime, formfielddatetime.Name)
	return formfielddatetime
}

// StagePreserveOrder puts formfielddatetime to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormFieldDateTimeOrder
// - update stage.FormFieldDateTimeOrder accordingly
func (formfielddatetime *FormFieldDateTime) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FormFieldDateTimes, stage.FormFieldDateTime_stagedOrder, stage.FormFieldDateTime_orderStaged, &stage.FormFieldDateTimeOrder, stage.FormFieldDateTimes_mapString, formfielddatetime, order, formfielddatetime.Name)
}

// Unstage removes formfielddatetime off the model stage
func (formfielddatetime *FormFieldDateTime) Unstage(stage *Stage) *FormFieldDateTime {
	__gong__unstage(stage.FormFieldDateTimes, stage.FormFieldDateTimes_mapString, formfielddatetime, formfielddatetime.Name)
	return formfielddatetime
}

// UnstageVoid removes formfielddatetime off the model stage
func (formfielddatetime *FormFieldDateTime) UnstageVoid(stage *Stage) {
	formfielddatetime.Unstage(stage)
}

func (formfielddatetime *FormFieldDateTime) StageVoid(stage *Stage) {
	formfielddatetime.Stage(stage)
}

// for satisfaction of GongStruct interface
func (formfielddatetime *FormFieldDateTime) GetName() (res string) {
	return formfielddatetime.Name
}

// for satisfaction of GongStruct interface
func (formfielddatetime *FormFieldDateTime) SetName(name string) {
	formfielddatetime.Name = name
}

// Stage puts formfieldfloat64 to the model stage
func (formfieldfloat64 *FormFieldFloat64) Stage(stage *Stage) *FormFieldFloat64 {
	__gong__stage(stage.FormFieldFloat64s, stage.FormFieldFloat64_stagedOrder, stage.FormFieldFloat64_orderStaged, &stage.FormFieldFloat64Order, stage.FormFieldFloat64s_mapString, formfieldfloat64, formfieldfloat64.Name)
	return formfieldfloat64
}

// StagePreserveOrder puts formfieldfloat64 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormFieldFloat64Order
// - update stage.FormFieldFloat64Order accordingly
func (formfieldfloat64 *FormFieldFloat64) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FormFieldFloat64s, stage.FormFieldFloat64_stagedOrder, stage.FormFieldFloat64_orderStaged, &stage.FormFieldFloat64Order, stage.FormFieldFloat64s_mapString, formfieldfloat64, order, formfieldfloat64.Name)
}

// Unstage removes formfieldfloat64 off the model stage
func (formfieldfloat64 *FormFieldFloat64) Unstage(stage *Stage) *FormFieldFloat64 {
	__gong__unstage(stage.FormFieldFloat64s, stage.FormFieldFloat64s_mapString, formfieldfloat64, formfieldfloat64.Name)
	return formfieldfloat64
}

// UnstageVoid removes formfieldfloat64 off the model stage
func (formfieldfloat64 *FormFieldFloat64) UnstageVoid(stage *Stage) {
	formfieldfloat64.Unstage(stage)
}

func (formfieldfloat64 *FormFieldFloat64) StageVoid(stage *Stage) {
	formfieldfloat64.Stage(stage)
}

// for satisfaction of GongStruct interface
func (formfieldfloat64 *FormFieldFloat64) GetName() (res string) {
	return formfieldfloat64.Name
}

// for satisfaction of GongStruct interface
func (formfieldfloat64 *FormFieldFloat64) SetName(name string) {
	formfieldfloat64.Name = name
}

// Stage puts formfieldint to the model stage
func (formfieldint *FormFieldInt) Stage(stage *Stage) *FormFieldInt {
	__gong__stage(stage.FormFieldInts, stage.FormFieldInt_stagedOrder, stage.FormFieldInt_orderStaged, &stage.FormFieldIntOrder, stage.FormFieldInts_mapString, formfieldint, formfieldint.Name)
	return formfieldint
}

// StagePreserveOrder puts formfieldint to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormFieldIntOrder
// - update stage.FormFieldIntOrder accordingly
func (formfieldint *FormFieldInt) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FormFieldInts, stage.FormFieldInt_stagedOrder, stage.FormFieldInt_orderStaged, &stage.FormFieldIntOrder, stage.FormFieldInts_mapString, formfieldint, order, formfieldint.Name)
}

// Unstage removes formfieldint off the model stage
func (formfieldint *FormFieldInt) Unstage(stage *Stage) *FormFieldInt {
	__gong__unstage(stage.FormFieldInts, stage.FormFieldInts_mapString, formfieldint, formfieldint.Name)
	return formfieldint
}

// UnstageVoid removes formfieldint off the model stage
func (formfieldint *FormFieldInt) UnstageVoid(stage *Stage) {
	formfieldint.Unstage(stage)
}

func (formfieldint *FormFieldInt) StageVoid(stage *Stage) {
	formfieldint.Stage(stage)
}

// for satisfaction of GongStruct interface
func (formfieldint *FormFieldInt) GetName() (res string) {
	return formfieldint.Name
}

// for satisfaction of GongStruct interface
func (formfieldint *FormFieldInt) SetName(name string) {
	formfieldint.Name = name
}

// Stage puts formfieldselect to the model stage
func (formfieldselect *FormFieldSelect) Stage(stage *Stage) *FormFieldSelect {
	__gong__stage(stage.FormFieldSelects, stage.FormFieldSelect_stagedOrder, stage.FormFieldSelect_orderStaged, &stage.FormFieldSelectOrder, stage.FormFieldSelects_mapString, formfieldselect, formfieldselect.Name)
	return formfieldselect
}

// StagePreserveOrder puts formfieldselect to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormFieldSelectOrder
// - update stage.FormFieldSelectOrder accordingly
func (formfieldselect *FormFieldSelect) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FormFieldSelects, stage.FormFieldSelect_stagedOrder, stage.FormFieldSelect_orderStaged, &stage.FormFieldSelectOrder, stage.FormFieldSelects_mapString, formfieldselect, order, formfieldselect.Name)
}

// Unstage removes formfieldselect off the model stage
func (formfieldselect *FormFieldSelect) Unstage(stage *Stage) *FormFieldSelect {
	__gong__unstage(stage.FormFieldSelects, stage.FormFieldSelects_mapString, formfieldselect, formfieldselect.Name)
	return formfieldselect
}

// UnstageVoid removes formfieldselect off the model stage
func (formfieldselect *FormFieldSelect) UnstageVoid(stage *Stage) {
	formfieldselect.Unstage(stage)
}

func (formfieldselect *FormFieldSelect) StageVoid(stage *Stage) {
	formfieldselect.Stage(stage)
}

// for satisfaction of GongStruct interface
func (formfieldselect *FormFieldSelect) GetName() (res string) {
	return formfieldselect.Name
}

// for satisfaction of GongStruct interface
func (formfieldselect *FormFieldSelect) SetName(name string) {
	formfieldselect.Name = name
}

// Stage puts formfieldstring to the model stage
func (formfieldstring *FormFieldString) Stage(stage *Stage) *FormFieldString {
	__gong__stage(stage.FormFieldStrings, stage.FormFieldString_stagedOrder, stage.FormFieldString_orderStaged, &stage.FormFieldStringOrder, stage.FormFieldStrings_mapString, formfieldstring, formfieldstring.Name)
	return formfieldstring
}

// StagePreserveOrder puts formfieldstring to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormFieldStringOrder
// - update stage.FormFieldStringOrder accordingly
func (formfieldstring *FormFieldString) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FormFieldStrings, stage.FormFieldString_stagedOrder, stage.FormFieldString_orderStaged, &stage.FormFieldStringOrder, stage.FormFieldStrings_mapString, formfieldstring, order, formfieldstring.Name)
}

// Unstage removes formfieldstring off the model stage
func (formfieldstring *FormFieldString) Unstage(stage *Stage) *FormFieldString {
	__gong__unstage(stage.FormFieldStrings, stage.FormFieldStrings_mapString, formfieldstring, formfieldstring.Name)
	return formfieldstring
}

// UnstageVoid removes formfieldstring off the model stage
func (formfieldstring *FormFieldString) UnstageVoid(stage *Stage) {
	formfieldstring.Unstage(stage)
}

func (formfieldstring *FormFieldString) StageVoid(stage *Stage) {
	formfieldstring.Stage(stage)
}

// for satisfaction of GongStruct interface
func (formfieldstring *FormFieldString) GetName() (res string) {
	return formfieldstring.Name
}

// for satisfaction of GongStruct interface
func (formfieldstring *FormFieldString) SetName(name string) {
	formfieldstring.Name = name
}

// Stage puts formfieldtime to the model stage
func (formfieldtime *FormFieldTime) Stage(stage *Stage) *FormFieldTime {
	__gong__stage(stage.FormFieldTimes, stage.FormFieldTime_stagedOrder, stage.FormFieldTime_orderStaged, &stage.FormFieldTimeOrder, stage.FormFieldTimes_mapString, formfieldtime, formfieldtime.Name)
	return formfieldtime
}

// StagePreserveOrder puts formfieldtime to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormFieldTimeOrder
// - update stage.FormFieldTimeOrder accordingly
func (formfieldtime *FormFieldTime) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FormFieldTimes, stage.FormFieldTime_stagedOrder, stage.FormFieldTime_orderStaged, &stage.FormFieldTimeOrder, stage.FormFieldTimes_mapString, formfieldtime, order, formfieldtime.Name)
}

// Unstage removes formfieldtime off the model stage
func (formfieldtime *FormFieldTime) Unstage(stage *Stage) *FormFieldTime {
	__gong__unstage(stage.FormFieldTimes, stage.FormFieldTimes_mapString, formfieldtime, formfieldtime.Name)
	return formfieldtime
}

// UnstageVoid removes formfieldtime off the model stage
func (formfieldtime *FormFieldTime) UnstageVoid(stage *Stage) {
	formfieldtime.Unstage(stage)
}

func (formfieldtime *FormFieldTime) StageVoid(stage *Stage) {
	formfieldtime.Stage(stage)
}

// for satisfaction of GongStruct interface
func (formfieldtime *FormFieldTime) GetName() (res string) {
	return formfieldtime.Name
}

// for satisfaction of GongStruct interface
func (formfieldtime *FormFieldTime) SetName(name string) {
	formfieldtime.Name = name
}

// Stage puts formgroup to the model stage
func (formgroup *FormGroup) Stage(stage *Stage) *FormGroup {
	__gong__stage(stage.FormGroups, stage.FormGroup_stagedOrder, stage.FormGroup_orderStaged, &stage.FormGroupOrder, stage.FormGroups_mapString, formgroup, formgroup.Name)
	return formgroup
}

// StagePreserveOrder puts formgroup to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormGroupOrder
// - update stage.FormGroupOrder accordingly
func (formgroup *FormGroup) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FormGroups, stage.FormGroup_stagedOrder, stage.FormGroup_orderStaged, &stage.FormGroupOrder, stage.FormGroups_mapString, formgroup, order, formgroup.Name)
}

// Unstage removes formgroup off the model stage
func (formgroup *FormGroup) Unstage(stage *Stage) *FormGroup {
	__gong__unstage(stage.FormGroups, stage.FormGroups_mapString, formgroup, formgroup.Name)
	return formgroup
}

// UnstageVoid removes formgroup off the model stage
func (formgroup *FormGroup) UnstageVoid(stage *Stage) {
	formgroup.Unstage(stage)
}

func (formgroup *FormGroup) StageVoid(stage *Stage) {
	formgroup.Stage(stage)
}

// for satisfaction of GongStruct interface
func (formgroup *FormGroup) GetName() (res string) {
	return formgroup.Name
}

// for satisfaction of GongStruct interface
func (formgroup *FormGroup) SetName(name string) {
	formgroup.Name = name
}

// Stage puts formsortassocbutton to the model stage
func (formsortassocbutton *FormSortAssocButton) Stage(stage *Stage) *FormSortAssocButton {
	__gong__stage(stage.FormSortAssocButtons, stage.FormSortAssocButton_stagedOrder, stage.FormSortAssocButton_orderStaged, &stage.FormSortAssocButtonOrder, stage.FormSortAssocButtons_mapString, formsortassocbutton, formsortassocbutton.Name)
	return formsortassocbutton
}

// StagePreserveOrder puts formsortassocbutton to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormSortAssocButtonOrder
// - update stage.FormSortAssocButtonOrder accordingly
func (formsortassocbutton *FormSortAssocButton) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FormSortAssocButtons, stage.FormSortAssocButton_stagedOrder, stage.FormSortAssocButton_orderStaged, &stage.FormSortAssocButtonOrder, stage.FormSortAssocButtons_mapString, formsortassocbutton, order, formsortassocbutton.Name)
}

// Unstage removes formsortassocbutton off the model stage
func (formsortassocbutton *FormSortAssocButton) Unstage(stage *Stage) *FormSortAssocButton {
	__gong__unstage(stage.FormSortAssocButtons, stage.FormSortAssocButtons_mapString, formsortassocbutton, formsortassocbutton.Name)
	return formsortassocbutton
}

// UnstageVoid removes formsortassocbutton off the model stage
func (formsortassocbutton *FormSortAssocButton) UnstageVoid(stage *Stage) {
	formsortassocbutton.Unstage(stage)
}

func (formsortassocbutton *FormSortAssocButton) StageVoid(stage *Stage) {
	formsortassocbutton.Stage(stage)
}

// for satisfaction of GongStruct interface
func (formsortassocbutton *FormSortAssocButton) GetName() (res string) {
	return formsortassocbutton.Name
}

// for satisfaction of GongStruct interface
func (formsortassocbutton *FormSortAssocButton) SetName(name string) {
	formsortassocbutton.Name = name
}

// Stage puts option to the model stage
func (option *Option) Stage(stage *Stage) *Option {
	__gong__stage(stage.Options, stage.Option_stagedOrder, stage.Option_orderStaged, &stage.OptionOrder, stage.Options_mapString, option, option.Name)
	return option
}

// StagePreserveOrder puts option to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.OptionOrder
// - update stage.OptionOrder accordingly
func (option *Option) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Options, stage.Option_stagedOrder, stage.Option_orderStaged, &stage.OptionOrder, stage.Options_mapString, option, order, option.Name)
}

// Unstage removes option off the model stage
func (option *Option) Unstage(stage *Stage) *Option {
	__gong__unstage(stage.Options, stage.Options_mapString, option, option.Name)
	return option
}

// UnstageVoid removes option off the model stage
func (option *Option) UnstageVoid(stage *Stage) {
	option.Unstage(stage)
}

func (option *Option) StageVoid(stage *Stage) {
	option.Stage(stage)
}

// for satisfaction of GongStruct interface
func (option *Option) GetName() (res string) {
	return option.Name
}

// for satisfaction of GongStruct interface
func (option *Option) SetName(name string) {
	option.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.CheckBoxs, &stage.CheckBoxs_mapString, &stage.CheckBox_stagedOrder, &stage.CheckBoxOrder)

	__gong__resetStageType(&stage.FormDivs, &stage.FormDivs_mapString, &stage.FormDiv_stagedOrder, &stage.FormDivOrder)

	__gong__resetStageType(&stage.FormEditAssocButtons, &stage.FormEditAssocButtons_mapString, &stage.FormEditAssocButton_stagedOrder, &stage.FormEditAssocButtonOrder)

	__gong__resetStageType(&stage.FormFields, &stage.FormFields_mapString, &stage.FormField_stagedOrder, &stage.FormFieldOrder)

	__gong__resetStageType(&stage.FormFieldDates, &stage.FormFieldDates_mapString, &stage.FormFieldDate_stagedOrder, &stage.FormFieldDateOrder)

	__gong__resetStageType(&stage.FormFieldDateTimes, &stage.FormFieldDateTimes_mapString, &stage.FormFieldDateTime_stagedOrder, &stage.FormFieldDateTimeOrder)

	__gong__resetStageType(&stage.FormFieldFloat64s, &stage.FormFieldFloat64s_mapString, &stage.FormFieldFloat64_stagedOrder, &stage.FormFieldFloat64Order)

	__gong__resetStageType(&stage.FormFieldInts, &stage.FormFieldInts_mapString, &stage.FormFieldInt_stagedOrder, &stage.FormFieldIntOrder)

	__gong__resetStageType(&stage.FormFieldSelects, &stage.FormFieldSelects_mapString, &stage.FormFieldSelect_stagedOrder, &stage.FormFieldSelectOrder)

	__gong__resetStageType(&stage.FormFieldStrings, &stage.FormFieldStrings_mapString, &stage.FormFieldString_stagedOrder, &stage.FormFieldStringOrder)

	__gong__resetStageType(&stage.FormFieldTimes, &stage.FormFieldTimes_mapString, &stage.FormFieldTime_stagedOrder, &stage.FormFieldTimeOrder)

	__gong__resetStageType(&stage.FormGroups, &stage.FormGroups_mapString, &stage.FormGroup_stagedOrder, &stage.FormGroupOrder)

	__gong__resetStageType(&stage.FormSortAssocButtons, &stage.FormSortAssocButtons_mapString, &stage.FormSortAssocButton_stagedOrder, &stage.FormSortAssocButtonOrder)

	__gong__resetStageType(&stage.Options, &stage.Options_mapString, &stage.Option_stagedOrder, &stage.OptionOrder)

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
	case *CheckBox:
		return any(stage.CheckBoxs_mapString).(map[string]Type)
	case *FormDiv:
		return any(stage.FormDivs_mapString).(map[string]Type)
	case *FormEditAssocButton:
		return any(stage.FormEditAssocButtons_mapString).(map[string]Type)
	case *FormField:
		return any(stage.FormFields_mapString).(map[string]Type)
	case *FormFieldDate:
		return any(stage.FormFieldDates_mapString).(map[string]Type)
	case *FormFieldDateTime:
		return any(stage.FormFieldDateTimes_mapString).(map[string]Type)
	case *FormFieldFloat64:
		return any(stage.FormFieldFloat64s_mapString).(map[string]Type)
	case *FormFieldInt:
		return any(stage.FormFieldInts_mapString).(map[string]Type)
	case *FormFieldSelect:
		return any(stage.FormFieldSelects_mapString).(map[string]Type)
	case *FormFieldString:
		return any(stage.FormFieldStrings_mapString).(map[string]Type)
	case *FormFieldTime:
		return any(stage.FormFieldTimes_mapString).(map[string]Type)
	case *FormGroup:
		return any(stage.FormGroups_mapString).(map[string]Type)
	case *FormSortAssocButton:
		return any(stage.FormSortAssocButtons_mapString).(map[string]Type)
	case *Option:
		return any(stage.Options_mapString).(map[string]Type)
	default:
		return nil
	}
}

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *CheckBox:
		return any(&stage.CheckBoxs).(*map[Type]struct{})
	case *FormDiv:
		return any(&stage.FormDivs).(*map[Type]struct{})
	case *FormEditAssocButton:
		return any(&stage.FormEditAssocButtons).(*map[Type]struct{})
	case *FormField:
		return any(&stage.FormFields).(*map[Type]struct{})
	case *FormFieldDate:
		return any(&stage.FormFieldDates).(*map[Type]struct{})
	case *FormFieldDateTime:
		return any(&stage.FormFieldDateTimes).(*map[Type]struct{})
	case *FormFieldFloat64:
		return any(&stage.FormFieldFloat64s).(*map[Type]struct{})
	case *FormFieldInt:
		return any(&stage.FormFieldInts).(*map[Type]struct{})
	case *FormFieldSelect:
		return any(&stage.FormFieldSelects).(*map[Type]struct{})
	case *FormFieldString:
		return any(&stage.FormFieldStrings).(*map[Type]struct{})
	case *FormFieldTime:
		return any(&stage.FormFieldTimes).(*map[Type]struct{})
	case *FormGroup:
		return any(&stage.FormGroups).(*map[Type]struct{})
	case *FormSortAssocButton:
		return any(&stage.FormSortAssocButtons).(*map[Type]struct{})
	case *Option:
		return any(&stage.Options).(*map[Type]struct{})
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
	case FormDiv:
		return any(&FormDiv{
			FormFields: []*FormField{{Name: "FormFields"}},
			CheckBoxs: []*CheckBox{{Name: "CheckBoxs"}},
			FormEditAssocButton: &FormEditAssocButton{Name: "FormEditAssocButton"},
			FormSortAssocButton: &FormSortAssocButton{Name: "FormSortAssocButton"},
		}).(*Type)
	case FormField:
		return any(&FormField{
			FormFieldString: &FormFieldString{Name: "FormFieldString"},
			FormFieldFloat64: &FormFieldFloat64{Name: "FormFieldFloat64"},
			FormFieldInt: &FormFieldInt{Name: "FormFieldInt"},
			FormFieldDate: &FormFieldDate{Name: "FormFieldDate"},
			FormFieldTime: &FormFieldTime{Name: "FormFieldTime"},
			FormFieldDateTime: &FormFieldDateTime{Name: "FormFieldDateTime"},
			FormFieldSelect: &FormFieldSelect{Name: "FormFieldSelect"},
		}).(*Type)
	case FormFieldSelect:
		return any(&FormFieldSelect{
			Value: &Option{Name: "Value"},
			Options: []*Option{{Name: "Options"}},
		}).(*Type)
	case FormGroup:
		return any(&FormGroup{
			FormDivs: []*FormDiv{{Name: "FormDivs"}},
		}).(*Type)
	case FormSortAssocButton:
		return any(&FormSortAssocButton{
			FormEditAssocButton: &FormEditAssocButton{Name: "FormEditAssocButton"},
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
	// reverse maps of direct associations of CheckBox
	case CheckBox:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormDiv
	case FormDiv:
		switch fieldname {
		// insertion point for per direct association field
		case "FormEditAssocButton":
			res := make(map[*FormEditAssocButton][]*FormDiv)
			for formdiv := range stage.FormDivs {
				if formdiv.FormEditAssocButton != nil {
					formeditassocbutton_ := formdiv.FormEditAssocButton
					var formdivs []*FormDiv
					_, ok := res[formeditassocbutton_]
					if ok {
						formdivs = res[formeditassocbutton_]
					} else {
						formdivs = make([]*FormDiv, 0)
					}
					formdivs = append(formdivs, formdiv)
					res[formeditassocbutton_] = formdivs
				}
			}
			return any(res).(map[*End][]*Start)
		case "FormSortAssocButton":
			res := make(map[*FormSortAssocButton][]*FormDiv)
			for formdiv := range stage.FormDivs {
				if formdiv.FormSortAssocButton != nil {
					formsortassocbutton_ := formdiv.FormSortAssocButton
					var formdivs []*FormDiv
					_, ok := res[formsortassocbutton_]
					if ok {
						formdivs = res[formsortassocbutton_]
					} else {
						formdivs = make([]*FormDiv, 0)
					}
					formdivs = append(formdivs, formdiv)
					res[formsortassocbutton_] = formdivs
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of FormEditAssocButton
	case FormEditAssocButton:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormField
	case FormField:
		switch fieldname {
		// insertion point for per direct association field
		case "FormFieldString":
			res := make(map[*FormFieldString][]*FormField)
			for formfield := range stage.FormFields {
				if formfield.FormFieldString != nil {
					formfieldstring_ := formfield.FormFieldString
					var formfields []*FormField
					_, ok := res[formfieldstring_]
					if ok {
						formfields = res[formfieldstring_]
					} else {
						formfields = make([]*FormField, 0)
					}
					formfields = append(formfields, formfield)
					res[formfieldstring_] = formfields
				}
			}
			return any(res).(map[*End][]*Start)
		case "FormFieldFloat64":
			res := make(map[*FormFieldFloat64][]*FormField)
			for formfield := range stage.FormFields {
				if formfield.FormFieldFloat64 != nil {
					formfieldfloat64_ := formfield.FormFieldFloat64
					var formfields []*FormField
					_, ok := res[formfieldfloat64_]
					if ok {
						formfields = res[formfieldfloat64_]
					} else {
						formfields = make([]*FormField, 0)
					}
					formfields = append(formfields, formfield)
					res[formfieldfloat64_] = formfields
				}
			}
			return any(res).(map[*End][]*Start)
		case "FormFieldInt":
			res := make(map[*FormFieldInt][]*FormField)
			for formfield := range stage.FormFields {
				if formfield.FormFieldInt != nil {
					formfieldint_ := formfield.FormFieldInt
					var formfields []*FormField
					_, ok := res[formfieldint_]
					if ok {
						formfields = res[formfieldint_]
					} else {
						formfields = make([]*FormField, 0)
					}
					formfields = append(formfields, formfield)
					res[formfieldint_] = formfields
				}
			}
			return any(res).(map[*End][]*Start)
		case "FormFieldDate":
			res := make(map[*FormFieldDate][]*FormField)
			for formfield := range stage.FormFields {
				if formfield.FormFieldDate != nil {
					formfielddate_ := formfield.FormFieldDate
					var formfields []*FormField
					_, ok := res[formfielddate_]
					if ok {
						formfields = res[formfielddate_]
					} else {
						formfields = make([]*FormField, 0)
					}
					formfields = append(formfields, formfield)
					res[formfielddate_] = formfields
				}
			}
			return any(res).(map[*End][]*Start)
		case "FormFieldTime":
			res := make(map[*FormFieldTime][]*FormField)
			for formfield := range stage.FormFields {
				if formfield.FormFieldTime != nil {
					formfieldtime_ := formfield.FormFieldTime
					var formfields []*FormField
					_, ok := res[formfieldtime_]
					if ok {
						formfields = res[formfieldtime_]
					} else {
						formfields = make([]*FormField, 0)
					}
					formfields = append(formfields, formfield)
					res[formfieldtime_] = formfields
				}
			}
			return any(res).(map[*End][]*Start)
		case "FormFieldDateTime":
			res := make(map[*FormFieldDateTime][]*FormField)
			for formfield := range stage.FormFields {
				if formfield.FormFieldDateTime != nil {
					formfielddatetime_ := formfield.FormFieldDateTime
					var formfields []*FormField
					_, ok := res[formfielddatetime_]
					if ok {
						formfields = res[formfielddatetime_]
					} else {
						formfields = make([]*FormField, 0)
					}
					formfields = append(formfields, formfield)
					res[formfielddatetime_] = formfields
				}
			}
			return any(res).(map[*End][]*Start)
		case "FormFieldSelect":
			res := make(map[*FormFieldSelect][]*FormField)
			for formfield := range stage.FormFields {
				if formfield.FormFieldSelect != nil {
					formfieldselect_ := formfield.FormFieldSelect
					var formfields []*FormField
					_, ok := res[formfieldselect_]
					if ok {
						formfields = res[formfieldselect_]
					} else {
						formfields = make([]*FormField, 0)
					}
					formfields = append(formfields, formfield)
					res[formfieldselect_] = formfields
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of FormFieldDate
	case FormFieldDate:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldDateTime
	case FormFieldDateTime:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldFloat64
	case FormFieldFloat64:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldInt
	case FormFieldInt:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldSelect
	case FormFieldSelect:
		switch fieldname {
		// insertion point for per direct association field
		case "Value":
			res := make(map[*Option][]*FormFieldSelect)
			for formfieldselect := range stage.FormFieldSelects {
				if formfieldselect.Value != nil {
					option_ := formfieldselect.Value
					var formfieldselects []*FormFieldSelect
					_, ok := res[option_]
					if ok {
						formfieldselects = res[option_]
					} else {
						formfieldselects = make([]*FormFieldSelect, 0)
					}
					formfieldselects = append(formfieldselects, formfieldselect)
					res[option_] = formfieldselects
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of FormFieldString
	case FormFieldString:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldTime
	case FormFieldTime:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormGroup
	case FormGroup:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormSortAssocButton
	case FormSortAssocButton:
		switch fieldname {
		// insertion point for per direct association field
		case "FormEditAssocButton":
			res := make(map[*FormEditAssocButton][]*FormSortAssocButton)
			for formsortassocbutton := range stage.FormSortAssocButtons {
				if formsortassocbutton.FormEditAssocButton != nil {
					formeditassocbutton_ := formsortassocbutton.FormEditAssocButton
					var formsortassocbuttons []*FormSortAssocButton
					_, ok := res[formeditassocbutton_]
					if ok {
						formsortassocbuttons = res[formeditassocbutton_]
					} else {
						formsortassocbuttons = make([]*FormSortAssocButton, 0)
					}
					formsortassocbuttons = append(formsortassocbuttons, formsortassocbutton)
					res[formeditassocbutton_] = formsortassocbuttons
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Option
	case Option:
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
	// reverse maps of direct associations of CheckBox
	case CheckBox:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormDiv
	case FormDiv:
		switch fieldname {
		// insertion point for per direct association field
		case "FormFields":
			res := make(map[*FormField][]*FormDiv)
			for formdiv := range stage.FormDivs {
				for _, formfield_ := range formdiv.FormFields {
					res[formfield_] = append(res[formfield_], formdiv)
				}
			}
			return any(res).(map[*End][]*Start)
		case "CheckBoxs":
			res := make(map[*CheckBox][]*FormDiv)
			for formdiv := range stage.FormDivs {
				for _, checkbox_ := range formdiv.CheckBoxs {
					res[checkbox_] = append(res[checkbox_], formdiv)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of FormEditAssocButton
	case FormEditAssocButton:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormField
	case FormField:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldDate
	case FormFieldDate:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldDateTime
	case FormFieldDateTime:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldFloat64
	case FormFieldFloat64:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldInt
	case FormFieldInt:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldSelect
	case FormFieldSelect:
		switch fieldname {
		// insertion point for per direct association field
		case "Options":
			res := make(map[*Option][]*FormFieldSelect)
			for formfieldselect := range stage.FormFieldSelects {
				for _, option_ := range formfieldselect.Options {
					res[option_] = append(res[option_], formfieldselect)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of FormFieldString
	case FormFieldString:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldTime
	case FormFieldTime:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormGroup
	case FormGroup:
		switch fieldname {
		// insertion point for per direct association field
		case "FormDivs":
			res := make(map[*FormDiv][]*FormGroup)
			for formgroup := range stage.FormGroups {
				for _, formdiv_ := range formgroup.FormDivs {
					res[formdiv_] = append(res[formdiv_], formgroup)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of FormSortAssocButton
	case FormSortAssocButton:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Option
	case Option:
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
	case *CheckBox:
		res = any(new(CheckBox)).(Type)
	case *FormDiv:
		res = any(new(FormDiv)).(Type)
	case *FormEditAssocButton:
		res = any(new(FormEditAssocButton)).(Type)
	case *FormField:
		res = any(new(FormField)).(Type)
	case *FormFieldDate:
		res = any(new(FormFieldDate)).(Type)
	case *FormFieldDateTime:
		res = any(new(FormFieldDateTime)).(Type)
	case *FormFieldFloat64:
		res = any(new(FormFieldFloat64)).(Type)
	case *FormFieldInt:
		res = any(new(FormFieldInt)).(Type)
	case *FormFieldSelect:
		res = any(new(FormFieldSelect)).(Type)
	case *FormFieldString:
		res = any(new(FormFieldString)).(Type)
	case *FormFieldTime:
		res = any(new(FormFieldTime)).(Type)
	case *FormGroup:
		res = any(new(FormGroup)).(Type)
	case *FormSortAssocButton:
		res = any(new(FormSortAssocButton)).(Type)
	case *Option:
		res = any(new(Option)).(Type)
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
	case *CheckBox:
		res = "CheckBox"
	case *FormDiv:
		res = "FormDiv"
	case *FormEditAssocButton:
		res = "FormEditAssocButton"
	case *FormField:
		res = "FormField"
	case *FormFieldDate:
		res = "FormFieldDate"
	case *FormFieldDateTime:
		res = "FormFieldDateTime"
	case *FormFieldFloat64:
		res = "FormFieldFloat64"
	case *FormFieldInt:
		res = "FormFieldInt"
	case *FormFieldSelect:
		res = "FormFieldSelect"
	case *FormFieldString:
		res = "FormFieldString"
	case *FormFieldTime:
		res = "FormFieldTime"
	case *FormGroup:
		res = "FormGroup"
	case *FormSortAssocButton:
		res = "FormSortAssocButton"
	case *Option:
		res = "Option"
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
	case *CheckBox:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "FormDiv"
		rf.Fieldname = "CheckBoxs"
		res = append(res, rf)
	case *FormDiv:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "FormGroup"
		rf.Fieldname = "FormDivs"
		res = append(res, rf)
	case *FormEditAssocButton:
		var rf ReverseField
		_ = rf
	case *FormField:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "FormDiv"
		rf.Fieldname = "FormFields"
		res = append(res, rf)
	case *FormFieldDate:
		var rf ReverseField
		_ = rf
	case *FormFieldDateTime:
		var rf ReverseField
		_ = rf
	case *FormFieldFloat64:
		var rf ReverseField
		_ = rf
	case *FormFieldInt:
		var rf ReverseField
		_ = rf
	case *FormFieldSelect:
		var rf ReverseField
		_ = rf
	case *FormFieldString:
		var rf ReverseField
		_ = rf
	case *FormFieldTime:
		var rf ReverseField
		_ = rf
	case *FormGroup:
		var rf ReverseField
		_ = rf
	case *FormSortAssocButton:
		var rf ReverseField
		_ = rf
	case *Option:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "FormFieldSelect"
		rf.Fieldname = "Options"
		res = append(res, rf)
	}
	return
}

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
}

// insertion point for get fields header method
func (checkbox *CheckBox) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (formdiv *FormDiv) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FormFields",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "FormField",
		},
		{
			Name:                 "CheckBoxs",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "CheckBox",
		},
		{
			Name:                 "FormEditAssocButton",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "FormEditAssocButton",
		},
		{
			Name:                 "FormSortAssocButton",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "FormSortAssocButton",
		},
		{
			Name:               "IsADivider",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsAStartAccordionGroup",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "AccordionGroupName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsAEndAccordionGroup",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (formeditassocbutton *FormEditAssocButton) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "AssociationStorage",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "HasChanged",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsForSavePurpose",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "HasToolTip",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ToolTipText",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MatTooltipShowDelay",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (formfield *FormField) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "InputTypeEnum",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "InputTypeEnum",
		},
		{
			Name:               "Label",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Placeholder",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FormFieldString",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "FormFieldString",
		},
		{
			Name:                 "FormFieldFloat64",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "FormFieldFloat64",
		},
		{
			Name:                 "FormFieldInt",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "FormFieldInt",
		},
		{
			Name:                 "FormFieldDate",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "FormFieldDate",
		},
		{
			Name:                 "FormFieldTime",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "FormFieldTime",
		},
		{
			Name:                 "FormFieldDateTime",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "FormFieldDateTime",
		},
		{
			Name:                 "FormFieldSelect",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "FormFieldSelect",
		},
		{
			Name:               "HasBespokeWidth",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "BespokeWidthPx",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "HasBespokeHeight",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "BespokeHeightPx",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (formfielddate *FormFieldDate) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeDate,
		},
	}
	return
}

func (formfielddatetime *FormFieldDateTime) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeDate,
		},
	}
	return
}

func (formfieldfloat64 *FormFieldFloat64) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "HasMinValidator",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "MinValue",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "HasMaxValidator",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "MaxValue",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (formfieldint *FormFieldInt) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "HasMinValidator",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "MinValue",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "HasMaxValidator",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "MaxValue",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (formfieldselect *FormFieldSelect) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Option",
		},
		{
			Name:                 "Options",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Option",
		},
		{
			Name:               "CanBeEmpty",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "PreserveInitialOrder",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (formfieldstring *FormFieldString) GongGetFieldHeaders() (res []GongFieldHeader) {
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
		{
			Name:               "IsTextArea",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (formfieldtime *FormFieldTime) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "Step",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (formgroup *FormGroup) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "TypeLabel",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FormDivs",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "FormDiv",
		},
		{
			Name:               "HasSuppressButton",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "HasSuppressButtonBeenPressed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (formsortassocbutton *FormSortAssocButton) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "HasToolTip",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ToolTipText",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MatTooltipShowDelay",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FormEditAssocButton",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "FormEditAssocButton",
		},
	}
	return
}

func (option *Option) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
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
func (checkbox *CheckBox) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = checkbox.Name
	case "Value":
		res.valueString = fmt.Sprintf("%t", checkbox.Value)
		res.valueBool = checkbox.Value
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (formdiv *FormDiv) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = formdiv.Name
	case "FormFields":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range formdiv.FormFields {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "CheckBoxs":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range formdiv.CheckBoxs {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "FormEditAssocButton":
		res.GongFieldValueType = GongFieldValueTypePointer
		if formdiv.FormEditAssocButton != nil {
			res.valueString = formdiv.FormEditAssocButton.Name
			res.ids = formdiv.FormEditAssocButton.GongGetUUID(stage)
		}
	case "FormSortAssocButton":
		res.GongFieldValueType = GongFieldValueTypePointer
		if formdiv.FormSortAssocButton != nil {
			res.valueString = formdiv.FormSortAssocButton.Name
			res.ids = formdiv.FormSortAssocButton.GongGetUUID(stage)
		}
	case "IsADivider":
		res.valueString = fmt.Sprintf("%t", formdiv.IsADivider)
		res.valueBool = formdiv.IsADivider
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsAStartAccordionGroup":
		res.valueString = fmt.Sprintf("%t", formdiv.IsAStartAccordionGroup)
		res.valueBool = formdiv.IsAStartAccordionGroup
		res.GongFieldValueType = GongFieldValueTypeBool
	case "AccordionGroupName":
		res.valueString = formdiv.AccordionGroupName
	case "IsAEndAccordionGroup":
		res.valueString = fmt.Sprintf("%t", formdiv.IsAEndAccordionGroup)
		res.valueBool = formdiv.IsAEndAccordionGroup
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (formeditassocbutton *FormEditAssocButton) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = formeditassocbutton.Name
	case "Label":
		res.valueString = formeditassocbutton.Label
	case "AssociationStorage":
		res.valueString = formeditassocbutton.AssociationStorage
	case "HasChanged":
		res.valueString = fmt.Sprintf("%t", formeditassocbutton.HasChanged)
		res.valueBool = formeditassocbutton.HasChanged
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsForSavePurpose":
		res.valueString = fmt.Sprintf("%t", formeditassocbutton.IsForSavePurpose)
		res.valueBool = formeditassocbutton.IsForSavePurpose
		res.GongFieldValueType = GongFieldValueTypeBool
	case "HasToolTip":
		res.valueString = fmt.Sprintf("%t", formeditassocbutton.HasToolTip)
		res.valueBool = formeditassocbutton.HasToolTip
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ToolTipText":
		res.valueString = formeditassocbutton.ToolTipText
	case "MatTooltipShowDelay":
		res.valueString = formeditassocbutton.MatTooltipShowDelay
	}
	return
}

func (formfield *FormField) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = formfield.Name
	case "InputTypeEnum":
		enum := formfield.InputTypeEnum
		res.valueString = enum.ToCodeString()
	case "Label":
		res.valueString = formfield.Label
	case "Placeholder":
		res.valueString = formfield.Placeholder
	case "FormFieldString":
		res.GongFieldValueType = GongFieldValueTypePointer
		if formfield.FormFieldString != nil {
			res.valueString = formfield.FormFieldString.Name
			res.ids = formfield.FormFieldString.GongGetUUID(stage)
		}
	case "FormFieldFloat64":
		res.GongFieldValueType = GongFieldValueTypePointer
		if formfield.FormFieldFloat64 != nil {
			res.valueString = formfield.FormFieldFloat64.Name
			res.ids = formfield.FormFieldFloat64.GongGetUUID(stage)
		}
	case "FormFieldInt":
		res.GongFieldValueType = GongFieldValueTypePointer
		if formfield.FormFieldInt != nil {
			res.valueString = formfield.FormFieldInt.Name
			res.ids = formfield.FormFieldInt.GongGetUUID(stage)
		}
	case "FormFieldDate":
		res.GongFieldValueType = GongFieldValueTypePointer
		if formfield.FormFieldDate != nil {
			res.valueString = formfield.FormFieldDate.Name
			res.ids = formfield.FormFieldDate.GongGetUUID(stage)
		}
	case "FormFieldTime":
		res.GongFieldValueType = GongFieldValueTypePointer
		if formfield.FormFieldTime != nil {
			res.valueString = formfield.FormFieldTime.Name
			res.ids = formfield.FormFieldTime.GongGetUUID(stage)
		}
	case "FormFieldDateTime":
		res.GongFieldValueType = GongFieldValueTypePointer
		if formfield.FormFieldDateTime != nil {
			res.valueString = formfield.FormFieldDateTime.Name
			res.ids = formfield.FormFieldDateTime.GongGetUUID(stage)
		}
	case "FormFieldSelect":
		res.GongFieldValueType = GongFieldValueTypePointer
		if formfield.FormFieldSelect != nil {
			res.valueString = formfield.FormFieldSelect.Name
			res.ids = formfield.FormFieldSelect.GongGetUUID(stage)
		}
	case "HasBespokeWidth":
		res.valueString = fmt.Sprintf("%t", formfield.HasBespokeWidth)
		res.valueBool = formfield.HasBespokeWidth
		res.GongFieldValueType = GongFieldValueTypeBool
	case "BespokeWidthPx":
		res.valueString = fmt.Sprintf("%d", formfield.BespokeWidthPx)
		res.valueInt = formfield.BespokeWidthPx
		res.GongFieldValueType = GongFieldValueTypeInt
	case "HasBespokeHeight":
		res.valueString = fmt.Sprintf("%t", formfield.HasBespokeHeight)
		res.valueBool = formfield.HasBespokeHeight
		res.GongFieldValueType = GongFieldValueTypeBool
	case "BespokeHeightPx":
		res.valueString = fmt.Sprintf("%d", formfield.BespokeHeightPx)
		res.valueInt = formfield.BespokeHeightPx
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (formfielddate *FormFieldDate) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = formfielddate.Name
	case "Value":
		res.valueString = formfielddate.Value.String()
	}
	return
}

func (formfielddatetime *FormFieldDateTime) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = formfielddatetime.Name
	case "Value":
		res.valueString = formfielddatetime.Value.String()
	}
	return
}

func (formfieldfloat64 *FormFieldFloat64) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = formfieldfloat64.Name
	case "Value":
		res.valueString = fmt.Sprintf("%f", formfieldfloat64.Value)
		res.valueFloat = formfieldfloat64.Value
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "HasMinValidator":
		res.valueString = fmt.Sprintf("%t", formfieldfloat64.HasMinValidator)
		res.valueBool = formfieldfloat64.HasMinValidator
		res.GongFieldValueType = GongFieldValueTypeBool
	case "MinValue":
		res.valueString = fmt.Sprintf("%f", formfieldfloat64.MinValue)
		res.valueFloat = formfieldfloat64.MinValue
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "HasMaxValidator":
		res.valueString = fmt.Sprintf("%t", formfieldfloat64.HasMaxValidator)
		res.valueBool = formfieldfloat64.HasMaxValidator
		res.GongFieldValueType = GongFieldValueTypeBool
	case "MaxValue":
		res.valueString = fmt.Sprintf("%f", formfieldfloat64.MaxValue)
		res.valueFloat = formfieldfloat64.MaxValue
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (formfieldint *FormFieldInt) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = formfieldint.Name
	case "Value":
		res.valueString = fmt.Sprintf("%d", formfieldint.Value)
		res.valueInt = formfieldint.Value
		res.GongFieldValueType = GongFieldValueTypeInt
	case "HasMinValidator":
		res.valueString = fmt.Sprintf("%t", formfieldint.HasMinValidator)
		res.valueBool = formfieldint.HasMinValidator
		res.GongFieldValueType = GongFieldValueTypeBool
	case "MinValue":
		res.valueString = fmt.Sprintf("%d", formfieldint.MinValue)
		res.valueInt = formfieldint.MinValue
		res.GongFieldValueType = GongFieldValueTypeInt
	case "HasMaxValidator":
		res.valueString = fmt.Sprintf("%t", formfieldint.HasMaxValidator)
		res.valueBool = formfieldint.HasMaxValidator
		res.GongFieldValueType = GongFieldValueTypeBool
	case "MaxValue":
		res.valueString = fmt.Sprintf("%d", formfieldint.MaxValue)
		res.valueInt = formfieldint.MaxValue
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (formfieldselect *FormFieldSelect) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = formfieldselect.Name
	case "Value":
		res.GongFieldValueType = GongFieldValueTypePointer
		if formfieldselect.Value != nil {
			res.valueString = formfieldselect.Value.Name
			res.ids = formfieldselect.Value.GongGetUUID(stage)
		}
	case "Options":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range formfieldselect.Options {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "CanBeEmpty":
		res.valueString = fmt.Sprintf("%t", formfieldselect.CanBeEmpty)
		res.valueBool = formfieldselect.CanBeEmpty
		res.GongFieldValueType = GongFieldValueTypeBool
	case "PreserveInitialOrder":
		res.valueString = fmt.Sprintf("%t", formfieldselect.PreserveInitialOrder)
		res.valueBool = formfieldselect.PreserveInitialOrder
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (formfieldstring *FormFieldString) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = formfieldstring.Name
	case "Value":
		res.valueString = formfieldstring.Value
	case "IsTextArea":
		res.valueString = fmt.Sprintf("%t", formfieldstring.IsTextArea)
		res.valueBool = formfieldstring.IsTextArea
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (formfieldtime *FormFieldTime) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = formfieldtime.Name
	case "Value":
		res.valueString = formfieldtime.Value.String()
	case "Step":
		res.valueString = fmt.Sprintf("%f", formfieldtime.Step)
		res.valueFloat = formfieldtime.Step
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (formgroup *FormGroup) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = formgroup.Name
	case "Label":
		res.valueString = formgroup.Label
	case "TypeLabel":
		res.valueString = formgroup.TypeLabel
	case "FormDivs":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range formgroup.FormDivs {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "HasSuppressButton":
		res.valueString = fmt.Sprintf("%t", formgroup.HasSuppressButton)
		res.valueBool = formgroup.HasSuppressButton
		res.GongFieldValueType = GongFieldValueTypeBool
	case "HasSuppressButtonBeenPressed":
		res.valueString = fmt.Sprintf("%t", formgroup.HasSuppressButtonBeenPressed)
		res.valueBool = formgroup.HasSuppressButtonBeenPressed
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (formsortassocbutton *FormSortAssocButton) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = formsortassocbutton.Name
	case "Label":
		res.valueString = formsortassocbutton.Label
	case "HasToolTip":
		res.valueString = fmt.Sprintf("%t", formsortassocbutton.HasToolTip)
		res.valueBool = formsortassocbutton.HasToolTip
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ToolTipText":
		res.valueString = formsortassocbutton.ToolTipText
	case "MatTooltipShowDelay":
		res.valueString = formsortassocbutton.MatTooltipShowDelay
	case "FormEditAssocButton":
		res.GongFieldValueType = GongFieldValueTypePointer
		if formsortassocbutton.FormEditAssocButton != nil {
			res.valueString = formsortassocbutton.FormEditAssocButton.Name
			res.ids = formsortassocbutton.FormEditAssocButton.GongGetUUID(stage)
		}
	}
	return
}

func (option *Option) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = option.Name
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
func (checkbox *CheckBox) GongGetGongstructName() string {
	return "CheckBox"
}

func (formdiv *FormDiv) GongGetGongstructName() string {
	return "FormDiv"
}

func (formeditassocbutton *FormEditAssocButton) GongGetGongstructName() string {
	return "FormEditAssocButton"
}

func (formfield *FormField) GongGetGongstructName() string {
	return "FormField"
}

func (formfielddate *FormFieldDate) GongGetGongstructName() string {
	return "FormFieldDate"
}

func (formfielddatetime *FormFieldDateTime) GongGetGongstructName() string {
	return "FormFieldDateTime"
}

func (formfieldfloat64 *FormFieldFloat64) GongGetGongstructName() string {
	return "FormFieldFloat64"
}

func (formfieldint *FormFieldInt) GongGetGongstructName() string {
	return "FormFieldInt"
}

func (formfieldselect *FormFieldSelect) GongGetGongstructName() string {
	return "FormFieldSelect"
}

func (formfieldstring *FormFieldString) GongGetGongstructName() string {
	return "FormFieldString"
}

func (formfieldtime *FormFieldTime) GongGetGongstructName() string {
	return "FormFieldTime"
}

func (formgroup *FormGroup) GongGetGongstructName() string {
	return "FormGroup"
}

func (formsortassocbutton *FormSortAssocButton) GongGetGongstructName() string {
	return "FormSortAssocButton"
}

func (option *Option) GongGetGongstructName() string {
	return "Option"
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
	__gong__rebuildMapString(stage.CheckBoxs, &stage.CheckBoxs_mapString)

	__gong__rebuildMapString(stage.FormDivs, &stage.FormDivs_mapString)

	__gong__rebuildMapString(stage.FormEditAssocButtons, &stage.FormEditAssocButtons_mapString)

	__gong__rebuildMapString(stage.FormFields, &stage.FormFields_mapString)

	__gong__rebuildMapString(stage.FormFieldDates, &stage.FormFieldDates_mapString)

	__gong__rebuildMapString(stage.FormFieldDateTimes, &stage.FormFieldDateTimes_mapString)

	__gong__rebuildMapString(stage.FormFieldFloat64s, &stage.FormFieldFloat64s_mapString)

	__gong__rebuildMapString(stage.FormFieldInts, &stage.FormFieldInts_mapString)

	__gong__rebuildMapString(stage.FormFieldSelects, &stage.FormFieldSelects_mapString)

	__gong__rebuildMapString(stage.FormFieldStrings, &stage.FormFieldStrings_mapString)

	__gong__rebuildMapString(stage.FormFieldTimes, &stage.FormFieldTimes_mapString)

	__gong__rebuildMapString(stage.FormGroups, &stage.FormGroups_mapString)

	__gong__rebuildMapString(stage.FormSortAssocButtons, &stage.FormSortAssocButtons_mapString)

	__gong__rebuildMapString(stage.Options, &stage.Options_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
