import {
  Component,
  ElementRef,
  Input,
  OnDestroy,
  OnInit,
  AfterViewInit,
  ViewChild,
  ViewEncapsulation,
  ChangeDetectorRef,
} from '@angular/core';
  import { CommonModule } from '@angular/common';
import { Subscription } from 'rxjs';
import {
  BillboardCollection,
  Cartesian2,
  Cartesian3,
  CesiumWidget,
  Color,
  ColorGeometryInstanceAttribute,
  DistanceDisplayCondition,
  Ellipsoid,
  GeometryInstance,
  Globe,
  HorizontalOrigin,
  ImageryLayer,
  LabelCollection,
  LabelStyle,
  Material,
  Math as CesiumMath,
  NearFarScalar,
  PerInstanceColorAppearance,
  PolygonGeometry,
  PolygonHierarchy,
  PolygonOutlineGeometry,
  PolylineCollection,
  Primitive,
  ScreenSpaceEventHandler,
  ScreenSpaceEventType,
  SkyAtmosphere,
  UrlTemplateImageryProvider,
  VerticalOrigin,
} from '@cesium/engine';

import * as cesium from '../../../../cesium/src/public-api';
import {
  OfflineCanvasTileImageryProvider,
  getOfflineCitiesList,
  getOfflineCountryCentroids,
  getOfflineRiverLabelsList,
} from './offline-world-map';

interface PickedFeatureRef {
  kind: 'polygon' | 'polyline' | 'marker';
  item: cesium.Polygon | cesium.Polyline | cesium.Marker;
  cartesian: any;
  popupHtml: string;
}

@Component({
  selector: 'lib-cesium-specific',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './cesium-specific.component.html',
  styleUrl: './cesium-specific.component.css',
  encapsulation: ViewEncapsulation.None,
})
export class CesiumSpecificComponent implements OnInit, AfterViewInit, OnDestroy {
  @Input() Name: string = '';

  @ViewChild('mapContainer', { static: true })
  mapContainer!: ElementRef<HTMLDivElement>;

  public frontRepo?: cesium.FrontRepo;
  public mapSingleton?: cesium.MapView;

  public isDarkTheme = false;
  public activePopupHtml = '';
  public popupScreenX = 0;
  public popupScreenY = 0;

  private widget?: CesiumWidget;
  private eventHandler?: ScreenSpaceEventHandler;
  private postRenderRemoveCallback?: () => void;
  private currentStyleKey = '';
  private frontRepoSub?: Subscription;
  private resizeObserver?: ResizeObserver;

  private baseLabelCollection?: LabelCollection;
  private baseBillboardCollection?: BillboardCollection;

  private featurePrimitives: any[] = [];
  private popupCartesian?: any;
  private svgDataUrlCache = new Map<string, string>();

  private lastCameraState = {
    lng: -9999,
    lat: -9999,
    zoom: -1,
    pitch: -9999,
    bearing: -9999,
  };

  constructor(
    private frontRepoService: cesium.FrontRepoService,
    private mapViewService: cesium.MapViewService,
    private polygonService: cesium.PolygonService,
    private polylineService: cesium.PolylineService,
    private markerService: cesium.MarkerService,
    private cdr: ChangeDetectorRef
  ) {}

  ngOnInit(): void {
    this.frontRepoSub = this.frontRepoService.connectToWebSocket(this.Name).subscribe({
      next: (frontRepo) => {
        this.frontRepo = frontRepo;

        const maps = this.frontRepo.getFrontArray<cesium.MapView>(cesium.MapView.GONGSTRUCT_NAME);
        this.mapSingleton = maps.length > 0 ? maps[0] : undefined;

        this.syncGlobe();
        this.cdr.markForCheck();
      },
      error: (err) => {
        console.error(`[lib-cesium-specific] WebSocket error for "${this.Name}"`, err);
      },
    });
  }

  ngAfterViewInit(): void {
    this.initGlobeIfNeeded();

    if (typeof ResizeObserver !== 'undefined' && this.mapContainer?.nativeElement) {
      this.resizeObserver = new ResizeObserver(() => {
        this.widget?.resize();
      });
      this.resizeObserver.observe(this.mapContainer.nativeElement);
    }
  }

