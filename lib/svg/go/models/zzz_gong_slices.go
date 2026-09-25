// generated code - do not edit
package models

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Circle
	// insertion point per field
	stage.Circle_Animations_reverseMap = make(map[*Animate]*Circle)
	for circle := range stage.Circles {
		_ = circle
		for _, _animate := range circle.Animations {
			stage.Circle_Animations_reverseMap[_animate] = circle
		}
	}

	// Compute reverse map for named struct Ellipse
	// insertion point per field
	stage.Ellipse_Animates_reverseMap = make(map[*Animate]*Ellipse)
	for ellipse := range stage.Ellipses {
		_ = ellipse
		for _, _animate := range ellipse.Animates {
			stage.Ellipse_Animates_reverseMap[_animate] = ellipse
		}
	}

	// Compute reverse map for named struct Layer
	// insertion point per field
	stage.Layer_Rects_reverseMap = make(map[*Rect]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _rect := range layer.Rects {
			stage.Layer_Rects_reverseMap[_rect] = layer
		}
	}
	stage.Layer_Texts_reverseMap = make(map[*Text]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _text := range layer.Texts {
			stage.Layer_Texts_reverseMap[_text] = layer
		}
	}
	stage.Layer_Circles_reverseMap = make(map[*Circle]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _circle := range layer.Circles {
			stage.Layer_Circles_reverseMap[_circle] = layer
		}
	}
	stage.Layer_Lines_reverseMap = make(map[*Line]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _line := range layer.Lines {
			stage.Layer_Lines_reverseMap[_line] = layer
		}
	}
	stage.Layer_Ellipses_reverseMap = make(map[*Ellipse]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _ellipse := range layer.Ellipses {
			stage.Layer_Ellipses_reverseMap[_ellipse] = layer
		}
	}
	stage.Layer_Polylines_reverseMap = make(map[*Polyline]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _polyline := range layer.Polylines {
			stage.Layer_Polylines_reverseMap[_polyline] = layer
		}
	}
	stage.Layer_Polygones_reverseMap = make(map[*Polygone]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _polygone := range layer.Polygones {
			stage.Layer_Polygones_reverseMap[_polygone] = layer
		}
	}
	stage.Layer_Paths_reverseMap = make(map[*Path]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _path := range layer.Paths {
			stage.Layer_Paths_reverseMap[_path] = layer
		}
	}
	stage.Layer_Links_reverseMap = make(map[*Link]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _link := range layer.Links {
			stage.Layer_Links_reverseMap[_link] = layer
		}
	}
	stage.Layer_RectLinkLinks_reverseMap = make(map[*RectLinkLink]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _rectlinklink := range layer.RectLinkLinks {
			stage.Layer_RectLinkLinks_reverseMap[_rectlinklink] = layer
		}
	}

	// Compute reverse map for named struct Line
	// insertion point per field
	stage.Line_Animates_reverseMap = make(map[*Animate]*Line)
	for line := range stage.Lines {
		_ = line
		for _, _animate := range line.Animates {
			stage.Line_Animates_reverseMap[_animate] = line
		}
	}

	// Compute reverse map for named struct Link
	// insertion point per field
	stage.Link_TextAtArrowStart_reverseMap = make(map[*LinkAnchoredText]*Link)
	for link := range stage.Links {
		_ = link
		for _, _linkanchoredtext := range link.TextAtArrowStart {
			stage.Link_TextAtArrowStart_reverseMap[_linkanchoredtext] = link
		}
	}
	stage.Link_TextAtArrowEnd_reverseMap = make(map[*LinkAnchoredText]*Link)
	for link := range stage.Links {
		_ = link
		for _, _linkanchoredtext := range link.TextAtArrowEnd {
			stage.Link_TextAtArrowEnd_reverseMap[_linkanchoredtext] = link
		}
	}
	stage.Link_TextAtCorner_reverseMap = make(map[*LinkAnchoredText]*Link)
	for link := range stage.Links {
		_ = link
		for _, _linkanchoredtext := range link.TextAtCorner {
			stage.Link_TextAtCorner_reverseMap[_linkanchoredtext] = link
		}
	}
	stage.Link_PathAtArrowStart_reverseMap = make(map[*LinkAnchoredPath]*Link)
	for link := range stage.Links {
		_ = link
		for _, _linkanchoredpath := range link.PathAtArrowStart {
			stage.Link_PathAtArrowStart_reverseMap[_linkanchoredpath] = link
		}
	}
	stage.Link_PathAtArrowEnd_reverseMap = make(map[*LinkAnchoredPath]*Link)
	for link := range stage.Links {
		_ = link
		for _, _linkanchoredpath := range link.PathAtArrowEnd {
			stage.Link_PathAtArrowEnd_reverseMap[_linkanchoredpath] = link
		}
	}
	stage.Link_PathAtCorner_reverseMap = make(map[*LinkAnchoredPath]*Link)
	for link := range stage.Links {
		_ = link
		for _, _linkanchoredpath := range link.PathAtCorner {
			stage.Link_PathAtCorner_reverseMap[_linkanchoredpath] = link
		}
	}
	stage.Link_ControlPoints_reverseMap = make(map[*ControlPoint]*Link)
	for link := range stage.Links {
		_ = link
		for _, _controlpoint := range link.ControlPoints {
			stage.Link_ControlPoints_reverseMap[_controlpoint] = link
		}
	}

	// Compute reverse map for named struct LinkAnchoredText
	// insertion point per field
	stage.LinkAnchoredText_Animates_reverseMap = make(map[*Animate]*LinkAnchoredText)
	for linkanchoredtext := range stage.LinkAnchoredTexts {
		_ = linkanchoredtext
		for _, _animate := range linkanchoredtext.Animates {
			stage.LinkAnchoredText_Animates_reverseMap[_animate] = linkanchoredtext
		}
	}

	// Compute reverse map for named struct Path
	// insertion point per field
	stage.Path_Animates_reverseMap = make(map[*Animate]*Path)
	for path := range stage.Paths {
		_ = path
		for _, _animate := range path.Animates {
			stage.Path_Animates_reverseMap[_animate] = path
		}
	}

	// Compute reverse map for named struct Polygone
	// insertion point per field
	stage.Polygone_Animates_reverseMap = make(map[*Animate]*Polygone)
	for polygone := range stage.Polygones {
		_ = polygone
		for _, _animate := range polygone.Animates {
			stage.Polygone_Animates_reverseMap[_animate] = polygone
		}
	}

	// Compute reverse map for named struct Polyline
	// insertion point per field
	stage.Polyline_Animates_reverseMap = make(map[*Animate]*Polyline)
	for polyline := range stage.Polylines {
		_ = polyline
		for _, _animate := range polyline.Animates {
			stage.Polyline_Animates_reverseMap[_animate] = polyline
		}
	}

	// Compute reverse map for named struct Rect
	// insertion point per field
	stage.Rect_Peers_reverseMap = make(map[*Rect]*Rect)
	for rect := range stage.Rects {
		_ = rect
		for _, _rect := range rect.Peers {
			stage.Rect_Peers_reverseMap[_rect] = rect
		}
	}
	stage.Rect_Obstacles_reverseMap = make(map[*Rect]*Rect)
	for rect := range stage.Rects {
		_ = rect
		for _, _rect := range rect.Obstacles {
			stage.Rect_Obstacles_reverseMap[_rect] = rect
		}
	}
	stage.Rect_HoveringTrigger_reverseMap = make(map[*Condition]*Rect)
	for rect := range stage.Rects {
		_ = rect
		for _, _condition := range rect.HoveringTrigger {
			stage.Rect_HoveringTrigger_reverseMap[_condition] = rect
		}
	}
	stage.Rect_DisplayConditions_reverseMap = make(map[*Condition]*Rect)
	for rect := range stage.Rects {
		_ = rect
		for _, _condition := range rect.DisplayConditions {
			stage.Rect_DisplayConditions_reverseMap[_condition] = rect
		}
	}
	stage.Rect_Animations_reverseMap = make(map[*Animate]*Rect)
	for rect := range stage.Rects {
		_ = rect
		for _, _animate := range rect.Animations {
			stage.Rect_Animations_reverseMap[_animate] = rect
		}
	}
	stage.Rect_RectAnchoredTexts_reverseMap = make(map[*RectAnchoredText]*Rect)
	for rect := range stage.Rects {
		_ = rect
		for _, _rectanchoredtext := range rect.RectAnchoredTexts {
			stage.Rect_RectAnchoredTexts_reverseMap[_rectanchoredtext] = rect
		}
	}
	stage.Rect_RectAnchoredRects_reverseMap = make(map[*RectAnchoredRect]*Rect)
	for rect := range stage.Rects {
		_ = rect
		for _, _rectanchoredrect := range rect.RectAnchoredRects {
			stage.Rect_RectAnchoredRects_reverseMap[_rectanchoredrect] = rect
		}
	}
	stage.Rect_RectAnchoredPaths_reverseMap = make(map[*RectAnchoredPath]*Rect)
	for rect := range stage.Rects {
		_ = rect
		for _, _rectanchoredpath := range rect.RectAnchoredPaths {
			stage.Rect_RectAnchoredPaths_reverseMap[_rectanchoredpath] = rect
		}
	}
	stage.Rect_RectAnchoredPngImages_reverseMap = make(map[*RectAnchoredPngImage]*Rect)
	for rect := range stage.Rects {
		_ = rect
		for _, _rectanchoredpngimage := range rect.RectAnchoredPngImages {
			stage.Rect_RectAnchoredPngImages_reverseMap[_rectanchoredpngimage] = rect
		}
	}

	// Compute reverse map for named struct RectAnchoredText
	// insertion point per field
	stage.RectAnchoredText_Animates_reverseMap = make(map[*Animate]*RectAnchoredText)
	for rectanchoredtext := range stage.RectAnchoredTexts {
		_ = rectanchoredtext
		for _, _animate := range rectanchoredtext.Animates {
			stage.RectAnchoredText_Animates_reverseMap[_animate] = rectanchoredtext
		}
	}

	// Compute reverse map for named struct SVG
	// insertion point per field
	stage.SVG_Layers_reverseMap = make(map[*Layer]*SVG)
	for svg := range stage.SVGs {
		_ = svg
		for _, _layer := range svg.Layers {
			stage.SVG_Layers_reverseMap[_layer] = svg
		}
	}

	// Compute reverse map for named struct Text
	// insertion point per field
	stage.Text_Animates_reverseMap = make(map[*Animate]*Text)
	for text := range stage.Texts {
		_ = text
		for _, _animate := range text.Animates {
			stage.Text_Animates_reverseMap[_animate] = text
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.Animates)

	res = __gong__appendInstances(res, stage.Circles)

	res = __gong__appendInstances(res, stage.Conditions)

	res = __gong__appendInstances(res, stage.ControlPoints)

	res = __gong__appendInstances(res, stage.Ellipses)

	res = __gong__appendInstances(res, stage.FileToDownloads)

	res = __gong__appendInstances(res, stage.Layers)

	res = __gong__appendInstances(res, stage.Lines)

	res = __gong__appendInstances(res, stage.Links)

	res = __gong__appendInstances(res, stage.LinkAnchoredPaths)

	res = __gong__appendInstances(res, stage.LinkAnchoredTexts)

	res = __gong__appendInstances(res, stage.Paths)

	res = __gong__appendInstances(res, stage.Points)

	res = __gong__appendInstances(res, stage.Polygones)

	res = __gong__appendInstances(res, stage.Polylines)

	res = __gong__appendInstances(res, stage.Rects)

	res = __gong__appendInstances(res, stage.RectAnchoredPaths)

	res = __gong__appendInstances(res, stage.RectAnchoredPngImages)

	res = __gong__appendInstances(res, stage.RectAnchoredRects)

	res = __gong__appendInstances(res, stage.RectAnchoredTexts)

	res = __gong__appendInstances(res, stage.RectLinkLinks)

	res = __gong__appendInstances(res, stage.SVGs)

	res = __gong__appendInstances(res, stage.SvgTexts)

	res = __gong__appendInstances(res, stage.Texts)

	return
}

