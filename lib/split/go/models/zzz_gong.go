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
	AsSplits                map[*AsSplit]struct{}
	AsSplits_instance       map[*AsSplit]*AsSplit
	AsSplits_mapString      map[string]*AsSplit
	AsSplitOrder            uint
	AsSplit_stagedOrder     map[*AsSplit]uint
	AsSplit_orderStaged     map[uint]*AsSplit
	AsSplits_reference      map[*AsSplit]*AsSplit
	AsSplits_referenceOrder map[*AsSplit]uint

	// insertion point for slice of pointers maps
	AsSplit_AsSplitAreas_reverseMap map[*AsSplitArea]*AsSplit

	OnAfterAsSplitCreateCallback GongOnAfterCreateInterface[AsSplit]
	OnAfterAsSplitUpdateCallback GongOnAfterUpdateInterface[AsSplit]
	OnAfterAsSplitDeleteCallback GongOnAfterDeleteInterface[AsSplit]

	AsSplitAreas                map[*AsSplitArea]struct{}
	AsSplitAreas_instance       map[*AsSplitArea]*AsSplitArea
	AsSplitAreas_mapString      map[string]*AsSplitArea
	AsSplitAreaOrder            uint
	AsSplitArea_stagedOrder     map[*AsSplitArea]uint
	AsSplitArea_orderStaged     map[uint]*AsSplitArea
	AsSplitAreas_reference      map[*AsSplitArea]*AsSplitArea
	AsSplitAreas_referenceOrder map[*AsSplitArea]uint

	// insertion point for slice of pointers maps
	OnAfterAsSplitAreaCreateCallback GongOnAfterCreateInterface[AsSplitArea]
	OnAfterAsSplitAreaUpdateCallback GongOnAfterUpdateInterface[AsSplitArea]
	OnAfterAsSplitAreaDeleteCallback GongOnAfterDeleteInterface[AsSplitArea]

	Buttons                map[*Button]struct{}
	Buttons_instance       map[*Button]*Button
	Buttons_mapString      map[string]*Button
	ButtonOrder            uint
	Button_stagedOrder     map[*Button]uint
	Button_orderStaged     map[uint]*Button
	Buttons_reference      map[*Button]*Button
	Buttons_referenceOrder map[*Button]uint

	// insertion point for slice of pointers maps
	OnAfterButtonCreateCallback GongOnAfterCreateInterface[Button]
	OnAfterButtonUpdateCallback GongOnAfterUpdateInterface[Button]
	OnAfterButtonDeleteCallback GongOnAfterDeleteInterface[Button]

	Cursors                map[*Cursor]struct{}
	Cursors_instance       map[*Cursor]*Cursor
	Cursors_mapString      map[string]*Cursor
	CursorOrder            uint
	Cursor_stagedOrder     map[*Cursor]uint
	Cursor_orderStaged     map[uint]*Cursor
	Cursors_reference      map[*Cursor]*Cursor
	Cursors_referenceOrder map[*Cursor]uint

	// insertion point for slice of pointers maps
	OnAfterCursorCreateCallback GongOnAfterCreateInterface[Cursor]
	OnAfterCursorUpdateCallback GongOnAfterUpdateInterface[Cursor]
	OnAfterCursorDeleteCallback GongOnAfterDeleteInterface[Cursor]

	FavIcons                map[*FavIcon]struct{}
	FavIcons_instance       map[*FavIcon]*FavIcon
	FavIcons_mapString      map[string]*FavIcon
	FavIconOrder            uint
	FavIcon_stagedOrder     map[*FavIcon]uint
	FavIcon_orderStaged     map[uint]*FavIcon
	FavIcons_reference      map[*FavIcon]*FavIcon
	FavIcons_referenceOrder map[*FavIcon]uint

	// insertion point for slice of pointers maps
	OnAfterFavIconCreateCallback GongOnAfterCreateInterface[FavIcon]
	OnAfterFavIconUpdateCallback GongOnAfterUpdateInterface[FavIcon]
	OnAfterFavIconDeleteCallback GongOnAfterDeleteInterface[FavIcon]

	Forms                map[*Form]struct{}
	Forms_instance       map[*Form]*Form
	Forms_mapString      map[string]*Form
	FormOrder            uint
	Form_stagedOrder     map[*Form]uint
	Form_orderStaged     map[uint]*Form
	Forms_reference      map[*Form]*Form
	Forms_referenceOrder map[*Form]uint

	// insertion point for slice of pointers maps
	OnAfterFormCreateCallback GongOnAfterCreateInterface[Form]
	OnAfterFormUpdateCallback GongOnAfterUpdateInterface[Form]
	OnAfterFormDeleteCallback GongOnAfterDeleteInterface[Form]

	Loads                map[*Load]struct{}
	Loads_instance       map[*Load]*Load
	Loads_mapString      map[string]*Load
	LoadOrder            uint
	Load_stagedOrder     map[*Load]uint
	Load_orderStaged     map[uint]*Load
	Loads_reference      map[*Load]*Load
	Loads_referenceOrder map[*Load]uint

	// insertion point for slice of pointers maps
	OnAfterLoadCreateCallback GongOnAfterCreateInterface[Load]
	OnAfterLoadUpdateCallback GongOnAfterUpdateInterface[Load]
	OnAfterLoadDeleteCallback GongOnAfterDeleteInterface[Load]

	LogoOnTheLefts                map[*LogoOnTheLeft]struct{}
	LogoOnTheLefts_instance       map[*LogoOnTheLeft]*LogoOnTheLeft
	LogoOnTheLefts_mapString      map[string]*LogoOnTheLeft
	LogoOnTheLeftOrder            uint
	LogoOnTheLeft_stagedOrder     map[*LogoOnTheLeft]uint
	LogoOnTheLeft_orderStaged     map[uint]*LogoOnTheLeft
	LogoOnTheLefts_reference      map[*LogoOnTheLeft]*LogoOnTheLeft
	LogoOnTheLefts_referenceOrder map[*LogoOnTheLeft]uint

	// insertion point for slice of pointers maps
	OnAfterLogoOnTheLeftCreateCallback GongOnAfterCreateInterface[LogoOnTheLeft]
	OnAfterLogoOnTheLeftUpdateCallback GongOnAfterUpdateInterface[LogoOnTheLeft]
	OnAfterLogoOnTheLeftDeleteCallback GongOnAfterDeleteInterface[LogoOnTheLeft]

	LogoOnTheRights                map[*LogoOnTheRight]struct{}
	LogoOnTheRights_instance       map[*LogoOnTheRight]*LogoOnTheRight
	LogoOnTheRights_mapString      map[string]*LogoOnTheRight
	LogoOnTheRightOrder            uint
	LogoOnTheRight_stagedOrder     map[*LogoOnTheRight]uint
	LogoOnTheRight_orderStaged     map[uint]*LogoOnTheRight
	LogoOnTheRights_reference      map[*LogoOnTheRight]*LogoOnTheRight
	LogoOnTheRights_referenceOrder map[*LogoOnTheRight]uint

	// insertion point for slice of pointers maps
	OnAfterLogoOnTheRightCreateCallback GongOnAfterCreateInterface[LogoOnTheRight]
	OnAfterLogoOnTheRightUpdateCallback GongOnAfterUpdateInterface[LogoOnTheRight]
	OnAfterLogoOnTheRightDeleteCallback GongOnAfterDeleteInterface[LogoOnTheRight]

	Markdowns                map[*Markdown]struct{}
	Markdowns_instance       map[*Markdown]*Markdown
	Markdowns_mapString      map[string]*Markdown
	MarkdownOrder            uint
	Markdown_stagedOrder     map[*Markdown]uint
	Markdown_orderStaged     map[uint]*Markdown
	Markdowns_reference      map[*Markdown]*Markdown
	Markdowns_referenceOrder map[*Markdown]uint

	// insertion point for slice of pointers maps
	OnAfterMarkdownCreateCallback GongOnAfterCreateInterface[Markdown]
	OnAfterMarkdownUpdateCallback GongOnAfterUpdateInterface[Markdown]
	OnAfterMarkdownDeleteCallback GongOnAfterDeleteInterface[Markdown]

	Sliders                map[*Slider]struct{}
	Sliders_instance       map[*Slider]*Slider
	Sliders_mapString      map[string]*Slider
	SliderOrder            uint
	Slider_stagedOrder     map[*Slider]uint
	Slider_orderStaged     map[uint]*Slider
	Sliders_reference      map[*Slider]*Slider
	Sliders_referenceOrder map[*Slider]uint

	// insertion point for slice of pointers maps
	OnAfterSliderCreateCallback GongOnAfterCreateInterface[Slider]
	OnAfterSliderUpdateCallback GongOnAfterUpdateInterface[Slider]
	OnAfterSliderDeleteCallback GongOnAfterDeleteInterface[Slider]

	Splits                map[*Split]struct{}
	Splits_instance       map[*Split]*Split
	Splits_mapString      map[string]*Split
	SplitOrder            uint
	Split_stagedOrder     map[*Split]uint
	Split_orderStaged     map[uint]*Split
	Splits_reference      map[*Split]*Split
	Splits_referenceOrder map[*Split]uint

	// insertion point for slice of pointers maps
	OnAfterSplitCreateCallback GongOnAfterCreateInterface[Split]
	OnAfterSplitUpdateCallback GongOnAfterUpdateInterface[Split]
	OnAfterSplitDeleteCallback GongOnAfterDeleteInterface[Split]

	Svgs                map[*Svg]struct{}
	Svgs_instance       map[*Svg]*Svg
	Svgs_mapString      map[string]*Svg
	SvgOrder            uint
	Svg_stagedOrder     map[*Svg]uint
	Svg_orderStaged     map[uint]*Svg
	Svgs_reference      map[*Svg]*Svg
	Svgs_referenceOrder map[*Svg]uint

	// insertion point for slice of pointers maps
	OnAfterSvgCreateCallback GongOnAfterCreateInterface[Svg]
	OnAfterSvgUpdateCallback GongOnAfterUpdateInterface[Svg]
	OnAfterSvgDeleteCallback GongOnAfterDeleteInterface[Svg]

	Tables                map[*Table]struct{}
	Tables_instance       map[*Table]*Table
	Tables_mapString      map[string]*Table
	TableOrder            uint
	Table_stagedOrder     map[*Table]uint
	Table_orderStaged     map[uint]*Table
	Tables_reference      map[*Table]*Table
	Tables_referenceOrder map[*Table]uint

	// insertion point for slice of pointers maps
	OnAfterTableCreateCallback GongOnAfterCreateInterface[Table]
	OnAfterTableUpdateCallback GongOnAfterUpdateInterface[Table]
	OnAfterTableDeleteCallback GongOnAfterDeleteInterface[Table]

	Threejss                map[*Threejs]struct{}
	Threejss_instance       map[*Threejs]*Threejs
	Threejss_mapString      map[string]*Threejs
	ThreejsOrder            uint
	Threejs_stagedOrder     map[*Threejs]uint
	Threejs_orderStaged     map[uint]*Threejs
	Threejss_reference      map[*Threejs]*Threejs
	Threejss_referenceOrder map[*Threejs]uint

	// insertion point for slice of pointers maps
	OnAfterThreejsCreateCallback GongOnAfterCreateInterface[Threejs]
	OnAfterThreejsUpdateCallback GongOnAfterUpdateInterface[Threejs]
	OnAfterThreejsDeleteCallback GongOnAfterDeleteInterface[Threejs]

	Titles                map[*Title]struct{}
	Titles_instance       map[*Title]*Title
	Titles_mapString      map[string]*Title
	TitleOrder            uint
	Title_stagedOrder     map[*Title]uint
	Title_orderStaged     map[uint]*Title
	Titles_reference      map[*Title]*Title
	Titles_referenceOrder map[*Title]uint

	// insertion point for slice of pointers maps
	OnAfterTitleCreateCallback GongOnAfterCreateInterface[Title]
	OnAfterTitleUpdateCallback GongOnAfterUpdateInterface[Title]
	OnAfterTitleDeleteCallback GongOnAfterDeleteInterface[Title]

	Tones                map[*Tone]struct{}
	Tones_instance       map[*Tone]*Tone
	Tones_mapString      map[string]*Tone
	ToneOrder            uint
	Tone_stagedOrder     map[*Tone]uint
	Tone_orderStaged     map[uint]*Tone
	Tones_reference      map[*Tone]*Tone
	Tones_referenceOrder map[*Tone]uint

	// insertion point for slice of pointers maps
	OnAfterToneCreateCallback GongOnAfterCreateInterface[Tone]
	OnAfterToneUpdateCallback GongOnAfterUpdateInterface[Tone]
	OnAfterToneDeleteCallback GongOnAfterDeleteInterface[Tone]

	Trees                map[*Tree]struct{}
	Trees_instance       map[*Tree]*Tree
	Trees_mapString      map[string]*Tree
	TreeOrder            uint
	Tree_stagedOrder     map[*Tree]uint
	Tree_orderStaged     map[uint]*Tree
	Trees_reference      map[*Tree]*Tree
	Trees_referenceOrder map[*Tree]uint

	// insertion point for slice of pointers maps
	OnAfterTreeCreateCallback GongOnAfterCreateInterface[Tree]
	OnAfterTreeUpdateCallback GongOnAfterUpdateInterface[Tree]
	OnAfterTreeDeleteCallback GongOnAfterDeleteInterface[Tree]

	Views                map[*View]struct{}
	Views_instance       map[*View]*View
	Views_mapString      map[string]*View
	ViewOrder            uint
	View_stagedOrder     map[*View]uint
	View_orderStaged     map[uint]*View
	Views_reference      map[*View]*View
	Views_referenceOrder map[*View]uint

	// insertion point for slice of pointers maps
	View_RootAsSplitAreas_reverseMap map[*AsSplitArea]*View

	OnAfterViewCreateCallback GongOnAfterCreateInterface[View]
	OnAfterViewUpdateCallback GongOnAfterUpdateInterface[View]
	OnAfterViewDeleteCallback GongOnAfterDeleteInterface[View]

	Xlsxs                map[*Xlsx]struct{}
	Xlsxs_instance       map[*Xlsx]*Xlsx
	Xlsxs_mapString      map[string]*Xlsx
	XlsxOrder            uint
	Xlsx_stagedOrder     map[*Xlsx]uint
	Xlsx_orderStaged     map[uint]*Xlsx
	Xlsxs_reference      map[*Xlsx]*Xlsx
	Xlsxs_referenceOrder map[*Xlsx]uint

	// insertion point for slice of pointers maps
	OnAfterXlsxCreateCallback GongOnAfterCreateInterface[Xlsx]
	OnAfterXlsxUpdateCallback GongOnAfterUpdateInterface[Xlsx]
	OnAfterXlsxDeleteCallback GongOnAfterDeleteInterface[Xlsx]

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
	__gong__clearReferences(&stage.AsSplits_reference, &stage.AsSplits_instance, &stage.AsSplits_referenceOrder)

	__gong__clearReferences(&stage.AsSplitAreas_reference, &stage.AsSplitAreas_instance, &stage.AsSplitAreas_referenceOrder)

	__gong__clearReferences(&stage.Buttons_reference, &stage.Buttons_instance, &stage.Buttons_referenceOrder)

	__gong__clearReferences(&stage.Cursors_reference, &stage.Cursors_instance, &stage.Cursors_referenceOrder)

	__gong__clearReferences(&stage.FavIcons_reference, &stage.FavIcons_instance, &stage.FavIcons_referenceOrder)

	__gong__clearReferences(&stage.Forms_reference, &stage.Forms_instance, &stage.Forms_referenceOrder)

	__gong__clearReferences(&stage.Loads_reference, &stage.Loads_instance, &stage.Loads_referenceOrder)

	__gong__clearReferences(&stage.LogoOnTheLefts_reference, &stage.LogoOnTheLefts_instance, &stage.LogoOnTheLefts_referenceOrder)

	__gong__clearReferences(&stage.LogoOnTheRights_reference, &stage.LogoOnTheRights_instance, &stage.LogoOnTheRights_referenceOrder)

	__gong__clearReferences(&stage.Markdowns_reference, &stage.Markdowns_instance, &stage.Markdowns_referenceOrder)

	__gong__clearReferences(&stage.Sliders_reference, &stage.Sliders_instance, &stage.Sliders_referenceOrder)

	__gong__clearReferences(&stage.Splits_reference, &stage.Splits_instance, &stage.Splits_referenceOrder)

	__gong__clearReferences(&stage.Svgs_reference, &stage.Svgs_instance, &stage.Svgs_referenceOrder)

	__gong__clearReferences(&stage.Tables_reference, &stage.Tables_instance, &stage.Tables_referenceOrder)

	__gong__clearReferences(&stage.Threejss_reference, &stage.Threejss_instance, &stage.Threejss_referenceOrder)

	__gong__clearReferences(&stage.Titles_reference, &stage.Titles_instance, &stage.Titles_referenceOrder)

	__gong__clearReferences(&stage.Tones_reference, &stage.Tones_instance, &stage.Tones_referenceOrder)

	__gong__clearReferences(&stage.Trees_reference, &stage.Trees_instance, &stage.Trees_referenceOrder)

	__gong__clearReferences(&stage.Views_reference, &stage.Views_instance, &stage.Views_referenceOrder)

	__gong__clearReferences(&stage.Xlsxs_reference, &stage.Xlsxs_instance, &stage.Xlsxs_referenceOrder)

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
	stage.AsSplitOrder = __gong__recomputeOrder(stage.AsSplit_stagedOrder)

	stage.AsSplitAreaOrder = __gong__recomputeOrder(stage.AsSplitArea_stagedOrder)

	stage.ButtonOrder = __gong__recomputeOrder(stage.Button_stagedOrder)

	stage.CursorOrder = __gong__recomputeOrder(stage.Cursor_stagedOrder)

	stage.FavIconOrder = __gong__recomputeOrder(stage.FavIcon_stagedOrder)

	stage.FormOrder = __gong__recomputeOrder(stage.Form_stagedOrder)

	stage.LoadOrder = __gong__recomputeOrder(stage.Load_stagedOrder)

	stage.LogoOnTheLeftOrder = __gong__recomputeOrder(stage.LogoOnTheLeft_stagedOrder)

	stage.LogoOnTheRightOrder = __gong__recomputeOrder(stage.LogoOnTheRight_stagedOrder)

	stage.MarkdownOrder = __gong__recomputeOrder(stage.Markdown_stagedOrder)

	stage.SliderOrder = __gong__recomputeOrder(stage.Slider_stagedOrder)

	stage.SplitOrder = __gong__recomputeOrder(stage.Split_stagedOrder)

	stage.SvgOrder = __gong__recomputeOrder(stage.Svg_stagedOrder)

	stage.TableOrder = __gong__recomputeOrder(stage.Table_stagedOrder)

	stage.ThreejsOrder = __gong__recomputeOrder(stage.Threejs_stagedOrder)

	stage.TitleOrder = __gong__recomputeOrder(stage.Title_stagedOrder)

	stage.ToneOrder = __gong__recomputeOrder(stage.Tone_stagedOrder)

	stage.TreeOrder = __gong__recomputeOrder(stage.Tree_stagedOrder)

	stage.ViewOrder = __gong__recomputeOrder(stage.View_stagedOrder)

	stage.XlsxOrder = __gong__recomputeOrder(stage.Xlsx_stagedOrder)

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
	case *AsSplit:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.AsSplits, stage.AsSplit_stagedOrder))
	case *AsSplitArea:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.AsSplitAreas, stage.AsSplitArea_stagedOrder))
	case *Button:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Buttons, stage.Button_stagedOrder))
	case *Cursor:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Cursors, stage.Cursor_stagedOrder))
	case *FavIcon:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.FavIcons, stage.FavIcon_stagedOrder))
	case *Form:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Forms, stage.Form_stagedOrder))
	case *Load:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Loads, stage.Load_stagedOrder))
	case *LogoOnTheLeft:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.LogoOnTheLefts, stage.LogoOnTheLeft_stagedOrder))
	case *LogoOnTheRight:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.LogoOnTheRights, stage.LogoOnTheRight_stagedOrder))
	case *Markdown:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Markdowns, stage.Markdown_stagedOrder))
	case *Slider:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Sliders, stage.Slider_stagedOrder))
	case *Split:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Splits, stage.Split_stagedOrder))
	case *Svg:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Svgs, stage.Svg_stagedOrder))
	case *Table:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Tables, stage.Table_stagedOrder))
	case *Threejs:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Threejss, stage.Threejs_stagedOrder))
	case *Title:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Titles, stage.Title_stagedOrder))
	case *Tone:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Tones, stage.Tone_stagedOrder))
	case *Tree:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Trees, stage.Tree_stagedOrder))
	case *View:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Views, stage.View_stagedOrder))
	case *Xlsx:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Xlsxs, stage.Xlsx_stagedOrder))

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
	return "github.com/fullstack-lang/gong/lib/split/go/models"
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
		AsSplits:           make(map[*AsSplit]struct{}),
		AsSplits_mapString: make(map[string]*AsSplit),

		AsSplitAreas:           make(map[*AsSplitArea]struct{}),
		AsSplitAreas_mapString: make(map[string]*AsSplitArea),

		Buttons:           make(map[*Button]struct{}),
		Buttons_mapString: make(map[string]*Button),

		Cursors:           make(map[*Cursor]struct{}),
		Cursors_mapString: make(map[string]*Cursor),

		FavIcons:           make(map[*FavIcon]struct{}),
		FavIcons_mapString: make(map[string]*FavIcon),

		Forms:           make(map[*Form]struct{}),
		Forms_mapString: make(map[string]*Form),

		Loads:           make(map[*Load]struct{}),
		Loads_mapString: make(map[string]*Load),

		LogoOnTheLefts:           make(map[*LogoOnTheLeft]struct{}),
		LogoOnTheLefts_mapString: make(map[string]*LogoOnTheLeft),

		LogoOnTheRights:           make(map[*LogoOnTheRight]struct{}),
		LogoOnTheRights_mapString: make(map[string]*LogoOnTheRight),

		Markdowns:           make(map[*Markdown]struct{}),
		Markdowns_mapString: make(map[string]*Markdown),

		Sliders:           make(map[*Slider]struct{}),
		Sliders_mapString: make(map[string]*Slider),

		Splits:           make(map[*Split]struct{}),
		Splits_mapString: make(map[string]*Split),

		Svgs:           make(map[*Svg]struct{}),
		Svgs_mapString: make(map[string]*Svg),

		Tables:           make(map[*Table]struct{}),
		Tables_mapString: make(map[string]*Table),

		Threejss:           make(map[*Threejs]struct{}),
		Threejss_mapString: make(map[string]*Threejs),

		Titles:           make(map[*Title]struct{}),
		Titles_mapString: make(map[string]*Title),

		Tones:           make(map[*Tone]struct{}),
		Tones_mapString: make(map[string]*Tone),

		Trees:           make(map[*Tree]struct{}),
		Trees_mapString: make(map[string]*Tree),

		Views:           make(map[*View]struct{}),
		Views_mapString: make(map[string]*View),

		Xlsxs:           make(map[*Xlsx]struct{}),
		Xlsxs_mapString: make(map[string]*Xlsx),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		AsSplit_stagedOrder: make(map[*AsSplit]uint),
		AsSplit_orderStaged: make(map[uint]*AsSplit),
		AsSplits_reference:  make(map[*AsSplit]*AsSplit),

		AsSplitArea_stagedOrder: make(map[*AsSplitArea]uint),
		AsSplitArea_orderStaged: make(map[uint]*AsSplitArea),
		AsSplitAreas_reference:  make(map[*AsSplitArea]*AsSplitArea),

		Button_stagedOrder: make(map[*Button]uint),
		Button_orderStaged: make(map[uint]*Button),
		Buttons_reference:  make(map[*Button]*Button),

		Cursor_stagedOrder: make(map[*Cursor]uint),
		Cursor_orderStaged: make(map[uint]*Cursor),
		Cursors_reference:  make(map[*Cursor]*Cursor),

		FavIcon_stagedOrder: make(map[*FavIcon]uint),
		FavIcon_orderStaged: make(map[uint]*FavIcon),
		FavIcons_reference:  make(map[*FavIcon]*FavIcon),

		Form_stagedOrder: make(map[*Form]uint),
		Form_orderStaged: make(map[uint]*Form),
		Forms_reference:  make(map[*Form]*Form),

		Load_stagedOrder: make(map[*Load]uint),
		Load_orderStaged: make(map[uint]*Load),
		Loads_reference:  make(map[*Load]*Load),

		LogoOnTheLeft_stagedOrder: make(map[*LogoOnTheLeft]uint),
		LogoOnTheLeft_orderStaged: make(map[uint]*LogoOnTheLeft),
		LogoOnTheLefts_reference:  make(map[*LogoOnTheLeft]*LogoOnTheLeft),

		LogoOnTheRight_stagedOrder: make(map[*LogoOnTheRight]uint),
		LogoOnTheRight_orderStaged: make(map[uint]*LogoOnTheRight),
		LogoOnTheRights_reference:  make(map[*LogoOnTheRight]*LogoOnTheRight),

		Markdown_stagedOrder: make(map[*Markdown]uint),
		Markdown_orderStaged: make(map[uint]*Markdown),
		Markdowns_reference:  make(map[*Markdown]*Markdown),

		Slider_stagedOrder: make(map[*Slider]uint),
		Slider_orderStaged: make(map[uint]*Slider),
		Sliders_reference:  make(map[*Slider]*Slider),

		Split_stagedOrder: make(map[*Split]uint),
		Split_orderStaged: make(map[uint]*Split),
		Splits_reference:  make(map[*Split]*Split),

		Svg_stagedOrder: make(map[*Svg]uint),
		Svg_orderStaged: make(map[uint]*Svg),
		Svgs_reference:  make(map[*Svg]*Svg),

		Table_stagedOrder: make(map[*Table]uint),
		Table_orderStaged: make(map[uint]*Table),
		Tables_reference:  make(map[*Table]*Table),

		Threejs_stagedOrder: make(map[*Threejs]uint),
		Threejs_orderStaged: make(map[uint]*Threejs),
		Threejss_reference:  make(map[*Threejs]*Threejs),

		Title_stagedOrder: make(map[*Title]uint),
		Title_orderStaged: make(map[uint]*Title),
		Titles_reference:  make(map[*Title]*Title),

		Tone_stagedOrder: make(map[*Tone]uint),
		Tone_orderStaged: make(map[uint]*Tone),
		Tones_reference:  make(map[*Tone]*Tone),

		Tree_stagedOrder: make(map[*Tree]uint),
		Tree_orderStaged: make(map[uint]*Tree),
		Trees_reference:  make(map[*Tree]*Tree),

		View_stagedOrder: make(map[*View]uint),
		View_orderStaged: make(map[uint]*View),
		Views_reference:  make(map[*View]*View),

		Xlsx_stagedOrder: make(map[*Xlsx]uint),
		Xlsx_orderStaged: make(map[uint]*Xlsx),
		Xlsxs_reference:  make(map[*Xlsx]*Xlsx),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"AsSplit": &AsSplitUnmarshaller{},

			"AsSplitArea": &AsSplitAreaUnmarshaller{},

			"Button": &ButtonUnmarshaller{},

			"Cursor": &CursorUnmarshaller{},

			"FavIcon": &FavIconUnmarshaller{},

			"Form": &FormUnmarshaller{},

			"Load": &LoadUnmarshaller{},

			"LogoOnTheLeft": &LogoOnTheLeftUnmarshaller{},

			"LogoOnTheRight": &LogoOnTheRightUnmarshaller{},

			"Markdown": &MarkdownUnmarshaller{},

			"Slider": &SliderUnmarshaller{},

			"Split": &SplitUnmarshaller{},

			"Svg": &SvgUnmarshaller{},

			"Table": &TableUnmarshaller{},

			"Threejs": &ThreejsUnmarshaller{},

			"Title": &TitleUnmarshaller{},

			"Tone": &ToneUnmarshaller{},

			"Tree": &TreeUnmarshaller{},

			"View": &ViewUnmarshaller{},

			"Xlsx": &XlsxUnmarshaller{},

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
	case *AsSplit:
		return any(stage.AsSplit_orderStaged[order]).(Type)
	case *AsSplitArea:
		return any(stage.AsSplitArea_orderStaged[order]).(Type)
	case *Button:
		return any(stage.Button_orderStaged[order]).(Type)
	case *Cursor:
		return any(stage.Cursor_orderStaged[order]).(Type)
	case *FavIcon:
		return any(stage.FavIcon_orderStaged[order]).(Type)
	case *Form:
		return any(stage.Form_orderStaged[order]).(Type)
	case *Load:
		return any(stage.Load_orderStaged[order]).(Type)
	case *LogoOnTheLeft:
		return any(stage.LogoOnTheLeft_orderStaged[order]).(Type)
	case *LogoOnTheRight:
		return any(stage.LogoOnTheRight_orderStaged[order]).(Type)
	case *Markdown:
		return any(stage.Markdown_orderStaged[order]).(Type)
	case *Slider:
		return any(stage.Slider_orderStaged[order]).(Type)
	case *Split:
		return any(stage.Split_orderStaged[order]).(Type)
	case *Svg:
		return any(stage.Svg_orderStaged[order]).(Type)
	case *Table:
		return any(stage.Table_orderStaged[order]).(Type)
	case *Threejs:
		return any(stage.Threejs_orderStaged[order]).(Type)
	case *Title:
		return any(stage.Title_orderStaged[order]).(Type)
	case *Tone:
		return any(stage.Tone_orderStaged[order]).(Type)
	case *Tree:
		return any(stage.Tree_orderStaged[order]).(Type)
	case *View:
		return any(stage.View_orderStaged[order]).(Type)
	case *Xlsx:
		return any(stage.Xlsx_orderStaged[order]).(Type)
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
	stage.Map_GongStructName_InstancesNb["AsSplit"] = len(stage.AsSplits)
	stage.Map_GongStructName_InstancesNb["AsSplitArea"] = len(stage.AsSplitAreas)
	stage.Map_GongStructName_InstancesNb["Button"] = len(stage.Buttons)
	stage.Map_GongStructName_InstancesNb["Cursor"] = len(stage.Cursors)
	stage.Map_GongStructName_InstancesNb["FavIcon"] = len(stage.FavIcons)
	stage.Map_GongStructName_InstancesNb["Form"] = len(stage.Forms)
	stage.Map_GongStructName_InstancesNb["Load"] = len(stage.Loads)
	stage.Map_GongStructName_InstancesNb["LogoOnTheLeft"] = len(stage.LogoOnTheLefts)
	stage.Map_GongStructName_InstancesNb["LogoOnTheRight"] = len(stage.LogoOnTheRights)
	stage.Map_GongStructName_InstancesNb["Markdown"] = len(stage.Markdowns)
	stage.Map_GongStructName_InstancesNb["Slider"] = len(stage.Sliders)
	stage.Map_GongStructName_InstancesNb["Split"] = len(stage.Splits)
	stage.Map_GongStructName_InstancesNb["Svg"] = len(stage.Svgs)
	stage.Map_GongStructName_InstancesNb["Table"] = len(stage.Tables)
	stage.Map_GongStructName_InstancesNb["Threejs"] = len(stage.Threejss)
	stage.Map_GongStructName_InstancesNb["Title"] = len(stage.Titles)
	stage.Map_GongStructName_InstancesNb["Tone"] = len(stage.Tones)
	stage.Map_GongStructName_InstancesNb["Tree"] = len(stage.Trees)
	stage.Map_GongStructName_InstancesNb["View"] = len(stage.Views)
	stage.Map_GongStructName_InstancesNb["Xlsx"] = len(stage.Xlsxs)
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
// Stage puts assplit to the model stage
func (assplit *AsSplit) Stage(stage *Stage) *AsSplit {
	__gong__stage(stage.AsSplits, stage.AsSplit_stagedOrder, stage.AsSplit_orderStaged, &stage.AsSplitOrder, stage.AsSplits_mapString, assplit, assplit.Name)
	return assplit
}

