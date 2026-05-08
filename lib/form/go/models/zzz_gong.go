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

	form_go "github.com/fullstack-lang/gong/lib/form/go"
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
	CheckBoxs                map[*CheckBox]struct{}
	CheckBoxs_instance       map[*CheckBox]*CheckBox
	CheckBoxs_mapString      map[string]*CheckBox
	CheckBoxOrder            uint
	CheckBox_stagedOrder     map[*CheckBox]uint
	CheckBox_orderStaged     map[uint]*CheckBox
	CheckBoxs_reference      map[*CheckBox]*CheckBox
	CheckBoxs_referenceOrder map[*CheckBox]uint

	// insertion point for slice of pointers maps
	OnAfterCheckBoxCreateCallback OnAfterCreateInterface[CheckBox]
	OnAfterCheckBoxUpdateCallback OnAfterUpdateInterface[CheckBox]
	OnAfterCheckBoxDeleteCallback OnAfterDeleteInterface[CheckBox]
	OnAfterCheckBoxReadCallback   OnAfterReadInterface[CheckBox]

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

	OnAfterFormDivCreateCallback OnAfterCreateInterface[FormDiv]
	OnAfterFormDivUpdateCallback OnAfterUpdateInterface[FormDiv]
	OnAfterFormDivDeleteCallback OnAfterDeleteInterface[FormDiv]
	OnAfterFormDivReadCallback   OnAfterReadInterface[FormDiv]

	FormEditAssocButtons                map[*FormEditAssocButton]struct{}
	FormEditAssocButtons_instance       map[*FormEditAssocButton]*FormEditAssocButton
	FormEditAssocButtons_mapString      map[string]*FormEditAssocButton
	FormEditAssocButtonOrder            uint
	FormEditAssocButton_stagedOrder     map[*FormEditAssocButton]uint
	FormEditAssocButton_orderStaged     map[uint]*FormEditAssocButton
	FormEditAssocButtons_reference      map[*FormEditAssocButton]*FormEditAssocButton
	FormEditAssocButtons_referenceOrder map[*FormEditAssocButton]uint

	// insertion point for slice of pointers maps
	OnAfterFormEditAssocButtonCreateCallback OnAfterCreateInterface[FormEditAssocButton]
	OnAfterFormEditAssocButtonUpdateCallback OnAfterUpdateInterface[FormEditAssocButton]
	OnAfterFormEditAssocButtonDeleteCallback OnAfterDeleteInterface[FormEditAssocButton]
	OnAfterFormEditAssocButtonReadCallback   OnAfterReadInterface[FormEditAssocButton]

	FormFields                map[*FormField]struct{}
	FormFields_instance       map[*FormField]*FormField
	FormFields_mapString      map[string]*FormField
	FormFieldOrder            uint
	FormField_stagedOrder     map[*FormField]uint
	FormField_orderStaged     map[uint]*FormField
	FormFields_reference      map[*FormField]*FormField
	FormFields_referenceOrder map[*FormField]uint

	// insertion point for slice of pointers maps
	OnAfterFormFieldCreateCallback OnAfterCreateInterface[FormField]
	OnAfterFormFieldUpdateCallback OnAfterUpdateInterface[FormField]
	OnAfterFormFieldDeleteCallback OnAfterDeleteInterface[FormField]
	OnAfterFormFieldReadCallback   OnAfterReadInterface[FormField]

	FormFieldDates                map[*FormFieldDate]struct{}
	FormFieldDates_instance       map[*FormFieldDate]*FormFieldDate
	FormFieldDates_mapString      map[string]*FormFieldDate
	FormFieldDateOrder            uint
	FormFieldDate_stagedOrder     map[*FormFieldDate]uint
	FormFieldDate_orderStaged     map[uint]*FormFieldDate
	FormFieldDates_reference      map[*FormFieldDate]*FormFieldDate
	FormFieldDates_referenceOrder map[*FormFieldDate]uint

	// insertion point for slice of pointers maps
	OnAfterFormFieldDateCreateCallback OnAfterCreateInterface[FormFieldDate]
	OnAfterFormFieldDateUpdateCallback OnAfterUpdateInterface[FormFieldDate]
	OnAfterFormFieldDateDeleteCallback OnAfterDeleteInterface[FormFieldDate]
	OnAfterFormFieldDateReadCallback   OnAfterReadInterface[FormFieldDate]

	FormFieldDateTimes                map[*FormFieldDateTime]struct{}
	FormFieldDateTimes_instance       map[*FormFieldDateTime]*FormFieldDateTime
	FormFieldDateTimes_mapString      map[string]*FormFieldDateTime
	FormFieldDateTimeOrder            uint
	FormFieldDateTime_stagedOrder     map[*FormFieldDateTime]uint
	FormFieldDateTime_orderStaged     map[uint]*FormFieldDateTime
	FormFieldDateTimes_reference      map[*FormFieldDateTime]*FormFieldDateTime
	FormFieldDateTimes_referenceOrder map[*FormFieldDateTime]uint

	// insertion point for slice of pointers maps
	OnAfterFormFieldDateTimeCreateCallback OnAfterCreateInterface[FormFieldDateTime]
	OnAfterFormFieldDateTimeUpdateCallback OnAfterUpdateInterface[FormFieldDateTime]
	OnAfterFormFieldDateTimeDeleteCallback OnAfterDeleteInterface[FormFieldDateTime]
	OnAfterFormFieldDateTimeReadCallback   OnAfterReadInterface[FormFieldDateTime]

	FormFieldFloat64s                map[*FormFieldFloat64]struct{}
	FormFieldFloat64s_instance       map[*FormFieldFloat64]*FormFieldFloat64
	FormFieldFloat64s_mapString      map[string]*FormFieldFloat64
	FormFieldFloat64Order            uint
	FormFieldFloat64_stagedOrder     map[*FormFieldFloat64]uint
	FormFieldFloat64_orderStaged     map[uint]*FormFieldFloat64
	FormFieldFloat64s_reference      map[*FormFieldFloat64]*FormFieldFloat64
	FormFieldFloat64s_referenceOrder map[*FormFieldFloat64]uint

	// insertion point for slice of pointers maps
	OnAfterFormFieldFloat64CreateCallback OnAfterCreateInterface[FormFieldFloat64]
	OnAfterFormFieldFloat64UpdateCallback OnAfterUpdateInterface[FormFieldFloat64]
	OnAfterFormFieldFloat64DeleteCallback OnAfterDeleteInterface[FormFieldFloat64]
	OnAfterFormFieldFloat64ReadCallback   OnAfterReadInterface[FormFieldFloat64]

	FormFieldInts                map[*FormFieldInt]struct{}
	FormFieldInts_instance       map[*FormFieldInt]*FormFieldInt
	FormFieldInts_mapString      map[string]*FormFieldInt
	FormFieldIntOrder            uint
	FormFieldInt_stagedOrder     map[*FormFieldInt]uint
	FormFieldInt_orderStaged     map[uint]*FormFieldInt
	FormFieldInts_reference      map[*FormFieldInt]*FormFieldInt
	FormFieldInts_referenceOrder map[*FormFieldInt]uint

	// insertion point for slice of pointers maps
	OnAfterFormFieldIntCreateCallback OnAfterCreateInterface[FormFieldInt]
	OnAfterFormFieldIntUpdateCallback OnAfterUpdateInterface[FormFieldInt]
	OnAfterFormFieldIntDeleteCallback OnAfterDeleteInterface[FormFieldInt]
	OnAfterFormFieldIntReadCallback   OnAfterReadInterface[FormFieldInt]

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

	OnAfterFormFieldSelectCreateCallback OnAfterCreateInterface[FormFieldSelect]
	OnAfterFormFieldSelectUpdateCallback OnAfterUpdateInterface[FormFieldSelect]
	OnAfterFormFieldSelectDeleteCallback OnAfterDeleteInterface[FormFieldSelect]
	OnAfterFormFieldSelectReadCallback   OnAfterReadInterface[FormFieldSelect]

	FormFieldStrings                map[*FormFieldString]struct{}
	FormFieldStrings_instance       map[*FormFieldString]*FormFieldString
	FormFieldStrings_mapString      map[string]*FormFieldString
	FormFieldStringOrder            uint
	FormFieldString_stagedOrder     map[*FormFieldString]uint
	FormFieldString_orderStaged     map[uint]*FormFieldString
	FormFieldStrings_reference      map[*FormFieldString]*FormFieldString
	FormFieldStrings_referenceOrder map[*FormFieldString]uint

	// insertion point for slice of pointers maps
	OnAfterFormFieldStringCreateCallback OnAfterCreateInterface[FormFieldString]
	OnAfterFormFieldStringUpdateCallback OnAfterUpdateInterface[FormFieldString]
	OnAfterFormFieldStringDeleteCallback OnAfterDeleteInterface[FormFieldString]
	OnAfterFormFieldStringReadCallback   OnAfterReadInterface[FormFieldString]

	FormFieldTimes                map[*FormFieldTime]struct{}
	FormFieldTimes_instance       map[*FormFieldTime]*FormFieldTime
	FormFieldTimes_mapString      map[string]*FormFieldTime
	FormFieldTimeOrder            uint
	FormFieldTime_stagedOrder     map[*FormFieldTime]uint
	FormFieldTime_orderStaged     map[uint]*FormFieldTime
	FormFieldTimes_reference      map[*FormFieldTime]*FormFieldTime
	FormFieldTimes_referenceOrder map[*FormFieldTime]uint

	// insertion point for slice of pointers maps
	OnAfterFormFieldTimeCreateCallback OnAfterCreateInterface[FormFieldTime]
	OnAfterFormFieldTimeUpdateCallback OnAfterUpdateInterface[FormFieldTime]
	OnAfterFormFieldTimeDeleteCallback OnAfterDeleteInterface[FormFieldTime]
	OnAfterFormFieldTimeReadCallback   OnAfterReadInterface[FormFieldTime]

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

	OnAfterFormGroupCreateCallback OnAfterCreateInterface[FormGroup]
	OnAfterFormGroupUpdateCallback OnAfterUpdateInterface[FormGroup]
	OnAfterFormGroupDeleteCallback OnAfterDeleteInterface[FormGroup]
	OnAfterFormGroupReadCallback   OnAfterReadInterface[FormGroup]

	FormSortAssocButtons                map[*FormSortAssocButton]struct{}
	FormSortAssocButtons_instance       map[*FormSortAssocButton]*FormSortAssocButton
	FormSortAssocButtons_mapString      map[string]*FormSortAssocButton
	FormSortAssocButtonOrder            uint
	FormSortAssocButton_stagedOrder     map[*FormSortAssocButton]uint
	FormSortAssocButton_orderStaged     map[uint]*FormSortAssocButton
	FormSortAssocButtons_reference      map[*FormSortAssocButton]*FormSortAssocButton
	FormSortAssocButtons_referenceOrder map[*FormSortAssocButton]uint

	// insertion point for slice of pointers maps
	OnAfterFormSortAssocButtonCreateCallback OnAfterCreateInterface[FormSortAssocButton]
	OnAfterFormSortAssocButtonUpdateCallback OnAfterUpdateInterface[FormSortAssocButton]
	OnAfterFormSortAssocButtonDeleteCallback OnAfterDeleteInterface[FormSortAssocButton]
	OnAfterFormSortAssocButtonReadCallback   OnAfterReadInterface[FormSortAssocButton]

	Options                map[*Option]struct{}
	Options_instance       map[*Option]*Option
	Options_mapString      map[string]*Option
	OptionOrder            uint
	Option_stagedOrder     map[*Option]uint
	Option_orderStaged     map[uint]*Option
	Options_reference      map[*Option]*Option
	Options_referenceOrder map[*Option]uint

	// insertion point for slice of pointers maps
	OnAfterOptionCreateCallback OnAfterCreateInterface[Option]
	OnAfterOptionUpdateCallback OnAfterUpdateInterface[Option]
	OnAfterOptionDeleteCallback OnAfterDeleteInterface[Option]
	OnAfterOptionReadCallback   OnAfterReadInterface[Option]

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
	stage.CheckBoxs_reference = make(map[*CheckBox]*CheckBox)
	stage.CheckBoxs_instance = make(map[*CheckBox]*CheckBox)
	stage.CheckBoxs_referenceOrder = make(map[*CheckBox]uint)

	stage.FormDivs_reference = make(map[*FormDiv]*FormDiv)
	stage.FormDivs_instance = make(map[*FormDiv]*FormDiv)
	stage.FormDivs_referenceOrder = make(map[*FormDiv]uint)

	stage.FormEditAssocButtons_reference = make(map[*FormEditAssocButton]*FormEditAssocButton)
	stage.FormEditAssocButtons_instance = make(map[*FormEditAssocButton]*FormEditAssocButton)
	stage.FormEditAssocButtons_referenceOrder = make(map[*FormEditAssocButton]uint)

	stage.FormFields_reference = make(map[*FormField]*FormField)
	stage.FormFields_instance = make(map[*FormField]*FormField)
	stage.FormFields_referenceOrder = make(map[*FormField]uint)

	stage.FormFieldDates_reference = make(map[*FormFieldDate]*FormFieldDate)
	stage.FormFieldDates_instance = make(map[*FormFieldDate]*FormFieldDate)
	stage.FormFieldDates_referenceOrder = make(map[*FormFieldDate]uint)

	stage.FormFieldDateTimes_reference = make(map[*FormFieldDateTime]*FormFieldDateTime)
	stage.FormFieldDateTimes_instance = make(map[*FormFieldDateTime]*FormFieldDateTime)
	stage.FormFieldDateTimes_referenceOrder = make(map[*FormFieldDateTime]uint)

	stage.FormFieldFloat64s_reference = make(map[*FormFieldFloat64]*FormFieldFloat64)
	stage.FormFieldFloat64s_instance = make(map[*FormFieldFloat64]*FormFieldFloat64)
	stage.FormFieldFloat64s_referenceOrder = make(map[*FormFieldFloat64]uint)

	stage.FormFieldInts_reference = make(map[*FormFieldInt]*FormFieldInt)
	stage.FormFieldInts_instance = make(map[*FormFieldInt]*FormFieldInt)
	stage.FormFieldInts_referenceOrder = make(map[*FormFieldInt]uint)

	stage.FormFieldSelects_reference = make(map[*FormFieldSelect]*FormFieldSelect)
	stage.FormFieldSelects_instance = make(map[*FormFieldSelect]*FormFieldSelect)
	stage.FormFieldSelects_referenceOrder = make(map[*FormFieldSelect]uint)

	stage.FormFieldStrings_reference = make(map[*FormFieldString]*FormFieldString)
	stage.FormFieldStrings_instance = make(map[*FormFieldString]*FormFieldString)
	stage.FormFieldStrings_referenceOrder = make(map[*FormFieldString]uint)

	stage.FormFieldTimes_reference = make(map[*FormFieldTime]*FormFieldTime)
	stage.FormFieldTimes_instance = make(map[*FormFieldTime]*FormFieldTime)
	stage.FormFieldTimes_referenceOrder = make(map[*FormFieldTime]uint)

	stage.FormGroups_reference = make(map[*FormGroup]*FormGroup)
	stage.FormGroups_instance = make(map[*FormGroup]*FormGroup)
	stage.FormGroups_referenceOrder = make(map[*FormGroup]uint)

	stage.FormSortAssocButtons_reference = make(map[*FormSortAssocButton]*FormSortAssocButton)
	stage.FormSortAssocButtons_instance = make(map[*FormSortAssocButton]*FormSortAssocButton)
	stage.FormSortAssocButtons_referenceOrder = make(map[*FormSortAssocButton]uint)

	stage.Options_reference = make(map[*Option]*Option)
	stage.Options_instance = make(map[*Option]*Option)
	stage.Options_referenceOrder = make(map[*Option]uint)

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
	var maxCheckBoxOrder uint
	var foundCheckBox bool
	for _, order := range stage.CheckBox_stagedOrder {
		if !foundCheckBox || order > maxCheckBoxOrder {
			maxCheckBoxOrder = order
			foundCheckBox = true
		}
	}
	if foundCheckBox {
		stage.CheckBoxOrder = maxCheckBoxOrder + 1
	} else {
		stage.CheckBoxOrder = 0
	}

	var maxFormDivOrder uint
	var foundFormDiv bool
	for _, order := range stage.FormDiv_stagedOrder {
		if !foundFormDiv || order > maxFormDivOrder {
			maxFormDivOrder = order
			foundFormDiv = true
		}
	}
	if foundFormDiv {
		stage.FormDivOrder = maxFormDivOrder + 1
	} else {
		stage.FormDivOrder = 0
	}

	var maxFormEditAssocButtonOrder uint
	var foundFormEditAssocButton bool
	for _, order := range stage.FormEditAssocButton_stagedOrder {
		if !foundFormEditAssocButton || order > maxFormEditAssocButtonOrder {
			maxFormEditAssocButtonOrder = order
			foundFormEditAssocButton = true
		}
	}
	if foundFormEditAssocButton {
		stage.FormEditAssocButtonOrder = maxFormEditAssocButtonOrder + 1
	} else {
		stage.FormEditAssocButtonOrder = 0
	}

	var maxFormFieldOrder uint
	var foundFormField bool
	for _, order := range stage.FormField_stagedOrder {
		if !foundFormField || order > maxFormFieldOrder {
			maxFormFieldOrder = order
			foundFormField = true
		}
	}
	if foundFormField {
		stage.FormFieldOrder = maxFormFieldOrder + 1
	} else {
		stage.FormFieldOrder = 0
	}

	var maxFormFieldDateOrder uint
	var foundFormFieldDate bool
	for _, order := range stage.FormFieldDate_stagedOrder {
		if !foundFormFieldDate || order > maxFormFieldDateOrder {
			maxFormFieldDateOrder = order
			foundFormFieldDate = true
		}
	}
	if foundFormFieldDate {
		stage.FormFieldDateOrder = maxFormFieldDateOrder + 1
	} else {
		stage.FormFieldDateOrder = 0
	}

	var maxFormFieldDateTimeOrder uint
	var foundFormFieldDateTime bool
	for _, order := range stage.FormFieldDateTime_stagedOrder {
		if !foundFormFieldDateTime || order > maxFormFieldDateTimeOrder {
			maxFormFieldDateTimeOrder = order
			foundFormFieldDateTime = true
		}
	}
	if foundFormFieldDateTime {
		stage.FormFieldDateTimeOrder = maxFormFieldDateTimeOrder + 1
	} else {
		stage.FormFieldDateTimeOrder = 0
	}

	var maxFormFieldFloat64Order uint
	var foundFormFieldFloat64 bool
	for _, order := range stage.FormFieldFloat64_stagedOrder {
		if !foundFormFieldFloat64 || order > maxFormFieldFloat64Order {
			maxFormFieldFloat64Order = order
			foundFormFieldFloat64 = true
		}
	}
	if foundFormFieldFloat64 {
		stage.FormFieldFloat64Order = maxFormFieldFloat64Order + 1
	} else {
		stage.FormFieldFloat64Order = 0
	}

	var maxFormFieldIntOrder uint
	var foundFormFieldInt bool
	for _, order := range stage.FormFieldInt_stagedOrder {
		if !foundFormFieldInt || order > maxFormFieldIntOrder {
			maxFormFieldIntOrder = order
			foundFormFieldInt = true
		}
	}
	if foundFormFieldInt {
		stage.FormFieldIntOrder = maxFormFieldIntOrder + 1
	} else {
		stage.FormFieldIntOrder = 0
	}

	var maxFormFieldSelectOrder uint
	var foundFormFieldSelect bool
	for _, order := range stage.FormFieldSelect_stagedOrder {
		if !foundFormFieldSelect || order > maxFormFieldSelectOrder {
			maxFormFieldSelectOrder = order
			foundFormFieldSelect = true
		}
	}
	if foundFormFieldSelect {
		stage.FormFieldSelectOrder = maxFormFieldSelectOrder + 1
	} else {
		stage.FormFieldSelectOrder = 0
	}

	var maxFormFieldStringOrder uint
	var foundFormFieldString bool
	for _, order := range stage.FormFieldString_stagedOrder {
		if !foundFormFieldString || order > maxFormFieldStringOrder {
			maxFormFieldStringOrder = order
			foundFormFieldString = true
		}
	}
	if foundFormFieldString {
		stage.FormFieldStringOrder = maxFormFieldStringOrder + 1
	} else {
		stage.FormFieldStringOrder = 0
	}

	var maxFormFieldTimeOrder uint
	var foundFormFieldTime bool
	for _, order := range stage.FormFieldTime_stagedOrder {
		if !foundFormFieldTime || order > maxFormFieldTimeOrder {
			maxFormFieldTimeOrder = order
			foundFormFieldTime = true
		}
	}
	if foundFormFieldTime {
		stage.FormFieldTimeOrder = maxFormFieldTimeOrder + 1
	} else {
		stage.FormFieldTimeOrder = 0
	}

	var maxFormGroupOrder uint
	var foundFormGroup bool
	for _, order := range stage.FormGroup_stagedOrder {
		if !foundFormGroup || order > maxFormGroupOrder {
			maxFormGroupOrder = order
			foundFormGroup = true
		}
	}
	if foundFormGroup {
		stage.FormGroupOrder = maxFormGroupOrder + 1
	} else {
		stage.FormGroupOrder = 0
	}

	var maxFormSortAssocButtonOrder uint
	var foundFormSortAssocButton bool
	for _, order := range stage.FormSortAssocButton_stagedOrder {
		if !foundFormSortAssocButton || order > maxFormSortAssocButtonOrder {
			maxFormSortAssocButtonOrder = order
			foundFormSortAssocButton = true
		}
	}
	if foundFormSortAssocButton {
		stage.FormSortAssocButtonOrder = maxFormSortAssocButtonOrder + 1
	} else {
		stage.FormSortAssocButtonOrder = 0
	}

	var maxOptionOrder uint
	var foundOption bool
	for _, order := range stage.Option_stagedOrder {
		if !foundOption || order > maxOptionOrder {
			maxOptionOrder = order
			foundOption = true
		}
	}
	if foundOption {
		stage.OptionOrder = maxOptionOrder + 1
	} else {
		stage.OptionOrder = 0
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
	case *CheckBox:
		tmp := GetStructInstancesByOrder(stage.CheckBoxs, stage.CheckBox_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *CheckBox implements.
			res = append(res, any(v).(T))
		}
		return res
	case *FormDiv:
		tmp := GetStructInstancesByOrder(stage.FormDivs, stage.FormDiv_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *FormDiv implements.
			res = append(res, any(v).(T))
		}
		return res
	case *FormEditAssocButton:
		tmp := GetStructInstancesByOrder(stage.FormEditAssocButtons, stage.FormEditAssocButton_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *FormEditAssocButton implements.
			res = append(res, any(v).(T))
		}
		return res
	case *FormField:
		tmp := GetStructInstancesByOrder(stage.FormFields, stage.FormField_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *FormField implements.
			res = append(res, any(v).(T))
		}
		return res
	case *FormFieldDate:
		tmp := GetStructInstancesByOrder(stage.FormFieldDates, stage.FormFieldDate_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *FormFieldDate implements.
			res = append(res, any(v).(T))
		}
		return res
	case *FormFieldDateTime:
		tmp := GetStructInstancesByOrder(stage.FormFieldDateTimes, stage.FormFieldDateTime_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *FormFieldDateTime implements.
			res = append(res, any(v).(T))
		}
		return res
	case *FormFieldFloat64:
		tmp := GetStructInstancesByOrder(stage.FormFieldFloat64s, stage.FormFieldFloat64_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *FormFieldFloat64 implements.
			res = append(res, any(v).(T))
		}
		return res
	case *FormFieldInt:
		tmp := GetStructInstancesByOrder(stage.FormFieldInts, stage.FormFieldInt_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *FormFieldInt implements.
			res = append(res, any(v).(T))
		}
		return res
	case *FormFieldSelect:
		tmp := GetStructInstancesByOrder(stage.FormFieldSelects, stage.FormFieldSelect_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *FormFieldSelect implements.
			res = append(res, any(v).(T))
		}
		return res
	case *FormFieldString:
		tmp := GetStructInstancesByOrder(stage.FormFieldStrings, stage.FormFieldString_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *FormFieldString implements.
			res = append(res, any(v).(T))
		}
		return res
	case *FormFieldTime:
		tmp := GetStructInstancesByOrder(stage.FormFieldTimes, stage.FormFieldTime_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *FormFieldTime implements.
			res = append(res, any(v).(T))
		}
		return res
	case *FormGroup:
		tmp := GetStructInstancesByOrder(stage.FormGroups, stage.FormGroup_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *FormGroup implements.
			res = append(res, any(v).(T))
		}
		return res
	case *FormSortAssocButton:
		tmp := GetStructInstancesByOrder(stage.FormSortAssocButtons, stage.FormSortAssocButton_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *FormSortAssocButton implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Option:
		tmp := GetStructInstancesByOrder(stage.Options, stage.Option_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Option implements.
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
	case "CheckBox":
		res = GetNamedStructInstances(stage.CheckBoxs, stage.CheckBox_stagedOrder)
	case "FormDiv":
		res = GetNamedStructInstances(stage.FormDivs, stage.FormDiv_stagedOrder)
	case "FormEditAssocButton":
		res = GetNamedStructInstances(stage.FormEditAssocButtons, stage.FormEditAssocButton_stagedOrder)
	case "FormField":
		res = GetNamedStructInstances(stage.FormFields, stage.FormField_stagedOrder)
	case "FormFieldDate":
		res = GetNamedStructInstances(stage.FormFieldDates, stage.FormFieldDate_stagedOrder)
	case "FormFieldDateTime":
		res = GetNamedStructInstances(stage.FormFieldDateTimes, stage.FormFieldDateTime_stagedOrder)
	case "FormFieldFloat64":
		res = GetNamedStructInstances(stage.FormFieldFloat64s, stage.FormFieldFloat64_stagedOrder)
	case "FormFieldInt":
		res = GetNamedStructInstances(stage.FormFieldInts, stage.FormFieldInt_stagedOrder)
	case "FormFieldSelect":
		res = GetNamedStructInstances(stage.FormFieldSelects, stage.FormFieldSelect_stagedOrder)
	case "FormFieldString":
		res = GetNamedStructInstances(stage.FormFieldStrings, stage.FormFieldString_stagedOrder)
	case "FormFieldTime":
		res = GetNamedStructInstances(stage.FormFieldTimes, stage.FormFieldTime_stagedOrder)
	case "FormGroup":
		res = GetNamedStructInstances(stage.FormGroups, stage.FormGroup_stagedOrder)
	case "FormSortAssocButton":
		res = GetNamedStructInstances(stage.FormSortAssocButtons, stage.FormSortAssocButton_stagedOrder)
	case "Option":
		res = GetNamedStructInstances(stage.Options, stage.Option_stagedOrder)
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
	return "github.com/fullstack-lang/gong/lib/form/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return form_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return form_go.GoDiagramsDir
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
	CommitCheckBox(checkbox *CheckBox)
	CheckoutCheckBox(checkbox *CheckBox)
	CommitFormDiv(formdiv *FormDiv)
	CheckoutFormDiv(formdiv *FormDiv)
	CommitFormEditAssocButton(formeditassocbutton *FormEditAssocButton)
	CheckoutFormEditAssocButton(formeditassocbutton *FormEditAssocButton)
	CommitFormField(formfield *FormField)
	CheckoutFormField(formfield *FormField)
	CommitFormFieldDate(formfielddate *FormFieldDate)
	CheckoutFormFieldDate(formfielddate *FormFieldDate)
	CommitFormFieldDateTime(formfielddatetime *FormFieldDateTime)
	CheckoutFormFieldDateTime(formfielddatetime *FormFieldDateTime)
	CommitFormFieldFloat64(formfieldfloat64 *FormFieldFloat64)
	CheckoutFormFieldFloat64(formfieldfloat64 *FormFieldFloat64)
	CommitFormFieldInt(formfieldint *FormFieldInt)
	CheckoutFormFieldInt(formfieldint *FormFieldInt)
	CommitFormFieldSelect(formfieldselect *FormFieldSelect)
	CheckoutFormFieldSelect(formfieldselect *FormFieldSelect)
	CommitFormFieldString(formfieldstring *FormFieldString)
	CheckoutFormFieldString(formfieldstring *FormFieldString)
	CommitFormFieldTime(formfieldtime *FormFieldTime)
	CheckoutFormFieldTime(formfieldtime *FormFieldTime)
	CommitFormGroup(formgroup *FormGroup)
	CheckoutFormGroup(formgroup *FormGroup)
	CommitFormSortAssocButton(formsortassocbutton *FormSortAssocButton)
	CheckoutFormSortAssocButton(formsortassocbutton *FormSortAssocButton)
	CommitOption(option *Option)
	CheckoutOption(option *Option)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

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
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
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

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "CheckBox"},
			{name: "FormDiv"},
			{name: "FormEditAssocButton"},
			{name: "FormField"},
			{name: "FormFieldDate"},
			{name: "FormFieldDateTime"},
			{name: "FormFieldFloat64"},
			{name: "FormFieldInt"},
			{name: "FormFieldSelect"},
			{name: "FormFieldString"},
			{name: "FormFieldTime"},
			{name: "FormGroup"},
			{name: "FormSortAssocButton"},
			{name: "Option"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *CheckBox:
		return stage.CheckBox_stagedOrder[instance]
	case *FormDiv:
		return stage.FormDiv_stagedOrder[instance]
	case *FormEditAssocButton:
		return stage.FormEditAssocButton_stagedOrder[instance]
	case *FormField:
		return stage.FormField_stagedOrder[instance]
	case *FormFieldDate:
		return stage.FormFieldDate_stagedOrder[instance]
	case *FormFieldDateTime:
		return stage.FormFieldDateTime_stagedOrder[instance]
	case *FormFieldFloat64:
		return stage.FormFieldFloat64_stagedOrder[instance]
	case *FormFieldInt:
		return stage.FormFieldInt_stagedOrder[instance]
	case *FormFieldSelect:
		return stage.FormFieldSelect_stagedOrder[instance]
	case *FormFieldString:
		return stage.FormFieldString_stagedOrder[instance]
	case *FormFieldTime:
		return stage.FormFieldTime_stagedOrder[instance]
	case *FormGroup:
		return stage.FormGroup_stagedOrder[instance]
	case *FormSortAssocButton:
		return stage.FormSortAssocButton_stagedOrder[instance]
	case *Option:
		return stage.Option_stagedOrder[instance]
	default:
		return 0 // should not happen
	}
}

func GongGetInstanceFromOrder[Type PointerToGongstruct](stage *Stage, order uint) (res Type) {
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

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *CheckBox:
		return stage.CheckBox_stagedOrder[instance]
	case *FormDiv:
		return stage.FormDiv_stagedOrder[instance]
	case *FormEditAssocButton:
		return stage.FormEditAssocButton_stagedOrder[instance]
	case *FormField:
		return stage.FormField_stagedOrder[instance]
	case *FormFieldDate:
		return stage.FormFieldDate_stagedOrder[instance]
	case *FormFieldDateTime:
		return stage.FormFieldDateTime_stagedOrder[instance]
	case *FormFieldFloat64:
		return stage.FormFieldFloat64_stagedOrder[instance]
	case *FormFieldInt:
		return stage.FormFieldInt_stagedOrder[instance]
	case *FormFieldSelect:
		return stage.FormFieldSelect_stagedOrder[instance]
	case *FormFieldString:
		return stage.FormFieldString_stagedOrder[instance]
	case *FormFieldTime:
		return stage.FormFieldTime_stagedOrder[instance]
	case *FormGroup:
		return stage.FormGroup_stagedOrder[instance]
	case *FormSortAssocButton:
		return stage.FormSortAssocButton_stagedOrder[instance]
	case *Option:
		return stage.Option_stagedOrder[instance]
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
	if _, ok := stage.CheckBoxs[checkbox]; !ok {
		stage.CheckBoxs[checkbox] = struct{}{}
		stage.CheckBox_stagedOrder[checkbox] = stage.CheckBoxOrder
		stage.CheckBox_orderStaged[stage.CheckBoxOrder] = checkbox
		stage.CheckBoxOrder++
	}
	stage.CheckBoxs_mapString[checkbox.Name] = checkbox

	return checkbox
}

// StagePreserveOrder puts checkbox to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CheckBoxOrder
// - update stage.CheckBoxOrder accordingly
func (checkbox *CheckBox) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.CheckBoxs[checkbox]; !ok {
		stage.CheckBoxs[checkbox] = struct{}{}

		if order > stage.CheckBoxOrder {
			stage.CheckBoxOrder = order
		}
		stage.CheckBox_stagedOrder[checkbox] = order
		stage.CheckBox_orderStaged[order] = checkbox
		stage.CheckBoxOrder++
	}
	stage.CheckBoxs_mapString[checkbox.Name] = checkbox
}

