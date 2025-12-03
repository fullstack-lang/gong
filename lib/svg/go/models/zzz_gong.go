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
	"time"

	svg_go "github.com/fullstack-lang/gong/lib/svg/go"
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

var _ = __Gong__Abs
var _ = strings.Clone("")

const ProbeTreeSidebarSuffix = ":sidebar of the probe"
const ProbeTableSuffix = ":table of the probe"
const ProbeFormSuffix = ":form of the probe"
const ProbeSplitSuffix = ":probe of the probe"

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeFormStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeFormSuffix
}

func (stage *Stage) GetProbeTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTableSuffix
}

func (stage *Stage) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeSplitSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")
var _ = errUnkownEnum

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

var _ = __dummy__fmt_variable

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

var _ = __dummy_math_variable

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void
var _ = __member

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
// swagger:ignore
type Stage struct {
	name               string
	commitId           uint // commitId is updated at each commit
	commitTimeStamp    time.Time
	contentWhenParsed  string
	commitIdWhenParsed uint

	// insertion point for definition of arrays registering instances
	Animates           map[*Animate]struct{}
	Animates_mapString map[string]*Animate

	// insertion point for slice of pointers maps
	OnAfterAnimateCreateCallback OnAfterCreateInterface[Animate]
	OnAfterAnimateUpdateCallback OnAfterUpdateInterface[Animate]
	OnAfterAnimateDeleteCallback OnAfterDeleteInterface[Animate]
	OnAfterAnimateReadCallback   OnAfterReadInterface[Animate]

	Circles           map[*Circle]struct{}
	Circles_mapString map[string]*Circle

	// insertion point for slice of pointers maps
	Circle_Animations_reverseMap map[*Animate]*Circle

	OnAfterCircleCreateCallback OnAfterCreateInterface[Circle]
	OnAfterCircleUpdateCallback OnAfterUpdateInterface[Circle]
	OnAfterCircleDeleteCallback OnAfterDeleteInterface[Circle]
	OnAfterCircleReadCallback   OnAfterReadInterface[Circle]

	Conditions           map[*Condition]struct{}
	Conditions_mapString map[string]*Condition

	// insertion point for slice of pointers maps
	OnAfterConditionCreateCallback OnAfterCreateInterface[Condition]
	OnAfterConditionUpdateCallback OnAfterUpdateInterface[Condition]
	OnAfterConditionDeleteCallback OnAfterDeleteInterface[Condition]
	OnAfterConditionReadCallback   OnAfterReadInterface[Condition]

	ControlPoints           map[*ControlPoint]struct{}
	ControlPoints_mapString map[string]*ControlPoint

	// insertion point for slice of pointers maps
	OnAfterControlPointCreateCallback OnAfterCreateInterface[ControlPoint]
	OnAfterControlPointUpdateCallback OnAfterUpdateInterface[ControlPoint]
	OnAfterControlPointDeleteCallback OnAfterDeleteInterface[ControlPoint]
	OnAfterControlPointReadCallback   OnAfterReadInterface[ControlPoint]

	Ellipses           map[*Ellipse]struct{}
	Ellipses_mapString map[string]*Ellipse

	// insertion point for slice of pointers maps
	Ellipse_Animates_reverseMap map[*Animate]*Ellipse

	OnAfterEllipseCreateCallback OnAfterCreateInterface[Ellipse]
	OnAfterEllipseUpdateCallback OnAfterUpdateInterface[Ellipse]
	OnAfterEllipseDeleteCallback OnAfterDeleteInterface[Ellipse]
	OnAfterEllipseReadCallback   OnAfterReadInterface[Ellipse]

	Layers           map[*Layer]struct{}
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

	Lines           map[*Line]struct{}
	Lines_mapString map[string]*Line

	// insertion point for slice of pointers maps
	Line_Animates_reverseMap map[*Animate]*Line

	OnAfterLineCreateCallback OnAfterCreateInterface[Line]
	OnAfterLineUpdateCallback OnAfterUpdateInterface[Line]
	OnAfterLineDeleteCallback OnAfterDeleteInterface[Line]
	OnAfterLineReadCallback   OnAfterReadInterface[Line]

	Links           map[*Link]struct{}
	Links_mapString map[string]*Link

	// insertion point for slice of pointers maps
	Link_TextAtArrowStart_reverseMap map[*LinkAnchoredText]*Link

	Link_TextAtArrowEnd_reverseMap map[*LinkAnchoredText]*Link

	Link_ControlPoints_reverseMap map[*ControlPoint]*Link

	OnAfterLinkCreateCallback OnAfterCreateInterface[Link]
	OnAfterLinkUpdateCallback OnAfterUpdateInterface[Link]
	OnAfterLinkDeleteCallback OnAfterDeleteInterface[Link]
	OnAfterLinkReadCallback   OnAfterReadInterface[Link]

	LinkAnchoredTexts           map[*LinkAnchoredText]struct{}
	LinkAnchoredTexts_mapString map[string]*LinkAnchoredText

	// insertion point for slice of pointers maps
	LinkAnchoredText_Animates_reverseMap map[*Animate]*LinkAnchoredText

	OnAfterLinkAnchoredTextCreateCallback OnAfterCreateInterface[LinkAnchoredText]
	OnAfterLinkAnchoredTextUpdateCallback OnAfterUpdateInterface[LinkAnchoredText]
	OnAfterLinkAnchoredTextDeleteCallback OnAfterDeleteInterface[LinkAnchoredText]
	OnAfterLinkAnchoredTextReadCallback   OnAfterReadInterface[LinkAnchoredText]

	Paths           map[*Path]struct{}
	Paths_mapString map[string]*Path

	// insertion point for slice of pointers maps
	Path_Animates_reverseMap map[*Animate]*Path

	OnAfterPathCreateCallback OnAfterCreateInterface[Path]
	OnAfterPathUpdateCallback OnAfterUpdateInterface[Path]
	OnAfterPathDeleteCallback OnAfterDeleteInterface[Path]
	OnAfterPathReadCallback   OnAfterReadInterface[Path]

	Points           map[*Point]struct{}
	Points_mapString map[string]*Point

	// insertion point for slice of pointers maps
	OnAfterPointCreateCallback OnAfterCreateInterface[Point]
	OnAfterPointUpdateCallback OnAfterUpdateInterface[Point]
	OnAfterPointDeleteCallback OnAfterDeleteInterface[Point]
	OnAfterPointReadCallback   OnAfterReadInterface[Point]

	Polygones           map[*Polygone]struct{}
	Polygones_mapString map[string]*Polygone

	// insertion point for slice of pointers maps
	Polygone_Animates_reverseMap map[*Animate]*Polygone

	OnAfterPolygoneCreateCallback OnAfterCreateInterface[Polygone]
	OnAfterPolygoneUpdateCallback OnAfterUpdateInterface[Polygone]
	OnAfterPolygoneDeleteCallback OnAfterDeleteInterface[Polygone]
	OnAfterPolygoneReadCallback   OnAfterReadInterface[Polygone]

	Polylines           map[*Polyline]struct{}
	Polylines_mapString map[string]*Polyline

	// insertion point for slice of pointers maps
	Polyline_Animates_reverseMap map[*Animate]*Polyline

	OnAfterPolylineCreateCallback OnAfterCreateInterface[Polyline]
	OnAfterPolylineUpdateCallback OnAfterUpdateInterface[Polyline]
	OnAfterPolylineDeleteCallback OnAfterDeleteInterface[Polyline]
	OnAfterPolylineReadCallback   OnAfterReadInterface[Polyline]

	Rects           map[*Rect]struct{}
	Rects_mapString map[string]*Rect

	// insertion point for slice of pointers maps
	Rect_HoveringTrigger_reverseMap map[*Condition]*Rect

	Rect_DisplayConditions_reverseMap map[*Condition]*Rect

	Rect_Animations_reverseMap map[*Animate]*Rect

	Rect_RectAnchoredTexts_reverseMap map[*RectAnchoredText]*Rect

	Rect_RectAnchoredRects_reverseMap map[*RectAnchoredRect]*Rect

	Rect_RectAnchoredPaths_reverseMap map[*RectAnchoredPath]*Rect

	OnAfterRectCreateCallback OnAfterCreateInterface[Rect]
	OnAfterRectUpdateCallback OnAfterUpdateInterface[Rect]
	OnAfterRectDeleteCallback OnAfterDeleteInterface[Rect]
	OnAfterRectReadCallback   OnAfterReadInterface[Rect]

	RectAnchoredPaths           map[*RectAnchoredPath]struct{}
	RectAnchoredPaths_mapString map[string]*RectAnchoredPath

	// insertion point for slice of pointers maps
	OnAfterRectAnchoredPathCreateCallback OnAfterCreateInterface[RectAnchoredPath]
	OnAfterRectAnchoredPathUpdateCallback OnAfterUpdateInterface[RectAnchoredPath]
	OnAfterRectAnchoredPathDeleteCallback OnAfterDeleteInterface[RectAnchoredPath]
	OnAfterRectAnchoredPathReadCallback   OnAfterReadInterface[RectAnchoredPath]

	RectAnchoredRects           map[*RectAnchoredRect]struct{}
	RectAnchoredRects_mapString map[string]*RectAnchoredRect

	// insertion point for slice of pointers maps
	OnAfterRectAnchoredRectCreateCallback OnAfterCreateInterface[RectAnchoredRect]
	OnAfterRectAnchoredRectUpdateCallback OnAfterUpdateInterface[RectAnchoredRect]
	OnAfterRectAnchoredRectDeleteCallback OnAfterDeleteInterface[RectAnchoredRect]
	OnAfterRectAnchoredRectReadCallback   OnAfterReadInterface[RectAnchoredRect]

	RectAnchoredTexts           map[*RectAnchoredText]struct{}
	RectAnchoredTexts_mapString map[string]*RectAnchoredText

	// insertion point for slice of pointers maps
	RectAnchoredText_Animates_reverseMap map[*Animate]*RectAnchoredText

	OnAfterRectAnchoredTextCreateCallback OnAfterCreateInterface[RectAnchoredText]
	OnAfterRectAnchoredTextUpdateCallback OnAfterUpdateInterface[RectAnchoredText]
	OnAfterRectAnchoredTextDeleteCallback OnAfterDeleteInterface[RectAnchoredText]
	OnAfterRectAnchoredTextReadCallback   OnAfterReadInterface[RectAnchoredText]

	RectLinkLinks           map[*RectLinkLink]struct{}
	RectLinkLinks_mapString map[string]*RectLinkLink

	// insertion point for slice of pointers maps
	OnAfterRectLinkLinkCreateCallback OnAfterCreateInterface[RectLinkLink]
	OnAfterRectLinkLinkUpdateCallback OnAfterUpdateInterface[RectLinkLink]
	OnAfterRectLinkLinkDeleteCallback OnAfterDeleteInterface[RectLinkLink]
	OnAfterRectLinkLinkReadCallback   OnAfterReadInterface[RectLinkLink]

	SVGs           map[*SVG]struct{}
	SVGs_mapString map[string]*SVG

	// insertion point for slice of pointers maps
	SVG_Layers_reverseMap map[*Layer]*SVG

	OnAfterSVGCreateCallback OnAfterCreateInterface[SVG]
	OnAfterSVGUpdateCallback OnAfterUpdateInterface[SVG]
	OnAfterSVGDeleteCallback OnAfterDeleteInterface[SVG]
	OnAfterSVGReadCallback   OnAfterReadInterface[SVG]

	SvgTexts           map[*SvgText]struct{}
	SvgTexts_mapString map[string]*SvgText

	// insertion point for slice of pointers maps
	OnAfterSvgTextCreateCallback OnAfterCreateInterface[SvgText]
	OnAfterSvgTextUpdateCallback OnAfterUpdateInterface[SvgText]
	OnAfterSvgTextDeleteCallback OnAfterDeleteInterface[SvgText]
	OnAfterSvgTextReadCallback   OnAfterReadInterface[SvgText]

	Texts           map[*Text]struct{}
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

	// store the stage order of each instance in order to
	// preserve this order when serializing them
	// insertion point for order fields declaration
	AnimateOrder            uint
	AnimateMap_Staged_Order map[*Animate]uint

	CircleOrder            uint
	CircleMap_Staged_Order map[*Circle]uint

	ConditionOrder            uint
	ConditionMap_Staged_Order map[*Condition]uint

	ControlPointOrder            uint
	ControlPointMap_Staged_Order map[*ControlPoint]uint

	EllipseOrder            uint
	EllipseMap_Staged_Order map[*Ellipse]uint

	LayerOrder            uint
	LayerMap_Staged_Order map[*Layer]uint

	LineOrder            uint
	LineMap_Staged_Order map[*Line]uint

	LinkOrder            uint
	LinkMap_Staged_Order map[*Link]uint

	LinkAnchoredTextOrder            uint
	LinkAnchoredTextMap_Staged_Order map[*LinkAnchoredText]uint

	PathOrder            uint
	PathMap_Staged_Order map[*Path]uint

	PointOrder            uint
	PointMap_Staged_Order map[*Point]uint

	PolygoneOrder            uint
	PolygoneMap_Staged_Order map[*Polygone]uint

	PolylineOrder            uint
	PolylineMap_Staged_Order map[*Polyline]uint

	RectOrder            uint
	RectMap_Staged_Order map[*Rect]uint

	RectAnchoredPathOrder            uint
	RectAnchoredPathMap_Staged_Order map[*RectAnchoredPath]uint

	RectAnchoredRectOrder            uint
	RectAnchoredRectMap_Staged_Order map[*RectAnchoredRect]uint

	RectAnchoredTextOrder            uint
	RectAnchoredTextMap_Staged_Order map[*RectAnchoredText]uint

	RectLinkLinkOrder            uint
	RectLinkLinkMap_Staged_Order map[*RectLinkLink]uint

	SVGOrder            uint
	SVGMap_Staged_Order map[*SVG]uint

	SvgTextOrder            uint
	SvgTextMap_Staged_Order map[*SvgText]uint

	TextOrder            uint
	TextMap_Staged_Order map[*Text]uint

	// end of insertion point

	NamedStructs []*NamedStruct

	// for the computation of the diff at each commit we need
	reference map[GongstructIF]GongstructIF
	modified  map[GongstructIF]struct{}
	new       map[GongstructIF]struct{}
	deleted   map[GongstructIF]struct{}
}

