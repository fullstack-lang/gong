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
import * as maplibregl from 'maplibre-gl';

import * as maplibre from '../../../../maplibre/src/public-api';
import { ensureMaplibreInlineWorker } from './maplibre-worker-inline';
import {
  buildOfflineVectorStyle,
  getOfflineCitiesList,
  getOfflineRiverLabelsList,
} from './offline-world-map';

@Component({
  selector: 'lib-maplibre-specific',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './maplibre-specific.component.html',
  styleUrl: './maplibre-specific.component.css',
  encapsulation: ViewEncapsulation.None,
})
export class MaplibreSpecificComponent implements OnInit, AfterViewInit, OnDestroy {
  @Input() Name: string = '';

  @ViewChild('mapContainer', { static: true })
  mapContainer!: ElementRef<HTMLDivElement>;

  public frontRepo?: maplibre.FrontRepo;
  public mapSingleton?: maplibre.MapView;

  private map?: maplibregl.Map;
  private isMapStyleLoaded = false;
  private mapInteractionsBound = false;
  private currentStyleKey = '';
  private frontRepoSub?: Subscription;
  private resizeObserver?: ResizeObserver;
  private moveDebounceTimer: any;

  private activeMarkers: maplibregl.Marker[] = [];
  private activeLabelMarkers: maplibregl.Marker[] = [];
  private activeBaseLabelMarkers: maplibregl.Marker[] = [];
  private polygonById = new Map<number, maplibre.Polygon>();
  private polylineById = new Map<number, maplibre.Polyline>();

  private lastCameraState = {
    lng: -9999,
    lat: -9999,
    zoom: -1,
    pitch: -1,
    bearing: -9999,
  };

  constructor(
    private frontRepoService: maplibre.FrontRepoService,
    private mapViewService: maplibre.MapViewService,
    private polygonService: maplibre.PolygonService,
    private polylineService: maplibre.PolylineService,
    private markerService: maplibre.MarkerService,
    private cdr: ChangeDetectorRef
  ) {}

  ngOnInit(): void {
    ensureMaplibreInlineWorker();

    this.frontRepoSub = this.frontRepoService.connectToWebSocket(this.Name).subscribe({
      next: (frontRepo) => {
        this.frontRepo = frontRepo;

        const maps = this.frontRepo.getFrontArray<maplibre.MapView>(maplibre.MapView.GONGSTRUCT_NAME);
        this.mapSingleton = maps.length > 0 ? maps[0] : undefined;

        this.syncMap();
        this.cdr.markForCheck();
      },
      error: (err) => {
        console.error(`[lib-maplibre-specific] WebSocket error for "${this.Name}"`, err);
      },
    });
  }

  ngAfterViewInit(): void {
    ensureMaplibreInlineWorker();
    this.initMapIfNeeded();

    if (typeof ResizeObserver !== 'undefined' && this.mapContainer?.nativeElement) {
      this.resizeObserver = new ResizeObserver(() => {
        this.map?.resize();
      });
      this.resizeObserver.observe(this.mapContainer.nativeElement);
    }
  }

  ngOnDestroy(): void {
    this.frontRepoSub?.unsubscribe();
    if (this.moveDebounceTimer) {
      clearTimeout(this.moveDebounceTimer);
    }
    this.resizeObserver?.disconnect();
    this.clearMarkers();
    this.clearBaseLabelMarkers();
    if (this.map) {
      this.map.remove();
      this.map = undefined;
    }
  }

  private resolveStyle(styleURL: string | undefined): string | maplibregl.StyleSpecification {
    const key = (styleURL || 'offline').trim();
    if (key.startsWith('http://') || key.startsWith('https://')) {
      return key;
    }
    if (key === 'offline-dark' || key === 'carto-dark' || key === 'dark') {
      return buildOfflineVectorStyle(true);
    }
    if (key === 'osm-online') {
      return {
        version: 8,
        sources: {
          osm: {
            type: 'raster',
            tiles: ['https://tile.openstreetmap.org/{z}/{x}/{y}.png'],
            tileSize: 256,
            attribution: '&copy; OpenStreetMap contributors',
          },
        },
        layers: [{ id: 'osm-base', type: 'raster', source: 'osm' }],
      };
    }
    // Default for all MapLibre applications: 100% offline embedded Natural Earth 1:50m vector world map
    return buildOfflineVectorStyle(false);
  }