// Unstage removes checkbox off the model stage
func (checkbox *CheckBox) Unstage(stage *Stage) *CheckBox {
	delete(stage.CheckBoxs, checkbox)
	// issue1150
	// delete(stage.CheckBox_stagedOrder, checkbox)
	delete(stage.CheckBoxs_mapString, checkbox.Name)

	return checkbox
}

// UnstageVoid removes checkbox off the model stage
func (checkbox *CheckBox) UnstageVoid(stage *Stage) {
	delete(stage.CheckBoxs, checkbox)
	// issue1150
	// delete(stage.CheckBox_stagedOrder, checkbox)
	delete(stage.CheckBoxs_mapString, checkbox.Name)
}

// commit checkbox to the back repo (if it is already staged)
func (checkbox *CheckBox) Commit(stage *Stage) *CheckBox {
	if _, ok := stage.CheckBoxs[checkbox]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCheckBox(checkbox)
		}
	}
	return checkbox
}

func (checkbox *CheckBox) CommitVoid(stage *Stage) {
	checkbox.Commit(stage)
}

func (checkbox *CheckBox) StageVoid(stage *Stage) {
	checkbox.Stage(stage)
}

// Checkout checkbox to the back repo (if it is already staged)
func (checkbox *CheckBox) Checkout(stage *Stage) *CheckBox {
	if _, ok := stage.CheckBoxs[checkbox]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCheckBox(checkbox)
		}
	}
	return checkbox
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
	if _, ok := stage.FormDivs[formdiv]; !ok {
		stage.FormDivs[formdiv] = struct{}{}
		stage.FormDiv_stagedOrder[formdiv] = stage.FormDivOrder
		stage.FormDiv_orderStaged[stage.FormDivOrder] = formdiv
		stage.FormDivOrder++
	}
	stage.FormDivs_mapString[formdiv.Name] = formdiv

	return formdiv
}

