// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type Coordinate_WOP struct {
	// insertion point

	Name string

	Longitude float64

	Latitude float64

	Altitude float64
}

func (from *Coordinate) GongCopyBasicFields(to *Coordinate) {
	// insertion point
	*to = *from
}

type Layer_WOP struct {
	// insertion point

	Name string

	IsVisible bool
}

func (from *Layer) GongCopyBasicFields(to *Layer) {
	// insertion point
	to.Name = from.Name
	to.IsVisible = from.IsVisible
}

type MapView_WOP struct {
	// insertion point

	Name string

	StyleURL string

	Longitude float64

	Latitude float64

	Zoom float64

	Pitch float64

	Bearing float64
}

func (from *MapView) GongCopyBasicFields(to *MapView) {
	// insertion point
	to.Name = from.Name
	to.StyleURL = from.StyleURL
	to.Longitude = from.Longitude
	to.Latitude = from.Latitude
	to.Zoom = from.Zoom
	to.Pitch = from.Pitch
	to.Bearing = from.Bearing
}

type Marker_WOP struct {
	// insertion point

	Name string

	Label string

	PopupHTML string

	Longitude float64

	Latitude float64

	Altitude float64

	SVG string

	Color string

	Scale float64

	Rotation float64

	RotationAlignment MarkerPitchAlignmentEnum

	PitchAlignment MarkerPitchAlignmentEnum

	Anchor MarkerAnchorEnum

	IsDraggable bool
}

func (from *Marker) GongCopyBasicFields(to *Marker) {
	// insertion point
	*to = *from
}

type Polygon_WOP struct {
	// insertion point

	Name string

	Label string

	PopupHTML string

	FillColor string

	FillOpacity float64

	StrokeColor string

	StrokeOpacity float64

	StrokeWidth float64

	IsExtruded bool

	BaseAltitude float64

	ExtrusionHeight float64
}

func (from *Polygon) GongCopyBasicFields(to *Polygon) {
	// insertion point
	to.Name = from.Name
	to.Label = from.Label
	to.PopupHTML = from.PopupHTML
	to.FillColor = from.FillColor
	to.FillOpacity = from.FillOpacity
	to.StrokeColor = from.StrokeColor
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.IsExtruded = from.IsExtruded
	to.BaseAltitude = from.BaseAltitude
	to.ExtrusionHeight = from.ExtrusionHeight
}

type Polyline_WOP struct {
	// insertion point

	Name string

	Label string

	PopupHTML string

	StrokeColor string

	StrokeOpacity float64

	StrokeWidth float64

	StrokeDashArray string
}

func (from *Polyline) GongCopyBasicFields(to *Polyline) {
	// insertion point
	to.Name = from.Name
	to.Label = from.Label
	to.PopupHTML = from.PopupHTML
	to.StrokeColor = from.StrokeColor
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
}

// end of insertion point
