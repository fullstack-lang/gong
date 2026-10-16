// generated code - do not edit
package attributes

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
	Attributess                map[*Attributes]struct{}
	Attributess_instance       map[*Attributes]*Attributes
	Attributess_mapString      map[string]*Attributes
	AttributesOrder            uint
	Attributes_stagedOrder     map[*Attributes]uint
	Attributes_orderStaged     map[uint]*Attributes
	Attributess_reference      map[*Attributes]*Attributes
	Attributess_referenceOrder map[*Attributes]uint

	// insertion point for slice of pointers maps
	Attributes_Key_reverseMap map[*Key]*Attributes

	Attributes_Time_reverseMap map[*Time]*Attributes

	Attributes_Clef_reverseMap map[*Clef]*Attributes

	Attributes_StaffDetails_reverseMap map[*StaffDetails]*Attributes

	Attributes_Transpose_reverseMap map[*Transpose]*Attributes

	Attributes_ForPart_reverseMap map[*ForPart]*Attributes

	Attributes_Directive_reverseMap map[*AttributesDirective]*Attributes

	Attributes_MeasureStyle_reverseMap map[*MeasureStyle]*Attributes

	OnAfterAttributesCreateCallback GongOnAfterCreateInterface[Attributes]
	OnAfterAttributesUpdateCallback GongOnAfterUpdateInterface[Attributes]
	OnAfterAttributesDeleteCallback GongOnAfterDeleteInterface[Attributes]

	AttributesDirectives                map[*AttributesDirective]struct{}
	AttributesDirectives_instance       map[*AttributesDirective]*AttributesDirective
	AttributesDirectives_mapString      map[string]*AttributesDirective
	AttributesDirectiveOrder            uint
	AttributesDirective_stagedOrder     map[*AttributesDirective]uint
	AttributesDirective_orderStaged     map[uint]*AttributesDirective
	AttributesDirectives_reference      map[*AttributesDirective]*AttributesDirective
	AttributesDirectives_referenceOrder map[*AttributesDirective]uint

	// insertion point for slice of pointers maps
	OnAfterAttributesDirectiveCreateCallback GongOnAfterCreateInterface[AttributesDirective]
	OnAfterAttributesDirectiveUpdateCallback GongOnAfterUpdateInterface[AttributesDirective]
	OnAfterAttributesDirectiveDeleteCallback GongOnAfterDeleteInterface[AttributesDirective]

	BeatRepeats                map[*BeatRepeat]struct{}
	BeatRepeats_instance       map[*BeatRepeat]*BeatRepeat
	BeatRepeats_mapString      map[string]*BeatRepeat
	BeatRepeatOrder            uint
	BeatRepeat_stagedOrder     map[*BeatRepeat]uint
	BeatRepeat_orderStaged     map[uint]*BeatRepeat
	BeatRepeats_reference      map[*BeatRepeat]*BeatRepeat
	BeatRepeats_referenceOrder map[*BeatRepeat]uint

	// insertion point for slice of pointers maps
	OnAfterBeatRepeatCreateCallback GongOnAfterCreateInterface[BeatRepeat]
	OnAfterBeatRepeatUpdateCallback GongOnAfterUpdateInterface[BeatRepeat]
	OnAfterBeatRepeatDeleteCallback GongOnAfterDeleteInterface[BeatRepeat]

	Cancels                map[*Cancel]struct{}
	Cancels_instance       map[*Cancel]*Cancel
	Cancels_mapString      map[string]*Cancel
	CancelOrder            uint
	Cancel_stagedOrder     map[*Cancel]uint
	Cancel_orderStaged     map[uint]*Cancel
	Cancels_reference      map[*Cancel]*Cancel
	Cancels_referenceOrder map[*Cancel]uint

	// insertion point for slice of pointers maps
	OnAfterCancelCreateCallback GongOnAfterCreateInterface[Cancel]
	OnAfterCancelUpdateCallback GongOnAfterUpdateInterface[Cancel]
	OnAfterCancelDeleteCallback GongOnAfterDeleteInterface[Cancel]

	Clefs                map[*Clef]struct{}
	Clefs_instance       map[*Clef]*Clef
	Clefs_mapString      map[string]*Clef
	ClefOrder            uint
	Clef_stagedOrder     map[*Clef]uint
	Clef_orderStaged     map[uint]*Clef
	Clefs_reference      map[*Clef]*Clef
	Clefs_referenceOrder map[*Clef]uint

	// insertion point for slice of pointers maps
	OnAfterClefCreateCallback GongOnAfterCreateInterface[Clef]
	OnAfterClefUpdateCallback GongOnAfterUpdateInterface[Clef]
	OnAfterClefDeleteCallback GongOnAfterDeleteInterface[Clef]

	Doubles                map[*Double]struct{}
	Doubles_instance       map[*Double]*Double
	Doubles_mapString      map[string]*Double
	DoubleOrder            uint
	Double_stagedOrder     map[*Double]uint
	Double_orderStaged     map[uint]*Double
	Doubles_reference      map[*Double]*Double
	Doubles_referenceOrder map[*Double]uint

	// insertion point for slice of pointers maps
	OnAfterDoubleCreateCallback GongOnAfterCreateInterface[Double]
	OnAfterDoubleUpdateCallback GongOnAfterUpdateInterface[Double]
	OnAfterDoubleDeleteCallback GongOnAfterDeleteInterface[Double]

	ForParts                map[*ForPart]struct{}
	ForParts_instance       map[*ForPart]*ForPart
	ForParts_mapString      map[string]*ForPart
	ForPartOrder            uint
	ForPart_stagedOrder     map[*ForPart]uint
	ForPart_orderStaged     map[uint]*ForPart
	ForParts_reference      map[*ForPart]*ForPart
	ForParts_referenceOrder map[*ForPart]uint

	// insertion point for slice of pointers maps
	OnAfterForPartCreateCallback GongOnAfterCreateInterface[ForPart]
	OnAfterForPartUpdateCallback GongOnAfterUpdateInterface[ForPart]
	OnAfterForPartDeleteCallback GongOnAfterDeleteInterface[ForPart]

	Interchangeables                map[*Interchangeable]struct{}
	Interchangeables_instance       map[*Interchangeable]*Interchangeable
	Interchangeables_mapString      map[string]*Interchangeable
	InterchangeableOrder            uint
	Interchangeable_stagedOrder     map[*Interchangeable]uint
	Interchangeable_orderStaged     map[uint]*Interchangeable
	Interchangeables_reference      map[*Interchangeable]*Interchangeable
	Interchangeables_referenceOrder map[*Interchangeable]uint

	// insertion point for slice of pointers maps
	OnAfterInterchangeableCreateCallback GongOnAfterCreateInterface[Interchangeable]
	OnAfterInterchangeableUpdateCallback GongOnAfterUpdateInterface[Interchangeable]
	OnAfterInterchangeableDeleteCallback GongOnAfterDeleteInterface[Interchangeable]

	Keys                map[*Key]struct{}
	Keys_instance       map[*Key]*Key
	Keys_mapString      map[string]*Key
	KeyOrder            uint
	Key_stagedOrder     map[*Key]uint
	Key_orderStaged     map[uint]*Key
	Keys_reference      map[*Key]*Key
	Keys_referenceOrder map[*Key]uint

	// insertion point for slice of pointers maps
	Key_KeyAccidental_reverseMap map[*KeyAccidental]*Key

	Key_KeyOctave_reverseMap map[*KeyOctave]*Key

	OnAfterKeyCreateCallback GongOnAfterCreateInterface[Key]
	OnAfterKeyUpdateCallback GongOnAfterUpdateInterface[Key]
	OnAfterKeyDeleteCallback GongOnAfterDeleteInterface[Key]

	KeyAccidentals                map[*KeyAccidental]struct{}
	KeyAccidentals_instance       map[*KeyAccidental]*KeyAccidental
	KeyAccidentals_mapString      map[string]*KeyAccidental
	KeyAccidentalOrder            uint
	KeyAccidental_stagedOrder     map[*KeyAccidental]uint
	KeyAccidental_orderStaged     map[uint]*KeyAccidental
	KeyAccidentals_reference      map[*KeyAccidental]*KeyAccidental
	KeyAccidentals_referenceOrder map[*KeyAccidental]uint

	// insertion point for slice of pointers maps
	OnAfterKeyAccidentalCreateCallback GongOnAfterCreateInterface[KeyAccidental]
	OnAfterKeyAccidentalUpdateCallback GongOnAfterUpdateInterface[KeyAccidental]
	OnAfterKeyAccidentalDeleteCallback GongOnAfterDeleteInterface[KeyAccidental]

	KeyOctaves                map[*KeyOctave]struct{}
	KeyOctaves_instance       map[*KeyOctave]*KeyOctave
	KeyOctaves_mapString      map[string]*KeyOctave
	KeyOctaveOrder            uint
	KeyOctave_stagedOrder     map[*KeyOctave]uint
	KeyOctave_orderStaged     map[uint]*KeyOctave
	KeyOctaves_reference      map[*KeyOctave]*KeyOctave
	KeyOctaves_referenceOrder map[*KeyOctave]uint

	// insertion point for slice of pointers maps
	OnAfterKeyOctaveCreateCallback GongOnAfterCreateInterface[KeyOctave]
	OnAfterKeyOctaveUpdateCallback GongOnAfterUpdateInterface[KeyOctave]
	OnAfterKeyOctaveDeleteCallback GongOnAfterDeleteInterface[KeyOctave]

	LineDetails                map[*LineDetail]struct{}
	LineDetails_instance       map[*LineDetail]*LineDetail
	LineDetails_mapString      map[string]*LineDetail
	LineDetailOrder            uint
	LineDetail_stagedOrder     map[*LineDetail]uint
	LineDetail_orderStaged     map[uint]*LineDetail
	LineDetails_reference      map[*LineDetail]*LineDetail
	LineDetails_referenceOrder map[*LineDetail]uint

	// insertion point for slice of pointers maps
	OnAfterLineDetailCreateCallback GongOnAfterCreateInterface[LineDetail]
	OnAfterLineDetailUpdateCallback GongOnAfterUpdateInterface[LineDetail]
	OnAfterLineDetailDeleteCallback GongOnAfterDeleteInterface[LineDetail]

	MeasureRepeats                map[*MeasureRepeat]struct{}
	MeasureRepeats_instance       map[*MeasureRepeat]*MeasureRepeat
	MeasureRepeats_mapString      map[string]*MeasureRepeat
	MeasureRepeatOrder            uint
	MeasureRepeat_stagedOrder     map[*MeasureRepeat]uint
	MeasureRepeat_orderStaged     map[uint]*MeasureRepeat
	MeasureRepeats_reference      map[*MeasureRepeat]*MeasureRepeat
	MeasureRepeats_referenceOrder map[*MeasureRepeat]uint

	// insertion point for slice of pointers maps
	OnAfterMeasureRepeatCreateCallback GongOnAfterCreateInterface[MeasureRepeat]
	OnAfterMeasureRepeatUpdateCallback GongOnAfterUpdateInterface[MeasureRepeat]
	OnAfterMeasureRepeatDeleteCallback GongOnAfterDeleteInterface[MeasureRepeat]

	MeasureStyles                map[*MeasureStyle]struct{}
	MeasureStyles_instance       map[*MeasureStyle]*MeasureStyle
	MeasureStyles_mapString      map[string]*MeasureStyle
	MeasureStyleOrder            uint
	MeasureStyle_stagedOrder     map[*MeasureStyle]uint
	MeasureStyle_orderStaged     map[uint]*MeasureStyle
	MeasureStyles_reference      map[*MeasureStyle]*MeasureStyle
	MeasureStyles_referenceOrder map[*MeasureStyle]uint

	// insertion point for slice of pointers maps
	OnAfterMeasureStyleCreateCallback GongOnAfterCreateInterface[MeasureStyle]
	OnAfterMeasureStyleUpdateCallback GongOnAfterUpdateInterface[MeasureStyle]
	OnAfterMeasureStyleDeleteCallback GongOnAfterDeleteInterface[MeasureStyle]

	MultipleRests                map[*MultipleRest]struct{}
	MultipleRests_instance       map[*MultipleRest]*MultipleRest
	MultipleRests_mapString      map[string]*MultipleRest
	MultipleRestOrder            uint
	MultipleRest_stagedOrder     map[*MultipleRest]uint
	MultipleRest_orderStaged     map[uint]*MultipleRest
	MultipleRests_reference      map[*MultipleRest]*MultipleRest
	MultipleRests_referenceOrder map[*MultipleRest]uint

	// insertion point for slice of pointers maps
	OnAfterMultipleRestCreateCallback GongOnAfterCreateInterface[MultipleRest]
	OnAfterMultipleRestUpdateCallback GongOnAfterUpdateInterface[MultipleRest]
	OnAfterMultipleRestDeleteCallback GongOnAfterDeleteInterface[MultipleRest]

	PartClefs                map[*PartClef]struct{}
	PartClefs_instance       map[*PartClef]*PartClef
	PartClefs_mapString      map[string]*PartClef
	PartClefOrder            uint
	PartClef_stagedOrder     map[*PartClef]uint
	PartClef_orderStaged     map[uint]*PartClef
	PartClefs_reference      map[*PartClef]*PartClef
	PartClefs_referenceOrder map[*PartClef]uint

	// insertion point for slice of pointers maps
	OnAfterPartClefCreateCallback GongOnAfterCreateInterface[PartClef]
	OnAfterPartClefUpdateCallback GongOnAfterUpdateInterface[PartClef]
	OnAfterPartClefDeleteCallback GongOnAfterDeleteInterface[PartClef]

	PartSymbols                map[*PartSymbol]struct{}
	PartSymbols_instance       map[*PartSymbol]*PartSymbol
	PartSymbols_mapString      map[string]*PartSymbol
	PartSymbolOrder            uint
	PartSymbol_stagedOrder     map[*PartSymbol]uint
	PartSymbol_orderStaged     map[uint]*PartSymbol
	PartSymbols_reference      map[*PartSymbol]*PartSymbol
	PartSymbols_referenceOrder map[*PartSymbol]uint

	// insertion point for slice of pointers maps
	OnAfterPartSymbolCreateCallback GongOnAfterCreateInterface[PartSymbol]
	OnAfterPartSymbolUpdateCallback GongOnAfterUpdateInterface[PartSymbol]
	OnAfterPartSymbolDeleteCallback GongOnAfterDeleteInterface[PartSymbol]

	PartTransposes                map[*PartTranspose]struct{}
	PartTransposes_instance       map[*PartTranspose]*PartTranspose
	PartTransposes_mapString      map[string]*PartTranspose
	PartTransposeOrder            uint
	PartTranspose_stagedOrder     map[*PartTranspose]uint
	PartTranspose_orderStaged     map[uint]*PartTranspose
	PartTransposes_reference      map[*PartTranspose]*PartTranspose
	PartTransposes_referenceOrder map[*PartTranspose]uint

	// insertion point for slice of pointers maps
	OnAfterPartTransposeCreateCallback GongOnAfterCreateInterface[PartTranspose]
	OnAfterPartTransposeUpdateCallback GongOnAfterUpdateInterface[PartTranspose]
	OnAfterPartTransposeDeleteCallback GongOnAfterDeleteInterface[PartTranspose]

	Slashs                map[*Slash]struct{}
	Slashs_instance       map[*Slash]*Slash
	Slashs_mapString      map[string]*Slash
	SlashOrder            uint
	Slash_stagedOrder     map[*Slash]uint
	Slash_orderStaged     map[uint]*Slash
	Slashs_reference      map[*Slash]*Slash
	Slashs_referenceOrder map[*Slash]uint

	// insertion point for slice of pointers maps
	OnAfterSlashCreateCallback GongOnAfterCreateInterface[Slash]
	OnAfterSlashUpdateCallback GongOnAfterUpdateInterface[Slash]
	OnAfterSlashDeleteCallback GongOnAfterDeleteInterface[Slash]

	StaffDetailss                map[*StaffDetails]struct{}
	StaffDetailss_instance       map[*StaffDetails]*StaffDetails
	StaffDetailss_mapString      map[string]*StaffDetails
	StaffDetailsOrder            uint
	StaffDetails_stagedOrder     map[*StaffDetails]uint
	StaffDetails_orderStaged     map[uint]*StaffDetails
	StaffDetailss_reference      map[*StaffDetails]*StaffDetails
	StaffDetailss_referenceOrder map[*StaffDetails]uint

	// insertion point for slice of pointers maps
	StaffDetails_LineDetail_reverseMap map[*LineDetail]*StaffDetails

	StaffDetails_StaffTuning_reverseMap map[*StaffTuning]*StaffDetails

	OnAfterStaffDetailsCreateCallback GongOnAfterCreateInterface[StaffDetails]
	OnAfterStaffDetailsUpdateCallback GongOnAfterUpdateInterface[StaffDetails]
	OnAfterStaffDetailsDeleteCallback GongOnAfterDeleteInterface[StaffDetails]

	StaffSizes                map[*StaffSize]struct{}
	StaffSizes_instance       map[*StaffSize]*StaffSize
	StaffSizes_mapString      map[string]*StaffSize
	StaffSizeOrder            uint
	StaffSize_stagedOrder     map[*StaffSize]uint
	StaffSize_orderStaged     map[uint]*StaffSize
	StaffSizes_reference      map[*StaffSize]*StaffSize
	StaffSizes_referenceOrder map[*StaffSize]uint

	// insertion point for slice of pointers maps
	OnAfterStaffSizeCreateCallback GongOnAfterCreateInterface[StaffSize]
	OnAfterStaffSizeUpdateCallback GongOnAfterUpdateInterface[StaffSize]
	OnAfterStaffSizeDeleteCallback GongOnAfterDeleteInterface[StaffSize]

	StaffTunings                map[*StaffTuning]struct{}
	StaffTunings_instance       map[*StaffTuning]*StaffTuning
	StaffTunings_mapString      map[string]*StaffTuning
	StaffTuningOrder            uint
	StaffTuning_stagedOrder     map[*StaffTuning]uint
	StaffTuning_orderStaged     map[uint]*StaffTuning
	StaffTunings_reference      map[*StaffTuning]*StaffTuning
	StaffTunings_referenceOrder map[*StaffTuning]uint

	// insertion point for slice of pointers maps
	OnAfterStaffTuningCreateCallback GongOnAfterCreateInterface[StaffTuning]
	OnAfterStaffTuningUpdateCallback GongOnAfterUpdateInterface[StaffTuning]
	OnAfterStaffTuningDeleteCallback GongOnAfterDeleteInterface[StaffTuning]

	Times                map[*Time]struct{}
	Times_instance       map[*Time]*Time
	Times_mapString      map[string]*Time
	TimeOrder            uint
	Time_stagedOrder     map[*Time]uint
	Time_orderStaged     map[uint]*Time
	Times_reference      map[*Time]*Time
	Times_referenceOrder map[*Time]uint

	// insertion point for slice of pointers maps
	OnAfterTimeCreateCallback GongOnAfterCreateInterface[Time]
	OnAfterTimeUpdateCallback GongOnAfterUpdateInterface[Time]
	OnAfterTimeDeleteCallback GongOnAfterDeleteInterface[Time]

	Transposes                map[*Transpose]struct{}
	Transposes_instance       map[*Transpose]*Transpose
	Transposes_mapString      map[string]*Transpose
	TransposeOrder            uint
	Transpose_stagedOrder     map[*Transpose]uint
	Transpose_orderStaged     map[uint]*Transpose
	Transposes_reference      map[*Transpose]*Transpose
	Transposes_referenceOrder map[*Transpose]uint

	// insertion point for slice of pointers maps
	OnAfterTransposeCreateCallback GongOnAfterCreateInterface[Transpose]
	OnAfterTransposeUpdateCallback GongOnAfterUpdateInterface[Transpose]
	OnAfterTransposeDeleteCallback GongOnAfterDeleteInterface[Transpose]

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
	__gong__clearReferences(&stage.Attributess_reference, &stage.Attributess_instance, &stage.Attributess_referenceOrder)

	__gong__clearReferences(&stage.AttributesDirectives_reference, &stage.AttributesDirectives_instance, &stage.AttributesDirectives_referenceOrder)

	__gong__clearReferences(&stage.BeatRepeats_reference, &stage.BeatRepeats_instance, &stage.BeatRepeats_referenceOrder)

	__gong__clearReferences(&stage.Cancels_reference, &stage.Cancels_instance, &stage.Cancels_referenceOrder)

	__gong__clearReferences(&stage.Clefs_reference, &stage.Clefs_instance, &stage.Clefs_referenceOrder)

	__gong__clearReferences(&stage.Doubles_reference, &stage.Doubles_instance, &stage.Doubles_referenceOrder)

	__gong__clearReferences(&stage.ForParts_reference, &stage.ForParts_instance, &stage.ForParts_referenceOrder)

	__gong__clearReferences(&stage.Interchangeables_reference, &stage.Interchangeables_instance, &stage.Interchangeables_referenceOrder)

	__gong__clearReferences(&stage.Keys_reference, &stage.Keys_instance, &stage.Keys_referenceOrder)

	__gong__clearReferences(&stage.KeyAccidentals_reference, &stage.KeyAccidentals_instance, &stage.KeyAccidentals_referenceOrder)

	__gong__clearReferences(&stage.KeyOctaves_reference, &stage.KeyOctaves_instance, &stage.KeyOctaves_referenceOrder)

	__gong__clearReferences(&stage.LineDetails_reference, &stage.LineDetails_instance, &stage.LineDetails_referenceOrder)

	__gong__clearReferences(&stage.MeasureRepeats_reference, &stage.MeasureRepeats_instance, &stage.MeasureRepeats_referenceOrder)

	__gong__clearReferences(&stage.MeasureStyles_reference, &stage.MeasureStyles_instance, &stage.MeasureStyles_referenceOrder)

	__gong__clearReferences(&stage.MultipleRests_reference, &stage.MultipleRests_instance, &stage.MultipleRests_referenceOrder)

	__gong__clearReferences(&stage.PartClefs_reference, &stage.PartClefs_instance, &stage.PartClefs_referenceOrder)

	__gong__clearReferences(&stage.PartSymbols_reference, &stage.PartSymbols_instance, &stage.PartSymbols_referenceOrder)

	__gong__clearReferences(&stage.PartTransposes_reference, &stage.PartTransposes_instance, &stage.PartTransposes_referenceOrder)

	__gong__clearReferences(&stage.Slashs_reference, &stage.Slashs_instance, &stage.Slashs_referenceOrder)

	__gong__clearReferences(&stage.StaffDetailss_reference, &stage.StaffDetailss_instance, &stage.StaffDetailss_referenceOrder)

	__gong__clearReferences(&stage.StaffSizes_reference, &stage.StaffSizes_instance, &stage.StaffSizes_referenceOrder)

	__gong__clearReferences(&stage.StaffTunings_reference, &stage.StaffTunings_instance, &stage.StaffTunings_referenceOrder)

	__gong__clearReferences(&stage.Times_reference, &stage.Times_instance, &stage.Times_referenceOrder)

	__gong__clearReferences(&stage.Transposes_reference, &stage.Transposes_instance, &stage.Transposes_referenceOrder)

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
	stage.AttributesOrder = __gong__recomputeOrder(stage.Attributes_stagedOrder)

	stage.AttributesDirectiveOrder = __gong__recomputeOrder(stage.AttributesDirective_stagedOrder)

	stage.BeatRepeatOrder = __gong__recomputeOrder(stage.BeatRepeat_stagedOrder)

	stage.CancelOrder = __gong__recomputeOrder(stage.Cancel_stagedOrder)

	stage.ClefOrder = __gong__recomputeOrder(stage.Clef_stagedOrder)

	stage.DoubleOrder = __gong__recomputeOrder(stage.Double_stagedOrder)

	stage.ForPartOrder = __gong__recomputeOrder(stage.ForPart_stagedOrder)

	stage.InterchangeableOrder = __gong__recomputeOrder(stage.Interchangeable_stagedOrder)

	stage.KeyOrder = __gong__recomputeOrder(stage.Key_stagedOrder)

	stage.KeyAccidentalOrder = __gong__recomputeOrder(stage.KeyAccidental_stagedOrder)

	stage.KeyOctaveOrder = __gong__recomputeOrder(stage.KeyOctave_stagedOrder)

	stage.LineDetailOrder = __gong__recomputeOrder(stage.LineDetail_stagedOrder)

	stage.MeasureRepeatOrder = __gong__recomputeOrder(stage.MeasureRepeat_stagedOrder)

	stage.MeasureStyleOrder = __gong__recomputeOrder(stage.MeasureStyle_stagedOrder)

	stage.MultipleRestOrder = __gong__recomputeOrder(stage.MultipleRest_stagedOrder)

	stage.PartClefOrder = __gong__recomputeOrder(stage.PartClef_stagedOrder)

	stage.PartSymbolOrder = __gong__recomputeOrder(stage.PartSymbol_stagedOrder)

	stage.PartTransposeOrder = __gong__recomputeOrder(stage.PartTranspose_stagedOrder)

	stage.SlashOrder = __gong__recomputeOrder(stage.Slash_stagedOrder)

	stage.StaffDetailsOrder = __gong__recomputeOrder(stage.StaffDetails_stagedOrder)

	stage.StaffSizeOrder = __gong__recomputeOrder(stage.StaffSize_stagedOrder)

	stage.StaffTuningOrder = __gong__recomputeOrder(stage.StaffTuning_stagedOrder)

	stage.TimeOrder = __gong__recomputeOrder(stage.Time_stagedOrder)

	stage.TransposeOrder = __gong__recomputeOrder(stage.Transpose_stagedOrder)

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
func (*Attributes) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Attributess, stage.Attributes_stagedOrder)
}