// StagePreserveOrder puts formdiv to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormDivOrder
// - update stage.FormDivOrder accordingly
func (formdiv *FormDiv) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.FormDivs[formdiv]; !ok {
		stage.FormDivs[formdiv] = struct{}{}

		if order > stage.FormDivOrder {
			stage.FormDivOrder = order
		}
		stage.FormDiv_stagedOrder[formdiv] = order
		stage.FormDiv_orderStaged[order] = formdiv
		stage.FormDivOrder++
	}
	stage.FormDivs_mapString[formdiv.Name] = formdiv
}

// Unstage removes formdiv off the model stage
func (formdiv *FormDiv) Unstage(stage *Stage) *FormDiv {
	delete(stage.FormDivs, formdiv)
	// issue1150
	// delete(stage.FormDiv_stagedOrder, formdiv)
	delete(stage.FormDivs_mapString, formdiv.Name)

	return formdiv
}

// UnstageVoid removes formdiv off the model stage
func (formdiv *FormDiv) UnstageVoid(stage *Stage) {
	delete(stage.FormDivs, formdiv)
	// issue1150
	// delete(stage.FormDiv_stagedOrder, formdiv)
	delete(stage.FormDivs_mapString, formdiv.Name)
}

// commit formdiv to the back repo (if it is already staged)
func (formdiv *FormDiv) Commit(stage *Stage) *FormDiv {
	if _, ok := stage.FormDivs[formdiv]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormDiv(formdiv)
		}
	}
	return formdiv
}

func (formdiv *FormDiv) CommitVoid(stage *Stage) {
	formdiv.Commit(stage)
}

func (formdiv *FormDiv) StageVoid(stage *Stage) {
	formdiv.Stage(stage)
}

