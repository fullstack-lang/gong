// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gongtree/go/db"
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

	nodeDBs map[uint]*NodeDB

	nextIDNodeDB uint

	svgiconDBs map[uint]*SVGIconDB

	nextIDSVGIconDB uint

	treeDBs map[uint]*TreeDB

	nextIDTreeDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		buttonDBs: make(map[uint]*ButtonDB),

		nodeDBs: make(map[uint]*NodeDB),

		svgiconDBs: make(map[uint]*SVGIconDB),

		treeDBs: make(map[uint]*TreeDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongtree/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *ButtonDB:
		db.nextIDButtonDB++
		v.ID = db.nextIDButtonDB
		db.buttonDBs[v.ID] = v
	case *NodeDB:
		db.nextIDNodeDB++
		v.ID = db.nextIDNodeDB
		db.nodeDBs[v.ID] = v
	case *SVGIconDB:
		db.nextIDSVGIconDB++
		v.ID = db.nextIDSVGIconDB
		db.svgiconDBs[v.ID] = v
	case *TreeDB:
		db.nextIDTreeDB++
		v.ID = db.nextIDTreeDB
		db.treeDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gongtree/go, unsupported type in Create")
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
		return nil, errors.New("github.com/fullstack-lang/gongtree/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ButtonDB:
		delete(db.buttonDBs, v.ID)
	case *NodeDB:
		delete(db.nodeDBs, v.ID)
	case *SVGIconDB:
		delete(db.svgiconDBs, v.ID)
	case *TreeDB:
		delete(db.treeDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongtree/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongtree/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ButtonDB:
		db.buttonDBs[v.ID] = v
		return db, nil
	case *NodeDB:
		db.nodeDBs[v.ID] = v
		return db, nil
	case *SVGIconDB:
		db.svgiconDBs[v.ID] = v
		return db, nil
	case *TreeDB:
		db.treeDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongtree/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongtree/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ButtonDB:
		if existing, ok := db.buttonDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Button github.com/fullstack-lang/gongtree/go, record not found")
		}
	case *NodeDB:
		if existing, ok := db.nodeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Node github.com/fullstack-lang/gongtree/go, record not found")
		}
	case *SVGIconDB:
		if existing, ok := db.svgiconDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SVGIcon github.com/fullstack-lang/gongtree/go, record not found")
		}
	case *TreeDB:
		if existing, ok := db.treeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Tree github.com/fullstack-lang/gongtree/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gongtree/go, unsupported type in Updates")
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
	case *[]NodeDB:
		*ptr = make([]NodeDB, 0, len(db.nodeDBs))
		for _, v := range db.nodeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SVGIconDB:
		*ptr = make([]SVGIconDB, 0, len(db.svgiconDBs))
		for _, v := range db.svgiconDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]TreeDB:
		*ptr = make([]TreeDB, 0, len(db.treeDBs))
		for _, v := range db.treeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongtree/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gongtree/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gongtree/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongtree/go, conds[0] is not a string or uint64")
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
		
	case *NodeDB:
		tmp, ok := db.nodeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Node Unkown entry %d", i))
		}

		nodeDB, _ := instanceDB.(*NodeDB)
		*nodeDB = *tmp
		
	case *SVGIconDB:
		tmp, ok := db.svgiconDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SVGIcon Unkown entry %d", i))
		}

		svgiconDB, _ := instanceDB.(*SVGIconDB)
		*svgiconDB = *tmp
		
	case *TreeDB:
		tmp, ok := db.treeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Tree Unkown entry %d", i))
		}

		treeDB, _ := instanceDB.(*TreeDB)
		*treeDB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gongtree/go, Unkown type")
	}
	
	return db, nil
}

