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

	split_go "github.com/fullstack-lang/gong/lib/split/go"
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
	ProbeTreeSidebarSuffix       = ":sidebar of the probe"
	ProbeTableSuffix             = ":table of the probe"
	ProbeNotificationTableSuffix = ":notification table of the probe"
	ProbeFormSuffix              = ":form of the probe"
	ProbeSplitSuffix             = ":probe of the probe"
)

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTreeSidebarSuffix
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

	// insertion point for definition of arrays registering instances
	AsSplits                map[*AsSplit]struct{}
	AsSplits_instance       map[*AsSplit]*AsSplit
	AsSplits_mapString      map[string]*AsSplit
	AsSplitOrder            uint
	AsSplit_stagedOrder     map[*AsSplit]uint
	AsSplits_reference      map[*AsSplit]*AsSplit
	AsSplits_referenceOrder map[*AsSplit]uint
	
	// insertion point for slice of pointers maps
	AsSplit_AsSplitAreas_reverseMap map[*AsSplitArea]*AsSplit

	OnAfterAsSplitCreateCallback OnAfterCreateInterface[AsSplit]
	OnAfterAsSplitUpdateCallback OnAfterUpdateInterface[AsSplit]
	OnAfterAsSplitDeleteCallback OnAfterDeleteInterface[AsSplit]
	OnAfterAsSplitReadCallback   OnAfterReadInterface[AsSplit]

	AsSplitAreas                map[*AsSplitArea]struct{}
	AsSplitAreas_instance       map[*AsSplitArea]*AsSplitArea
	AsSplitAreas_mapString      map[string]*AsSplitArea
	AsSplitAreaOrder            uint
	AsSplitArea_stagedOrder     map[*AsSplitArea]uint
	AsSplitAreas_reference      map[*AsSplitArea]*AsSplitArea
	AsSplitAreas_referenceOrder map[*AsSplitArea]uint
	
	// insertion point for slice of pointers maps
	OnAfterAsSplitAreaCreateCallback OnAfterCreateInterface[AsSplitArea]
	OnAfterAsSplitAreaUpdateCallback OnAfterUpdateInterface[AsSplitArea]
	OnAfterAsSplitAreaDeleteCallback OnAfterDeleteInterface[AsSplitArea]
	OnAfterAsSplitAreaReadCallback   OnAfterReadInterface[AsSplitArea]

	Buttons                map[*Button]struct{}
	Buttons_instance       map[*Button]*Button
	Buttons_mapString      map[string]*Button
	ButtonOrder            uint
	Button_stagedOrder     map[*Button]uint
	Buttons_reference      map[*Button]*Button
	Buttons_referenceOrder map[*Button]uint
	
	// insertion point for slice of pointers maps
	OnAfterButtonCreateCallback OnAfterCreateInterface[Button]
	OnAfterButtonUpdateCallback OnAfterUpdateInterface[Button]
	OnAfterButtonDeleteCallback OnAfterDeleteInterface[Button]
	OnAfterButtonReadCallback   OnAfterReadInterface[Button]

	Cursors                map[*Cursor]struct{}
	Cursors_instance       map[*Cursor]*Cursor
	Cursors_mapString      map[string]*Cursor
	CursorOrder            uint
	Cursor_stagedOrder     map[*Cursor]uint
	Cursors_reference      map[*Cursor]*Cursor
	Cursors_referenceOrder map[*Cursor]uint
	
	// insertion point for slice of pointers maps
	OnAfterCursorCreateCallback OnAfterCreateInterface[Cursor]
	OnAfterCursorUpdateCallback OnAfterUpdateInterface[Cursor]
	OnAfterCursorDeleteCallback OnAfterDeleteInterface[Cursor]
	OnAfterCursorReadCallback   OnAfterReadInterface[Cursor]

	FavIcons                map[*FavIcon]struct{}
	FavIcons_instance       map[*FavIcon]*FavIcon
	FavIcons_mapString      map[string]*FavIcon
	FavIconOrder            uint
	FavIcon_stagedOrder     map[*FavIcon]uint
	FavIcons_reference      map[*FavIcon]*FavIcon
	FavIcons_referenceOrder map[*FavIcon]uint
	
	// insertion point for slice of pointers maps
	OnAfterFavIconCreateCallback OnAfterCreateInterface[FavIcon]
	OnAfterFavIconUpdateCallback OnAfterUpdateInterface[FavIcon]
	OnAfterFavIconDeleteCallback OnAfterDeleteInterface[FavIcon]
	OnAfterFavIconReadCallback   OnAfterReadInterface[FavIcon]

	Forms                map[*Form]struct{}
	Forms_instance       map[*Form]*Form
	Forms_mapString      map[string]*Form
	FormOrder            uint
	Form_stagedOrder     map[*Form]uint
	Forms_reference      map[*Form]*Form
	Forms_referenceOrder map[*Form]uint
	
	// insertion point for slice of pointers maps
	OnAfterFormCreateCallback OnAfterCreateInterface[Form]
	OnAfterFormUpdateCallback OnAfterUpdateInterface[Form]
	OnAfterFormDeleteCallback OnAfterDeleteInterface[Form]
	OnAfterFormReadCallback   OnAfterReadInterface[Form]

	Loads                map[*Load]struct{}
	Loads_instance       map[*Load]*Load
	Loads_mapString      map[string]*Load
	LoadOrder            uint
	Load_stagedOrder     map[*Load]uint
	Loads_reference      map[*Load]*Load
	Loads_referenceOrder map[*Load]uint
	
	// insertion point for slice of pointers maps
	OnAfterLoadCreateCallback OnAfterCreateInterface[Load]
	OnAfterLoadUpdateCallback OnAfterUpdateInterface[Load]
	OnAfterLoadDeleteCallback OnAfterDeleteInterface[Load]
	OnAfterLoadReadCallback   OnAfterReadInterface[Load]

	LogoOnTheLefts                map[*LogoOnTheLeft]struct{}
	LogoOnTheLefts_instance       map[*LogoOnTheLeft]*LogoOnTheLeft
	LogoOnTheLefts_mapString      map[string]*LogoOnTheLeft
	LogoOnTheLeftOrder            uint
	LogoOnTheLeft_stagedOrder     map[*LogoOnTheLeft]uint
	LogoOnTheLefts_reference      map[*LogoOnTheLeft]*LogoOnTheLeft
	LogoOnTheLefts_referenceOrder map[*LogoOnTheLeft]uint
	
	// insertion point for slice of pointers maps
	OnAfterLogoOnTheLeftCreateCallback OnAfterCreateInterface[LogoOnTheLeft]
	OnAfterLogoOnTheLeftUpdateCallback OnAfterUpdateInterface[LogoOnTheLeft]
	OnAfterLogoOnTheLeftDeleteCallback OnAfterDeleteInterface[LogoOnTheLeft]
	OnAfterLogoOnTheLeftReadCallback   OnAfterReadInterface[LogoOnTheLeft]

	LogoOnTheRights                map[*LogoOnTheRight]struct{}
	LogoOnTheRights_instance       map[*LogoOnTheRight]*LogoOnTheRight
	LogoOnTheRights_mapString      map[string]*LogoOnTheRight
	LogoOnTheRightOrder            uint
	LogoOnTheRight_stagedOrder     map[*LogoOnTheRight]uint
	LogoOnTheRights_reference      map[*LogoOnTheRight]*LogoOnTheRight
	LogoOnTheRights_referenceOrder map[*LogoOnTheRight]uint
	
	// insertion point for slice of pointers maps
	OnAfterLogoOnTheRightCreateCallback OnAfterCreateInterface[LogoOnTheRight]
	OnAfterLogoOnTheRightUpdateCallback OnAfterUpdateInterface[LogoOnTheRight]
	OnAfterLogoOnTheRightDeleteCallback OnAfterDeleteInterface[LogoOnTheRight]
	OnAfterLogoOnTheRightReadCallback   OnAfterReadInterface[LogoOnTheRight]

	Markdowns                map[*Markdown]struct{}
	Markdowns_instance       map[*Markdown]*Markdown
	Markdowns_mapString      map[string]*Markdown
	MarkdownOrder            uint
	Markdown_stagedOrder     map[*Markdown]uint
	Markdowns_reference      map[*Markdown]*Markdown
	Markdowns_referenceOrder map[*Markdown]uint
	
	// insertion point for slice of pointers maps
	OnAfterMarkdownCreateCallback OnAfterCreateInterface[Markdown]
	OnAfterMarkdownUpdateCallback OnAfterUpdateInterface[Markdown]
	OnAfterMarkdownDeleteCallback OnAfterDeleteInterface[Markdown]
	OnAfterMarkdownReadCallback   OnAfterReadInterface[Markdown]

	Sliders                map[*Slider]struct{}
	Sliders_instance       map[*Slider]*Slider
	Sliders_mapString      map[string]*Slider
	SliderOrder            uint
	Slider_stagedOrder     map[*Slider]uint
	Sliders_reference      map[*Slider]*Slider
	Sliders_referenceOrder map[*Slider]uint
	
	// insertion point for slice of pointers maps
	OnAfterSliderCreateCallback OnAfterCreateInterface[Slider]
	OnAfterSliderUpdateCallback OnAfterUpdateInterface[Slider]
	OnAfterSliderDeleteCallback OnAfterDeleteInterface[Slider]
	OnAfterSliderReadCallback   OnAfterReadInterface[Slider]

	Splits                map[*Split]struct{}
	Splits_instance       map[*Split]*Split
	Splits_mapString      map[string]*Split
	SplitOrder            uint
	Split_stagedOrder     map[*Split]uint
	Splits_reference      map[*Split]*Split
	Splits_referenceOrder map[*Split]uint
	
	// insertion point for slice of pointers maps
	OnAfterSplitCreateCallback OnAfterCreateInterface[Split]
	OnAfterSplitUpdateCallback OnAfterUpdateInterface[Split]
	OnAfterSplitDeleteCallback OnAfterDeleteInterface[Split]
	OnAfterSplitReadCallback   OnAfterReadInterface[Split]

	Svgs                map[*Svg]struct{}
	Svgs_instance       map[*Svg]*Svg
	Svgs_mapString      map[string]*Svg
	SvgOrder            uint
	Svg_stagedOrder     map[*Svg]uint
	Svgs_reference      map[*Svg]*Svg
	Svgs_referenceOrder map[*Svg]uint
	
	// insertion point for slice of pointers maps
	OnAfterSvgCreateCallback OnAfterCreateInterface[Svg]
	OnAfterSvgUpdateCallback OnAfterUpdateInterface[Svg]
	OnAfterSvgDeleteCallback OnAfterDeleteInterface[Svg]
	OnAfterSvgReadCallback   OnAfterReadInterface[Svg]

	Tables                map[*Table]struct{}
	Tables_instance       map[*Table]*Table
	Tables_mapString      map[string]*Table
	TableOrder            uint
	Table_stagedOrder     map[*Table]uint
	Tables_reference      map[*Table]*Table
	Tables_referenceOrder map[*Table]uint
	
	// insertion point for slice of pointers maps
	OnAfterTableCreateCallback OnAfterCreateInterface[Table]
	OnAfterTableUpdateCallback OnAfterUpdateInterface[Table]
	OnAfterTableDeleteCallback OnAfterDeleteInterface[Table]
	OnAfterTableReadCallback   OnAfterReadInterface[Table]

	Titles                map[*Title]struct{}
	Titles_instance       map[*Title]*Title
	Titles_mapString      map[string]*Title
	TitleOrder            uint
	Title_stagedOrder     map[*Title]uint
	Titles_reference      map[*Title]*Title
	Titles_referenceOrder map[*Title]uint
	
	// insertion point for slice of pointers maps
	OnAfterTitleCreateCallback OnAfterCreateInterface[Title]
	OnAfterTitleUpdateCallback OnAfterUpdateInterface[Title]
	OnAfterTitleDeleteCallback OnAfterDeleteInterface[Title]
	OnAfterTitleReadCallback   OnAfterReadInterface[Title]

	Tones                map[*Tone]struct{}
	Tones_instance       map[*Tone]*Tone
	Tones_mapString      map[string]*Tone
	ToneOrder            uint
	Tone_stagedOrder     map[*Tone]uint
	Tones_reference      map[*Tone]*Tone
	Tones_referenceOrder map[*Tone]uint
	
	// insertion point for slice of pointers maps
	OnAfterToneCreateCallback OnAfterCreateInterface[Tone]
	OnAfterToneUpdateCallback OnAfterUpdateInterface[Tone]
	OnAfterToneDeleteCallback OnAfterDeleteInterface[Tone]
	OnAfterToneReadCallback   OnAfterReadInterface[Tone]

	Trees                map[*Tree]struct{}
	Trees_instance       map[*Tree]*Tree
	Trees_mapString      map[string]*Tree
	TreeOrder            uint
	Tree_stagedOrder     map[*Tree]uint
	Trees_reference      map[*Tree]*Tree
	Trees_referenceOrder map[*Tree]uint
	
	// insertion point for slice of pointers maps
	OnAfterTreeCreateCallback OnAfterCreateInterface[Tree]
	OnAfterTreeUpdateCallback OnAfterUpdateInterface[Tree]
	OnAfterTreeDeleteCallback OnAfterDeleteInterface[Tree]
	OnAfterTreeReadCallback   OnAfterReadInterface[Tree]

	Views                map[*View]struct{}
	Views_instance       map[*View]*View
	Views_mapString      map[string]*View
	ViewOrder            uint
	View_stagedOrder     map[*View]uint
	Views_reference      map[*View]*View
	Views_referenceOrder map[*View]uint
	
	// insertion point for slice of pointers maps
	View_RootAsSplitAreas_reverseMap map[*AsSplitArea]*View

	OnAfterViewCreateCallback OnAfterCreateInterface[View]
	OnAfterViewUpdateCallback OnAfterUpdateInterface[View]
	OnAfterViewDeleteCallback OnAfterDeleteInterface[View]
	OnAfterViewReadCallback   OnAfterReadInterface[View]

	Xlsxs                map[*Xlsx]struct{}
	Xlsxs_instance       map[*Xlsx]*Xlsx
	Xlsxs_mapString      map[string]*Xlsx
	XlsxOrder            uint
	Xlsx_stagedOrder     map[*Xlsx]uint
	Xlsxs_reference      map[*Xlsx]*Xlsx
	Xlsxs_referenceOrder map[*Xlsx]uint
	
	// insertion point for slice of pointers maps
	OnAfterXlsxCreateCallback OnAfterCreateInterface[Xlsx]
	OnAfterXlsxUpdateCallback OnAfterUpdateInterface[Xlsx]
	OnAfterXlsxDeleteCallback OnAfterDeleteInterface[Xlsx]
	OnAfterXlsxReadCallback   OnAfterReadInterface[Xlsx]

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

	lock sync.RWMutex
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

	if stage.commitsBehind >= len(stage.backwardCommits) {
		return errors.New("no more backward commit to apply")
	}

	commitToApply := stage.backwardCommits[len(stage.backwardCommits)-1-stage.commitsBehind]

	// umarshall the backward commit to the stage

	// the parsing of the commit will call the UX update
	// therefore, it is important to stage.commitsBehind before because it is used in the
	// UX
	stage.commitsBehind++
	err := GongParseAstString(stage, commitToApply, true)
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
	err := GongParseAstString(stage, commitToApply, true)
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

