// generated code - do not edit
package models

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

var __GongSliceTemplate_time__dummyDeclaration time.Duration
var _ = __GongSliceTemplate_time__dummyDeclaration

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Animate
	// insertion point per field

	// Compute reverse map for named struct Circle
	// insertion point per field
	stage.Circle_Animations_reverseMap = make(map[*Animate]*Circle)
	for circle := range stage.Circles {
		_ = circle
		for _, _animate := range circle.Animations {
			stage.Circle_Animations_reverseMap[_animate] = circle
		}
	}

	// Compute reverse map for named struct Condition
	// insertion point per field

	// Compute reverse map for named struct ControlPoint
	// insertion point per field

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

	// Compute reverse map for named struct Point
	// insertion point per field

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

	// Compute reverse map for named struct RectAnchoredPath
	// insertion point per field

	// Compute reverse map for named struct RectAnchoredRect
	// insertion point per field

	// Compute reverse map for named struct RectAnchoredText
	// insertion point per field
	stage.RectAnchoredText_Animates_reverseMap = make(map[*Animate]*RectAnchoredText)
	for rectanchoredtext := range stage.RectAnchoredTexts {
		_ = rectanchoredtext
		for _, _animate := range rectanchoredtext.Animates {
			stage.RectAnchoredText_Animates_reverseMap[_animate] = rectanchoredtext
		}
	}

	// Compute reverse map for named struct RectLinkLink
	// insertion point per field

	// Compute reverse map for named struct SVG
	// insertion point per field
	stage.SVG_Layers_reverseMap = make(map[*Layer]*SVG)
	for svg := range stage.SVGs {
		_ = svg
		for _, _layer := range svg.Layers {
			stage.SVG_Layers_reverseMap[_layer] = svg
		}
	}

	// Compute reverse map for named struct SvgText
	// insertion point per field

	// Compute reverse map for named struct Text
	// insertion point per field
	stage.Text_Animates_reverseMap = make(map[*Animate]*Text)
	for text := range stage.Texts {
		_ = text
		for _, _animate := range text.Animates {
			stage.Text_Animates_reverseMap[_animate] = text
		}
	}

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Animates {
		res = append(res, instance)
	}

	for instance := range stage.Circles {
		res = append(res, instance)
	}

	for instance := range stage.Conditions {
		res = append(res, instance)
	}

	for instance := range stage.ControlPoints {
		res = append(res, instance)
	}

	for instance := range stage.Ellipses {
		res = append(res, instance)
	}

	for instance := range stage.Layers {
		res = append(res, instance)
	}

	for instance := range stage.Lines {
		res = append(res, instance)
	}

	for instance := range stage.Links {
		res = append(res, instance)
	}

	for instance := range stage.LinkAnchoredTexts {
		res = append(res, instance)
	}

	for instance := range stage.Paths {
		res = append(res, instance)
	}

	for instance := range stage.Points {
		res = append(res, instance)
	}

	for instance := range stage.Polygones {
		res = append(res, instance)
	}

	for instance := range stage.Polylines {
		res = append(res, instance)
	}

	for instance := range stage.Rects {
		res = append(res, instance)
	}

	for instance := range stage.RectAnchoredPaths {
		res = append(res, instance)
	}

	for instance := range stage.RectAnchoredRects {
		res = append(res, instance)
	}

	for instance := range stage.RectAnchoredTexts {
		res = append(res, instance)
	}

	for instance := range stage.RectLinkLinks {
		res = append(res, instance)
	}

	for instance := range stage.SVGs {
		res = append(res, instance)
	}

	for instance := range stage.SvgTexts {
		res = append(res, instance)
	}

	for instance := range stage.Texts {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (animate *Animate) GongCopy() GongstructIF {
	newInstance := *animate
	return &newInstance
}

func (circle *Circle) GongCopy() GongstructIF {
	newInstance := *circle
	return &newInstance
}

func (condition *Condition) GongCopy() GongstructIF {
	newInstance := *condition
	return &newInstance
}

func (controlpoint *ControlPoint) GongCopy() GongstructIF {
	newInstance := *controlpoint
	return &newInstance
}

func (ellipse *Ellipse) GongCopy() GongstructIF {
	newInstance := *ellipse
	return &newInstance
}

func (layer *Layer) GongCopy() GongstructIF {
	newInstance := *layer
	return &newInstance
}

func (line *Line) GongCopy() GongstructIF {
	newInstance := *line
	return &newInstance
}

func (link *Link) GongCopy() GongstructIF {
	newInstance := *link
	return &newInstance
}

func (linkanchoredtext *LinkAnchoredText) GongCopy() GongstructIF {
	newInstance := *linkanchoredtext
	return &newInstance
}

func (path *Path) GongCopy() GongstructIF {
	newInstance := *path
	return &newInstance
}

func (point *Point) GongCopy() GongstructIF {
	newInstance := *point
	return &newInstance
}

func (polygone *Polygone) GongCopy() GongstructIF {
	newInstance := *polygone
	return &newInstance
}

func (polyline *Polyline) GongCopy() GongstructIF {
	newInstance := *polyline
	return &newInstance
}

func (rect *Rect) GongCopy() GongstructIF {
	newInstance := *rect
	return &newInstance
}

func (rectanchoredpath *RectAnchoredPath) GongCopy() GongstructIF {
	newInstance := *rectanchoredpath
	return &newInstance
}

func (rectanchoredrect *RectAnchoredRect) GongCopy() GongstructIF {
	newInstance := *rectanchoredrect
	return &newInstance
}

func (rectanchoredtext *RectAnchoredText) GongCopy() GongstructIF {
	newInstance := *rectanchoredtext
	return &newInstance
}

func (rectlinklink *RectLinkLink) GongCopy() GongstructIF {
	newInstance := *rectlinklink
	return &newInstance
}

func (svg *SVG) GongCopy() GongstructIF {
	newInstance := *svg
	return &newInstance
}

func (svgtext *SvgText) GongCopy() GongstructIF {
	newInstance := *svgtext
	return &newInstance
}

func (text *Text) GongCopy() GongstructIF {
	newInstance := *text
	return &newInstance
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
	var animates_newInstances []*Animate
	var animates_deletedInstances []*Animate

	// parse all staged instances and check if they have a reference
	for animate := range stage.Animates {
		if ref, ok := stage.Animates_reference[animate]; !ok {
			animates_newInstances = append(animates_newInstances, animate)
			newInstancesSlice = append(newInstancesSlice, animate.GongMarshallIdentifier(stage))
			if stage.Animates_referenceOrder == nil {
				stage.Animates_referenceOrder = make(map[*Animate]uint)
			}
			stage.Animates_referenceOrder[animate] = stage.AnimateMap_Staged_Order[animate]
			newInstancesReverseSlice = append(newInstancesReverseSlice, animate.GongMarshallUnstaging(stage))
			delete(stage.Animates_referenceOrder, animate)
			fieldInitializers, pointersInitializations := animate.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.AnimateMap_Staged_Order[ref] = stage.AnimateMap_Staged_Order[animate]
			diffs := animate.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, animate)
			delete(stage.AnimateMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", animate.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Animates_reference {
		if _, ok := stage.Animates[ref]; !ok {
			animates_deletedInstances = append(animates_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(animates_newInstances)
	lenDeletedInstances += len(animates_deletedInstances)
	var circles_newInstances []*Circle
	var circles_deletedInstances []*Circle

	// parse all staged instances and check if they have a reference
	for circle := range stage.Circles {
		if ref, ok := stage.Circles_reference[circle]; !ok {
			circles_newInstances = append(circles_newInstances, circle)
			newInstancesSlice = append(newInstancesSlice, circle.GongMarshallIdentifier(stage))
			if stage.Circles_referenceOrder == nil {
				stage.Circles_referenceOrder = make(map[*Circle]uint)
			}
			stage.Circles_referenceOrder[circle] = stage.CircleMap_Staged_Order[circle]
			newInstancesReverseSlice = append(newInstancesReverseSlice, circle.GongMarshallUnstaging(stage))
			delete(stage.Circles_referenceOrder, circle)
			fieldInitializers, pointersInitializations := circle.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.CircleMap_Staged_Order[ref] = stage.CircleMap_Staged_Order[circle]
			diffs := circle.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, circle)
			delete(stage.CircleMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", circle.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Circles_reference {
		if _, ok := stage.Circles[ref]; !ok {
			circles_deletedInstances = append(circles_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(circles_newInstances)
	lenDeletedInstances += len(circles_deletedInstances)
	var conditions_newInstances []*Condition
	var conditions_deletedInstances []*Condition

	// parse all staged instances and check if they have a reference
	for condition := range stage.Conditions {
		if ref, ok := stage.Conditions_reference[condition]; !ok {
			conditions_newInstances = append(conditions_newInstances, condition)
			newInstancesSlice = append(newInstancesSlice, condition.GongMarshallIdentifier(stage))
			if stage.Conditions_referenceOrder == nil {
				stage.Conditions_referenceOrder = make(map[*Condition]uint)
			}
			stage.Conditions_referenceOrder[condition] = stage.ConditionMap_Staged_Order[condition]
			newInstancesReverseSlice = append(newInstancesReverseSlice, condition.GongMarshallUnstaging(stage))
			delete(stage.Conditions_referenceOrder, condition)
			fieldInitializers, pointersInitializations := condition.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ConditionMap_Staged_Order[ref] = stage.ConditionMap_Staged_Order[condition]
			diffs := condition.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, condition)
			delete(stage.ConditionMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", condition.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Conditions_reference {
		if _, ok := stage.Conditions[ref]; !ok {
			conditions_deletedInstances = append(conditions_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(conditions_newInstances)
	lenDeletedInstances += len(conditions_deletedInstances)
	var controlpoints_newInstances []*ControlPoint
	var controlpoints_deletedInstances []*ControlPoint

	// parse all staged instances and check if they have a reference
	for controlpoint := range stage.ControlPoints {
		if ref, ok := stage.ControlPoints_reference[controlpoint]; !ok {
			controlpoints_newInstances = append(controlpoints_newInstances, controlpoint)
			newInstancesSlice = append(newInstancesSlice, controlpoint.GongMarshallIdentifier(stage))
			if stage.ControlPoints_referenceOrder == nil {
				stage.ControlPoints_referenceOrder = make(map[*ControlPoint]uint)
			}
			stage.ControlPoints_referenceOrder[controlpoint] = stage.ControlPointMap_Staged_Order[controlpoint]
			newInstancesReverseSlice = append(newInstancesReverseSlice, controlpoint.GongMarshallUnstaging(stage))
			delete(stage.ControlPoints_referenceOrder, controlpoint)
			fieldInitializers, pointersInitializations := controlpoint.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ControlPointMap_Staged_Order[ref] = stage.ControlPointMap_Staged_Order[controlpoint]
			diffs := controlpoint.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, controlpoint)
			delete(stage.ControlPointMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", controlpoint.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ControlPoints_reference {
		if _, ok := stage.ControlPoints[ref]; !ok {
			controlpoints_deletedInstances = append(controlpoints_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(controlpoints_newInstances)
	lenDeletedInstances += len(controlpoints_deletedInstances)
	var ellipses_newInstances []*Ellipse
	var ellipses_deletedInstances []*Ellipse

	// parse all staged instances and check if they have a reference
	for ellipse := range stage.Ellipses {
		if ref, ok := stage.Ellipses_reference[ellipse]; !ok {
			ellipses_newInstances = append(ellipses_newInstances, ellipse)
			newInstancesSlice = append(newInstancesSlice, ellipse.GongMarshallIdentifier(stage))
			if stage.Ellipses_referenceOrder == nil {
				stage.Ellipses_referenceOrder = make(map[*Ellipse]uint)
			}
			stage.Ellipses_referenceOrder[ellipse] = stage.EllipseMap_Staged_Order[ellipse]
			newInstancesReverseSlice = append(newInstancesReverseSlice, ellipse.GongMarshallUnstaging(stage))
			delete(stage.Ellipses_referenceOrder, ellipse)
			fieldInitializers, pointersInitializations := ellipse.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.EllipseMap_Staged_Order[ref] = stage.EllipseMap_Staged_Order[ellipse]
			diffs := ellipse.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, ellipse)
			delete(stage.EllipseMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", ellipse.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Ellipses_reference {
		if _, ok := stage.Ellipses[ref]; !ok {
			ellipses_deletedInstances = append(ellipses_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(ellipses_newInstances)
	lenDeletedInstances += len(ellipses_deletedInstances)
	var layers_newInstances []*Layer
	var layers_deletedInstances []*Layer

	// parse all staged instances and check if they have a reference
	for layer := range stage.Layers {
		if ref, ok := stage.Layers_reference[layer]; !ok {
			layers_newInstances = append(layers_newInstances, layer)
			newInstancesSlice = append(newInstancesSlice, layer.GongMarshallIdentifier(stage))
			if stage.Layers_referenceOrder == nil {
				stage.Layers_referenceOrder = make(map[*Layer]uint)
			}
			stage.Layers_referenceOrder[layer] = stage.LayerMap_Staged_Order[layer]
			newInstancesReverseSlice = append(newInstancesReverseSlice, layer.GongMarshallUnstaging(stage))
			delete(stage.Layers_referenceOrder, layer)
			fieldInitializers, pointersInitializations := layer.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.LayerMap_Staged_Order[ref] = stage.LayerMap_Staged_Order[layer]
			diffs := layer.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, layer)
			delete(stage.LayerMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", layer.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Layers_reference {
		if _, ok := stage.Layers[ref]; !ok {
			layers_deletedInstances = append(layers_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(layers_newInstances)
	lenDeletedInstances += len(layers_deletedInstances)
	var lines_newInstances []*Line
	var lines_deletedInstances []*Line

	// parse all staged instances and check if they have a reference
	for line := range stage.Lines {
		if ref, ok := stage.Lines_reference[line]; !ok {
			lines_newInstances = append(lines_newInstances, line)
			newInstancesSlice = append(newInstancesSlice, line.GongMarshallIdentifier(stage))
			if stage.Lines_referenceOrder == nil {
				stage.Lines_referenceOrder = make(map[*Line]uint)
			}
			stage.Lines_referenceOrder[line] = stage.LineMap_Staged_Order[line]
			newInstancesReverseSlice = append(newInstancesReverseSlice, line.GongMarshallUnstaging(stage))
			delete(stage.Lines_referenceOrder, line)
			fieldInitializers, pointersInitializations := line.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.LineMap_Staged_Order[ref] = stage.LineMap_Staged_Order[line]
			diffs := line.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, line)
			delete(stage.LineMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", line.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Lines_reference {
		if _, ok := stage.Lines[ref]; !ok {
			lines_deletedInstances = append(lines_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(lines_newInstances)
	lenDeletedInstances += len(lines_deletedInstances)
	var links_newInstances []*Link
	var links_deletedInstances []*Link

	// parse all staged instances and check if they have a reference
	for link := range stage.Links {
		if ref, ok := stage.Links_reference[link]; !ok {
			links_newInstances = append(links_newInstances, link)
			newInstancesSlice = append(newInstancesSlice, link.GongMarshallIdentifier(stage))
			if stage.Links_referenceOrder == nil {
				stage.Links_referenceOrder = make(map[*Link]uint)
			}
			stage.Links_referenceOrder[link] = stage.LinkMap_Staged_Order[link]
			newInstancesReverseSlice = append(newInstancesReverseSlice, link.GongMarshallUnstaging(stage))
			delete(stage.Links_referenceOrder, link)
			fieldInitializers, pointersInitializations := link.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.LinkMap_Staged_Order[ref] = stage.LinkMap_Staged_Order[link]
			diffs := link.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, link)
			delete(stage.LinkMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", link.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Links_reference {
		if _, ok := stage.Links[ref]; !ok {
			links_deletedInstances = append(links_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(links_newInstances)
	lenDeletedInstances += len(links_deletedInstances)
	var linkanchoredtexts_newInstances []*LinkAnchoredText
	var linkanchoredtexts_deletedInstances []*LinkAnchoredText

	// parse all staged instances and check if they have a reference
	for linkanchoredtext := range stage.LinkAnchoredTexts {
		if ref, ok := stage.LinkAnchoredTexts_reference[linkanchoredtext]; !ok {
			linkanchoredtexts_newInstances = append(linkanchoredtexts_newInstances, linkanchoredtext)
			newInstancesSlice = append(newInstancesSlice, linkanchoredtext.GongMarshallIdentifier(stage))
			if stage.LinkAnchoredTexts_referenceOrder == nil {
				stage.LinkAnchoredTexts_referenceOrder = make(map[*LinkAnchoredText]uint)
			}
			stage.LinkAnchoredTexts_referenceOrder[linkanchoredtext] = stage.LinkAnchoredTextMap_Staged_Order[linkanchoredtext]
			newInstancesReverseSlice = append(newInstancesReverseSlice, linkanchoredtext.GongMarshallUnstaging(stage))
			delete(stage.LinkAnchoredTexts_referenceOrder, linkanchoredtext)
			fieldInitializers, pointersInitializations := linkanchoredtext.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.LinkAnchoredTextMap_Staged_Order[ref] = stage.LinkAnchoredTextMap_Staged_Order[linkanchoredtext]
			diffs := linkanchoredtext.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, linkanchoredtext)
			delete(stage.LinkAnchoredTextMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", linkanchoredtext.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.LinkAnchoredTexts_reference {
		if _, ok := stage.LinkAnchoredTexts[ref]; !ok {
			linkanchoredtexts_deletedInstances = append(linkanchoredtexts_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(linkanchoredtexts_newInstances)
	lenDeletedInstances += len(linkanchoredtexts_deletedInstances)
	var paths_newInstances []*Path
	var paths_deletedInstances []*Path

	// parse all staged instances and check if they have a reference
	for path := range stage.Paths {
		if ref, ok := stage.Paths_reference[path]; !ok {
			paths_newInstances = append(paths_newInstances, path)
			newInstancesSlice = append(newInstancesSlice, path.GongMarshallIdentifier(stage))
			if stage.Paths_referenceOrder == nil {
				stage.Paths_referenceOrder = make(map[*Path]uint)
			}
			stage.Paths_referenceOrder[path] = stage.PathMap_Staged_Order[path]
			newInstancesReverseSlice = append(newInstancesReverseSlice, path.GongMarshallUnstaging(stage))
			delete(stage.Paths_referenceOrder, path)
			fieldInitializers, pointersInitializations := path.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PathMap_Staged_Order[ref] = stage.PathMap_Staged_Order[path]
			diffs := path.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, path)
			delete(stage.PathMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", path.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Paths_reference {
		if _, ok := stage.Paths[ref]; !ok {
			paths_deletedInstances = append(paths_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(paths_newInstances)
	lenDeletedInstances += len(paths_deletedInstances)
	var points_newInstances []*Point
	var points_deletedInstances []*Point

	// parse all staged instances and check if they have a reference
	for point := range stage.Points {
		if ref, ok := stage.Points_reference[point]; !ok {
			points_newInstances = append(points_newInstances, point)
			newInstancesSlice = append(newInstancesSlice, point.GongMarshallIdentifier(stage))
			if stage.Points_referenceOrder == nil {
				stage.Points_referenceOrder = make(map[*Point]uint)
			}
			stage.Points_referenceOrder[point] = stage.PointMap_Staged_Order[point]
			newInstancesReverseSlice = append(newInstancesReverseSlice, point.GongMarshallUnstaging(stage))
			delete(stage.Points_referenceOrder, point)
			fieldInitializers, pointersInitializations := point.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PointMap_Staged_Order[ref] = stage.PointMap_Staged_Order[point]
			diffs := point.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, point)
			delete(stage.PointMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", point.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Points_reference {
		if _, ok := stage.Points[ref]; !ok {
			points_deletedInstances = append(points_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(points_newInstances)
	lenDeletedInstances += len(points_deletedInstances)
	var polygones_newInstances []*Polygone
	var polygones_deletedInstances []*Polygone

	// parse all staged instances and check if they have a reference
	for polygone := range stage.Polygones {
		if ref, ok := stage.Polygones_reference[polygone]; !ok {
			polygones_newInstances = append(polygones_newInstances, polygone)
			newInstancesSlice = append(newInstancesSlice, polygone.GongMarshallIdentifier(stage))
			if stage.Polygones_referenceOrder == nil {
				stage.Polygones_referenceOrder = make(map[*Polygone]uint)
			}
			stage.Polygones_referenceOrder[polygone] = stage.PolygoneMap_Staged_Order[polygone]
			newInstancesReverseSlice = append(newInstancesReverseSlice, polygone.GongMarshallUnstaging(stage))
			delete(stage.Polygones_referenceOrder, polygone)
			fieldInitializers, pointersInitializations := polygone.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PolygoneMap_Staged_Order[ref] = stage.PolygoneMap_Staged_Order[polygone]
			diffs := polygone.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, polygone)
			delete(stage.PolygoneMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", polygone.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Polygones_reference {
		if _, ok := stage.Polygones[ref]; !ok {
			polygones_deletedInstances = append(polygones_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(polygones_newInstances)
	lenDeletedInstances += len(polygones_deletedInstances)
	var polylines_newInstances []*Polyline
	var polylines_deletedInstances []*Polyline

	// parse all staged instances and check if they have a reference
	for polyline := range stage.Polylines {
		if ref, ok := stage.Polylines_reference[polyline]; !ok {
			polylines_newInstances = append(polylines_newInstances, polyline)
			newInstancesSlice = append(newInstancesSlice, polyline.GongMarshallIdentifier(stage))
			if stage.Polylines_referenceOrder == nil {
				stage.Polylines_referenceOrder = make(map[*Polyline]uint)
			}
			stage.Polylines_referenceOrder[polyline] = stage.PolylineMap_Staged_Order[polyline]
			newInstancesReverseSlice = append(newInstancesReverseSlice, polyline.GongMarshallUnstaging(stage))
			delete(stage.Polylines_referenceOrder, polyline)
			fieldInitializers, pointersInitializations := polyline.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PolylineMap_Staged_Order[ref] = stage.PolylineMap_Staged_Order[polyline]
			diffs := polyline.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, polyline)
			delete(stage.PolylineMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", polyline.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Polylines_reference {
		if _, ok := stage.Polylines[ref]; !ok {
			polylines_deletedInstances = append(polylines_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(polylines_newInstances)
	lenDeletedInstances += len(polylines_deletedInstances)
	var rects_newInstances []*Rect
	var rects_deletedInstances []*Rect

	// parse all staged instances and check if they have a reference
	for rect := range stage.Rects {
		if ref, ok := stage.Rects_reference[rect]; !ok {
			rects_newInstances = append(rects_newInstances, rect)
			newInstancesSlice = append(newInstancesSlice, rect.GongMarshallIdentifier(stage))
			if stage.Rects_referenceOrder == nil {
				stage.Rects_referenceOrder = make(map[*Rect]uint)
			}
			stage.Rects_referenceOrder[rect] = stage.RectMap_Staged_Order[rect]
			newInstancesReverseSlice = append(newInstancesReverseSlice, rect.GongMarshallUnstaging(stage))
			delete(stage.Rects_referenceOrder, rect)
			fieldInitializers, pointersInitializations := rect.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.RectMap_Staged_Order[ref] = stage.RectMap_Staged_Order[rect]
			diffs := rect.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, rect)
			delete(stage.RectMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", rect.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Rects_reference {
		if _, ok := stage.Rects[ref]; !ok {
			rects_deletedInstances = append(rects_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(rects_newInstances)
	lenDeletedInstances += len(rects_deletedInstances)
	var rectanchoredpaths_newInstances []*RectAnchoredPath
	var rectanchoredpaths_deletedInstances []*RectAnchoredPath

	// parse all staged instances and check if they have a reference
	for rectanchoredpath := range stage.RectAnchoredPaths {
		if ref, ok := stage.RectAnchoredPaths_reference[rectanchoredpath]; !ok {
			rectanchoredpaths_newInstances = append(rectanchoredpaths_newInstances, rectanchoredpath)
			newInstancesSlice = append(newInstancesSlice, rectanchoredpath.GongMarshallIdentifier(stage))
			if stage.RectAnchoredPaths_referenceOrder == nil {
				stage.RectAnchoredPaths_referenceOrder = make(map[*RectAnchoredPath]uint)
			}
			stage.RectAnchoredPaths_referenceOrder[rectanchoredpath] = stage.RectAnchoredPathMap_Staged_Order[rectanchoredpath]
			newInstancesReverseSlice = append(newInstancesReverseSlice, rectanchoredpath.GongMarshallUnstaging(stage))
			delete(stage.RectAnchoredPaths_referenceOrder, rectanchoredpath)
			fieldInitializers, pointersInitializations := rectanchoredpath.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.RectAnchoredPathMap_Staged_Order[ref] = stage.RectAnchoredPathMap_Staged_Order[rectanchoredpath]
			diffs := rectanchoredpath.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, rectanchoredpath)
			delete(stage.RectAnchoredPathMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", rectanchoredpath.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.RectAnchoredPaths_reference {
		if _, ok := stage.RectAnchoredPaths[ref]; !ok {
			rectanchoredpaths_deletedInstances = append(rectanchoredpaths_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(rectanchoredpaths_newInstances)
	lenDeletedInstances += len(rectanchoredpaths_deletedInstances)
	var rectanchoredrects_newInstances []*RectAnchoredRect
	var rectanchoredrects_deletedInstances []*RectAnchoredRect

	// parse all staged instances and check if they have a reference
	for rectanchoredrect := range stage.RectAnchoredRects {
		if ref, ok := stage.RectAnchoredRects_reference[rectanchoredrect]; !ok {
			rectanchoredrects_newInstances = append(rectanchoredrects_newInstances, rectanchoredrect)
			newInstancesSlice = append(newInstancesSlice, rectanchoredrect.GongMarshallIdentifier(stage))
			if stage.RectAnchoredRects_referenceOrder == nil {
				stage.RectAnchoredRects_referenceOrder = make(map[*RectAnchoredRect]uint)
			}
			stage.RectAnchoredRects_referenceOrder[rectanchoredrect] = stage.RectAnchoredRectMap_Staged_Order[rectanchoredrect]
			newInstancesReverseSlice = append(newInstancesReverseSlice, rectanchoredrect.GongMarshallUnstaging(stage))
			delete(stage.RectAnchoredRects_referenceOrder, rectanchoredrect)
			fieldInitializers, pointersInitializations := rectanchoredrect.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.RectAnchoredRectMap_Staged_Order[ref] = stage.RectAnchoredRectMap_Staged_Order[rectanchoredrect]
			diffs := rectanchoredrect.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, rectanchoredrect)
			delete(stage.RectAnchoredRectMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", rectanchoredrect.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.RectAnchoredRects_reference {
		if _, ok := stage.RectAnchoredRects[ref]; !ok {
			rectanchoredrects_deletedInstances = append(rectanchoredrects_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(rectanchoredrects_newInstances)
	lenDeletedInstances += len(rectanchoredrects_deletedInstances)
	var rectanchoredtexts_newInstances []*RectAnchoredText
	var rectanchoredtexts_deletedInstances []*RectAnchoredText

	// parse all staged instances and check if they have a reference
	for rectanchoredtext := range stage.RectAnchoredTexts {
		if ref, ok := stage.RectAnchoredTexts_reference[rectanchoredtext]; !ok {
			rectanchoredtexts_newInstances = append(rectanchoredtexts_newInstances, rectanchoredtext)
			newInstancesSlice = append(newInstancesSlice, rectanchoredtext.GongMarshallIdentifier(stage))
			if stage.RectAnchoredTexts_referenceOrder == nil {
				stage.RectAnchoredTexts_referenceOrder = make(map[*RectAnchoredText]uint)
			}
			stage.RectAnchoredTexts_referenceOrder[rectanchoredtext] = stage.RectAnchoredTextMap_Staged_Order[rectanchoredtext]
			newInstancesReverseSlice = append(newInstancesReverseSlice, rectanchoredtext.GongMarshallUnstaging(stage))
			delete(stage.RectAnchoredTexts_referenceOrder, rectanchoredtext)
			fieldInitializers, pointersInitializations := rectanchoredtext.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.RectAnchoredTextMap_Staged_Order[ref] = stage.RectAnchoredTextMap_Staged_Order[rectanchoredtext]
			diffs := rectanchoredtext.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, rectanchoredtext)
			delete(stage.RectAnchoredTextMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", rectanchoredtext.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.RectAnchoredTexts_reference {
		if _, ok := stage.RectAnchoredTexts[ref]; !ok {
			rectanchoredtexts_deletedInstances = append(rectanchoredtexts_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(rectanchoredtexts_newInstances)
	lenDeletedInstances += len(rectanchoredtexts_deletedInstances)
	var rectlinklinks_newInstances []*RectLinkLink
	var rectlinklinks_deletedInstances []*RectLinkLink

	// parse all staged instances and check if they have a reference
	for rectlinklink := range stage.RectLinkLinks {
		if ref, ok := stage.RectLinkLinks_reference[rectlinklink]; !ok {
			rectlinklinks_newInstances = append(rectlinklinks_newInstances, rectlinklink)
			newInstancesSlice = append(newInstancesSlice, rectlinklink.GongMarshallIdentifier(stage))
			if stage.RectLinkLinks_referenceOrder == nil {
				stage.RectLinkLinks_referenceOrder = make(map[*RectLinkLink]uint)
			}
			stage.RectLinkLinks_referenceOrder[rectlinklink] = stage.RectLinkLinkMap_Staged_Order[rectlinklink]
			newInstancesReverseSlice = append(newInstancesReverseSlice, rectlinklink.GongMarshallUnstaging(stage))
			delete(stage.RectLinkLinks_referenceOrder, rectlinklink)
			fieldInitializers, pointersInitializations := rectlinklink.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.RectLinkLinkMap_Staged_Order[ref] = stage.RectLinkLinkMap_Staged_Order[rectlinklink]
			diffs := rectlinklink.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, rectlinklink)
			delete(stage.RectLinkLinkMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", rectlinklink.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.RectLinkLinks_reference {
		if _, ok := stage.RectLinkLinks[ref]; !ok {
			rectlinklinks_deletedInstances = append(rectlinklinks_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(rectlinklinks_newInstances)
	lenDeletedInstances += len(rectlinklinks_deletedInstances)
	var svgs_newInstances []*SVG
	var svgs_deletedInstances []*SVG

	// parse all staged instances and check if they have a reference
	for svg := range stage.SVGs {
		if ref, ok := stage.SVGs_reference[svg]; !ok {
			svgs_newInstances = append(svgs_newInstances, svg)
			newInstancesSlice = append(newInstancesSlice, svg.GongMarshallIdentifier(stage))
			if stage.SVGs_referenceOrder == nil {
				stage.SVGs_referenceOrder = make(map[*SVG]uint)
			}
			stage.SVGs_referenceOrder[svg] = stage.SVGMap_Staged_Order[svg]
			newInstancesReverseSlice = append(newInstancesReverseSlice, svg.GongMarshallUnstaging(stage))
			delete(stage.SVGs_referenceOrder, svg)
			fieldInitializers, pointersInitializations := svg.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SVGMap_Staged_Order[ref] = stage.SVGMap_Staged_Order[svg]
			diffs := svg.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, svg)
			delete(stage.SVGMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", svg.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.SVGs_reference {
		if _, ok := stage.SVGs[ref]; !ok {
			svgs_deletedInstances = append(svgs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(svgs_newInstances)
	lenDeletedInstances += len(svgs_deletedInstances)
	var svgtexts_newInstances []*SvgText
	var svgtexts_deletedInstances []*SvgText

	// parse all staged instances and check if they have a reference
	for svgtext := range stage.SvgTexts {
		if ref, ok := stage.SvgTexts_reference[svgtext]; !ok {
			svgtexts_newInstances = append(svgtexts_newInstances, svgtext)
			newInstancesSlice = append(newInstancesSlice, svgtext.GongMarshallIdentifier(stage))
			if stage.SvgTexts_referenceOrder == nil {
				stage.SvgTexts_referenceOrder = make(map[*SvgText]uint)
			}
			stage.SvgTexts_referenceOrder[svgtext] = stage.SvgTextMap_Staged_Order[svgtext]
			newInstancesReverseSlice = append(newInstancesReverseSlice, svgtext.GongMarshallUnstaging(stage))
			delete(stage.SvgTexts_referenceOrder, svgtext)
			fieldInitializers, pointersInitializations := svgtext.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SvgTextMap_Staged_Order[ref] = stage.SvgTextMap_Staged_Order[svgtext]
			diffs := svgtext.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, svgtext)
			delete(stage.SvgTextMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", svgtext.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.SvgTexts_reference {
		if _, ok := stage.SvgTexts[ref]; !ok {
			svgtexts_deletedInstances = append(svgtexts_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(svgtexts_newInstances)
	lenDeletedInstances += len(svgtexts_deletedInstances)
	var texts_newInstances []*Text
	var texts_deletedInstances []*Text

	// parse all staged instances and check if they have a reference
	for text := range stage.Texts {
		if ref, ok := stage.Texts_reference[text]; !ok {
			texts_newInstances = append(texts_newInstances, text)
			newInstancesSlice = append(newInstancesSlice, text.GongMarshallIdentifier(stage))
			if stage.Texts_referenceOrder == nil {
				stage.Texts_referenceOrder = make(map[*Text]uint)
			}
			stage.Texts_referenceOrder[text] = stage.TextMap_Staged_Order[text]
			newInstancesReverseSlice = append(newInstancesReverseSlice, text.GongMarshallUnstaging(stage))
			delete(stage.Texts_referenceOrder, text)
			fieldInitializers, pointersInitializations := text.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TextMap_Staged_Order[ref] = stage.TextMap_Staged_Order[text]
			diffs := text.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, text)
			delete(stage.TextMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", text.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Texts_reference {
		if _, ok := stage.Texts[ref]; !ok {
			texts_deletedInstances = append(texts_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(texts_newInstances)
	lenDeletedInstances += len(texts_deletedInstances)

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

		if stage.GetProbeIF() != nil {
			var mergedCommits string
			for _, commit := range stage.forwardCommits {
				mergedCommits += commit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Forward commits:\n"+
					mergedCommits,
			)

			var reverseMergedCommits string
			for _, reverserCommit := range stage.backwardCommits {
				reverseMergedCommits += reverserCommit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Backward commits:\n"+
					reverseMergedCommits,
			)

			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Animates_reference = make(map[*Animate]*Animate)
	stage.Animates_referenceOrder = make(map[*Animate]uint) // diff Unstage needs the reference order
	for instance := range stage.Animates {
		stage.Animates_reference[instance] = instance.GongCopy().(*Animate)
		stage.Animates_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Circles_reference = make(map[*Circle]*Circle)
	stage.Circles_referenceOrder = make(map[*Circle]uint) // diff Unstage needs the reference order
	for instance := range stage.Circles {
		stage.Circles_reference[instance] = instance.GongCopy().(*Circle)
		stage.Circles_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Conditions_reference = make(map[*Condition]*Condition)
	stage.Conditions_referenceOrder = make(map[*Condition]uint) // diff Unstage needs the reference order
	for instance := range stage.Conditions {
		stage.Conditions_reference[instance] = instance.GongCopy().(*Condition)
		stage.Conditions_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ControlPoints_reference = make(map[*ControlPoint]*ControlPoint)
	stage.ControlPoints_referenceOrder = make(map[*ControlPoint]uint) // diff Unstage needs the reference order
	for instance := range stage.ControlPoints {
		stage.ControlPoints_reference[instance] = instance.GongCopy().(*ControlPoint)
		stage.ControlPoints_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Ellipses_reference = make(map[*Ellipse]*Ellipse)
	stage.Ellipses_referenceOrder = make(map[*Ellipse]uint) // diff Unstage needs the reference order
	for instance := range stage.Ellipses {
		stage.Ellipses_reference[instance] = instance.GongCopy().(*Ellipse)
		stage.Ellipses_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Layers_reference = make(map[*Layer]*Layer)
	stage.Layers_referenceOrder = make(map[*Layer]uint) // diff Unstage needs the reference order
	for instance := range stage.Layers {
		stage.Layers_reference[instance] = instance.GongCopy().(*Layer)
		stage.Layers_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Lines_reference = make(map[*Line]*Line)
	stage.Lines_referenceOrder = make(map[*Line]uint) // diff Unstage needs the reference order
	for instance := range stage.Lines {
		stage.Lines_reference[instance] = instance.GongCopy().(*Line)
		stage.Lines_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Links_reference = make(map[*Link]*Link)
	stage.Links_referenceOrder = make(map[*Link]uint) // diff Unstage needs the reference order
	for instance := range stage.Links {
		stage.Links_reference[instance] = instance.GongCopy().(*Link)
		stage.Links_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.LinkAnchoredTexts_reference = make(map[*LinkAnchoredText]*LinkAnchoredText)
	stage.LinkAnchoredTexts_referenceOrder = make(map[*LinkAnchoredText]uint) // diff Unstage needs the reference order
	for instance := range stage.LinkAnchoredTexts {
		stage.LinkAnchoredTexts_reference[instance] = instance.GongCopy().(*LinkAnchoredText)
		stage.LinkAnchoredTexts_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Paths_reference = make(map[*Path]*Path)
	stage.Paths_referenceOrder = make(map[*Path]uint) // diff Unstage needs the reference order
	for instance := range stage.Paths {
		stage.Paths_reference[instance] = instance.GongCopy().(*Path)
		stage.Paths_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Points_reference = make(map[*Point]*Point)
	stage.Points_referenceOrder = make(map[*Point]uint) // diff Unstage needs the reference order
	for instance := range stage.Points {
		stage.Points_reference[instance] = instance.GongCopy().(*Point)
		stage.Points_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Polygones_reference = make(map[*Polygone]*Polygone)
	stage.Polygones_referenceOrder = make(map[*Polygone]uint) // diff Unstage needs the reference order
	for instance := range stage.Polygones {
		stage.Polygones_reference[instance] = instance.GongCopy().(*Polygone)
		stage.Polygones_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Polylines_reference = make(map[*Polyline]*Polyline)
	stage.Polylines_referenceOrder = make(map[*Polyline]uint) // diff Unstage needs the reference order
	for instance := range stage.Polylines {
		stage.Polylines_reference[instance] = instance.GongCopy().(*Polyline)
		stage.Polylines_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Rects_reference = make(map[*Rect]*Rect)
	stage.Rects_referenceOrder = make(map[*Rect]uint) // diff Unstage needs the reference order
	for instance := range stage.Rects {
		stage.Rects_reference[instance] = instance.GongCopy().(*Rect)
		stage.Rects_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.RectAnchoredPaths_reference = make(map[*RectAnchoredPath]*RectAnchoredPath)
	stage.RectAnchoredPaths_referenceOrder = make(map[*RectAnchoredPath]uint) // diff Unstage needs the reference order
	for instance := range stage.RectAnchoredPaths {
		stage.RectAnchoredPaths_reference[instance] = instance.GongCopy().(*RectAnchoredPath)
		stage.RectAnchoredPaths_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.RectAnchoredRects_reference = make(map[*RectAnchoredRect]*RectAnchoredRect)
	stage.RectAnchoredRects_referenceOrder = make(map[*RectAnchoredRect]uint) // diff Unstage needs the reference order
	for instance := range stage.RectAnchoredRects {
		stage.RectAnchoredRects_reference[instance] = instance.GongCopy().(*RectAnchoredRect)
		stage.RectAnchoredRects_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.RectAnchoredTexts_reference = make(map[*RectAnchoredText]*RectAnchoredText)
	stage.RectAnchoredTexts_referenceOrder = make(map[*RectAnchoredText]uint) // diff Unstage needs the reference order
	for instance := range stage.RectAnchoredTexts {
		stage.RectAnchoredTexts_reference[instance] = instance.GongCopy().(*RectAnchoredText)
		stage.RectAnchoredTexts_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.RectLinkLinks_reference = make(map[*RectLinkLink]*RectLinkLink)
	stage.RectLinkLinks_referenceOrder = make(map[*RectLinkLink]uint) // diff Unstage needs the reference order
	for instance := range stage.RectLinkLinks {
		stage.RectLinkLinks_reference[instance] = instance.GongCopy().(*RectLinkLink)
		stage.RectLinkLinks_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SVGs_reference = make(map[*SVG]*SVG)
	stage.SVGs_referenceOrder = make(map[*SVG]uint) // diff Unstage needs the reference order
	for instance := range stage.SVGs {
		stage.SVGs_reference[instance] = instance.GongCopy().(*SVG)
		stage.SVGs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SvgTexts_reference = make(map[*SvgText]*SvgText)
	stage.SvgTexts_referenceOrder = make(map[*SvgText]uint) // diff Unstage needs the reference order
	for instance := range stage.SvgTexts {
		stage.SvgTexts_reference[instance] = instance.GongCopy().(*SvgText)
		stage.SvgTexts_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Texts_reference = make(map[*Text]*Text)
	stage.Texts_referenceOrder = make(map[*Text]uint) // diff Unstage needs the reference order
	for instance := range stage.Texts {
		stage.Texts_reference[instance] = instance.GongCopy().(*Text)
		stage.Texts_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (animate *Animate) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.AnimateMap_Staged_Order[animate]; ok {
		return order
	}
	return stage.Animates_referenceOrder[animate]
}

func (animate *Animate) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Animates_referenceOrder[animate]
}

func (circle *Circle) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.CircleMap_Staged_Order[circle]; ok {
		return order
	}
	return stage.Circles_referenceOrder[circle]
}

func (circle *Circle) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Circles_referenceOrder[circle]
}

func (condition *Condition) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ConditionMap_Staged_Order[condition]; ok {
		return order
	}
	return stage.Conditions_referenceOrder[condition]
}

func (condition *Condition) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Conditions_referenceOrder[condition]
}

func (controlpoint *ControlPoint) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ControlPointMap_Staged_Order[controlpoint]; ok {
		return order
	}
	return stage.ControlPoints_referenceOrder[controlpoint]
}

func (controlpoint *ControlPoint) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ControlPoints_referenceOrder[controlpoint]
}

func (ellipse *Ellipse) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EllipseMap_Staged_Order[ellipse]; ok {
		return order
	}
	return stage.Ellipses_referenceOrder[ellipse]
}

func (ellipse *Ellipse) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Ellipses_referenceOrder[ellipse]
}

func (layer *Layer) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.LayerMap_Staged_Order[layer]; ok {
		return order
	}
	return stage.Layers_referenceOrder[layer]
}

func (layer *Layer) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Layers_referenceOrder[layer]
}

func (line *Line) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.LineMap_Staged_Order[line]; ok {
		return order
	}
	return stage.Lines_referenceOrder[line]
}

func (line *Line) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Lines_referenceOrder[line]
}

func (link *Link) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.LinkMap_Staged_Order[link]; ok {
		return order
	}
	return stage.Links_referenceOrder[link]
}

func (link *Link) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Links_referenceOrder[link]
}

func (linkanchoredtext *LinkAnchoredText) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.LinkAnchoredTextMap_Staged_Order[linkanchoredtext]; ok {
		return order
	}
	return stage.LinkAnchoredTexts_referenceOrder[linkanchoredtext]
}

func (linkanchoredtext *LinkAnchoredText) GongGetReferenceOrder(stage *Stage) uint {
	return stage.LinkAnchoredTexts_referenceOrder[linkanchoredtext]
}

func (path *Path) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PathMap_Staged_Order[path]; ok {
		return order
	}
	return stage.Paths_referenceOrder[path]
}

func (path *Path) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Paths_referenceOrder[path]
}

func (point *Point) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PointMap_Staged_Order[point]; ok {
		return order
	}
	return stage.Points_referenceOrder[point]
}

func (point *Point) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Points_referenceOrder[point]
}

func (polygone *Polygone) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PolygoneMap_Staged_Order[polygone]; ok {
		return order
	}
	return stage.Polygones_referenceOrder[polygone]
}

func (polygone *Polygone) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Polygones_referenceOrder[polygone]
}

func (polyline *Polyline) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PolylineMap_Staged_Order[polyline]; ok {
		return order
	}
	return stage.Polylines_referenceOrder[polyline]
}

func (polyline *Polyline) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Polylines_referenceOrder[polyline]
}

func (rect *Rect) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RectMap_Staged_Order[rect]; ok {
		return order
	}
	return stage.Rects_referenceOrder[rect]
}

func (rect *Rect) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Rects_referenceOrder[rect]
}

func (rectanchoredpath *RectAnchoredPath) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RectAnchoredPathMap_Staged_Order[rectanchoredpath]; ok {
		return order
	}
	return stage.RectAnchoredPaths_referenceOrder[rectanchoredpath]
}

func (rectanchoredpath *RectAnchoredPath) GongGetReferenceOrder(stage *Stage) uint {
	return stage.RectAnchoredPaths_referenceOrder[rectanchoredpath]
}

func (rectanchoredrect *RectAnchoredRect) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RectAnchoredRectMap_Staged_Order[rectanchoredrect]; ok {
		return order
	}
	return stage.RectAnchoredRects_referenceOrder[rectanchoredrect]
}

func (rectanchoredrect *RectAnchoredRect) GongGetReferenceOrder(stage *Stage) uint {
	return stage.RectAnchoredRects_referenceOrder[rectanchoredrect]
}

func (rectanchoredtext *RectAnchoredText) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RectAnchoredTextMap_Staged_Order[rectanchoredtext]; ok {
		return order
	}
	return stage.RectAnchoredTexts_referenceOrder[rectanchoredtext]
}

func (rectanchoredtext *RectAnchoredText) GongGetReferenceOrder(stage *Stage) uint {
	return stage.RectAnchoredTexts_referenceOrder[rectanchoredtext]
}

func (rectlinklink *RectLinkLink) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RectLinkLinkMap_Staged_Order[rectlinklink]; ok {
		return order
	}
	return stage.RectLinkLinks_referenceOrder[rectlinklink]
}

func (rectlinklink *RectLinkLink) GongGetReferenceOrder(stage *Stage) uint {
	return stage.RectLinkLinks_referenceOrder[rectlinklink]
}

func (svg *SVG) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SVGMap_Staged_Order[svg]; ok {
		return order
	}
	return stage.SVGs_referenceOrder[svg]
}

func (svg *SVG) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SVGs_referenceOrder[svg]
}

func (svgtext *SvgText) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SvgTextMap_Staged_Order[svgtext]; ok {
		return order
	}
	return stage.SvgTexts_referenceOrder[svgtext]
}

func (svgtext *SvgText) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SvgTexts_referenceOrder[svgtext]
}

func (text *Text) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TextMap_Staged_Order[text]; ok {
		return order
	}
	return stage.Texts_referenceOrder[text]
}

func (text *Text) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Texts_referenceOrder[text]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (animate *Animate) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", animate.GongGetGongstructName(), animate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (animate *Animate) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", animate.GongGetGongstructName(), animate.GongGetReferenceOrder(stage))
}

func (circle *Circle) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", circle.GongGetGongstructName(), circle.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (circle *Circle) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", circle.GongGetGongstructName(), circle.GongGetReferenceOrder(stage))
}

func (condition *Condition) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", condition.GongGetGongstructName(), condition.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (condition *Condition) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", condition.GongGetGongstructName(), condition.GongGetReferenceOrder(stage))
}

func (controlpoint *ControlPoint) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", controlpoint.GongGetGongstructName(), controlpoint.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (controlpoint *ControlPoint) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", controlpoint.GongGetGongstructName(), controlpoint.GongGetReferenceOrder(stage))
}

func (ellipse *Ellipse) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", ellipse.GongGetGongstructName(), ellipse.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ellipse *Ellipse) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", ellipse.GongGetGongstructName(), ellipse.GongGetReferenceOrder(stage))
}

func (layer *Layer) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", layer.GongGetGongstructName(), layer.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (layer *Layer) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", layer.GongGetGongstructName(), layer.GongGetReferenceOrder(stage))
}

func (line *Line) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", line.GongGetGongstructName(), line.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (line *Line) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", line.GongGetGongstructName(), line.GongGetReferenceOrder(stage))
}

func (link *Link) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", link.GongGetGongstructName(), link.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (link *Link) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", link.GongGetGongstructName(), link.GongGetReferenceOrder(stage))
}

func (linkanchoredtext *LinkAnchoredText) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", linkanchoredtext.GongGetGongstructName(), linkanchoredtext.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (linkanchoredtext *LinkAnchoredText) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", linkanchoredtext.GongGetGongstructName(), linkanchoredtext.GongGetReferenceOrder(stage))
}

func (path *Path) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", path.GongGetGongstructName(), path.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (path *Path) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", path.GongGetGongstructName(), path.GongGetReferenceOrder(stage))
}

func (point *Point) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", point.GongGetGongstructName(), point.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (point *Point) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", point.GongGetGongstructName(), point.GongGetReferenceOrder(stage))
}

func (polygone *Polygone) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", polygone.GongGetGongstructName(), polygone.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (polygone *Polygone) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", polygone.GongGetGongstructName(), polygone.GongGetReferenceOrder(stage))
}

func (polyline *Polyline) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", polyline.GongGetGongstructName(), polyline.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (polyline *Polyline) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", polyline.GongGetGongstructName(), polyline.GongGetReferenceOrder(stage))
}

func (rect *Rect) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rect.GongGetGongstructName(), rect.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rect *Rect) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rect.GongGetGongstructName(), rect.GongGetReferenceOrder(stage))
}

func (rectanchoredpath *RectAnchoredPath) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rectanchoredpath.GongGetGongstructName(), rectanchoredpath.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rectanchoredpath *RectAnchoredPath) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rectanchoredpath.GongGetGongstructName(), rectanchoredpath.GongGetReferenceOrder(stage))
}

func (rectanchoredrect *RectAnchoredRect) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rectanchoredrect.GongGetGongstructName(), rectanchoredrect.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rectanchoredrect *RectAnchoredRect) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rectanchoredrect.GongGetGongstructName(), rectanchoredrect.GongGetReferenceOrder(stage))
}

func (rectanchoredtext *RectAnchoredText) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rectanchoredtext.GongGetGongstructName(), rectanchoredtext.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rectanchoredtext *RectAnchoredText) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rectanchoredtext.GongGetGongstructName(), rectanchoredtext.GongGetReferenceOrder(stage))
}

func (rectlinklink *RectLinkLink) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rectlinklink.GongGetGongstructName(), rectlinklink.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rectlinklink *RectLinkLink) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rectlinklink.GongGetGongstructName(), rectlinklink.GongGetReferenceOrder(stage))
}

func (svg *SVG) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svg.GongGetGongstructName(), svg.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (svg *SVG) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svg.GongGetGongstructName(), svg.GongGetReferenceOrder(stage))
}

func (svgtext *SvgText) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svgtext.GongGetGongstructName(), svgtext.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (svgtext *SvgText) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svgtext.GongGetGongstructName(), svgtext.GongGetReferenceOrder(stage))
}

func (text *Text) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", text.GongGetGongstructName(), text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (text *Text) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", text.GongGetGongstructName(), text.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (animate *Animate) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", animate.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Animate")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", animate.Name)
	return
}
func (circle *Circle) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", circle.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Circle")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", circle.Name)
	return
}
func (condition *Condition) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", condition.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Condition")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", condition.Name)
	return
}
func (controlpoint *ControlPoint) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", controlpoint.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ControlPoint")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", controlpoint.Name)
	return
}
func (ellipse *Ellipse) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Ellipse")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ellipse.Name)
	return
}
func (layer *Layer) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", layer.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Layer")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", layer.Name)
	return
}
func (line *Line) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", line.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Line")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", line.Name)
	return
}
func (link *Link) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", link.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Link")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", link.Name)
	return
}
func (linkanchoredtext *LinkAnchoredText) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LinkAnchoredText")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", linkanchoredtext.Name)
	return
}
func (path *Path) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", path.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Path")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", path.Name)
	return
}
func (point *Point) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", point.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Point")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", point.Name)
	return
}
func (polygone *Polygone) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", polygone.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Polygone")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", polygone.Name)
	return
}
func (polyline *Polyline) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", polyline.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Polyline")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", polyline.Name)
	return
}
func (rect *Rect) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rect.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Rect")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", rect.Name)
	return
}
func (rectanchoredpath *RectAnchoredPath) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RectAnchoredPath")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", rectanchoredpath.Name)
	return
}
func (rectanchoredrect *RectAnchoredRect) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RectAnchoredRect")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", rectanchoredrect.Name)
	return
}
func (rectanchoredtext *RectAnchoredText) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RectAnchoredText")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", rectanchoredtext.Name)
	return
}
func (rectlinklink *RectLinkLink) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RectLinkLink")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", rectlinklink.Name)
	return
}
func (svg *SVG) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", svg.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SVG")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", svg.Name)
	return
}
func (svgtext *SvgText) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", svgtext.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SvgText")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", svgtext.Name)
	return
}
func (text *Text) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", text.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Text")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", text.Name)
	return
}

