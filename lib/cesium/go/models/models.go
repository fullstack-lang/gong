package models

// MapView is the root singleton/view configuration for a Cesium instance.
type MapView struct {
	Name string

	// StyleURL specifies the Cesium style JSON URL or preset name.
	// If empty or "osm", defaults to an OpenStreetMap raster style requiring no API key.
	// Other presets: "carto-positron", "carto-dark", "openfreemap-liberty", or any https:// style.json URL.
	StyleURL string

	// Center coordinates in degrees (WGS84)
	Longitude float64
	Latitude  float64

	// Viewport parameters
	Zoom    float64
	Pitch   float64
	Bearing float64

	// Layers ordered from bottom to top
	Layers []*Layer

	// OnUpdate is triggered when the user pans, zooms, tilts, or rotates the map in the frontend
	OnUpdate func(updatedMap *MapView)
}

func (m *MapView) OnAfterUpdate(stage *Stage, stagedMap, frontMap *MapView) {
	stagedMap.Longitude = frontMap.Longitude
	stagedMap.Latitude = frontMap.Latitude
	stagedMap.Zoom = frontMap.Zoom
	stagedMap.Pitch = frontMap.Pitch
	stagedMap.Bearing = frontMap.Bearing

	if stagedMap.OnUpdate != nil {
		stagedMap.OnUpdate(stagedMap)
	}
}

// Layer groups geographic primitives and SVG markers that can be toggled together.
type Layer struct {
	Name      string
	IsVisible bool

	Polygons  []*Polygon
	Polylines []*Polyline
	Markers   []*Marker
}

// Coordinate represents a geographic point in WGS84 (Longitude, Latitude in degrees, Altitude in meters).
type Coordinate struct {
	Name      string
	Longitude float64
	Latitude  float64
	Altitude  float64
}

// Polygon represents a closed geographic area (e.g. an AIXM Airspace volume or airport surface).
// Supports both 2D fill/stroke rendering and 3D extrusion when IsExtruded is true.
type Polygon struct {
	Name      string
	Label     string
	PopupHTML string

	Coordinates []*Coordinate

	FillColor     string
	FillOpacity   float64
	StrokeColor   string
	StrokeOpacity float64
	StrokeWidth   float64

	// 3D extrusion parameters (in meters)
	IsExtruded      bool
	BaseAltitude    float64
	ExtrusionHeight float64

	OnClick func(polygon *Polygon)
}

func (p *Polygon) OnAfterUpdate(stage *Stage, stagedPolygon, frontPolygon *Polygon) {
	if stagedPolygon.OnClick != nil {
		stagedPolygon.OnClick(stagedPolygon)
	}
}

// Polyline represents a geographic sequence of connected segments (e.g. AIXM routes, airways, procedures).
type Polyline struct {
	Name      string
	Label     string
	PopupHTML string

	Coordinates []*Coordinate

	StrokeColor     string
	StrokeOpacity   float64
	StrokeWidth     float64
	StrokeDashArray string

	OnClick func(polyline *Polyline)
}

func (pl *Polyline) OnAfterUpdate(stage *Stage, stagedPolyline, frontPolyline *Polyline) {
	if stagedPolyline.OnClick != nil {
		stagedPolyline.OnClick(stagedPolyline)
	}
}

type MarkerPitchAlignmentEnum string

const (
	MarkerAlignmentAuto     MarkerPitchAlignmentEnum = "auto"
	MarkerAlignmentMap      MarkerPitchAlignmentEnum = "map"
	MarkerAlignmentViewport MarkerPitchAlignmentEnum = "viewport"
)

type MarkerAnchorEnum string

const (
	MarkerAnchorCenter      MarkerAnchorEnum = "center"
	MarkerAnchorTop         MarkerAnchorEnum = "top"
	MarkerAnchorBottom      MarkerAnchorEnum = "bottom"
	MarkerAnchorLeft        MarkerAnchorEnum = "left"
	MarkerAnchorRight       MarkerAnchorEnum = "right"
	MarkerAnchorTopLeft     MarkerAnchorEnum = "top-left"
	MarkerAnchorTopRight    MarkerAnchorEnum = "top-right"
	MarkerAnchorBottomLeft  MarkerAnchorEnum = "bottom-left"
	MarkerAnchorBottomRight MarkerAnchorEnum = "bottom-right"
)

// Marker places a point or a custom SVG object at (Longitude, Latitude) on the map.
type Marker struct {
	Name      string
	Label     string
	PopupHTML string

	Longitude float64
	Latitude  float64
	Altitude  float64

	// Raw SVG markup rendered on the map. If empty, a default marker pin is used.
	//gong:width 600 gong:height 300
	SVG string

	Color    string
	Scale    float64
	Rotation float64

	RotationAlignment MarkerPitchAlignmentEnum
	PitchAlignment    MarkerPitchAlignmentEnum
	Anchor            MarkerAnchorEnum

	IsDraggable bool

	OnClick func(marker *Marker)
}

func (mk *Marker) OnAfterUpdate(stage *Stage, stagedMarker, frontMarker *Marker) {
	stagedMarker.Longitude = frontMarker.Longitude
	stagedMarker.Latitude = frontMarker.Latitude

	if stagedMarker.OnClick != nil {
		stagedMarker.OnClick(stagedMarker)
	}
}