// Checkout formdiv to the back repo (if it is already staged)
func (formdiv *FormDiv) Checkout(stage *Stage) *FormDiv {
	if _, ok := stage.FormDivs[formdiv]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormDiv(formdiv)
		}
	}
	return formdiv
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
	if _, ok := stage.FormEditAssocButtons[formeditassocbutton]; !ok {
		stage.FormEditAssocButtons[formeditassocbutton] = struct{}{}
		stage.FormEditAssocButton_stagedOrder[formeditassocbutton] = stage.FormEditAssocButtonOrder
		stage.FormEditAssocButton_orderStaged[stage.FormEditAssocButtonOrder] = formeditassocbutton
		stage.FormEditAssocButtonOrder++
	}
	stage.FormEditAssocButtons_mapString[formeditassocbutton.Name] = formeditassocbutton

	return formeditassocbutton
}

// StagePreserveOrder puts formeditassocbutton to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormEditAssocButtonOrder
// - update stage.FormEditAssocButtonOrder accordingly
func (formeditassocbutton *FormEditAssocButton) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.FormEditAssocButtons[formeditassocbutton]; !ok {
		stage.FormEditAssocButtons[formeditassocbutton] = struct{}{}

		if order > stage.FormEditAssocButtonOrder {
			stage.FormEditAssocButtonOrder = order
		}
		stage.FormEditAssocButton_stagedOrder[formeditassocbutton] = order
		stage.FormEditAssocButton_orderStaged[order] = formeditassocbutton
		stage.FormEditAssocButtonOrder++
	}
	stage.FormEditAssocButtons_mapString[formeditassocbutton.Name] = formeditassocbutton
}

// Unstage removes formeditassocbutton off the model stage
func (formeditassocbutton *FormEditAssocButton) Unstage(stage *Stage) *FormEditAssocButton {
	delete(stage.FormEditAssocButtons, formeditassocbutton)
	// issue1150
	// delete(stage.FormEditAssocButton_stagedOrder, formeditassocbutton)
	delete(stage.FormEditAssocButtons_mapString, formeditassocbutton.Name)

	return formeditassocbutton
}

// UnstageVoid removes formeditassocbutton off the model stage
func (formeditassocbutton *FormEditAssocButton) UnstageVoid(stage *Stage) {
	delete(stage.FormEditAssocButtons, formeditassocbutton)
	// issue1150
	// delete(stage.FormEditAssocButton_stagedOrder, formeditassocbutton)
	delete(stage.FormEditAssocButtons_mapString, formeditassocbutton.Name)
}

// commit formeditassocbutton to the back repo (if it is already staged)
func (formeditassocbutton *FormEditAssocButton) Commit(stage *Stage) *FormEditAssocButton {
	if _, ok := stage.FormEditAssocButtons[formeditassocbutton]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormEditAssocButton(formeditassocbutton)
		}
	}
	return formeditassocbutton
}

func (formeditassocbutton *FormEditAssocButton) CommitVoid(stage *Stage) {
	formeditassocbutton.Commit(stage)
}

func (formeditassocbutton *FormEditAssocButton) StageVoid(stage *Stage) {
	formeditassocbutton.Stage(stage)
}

// Checkout formeditassocbutton to the back repo (if it is already staged)
func (formeditassocbutton *FormEditAssocButton) Checkout(stage *Stage) *FormEditAssocButton {
	if _, ok := stage.FormEditAssocButtons[formeditassocbutton]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormEditAssocButton(formeditassocbutton)
		}
	}
	return formeditassocbutton
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
	if _, ok := stage.FormFields[formfield]; !ok {
		stage.FormFields[formfield] = struct{}{}
		stage.FormField_stagedOrder[formfield] = stage.FormFieldOrder
		stage.FormField_orderStaged[stage.FormFieldOrder] = formfield
		stage.FormFieldOrder++
	}
	stage.FormFields_mapString[formfield.Name] = formfield

	return formfield
}

// StagePreserveOrder puts formfield to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormFieldOrder
// - update stage.FormFieldOrder accordingly
func (formfield *FormField) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.FormFields[formfield]; !ok {
		stage.FormFields[formfield] = struct{}{}

		if order > stage.FormFieldOrder {
			stage.FormFieldOrder = order
		}
		stage.FormField_stagedOrder[formfield] = order
		stage.FormField_orderStaged[order] = formfield
		stage.FormFieldOrder++
	}
	stage.FormFields_mapString[formfield.Name] = formfield
}

// Unstage removes formfield off the model stage
func (formfield *FormField) Unstage(stage *Stage) *FormField {
	delete(stage.FormFields, formfield)
	// issue1150
	// delete(stage.FormField_stagedOrder, formfield)
	delete(stage.FormFields_mapString, formfield.Name)

	return formfield
}

// UnstageVoid removes formfield off the model stage
func (formfield *FormField) UnstageVoid(stage *Stage) {
	delete(stage.FormFields, formfield)
	// issue1150
	// delete(stage.FormField_stagedOrder, formfield)
	delete(stage.FormFields_mapString, formfield.Name)
}

// commit formfield to the back repo (if it is already staged)
func (formfield *FormField) Commit(stage *Stage) *FormField {
	if _, ok := stage.FormFields[formfield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormField(formfield)
		}
	}
	return formfield
}

func (formfield *FormField) CommitVoid(stage *Stage) {
	formfield.Commit(stage)
}

func (formfield *FormField) StageVoid(stage *Stage) {
	formfield.Stage(stage)
}

// Checkout formfield to the back repo (if it is already staged)
func (formfield *FormField) Checkout(stage *Stage) *FormField {
	if _, ok := stage.FormFields[formfield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormField(formfield)
		}
	}
	return formfield
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
	if _, ok := stage.FormFieldDates[formfielddate]; !ok {
		stage.FormFieldDates[formfielddate] = struct{}{}
		stage.FormFieldDate_stagedOrder[formfielddate] = stage.FormFieldDateOrder
		stage.FormFieldDate_orderStaged[stage.FormFieldDateOrder] = formfielddate
		stage.FormFieldDateOrder++
	}
	stage.FormFieldDates_mapString[formfielddate.Name] = formfielddate

	return formfielddate
}

// StagePreserveOrder puts formfielddate to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormFieldDateOrder
// - update stage.FormFieldDateOrder accordingly
func (formfielddate *FormFieldDate) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.FormFieldDates[formfielddate]; !ok {
		stage.FormFieldDates[formfielddate] = struct{}{}

		if order > stage.FormFieldDateOrder {
			stage.FormFieldDateOrder = order
		}
		stage.FormFieldDate_stagedOrder[formfielddate] = order
		stage.FormFieldDate_orderStaged[order] = formfielddate
		stage.FormFieldDateOrder++
	}
	stage.FormFieldDates_mapString[formfielddate.Name] = formfielddate
}

// Unstage removes formfielddate off the model stage
func (formfielddate *FormFieldDate) Unstage(stage *Stage) *FormFieldDate {
	delete(stage.FormFieldDates, formfielddate)
	// issue1150
	// delete(stage.FormFieldDate_stagedOrder, formfielddate)
	delete(stage.FormFieldDates_mapString, formfielddate.Name)

	return formfielddate
}

// UnstageVoid removes formfielddate off the model stage
func (formfielddate *FormFieldDate) UnstageVoid(stage *Stage) {
	delete(stage.FormFieldDates, formfielddate)
	// issue1150
	// delete(stage.FormFieldDate_stagedOrder, formfielddate)
	delete(stage.FormFieldDates_mapString, formfielddate.Name)
}

// commit formfielddate to the back repo (if it is already staged)
func (formfielddate *FormFieldDate) Commit(stage *Stage) *FormFieldDate {
	if _, ok := stage.FormFieldDates[formfielddate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormFieldDate(formfielddate)
		}
	}
	return formfielddate
}

func (formfielddate *FormFieldDate) CommitVoid(stage *Stage) {
	formfielddate.Commit(stage)
}

func (formfielddate *FormFieldDate) StageVoid(stage *Stage) {
	formfielddate.Stage(stage)
}

// Checkout formfielddate to the back repo (if it is already staged)
func (formfielddate *FormFieldDate) Checkout(stage *Stage) *FormFieldDate {
	if _, ok := stage.FormFieldDates[formfielddate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormFieldDate(formfielddate)
		}
	}
	return formfielddate
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
	if _, ok := stage.FormFieldDateTimes[formfielddatetime]; !ok {
		stage.FormFieldDateTimes[formfielddatetime] = struct{}{}
		stage.FormFieldDateTime_stagedOrder[formfielddatetime] = stage.FormFieldDateTimeOrder
		stage.FormFieldDateTime_orderStaged[stage.FormFieldDateTimeOrder] = formfielddatetime
		stage.FormFieldDateTimeOrder++
	}
	stage.FormFieldDateTimes_mapString[formfielddatetime.Name] = formfielddatetime

	return formfielddatetime
}

// StagePreserveOrder puts formfielddatetime to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormFieldDateTimeOrder
// - update stage.FormFieldDateTimeOrder accordingly
func (formfielddatetime *FormFieldDateTime) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.FormFieldDateTimes[formfielddatetime]; !ok {
		stage.FormFieldDateTimes[formfielddatetime] = struct{}{}

		if order > stage.FormFieldDateTimeOrder {
			stage.FormFieldDateTimeOrder = order
		}
		stage.FormFieldDateTime_stagedOrder[formfielddatetime] = order
		stage.FormFieldDateTime_orderStaged[order] = formfielddatetime
		stage.FormFieldDateTimeOrder++
	}
	stage.FormFieldDateTimes_mapString[formfielddatetime.Name] = formfielddatetime
}

// Unstage removes formfielddatetime off the model stage
func (formfielddatetime *FormFieldDateTime) Unstage(stage *Stage) *FormFieldDateTime {
	delete(stage.FormFieldDateTimes, formfielddatetime)
	// issue1150
	// delete(stage.FormFieldDateTime_stagedOrder, formfielddatetime)
	delete(stage.FormFieldDateTimes_mapString, formfielddatetime.Name)

	return formfielddatetime
}

// UnstageVoid removes formfielddatetime off the model stage
func (formfielddatetime *FormFieldDateTime) UnstageVoid(stage *Stage) {
	delete(stage.FormFieldDateTimes, formfielddatetime)
	// issue1150
	// delete(stage.FormFieldDateTime_stagedOrder, formfielddatetime)
	delete(stage.FormFieldDateTimes_mapString, formfielddatetime.Name)
}

// commit formfielddatetime to the back repo (if it is already staged)
func (formfielddatetime *FormFieldDateTime) Commit(stage *Stage) *FormFieldDateTime {
	if _, ok := stage.FormFieldDateTimes[formfielddatetime]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormFieldDateTime(formfielddatetime)
		}
	}
	return formfielddatetime
}

func (formfielddatetime *FormFieldDateTime) CommitVoid(stage *Stage) {
	formfielddatetime.Commit(stage)
}

func (formfielddatetime *FormFieldDateTime) StageVoid(stage *Stage) {
	formfielddatetime.Stage(stage)
}

// Checkout formfielddatetime to the back repo (if it is already staged)
func (formfielddatetime *FormFieldDateTime) Checkout(stage *Stage) *FormFieldDateTime {
	if _, ok := stage.FormFieldDateTimes[formfielddatetime]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormFieldDateTime(formfielddatetime)
		}
	}
	return formfielddatetime
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
	if _, ok := stage.FormFieldFloat64s[formfieldfloat64]; !ok {
		stage.FormFieldFloat64s[formfieldfloat64] = struct{}{}
		stage.FormFieldFloat64_stagedOrder[formfieldfloat64] = stage.FormFieldFloat64Order
		stage.FormFieldFloat64_orderStaged[stage.FormFieldFloat64Order] = formfieldfloat64
		stage.FormFieldFloat64Order++
	}
	stage.FormFieldFloat64s_mapString[formfieldfloat64.Name] = formfieldfloat64

	return formfieldfloat64
}

