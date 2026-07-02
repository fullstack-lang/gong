// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	AmbiantLightAPIs []*AmbiantLightAPI

	BoxGeometryAPIs []*BoxGeometryAPI

	CanvasAPIs []*CanvasAPI

	CurveAPIs []*CurveAPI

	CylinderGeometryAPIs []*CylinderGeometryAPI

	DirectionalLightAPIs []*DirectionalLightAPI

	ExtrudeGeometryAPIs []*ExtrudeGeometryAPI

	MeshAPIs []*MeshAPI

	MeshMaterialBasicAPIs []*MeshMaterialBasicAPI

	MeshPhysicalMaterialAPIs []*MeshPhysicalMaterialAPI

	PlaneGeometryAPIs []*PlaneGeometryAPI

	ShapeAPIs []*ShapeAPI

	SphereGeometryAPIs []*SphereGeometryAPI

	TorusGeometryAPIs []*TorusGeometryAPI

	TubeGeometryAPIs []*TubeGeometryAPI

	Vector2APIs []*Vector2API

	Vector3APIs []*Vector3API

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, ambiantlightDB := range backRepo.BackRepoAmbiantLight.Map_AmbiantLightDBID_AmbiantLightDB {

		var ambiantlightAPI AmbiantLightAPI
		ambiantlightAPI.ID = ambiantlightDB.ID
		ambiantlightAPI.AmbiantLightPointersEncoding = ambiantlightDB.AmbiantLightPointersEncoding
		ambiantlightDB.CopyBasicFieldsToAmbiantLight_WOP(&ambiantlightAPI.AmbiantLight_WOP)

		backRepoData.AmbiantLightAPIs = append(backRepoData.AmbiantLightAPIs, &ambiantlightAPI)
	}

	for _, boxgeometryDB := range backRepo.BackRepoBoxGeometry.Map_BoxGeometryDBID_BoxGeometryDB {

		var boxgeometryAPI BoxGeometryAPI
		boxgeometryAPI.ID = boxgeometryDB.ID
		boxgeometryAPI.BoxGeometryPointersEncoding = boxgeometryDB.BoxGeometryPointersEncoding
		boxgeometryDB.CopyBasicFieldsToBoxGeometry_WOP(&boxgeometryAPI.BoxGeometry_WOP)

		backRepoData.BoxGeometryAPIs = append(backRepoData.BoxGeometryAPIs, &boxgeometryAPI)
	}

	for _, canvasDB := range backRepo.BackRepoCanvas.Map_CanvasDBID_CanvasDB {

		var canvasAPI CanvasAPI
		canvasAPI.ID = canvasDB.ID
		canvasAPI.CanvasPointersEncoding = canvasDB.CanvasPointersEncoding
		canvasDB.CopyBasicFieldsToCanvas_WOP(&canvasAPI.Canvas_WOP)

		backRepoData.CanvasAPIs = append(backRepoData.CanvasAPIs, &canvasAPI)
	}

	for _, curveDB := range backRepo.BackRepoCurve.Map_CurveDBID_CurveDB {

		var curveAPI CurveAPI
		curveAPI.ID = curveDB.ID
		curveAPI.CurvePointersEncoding = curveDB.CurvePointersEncoding
		curveDB.CopyBasicFieldsToCurve_WOP(&curveAPI.Curve_WOP)

		backRepoData.CurveAPIs = append(backRepoData.CurveAPIs, &curveAPI)
	}

	for _, cylindergeometryDB := range backRepo.BackRepoCylinderGeometry.Map_CylinderGeometryDBID_CylinderGeometryDB {

		var cylindergeometryAPI CylinderGeometryAPI
		cylindergeometryAPI.ID = cylindergeometryDB.ID
		cylindergeometryAPI.CylinderGeometryPointersEncoding = cylindergeometryDB.CylinderGeometryPointersEncoding
		cylindergeometryDB.CopyBasicFieldsToCylinderGeometry_WOP(&cylindergeometryAPI.CylinderGeometry_WOP)

		backRepoData.CylinderGeometryAPIs = append(backRepoData.CylinderGeometryAPIs, &cylindergeometryAPI)
	}

	for _, directionallightDB := range backRepo.BackRepoDirectionalLight.Map_DirectionalLightDBID_DirectionalLightDB {

		var directionallightAPI DirectionalLightAPI
		directionallightAPI.ID = directionallightDB.ID
		directionallightAPI.DirectionalLightPointersEncoding = directionallightDB.DirectionalLightPointersEncoding
		directionallightDB.CopyBasicFieldsToDirectionalLight_WOP(&directionallightAPI.DirectionalLight_WOP)

		backRepoData.DirectionalLightAPIs = append(backRepoData.DirectionalLightAPIs, &directionallightAPI)
	}

	for _, extrudegeometryDB := range backRepo.BackRepoExtrudeGeometry.Map_ExtrudeGeometryDBID_ExtrudeGeometryDB {

		var extrudegeometryAPI ExtrudeGeometryAPI
		extrudegeometryAPI.ID = extrudegeometryDB.ID
		extrudegeometryAPI.ExtrudeGeometryPointersEncoding = extrudegeometryDB.ExtrudeGeometryPointersEncoding
		extrudegeometryDB.CopyBasicFieldsToExtrudeGeometry_WOP(&extrudegeometryAPI.ExtrudeGeometry_WOP)

		backRepoData.ExtrudeGeometryAPIs = append(backRepoData.ExtrudeGeometryAPIs, &extrudegeometryAPI)
	}

	for _, meshDB := range backRepo.BackRepoMesh.Map_MeshDBID_MeshDB {

		var meshAPI MeshAPI
		meshAPI.ID = meshDB.ID
		meshAPI.MeshPointersEncoding = meshDB.MeshPointersEncoding
		meshDB.CopyBasicFieldsToMesh_WOP(&meshAPI.Mesh_WOP)

		backRepoData.MeshAPIs = append(backRepoData.MeshAPIs, &meshAPI)
	}

	for _, meshmaterialbasicDB := range backRepo.BackRepoMeshMaterialBasic.Map_MeshMaterialBasicDBID_MeshMaterialBasicDB {

		var meshmaterialbasicAPI MeshMaterialBasicAPI
		meshmaterialbasicAPI.ID = meshmaterialbasicDB.ID
		meshmaterialbasicAPI.MeshMaterialBasicPointersEncoding = meshmaterialbasicDB.MeshMaterialBasicPointersEncoding
		meshmaterialbasicDB.CopyBasicFieldsToMeshMaterialBasic_WOP(&meshmaterialbasicAPI.MeshMaterialBasic_WOP)

		backRepoData.MeshMaterialBasicAPIs = append(backRepoData.MeshMaterialBasicAPIs, &meshmaterialbasicAPI)
	}

	for _, meshphysicalmaterialDB := range backRepo.BackRepoMeshPhysicalMaterial.Map_MeshPhysicalMaterialDBID_MeshPhysicalMaterialDB {

		var meshphysicalmaterialAPI MeshPhysicalMaterialAPI
		meshphysicalmaterialAPI.ID = meshphysicalmaterialDB.ID
		meshphysicalmaterialAPI.MeshPhysicalMaterialPointersEncoding = meshphysicalmaterialDB.MeshPhysicalMaterialPointersEncoding
		meshphysicalmaterialDB.CopyBasicFieldsToMeshPhysicalMaterial_WOP(&meshphysicalmaterialAPI.MeshPhysicalMaterial_WOP)

		backRepoData.MeshPhysicalMaterialAPIs = append(backRepoData.MeshPhysicalMaterialAPIs, &meshphysicalmaterialAPI)
	}

	for _, planegeometryDB := range backRepo.BackRepoPlaneGeometry.Map_PlaneGeometryDBID_PlaneGeometryDB {

		var planegeometryAPI PlaneGeometryAPI
		planegeometryAPI.ID = planegeometryDB.ID
		planegeometryAPI.PlaneGeometryPointersEncoding = planegeometryDB.PlaneGeometryPointersEncoding
		planegeometryDB.CopyBasicFieldsToPlaneGeometry_WOP(&planegeometryAPI.PlaneGeometry_WOP)

		backRepoData.PlaneGeometryAPIs = append(backRepoData.PlaneGeometryAPIs, &planegeometryAPI)
	}

	for _, shapeDB := range backRepo.BackRepoShape.Map_ShapeDBID_ShapeDB {

		var shapeAPI ShapeAPI
		shapeAPI.ID = shapeDB.ID
		shapeAPI.ShapePointersEncoding = shapeDB.ShapePointersEncoding
		shapeDB.CopyBasicFieldsToShape_WOP(&shapeAPI.Shape_WOP)

		backRepoData.ShapeAPIs = append(backRepoData.ShapeAPIs, &shapeAPI)
	}

	for _, spheregeometryDB := range backRepo.BackRepoSphereGeometry.Map_SphereGeometryDBID_SphereGeometryDB {

		var spheregeometryAPI SphereGeometryAPI
		spheregeometryAPI.ID = spheregeometryDB.ID
		spheregeometryAPI.SphereGeometryPointersEncoding = spheregeometryDB.SphereGeometryPointersEncoding
		spheregeometryDB.CopyBasicFieldsToSphereGeometry_WOP(&spheregeometryAPI.SphereGeometry_WOP)

		backRepoData.SphereGeometryAPIs = append(backRepoData.SphereGeometryAPIs, &spheregeometryAPI)
	}

	for _, torusgeometryDB := range backRepo.BackRepoTorusGeometry.Map_TorusGeometryDBID_TorusGeometryDB {

		var torusgeometryAPI TorusGeometryAPI
		torusgeometryAPI.ID = torusgeometryDB.ID
		torusgeometryAPI.TorusGeometryPointersEncoding = torusgeometryDB.TorusGeometryPointersEncoding
		torusgeometryDB.CopyBasicFieldsToTorusGeometry_WOP(&torusgeometryAPI.TorusGeometry_WOP)

		backRepoData.TorusGeometryAPIs = append(backRepoData.TorusGeometryAPIs, &torusgeometryAPI)
	}

	for _, tubegeometryDB := range backRepo.BackRepoTubeGeometry.Map_TubeGeometryDBID_TubeGeometryDB {

		var tubegeometryAPI TubeGeometryAPI
		tubegeometryAPI.ID = tubegeometryDB.ID
		tubegeometryAPI.TubeGeometryPointersEncoding = tubegeometryDB.TubeGeometryPointersEncoding
		tubegeometryDB.CopyBasicFieldsToTubeGeometry_WOP(&tubegeometryAPI.TubeGeometry_WOP)

		backRepoData.TubeGeometryAPIs = append(backRepoData.TubeGeometryAPIs, &tubegeometryAPI)
	}

	for _, vector2DB := range backRepo.BackRepoVector2.Map_Vector2DBID_Vector2DB {

		var vector2API Vector2API
		vector2API.ID = vector2DB.ID
		vector2API.Vector2PointersEncoding = vector2DB.Vector2PointersEncoding
		vector2DB.CopyBasicFieldsToVector2_WOP(&vector2API.Vector2_WOP)

		backRepoData.Vector2APIs = append(backRepoData.Vector2APIs, &vector2API)
	}

	for _, vector3DB := range backRepo.BackRepoVector3.Map_Vector3DBID_Vector3DB {

		var vector3API Vector3API
		vector3API.ID = vector3DB.ID
		vector3API.Vector3PointersEncoding = vector3DB.Vector3PointersEncoding
		vector3DB.CopyBasicFieldsToVector3_WOP(&vector3API.Vector3_WOP)

		backRepoData.Vector3APIs = append(backRepoData.Vector3APIs, &vector3API)
	}

}