func (*Attributes) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Attributes_orderStaged[order]
}

func (*Attributes) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Attributess_mapString
}

func (*Attributes) GongGetInstancesSet(stage *Stage) any {
	return &stage.Attributess
}

func (*Attributes) GongNewInstance() any {
	return new(Attributes)
}

func (*AttributesDirective) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AttributesDirectives, stage.AttributesDirective_stagedOrder)
}

func (*AttributesDirective) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AttributesDirective_orderStaged[order]
}

func (*AttributesDirective) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AttributesDirectives_mapString
}

func (*AttributesDirective) GongGetInstancesSet(stage *Stage) any {
	return &stage.AttributesDirectives
}

func (*AttributesDirective) GongNewInstance() any {
	return new(AttributesDirective)
}

func (*BeatRepeat) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.BeatRepeats, stage.BeatRepeat_stagedOrder)
}

func (*BeatRepeat) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.BeatRepeat_orderStaged[order]
}

func (*BeatRepeat) GongGetInstancesMapByName(stage *Stage) any {
	return stage.BeatRepeats_mapString
}

func (*BeatRepeat) GongGetInstancesSet(stage *Stage) any {
	return &stage.BeatRepeats
}

func (*BeatRepeat) GongNewInstance() any {
	return new(BeatRepeat)
}

func (*Cancel) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Cancels, stage.Cancel_stagedOrder)
}

func (*Cancel) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Cancel_orderStaged[order]
}

func (*Cancel) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Cancels_mapString
}

func (*Cancel) GongGetInstancesSet(stage *Stage) any {
	return &stage.Cancels
}

func (*Cancel) GongNewInstance() any {
	return new(Cancel)
}

func (*Clef) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Clefs, stage.Clef_stagedOrder)
}

func (*Clef) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Clef_orderStaged[order]
}

func (*Clef) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Clefs_mapString
}

func (*Clef) GongGetInstancesSet(stage *Stage) any {
	return &stage.Clefs
}

func (*Clef) GongNewInstance() any {
	return new(Clef)
}

func (*Double) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Doubles, stage.Double_stagedOrder)
}

func (*Double) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Double_orderStaged[order]
}

func (*Double) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Doubles_mapString
}

func (*Double) GongGetInstancesSet(stage *Stage) any {
	return &stage.Doubles
}

func (*Double) GongNewInstance() any {
	return new(Double)
}

func (*ForPart) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ForParts, stage.ForPart_stagedOrder)
}

func (*ForPart) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ForPart_orderStaged[order]
}

func (*ForPart) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ForParts_mapString
}

func (*ForPart) GongGetInstancesSet(stage *Stage) any {
	return &stage.ForParts
}

func (*ForPart) GongNewInstance() any {
	return new(ForPart)
}

func (*Interchangeable) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Interchangeables, stage.Interchangeable_stagedOrder)
}

func (*Interchangeable) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Interchangeable_orderStaged[order]
}

func (*Interchangeable) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Interchangeables_mapString
}

func (*Interchangeable) GongGetInstancesSet(stage *Stage) any {
	return &stage.Interchangeables
}

func (*Interchangeable) GongNewInstance() any {
	return new(Interchangeable)
}

func (*Key) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Keys, stage.Key_stagedOrder)
}

func (*Key) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Key_orderStaged[order]
}

func (*Key) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Keys_mapString
}

func (*Key) GongGetInstancesSet(stage *Stage) any {
	return &stage.Keys
}

func (*Key) GongNewInstance() any {
	return new(Key)
}

func (*KeyAccidental) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.KeyAccidentals, stage.KeyAccidental_stagedOrder)
}

func (*KeyAccidental) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.KeyAccidental_orderStaged[order]
}

func (*KeyAccidental) GongGetInstancesMapByName(stage *Stage) any {
	return stage.KeyAccidentals_mapString
}

func (*KeyAccidental) GongGetInstancesSet(stage *Stage) any {
	return &stage.KeyAccidentals
}

func (*KeyAccidental) GongNewInstance() any {
	return new(KeyAccidental)
}

func (*KeyOctave) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.KeyOctaves, stage.KeyOctave_stagedOrder)
}

func (*KeyOctave) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.KeyOctave_orderStaged[order]
}

func (*KeyOctave) GongGetInstancesMapByName(stage *Stage) any {
	return stage.KeyOctaves_mapString
}

func (*KeyOctave) GongGetInstancesSet(stage *Stage) any {
	return &stage.KeyOctaves
}

func (*KeyOctave) GongNewInstance() any {
	return new(KeyOctave)
}

func (*LineDetail) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.LineDetails, stage.LineDetail_stagedOrder)
}

func (*LineDetail) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.LineDetail_orderStaged[order]
}

func (*LineDetail) GongGetInstancesMapByName(stage *Stage) any {
	return stage.LineDetails_mapString
}

func (*LineDetail) GongGetInstancesSet(stage *Stage) any {
	return &stage.LineDetails
}

func (*LineDetail) GongNewInstance() any {
	return new(LineDetail)
}

func (*MeasureRepeat) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MeasureRepeats, stage.MeasureRepeat_stagedOrder)
}

func (*MeasureRepeat) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MeasureRepeat_orderStaged[order]
}

func (*MeasureRepeat) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MeasureRepeats_mapString
}

func (*MeasureRepeat) GongGetInstancesSet(stage *Stage) any {
	return &stage.MeasureRepeats
}

func (*MeasureRepeat) GongNewInstance() any {
	return new(MeasureRepeat)
}

func (*MeasureStyle) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MeasureStyles, stage.MeasureStyle_stagedOrder)
}

func (*MeasureStyle) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MeasureStyle_orderStaged[order]
}

func (*MeasureStyle) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MeasureStyles_mapString
}

func (*MeasureStyle) GongGetInstancesSet(stage *Stage) any {
	return &stage.MeasureStyles
}

func (*MeasureStyle) GongNewInstance() any {
	return new(MeasureStyle)
}