// Orphans removes all commits
func (stage *Stage) Orphans() {
	stage.forwardCommits = stage.forwardCommits[:0]
	stage.backwardCommits = stage.backwardCommits[:0]
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

// recomputeOrders recomputes the next order for each struct
// this is necessary because the order might have been incremented
// during the commits that have been discarded
// insertion point for max order recomputation
func (stage *Stage) recomputeOrders() {
	// insertion point for max order recomputation
	var maxAsSplitOrder uint
	var foundAsSplit bool
	for _, order := range stage.AsSplit_stagedOrder {
		if !foundAsSplit || order > maxAsSplitOrder {
			maxAsSplitOrder = order
			foundAsSplit = true
		}
	}
	if foundAsSplit {
		stage.AsSplitOrder = maxAsSplitOrder + 1
	} else {
		stage.AsSplitOrder = 0
	}

	var maxAsSplitAreaOrder uint
	var foundAsSplitArea bool
	for _, order := range stage.AsSplitArea_stagedOrder {
		if !foundAsSplitArea || order > maxAsSplitAreaOrder {
			maxAsSplitAreaOrder = order
			foundAsSplitArea = true
		}
	}
	if foundAsSplitArea {
		stage.AsSplitAreaOrder = maxAsSplitAreaOrder + 1
	} else {
		stage.AsSplitAreaOrder = 0
	}

	var maxButtonOrder uint
	var foundButton bool
	for _, order := range stage.Button_stagedOrder {
		if !foundButton || order > maxButtonOrder {
			maxButtonOrder = order
			foundButton = true
		}
	}
	if foundButton {
		stage.ButtonOrder = maxButtonOrder + 1
	} else {
		stage.ButtonOrder = 0
	}

	var maxCursorOrder uint
	var foundCursor bool
	for _, order := range stage.Cursor_stagedOrder {
		if !foundCursor || order > maxCursorOrder {
			maxCursorOrder = order
			foundCursor = true
		}
	}
	if foundCursor {
		stage.CursorOrder = maxCursorOrder + 1
	} else {
		stage.CursorOrder = 0
	}

	var maxFavIconOrder uint
	var foundFavIcon bool
	for _, order := range stage.FavIcon_stagedOrder {
		if !foundFavIcon || order > maxFavIconOrder {
			maxFavIconOrder = order
			foundFavIcon = true
		}
	}
	if foundFavIcon {
		stage.FavIconOrder = maxFavIconOrder + 1
	} else {
		stage.FavIconOrder = 0
	}

	var maxFormOrder uint
	var foundForm bool
	for _, order := range stage.Form_stagedOrder {
		if !foundForm || order > maxFormOrder {
			maxFormOrder = order
			foundForm = true
		}
	}
	if foundForm {
		stage.FormOrder = maxFormOrder + 1
	} else {
		stage.FormOrder = 0
	}

	var maxLoadOrder uint
	var foundLoad bool
	for _, order := range stage.Load_stagedOrder {
		if !foundLoad || order > maxLoadOrder {
			maxLoadOrder = order
			foundLoad = true
		}
	}
	if foundLoad {
		stage.LoadOrder = maxLoadOrder + 1
	} else {
		stage.LoadOrder = 0
	}

	var maxLogoOnTheLeftOrder uint
	var foundLogoOnTheLeft bool
	for _, order := range stage.LogoOnTheLeft_stagedOrder {
		if !foundLogoOnTheLeft || order > maxLogoOnTheLeftOrder {
			maxLogoOnTheLeftOrder = order
			foundLogoOnTheLeft = true
		}
	}
	if foundLogoOnTheLeft {
		stage.LogoOnTheLeftOrder = maxLogoOnTheLeftOrder + 1
	} else {
		stage.LogoOnTheLeftOrder = 0
	}

	var maxLogoOnTheRightOrder uint
	var foundLogoOnTheRight bool
	for _, order := range stage.LogoOnTheRight_stagedOrder {
		if !foundLogoOnTheRight || order > maxLogoOnTheRightOrder {
			maxLogoOnTheRightOrder = order
			foundLogoOnTheRight = true
		}
	}
	if foundLogoOnTheRight {
		stage.LogoOnTheRightOrder = maxLogoOnTheRightOrder + 1
	} else {
		stage.LogoOnTheRightOrder = 0
	}

	var maxMarkdownOrder uint
	var foundMarkdown bool
	for _, order := range stage.Markdown_stagedOrder {
		if !foundMarkdown || order > maxMarkdownOrder {
			maxMarkdownOrder = order
			foundMarkdown = true
		}
	}
	if foundMarkdown {
		stage.MarkdownOrder = maxMarkdownOrder + 1
	} else {
		stage.MarkdownOrder = 0
	}

	var maxSliderOrder uint
	var foundSlider bool
	for _, order := range stage.Slider_stagedOrder {
		if !foundSlider || order > maxSliderOrder {
			maxSliderOrder = order
			foundSlider = true
		}
	}
	if foundSlider {
		stage.SliderOrder = maxSliderOrder + 1
	} else {
		stage.SliderOrder = 0
	}

	var maxSplitOrder uint
	var foundSplit bool
	for _, order := range stage.Split_stagedOrder {
		if !foundSplit || order > maxSplitOrder {
			maxSplitOrder = order
			foundSplit = true
		}
	}
	if foundSplit {
		stage.SplitOrder = maxSplitOrder + 1
	} else {
		stage.SplitOrder = 0
	}

	var maxSvgOrder uint
	var foundSvg bool
	for _, order := range stage.Svg_stagedOrder {
		if !foundSvg || order > maxSvgOrder {
			maxSvgOrder = order
			foundSvg = true
		}
	}
	if foundSvg {
		stage.SvgOrder = maxSvgOrder + 1
	} else {
		stage.SvgOrder = 0
	}

	var maxTableOrder uint
	var foundTable bool
	for _, order := range stage.Table_stagedOrder {
		if !foundTable || order > maxTableOrder {
			maxTableOrder = order
			foundTable = true
		}
	}
	if foundTable {
		stage.TableOrder = maxTableOrder + 1
	} else {
		stage.TableOrder = 0
	}

	var maxTitleOrder uint
	var foundTitle bool
	for _, order := range stage.Title_stagedOrder {
		if !foundTitle || order > maxTitleOrder {
			maxTitleOrder = order
			foundTitle = true
		}
	}
	if foundTitle {
		stage.TitleOrder = maxTitleOrder + 1
	} else {
		stage.TitleOrder = 0
	}

	var maxToneOrder uint
	var foundTone bool
	for _, order := range stage.Tone_stagedOrder {
		if !foundTone || order > maxToneOrder {
			maxToneOrder = order
			foundTone = true
		}
	}
	if foundTone {
		stage.ToneOrder = maxToneOrder + 1
	} else {
		stage.ToneOrder = 0
	}

	var maxTreeOrder uint
	var foundTree bool
	for _, order := range stage.Tree_stagedOrder {
		if !foundTree || order > maxTreeOrder {
			maxTreeOrder = order
			foundTree = true
		}
	}
	if foundTree {
		stage.TreeOrder = maxTreeOrder + 1
	} else {
		stage.TreeOrder = 0
	}

	var maxViewOrder uint
	var foundView bool
	for _, order := range stage.View_stagedOrder {
		if !foundView || order > maxViewOrder {
			maxViewOrder = order
			foundView = true
		}
	}
	if foundView {
		stage.ViewOrder = maxViewOrder + 1
	} else {
		stage.ViewOrder = 0
	}

	var maxXlsxOrder uint
	var foundXlsx bool
	for _, order := range stage.Xlsx_stagedOrder {
		if !foundXlsx || order > maxXlsxOrder {
			maxXlsxOrder = order
			foundXlsx = true
		}
	}
	if foundXlsx {
		stage.XlsxOrder = maxXlsxOrder + 1
	} else {
		stage.XlsxOrder = 0
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

func GetStructInstancesByOrderAuto[T PointerToGongstruct](stage *Stage) (res []T) {
	var t T
	switch any(t).(type) {
	// insertion point for case
	case *AsSplit:
		tmp := GetStructInstancesByOrder(stage.AsSplits, stage.AsSplit_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *AsSplit implements.
			res = append(res, any(v).(T))
		}
		return res
	case *AsSplitArea:
		tmp := GetStructInstancesByOrder(stage.AsSplitAreas, stage.AsSplitArea_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *AsSplitArea implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Button:
		tmp := GetStructInstancesByOrder(stage.Buttons, stage.Button_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Button implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Cursor:
		tmp := GetStructInstancesByOrder(stage.Cursors, stage.Cursor_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Cursor implements.
			res = append(res, any(v).(T))
		}
		return res
	case *FavIcon:
		tmp := GetStructInstancesByOrder(stage.FavIcons, stage.FavIcon_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *FavIcon implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Form:
		tmp := GetStructInstancesByOrder(stage.Forms, stage.Form_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Form implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Load:
		tmp := GetStructInstancesByOrder(stage.Loads, stage.Load_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Load implements.
			res = append(res, any(v).(T))
		}
		return res
	case *LogoOnTheLeft:
		tmp := GetStructInstancesByOrder(stage.LogoOnTheLefts, stage.LogoOnTheLeft_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *LogoOnTheLeft implements.
			res = append(res, any(v).(T))
		}
		return res
	case *LogoOnTheRight:
		tmp := GetStructInstancesByOrder(stage.LogoOnTheRights, stage.LogoOnTheRight_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *LogoOnTheRight implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Markdown:
		tmp := GetStructInstancesByOrder(stage.Markdowns, stage.Markdown_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Markdown implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Slider:
		tmp := GetStructInstancesByOrder(stage.Sliders, stage.Slider_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Slider implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Split:
		tmp := GetStructInstancesByOrder(stage.Splits, stage.Split_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Split implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Svg:
		tmp := GetStructInstancesByOrder(stage.Svgs, stage.Svg_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Svg implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Table:
		tmp := GetStructInstancesByOrder(stage.Tables, stage.Table_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Table implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Title:
		tmp := GetStructInstancesByOrder(stage.Titles, stage.Title_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Title implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Tone:
		tmp := GetStructInstancesByOrder(stage.Tones, stage.Tone_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Tone implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Tree:
		tmp := GetStructInstancesByOrder(stage.Trees, stage.Tree_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Tree implements.
			res = append(res, any(v).(T))
		}
		return res
	case *View:
		tmp := GetStructInstancesByOrder(stage.Views, stage.View_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *View implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Xlsx:
		tmp := GetStructInstancesByOrder(stage.Xlsxs, stage.Xlsx_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Xlsx implements.
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
	case "AsSplit":
		res = GetNamedStructInstances(stage.AsSplits, stage.AsSplit_stagedOrder)
	case "AsSplitArea":
		res = GetNamedStructInstances(stage.AsSplitAreas, stage.AsSplitArea_stagedOrder)
	case "Button":
		res = GetNamedStructInstances(stage.Buttons, stage.Button_stagedOrder)
	case "Cursor":
		res = GetNamedStructInstances(stage.Cursors, stage.Cursor_stagedOrder)
	case "FavIcon":
		res = GetNamedStructInstances(stage.FavIcons, stage.FavIcon_stagedOrder)
	case "Form":
		res = GetNamedStructInstances(stage.Forms, stage.Form_stagedOrder)
	case "Load":
		res = GetNamedStructInstances(stage.Loads, stage.Load_stagedOrder)
	case "LogoOnTheLeft":
		res = GetNamedStructInstances(stage.LogoOnTheLefts, stage.LogoOnTheLeft_stagedOrder)
	case "LogoOnTheRight":
		res = GetNamedStructInstances(stage.LogoOnTheRights, stage.LogoOnTheRight_stagedOrder)
	case "Markdown":
		res = GetNamedStructInstances(stage.Markdowns, stage.Markdown_stagedOrder)
	case "Slider":
		res = GetNamedStructInstances(stage.Sliders, stage.Slider_stagedOrder)
	case "Split":
		res = GetNamedStructInstances(stage.Splits, stage.Split_stagedOrder)
	case "Svg":
		res = GetNamedStructInstances(stage.Svgs, stage.Svg_stagedOrder)
	case "Table":
		res = GetNamedStructInstances(stage.Tables, stage.Table_stagedOrder)
	case "Title":
		res = GetNamedStructInstances(stage.Titles, stage.Title_stagedOrder)
	case "Tone":
		res = GetNamedStructInstances(stage.Tones, stage.Tone_stagedOrder)
	case "Tree":
		res = GetNamedStructInstances(stage.Trees, stage.Tree_stagedOrder)
	case "View":
		res = GetNamedStructInstances(stage.Views, stage.View_stagedOrder)
	case "Xlsx":
		res = GetNamedStructInstances(stage.Xlsxs, stage.Xlsx_stagedOrder)
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
	return "github.com/fullstack-lang/gong/lib/split/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return split_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return split_go.GoDiagramsDir
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
	CommitAsSplit(assplit *AsSplit)
	CheckoutAsSplit(assplit *AsSplit)
	CommitAsSplitArea(assplitarea *AsSplitArea)
	CheckoutAsSplitArea(assplitarea *AsSplitArea)
	CommitButton(button *Button)
	CheckoutButton(button *Button)
	CommitCursor(cursor *Cursor)
	CheckoutCursor(cursor *Cursor)
	CommitFavIcon(favicon *FavIcon)
	CheckoutFavIcon(favicon *FavIcon)
	CommitForm(form *Form)
	CheckoutForm(form *Form)
	CommitLoad(load *Load)
	CheckoutLoad(load *Load)
	CommitLogoOnTheLeft(logoontheleft *LogoOnTheLeft)
	CheckoutLogoOnTheLeft(logoontheleft *LogoOnTheLeft)
	CommitLogoOnTheRight(logoontheright *LogoOnTheRight)
	CheckoutLogoOnTheRight(logoontheright *LogoOnTheRight)
	CommitMarkdown(markdown *Markdown)
	CheckoutMarkdown(markdown *Markdown)
	CommitSlider(slider *Slider)
	CheckoutSlider(slider *Slider)
	CommitSplit(split *Split)
	CheckoutSplit(split *Split)
	CommitSvg(svg *Svg)
	CheckoutSvg(svg *Svg)
	CommitTable(table *Table)
	CheckoutTable(table *Table)
	CommitTitle(title *Title)
	CheckoutTitle(title *Title)
	CommitTone(tone *Tone)
	CheckoutTone(tone *Tone)
	CommitTree(tree *Tree)
	CheckoutTree(tree *Tree)
	CommitView(view *View)
	CheckoutView(view *View)
	CommitXlsx(xlsx *Xlsx)
	CheckoutXlsx(xlsx *Xlsx)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

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

		AsSplitArea_stagedOrder: make(map[*AsSplitArea]uint),

		Button_stagedOrder: make(map[*Button]uint),

		Cursor_stagedOrder: make(map[*Cursor]uint),

		FavIcon_stagedOrder: make(map[*FavIcon]uint),

		Form_stagedOrder: make(map[*Form]uint),

		Load_stagedOrder: make(map[*Load]uint),

		LogoOnTheLeft_stagedOrder: make(map[*LogoOnTheLeft]uint),

		LogoOnTheRight_stagedOrder: make(map[*LogoOnTheRight]uint),

		Markdown_stagedOrder: make(map[*Markdown]uint),

		Slider_stagedOrder: make(map[*Slider]uint),

		Split_stagedOrder: make(map[*Split]uint),

		Svg_stagedOrder: make(map[*Svg]uint),

		Table_stagedOrder: make(map[*Table]uint),

		Title_stagedOrder: make(map[*Title]uint),

		Tone_stagedOrder: make(map[*Tone]uint),

		Tree_stagedOrder: make(map[*Tree]uint),

		View_stagedOrder: make(map[*View]uint),

		Xlsx_stagedOrder: make(map[*Xlsx]uint),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
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

			"Title": &TitleUnmarshaller{},

			"Tone": &ToneUnmarshaller{},

			"Tree": &TreeUnmarshaller{},

			"View": &ViewUnmarshaller{},

			"Xlsx": &XlsxUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "AsSplit"},
			{name: "AsSplitArea"},
			{name: "Button"},
			{name: "Cursor"},
			{name: "FavIcon"},
			{name: "Form"},
			{name: "Load"},
			{name: "LogoOnTheLeft"},
			{name: "LogoOnTheRight"},
			{name: "Markdown"},
			{name: "Slider"},
			{name: "Split"},
			{name: "Svg"},
			{name: "Table"},
			{name: "Title"},
			{name: "Tone"},
			{name: "Tree"},
			{name: "View"},
			{name: "Xlsx"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *AsSplit:
		return stage.AsSplit_stagedOrder[instance]
	case *AsSplitArea:
		return stage.AsSplitArea_stagedOrder[instance]
	case *Button:
		return stage.Button_stagedOrder[instance]
	case *Cursor:
		return stage.Cursor_stagedOrder[instance]
	case *FavIcon:
		return stage.FavIcon_stagedOrder[instance]
	case *Form:
		return stage.Form_stagedOrder[instance]
	case *Load:
		return stage.Load_stagedOrder[instance]
	case *LogoOnTheLeft:
		return stage.LogoOnTheLeft_stagedOrder[instance]
	case *LogoOnTheRight:
		return stage.LogoOnTheRight_stagedOrder[instance]
	case *Markdown:
		return stage.Markdown_stagedOrder[instance]
	case *Slider:
		return stage.Slider_stagedOrder[instance]
	case *Split:
		return stage.Split_stagedOrder[instance]
	case *Svg:
		return stage.Svg_stagedOrder[instance]
	case *Table:
		return stage.Table_stagedOrder[instance]
	case *Title:
		return stage.Title_stagedOrder[instance]
	case *Tone:
		return stage.Tone_stagedOrder[instance]
	case *Tree:
		return stage.Tree_stagedOrder[instance]
	case *View:
		return stage.View_stagedOrder[instance]
	case *Xlsx:
		return stage.Xlsx_stagedOrder[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *AsSplit:
		return stage.AsSplit_stagedOrder[instance]
	case *AsSplitArea:
		return stage.AsSplitArea_stagedOrder[instance]
	case *Button:
		return stage.Button_stagedOrder[instance]
	case *Cursor:
		return stage.Cursor_stagedOrder[instance]
	case *FavIcon:
		return stage.FavIcon_stagedOrder[instance]
	case *Form:
		return stage.Form_stagedOrder[instance]
	case *Load:
		return stage.Load_stagedOrder[instance]
	case *LogoOnTheLeft:
		return stage.LogoOnTheLeft_stagedOrder[instance]
	case *LogoOnTheRight:
		return stage.LogoOnTheRight_stagedOrder[instance]
	case *Markdown:
		return stage.Markdown_stagedOrder[instance]
	case *Slider:
		return stage.Slider_stagedOrder[instance]
	case *Split:
		return stage.Split_stagedOrder[instance]
	case *Svg:
		return stage.Svg_stagedOrder[instance]
	case *Table:
		return stage.Table_stagedOrder[instance]
	case *Title:
		return stage.Title_stagedOrder[instance]
	case *Tone:
		return stage.Tone_stagedOrder[instance]
	case *Tree:
		return stage.Tree_stagedOrder[instance]
	case *View:
		return stage.View_stagedOrder[instance]
	case *Xlsx:
		return stage.Xlsx_stagedOrder[instance]
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
	if _, ok := stage.AsSplits[assplit]; !ok {
		stage.AsSplits[assplit] = struct{}{}
		stage.AsSplit_stagedOrder[assplit] = stage.AsSplitOrder
		stage.AsSplitOrder++
	}
	stage.AsSplits_mapString[assplit.Name] = assplit

	return assplit
}

// StagePreserveOrder puts assplit to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AsSplitOrder
// - update stage.AsSplitOrder accordingly
func (assplit *AsSplit) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.AsSplits[assplit]; !ok {
		stage.AsSplits[assplit] = struct{}{}

		if order > stage.AsSplitOrder {
			stage.AsSplitOrder = order
		}
		stage.AsSplit_stagedOrder[assplit] = order
		stage.AsSplitOrder++
	}
	stage.AsSplits_mapString[assplit.Name] = assplit
}

// Unstage removes assplit off the model stage
func (assplit *AsSplit) Unstage(stage *Stage) *AsSplit {
	delete(stage.AsSplits, assplit)
	// issue1150
	// delete(stage.AsSplit_stagedOrder, assplit)
	delete(stage.AsSplits_mapString, assplit.Name)

	return assplit
}

// UnstageVoid removes assplit off the model stage
func (assplit *AsSplit) UnstageVoid(stage *Stage) {
	delete(stage.AsSplits, assplit)
	// issue1150
	// delete(stage.AsSplit_stagedOrder, assplit)
	delete(stage.AsSplits_mapString, assplit.Name)
}

// commit assplit to the back repo (if it is already staged)
func (assplit *AsSplit) Commit(stage *Stage) *AsSplit {
	if _, ok := stage.AsSplits[assplit]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAsSplit(assplit)
		}
	}
	return assplit
}

func (assplit *AsSplit) CommitVoid(stage *Stage) {
	assplit.Commit(stage)
}

func (assplit *AsSplit) StageVoid(stage *Stage) {
	assplit.Stage(stage)
}

// Checkout assplit to the back repo (if it is already staged)
func (assplit *AsSplit) Checkout(stage *Stage) *AsSplit {
	if _, ok := stage.AsSplits[assplit]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAsSplit(assplit)
		}
	}
	return assplit
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
	if _, ok := stage.AsSplitAreas[assplitarea]; !ok {
		stage.AsSplitAreas[assplitarea] = struct{}{}
		stage.AsSplitArea_stagedOrder[assplitarea] = stage.AsSplitAreaOrder
		stage.AsSplitAreaOrder++
	}
	stage.AsSplitAreas_mapString[assplitarea.Name] = assplitarea

	return assplitarea
}

// StagePreserveOrder puts assplitarea to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AsSplitAreaOrder
// - update stage.AsSplitAreaOrder accordingly
func (assplitarea *AsSplitArea) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.AsSplitAreas[assplitarea]; !ok {
		stage.AsSplitAreas[assplitarea] = struct{}{}

		if order > stage.AsSplitAreaOrder {
			stage.AsSplitAreaOrder = order
		}
		stage.AsSplitArea_stagedOrder[assplitarea] = order
		stage.AsSplitAreaOrder++
	}
	stage.AsSplitAreas_mapString[assplitarea.Name] = assplitarea
}

// Unstage removes assplitarea off the model stage
func (assplitarea *AsSplitArea) Unstage(stage *Stage) *AsSplitArea {
	delete(stage.AsSplitAreas, assplitarea)
	// issue1150
	// delete(stage.AsSplitArea_stagedOrder, assplitarea)
	delete(stage.AsSplitAreas_mapString, assplitarea.Name)

	return assplitarea
}

// UnstageVoid removes assplitarea off the model stage
func (assplitarea *AsSplitArea) UnstageVoid(stage *Stage) {
	delete(stage.AsSplitAreas, assplitarea)
	// issue1150
	// delete(stage.AsSplitArea_stagedOrder, assplitarea)
	delete(stage.AsSplitAreas_mapString, assplitarea.Name)
}

// commit assplitarea to the back repo (if it is already staged)
func (assplitarea *AsSplitArea) Commit(stage *Stage) *AsSplitArea {
	if _, ok := stage.AsSplitAreas[assplitarea]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAsSplitArea(assplitarea)
		}
	}
	return assplitarea
}

func (assplitarea *AsSplitArea) CommitVoid(stage *Stage) {
	assplitarea.Commit(stage)
}

func (assplitarea *AsSplitArea) StageVoid(stage *Stage) {
	assplitarea.Stage(stage)
}

// Checkout assplitarea to the back repo (if it is already staged)
func (assplitarea *AsSplitArea) Checkout(stage *Stage) *AsSplitArea {
	if _, ok := stage.AsSplitAreas[assplitarea]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAsSplitArea(assplitarea)
		}
	}
	return assplitarea
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
	if _, ok := stage.Buttons[button]; !ok {
		stage.Buttons[button] = struct{}{}
		stage.Button_stagedOrder[button] = stage.ButtonOrder
		stage.ButtonOrder++
	}
	stage.Buttons_mapString[button.Name] = button

	return button
}

// StagePreserveOrder puts button to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ButtonOrder
// - update stage.ButtonOrder accordingly
func (button *Button) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Buttons[button]; !ok {
		stage.Buttons[button] = struct{}{}

		if order > stage.ButtonOrder {
			stage.ButtonOrder = order
		}
		stage.Button_stagedOrder[button] = order
		stage.ButtonOrder++
	}
	stage.Buttons_mapString[button.Name] = button
}

// Unstage removes button off the model stage
func (button *Button) Unstage(stage *Stage) *Button {
	delete(stage.Buttons, button)
	// issue1150
	// delete(stage.Button_stagedOrder, button)
	delete(stage.Buttons_mapString, button.Name)

	return button
}

// UnstageVoid removes button off the model stage
func (button *Button) UnstageVoid(stage *Stage) {
	delete(stage.Buttons, button)
	// issue1150
	// delete(stage.Button_stagedOrder, button)
	delete(stage.Buttons_mapString, button.Name)
}

// commit button to the back repo (if it is already staged)
func (button *Button) Commit(stage *Stage) *Button {
	if _, ok := stage.Buttons[button]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitButton(button)
		}
	}
	return button
}

func (button *Button) CommitVoid(stage *Stage) {
	button.Commit(stage)
}

func (button *Button) StageVoid(stage *Stage) {
	button.Stage(stage)
}

// Checkout button to the back repo (if it is already staged)
func (button *Button) Checkout(stage *Stage) *Button {
	if _, ok := stage.Buttons[button]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutButton(button)
		}
	}
	return button
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
	if _, ok := stage.Cursors[cursor]; !ok {
		stage.Cursors[cursor] = struct{}{}
		stage.Cursor_stagedOrder[cursor] = stage.CursorOrder
		stage.CursorOrder++
	}
	stage.Cursors_mapString[cursor.Name] = cursor

	return cursor
}

// StagePreserveOrder puts cursor to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CursorOrder
// - update stage.CursorOrder accordingly
func (cursor *Cursor) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Cursors[cursor]; !ok {
		stage.Cursors[cursor] = struct{}{}

		if order > stage.CursorOrder {
			stage.CursorOrder = order
		}
		stage.Cursor_stagedOrder[cursor] = order
		stage.CursorOrder++
	}
	stage.Cursors_mapString[cursor.Name] = cursor
}

// Unstage removes cursor off the model stage
func (cursor *Cursor) Unstage(stage *Stage) *Cursor {
	delete(stage.Cursors, cursor)
	// issue1150
	// delete(stage.Cursor_stagedOrder, cursor)
	delete(stage.Cursors_mapString, cursor.Name)

	return cursor
}

// UnstageVoid removes cursor off the model stage
func (cursor *Cursor) UnstageVoid(stage *Stage) {
	delete(stage.Cursors, cursor)
	// issue1150
	// delete(stage.Cursor_stagedOrder, cursor)
	delete(stage.Cursors_mapString, cursor.Name)
}

// commit cursor to the back repo (if it is already staged)
func (cursor *Cursor) Commit(stage *Stage) *Cursor {
	if _, ok := stage.Cursors[cursor]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCursor(cursor)
		}
	}
	return cursor
}

func (cursor *Cursor) CommitVoid(stage *Stage) {
	cursor.Commit(stage)
}

func (cursor *Cursor) StageVoid(stage *Stage) {
	cursor.Stage(stage)
}

// Checkout cursor to the back repo (if it is already staged)
func (cursor *Cursor) Checkout(stage *Stage) *Cursor {
	if _, ok := stage.Cursors[cursor]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCursor(cursor)
		}
	}
	return cursor
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
	if _, ok := stage.FavIcons[favicon]; !ok {
		stage.FavIcons[favicon] = struct{}{}
		stage.FavIcon_stagedOrder[favicon] = stage.FavIconOrder
		stage.FavIconOrder++
	}
	stage.FavIcons_mapString[favicon.Name] = favicon

	return favicon
}