// StagePreserveOrder puts assplit to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AsSplitOrder
// - update stage.AsSplitOrder accordingly
func (assplit *AsSplit) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AsSplits, stage.AsSplit_stagedOrder, stage.AsSplit_orderStaged, &stage.AsSplitOrder, stage.AsSplits_mapString, assplit, order, assplit.Name)
}

// Unstage removes assplit off the model stage
func (assplit *AsSplit) Unstage(stage *Stage) *AsSplit {
	__gong__unstage(stage.AsSplits, stage.AsSplits_mapString, assplit, assplit.Name)
	return assplit
}

// UnstageVoid removes assplit off the model stage
func (assplit *AsSplit) UnstageVoid(stage *Stage) {
	assplit.Unstage(stage)
}

func (assplit *AsSplit) StageVoid(stage *Stage) {
	assplit.Stage(stage)
}

// for satisfaction of GongStruct interface
func (assplit *AsSplit) GetName() (res string) {
	return assplit.Name
}

// for satisfaction of GongStruct interface
func (assplit *AsSplit) SetName(name string) {
	assplit.Name = name
}

// Stage puts assplitarea to the model stage
func (assplitarea *AsSplitArea) Stage(stage *Stage) *AsSplitArea {
	__gong__stage(stage.AsSplitAreas, stage.AsSplitArea_stagedOrder, stage.AsSplitArea_orderStaged, &stage.AsSplitAreaOrder, stage.AsSplitAreas_mapString, assplitarea, assplitarea.Name)
	return assplitarea
}