func (stage *Stage) GetCommitId() uint {
	return stage.commitId
}

func (stage *Stage) GetCommitTS() time.Time {
	return stage.commitTimeStamp
}

// GetNamedStructs implements models.ProbebStage.
func (stage *Stage) GetNamedStructsNames() (res []string) {

	for _, namedStruct := range stage.NamedStructs {
		res = append(res, namedStruct.name)
	}

	return
}

func (stage *Stage) GetReference() map[GongstructIF]GongstructIF {
	return stage.reference
}

func (stage *Stage) GetModified() map[GongstructIF]struct{} {
	return stage.modified
}

func (stage *Stage) GetNew() map[GongstructIF]struct{} {
	return stage.new
}

func (stage *Stage) GetDeleted() map[GongstructIF]struct{} {
	return stage.deleted
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
	case *Animate:
		tmp := GetStructInstancesByOrder(stage.Animates, stage.AnimateMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Animate implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Circle:
		tmp := GetStructInstancesByOrder(stage.Circles, stage.CircleMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Circle implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Condition:
		tmp := GetStructInstancesByOrder(stage.Conditions, stage.ConditionMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Condition implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ControlPoint:
		tmp := GetStructInstancesByOrder(stage.ControlPoints, stage.ControlPointMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ControlPoint implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Ellipse:
		tmp := GetStructInstancesByOrder(stage.Ellipses, stage.EllipseMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Ellipse implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Layer:
		tmp := GetStructInstancesByOrder(stage.Layers, stage.LayerMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Layer implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Line:
		tmp := GetStructInstancesByOrder(stage.Lines, stage.LineMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Line implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Link:
		tmp := GetStructInstancesByOrder(stage.Links, stage.LinkMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Link implements.
			res = append(res, any(v).(T))
		}
		return res
	case *LinkAnchoredText:
		tmp := GetStructInstancesByOrder(stage.LinkAnchoredTexts, stage.LinkAnchoredTextMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *LinkAnchoredText implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Path:
		tmp := GetStructInstancesByOrder(stage.Paths, stage.PathMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Path implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Point:
		tmp := GetStructInstancesByOrder(stage.Points, stage.PointMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Point implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Polygone:
		tmp := GetStructInstancesByOrder(stage.Polygones, stage.PolygoneMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Polygone implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Polyline:
		tmp := GetStructInstancesByOrder(stage.Polylines, stage.PolylineMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Polyline implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Rect:
		tmp := GetStructInstancesByOrder(stage.Rects, stage.RectMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Rect implements.
			res = append(res, any(v).(T))
		}
		return res
	case *RectAnchoredPath:
		tmp := GetStructInstancesByOrder(stage.RectAnchoredPaths, stage.RectAnchoredPathMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *RectAnchoredPath implements.
			res = append(res, any(v).(T))
		}
		return res
	case *RectAnchoredRect:
		tmp := GetStructInstancesByOrder(stage.RectAnchoredRects, stage.RectAnchoredRectMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *RectAnchoredRect implements.
			res = append(res, any(v).(T))
		}
		return res
	case *RectAnchoredText:
		tmp := GetStructInstancesByOrder(stage.RectAnchoredTexts, stage.RectAnchoredTextMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *RectAnchoredText implements.
			res = append(res, any(v).(T))
		}
		return res
	case *RectLinkLink:
		tmp := GetStructInstancesByOrder(stage.RectLinkLinks, stage.RectLinkLinkMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *RectLinkLink implements.
			res = append(res, any(v).(T))
		}
		return res
	case *SVG:
		tmp := GetStructInstancesByOrder(stage.SVGs, stage.SVGMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SVG implements.
			res = append(res, any(v).(T))
		}
		return res
	case *SvgText:
		tmp := GetStructInstancesByOrder(stage.SvgTexts, stage.SvgTextMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SvgText implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Text:
		tmp := GetStructInstancesByOrder(stage.Texts, stage.TextMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Text implements.
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
	case "Animate":
		res = GetNamedStructInstances(stage.Animates, stage.AnimateMap_Staged_Order)
	case "Circle":
		res = GetNamedStructInstances(stage.Circles, stage.CircleMap_Staged_Order)
	case "Condition":
		res = GetNamedStructInstances(stage.Conditions, stage.ConditionMap_Staged_Order)
	case "ControlPoint":
		res = GetNamedStructInstances(stage.ControlPoints, stage.ControlPointMap_Staged_Order)
	case "Ellipse":
		res = GetNamedStructInstances(stage.Ellipses, stage.EllipseMap_Staged_Order)
	case "Layer":
		res = GetNamedStructInstances(stage.Layers, stage.LayerMap_Staged_Order)
	case "Line":
		res = GetNamedStructInstances(stage.Lines, stage.LineMap_Staged_Order)
	case "Link":
		res = GetNamedStructInstances(stage.Links, stage.LinkMap_Staged_Order)
	case "LinkAnchoredText":
		res = GetNamedStructInstances(stage.LinkAnchoredTexts, stage.LinkAnchoredTextMap_Staged_Order)
	case "Path":
		res = GetNamedStructInstances(stage.Paths, stage.PathMap_Staged_Order)
	case "Point":
		res = GetNamedStructInstances(stage.Points, stage.PointMap_Staged_Order)
	case "Polygone":
		res = GetNamedStructInstances(stage.Polygones, stage.PolygoneMap_Staged_Order)
	case "Polyline":
		res = GetNamedStructInstances(stage.Polylines, stage.PolylineMap_Staged_Order)
	case "Rect":
		res = GetNamedStructInstances(stage.Rects, stage.RectMap_Staged_Order)
	case "RectAnchoredPath":
		res = GetNamedStructInstances(stage.RectAnchoredPaths, stage.RectAnchoredPathMap_Staged_Order)
	case "RectAnchoredRect":
		res = GetNamedStructInstances(stage.RectAnchoredRects, stage.RectAnchoredRectMap_Staged_Order)
	case "RectAnchoredText":
		res = GetNamedStructInstances(stage.RectAnchoredTexts, stage.RectAnchoredTextMap_Staged_Order)
	case "RectLinkLink":
		res = GetNamedStructInstances(stage.RectLinkLinks, stage.RectLinkLinkMap_Staged_Order)
	case "SVG":
		res = GetNamedStructInstances(stage.SVGs, stage.SVGMap_Staged_Order)
	case "SvgText":
		res = GetNamedStructInstances(stage.SvgTexts, stage.SvgTextMap_Staged_Order)
	case "Text":
		res = GetNamedStructInstances(stage.Texts, stage.TextMap_Staged_Order)
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
	return "github.com/fullstack-lang/gong/lib/svg/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return svg_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return svg_go.GoDiagramsDir
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
	CommitAnimate(animate *Animate)
	CheckoutAnimate(animate *Animate)
	CommitCircle(circle *Circle)
	CheckoutCircle(circle *Circle)
	CommitCondition(condition *Condition)
	CheckoutCondition(condition *Condition)
	CommitControlPoint(controlpoint *ControlPoint)
	CheckoutControlPoint(controlpoint *ControlPoint)
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
	CommitSvgText(svgtext *SvgText)
	CheckoutSvgText(svgtext *SvgText)
	CommitText(text *Text)
	CheckoutText(text *Text)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

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

		Layers:           make(map[*Layer]struct{}),
		Layers_mapString: make(map[string]*Layer),

		Lines:           make(map[*Line]struct{}),
		Lines_mapString: make(map[string]*Line),

		Links:           make(map[*Link]struct{}),
		Links_mapString: make(map[string]*Link),

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
		AnimateMap_Staged_Order: make(map[*Animate]uint),

		CircleMap_Staged_Order: make(map[*Circle]uint),

		ConditionMap_Staged_Order: make(map[*Condition]uint),

		ControlPointMap_Staged_Order: make(map[*ControlPoint]uint),

		EllipseMap_Staged_Order: make(map[*Ellipse]uint),

		LayerMap_Staged_Order: make(map[*Layer]uint),

		LineMap_Staged_Order: make(map[*Line]uint),

		LinkMap_Staged_Order: make(map[*Link]uint),

		LinkAnchoredTextMap_Staged_Order: make(map[*LinkAnchoredText]uint),

		PathMap_Staged_Order: make(map[*Path]uint),

		PointMap_Staged_Order: make(map[*Point]uint),

		PolygoneMap_Staged_Order: make(map[*Polygone]uint),

		PolylineMap_Staged_Order: make(map[*Polyline]uint),

		RectMap_Staged_Order: make(map[*Rect]uint),

		RectAnchoredPathMap_Staged_Order: make(map[*RectAnchoredPath]uint),

		RectAnchoredRectMap_Staged_Order: make(map[*RectAnchoredRect]uint),

		RectAnchoredTextMap_Staged_Order: make(map[*RectAnchoredText]uint),

		RectLinkLinkMap_Staged_Order: make(map[*RectLinkLink]uint),

		SVGMap_Staged_Order: make(map[*SVG]uint),

		SvgTextMap_Staged_Order: make(map[*SvgText]uint),

		TextMap_Staged_Order: make(map[*Text]uint),

		// end of insertion point

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Animate"},
			{name: "Circle"},
			{name: "Condition"},
			{name: "ControlPoint"},
			{name: "Ellipse"},
			{name: "Layer"},
			{name: "Line"},
			{name: "Link"},
			{name: "LinkAnchoredText"},
			{name: "Path"},
			{name: "Point"},
			{name: "Polygone"},
			{name: "Polyline"},
			{name: "Rect"},
			{name: "RectAnchoredPath"},
			{name: "RectAnchoredRect"},
			{name: "RectAnchoredText"},
			{name: "RectLinkLink"},
			{name: "SVG"},
			{name: "SvgText"},
			{name: "Text"},
		}, // end of insertion point

		reference: make(map[GongstructIF]GongstructIF),
		new:       make(map[GongstructIF]struct{}),
		modified:  make(map[GongstructIF]struct{}),
		deleted:   make(map[GongstructIF]struct{}),
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Animate:
		return stage.AnimateMap_Staged_Order[instance]
	case *Circle:
		return stage.CircleMap_Staged_Order[instance]
	case *Condition:
		return stage.ConditionMap_Staged_Order[instance]
	case *ControlPoint:
		return stage.ControlPointMap_Staged_Order[instance]
	case *Ellipse:
		return stage.EllipseMap_Staged_Order[instance]
	case *Layer:
		return stage.LayerMap_Staged_Order[instance]
	case *Line:
		return stage.LineMap_Staged_Order[instance]
	case *Link:
		return stage.LinkMap_Staged_Order[instance]
	case *LinkAnchoredText:
		return stage.LinkAnchoredTextMap_Staged_Order[instance]
	case *Path:
		return stage.PathMap_Staged_Order[instance]
	case *Point:
		return stage.PointMap_Staged_Order[instance]
	case *Polygone:
		return stage.PolygoneMap_Staged_Order[instance]
	case *Polyline:
		return stage.PolylineMap_Staged_Order[instance]
	case *Rect:
		return stage.RectMap_Staged_Order[instance]
	case *RectAnchoredPath:
		return stage.RectAnchoredPathMap_Staged_Order[instance]
	case *RectAnchoredRect:
		return stage.RectAnchoredRectMap_Staged_Order[instance]
	case *RectAnchoredText:
		return stage.RectAnchoredTextMap_Staged_Order[instance]
	case *RectLinkLink:
		return stage.RectLinkLinkMap_Staged_Order[instance]
	case *SVG:
		return stage.SVGMap_Staged_Order[instance]
	case *SvgText:
		return stage.SvgTextMap_Staged_Order[instance]
	case *Text:
		return stage.TextMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Animate:
		return stage.AnimateMap_Staged_Order[instance]
	case *Circle:
		return stage.CircleMap_Staged_Order[instance]
	case *Condition:
		return stage.ConditionMap_Staged_Order[instance]
	case *ControlPoint:
		return stage.ControlPointMap_Staged_Order[instance]
	case *Ellipse:
		return stage.EllipseMap_Staged_Order[instance]
	case *Layer:
		return stage.LayerMap_Staged_Order[instance]
	case *Line:
		return stage.LineMap_Staged_Order[instance]
	case *Link:
		return stage.LinkMap_Staged_Order[instance]
	case *LinkAnchoredText:
		return stage.LinkAnchoredTextMap_Staged_Order[instance]
	case *Path:
		return stage.PathMap_Staged_Order[instance]
	case *Point:
		return stage.PointMap_Staged_Order[instance]
	case *Polygone:
		return stage.PolygoneMap_Staged_Order[instance]
	case *Polyline:
		return stage.PolylineMap_Staged_Order[instance]
	case *Rect:
		return stage.RectMap_Staged_Order[instance]
	case *RectAnchoredPath:
		return stage.RectAnchoredPathMap_Staged_Order[instance]
	case *RectAnchoredRect:
		return stage.RectAnchoredRectMap_Staged_Order[instance]
	case *RectAnchoredText:
		return stage.RectAnchoredTextMap_Staged_Order[instance]
	case *RectLinkLink:
		return stage.RectLinkLinkMap_Staged_Order[instance]
	case *SVG:
		return stage.SVGMap_Staged_Order[instance]
	case *SvgText:
		return stage.SvgTextMap_Staged_Order[instance]
	case *Text:
		return stage.TextMap_Staged_Order[instance]
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
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *Stage) Commit() {
	stage.ComputeReverseMaps()
	stage.commitId++
	stage.commitTimeStamp = time.Now()

	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}
	stage.ComputeInstancesNb()
	stage.ComputeReference()
}

func (stage *Stage) ComputeInstancesNb() {
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Animate"] = len(stage.Animates)
	stage.Map_GongStructName_InstancesNb["Circle"] = len(stage.Circles)
	stage.Map_GongStructName_InstancesNb["Condition"] = len(stage.Conditions)
	stage.Map_GongStructName_InstancesNb["ControlPoint"] = len(stage.ControlPoints)
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

	if _, ok := stage.Animates[animate]; !ok {
		stage.Animates[animate] = struct{}{}
		stage.AnimateMap_Staged_Order[animate] = stage.AnimateOrder
		stage.AnimateOrder++
		stage.new[animate] = struct{}{}
		delete(stage.deleted, animate)
	} else {
		if _, ok := stage.new[animate]; !ok {
			stage.modified[animate] = struct{}{}
		}
	}
	stage.Animates_mapString[animate.Name] = animate

	return animate
}

// Unstage removes animate off the model stage
func (animate *Animate) Unstage(stage *Stage) *Animate {
	delete(stage.Animates, animate)
	delete(stage.Animates_mapString, animate.Name)

	if _, ok := stage.reference[animate]; ok {
		stage.deleted[animate] = struct{}{}
	} else {
		delete(stage.new, animate)
	}
	return animate
}

// UnstageVoid removes animate off the model stage
func (animate *Animate) UnstageVoid(stage *Stage) {
	delete(stage.Animates, animate)
	delete(stage.Animates_mapString, animate.Name)
}

// commit animate to the back repo (if it is already staged)
func (animate *Animate) Commit(stage *Stage) *Animate {
	if _, ok := stage.Animates[animate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAnimate(animate)
		}
	}
	return animate
}

func (animate *Animate) CommitVoid(stage *Stage) {
	animate.Commit(stage)
}

func (animate *Animate) StageVoid(stage *Stage) {
	animate.Stage(stage)
}

// Checkout animate to the back repo (if it is already staged)
func (animate *Animate) Checkout(stage *Stage) *Animate {
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
func (circle *Circle) Stage(stage *Stage) *Circle {

	if _, ok := stage.Circles[circle]; !ok {
		stage.Circles[circle] = struct{}{}
		stage.CircleMap_Staged_Order[circle] = stage.CircleOrder
		stage.CircleOrder++
		stage.new[circle] = struct{}{}
		delete(stage.deleted, circle)
	} else {
		if _, ok := stage.new[circle]; !ok {
			stage.modified[circle] = struct{}{}
		}
	}
	stage.Circles_mapString[circle.Name] = circle

	return circle
}

// Unstage removes circle off the model stage
func (circle *Circle) Unstage(stage *Stage) *Circle {
	delete(stage.Circles, circle)
	delete(stage.Circles_mapString, circle.Name)

	if _, ok := stage.reference[circle]; ok {
		stage.deleted[circle] = struct{}{}
	} else {
		delete(stage.new, circle)
	}
	return circle
}

// UnstageVoid removes circle off the model stage
func (circle *Circle) UnstageVoid(stage *Stage) {
	delete(stage.Circles, circle)
	delete(stage.Circles_mapString, circle.Name)
}

// commit circle to the back repo (if it is already staged)
func (circle *Circle) Commit(stage *Stage) *Circle {
	if _, ok := stage.Circles[circle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCircle(circle)
		}
	}
	return circle
}

func (circle *Circle) CommitVoid(stage *Stage) {
	circle.Commit(stage)
}

func (circle *Circle) StageVoid(stage *Stage) {
	circle.Stage(stage)
}

// Checkout circle to the back repo (if it is already staged)
func (circle *Circle) Checkout(stage *Stage) *Circle {
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

// Stage puts condition to the model stage
func (condition *Condition) Stage(stage *Stage) *Condition {

	if _, ok := stage.Conditions[condition]; !ok {
		stage.Conditions[condition] = struct{}{}
		stage.ConditionMap_Staged_Order[condition] = stage.ConditionOrder
		stage.ConditionOrder++
		stage.new[condition] = struct{}{}
		delete(stage.deleted, condition)
	} else {
		if _, ok := stage.new[condition]; !ok {
			stage.modified[condition] = struct{}{}
		}
	}
	stage.Conditions_mapString[condition.Name] = condition

	return condition
}

// Unstage removes condition off the model stage
func (condition *Condition) Unstage(stage *Stage) *Condition {
	delete(stage.Conditions, condition)
	delete(stage.Conditions_mapString, condition.Name)

	if _, ok := stage.reference[condition]; ok {
		stage.deleted[condition] = struct{}{}
	} else {
		delete(stage.new, condition)
	}
	return condition
}

// UnstageVoid removes condition off the model stage
func (condition *Condition) UnstageVoid(stage *Stage) {
	delete(stage.Conditions, condition)
	delete(stage.Conditions_mapString, condition.Name)
}

// commit condition to the back repo (if it is already staged)
func (condition *Condition) Commit(stage *Stage) *Condition {
	if _, ok := stage.Conditions[condition]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCondition(condition)
		}
	}
	return condition
}

func (condition *Condition) CommitVoid(stage *Stage) {
	condition.Commit(stage)
}

func (condition *Condition) StageVoid(stage *Stage) {
	condition.Stage(stage)
}

// Checkout condition to the back repo (if it is already staged)
func (condition *Condition) Checkout(stage *Stage) *Condition {
	if _, ok := stage.Conditions[condition]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCondition(condition)
		}
	}
	return condition
}

// for satisfaction of GongStruct interface
func (condition *Condition) GetName() (res string) {
	return condition.Name
}

// Stage puts controlpoint to the model stage
func (controlpoint *ControlPoint) Stage(stage *Stage) *ControlPoint {

	if _, ok := stage.ControlPoints[controlpoint]; !ok {
		stage.ControlPoints[controlpoint] = struct{}{}
		stage.ControlPointMap_Staged_Order[controlpoint] = stage.ControlPointOrder
		stage.ControlPointOrder++
		stage.new[controlpoint] = struct{}{}
		delete(stage.deleted, controlpoint)
	} else {
		if _, ok := stage.new[controlpoint]; !ok {
			stage.modified[controlpoint] = struct{}{}
		}
	}
	stage.ControlPoints_mapString[controlpoint.Name] = controlpoint

	return controlpoint
}

// Unstage removes controlpoint off the model stage
func (controlpoint *ControlPoint) Unstage(stage *Stage) *ControlPoint {
	delete(stage.ControlPoints, controlpoint)
	delete(stage.ControlPoints_mapString, controlpoint.Name)

	if _, ok := stage.reference[controlpoint]; ok {
		stage.deleted[controlpoint] = struct{}{}
	} else {
		delete(stage.new, controlpoint)
	}
	return controlpoint
}

// UnstageVoid removes controlpoint off the model stage
func (controlpoint *ControlPoint) UnstageVoid(stage *Stage) {
	delete(stage.ControlPoints, controlpoint)
	delete(stage.ControlPoints_mapString, controlpoint.Name)
}

// commit controlpoint to the back repo (if it is already staged)
func (controlpoint *ControlPoint) Commit(stage *Stage) *ControlPoint {
	if _, ok := stage.ControlPoints[controlpoint]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitControlPoint(controlpoint)
		}
	}
	return controlpoint
}

func (controlpoint *ControlPoint) CommitVoid(stage *Stage) {
	controlpoint.Commit(stage)
}

func (controlpoint *ControlPoint) StageVoid(stage *Stage) {
	controlpoint.Stage(stage)
}

// Checkout controlpoint to the back repo (if it is already staged)
func (controlpoint *ControlPoint) Checkout(stage *Stage) *ControlPoint {
	if _, ok := stage.ControlPoints[controlpoint]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutControlPoint(controlpoint)
		}
	}
	return controlpoint
}

// for satisfaction of GongStruct interface
func (controlpoint *ControlPoint) GetName() (res string) {
	return controlpoint.Name
}

// Stage puts ellipse to the model stage
func (ellipse *Ellipse) Stage(stage *Stage) *Ellipse {

	if _, ok := stage.Ellipses[ellipse]; !ok {
		stage.Ellipses[ellipse] = struct{}{}
		stage.EllipseMap_Staged_Order[ellipse] = stage.EllipseOrder
		stage.EllipseOrder++
		stage.new[ellipse] = struct{}{}
		delete(stage.deleted, ellipse)
	} else {
		if _, ok := stage.new[ellipse]; !ok {
			stage.modified[ellipse] = struct{}{}
		}
	}
	stage.Ellipses_mapString[ellipse.Name] = ellipse

	return ellipse
}

// Unstage removes ellipse off the model stage
func (ellipse *Ellipse) Unstage(stage *Stage) *Ellipse {
	delete(stage.Ellipses, ellipse)
	delete(stage.Ellipses_mapString, ellipse.Name)

	if _, ok := stage.reference[ellipse]; ok {
		stage.deleted[ellipse] = struct{}{}
	} else {
		delete(stage.new, ellipse)
	}
	return ellipse
}

// UnstageVoid removes ellipse off the model stage
func (ellipse *Ellipse) UnstageVoid(stage *Stage) {
	delete(stage.Ellipses, ellipse)
	delete(stage.Ellipses_mapString, ellipse.Name)
}

// commit ellipse to the back repo (if it is already staged)
func (ellipse *Ellipse) Commit(stage *Stage) *Ellipse {
	if _, ok := stage.Ellipses[ellipse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitEllipse(ellipse)
		}
	}
	return ellipse
}

func (ellipse *Ellipse) CommitVoid(stage *Stage) {
	ellipse.Commit(stage)
}

func (ellipse *Ellipse) StageVoid(stage *Stage) {
	ellipse.Stage(stage)
}

// Checkout ellipse to the back repo (if it is already staged)
func (ellipse *Ellipse) Checkout(stage *Stage) *Ellipse {
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
func (layer *Layer) Stage(stage *Stage) *Layer {

	if _, ok := stage.Layers[layer]; !ok {
		stage.Layers[layer] = struct{}{}
		stage.LayerMap_Staged_Order[layer] = stage.LayerOrder
		stage.LayerOrder++
		stage.new[layer] = struct{}{}
		delete(stage.deleted, layer)
	} else {
		if _, ok := stage.new[layer]; !ok {
			stage.modified[layer] = struct{}{}
		}
	}
	stage.Layers_mapString[layer.Name] = layer

	return layer
}

// Unstage removes layer off the model stage
func (layer *Layer) Unstage(stage *Stage) *Layer {
	delete(stage.Layers, layer)
	delete(stage.Layers_mapString, layer.Name)

	if _, ok := stage.reference[layer]; ok {
		stage.deleted[layer] = struct{}{}
	} else {
		delete(stage.new, layer)
	}
	return layer
}

// UnstageVoid removes layer off the model stage
func (layer *Layer) UnstageVoid(stage *Stage) {
	delete(stage.Layers, layer)
	delete(stage.Layers_mapString, layer.Name)
}

// commit layer to the back repo (if it is already staged)
func (layer *Layer) Commit(stage *Stage) *Layer {
	if _, ok := stage.Layers[layer]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLayer(layer)
		}
	}
	return layer
}

func (layer *Layer) CommitVoid(stage *Stage) {
	layer.Commit(stage)
}

func (layer *Layer) StageVoid(stage *Stage) {
	layer.Stage(stage)
}

// Checkout layer to the back repo (if it is already staged)
func (layer *Layer) Checkout(stage *Stage) *Layer {
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
func (line *Line) Stage(stage *Stage) *Line {

	if _, ok := stage.Lines[line]; !ok {
		stage.Lines[line] = struct{}{}
		stage.LineMap_Staged_Order[line] = stage.LineOrder
		stage.LineOrder++
		stage.new[line] = struct{}{}
		delete(stage.deleted, line)
	} else {
		if _, ok := stage.new[line]; !ok {
			stage.modified[line] = struct{}{}
		}
	}
	stage.Lines_mapString[line.Name] = line

	return line
}

// Unstage removes line off the model stage
func (line *Line) Unstage(stage *Stage) *Line {
	delete(stage.Lines, line)
	delete(stage.Lines_mapString, line.Name)

	if _, ok := stage.reference[line]; ok {
		stage.deleted[line] = struct{}{}
	} else {
		delete(stage.new, line)
	}
	return line
}

// UnstageVoid removes line off the model stage
func (line *Line) UnstageVoid(stage *Stage) {
	delete(stage.Lines, line)
	delete(stage.Lines_mapString, line.Name)
}

// commit line to the back repo (if it is already staged)
func (line *Line) Commit(stage *Stage) *Line {
	if _, ok := stage.Lines[line]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLine(line)
		}
	}
	return line
}

func (line *Line) CommitVoid(stage *Stage) {
	line.Commit(stage)
}

func (line *Line) StageVoid(stage *Stage) {
	line.Stage(stage)
}

// Checkout line to the back repo (if it is already staged)
func (line *Line) Checkout(stage *Stage) *Line {
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
func (link *Link) Stage(stage *Stage) *Link {

	if _, ok := stage.Links[link]; !ok {
		stage.Links[link] = struct{}{}
		stage.LinkMap_Staged_Order[link] = stage.LinkOrder
		stage.LinkOrder++
		stage.new[link] = struct{}{}
		delete(stage.deleted, link)
	} else {
		if _, ok := stage.new[link]; !ok {
			stage.modified[link] = struct{}{}
		}
	}
	stage.Links_mapString[link.Name] = link

	return link
}

// Unstage removes link off the model stage
func (link *Link) Unstage(stage *Stage) *Link {
	delete(stage.Links, link)
	delete(stage.Links_mapString, link.Name)

	if _, ok := stage.reference[link]; ok {
		stage.deleted[link] = struct{}{}
	} else {
		delete(stage.new, link)
	}
	return link
}

// UnstageVoid removes link off the model stage
func (link *Link) UnstageVoid(stage *Stage) {
	delete(stage.Links, link)
	delete(stage.Links_mapString, link.Name)
}

// commit link to the back repo (if it is already staged)
func (link *Link) Commit(stage *Stage) *Link {
	if _, ok := stage.Links[link]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLink(link)
		}
	}
	return link
}

