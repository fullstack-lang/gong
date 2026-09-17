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
func (ambiantlight *AmbiantLight) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAmbiantLightCreateCallback != nil {
		stage.OnAfterAmbiantLightCreateCallback.OnAfterCreate(stage, ambiantlight)
	}
}

func (ambiantlight *AmbiantLight) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAmbiantLightUpdateCallback != nil {
		var frontAmbiantLight *AmbiantLight
		if front != nil {
			frontAmbiantLight, _ = front.(*AmbiantLight)
		}
		stage.OnAfterAmbiantLightUpdateCallback.OnAfterUpdate(stage, ambiantlight, frontAmbiantLight)
	}
}

func (ambiantlight *AmbiantLight) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAmbiantLightDeleteCallback != nil {
		var frontAmbiantLight *AmbiantLight
		if front != nil {
			frontAmbiantLight, _ = front.(*AmbiantLight)
		}
		stage.OnAfterAmbiantLightDeleteCallback.OnAfterDelete(stage, ambiantlight, frontAmbiantLight)
	}
}

func (boxgeometry *BoxGeometry) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBoxGeometryCreateCallback != nil {
		stage.OnAfterBoxGeometryCreateCallback.OnAfterCreate(stage, boxgeometry)
	}
}

func (boxgeometry *BoxGeometry) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBoxGeometryUpdateCallback != nil {
		var frontBoxGeometry *BoxGeometry
		if front != nil {
			frontBoxGeometry, _ = front.(*BoxGeometry)
		}
		stage.OnAfterBoxGeometryUpdateCallback.OnAfterUpdate(stage, boxgeometry, frontBoxGeometry)
	}
}

func (boxgeometry *BoxGeometry) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBoxGeometryDeleteCallback != nil {
		var frontBoxGeometry *BoxGeometry
		if front != nil {
			frontBoxGeometry, _ = front.(*BoxGeometry)
		}
		stage.OnAfterBoxGeometryDeleteCallback.OnAfterDelete(stage, boxgeometry, frontBoxGeometry)
	}
}

func (buffergeometry *BufferGeometry) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBufferGeometryCreateCallback != nil {
		stage.OnAfterBufferGeometryCreateCallback.OnAfterCreate(stage, buffergeometry)
	}
}

func (buffergeometry *BufferGeometry) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBufferGeometryUpdateCallback != nil {
		var frontBufferGeometry *BufferGeometry
		if front != nil {
			frontBufferGeometry, _ = front.(*BufferGeometry)
		}
		stage.OnAfterBufferGeometryUpdateCallback.OnAfterUpdate(stage, buffergeometry, frontBufferGeometry)
	}
}

func (buffergeometry *BufferGeometry) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBufferGeometryDeleteCallback != nil {
		var frontBufferGeometry *BufferGeometry
		if front != nil {
			frontBufferGeometry, _ = front.(*BufferGeometry)
		}
		stage.OnAfterBufferGeometryDeleteCallback.OnAfterDelete(stage, buffergeometry, frontBufferGeometry)
	}
}

func (camera *Camera) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCameraCreateCallback != nil {
		stage.OnAfterCameraCreateCallback.OnAfterCreate(stage, camera)
	}
}

func (camera *Camera) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCameraUpdateCallback != nil {
		var frontCamera *Camera
		if front != nil {
			frontCamera, _ = front.(*Camera)
		}
		stage.OnAfterCameraUpdateCallback.OnAfterUpdate(stage, camera, frontCamera)
	}
}

func (camera *Camera) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCameraDeleteCallback != nil {
		var frontCamera *Camera
		if front != nil {
			frontCamera, _ = front.(*Camera)
		}
		stage.OnAfterCameraDeleteCallback.OnAfterDelete(stage, camera, frontCamera)
	}
}

func (canvas *Canvas) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCanvasCreateCallback != nil {
		stage.OnAfterCanvasCreateCallback.OnAfterCreate(stage, canvas)
	}
}

func (canvas *Canvas) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCanvasUpdateCallback != nil {
		var frontCanvas *Canvas
		if front != nil {
			frontCanvas, _ = front.(*Canvas)
		}
		stage.OnAfterCanvasUpdateCallback.OnAfterUpdate(stage, canvas, frontCanvas)
	}
}

func (canvas *Canvas) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCanvasDeleteCallback != nil {
		var frontCanvas *Canvas
		if front != nil {
			frontCanvas, _ = front.(*Canvas)
		}
		stage.OnAfterCanvasDeleteCallback.OnAfterDelete(stage, canvas, frontCanvas)
	}
}

