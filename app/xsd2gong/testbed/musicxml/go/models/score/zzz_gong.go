// generated code - do not edit
package score

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
	Credits                map[*Credit]struct{}
	Credits_instance       map[*Credit]*Credit
	Credits_mapString      map[string]*Credit
	CreditOrder            uint
	Credit_stagedOrder     map[*Credit]uint
	Credit_orderStaged     map[uint]*Credit
	Credits_reference      map[*Credit]*Credit
	Credits_referenceOrder map[*Credit]uint

	// insertion point for slice of pointers maps
	OnAfterCreditCreateCallback GongOnAfterCreateInterface[Credit]
	OnAfterCreditUpdateCallback GongOnAfterUpdateInterface[Credit]
	OnAfterCreditDeleteCallback GongOnAfterDeleteInterface[Credit]

	Defaultss                map[*Defaults]struct{}
	Defaultss_instance       map[*Defaults]*Defaults
	Defaultss_mapString      map[string]*Defaults
	DefaultsOrder            uint
	Defaults_stagedOrder     map[*Defaults]uint
	Defaults_orderStaged     map[uint]*Defaults
	Defaultss_reference      map[*Defaults]*Defaults
	Defaultss_referenceOrder map[*Defaults]uint

	// insertion point for slice of pointers maps
	Defaults_LyricFont_reverseMap map[*LyricFont]*Defaults

	Defaults_LyricLanguage_reverseMap map[*LyricLanguage]*Defaults

	OnAfterDefaultsCreateCallback GongOnAfterCreateInterface[Defaults]
	OnAfterDefaultsUpdateCallback GongOnAfterUpdateInterface[Defaults]
	OnAfterDefaultsDeleteCallback GongOnAfterDeleteInterface[Defaults]

	EmptyFonts                map[*EmptyFont]struct{}
	EmptyFonts_instance       map[*EmptyFont]*EmptyFont
	EmptyFonts_mapString      map[string]*EmptyFont
	EmptyFontOrder            uint
	EmptyFont_stagedOrder     map[*EmptyFont]uint
	EmptyFont_orderStaged     map[uint]*EmptyFont
	EmptyFonts_reference      map[*EmptyFont]*EmptyFont
	EmptyFonts_referenceOrder map[*EmptyFont]uint

	// insertion point for slice of pointers maps
	OnAfterEmptyFontCreateCallback GongOnAfterCreateInterface[EmptyFont]
	OnAfterEmptyFontUpdateCallback GongOnAfterUpdateInterface[EmptyFont]
	OnAfterEmptyFontDeleteCallback GongOnAfterDeleteInterface[EmptyFont]

	GroupBarlines                map[*GroupBarline]struct{}
	GroupBarlines_instance       map[*GroupBarline]*GroupBarline
	GroupBarlines_mapString      map[string]*GroupBarline
	GroupBarlineOrder            uint
	GroupBarline_stagedOrder     map[*GroupBarline]uint
	GroupBarline_orderStaged     map[uint]*GroupBarline
	GroupBarlines_reference      map[*GroupBarline]*GroupBarline
	GroupBarlines_referenceOrder map[*GroupBarline]uint

	// insertion point for slice of pointers maps
	OnAfterGroupBarlineCreateCallback GongOnAfterCreateInterface[GroupBarline]
	OnAfterGroupBarlineUpdateCallback GongOnAfterUpdateInterface[GroupBarline]
	OnAfterGroupBarlineDeleteCallback GongOnAfterDeleteInterface[GroupBarline]

	GroupNames                map[*GroupName]struct{}
	GroupNames_instance       map[*GroupName]*GroupName
	GroupNames_mapString      map[string]*GroupName
	GroupNameOrder            uint
	GroupName_stagedOrder     map[*GroupName]uint
	GroupName_orderStaged     map[uint]*GroupName
	GroupNames_reference      map[*GroupName]*GroupName
	GroupNames_referenceOrder map[*GroupName]uint

	// insertion point for slice of pointers maps
	OnAfterGroupNameCreateCallback GongOnAfterCreateInterface[GroupName]
	OnAfterGroupNameUpdateCallback GongOnAfterUpdateInterface[GroupName]
	OnAfterGroupNameDeleteCallback GongOnAfterDeleteInterface[GroupName]

	GroupSymbols                map[*GroupSymbol]struct{}
	GroupSymbols_instance       map[*GroupSymbol]*GroupSymbol
	GroupSymbols_mapString      map[string]*GroupSymbol
	GroupSymbolOrder            uint
	GroupSymbol_stagedOrder     map[*GroupSymbol]uint
	GroupSymbol_orderStaged     map[uint]*GroupSymbol
	GroupSymbols_reference      map[*GroupSymbol]*GroupSymbol
	GroupSymbols_referenceOrder map[*GroupSymbol]uint

	// insertion point for slice of pointers maps
	OnAfterGroupSymbolCreateCallback GongOnAfterCreateInterface[GroupSymbol]
	OnAfterGroupSymbolUpdateCallback GongOnAfterUpdateInterface[GroupSymbol]
	OnAfterGroupSymbolDeleteCallback GongOnAfterDeleteInterface[GroupSymbol]

	InstrumentLinks                map[*InstrumentLink]struct{}
	InstrumentLinks_instance       map[*InstrumentLink]*InstrumentLink
	InstrumentLinks_mapString      map[string]*InstrumentLink
	InstrumentLinkOrder            uint
	InstrumentLink_stagedOrder     map[*InstrumentLink]uint
	InstrumentLink_orderStaged     map[uint]*InstrumentLink
	InstrumentLinks_reference      map[*InstrumentLink]*InstrumentLink
	InstrumentLinks_referenceOrder map[*InstrumentLink]uint

	// insertion point for slice of pointers maps
	OnAfterInstrumentLinkCreateCallback GongOnAfterCreateInterface[InstrumentLink]
	OnAfterInstrumentLinkUpdateCallback GongOnAfterUpdateInterface[InstrumentLink]
	OnAfterInstrumentLinkDeleteCallback GongOnAfterDeleteInterface[InstrumentLink]

	LyricFonts                map[*LyricFont]struct{}
	LyricFonts_instance       map[*LyricFont]*LyricFont
	LyricFonts_mapString      map[string]*LyricFont
	LyricFontOrder            uint
	LyricFont_stagedOrder     map[*LyricFont]uint
	LyricFont_orderStaged     map[uint]*LyricFont
	LyricFonts_reference      map[*LyricFont]*LyricFont
	LyricFonts_referenceOrder map[*LyricFont]uint

	// insertion point for slice of pointers maps
	OnAfterLyricFontCreateCallback GongOnAfterCreateInterface[LyricFont]
	OnAfterLyricFontUpdateCallback GongOnAfterUpdateInterface[LyricFont]
	OnAfterLyricFontDeleteCallback GongOnAfterDeleteInterface[LyricFont]

	LyricLanguages                map[*LyricLanguage]struct{}
	LyricLanguages_instance       map[*LyricLanguage]*LyricLanguage
	LyricLanguages_mapString      map[string]*LyricLanguage
	LyricLanguageOrder            uint
	LyricLanguage_stagedOrder     map[*LyricLanguage]uint
	LyricLanguage_orderStaged     map[uint]*LyricLanguage
	LyricLanguages_reference      map[*LyricLanguage]*LyricLanguage
	LyricLanguages_referenceOrder map[*LyricLanguage]uint

	// insertion point for slice of pointers maps
	OnAfterLyricLanguageCreateCallback GongOnAfterCreateInterface[LyricLanguage]
	OnAfterLyricLanguageUpdateCallback GongOnAfterUpdateInterface[LyricLanguage]
	OnAfterLyricLanguageDeleteCallback GongOnAfterDeleteInterface[LyricLanguage]

	Measures                map[*Measure]struct{}
	Measures_instance       map[*Measure]*Measure
	Measures_mapString      map[string]*Measure
	MeasureOrder            uint
	Measure_stagedOrder     map[*Measure]uint
	Measure_orderStaged     map[uint]*Measure
	Measures_reference      map[*Measure]*Measure
	Measures_referenceOrder map[*Measure]uint

	// insertion point for slice of pointers maps
	Measure_MusicData_reverseMap map[*MeasureMusicDataChoice]*Measure

	OnAfterMeasureCreateCallback GongOnAfterCreateInterface[Measure]
	OnAfterMeasureUpdateCallback GongOnAfterUpdateInterface[Measure]
	OnAfterMeasureDeleteCallback GongOnAfterDeleteInterface[Measure]

	MeasureMusicDataChoices                map[*MeasureMusicDataChoice]struct{}
	MeasureMusicDataChoices_instance       map[*MeasureMusicDataChoice]*MeasureMusicDataChoice
	MeasureMusicDataChoices_mapString      map[string]*MeasureMusicDataChoice
	MeasureMusicDataChoiceOrder            uint
	MeasureMusicDataChoice_stagedOrder     map[*MeasureMusicDataChoice]uint
	MeasureMusicDataChoice_orderStaged     map[uint]*MeasureMusicDataChoice
	MeasureMusicDataChoices_reference      map[*MeasureMusicDataChoice]*MeasureMusicDataChoice
	MeasureMusicDataChoices_referenceOrder map[*MeasureMusicDataChoice]uint

	// insertion point for slice of pointers maps
	OnAfterMeasureMusicDataChoiceCreateCallback GongOnAfterCreateInterface[MeasureMusicDataChoice]
	OnAfterMeasureMusicDataChoiceUpdateCallback GongOnAfterUpdateInterface[MeasureMusicDataChoice]
	OnAfterMeasureMusicDataChoiceDeleteCallback GongOnAfterDeleteInterface[MeasureMusicDataChoice]

	Opuss                map[*Opus]struct{}
	Opuss_instance       map[*Opus]*Opus
	Opuss_mapString      map[string]*Opus
	OpusOrder            uint
	Opus_stagedOrder     map[*Opus]uint
	Opus_orderStaged     map[uint]*Opus
	Opuss_reference      map[*Opus]*Opus
	Opuss_referenceOrder map[*Opus]uint

	// insertion point for slice of pointers maps
	OnAfterOpusCreateCallback GongOnAfterCreateInterface[Opus]
	OnAfterOpusUpdateCallback GongOnAfterUpdateInterface[Opus]
	OnAfterOpusDeleteCallback GongOnAfterDeleteInterface[Opus]

	Parts                map[*Part]struct{}
	Parts_instance       map[*Part]*Part
	Parts_mapString      map[string]*Part
	PartOrder            uint
	Part_stagedOrder     map[*Part]uint
	Part_orderStaged     map[uint]*Part
	Parts_reference      map[*Part]*Part
	Parts_referenceOrder map[*Part]uint

	// insertion point for slice of pointers maps
	Part_Measure_reverseMap map[*Measure]*Part

	OnAfterPartCreateCallback GongOnAfterCreateInterface[Part]
	OnAfterPartUpdateCallback GongOnAfterUpdateInterface[Part]
	OnAfterPartDeleteCallback GongOnAfterDeleteInterface[Part]

	PartGroups                map[*PartGroup]struct{}
	PartGroups_instance       map[*PartGroup]*PartGroup
	PartGroups_mapString      map[string]*PartGroup
	PartGroupOrder            uint
	PartGroup_stagedOrder     map[*PartGroup]uint
	PartGroup_orderStaged     map[uint]*PartGroup
	PartGroups_reference      map[*PartGroup]*PartGroup
	PartGroups_referenceOrder map[*PartGroup]uint

	// insertion point for slice of pointers maps
	OnAfterPartGroupCreateCallback GongOnAfterCreateInterface[PartGroup]
	OnAfterPartGroupUpdateCallback GongOnAfterUpdateInterface[PartGroup]
	OnAfterPartGroupDeleteCallback GongOnAfterDeleteInterface[PartGroup]

	PartLinks                map[*PartLink]struct{}
	PartLinks_instance       map[*PartLink]*PartLink
	PartLinks_mapString      map[string]*PartLink
	PartLinkOrder            uint
	PartLink_stagedOrder     map[*PartLink]uint
	PartLink_orderStaged     map[uint]*PartLink
	PartLinks_reference      map[*PartLink]*PartLink
	PartLinks_referenceOrder map[*PartLink]uint

	// insertion point for slice of pointers maps
	PartLink_InstrumentLink_reverseMap map[*InstrumentLink]*PartLink

	OnAfterPartLinkCreateCallback GongOnAfterCreateInterface[PartLink]
	OnAfterPartLinkUpdateCallback GongOnAfterUpdateInterface[PartLink]
	OnAfterPartLinkDeleteCallback GongOnAfterDeleteInterface[PartLink]

	PartLists                map[*PartList]struct{}
	PartLists_instance       map[*PartList]*PartList
	PartLists_mapString      map[string]*PartList
	PartListOrder            uint
	PartList_stagedOrder     map[*PartList]uint
	PartList_orderStaged     map[uint]*PartList
	PartLists_reference      map[*PartList]*PartList
	PartLists_referenceOrder map[*PartList]uint

	// insertion point for slice of pointers maps
	PartList_PartGroup_reverseMap map[*PartGroup]*PartList

	PartList_ScorePart_reverseMap map[*ScorePart]*PartList

	OnAfterPartListCreateCallback GongOnAfterCreateInterface[PartList]
	OnAfterPartListUpdateCallback GongOnAfterUpdateInterface[PartList]
	OnAfterPartListDeleteCallback GongOnAfterDeleteInterface[PartList]

	PartNames                map[*PartName]struct{}
	PartNames_instance       map[*PartName]*PartName
	PartNames_mapString      map[string]*PartName
	PartNameOrder            uint
	PartName_stagedOrder     map[*PartName]uint
	PartName_orderStaged     map[uint]*PartName
	PartNames_reference      map[*PartName]*PartName
	PartNames_referenceOrder map[*PartName]uint

	// insertion point for slice of pointers maps
	OnAfterPartNameCreateCallback GongOnAfterCreateInterface[PartName]
	OnAfterPartNameUpdateCallback GongOnAfterUpdateInterface[PartName]
	OnAfterPartNameDeleteCallback GongOnAfterDeleteInterface[PartName]

	Players                map[*Player]struct{}
	Players_instance       map[*Player]*Player
	Players_mapString      map[string]*Player
	PlayerOrder            uint
	Player_stagedOrder     map[*Player]uint
	Player_orderStaged     map[uint]*Player
	Players_reference      map[*Player]*Player
	Players_referenceOrder map[*Player]uint

	// insertion point for slice of pointers maps
	OnAfterPlayerCreateCallback GongOnAfterCreateInterface[Player]
	OnAfterPlayerUpdateCallback GongOnAfterUpdateInterface[Player]
	OnAfterPlayerDeleteCallback GongOnAfterDeleteInterface[Player]

	ScoreInstruments                map[*ScoreInstrument]struct{}
	ScoreInstruments_instance       map[*ScoreInstrument]*ScoreInstrument
	ScoreInstruments_mapString      map[string]*ScoreInstrument
	ScoreInstrumentOrder            uint
	ScoreInstrument_stagedOrder     map[*ScoreInstrument]uint
	ScoreInstrument_orderStaged     map[uint]*ScoreInstrument
	ScoreInstruments_reference      map[*ScoreInstrument]*ScoreInstrument
	ScoreInstruments_referenceOrder map[*ScoreInstrument]uint

	// insertion point for slice of pointers maps
	OnAfterScoreInstrumentCreateCallback GongOnAfterCreateInterface[ScoreInstrument]
	OnAfterScoreInstrumentUpdateCallback GongOnAfterUpdateInterface[ScoreInstrument]
	OnAfterScoreInstrumentDeleteCallback GongOnAfterDeleteInterface[ScoreInstrument]

	ScoreParts                map[*ScorePart]struct{}
	ScoreParts_instance       map[*ScorePart]*ScorePart
	ScoreParts_mapString      map[string]*ScorePart
	ScorePartOrder            uint
	ScorePart_stagedOrder     map[*ScorePart]uint
	ScorePart_orderStaged     map[uint]*ScorePart
	ScoreParts_reference      map[*ScorePart]*ScorePart
	ScoreParts_referenceOrder map[*ScorePart]uint

	// insertion point for slice of pointers maps
	ScorePart_PartLink_reverseMap map[*PartLink]*ScorePart

	ScorePart_ScoreInstrument_reverseMap map[*ScoreInstrument]*ScorePart

	ScorePart_Player_reverseMap map[*Player]*ScorePart

	OnAfterScorePartCreateCallback GongOnAfterCreateInterface[ScorePart]
	OnAfterScorePartUpdateCallback GongOnAfterUpdateInterface[ScorePart]
	OnAfterScorePartDeleteCallback GongOnAfterDeleteInterface[ScorePart]

	ScorePartwises                map[*ScorePartwise]struct{}
	ScorePartwises_instance       map[*ScorePartwise]*ScorePartwise
	ScorePartwises_mapString      map[string]*ScorePartwise
	ScorePartwiseOrder            uint
	ScorePartwise_stagedOrder     map[*ScorePartwise]uint
	ScorePartwise_orderStaged     map[uint]*ScorePartwise
	ScorePartwises_reference      map[*ScorePartwise]*ScorePartwise
	ScorePartwises_referenceOrder map[*ScorePartwise]uint

	// insertion point for slice of pointers maps
	ScorePartwise_Credit_reverseMap map[*Credit]*ScorePartwise

	ScorePartwise_Part_reverseMap map[*Part]*ScorePartwise

	OnAfterScorePartwiseCreateCallback GongOnAfterCreateInterface[ScorePartwise]
	OnAfterScorePartwiseUpdateCallback GongOnAfterUpdateInterface[ScorePartwise]
	OnAfterScorePartwiseDeleteCallback GongOnAfterDeleteInterface[ScorePartwise]

	ScoreTimewises                map[*ScoreTimewise]struct{}
	ScoreTimewises_instance       map[*ScoreTimewise]*ScoreTimewise
	ScoreTimewises_mapString      map[string]*ScoreTimewise
	ScoreTimewiseOrder            uint
	ScoreTimewise_stagedOrder     map[*ScoreTimewise]uint
	ScoreTimewise_orderStaged     map[uint]*ScoreTimewise
	ScoreTimewises_reference      map[*ScoreTimewise]*ScoreTimewise
	ScoreTimewises_referenceOrder map[*ScoreTimewise]uint

	// insertion point for slice of pointers maps
	ScoreTimewise_Credit_reverseMap map[*Credit]*ScoreTimewise

	ScoreTimewise_Measure_reverseMap map[*ScoreTimewiseMeasure]*ScoreTimewise

	OnAfterScoreTimewiseCreateCallback GongOnAfterCreateInterface[ScoreTimewise]
	OnAfterScoreTimewiseUpdateCallback GongOnAfterUpdateInterface[ScoreTimewise]
	OnAfterScoreTimewiseDeleteCallback GongOnAfterDeleteInterface[ScoreTimewise]

	ScoreTimewiseMeasures                map[*ScoreTimewiseMeasure]struct{}
	ScoreTimewiseMeasures_instance       map[*ScoreTimewiseMeasure]*ScoreTimewiseMeasure
	ScoreTimewiseMeasures_mapString      map[string]*ScoreTimewiseMeasure
	ScoreTimewiseMeasureOrder            uint
	ScoreTimewiseMeasure_stagedOrder     map[*ScoreTimewiseMeasure]uint
	ScoreTimewiseMeasure_orderStaged     map[uint]*ScoreTimewiseMeasure
	ScoreTimewiseMeasures_reference      map[*ScoreTimewiseMeasure]*ScoreTimewiseMeasure
	ScoreTimewiseMeasures_referenceOrder map[*ScoreTimewiseMeasure]uint

	// insertion point for slice of pointers maps
	ScoreTimewiseMeasure_Part_reverseMap map[*ScoreTimewiseMeasurePart]*ScoreTimewiseMeasure

	OnAfterScoreTimewiseMeasureCreateCallback GongOnAfterCreateInterface[ScoreTimewiseMeasure]
	OnAfterScoreTimewiseMeasureUpdateCallback GongOnAfterUpdateInterface[ScoreTimewiseMeasure]
	OnAfterScoreTimewiseMeasureDeleteCallback GongOnAfterDeleteInterface[ScoreTimewiseMeasure]

	ScoreTimewiseMeasureParts                map[*ScoreTimewiseMeasurePart]struct{}
	ScoreTimewiseMeasureParts_instance       map[*ScoreTimewiseMeasurePart]*ScoreTimewiseMeasurePart
	ScoreTimewiseMeasureParts_mapString      map[string]*ScoreTimewiseMeasurePart
	ScoreTimewiseMeasurePartOrder            uint
	ScoreTimewiseMeasurePart_stagedOrder     map[*ScoreTimewiseMeasurePart]uint
	ScoreTimewiseMeasurePart_orderStaged     map[uint]*ScoreTimewiseMeasurePart
	ScoreTimewiseMeasureParts_reference      map[*ScoreTimewiseMeasurePart]*ScoreTimewiseMeasurePart
	ScoreTimewiseMeasureParts_referenceOrder map[*ScoreTimewiseMeasurePart]uint

	// insertion point for slice of pointers maps
	OnAfterScoreTimewiseMeasurePartCreateCallback GongOnAfterCreateInterface[ScoreTimewiseMeasurePart]
	OnAfterScoreTimewiseMeasurePartUpdateCallback GongOnAfterUpdateInterface[ScoreTimewiseMeasurePart]
	OnAfterScoreTimewiseMeasurePartDeleteCallback GongOnAfterDeleteInterface[ScoreTimewiseMeasurePart]

	Works                map[*Work]struct{}
	Works_instance       map[*Work]*Work
	Works_mapString      map[string]*Work
	WorkOrder            uint
	Work_stagedOrder     map[*Work]uint
	Work_orderStaged     map[uint]*Work
	Works_reference      map[*Work]*Work
	Works_referenceOrder map[*Work]uint

	// insertion point for slice of pointers maps
	OnAfterWorkCreateCallback GongOnAfterCreateInterface[Work]
	OnAfterWorkUpdateCallback GongOnAfterUpdateInterface[Work]
	OnAfterWorkDeleteCallback GongOnAfterDeleteInterface[Work]

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
	__gong__clearReferences(&stage.Credits_reference, &stage.Credits_instance, &stage.Credits_referenceOrder)

	__gong__clearReferences(&stage.Defaultss_reference, &stage.Defaultss_instance, &stage.Defaultss_referenceOrder)

	__gong__clearReferences(&stage.EmptyFonts_reference, &stage.EmptyFonts_instance, &stage.EmptyFonts_referenceOrder)

	__gong__clearReferences(&stage.GroupBarlines_reference, &stage.GroupBarlines_instance, &stage.GroupBarlines_referenceOrder)

	__gong__clearReferences(&stage.GroupNames_reference, &stage.GroupNames_instance, &stage.GroupNames_referenceOrder)

	__gong__clearReferences(&stage.GroupSymbols_reference, &stage.GroupSymbols_instance, &stage.GroupSymbols_referenceOrder)

	__gong__clearReferences(&stage.InstrumentLinks_reference, &stage.InstrumentLinks_instance, &stage.InstrumentLinks_referenceOrder)

	__gong__clearReferences(&stage.LyricFonts_reference, &stage.LyricFonts_instance, &stage.LyricFonts_referenceOrder)

	__gong__clearReferences(&stage.LyricLanguages_reference, &stage.LyricLanguages_instance, &stage.LyricLanguages_referenceOrder)

	__gong__clearReferences(&stage.Measures_reference, &stage.Measures_instance, &stage.Measures_referenceOrder)

	__gong__clearReferences(&stage.MeasureMusicDataChoices_reference, &stage.MeasureMusicDataChoices_instance, &stage.MeasureMusicDataChoices_referenceOrder)

	__gong__clearReferences(&stage.Opuss_reference, &stage.Opuss_instance, &stage.Opuss_referenceOrder)

	__gong__clearReferences(&stage.Parts_reference, &stage.Parts_instance, &stage.Parts_referenceOrder)

	__gong__clearReferences(&stage.PartGroups_reference, &stage.PartGroups_instance, &stage.PartGroups_referenceOrder)

	__gong__clearReferences(&stage.PartLinks_reference, &stage.PartLinks_instance, &stage.PartLinks_referenceOrder)

	__gong__clearReferences(&stage.PartLists_reference, &stage.PartLists_instance, &stage.PartLists_referenceOrder)

	__gong__clearReferences(&stage.PartNames_reference, &stage.PartNames_instance, &stage.PartNames_referenceOrder)

	__gong__clearReferences(&stage.Players_reference, &stage.Players_instance, &stage.Players_referenceOrder)

	__gong__clearReferences(&stage.ScoreInstruments_reference, &stage.ScoreInstruments_instance, &stage.ScoreInstruments_referenceOrder)

	__gong__clearReferences(&stage.ScoreParts_reference, &stage.ScoreParts_instance, &stage.ScoreParts_referenceOrder)

	__gong__clearReferences(&stage.ScorePartwises_reference, &stage.ScorePartwises_instance, &stage.ScorePartwises_referenceOrder)

	__gong__clearReferences(&stage.ScoreTimewises_reference, &stage.ScoreTimewises_instance, &stage.ScoreTimewises_referenceOrder)

	__gong__clearReferences(&stage.ScoreTimewiseMeasures_reference, &stage.ScoreTimewiseMeasures_instance, &stage.ScoreTimewiseMeasures_referenceOrder)

	__gong__clearReferences(&stage.ScoreTimewiseMeasureParts_reference, &stage.ScoreTimewiseMeasureParts_instance, &stage.ScoreTimewiseMeasureParts_referenceOrder)

	__gong__clearReferences(&stage.Works_reference, &stage.Works_instance, &stage.Works_referenceOrder)

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
	stage.CreditOrder = __gong__recomputeOrder(stage.Credit_stagedOrder)

	stage.DefaultsOrder = __gong__recomputeOrder(stage.Defaults_stagedOrder)

	stage.EmptyFontOrder = __gong__recomputeOrder(stage.EmptyFont_stagedOrder)

	stage.GroupBarlineOrder = __gong__recomputeOrder(stage.GroupBarline_stagedOrder)

	stage.GroupNameOrder = __gong__recomputeOrder(stage.GroupName_stagedOrder)

	stage.GroupSymbolOrder = __gong__recomputeOrder(stage.GroupSymbol_stagedOrder)

	stage.InstrumentLinkOrder = __gong__recomputeOrder(stage.InstrumentLink_stagedOrder)

	stage.LyricFontOrder = __gong__recomputeOrder(stage.LyricFont_stagedOrder)

	stage.LyricLanguageOrder = __gong__recomputeOrder(stage.LyricLanguage_stagedOrder)

	stage.MeasureOrder = __gong__recomputeOrder(stage.Measure_stagedOrder)

	stage.MeasureMusicDataChoiceOrder = __gong__recomputeOrder(stage.MeasureMusicDataChoice_stagedOrder)

	stage.OpusOrder = __gong__recomputeOrder(stage.Opus_stagedOrder)

	stage.PartOrder = __gong__recomputeOrder(stage.Part_stagedOrder)

	stage.PartGroupOrder = __gong__recomputeOrder(stage.PartGroup_stagedOrder)

	stage.PartLinkOrder = __gong__recomputeOrder(stage.PartLink_stagedOrder)

	stage.PartListOrder = __gong__recomputeOrder(stage.PartList_stagedOrder)

	stage.PartNameOrder = __gong__recomputeOrder(stage.PartName_stagedOrder)

	stage.PlayerOrder = __gong__recomputeOrder(stage.Player_stagedOrder)

	stage.ScoreInstrumentOrder = __gong__recomputeOrder(stage.ScoreInstrument_stagedOrder)

	stage.ScorePartOrder = __gong__recomputeOrder(stage.ScorePart_stagedOrder)

	stage.ScorePartwiseOrder = __gong__recomputeOrder(stage.ScorePartwise_stagedOrder)

	stage.ScoreTimewiseOrder = __gong__recomputeOrder(stage.ScoreTimewise_stagedOrder)

	stage.ScoreTimewiseMeasureOrder = __gong__recomputeOrder(stage.ScoreTimewiseMeasure_stagedOrder)

	stage.ScoreTimewiseMeasurePartOrder = __gong__recomputeOrder(stage.ScoreTimewiseMeasurePart_stagedOrder)

	stage.WorkOrder = __gong__recomputeOrder(stage.Work_stagedOrder)

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
func (*Credit) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Credits, stage.Credit_stagedOrder)
}