  ngOnDestroy(): void {
    this.frontRepoSub?.unsubscribe();
    this.resizeObserver?.disconnect();
    if (this.postRenderRemoveCallback) {
      this.postRenderRemoveCallback();
      this.postRenderRemoveCallback = undefined;
    }
    this.eventHandler?.destroy();
    if (this.widget && !this.widget.isDestroyed()) {
      this.widget.destroy();
    }
  }

  public resetCameraView(): void {
    if (!this.widget) return;
    const lng = this.mapSingleton?.Longitude ?? 2.3522;
    const lat = this.mapSingleton?.Latitude ?? 46.5;
    const zoom = this.mapSingleton?.Zoom || 5.2;
    const pitchDeg = this.mapSingleton?.Pitch ? -Math.max(20, 90 - this.mapSingleton.Pitch) : -45;
    const headingDeg = this.mapSingleton?.Bearing ?? 0;
    this.flyToView(lng, lat, zoom, pitchDeg, headingDeg, 1.0);
  }

  public tiltCamera(): void {
    if (!this.widget) return;
    const lng = this.mapSingleton?.Longitude ?? 2.3522;
    const lat = (this.mapSingleton?.Latitude ?? 46.5) - 2.2;
    const zoom = Math.max(this.mapSingleton?.Zoom || 5.5, 5.5);
    this.flyToView(lng, lat, zoom, -35, 0, 0.9);
  }

  public topDownCamera(): void {
    if (!this.widget) return;
    const lng = this.mapSingleton?.Longitude ?? 2.3522;
    const lat = this.mapSingleton?.Latitude ?? 46.5;
    const zoom = this.mapSingleton?.Zoom || 5.2;
    this.flyToView(lng, lat, zoom, -89.9, 0, 0.9);
  }

  public toggleBaseTheme(): void {
    this.isDarkTheme = !this.isDarkTheme;
    this.applyBaseLayer(this.isDarkTheme ? 'offline-dark' : 'offline-light');
    this.rebuildBaseWorldLabels();
  }

  public closePopup(): void {
    this.activePopupHtml = '';
    this.popupCartesian = undefined;
    this.cdr.markForCheck();
  }

  private initGlobeIfNeeded(): void {
    if (this.widget || !this.mapContainer?.nativeElement) {
      return;
    }

    const styleKey = this.resolveStyleKey(this.mapSingleton?.StyleURL);
    this.isDarkTheme = styleKey === 'offline-dark';
    this.currentStyleKey = styleKey;

    const baseProvider = this.createImageryProvider(styleKey);
    const globe = new Globe(Ellipsoid.WGS84);
    globe.baseColor = this.isDarkTheme
      ? Color.fromCssColorString('#0b192c')
      : Color.fromCssColorString('#cbe0ee');
    globe.depthTestAgainstTerrain = false;
    globe.showGroundAtmosphere = true;

    this.widget = new CesiumWidget(this.mapContainer.nativeElement, {
      baseLayer: new ImageryLayer(baseProvider as any),
      skyBox: false,
      skyAtmosphere: new SkyAtmosphere(),
      globe,
      requestRenderMode: false,
      contextOptions: {
        webgl: {
          alpha: false,
          antialias: true,
        },
      },
    });

    this.widget.scene.backgroundColor = Color.fromCssColorString('#050b14');

    this.rebuildBaseWorldLabels();
    this.bindInteractions();

    const initialLng = this.mapSingleton?.Longitude ?? 2.3522;
    const initialLat = this.mapSingleton?.Latitude ?? 46.5;
    const initialZoom = this.mapSingleton?.Zoom || 5.2;
    const initialPitch =
      this.mapSingleton?.Pitch && this.mapSingleton.Pitch > 0
        ? -Math.max(18, 90 - this.mapSingleton.Pitch)
        : -48;
    const initialBearing = this.mapSingleton?.Bearing ?? 0;

    this.flyToView(initialLng, initialLat, initialZoom, initialPitch, initialBearing, 0);

    if (this.frontRepo) {
      this.syncGlobe();
    }
  }