  private initMapIfNeeded(): void {
    if (this.map || !this.mapContainer?.nativeElement) {
      return;
    }

    ensureMaplibreInlineWorker();

    const lng = this.mapSingleton?.Longitude ?? 2.35;
    const lat = this.mapSingleton?.Latitude ?? 48.85;
    const zoom = this.mapSingleton?.Zoom || 6;
    const pitch = this.mapSingleton?.Pitch ?? 0;
    const bearing = this.mapSingleton?.Bearing ?? 0;
    const styleKey = (this.mapSingleton?.StyleURL || 'offline').trim();
    this.currentStyleKey = styleKey;

    this.map = new maplibregl.Map({
      container: this.mapContainer.nativeElement,
      style: this.resolveStyle(styleKey),
      center: [lng, lat],
      zoom: zoom,
      pitch: pitch,
      bearing: bearing,
    });

    this.lastCameraState = { lng, lat, zoom, pitch, bearing };

    this.map.addControl(
      new maplibregl.NavigationControl({ visualizePitch: true }),
      'top-right'
    );
    this.map.addControl(new maplibregl.ScaleControl({ unit: 'nautical' }), 'bottom-left');

    this.map.on('load', () => {
      this.isMapStyleLoaded = true;
      this.ensureGeoJSONLayers();
      this.renderScene();
      this.updateOfflineBaseLabels();
    });

    this.map.on('style.load', () => {
      this.isMapStyleLoaded = true;
      this.ensureGeoJSONLayers();
      this.renderScene();
      this.updateOfflineBaseLabels();
    });

    this.map.on('moveend', (event: any) => {
      this.updateOfflineBaseLabels();

      // Only push camera updates when initiated by user interaction
      if (!event || !event.originalEvent || !this.map || !this.mapSingleton) {
        return;
      }

      const center = this.map.getCenter();
      const newZoom = this.map.getZoom();
      const newPitch = this.map.getPitch();
      const newBearing = this.map.getBearing();

      this.lastCameraState = {
        lng: center.lng,
        lat: center.lat,
        zoom: newZoom,
        pitch: newPitch,
        bearing: newBearing,
      };

      if (this.moveDebounceTimer) {
        clearTimeout(this.moveDebounceTimer);
      }
      this.moveDebounceTimer = setTimeout(() => {
        if (!this.mapSingleton) return;
        this.mapSingleton.Longitude = center.lng;
        this.mapSingleton.Latitude = center.lat;
        this.mapSingleton.Zoom = newZoom;
        this.mapSingleton.Pitch = newPitch;
        this.mapSingleton.Bearing = newBearing;
        this.mapViewService.updateFront(this.mapSingleton, this.Name).subscribe();
      }, 250);
    });
  }