func (curve *Curve) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCurveCreateCallback != nil {
		stage.OnAfterCurveCreateCallback.OnAfterCreate(stage, curve)
	}
}

func (curve *Curve) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurveUpdateCallback != nil {
		var frontCurve *Curve
		if front != nil {
			frontCurve, _ = front.(*Curve)
		}
		stage.OnAfterCurveUpdateCallback.OnAfterUpdate(stage, curve, frontCurve)
	}
}

func (curve *Curve) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurveDeleteCallback != nil {
		var frontCurve *Curve
		if front != nil {
			frontCurve, _ = front.(*Curve)
		}
		stage.OnAfterCurveDeleteCallback.OnAfterDelete(stage, curve, frontCurve)
	}
}

func (cylindergeometry *CylinderGeometry) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCylinderGeometryCreateCallback != nil {
		stage.OnAfterCylinderGeometryCreateCallback.OnAfterCreate(stage, cylindergeometry)
	}
}

func (cylindergeometry *CylinderGeometry) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCylinderGeometryUpdateCallback != nil {
		var frontCylinderGeometry *CylinderGeometry
		if front != nil {
			frontCylinderGeometry, _ = front.(*CylinderGeometry)
		}
		stage.OnAfterCylinderGeometryUpdateCallback.OnAfterUpdate(stage, cylindergeometry, frontCylinderGeometry)
	}
}

func (cylindergeometry *CylinderGeometry) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCylinderGeometryDeleteCallback != nil {
		var frontCylinderGeometry *CylinderGeometry
		if front != nil {
			frontCylinderGeometry, _ = front.(*CylinderGeometry)
		}
		stage.OnAfterCylinderGeometryDeleteCallback.OnAfterDelete(stage, cylindergeometry, frontCylinderGeometry)
	}
}

func (directionallight *DirectionalLight) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectionalLightCreateCallback != nil {
		stage.OnAfterDirectionalLightCreateCallback.OnAfterCreate(stage, directionallight)
	}
}

func (directionallight *DirectionalLight) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectionalLightUpdateCallback != nil {
		var frontDirectionalLight *DirectionalLight
		if front != nil {
			frontDirectionalLight, _ = front.(*DirectionalLight)
		}
		stage.OnAfterDirectionalLightUpdateCallback.OnAfterUpdate(stage, directionallight, frontDirectionalLight)
	}
}

func (directionallight *DirectionalLight) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectionalLightDeleteCallback != nil {
		var frontDirectionalLight *DirectionalLight
		if front != nil {
			frontDirectionalLight, _ = front.(*DirectionalLight)
		}
		stage.OnAfterDirectionalLightDeleteCallback.OnAfterDelete(stage, directionallight, frontDirectionalLight)
	}
}

func (extrudegeometry *ExtrudeGeometry) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterExtrudeGeometryCreateCallback != nil {
		stage.OnAfterExtrudeGeometryCreateCallback.OnAfterCreate(stage, extrudegeometry)
	}
}

func (extrudegeometry *ExtrudeGeometry) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterExtrudeGeometryUpdateCallback != nil {
		var frontExtrudeGeometry *ExtrudeGeometry
		if front != nil {
			frontExtrudeGeometry, _ = front.(*ExtrudeGeometry)
		}
		stage.OnAfterExtrudeGeometryUpdateCallback.OnAfterUpdate(stage, extrudegeometry, frontExtrudeGeometry)
	}
}

func (extrudegeometry *ExtrudeGeometry) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterExtrudeGeometryDeleteCallback != nil {
		var frontExtrudeGeometry *ExtrudeGeometry
		if front != nil {
			frontExtrudeGeometry, _ = front.(*ExtrudeGeometry)
		}
		stage.OnAfterExtrudeGeometryDeleteCallback.OnAfterDelete(stage, extrudegeometry, frontExtrudeGeometry)
	}
}

func (mesh *Mesh) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMeshCreateCallback != nil {
		stage.OnAfterMeshCreateCallback.OnAfterCreate(stage, mesh)
	}
}

func (mesh *Mesh) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeshUpdateCallback != nil {
		var frontMesh *Mesh
		if front != nil {
			frontMesh, _ = front.(*Mesh)
		}
		stage.OnAfterMeshUpdateCallback.OnAfterUpdate(stage, mesh, frontMesh)
	}
}

func (mesh *Mesh) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeshDeleteCallback != nil {
		var frontMesh *Mesh
		if front != nil {
			frontMesh, _ = front.(*Mesh)
		}
		stage.OnAfterMeshDeleteCallback.OnAfterDelete(stage, mesh, frontMesh)
	}
}