func (link *Link) CommitVoid(stage *Stage) {
	link.Commit(stage)
}

func (link *Link) StageVoid(stage *Stage) {
	link.Stage(stage)
}

// Checkout link to the back repo (if it is already staged)
func (link *Link) Checkout(stage *Stage) *Link {
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
func (linkanchoredtext *LinkAnchoredText) Stage(stage *Stage) *LinkAnchoredText {

	if _, ok := stage.LinkAnchoredTexts[linkanchoredtext]; !ok {
		stage.LinkAnchoredTexts[linkanchoredtext] = struct{}{}
		stage.LinkAnchoredTextMap_Staged_Order[linkanchoredtext] = stage.LinkAnchoredTextOrder
		stage.LinkAnchoredTextOrder++
		stage.new[linkanchoredtext] = struct{}{}
		delete(stage.deleted, linkanchoredtext)
	} else {
		if _, ok := stage.new[linkanchoredtext]; !ok {
			stage.modified[linkanchoredtext] = struct{}{}
		}
	}
	stage.LinkAnchoredTexts_mapString[linkanchoredtext.Name] = linkanchoredtext

	return linkanchoredtext
}

// Unstage removes linkanchoredtext off the model stage
func (linkanchoredtext *LinkAnchoredText) Unstage(stage *Stage) *LinkAnchoredText {
	delete(stage.LinkAnchoredTexts, linkanchoredtext)
	delete(stage.LinkAnchoredTexts_mapString, linkanchoredtext.Name)

	if _, ok := stage.reference[linkanchoredtext]; ok {
		stage.deleted[linkanchoredtext] = struct{}{}
	} else {
		delete(stage.new, linkanchoredtext)
	}
	return linkanchoredtext
}

// UnstageVoid removes linkanchoredtext off the model stage
func (linkanchoredtext *LinkAnchoredText) UnstageVoid(stage *Stage) {
	delete(stage.LinkAnchoredTexts, linkanchoredtext)
	delete(stage.LinkAnchoredTexts_mapString, linkanchoredtext.Name)
}

// commit linkanchoredtext to the back repo (if it is already staged)
func (linkanchoredtext *LinkAnchoredText) Commit(stage *Stage) *LinkAnchoredText {
	if _, ok := stage.LinkAnchoredTexts[linkanchoredtext]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLinkAnchoredText(linkanchoredtext)
		}
	}
	return linkanchoredtext
}