// insertion point per named struct
func (animate *Animate) GongCopy() GongstructIF {
	newInstance := new(Animate)
	animate.GongCopyBasicFields(newInstance)
	return newInstance
}

func (circle *Circle) GongCopy() GongstructIF {
	newInstance := new(Circle)
	circle.GongCopyBasicFields(newInstance)
	return newInstance
}

func (condition *Condition) GongCopy() GongstructIF {
	newInstance := new(Condition)
	condition.GongCopyBasicFields(newInstance)
	return newInstance
}

func (controlpoint *ControlPoint) GongCopy() GongstructIF {
	newInstance := new(ControlPoint)
	controlpoint.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ellipse *Ellipse) GongCopy() GongstructIF {
	newInstance := new(Ellipse)
	ellipse.GongCopyBasicFields(newInstance)
	return newInstance
}

func (filetodownload *FileToDownload) GongCopy() GongstructIF {
	newInstance := new(FileToDownload)
	filetodownload.GongCopyBasicFields(newInstance)
	return newInstance
}

func (layer *Layer) GongCopy() GongstructIF {
	newInstance := new(Layer)
	layer.GongCopyBasicFields(newInstance)
	return newInstance
}

func (line *Line) GongCopy() GongstructIF {
	newInstance := new(Line)
	line.GongCopyBasicFields(newInstance)
	return newInstance
}