func (*MultipleRest) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MultipleRests, stage.MultipleRest_stagedOrder)
}

func (*MultipleRest) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MultipleRest_orderStaged[order]
}

func (*MultipleRest) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MultipleRests_mapString
}

func (*MultipleRest) GongGetInstancesSet(stage *Stage) any {
	return &stage.MultipleRests
}

func (*MultipleRest) GongNewInstance() any {
	return new(MultipleRest)
}

func (*PartClef) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.PartClefs, stage.PartClef_stagedOrder)
}

func (*PartClef) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.PartClef_orderStaged[order]
}

func (*PartClef) GongGetInstancesMapByName(stage *Stage) any {
	return stage.PartClefs_mapString
}

func (*PartClef) GongGetInstancesSet(stage *Stage) any {
	return &stage.PartClefs
}

func (*PartClef) GongNewInstance() any {
	return new(PartClef)
}

func (*PartSymbol) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.PartSymbols, stage.PartSymbol_stagedOrder)
}

func (*PartSymbol) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.PartSymbol_orderStaged[order]
}

func (*PartSymbol) GongGetInstancesMapByName(stage *Stage) any {
	return stage.PartSymbols_mapString
}

func (*PartSymbol) GongGetInstancesSet(stage *Stage) any {
	return &stage.PartSymbols
}

func (*PartSymbol) GongNewInstance() any {
	return new(PartSymbol)
}

func (*PartTranspose) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.PartTransposes, stage.PartTranspose_stagedOrder)
}

func (*PartTranspose) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.PartTranspose_orderStaged[order]
}

func (*PartTranspose) GongGetInstancesMapByName(stage *Stage) any {
	return stage.PartTransposes_mapString
}

func (*PartTranspose) GongGetInstancesSet(stage *Stage) any {
	return &stage.PartTransposes
}

func (*PartTranspose) GongNewInstance() any {
	return new(PartTranspose)
}

func (*Slash) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Slashs, stage.Slash_stagedOrder)
}

func (*Slash) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Slash_orderStaged[order]
}

func (*Slash) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Slashs_mapString
}

func (*Slash) GongGetInstancesSet(stage *Stage) any {
	return &stage.Slashs
}

func (*Slash) GongNewInstance() any {
	return new(Slash)
}

func (*StaffDetails) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.StaffDetailss, stage.StaffDetails_stagedOrder)
}

func (*StaffDetails) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.StaffDetails_orderStaged[order]
}

func (*StaffDetails) GongGetInstancesMapByName(stage *Stage) any {
	return stage.StaffDetailss_mapString
}

func (*StaffDetails) GongGetInstancesSet(stage *Stage) any {
	return &stage.StaffDetailss
}

func (*StaffDetails) GongNewInstance() any {
	return new(StaffDetails)
}

func (*StaffSize) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.StaffSizes, stage.StaffSize_stagedOrder)
}

func (*StaffSize) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.StaffSize_orderStaged[order]
}

func (*StaffSize) GongGetInstancesMapByName(stage *Stage) any {
	return stage.StaffSizes_mapString
}

func (*StaffSize) GongGetInstancesSet(stage *Stage) any {
	return &stage.StaffSizes
}

func (*StaffSize) GongNewInstance() any {
	return new(StaffSize)
}

func (*StaffTuning) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.StaffTunings, stage.StaffTuning_stagedOrder)
}

func (*StaffTuning) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.StaffTuning_orderStaged[order]
}

func (*StaffTuning) GongGetInstancesMapByName(stage *Stage) any {
	return stage.StaffTunings_mapString
}

func (*StaffTuning) GongGetInstancesSet(stage *Stage) any {
	return &stage.StaffTunings
}

func (*StaffTuning) GongNewInstance() any {
	return new(StaffTuning)
}

func (*Time) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Times, stage.Time_stagedOrder)
}

func (*Time) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Time_orderStaged[order]
}

func (*Time) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Times_mapString
}

func (*Time) GongGetInstancesSet(stage *Stage) any {
	return &stage.Times
}

func (*Time) GongNewInstance() any {
	return new(Time)
}

func (*Transpose) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Transposes, stage.Transpose_stagedOrder)
}

func (*Transpose) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Transpose_orderStaged[order]
}

func (*Transpose) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Transposes_mapString
}

func (*Transpose) GongGetInstancesSet(stage *Stage) any {
	return &stage.Transposes
}

func (*Transpose) GongNewInstance() any {
	return new(Transpose)
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
	return "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/attributes"
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
		Attributess:           make(map[*Attributes]struct{}),
		Attributess_mapString: make(map[string]*Attributes),

		AttributesDirectives:           make(map[*AttributesDirective]struct{}),
		AttributesDirectives_mapString: make(map[string]*AttributesDirective),

		BeatRepeats:           make(map[*BeatRepeat]struct{}),
		BeatRepeats_mapString: make(map[string]*BeatRepeat),

		Cancels:           make(map[*Cancel]struct{}),
		Cancels_mapString: make(map[string]*Cancel),

		Clefs:           make(map[*Clef]struct{}),
		Clefs_mapString: make(map[string]*Clef),

		Doubles:           make(map[*Double]struct{}),
		Doubles_mapString: make(map[string]*Double),

		ForParts:           make(map[*ForPart]struct{}),
		ForParts_mapString: make(map[string]*ForPart),

		Interchangeables:           make(map[*Interchangeable]struct{}),
		Interchangeables_mapString: make(map[string]*Interchangeable),

		Keys:           make(map[*Key]struct{}),
		Keys_mapString: make(map[string]*Key),

		KeyAccidentals:           make(map[*KeyAccidental]struct{}),
		KeyAccidentals_mapString: make(map[string]*KeyAccidental),

		KeyOctaves:           make(map[*KeyOctave]struct{}),
		KeyOctaves_mapString: make(map[string]*KeyOctave),

		LineDetails:           make(map[*LineDetail]struct{}),
		LineDetails_mapString: make(map[string]*LineDetail),

		MeasureRepeats:           make(map[*MeasureRepeat]struct{}),
		MeasureRepeats_mapString: make(map[string]*MeasureRepeat),

		MeasureStyles:           make(map[*MeasureStyle]struct{}),
		MeasureStyles_mapString: make(map[string]*MeasureStyle),

		MultipleRests:           make(map[*MultipleRest]struct{}),
		MultipleRests_mapString: make(map[string]*MultipleRest),

		PartClefs:           make(map[*PartClef]struct{}),
		PartClefs_mapString: make(map[string]*PartClef),

		PartSymbols:           make(map[*PartSymbol]struct{}),
		PartSymbols_mapString: make(map[string]*PartSymbol),

		PartTransposes:           make(map[*PartTranspose]struct{}),
		PartTransposes_mapString: make(map[string]*PartTranspose),

		Slashs:           make(map[*Slash]struct{}),
		Slashs_mapString: make(map[string]*Slash),

		StaffDetailss:           make(map[*StaffDetails]struct{}),
		StaffDetailss_mapString: make(map[string]*StaffDetails),

		StaffSizes:           make(map[*StaffSize]struct{}),
		StaffSizes_mapString: make(map[string]*StaffSize),

		StaffTunings:           make(map[*StaffTuning]struct{}),
		StaffTunings_mapString: make(map[string]*StaffTuning),

		Times:           make(map[*Time]struct{}),
		Times_mapString: make(map[string]*Time),

		Transposes:           make(map[*Transpose]struct{}),
		Transposes_mapString: make(map[string]*Transpose),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		Attributes_stagedOrder: make(map[*Attributes]uint),
		Attributes_orderStaged: make(map[uint]*Attributes),
		Attributess_reference:  make(map[*Attributes]*Attributes),

		AttributesDirective_stagedOrder: make(map[*AttributesDirective]uint),
		AttributesDirective_orderStaged: make(map[uint]*AttributesDirective),
		AttributesDirectives_reference:  make(map[*AttributesDirective]*AttributesDirective),

		BeatRepeat_stagedOrder: make(map[*BeatRepeat]uint),
		BeatRepeat_orderStaged: make(map[uint]*BeatRepeat),
		BeatRepeats_reference:  make(map[*BeatRepeat]*BeatRepeat),

		Cancel_stagedOrder: make(map[*Cancel]uint),
		Cancel_orderStaged: make(map[uint]*Cancel),
		Cancels_reference:  make(map[*Cancel]*Cancel),

		Clef_stagedOrder: make(map[*Clef]uint),
		Clef_orderStaged: make(map[uint]*Clef),
		Clefs_reference:  make(map[*Clef]*Clef),

		Double_stagedOrder: make(map[*Double]uint),
		Double_orderStaged: make(map[uint]*Double),
		Doubles_reference:  make(map[*Double]*Double),

		ForPart_stagedOrder: make(map[*ForPart]uint),
		ForPart_orderStaged: make(map[uint]*ForPart),
		ForParts_reference:  make(map[*ForPart]*ForPart),

		Interchangeable_stagedOrder: make(map[*Interchangeable]uint),
		Interchangeable_orderStaged: make(map[uint]*Interchangeable),
		Interchangeables_reference:  make(map[*Interchangeable]*Interchangeable),

		Key_stagedOrder: make(map[*Key]uint),
		Key_orderStaged: make(map[uint]*Key),
		Keys_reference:  make(map[*Key]*Key),

		KeyAccidental_stagedOrder: make(map[*KeyAccidental]uint),
		KeyAccidental_orderStaged: make(map[uint]*KeyAccidental),
		KeyAccidentals_reference:  make(map[*KeyAccidental]*KeyAccidental),

		KeyOctave_stagedOrder: make(map[*KeyOctave]uint),
		KeyOctave_orderStaged: make(map[uint]*KeyOctave),
		KeyOctaves_reference:  make(map[*KeyOctave]*KeyOctave),

		LineDetail_stagedOrder: make(map[*LineDetail]uint),
		LineDetail_orderStaged: make(map[uint]*LineDetail),
		LineDetails_reference:  make(map[*LineDetail]*LineDetail),

		MeasureRepeat_stagedOrder: make(map[*MeasureRepeat]uint),
		MeasureRepeat_orderStaged: make(map[uint]*MeasureRepeat),
		MeasureRepeats_reference:  make(map[*MeasureRepeat]*MeasureRepeat),

		MeasureStyle_stagedOrder: make(map[*MeasureStyle]uint),
		MeasureStyle_orderStaged: make(map[uint]*MeasureStyle),
		MeasureStyles_reference:  make(map[*MeasureStyle]*MeasureStyle),

		MultipleRest_stagedOrder: make(map[*MultipleRest]uint),
		MultipleRest_orderStaged: make(map[uint]*MultipleRest),
		MultipleRests_reference:  make(map[*MultipleRest]*MultipleRest),

		PartClef_stagedOrder: make(map[*PartClef]uint),
		PartClef_orderStaged: make(map[uint]*PartClef),
		PartClefs_reference:  make(map[*PartClef]*PartClef),

		PartSymbol_stagedOrder: make(map[*PartSymbol]uint),
		PartSymbol_orderStaged: make(map[uint]*PartSymbol),
		PartSymbols_reference:  make(map[*PartSymbol]*PartSymbol),

		PartTranspose_stagedOrder: make(map[*PartTranspose]uint),
		PartTranspose_orderStaged: make(map[uint]*PartTranspose),
		PartTransposes_reference:  make(map[*PartTranspose]*PartTranspose),

		Slash_stagedOrder: make(map[*Slash]uint),
		Slash_orderStaged: make(map[uint]*Slash),
		Slashs_reference:  make(map[*Slash]*Slash),

		StaffDetails_stagedOrder: make(map[*StaffDetails]uint),
		StaffDetails_orderStaged: make(map[uint]*StaffDetails),
		StaffDetailss_reference:  make(map[*StaffDetails]*StaffDetails),

		StaffSize_stagedOrder: make(map[*StaffSize]uint),
		StaffSize_orderStaged: make(map[uint]*StaffSize),
		StaffSizes_reference:  make(map[*StaffSize]*StaffSize),

		StaffTuning_stagedOrder: make(map[*StaffTuning]uint),
		StaffTuning_orderStaged: make(map[uint]*StaffTuning),
		StaffTunings_reference:  make(map[*StaffTuning]*StaffTuning),

		Time_stagedOrder: make(map[*Time]uint),
		Time_orderStaged: make(map[uint]*Time),
		Times_reference:  make(map[*Time]*Time),

		Transpose_stagedOrder: make(map[*Transpose]uint),
		Transpose_orderStaged: make(map[uint]*Transpose),
		Transposes_reference:  make(map[*Transpose]*Transpose),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"Attributes": &AttributesUnmarshaller{},

			"AttributesDirective": &AttributesDirectiveUnmarshaller{},

			"BeatRepeat": &BeatRepeatUnmarshaller{},

			"Cancel": &CancelUnmarshaller{},

			"Clef": &ClefUnmarshaller{},

			"Double": &DoubleUnmarshaller{},

			"ForPart": &ForPartUnmarshaller{},

			"Interchangeable": &InterchangeableUnmarshaller{},

			"Key": &KeyUnmarshaller{},

			"KeyAccidental": &KeyAccidentalUnmarshaller{},

			"KeyOctave": &KeyOctaveUnmarshaller{},

			"LineDetail": &LineDetailUnmarshaller{},

			"MeasureRepeat": &MeasureRepeatUnmarshaller{},

			"MeasureStyle": &MeasureStyleUnmarshaller{},

			"MultipleRest": &MultipleRestUnmarshaller{},

			"PartClef": &PartClefUnmarshaller{},

			"PartSymbol": &PartSymbolUnmarshaller{},

			"PartTranspose": &PartTransposeUnmarshaller{},

			"Slash": &SlashUnmarshaller{},

			"StaffDetails": &StaffDetailsUnmarshaller{},

			"StaffSize": &StaffSizeUnmarshaller{},

			"StaffTuning": &StaffTuningUnmarshaller{},

			"Time": &TimeUnmarshaller{},

			"Transpose": &TransposeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["Attributes"] = len(stage.Attributess)
	stage.Map_GongStructName_InstancesNb["AttributesDirective"] = len(stage.AttributesDirectives)
	stage.Map_GongStructName_InstancesNb["BeatRepeat"] = len(stage.BeatRepeats)
	stage.Map_GongStructName_InstancesNb["Cancel"] = len(stage.Cancels)
	stage.Map_GongStructName_InstancesNb["Clef"] = len(stage.Clefs)
	stage.Map_GongStructName_InstancesNb["Double"] = len(stage.Doubles)
	stage.Map_GongStructName_InstancesNb["ForPart"] = len(stage.ForParts)
	stage.Map_GongStructName_InstancesNb["Interchangeable"] = len(stage.Interchangeables)
	stage.Map_GongStructName_InstancesNb["Key"] = len(stage.Keys)
	stage.Map_GongStructName_InstancesNb["KeyAccidental"] = len(stage.KeyAccidentals)
	stage.Map_GongStructName_InstancesNb["KeyOctave"] = len(stage.KeyOctaves)
	stage.Map_GongStructName_InstancesNb["LineDetail"] = len(stage.LineDetails)
	stage.Map_GongStructName_InstancesNb["MeasureRepeat"] = len(stage.MeasureRepeats)
	stage.Map_GongStructName_InstancesNb["MeasureStyle"] = len(stage.MeasureStyles)
	stage.Map_GongStructName_InstancesNb["MultipleRest"] = len(stage.MultipleRests)
	stage.Map_GongStructName_InstancesNb["PartClef"] = len(stage.PartClefs)
	stage.Map_GongStructName_InstancesNb["PartSymbol"] = len(stage.PartSymbols)
	stage.Map_GongStructName_InstancesNb["PartTranspose"] = len(stage.PartTransposes)
	stage.Map_GongStructName_InstancesNb["Slash"] = len(stage.Slashs)
	stage.Map_GongStructName_InstancesNb["StaffDetails"] = len(stage.StaffDetailss)
	stage.Map_GongStructName_InstancesNb["StaffSize"] = len(stage.StaffSizes)
	stage.Map_GongStructName_InstancesNb["StaffTuning"] = len(stage.StaffTunings)
	stage.Map_GongStructName_InstancesNb["Time"] = len(stage.Times)
	stage.Map_GongStructName_InstancesNb["Transpose"] = len(stage.Transposes)
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
// Stage puts attributes to the model stage
func (attributes *Attributes) Stage(stage *Stage) *Attributes {
	__gong__stage(stage.Attributess, stage.Attributes_stagedOrder, stage.Attributes_orderStaged, &stage.AttributesOrder, stage.Attributess_mapString, attributes, attributes.Name)
	return attributes
}

// StagePreserveOrder puts attributes to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AttributesOrder
// - update stage.AttributesOrder accordingly
func (attributes *Attributes) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Attributess, stage.Attributes_stagedOrder, stage.Attributes_orderStaged, &stage.AttributesOrder, stage.Attributess_mapString, attributes, order, attributes.Name)
}

