// generated code - do not edit

//insertion point for imports
import { CoordinateAPI } from './coordinate-api'

import { LayerAPI } from './layer-api'

import { MapViewAPI } from './mapview-api'

import { MarkerAPI } from './marker-api'

import { PolygonAPI } from './polygon-api'

import { PolylineAPI } from './polyline-api'


export class BackRepoData {
	// insertion point for declarations
	CoordinateAPIs = new Array<CoordinateAPI>()

	LayerAPIs = new Array<LayerAPI>()

	MapViewAPIs = new Array<MapViewAPI>()

	MarkerAPIs = new Array<MarkerAPI>()

	PolygonAPIs = new Array<PolygonAPI>()

	PolylineAPIs = new Array<PolylineAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.CoordinateAPIs = data?.CoordinateAPIs || [];

		this.LayerAPIs = data?.LayerAPIs || [];

		this.MapViewAPIs = data?.MapViewAPIs || [];

		this.MarkerAPIs = data?.MarkerAPIs || [];

		this.PolygonAPIs = data?.PolygonAPIs || [];

		this.PolylineAPIs = data?.PolylineAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}