import * as i0 from '@angular/core';
import { NgModule, Injectable, DOCUMENT, Inject } from '@angular/core';
import * as i1 from '@angular/common/http';
import { HttpHeaders, HttpParams } from '@angular/common/http';
import { of, Observable, interval, BehaviorSubject } from 'rxjs';
import { shareReplay, switchMap, catchError, tap } from 'rxjs/operators';
import * as i2 from '@angular/common';
import { DOCUMENT as DOCUMENT$1 } from '@angular/common';

// generated code - do not edit
class MaplibreModule {
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: MaplibreModule, deps: [], target: i0.ɵɵFactoryTarget.NgModule });
    static ɵmod = i0.ɵɵngDeclareNgModule({ minVersion: "14.0.0", version: "22.1.4", ngImport: i0, type: MaplibreModule });
    static ɵinj = i0.ɵɵngDeclareInjector({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: MaplibreModule });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: MaplibreModule, decorators: [{
            type: NgModule,
            args: [{
                    declarations: [],
                    imports: [],
                    exports: [],
                    providers: [],
                }]
        }] });

// generated code - do not edit
class Coordinate {
    static GONGSTRUCT_NAME = "Coordinate";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    Longitude = 0;
    Latitude = 0;
    Altitude = 0;
    // insertion point for pointers and slices of pointers declarations
    CreatedAt;
    DeletedAt;
}
function CopyCoordinateToCoordinateAPI(coordinate, coordinateAPI) {
    coordinateAPI.CreatedAt = coordinate.CreatedAt;
    coordinateAPI.DeletedAt = coordinate.DeletedAt;
    coordinateAPI.ID = coordinate.ID;
    // insertion point for basic fields copy operations
    coordinateAPI.Name = coordinate.Name;
    coordinateAPI.Longitude = coordinate.Longitude;
    coordinateAPI.Latitude = coordinate.Latitude;
    coordinateAPI.Altitude = coordinate.Altitude;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}
// CopyCoordinateAPIToCoordinate update basic, pointers and slice of pointers fields of coordinate
// from respectively the basic fields and encoded fields of pointers and slices of pointers of coordinateAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyCoordinateAPIToCoordinate(coordinateAPI, coordinate, frontRepo) {
    coordinate.CreatedAt = coordinateAPI.CreatedAt;
    coordinate.DeletedAt = coordinateAPI.DeletedAt;
    coordinate.ID = coordinateAPI.ID;
    // insertion point for basic fields copy operations
    coordinate.Name = coordinateAPI.Name;
    coordinate.Longitude = coordinateAPI.Longitude;
    coordinate.Latitude = coordinateAPI.Latitude;
    coordinate.Altitude = coordinateAPI.Altitude;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}

// generated code - do not edit
class Layer {
    static GONGSTRUCT_NAME = "Layer";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    IsVisible = false;
    // insertion point for pointers and slices of pointers declarations
    Polygons = [];
    Polylines = [];
    Markers = [];
    CreatedAt;
    DeletedAt;
}
function CopyLayerToLayerAPI(layer, layerAPI) {
    layerAPI.CreatedAt = layer.CreatedAt;
    layerAPI.DeletedAt = layer.DeletedAt;
    layerAPI.ID = layer.ID;
    // insertion point for basic fields copy operations
    layerAPI.Name = layer.Name;
    layerAPI.IsVisible = layer.IsVisible;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
    layerAPI.LayerPointersEncoding.Polygons = [];
    for (let _polygon of layer.Polygons) {
        layerAPI.LayerPointersEncoding.Polygons.push(_polygon.ID);
    }
    layerAPI.LayerPointersEncoding.Polylines = [];
    for (let _polyline of layer.Polylines) {
        layerAPI.LayerPointersEncoding.Polylines.push(_polyline.ID);
    }
    layerAPI.LayerPointersEncoding.Markers = [];
    for (let _marker of layer.Markers) {
        layerAPI.LayerPointersEncoding.Markers.push(_marker.ID);
    }
}
// CopyLayerAPIToLayer update basic, pointers and slice of pointers fields of layer
// from respectively the basic fields and encoded fields of pointers and slices of pointers of layerAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyLayerAPIToLayer(layerAPI, layer, frontRepo) {
    layer.CreatedAt = layerAPI.CreatedAt;
    layer.DeletedAt = layerAPI.DeletedAt;
    layer.ID = layerAPI.ID;
    // insertion point for basic fields copy operations
    layer.Name = layerAPI.Name;
    layer.IsVisible = layerAPI.IsVisible;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
    if (!Array.isArray(layerAPI.LayerPointersEncoding.Polygons)) {
        console.error('Rects is not an array:', layerAPI.LayerPointersEncoding.Polygons);
        return;
    }
    layer.Polygons = new Array();
    for (let _id of layerAPI.LayerPointersEncoding.Polygons) {
        let _polygon = frontRepo.map_ID_Polygon.get(_id);
        if (_polygon != undefined) {
            layer.Polygons.push(_polygon);
        }
    }
    if (!Array.isArray(layerAPI.LayerPointersEncoding.Polylines)) {
        console.error('Rects is not an array:', layerAPI.LayerPointersEncoding.Polylines);
        return;
    }
    layer.Polylines = new Array();
    for (let _id of layerAPI.LayerPointersEncoding.Polylines) {
        let _polyline = frontRepo.map_ID_Polyline.get(_id);
        if (_polyline != undefined) {
            layer.Polylines.push(_polyline);
        }
    }
    if (!Array.isArray(layerAPI.LayerPointersEncoding.Markers)) {
        console.error('Rects is not an array:', layerAPI.LayerPointersEncoding.Markers);
        return;
    }
    layer.Markers = new Array();
    for (let _id of layerAPI.LayerPointersEncoding.Markers) {
        let _marker = frontRepo.map_ID_Marker.get(_id);
        if (_marker != undefined) {
            layer.Markers.push(_marker);
        }
    }
}