// StagePreserveOrder puts favicon to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FavIconOrder
// - update stage.FavIconOrder accordingly
func (favicon *FavIcon) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.FavIcons[favicon]; !ok {
		stage.FavIcons[favicon] = struct{}{}

		if order > stage.FavIconOrder {
			stage.FavIconOrder = order
		}
		stage.FavIcon_stagedOrder[favicon] = order
		stage.FavIconOrder++
	}
	stage.FavIcons_mapString[favicon.Name] = favicon
}

// Unstage removes favicon off the model stage
func (favicon *FavIcon) Unstage(stage *Stage) *FavIcon {
	delete(stage.FavIcons, favicon)
	// issue1150
	// delete(stage.FavIcon_stagedOrder, favicon)
	delete(stage.FavIcons_mapString, favicon.Name)

	return favicon
}

// UnstageVoid removes favicon off the model stage
func (favicon *FavIcon) UnstageVoid(stage *Stage) {
	delete(stage.FavIcons, favicon)
	// issue1150
	// delete(stage.FavIcon_stagedOrder, favicon)
	delete(stage.FavIcons_mapString, favicon.Name)
}

// commit favicon to the back repo (if it is already staged)
func (favicon *FavIcon) Commit(stage *Stage) *FavIcon {
	if _, ok := stage.FavIcons[favicon]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFavIcon(favicon)
		}
	}
	return favicon
}