// Unstage removes attributes off the model stage
func (attributes *Attributes) Unstage(stage *Stage) *Attributes {
	__gong__unstage(stage.Attributess, stage.Attributess_mapString, attributes, attributes.Name)
	return attributes
}

// UnstageVoid removes attributes off the model stage
func (attributes *Attributes) UnstageVoid(stage *Stage) {
	attributes.Unstage(stage)
}

func (attributes *Attributes) StageVoid(stage *Stage) {
	attributes.Stage(stage)
}

// for satisfaction of GongStruct interface
func (attributes *Attributes) GetName() (res string) {
	return attributes.Name
}

// for satisfaction of GongStruct interface
func (attributes *Attributes) SetName(name string) {
	attributes.Name = name
}

// Stage puts attributesdirective to the model stage
func (attributesdirective *AttributesDirective) Stage(stage *Stage) *AttributesDirective {
	__gong__stage(stage.AttributesDirectives, stage.AttributesDirective_stagedOrder, stage.AttributesDirective_orderStaged, &stage.AttributesDirectiveOrder, stage.AttributesDirectives_mapString, attributesdirective, attributesdirective.Name)
	return attributesdirective
}

// StagePreserveOrder puts attributesdirective to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AttributesDirectiveOrder
// - update stage.AttributesDirectiveOrder accordingly
func (attributesdirective *AttributesDirective) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AttributesDirectives, stage.AttributesDirective_stagedOrder, stage.AttributesDirective_orderStaged, &stage.AttributesDirectiveOrder, stage.AttributesDirectives_mapString, attributesdirective, order, attributesdirective.Name)
}

// Unstage removes attributesdirective off the model stage
func (attributesdirective *AttributesDirective) Unstage(stage *Stage) *AttributesDirective {
	__gong__unstage(stage.AttributesDirectives, stage.AttributesDirectives_mapString, attributesdirective, attributesdirective.Name)
	return attributesdirective
}

// UnstageVoid removes attributesdirective off the model stage
func (attributesdirective *AttributesDirective) UnstageVoid(stage *Stage) {
	attributesdirective.Unstage(stage)
}

func (attributesdirective *AttributesDirective) StageVoid(stage *Stage) {
	attributesdirective.Stage(stage)
}

// for satisfaction of GongStruct interface
func (attributesdirective *AttributesDirective) GetName() (res string) {
	return attributesdirective.Name
}

// for satisfaction of GongStruct interface
func (attributesdirective *AttributesDirective) SetName(name string) {
	attributesdirective.Name = name
}

// Stage puts beatrepeat to the model stage
func (beatrepeat *BeatRepeat) Stage(stage *Stage) *BeatRepeat {
	__gong__stage(stage.BeatRepeats, stage.BeatRepeat_stagedOrder, stage.BeatRepeat_orderStaged, &stage.BeatRepeatOrder, stage.BeatRepeats_mapString, beatrepeat, beatrepeat.Name)
	return beatrepeat
}

// StagePreserveOrder puts beatrepeat to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BeatRepeatOrder
// - update stage.BeatRepeatOrder accordingly
func (beatrepeat *BeatRepeat) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.BeatRepeats, stage.BeatRepeat_stagedOrder, stage.BeatRepeat_orderStaged, &stage.BeatRepeatOrder, stage.BeatRepeats_mapString, beatrepeat, order, beatrepeat.Name)
}

// Unstage removes beatrepeat off the model stage
func (beatrepeat *BeatRepeat) Unstage(stage *Stage) *BeatRepeat {
	__gong__unstage(stage.BeatRepeats, stage.BeatRepeats_mapString, beatrepeat, beatrepeat.Name)
	return beatrepeat
}

// UnstageVoid removes beatrepeat off the model stage
func (beatrepeat *BeatRepeat) UnstageVoid(stage *Stage) {
	beatrepeat.Unstage(stage)
}

func (beatrepeat *BeatRepeat) StageVoid(stage *Stage) {
	beatrepeat.Stage(stage)
}

// for satisfaction of GongStruct interface
func (beatrepeat *BeatRepeat) GetName() (res string) {
	return beatrepeat.Name
}

// for satisfaction of GongStruct interface
func (beatrepeat *BeatRepeat) SetName(name string) {
	beatrepeat.Name = name
}

// Stage puts cancel to the model stage
func (cancel *Cancel) Stage(stage *Stage) *Cancel {
	__gong__stage(stage.Cancels, stage.Cancel_stagedOrder, stage.Cancel_orderStaged, &stage.CancelOrder, stage.Cancels_mapString, cancel, cancel.Name)
	return cancel
}

// StagePreserveOrder puts cancel to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CancelOrder
// - update stage.CancelOrder accordingly
func (cancel *Cancel) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Cancels, stage.Cancel_stagedOrder, stage.Cancel_orderStaged, &stage.CancelOrder, stage.Cancels_mapString, cancel, order, cancel.Name)
}

// Unstage removes cancel off the model stage
func (cancel *Cancel) Unstage(stage *Stage) *Cancel {
	__gong__unstage(stage.Cancels, stage.Cancels_mapString, cancel, cancel.Name)
	return cancel
}

// UnstageVoid removes cancel off the model stage
func (cancel *Cancel) UnstageVoid(stage *Stage) {
	cancel.Unstage(stage)
}

func (cancel *Cancel) StageVoid(stage *Stage) {
	cancel.Stage(stage)
}

// for satisfaction of GongStruct interface
func (cancel *Cancel) GetName() (res string) {
	return cancel.Name
}

// for satisfaction of GongStruct interface
func (cancel *Cancel) SetName(name string) {
	cancel.Name = name
}

// Stage puts clef to the model stage
func (clef *Clef) Stage(stage *Stage) *Clef {
	__gong__stage(stage.Clefs, stage.Clef_stagedOrder, stage.Clef_orderStaged, &stage.ClefOrder, stage.Clefs_mapString, clef, clef.Name)
	return clef
}

// StagePreserveOrder puts clef to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ClefOrder
// - update stage.ClefOrder accordingly
func (clef *Clef) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Clefs, stage.Clef_stagedOrder, stage.Clef_orderStaged, &stage.ClefOrder, stage.Clefs_mapString, clef, order, clef.Name)
}

// Unstage removes clef off the model stage
func (clef *Clef) Unstage(stage *Stage) *Clef {
	__gong__unstage(stage.Clefs, stage.Clefs_mapString, clef, clef.Name)
	return clef
}

// UnstageVoid removes clef off the model stage
func (clef *Clef) UnstageVoid(stage *Stage) {
	clef.Unstage(stage)
}

func (clef *Clef) StageVoid(stage *Stage) {
	clef.Stage(stage)
}

// for satisfaction of GongStruct interface
func (clef *Clef) GetName() (res string) {
	return clef.Name
}

// for satisfaction of GongStruct interface
func (clef *Clef) SetName(name string) {
	clef.Name = name
}

// Stage puts double to the model stage
func (double *Double) Stage(stage *Stage) *Double {
	__gong__stage(stage.Doubles, stage.Double_stagedOrder, stage.Double_orderStaged, &stage.DoubleOrder, stage.Doubles_mapString, double, double.Name)
	return double
}

// StagePreserveOrder puts double to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DoubleOrder
// - update stage.DoubleOrder accordingly
func (double *Double) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Doubles, stage.Double_stagedOrder, stage.Double_orderStaged, &stage.DoubleOrder, stage.Doubles_mapString, double, order, double.Name)
}

// Unstage removes double off the model stage
func (double *Double) Unstage(stage *Stage) *Double {
	__gong__unstage(stage.Doubles, stage.Doubles_mapString, double, double.Name)
	return double
}

// UnstageVoid removes double off the model stage
func (double *Double) UnstageVoid(stage *Stage) {
	double.Unstage(stage)
}

func (double *Double) StageVoid(stage *Stage) {
	double.Stage(stage)
}

// for satisfaction of GongStruct interface
func (double *Double) GetName() (res string) {
	return double.Name
}

// for satisfaction of GongStruct interface
func (double *Double) SetName(name string) {
	double.Name = name
}

// Stage puts forpart to the model stage
func (forpart *ForPart) Stage(stage *Stage) *ForPart {
	__gong__stage(stage.ForParts, stage.ForPart_stagedOrder, stage.ForPart_orderStaged, &stage.ForPartOrder, stage.ForParts_mapString, forpart, forpart.Name)
	return forpart
}

// StagePreserveOrder puts forpart to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ForPartOrder
// - update stage.ForPartOrder accordingly
func (forpart *ForPart) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ForParts, stage.ForPart_stagedOrder, stage.ForPart_orderStaged, &stage.ForPartOrder, stage.ForParts_mapString, forpart, order, forpart.Name)
}

// Unstage removes forpart off the model stage
func (forpart *ForPart) Unstage(stage *Stage) *ForPart {
	__gong__unstage(stage.ForParts, stage.ForParts_mapString, forpart, forpart.Name)
	return forpart
}

// UnstageVoid removes forpart off the model stage
func (forpart *ForPart) UnstageVoid(stage *Stage) {
	forpart.Unstage(stage)
}

func (forpart *ForPart) StageVoid(stage *Stage) {
	forpart.Stage(stage)
}

// for satisfaction of GongStruct interface
func (forpart *ForPart) GetName() (res string) {
	return forpart.Name
}

// for satisfaction of GongStruct interface
func (forpart *ForPart) SetName(name string) {
	forpart.Name = name
}

// Stage puts interchangeable to the model stage
func (interchangeable *Interchangeable) Stage(stage *Stage) *Interchangeable {
	__gong__stage(stage.Interchangeables, stage.Interchangeable_stagedOrder, stage.Interchangeable_orderStaged, &stage.InterchangeableOrder, stage.Interchangeables_mapString, interchangeable, interchangeable.Name)
	return interchangeable
}

// StagePreserveOrder puts interchangeable to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.InterchangeableOrder
// - update stage.InterchangeableOrder accordingly
func (interchangeable *Interchangeable) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Interchangeables, stage.Interchangeable_stagedOrder, stage.Interchangeable_orderStaged, &stage.InterchangeableOrder, stage.Interchangeables_mapString, interchangeable, order, interchangeable.Name)
}

// Unstage removes interchangeable off the model stage
func (interchangeable *Interchangeable) Unstage(stage *Stage) *Interchangeable {
	__gong__unstage(stage.Interchangeables, stage.Interchangeables_mapString, interchangeable, interchangeable.Name)
	return interchangeable
}

// UnstageVoid removes interchangeable off the model stage
func (interchangeable *Interchangeable) UnstageVoid(stage *Stage) {
	interchangeable.Unstage(stage)
}

func (interchangeable *Interchangeable) StageVoid(stage *Stage) {
	interchangeable.Stage(stage)
}

// for satisfaction of GongStruct interface
func (interchangeable *Interchangeable) GetName() (res string) {
	return interchangeable.Name
}

// for satisfaction of GongStruct interface
func (interchangeable *Interchangeable) SetName(name string) {
	interchangeable.Name = name
}

// Stage puts key to the model stage
func (key *Key) Stage(stage *Stage) *Key {
	__gong__stage(stage.Keys, stage.Key_stagedOrder, stage.Key_orderStaged, &stage.KeyOrder, stage.Keys_mapString, key, key.Name)
	return key
}

// StagePreserveOrder puts key to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.KeyOrder
// - update stage.KeyOrder accordingly
func (key *Key) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Keys, stage.Key_stagedOrder, stage.Key_orderStaged, &stage.KeyOrder, stage.Keys_mapString, key, order, key.Name)
}

// Unstage removes key off the model stage
func (key *Key) Unstage(stage *Stage) *Key {
	__gong__unstage(stage.Keys, stage.Keys_mapString, key, key.Name)
	return key
}

// UnstageVoid removes key off the model stage
func (key *Key) UnstageVoid(stage *Stage) {
	key.Unstage(stage)
}

func (key *Key) StageVoid(stage *Stage) {
	key.Stage(stage)
}

// for satisfaction of GongStruct interface
func (key *Key) GetName() (res string) {
	return key.Name
}

// for satisfaction of GongStruct interface
func (key *Key) SetName(name string) {
	key.Name = name
}

// Stage puts keyaccidental to the model stage
func (keyaccidental *KeyAccidental) Stage(stage *Stage) *KeyAccidental {
	__gong__stage(stage.KeyAccidentals, stage.KeyAccidental_stagedOrder, stage.KeyAccidental_orderStaged, &stage.KeyAccidentalOrder, stage.KeyAccidentals_mapString, keyaccidental, keyaccidental.Name)
	return keyaccidental
}

// StagePreserveOrder puts keyaccidental to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.KeyAccidentalOrder
// - update stage.KeyAccidentalOrder accordingly
func (keyaccidental *KeyAccidental) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.KeyAccidentals, stage.KeyAccidental_stagedOrder, stage.KeyAccidental_orderStaged, &stage.KeyAccidentalOrder, stage.KeyAccidentals_mapString, keyaccidental, order, keyaccidental.Name)
}

// Unstage removes keyaccidental off the model stage
func (keyaccidental *KeyAccidental) Unstage(stage *Stage) *KeyAccidental {
	__gong__unstage(stage.KeyAccidentals, stage.KeyAccidentals_mapString, keyaccidental, keyaccidental.Name)
	return keyaccidental
}

// UnstageVoid removes keyaccidental off the model stage
func (keyaccidental *KeyAccidental) UnstageVoid(stage *Stage) {
	keyaccidental.Unstage(stage)
}

func (keyaccidental *KeyAccidental) StageVoid(stage *Stage) {
	keyaccidental.Stage(stage)
}

// for satisfaction of GongStruct interface
func (keyaccidental *KeyAccidental) GetName() (res string) {
	return keyaccidental.Name
}

// for satisfaction of GongStruct interface
func (keyaccidental *KeyAccidental) SetName(name string) {
	keyaccidental.Name = name
}

// Stage puts keyoctave to the model stage
func (keyoctave *KeyOctave) Stage(stage *Stage) *KeyOctave {
	__gong__stage(stage.KeyOctaves, stage.KeyOctave_stagedOrder, stage.KeyOctave_orderStaged, &stage.KeyOctaveOrder, stage.KeyOctaves_mapString, keyoctave, keyoctave.Name)
	return keyoctave
}

// StagePreserveOrder puts keyoctave to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.KeyOctaveOrder
// - update stage.KeyOctaveOrder accordingly
func (keyoctave *KeyOctave) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.KeyOctaves, stage.KeyOctave_stagedOrder, stage.KeyOctave_orderStaged, &stage.KeyOctaveOrder, stage.KeyOctaves_mapString, keyoctave, order, keyoctave.Name)
}

// Unstage removes keyoctave off the model stage
func (keyoctave *KeyOctave) Unstage(stage *Stage) *KeyOctave {
	__gong__unstage(stage.KeyOctaves, stage.KeyOctaves_mapString, keyoctave, keyoctave.Name)
	return keyoctave
}

