// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct {
	path string

	// insertion point for definition of arrays registering instances
	Animates           map[*Animate]any
	Animates_mapString map[string]*Animate

	// insertion point for slice of pointers maps
	OnAfterAnimateCreateCallback OnAfterCreateInterface[Animate]
	OnAfterAnimateUpdateCallback OnAfterUpdateInterface[Animate]
	OnAfterAnimateDeleteCallback OnAfterDeleteInterface[Animate]
	OnAfterAnimateReadCallback   OnAfterReadInterface[Animate]

	Circles           map[*Circle]any
	Circles_mapString map[string]*Circle

	// insertion point for slice of pointers maps
	Circle_Animations_reverseMap map[*Animate]*Circle

	OnAfterCircleCreateCallback OnAfterCreateInterface[Circle]
	OnAfterCircleUpdateCallback OnAfterUpdateInterface[Circle]
	OnAfterCircleDeleteCallback OnAfterDeleteInterface[Circle]
	OnAfterCircleReadCallback   OnAfterReadInterface[Circle]

	Ellipses           map[*Ellipse]any
	Ellipses_mapString map[string]*Ellipse

	// insertion point for slice of pointers maps
	Ellipse_Animates_reverseMap map[*Animate]*Ellipse

	OnAfterEllipseCreateCallback OnAfterCreateInterface[Ellipse]
	OnAfterEllipseUpdateCallback OnAfterUpdateInterface[Ellipse]
	OnAfterEllipseDeleteCallback OnAfterDeleteInterface[Ellipse]
	OnAfterEllipseReadCallback   OnAfterReadInterface[Ellipse]

	Layers           map[*Layer]any
	Layers_mapString map[string]*Layer

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

	OnAfterLayerCreateCallback OnAfterCreateInterface[Layer]
	OnAfterLayerUpdateCallback OnAfterUpdateInterface[Layer]
	OnAfterLayerDeleteCallback OnAfterDeleteInterface[Layer]
	OnAfterLayerReadCallback   OnAfterReadInterface[Layer]

	Lines           map[*Line]any
	Lines_mapString map[string]*Line

	// insertion point for slice of pointers maps
	Line_Animates_reverseMap map[*Animate]*Line

	OnAfterLineCreateCallback OnAfterCreateInterface[Line]
	OnAfterLineUpdateCallback OnAfterUpdateInterface[Line]
	OnAfterLineDeleteCallback OnAfterDeleteInterface[Line]
	OnAfterLineReadCallback   OnAfterReadInterface[Line]

	Links           map[*Link]any
	Links_mapString map[string]*Link

	// insertion point for slice of pointers maps
	Link_TextAtArrowEnd_reverseMap map[*LinkAnchoredText]*Link

	Link_TextAtArrowStart_reverseMap map[*LinkAnchoredText]*Link

	Link_ControlPoints_reverseMap map[*Point]*Link

	OnAfterLinkCreateCallback OnAfterCreateInterface[Link]
	OnAfterLinkUpdateCallback OnAfterUpdateInterface[Link]
	OnAfterLinkDeleteCallback OnAfterDeleteInterface[Link]
	OnAfterLinkReadCallback   OnAfterReadInterface[Link]

	LinkAnchoredTexts           map[*LinkAnchoredText]any
	LinkAnchoredTexts_mapString map[string]*LinkAnchoredText

	// insertion point for slice of pointers maps
	LinkAnchoredText_Animates_reverseMap map[*Animate]*LinkAnchoredText

	OnAfterLinkAnchoredTextCreateCallback OnAfterCreateInterface[LinkAnchoredText]
	OnAfterLinkAnchoredTextUpdateCallback OnAfterUpdateInterface[LinkAnchoredText]
	OnAfterLinkAnchoredTextDeleteCallback OnAfterDeleteInterface[LinkAnchoredText]
	OnAfterLinkAnchoredTextReadCallback   OnAfterReadInterface[LinkAnchoredText]

	Paths           map[*Path]any
	Paths_mapString map[string]*Path

	// insertion point for slice of pointers maps
	Path_Animates_reverseMap map[*Animate]*Path

	OnAfterPathCreateCallback OnAfterCreateInterface[Path]
	OnAfterPathUpdateCallback OnAfterUpdateInterface[Path]
	OnAfterPathDeleteCallback OnAfterDeleteInterface[Path]
	OnAfterPathReadCallback   OnAfterReadInterface[Path]

	Points           map[*Point]any
	Points_mapString map[string]*Point

	// insertion point for slice of pointers maps
	OnAfterPointCreateCallback OnAfterCreateInterface[Point]
	OnAfterPointUpdateCallback OnAfterUpdateInterface[Point]
	OnAfterPointDeleteCallback OnAfterDeleteInterface[Point]
	OnAfterPointReadCallback   OnAfterReadInterface[Point]

	Polygones           map[*Polygone]any
	Polygones_mapString map[string]*Polygone

	// insertion point for slice of pointers maps
	Polygone_Animates_reverseMap map[*Animate]*Polygone

	OnAfterPolygoneCreateCallback OnAfterCreateInterface[Polygone]
	OnAfterPolygoneUpdateCallback OnAfterUpdateInterface[Polygone]
	OnAfterPolygoneDeleteCallback OnAfterDeleteInterface[Polygone]
	OnAfterPolygoneReadCallback   OnAfterReadInterface[Polygone]

	Polylines           map[*Polyline]any
	Polylines_mapString map[string]*Polyline

	// insertion point for slice of pointers maps
	Polyline_Animates_reverseMap map[*Animate]*Polyline

	OnAfterPolylineCreateCallback OnAfterCreateInterface[Polyline]
	OnAfterPolylineUpdateCallback OnAfterUpdateInterface[Polyline]
	OnAfterPolylineDeleteCallback OnAfterDeleteInterface[Polyline]
	OnAfterPolylineReadCallback   OnAfterReadInterface[Polyline]

	Rects           map[*Rect]any
	Rects_mapString map[string]*Rect

	// insertion point for slice of pointers maps
	Rect_Animations_reverseMap map[*Animate]*Rect

	Rect_RectAnchoredTexts_reverseMap map[*RectAnchoredText]*Rect

	Rect_RectAnchoredRects_reverseMap map[*RectAnchoredRect]*Rect

	Rect_RectAnchoredPaths_reverseMap map[*RectAnchoredPath]*Rect

	OnAfterRectCreateCallback OnAfterCreateInterface[Rect]
	OnAfterRectUpdateCallback OnAfterUpdateInterface[Rect]
	OnAfterRectDeleteCallback OnAfterDeleteInterface[Rect]
	OnAfterRectReadCallback   OnAfterReadInterface[Rect]

	RectAnchoredPaths           map[*RectAnchoredPath]any
	RectAnchoredPaths_mapString map[string]*RectAnchoredPath

	// insertion point for slice of pointers maps
	OnAfterRectAnchoredPathCreateCallback OnAfterCreateInterface[RectAnchoredPath]
	OnAfterRectAnchoredPathUpdateCallback OnAfterUpdateInterface[RectAnchoredPath]
	OnAfterRectAnchoredPathDeleteCallback OnAfterDeleteInterface[RectAnchoredPath]
	OnAfterRectAnchoredPathReadCallback   OnAfterReadInterface[RectAnchoredPath]

	RectAnchoredRects           map[*RectAnchoredRect]any
	RectAnchoredRects_mapString map[string]*RectAnchoredRect

	// insertion point for slice of pointers maps
	OnAfterRectAnchoredRectCreateCallback OnAfterCreateInterface[RectAnchoredRect]
	OnAfterRectAnchoredRectUpdateCallback OnAfterUpdateInterface[RectAnchoredRect]
	OnAfterRectAnchoredRectDeleteCallback OnAfterDeleteInterface[RectAnchoredRect]
	OnAfterRectAnchoredRectReadCallback   OnAfterReadInterface[RectAnchoredRect]

	RectAnchoredTexts           map[*RectAnchoredText]any
	RectAnchoredTexts_mapString map[string]*RectAnchoredText

	// insertion point for slice of pointers maps
	RectAnchoredText_Animates_reverseMap map[*Animate]*RectAnchoredText

	OnAfterRectAnchoredTextCreateCallback OnAfterCreateInterface[RectAnchoredText]
	OnAfterRectAnchoredTextUpdateCallback OnAfterUpdateInterface[RectAnchoredText]
	OnAfterRectAnchoredTextDeleteCallback OnAfterDeleteInterface[RectAnchoredText]
	OnAfterRectAnchoredTextReadCallback   OnAfterReadInterface[RectAnchoredText]

	RectLinkLinks           map[*RectLinkLink]any
	RectLinkLinks_mapString map[string]*RectLinkLink

	// insertion point for slice of pointers maps
	OnAfterRectLinkLinkCreateCallback OnAfterCreateInterface[RectLinkLink]
	OnAfterRectLinkLinkUpdateCallback OnAfterUpdateInterface[RectLinkLink]
	OnAfterRectLinkLinkDeleteCallback OnAfterDeleteInterface[RectLinkLink]
	OnAfterRectLinkLinkReadCallback   OnAfterReadInterface[RectLinkLink]

	SVGs           map[*SVG]any
	SVGs_mapString map[string]*SVG

	// insertion point for slice of pointers maps
	SVG_Layers_reverseMap map[*Layer]*SVG

	OnAfterSVGCreateCallback OnAfterCreateInterface[SVG]
	OnAfterSVGUpdateCallback OnAfterUpdateInterface[SVG]
	OnAfterSVGDeleteCallback OnAfterDeleteInterface[SVG]
	OnAfterSVGReadCallback   OnAfterReadInterface[SVG]

	Texts           map[*Text]any
	Texts_mapString map[string]*Text

	// insertion point for slice of pointers maps
	Text_Animates_reverseMap map[*Animate]*Text

	OnAfterTextCreateCallback OnAfterCreateInterface[Text]
	OnAfterTextUpdateCallback OnAfterUpdateInterface[Text]
	OnAfterTextDeleteCallback OnAfterDeleteInterface[Text]
	OnAfterTextReadCallback   OnAfterReadInterface[Text]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here
}

func (stage *StageStruct) GetType() string {
	return "github.com/fullstack-lang/gongsvg/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitAnimate(animate *Animate)
	CheckoutAnimate(animate *Animate)
	CommitCircle(circle *Circle)
	CheckoutCircle(circle *Circle)
	CommitEllipse(ellipse *Ellipse)
	CheckoutEllipse(ellipse *Ellipse)
	CommitLayer(layer *Layer)
	CheckoutLayer(layer *Layer)
	CommitLine(line *Line)
	CheckoutLine(line *Line)
	CommitLink(link *Link)
	CheckoutLink(link *Link)
	CommitLinkAnchoredText(linkanchoredtext *LinkAnchoredText)
	CheckoutLinkAnchoredText(linkanchoredtext *LinkAnchoredText)
	CommitPath(path *Path)
	CheckoutPath(path *Path)
	CommitPoint(point *Point)
	CheckoutPoint(point *Point)
	CommitPolygone(polygone *Polygone)
	CheckoutPolygone(polygone *Polygone)
	CommitPolyline(polyline *Polyline)
	CheckoutPolyline(polyline *Polyline)
	CommitRect(rect *Rect)
	CheckoutRect(rect *Rect)
	CommitRectAnchoredPath(rectanchoredpath *RectAnchoredPath)
	CheckoutRectAnchoredPath(rectanchoredpath *RectAnchoredPath)
	CommitRectAnchoredRect(rectanchoredrect *RectAnchoredRect)
	CheckoutRectAnchoredRect(rectanchoredrect *RectAnchoredRect)
	CommitRectAnchoredText(rectanchoredtext *RectAnchoredText)
	CheckoutRectAnchoredText(rectanchoredtext *RectAnchoredText)
	CommitRectLinkLink(rectlinklink *RectLinkLink)
	CheckoutRectLinkLink(rectlinklink *RectLinkLink)
	CommitSVG(svg *SVG)
	CheckoutSVG(svg *SVG)
	CommitText(text *Text)
	CheckoutText(text *Text)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		Animates:           make(map[*Animate]any),
		Animates_mapString: make(map[string]*Animate),

		Circles:           make(map[*Circle]any),
		Circles_mapString: make(map[string]*Circle),

		Ellipses:           make(map[*Ellipse]any),
		Ellipses_mapString: make(map[string]*Ellipse),

		Layers:           make(map[*Layer]any),
		Layers_mapString: make(map[string]*Layer),

		Lines:           make(map[*Line]any),
		Lines_mapString: make(map[string]*Line),

		Links:           make(map[*Link]any),
		Links_mapString: make(map[string]*Link),

		LinkAnchoredTexts:           make(map[*LinkAnchoredText]any),
		LinkAnchoredTexts_mapString: make(map[string]*LinkAnchoredText),

		Paths:           make(map[*Path]any),
		Paths_mapString: make(map[string]*Path),

		Points:           make(map[*Point]any),
		Points_mapString: make(map[string]*Point),

		Polygones:           make(map[*Polygone]any),
		Polygones_mapString: make(map[string]*Polygone),

		Polylines:           make(map[*Polyline]any),
		Polylines_mapString: make(map[string]*Polyline),

		Rects:           make(map[*Rect]any),
		Rects_mapString: make(map[string]*Rect),

		RectAnchoredPaths:           make(map[*RectAnchoredPath]any),
		RectAnchoredPaths_mapString: make(map[string]*RectAnchoredPath),

		RectAnchoredRects:           make(map[*RectAnchoredRect]any),
		RectAnchoredRects_mapString: make(map[string]*RectAnchoredRect),

		RectAnchoredTexts:           make(map[*RectAnchoredText]any),
		RectAnchoredTexts_mapString: make(map[string]*RectAnchoredText),

		RectLinkLinks:           make(map[*RectLinkLink]any),
		RectLinkLinks_mapString: make(map[string]*RectLinkLink),

		SVGs:           make(map[*SVG]any),
		SVGs_mapString: make(map[string]*SVG),

		Texts:           make(map[*Text]any),
		Texts_mapString: make(map[string]*Text),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		path: path,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here
	}

	return
}