func (favicon *FavIcon) CommitVoid(stage *Stage) {
	favicon.Commit(stage)
}

func (favicon *FavIcon) StageVoid(stage *Stage) {
	favicon.Stage(stage)
}

// Checkout favicon to the back repo (if it is already staged)
func (favicon *FavIcon) Checkout(stage *Stage) *FavIcon {
	if _, ok := stage.FavIcons[favicon]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFavIcon(favicon)
		}
	}
	return favicon
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
	if _, ok := stage.Forms[form]; !ok {
		stage.Forms[form] = struct{}{}
		stage.Form_stagedOrder[form] = stage.FormOrder
		stage.FormOrder++
	}
	stage.Forms_mapString[form.Name] = form

	return form
}

// StagePreserveOrder puts form to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormOrder
// - update stage.FormOrder accordingly
func (form *Form) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Forms[form]; !ok {
		stage.Forms[form] = struct{}{}

		if order > stage.FormOrder {
			stage.FormOrder = order
		}
		stage.Form_stagedOrder[form] = order
		stage.FormOrder++
	}
	stage.Forms_mapString[form.Name] = form
}

// Unstage removes form off the model stage
func (form *Form) Unstage(stage *Stage) *Form {
	delete(stage.Forms, form)
	// issue1150
	// delete(stage.Form_stagedOrder, form)
	delete(stage.Forms_mapString, form.Name)

	return form
}

// UnstageVoid removes form off the model stage
func (form *Form) UnstageVoid(stage *Stage) {
	delete(stage.Forms, form)
	// issue1150
	// delete(stage.Form_stagedOrder, form)
	delete(stage.Forms_mapString, form.Name)
}

// commit form to the back repo (if it is already staged)
func (form *Form) Commit(stage *Stage) *Form {
	if _, ok := stage.Forms[form]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitForm(form)
		}
	}
	return form
}

func (form *Form) CommitVoid(stage *Stage) {
	form.Commit(stage)
}

func (form *Form) StageVoid(stage *Stage) {
	form.Stage(stage)
}

// Checkout form to the back repo (if it is already staged)
func (form *Form) Checkout(stage *Stage) *Form {
	if _, ok := stage.Forms[form]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutForm(form)
		}
	}
	return form
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
	if _, ok := stage.Loads[load]; !ok {
		stage.Loads[load] = struct{}{}
		stage.Load_stagedOrder[load] = stage.LoadOrder
		stage.LoadOrder++
	}
	stage.Loads_mapString[load.Name] = load

	return load
}

// StagePreserveOrder puts load to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LoadOrder
// - update stage.LoadOrder accordingly
func (load *Load) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Loads[load]; !ok {
		stage.Loads[load] = struct{}{}

		if order > stage.LoadOrder {
			stage.LoadOrder = order
		}
		stage.Load_stagedOrder[load] = order
		stage.LoadOrder++
	}
	stage.Loads_mapString[load.Name] = load
}

// Unstage removes load off the model stage
func (load *Load) Unstage(stage *Stage) *Load {
	delete(stage.Loads, load)
	// issue1150
	// delete(stage.Load_stagedOrder, load)
	delete(stage.Loads_mapString, load.Name)

	return load
}

// UnstageVoid removes load off the model stage
func (load *Load) UnstageVoid(stage *Stage) {
	delete(stage.Loads, load)
	// issue1150
	// delete(stage.Load_stagedOrder, load)
	delete(stage.Loads_mapString, load.Name)
}

// commit load to the back repo (if it is already staged)
func (load *Load) Commit(stage *Stage) *Load {
	if _, ok := stage.Loads[load]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLoad(load)
		}
	}
	return load
}

func (load *Load) CommitVoid(stage *Stage) {
	load.Commit(stage)
}

func (load *Load) StageVoid(stage *Stage) {
	load.Stage(stage)
}

// Checkout load to the back repo (if it is already staged)
func (load *Load) Checkout(stage *Stage) *Load {
	if _, ok := stage.Loads[load]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLoad(load)
		}
	}
	return load
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
	if _, ok := stage.LogoOnTheLefts[logoontheleft]; !ok {
		stage.LogoOnTheLefts[logoontheleft] = struct{}{}
		stage.LogoOnTheLeft_stagedOrder[logoontheleft] = stage.LogoOnTheLeftOrder
		stage.LogoOnTheLeftOrder++
	}
	stage.LogoOnTheLefts_mapString[logoontheleft.Name] = logoontheleft

	return logoontheleft
}

// StagePreserveOrder puts logoontheleft to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LogoOnTheLeftOrder
// - update stage.LogoOnTheLeftOrder accordingly
func (logoontheleft *LogoOnTheLeft) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.LogoOnTheLefts[logoontheleft]; !ok {
		stage.LogoOnTheLefts[logoontheleft] = struct{}{}

		if order > stage.LogoOnTheLeftOrder {
			stage.LogoOnTheLeftOrder = order
		}
		stage.LogoOnTheLeft_stagedOrder[logoontheleft] = order
		stage.LogoOnTheLeftOrder++
	}
	stage.LogoOnTheLefts_mapString[logoontheleft.Name] = logoontheleft
}

// Unstage removes logoontheleft off the model stage
func (logoontheleft *LogoOnTheLeft) Unstage(stage *Stage) *LogoOnTheLeft {
	delete(stage.LogoOnTheLefts, logoontheleft)
	// issue1150
	// delete(stage.LogoOnTheLeft_stagedOrder, logoontheleft)
	delete(stage.LogoOnTheLefts_mapString, logoontheleft.Name)

	return logoontheleft
}

// UnstageVoid removes logoontheleft off the model stage
func (logoontheleft *LogoOnTheLeft) UnstageVoid(stage *Stage) {
	delete(stage.LogoOnTheLefts, logoontheleft)
	// issue1150
	// delete(stage.LogoOnTheLeft_stagedOrder, logoontheleft)
	delete(stage.LogoOnTheLefts_mapString, logoontheleft.Name)
}

// commit logoontheleft to the back repo (if it is already staged)
func (logoontheleft *LogoOnTheLeft) Commit(stage *Stage) *LogoOnTheLeft {
	if _, ok := stage.LogoOnTheLefts[logoontheleft]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLogoOnTheLeft(logoontheleft)
		}
	}
	return logoontheleft
}

func (logoontheleft *LogoOnTheLeft) CommitVoid(stage *Stage) {
	logoontheleft.Commit(stage)
}

func (logoontheleft *LogoOnTheLeft) StageVoid(stage *Stage) {
	logoontheleft.Stage(stage)
}

// Checkout logoontheleft to the back repo (if it is already staged)
func (logoontheleft *LogoOnTheLeft) Checkout(stage *Stage) *LogoOnTheLeft {
	if _, ok := stage.LogoOnTheLefts[logoontheleft]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLogoOnTheLeft(logoontheleft)
		}
	}
	return logoontheleft
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
	if _, ok := stage.LogoOnTheRights[logoontheright]; !ok {
		stage.LogoOnTheRights[logoontheright] = struct{}{}
		stage.LogoOnTheRight_stagedOrder[logoontheright] = stage.LogoOnTheRightOrder
		stage.LogoOnTheRightOrder++
	}
	stage.LogoOnTheRights_mapString[logoontheright.Name] = logoontheright

	return logoontheright
}

