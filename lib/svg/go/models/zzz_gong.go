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
	Animates                map[*Animate]struct{}
	Animates_instance       map[*Animate]*Animate
	Animates_mapString      map[string]*Animate
	AnimateOrder            uint
	Animate_stagedOrder     map[*Animate]uint
	Animate_orderStaged     map[uint]*Animate
	Animates_reference      map[*Animate]*Animate
	Animates_referenceOrder map[*Animate]uint

	// insertion point for slice of pointers maps
	OnAfterAnimateCreateCallback GongOnAfterCreateInterface[Animate]
	OnAfterAnimateUpdateCallback GongOnAfterUpdateInterface[Animate]
	OnAfterAnimateDeleteCallback GongOnAfterDeleteInterface[Animate]

	Circles                map[*Circle]struct{}
	Circles_instance       map[*Circle]*Circle
	Circles_mapString      map[string]*Circle
	CircleOrder            uint
	Circle_stagedOrder     map[*Circle]uint
	Circle_orderStaged     map[uint]*Circle
	Circles_reference      map[*Circle]*Circle
	Circles_referenceOrder map[*Circle]uint

	// insertion point for slice of pointers maps
	Circle_Animations_reverseMap map[*Animate]*Circle

	OnAfterCircleCreateCallback GongOnAfterCreateInterface[Circle]
	OnAfterCircleUpdateCallback GongOnAfterUpdateInterface[Circle]
	OnAfterCircleDeleteCallback GongOnAfterDeleteInterface[Circle]

	Conditions                map[*Condition]struct{}
	Conditions_instance       map[*Condition]*Condition
	Conditions_mapString      map[string]*Condition
	ConditionOrder            uint
	Condition_stagedOrder     map[*Condition]uint
	Condition_orderStaged     map[uint]*Condition
	Conditions_reference      map[*Condition]*Condition
	Conditions_referenceOrder map[*Condition]uint

	// insertion point for slice of pointers maps
	OnAfterConditionCreateCallback GongOnAfterCreateInterface[Condition]
	OnAfterConditionUpdateCallback GongOnAfterUpdateInterface[Condition]
	OnAfterConditionDeleteCallback GongOnAfterDeleteInterface[Condition]

	ControlPoints                map[*ControlPoint]struct{}
	ControlPoints_instance       map[*ControlPoint]*ControlPoint
	ControlPoints_mapString      map[string]*ControlPoint
	ControlPointOrder            uint
	ControlPoint_stagedOrder     map[*ControlPoint]uint
	ControlPoint_orderStaged     map[uint]*ControlPoint
	ControlPoints_reference      map[*ControlPoint]*ControlPoint
	ControlPoints_referenceOrder map[*ControlPoint]uint

	// insertion point for slice of pointers maps
	OnAfterControlPointCreateCallback GongOnAfterCreateInterface[ControlPoint]
	OnAfterControlPointUpdateCallback GongOnAfterUpdateInterface[ControlPoint]
	OnAfterControlPointDeleteCallback GongOnAfterDeleteInterface[ControlPoint]

	Ellipses                map[*Ellipse]struct{}
	Ellipses_instance       map[*Ellipse]*Ellipse
	Ellipses_mapString      map[string]*Ellipse
	EllipseOrder            uint
	Ellipse_stagedOrder     map[*Ellipse]uint
	Ellipse_orderStaged     map[uint]*Ellipse
	Ellipses_reference      map[*Ellipse]*Ellipse
	Ellipses_referenceOrder map[*Ellipse]uint

	// insertion point for slice of pointers maps
	Ellipse_Animates_reverseMap map[*Animate]*Ellipse

	OnAfterEllipseCreateCallback GongOnAfterCreateInterface[Ellipse]
	OnAfterEllipseUpdateCallback GongOnAfterUpdateInterface[Ellipse]
	OnAfterEllipseDeleteCallback GongOnAfterDeleteInterface[Ellipse]

	FileToDownloads                map[*FileToDownload]struct{}
	FileToDownloads_instance       map[*FileToDownload]*FileToDownload
	FileToDownloads_mapString      map[string]*FileToDownload
	FileToDownloadOrder            uint
	FileToDownload_stagedOrder     map[*FileToDownload]uint
	FileToDownload_orderStaged     map[uint]*FileToDownload
	FileToDownloads_reference      map[*FileToDownload]*FileToDownload
	FileToDownloads_referenceOrder map[*FileToDownload]uint

	// insertion point for slice of pointers maps
	OnAfterFileToDownloadCreateCallback GongOnAfterCreateInterface[FileToDownload]
	OnAfterFileToDownloadUpdateCallback GongOnAfterUpdateInterface[FileToDownload]
	OnAfterFileToDownloadDeleteCallback GongOnAfterDeleteInterface[FileToDownload]

	Layers                map[*Layer]struct{}
	Layers_instance       map[*Layer]*Layer
	Layers_mapString      map[string]*Layer
	LayerOrder            uint
	Layer_stagedOrder     map[*Layer]uint
	Layer_orderStaged     map[uint]*Layer
	Layers_reference      map[*Layer]*Layer
	Layers_referenceOrder map[*Layer]uint

	// insertion point for slice of pointers maps
	Layer_Rects_reverseMap map[*Rect]*Layer

	Layer_Texts_reverseMap map[*Text]*Layer

	Layer_Circles_reverseMap map[*Circle]*Layer

	Layer_Lines_reverseMap map[*Line]*Layer

	Layer_Ellipses_reverseMap map[*Ellipse]*Layer

	Layer_Polylines_reverseMap map[*Polyline]*Layer

	Layer_Polygones_reverseMap map[*Polygone]*Layer

	Layer_Paths_reverseMap map[*Path]*Layer

	Layer_Links_reverseMap map[*Link]*Layer

	Layer_RectLinkLinks_reverseMap map[*RectLinkLink]*Layer

	OnAfterLayerCreateCallback GongOnAfterCreateInterface[Layer]
	OnAfterLayerUpdateCallback GongOnAfterUpdateInterface[Layer]
	OnAfterLayerDeleteCallback GongOnAfterDeleteInterface[Layer]

	Lines                map[*Line]struct{}
	Lines_instance       map[*Line]*Line
	Lines_mapString      map[string]*Line
	LineOrder            uint
	Line_stagedOrder     map[*Line]uint
	Line_orderStaged     map[uint]*Line
	Lines_reference      map[*Line]*Line
	Lines_referenceOrder map[*Line]uint

	// insertion point for slice of pointers maps
	Line_Animates_reverseMap map[*Animate]*Line

	OnAfterLineCreateCallback GongOnAfterCreateInterface[Line]
	OnAfterLineUpdateCallback GongOnAfterUpdateInterface[Line]
	OnAfterLineDeleteCallback GongOnAfterDeleteInterface[Line]

	Links                map[*Link]struct{}
	Links_instance       map[*Link]*Link
	Links_mapString      map[string]*Link
	LinkOrder            uint
	Link_stagedOrder     map[*Link]uint
	Link_orderStaged     map[uint]*Link
	Links_reference      map[*Link]*Link
	Links_referenceOrder map[*Link]uint

	// insertion point for slice of pointers maps
	Link_TextAtArrowStart_reverseMap map[*LinkAnchoredText]*Link

	Link_TextAtArrowEnd_reverseMap map[*LinkAnchoredText]*Link

	Link_TextAtCorner_reverseMap map[*LinkAnchoredText]*Link

	Link_PathAtArrowStart_reverseMap map[*LinkAnchoredPath]*Link

	Link_PathAtArrowEnd_reverseMap map[*LinkAnchoredPath]*Link

	Link_PathAtCorner_reverseMap map[*LinkAnchoredPath]*Link

	Link_ControlPoints_reverseMap map[*ControlPoint]*Link

	OnAfterLinkCreateCallback GongOnAfterCreateInterface[Link]
	OnAfterLinkUpdateCallback GongOnAfterUpdateInterface[Link]
	OnAfterLinkDeleteCallback GongOnAfterDeleteInterface[Link]

	LinkAnchoredPaths                map[*LinkAnchoredPath]struct{}
	LinkAnchoredPaths_instance       map[*LinkAnchoredPath]*LinkAnchoredPath
	LinkAnchoredPaths_mapString      map[string]*LinkAnchoredPath
	LinkAnchoredPathOrder            uint
	LinkAnchoredPath_stagedOrder     map[*LinkAnchoredPath]uint
	LinkAnchoredPath_orderStaged     map[uint]*LinkAnchoredPath
	LinkAnchoredPaths_reference      map[*LinkAnchoredPath]*LinkAnchoredPath
	LinkAnchoredPaths_referenceOrder map[*LinkAnchoredPath]uint

	// insertion point for slice of pointers maps
	OnAfterLinkAnchoredPathCreateCallback GongOnAfterCreateInterface[LinkAnchoredPath]
	OnAfterLinkAnchoredPathUpdateCallback GongOnAfterUpdateInterface[LinkAnchoredPath]
	OnAfterLinkAnchoredPathDeleteCallback GongOnAfterDeleteInterface[LinkAnchoredPath]

	LinkAnchoredTexts                map[*LinkAnchoredText]struct{}
	LinkAnchoredTexts_instance       map[*LinkAnchoredText]*LinkAnchoredText
	LinkAnchoredTexts_mapString      map[string]*LinkAnchoredText
	LinkAnchoredTextOrder            uint
	LinkAnchoredText_stagedOrder     map[*LinkAnchoredText]uint
	LinkAnchoredText_orderStaged     map[uint]*LinkAnchoredText
	LinkAnchoredTexts_reference      map[*LinkAnchoredText]*LinkAnchoredText
	LinkAnchoredTexts_referenceOrder map[*LinkAnchoredText]uint

	// insertion point for slice of pointers maps
	LinkAnchoredText_Animates_reverseMap map[*Animate]*LinkAnchoredText

	OnAfterLinkAnchoredTextCreateCallback GongOnAfterCreateInterface[LinkAnchoredText]
	OnAfterLinkAnchoredTextUpdateCallback GongOnAfterUpdateInterface[LinkAnchoredText]
	OnAfterLinkAnchoredTextDeleteCallback GongOnAfterDeleteInterface[LinkAnchoredText]

	Paths                map[*Path]struct{}
	Paths_instance       map[*Path]*Path
	Paths_mapString      map[string]*Path
	PathOrder            uint
	Path_stagedOrder     map[*Path]uint
	Path_orderStaged     map[uint]*Path
	Paths_reference      map[*Path]*Path
	Paths_referenceOrder map[*Path]uint

	// insertion point for slice of pointers maps
	Path_Animates_reverseMap map[*Animate]*Path

	OnAfterPathCreateCallback GongOnAfterCreateInterface[Path]
	OnAfterPathUpdateCallback GongOnAfterUpdateInterface[Path]
	OnAfterPathDeleteCallback GongOnAfterDeleteInterface[Path]

	Points                map[*Point]struct{}
	Points_instance       map[*Point]*Point
	Points_mapString      map[string]*Point
	PointOrder            uint
	Point_stagedOrder     map[*Point]uint
	Point_orderStaged     map[uint]*Point
	Points_reference      map[*Point]*Point
	Points_referenceOrder map[*Point]uint

	// insertion point for slice of pointers maps
	OnAfterPointCreateCallback GongOnAfterCreateInterface[Point]
	OnAfterPointUpdateCallback GongOnAfterUpdateInterface[Point]
	OnAfterPointDeleteCallback GongOnAfterDeleteInterface[Point]

	Polygones                map[*Polygone]struct{}
	Polygones_instance       map[*Polygone]*Polygone
	Polygones_mapString      map[string]*Polygone
	PolygoneOrder            uint
	Polygone_stagedOrder     map[*Polygone]uint
	Polygone_orderStaged     map[uint]*Polygone
	Polygones_reference      map[*Polygone]*Polygone
	Polygones_referenceOrder map[*Polygone]uint

	// insertion point for slice of pointers maps
	Polygone_Animates_reverseMap map[*Animate]*Polygone

	OnAfterPolygoneCreateCallback GongOnAfterCreateInterface[Polygone]
	OnAfterPolygoneUpdateCallback GongOnAfterUpdateInterface[Polygone]
	OnAfterPolygoneDeleteCallback GongOnAfterDeleteInterface[Polygone]

	Polylines                map[*Polyline]struct{}
	Polylines_instance       map[*Polyline]*Polyline
	Polylines_mapString      map[string]*Polyline
	PolylineOrder            uint
	Polyline_stagedOrder     map[*Polyline]uint
	Polyline_orderStaged     map[uint]*Polyline
	Polylines_reference      map[*Polyline]*Polyline
	Polylines_referenceOrder map[*Polyline]uint

	// insertion point for slice of pointers maps
	Polyline_Animates_reverseMap map[*Animate]*Polyline

	OnAfterPolylineCreateCallback GongOnAfterCreateInterface[Polyline]
	OnAfterPolylineUpdateCallback GongOnAfterUpdateInterface[Polyline]
	OnAfterPolylineDeleteCallback GongOnAfterDeleteInterface[Polyline]

	Rects                map[*Rect]struct{}
	Rects_instance       map[*Rect]*Rect
	Rects_mapString      map[string]*Rect
	RectOrder            uint
	Rect_stagedOrder     map[*Rect]uint
	Rect_orderStaged     map[uint]*Rect
	Rects_reference      map[*Rect]*Rect
	Rects_referenceOrder map[*Rect]uint

	// insertion point for slice of pointers maps
	Rect_Peers_reverseMap map[*Rect]*Rect

	Rect_Obstacles_reverseMap map[*Rect]*Rect

	Rect_HoveringTrigger_reverseMap map[*Condition]*Rect

	Rect_DisplayConditions_reverseMap map[*Condition]*Rect

	Rect_Animations_reverseMap map[*Animate]*Rect

	Rect_RectAnchoredTexts_reverseMap map[*RectAnchoredText]*Rect

	Rect_RectAnchoredRects_reverseMap map[*RectAnchoredRect]*Rect

	Rect_RectAnchoredPaths_reverseMap map[*RectAnchoredPath]*Rect

	Rect_RectAnchoredPngImages_reverseMap map[*RectAnchoredPngImage]*Rect

	OnAfterRectCreateCallback GongOnAfterCreateInterface[Rect]
	OnAfterRectUpdateCallback GongOnAfterUpdateInterface[Rect]
	OnAfterRectDeleteCallback GongOnAfterDeleteInterface[Rect]

	RectAnchoredPaths                map[*RectAnchoredPath]struct{}
	RectAnchoredPaths_instance       map[*RectAnchoredPath]*RectAnchoredPath
	RectAnchoredPaths_mapString      map[string]*RectAnchoredPath
	RectAnchoredPathOrder            uint
	RectAnchoredPath_stagedOrder     map[*RectAnchoredPath]uint
	RectAnchoredPath_orderStaged     map[uint]*RectAnchoredPath
	RectAnchoredPaths_reference      map[*RectAnchoredPath]*RectAnchoredPath
	RectAnchoredPaths_referenceOrder map[*RectAnchoredPath]uint

	// insertion point for slice of pointers maps
	OnAfterRectAnchoredPathCreateCallback GongOnAfterCreateInterface[RectAnchoredPath]
	OnAfterRectAnchoredPathUpdateCallback GongOnAfterUpdateInterface[RectAnchoredPath]
	OnAfterRectAnchoredPathDeleteCallback GongOnAfterDeleteInterface[RectAnchoredPath]

	RectAnchoredPngImages                map[*RectAnchoredPngImage]struct{}
	RectAnchoredPngImages_instance       map[*RectAnchoredPngImage]*RectAnchoredPngImage
	RectAnchoredPngImages_mapString      map[string]*RectAnchoredPngImage
	RectAnchoredPngImageOrder            uint
	RectAnchoredPngImage_stagedOrder     map[*RectAnchoredPngImage]uint
	RectAnchoredPngImage_orderStaged     map[uint]*RectAnchoredPngImage
	RectAnchoredPngImages_reference      map[*RectAnchoredPngImage]*RectAnchoredPngImage
	RectAnchoredPngImages_referenceOrder map[*RectAnchoredPngImage]uint

	// insertion point for slice of pointers maps
	OnAfterRectAnchoredPngImageCreateCallback GongOnAfterCreateInterface[RectAnchoredPngImage]
	OnAfterRectAnchoredPngImageUpdateCallback GongOnAfterUpdateInterface[RectAnchoredPngImage]
	OnAfterRectAnchoredPngImageDeleteCallback GongOnAfterDeleteInterface[RectAnchoredPngImage]

	RectAnchoredRects                map[*RectAnchoredRect]struct{}
	RectAnchoredRects_instance       map[*RectAnchoredRect]*RectAnchoredRect
	RectAnchoredRects_mapString      map[string]*RectAnchoredRect
	RectAnchoredRectOrder            uint
	RectAnchoredRect_stagedOrder     map[*RectAnchoredRect]uint
	RectAnchoredRect_orderStaged     map[uint]*RectAnchoredRect
	RectAnchoredRects_reference      map[*RectAnchoredRect]*RectAnchoredRect
	RectAnchoredRects_referenceOrder map[*RectAnchoredRect]uint

	// insertion point for slice of pointers maps
	OnAfterRectAnchoredRectCreateCallback GongOnAfterCreateInterface[RectAnchoredRect]
	OnAfterRectAnchoredRectUpdateCallback GongOnAfterUpdateInterface[RectAnchoredRect]
	OnAfterRectAnchoredRectDeleteCallback GongOnAfterDeleteInterface[RectAnchoredRect]

	RectAnchoredTexts                map[*RectAnchoredText]struct{}
	RectAnchoredTexts_instance       map[*RectAnchoredText]*RectAnchoredText
	RectAnchoredTexts_mapString      map[string]*RectAnchoredText
	RectAnchoredTextOrder            uint
	RectAnchoredText_stagedOrder     map[*RectAnchoredText]uint
	RectAnchoredText_orderStaged     map[uint]*RectAnchoredText
	RectAnchoredTexts_reference      map[*RectAnchoredText]*RectAnchoredText
	RectAnchoredTexts_referenceOrder map[*RectAnchoredText]uint

	// insertion point for slice of pointers maps
	RectAnchoredText_Animates_reverseMap map[*Animate]*RectAnchoredText

	OnAfterRectAnchoredTextCreateCallback GongOnAfterCreateInterface[RectAnchoredText]
	OnAfterRectAnchoredTextUpdateCallback GongOnAfterUpdateInterface[RectAnchoredText]
	OnAfterRectAnchoredTextDeleteCallback GongOnAfterDeleteInterface[RectAnchoredText]

	RectLinkLinks                map[*RectLinkLink]struct{}
	RectLinkLinks_instance       map[*RectLinkLink]*RectLinkLink
	RectLinkLinks_mapString      map[string]*RectLinkLink
	RectLinkLinkOrder            uint
	RectLinkLink_stagedOrder     map[*RectLinkLink]uint
	RectLinkLink_orderStaged     map[uint]*RectLinkLink
	RectLinkLinks_reference      map[*RectLinkLink]*RectLinkLink
	RectLinkLinks_referenceOrder map[*RectLinkLink]uint

	// insertion point for slice of pointers maps
	OnAfterRectLinkLinkCreateCallback GongOnAfterCreateInterface[RectLinkLink]
	OnAfterRectLinkLinkUpdateCallback GongOnAfterUpdateInterface[RectLinkLink]
	OnAfterRectLinkLinkDeleteCallback GongOnAfterDeleteInterface[RectLinkLink]

	SVGs                map[*SVG]struct{}
	SVGs_instance       map[*SVG]*SVG
	SVGs_mapString      map[string]*SVG
	SVGOrder            uint
	SVG_stagedOrder     map[*SVG]uint
	SVG_orderStaged     map[uint]*SVG
	SVGs_reference      map[*SVG]*SVG
	SVGs_referenceOrder map[*SVG]uint

	// insertion point for slice of pointers maps
	SVG_Layers_reverseMap map[*Layer]*SVG

	OnAfterSVGCreateCallback GongOnAfterCreateInterface[SVG]
	OnAfterSVGUpdateCallback GongOnAfterUpdateInterface[SVG]
	OnAfterSVGDeleteCallback GongOnAfterDeleteInterface[SVG]

	SvgTexts                map[*SvgText]struct{}
	SvgTexts_instance       map[*SvgText]*SvgText
	SvgTexts_mapString      map[string]*SvgText
	SvgTextOrder            uint
	SvgText_stagedOrder     map[*SvgText]uint
	SvgText_orderStaged     map[uint]*SvgText
	SvgTexts_reference      map[*SvgText]*SvgText
	SvgTexts_referenceOrder map[*SvgText]uint

	// insertion point for slice of pointers maps
	OnAfterSvgTextCreateCallback GongOnAfterCreateInterface[SvgText]
	OnAfterSvgTextUpdateCallback GongOnAfterUpdateInterface[SvgText]
	OnAfterSvgTextDeleteCallback GongOnAfterDeleteInterface[SvgText]

	Texts                map[*Text]struct{}
	Texts_instance       map[*Text]*Text
	Texts_mapString      map[string]*Text
	TextOrder            uint
	Text_stagedOrder     map[*Text]uint
	Text_orderStaged     map[uint]*Text
	Texts_reference      map[*Text]*Text
	Texts_referenceOrder map[*Text]uint

	// insertion point for slice of pointers maps
	Text_Animates_reverseMap map[*Animate]*Text

	OnAfterTextCreateCallback GongOnAfterCreateInterface[Text]
	OnAfterTextUpdateCallback GongOnAfterUpdateInterface[Text]
	OnAfterTextDeleteCallback GongOnAfterDeleteInterface[Text]

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
	__gong__clearReferences(&stage.Animates_reference, &stage.Animates_instance, &stage.Animates_referenceOrder)

	__gong__clearReferences(&stage.Circles_reference, &stage.Circles_instance, &stage.Circles_referenceOrder)

	__gong__clearReferences(&stage.Conditions_reference, &stage.Conditions_instance, &stage.Conditions_referenceOrder)

	__gong__clearReferences(&stage.ControlPoints_reference, &stage.ControlPoints_instance, &stage.ControlPoints_referenceOrder)

	__gong__clearReferences(&stage.Ellipses_reference, &stage.Ellipses_instance, &stage.Ellipses_referenceOrder)

	__gong__clearReferences(&stage.FileToDownloads_reference, &stage.FileToDownloads_instance, &stage.FileToDownloads_referenceOrder)

	__gong__clearReferences(&stage.Layers_reference, &stage.Layers_instance, &stage.Layers_referenceOrder)

	__gong__clearReferences(&stage.Lines_reference, &stage.Lines_instance, &stage.Lines_referenceOrder)

	__gong__clearReferences(&stage.Links_reference, &stage.Links_instance, &stage.Links_referenceOrder)

	__gong__clearReferences(&stage.LinkAnchoredPaths_reference, &stage.LinkAnchoredPaths_instance, &stage.LinkAnchoredPaths_referenceOrder)

	__gong__clearReferences(&stage.LinkAnchoredTexts_reference, &stage.LinkAnchoredTexts_instance, &stage.LinkAnchoredTexts_referenceOrder)

	__gong__clearReferences(&stage.Paths_reference, &stage.Paths_instance, &stage.Paths_referenceOrder)

	__gong__clearReferences(&stage.Points_reference, &stage.Points_instance, &stage.Points_referenceOrder)

	__gong__clearReferences(&stage.Polygones_reference, &stage.Polygones_instance, &stage.Polygones_referenceOrder)

	__gong__clearReferences(&stage.Polylines_reference, &stage.Polylines_instance, &stage.Polylines_referenceOrder)

	__gong__clearReferences(&stage.Rects_reference, &stage.Rects_instance, &stage.Rects_referenceOrder)

	__gong__clearReferences(&stage.RectAnchoredPaths_reference, &stage.RectAnchoredPaths_instance, &stage.RectAnchoredPaths_referenceOrder)

	__gong__clearReferences(&stage.RectAnchoredPngImages_reference, &stage.RectAnchoredPngImages_instance, &stage.RectAnchoredPngImages_referenceOrder)

	__gong__clearReferences(&stage.RectAnchoredRects_reference, &stage.RectAnchoredRects_instance, &stage.RectAnchoredRects_referenceOrder)

	__gong__clearReferences(&stage.RectAnchoredTexts_reference, &stage.RectAnchoredTexts_instance, &stage.RectAnchoredTexts_referenceOrder)

	__gong__clearReferences(&stage.RectLinkLinks_reference, &stage.RectLinkLinks_instance, &stage.RectLinkLinks_referenceOrder)

	__gong__clearReferences(&stage.SVGs_reference, &stage.SVGs_instance, &stage.SVGs_referenceOrder)

	__gong__clearReferences(&stage.SvgTexts_reference, &stage.SvgTexts_instance, &stage.SvgTexts_referenceOrder)

	__gong__clearReferences(&stage.Texts_reference, &stage.Texts_instance, &stage.Texts_referenceOrder)

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
	stage.AnimateOrder = __gong__recomputeOrder(stage.Animate_stagedOrder)

	stage.CircleOrder = __gong__recomputeOrder(stage.Circle_stagedOrder)

	stage.ConditionOrder = __gong__recomputeOrder(stage.Condition_stagedOrder)

	stage.ControlPointOrder = __gong__recomputeOrder(stage.ControlPoint_stagedOrder)

	stage.EllipseOrder = __gong__recomputeOrder(stage.Ellipse_stagedOrder)

	stage.FileToDownloadOrder = __gong__recomputeOrder(stage.FileToDownload_stagedOrder)

	stage.LayerOrder = __gong__recomputeOrder(stage.Layer_stagedOrder)

	stage.LineOrder = __gong__recomputeOrder(stage.Line_stagedOrder)

	stage.LinkOrder = __gong__recomputeOrder(stage.Link_stagedOrder)

	stage.LinkAnchoredPathOrder = __gong__recomputeOrder(stage.LinkAnchoredPath_stagedOrder)

	stage.LinkAnchoredTextOrder = __gong__recomputeOrder(stage.LinkAnchoredText_stagedOrder)

	stage.PathOrder = __gong__recomputeOrder(stage.Path_stagedOrder)

	stage.PointOrder = __gong__recomputeOrder(stage.Point_stagedOrder)

	stage.PolygoneOrder = __gong__recomputeOrder(stage.Polygone_stagedOrder)

	stage.PolylineOrder = __gong__recomputeOrder(stage.Polyline_stagedOrder)

	stage.RectOrder = __gong__recomputeOrder(stage.Rect_stagedOrder)

	stage.RectAnchoredPathOrder = __gong__recomputeOrder(stage.RectAnchoredPath_stagedOrder)

	stage.RectAnchoredPngImageOrder = __gong__recomputeOrder(stage.RectAnchoredPngImage_stagedOrder)

	stage.RectAnchoredRectOrder = __gong__recomputeOrder(stage.RectAnchoredRect_stagedOrder)

	stage.RectAnchoredTextOrder = __gong__recomputeOrder(stage.RectAnchoredText_stagedOrder)

	stage.RectLinkLinkOrder = __gong__recomputeOrder(stage.RectLinkLink_stagedOrder)

	stage.SVGOrder = __gong__recomputeOrder(stage.SVG_stagedOrder)

	stage.SvgTextOrder = __gong__recomputeOrder(stage.SvgText_stagedOrder)

	stage.TextOrder = __gong__recomputeOrder(stage.Text_stagedOrder)

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
	case *Animate:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Animates, stage.Animate_stagedOrder))
	case *Circle:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Circles, stage.Circle_stagedOrder))
	case *Condition:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Conditions, stage.Condition_stagedOrder))
	case *ControlPoint:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ControlPoints, stage.ControlPoint_stagedOrder))
	case *Ellipse:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Ellipses, stage.Ellipse_stagedOrder))
	case *FileToDownload:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.FileToDownloads, stage.FileToDownload_stagedOrder))
	case *Layer:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Layers, stage.Layer_stagedOrder))
	case *Line:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Lines, stage.Line_stagedOrder))
	case *Link:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Links, stage.Link_stagedOrder))
	case *LinkAnchoredPath:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.LinkAnchoredPaths, stage.LinkAnchoredPath_stagedOrder))
	case *LinkAnchoredText:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.LinkAnchoredTexts, stage.LinkAnchoredText_stagedOrder))
	case *Path:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Paths, stage.Path_stagedOrder))
	case *Point:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Points, stage.Point_stagedOrder))
	case *Polygone:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Polygones, stage.Polygone_stagedOrder))
	case *Polyline:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Polylines, stage.Polyline_stagedOrder))
	case *Rect:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Rects, stage.Rect_stagedOrder))
	case *RectAnchoredPath:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.RectAnchoredPaths, stage.RectAnchoredPath_stagedOrder))
	case *RectAnchoredPngImage:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.RectAnchoredPngImages, stage.RectAnchoredPngImage_stagedOrder))
	case *RectAnchoredRect:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.RectAnchoredRects, stage.RectAnchoredRect_stagedOrder))
	case *RectAnchoredText:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.RectAnchoredTexts, stage.RectAnchoredText_stagedOrder))
	case *RectLinkLink:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.RectLinkLinks, stage.RectLinkLink_stagedOrder))
	case *SVG:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.SVGs, stage.SVG_stagedOrder))
	case *SvgText:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.SvgTexts, stage.SvgText_stagedOrder))
	case *Text:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Texts, stage.Text_stagedOrder))

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
	return "github.com/fullstack-lang/gong/lib/svg/go/models"
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
		Animates:           make(map[*Animate]struct{}),
		Animates_mapString: make(map[string]*Animate),

		Circles:           make(map[*Circle]struct{}),
		Circles_mapString: make(map[string]*Circle),

		Conditions:           make(map[*Condition]struct{}),
		Conditions_mapString: make(map[string]*Condition),

		ControlPoints:           make(map[*ControlPoint]struct{}),
		ControlPoints_mapString: make(map[string]*ControlPoint),

		Ellipses:           make(map[*Ellipse]struct{}),
		Ellipses_mapString: make(map[string]*Ellipse),

		FileToDownloads:           make(map[*FileToDownload]struct{}),
		FileToDownloads_mapString: make(map[string]*FileToDownload),

		Layers:           make(map[*Layer]struct{}),
		Layers_mapString: make(map[string]*Layer),

		Lines:           make(map[*Line]struct{}),
		Lines_mapString: make(map[string]*Line),

		Links:           make(map[*Link]struct{}),
		Links_mapString: make(map[string]*Link),

		LinkAnchoredPaths:           make(map[*LinkAnchoredPath]struct{}),
		LinkAnchoredPaths_mapString: make(map[string]*LinkAnchoredPath),

		LinkAnchoredTexts:           make(map[*LinkAnchoredText]struct{}),
		LinkAnchoredTexts_mapString: make(map[string]*LinkAnchoredText),

		Paths:           make(map[*Path]struct{}),
		Paths_mapString: make(map[string]*Path),

		Points:           make(map[*Point]struct{}),
		Points_mapString: make(map[string]*Point),

		Polygones:           make(map[*Polygone]struct{}),
		Polygones_mapString: make(map[string]*Polygone),

		Polylines:           make(map[*Polyline]struct{}),
		Polylines_mapString: make(map[string]*Polyline),

		Rects:           make(map[*Rect]struct{}),
		Rects_mapString: make(map[string]*Rect),

		RectAnchoredPaths:           make(map[*RectAnchoredPath]struct{}),
		RectAnchoredPaths_mapString: make(map[string]*RectAnchoredPath),

		RectAnchoredPngImages:           make(map[*RectAnchoredPngImage]struct{}),
		RectAnchoredPngImages_mapString: make(map[string]*RectAnchoredPngImage),

		RectAnchoredRects:           make(map[*RectAnchoredRect]struct{}),
		RectAnchoredRects_mapString: make(map[string]*RectAnchoredRect),

		RectAnchoredTexts:           make(map[*RectAnchoredText]struct{}),
		RectAnchoredTexts_mapString: make(map[string]*RectAnchoredText),

		RectLinkLinks:           make(map[*RectLinkLink]struct{}),
		RectLinkLinks_mapString: make(map[string]*RectLinkLink),

		SVGs:           make(map[*SVG]struct{}),
		SVGs_mapString: make(map[string]*SVG),

		SvgTexts:           make(map[*SvgText]struct{}),
		SvgTexts_mapString: make(map[string]*SvgText),

		Texts:           make(map[*Text]struct{}),
		Texts_mapString: make(map[string]*Text),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		Animate_stagedOrder: make(map[*Animate]uint),
		Animate_orderStaged: make(map[uint]*Animate),
		Animates_reference:  make(map[*Animate]*Animate),

		Circle_stagedOrder: make(map[*Circle]uint),
		Circle_orderStaged: make(map[uint]*Circle),
		Circles_reference:  make(map[*Circle]*Circle),

		Condition_stagedOrder: make(map[*Condition]uint),
		Condition_orderStaged: make(map[uint]*Condition),
		Conditions_reference:  make(map[*Condition]*Condition),

		ControlPoint_stagedOrder: make(map[*ControlPoint]uint),
		ControlPoint_orderStaged: make(map[uint]*ControlPoint),
		ControlPoints_reference:  make(map[*ControlPoint]*ControlPoint),

		Ellipse_stagedOrder: make(map[*Ellipse]uint),
		Ellipse_orderStaged: make(map[uint]*Ellipse),
		Ellipses_reference:  make(map[*Ellipse]*Ellipse),

		FileToDownload_stagedOrder: make(map[*FileToDownload]uint),
		FileToDownload_orderStaged: make(map[uint]*FileToDownload),
		FileToDownloads_reference:  make(map[*FileToDownload]*FileToDownload),

		Layer_stagedOrder: make(map[*Layer]uint),
		Layer_orderStaged: make(map[uint]*Layer),
		Layers_reference:  make(map[*Layer]*Layer),

		Line_stagedOrder: make(map[*Line]uint),
		Line_orderStaged: make(map[uint]*Line),
		Lines_reference:  make(map[*Line]*Line),

		Link_stagedOrder: make(map[*Link]uint),
		Link_orderStaged: make(map[uint]*Link),
		Links_reference:  make(map[*Link]*Link),

		LinkAnchoredPath_stagedOrder: make(map[*LinkAnchoredPath]uint),
		LinkAnchoredPath_orderStaged: make(map[uint]*LinkAnchoredPath),
		LinkAnchoredPaths_reference:  make(map[*LinkAnchoredPath]*LinkAnchoredPath),

		LinkAnchoredText_stagedOrder: make(map[*LinkAnchoredText]uint),
		LinkAnchoredText_orderStaged: make(map[uint]*LinkAnchoredText),
		LinkAnchoredTexts_reference:  make(map[*LinkAnchoredText]*LinkAnchoredText),

		Path_stagedOrder: make(map[*Path]uint),
		Path_orderStaged: make(map[uint]*Path),
		Paths_reference:  make(map[*Path]*Path),

		Point_stagedOrder: make(map[*Point]uint),
		Point_orderStaged: make(map[uint]*Point),
		Points_reference:  make(map[*Point]*Point),

		Polygone_stagedOrder: make(map[*Polygone]uint),
		Polygone_orderStaged: make(map[uint]*Polygone),
		Polygones_reference:  make(map[*Polygone]*Polygone),

		Polyline_stagedOrder: make(map[*Polyline]uint),
		Polyline_orderStaged: make(map[uint]*Polyline),
		Polylines_reference:  make(map[*Polyline]*Polyline),

		Rect_stagedOrder: make(map[*Rect]uint),
		Rect_orderStaged: make(map[uint]*Rect),
		Rects_reference:  make(map[*Rect]*Rect),

		RectAnchoredPath_stagedOrder: make(map[*RectAnchoredPath]uint),
		RectAnchoredPath_orderStaged: make(map[uint]*RectAnchoredPath),
		RectAnchoredPaths_reference:  make(map[*RectAnchoredPath]*RectAnchoredPath),

		RectAnchoredPngImage_stagedOrder: make(map[*RectAnchoredPngImage]uint),
		RectAnchoredPngImage_orderStaged: make(map[uint]*RectAnchoredPngImage),
		RectAnchoredPngImages_reference:  make(map[*RectAnchoredPngImage]*RectAnchoredPngImage),

		RectAnchoredRect_stagedOrder: make(map[*RectAnchoredRect]uint),
		RectAnchoredRect_orderStaged: make(map[uint]*RectAnchoredRect),
		RectAnchoredRects_reference:  make(map[*RectAnchoredRect]*RectAnchoredRect),

		RectAnchoredText_stagedOrder: make(map[*RectAnchoredText]uint),
		RectAnchoredText_orderStaged: make(map[uint]*RectAnchoredText),
		RectAnchoredTexts_reference:  make(map[*RectAnchoredText]*RectAnchoredText),

		RectLinkLink_stagedOrder: make(map[*RectLinkLink]uint),
		RectLinkLink_orderStaged: make(map[uint]*RectLinkLink),
		RectLinkLinks_reference:  make(map[*RectLinkLink]*RectLinkLink),

		SVG_stagedOrder: make(map[*SVG]uint),
		SVG_orderStaged: make(map[uint]*SVG),
		SVGs_reference:  make(map[*SVG]*SVG),

		SvgText_stagedOrder: make(map[*SvgText]uint),
		SvgText_orderStaged: make(map[uint]*SvgText),
		SvgTexts_reference:  make(map[*SvgText]*SvgText),

		Text_stagedOrder: make(map[*Text]uint),
		Text_orderStaged: make(map[uint]*Text),
		Texts_reference:  make(map[*Text]*Text),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"Animate": &AnimateUnmarshaller{},

			"Circle": &CircleUnmarshaller{},

			"Condition": &ConditionUnmarshaller{},

			"ControlPoint": &ControlPointUnmarshaller{},

			"Ellipse": &EllipseUnmarshaller{},

			"FileToDownload": &FileToDownloadUnmarshaller{},

			"Layer": &LayerUnmarshaller{},

			"Line": &LineUnmarshaller{},

			"Link": &LinkUnmarshaller{},

			"LinkAnchoredPath": &LinkAnchoredPathUnmarshaller{},

			"LinkAnchoredText": &LinkAnchoredTextUnmarshaller{},

			"Path": &PathUnmarshaller{},

			"Point": &PointUnmarshaller{},

			"Polygone": &PolygoneUnmarshaller{},

			"Polyline": &PolylineUnmarshaller{},

			"Rect": &RectUnmarshaller{},

			"RectAnchoredPath": &RectAnchoredPathUnmarshaller{},

			"RectAnchoredPngImage": &RectAnchoredPngImageUnmarshaller{},

			"RectAnchoredRect": &RectAnchoredRectUnmarshaller{},

			"RectAnchoredText": &RectAnchoredTextUnmarshaller{},

			"RectLinkLink": &RectLinkLinkUnmarshaller{},

			"SVG": &SVGUnmarshaller{},

			"SvgText": &SvgTextUnmarshaller{},

			"Text": &TextUnmarshaller{},

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
	case *Animate:
		return any(stage.Animate_orderStaged[order]).(Type)
	case *Circle:
		return any(stage.Circle_orderStaged[order]).(Type)
	case *Condition:
		return any(stage.Condition_orderStaged[order]).(Type)
	case *ControlPoint:
		return any(stage.ControlPoint_orderStaged[order]).(Type)
	case *Ellipse:
		return any(stage.Ellipse_orderStaged[order]).(Type)
	case *FileToDownload:
		return any(stage.FileToDownload_orderStaged[order]).(Type)
	case *Layer:
		return any(stage.Layer_orderStaged[order]).(Type)
	case *Line:
		return any(stage.Line_orderStaged[order]).(Type)
	case *Link:
		return any(stage.Link_orderStaged[order]).(Type)
	case *LinkAnchoredPath:
		return any(stage.LinkAnchoredPath_orderStaged[order]).(Type)
	case *LinkAnchoredText:
		return any(stage.LinkAnchoredText_orderStaged[order]).(Type)
	case *Path:
		return any(stage.Path_orderStaged[order]).(Type)
	case *Point:
		return any(stage.Point_orderStaged[order]).(Type)
	case *Polygone:
		return any(stage.Polygone_orderStaged[order]).(Type)
	case *Polyline:
		return any(stage.Polyline_orderStaged[order]).(Type)
	case *Rect:
		return any(stage.Rect_orderStaged[order]).(Type)
	case *RectAnchoredPath:
		return any(stage.RectAnchoredPath_orderStaged[order]).(Type)
	case *RectAnchoredPngImage:
		return any(stage.RectAnchoredPngImage_orderStaged[order]).(Type)
	case *RectAnchoredRect:
		return any(stage.RectAnchoredRect_orderStaged[order]).(Type)
	case *RectAnchoredText:
		return any(stage.RectAnchoredText_orderStaged[order]).(Type)
	case *RectLinkLink:
		return any(stage.RectLinkLink_orderStaged[order]).(Type)
	case *SVG:
		return any(stage.SVG_orderStaged[order]).(Type)
	case *SvgText:
		return any(stage.SvgText_orderStaged[order]).(Type)
	case *Text:
		return any(stage.Text_orderStaged[order]).(Type)
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
	stage.Map_GongStructName_InstancesNb["Animate"] = len(stage.Animates)
	stage.Map_GongStructName_InstancesNb["Circle"] = len(stage.Circles)
	stage.Map_GongStructName_InstancesNb["Condition"] = len(stage.Conditions)
	stage.Map_GongStructName_InstancesNb["ControlPoint"] = len(stage.ControlPoints)
	stage.Map_GongStructName_InstancesNb["Ellipse"] = len(stage.Ellipses)
	stage.Map_GongStructName_InstancesNb["FileToDownload"] = len(stage.FileToDownloads)
	stage.Map_GongStructName_InstancesNb["Layer"] = len(stage.Layers)
	stage.Map_GongStructName_InstancesNb["Line"] = len(stage.Lines)
	stage.Map_GongStructName_InstancesNb["Link"] = len(stage.Links)
	stage.Map_GongStructName_InstancesNb["LinkAnchoredPath"] = len(stage.LinkAnchoredPaths)
	stage.Map_GongStructName_InstancesNb["LinkAnchoredText"] = len(stage.LinkAnchoredTexts)
	stage.Map_GongStructName_InstancesNb["Path"] = len(stage.Paths)
	stage.Map_GongStructName_InstancesNb["Point"] = len(stage.Points)
	stage.Map_GongStructName_InstancesNb["Polygone"] = len(stage.Polygones)
	stage.Map_GongStructName_InstancesNb["Polyline"] = len(stage.Polylines)
	stage.Map_GongStructName_InstancesNb["Rect"] = len(stage.Rects)
	stage.Map_GongStructName_InstancesNb["RectAnchoredPath"] = len(stage.RectAnchoredPaths)
	stage.Map_GongStructName_InstancesNb["RectAnchoredPngImage"] = len(stage.RectAnchoredPngImages)
	stage.Map_GongStructName_InstancesNb["RectAnchoredRect"] = len(stage.RectAnchoredRects)
	stage.Map_GongStructName_InstancesNb["RectAnchoredText"] = len(stage.RectAnchoredTexts)
	stage.Map_GongStructName_InstancesNb["RectLinkLink"] = len(stage.RectLinkLinks)
	stage.Map_GongStructName_InstancesNb["SVG"] = len(stage.SVGs)
	stage.Map_GongStructName_InstancesNb["SvgText"] = len(stage.SvgTexts)
	stage.Map_GongStructName_InstancesNb["Text"] = len(stage.Texts)
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
// Stage puts animate to the model stage
func (animate *Animate) Stage(stage *Stage) *Animate {
	__gong__stage(stage.Animates, stage.Animate_stagedOrder, stage.Animate_orderStaged, &stage.AnimateOrder, stage.Animates_mapString, animate, animate.Name)
	return animate
}