func (*Credit) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Credit_orderStaged[order]
}

func (*Credit) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Credits_mapString
}

func (*Credit) GongGetInstancesSet(stage *Stage) any {
	return &stage.Credits
}

func (*Credit) GongNewInstance() any {
	return new(Credit)
}

func (*Defaults) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Defaultss, stage.Defaults_stagedOrder)
}

func (*Defaults) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Defaults_orderStaged[order]
}

func (*Defaults) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Defaultss_mapString
}

func (*Defaults) GongGetInstancesSet(stage *Stage) any {
	return &stage.Defaultss
}

func (*Defaults) GongNewInstance() any {
	return new(Defaults)
}

func (*EmptyFont) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.EmptyFonts, stage.EmptyFont_stagedOrder)
}

func (*EmptyFont) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.EmptyFont_orderStaged[order]
}

func (*EmptyFont) GongGetInstancesMapByName(stage *Stage) any {
	return stage.EmptyFonts_mapString
}

func (*EmptyFont) GongGetInstancesSet(stage *Stage) any {
	return &stage.EmptyFonts
}

func (*EmptyFont) GongNewInstance() any {
	return new(EmptyFont)
}

func (*GroupBarline) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.GroupBarlines, stage.GroupBarline_stagedOrder)
}

func (*GroupBarline) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.GroupBarline_orderStaged[order]
}

func (*GroupBarline) GongGetInstancesMapByName(stage *Stage) any {
	return stage.GroupBarlines_mapString
}

func (*GroupBarline) GongGetInstancesSet(stage *Stage) any {
	return &stage.GroupBarlines
}

func (*GroupBarline) GongNewInstance() any {
	return new(GroupBarline)
}

func (*GroupName) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.GroupNames, stage.GroupName_stagedOrder)
}

func (*GroupName) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.GroupName_orderStaged[order]
}

func (*GroupName) GongGetInstancesMapByName(stage *Stage) any {
	return stage.GroupNames_mapString
}

func (*GroupName) GongGetInstancesSet(stage *Stage) any {
	return &stage.GroupNames
}

func (*GroupName) GongNewInstance() any {
	return new(GroupName)
}

func (*GroupSymbol) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.GroupSymbols, stage.GroupSymbol_stagedOrder)
}

func (*GroupSymbol) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.GroupSymbol_orderStaged[order]
}

func (*GroupSymbol) GongGetInstancesMapByName(stage *Stage) any {
	return stage.GroupSymbols_mapString
}

func (*GroupSymbol) GongGetInstancesSet(stage *Stage) any {
	return &stage.GroupSymbols
}

func (*GroupSymbol) GongNewInstance() any {
	return new(GroupSymbol)
}

func (*InstrumentLink) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.InstrumentLinks, stage.InstrumentLink_stagedOrder)
}

func (*InstrumentLink) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.InstrumentLink_orderStaged[order]
}

func (*InstrumentLink) GongGetInstancesMapByName(stage *Stage) any {
	return stage.InstrumentLinks_mapString
}

func (*InstrumentLink) GongGetInstancesSet(stage *Stage) any {
	return &stage.InstrumentLinks
}

func (*InstrumentLink) GongNewInstance() any {
	return new(InstrumentLink)
}

func (*LyricFont) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.LyricFonts, stage.LyricFont_stagedOrder)
}

func (*LyricFont) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.LyricFont_orderStaged[order]
}

func (*LyricFont) GongGetInstancesMapByName(stage *Stage) any {
	return stage.LyricFonts_mapString
}

func (*LyricFont) GongGetInstancesSet(stage *Stage) any {
	return &stage.LyricFonts
}

func (*LyricFont) GongNewInstance() any {
	return new(LyricFont)
}

func (*LyricLanguage) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.LyricLanguages, stage.LyricLanguage_stagedOrder)
}

func (*LyricLanguage) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.LyricLanguage_orderStaged[order]
}

func (*LyricLanguage) GongGetInstancesMapByName(stage *Stage) any {
	return stage.LyricLanguages_mapString
}

func (*LyricLanguage) GongGetInstancesSet(stage *Stage) any {
	return &stage.LyricLanguages
}

func (*LyricLanguage) GongNewInstance() any {
	return new(LyricLanguage)
}

func (*Measure) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Measures, stage.Measure_stagedOrder)
}

func (*Measure) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Measure_orderStaged[order]
}

func (*Measure) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Measures_mapString
}

func (*Measure) GongGetInstancesSet(stage *Stage) any {
	return &stage.Measures
}

func (*Measure) GongNewInstance() any {
	return new(Measure)
}

func (*MeasureMusicDataChoice) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MeasureMusicDataChoices, stage.MeasureMusicDataChoice_stagedOrder)
}

func (*MeasureMusicDataChoice) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MeasureMusicDataChoice_orderStaged[order]
}

func (*MeasureMusicDataChoice) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MeasureMusicDataChoices_mapString
}

func (*MeasureMusicDataChoice) GongGetInstancesSet(stage *Stage) any {
	return &stage.MeasureMusicDataChoices
}

func (*MeasureMusicDataChoice) GongNewInstance() any {
	return new(MeasureMusicDataChoice)
}

func (*Opus) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Opuss, stage.Opus_stagedOrder)
}

func (*Opus) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Opus_orderStaged[order]
}

func (*Opus) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Opuss_mapString
}

func (*Opus) GongGetInstancesSet(stage *Stage) any {
	return &stage.Opuss
}

func (*Opus) GongNewInstance() any {
	return new(Opus)
}

func (*Part) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Parts, stage.Part_stagedOrder)
}

func (*Part) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Part_orderStaged[order]
}

func (*Part) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Parts_mapString
}

func (*Part) GongGetInstancesSet(stage *Stage) any {
	return &stage.Parts
}

func (*Part) GongNewInstance() any {
	return new(Part)
}

func (*PartGroup) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.PartGroups, stage.PartGroup_stagedOrder)
}

func (*PartGroup) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.PartGroup_orderStaged[order]
}

func (*PartGroup) GongGetInstancesMapByName(stage *Stage) any {
	return stage.PartGroups_mapString
}

func (*PartGroup) GongGetInstancesSet(stage *Stage) any {
	return &stage.PartGroups
}