func (linkanchoredtext *LinkAnchoredText) CommitVoid(stage *Stage) {
	linkanchoredtext.Commit(stage)
}

func (linkanchoredtext *LinkAnchoredText) StageVoid(stage *Stage) {
	linkanchoredtext.Stage(stage)
}

// Checkout linkanchoredtext to the back repo (if it is already staged)
func (linkanchoredtext *LinkAnchoredText) Checkout(stage *Stage) *LinkAnchoredText {
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
func (path *Path) Stage(stage *Stage) *Path {

	if _, ok := stage.Paths[path]; !ok {
		stage.Paths[path] = struct{}{}
		stage.PathMap_Staged_Order[path] = stage.PathOrder
		stage.PathOrder++
		stage.new[path] = struct{}{}
		delete(stage.deleted, path)
	} else {
		if _, ok := stage.new[path]; !ok {
			stage.modified[path] = struct{}{}
		}
	}
	stage.Paths_mapString[path.Name] = path

	return path
}

// Unstage removes path off the model stage
func (path *Path) Unstage(stage *Stage) *Path {
	delete(stage.Paths, path)
	delete(stage.Paths_mapString, path.Name)

	if _, ok := stage.reference[path]; ok {
		stage.deleted[path] = struct{}{}
	} else {
		delete(stage.new, path)
	}
	return path
}

// UnstageVoid removes path off the model stage
func (path *Path) UnstageVoid(stage *Stage) {
	delete(stage.Paths, path)
	delete(stage.Paths_mapString, path.Name)
}

// commit path to the back repo (if it is already staged)
func (path *Path) Commit(stage *Stage) *Path {
	if _, ok := stage.Paths[path]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPath(path)
		}
	}
	return path
}

func (path *Path) CommitVoid(stage *Stage) {
	path.Commit(stage)
}

func (path *Path) StageVoid(stage *Stage) {
	path.Stage(stage)
}

// Checkout path to the back repo (if it is already staged)
func (path *Path) Checkout(stage *Stage) *Path {
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
func (point *Point) Stage(stage *Stage) *Point {

	if _, ok := stage.Points[point]; !ok {
		stage.Points[point] = struct{}{}
		stage.PointMap_Staged_Order[point] = stage.PointOrder
		stage.PointOrder++
		stage.new[point] = struct{}{}
		delete(stage.deleted, point)
	} else {
		if _, ok := stage.new[point]; !ok {
			stage.modified[point] = struct{}{}
		}
	}
	stage.Points_mapString[point.Name] = point

	return point
}

// Unstage removes point off the model stage
func (point *Point) Unstage(stage *Stage) *Point {
	delete(stage.Points, point)
	delete(stage.Points_mapString, point.Name)

	if _, ok := stage.reference[point]; ok {
		stage.deleted[point] = struct{}{}
	} else {
		delete(stage.new, point)
	}
	return point
}

// UnstageVoid removes point off the model stage
func (point *Point) UnstageVoid(stage *Stage) {
	delete(stage.Points, point)
	delete(stage.Points_mapString, point.Name)
}

// commit point to the back repo (if it is already staged)
func (point *Point) Commit(stage *Stage) *Point {
	if _, ok := stage.Points[point]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPoint(point)
		}
	}
	return point
}

func (point *Point) CommitVoid(stage *Stage) {
	point.Commit(stage)
}

func (point *Point) StageVoid(stage *Stage) {
	point.Stage(stage)
}

// Checkout point to the back repo (if it is already staged)
func (point *Point) Checkout(stage *Stage) *Point {
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
func (polygone *Polygone) Stage(stage *Stage) *Polygone {

	if _, ok := stage.Polygones[polygone]; !ok {
		stage.Polygones[polygone] = struct{}{}
		stage.PolygoneMap_Staged_Order[polygone] = stage.PolygoneOrder
		stage.PolygoneOrder++
		stage.new[polygone] = struct{}{}
		delete(stage.deleted, polygone)
	} else {
		if _, ok := stage.new[polygone]; !ok {
			stage.modified[polygone] = struct{}{}
		}
	}
	stage.Polygones_mapString[polygone.Name] = polygone

	return polygone
}

// Unstage removes polygone off the model stage
func (polygone *Polygone) Unstage(stage *Stage) *Polygone {
	delete(stage.Polygones, polygone)
	delete(stage.Polygones_mapString, polygone.Name)

	if _, ok := stage.reference[polygone]; ok {
		stage.deleted[polygone] = struct{}{}
	} else {
		delete(stage.new, polygone)
	}
	return polygone
}

// UnstageVoid removes polygone off the model stage
func (polygone *Polygone) UnstageVoid(stage *Stage) {
	delete(stage.Polygones, polygone)
	delete(stage.Polygones_mapString, polygone.Name)
}

// commit polygone to the back repo (if it is already staged)
func (polygone *Polygone) Commit(stage *Stage) *Polygone {
	if _, ok := stage.Polygones[polygone]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPolygone(polygone)
		}
	}
	return polygone
}

func (polygone *Polygone) CommitVoid(stage *Stage) {
	polygone.Commit(stage)
}

func (polygone *Polygone) StageVoid(stage *Stage) {
	polygone.Stage(stage)
}

// Checkout polygone to the back repo (if it is already staged)
func (polygone *Polygone) Checkout(stage *Stage) *Polygone {
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
func (polyline *Polyline) Stage(stage *Stage) *Polyline {

	if _, ok := stage.Polylines[polyline]; !ok {
		stage.Polylines[polyline] = struct{}{}
		stage.PolylineMap_Staged_Order[polyline] = stage.PolylineOrder
		stage.PolylineOrder++
		stage.new[polyline] = struct{}{}
		delete(stage.deleted, polyline)
	} else {
		if _, ok := stage.new[polyline]; !ok {
			stage.modified[polyline] = struct{}{}
		}
	}
	stage.Polylines_mapString[polyline.Name] = polyline

	return polyline
}

// Unstage removes polyline off the model stage
func (polyline *Polyline) Unstage(stage *Stage) *Polyline {
	delete(stage.Polylines, polyline)
	delete(stage.Polylines_mapString, polyline.Name)

	if _, ok := stage.reference[polyline]; ok {
		stage.deleted[polyline] = struct{}{}
	} else {
		delete(stage.new, polyline)
	}
	return polyline
}

// UnstageVoid removes polyline off the model stage
func (polyline *Polyline) UnstageVoid(stage *Stage) {
	delete(stage.Polylines, polyline)
	delete(stage.Polylines_mapString, polyline.Name)
}

// commit polyline to the back repo (if it is already staged)
func (polyline *Polyline) Commit(stage *Stage) *Polyline {
	if _, ok := stage.Polylines[polyline]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPolyline(polyline)
		}
	}
	return polyline
}

func (polyline *Polyline) CommitVoid(stage *Stage) {
	polyline.Commit(stage)
}

func (polyline *Polyline) StageVoid(stage *Stage) {
	polyline.Stage(stage)
}

// Checkout polyline to the back repo (if it is already staged)
func (polyline *Polyline) Checkout(stage *Stage) *Polyline {
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
func (rect *Rect) Stage(stage *Stage) *Rect {

	if _, ok := stage.Rects[rect]; !ok {
		stage.Rects[rect] = struct{}{}
		stage.RectMap_Staged_Order[rect] = stage.RectOrder
		stage.RectOrder++
		stage.new[rect] = struct{}{}
		delete(stage.deleted, rect)
	} else {
		if _, ok := stage.new[rect]; !ok {
			stage.modified[rect] = struct{}{}
		}
	}
	stage.Rects_mapString[rect.Name] = rect

	return rect
}

// Unstage removes rect off the model stage
func (rect *Rect) Unstage(stage *Stage) *Rect {
	delete(stage.Rects, rect)
	delete(stage.Rects_mapString, rect.Name)

	if _, ok := stage.reference[rect]; ok {
		stage.deleted[rect] = struct{}{}
	} else {
		delete(stage.new, rect)
	}
	return rect
}

// UnstageVoid removes rect off the model stage
func (rect *Rect) UnstageVoid(stage *Stage) {
	delete(stage.Rects, rect)
	delete(stage.Rects_mapString, rect.Name)
}

// commit rect to the back repo (if it is already staged)
func (rect *Rect) Commit(stage *Stage) *Rect {
	if _, ok := stage.Rects[rect]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRect(rect)
		}
	}
	return rect
}

func (rect *Rect) CommitVoid(stage *Stage) {
	rect.Commit(stage)
}

func (rect *Rect) StageVoid(stage *Stage) {
	rect.Stage(stage)
}

// Checkout rect to the back repo (if it is already staged)
func (rect *Rect) Checkout(stage *Stage) *Rect {
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
func (rectanchoredpath *RectAnchoredPath) Stage(stage *Stage) *RectAnchoredPath {

	if _, ok := stage.RectAnchoredPaths[rectanchoredpath]; !ok {
		stage.RectAnchoredPaths[rectanchoredpath] = struct{}{}
		stage.RectAnchoredPathMap_Staged_Order[rectanchoredpath] = stage.RectAnchoredPathOrder
		stage.RectAnchoredPathOrder++
		stage.new[rectanchoredpath] = struct{}{}
		delete(stage.deleted, rectanchoredpath)
	} else {
		if _, ok := stage.new[rectanchoredpath]; !ok {
			stage.modified[rectanchoredpath] = struct{}{}
		}
	}
	stage.RectAnchoredPaths_mapString[rectanchoredpath.Name] = rectanchoredpath

	return rectanchoredpath
}

// Unstage removes rectanchoredpath off the model stage
func (rectanchoredpath *RectAnchoredPath) Unstage(stage *Stage) *RectAnchoredPath {
	delete(stage.RectAnchoredPaths, rectanchoredpath)
	delete(stage.RectAnchoredPaths_mapString, rectanchoredpath.Name)

	if _, ok := stage.reference[rectanchoredpath]; ok {
		stage.deleted[rectanchoredpath] = struct{}{}
	} else {
		delete(stage.new, rectanchoredpath)
	}
	return rectanchoredpath
}

// UnstageVoid removes rectanchoredpath off the model stage
func (rectanchoredpath *RectAnchoredPath) UnstageVoid(stage *Stage) {
	delete(stage.RectAnchoredPaths, rectanchoredpath)
	delete(stage.RectAnchoredPaths_mapString, rectanchoredpath.Name)
}

// commit rectanchoredpath to the back repo (if it is already staged)
func (rectanchoredpath *RectAnchoredPath) Commit(stage *Stage) *RectAnchoredPath {
	if _, ok := stage.RectAnchoredPaths[rectanchoredpath]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRectAnchoredPath(rectanchoredpath)
		}
	}
	return rectanchoredpath
}

func (rectanchoredpath *RectAnchoredPath) CommitVoid(stage *Stage) {
	rectanchoredpath.Commit(stage)
}

func (rectanchoredpath *RectAnchoredPath) StageVoid(stage *Stage) {
	rectanchoredpath.Stage(stage)
}

