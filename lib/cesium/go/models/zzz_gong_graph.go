// generated code - do not edit
package models

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (coordinate *Coordinate) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Coordinates[coordinate]
	return ok
}

func (layer *Layer) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Layers[layer]
	return ok
}

func (mapview *MapView) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MapViews[mapview]
	return ok
}

func (marker *Marker) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Markers[marker]
	return ok
}

func (polygon *Polygon) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Polygons[polygon]
	return ok
}

func (polyline *Polyline) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Polylines[polyline]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (coordinate *Coordinate) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coordinate) {
		return
	}

	coordinate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (layer *Layer) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(layer) {
		return
	}

	layer.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _polygon := range layer.Polygons {
		stage.StageBranch(_polygon)
	}
	for _, _polyline := range layer.Polylines {
		stage.StageBranch(_polyline)
	}
	for _, _marker := range layer.Markers {
		stage.StageBranch(_marker)
	}

}

func (mapview *MapView) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(mapview) {
		return
	}

	mapview.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _layer := range mapview.Layers {
		stage.StageBranch(_layer)
	}

}

func (marker *Marker) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(marker) {
		return
	}

	marker.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (polygon *Polygon) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(polygon) {
		return
	}

	polygon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _coordinate := range polygon.Coordinates {
		stage.StageBranch(_coordinate)
	}

}

func (polyline *Polyline) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(polyline) {
		return
	}

	polyline.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _coordinate := range polyline.Coordinates {
		stage.StageBranch(_coordinate)
	}

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Coordinate:
		toT := GongCopyBranchCoordinate(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Layer:
		toT := GongCopyBranchLayer(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MapView:
		toT := GongCopyBranchMapView(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Marker:
		toT := GongCopyBranchMarker(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Polygon:
		toT := GongCopyBranchPolygon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Polyline:
		toT := GongCopyBranchPolyline(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchCoordinate(mapOrigCopy map[any]any, coordinateFrom *Coordinate) (coordinateTo *Coordinate) {
	var alreadyCopied bool
	coordinateTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coordinateFrom)
	if alreadyCopied {
		return
	}
	coordinateFrom.GongCopyBasicFields(coordinateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLayer(mapOrigCopy map[any]any, layerFrom *Layer) (layerTo *Layer) {
	var alreadyCopied bool
	layerTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, layerFrom)
	if alreadyCopied {
		return
	}
	layerFrom.GongCopyBasicFields(layerTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _polygon := range layerFrom.Polygons {
		layerTo.Polygons = append(layerTo.Polygons, GongCopyBranchPolygon(mapOrigCopy, _polygon))
	}
	for _, _polyline := range layerFrom.Polylines {
		layerTo.Polylines = append(layerTo.Polylines, GongCopyBranchPolyline(mapOrigCopy, _polyline))
	}
	for _, _marker := range layerFrom.Markers {
		layerTo.Markers = append(layerTo.Markers, GongCopyBranchMarker(mapOrigCopy, _marker))
	}

	return
}

func GongCopyBranchMapView(mapOrigCopy map[any]any, mapviewFrom *MapView) (mapviewTo *MapView) {
	var alreadyCopied bool
	mapviewTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, mapviewFrom)
	if alreadyCopied {
		return
	}
	mapviewFrom.GongCopyBasicFields(mapviewTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _layer := range mapviewFrom.Layers {
		mapviewTo.Layers = append(mapviewTo.Layers, GongCopyBranchLayer(mapOrigCopy, _layer))
	}

	return
}

func GongCopyBranchMarker(mapOrigCopy map[any]any, markerFrom *Marker) (markerTo *Marker) {
	var alreadyCopied bool
	markerTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, markerFrom)
	if alreadyCopied {
		return
	}
	markerFrom.GongCopyBasicFields(markerTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPolygon(mapOrigCopy map[any]any, polygonFrom *Polygon) (polygonTo *Polygon) {
	var alreadyCopied bool
	polygonTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, polygonFrom)
	if alreadyCopied {
		return
	}
	polygonFrom.GongCopyBasicFields(polygonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _coordinate := range polygonFrom.Coordinates {
		polygonTo.Coordinates = append(polygonTo.Coordinates, GongCopyBranchCoordinate(mapOrigCopy, _coordinate))
	}

	return
}

func GongCopyBranchPolyline(mapOrigCopy map[any]any, polylineFrom *Polyline) (polylineTo *Polyline) {
	var alreadyCopied bool
	polylineTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, polylineFrom)
	if alreadyCopied {
		return
	}
	polylineFrom.GongCopyBasicFields(polylineTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _coordinate := range polylineFrom.Coordinates {
		polylineTo.Coordinates = append(polylineTo.Coordinates, GongCopyBranchCoordinate(mapOrigCopy, _coordinate))
	}

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// insertion point for unstage branch per struct
func (coordinate *Coordinate) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coordinate) {
		return
	}

	coordinate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (layer *Layer) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(layer) {
		return
	}

	layer.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _polygon := range layer.Polygons {
		stage.UnstageBranch(_polygon)
	}
	for _, _polyline := range layer.Polylines {
		stage.UnstageBranch(_polyline)
	}
	for _, _marker := range layer.Markers {
		stage.UnstageBranch(_marker)
	}

}

func (mapview *MapView) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(mapview) {
		return
	}

	mapview.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _layer := range mapview.Layers {
		stage.UnstageBranch(_layer)
	}

}

func (marker *Marker) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(marker) {
		return
	}

	marker.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (polygon *Polygon) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(polygon) {
		return
	}

	polygon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _coordinate := range polygon.Coordinates {
		stage.UnstageBranch(_coordinate)
	}

}

func (polyline *Polyline) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(polyline) {
		return
	}

	polyline.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _coordinate := range polyline.Coordinates {
		stage.UnstageBranch(_coordinate)
	}

}

// insertion point for pointer reconstruction from references
func (reference *Coordinate) GongReconstructPointersFromReferences(stage *Stage, instance *Coordinate) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Layer) GongReconstructPointersFromReferences(stage *Stage, instance *Layer) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Polygons, stage.Polygons_reference, instance.Polygons)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Polylines, stage.Polylines_reference, instance.Polylines)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Markers, stage.Markers_reference, instance.Markers)
}

