// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gong/lib/markdown/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	contentDBs map[uint]*ContentDB

	nextIDContentDB uint

	svgimageDBs map[uint]*SvgImageDB

	nextIDSvgImageDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		contentDBs: make(map[uint]*ContentDB),

		svgimageDBs: make(map[uint]*SvgImageDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/markdown/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *ContentDB:
		db.nextIDContentDB++
		v.ID = db.nextIDContentDB
		db.contentDBs[v.ID] = v
	case *SvgImageDB:
		db.nextIDSvgImageDB++
		v.ID = db.nextIDSvgImageDB
		db.svgimageDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/markdown/go, unsupported type in Create")
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
		return nil, errors.New("github.com/fullstack-lang/gong/lib/markdown/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ContentDB:
		delete(db.contentDBs, v.ID)
	case *SvgImageDB:
		delete(db.svgimageDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/markdown/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/markdown/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ContentDB:
		db.contentDBs[v.ID] = v
		return db, nil
	case *SvgImageDB:
		db.svgimageDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/markdown/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/markdown/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ContentDB:
		if existing, ok := db.contentDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Content github.com/fullstack-lang/gong/lib/markdown/go, record not found")
		}
	case *SvgImageDB:
		if existing, ok := db.svgimageDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SvgImage github.com/fullstack-lang/gong/lib/markdown/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/markdown/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]ContentDB:
		*ptr = make([]ContentDB, 0, len(db.contentDBs))
		for _, v := range db.contentDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SvgImageDB:
		*ptr = make([]SvgImageDB, 0, len(db.svgimageDBs))
		for _, v := range db.svgimageDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/markdown/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/markdown/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gong/lib/markdown/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/markdown/go, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *ContentDB:
		tmp, ok := db.contentDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Content Unkown entry %d", i))
		}

		contentDB, _ := instanceDB.(*ContentDB)
		*contentDB = *tmp
		
	case *SvgImageDB:
		tmp, ok := db.svgimageDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SvgImage Unkown entry %d", i))
		}

		svgimageDB, _ := instanceDB.(*SvgImageDB)
		*svgimageDB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/markdown/go, Unkown type")
	}
	
	return db, nil
}