// Checkout rectanchoredpath to the back repo (if it is already staged)
func (rectanchoredpath *RectAnchoredPath) Checkout(stage *Stage) *RectAnchoredPath {
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
func (rectanchoredrect *RectAnchoredRect) Stage(stage *Stage) *RectAnchoredRect {

	if _, ok := stage.RectAnchoredRects[rectanchoredrect]; !ok {
		stage.RectAnchoredRects[rectanchoredrect] = struct{}{}
		stage.RectAnchoredRectMap_Staged_Order[rectanchoredrect] = stage.RectAnchoredRectOrder
		stage.RectAnchoredRectOrder++
		stage.new[rectanchoredrect] = struct{}{}
		delete(stage.deleted, rectanchoredrect)
	} else {
		if _, ok := stage.new[rectanchoredrect]; !ok {
			stage.modified[rectanchoredrect] = struct{}{}
		}
	}
	stage.RectAnchoredRects_mapString[rectanchoredrect.Name] = rectanchoredrect

	return rectanchoredrect
}

// Unstage removes rectanchoredrect off the model stage
func (rectanchoredrect *RectAnchoredRect) Unstage(stage *Stage) *RectAnchoredRect {
	delete(stage.RectAnchoredRects, rectanchoredrect)
	delete(stage.RectAnchoredRects_mapString, rectanchoredrect.Name)

	if _, ok := stage.reference[rectanchoredrect]; ok {
		stage.deleted[rectanchoredrect] = struct{}{}
	} else {
		delete(stage.new, rectanchoredrect)
	}
	return rectanchoredrect
}

// UnstageVoid removes rectanchoredrect off the model stage
func (rectanchoredrect *RectAnchoredRect) UnstageVoid(stage *Stage) {
	delete(stage.RectAnchoredRects, rectanchoredrect)
	delete(stage.RectAnchoredRects_mapString, rectanchoredrect.Name)
}

// commit rectanchoredrect to the back repo (if it is already staged)
func (rectanchoredrect *RectAnchoredRect) Commit(stage *Stage) *RectAnchoredRect {
	if _, ok := stage.RectAnchoredRects[rectanchoredrect]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRectAnchoredRect(rectanchoredrect)
		}
	}
	return rectanchoredrect
}

func (rectanchoredrect *RectAnchoredRect) CommitVoid(stage *Stage) {
	rectanchoredrect.Commit(stage)
}

func (rectanchoredrect *RectAnchoredRect) StageVoid(stage *Stage) {
	rectanchoredrect.Stage(stage)
}

// Checkout rectanchoredrect to the back repo (if it is already staged)
func (rectanchoredrect *RectAnchoredRect) Checkout(stage *Stage) *RectAnchoredRect {
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
func (rectanchoredtext *RectAnchoredText) Stage(stage *Stage) *RectAnchoredText {

	if _, ok := stage.RectAnchoredTexts[rectanchoredtext]; !ok {
		stage.RectAnchoredTexts[rectanchoredtext] = struct{}{}
		stage.RectAnchoredTextMap_Staged_Order[rectanchoredtext] = stage.RectAnchoredTextOrder
		stage.RectAnchoredTextOrder++
		stage.new[rectanchoredtext] = struct{}{}
		delete(stage.deleted, rectanchoredtext)
	} else {
		if _, ok := stage.new[rectanchoredtext]; !ok {
			stage.modified[rectanchoredtext] = struct{}{}
		}
	}
	stage.RectAnchoredTexts_mapString[rectanchoredtext.Name] = rectanchoredtext

	return rectanchoredtext
}

// Unstage removes rectanchoredtext off the model stage
func (rectanchoredtext *RectAnchoredText) Unstage(stage *Stage) *RectAnchoredText {
	delete(stage.RectAnchoredTexts, rectanchoredtext)
	delete(stage.RectAnchoredTexts_mapString, rectanchoredtext.Name)

	if _, ok := stage.reference[rectanchoredtext]; ok {
		stage.deleted[rectanchoredtext] = struct{}{}
	} else {
		delete(stage.new, rectanchoredtext)
	}
	return rectanchoredtext
}

// UnstageVoid removes rectanchoredtext off the model stage
func (rectanchoredtext *RectAnchoredText) UnstageVoid(stage *Stage) {
	delete(stage.RectAnchoredTexts, rectanchoredtext)
	delete(stage.RectAnchoredTexts_mapString, rectanchoredtext.Name)
}

// commit rectanchoredtext to the back repo (if it is already staged)
func (rectanchoredtext *RectAnchoredText) Commit(stage *Stage) *RectAnchoredText {
	if _, ok := stage.RectAnchoredTexts[rectanchoredtext]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRectAnchoredText(rectanchoredtext)
		}
	}
	return rectanchoredtext
}

func (rectanchoredtext *RectAnchoredText) CommitVoid(stage *Stage) {
	rectanchoredtext.Commit(stage)
}

func (rectanchoredtext *RectAnchoredText) StageVoid(stage *Stage) {
	rectanchoredtext.Stage(stage)
}

// Checkout rectanchoredtext to the back repo (if it is already staged)
func (rectanchoredtext *RectAnchoredText) Checkout(stage *Stage) *RectAnchoredText {
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
func (rectlinklink *RectLinkLink) Stage(stage *Stage) *RectLinkLink {

	if _, ok := stage.RectLinkLinks[rectlinklink]; !ok {
		stage.RectLinkLinks[rectlinklink] = struct{}{}
		stage.RectLinkLinkMap_Staged_Order[rectlinklink] = stage.RectLinkLinkOrder
		stage.RectLinkLinkOrder++
		stage.new[rectlinklink] = struct{}{}
		delete(stage.deleted, rectlinklink)
	} else {
		if _, ok := stage.new[rectlinklink]; !ok {
			stage.modified[rectlinklink] = struct{}{}
		}
	}
	stage.RectLinkLinks_mapString[rectlinklink.Name] = rectlinklink

	return rectlinklink
}

// Unstage removes rectlinklink off the model stage
func (rectlinklink *RectLinkLink) Unstage(stage *Stage) *RectLinkLink {
	delete(stage.RectLinkLinks, rectlinklink)
	delete(stage.RectLinkLinks_mapString, rectlinklink.Name)

	if _, ok := stage.reference[rectlinklink]; ok {
		stage.deleted[rectlinklink] = struct{}{}
	} else {
		delete(stage.new, rectlinklink)
	}
	return rectlinklink
}

// UnstageVoid removes rectlinklink off the model stage
func (rectlinklink *RectLinkLink) UnstageVoid(stage *Stage) {
	delete(stage.RectLinkLinks, rectlinklink)
	delete(stage.RectLinkLinks_mapString, rectlinklink.Name)
}

// commit rectlinklink to the back repo (if it is already staged)
func (rectlinklink *RectLinkLink) Commit(stage *Stage) *RectLinkLink {
	if _, ok := stage.RectLinkLinks[rectlinklink]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRectLinkLink(rectlinklink)
		}
	}
	return rectlinklink
}

func (rectlinklink *RectLinkLink) CommitVoid(stage *Stage) {
	rectlinklink.Commit(stage)
}

func (rectlinklink *RectLinkLink) StageVoid(stage *Stage) {
	rectlinklink.Stage(stage)
}

// Checkout rectlinklink to the back repo (if it is already staged)
func (rectlinklink *RectLinkLink) Checkout(stage *Stage) *RectLinkLink {
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
func (svg *SVG) Stage(stage *Stage) *SVG {

	if _, ok := stage.SVGs[svg]; !ok {
		stage.SVGs[svg] = struct{}{}
		stage.SVGMap_Staged_Order[svg] = stage.SVGOrder
		stage.SVGOrder++
		stage.new[svg] = struct{}{}
		delete(stage.deleted, svg)
	} else {
		if _, ok := stage.new[svg]; !ok {
			stage.modified[svg] = struct{}{}
		}
	}
	stage.SVGs_mapString[svg.Name] = svg

	return svg
}

// Unstage removes svg off the model stage
func (svg *SVG) Unstage(stage *Stage) *SVG {
	delete(stage.SVGs, svg)
	delete(stage.SVGs_mapString, svg.Name)

	if _, ok := stage.reference[svg]; ok {
		stage.deleted[svg] = struct{}{}
	} else {
		delete(stage.new, svg)
	}
	return svg
}

// UnstageVoid removes svg off the model stage
func (svg *SVG) UnstageVoid(stage *Stage) {
	delete(stage.SVGs, svg)
	delete(stage.SVGs_mapString, svg.Name)
}

// commit svg to the back repo (if it is already staged)
func (svg *SVG) Commit(stage *Stage) *SVG {
	if _, ok := stage.SVGs[svg]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSVG(svg)
		}
	}
	return svg
}

func (svg *SVG) CommitVoid(stage *Stage) {
	svg.Commit(stage)
}

func (svg *SVG) StageVoid(stage *Stage) {
	svg.Stage(stage)
}

// Checkout svg to the back repo (if it is already staged)
func (svg *SVG) Checkout(stage *Stage) *SVG {
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

// Stage puts svgtext to the model stage
func (svgtext *SvgText) Stage(stage *Stage) *SvgText {

	if _, ok := stage.SvgTexts[svgtext]; !ok {
		stage.SvgTexts[svgtext] = struct{}{}
		stage.SvgTextMap_Staged_Order[svgtext] = stage.SvgTextOrder
		stage.SvgTextOrder++
		stage.new[svgtext] = struct{}{}
		delete(stage.deleted, svgtext)
	} else {
		if _, ok := stage.new[svgtext]; !ok {
			stage.modified[svgtext] = struct{}{}
		}
	}
	stage.SvgTexts_mapString[svgtext.Name] = svgtext

	return svgtext
}

// Unstage removes svgtext off the model stage
func (svgtext *SvgText) Unstage(stage *Stage) *SvgText {
	delete(stage.SvgTexts, svgtext)
	delete(stage.SvgTexts_mapString, svgtext.Name)

	if _, ok := stage.reference[svgtext]; ok {
		stage.deleted[svgtext] = struct{}{}
	} else {
		delete(stage.new, svgtext)
	}
	return svgtext
}

// UnstageVoid removes svgtext off the model stage
func (svgtext *SvgText) UnstageVoid(stage *Stage) {
	delete(stage.SvgTexts, svgtext)
	delete(stage.SvgTexts_mapString, svgtext.Name)
}

// commit svgtext to the back repo (if it is already staged)
func (svgtext *SvgText) Commit(stage *Stage) *SvgText {
	if _, ok := stage.SvgTexts[svgtext]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSvgText(svgtext)
		}
	}
	return svgtext
}

func (svgtext *SvgText) CommitVoid(stage *Stage) {
	svgtext.Commit(stage)
}

func (svgtext *SvgText) StageVoid(stage *Stage) {
	svgtext.Stage(stage)
}

// Checkout svgtext to the back repo (if it is already staged)
func (svgtext *SvgText) Checkout(stage *Stage) *SvgText {
	if _, ok := stage.SvgTexts[svgtext]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSvgText(svgtext)
		}
	}
	return svgtext
}

// for satisfaction of GongStruct interface
func (svgtext *SvgText) GetName() (res string) {
	return svgtext.Name
}

// Stage puts text to the model stage
func (text *Text) Stage(stage *Stage) *Text {

	if _, ok := stage.Texts[text]; !ok {
		stage.Texts[text] = struct{}{}
		stage.TextMap_Staged_Order[text] = stage.TextOrder
		stage.TextOrder++
		stage.new[text] = struct{}{}
		delete(stage.deleted, text)
	} else {
		if _, ok := stage.new[text]; !ok {
			stage.modified[text] = struct{}{}
		}
	}
	stage.Texts_mapString[text.Name] = text

	return text
}

// Unstage removes text off the model stage
func (text *Text) Unstage(stage *Stage) *Text {
	delete(stage.Texts, text)
	delete(stage.Texts_mapString, text.Name)

	if _, ok := stage.reference[text]; ok {
		stage.deleted[text] = struct{}{}
	} else {
		delete(stage.new, text)
	}
	return text
}

// UnstageVoid removes text off the model stage
func (text *Text) UnstageVoid(stage *Stage) {
	delete(stage.Texts, text)
	delete(stage.Texts_mapString, text.Name)
}

// commit text to the back repo (if it is already staged)
func (text *Text) Commit(stage *Stage) *Text {
	if _, ok := stage.Texts[text]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitText(text)
		}
	}
	return text
}

func (text *Text) CommitVoid(stage *Stage) {
	text.Commit(stage)
}

func (text *Text) StageVoid(stage *Stage) {
	text.Stage(stage)
}