// StagePreserveOrder puts logoontheright to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LogoOnTheRightOrder
// - update stage.LogoOnTheRightOrder accordingly
func (logoontheright *LogoOnTheRight) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.LogoOnTheRights[logoontheright]; !ok {
		stage.LogoOnTheRights[logoontheright] = struct{}{}

		if order > stage.LogoOnTheRightOrder {
			stage.LogoOnTheRightOrder = order
		}
		stage.LogoOnTheRight_stagedOrder[logoontheright] = order
		stage.LogoOnTheRightOrder++
	}
	stage.LogoOnTheRights_mapString[logoontheright.Name] = logoontheright
}

// Unstage removes logoontheright off the model stage
func (logoontheright *LogoOnTheRight) Unstage(stage *Stage) *LogoOnTheRight {
	delete(stage.LogoOnTheRights, logoontheright)
	// issue1150
	// delete(stage.LogoOnTheRight_stagedOrder, logoontheright)
	delete(stage.LogoOnTheRights_mapString, logoontheright.Name)

	return logoontheright
}

// UnstageVoid removes logoontheright off the model stage
func (logoontheright *LogoOnTheRight) UnstageVoid(stage *Stage) {
	delete(stage.LogoOnTheRights, logoontheright)
	// issue1150
	// delete(stage.LogoOnTheRight_stagedOrder, logoontheright)
	delete(stage.LogoOnTheRights_mapString, logoontheright.Name)
}

// commit logoontheright to the back repo (if it is already staged)
func (logoontheright *LogoOnTheRight) Commit(stage *Stage) *LogoOnTheRight {
	if _, ok := stage.LogoOnTheRights[logoontheright]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLogoOnTheRight(logoontheright)
		}
	}
	return logoontheright
}

func (logoontheright *LogoOnTheRight) CommitVoid(stage *Stage) {
	logoontheright.Commit(stage)
}

func (logoontheright *LogoOnTheRight) StageVoid(stage *Stage) {
	logoontheright.Stage(stage)
}

// Checkout logoontheright to the back repo (if it is already staged)
func (logoontheright *LogoOnTheRight) Checkout(stage *Stage) *LogoOnTheRight {
	if _, ok := stage.LogoOnTheRights[logoontheright]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLogoOnTheRight(logoontheright)
		}
	}
	return logoontheright
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
	if _, ok := stage.Markdowns[markdown]; !ok {
		stage.Markdowns[markdown] = struct{}{}
		stage.Markdown_stagedOrder[markdown] = stage.MarkdownOrder
		stage.MarkdownOrder++
	}
	stage.Markdowns_mapString[markdown.Name] = markdown

	return markdown
}

// StagePreserveOrder puts markdown to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MarkdownOrder
// - update stage.MarkdownOrder accordingly
func (markdown *Markdown) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Markdowns[markdown]; !ok {
		stage.Markdowns[markdown] = struct{}{}

		if order > stage.MarkdownOrder {
			stage.MarkdownOrder = order
		}
		stage.Markdown_stagedOrder[markdown] = order
		stage.MarkdownOrder++
	}
	stage.Markdowns_mapString[markdown.Name] = markdown
}

// Unstage removes markdown off the model stage
func (markdown *Markdown) Unstage(stage *Stage) *Markdown {
	delete(stage.Markdowns, markdown)
	// issue1150
	// delete(stage.Markdown_stagedOrder, markdown)
	delete(stage.Markdowns_mapString, markdown.Name)

	return markdown
}

// UnstageVoid removes markdown off the model stage
func (markdown *Markdown) UnstageVoid(stage *Stage) {
	delete(stage.Markdowns, markdown)
	// issue1150
	// delete(stage.Markdown_stagedOrder, markdown)
	delete(stage.Markdowns_mapString, markdown.Name)
}

// commit markdown to the back repo (if it is already staged)
func (markdown *Markdown) Commit(stage *Stage) *Markdown {
	if _, ok := stage.Markdowns[markdown]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMarkdown(markdown)
		}
	}
	return markdown
}

func (markdown *Markdown) CommitVoid(stage *Stage) {
	markdown.Commit(stage)
}

func (markdown *Markdown) StageVoid(stage *Stage) {
	markdown.Stage(stage)
}

// Checkout markdown to the back repo (if it is already staged)
func (markdown *Markdown) Checkout(stage *Stage) *Markdown {
	if _, ok := stage.Markdowns[markdown]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMarkdown(markdown)
		}
	}
	return markdown
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
	if _, ok := stage.Sliders[slider]; !ok {
		stage.Sliders[slider] = struct{}{}
		stage.Slider_stagedOrder[slider] = stage.SliderOrder
		stage.SliderOrder++
	}
	stage.Sliders_mapString[slider.Name] = slider

	return slider
}

// StagePreserveOrder puts slider to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SliderOrder
// - update stage.SliderOrder accordingly
func (slider *Slider) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Sliders[slider]; !ok {
		stage.Sliders[slider] = struct{}{}

		if order > stage.SliderOrder {
			stage.SliderOrder = order
		}
		stage.Slider_stagedOrder[slider] = order
		stage.SliderOrder++
	}
	stage.Sliders_mapString[slider.Name] = slider
}

// Unstage removes slider off the model stage
func (slider *Slider) Unstage(stage *Stage) *Slider {
	delete(stage.Sliders, slider)
	// issue1150
	// delete(stage.Slider_stagedOrder, slider)
	delete(stage.Sliders_mapString, slider.Name)

	return slider
}

// UnstageVoid removes slider off the model stage
func (slider *Slider) UnstageVoid(stage *Stage) {
	delete(stage.Sliders, slider)
	// issue1150
	// delete(stage.Slider_stagedOrder, slider)
	delete(stage.Sliders_mapString, slider.Name)
}

// commit slider to the back repo (if it is already staged)
func (slider *Slider) Commit(stage *Stage) *Slider {
	if _, ok := stage.Sliders[slider]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSlider(slider)
		}
	}
	return slider
}

func (slider *Slider) CommitVoid(stage *Stage) {
	slider.Commit(stage)
}

func (slider *Slider) StageVoid(stage *Stage) {
	slider.Stage(stage)
}

// Checkout slider to the back repo (if it is already staged)
func (slider *Slider) Checkout(stage *Stage) *Slider {
	if _, ok := stage.Sliders[slider]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSlider(slider)
		}
	}
	return slider
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
	if _, ok := stage.Splits[split]; !ok {
		stage.Splits[split] = struct{}{}
		stage.Split_stagedOrder[split] = stage.SplitOrder
		stage.SplitOrder++
	}
	stage.Splits_mapString[split.Name] = split

	return split
}

// StagePreserveOrder puts split to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SplitOrder
// - update stage.SplitOrder accordingly
func (split *Split) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Splits[split]; !ok {
		stage.Splits[split] = struct{}{}

		if order > stage.SplitOrder {
			stage.SplitOrder = order
		}
		stage.Split_stagedOrder[split] = order
		stage.SplitOrder++
	}
	stage.Splits_mapString[split.Name] = split
}

// Unstage removes split off the model stage
func (split *Split) Unstage(stage *Stage) *Split {
	delete(stage.Splits, split)
	// issue1150
	// delete(stage.Split_stagedOrder, split)
	delete(stage.Splits_mapString, split.Name)

	return split
}

// UnstageVoid removes split off the model stage
func (split *Split) UnstageVoid(stage *Stage) {
	delete(stage.Splits, split)
	// issue1150
	// delete(stage.Split_stagedOrder, split)
	delete(stage.Splits_mapString, split.Name)
}

// commit split to the back repo (if it is already staged)
func (split *Split) Commit(stage *Stage) *Split {
	if _, ok := stage.Splits[split]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSplit(split)
		}
	}
	return split
}

func (split *Split) CommitVoid(stage *Stage) {
	split.Commit(stage)
}

func (split *Split) StageVoid(stage *Stage) {
	split.Stage(stage)
}

// Checkout split to the back repo (if it is already staged)
func (split *Split) Checkout(stage *Stage) *Split {
	if _, ok := stage.Splits[split]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSplit(split)
		}
	}
	return split
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
	if _, ok := stage.Svgs[svg]; !ok {
		stage.Svgs[svg] = struct{}{}
		stage.Svg_stagedOrder[svg] = stage.SvgOrder
		stage.SvgOrder++
	}
	stage.Svgs_mapString[svg.Name] = svg

	return svg
}

// StagePreserveOrder puts svg to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SvgOrder
// - update stage.SvgOrder accordingly
func (svg *Svg) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Svgs[svg]; !ok {
		stage.Svgs[svg] = struct{}{}

		if order > stage.SvgOrder {
			stage.SvgOrder = order
		}
		stage.Svg_stagedOrder[svg] = order
		stage.SvgOrder++
	}
	stage.Svgs_mapString[svg.Name] = svg
}

// Unstage removes svg off the model stage
func (svg *Svg) Unstage(stage *Stage) *Svg {
	delete(stage.Svgs, svg)
	// issue1150
	// delete(stage.Svg_stagedOrder, svg)
	delete(stage.Svgs_mapString, svg.Name)

	return svg
}

// UnstageVoid removes svg off the model stage
func (svg *Svg) UnstageVoid(stage *Stage) {
	delete(stage.Svgs, svg)
	// issue1150
	// delete(stage.Svg_stagedOrder, svg)
	delete(stage.Svgs_mapString, svg.Name)
}

// commit svg to the back repo (if it is already staged)
func (svg *Svg) Commit(stage *Stage) *Svg {
	if _, ok := stage.Svgs[svg]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSvg(svg)
		}
	}
	return svg
}

func (svg *Svg) CommitVoid(stage *Stage) {
	svg.Commit(stage)
}

func (svg *Svg) StageVoid(stage *Stage) {
	svg.Stage(stage)
}

// Checkout svg to the back repo (if it is already staged)
func (svg *Svg) Checkout(stage *Stage) *Svg {
	if _, ok := stage.Svgs[svg]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSvg(svg)
		}
	}
	return svg
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
	if _, ok := stage.Tables[table]; !ok {
		stage.Tables[table] = struct{}{}
		stage.Table_stagedOrder[table] = stage.TableOrder
		stage.TableOrder++
	}
	stage.Tables_mapString[table.Name] = table

	return table
}

// StagePreserveOrder puts table to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TableOrder
// - update stage.TableOrder accordingly
func (table *Table) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Tables[table]; !ok {
		stage.Tables[table] = struct{}{}

		if order > stage.TableOrder {
			stage.TableOrder = order
		}
		stage.Table_stagedOrder[table] = order
		stage.TableOrder++
	}
	stage.Tables_mapString[table.Name] = table
}

// Unstage removes table off the model stage
func (table *Table) Unstage(stage *Stage) *Table {
	delete(stage.Tables, table)
	// issue1150
	// delete(stage.Table_stagedOrder, table)
	delete(stage.Tables_mapString, table.Name)

	return table
}

// UnstageVoid removes table off the model stage
func (table *Table) UnstageVoid(stage *Stage) {
	delete(stage.Tables, table)
	// issue1150
	// delete(stage.Table_stagedOrder, table)
	delete(stage.Tables_mapString, table.Name)
}

// commit table to the back repo (if it is already staged)
func (table *Table) Commit(stage *Stage) *Table {
	if _, ok := stage.Tables[table]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTable(table)
		}
	}
	return table
}

func (table *Table) CommitVoid(stage *Stage) {
	table.Commit(stage)
}

func (table *Table) StageVoid(stage *Stage) {
	table.Stage(stage)
}

// Checkout table to the back repo (if it is already staged)
func (table *Table) Checkout(stage *Stage) *Table {
	if _, ok := stage.Tables[table]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTable(table)
		}
	}
	return table
}

// for satisfaction of GongStruct interface
func (table *Table) GetName() (res string) {
	return table.Name
}

// for satisfaction of GongStruct interface
func (table *Table) SetName(name string) {
	table.Name = name
}

// Stage puts title to the model stage
func (title *Title) Stage(stage *Stage) *Title {
	if _, ok := stage.Titles[title]; !ok {
		stage.Titles[title] = struct{}{}
		stage.Title_stagedOrder[title] = stage.TitleOrder
		stage.TitleOrder++
	}
	stage.Titles_mapString[title.Name] = title

	return title
}

// StagePreserveOrder puts title to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TitleOrder
// - update stage.TitleOrder accordingly
func (title *Title) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Titles[title]; !ok {
		stage.Titles[title] = struct{}{}

		if order > stage.TitleOrder {
			stage.TitleOrder = order
		}
		stage.Title_stagedOrder[title] = order
		stage.TitleOrder++
	}
	stage.Titles_mapString[title.Name] = title
}

// Unstage removes title off the model stage
func (title *Title) Unstage(stage *Stage) *Title {
	delete(stage.Titles, title)
	// issue1150
	// delete(stage.Title_stagedOrder, title)
	delete(stage.Titles_mapString, title.Name)

	return title
}