func (stage *StageStruct) GetPath() string {
	return stage.path
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
	stage.ComputeReverseMaps()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Animate"] = len(stage.Animates)
	stage.Map_GongStructName_InstancesNb["Circle"] = len(stage.Circles)
	stage.Map_GongStructName_InstancesNb["Ellipse"] = len(stage.Ellipses)
	stage.Map_GongStructName_InstancesNb["Layer"] = len(stage.Layers)
	stage.Map_GongStructName_InstancesNb["Line"] = len(stage.Lines)
	stage.Map_GongStructName_InstancesNb["Link"] = len(stage.Links)
	stage.Map_GongStructName_InstancesNb["LinkAnchoredText"] = len(stage.LinkAnchoredTexts)
	stage.Map_GongStructName_InstancesNb["Path"] = len(stage.Paths)
	stage.Map_GongStructName_InstancesNb["Point"] = len(stage.Points)
	stage.Map_GongStructName_InstancesNb["Polygone"] = len(stage.Polygones)
	stage.Map_GongStructName_InstancesNb["Polyline"] = len(stage.Polylines)
	stage.Map_GongStructName_InstancesNb["Rect"] = len(stage.Rects)
	stage.Map_GongStructName_InstancesNb["RectAnchoredPath"] = len(stage.RectAnchoredPaths)
	stage.Map_GongStructName_InstancesNb["RectAnchoredRect"] = len(stage.RectAnchoredRects)
	stage.Map_GongStructName_InstancesNb["RectAnchoredText"] = len(stage.RectAnchoredTexts)
	stage.Map_GongStructName_InstancesNb["RectLinkLink"] = len(stage.RectLinkLinks)
	stage.Map_GongStructName_InstancesNb["SVG"] = len(stage.SVGs)
	stage.Map_GongStructName_InstancesNb["Text"] = len(stage.Texts)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Animate"] = len(stage.Animates)
	stage.Map_GongStructName_InstancesNb["Circle"] = len(stage.Circles)
	stage.Map_GongStructName_InstancesNb["Ellipse"] = len(stage.Ellipses)
	stage.Map_GongStructName_InstancesNb["Layer"] = len(stage.Layers)
	stage.Map_GongStructName_InstancesNb["Line"] = len(stage.Lines)
	stage.Map_GongStructName_InstancesNb["Link"] = len(stage.Links)
	stage.Map_GongStructName_InstancesNb["LinkAnchoredText"] = len(stage.LinkAnchoredTexts)
	stage.Map_GongStructName_InstancesNb["Path"] = len(stage.Paths)
	stage.Map_GongStructName_InstancesNb["Point"] = len(stage.Points)
	stage.Map_GongStructName_InstancesNb["Polygone"] = len(stage.Polygones)
	stage.Map_GongStructName_InstancesNb["Polyline"] = len(stage.Polylines)
	stage.Map_GongStructName_InstancesNb["Rect"] = len(stage.Rects)
	stage.Map_GongStructName_InstancesNb["RectAnchoredPath"] = len(stage.RectAnchoredPaths)
	stage.Map_GongStructName_InstancesNb["RectAnchoredRect"] = len(stage.RectAnchoredRects)
	stage.Map_GongStructName_InstancesNb["RectAnchoredText"] = len(stage.RectAnchoredTexts)
	stage.Map_GongStructName_InstancesNb["RectLinkLink"] = len(stage.RectLinkLinks)
	stage.Map_GongStructName_InstancesNb["SVG"] = len(stage.SVGs)
	stage.Map_GongStructName_InstancesNb["Text"] = len(stage.Texts)

}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts animate to the model stage
func (animate *Animate) Stage(stage *StageStruct) *Animate {
	stage.Animates[animate] = __member
	stage.Animates_mapString[animate.Name] = animate

	return animate
}

// Unstage removes animate off the model stage
func (animate *Animate) Unstage(stage *StageStruct) *Animate {
	delete(stage.Animates, animate)
	delete(stage.Animates_mapString, animate.Name)
	return animate
}

// UnstageVoid removes animate off the model stage
func (animate *Animate) UnstageVoid(stage *StageStruct) {
	delete(stage.Animates, animate)
	delete(stage.Animates_mapString, animate.Name)
}

// commit animate to the back repo (if it is already staged)
func (animate *Animate) Commit(stage *StageStruct) *Animate {
	if _, ok := stage.Animates[animate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAnimate(animate)
		}
	}
	return animate
}

func (animate *Animate) CommitVoid(stage *StageStruct) {
	animate.Commit(stage)
}

// Checkout animate to the back repo (if it is already staged)
func (animate *Animate) Checkout(stage *StageStruct) *Animate {
	if _, ok := stage.Animates[animate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAnimate(animate)
		}
	}
	return animate
}

// for satisfaction of GongStruct interface
func (animate *Animate) GetName() (res string) {
	return animate.Name
}

// Stage puts circle to the model stage
func (circle *Circle) Stage(stage *StageStruct) *Circle {
	stage.Circles[circle] = __member
	stage.Circles_mapString[circle.Name] = circle

	return circle
}

// Unstage removes circle off the model stage
func (circle *Circle) Unstage(stage *StageStruct) *Circle {
	delete(stage.Circles, circle)
	delete(stage.Circles_mapString, circle.Name)
	return circle
}

// UnstageVoid removes circle off the model stage
func (circle *Circle) UnstageVoid(stage *StageStruct) {
	delete(stage.Circles, circle)
	delete(stage.Circles_mapString, circle.Name)
}

// commit circle to the back repo (if it is already staged)
func (circle *Circle) Commit(stage *StageStruct) *Circle {
	if _, ok := stage.Circles[circle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCircle(circle)
		}
	}
	return circle
}

func (circle *Circle) CommitVoid(stage *StageStruct) {
	circle.Commit(stage)
}

// Checkout circle to the back repo (if it is already staged)
func (circle *Circle) Checkout(stage *StageStruct) *Circle {
	if _, ok := stage.Circles[circle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCircle(circle)
		}
	}
	return circle
}

// for satisfaction of GongStruct interface
func (circle *Circle) GetName() (res string) {
	return circle.Name
}

// Stage puts ellipse to the model stage
func (ellipse *Ellipse) Stage(stage *StageStruct) *Ellipse {
	stage.Ellipses[ellipse] = __member
	stage.Ellipses_mapString[ellipse.Name] = ellipse

	return ellipse
}

// Unstage removes ellipse off the model stage
func (ellipse *Ellipse) Unstage(stage *StageStruct) *Ellipse {
	delete(stage.Ellipses, ellipse)
	delete(stage.Ellipses_mapString, ellipse.Name)
	return ellipse
}

// UnstageVoid removes ellipse off the model stage
func (ellipse *Ellipse) UnstageVoid(stage *StageStruct) {
	delete(stage.Ellipses, ellipse)
	delete(stage.Ellipses_mapString, ellipse.Name)
}

// commit ellipse to the back repo (if it is already staged)
func (ellipse *Ellipse) Commit(stage *StageStruct) *Ellipse {
	if _, ok := stage.Ellipses[ellipse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitEllipse(ellipse)
		}
	}
	return ellipse
}

func (ellipse *Ellipse) CommitVoid(stage *StageStruct) {
	ellipse.Commit(stage)
}

// Checkout ellipse to the back repo (if it is already staged)
func (ellipse *Ellipse) Checkout(stage *StageStruct) *Ellipse {
	if _, ok := stage.Ellipses[ellipse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutEllipse(ellipse)
		}
	}
	return ellipse
}

// for satisfaction of GongStruct interface
func (ellipse *Ellipse) GetName() (res string) {
	return ellipse.Name
}

// Stage puts layer to the model stage
func (layer *Layer) Stage(stage *StageStruct) *Layer {
	stage.Layers[layer] = __member
	stage.Layers_mapString[layer.Name] = layer

	return layer
}

// Unstage removes layer off the model stage
func (layer *Layer) Unstage(stage *StageStruct) *Layer {
	delete(stage.Layers, layer)
	delete(stage.Layers_mapString, layer.Name)
	return layer
}

// UnstageVoid removes layer off the model stage
func (layer *Layer) UnstageVoid(stage *StageStruct) {
	delete(stage.Layers, layer)
	delete(stage.Layers_mapString, layer.Name)
}

// commit layer to the back repo (if it is already staged)
func (layer *Layer) Commit(stage *StageStruct) *Layer {
	if _, ok := stage.Layers[layer]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLayer(layer)
		}
	}
	return layer
}

func (layer *Layer) CommitVoid(stage *StageStruct) {
	layer.Commit(stage)
}

// Checkout layer to the back repo (if it is already staged)
func (layer *Layer) Checkout(stage *StageStruct) *Layer {
	if _, ok := stage.Layers[layer]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLayer(layer)
		}
	}
	return layer
}

// for satisfaction of GongStruct interface
func (layer *Layer) GetName() (res string) {
	return layer.Name
}

// Stage puts line to the model stage
func (line *Line) Stage(stage *StageStruct) *Line {
	stage.Lines[line] = __member
	stage.Lines_mapString[line.Name] = line

	return line
}

// Unstage removes line off the model stage
func (line *Line) Unstage(stage *StageStruct) *Line {
	delete(stage.Lines, line)
	delete(stage.Lines_mapString, line.Name)
	return line
}

// UnstageVoid removes line off the model stage
func (line *Line) UnstageVoid(stage *StageStruct) {
	delete(stage.Lines, line)
	delete(stage.Lines_mapString, line.Name)
}

// commit line to the back repo (if it is already staged)
func (line *Line) Commit(stage *StageStruct) *Line {
	if _, ok := stage.Lines[line]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLine(line)
		}
	}
	return line
}

func (line *Line) CommitVoid(stage *StageStruct) {
	line.Commit(stage)
}

// Checkout line to the back repo (if it is already staged)
func (line *Line) Checkout(stage *StageStruct) *Line {
	if _, ok := stage.Lines[line]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLine(line)
		}
	}
	return line
}

// for satisfaction of GongStruct interface
func (line *Line) GetName() (res string) {
	return line.Name
}

// Stage puts link to the model stage
func (link *Link) Stage(stage *StageStruct) *Link {
	stage.Links[link] = __member
	stage.Links_mapString[link.Name] = link

	return link
}

// Unstage removes link off the model stage
func (link *Link) Unstage(stage *StageStruct) *Link {
	delete(stage.Links, link)
	delete(stage.Links_mapString, link.Name)
	return link
}

// UnstageVoid removes link off the model stage
func (link *Link) UnstageVoid(stage *StageStruct) {
	delete(stage.Links, link)
	delete(stage.Links_mapString, link.Name)
}

// commit link to the back repo (if it is already staged)
func (link *Link) Commit(stage *StageStruct) *Link {
	if _, ok := stage.Links[link]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLink(link)
		}
	}
	return link
}

func (link *Link) CommitVoid(stage *StageStruct) {
	link.Commit(stage)
}

// Checkout link to the back repo (if it is already staged)
func (link *Link) Checkout(stage *StageStruct) *Link {
	if _, ok := stage.Links[link]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLink(link)
		}
	}
	return link
}

// for satisfaction of GongStruct interface
func (link *Link) GetName() (res string) {
	return link.Name
}

// Stage puts linkanchoredtext to the model stage
func (linkanchoredtext *LinkAnchoredText) Stage(stage *StageStruct) *LinkAnchoredText {
	stage.LinkAnchoredTexts[linkanchoredtext] = __member
	stage.LinkAnchoredTexts_mapString[linkanchoredtext.Name] = linkanchoredtext

	return linkanchoredtext
}

// Unstage removes linkanchoredtext off the model stage
func (linkanchoredtext *LinkAnchoredText) Unstage(stage *StageStruct) *LinkAnchoredText {
	delete(stage.LinkAnchoredTexts, linkanchoredtext)
	delete(stage.LinkAnchoredTexts_mapString, linkanchoredtext.Name)
	return linkanchoredtext
}

// UnstageVoid removes linkanchoredtext off the model stage
func (linkanchoredtext *LinkAnchoredText) UnstageVoid(stage *StageStruct) {
	delete(stage.LinkAnchoredTexts, linkanchoredtext)
	delete(stage.LinkAnchoredTexts_mapString, linkanchoredtext.Name)
}

// commit linkanchoredtext to the back repo (if it is already staged)
func (linkanchoredtext *LinkAnchoredText) Commit(stage *StageStruct) *LinkAnchoredText {
	if _, ok := stage.LinkAnchoredTexts[linkanchoredtext]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLinkAnchoredText(linkanchoredtext)
		}
	}
	return linkanchoredtext
}

func (linkanchoredtext *LinkAnchoredText) CommitVoid(stage *StageStruct) {
	linkanchoredtext.Commit(stage)
}

// Checkout linkanchoredtext to the back repo (if it is already staged)
func (linkanchoredtext *LinkAnchoredText) Checkout(stage *StageStruct) *LinkAnchoredText {
	if _, ok := stage.LinkAnchoredTexts[linkanchoredtext]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLinkAnchoredText(linkanchoredtext)
		}
	}
	return linkanchoredtext
}

// for satisfaction of GongStruct interface
func (linkanchoredtext *LinkAnchoredText) GetName() (res string) {
	return linkanchoredtext.Name
}

