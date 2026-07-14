// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gong/lib/threejs/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	ambiantlightDBs map[uint]*AmbiantLightDB

	nextIDAmbiantLightDB uint

	boxgeometryDBs map[uint]*BoxGeometryDB

	nextIDBoxGeometryDB uint

	buffergeometryDBs map[uint]*BufferGeometryDB

	nextIDBufferGeometryDB uint

	cameraDBs map[uint]*CameraDB

	nextIDCameraDB uint

	canvasDBs map[uint]*CanvasDB

	nextIDCanvasDB uint

	curveDBs map[uint]*CurveDB

	nextIDCurveDB uint

	cylindergeometryDBs map[uint]*CylinderGeometryDB

	nextIDCylinderGeometryDB uint

	directionallightDBs map[uint]*DirectionalLightDB

	nextIDDirectionalLightDB uint

	extrudegeometryDBs map[uint]*ExtrudeGeometryDB

	nextIDExtrudeGeometryDB uint

	meshDBs map[uint]*MeshDB

	nextIDMeshDB uint

	meshmaterialbasicDBs map[uint]*MeshMaterialBasicDB

	nextIDMeshMaterialBasicDB uint

	meshphysicalmaterialDBs map[uint]*MeshPhysicalMaterialDB

	nextIDMeshPhysicalMaterialDB uint

	planegeometryDBs map[uint]*PlaneGeometryDB

	nextIDPlaneGeometryDB uint

	shapeDBs map[uint]*ShapeDB

	nextIDShapeDB uint

	spheregeometryDBs map[uint]*SphereGeometryDB

	nextIDSphereGeometryDB uint

	torusgeometryDBs map[uint]*TorusGeometryDB

	nextIDTorusGeometryDB uint

	triangleDBs map[uint]*TriangleDB

	nextIDTriangleDB uint

	tubegeometryDBs map[uint]*TubeGeometryDB

	nextIDTubeGeometryDB uint

	vector2DBs map[uint]*Vector2DB

	nextIDVector2DB uint

	vector3DBs map[uint]*Vector3DB

	nextIDVector3DB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		ambiantlightDBs: make(map[uint]*AmbiantLightDB),

		boxgeometryDBs: make(map[uint]*BoxGeometryDB),

		buffergeometryDBs: make(map[uint]*BufferGeometryDB),

		cameraDBs: make(map[uint]*CameraDB),

		canvasDBs: make(map[uint]*CanvasDB),

		curveDBs: make(map[uint]*CurveDB),

		cylindergeometryDBs: make(map[uint]*CylinderGeometryDB),

		directionallightDBs: make(map[uint]*DirectionalLightDB),

		extrudegeometryDBs: make(map[uint]*ExtrudeGeometryDB),

		meshDBs: make(map[uint]*MeshDB),

		meshmaterialbasicDBs: make(map[uint]*MeshMaterialBasicDB),

		meshphysicalmaterialDBs: make(map[uint]*MeshPhysicalMaterialDB),

		planegeometryDBs: make(map[uint]*PlaneGeometryDB),

		shapeDBs: make(map[uint]*ShapeDB),

		spheregeometryDBs: make(map[uint]*SphereGeometryDB),

		torusgeometryDBs: make(map[uint]*TorusGeometryDB),

		triangleDBs: make(map[uint]*TriangleDB),

		tubegeometryDBs: make(map[uint]*TubeGeometryDB),

		vector2DBs: make(map[uint]*Vector2DB),

		vector3DBs: make(map[uint]*Vector3DB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/threejs/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *AmbiantLightDB:
		db.nextIDAmbiantLightDB++
		v.ID = db.nextIDAmbiantLightDB
		db.ambiantlightDBs[v.ID] = v
	case *BoxGeometryDB:
		db.nextIDBoxGeometryDB++
		v.ID = db.nextIDBoxGeometryDB
		db.boxgeometryDBs[v.ID] = v
	case *BufferGeometryDB:
		db.nextIDBufferGeometryDB++
		v.ID = db.nextIDBufferGeometryDB
		db.buffergeometryDBs[v.ID] = v
	case *CameraDB:
		db.nextIDCameraDB++
		v.ID = db.nextIDCameraDB
		db.cameraDBs[v.ID] = v
	case *CanvasDB:
		db.nextIDCanvasDB++
		v.ID = db.nextIDCanvasDB
		db.canvasDBs[v.ID] = v
	case *CurveDB:
		db.nextIDCurveDB++
		v.ID = db.nextIDCurveDB
		db.curveDBs[v.ID] = v
	case *CylinderGeometryDB:
		db.nextIDCylinderGeometryDB++
		v.ID = db.nextIDCylinderGeometryDB
		db.cylindergeometryDBs[v.ID] = v
	case *DirectionalLightDB:
		db.nextIDDirectionalLightDB++
		v.ID = db.nextIDDirectionalLightDB
		db.directionallightDBs[v.ID] = v
	case *ExtrudeGeometryDB:
		db.nextIDExtrudeGeometryDB++
		v.ID = db.nextIDExtrudeGeometryDB
		db.extrudegeometryDBs[v.ID] = v
	case *MeshDB:
		db.nextIDMeshDB++
		v.ID = db.nextIDMeshDB
		db.meshDBs[v.ID] = v
	case *MeshMaterialBasicDB:
		db.nextIDMeshMaterialBasicDB++
		v.ID = db.nextIDMeshMaterialBasicDB
		db.meshmaterialbasicDBs[v.ID] = v
	case *MeshPhysicalMaterialDB:
		db.nextIDMeshPhysicalMaterialDB++
		v.ID = db.nextIDMeshPhysicalMaterialDB
		db.meshphysicalmaterialDBs[v.ID] = v
	case *PlaneGeometryDB:
		db.nextIDPlaneGeometryDB++
		v.ID = db.nextIDPlaneGeometryDB
		db.planegeometryDBs[v.ID] = v
	case *ShapeDB:
		db.nextIDShapeDB++
		v.ID = db.nextIDShapeDB
		db.shapeDBs[v.ID] = v
	case *SphereGeometryDB:
		db.nextIDSphereGeometryDB++
		v.ID = db.nextIDSphereGeometryDB
		db.spheregeometryDBs[v.ID] = v
	case *TorusGeometryDB:
		db.nextIDTorusGeometryDB++
		v.ID = db.nextIDTorusGeometryDB
		db.torusgeometryDBs[v.ID] = v
	case *TriangleDB:
		db.nextIDTriangleDB++
		v.ID = db.nextIDTriangleDB
		db.triangleDBs[v.ID] = v
	case *TubeGeometryDB:
		db.nextIDTubeGeometryDB++
		v.ID = db.nextIDTubeGeometryDB
		db.tubegeometryDBs[v.ID] = v
	case *Vector2DB:
		db.nextIDVector2DB++
		v.ID = db.nextIDVector2DB
		db.vector2DBs[v.ID] = v
	case *Vector3DB:
		db.nextIDVector3DB++
		v.ID = db.nextIDVector3DB
		db.vector3DBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/threejs/go, unsupported type in Create")
	}
	return db, nil
}

