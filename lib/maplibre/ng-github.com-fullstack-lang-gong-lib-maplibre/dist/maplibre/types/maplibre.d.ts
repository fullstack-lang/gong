import * as i0 from '@angular/core';
import { NgZone } from '@angular/core';
import { HttpHeaders, HttpClient } from '@angular/common/http';
import { Observable, BehaviorSubject } from 'rxjs';
import { Location } from '@angular/common';

declare class MaplibreModule {
    static ɵfac: i0.ɵɵFactoryDeclaration<MaplibreModule, never>;
    static ɵmod: i0.ɵɵNgModuleDeclaration<MaplibreModule, never, never, never>;
    static ɵinj: i0.ɵɵInjectorDeclaration<MaplibreModule>;
}

declare class CoordinateAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    Longitude: number;
    Latitude: number;
    Altitude: number;
    CoordinatePointersEncoding: CoordinatePointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class CoordinatePointersEncoding {
}

declare class Coordinate {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    Longitude: number;
    Latitude: number;
    Altitude: number;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyCoordinateToCoordinateAPI(coordinate: Coordinate, coordinateAPI: CoordinateAPI): void;
declare function CopyCoordinateAPIToCoordinate(coordinateAPI: CoordinateAPI, coordinate: Coordinate, frontRepo: FrontRepo): void;

declare class LayerAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    IsVisible: boolean;
    LayerPointersEncoding: LayerPointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class LayerPointersEncoding {
    Polygons: number[];
    Polylines: number[];
    Markers: number[];
}

declare class PolygonAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    Label: string;
    PopupHTML: string;
    FillColor: string;
    FillOpacity: number;
    StrokeColor: string;
    StrokeOpacity: number;
    StrokeWidth: number;
    IsExtruded: boolean;
    BaseAltitude: number;
    ExtrusionHeight: number;
    PolygonPointersEncoding: PolygonPointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class PolygonPointersEncoding {
    Coordinates: number[];
}

declare class Polygon {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    Label: string;
    PopupHTML: string;
    FillColor: string;
    FillOpacity: number;
    StrokeColor: string;
    StrokeOpacity: number;
    StrokeWidth: number;
    IsExtruded: boolean;
    BaseAltitude: number;
    ExtrusionHeight: number;
    Coordinates: Array<Coordinate>;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyPolygonToPolygonAPI(polygon: Polygon, polygonAPI: PolygonAPI): void;
declare function CopyPolygonAPIToPolygon(polygonAPI: PolygonAPI, polygon: Polygon, frontRepo: FrontRepo): void;

declare class PolylineAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    Label: string;
    PopupHTML: string;
    StrokeColor: string;
    StrokeOpacity: number;
    StrokeWidth: number;
    StrokeDashArray: string;
    PolylinePointersEncoding: PolylinePointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class PolylinePointersEncoding {
    Coordinates: number[];
}

declare class Polyline {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    Label: string;
    PopupHTML: string;
    StrokeColor: string;
    StrokeOpacity: number;
    StrokeWidth: number;
    StrokeDashArray: string;
    Coordinates: Array<Coordinate>;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyPolylineToPolylineAPI(polyline: Polyline, polylineAPI: PolylineAPI): void;
declare function CopyPolylineAPIToPolyline(polylineAPI: PolylineAPI, polyline: Polyline, frontRepo: FrontRepo): void;

declare class MarkerAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    Label: string;
    PopupHTML: string;
    Longitude: number;
    Latitude: number;
    Altitude: number;
    SVG: string;
    Color: string;
    Scale: number;
    Rotation: number;
    RotationAlignment: string;
    PitchAlignment: string;
    Anchor: string;
    IsDraggable: boolean;
    MarkerPointersEncoding: MarkerPointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class MarkerPointersEncoding {
}

declare class Marker {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    Label: string;
    PopupHTML: string;
    Longitude: number;
    Latitude: number;
    Altitude: number;
    SVG: string;
    Color: string;
    Scale: number;
    Rotation: number;
    RotationAlignment: string;
    PitchAlignment: string;
    Anchor: string;
    IsDraggable: boolean;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyMarkerToMarkerAPI(marker: Marker, markerAPI: MarkerAPI): void;
declare function CopyMarkerAPIToMarker(markerAPI: MarkerAPI, marker: Marker, frontRepo: FrontRepo): void;

declare class Layer {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    IsVisible: boolean;
    Polygons: Array<Polygon>;
    Polylines: Array<Polyline>;
    Markers: Array<Marker>;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyLayerToLayerAPI(layer: Layer, layerAPI: LayerAPI): void;
declare function CopyLayerAPIToLayer(layerAPI: LayerAPI, layer: Layer, frontRepo: FrontRepo): void;

declare class MapViewAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StyleURL: string;
    Longitude: number;
    Latitude: number;
    Zoom: number;
    Pitch: number;
    Bearing: number;
    MapViewPointersEncoding: MapViewPointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class MapViewPointersEncoding {
    Layers: number[];
}

declare class MapView {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StyleURL: string;
    Longitude: number;
    Latitude: number;
    Zoom: number;
    Pitch: number;
    Bearing: number;
    Layers: Array<Layer>;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyMapViewToMapViewAPI(mapview: MapView, mapviewAPI: MapViewAPI): void;
declare function CopyMapViewAPIToMapView(mapviewAPI: MapViewAPI, mapview: MapView, frontRepo: FrontRepo): void;

declare const StackType = "github.com/fullstack-lang/gong/lib/maplibre/go/models";
declare class FrontRepo {
    array_Coordinates: Coordinate[];
    map_ID_Coordinate: Map<number, Coordinate>;
    array_Layers: Layer[];
    map_ID_Layer: Map<number, Layer>;
    array_MapViews: MapView[];
    map_ID_MapView: Map<number, MapView>;
    array_Markers: Marker[];
    map_ID_Marker: Map<number, Marker>;
    array_Polygons: Polygon[];
    map_ID_Polygon: Map<number, Polygon>;
    array_Polylines: Polyline[];
    map_ID_Polyline: Map<number, Polyline>;
    GONG__Index: number;
    getFrontArray<Type>(gongStructName: string): Array<Type>;
    getFrontMap<Type>(gongStructName: string): Map<number, Type>;
}
declare class DialogData {
    ID: number;
    ReversePointer: string;
    OrderingMode: boolean;
    SelectionMode: SelectionMode;
    SourceStruct: string;
    SourceField: string;
    IntermediateStruct: string;
    IntermediateStructField: string;
    NextAssociationStruct: string;
    Name: string;
}
declare enum SelectionMode {
    ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
    MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE"
}
declare class FrontRepoService {
    private http;
    private ngZone;
    Name: string;
    httpOptions: {
        headers: HttpHeaders;
    };
    frontRepo: FrontRepo;
    private webSocketConnections;
    constructor(http: HttpClient, ngZone: NgZone);
    pull(Name?: string): Observable<FrontRepo>;
    connectToWebSocket(Name: string): Observable<FrontRepo>;
    static ɵfac: i0.ɵɵFactoryDeclaration<FrontRepoService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<FrontRepoService>;
}
declare function getCoordinateUniqueID(id: number): number;
declare function getLayerUniqueID(id: number): number;
declare function getMapViewUniqueID(id: number): number;
declare function getMarkerUniqueID(id: number): number;
declare function getPolygonUniqueID(id: number): number;
declare function getPolylineUniqueID(id: number): number;

declare class NullInt64 {
    Int64: number;
    Valid: boolean;
}

declare class CommitNbFromBackService {
    private http;
    private location;
    private document;
    httpOptions: {
        headers: HttpHeaders;
    };
    private commitNbFromBackUrl;
    constructor(http: HttpClient, location: Location, document: Document);
    getCommitNbFromBack(intervalMs: number, Name?: string): Observable<number>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<CommitNbFromBackService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<CommitNbFromBackService>;
}

declare class PushFromFrontNbService {
    private http;
    private location;
    private document;
    httpOptions: {
        headers: HttpHeaders;
    };
    private pushFromFrontNbURL;
    constructor(http: HttpClient, location: Location, document: Document);
    getPushFromFrontNb(): Observable<number>;
    getPushNbFromFront(intervalMs: number, Name?: string): Observable<number>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<PushFromFrontNbService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<PushFromFrontNbService>;
}

declare class BackRepoData {
    CoordinateAPIs: CoordinateAPI[];
    LayerAPIs: LayerAPI[];
    MapViewAPIs: MapViewAPI[];
    MarkerAPIs: MarkerAPI[];
    PolygonAPIs: PolygonAPI[];
    PolylineAPIs: PolylineAPI[];
    GONG__Index: number;
    constructor(data?: Partial<BackRepoData>);
}

declare class WebSocketService {
    private document;
    constructor(document: Document);
    connect(stackPath: string): Observable<any>;
    static ɵfac: i0.ɵɵFactoryDeclaration<WebSocketService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<WebSocketService>;
}

declare class CoordinateService {
    private http;
    private document;
    CoordinateServiceChanged: BehaviorSubject<string>;
    private coordinatesUrl;
    constructor(http: HttpClient, document: Document);
    updateFront(coordinate: Coordinate, Name: string): Observable<CoordinateAPI>;
    /** PUT: update the coordinatedb on the server */
    update(coordinatedb: CoordinateAPI, Name: string, frontRepo: FrontRepo): Observable<CoordinateAPI>;
    updateCoordinate(coordinatedb: CoordinateAPI, Name: string, frontRepo: FrontRepo): Observable<CoordinateAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<CoordinateService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<CoordinateService>;
}

declare class LayerService {
    private http;
    private document;
    LayerServiceChanged: BehaviorSubject<string>;
    private layersUrl;
    constructor(http: HttpClient, document: Document);
    updateFront(layer: Layer, Name: string): Observable<LayerAPI>;
    /** PUT: update the layerdb on the server */
    update(layerdb: LayerAPI, Name: string, frontRepo: FrontRepo): Observable<LayerAPI>;
    updateLayer(layerdb: LayerAPI, Name: string, frontRepo: FrontRepo): Observable<LayerAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<LayerService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<LayerService>;
}

declare class MapViewService {
    private http;
    private document;
    MapViewServiceChanged: BehaviorSubject<string>;
    private mapviewsUrl;
    constructor(http: HttpClient, document: Document);
    updateFront(mapview: MapView, Name: string): Observable<MapViewAPI>;
    /** PUT: update the mapviewdb on the server */
    update(mapviewdb: MapViewAPI, Name: string, frontRepo: FrontRepo): Observable<MapViewAPI>;
    updateMapView(mapviewdb: MapViewAPI, Name: string, frontRepo: FrontRepo): Observable<MapViewAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<MapViewService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<MapViewService>;
}

declare class MarkerService {
    private http;
    private document;
    MarkerServiceChanged: BehaviorSubject<string>;
    private markersUrl;
    constructor(http: HttpClient, document: Document);
    updateFront(marker: Marker, Name: string): Observable<MarkerAPI>;
    /** PUT: update the markerdb on the server */
    update(markerdb: MarkerAPI, Name: string, frontRepo: FrontRepo): Observable<MarkerAPI>;
    updateMarker(markerdb: MarkerAPI, Name: string, frontRepo: FrontRepo): Observable<MarkerAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<MarkerService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<MarkerService>;
}

declare class PolygonService {
    private http;
    private document;
    PolygonServiceChanged: BehaviorSubject<string>;
    private polygonsUrl;
    constructor(http: HttpClient, document: Document);
    updateFront(polygon: Polygon, Name: string): Observable<PolygonAPI>;
    /** PUT: update the polygondb on the server */
    update(polygondb: PolygonAPI, Name: string, frontRepo: FrontRepo): Observable<PolygonAPI>;
    updatePolygon(polygondb: PolygonAPI, Name: string, frontRepo: FrontRepo): Observable<PolygonAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<PolygonService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<PolygonService>;
}

declare class PolylineService {
    private http;
    private document;
    PolylineServiceChanged: BehaviorSubject<string>;
    private polylinesUrl;
    constructor(http: HttpClient, document: Document);
    updateFront(polyline: Polyline, Name: string): Observable<PolylineAPI>;
    /** PUT: update the polylinedb on the server */
    update(polylinedb: PolylineAPI, Name: string, frontRepo: FrontRepo): Observable<PolylineAPI>;
    updatePolyline(polylinedb: PolylineAPI, Name: string, frontRepo: FrontRepo): Observable<PolylineAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<PolylineService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<PolylineService>;
}

declare enum MarkerAnchorEnum {
    MarkerAnchorCenter = "center",
    MarkerAnchorTop = "top",
    MarkerAnchorBottom = "bottom",
    MarkerAnchorLeft = "left",
    MarkerAnchorRight = "right",
    MarkerAnchorTopLeft = "top-left",
    MarkerAnchorTopRight = "top-right",
    MarkerAnchorBottomLeft = "bottom-left",
    MarkerAnchorBottomRight = "bottom-right"
}
interface MarkerAnchorEnumSelect {
    value: string;
    viewValue: string;
}
declare const MarkerAnchorEnumList: MarkerAnchorEnumSelect[];

declare enum MarkerPitchAlignmentEnum {
    MarkerAlignmentAuto = "auto",
    MarkerAlignmentMap = "map",
    MarkerAlignmentViewport = "viewport"
}
interface MarkerPitchAlignmentEnumSelect {
    value: string;
    viewValue: string;
}
declare const MarkerPitchAlignmentEnumList: MarkerPitchAlignmentEnumSelect[];

export { BackRepoData, CommitNbFromBackService, Coordinate, CoordinateAPI, CoordinatePointersEncoding, CoordinateService, CopyCoordinateAPIToCoordinate, CopyCoordinateToCoordinateAPI, CopyLayerAPIToLayer, CopyLayerToLayerAPI, CopyMapViewAPIToMapView, CopyMapViewToMapViewAPI, CopyMarkerAPIToMarker, CopyMarkerToMarkerAPI, CopyPolygonAPIToPolygon, CopyPolygonToPolygonAPI, CopyPolylineAPIToPolyline, CopyPolylineToPolylineAPI, DialogData, FrontRepo, FrontRepoService, Layer, LayerAPI, LayerPointersEncoding, LayerService, MapView, MapViewAPI, MapViewPointersEncoding, MapViewService, MaplibreModule, Marker, MarkerAPI, MarkerAnchorEnum, MarkerAnchorEnumList, MarkerPitchAlignmentEnum, MarkerPitchAlignmentEnumList, MarkerPointersEncoding, MarkerService, NullInt64, Polygon, PolygonAPI, PolygonPointersEncoding, PolygonService, Polyline, PolylineAPI, PolylinePointersEncoding, PolylineService, PushFromFrontNbService, SelectionMode, StackType, WebSocketService, getCoordinateUniqueID, getLayerUniqueID, getMapViewUniqueID, getMarkerUniqueID, getPolygonUniqueID, getPolylineUniqueID };
export type { MarkerAnchorEnumSelect, MarkerPitchAlignmentEnumSelect };