// Stage puts path to the model stage
func (path *Path) Stage(stage *StageStruct) *Path {
	stage.Paths[path] = __member
	stage.Paths_mapString[path.Name] = path

	return path
}

// Unstage removes path off the model stage
func (path *Path) Unstage(stage *StageStruct) *Path {
	delete(stage.Paths, path)
	delete(stage.Paths_mapString, path.Name)
	return path
}

// UnstageVoid removes path off the model stage
func (path *Path) UnstageVoid(stage *StageStruct) {
	delete(stage.Paths, path)
	delete(stage.Paths_mapString, path.Name)
}

// commit path to the back repo (if it is already staged)
func (path *Path) Commit(stage *StageStruct) *Path {
	if _, ok := stage.Paths[path]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPath(path)
		}
	}
	return path
}

func (path *Path) CommitVoid(stage *StageStruct) {
	path.Commit(stage)
}

// Checkout path to the back repo (if it is already staged)
func (path *Path) Checkout(stage *StageStruct) *Path {
	if _, ok := stage.Paths[path]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPath(path)
		}
	}
	return path
}

// for satisfaction of GongStruct interface
func (path *Path) GetName() (res string) {
	return path.Name
}

// Stage puts point to the model stage
func (point *Point) Stage(stage *StageStruct) *Point {
	stage.Points[point] = __member
	stage.Points_mapString[point.Name] = point

	return point
}

// Unstage removes point off the model stage
func (point *Point) Unstage(stage *StageStruct) *Point {
	delete(stage.Points, point)
	delete(stage.Points_mapString, point.Name)
	return point
}

// UnstageVoid removes point off the model stage
func (point *Point) UnstageVoid(stage *StageStruct) {
	delete(stage.Points, point)
	delete(stage.Points_mapString, point.Name)
}

// commit point to the back repo (if it is already staged)
func (point *Point) Commit(stage *StageStruct) *Point {
	if _, ok := stage.Points[point]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPoint(point)
		}
	}
	return point
}

func (point *Point) CommitVoid(stage *StageStruct) {
	point.Commit(stage)
}

// Checkout point to the back repo (if it is already staged)
func (point *Point) Checkout(stage *StageStruct) *Point {
	if _, ok := stage.Points[point]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPoint(point)
		}
	}
	return point
}

// for satisfaction of GongStruct interface
func (point *Point) GetName() (res string) {
	return point.Name
}

// Stage puts polygone to the model stage
func (polygone *Polygone) Stage(stage *StageStruct) *Polygone {
	stage.Polygones[polygone] = __member
	stage.Polygones_mapString[polygone.Name] = polygone

	return polygone
}

// Unstage removes polygone off the model stage
func (polygone *Polygone) Unstage(stage *StageStruct) *Polygone {
	delete(stage.Polygones, polygone)
	delete(stage.Polygones_mapString, polygone.Name)
	return polygone
}

// UnstageVoid removes polygone off the model stage
func (polygone *Polygone) UnstageVoid(stage *StageStruct) {
	delete(stage.Polygones, polygone)
	delete(stage.Polygones_mapString, polygone.Name)
}

// commit polygone to the back repo (if it is already staged)
func (polygone *Polygone) Commit(stage *StageStruct) *Polygone {
	if _, ok := stage.Polygones[polygone]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPolygone(polygone)
		}
	}
	return polygone
}

func (polygone *Polygone) CommitVoid(stage *StageStruct) {
	polygone.Commit(stage)
}

// Checkout polygone to the back repo (if it is already staged)
func (polygone *Polygone) Checkout(stage *StageStruct) *Polygone {
	if _, ok := stage.Polygones[polygone]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPolygone(polygone)
		}
	}
	return polygone
}

// for satisfaction of GongStruct interface
func (polygone *Polygone) GetName() (res string) {
	return polygone.Name
}

// Stage puts polyline to the model stage
func (polyline *Polyline) Stage(stage *StageStruct) *Polyline {
	stage.Polylines[polyline] = __member
	stage.Polylines_mapString[polyline.Name] = polyline

	return polyline
}

// Unstage removes polyline off the model stage
func (polyline *Polyline) Unstage(stage *StageStruct) *Polyline {
	delete(stage.Polylines, polyline)
	delete(stage.Polylines_mapString, polyline.Name)
	return polyline
}

// UnstageVoid removes polyline off the model stage
func (polyline *Polyline) UnstageVoid(stage *StageStruct) {
	delete(stage.Polylines, polyline)
	delete(stage.Polylines_mapString, polyline.Name)
}

// commit polyline to the back repo (if it is already staged)
func (polyline *Polyline) Commit(stage *StageStruct) *Polyline {
	if _, ok := stage.Polylines[polyline]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPolyline(polyline)
		}
	}
	return polyline
}

func (polyline *Polyline) CommitVoid(stage *StageStruct) {
	polyline.Commit(stage)
}

// Checkout polyline to the back repo (if it is already staged)
func (polyline *Polyline) Checkout(stage *StageStruct) *Polyline {
	if _, ok := stage.Polylines[polyline]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPolyline(polyline)
		}
	}
	return polyline
}

// for satisfaction of GongStruct interface
func (polyline *Polyline) GetName() (res string) {
	return polyline.Name
}

// Stage puts rect to the model stage
func (rect *Rect) Stage(stage *StageStruct) *Rect {
	stage.Rects[rect] = __member
	stage.Rects_mapString[rect.Name] = rect

	return rect
}

// Unstage removes rect off the model stage
func (rect *Rect) Unstage(stage *StageStruct) *Rect {
	delete(stage.Rects, rect)
	delete(stage.Rects_mapString, rect.Name)
	return rect
}

// UnstageVoid removes rect off the model stage
func (rect *Rect) UnstageVoid(stage *StageStruct) {
	delete(stage.Rects, rect)
	delete(stage.Rects_mapString, rect.Name)
}

// commit rect to the back repo (if it is already staged)
func (rect *Rect) Commit(stage *StageStruct) *Rect {
	if _, ok := stage.Rects[rect]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRect(rect)
		}
	}
	return rect
}

func (rect *Rect) CommitVoid(stage *StageStruct) {
	rect.Commit(stage)
}

// Checkout rect to the back repo (if it is already staged)
func (rect *Rect) Checkout(stage *StageStruct) *Rect {
	if _, ok := stage.Rects[rect]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRect(rect)
		}
	}
	return rect
}

// for satisfaction of GongStruct interface
func (rect *Rect) GetName() (res string) {
	return rect.Name
}

// Stage puts rectanchoredpath to the model stage
func (rectanchoredpath *RectAnchoredPath) Stage(stage *StageStruct) *RectAnchoredPath {
	stage.RectAnchoredPaths[rectanchoredpath] = __member
	stage.RectAnchoredPaths_mapString[rectanchoredpath.Name] = rectanchoredpath

	return rectanchoredpath
}

// Unstage removes rectanchoredpath off the model stage
func (rectanchoredpath *RectAnchoredPath) Unstage(stage *StageStruct) *RectAnchoredPath {
	delete(stage.RectAnchoredPaths, rectanchoredpath)
	delete(stage.RectAnchoredPaths_mapString, rectanchoredpath.Name)
	return rectanchoredpath
}

// UnstageVoid removes rectanchoredpath off the model stage
func (rectanchoredpath *RectAnchoredPath) UnstageVoid(stage *StageStruct) {
	delete(stage.RectAnchoredPaths, rectanchoredpath)
	delete(stage.RectAnchoredPaths_mapString, rectanchoredpath.Name)
}

// commit rectanchoredpath to the back repo (if it is already staged)
func (rectanchoredpath *RectAnchoredPath) Commit(stage *StageStruct) *RectAnchoredPath {
	if _, ok := stage.RectAnchoredPaths[rectanchoredpath]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRectAnchoredPath(rectanchoredpath)
		}
	}
	return rectanchoredpath
}

func (rectanchoredpath *RectAnchoredPath) CommitVoid(stage *StageStruct) {
	rectanchoredpath.Commit(stage)
}

// Checkout rectanchoredpath to the back repo (if it is already staged)
func (rectanchoredpath *RectAnchoredPath) Checkout(stage *StageStruct) *RectAnchoredPath {
	if _, ok := stage.RectAnchoredPaths[rectanchoredpath]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRectAnchoredPath(rectanchoredpath)
		}
	}
	return rectanchoredpath
}

// for satisfaction of GongStruct interface
func (rectanchoredpath *RectAnchoredPath) GetName() (res string) {
	return rectanchoredpath.Name
}

// Stage puts rectanchoredrect to the model stage
func (rectanchoredrect *RectAnchoredRect) Stage(stage *StageStruct) *RectAnchoredRect {
	stage.RectAnchoredRects[rectanchoredrect] = __member
	stage.RectAnchoredRects_mapString[rectanchoredrect.Name] = rectanchoredrect

	return rectanchoredrect
}

// Unstage removes rectanchoredrect off the model stage
func (rectanchoredrect *RectAnchoredRect) Unstage(stage *StageStruct) *RectAnchoredRect {
	delete(stage.RectAnchoredRects, rectanchoredrect)
	delete(stage.RectAnchoredRects_mapString, rectanchoredrect.Name)
	return rectanchoredrect
}

// UnstageVoid removes rectanchoredrect off the model stage
func (rectanchoredrect *RectAnchoredRect) UnstageVoid(stage *StageStruct) {
	delete(stage.RectAnchoredRects, rectanchoredrect)
	delete(stage.RectAnchoredRects_mapString, rectanchoredrect.Name)
}

// commit rectanchoredrect to the back repo (if it is already staged)
func (rectanchoredrect *RectAnchoredRect) Commit(stage *StageStruct) *RectAnchoredRect {
	if _, ok := stage.RectAnchoredRects[rectanchoredrect]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRectAnchoredRect(rectanchoredrect)
		}
	}
	return rectanchoredrect
}

func (rectanchoredrect *RectAnchoredRect) CommitVoid(stage *StageStruct) {
	rectanchoredrect.Commit(stage)
}

// Checkout rectanchoredrect to the back repo (if it is already staged)
func (rectanchoredrect *RectAnchoredRect) Checkout(stage *StageStruct) *RectAnchoredRect {
	if _, ok := stage.RectAnchoredRects[rectanchoredrect]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRectAnchoredRect(rectanchoredrect)
		}
	}
	return rectanchoredrect
}

// for satisfaction of GongStruct interface
func (rectanchoredrect *RectAnchoredRect) GetName() (res string) {
	return rectanchoredrect.Name
}

// Stage puts rectanchoredtext to the model stage
func (rectanchoredtext *RectAnchoredText) Stage(stage *StageStruct) *RectAnchoredText {
	stage.RectAnchoredTexts[rectanchoredtext] = __member
	stage.RectAnchoredTexts_mapString[rectanchoredtext.Name] = rectanchoredtext

	return rectanchoredtext
}

// Unstage removes rectanchoredtext off the model stage
func (rectanchoredtext *RectAnchoredText) Unstage(stage *StageStruct) *RectAnchoredText {
	delete(stage.RectAnchoredTexts, rectanchoredtext)
	delete(stage.RectAnchoredTexts_mapString, rectanchoredtext.Name)
	return rectanchoredtext
}

// UnstageVoid removes rectanchoredtext off the model stage
func (rectanchoredtext *RectAnchoredText) UnstageVoid(stage *StageStruct) {
	delete(stage.RectAnchoredTexts, rectanchoredtext)
	delete(stage.RectAnchoredTexts_mapString, rectanchoredtext.Name)
}

// commit rectanchoredtext to the back repo (if it is already staged)
func (rectanchoredtext *RectAnchoredText) Commit(stage *StageStruct) *RectAnchoredText {
	if _, ok := stage.RectAnchoredTexts[rectanchoredtext]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRectAnchoredText(rectanchoredtext)
		}
	}
	return rectanchoredtext
}

func (rectanchoredtext *RectAnchoredText) CommitVoid(stage *StageStruct) {
	rectanchoredtext.Commit(stage)
}

// Checkout rectanchoredtext to the back repo (if it is already staged)
func (rectanchoredtext *RectAnchoredText) Checkout(stage *StageStruct) *RectAnchoredText {
	if _, ok := stage.RectAnchoredTexts[rectanchoredtext]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRectAnchoredText(rectanchoredtext)
		}
	}
	return rectanchoredtext
}

// for satisfaction of GongStruct interface
func (rectanchoredtext *RectAnchoredText) GetName() (res string) {
	return rectanchoredtext.Name
}

// Stage puts rectlinklink to the model stage
func (rectlinklink *RectLinkLink) Stage(stage *StageStruct) *RectLinkLink {
	stage.RectLinkLinks[rectlinklink] = __member
	stage.RectLinkLinks_mapString[rectlinklink.Name] = rectlinklink

	return rectlinklink
}

// Unstage removes rectlinklink off the model stage
func (rectlinklink *RectLinkLink) Unstage(stage *StageStruct) *RectLinkLink {
	delete(stage.RectLinkLinks, rectlinklink)
	delete(stage.RectLinkLinks_mapString, rectlinklink.Name)
	return rectlinklink
}

// UnstageVoid removes rectlinklink off the model stage
func (rectlinklink *RectLinkLink) UnstageVoid(stage *StageStruct) {
	delete(stage.RectLinkLinks, rectlinklink)
	delete(stage.RectLinkLinks_mapString, rectlinklink.Name)
}

// commit rectlinklink to the back repo (if it is already staged)
func (rectlinklink *RectLinkLink) Commit(stage *StageStruct) *RectLinkLink {
	if _, ok := stage.RectLinkLinks[rectlinklink]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRectLinkLink(rectlinklink)
		}
	}
	return rectlinklink
}

