import { Component, ChangeDetectionStrategy, CUSTOM_ELEMENTS_SCHEMA, Input, ChangeDetectorRef, ViewChild, Directive, inject, OnChanges, SimpleChanges } from '@angular/core';
import { NgtCanvas } from 'angular-three/dom';
import { extend, NgtArgs } from 'angular-three';
import * as THREE from 'three';
import { NgtsOrbitControls } from 'angular-three-soba/controls';

import * as threejs from '../../../../threejs/src/public-api'


extend(THREE);

class HelixCurve extends THREE.Curve<THREE.Vector3> {
  constructor(public radius: number = 1, public height: number = 10, public turns: number = 5) {
    super();
  }
  override getPoint(t: number, optionalTarget = new THREE.Vector3()) {
    const angle = t * Math.PI * 2 * this.turns;
    const x = Math.cos(angle) * this.radius;
    const z = Math.sin(angle) * this.radius;
    const y = t * this.height;
    return optionalTarget.set(x, y, z);
  }
}

class UndulatingTorusCurve extends THREE.Curve<THREE.Vector3> {
  constructor(public radius: number = 10, public waves: number = 5, public amplitude: number = 2) {
    super();
  }
  override getPoint(t: number, optionalTarget = new THREE.Vector3()) {
    const angle = t * Math.PI * 2;
    const x = Math.cos(angle) * this.radius;
    const z = Math.sin(angle) * this.radius;
    const y = Math.sin(angle * this.waves) * this.amplitude;
    return optionalTarget.set(x, y, z);
  }
}

import { injectStore } from 'angular-three';

@Directive({
  selector: '[cameraUpdater]',
  standalone: true
})
export class CameraUpdaterDirective implements OnChanges {
  @Input('cameraUpdater') cam: any;
  
  private store = injectStore();
  
  ngOnChanges(changes: SimpleChanges) {
    this.updateCamera();
  }

  updateCamera() {
    if (this.cam) {
      const state = this.store();
      const camera = state.camera;
      const controls = state.controls as any;
      if (camera) {
        camera.position.set(this.cam.X ?? 15, this.cam.Y ?? 15, this.cam.Z ?? 15);
        camera.lookAt(this.cam.TargetX ?? 0, this.cam.TargetY ?? 0, this.cam.TargetZ ?? 0);
        
        // Use type assertion to set fov since state.camera is of type THREE.Camera 
        // which might not have 'fov' if it's not a PerspectiveCamera
        if ('fov' in camera) {
          (camera as any).fov = this.cam.Fov || 50;
        }

        camera.updateProjectionMatrix();
      }
      if (controls) {
        controls.target.set(this.cam.TargetX ?? 0, this.cam.TargetY ?? 0, this.cam.TargetZ ?? 0);
        controls.update();
      } else {
        // Retry for controls
        setTimeout(() => {
          const c = this.store().controls as any;
          if (c) {
            c.target.set(this.cam.TargetX ?? 0, this.cam.TargetY ?? 0, this.cam.TargetZ ?? 0);
            c.update();
          }
        }, 100);
      }
    }
  }
}