func (link *Link) GongCopy() GongstructIF {
	newInstance := new(Link)
	link.GongCopyBasicFields(newInstance)
	return newInstance
}

func (linkanchoredpath *LinkAnchoredPath) GongCopy() GongstructIF {
	newInstance := new(LinkAnchoredPath)
	linkanchoredpath.GongCopyBasicFields(newInstance)
	return newInstance
}

func (linkanchoredtext *LinkAnchoredText) GongCopy() GongstructIF {
	newInstance := new(LinkAnchoredText)
	linkanchoredtext.GongCopyBasicFields(newInstance)
	return newInstance
}

func (path *Path) GongCopy() GongstructIF {
	newInstance := new(Path)
	path.GongCopyBasicFields(newInstance)
	return newInstance
}

func (point *Point) GongCopy() GongstructIF {
	newInstance := new(Point)
	point.GongCopyBasicFields(newInstance)
	return newInstance
}

func (polygone *Polygone) GongCopy() GongstructIF {
	newInstance := new(Polygone)
	polygone.GongCopyBasicFields(newInstance)
	return newInstance
}

func (polyline *Polyline) GongCopy() GongstructIF {
	newInstance := new(Polyline)
	polyline.GongCopyBasicFields(newInstance)
	return newInstance
}

func (rect *Rect) GongCopy() GongstructIF {
	newInstance := new(Rect)
	rect.GongCopyBasicFields(newInstance)
	return newInstance
}

func (rectanchoredpath *RectAnchoredPath) GongCopy() GongstructIF {
	newInstance := new(RectAnchoredPath)
	rectanchoredpath.GongCopyBasicFields(newInstance)
	return newInstance
}

func (rectanchoredpngimage *RectAnchoredPngImage) GongCopy() GongstructIF {
	newInstance := new(RectAnchoredPngImage)
	rectanchoredpngimage.GongCopyBasicFields(newInstance)
	return newInstance
}

func (rectanchoredrect *RectAnchoredRect) GongCopy() GongstructIF {
	newInstance := new(RectAnchoredRect)
	rectanchoredrect.GongCopyBasicFields(newInstance)
	return newInstance
}

func (rectanchoredtext *RectAnchoredText) GongCopy() GongstructIF {
	newInstance := new(RectAnchoredText)
	rectanchoredtext.GongCopyBasicFields(newInstance)
	return newInstance
}

func (rectlinklink *RectLinkLink) GongCopy() GongstructIF {
	newInstance := new(RectLinkLink)
	rectlinklink.GongCopyBasicFields(newInstance)
	return newInstance
}