  private resolveStyleKey(rawStyleUrl: string | undefined): string {
    const trimmed = (rawStyleUrl || '').trim();
    if (!trimmed || trimmed === 'offline' || trimmed === 'offline-light') {
      return 'offline-light';
    }
    if (trimmed === 'offline-dark' || trimmed === 'dark') {
      return 'offline-dark';
    }
    return trimmed;
  }

  private createImageryProvider(styleKey: string): any {
    if (styleKey === 'offline-dark') {
      return new OfflineCanvasTileImageryProvider(true);
    }
    if (styleKey === 'offline-light') {
      return new OfflineCanvasTileImageryProvider(false);
    }
    if (styleKey === 'osm-online') {
      return new UrlTemplateImageryProvider({
        url: 'https://tile.openstreetmap.org/{z}/{x}/{y}.png',
        maximumLevel: 18,
      });
    }
    if (styleKey.startsWith('http://') || styleKey.startsWith('https://')) {
      return new UrlTemplateImageryProvider({
        url: styleKey,
        maximumLevel: 18,
      });
    }
    return new OfflineCanvasTileImageryProvider(false);
  }

  private applyBaseLayer(styleKey: string): void {
    if (!this.widget) return;
    this.currentStyleKey = styleKey;
    const layers = this.widget.imageryLayers;
    layers.removeAll(true);
    layers.addImageryProvider(this.createImageryProvider(styleKey));
    this.widget.scene.globe.baseColor =
      styleKey === 'offline-dark'
        ? Color.fromCssColorString('#0b192c')
        : Color.fromCssColorString('#cbe0ee');
  }

  private rebuildBaseWorldLabels(): void {
    if (!this.widget) return;
    const primitives = this.widget.scene.primitives;
    if (this.baseLabelCollection) {
      primitives.remove(this.baseLabelCollection);
      this.baseLabelCollection = undefined;
    }
    if (this.baseBillboardCollection) {
      primitives.remove(this.baseBillboardCollection);
      this.baseBillboardCollection = undefined;
    }

    const labels = new LabelCollection();
    const billboards = new BillboardCollection();

    const dark = this.isDarkTheme;
    const countryFill = dark
      ? Color.fromCssColorString('#94a3b8')
      : Color.fromCssColorString('#475569');
    const countryOutline = dark
      ? Color.fromCssColorString('#0f172a')
      : Color.fromCssColorString('#ffffff');

    // 1. Country names on the 3D globe
    const countries = getOfflineCountryCentroids();
    for (const c of countries) {
      if (c.areaDeg2 < 1.5) continue;
      const maxDist = c.areaDeg2 > 40 ? 14000000 : c.areaDeg2 > 8 ? 7500000 : 3500000;
      labels.add({
        position: Cartesian3.fromDegrees(c.lon, c.lat, 50),
        text: c.name.toUpperCase(),
        font: c.areaDeg2 > 25 ? '700 12px Roboto, sans-serif' : '600 10px Roboto, sans-serif',
        fillColor: countryFill,
        outlineColor: countryOutline,
        outlineWidth: 3,
        style: LabelStyle.FILL_AND_OUTLINE,
        horizontalOrigin: HorizontalOrigin.CENTER,
        verticalOrigin: VerticalOrigin.CENTER,
        distanceDisplayCondition: new DistanceDisplayCondition(150000, maxDist),
        scaleByDistance: new NearFarScalar(300000, 1.0, maxDist, 0.75),
      });
    }

    // 2. Capitals & Major Cities
    const cities = getOfflineCitiesList();
    const capDotUrl = this.getCapitalDotDataUrl(dark);
    const cityDotUrl = this.getCityDotDataUrl(dark);
    const capTextColor = dark
      ? Color.fromCssColorString('#f8fafc')
      : Color.fromCssColorString('#0f172a');
    const cityTextColor = dark
      ? Color.fromCssColorString('#cbd5e1')
      : Color.fromCssColorString('#334155');

    for (const city of cities) {
      const isCap = city.c === 1;
      const maxDist = isCap
        ? 5500000
        : city.s <= 2
        ? 3500000
        : city.s <= 5
        ? 1800000
        : 900000;
      const pos = Cartesian3.fromDegrees(city.lon, city.lat, 30);

      billboards.add({
        position: pos,
        image: isCap ? capDotUrl : cityDotUrl,
        width: isCap ? 10 : 6,
        height: isCap ? 10 : 6,
        distanceDisplayCondition: new DistanceDisplayCondition(1000, maxDist),
      });

      labels.add({
        position: pos,
        text: city.n,
        font: isCap ? '700 11px Roboto, sans-serif' : '500 10px Roboto, sans-serif',
        fillColor: isCap ? capTextColor : cityTextColor,
        outlineColor: countryOutline,
        outlineWidth: 3,
        style: LabelStyle.FILL_AND_OUTLINE,
        pixelOffset: new Cartesian2(0, -10),
        horizontalOrigin: HorizontalOrigin.CENTER,
        verticalOrigin: VerticalOrigin.BOTTOM,
        distanceDisplayCondition: new DistanceDisplayCondition(1000, maxDist),
      });
    }

    // 3. Major River labels
    const riverLabels = getOfflineRiverLabelsList();
    const riverColor = dark
      ? Color.fromCssColorString('#38bdf8')
      : Color.fromCssColorString('#2563eb');
    for (const riv of riverLabels) {
      const maxDist = riv.s <= 3 ? 2800000 : riv.s <= 6 ? 1500000 : 800000;
      labels.add({
        position: Cartesian3.fromDegrees(riv.lon, riv.lat, 20),
        text: riv.n,
        font: 'italic 600 10px Georgia, serif',
        fillColor: riverColor,
        outlineColor: countryOutline,
        outlineWidth: 2,
        style: LabelStyle.FILL_AND_OUTLINE,
        horizontalOrigin: HorizontalOrigin.CENTER,
        verticalOrigin: VerticalOrigin.CENTER,
        distanceDisplayCondition: new DistanceDisplayCondition(10000, maxDist),
      });
    }

    this.baseBillboardCollection = primitives.add(billboards);
    this.baseLabelCollection = primitives.add(labels);
  }