// UnstageVoid removes keyoctave off the model stage
func (keyoctave *KeyOctave) UnstageVoid(stage *Stage) {
	keyoctave.Unstage(stage)
}

func (keyoctave *KeyOctave) StageVoid(stage *Stage) {
	keyoctave.Stage(stage)
}

// for satisfaction of GongStruct interface
func (keyoctave *KeyOctave) GetName() (res string) {
	return keyoctave.Name
}

// for satisfaction of GongStruct interface
func (keyoctave *KeyOctave) SetName(name string) {
	keyoctave.Name = name
}

// Stage puts linedetail to the model stage
func (linedetail *LineDetail) Stage(stage *Stage) *LineDetail {
	__gong__stage(stage.LineDetails, stage.LineDetail_stagedOrder, stage.LineDetail_orderStaged, &stage.LineDetailOrder, stage.LineDetails_mapString, linedetail, linedetail.Name)
	return linedetail
}

// StagePreserveOrder puts linedetail to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LineDetailOrder
// - update stage.LineDetailOrder accordingly
func (linedetail *LineDetail) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.LineDetails, stage.LineDetail_stagedOrder, stage.LineDetail_orderStaged, &stage.LineDetailOrder, stage.LineDetails_mapString, linedetail, order, linedetail.Name)
}

// Unstage removes linedetail off the model stage
func (linedetail *LineDetail) Unstage(stage *Stage) *LineDetail {
	__gong__unstage(stage.LineDetails, stage.LineDetails_mapString, linedetail, linedetail.Name)
	return linedetail
}

// UnstageVoid removes linedetail off the model stage
func (linedetail *LineDetail) UnstageVoid(stage *Stage) {
	linedetail.Unstage(stage)
}

func (linedetail *LineDetail) StageVoid(stage *Stage) {
	linedetail.Stage(stage)
}

// for satisfaction of GongStruct interface
func (linedetail *LineDetail) GetName() (res string) {
	return linedetail.Name
}

// for satisfaction of GongStruct interface
func (linedetail *LineDetail) SetName(name string) {
	linedetail.Name = name
}

// Stage puts measurerepeat to the model stage
func (measurerepeat *MeasureRepeat) Stage(stage *Stage) *MeasureRepeat {
	__gong__stage(stage.MeasureRepeats, stage.MeasureRepeat_stagedOrder, stage.MeasureRepeat_orderStaged, &stage.MeasureRepeatOrder, stage.MeasureRepeats_mapString, measurerepeat, measurerepeat.Name)
	return measurerepeat
}

// StagePreserveOrder puts measurerepeat to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MeasureRepeatOrder
// - update stage.MeasureRepeatOrder accordingly
func (measurerepeat *MeasureRepeat) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MeasureRepeats, stage.MeasureRepeat_stagedOrder, stage.MeasureRepeat_orderStaged, &stage.MeasureRepeatOrder, stage.MeasureRepeats_mapString, measurerepeat, order, measurerepeat.Name)
}

// Unstage removes measurerepeat off the model stage
func (measurerepeat *MeasureRepeat) Unstage(stage *Stage) *MeasureRepeat {
	__gong__unstage(stage.MeasureRepeats, stage.MeasureRepeats_mapString, measurerepeat, measurerepeat.Name)
	return measurerepeat
}

// UnstageVoid removes measurerepeat off the model stage
func (measurerepeat *MeasureRepeat) UnstageVoid(stage *Stage) {
	measurerepeat.Unstage(stage)
}

func (measurerepeat *MeasureRepeat) StageVoid(stage *Stage) {
	measurerepeat.Stage(stage)
}

// for satisfaction of GongStruct interface
func (measurerepeat *MeasureRepeat) GetName() (res string) {
	return measurerepeat.Name
}

// for satisfaction of GongStruct interface
func (measurerepeat *MeasureRepeat) SetName(name string) {
	measurerepeat.Name = name
}

// Stage puts measurestyle to the model stage
func (measurestyle *MeasureStyle) Stage(stage *Stage) *MeasureStyle {
	__gong__stage(stage.MeasureStyles, stage.MeasureStyle_stagedOrder, stage.MeasureStyle_orderStaged, &stage.MeasureStyleOrder, stage.MeasureStyles_mapString, measurestyle, measurestyle.Name)
	return measurestyle
}

// StagePreserveOrder puts measurestyle to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MeasureStyleOrder
// - update stage.MeasureStyleOrder accordingly
func (measurestyle *MeasureStyle) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MeasureStyles, stage.MeasureStyle_stagedOrder, stage.MeasureStyle_orderStaged, &stage.MeasureStyleOrder, stage.MeasureStyles_mapString, measurestyle, order, measurestyle.Name)
}

// Unstage removes measurestyle off the model stage
func (measurestyle *MeasureStyle) Unstage(stage *Stage) *MeasureStyle {
	__gong__unstage(stage.MeasureStyles, stage.MeasureStyles_mapString, measurestyle, measurestyle.Name)
	return measurestyle
}

// UnstageVoid removes measurestyle off the model stage
func (measurestyle *MeasureStyle) UnstageVoid(stage *Stage) {
	measurestyle.Unstage(stage)
}

func (measurestyle *MeasureStyle) StageVoid(stage *Stage) {
	measurestyle.Stage(stage)
}

// for satisfaction of GongStruct interface
func (measurestyle *MeasureStyle) GetName() (res string) {
	return measurestyle.Name
}

// for satisfaction of GongStruct interface
func (measurestyle *MeasureStyle) SetName(name string) {
	measurestyle.Name = name
}

// Stage puts multiplerest to the model stage
func (multiplerest *MultipleRest) Stage(stage *Stage) *MultipleRest {
	__gong__stage(stage.MultipleRests, stage.MultipleRest_stagedOrder, stage.MultipleRest_orderStaged, &stage.MultipleRestOrder, stage.MultipleRests_mapString, multiplerest, multiplerest.Name)
	return multiplerest
}

// StagePreserveOrder puts multiplerest to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MultipleRestOrder
// - update stage.MultipleRestOrder accordingly
func (multiplerest *MultipleRest) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MultipleRests, stage.MultipleRest_stagedOrder, stage.MultipleRest_orderStaged, &stage.MultipleRestOrder, stage.MultipleRests_mapString, multiplerest, order, multiplerest.Name)
}

// Unstage removes multiplerest off the model stage
func (multiplerest *MultipleRest) Unstage(stage *Stage) *MultipleRest {
	__gong__unstage(stage.MultipleRests, stage.MultipleRests_mapString, multiplerest, multiplerest.Name)
	return multiplerest
}

// UnstageVoid removes multiplerest off the model stage
func (multiplerest *MultipleRest) UnstageVoid(stage *Stage) {
	multiplerest.Unstage(stage)
}

func (multiplerest *MultipleRest) StageVoid(stage *Stage) {
	multiplerest.Stage(stage)
}

// for satisfaction of GongStruct interface
func (multiplerest *MultipleRest) GetName() (res string) {
	return multiplerest.Name
}

// for satisfaction of GongStruct interface
func (multiplerest *MultipleRest) SetName(name string) {
	multiplerest.Name = name
}

// Stage puts partclef to the model stage
func (partclef *PartClef) Stage(stage *Stage) *PartClef {
	__gong__stage(stage.PartClefs, stage.PartClef_stagedOrder, stage.PartClef_orderStaged, &stage.PartClefOrder, stage.PartClefs_mapString, partclef, partclef.Name)
	return partclef
}

// StagePreserveOrder puts partclef to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PartClefOrder
// - update stage.PartClefOrder accordingly
func (partclef *PartClef) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.PartClefs, stage.PartClef_stagedOrder, stage.PartClef_orderStaged, &stage.PartClefOrder, stage.PartClefs_mapString, partclef, order, partclef.Name)
}

// Unstage removes partclef off the model stage
func (partclef *PartClef) Unstage(stage *Stage) *PartClef {
	__gong__unstage(stage.PartClefs, stage.PartClefs_mapString, partclef, partclef.Name)
	return partclef
}

// UnstageVoid removes partclef off the model stage
func (partclef *PartClef) UnstageVoid(stage *Stage) {
	partclef.Unstage(stage)
}

func (partclef *PartClef) StageVoid(stage *Stage) {
	partclef.Stage(stage)
}

// for satisfaction of GongStruct interface
func (partclef *PartClef) GetName() (res string) {
	return partclef.Name
}

// for satisfaction of GongStruct interface
func (partclef *PartClef) SetName(name string) {
	partclef.Name = name
}

// Stage puts partsymbol to the model stage
func (partsymbol *PartSymbol) Stage(stage *Stage) *PartSymbol {
	__gong__stage(stage.PartSymbols, stage.PartSymbol_stagedOrder, stage.PartSymbol_orderStaged, &stage.PartSymbolOrder, stage.PartSymbols_mapString, partsymbol, partsymbol.Name)
	return partsymbol
}

// StagePreserveOrder puts partsymbol to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PartSymbolOrder
// - update stage.PartSymbolOrder accordingly
func (partsymbol *PartSymbol) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.PartSymbols, stage.PartSymbol_stagedOrder, stage.PartSymbol_orderStaged, &stage.PartSymbolOrder, stage.PartSymbols_mapString, partsymbol, order, partsymbol.Name)
}

// Unstage removes partsymbol off the model stage
func (partsymbol *PartSymbol) Unstage(stage *Stage) *PartSymbol {
	__gong__unstage(stage.PartSymbols, stage.PartSymbols_mapString, partsymbol, partsymbol.Name)
	return partsymbol
}

// UnstageVoid removes partsymbol off the model stage
func (partsymbol *PartSymbol) UnstageVoid(stage *Stage) {
	partsymbol.Unstage(stage)
}

func (partsymbol *PartSymbol) StageVoid(stage *Stage) {
	partsymbol.Stage(stage)
}

// for satisfaction of GongStruct interface
func (partsymbol *PartSymbol) GetName() (res string) {
	return partsymbol.Name
}

// for satisfaction of GongStruct interface
func (partsymbol *PartSymbol) SetName(name string) {
	partsymbol.Name = name
}

// Stage puts parttranspose to the model stage
func (parttranspose *PartTranspose) Stage(stage *Stage) *PartTranspose {
	__gong__stage(stage.PartTransposes, stage.PartTranspose_stagedOrder, stage.PartTranspose_orderStaged, &stage.PartTransposeOrder, stage.PartTransposes_mapString, parttranspose, parttranspose.Name)
	return parttranspose
}

// StagePreserveOrder puts parttranspose to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PartTransposeOrder
// - update stage.PartTransposeOrder accordingly
func (parttranspose *PartTranspose) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.PartTransposes, stage.PartTranspose_stagedOrder, stage.PartTranspose_orderStaged, &stage.PartTransposeOrder, stage.PartTransposes_mapString, parttranspose, order, parttranspose.Name)
}

// Unstage removes parttranspose off the model stage
func (parttranspose *PartTranspose) Unstage(stage *Stage) *PartTranspose {
	__gong__unstage(stage.PartTransposes, stage.PartTransposes_mapString, parttranspose, parttranspose.Name)
	return parttranspose
}

// UnstageVoid removes parttranspose off the model stage
func (parttranspose *PartTranspose) UnstageVoid(stage *Stage) {
	parttranspose.Unstage(stage)
}

func (parttranspose *PartTranspose) StageVoid(stage *Stage) {
	parttranspose.Stage(stage)
}

// for satisfaction of GongStruct interface
func (parttranspose *PartTranspose) GetName() (res string) {
	return parttranspose.Name
}

// for satisfaction of GongStruct interface
func (parttranspose *PartTranspose) SetName(name string) {
	parttranspose.Name = name
}

// Stage puts slash to the model stage
func (slash *Slash) Stage(stage *Stage) *Slash {
	__gong__stage(stage.Slashs, stage.Slash_stagedOrder, stage.Slash_orderStaged, &stage.SlashOrder, stage.Slashs_mapString, slash, slash.Name)
	return slash
}

// StagePreserveOrder puts slash to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SlashOrder
// - update stage.SlashOrder accordingly
func (slash *Slash) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Slashs, stage.Slash_stagedOrder, stage.Slash_orderStaged, &stage.SlashOrder, stage.Slashs_mapString, slash, order, slash.Name)
}

// Unstage removes slash off the model stage
func (slash *Slash) Unstage(stage *Stage) *Slash {
	__gong__unstage(stage.Slashs, stage.Slashs_mapString, slash, slash.Name)
	return slash
}

// UnstageVoid removes slash off the model stage
func (slash *Slash) UnstageVoid(stage *Stage) {
	slash.Unstage(stage)
}

func (slash *Slash) StageVoid(stage *Stage) {
	slash.Stage(stage)
}

// for satisfaction of GongStruct interface
func (slash *Slash) GetName() (res string) {
	return slash.Name
}

// for satisfaction of GongStruct interface
func (slash *Slash) SetName(name string) {
	slash.Name = name
}

// Stage puts staffdetails to the model stage
func (staffdetails *StaffDetails) Stage(stage *Stage) *StaffDetails {
	__gong__stage(stage.StaffDetailss, stage.StaffDetails_stagedOrder, stage.StaffDetails_orderStaged, &stage.StaffDetailsOrder, stage.StaffDetailss_mapString, staffdetails, staffdetails.Name)
	return staffdetails
}

// StagePreserveOrder puts staffdetails to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StaffDetailsOrder
// - update stage.StaffDetailsOrder accordingly
func (staffdetails *StaffDetails) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.StaffDetailss, stage.StaffDetails_stagedOrder, stage.StaffDetails_orderStaged, &stage.StaffDetailsOrder, stage.StaffDetailss_mapString, staffdetails, order, staffdetails.Name)
}

// Unstage removes staffdetails off the model stage
func (staffdetails *StaffDetails) Unstage(stage *Stage) *StaffDetails {
	__gong__unstage(stage.StaffDetailss, stage.StaffDetailss_mapString, staffdetails, staffdetails.Name)
	return staffdetails
}

// UnstageVoid removes staffdetails off the model stage
func (staffdetails *StaffDetails) UnstageVoid(stage *Stage) {
	staffdetails.Unstage(stage)
}

func (staffdetails *StaffDetails) StageVoid(stage *Stage) {
	staffdetails.Stage(stage)
}

// for satisfaction of GongStruct interface
func (staffdetails *StaffDetails) GetName() (res string) {
	return staffdetails.Name
}

// for satisfaction of GongStruct interface
func (staffdetails *StaffDetails) SetName(name string) {
	staffdetails.Name = name
}

// Stage puts staffsize to the model stage
func (staffsize *StaffSize) Stage(stage *Stage) *StaffSize {
	__gong__stage(stage.StaffSizes, stage.StaffSize_stagedOrder, stage.StaffSize_orderStaged, &stage.StaffSizeOrder, stage.StaffSizes_mapString, staffsize, staffsize.Name)
	return staffsize
}

// StagePreserveOrder puts staffsize to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StaffSizeOrder
// - update stage.StaffSizeOrder accordingly
func (staffsize *StaffSize) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.StaffSizes, stage.StaffSize_stagedOrder, stage.StaffSize_orderStaged, &stage.StaffSizeOrder, stage.StaffSizes_mapString, staffsize, order, staffsize.Name)
}

// Unstage removes staffsize off the model stage
func (staffsize *StaffSize) Unstage(stage *Stage) *StaffSize {
	__gong__unstage(stage.StaffSizes, stage.StaffSizes_mapString, staffsize, staffsize.Name)
	return staffsize
}

// UnstageVoid removes staffsize off the model stage
func (staffsize *StaffSize) UnstageVoid(stage *Stage) {
	staffsize.Unstage(stage)
}

func (staffsize *StaffSize) StageVoid(stage *Stage) {
	staffsize.Stage(stage)
}

