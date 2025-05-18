// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gong/test/test/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	astructDBs map[uint]*AstructDB

	nextIDAstructDB uint

	astructbstruct2useDBs map[uint]*AstructBstruct2UseDB

	nextIDAstructBstruct2UseDB uint

	astructbstructuseDBs map[uint]*AstructBstructUseDB

	nextIDAstructBstructUseDB uint

	bstructDBs map[uint]*BstructDB

	nextIDBstructDB uint

	dstructDBs map[uint]*DstructDB

	nextIDDstructDB uint

	f0123456789012345678901234567890DBs map[uint]*F0123456789012345678901234567890DB

	nextIDF0123456789012345678901234567890DB uint

	gstructDBs map[uint]*GstructDB

	nextIDGstructDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		astructDBs: make(map[uint]*AstructDB),

		astructbstruct2useDBs: make(map[uint]*AstructBstruct2UseDB),

		astructbstructuseDBs: make(map[uint]*AstructBstructUseDB),

		bstructDBs: make(map[uint]*BstructDB),

		dstructDBs: make(map[uint]*DstructDB),

		f0123456789012345678901234567890DBs: make(map[uint]*F0123456789012345678901234567890DB),

		gstructDBs: make(map[uint]*GstructDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/test/test/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *AstructDB:
		db.nextIDAstructDB++
		v.ID = db.nextIDAstructDB
		db.astructDBs[v.ID] = v
	case *AstructBstruct2UseDB:
		db.nextIDAstructBstruct2UseDB++
		v.ID = db.nextIDAstructBstruct2UseDB
		db.astructbstruct2useDBs[v.ID] = v
	case *AstructBstructUseDB:
		db.nextIDAstructBstructUseDB++
		v.ID = db.nextIDAstructBstructUseDB
		db.astructbstructuseDBs[v.ID] = v
	case *BstructDB:
		db.nextIDBstructDB++
		v.ID = db.nextIDBstructDB
		db.bstructDBs[v.ID] = v
	case *DstructDB:
		db.nextIDDstructDB++
		v.ID = db.nextIDDstructDB
		db.dstructDBs[v.ID] = v
	case *F0123456789012345678901234567890DB:
		db.nextIDF0123456789012345678901234567890DB++
		v.ID = db.nextIDF0123456789012345678901234567890DB
		db.f0123456789012345678901234567890DBs[v.ID] = v
	case *GstructDB:
		db.nextIDGstructDB++
		v.ID = db.nextIDGstructDB
		db.gstructDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/test/test/go, unsupported type in Create")
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
		return nil, errors.New("github.com/fullstack-lang/gong/test/test/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AstructDB:
		delete(db.astructDBs, v.ID)
	case *AstructBstruct2UseDB:
		delete(db.astructbstruct2useDBs, v.ID)
	case *AstructBstructUseDB:
		delete(db.astructbstructuseDBs, v.ID)
	case *BstructDB:
		delete(db.bstructDBs, v.ID)
	case *DstructDB:
		delete(db.dstructDBs, v.ID)
	case *F0123456789012345678901234567890DB:
		delete(db.f0123456789012345678901234567890DBs, v.ID)
	case *GstructDB:
		delete(db.gstructDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/test/test/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/test/test/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AstructDB:
		db.astructDBs[v.ID] = v
		return db, nil
	case *AstructBstruct2UseDB:
		db.astructbstruct2useDBs[v.ID] = v
		return db, nil
	case *AstructBstructUseDB:
		db.astructbstructuseDBs[v.ID] = v
		return db, nil
	case *BstructDB:
		db.bstructDBs[v.ID] = v
		return db, nil
	case *DstructDB:
		db.dstructDBs[v.ID] = v
		return db, nil
	case *F0123456789012345678901234567890DB:
		db.f0123456789012345678901234567890DBs[v.ID] = v
		return db, nil
	case *GstructDB:
		db.gstructDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/test/test/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/test/test/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AstructDB:
		if existing, ok := db.astructDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Astruct github.com/fullstack-lang/gong/test/test/go, record not found")
		}
	case *AstructBstruct2UseDB:
		if existing, ok := db.astructbstruct2useDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db AstructBstruct2Use github.com/fullstack-lang/gong/test/test/go, record not found")
		}
	case *AstructBstructUseDB:
		if existing, ok := db.astructbstructuseDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db AstructBstructUse github.com/fullstack-lang/gong/test/test/go, record not found")
		}
	case *BstructDB:
		if existing, ok := db.bstructDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Bstruct github.com/fullstack-lang/gong/test/test/go, record not found")
		}
	case *DstructDB:
		if existing, ok := db.dstructDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Dstruct github.com/fullstack-lang/gong/test/test/go, record not found")
		}
	case *F0123456789012345678901234567890DB:
		if existing, ok := db.f0123456789012345678901234567890DBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db F0123456789012345678901234567890 github.com/fullstack-lang/gong/test/test/go, record not found")
		}
	case *GstructDB:
		if existing, ok := db.gstructDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Gstruct github.com/fullstack-lang/gong/test/test/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/test/test/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]AstructDB:
		*ptr = make([]AstructDB, 0, len(db.astructDBs))
		for _, v := range db.astructDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]AstructBstruct2UseDB:
		*ptr = make([]AstructBstruct2UseDB, 0, len(db.astructbstruct2useDBs))
		for _, v := range db.astructbstruct2useDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]AstructBstructUseDB:
		*ptr = make([]AstructBstructUseDB, 0, len(db.astructbstructuseDBs))
		for _, v := range db.astructbstructuseDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]BstructDB:
		*ptr = make([]BstructDB, 0, len(db.bstructDBs))
		for _, v := range db.bstructDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DstructDB:
		*ptr = make([]DstructDB, 0, len(db.dstructDBs))
		for _, v := range db.dstructDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]F0123456789012345678901234567890DB:
		*ptr = make([]F0123456789012345678901234567890DB, 0, len(db.f0123456789012345678901234567890DBs))
		for _, v := range db.f0123456789012345678901234567890DBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]GstructDB:
		*ptr = make([]GstructDB, 0, len(db.gstructDBs))
		for _, v := range db.gstructDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/test/test/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gong/test/test/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gong/test/test/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/test/test/go, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *AstructDB:
		tmp, ok := db.astructDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Astruct Unkown entry %d", i))
		}

		astructDB, _ := instanceDB.(*AstructDB)
		*astructDB = *tmp
		
	case *AstructBstruct2UseDB:
		tmp, ok := db.astructbstruct2useDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First AstructBstruct2Use Unkown entry %d", i))
		}

		astructbstruct2useDB, _ := instanceDB.(*AstructBstruct2UseDB)
		*astructbstruct2useDB = *tmp
		
	case *AstructBstructUseDB:
		tmp, ok := db.astructbstructuseDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First AstructBstructUse Unkown entry %d", i))
		}

		astructbstructuseDB, _ := instanceDB.(*AstructBstructUseDB)
		*astructbstructuseDB = *tmp
		
	case *BstructDB:
		tmp, ok := db.bstructDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Bstruct Unkown entry %d", i))
		}

		bstructDB, _ := instanceDB.(*BstructDB)
		*bstructDB = *tmp
		
	case *DstructDB:
		tmp, ok := db.dstructDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Dstruct Unkown entry %d", i))
		}

		dstructDB, _ := instanceDB.(*DstructDB)
		*dstructDB = *tmp
		
	case *F0123456789012345678901234567890DB:
		tmp, ok := db.f0123456789012345678901234567890DBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First F0123456789012345678901234567890 Unkown entry %d", i))
		}

		f0123456789012345678901234567890DB, _ := instanceDB.(*F0123456789012345678901234567890DB)
		*f0123456789012345678901234567890DB = *tmp
		
	case *GstructDB:
		tmp, ok := db.gstructDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Gstruct Unkown entry %d", i))
		}

		gstructDB, _ := instanceDB.(*GstructDB)
		*gstructDB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/test/test/go, Unkown type")
	}
	
	return db, nil
}