  private syncMap(): void {
    this.initMapIfNeeded();
    if (!this.map) {
      return;
    }

    if (this.mapSingleton) {
      const nextStyleKey = (this.mapSingleton.StyleURL || 'offline').trim();
      if (nextStyleKey !== this.currentStyleKey) {
        this.currentStyleKey = nextStyleKey;
        this.isMapStyleLoaded = false;
        this.map.setStyle(this.resolveStyle(nextStyleKey));
      }

      const lng = this.mapSingleton.Longitude ?? 0;
      const lat = this.mapSingleton.Latitude ?? 0;
      const zoom = this.mapSingleton.Zoom || 6;
      const pitch = this.mapSingleton.Pitch ?? 0;
      const bearing = this.mapSingleton.Bearing ?? 0;

      const diff =
        Math.abs(lng - this.lastCameraState.lng) +
        Math.abs(lat - this.lastCameraState.lat) +
        Math.abs(zoom - this.lastCameraState.zoom) +
        Math.abs(pitch - this.lastCameraState.pitch) +
        Math.abs(bearing - this.lastCameraState.bearing);

      if (diff > 0.0005) {
        this.lastCameraState = { lng, lat, zoom, pitch, bearing };
        this.map.jumpTo({
          center: [lng, lat],
          zoom: zoom,
          pitch: pitch,
          bearing: bearing,
        });
      }
    }

    if (this.isMapStyleLoaded || this.map.isStyleLoaded()) {
      this.isMapStyleLoaded = true;
      this.ensureGeoJSONLayers();
      this.renderScene();
    }
  }

  private ensureGeoJSONLayers(): void {
    if (!this.map) return;

    const emptyFC: any = {
      type: 'FeatureCollection',
      features: [],
    };

    if (!this.map.getSource('gong-polygons-2d-src')) {
      this.map.addSource('gong-polygons-2d-src', {
        type: 'geojson',
        data: emptyFC,
      });
    }
    if (!this.map.getSource('gong-polygons-3d-src')) {
      this.map.addSource('gong-polygons-3d-src', {
        type: 'geojson',
        data: emptyFC,
      });
    }
    if (!this.map.getSource('gong-polylines-solid-src')) {
      this.map.addSource('gong-polylines-solid-src', {
        type: 'geojson',
        data: emptyFC,
      });
    }
    if (!this.map.getSource('gong-polylines-dashed-src')) {
      this.map.addSource('gong-polylines-dashed-src', {
        type: 'geojson',
        data: emptyFC,
      });
    }

    // 1. 2D Polygon Fill
    if (!this.map.getLayer('gong-polygons-2d-fill')) {
      this.map.addLayer({
        id: 'gong-polygons-2d-fill',
        type: 'fill',
        source: 'gong-polygons-2d-src',
        paint: {
          'fill-color': ['get', 'fillColor'],
          'fill-opacity': ['get', 'fillOpacity'],
        },
      });
    }

    // 2. 2D Polygon Outline
    if (!this.map.getLayer('gong-polygons-2d-stroke')) {
      this.map.addLayer({
        id: 'gong-polygons-2d-stroke',
        type: 'line',
        source: 'gong-polygons-2d-src',
        layout: {
          'line-cap': 'round',
          'line-join': 'round',
        },
        paint: {
          'line-color': ['get', 'strokeColor'],
          'line-opacity': ['get', 'strokeOpacity'],
          'line-width': ['get', 'strokeWidth'],
        },
      });
    }

    // 3. 3D Extruded Polygons (fill-extrusion-opacity is data-constant in MapLibre Style Spec)
    if (!this.map.getLayer('gong-polygons-3d-extrusion')) {
      this.map.addLayer({
        id: 'gong-polygons-3d-extrusion',
        type: 'fill-extrusion',
        source: 'gong-polygons-3d-src',
        paint: {
          'fill-extrusion-color': ['get', 'fillColor'],
          'fill-extrusion-opacity': 0.68,
          'fill-extrusion-base': ['get', 'baseAltitude'],
          'fill-extrusion-height': ['get', 'extrusionHeight'],
        },
      });
    }

    // 4. 3D Polygon Outline
    if (!this.map.getLayer('gong-polygons-3d-stroke')) {
      this.map.addLayer({
        id: 'gong-polygons-3d-stroke',
        type: 'line',
        source: 'gong-polygons-3d-src',
        layout: {
          'line-cap': 'round',
          'line-join': 'round',
        },
        paint: {
          'line-color': ['get', 'strokeColor'],
          'line-opacity': ['get', 'strokeOpacity'],
          'line-width': ['get', 'strokeWidth'],
        },
      });
    }

    // 5. Solid Polylines
    if (!this.map.getLayer('gong-polylines-solid-layer')) {
      this.map.addLayer({
        id: 'gong-polylines-solid-layer',
        type: 'line',
        source: 'gong-polylines-solid-src',
        layout: {
          'line-cap': 'round',
          'line-join': 'round',
        },
        paint: {
          'line-color': ['get', 'strokeColor'],
          'line-opacity': ['get', 'strokeOpacity'],
          'line-width': ['get', 'strokeWidth'],
        },
      });
    }

    // 6. Dashed Polylines
    if (!this.map.getLayer('gong-polylines-dashed-layer')) {
      this.map.addLayer({
        id: 'gong-polylines-dashed-layer',
        type: 'line',
        source: 'gong-polylines-dashed-src',
        layout: {
          'line-cap': 'round',
          'line-join': 'round',
        },
        paint: {
          'line-color': ['get', 'strokeColor'],
          'line-opacity': ['get', 'strokeOpacity'],
          'line-width': ['get', 'strokeWidth'],
          'line-dasharray': [4, 2],
        },
      });
    }

    this.bindMapInteractionsOnce();
  }