  private bindInteractions(): void {
    if (!this.widget) return;
    const scene = this.widget.scene;

    this.eventHandler = new ScreenSpaceEventHandler(scene.canvas);
    this.eventHandler.setInputAction((movement: { position: any }) => {
      if (!movement?.position) return;
      const picked = scene.pick(movement.position);
      const ref: PickedFeatureRef | undefined = picked?.id;
      if (!ref || !ref.kind) {
        this.closePopup();
        return;
      }

      if (ref.popupHtml) {
        this.activePopupHtml = ref.popupHtml;
        this.popupCartesian = ref.cartesian;
        this.updatePopupScreenPosition();
      } else {
        this.closePopup();
      }

      if (ref.kind === 'polygon') {
        this.polygonService.updateFront(ref.item as cesium.Polygon, this.Name).subscribe();
      } else if (ref.kind === 'polyline') {
        this.polylineService.updateFront(ref.item as cesium.Polyline, this.Name).subscribe();
      } else if (ref.kind === 'marker') {
        this.markerService.updateFront(ref.item as cesium.Marker, this.Name).subscribe();
      }
      this.cdr.markForCheck();
    }, ScreenSpaceEventType.LEFT_CLICK);

    const onPostRender = () => {
      if (this.popupCartesian && this.activePopupHtml) {
        this.updatePopupScreenPosition();
      }
    };
    scene.postRender.addEventListener(onPostRender);
    this.postRenderRemoveCallback = () => {
      scene.postRender.removeEventListener(onPostRender);
    };
  }

  private updatePopupScreenPosition(): void {
    if (!this.widget || !this.popupCartesian) return;
    const windowCoord = this.widget.scene.cartesianToCanvasCoordinates(this.popupCartesian);
    if (windowCoord) {
      this.popupScreenX = windowCoord.x;
      this.popupScreenY = windowCoord.y;
      this.cdr.detectChanges();
    }
  }