// generated code - do not edit
class MapView {
    static GONGSTRUCT_NAME = "MapView";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StyleURL = "";
    Longitude = 0;
    Latitude = 0;
    Zoom = 0;
    Pitch = 0;
    Bearing = 0;
    // insertion point for pointers and slices of pointers declarations
    Layers = [];
    CreatedAt;
    DeletedAt;
}
function CopyMapViewToMapViewAPI(mapview, mapviewAPI) {
    mapviewAPI.CreatedAt = mapview.CreatedAt;
    mapviewAPI.DeletedAt = mapview.DeletedAt;
    mapviewAPI.ID = mapview.ID;
    // insertion point for basic fields copy operations
    mapviewAPI.Name = mapview.Name;
    mapviewAPI.StyleURL = mapview.StyleURL;
    mapviewAPI.Longitude = mapview.Longitude;
    mapviewAPI.Latitude = mapview.Latitude;
    mapviewAPI.Zoom = mapview.Zoom;
    mapviewAPI.Pitch = mapview.Pitch;
    mapviewAPI.Bearing = mapview.Bearing;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
    mapviewAPI.MapViewPointersEncoding.Layers = [];
    for (let _layer of mapview.Layers) {
        mapviewAPI.MapViewPointersEncoding.Layers.push(_layer.ID);
    }
}
// CopyMapViewAPIToMapView update basic, pointers and slice of pointers fields of mapview
// from respectively the basic fields and encoded fields of pointers and slices of pointers of mapviewAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyMapViewAPIToMapView(mapviewAPI, mapview, frontRepo) {
    mapview.CreatedAt = mapviewAPI.CreatedAt;
    mapview.DeletedAt = mapviewAPI.DeletedAt;
    mapview.ID = mapviewAPI.ID;
    // insertion point for basic fields copy operations
    mapview.Name = mapviewAPI.Name;
    mapview.StyleURL = mapviewAPI.StyleURL;
    mapview.Longitude = mapviewAPI.Longitude;
    mapview.Latitude = mapviewAPI.Latitude;
    mapview.Zoom = mapviewAPI.Zoom;
    mapview.Pitch = mapviewAPI.Pitch;
    mapview.Bearing = mapviewAPI.Bearing;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
    if (!Array.isArray(mapviewAPI.MapViewPointersEncoding.Layers)) {
        console.error('Rects is not an array:', mapviewAPI.MapViewPointersEncoding.Layers);
        return;
    }
    mapview.Layers = new Array();
    for (let _id of mapviewAPI.MapViewPointersEncoding.Layers) {
        let _layer = frontRepo.map_ID_Layer.get(_id);
        if (_layer != undefined) {
            mapview.Layers.push(_layer);
        }
    }
}

// generated code - do not edit
class Marker {
    static GONGSTRUCT_NAME = "Marker";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    Label = "";
    PopupHTML = "";
    Longitude = 0;
    Latitude = 0;
    Altitude = 0;
    SVG = "";
    Color = "";
    Scale = 0;
    Rotation = 0;
    RotationAlignment = "";
    PitchAlignment = "";
    Anchor = "";
    IsDraggable = false;
    // insertion point for pointers and slices of pointers declarations
    CreatedAt;
    DeletedAt;
}
function CopyMarkerToMarkerAPI(marker, markerAPI) {
    markerAPI.CreatedAt = marker.CreatedAt;
    markerAPI.DeletedAt = marker.DeletedAt;
    markerAPI.ID = marker.ID;
    // insertion point for basic fields copy operations
    markerAPI.Name = marker.Name;
    markerAPI.Label = marker.Label;
    markerAPI.PopupHTML = marker.PopupHTML;
    markerAPI.Longitude = marker.Longitude;
    markerAPI.Latitude = marker.Latitude;
    markerAPI.Altitude = marker.Altitude;
    markerAPI.SVG = marker.SVG;
    markerAPI.Color = marker.Color;
    markerAPI.Scale = marker.Scale;
    markerAPI.Rotation = marker.Rotation;
    markerAPI.RotationAlignment = marker.RotationAlignment;
    markerAPI.PitchAlignment = marker.PitchAlignment;
    markerAPI.Anchor = marker.Anchor;
    markerAPI.IsDraggable = marker.IsDraggable;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}
// CopyMarkerAPIToMarker update basic, pointers and slice of pointers fields of marker
// from respectively the basic fields and encoded fields of pointers and slices of pointers of markerAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyMarkerAPIToMarker(markerAPI, marker, frontRepo) {
    marker.CreatedAt = markerAPI.CreatedAt;
    marker.DeletedAt = markerAPI.DeletedAt;
    marker.ID = markerAPI.ID;
    // insertion point for basic fields copy operations
    marker.Name = markerAPI.Name;
    marker.Label = markerAPI.Label;
    marker.PopupHTML = markerAPI.PopupHTML;
    marker.Longitude = markerAPI.Longitude;
    marker.Latitude = markerAPI.Latitude;
    marker.Altitude = markerAPI.Altitude;
    marker.SVG = markerAPI.SVG;
    marker.Color = markerAPI.Color;
    marker.Scale = markerAPI.Scale;
    marker.Rotation = markerAPI.Rotation;
    marker.RotationAlignment = markerAPI.RotationAlignment;
    marker.PitchAlignment = markerAPI.PitchAlignment;
    marker.Anchor = markerAPI.Anchor;
    marker.IsDraggable = markerAPI.IsDraggable;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}