// StagePreserveOrder puts formfieldfloat64 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormFieldFloat64Order
// - update stage.FormFieldFloat64Order accordingly
func (formfieldfloat64 *FormFieldFloat64) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.FormFieldFloat64s[formfieldfloat64]; !ok {
		stage.FormFieldFloat64s[formfieldfloat64] = struct{}{}

		if order > stage.FormFieldFloat64Order {
			stage.FormFieldFloat64Order = order
		}
		stage.FormFieldFloat64_stagedOrder[formfieldfloat64] = order
		stage.FormFieldFloat64_orderStaged[order] = formfieldfloat64
		stage.FormFieldFloat64Order++
	}
	stage.FormFieldFloat64s_mapString[formfieldfloat64.Name] = formfieldfloat64
}

// Unstage removes formfieldfloat64 off the model stage
func (formfieldfloat64 *FormFieldFloat64) Unstage(stage *Stage) *FormFieldFloat64 {
	delete(stage.FormFieldFloat64s, formfieldfloat64)
	// issue1150
	// delete(stage.FormFieldFloat64_stagedOrder, formfieldfloat64)
	delete(stage.FormFieldFloat64s_mapString, formfieldfloat64.Name)

	return formfieldfloat64
}

// UnstageVoid removes formfieldfloat64 off the model stage
func (formfieldfloat64 *FormFieldFloat64) UnstageVoid(stage *Stage) {
	delete(stage.FormFieldFloat64s, formfieldfloat64)
	// issue1150
	// delete(stage.FormFieldFloat64_stagedOrder, formfieldfloat64)
	delete(stage.FormFieldFloat64s_mapString, formfieldfloat64.Name)
}

// commit formfieldfloat64 to the back repo (if it is already staged)
func (formfieldfloat64 *FormFieldFloat64) Commit(stage *Stage) *FormFieldFloat64 {
	if _, ok := stage.FormFieldFloat64s[formfieldfloat64]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormFieldFloat64(formfieldfloat64)
		}
	}
	return formfieldfloat64
}

func (formfieldfloat64 *FormFieldFloat64) CommitVoid(stage *Stage) {
	formfieldfloat64.Commit(stage)
}

func (formfieldfloat64 *FormFieldFloat64) StageVoid(stage *Stage) {
	formfieldfloat64.Stage(stage)
}

// Checkout formfieldfloat64 to the back repo (if it is already staged)
func (formfieldfloat64 *FormFieldFloat64) Checkout(stage *Stage) *FormFieldFloat64 {
	if _, ok := stage.FormFieldFloat64s[formfieldfloat64]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormFieldFloat64(formfieldfloat64)
		}
	}
	return formfieldfloat64
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
	if _, ok := stage.FormFieldInts[formfieldint]; !ok {
		stage.FormFieldInts[formfieldint] = struct{}{}
		stage.FormFieldInt_stagedOrder[formfieldint] = stage.FormFieldIntOrder
		stage.FormFieldInt_orderStaged[stage.FormFieldIntOrder] = formfieldint
		stage.FormFieldIntOrder++
	}
	stage.FormFieldInts_mapString[formfieldint.Name] = formfieldint

	return formfieldint
}

// StagePreserveOrder puts formfieldint to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormFieldIntOrder
// - update stage.FormFieldIntOrder accordingly
func (formfieldint *FormFieldInt) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.FormFieldInts[formfieldint]; !ok {
		stage.FormFieldInts[formfieldint] = struct{}{}

		if order > stage.FormFieldIntOrder {
			stage.FormFieldIntOrder = order
		}
		stage.FormFieldInt_stagedOrder[formfieldint] = order
		stage.FormFieldInt_orderStaged[order] = formfieldint
		stage.FormFieldIntOrder++
	}
	stage.FormFieldInts_mapString[formfieldint.Name] = formfieldint
}

// Unstage removes formfieldint off the model stage
func (formfieldint *FormFieldInt) Unstage(stage *Stage) *FormFieldInt {
	delete(stage.FormFieldInts, formfieldint)
	// issue1150
	// delete(stage.FormFieldInt_stagedOrder, formfieldint)
	delete(stage.FormFieldInts_mapString, formfieldint.Name)

	return formfieldint
}

// UnstageVoid removes formfieldint off the model stage
func (formfieldint *FormFieldInt) UnstageVoid(stage *Stage) {
	delete(stage.FormFieldInts, formfieldint)
	// issue1150
	// delete(stage.FormFieldInt_stagedOrder, formfieldint)
	delete(stage.FormFieldInts_mapString, formfieldint.Name)
}

// commit formfieldint to the back repo (if it is already staged)
func (formfieldint *FormFieldInt) Commit(stage *Stage) *FormFieldInt {
	if _, ok := stage.FormFieldInts[formfieldint]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormFieldInt(formfieldint)
		}
	}
	return formfieldint
}

func (formfieldint *FormFieldInt) CommitVoid(stage *Stage) {
	formfieldint.Commit(stage)
}

func (formfieldint *FormFieldInt) StageVoid(stage *Stage) {
	formfieldint.Stage(stage)
}

// Checkout formfieldint to the back repo (if it is already staged)
func (formfieldint *FormFieldInt) Checkout(stage *Stage) *FormFieldInt {
	if _, ok := stage.FormFieldInts[formfieldint]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormFieldInt(formfieldint)
		}
	}
	return formfieldint
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
	if _, ok := stage.FormFieldSelects[formfieldselect]; !ok {
		stage.FormFieldSelects[formfieldselect] = struct{}{}
		stage.FormFieldSelect_stagedOrder[formfieldselect] = stage.FormFieldSelectOrder
		stage.FormFieldSelect_orderStaged[stage.FormFieldSelectOrder] = formfieldselect
		stage.FormFieldSelectOrder++
	}
	stage.FormFieldSelects_mapString[formfieldselect.Name] = formfieldselect

	return formfieldselect
}

// StagePreserveOrder puts formfieldselect to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormFieldSelectOrder
// - update stage.FormFieldSelectOrder accordingly
func (formfieldselect *FormFieldSelect) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.FormFieldSelects[formfieldselect]; !ok {
		stage.FormFieldSelects[formfieldselect] = struct{}{}

		if order > stage.FormFieldSelectOrder {
			stage.FormFieldSelectOrder = order
		}
		stage.FormFieldSelect_stagedOrder[formfieldselect] = order
		stage.FormFieldSelect_orderStaged[order] = formfieldselect
		stage.FormFieldSelectOrder++
	}
	stage.FormFieldSelects_mapString[formfieldselect.Name] = formfieldselect
}

// Unstage removes formfieldselect off the model stage
func (formfieldselect *FormFieldSelect) Unstage(stage *Stage) *FormFieldSelect {
	delete(stage.FormFieldSelects, formfieldselect)
	// issue1150
	// delete(stage.FormFieldSelect_stagedOrder, formfieldselect)
	delete(stage.FormFieldSelects_mapString, formfieldselect.Name)

	return formfieldselect
}

// UnstageVoid removes formfieldselect off the model stage
func (formfieldselect *FormFieldSelect) UnstageVoid(stage *Stage) {
	delete(stage.FormFieldSelects, formfieldselect)
	// issue1150
	// delete(stage.FormFieldSelect_stagedOrder, formfieldselect)
	delete(stage.FormFieldSelects_mapString, formfieldselect.Name)
}

// commit formfieldselect to the back repo (if it is already staged)
func (formfieldselect *FormFieldSelect) Commit(stage *Stage) *FormFieldSelect {
	if _, ok := stage.FormFieldSelects[formfieldselect]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormFieldSelect(formfieldselect)
		}
	}
	return formfieldselect
}

func (formfieldselect *FormFieldSelect) CommitVoid(stage *Stage) {
	formfieldselect.Commit(stage)
}

func (formfieldselect *FormFieldSelect) StageVoid(stage *Stage) {
	formfieldselect.Stage(stage)
}

// Checkout formfieldselect to the back repo (if it is already staged)
func (formfieldselect *FormFieldSelect) Checkout(stage *Stage) *FormFieldSelect {
	if _, ok := stage.FormFieldSelects[formfieldselect]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormFieldSelect(formfieldselect)
		}
	}
	return formfieldselect
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
	if _, ok := stage.FormFieldStrings[formfieldstring]; !ok {
		stage.FormFieldStrings[formfieldstring] = struct{}{}
		stage.FormFieldString_stagedOrder[formfieldstring] = stage.FormFieldStringOrder
		stage.FormFieldString_orderStaged[stage.FormFieldStringOrder] = formfieldstring
		stage.FormFieldStringOrder++
	}
	stage.FormFieldStrings_mapString[formfieldstring.Name] = formfieldstring

	return formfieldstring
}

// StagePreserveOrder puts formfieldstring to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormFieldStringOrder
// - update stage.FormFieldStringOrder accordingly
func (formfieldstring *FormFieldString) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.FormFieldStrings[formfieldstring]; !ok {
		stage.FormFieldStrings[formfieldstring] = struct{}{}

		if order > stage.FormFieldStringOrder {
			stage.FormFieldStringOrder = order
		}
		stage.FormFieldString_stagedOrder[formfieldstring] = order
		stage.FormFieldString_orderStaged[order] = formfieldstring
		stage.FormFieldStringOrder++
	}
	stage.FormFieldStrings_mapString[formfieldstring.Name] = formfieldstring
}

// Unstage removes formfieldstring off the model stage
func (formfieldstring *FormFieldString) Unstage(stage *Stage) *FormFieldString {
	delete(stage.FormFieldStrings, formfieldstring)
	// issue1150
	// delete(stage.FormFieldString_stagedOrder, formfieldstring)
	delete(stage.FormFieldStrings_mapString, formfieldstring.Name)

	return formfieldstring
}

// UnstageVoid removes formfieldstring off the model stage
func (formfieldstring *FormFieldString) UnstageVoid(stage *Stage) {
	delete(stage.FormFieldStrings, formfieldstring)
	// issue1150
	// delete(stage.FormFieldString_stagedOrder, formfieldstring)
	delete(stage.FormFieldStrings_mapString, formfieldstring.Name)
}

// commit formfieldstring to the back repo (if it is already staged)
func (formfieldstring *FormFieldString) Commit(stage *Stage) *FormFieldString {
	if _, ok := stage.FormFieldStrings[formfieldstring]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormFieldString(formfieldstring)
		}
	}
	return formfieldstring
}

func (formfieldstring *FormFieldString) CommitVoid(stage *Stage) {
	formfieldstring.Commit(stage)
}

func (formfieldstring *FormFieldString) StageVoid(stage *Stage) {
	formfieldstring.Stage(stage)
}

// Checkout formfieldstring to the back repo (if it is already staged)
func (formfieldstring *FormFieldString) Checkout(stage *Stage) *FormFieldString {
	if _, ok := stage.FormFieldStrings[formfieldstring]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormFieldString(formfieldstring)
		}
	}
	return formfieldstring
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
	if _, ok := stage.FormFieldTimes[formfieldtime]; !ok {
		stage.FormFieldTimes[formfieldtime] = struct{}{}
		stage.FormFieldTime_stagedOrder[formfieldtime] = stage.FormFieldTimeOrder
		stage.FormFieldTime_orderStaged[stage.FormFieldTimeOrder] = formfieldtime
		stage.FormFieldTimeOrder++
	}
	stage.FormFieldTimes_mapString[formfieldtime.Name] = formfieldtime

	return formfieldtime
}

// StagePreserveOrder puts formfieldtime to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormFieldTimeOrder
// - update stage.FormFieldTimeOrder accordingly
func (formfieldtime *FormFieldTime) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.FormFieldTimes[formfieldtime]; !ok {
		stage.FormFieldTimes[formfieldtime] = struct{}{}

		if order > stage.FormFieldTimeOrder {
			stage.FormFieldTimeOrder = order
		}
		stage.FormFieldTime_stagedOrder[formfieldtime] = order
		stage.FormFieldTime_orderStaged[order] = formfieldtime
		stage.FormFieldTimeOrder++
	}
	stage.FormFieldTimes_mapString[formfieldtime.Name] = formfieldtime
}

// Unstage removes formfieldtime off the model stage
func (formfieldtime *FormFieldTime) Unstage(stage *Stage) *FormFieldTime {
	delete(stage.FormFieldTimes, formfieldtime)
	// issue1150
	// delete(stage.FormFieldTime_stagedOrder, formfieldtime)
	delete(stage.FormFieldTimes_mapString, formfieldtime.Name)

	return formfieldtime
}