@Component({
  selector: 'lib-threejs-specific',
  templateUrl: './gongthreejs-specific.html',
  styleUrl: './gongthreejs-specific.css',
  imports: [NgtCanvas, NgtArgs, NgtsOrbitControls, CameraUpdaterDirective],
  schemas: [CUSTOM_ELEMENTS_SCHEMA],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class GongthreejsSpecific {

  @Input() Name: string = ""
  public frontRepo?: threejs.FrontRepo

  canvasSingloton?: threejs.Canvas


  constructor(
    private threejsFrontRepoService: threejs.FrontRepoService,
    private threejsCommitNbFromBackService: threejs.CommitNbFromBackService,
    private threejsPushFromFrontNbService: threejs.PushFromFrontNbService,
    private cdr: ChangeDetectorRef,
  ) {
  }

  getCurve(curve: threejs.Curve | undefined): THREE.Curve<THREE.Vector3> {
    if (!curve || !curve.Points || curve.Points.length === 0) {
      // return a default line curve if no points exist
      return new THREE.LineCurve3(new THREE.Vector3(0, 0, 0), new THREE.Vector3(0, 1, 0));
    }
    const points = curve.Points.map(v => new THREE.Vector3(v.X, v.Y, v.Z));
    return new THREE.CatmullRomCurve3(points);
  }

  getShape(shape: threejs.Shape | undefined): THREE.Shape {
    const s = new THREE.Shape();
    if (!shape || !shape.Points || shape.Points.length === 0) return s;
    s.moveTo(shape.Points[0].X, shape.Points[0].Y);
    for (let i = 1; i < shape.Points.length; i++) {
      s.lineTo(shape.Points[i].X, shape.Points[i].Y);
    }
    return s;
  }

  getExtrudeArgs(extrude: threejs.ExtrudeGeometry): any[] {
    const shape = this.getShape(extrude.Shape);
    const settings = {
      steps: extrude.Steps,
      extrudePath: this.getCurve(extrude.ExtrudePath)
    };
    return [shape, settings];
  }

  private bufferGeometryCache = new Map<number, THREE.BufferGeometry>();

  getBufferGeometryArgs(bufferGeometry: threejs.BufferGeometry): THREE.BufferGeometry {
    if (!bufferGeometry) return new THREE.BufferGeometry();

    if (this.bufferGeometryCache.has(bufferGeometry.ID)) {
      return this.bufferGeometryCache.get(bufferGeometry.ID)!;
    }

    const geometry = new THREE.BufferGeometry();

    if (bufferGeometry.Vertices && bufferGeometry.Vertices.length > 0) {
      // The backend preserves the append order of Vertices using IntSlice.
      // We must NOT sort by ID because IDs are assigned randomly during stage.Commit() map iteration!
      // Sorting by ID completely scrambles the array, which breaks the Triangle face indices (V1, V2, V3).
      const vertices = bufferGeometry.Vertices;
      const positions = new Float32Array(vertices.length * 3);
      
      vertices.forEach((v, i) => {
        positions[i * 3] = v.X;
        positions[i * 3 + 1] = v.Y;
        positions[i * 3 + 2] = v.Z;
      });
      geometry.setAttribute('position', new THREE.BufferAttribute(positions, 3));

      if (bufferGeometry.Faces && bufferGeometry.Faces.length > 0) {
        const indices = new Uint32Array(bufferGeometry.Faces.length * 3);
        bufferGeometry.Faces.forEach((f, i) => {
          indices[i * 3] = f.V1;
          indices[i * 3 + 1] = f.V2;
          indices[i * 3 + 2] = f.V3;
        });
        geometry.setIndex(new THREE.BufferAttribute(indices, 1));
      }
      
      // Calculate normals so that MeshPhysicalMaterial can render properly
      geometry.computeVertexNormals();
    }

    geometry.computeVertexNormals();
    this.bufferGeometryCache.set(bufferGeometry.ID, geometry);
    return geometry;
  }

  ngOnInit(): void {
    this.threejsFrontRepoService.connectToWebSocket(this.Name).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        var canvasSingloton: threejs.Canvas = new (threejs.Canvas)

        for (let instance of this.frontRepo.getFrontArray<threejs.Canvas>(threejs.Canvas.GONGSTRUCT_NAME)) {
          canvasSingloton = instance
        }

        this.canvasSingloton = canvasSingloton
        this.cdr.detectChanges()
      }
    )
  }

  readonly curve = new THREE.CubicBezierCurve3(
      new THREE.Vector3(-10, 0, 5),
      new THREE.Vector3(-5, 15, 5),
      new THREE.Vector3(5, -15, 5),
      new THREE.Vector3(10, 0, 5)
    );

  readonly helixCurve = new HelixCurve(2, 6, 4);

  readonly torusLikeCurve = new UndulatingTorusCurve(10, 5, 2);

  readonly squareShape = (() => {
    const shape = new THREE.Shape();
    const size = 0.75;
    shape.moveTo(-size, -size);
    shape.lineTo(size, -size);
    shape.lineTo(size, size);
    shape.lineTo(-size, size);
    shape.lineTo(-size, -size);
    return shape;
  })();

  readonly extrudeSettings = {
    extrudePath: this.torusLikeCurve,
    steps: 500,
  };
}