// Checkout text to the back repo (if it is already staged)
func (text *Text) Checkout(stage *Stage) *Text {
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
	CreateORMCondition(Condition *Condition)
	CreateORMControlPoint(ControlPoint *ControlPoint)
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
	CreateORMSvgText(SvgText *SvgText)
	CreateORMText(Text *Text)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAnimate(Animate *Animate)
	DeleteORMCircle(Circle *Circle)
	DeleteORMCondition(Condition *Condition)
	DeleteORMControlPoint(ControlPoint *ControlPoint)
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
	DeleteORMSvgText(SvgText *SvgText)
	DeleteORMText(Text *Text)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Animates = make(map[*Animate]struct{})
	stage.Animates_mapString = make(map[string]*Animate)
	stage.AnimateMap_Staged_Order = make(map[*Animate]uint)
	stage.AnimateOrder = 0

	stage.Circles = make(map[*Circle]struct{})
	stage.Circles_mapString = make(map[string]*Circle)
	stage.CircleMap_Staged_Order = make(map[*Circle]uint)
	stage.CircleOrder = 0

	stage.Conditions = make(map[*Condition]struct{})
	stage.Conditions_mapString = make(map[string]*Condition)
	stage.ConditionMap_Staged_Order = make(map[*Condition]uint)
	stage.ConditionOrder = 0

	stage.ControlPoints = make(map[*ControlPoint]struct{})
	stage.ControlPoints_mapString = make(map[string]*ControlPoint)
	stage.ControlPointMap_Staged_Order = make(map[*ControlPoint]uint)
	stage.ControlPointOrder = 0

	stage.Ellipses = make(map[*Ellipse]struct{})
	stage.Ellipses_mapString = make(map[string]*Ellipse)
	stage.EllipseMap_Staged_Order = make(map[*Ellipse]uint)
	stage.EllipseOrder = 0

	stage.Layers = make(map[*Layer]struct{})
	stage.Layers_mapString = make(map[string]*Layer)
	stage.LayerMap_Staged_Order = make(map[*Layer]uint)
	stage.LayerOrder = 0

	stage.Lines = make(map[*Line]struct{})
	stage.Lines_mapString = make(map[string]*Line)
	stage.LineMap_Staged_Order = make(map[*Line]uint)
	stage.LineOrder = 0

	stage.Links = make(map[*Link]struct{})
	stage.Links_mapString = make(map[string]*Link)
	stage.LinkMap_Staged_Order = make(map[*Link]uint)
	stage.LinkOrder = 0

	stage.LinkAnchoredTexts = make(map[*LinkAnchoredText]struct{})
	stage.LinkAnchoredTexts_mapString = make(map[string]*LinkAnchoredText)
	stage.LinkAnchoredTextMap_Staged_Order = make(map[*LinkAnchoredText]uint)
	stage.LinkAnchoredTextOrder = 0

	stage.Paths = make(map[*Path]struct{})
	stage.Paths_mapString = make(map[string]*Path)
	stage.PathMap_Staged_Order = make(map[*Path]uint)
	stage.PathOrder = 0

	stage.Points = make(map[*Point]struct{})
	stage.Points_mapString = make(map[string]*Point)
	stage.PointMap_Staged_Order = make(map[*Point]uint)
	stage.PointOrder = 0

	stage.Polygones = make(map[*Polygone]struct{})
	stage.Polygones_mapString = make(map[string]*Polygone)
	stage.PolygoneMap_Staged_Order = make(map[*Polygone]uint)
	stage.PolygoneOrder = 0

	stage.Polylines = make(map[*Polyline]struct{})
	stage.Polylines_mapString = make(map[string]*Polyline)
	stage.PolylineMap_Staged_Order = make(map[*Polyline]uint)
	stage.PolylineOrder = 0

	stage.Rects = make(map[*Rect]struct{})
	stage.Rects_mapString = make(map[string]*Rect)
	stage.RectMap_Staged_Order = make(map[*Rect]uint)
	stage.RectOrder = 0

	stage.RectAnchoredPaths = make(map[*RectAnchoredPath]struct{})
	stage.RectAnchoredPaths_mapString = make(map[string]*RectAnchoredPath)
	stage.RectAnchoredPathMap_Staged_Order = make(map[*RectAnchoredPath]uint)
	stage.RectAnchoredPathOrder = 0

	stage.RectAnchoredRects = make(map[*RectAnchoredRect]struct{})
	stage.RectAnchoredRects_mapString = make(map[string]*RectAnchoredRect)
	stage.RectAnchoredRectMap_Staged_Order = make(map[*RectAnchoredRect]uint)
	stage.RectAnchoredRectOrder = 0

	stage.RectAnchoredTexts = make(map[*RectAnchoredText]struct{})
	stage.RectAnchoredTexts_mapString = make(map[string]*RectAnchoredText)
	stage.RectAnchoredTextMap_Staged_Order = make(map[*RectAnchoredText]uint)
	stage.RectAnchoredTextOrder = 0

	stage.RectLinkLinks = make(map[*RectLinkLink]struct{})
	stage.RectLinkLinks_mapString = make(map[string]*RectLinkLink)
	stage.RectLinkLinkMap_Staged_Order = make(map[*RectLinkLink]uint)
	stage.RectLinkLinkOrder = 0

	stage.SVGs = make(map[*SVG]struct{})
	stage.SVGs_mapString = make(map[string]*SVG)
	stage.SVGMap_Staged_Order = make(map[*SVG]uint)
	stage.SVGOrder = 0

	stage.SvgTexts = make(map[*SvgText]struct{})
	stage.SvgTexts_mapString = make(map[string]*SvgText)
	stage.SvgTextMap_Staged_Order = make(map[*SvgText]uint)
	stage.SvgTextOrder = 0

	stage.Texts = make(map[*Text]struct{})
	stage.Texts_mapString = make(map[string]*Text)
	stage.TextMap_Staged_Order = make(map[*Text]uint)
	stage.TextOrder = 0

	stage.ComputeReference()
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.Animates = nil
	stage.Animates_mapString = nil

	stage.Circles = nil
	stage.Circles_mapString = nil

	stage.Conditions = nil
	stage.Conditions_mapString = nil

	stage.ControlPoints = nil
	stage.ControlPoints_mapString = nil

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

	stage.SvgTexts = nil
	stage.SvgTexts_mapString = nil

	stage.Texts = nil
	stage.Texts_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for animate := range stage.Animates {
		animate.Unstage(stage)
	}

	for circle := range stage.Circles {
		circle.Unstage(stage)
	}

	for condition := range stage.Conditions {
		condition.Unstage(stage)
	}

	for controlpoint := range stage.ControlPoints {
		controlpoint.Unstage(stage)
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

	for svgtext := range stage.SvgTexts {
		svgtext.Unstage(stage)
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
type GongstructIF interface {
	GetName() string
	CommitVoid(*Stage)
	StageVoid(*Stage)
	UnstageVoid(stage *Stage)
	GongGetFieldHeaders() []GongFieldHeader
	GongClean(stage *Stage)
	GongGetFieldValue(fieldName string, stage *Stage) GongFieldValue
	GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error
	GongGetGongstructName() string
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
	case map[*Animate]any:
		return any(&stage.Animates).(*Type)
	case map[*Circle]any:
		return any(&stage.Circles).(*Type)
	case map[*Condition]any:
		return any(&stage.Conditions).(*Type)
	case map[*ControlPoint]any:
		return any(&stage.ControlPoints).(*Type)
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
	case map[*SvgText]any:
		return any(&stage.SvgTexts).(*Type)
	case map[*Text]any:
		return any(&stage.Texts).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *Stage) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*Animate:
		return any(&stage.Animates_mapString).(*Type)
	case map[string]*Circle:
		return any(&stage.Circles_mapString).(*Type)
	case map[string]*Condition:
		return any(&stage.Conditions_mapString).(*Type)
	case map[string]*ControlPoint:
		return any(&stage.ControlPoints_mapString).(*Type)
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
	case map[string]*SvgText:
		return any(&stage.SvgTexts_mapString).(*Type)
	case map[string]*Text:
		return any(&stage.Texts_mapString).(*Type)
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
	case Animate:
		return any(&stage.Animates).(*map[*Type]struct{})
	case Circle:
		return any(&stage.Circles).(*map[*Type]struct{})
	case Condition:
		return any(&stage.Conditions).(*map[*Type]struct{})
	case ControlPoint:
		return any(&stage.ControlPoints).(*map[*Type]struct{})
	case Ellipse:
		return any(&stage.Ellipses).(*map[*Type]struct{})
	case Layer:
		return any(&stage.Layers).(*map[*Type]struct{})
	case Line:
		return any(&stage.Lines).(*map[*Type]struct{})
	case Link:
		return any(&stage.Links).(*map[*Type]struct{})
	case LinkAnchoredText:
		return any(&stage.LinkAnchoredTexts).(*map[*Type]struct{})
	case Path:
		return any(&stage.Paths).(*map[*Type]struct{})
	case Point:
		return any(&stage.Points).(*map[*Type]struct{})
	case Polygone:
		return any(&stage.Polygones).(*map[*Type]struct{})
	case Polyline:
		return any(&stage.Polylines).(*map[*Type]struct{})
	case Rect:
		return any(&stage.Rects).(*map[*Type]struct{})
	case RectAnchoredPath:
		return any(&stage.RectAnchoredPaths).(*map[*Type]struct{})
	case RectAnchoredRect:
		return any(&stage.RectAnchoredRects).(*map[*Type]struct{})
	case RectAnchoredText:
		return any(&stage.RectAnchoredTexts).(*map[*Type]struct{})
	case RectLinkLink:
		return any(&stage.RectLinkLinks).(*map[*Type]struct{})
	case SVG:
		return any(&stage.SVGs).(*map[*Type]struct{})
	case SvgText:
		return any(&stage.SvgTexts).(*map[*Type]struct{})
	case Text:
		return any(&stage.Texts).(*map[*Type]struct{})
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
	case *Layer:
		return any(&stage.Layers).(*map[Type]struct{})
	case *Line:
		return any(&stage.Lines).(*map[Type]struct{})
	case *Link:
		return any(&stage.Links).(*map[Type]struct{})
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

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *Stage) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Animate:
		return any(&stage.Animates_mapString).(*map[string]*Type)
	case Circle:
		return any(&stage.Circles_mapString).(*map[string]*Type)
	case Condition:
		return any(&stage.Conditions_mapString).(*map[string]*Type)
	case ControlPoint:
		return any(&stage.ControlPoints_mapString).(*map[string]*Type)
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
	case SvgText:
		return any(&stage.SvgTexts_mapString).(*map[string]*Type)
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
	case Condition:
		return any(&Condition{
			// Initialisation of associations
		}).(*Type)
	case ControlPoint:
		return any(&ControlPoint{
			// Initialisation of associations
			// field is initialized with an instance of Rect with the name of the field
			ClosestRect: &Rect{Name: "ClosestRect"},
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
			TextAtArrowStart: []*LinkAnchoredText{{Name: "TextAtArrowStart"}},
			// field is initialized with an instance of LinkAnchoredText with the name of the field
			TextAtArrowEnd: []*LinkAnchoredText{{Name: "TextAtArrowEnd"}},
			// field is initialized with an instance of ControlPoint with the name of the field
			ControlPoints: []*ControlPoint{{Name: "ControlPoints"}},
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
	case SvgText:
		return any(&SvgText{
			// Initialisation of associations
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
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End][]*Start {

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
		case "ControlPoints":
			res := make(map[*ControlPoint][]*Link)
			for link := range stage.Links {
				for _, controlpoint_ := range link.ControlPoints {
					res[controlpoint_] = append(res[controlpoint_], link)
				}
			}
			return any(res).(map[*End][]*Start)
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
	case *Condition:
		res = "Condition"
	case *ControlPoint:
		res = "ControlPoint"
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
	case *SvgText:
		res = "SvgText"
	case *Text:
		res = "Text"
	}
	return res
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type PointerToGongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

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
	case *LinkAnchoredText:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Link"
		rf.Fieldname = "TextAtArrowStart"
		res = append(res, rf)
		rf.GongstructName = "Link"
		rf.Fieldname = "TextAtArrowEnd"
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
	case *RectAnchoredPath:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Rect"
		rf.Fieldname = "RectAnchoredPaths"
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

// insertion point for get fields header method
func (animate *Animate) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "AttributeName",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Values",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "From",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "To",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Dur",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "RepeatCount",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (circle *Circle) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CX",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CY",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Radius",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (controlpoint *ControlPoint) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "X_Relative",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Y_Relative",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CX",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CY",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "RX",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "RY",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Animates",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Animate",
		},
	}
	return
}

func (layer *Layer) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "X1",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Y1",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "X2",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Y2",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Animates",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Animate",
		},
		{
			Name:               "MouseClickX",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "MouseClickY",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (link *Link) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Type",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsBezierCurve",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Start",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Rect",
		},
		{
			Name:               "StartAnchorType",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "End",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Rect",
		},
		{
			Name:               "EndAnchorType",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StartOrientation",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StartRatio",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "EndOrientation",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "EndRatio",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CornerOffsetRatio",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CornerRadius",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "HasEndArrow",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "EndArrowSize",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "EndArrowOffset",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "HasStartArrow",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StartArrowSize",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StartArrowOffset",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			Name:                 "ControlPoints",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlPoint",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "MouseX",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "MouseY",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "MouseEventKey",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (linkanchoredtext *LinkAnchoredText) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Content",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "AutomaticLayout",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "LinkAnchorType",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "X_Offset",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Y_Offset",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FontWeight",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FontStyle",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "LetterSpacing",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "WhiteSpace",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Definition",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (polygone *Polygone) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Points",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Points",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "RX",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsSelected",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CanHaveLeftHandle",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "HasLeftHandle",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CanHaveRightHandle",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "HasRightHandle",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CanHaveTopHandle",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "HasTopHandle",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsScalingProportionally",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CanHaveBottomHandle",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "HasBottomHandle",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CanMoveHorizontaly",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CanMoveVerticaly",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			Name:               "ChangeColorWhenHovered",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ColorWhenHovered",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "OriginalColor",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FillOpacityWhenHovered",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "OriginalFillOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "HasToolTip",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ToolTipText",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ToolTipPosition",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "MouseX",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "MouseY",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "MouseEventKey",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (rectanchoredpath *RectAnchoredPath) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Definition",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "X_Offset",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Y_Offset",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "RectAnchorType",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ScalePropotionnally",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "AppliedScaling",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (rectanchoredrect *RectAnchoredRect) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "RX",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "X_Offset",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Y_Offset",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "RectAnchorType",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "WidthFollowRect",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "HeightFollowRect",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "HasToolTip",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ToolTipText",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (rectanchoredtext *RectAnchoredText) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Content",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FontWeight",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FontStyle",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "LetterSpacing",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "WhiteSpace",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "X_Offset",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Y_Offset",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "RectAnchorType",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "TextAnchorType",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "DominantBaseline",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "WritingMode",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Animates",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Animate",
		},
	}
	return
}

