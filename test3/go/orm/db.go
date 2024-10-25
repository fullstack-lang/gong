// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gong/test3/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	aDBs map[uint]*ADB

	nextIDADB uint

	bDBs map[uint]*BDB

	nextIDBDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		aDBs: make(map[uint]*ADB),

		bDBs: make(map[uint]*BDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *ADB:
		db.nextIDADB++
		v.ID = db.nextIDADB
		db.aDBs[v.ID] = v
	case *BDB:
		db.nextIDBDB++
		v.ID = db.nextIDBDB
		db.bDBs[v.ID] = v
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

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ADB:
		delete(db.aDBs, v.ID)
	case *BDB:
		delete(db.bDBs, v.ID)
	default:
		return nil, errors.New("unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ADB:
		db.aDBs[v.ID] = v
		return db, nil
	case *BDB:
		db.bDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ADB:
		if existing, ok := db.aDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("record not found")
		}
	case *BDB:
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

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]ADB:
        *ptr = make([]ADB, 0, len(db.aDBs))
        for _, v := range db.aDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]BDB:
        *ptr = make([]BDB, 0, len(db.bDBs))
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

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *ADB:
		tmp, ok := db.aDBs[uint(i)]

		aDB, _ := instanceDB.(*ADB)
		*aDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *BDB:
		tmp, ok := db.bDBs[uint(i)]

		bDB, _ := instanceDB.(*BDB)
		*bDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	default:
		return nil, errors.New("Unkown type")
	}
	
	return db, nil
}