  private syncGlobe(): void {
    this.initGlobeIfNeeded();
    if (!this.widget || !this.frontRepo) {
      return;
    }

    if (this.mapSingleton) {
      const targetStyleKey = this.resolveStyleKey(this.mapSingleton.StyleURL);
      if (targetStyleKey !== this.currentStyleKey) {
        this.isDarkTheme = targetStyleKey === 'offline-dark';
        this.applyBaseLayer(targetStyleKey);
        this.rebuildBaseWorldLabels();
      }

      const targetLng = this.mapSingleton.Longitude || 0;
      const targetLat = this.mapSingleton.Latitude || 0;
      const targetZoom = this.mapSingleton.Zoom || 5;
      const targetPitch = this.mapSingleton.Pitch || 0;
      const targetBearing = this.mapSingleton.Bearing || 0;

      const cameraChanged =
        Math.abs(targetLng - this.lastCameraState.lng) > 1e-4 ||
        Math.abs(targetLat - this.lastCameraState.lat) > 1e-4 ||
        Math.abs(targetZoom - this.lastCameraState.zoom) > 1e-2 ||
        Math.abs(targetPitch - this.lastCameraState.pitch) > 1e-1 ||
        Math.abs(targetBearing - this.lastCameraState.bearing) > 1e-1;

      if (cameraChanged) {
        this.lastCameraState = {
          lng: targetLng,
          lat: targetLat,
          zoom: targetZoom,
          pitch: targetPitch,
          bearing: targetBearing,
        };
        const pitchDeg = targetPitch > 0 ? -Math.max(18, 90 - targetPitch) : -50;
        // Offset latitude slightly south when tilted so the target airspace is centered in the 3D viewport
        const altitudeMeters = this.zoomToAltitudeMeters(targetZoom);
        const latOffsetDeg =
          pitchDeg > -80 ? -(altitudeMeters / 111320) * Math.cos(CesiumMath.toRadians(-pitchDeg)) * 0.42 : 0;

        this.flyToView(
          targetLng,
          Math.max(-85, Math.min(85, targetLat + latOffsetDeg)),
          targetZoom,
          pitchDeg,
          targetBearing,
          0.8
        );
      }
    }

    this.renderLayers();
  }

  private zoomToAltitudeMeters(zoom: number): number {
    const clampedZoom = Math.max(1, Math.min(18, zoom || 5));
    return Math.max(800, 32000000 / Math.pow(2, clampedZoom));
  }

  private flyToView(
    lng: number,
    lat: number,
    zoom: number,
    pitchDeg: number,
    headingDeg: number,
    durationSeconds: number
  ): void {
    if (!this.widget) return;
    const height = this.zoomToAltitudeMeters(zoom);
    this.widget.camera.flyTo({
      destination: Cartesian3.fromDegrees(lng, lat, height),
      orientation: {
        heading: CesiumMath.toRadians(headingDeg),
        pitch: CesiumMath.toRadians(pitchDeg),
        roll: 0,
      },
      duration: durationSeconds,
    });
  }

  private clearFeaturePrimitives(): void {
    if (!this.widget) return;
    const primitives = this.widget.scene.primitives;
    for (const p of this.featurePrimitives) {
      primitives.remove(p);
    }
    this.featurePrimitives = [];
  }