func (*PartGroup) GongNewInstance() any {
	return new(PartGroup)
}

func (*PartLink) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.PartLinks, stage.PartLink_stagedOrder)
}

func (*PartLink) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.PartLink_orderStaged[order]
}

func (*PartLink) GongGetInstancesMapByName(stage *Stage) any {
	return stage.PartLinks_mapString
}

func (*PartLink) GongGetInstancesSet(stage *Stage) any {
	return &stage.PartLinks
}

func (*PartLink) GongNewInstance() any {
	return new(PartLink)
}

func (*PartList) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.PartLists, stage.PartList_stagedOrder)
}

func (*PartList) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.PartList_orderStaged[order]
}

func (*PartList) GongGetInstancesMapByName(stage *Stage) any {
	return stage.PartLists_mapString
}

func (*PartList) GongGetInstancesSet(stage *Stage) any {
	return &stage.PartLists
}

func (*PartList) GongNewInstance() any {
	return new(PartList)
}

func (*PartName) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.PartNames, stage.PartName_stagedOrder)
}

func (*PartName) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.PartName_orderStaged[order]
}

func (*PartName) GongGetInstancesMapByName(stage *Stage) any {
	return stage.PartNames_mapString
}

func (*PartName) GongGetInstancesSet(stage *Stage) any {
	return &stage.PartNames
}

func (*PartName) GongNewInstance() any {
	return new(PartName)
}

func (*Player) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Players, stage.Player_stagedOrder)
}

func (*Player) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Player_orderStaged[order]
}

func (*Player) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Players_mapString
}

func (*Player) GongGetInstancesSet(stage *Stage) any {
	return &stage.Players
}

func (*Player) GongNewInstance() any {
	return new(Player)
}

func (*ScoreInstrument) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ScoreInstruments, stage.ScoreInstrument_stagedOrder)
}

func (*ScoreInstrument) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ScoreInstrument_orderStaged[order]
}

func (*ScoreInstrument) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ScoreInstruments_mapString
}

func (*ScoreInstrument) GongGetInstancesSet(stage *Stage) any {
	return &stage.ScoreInstruments
}

func (*ScoreInstrument) GongNewInstance() any {
	return new(ScoreInstrument)
}

func (*ScorePart) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ScoreParts, stage.ScorePart_stagedOrder)
}

func (*ScorePart) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ScorePart_orderStaged[order]
}

func (*ScorePart) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ScoreParts_mapString
}

func (*ScorePart) GongGetInstancesSet(stage *Stage) any {
	return &stage.ScoreParts
}

func (*ScorePart) GongNewInstance() any {
	return new(ScorePart)
}

func (*ScorePartwise) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ScorePartwises, stage.ScorePartwise_stagedOrder)
}

func (*ScorePartwise) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ScorePartwise_orderStaged[order]
}

func (*ScorePartwise) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ScorePartwises_mapString
}

func (*ScorePartwise) GongGetInstancesSet(stage *Stage) any {
	return &stage.ScorePartwises
}

func (*ScorePartwise) GongNewInstance() any {
	return new(ScorePartwise)
}

func (*ScoreTimewise) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ScoreTimewises, stage.ScoreTimewise_stagedOrder)
}

func (*ScoreTimewise) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ScoreTimewise_orderStaged[order]
}

func (*ScoreTimewise) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ScoreTimewises_mapString
}

func (*ScoreTimewise) GongGetInstancesSet(stage *Stage) any {
	return &stage.ScoreTimewises
}

func (*ScoreTimewise) GongNewInstance() any {
	return new(ScoreTimewise)
}

func (*ScoreTimewiseMeasure) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ScoreTimewiseMeasures, stage.ScoreTimewiseMeasure_stagedOrder)
}

func (*ScoreTimewiseMeasure) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ScoreTimewiseMeasure_orderStaged[order]
}

func (*ScoreTimewiseMeasure) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ScoreTimewiseMeasures_mapString
}

func (*ScoreTimewiseMeasure) GongGetInstancesSet(stage *Stage) any {
	return &stage.ScoreTimewiseMeasures
}

func (*ScoreTimewiseMeasure) GongNewInstance() any {
	return new(ScoreTimewiseMeasure)
}

func (*ScoreTimewiseMeasurePart) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ScoreTimewiseMeasureParts, stage.ScoreTimewiseMeasurePart_stagedOrder)
}

func (*ScoreTimewiseMeasurePart) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ScoreTimewiseMeasurePart_orderStaged[order]
}

func (*ScoreTimewiseMeasurePart) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ScoreTimewiseMeasureParts_mapString
}

func (*ScoreTimewiseMeasurePart) GongGetInstancesSet(stage *Stage) any {
	return &stage.ScoreTimewiseMeasureParts
}

func (*ScoreTimewiseMeasurePart) GongNewInstance() any {
	return new(ScoreTimewiseMeasurePart)
}

func (*Work) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Works, stage.Work_stagedOrder)
}

func (*Work) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Work_orderStaged[order]
}

func (*Work) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Works_mapString
}

func (*Work) GongGetInstancesSet(stage *Stage) any {
	return &stage.Works
}

func (*Work) GongNewInstance() any {
	return new(Work)
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
	return "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/score"
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
		Credits:           make(map[*Credit]struct{}),
		Credits_mapString: make(map[string]*Credit),

		Defaultss:           make(map[*Defaults]struct{}),
		Defaultss_mapString: make(map[string]*Defaults),

		EmptyFonts:           make(map[*EmptyFont]struct{}),
		EmptyFonts_mapString: make(map[string]*EmptyFont),

		GroupBarlines:           make(map[*GroupBarline]struct{}),
		GroupBarlines_mapString: make(map[string]*GroupBarline),

		GroupNames:           make(map[*GroupName]struct{}),
		GroupNames_mapString: make(map[string]*GroupName),

		GroupSymbols:           make(map[*GroupSymbol]struct{}),
		GroupSymbols_mapString: make(map[string]*GroupSymbol),

		InstrumentLinks:           make(map[*InstrumentLink]struct{}),
		InstrumentLinks_mapString: make(map[string]*InstrumentLink),

		LyricFonts:           make(map[*LyricFont]struct{}),
		LyricFonts_mapString: make(map[string]*LyricFont),

		LyricLanguages:           make(map[*LyricLanguage]struct{}),
		LyricLanguages_mapString: make(map[string]*LyricLanguage),

		Measures:           make(map[*Measure]struct{}),
		Measures_mapString: make(map[string]*Measure),

		MeasureMusicDataChoices:           make(map[*MeasureMusicDataChoice]struct{}),
		MeasureMusicDataChoices_mapString: make(map[string]*MeasureMusicDataChoice),

		Opuss:           make(map[*Opus]struct{}),
		Opuss_mapString: make(map[string]*Opus),

		Parts:           make(map[*Part]struct{}),
		Parts_mapString: make(map[string]*Part),

		PartGroups:           make(map[*PartGroup]struct{}),
		PartGroups_mapString: make(map[string]*PartGroup),

		PartLinks:           make(map[*PartLink]struct{}),
		PartLinks_mapString: make(map[string]*PartLink),

		PartLists:           make(map[*PartList]struct{}),
		PartLists_mapString: make(map[string]*PartList),

		PartNames:           make(map[*PartName]struct{}),
		PartNames_mapString: make(map[string]*PartName),

		Players:           make(map[*Player]struct{}),
		Players_mapString: make(map[string]*Player),

		ScoreInstruments:           make(map[*ScoreInstrument]struct{}),
		ScoreInstruments_mapString: make(map[string]*ScoreInstrument),

		ScoreParts:           make(map[*ScorePart]struct{}),
		ScoreParts_mapString: make(map[string]*ScorePart),

		ScorePartwises:           make(map[*ScorePartwise]struct{}),
		ScorePartwises_mapString: make(map[string]*ScorePartwise),

		ScoreTimewises:           make(map[*ScoreTimewise]struct{}),
		ScoreTimewises_mapString: make(map[string]*ScoreTimewise),

		ScoreTimewiseMeasures:           make(map[*ScoreTimewiseMeasure]struct{}),
		ScoreTimewiseMeasures_mapString: make(map[string]*ScoreTimewiseMeasure),

		ScoreTimewiseMeasureParts:           make(map[*ScoreTimewiseMeasurePart]struct{}),
		ScoreTimewiseMeasureParts_mapString: make(map[string]*ScoreTimewiseMeasurePart),

		Works:           make(map[*Work]struct{}),
		Works_mapString: make(map[string]*Work),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		Credit_stagedOrder: make(map[*Credit]uint),
		Credit_orderStaged: make(map[uint]*Credit),
		Credits_reference:  make(map[*Credit]*Credit),

		Defaults_stagedOrder: make(map[*Defaults]uint),
		Defaults_orderStaged: make(map[uint]*Defaults),
		Defaultss_reference:  make(map[*Defaults]*Defaults),

		EmptyFont_stagedOrder: make(map[*EmptyFont]uint),
		EmptyFont_orderStaged: make(map[uint]*EmptyFont),
		EmptyFonts_reference:  make(map[*EmptyFont]*EmptyFont),

		GroupBarline_stagedOrder: make(map[*GroupBarline]uint),
		GroupBarline_orderStaged: make(map[uint]*GroupBarline),
		GroupBarlines_reference:  make(map[*GroupBarline]*GroupBarline),

		GroupName_stagedOrder: make(map[*GroupName]uint),
		GroupName_orderStaged: make(map[uint]*GroupName),
		GroupNames_reference:  make(map[*GroupName]*GroupName),

		GroupSymbol_stagedOrder: make(map[*GroupSymbol]uint),
		GroupSymbol_orderStaged: make(map[uint]*GroupSymbol),
		GroupSymbols_reference:  make(map[*GroupSymbol]*GroupSymbol),

		InstrumentLink_stagedOrder: make(map[*InstrumentLink]uint),
		InstrumentLink_orderStaged: make(map[uint]*InstrumentLink),
		InstrumentLinks_reference:  make(map[*InstrumentLink]*InstrumentLink),

		LyricFont_stagedOrder: make(map[*LyricFont]uint),
		LyricFont_orderStaged: make(map[uint]*LyricFont),
		LyricFonts_reference:  make(map[*LyricFont]*LyricFont),

		LyricLanguage_stagedOrder: make(map[*LyricLanguage]uint),
		LyricLanguage_orderStaged: make(map[uint]*LyricLanguage),
		LyricLanguages_reference:  make(map[*LyricLanguage]*LyricLanguage),

		Measure_stagedOrder: make(map[*Measure]uint),
		Measure_orderStaged: make(map[uint]*Measure),
		Measures_reference:  make(map[*Measure]*Measure),

		MeasureMusicDataChoice_stagedOrder: make(map[*MeasureMusicDataChoice]uint),
		MeasureMusicDataChoice_orderStaged: make(map[uint]*MeasureMusicDataChoice),
		MeasureMusicDataChoices_reference:  make(map[*MeasureMusicDataChoice]*MeasureMusicDataChoice),

		Opus_stagedOrder: make(map[*Opus]uint),
		Opus_orderStaged: make(map[uint]*Opus),
		Opuss_reference:  make(map[*Opus]*Opus),

		Part_stagedOrder: make(map[*Part]uint),
		Part_orderStaged: make(map[uint]*Part),
		Parts_reference:  make(map[*Part]*Part),

		PartGroup_stagedOrder: make(map[*PartGroup]uint),
		PartGroup_orderStaged: make(map[uint]*PartGroup),
		PartGroups_reference:  make(map[*PartGroup]*PartGroup),

		PartLink_stagedOrder: make(map[*PartLink]uint),
		PartLink_orderStaged: make(map[uint]*PartLink),
		PartLinks_reference:  make(map[*PartLink]*PartLink),

		PartList_stagedOrder: make(map[*PartList]uint),
		PartList_orderStaged: make(map[uint]*PartList),
		PartLists_reference:  make(map[*PartList]*PartList),

		PartName_stagedOrder: make(map[*PartName]uint),
		PartName_orderStaged: make(map[uint]*PartName),
		PartNames_reference:  make(map[*PartName]*PartName),

		Player_stagedOrder: make(map[*Player]uint),
		Player_orderStaged: make(map[uint]*Player),
		Players_reference:  make(map[*Player]*Player),

		ScoreInstrument_stagedOrder: make(map[*ScoreInstrument]uint),
		ScoreInstrument_orderStaged: make(map[uint]*ScoreInstrument),
		ScoreInstruments_reference:  make(map[*ScoreInstrument]*ScoreInstrument),

		ScorePart_stagedOrder: make(map[*ScorePart]uint),
		ScorePart_orderStaged: make(map[uint]*ScorePart),
		ScoreParts_reference:  make(map[*ScorePart]*ScorePart),

		ScorePartwise_stagedOrder: make(map[*ScorePartwise]uint),
		ScorePartwise_orderStaged: make(map[uint]*ScorePartwise),
		ScorePartwises_reference:  make(map[*ScorePartwise]*ScorePartwise),

		ScoreTimewise_stagedOrder: make(map[*ScoreTimewise]uint),
		ScoreTimewise_orderStaged: make(map[uint]*ScoreTimewise),
		ScoreTimewises_reference:  make(map[*ScoreTimewise]*ScoreTimewise),

		ScoreTimewiseMeasure_stagedOrder: make(map[*ScoreTimewiseMeasure]uint),
		ScoreTimewiseMeasure_orderStaged: make(map[uint]*ScoreTimewiseMeasure),
		ScoreTimewiseMeasures_reference:  make(map[*ScoreTimewiseMeasure]*ScoreTimewiseMeasure),

		ScoreTimewiseMeasurePart_stagedOrder: make(map[*ScoreTimewiseMeasurePart]uint),
		ScoreTimewiseMeasurePart_orderStaged: make(map[uint]*ScoreTimewiseMeasurePart),
		ScoreTimewiseMeasureParts_reference:  make(map[*ScoreTimewiseMeasurePart]*ScoreTimewiseMeasurePart),

		Work_stagedOrder: make(map[*Work]uint),
		Work_orderStaged: make(map[uint]*Work),
		Works_reference:  make(map[*Work]*Work),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"Credit": &CreditUnmarshaller{},

			"Defaults": &DefaultsUnmarshaller{},

			"EmptyFont": &EmptyFontUnmarshaller{},

			"GroupBarline": &GroupBarlineUnmarshaller{},

			"GroupName": &GroupNameUnmarshaller{},

			"GroupSymbol": &GroupSymbolUnmarshaller{},

			"InstrumentLink": &InstrumentLinkUnmarshaller{},

			"LyricFont": &LyricFontUnmarshaller{},

			"LyricLanguage": &LyricLanguageUnmarshaller{},

			"Measure": &MeasureUnmarshaller{},

			"MeasureMusicDataChoice": &MeasureMusicDataChoiceUnmarshaller{},

			"Opus": &OpusUnmarshaller{},

			"Part": &PartUnmarshaller{},

			"PartGroup": &PartGroupUnmarshaller{},

			"PartLink": &PartLinkUnmarshaller{},

			"PartList": &PartListUnmarshaller{},

			"PartName": &PartNameUnmarshaller{},

			"Player": &PlayerUnmarshaller{},

			"ScoreInstrument": &ScoreInstrumentUnmarshaller{},

			"ScorePart": &ScorePartUnmarshaller{},

			"ScorePartwise": &ScorePartwiseUnmarshaller{},

			"ScoreTimewise": &ScoreTimewiseUnmarshaller{},

			"ScoreTimewiseMeasure": &ScoreTimewiseMeasureUnmarshaller{},

			"ScoreTimewiseMeasurePart": &ScoreTimewiseMeasurePartUnmarshaller{},

			"Work": &WorkUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["Credit"] = len(stage.Credits)
	stage.Map_GongStructName_InstancesNb["Defaults"] = len(stage.Defaultss)
	stage.Map_GongStructName_InstancesNb["EmptyFont"] = len(stage.EmptyFonts)
	stage.Map_GongStructName_InstancesNb["GroupBarline"] = len(stage.GroupBarlines)
	stage.Map_GongStructName_InstancesNb["GroupName"] = len(stage.GroupNames)
	stage.Map_GongStructName_InstancesNb["GroupSymbol"] = len(stage.GroupSymbols)
	stage.Map_GongStructName_InstancesNb["InstrumentLink"] = len(stage.InstrumentLinks)
	stage.Map_GongStructName_InstancesNb["LyricFont"] = len(stage.LyricFonts)
	stage.Map_GongStructName_InstancesNb["LyricLanguage"] = len(stage.LyricLanguages)
	stage.Map_GongStructName_InstancesNb["Measure"] = len(stage.Measures)
	stage.Map_GongStructName_InstancesNb["MeasureMusicDataChoice"] = len(stage.MeasureMusicDataChoices)
	stage.Map_GongStructName_InstancesNb["Opus"] = len(stage.Opuss)
	stage.Map_GongStructName_InstancesNb["Part"] = len(stage.Parts)
	stage.Map_GongStructName_InstancesNb["PartGroup"] = len(stage.PartGroups)
	stage.Map_GongStructName_InstancesNb["PartLink"] = len(stage.PartLinks)
	stage.Map_GongStructName_InstancesNb["PartList"] = len(stage.PartLists)
	stage.Map_GongStructName_InstancesNb["PartName"] = len(stage.PartNames)
	stage.Map_GongStructName_InstancesNb["Player"] = len(stage.Players)
	stage.Map_GongStructName_InstancesNb["ScoreInstrument"] = len(stage.ScoreInstruments)
	stage.Map_GongStructName_InstancesNb["ScorePart"] = len(stage.ScoreParts)
	stage.Map_GongStructName_InstancesNb["ScorePartwise"] = len(stage.ScorePartwises)
	stage.Map_GongStructName_InstancesNb["ScoreTimewise"] = len(stage.ScoreTimewises)
	stage.Map_GongStructName_InstancesNb["ScoreTimewiseMeasure"] = len(stage.ScoreTimewiseMeasures)
	stage.Map_GongStructName_InstancesNb["ScoreTimewiseMeasurePart"] = len(stage.ScoreTimewiseMeasureParts)
	stage.Map_GongStructName_InstancesNb["Work"] = len(stage.Works)
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
// Stage puts credit to the model stage
func (credit *Credit) Stage(stage *Stage) *Credit {
	__gong__stage(stage.Credits, stage.Credit_stagedOrder, stage.Credit_orderStaged, &stage.CreditOrder, stage.Credits_mapString, credit, credit.Name)
	return credit
}

// StagePreserveOrder puts credit to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CreditOrder
// - update stage.CreditOrder accordingly
func (credit *Credit) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Credits, stage.Credit_stagedOrder, stage.Credit_orderStaged, &stage.CreditOrder, stage.Credits_mapString, credit, order, credit.Name)
}