func (meshmaterialbasic *MeshMaterialBasic) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMeshMaterialBasicCreateCallback != nil {
		stage.OnAfterMeshMaterialBasicCreateCallback.OnAfterCreate(stage, meshmaterialbasic)
	}
}

func (meshmaterialbasic *MeshMaterialBasic) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeshMaterialBasicUpdateCallback != nil {
		var frontMeshMaterialBasic *MeshMaterialBasic
		if front != nil {
			frontMeshMaterialBasic, _ = front.(*MeshMaterialBasic)
		}
		stage.OnAfterMeshMaterialBasicUpdateCallback.OnAfterUpdate(stage, meshmaterialbasic, frontMeshMaterialBasic)
	}
}

func (meshmaterialbasic *MeshMaterialBasic) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeshMaterialBasicDeleteCallback != nil {
		var frontMeshMaterialBasic *MeshMaterialBasic
		if front != nil {
			frontMeshMaterialBasic, _ = front.(*MeshMaterialBasic)
		}
		stage.OnAfterMeshMaterialBasicDeleteCallback.OnAfterDelete(stage, meshmaterialbasic, frontMeshMaterialBasic)
	}
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMeshPhysicalMaterialCreateCallback != nil {
		stage.OnAfterMeshPhysicalMaterialCreateCallback.OnAfterCreate(stage, meshphysicalmaterial)
	}
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeshPhysicalMaterialUpdateCallback != nil {
		var frontMeshPhysicalMaterial *MeshPhysicalMaterial
		if front != nil {
			frontMeshPhysicalMaterial, _ = front.(*MeshPhysicalMaterial)
		}
		stage.OnAfterMeshPhysicalMaterialUpdateCallback.OnAfterUpdate(stage, meshphysicalmaterial, frontMeshPhysicalMaterial)
	}
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeshPhysicalMaterialDeleteCallback != nil {
		var frontMeshPhysicalMaterial *MeshPhysicalMaterial
		if front != nil {
			frontMeshPhysicalMaterial, _ = front.(*MeshPhysicalMaterial)
		}
		stage.OnAfterMeshPhysicalMaterialDeleteCallback.OnAfterDelete(stage, meshphysicalmaterial, frontMeshPhysicalMaterial)
	}
}

func (planegeometry *PlaneGeometry) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPlaneGeometryCreateCallback != nil {
		stage.OnAfterPlaneGeometryCreateCallback.OnAfterCreate(stage, planegeometry)
	}
}

func (planegeometry *PlaneGeometry) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlaneGeometryUpdateCallback != nil {
		var frontPlaneGeometry *PlaneGeometry
		if front != nil {
			frontPlaneGeometry, _ = front.(*PlaneGeometry)
		}
		stage.OnAfterPlaneGeometryUpdateCallback.OnAfterUpdate(stage, planegeometry, frontPlaneGeometry)
	}
}

func (planegeometry *PlaneGeometry) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlaneGeometryDeleteCallback != nil {
		var frontPlaneGeometry *PlaneGeometry
		if front != nil {
			frontPlaneGeometry, _ = front.(*PlaneGeometry)
		}
		stage.OnAfterPlaneGeometryDeleteCallback.OnAfterDelete(stage, planegeometry, frontPlaneGeometry)
	}
}

func (shape *Shape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterShapeCreateCallback != nil {
		stage.OnAfterShapeCreateCallback.OnAfterCreate(stage, shape)
	}
}

func (shape *Shape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShapeUpdateCallback != nil {
		var frontShape *Shape
		if front != nil {
			frontShape, _ = front.(*Shape)
		}
		stage.OnAfterShapeUpdateCallback.OnAfterUpdate(stage, shape, frontShape)
	}
}

func (shape *Shape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShapeDeleteCallback != nil {
		var frontShape *Shape
		if front != nil {
			frontShape, _ = front.(*Shape)
		}
		stage.OnAfterShapeDeleteCallback.OnAfterDelete(stage, shape, frontShape)
	}
}

func (spheregeometry *SphereGeometry) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSphereGeometryCreateCallback != nil {
		stage.OnAfterSphereGeometryCreateCallback.OnAfterCreate(stage, spheregeometry)
	}
}

func (spheregeometry *SphereGeometry) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSphereGeometryUpdateCallback != nil {
		var frontSphereGeometry *SphereGeometry
		if front != nil {
			frontSphereGeometry, _ = front.(*SphereGeometry)
		}
		stage.OnAfterSphereGeometryUpdateCallback.OnAfterUpdate(stage, spheregeometry, frontSphereGeometry)
	}
}

