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

	// Compute reverse map for named struct FileToDownload
	// insertion point per field

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

	// Compute reverse map for named struct LinkAnchoredPath
	// insertion point per field

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

	// Compute reverse map for named struct RectAnchoredPath
	// insertion point per field

	// Compute reverse map for named struct RectAnchoredPngImage
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

	// end of insertion point per named struct
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

	for instance := range stage.FileToDownloads {
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

	for instance := range stage.LinkAnchoredPaths {
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

	for instance := range stage.RectAnchoredPngImages {
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
func (animate *Animate) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(animate).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(animate), uint64(stage.GetOrder(animate)))
	return
}

func (circle *Circle) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(circle).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(circle), uint64(stage.GetOrder(circle)))
	return
}

func (condition *Condition) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(condition).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(condition), uint64(stage.GetOrder(condition)))
	return
}

func (controlpoint *ControlPoint) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(controlpoint).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(controlpoint), uint64(stage.GetOrder(controlpoint)))
	return
}

func (ellipse *Ellipse) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(ellipse).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(ellipse), uint64(stage.GetOrder(ellipse)))
	return
}

func (filetodownload *FileToDownload) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(filetodownload).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(filetodownload), uint64(stage.GetOrder(filetodownload)))
	return
}

func (layer *Layer) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(layer).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(layer), uint64(stage.GetOrder(layer)))
	return
}

func (line *Line) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(line).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(line), uint64(stage.GetOrder(line)))
	return
}

func (link *Link) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(link).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(link), uint64(stage.GetOrder(link)))
	return
}

func (linkanchoredpath *LinkAnchoredPath) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(linkanchoredpath).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(linkanchoredpath), uint64(stage.GetOrder(linkanchoredpath)))
	return
}

func (linkanchoredtext *LinkAnchoredText) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(linkanchoredtext).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(linkanchoredtext), uint64(stage.GetOrder(linkanchoredtext)))
	return
}

func (path *Path) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(path).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(path), uint64(stage.GetOrder(path)))
	return
}

func (point *Point) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(point).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(point), uint64(stage.GetOrder(point)))
	return
}

func (polygone *Polygone) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(polygone).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(polygone), uint64(stage.GetOrder(polygone)))
	return
}

func (polyline *Polyline) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(polyline).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(polyline), uint64(stage.GetOrder(polyline)))
	return
}

func (rect *Rect) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rect).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(rect), uint64(stage.GetOrder(rect)))
	return
}

func (rectanchoredpath *RectAnchoredPath) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rectanchoredpath).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(rectanchoredpath), uint64(stage.GetOrder(rectanchoredpath)))
	return
}

func (rectanchoredpngimage *RectAnchoredPngImage) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rectanchoredpngimage).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(rectanchoredpngimage), uint64(stage.GetOrder(rectanchoredpngimage)))
	return
}

func (rectanchoredrect *RectAnchoredRect) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rectanchoredrect).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(rectanchoredrect), uint64(stage.GetOrder(rectanchoredrect)))
	return
}

func (rectanchoredtext *RectAnchoredText) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rectanchoredtext).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(rectanchoredtext), uint64(stage.GetOrder(rectanchoredtext)))
	return
}

func (rectlinklink *RectLinkLink) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rectlinklink).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(rectlinklink), uint64(stage.GetOrder(rectlinklink)))
	return
}

func (svg *SVG) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(svg).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(svg), uint64(stage.GetOrder(svg)))
	return
}

func (svgtext *SvgText) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(svgtext).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(svgtext), uint64(stage.GetOrder(svgtext)))
	return
}

func (text *Text) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(text).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(text), uint64(stage.GetOrder(text)))
	return
}