  private renderLayers(): void {
    if (!this.widget || !this.frontRepo) {
      return;
    }

    this.clearFeaturePrimitives();

    const layers =
      this.mapSingleton?.Layers && this.mapSingleton.Layers.length > 0
        ? this.mapSingleton.Layers
        : this.frontRepo.getFrontArray<cesium.Layer>(cesium.Layer.GONGSTRUCT_NAME);
    const visibleLayers = layers.filter((layer) => layer.IsVisible);

    const volumeInstances: any[] = [];
    const outlineInstances: any[] = [];
    const polylines = new PolylineCollection();
    const billboards = new BillboardCollection();
    const labels = new LabelCollection();

    for (const layer of visibleLayers) {
      // 1. Polygons (3D Extruded Airspace Volumes + Wireframe Outlines + Roof/Floor Contours)
      for (const poly of layer.Polygons || []) {
        const coords = poly.Coordinates || [];
        if (coords.length < 3) continue;

        const degs: number[] = [];
        let sumLon = 0;
        let sumLat = 0;
        let prevLon = NaN;
        let prevLat = NaN;
        for (const c of coords) {
          if (Math.abs(c.Longitude - prevLon) < 1e-6 && Math.abs(c.Latitude - prevLat) < 1e-6) {
            continue;
          }
          degs.push(c.Longitude, c.Latitude);
          sumLon += c.Longitude;
          sumLat += c.Latitude;
          prevLon = c.Longitude;
          prevLat = c.Latitude;
        }
        const uniquePts = degs.length / 2;
        if (uniquePts < 3) continue;

        const centroidLon = sumLon / uniquePts;
        const centroidLat = sumLat / uniquePts;
        const baseAlt = Math.max(0, poly.BaseAltitude || 0);
        const extrudeHeight = poly.IsExtruded ? Math.max(poly.ExtrusionHeight || 0, 120) : 60;
        const topAlt = baseAlt + extrudeHeight;

        const fillCss = poly.FillColor || '#2563eb';
        const fillOpacity =
          poly.FillOpacity !== undefined && poly.FillOpacity > 0
            ? Math.min(0.85, Math.max(0.22, poly.FillOpacity))
            : 0.35;
        const fillColor = Color.fromCssColorString(fillCss).withAlpha(fillOpacity);

        const lineCss = poly.StrokeColor || '#1e40af';
        const lineOpacity =
          poly.StrokeOpacity !== undefined && poly.StrokeOpacity > 0 ? poly.StrokeOpacity : 0.95;
        const lineColor = Color.fromCssColorString(lineCss).withAlpha(lineOpacity);
        const lineWidth = Math.max(1.5, poly.StrokeWidth || 2.0);

        const pickRef: PickedFeatureRef = {
          kind: 'polygon',
          item: poly,
          cartesian: Cartesian3.fromDegrees(centroidLon, centroidLat, topAlt),
          popupHtml: poly.PopupHTML || `<strong>${this.escapeHtml(poly.Name)}</strong>`,
        };

        const positions = Cartesian3.fromDegreesArray(degs);
        const hierarchy = new PolygonHierarchy(positions);

        volumeInstances.push(
          new GeometryInstance({
            geometry: new PolygonGeometry({
              polygonHierarchy: hierarchy,
              height: baseAlt,
              extrudedHeight: topAlt,
              vertexFormat: PerInstanceColorAppearance.VERTEX_FORMAT,
            }),
            attributes: {
              color: ColorGeometryInstanceAttribute.fromColor(fillColor),
            },
            id: pickRef,
          })
        );

        outlineInstances.push(
          new GeometryInstance({
            geometry: new PolygonOutlineGeometry({
              polygonHierarchy: hierarchy,
              height: baseAlt,
              extrudedHeight: topAlt,
            }),
            attributes: {
              color: ColorGeometryInstanceAttribute.fromColor(lineColor),
            },
            id: pickRef,
          })
        );

        // Top roof contour polyline with configurable width
        const roofPositions: any[] = [];
        for (let i = 0; i < degs.length; i += 2) {
          roofPositions.push(Cartesian3.fromDegrees(degs[i], degs[i + 1], topAlt));
        }
        if (roofPositions.length > 0) {
          roofPositions.push(roofPositions[0]);
        }
        polylines.add({
          positions: roofPositions,
          width: lineWidth,
          material: Material.fromType('Color', { color: lineColor }),
          id: pickRef,
        });

        // Base footprint contour if elevated above ground
        if (baseAlt > 50) {
          const floorPositions: any[] = [];
          for (let i = 0; i < degs.length; i += 2) {
            floorPositions.push(Cartesian3.fromDegrees(degs[i], degs[i + 1], baseAlt));
          }
          if (floorPositions.length > 0) {
            floorPositions.push(floorPositions[0]);
          }
          polylines.add({
            positions: floorPositions,
            width: Math.max(1.0, lineWidth * 0.75),
            material: Material.fromType('Color', {
              color: lineColor.withAlpha(lineOpacity * 0.75),
            }),
            id: pickRef,
          });
        }

        if (poly.Label) {
          labels.add({
            position: Cartesian3.fromDegrees(centroidLon, centroidLat, topAlt + 150),
            text: poly.Label,
            font: '600 11px Roboto, sans-serif',
            fillColor: Color.fromCssColorString('#f8fafc'),
            outlineColor: Color.fromCssColorString('#0f172a'),
            outlineWidth: 3,
            style: LabelStyle.FILL_AND_OUTLINE,
            horizontalOrigin: HorizontalOrigin.CENTER,
            verticalOrigin: VerticalOrigin.BOTTOM,
            distanceDisplayCondition: new DistanceDisplayCondition(0, 3500000),
            id: pickRef,
          });
        }
      }

      // 2. Polylines (GeoBorders, Flight Paths, Routes)
      for (const line of layer.Polylines || []) {
        const coords = line.Coordinates || [];
        if (coords.length < 2) continue;

        const positions: any[] = [];
        for (const c of coords) {
          positions.push(Cartesian3.fromDegrees(c.Longitude, c.Latitude, c.Altitude || 60));
        }
        const midIdx = Math.floor(coords.length / 2);
        const midCoord = coords[midIdx];
        const midCartesian = Cartesian3.fromDegrees(
          midCoord.Longitude,
          midCoord.Latitude,
          (midCoord.Altitude || 60) + 100
        );

        const lineColor = Color.fromCssColorString(line.StrokeColor || '#0284c7').withAlpha(
          line.StrokeOpacity !== undefined && line.StrokeOpacity > 0 ? line.StrokeOpacity : 0.95
        );

        const pickRef: PickedFeatureRef = {
          kind: 'polyline',
          item: line,
          cartesian: midCartesian,
          popupHtml: line.PopupHTML || `<strong>${this.escapeHtml(line.Name)}</strong>`,
        };

        polylines.add({
          positions,
          width: Math.max(1.5, line.StrokeWidth || 2.5),
          material: Material.fromType('Color', { color: lineColor }),
          id: pickRef,
        });

        if (line.Label) {
          labels.add({
            position: midCartesian,
            text: line.Label,
            font: '600 11px Roboto, sans-serif',
            fillColor: Color.fromCssColorString('#f8fafc'),
            outlineColor: Color.fromCssColorString('#0f172a'),
            outlineWidth: 3,
            style: LabelStyle.FILL_AND_OUTLINE,
            horizontalOrigin: HorizontalOrigin.CENTER,
            verticalOrigin: VerticalOrigin.BOTTOM,
            distanceDisplayCondition: new DistanceDisplayCondition(0, 3000000),
            id: pickRef,
          });
        }
      }

      // 3. Markers (3D SVG Billboards + Labels)
      for (const marker of layer.Markers || []) {
        const alt = marker.Altitude || 80;
        const pos = Cartesian3.fromDegrees(marker.Longitude, marker.Latitude, alt);
        const pickRef: PickedFeatureRef = {
          kind: 'marker',
          item: marker,
          cartesian: pos,
          popupHtml: marker.PopupHTML || `<strong>${this.escapeHtml(marker.Name)}</strong>`,
        };

        const iconUrl = this.getMarkerSvgDataUrl(marker.SVG, marker.Color || '#ef4444');
        const scale = marker.Scale && marker.Scale > 0 ? marker.Scale : 1.0;

        billboards.add({
          position: pos,
          image: iconUrl,
          scale,
          rotation: marker.Rotation ? -CesiumMath.toRadians(marker.Rotation) : 0,
          horizontalOrigin: HorizontalOrigin.CENTER,
          verticalOrigin: VerticalOrigin.CENTER,
          id: pickRef,
        });

        if (marker.Label) {
          labels.add({
            position: pos,
            text: marker.Label,
            font: '600 11px Roboto, sans-serif',
            fillColor: Color.fromCssColorString('#0f172a'),
            outlineColor: Color.fromCssColorString('#ffffff'),
            outlineWidth: 3,
            style: LabelStyle.FILL_AND_OUTLINE,
            pixelOffset: new Cartesian2(0, 18),
            horizontalOrigin: HorizontalOrigin.CENTER,
            verticalOrigin: VerticalOrigin.TOP,
            distanceDisplayCondition: new DistanceDisplayCondition(0, 3500000),
            id: pickRef,
          });
        }
      }
    }

    const primitives = this.widget.scene.primitives;

    if (volumeInstances.length > 0) {
      const volPrim = primitives.add(
        new Primitive({
          geometryInstances: volumeInstances,
          appearance: new PerInstanceColorAppearance({
            translucent: true,
            closed: false,
          }),
          asynchronous: false,
        })
      );
      this.featurePrimitives.push(volPrim);
    }

    if (outlineInstances.length > 0) {
      const outPrim = primitives.add(
        new Primitive({
          geometryInstances: outlineInstances,
          appearance: new PerInstanceColorAppearance({
            flat: true,
            translucent: true,
          }),
          asynchronous: false,
        })
      );
      this.featurePrimitives.push(outPrim);
    }

    if (polylines.length > 0) {
      this.featurePrimitives.push(primitives.add(polylines));
    }
    if (billboards.length > 0) {
      this.featurePrimitives.push(primitives.add(billboards));
    }
    if (labels.length > 0) {
      this.featurePrimitives.push(primitives.add(labels));
    }
  }

