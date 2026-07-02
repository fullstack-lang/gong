// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AmbiantLight:
		if stage.OnAfterAmbiantLightCreateCallback != nil {
			stage.OnAfterAmbiantLightCreateCallback.OnAfterCreate(stage, target)
		}
	case *BoxGeometry:
		if stage.OnAfterBoxGeometryCreateCallback != nil {
			stage.OnAfterBoxGeometryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Canvas:
		if stage.OnAfterCanvasCreateCallback != nil {
			stage.OnAfterCanvasCreateCallback.OnAfterCreate(stage, target)
		}
	case *Curve:
		if stage.OnAfterCurveCreateCallback != nil {
			stage.OnAfterCurveCreateCallback.OnAfterCreate(stage, target)
		}
	case *CylinderGeometry:
		if stage.OnAfterCylinderGeometryCreateCallback != nil {
			stage.OnAfterCylinderGeometryCreateCallback.OnAfterCreate(stage, target)
		}
	case *DirectionalLight:
		if stage.OnAfterDirectionalLightCreateCallback != nil {
			stage.OnAfterDirectionalLightCreateCallback.OnAfterCreate(stage, target)
		}
	case *ExtrudeGeometry:
		if stage.OnAfterExtrudeGeometryCreateCallback != nil {
			stage.OnAfterExtrudeGeometryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Mesh:
		if stage.OnAfterMeshCreateCallback != nil {
			stage.OnAfterMeshCreateCallback.OnAfterCreate(stage, target)
		}
	case *MeshMaterialBasic:
		if stage.OnAfterMeshMaterialBasicCreateCallback != nil {
			stage.OnAfterMeshMaterialBasicCreateCallback.OnAfterCreate(stage, target)
		}
	case *MeshPhysicalMaterial:
		if stage.OnAfterMeshPhysicalMaterialCreateCallback != nil {
			stage.OnAfterMeshPhysicalMaterialCreateCallback.OnAfterCreate(stage, target)
		}
	case *PlaneGeometry:
		if stage.OnAfterPlaneGeometryCreateCallback != nil {
			stage.OnAfterPlaneGeometryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Shape:
		if stage.OnAfterShapeCreateCallback != nil {
			stage.OnAfterShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *SphereGeometry:
		if stage.OnAfterSphereGeometryCreateCallback != nil {
			stage.OnAfterSphereGeometryCreateCallback.OnAfterCreate(stage, target)
		}
	case *TorusGeometry:
		if stage.OnAfterTorusGeometryCreateCallback != nil {
			stage.OnAfterTorusGeometryCreateCallback.OnAfterCreate(stage, target)
		}
	case *TubeGeometry:
		if stage.OnAfterTubeGeometryCreateCallback != nil {
			stage.OnAfterTubeGeometryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Vector2:
		if stage.OnAfterVector2CreateCallback != nil {
			stage.OnAfterVector2CreateCallback.OnAfterCreate(stage, target)
		}
	case *Vector3:
		if stage.OnAfterVector3CreateCallback != nil {
			stage.OnAfterVector3CreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *AmbiantLight:
		newTarget := any(new).(*AmbiantLight)
		if stage.OnAfterAmbiantLightUpdateCallback != nil {
			stage.OnAfterAmbiantLightUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *BoxGeometry:
		newTarget := any(new).(*BoxGeometry)
		if stage.OnAfterBoxGeometryUpdateCallback != nil {
			stage.OnAfterBoxGeometryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Canvas:
		newTarget := any(new).(*Canvas)
		if stage.OnAfterCanvasUpdateCallback != nil {
			stage.OnAfterCanvasUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Curve:
		newTarget := any(new).(*Curve)
		if stage.OnAfterCurveUpdateCallback != nil {
			stage.OnAfterCurveUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *CylinderGeometry:
		newTarget := any(new).(*CylinderGeometry)
		if stage.OnAfterCylinderGeometryUpdateCallback != nil {
			stage.OnAfterCylinderGeometryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DirectionalLight:
		newTarget := any(new).(*DirectionalLight)
		if stage.OnAfterDirectionalLightUpdateCallback != nil {
			stage.OnAfterDirectionalLightUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ExtrudeGeometry:
		newTarget := any(new).(*ExtrudeGeometry)
		if stage.OnAfterExtrudeGeometryUpdateCallback != nil {
			stage.OnAfterExtrudeGeometryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Mesh:
		newTarget := any(new).(*Mesh)
		if stage.OnAfterMeshUpdateCallback != nil {
			stage.OnAfterMeshUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MeshMaterialBasic:
		newTarget := any(new).(*MeshMaterialBasic)
		if stage.OnAfterMeshMaterialBasicUpdateCallback != nil {
			stage.OnAfterMeshMaterialBasicUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MeshPhysicalMaterial:
		newTarget := any(new).(*MeshPhysicalMaterial)
		if stage.OnAfterMeshPhysicalMaterialUpdateCallback != nil {
			stage.OnAfterMeshPhysicalMaterialUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PlaneGeometry:
		newTarget := any(new).(*PlaneGeometry)
		if stage.OnAfterPlaneGeometryUpdateCallback != nil {
			stage.OnAfterPlaneGeometryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Shape:
		newTarget := any(new).(*Shape)
		if stage.OnAfterShapeUpdateCallback != nil {
			stage.OnAfterShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SphereGeometry:
		newTarget := any(new).(*SphereGeometry)
		if stage.OnAfterSphereGeometryUpdateCallback != nil {
			stage.OnAfterSphereGeometryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TorusGeometry:
		newTarget := any(new).(*TorusGeometry)
		if stage.OnAfterTorusGeometryUpdateCallback != nil {
			stage.OnAfterTorusGeometryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TubeGeometry:
		newTarget := any(new).(*TubeGeometry)
		if stage.OnAfterTubeGeometryUpdateCallback != nil {
			stage.OnAfterTubeGeometryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Vector2:
		newTarget := any(new).(*Vector2)
		if stage.OnAfterVector2UpdateCallback != nil {
			stage.OnAfterVector2UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Vector3:
		newTarget := any(new).(*Vector3)
		if stage.OnAfterVector3UpdateCallback != nil {
			stage.OnAfterVector3UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *AmbiantLight:
		if stage.OnAfterAmbiantLightDeleteCallback != nil {
			staged := any(staged).(*AmbiantLight)
			stage.OnAfterAmbiantLightDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *BoxGeometry:
		if stage.OnAfterBoxGeometryDeleteCallback != nil {
			staged := any(staged).(*BoxGeometry)
			stage.OnAfterBoxGeometryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Canvas:
		if stage.OnAfterCanvasDeleteCallback != nil {
			staged := any(staged).(*Canvas)
			stage.OnAfterCanvasDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Curve:
		if stage.OnAfterCurveDeleteCallback != nil {
			staged := any(staged).(*Curve)
			stage.OnAfterCurveDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *CylinderGeometry:
		if stage.OnAfterCylinderGeometryDeleteCallback != nil {
			staged := any(staged).(*CylinderGeometry)
			stage.OnAfterCylinderGeometryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DirectionalLight:
		if stage.OnAfterDirectionalLightDeleteCallback != nil {
			staged := any(staged).(*DirectionalLight)
			stage.OnAfterDirectionalLightDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ExtrudeGeometry:
		if stage.OnAfterExtrudeGeometryDeleteCallback != nil {
			staged := any(staged).(*ExtrudeGeometry)
			stage.OnAfterExtrudeGeometryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Mesh:
		if stage.OnAfterMeshDeleteCallback != nil {
			staged := any(staged).(*Mesh)
			stage.OnAfterMeshDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MeshMaterialBasic:
		if stage.OnAfterMeshMaterialBasicDeleteCallback != nil {
			staged := any(staged).(*MeshMaterialBasic)
			stage.OnAfterMeshMaterialBasicDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MeshPhysicalMaterial:
		if stage.OnAfterMeshPhysicalMaterialDeleteCallback != nil {
			staged := any(staged).(*MeshPhysicalMaterial)
			stage.OnAfterMeshPhysicalMaterialDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PlaneGeometry:
		if stage.OnAfterPlaneGeometryDeleteCallback != nil {
			staged := any(staged).(*PlaneGeometry)
			stage.OnAfterPlaneGeometryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Shape:
		if stage.OnAfterShapeDeleteCallback != nil {
			staged := any(staged).(*Shape)
			stage.OnAfterShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SphereGeometry:
		if stage.OnAfterSphereGeometryDeleteCallback != nil {
			staged := any(staged).(*SphereGeometry)
			stage.OnAfterSphereGeometryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TorusGeometry:
		if stage.OnAfterTorusGeometryDeleteCallback != nil {
			staged := any(staged).(*TorusGeometry)
			stage.OnAfterTorusGeometryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TubeGeometry:
		if stage.OnAfterTubeGeometryDeleteCallback != nil {
			staged := any(staged).(*TubeGeometry)
			stage.OnAfterTubeGeometryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Vector2:
		if stage.OnAfterVector2DeleteCallback != nil {
			staged := any(staged).(*Vector2)
			stage.OnAfterVector2DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Vector3:
		if stage.OnAfterVector3DeleteCallback != nil {
			staged := any(staged).(*Vector3)
			stage.OnAfterVector3DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AmbiantLight:
		if stage.OnAfterAmbiantLightReadCallback != nil {
			stage.OnAfterAmbiantLightReadCallback.OnAfterRead(stage, target)
		}
	case *BoxGeometry:
		if stage.OnAfterBoxGeometryReadCallback != nil {
			stage.OnAfterBoxGeometryReadCallback.OnAfterRead(stage, target)
		}
	case *Canvas:
		if stage.OnAfterCanvasReadCallback != nil {
			stage.OnAfterCanvasReadCallback.OnAfterRead(stage, target)
		}
	case *Curve:
		if stage.OnAfterCurveReadCallback != nil {
			stage.OnAfterCurveReadCallback.OnAfterRead(stage, target)
		}
	case *CylinderGeometry:
		if stage.OnAfterCylinderGeometryReadCallback != nil {
			stage.OnAfterCylinderGeometryReadCallback.OnAfterRead(stage, target)
		}
	case *DirectionalLight:
		if stage.OnAfterDirectionalLightReadCallback != nil {
			stage.OnAfterDirectionalLightReadCallback.OnAfterRead(stage, target)
		}
	case *ExtrudeGeometry:
		if stage.OnAfterExtrudeGeometryReadCallback != nil {
			stage.OnAfterExtrudeGeometryReadCallback.OnAfterRead(stage, target)
		}
	case *Mesh:
		if stage.OnAfterMeshReadCallback != nil {
			stage.OnAfterMeshReadCallback.OnAfterRead(stage, target)
		}
	case *MeshMaterialBasic:
		if stage.OnAfterMeshMaterialBasicReadCallback != nil {
			stage.OnAfterMeshMaterialBasicReadCallback.OnAfterRead(stage, target)
		}
	case *MeshPhysicalMaterial:
		if stage.OnAfterMeshPhysicalMaterialReadCallback != nil {
			stage.OnAfterMeshPhysicalMaterialReadCallback.OnAfterRead(stage, target)
		}
	case *PlaneGeometry:
		if stage.OnAfterPlaneGeometryReadCallback != nil {
			stage.OnAfterPlaneGeometryReadCallback.OnAfterRead(stage, target)
		}
	case *Shape:
		if stage.OnAfterShapeReadCallback != nil {
			stage.OnAfterShapeReadCallback.OnAfterRead(stage, target)
		}
	case *SphereGeometry:
		if stage.OnAfterSphereGeometryReadCallback != nil {
			stage.OnAfterSphereGeometryReadCallback.OnAfterRead(stage, target)
		}
	case *TorusGeometry:
		if stage.OnAfterTorusGeometryReadCallback != nil {
			stage.OnAfterTorusGeometryReadCallback.OnAfterRead(stage, target)
		}
	case *TubeGeometry:
		if stage.OnAfterTubeGeometryReadCallback != nil {
			stage.OnAfterTubeGeometryReadCallback.OnAfterRead(stage, target)
		}
	case *Vector2:
		if stage.OnAfterVector2ReadCallback != nil {
			stage.OnAfterVector2ReadCallback.OnAfterRead(stage, target)
		}
	case *Vector3:
		if stage.OnAfterVector3ReadCallback != nil {
			stage.OnAfterVector3ReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *Stage, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AmbiantLight:
		stage.OnAfterAmbiantLightUpdateCallback = any(callback).(OnAfterUpdateInterface[AmbiantLight])
	case *BoxGeometry:
		stage.OnAfterBoxGeometryUpdateCallback = any(callback).(OnAfterUpdateInterface[BoxGeometry])
	case *Canvas:
		stage.OnAfterCanvasUpdateCallback = any(callback).(OnAfterUpdateInterface[Canvas])
	case *Curve:
		stage.OnAfterCurveUpdateCallback = any(callback).(OnAfterUpdateInterface[Curve])
	case *CylinderGeometry:
		stage.OnAfterCylinderGeometryUpdateCallback = any(callback).(OnAfterUpdateInterface[CylinderGeometry])
	case *DirectionalLight:
		stage.OnAfterDirectionalLightUpdateCallback = any(callback).(OnAfterUpdateInterface[DirectionalLight])
	case *ExtrudeGeometry:
		stage.OnAfterExtrudeGeometryUpdateCallback = any(callback).(OnAfterUpdateInterface[ExtrudeGeometry])
	case *Mesh:
		stage.OnAfterMeshUpdateCallback = any(callback).(OnAfterUpdateInterface[Mesh])
	case *MeshMaterialBasic:
		stage.OnAfterMeshMaterialBasicUpdateCallback = any(callback).(OnAfterUpdateInterface[MeshMaterialBasic])
	case *MeshPhysicalMaterial:
		stage.OnAfterMeshPhysicalMaterialUpdateCallback = any(callback).(OnAfterUpdateInterface[MeshPhysicalMaterial])
	case *PlaneGeometry:
		stage.OnAfterPlaneGeometryUpdateCallback = any(callback).(OnAfterUpdateInterface[PlaneGeometry])
	case *Shape:
		stage.OnAfterShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[Shape])
	case *SphereGeometry:
		stage.OnAfterSphereGeometryUpdateCallback = any(callback).(OnAfterUpdateInterface[SphereGeometry])
	case *TorusGeometry:
		stage.OnAfterTorusGeometryUpdateCallback = any(callback).(OnAfterUpdateInterface[TorusGeometry])
	case *TubeGeometry:
		stage.OnAfterTubeGeometryUpdateCallback = any(callback).(OnAfterUpdateInterface[TubeGeometry])
	case *Vector2:
		stage.OnAfterVector2UpdateCallback = any(callback).(OnAfterUpdateInterface[Vector2])
	case *Vector3:
		stage.OnAfterVector3UpdateCallback = any(callback).(OnAfterUpdateInterface[Vector3])
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AmbiantLight:
		stage.OnAfterAmbiantLightCreateCallback = any(callback).(OnAfterCreateInterface[AmbiantLight])
	case *BoxGeometry:
		stage.OnAfterBoxGeometryCreateCallback = any(callback).(OnAfterCreateInterface[BoxGeometry])
	case *Canvas:
		stage.OnAfterCanvasCreateCallback = any(callback).(OnAfterCreateInterface[Canvas])
	case *Curve:
		stage.OnAfterCurveCreateCallback = any(callback).(OnAfterCreateInterface[Curve])
	case *CylinderGeometry:
		stage.OnAfterCylinderGeometryCreateCallback = any(callback).(OnAfterCreateInterface[CylinderGeometry])
	case *DirectionalLight:
		stage.OnAfterDirectionalLightCreateCallback = any(callback).(OnAfterCreateInterface[DirectionalLight])
	case *ExtrudeGeometry:
		stage.OnAfterExtrudeGeometryCreateCallback = any(callback).(OnAfterCreateInterface[ExtrudeGeometry])
	case *Mesh:
		stage.OnAfterMeshCreateCallback = any(callback).(OnAfterCreateInterface[Mesh])
	case *MeshMaterialBasic:
		stage.OnAfterMeshMaterialBasicCreateCallback = any(callback).(OnAfterCreateInterface[MeshMaterialBasic])
	case *MeshPhysicalMaterial:
		stage.OnAfterMeshPhysicalMaterialCreateCallback = any(callback).(OnAfterCreateInterface[MeshPhysicalMaterial])
	case *PlaneGeometry:
		stage.OnAfterPlaneGeometryCreateCallback = any(callback).(OnAfterCreateInterface[PlaneGeometry])
	case *Shape:
		stage.OnAfterShapeCreateCallback = any(callback).(OnAfterCreateInterface[Shape])
	case *SphereGeometry:
		stage.OnAfterSphereGeometryCreateCallback = any(callback).(OnAfterCreateInterface[SphereGeometry])
	case *TorusGeometry:
		stage.OnAfterTorusGeometryCreateCallback = any(callback).(OnAfterCreateInterface[TorusGeometry])
	case *TubeGeometry:
		stage.OnAfterTubeGeometryCreateCallback = any(callback).(OnAfterCreateInterface[TubeGeometry])
	case *Vector2:
		stage.OnAfterVector2CreateCallback = any(callback).(OnAfterCreateInterface[Vector2])
	case *Vector3:
		stage.OnAfterVector3CreateCallback = any(callback).(OnAfterCreateInterface[Vector3])
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AmbiantLight:
		stage.OnAfterAmbiantLightDeleteCallback = any(callback).(OnAfterDeleteInterface[AmbiantLight])
	case *BoxGeometry:
		stage.OnAfterBoxGeometryDeleteCallback = any(callback).(OnAfterDeleteInterface[BoxGeometry])
	case *Canvas:
		stage.OnAfterCanvasDeleteCallback = any(callback).(OnAfterDeleteInterface[Canvas])
	case *Curve:
		stage.OnAfterCurveDeleteCallback = any(callback).(OnAfterDeleteInterface[Curve])
	case *CylinderGeometry:
		stage.OnAfterCylinderGeometryDeleteCallback = any(callback).(OnAfterDeleteInterface[CylinderGeometry])
	case *DirectionalLight:
		stage.OnAfterDirectionalLightDeleteCallback = any(callback).(OnAfterDeleteInterface[DirectionalLight])
	case *ExtrudeGeometry:
		stage.OnAfterExtrudeGeometryDeleteCallback = any(callback).(OnAfterDeleteInterface[ExtrudeGeometry])
	case *Mesh:
		stage.OnAfterMeshDeleteCallback = any(callback).(OnAfterDeleteInterface[Mesh])
	case *MeshMaterialBasic:
		stage.OnAfterMeshMaterialBasicDeleteCallback = any(callback).(OnAfterDeleteInterface[MeshMaterialBasic])
	case *MeshPhysicalMaterial:
		stage.OnAfterMeshPhysicalMaterialDeleteCallback = any(callback).(OnAfterDeleteInterface[MeshPhysicalMaterial])
	case *PlaneGeometry:
		stage.OnAfterPlaneGeometryDeleteCallback = any(callback).(OnAfterDeleteInterface[PlaneGeometry])
	case *Shape:
		stage.OnAfterShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[Shape])
	case *SphereGeometry:
		stage.OnAfterSphereGeometryDeleteCallback = any(callback).(OnAfterDeleteInterface[SphereGeometry])
	case *TorusGeometry:
		stage.OnAfterTorusGeometryDeleteCallback = any(callback).(OnAfterDeleteInterface[TorusGeometry])
	case *TubeGeometry:
		stage.OnAfterTubeGeometryDeleteCallback = any(callback).(OnAfterDeleteInterface[TubeGeometry])
	case *Vector2:
		stage.OnAfterVector2DeleteCallback = any(callback).(OnAfterDeleteInterface[Vector2])
	case *Vector3:
		stage.OnAfterVector3DeleteCallback = any(callback).(OnAfterDeleteInterface[Vector3])
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AmbiantLight:
		stage.OnAfterAmbiantLightReadCallback = any(callback).(OnAfterReadInterface[AmbiantLight])
	case *BoxGeometry:
		stage.OnAfterBoxGeometryReadCallback = any(callback).(OnAfterReadInterface[BoxGeometry])
	case *Canvas:
		stage.OnAfterCanvasReadCallback = any(callback).(OnAfterReadInterface[Canvas])
	case *Curve:
		stage.OnAfterCurveReadCallback = any(callback).(OnAfterReadInterface[Curve])
	case *CylinderGeometry:
		stage.OnAfterCylinderGeometryReadCallback = any(callback).(OnAfterReadInterface[CylinderGeometry])
	case *DirectionalLight:
		stage.OnAfterDirectionalLightReadCallback = any(callback).(OnAfterReadInterface[DirectionalLight])
	case *ExtrudeGeometry:
		stage.OnAfterExtrudeGeometryReadCallback = any(callback).(OnAfterReadInterface[ExtrudeGeometry])
	case *Mesh:
		stage.OnAfterMeshReadCallback = any(callback).(OnAfterReadInterface[Mesh])
	case *MeshMaterialBasic:
		stage.OnAfterMeshMaterialBasicReadCallback = any(callback).(OnAfterReadInterface[MeshMaterialBasic])
	case *MeshPhysicalMaterial:
		stage.OnAfterMeshPhysicalMaterialReadCallback = any(callback).(OnAfterReadInterface[MeshPhysicalMaterial])
	case *PlaneGeometry:
		stage.OnAfterPlaneGeometryReadCallback = any(callback).(OnAfterReadInterface[PlaneGeometry])
	case *Shape:
		stage.OnAfterShapeReadCallback = any(callback).(OnAfterReadInterface[Shape])
	case *SphereGeometry:
		stage.OnAfterSphereGeometryReadCallback = any(callback).(OnAfterReadInterface[SphereGeometry])
	case *TorusGeometry:
		stage.OnAfterTorusGeometryReadCallback = any(callback).(OnAfterReadInterface[TorusGeometry])
	case *TubeGeometry:
		stage.OnAfterTubeGeometryReadCallback = any(callback).(OnAfterReadInterface[TubeGeometry])
	case *Vector2:
		stage.OnAfterVector2ReadCallback = any(callback).(OnAfterReadInterface[Vector2])
	case *Vector3:
		stage.OnAfterVector3ReadCallback = any(callback).(OnAfterReadInterface[Vector3])
	}
}