  private getInteractiveLayerIds(): string[] {
    if (!this.map) return [];
    const candidateIds = [
      'gong-polylines-dashed-layer',
      'gong-polylines-solid-layer',
      'gong-polygons-3d-stroke',
      'gong-polygons-3d-extrusion',
      'gong-polygons-2d-stroke',
      'gong-polygons-2d-fill',
    ];
    return candidateIds.filter((id) => !!this.map?.getLayer(id));
  }

  private bindMapInteractionsOnce(): void {
    if (!this.map || this.mapInteractionsBound) return;
    this.mapInteractionsBound = true;

    this.map.on('click', (e) => {
      if (!this.map) return;
      const layers = this.getInteractiveLayerIds();
      if (layers.length === 0) return;

      const features = this.map.queryRenderedFeatures(e.point, { layers });
      if (!features || features.length === 0) return;

      const feature = features[0];
      const gongId = feature.properties?.['gongId'];
      const popupHTML = feature.properties?.['popupHTML'];

      if (popupHTML && this.map) {
        new maplibregl.Popup({ maxWidth: '320px' })
          .setLngLat(e.lngLat)
          .setHTML(popupHTML)
          .addTo(this.map);
      }

      if (typeof gongId === 'number') {
        const layerId = feature.layer?.id || '';
        if (layerId.startsWith('gong-polylines-')) {
          const line = this.polylineById.get(gongId);
          if (line) {
            this.polylineService.updateFront(line, this.Name).subscribe();
          }
        } else {
          const poly = this.polygonById.get(gongId);
          if (poly) {
            this.polygonService.updateFront(poly, this.Name).subscribe();
          }
        }
      }
    });

    this.map.on('mousemove', (e) => {
      if (!this.map) return;
      const layers = this.getInteractiveLayerIds();
      if (layers.length === 0) return;
      const features = this.map.queryRenderedFeatures(e.point, { layers });
      this.map.getCanvas().style.cursor = features && features.length > 0 ? 'pointer' : '';
    });
  }

  private clearMarkers(): void {
    for (const m of this.activeMarkers) {
      m.remove();
    }
    this.activeMarkers = [];

    for (const lm of this.activeLabelMarkers) {
      lm.remove();
    }
    this.activeLabelMarkers = [];
  }

  private clearBaseLabelMarkers(): void {
    for (const bm of this.activeBaseLabelMarkers) {
      bm.remove();
    }
    this.activeBaseLabelMarkers = [];
  }

  private isOfflineStyleActive(): boolean {
    const key = (this.currentStyleKey || 'offline').trim();
    if (key.startsWith('http://') || key.startsWith('https://') || key === 'osm-online') {
      return false;
    }
    return true;
  }

  private isDarkOfflineStyleActive(): boolean {
    const key = (this.currentStyleKey || 'offline').trim();
    return key === 'offline-dark' || key === 'carto-dark' || key === 'dark';
  }