func (rectlinklink *RectLinkLink) CommitVoid(stage *StageStruct) {
	rectlinklink.Commit(stage)
}

// Checkout rectlinklink to the back repo (if it is already staged)
func (rectlinklink *RectLinkLink) Checkout(stage *StageStruct) *RectLinkLink {
	if _, ok := stage.RectLinkLinks[rectlinklink]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRectLinkLink(rectlinklink)
		}
	}
	return rectlinklink
}

// for satisfaction of GongStruct interface
func (rectlinklink *RectLinkLink) GetName() (res string) {
	return rectlinklink.Name
}

// Stage puts svg to the model stage
func (svg *SVG) Stage(stage *StageStruct) *SVG {
	stage.SVGs[svg] = __member
	stage.SVGs_mapString[svg.Name] = svg

	return svg
}

// Unstage removes svg off the model stage
func (svg *SVG) Unstage(stage *StageStruct) *SVG {
	delete(stage.SVGs, svg)
	delete(stage.SVGs_mapString, svg.Name)
	return svg
}

// UnstageVoid removes svg off the model stage
func (svg *SVG) UnstageVoid(stage *StageStruct) {
	delete(stage.SVGs, svg)
	delete(stage.SVGs_mapString, svg.Name)
}

// commit svg to the back repo (if it is already staged)
func (svg *SVG) Commit(stage *StageStruct) *SVG {
	if _, ok := stage.SVGs[svg]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSVG(svg)
		}
	}
	return svg
}

func (svg *SVG) CommitVoid(stage *StageStruct) {
	svg.Commit(stage)
}

// Checkout svg to the back repo (if it is already staged)
func (svg *SVG) Checkout(stage *StageStruct) *SVG {
	if _, ok := stage.SVGs[svg]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSVG(svg)
		}
	}
	return svg
}

// for satisfaction of GongStruct interface
func (svg *SVG) GetName() (res string) {
	return svg.Name
}

// Stage puts text to the model stage
func (text *Text) Stage(stage *StageStruct) *Text {
	stage.Texts[text] = __member
	stage.Texts_mapString[text.Name] = text

	return text
}

// Unstage removes text off the model stage
func (text *Text) Unstage(stage *StageStruct) *Text {
	delete(stage.Texts, text)
	delete(stage.Texts_mapString, text.Name)
	return text
}

// UnstageVoid removes text off the model stage
func (text *Text) UnstageVoid(stage *StageStruct) {
	delete(stage.Texts, text)
	delete(stage.Texts_mapString, text.Name)
}

// commit text to the back repo (if it is already staged)
func (text *Text) Commit(stage *StageStruct) *Text {
	if _, ok := stage.Texts[text]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitText(text)
		}
	}
	return text
}

func (text *Text) CommitVoid(stage *StageStruct) {
	text.Commit(stage)
}

// Checkout text to the back repo (if it is already staged)
func (text *Text) Checkout(stage *StageStruct) *Text {
	if _, ok := stage.Texts[text]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutText(text)
		}
	}
	return text
}