func (svg *SVG) GongCopy() GongstructIF {
	newInstance := new(SVG)
	svg.GongCopyBasicFields(newInstance)
	return newInstance
}

func (svgtext *SvgText) GongCopy() GongstructIF {
	newInstance := new(SvgText)
	svgtext.GongCopyBasicFields(newInstance)
	return newInstance
}

func (text *Text) GongCopy() GongstructIF {
	newInstance := new(Text)
	text.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (animate *Animate) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, animate)
}

func (circle *Circle) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, circle)
}

func (condition *Condition) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, condition)
}

func (controlpoint *ControlPoint) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, controlpoint)
}

func (ellipse *Ellipse) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ellipse)
}

func (filetodownload *FileToDownload) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, filetodownload)
}

func (layer *Layer) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, layer)
}

func (line *Line) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, line)
}

func (link *Link) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, link)
}

func (linkanchoredpath *LinkAnchoredPath) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, linkanchoredpath)
}

func (linkanchoredtext *LinkAnchoredText) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, linkanchoredtext)
}

func (path *Path) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, path)
}

func (point *Point) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, point)
}

func (polygone *Polygone) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, polygone)
}

func (polyline *Polyline) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, polyline)
}

func (rect *Rect) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, rect)
}

func (rectanchoredpath *RectAnchoredPath) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, rectanchoredpath)
}

func (rectanchoredpngimage *RectAnchoredPngImage) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, rectanchoredpngimage)
}

func (rectanchoredrect *RectAnchoredRect) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, rectanchoredrect)
}

func (rectanchoredtext *RectAnchoredText) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, rectanchoredtext)
}

func (rectlinklink *RectLinkLink) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, rectlinklink)
}

func (svg *SVG) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, svg)
}

func (svgtext *SvgText) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, svgtext)
}

