package models

import "fmt"

// GenerateReferenceScene populates a stage with sample aviation/AIXM airspaces, routes, and SVG markers.
func GenerateReferenceScene(stage *Stage) {
	stage.Reset()

	// 1. Airspace Layer (Polygons)
	airspaceLayer := (&Layer{
		Name:      "AIXM Airspaces",
		IsVisible: true,
	}).Stage(stage)

	// Sample TMA Polygon around Paris area
	tmaCoords := [][2]float64{
		{1.65, 49.25},
		{2.95, 49.25},
		{3.15, 48.65},
		{2.55, 48.35},
		{1.80, 48.45},
		{1.65, 49.25},
	}
	var tmaPolygonCoords []*Coordinate
	for i, pt := range tmaCoords {
		c := (&Coordinate{
			Name:      fmt.Sprintf("TMA-PT-%d", i+1),
			Longitude: pt[0],
			Latitude:  pt[1],
		}).Stage(stage)
		tmaPolygonCoords = append(tmaPolygonCoords, c)
	}

	tmaPolygon := (&Polygon{
		Name:          "LFPG TMA Sector 1",
		Label:         "LFPG TMA 1 (FL065 / 2000ft)",
		PopupHTML:     "<b>LFPG TMA Sector 1</b><br/>Class A<br/>Upper: FL065 / Lower: 2000ft AMSL",
		Coordinates:   tmaPolygonCoords,
		FillColor:     "#2563eb",
		FillOpacity:   0.22,
		StrokeColor:   "#1d4ed8",
		StrokeOpacity: 0.9,
		StrokeWidth:   2.5,
	}).Stage(stage)
	airspaceLayer.Polygons = append(airspaceLayer.Polygons, tmaPolygon)

	// Sample CTR Polygon (3D Extruded)
	ctrCoords := [][2]float64{
		{2.35, 49.08},
		{2.72, 49.08},
		{2.72, 48.94},
		{2.35, 48.94},
		{2.35, 49.08},
	}
	var ctrPolygonCoords []*Coordinate
	for i, pt := range ctrCoords {
		c := (&Coordinate{
			Name:      fmt.Sprintf("CTR-PT-%d", i+1),
			Longitude: pt[0],
			Latitude:  pt[1],
		}).Stage(stage)
		ctrPolygonCoords = append(ctrPolygonCoords, c)
	}

	ctrPolygon := (&Polygon{
		Name:            "LFPG CTR",
		Label:           "LFPG CTR (SFC - 1500ft)",
		PopupHTML:       "<b>LFPG CTR</b><br/>Class D<br/>SFC to 1500ft AMSL",
		Coordinates:     ctrPolygonCoords,
		FillColor:       "#f59e0b",
		FillOpacity:     0.45,
		StrokeColor:     "#b45309",
		StrokeOpacity:   1.0,
		StrokeWidth:     2.0,
		IsExtruded:      true,
		BaseAltitude:    0,
		ExtrusionHeight: 1500,
	}).Stage(stage)
	airspaceLayer.Polygons = append(airspaceLayer.Polygons, ctrPolygon)

	// 2. ATS Routes Layer (Polylines)
	routeLayer := (&Layer{
		Name:      "ATS Routes",
		IsVisible: true,
	}).Stage(stage)

	routePoints := [][2]float64{
		{1.38, 49.00}, // EVX
		{2.03, 49.09}, // PON
		{2.55, 49.01}, // PGS
		{3.10, 48.85}, // CLM
	}
	var routeCoords []*Coordinate
	for i, pt := range routePoints {
		c := (&Coordinate{
			Name:      fmt.Sprintf("UM164-WP-%d", i+1),
			Longitude: pt[0],
			Latitude:  pt[1],
		}).Stage(stage)
		routeCoords = append(routeCoords, c)
	}

	atsRoute := (&Polyline{
		Name:            "ATS Route UM164",
		Label:           "UM164",
		PopupHTML:       "<b>ATS Route UM164</b><br/>EVX - PON - PGS - CLM",
		Coordinates:     routeCoords,
		StrokeColor:     "#0f766e",
		StrokeOpacity:   0.95,
		StrokeWidth:     3.0,
		StrokeDashArray: "6, 3",
	}).Stage(stage)
	routeLayer.Polylines = append(routeLayer.Polylines, atsRoute)

	// 3. SVG Objects & Navaids Layer (Markers)
	svgLayer := (&Layer{
		Name:      "SVG Symbols & Navaids",
		IsVisible: true,
	}).Stage(stage)

	vorSVG := `<svg width="36" height="36" viewBox="0 0 36 36" xmlns="http://www.w3.org/2000/svg">
  <polygon points="18,3 31,10.5 31,25.5 18,33 5,25.5 5,10.5" fill="rgba(255,255,255,0.85)" stroke="#1e3a8a" stroke-width="2.5"/>
  <circle cx="18" cy="18" r="3" fill="#1e3a8a"/>
</svg>`

	pgsVor := (&Marker{
		Name:              "PGS VOR/DME",
		Label:             "PGS 117.05",
		PopupHTML:         "<b>PARIS CHARLES DE GAULLE (PGS)</b><br/>VOR/DME 117.05 MHz",
		Longitude:         2.548,
		Latitude:          49.010,
		SVG:               vorSVG,
		Scale:             1.0,
		Anchor:            MarkerAnchorCenter,
		RotationAlignment: MarkerAlignmentViewport,
		PitchAlignment:    MarkerAlignmentViewport,
	}).Stage(stage)

	ponVor := (&Marker{
		Name:              "PON VOR/DME",
		Label:             "PON 111.60",
		PopupHTML:         "<b>PONTOISE (PON)</b><br/>VOR/DME 111.60 MHz",
		Longitude:         2.035,
		Latitude:          49.096,
		SVG:               vorSVG,
		Scale:             1.0,
		Anchor:            MarkerAnchorCenter,
		RotationAlignment: MarkerAlignmentViewport,
		PitchAlignment:    MarkerAlignmentViewport,
	}).Stage(stage)

	aircraftSVG := `<svg width="40" height="40" viewBox="0 0 40 40" xmlns="http://www.w3.org/2000/svg">
  <path d="M20 3 L23 15 L36 21 L36 24 L23 21 L22 31 L27 35 L27 37 L20 35 L13 37 L13 35 L18 31 L17 21 L4 24 L4 21 L17 15 Z" fill="#dc2626" stroke="#ffffff" stroke-width="1.5"/>
</svg>`

	aircraftMarker := (&Marker{
		Name:              "AFR123",
		Label:             "AFR123 · FL180",
		PopupHTML:         "<b>AFR123</b><br/>A320 · Heading 095° · FL180<br/><i>(Drag marker to update coordinates)</i>",
		Longitude:         2.25,
		Latitude:          49.05,
		SVG:               aircraftSVG,
		Scale:             1.1,
		Rotation:          95,
		RotationAlignment: MarkerAlignmentMap,
		PitchAlignment:    MarkerAlignmentMap,
		Anchor:            MarkerAnchorCenter,
		IsDraggable:       true,
	}).Stage(stage)

	svgLayer.Markers = append(svgLayer.Markers, pgsVor, ponVor, aircraftMarker)

	// Root MapView instance
	(&MapView{
		Name:      "AIXM & SVG Map",
		StyleURL:  "offline",
		Longitude: 2.35,
		Latitude:  48.90,
		Zoom:      8.5,
		Pitch:     25,
		Bearing:   0,
		Layers: []*Layer{
			airspaceLayer,
			routeLayer,
			svgLayer,
		},
	}).Stage(stage)

	stage.Commit()
}