  private updateOfflineBaseLabels(): void {
    this.clearBaseLabelMarkers();
    if (!this.map || !this.isOfflineStyleActive()) {
      return;
    }

    const zoom = this.map.getZoom();
    if (zoom < 2.2) {
      return;
    }

    const bounds = this.map.getBounds();
    const west = bounds.getWest();
    const east = bounds.getEast();
    const south = bounds.getSouth();
    const north = bounds.getNorth();
    const isDark = this.isDarkOfflineStyleActive();

    const maxCityScalerank =
      zoom < 3.2 ? 1 : zoom < 4.3 ? 2 : zoom < 5.3 ? 4 : zoom < 6.4 ? 6 : 10;
    const maxRiverScalerank =
      zoom < 4.2 ? -1 : zoom < 5.3 ? 5 : zoom < 6.5 ? 8 : 10;
    const maxLabels =
      zoom < 3.5 ? 24 : zoom < 5.0 ? 42 : zoom < 6.5 ? 60 : 80;

    const occupiedCells = new Set<string>();
    const cellW = 74;
    const cellH = 22;

    const tryReserveScreenBox = (lon: number, lat: number): boolean => {
      if (!this.map) return false;
      const pt = this.map.project([lon, lat]);
      const cx = Math.floor(pt.x / cellW);
      const cy = Math.floor(pt.y / cellH);
      const key = `${cx}:${cy}`;
      if (occupiedCells.has(key)) {
        return false;
      }
      occupiedCells.add(key);
      return true;
    };

    let count = 0;
    const cities = getOfflineCitiesList();
    for (let i = 0; i < cities.length && count < maxLabels; i++) {
      const city = cities[i];
      if (city.lon < west || city.lon > east || city.lat < south || city.lat > north) {
        continue;
      }
      if (city.c !== 1 && city.s > maxCityScalerank) {
        continue;
      }
      if (zoom < 3.0 && city.c === 1 && city.s > 3) {
        continue;
      }
      if (!tryReserveScreenBox(city.lon, city.lat)) {
        continue;
      }

      const el = document.createElement('div');
      if (city.c === 1) {
        el.className = isDark
          ? 'gong-maplibre-base-label gong-maplibre-capital-label gong-maplibre-capital-label-dark'
          : 'gong-maplibre-base-label gong-maplibre-capital-label';
      } else {
        el.className = isDark
          ? 'gong-maplibre-base-label gong-maplibre-city-label gong-maplibre-city-label-dark'
          : 'gong-maplibre-base-label gong-maplibre-city-label';
      }
      el.textContent = city.n;

      const marker = new maplibregl.Marker({
        element: el,
        anchor: 'bottom-left',
        offset: [5, -3],
      })
        .setLngLat([city.lon, city.lat])
        .addTo(this.map);

      this.activeBaseLabelMarkers.push(marker);
      count++;
    }

    if (maxRiverScalerank >= 0 && count < maxLabels) {
      const rivers = getOfflineRiverLabelsList();
      for (let i = 0; i < rivers.length && count < maxLabels; i++) {
        const rv = rivers[i];
        if (rv.s > maxRiverScalerank) {
          continue;
        }
        if (rv.lon < west || rv.lon > east || rv.lat < south || rv.lat > north) {
          continue;
        }
        if (!tryReserveScreenBox(rv.lon, rv.lat)) {
          continue;
        }

        const el = document.createElement('div');
        el.className = isDark
          ? 'gong-maplibre-base-label gong-maplibre-river-label gong-maplibre-river-label-dark'
          : 'gong-maplibre-base-label gong-maplibre-river-label';
        el.textContent = rv.n;

        const marker = new maplibregl.Marker({
          element: el,
          anchor: 'center',
        })
          .setLngLat([rv.lon, rv.lat])
          .addTo(this.map);

        this.activeBaseLabelMarkers.push(marker);
        count++;
      }
    }
  }

