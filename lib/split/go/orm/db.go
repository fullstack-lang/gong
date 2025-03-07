// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gong/lib/split/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	assplitDBs map[uint]*AsSplitDB

	nextIDAsSplitDB uint

	assplitareaDBs map[uint]*AsSplitAreaDB

	nextIDAsSplitAreaDB uint

	treeDBs map[uint]*TreeDB

	nextIDTreeDB uint

	viewDBs map[uint]*ViewDB

	nextIDViewDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		assplitDBs: make(map[uint]*AsSplitDB),

		assplitareaDBs: make(map[uint]*AsSplitAreaDB),

		treeDBs: make(map[uint]*TreeDB),

		viewDBs: make(map[uint]*ViewDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/split/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *AsSplitDB:
		db.nextIDAsSplitDB++
		v.ID = db.nextIDAsSplitDB
		db.assplitDBs[v.ID] = v
	case *AsSplitAreaDB:
		db.nextIDAsSplitAreaDB++
		v.ID = db.nextIDAsSplitAreaDB
		db.assplitareaDBs[v.ID] = v
	case *TreeDB:
		db.nextIDTreeDB++
		v.ID = db.nextIDTreeDB
		db.treeDBs[v.ID] = v
	case *ViewDB:
		db.nextIDViewDB++
		v.ID = db.nextIDViewDB
		db.viewDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/split/go, unsupported type in Create")
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
		return nil, errors.New("github.com/fullstack-lang/gong/lib/split/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AsSplitDB:
		delete(db.assplitDBs, v.ID)
	case *AsSplitAreaDB:
		delete(db.assplitareaDBs, v.ID)
	case *TreeDB:
		delete(db.treeDBs, v.ID)
	case *ViewDB:
		delete(db.viewDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/split/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/split/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AsSplitDB:
		db.assplitDBs[v.ID] = v
		return db, nil
	case *AsSplitAreaDB:
		db.assplitareaDBs[v.ID] = v
		return db, nil
	case *TreeDB:
		db.treeDBs[v.ID] = v
		return db, nil
	case *ViewDB:
		db.viewDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/split/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/split/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AsSplitDB:
		if existing, ok := db.assplitDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db AsSplit github.com/fullstack-lang/gong/lib/split/go, record not found")
		}
	case *AsSplitAreaDB:
		if existing, ok := db.assplitareaDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db AsSplitArea github.com/fullstack-lang/gong/lib/split/go, record not found")
		}
	case *TreeDB:
		if existing, ok := db.treeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Tree github.com/fullstack-lang/gong/lib/split/go, record not found")
		}
	case *ViewDB:
		if existing, ok := db.viewDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db View github.com/fullstack-lang/gong/lib/split/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/split/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]AsSplitDB:
		*ptr = make([]AsSplitDB, 0, len(db.assplitDBs))
		for _, v := range db.assplitDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]AsSplitAreaDB:
		*ptr = make([]AsSplitAreaDB, 0, len(db.assplitareaDBs))
		for _, v := range db.assplitareaDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]TreeDB:
		*ptr = make([]TreeDB, 0, len(db.treeDBs))
		for _, v := range db.treeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ViewDB:
		*ptr = make([]ViewDB, 0, len(db.viewDBs))
		for _, v := range db.viewDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/split/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/split/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gong/lib/split/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/split/go, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *AsSplitDB:
		tmp, ok := db.assplitDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First AsSplit Unkown entry %d", i))
		}

		assplitDB, _ := instanceDB.(*AsSplitDB)
		*assplitDB = *tmp
		
	case *AsSplitAreaDB:
		tmp, ok := db.assplitareaDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First AsSplitArea Unkown entry %d", i))
		}

		assplitareaDB, _ := instanceDB.(*AsSplitAreaDB)
		*assplitareaDB = *tmp
		
	case *TreeDB:
		tmp, ok := db.treeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Tree Unkown entry %d", i))
		}

		treeDB, _ := instanceDB.(*TreeDB)
		*treeDB = *tmp
		
	case *ViewDB:
		tmp, ok := db.viewDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First View Unkown entry %d", i))
		}

		viewDB, _ := instanceDB.(*ViewDB)
		*viewDB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/split/go, Unkown type")
	}
	
	return db, nil
}

