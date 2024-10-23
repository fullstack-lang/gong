package dblite

import (
	"errors"
	"sync"

	"github.com/fullstack-lang/gong/test3/go/db"
	"github.com/fullstack-lang/gong/test3/go/orm"
)

// Ensure DBlite implements db.DBInterface
var _ db.DBInterface = &DBlite{}

// DBlite represents an in-memory database for ADB instances
type DBlite struct {
	data   map[uint]*orm.ADB
	mutex  sync.RWMutex
	nextID uint
}

// NewDBlite creates a new instance of DBlite
func NewDBlite() *DBlite {
	return &DBlite{
		data:   make(map[uint]*orm.ADB),
		nextID: 1,
	}
}

// Create adds a new ADB instance to the database
func (db *DBlite) Create(instanceDB any) (db.DBInterface, error) {
	adb, ok := instanceDB.(*orm.ADB)
	if !ok {
		return db, errors.New("Create: instanceDB is not *orm.ADB")
	}
	db.mutex.Lock()
	defer db.mutex.Unlock()
	adb.ID = db.nextID
	db.nextID++
	db.data[adb.ID] = adb
	return db, nil
}

// Unscoped returns the database instance (no operation needed for in-memory DB)
func (db *DBlite) Unscoped() (db.DBInterface, error) {
	return db, nil
}

// Model sets the model type (not needed here as we only deal with ADB)
func (db *DBlite) Model(instanceDB any) (db.DBInterface, error) {
	return db, nil
}

// Delete removes an ADB instance from the database
func (db *DBlite) Delete(instanceDB any) (db.DBInterface, error) {
	adb, ok := instanceDB.(*orm.ADB)
	if !ok {
		return db, errors.New("Delete: instanceDB is not *orm.ADB")
	}
	db.mutex.Lock()
	defer db.mutex.Unlock()
	delete(db.data, adb.ID)
	return db, nil
}

// Save adds or updates an ADB instance in the database
func (db *DBlite) Save(instanceDB any) (db.DBInterface, error) {
	adb, ok := instanceDB.(*orm.ADB)
	if !ok {
		return db, errors.New("Save: instanceDB is not *orm.ADB")
	}
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if adb.ID == 0 {
		adb.ID = db.nextID
		db.nextID++
	}
	db.data[adb.ID] = adb
	return db, nil
}

// Updates modifies an existing ADB instance in the database
func (db *DBlite) Updates(instanceDB any) (db.DBInterface, error) {
	adb, ok := instanceDB.(*orm.ADB)
	if !ok {
		return db, errors.New("Updates: instanceDB is not *orm.ADB")
	}
	db.mutex.Lock()
	defer db.mutex.Unlock()
	existing, exists := db.data[adb.ID]
	if !exists {
		return db, errors.New("Updates: record not found")
	}
	existing.Name_Data = adb.Name_Data
	return db, nil
}

// Find retrieves all ADB instances from the database
func (db *DBlite) Find(instanceDBs any) (db.DBInterface, error) {
	slice, ok := instanceDBs.(*[]orm.ADB)
	if !ok {
		return db, errors.New("Find: instanceDBs is not *[]orm.ADB")
	}
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	*slice = make([]orm.ADB, 0, len(db.data))
	for _, adb := range db.data {
		*slice = append(*slice, *adb)
	}
	return db, nil
}

// First retrieves the first ADB instance matching the condition
func (db *DBlite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	adb, ok := instanceDB.(*orm.ADB)
	if !ok {
		return db, errors.New("First: instanceDB is not *orm.ADB")
	}
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	if len(conds) > 0 {
		id, ok := conds[0].(uint)
		if !ok {
			return db, errors.New("First: conds[0] is not uint ID")
		}
		record, exists := db.data[id]
		if !exists {
			return db, errors.New("First: record not found")
		}
		*adb = *record
		return db, nil
	}
	for _, record := range db.data {
		*adb = *record
		return db, nil
	}
	return db, errors.New("First: no records found")
}