// StagePreserveOrder puts assplitarea to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AsSplitAreaOrder
// - update stage.AsSplitAreaOrder accordingly
func (assplitarea *AsSplitArea) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AsSplitAreas, stage.AsSplitArea_stagedOrder, stage.AsSplitArea_orderStaged, &stage.AsSplitAreaOrder, stage.AsSplitAreas_mapString, assplitarea, order, assplitarea.Name)
}

// Unstage removes assplitarea off the model stage
func (assplitarea *AsSplitArea) Unstage(stage *Stage) *AsSplitArea {
	__gong__unstage(stage.AsSplitAreas, stage.AsSplitAreas_mapString, assplitarea, assplitarea.Name)
	return assplitarea
}

// UnstageVoid removes assplitarea off the model stage
func (assplitarea *AsSplitArea) UnstageVoid(stage *Stage) {
	assplitarea.Unstage(stage)
}

func (assplitarea *AsSplitArea) StageVoid(stage *Stage) {
	assplitarea.Stage(stage)
}

// for satisfaction of GongStruct interface
func (assplitarea *AsSplitArea) GetName() (res string) {
	return assplitarea.Name
}

// for satisfaction of GongStruct interface
func (assplitarea *AsSplitArea) SetName(name string) {
	assplitarea.Name = name
}