// Unstage removes credit off the model stage
func (credit *Credit) Unstage(stage *Stage) *Credit {
	__gong__unstage(stage.Credits, stage.Credits_mapString, credit, credit.Name)
	return credit
}

// UnstageVoid removes credit off the model stage
func (credit *Credit) UnstageVoid(stage *Stage) {
	credit.Unstage(stage)
}

func (credit *Credit) StageVoid(stage *Stage) {
	credit.Stage(stage)
}

// for satisfaction of GongStruct interface
func (credit *Credit) GetName() (res string) {
	return credit.Name
}

// for satisfaction of GongStruct interface
func (credit *Credit) SetName(name string) {
	credit.Name = name
}

// Stage puts defaults to the model stage
func (defaults *Defaults) Stage(stage *Stage) *Defaults {
	__gong__stage(stage.Defaultss, stage.Defaults_stagedOrder, stage.Defaults_orderStaged, &stage.DefaultsOrder, stage.Defaultss_mapString, defaults, defaults.Name)
	return defaults
}

// StagePreserveOrder puts defaults to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DefaultsOrder
// - update stage.DefaultsOrder accordingly
func (defaults *Defaults) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Defaultss, stage.Defaults_stagedOrder, stage.Defaults_orderStaged, &stage.DefaultsOrder, stage.Defaultss_mapString, defaults, order, defaults.Name)
}

// Unstage removes defaults off the model stage
func (defaults *Defaults) Unstage(stage *Stage) *Defaults {
	__gong__unstage(stage.Defaultss, stage.Defaultss_mapString, defaults, defaults.Name)
	return defaults
}

// UnstageVoid removes defaults off the model stage
func (defaults *Defaults) UnstageVoid(stage *Stage) {
	defaults.Unstage(stage)
}

func (defaults *Defaults) StageVoid(stage *Stage) {
	defaults.Stage(stage)
}

// for satisfaction of GongStruct interface
func (defaults *Defaults) GetName() (res string) {
	return defaults.Name
}

// for satisfaction of GongStruct interface
func (defaults *Defaults) SetName(name string) {
	defaults.Name = name
}

// Stage puts emptyfont to the model stage
func (emptyfont *EmptyFont) Stage(stage *Stage) *EmptyFont {
	__gong__stage(stage.EmptyFonts, stage.EmptyFont_stagedOrder, stage.EmptyFont_orderStaged, &stage.EmptyFontOrder, stage.EmptyFonts_mapString, emptyfont, emptyfont.Name)
	return emptyfont
}

// StagePreserveOrder puts emptyfont to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EmptyFontOrder
// - update stage.EmptyFontOrder accordingly
func (emptyfont *EmptyFont) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.EmptyFonts, stage.EmptyFont_stagedOrder, stage.EmptyFont_orderStaged, &stage.EmptyFontOrder, stage.EmptyFonts_mapString, emptyfont, order, emptyfont.Name)
}

// Unstage removes emptyfont off the model stage
func (emptyfont *EmptyFont) Unstage(stage *Stage) *EmptyFont {
	__gong__unstage(stage.EmptyFonts, stage.EmptyFonts_mapString, emptyfont, emptyfont.Name)
	return emptyfont
}

// UnstageVoid removes emptyfont off the model stage
func (emptyfont *EmptyFont) UnstageVoid(stage *Stage) {
	emptyfont.Unstage(stage)
}

func (emptyfont *EmptyFont) StageVoid(stage *Stage) {
	emptyfont.Stage(stage)
}

// for satisfaction of GongStruct interface
func (emptyfont *EmptyFont) GetName() (res string) {
	return emptyfont.Name
}

// for satisfaction of GongStruct interface
func (emptyfont *EmptyFont) SetName(name string) {
	emptyfont.Name = name
}

// Stage puts groupbarline to the model stage
func (groupbarline *GroupBarline) Stage(stage *Stage) *GroupBarline {
	__gong__stage(stage.GroupBarlines, stage.GroupBarline_stagedOrder, stage.GroupBarline_orderStaged, &stage.GroupBarlineOrder, stage.GroupBarlines_mapString, groupbarline, groupbarline.Name)
	return groupbarline
}

// StagePreserveOrder puts groupbarline to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GroupBarlineOrder
// - update stage.GroupBarlineOrder accordingly
func (groupbarline *GroupBarline) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.GroupBarlines, stage.GroupBarline_stagedOrder, stage.GroupBarline_orderStaged, &stage.GroupBarlineOrder, stage.GroupBarlines_mapString, groupbarline, order, groupbarline.Name)
}

// Unstage removes groupbarline off the model stage
func (groupbarline *GroupBarline) Unstage(stage *Stage) *GroupBarline {
	__gong__unstage(stage.GroupBarlines, stage.GroupBarlines_mapString, groupbarline, groupbarline.Name)
	return groupbarline
}

// UnstageVoid removes groupbarline off the model stage
func (groupbarline *GroupBarline) UnstageVoid(stage *Stage) {
	groupbarline.Unstage(stage)
}

func (groupbarline *GroupBarline) StageVoid(stage *Stage) {
	groupbarline.Stage(stage)
}

// for satisfaction of GongStruct interface
func (groupbarline *GroupBarline) GetName() (res string) {
	return groupbarline.Name
}

// for satisfaction of GongStruct interface
func (groupbarline *GroupBarline) SetName(name string) {
	groupbarline.Name = name
}

// Stage puts groupname to the model stage
func (groupname *GroupName) Stage(stage *Stage) *GroupName {
	__gong__stage(stage.GroupNames, stage.GroupName_stagedOrder, stage.GroupName_orderStaged, &stage.GroupNameOrder, stage.GroupNames_mapString, groupname, groupname.Name)
	return groupname
}

// StagePreserveOrder puts groupname to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GroupNameOrder
// - update stage.GroupNameOrder accordingly
func (groupname *GroupName) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.GroupNames, stage.GroupName_stagedOrder, stage.GroupName_orderStaged, &stage.GroupNameOrder, stage.GroupNames_mapString, groupname, order, groupname.Name)
}

// Unstage removes groupname off the model stage
func (groupname *GroupName) Unstage(stage *Stage) *GroupName {
	__gong__unstage(stage.GroupNames, stage.GroupNames_mapString, groupname, groupname.Name)
	return groupname
}

// UnstageVoid removes groupname off the model stage
func (groupname *GroupName) UnstageVoid(stage *Stage) {
	groupname.Unstage(stage)
}

func (groupname *GroupName) StageVoid(stage *Stage) {
	groupname.Stage(stage)
}

// for satisfaction of GongStruct interface
func (groupname *GroupName) GetName() (res string) {
	return groupname.Name
}

// for satisfaction of GongStruct interface
func (groupname *GroupName) SetName(name string) {
	groupname.Name = name
}

// Stage puts groupsymbol to the model stage
func (groupsymbol *GroupSymbol) Stage(stage *Stage) *GroupSymbol {
	__gong__stage(stage.GroupSymbols, stage.GroupSymbol_stagedOrder, stage.GroupSymbol_orderStaged, &stage.GroupSymbolOrder, stage.GroupSymbols_mapString, groupsymbol, groupsymbol.Name)
	return groupsymbol
}

// StagePreserveOrder puts groupsymbol to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GroupSymbolOrder
// - update stage.GroupSymbolOrder accordingly
func (groupsymbol *GroupSymbol) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.GroupSymbols, stage.GroupSymbol_stagedOrder, stage.GroupSymbol_orderStaged, &stage.GroupSymbolOrder, stage.GroupSymbols_mapString, groupsymbol, order, groupsymbol.Name)
}

// Unstage removes groupsymbol off the model stage
func (groupsymbol *GroupSymbol) Unstage(stage *Stage) *GroupSymbol {
	__gong__unstage(stage.GroupSymbols, stage.GroupSymbols_mapString, groupsymbol, groupsymbol.Name)
	return groupsymbol
}

// UnstageVoid removes groupsymbol off the model stage
func (groupsymbol *GroupSymbol) UnstageVoid(stage *Stage) {
	groupsymbol.Unstage(stage)
}

func (groupsymbol *GroupSymbol) StageVoid(stage *Stage) {
	groupsymbol.Stage(stage)
}

// for satisfaction of GongStruct interface
func (groupsymbol *GroupSymbol) GetName() (res string) {
	return groupsymbol.Name
}

// for satisfaction of GongStruct interface
func (groupsymbol *GroupSymbol) SetName(name string) {
	groupsymbol.Name = name
}

// Stage puts instrumentlink to the model stage
func (instrumentlink *InstrumentLink) Stage(stage *Stage) *InstrumentLink {
	__gong__stage(stage.InstrumentLinks, stage.InstrumentLink_stagedOrder, stage.InstrumentLink_orderStaged, &stage.InstrumentLinkOrder, stage.InstrumentLinks_mapString, instrumentlink, instrumentlink.Name)
	return instrumentlink
}

// StagePreserveOrder puts instrumentlink to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.InstrumentLinkOrder
// - update stage.InstrumentLinkOrder accordingly
func (instrumentlink *InstrumentLink) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.InstrumentLinks, stage.InstrumentLink_stagedOrder, stage.InstrumentLink_orderStaged, &stage.InstrumentLinkOrder, stage.InstrumentLinks_mapString, instrumentlink, order, instrumentlink.Name)
}

// Unstage removes instrumentlink off the model stage
func (instrumentlink *InstrumentLink) Unstage(stage *Stage) *InstrumentLink {
	__gong__unstage(stage.InstrumentLinks, stage.InstrumentLinks_mapString, instrumentlink, instrumentlink.Name)
	return instrumentlink
}

// UnstageVoid removes instrumentlink off the model stage
func (instrumentlink *InstrumentLink) UnstageVoid(stage *Stage) {
	instrumentlink.Unstage(stage)
}

func (instrumentlink *InstrumentLink) StageVoid(stage *Stage) {
	instrumentlink.Stage(stage)
}

// for satisfaction of GongStruct interface
func (instrumentlink *InstrumentLink) GetName() (res string) {
	return instrumentlink.Name
}

// for satisfaction of GongStruct interface
func (instrumentlink *InstrumentLink) SetName(name string) {
	instrumentlink.Name = name
}

// Stage puts lyricfont to the model stage
func (lyricfont *LyricFont) Stage(stage *Stage) *LyricFont {
	__gong__stage(stage.LyricFonts, stage.LyricFont_stagedOrder, stage.LyricFont_orderStaged, &stage.LyricFontOrder, stage.LyricFonts_mapString, lyricfont, lyricfont.Name)
	return lyricfont
}

// StagePreserveOrder puts lyricfont to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LyricFontOrder
// - update stage.LyricFontOrder accordingly
func (lyricfont *LyricFont) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.LyricFonts, stage.LyricFont_stagedOrder, stage.LyricFont_orderStaged, &stage.LyricFontOrder, stage.LyricFonts_mapString, lyricfont, order, lyricfont.Name)
}

// Unstage removes lyricfont off the model stage
func (lyricfont *LyricFont) Unstage(stage *Stage) *LyricFont {
	__gong__unstage(stage.LyricFonts, stage.LyricFonts_mapString, lyricfont, lyricfont.Name)
	return lyricfont
}

// UnstageVoid removes lyricfont off the model stage
func (lyricfont *LyricFont) UnstageVoid(stage *Stage) {
	lyricfont.Unstage(stage)
}

func (lyricfont *LyricFont) StageVoid(stage *Stage) {
	lyricfont.Stage(stage)
}

// for satisfaction of GongStruct interface
func (lyricfont *LyricFont) GetName() (res string) {
	return lyricfont.Name
}

// for satisfaction of GongStruct interface
func (lyricfont *LyricFont) SetName(name string) {
	lyricfont.Name = name
}

// Stage puts lyriclanguage to the model stage
func (lyriclanguage *LyricLanguage) Stage(stage *Stage) *LyricLanguage {
	__gong__stage(stage.LyricLanguages, stage.LyricLanguage_stagedOrder, stage.LyricLanguage_orderStaged, &stage.LyricLanguageOrder, stage.LyricLanguages_mapString, lyriclanguage, lyriclanguage.Name)
	return lyriclanguage
}

// StagePreserveOrder puts lyriclanguage to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LyricLanguageOrder
// - update stage.LyricLanguageOrder accordingly
func (lyriclanguage *LyricLanguage) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.LyricLanguages, stage.LyricLanguage_stagedOrder, stage.LyricLanguage_orderStaged, &stage.LyricLanguageOrder, stage.LyricLanguages_mapString, lyriclanguage, order, lyriclanguage.Name)
}

// Unstage removes lyriclanguage off the model stage
func (lyriclanguage *LyricLanguage) Unstage(stage *Stage) *LyricLanguage {
	__gong__unstage(stage.LyricLanguages, stage.LyricLanguages_mapString, lyriclanguage, lyriclanguage.Name)
	return lyriclanguage
}

// UnstageVoid removes lyriclanguage off the model stage
func (lyriclanguage *LyricLanguage) UnstageVoid(stage *Stage) {
	lyriclanguage.Unstage(stage)
}

func (lyriclanguage *LyricLanguage) StageVoid(stage *Stage) {
	lyriclanguage.Stage(stage)
}

// for satisfaction of GongStruct interface
func (lyriclanguage *LyricLanguage) GetName() (res string) {
	return lyriclanguage.Name
}

// for satisfaction of GongStruct interface
func (lyriclanguage *LyricLanguage) SetName(name string) {
	lyriclanguage.Name = name
}

// Stage puts measure to the model stage
func (measure *Measure) Stage(stage *Stage) *Measure {
	__gong__stage(stage.Measures, stage.Measure_stagedOrder, stage.Measure_orderStaged, &stage.MeasureOrder, stage.Measures_mapString, measure, measure.Name)
	return measure
}

// StagePreserveOrder puts measure to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MeasureOrder
// - update stage.MeasureOrder accordingly
func (measure *Measure) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Measures, stage.Measure_stagedOrder, stage.Measure_orderStaged, &stage.MeasureOrder, stage.Measures_mapString, measure, order, measure.Name)
}

// Unstage removes measure off the model stage
func (measure *Measure) Unstage(stage *Stage) *Measure {
	__gong__unstage(stage.Measures, stage.Measures_mapString, measure, measure.Name)
	return measure
}

// UnstageVoid removes measure off the model stage
func (measure *Measure) UnstageVoid(stage *Stage) {
	measure.Unstage(stage)
}

func (measure *Measure) StageVoid(stage *Stage) {
	measure.Stage(stage)
}

// for satisfaction of GongStruct interface
func (measure *Measure) GetName() (res string) {
	return measure.Name
}

// for satisfaction of GongStruct interface
func (measure *Measure) SetName(name string) {
	measure.Name = name
}

// Stage puts measuremusicdatachoice to the model stage
func (measuremusicdatachoice *MeasureMusicDataChoice) Stage(stage *Stage) *MeasureMusicDataChoice {
	__gong__stage(stage.MeasureMusicDataChoices, stage.MeasureMusicDataChoice_stagedOrder, stage.MeasureMusicDataChoice_orderStaged, &stage.MeasureMusicDataChoiceOrder, stage.MeasureMusicDataChoices_mapString, measuremusicdatachoice, measuremusicdatachoice.Name)
	return measuremusicdatachoice
}

// StagePreserveOrder puts measuremusicdatachoice to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MeasureMusicDataChoiceOrder
// - update stage.MeasureMusicDataChoiceOrder accordingly
func (measuremusicdatachoice *MeasureMusicDataChoice) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MeasureMusicDataChoices, stage.MeasureMusicDataChoice_stagedOrder, stage.MeasureMusicDataChoice_orderStaged, &stage.MeasureMusicDataChoiceOrder, stage.MeasureMusicDataChoices_mapString, measuremusicdatachoice, order, measuremusicdatachoice.Name)
}

