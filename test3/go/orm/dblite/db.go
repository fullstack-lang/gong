// generated code - do not edit
package dblite

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/fullstack-lang/gong/test3/go/db"
	"github.com/fullstack-lang/gong/test3/go/orm"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// insertion point definitions

	aDBs map[uint]*orm.ADB

	nextIDADB uint

	bDBs map[uint]*orm.BDB

	nextIDBDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		aDBs: make(map[uint]*orm.ADB),

		bDBs: make(map[uint]*orm.BDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("instanceDB cannot be nil")
	}
	switch v := instanceDB.(type) {
	// insertion point create
	case **orm.ADB:
		db.nextIDADB++
		(*v).ID = db.nextIDADB
		db.aDBs[(*v).ID] = *v
	case **orm.BDB:
		db.nextIDBDB++
		(*v).ID = db.nextIDBDB
		db.bDBs[(*v).ID] = *v
	default:
		return nil, errors.New("unsupported type in Create")
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
		return nil, errors.New("instanceDB cannot be nil")
	}
	switch v := instanceDB.(type) {
	// insertion point delete
	case *orm.ADB:
		delete(db.aDBs, v.ID)
	case *orm.BDB:
		delete(db.bDBs, v.ID)
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
	// insertion point delete
	case *orm.ADB:
		if existing, ok := db.aDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("record not found")
		}
	case *orm.BDB:
		if existing, ok := db.bDBs[v.ID]; ok {
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
	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]orm.ADB:
        *ptr = make([]orm.ADB, 0, len(db.aDBs))
        for _, v := range db.aDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]orm.BDB:
        *ptr = make([]orm.BDB, 0, len(db.bDBs))
        for _, v := range db.bDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
    default:
        return nil, errors.New("Find: unsupported type")
    }
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("Do not process when conds is not a single parameter")
	}

	str, ok := conds[0].(string)

	if !ok {
		return nil, errors.New("conds[0] is not a string")
	}

	i, err := strconv.ParseUint(str, 10, 32) // Base 10, 32-bit unsigned int
	if err != nil {
		return nil, errors.New("conds[0] is not a string number")
	}

	switch instanceDB.(type) {
	// insertion point first
	case *orm.ADB:
		tmp, ok := db.aDBs[uint(i)]

		aDB, _ := instanceDB.(*orm.ADB)
		*aDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *orm.BDB:
		tmp, ok := db.bDBs[uint(i)]

		bDB, _ := instanceDB.(*orm.BDB)
		*bDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	default:
		return nil, errors.New("Unkown type")
	}
	
	return db, nil
}