  private renderScene(): void {
    if (!this.map || !this.frontRepo) return;

    this.clearMarkers();
    this.polygonById.clear();
    this.polylineById.clear();

    const polygons2D: any[] = [];
    const polygons3D: any[] = [];
    const polylinesSolid: any[] = [];
    const polylinesDashed: any[] = [];

    const layers: maplibre.Layer[] =
      this.mapSingleton && this.mapSingleton.Layers && this.mapSingleton.Layers.length > 0
        ? this.mapSingleton.Layers
        : this.frontRepo.getFrontArray<maplibre.Layer>(maplibre.Layer.GONGSTRUCT_NAME);

    for (const layer of layers) {
      if (!layer.IsVisible) {
        continue;
      }

      // Polygons
      if (layer.Polygons) {
        for (const poly of layer.Polygons) {
          if (!poly.Coordinates || poly.Coordinates.length < 3) continue;

          this.polygonById.set(poly.ID, poly);
          const ring: [number, number][] = poly.Coordinates.map((c) => [
            c.Longitude,
            c.Latitude,
          ]);
          const first = ring[0];
          const last = ring[ring.length - 1];
          if (first[0] !== last[0] || first[1] !== last[1]) {
            ring.push([first[0], first[1]]);
          }

          const feature: any = {
            type: 'Feature',
            geometry: {
              type: 'Polygon',
              coordinates: [ring],
            },
            properties: {
              gongId: poly.ID,
              name: poly.Name,
              label: poly.Label || '',
              popupHTML:
                poly.PopupHTML ||
                (poly.Label ? `<b>${poly.Label}</b>` : `<b>${poly.Name}</b>`),
              fillColor: poly.FillColor || '#3b82f6',
              fillOpacity: poly.FillOpacity !== undefined ? poly.FillOpacity : 0.3,
              strokeColor: poly.StrokeColor || '#1d4ed8',
              strokeOpacity: poly.StrokeOpacity !== undefined ? poly.StrokeOpacity : 1.0,
              strokeWidth: poly.StrokeWidth || 2.0,
              baseAltitude: poly.BaseAltitude || 0,
              extrusionHeight: poly.ExtrusionHeight || 0,
            },
          };

          if (poly.IsExtruded) {
            polygons3D.push(feature);
          } else {
            polygons2D.push(feature);
          }
        }
      }

      // Polylines
      if (layer.Polylines) {
        for (const line of layer.Polylines) {
          if (!line.Coordinates || line.Coordinates.length < 2) continue;

          this.polylineById.set(line.ID, line);
          const coords: [number, number][] = line.Coordinates.map((c) => [
            c.Longitude,
            c.Latitude,
          ]);

          const feature: any = {
            type: 'Feature',
            geometry: {
              type: 'LineString',
              coordinates: coords,
            },
            properties: {
              gongId: line.ID,
              name: line.Name,
              label: line.Label || '',
              popupHTML:
                line.PopupHTML ||
                (line.Label ? `<b>${line.Label}</b>` : `<b>${line.Name}</b>`),
              strokeColor: line.StrokeColor || '#0f766e',
              strokeOpacity: line.StrokeOpacity !== undefined ? line.StrokeOpacity : 1.0,
              strokeWidth: line.StrokeWidth || 2.5,
            },
          };

          if (line.StrokeDashArray && line.StrokeDashArray.trim().length > 0) {
            polylinesDashed.push(feature);
          } else {
            polylinesSolid.push(feature);
          }
        }
      }

      // Markers & SVG Objects
      if (layer.Markers) {
        for (const mk of layer.Markers) {
          this.createAndAddMarker(mk);
        }
      }
    }

    const poly2dSource = this.map.getSource('gong-polygons-2d-src') as
      | maplibregl.GeoJSONSource
      | undefined;
    poly2dSource?.setData({ type: 'FeatureCollection', features: polygons2D });

    const poly3dSource = this.map.getSource('gong-polygons-3d-src') as
      | maplibregl.GeoJSONSource
      | undefined;
    poly3dSource?.setData({ type: 'FeatureCollection', features: polygons3D });

    const lineSolidSource = this.map.getSource('gong-polylines-solid-src') as
      | maplibregl.GeoJSONSource
      | undefined;
    lineSolidSource?.setData({ type: 'FeatureCollection', features: polylinesSolid });

    const lineDashedSource = this.map.getSource('gong-polylines-dashed-src') as
      | maplibregl.GeoJSONSource
      | undefined;
    lineDashedSource?.setData({ type: 'FeatureCollection', features: polylinesDashed });

    this.updateOfflineBaseLabels();
  }