// UnstageVoid removes formfieldtime off the model stage
func (formfieldtime *FormFieldTime) UnstageVoid(stage *Stage) {
	delete(stage.FormFieldTimes, formfieldtime)
	// issue1150
	// delete(stage.FormFieldTime_stagedOrder, formfieldtime)
	delete(stage.FormFieldTimes_mapString, formfieldtime.Name)
}

// commit formfieldtime to the back repo (if it is already staged)
func (formfieldtime *FormFieldTime) Commit(stage *Stage) *FormFieldTime {
	if _, ok := stage.FormFieldTimes[formfieldtime]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormFieldTime(formfieldtime)
		}
	}
	return formfieldtime
}

func (formfieldtime *FormFieldTime) CommitVoid(stage *Stage) {
	formfieldtime.Commit(stage)
}

func (formfieldtime *FormFieldTime) StageVoid(stage *Stage) {
	formfieldtime.Stage(stage)
}

// Checkout formfieldtime to the back repo (if it is already staged)
func (formfieldtime *FormFieldTime) Checkout(stage *Stage) *FormFieldTime {
	if _, ok := stage.FormFieldTimes[formfieldtime]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormFieldTime(formfieldtime)
		}
	}
	return formfieldtime
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
	if _, ok := stage.FormGroups[formgroup]; !ok {
		stage.FormGroups[formgroup] = struct{}{}
		stage.FormGroup_stagedOrder[formgroup] = stage.FormGroupOrder
		stage.FormGroup_orderStaged[stage.FormGroupOrder] = formgroup
		stage.FormGroupOrder++
	}
	stage.FormGroups_mapString[formgroup.Name] = formgroup

	return formgroup
}

// StagePreserveOrder puts formgroup to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormGroupOrder
// - update stage.FormGroupOrder accordingly
func (formgroup *FormGroup) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.FormGroups[formgroup]; !ok {
		stage.FormGroups[formgroup] = struct{}{}

		if order > stage.FormGroupOrder {
			stage.FormGroupOrder = order
		}
		stage.FormGroup_stagedOrder[formgroup] = order
		stage.FormGroup_orderStaged[order] = formgroup
		stage.FormGroupOrder++
	}
	stage.FormGroups_mapString[formgroup.Name] = formgroup
}

// Unstage removes formgroup off the model stage
func (formgroup *FormGroup) Unstage(stage *Stage) *FormGroup {
	delete(stage.FormGroups, formgroup)
	// issue1150
	// delete(stage.FormGroup_stagedOrder, formgroup)
	delete(stage.FormGroups_mapString, formgroup.Name)

	return formgroup
}

// UnstageVoid removes formgroup off the model stage
func (formgroup *FormGroup) UnstageVoid(stage *Stage) {
	delete(stage.FormGroups, formgroup)
	// issue1150
	// delete(stage.FormGroup_stagedOrder, formgroup)
	delete(stage.FormGroups_mapString, formgroup.Name)
}

// commit formgroup to the back repo (if it is already staged)
func (formgroup *FormGroup) Commit(stage *Stage) *FormGroup {
	if _, ok := stage.FormGroups[formgroup]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormGroup(formgroup)
		}
	}
	return formgroup
}

func (formgroup *FormGroup) CommitVoid(stage *Stage) {
	formgroup.Commit(stage)
}

func (formgroup *FormGroup) StageVoid(stage *Stage) {
	formgroup.Stage(stage)
}

// Checkout formgroup to the back repo (if it is already staged)
func (formgroup *FormGroup) Checkout(stage *Stage) *FormGroup {
	if _, ok := stage.FormGroups[formgroup]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormGroup(formgroup)
		}
	}
	return formgroup
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
	if _, ok := stage.FormSortAssocButtons[formsortassocbutton]; !ok {
		stage.FormSortAssocButtons[formsortassocbutton] = struct{}{}
		stage.FormSortAssocButton_stagedOrder[formsortassocbutton] = stage.FormSortAssocButtonOrder
		stage.FormSortAssocButton_orderStaged[stage.FormSortAssocButtonOrder] = formsortassocbutton
		stage.FormSortAssocButtonOrder++
	}
	stage.FormSortAssocButtons_mapString[formsortassocbutton.Name] = formsortassocbutton

	return formsortassocbutton
}

// StagePreserveOrder puts formsortassocbutton to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormSortAssocButtonOrder
// - update stage.FormSortAssocButtonOrder accordingly
func (formsortassocbutton *FormSortAssocButton) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.FormSortAssocButtons[formsortassocbutton]; !ok {
		stage.FormSortAssocButtons[formsortassocbutton] = struct{}{}

		if order > stage.FormSortAssocButtonOrder {
			stage.FormSortAssocButtonOrder = order
		}
		stage.FormSortAssocButton_stagedOrder[formsortassocbutton] = order
		stage.FormSortAssocButton_orderStaged[order] = formsortassocbutton
		stage.FormSortAssocButtonOrder++
	}
	stage.FormSortAssocButtons_mapString[formsortassocbutton.Name] = formsortassocbutton
}

// Unstage removes formsortassocbutton off the model stage
func (formsortassocbutton *FormSortAssocButton) Unstage(stage *Stage) *FormSortAssocButton {
	delete(stage.FormSortAssocButtons, formsortassocbutton)
	// issue1150
	// delete(stage.FormSortAssocButton_stagedOrder, formsortassocbutton)
	delete(stage.FormSortAssocButtons_mapString, formsortassocbutton.Name)

	return formsortassocbutton
}

// UnstageVoid removes formsortassocbutton off the model stage
func (formsortassocbutton *FormSortAssocButton) UnstageVoid(stage *Stage) {
	delete(stage.FormSortAssocButtons, formsortassocbutton)
	// issue1150
	// delete(stage.FormSortAssocButton_stagedOrder, formsortassocbutton)
	delete(stage.FormSortAssocButtons_mapString, formsortassocbutton.Name)
}

// commit formsortassocbutton to the back repo (if it is already staged)
func (formsortassocbutton *FormSortAssocButton) Commit(stage *Stage) *FormSortAssocButton {
	if _, ok := stage.FormSortAssocButtons[formsortassocbutton]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormSortAssocButton(formsortassocbutton)
		}
	}
	return formsortassocbutton
}

func (formsortassocbutton *FormSortAssocButton) CommitVoid(stage *Stage) {
	formsortassocbutton.Commit(stage)
}

func (formsortassocbutton *FormSortAssocButton) StageVoid(stage *Stage) {
	formsortassocbutton.Stage(stage)
}

// Checkout formsortassocbutton to the back repo (if it is already staged)
func (formsortassocbutton *FormSortAssocButton) Checkout(stage *Stage) *FormSortAssocButton {
	if _, ok := stage.FormSortAssocButtons[formsortassocbutton]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormSortAssocButton(formsortassocbutton)
		}
	}
	return formsortassocbutton
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
	if _, ok := stage.Options[option]; !ok {
		stage.Options[option] = struct{}{}
		stage.Option_stagedOrder[option] = stage.OptionOrder
		stage.Option_orderStaged[stage.OptionOrder] = option
		stage.OptionOrder++
	}
	stage.Options_mapString[option.Name] = option

	return option
}

// StagePreserveOrder puts option to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.OptionOrder
// - update stage.OptionOrder accordingly
func (option *Option) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Options[option]; !ok {
		stage.Options[option] = struct{}{}

		if order > stage.OptionOrder {
			stage.OptionOrder = order
		}
		stage.Option_stagedOrder[option] = order
		stage.Option_orderStaged[order] = option
		stage.OptionOrder++
	}
	stage.Options_mapString[option.Name] = option
}

// Unstage removes option off the model stage
func (option *Option) Unstage(stage *Stage) *Option {
	delete(stage.Options, option)
	// issue1150
	// delete(stage.Option_stagedOrder, option)
	delete(stage.Options_mapString, option.Name)

	return option
}

// UnstageVoid removes option off the model stage
func (option *Option) UnstageVoid(stage *Stage) {
	delete(stage.Options, option)
	// issue1150
	// delete(stage.Option_stagedOrder, option)
	delete(stage.Options_mapString, option.Name)
}

// commit option to the back repo (if it is already staged)
func (option *Option) Commit(stage *Stage) *Option {
	if _, ok := stage.Options[option]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitOption(option)
		}
	}
	return option
}

func (option *Option) CommitVoid(stage *Stage) {
	option.Commit(stage)
}

func (option *Option) StageVoid(stage *Stage) {
	option.Stage(stage)
}

// Checkout option to the back repo (if it is already staged)
func (option *Option) Checkout(stage *Stage) *Option {
	if _, ok := stage.Options[option]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutOption(option)
		}
	}
	return option
}

// for satisfaction of GongStruct interface
func (option *Option) GetName() (res string) {
	return option.Name
}