// Unstage removes measuremusicdatachoice off the model stage
func (measuremusicdatachoice *MeasureMusicDataChoice) Unstage(stage *Stage) *MeasureMusicDataChoice {
	__gong__unstage(stage.MeasureMusicDataChoices, stage.MeasureMusicDataChoices_mapString, measuremusicdatachoice, measuremusicdatachoice.Name)
	return measuremusicdatachoice
}

// UnstageVoid removes measuremusicdatachoice off the model stage
func (measuremusicdatachoice *MeasureMusicDataChoice) UnstageVoid(stage *Stage) {
	measuremusicdatachoice.Unstage(stage)
}

func (measuremusicdatachoice *MeasureMusicDataChoice) StageVoid(stage *Stage) {
	measuremusicdatachoice.Stage(stage)
}

// for satisfaction of GongStruct interface
func (measuremusicdatachoice *MeasureMusicDataChoice) GetName() (res string) {
	return measuremusicdatachoice.Name
}

// for satisfaction of GongStruct interface
func (measuremusicdatachoice *MeasureMusicDataChoice) SetName(name string) {
	measuremusicdatachoice.Name = name
}

// Stage puts opus to the model stage
func (opus *Opus) Stage(stage *Stage) *Opus {
	__gong__stage(stage.Opuss, stage.Opus_stagedOrder, stage.Opus_orderStaged, &stage.OpusOrder, stage.Opuss_mapString, opus, opus.Name)
	return opus
}

// StagePreserveOrder puts opus to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.OpusOrder
// - update stage.OpusOrder accordingly
func (opus *Opus) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Opuss, stage.Opus_stagedOrder, stage.Opus_orderStaged, &stage.OpusOrder, stage.Opuss_mapString, opus, order, opus.Name)
}

// Unstage removes opus off the model stage
func (opus *Opus) Unstage(stage *Stage) *Opus {
	__gong__unstage(stage.Opuss, stage.Opuss_mapString, opus, opus.Name)
	return opus
}

// UnstageVoid removes opus off the model stage
func (opus *Opus) UnstageVoid(stage *Stage) {
	opus.Unstage(stage)
}

func (opus *Opus) StageVoid(stage *Stage) {
	opus.Stage(stage)
}

// for satisfaction of GongStruct interface
func (opus *Opus) GetName() (res string) {
	return opus.Name
}

// for satisfaction of GongStruct interface
func (opus *Opus) SetName(name string) {
	opus.Name = name
}

// Stage puts part to the model stage
func (part *Part) Stage(stage *Stage) *Part {
	__gong__stage(stage.Parts, stage.Part_stagedOrder, stage.Part_orderStaged, &stage.PartOrder, stage.Parts_mapString, part, part.Name)
	return part
}

// StagePreserveOrder puts part to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PartOrder
// - update stage.PartOrder accordingly
func (part *Part) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Parts, stage.Part_stagedOrder, stage.Part_orderStaged, &stage.PartOrder, stage.Parts_mapString, part, order, part.Name)
}

// Unstage removes part off the model stage
func (part *Part) Unstage(stage *Stage) *Part {
	__gong__unstage(stage.Parts, stage.Parts_mapString, part, part.Name)
	return part
}

// UnstageVoid removes part off the model stage
func (part *Part) UnstageVoid(stage *Stage) {
	part.Unstage(stage)
}

func (part *Part) StageVoid(stage *Stage) {
	part.Stage(stage)
}

// for satisfaction of GongStruct interface
func (part *Part) GetName() (res string) {
	return part.Name
}

// for satisfaction of GongStruct interface
func (part *Part) SetName(name string) {
	part.Name = name
}

// Stage puts partgroup to the model stage
func (partgroup *PartGroup) Stage(stage *Stage) *PartGroup {
	__gong__stage(stage.PartGroups, stage.PartGroup_stagedOrder, stage.PartGroup_orderStaged, &stage.PartGroupOrder, stage.PartGroups_mapString, partgroup, partgroup.Name)
	return partgroup
}

// StagePreserveOrder puts partgroup to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PartGroupOrder
// - update stage.PartGroupOrder accordingly
func (partgroup *PartGroup) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.PartGroups, stage.PartGroup_stagedOrder, stage.PartGroup_orderStaged, &stage.PartGroupOrder, stage.PartGroups_mapString, partgroup, order, partgroup.Name)
}

// Unstage removes partgroup off the model stage
func (partgroup *PartGroup) Unstage(stage *Stage) *PartGroup {
	__gong__unstage(stage.PartGroups, stage.PartGroups_mapString, partgroup, partgroup.Name)
	return partgroup
}

// UnstageVoid removes partgroup off the model stage
func (partgroup *PartGroup) UnstageVoid(stage *Stage) {
	partgroup.Unstage(stage)
}

func (partgroup *PartGroup) StageVoid(stage *Stage) {
	partgroup.Stage(stage)
}

// for satisfaction of GongStruct interface
func (partgroup *PartGroup) GetName() (res string) {
	return partgroup.Name
}

// for satisfaction of GongStruct interface
func (partgroup *PartGroup) SetName(name string) {
	partgroup.Name = name
}

// Stage puts partlink to the model stage
func (partlink *PartLink) Stage(stage *Stage) *PartLink {
	__gong__stage(stage.PartLinks, stage.PartLink_stagedOrder, stage.PartLink_orderStaged, &stage.PartLinkOrder, stage.PartLinks_mapString, partlink, partlink.Name)
	return partlink
}

// StagePreserveOrder puts partlink to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PartLinkOrder
// - update stage.PartLinkOrder accordingly
func (partlink *PartLink) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.PartLinks, stage.PartLink_stagedOrder, stage.PartLink_orderStaged, &stage.PartLinkOrder, stage.PartLinks_mapString, partlink, order, partlink.Name)
}

// Unstage removes partlink off the model stage
func (partlink *PartLink) Unstage(stage *Stage) *PartLink {
	__gong__unstage(stage.PartLinks, stage.PartLinks_mapString, partlink, partlink.Name)
	return partlink
}

// UnstageVoid removes partlink off the model stage
func (partlink *PartLink) UnstageVoid(stage *Stage) {
	partlink.Unstage(stage)
}

func (partlink *PartLink) StageVoid(stage *Stage) {
	partlink.Stage(stage)
}

// for satisfaction of GongStruct interface
func (partlink *PartLink) GetName() (res string) {
	return partlink.Name
}

// for satisfaction of GongStruct interface
func (partlink *PartLink) SetName(name string) {
	partlink.Name = name
}

// Stage puts partlist to the model stage
func (partlist *PartList) Stage(stage *Stage) *PartList {
	__gong__stage(stage.PartLists, stage.PartList_stagedOrder, stage.PartList_orderStaged, &stage.PartListOrder, stage.PartLists_mapString, partlist, partlist.Name)
	return partlist
}

// StagePreserveOrder puts partlist to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PartListOrder
// - update stage.PartListOrder accordingly
func (partlist *PartList) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.PartLists, stage.PartList_stagedOrder, stage.PartList_orderStaged, &stage.PartListOrder, stage.PartLists_mapString, partlist, order, partlist.Name)
}

// Unstage removes partlist off the model stage
func (partlist *PartList) Unstage(stage *Stage) *PartList {
	__gong__unstage(stage.PartLists, stage.PartLists_mapString, partlist, partlist.Name)
	return partlist
}

// UnstageVoid removes partlist off the model stage
func (partlist *PartList) UnstageVoid(stage *Stage) {
	partlist.Unstage(stage)
}

func (partlist *PartList) StageVoid(stage *Stage) {
	partlist.Stage(stage)
}

// for satisfaction of GongStruct interface
func (partlist *PartList) GetName() (res string) {
	return partlist.Name
}

// for satisfaction of GongStruct interface
func (partlist *PartList) SetName(name string) {
	partlist.Name = name
}

// Stage puts partname to the model stage
func (partname *PartName) Stage(stage *Stage) *PartName {
	__gong__stage(stage.PartNames, stage.PartName_stagedOrder, stage.PartName_orderStaged, &stage.PartNameOrder, stage.PartNames_mapString, partname, partname.Name)
	return partname
}

// StagePreserveOrder puts partname to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PartNameOrder
// - update stage.PartNameOrder accordingly
func (partname *PartName) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.PartNames, stage.PartName_stagedOrder, stage.PartName_orderStaged, &stage.PartNameOrder, stage.PartNames_mapString, partname, order, partname.Name)
}

// Unstage removes partname off the model stage
func (partname *PartName) Unstage(stage *Stage) *PartName {
	__gong__unstage(stage.PartNames, stage.PartNames_mapString, partname, partname.Name)
	return partname
}

// UnstageVoid removes partname off the model stage
func (partname *PartName) UnstageVoid(stage *Stage) {
	partname.Unstage(stage)
}

func (partname *PartName) StageVoid(stage *Stage) {
	partname.Stage(stage)
}

// for satisfaction of GongStruct interface
func (partname *PartName) GetName() (res string) {
	return partname.Name
}

// for satisfaction of GongStruct interface
func (partname *PartName) SetName(name string) {
	partname.Name = name
}

// Stage puts player to the model stage
func (player *Player) Stage(stage *Stage) *Player {
	__gong__stage(stage.Players, stage.Player_stagedOrder, stage.Player_orderStaged, &stage.PlayerOrder, stage.Players_mapString, player, player.Name)
	return player
}

// StagePreserveOrder puts player to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PlayerOrder
// - update stage.PlayerOrder accordingly
func (player *Player) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Players, stage.Player_stagedOrder, stage.Player_orderStaged, &stage.PlayerOrder, stage.Players_mapString, player, order, player.Name)
}

// Unstage removes player off the model stage
func (player *Player) Unstage(stage *Stage) *Player {
	__gong__unstage(stage.Players, stage.Players_mapString, player, player.Name)
	return player
}

// UnstageVoid removes player off the model stage
func (player *Player) UnstageVoid(stage *Stage) {
	player.Unstage(stage)
}

func (player *Player) StageVoid(stage *Stage) {
	player.Stage(stage)
}

// for satisfaction of GongStruct interface
func (player *Player) GetName() (res string) {
	return player.Name
}

// for satisfaction of GongStruct interface
func (player *Player) SetName(name string) {
	player.Name = name
}

// Stage puts scoreinstrument to the model stage
func (scoreinstrument *ScoreInstrument) Stage(stage *Stage) *ScoreInstrument {
	__gong__stage(stage.ScoreInstruments, stage.ScoreInstrument_stagedOrder, stage.ScoreInstrument_orderStaged, &stage.ScoreInstrumentOrder, stage.ScoreInstruments_mapString, scoreinstrument, scoreinstrument.Name)
	return scoreinstrument
}

// StagePreserveOrder puts scoreinstrument to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ScoreInstrumentOrder
// - update stage.ScoreInstrumentOrder accordingly
func (scoreinstrument *ScoreInstrument) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ScoreInstruments, stage.ScoreInstrument_stagedOrder, stage.ScoreInstrument_orderStaged, &stage.ScoreInstrumentOrder, stage.ScoreInstruments_mapString, scoreinstrument, order, scoreinstrument.Name)
}

// Unstage removes scoreinstrument off the model stage
func (scoreinstrument *ScoreInstrument) Unstage(stage *Stage) *ScoreInstrument {
	__gong__unstage(stage.ScoreInstruments, stage.ScoreInstruments_mapString, scoreinstrument, scoreinstrument.Name)
	return scoreinstrument
}

// UnstageVoid removes scoreinstrument off the model stage
func (scoreinstrument *ScoreInstrument) UnstageVoid(stage *Stage) {
	scoreinstrument.Unstage(stage)
}

func (scoreinstrument *ScoreInstrument) StageVoid(stage *Stage) {
	scoreinstrument.Stage(stage)
}

// for satisfaction of GongStruct interface
func (scoreinstrument *ScoreInstrument) GetName() (res string) {
	return scoreinstrument.Name
}

// for satisfaction of GongStruct interface
func (scoreinstrument *ScoreInstrument) SetName(name string) {
	scoreinstrument.Name = name
}

// Stage puts scorepart to the model stage
func (scorepart *ScorePart) Stage(stage *Stage) *ScorePart {
	__gong__stage(stage.ScoreParts, stage.ScorePart_stagedOrder, stage.ScorePart_orderStaged, &stage.ScorePartOrder, stage.ScoreParts_mapString, scorepart, scorepart.Name)
	return scorepart
}

// StagePreserveOrder puts scorepart to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ScorePartOrder
// - update stage.ScorePartOrder accordingly
func (scorepart *ScorePart) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ScoreParts, stage.ScorePart_stagedOrder, stage.ScorePart_orderStaged, &stage.ScorePartOrder, stage.ScoreParts_mapString, scorepart, order, scorepart.Name)
}

// Unstage removes scorepart off the model stage
func (scorepart *ScorePart) Unstage(stage *Stage) *ScorePart {
	__gong__unstage(stage.ScoreParts, stage.ScoreParts_mapString, scorepart, scorepart.Name)
	return scorepart
}

// UnstageVoid removes scorepart off the model stage
func (scorepart *ScorePart) UnstageVoid(stage *Stage) {
	scorepart.Unstage(stage)
}

func (scorepart *ScorePart) StageVoid(stage *Stage) {
	scorepart.Stage(stage)
}

// for satisfaction of GongStruct interface
func (scorepart *ScorePart) GetName() (res string) {
	return scorepart.Name
}

// for satisfaction of GongStruct interface
func (scorepart *ScorePart) SetName(name string) {
	scorepart.Name = name
}

// Stage puts scorepartwise to the model stage
func (scorepartwise *ScorePartwise) Stage(stage *Stage) *ScorePartwise {
	__gong__stage(stage.ScorePartwises, stage.ScorePartwise_stagedOrder, stage.ScorePartwise_orderStaged, &stage.ScorePartwiseOrder, stage.ScorePartwises_mapString, scorepartwise, scorepartwise.Name)
	return scorepartwise
}

// StagePreserveOrder puts scorepartwise to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ScorePartwiseOrder
// - update stage.ScorePartwiseOrder accordingly
func (scorepartwise *ScorePartwise) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ScorePartwises, stage.ScorePartwise_stagedOrder, stage.ScorePartwise_orderStaged, &stage.ScorePartwiseOrder, stage.ScorePartwises_mapString, scorepartwise, order, scorepartwise.Name)
}

// Unstage removes scorepartwise off the model stage
func (scorepartwise *ScorePartwise) Unstage(stage *Stage) *ScorePartwise {
	__gong__unstage(stage.ScorePartwises, stage.ScorePartwises_mapString, scorepartwise, scorepartwise.Name)
	return scorepartwise
}

// UnstageVoid removes scorepartwise off the model stage
func (scorepartwise *ScorePartwise) UnstageVoid(stage *Stage) {
	scorepartwise.Unstage(stage)
}

func (scorepartwise *ScorePartwise) StageVoid(stage *Stage) {
	scorepartwise.Stage(stage)
}

// for satisfaction of GongStruct interface
func (scorepartwise *ScorePartwise) GetName() (res string) {
	return scorepartwise.Name
}

// for satisfaction of GongStruct interface
func (scorepartwise *ScorePartwise) SetName(name string) {
	scorepartwise.Name = name
}

// Stage puts scoretimewise to the model stage
func (scoretimewise *ScoreTimewise) Stage(stage *Stage) *ScoreTimewise {
	__gong__stage(stage.ScoreTimewises, stage.ScoreTimewise_stagedOrder, stage.ScoreTimewise_orderStaged, &stage.ScoreTimewiseOrder, stage.ScoreTimewises_mapString, scoretimewise, scoretimewise.Name)
	return scoretimewise
}

// StagePreserveOrder puts scoretimewise to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ScoreTimewiseOrder
// - update stage.ScoreTimewiseOrder accordingly
func (scoretimewise *ScoreTimewise) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ScoreTimewises, stage.ScoreTimewise_stagedOrder, stage.ScoreTimewise_orderStaged, &stage.ScoreTimewiseOrder, stage.ScoreTimewises_mapString, scoretimewise, order, scoretimewise.Name)
}

// Unstage removes scoretimewise off the model stage
func (scoretimewise *ScoreTimewise) Unstage(stage *Stage) *ScoreTimewise {
	__gong__unstage(stage.ScoreTimewises, stage.ScoreTimewises_mapString, scoretimewise, scoretimewise.Name)
	return scoretimewise
}

// UnstageVoid removes scoretimewise off the model stage
func (scoretimewise *ScoreTimewise) UnstageVoid(stage *Stage) {
	scoretimewise.Unstage(stage)
}

func (scoretimewise *ScoreTimewise) StageVoid(stage *Stage) {
	scoretimewise.Stage(stage)
}

// for satisfaction of GongStruct interface
func (scoretimewise *ScoreTimewise) GetName() (res string) {
	return scoretimewise.Name
}

// for satisfaction of GongStruct interface
func (scoretimewise *ScoreTimewise) SetName(name string) {
	scoretimewise.Name = name
}

// Stage puts scoretimewisemeasure to the model stage
func (scoretimewisemeasure *ScoreTimewiseMeasure) Stage(stage *Stage) *ScoreTimewiseMeasure {
	__gong__stage(stage.ScoreTimewiseMeasures, stage.ScoreTimewiseMeasure_stagedOrder, stage.ScoreTimewiseMeasure_orderStaged, &stage.ScoreTimewiseMeasureOrder, stage.ScoreTimewiseMeasures_mapString, scoretimewisemeasure, scoretimewisemeasure.Name)
	return scoretimewisemeasure
}

