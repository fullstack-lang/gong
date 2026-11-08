// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gong/lib/maplibre/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	coordinateDBs map[uint]*CoordinateDB

	nextIDCoordinateDB uint

	layerDBs map[uint]*LayerDB

	nextIDLayerDB uint

	mapviewDBs map[uint]*MapViewDB

	nextIDMapViewDB uint

	markerDBs map[uint]*MarkerDB

	nextIDMarkerDB uint

	polygonDBs map[uint]*PolygonDB

	nextIDPolygonDB uint

	polylineDBs map[uint]*PolylineDB

	nextIDPolylineDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		coordinateDBs: make(map[uint]*CoordinateDB),

		layerDBs: make(map[uint]*LayerDB),

		mapviewDBs: make(map[uint]*MapViewDB),

		markerDBs: make(map[uint]*MarkerDB),

		polygonDBs: make(map[uint]*PolygonDB),

		polylineDBs: make(map[uint]*PolylineDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/maplibre/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *CoordinateDB:
		db.nextIDCoordinateDB++
		v.ID = db.nextIDCoordinateDB
		db.coordinateDBs[v.ID] = v
	case *LayerDB:
		db.nextIDLayerDB++
		v.ID = db.nextIDLayerDB
		db.layerDBs[v.ID] = v
	case *MapViewDB:
		db.nextIDMapViewDB++
		v.ID = db.nextIDMapViewDB
		db.mapviewDBs[v.ID] = v
	case *MarkerDB:
		db.nextIDMarkerDB++
		v.ID = db.nextIDMarkerDB
		db.markerDBs[v.ID] = v
	case *PolygonDB:
		db.nextIDPolygonDB++
		v.ID = db.nextIDPolygonDB
		db.polygonDBs[v.ID] = v
	case *PolylineDB:
		db.nextIDPolylineDB++
		v.ID = db.nextIDPolylineDB
		db.polylineDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/maplibre/go, unsupported type in Create")
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
		return nil, errors.New("github.com/fullstack-lang/gong/lib/maplibre/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *CoordinateDB:
		delete(db.coordinateDBs, v.ID)
	case *LayerDB:
		delete(db.layerDBs, v.ID)
	case *MapViewDB:
		delete(db.mapviewDBs, v.ID)
	case *MarkerDB:
		delete(db.markerDBs, v.ID)
	case *PolygonDB:
		delete(db.polygonDBs, v.ID)
	case *PolylineDB:
		delete(db.polylineDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/maplibre/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/maplibre/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *CoordinateDB:
		db.coordinateDBs[v.ID] = v
		return db, nil
	case *LayerDB:
		db.layerDBs[v.ID] = v
		return db, nil
	case *MapViewDB:
		db.mapviewDBs[v.ID] = v
		return db, nil
	case *MarkerDB:
		db.markerDBs[v.ID] = v
		return db, nil
	case *PolygonDB:
		db.polygonDBs[v.ID] = v
		return db, nil
	case *PolylineDB:
		db.polylineDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/maplibre/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/maplibre/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *CoordinateDB:
		if existing, ok := db.coordinateDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Coordinate github.com/fullstack-lang/gong/lib/maplibre/go, record not found")
		}
	case *LayerDB:
		if existing, ok := db.layerDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Layer github.com/fullstack-lang/gong/lib/maplibre/go, record not found")
		}
	case *MapViewDB:
		if existing, ok := db.mapviewDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db MapView github.com/fullstack-lang/gong/lib/maplibre/go, record not found")
		}
	case *MarkerDB:
		if existing, ok := db.markerDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Marker github.com/fullstack-lang/gong/lib/maplibre/go, record not found")
		}
	case *PolygonDB:
		if existing, ok := db.polygonDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Polygon github.com/fullstack-lang/gong/lib/maplibre/go, record not found")
		}
	case *PolylineDB:
		if existing, ok := db.polylineDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Polyline github.com/fullstack-lang/gong/lib/maplibre/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/maplibre/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]CoordinateDB:
		*ptr = make([]CoordinateDB, 0, len(db.coordinateDBs))
		for _, v := range db.coordinateDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]LayerDB:
		*ptr = make([]LayerDB, 0, len(db.layerDBs))
		for _, v := range db.layerDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]MapViewDB:
		*ptr = make([]MapViewDB, 0, len(db.mapviewDBs))
		for _, v := range db.mapviewDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]MarkerDB:
		*ptr = make([]MarkerDB, 0, len(db.markerDBs))
		for _, v := range db.markerDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]PolygonDB:
		*ptr = make([]PolygonDB, 0, len(db.polygonDBs))
		for _, v := range db.polygonDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]PolylineDB:
		*ptr = make([]PolylineDB, 0, len(db.polylineDBs))
		for _, v := range db.polylineDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/maplibre/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/maplibre/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gong/lib/maplibre/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/maplibre/go, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *CoordinateDB:
		tmp, ok := db.coordinateDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Coordinate Unkown entry %d", i))
		}

		coordinateDB, _ := instanceDB.(*CoordinateDB)
		*coordinateDB = *tmp

	case *LayerDB:
		tmp, ok := db.layerDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Layer Unkown entry %d", i))
		}

		layerDB, _ := instanceDB.(*LayerDB)
		*layerDB = *tmp

	case *MapViewDB:
		tmp, ok := db.mapviewDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First MapView Unkown entry %d", i))
		}

		mapviewDB, _ := instanceDB.(*MapViewDB)
		*mapviewDB = *tmp

	case *MarkerDB:
		tmp, ok := db.markerDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Marker Unkown entry %d", i))
		}

		markerDB, _ := instanceDB.(*MarkerDB)
		*markerDB = *tmp

	case *PolygonDB:
		tmp, ok := db.polygonDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Polygon Unkown entry %d", i))
		}

		polygonDB, _ := instanceDB.(*PolygonDB)
		*polygonDB = *tmp

	case *PolylineDB:
		tmp, ok := db.polylineDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Polyline Unkown entry %d", i))
		}

		polylineDB, _ := instanceDB.(*PolylineDB)
		*polylineDB = *tmp

	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/maplibre/go, Unkown type")
	}

	return db, nil
}