// for satisfaction of GongStruct interface
func (text *Text) GetName() (res string) {
	return text.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMAnimate(Animate *Animate)
	CreateORMCircle(Circle *Circle)
	CreateORMEllipse(Ellipse *Ellipse)
	CreateORMLayer(Layer *Layer)
	CreateORMLine(Line *Line)
	CreateORMLink(Link *Link)
	CreateORMLinkAnchoredText(LinkAnchoredText *LinkAnchoredText)
	CreateORMPath(Path *Path)
	CreateORMPoint(Point *Point)
	CreateORMPolygone(Polygone *Polygone)
	CreateORMPolyline(Polyline *Polyline)
	CreateORMRect(Rect *Rect)
	CreateORMRectAnchoredPath(RectAnchoredPath *RectAnchoredPath)
	CreateORMRectAnchoredRect(RectAnchoredRect *RectAnchoredRect)
	CreateORMRectAnchoredText(RectAnchoredText *RectAnchoredText)
	CreateORMRectLinkLink(RectLinkLink *RectLinkLink)
	CreateORMSVG(SVG *SVG)
	CreateORMText(Text *Text)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAnimate(Animate *Animate)
	DeleteORMCircle(Circle *Circle)
	DeleteORMEllipse(Ellipse *Ellipse)
	DeleteORMLayer(Layer *Layer)
	DeleteORMLine(Line *Line)
	DeleteORMLink(Link *Link)
	DeleteORMLinkAnchoredText(LinkAnchoredText *LinkAnchoredText)
	DeleteORMPath(Path *Path)
	DeleteORMPoint(Point *Point)
	DeleteORMPolygone(Polygone *Polygone)
	DeleteORMPolyline(Polyline *Polyline)
	DeleteORMRect(Rect *Rect)
	DeleteORMRectAnchoredPath(RectAnchoredPath *RectAnchoredPath)
	DeleteORMRectAnchoredRect(RectAnchoredRect *RectAnchoredRect)
	DeleteORMRectAnchoredText(RectAnchoredText *RectAnchoredText)
	DeleteORMRectLinkLink(RectLinkLink *RectLinkLink)
	DeleteORMSVG(SVG *SVG)
	DeleteORMText(Text *Text)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Animates = make(map[*Animate]any)
	stage.Animates_mapString = make(map[string]*Animate)

	stage.Circles = make(map[*Circle]any)
	stage.Circles_mapString = make(map[string]*Circle)

	stage.Ellipses = make(map[*Ellipse]any)
	stage.Ellipses_mapString = make(map[string]*Ellipse)

	stage.Layers = make(map[*Layer]any)
	stage.Layers_mapString = make(map[string]*Layer)

	stage.Lines = make(map[*Line]any)
	stage.Lines_mapString = make(map[string]*Line)

	stage.Links = make(map[*Link]any)
	stage.Links_mapString = make(map[string]*Link)

	stage.LinkAnchoredTexts = make(map[*LinkAnchoredText]any)
	stage.LinkAnchoredTexts_mapString = make(map[string]*LinkAnchoredText)

	stage.Paths = make(map[*Path]any)
	stage.Paths_mapString = make(map[string]*Path)

	stage.Points = make(map[*Point]any)
	stage.Points_mapString = make(map[string]*Point)

	stage.Polygones = make(map[*Polygone]any)
	stage.Polygones_mapString = make(map[string]*Polygone)

	stage.Polylines = make(map[*Polyline]any)
	stage.Polylines_mapString = make(map[string]*Polyline)

	stage.Rects = make(map[*Rect]any)
	stage.Rects_mapString = make(map[string]*Rect)

	stage.RectAnchoredPaths = make(map[*RectAnchoredPath]any)
	stage.RectAnchoredPaths_mapString = make(map[string]*RectAnchoredPath)

	stage.RectAnchoredRects = make(map[*RectAnchoredRect]any)
	stage.RectAnchoredRects_mapString = make(map[string]*RectAnchoredRect)

	stage.RectAnchoredTexts = make(map[*RectAnchoredText]any)
	stage.RectAnchoredTexts_mapString = make(map[string]*RectAnchoredText)

	stage.RectLinkLinks = make(map[*RectLinkLink]any)
	stage.RectLinkLinks_mapString = make(map[string]*RectLinkLink)

	stage.SVGs = make(map[*SVG]any)
	stage.SVGs_mapString = make(map[string]*SVG)

	stage.Texts = make(map[*Text]any)
	stage.Texts_mapString = make(map[string]*Text)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.Animates = nil
	stage.Animates_mapString = nil

	stage.Circles = nil
	stage.Circles_mapString = nil

	stage.Ellipses = nil
	stage.Ellipses_mapString = nil

	stage.Layers = nil
	stage.Layers_mapString = nil

	stage.Lines = nil
	stage.Lines_mapString = nil

	stage.Links = nil
	stage.Links_mapString = nil

	stage.LinkAnchoredTexts = nil
	stage.LinkAnchoredTexts_mapString = nil

	stage.Paths = nil
	stage.Paths_mapString = nil

	stage.Points = nil
	stage.Points_mapString = nil

	stage.Polygones = nil
	stage.Polygones_mapString = nil

	stage.Polylines = nil
	stage.Polylines_mapString = nil

	stage.Rects = nil
	stage.Rects_mapString = nil

	stage.RectAnchoredPaths = nil
	stage.RectAnchoredPaths_mapString = nil

	stage.RectAnchoredRects = nil
	stage.RectAnchoredRects_mapString = nil

	stage.RectAnchoredTexts = nil
	stage.RectAnchoredTexts_mapString = nil

	stage.RectLinkLinks = nil
	stage.RectLinkLinks_mapString = nil

	stage.SVGs = nil
	stage.SVGs_mapString = nil

	stage.Texts = nil
	stage.Texts_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for animate := range stage.Animates {
		animate.Unstage(stage)
	}

	for circle := range stage.Circles {
		circle.Unstage(stage)
	}

	for ellipse := range stage.Ellipses {
		ellipse.Unstage(stage)
	}

	for layer := range stage.Layers {
		layer.Unstage(stage)
	}

	for line := range stage.Lines {
		line.Unstage(stage)
	}

	for link := range stage.Links {
		link.Unstage(stage)
	}

	for linkanchoredtext := range stage.LinkAnchoredTexts {
		linkanchoredtext.Unstage(stage)
	}

	for path := range stage.Paths {
		path.Unstage(stage)
	}

	for point := range stage.Points {
		point.Unstage(stage)
	}

	for polygone := range stage.Polygones {
		polygone.Unstage(stage)
	}

	for polyline := range stage.Polylines {
		polyline.Unstage(stage)
	}

	for rect := range stage.Rects {
		rect.Unstage(stage)
	}

	for rectanchoredpath := range stage.RectAnchoredPaths {
		rectanchoredpath.Unstage(stage)
	}

	for rectanchoredrect := range stage.RectAnchoredRects {
		rectanchoredrect.Unstage(stage)
	}

	for rectanchoredtext := range stage.RectAnchoredTexts {
		rectanchoredtext.Unstage(stage)
	}

	for rectlinklink := range stage.RectLinkLinks {
		rectlinklink.Unstage(stage)
	}

	for svg := range stage.SVGs {
		svg.Unstage(stage)
	}

	for text := range stage.Texts {
		text.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	GetName() string
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

	for key := range set {
		sortedSlice = append(sortedSlice, key)
	}
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *StageStruct) (sortedSlice []T) {

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
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*Animate]any:
		return any(&stage.Animates).(*Type)
	case map[*Circle]any:
		return any(&stage.Circles).(*Type)
	case map[*Ellipse]any:
		return any(&stage.Ellipses).(*Type)
	case map[*Layer]any:
		return any(&stage.Layers).(*Type)
	case map[*Line]any:
		return any(&stage.Lines).(*Type)
	case map[*Link]any:
		return any(&stage.Links).(*Type)
	case map[*LinkAnchoredText]any:
		return any(&stage.LinkAnchoredTexts).(*Type)
	case map[*Path]any:
		return any(&stage.Paths).(*Type)
	case map[*Point]any:
		return any(&stage.Points).(*Type)
	case map[*Polygone]any:
		return any(&stage.Polygones).(*Type)
	case map[*Polyline]any:
		return any(&stage.Polylines).(*Type)
	case map[*Rect]any:
		return any(&stage.Rects).(*Type)
	case map[*RectAnchoredPath]any:
		return any(&stage.RectAnchoredPaths).(*Type)
	case map[*RectAnchoredRect]any:
		return any(&stage.RectAnchoredRects).(*Type)
	case map[*RectAnchoredText]any:
		return any(&stage.RectAnchoredTexts).(*Type)
	case map[*RectLinkLink]any:
		return any(&stage.RectLinkLinks).(*Type)
	case map[*SVG]any:
		return any(&stage.SVGs).(*Type)
	case map[*Text]any:
		return any(&stage.Texts).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*Animate:
		return any(&stage.Animates_mapString).(*Type)
	case map[string]*Circle:
		return any(&stage.Circles_mapString).(*Type)
	case map[string]*Ellipse:
		return any(&stage.Ellipses_mapString).(*Type)
	case map[string]*Layer:
		return any(&stage.Layers_mapString).(*Type)
	case map[string]*Line:
		return any(&stage.Lines_mapString).(*Type)
	case map[string]*Link:
		return any(&stage.Links_mapString).(*Type)
	case map[string]*LinkAnchoredText:
		return any(&stage.LinkAnchoredTexts_mapString).(*Type)
	case map[string]*Path:
		return any(&stage.Paths_mapString).(*Type)
	case map[string]*Point:
		return any(&stage.Points_mapString).(*Type)
	case map[string]*Polygone:
		return any(&stage.Polygones_mapString).(*Type)
	case map[string]*Polyline:
		return any(&stage.Polylines_mapString).(*Type)
	case map[string]*Rect:
		return any(&stage.Rects_mapString).(*Type)
	case map[string]*RectAnchoredPath:
		return any(&stage.RectAnchoredPaths_mapString).(*Type)
	case map[string]*RectAnchoredRect:
		return any(&stage.RectAnchoredRects_mapString).(*Type)
	case map[string]*RectAnchoredText:
		return any(&stage.RectAnchoredTexts_mapString).(*Type)
	case map[string]*RectLinkLink:
		return any(&stage.RectLinkLinks_mapString).(*Type)
	case map[string]*SVG:
		return any(&stage.SVGs_mapString).(*Type)
	case map[string]*Text:
		return any(&stage.Texts_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Animate:
		return any(&stage.Animates).(*map[*Type]any)
	case Circle:
		return any(&stage.Circles).(*map[*Type]any)
	case Ellipse:
		return any(&stage.Ellipses).(*map[*Type]any)
	case Layer:
		return any(&stage.Layers).(*map[*Type]any)
	case Line:
		return any(&stage.Lines).(*map[*Type]any)
	case Link:
		return any(&stage.Links).(*map[*Type]any)
	case LinkAnchoredText:
		return any(&stage.LinkAnchoredTexts).(*map[*Type]any)
	case Path:
		return any(&stage.Paths).(*map[*Type]any)
	case Point:
		return any(&stage.Points).(*map[*Type]any)
	case Polygone:
		return any(&stage.Polygones).(*map[*Type]any)
	case Polyline:
		return any(&stage.Polylines).(*map[*Type]any)
	case Rect:
		return any(&stage.Rects).(*map[*Type]any)
	case RectAnchoredPath:
		return any(&stage.RectAnchoredPaths).(*map[*Type]any)
	case RectAnchoredRect:
		return any(&stage.RectAnchoredRects).(*map[*Type]any)
	case RectAnchoredText:
		return any(&stage.RectAnchoredTexts).(*map[*Type]any)
	case RectLinkLink:
		return any(&stage.RectLinkLinks).(*map[*Type]any)
	case SVG:
		return any(&stage.SVGs).(*map[*Type]any)
	case Text:
		return any(&stage.Texts).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *StageStruct) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Animate:
		return any(&stage.Animates).(*map[Type]any)
	case *Circle:
		return any(&stage.Circles).(*map[Type]any)
	case *Ellipse:
		return any(&stage.Ellipses).(*map[Type]any)
	case *Layer:
		return any(&stage.Layers).(*map[Type]any)
	case *Line:
		return any(&stage.Lines).(*map[Type]any)
	case *Link:
		return any(&stage.Links).(*map[Type]any)
	case *LinkAnchoredText:
		return any(&stage.LinkAnchoredTexts).(*map[Type]any)
	case *Path:
		return any(&stage.Paths).(*map[Type]any)
	case *Point:
		return any(&stage.Points).(*map[Type]any)
	case *Polygone:
		return any(&stage.Polygones).(*map[Type]any)
	case *Polyline:
		return any(&stage.Polylines).(*map[Type]any)
	case *Rect:
		return any(&stage.Rects).(*map[Type]any)
	case *RectAnchoredPath:
		return any(&stage.RectAnchoredPaths).(*map[Type]any)
	case *RectAnchoredRect:
		return any(&stage.RectAnchoredRects).(*map[Type]any)
	case *RectAnchoredText:
		return any(&stage.RectAnchoredTexts).(*map[Type]any)
	case *RectLinkLink:
		return any(&stage.RectLinkLinks).(*map[Type]any)
	case *SVG:
		return any(&stage.SVGs).(*map[Type]any)
	case *Text:
		return any(&stage.Texts).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Animate:
		return any(&stage.Animates_mapString).(*map[string]*Type)
	case Circle:
		return any(&stage.Circles_mapString).(*map[string]*Type)
	case Ellipse:
		return any(&stage.Ellipses_mapString).(*map[string]*Type)
	case Layer:
		return any(&stage.Layers_mapString).(*map[string]*Type)
	case Line:
		return any(&stage.Lines_mapString).(*map[string]*Type)
	case Link:
		return any(&stage.Links_mapString).(*map[string]*Type)
	case LinkAnchoredText:
		return any(&stage.LinkAnchoredTexts_mapString).(*map[string]*Type)
	case Path:
		return any(&stage.Paths_mapString).(*map[string]*Type)
	case Point:
		return any(&stage.Points_mapString).(*map[string]*Type)
	case Polygone:
		return any(&stage.Polygones_mapString).(*map[string]*Type)
	case Polyline:
		return any(&stage.Polylines_mapString).(*map[string]*Type)
	case Rect:
		return any(&stage.Rects_mapString).(*map[string]*Type)
	case RectAnchoredPath:
		return any(&stage.RectAnchoredPaths_mapString).(*map[string]*Type)
	case RectAnchoredRect:
		return any(&stage.RectAnchoredRects_mapString).(*map[string]*Type)
	case RectAnchoredText:
		return any(&stage.RectAnchoredTexts_mapString).(*map[string]*Type)
	case RectLinkLink:
		return any(&stage.RectLinkLinks_mapString).(*map[string]*Type)
	case SVG:
		return any(&stage.SVGs_mapString).(*map[string]*Type)
	case Text:
		return any(&stage.Texts_mapString).(*map[string]*Type)
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
	case Animate:
		return any(&Animate{
			// Initialisation of associations
		}).(*Type)
	case Circle:
		return any(&Circle{
			// Initialisation of associations
			// field is initialized with an instance of Animate with the name of the field
			Animations: []*Animate{{Name: "Animations"}},
		}).(*Type)
	case Ellipse:
		return any(&Ellipse{
			// Initialisation of associations
			// field is initialized with an instance of Animate with the name of the field
			Animates: []*Animate{{Name: "Animates"}},
		}).(*Type)
	case Layer:
		return any(&Layer{
			// Initialisation of associations
			// field is initialized with an instance of Rect with the name of the field
			Rects: []*Rect{{Name: "Rects"}},
			// field is initialized with an instance of Text with the name of the field
			Texts: []*Text{{Name: "Texts"}},
			// field is initialized with an instance of Circle with the name of the field
			Circles: []*Circle{{Name: "Circles"}},
			// field is initialized with an instance of Line with the name of the field
			Lines: []*Line{{Name: "Lines"}},
			// field is initialized with an instance of Ellipse with the name of the field
			Ellipses: []*Ellipse{{Name: "Ellipses"}},
			// field is initialized with an instance of Polyline with the name of the field
			Polylines: []*Polyline{{Name: "Polylines"}},
			// field is initialized with an instance of Polygone with the name of the field
			Polygones: []*Polygone{{Name: "Polygones"}},
			// field is initialized with an instance of Path with the name of the field
			Paths: []*Path{{Name: "Paths"}},
			// field is initialized with an instance of Link with the name of the field
			Links: []*Link{{Name: "Links"}},
			// field is initialized with an instance of RectLinkLink with the name of the field
			RectLinkLinks: []*RectLinkLink{{Name: "RectLinkLinks"}},
		}).(*Type)
	case Line:
		return any(&Line{
			// Initialisation of associations
			// field is initialized with an instance of Animate with the name of the field
			Animates: []*Animate{{Name: "Animates"}},
		}).(*Type)
	case Link:
		return any(&Link{
			// Initialisation of associations
			// field is initialized with an instance of Rect with the name of the field
			Start: &Rect{Name: "Start"},
			// field is initialized with an instance of Rect with the name of the field
			End: &Rect{Name: "End"},
			// field is initialized with an instance of LinkAnchoredText with the name of the field
			TextAtArrowEnd: []*LinkAnchoredText{{Name: "TextAtArrowEnd"}},
			// field is initialized with an instance of LinkAnchoredText with the name of the field
			TextAtArrowStart: []*LinkAnchoredText{{Name: "TextAtArrowStart"}},
			// field is initialized with an instance of Point with the name of the field
			ControlPoints: []*Point{{Name: "ControlPoints"}},
		}).(*Type)
	case LinkAnchoredText:
		return any(&LinkAnchoredText{
			// Initialisation of associations
			// field is initialized with an instance of Animate with the name of the field
			Animates: []*Animate{{Name: "Animates"}},
		}).(*Type)
	case Path:
		return any(&Path{
			// Initialisation of associations
			// field is initialized with an instance of Animate with the name of the field
			Animates: []*Animate{{Name: "Animates"}},
		}).(*Type)
	case Point:
		return any(&Point{
			// Initialisation of associations
		}).(*Type)
	case Polygone:
		return any(&Polygone{
			// Initialisation of associations
			// field is initialized with an instance of Animate with the name of the field
			Animates: []*Animate{{Name: "Animates"}},
		}).(*Type)
	case Polyline:
		return any(&Polyline{
			// Initialisation of associations
			// field is initialized with an instance of Animate with the name of the field
			Animates: []*Animate{{Name: "Animates"}},
		}).(*Type)
	case Rect:
		return any(&Rect{
			// Initialisation of associations
			// field is initialized with an instance of Animate with the name of the field
			Animations: []*Animate{{Name: "Animations"}},
			// field is initialized with an instance of RectAnchoredText with the name of the field
			RectAnchoredTexts: []*RectAnchoredText{{Name: "RectAnchoredTexts"}},
			// field is initialized with an instance of RectAnchoredRect with the name of the field
			RectAnchoredRects: []*RectAnchoredRect{{Name: "RectAnchoredRects"}},
			// field is initialized with an instance of RectAnchoredPath with the name of the field
			RectAnchoredPaths: []*RectAnchoredPath{{Name: "RectAnchoredPaths"}},
		}).(*Type)
	case RectAnchoredPath:
		return any(&RectAnchoredPath{
			// Initialisation of associations
		}).(*Type)
	case RectAnchoredRect:
		return any(&RectAnchoredRect{
			// Initialisation of associations
		}).(*Type)
	case RectAnchoredText:
		return any(&RectAnchoredText{
			// Initialisation of associations
			// field is initialized with an instance of Animate with the name of the field
			Animates: []*Animate{{Name: "Animates"}},
		}).(*Type)
	case RectLinkLink:
		return any(&RectLinkLink{
			// Initialisation of associations
			// field is initialized with an instance of Rect with the name of the field
			Start: &Rect{Name: "Start"},
			// field is initialized with an instance of Link with the name of the field
			End: &Link{Name: "End"},
		}).(*Type)
	case SVG:
		return any(&SVG{
			// Initialisation of associations
			// field is initialized with an instance of Layer with the name of the field
			Layers: []*Layer{{Name: "Layers"}},
			// field is initialized with an instance of Rect with the name of the field
			StartRect: &Rect{Name: "StartRect"},
			// field is initialized with an instance of Rect with the name of the field
			EndRect: &Rect{Name: "EndRect"},
		}).(*Type)
	case Text:
		return any(&Text{
			// Initialisation of associations
			// field is initialized with an instance of Animate with the name of the field
			Animates: []*Animate{{Name: "Animates"}},
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
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

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
	// reverse maps of direct associations of Ellipse
	case Ellipse:
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
		}
	// reverse maps of direct associations of RectAnchoredPath
	case RectAnchoredPath:
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
	// reverse maps of direct associations of Text
	case Text:
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
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

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
			res := make(map[*Animate]*Circle)
			for circle := range stage.Circles {
				for _, animate_ := range circle.Animations {
					res[animate_] = circle
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Ellipse
	case Ellipse:
		switch fieldname {
		// insertion point for per direct association field
		case "Animates":
			res := make(map[*Animate]*Ellipse)
			for ellipse := range stage.Ellipses {
				for _, animate_ := range ellipse.Animates {
					res[animate_] = ellipse
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Layer
	case Layer:
		switch fieldname {
		// insertion point for per direct association field
		case "Rects":
			res := make(map[*Rect]*Layer)
			for layer := range stage.Layers {
				for _, rect_ := range layer.Rects {
					res[rect_] = layer
				}
			}
			return any(res).(map[*End]*Start)
		case "Texts":
			res := make(map[*Text]*Layer)
			for layer := range stage.Layers {
				for _, text_ := range layer.Texts {
					res[text_] = layer
				}
			}
			return any(res).(map[*End]*Start)
		case "Circles":
			res := make(map[*Circle]*Layer)
			for layer := range stage.Layers {
				for _, circle_ := range layer.Circles {
					res[circle_] = layer
				}
			}
			return any(res).(map[*End]*Start)
		case "Lines":
			res := make(map[*Line]*Layer)
			for layer := range stage.Layers {
				for _, line_ := range layer.Lines {
					res[line_] = layer
				}
			}
			return any(res).(map[*End]*Start)
		case "Ellipses":
			res := make(map[*Ellipse]*Layer)
			for layer := range stage.Layers {
				for _, ellipse_ := range layer.Ellipses {
					res[ellipse_] = layer
				}
			}
			return any(res).(map[*End]*Start)
		case "Polylines":
			res := make(map[*Polyline]*Layer)
			for layer := range stage.Layers {
				for _, polyline_ := range layer.Polylines {
					res[polyline_] = layer
				}
			}
			return any(res).(map[*End]*Start)
		case "Polygones":
			res := make(map[*Polygone]*Layer)
			for layer := range stage.Layers {
				for _, polygone_ := range layer.Polygones {
					res[polygone_] = layer
				}
			}
			return any(res).(map[*End]*Start)
		case "Paths":
			res := make(map[*Path]*Layer)
			for layer := range stage.Layers {
				for _, path_ := range layer.Paths {
					res[path_] = layer
				}
			}
			return any(res).(map[*End]*Start)
		case "Links":
			res := make(map[*Link]*Layer)
			for layer := range stage.Layers {
				for _, link_ := range layer.Links {
					res[link_] = layer
				}
			}
			return any(res).(map[*End]*Start)
		case "RectLinkLinks":
			res := make(map[*RectLinkLink]*Layer)
			for layer := range stage.Layers {
				for _, rectlinklink_ := range layer.RectLinkLinks {
					res[rectlinklink_] = layer
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Line
	case Line:
		switch fieldname {
		// insertion point for per direct association field
		case "Animates":
			res := make(map[*Animate]*Line)
			for line := range stage.Lines {
				for _, animate_ := range line.Animates {
					res[animate_] = line
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Link
	case Link:
		switch fieldname {
		// insertion point for per direct association field
		case "TextAtArrowEnd":
			res := make(map[*LinkAnchoredText]*Link)
			for link := range stage.Links {
				for _, linkanchoredtext_ := range link.TextAtArrowEnd {
					res[linkanchoredtext_] = link
				}
			}
			return any(res).(map[*End]*Start)
		case "TextAtArrowStart":
			res := make(map[*LinkAnchoredText]*Link)
			for link := range stage.Links {
				for _, linkanchoredtext_ := range link.TextAtArrowStart {
					res[linkanchoredtext_] = link
				}
			}
			return any(res).(map[*End]*Start)
		case "ControlPoints":
			res := make(map[*Point]*Link)
			for link := range stage.Links {
				for _, point_ := range link.ControlPoints {
					res[point_] = link
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of LinkAnchoredText
	case LinkAnchoredText:
		switch fieldname {
		// insertion point for per direct association field
		case "Animates":
			res := make(map[*Animate]*LinkAnchoredText)
			for linkanchoredtext := range stage.LinkAnchoredTexts {
				for _, animate_ := range linkanchoredtext.Animates {
					res[animate_] = linkanchoredtext
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Path
	case Path:
		switch fieldname {
		// insertion point for per direct association field
		case "Animates":
			res := make(map[*Animate]*Path)
			for path := range stage.Paths {
				for _, animate_ := range path.Animates {
					res[animate_] = path
				}
			}
			return any(res).(map[*End]*Start)
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
			res := make(map[*Animate]*Polygone)
			for polygone := range stage.Polygones {
				for _, animate_ := range polygone.Animates {
					res[animate_] = polygone
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Polyline
	case Polyline:
		switch fieldname {
		// insertion point for per direct association field
		case "Animates":
			res := make(map[*Animate]*Polyline)
			for polyline := range stage.Polylines {
				for _, animate_ := range polyline.Animates {
					res[animate_] = polyline
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Rect
	case Rect:
		switch fieldname {
		// insertion point for per direct association field
		case "Animations":
			res := make(map[*Animate]*Rect)
			for rect := range stage.Rects {
				for _, animate_ := range rect.Animations {
					res[animate_] = rect
				}
			}
			return any(res).(map[*End]*Start)
		case "RectAnchoredTexts":
			res := make(map[*RectAnchoredText]*Rect)
			for rect := range stage.Rects {
				for _, rectanchoredtext_ := range rect.RectAnchoredTexts {
					res[rectanchoredtext_] = rect
				}
			}
			return any(res).(map[*End]*Start)
		case "RectAnchoredRects":
			res := make(map[*RectAnchoredRect]*Rect)
			for rect := range stage.Rects {
				for _, rectanchoredrect_ := range rect.RectAnchoredRects {
					res[rectanchoredrect_] = rect
				}
			}
			return any(res).(map[*End]*Start)
		case "RectAnchoredPaths":
			res := make(map[*RectAnchoredPath]*Rect)
			for rect := range stage.Rects {
				for _, rectanchoredpath_ := range rect.RectAnchoredPaths {
					res[rectanchoredpath_] = rect
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of RectAnchoredPath
	case RectAnchoredPath:
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
			res := make(map[*Animate]*RectAnchoredText)
			for rectanchoredtext := range stage.RectAnchoredTexts {
				for _, animate_ := range rectanchoredtext.Animates {
					res[animate_] = rectanchoredtext
				}
			}
			return any(res).(map[*End]*Start)
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
			res := make(map[*Layer]*SVG)
			for svg := range stage.SVGs {
				for _, layer_ := range svg.Layers {
					res[layer_] = svg
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Text
	case Text:
		switch fieldname {
		// insertion point for per direct association field
		case "Animates":
			res := make(map[*Animate]*Text)
			for text := range stage.Texts {
				for _, animate_ := range text.Animates {
					res[animate_] = text
				}
			}
			return any(res).(map[*End]*Start)
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Animate:
		res = "Animate"
	case Circle:
		res = "Circle"
	case Ellipse:
		res = "Ellipse"
	case Layer:
		res = "Layer"
	case Line:
		res = "Line"
	case Link:
		res = "Link"
	case LinkAnchoredText:
		res = "LinkAnchoredText"
	case Path:
		res = "Path"
	case Point:
		res = "Point"
	case Polygone:
		res = "Polygone"
	case Polyline:
		res = "Polyline"
	case Rect:
		res = "Rect"
	case RectAnchoredPath:
		res = "RectAnchoredPath"
	case RectAnchoredRect:
		res = "RectAnchoredRect"
	case RectAnchoredText:
		res = "RectAnchoredText"
	case RectLinkLink:
		res = "RectLinkLink"
	case SVG:
		res = "SVG"
	case Text:
		res = "Text"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Animate:
		res = "Animate"
	case *Circle:
		res = "Circle"
	case *Ellipse:
		res = "Ellipse"
	case *Layer:
		res = "Layer"
	case *Line:
		res = "Line"
	case *Link:
		res = "Link"
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
	case *RectAnchoredRect:
		res = "RectAnchoredRect"
	case *RectAnchoredText:
		res = "RectAnchoredText"
	case *RectLinkLink:
		res = "RectLinkLink"
	case *SVG:
		res = "SVG"
	case *Text:
		res = "Text"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Animate:
		res = []string{"Name", "AttributeName", "Values", "From", "To", "Dur", "RepeatCount"}
	case Circle:
		res = []string{"Name", "CX", "CY", "Radius", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Animations"}
	case Ellipse:
		res = []string{"Name", "CX", "CY", "RX", "RY", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Animates"}
	case Layer:
		res = []string{"Display", "Name", "Rects", "Texts", "Circles", "Lines", "Ellipses", "Polylines", "Polygones", "Paths", "Links", "RectLinkLinks"}
	case Line:
		res = []string{"Name", "X1", "Y1", "X2", "Y2", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Animates", "MouseClickX", "MouseClickY"}
	case Link:
		res = []string{"Name", "Type", "IsBezierCurve", "Start", "StartAnchorType", "End", "EndAnchorType", "StartOrientation", "StartRatio", "EndOrientation", "EndRatio", "CornerOffsetRatio", "CornerRadius", "HasEndArrow", "EndArrowSize", "HasStartArrow", "StartArrowSize", "TextAtArrowEnd", "TextAtArrowStart", "ControlPoints", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case LinkAnchoredText:
		res = []string{"Name", "Content", "AutomaticLayout", "LinkAnchorType", "X_Offset", "Y_Offset", "FontWeight", "FontSize", "LetterSpacing", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Animates"}
	case Path:
		res = []string{"Name", "Definition", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Animates"}
	case Point:
		res = []string{"Name", "X", "Y"}
	case Polygone:
		res = []string{"Name", "Points", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Animates"}
	case Polyline:
		res = []string{"Name", "Points", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Animates"}
	case Rect:
		res = []string{"Name", "X", "Y", "Width", "Height", "RX", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Animations", "IsSelectable", "IsSelected", "CanHaveLeftHandle", "HasLeftHandle", "CanHaveRightHandle", "HasRightHandle", "CanHaveTopHandle", "HasTopHandle", "IsScalingProportionally", "CanHaveBottomHandle", "HasBottomHandle", "CanMoveHorizontaly", "CanMoveVerticaly", "RectAnchoredTexts", "RectAnchoredRects", "RectAnchoredPaths"}
	case RectAnchoredPath:
		res = []string{"Name", "Definition", "X_Offset", "Y_Offset", "RectAnchorType", "ScalePropotionnally", "AppliedScaling", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case RectAnchoredRect:
		res = []string{"Name", "X", "Y", "Width", "Height", "RX", "X_Offset", "Y_Offset", "RectAnchorType", "WidthFollowRect", "HeightFollowRect", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case RectAnchoredText:
		res = []string{"Name", "Content", "FontWeight", "FontSize", "FontStyle", "X_Offset", "Y_Offset", "RectAnchorType", "TextAnchorType", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Animates"}
	case RectLinkLink:
		res = []string{"Name", "Start", "End", "TargetAnchorPosition", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case SVG:
		res = []string{"Name", "Layers", "DrawingState", "StartRect", "EndRect", "IsEditable"}
	case Text:
		res = []string{"Name", "X", "Y", "Content", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Animates"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case Animate:
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
	case Circle:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layer"
		rf.Fieldname = "Circles"
		res = append(res, rf)
	case Ellipse:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layer"
		rf.Fieldname = "Ellipses"
		res = append(res, rf)
	case Layer:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SVG"
		rf.Fieldname = "Layers"
		res = append(res, rf)
	case Line:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layer"
		rf.Fieldname = "Lines"
		res = append(res, rf)
	case Link:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layer"
		rf.Fieldname = "Links"
		res = append(res, rf)
	case LinkAnchoredText:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Link"
		rf.Fieldname = "TextAtArrowEnd"
		res = append(res, rf)
		rf.GongstructName = "Link"
		rf.Fieldname = "TextAtArrowStart"
		res = append(res, rf)
	case Path:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layer"
		rf.Fieldname = "Paths"
		res = append(res, rf)
	case Point:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Link"
		rf.Fieldname = "ControlPoints"
		res = append(res, rf)
	case Polygone:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layer"
		rf.Fieldname = "Polygones"
		res = append(res, rf)
	case Polyline:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layer"
		rf.Fieldname = "Polylines"
		res = append(res, rf)
	case Rect:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layer"
		rf.Fieldname = "Rects"
		res = append(res, rf)
	case RectAnchoredPath:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Rect"
		rf.Fieldname = "RectAnchoredPaths"
		res = append(res, rf)
	case RectAnchoredRect:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Rect"
		rf.Fieldname = "RectAnchoredRects"
		res = append(res, rf)
	case RectAnchoredText:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Rect"
		rf.Fieldname = "RectAnchoredTexts"
		res = append(res, rf)
	case RectLinkLink:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layer"
		rf.Fieldname = "RectLinkLinks"
		res = append(res, rf)
	case SVG:
		var rf ReverseField
		_ = rf
	case Text:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layer"
		rf.Fieldname = "Texts"
		res = append(res, rf)
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Animate:
		res = []string{"Name", "AttributeName", "Values", "From", "To", "Dur", "RepeatCount"}
	case *Circle:
		res = []string{"Name", "CX", "CY", "Radius", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Animations"}
	case *Ellipse:
		res = []string{"Name", "CX", "CY", "RX", "RY", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Animates"}
	case *Layer:
		res = []string{"Display", "Name", "Rects", "Texts", "Circles", "Lines", "Ellipses", "Polylines", "Polygones", "Paths", "Links", "RectLinkLinks"}
	case *Line:
		res = []string{"Name", "X1", "Y1", "X2", "Y2", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Animates", "MouseClickX", "MouseClickY"}
	case *Link:
		res = []string{"Name", "Type", "IsBezierCurve", "Start", "StartAnchorType", "End", "EndAnchorType", "StartOrientation", "StartRatio", "EndOrientation", "EndRatio", "CornerOffsetRatio", "CornerRadius", "HasEndArrow", "EndArrowSize", "HasStartArrow", "StartArrowSize", "TextAtArrowEnd", "TextAtArrowStart", "ControlPoints", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case *LinkAnchoredText:
		res = []string{"Name", "Content", "AutomaticLayout", "LinkAnchorType", "X_Offset", "Y_Offset", "FontWeight", "FontSize", "LetterSpacing", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Animates"}
	case *Path:
		res = []string{"Name", "Definition", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Animates"}
	case *Point:
		res = []string{"Name", "X", "Y"}
	case *Polygone:
		res = []string{"Name", "Points", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Animates"}
	case *Polyline:
		res = []string{"Name", "Points", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Animates"}
	case *Rect:
		res = []string{"Name", "X", "Y", "Width", "Height", "RX", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Animations", "IsSelectable", "IsSelected", "CanHaveLeftHandle", "HasLeftHandle", "CanHaveRightHandle", "HasRightHandle", "CanHaveTopHandle", "HasTopHandle", "IsScalingProportionally", "CanHaveBottomHandle", "HasBottomHandle", "CanMoveHorizontaly", "CanMoveVerticaly", "RectAnchoredTexts", "RectAnchoredRects", "RectAnchoredPaths"}
	case *RectAnchoredPath:
		res = []string{"Name", "Definition", "X_Offset", "Y_Offset", "RectAnchorType", "ScalePropotionnally", "AppliedScaling", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case *RectAnchoredRect:
		res = []string{"Name", "X", "Y", "Width", "Height", "RX", "X_Offset", "Y_Offset", "RectAnchorType", "WidthFollowRect", "HeightFollowRect", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case *RectAnchoredText:
		res = []string{"Name", "Content", "FontWeight", "FontSize", "FontStyle", "X_Offset", "Y_Offset", "RectAnchorType", "TextAnchorType", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Animates"}
	case *RectLinkLink:
		res = []string{"Name", "Start", "End", "TargetAnchorPosition", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case *SVG:
		res = []string{"Name", "Layers", "DrawingState", "StartRect", "EndRect", "IsEditable"}
	case *Text:
		res = []string{"Name", "X", "Y", "Content", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Animates"}
	}
	return
}

type GongFieldValueType string

const (
	GongFieldValueTypeInt     GongFieldValueType = "GongFieldValueTypeInt"
	GongFieldValueTypeFloat   GongFieldValueType = "GongFieldValueTypeFloat"
	GongFieldValueTypeBool    GongFieldValueType = "GongFieldValueTypeBool"
	GongFieldValueTypeOthers  GongFieldValueType = "GongFieldValueTypeOthers"
)

type GongFieldValue struct {
	valueString string
	GongFieldValueType
	valueInt   int
	valueFloat float64
	valueBool  bool
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

func GetFieldStringValueFromPointer(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *Animate:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "AttributeName":
			res.valueString = inferedInstance.AttributeName
		case "Values":
			res.valueString = inferedInstance.Values
		case "From":
			res.valueString = inferedInstance.From
		case "To":
			res.valueString = inferedInstance.To
		case "Dur":
			res.valueString = inferedInstance.Dur
		case "RepeatCount":
			res.valueString = inferedInstance.RepeatCount
		}
	case *Circle:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "CX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CX)
			res.valueFloat = inferedInstance.CX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "CY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CY)
			res.valueFloat = inferedInstance.CY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Radius":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Radius)
			res.valueFloat = inferedInstance.Radius
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Animations":
			for idx, __instance__ := range inferedInstance.Animations {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *Ellipse:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "CX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CX)
			res.valueFloat = inferedInstance.CX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "CY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CY)
			res.valueFloat = inferedInstance.CY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "RX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.RX)
			res.valueFloat = inferedInstance.RX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "RY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.RY)
			res.valueFloat = inferedInstance.RY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Animates":
			for idx, __instance__ := range inferedInstance.Animates {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *Layer:
		switch fieldName {
		// string value of fields
		case "Display":
			res.valueString = fmt.Sprintf("%t", inferedInstance.Display)
			res.valueBool = inferedInstance.Display
			res.GongFieldValueType = GongFieldValueTypeBool
		case "Name":
			res.valueString = inferedInstance.Name
		case "Rects":
			for idx, __instance__ := range inferedInstance.Rects {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Texts":
			for idx, __instance__ := range inferedInstance.Texts {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Circles":
			for idx, __instance__ := range inferedInstance.Circles {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Lines":
			for idx, __instance__ := range inferedInstance.Lines {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Ellipses":
			for idx, __instance__ := range inferedInstance.Ellipses {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Polylines":
			for idx, __instance__ := range inferedInstance.Polylines {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Polygones":
			for idx, __instance__ := range inferedInstance.Polygones {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Paths":
			for idx, __instance__ := range inferedInstance.Paths {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Links":
			for idx, __instance__ := range inferedInstance.Links {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "RectLinkLinks":
			for idx, __instance__ := range inferedInstance.RectLinkLinks {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *Line:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "X1":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X1)
			res.valueFloat = inferedInstance.X1
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y1":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y1)
			res.valueFloat = inferedInstance.Y1
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "X2":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X2)
			res.valueFloat = inferedInstance.X2
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y2":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y2)
			res.valueFloat = inferedInstance.Y2
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Animates":
			for idx, __instance__ := range inferedInstance.Animates {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "MouseClickX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.MouseClickX)
			res.valueFloat = inferedInstance.MouseClickX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "MouseClickY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.MouseClickY)
			res.valueFloat = inferedInstance.MouseClickY
			res.GongFieldValueType = GongFieldValueTypeFloat
		}
	case *Link:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Type":
			enum := inferedInstance.Type
			res.valueString = enum.ToCodeString()
		case "IsBezierCurve":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsBezierCurve)
			res.valueBool = inferedInstance.IsBezierCurve
			res.GongFieldValueType = GongFieldValueTypeBool
		case "Start":
			if inferedInstance.Start != nil {
				res.valueString = inferedInstance.Start.Name
			}
		case "StartAnchorType":
			enum := inferedInstance.StartAnchorType
			res.valueString = enum.ToCodeString()
		case "End":
			if inferedInstance.End != nil {
				res.valueString = inferedInstance.End.Name
			}
		case "EndAnchorType":
			enum := inferedInstance.EndAnchorType
			res.valueString = enum.ToCodeString()
		case "StartOrientation":
			enum := inferedInstance.StartOrientation
			res.valueString = enum.ToCodeString()
		case "StartRatio":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StartRatio)
			res.valueFloat = inferedInstance.StartRatio
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndOrientation":
			enum := inferedInstance.EndOrientation
			res.valueString = enum.ToCodeString()
		case "EndRatio":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndRatio)
			res.valueFloat = inferedInstance.EndRatio
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "CornerOffsetRatio":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CornerOffsetRatio)
			res.valueFloat = inferedInstance.CornerOffsetRatio
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "CornerRadius":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CornerRadius)
			res.valueFloat = inferedInstance.CornerRadius
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "HasEndArrow":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasEndArrow)
			res.valueBool = inferedInstance.HasEndArrow
			res.GongFieldValueType = GongFieldValueTypeBool
		case "EndArrowSize":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndArrowSize)
			res.valueFloat = inferedInstance.EndArrowSize
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "HasStartArrow":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasStartArrow)
			res.valueBool = inferedInstance.HasStartArrow
			res.GongFieldValueType = GongFieldValueTypeBool
		case "StartArrowSize":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StartArrowSize)
			res.valueFloat = inferedInstance.StartArrowSize
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "TextAtArrowEnd":
			for idx, __instance__ := range inferedInstance.TextAtArrowEnd {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "TextAtArrowStart":
			for idx, __instance__ := range inferedInstance.TextAtArrowStart {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "ControlPoints":
			for idx, __instance__ := range inferedInstance.ControlPoints {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case *LinkAnchoredText:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		case "AutomaticLayout":
			res.valueString = fmt.Sprintf("%t", inferedInstance.AutomaticLayout)
			res.valueBool = inferedInstance.AutomaticLayout
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LinkAnchorType":
			enum := inferedInstance.LinkAnchorType
			res.valueString = enum.ToCodeString()
		case "X_Offset":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X_Offset)
			res.valueFloat = inferedInstance.X_Offset
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y_Offset":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y_Offset)
			res.valueFloat = inferedInstance.Y_Offset
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "FontWeight":
			res.valueString = inferedInstance.FontWeight
		case "FontSize":
			res.valueString = inferedInstance.FontSize
		case "LetterSpacing":
			res.valueString = inferedInstance.LetterSpacing
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Animates":
			for idx, __instance__ := range inferedInstance.Animates {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *Path:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Definition":
			res.valueString = inferedInstance.Definition
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Animates":
			for idx, __instance__ := range inferedInstance.Animates {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *Point:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "X":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X)
			res.valueFloat = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y)
			res.valueFloat = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeFloat
		}
	case *Polygone:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Points":
			res.valueString = inferedInstance.Points
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Animates":
			for idx, __instance__ := range inferedInstance.Animates {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *Polyline:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Points":
			res.valueString = inferedInstance.Points
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Animates":
			for idx, __instance__ := range inferedInstance.Animates {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *Rect:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "X":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X)
			res.valueFloat = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y)
			res.valueFloat = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Width":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Width)
			res.valueFloat = inferedInstance.Width
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Height":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Height)
			res.valueFloat = inferedInstance.Height
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "RX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.RX)
			res.valueFloat = inferedInstance.RX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Animations":
			for idx, __instance__ := range inferedInstance.Animations {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "IsSelectable":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsSelectable)
			res.valueBool = inferedInstance.IsSelectable
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsSelected":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsSelected)
			res.valueBool = inferedInstance.IsSelected
			res.GongFieldValueType = GongFieldValueTypeBool
		case "CanHaveLeftHandle":
			res.valueString = fmt.Sprintf("%t", inferedInstance.CanHaveLeftHandle)
			res.valueBool = inferedInstance.CanHaveLeftHandle
			res.GongFieldValueType = GongFieldValueTypeBool
		case "HasLeftHandle":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasLeftHandle)
			res.valueBool = inferedInstance.HasLeftHandle
			res.GongFieldValueType = GongFieldValueTypeBool
		case "CanHaveRightHandle":
			res.valueString = fmt.Sprintf("%t", inferedInstance.CanHaveRightHandle)
			res.valueBool = inferedInstance.CanHaveRightHandle
			res.GongFieldValueType = GongFieldValueTypeBool
		case "HasRightHandle":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasRightHandle)
			res.valueBool = inferedInstance.HasRightHandle
			res.GongFieldValueType = GongFieldValueTypeBool
		case "CanHaveTopHandle":
			res.valueString = fmt.Sprintf("%t", inferedInstance.CanHaveTopHandle)
			res.valueBool = inferedInstance.CanHaveTopHandle
			res.GongFieldValueType = GongFieldValueTypeBool
		case "HasTopHandle":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasTopHandle)
			res.valueBool = inferedInstance.HasTopHandle
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsScalingProportionally":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsScalingProportionally)
			res.valueBool = inferedInstance.IsScalingProportionally
			res.GongFieldValueType = GongFieldValueTypeBool
		case "CanHaveBottomHandle":
			res.valueString = fmt.Sprintf("%t", inferedInstance.CanHaveBottomHandle)
			res.valueBool = inferedInstance.CanHaveBottomHandle
			res.GongFieldValueType = GongFieldValueTypeBool
		case "HasBottomHandle":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasBottomHandle)
			res.valueBool = inferedInstance.HasBottomHandle
			res.GongFieldValueType = GongFieldValueTypeBool
		case "CanMoveHorizontaly":
			res.valueString = fmt.Sprintf("%t", inferedInstance.CanMoveHorizontaly)
			res.valueBool = inferedInstance.CanMoveHorizontaly
			res.GongFieldValueType = GongFieldValueTypeBool
		case "CanMoveVerticaly":
			res.valueString = fmt.Sprintf("%t", inferedInstance.CanMoveVerticaly)
			res.valueBool = inferedInstance.CanMoveVerticaly
			res.GongFieldValueType = GongFieldValueTypeBool
		case "RectAnchoredTexts":
			for idx, __instance__ := range inferedInstance.RectAnchoredTexts {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "RectAnchoredRects":
			for idx, __instance__ := range inferedInstance.RectAnchoredRects {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "RectAnchoredPaths":
			for idx, __instance__ := range inferedInstance.RectAnchoredPaths {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *RectAnchoredPath:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Definition":
			res.valueString = inferedInstance.Definition
		case "X_Offset":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X_Offset)
			res.valueFloat = inferedInstance.X_Offset
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y_Offset":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y_Offset)
			res.valueFloat = inferedInstance.Y_Offset
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "RectAnchorType":
			enum := inferedInstance.RectAnchorType
			res.valueString = enum.ToCodeString()
		case "ScalePropotionnally":
			res.valueString = fmt.Sprintf("%t", inferedInstance.ScalePropotionnally)
			res.valueBool = inferedInstance.ScalePropotionnally
			res.GongFieldValueType = GongFieldValueTypeBool
		case "AppliedScaling":
			res.valueString = fmt.Sprintf("%f", inferedInstance.AppliedScaling)
			res.valueFloat = inferedInstance.AppliedScaling
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case *RectAnchoredRect:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "X":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X)
			res.valueFloat = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y)
			res.valueFloat = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Width":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Width)
			res.valueFloat = inferedInstance.Width
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Height":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Height)
			res.valueFloat = inferedInstance.Height
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "RX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.RX)
			res.valueFloat = inferedInstance.RX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "X_Offset":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X_Offset)
			res.valueFloat = inferedInstance.X_Offset
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y_Offset":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y_Offset)
			res.valueFloat = inferedInstance.Y_Offset
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "RectAnchorType":
			enum := inferedInstance.RectAnchorType
			res.valueString = enum.ToCodeString()
		case "WidthFollowRect":
			res.valueString = fmt.Sprintf("%t", inferedInstance.WidthFollowRect)
			res.valueBool = inferedInstance.WidthFollowRect
			res.GongFieldValueType = GongFieldValueTypeBool
		case "HeightFollowRect":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HeightFollowRect)
			res.valueBool = inferedInstance.HeightFollowRect
			res.GongFieldValueType = GongFieldValueTypeBool
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case *RectAnchoredText:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		case "FontWeight":
			res.valueString = inferedInstance.FontWeight
		case "FontSize":
			res.valueString = fmt.Sprintf("%d", inferedInstance.FontSize)
			res.valueInt = inferedInstance.FontSize
			res.GongFieldValueType = GongFieldValueTypeInt
		case "FontStyle":
			res.valueString = inferedInstance.FontStyle
		case "X_Offset":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X_Offset)
			res.valueFloat = inferedInstance.X_Offset
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y_Offset":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y_Offset)
			res.valueFloat = inferedInstance.Y_Offset
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "RectAnchorType":
			enum := inferedInstance.RectAnchorType
			res.valueString = enum.ToCodeString()
		case "TextAnchorType":
			enum := inferedInstance.TextAnchorType
			res.valueString = enum.ToCodeString()
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Animates":
			for idx, __instance__ := range inferedInstance.Animates {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *RectLinkLink:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Start":
			if inferedInstance.Start != nil {
				res.valueString = inferedInstance.Start.Name
			}
		case "End":
			if inferedInstance.End != nil {
				res.valueString = inferedInstance.End.Name
			}
		case "TargetAnchorPosition":
			res.valueString = fmt.Sprintf("%f", inferedInstance.TargetAnchorPosition)
			res.valueFloat = inferedInstance.TargetAnchorPosition
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case *SVG:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Layers":
			for idx, __instance__ := range inferedInstance.Layers {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DrawingState":
			enum := inferedInstance.DrawingState
			res.valueString = enum.ToCodeString()
		case "StartRect":
			if inferedInstance.StartRect != nil {
				res.valueString = inferedInstance.StartRect.Name
			}
		case "EndRect":
			if inferedInstance.EndRect != nil {
				res.valueString = inferedInstance.EndRect.Name
			}
		case "IsEditable":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsEditable)
			res.valueBool = inferedInstance.IsEditable
			res.GongFieldValueType = GongFieldValueTypeBool
		}
	case *Text:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "X":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X)
			res.valueFloat = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y)
			res.valueFloat = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Content":
			res.valueString = inferedInstance.Content
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Animates":
			for idx, __instance__ := range inferedInstance.Animates {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Animate:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "AttributeName":
			res.valueString = inferedInstance.AttributeName
		case "Values":
			res.valueString = inferedInstance.Values
		case "From":
			res.valueString = inferedInstance.From
		case "To":
			res.valueString = inferedInstance.To
		case "Dur":
			res.valueString = inferedInstance.Dur
		case "RepeatCount":
			res.valueString = inferedInstance.RepeatCount
		}
	case Circle:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "CX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CX)
			res.valueFloat = inferedInstance.CX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "CY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CY)
			res.valueFloat = inferedInstance.CY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Radius":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Radius)
			res.valueFloat = inferedInstance.Radius
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Animations":
			for idx, __instance__ := range inferedInstance.Animations {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case Ellipse:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "CX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CX)
			res.valueFloat = inferedInstance.CX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "CY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CY)
			res.valueFloat = inferedInstance.CY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "RX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.RX)
			res.valueFloat = inferedInstance.RX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "RY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.RY)
			res.valueFloat = inferedInstance.RY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Animates":
			for idx, __instance__ := range inferedInstance.Animates {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case Layer:
		switch fieldName {
		// string value of fields
		case "Display":
			res.valueString = fmt.Sprintf("%t", inferedInstance.Display)
			res.valueBool = inferedInstance.Display
			res.GongFieldValueType = GongFieldValueTypeBool
		case "Name":
			res.valueString = inferedInstance.Name
		case "Rects":
			for idx, __instance__ := range inferedInstance.Rects {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Texts":
			for idx, __instance__ := range inferedInstance.Texts {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Circles":
			for idx, __instance__ := range inferedInstance.Circles {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Lines":
			for idx, __instance__ := range inferedInstance.Lines {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Ellipses":
			for idx, __instance__ := range inferedInstance.Ellipses {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Polylines":
			for idx, __instance__ := range inferedInstance.Polylines {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Polygones":
			for idx, __instance__ := range inferedInstance.Polygones {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Paths":
			for idx, __instance__ := range inferedInstance.Paths {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Links":
			for idx, __instance__ := range inferedInstance.Links {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "RectLinkLinks":
			for idx, __instance__ := range inferedInstance.RectLinkLinks {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case Line:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "X1":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X1)
			res.valueFloat = inferedInstance.X1
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y1":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y1)
			res.valueFloat = inferedInstance.Y1
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "X2":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X2)
			res.valueFloat = inferedInstance.X2
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y2":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y2)
			res.valueFloat = inferedInstance.Y2
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Animates":
			for idx, __instance__ := range inferedInstance.Animates {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "MouseClickX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.MouseClickX)
			res.valueFloat = inferedInstance.MouseClickX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "MouseClickY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.MouseClickY)
			res.valueFloat = inferedInstance.MouseClickY
			res.GongFieldValueType = GongFieldValueTypeFloat
		}
	case Link:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Type":
			enum := inferedInstance.Type
			res.valueString = enum.ToCodeString()
		case "IsBezierCurve":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsBezierCurve)
			res.valueBool = inferedInstance.IsBezierCurve
			res.GongFieldValueType = GongFieldValueTypeBool
		case "Start":
			if inferedInstance.Start != nil {
				res.valueString = inferedInstance.Start.Name
			}
		case "StartAnchorType":
			enum := inferedInstance.StartAnchorType
			res.valueString = enum.ToCodeString()
		case "End":
			if inferedInstance.End != nil {
				res.valueString = inferedInstance.End.Name
			}
		case "EndAnchorType":
			enum := inferedInstance.EndAnchorType
			res.valueString = enum.ToCodeString()
		case "StartOrientation":
			enum := inferedInstance.StartOrientation
			res.valueString = enum.ToCodeString()
		case "StartRatio":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StartRatio)
			res.valueFloat = inferedInstance.StartRatio
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndOrientation":
			enum := inferedInstance.EndOrientation
			res.valueString = enum.ToCodeString()
		case "EndRatio":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndRatio)
			res.valueFloat = inferedInstance.EndRatio
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "CornerOffsetRatio":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CornerOffsetRatio)
			res.valueFloat = inferedInstance.CornerOffsetRatio
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "CornerRadius":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CornerRadius)
			res.valueFloat = inferedInstance.CornerRadius
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "HasEndArrow":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasEndArrow)
			res.valueBool = inferedInstance.HasEndArrow
			res.GongFieldValueType = GongFieldValueTypeBool
		case "EndArrowSize":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndArrowSize)
			res.valueFloat = inferedInstance.EndArrowSize
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "HasStartArrow":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasStartArrow)
			res.valueBool = inferedInstance.HasStartArrow
			res.GongFieldValueType = GongFieldValueTypeBool
		case "StartArrowSize":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StartArrowSize)
			res.valueFloat = inferedInstance.StartArrowSize
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "TextAtArrowEnd":
			for idx, __instance__ := range inferedInstance.TextAtArrowEnd {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "TextAtArrowStart":
			for idx, __instance__ := range inferedInstance.TextAtArrowStart {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "ControlPoints":
			for idx, __instance__ := range inferedInstance.ControlPoints {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case LinkAnchoredText:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		case "AutomaticLayout":
			res.valueString = fmt.Sprintf("%t", inferedInstance.AutomaticLayout)
			res.valueBool = inferedInstance.AutomaticLayout
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LinkAnchorType":
			enum := inferedInstance.LinkAnchorType
			res.valueString = enum.ToCodeString()
		case "X_Offset":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X_Offset)
			res.valueFloat = inferedInstance.X_Offset
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y_Offset":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y_Offset)
			res.valueFloat = inferedInstance.Y_Offset
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "FontWeight":
			res.valueString = inferedInstance.FontWeight
		case "FontSize":
			res.valueString = inferedInstance.FontSize
		case "LetterSpacing":
			res.valueString = inferedInstance.LetterSpacing
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Animates":
			for idx, __instance__ := range inferedInstance.Animates {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case Path:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Definition":
			res.valueString = inferedInstance.Definition
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Animates":
			for idx, __instance__ := range inferedInstance.Animates {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case Point:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "X":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X)
			res.valueFloat = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y)
			res.valueFloat = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeFloat
		}
	case Polygone:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Points":
			res.valueString = inferedInstance.Points
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Animates":
			for idx, __instance__ := range inferedInstance.Animates {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case Polyline:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Points":
			res.valueString = inferedInstance.Points
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Animates":
			for idx, __instance__ := range inferedInstance.Animates {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case Rect:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "X":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X)
			res.valueFloat = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y)
			res.valueFloat = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Width":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Width)
			res.valueFloat = inferedInstance.Width
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Height":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Height)
			res.valueFloat = inferedInstance.Height
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "RX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.RX)
			res.valueFloat = inferedInstance.RX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Animations":
			for idx, __instance__ := range inferedInstance.Animations {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "IsSelectable":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsSelectable)
			res.valueBool = inferedInstance.IsSelectable
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsSelected":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsSelected)
			res.valueBool = inferedInstance.IsSelected
			res.GongFieldValueType = GongFieldValueTypeBool
		case "CanHaveLeftHandle":
			res.valueString = fmt.Sprintf("%t", inferedInstance.CanHaveLeftHandle)
			res.valueBool = inferedInstance.CanHaveLeftHandle
			res.GongFieldValueType = GongFieldValueTypeBool
		case "HasLeftHandle":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasLeftHandle)
			res.valueBool = inferedInstance.HasLeftHandle
			res.GongFieldValueType = GongFieldValueTypeBool
		case "CanHaveRightHandle":
			res.valueString = fmt.Sprintf("%t", inferedInstance.CanHaveRightHandle)
			res.valueBool = inferedInstance.CanHaveRightHandle
			res.GongFieldValueType = GongFieldValueTypeBool
		case "HasRightHandle":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasRightHandle)
			res.valueBool = inferedInstance.HasRightHandle
			res.GongFieldValueType = GongFieldValueTypeBool
		case "CanHaveTopHandle":
			res.valueString = fmt.Sprintf("%t", inferedInstance.CanHaveTopHandle)
			res.valueBool = inferedInstance.CanHaveTopHandle
			res.GongFieldValueType = GongFieldValueTypeBool
		case "HasTopHandle":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasTopHandle)
			res.valueBool = inferedInstance.HasTopHandle
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsScalingProportionally":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsScalingProportionally)
			res.valueBool = inferedInstance.IsScalingProportionally
			res.GongFieldValueType = GongFieldValueTypeBool
		case "CanHaveBottomHandle":
			res.valueString = fmt.Sprintf("%t", inferedInstance.CanHaveBottomHandle)
			res.valueBool = inferedInstance.CanHaveBottomHandle
			res.GongFieldValueType = GongFieldValueTypeBool
		case "HasBottomHandle":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasBottomHandle)
			res.valueBool = inferedInstance.HasBottomHandle
			res.GongFieldValueType = GongFieldValueTypeBool
		case "CanMoveHorizontaly":
			res.valueString = fmt.Sprintf("%t", inferedInstance.CanMoveHorizontaly)
			res.valueBool = inferedInstance.CanMoveHorizontaly
			res.GongFieldValueType = GongFieldValueTypeBool
		case "CanMoveVerticaly":
			res.valueString = fmt.Sprintf("%t", inferedInstance.CanMoveVerticaly)
			res.valueBool = inferedInstance.CanMoveVerticaly
			res.GongFieldValueType = GongFieldValueTypeBool
		case "RectAnchoredTexts":
			for idx, __instance__ := range inferedInstance.RectAnchoredTexts {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "RectAnchoredRects":
			for idx, __instance__ := range inferedInstance.RectAnchoredRects {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "RectAnchoredPaths":
			for idx, __instance__ := range inferedInstance.RectAnchoredPaths {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case RectAnchoredPath:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Definition":
			res.valueString = inferedInstance.Definition
		case "X_Offset":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X_Offset)
			res.valueFloat = inferedInstance.X_Offset
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y_Offset":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y_Offset)
			res.valueFloat = inferedInstance.Y_Offset
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "RectAnchorType":
			enum := inferedInstance.RectAnchorType
			res.valueString = enum.ToCodeString()
		case "ScalePropotionnally":
			res.valueString = fmt.Sprintf("%t", inferedInstance.ScalePropotionnally)
			res.valueBool = inferedInstance.ScalePropotionnally
			res.GongFieldValueType = GongFieldValueTypeBool
		case "AppliedScaling":
			res.valueString = fmt.Sprintf("%f", inferedInstance.AppliedScaling)
			res.valueFloat = inferedInstance.AppliedScaling
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case RectAnchoredRect:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "X":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X)
			res.valueFloat = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y)
			res.valueFloat = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Width":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Width)
			res.valueFloat = inferedInstance.Width
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Height":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Height)
			res.valueFloat = inferedInstance.Height
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "RX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.RX)
			res.valueFloat = inferedInstance.RX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "X_Offset":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X_Offset)
			res.valueFloat = inferedInstance.X_Offset
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y_Offset":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y_Offset)
			res.valueFloat = inferedInstance.Y_Offset
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "RectAnchorType":
			enum := inferedInstance.RectAnchorType
			res.valueString = enum.ToCodeString()
		case "WidthFollowRect":
			res.valueString = fmt.Sprintf("%t", inferedInstance.WidthFollowRect)
			res.valueBool = inferedInstance.WidthFollowRect
			res.GongFieldValueType = GongFieldValueTypeBool
		case "HeightFollowRect":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HeightFollowRect)
			res.valueBool = inferedInstance.HeightFollowRect
			res.GongFieldValueType = GongFieldValueTypeBool
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case RectAnchoredText:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		case "FontWeight":
			res.valueString = inferedInstance.FontWeight
		case "FontSize":
			res.valueString = fmt.Sprintf("%d", inferedInstance.FontSize)
			res.valueInt = inferedInstance.FontSize
			res.GongFieldValueType = GongFieldValueTypeInt
		case "FontStyle":
			res.valueString = inferedInstance.FontStyle
		case "X_Offset":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X_Offset)
			res.valueFloat = inferedInstance.X_Offset
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y_Offset":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y_Offset)
			res.valueFloat = inferedInstance.Y_Offset
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "RectAnchorType":
			enum := inferedInstance.RectAnchorType
			res.valueString = enum.ToCodeString()
		case "TextAnchorType":
			enum := inferedInstance.TextAnchorType
			res.valueString = enum.ToCodeString()
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Animates":
			for idx, __instance__ := range inferedInstance.Animates {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case RectLinkLink:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Start":
			if inferedInstance.Start != nil {
				res.valueString = inferedInstance.Start.Name
			}
		case "End":
			if inferedInstance.End != nil {
				res.valueString = inferedInstance.End.Name
			}
		case "TargetAnchorPosition":
			res.valueString = fmt.Sprintf("%f", inferedInstance.TargetAnchorPosition)
			res.valueFloat = inferedInstance.TargetAnchorPosition
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case SVG:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Layers":
			for idx, __instance__ := range inferedInstance.Layers {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DrawingState":
			enum := inferedInstance.DrawingState
			res.valueString = enum.ToCodeString()
		case "StartRect":
			if inferedInstance.StartRect != nil {
				res.valueString = inferedInstance.StartRect.Name
			}
		case "EndRect":
			if inferedInstance.EndRect != nil {
				res.valueString = inferedInstance.EndRect.Name
			}
		case "IsEditable":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsEditable)
			res.valueBool = inferedInstance.IsEditable
			res.GongFieldValueType = GongFieldValueTypeBool
		}
	case Text:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "X":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X)
			res.valueFloat = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y)
			res.valueFloat = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Content":
			res.valueString = inferedInstance.Content
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Animates":
			for idx, __instance__ := range inferedInstance.Animates {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