// for satisfaction of GongStruct interface
func (option *Option) SetName(name string) {
	option.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMCheckBox(CheckBox *CheckBox)
	CreateORMFormDiv(FormDiv *FormDiv)
	CreateORMFormEditAssocButton(FormEditAssocButton *FormEditAssocButton)
	CreateORMFormField(FormField *FormField)
	CreateORMFormFieldDate(FormFieldDate *FormFieldDate)
	CreateORMFormFieldDateTime(FormFieldDateTime *FormFieldDateTime)
	CreateORMFormFieldFloat64(FormFieldFloat64 *FormFieldFloat64)
	CreateORMFormFieldInt(FormFieldInt *FormFieldInt)
	CreateORMFormFieldSelect(FormFieldSelect *FormFieldSelect)
	CreateORMFormFieldString(FormFieldString *FormFieldString)
	CreateORMFormFieldTime(FormFieldTime *FormFieldTime)
	CreateORMFormGroup(FormGroup *FormGroup)
	CreateORMFormSortAssocButton(FormSortAssocButton *FormSortAssocButton)
	CreateORMOption(Option *Option)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMCheckBox(CheckBox *CheckBox)
	DeleteORMFormDiv(FormDiv *FormDiv)
	DeleteORMFormEditAssocButton(FormEditAssocButton *FormEditAssocButton)
	DeleteORMFormField(FormField *FormField)
	DeleteORMFormFieldDate(FormFieldDate *FormFieldDate)
	DeleteORMFormFieldDateTime(FormFieldDateTime *FormFieldDateTime)
	DeleteORMFormFieldFloat64(FormFieldFloat64 *FormFieldFloat64)
	DeleteORMFormFieldInt(FormFieldInt *FormFieldInt)
	DeleteORMFormFieldSelect(FormFieldSelect *FormFieldSelect)
	DeleteORMFormFieldString(FormFieldString *FormFieldString)
	DeleteORMFormFieldTime(FormFieldTime *FormFieldTime)
	DeleteORMFormGroup(FormGroup *FormGroup)
	DeleteORMFormSortAssocButton(FormSortAssocButton *FormSortAssocButton)
	DeleteORMOption(Option *Option)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.CheckBoxs = make(map[*CheckBox]struct{})
	stage.CheckBoxs_mapString = make(map[string]*CheckBox)
	stage.CheckBox_stagedOrder = make(map[*CheckBox]uint)
	stage.CheckBoxOrder = 0

	stage.FormDivs = make(map[*FormDiv]struct{})
	stage.FormDivs_mapString = make(map[string]*FormDiv)
	stage.FormDiv_stagedOrder = make(map[*FormDiv]uint)
	stage.FormDivOrder = 0

	stage.FormEditAssocButtons = make(map[*FormEditAssocButton]struct{})
	stage.FormEditAssocButtons_mapString = make(map[string]*FormEditAssocButton)
	stage.FormEditAssocButton_stagedOrder = make(map[*FormEditAssocButton]uint)
	stage.FormEditAssocButtonOrder = 0

	stage.FormFields = make(map[*FormField]struct{})
	stage.FormFields_mapString = make(map[string]*FormField)
	stage.FormField_stagedOrder = make(map[*FormField]uint)
	stage.FormFieldOrder = 0

	stage.FormFieldDates = make(map[*FormFieldDate]struct{})
	stage.FormFieldDates_mapString = make(map[string]*FormFieldDate)
	stage.FormFieldDate_stagedOrder = make(map[*FormFieldDate]uint)
	stage.FormFieldDateOrder = 0

	stage.FormFieldDateTimes = make(map[*FormFieldDateTime]struct{})
	stage.FormFieldDateTimes_mapString = make(map[string]*FormFieldDateTime)
	stage.FormFieldDateTime_stagedOrder = make(map[*FormFieldDateTime]uint)
	stage.FormFieldDateTimeOrder = 0

	stage.FormFieldFloat64s = make(map[*FormFieldFloat64]struct{})
	stage.FormFieldFloat64s_mapString = make(map[string]*FormFieldFloat64)
	stage.FormFieldFloat64_stagedOrder = make(map[*FormFieldFloat64]uint)
	stage.FormFieldFloat64Order = 0

	stage.FormFieldInts = make(map[*FormFieldInt]struct{})
	stage.FormFieldInts_mapString = make(map[string]*FormFieldInt)
	stage.FormFieldInt_stagedOrder = make(map[*FormFieldInt]uint)
	stage.FormFieldIntOrder = 0

	stage.FormFieldSelects = make(map[*FormFieldSelect]struct{})
	stage.FormFieldSelects_mapString = make(map[string]*FormFieldSelect)
	stage.FormFieldSelect_stagedOrder = make(map[*FormFieldSelect]uint)
	stage.FormFieldSelectOrder = 0

	stage.FormFieldStrings = make(map[*FormFieldString]struct{})
	stage.FormFieldStrings_mapString = make(map[string]*FormFieldString)
	stage.FormFieldString_stagedOrder = make(map[*FormFieldString]uint)
	stage.FormFieldStringOrder = 0

	stage.FormFieldTimes = make(map[*FormFieldTime]struct{})
	stage.FormFieldTimes_mapString = make(map[string]*FormFieldTime)
	stage.FormFieldTime_stagedOrder = make(map[*FormFieldTime]uint)
	stage.FormFieldTimeOrder = 0

	stage.FormGroups = make(map[*FormGroup]struct{})
	stage.FormGroups_mapString = make(map[string]*FormGroup)
	stage.FormGroup_stagedOrder = make(map[*FormGroup]uint)
	stage.FormGroupOrder = 0

	stage.FormSortAssocButtons = make(map[*FormSortAssocButton]struct{})
	stage.FormSortAssocButtons_mapString = make(map[string]*FormSortAssocButton)
	stage.FormSortAssocButton_stagedOrder = make(map[*FormSortAssocButton]uint)
	stage.FormSortAssocButtonOrder = 0

	stage.Options = make(map[*Option]struct{})
	stage.Options_mapString = make(map[string]*Option)
	stage.Option_stagedOrder = make(map[*Option]uint)
	stage.OptionOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.CheckBoxs = nil
	stage.CheckBoxs_mapString = nil

	stage.FormDivs = nil
	stage.FormDivs_mapString = nil

	stage.FormEditAssocButtons = nil
	stage.FormEditAssocButtons_mapString = nil

	stage.FormFields = nil
	stage.FormFields_mapString = nil

	stage.FormFieldDates = nil
	stage.FormFieldDates_mapString = nil

	stage.FormFieldDateTimes = nil
	stage.FormFieldDateTimes_mapString = nil

	stage.FormFieldFloat64s = nil
	stage.FormFieldFloat64s_mapString = nil

	stage.FormFieldInts = nil
	stage.FormFieldInts_mapString = nil

	stage.FormFieldSelects = nil
	stage.FormFieldSelects_mapString = nil

	stage.FormFieldStrings = nil
	stage.FormFieldStrings_mapString = nil

	stage.FormFieldTimes = nil
	stage.FormFieldTimes_mapString = nil

	stage.FormGroups = nil
	stage.FormGroups_mapString = nil

	stage.FormSortAssocButtons = nil
	stage.FormSortAssocButtons_mapString = nil

	stage.Options = nil
	stage.Options_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for checkbox := range stage.CheckBoxs {
		checkbox.Unstage(stage)
	}

	for formdiv := range stage.FormDivs {
		formdiv.Unstage(stage)
	}

	for formeditassocbutton := range stage.FormEditAssocButtons {
		formeditassocbutton.Unstage(stage)
	}

	for formfield := range stage.FormFields {
		formfield.Unstage(stage)
	}

	for formfielddate := range stage.FormFieldDates {
		formfielddate.Unstage(stage)
	}

	for formfielddatetime := range stage.FormFieldDateTimes {
		formfielddatetime.Unstage(stage)
	}

	for formfieldfloat64 := range stage.FormFieldFloat64s {
		formfieldfloat64.Unstage(stage)
	}

	for formfieldint := range stage.FormFieldInts {
		formfieldint.Unstage(stage)
	}

	for formfieldselect := range stage.FormFieldSelects {
		formfieldselect.Unstage(stage)
	}

	for formfieldstring := range stage.FormFieldStrings {
		formfieldstring.Unstage(stage)
	}

	for formfieldtime := range stage.FormFieldTimes {
		formfieldtime.Unstage(stage)
	}

	for formgroup := range stage.FormGroups {
		formgroup.Unstage(stage)
	}

	for formsortassocbutton := range stage.FormSortAssocButtons {
		formsortassocbutton.Unstage(stage)
	}

	for option := range stage.Options {
		option.Unstage(stage)
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
	case map[*CheckBox]any:
		return any(&stage.CheckBoxs).(*Type)
	case map[*FormDiv]any:
		return any(&stage.FormDivs).(*Type)
	case map[*FormEditAssocButton]any:
		return any(&stage.FormEditAssocButtons).(*Type)
	case map[*FormField]any:
		return any(&stage.FormFields).(*Type)
	case map[*FormFieldDate]any:
		return any(&stage.FormFieldDates).(*Type)
	case map[*FormFieldDateTime]any:
		return any(&stage.FormFieldDateTimes).(*Type)
	case map[*FormFieldFloat64]any:
		return any(&stage.FormFieldFloat64s).(*Type)
	case map[*FormFieldInt]any:
		return any(&stage.FormFieldInts).(*Type)
	case map[*FormFieldSelect]any:
		return any(&stage.FormFieldSelects).(*Type)
	case map[*FormFieldString]any:
		return any(&stage.FormFieldStrings).(*Type)
	case map[*FormFieldTime]any:
		return any(&stage.FormFieldTimes).(*Type)
	case map[*FormGroup]any:
		return any(&stage.FormGroups).(*Type)
	case map[*FormSortAssocButton]any:
		return any(&stage.FormSortAssocButtons).(*Type)
	case map[*Option]any:
		return any(&stage.Options).(*Type)
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

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case CheckBox:
		return any(&stage.CheckBoxs).(*map[*Type]struct{})
	case FormDiv:
		return any(&stage.FormDivs).(*map[*Type]struct{})
	case FormEditAssocButton:
		return any(&stage.FormEditAssocButtons).(*map[*Type]struct{})
	case FormField:
		return any(&stage.FormFields).(*map[*Type]struct{})
	case FormFieldDate:
		return any(&stage.FormFieldDates).(*map[*Type]struct{})
	case FormFieldDateTime:
		return any(&stage.FormFieldDateTimes).(*map[*Type]struct{})
	case FormFieldFloat64:
		return any(&stage.FormFieldFloat64s).(*map[*Type]struct{})
	case FormFieldInt:
		return any(&stage.FormFieldInts).(*map[*Type]struct{})
	case FormFieldSelect:
		return any(&stage.FormFieldSelects).(*map[*Type]struct{})
	case FormFieldString:
		return any(&stage.FormFieldStrings).(*map[*Type]struct{})
	case FormFieldTime:
		return any(&stage.FormFieldTimes).(*map[*Type]struct{})
	case FormGroup:
		return any(&stage.FormGroups).(*map[*Type]struct{})
	case FormSortAssocButton:
		return any(&stage.FormSortAssocButtons).(*map[*Type]struct{})
	case Option:
		return any(&stage.Options).(*map[*Type]struct{})
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

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *Stage) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case CheckBox:
		return any(&stage.CheckBoxs_mapString).(*map[string]*Type)
	case FormDiv:
		return any(&stage.FormDivs_mapString).(*map[string]*Type)
	case FormEditAssocButton:
		return any(&stage.FormEditAssocButtons_mapString).(*map[string]*Type)
	case FormField:
		return any(&stage.FormFields_mapString).(*map[string]*Type)
	case FormFieldDate:
		return any(&stage.FormFieldDates_mapString).(*map[string]*Type)
	case FormFieldDateTime:
		return any(&stage.FormFieldDateTimes_mapString).(*map[string]*Type)
	case FormFieldFloat64:
		return any(&stage.FormFieldFloat64s_mapString).(*map[string]*Type)
	case FormFieldInt:
		return any(&stage.FormFieldInts_mapString).(*map[string]*Type)
	case FormFieldSelect:
		return any(&stage.FormFieldSelects_mapString).(*map[string]*Type)
	case FormFieldString:
		return any(&stage.FormFieldStrings_mapString).(*map[string]*Type)
	case FormFieldTime:
		return any(&stage.FormFieldTimes_mapString).(*map[string]*Type)
	case FormGroup:
		return any(&stage.FormGroups_mapString).(*map[string]*Type)
	case FormSortAssocButton:
		return any(&stage.FormSortAssocButtons_mapString).(*map[string]*Type)
	case Option:
		return any(&stage.Options_mapString).(*map[string]*Type)
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
	case CheckBox:
		return any(&CheckBox{
			// Initialisation of associations
		}).(*Type)
	case FormDiv:
		return any(&FormDiv{
			// Initialisation of associations
			// field is initialized with an instance of FormField with the name of the field
			FormFields: []*FormField{{Name: "FormFields"}},
			// field is initialized with an instance of CheckBox with the name of the field
			CheckBoxs: []*CheckBox{{Name: "CheckBoxs"}},
			// field is initialized with an instance of FormEditAssocButton with the name of the field
			FormEditAssocButton: &FormEditAssocButton{Name: "FormEditAssocButton"},
			// field is initialized with an instance of FormSortAssocButton with the name of the field
			FormSortAssocButton: &FormSortAssocButton{Name: "FormSortAssocButton"},
		}).(*Type)
	case FormEditAssocButton:
		return any(&FormEditAssocButton{
			// Initialisation of associations
		}).(*Type)
	case FormField:
		return any(&FormField{
			// Initialisation of associations
			// field is initialized with an instance of FormFieldString with the name of the field
			FormFieldString: &FormFieldString{Name: "FormFieldString"},
			// field is initialized with an instance of FormFieldFloat64 with the name of the field
			FormFieldFloat64: &FormFieldFloat64{Name: "FormFieldFloat64"},
			// field is initialized with an instance of FormFieldInt with the name of the field
			FormFieldInt: &FormFieldInt{Name: "FormFieldInt"},
			// field is initialized with an instance of FormFieldDate with the name of the field
			FormFieldDate: &FormFieldDate{Name: "FormFieldDate"},
			// field is initialized with an instance of FormFieldTime with the name of the field
			FormFieldTime: &FormFieldTime{Name: "FormFieldTime"},
			// field is initialized with an instance of FormFieldDateTime with the name of the field
			FormFieldDateTime: &FormFieldDateTime{Name: "FormFieldDateTime"},
			// field is initialized with an instance of FormFieldSelect with the name of the field
			FormFieldSelect: &FormFieldSelect{Name: "FormFieldSelect"},
		}).(*Type)
	case FormFieldDate:
		return any(&FormFieldDate{
			// Initialisation of associations
		}).(*Type)
	case FormFieldDateTime:
		return any(&FormFieldDateTime{
			// Initialisation of associations
		}).(*Type)
	case FormFieldFloat64:
		return any(&FormFieldFloat64{
			// Initialisation of associations
		}).(*Type)
	case FormFieldInt:
		return any(&FormFieldInt{
			// Initialisation of associations
		}).(*Type)
	case FormFieldSelect:
		return any(&FormFieldSelect{
			// Initialisation of associations
			// field is initialized with an instance of Option with the name of the field
			Value: &Option{Name: "Value"},
			// field is initialized with an instance of Option with the name of the field
			Options: []*Option{{Name: "Options"}},
		}).(*Type)
	case FormFieldString:
		return any(&FormFieldString{
			// Initialisation of associations
		}).(*Type)
	case FormFieldTime:
		return any(&FormFieldTime{
			// Initialisation of associations
		}).(*Type)
	case FormGroup:
		return any(&FormGroup{
			// Initialisation of associations
			// field is initialized with an instance of FormDiv with the name of the field
			FormDivs: []*FormDiv{{Name: "FormDivs"}},
		}).(*Type)
	case FormSortAssocButton:
		return any(&FormSortAssocButton{
			// Initialisation of associations
			// field is initialized with an instance of FormEditAssocButton with the name of the field
			FormEditAssocButton: &FormEditAssocButton{Name: "FormEditAssocButton"},
		}).(*Type)
	case Option:
		return any(&Option{
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

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type GongstructIF]() (res string) {
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

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type GongstructIF]() (res []ReverseField) {
	res = make([]ReverseField, 0)

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

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (checkbox *CheckBox) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		checkbox.Name = value.GetValueString()
	case "Value":
		checkbox.Value = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (formdiv *FormDiv) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		formdiv.Name = value.GetValueString()
	case "FormFields":
		formdiv.FormFields = make([]*FormField, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.FormFields {
					if stage.FormField_stagedOrder[__instance__] == uint(id) {
						formdiv.FormFields = append(formdiv.FormFields, __instance__)
						break
					}
				}
			}
		}
	case "CheckBoxs":
		formdiv.CheckBoxs = make([]*CheckBox, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.CheckBoxs {
					if stage.CheckBox_stagedOrder[__instance__] == uint(id) {
						formdiv.CheckBoxs = append(formdiv.CheckBoxs, __instance__)
						break
					}
				}
			}
		}
	case "FormEditAssocButton":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			formdiv.FormEditAssocButton = nil
			for __instance__ := range stage.FormEditAssocButtons {
				if stage.FormEditAssocButton_stagedOrder[__instance__] == uint(id) {
					formdiv.FormEditAssocButton = __instance__
					break
				}
			}
		}
	case "FormSortAssocButton":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			formdiv.FormSortAssocButton = nil
			for __instance__ := range stage.FormSortAssocButtons {
				if stage.FormSortAssocButton_stagedOrder[__instance__] == uint(id) {
					formdiv.FormSortAssocButton = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (formeditassocbutton *FormEditAssocButton) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		formeditassocbutton.Name = value.GetValueString()
	case "Label":
		formeditassocbutton.Label = value.GetValueString()
	case "AssociationStorage":
		formeditassocbutton.AssociationStorage = value.GetValueString()
	case "HasChanged":
		formeditassocbutton.HasChanged = value.GetValueBool()
	case "IsForSavePurpose":
		formeditassocbutton.IsForSavePurpose = value.GetValueBool()
	case "HasToolTip":
		formeditassocbutton.HasToolTip = value.GetValueBool()
	case "ToolTipText":
		formeditassocbutton.ToolTipText = value.GetValueString()
	case "MatTooltipShowDelay":
		formeditassocbutton.MatTooltipShowDelay = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (formfield *FormField) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		formfield.Name = value.GetValueString()
	case "InputTypeEnum":
		formfield.InputTypeEnum.FromCodeString(value.GetValueString())
	case "Label":
		formfield.Label = value.GetValueString()
	case "Placeholder":
		formfield.Placeholder = value.GetValueString()
	case "FormFieldString":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			formfield.FormFieldString = nil
			for __instance__ := range stage.FormFieldStrings {
				if stage.FormFieldString_stagedOrder[__instance__] == uint(id) {
					formfield.FormFieldString = __instance__
					break
				}
			}
		}
	case "FormFieldFloat64":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			formfield.FormFieldFloat64 = nil
			for __instance__ := range stage.FormFieldFloat64s {
				if stage.FormFieldFloat64_stagedOrder[__instance__] == uint(id) {
					formfield.FormFieldFloat64 = __instance__
					break
				}
			}
		}
	case "FormFieldInt":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			formfield.FormFieldInt = nil
			for __instance__ := range stage.FormFieldInts {
				if stage.FormFieldInt_stagedOrder[__instance__] == uint(id) {
					formfield.FormFieldInt = __instance__
					break
				}
			}
		}
	case "FormFieldDate":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			formfield.FormFieldDate = nil
			for __instance__ := range stage.FormFieldDates {
				if stage.FormFieldDate_stagedOrder[__instance__] == uint(id) {
					formfield.FormFieldDate = __instance__
					break
				}
			}
		}
	case "FormFieldTime":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			formfield.FormFieldTime = nil
			for __instance__ := range stage.FormFieldTimes {
				if stage.FormFieldTime_stagedOrder[__instance__] == uint(id) {
					formfield.FormFieldTime = __instance__
					break
				}
			}
		}
	case "FormFieldDateTime":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			formfield.FormFieldDateTime = nil
			for __instance__ := range stage.FormFieldDateTimes {
				if stage.FormFieldDateTime_stagedOrder[__instance__] == uint(id) {
					formfield.FormFieldDateTime = __instance__
					break
				}
			}
		}
	case "FormFieldSelect":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			formfield.FormFieldSelect = nil
			for __instance__ := range stage.FormFieldSelects {
				if stage.FormFieldSelect_stagedOrder[__instance__] == uint(id) {
					formfield.FormFieldSelect = __instance__
					break
				}
			}
		}
	case "HasBespokeWidth":
		formfield.HasBespokeWidth = value.GetValueBool()
	case "BespokeWidthPx":
		formfield.BespokeWidthPx = int(value.GetValueInt())
	case "HasBespokeHeight":
		formfield.HasBespokeHeight = value.GetValueBool()
	case "BespokeHeightPx":
		formfield.BespokeHeightPx = int(value.GetValueInt())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (formfielddate *FormFieldDate) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		formfielddate.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (formfielddatetime *FormFieldDateTime) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		formfielddatetime.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (formfieldfloat64 *FormFieldFloat64) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		formfieldfloat64.Name = value.GetValueString()
	case "Value":
		formfieldfloat64.Value = value.GetValueFloat()
	case "HasMinValidator":
		formfieldfloat64.HasMinValidator = value.GetValueBool()
	case "MinValue":
		formfieldfloat64.MinValue = value.GetValueFloat()
	case "HasMaxValidator":
		formfieldfloat64.HasMaxValidator = value.GetValueBool()
	case "MaxValue":
		formfieldfloat64.MaxValue = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (formfieldint *FormFieldInt) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		formfieldint.Name = value.GetValueString()
	case "Value":
		formfieldint.Value = int(value.GetValueInt())
	case "HasMinValidator":
		formfieldint.HasMinValidator = value.GetValueBool()
	case "MinValue":
		formfieldint.MinValue = int(value.GetValueInt())
	case "HasMaxValidator":
		formfieldint.HasMaxValidator = value.GetValueBool()
	case "MaxValue":
		formfieldint.MaxValue = int(value.GetValueInt())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (formfieldselect *FormFieldSelect) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		formfieldselect.Name = value.GetValueString()
	case "Value":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			formfieldselect.Value = nil
			for __instance__ := range stage.Options {
				if stage.Option_stagedOrder[__instance__] == uint(id) {
					formfieldselect.Value = __instance__
					break
				}
			}
		}
	case "Options":
		formfieldselect.Options = make([]*Option, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Options {
					if stage.Option_stagedOrder[__instance__] == uint(id) {
						formfieldselect.Options = append(formfieldselect.Options, __instance__)
						break
					}
				}
			}
		}
	case "CanBeEmpty":
		formfieldselect.CanBeEmpty = value.GetValueBool()
	case "PreserveInitialOrder":
		formfieldselect.PreserveInitialOrder = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (formfieldstring *FormFieldString) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		formfieldstring.Name = value.GetValueString()
	case "Value":
		formfieldstring.Value = value.GetValueString()
	case "IsTextArea":
		formfieldstring.IsTextArea = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (formfieldtime *FormFieldTime) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		formfieldtime.Name = value.GetValueString()
	case "Step":
		formfieldtime.Step = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (formgroup *FormGroup) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		formgroup.Name = value.GetValueString()
	case "Label":
		formgroup.Label = value.GetValueString()
	case "FormDivs":
		formgroup.FormDivs = make([]*FormDiv, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.FormDivs {
					if stage.FormDiv_stagedOrder[__instance__] == uint(id) {
						formgroup.FormDivs = append(formgroup.FormDivs, __instance__)
						break
					}
				}
			}
		}
	case "HasSuppressButton":
		formgroup.HasSuppressButton = value.GetValueBool()
	case "HasSuppressButtonBeenPressed":
		formgroup.HasSuppressButtonBeenPressed = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (formsortassocbutton *FormSortAssocButton) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		formsortassocbutton.Name = value.GetValueString()
	case "Label":
		formsortassocbutton.Label = value.GetValueString()
	case "HasToolTip":
		formsortassocbutton.HasToolTip = value.GetValueBool()
	case "ToolTipText":
		formsortassocbutton.ToolTipText = value.GetValueString()
	case "MatTooltipShowDelay":
		formsortassocbutton.MatTooltipShowDelay = value.GetValueString()
	case "FormEditAssocButton":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			formsortassocbutton.FormEditAssocButton = nil
			for __instance__ := range stage.FormEditAssocButtons {
				if stage.FormEditAssocButton_stagedOrder[__instance__] == uint(id) {
					formsortassocbutton.FormEditAssocButton = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (option *Option) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		option.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
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

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	stage.CheckBoxs_mapString = make(map[string]*CheckBox)
	for checkbox := range stage.CheckBoxs {
		stage.CheckBoxs_mapString[checkbox.Name] = checkbox
	}

	stage.FormDivs_mapString = make(map[string]*FormDiv)
	for formdiv := range stage.FormDivs {
		stage.FormDivs_mapString[formdiv.Name] = formdiv
	}

	stage.FormEditAssocButtons_mapString = make(map[string]*FormEditAssocButton)
	for formeditassocbutton := range stage.FormEditAssocButtons {
		stage.FormEditAssocButtons_mapString[formeditassocbutton.Name] = formeditassocbutton
	}

	stage.FormFields_mapString = make(map[string]*FormField)
	for formfield := range stage.FormFields {
		stage.FormFields_mapString[formfield.Name] = formfield
	}

	stage.FormFieldDates_mapString = make(map[string]*FormFieldDate)
	for formfielddate := range stage.FormFieldDates {
		stage.FormFieldDates_mapString[formfielddate.Name] = formfielddate
	}

	stage.FormFieldDateTimes_mapString = make(map[string]*FormFieldDateTime)
	for formfielddatetime := range stage.FormFieldDateTimes {
		stage.FormFieldDateTimes_mapString[formfielddatetime.Name] = formfielddatetime
	}

	stage.FormFieldFloat64s_mapString = make(map[string]*FormFieldFloat64)
	for formfieldfloat64 := range stage.FormFieldFloat64s {
		stage.FormFieldFloat64s_mapString[formfieldfloat64.Name] = formfieldfloat64
	}

	stage.FormFieldInts_mapString = make(map[string]*FormFieldInt)
	for formfieldint := range stage.FormFieldInts {
		stage.FormFieldInts_mapString[formfieldint.Name] = formfieldint
	}

	stage.FormFieldSelects_mapString = make(map[string]*FormFieldSelect)
	for formfieldselect := range stage.FormFieldSelects {
		stage.FormFieldSelects_mapString[formfieldselect.Name] = formfieldselect
	}

	stage.FormFieldStrings_mapString = make(map[string]*FormFieldString)
	for formfieldstring := range stage.FormFieldStrings {
		stage.FormFieldStrings_mapString[formfieldstring.Name] = formfieldstring
	}

	stage.FormFieldTimes_mapString = make(map[string]*FormFieldTime)
	for formfieldtime := range stage.FormFieldTimes {
		stage.FormFieldTimes_mapString[formfieldtime.Name] = formfieldtime
	}

	stage.FormGroups_mapString = make(map[string]*FormGroup)
	for formgroup := range stage.FormGroups {
		stage.FormGroups_mapString[formgroup.Name] = formgroup
	}

	stage.FormSortAssocButtons_mapString = make(map[string]*FormSortAssocButton)
	for formsortassocbutton := range stage.FormSortAssocButtons {
		stage.FormSortAssocButtons_mapString[formsortassocbutton.Name] = formsortassocbutton
	}

	stage.Options_mapString = make(map[string]*Option)
	for option := range stage.Options {
		stage.Options_mapString[option.Name] = option
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