// Unscoped sets the unscoped flag for soft-deletes (not used in this implementation)
func (db *DBLite) Unscoped() (db.DBInterface, error) {
	return db, nil
}

// Model is a placeholder in this implementation
func (db *DBLite) Model(instanceDB any) (db.DBInterface, error) {
	// Not implemented as types are handled directly
	return db, nil
}

// Delete removes a record from the database
func (db *DBLite) Delete(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/threejs/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AmbiantLightDB:
		delete(db.ambiantlightDBs, v.ID)
	case *BoxGeometryDB:
		delete(db.boxgeometryDBs, v.ID)
	case *BufferGeometryDB:
		delete(db.buffergeometryDBs, v.ID)
	case *CameraDB:
		delete(db.cameraDBs, v.ID)
	case *CanvasDB:
		delete(db.canvasDBs, v.ID)
	case *CurveDB:
		delete(db.curveDBs, v.ID)
	case *CylinderGeometryDB:
		delete(db.cylindergeometryDBs, v.ID)
	case *DirectionalLightDB:
		delete(db.directionallightDBs, v.ID)
	case *ExtrudeGeometryDB:
		delete(db.extrudegeometryDBs, v.ID)
	case *MeshDB:
		delete(db.meshDBs, v.ID)
	case *MeshMaterialBasicDB:
		delete(db.meshmaterialbasicDBs, v.ID)
	case *MeshPhysicalMaterialDB:
		delete(db.meshphysicalmaterialDBs, v.ID)
	case *PlaneGeometryDB:
		delete(db.planegeometryDBs, v.ID)
	case *ShapeDB:
		delete(db.shapeDBs, v.ID)
	case *SphereGeometryDB:
		delete(db.spheregeometryDBs, v.ID)
	case *TorusGeometryDB:
		delete(db.torusgeometryDBs, v.ID)
	case *TriangleDB:
		delete(db.triangleDBs, v.ID)
	case *TubeGeometryDB:
		delete(db.tubegeometryDBs, v.ID)
	case *Vector2DB:
		delete(db.vector2DBs, v.ID)
	case *Vector3DB:
		delete(db.vector3DBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/threejs/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/threejs/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AmbiantLightDB:
		db.ambiantlightDBs[v.ID] = v
		return db, nil
	case *BoxGeometryDB:
		db.boxgeometryDBs[v.ID] = v
		return db, nil
	case *BufferGeometryDB:
		db.buffergeometryDBs[v.ID] = v
		return db, nil
	case *CameraDB:
		db.cameraDBs[v.ID] = v
		return db, nil
	case *CanvasDB:
		db.canvasDBs[v.ID] = v
		return db, nil
	case *CurveDB:
		db.curveDBs[v.ID] = v
		return db, nil
	case *CylinderGeometryDB:
		db.cylindergeometryDBs[v.ID] = v
		return db, nil
	case *DirectionalLightDB:
		db.directionallightDBs[v.ID] = v
		return db, nil
	case *ExtrudeGeometryDB:
		db.extrudegeometryDBs[v.ID] = v
		return db, nil
	case *MeshDB:
		db.meshDBs[v.ID] = v
		return db, nil
	case *MeshMaterialBasicDB:
		db.meshmaterialbasicDBs[v.ID] = v
		return db, nil
	case *MeshPhysicalMaterialDB:
		db.meshphysicalmaterialDBs[v.ID] = v
		return db, nil
	case *PlaneGeometryDB:
		db.planegeometryDBs[v.ID] = v
		return db, nil
	case *ShapeDB:
		db.shapeDBs[v.ID] = v
		return db, nil
	case *SphereGeometryDB:
		db.spheregeometryDBs[v.ID] = v
		return db, nil
	case *TorusGeometryDB:
		db.torusgeometryDBs[v.ID] = v
		return db, nil
	case *TriangleDB:
		db.triangleDBs[v.ID] = v
		return db, nil
	case *TubeGeometryDB:
		db.tubegeometryDBs[v.ID] = v
		return db, nil
	case *Vector2DB:
		db.vector2DBs[v.ID] = v
		return db, nil
	case *Vector3DB:
		db.vector3DBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/threejs/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/threejs/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AmbiantLightDB:
		if existing, ok := db.ambiantlightDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db AmbiantLight github.com/fullstack-lang/gong/lib/threejs/go, record not found")
		}
	case *BoxGeometryDB:
		if existing, ok := db.boxgeometryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db BoxGeometry github.com/fullstack-lang/gong/lib/threejs/go, record not found")
		}
	case *BufferGeometryDB:
		if existing, ok := db.buffergeometryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db BufferGeometry github.com/fullstack-lang/gong/lib/threejs/go, record not found")
		}
	case *CameraDB:
		if existing, ok := db.cameraDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Camera github.com/fullstack-lang/gong/lib/threejs/go, record not found")
		}
	case *CanvasDB:
		if existing, ok := db.canvasDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Canvas github.com/fullstack-lang/gong/lib/threejs/go, record not found")
		}
	case *CurveDB:
		if existing, ok := db.curveDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Curve github.com/fullstack-lang/gong/lib/threejs/go, record not found")
		}
	case *CylinderGeometryDB:
		if existing, ok := db.cylindergeometryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db CylinderGeometry github.com/fullstack-lang/gong/lib/threejs/go, record not found")
		}
	case *DirectionalLightDB:
		if existing, ok := db.directionallightDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DirectionalLight github.com/fullstack-lang/gong/lib/threejs/go, record not found")
		}
	case *ExtrudeGeometryDB:
		if existing, ok := db.extrudegeometryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ExtrudeGeometry github.com/fullstack-lang/gong/lib/threejs/go, record not found")
		}
	case *MeshDB:
		if existing, ok := db.meshDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Mesh github.com/fullstack-lang/gong/lib/threejs/go, record not found")
		}
	case *MeshMaterialBasicDB:
		if existing, ok := db.meshmaterialbasicDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db MeshMaterialBasic github.com/fullstack-lang/gong/lib/threejs/go, record not found")
		}
	case *MeshPhysicalMaterialDB:
		if existing, ok := db.meshphysicalmaterialDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db MeshPhysicalMaterial github.com/fullstack-lang/gong/lib/threejs/go, record not found")
		}
	case *PlaneGeometryDB:
		if existing, ok := db.planegeometryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db PlaneGeometry github.com/fullstack-lang/gong/lib/threejs/go, record not found")
		}
	case *ShapeDB:
		if existing, ok := db.shapeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Shape github.com/fullstack-lang/gong/lib/threejs/go, record not found")
		}
	case *SphereGeometryDB:
		if existing, ok := db.spheregeometryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SphereGeometry github.com/fullstack-lang/gong/lib/threejs/go, record not found")
		}
	case *TorusGeometryDB:
		if existing, ok := db.torusgeometryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db TorusGeometry github.com/fullstack-lang/gong/lib/threejs/go, record not found")
		}
	case *TriangleDB:
		if existing, ok := db.triangleDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Triangle github.com/fullstack-lang/gong/lib/threejs/go, record not found")
		}
	case *TubeGeometryDB:
		if existing, ok := db.tubegeometryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db TubeGeometry github.com/fullstack-lang/gong/lib/threejs/go, record not found")
		}
	case *Vector2DB:
		if existing, ok := db.vector2DBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Vector2 github.com/fullstack-lang/gong/lib/threejs/go, record not found")
		}
	case *Vector3DB:
		if existing, ok := db.vector3DBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Vector3 github.com/fullstack-lang/gong/lib/threejs/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/threejs/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]AmbiantLightDB:
		*ptr = make([]AmbiantLightDB, 0, len(db.ambiantlightDBs))
		for _, v := range db.ambiantlightDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]BoxGeometryDB:
		*ptr = make([]BoxGeometryDB, 0, len(db.boxgeometryDBs))
		for _, v := range db.boxgeometryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]BufferGeometryDB:
		*ptr = make([]BufferGeometryDB, 0, len(db.buffergeometryDBs))
		for _, v := range db.buffergeometryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]CameraDB:
		*ptr = make([]CameraDB, 0, len(db.cameraDBs))
		for _, v := range db.cameraDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]CanvasDB:
		*ptr = make([]CanvasDB, 0, len(db.canvasDBs))
		for _, v := range db.canvasDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]CurveDB:
		*ptr = make([]CurveDB, 0, len(db.curveDBs))
		for _, v := range db.curveDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]CylinderGeometryDB:
		*ptr = make([]CylinderGeometryDB, 0, len(db.cylindergeometryDBs))
		for _, v := range db.cylindergeometryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DirectionalLightDB:
		*ptr = make([]DirectionalLightDB, 0, len(db.directionallightDBs))
		for _, v := range db.directionallightDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ExtrudeGeometryDB:
		*ptr = make([]ExtrudeGeometryDB, 0, len(db.extrudegeometryDBs))
		for _, v := range db.extrudegeometryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]MeshDB:
		*ptr = make([]MeshDB, 0, len(db.meshDBs))
		for _, v := range db.meshDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]MeshMaterialBasicDB:
		*ptr = make([]MeshMaterialBasicDB, 0, len(db.meshmaterialbasicDBs))
		for _, v := range db.meshmaterialbasicDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]MeshPhysicalMaterialDB:
		*ptr = make([]MeshPhysicalMaterialDB, 0, len(db.meshphysicalmaterialDBs))
		for _, v := range db.meshphysicalmaterialDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]PlaneGeometryDB:
		*ptr = make([]PlaneGeometryDB, 0, len(db.planegeometryDBs))
		for _, v := range db.planegeometryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ShapeDB:
		*ptr = make([]ShapeDB, 0, len(db.shapeDBs))
		for _, v := range db.shapeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SphereGeometryDB:
		*ptr = make([]SphereGeometryDB, 0, len(db.spheregeometryDBs))
		for _, v := range db.spheregeometryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]TorusGeometryDB:
		*ptr = make([]TorusGeometryDB, 0, len(db.torusgeometryDBs))
		for _, v := range db.torusgeometryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]TriangleDB:
		*ptr = make([]TriangleDB, 0, len(db.triangleDBs))
		for _, v := range db.triangleDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]TubeGeometryDB:
		*ptr = make([]TubeGeometryDB, 0, len(db.tubegeometryDBs))
		for _, v := range db.tubegeometryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Vector2DB:
		*ptr = make([]Vector2DB, 0, len(db.vector2DBs))
		for _, v := range db.vector2DBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Vector3DB:
		*ptr = make([]Vector3DB, 0, len(db.vector3DBs))
		for _, v := range db.vector3DBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/threejs/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/threejs/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gong/lib/threejs/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/threejs/go, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *AmbiantLightDB:
		tmp, ok := db.ambiantlightDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First AmbiantLight Unkown entry %d", i))
		}

		ambiantlightDB, _ := instanceDB.(*AmbiantLightDB)
		*ambiantlightDB = *tmp

	case *BoxGeometryDB:
		tmp, ok := db.boxgeometryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First BoxGeometry Unkown entry %d", i))
		}

		boxgeometryDB, _ := instanceDB.(*BoxGeometryDB)
		*boxgeometryDB = *tmp

	case *BufferGeometryDB:
		tmp, ok := db.buffergeometryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First BufferGeometry Unkown entry %d", i))
		}

		buffergeometryDB, _ := instanceDB.(*BufferGeometryDB)
		*buffergeometryDB = *tmp

	case *CameraDB:
		tmp, ok := db.cameraDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Camera Unkown entry %d", i))
		}

		cameraDB, _ := instanceDB.(*CameraDB)
		*cameraDB = *tmp

	case *CanvasDB:
		tmp, ok := db.canvasDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Canvas Unkown entry %d", i))
		}

		canvasDB, _ := instanceDB.(*CanvasDB)
		*canvasDB = *tmp

	case *CurveDB:
		tmp, ok := db.curveDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Curve Unkown entry %d", i))
		}

		curveDB, _ := instanceDB.(*CurveDB)
		*curveDB = *tmp

	case *CylinderGeometryDB:
		tmp, ok := db.cylindergeometryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First CylinderGeometry Unkown entry %d", i))
		}

		cylindergeometryDB, _ := instanceDB.(*CylinderGeometryDB)
		*cylindergeometryDB = *tmp

	case *DirectionalLightDB:
		tmp, ok := db.directionallightDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DirectionalLight Unkown entry %d", i))
		}

		directionallightDB, _ := instanceDB.(*DirectionalLightDB)
		*directionallightDB = *tmp

	case *ExtrudeGeometryDB:
		tmp, ok := db.extrudegeometryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ExtrudeGeometry Unkown entry %d", i))
		}

		extrudegeometryDB, _ := instanceDB.(*ExtrudeGeometryDB)
		*extrudegeometryDB = *tmp

	case *MeshDB:
		tmp, ok := db.meshDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Mesh Unkown entry %d", i))
		}

		meshDB, _ := instanceDB.(*MeshDB)
		*meshDB = *tmp

	case *MeshMaterialBasicDB:
		tmp, ok := db.meshmaterialbasicDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First MeshMaterialBasic Unkown entry %d", i))
		}

		meshmaterialbasicDB, _ := instanceDB.(*MeshMaterialBasicDB)
		*meshmaterialbasicDB = *tmp

	case *MeshPhysicalMaterialDB:
		tmp, ok := db.meshphysicalmaterialDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First MeshPhysicalMaterial Unkown entry %d", i))
		}

		meshphysicalmaterialDB, _ := instanceDB.(*MeshPhysicalMaterialDB)
		*meshphysicalmaterialDB = *tmp

	case *PlaneGeometryDB:
		tmp, ok := db.planegeometryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First PlaneGeometry Unkown entry %d", i))
		}

		planegeometryDB, _ := instanceDB.(*PlaneGeometryDB)
		*planegeometryDB = *tmp

	case *ShapeDB:
		tmp, ok := db.shapeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Shape Unkown entry %d", i))
		}

		shapeDB, _ := instanceDB.(*ShapeDB)
		*shapeDB = *tmp

	case *SphereGeometryDB:
		tmp, ok := db.spheregeometryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SphereGeometry Unkown entry %d", i))
		}

		spheregeometryDB, _ := instanceDB.(*SphereGeometryDB)
		*spheregeometryDB = *tmp

	case *TorusGeometryDB:
		tmp, ok := db.torusgeometryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First TorusGeometry Unkown entry %d", i))
		}

		torusgeometryDB, _ := instanceDB.(*TorusGeometryDB)
		*torusgeometryDB = *tmp

	case *TriangleDB:
		tmp, ok := db.triangleDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Triangle Unkown entry %d", i))
		}

		triangleDB, _ := instanceDB.(*TriangleDB)
		*triangleDB = *tmp

	case *TubeGeometryDB:
		tmp, ok := db.tubegeometryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First TubeGeometry Unkown entry %d", i))
		}

		tubegeometryDB, _ := instanceDB.(*TubeGeometryDB)
		*tubegeometryDB = *tmp

	case *Vector2DB:
		tmp, ok := db.vector2DBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Vector2 Unkown entry %d", i))
		}

		vector2DB, _ := instanceDB.(*Vector2DB)
		*vector2DB = *tmp

	case *Vector3DB:
		tmp, ok := db.vector3DBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Vector3 Unkown entry %d", i))
		}

		vector3DB, _ := instanceDB.(*Vector3DB)
		*vector3DB = *tmp

	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/threejs/go, Unkown type")
	}

	return db, nil
}