type GongstructDiffable[T any] interface {
	PointerToGongstruct
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
	stage.Animates_reference = make(map[*Animate]*Animate)
	stage.Animates_referenceOrder = make(map[*Animate]uint) // diff Unstage needs the reference order
	stage.Animates_instance = make(map[*Animate]*Animate)
	for instance := range stage.Animates {
		_copy := instance.GongCopy().(*Animate)
		stage.Animates_reference[instance] = _copy
		stage.Animates_instance[_copy] = instance
		stage.Animates_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Circles_reference = make(map[*Circle]*Circle)
	stage.Circles_referenceOrder = make(map[*Circle]uint) // diff Unstage needs the reference order
	stage.Circles_instance = make(map[*Circle]*Circle)
	for instance := range stage.Circles {
		_copy := instance.GongCopy().(*Circle)
		stage.Circles_reference[instance] = _copy
		stage.Circles_instance[_copy] = instance
		stage.Circles_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Conditions_reference = make(map[*Condition]*Condition)
	stage.Conditions_referenceOrder = make(map[*Condition]uint) // diff Unstage needs the reference order
	stage.Conditions_instance = make(map[*Condition]*Condition)
	for instance := range stage.Conditions {
		_copy := instance.GongCopy().(*Condition)
		stage.Conditions_reference[instance] = _copy
		stage.Conditions_instance[_copy] = instance
		stage.Conditions_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ControlPoints_reference = make(map[*ControlPoint]*ControlPoint)
	stage.ControlPoints_referenceOrder = make(map[*ControlPoint]uint) // diff Unstage needs the reference order
	stage.ControlPoints_instance = make(map[*ControlPoint]*ControlPoint)
	for instance := range stage.ControlPoints {
		_copy := instance.GongCopy().(*ControlPoint)
		stage.ControlPoints_reference[instance] = _copy
		stage.ControlPoints_instance[_copy] = instance
		stage.ControlPoints_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Ellipses_reference = make(map[*Ellipse]*Ellipse)
	stage.Ellipses_referenceOrder = make(map[*Ellipse]uint) // diff Unstage needs the reference order
	stage.Ellipses_instance = make(map[*Ellipse]*Ellipse)
	for instance := range stage.Ellipses {
		_copy := instance.GongCopy().(*Ellipse)
		stage.Ellipses_reference[instance] = _copy
		stage.Ellipses_instance[_copy] = instance
		stage.Ellipses_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FileToDownloads_reference = make(map[*FileToDownload]*FileToDownload)
	stage.FileToDownloads_referenceOrder = make(map[*FileToDownload]uint) // diff Unstage needs the reference order
	stage.FileToDownloads_instance = make(map[*FileToDownload]*FileToDownload)
	for instance := range stage.FileToDownloads {
		_copy := instance.GongCopy().(*FileToDownload)
		stage.FileToDownloads_reference[instance] = _copy
		stage.FileToDownloads_instance[_copy] = instance
		stage.FileToDownloads_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Layers_reference = make(map[*Layer]*Layer)
	stage.Layers_referenceOrder = make(map[*Layer]uint) // diff Unstage needs the reference order
	stage.Layers_instance = make(map[*Layer]*Layer)
	for instance := range stage.Layers {
		_copy := instance.GongCopy().(*Layer)
		stage.Layers_reference[instance] = _copy
		stage.Layers_instance[_copy] = instance
		stage.Layers_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Lines_reference = make(map[*Line]*Line)
	stage.Lines_referenceOrder = make(map[*Line]uint) // diff Unstage needs the reference order
	stage.Lines_instance = make(map[*Line]*Line)
	for instance := range stage.Lines {
		_copy := instance.GongCopy().(*Line)
		stage.Lines_reference[instance] = _copy
		stage.Lines_instance[_copy] = instance
		stage.Lines_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Links_reference = make(map[*Link]*Link)
	stage.Links_referenceOrder = make(map[*Link]uint) // diff Unstage needs the reference order
	stage.Links_instance = make(map[*Link]*Link)
	for instance := range stage.Links {
		_copy := instance.GongCopy().(*Link)
		stage.Links_reference[instance] = _copy
		stage.Links_instance[_copy] = instance
		stage.Links_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.LinkAnchoredPaths_reference = make(map[*LinkAnchoredPath]*LinkAnchoredPath)
	stage.LinkAnchoredPaths_referenceOrder = make(map[*LinkAnchoredPath]uint) // diff Unstage needs the reference order
	stage.LinkAnchoredPaths_instance = make(map[*LinkAnchoredPath]*LinkAnchoredPath)
	for instance := range stage.LinkAnchoredPaths {
		_copy := instance.GongCopy().(*LinkAnchoredPath)
		stage.LinkAnchoredPaths_reference[instance] = _copy
		stage.LinkAnchoredPaths_instance[_copy] = instance
		stage.LinkAnchoredPaths_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.LinkAnchoredTexts_reference = make(map[*LinkAnchoredText]*LinkAnchoredText)
	stage.LinkAnchoredTexts_referenceOrder = make(map[*LinkAnchoredText]uint) // diff Unstage needs the reference order
	stage.LinkAnchoredTexts_instance = make(map[*LinkAnchoredText]*LinkAnchoredText)
	for instance := range stage.LinkAnchoredTexts {
		_copy := instance.GongCopy().(*LinkAnchoredText)
		stage.LinkAnchoredTexts_reference[instance] = _copy
		stage.LinkAnchoredTexts_instance[_copy] = instance
		stage.LinkAnchoredTexts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Paths_reference = make(map[*Path]*Path)
	stage.Paths_referenceOrder = make(map[*Path]uint) // diff Unstage needs the reference order
	stage.Paths_instance = make(map[*Path]*Path)
	for instance := range stage.Paths {
		_copy := instance.GongCopy().(*Path)
		stage.Paths_reference[instance] = _copy
		stage.Paths_instance[_copy] = instance
		stage.Paths_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Points_reference = make(map[*Point]*Point)
	stage.Points_referenceOrder = make(map[*Point]uint) // diff Unstage needs the reference order
	stage.Points_instance = make(map[*Point]*Point)
	for instance := range stage.Points {
		_copy := instance.GongCopy().(*Point)
		stage.Points_reference[instance] = _copy
		stage.Points_instance[_copy] = instance
		stage.Points_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Polygones_reference = make(map[*Polygone]*Polygone)
	stage.Polygones_referenceOrder = make(map[*Polygone]uint) // diff Unstage needs the reference order
	stage.Polygones_instance = make(map[*Polygone]*Polygone)
	for instance := range stage.Polygones {
		_copy := instance.GongCopy().(*Polygone)
		stage.Polygones_reference[instance] = _copy
		stage.Polygones_instance[_copy] = instance
		stage.Polygones_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Polylines_reference = make(map[*Polyline]*Polyline)
	stage.Polylines_referenceOrder = make(map[*Polyline]uint) // diff Unstage needs the reference order
	stage.Polylines_instance = make(map[*Polyline]*Polyline)
	for instance := range stage.Polylines {
		_copy := instance.GongCopy().(*Polyline)
		stage.Polylines_reference[instance] = _copy
		stage.Polylines_instance[_copy] = instance
		stage.Polylines_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Rects_reference = make(map[*Rect]*Rect)
	stage.Rects_referenceOrder = make(map[*Rect]uint) // diff Unstage needs the reference order
	stage.Rects_instance = make(map[*Rect]*Rect)
	for instance := range stage.Rects {
		_copy := instance.GongCopy().(*Rect)
		stage.Rects_reference[instance] = _copy
		stage.Rects_instance[_copy] = instance
		stage.Rects_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.RectAnchoredPaths_reference = make(map[*RectAnchoredPath]*RectAnchoredPath)
	stage.RectAnchoredPaths_referenceOrder = make(map[*RectAnchoredPath]uint) // diff Unstage needs the reference order
	stage.RectAnchoredPaths_instance = make(map[*RectAnchoredPath]*RectAnchoredPath)
	for instance := range stage.RectAnchoredPaths {
		_copy := instance.GongCopy().(*RectAnchoredPath)
		stage.RectAnchoredPaths_reference[instance] = _copy
		stage.RectAnchoredPaths_instance[_copy] = instance
		stage.RectAnchoredPaths_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.RectAnchoredPngImages_reference = make(map[*RectAnchoredPngImage]*RectAnchoredPngImage)
	stage.RectAnchoredPngImages_referenceOrder = make(map[*RectAnchoredPngImage]uint) // diff Unstage needs the reference order
	stage.RectAnchoredPngImages_instance = make(map[*RectAnchoredPngImage]*RectAnchoredPngImage)
	for instance := range stage.RectAnchoredPngImages {
		_copy := instance.GongCopy().(*RectAnchoredPngImage)
		stage.RectAnchoredPngImages_reference[instance] = _copy
		stage.RectAnchoredPngImages_instance[_copy] = instance
		stage.RectAnchoredPngImages_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.RectAnchoredRects_reference = make(map[*RectAnchoredRect]*RectAnchoredRect)
	stage.RectAnchoredRects_referenceOrder = make(map[*RectAnchoredRect]uint) // diff Unstage needs the reference order
	stage.RectAnchoredRects_instance = make(map[*RectAnchoredRect]*RectAnchoredRect)
	for instance := range stage.RectAnchoredRects {
		_copy := instance.GongCopy().(*RectAnchoredRect)
		stage.RectAnchoredRects_reference[instance] = _copy
		stage.RectAnchoredRects_instance[_copy] = instance
		stage.RectAnchoredRects_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.RectAnchoredTexts_reference = make(map[*RectAnchoredText]*RectAnchoredText)
	stage.RectAnchoredTexts_referenceOrder = make(map[*RectAnchoredText]uint) // diff Unstage needs the reference order
	stage.RectAnchoredTexts_instance = make(map[*RectAnchoredText]*RectAnchoredText)
	for instance := range stage.RectAnchoredTexts {
		_copy := instance.GongCopy().(*RectAnchoredText)
		stage.RectAnchoredTexts_reference[instance] = _copy
		stage.RectAnchoredTexts_instance[_copy] = instance
		stage.RectAnchoredTexts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.RectLinkLinks_reference = make(map[*RectLinkLink]*RectLinkLink)
	stage.RectLinkLinks_referenceOrder = make(map[*RectLinkLink]uint) // diff Unstage needs the reference order
	stage.RectLinkLinks_instance = make(map[*RectLinkLink]*RectLinkLink)
	for instance := range stage.RectLinkLinks {
		_copy := instance.GongCopy().(*RectLinkLink)
		stage.RectLinkLinks_reference[instance] = _copy
		stage.RectLinkLinks_instance[_copy] = instance
		stage.RectLinkLinks_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SVGs_reference = make(map[*SVG]*SVG)
	stage.SVGs_referenceOrder = make(map[*SVG]uint) // diff Unstage needs the reference order
	stage.SVGs_instance = make(map[*SVG]*SVG)
	for instance := range stage.SVGs {
		_copy := instance.GongCopy().(*SVG)
		stage.SVGs_reference[instance] = _copy
		stage.SVGs_instance[_copy] = instance
		stage.SVGs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SvgTexts_reference = make(map[*SvgText]*SvgText)
	stage.SvgTexts_referenceOrder = make(map[*SvgText]uint) // diff Unstage needs the reference order
	stage.SvgTexts_instance = make(map[*SvgText]*SvgText)
	for instance := range stage.SvgTexts {
		_copy := instance.GongCopy().(*SvgText)
		stage.SvgTexts_reference[instance] = _copy
		stage.SvgTexts_instance[_copy] = instance
		stage.SvgTexts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Texts_reference = make(map[*Text]*Text)
	stage.Texts_referenceOrder = make(map[*Text]uint) // diff Unstage needs the reference order
	stage.Texts_instance = make(map[*Text]*Text)
	for instance := range stage.Texts {
		_copy := instance.GongCopy().(*Text)
		stage.Texts_reference[instance] = _copy
		stage.Texts_instance[_copy] = instance
		stage.Texts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.Animates {
		reference := stage.Animates_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Circles {
		reference := stage.Circles_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Conditions {
		reference := stage.Conditions_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ControlPoints {
		reference := stage.ControlPoints_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Ellipses {
		reference := stage.Ellipses_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FileToDownloads {
		reference := stage.FileToDownloads_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Layers {
		reference := stage.Layers_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Lines {
		reference := stage.Lines_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Links {
		reference := stage.Links_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.LinkAnchoredPaths {
		reference := stage.LinkAnchoredPaths_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.LinkAnchoredTexts {
		reference := stage.LinkAnchoredTexts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Paths {
		reference := stage.Paths_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Points {
		reference := stage.Points_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Polygones {
		reference := stage.Polygones_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Polylines {
		reference := stage.Polylines_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Rects {
		reference := stage.Rects_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.RectAnchoredPaths {
		reference := stage.RectAnchoredPaths_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.RectAnchoredPngImages {
		reference := stage.RectAnchoredPngImages_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.RectAnchoredRects {
		reference := stage.RectAnchoredRects_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.RectAnchoredTexts {
		reference := stage.RectAnchoredTexts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.RectLinkLinks {
		reference := stage.RectLinkLinks_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SVGs {
		reference := stage.SVGs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SvgTexts {
		reference := stage.SvgTexts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Texts {
		reference := stage.Texts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

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
	if order, ok := stage.Animate_stagedOrder[animate]; ok {
		return order
	}
	if order, ok := stage.Animates_referenceOrder[animate]; ok {
		return order
	} else {
		log.Printf("instance %p of type Animate was not staged and does not have a reference order", animate)
		return 0
	}
}

func (circle *Circle) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Circle_stagedOrder[circle]; ok {
		return order
	}
	if order, ok := stage.Circles_referenceOrder[circle]; ok {
		return order
	} else {
		log.Printf("instance %p of type Circle was not staged and does not have a reference order", circle)
		return 0
	}
}

func (condition *Condition) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Condition_stagedOrder[condition]; ok {
		return order
	}
	if order, ok := stage.Conditions_referenceOrder[condition]; ok {
		return order
	} else {
		log.Printf("instance %p of type Condition was not staged and does not have a reference order", condition)
		return 0
	}
}

func (controlpoint *ControlPoint) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ControlPoint_stagedOrder[controlpoint]; ok {
		return order
	}
	if order, ok := stage.ControlPoints_referenceOrder[controlpoint]; ok {
		return order
	} else {
		log.Printf("instance %p of type ControlPoint was not staged and does not have a reference order", controlpoint)
		return 0
	}
}

func (ellipse *Ellipse) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Ellipse_stagedOrder[ellipse]; ok {
		return order
	}
	if order, ok := stage.Ellipses_referenceOrder[ellipse]; ok {
		return order
	} else {
		log.Printf("instance %p of type Ellipse was not staged and does not have a reference order", ellipse)
		return 0
	}
}

func (filetodownload *FileToDownload) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FileToDownload_stagedOrder[filetodownload]; ok {
		return order
	}
	if order, ok := stage.FileToDownloads_referenceOrder[filetodownload]; ok {
		return order
	} else {
		log.Printf("instance %p of type FileToDownload was not staged and does not have a reference order", filetodownload)
		return 0
	}
}

func (layer *Layer) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Layer_stagedOrder[layer]; ok {
		return order
	}
	if order, ok := stage.Layers_referenceOrder[layer]; ok {
		return order
	} else {
		log.Printf("instance %p of type Layer was not staged and does not have a reference order", layer)
		return 0
	}
}

func (line *Line) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Line_stagedOrder[line]; ok {
		return order
	}
	if order, ok := stage.Lines_referenceOrder[line]; ok {
		return order
	} else {
		log.Printf("instance %p of type Line was not staged and does not have a reference order", line)
		return 0
	}
}

func (link *Link) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Link_stagedOrder[link]; ok {
		return order
	}
	if order, ok := stage.Links_referenceOrder[link]; ok {
		return order
	} else {
		log.Printf("instance %p of type Link was not staged and does not have a reference order", link)
		return 0
	}
}

func (linkanchoredpath *LinkAnchoredPath) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.LinkAnchoredPath_stagedOrder[linkanchoredpath]; ok {
		return order
	}
	if order, ok := stage.LinkAnchoredPaths_referenceOrder[linkanchoredpath]; ok {
		return order
	} else {
		log.Printf("instance %p of type LinkAnchoredPath was not staged and does not have a reference order", linkanchoredpath)
		return 0
	}
}

func (linkanchoredtext *LinkAnchoredText) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.LinkAnchoredText_stagedOrder[linkanchoredtext]; ok {
		return order
	}
	if order, ok := stage.LinkAnchoredTexts_referenceOrder[linkanchoredtext]; ok {
		return order
	} else {
		log.Printf("instance %p of type LinkAnchoredText was not staged and does not have a reference order", linkanchoredtext)
		return 0
	}
}

func (path *Path) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Path_stagedOrder[path]; ok {
		return order
	}
	if order, ok := stage.Paths_referenceOrder[path]; ok {
		return order
	} else {
		log.Printf("instance %p of type Path was not staged and does not have a reference order", path)
		return 0
	}
}

func (point *Point) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Point_stagedOrder[point]; ok {
		return order
	}
	if order, ok := stage.Points_referenceOrder[point]; ok {
		return order
	} else {
		log.Printf("instance %p of type Point was not staged and does not have a reference order", point)
		return 0
	}
}

func (polygone *Polygone) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Polygone_stagedOrder[polygone]; ok {
		return order
	}
	if order, ok := stage.Polygones_referenceOrder[polygone]; ok {
		return order
	} else {
		log.Printf("instance %p of type Polygone was not staged and does not have a reference order", polygone)
		return 0
	}
}

func (polyline *Polyline) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Polyline_stagedOrder[polyline]; ok {
		return order
	}
	if order, ok := stage.Polylines_referenceOrder[polyline]; ok {
		return order
	} else {
		log.Printf("instance %p of type Polyline was not staged and does not have a reference order", polyline)
		return 0
	}
}

func (rect *Rect) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Rect_stagedOrder[rect]; ok {
		return order
	}
	if order, ok := stage.Rects_referenceOrder[rect]; ok {
		return order
	} else {
		log.Printf("instance %p of type Rect was not staged and does not have a reference order", rect)
		return 0
	}
}

func (rectanchoredpath *RectAnchoredPath) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RectAnchoredPath_stagedOrder[rectanchoredpath]; ok {
		return order
	}
	if order, ok := stage.RectAnchoredPaths_referenceOrder[rectanchoredpath]; ok {
		return order
	} else {
		log.Printf("instance %p of type RectAnchoredPath was not staged and does not have a reference order", rectanchoredpath)
		return 0
	}
}

func (rectanchoredpngimage *RectAnchoredPngImage) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RectAnchoredPngImage_stagedOrder[rectanchoredpngimage]; ok {
		return order
	}
	if order, ok := stage.RectAnchoredPngImages_referenceOrder[rectanchoredpngimage]; ok {
		return order
	} else {
		log.Printf("instance %p of type RectAnchoredPngImage was not staged and does not have a reference order", rectanchoredpngimage)
		return 0
	}
}

func (rectanchoredrect *RectAnchoredRect) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RectAnchoredRect_stagedOrder[rectanchoredrect]; ok {
		return order
	}
	if order, ok := stage.RectAnchoredRects_referenceOrder[rectanchoredrect]; ok {
		return order
	} else {
		log.Printf("instance %p of type RectAnchoredRect was not staged and does not have a reference order", rectanchoredrect)
		return 0
	}
}

func (rectanchoredtext *RectAnchoredText) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RectAnchoredText_stagedOrder[rectanchoredtext]; ok {
		return order
	}
	if order, ok := stage.RectAnchoredTexts_referenceOrder[rectanchoredtext]; ok {
		return order
	} else {
		log.Printf("instance %p of type RectAnchoredText was not staged and does not have a reference order", rectanchoredtext)
		return 0
	}
}

func (rectlinklink *RectLinkLink) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RectLinkLink_stagedOrder[rectlinklink]; ok {
		return order
	}
	if order, ok := stage.RectLinkLinks_referenceOrder[rectlinklink]; ok {
		return order
	} else {
		log.Printf("instance %p of type RectLinkLink was not staged and does not have a reference order", rectlinklink)
		return 0
	}
}

func (svg *SVG) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SVG_stagedOrder[svg]; ok {
		return order
	}
	if order, ok := stage.SVGs_referenceOrder[svg]; ok {
		return order
	} else {
		log.Printf("instance %p of type SVG was not staged and does not have a reference order", svg)
		return 0
	}
}

func (svgtext *SvgText) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SvgText_stagedOrder[svgtext]; ok {
		return order
	}
	if order, ok := stage.SvgTexts_referenceOrder[svgtext]; ok {
		return order
	} else {
		log.Printf("instance %p of type SvgText was not staged and does not have a reference order", svgtext)
		return 0
	}
}

func (text *Text) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Text_stagedOrder[text]; ok {
		return order
	}
	if order, ok := stage.Texts_referenceOrder[text]; ok {
		return order
	} else {
		log.Printf("instance %p of type Text was not staged and does not have a reference order", text)
		return 0
	}
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
	return fmt.Sprintf("__%s__%08d_", animate.GongGetGongstructName(), animate.GongGetOrder(stage))
}

func (circle *Circle) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", circle.GongGetGongstructName(), circle.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (circle *Circle) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", circle.GongGetGongstructName(), circle.GongGetOrder(stage))
}

func (condition *Condition) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", condition.GongGetGongstructName(), condition.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (condition *Condition) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", condition.GongGetGongstructName(), condition.GongGetOrder(stage))
}

func (controlpoint *ControlPoint) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", controlpoint.GongGetGongstructName(), controlpoint.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (controlpoint *ControlPoint) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", controlpoint.GongGetGongstructName(), controlpoint.GongGetOrder(stage))
}

func (ellipse *Ellipse) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", ellipse.GongGetGongstructName(), ellipse.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ellipse *Ellipse) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", ellipse.GongGetGongstructName(), ellipse.GongGetOrder(stage))
}

func (filetodownload *FileToDownload) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", filetodownload.GongGetGongstructName(), filetodownload.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (filetodownload *FileToDownload) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", filetodownload.GongGetGongstructName(), filetodownload.GongGetOrder(stage))
}

func (layer *Layer) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", layer.GongGetGongstructName(), layer.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (layer *Layer) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", layer.GongGetGongstructName(), layer.GongGetOrder(stage))
}

func (line *Line) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", line.GongGetGongstructName(), line.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (line *Line) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", line.GongGetGongstructName(), line.GongGetOrder(stage))
}

func (link *Link) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", link.GongGetGongstructName(), link.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (link *Link) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", link.GongGetGongstructName(), link.GongGetOrder(stage))
}

func (linkanchoredpath *LinkAnchoredPath) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", linkanchoredpath.GongGetGongstructName(), linkanchoredpath.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (linkanchoredpath *LinkAnchoredPath) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", linkanchoredpath.GongGetGongstructName(), linkanchoredpath.GongGetOrder(stage))
}

func (linkanchoredtext *LinkAnchoredText) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", linkanchoredtext.GongGetGongstructName(), linkanchoredtext.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (linkanchoredtext *LinkAnchoredText) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", linkanchoredtext.GongGetGongstructName(), linkanchoredtext.GongGetOrder(stage))
}

func (path *Path) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", path.GongGetGongstructName(), path.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (path *Path) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", path.GongGetGongstructName(), path.GongGetOrder(stage))
}

func (point *Point) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", point.GongGetGongstructName(), point.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (point *Point) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", point.GongGetGongstructName(), point.GongGetOrder(stage))
}

func (polygone *Polygone) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", polygone.GongGetGongstructName(), polygone.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (polygone *Polygone) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", polygone.GongGetGongstructName(), polygone.GongGetOrder(stage))
}