func (spheregeometry *SphereGeometry) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSphereGeometryDeleteCallback != nil {
		var frontSphereGeometry *SphereGeometry
		if front != nil {
			frontSphereGeometry, _ = front.(*SphereGeometry)
		}
		stage.OnAfterSphereGeometryDeleteCallback.OnAfterDelete(stage, spheregeometry, frontSphereGeometry)
	}
}

func (torusgeometry *TorusGeometry) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTorusGeometryCreateCallback != nil {
		stage.OnAfterTorusGeometryCreateCallback.OnAfterCreate(stage, torusgeometry)
	}
}

func (torusgeometry *TorusGeometry) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTorusGeometryUpdateCallback != nil {
		var frontTorusGeometry *TorusGeometry
		if front != nil {
			frontTorusGeometry, _ = front.(*TorusGeometry)
		}
		stage.OnAfterTorusGeometryUpdateCallback.OnAfterUpdate(stage, torusgeometry, frontTorusGeometry)
	}
}

func (torusgeometry *TorusGeometry) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTorusGeometryDeleteCallback != nil {
		var frontTorusGeometry *TorusGeometry
		if front != nil {
			frontTorusGeometry, _ = front.(*TorusGeometry)
		}
		stage.OnAfterTorusGeometryDeleteCallback.OnAfterDelete(stage, torusgeometry, frontTorusGeometry)
	}
}

func (triangle *Triangle) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTriangleCreateCallback != nil {
		stage.OnAfterTriangleCreateCallback.OnAfterCreate(stage, triangle)
	}
}

func (triangle *Triangle) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTriangleUpdateCallback != nil {
		var frontTriangle *Triangle
		if front != nil {
			frontTriangle, _ = front.(*Triangle)
		}
		stage.OnAfterTriangleUpdateCallback.OnAfterUpdate(stage, triangle, frontTriangle)
	}
}

func (triangle *Triangle) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTriangleDeleteCallback != nil {
		var frontTriangle *Triangle
		if front != nil {
			frontTriangle, _ = front.(*Triangle)
		}
		stage.OnAfterTriangleDeleteCallback.OnAfterDelete(stage, triangle, frontTriangle)
	}
}

func (tubegeometry *TubeGeometry) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTubeGeometryCreateCallback != nil {
		stage.OnAfterTubeGeometryCreateCallback.OnAfterCreate(stage, tubegeometry)
	}
}

func (tubegeometry *TubeGeometry) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTubeGeometryUpdateCallback != nil {
		var frontTubeGeometry *TubeGeometry
		if front != nil {
			frontTubeGeometry, _ = front.(*TubeGeometry)
		}
		stage.OnAfterTubeGeometryUpdateCallback.OnAfterUpdate(stage, tubegeometry, frontTubeGeometry)
	}
}

func (tubegeometry *TubeGeometry) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTubeGeometryDeleteCallback != nil {
		var frontTubeGeometry *TubeGeometry
		if front != nil {
			frontTubeGeometry, _ = front.(*TubeGeometry)
		}
		stage.OnAfterTubeGeometryDeleteCallback.OnAfterDelete(stage, tubegeometry, frontTubeGeometry)
	}
}

func (vector2 *Vector2) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVector2CreateCallback != nil {
		stage.OnAfterVector2CreateCallback.OnAfterCreate(stage, vector2)
	}
}

func (vector2 *Vector2) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVector2UpdateCallback != nil {
		var frontVector2 *Vector2
		if front != nil {
			frontVector2, _ = front.(*Vector2)
		}
		stage.OnAfterVector2UpdateCallback.OnAfterUpdate(stage, vector2, frontVector2)
	}
}

func (vector2 *Vector2) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVector2DeleteCallback != nil {
		var frontVector2 *Vector2
		if front != nil {
			frontVector2, _ = front.(*Vector2)
		}
		stage.OnAfterVector2DeleteCallback.OnAfterDelete(stage, vector2, frontVector2)
	}
}

func (vector3 *Vector3) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVector3CreateCallback != nil {
		stage.OnAfterVector3CreateCallback.OnAfterCreate(stage, vector3)
	}
}

func (vector3 *Vector3) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVector3UpdateCallback != nil {
		var frontVector3 *Vector3
		if front != nil {
			frontVector3, _ = front.(*Vector3)
		}
		stage.OnAfterVector3UpdateCallback.OnAfterUpdate(stage, vector3, frontVector3)
	}
}

func (vector3 *Vector3) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVector3DeleteCallback != nil {
		var frontVector3 *Vector3
		if front != nil {
			frontVector3, _ = front.(*Vector3)
		}
		stage.OnAfterVector3DeleteCallback.OnAfterDelete(stage, vector3, frontVector3)
	}
}

