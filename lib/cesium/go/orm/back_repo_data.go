// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	CoordinateAPIs []*CoordinateAPI

	LayerAPIs []*LayerAPI

	MapViewAPIs []*MapViewAPI

	MarkerAPIs []*MarkerAPI

	PolygonAPIs []*PolygonAPI

	PolylineAPIs []*PolylineAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, coordinateDB := range backRepo.BackRepoCoordinate.Map_CoordinateDBID_CoordinateDB {

		var coordinateAPI CoordinateAPI
		coordinateAPI.ID = coordinateDB.ID
		coordinateAPI.CoordinatePointersEncoding = coordinateDB.CoordinatePointersEncoding
		coordinateDB.CopyBasicFieldsToCoordinate_WOP(&coordinateAPI.Coordinate_WOP)

		backRepoData.CoordinateAPIs = append(backRepoData.CoordinateAPIs, &coordinateAPI)
	}

	for _, layerDB := range backRepo.BackRepoLayer.Map_LayerDBID_LayerDB {

		var layerAPI LayerAPI
		layerAPI.ID = layerDB.ID
		layerAPI.LayerPointersEncoding = layerDB.LayerPointersEncoding
		layerDB.CopyBasicFieldsToLayer_WOP(&layerAPI.Layer_WOP)

		backRepoData.LayerAPIs = append(backRepoData.LayerAPIs, &layerAPI)
	}

	for _, mapviewDB := range backRepo.BackRepoMapView.Map_MapViewDBID_MapViewDB {

		var mapviewAPI MapViewAPI
		mapviewAPI.ID = mapviewDB.ID
		mapviewAPI.MapViewPointersEncoding = mapviewDB.MapViewPointersEncoding
		mapviewDB.CopyBasicFieldsToMapView_WOP(&mapviewAPI.MapView_WOP)

		backRepoData.MapViewAPIs = append(backRepoData.MapViewAPIs, &mapviewAPI)
	}

	for _, markerDB := range backRepo.BackRepoMarker.Map_MarkerDBID_MarkerDB {

		var markerAPI MarkerAPI
		markerAPI.ID = markerDB.ID
		markerAPI.MarkerPointersEncoding = markerDB.MarkerPointersEncoding
		markerDB.CopyBasicFieldsToMarker_WOP(&markerAPI.Marker_WOP)

		backRepoData.MarkerAPIs = append(backRepoData.MarkerAPIs, &markerAPI)
	}

	for _, polygonDB := range backRepo.BackRepoPolygon.Map_PolygonDBID_PolygonDB {

		var polygonAPI PolygonAPI
		polygonAPI.ID = polygonDB.ID
		polygonAPI.PolygonPointersEncoding = polygonDB.PolygonPointersEncoding
		polygonDB.CopyBasicFieldsToPolygon_WOP(&polygonAPI.Polygon_WOP)

		backRepoData.PolygonAPIs = append(backRepoData.PolygonAPIs, &polygonAPI)
	}

	for _, polylineDB := range backRepo.BackRepoPolyline.Map_PolylineDBID_PolylineDB {

		var polylineAPI PolylineAPI
		polylineAPI.ID = polylineDB.ID
		polylineAPI.PolylinePointersEncoding = polylineDB.PolylinePointersEncoding
		polylineDB.CopyBasicFieldsToPolyline_WOP(&polylineAPI.Polyline_WOP)

		backRepoData.PolylineAPIs = append(backRepoData.PolylineAPIs, &polylineAPI)
	}

}