// generated code - do not edit
class Polygon {
    static GONGSTRUCT_NAME = "Polygon";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    Label = "";
    PopupHTML = "";
    FillColor = "";
    FillOpacity = 0;
    StrokeColor = "";
    StrokeOpacity = 0;
    StrokeWidth = 0;
    IsExtruded = false;
    BaseAltitude = 0;
    ExtrusionHeight = 0;
    // insertion point for pointers and slices of pointers declarations
    Coordinates = [];
    CreatedAt;
    DeletedAt;
}
function CopyPolygonToPolygonAPI(polygon, polygonAPI) {
    polygonAPI.CreatedAt = polygon.CreatedAt;
    polygonAPI.DeletedAt = polygon.DeletedAt;
    polygonAPI.ID = polygon.ID;
    // insertion point for basic fields copy operations
    polygonAPI.Name = polygon.Name;
    polygonAPI.Label = polygon.Label;
    polygonAPI.PopupHTML = polygon.PopupHTML;
    polygonAPI.FillColor = polygon.FillColor;
    polygonAPI.FillOpacity = polygon.FillOpacity;
    polygonAPI.StrokeColor = polygon.StrokeColor;
    polygonAPI.StrokeOpacity = polygon.StrokeOpacity;
    polygonAPI.StrokeWidth = polygon.StrokeWidth;
    polygonAPI.IsExtruded = polygon.IsExtruded;
    polygonAPI.BaseAltitude = polygon.BaseAltitude;
    polygonAPI.ExtrusionHeight = polygon.ExtrusionHeight;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
    polygonAPI.PolygonPointersEncoding.Coordinates = [];
    for (let _coordinate of polygon.Coordinates) {
        polygonAPI.PolygonPointersEncoding.Coordinates.push(_coordinate.ID);
    }
}
// CopyPolygonAPIToPolygon update basic, pointers and slice of pointers fields of polygon
// from respectively the basic fields and encoded fields of pointers and slices of pointers of polygonAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyPolygonAPIToPolygon(polygonAPI, polygon, frontRepo) {
    polygon.CreatedAt = polygonAPI.CreatedAt;
    polygon.DeletedAt = polygonAPI.DeletedAt;
    polygon.ID = polygonAPI.ID;
    // insertion point for basic fields copy operations
    polygon.Name = polygonAPI.Name;
    polygon.Label = polygonAPI.Label;
    polygon.PopupHTML = polygonAPI.PopupHTML;
    polygon.FillColor = polygonAPI.FillColor;
    polygon.FillOpacity = polygonAPI.FillOpacity;
    polygon.StrokeColor = polygonAPI.StrokeColor;
    polygon.StrokeOpacity = polygonAPI.StrokeOpacity;
    polygon.StrokeWidth = polygonAPI.StrokeWidth;
    polygon.IsExtruded = polygonAPI.IsExtruded;
    polygon.BaseAltitude = polygonAPI.BaseAltitude;
    polygon.ExtrusionHeight = polygonAPI.ExtrusionHeight;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
    if (!Array.isArray(polygonAPI.PolygonPointersEncoding.Coordinates)) {
        console.error('Rects is not an array:', polygonAPI.PolygonPointersEncoding.Coordinates);
        return;
    }
    polygon.Coordinates = new Array();
    for (let _id of polygonAPI.PolygonPointersEncoding.Coordinates) {
        let _coordinate = frontRepo.map_ID_Coordinate.get(_id);
        if (_coordinate != undefined) {
            polygon.Coordinates.push(_coordinate);
        }
    }
}

// generated code - do not edit
class Polyline {
    static GONGSTRUCT_NAME = "Polyline";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    Label = "";
    PopupHTML = "";
    StrokeColor = "";
    StrokeOpacity = 0;
    StrokeWidth = 0;
    StrokeDashArray = "";
    // insertion point for pointers and slices of pointers declarations
    Coordinates = [];
    CreatedAt;
    DeletedAt;
}
function CopyPolylineToPolylineAPI(polyline, polylineAPI) {
    polylineAPI.CreatedAt = polyline.CreatedAt;
    polylineAPI.DeletedAt = polyline.DeletedAt;
    polylineAPI.ID = polyline.ID;
    // insertion point for basic fields copy operations
    polylineAPI.Name = polyline.Name;
    polylineAPI.Label = polyline.Label;
    polylineAPI.PopupHTML = polyline.PopupHTML;
    polylineAPI.StrokeColor = polyline.StrokeColor;
    polylineAPI.StrokeOpacity = polyline.StrokeOpacity;
    polylineAPI.StrokeWidth = polyline.StrokeWidth;
    polylineAPI.StrokeDashArray = polyline.StrokeDashArray;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
    polylineAPI.PolylinePointersEncoding.Coordinates = [];
    for (let _coordinate of polyline.Coordinates) {
        polylineAPI.PolylinePointersEncoding.Coordinates.push(_coordinate.ID);
    }
}
// CopyPolylineAPIToPolyline update basic, pointers and slice of pointers fields of polyline
// from respectively the basic fields and encoded fields of pointers and slices of pointers of polylineAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyPolylineAPIToPolyline(polylineAPI, polyline, frontRepo) {
    polyline.CreatedAt = polylineAPI.CreatedAt;
    polyline.DeletedAt = polylineAPI.DeletedAt;
    polyline.ID = polylineAPI.ID;
    // insertion point for basic fields copy operations
    polyline.Name = polylineAPI.Name;
    polyline.Label = polylineAPI.Label;
    polyline.PopupHTML = polylineAPI.PopupHTML;
    polyline.StrokeColor = polylineAPI.StrokeColor;
    polyline.StrokeOpacity = polylineAPI.StrokeOpacity;
    polyline.StrokeWidth = polylineAPI.StrokeWidth;
    polyline.StrokeDashArray = polylineAPI.StrokeDashArray;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
    if (!Array.isArray(polylineAPI.PolylinePointersEncoding.Coordinates)) {
        console.error('Rects is not an array:', polylineAPI.PolylinePointersEncoding.Coordinates);
        return;
    }
    polyline.Coordinates = new Array();
    for (let _id of polylineAPI.PolylinePointersEncoding.Coordinates) {
        let _coordinate = frontRepo.map_ID_Coordinate.get(_id);
        if (_coordinate != undefined) {
            polyline.Coordinates.push(_coordinate);
        }
    }
}

// generated code - do not edit
class BackRepoData {
    // insertion point for declarations
    CoordinateAPIs = new Array();
    LayerAPIs = new Array();
    MapViewAPIs = new Array();
    MarkerAPIs = new Array();
    PolygonAPIs = new Array();
    PolylineAPIs = new Array();
    // index of the web socket for this stack type (unique among all stack instances)
    GONG__Index;
    constructor(data) {
        // insertion point for copies
        this.CoordinateAPIs = data?.CoordinateAPIs || [];
        this.LayerAPIs = data?.LayerAPIs || [];
        this.MapViewAPIs = data?.MapViewAPIs || [];
        this.MarkerAPIs = data?.MarkerAPIs || [];
        this.PolygonAPIs = data?.PolygonAPIs || [];
        this.PolylineAPIs = data?.PolylineAPIs || [];
        this.GONG__Index = data?.GONG__Index ?? -1; // Assign Index here
    }
}