func (rectlinklink *RectLinkLink) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (svg *SVG) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Layers",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Layer",
		},
		{
			Name:               "DrawingState",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsSVGFrontEndFileGenerated",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsSVGBackEndFileGenerated",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "DefaultDirectoryForGeneratedImages",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsControlBannerHidden",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (svgtext *SvgText) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Text",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (text *Text) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Content",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FontWeight",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FontStyle",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "LetterSpacing",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "WhiteSpace",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Animates",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Animate",
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
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
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, controlpoint.ClosestRect))
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Texts":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.Texts {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Circles":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.Circles {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Lines":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.Lines {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Ellipses":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.Ellipses {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Polylines":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.Polylines {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Polygones":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.Polygones {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Paths":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.Paths {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Links":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.Links {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "RectLinkLinks":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layer.RectLinkLinks {
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
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
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, link.Start))
		}
	case "StartAnchorType":
		enum := link.StartAnchorType
		res.valueString = enum.ToCodeString()
	case "End":
		res.GongFieldValueType = GongFieldValueTypePointer
		if link.End != nil {
			res.valueString = link.End.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, link.End))
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "TextAtArrowEnd":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range link.TextAtArrowEnd {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "ControlPoints":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range link.ControlPoints {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "DisplayConditions":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range rect.DisplayConditions {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Animations":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range rect.Animations {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "RectAnchoredRects":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range rect.RectAnchoredRects {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "RectAnchoredPaths":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range rect.RectAnchoredPaths {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
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
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, rectlinklink.Start))
		}
	case "End":
		res.GongFieldValueType = GongFieldValueTypePointer
		if rectlinklink.End != nil {
			res.valueString = rectlinklink.End.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, rectlinklink.End))
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "DrawingState":
		enum := svg.DrawingState
		res.valueString = enum.ToCodeString()
	case "StartRect":
		res.GongFieldValueType = GongFieldValueTypePointer
		if svg.StartRect != nil {
			res.valueString = svg.StartRect.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, svg.StartRect))
		}
	case "EndRect":
		res.GongFieldValueType = GongFieldValueTypePointer
		if svg.EndRect != nil {
			res.valueString = svg.EndRect.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, svg.EndRect))
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	}
	return
}
func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {

	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (animate *Animate) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		animate.Name = value.GetValueString()
	case "AttributeName":
		animate.AttributeName = value.GetValueString()
	case "Values":
		animate.Values = value.GetValueString()
	case "From":
		animate.From = value.GetValueString()
	case "To":
		animate.To = value.GetValueString()
	case "Dur":
		animate.Dur = value.GetValueString()
	case "RepeatCount":
		animate.RepeatCount = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (circle *Circle) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		circle.Name = value.GetValueString()
	case "CX":
		circle.CX = value.GetValueFloat()
	case "CY":
		circle.CY = value.GetValueFloat()
	case "Radius":
		circle.Radius = value.GetValueFloat()
	case "Color":
		circle.Color = value.GetValueString()
	case "FillOpacity":
		circle.FillOpacity = value.GetValueFloat()
	case "Stroke":
		circle.Stroke = value.GetValueString()
	case "StrokeOpacity":
		circle.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		circle.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		circle.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		circle.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		circle.Transform = value.GetValueString()
	case "Animations":
		circle.Animations = make([]*Animate, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Animates {
					if stage.AnimateMap_Staged_Order[__instance__] == uint(id) {
						circle.Animations = append(circle.Animations, __instance__)
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

func (condition *Condition) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		condition.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (controlpoint *ControlPoint) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		controlpoint.Name = value.GetValueString()
	case "X_Relative":
		controlpoint.X_Relative = value.GetValueFloat()
	case "Y_Relative":
		controlpoint.Y_Relative = value.GetValueFloat()
	case "ClosestRect":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			controlpoint.ClosestRect = nil
			for __instance__ := range stage.Rects {
				if stage.RectMap_Staged_Order[__instance__] == uint(id) {
					controlpoint.ClosestRect = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (ellipse *Ellipse) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		ellipse.Name = value.GetValueString()
	case "CX":
		ellipse.CX = value.GetValueFloat()
	case "CY":
		ellipse.CY = value.GetValueFloat()
	case "RX":
		ellipse.RX = value.GetValueFloat()
	case "RY":
		ellipse.RY = value.GetValueFloat()
	case "Color":
		ellipse.Color = value.GetValueString()
	case "FillOpacity":
		ellipse.FillOpacity = value.GetValueFloat()
	case "Stroke":
		ellipse.Stroke = value.GetValueString()
	case "StrokeOpacity":
		ellipse.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		ellipse.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		ellipse.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		ellipse.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		ellipse.Transform = value.GetValueString()
	case "Animates":
		ellipse.Animates = make([]*Animate, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Animates {
					if stage.AnimateMap_Staged_Order[__instance__] == uint(id) {
						ellipse.Animates = append(ellipse.Animates, __instance__)
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

func (layer *Layer) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		layer.Name = value.GetValueString()
	case "Rects":
		layer.Rects = make([]*Rect, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Rects {
					if stage.RectMap_Staged_Order[__instance__] == uint(id) {
						layer.Rects = append(layer.Rects, __instance__)
						break
					}
				}
			}
		}
	case "Texts":
		layer.Texts = make([]*Text, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Texts {
					if stage.TextMap_Staged_Order[__instance__] == uint(id) {
						layer.Texts = append(layer.Texts, __instance__)
						break
					}
				}
			}
		}
	case "Circles":
		layer.Circles = make([]*Circle, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Circles {
					if stage.CircleMap_Staged_Order[__instance__] == uint(id) {
						layer.Circles = append(layer.Circles, __instance__)
						break
					}
				}
			}
		}
	case "Lines":
		layer.Lines = make([]*Line, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Lines {
					if stage.LineMap_Staged_Order[__instance__] == uint(id) {
						layer.Lines = append(layer.Lines, __instance__)
						break
					}
				}
			}
		}
	case "Ellipses":
		layer.Ellipses = make([]*Ellipse, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Ellipses {
					if stage.EllipseMap_Staged_Order[__instance__] == uint(id) {
						layer.Ellipses = append(layer.Ellipses, __instance__)
						break
					}
				}
			}
		}
	case "Polylines":
		layer.Polylines = make([]*Polyline, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Polylines {
					if stage.PolylineMap_Staged_Order[__instance__] == uint(id) {
						layer.Polylines = append(layer.Polylines, __instance__)
						break
					}
				}
			}
		}
	case "Polygones":
		layer.Polygones = make([]*Polygone, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Polygones {
					if stage.PolygoneMap_Staged_Order[__instance__] == uint(id) {
						layer.Polygones = append(layer.Polygones, __instance__)
						break
					}
				}
			}
		}
	case "Paths":
		layer.Paths = make([]*Path, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Paths {
					if stage.PathMap_Staged_Order[__instance__] == uint(id) {
						layer.Paths = append(layer.Paths, __instance__)
						break
					}
				}
			}
		}
	case "Links":
		layer.Links = make([]*Link, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Links {
					if stage.LinkMap_Staged_Order[__instance__] == uint(id) {
						layer.Links = append(layer.Links, __instance__)
						break
					}
				}
			}
		}
	case "RectLinkLinks":
		layer.RectLinkLinks = make([]*RectLinkLink, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.RectLinkLinks {
					if stage.RectLinkLinkMap_Staged_Order[__instance__] == uint(id) {
						layer.RectLinkLinks = append(layer.RectLinkLinks, __instance__)
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

func (line *Line) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		line.Name = value.GetValueString()
	case "X1":
		line.X1 = value.GetValueFloat()
	case "Y1":
		line.Y1 = value.GetValueFloat()
	case "X2":
		line.X2 = value.GetValueFloat()
	case "Y2":
		line.Y2 = value.GetValueFloat()
	case "Color":
		line.Color = value.GetValueString()
	case "FillOpacity":
		line.FillOpacity = value.GetValueFloat()
	case "Stroke":
		line.Stroke = value.GetValueString()
	case "StrokeOpacity":
		line.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		line.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		line.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		line.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		line.Transform = value.GetValueString()
	case "Animates":
		line.Animates = make([]*Animate, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Animates {
					if stage.AnimateMap_Staged_Order[__instance__] == uint(id) {
						line.Animates = append(line.Animates, __instance__)
						break
					}
				}
			}
		}
	case "MouseClickX":
		line.MouseClickX = value.GetValueFloat()
	case "MouseClickY":
		line.MouseClickY = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (link *Link) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		link.Name = value.GetValueString()
	case "Type":
		link.Type.FromCodeString(value.GetValueString())
	case "IsBezierCurve":
		link.IsBezierCurve = value.GetValueBool()
	case "Start":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			link.Start = nil
			for __instance__ := range stage.Rects {
				if stage.RectMap_Staged_Order[__instance__] == uint(id) {
					link.Start = __instance__
					break
				}
			}
		}
	case "StartAnchorType":
		link.StartAnchorType.FromCodeString(value.GetValueString())
	case "End":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			link.End = nil
			for __instance__ := range stage.Rects {
				if stage.RectMap_Staged_Order[__instance__] == uint(id) {
					link.End = __instance__
					break
				}
			}
		}
	case "EndAnchorType":
		link.EndAnchorType.FromCodeString(value.GetValueString())
	case "StartOrientation":
		link.StartOrientation.FromCodeString(value.GetValueString())
	case "StartRatio":
		link.StartRatio = value.GetValueFloat()
	case "EndOrientation":
		link.EndOrientation.FromCodeString(value.GetValueString())
	case "EndRatio":
		link.EndRatio = value.GetValueFloat()
	case "CornerOffsetRatio":
		link.CornerOffsetRatio = value.GetValueFloat()
	case "CornerRadius":
		link.CornerRadius = value.GetValueFloat()
	case "HasEndArrow":
		link.HasEndArrow = value.GetValueBool()
	case "EndArrowSize":
		link.EndArrowSize = value.GetValueFloat()
	case "EndArrowOffset":
		link.EndArrowOffset = value.GetValueFloat()
	case "HasStartArrow":
		link.HasStartArrow = value.GetValueBool()
	case "StartArrowSize":
		link.StartArrowSize = value.GetValueFloat()
	case "StartArrowOffset":
		link.StartArrowOffset = value.GetValueFloat()
	case "TextAtArrowStart":
		link.TextAtArrowStart = make([]*LinkAnchoredText, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.LinkAnchoredTexts {
					if stage.LinkAnchoredTextMap_Staged_Order[__instance__] == uint(id) {
						link.TextAtArrowStart = append(link.TextAtArrowStart, __instance__)
						break
					}
				}
			}
		}
	case "TextAtArrowEnd":
		link.TextAtArrowEnd = make([]*LinkAnchoredText, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.LinkAnchoredTexts {
					if stage.LinkAnchoredTextMap_Staged_Order[__instance__] == uint(id) {
						link.TextAtArrowEnd = append(link.TextAtArrowEnd, __instance__)
						break
					}
				}
			}
		}
	case "ControlPoints":
		link.ControlPoints = make([]*ControlPoint, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ControlPoints {
					if stage.ControlPointMap_Staged_Order[__instance__] == uint(id) {
						link.ControlPoints = append(link.ControlPoints, __instance__)
						break
					}
				}
			}
		}
	case "Color":
		link.Color = value.GetValueString()
	case "FillOpacity":
		link.FillOpacity = value.GetValueFloat()
	case "Stroke":
		link.Stroke = value.GetValueString()
	case "StrokeOpacity":
		link.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		link.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		link.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		link.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		link.Transform = value.GetValueString()
	case "MouseX":
		link.MouseX = value.GetValueFloat()
	case "MouseY":
		link.MouseY = value.GetValueFloat()
	case "MouseEventKey":
		link.MouseEventKey.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (linkanchoredtext *LinkAnchoredText) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		linkanchoredtext.Name = value.GetValueString()
	case "Content":
		linkanchoredtext.Content = value.GetValueString()
	case "AutomaticLayout":
		linkanchoredtext.AutomaticLayout = value.GetValueBool()
	case "LinkAnchorType":
		linkanchoredtext.LinkAnchorType.FromCodeString(value.GetValueString())
	case "X_Offset":
		linkanchoredtext.X_Offset = value.GetValueFloat()
	case "Y_Offset":
		linkanchoredtext.Y_Offset = value.GetValueFloat()
	case "FontWeight":
		linkanchoredtext.FontWeight = value.GetValueString()
	case "FontSize":
		linkanchoredtext.FontSize = value.GetValueString()
	case "FontStyle":
		linkanchoredtext.FontStyle = value.GetValueString()
	case "LetterSpacing":
		linkanchoredtext.LetterSpacing = value.GetValueString()
	case "FontFamily":
		linkanchoredtext.FontFamily = value.GetValueString()
	case "WhiteSpace":
		linkanchoredtext.WhiteSpace.FromCodeString(value.GetValueString())
	case "Color":
		linkanchoredtext.Color = value.GetValueString()
	case "FillOpacity":
		linkanchoredtext.FillOpacity = value.GetValueFloat()
	case "Stroke":
		linkanchoredtext.Stroke = value.GetValueString()
	case "StrokeOpacity":
		linkanchoredtext.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		linkanchoredtext.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		linkanchoredtext.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		linkanchoredtext.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		linkanchoredtext.Transform = value.GetValueString()
	case "Animates":
		linkanchoredtext.Animates = make([]*Animate, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Animates {
					if stage.AnimateMap_Staged_Order[__instance__] == uint(id) {
						linkanchoredtext.Animates = append(linkanchoredtext.Animates, __instance__)
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

func (path *Path) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		path.Name = value.GetValueString()
	case "Definition":
		path.Definition = value.GetValueString()
	case "Color":
		path.Color = value.GetValueString()
	case "FillOpacity":
		path.FillOpacity = value.GetValueFloat()
	case "Stroke":
		path.Stroke = value.GetValueString()
	case "StrokeOpacity":
		path.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		path.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		path.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		path.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		path.Transform = value.GetValueString()
	case "Animates":
		path.Animates = make([]*Animate, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Animates {
					if stage.AnimateMap_Staged_Order[__instance__] == uint(id) {
						path.Animates = append(path.Animates, __instance__)
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

func (point *Point) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		point.Name = value.GetValueString()
	case "X":
		point.X = value.GetValueFloat()
	case "Y":
		point.Y = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (polygone *Polygone) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		polygone.Name = value.GetValueString()
	case "Points":
		polygone.Points = value.GetValueString()
	case "Color":
		polygone.Color = value.GetValueString()
	case "FillOpacity":
		polygone.FillOpacity = value.GetValueFloat()
	case "Stroke":
		polygone.Stroke = value.GetValueString()
	case "StrokeOpacity":
		polygone.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		polygone.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		polygone.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		polygone.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		polygone.Transform = value.GetValueString()
	case "Animates":
		polygone.Animates = make([]*Animate, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Animates {
					if stage.AnimateMap_Staged_Order[__instance__] == uint(id) {
						polygone.Animates = append(polygone.Animates, __instance__)
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

func (polyline *Polyline) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		polyline.Name = value.GetValueString()
	case "Points":
		polyline.Points = value.GetValueString()
	case "Color":
		polyline.Color = value.GetValueString()
	case "FillOpacity":
		polyline.FillOpacity = value.GetValueFloat()
	case "Stroke":
		polyline.Stroke = value.GetValueString()
	case "StrokeOpacity":
		polyline.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		polyline.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		polyline.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		polyline.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		polyline.Transform = value.GetValueString()
	case "Animates":
		polyline.Animates = make([]*Animate, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Animates {
					if stage.AnimateMap_Staged_Order[__instance__] == uint(id) {
						polyline.Animates = append(polyline.Animates, __instance__)
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

func (rect *Rect) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		rect.Name = value.GetValueString()
	case "X":
		rect.X = value.GetValueFloat()
	case "Y":
		rect.Y = value.GetValueFloat()
	case "Width":
		rect.Width = value.GetValueFloat()
	case "Height":
		rect.Height = value.GetValueFloat()
	case "RX":
		rect.RX = value.GetValueFloat()
	case "Color":
		rect.Color = value.GetValueString()
	case "FillOpacity":
		rect.FillOpacity = value.GetValueFloat()
	case "Stroke":
		rect.Stroke = value.GetValueString()
	case "StrokeOpacity":
		rect.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		rect.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		rect.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		rect.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		rect.Transform = value.GetValueString()
	case "HoveringTrigger":
		rect.HoveringTrigger = make([]*Condition, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Conditions {
					if stage.ConditionMap_Staged_Order[__instance__] == uint(id) {
						rect.HoveringTrigger = append(rect.HoveringTrigger, __instance__)
						break
					}
				}
			}
		}
	case "DisplayConditions":
		rect.DisplayConditions = make([]*Condition, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Conditions {
					if stage.ConditionMap_Staged_Order[__instance__] == uint(id) {
						rect.DisplayConditions = append(rect.DisplayConditions, __instance__)
						break
					}
				}
			}
		}
	case "Animations":
		rect.Animations = make([]*Animate, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Animates {
					if stage.AnimateMap_Staged_Order[__instance__] == uint(id) {
						rect.Animations = append(rect.Animations, __instance__)
						break
					}
				}
			}
		}
	case "IsSelectable":
		rect.IsSelectable = value.GetValueBool()
	case "IsSelected":
		rect.IsSelected = value.GetValueBool()
	case "CanHaveLeftHandle":
		rect.CanHaveLeftHandle = value.GetValueBool()
	case "HasLeftHandle":
		rect.HasLeftHandle = value.GetValueBool()
	case "CanHaveRightHandle":
		rect.CanHaveRightHandle = value.GetValueBool()
	case "HasRightHandle":
		rect.HasRightHandle = value.GetValueBool()
	case "CanHaveTopHandle":
		rect.CanHaveTopHandle = value.GetValueBool()
	case "HasTopHandle":
		rect.HasTopHandle = value.GetValueBool()
	case "IsScalingProportionally":
		rect.IsScalingProportionally = value.GetValueBool()
	case "CanHaveBottomHandle":
		rect.CanHaveBottomHandle = value.GetValueBool()
	case "HasBottomHandle":
		rect.HasBottomHandle = value.GetValueBool()
	case "CanMoveHorizontaly":
		rect.CanMoveHorizontaly = value.GetValueBool()
	case "CanMoveVerticaly":
		rect.CanMoveVerticaly = value.GetValueBool()
	case "RectAnchoredTexts":
		rect.RectAnchoredTexts = make([]*RectAnchoredText, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.RectAnchoredTexts {
					if stage.RectAnchoredTextMap_Staged_Order[__instance__] == uint(id) {
						rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, __instance__)
						break
					}
				}
			}
		}
	case "RectAnchoredRects":
		rect.RectAnchoredRects = make([]*RectAnchoredRect, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.RectAnchoredRects {
					if stage.RectAnchoredRectMap_Staged_Order[__instance__] == uint(id) {
						rect.RectAnchoredRects = append(rect.RectAnchoredRects, __instance__)
						break
					}
				}
			}
		}
	case "RectAnchoredPaths":
		rect.RectAnchoredPaths = make([]*RectAnchoredPath, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.RectAnchoredPaths {
					if stage.RectAnchoredPathMap_Staged_Order[__instance__] == uint(id) {
						rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, __instance__)
						break
					}
				}
			}
		}
	case "ChangeColorWhenHovered":
		rect.ChangeColorWhenHovered = value.GetValueBool()
	case "ColorWhenHovered":
		rect.ColorWhenHovered = value.GetValueString()
	case "OriginalColor":
		rect.OriginalColor = value.GetValueString()
	case "FillOpacityWhenHovered":
		rect.FillOpacityWhenHovered = value.GetValueFloat()
	case "OriginalFillOpacity":
		rect.OriginalFillOpacity = value.GetValueFloat()
	case "HasToolTip":
		rect.HasToolTip = value.GetValueBool()
	case "ToolTipText":
		rect.ToolTipText = value.GetValueString()
	case "ToolTipPosition":
		rect.ToolTipPosition.FromCodeString(value.GetValueString())
	case "MouseX":
		rect.MouseX = value.GetValueFloat()
	case "MouseY":
		rect.MouseY = value.GetValueFloat()
	case "MouseEventKey":
		rect.MouseEventKey.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (rectanchoredpath *RectAnchoredPath) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		rectanchoredpath.Name = value.GetValueString()
	case "Definition":
		rectanchoredpath.Definition = value.GetValueString()
	case "X_Offset":
		rectanchoredpath.X_Offset = value.GetValueFloat()
	case "Y_Offset":
		rectanchoredpath.Y_Offset = value.GetValueFloat()
	case "RectAnchorType":
		rectanchoredpath.RectAnchorType.FromCodeString(value.GetValueString())
	case "ScalePropotionnally":
		rectanchoredpath.ScalePropotionnally = value.GetValueBool()
	case "AppliedScaling":
		rectanchoredpath.AppliedScaling = value.GetValueFloat()
	case "Color":
		rectanchoredpath.Color = value.GetValueString()
	case "FillOpacity":
		rectanchoredpath.FillOpacity = value.GetValueFloat()
	case "Stroke":
		rectanchoredpath.Stroke = value.GetValueString()
	case "StrokeOpacity":
		rectanchoredpath.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		rectanchoredpath.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		rectanchoredpath.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		rectanchoredpath.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		rectanchoredpath.Transform = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (rectanchoredrect *RectAnchoredRect) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		rectanchoredrect.Name = value.GetValueString()
	case "X":
		rectanchoredrect.X = value.GetValueFloat()
	case "Y":
		rectanchoredrect.Y = value.GetValueFloat()
	case "Width":
		rectanchoredrect.Width = value.GetValueFloat()
	case "Height":
		rectanchoredrect.Height = value.GetValueFloat()
	case "RX":
		rectanchoredrect.RX = value.GetValueFloat()
	case "X_Offset":
		rectanchoredrect.X_Offset = value.GetValueFloat()
	case "Y_Offset":
		rectanchoredrect.Y_Offset = value.GetValueFloat()
	case "RectAnchorType":
		rectanchoredrect.RectAnchorType.FromCodeString(value.GetValueString())
	case "WidthFollowRect":
		rectanchoredrect.WidthFollowRect = value.GetValueBool()
	case "HeightFollowRect":
		rectanchoredrect.HeightFollowRect = value.GetValueBool()
	case "HasToolTip":
		rectanchoredrect.HasToolTip = value.GetValueBool()
	case "ToolTipText":
		rectanchoredrect.ToolTipText = value.GetValueString()
	case "Color":
		rectanchoredrect.Color = value.GetValueString()
	case "FillOpacity":
		rectanchoredrect.FillOpacity = value.GetValueFloat()
	case "Stroke":
		rectanchoredrect.Stroke = value.GetValueString()
	case "StrokeOpacity":
		rectanchoredrect.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		rectanchoredrect.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		rectanchoredrect.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		rectanchoredrect.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		rectanchoredrect.Transform = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (rectanchoredtext *RectAnchoredText) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		rectanchoredtext.Name = value.GetValueString()
	case "Content":
		rectanchoredtext.Content = value.GetValueString()
	case "FontWeight":
		rectanchoredtext.FontWeight = value.GetValueString()
	case "FontSize":
		rectanchoredtext.FontSize = value.GetValueString()
	case "FontStyle":
		rectanchoredtext.FontStyle = value.GetValueString()
	case "LetterSpacing":
		rectanchoredtext.LetterSpacing = value.GetValueString()
	case "FontFamily":
		rectanchoredtext.FontFamily = value.GetValueString()
	case "WhiteSpace":
		rectanchoredtext.WhiteSpace.FromCodeString(value.GetValueString())
	case "X_Offset":
		rectanchoredtext.X_Offset = value.GetValueFloat()
	case "Y_Offset":
		rectanchoredtext.Y_Offset = value.GetValueFloat()
	case "RectAnchorType":
		rectanchoredtext.RectAnchorType.FromCodeString(value.GetValueString())
	case "TextAnchorType":
		rectanchoredtext.TextAnchorType.FromCodeString(value.GetValueString())
	case "DominantBaseline":
		rectanchoredtext.DominantBaseline.FromCodeString(value.GetValueString())
	case "WritingMode":
		rectanchoredtext.WritingMode.FromCodeString(value.GetValueString())
	case "Color":
		rectanchoredtext.Color = value.GetValueString()
	case "FillOpacity":
		rectanchoredtext.FillOpacity = value.GetValueFloat()
	case "Stroke":
		rectanchoredtext.Stroke = value.GetValueString()
	case "StrokeOpacity":
		rectanchoredtext.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		rectanchoredtext.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		rectanchoredtext.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		rectanchoredtext.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		rectanchoredtext.Transform = value.GetValueString()
	case "Animates":
		rectanchoredtext.Animates = make([]*Animate, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Animates {
					if stage.AnimateMap_Staged_Order[__instance__] == uint(id) {
						rectanchoredtext.Animates = append(rectanchoredtext.Animates, __instance__)
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

func (rectlinklink *RectLinkLink) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		rectlinklink.Name = value.GetValueString()
	case "Start":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			rectlinklink.Start = nil
			for __instance__ := range stage.Rects {
				if stage.RectMap_Staged_Order[__instance__] == uint(id) {
					rectlinklink.Start = __instance__
					break
				}
			}
		}
	case "End":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			rectlinklink.End = nil
			for __instance__ := range stage.Links {
				if stage.LinkMap_Staged_Order[__instance__] == uint(id) {
					rectlinklink.End = __instance__
					break
				}
			}
		}
	case "TargetAnchorPosition":
		rectlinklink.TargetAnchorPosition = value.GetValueFloat()
	case "Color":
		rectlinklink.Color = value.GetValueString()
	case "FillOpacity":
		rectlinklink.FillOpacity = value.GetValueFloat()
	case "Stroke":
		rectlinklink.Stroke = value.GetValueString()
	case "StrokeOpacity":
		rectlinklink.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		rectlinklink.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		rectlinklink.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		rectlinklink.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		rectlinklink.Transform = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (svg *SVG) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		svg.Name = value.GetValueString()
	case "Layers":
		svg.Layers = make([]*Layer, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Layers {
					if stage.LayerMap_Staged_Order[__instance__] == uint(id) {
						svg.Layers = append(svg.Layers, __instance__)
						break
					}
				}
			}
		}
	case "DrawingState":
		svg.DrawingState.FromCodeString(value.GetValueString())
	case "StartRect":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			svg.StartRect = nil
			for __instance__ := range stage.Rects {
				if stage.RectMap_Staged_Order[__instance__] == uint(id) {
					svg.StartRect = __instance__
					break
				}
			}
		}
	case "EndRect":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			svg.EndRect = nil
			for __instance__ := range stage.Rects {
				if stage.RectMap_Staged_Order[__instance__] == uint(id) {
					svg.EndRect = __instance__
					break
				}
			}
		}
	case "IsEditable":
		svg.IsEditable = value.GetValueBool()
	case "IsSVGFrontEndFileGenerated":
		svg.IsSVGFrontEndFileGenerated = value.GetValueBool()
	case "IsSVGBackEndFileGenerated":
		svg.IsSVGBackEndFileGenerated = value.GetValueBool()
	case "DefaultDirectoryForGeneratedImages":
		svg.DefaultDirectoryForGeneratedImages = value.GetValueString()
	case "IsControlBannerHidden":
		svg.IsControlBannerHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (svgtext *SvgText) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		svgtext.Name = value.GetValueString()
	case "Text":
		svgtext.Text = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (text *Text) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		text.Name = value.GetValueString()
	case "X":
		text.X = value.GetValueFloat()
	case "Y":
		text.Y = value.GetValueFloat()
	case "Content":
		text.Content = value.GetValueString()
	case "Color":
		text.Color = value.GetValueString()
	case "FillOpacity":
		text.FillOpacity = value.GetValueFloat()
	case "Stroke":
		text.Stroke = value.GetValueString()
	case "StrokeOpacity":
		text.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		text.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		text.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		text.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		text.Transform = value.GetValueString()
	case "FontWeight":
		text.FontWeight = value.GetValueString()
	case "FontSize":
		text.FontSize = value.GetValueString()
	case "FontStyle":
		text.FontStyle = value.GetValueString()
	case "LetterSpacing":
		text.LetterSpacing = value.GetValueString()
	case "FontFamily":
		text.FontFamily = value.GetValueString()
	case "WhiteSpace":
		text.WhiteSpace.FromCodeString(value.GetValueString())
	case "Animates":
		text.Animates = make([]*Animate, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Animates {
					if stage.AnimateMap_Staged_Order[__instance__] == uint(id) {
						text.Animates = append(text.Animates, __instance__)
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

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
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

func (layer *Layer) GongGetGongstructName() string {
	return "Layer"
}

func (line *Line) GongGetGongstructName() string {
	return "Line"
}

func (link *Link) GongGetGongstructName() string {
	return "Link"
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

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

// Last line of the template