// for satisfaction of GongStruct interface
func (staffsize *StaffSize) GetName() (res string) {
	return staffsize.Name
}

// for satisfaction of GongStruct interface
func (staffsize *StaffSize) SetName(name string) {
	staffsize.Name = name
}

// Stage puts stafftuning to the model stage
func (stafftuning *StaffTuning) Stage(stage *Stage) *StaffTuning {
	__gong__stage(stage.StaffTunings, stage.StaffTuning_stagedOrder, stage.StaffTuning_orderStaged, &stage.StaffTuningOrder, stage.StaffTunings_mapString, stafftuning, stafftuning.Name)
	return stafftuning
}

// StagePreserveOrder puts stafftuning to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StaffTuningOrder
// - update stage.StaffTuningOrder accordingly
func (stafftuning *StaffTuning) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.StaffTunings, stage.StaffTuning_stagedOrder, stage.StaffTuning_orderStaged, &stage.StaffTuningOrder, stage.StaffTunings_mapString, stafftuning, order, stafftuning.Name)
}

// Unstage removes stafftuning off the model stage
func (stafftuning *StaffTuning) Unstage(stage *Stage) *StaffTuning {
	__gong__unstage(stage.StaffTunings, stage.StaffTunings_mapString, stafftuning, stafftuning.Name)
	return stafftuning
}

// UnstageVoid removes stafftuning off the model stage
func (stafftuning *StaffTuning) UnstageVoid(stage *Stage) {
	stafftuning.Unstage(stage)
}

func (stafftuning *StaffTuning) StageVoid(stage *Stage) {
	stafftuning.Stage(stage)
}

// for satisfaction of GongStruct interface
func (stafftuning *StaffTuning) GetName() (res string) {
	return stafftuning.Name
}

// for satisfaction of GongStruct interface
func (stafftuning *StaffTuning) SetName(name string) {
	stafftuning.Name = name
}

// Stage puts time to the model stage
func (time *Time) Stage(stage *Stage) *Time {
	__gong__stage(stage.Times, stage.Time_stagedOrder, stage.Time_orderStaged, &stage.TimeOrder, stage.Times_mapString, time, time.Name)
	return time
}

// StagePreserveOrder puts time to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TimeOrder
// - update stage.TimeOrder accordingly
func (time *Time) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Times, stage.Time_stagedOrder, stage.Time_orderStaged, &stage.TimeOrder, stage.Times_mapString, time, order, time.Name)
}

// Unstage removes time off the model stage
func (time *Time) Unstage(stage *Stage) *Time {
	__gong__unstage(stage.Times, stage.Times_mapString, time, time.Name)
	return time
}

// UnstageVoid removes time off the model stage
func (time *Time) UnstageVoid(stage *Stage) {
	time.Unstage(stage)
}

func (time *Time) StageVoid(stage *Stage) {
	time.Stage(stage)
}

// for satisfaction of GongStruct interface
func (time *Time) GetName() (res string) {
	return time.Name
}

// for satisfaction of GongStruct interface
func (time *Time) SetName(name string) {
	time.Name = name
}

// Stage puts transpose to the model stage
func (transpose *Transpose) Stage(stage *Stage) *Transpose {
	__gong__stage(stage.Transposes, stage.Transpose_stagedOrder, stage.Transpose_orderStaged, &stage.TransposeOrder, stage.Transposes_mapString, transpose, transpose.Name)
	return transpose
}

// StagePreserveOrder puts transpose to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TransposeOrder
// - update stage.TransposeOrder accordingly
func (transpose *Transpose) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Transposes, stage.Transpose_stagedOrder, stage.Transpose_orderStaged, &stage.TransposeOrder, stage.Transposes_mapString, transpose, order, transpose.Name)
}

// Unstage removes transpose off the model stage
func (transpose *Transpose) Unstage(stage *Stage) *Transpose {
	__gong__unstage(stage.Transposes, stage.Transposes_mapString, transpose, transpose.Name)
	return transpose
}

// UnstageVoid removes transpose off the model stage
func (transpose *Transpose) UnstageVoid(stage *Stage) {
	transpose.Unstage(stage)
}

func (transpose *Transpose) StageVoid(stage *Stage) {
	transpose.Stage(stage)
}

// for satisfaction of GongStruct interface
func (transpose *Transpose) GetName() (res string) {
	return transpose.Name
}

// for satisfaction of GongStruct interface
func (transpose *Transpose) SetName(name string) {
	transpose.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.Attributess, &stage.Attributess_mapString, &stage.Attributes_stagedOrder, &stage.AttributesOrder)

	__gong__resetStageType(&stage.AttributesDirectives, &stage.AttributesDirectives_mapString, &stage.AttributesDirective_stagedOrder, &stage.AttributesDirectiveOrder)

	__gong__resetStageType(&stage.BeatRepeats, &stage.BeatRepeats_mapString, &stage.BeatRepeat_stagedOrder, &stage.BeatRepeatOrder)

	__gong__resetStageType(&stage.Cancels, &stage.Cancels_mapString, &stage.Cancel_stagedOrder, &stage.CancelOrder)

	__gong__resetStageType(&stage.Clefs, &stage.Clefs_mapString, &stage.Clef_stagedOrder, &stage.ClefOrder)

	__gong__resetStageType(&stage.Doubles, &stage.Doubles_mapString, &stage.Double_stagedOrder, &stage.DoubleOrder)

	__gong__resetStageType(&stage.ForParts, &stage.ForParts_mapString, &stage.ForPart_stagedOrder, &stage.ForPartOrder)

	__gong__resetStageType(&stage.Interchangeables, &stage.Interchangeables_mapString, &stage.Interchangeable_stagedOrder, &stage.InterchangeableOrder)

	__gong__resetStageType(&stage.Keys, &stage.Keys_mapString, &stage.Key_stagedOrder, &stage.KeyOrder)

	__gong__resetStageType(&stage.KeyAccidentals, &stage.KeyAccidentals_mapString, &stage.KeyAccidental_stagedOrder, &stage.KeyAccidentalOrder)

	__gong__resetStageType(&stage.KeyOctaves, &stage.KeyOctaves_mapString, &stage.KeyOctave_stagedOrder, &stage.KeyOctaveOrder)

	__gong__resetStageType(&stage.LineDetails, &stage.LineDetails_mapString, &stage.LineDetail_stagedOrder, &stage.LineDetailOrder)

	__gong__resetStageType(&stage.MeasureRepeats, &stage.MeasureRepeats_mapString, &stage.MeasureRepeat_stagedOrder, &stage.MeasureRepeatOrder)

	__gong__resetStageType(&stage.MeasureStyles, &stage.MeasureStyles_mapString, &stage.MeasureStyle_stagedOrder, &stage.MeasureStyleOrder)

	__gong__resetStageType(&stage.MultipleRests, &stage.MultipleRests_mapString, &stage.MultipleRest_stagedOrder, &stage.MultipleRestOrder)

	__gong__resetStageType(&stage.PartClefs, &stage.PartClefs_mapString, &stage.PartClef_stagedOrder, &stage.PartClefOrder)

	__gong__resetStageType(&stage.PartSymbols, &stage.PartSymbols_mapString, &stage.PartSymbol_stagedOrder, &stage.PartSymbolOrder)

	__gong__resetStageType(&stage.PartTransposes, &stage.PartTransposes_mapString, &stage.PartTranspose_stagedOrder, &stage.PartTransposeOrder)

	__gong__resetStageType(&stage.Slashs, &stage.Slashs_mapString, &stage.Slash_stagedOrder, &stage.SlashOrder)

	__gong__resetStageType(&stage.StaffDetailss, &stage.StaffDetailss_mapString, &stage.StaffDetails_stagedOrder, &stage.StaffDetailsOrder)

	__gong__resetStageType(&stage.StaffSizes, &stage.StaffSizes_mapString, &stage.StaffSize_stagedOrder, &stage.StaffSizeOrder)

	__gong__resetStageType(&stage.StaffTunings, &stage.StaffTunings_mapString, &stage.StaffTuning_stagedOrder, &stage.StaffTuningOrder)

	__gong__resetStageType(&stage.Times, &stage.Times_mapString, &stage.Time_stagedOrder, &stage.TimeOrder)

	__gong__resetStageType(&stage.Transposes, &stage.Transposes_mapString, &stage.Transpose_stagedOrder, &stage.TransposeOrder)

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
func (Attributes) GongGetAssociationName() any {
	return &Attributes{
			Key: []*Key{{Name: "Key"}},
			Time: []*Time{{Name: "Time"}},
			PartSymbol: &PartSymbol{Name: "PartSymbol"},
			Clef: []*Clef{{Name: "Clef"}},
			StaffDetails: []*StaffDetails{{Name: "StaffDetails"}},
			Transpose: []*Transpose{{Name: "Transpose"}},
			ForPart: []*ForPart{{Name: "ForPart"}},
			Directive: []*AttributesDirective{{Name: "Directive"}},
			MeasureStyle: []*MeasureStyle{{Name: "MeasureStyle"}},
	}
}

func (AttributesDirective) GongGetAssociationName() any {
	return &AttributesDirective{
	}
}

func (BeatRepeat) GongGetAssociationName() any {
	return &BeatRepeat{
	}
}

func (Cancel) GongGetAssociationName() any {
	return &Cancel{
	}
}

func (Clef) GongGetAssociationName() any {
	return &Clef{
	}
}

func (Double) GongGetAssociationName() any {
	return &Double{
	}
}

func (ForPart) GongGetAssociationName() any {
	return &ForPart{
			PartClef: &PartClef{Name: "PartClef"},
			PartTranspose: &PartTranspose{Name: "PartTranspose"},
	}
}

func (Interchangeable) GongGetAssociationName() any {
	return &Interchangeable{
	}
}

func (Key) GongGetAssociationName() any {
	return &Key{
			Cancel: &Cancel{Name: "Cancel"},
			KeyAccidental: []*KeyAccidental{{Name: "KeyAccidental"}},
			KeyOctave: []*KeyOctave{{Name: "KeyOctave"}},
	}
}

func (KeyAccidental) GongGetAssociationName() any {
	return &KeyAccidental{
	}
}

func (KeyOctave) GongGetAssociationName() any {
	return &KeyOctave{
	}
}

func (LineDetail) GongGetAssociationName() any {
	return &LineDetail{
	}
}

func (MeasureRepeat) GongGetAssociationName() any {
	return &MeasureRepeat{
	}
}

func (MeasureStyle) GongGetAssociationName() any {
	return &MeasureStyle{
			MultipleRest: &MultipleRest{Name: "MultipleRest"},
			MeasureRepeat: &MeasureRepeat{Name: "MeasureRepeat"},
			BeatRepeat: &BeatRepeat{Name: "BeatRepeat"},
			Slash: &Slash{Name: "Slash"},
	}
}

func (MultipleRest) GongGetAssociationName() any {
	return &MultipleRest{
	}
}

func (PartClef) GongGetAssociationName() any {
	return &PartClef{
	}
}

func (PartSymbol) GongGetAssociationName() any {
	return &PartSymbol{
	}
}

func (PartTranspose) GongGetAssociationName() any {
	return &PartTranspose{
			Double: &Double{Name: "Double"},
	}
}

func (Slash) GongGetAssociationName() any {
	return &Slash{
	}
}

func (StaffDetails) GongGetAssociationName() any {
	return &StaffDetails{
			LineDetail: []*LineDetail{{Name: "LineDetail"}},
			StaffTuning: []*StaffTuning{{Name: "StaffTuning"}},
			StaffSize: &StaffSize{Name: "StaffSize"},
	}
}

func (StaffSize) GongGetAssociationName() any {
	return &StaffSize{
	}
}

func (StaffTuning) GongGetAssociationName() any {
	return &StaffTuning{
	}
}

func (Time) GongGetAssociationName() any {
	return &Time{
			Interchangeable: &Interchangeable{Name: "Interchangeable"},
	}
}

