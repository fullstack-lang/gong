// generated code - do not edit
import { Injectable, NgZone } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, BehaviorSubject, of } from 'rxjs'
import { shareReplay } from 'rxjs/operators'

// insertion point sub template for services imports
import { CoordinateAPI } from './coordinate-api'
import { Coordinate, CopyCoordinateAPIToCoordinate } from './coordinate'

import { LayerAPI } from './layer-api'
import { Layer, CopyLayerAPIToLayer } from './layer'

import { MapViewAPI } from './mapview-api'
import { MapView, CopyMapViewAPIToMapView } from './mapview'

import { MarkerAPI } from './marker-api'
import { Marker, CopyMarkerAPIToMarker } from './marker'

import { PolygonAPI } from './polygon-api'
import { Polygon, CopyPolygonAPIToPolygon } from './polygon'

import { PolylineAPI } from './polyline-api'
import { Polyline, CopyPolylineAPIToPolyline } from './polyline'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/lib/maplibre/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Coordinates = new Array<Coordinate>() // array of front instances
	map_ID_Coordinate = new Map<number, Coordinate>() // map of front instances

	array_Layers = new Array<Layer>() // array of front instances
	map_ID_Layer = new Map<number, Layer>() // map of front instances

	array_MapViews = new Array<MapView>() // array of front instances
	map_ID_MapView = new Map<number, MapView>() // map of front instances

	array_Markers = new Array<Marker>() // array of front instances
	map_ID_Marker = new Map<number, Marker>() // map of front instances

	array_Polygons = new Array<Polygon>() // array of front instances
	map_ID_Polygon = new Map<number, Polygon>() // map of front instances

	array_Polylines = new Array<Polyline>() // array of front instances
	map_ID_Polyline = new Map<number, Polyline>() // map of front instances


	public GONG__Index = -1

	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'Coordinate':
				return this.array_Coordinates as unknown as Array<Type>
			case 'Layer':
				return this.array_Layers as unknown as Array<Type>
			case 'MapView':
				return this.array_MapViews as unknown as Array<Type>
			case 'Marker':
				return this.array_Markers as unknown as Array<Type>
			case 'Polygon':
				return this.array_Polygons as unknown as Array<Type>
			case 'Polyline':
				return this.array_Polylines as unknown as Array<Type>
			default:
				throw new Error("Type not recognized")
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'Coordinate':
				return this.map_ID_Coordinate as unknown as Map<number, Type>
			case 'Layer':
				return this.map_ID_Layer as unknown as Map<number, Type>
			case 'MapView':
				return this.map_ID_MapView as unknown as Map<number, Type>
			case 'Marker':
				return this.map_ID_Marker as unknown as Map<number, Type>
			case 'Polygon':
				return this.map_ID_Polygon as unknown as Map<number, Type>
			case 'Polyline':
				return this.map_ID_Polyline as unknown as Map<number, Type>
			default:
				throw new Error("Type not recognized")
		}
	}
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
	ID: number = 0 // ID of the calling instance

	// the reverse pointer is the name of the generated field on the destination
	// struct of the ONE-MANY association
	ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
	OrderingMode: boolean = false // if true, this is for ordering items

	// there are different selection mode : ONE_MANY or MANY_MANY
	SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

	// used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
	//
	// In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
	// 
	// in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
	// at the end of the ONE-MANY association
	SourceStruct: string = ""	// The "Aclass"
	SourceField: string = "" // the "AnarrayofbUse"
	IntermediateStruct: string = "" // the "AclassBclassUse" 
	IntermediateStructField: string = "" // the "Bclass" as field
	NextAssociationStruct: string = "" // the "Bclass"

	Name: string = ""
}

export enum SelectionMode {
	ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
	MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
	providedIn: 'root'
})
export class FrontRepoService {

	Name: string = ""

	httpOptions = {
		headers: new HttpHeaders({ 'Content-Type': 'application/json' })
	}

	//
	// Store of all instances of the stack
	//
	frontRepo = new (FrontRepo)

	// Manage open WebSocket connections
	private webSocketConnections = new Map<string, Observable<FrontRepo>>()


	constructor(
		private http: HttpClient,
		private ngZone: NgZone,
	) { }

	//
	// pull returns the FrontRepo Observable
	//
	pull(Name: string = ""): Observable<FrontRepo> {
		this.Name = Name
		return of(this.frontRepo)
	}