func (text *Text) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, text)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	computeCommitsForType(
		stage,
		stage.Animates,
		stage.Animate_stagedOrder,
		stage.Animates_reference,
		&stage.Animates_referenceOrder,
		stage.Animates_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Circles,
		stage.Circle_stagedOrder,
		stage.Circles_reference,
		&stage.Circles_referenceOrder,
		stage.Circles_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Conditions,
		stage.Condition_stagedOrder,
		stage.Conditions_reference,
		&stage.Conditions_referenceOrder,
		stage.Conditions_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ControlPoints,
		stage.ControlPoint_stagedOrder,
		stage.ControlPoints_reference,
		&stage.ControlPoints_referenceOrder,
		stage.ControlPoints_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Ellipses,
		stage.Ellipse_stagedOrder,
		stage.Ellipses_reference,
		&stage.Ellipses_referenceOrder,
		stage.Ellipses_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FileToDownloads,
		stage.FileToDownload_stagedOrder,
		stage.FileToDownloads_reference,
		&stage.FileToDownloads_referenceOrder,
		stage.FileToDownloads_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Layers,
		stage.Layer_stagedOrder,
		stage.Layers_reference,
		&stage.Layers_referenceOrder,
		stage.Layers_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Lines,
		stage.Line_stagedOrder,
		stage.Lines_reference,
		&stage.Lines_referenceOrder,
		stage.Lines_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Links,
		stage.Link_stagedOrder,
		stage.Links_reference,
		&stage.Links_referenceOrder,
		stage.Links_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.LinkAnchoredPaths,
		stage.LinkAnchoredPath_stagedOrder,
		stage.LinkAnchoredPaths_reference,
		&stage.LinkAnchoredPaths_referenceOrder,
		stage.LinkAnchoredPaths_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.LinkAnchoredTexts,
		stage.LinkAnchoredText_stagedOrder,
		stage.LinkAnchoredTexts_reference,
		&stage.LinkAnchoredTexts_referenceOrder,
		stage.LinkAnchoredTexts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Paths,
		stage.Path_stagedOrder,
		stage.Paths_reference,
		&stage.Paths_referenceOrder,
		stage.Paths_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Points,
		stage.Point_stagedOrder,
		stage.Points_reference,
		&stage.Points_referenceOrder,
		stage.Points_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Polygones,
		stage.Polygone_stagedOrder,
		stage.Polygones_reference,
		&stage.Polygones_referenceOrder,
		stage.Polygones_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Polylines,
		stage.Polyline_stagedOrder,
		stage.Polylines_reference,
		&stage.Polylines_referenceOrder,
		stage.Polylines_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Rects,
		stage.Rect_stagedOrder,
		stage.Rects_reference,
		&stage.Rects_referenceOrder,
		stage.Rects_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.RectAnchoredPaths,
		stage.RectAnchoredPath_stagedOrder,
		stage.RectAnchoredPaths_reference,
		&stage.RectAnchoredPaths_referenceOrder,
		stage.RectAnchoredPaths_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.RectAnchoredPngImages,
		stage.RectAnchoredPngImage_stagedOrder,
		stage.RectAnchoredPngImages_reference,
		&stage.RectAnchoredPngImages_referenceOrder,
		stage.RectAnchoredPngImages_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.RectAnchoredRects,
		stage.RectAnchoredRect_stagedOrder,
		stage.RectAnchoredRects_reference,
		&stage.RectAnchoredRects_referenceOrder,
		stage.RectAnchoredRects_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.RectAnchoredTexts,
		stage.RectAnchoredText_stagedOrder,
		stage.RectAnchoredTexts_reference,
		&stage.RectAnchoredTexts_referenceOrder,
		stage.RectAnchoredTexts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.RectLinkLinks,
		stage.RectLinkLink_stagedOrder,
		stage.RectLinkLinks_reference,
		&stage.RectLinkLinks_referenceOrder,
		stage.RectLinkLinks_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SVGs,
		stage.SVG_stagedOrder,
		stage.SVGs_reference,
		&stage.SVGs_referenceOrder,
		stage.SVGs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SvgTexts,
		stage.SvgText_stagedOrder,
		stage.SvgTexts_reference,
		&stage.SvgTexts_referenceOrder,
		stage.SvgTexts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Texts,
		stage.Text_stagedOrder,
		stage.Texts_reference,
		&stage.Texts_referenceOrder,
		stage.Texts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)
		stage.modified = true
	} else {
		stage.modified = false
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.Animates, &stage.Animates_reference, &stage.Animates_referenceOrder, &stage.Animates_instance)

	__gong__computeReferencePass1(stage, stage.Circles, &stage.Circles_reference, &stage.Circles_referenceOrder, &stage.Circles_instance)

	__gong__computeReferencePass1(stage, stage.Conditions, &stage.Conditions_reference, &stage.Conditions_referenceOrder, &stage.Conditions_instance)

	__gong__computeReferencePass1(stage, stage.ControlPoints, &stage.ControlPoints_reference, &stage.ControlPoints_referenceOrder, &stage.ControlPoints_instance)

	__gong__computeReferencePass1(stage, stage.Ellipses, &stage.Ellipses_reference, &stage.Ellipses_referenceOrder, &stage.Ellipses_instance)

	__gong__computeReferencePass1(stage, stage.FileToDownloads, &stage.FileToDownloads_reference, &stage.FileToDownloads_referenceOrder, &stage.FileToDownloads_instance)

	__gong__computeReferencePass1(stage, stage.Layers, &stage.Layers_reference, &stage.Layers_referenceOrder, &stage.Layers_instance)

	__gong__computeReferencePass1(stage, stage.Lines, &stage.Lines_reference, &stage.Lines_referenceOrder, &stage.Lines_instance)

	__gong__computeReferencePass1(stage, stage.Links, &stage.Links_reference, &stage.Links_referenceOrder, &stage.Links_instance)

	__gong__computeReferencePass1(stage, stage.LinkAnchoredPaths, &stage.LinkAnchoredPaths_reference, &stage.LinkAnchoredPaths_referenceOrder, &stage.LinkAnchoredPaths_instance)

	__gong__computeReferencePass1(stage, stage.LinkAnchoredTexts, &stage.LinkAnchoredTexts_reference, &stage.LinkAnchoredTexts_referenceOrder, &stage.LinkAnchoredTexts_instance)

	__gong__computeReferencePass1(stage, stage.Paths, &stage.Paths_reference, &stage.Paths_referenceOrder, &stage.Paths_instance)

	__gong__computeReferencePass1(stage, stage.Points, &stage.Points_reference, &stage.Points_referenceOrder, &stage.Points_instance)

	__gong__computeReferencePass1(stage, stage.Polygones, &stage.Polygones_reference, &stage.Polygones_referenceOrder, &stage.Polygones_instance)

	__gong__computeReferencePass1(stage, stage.Polylines, &stage.Polylines_reference, &stage.Polylines_referenceOrder, &stage.Polylines_instance)

	__gong__computeReferencePass1(stage, stage.Rects, &stage.Rects_reference, &stage.Rects_referenceOrder, &stage.Rects_instance)

	__gong__computeReferencePass1(stage, stage.RectAnchoredPaths, &stage.RectAnchoredPaths_reference, &stage.RectAnchoredPaths_referenceOrder, &stage.RectAnchoredPaths_instance)

	__gong__computeReferencePass1(stage, stage.RectAnchoredPngImages, &stage.RectAnchoredPngImages_reference, &stage.RectAnchoredPngImages_referenceOrder, &stage.RectAnchoredPngImages_instance)

	__gong__computeReferencePass1(stage, stage.RectAnchoredRects, &stage.RectAnchoredRects_reference, &stage.RectAnchoredRects_referenceOrder, &stage.RectAnchoredRects_instance)

	__gong__computeReferencePass1(stage, stage.RectAnchoredTexts, &stage.RectAnchoredTexts_reference, &stage.RectAnchoredTexts_referenceOrder, &stage.RectAnchoredTexts_instance)

	__gong__computeReferencePass1(stage, stage.RectLinkLinks, &stage.RectLinkLinks_reference, &stage.RectLinkLinks_referenceOrder, &stage.RectLinkLinks_instance)

	__gong__computeReferencePass1(stage, stage.SVGs, &stage.SVGs_reference, &stage.SVGs_referenceOrder, &stage.SVGs_instance)

	__gong__computeReferencePass1(stage, stage.SvgTexts, &stage.SvgTexts_reference, &stage.SvgTexts_referenceOrder, &stage.SvgTexts_instance)

	__gong__computeReferencePass1(stage, stage.Texts, &stage.Texts_reference, &stage.Texts_referenceOrder, &stage.Texts_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.Animates, stage.Animates_reference, stage)

	__gong__computeReferencePass2(stage.Circles, stage.Circles_reference, stage)

	__gong__computeReferencePass2(stage.Conditions, stage.Conditions_reference, stage)

	__gong__computeReferencePass2(stage.ControlPoints, stage.ControlPoints_reference, stage)

	__gong__computeReferencePass2(stage.Ellipses, stage.Ellipses_reference, stage)

	__gong__computeReferencePass2(stage.FileToDownloads, stage.FileToDownloads_reference, stage)

	__gong__computeReferencePass2(stage.Layers, stage.Layers_reference, stage)

	__gong__computeReferencePass2(stage.Lines, stage.Lines_reference, stage)

	__gong__computeReferencePass2(stage.Links, stage.Links_reference, stage)

	__gong__computeReferencePass2(stage.LinkAnchoredPaths, stage.LinkAnchoredPaths_reference, stage)

	__gong__computeReferencePass2(stage.LinkAnchoredTexts, stage.LinkAnchoredTexts_reference, stage)

	__gong__computeReferencePass2(stage.Paths, stage.Paths_reference, stage)

	__gong__computeReferencePass2(stage.Points, stage.Points_reference, stage)

	__gong__computeReferencePass2(stage.Polygones, stage.Polygones_reference, stage)

	__gong__computeReferencePass2(stage.Polylines, stage.Polylines_reference, stage)

	__gong__computeReferencePass2(stage.Rects, stage.Rects_reference, stage)

	__gong__computeReferencePass2(stage.RectAnchoredPaths, stage.RectAnchoredPaths_reference, stage)

	__gong__computeReferencePass2(stage.RectAnchoredPngImages, stage.RectAnchoredPngImages_reference, stage)

	__gong__computeReferencePass2(stage.RectAnchoredRects, stage.RectAnchoredRects_reference, stage)

	__gong__computeReferencePass2(stage.RectAnchoredTexts, stage.RectAnchoredTexts_reference, stage)

	__gong__computeReferencePass2(stage.RectLinkLinks, stage.RectLinkLinks_reference, stage)

	__gong__computeReferencePass2(stage.SVGs, stage.SVGs_reference, stage)

	__gong__computeReferencePass2(stage.SvgTexts, stage.SvgTexts_reference, stage)

	__gong__computeReferencePass2(stage.Texts, stage.Texts_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (animate *Animate) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Animate_stagedOrder, stage.Animates_referenceOrder, animate, "Animate")
}