// UnstageVoid removes title off the model stage
func (title *Title) UnstageVoid(stage *Stage) {
	delete(stage.Titles, title)
	// issue1150
	// delete(stage.Title_stagedOrder, title)
	delete(stage.Titles_mapString, title.Name)
}

// commit title to the back repo (if it is already staged)
func (title *Title) Commit(stage *Stage) *Title {
	if _, ok := stage.Titles[title]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTitle(title)
		}
	}
	return title
}

func (title *Title) CommitVoid(stage *Stage) {
	title.Commit(stage)
}

func (title *Title) StageVoid(stage *Stage) {
	title.Stage(stage)
}

// Checkout title to the back repo (if it is already staged)
func (title *Title) Checkout(stage *Stage) *Title {
	if _, ok := stage.Titles[title]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTitle(title)
		}
	}
	return title
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
	if _, ok := stage.Tones[tone]; !ok {
		stage.Tones[tone] = struct{}{}
		stage.Tone_stagedOrder[tone] = stage.ToneOrder
		stage.ToneOrder++
	}
	stage.Tones_mapString[tone.Name] = tone

	return tone
}

// StagePreserveOrder puts tone to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ToneOrder
// - update stage.ToneOrder accordingly
func (tone *Tone) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Tones[tone]; !ok {
		stage.Tones[tone] = struct{}{}

		if order > stage.ToneOrder {
			stage.ToneOrder = order
		}
		stage.Tone_stagedOrder[tone] = order
		stage.ToneOrder++
	}
	stage.Tones_mapString[tone.Name] = tone
}

// Unstage removes tone off the model stage
func (tone *Tone) Unstage(stage *Stage) *Tone {
	delete(stage.Tones, tone)
	// issue1150
	// delete(stage.Tone_stagedOrder, tone)
	delete(stage.Tones_mapString, tone.Name)

	return tone
}

// UnstageVoid removes tone off the model stage
func (tone *Tone) UnstageVoid(stage *Stage) {
	delete(stage.Tones, tone)
	// issue1150
	// delete(stage.Tone_stagedOrder, tone)
	delete(stage.Tones_mapString, tone.Name)
}

// commit tone to the back repo (if it is already staged)
func (tone *Tone) Commit(stage *Stage) *Tone {
	if _, ok := stage.Tones[tone]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTone(tone)
		}
	}
	return tone
}

func (tone *Tone) CommitVoid(stage *Stage) {
	tone.Commit(stage)
}

func (tone *Tone) StageVoid(stage *Stage) {
	tone.Stage(stage)
}

// Checkout tone to the back repo (if it is already staged)
func (tone *Tone) Checkout(stage *Stage) *Tone {
	if _, ok := stage.Tones[tone]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTone(tone)
		}
	}
	return tone
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
	if _, ok := stage.Trees[tree]; !ok {
		stage.Trees[tree] = struct{}{}
		stage.Tree_stagedOrder[tree] = stage.TreeOrder
		stage.TreeOrder++
	}
	stage.Trees_mapString[tree.Name] = tree

	return tree
}

// StagePreserveOrder puts tree to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TreeOrder
// - update stage.TreeOrder accordingly
func (tree *Tree) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Trees[tree]; !ok {
		stage.Trees[tree] = struct{}{}

		if order > stage.TreeOrder {
			stage.TreeOrder = order
		}
		stage.Tree_stagedOrder[tree] = order
		stage.TreeOrder++
	}
	stage.Trees_mapString[tree.Name] = tree
}

// Unstage removes tree off the model stage
func (tree *Tree) Unstage(stage *Stage) *Tree {
	delete(stage.Trees, tree)
	// issue1150
	// delete(stage.Tree_stagedOrder, tree)
	delete(stage.Trees_mapString, tree.Name)

	return tree
}

// UnstageVoid removes tree off the model stage
func (tree *Tree) UnstageVoid(stage *Stage) {
	delete(stage.Trees, tree)
	// issue1150
	// delete(stage.Tree_stagedOrder, tree)
	delete(stage.Trees_mapString, tree.Name)
}

// commit tree to the back repo (if it is already staged)
func (tree *Tree) Commit(stage *Stage) *Tree {
	if _, ok := stage.Trees[tree]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTree(tree)
		}
	}
	return tree
}

func (tree *Tree) CommitVoid(stage *Stage) {
	tree.Commit(stage)
}

func (tree *Tree) StageVoid(stage *Stage) {
	tree.Stage(stage)
}

// Checkout tree to the back repo (if it is already staged)
func (tree *Tree) Checkout(stage *Stage) *Tree {
	if _, ok := stage.Trees[tree]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTree(tree)
		}
	}
	return tree
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
	if _, ok := stage.Views[view]; !ok {
		stage.Views[view] = struct{}{}
		stage.View_stagedOrder[view] = stage.ViewOrder
		stage.ViewOrder++
	}
	stage.Views_mapString[view.Name] = view

	return view
}

// StagePreserveOrder puts view to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ViewOrder
// - update stage.ViewOrder accordingly
func (view *View) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Views[view]; !ok {
		stage.Views[view] = struct{}{}

		if order > stage.ViewOrder {
			stage.ViewOrder = order
		}
		stage.View_stagedOrder[view] = order
		stage.ViewOrder++
	}
	stage.Views_mapString[view.Name] = view
}

// Unstage removes view off the model stage
func (view *View) Unstage(stage *Stage) *View {
	delete(stage.Views, view)
	// issue1150
	// delete(stage.View_stagedOrder, view)
	delete(stage.Views_mapString, view.Name)

	return view
}

// UnstageVoid removes view off the model stage
func (view *View) UnstageVoid(stage *Stage) {
	delete(stage.Views, view)
	// issue1150
	// delete(stage.View_stagedOrder, view)
	delete(stage.Views_mapString, view.Name)
}

// commit view to the back repo (if it is already staged)
func (view *View) Commit(stage *Stage) *View {
	if _, ok := stage.Views[view]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitView(view)
		}
	}
	return view
}

func (view *View) CommitVoid(stage *Stage) {
	view.Commit(stage)
}

func (view *View) StageVoid(stage *Stage) {
	view.Stage(stage)
}

// Checkout view to the back repo (if it is already staged)
func (view *View) Checkout(stage *Stage) *View {
	if _, ok := stage.Views[view]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutView(view)
		}
	}
	return view
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
	if _, ok := stage.Xlsxs[xlsx]; !ok {
		stage.Xlsxs[xlsx] = struct{}{}
		stage.Xlsx_stagedOrder[xlsx] = stage.XlsxOrder
		stage.XlsxOrder++
	}
	stage.Xlsxs_mapString[xlsx.Name] = xlsx

	return xlsx
}

// StagePreserveOrder puts xlsx to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.XlsxOrder
// - update stage.XlsxOrder accordingly
func (xlsx *Xlsx) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Xlsxs[xlsx]; !ok {
		stage.Xlsxs[xlsx] = struct{}{}

		if order > stage.XlsxOrder {
			stage.XlsxOrder = order
		}
		stage.Xlsx_stagedOrder[xlsx] = order
		stage.XlsxOrder++
	}
	stage.Xlsxs_mapString[xlsx.Name] = xlsx
}

// Unstage removes xlsx off the model stage
func (xlsx *Xlsx) Unstage(stage *Stage) *Xlsx {
	delete(stage.Xlsxs, xlsx)
	// issue1150
	// delete(stage.Xlsx_stagedOrder, xlsx)
	delete(stage.Xlsxs_mapString, xlsx.Name)

	return xlsx
}

// UnstageVoid removes xlsx off the model stage
func (xlsx *Xlsx) UnstageVoid(stage *Stage) {
	delete(stage.Xlsxs, xlsx)
	// issue1150
	// delete(stage.Xlsx_stagedOrder, xlsx)
	delete(stage.Xlsxs_mapString, xlsx.Name)
}

// commit xlsx to the back repo (if it is already staged)
func (xlsx *Xlsx) Commit(stage *Stage) *Xlsx {
	if _, ok := stage.Xlsxs[xlsx]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXlsx(xlsx)
		}
	}
	return xlsx
}

func (xlsx *Xlsx) CommitVoid(stage *Stage) {
	xlsx.Commit(stage)
}

func (xlsx *Xlsx) StageVoid(stage *Stage) {
	xlsx.Stage(stage)
}

// Checkout xlsx to the back repo (if it is already staged)
func (xlsx *Xlsx) Checkout(stage *Stage) *Xlsx {
	if _, ok := stage.Xlsxs[xlsx]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXlsx(xlsx)
		}
	}
	return xlsx
}

// for satisfaction of GongStruct interface
func (xlsx *Xlsx) GetName() (res string) {
	return xlsx.Name
}