// StagePreserveOrder puts scoretimewisemeasure to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ScoreTimewiseMeasureOrder
// - update stage.ScoreTimewiseMeasureOrder accordingly
func (scoretimewisemeasure *ScoreTimewiseMeasure) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ScoreTimewiseMeasures, stage.ScoreTimewiseMeasure_stagedOrder, stage.ScoreTimewiseMeasure_orderStaged, &stage.ScoreTimewiseMeasureOrder, stage.ScoreTimewiseMeasures_mapString, scoretimewisemeasure, order, scoretimewisemeasure.Name)
}

// Unstage removes scoretimewisemeasure off the model stage
func (scoretimewisemeasure *ScoreTimewiseMeasure) Unstage(stage *Stage) *ScoreTimewiseMeasure {
	__gong__unstage(stage.ScoreTimewiseMeasures, stage.ScoreTimewiseMeasures_mapString, scoretimewisemeasure, scoretimewisemeasure.Name)
	return scoretimewisemeasure
}

// UnstageVoid removes scoretimewisemeasure off the model stage
func (scoretimewisemeasure *ScoreTimewiseMeasure) UnstageVoid(stage *Stage) {
	scoretimewisemeasure.Unstage(stage)
}

func (scoretimewisemeasure *ScoreTimewiseMeasure) StageVoid(stage *Stage) {
	scoretimewisemeasure.Stage(stage)
}

// for satisfaction of GongStruct interface
func (scoretimewisemeasure *ScoreTimewiseMeasure) GetName() (res string) {
	return scoretimewisemeasure.Name
}

// for satisfaction of GongStruct interface
func (scoretimewisemeasure *ScoreTimewiseMeasure) SetName(name string) {
	scoretimewisemeasure.Name = name
}

// Stage puts scoretimewisemeasurepart to the model stage
func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) Stage(stage *Stage) *ScoreTimewiseMeasurePart {
	__gong__stage(stage.ScoreTimewiseMeasureParts, stage.ScoreTimewiseMeasurePart_stagedOrder, stage.ScoreTimewiseMeasurePart_orderStaged, &stage.ScoreTimewiseMeasurePartOrder, stage.ScoreTimewiseMeasureParts_mapString, scoretimewisemeasurepart, scoretimewisemeasurepart.Name)
	return scoretimewisemeasurepart
}

// StagePreserveOrder puts scoretimewisemeasurepart to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ScoreTimewiseMeasurePartOrder
// - update stage.ScoreTimewiseMeasurePartOrder accordingly
func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ScoreTimewiseMeasureParts, stage.ScoreTimewiseMeasurePart_stagedOrder, stage.ScoreTimewiseMeasurePart_orderStaged, &stage.ScoreTimewiseMeasurePartOrder, stage.ScoreTimewiseMeasureParts_mapString, scoretimewisemeasurepart, order, scoretimewisemeasurepart.Name)
}

// Unstage removes scoretimewisemeasurepart off the model stage
func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) Unstage(stage *Stage) *ScoreTimewiseMeasurePart {
	__gong__unstage(stage.ScoreTimewiseMeasureParts, stage.ScoreTimewiseMeasureParts_mapString, scoretimewisemeasurepart, scoretimewisemeasurepart.Name)
	return scoretimewisemeasurepart
}

// UnstageVoid removes scoretimewisemeasurepart off the model stage
func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) UnstageVoid(stage *Stage) {
	scoretimewisemeasurepart.Unstage(stage)
}

func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) StageVoid(stage *Stage) {
	scoretimewisemeasurepart.Stage(stage)
}

// for satisfaction of GongStruct interface
func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) GetName() (res string) {
	return scoretimewisemeasurepart.Name
}

// for satisfaction of GongStruct interface
func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) SetName(name string) {
	scoretimewisemeasurepart.Name = name
}

// Stage puts work to the model stage
func (work *Work) Stage(stage *Stage) *Work {
	__gong__stage(stage.Works, stage.Work_stagedOrder, stage.Work_orderStaged, &stage.WorkOrder, stage.Works_mapString, work, work.Name)
	return work
}

// StagePreserveOrder puts work to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.WorkOrder
// - update stage.WorkOrder accordingly
func (work *Work) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Works, stage.Work_stagedOrder, stage.Work_orderStaged, &stage.WorkOrder, stage.Works_mapString, work, order, work.Name)
}

// Unstage removes work off the model stage
func (work *Work) Unstage(stage *Stage) *Work {
	__gong__unstage(stage.Works, stage.Works_mapString, work, work.Name)
	return work
}

// UnstageVoid removes work off the model stage
func (work *Work) UnstageVoid(stage *Stage) {
	work.Unstage(stage)
}

func (work *Work) StageVoid(stage *Stage) {
	work.Stage(stage)
}

// for satisfaction of GongStruct interface
func (work *Work) GetName() (res string) {
	return work.Name
}

// for satisfaction of GongStruct interface
func (work *Work) SetName(name string) {
	work.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.Credits, &stage.Credits_mapString, &stage.Credit_stagedOrder, &stage.CreditOrder)

	__gong__resetStageType(&stage.Defaultss, &stage.Defaultss_mapString, &stage.Defaults_stagedOrder, &stage.DefaultsOrder)

	__gong__resetStageType(&stage.EmptyFonts, &stage.EmptyFonts_mapString, &stage.EmptyFont_stagedOrder, &stage.EmptyFontOrder)

	__gong__resetStageType(&stage.GroupBarlines, &stage.GroupBarlines_mapString, &stage.GroupBarline_stagedOrder, &stage.GroupBarlineOrder)

	__gong__resetStageType(&stage.GroupNames, &stage.GroupNames_mapString, &stage.GroupName_stagedOrder, &stage.GroupNameOrder)

	__gong__resetStageType(&stage.GroupSymbols, &stage.GroupSymbols_mapString, &stage.GroupSymbol_stagedOrder, &stage.GroupSymbolOrder)

	__gong__resetStageType(&stage.InstrumentLinks, &stage.InstrumentLinks_mapString, &stage.InstrumentLink_stagedOrder, &stage.InstrumentLinkOrder)

	__gong__resetStageType(&stage.LyricFonts, &stage.LyricFonts_mapString, &stage.LyricFont_stagedOrder, &stage.LyricFontOrder)

	__gong__resetStageType(&stage.LyricLanguages, &stage.LyricLanguages_mapString, &stage.LyricLanguage_stagedOrder, &stage.LyricLanguageOrder)

	__gong__resetStageType(&stage.Measures, &stage.Measures_mapString, &stage.Measure_stagedOrder, &stage.MeasureOrder)

	__gong__resetStageType(&stage.MeasureMusicDataChoices, &stage.MeasureMusicDataChoices_mapString, &stage.MeasureMusicDataChoice_stagedOrder, &stage.MeasureMusicDataChoiceOrder)

	__gong__resetStageType(&stage.Opuss, &stage.Opuss_mapString, &stage.Opus_stagedOrder, &stage.OpusOrder)

	__gong__resetStageType(&stage.Parts, &stage.Parts_mapString, &stage.Part_stagedOrder, &stage.PartOrder)

	__gong__resetStageType(&stage.PartGroups, &stage.PartGroups_mapString, &stage.PartGroup_stagedOrder, &stage.PartGroupOrder)

	__gong__resetStageType(&stage.PartLinks, &stage.PartLinks_mapString, &stage.PartLink_stagedOrder, &stage.PartLinkOrder)

	__gong__resetStageType(&stage.PartLists, &stage.PartLists_mapString, &stage.PartList_stagedOrder, &stage.PartListOrder)

	__gong__resetStageType(&stage.PartNames, &stage.PartNames_mapString, &stage.PartName_stagedOrder, &stage.PartNameOrder)

	__gong__resetStageType(&stage.Players, &stage.Players_mapString, &stage.Player_stagedOrder, &stage.PlayerOrder)

	__gong__resetStageType(&stage.ScoreInstruments, &stage.ScoreInstruments_mapString, &stage.ScoreInstrument_stagedOrder, &stage.ScoreInstrumentOrder)

	__gong__resetStageType(&stage.ScoreParts, &stage.ScoreParts_mapString, &stage.ScorePart_stagedOrder, &stage.ScorePartOrder)

	__gong__resetStageType(&stage.ScorePartwises, &stage.ScorePartwises_mapString, &stage.ScorePartwise_stagedOrder, &stage.ScorePartwiseOrder)

	__gong__resetStageType(&stage.ScoreTimewises, &stage.ScoreTimewises_mapString, &stage.ScoreTimewise_stagedOrder, &stage.ScoreTimewiseOrder)

	__gong__resetStageType(&stage.ScoreTimewiseMeasures, &stage.ScoreTimewiseMeasures_mapString, &stage.ScoreTimewiseMeasure_stagedOrder, &stage.ScoreTimewiseMeasureOrder)

	__gong__resetStageType(&stage.ScoreTimewiseMeasureParts, &stage.ScoreTimewiseMeasureParts_mapString, &stage.ScoreTimewiseMeasurePart_stagedOrder, &stage.ScoreTimewiseMeasurePartOrder)

	__gong__resetStageType(&stage.Works, &stage.Works_mapString, &stage.Work_stagedOrder, &stage.WorkOrder)

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
func (Credit) GongGetAssociationName() any {
	return &Credit{
	}
}

func (Defaults) GongGetAssociationName() any {
	return &Defaults{
			MusicFont: &EmptyFont{Name: "MusicFont"},
			WordFont: &EmptyFont{Name: "WordFont"},
			LyricFont: []*LyricFont{{Name: "LyricFont"}},
			LyricLanguage: []*LyricLanguage{{Name: "LyricLanguage"}},
	}
}

func (EmptyFont) GongGetAssociationName() any {
	return &EmptyFont{
	}
}

func (GroupBarline) GongGetAssociationName() any {
	return &GroupBarline{
	}
}

func (GroupName) GongGetAssociationName() any {
	return &GroupName{
	}
}

func (GroupSymbol) GongGetAssociationName() any {
	return &GroupSymbol{
	}
}

func (InstrumentLink) GongGetAssociationName() any {
	return &InstrumentLink{
			RefScoreInstrument: &ScoreInstrument{Name: "RefScoreInstrument"},
	}
}

func (LyricFont) GongGetAssociationName() any {
	return &LyricFont{
	}
}

func (LyricLanguage) GongGetAssociationName() any {
	return &LyricLanguage{
	}
}

func (Measure) GongGetAssociationName() any {
	return &Measure{
			MusicData: []*MeasureMusicDataChoice{{Name: "MusicData"}},
	}
}

func (MeasureMusicDataChoice) GongGetAssociationName() any {
	return &MeasureMusicDataChoice{
	}
}

func (Opus) GongGetAssociationName() any {
	return &Opus{
	}
}

func (Part) GongGetAssociationName() any {
	return &Part{
			RefScorePart: &ScorePart{Name: "RefScorePart"},
			Measure: []*Measure{{Name: "Measure"}},
	}
}

func (PartGroup) GongGetAssociationName() any {
	return &PartGroup{
			GroupName: &GroupName{Name: "GroupName"},
			GroupAbbreviation: &GroupName{Name: "GroupAbbreviation"},
			GroupSymbol: &GroupSymbol{Name: "GroupSymbol"},
			GroupBarline: &GroupBarline{Name: "GroupBarline"},
	}
}

func (PartLink) GongGetAssociationName() any {
	return &PartLink{
			InstrumentLink: []*InstrumentLink{{Name: "InstrumentLink"}},
	}
}

func (PartList) GongGetAssociationName() any {
	return &PartList{
			PartGroup: []*PartGroup{{Name: "PartGroup"}},
			ScorePart: []*ScorePart{{Name: "ScorePart"}},
	}
}

func (PartName) GongGetAssociationName() any {
	return &PartName{
	}
}

func (Player) GongGetAssociationName() any {
	return &Player{
	}
}

func (ScoreInstrument) GongGetAssociationName() any {
	return &ScoreInstrument{
	}
}

func (ScorePart) GongGetAssociationName() any {
	return &ScorePart{
			PartLink: []*PartLink{{Name: "PartLink"}},
			PartName: &PartName{Name: "PartName"},
			PartAbbreviation: &PartName{Name: "PartAbbreviation"},
			ScoreInstrument: []*ScoreInstrument{{Name: "ScoreInstrument"}},
			Player: []*Player{{Name: "Player"}},
	}
}

func (ScorePartwise) GongGetAssociationName() any {
	return &ScorePartwise{
			Work: &Work{Name: "Work"},
			Defaults: &Defaults{Name: "Defaults"},
			Credit: []*Credit{{Name: "Credit"}},
			PartList: &PartList{Name: "PartList"},
			Part: []*Part{{Name: "Part"}},
	}
}

func (ScoreTimewise) GongGetAssociationName() any {
	return &ScoreTimewise{
			Work: &Work{Name: "Work"},
			Defaults: &Defaults{Name: "Defaults"},
			Credit: []*Credit{{Name: "Credit"}},
			PartList: &PartList{Name: "PartList"},
			Measure: []*ScoreTimewiseMeasure{{Name: "Measure"}},
	}
}

func (ScoreTimewiseMeasure) GongGetAssociationName() any {
	return &ScoreTimewiseMeasure{
			Part: []*ScoreTimewiseMeasurePart{{Name: "Part"}},
	}
}

func (ScoreTimewiseMeasurePart) GongGetAssociationName() any {
	return &ScoreTimewiseMeasurePart{
			RefScorePart: &ScorePart{Name: "RefScorePart"},
	}
}