// Stage puts button to the model stage
func (button *Button) Stage(stage *Stage) *Button {
	__gong__stage(stage.Buttons, stage.Button_stagedOrder, stage.Button_orderStaged, &stage.ButtonOrder, stage.Buttons_mapString, button, button.Name)
	return button
}

// StagePreserveOrder puts button to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ButtonOrder
// - update stage.ButtonOrder accordingly
func (button *Button) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Buttons, stage.Button_stagedOrder, stage.Button_orderStaged, &stage.ButtonOrder, stage.Buttons_mapString, button, order, button.Name)
}

// Unstage removes button off the model stage
func (button *Button) Unstage(stage *Stage) *Button {
	__gong__unstage(stage.Buttons, stage.Buttons_mapString, button, button.Name)
	return button
}

// UnstageVoid removes button off the model stage
func (button *Button) UnstageVoid(stage *Stage) {
	button.Unstage(stage)
}

func (button *Button) StageVoid(stage *Stage) {
	button.Stage(stage)
}

// for satisfaction of GongStruct interface
func (button *Button) GetName() (res string) {
	return button.Name
}

// for satisfaction of GongStruct interface
func (button *Button) SetName(name string) {
	button.Name = name
}

// Stage puts cursor to the model stage
func (cursor *Cursor) Stage(stage *Stage) *Cursor {
	__gong__stage(stage.Cursors, stage.Cursor_stagedOrder, stage.Cursor_orderStaged, &stage.CursorOrder, stage.Cursors_mapString, cursor, cursor.Name)
	return cursor
}

// StagePreserveOrder puts cursor to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CursorOrder
// - update stage.CursorOrder accordingly
func (cursor *Cursor) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Cursors, stage.Cursor_stagedOrder, stage.Cursor_orderStaged, &stage.CursorOrder, stage.Cursors_mapString, cursor, order, cursor.Name)
}

// Unstage removes cursor off the model stage
func (cursor *Cursor) Unstage(stage *Stage) *Cursor {
	__gong__unstage(stage.Cursors, stage.Cursors_mapString, cursor, cursor.Name)
	return cursor
}

// UnstageVoid removes cursor off the model stage
func (cursor *Cursor) UnstageVoid(stage *Stage) {
	cursor.Unstage(stage)
}

func (cursor *Cursor) StageVoid(stage *Stage) {
	cursor.Stage(stage)
}

// for satisfaction of GongStruct interface
func (cursor *Cursor) GetName() (res string) {
	return cursor.Name
}

// for satisfaction of GongStruct interface
func (cursor *Cursor) SetName(name string) {
	cursor.Name = name
}

// Stage puts favicon to the model stage
func (favicon *FavIcon) Stage(stage *Stage) *FavIcon {
	__gong__stage(stage.FavIcons, stage.FavIcon_stagedOrder, stage.FavIcon_orderStaged, &stage.FavIconOrder, stage.FavIcons_mapString, favicon, favicon.Name)
	return favicon
}

// StagePreserveOrder puts favicon to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FavIconOrder
// - update stage.FavIconOrder accordingly
func (favicon *FavIcon) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FavIcons, stage.FavIcon_stagedOrder, stage.FavIcon_orderStaged, &stage.FavIconOrder, stage.FavIcons_mapString, favicon, order, favicon.Name)
}

// Unstage removes favicon off the model stage
func (favicon *FavIcon) Unstage(stage *Stage) *FavIcon {
	__gong__unstage(stage.FavIcons, stage.FavIcons_mapString, favicon, favicon.Name)
	return favicon
}

// UnstageVoid removes favicon off the model stage
func (favicon *FavIcon) UnstageVoid(stage *Stage) {
	favicon.Unstage(stage)
}

func (favicon *FavIcon) StageVoid(stage *Stage) {
	favicon.Stage(stage)
}

// for satisfaction of GongStruct interface
func (favicon *FavIcon) GetName() (res string) {
	return favicon.Name
}

// for satisfaction of GongStruct interface
func (favicon *FavIcon) SetName(name string) {
	favicon.Name = name
}

// Stage puts form to the model stage
func (form *Form) Stage(stage *Stage) *Form {
	__gong__stage(stage.Forms, stage.Form_stagedOrder, stage.Form_orderStaged, &stage.FormOrder, stage.Forms_mapString, form, form.Name)
	return form
}

// StagePreserveOrder puts form to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormOrder
// - update stage.FormOrder accordingly
func (form *Form) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Forms, stage.Form_stagedOrder, stage.Form_orderStaged, &stage.FormOrder, stage.Forms_mapString, form, order, form.Name)
}

// Unstage removes form off the model stage
func (form *Form) Unstage(stage *Stage) *Form {
	__gong__unstage(stage.Forms, stage.Forms_mapString, form, form.Name)
	return form
}

// UnstageVoid removes form off the model stage
func (form *Form) UnstageVoid(stage *Stage) {
	form.Unstage(stage)
}

func (form *Form) StageVoid(stage *Stage) {
	form.Stage(stage)
}

// for satisfaction of GongStruct interface
func (form *Form) GetName() (res string) {
	return form.Name
}

// for satisfaction of GongStruct interface
func (form *Form) SetName(name string) {
	form.Name = name
}

// Stage puts load to the model stage
func (load *Load) Stage(stage *Stage) *Load {
	__gong__stage(stage.Loads, stage.Load_stagedOrder, stage.Load_orderStaged, &stage.LoadOrder, stage.Loads_mapString, load, load.Name)
	return load
}

// StagePreserveOrder puts load to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LoadOrder
// - update stage.LoadOrder accordingly
func (load *Load) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Loads, stage.Load_stagedOrder, stage.Load_orderStaged, &stage.LoadOrder, stage.Loads_mapString, load, order, load.Name)
}

// Unstage removes load off the model stage
func (load *Load) Unstage(stage *Stage) *Load {
	__gong__unstage(stage.Loads, stage.Loads_mapString, load, load.Name)
	return load
}

// UnstageVoid removes load off the model stage
func (load *Load) UnstageVoid(stage *Stage) {
	load.Unstage(stage)
}

func (load *Load) StageVoid(stage *Stage) {
	load.Stage(stage)
}

// for satisfaction of GongStruct interface
func (load *Load) GetName() (res string) {
	return load.Name
}

// for satisfaction of GongStruct interface
func (load *Load) SetName(name string) {
	load.Name = name
}

// Stage puts logoontheleft to the model stage
func (logoontheleft *LogoOnTheLeft) Stage(stage *Stage) *LogoOnTheLeft {
	__gong__stage(stage.LogoOnTheLefts, stage.LogoOnTheLeft_stagedOrder, stage.LogoOnTheLeft_orderStaged, &stage.LogoOnTheLeftOrder, stage.LogoOnTheLefts_mapString, logoontheleft, logoontheleft.Name)
	return logoontheleft
}

// StagePreserveOrder puts logoontheleft to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LogoOnTheLeftOrder
// - update stage.LogoOnTheLeftOrder accordingly
func (logoontheleft *LogoOnTheLeft) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.LogoOnTheLefts, stage.LogoOnTheLeft_stagedOrder, stage.LogoOnTheLeft_orderStaged, &stage.LogoOnTheLeftOrder, stage.LogoOnTheLefts_mapString, logoontheleft, order, logoontheleft.Name)
}

// Unstage removes logoontheleft off the model stage
func (logoontheleft *LogoOnTheLeft) Unstage(stage *Stage) *LogoOnTheLeft {
	__gong__unstage(stage.LogoOnTheLefts, stage.LogoOnTheLefts_mapString, logoontheleft, logoontheleft.Name)
	return logoontheleft
}

// UnstageVoid removes logoontheleft off the model stage
func (logoontheleft *LogoOnTheLeft) UnstageVoid(stage *Stage) {
	logoontheleft.Unstage(stage)
}

func (logoontheleft *LogoOnTheLeft) StageVoid(stage *Stage) {
	logoontheleft.Stage(stage)
}

// for satisfaction of GongStruct interface
func (logoontheleft *LogoOnTheLeft) GetName() (res string) {
	return logoontheleft.Name
}

// for satisfaction of GongStruct interface
func (logoontheleft *LogoOnTheLeft) SetName(name string) {
	logoontheleft.Name = name
}

// Stage puts logoontheright to the model stage
func (logoontheright *LogoOnTheRight) Stage(stage *Stage) *LogoOnTheRight {
	__gong__stage(stage.LogoOnTheRights, stage.LogoOnTheRight_stagedOrder, stage.LogoOnTheRight_orderStaged, &stage.LogoOnTheRightOrder, stage.LogoOnTheRights_mapString, logoontheright, logoontheright.Name)
	return logoontheright
}

// StagePreserveOrder puts logoontheright to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LogoOnTheRightOrder
// - update stage.LogoOnTheRightOrder accordingly
func (logoontheright *LogoOnTheRight) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.LogoOnTheRights, stage.LogoOnTheRight_stagedOrder, stage.LogoOnTheRight_orderStaged, &stage.LogoOnTheRightOrder, stage.LogoOnTheRights_mapString, logoontheright, order, logoontheright.Name)
}

// Unstage removes logoontheright off the model stage
func (logoontheright *LogoOnTheRight) Unstage(stage *Stage) *LogoOnTheRight {
	__gong__unstage(stage.LogoOnTheRights, stage.LogoOnTheRights_mapString, logoontheright, logoontheright.Name)
	return logoontheright
}

// UnstageVoid removes logoontheright off the model stage
func (logoontheright *LogoOnTheRight) UnstageVoid(stage *Stage) {
	logoontheright.Unstage(stage)
}