func (Transpose) GongGetAssociationName() any {
	return &Transpose{
			Double: &Double{Name: "Double"},
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
	// reverse maps of direct associations of Attributes
	case Attributes:
		switch fieldname {
		// insertion point for per direct association field
		case "PartSymbol":
			res := make(map[*PartSymbol][]*Attributes)
			for attributes := range stage.Attributess {
				if attributes.PartSymbol != nil {
					partsymbol_ := attributes.PartSymbol
					var attributess []*Attributes
					_, ok := res[partsymbol_]
					if ok {
						attributess = res[partsymbol_]
					} else {
						attributess = make([]*Attributes, 0)
					}
					attributess = append(attributess, attributes)
					res[partsymbol_] = attributess
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AttributesDirective
	case AttributesDirective:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BeatRepeat
	case BeatRepeat:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Cancel
	case Cancel:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Clef
	case Clef:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Double
	case Double:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ForPart
	case ForPart:
		switch fieldname {
		// insertion point for per direct association field
		case "PartClef":
			res := make(map[*PartClef][]*ForPart)
			for forpart := range stage.ForParts {
				if forpart.PartClef != nil {
					partclef_ := forpart.PartClef
					var forparts []*ForPart
					_, ok := res[partclef_]
					if ok {
						forparts = res[partclef_]
					} else {
						forparts = make([]*ForPart, 0)
					}
					forparts = append(forparts, forpart)
					res[partclef_] = forparts
				}
			}
			return any(res).(map[*End][]*Start)
		case "PartTranspose":
			res := make(map[*PartTranspose][]*ForPart)
			for forpart := range stage.ForParts {
				if forpart.PartTranspose != nil {
					parttranspose_ := forpart.PartTranspose
					var forparts []*ForPart
					_, ok := res[parttranspose_]
					if ok {
						forparts = res[parttranspose_]
					} else {
						forparts = make([]*ForPart, 0)
					}
					forparts = append(forparts, forpart)
					res[parttranspose_] = forparts
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Interchangeable
	case Interchangeable:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Key
	case Key:
		switch fieldname {
		// insertion point for per direct association field
		case "Cancel":
			res := make(map[*Cancel][]*Key)
			for key := range stage.Keys {
				if key.Cancel != nil {
					cancel_ := key.Cancel
					var keys []*Key
					_, ok := res[cancel_]
					if ok {
						keys = res[cancel_]
					} else {
						keys = make([]*Key, 0)
					}
					keys = append(keys, key)
					res[cancel_] = keys
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of KeyAccidental
	case KeyAccidental:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of KeyOctave
	case KeyOctave:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LineDetail
	case LineDetail:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MeasureRepeat
	case MeasureRepeat:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MeasureStyle
	case MeasureStyle:
		switch fieldname {
		// insertion point for per direct association field
		case "MultipleRest":
			res := make(map[*MultipleRest][]*MeasureStyle)
			for measurestyle := range stage.MeasureStyles {
				if measurestyle.MultipleRest != nil {
					multiplerest_ := measurestyle.MultipleRest
					var measurestyles []*MeasureStyle
					_, ok := res[multiplerest_]
					if ok {
						measurestyles = res[multiplerest_]
					} else {
						measurestyles = make([]*MeasureStyle, 0)
					}
					measurestyles = append(measurestyles, measurestyle)
					res[multiplerest_] = measurestyles
				}
			}
			return any(res).(map[*End][]*Start)
		case "MeasureRepeat":
			res := make(map[*MeasureRepeat][]*MeasureStyle)
			for measurestyle := range stage.MeasureStyles {
				if measurestyle.MeasureRepeat != nil {
					measurerepeat_ := measurestyle.MeasureRepeat
					var measurestyles []*MeasureStyle
					_, ok := res[measurerepeat_]
					if ok {
						measurestyles = res[measurerepeat_]
					} else {
						measurestyles = make([]*MeasureStyle, 0)
					}
					measurestyles = append(measurestyles, measurestyle)
					res[measurerepeat_] = measurestyles
				}
			}
			return any(res).(map[*End][]*Start)
		case "BeatRepeat":
			res := make(map[*BeatRepeat][]*MeasureStyle)
			for measurestyle := range stage.MeasureStyles {
				if measurestyle.BeatRepeat != nil {
					beatrepeat_ := measurestyle.BeatRepeat
					var measurestyles []*MeasureStyle
					_, ok := res[beatrepeat_]
					if ok {
						measurestyles = res[beatrepeat_]
					} else {
						measurestyles = make([]*MeasureStyle, 0)
					}
					measurestyles = append(measurestyles, measurestyle)
					res[beatrepeat_] = measurestyles
				}
			}
			return any(res).(map[*End][]*Start)
		case "Slash":
			res := make(map[*Slash][]*MeasureStyle)
			for measurestyle := range stage.MeasureStyles {
				if measurestyle.Slash != nil {
					slash_ := measurestyle.Slash
					var measurestyles []*MeasureStyle
					_, ok := res[slash_]
					if ok {
						measurestyles = res[slash_]
					} else {
						measurestyles = make([]*MeasureStyle, 0)
					}
					measurestyles = append(measurestyles, measurestyle)
					res[slash_] = measurestyles
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MultipleRest
	case MultipleRest:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PartClef
	case PartClef:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PartSymbol
	case PartSymbol:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PartTranspose
	case PartTranspose:
		switch fieldname {
		// insertion point for per direct association field
		case "Double":
			res := make(map[*Double][]*PartTranspose)
			for parttranspose := range stage.PartTransposes {
				if parttranspose.Double != nil {
					double_ := parttranspose.Double
					var parttransposes []*PartTranspose
					_, ok := res[double_]
					if ok {
						parttransposes = res[double_]
					} else {
						parttransposes = make([]*PartTranspose, 0)
					}
					parttransposes = append(parttransposes, parttranspose)
					res[double_] = parttransposes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Slash
	case Slash:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StaffDetails
	case StaffDetails:
		switch fieldname {
		// insertion point for per direct association field
		case "StaffSize":
			res := make(map[*StaffSize][]*StaffDetails)
			for staffdetails := range stage.StaffDetailss {
				if staffdetails.StaffSize != nil {
					staffsize_ := staffdetails.StaffSize
					var staffdetailss []*StaffDetails
					_, ok := res[staffsize_]
					if ok {
						staffdetailss = res[staffsize_]
					} else {
						staffdetailss = make([]*StaffDetails, 0)
					}
					staffdetailss = append(staffdetailss, staffdetails)
					res[staffsize_] = staffdetailss
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of StaffSize
	case StaffSize:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StaffTuning
	case StaffTuning:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Time
	case Time:
		switch fieldname {
		// insertion point for per direct association field
		case "Interchangeable":
			res := make(map[*Interchangeable][]*Time)
			for time := range stage.Times {
				if time.Interchangeable != nil {
					interchangeable_ := time.Interchangeable
					var times []*Time
					_, ok := res[interchangeable_]
					if ok {
						times = res[interchangeable_]
					} else {
						times = make([]*Time, 0)
					}
					times = append(times, time)
					res[interchangeable_] = times
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Transpose
	case Transpose:
		switch fieldname {
		// insertion point for per direct association field
		case "Double":
			res := make(map[*Double][]*Transpose)
			for transpose := range stage.Transposes {
				if transpose.Double != nil {
					double_ := transpose.Double
					var transposes []*Transpose
					_, ok := res[double_]
					if ok {
						transposes = res[double_]
					} else {
						transposes = make([]*Transpose, 0)
					}
					transposes = append(transposes, transpose)
					res[double_] = transposes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap is the Stage method for backtrack navigation of slice-of-pointers associations.
func (stage *Stage) GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string) map[*End][]*Start {
	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Attributes
	case Attributes:
		switch fieldname {
		// insertion point for per direct association field
		case "Key":
			res := make(map[*Key][]*Attributes)
			for attributes := range stage.Attributess {
				for _, key_ := range attributes.Key {
					res[key_] = append(res[key_], attributes)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Time":
			res := make(map[*Time][]*Attributes)
			for attributes := range stage.Attributess {
				for _, time_ := range attributes.Time {
					res[time_] = append(res[time_], attributes)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Clef":
			res := make(map[*Clef][]*Attributes)
			for attributes := range stage.Attributess {
				for _, clef_ := range attributes.Clef {
					res[clef_] = append(res[clef_], attributes)
				}
			}
			return any(res).(map[*End][]*Start)
		case "StaffDetails":
			res := make(map[*StaffDetails][]*Attributes)
			for attributes := range stage.Attributess {
				for _, staffdetails_ := range attributes.StaffDetails {
					res[staffdetails_] = append(res[staffdetails_], attributes)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Transpose":
			res := make(map[*Transpose][]*Attributes)
			for attributes := range stage.Attributess {
				for _, transpose_ := range attributes.Transpose {
					res[transpose_] = append(res[transpose_], attributes)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ForPart":
			res := make(map[*ForPart][]*Attributes)
			for attributes := range stage.Attributess {
				for _, forpart_ := range attributes.ForPart {
					res[forpart_] = append(res[forpart_], attributes)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Directive":
			res := make(map[*AttributesDirective][]*Attributes)
			for attributes := range stage.Attributess {
				for _, attributesdirective_ := range attributes.Directive {
					res[attributesdirective_] = append(res[attributesdirective_], attributes)
				}
			}
			return any(res).(map[*End][]*Start)
		case "MeasureStyle":
			res := make(map[*MeasureStyle][]*Attributes)
			for attributes := range stage.Attributess {
				for _, measurestyle_ := range attributes.MeasureStyle {
					res[measurestyle_] = append(res[measurestyle_], attributes)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AttributesDirective
	case AttributesDirective:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BeatRepeat
	case BeatRepeat:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Cancel
	case Cancel:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Clef
	case Clef:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Double
	case Double:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ForPart
	case ForPart:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Interchangeable
	case Interchangeable:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Key
	case Key:
		switch fieldname {
		// insertion point for per direct association field
		case "KeyAccidental":
			res := make(map[*KeyAccidental][]*Key)
			for key := range stage.Keys {
				for _, keyaccidental_ := range key.KeyAccidental {
					res[keyaccidental_] = append(res[keyaccidental_], key)
				}
			}
			return any(res).(map[*End][]*Start)
		case "KeyOctave":
			res := make(map[*KeyOctave][]*Key)
			for key := range stage.Keys {
				for _, keyoctave_ := range key.KeyOctave {
					res[keyoctave_] = append(res[keyoctave_], key)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of KeyAccidental
	case KeyAccidental:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of KeyOctave
	case KeyOctave:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LineDetail
	case LineDetail:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MeasureRepeat
	case MeasureRepeat:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MeasureStyle
	case MeasureStyle:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MultipleRest
	case MultipleRest:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PartClef
	case PartClef:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PartSymbol
	case PartSymbol:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PartTranspose
	case PartTranspose:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Slash
	case Slash:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StaffDetails
	case StaffDetails:
		switch fieldname {
		// insertion point for per direct association field
		case "LineDetail":
			res := make(map[*LineDetail][]*StaffDetails)
			for staffdetails := range stage.StaffDetailss {
				for _, linedetail_ := range staffdetails.LineDetail {
					res[linedetail_] = append(res[linedetail_], staffdetails)
				}
			}
			return any(res).(map[*End][]*Start)
		case "StaffTuning":
			res := make(map[*StaffTuning][]*StaffDetails)
			for staffdetails := range stage.StaffDetailss {
				for _, stafftuning_ := range staffdetails.StaffTuning {
					res[stafftuning_] = append(res[stafftuning_], staffdetails)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of StaffSize
	case StaffSize:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StaffTuning
	case StaffTuning:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Time
	case Time:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Transpose
	case Transpose:
		switch fieldname {
		// insertion point for per direct association field
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
func (*Attributes) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AttributesDirective) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Attributes",
			Fieldname: "Directive",
		},
	}
}

func (*BeatRepeat) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Cancel) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Clef) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Attributes",
			Fieldname: "Clef",
		},
	}
}

func (*Double) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ForPart) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Attributes",
			Fieldname: "ForPart",
		},
	}
}

func (*Interchangeable) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Key) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Attributes",
			Fieldname: "Key",
		},
	}
}

func (*KeyAccidental) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Key",
			Fieldname: "KeyAccidental",
		},
	}
}

func (*KeyOctave) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Key",
			Fieldname: "KeyOctave",
		},
	}
}

func (*LineDetail) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "StaffDetails",
			Fieldname: "LineDetail",
		},
	}
}

func (*MeasureRepeat) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MeasureStyle) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Attributes",
			Fieldname: "MeasureStyle",
		},
	}
}

func (*MultipleRest) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*PartClef) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*PartSymbol) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*PartTranspose) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Slash) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*StaffDetails) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Attributes",
			Fieldname: "StaffDetails",
		},
	}
}

func (*StaffSize) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*StaffTuning) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "StaffDetails",
			Fieldname: "StaffTuning",
		},
	}
}

func (*Time) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Attributes",
			Fieldname: "Time",
		},
	}
}

func (*Transpose) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Attributes",
			Fieldname: "Transpose",
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
func (attributes *Attributes) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Divisions",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Key",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Key",
		},
		{
			Name:                 "Time",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Time",
		},
		{
			Name:               "Staves",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "PartSymbol",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PartSymbol",
		},
		{
			Name:               "Instruments",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Clef",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Clef",
		},
		{
			Name:                 "StaffDetails",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "StaffDetails",
		},
		{
			Name:                 "Transpose",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Transpose",
		},
		{
			Name:                 "ForPart",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ForPart",
		},
		{
			Name:                 "Directive",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AttributesDirective",
		},
		{
			Name:                 "MeasureStyle",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MeasureStyle",
		},
	}
	return
}

func (attributesdirective *AttributesDirective) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Lang",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (beatrepeat *BeatRepeat) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Slashes",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "ExceptVoice",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (cancel *Cancel) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Location",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumCancelLocation",
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (clef *Clef) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Sign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumClefSign",
		},
		{
			Name:               "Line",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "ClefOctaveChange",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (double *Double) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (forpart *ForPart) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "PartClef",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PartClef",
		},
		{
			Name:                 "PartTranspose",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PartTranspose",
		},
	}
	return
}

func (interchangeable *Interchangeable) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Symbol",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumTimeSymbol",
		},
		{
			Name:                 "Separator",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumTimeSeparator",
		},
		{
			Name:                 "TimeRelation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumTimeRelation",
		},
		{
			Name:               "Beats",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BeatType",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (key *Key) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Cancel",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Cancel",
		},
		{
			Name:               "Fifths",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Mode",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "KeyStep",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumStep",
		},
		{
			Name:               "KeyAlter",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "KeyAccidental",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "KeyAccidental",
		},
		{
			Name:                 "KeyOctave",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "KeyOctave",
		},
	}
	return
}

func (keyaccidental *KeyAccidental) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Smufl",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (keyoctave *KeyOctave) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (linedetail *LineDetail) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Line",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (measurerepeat *MeasureRepeat) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Slashes",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (measurestyle *MeasureStyle) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "MultipleRest",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MultipleRest",
		},
		{
			Name:                 "MeasureRepeat",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MeasureRepeat",
		},
		{
			Name:                 "BeatRepeat",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "BeatRepeat",
		},
		{
			Name:                 "Slash",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Slash",
		},
	}
	return
}

func (multiplerest *MultipleRest) GongGetFieldHeaders() (res []GongFieldHeader) {
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
	}
	return
}

func (partclef *PartClef) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Sign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumClefSign",
		},
		{
			Name:               "Line",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "ClefOctaveChange",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (partsymbol *PartSymbol) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "TopStaff",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "BottomStaff",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumGroupSymbolValue",
		},
	}
	return
}

func (parttranspose *PartTranspose) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Diatonic",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Chromatic",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "OctaveChange",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Double",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Double",
		},
	}
	return
}

func (slash *Slash) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ExceptVoice",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (staffdetails *StaffDetails) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "ShowFrets",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumShowFrets",
		},
		{
			Name:                 "StaffType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumStaffType",
		},
		{
			Name:               "StaffLines",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "LineDetail",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "LineDetail",
		},
		{
			Name:                 "StaffTuning",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "StaffTuning",
		},
		{
			Name:               "Capo",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "StaffSize",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "StaffSize",
		},
	}
	return
}

func (staffsize *StaffSize) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Scaling",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (stafftuning *StaffTuning) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Line",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "TuningStep",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumStep",
		},
		{
			Name:               "TuningAlter",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "TuningOctave",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (time *Time) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Symbol",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumTimeSymbol",
		},
		{
			Name:                 "Separator",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumTimeSeparator",
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Beats",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BeatType",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Interchangeable",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Interchangeable",
		},
		{
			Name:               "SenzaMisura",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (transpose *Transpose) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Diatonic",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Chromatic",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "OctaveChange",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Double",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Double",
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
func (attributes *Attributes) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = attributes.Name
	case "Divisions":
		res.valueString = attributes.Divisions
	case "Key":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range attributes.Key {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Time":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range attributes.Time {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Staves":
		res.valueString = fmt.Sprintf("%d", attributes.Staves)
		res.valueInt = attributes.Staves
		res.GongFieldValueType = GongFieldValueTypeInt
	case "PartSymbol":
		res.GongFieldValueType = GongFieldValueTypePointer
		if attributes.PartSymbol != nil {
			res.valueString = attributes.PartSymbol.Name
			res.ids = attributes.PartSymbol.GongGetUUID(stage)
		}
	case "Instruments":
		res.valueString = fmt.Sprintf("%d", attributes.Instruments)
		res.valueInt = attributes.Instruments
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Clef":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range attributes.Clef {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "StaffDetails":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range attributes.StaffDetails {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Transpose":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range attributes.Transpose {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ForPart":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range attributes.ForPart {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Directive":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range attributes.Directive {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "MeasureStyle":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range attributes.MeasureStyle {
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

func (attributesdirective *AttributesDirective) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = attributesdirective.Name
	case "Lang":
		res.valueString = attributesdirective.Lang
	case "DefaultX":
		res.valueString = attributesdirective.DefaultX
	case "DefaultY":
		res.valueString = attributesdirective.DefaultY
	case "RelativeX":
		res.valueString = attributesdirective.RelativeX
	case "RelativeY":
		res.valueString = attributesdirective.RelativeY
	case "FontFamily":
		res.valueString = attributesdirective.FontFamily
	case "FontSize":
		res.valueString = attributesdirective.FontSize
	case "Color":
		res.valueString = attributesdirective.Color
	case "Value":
		res.valueString = attributesdirective.Value
	}
	return
}

func (beatrepeat *BeatRepeat) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = beatrepeat.Name
	case "Slashes":
		res.valueString = fmt.Sprintf("%d", beatrepeat.Slashes)
		res.valueInt = beatrepeat.Slashes
		res.GongFieldValueType = GongFieldValueTypeInt
	case "ExceptVoice":
		res.valueString = beatrepeat.ExceptVoice
	}
	return
}

func (cancel *Cancel) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = cancel.Name
	case "Location":
		enum := cancel.Location
		res.valueString = enum.ToCodeString()
	case "Value":
		res.valueString = fmt.Sprintf("%d", cancel.Value)
		res.valueInt = cancel.Value
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (clef *Clef) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = clef.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", clef.Number)
		res.valueInt = clef.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "DefaultX":
		res.valueString = clef.DefaultX
	case "DefaultY":
		res.valueString = clef.DefaultY
	case "RelativeX":
		res.valueString = clef.RelativeX
	case "RelativeY":
		res.valueString = clef.RelativeY
	case "FontFamily":
		res.valueString = clef.FontFamily
	case "FontSize":
		res.valueString = clef.FontSize
	case "Color":
		res.valueString = clef.Color
	case "Id":
		res.valueString = clef.Id
	case "Sign":
		enum := clef.Sign
		res.valueString = enum.ToCodeString()
	case "Line":
		res.valueString = fmt.Sprintf("%d", clef.Line)
		res.valueInt = clef.Line
		res.GongFieldValueType = GongFieldValueTypeInt
	case "ClefOctaveChange":
		res.valueString = fmt.Sprintf("%d", clef.ClefOctaveChange)
		res.valueInt = clef.ClefOctaveChange
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (double *Double) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = double.Name
	}
	return
}