func (Work) GongGetAssociationName() any {
	return &Work{
			Opus: &Opus{Name: "Opus"},
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
	// reverse maps of direct associations of Credit
	case Credit:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Defaults
	case Defaults:
		switch fieldname {
		// insertion point for per direct association field
		case "MusicFont":
			res := make(map[*EmptyFont][]*Defaults)
			for defaults := range stage.Defaultss {
				if defaults.MusicFont != nil {
					emptyfont_ := defaults.MusicFont
					var defaultss []*Defaults
					_, ok := res[emptyfont_]
					if ok {
						defaultss = res[emptyfont_]
					} else {
						defaultss = make([]*Defaults, 0)
					}
					defaultss = append(defaultss, defaults)
					res[emptyfont_] = defaultss
				}
			}
			return any(res).(map[*End][]*Start)
		case "WordFont":
			res := make(map[*EmptyFont][]*Defaults)
			for defaults := range stage.Defaultss {
				if defaults.WordFont != nil {
					emptyfont_ := defaults.WordFont
					var defaultss []*Defaults
					_, ok := res[emptyfont_]
					if ok {
						defaultss = res[emptyfont_]
					} else {
						defaultss = make([]*Defaults, 0)
					}
					defaultss = append(defaultss, defaults)
					res[emptyfont_] = defaultss
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of EmptyFont
	case EmptyFont:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GroupBarline
	case GroupBarline:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GroupName
	case GroupName:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GroupSymbol
	case GroupSymbol:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of InstrumentLink
	case InstrumentLink:
		switch fieldname {
		// insertion point for per direct association field
		case "RefScoreInstrument":
			res := make(map[*ScoreInstrument][]*InstrumentLink)
			for instrumentlink := range stage.InstrumentLinks {
				if instrumentlink.RefScoreInstrument != nil {
					scoreinstrument_ := instrumentlink.RefScoreInstrument
					var instrumentlinks []*InstrumentLink
					_, ok := res[scoreinstrument_]
					if ok {
						instrumentlinks = res[scoreinstrument_]
					} else {
						instrumentlinks = make([]*InstrumentLink, 0)
					}
					instrumentlinks = append(instrumentlinks, instrumentlink)
					res[scoreinstrument_] = instrumentlinks
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of LyricFont
	case LyricFont:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LyricLanguage
	case LyricLanguage:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Measure
	case Measure:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MeasureMusicDataChoice
	case MeasureMusicDataChoice:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Opus
	case Opus:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Part
	case Part:
		switch fieldname {
		// insertion point for per direct association field
		case "RefScorePart":
			res := make(map[*ScorePart][]*Part)
			for part := range stage.Parts {
				if part.RefScorePart != nil {
					scorepart_ := part.RefScorePart
					var parts []*Part
					_, ok := res[scorepart_]
					if ok {
						parts = res[scorepart_]
					} else {
						parts = make([]*Part, 0)
					}
					parts = append(parts, part)
					res[scorepart_] = parts
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of PartGroup
	case PartGroup:
		switch fieldname {
		// insertion point for per direct association field
		case "GroupName":
			res := make(map[*GroupName][]*PartGroup)
			for partgroup := range stage.PartGroups {
				if partgroup.GroupName != nil {
					groupname_ := partgroup.GroupName
					var partgroups []*PartGroup
					_, ok := res[groupname_]
					if ok {
						partgroups = res[groupname_]
					} else {
						partgroups = make([]*PartGroup, 0)
					}
					partgroups = append(partgroups, partgroup)
					res[groupname_] = partgroups
				}
			}
			return any(res).(map[*End][]*Start)
		case "GroupAbbreviation":
			res := make(map[*GroupName][]*PartGroup)
			for partgroup := range stage.PartGroups {
				if partgroup.GroupAbbreviation != nil {
					groupname_ := partgroup.GroupAbbreviation
					var partgroups []*PartGroup
					_, ok := res[groupname_]
					if ok {
						partgroups = res[groupname_]
					} else {
						partgroups = make([]*PartGroup, 0)
					}
					partgroups = append(partgroups, partgroup)
					res[groupname_] = partgroups
				}
			}
			return any(res).(map[*End][]*Start)
		case "GroupSymbol":
			res := make(map[*GroupSymbol][]*PartGroup)
			for partgroup := range stage.PartGroups {
				if partgroup.GroupSymbol != nil {
					groupsymbol_ := partgroup.GroupSymbol
					var partgroups []*PartGroup
					_, ok := res[groupsymbol_]
					if ok {
						partgroups = res[groupsymbol_]
					} else {
						partgroups = make([]*PartGroup, 0)
					}
					partgroups = append(partgroups, partgroup)
					res[groupsymbol_] = partgroups
				}
			}
			return any(res).(map[*End][]*Start)
		case "GroupBarline":
			res := make(map[*GroupBarline][]*PartGroup)
			for partgroup := range stage.PartGroups {
				if partgroup.GroupBarline != nil {
					groupbarline_ := partgroup.GroupBarline
					var partgroups []*PartGroup
					_, ok := res[groupbarline_]
					if ok {
						partgroups = res[groupbarline_]
					} else {
						partgroups = make([]*PartGroup, 0)
					}
					partgroups = append(partgroups, partgroup)
					res[groupbarline_] = partgroups
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of PartLink
	case PartLink:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PartList
	case PartList:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PartName
	case PartName:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Player
	case Player:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ScoreInstrument
	case ScoreInstrument:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ScorePart
	case ScorePart:
		switch fieldname {
		// insertion point for per direct association field
		case "PartName":
			res := make(map[*PartName][]*ScorePart)
			for scorepart := range stage.ScoreParts {
				if scorepart.PartName != nil {
					partname_ := scorepart.PartName
					var scoreparts []*ScorePart
					_, ok := res[partname_]
					if ok {
						scoreparts = res[partname_]
					} else {
						scoreparts = make([]*ScorePart, 0)
					}
					scoreparts = append(scoreparts, scorepart)
					res[partname_] = scoreparts
				}
			}
			return any(res).(map[*End][]*Start)
		case "PartAbbreviation":
			res := make(map[*PartName][]*ScorePart)
			for scorepart := range stage.ScoreParts {
				if scorepart.PartAbbreviation != nil {
					partname_ := scorepart.PartAbbreviation
					var scoreparts []*ScorePart
					_, ok := res[partname_]
					if ok {
						scoreparts = res[partname_]
					} else {
						scoreparts = make([]*ScorePart, 0)
					}
					scoreparts = append(scoreparts, scorepart)
					res[partname_] = scoreparts
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ScorePartwise
	case ScorePartwise:
		switch fieldname {
		// insertion point for per direct association field
		case "Work":
			res := make(map[*Work][]*ScorePartwise)
			for scorepartwise := range stage.ScorePartwises {
				if scorepartwise.Work != nil {
					work_ := scorepartwise.Work
					var scorepartwises []*ScorePartwise
					_, ok := res[work_]
					if ok {
						scorepartwises = res[work_]
					} else {
						scorepartwises = make([]*ScorePartwise, 0)
					}
					scorepartwises = append(scorepartwises, scorepartwise)
					res[work_] = scorepartwises
				}
			}
			return any(res).(map[*End][]*Start)
		case "Defaults":
			res := make(map[*Defaults][]*ScorePartwise)
			for scorepartwise := range stage.ScorePartwises {
				if scorepartwise.Defaults != nil {
					defaults_ := scorepartwise.Defaults
					var scorepartwises []*ScorePartwise
					_, ok := res[defaults_]
					if ok {
						scorepartwises = res[defaults_]
					} else {
						scorepartwises = make([]*ScorePartwise, 0)
					}
					scorepartwises = append(scorepartwises, scorepartwise)
					res[defaults_] = scorepartwises
				}
			}
			return any(res).(map[*End][]*Start)
		case "PartList":
			res := make(map[*PartList][]*ScorePartwise)
			for scorepartwise := range stage.ScorePartwises {
				if scorepartwise.PartList != nil {
					partlist_ := scorepartwise.PartList
					var scorepartwises []*ScorePartwise
					_, ok := res[partlist_]
					if ok {
						scorepartwises = res[partlist_]
					} else {
						scorepartwises = make([]*ScorePartwise, 0)
					}
					scorepartwises = append(scorepartwises, scorepartwise)
					res[partlist_] = scorepartwises
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ScoreTimewise
	case ScoreTimewise:
		switch fieldname {
		// insertion point for per direct association field
		case "Work":
			res := make(map[*Work][]*ScoreTimewise)
			for scoretimewise := range stage.ScoreTimewises {
				if scoretimewise.Work != nil {
					work_ := scoretimewise.Work
					var scoretimewises []*ScoreTimewise
					_, ok := res[work_]
					if ok {
						scoretimewises = res[work_]
					} else {
						scoretimewises = make([]*ScoreTimewise, 0)
					}
					scoretimewises = append(scoretimewises, scoretimewise)
					res[work_] = scoretimewises
				}
			}
			return any(res).(map[*End][]*Start)
		case "Defaults":
			res := make(map[*Defaults][]*ScoreTimewise)
			for scoretimewise := range stage.ScoreTimewises {
				if scoretimewise.Defaults != nil {
					defaults_ := scoretimewise.Defaults
					var scoretimewises []*ScoreTimewise
					_, ok := res[defaults_]
					if ok {
						scoretimewises = res[defaults_]
					} else {
						scoretimewises = make([]*ScoreTimewise, 0)
					}
					scoretimewises = append(scoretimewises, scoretimewise)
					res[defaults_] = scoretimewises
				}
			}
			return any(res).(map[*End][]*Start)
		case "PartList":
			res := make(map[*PartList][]*ScoreTimewise)
			for scoretimewise := range stage.ScoreTimewises {
				if scoretimewise.PartList != nil {
					partlist_ := scoretimewise.PartList
					var scoretimewises []*ScoreTimewise
					_, ok := res[partlist_]
					if ok {
						scoretimewises = res[partlist_]
					} else {
						scoretimewises = make([]*ScoreTimewise, 0)
					}
					scoretimewises = append(scoretimewises, scoretimewise)
					res[partlist_] = scoretimewises
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ScoreTimewiseMeasure
	case ScoreTimewiseMeasure:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ScoreTimewiseMeasurePart
	case ScoreTimewiseMeasurePart:
		switch fieldname {
		// insertion point for per direct association field
		case "RefScorePart":
			res := make(map[*ScorePart][]*ScoreTimewiseMeasurePart)
			for scoretimewisemeasurepart := range stage.ScoreTimewiseMeasureParts {
				if scoretimewisemeasurepart.RefScorePart != nil {
					scorepart_ := scoretimewisemeasurepart.RefScorePart
					var scoretimewisemeasureparts []*ScoreTimewiseMeasurePart
					_, ok := res[scorepart_]
					if ok {
						scoretimewisemeasureparts = res[scorepart_]
					} else {
						scoretimewisemeasureparts = make([]*ScoreTimewiseMeasurePart, 0)
					}
					scoretimewisemeasureparts = append(scoretimewisemeasureparts, scoretimewisemeasurepart)
					res[scorepart_] = scoretimewisemeasureparts
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Work
	case Work:
		switch fieldname {
		// insertion point for per direct association field
		case "Opus":
			res := make(map[*Opus][]*Work)
			for work := range stage.Works {
				if work.Opus != nil {
					opus_ := work.Opus
					var works []*Work
					_, ok := res[opus_]
					if ok {
						works = res[opus_]
					} else {
						works = make([]*Work, 0)
					}
					works = append(works, work)
					res[opus_] = works
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
	// reverse maps of direct associations of Credit
	case Credit:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Defaults
	case Defaults:
		switch fieldname {
		// insertion point for per direct association field
		case "LyricFont":
			res := make(map[*LyricFont][]*Defaults)
			for defaults := range stage.Defaultss {
				for _, lyricfont_ := range defaults.LyricFont {
					res[lyricfont_] = append(res[lyricfont_], defaults)
				}
			}
			return any(res).(map[*End][]*Start)
		case "LyricLanguage":
			res := make(map[*LyricLanguage][]*Defaults)
			for defaults := range stage.Defaultss {
				for _, lyriclanguage_ := range defaults.LyricLanguage {
					res[lyriclanguage_] = append(res[lyriclanguage_], defaults)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of EmptyFont
	case EmptyFont:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GroupBarline
	case GroupBarline:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GroupName
	case GroupName:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GroupSymbol
	case GroupSymbol:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of InstrumentLink
	case InstrumentLink:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LyricFont
	case LyricFont:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LyricLanguage
	case LyricLanguage:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Measure
	case Measure:
		switch fieldname {
		// insertion point for per direct association field
		case "MusicData":
			res := make(map[*MeasureMusicDataChoice][]*Measure)
			for measure := range stage.Measures {
				for _, measuremusicdatachoice_ := range measure.MusicData {
					res[measuremusicdatachoice_] = append(res[measuremusicdatachoice_], measure)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MeasureMusicDataChoice
	case MeasureMusicDataChoice:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Opus
	case Opus:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Part
	case Part:
		switch fieldname {
		// insertion point for per direct association field
		case "Measure":
			res := make(map[*Measure][]*Part)
			for part := range stage.Parts {
				for _, measure_ := range part.Measure {
					res[measure_] = append(res[measure_], part)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of PartGroup
	case PartGroup:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PartLink
	case PartLink:
		switch fieldname {
		// insertion point for per direct association field
		case "InstrumentLink":
			res := make(map[*InstrumentLink][]*PartLink)
			for partlink := range stage.PartLinks {
				for _, instrumentlink_ := range partlink.InstrumentLink {
					res[instrumentlink_] = append(res[instrumentlink_], partlink)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of PartList
	case PartList:
		switch fieldname {
		// insertion point for per direct association field
		case "PartGroup":
			res := make(map[*PartGroup][]*PartList)
			for partlist := range stage.PartLists {
				for _, partgroup_ := range partlist.PartGroup {
					res[partgroup_] = append(res[partgroup_], partlist)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ScorePart":
			res := make(map[*ScorePart][]*PartList)
			for partlist := range stage.PartLists {
				for _, scorepart_ := range partlist.ScorePart {
					res[scorepart_] = append(res[scorepart_], partlist)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of PartName
	case PartName:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Player
	case Player:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ScoreInstrument
	case ScoreInstrument:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ScorePart
	case ScorePart:
		switch fieldname {
		// insertion point for per direct association field
		case "PartLink":
			res := make(map[*PartLink][]*ScorePart)
			for scorepart := range stage.ScoreParts {
				for _, partlink_ := range scorepart.PartLink {
					res[partlink_] = append(res[partlink_], scorepart)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ScoreInstrument":
			res := make(map[*ScoreInstrument][]*ScorePart)
			for scorepart := range stage.ScoreParts {
				for _, scoreinstrument_ := range scorepart.ScoreInstrument {
					res[scoreinstrument_] = append(res[scoreinstrument_], scorepart)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Player":
			res := make(map[*Player][]*ScorePart)
			for scorepart := range stage.ScoreParts {
				for _, player_ := range scorepart.Player {
					res[player_] = append(res[player_], scorepart)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ScorePartwise
	case ScorePartwise:
		switch fieldname {
		// insertion point for per direct association field
		case "Credit":
			res := make(map[*Credit][]*ScorePartwise)
			for scorepartwise := range stage.ScorePartwises {
				for _, credit_ := range scorepartwise.Credit {
					res[credit_] = append(res[credit_], scorepartwise)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Part":
			res := make(map[*Part][]*ScorePartwise)
			for scorepartwise := range stage.ScorePartwises {
				for _, part_ := range scorepartwise.Part {
					res[part_] = append(res[part_], scorepartwise)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ScoreTimewise
	case ScoreTimewise:
		switch fieldname {
		// insertion point for per direct association field
		case "Credit":
			res := make(map[*Credit][]*ScoreTimewise)
			for scoretimewise := range stage.ScoreTimewises {
				for _, credit_ := range scoretimewise.Credit {
					res[credit_] = append(res[credit_], scoretimewise)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Measure":
			res := make(map[*ScoreTimewiseMeasure][]*ScoreTimewise)
			for scoretimewise := range stage.ScoreTimewises {
				for _, scoretimewisemeasure_ := range scoretimewise.Measure {
					res[scoretimewisemeasure_] = append(res[scoretimewisemeasure_], scoretimewise)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ScoreTimewiseMeasure
	case ScoreTimewiseMeasure:
		switch fieldname {
		// insertion point for per direct association field
		case "Part":
			res := make(map[*ScoreTimewiseMeasurePart][]*ScoreTimewiseMeasure)
			for scoretimewisemeasure := range stage.ScoreTimewiseMeasures {
				for _, scoretimewisemeasurepart_ := range scoretimewisemeasure.Part {
					res[scoretimewisemeasurepart_] = append(res[scoretimewisemeasurepart_], scoretimewisemeasure)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ScoreTimewiseMeasurePart
	case ScoreTimewiseMeasurePart:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Work
	case Work:
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
func (*Credit) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "ScorePartwise",
			Fieldname: "Credit",
		},
		{
			GongstructName: "ScoreTimewise",
			Fieldname: "Credit",
		},
	}
}

func (*Defaults) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*EmptyFont) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*GroupBarline) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*GroupName) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*GroupSymbol) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*InstrumentLink) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "PartLink",
			Fieldname: "InstrumentLink",
		},
	}
}

func (*LyricFont) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Defaults",
			Fieldname: "LyricFont",
		},
	}
}

func (*LyricLanguage) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Defaults",
			Fieldname: "LyricLanguage",
		},
	}
}

func (*Measure) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Part",
			Fieldname: "Measure",
		},
	}
}

func (*MeasureMusicDataChoice) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Measure",
			Fieldname: "MusicData",
		},
	}
}

func (*Opus) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Part) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "ScorePartwise",
			Fieldname: "Part",
		},
	}
}

func (*PartGroup) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "PartList",
			Fieldname: "PartGroup",
		},
	}
}

func (*PartLink) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "ScorePart",
			Fieldname: "PartLink",
		},
	}
}

func (*PartList) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*PartName) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Player) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "ScorePart",
			Fieldname: "Player",
		},
	}
}

func (*ScoreInstrument) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "ScorePart",
			Fieldname: "ScoreInstrument",
		},
	}
}

func (*ScorePart) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "PartList",
			Fieldname: "ScorePart",
		},
	}
}

func (*ScorePartwise) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ScoreTimewise) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ScoreTimewiseMeasure) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "ScoreTimewise",
			Fieldname: "Measure",
		},
	}
}

func (*ScoreTimewiseMeasurePart) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "ScoreTimewiseMeasure",
			Fieldname: "Part",
		},
	}
}

func (*Work) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
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
func (credit *Credit) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Page",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "CreditType",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (defaults *Defaults) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "MusicFont",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "EmptyFont",
		},
		{
			Name:                 "WordFont",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "EmptyFont",
		},
		{
			Name:                 "LyricFont",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "LyricFont",
		},
		{
			Name:                 "LyricLanguage",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "LyricLanguage",
		},
	}
	return
}

func (emptyfont *EmptyFont) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
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
	}
	return
}

func (groupbarline *GroupBarline) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumGroupBarlineValue",
		},
	}
	return
}

func (groupname *GroupName) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
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

func (groupsymbol *GroupSymbol) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
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
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (instrumentlink *InstrumentLink) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "RefScoreInstrument",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ScoreInstrument",
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (lyricfont *LyricFont) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NameAttr",
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
	}
	return
}

func (lyriclanguage *LyricLanguage) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NameAttr",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Lang",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (measure *Measure) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Text",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "MusicData",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MeasureMusicDataChoice",
		},
	}
	return
}

func (measuremusicdatachoice *MeasureMusicDataChoice) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (opus *Opus) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Href",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Role",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Title",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (part *Part) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "RefScorePart",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ScorePart",
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Measure",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Measure",
		},
	}
	return
}

func (partgroup *PartGroup) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "GroupName",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "GroupName",
		},
		{
			Name:                 "GroupAbbreviation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "GroupName",
		},
		{
			Name:                 "GroupSymbol",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "GroupSymbol",
		},
		{
			Name:                 "GroupBarline",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "GroupBarline",
		},
	}
	return
}

func (partlink *PartLink) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Href",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Role",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Title",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "InstrumentLink",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "InstrumentLink",
		},
		{
			Name:               "GroupLink",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (partlist *PartList) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "PartGroup",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "PartGroup",
		},
		{
			Name:                 "ScorePart",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ScorePart",
		},
	}
	return
}

func (partname *PartName) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
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

func (player *Player) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "PlayerName",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (scoreinstrument *ScoreInstrument) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "InstrumentName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "InstrumentAbbreviation",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "InstrumentSound",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Ensemble",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (scorepart *ScorePart) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "PartLink",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "PartLink",
		},
		{
			Name:                 "PartName",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PartName",
		},
		{
			Name:                 "PartAbbreviation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PartName",
		},
		{
			Name:               "Group",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "ScoreInstrument",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ScoreInstrument",
		},
		{
			Name:                 "Player",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Player",
		},
	}
	return
}

