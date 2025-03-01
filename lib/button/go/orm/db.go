// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gong/lib/button/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	buttonDBs map[uint]*ButtonDB

	nextIDButtonDB uint

	groupDBs map[uint]*GroupDB

	nextIDGroupDB uint

	layoutDBs map[uint]*LayoutDB

	nextIDLayoutDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		buttonDBs: make(map[uint]*ButtonDB),

		groupDBs: make(map[uint]*GroupDB),

		layoutDBs: make(map[uint]*LayoutDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/button/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *ButtonDB:
		db.nextIDButtonDB++
		v.ID = db.nextIDButtonDB
		db.buttonDBs[v.ID] = v
	case *GroupDB:
		db.nextIDGroupDB++
		v.ID = db.nextIDGroupDB
		db.groupDBs[v.ID] = v
	case *LayoutDB:
		db.nextIDLayoutDB++
		v.ID = db.nextIDLayoutDB
		db.layoutDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/button/go, unsupported type in Create")
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
		return nil, errors.New("github.com/fullstack-lang/gong/lib/button/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ButtonDB:
		delete(db.buttonDBs, v.ID)
	case *GroupDB:
		delete(db.groupDBs, v.ID)
	case *LayoutDB:
		delete(db.layoutDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/button/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/button/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ButtonDB:
		db.buttonDBs[v.ID] = v
		return db, nil
	case *GroupDB:
		db.groupDBs[v.ID] = v
		return db, nil
	case *LayoutDB:
		db.layoutDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/button/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/button/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ButtonDB:
		if existing, ok := db.buttonDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Button github.com/fullstack-lang/gong/lib/button/go, record not found")
		}
	case *GroupDB:
		if existing, ok := db.groupDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Group github.com/fullstack-lang/gong/lib/button/go, record not found")
		}
	case *LayoutDB:
		if existing, ok := db.layoutDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Layout github.com/fullstack-lang/gong/lib/button/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/button/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]ButtonDB:
		*ptr = make([]ButtonDB, 0, len(db.buttonDBs))
		for _, v := range db.buttonDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]GroupDB:
		*ptr = make([]GroupDB, 0, len(db.groupDBs))
		for _, v := range db.groupDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]LayoutDB:
		*ptr = make([]LayoutDB, 0, len(db.layoutDBs))
		for _, v := range db.layoutDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/button/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/button/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gong/lib/button/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/button/go, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *ButtonDB:
		tmp, ok := db.buttonDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Button Unkown entry %d", i))
		}

		buttonDB, _ := instanceDB.(*ButtonDB)
		*buttonDB = *tmp
		
	case *GroupDB:
		tmp, ok := db.groupDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Group Unkown entry %d", i))
		}

		groupDB, _ := instanceDB.(*GroupDB)
		*groupDB = *tmp
		
	case *LayoutDB:
		tmp, ok := db.layoutDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Layout Unkown entry %d", i))
		}

		layoutDB, _ := instanceDB.(*LayoutDB)
		*layoutDB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/button/go, Unkown type")
	}
	
	return db, nil
}