  private getMarkerSvgDataUrl(rawSvg: string | undefined, color: string): string {
    const key = `${rawSvg || ''}::${color}`;
    const cached = this.svgDataUrlCache.get(key);
    if (cached) return cached;

    let svg = (rawSvg || '').trim();
    if (!svg) {
      svg = `<svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24" fill="${color}" stroke="#ffffff" stroke-width="2"><path d="M12 2C8.13 2 5 5.13 5 9c0 5.25 7 13 7 13s7-7.75 7-13c0-3.87-3.13-7-7-7z"/><circle cx="12" cy="9" r="2.5" fill="#ffffff"/></svg>`;
    } else if (!svg.includes('xmlns=')) {
      svg = svg.replace('<svg', '<svg xmlns="http://www.w3.org/2000/svg"');
    }

    const dataUrl = `data:image/svg+xml;utf8,${encodeURIComponent(svg)}`;
    this.svgDataUrlCache.set(key, dataUrl);
    return dataUrl;
  }

  private getCapitalDotDataUrl(dark: boolean): string {
    const key = `cap-dot-${dark}`;
    const cached = this.svgDataUrlCache.get(key);
    if (cached) return cached;
    const svg = `<svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 12 12"><circle cx="6" cy="6" r="5" fill="#ffffff" stroke="#dc2626" stroke-width="1.8"/><circle cx="6" cy="6" r="2.2" fill="#dc2626"/></svg>`;
    const url = `data:image/svg+xml;utf8,${encodeURIComponent(svg)}`;
    this.svgDataUrlCache.set(key, url);
    return url;
  }

  private getCityDotDataUrl(dark: boolean): string {
    const key = `city-dot-${dark}`;
    const cached = this.svgDataUrlCache.get(key);
    if (cached) return cached;
    const fill = dark ? '#cbd5e1' : '#334155';
    const stroke = dark ? '#0f172a' : '#ffffff';
    const svg = `<svg xmlns="http://www.w3.org/2000/svg" width="8" height="8" viewBox="0 0 8 8"><circle cx="4" cy="4" r="3" fill="${fill}" stroke="${stroke}" stroke-width="1.2"/></svg>`;
    const url = `data:image/svg+xml;utf8,${encodeURIComponent(svg)}`;
    this.svgDataUrlCache.set(key, url);
    return url;
  }

  private escapeHtml(text: string): string {
    return (text || '')
      .replace(/&/g, '&amp;')
      .replace(/</g, '&lt;')
      .replace(/>/g, '&gt;');
  }
}