// StagePreserveOrder puts animate to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AnimateOrder
// - update stage.AnimateOrder accordingly
func (animate *Animate) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Animates, stage.Animate_stagedOrder, stage.Animate_orderStaged, &stage.AnimateOrder, stage.Animates_mapString, animate, order, animate.Name)
}

// Unstage removes animate off the model stage
func (animate *Animate) Unstage(stage *Stage) *Animate {
	__gong__unstage(stage.Animates, stage.Animates_mapString, animate, animate.Name)
	return animate
}

// UnstageVoid removes animate off the model stage
func (animate *Animate) UnstageVoid(stage *Stage) {
	animate.Unstage(stage)
}

func (animate *Animate) StageVoid(stage *Stage) {
	animate.Stage(stage)
}

// for satisfaction of GongStruct interface
func (animate *Animate) GetName() (res string) {
	return animate.Name
}

// for satisfaction of GongStruct interface
func (animate *Animate) SetName(name string) {
	animate.Name = name
}

// Stage puts circle to the model stage
func (circle *Circle) Stage(stage *Stage) *Circle {
	__gong__stage(stage.Circles, stage.Circle_stagedOrder, stage.Circle_orderStaged, &stage.CircleOrder, stage.Circles_mapString, circle, circle.Name)
	return circle
}