  private createAndAddMarker(mk: maplibre.Marker): void {
    if (!this.map) return;

    const rootEl = document.createElement('div');
    rootEl.className = 'gong-maplibre-marker';

    const svgContainer = document.createElement('div');
    svgContainer.className = 'gong-maplibre-marker-svg';
    const scale = mk.Scale && mk.Scale > 0 ? mk.Scale : 1.0;
    svgContainer.style.transform = `scale(${scale})`;

    const hasCustomSVG = !!(mk.SVG && mk.SVG.trim().length > 0);
    if (hasCustomSVG) {
      svgContainer.innerHTML = mk.SVG;
      if (mk.Color) {
        svgContainer.style.color = mk.Color;
      }
    } else {
      const pinColor = mk.Color || '#2563eb';
      svgContainer.innerHTML = `<svg width="28" height="36" viewBox="0 0 28 36" xmlns="http://www.w3.org/2000/svg">
        <path d="M14 1 C6.8 1 1 6.8 1 14 C1 23.5 14 35 14 35 C14 35 27 23.5 27 14 C27 6.8 21.2 1 14 1 Z" fill="${pinColor}" stroke="#ffffff" stroke-width="2"/>
        <circle cx="14" cy="14" r="5" fill="#ffffff"/>
      </svg>`;
    }

    rootEl.appendChild(svgContainer);

    // Only append a separate DOM text pill when no custom SVG badge is provided
    if (!hasCustomSVG && mk.Label && mk.Label.trim().length > 0) {
      const labelEl = document.createElement('div');
      labelEl.className = 'gong-maplibre-marker-label';
      labelEl.textContent = mk.Label;
      rootEl.appendChild(labelEl);
    }

    const anchor = (mk.Anchor || 'center') as maplibregl.PositionAnchor;
    const rotAlign = (mk.RotationAlignment || 'auto') as 'map' | 'viewport' | 'auto';
    const pitchAlign = (mk.PitchAlignment || 'auto') as 'map' | 'viewport' | 'auto';

    const mlMarker = new maplibregl.Marker({
      element: rootEl,
      anchor: anchor,
      rotation: mk.Rotation || 0,
      rotationAlignment: rotAlign,
      pitchAlignment: pitchAlign,
      draggable: !!mk.IsDraggable,
    }).setLngLat([mk.Longitude, mk.Latitude]);

    if (mk.PopupHTML && mk.PopupHTML.trim().length > 0) {
      const popup = new maplibregl.Popup({ offset: 16, maxWidth: '300px' }).setHTML(
        mk.PopupHTML
      );
      mlMarker.setPopup(popup);
    }

    rootEl.addEventListener('click', (ev: MouseEvent) => {
      ev.stopPropagation();
      this.markerService.updateFront(mk, this.Name).subscribe();
    });

    if (mk.IsDraggable) {
      mlMarker.on('dragend', () => {
        const lngLat = mlMarker.getLngLat();
        mk.Longitude = lngLat.lng;
        mk.Latitude = lngLat.lat;
        this.markerService.updateFront(mk, this.Name).subscribe();
      });
    }

    mlMarker.addTo(this.map);
    this.activeMarkers.push(mlMarker);
  }
}