func (polyline *Polyline) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", polyline.GongGetGongstructName(), polyline.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (polyline *Polyline) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", polyline.GongGetGongstructName(), polyline.GongGetOrder(stage))
}

func (rect *Rect) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rect.GongGetGongstructName(), rect.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rect *Rect) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rect.GongGetGongstructName(), rect.GongGetOrder(stage))
}

func (rectanchoredpath *RectAnchoredPath) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rectanchoredpath.GongGetGongstructName(), rectanchoredpath.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rectanchoredpath *RectAnchoredPath) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rectanchoredpath.GongGetGongstructName(), rectanchoredpath.GongGetOrder(stage))
}

func (rectanchoredpngimage *RectAnchoredPngImage) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rectanchoredpngimage.GongGetGongstructName(), rectanchoredpngimage.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rectanchoredpngimage *RectAnchoredPngImage) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rectanchoredpngimage.GongGetGongstructName(), rectanchoredpngimage.GongGetOrder(stage))
}

func (rectanchoredrect *RectAnchoredRect) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rectanchoredrect.GongGetGongstructName(), rectanchoredrect.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rectanchoredrect *RectAnchoredRect) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rectanchoredrect.GongGetGongstructName(), rectanchoredrect.GongGetOrder(stage))
}

func (rectanchoredtext *RectAnchoredText) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rectanchoredtext.GongGetGongstructName(), rectanchoredtext.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rectanchoredtext *RectAnchoredText) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rectanchoredtext.GongGetGongstructName(), rectanchoredtext.GongGetOrder(stage))
}