func (circle *Circle) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Circle_stagedOrder, stage.Circles_referenceOrder, circle, "Circle")
}

func (condition *Condition) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Condition_stagedOrder, stage.Conditions_referenceOrder, condition, "Condition")
}

func (controlpoint *ControlPoint) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ControlPoint_stagedOrder, stage.ControlPoints_referenceOrder, controlpoint, "ControlPoint")
}

func (ellipse *Ellipse) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Ellipse_stagedOrder, stage.Ellipses_referenceOrder, ellipse, "Ellipse")
}

func (filetodownload *FileToDownload) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FileToDownload_stagedOrder, stage.FileToDownloads_referenceOrder, filetodownload, "FileToDownload")
}

func (layer *Layer) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Layer_stagedOrder, stage.Layers_referenceOrder, layer, "Layer")
}

func (line *Line) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Line_stagedOrder, stage.Lines_referenceOrder, line, "Line")
}

func (link *Link) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Link_stagedOrder, stage.Links_referenceOrder, link, "Link")
}

func (linkanchoredpath *LinkAnchoredPath) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LinkAnchoredPath_stagedOrder, stage.LinkAnchoredPaths_referenceOrder, linkanchoredpath, "LinkAnchoredPath")
}

func (linkanchoredtext *LinkAnchoredText) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LinkAnchoredText_stagedOrder, stage.LinkAnchoredTexts_referenceOrder, linkanchoredtext, "LinkAnchoredText")
}

func (path *Path) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Path_stagedOrder, stage.Paths_referenceOrder, path, "Path")
}

func (point *Point) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Point_stagedOrder, stage.Points_referenceOrder, point, "Point")
}

func (polygone *Polygone) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Polygone_stagedOrder, stage.Polygones_referenceOrder, polygone, "Polygone")
}

func (polyline *Polyline) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Polyline_stagedOrder, stage.Polylines_referenceOrder, polyline, "Polyline")
}

func (rect *Rect) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Rect_stagedOrder, stage.Rects_referenceOrder, rect, "Rect")
}

func (rectanchoredpath *RectAnchoredPath) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RectAnchoredPath_stagedOrder, stage.RectAnchoredPaths_referenceOrder, rectanchoredpath, "RectAnchoredPath")
}

func (rectanchoredpngimage *RectAnchoredPngImage) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RectAnchoredPngImage_stagedOrder, stage.RectAnchoredPngImages_referenceOrder, rectanchoredpngimage, "RectAnchoredPngImage")
}

func (rectanchoredrect *RectAnchoredRect) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RectAnchoredRect_stagedOrder, stage.RectAnchoredRects_referenceOrder, rectanchoredrect, "RectAnchoredRect")
}

func (rectanchoredtext *RectAnchoredText) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RectAnchoredText_stagedOrder, stage.RectAnchoredTexts_referenceOrder, rectanchoredtext, "RectAnchoredText")
}

func (rectlinklink *RectLinkLink) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RectLinkLink_stagedOrder, stage.RectLinkLinks_referenceOrder, rectlinklink, "RectLinkLink")
}

func (svg *SVG) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SVG_stagedOrder, stage.SVGs_referenceOrder, svg, "SVG")
}