// insertion point for unstaging
func (animate *Animate) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", animate.GongGetReferenceIdentifier(stage))
	return
}
func (circle *Circle) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", circle.GongGetReferenceIdentifier(stage))
	return
}
func (condition *Condition) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", condition.GongGetReferenceIdentifier(stage))
	return
}
func (controlpoint *ControlPoint) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", controlpoint.GongGetReferenceIdentifier(stage))
	return
}
func (ellipse *Ellipse) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", ellipse.GongGetReferenceIdentifier(stage))
	return
}
func (layer *Layer) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", layer.GongGetReferenceIdentifier(stage))
	return
}
func (line *Line) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", line.GongGetReferenceIdentifier(stage))
	return
}
func (link *Link) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", link.GongGetReferenceIdentifier(stage))
	return
}
func (linkanchoredtext *LinkAnchoredText) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", linkanchoredtext.GongGetReferenceIdentifier(stage))
	return
}
func (path *Path) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", path.GongGetReferenceIdentifier(stage))
	return
}
func (point *Point) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", point.GongGetReferenceIdentifier(stage))
	return
}
func (polygone *Polygone) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", polygone.GongGetReferenceIdentifier(stage))
	return
}
func (polyline *Polyline) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", polyline.GongGetReferenceIdentifier(stage))
	return
}
func (rect *Rect) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rect.GongGetReferenceIdentifier(stage))
	return
}
func (rectanchoredpath *RectAnchoredPath) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rectanchoredpath.GongGetReferenceIdentifier(stage))
	return
}
func (rectanchoredrect *RectAnchoredRect) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rectanchoredrect.GongGetReferenceIdentifier(stage))
	return
}
func (rectanchoredtext *RectAnchoredText) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rectanchoredtext.GongGetReferenceIdentifier(stage))
	return
}
func (rectlinklink *RectLinkLink) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rectlinklink.GongGetReferenceIdentifier(stage))
	return
}
func (svg *SVG) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", svg.GongGetReferenceIdentifier(stage))
	return
}
func (svgtext *SvgText) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", svgtext.GongGetReferenceIdentifier(stage))
	return
}
func (text *Text) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", text.GongGetReferenceIdentifier(stage))
	return
}