// StagePreserveOrder puts circle to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CircleOrder
// - update stage.CircleOrder accordingly
func (circle *Circle) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Circles, stage.Circle_stagedOrder, stage.Circle_orderStaged, &stage.CircleOrder, stage.Circles_mapString, circle, order, circle.Name)
}

// Unstage removes circle off the model stage
func (circle *Circle) Unstage(stage *Stage) *Circle {
	__gong__unstage(stage.Circles, stage.Circles_mapString, circle, circle.Name)
	return circle
}

// UnstageVoid removes circle off the model stage
func (circle *Circle) UnstageVoid(stage *Stage) {
	circle.Unstage(stage)
}

func (circle *Circle) StageVoid(stage *Stage) {
	circle.Stage(stage)
}

// for satisfaction of GongStruct interface
func (circle *Circle) GetName() (res string) {
	return circle.Name
}

// for satisfaction of GongStruct interface
func (circle *Circle) SetName(name string) {
	circle.Name = name
}

// Stage puts condition to the model stage
func (condition *Condition) Stage(stage *Stage) *Condition {
	__gong__stage(stage.Conditions, stage.Condition_stagedOrder, stage.Condition_orderStaged, &stage.ConditionOrder, stage.Conditions_mapString, condition, condition.Name)
	return condition
}

// StagePreserveOrder puts condition to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ConditionOrder
// - update stage.ConditionOrder accordingly
func (condition *Condition) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Conditions, stage.Condition_stagedOrder, stage.Condition_orderStaged, &stage.ConditionOrder, stage.Conditions_mapString, condition, order, condition.Name)
}

// Unstage removes condition off the model stage
func (condition *Condition) Unstage(stage *Stage) *Condition {
	__gong__unstage(stage.Conditions, stage.Conditions_mapString, condition, condition.Name)
	return condition
}

// UnstageVoid removes condition off the model stage
func (condition *Condition) UnstageVoid(stage *Stage) {
	condition.Unstage(stage)
}

func (condition *Condition) StageVoid(stage *Stage) {
	condition.Stage(stage)
}

// for satisfaction of GongStruct interface
func (condition *Condition) GetName() (res string) {
	return condition.Name
}

// for satisfaction of GongStruct interface
func (condition *Condition) SetName(name string) {
	condition.Name = name
}

// Stage puts controlpoint to the model stage
func (controlpoint *ControlPoint) Stage(stage *Stage) *ControlPoint {
	__gong__stage(stage.ControlPoints, stage.ControlPoint_stagedOrder, stage.ControlPoint_orderStaged, &stage.ControlPointOrder, stage.ControlPoints_mapString, controlpoint, controlpoint.Name)
	return controlpoint
}

// StagePreserveOrder puts controlpoint to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ControlPointOrder
// - update stage.ControlPointOrder accordingly
func (controlpoint *ControlPoint) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ControlPoints, stage.ControlPoint_stagedOrder, stage.ControlPoint_orderStaged, &stage.ControlPointOrder, stage.ControlPoints_mapString, controlpoint, order, controlpoint.Name)
}

// Unstage removes controlpoint off the model stage
func (controlpoint *ControlPoint) Unstage(stage *Stage) *ControlPoint {
	__gong__unstage(stage.ControlPoints, stage.ControlPoints_mapString, controlpoint, controlpoint.Name)
	return controlpoint
}

// UnstageVoid removes controlpoint off the model stage
func (controlpoint *ControlPoint) UnstageVoid(stage *Stage) {
	controlpoint.Unstage(stage)
}

func (controlpoint *ControlPoint) StageVoid(stage *Stage) {
	controlpoint.Stage(stage)
}

// for satisfaction of GongStruct interface
func (controlpoint *ControlPoint) GetName() (res string) {
	return controlpoint.Name
}

// for satisfaction of GongStruct interface
func (controlpoint *ControlPoint) SetName(name string) {
	controlpoint.Name = name
}

// Stage puts ellipse to the model stage
func (ellipse *Ellipse) Stage(stage *Stage) *Ellipse {
	__gong__stage(stage.Ellipses, stage.Ellipse_stagedOrder, stage.Ellipse_orderStaged, &stage.EllipseOrder, stage.Ellipses_mapString, ellipse, ellipse.Name)
	return ellipse
}

// StagePreserveOrder puts ellipse to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EllipseOrder
// - update stage.EllipseOrder accordingly
func (ellipse *Ellipse) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Ellipses, stage.Ellipse_stagedOrder, stage.Ellipse_orderStaged, &stage.EllipseOrder, stage.Ellipses_mapString, ellipse, order, ellipse.Name)
}

// Unstage removes ellipse off the model stage
func (ellipse *Ellipse) Unstage(stage *Stage) *Ellipse {
	__gong__unstage(stage.Ellipses, stage.Ellipses_mapString, ellipse, ellipse.Name)
	return ellipse
}

// UnstageVoid removes ellipse off the model stage
func (ellipse *Ellipse) UnstageVoid(stage *Stage) {
	ellipse.Unstage(stage)
}

func (ellipse *Ellipse) StageVoid(stage *Stage) {
	ellipse.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ellipse *Ellipse) GetName() (res string) {
	return ellipse.Name
}

// for satisfaction of GongStruct interface
func (ellipse *Ellipse) SetName(name string) {
	ellipse.Name = name
}

// Stage puts filetodownload to the model stage
func (filetodownload *FileToDownload) Stage(stage *Stage) *FileToDownload {
	__gong__stage(stage.FileToDownloads, stage.FileToDownload_stagedOrder, stage.FileToDownload_orderStaged, &stage.FileToDownloadOrder, stage.FileToDownloads_mapString, filetodownload, filetodownload.Name)
	return filetodownload
}

// StagePreserveOrder puts filetodownload to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FileToDownloadOrder
// - update stage.FileToDownloadOrder accordingly
func (filetodownload *FileToDownload) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FileToDownloads, stage.FileToDownload_stagedOrder, stage.FileToDownload_orderStaged, &stage.FileToDownloadOrder, stage.FileToDownloads_mapString, filetodownload, order, filetodownload.Name)
}

// Unstage removes filetodownload off the model stage
func (filetodownload *FileToDownload) Unstage(stage *Stage) *FileToDownload {
	__gong__unstage(stage.FileToDownloads, stage.FileToDownloads_mapString, filetodownload, filetodownload.Name)
	return filetodownload
}

// UnstageVoid removes filetodownload off the model stage
func (filetodownload *FileToDownload) UnstageVoid(stage *Stage) {
	filetodownload.Unstage(stage)
}

func (filetodownload *FileToDownload) StageVoid(stage *Stage) {
	filetodownload.Stage(stage)
}

// for satisfaction of GongStruct interface
func (filetodownload *FileToDownload) GetName() (res string) {
	return filetodownload.Name
}

// for satisfaction of GongStruct interface
func (filetodownload *FileToDownload) SetName(name string) {
	filetodownload.Name = name
}

// Stage puts layer to the model stage
func (layer *Layer) Stage(stage *Stage) *Layer {
	__gong__stage(stage.Layers, stage.Layer_stagedOrder, stage.Layer_orderStaged, &stage.LayerOrder, stage.Layers_mapString, layer, layer.Name)
	return layer
}

// StagePreserveOrder puts layer to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LayerOrder
// - update stage.LayerOrder accordingly
func (layer *Layer) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Layers, stage.Layer_stagedOrder, stage.Layer_orderStaged, &stage.LayerOrder, stage.Layers_mapString, layer, order, layer.Name)
}

// Unstage removes layer off the model stage
func (layer *Layer) Unstage(stage *Stage) *Layer {
	__gong__unstage(stage.Layers, stage.Layers_mapString, layer, layer.Name)
	return layer
}

// UnstageVoid removes layer off the model stage
func (layer *Layer) UnstageVoid(stage *Stage) {
	layer.Unstage(stage)
}

func (layer *Layer) StageVoid(stage *Stage) {
	layer.Stage(stage)
}

// for satisfaction of GongStruct interface
func (layer *Layer) GetName() (res string) {
	return layer.Name
}

// for satisfaction of GongStruct interface
func (layer *Layer) SetName(name string) {
	layer.Name = name
}

// Stage puts line to the model stage
func (line *Line) Stage(stage *Stage) *Line {
	__gong__stage(stage.Lines, stage.Line_stagedOrder, stage.Line_orderStaged, &stage.LineOrder, stage.Lines_mapString, line, line.Name)
	return line
}

// StagePreserveOrder puts line to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LineOrder
// - update stage.LineOrder accordingly
func (line *Line) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Lines, stage.Line_stagedOrder, stage.Line_orderStaged, &stage.LineOrder, stage.Lines_mapString, line, order, line.Name)
}

// Unstage removes line off the model stage
func (line *Line) Unstage(stage *Stage) *Line {
	__gong__unstage(stage.Lines, stage.Lines_mapString, line, line.Name)
	return line
}

// UnstageVoid removes line off the model stage
func (line *Line) UnstageVoid(stage *Stage) {
	line.Unstage(stage)
}

func (line *Line) StageVoid(stage *Stage) {
	line.Stage(stage)
}

// for satisfaction of GongStruct interface
func (line *Line) GetName() (res string) {
	return line.Name
}

// for satisfaction of GongStruct interface
func (line *Line) SetName(name string) {
	line.Name = name
}

// Stage puts link to the model stage
func (link *Link) Stage(stage *Stage) *Link {
	__gong__stage(stage.Links, stage.Link_stagedOrder, stage.Link_orderStaged, &stage.LinkOrder, stage.Links_mapString, link, link.Name)
	return link
}

// StagePreserveOrder puts link to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LinkOrder
// - update stage.LinkOrder accordingly
func (link *Link) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Links, stage.Link_stagedOrder, stage.Link_orderStaged, &stage.LinkOrder, stage.Links_mapString, link, order, link.Name)
}

// Unstage removes link off the model stage
func (link *Link) Unstage(stage *Stage) *Link {
	__gong__unstage(stage.Links, stage.Links_mapString, link, link.Name)
	return link
}

// UnstageVoid removes link off the model stage
func (link *Link) UnstageVoid(stage *Stage) {
	link.Unstage(stage)
}

func (link *Link) StageVoid(stage *Stage) {
	link.Stage(stage)
}

// for satisfaction of GongStruct interface
func (link *Link) GetName() (res string) {
	return link.Name
}

// for satisfaction of GongStruct interface
func (link *Link) SetName(name string) {
	link.Name = name
}

// Stage puts linkanchoredpath to the model stage
func (linkanchoredpath *LinkAnchoredPath) Stage(stage *Stage) *LinkAnchoredPath {
	__gong__stage(stage.LinkAnchoredPaths, stage.LinkAnchoredPath_stagedOrder, stage.LinkAnchoredPath_orderStaged, &stage.LinkAnchoredPathOrder, stage.LinkAnchoredPaths_mapString, linkanchoredpath, linkanchoredpath.Name)
	return linkanchoredpath
}

// StagePreserveOrder puts linkanchoredpath to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LinkAnchoredPathOrder
// - update stage.LinkAnchoredPathOrder accordingly
func (linkanchoredpath *LinkAnchoredPath) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.LinkAnchoredPaths, stage.LinkAnchoredPath_stagedOrder, stage.LinkAnchoredPath_orderStaged, &stage.LinkAnchoredPathOrder, stage.LinkAnchoredPaths_mapString, linkanchoredpath, order, linkanchoredpath.Name)
}

// Unstage removes linkanchoredpath off the model stage
func (linkanchoredpath *LinkAnchoredPath) Unstage(stage *Stage) *LinkAnchoredPath {
	__gong__unstage(stage.LinkAnchoredPaths, stage.LinkAnchoredPaths_mapString, linkanchoredpath, linkanchoredpath.Name)
	return linkanchoredpath
}

// UnstageVoid removes linkanchoredpath off the model stage
func (linkanchoredpath *LinkAnchoredPath) UnstageVoid(stage *Stage) {
	linkanchoredpath.Unstage(stage)
}

func (linkanchoredpath *LinkAnchoredPath) StageVoid(stage *Stage) {
	linkanchoredpath.Stage(stage)
}

// for satisfaction of GongStruct interface
func (linkanchoredpath *LinkAnchoredPath) GetName() (res string) {
	return linkanchoredpath.Name
}

// for satisfaction of GongStruct interface
func (linkanchoredpath *LinkAnchoredPath) SetName(name string) {
	linkanchoredpath.Name = name
}

// Stage puts linkanchoredtext to the model stage
func (linkanchoredtext *LinkAnchoredText) Stage(stage *Stage) *LinkAnchoredText {
	__gong__stage(stage.LinkAnchoredTexts, stage.LinkAnchoredText_stagedOrder, stage.LinkAnchoredText_orderStaged, &stage.LinkAnchoredTextOrder, stage.LinkAnchoredTexts_mapString, linkanchoredtext, linkanchoredtext.Name)
	return linkanchoredtext
}

// StagePreserveOrder puts linkanchoredtext to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LinkAnchoredTextOrder
// - update stage.LinkAnchoredTextOrder accordingly
func (linkanchoredtext *LinkAnchoredText) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.LinkAnchoredTexts, stage.LinkAnchoredText_stagedOrder, stage.LinkAnchoredText_orderStaged, &stage.LinkAnchoredTextOrder, stage.LinkAnchoredTexts_mapString, linkanchoredtext, order, linkanchoredtext.Name)
}