func (rectlinklink *RectLinkLink) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rectlinklink.GongGetGongstructName(), rectlinklink.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rectlinklink *RectLinkLink) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rectlinklink.GongGetGongstructName(), rectlinklink.GongGetOrder(stage))
}

func (svg *SVG) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svg.GongGetGongstructName(), svg.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (svg *SVG) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svg.GongGetGongstructName(), svg.GongGetOrder(stage))
}

func (svgtext *SvgText) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svgtext.GongGetGongstructName(), svgtext.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (svgtext *SvgText) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svgtext.GongGetGongstructName(), svgtext.GongGetOrder(stage))
}

func (text *Text) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", text.GongGetGongstructName(), text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (text *Text) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", text.GongGetGongstructName(), text.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (animate *Animate) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", animate.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Animate")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(animate.Name))
	return
}

func (circle *Circle) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", circle.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Circle")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(circle.Name))
	return
}

func (condition *Condition) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", condition.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Condition")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(condition.Name))
	return
}

func (controlpoint *ControlPoint) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", controlpoint.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ControlPoint")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(controlpoint.Name))
	return
}

func (ellipse *Ellipse) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Ellipse")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(ellipse.Name))
	return
}

func (filetodownload *FileToDownload) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", filetodownload.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FileToDownload")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(filetodownload.Name))
	return
}

