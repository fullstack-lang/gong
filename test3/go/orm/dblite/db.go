package dblite

import (
	"errors"
	"reflect"

	"github.com/fullstack-lang/gong/test3/go/db"
	"github.com/fullstack-lang/gong/test3/go/orm"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	adbs     map[uint]*orm.ADB
	bdbs     map[uint]*orm.BDB
	unscoped bool

	nextIDADB uint
	nextIDBDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		adbs: make(map[uint]*orm.ADB),
		bdbs: make(map[uint]*orm.BDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("instanceDB cannot be nil")
	}
	switch v := instanceDB.(type) {
	case *orm.ADB:
		db.nextIDADB++
		v.ID = db.nextIDADB
		db.adbs[v.ID] = v
	case *orm.BDB:
		db.nextIDBDB++
		v.ID = db.nextIDBDB
		db.bdbs[v.ID] = v
	default:
		return nil, errors.New("unsupported type in Create")
	}
	return db, nil
}

// Unscoped sets the unscoped flag for soft-deletes (not used in this implementation)
func (db *DBLite) Unscoped() (db.DBInterface, error) {
	db.unscoped = true
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
		return nil, errors.New("instanceDB cannot be nil")
	}
	switch v := instanceDB.(type) {
	case *orm.ADB:
		delete(db.adbs, v.ID)
	case *orm.BDB:
		delete(db.bdbs, v.ID)
	default:
		return nil, errors.New("unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {
	return db.Create(instanceDB)
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("instanceDB cannot be nil")
	}
	switch v := instanceDB.(type) {
	case *orm.ADB:
		if existing, ok := db.adbs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("record not found")
		}
	case *orm.BDB:
		if existing, ok := db.bdbs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("record not found")
		}
	default:
		return nil, errors.New("unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {
	if instanceDBs == nil {
		return nil, errors.New("instanceDBs cannot be nil")
	}
	val := reflect.ValueOf(instanceDBs)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Slice {
		return nil, errors.New("instanceDBs must be a pointer to a slice")
	}
	sliceVal := val.Elem()
	elemType := sliceVal.Type().Elem()

	switch elemType {
	case reflect.TypeOf(&orm.ADB{}):
		for _, v := range db.adbs {
			sliceVal.Set(reflect.Append(sliceVal, reflect.ValueOf(v)))
		}
	case reflect.TypeOf(&orm.BDB{}):
		for _, v := range db.bdbs {
			sliceVal.Set(reflect.Append(sliceVal, reflect.ValueOf(v)))
		}
	default:
		return nil, errors.New("unsupported type in Find")
	}
	return db, nil
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("instanceDB cannot be nil")
	}
	switch v := instanceDB.(type) {
	case *orm.ADB:
		for _, adb := range db.adbs {
			*v = *adb
			return db, nil
		}
	case *orm.BDB:
		for _, bdb := range db.bdbs {
			*v = *bdb
			return db, nil
		}
	default:
		return nil, errors.New("unsupported type in First")
	}
	return nil, errors.New("no records found")
}