func (logoontheright *LogoOnTheRight) StageVoid(stage *Stage) {
	logoontheright.Stage(stage)
}

// for satisfaction of GongStruct interface
func (logoontheright *LogoOnTheRight) GetName() (res string) {
	return logoontheright.Name
}

// for satisfaction of GongStruct interface
func (logoontheright *LogoOnTheRight) SetName(name string) {
	logoontheright.Name = name
}

// Stage puts markdown to the model stage
func (markdown *Markdown) Stage(stage *Stage) *Markdown {
	__gong__stage(stage.Markdowns, stage.Markdown_stagedOrder, stage.Markdown_orderStaged, &stage.MarkdownOrder, stage.Markdowns_mapString, markdown, markdown.Name)
	return markdown
}

// StagePreserveOrder puts markdown to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MarkdownOrder
// - update stage.MarkdownOrder accordingly
func (markdown *Markdown) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Markdowns, stage.Markdown_stagedOrder, stage.Markdown_orderStaged, &stage.MarkdownOrder, stage.Markdowns_mapString, markdown, order, markdown.Name)
}

// Unstage removes markdown off the model stage
func (markdown *Markdown) Unstage(stage *Stage) *Markdown {
	__gong__unstage(stage.Markdowns, stage.Markdowns_mapString, markdown, markdown.Name)
	return markdown
}

// UnstageVoid removes markdown off the model stage
func (markdown *Markdown) UnstageVoid(stage *Stage) {
	markdown.Unstage(stage)
}

func (markdown *Markdown) StageVoid(stage *Stage) {
	markdown.Stage(stage)
}

// for satisfaction of GongStruct interface
func (markdown *Markdown) GetName() (res string) {
	return markdown.Name
}

// for satisfaction of GongStruct interface
func (markdown *Markdown) SetName(name string) {
	markdown.Name = name
}

// Stage puts slider to the model stage
func (slider *Slider) Stage(stage *Stage) *Slider {
	__gong__stage(stage.Sliders, stage.Slider_stagedOrder, stage.Slider_orderStaged, &stage.SliderOrder, stage.Sliders_mapString, slider, slider.Name)
	return slider
}

// StagePreserveOrder puts slider to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SliderOrder
// - update stage.SliderOrder accordingly
func (slider *Slider) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Sliders, stage.Slider_stagedOrder, stage.Slider_orderStaged, &stage.SliderOrder, stage.Sliders_mapString, slider, order, slider.Name)
}

// Unstage removes slider off the model stage
func (slider *Slider) Unstage(stage *Stage) *Slider {
	__gong__unstage(stage.Sliders, stage.Sliders_mapString, slider, slider.Name)
	return slider
}

// UnstageVoid removes slider off the model stage
func (slider *Slider) UnstageVoid(stage *Stage) {
	slider.Unstage(stage)
}

func (slider *Slider) StageVoid(stage *Stage) {
	slider.Stage(stage)
}

// for satisfaction of GongStruct interface
func (slider *Slider) GetName() (res string) {
	return slider.Name
}

// for satisfaction of GongStruct interface
func (slider *Slider) SetName(name string) {
	slider.Name = name
}

// Stage puts split to the model stage
func (split *Split) Stage(stage *Stage) *Split {
	__gong__stage(stage.Splits, stage.Split_stagedOrder, stage.Split_orderStaged, &stage.SplitOrder, stage.Splits_mapString, split, split.Name)
	return split
}

// StagePreserveOrder puts split to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SplitOrder
// - update stage.SplitOrder accordingly
func (split *Split) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Splits, stage.Split_stagedOrder, stage.Split_orderStaged, &stage.SplitOrder, stage.Splits_mapString, split, order, split.Name)
}

// Unstage removes split off the model stage
func (split *Split) Unstage(stage *Stage) *Split {
	__gong__unstage(stage.Splits, stage.Splits_mapString, split, split.Name)
	return split
}

// UnstageVoid removes split off the model stage
func (split *Split) UnstageVoid(stage *Stage) {
	split.Unstage(stage)
}

func (split *Split) StageVoid(stage *Stage) {
	split.Stage(stage)
}

// for satisfaction of GongStruct interface
func (split *Split) GetName() (res string) {
	return split.Name
}

// for satisfaction of GongStruct interface
func (split *Split) SetName(name string) {
	split.Name = name
}

// Stage puts svg to the model stage
func (svg *Svg) Stage(stage *Stage) *Svg {
	__gong__stage(stage.Svgs, stage.Svg_stagedOrder, stage.Svg_orderStaged, &stage.SvgOrder, stage.Svgs_mapString, svg, svg.Name)
	return svg
}

// StagePreserveOrder puts svg to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SvgOrder
// - update stage.SvgOrder accordingly
func (svg *Svg) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Svgs, stage.Svg_stagedOrder, stage.Svg_orderStaged, &stage.SvgOrder, stage.Svgs_mapString, svg, order, svg.Name)
}

// Unstage removes svg off the model stage
func (svg *Svg) Unstage(stage *Stage) *Svg {
	__gong__unstage(stage.Svgs, stage.Svgs_mapString, svg, svg.Name)
	return svg
}

// UnstageVoid removes svg off the model stage
func (svg *Svg) UnstageVoid(stage *Stage) {
	svg.Unstage(stage)
}

func (svg *Svg) StageVoid(stage *Stage) {
	svg.Stage(stage)
}

// for satisfaction of GongStruct interface
func (svg *Svg) GetName() (res string) {
	return svg.Name
}

// for satisfaction of GongStruct interface
func (svg *Svg) SetName(name string) {
	svg.Name = name
}

// Stage puts table to the model stage
func (table *Table) Stage(stage *Stage) *Table {
	__gong__stage(stage.Tables, stage.Table_stagedOrder, stage.Table_orderStaged, &stage.TableOrder, stage.Tables_mapString, table, table.Name)
	return table
}

// StagePreserveOrder puts table to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TableOrder
// - update stage.TableOrder accordingly
func (table *Table) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Tables, stage.Table_stagedOrder, stage.Table_orderStaged, &stage.TableOrder, stage.Tables_mapString, table, order, table.Name)
}

// Unstage removes table off the model stage
func (table *Table) Unstage(stage *Stage) *Table {
	__gong__unstage(stage.Tables, stage.Tables_mapString, table, table.Name)
	return table
}

// UnstageVoid removes table off the model stage
func (table *Table) UnstageVoid(stage *Stage) {
	table.Unstage(stage)
}

func (table *Table) StageVoid(stage *Stage) {
	table.Stage(stage)
}

// for satisfaction of GongStruct interface
func (table *Table) GetName() (res string) {
	return table.Name
}

// for satisfaction of GongStruct interface
func (table *Table) SetName(name string) {
	table.Name = name
}

// Stage puts threejs to the model stage
func (threejs *Threejs) Stage(stage *Stage) *Threejs {
	__gong__stage(stage.Threejss, stage.Threejs_stagedOrder, stage.Threejs_orderStaged, &stage.ThreejsOrder, stage.Threejss_mapString, threejs, threejs.Name)
	return threejs
}

// StagePreserveOrder puts threejs to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ThreejsOrder
// - update stage.ThreejsOrder accordingly
func (threejs *Threejs) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Threejss, stage.Threejs_stagedOrder, stage.Threejs_orderStaged, &stage.ThreejsOrder, stage.Threejss_mapString, threejs, order, threejs.Name)
}

// Unstage removes threejs off the model stage
func (threejs *Threejs) Unstage(stage *Stage) *Threejs {
	__gong__unstage(stage.Threejss, stage.Threejss_mapString, threejs, threejs.Name)
	return threejs
}

// UnstageVoid removes threejs off the model stage
func (threejs *Threejs) UnstageVoid(stage *Stage) {
	threejs.Unstage(stage)
}

func (threejs *Threejs) StageVoid(stage *Stage) {
	threejs.Stage(stage)
}

// for satisfaction of GongStruct interface
func (threejs *Threejs) GetName() (res string) {
	return threejs.Name
}

// for satisfaction of GongStruct interface
func (threejs *Threejs) SetName(name string) {
	threejs.Name = name
}

// Stage puts title to the model stage
func (title *Title) Stage(stage *Stage) *Title {
	__gong__stage(stage.Titles, stage.Title_stagedOrder, stage.Title_orderStaged, &stage.TitleOrder, stage.Titles_mapString, title, title.Name)
	return title
}

// StagePreserveOrder puts title to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TitleOrder
// - update stage.TitleOrder accordingly
func (title *Title) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Titles, stage.Title_stagedOrder, stage.Title_orderStaged, &stage.TitleOrder, stage.Titles_mapString, title, order, title.Name)
}

// Unstage removes title off the model stage
func (title *Title) Unstage(stage *Stage) *Title {
	__gong__unstage(stage.Titles, stage.Titles_mapString, title, title.Name)
	return title
}

// UnstageVoid removes title off the model stage
func (title *Title) UnstageVoid(stage *Stage) {
	title.Unstage(stage)
}

func (title *Title) StageVoid(stage *Stage) {
	title.Stage(stage)
}

// for satisfaction of GongStruct interface
func (title *Title) GetName() (res string) {
	return title.Name
}

// for satisfaction of GongStruct interface
func (title *Title) SetName(name string) {
	title.Name = name
}

// Stage puts tone to the model stage
func (tone *Tone) Stage(stage *Stage) *Tone {
	__gong__stage(stage.Tones, stage.Tone_stagedOrder, stage.Tone_orderStaged, &stage.ToneOrder, stage.Tones_mapString, tone, tone.Name)
	return tone
}

// StagePreserveOrder puts tone to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ToneOrder
// - update stage.ToneOrder accordingly
func (tone *Tone) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Tones, stage.Tone_stagedOrder, stage.Tone_orderStaged, &stage.ToneOrder, stage.Tones_mapString, tone, order, tone.Name)
}

// Unstage removes tone off the model stage
func (tone *Tone) Unstage(stage *Stage) *Tone {
	__gong__unstage(stage.Tones, stage.Tones_mapString, tone, tone.Name)
	return tone
}

// UnstageVoid removes tone off the model stage
func (tone *Tone) UnstageVoid(stage *Stage) {
	tone.Unstage(stage)
}

func (tone *Tone) StageVoid(stage *Stage) {
	tone.Stage(stage)
}

// for satisfaction of GongStruct interface
func (tone *Tone) GetName() (res string) {
	return tone.Name
}

// for satisfaction of GongStruct interface
func (tone *Tone) SetName(name string) {
	tone.Name = name
}

// Stage puts tree to the model stage
func (tree *Tree) Stage(stage *Stage) *Tree {
	__gong__stage(stage.Trees, stage.Tree_stagedOrder, stage.Tree_orderStaged, &stage.TreeOrder, stage.Trees_mapString, tree, tree.Name)
	return tree
}

// StagePreserveOrder puts tree to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TreeOrder
// - update stage.TreeOrder accordingly
func (tree *Tree) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Trees, stage.Tree_stagedOrder, stage.Tree_orderStaged, &stage.TreeOrder, stage.Trees_mapString, tree, order, tree.Name)
}