// for satisfaction of GongStruct interface
func (xlsx *Xlsx) SetName(name string) {
	xlsx.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMAsSplit(AsSplit *AsSplit)
	CreateORMAsSplitArea(AsSplitArea *AsSplitArea)
	CreateORMButton(Button *Button)
	CreateORMCursor(Cursor *Cursor)
	CreateORMFavIcon(FavIcon *FavIcon)
	CreateORMForm(Form *Form)
	CreateORMLoad(Load *Load)
	CreateORMLogoOnTheLeft(LogoOnTheLeft *LogoOnTheLeft)
	CreateORMLogoOnTheRight(LogoOnTheRight *LogoOnTheRight)
	CreateORMMarkdown(Markdown *Markdown)
	CreateORMSlider(Slider *Slider)
	CreateORMSplit(Split *Split)
	CreateORMSvg(Svg *Svg)
	CreateORMTable(Table *Table)
	CreateORMTitle(Title *Title)
	CreateORMTone(Tone *Tone)
	CreateORMTree(Tree *Tree)
	CreateORMView(View *View)
	CreateORMXlsx(Xlsx *Xlsx)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAsSplit(AsSplit *AsSplit)
	DeleteORMAsSplitArea(AsSplitArea *AsSplitArea)
	DeleteORMButton(Button *Button)
	DeleteORMCursor(Cursor *Cursor)
	DeleteORMFavIcon(FavIcon *FavIcon)
	DeleteORMForm(Form *Form)
	DeleteORMLoad(Load *Load)
	DeleteORMLogoOnTheLeft(LogoOnTheLeft *LogoOnTheLeft)
	DeleteORMLogoOnTheRight(LogoOnTheRight *LogoOnTheRight)
	DeleteORMMarkdown(Markdown *Markdown)
	DeleteORMSlider(Slider *Slider)
	DeleteORMSplit(Split *Split)
	DeleteORMSvg(Svg *Svg)
	DeleteORMTable(Table *Table)
	DeleteORMTitle(Title *Title)
	DeleteORMTone(Tone *Tone)
	DeleteORMTree(Tree *Tree)
	DeleteORMView(View *View)
	DeleteORMXlsx(Xlsx *Xlsx)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.AsSplits = make(map[*AsSplit]struct{})
	stage.AsSplits_mapString = make(map[string]*AsSplit)
	stage.AsSplit_stagedOrder = make(map[*AsSplit]uint)
	stage.AsSplitOrder = 0

	stage.AsSplitAreas = make(map[*AsSplitArea]struct{})
	stage.AsSplitAreas_mapString = make(map[string]*AsSplitArea)
	stage.AsSplitArea_stagedOrder = make(map[*AsSplitArea]uint)
	stage.AsSplitAreaOrder = 0

	stage.Buttons = make(map[*Button]struct{})
	stage.Buttons_mapString = make(map[string]*Button)
	stage.Button_stagedOrder = make(map[*Button]uint)
	stage.ButtonOrder = 0

	stage.Cursors = make(map[*Cursor]struct{})
	stage.Cursors_mapString = make(map[string]*Cursor)
	stage.Cursor_stagedOrder = make(map[*Cursor]uint)
	stage.CursorOrder = 0

	stage.FavIcons = make(map[*FavIcon]struct{})
	stage.FavIcons_mapString = make(map[string]*FavIcon)
	stage.FavIcon_stagedOrder = make(map[*FavIcon]uint)
	stage.FavIconOrder = 0

	stage.Forms = make(map[*Form]struct{})
	stage.Forms_mapString = make(map[string]*Form)
	stage.Form_stagedOrder = make(map[*Form]uint)
	stage.FormOrder = 0

	stage.Loads = make(map[*Load]struct{})
	stage.Loads_mapString = make(map[string]*Load)
	stage.Load_stagedOrder = make(map[*Load]uint)
	stage.LoadOrder = 0

	stage.LogoOnTheLefts = make(map[*LogoOnTheLeft]struct{})
	stage.LogoOnTheLefts_mapString = make(map[string]*LogoOnTheLeft)
	stage.LogoOnTheLeft_stagedOrder = make(map[*LogoOnTheLeft]uint)
	stage.LogoOnTheLeftOrder = 0

	stage.LogoOnTheRights = make(map[*LogoOnTheRight]struct{})
	stage.LogoOnTheRights_mapString = make(map[string]*LogoOnTheRight)
	stage.LogoOnTheRight_stagedOrder = make(map[*LogoOnTheRight]uint)
	stage.LogoOnTheRightOrder = 0

	stage.Markdowns = make(map[*Markdown]struct{})
	stage.Markdowns_mapString = make(map[string]*Markdown)
	stage.Markdown_stagedOrder = make(map[*Markdown]uint)
	stage.MarkdownOrder = 0

	stage.Sliders = make(map[*Slider]struct{})
	stage.Sliders_mapString = make(map[string]*Slider)
	stage.Slider_stagedOrder = make(map[*Slider]uint)
	stage.SliderOrder = 0

	stage.Splits = make(map[*Split]struct{})
	stage.Splits_mapString = make(map[string]*Split)
	stage.Split_stagedOrder = make(map[*Split]uint)
	stage.SplitOrder = 0

	stage.Svgs = make(map[*Svg]struct{})
	stage.Svgs_mapString = make(map[string]*Svg)
	stage.Svg_stagedOrder = make(map[*Svg]uint)
	stage.SvgOrder = 0

	stage.Tables = make(map[*Table]struct{})
	stage.Tables_mapString = make(map[string]*Table)
	stage.Table_stagedOrder = make(map[*Table]uint)
	stage.TableOrder = 0

	stage.Titles = make(map[*Title]struct{})
	stage.Titles_mapString = make(map[string]*Title)
	stage.Title_stagedOrder = make(map[*Title]uint)
	stage.TitleOrder = 0

	stage.Tones = make(map[*Tone]struct{})
	stage.Tones_mapString = make(map[string]*Tone)
	stage.Tone_stagedOrder = make(map[*Tone]uint)
	stage.ToneOrder = 0

	stage.Trees = make(map[*Tree]struct{})
	stage.Trees_mapString = make(map[string]*Tree)
	stage.Tree_stagedOrder = make(map[*Tree]uint)
	stage.TreeOrder = 0

	stage.Views = make(map[*View]struct{})
	stage.Views_mapString = make(map[string]*View)
	stage.View_stagedOrder = make(map[*View]uint)
	stage.ViewOrder = 0

	stage.Xlsxs = make(map[*Xlsx]struct{})
	stage.Xlsxs_mapString = make(map[string]*Xlsx)
	stage.Xlsx_stagedOrder = make(map[*Xlsx]uint)
	stage.XlsxOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.AsSplits = nil
	stage.AsSplits_mapString = nil

	stage.AsSplitAreas = nil
	stage.AsSplitAreas_mapString = nil

	stage.Buttons = nil
	stage.Buttons_mapString = nil

	stage.Cursors = nil
	stage.Cursors_mapString = nil

	stage.FavIcons = nil
	stage.FavIcons_mapString = nil

	stage.Forms = nil
	stage.Forms_mapString = nil

	stage.Loads = nil
	stage.Loads_mapString = nil

	stage.LogoOnTheLefts = nil
	stage.LogoOnTheLefts_mapString = nil

	stage.LogoOnTheRights = nil
	stage.LogoOnTheRights_mapString = nil

	stage.Markdowns = nil
	stage.Markdowns_mapString = nil

	stage.Sliders = nil
	stage.Sliders_mapString = nil

	stage.Splits = nil
	stage.Splits_mapString = nil

	stage.Svgs = nil
	stage.Svgs_mapString = nil

	stage.Tables = nil
	stage.Tables_mapString = nil

	stage.Titles = nil
	stage.Titles_mapString = nil

	stage.Tones = nil
	stage.Tones_mapString = nil

	stage.Trees = nil
	stage.Trees_mapString = nil

	stage.Views = nil
	stage.Views_mapString = nil

	stage.Xlsxs = nil
	stage.Xlsxs_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for assplit := range stage.AsSplits {
		assplit.Unstage(stage)
	}

	for assplitarea := range stage.AsSplitAreas {
		assplitarea.Unstage(stage)
	}

	for button := range stage.Buttons {
		button.Unstage(stage)
	}

	for cursor := range stage.Cursors {
		cursor.Unstage(stage)
	}

	for favicon := range stage.FavIcons {
		favicon.Unstage(stage)
	}

	for form := range stage.Forms {
		form.Unstage(stage)
	}

	for load := range stage.Loads {
		load.Unstage(stage)
	}

	for logoontheleft := range stage.LogoOnTheLefts {
		logoontheleft.Unstage(stage)
	}

	for logoontheright := range stage.LogoOnTheRights {
		logoontheright.Unstage(stage)
	}

	for markdown := range stage.Markdowns {
		markdown.Unstage(stage)
	}

	for slider := range stage.Sliders {
		slider.Unstage(stage)
	}

	for split := range stage.Splits {
		split.Unstage(stage)
	}

	for svg := range stage.Svgs {
		svg.Unstage(stage)
	}

	for table := range stage.Tables {
		table.Unstage(stage)
	}

	for title := range stage.Titles {
		title.Unstage(stage)
	}

	for tone := range stage.Tones {
		tone.Unstage(stage)
	}

	for tree := range stage.Trees {
		tree.Unstage(stage)
	}

	for view := range stage.Views {
		view.Unstage(stage)
	}

	for xlsx := range stage.Xlsxs {
		xlsx.Unstage(stage)
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
	case map[*AsSplit]any:
		return any(&stage.AsSplits).(*Type)
	case map[*AsSplitArea]any:
		return any(&stage.AsSplitAreas).(*Type)
	case map[*Button]any:
		return any(&stage.Buttons).(*Type)
	case map[*Cursor]any:
		return any(&stage.Cursors).(*Type)
	case map[*FavIcon]any:
		return any(&stage.FavIcons).(*Type)
	case map[*Form]any:
		return any(&stage.Forms).(*Type)
	case map[*Load]any:
		return any(&stage.Loads).(*Type)
	case map[*LogoOnTheLeft]any:
		return any(&stage.LogoOnTheLefts).(*Type)
	case map[*LogoOnTheRight]any:
		return any(&stage.LogoOnTheRights).(*Type)
	case map[*Markdown]any:
		return any(&stage.Markdowns).(*Type)
	case map[*Slider]any:
		return any(&stage.Sliders).(*Type)
	case map[*Split]any:
		return any(&stage.Splits).(*Type)
	case map[*Svg]any:
		return any(&stage.Svgs).(*Type)
	case map[*Table]any:
		return any(&stage.Tables).(*Type)
	case map[*Title]any:
		return any(&stage.Titles).(*Type)
	case map[*Tone]any:
		return any(&stage.Tones).(*Type)
	case map[*Tree]any:
		return any(&stage.Trees).(*Type)
	case map[*View]any:
		return any(&stage.Views).(*Type)
	case map[*Xlsx]any:
		return any(&stage.Xlsxs).(*Type)
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

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case AsSplit:
		return any(&stage.AsSplits).(*map[*Type]struct{})
	case AsSplitArea:
		return any(&stage.AsSplitAreas).(*map[*Type]struct{})
	case Button:
		return any(&stage.Buttons).(*map[*Type]struct{})
	case Cursor:
		return any(&stage.Cursors).(*map[*Type]struct{})
	case FavIcon:
		return any(&stage.FavIcons).(*map[*Type]struct{})
	case Form:
		return any(&stage.Forms).(*map[*Type]struct{})
	case Load:
		return any(&stage.Loads).(*map[*Type]struct{})
	case LogoOnTheLeft:
		return any(&stage.LogoOnTheLefts).(*map[*Type]struct{})
	case LogoOnTheRight:
		return any(&stage.LogoOnTheRights).(*map[*Type]struct{})
	case Markdown:
		return any(&stage.Markdowns).(*map[*Type]struct{})
	case Slider:
		return any(&stage.Sliders).(*map[*Type]struct{})
	case Split:
		return any(&stage.Splits).(*map[*Type]struct{})
	case Svg:
		return any(&stage.Svgs).(*map[*Type]struct{})
	case Table:
		return any(&stage.Tables).(*map[*Type]struct{})
	case Title:
		return any(&stage.Titles).(*map[*Type]struct{})
	case Tone:
		return any(&stage.Tones).(*map[*Type]struct{})
	case Tree:
		return any(&stage.Trees).(*map[*Type]struct{})
	case View:
		return any(&stage.Views).(*map[*Type]struct{})
	case Xlsx:
		return any(&stage.Xlsxs).(*map[*Type]struct{})
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

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *Stage) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case AsSplit:
		return any(&stage.AsSplits_mapString).(*map[string]*Type)
	case AsSplitArea:
		return any(&stage.AsSplitAreas_mapString).(*map[string]*Type)
	case Button:
		return any(&stage.Buttons_mapString).(*map[string]*Type)
	case Cursor:
		return any(&stage.Cursors_mapString).(*map[string]*Type)
	case FavIcon:
		return any(&stage.FavIcons_mapString).(*map[string]*Type)
	case Form:
		return any(&stage.Forms_mapString).(*map[string]*Type)
	case Load:
		return any(&stage.Loads_mapString).(*map[string]*Type)
	case LogoOnTheLeft:
		return any(&stage.LogoOnTheLefts_mapString).(*map[string]*Type)
	case LogoOnTheRight:
		return any(&stage.LogoOnTheRights_mapString).(*map[string]*Type)
	case Markdown:
		return any(&stage.Markdowns_mapString).(*map[string]*Type)
	case Slider:
		return any(&stage.Sliders_mapString).(*map[string]*Type)
	case Split:
		return any(&stage.Splits_mapString).(*map[string]*Type)
	case Svg:
		return any(&stage.Svgs_mapString).(*map[string]*Type)
	case Table:
		return any(&stage.Tables_mapString).(*map[string]*Type)
	case Title:
		return any(&stage.Titles_mapString).(*map[string]*Type)
	case Tone:
		return any(&stage.Tones_mapString).(*map[string]*Type)
	case Tree:
		return any(&stage.Trees_mapString).(*map[string]*Type)
	case View:
		return any(&stage.Views_mapString).(*map[string]*Type)
	case Xlsx:
		return any(&stage.Xlsxs_mapString).(*map[string]*Type)
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
	case AsSplit:
		return any(&AsSplit{
			// Initialisation of associations
			// field is initialized with an instance of AsSplitArea with the name of the field
			AsSplitAreas: []*AsSplitArea{{Name: "AsSplitAreas"}},
		}).(*Type)
	case AsSplitArea:
		return any(&AsSplitArea{
			// Initialisation of associations
			// field is initialized with an instance of AsSplit with the name of the field
			AsSplit: &AsSplit{Name: "AsSplit"},
			// field is initialized with an instance of Button with the name of the field
			Button: &Button{Name: "Button"},
			// field is initialized with an instance of Cursor with the name of the field
			Cursor: &Cursor{Name: "Cursor"},
			// field is initialized with an instance of Form with the name of the field
			Form: &Form{Name: "Form"},
			// field is initialized with an instance of Load with the name of the field
			Load: &Load{Name: "Load"},
			// field is initialized with an instance of Markdown with the name of the field
			Markdown: &Markdown{Name: "Markdown"},
			// field is initialized with an instance of Slider with the name of the field
			Slider: &Slider{Name: "Slider"},
			// field is initialized with an instance of Split with the name of the field
			Split: &Split{Name: "Split"},
			// field is initialized with an instance of Svg with the name of the field
			Svg: &Svg{Name: "Svg"},
			// field is initialized with an instance of Table with the name of the field
			Table: &Table{Name: "Table"},
			// field is initialized with an instance of Tone with the name of the field
			Tone: &Tone{Name: "Tone"},
			// field is initialized with an instance of Tree with the name of the field
			Tree: &Tree{Name: "Tree"},
			// field is initialized with an instance of Xlsx with the name of the field
			Xlsx: &Xlsx{Name: "Xlsx"},
		}).(*Type)
	case Button:
		return any(&Button{
			// Initialisation of associations
		}).(*Type)
	case Cursor:
		return any(&Cursor{
			// Initialisation of associations
		}).(*Type)
	case FavIcon:
		return any(&FavIcon{
			// Initialisation of associations
		}).(*Type)
	case Form:
		return any(&Form{
			// Initialisation of associations
		}).(*Type)
	case Load:
		return any(&Load{
			// Initialisation of associations
		}).(*Type)
	case LogoOnTheLeft:
		return any(&LogoOnTheLeft{
			// Initialisation of associations
		}).(*Type)
	case LogoOnTheRight:
		return any(&LogoOnTheRight{
			// Initialisation of associations
		}).(*Type)
	case Markdown:
		return any(&Markdown{
			// Initialisation of associations
		}).(*Type)
	case Slider:
		return any(&Slider{
			// Initialisation of associations
		}).(*Type)
	case Split:
		return any(&Split{
			// Initialisation of associations
		}).(*Type)
	case Svg:
		return any(&Svg{
			// Initialisation of associations
		}).(*Type)
	case Table:
		return any(&Table{
			// Initialisation of associations
		}).(*Type)
	case Title:
		return any(&Title{
			// Initialisation of associations
		}).(*Type)
	case Tone:
		return any(&Tone{
			// Initialisation of associations
		}).(*Type)
	case Tree:
		return any(&Tree{
			// Initialisation of associations
		}).(*Type)
	case View:
		return any(&View{
			// Initialisation of associations
			// field is initialized with an instance of AsSplitArea with the name of the field
			RootAsSplitAreas: []*AsSplitArea{{Name: "RootAsSplitAreas"}},
		}).(*Type)
	case Xlsx:
		return any(&Xlsx{
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

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type GongstructIF]() (res string) {
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

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type GongstructIF]() (res []ReverseField) {
	res = make([]ReverseField, 0)

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

// insertion point for get fields header method
func (assplit *AsSplit) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "EnumString",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumTypeString",
		},
		{
			Name:                 "AsSplitAreas",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AsSplitArea",
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
			Name:                 "EnumString",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumTypeString",
		},
		{
			Name:               "IsSecondatyView",
			GongFieldValueType: GongFieldValueTypeBool,
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
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
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, assplitarea.AsSplit))
		}
	case "Button":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Button != nil {
			res.valueString = assplitarea.Button.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, assplitarea.Button))
		}
	case "Cursor":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Cursor != nil {
			res.valueString = assplitarea.Cursor.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, assplitarea.Cursor))
		}
	case "Form":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Form != nil {
			res.valueString = assplitarea.Form.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, assplitarea.Form))
		}
	case "Load":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Load != nil {
			res.valueString = assplitarea.Load.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, assplitarea.Load))
		}
	case "Markdown":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Markdown != nil {
			res.valueString = assplitarea.Markdown.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, assplitarea.Markdown))
		}
	case "Slider":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Slider != nil {
			res.valueString = assplitarea.Slider.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, assplitarea.Slider))
		}
	case "Split":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Split != nil {
			res.valueString = assplitarea.Split.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, assplitarea.Split))
		}
	case "Svg":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Svg != nil {
			res.valueString = assplitarea.Svg.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, assplitarea.Svg))
		}
	case "Table":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Table != nil {
			res.valueString = assplitarea.Table.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, assplitarea.Table))
		}
	case "Tone":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Tone != nil {
			res.valueString = assplitarea.Tone.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, assplitarea.Tone))
		}
	case "Tree":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Tree != nil {
			res.valueString = assplitarea.Tree.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, assplitarea.Tree))
		}
	case "Xlsx":
		res.GongFieldValueType = GongFieldValueTypePointer
		if assplitarea.Xlsx != nil {
			res.valueString = assplitarea.Xlsx.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, assplitarea.Xlsx))
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "IsSelectedView":
		res.valueString = fmt.Sprintf("%t", view.IsSelectedView)
		res.valueBool = view.IsSelectedView
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Direction":
		enum := view.Direction
		res.valueString = enum.ToCodeString()
	case "IsSecondatyView":
		res.valueString = fmt.Sprintf("%t", view.IsSecondatyView)
		res.valueBool = view.IsSecondatyView
		res.GongFieldValueType = GongFieldValueTypeBool
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

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (assplit *AsSplit) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		assplit.Name = value.GetValueString()
	case "Direction":
		assplit.Direction.FromCodeString(value.GetValueString())
	case "AsSplitAreas":
		assplit.AsSplitAreas = make([]*AsSplitArea, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.AsSplitAreas {
					if stage.AsSplitArea_stagedOrder[__instance__] == uint(id) {
						assplit.AsSplitAreas = append(assplit.AsSplitAreas, __instance__)
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

func (assplitarea *AsSplitArea) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		assplitarea.Name = value.GetValueString()
	case "ShowNameInHeader":
		assplitarea.ShowNameInHeader = value.GetValueBool()
	case "Size":
		assplitarea.Size = value.GetValueFloat()
	case "IsAny":
		assplitarea.IsAny = value.GetValueBool()
	case "AsSplit":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			assplitarea.AsSplit = nil
			for __instance__ := range stage.AsSplits {
				if stage.AsSplit_stagedOrder[__instance__] == uint(id) {
					assplitarea.AsSplit = __instance__
					break
				}
			}
		}
	case "Button":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			assplitarea.Button = nil
			for __instance__ := range stage.Buttons {
				if stage.Button_stagedOrder[__instance__] == uint(id) {
					assplitarea.Button = __instance__
					break
				}
			}
		}
	case "Cursor":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			assplitarea.Cursor = nil
			for __instance__ := range stage.Cursors {
				if stage.Cursor_stagedOrder[__instance__] == uint(id) {
					assplitarea.Cursor = __instance__
					break
				}
			}
		}
	case "Form":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			assplitarea.Form = nil
			for __instance__ := range stage.Forms {
				if stage.Form_stagedOrder[__instance__] == uint(id) {
					assplitarea.Form = __instance__
					break
				}
			}
		}
	case "Load":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			assplitarea.Load = nil
			for __instance__ := range stage.Loads {
				if stage.Load_stagedOrder[__instance__] == uint(id) {
					assplitarea.Load = __instance__
					break
				}
			}
		}
	case "Markdown":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			assplitarea.Markdown = nil
			for __instance__ := range stage.Markdowns {
				if stage.Markdown_stagedOrder[__instance__] == uint(id) {
					assplitarea.Markdown = __instance__
					break
				}
			}
		}
	case "Slider":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			assplitarea.Slider = nil
			for __instance__ := range stage.Sliders {
				if stage.Slider_stagedOrder[__instance__] == uint(id) {
					assplitarea.Slider = __instance__
					break
				}
			}
		}
	case "Split":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			assplitarea.Split = nil
			for __instance__ := range stage.Splits {
				if stage.Split_stagedOrder[__instance__] == uint(id) {
					assplitarea.Split = __instance__
					break
				}
			}
		}
	case "Svg":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			assplitarea.Svg = nil
			for __instance__ := range stage.Svgs {
				if stage.Svg_stagedOrder[__instance__] == uint(id) {
					assplitarea.Svg = __instance__
					break
				}
			}
		}
	case "Table":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			assplitarea.Table = nil
			for __instance__ := range stage.Tables {
				if stage.Table_stagedOrder[__instance__] == uint(id) {
					assplitarea.Table = __instance__
					break
				}
			}
		}
	case "Tone":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			assplitarea.Tone = nil
			for __instance__ := range stage.Tones {
				if stage.Tone_stagedOrder[__instance__] == uint(id) {
					assplitarea.Tone = __instance__
					break
				}
			}
		}
	case "Tree":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			assplitarea.Tree = nil
			for __instance__ := range stage.Trees {
				if stage.Tree_stagedOrder[__instance__] == uint(id) {
					assplitarea.Tree = __instance__
					break
				}
			}
		}
	case "Xlsx":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			assplitarea.Xlsx = nil
			for __instance__ := range stage.Xlsxs {
				if stage.Xlsx_stagedOrder[__instance__] == uint(id) {
					assplitarea.Xlsx = __instance__
					break
				}
			}
		}
	case "HasDiv":
		assplitarea.HasDiv = value.GetValueBool()
	case "DivStyle":
		assplitarea.DivStyle = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (button *Button) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		button.Name = value.GetValueString()
	case "StackName":
		button.StackName = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (cursor *Cursor) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		cursor.Name = value.GetValueString()
	case "StackName":
		cursor.StackName = value.GetValueString()
	case "Style":
		cursor.Style = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (favicon *FavIcon) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		favicon.Name = value.GetValueString()
	case "SVG":
		favicon.SVG = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (form *Form) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		form.Name = value.GetValueString()
	case "StackName":
		form.StackName = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (load *Load) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		load.Name = value.GetValueString()
	case "StackName":
		load.StackName = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (logoontheleft *LogoOnTheLeft) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		logoontheleft.Name = value.GetValueString()
	case "Width":
		logoontheleft.Width = int(value.GetValueInt())
	case "Height":
		logoontheleft.Height = int(value.GetValueInt())
	case "SVG":
		logoontheleft.SVG = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (logoontheright *LogoOnTheRight) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		logoontheright.Name = value.GetValueString()
	case "Width":
		logoontheright.Width = int(value.GetValueInt())
	case "Height":
		logoontheright.Height = int(value.GetValueInt())
	case "SVG":
		logoontheright.SVG = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (markdown *Markdown) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		markdown.Name = value.GetValueString()
	case "StackName":
		markdown.StackName = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (slider *Slider) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		slider.Name = value.GetValueString()
	case "StackName":
		slider.StackName = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (split *Split) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		split.Name = value.GetValueString()
	case "StackName":
		split.StackName = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (svg *Svg) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		svg.Name = value.GetValueString()
	case "StackName":
		svg.StackName = value.GetValueString()
	case "Style":
		svg.Style = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (table *Table) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		table.Name = value.GetValueString()
	case "StackName":
		table.StackName = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (title *Title) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		title.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (tone *Tone) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		tone.Name = value.GetValueString()
	case "StackName":
		tone.StackName = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (tree *Tree) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		tree.Name = value.GetValueString()
	case "StackName":
		tree.StackName = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (view *View) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		view.Name = value.GetValueString()
	case "ShowViewName":
		view.ShowViewName = value.GetValueBool()
	case "RootAsSplitAreas":
		view.RootAsSplitAreas = make([]*AsSplitArea, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.AsSplitAreas {
					if stage.AsSplitArea_stagedOrder[__instance__] == uint(id) {
						view.RootAsSplitAreas = append(view.RootAsSplitAreas, __instance__)
						break
					}
				}
			}
		}
	case "IsSelectedView":
		view.IsSelectedView = value.GetValueBool()
	case "Direction":
		view.Direction.FromCodeString(value.GetValueString())
	case "IsSecondatyView":
		view.IsSecondatyView = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (xlsx *Xlsx) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		xlsx.Name = value.GetValueString()
	case "StackName":
		xlsx.StackName = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
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

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	stage.AsSplits_mapString = make(map[string]*AsSplit)
	for assplit := range stage.AsSplits {
		stage.AsSplits_mapString[assplit.Name] = assplit
	}

	stage.AsSplitAreas_mapString = make(map[string]*AsSplitArea)
	for assplitarea := range stage.AsSplitAreas {
		stage.AsSplitAreas_mapString[assplitarea.Name] = assplitarea
	}

	stage.Buttons_mapString = make(map[string]*Button)
	for button := range stage.Buttons {
		stage.Buttons_mapString[button.Name] = button
	}

	stage.Cursors_mapString = make(map[string]*Cursor)
	for cursor := range stage.Cursors {
		stage.Cursors_mapString[cursor.Name] = cursor
	}

	stage.FavIcons_mapString = make(map[string]*FavIcon)
	for favicon := range stage.FavIcons {
		stage.FavIcons_mapString[favicon.Name] = favicon
	}

	stage.Forms_mapString = make(map[string]*Form)
	for form := range stage.Forms {
		stage.Forms_mapString[form.Name] = form
	}

	stage.Loads_mapString = make(map[string]*Load)
	for load := range stage.Loads {
		stage.Loads_mapString[load.Name] = load
	}

	stage.LogoOnTheLefts_mapString = make(map[string]*LogoOnTheLeft)
	for logoontheleft := range stage.LogoOnTheLefts {
		stage.LogoOnTheLefts_mapString[logoontheleft.Name] = logoontheleft
	}

	stage.LogoOnTheRights_mapString = make(map[string]*LogoOnTheRight)
	for logoontheright := range stage.LogoOnTheRights {
		stage.LogoOnTheRights_mapString[logoontheright.Name] = logoontheright
	}

	stage.Markdowns_mapString = make(map[string]*Markdown)
	for markdown := range stage.Markdowns {
		stage.Markdowns_mapString[markdown.Name] = markdown
	}

	stage.Sliders_mapString = make(map[string]*Slider)
	for slider := range stage.Sliders {
		stage.Sliders_mapString[slider.Name] = slider
	}

	stage.Splits_mapString = make(map[string]*Split)
	for split := range stage.Splits {
		stage.Splits_mapString[split.Name] = split
	}

	stage.Svgs_mapString = make(map[string]*Svg)
	for svg := range stage.Svgs {
		stage.Svgs_mapString[svg.Name] = svg
	}

	stage.Tables_mapString = make(map[string]*Table)
	for table := range stage.Tables {
		stage.Tables_mapString[table.Name] = table
	}

	stage.Titles_mapString = make(map[string]*Title)
	for title := range stage.Titles {
		stage.Titles_mapString[title.Name] = title
	}

	stage.Tones_mapString = make(map[string]*Tone)
	for tone := range stage.Tones {
		stage.Tones_mapString[tone.Name] = tone
	}

	stage.Trees_mapString = make(map[string]*Tree)
	for tree := range stage.Trees {
		stage.Trees_mapString[tree.Name] = tree
	}

	stage.Views_mapString = make(map[string]*View)
	for view := range stage.Views {
		stage.Views_mapString[view.Name] = view
	}

	stage.Xlsxs_mapString = make(map[string]*Xlsx)
	for xlsx := range stage.Xlsxs {
		stage.Xlsxs_mapString[xlsx.Name] = xlsx
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
