// generated code - do not edit
package models

import (
	"fmt"
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

func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesStmt string
	_ = newInstancesStmt
	var fieldsEditStmt string
	_ = fieldsEditStmt
	var deletedInstancesStmt string
	_ = deletedInstancesStmt

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
			newInstancesStmt += animate.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := animate.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := animate.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", animate.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for animate := range stage.Animates_reference {
		if _, ok := stage.Animates[animate]; !ok {
			animates_deletedInstances = append(animates_deletedInstances, animate)
			deletedInstancesStmt += animate.GongMarshallUnstaging(stage)
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
			newInstancesStmt += circle.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := circle.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := circle.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", circle.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for circle := range stage.Circles_reference {
		if _, ok := stage.Circles[circle]; !ok {
			circles_deletedInstances = append(circles_deletedInstances, circle)
			deletedInstancesStmt += circle.GongMarshallUnstaging(stage)
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
			newInstancesStmt += condition.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := condition.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := condition.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", condition.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for condition := range stage.Conditions_reference {
		if _, ok := stage.Conditions[condition]; !ok {
			conditions_deletedInstances = append(conditions_deletedInstances, condition)
			deletedInstancesStmt += condition.GongMarshallUnstaging(stage)
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
			newInstancesStmt += controlpoint.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := controlpoint.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := controlpoint.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", controlpoint.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for controlpoint := range stage.ControlPoints_reference {
		if _, ok := stage.ControlPoints[controlpoint]; !ok {
			controlpoints_deletedInstances = append(controlpoints_deletedInstances, controlpoint)
			deletedInstancesStmt += controlpoint.GongMarshallUnstaging(stage)
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
			newInstancesStmt += ellipse.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ellipse.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := ellipse.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", ellipse.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ellipse := range stage.Ellipses_reference {
		if _, ok := stage.Ellipses[ellipse]; !ok {
			ellipses_deletedInstances = append(ellipses_deletedInstances, ellipse)
			deletedInstancesStmt += ellipse.GongMarshallUnstaging(stage)
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
			newInstancesStmt += layer.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := layer.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := layer.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", layer.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for layer := range stage.Layers_reference {
		if _, ok := stage.Layers[layer]; !ok {
			layers_deletedInstances = append(layers_deletedInstances, layer)
			deletedInstancesStmt += layer.GongMarshallUnstaging(stage)
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
			newInstancesStmt += line.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := line.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := line.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", line.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for line := range stage.Lines_reference {
		if _, ok := stage.Lines[line]; !ok {
			lines_deletedInstances = append(lines_deletedInstances, line)
			deletedInstancesStmt += line.GongMarshallUnstaging(stage)
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
			newInstancesStmt += link.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := link.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := link.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", link.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for link := range stage.Links_reference {
		if _, ok := stage.Links[link]; !ok {
			links_deletedInstances = append(links_deletedInstances, link)
			deletedInstancesStmt += link.GongMarshallUnstaging(stage)
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
			newInstancesStmt += linkanchoredtext.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := linkanchoredtext.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := linkanchoredtext.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", linkanchoredtext.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for linkanchoredtext := range stage.LinkAnchoredTexts_reference {
		if _, ok := stage.LinkAnchoredTexts[linkanchoredtext]; !ok {
			linkanchoredtexts_deletedInstances = append(linkanchoredtexts_deletedInstances, linkanchoredtext)
			deletedInstancesStmt += linkanchoredtext.GongMarshallUnstaging(stage)
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
			newInstancesStmt += path.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := path.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := path.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", path.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for path := range stage.Paths_reference {
		if _, ok := stage.Paths[path]; !ok {
			paths_deletedInstances = append(paths_deletedInstances, path)
			deletedInstancesStmt += path.GongMarshallUnstaging(stage)
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
			newInstancesStmt += point.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := point.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := point.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", point.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for point := range stage.Points_reference {
		if _, ok := stage.Points[point]; !ok {
			points_deletedInstances = append(points_deletedInstances, point)
			deletedInstancesStmt += point.GongMarshallUnstaging(stage)
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
			newInstancesStmt += polygone.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := polygone.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := polygone.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", polygone.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for polygone := range stage.Polygones_reference {
		if _, ok := stage.Polygones[polygone]; !ok {
			polygones_deletedInstances = append(polygones_deletedInstances, polygone)
			deletedInstancesStmt += polygone.GongMarshallUnstaging(stage)
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
			newInstancesStmt += polyline.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := polyline.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := polyline.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", polyline.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for polyline := range stage.Polylines_reference {
		if _, ok := stage.Polylines[polyline]; !ok {
			polylines_deletedInstances = append(polylines_deletedInstances, polyline)
			deletedInstancesStmt += polyline.GongMarshallUnstaging(stage)
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
			newInstancesStmt += rect.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := rect.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := rect.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", rect.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for rect := range stage.Rects_reference {
		if _, ok := stage.Rects[rect]; !ok {
			rects_deletedInstances = append(rects_deletedInstances, rect)
			deletedInstancesStmt += rect.GongMarshallUnstaging(stage)
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
			newInstancesStmt += rectanchoredpath.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := rectanchoredpath.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := rectanchoredpath.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", rectanchoredpath.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for rectanchoredpath := range stage.RectAnchoredPaths_reference {
		if _, ok := stage.RectAnchoredPaths[rectanchoredpath]; !ok {
			rectanchoredpaths_deletedInstances = append(rectanchoredpaths_deletedInstances, rectanchoredpath)
			deletedInstancesStmt += rectanchoredpath.GongMarshallUnstaging(stage)
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
			newInstancesStmt += rectanchoredrect.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := rectanchoredrect.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := rectanchoredrect.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", rectanchoredrect.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for rectanchoredrect := range stage.RectAnchoredRects_reference {
		if _, ok := stage.RectAnchoredRects[rectanchoredrect]; !ok {
			rectanchoredrects_deletedInstances = append(rectanchoredrects_deletedInstances, rectanchoredrect)
			deletedInstancesStmt += rectanchoredrect.GongMarshallUnstaging(stage)
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
			newInstancesStmt += rectanchoredtext.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := rectanchoredtext.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := rectanchoredtext.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", rectanchoredtext.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for rectanchoredtext := range stage.RectAnchoredTexts_reference {
		if _, ok := stage.RectAnchoredTexts[rectanchoredtext]; !ok {
			rectanchoredtexts_deletedInstances = append(rectanchoredtexts_deletedInstances, rectanchoredtext)
			deletedInstancesStmt += rectanchoredtext.GongMarshallUnstaging(stage)
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
			newInstancesStmt += rectlinklink.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := rectlinklink.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := rectlinklink.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", rectlinklink.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for rectlinklink := range stage.RectLinkLinks_reference {
		if _, ok := stage.RectLinkLinks[rectlinklink]; !ok {
			rectlinklinks_deletedInstances = append(rectlinklinks_deletedInstances, rectlinklink)
			deletedInstancesStmt += rectlinklink.GongMarshallUnstaging(stage)
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
			newInstancesStmt += svg.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := svg.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := svg.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", svg.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for svg := range stage.SVGs_reference {
		if _, ok := stage.SVGs[svg]; !ok {
			svgs_deletedInstances = append(svgs_deletedInstances, svg)
			deletedInstancesStmt += svg.GongMarshallUnstaging(stage)
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
			newInstancesStmt += svgtext.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := svgtext.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := svgtext.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", svgtext.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for svgtext := range stage.SvgTexts_reference {
		if _, ok := stage.SvgTexts[svgtext]; !ok {
			svgtexts_deletedInstances = append(svgtexts_deletedInstances, svgtext)
			deletedInstancesStmt += svgtext.GongMarshallUnstaging(stage)
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
			newInstancesStmt += text.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := text.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := text.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", text.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for text := range stage.Texts_reference {
		if _, ok := stage.Texts[text]; !ok {
			texts_deletedInstances = append(texts_deletedInstances, text)
			deletedInstancesStmt += text.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(texts_newInstances)
	lenDeletedInstances += len(texts_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		notif := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		notif += fmt.Sprintf("\n\t// %s", time.Now().Format(time.RFC3339Nano))
		notif += "\n\tstage.Commit()"
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().AddNotification(
				time.Now(),
				notif,
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
	return stage.AnimateMap_Staged_Order[animate]
}

func (animate *Animate) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Animates_referenceOrder[animate]
}

func (circle *Circle) GongGetOrder(stage *Stage) uint {
	return stage.CircleMap_Staged_Order[circle]
}

func (circle *Circle) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Circles_referenceOrder[circle]
}

func (condition *Condition) GongGetOrder(stage *Stage) uint {
	return stage.ConditionMap_Staged_Order[condition]
}

func (condition *Condition) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Conditions_referenceOrder[condition]
}

func (controlpoint *ControlPoint) GongGetOrder(stage *Stage) uint {
	return stage.ControlPointMap_Staged_Order[controlpoint]
}

func (controlpoint *ControlPoint) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ControlPoints_referenceOrder[controlpoint]
}

func (ellipse *Ellipse) GongGetOrder(stage *Stage) uint {
	return stage.EllipseMap_Staged_Order[ellipse]
}

func (ellipse *Ellipse) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Ellipses_referenceOrder[ellipse]
}

func (layer *Layer) GongGetOrder(stage *Stage) uint {
	return stage.LayerMap_Staged_Order[layer]
}

func (layer *Layer) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Layers_referenceOrder[layer]
}

func (line *Line) GongGetOrder(stage *Stage) uint {
	return stage.LineMap_Staged_Order[line]
}

func (line *Line) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Lines_referenceOrder[line]
}

func (link *Link) GongGetOrder(stage *Stage) uint {
	return stage.LinkMap_Staged_Order[link]
}

func (link *Link) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Links_referenceOrder[link]
}

func (linkanchoredtext *LinkAnchoredText) GongGetOrder(stage *Stage) uint {
	return stage.LinkAnchoredTextMap_Staged_Order[linkanchoredtext]
}

func (linkanchoredtext *LinkAnchoredText) GongGetReferenceOrder(stage *Stage) uint {
	return stage.LinkAnchoredTexts_referenceOrder[linkanchoredtext]
}

func (path *Path) GongGetOrder(stage *Stage) uint {
	return stage.PathMap_Staged_Order[path]
}

func (path *Path) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Paths_referenceOrder[path]
}

func (point *Point) GongGetOrder(stage *Stage) uint {
	return stage.PointMap_Staged_Order[point]
}

func (point *Point) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Points_referenceOrder[point]
}

func (polygone *Polygone) GongGetOrder(stage *Stage) uint {
	return stage.PolygoneMap_Staged_Order[polygone]
}

func (polygone *Polygone) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Polygones_referenceOrder[polygone]
}

func (polyline *Polyline) GongGetOrder(stage *Stage) uint {
	return stage.PolylineMap_Staged_Order[polyline]
}

func (polyline *Polyline) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Polylines_referenceOrder[polyline]
}

func (rect *Rect) GongGetOrder(stage *Stage) uint {
	return stage.RectMap_Staged_Order[rect]
}

func (rect *Rect) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Rects_referenceOrder[rect]
}

func (rectanchoredpath *RectAnchoredPath) GongGetOrder(stage *Stage) uint {
	return stage.RectAnchoredPathMap_Staged_Order[rectanchoredpath]
}

func (rectanchoredpath *RectAnchoredPath) GongGetReferenceOrder(stage *Stage) uint {
	return stage.RectAnchoredPaths_referenceOrder[rectanchoredpath]
}

func (rectanchoredrect *RectAnchoredRect) GongGetOrder(stage *Stage) uint {
	return stage.RectAnchoredRectMap_Staged_Order[rectanchoredrect]
}

func (rectanchoredrect *RectAnchoredRect) GongGetReferenceOrder(stage *Stage) uint {
	return stage.RectAnchoredRects_referenceOrder[rectanchoredrect]
}

func (rectanchoredtext *RectAnchoredText) GongGetOrder(stage *Stage) uint {
	return stage.RectAnchoredTextMap_Staged_Order[rectanchoredtext]
}

func (rectanchoredtext *RectAnchoredText) GongGetReferenceOrder(stage *Stage) uint {
	return stage.RectAnchoredTexts_referenceOrder[rectanchoredtext]
}

func (rectlinklink *RectLinkLink) GongGetOrder(stage *Stage) uint {
	return stage.RectLinkLinkMap_Staged_Order[rectlinklink]
}

func (rectlinklink *RectLinkLink) GongGetReferenceOrder(stage *Stage) uint {
	return stage.RectLinkLinks_referenceOrder[rectlinklink]
}

func (svg *SVG) GongGetOrder(stage *Stage) uint {
	return stage.SVGMap_Staged_Order[svg]
}

func (svg *SVG) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SVGs_referenceOrder[svg]
}

func (svgtext *SvgText) GongGetOrder(stage *Stage) uint {
	return stage.SvgTextMap_Staged_Order[svgtext]
}

func (svgtext *SvgText) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SvgTexts_referenceOrder[svgtext]
}

func (text *Text) GongGetOrder(stage *Stage) uint {
	return stage.TextMap_Staged_Order[text]
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