func (svgtext *SvgText) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SvgText_stagedOrder, stage.SvgTexts_referenceOrder, svgtext, "SvgText")
}

func (text *Text) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Text_stagedOrder, stage.Texts_referenceOrder, text, "Text")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (animate *Animate) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(animate, animate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (animate *Animate) GongGetReferenceIdentifier(stage *Stage) string {
	return animate.GongGetIdentifier(stage)
}

func (circle *Circle) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(circle, circle.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (circle *Circle) GongGetReferenceIdentifier(stage *Stage) string {
	return circle.GongGetIdentifier(stage)
}

func (condition *Condition) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(condition, condition.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (condition *Condition) GongGetReferenceIdentifier(stage *Stage) string {
	return condition.GongGetIdentifier(stage)
}

func (controlpoint *ControlPoint) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(controlpoint, controlpoint.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (controlpoint *ControlPoint) GongGetReferenceIdentifier(stage *Stage) string {
	return controlpoint.GongGetIdentifier(stage)
}

func (ellipse *Ellipse) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ellipse, ellipse.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ellipse *Ellipse) GongGetReferenceIdentifier(stage *Stage) string {
	return ellipse.GongGetIdentifier(stage)
}

func (filetodownload *FileToDownload) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(filetodownload, filetodownload.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (filetodownload *FileToDownload) GongGetReferenceIdentifier(stage *Stage) string {
	return filetodownload.GongGetIdentifier(stage)
}

func (layer *Layer) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(layer, layer.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (layer *Layer) GongGetReferenceIdentifier(stage *Stage) string {
	return layer.GongGetIdentifier(stage)
}

func (line *Line) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(line, line.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (line *Line) GongGetReferenceIdentifier(stage *Stage) string {
	return line.GongGetIdentifier(stage)
}

func (link *Link) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(link, link.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (link *Link) GongGetReferenceIdentifier(stage *Stage) string {
	return link.GongGetIdentifier(stage)
}

func (linkanchoredpath *LinkAnchoredPath) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(linkanchoredpath, linkanchoredpath.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (linkanchoredpath *LinkAnchoredPath) GongGetReferenceIdentifier(stage *Stage) string {
	return linkanchoredpath.GongGetIdentifier(stage)
}

func (linkanchoredtext *LinkAnchoredText) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(linkanchoredtext, linkanchoredtext.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (linkanchoredtext *LinkAnchoredText) GongGetReferenceIdentifier(stage *Stage) string {
	return linkanchoredtext.GongGetIdentifier(stage)
}

func (path *Path) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(path, path.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (path *Path) GongGetReferenceIdentifier(stage *Stage) string {
	return path.GongGetIdentifier(stage)
}

func (point *Point) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(point, point.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (point *Point) GongGetReferenceIdentifier(stage *Stage) string {
	return point.GongGetIdentifier(stage)
}

func (polygone *Polygone) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(polygone, polygone.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (polygone *Polygone) GongGetReferenceIdentifier(stage *Stage) string {
	return polygone.GongGetIdentifier(stage)
}

func (polyline *Polyline) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(polyline, polyline.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (polyline *Polyline) GongGetReferenceIdentifier(stage *Stage) string {
	return polyline.GongGetIdentifier(stage)
}

func (rect *Rect) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(rect, rect.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rect *Rect) GongGetReferenceIdentifier(stage *Stage) string {
	return rect.GongGetIdentifier(stage)
}

func (rectanchoredpath *RectAnchoredPath) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(rectanchoredpath, rectanchoredpath.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rectanchoredpath *RectAnchoredPath) GongGetReferenceIdentifier(stage *Stage) string {
	return rectanchoredpath.GongGetIdentifier(stage)
}

func (rectanchoredpngimage *RectAnchoredPngImage) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(rectanchoredpngimage, rectanchoredpngimage.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rectanchoredpngimage *RectAnchoredPngImage) GongGetReferenceIdentifier(stage *Stage) string {
	return rectanchoredpngimage.GongGetIdentifier(stage)
}

func (rectanchoredrect *RectAnchoredRect) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(rectanchoredrect, rectanchoredrect.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rectanchoredrect *RectAnchoredRect) GongGetReferenceIdentifier(stage *Stage) string {
	return rectanchoredrect.GongGetIdentifier(stage)
}

func (rectanchoredtext *RectAnchoredText) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(rectanchoredtext, rectanchoredtext.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rectanchoredtext *RectAnchoredText) GongGetReferenceIdentifier(stage *Stage) string {
	return rectanchoredtext.GongGetIdentifier(stage)
}

func (rectlinklink *RectLinkLink) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(rectlinklink, rectlinklink.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rectlinklink *RectLinkLink) GongGetReferenceIdentifier(stage *Stage) string {
	return rectlinklink.GongGetIdentifier(stage)
}

func (svg *SVG) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(svg, svg.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (svg *SVG) GongGetReferenceIdentifier(stage *Stage) string {
	return svg.GongGetIdentifier(stage)
}

func (svgtext *SvgText) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(svgtext, svgtext.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (svgtext *SvgText) GongGetReferenceIdentifier(stage *Stage) string {
	return svgtext.GongGetIdentifier(stage)
}

func (text *Text) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(text, text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (text *Text) GongGetReferenceIdentifier(stage *Stage) string {
	return text.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (animate *Animate) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(animate.GongGetIdentifier(stage), "Animate", animate.Name)
}

func (circle *Circle) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(circle.GongGetIdentifier(stage), "Circle", circle.Name)
}

func (condition *Condition) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(condition.GongGetIdentifier(stage), "Condition", condition.Name)
}

func (controlpoint *ControlPoint) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(controlpoint.GongGetIdentifier(stage), "ControlPoint", controlpoint.Name)
}

func (ellipse *Ellipse) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ellipse.GongGetIdentifier(stage), "Ellipse", ellipse.Name)
}

func (filetodownload *FileToDownload) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(filetodownload.GongGetIdentifier(stage), "FileToDownload", filetodownload.Name)
}

func (layer *Layer) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(layer.GongGetIdentifier(stage), "Layer", layer.Name)
}

func (line *Line) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(line.GongGetIdentifier(stage), "Line", line.Name)
}

func (link *Link) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(link.GongGetIdentifier(stage), "Link", link.Name)
}

func (linkanchoredpath *LinkAnchoredPath) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(linkanchoredpath.GongGetIdentifier(stage), "LinkAnchoredPath", linkanchoredpath.Name)
}

func (linkanchoredtext *LinkAnchoredText) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(linkanchoredtext.GongGetIdentifier(stage), "LinkAnchoredText", linkanchoredtext.Name)
}

func (path *Path) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(path.GongGetIdentifier(stage), "Path", path.Name)
}

func (point *Point) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(point.GongGetIdentifier(stage), "Point", point.Name)
}

func (polygone *Polygone) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(polygone.GongGetIdentifier(stage), "Polygone", polygone.Name)
}

func (polyline *Polyline) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(polyline.GongGetIdentifier(stage), "Polyline", polyline.Name)
}

func (rect *Rect) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(rect.GongGetIdentifier(stage), "Rect", rect.Name)
}