// Unstage removes tree off the model stage
func (tree *Tree) Unstage(stage *Stage) *Tree {
	__gong__unstage(stage.Trees, stage.Trees_mapString, tree, tree.Name)
	return tree
}

// UnstageVoid removes tree off the model stage
func (tree *Tree) UnstageVoid(stage *Stage) {
	tree.Unstage(stage)
}

func (tree *Tree) StageVoid(stage *Stage) {
	tree.Stage(stage)
}

// for satisfaction of GongStruct interface
func (tree *Tree) GetName() (res string) {
	return tree.Name
}

// for satisfaction of GongStruct interface
func (tree *Tree) SetName(name string) {
	tree.Name = name
}

// Stage puts view to the model stage
func (view *View) Stage(stage *Stage) *View {
	__gong__stage(stage.Views, stage.View_stagedOrder, stage.View_orderStaged, &stage.ViewOrder, stage.Views_mapString, view, view.Name)
	return view
}

// StagePreserveOrder puts view to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ViewOrder
// - update stage.ViewOrder accordingly
func (view *View) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Views, stage.View_stagedOrder, stage.View_orderStaged, &stage.ViewOrder, stage.Views_mapString, view, order, view.Name)
}

// Unstage removes view off the model stage
func (view *View) Unstage(stage *Stage) *View {
	__gong__unstage(stage.Views, stage.Views_mapString, view, view.Name)
	return view
}

// UnstageVoid removes view off the model stage
func (view *View) UnstageVoid(stage *Stage) {
	view.Unstage(stage)
}

func (view *View) StageVoid(stage *Stage) {
	view.Stage(stage)
}

// for satisfaction of GongStruct interface
func (view *View) GetName() (res string) {
	return view.Name
}

// for satisfaction of GongStruct interface
func (view *View) SetName(name string) {
	view.Name = name
}

// Stage puts xlsx to the model stage
func (xlsx *Xlsx) Stage(stage *Stage) *Xlsx {
	__gong__stage(stage.Xlsxs, stage.Xlsx_stagedOrder, stage.Xlsx_orderStaged, &stage.XlsxOrder, stage.Xlsxs_mapString, xlsx, xlsx.Name)
	return xlsx
}

// StagePreserveOrder puts xlsx to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.XlsxOrder
// - update stage.XlsxOrder accordingly
func (xlsx *Xlsx) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Xlsxs, stage.Xlsx_stagedOrder, stage.Xlsx_orderStaged, &stage.XlsxOrder, stage.Xlsxs_mapString, xlsx, order, xlsx.Name)
}

// Unstage removes xlsx off the model stage
func (xlsx *Xlsx) Unstage(stage *Stage) *Xlsx {
	__gong__unstage(stage.Xlsxs, stage.Xlsxs_mapString, xlsx, xlsx.Name)
	return xlsx
}

// UnstageVoid removes xlsx off the model stage
func (xlsx *Xlsx) UnstageVoid(stage *Stage) {
	xlsx.Unstage(stage)
}

func (xlsx *Xlsx) StageVoid(stage *Stage) {
	xlsx.Stage(stage)
}

// for satisfaction of GongStruct interface
func (xlsx *Xlsx) GetName() (res string) {
	return xlsx.Name
}