// Unstage removes linkanchoredtext off the model stage
func (linkanchoredtext *LinkAnchoredText) Unstage(stage *Stage) *LinkAnchoredText {
	__gong__unstage(stage.LinkAnchoredTexts, stage.LinkAnchoredTexts_mapString, linkanchoredtext, linkanchoredtext.Name)
	return linkanchoredtext
}

// UnstageVoid removes linkanchoredtext off the model stage
func (linkanchoredtext *LinkAnchoredText) UnstageVoid(stage *Stage) {
	linkanchoredtext.Unstage(stage)
}

func (linkanchoredtext *LinkAnchoredText) StageVoid(stage *Stage) {
	linkanchoredtext.Stage(stage)
}

// for satisfaction of GongStruct interface
func (linkanchoredtext *LinkAnchoredText) GetName() (res string) {
	return linkanchoredtext.Name
}

// for satisfaction of GongStruct interface
func (linkanchoredtext *LinkAnchoredText) SetName(name string) {
	linkanchoredtext.Name = name
}

// Stage puts path to the model stage
func (path *Path) Stage(stage *Stage) *Path {
	__gong__stage(stage.Paths, stage.Path_stagedOrder, stage.Path_orderStaged, &stage.PathOrder, stage.Paths_mapString, path, path.Name)
	return path
}

// StagePreserveOrder puts path to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PathOrder
// - update stage.PathOrder accordingly
func (path *Path) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Paths, stage.Path_stagedOrder, stage.Path_orderStaged, &stage.PathOrder, stage.Paths_mapString, path, order, path.Name)
}

// Unstage removes path off the model stage
func (path *Path) Unstage(stage *Stage) *Path {
	__gong__unstage(stage.Paths, stage.Paths_mapString, path, path.Name)
	return path
}

// UnstageVoid removes path off the model stage
func (path *Path) UnstageVoid(stage *Stage) {
	path.Unstage(stage)
}

func (path *Path) StageVoid(stage *Stage) {
	path.Stage(stage)
}

// for satisfaction of GongStruct interface
func (path *Path) GetName() (res string) {
	return path.Name
}

// for satisfaction of GongStruct interface
func (path *Path) SetName(name string) {
	path.Name = name
}

// Stage puts point to the model stage
func (point *Point) Stage(stage *Stage) *Point {
	__gong__stage(stage.Points, stage.Point_stagedOrder, stage.Point_orderStaged, &stage.PointOrder, stage.Points_mapString, point, point.Name)
	return point
}

// StagePreserveOrder puts point to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PointOrder
// - update stage.PointOrder accordingly
func (point *Point) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Points, stage.Point_stagedOrder, stage.Point_orderStaged, &stage.PointOrder, stage.Points_mapString, point, order, point.Name)
}

// Unstage removes point off the model stage
func (point *Point) Unstage(stage *Stage) *Point {
	__gong__unstage(stage.Points, stage.Points_mapString, point, point.Name)
	return point
}

// UnstageVoid removes point off the model stage
func (point *Point) UnstageVoid(stage *Stage) {
	point.Unstage(stage)
}

func (point *Point) StageVoid(stage *Stage) {
	point.Stage(stage)
}

// for satisfaction of GongStruct interface
func (point *Point) GetName() (res string) {
	return point.Name
}

// for satisfaction of GongStruct interface
func (point *Point) SetName(name string) {
	point.Name = name
}

// Stage puts polygone to the model stage
func (polygone *Polygone) Stage(stage *Stage) *Polygone {
	__gong__stage(stage.Polygones, stage.Polygone_stagedOrder, stage.Polygone_orderStaged, &stage.PolygoneOrder, stage.Polygones_mapString, polygone, polygone.Name)
	return polygone
}

// StagePreserveOrder puts polygone to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PolygoneOrder
// - update stage.PolygoneOrder accordingly
func (polygone *Polygone) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Polygones, stage.Polygone_stagedOrder, stage.Polygone_orderStaged, &stage.PolygoneOrder, stage.Polygones_mapString, polygone, order, polygone.Name)
}

// Unstage removes polygone off the model stage
func (polygone *Polygone) Unstage(stage *Stage) *Polygone {
	__gong__unstage(stage.Polygones, stage.Polygones_mapString, polygone, polygone.Name)
	return polygone
}

// UnstageVoid removes polygone off the model stage
func (polygone *Polygone) UnstageVoid(stage *Stage) {
	polygone.Unstage(stage)
}

func (polygone *Polygone) StageVoid(stage *Stage) {
	polygone.Stage(stage)
}

// for satisfaction of GongStruct interface
func (polygone *Polygone) GetName() (res string) {
	return polygone.Name
}

// for satisfaction of GongStruct interface
func (polygone *Polygone) SetName(name string) {
	polygone.Name = name
}

// Stage puts polyline to the model stage
func (polyline *Polyline) Stage(stage *Stage) *Polyline {
	__gong__stage(stage.Polylines, stage.Polyline_stagedOrder, stage.Polyline_orderStaged, &stage.PolylineOrder, stage.Polylines_mapString, polyline, polyline.Name)
	return polyline
}

// StagePreserveOrder puts polyline to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PolylineOrder
// - update stage.PolylineOrder accordingly
func (polyline *Polyline) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Polylines, stage.Polyline_stagedOrder, stage.Polyline_orderStaged, &stage.PolylineOrder, stage.Polylines_mapString, polyline, order, polyline.Name)
}

// Unstage removes polyline off the model stage
func (polyline *Polyline) Unstage(stage *Stage) *Polyline {
	__gong__unstage(stage.Polylines, stage.Polylines_mapString, polyline, polyline.Name)
	return polyline
}

// UnstageVoid removes polyline off the model stage
func (polyline *Polyline) UnstageVoid(stage *Stage) {
	polyline.Unstage(stage)
}

func (polyline *Polyline) StageVoid(stage *Stage) {
	polyline.Stage(stage)
}

// for satisfaction of GongStruct interface
func (polyline *Polyline) GetName() (res string) {
	return polyline.Name
}

// for satisfaction of GongStruct interface
func (polyline *Polyline) SetName(name string) {
	polyline.Name = name
}

// Stage puts rect to the model stage
func (rect *Rect) Stage(stage *Stage) *Rect {
	__gong__stage(stage.Rects, stage.Rect_stagedOrder, stage.Rect_orderStaged, &stage.RectOrder, stage.Rects_mapString, rect, rect.Name)
	return rect
}

// StagePreserveOrder puts rect to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RectOrder
// - update stage.RectOrder accordingly
func (rect *Rect) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Rects, stage.Rect_stagedOrder, stage.Rect_orderStaged, &stage.RectOrder, stage.Rects_mapString, rect, order, rect.Name)
}

// Unstage removes rect off the model stage
func (rect *Rect) Unstage(stage *Stage) *Rect {
	__gong__unstage(stage.Rects, stage.Rects_mapString, rect, rect.Name)
	return rect
}

// UnstageVoid removes rect off the model stage
func (rect *Rect) UnstageVoid(stage *Stage) {
	rect.Unstage(stage)
}

func (rect *Rect) StageVoid(stage *Stage) {
	rect.Stage(stage)
}

// for satisfaction of GongStruct interface
func (rect *Rect) GetName() (res string) {
	return rect.Name
}

// for satisfaction of GongStruct interface
func (rect *Rect) SetName(name string) {
	rect.Name = name
}

// Stage puts rectanchoredpath to the model stage
func (rectanchoredpath *RectAnchoredPath) Stage(stage *Stage) *RectAnchoredPath {
	__gong__stage(stage.RectAnchoredPaths, stage.RectAnchoredPath_stagedOrder, stage.RectAnchoredPath_orderStaged, &stage.RectAnchoredPathOrder, stage.RectAnchoredPaths_mapString, rectanchoredpath, rectanchoredpath.Name)
	return rectanchoredpath
}

// StagePreserveOrder puts rectanchoredpath to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RectAnchoredPathOrder
// - update stage.RectAnchoredPathOrder accordingly
func (rectanchoredpath *RectAnchoredPath) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.RectAnchoredPaths, stage.RectAnchoredPath_stagedOrder, stage.RectAnchoredPath_orderStaged, &stage.RectAnchoredPathOrder, stage.RectAnchoredPaths_mapString, rectanchoredpath, order, rectanchoredpath.Name)
}

// Unstage removes rectanchoredpath off the model stage
func (rectanchoredpath *RectAnchoredPath) Unstage(stage *Stage) *RectAnchoredPath {
	__gong__unstage(stage.RectAnchoredPaths, stage.RectAnchoredPaths_mapString, rectanchoredpath, rectanchoredpath.Name)
	return rectanchoredpath
}

// UnstageVoid removes rectanchoredpath off the model stage
func (rectanchoredpath *RectAnchoredPath) UnstageVoid(stage *Stage) {
	rectanchoredpath.Unstage(stage)
}

func (rectanchoredpath *RectAnchoredPath) StageVoid(stage *Stage) {
	rectanchoredpath.Stage(stage)
}

// for satisfaction of GongStruct interface
func (rectanchoredpath *RectAnchoredPath) GetName() (res string) {
	return rectanchoredpath.Name
}

// for satisfaction of GongStruct interface
func (rectanchoredpath *RectAnchoredPath) SetName(name string) {
	rectanchoredpath.Name = name
}

// Stage puts rectanchoredpngimage to the model stage
func (rectanchoredpngimage *RectAnchoredPngImage) Stage(stage *Stage) *RectAnchoredPngImage {
	__gong__stage(stage.RectAnchoredPngImages, stage.RectAnchoredPngImage_stagedOrder, stage.RectAnchoredPngImage_orderStaged, &stage.RectAnchoredPngImageOrder, stage.RectAnchoredPngImages_mapString, rectanchoredpngimage, rectanchoredpngimage.Name)
	return rectanchoredpngimage
}

// StagePreserveOrder puts rectanchoredpngimage to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RectAnchoredPngImageOrder
// - update stage.RectAnchoredPngImageOrder accordingly
func (rectanchoredpngimage *RectAnchoredPngImage) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.RectAnchoredPngImages, stage.RectAnchoredPngImage_stagedOrder, stage.RectAnchoredPngImage_orderStaged, &stage.RectAnchoredPngImageOrder, stage.RectAnchoredPngImages_mapString, rectanchoredpngimage, order, rectanchoredpngimage.Name)
}

// Unstage removes rectanchoredpngimage off the model stage
func (rectanchoredpngimage *RectAnchoredPngImage) Unstage(stage *Stage) *RectAnchoredPngImage {
	__gong__unstage(stage.RectAnchoredPngImages, stage.RectAnchoredPngImages_mapString, rectanchoredpngimage, rectanchoredpngimage.Name)
	return rectanchoredpngimage
}

// UnstageVoid removes rectanchoredpngimage off the model stage
func (rectanchoredpngimage *RectAnchoredPngImage) UnstageVoid(stage *Stage) {
	rectanchoredpngimage.Unstage(stage)
}

func (rectanchoredpngimage *RectAnchoredPngImage) StageVoid(stage *Stage) {
	rectanchoredpngimage.Stage(stage)
}

// for satisfaction of GongStruct interface
func (rectanchoredpngimage *RectAnchoredPngImage) GetName() (res string) {
	return rectanchoredpngimage.Name
}

// for satisfaction of GongStruct interface
func (rectanchoredpngimage *RectAnchoredPngImage) SetName(name string) {
	rectanchoredpngimage.Name = name
}

// Stage puts rectanchoredrect to the model stage
func (rectanchoredrect *RectAnchoredRect) Stage(stage *Stage) *RectAnchoredRect {
	__gong__stage(stage.RectAnchoredRects, stage.RectAnchoredRect_stagedOrder, stage.RectAnchoredRect_orderStaged, &stage.RectAnchoredRectOrder, stage.RectAnchoredRects_mapString, rectanchoredrect, rectanchoredrect.Name)
	return rectanchoredrect
}

// StagePreserveOrder puts rectanchoredrect to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RectAnchoredRectOrder
// - update stage.RectAnchoredRectOrder accordingly
func (rectanchoredrect *RectAnchoredRect) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.RectAnchoredRects, stage.RectAnchoredRect_stagedOrder, stage.RectAnchoredRect_orderStaged, &stage.RectAnchoredRectOrder, stage.RectAnchoredRects_mapString, rectanchoredrect, order, rectanchoredrect.Name)
}

// Unstage removes rectanchoredrect off the model stage
func (rectanchoredrect *RectAnchoredRect) Unstage(stage *Stage) *RectAnchoredRect {
	__gong__unstage(stage.RectAnchoredRects, stage.RectAnchoredRects_mapString, rectanchoredrect, rectanchoredrect.Name)
	return rectanchoredrect
}

// UnstageVoid removes rectanchoredrect off the model stage
func (rectanchoredrect *RectAnchoredRect) UnstageVoid(stage *Stage) {
	rectanchoredrect.Unstage(stage)
}

func (rectanchoredrect *RectAnchoredRect) StageVoid(stage *Stage) {
	rectanchoredrect.Stage(stage)
}

// for satisfaction of GongStruct interface
func (rectanchoredrect *RectAnchoredRect) GetName() (res string) {
	return rectanchoredrect.Name
}

// for satisfaction of GongStruct interface
func (rectanchoredrect *RectAnchoredRect) SetName(name string) {
	rectanchoredrect.Name = name
}

// Stage puts rectanchoredtext to the model stage
func (rectanchoredtext *RectAnchoredText) Stage(stage *Stage) *RectAnchoredText {
	__gong__stage(stage.RectAnchoredTexts, stage.RectAnchoredText_stagedOrder, stage.RectAnchoredText_orderStaged, &stage.RectAnchoredTextOrder, stage.RectAnchoredTexts_mapString, rectanchoredtext, rectanchoredtext.Name)
	return rectanchoredtext
}

// StagePreserveOrder puts rectanchoredtext to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RectAnchoredTextOrder
// - update stage.RectAnchoredTextOrder accordingly
func (rectanchoredtext *RectAnchoredText) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.RectAnchoredTexts, stage.RectAnchoredText_stagedOrder, stage.RectAnchoredText_orderStaged, &stage.RectAnchoredTextOrder, stage.RectAnchoredTexts_mapString, rectanchoredtext, order, rectanchoredtext.Name)
}

// Unstage removes rectanchoredtext off the model stage
func (rectanchoredtext *RectAnchoredText) Unstage(stage *Stage) *RectAnchoredText {
	__gong__unstage(stage.RectAnchoredTexts, stage.RectAnchoredTexts_mapString, rectanchoredtext, rectanchoredtext.Name)
	return rectanchoredtext
}

// UnstageVoid removes rectanchoredtext off the model stage
func (rectanchoredtext *RectAnchoredText) UnstageVoid(stage *Stage) {
	rectanchoredtext.Unstage(stage)
}

func (rectanchoredtext *RectAnchoredText) StageVoid(stage *Stage) {
	rectanchoredtext.Stage(stage)
}

// for satisfaction of GongStruct interface
func (rectanchoredtext *RectAnchoredText) GetName() (res string) {
	return rectanchoredtext.Name
}

// for satisfaction of GongStruct interface
func (rectanchoredtext *RectAnchoredText) SetName(name string) {
	rectanchoredtext.Name = name
}

// Stage puts rectlinklink to the model stage
func (rectlinklink *RectLinkLink) Stage(stage *Stage) *RectLinkLink {
	__gong__stage(stage.RectLinkLinks, stage.RectLinkLink_stagedOrder, stage.RectLinkLink_orderStaged, &stage.RectLinkLinkOrder, stage.RectLinkLinks_mapString, rectlinklink, rectlinklink.Name)
	return rectlinklink
}

// StagePreserveOrder puts rectlinklink to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RectLinkLinkOrder
// - update stage.RectLinkLinkOrder accordingly
func (rectlinklink *RectLinkLink) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.RectLinkLinks, stage.RectLinkLink_stagedOrder, stage.RectLinkLink_orderStaged, &stage.RectLinkLinkOrder, stage.RectLinkLinks_mapString, rectlinklink, order, rectlinklink.Name)
}

// Unstage removes rectlinklink off the model stage
func (rectlinklink *RectLinkLink) Unstage(stage *Stage) *RectLinkLink {
	__gong__unstage(stage.RectLinkLinks, stage.RectLinkLinks_mapString, rectlinklink, rectlinklink.Name)
	return rectlinklink
}

// UnstageVoid removes rectlinklink off the model stage
func (rectlinklink *RectLinkLink) UnstageVoid(stage *Stage) {
	rectlinklink.Unstage(stage)
}

func (rectlinklink *RectLinkLink) StageVoid(stage *Stage) {
	rectlinklink.Stage(stage)
}

// for satisfaction of GongStruct interface
func (rectlinklink *RectLinkLink) GetName() (res string) {
	return rectlinklink.Name
}

// for satisfaction of GongStruct interface
func (rectlinklink *RectLinkLink) SetName(name string) {
	rectlinklink.Name = name
}

// Stage puts svg to the model stage
func (svg *SVG) Stage(stage *Stage) *SVG {
	__gong__stage(stage.SVGs, stage.SVG_stagedOrder, stage.SVG_orderStaged, &stage.SVGOrder, stage.SVGs_mapString, svg, svg.Name)
	return svg
}

// StagePreserveOrder puts svg to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SVGOrder
// - update stage.SVGOrder accordingly
func (svg *SVG) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.SVGs, stage.SVG_stagedOrder, stage.SVG_orderStaged, &stage.SVGOrder, stage.SVGs_mapString, svg, order, svg.Name)
}

// Unstage removes svg off the model stage
func (svg *SVG) Unstage(stage *Stage) *SVG {
	__gong__unstage(stage.SVGs, stage.SVGs_mapString, svg, svg.Name)
	return svg
}

// UnstageVoid removes svg off the model stage
func (svg *SVG) UnstageVoid(stage *Stage) {
	svg.Unstage(stage)
}

func (svg *SVG) StageVoid(stage *Stage) {
	svg.Stage(stage)
}

// for satisfaction of GongStruct interface
func (svg *SVG) GetName() (res string) {
	return svg.Name
}

// for satisfaction of GongStruct interface
func (svg *SVG) SetName(name string) {
	svg.Name = name
}

// Stage puts svgtext to the model stage
func (svgtext *SvgText) Stage(stage *Stage) *SvgText {
	__gong__stage(stage.SvgTexts, stage.SvgText_stagedOrder, stage.SvgText_orderStaged, &stage.SvgTextOrder, stage.SvgTexts_mapString, svgtext, svgtext.Name)
	return svgtext
}

// StagePreserveOrder puts svgtext to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SvgTextOrder
// - update stage.SvgTextOrder accordingly
func (svgtext *SvgText) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.SvgTexts, stage.SvgText_stagedOrder, stage.SvgText_orderStaged, &stage.SvgTextOrder, stage.SvgTexts_mapString, svgtext, order, svgtext.Name)
}

// Unstage removes svgtext off the model stage
func (svgtext *SvgText) Unstage(stage *Stage) *SvgText {
	__gong__unstage(stage.SvgTexts, stage.SvgTexts_mapString, svgtext, svgtext.Name)
	return svgtext
}

// UnstageVoid removes svgtext off the model stage
func (svgtext *SvgText) UnstageVoid(stage *Stage) {
	svgtext.Unstage(stage)
}

func (svgtext *SvgText) StageVoid(stage *Stage) {
	svgtext.Stage(stage)
}

// for satisfaction of GongStruct interface
func (svgtext *SvgText) GetName() (res string) {
	return svgtext.Name
}

// for satisfaction of GongStruct interface
func (svgtext *SvgText) SetName(name string) {
	svgtext.Name = name
}