// generated code - do not edit
const StackType = "github.com/fullstack-lang/gong/lib/maplibre/go/models";
// FrontRepo stores all instances in a front repository (design pattern repository)
class FrontRepo {
    array_Coordinates = new Array(); // array of front instances
    map_ID_Coordinate = new Map(); // map of front instances
    array_Layers = new Array(); // array of front instances
    map_ID_Layer = new Map(); // map of front instances
    array_MapViews = new Array(); // array of front instances
    map_ID_MapView = new Map(); // map of front instances
    array_Markers = new Array(); // array of front instances
    map_ID_Marker = new Map(); // map of front instances
    array_Polygons = new Array(); // array of front instances
    map_ID_Polygon = new Map(); // map of front instances
    array_Polylines = new Array(); // array of front instances
    map_ID_Polyline = new Map(); // map of front instances
    GONG__Index = -1;
    // getFrontArray allows for a get function that is robust to refactoring of the named struct name
    // for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
    // contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
    getFrontArray(gongStructName) {
        switch (gongStructName) {
            // insertion point
            case 'Coordinate':
                return this.array_Coordinates;
            case 'Layer':
                return this.array_Layers;
            case 'MapView':
                return this.array_MapViews;
            case 'Marker':
                return this.array_Markers;
            case 'Polygon':
                return this.array_Polygons;
            case 'Polyline':
                return this.array_Polylines;
            default:
                throw new Error("Type not recognized");
        }
    }
    getFrontMap(gongStructName) {
        switch (gongStructName) {
            // insertion point
            case 'Coordinate':
                return this.map_ID_Coordinate;
            case 'Layer':
                return this.map_ID_Layer;
            case 'MapView':
                return this.map_ID_MapView;
            case 'Marker':
                return this.map_ID_Marker;
            case 'Polygon':
                return this.map_ID_Polygon;
            case 'Polyline':
                return this.map_ID_Polyline;
            default:
                throw new Error("Type not recognized");
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
class DialogData {
    ID = 0; // ID of the calling instance
    // the reverse pointer is the name of the generated field on the destination
    // struct of the ONE-MANY association
    ReversePointer = ""; // field of {{Structname}} that serve as reverse pointer
    OrderingMode = false; // if true, this is for ordering items
    // there are different selection mode : ONE_MANY or MANY_MANY
    SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE;
    // used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
    //
    // In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
    // 
    // in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
    // at the end of the ONE-MANY association
    SourceStruct = ""; // The "Aclass"
    SourceField = ""; // the "AnarrayofbUse"
    IntermediateStruct = ""; // the "AclassBclassUse" 
    IntermediateStructField = ""; // the "Bclass" as field
    NextAssociationStruct = ""; // the "Bclass"
    Name = "";
}
var SelectionMode;
(function (SelectionMode) {
    SelectionMode["ONE_MANY_ASSOCIATION_MODE"] = "ONE_MANY_ASSOCIATION_MODE";
    SelectionMode["MANY_MANY_ASSOCIATION_MODE"] = "MANY_MANY_ASSOCIATION_MODE";
})(SelectionMode || (SelectionMode = {}));
//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
class FrontRepoService {
    http;
    ngZone;
    Name = "";
    httpOptions = {
        headers: new HttpHeaders({ 'Content-Type': 'application/json' })
    };
    //
    // Store of all instances of the stack
    //
    frontRepo = new (FrontRepo);
    // Manage open WebSocket connections
    webSocketConnections = new Map();
    constructor(http, ngZone) {
        this.http = http;
        this.ngZone = ngZone;
    }
    //
    // pull returns the FrontRepo Observable
    //
    pull(Name = "") {
        this.Name = Name;
        return of(this.frontRepo);
    }
    connectToWebSocket(Name) {
        // console.log("github.com/fullstack-lang/gong/lib/maplibre/go; connectToWebSocket: started", Name)
        // Check if a connection for this name already exists
        if (this.webSocketConnections.has(Name)) {
            // console.log("github.com/fullstack-lang/gong/lib/maplibre/go; connectToWebSocket: returning existing connection")
            return this.webSocketConnections.get(Name);
        }
        //
        // Create a new connection
        //
        let host = window.location.host;
        const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
        if (host === 'localhost:4200') {
            host = 'localhost:8080';
        }
        // Construct the base path using the dynamic host and protocol
        // The API path remains the same.
        let basePath = `${protocol}//${host}/api/github.com/fullstack-lang/gong/lib/maplibre/go/v1/ws/stage`;
        let params = new HttpParams().set("Name", Name);
        let paramString = params.toString();
        let url = `${basePath}?${paramString}`;
        const newConnection$ = new Observable(observer => {
            // console.log("github.com/fullstack-lang/gong/lib/maplibre/go; connectToWebSocket: new Observable created")
            let socket;
            const isOfflineMode = window.location.protocol === 'file:' || window.document.getElementById('wasm-progress-container') !== null;
            const processData = (dataString) => {
                // console.log("github.com/fullstack-lang/gong/lib/maplibre/go; connectToWebSocket: processData called")
                const backRepoData = new BackRepoData(JSON.parse(dataString));
                let frontRepo = new (FrontRepo)();
                frontRepo.GONG__Index = backRepoData.GONG__Index;
                // 
                // First Step: init map of instances
                // insertion point sub template for init 
                // init the arrays
                frontRepo.array_Coordinates = [];
                frontRepo.map_ID_Coordinate.clear();
                backRepoData.CoordinateAPIs.forEach(coordinateAPI => {
                    let coordinate = new Coordinate;
                    frontRepo.array_Coordinates.push(coordinate);
                    frontRepo.map_ID_Coordinate.set(coordinateAPI.ID, coordinate);
                });
                // init the arrays
                frontRepo.array_Layers = [];
                frontRepo.map_ID_Layer.clear();
                backRepoData.LayerAPIs.forEach(layerAPI => {
                    let layer = new Layer;
                    frontRepo.array_Layers.push(layer);
                    frontRepo.map_ID_Layer.set(layerAPI.ID, layer);
                });
                // init the arrays
                frontRepo.array_MapViews = [];
                frontRepo.map_ID_MapView.clear();
                backRepoData.MapViewAPIs.forEach(mapviewAPI => {
                    let mapview = new MapView;
                    frontRepo.array_MapViews.push(mapview);
                    frontRepo.map_ID_MapView.set(mapviewAPI.ID, mapview);
                });
                // init the arrays
                frontRepo.array_Markers = [];
                frontRepo.map_ID_Marker.clear();
                backRepoData.MarkerAPIs.forEach(markerAPI => {
                    let marker = new Marker;
                    frontRepo.array_Markers.push(marker);
                    frontRepo.map_ID_Marker.set(markerAPI.ID, marker);
                });
                // init the arrays
                frontRepo.array_Polygons = [];
                frontRepo.map_ID_Polygon.clear();
                backRepoData.PolygonAPIs.forEach(polygonAPI => {
                    let polygon = new Polygon;
                    frontRepo.array_Polygons.push(polygon);
                    frontRepo.map_ID_Polygon.set(polygonAPI.ID, polygon);
                });
                // init the arrays
                frontRepo.array_Polylines = [];
                frontRepo.map_ID_Polyline.clear();
                backRepoData.PolylineAPIs.forEach(polylineAPI => {
                    let polyline = new Polyline;
                    frontRepo.array_Polylines.push(polyline);
                    frontRepo.map_ID_Polyline.set(polylineAPI.ID, polyline);
                });
                // 
                // Second Step: reddeem front objects
                // insertion point sub template for redeem 
                // fill up front objects
                backRepoData.CoordinateAPIs.forEach(coordinateAPI => {
                    let coordinate = frontRepo.map_ID_Coordinate.get(coordinateAPI.ID);
                    CopyCoordinateAPIToCoordinate(coordinateAPI, coordinate, frontRepo);
                });
                // fill up front objects
                backRepoData.LayerAPIs.forEach(layerAPI => {
                    let layer = frontRepo.map_ID_Layer.get(layerAPI.ID);
                    CopyLayerAPIToLayer(layerAPI, layer, frontRepo);
                });
                // fill up front objects
                backRepoData.MapViewAPIs.forEach(mapviewAPI => {
                    let mapview = frontRepo.map_ID_MapView.get(mapviewAPI.ID);
                    CopyMapViewAPIToMapView(mapviewAPI, mapview, frontRepo);
                });
                // fill up front objects
                backRepoData.MarkerAPIs.forEach(markerAPI => {
                    let marker = frontRepo.map_ID_Marker.get(markerAPI.ID);
                    CopyMarkerAPIToMarker(markerAPI, marker, frontRepo);
                });
                // fill up front objects
                backRepoData.PolygonAPIs.forEach(polygonAPI => {
                    let polygon = frontRepo.map_ID_Polygon.get(polygonAPI.ID);
                    CopyPolygonAPIToPolygon(polygonAPI, polygon, frontRepo);
                });
                // fill up front objects
                backRepoData.PolylineAPIs.forEach(polylineAPI => {
                    let polyline = frontRepo.map_ID_Polyline.get(polylineAPI.ID);
                    CopyPolylineAPIToPolyline(polylineAPI, polyline, frontRepo);
                });
                this.ngZone.run(() => {
                    observer.next(frontRepo);
                });
            };
            // 3. Connection Loop
            const attemptConnection = (retries) => {
                // console.log("github.com/fullstack-lang/gong/lib/maplibre/go; attemptConnection: retries =", retries, "isOfflineMode =", isOfflineMode)
                // A. WASM OFFLINE MODE (Check if Go is ready)
                if (window.openWasmSocket) {
                    // console.log("github.com/fullstack-lang/gong/lib/maplibre/go; attemptConnection: openWasmSocket exists, calling it");
                    window.openWasmSocket("github.com/fullstack-lang/gong/lib/maplibre/go", Name, processData);
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
                    socket = new WebSocket(url);
                    socket.onopen = (event) => {
                        // console.log("github.com/fullstack-lang/gong/lib/maplibre/go; WebSocket: onopen", event)
                    };
                    socket.onmessage = event => {
                        // console.log("github.com/fullstack-lang/gong/lib/maplibre/go; WebSocket: onmessage")
                        processData(event.data);
                    };
                    socket.onerror = event => {
                        console.error("github.com/fullstack-lang/gong/lib/maplibre/go WebSocket: onerror", event);
                        observer.error(event);
                    };
                    socket.onclose = (event) => {
                        // console.log("github.com/fullstack-lang/gong/lib/maplibre/go; WebSocket: onclose", event)
                        observer.complete();
                    };
                }
                else {
                    console.error("github.com/fullstack-lang/gong/lib/maplibre/go, attemptConnection: Offline mode detected, but WASM backend failed to load.");
                    observer.error("Offline mode detected, but WASM backend failed to load.");
                }
            };
            attemptConnection(50);
            // Teardown logic: Called when the last subscriber unsubscribes.
            return () => {
                this.webSocketConnections.delete(Name); // Remove from cache
                if (socket) {
                    socket.close();
                }
            };
        }).pipe(
        // This is the key:
        // - shareReplay makes this a "multicast" observable, sharing the single WebSocket among subscribers.
        // - { bufferSize: 1, refCount: true } means:
        //   - bufferSize: 1 => new subscribers get the last emitted value immediately.
        //   - refCount: true => the connection starts with the first subscriber and stops with the last.
        shareReplay({ bufferSize: 1, refCount: true }));
        // Store the new connection observable in the map
        this.webSocketConnections.set(Name, newConnection$);
        return newConnection$;
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: FrontRepoService, deps: [{ token: i1.HttpClient }, { token: i0.NgZone }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: FrontRepoService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: FrontRepoService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: i0.NgZone }] });
// insertion point for get unique ID per struct 
function getCoordinateUniqueID(id) {
    return 31 * id;
}
function getLayerUniqueID(id) {
    return 37 * id;
}
function getMapViewUniqueID(id) {
    return 41 * id;
}
function getMarkerUniqueID(id) {
    return 43 * id;
}
function getPolygonUniqueID(id) {
    return 47 * id;
}
function getPolylineUniqueID(id) {
    return 53 * id;
}

// define the type of nullable Int64 in order to support back pointers IDs
class NullInt64 {
    Int64 = 0;
    Valid = false;
}

// generated file, do not edit
class CommitNbFromBackService {
    http;
    location;
    document;
    httpOptions = {
        headers: new HttpHeaders({ 'Content-Type': 'application/json' })
    };
    commitNbFromBackUrl;
    constructor(http, location, document) {
        this.http = http;
        this.location = location;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.commitNbFromBackUrl = origin + '/api/github.com/fullstack-lang/gong/lib/maplibre/go/v1/commitfrombacknb';
    }
    getCommitNbFromBack(intervalMs, Name = "") {
        let params = new HttpParams().set("Name", Name);
        return interval(intervalMs).pipe(switchMap(() => this.http.get(this.commitNbFromBackUrl, { params: params }).pipe(catchError(error => {
            // Handle the error here, e.g. log it, show a notification, etc.
            console.error('Error fetching commit number:', error);
            // Return a default value, a new Observable, or rethrow the error
            return of(0); // Here, we return 0 as a default value
        }))));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in CommitNbFromBackService', result) {
        return (error) => {
            // TODO: send the error to remote logging infrastructure
            console.error("in CommitNbFromBackService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log('${operation} failed: ${error.message}');
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: CommitNbFromBackService, deps: [{ token: i1.HttpClient }, { token: i2.Location }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: CommitNbFromBackService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: CommitNbFromBackService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: i2.Location }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// generated file, do not edit
class PushFromFrontNbService {
    http;
    location;
    document;
    httpOptions = {
        headers: new HttpHeaders({ 'Content-Type': 'application/json' })
    };
    pushFromFrontNbURL;
    constructor(http, location, document) {
        this.http = http;
        this.location = location;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.pushFromFrontNbURL = origin + '/api/github.com/fullstack-lang/gong/lib/maplibre/go/v1/pushfromfrontnb';
    }
    // observable of the commit nb getter
    getPushFromFrontNb() {
        return this.http.get(this.pushFromFrontNbURL)
            .pipe(tap(_ => this.log('fetched commit nb')), catchError(this.handleError('getPushFromFrontNb', -1)));
    }
    getPushNbFromFront(intervalMs, Name = "") {
        let params = new HttpParams().set("Name", Name);
        return interval(intervalMs).pipe(switchMap(() => this.http.get(this.pushFromFrontNbURL, { params: params }).pipe(catchError(error => {
            // Handle the error here, e.g. log it, show a notification, etc.
            console.error('Error fetching commit number:', error);
            // Return a default value, a new Observable, or rethrow the error
            return of(0); // Here, we return 0 as a default value
        }))));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation', result) {
        return (error) => {
            // TODO: send the error to remote logging infrastructure
            console.error(error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log('${operation} failed: ${error.message}');
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: PushFromFrontNbService, deps: [{ token: i1.HttpClient }, { token: i2.Location }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: PushFromFrontNbService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: PushFromFrontNbService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: i2.Location }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// generated code - do not edit
class WebSocketService {
    document;
    constructor(document) {
        this.document = document;
    }
    connect(stackPath) {
        return new Observable(subscriber => {
            // Determine if we are running in the offline HTML file
            const isOfflineMode = this.document.location.protocol === 'file:' || window.document.getElementById('wasm-progress-container') !== null;
            const attemptConnection = (retries) => {
                // 1. WASM OFFLINE MODE (Check if Go is ready)
                if (window.openWasmSocket) {
                    const onMessageFromGo = (message) => {
                        subscriber.next(JSON.parse(message));
                    };
                    window.openWasmSocket(stackPath, onMessageFromGo);
                    return; // Successfully connected to WASM
                }
                // 2. STILL WAITING FOR WASM
                if (isOfflineMode && retries > 0) {
                    // If we are offline, WASM might just be decoding. Wait 100ms and check again.
                    setTimeout(() => attemptConnection(retries - 1), 100);
                    return;
                }
                // 3. STANDARD SERVER MODE (Fallback)
                if (!isOfflineMode) {
                    let protocol = this.document.location.protocol === 'https:' ? 'wss://' : 'ws://';
                    let port = this.document.location.port ? ':' + this.document.location.port : '';
                    let host = this.document.location.hostname;
                    const wsUrl = `${protocol}${host}${port}/api/github.com/fullstack-lang/gong/lib/maplibre/go/v1/ws/stage?Name=${stackPath}`;
                    const ws = new WebSocket(wsUrl);
                    ws.onmessage = (event) => subscriber.next(JSON.parse(event.data));
                    ws.onerror = (error) => subscriber.error(error);
                    ws.onclose = () => subscriber.complete();
                }
                else {
                    subscriber.error("Offline mode detected, but WASM backend failed to load.");
                }
            };
            // Start the connection loop, allowing up to 5 seconds (50 * 100ms) for WASM to boot
            attemptConnection(50);
        });
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: WebSocketService, deps: [{ token: DOCUMENT$1 }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: WebSocketService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: WebSocketService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT$1]
                }] }] });

// generated code - do not edit
// insertion point for imports
class CoordinateAPI {
    static GONGSTRUCT_NAME = "Coordinate";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    Longitude = 0;
    Latitude = 0;
    Altitude = 0;
    // insertion point for other decls
    CoordinatePointersEncoding = new CoordinatePointersEncoding;
    CreatedAt;
    DeletedAt;
}
class CoordinatePointersEncoding {
}

// generated code, do not edit
// generated by ng_file_service_ts
// insertion point for imports
class CoordinateService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    CoordinateServiceChanged = new BehaviorSubject("");
    coordinatesUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.coordinatesUrl = origin + '/api/github.com/fullstack-lang/gong/lib/maplibre/go/v1/coordinates';
    }
    // updateFront copy coordinate to a version with encoded pointers and update to the back
    updateFront(coordinate, Name) {
        let coordinateAPI = new CoordinateAPI;
        CopyCoordinateToCoordinateAPI(coordinate, coordinateAPI);
        const id = typeof coordinateAPI === 'number' ? coordinateAPI : coordinateAPI.ID;
        const url = `${this.coordinatesUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, coordinateAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateCoordinate')));
    }
    /** PUT: update the coordinatedb on the server */
    update(coordinatedb, Name, frontRepo) {
        return this.updateCoordinate(coordinatedb, Name, frontRepo);
    }
    updateCoordinate(coordinatedb, Name, frontRepo) {
        const id = typeof coordinatedb === 'number' ? coordinatedb : coordinatedb.ID;
        const url = `${this.coordinatesUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, coordinatedb, httpOptions).pipe(tap(_ => {
            // this.log(`updated coordinatedb id=${coordinatedb.ID}`)
        }), catchError(this.handleError('updateCoordinate')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in CoordinateService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("CoordinateService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: CoordinateService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: CoordinateService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: CoordinateService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

class LayerAPI {
    static GONGSTRUCT_NAME = "Layer";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    IsVisible = false;
    // insertion point for other decls
    LayerPointersEncoding = new LayerPointersEncoding;
    CreatedAt;
    DeletedAt;
}
class LayerPointersEncoding {
    // insertion point for pointers and slices of pointers encoding fields
    Polygons = [];
    Polylines = [];
    Markers = [];
}

// generated code, do not edit
// generated by ng_file_service_ts
class LayerService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    LayerServiceChanged = new BehaviorSubject("");
    layersUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.layersUrl = origin + '/api/github.com/fullstack-lang/gong/lib/maplibre/go/v1/layers';
    }
    // updateFront copy layer to a version with encoded pointers and update to the back
    updateFront(layer, Name) {
        let layerAPI = new LayerAPI;
        CopyLayerToLayerAPI(layer, layerAPI);
        const id = typeof layerAPI === 'number' ? layerAPI : layerAPI.ID;
        const url = `${this.layersUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, layerAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateLayer')));
    }
    /** PUT: update the layerdb on the server */
    update(layerdb, Name, frontRepo) {
        return this.updateLayer(layerdb, Name, frontRepo);
    }
    updateLayer(layerdb, Name, frontRepo) {
        const id = typeof layerdb === 'number' ? layerdb : layerdb.ID;
        const url = `${this.layersUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, layerdb, httpOptions).pipe(tap(_ => {
            // this.log(`updated layerdb id=${layerdb.ID}`)
        }), catchError(this.handleError('updateLayer')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in LayerService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("LayerService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: LayerService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: LayerService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: LayerService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

class MapViewAPI {
    static GONGSTRUCT_NAME = "MapView";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StyleURL = "";
    Longitude = 0;
    Latitude = 0;
    Zoom = 0;
    Pitch = 0;
    Bearing = 0;
    // insertion point for other decls
    MapViewPointersEncoding = new MapViewPointersEncoding;
    CreatedAt;
    DeletedAt;
}
class MapViewPointersEncoding {
    // insertion point for pointers and slices of pointers encoding fields
    Layers = [];
}

// generated code, do not edit
// generated by ng_file_service_ts
class MapViewService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    MapViewServiceChanged = new BehaviorSubject("");
    mapviewsUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.mapviewsUrl = origin + '/api/github.com/fullstack-lang/gong/lib/maplibre/go/v1/mapviews';
    }
    // updateFront copy mapview to a version with encoded pointers and update to the back
    updateFront(mapview, Name) {
        let mapviewAPI = new MapViewAPI;
        CopyMapViewToMapViewAPI(mapview, mapviewAPI);
        const id = typeof mapviewAPI === 'number' ? mapviewAPI : mapviewAPI.ID;
        const url = `${this.mapviewsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, mapviewAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateMapView')));
    }
    /** PUT: update the mapviewdb on the server */
    update(mapviewdb, Name, frontRepo) {
        return this.updateMapView(mapviewdb, Name, frontRepo);
    }
    updateMapView(mapviewdb, Name, frontRepo) {
        const id = typeof mapviewdb === 'number' ? mapviewdb : mapviewdb.ID;
        const url = `${this.mapviewsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, mapviewdb, httpOptions).pipe(tap(_ => {
            // this.log(`updated mapviewdb id=${mapviewdb.ID}`)
        }), catchError(this.handleError('updateMapView')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in MapViewService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("MapViewService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: MapViewService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: MapViewService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: MapViewService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// generated code - do not edit
// insertion point for imports
class MarkerAPI {
    static GONGSTRUCT_NAME = "Marker";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    Label = "";
    PopupHTML = "";
    Longitude = 0;
    Latitude = 0;
    Altitude = 0;
    SVG = "";
    Color = "";
    Scale = 0;
    Rotation = 0;
    RotationAlignment = "";
    PitchAlignment = "";
    Anchor = "";
    IsDraggable = false;
    // insertion point for other decls
    MarkerPointersEncoding = new MarkerPointersEncoding;
    CreatedAt;
    DeletedAt;
}
class MarkerPointersEncoding {
}

// generated code, do not edit
// generated by ng_file_service_ts
// insertion point for imports
class MarkerService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    MarkerServiceChanged = new BehaviorSubject("");
    markersUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.markersUrl = origin + '/api/github.com/fullstack-lang/gong/lib/maplibre/go/v1/markers';
    }
    // updateFront copy marker to a version with encoded pointers and update to the back
    updateFront(marker, Name) {
        let markerAPI = new MarkerAPI;
        CopyMarkerToMarkerAPI(marker, markerAPI);
        const id = typeof markerAPI === 'number' ? markerAPI : markerAPI.ID;
        const url = `${this.markersUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, markerAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateMarker')));
    }
    /** PUT: update the markerdb on the server */
    update(markerdb, Name, frontRepo) {
        return this.updateMarker(markerdb, Name, frontRepo);
    }
    updateMarker(markerdb, Name, frontRepo) {
        const id = typeof markerdb === 'number' ? markerdb : markerdb.ID;
        const url = `${this.markersUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, markerdb, httpOptions).pipe(tap(_ => {
            // this.log(`updated markerdb id=${markerdb.ID}`)
        }), catchError(this.handleError('updateMarker')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in MarkerService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("MarkerService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: MarkerService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: MarkerService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: MarkerService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

class PolygonAPI {
    static GONGSTRUCT_NAME = "Polygon";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    Label = "";
    PopupHTML = "";
    FillColor = "";
    FillOpacity = 0;
    StrokeColor = "";
    StrokeOpacity = 0;
    StrokeWidth = 0;
    IsExtruded = false;
    BaseAltitude = 0;
    ExtrusionHeight = 0;
    // insertion point for other decls
    PolygonPointersEncoding = new PolygonPointersEncoding;
    CreatedAt;
    DeletedAt;
}
class PolygonPointersEncoding {
    // insertion point for pointers and slices of pointers encoding fields
    Coordinates = [];
}

// generated code, do not edit
// generated by ng_file_service_ts
class PolygonService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    PolygonServiceChanged = new BehaviorSubject("");
    polygonsUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.polygonsUrl = origin + '/api/github.com/fullstack-lang/gong/lib/maplibre/go/v1/polygons';
    }
    // updateFront copy polygon to a version with encoded pointers and update to the back
    updateFront(polygon, Name) {
        let polygonAPI = new PolygonAPI;
        CopyPolygonToPolygonAPI(polygon, polygonAPI);
        const id = typeof polygonAPI === 'number' ? polygonAPI : polygonAPI.ID;
        const url = `${this.polygonsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, polygonAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updatePolygon')));
    }
    /** PUT: update the polygondb on the server */
    update(polygondb, Name, frontRepo) {
        return this.updatePolygon(polygondb, Name, frontRepo);
    }
    updatePolygon(polygondb, Name, frontRepo) {
        const id = typeof polygondb === 'number' ? polygondb : polygondb.ID;
        const url = `${this.polygonsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, polygondb, httpOptions).pipe(tap(_ => {
            // this.log(`updated polygondb id=${polygondb.ID}`)
        }), catchError(this.handleError('updatePolygon')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in PolygonService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("PolygonService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: PolygonService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: PolygonService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: PolygonService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

class PolylineAPI {
    static GONGSTRUCT_NAME = "Polyline";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    Label = "";
    PopupHTML = "";
    StrokeColor = "";
    StrokeOpacity = 0;
    StrokeWidth = 0;
    StrokeDashArray = "";
    // insertion point for other decls
    PolylinePointersEncoding = new PolylinePointersEncoding;
    CreatedAt;
    DeletedAt;
}
class PolylinePointersEncoding {
    // insertion point for pointers and slices of pointers encoding fields
    Coordinates = [];
}

// generated code, do not edit
// generated by ng_file_service_ts
class PolylineService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    PolylineServiceChanged = new BehaviorSubject("");
    polylinesUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.polylinesUrl = origin + '/api/github.com/fullstack-lang/gong/lib/maplibre/go/v1/polylines';
    }
    // updateFront copy polyline to a version with encoded pointers and update to the back
    updateFront(polyline, Name) {
        let polylineAPI = new PolylineAPI;
        CopyPolylineToPolylineAPI(polyline, polylineAPI);
        const id = typeof polylineAPI === 'number' ? polylineAPI : polylineAPI.ID;
        const url = `${this.polylinesUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, polylineAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updatePolyline')));
    }
    /** PUT: update the polylinedb on the server */
    update(polylinedb, Name, frontRepo) {
        return this.updatePolyline(polylinedb, Name, frontRepo);
    }
    updatePolyline(polylinedb, Name, frontRepo) {
        const id = typeof polylinedb === 'number' ? polylinedb : polylinedb.ID;
        const url = `${this.polylinesUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, polylinedb, httpOptions).pipe(tap(_ => {
            // this.log(`updated polylinedb id=${polylinedb.ID}`)
        }), catchError(this.handleError('updatePolyline')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in PolylineService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("PolylineService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: PolylineService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: PolylineService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "22.1.4", ngImport: i0, type: PolylineService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// generated from ng_file_enum.ts.go
var MarkerAnchorEnum;
(function (MarkerAnchorEnum) {
    // insertion point	
    MarkerAnchorEnum["MarkerAnchorCenter"] = "center";
    MarkerAnchorEnum["MarkerAnchorTop"] = "top";
    MarkerAnchorEnum["MarkerAnchorBottom"] = "bottom";
    MarkerAnchorEnum["MarkerAnchorLeft"] = "left";
    MarkerAnchorEnum["MarkerAnchorRight"] = "right";
    MarkerAnchorEnum["MarkerAnchorTopLeft"] = "top-left";
    MarkerAnchorEnum["MarkerAnchorTopRight"] = "top-right";
    MarkerAnchorEnum["MarkerAnchorBottomLeft"] = "bottom-left";
    MarkerAnchorEnum["MarkerAnchorBottomRight"] = "bottom-right";
})(MarkerAnchorEnum || (MarkerAnchorEnum = {}));
const MarkerAnchorEnumList = [
    { value: MarkerAnchorEnum.MarkerAnchorCenter, viewValue: "center" },
    { value: MarkerAnchorEnum.MarkerAnchorTop, viewValue: "top" },
    { value: MarkerAnchorEnum.MarkerAnchorBottom, viewValue: "bottom" },
    { value: MarkerAnchorEnum.MarkerAnchorLeft, viewValue: "left" },
    { value: MarkerAnchorEnum.MarkerAnchorRight, viewValue: "right" },
    { value: MarkerAnchorEnum.MarkerAnchorTopLeft, viewValue: "top-left" },
    { value: MarkerAnchorEnum.MarkerAnchorTopRight, viewValue: "top-right" },
    { value: MarkerAnchorEnum.MarkerAnchorBottomLeft, viewValue: "bottom-left" },
    { value: MarkerAnchorEnum.MarkerAnchorBottomRight, viewValue: "bottom-right" },
];

// generated from ng_file_enum.ts.go
var MarkerPitchAlignmentEnum;
(function (MarkerPitchAlignmentEnum) {
    // insertion point	
    MarkerPitchAlignmentEnum["MarkerAlignmentAuto"] = "auto";
    MarkerPitchAlignmentEnum["MarkerAlignmentMap"] = "map";
    MarkerPitchAlignmentEnum["MarkerAlignmentViewport"] = "viewport";
})(MarkerPitchAlignmentEnum || (MarkerPitchAlignmentEnum = {}));
const MarkerPitchAlignmentEnumList = [
    { value: MarkerPitchAlignmentEnum.MarkerAlignmentAuto, viewValue: "auto" },
    { value: MarkerPitchAlignmentEnum.MarkerAlignmentMap, viewValue: "map" },
    { value: MarkerPitchAlignmentEnum.MarkerAlignmentViewport, viewValue: "viewport" },
];

// generated from ng_file_public_api_ts.go
/*
* Public API Surface of maplibre
*/

/**
 * Generated bundle index. Do not edit.
 */

export { BackRepoData, CommitNbFromBackService, Coordinate, CoordinateAPI, CoordinatePointersEncoding, CoordinateService, CopyCoordinateAPIToCoordinate, CopyCoordinateToCoordinateAPI, CopyLayerAPIToLayer, CopyLayerToLayerAPI, CopyMapViewAPIToMapView, CopyMapViewToMapViewAPI, CopyMarkerAPIToMarker, CopyMarkerToMarkerAPI, CopyPolygonAPIToPolygon, CopyPolygonToPolygonAPI, CopyPolylineAPIToPolyline, CopyPolylineToPolylineAPI, DialogData, FrontRepo, FrontRepoService, Layer, LayerAPI, LayerPointersEncoding, LayerService, MapView, MapViewAPI, MapViewPointersEncoding, MapViewService, MaplibreModule, Marker, MarkerAPI, MarkerAnchorEnum, MarkerAnchorEnumList, MarkerPitchAlignmentEnum, MarkerPitchAlignmentEnumList, MarkerPointersEncoding, MarkerService, NullInt64, Polygon, PolygonAPI, PolygonPointersEncoding, PolygonService, Polyline, PolylineAPI, PolylinePointersEncoding, PolylineService, PushFromFrontNbService, SelectionMode, StackType, WebSocketService, getCoordinateUniqueID, getLayerUniqueID, getMapViewUniqueID, getMarkerUniqueID, getPolygonUniqueID, getPolylineUniqueID };
//# sourceMappingURL=maplibre.mjs.map