// for satisfaction of GongStruct interface
func (xlsx *Xlsx) SetName(name string) {
	xlsx.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.AsSplits, &stage.AsSplits_mapString, &stage.AsSplit_stagedOrder, &stage.AsSplitOrder)

	__gong__resetStageType(&stage.AsSplitAreas, &stage.AsSplitAreas_mapString, &stage.AsSplitArea_stagedOrder, &stage.AsSplitAreaOrder)

	__gong__resetStageType(&stage.Buttons, &stage.Buttons_mapString, &stage.Button_stagedOrder, &stage.ButtonOrder)

	__gong__resetStageType(&stage.Cursors, &stage.Cursors_mapString, &stage.Cursor_stagedOrder, &stage.CursorOrder)

	__gong__resetStageType(&stage.FavIcons, &stage.FavIcons_mapString, &stage.FavIcon_stagedOrder, &stage.FavIconOrder)

	__gong__resetStageType(&stage.Forms, &stage.Forms_mapString, &stage.Form_stagedOrder, &stage.FormOrder)

	__gong__resetStageType(&stage.Loads, &stage.Loads_mapString, &stage.Load_stagedOrder, &stage.LoadOrder)

	__gong__resetStageType(&stage.LogoOnTheLefts, &stage.LogoOnTheLefts_mapString, &stage.LogoOnTheLeft_stagedOrder, &stage.LogoOnTheLeftOrder)

	__gong__resetStageType(&stage.LogoOnTheRights, &stage.LogoOnTheRights_mapString, &stage.LogoOnTheRight_stagedOrder, &stage.LogoOnTheRightOrder)

	__gong__resetStageType(&stage.Markdowns, &stage.Markdowns_mapString, &stage.Markdown_stagedOrder, &stage.MarkdownOrder)

	__gong__resetStageType(&stage.Sliders, &stage.Sliders_mapString, &stage.Slider_stagedOrder, &stage.SliderOrder)

	__gong__resetStageType(&stage.Splits, &stage.Splits_mapString, &stage.Split_stagedOrder, &stage.SplitOrder)

	__gong__resetStageType(&stage.Svgs, &stage.Svgs_mapString, &stage.Svg_stagedOrder, &stage.SvgOrder)

	__gong__resetStageType(&stage.Tables, &stage.Tables_mapString, &stage.Table_stagedOrder, &stage.TableOrder)

	__gong__resetStageType(&stage.Threejss, &stage.Threejss_mapString, &stage.Threejs_stagedOrder, &stage.ThreejsOrder)

	__gong__resetStageType(&stage.Titles, &stage.Titles_mapString, &stage.Title_stagedOrder, &stage.TitleOrder)

	__gong__resetStageType(&stage.Tones, &stage.Tones_mapString, &stage.Tone_stagedOrder, &stage.ToneOrder)

	__gong__resetStageType(&stage.Trees, &stage.Trees_mapString, &stage.Tree_stagedOrder, &stage.TreeOrder)

	__gong__resetStageType(&stage.Views, &stage.Views_mapString, &stage.View_stagedOrder, &stage.ViewOrder)

	__gong__resetStageType(&stage.Xlsxs, &stage.Xlsxs_mapString, &stage.Xlsx_stagedOrder, &stage.XlsxOrder)

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
	case *AsSplit:
		return any(stage.AsSplits_mapString).(map[string]Type)
	case *AsSplitArea:
		return any(stage.AsSplitAreas_mapString).(map[string]Type)
	case *Button:
		return any(stage.Buttons_mapString).(map[string]Type)
	case *Cursor:
		return any(stage.Cursors_mapString).(map[string]Type)
	case *FavIcon:
		return any(stage.FavIcons_mapString).(map[string]Type)
	case *Form:
		return any(stage.Forms_mapString).(map[string]Type)
	case *Load:
		return any(stage.Loads_mapString).(map[string]Type)
	case *LogoOnTheLeft:
		return any(stage.LogoOnTheLefts_mapString).(map[string]Type)
	case *LogoOnTheRight:
		return any(stage.LogoOnTheRights_mapString).(map[string]Type)
	case *Markdown:
		return any(stage.Markdowns_mapString).(map[string]Type)
	case *Slider:
		return any(stage.Sliders_mapString).(map[string]Type)
	case *Split:
		return any(stage.Splits_mapString).(map[string]Type)
	case *Svg:
		return any(stage.Svgs_mapString).(map[string]Type)
	case *Table:
		return any(stage.Tables_mapString).(map[string]Type)
	case *Threejs:
		return any(stage.Threejss_mapString).(map[string]Type)
	case *Title:
		return any(stage.Titles_mapString).(map[string]Type)
	case *Tone:
		return any(stage.Tones_mapString).(map[string]Type)
	case *Tree:
		return any(stage.Trees_mapString).(map[string]Type)
	case *View:
		return any(stage.Views_mapString).(map[string]Type)
	case *Xlsx:
		return any(stage.Xlsxs_mapString).(map[string]Type)
	default:
		return nil
	}
}

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *AsSplit:
		return any(&stage.AsSplits).(*map[Type]struct{})
	case *AsSplitArea:
		return any(&stage.AsSplitAreas).(*map[Type]struct{})
	case *Button:
		return any(&stage.Buttons).(*map[Type]struct{})
	case *Cursor:
		return any(&stage.Cursors).(*map[Type]struct{})
	case *FavIcon:
		return any(&stage.FavIcons).(*map[Type]struct{})
	case *Form:
		return any(&stage.Forms).(*map[Type]struct{})
	case *Load:
		return any(&stage.Loads).(*map[Type]struct{})
	case *LogoOnTheLeft:
		return any(&stage.LogoOnTheLefts).(*map[Type]struct{})
	case *LogoOnTheRight:
		return any(&stage.LogoOnTheRights).(*map[Type]struct{})
	case *Markdown:
		return any(&stage.Markdowns).(*map[Type]struct{})
	case *Slider:
		return any(&stage.Sliders).(*map[Type]struct{})
	case *Split:
		return any(&stage.Splits).(*map[Type]struct{})
	case *Svg:
		return any(&stage.Svgs).(*map[Type]struct{})
	case *Table:
		return any(&stage.Tables).(*map[Type]struct{})
	case *Threejs:
		return any(&stage.Threejss).(*map[Type]struct{})
	case *Title:
		return any(&stage.Titles).(*map[Type]struct{})
	case *Tone:
		return any(&stage.Tones).(*map[Type]struct{})
	case *Tree:
		return any(&stage.Trees).(*map[Type]struct{})
	case *View:
		return any(&stage.Views).(*map[Type]struct{})
	case *Xlsx:
		return any(&stage.Xlsxs).(*map[Type]struct{})
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
	case AsSplit:
		return any(&AsSplit{
			AsSplitAreas: []*AsSplitArea{{Name: "AsSplitAreas"}},
		}).(*Type)
	case AsSplitArea:
		return any(&AsSplitArea{
			AsSplit: &AsSplit{Name: "AsSplit"},
			Button: &Button{Name: "Button"},
			Cursor: &Cursor{Name: "Cursor"},
			Form: &Form{Name: "Form"},
			Load: &Load{Name: "Load"},
			Markdown: &Markdown{Name: "Markdown"},
			Slider: &Slider{Name: "Slider"},
			Split: &Split{Name: "Split"},
			Svg: &Svg{Name: "Svg"},
			Table: &Table{Name: "Table"},
			Tone: &Tone{Name: "Tone"},
			Tree: &Tree{Name: "Tree"},
			Threejs: &Threejs{Name: "Threejs"},
			Xlsx: &Xlsx{Name: "Xlsx"},
		}).(*Type)
	case View:
		return any(&View{
			RootAsSplitAreas: []*AsSplitArea{{Name: "RootAsSplitAreas"}},
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
	// reverse maps of direct associations of AsSplit
	case AsSplit:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AsSplitArea
	case AsSplitArea:
		switch fieldname {
		// insertion point for per direct association field
		case "AsSplit":
			res := make(map[*AsSplit][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.AsSplit != nil {
					assplit_ := assplitarea.AsSplit
					var assplitareas []*AsSplitArea
					_, ok := res[assplit_]
					if ok {
						assplitareas = res[assplit_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[assplit_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Button":
			res := make(map[*Button][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Button != nil {
					button_ := assplitarea.Button
					var assplitareas []*AsSplitArea
					_, ok := res[button_]
					if ok {
						assplitareas = res[button_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[button_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Cursor":
			res := make(map[*Cursor][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Cursor != nil {
					cursor_ := assplitarea.Cursor
					var assplitareas []*AsSplitArea
					_, ok := res[cursor_]
					if ok {
						assplitareas = res[cursor_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[cursor_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Form":
			res := make(map[*Form][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Form != nil {
					form_ := assplitarea.Form
					var assplitareas []*AsSplitArea
					_, ok := res[form_]
					if ok {
						assplitareas = res[form_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[form_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Load":
			res := make(map[*Load][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Load != nil {
					load_ := assplitarea.Load
					var assplitareas []*AsSplitArea
					_, ok := res[load_]
					if ok {
						assplitareas = res[load_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[load_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Markdown":
			res := make(map[*Markdown][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Markdown != nil {
					markdown_ := assplitarea.Markdown
					var assplitareas []*AsSplitArea
					_, ok := res[markdown_]
					if ok {
						assplitareas = res[markdown_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[markdown_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Slider":
			res := make(map[*Slider][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Slider != nil {
					slider_ := assplitarea.Slider
					var assplitareas []*AsSplitArea
					_, ok := res[slider_]
					if ok {
						assplitareas = res[slider_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[slider_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Split":
			res := make(map[*Split][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Split != nil {
					split_ := assplitarea.Split
					var assplitareas []*AsSplitArea
					_, ok := res[split_]
					if ok {
						assplitareas = res[split_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[split_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Svg":
			res := make(map[*Svg][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Svg != nil {
					svg_ := assplitarea.Svg
					var assplitareas []*AsSplitArea
					_, ok := res[svg_]
					if ok {
						assplitareas = res[svg_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[svg_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Table":
			res := make(map[*Table][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Table != nil {
					table_ := assplitarea.Table
					var assplitareas []*AsSplitArea
					_, ok := res[table_]
					if ok {
						assplitareas = res[table_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[table_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Tone":
			res := make(map[*Tone][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Tone != nil {
					tone_ := assplitarea.Tone
					var assplitareas []*AsSplitArea
					_, ok := res[tone_]
					if ok {
						assplitareas = res[tone_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[tone_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Tree":
			res := make(map[*Tree][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Tree != nil {
					tree_ := assplitarea.Tree
					var assplitareas []*AsSplitArea
					_, ok := res[tree_]
					if ok {
						assplitareas = res[tree_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[tree_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Threejs":
			res := make(map[*Threejs][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Threejs != nil {
					threejs_ := assplitarea.Threejs
					var assplitareas []*AsSplitArea
					_, ok := res[threejs_]
					if ok {
						assplitareas = res[threejs_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[threejs_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Xlsx":
			res := make(map[*Xlsx][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Xlsx != nil {
					xlsx_ := assplitarea.Xlsx
					var assplitareas []*AsSplitArea
					_, ok := res[xlsx_]
					if ok {
						assplitareas = res[xlsx_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[xlsx_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Button
	case Button:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Cursor
	case Cursor:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FavIcon
	case FavIcon:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Form
	case Form:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Load
	case Load:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LogoOnTheLeft
	case LogoOnTheLeft:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LogoOnTheRight
	case LogoOnTheRight:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Markdown
	case Markdown:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Slider
	case Slider:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Split
	case Split:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Svg
	case Svg:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Table
	case Table:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Threejs
	case Threejs:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Title
	case Title:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tone
	case Tone:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tree
	case Tree:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of View
	case View:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xlsx
	case Xlsx:
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
	// reverse maps of direct associations of AsSplit
	case AsSplit:
		switch fieldname {
		// insertion point for per direct association field
		case "AsSplitAreas":
			res := make(map[*AsSplitArea][]*AsSplit)
			for assplit := range stage.AsSplits {
				for _, assplitarea_ := range assplit.AsSplitAreas {
					res[assplitarea_] = append(res[assplitarea_], assplit)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AsSplitArea
	case AsSplitArea:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Button
	case Button:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Cursor
	case Cursor:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FavIcon
	case FavIcon:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Form
	case Form:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Load
	case Load:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LogoOnTheLeft
	case LogoOnTheLeft:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LogoOnTheRight
	case LogoOnTheRight:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Markdown
	case Markdown:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Slider
	case Slider:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Split
	case Split:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Svg
	case Svg:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Table
	case Table:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Threejs
	case Threejs:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Title
	case Title:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tone
	case Tone:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tree
	case Tree:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of View
	case View:
		switch fieldname {
		// insertion point for per direct association field
		case "RootAsSplitAreas":
			res := make(map[*AsSplitArea][]*View)
			for view := range stage.Views {
				for _, assplitarea_ := range view.RootAsSplitAreas {
					res[assplitarea_] = append(res[assplitarea_], view)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Xlsx
	case Xlsx:
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
	case *AsSplit:
		res = any(new(AsSplit)).(Type)
	case *AsSplitArea:
		res = any(new(AsSplitArea)).(Type)
	case *Button:
		res = any(new(Button)).(Type)
	case *Cursor:
		res = any(new(Cursor)).(Type)
	case *FavIcon:
		res = any(new(FavIcon)).(Type)
	case *Form:
		res = any(new(Form)).(Type)
	case *Load:
		res = any(new(Load)).(Type)
	case *LogoOnTheLeft:
		res = any(new(LogoOnTheLeft)).(Type)
	case *LogoOnTheRight:
		res = any(new(LogoOnTheRight)).(Type)
	case *Markdown:
		res = any(new(Markdown)).(Type)
	case *Slider:
		res = any(new(Slider)).(Type)
	case *Split:
		res = any(new(Split)).(Type)
	case *Svg:
		res = any(new(Svg)).(Type)
	case *Table:
		res = any(new(Table)).(Type)
	case *Threejs:
		res = any(new(Threejs)).(Type)
	case *Title:
		res = any(new(Title)).(Type)
	case *Tone:
		res = any(new(Tone)).(Type)
	case *Tree:
		res = any(new(Tree)).(Type)
	case *View:
		res = any(new(View)).(Type)
	case *Xlsx:
		res = any(new(Xlsx)).(Type)
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
	case *AsSplit:
		res = "AsSplit"
	case *AsSplitArea:
		res = "AsSplitArea"
	case *Button:
		res = "Button"
	case *Cursor:
		res = "Cursor"
	case *FavIcon:
		res = "FavIcon"
	case *Form:
		res = "Form"
	case *Load:
		res = "Load"
	case *LogoOnTheLeft:
		res = "LogoOnTheLeft"
	case *LogoOnTheRight:
		res = "LogoOnTheRight"
	case *Markdown:
		res = "Markdown"
	case *Slider:
		res = "Slider"
	case *Split:
		res = "Split"
	case *Svg:
		res = "Svg"
	case *Table:
		res = "Table"
	case *Threejs:
		res = "Threejs"
	case *Title:
		res = "Title"
	case *Tone:
		res = "Tone"
	case *Tree:
		res = "Tree"
	case *View:
		res = "View"
	case *Xlsx:
		res = "Xlsx"
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
	case *AsSplit:
		var rf ReverseField
		_ = rf
	case *AsSplitArea:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "AsSplit"
		rf.Fieldname = "AsSplitAreas"
		res = append(res, rf)
		rf.GongstructName = "View"
		rf.Fieldname = "RootAsSplitAreas"
		res = append(res, rf)
	case *Button:
		var rf ReverseField
		_ = rf
	case *Cursor:
		var rf ReverseField
		_ = rf
	case *FavIcon:
		var rf ReverseField
		_ = rf
	case *Form:
		var rf ReverseField
		_ = rf
	case *Load:
		var rf ReverseField
		_ = rf
	case *LogoOnTheLeft:
		var rf ReverseField
		_ = rf
	case *LogoOnTheRight:
		var rf ReverseField
		_ = rf
	case *Markdown:
		var rf ReverseField
		_ = rf
	case *Slider:
		var rf ReverseField
		_ = rf
	case *Split:
		var rf ReverseField
		_ = rf
	case *Svg:
		var rf ReverseField
		_ = rf
	case *Table:
		var rf ReverseField
		_ = rf
	case *Threejs:
		var rf ReverseField
		_ = rf
	case *Title:
		var rf ReverseField
		_ = rf
	case *Tone:
		var rf ReverseField
		_ = rf
	case *Tree:
		var rf ReverseField
		_ = rf
	case *View:
		var rf ReverseField
		_ = rf
	case *Xlsx:
		var rf ReverseField
		_ = rf
	}
	return
}

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
}

// insertion point for get fields header method
func (assplit *AsSplit) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Direction",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "Direction",
		},
		{
			Name:                 "AsSplitAreas",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AsSplitArea",
		},
		{
			Name:               "IsSizeInPixel",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsWithCustomGutterSize",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "GutterSize",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (assplitarea *AsSplitArea) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ShowNameInHeader",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "Size",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsAny",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "AsSplit",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AsSplit",
		},
		{
			Name:                 "Button",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Button",
		},
		{
			Name:                 "Cursor",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Cursor",
		},
		{
			Name:                 "Form",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Form",
		},
		{
			Name:                 "Load",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Load",
		},
		{
			Name:                 "Markdown",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Markdown",
		},
		{
			Name:                 "Slider",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Slider",
		},
		{
			Name:                 "Split",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Split",
		},
		{
			Name:                 "Svg",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Svg",
		},
		{
			Name:                 "Table",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Table",
		},
		{
			Name:                 "Tone",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Tone",
		},
		{
			Name:                 "Tree",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Tree",
		},
		{
			Name:                 "Threejs",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Threejs",
		},
		{
			Name:                 "Xlsx",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Xlsx",
		},
		{
			Name:               "HasDiv",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "DivStyle",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (button *Button) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StackName",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (cursor *Cursor) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StackName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Style",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (favicon *FavIcon) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "SVG",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (form *Form) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StackName",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (load *Load) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StackName",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (logoontheleft *LogoOnTheLeft) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "SVG",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (logoontheright *LogoOnTheRight) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "SVG",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (markdown *Markdown) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StackName",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (slider *Slider) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StackName",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (split *Split) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StackName",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (svg *Svg) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StackName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Style",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (table *Table) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StackName",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (threejs *Threejs) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StackName",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (title *Title) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (tone *Tone) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StackName",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (tree *Tree) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StackName",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (view *View) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ShowViewName",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "RootAsSplitAreas",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AsSplitArea",
		},
		{
			Name:               "IsSelectedView",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Direction",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "Direction",
		},
		{
			Name:               "IsSecondaryView",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsSizeInPixel",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsWithCustomGutterSize",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "GutterSize",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (xlsx *Xlsx) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StackName",
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
func (assplit *AsSplit) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = assplit.Name
	case "Direction":
		enum := assplit.Direction
		res.valueString = enum.ToCodeString()
	case "AsSplitAreas":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range assplit.AsSplitAreas {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsSizeInPixel":
		res.valueString = fmt.Sprintf("%t", assplit.IsSizeInPixel)
		res.valueBool = assplit.IsSizeInPixel
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsWithCustomGutterSize":
		res.valueString = fmt.Sprintf("%t", assplit.IsWithCustomGutterSize)
		res.valueBool = assplit.IsWithCustomGutterSize
		res.GongFieldValueType = GongFieldValueTypeBool
	case "GutterSize":
		res.valueString = fmt.Sprintf("%f", assplit.GutterSize)
		res.valueFloat = assplit.GutterSize
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (assplitarea *AsSplitArea) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = assplitarea.Name
	case "ShowNameInHeader":
		res.valueString = fmt.Sprintf("%t", assplitarea.ShowNameInHeader)
		res.valueBool = assplitarea.ShowNameInHeader
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Size":
		res.valueString = fmt.Sprintf("%f", assplitarea.Size)
		res.valueFloat = assplitarea.Size
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsAny":
		res.valueString = fmt.Sprintf("%t", assplitarea.IsAny)
		res.valueBool = assplitarea.IsAny
		res.GongFieldValueType = GongFieldValueTypeBool
	case "AsSplit":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.AsSplit != nil {
			res.valueString = assplitarea.AsSplit.Name
			res.ids = assplitarea.AsSplit.GongGetUUID(stage)
		}
	case "Button":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Button != nil {
			res.valueString = assplitarea.Button.Name
			res.ids = assplitarea.Button.GongGetUUID(stage)
		}
	case "Cursor":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Cursor != nil {
			res.valueString = assplitarea.Cursor.Name
			res.ids = assplitarea.Cursor.GongGetUUID(stage)
		}
	case "Form":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Form != nil {
			res.valueString = assplitarea.Form.Name
			res.ids = assplitarea.Form.GongGetUUID(stage)
		}
	case "Load":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Load != nil {
			res.valueString = assplitarea.Load.Name
			res.ids = assplitarea.Load.GongGetUUID(stage)
		}
	case "Markdown":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Markdown != nil {
			res.valueString = assplitarea.Markdown.Name
			res.ids = assplitarea.Markdown.GongGetUUID(stage)
		}
	case "Slider":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Slider != nil {
			res.valueString = assplitarea.Slider.Name
			res.ids = assplitarea.Slider.GongGetUUID(stage)
		}
	case "Split":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Split != nil {
			res.valueString = assplitarea.Split.Name
			res.ids = assplitarea.Split.GongGetUUID(stage)
		}
	case "Svg":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Svg != nil {
			res.valueString = assplitarea.Svg.Name
			res.ids = assplitarea.Svg.GongGetUUID(stage)
		}
	case "Table":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Table != nil {
			res.valueString = assplitarea.Table.Name
			res.ids = assplitarea.Table.GongGetUUID(stage)
		}
	case "Tone":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Tone != nil {
			res.valueString = assplitarea.Tone.Name
			res.ids = assplitarea.Tone.GongGetUUID(stage)
		}
	case "Tree":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Tree != nil {
			res.valueString = assplitarea.Tree.Name
			res.ids = assplitarea.Tree.GongGetUUID(stage)
		}
	case "Threejs":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Threejs != nil {
			res.valueString = assplitarea.Threejs.Name
			res.ids = assplitarea.Threejs.GongGetUUID(stage)
		}
	case "Xlsx":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Xlsx != nil {
			res.valueString = assplitarea.Xlsx.Name
			res.ids = assplitarea.Xlsx.GongGetUUID(stage)
		}
	case "HasDiv":
		res.valueString = fmt.Sprintf("%t", assplitarea.HasDiv)
		res.valueBool = assplitarea.HasDiv
		res.GongFieldValueType = GongFieldValueTypeBool
	case "DivStyle":
		res.valueString = assplitarea.DivStyle
	}
	return
}

func (button *Button) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = button.Name
	case "StackName":
		res.valueString = button.StackName
	}
	return
}

func (cursor *Cursor) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = cursor.Name
	case "StackName":
		res.valueString = cursor.StackName
	case "Style":
		res.valueString = cursor.Style
	}
	return
}

func (favicon *FavIcon) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = favicon.Name
	case "SVG":
		res.valueString = favicon.SVG
	}
	return
}

func (form *Form) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = form.Name
	case "StackName":
		res.valueString = form.StackName
	}
	return
}

func (load *Load) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = load.Name
	case "StackName":
		res.valueString = load.StackName
	}
	return
}

func (logoontheleft *LogoOnTheLeft) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = logoontheleft.Name
	case "Width":
		res.valueString = fmt.Sprintf("%d", logoontheleft.Width)
		res.valueInt = logoontheleft.Width
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Height":
		res.valueString = fmt.Sprintf("%d", logoontheleft.Height)
		res.valueInt = logoontheleft.Height
		res.GongFieldValueType = GongFieldValueTypeInt
	case "SVG":
		res.valueString = logoontheleft.SVG
	}
	return
}

func (logoontheright *LogoOnTheRight) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = logoontheright.Name
	case "Width":
		res.valueString = fmt.Sprintf("%d", logoontheright.Width)
		res.valueInt = logoontheright.Width
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Height":
		res.valueString = fmt.Sprintf("%d", logoontheright.Height)
		res.valueInt = logoontheright.Height
		res.GongFieldValueType = GongFieldValueTypeInt
	case "SVG":
		res.valueString = logoontheright.SVG
	}
	return
}

func (markdown *Markdown) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = markdown.Name
	case "StackName":
		res.valueString = markdown.StackName
	}
	return
}

func (slider *Slider) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = slider.Name
	case "StackName":
		res.valueString = slider.StackName
	}
	return
}

func (split *Split) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = split.Name
	case "StackName":
		res.valueString = split.StackName
	}
	return
}

func (svg *Svg) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = svg.Name
	case "StackName":
		res.valueString = svg.StackName
	case "Style":
		res.valueString = svg.Style
	}
	return
}

func (table *Table) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = table.Name
	case "StackName":
		res.valueString = table.StackName
	}
	return
}

func (threejs *Threejs) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = threejs.Name
	case "StackName":
		res.valueString = threejs.StackName
	}
	return
}

func (title *Title) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = title.Name
	}
	return
}

func (tone *Tone) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = tone.Name
	case "StackName":
		res.valueString = tone.StackName
	}
	return
}

func (tree *Tree) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = tree.Name
	case "StackName":
		res.valueString = tree.StackName
	}
	return
}

func (view *View) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = view.Name
	case "ShowViewName":
		res.valueString = fmt.Sprintf("%t", view.ShowViewName)
		res.valueBool = view.ShowViewName
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RootAsSplitAreas":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range view.RootAsSplitAreas {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsSelectedView":
		res.valueString = fmt.Sprintf("%t", view.IsSelectedView)
		res.valueBool = view.IsSelectedView
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Direction":
		enum := view.Direction
		res.valueString = enum.ToCodeString()
	case "IsSecondaryView":
		res.valueString = fmt.Sprintf("%t", view.IsSecondaryView)
		res.valueBool = view.IsSecondaryView
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsSizeInPixel":
		res.valueString = fmt.Sprintf("%t", view.IsSizeInPixel)
		res.valueBool = view.IsSizeInPixel
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsWithCustomGutterSize":
		res.valueString = fmt.Sprintf("%t", view.IsWithCustomGutterSize)
		res.valueBool = view.IsWithCustomGutterSize
		res.GongFieldValueType = GongFieldValueTypeBool
	case "GutterSize":
		res.valueString = fmt.Sprintf("%f", view.GutterSize)
		res.valueFloat = view.GutterSize
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (xlsx *Xlsx) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = xlsx.Name
	case "StackName":
		res.valueString = xlsx.StackName
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
func (assplit *AsSplit) GongGetGongstructName() string {
	return "AsSplit"
}

func (assplitarea *AsSplitArea) GongGetGongstructName() string {
	return "AsSplitArea"
}

func (button *Button) GongGetGongstructName() string {
	return "Button"
}

func (cursor *Cursor) GongGetGongstructName() string {
	return "Cursor"
}

func (favicon *FavIcon) GongGetGongstructName() string {
	return "FavIcon"
}

func (form *Form) GongGetGongstructName() string {
	return "Form"
}

func (load *Load) GongGetGongstructName() string {
	return "Load"
}

func (logoontheleft *LogoOnTheLeft) GongGetGongstructName() string {
	return "LogoOnTheLeft"
}

func (logoontheright *LogoOnTheRight) GongGetGongstructName() string {
	return "LogoOnTheRight"
}

func (markdown *Markdown) GongGetGongstructName() string {
	return "Markdown"
}

func (slider *Slider) GongGetGongstructName() string {
	return "Slider"
}

func (split *Split) GongGetGongstructName() string {
	return "Split"
}

func (svg *Svg) GongGetGongstructName() string {
	return "Svg"
}

func (table *Table) GongGetGongstructName() string {
	return "Table"
}

func (threejs *Threejs) GongGetGongstructName() string {
	return "Threejs"
}

func (title *Title) GongGetGongstructName() string {
	return "Title"
}

func (tone *Tone) GongGetGongstructName() string {
	return "Tone"
}

func (tree *Tree) GongGetGongstructName() string {
	return "Tree"
}

func (view *View) GongGetGongstructName() string {
	return "View"
}

func (xlsx *Xlsx) GongGetGongstructName() string {
	return "Xlsx"
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
	__gong__rebuildMapString(stage.AsSplits, &stage.AsSplits_mapString)

	__gong__rebuildMapString(stage.AsSplitAreas, &stage.AsSplitAreas_mapString)

	__gong__rebuildMapString(stage.Buttons, &stage.Buttons_mapString)

	__gong__rebuildMapString(stage.Cursors, &stage.Cursors_mapString)

	__gong__rebuildMapString(stage.FavIcons, &stage.FavIcons_mapString)

	__gong__rebuildMapString(stage.Forms, &stage.Forms_mapString)

	__gong__rebuildMapString(stage.Loads, &stage.Loads_mapString)

	__gong__rebuildMapString(stage.LogoOnTheLefts, &stage.LogoOnTheLefts_mapString)

	__gong__rebuildMapString(stage.LogoOnTheRights, &stage.LogoOnTheRights_mapString)

	__gong__rebuildMapString(stage.Markdowns, &stage.Markdowns_mapString)

	__gong__rebuildMapString(stage.Sliders, &stage.Sliders_mapString)

	__gong__rebuildMapString(stage.Splits, &stage.Splits_mapString)

	__gong__rebuildMapString(stage.Svgs, &stage.Svgs_mapString)

	__gong__rebuildMapString(stage.Tables, &stage.Tables_mapString)

	__gong__rebuildMapString(stage.Threejss, &stage.Threejss_mapString)

	__gong__rebuildMapString(stage.Titles, &stage.Titles_mapString)

	__gong__rebuildMapString(stage.Tones, &stage.Tones_mapString)

	__gong__rebuildMapString(stage.Trees, &stage.Trees_mapString)

	__gong__rebuildMapString(stage.Views, &stage.Views_mapString)

	__gong__rebuildMapString(stage.Xlsxs, &stage.Xlsxs_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