func (layer *Layer) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", layer.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Layer")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(layer.Name))
	return
}

func (line *Line) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", line.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Line")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(line.Name))
	return
}

func (link *Link) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", link.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Link")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(link.Name))
	return
}

func (linkanchoredpath *LinkAnchoredPath) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", linkanchoredpath.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LinkAnchoredPath")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(linkanchoredpath.Name))
	return
}

func (linkanchoredtext *LinkAnchoredText) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LinkAnchoredText")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(linkanchoredtext.Name))
	return
}

func (path *Path) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", path.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Path")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(path.Name))
	return
}

func (point *Point) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", point.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Point")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(point.Name))
	return
}

func (polygone *Polygone) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", polygone.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Polygone")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(polygone.Name))
	return
}

func (polyline *Polyline) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", polyline.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Polyline")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(polyline.Name))
	return
}

func (rect *Rect) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rect.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Rect")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(rect.Name))
	return
}

func (rectanchoredpath *RectAnchoredPath) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RectAnchoredPath")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(rectanchoredpath.Name))
	return
}

func (rectanchoredpngimage *RectAnchoredPngImage) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rectanchoredpngimage.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RectAnchoredPngImage")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(rectanchoredpngimage.Name))
	return
}

func (rectanchoredrect *RectAnchoredRect) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RectAnchoredRect")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(rectanchoredrect.Name))
	return
}

func (rectanchoredtext *RectAnchoredText) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RectAnchoredText")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(rectanchoredtext.Name))
	return
}

func (rectlinklink *RectLinkLink) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RectLinkLink")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(rectlinklink.Name))
	return
}

func (svg *SVG) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", svg.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SVG")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(svg.Name))
	return
}

func (svgtext *SvgText) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", svgtext.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SvgText")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(svgtext.Name))
	return
}

func (text *Text) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", text.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Text")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(text.Name))
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

func (filetodownload *FileToDownload) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", filetodownload.GongGetReferenceIdentifier(stage))
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

func (linkanchoredpath *LinkAnchoredPath) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", linkanchoredpath.GongGetReferenceIdentifier(stage))
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

func (rectanchoredpngimage *RectAnchoredPngImage) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rectanchoredpngimage.GongGetReferenceIdentifier(stage))
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

// end of template
