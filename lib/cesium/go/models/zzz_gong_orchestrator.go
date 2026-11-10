// generated code - do not edit
package models

// insertion point
// MapViewOrchestrator
type MapViewOrchestrator struct {
}

func (orchestrator *MapViewOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedMapView, backRepoMapView *MapView) {

	stagedMapView.OnAfterUpdate(gongsvgStage, stagedMapView, backRepoMapView)
}

// MarkerOrchestrator
type MarkerOrchestrator struct {
}

func (orchestrator *MarkerOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedMarker, backRepoMarker *Marker) {

	stagedMarker.OnAfterUpdate(gongsvgStage, stagedMarker, backRepoMarker)
}

// PolygonOrchestrator
type PolygonOrchestrator struct {
}

func (orchestrator *PolygonOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedPolygon, backRepoPolygon *Polygon) {

	stagedPolygon.OnAfterUpdate(gongsvgStage, stagedPolygon, backRepoPolygon)
}

// PolylineOrchestrator
type PolylineOrchestrator struct {
}

func (orchestrator *PolylineOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedPolyline, backRepoPolyline *Polyline) {

	stagedPolyline.OnAfterUpdate(gongsvgStage, stagedPolyline, backRepoPolyline)
}

// SetOrchestratorOnAfterUpdate is the Stage method for setting orchestrators.
func (stage *Stage) SetOrchestratorOnAfterUpdate[Type Gongstruct]() {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case MapView:
		stage.OnAfterMapViewUpdateCallback = new(MapViewOrchestrator)
	case Marker:
		stage.OnAfterMarkerUpdateCallback = new(MarkerOrchestrator)
	case Polygon:
		stage.OnAfterPolygonUpdateCallback = new(PolygonOrchestrator)
	case Polyline:
		stage.OnAfterPolylineUpdateCallback = new(PolylineOrchestrator)
	}

}
