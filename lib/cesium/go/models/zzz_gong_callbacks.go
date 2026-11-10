// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (coordinate *Coordinate) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCoordinateCreateCallback != nil {
		stage.OnAfterCoordinateCreateCallback.OnAfterCreate(stage, coordinate)
	}
}

func (coordinate *Coordinate) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCoordinateUpdateCallback != nil {
		var frontCoordinate *Coordinate
		if front != nil {
			frontCoordinate, _ = front.(*Coordinate)
		}
		stage.OnAfterCoordinateUpdateCallback.OnAfterUpdate(stage, coordinate, frontCoordinate)
	}
}

func (coordinate *Coordinate) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCoordinateDeleteCallback != nil {
		var frontCoordinate *Coordinate
		if front != nil {
			frontCoordinate, _ = front.(*Coordinate)
		}
		stage.OnAfterCoordinateDeleteCallback.OnAfterDelete(stage, coordinate, frontCoordinate)
	}
}

func (layer *Layer) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLayerCreateCallback != nil {
		stage.OnAfterLayerCreateCallback.OnAfterCreate(stage, layer)
	}
}

func (layer *Layer) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLayerUpdateCallback != nil {
		var frontLayer *Layer
		if front != nil {
			frontLayer, _ = front.(*Layer)
		}
		stage.OnAfterLayerUpdateCallback.OnAfterUpdate(stage, layer, frontLayer)
	}
}

func (layer *Layer) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLayerDeleteCallback != nil {
		var frontLayer *Layer
		if front != nil {
			frontLayer, _ = front.(*Layer)
		}
		stage.OnAfterLayerDeleteCallback.OnAfterDelete(stage, layer, frontLayer)
	}
}

func (mapview *MapView) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMapViewCreateCallback != nil {
		stage.OnAfterMapViewCreateCallback.OnAfterCreate(stage, mapview)
	}
}

func (mapview *MapView) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMapViewUpdateCallback != nil {
		var frontMapView *MapView
		if front != nil {
			frontMapView, _ = front.(*MapView)
		}
		stage.OnAfterMapViewUpdateCallback.OnAfterUpdate(stage, mapview, frontMapView)
	}
}

func (mapview *MapView) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMapViewDeleteCallback != nil {
		var frontMapView *MapView
		if front != nil {
			frontMapView, _ = front.(*MapView)
		}
		stage.OnAfterMapViewDeleteCallback.OnAfterDelete(stage, mapview, frontMapView)
	}
}

func (marker *Marker) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMarkerCreateCallback != nil {
		stage.OnAfterMarkerCreateCallback.OnAfterCreate(stage, marker)
	}
}

func (marker *Marker) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkerUpdateCallback != nil {
		var frontMarker *Marker
		if front != nil {
			frontMarker, _ = front.(*Marker)
		}
		stage.OnAfterMarkerUpdateCallback.OnAfterUpdate(stage, marker, frontMarker)
	}
}

func (marker *Marker) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkerDeleteCallback != nil {
		var frontMarker *Marker
		if front != nil {
			frontMarker, _ = front.(*Marker)
		}
		stage.OnAfterMarkerDeleteCallback.OnAfterDelete(stage, marker, frontMarker)
	}
}

func (polygon *Polygon) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPolygonCreateCallback != nil {
		stage.OnAfterPolygonCreateCallback.OnAfterCreate(stage, polygon)
	}
}

func (polygon *Polygon) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPolygonUpdateCallback != nil {
		var frontPolygon *Polygon
		if front != nil {
			frontPolygon, _ = front.(*Polygon)
		}
		stage.OnAfterPolygonUpdateCallback.OnAfterUpdate(stage, polygon, frontPolygon)
	}
}

func (polygon *Polygon) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPolygonDeleteCallback != nil {
		var frontPolygon *Polygon
		if front != nil {
			frontPolygon, _ = front.(*Polygon)
		}
		stage.OnAfterPolygonDeleteCallback.OnAfterDelete(stage, polygon, frontPolygon)
	}
}

func (polyline *Polyline) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPolylineCreateCallback != nil {
		stage.OnAfterPolylineCreateCallback.OnAfterCreate(stage, polyline)
	}
}

func (polyline *Polyline) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPolylineUpdateCallback != nil {
		var frontPolyline *Polyline
		if front != nil {
			frontPolyline, _ = front.(*Polyline)
		}
		stage.OnAfterPolylineUpdateCallback.OnAfterUpdate(stage, polyline, frontPolyline)
	}
}

func (polyline *Polyline) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPolylineDeleteCallback != nil {
		var frontPolyline *Polyline
		if front != nil {
			frontPolyline, _ = front.(*Polyline)
		}
		stage.OnAfterPolylineDeleteCallback.OnAfterDelete(stage, polyline, frontPolyline)
	}
}