// Stage puts text to the model stage
func (text *Text) Stage(stage *Stage) *Text {
	__gong__stage(stage.Texts, stage.Text_stagedOrder, stage.Text_orderStaged, &stage.TextOrder, stage.Texts_mapString, text, text.Name)
	return text
}

// StagePreserveOrder puts text to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TextOrder
// - update stage.TextOrder accordingly
func (text *Text) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Texts, stage.Text_stagedOrder, stage.Text_orderStaged, &stage.TextOrder, stage.Texts_mapString, text, order, text.Name)
}

// Unstage removes text off the model stage
func (text *Text) Unstage(stage *Stage) *Text {
	__gong__unstage(stage.Texts, stage.Texts_mapString, text, text.Name)
	return text
}

// UnstageVoid removes text off the model stage
func (text *Text) UnstageVoid(stage *Stage) {
	text.Unstage(stage)
}

func (text *Text) StageVoid(stage *Stage) {
	text.Stage(stage)
}

// for satisfaction of GongStruct interface
func (text *Text) GetName() (res string) {
	return text.Name
}

// for satisfaction of GongStruct interface
func (text *Text) SetName(name string) {
	text.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.Animates, &stage.Animates_mapString, &stage.Animate_stagedOrder, &stage.AnimateOrder)

	__gong__resetStageType(&stage.Circles, &stage.Circles_mapString, &stage.Circle_stagedOrder, &stage.CircleOrder)

	__gong__resetStageType(&stage.Conditions, &stage.Conditions_mapString, &stage.Condition_stagedOrder, &stage.ConditionOrder)

	__gong__resetStageType(&stage.ControlPoints, &stage.ControlPoints_mapString, &stage.ControlPoint_stagedOrder, &stage.ControlPointOrder)

	__gong__resetStageType(&stage.Ellipses, &stage.Ellipses_mapString, &stage.Ellipse_stagedOrder, &stage.EllipseOrder)

	__gong__resetStageType(&stage.FileToDownloads, &stage.FileToDownloads_mapString, &stage.FileToDownload_stagedOrder, &stage.FileToDownloadOrder)

	__gong__resetStageType(&stage.Layers, &stage.Layers_mapString, &stage.Layer_stagedOrder, &stage.LayerOrder)

	__gong__resetStageType(&stage.Lines, &stage.Lines_mapString, &stage.Line_stagedOrder, &stage.LineOrder)

	__gong__resetStageType(&stage.Links, &stage.Links_mapString, &stage.Link_stagedOrder, &stage.LinkOrder)

	__gong__resetStageType(&stage.LinkAnchoredPaths, &stage.LinkAnchoredPaths_mapString, &stage.LinkAnchoredPath_stagedOrder, &stage.LinkAnchoredPathOrder)

	__gong__resetStageType(&stage.LinkAnchoredTexts, &stage.LinkAnchoredTexts_mapString, &stage.LinkAnchoredText_stagedOrder, &stage.LinkAnchoredTextOrder)

	__gong__resetStageType(&stage.Paths, &stage.Paths_mapString, &stage.Path_stagedOrder, &stage.PathOrder)

	__gong__resetStageType(&stage.Points, &stage.Points_mapString, &stage.Point_stagedOrder, &stage.PointOrder)

	__gong__resetStageType(&stage.Polygones, &stage.Polygones_mapString, &stage.Polygone_stagedOrder, &stage.PolygoneOrder)

	__gong__resetStageType(&stage.Polylines, &stage.Polylines_mapString, &stage.Polyline_stagedOrder, &stage.PolylineOrder)

	__gong__resetStageType(&stage.Rects, &stage.Rects_mapString, &stage.Rect_stagedOrder, &stage.RectOrder)

	__gong__resetStageType(&stage.RectAnchoredPaths, &stage.RectAnchoredPaths_mapString, &stage.RectAnchoredPath_stagedOrder, &stage.RectAnchoredPathOrder)

	__gong__resetStageType(&stage.RectAnchoredPngImages, &stage.RectAnchoredPngImages_mapString, &stage.RectAnchoredPngImage_stagedOrder, &stage.RectAnchoredPngImageOrder)

	__gong__resetStageType(&stage.RectAnchoredRects, &stage.RectAnchoredRects_mapString, &stage.RectAnchoredRect_stagedOrder, &stage.RectAnchoredRectOrder)

	__gong__resetStageType(&stage.RectAnchoredTexts, &stage.RectAnchoredTexts_mapString, &stage.RectAnchoredText_stagedOrder, &stage.RectAnchoredTextOrder)

	__gong__resetStageType(&stage.RectLinkLinks, &stage.RectLinkLinks_mapString, &stage.RectLinkLink_stagedOrder, &stage.RectLinkLinkOrder)

	__gong__resetStageType(&stage.SVGs, &stage.SVGs_mapString, &stage.SVG_stagedOrder, &stage.SVGOrder)

	__gong__resetStageType(&stage.SvgTexts, &stage.SvgTexts_mapString, &stage.SvgText_stagedOrder, &stage.SvgTextOrder)

	__gong__resetStageType(&stage.Texts, &stage.Texts_mapString, &stage.Text_stagedOrder, &stage.TextOrder)

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
	case *Animate:
		return any(stage.Animates_mapString).(map[string]Type)
	case *Circle:
		return any(stage.Circles_mapString).(map[string]Type)
	case *Condition:
		return any(stage.Conditions_mapString).(map[string]Type)
	case *ControlPoint:
		return any(stage.ControlPoints_mapString).(map[string]Type)
	case *Ellipse:
		return any(stage.Ellipses_mapString).(map[string]Type)
	case *FileToDownload:
		return any(stage.FileToDownloads_mapString).(map[string]Type)
	case *Layer:
		return any(stage.Layers_mapString).(map[string]Type)
	case *Line:
		return any(stage.Lines_mapString).(map[string]Type)
	case *Link:
		return any(stage.Links_mapString).(map[string]Type)
	case *LinkAnchoredPath:
		return any(stage.LinkAnchoredPaths_mapString).(map[string]Type)
	case *LinkAnchoredText:
		return any(stage.LinkAnchoredTexts_mapString).(map[string]Type)
	case *Path:
		return any(stage.Paths_mapString).(map[string]Type)
	case *Point:
		return any(stage.Points_mapString).(map[string]Type)
	case *Polygone:
		return any(stage.Polygones_mapString).(map[string]Type)
	case *Polyline:
		return any(stage.Polylines_mapString).(map[string]Type)
	case *Rect:
		return any(stage.Rects_mapString).(map[string]Type)
	case *RectAnchoredPath:
		return any(stage.RectAnchoredPaths_mapString).(map[string]Type)
	case *RectAnchoredPngImage:
		return any(stage.RectAnchoredPngImages_mapString).(map[string]Type)
	case *RectAnchoredRect:
		return any(stage.RectAnchoredRects_mapString).(map[string]Type)
	case *RectAnchoredText:
		return any(stage.RectAnchoredTexts_mapString).(map[string]Type)
	case *RectLinkLink:
		return any(stage.RectLinkLinks_mapString).(map[string]Type)
	case *SVG:
		return any(stage.SVGs_mapString).(map[string]Type)
	case *SvgText:
		return any(stage.SvgTexts_mapString).(map[string]Type)
	case *Text:
		return any(stage.Texts_mapString).(map[string]Type)
	default:
		return nil
	}
}

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Animate:
		return any(&stage.Animates).(*map[Type]struct{})
	case *Circle:
		return any(&stage.Circles).(*map[Type]struct{})
	case *Condition:
		return any(&stage.Conditions).(*map[Type]struct{})
	case *ControlPoint:
		return any(&stage.ControlPoints).(*map[Type]struct{})
	case *Ellipse:
		return any(&stage.Ellipses).(*map[Type]struct{})
	case *FileToDownload:
		return any(&stage.FileToDownloads).(*map[Type]struct{})
	case *Layer:
		return any(&stage.Layers).(*map[Type]struct{})
	case *Line:
		return any(&stage.Lines).(*map[Type]struct{})
	case *Link:
		return any(&stage.Links).(*map[Type]struct{})
	case *LinkAnchoredPath:
		return any(&stage.LinkAnchoredPaths).(*map[Type]struct{})
	case *LinkAnchoredText:
		return any(&stage.LinkAnchoredTexts).(*map[Type]struct{})
	case *Path:
		return any(&stage.Paths).(*map[Type]struct{})
	case *Point:
		return any(&stage.Points).(*map[Type]struct{})
	case *Polygone:
		return any(&stage.Polygones).(*map[Type]struct{})
	case *Polyline:
		return any(&stage.Polylines).(*map[Type]struct{})
	case *Rect:
		return any(&stage.Rects).(*map[Type]struct{})
	case *RectAnchoredPath:
		return any(&stage.RectAnchoredPaths).(*map[Type]struct{})
	case *RectAnchoredPngImage:
		return any(&stage.RectAnchoredPngImages).(*map[Type]struct{})
	case *RectAnchoredRect:
		return any(&stage.RectAnchoredRects).(*map[Type]struct{})
	case *RectAnchoredText:
		return any(&stage.RectAnchoredTexts).(*map[Type]struct{})
	case *RectLinkLink:
		return any(&stage.RectLinkLinks).(*map[Type]struct{})
	case *SVG:
		return any(&stage.SVGs).(*map[Type]struct{})
	case *SvgText:
		return any(&stage.SvgTexts).(*map[Type]struct{})
	case *Text:
		return any(&stage.Texts).(*map[Type]struct{})
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
	case Circle:
		return any(&Circle{
			Animations: []*Animate{{Name: "Animations"}},
		}).(*Type)
	case ControlPoint:
		return any(&ControlPoint{
			ClosestRect: &Rect{Name: "ClosestRect"},
		}).(*Type)
	case Ellipse:
		return any(&Ellipse{
			Animates: []*Animate{{Name: "Animates"}},
		}).(*Type)
	case Layer:
		return any(&Layer{
			Rects: []*Rect{{Name: "Rects"}},
			Texts: []*Text{{Name: "Texts"}},
			Circles: []*Circle{{Name: "Circles"}},
			Lines: []*Line{{Name: "Lines"}},
			Ellipses: []*Ellipse{{Name: "Ellipses"}},
			Polylines: []*Polyline{{Name: "Polylines"}},
			Polygones: []*Polygone{{Name: "Polygones"}},
			Paths: []*Path{{Name: "Paths"}},
			Links: []*Link{{Name: "Links"}},
			RectLinkLinks: []*RectLinkLink{{Name: "RectLinkLinks"}},
		}).(*Type)
	case Line:
		return any(&Line{
			Animates: []*Animate{{Name: "Animates"}},
		}).(*Type)
	case Link:
		return any(&Link{
			Start: &Rect{Name: "Start"},
			End: &Rect{Name: "End"},
			TextAtArrowStart: []*LinkAnchoredText{{Name: "TextAtArrowStart"}},
			TextAtArrowEnd: []*LinkAnchoredText{{Name: "TextAtArrowEnd"}},
			TextAtCorner: []*LinkAnchoredText{{Name: "TextAtCorner"}},
			PathAtArrowStart: []*LinkAnchoredPath{{Name: "PathAtArrowStart"}},
			PathAtArrowEnd: []*LinkAnchoredPath{{Name: "PathAtArrowEnd"}},
			PathAtCorner: []*LinkAnchoredPath{{Name: "PathAtCorner"}},
			ControlPoints: []*ControlPoint{{Name: "ControlPoints"}},
		}).(*Type)
	case LinkAnchoredText:
		return any(&LinkAnchoredText{
			Animates: []*Animate{{Name: "Animates"}},
		}).(*Type)
	case Path:
		return any(&Path{
			Animates: []*Animate{{Name: "Animates"}},
		}).(*Type)
	case Polygone:
		return any(&Polygone{
			Animates: []*Animate{{Name: "Animates"}},
		}).(*Type)
	case Polyline:
		return any(&Polyline{
			Animates: []*Animate{{Name: "Animates"}},
		}).(*Type)
	case Rect:
		return any(&Rect{
			Peers: []*Rect{{Name: "Peers"}},
			EnclosingRect: &Rect{Name: "EnclosingRect"},
			Obstacles: []*Rect{{Name: "Obstacles"}},
			AnchoredTo: &Rect{Name: "AnchoredTo"},
			HoveringTrigger: []*Condition{{Name: "HoveringTrigger"}},
			DisplayConditions: []*Condition{{Name: "DisplayConditions"}},
			Animations: []*Animate{{Name: "Animations"}},
			RectAnchoredTexts: []*RectAnchoredText{{Name: "RectAnchoredTexts"}},
			RectAnchoredRects: []*RectAnchoredRect{{Name: "RectAnchoredRects"}},
			RectAnchoredPaths: []*RectAnchoredPath{{Name: "RectAnchoredPaths"}},
			RectAnchoredPngImages: []*RectAnchoredPngImage{{Name: "RectAnchoredPngImages"}},
		}).(*Type)
	case RectAnchoredText:
		return any(&RectAnchoredText{
			Animates: []*Animate{{Name: "Animates"}},
		}).(*Type)
	case RectLinkLink:
		return any(&RectLinkLink{
			Start: &Rect{Name: "Start"},
			End: &Link{Name: "End"},
		}).(*Type)
	case SVG:
		return any(&SVG{
			Layers: []*Layer{{Name: "Layers"}},
			StartRect: &Rect{Name: "StartRect"},
			EndRect: &Rect{Name: "EndRect"},
		}).(*Type)
	case Text:
		return any(&Text{
			Animates: []*Animate{{Name: "Animates"}},
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
	// reverse maps of direct associations of Animate
	case Animate:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Circle
	case Circle:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Condition
	case Condition:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ControlPoint
	case ControlPoint:
		switch fieldname {
		// insertion point for per direct association field
		case "ClosestRect":
			res := make(map[*Rect][]*ControlPoint)
			for controlpoint := range stage.ControlPoints {
				if controlpoint.ClosestRect != nil {
					rect_ := controlpoint.ClosestRect
					var controlpoints []*ControlPoint
					_, ok := res[rect_]
					if ok {
						controlpoints = res[rect_]
					} else {
						controlpoints = make([]*ControlPoint, 0)
					}
					controlpoints = append(controlpoints, controlpoint)
					res[rect_] = controlpoints
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Ellipse
	case Ellipse:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FileToDownload
	case FileToDownload:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Layer
	case Layer:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Line
	case Line:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Link
	case Link:
		switch fieldname {
		// insertion point for per direct association field
		case "Start":
			res := make(map[*Rect][]*Link)
			for link := range stage.Links {
				if link.Start != nil {
					rect_ := link.Start
					var links []*Link
					_, ok := res[rect_]
					if ok {
						links = res[rect_]
					} else {
						links = make([]*Link, 0)
					}
					links = append(links, link)
					res[rect_] = links
				}
			}
			return any(res).(map[*End][]*Start)
		case "End":
			res := make(map[*Rect][]*Link)
			for link := range stage.Links {
				if link.End != nil {
					rect_ := link.End
					var links []*Link
					_, ok := res[rect_]
					if ok {
						links = res[rect_]
					} else {
						links = make([]*Link, 0)
					}
					links = append(links, link)
					res[rect_] = links
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of LinkAnchoredPath
	case LinkAnchoredPath:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LinkAnchoredText
	case LinkAnchoredText:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Path
	case Path:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Point
	case Point:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Polygone
	case Polygone:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Polyline
	case Polyline:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Rect
	case Rect:
		switch fieldname {
		// insertion point for per direct association field
		case "EnclosingRect":
			res := make(map[*Rect][]*Rect)
			for rect := range stage.Rects {
				if rect.EnclosingRect != nil {
					rect_ := rect.EnclosingRect
					var rects []*Rect
					_, ok := res[rect_]
					if ok {
						rects = res[rect_]
					} else {
						rects = make([]*Rect, 0)
					}
					rects = append(rects, rect)
					res[rect_] = rects
				}
			}
			return any(res).(map[*End][]*Start)
		case "AnchoredTo":
			res := make(map[*Rect][]*Rect)
			for rect := range stage.Rects {
				if rect.AnchoredTo != nil {
					rect_ := rect.AnchoredTo
					var rects []*Rect
					_, ok := res[rect_]
					if ok {
						rects = res[rect_]
					} else {
						rects = make([]*Rect, 0)
					}
					rects = append(rects, rect)
					res[rect_] = rects
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of RectAnchoredPath
	case RectAnchoredPath:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RectAnchoredPngImage
	case RectAnchoredPngImage:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RectAnchoredRect
	case RectAnchoredRect:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RectAnchoredText
	case RectAnchoredText:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RectLinkLink
	case RectLinkLink:
		switch fieldname {
		// insertion point for per direct association field
		case "Start":
			res := make(map[*Rect][]*RectLinkLink)
			for rectlinklink := range stage.RectLinkLinks {
				if rectlinklink.Start != nil {
					rect_ := rectlinklink.Start
					var rectlinklinks []*RectLinkLink
					_, ok := res[rect_]
					if ok {
						rectlinklinks = res[rect_]
					} else {
						rectlinklinks = make([]*RectLinkLink, 0)
					}
					rectlinklinks = append(rectlinklinks, rectlinklink)
					res[rect_] = rectlinklinks
				}
			}
			return any(res).(map[*End][]*Start)
		case "End":
			res := make(map[*Link][]*RectLinkLink)
			for rectlinklink := range stage.RectLinkLinks {
				if rectlinklink.End != nil {
					link_ := rectlinklink.End
					var rectlinklinks []*RectLinkLink
					_, ok := res[link_]
					if ok {
						rectlinklinks = res[link_]
					} else {
						rectlinklinks = make([]*RectLinkLink, 0)
					}
					rectlinklinks = append(rectlinklinks, rectlinklink)
					res[link_] = rectlinklinks
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SVG
	case SVG:
		switch fieldname {
		// insertion point for per direct association field
		case "StartRect":
			res := make(map[*Rect][]*SVG)
			for svg := range stage.SVGs {
				if svg.StartRect != nil {
					rect_ := svg.StartRect
					var svgs []*SVG
					_, ok := res[rect_]
					if ok {
						svgs = res[rect_]
					} else {
						svgs = make([]*SVG, 0)
					}
					svgs = append(svgs, svg)
					res[rect_] = svgs
				}
			}
			return any(res).(map[*End][]*Start)
		case "EndRect":
			res := make(map[*Rect][]*SVG)
			for svg := range stage.SVGs {
				if svg.EndRect != nil {
					rect_ := svg.EndRect
					var svgs []*SVG
					_, ok := res[rect_]
					if ok {
						svgs = res[rect_]
					} else {
						svgs = make([]*SVG, 0)
					}
					svgs = append(svgs, svg)
					res[rect_] = svgs
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SvgText
	case SvgText:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Text
	case Text:
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
	// reverse maps of direct associations of Animate
	case Animate:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Circle
	case Circle:
		switch fieldname {
		// insertion point for per direct association field
		case "Animations":
			res := make(map[*Animate][]*Circle)
			for circle := range stage.Circles {
				for _, animate_ := range circle.Animations {
					res[animate_] = append(res[animate_], circle)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Condition
	case Condition:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ControlPoint
	case ControlPoint:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Ellipse
	case Ellipse:
		switch fieldname {
		// insertion point for per direct association field
		case "Animates":
			res := make(map[*Animate][]*Ellipse)
			for ellipse := range stage.Ellipses {
				for _, animate_ := range ellipse.Animates {
					res[animate_] = append(res[animate_], ellipse)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of FileToDownload
	case FileToDownload:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Layer
	case Layer:
		switch fieldname {
		// insertion point for per direct association field
		case "Rects":
			res := make(map[*Rect][]*Layer)
			for layer := range stage.Layers {
				for _, rect_ := range layer.Rects {
					res[rect_] = append(res[rect_], layer)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Texts":
			res := make(map[*Text][]*Layer)
			for layer := range stage.Layers {
				for _, text_ := range layer.Texts {
					res[text_] = append(res[text_], layer)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Circles":
			res := make(map[*Circle][]*Layer)
			for layer := range stage.Layers {
				for _, circle_ := range layer.Circles {
					res[circle_] = append(res[circle_], layer)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Lines":
			res := make(map[*Line][]*Layer)
			for layer := range stage.Layers {
				for _, line_ := range layer.Lines {
					res[line_] = append(res[line_], layer)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Ellipses":
			res := make(map[*Ellipse][]*Layer)
			for layer := range stage.Layers {
				for _, ellipse_ := range layer.Ellipses {
					res[ellipse_] = append(res[ellipse_], layer)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Polylines":
			res := make(map[*Polyline][]*Layer)
			for layer := range stage.Layers {
				for _, polyline_ := range layer.Polylines {
					res[polyline_] = append(res[polyline_], layer)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Polygones":
			res := make(map[*Polygone][]*Layer)
			for layer := range stage.Layers {
				for _, polygone_ := range layer.Polygones {
					res[polygone_] = append(res[polygone_], layer)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Paths":
			res := make(map[*Path][]*Layer)
			for layer := range stage.Layers {
				for _, path_ := range layer.Paths {
					res[path_] = append(res[path_], layer)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Links":
			res := make(map[*Link][]*Layer)
			for layer := range stage.Layers {
				for _, link_ := range layer.Links {
					res[link_] = append(res[link_], layer)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RectLinkLinks":
			res := make(map[*RectLinkLink][]*Layer)
			for layer := range stage.Layers {
				for _, rectlinklink_ := range layer.RectLinkLinks {
					res[rectlinklink_] = append(res[rectlinklink_], layer)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Line
	case Line:
		switch fieldname {
		// insertion point for per direct association field
		case "Animates":
			res := make(map[*Animate][]*Line)
			for line := range stage.Lines {
				for _, animate_ := range line.Animates {
					res[animate_] = append(res[animate_], line)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Link
	case Link:
		switch fieldname {
		// insertion point for per direct association field
		case "TextAtArrowStart":
			res := make(map[*LinkAnchoredText][]*Link)
			for link := range stage.Links {
				for _, linkanchoredtext_ := range link.TextAtArrowStart {
					res[linkanchoredtext_] = append(res[linkanchoredtext_], link)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TextAtArrowEnd":
			res := make(map[*LinkAnchoredText][]*Link)
			for link := range stage.Links {
				for _, linkanchoredtext_ := range link.TextAtArrowEnd {
					res[linkanchoredtext_] = append(res[linkanchoredtext_], link)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TextAtCorner":
			res := make(map[*LinkAnchoredText][]*Link)
			for link := range stage.Links {
				for _, linkanchoredtext_ := range link.TextAtCorner {
					res[linkanchoredtext_] = append(res[linkanchoredtext_], link)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PathAtArrowStart":
			res := make(map[*LinkAnchoredPath][]*Link)
			for link := range stage.Links {
				for _, linkanchoredpath_ := range link.PathAtArrowStart {
					res[linkanchoredpath_] = append(res[linkanchoredpath_], link)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PathAtArrowEnd":
			res := make(map[*LinkAnchoredPath][]*Link)
			for link := range stage.Links {
				for _, linkanchoredpath_ := range link.PathAtArrowEnd {
					res[linkanchoredpath_] = append(res[linkanchoredpath_], link)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PathAtCorner":
			res := make(map[*LinkAnchoredPath][]*Link)
			for link := range stage.Links {
				for _, linkanchoredpath_ := range link.PathAtCorner {
					res[linkanchoredpath_] = append(res[linkanchoredpath_], link)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ControlPoints":
			res := make(map[*ControlPoint][]*Link)
			for link := range stage.Links {
				for _, controlpoint_ := range link.ControlPoints {
					res[controlpoint_] = append(res[controlpoint_], link)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of LinkAnchoredPath
	case LinkAnchoredPath:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LinkAnchoredText
	case LinkAnchoredText:
		switch fieldname {
		// insertion point for per direct association field
		case "Animates":
			res := make(map[*Animate][]*LinkAnchoredText)
			for linkanchoredtext := range stage.LinkAnchoredTexts {
				for _, animate_ := range linkanchoredtext.Animates {
					res[animate_] = append(res[animate_], linkanchoredtext)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Path
	case Path:
		switch fieldname {
		// insertion point for per direct association field
		case "Animates":
			res := make(map[*Animate][]*Path)
			for path := range stage.Paths {
				for _, animate_ := range path.Animates {
					res[animate_] = append(res[animate_], path)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Point
	case Point:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Polygone
	case Polygone:
		switch fieldname {
		// insertion point for per direct association field
		case "Animates":
			res := make(map[*Animate][]*Polygone)
			for polygone := range stage.Polygones {
				for _, animate_ := range polygone.Animates {
					res[animate_] = append(res[animate_], polygone)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Polyline
	case Polyline:
		switch fieldname {
		// insertion point for per direct association field
		case "Animates":
			res := make(map[*Animate][]*Polyline)
			for polyline := range stage.Polylines {
				for _, animate_ := range polyline.Animates {
					res[animate_] = append(res[animate_], polyline)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Rect
	case Rect:
		switch fieldname {
		// insertion point for per direct association field
		case "Peers":
			res := make(map[*Rect][]*Rect)
			for rect := range stage.Rects {
				for _, rect_ := range rect.Peers {
					res[rect_] = append(res[rect_], rect)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Obstacles":
			res := make(map[*Rect][]*Rect)
			for rect := range stage.Rects {
				for _, rect_ := range rect.Obstacles {
					res[rect_] = append(res[rect_], rect)
				}
			}
			return any(res).(map[*End][]*Start)
		case "HoveringTrigger":
			res := make(map[*Condition][]*Rect)
			for rect := range stage.Rects {
				for _, condition_ := range rect.HoveringTrigger {
					res[condition_] = append(res[condition_], rect)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DisplayConditions":
			res := make(map[*Condition][]*Rect)
			for rect := range stage.Rects {
				for _, condition_ := range rect.DisplayConditions {
					res[condition_] = append(res[condition_], rect)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Animations":
			res := make(map[*Animate][]*Rect)
			for rect := range stage.Rects {
				for _, animate_ := range rect.Animations {
					res[animate_] = append(res[animate_], rect)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RectAnchoredTexts":
			res := make(map[*RectAnchoredText][]*Rect)
			for rect := range stage.Rects {
				for _, rectanchoredtext_ := range rect.RectAnchoredTexts {
					res[rectanchoredtext_] = append(res[rectanchoredtext_], rect)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RectAnchoredRects":
			res := make(map[*RectAnchoredRect][]*Rect)
			for rect := range stage.Rects {
				for _, rectanchoredrect_ := range rect.RectAnchoredRects {
					res[rectanchoredrect_] = append(res[rectanchoredrect_], rect)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RectAnchoredPaths":
			res := make(map[*RectAnchoredPath][]*Rect)
			for rect := range stage.Rects {
				for _, rectanchoredpath_ := range rect.RectAnchoredPaths {
					res[rectanchoredpath_] = append(res[rectanchoredpath_], rect)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RectAnchoredPngImages":
			res := make(map[*RectAnchoredPngImage][]*Rect)
			for rect := range stage.Rects {
				for _, rectanchoredpngimage_ := range rect.RectAnchoredPngImages {
					res[rectanchoredpngimage_] = append(res[rectanchoredpngimage_], rect)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of RectAnchoredPath
	case RectAnchoredPath:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RectAnchoredPngImage
	case RectAnchoredPngImage:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RectAnchoredRect
	case RectAnchoredRect:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RectAnchoredText
	case RectAnchoredText:
		switch fieldname {
		// insertion point for per direct association field
		case "Animates":
			res := make(map[*Animate][]*RectAnchoredText)
			for rectanchoredtext := range stage.RectAnchoredTexts {
				for _, animate_ := range rectanchoredtext.Animates {
					res[animate_] = append(res[animate_], rectanchoredtext)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of RectLinkLink
	case RectLinkLink:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SVG
	case SVG:
		switch fieldname {
		// insertion point for per direct association field
		case "Layers":
			res := make(map[*Layer][]*SVG)
			for svg := range stage.SVGs {
				for _, layer_ := range svg.Layers {
					res[layer_] = append(res[layer_], svg)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SvgText
	case SvgText:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Text
	case Text:
		switch fieldname {
		// insertion point for per direct association field
		case "Animates":
			res := make(map[*Animate][]*Text)
			for text := range stage.Texts {
				for _, animate_ := range text.Animates {
					res[animate_] = append(res[animate_], text)
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
	case *Animate:
		res = any(new(Animate)).(Type)
	case *Circle:
		res = any(new(Circle)).(Type)
	case *Condition:
		res = any(new(Condition)).(Type)
	case *ControlPoint:
		res = any(new(ControlPoint)).(Type)
	case *Ellipse:
		res = any(new(Ellipse)).(Type)
	case *FileToDownload:
		res = any(new(FileToDownload)).(Type)
	case *Layer:
		res = any(new(Layer)).(Type)
	case *Line:
		res = any(new(Line)).(Type)
	case *Link:
		res = any(new(Link)).(Type)
	case *LinkAnchoredPath:
		res = any(new(LinkAnchoredPath)).(Type)
	case *LinkAnchoredText:
		res = any(new(LinkAnchoredText)).(Type)
	case *Path:
		res = any(new(Path)).(Type)
	case *Point:
		res = any(new(Point)).(Type)
	case *Polygone:
		res = any(new(Polygone)).(Type)
	case *Polyline:
		res = any(new(Polyline)).(Type)
	case *Rect:
		res = any(new(Rect)).(Type)
	case *RectAnchoredPath:
		res = any(new(RectAnchoredPath)).(Type)
	case *RectAnchoredPngImage:
		res = any(new(RectAnchoredPngImage)).(Type)
	case *RectAnchoredRect:
		res = any(new(RectAnchoredRect)).(Type)
	case *RectAnchoredText:
		res = any(new(RectAnchoredText)).(Type)
	case *RectLinkLink:
		res = any(new(RectLinkLink)).(Type)
	case *SVG:
		res = any(new(SVG)).(Type)
	case *SvgText:
		res = any(new(SvgText)).(Type)
	case *Text:
		res = any(new(Text)).(Type)
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
	case *Animate:
		res = "Animate"
	case *Circle:
		res = "Circle"
	case *Condition:
		res = "Condition"
	case *ControlPoint:
		res = "ControlPoint"
	case *Ellipse:
		res = "Ellipse"
	case *FileToDownload:
		res = "FileToDownload"
	case *Layer:
		res = "Layer"
	case *Line:
		res = "Line"
	case *Link:
		res = "Link"
	case *LinkAnchoredPath:
		res = "LinkAnchoredPath"
	case *LinkAnchoredText:
		res = "LinkAnchoredText"
	case *Path:
		res = "Path"
	case *Point:
		res = "Point"
	case *Polygone:
		res = "Polygone"
	case *Polyline:
		res = "Polyline"
	case *Rect:
		res = "Rect"
	case *RectAnchoredPath:
		res = "RectAnchoredPath"
	case *RectAnchoredPngImage:
		res = "RectAnchoredPngImage"
	case *RectAnchoredRect:
		res = "RectAnchoredRect"
	case *RectAnchoredText:
		res = "RectAnchoredText"
	case *RectLinkLink:
		res = "RectLinkLink"
	case *SVG:
		res = "SVG"
	case *SvgText:
		res = "SvgText"
	case *Text:
		res = "Text"
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
	case *Animate:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Circle"
		rf.Fieldname = "Animations"
		res = append(res, rf)
		rf.GongstructName = "Ellipse"
		rf.Fieldname = "Animates"
		res = append(res, rf)
		rf.GongstructName = "Line"
		rf.Fieldname = "Animates"
		res = append(res, rf)
		rf.GongstructName = "LinkAnchoredText"
		rf.Fieldname = "Animates"
		res = append(res, rf)
		rf.GongstructName = "Path"
		rf.Fieldname = "Animates"
		res = append(res, rf)
		rf.GongstructName = "Polygone"
		rf.Fieldname = "Animates"
		res = append(res, rf)
		rf.GongstructName = "Polyline"
		rf.Fieldname = "Animates"
		res = append(res, rf)
		rf.GongstructName = "Rect"
		rf.Fieldname = "Animations"
		res = append(res, rf)
		rf.GongstructName = "RectAnchoredText"
		rf.Fieldname = "Animates"
		res = append(res, rf)
		rf.GongstructName = "Text"
		rf.Fieldname = "Animates"
		res = append(res, rf)
	case *Circle:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layer"
		rf.Fieldname = "Circles"
		res = append(res, rf)
	case *Condition:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Rect"
		rf.Fieldname = "HoveringTrigger"
		res = append(res, rf)
		rf.GongstructName = "Rect"
		rf.Fieldname = "DisplayConditions"
		res = append(res, rf)
	case *ControlPoint:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Link"
		rf.Fieldname = "ControlPoints"
		res = append(res, rf)
	case *Ellipse:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layer"
		rf.Fieldname = "Ellipses"
		res = append(res, rf)
	case *FileToDownload:
		var rf ReverseField
		_ = rf
	case *Layer:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SVG"
		rf.Fieldname = "Layers"
		res = append(res, rf)
	case *Line:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layer"
		rf.Fieldname = "Lines"
		res = append(res, rf)
	case *Link:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layer"
		rf.Fieldname = "Links"
		res = append(res, rf)
	case *LinkAnchoredPath:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Link"
		rf.Fieldname = "PathAtArrowStart"
		res = append(res, rf)
		rf.GongstructName = "Link"
		rf.Fieldname = "PathAtArrowEnd"
		res = append(res, rf)
		rf.GongstructName = "Link"
		rf.Fieldname = "PathAtCorner"
		res = append(res, rf)
	case *LinkAnchoredText:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Link"
		rf.Fieldname = "TextAtArrowStart"
		res = append(res, rf)
		rf.GongstructName = "Link"
		rf.Fieldname = "TextAtArrowEnd"
		res = append(res, rf)
		rf.GongstructName = "Link"
		rf.Fieldname = "TextAtCorner"
		res = append(res, rf)
	case *Path:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layer"
		rf.Fieldname = "Paths"
		res = append(res, rf)
	case *Point:
		var rf ReverseField
		_ = rf
	case *Polygone:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layer"
		rf.Fieldname = "Polygones"
		res = append(res, rf)
	case *Polyline:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layer"
		rf.Fieldname = "Polylines"
		res = append(res, rf)
	case *Rect:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layer"
		rf.Fieldname = "Rects"
		res = append(res, rf)
		rf.GongstructName = "Rect"
		rf.Fieldname = "Peers"
		res = append(res, rf)
		rf.GongstructName = "Rect"
		rf.Fieldname = "Obstacles"
		res = append(res, rf)
	case *RectAnchoredPath:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Rect"
		rf.Fieldname = "RectAnchoredPaths"
		res = append(res, rf)
	case *RectAnchoredPngImage:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Rect"
		rf.Fieldname = "RectAnchoredPngImages"
		res = append(res, rf)
	case *RectAnchoredRect:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Rect"
		rf.Fieldname = "RectAnchoredRects"
		res = append(res, rf)
	case *RectAnchoredText:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Rect"
		rf.Fieldname = "RectAnchoredTexts"
		res = append(res, rf)
	case *RectLinkLink:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layer"
		rf.Fieldname = "RectLinkLinks"
		res = append(res, rf)
	case *SVG:
		var rf ReverseField
		_ = rf
	case *SvgText:
		var rf ReverseField
		_ = rf
	case *Text:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layer"
		rf.Fieldname = "Texts"
		res = append(res, rf)
	}
	return
}

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
}

// insertion point for get fields header method
func (animate *Animate) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "AttributeName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Values",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "From",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "To",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Dur",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RepeatCount",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (circle *Circle) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "CX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "CY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Radius",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Animations",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Animate",
		},
	}
	return
}

func (condition *Condition) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (controlpoint *ControlPoint) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "X_Relative",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y_Relative",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "ClosestRect",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Rect",
		},
	}
	return
}

func (ellipse *Ellipse) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "CX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "CY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "RX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "RY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Animates",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Animate",
		},
	}
	return
}

func (filetodownload *FileToDownload) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Base64EncodedContent",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (layer *Layer) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Rects",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Rect",
		},
		{
			Name:                 "Texts",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Text",
		},
		{
			Name:                 "Circles",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Circle",
		},
		{
			Name:                 "Lines",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Line",
		},
		{
			Name:                 "Ellipses",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Ellipse",
		},
		{
			Name:                 "Polylines",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Polyline",
		},
		{
			Name:                 "Polygones",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Polygone",
		},
		{
			Name:                 "Paths",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Path",
		},
		{
			Name:                 "Links",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Link",
		},
		{
			Name:                 "RectLinkLinks",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "RectLinkLink",
		},
	}
	return
}

func (line *Line) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "X1",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y1",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "X2",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y2",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Animates",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Animate",
		},
		{
			Name:               "MouseClickX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "MouseClickY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (link *Link) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Type",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "LinkType",
		},
		{
			Name:               "IsBezierCurve",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Start",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Rect",
		},
		{
			Name:                 "StartAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "AnchorType",
		},
		{
			Name:                 "End",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Rect",
		},
		{
			Name:                 "EndAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "AnchorType",
		},
		{
			Name:                 "StartOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:               "StartRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "EndOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:               "EndRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "CornerOffsetRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "CornerRadius",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "HasEndArrow",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "EndArrowSize",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndArrowOffset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "HasStartArrow",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "StartArrowSize",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StartArrowOffset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "TextAtArrowStart",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "LinkAnchoredText",
		},
		{
			Name:                 "TextAtArrowEnd",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "LinkAnchoredText",
		},
		{
			Name:                 "TextAtCorner",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "LinkAnchoredText",
		},
		{
			Name:                 "PathAtArrowStart",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "LinkAnchoredPath",
		},
		{
			Name:                 "PathAtArrowEnd",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "LinkAnchoredPath",
		},
		{
			Name:                 "PathAtCorner",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "LinkAnchoredPath",
		},
		{
			Name:                 "ControlPoints",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlPoint",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MouseX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "MouseY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "MouseEventKey",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "MouseEventKey",
		},
	}
	return
}

func (linkanchoredpath *LinkAnchoredPath) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Definition",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "X_Offset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y_Offset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ScalePropotionnally",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "AppliedScaling",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (linkanchoredtext *LinkAnchoredText) GongGetFieldHeaders() (res []GongFieldHeader) {
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
		{
			Name:               "AutomaticLayout",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LinkAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "LinkAnchorType",
		},
		{
			Name:               "X_Offset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y_Offset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "FontWeight",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontStyle",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "LetterSpacing",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "WhiteSpace",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "WhiteSpaceEnum",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Animates",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Animate",
		},
	}
	return
}

func (path *Path) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Definition",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Animates",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Animate",
		},
	}
	return
}

func (point *Point) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (polygone *Polygone) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Points",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Animates",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Animate",
		},
	}
	return
}

func (polyline *Polyline) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Points",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Animates",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Animate",
		},
	}
	return
}

func (rect *Rect) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
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
			Name:               "RX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "Peers",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Rect",
		},
		{
			Name:                 "EnclosingRect",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Rect",
		},
		{
			Name:                 "Obstacles",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Rect",
		},
		{
			Name:                 "AnchoredTo",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Rect",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "HoveringTrigger",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Condition",
		},
		{
			Name:                 "DisplayConditions",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Condition",
		},
		{
			Name:                 "Animations",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Animate",
		},
		{
			Name:               "IsSelectable",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsSelected",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "CanHaveLeftHandle",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "HasLeftHandle",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "CanHaveRightHandle",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "HasRightHandle",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "CanHaveTopHandle",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "HasTopHandle",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsScalingProportionally",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "CanHaveBottomHandle",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "HasBottomHandle",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "CanMoveHorizontaly",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "CanMoveVerticaly",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "RectAnchoredTexts",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "RectAnchoredText",
		},
		{
			Name:                 "RectAnchoredRects",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "RectAnchoredRect",
		},
		{
			Name:                 "RectAnchoredPaths",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "RectAnchoredPath",
		},
		{
			Name:                 "RectAnchoredPngImages",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "RectAnchoredPngImage",
		},
		{
			Name:               "ChangeColorWhenHovered",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ColorWhenHovered",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "OriginalColor",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacityWhenHovered",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "OriginalFillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
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
			Name:                 "ToolTipPosition",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "ToolTipPositionEnum",
		},
		{
			Name:               "MouseX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "MouseY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "MouseEventKey",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "MouseEventKey",
		},
		{
			Name:               "URLPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "URLTarget",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "LinkTargetType",
		},
	}
	return
}

func (rectanchoredpath *RectAnchoredPath) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Definition",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "X_Offset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y_Offset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "RectAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "RectAnchorType",
		},
		{
			Name:               "ScalePropotionnally",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "AppliedScaling",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (rectanchoredpngimage *RectAnchoredPngImage) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
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
			Name:               "RX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "X_Offset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y_Offset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "RectAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "RectAnchorType",
		},
		{
			Name:               "Base64Content",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (rectanchoredrect *RectAnchoredRect) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
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
			Name:               "RX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "X_Offset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y_Offset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "RectAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "RectAnchorType",
		},
		{
			Name:               "WidthFollowRect",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "HeightFollowRect",
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
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (rectanchoredtext *RectAnchoredText) GongGetFieldHeaders() (res []GongFieldHeader) {
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
		{
			Name:               "FontWeight",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontStyle",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "LetterSpacing",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "WhiteSpace",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "WhiteSpaceEnum",
		},
		{
			Name:               "X_Offset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y_Offset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "RectAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "RectAnchorType",
		},
		{
			Name:                 "TextAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "TextAnchorType",
		},
		{
			Name:                 "DominantBaseline",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "DominantBaselineType",
		},
		{
			Name:                 "WritingMode",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "WritingMode",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Animates",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Animate",
		},
		{
			Name:               "URLPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "URLTarget",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "LinkTargetType",
		},
	}
	return
}

func (rectlinklink *RectLinkLink) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Start",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Rect",
		},
		{
			Name:                 "End",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Link",
		},
		{
			Name:               "TargetAnchorPosition",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (svg *SVG) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Layers",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Layer",
		},
		{
			Name:                 "DrawingState",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "DrawingState",
		},
		{
			Name:                 "StartRect",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Rect",
		},
		{
			Name:                 "EndRect",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Rect",
		},
		{
			Name:               "IsEditable",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsSVGFrontEndFileGenerated",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsSVGBackEndFileGenerated",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "DefaultDirectoryForGeneratedImages",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsControlBannerHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "PanX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "PanY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Zoom",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "OverrideWidth",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "OverriddenWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "OverrideHeight",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "OverriddenHeight",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (svgtext *SvgText) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Text",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (text *Text) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
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
			Name:               "Content",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontWeight",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontStyle",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "LetterSpacing",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "WhiteSpace",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "WhiteSpaceEnum",
		},
		{
			Name:                 "Animates",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Animate",
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
func (animate *Animate) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = animate.Name
	case "AttributeName":
		res.valueString = animate.AttributeName
	case "Values":
		res.valueString = animate.Values
	case "From":
		res.valueString = animate.From
	case "To":
		res.valueString = animate.To
	case "Dur":
		res.valueString = animate.Dur
	case "RepeatCount":
		res.valueString = animate.RepeatCount
	}
	return
}

func (circle *Circle) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = circle.Name
	case "CX":
		res.valueString = fmt.Sprintf("%f", circle.CX)
		res.valueFloat = circle.CX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "CY":
		res.valueString = fmt.Sprintf("%f", circle.CY)
		res.valueFloat = circle.CY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Radius":
		res.valueString = fmt.Sprintf("%f", circle.Radius)
		res.valueFloat = circle.Radius
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Color":
		res.valueString = circle.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", circle.FillOpacity)
		res.valueFloat = circle.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = circle.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", circle.StrokeOpacity)
		res.valueFloat = circle.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", circle.StrokeWidth)
		res.valueFloat = circle.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = circle.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = circle.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = circle.Transform
	case "Animations":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range circle.Animations {
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

func (condition *Condition) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = condition.Name
	}
	return
}

func (controlpoint *ControlPoint) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = controlpoint.Name
	case "X_Relative":
		res.valueString = fmt.Sprintf("%f", controlpoint.X_Relative)
		res.valueFloat = controlpoint.X_Relative
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y_Relative":
		res.valueString = fmt.Sprintf("%f", controlpoint.Y_Relative)
		res.valueFloat = controlpoint.Y_Relative
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ClosestRect":
		res.GongFieldValueType = GongFieldValueTypePointer
		if controlpoint.ClosestRect != nil {
			res.valueString = controlpoint.ClosestRect.Name
			res.ids = controlpoint.ClosestRect.GongGetUUID(stage)
		}
	}
	return
}

func (ellipse *Ellipse) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ellipse.Name
	case "CX":
		res.valueString = fmt.Sprintf("%f", ellipse.CX)
		res.valueFloat = ellipse.CX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "CY":
		res.valueString = fmt.Sprintf("%f", ellipse.CY)
		res.valueFloat = ellipse.CY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RX":
		res.valueString = fmt.Sprintf("%f", ellipse.RX)
		res.valueFloat = ellipse.RX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RY":
		res.valueString = fmt.Sprintf("%f", ellipse.RY)
		res.valueFloat = ellipse.RY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Color":
		res.valueString = ellipse.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", ellipse.FillOpacity)
		res.valueFloat = ellipse.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = ellipse.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", ellipse.StrokeOpacity)
		res.valueFloat = ellipse.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", ellipse.StrokeWidth)
		res.valueFloat = ellipse.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = ellipse.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = ellipse.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = ellipse.Transform
	case "Animates":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ellipse.Animates {
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

func (filetodownload *FileToDownload) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = filetodownload.Name
	case "Base64EncodedContent":
		res.valueString = filetodownload.Base64EncodedContent
	}
	return
}

func (layer *Layer) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = layer.Name
	case "Rects":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.Rects {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Texts":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.Texts {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Circles":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.Circles {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Lines":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.Lines {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Ellipses":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.Ellipses {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Polylines":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.Polylines {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Polygones":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.Polygones {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Paths":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.Paths {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Links":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.Links {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RectLinkLinks":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.RectLinkLinks {
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

func (line *Line) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = line.Name
	case "X1":
		res.valueString = fmt.Sprintf("%f", line.X1)
		res.valueFloat = line.X1
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y1":
		res.valueString = fmt.Sprintf("%f", line.Y1)
		res.valueFloat = line.Y1
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "X2":
		res.valueString = fmt.Sprintf("%f", line.X2)
		res.valueFloat = line.X2
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y2":
		res.valueString = fmt.Sprintf("%f", line.Y2)
		res.valueFloat = line.Y2
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Color":
		res.valueString = line.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", line.FillOpacity)
		res.valueFloat = line.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = line.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", line.StrokeOpacity)
		res.valueFloat = line.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", line.StrokeWidth)
		res.valueFloat = line.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = line.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = line.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = line.Transform
	case "Animates":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range line.Animates {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "MouseClickX":
		res.valueString = fmt.Sprintf("%f", line.MouseClickX)
		res.valueFloat = line.MouseClickX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "MouseClickY":
		res.valueString = fmt.Sprintf("%f", line.MouseClickY)
		res.valueFloat = line.MouseClickY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (link *Link) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = link.Name
	case "Type":
		enum := link.Type
		res.valueString = enum.ToCodeString()
	case "IsBezierCurve":
		res.valueString = fmt.Sprintf("%t", link.IsBezierCurve)
		res.valueBool = link.IsBezierCurve
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Start":
		res.GongFieldValueType = GongFieldValueTypePointer
		if link.Start != nil {
			res.valueString = link.Start.Name
			res.ids = link.Start.GongGetUUID(stage)
		}
	case "StartAnchorType":
		enum := link.StartAnchorType
		res.valueString = enum.ToCodeString()
	case "End":
		res.GongFieldValueType = GongFieldValueTypePointer
		if link.End != nil {
			res.valueString = link.End.Name
			res.ids = link.End.GongGetUUID(stage)
		}
	case "EndAnchorType":
		enum := link.EndAnchorType
		res.valueString = enum.ToCodeString()
	case "StartOrientation":
		enum := link.StartOrientation
		res.valueString = enum.ToCodeString()
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", link.StartRatio)
		res.valueFloat = link.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndOrientation":
		enum := link.EndOrientation
		res.valueString = enum.ToCodeString()
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", link.EndRatio)
		res.valueFloat = link.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", link.CornerOffsetRatio)
		res.valueFloat = link.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "CornerRadius":
		res.valueString = fmt.Sprintf("%f", link.CornerRadius)
		res.valueFloat = link.CornerRadius
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "HasEndArrow":
		res.valueString = fmt.Sprintf("%t", link.HasEndArrow)
		res.valueBool = link.HasEndArrow
		res.GongFieldValueType = GongFieldValueTypeBool
	case "EndArrowSize":
		res.valueString = fmt.Sprintf("%f", link.EndArrowSize)
		res.valueFloat = link.EndArrowSize
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndArrowOffset":
		res.valueString = fmt.Sprintf("%f", link.EndArrowOffset)
		res.valueFloat = link.EndArrowOffset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "HasStartArrow":
		res.valueString = fmt.Sprintf("%t", link.HasStartArrow)
		res.valueBool = link.HasStartArrow
		res.GongFieldValueType = GongFieldValueTypeBool
	case "StartArrowSize":
		res.valueString = fmt.Sprintf("%f", link.StartArrowSize)
		res.valueFloat = link.StartArrowSize
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartArrowOffset":
		res.valueString = fmt.Sprintf("%f", link.StartArrowOffset)
		res.valueFloat = link.StartArrowOffset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "TextAtArrowStart":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range link.TextAtArrowStart {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "TextAtArrowEnd":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range link.TextAtArrowEnd {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "TextAtCorner":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range link.TextAtCorner {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PathAtArrowStart":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range link.PathAtArrowStart {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PathAtArrowEnd":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range link.PathAtArrowEnd {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PathAtCorner":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range link.PathAtCorner {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ControlPoints":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range link.ControlPoints {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Color":
		res.valueString = link.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", link.FillOpacity)
		res.valueFloat = link.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = link.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", link.StrokeOpacity)
		res.valueFloat = link.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", link.StrokeWidth)
		res.valueFloat = link.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = link.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = link.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = link.Transform
	case "MouseX":
		res.valueString = fmt.Sprintf("%f", link.MouseX)
		res.valueFloat = link.MouseX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "MouseY":
		res.valueString = fmt.Sprintf("%f", link.MouseY)
		res.valueFloat = link.MouseY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "MouseEventKey":
		enum := link.MouseEventKey
		res.valueString = enum.ToCodeString()
	}
	return
}

func (linkanchoredpath *LinkAnchoredPath) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = linkanchoredpath.Name
	case "Definition":
		res.valueString = linkanchoredpath.Definition
	case "X_Offset":
		res.valueString = fmt.Sprintf("%f", linkanchoredpath.X_Offset)
		res.valueFloat = linkanchoredpath.X_Offset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y_Offset":
		res.valueString = fmt.Sprintf("%f", linkanchoredpath.Y_Offset)
		res.valueFloat = linkanchoredpath.Y_Offset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ScalePropotionnally":
		res.valueString = fmt.Sprintf("%t", linkanchoredpath.ScalePropotionnally)
		res.valueBool = linkanchoredpath.ScalePropotionnally
		res.GongFieldValueType = GongFieldValueTypeBool
	case "AppliedScaling":
		res.valueString = fmt.Sprintf("%f", linkanchoredpath.AppliedScaling)
		res.valueFloat = linkanchoredpath.AppliedScaling
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Color":
		res.valueString = linkanchoredpath.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", linkanchoredpath.FillOpacity)
		res.valueFloat = linkanchoredpath.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = linkanchoredpath.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", linkanchoredpath.StrokeOpacity)
		res.valueFloat = linkanchoredpath.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", linkanchoredpath.StrokeWidth)
		res.valueFloat = linkanchoredpath.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = linkanchoredpath.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = linkanchoredpath.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = linkanchoredpath.Transform
	}
	return
}

func (linkanchoredtext *LinkAnchoredText) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = linkanchoredtext.Name
	case "Content":
		res.valueString = linkanchoredtext.Content
	case "AutomaticLayout":
		res.valueString = fmt.Sprintf("%t", linkanchoredtext.AutomaticLayout)
		res.valueBool = linkanchoredtext.AutomaticLayout
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LinkAnchorType":
		enum := linkanchoredtext.LinkAnchorType
		res.valueString = enum.ToCodeString()
	case "X_Offset":
		res.valueString = fmt.Sprintf("%f", linkanchoredtext.X_Offset)
		res.valueFloat = linkanchoredtext.X_Offset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y_Offset":
		res.valueString = fmt.Sprintf("%f", linkanchoredtext.Y_Offset)
		res.valueFloat = linkanchoredtext.Y_Offset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "FontWeight":
		res.valueString = linkanchoredtext.FontWeight
	case "FontSize":
		res.valueString = linkanchoredtext.FontSize
	case "FontStyle":
		res.valueString = linkanchoredtext.FontStyle
	case "LetterSpacing":
		res.valueString = linkanchoredtext.LetterSpacing
	case "FontFamily":
		res.valueString = linkanchoredtext.FontFamily
	case "WhiteSpace":
		enum := linkanchoredtext.WhiteSpace
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = linkanchoredtext.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", linkanchoredtext.FillOpacity)
		res.valueFloat = linkanchoredtext.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = linkanchoredtext.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", linkanchoredtext.StrokeOpacity)
		res.valueFloat = linkanchoredtext.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", linkanchoredtext.StrokeWidth)
		res.valueFloat = linkanchoredtext.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = linkanchoredtext.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = linkanchoredtext.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = linkanchoredtext.Transform
	case "Animates":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range linkanchoredtext.Animates {
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

func (path *Path) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = path.Name
	case "Definition":
		res.valueString = path.Definition
	case "Color":
		res.valueString = path.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", path.FillOpacity)
		res.valueFloat = path.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = path.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", path.StrokeOpacity)
		res.valueFloat = path.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", path.StrokeWidth)
		res.valueFloat = path.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = path.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = path.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = path.Transform
	case "Animates":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range path.Animates {
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

func (point *Point) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = point.Name
	case "X":
		res.valueString = fmt.Sprintf("%f", point.X)
		res.valueFloat = point.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", point.Y)
		res.valueFloat = point.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (polygone *Polygone) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = polygone.Name
	case "Points":
		res.valueString = polygone.Points
	case "Color":
		res.valueString = polygone.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", polygone.FillOpacity)
		res.valueFloat = polygone.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = polygone.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", polygone.StrokeOpacity)
		res.valueFloat = polygone.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", polygone.StrokeWidth)
		res.valueFloat = polygone.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = polygone.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = polygone.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = polygone.Transform
	case "Animates":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range polygone.Animates {
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

func (polyline *Polyline) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = polyline.Name
	case "Points":
		res.valueString = polyline.Points
	case "Color":
		res.valueString = polyline.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", polyline.FillOpacity)
		res.valueFloat = polyline.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = polyline.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", polyline.StrokeOpacity)
		res.valueFloat = polyline.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", polyline.StrokeWidth)
		res.valueFloat = polyline.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = polyline.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = polyline.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = polyline.Transform
	case "Animates":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range polyline.Animates {
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

func (rect *Rect) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = rect.Name
	case "X":
		res.valueString = fmt.Sprintf("%f", rect.X)
		res.valueFloat = rect.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", rect.Y)
		res.valueFloat = rect.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", rect.Width)
		res.valueFloat = rect.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", rect.Height)
		res.valueFloat = rect.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RX":
		res.valueString = fmt.Sprintf("%f", rect.RX)
		res.valueFloat = rect.RX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Peers":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range rect.Peers {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "EnclosingRect":
		res.GongFieldValueType = GongFieldValueTypePointer
		if rect.EnclosingRect != nil {
			res.valueString = rect.EnclosingRect.Name
			res.ids = rect.EnclosingRect.GongGetUUID(stage)
		}
	case "Obstacles":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range rect.Obstacles {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "AnchoredTo":
		res.GongFieldValueType = GongFieldValueTypePointer
		if rect.AnchoredTo != nil {
			res.valueString = rect.AnchoredTo.Name
			res.ids = rect.AnchoredTo.GongGetUUID(stage)
		}
	case "Color":
		res.valueString = rect.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", rect.FillOpacity)
		res.valueFloat = rect.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = rect.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", rect.StrokeOpacity)
		res.valueFloat = rect.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", rect.StrokeWidth)
		res.valueFloat = rect.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = rect.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = rect.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = rect.Transform
	case "HoveringTrigger":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range rect.HoveringTrigger {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DisplayConditions":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range rect.DisplayConditions {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Animations":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range rect.Animations {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsSelectable":
		res.valueString = fmt.Sprintf("%t", rect.IsSelectable)
		res.valueBool = rect.IsSelectable
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsSelected":
		res.valueString = fmt.Sprintf("%t", rect.IsSelected)
		res.valueBool = rect.IsSelected
		res.GongFieldValueType = GongFieldValueTypeBool
	case "CanHaveLeftHandle":
		res.valueString = fmt.Sprintf("%t", rect.CanHaveLeftHandle)
		res.valueBool = rect.CanHaveLeftHandle
		res.GongFieldValueType = GongFieldValueTypeBool
	case "HasLeftHandle":
		res.valueString = fmt.Sprintf("%t", rect.HasLeftHandle)
		res.valueBool = rect.HasLeftHandle
		res.GongFieldValueType = GongFieldValueTypeBool
	case "CanHaveRightHandle":
		res.valueString = fmt.Sprintf("%t", rect.CanHaveRightHandle)
		res.valueBool = rect.CanHaveRightHandle
		res.GongFieldValueType = GongFieldValueTypeBool
	case "HasRightHandle":
		res.valueString = fmt.Sprintf("%t", rect.HasRightHandle)
		res.valueBool = rect.HasRightHandle
		res.GongFieldValueType = GongFieldValueTypeBool
	case "CanHaveTopHandle":
		res.valueString = fmt.Sprintf("%t", rect.CanHaveTopHandle)
		res.valueBool = rect.CanHaveTopHandle
		res.GongFieldValueType = GongFieldValueTypeBool
	case "HasTopHandle":
		res.valueString = fmt.Sprintf("%t", rect.HasTopHandle)
		res.valueBool = rect.HasTopHandle
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsScalingProportionally":
		res.valueString = fmt.Sprintf("%t", rect.IsScalingProportionally)
		res.valueBool = rect.IsScalingProportionally
		res.GongFieldValueType = GongFieldValueTypeBool
	case "CanHaveBottomHandle":
		res.valueString = fmt.Sprintf("%t", rect.CanHaveBottomHandle)
		res.valueBool = rect.CanHaveBottomHandle
		res.GongFieldValueType = GongFieldValueTypeBool
	case "HasBottomHandle":
		res.valueString = fmt.Sprintf("%t", rect.HasBottomHandle)
		res.valueBool = rect.HasBottomHandle
		res.GongFieldValueType = GongFieldValueTypeBool
	case "CanMoveHorizontaly":
		res.valueString = fmt.Sprintf("%t", rect.CanMoveHorizontaly)
		res.valueBool = rect.CanMoveHorizontaly
		res.GongFieldValueType = GongFieldValueTypeBool
	case "CanMoveVerticaly":
		res.valueString = fmt.Sprintf("%t", rect.CanMoveVerticaly)
		res.valueBool = rect.CanMoveVerticaly
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RectAnchoredTexts":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range rect.RectAnchoredTexts {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RectAnchoredRects":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range rect.RectAnchoredRects {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RectAnchoredPaths":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range rect.RectAnchoredPaths {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RectAnchoredPngImages":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range rect.RectAnchoredPngImages {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ChangeColorWhenHovered":
		res.valueString = fmt.Sprintf("%t", rect.ChangeColorWhenHovered)
		res.valueBool = rect.ChangeColorWhenHovered
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ColorWhenHovered":
		res.valueString = rect.ColorWhenHovered
	case "OriginalColor":
		res.valueString = rect.OriginalColor
	case "FillOpacityWhenHovered":
		res.valueString = fmt.Sprintf("%f", rect.FillOpacityWhenHovered)
		res.valueFloat = rect.FillOpacityWhenHovered
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "OriginalFillOpacity":
		res.valueString = fmt.Sprintf("%f", rect.OriginalFillOpacity)
		res.valueFloat = rect.OriginalFillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "HasToolTip":
		res.valueString = fmt.Sprintf("%t", rect.HasToolTip)
		res.valueBool = rect.HasToolTip
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ToolTipText":
		res.valueString = rect.ToolTipText
	case "ToolTipPosition":
		enum := rect.ToolTipPosition
		res.valueString = enum.ToCodeString()
	case "MouseX":
		res.valueString = fmt.Sprintf("%f", rect.MouseX)
		res.valueFloat = rect.MouseX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "MouseY":
		res.valueString = fmt.Sprintf("%f", rect.MouseY)
		res.valueFloat = rect.MouseY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "MouseEventKey":
		enum := rect.MouseEventKey
		res.valueString = enum.ToCodeString()
	case "URLPath":
		res.valueString = rect.URLPath
	case "URLTarget":
		enum := rect.URLTarget
		res.valueString = enum.ToCodeString()
	}
	return
}

func (rectanchoredpath *RectAnchoredPath) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = rectanchoredpath.Name
	case "Definition":
		res.valueString = rectanchoredpath.Definition
	case "X_Offset":
		res.valueString = fmt.Sprintf("%f", rectanchoredpath.X_Offset)
		res.valueFloat = rectanchoredpath.X_Offset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y_Offset":
		res.valueString = fmt.Sprintf("%f", rectanchoredpath.Y_Offset)
		res.valueFloat = rectanchoredpath.Y_Offset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RectAnchorType":
		enum := rectanchoredpath.RectAnchorType
		res.valueString = enum.ToCodeString()
	case "ScalePropotionnally":
		res.valueString = fmt.Sprintf("%t", rectanchoredpath.ScalePropotionnally)
		res.valueBool = rectanchoredpath.ScalePropotionnally
		res.GongFieldValueType = GongFieldValueTypeBool
	case "AppliedScaling":
		res.valueString = fmt.Sprintf("%f", rectanchoredpath.AppliedScaling)
		res.valueFloat = rectanchoredpath.AppliedScaling
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Color":
		res.valueString = rectanchoredpath.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", rectanchoredpath.FillOpacity)
		res.valueFloat = rectanchoredpath.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = rectanchoredpath.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", rectanchoredpath.StrokeOpacity)
		res.valueFloat = rectanchoredpath.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", rectanchoredpath.StrokeWidth)
		res.valueFloat = rectanchoredpath.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = rectanchoredpath.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = rectanchoredpath.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = rectanchoredpath.Transform
	}
	return
}

func (rectanchoredpngimage *RectAnchoredPngImage) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = rectanchoredpngimage.Name
	case "X":
		res.valueString = fmt.Sprintf("%f", rectanchoredpngimage.X)
		res.valueFloat = rectanchoredpngimage.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", rectanchoredpngimage.Y)
		res.valueFloat = rectanchoredpngimage.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", rectanchoredpngimage.Width)
		res.valueFloat = rectanchoredpngimage.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", rectanchoredpngimage.Height)
		res.valueFloat = rectanchoredpngimage.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RX":
		res.valueString = fmt.Sprintf("%f", rectanchoredpngimage.RX)
		res.valueFloat = rectanchoredpngimage.RX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "X_Offset":
		res.valueString = fmt.Sprintf("%f", rectanchoredpngimage.X_Offset)
		res.valueFloat = rectanchoredpngimage.X_Offset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y_Offset":
		res.valueString = fmt.Sprintf("%f", rectanchoredpngimage.Y_Offset)
		res.valueFloat = rectanchoredpngimage.Y_Offset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RectAnchorType":
		enum := rectanchoredpngimage.RectAnchorType
		res.valueString = enum.ToCodeString()
	case "Base64Content":
		res.valueString = rectanchoredpngimage.Base64Content
	}
	return
}

func (rectanchoredrect *RectAnchoredRect) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = rectanchoredrect.Name
	case "X":
		res.valueString = fmt.Sprintf("%f", rectanchoredrect.X)
		res.valueFloat = rectanchoredrect.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", rectanchoredrect.Y)
		res.valueFloat = rectanchoredrect.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", rectanchoredrect.Width)
		res.valueFloat = rectanchoredrect.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", rectanchoredrect.Height)
		res.valueFloat = rectanchoredrect.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RX":
		res.valueString = fmt.Sprintf("%f", rectanchoredrect.RX)
		res.valueFloat = rectanchoredrect.RX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "X_Offset":
		res.valueString = fmt.Sprintf("%f", rectanchoredrect.X_Offset)
		res.valueFloat = rectanchoredrect.X_Offset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y_Offset":
		res.valueString = fmt.Sprintf("%f", rectanchoredrect.Y_Offset)
		res.valueFloat = rectanchoredrect.Y_Offset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RectAnchorType":
		enum := rectanchoredrect.RectAnchorType
		res.valueString = enum.ToCodeString()
	case "WidthFollowRect":
		res.valueString = fmt.Sprintf("%t", rectanchoredrect.WidthFollowRect)
		res.valueBool = rectanchoredrect.WidthFollowRect
		res.GongFieldValueType = GongFieldValueTypeBool
	case "HeightFollowRect":
		res.valueString = fmt.Sprintf("%t", rectanchoredrect.HeightFollowRect)
		res.valueBool = rectanchoredrect.HeightFollowRect
		res.GongFieldValueType = GongFieldValueTypeBool
	case "HasToolTip":
		res.valueString = fmt.Sprintf("%t", rectanchoredrect.HasToolTip)
		res.valueBool = rectanchoredrect.HasToolTip
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ToolTipText":
		res.valueString = rectanchoredrect.ToolTipText
	case "Color":
		res.valueString = rectanchoredrect.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", rectanchoredrect.FillOpacity)
		res.valueFloat = rectanchoredrect.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = rectanchoredrect.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", rectanchoredrect.StrokeOpacity)
		res.valueFloat = rectanchoredrect.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", rectanchoredrect.StrokeWidth)
		res.valueFloat = rectanchoredrect.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = rectanchoredrect.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = rectanchoredrect.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = rectanchoredrect.Transform
	}
	return
}

func (rectanchoredtext *RectAnchoredText) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = rectanchoredtext.Name
	case "Content":
		res.valueString = rectanchoredtext.Content
	case "FontWeight":
		res.valueString = rectanchoredtext.FontWeight
	case "FontSize":
		res.valueString = rectanchoredtext.FontSize
	case "FontStyle":
		res.valueString = rectanchoredtext.FontStyle
	case "LetterSpacing":
		res.valueString = rectanchoredtext.LetterSpacing
	case "FontFamily":
		res.valueString = rectanchoredtext.FontFamily
	case "WhiteSpace":
		enum := rectanchoredtext.WhiteSpace
		res.valueString = enum.ToCodeString()
	case "X_Offset":
		res.valueString = fmt.Sprintf("%f", rectanchoredtext.X_Offset)
		res.valueFloat = rectanchoredtext.X_Offset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y_Offset":
		res.valueString = fmt.Sprintf("%f", rectanchoredtext.Y_Offset)
		res.valueFloat = rectanchoredtext.Y_Offset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RectAnchorType":
		enum := rectanchoredtext.RectAnchorType
		res.valueString = enum.ToCodeString()
	case "TextAnchorType":
		enum := rectanchoredtext.TextAnchorType
		res.valueString = enum.ToCodeString()
	case "DominantBaseline":
		enum := rectanchoredtext.DominantBaseline
		res.valueString = enum.ToCodeString()
	case "WritingMode":
		enum := rectanchoredtext.WritingMode
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = rectanchoredtext.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", rectanchoredtext.FillOpacity)
		res.valueFloat = rectanchoredtext.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = rectanchoredtext.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", rectanchoredtext.StrokeOpacity)
		res.valueFloat = rectanchoredtext.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", rectanchoredtext.StrokeWidth)
		res.valueFloat = rectanchoredtext.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = rectanchoredtext.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = rectanchoredtext.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = rectanchoredtext.Transform
	case "Animates":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range rectanchoredtext.Animates {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "URLPath":
		res.valueString = rectanchoredtext.URLPath
	case "URLTarget":
		enum := rectanchoredtext.URLTarget
		res.valueString = enum.ToCodeString()
	}
	return
}

func (rectlinklink *RectLinkLink) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = rectlinklink.Name
	case "Start":
		res.GongFieldValueType = GongFieldValueTypePointer
		if rectlinklink.Start != nil {
			res.valueString = rectlinklink.Start.Name
			res.ids = rectlinklink.Start.GongGetUUID(stage)
		}
	case "End":
		res.GongFieldValueType = GongFieldValueTypePointer
		if rectlinklink.End != nil {
			res.valueString = rectlinklink.End.Name
			res.ids = rectlinklink.End.GongGetUUID(stage)
		}
	case "TargetAnchorPosition":
		res.valueString = fmt.Sprintf("%f", rectlinklink.TargetAnchorPosition)
		res.valueFloat = rectlinklink.TargetAnchorPosition
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Color":
		res.valueString = rectlinklink.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", rectlinklink.FillOpacity)
		res.valueFloat = rectlinklink.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = rectlinklink.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", rectlinklink.StrokeOpacity)
		res.valueFloat = rectlinklink.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", rectlinklink.StrokeWidth)
		res.valueFloat = rectlinklink.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = rectlinklink.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = rectlinklink.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = rectlinklink.Transform
	}
	return
}

func (svg *SVG) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = svg.Name
	case "Layers":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range svg.Layers {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DrawingState":
		enum := svg.DrawingState
		res.valueString = enum.ToCodeString()
	case "StartRect":
		res.GongFieldValueType = GongFieldValueTypePointer
		if svg.StartRect != nil {
			res.valueString = svg.StartRect.Name
			res.ids = svg.StartRect.GongGetUUID(stage)
		}
	case "EndRect":
		res.GongFieldValueType = GongFieldValueTypePointer
		if svg.EndRect != nil {
			res.valueString = svg.EndRect.Name
			res.ids = svg.EndRect.GongGetUUID(stage)
		}
	case "IsEditable":
		res.valueString = fmt.Sprintf("%t", svg.IsEditable)
		res.valueBool = svg.IsEditable
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsSVGFrontEndFileGenerated":
		res.valueString = fmt.Sprintf("%t", svg.IsSVGFrontEndFileGenerated)
		res.valueBool = svg.IsSVGFrontEndFileGenerated
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsSVGBackEndFileGenerated":
		res.valueString = fmt.Sprintf("%t", svg.IsSVGBackEndFileGenerated)
		res.valueBool = svg.IsSVGBackEndFileGenerated
		res.GongFieldValueType = GongFieldValueTypeBool
	case "DefaultDirectoryForGeneratedImages":
		res.valueString = svg.DefaultDirectoryForGeneratedImages
	case "IsControlBannerHidden":
		res.valueString = fmt.Sprintf("%t", svg.IsControlBannerHidden)
		res.valueBool = svg.IsControlBannerHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	case "PanX":
		res.valueString = fmt.Sprintf("%f", svg.PanX)
		res.valueFloat = svg.PanX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "PanY":
		res.valueString = fmt.Sprintf("%f", svg.PanY)
		res.valueFloat = svg.PanY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Zoom":
		res.valueString = fmt.Sprintf("%f", svg.Zoom)
		res.valueFloat = svg.Zoom
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "OverrideWidth":
		res.valueString = fmt.Sprintf("%t", svg.OverrideWidth)
		res.valueBool = svg.OverrideWidth
		res.GongFieldValueType = GongFieldValueTypeBool
	case "OverriddenWidth":
		res.valueString = fmt.Sprintf("%f", svg.OverriddenWidth)
		res.valueFloat = svg.OverriddenWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "OverrideHeight":
		res.valueString = fmt.Sprintf("%t", svg.OverrideHeight)
		res.valueBool = svg.OverrideHeight
		res.GongFieldValueType = GongFieldValueTypeBool
	case "OverriddenHeight":
		res.valueString = fmt.Sprintf("%f", svg.OverriddenHeight)
		res.valueFloat = svg.OverriddenHeight
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (svgtext *SvgText) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = svgtext.Name
	case "Text":
		res.valueString = svgtext.Text
	}
	return
}

func (text *Text) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = text.Name
	case "X":
		res.valueString = fmt.Sprintf("%f", text.X)
		res.valueFloat = text.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", text.Y)
		res.valueFloat = text.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Content":
		res.valueString = text.Content
	case "Color":
		res.valueString = text.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", text.FillOpacity)
		res.valueFloat = text.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = text.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", text.StrokeOpacity)
		res.valueFloat = text.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", text.StrokeWidth)
		res.valueFloat = text.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = text.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = text.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = text.Transform
	case "FontWeight":
		res.valueString = text.FontWeight
	case "FontSize":
		res.valueString = text.FontSize
	case "FontStyle":
		res.valueString = text.FontStyle
	case "LetterSpacing":
		res.valueString = text.LetterSpacing
	case "FontFamily":
		res.valueString = text.FontFamily
	case "WhiteSpace":
		enum := text.WhiteSpace
		res.valueString = enum.ToCodeString()
	case "Animates":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range text.Animates {
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

func (stage *Stage) GetFieldStringValueFromPointer(instance GongstructIF, fieldName string) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	return stage.GetFieldStringValueFromPointer(instance, fieldName)
}

// insertion point for generic get gongstruct name
func (animate *Animate) GongGetGongstructName() string {
	return "Animate"
}

func (circle *Circle) GongGetGongstructName() string {
	return "Circle"
}

func (condition *Condition) GongGetGongstructName() string {
	return "Condition"
}

func (controlpoint *ControlPoint) GongGetGongstructName() string {
	return "ControlPoint"
}

func (ellipse *Ellipse) GongGetGongstructName() string {
	return "Ellipse"
}

func (filetodownload *FileToDownload) GongGetGongstructName() string {
	return "FileToDownload"
}

func (layer *Layer) GongGetGongstructName() string {
	return "Layer"
}

func (line *Line) GongGetGongstructName() string {
	return "Line"
}

func (link *Link) GongGetGongstructName() string {
	return "Link"
}

func (linkanchoredpath *LinkAnchoredPath) GongGetGongstructName() string {
	return "LinkAnchoredPath"
}

func (linkanchoredtext *LinkAnchoredText) GongGetGongstructName() string {
	return "LinkAnchoredText"
}

func (path *Path) GongGetGongstructName() string {
	return "Path"
}

func (point *Point) GongGetGongstructName() string {
	return "Point"
}

func (polygone *Polygone) GongGetGongstructName() string {
	return "Polygone"
}

func (polyline *Polyline) GongGetGongstructName() string {
	return "Polyline"
}

func (rect *Rect) GongGetGongstructName() string {
	return "Rect"
}

func (rectanchoredpath *RectAnchoredPath) GongGetGongstructName() string {
	return "RectAnchoredPath"
}

func (rectanchoredpngimage *RectAnchoredPngImage) GongGetGongstructName() string {
	return "RectAnchoredPngImage"
}

func (rectanchoredrect *RectAnchoredRect) GongGetGongstructName() string {
	return "RectAnchoredRect"
}

func (rectanchoredtext *RectAnchoredText) GongGetGongstructName() string {
	return "RectAnchoredText"
}

func (rectlinklink *RectLinkLink) GongGetGongstructName() string {
	return "RectLinkLink"
}

func (svg *SVG) GongGetGongstructName() string {
	return "SVG"
}

func (svgtext *SvgText) GongGetGongstructName() string {
	return "SvgText"
}

func (text *Text) GongGetGongstructName() string {
	return "Text"
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
	__gong__rebuildMapString(stage.Animates, &stage.Animates_mapString)

	__gong__rebuildMapString(stage.Circles, &stage.Circles_mapString)

	__gong__rebuildMapString(stage.Conditions, &stage.Conditions_mapString)

	__gong__rebuildMapString(stage.ControlPoints, &stage.ControlPoints_mapString)

	__gong__rebuildMapString(stage.Ellipses, &stage.Ellipses_mapString)

	__gong__rebuildMapString(stage.FileToDownloads, &stage.FileToDownloads_mapString)

	__gong__rebuildMapString(stage.Layers, &stage.Layers_mapString)

	__gong__rebuildMapString(stage.Lines, &stage.Lines_mapString)

	__gong__rebuildMapString(stage.Links, &stage.Links_mapString)

	__gong__rebuildMapString(stage.LinkAnchoredPaths, &stage.LinkAnchoredPaths_mapString)

	__gong__rebuildMapString(stage.LinkAnchoredTexts, &stage.LinkAnchoredTexts_mapString)

	__gong__rebuildMapString(stage.Paths, &stage.Paths_mapString)

	__gong__rebuildMapString(stage.Points, &stage.Points_mapString)

	__gong__rebuildMapString(stage.Polygones, &stage.Polygones_mapString)

	__gong__rebuildMapString(stage.Polylines, &stage.Polylines_mapString)

	__gong__rebuildMapString(stage.Rects, &stage.Rects_mapString)

	__gong__rebuildMapString(stage.RectAnchoredPaths, &stage.RectAnchoredPaths_mapString)

	__gong__rebuildMapString(stage.RectAnchoredPngImages, &stage.RectAnchoredPngImages_mapString)

	__gong__rebuildMapString(stage.RectAnchoredRects, &stage.RectAnchoredRects_mapString)

	__gong__rebuildMapString(stage.RectAnchoredTexts, &stage.RectAnchoredTexts_mapString)

	__gong__rebuildMapString(stage.RectLinkLinks, &stage.RectLinkLinks_mapString)

	__gong__rebuildMapString(stage.SVGs, &stage.SVGs_mapString)

	__gong__rebuildMapString(stage.SvgTexts, &stage.SvgTexts_mapString)

	__gong__rebuildMapString(stage.Texts, &stage.Texts_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