func (rectanchoredpath *RectAnchoredPath) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(rectanchoredpath.GongGetIdentifier(stage), "RectAnchoredPath", rectanchoredpath.Name)
}

func (rectanchoredpngimage *RectAnchoredPngImage) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(rectanchoredpngimage.GongGetIdentifier(stage), "RectAnchoredPngImage", rectanchoredpngimage.Name)
}

func (rectanchoredrect *RectAnchoredRect) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(rectanchoredrect.GongGetIdentifier(stage), "RectAnchoredRect", rectanchoredrect.Name)
}

func (rectanchoredtext *RectAnchoredText) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(rectanchoredtext.GongGetIdentifier(stage), "RectAnchoredText", rectanchoredtext.Name)
}

func (rectlinklink *RectLinkLink) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(rectlinklink.GongGetIdentifier(stage), "RectLinkLink", rectlinklink.Name)
}

func (svg *SVG) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(svg.GongGetIdentifier(stage), "SVG", svg.Name)
}

func (svgtext *SvgText) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(svgtext.GongGetIdentifier(stage), "SvgText", svgtext.Name)
}

func (text *Text) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(text.GongGetIdentifier(stage), "Text", text.Name)
}

// insertion point for unstaging
func (animate *Animate) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(animate.GongGetReferenceIdentifier(stage))
}

func (circle *Circle) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(circle.GongGetReferenceIdentifier(stage))
}

func (condition *Condition) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(condition.GongGetReferenceIdentifier(stage))
}

func (controlpoint *ControlPoint) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(controlpoint.GongGetReferenceIdentifier(stage))
}

func (ellipse *Ellipse) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ellipse.GongGetReferenceIdentifier(stage))
}

func (filetodownload *FileToDownload) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(filetodownload.GongGetReferenceIdentifier(stage))
}

func (layer *Layer) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(layer.GongGetReferenceIdentifier(stage))
}

func (line *Line) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(line.GongGetReferenceIdentifier(stage))
}

func (link *Link) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(link.GongGetReferenceIdentifier(stage))
}

func (linkanchoredpath *LinkAnchoredPath) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(linkanchoredpath.GongGetReferenceIdentifier(stage))
}

func (linkanchoredtext *LinkAnchoredText) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(linkanchoredtext.GongGetReferenceIdentifier(stage))
}

func (path *Path) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(path.GongGetReferenceIdentifier(stage))
}

func (point *Point) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(point.GongGetReferenceIdentifier(stage))
}

func (polygone *Polygone) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(polygone.GongGetReferenceIdentifier(stage))
}

func (polyline *Polyline) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(polyline.GongGetReferenceIdentifier(stage))
}

func (rect *Rect) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(rect.GongGetReferenceIdentifier(stage))
}

func (rectanchoredpath *RectAnchoredPath) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(rectanchoredpath.GongGetReferenceIdentifier(stage))
}

func (rectanchoredpngimage *RectAnchoredPngImage) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(rectanchoredpngimage.GongGetReferenceIdentifier(stage))
}

func (rectanchoredrect *RectAnchoredRect) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(rectanchoredrect.GongGetReferenceIdentifier(stage))
}

func (rectanchoredtext *RectAnchoredText) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(rectanchoredtext.GongGetReferenceIdentifier(stage))
}

func (rectlinklink *RectLinkLink) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(rectlinklink.GongGetReferenceIdentifier(stage))
}

func (svg *SVG) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(svg.GongGetReferenceIdentifier(stage))
}

func (svgtext *SvgText) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(svgtext.GongGetReferenceIdentifier(stage))
}

func (text *Text) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(text.GongGetReferenceIdentifier(stage))
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