func (forpart *ForPart) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = forpart.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", forpart.Number)
		res.valueInt = forpart.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Id":
		res.valueString = forpart.Id
	case "PartClef":
		res.GongFieldValueType = GongFieldValueTypePointer
		if forpart.PartClef != nil {
			res.valueString = forpart.PartClef.Name
			res.ids = forpart.PartClef.GongGetUUID(stage)
		}
	case "PartTranspose":
		res.GongFieldValueType = GongFieldValueTypePointer
		if forpart.PartTranspose != nil {
			res.valueString = forpart.PartTranspose.Name
			res.ids = forpart.PartTranspose.GongGetUUID(stage)
		}
	}
	return
}

func (interchangeable *Interchangeable) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = interchangeable.Name
	case "Symbol":
		enum := interchangeable.Symbol
		res.valueString = enum.ToCodeString()
	case "Separator":
		enum := interchangeable.Separator
		res.valueString = enum.ToCodeString()
	case "TimeRelation":
		enum := interchangeable.TimeRelation
		res.valueString = enum.ToCodeString()
	case "Beats":
		res.valueString = interchangeable.Beats
	case "BeatType":
		res.valueString = interchangeable.BeatType
	}
	return
}

func (key *Key) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = key.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", key.Number)
		res.valueInt = key.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "DefaultX":
		res.valueString = key.DefaultX
	case "DefaultY":
		res.valueString = key.DefaultY
	case "RelativeX":
		res.valueString = key.RelativeX
	case "RelativeY":
		res.valueString = key.RelativeY
	case "FontFamily":
		res.valueString = key.FontFamily
	case "FontSize":
		res.valueString = key.FontSize
	case "Color":
		res.valueString = key.Color
	case "Id":
		res.valueString = key.Id
	case "Cancel":
		res.GongFieldValueType = GongFieldValueTypePointer
		if key.Cancel != nil {
			res.valueString = key.Cancel.Name
			res.ids = key.Cancel.GongGetUUID(stage)
		}
	case "Fifths":
		res.valueString = fmt.Sprintf("%d", key.Fifths)
		res.valueInt = key.Fifths
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Mode":
		res.valueString = key.Mode
	case "KeyStep":
		enum := key.KeyStep
		res.valueString = enum.ToCodeString()
	case "KeyAlter":
		res.valueString = key.KeyAlter
	case "KeyAccidental":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range key.KeyAccidental {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "KeyOctave":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range key.KeyOctave {
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

func (keyaccidental *KeyAccidental) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = keyaccidental.Name
	case "Smufl":
		res.valueString = keyaccidental.Smufl
	}
	return
}

func (keyoctave *KeyOctave) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = keyoctave.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", keyoctave.Number)
		res.valueInt = keyoctave.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Value":
		res.valueString = fmt.Sprintf("%d", keyoctave.Value)
		res.valueInt = keyoctave.Value
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (linedetail *LineDetail) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = linedetail.Name
	case "Line":
		res.valueString = fmt.Sprintf("%d", linedetail.Line)
		res.valueInt = linedetail.Line
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Width":
		res.valueString = linedetail.Width
	case "Color":
		res.valueString = linedetail.Color
	}
	return
}

func (measurerepeat *MeasureRepeat) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = measurerepeat.Name
	case "Slashes":
		res.valueString = fmt.Sprintf("%d", measurerepeat.Slashes)
		res.valueInt = measurerepeat.Slashes
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Value":
		res.valueString = measurerepeat.Value
	}
	return
}

func (measurestyle *MeasureStyle) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = measurestyle.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", measurestyle.Number)
		res.valueInt = measurestyle.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "FontFamily":
		res.valueString = measurestyle.FontFamily
	case "FontSize":
		res.valueString = measurestyle.FontSize
	case "Color":
		res.valueString = measurestyle.Color
	case "Id":
		res.valueString = measurestyle.Id
	case "MultipleRest":
		res.GongFieldValueType = GongFieldValueTypePointer
		if measurestyle.MultipleRest != nil {
			res.valueString = measurestyle.MultipleRest.Name
			res.ids = measurestyle.MultipleRest.GongGetUUID(stage)
		}
	case "MeasureRepeat":
		res.GongFieldValueType = GongFieldValueTypePointer
		if measurestyle.MeasureRepeat != nil {
			res.valueString = measurestyle.MeasureRepeat.Name
			res.ids = measurestyle.MeasureRepeat.GongGetUUID(stage)
		}
	case "BeatRepeat":
		res.GongFieldValueType = GongFieldValueTypePointer
		if measurestyle.BeatRepeat != nil {
			res.valueString = measurestyle.BeatRepeat.Name
			res.ids = measurestyle.BeatRepeat.GongGetUUID(stage)
		}
	case "Slash":
		res.GongFieldValueType = GongFieldValueTypePointer
		if measurestyle.Slash != nil {
			res.valueString = measurestyle.Slash.Name
			res.ids = measurestyle.Slash.GongGetUUID(stage)
		}
	}
	return
}

func (multiplerest *MultipleRest) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = multiplerest.Name
	case "Value":
		res.valueString = fmt.Sprintf("%d", multiplerest.Value)
		res.valueInt = multiplerest.Value
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (partclef *PartClef) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = partclef.Name
	case "Sign":
		enum := partclef.Sign
		res.valueString = enum.ToCodeString()
	case "Line":
		res.valueString = fmt.Sprintf("%d", partclef.Line)
		res.valueInt = partclef.Line
		res.GongFieldValueType = GongFieldValueTypeInt
	case "ClefOctaveChange":
		res.valueString = fmt.Sprintf("%d", partclef.ClefOctaveChange)
		res.valueInt = partclef.ClefOctaveChange
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (partsymbol *PartSymbol) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = partsymbol.Name
	case "TopStaff":
		res.valueString = fmt.Sprintf("%d", partsymbol.TopStaff)
		res.valueInt = partsymbol.TopStaff
		res.GongFieldValueType = GongFieldValueTypeInt
	case "BottomStaff":
		res.valueString = fmt.Sprintf("%d", partsymbol.BottomStaff)
		res.valueInt = partsymbol.BottomStaff
		res.GongFieldValueType = GongFieldValueTypeInt
	case "DefaultX":
		res.valueString = partsymbol.DefaultX
	case "DefaultY":
		res.valueString = partsymbol.DefaultY
	case "RelativeX":
		res.valueString = partsymbol.RelativeX
	case "RelativeY":
		res.valueString = partsymbol.RelativeY
	case "Color":
		res.valueString = partsymbol.Color
	case "Value":
		enum := partsymbol.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (parttranspose *PartTranspose) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = parttranspose.Name
	case "Diatonic":
		res.valueString = fmt.Sprintf("%d", parttranspose.Diatonic)
		res.valueInt = parttranspose.Diatonic
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Chromatic":
		res.valueString = parttranspose.Chromatic
	case "OctaveChange":
		res.valueString = fmt.Sprintf("%d", parttranspose.OctaveChange)
		res.valueInt = parttranspose.OctaveChange
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Double":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parttranspose.Double != nil {
			res.valueString = parttranspose.Double.Name
			res.ids = parttranspose.Double.GongGetUUID(stage)
		}
	}
	return
}

func (slash *Slash) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = slash.Name
	case "ExceptVoice":
		res.valueString = slash.ExceptVoice
	}
	return
}

func (staffdetails *StaffDetails) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = staffdetails.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", staffdetails.Number)
		res.valueInt = staffdetails.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "ShowFrets":
		enum := staffdetails.ShowFrets
		res.valueString = enum.ToCodeString()
	case "StaffType":
		enum := staffdetails.StaffType
		res.valueString = enum.ToCodeString()
	case "StaffLines":
		res.valueString = fmt.Sprintf("%d", staffdetails.StaffLines)
		res.valueInt = staffdetails.StaffLines
		res.GongFieldValueType = GongFieldValueTypeInt
	case "LineDetail":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range staffdetails.LineDetail {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "StaffTuning":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range staffdetails.StaffTuning {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Capo":
		res.valueString = fmt.Sprintf("%d", staffdetails.Capo)
		res.valueInt = staffdetails.Capo
		res.GongFieldValueType = GongFieldValueTypeInt
	case "StaffSize":
		res.GongFieldValueType = GongFieldValueTypePointer
		if staffdetails.StaffSize != nil {
			res.valueString = staffdetails.StaffSize.Name
			res.ids = staffdetails.StaffSize.GongGetUUID(stage)
		}
	}
	return
}

func (staffsize *StaffSize) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = staffsize.Name
	case "Scaling":
		res.valueString = staffsize.Scaling
	case "Value":
		res.valueString = staffsize.Value
	}
	return
}

func (stafftuning *StaffTuning) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = stafftuning.Name
	case "Line":
		res.valueString = fmt.Sprintf("%d", stafftuning.Line)
		res.valueInt = stafftuning.Line
		res.GongFieldValueType = GongFieldValueTypeInt
	case "TuningStep":
		enum := stafftuning.TuningStep
		res.valueString = enum.ToCodeString()
	case "TuningAlter":
		res.valueString = stafftuning.TuningAlter
	case "TuningOctave":
		res.valueString = fmt.Sprintf("%d", stafftuning.TuningOctave)
		res.valueInt = stafftuning.TuningOctave
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (time *Time) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = time.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", time.Number)
		res.valueInt = time.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Symbol":
		enum := time.Symbol
		res.valueString = enum.ToCodeString()
	case "Separator":
		enum := time.Separator
		res.valueString = enum.ToCodeString()
	case "DefaultX":
		res.valueString = time.DefaultX
	case "DefaultY":
		res.valueString = time.DefaultY
	case "RelativeX":
		res.valueString = time.RelativeX
	case "RelativeY":
		res.valueString = time.RelativeY
	case "FontFamily":
		res.valueString = time.FontFamily
	case "FontSize":
		res.valueString = time.FontSize
	case "Color":
		res.valueString = time.Color
	case "Id":
		res.valueString = time.Id
	case "Beats":
		res.valueString = time.Beats
	case "BeatType":
		res.valueString = time.BeatType
	case "Interchangeable":
		res.GongFieldValueType = GongFieldValueTypePointer
		if time.Interchangeable != nil {
			res.valueString = time.Interchangeable.Name
			res.ids = time.Interchangeable.GongGetUUID(stage)
		}
	case "SenzaMisura":
		res.valueString = time.SenzaMisura
	}
	return
}

func (transpose *Transpose) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = transpose.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", transpose.Number)
		res.valueInt = transpose.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Id":
		res.valueString = transpose.Id
	case "Diatonic":
		res.valueString = fmt.Sprintf("%d", transpose.Diatonic)
		res.valueInt = transpose.Diatonic
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Chromatic":
		res.valueString = transpose.Chromatic
	case "OctaveChange":
		res.valueString = fmt.Sprintf("%d", transpose.OctaveChange)
		res.valueInt = transpose.OctaveChange
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Double":
		res.GongFieldValueType = GongFieldValueTypePointer
		if transpose.Double != nil {
			res.valueString = transpose.Double.Name
			res.ids = transpose.Double.GongGetUUID(stage)
		}
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
func (attributes *Attributes) GongGetGongstructName() string {
	return "Attributes"
}

func (attributesdirective *AttributesDirective) GongGetGongstructName() string {
	return "AttributesDirective"
}

func (beatrepeat *BeatRepeat) GongGetGongstructName() string {
	return "BeatRepeat"
}

func (cancel *Cancel) GongGetGongstructName() string {
	return "Cancel"
}

func (clef *Clef) GongGetGongstructName() string {
	return "Clef"
}

func (double *Double) GongGetGongstructName() string {
	return "Double"
}

func (forpart *ForPart) GongGetGongstructName() string {
	return "ForPart"
}

func (interchangeable *Interchangeable) GongGetGongstructName() string {
	return "Interchangeable"
}

func (key *Key) GongGetGongstructName() string {
	return "Key"
}

func (keyaccidental *KeyAccidental) GongGetGongstructName() string {
	return "KeyAccidental"
}

func (keyoctave *KeyOctave) GongGetGongstructName() string {
	return "KeyOctave"
}

func (linedetail *LineDetail) GongGetGongstructName() string {
	return "LineDetail"
}

func (measurerepeat *MeasureRepeat) GongGetGongstructName() string {
	return "MeasureRepeat"
}

func (measurestyle *MeasureStyle) GongGetGongstructName() string {
	return "MeasureStyle"
}

func (multiplerest *MultipleRest) GongGetGongstructName() string {
	return "MultipleRest"
}

func (partclef *PartClef) GongGetGongstructName() string {
	return "PartClef"
}

func (partsymbol *PartSymbol) GongGetGongstructName() string {
	return "PartSymbol"
}

func (parttranspose *PartTranspose) GongGetGongstructName() string {
	return "PartTranspose"
}

func (slash *Slash) GongGetGongstructName() string {
	return "Slash"
}

func (staffdetails *StaffDetails) GongGetGongstructName() string {
	return "StaffDetails"
}

func (staffsize *StaffSize) GongGetGongstructName() string {
	return "StaffSize"
}

func (stafftuning *StaffTuning) GongGetGongstructName() string {
	return "StaffTuning"
}

func (time *Time) GongGetGongstructName() string {
	return "Time"
}

func (transpose *Transpose) GongGetGongstructName() string {
	return "Transpose"
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
	__gong__rebuildMapString(stage.Attributess, &stage.Attributess_mapString)

	__gong__rebuildMapString(stage.AttributesDirectives, &stage.AttributesDirectives_mapString)

	__gong__rebuildMapString(stage.BeatRepeats, &stage.BeatRepeats_mapString)

	__gong__rebuildMapString(stage.Cancels, &stage.Cancels_mapString)

	__gong__rebuildMapString(stage.Clefs, &stage.Clefs_mapString)

	__gong__rebuildMapString(stage.Doubles, &stage.Doubles_mapString)

	__gong__rebuildMapString(stage.ForParts, &stage.ForParts_mapString)

	__gong__rebuildMapString(stage.Interchangeables, &stage.Interchangeables_mapString)

	__gong__rebuildMapString(stage.Keys, &stage.Keys_mapString)

	__gong__rebuildMapString(stage.KeyAccidentals, &stage.KeyAccidentals_mapString)

	__gong__rebuildMapString(stage.KeyOctaves, &stage.KeyOctaves_mapString)

	__gong__rebuildMapString(stage.LineDetails, &stage.LineDetails_mapString)

	__gong__rebuildMapString(stage.MeasureRepeats, &stage.MeasureRepeats_mapString)

	__gong__rebuildMapString(stage.MeasureStyles, &stage.MeasureStyles_mapString)

	__gong__rebuildMapString(stage.MultipleRests, &stage.MultipleRests_mapString)

	__gong__rebuildMapString(stage.PartClefs, &stage.PartClefs_mapString)

	__gong__rebuildMapString(stage.PartSymbols, &stage.PartSymbols_mapString)

	__gong__rebuildMapString(stage.PartTransposes, &stage.PartTransposes_mapString)

	__gong__rebuildMapString(stage.Slashs, &stage.Slashs_mapString)

	__gong__rebuildMapString(stage.StaffDetailss, &stage.StaffDetailss_mapString)

	__gong__rebuildMapString(stage.StaffSizes, &stage.StaffSizes_mapString)

	__gong__rebuildMapString(stage.StaffTunings, &stage.StaffTunings_mapString)

	__gong__rebuildMapString(stage.Times, &stage.Times_mapString)

	__gong__rebuildMapString(stage.Transposes, &stage.Transposes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