func (reference *MapView) GongReconstructPointersFromReferences(stage *Stage, instance *MapView) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Layers, stage.Layers_reference, instance.Layers)
}

func (reference *Marker) GongReconstructPointersFromReferences(stage *Stage, instance *Marker) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Polygon) GongReconstructPointersFromReferences(stage *Stage, instance *Polygon) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Coordinates, stage.Coordinates_reference, instance.Coordinates)
}

func (reference *Polyline) GongReconstructPointersFromReferences(stage *Stage, instance *Polyline) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Coordinates, stage.Coordinates_reference, instance.Coordinates)
}

// insertion point for pointer reconstruction from instances
func (reference *Coordinate) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Layer) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Polygons, stage.Polygons_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Polylines, stage.Polylines_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Markers, stage.Markers_instance)
}

func (reference *MapView) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Layers, stage.Layers_instance)
}

func (reference *Marker) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Polygon) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Coordinates, stage.Coordinates_instance)
}

func (reference *Polyline) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Coordinates, stage.Coordinates_instance)
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coordinate *Coordinate) GongDiff(stage *Stage, coordinateOther *Coordinate) (diffs []string) {
	// insertion point for field diffs
	if coordinate.Name != coordinateOther.Name {
		diffs = append(diffs, coordinate.GongMarshallField(stage, "Name"))
	}
	if coordinate.Longitude != coordinateOther.Longitude {
		diffs = append(diffs, coordinate.GongMarshallField(stage, "Longitude"))
	}
	if coordinate.Latitude != coordinateOther.Latitude {
		diffs = append(diffs, coordinate.GongMarshallField(stage, "Latitude"))
	}
	if coordinate.Altitude != coordinateOther.Altitude {
		diffs = append(diffs, coordinate.GongMarshallField(stage, "Altitude"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (layer *Layer) GongDiff(stage *Stage, layerOther *Layer) (diffs []string) {
	// insertion point for field diffs
	if layer.Name != layerOther.Name {
		diffs = append(diffs, layer.GongMarshallField(stage, "Name"))
	}
	if layer.IsVisible != layerOther.IsVisible {
		diffs = append(diffs, layer.GongMarshallField(stage, "IsVisible"))
	}
	if ops := __gong__diffSliceOfPointers(stage, layer, "Polygons", layerOther.Polygons, layer.Polygons); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, layer, "Polylines", layerOther.Polylines, layer.Polylines); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, layer, "Markers", layerOther.Markers, layer.Markers); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (mapview *MapView) GongDiff(stage *Stage, mapviewOther *MapView) (diffs []string) {
	// insertion point for field diffs
	if mapview.Name != mapviewOther.Name {
		diffs = append(diffs, mapview.GongMarshallField(stage, "Name"))
	}
	if mapview.StyleURL != mapviewOther.StyleURL {
		diffs = append(diffs, mapview.GongMarshallField(stage, "StyleURL"))
	}
	if mapview.Longitude != mapviewOther.Longitude {
		diffs = append(diffs, mapview.GongMarshallField(stage, "Longitude"))
	}
	if mapview.Latitude != mapviewOther.Latitude {
		diffs = append(diffs, mapview.GongMarshallField(stage, "Latitude"))
	}
	if mapview.Zoom != mapviewOther.Zoom {
		diffs = append(diffs, mapview.GongMarshallField(stage, "Zoom"))
	}
	if mapview.Pitch != mapviewOther.Pitch {
		diffs = append(diffs, mapview.GongMarshallField(stage, "Pitch"))
	}
	if mapview.Bearing != mapviewOther.Bearing {
		diffs = append(diffs, mapview.GongMarshallField(stage, "Bearing"))
	}
	if ops := __gong__diffSliceOfPointers(stage, mapview, "Layers", mapviewOther.Layers, mapview.Layers); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (marker *Marker) GongDiff(stage *Stage, markerOther *Marker) (diffs []string) {
	// insertion point for field diffs
	if marker.Name != markerOther.Name {
		diffs = append(diffs, marker.GongMarshallField(stage, "Name"))
	}
	if marker.Label != markerOther.Label {
		diffs = append(diffs, marker.GongMarshallField(stage, "Label"))
	}
	if marker.PopupHTML != markerOther.PopupHTML {
		diffs = append(diffs, marker.GongMarshallField(stage, "PopupHTML"))
	}
	if marker.Longitude != markerOther.Longitude {
		diffs = append(diffs, marker.GongMarshallField(stage, "Longitude"))
	}
	if marker.Latitude != markerOther.Latitude {
		diffs = append(diffs, marker.GongMarshallField(stage, "Latitude"))
	}
	if marker.Altitude != markerOther.Altitude {
		diffs = append(diffs, marker.GongMarshallField(stage, "Altitude"))
	}
	if marker.SVG != markerOther.SVG {
		diffs = append(diffs, marker.GongMarshallField(stage, "SVG"))
	}
	if marker.Color != markerOther.Color {
		diffs = append(diffs, marker.GongMarshallField(stage, "Color"))
	}
	if marker.Scale != markerOther.Scale {
		diffs = append(diffs, marker.GongMarshallField(stage, "Scale"))
	}
	if marker.Rotation != markerOther.Rotation {
		diffs = append(diffs, marker.GongMarshallField(stage, "Rotation"))
	}
	if marker.RotationAlignment != markerOther.RotationAlignment {
		diffs = append(diffs, marker.GongMarshallField(stage, "RotationAlignment"))
	}
	if marker.PitchAlignment != markerOther.PitchAlignment {
		diffs = append(diffs, marker.GongMarshallField(stage, "PitchAlignment"))
	}
	if marker.Anchor != markerOther.Anchor {
		diffs = append(diffs, marker.GongMarshallField(stage, "Anchor"))
	}
	if marker.IsDraggable != markerOther.IsDraggable {
		diffs = append(diffs, marker.GongMarshallField(stage, "IsDraggable"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (polygon *Polygon) GongDiff(stage *Stage, polygonOther *Polygon) (diffs []string) {
	// insertion point for field diffs
	if polygon.Name != polygonOther.Name {
		diffs = append(diffs, polygon.GongMarshallField(stage, "Name"))
	}
	if polygon.Label != polygonOther.Label {
		diffs = append(diffs, polygon.GongMarshallField(stage, "Label"))
	}
	if polygon.PopupHTML != polygonOther.PopupHTML {
		diffs = append(diffs, polygon.GongMarshallField(stage, "PopupHTML"))
	}
	if ops := __gong__diffSliceOfPointers(stage, polygon, "Coordinates", polygonOther.Coordinates, polygon.Coordinates); ops != "" {
		diffs = append(diffs, ops)
	}
	if polygon.FillColor != polygonOther.FillColor {
		diffs = append(diffs, polygon.GongMarshallField(stage, "FillColor"))
	}
	if polygon.FillOpacity != polygonOther.FillOpacity {
		diffs = append(diffs, polygon.GongMarshallField(stage, "FillOpacity"))
	}
	if polygon.StrokeColor != polygonOther.StrokeColor {
		diffs = append(diffs, polygon.GongMarshallField(stage, "StrokeColor"))
	}
	if polygon.StrokeOpacity != polygonOther.StrokeOpacity {
		diffs = append(diffs, polygon.GongMarshallField(stage, "StrokeOpacity"))
	}
	if polygon.StrokeWidth != polygonOther.StrokeWidth {
		diffs = append(diffs, polygon.GongMarshallField(stage, "StrokeWidth"))
	}
	if polygon.IsExtruded != polygonOther.IsExtruded {
		diffs = append(diffs, polygon.GongMarshallField(stage, "IsExtruded"))
	}
	if polygon.BaseAltitude != polygonOther.BaseAltitude {
		diffs = append(diffs, polygon.GongMarshallField(stage, "BaseAltitude"))
	}
	if polygon.ExtrusionHeight != polygonOther.ExtrusionHeight {
		diffs = append(diffs, polygon.GongMarshallField(stage, "ExtrusionHeight"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (polyline *Polyline) GongDiff(stage *Stage, polylineOther *Polyline) (diffs []string) {
	// insertion point for field diffs
	if polyline.Name != polylineOther.Name {
		diffs = append(diffs, polyline.GongMarshallField(stage, "Name"))
	}
	if polyline.Label != polylineOther.Label {
		diffs = append(diffs, polyline.GongMarshallField(stage, "Label"))
	}
	if polyline.PopupHTML != polylineOther.PopupHTML {
		diffs = append(diffs, polyline.GongMarshallField(stage, "PopupHTML"))
	}
	if ops := __gong__diffSliceOfPointers(stage, polyline, "Coordinates", polylineOther.Coordinates, polyline.Coordinates); ops != "" {
		diffs = append(diffs, ops)
	}
	if polyline.StrokeColor != polylineOther.StrokeColor {
		diffs = append(diffs, polyline.GongMarshallField(stage, "StrokeColor"))
	}
	if polyline.StrokeOpacity != polylineOther.StrokeOpacity {
		diffs = append(diffs, polyline.GongMarshallField(stage, "StrokeOpacity"))
	}
	if polyline.StrokeWidth != polylineOther.StrokeWidth {
		diffs = append(diffs, polyline.GongMarshallField(stage, "StrokeWidth"))
	}
	if polyline.StrokeDashArray != polylineOther.StrokeDashArray {
		diffs = append(diffs, polyline.GongMarshallField(stage, "StrokeDashArray"))
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