func (scorepartwise *ScorePartwise) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Version",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Work",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Work",
		},
		{
			Name:               "MovementNumber",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MovementTitle",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Defaults",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Defaults",
		},
		{
			Name:                 "Credit",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Credit",
		},
		{
			Name:                 "PartList",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PartList",
		},
		{
			Name:                 "Part",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Part",
		},
	}
	return
}

func (scoretimewise *ScoreTimewise) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Version",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Work",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Work",
		},
		{
			Name:               "MovementNumber",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MovementTitle",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Defaults",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Defaults",
		},
		{
			Name:                 "Credit",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Credit",
		},
		{
			Name:                 "PartList",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PartList",
		},
		{
			Name:                 "Measure",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ScoreTimewiseMeasure",
		},
	}
	return
}

func (scoretimewisemeasure *ScoreTimewiseMeasure) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Text",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Part",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ScoreTimewiseMeasurePart",
		},
	}
	return
}

func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "RefScorePart",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ScorePart",
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (work *Work) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "WorkNumber",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "WorkTitle",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Opus",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Opus",
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
func (credit *Credit) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = credit.Name
	case "Page":
		res.valueString = fmt.Sprintf("%d", credit.Page)
		res.valueInt = credit.Page
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Id":
		res.valueString = credit.Id
	case "CreditType":
		res.valueString = credit.CreditType
	}
	return
}

func (defaults *Defaults) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = defaults.Name
	case "MusicFont":
		res.GongFieldValueType = GongFieldValueTypePointer
		if defaults.MusicFont != nil {
			res.valueString = defaults.MusicFont.Name
			res.ids = defaults.MusicFont.GongGetUUID(stage)
		}
	case "WordFont":
		res.GongFieldValueType = GongFieldValueTypePointer
		if defaults.WordFont != nil {
			res.valueString = defaults.WordFont.Name
			res.ids = defaults.WordFont.GongGetUUID(stage)
		}
	case "LyricFont":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range defaults.LyricFont {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "LyricLanguage":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range defaults.LyricLanguage {
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

func (emptyfont *EmptyFont) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = emptyfont.Name
	case "FontFamily":
		res.valueString = emptyfont.FontFamily
	case "FontSize":
		res.valueString = emptyfont.FontSize
	}
	return
}

func (groupbarline *GroupBarline) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = groupbarline.Name
	case "Color":
		res.valueString = groupbarline.Color
	case "Value":
		enum := groupbarline.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (groupname *GroupName) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = groupname.Name
	case "DefaultX":
		res.valueString = groupname.DefaultX
	case "DefaultY":
		res.valueString = groupname.DefaultY
	case "RelativeX":
		res.valueString = groupname.RelativeX
	case "RelativeY":
		res.valueString = groupname.RelativeY
	case "FontFamily":
		res.valueString = groupname.FontFamily
	case "FontSize":
		res.valueString = groupname.FontSize
	case "Color":
		res.valueString = groupname.Color
	case "Value":
		res.valueString = groupname.Value
	}
	return
}

func (groupsymbol *GroupSymbol) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = groupsymbol.Name
	case "DefaultX":
		res.valueString = groupsymbol.DefaultX
	case "DefaultY":
		res.valueString = groupsymbol.DefaultY
	case "RelativeX":
		res.valueString = groupsymbol.RelativeX
	case "RelativeY":
		res.valueString = groupsymbol.RelativeY
	case "Color":
		res.valueString = groupsymbol.Color
	}
	return
}

func (instrumentlink *InstrumentLink) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = instrumentlink.Name
	case "RefScoreInstrument":
		res.GongFieldValueType = GongFieldValueTypePointer
		if instrumentlink.RefScoreInstrument != nil {
			res.valueString = instrumentlink.RefScoreInstrument.Name
			res.ids = instrumentlink.RefScoreInstrument.GongGetUUID(stage)
		}
	case "Id":
		res.valueString = instrumentlink.Id
	}
	return
}

func (lyricfont *LyricFont) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = lyricfont.Name
	case "Number":
		res.valueString = lyricfont.Number
	case "NameAttr":
		res.valueString = lyricfont.NameAttr
	case "FontFamily":
		res.valueString = lyricfont.FontFamily
	case "FontSize":
		res.valueString = lyricfont.FontSize
	}
	return
}

func (lyriclanguage *LyricLanguage) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = lyriclanguage.Name
	case "Number":
		res.valueString = lyriclanguage.Number
	case "NameAttr":
		res.valueString = lyriclanguage.NameAttr
	case "Lang":
		res.valueString = lyriclanguage.Lang
	}
	return
}

func (measure *Measure) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = measure.Name
	case "Number":
		res.valueString = measure.Number
	case "Text":
		res.valueString = measure.Text
	case "Width":
		res.valueString = measure.Width
	case "Id":
		res.valueString = measure.Id
	case "MusicData":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range measure.MusicData {
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

func (measuremusicdatachoice *MeasureMusicDataChoice) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = measuremusicdatachoice.Name
	}
	return
}

func (opus *Opus) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = opus.Name
	case "Href":
		res.valueString = opus.Href
	case "Role":
		res.valueString = opus.Role
	case "Title":
		res.valueString = opus.Title
	}
	return
}

func (part *Part) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = part.Name
	case "RefScorePart":
		res.GongFieldValueType = GongFieldValueTypePointer
		if part.RefScorePart != nil {
			res.valueString = part.RefScorePart.Name
			res.ids = part.RefScorePart.GongGetUUID(stage)
		}
	case "Id":
		res.valueString = part.Id
	case "Measure":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range part.Measure {
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

func (partgroup *PartGroup) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = partgroup.Name
	case "Number":
		res.valueString = partgroup.Number
	case "GroupName":
		res.GongFieldValueType = GongFieldValueTypePointer
		if partgroup.GroupName != nil {
			res.valueString = partgroup.GroupName.Name
			res.ids = partgroup.GroupName.GongGetUUID(stage)
		}
	case "GroupAbbreviation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if partgroup.GroupAbbreviation != nil {
			res.valueString = partgroup.GroupAbbreviation.Name
			res.ids = partgroup.GroupAbbreviation.GongGetUUID(stage)
		}
	case "GroupSymbol":
		res.GongFieldValueType = GongFieldValueTypePointer
		if partgroup.GroupSymbol != nil {
			res.valueString = partgroup.GroupSymbol.Name
			res.ids = partgroup.GroupSymbol.GongGetUUID(stage)
		}
	case "GroupBarline":
		res.GongFieldValueType = GongFieldValueTypePointer
		if partgroup.GroupBarline != nil {
			res.valueString = partgroup.GroupBarline.Name
			res.ids = partgroup.GroupBarline.GongGetUUID(stage)
		}
	}
	return
}

func (partlink *PartLink) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = partlink.Name
	case "Href":
		res.valueString = partlink.Href
	case "Role":
		res.valueString = partlink.Role
	case "Title":
		res.valueString = partlink.Title
	case "InstrumentLink":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range partlink.InstrumentLink {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "GroupLink":
		res.valueString = partlink.GroupLink
	}
	return
}

func (partlist *PartList) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = partlist.Name
	case "PartGroup":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range partlist.PartGroup {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ScorePart":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range partlist.ScorePart {
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

func (partname *PartName) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = partname.Name
	case "DefaultX":
		res.valueString = partname.DefaultX
	case "DefaultY":
		res.valueString = partname.DefaultY
	case "RelativeX":
		res.valueString = partname.RelativeX
	case "RelativeY":
		res.valueString = partname.RelativeY
	case "FontFamily":
		res.valueString = partname.FontFamily
	case "FontSize":
		res.valueString = partname.FontSize
	case "Color":
		res.valueString = partname.Color
	case "Value":
		res.valueString = partname.Value
	}
	return
}

func (player *Player) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = player.Name
	case "Id":
		res.valueString = player.Id
	case "PlayerName":
		res.valueString = player.PlayerName
	}
	return
}

func (scoreinstrument *ScoreInstrument) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = scoreinstrument.Name
	case "Id":
		res.valueString = scoreinstrument.Id
	case "InstrumentName":
		res.valueString = scoreinstrument.InstrumentName
	case "InstrumentAbbreviation":
		res.valueString = scoreinstrument.InstrumentAbbreviation
	case "InstrumentSound":
		res.valueString = scoreinstrument.InstrumentSound
	case "Ensemble":
		res.valueString = scoreinstrument.Ensemble
	}
	return
}

func (scorepart *ScorePart) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = scorepart.Name
	case "Id":
		res.valueString = scorepart.Id
	case "PartLink":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range scorepart.PartLink {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PartName":
		res.GongFieldValueType = GongFieldValueTypePointer
		if scorepart.PartName != nil {
			res.valueString = scorepart.PartName.Name
			res.ids = scorepart.PartName.GongGetUUID(stage)
		}
	case "PartAbbreviation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if scorepart.PartAbbreviation != nil {
			res.valueString = scorepart.PartAbbreviation.Name
			res.ids = scorepart.PartAbbreviation.GongGetUUID(stage)
		}
	case "Group":
		res.valueString = scorepart.Group
	case "ScoreInstrument":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range scorepart.ScoreInstrument {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Player":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range scorepart.Player {
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

func (scorepartwise *ScorePartwise) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = scorepartwise.Name
	case "Version":
		res.valueString = scorepartwise.Version
	case "Work":
		res.GongFieldValueType = GongFieldValueTypePointer
		if scorepartwise.Work != nil {
			res.valueString = scorepartwise.Work.Name
			res.ids = scorepartwise.Work.GongGetUUID(stage)
		}
	case "MovementNumber":
		res.valueString = scorepartwise.MovementNumber
	case "MovementTitle":
		res.valueString = scorepartwise.MovementTitle
	case "Defaults":
		res.GongFieldValueType = GongFieldValueTypePointer
		if scorepartwise.Defaults != nil {
			res.valueString = scorepartwise.Defaults.Name
			res.ids = scorepartwise.Defaults.GongGetUUID(stage)
		}
	case "Credit":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range scorepartwise.Credit {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PartList":
		res.GongFieldValueType = GongFieldValueTypePointer
		if scorepartwise.PartList != nil {
			res.valueString = scorepartwise.PartList.Name
			res.ids = scorepartwise.PartList.GongGetUUID(stage)
		}
	case "Part":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range scorepartwise.Part {
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

func (scoretimewise *ScoreTimewise) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = scoretimewise.Name
	case "Version":
		res.valueString = scoretimewise.Version
	case "Work":
		res.GongFieldValueType = GongFieldValueTypePointer
		if scoretimewise.Work != nil {
			res.valueString = scoretimewise.Work.Name
			res.ids = scoretimewise.Work.GongGetUUID(stage)
		}
	case "MovementNumber":
		res.valueString = scoretimewise.MovementNumber
	case "MovementTitle":
		res.valueString = scoretimewise.MovementTitle
	case "Defaults":
		res.GongFieldValueType = GongFieldValueTypePointer
		if scoretimewise.Defaults != nil {
			res.valueString = scoretimewise.Defaults.Name
			res.ids = scoretimewise.Defaults.GongGetUUID(stage)
		}
	case "Credit":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range scoretimewise.Credit {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PartList":
		res.GongFieldValueType = GongFieldValueTypePointer
		if scoretimewise.PartList != nil {
			res.valueString = scoretimewise.PartList.Name
			res.ids = scoretimewise.PartList.GongGetUUID(stage)
		}
	case "Measure":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range scoretimewise.Measure {
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

func (scoretimewisemeasure *ScoreTimewiseMeasure) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = scoretimewisemeasure.Name
	case "Number":
		res.valueString = scoretimewisemeasure.Number
	case "Text":
		res.valueString = scoretimewisemeasure.Text
	case "Width":
		res.valueString = scoretimewisemeasure.Width
	case "Id":
		res.valueString = scoretimewisemeasure.Id
	case "Part":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range scoretimewisemeasure.Part {
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

func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = scoretimewisemeasurepart.Name
	case "RefScorePart":
		res.GongFieldValueType = GongFieldValueTypePointer
		if scoretimewisemeasurepart.RefScorePart != nil {
			res.valueString = scoretimewisemeasurepart.RefScorePart.Name
			res.ids = scoretimewisemeasurepart.RefScorePart.GongGetUUID(stage)
		}
	case "Id":
		res.valueString = scoretimewisemeasurepart.Id
	}
	return
}

func (work *Work) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = work.Name
	case "WorkNumber":
		res.valueString = work.WorkNumber
	case "WorkTitle":
		res.valueString = work.WorkTitle
	case "Opus":
		res.GongFieldValueType = GongFieldValueTypePointer
		if work.Opus != nil {
			res.valueString = work.Opus.Name
			res.ids = work.Opus.GongGetUUID(stage)
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
func (credit *Credit) GongGetGongstructName() string {
	return "Credit"
}

func (defaults *Defaults) GongGetGongstructName() string {
	return "Defaults"
}

func (emptyfont *EmptyFont) GongGetGongstructName() string {
	return "EmptyFont"
}

func (groupbarline *GroupBarline) GongGetGongstructName() string {
	return "GroupBarline"
}

func (groupname *GroupName) GongGetGongstructName() string {
	return "GroupName"
}

func (groupsymbol *GroupSymbol) GongGetGongstructName() string {
	return "GroupSymbol"
}

func (instrumentlink *InstrumentLink) GongGetGongstructName() string {
	return "InstrumentLink"
}

func (lyricfont *LyricFont) GongGetGongstructName() string {
	return "LyricFont"
}

func (lyriclanguage *LyricLanguage) GongGetGongstructName() string {
	return "LyricLanguage"
}

func (measure *Measure) GongGetGongstructName() string {
	return "Measure"
}

func (measuremusicdatachoice *MeasureMusicDataChoice) GongGetGongstructName() string {
	return "MeasureMusicDataChoice"
}

func (opus *Opus) GongGetGongstructName() string {
	return "Opus"
}

func (part *Part) GongGetGongstructName() string {
	return "Part"
}

func (partgroup *PartGroup) GongGetGongstructName() string {
	return "PartGroup"
}

func (partlink *PartLink) GongGetGongstructName() string {
	return "PartLink"
}

func (partlist *PartList) GongGetGongstructName() string {
	return "PartList"
}

func (partname *PartName) GongGetGongstructName() string {
	return "PartName"
}

func (player *Player) GongGetGongstructName() string {
	return "Player"
}

func (scoreinstrument *ScoreInstrument) GongGetGongstructName() string {
	return "ScoreInstrument"
}

func (scorepart *ScorePart) GongGetGongstructName() string {
	return "ScorePart"
}

func (scorepartwise *ScorePartwise) GongGetGongstructName() string {
	return "ScorePartwise"
}

func (scoretimewise *ScoreTimewise) GongGetGongstructName() string {
	return "ScoreTimewise"
}

func (scoretimewisemeasure *ScoreTimewiseMeasure) GongGetGongstructName() string {
	return "ScoreTimewiseMeasure"
}

func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) GongGetGongstructName() string {
	return "ScoreTimewiseMeasurePart"
}

func (work *Work) GongGetGongstructName() string {
	return "Work"
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
	__gong__rebuildMapString(stage.Credits, &stage.Credits_mapString)

	__gong__rebuildMapString(stage.Defaultss, &stage.Defaultss_mapString)

	__gong__rebuildMapString(stage.EmptyFonts, &stage.EmptyFonts_mapString)

	__gong__rebuildMapString(stage.GroupBarlines, &stage.GroupBarlines_mapString)

	__gong__rebuildMapString(stage.GroupNames, &stage.GroupNames_mapString)

	__gong__rebuildMapString(stage.GroupSymbols, &stage.GroupSymbols_mapString)

	__gong__rebuildMapString(stage.InstrumentLinks, &stage.InstrumentLinks_mapString)

	__gong__rebuildMapString(stage.LyricFonts, &stage.LyricFonts_mapString)

	__gong__rebuildMapString(stage.LyricLanguages, &stage.LyricLanguages_mapString)

	__gong__rebuildMapString(stage.Measures, &stage.Measures_mapString)

	__gong__rebuildMapString(stage.MeasureMusicDataChoices, &stage.MeasureMusicDataChoices_mapString)

	__gong__rebuildMapString(stage.Opuss, &stage.Opuss_mapString)

	__gong__rebuildMapString(stage.Parts, &stage.Parts_mapString)

	__gong__rebuildMapString(stage.PartGroups, &stage.PartGroups_mapString)

	__gong__rebuildMapString(stage.PartLinks, &stage.PartLinks_mapString)

	__gong__rebuildMapString(stage.PartLists, &stage.PartLists_mapString)

	__gong__rebuildMapString(stage.PartNames, &stage.PartNames_mapString)

	__gong__rebuildMapString(stage.Players, &stage.Players_mapString)

	__gong__rebuildMapString(stage.ScoreInstruments, &stage.ScoreInstruments_mapString)

	__gong__rebuildMapString(stage.ScoreParts, &stage.ScoreParts_mapString)

	__gong__rebuildMapString(stage.ScorePartwises, &stage.ScorePartwises_mapString)

	__gong__rebuildMapString(stage.ScoreTimewises, &stage.ScoreTimewises_mapString)

	__gong__rebuildMapString(stage.ScoreTimewiseMeasures, &stage.ScoreTimewiseMeasures_mapString)

	__gong__rebuildMapString(stage.ScoreTimewiseMeasureParts, &stage.ScoreTimewiseMeasureParts_mapString)

	__gong__rebuildMapString(stage.Works, &stage.Works_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