	public connectToWebSocket(Name: string): Observable<FrontRepo> {

		// console.log("github.com/fullstack-lang/gong/lib/maplibre/go; connectToWebSocket: started", Name)

		// Check if a connection for this name already exists
		if (this.webSocketConnections.has(Name)) {
			// console.log("github.com/fullstack-lang/gong/lib/maplibre/go; connectToWebSocket: returning existing connection")
			return this.webSocketConnections.get(Name)!
		}

		//
		// Create a new connection
		//
		let host = window.location.host
		const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'

		if (host === 'localhost:4200') {
			host = 'localhost:8080'
		}

		// Construct the base path using the dynamic host and protocol
		// The API path remains the same.
		let basePath = `${protocol}//${host}/api/github.com/fullstack-lang/gong/lib/maplibre/go/v1/ws/stage`

		let params = new HttpParams().set("Name", Name)
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`

		const newConnection$ = new Observable<FrontRepo>(observer => {
			// console.log("github.com/fullstack-lang/gong/lib/maplibre/go; connectToWebSocket: new Observable created")

			let socket: WebSocket | undefined

			const isOfflineMode = window.location.protocol === 'file:' || window.document.getElementById('wasm-progress-container') !== null

			const processData = (dataString: string) => {
				// console.log("github.com/fullstack-lang/gong/lib/maplibre/go; connectToWebSocket: processData called")
				const backRepoData = new BackRepoData(JSON.parse(dataString))
				let frontRepo = new (FrontRepo)()
				frontRepo.GONG__Index = backRepoData.GONG__Index

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				frontRepo.array_Coordinates = []
				frontRepo.map_ID_Coordinate.clear()

				backRepoData.CoordinateAPIs.forEach(
					coordinateAPI => {
						let coordinate = new Coordinate
						frontRepo.array_Coordinates.push(coordinate)
						frontRepo.map_ID_Coordinate.set(coordinateAPI.ID, coordinate)
					}
				)

				// init the arrays
				frontRepo.array_Layers = []
				frontRepo.map_ID_Layer.clear()

				backRepoData.LayerAPIs.forEach(
					layerAPI => {
						let layer = new Layer
						frontRepo.array_Layers.push(layer)
						frontRepo.map_ID_Layer.set(layerAPI.ID, layer)
					}
				)

				// init the arrays
				frontRepo.array_MapViews = []
				frontRepo.map_ID_MapView.clear()

				backRepoData.MapViewAPIs.forEach(
					mapviewAPI => {
						let mapview = new MapView
						frontRepo.array_MapViews.push(mapview)
						frontRepo.map_ID_MapView.set(mapviewAPI.ID, mapview)
					}
				)

				// init the arrays
				frontRepo.array_Markers = []
				frontRepo.map_ID_Marker.clear()

				backRepoData.MarkerAPIs.forEach(
					markerAPI => {
						let marker = new Marker
						frontRepo.array_Markers.push(marker)
						frontRepo.map_ID_Marker.set(markerAPI.ID, marker)
					}
				)

				// init the arrays
				frontRepo.array_Polygons = []
				frontRepo.map_ID_Polygon.clear()

				backRepoData.PolygonAPIs.forEach(
					polygonAPI => {
						let polygon = new Polygon
						frontRepo.array_Polygons.push(polygon)
						frontRepo.map_ID_Polygon.set(polygonAPI.ID, polygon)
					}
				)

				// init the arrays
				frontRepo.array_Polylines = []
				frontRepo.map_ID_Polyline.clear()

				backRepoData.PolylineAPIs.forEach(
					polylineAPI => {
						let polyline = new Polyline
						frontRepo.array_Polylines.push(polyline)
						frontRepo.map_ID_Polyline.set(polylineAPI.ID, polyline)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.CoordinateAPIs.forEach(
					coordinateAPI => {
						let coordinate = frontRepo.map_ID_Coordinate.get(coordinateAPI.ID)
						CopyCoordinateAPIToCoordinate(coordinateAPI, coordinate!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.LayerAPIs.forEach(
					layerAPI => {
						let layer = frontRepo.map_ID_Layer.get(layerAPI.ID)
						CopyLayerAPIToLayer(layerAPI, layer!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.MapViewAPIs.forEach(
					mapviewAPI => {
						let mapview = frontRepo.map_ID_MapView.get(mapviewAPI.ID)
						CopyMapViewAPIToMapView(mapviewAPI, mapview!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.MarkerAPIs.forEach(
					markerAPI => {
						let marker = frontRepo.map_ID_Marker.get(markerAPI.ID)
						CopyMarkerAPIToMarker(markerAPI, marker!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.PolygonAPIs.forEach(
					polygonAPI => {
						let polygon = frontRepo.map_ID_Polygon.get(polygonAPI.ID)
						CopyPolygonAPIToPolygon(polygonAPI, polygon!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.PolylineAPIs.forEach(
					polylineAPI => {
						let polyline = frontRepo.map_ID_Polyline.get(polylineAPI.ID)
						CopyPolylineAPIToPolyline(polylineAPI, polyline!, frontRepo)
					}
				)


				this.ngZone.run(() => {
					observer.next(frontRepo)
				})
			}

			// 3. Connection Loop
			const attemptConnection = (retries: number): void => {
				// console.log("github.com/fullstack-lang/gong/lib/maplibre/go; attemptConnection: retries =", retries, "isOfflineMode =", isOfflineMode)

				// A. WASM OFFLINE MODE (Check if Go is ready)
				if ((window as any).openWasmSocket) {
					// console.log("github.com/fullstack-lang/gong/lib/maplibre/go; attemptConnection: openWasmSocket exists, calling it");
					(window as any).openWasmSocket("github.com/fullstack-lang/gong/lib/maplibre/go", Name, processData);
					return;
				}

				// B. WAITING FOR WASM
				if (isOfflineMode && retries > 0) {
					// console.log("github.com/fullstack-lang/gong/lib/maplibre/go; attemptConnection: WAITING FOR WASM. Retries left:", retries)
					setTimeout(() => attemptConnection(retries - 1), 100);
					return;
				}

				// C. STANDARD SERVER MODE
				if (!isOfflineMode) {
					// console.log("github.com/fullstack-lang/gong/lib/maplibre/go; attemptConnection: STANDARD SERVER MODE. url =", url)
					socket = new WebSocket(url)
					socket.onopen = (event) => {
						// console.log("github.com/fullstack-lang/gong/lib/maplibre/go; WebSocket: onopen", event)
					}
					socket.onmessage = event => {
						// console.log("github.com/fullstack-lang/gong/lib/maplibre/go; WebSocket: onmessage")
						processData(event.data)
					}
					socket.onerror = event => {
						console.error("github.com/fullstack-lang/gong/lib/maplibre/go WebSocket: onerror", event)
						observer.error(event)
					}
					socket.onclose = (event) => {
						// console.log("github.com/fullstack-lang/gong/lib/maplibre/go; WebSocket: onclose", event)
						observer.complete()
					}
				} else {
					console.error("github.com/fullstack-lang/gong/lib/maplibre/go, attemptConnection: Offline mode detected, but WASM backend failed to load.")
					observer.error("Offline mode detected, but WASM backend failed to load.");
				}
			};

			attemptConnection(50);

			// Teardown logic: Called when the last subscriber unsubscribes.
			return () => {
				this.webSocketConnections.delete(Name) // Remove from cache
				if (socket) {
					socket.close()
				}
			}
		}).pipe(
			// This is the key:
			// - shareReplay makes this a "multicast" observable, sharing the single WebSocket among subscribers.
			// - { bufferSize: 1, refCount: true } means:
			//   - bufferSize: 1 => new subscribers get the last emitted value immediately.
			//   - refCount: true => the connection starts with the first subscriber and stops with the last.
			shareReplay({ bufferSize: 1, refCount: true })
		)

		// Store the new connection observable in the map
		this.webSocketConnections.set(Name, newConnection$)
		return newConnection$
	}
}

// insertion point for get unique ID per struct 
export function getCoordinateUniqueID(id: number): number {
	return 31 * id
}
export function getLayerUniqueID(id: number): number {
	return 37 * id
}
export function getMapViewUniqueID(id: number): number {
	return 41 * id
}
export function getMarkerUniqueID(id: number): number {
	return 43 * id
}
export function getPolygonUniqueID(id: number): number {
	return 47 * id
}
export function getPolylineUniqueID(id: number): number {
	return 53 * id
}
