// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gong/lib/tone/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	freqencyDBs map[uint]*FreqencyDB

	nextIDFreqencyDB uint

	noteDBs map[uint]*NoteDB

	nextIDNoteDB uint

	playerDBs map[uint]*PlayerDB

	nextIDPlayerDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		freqencyDBs: make(map[uint]*FreqencyDB),

		noteDBs: make(map[uint]*NoteDB),

		playerDBs: make(map[uint]*PlayerDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/tone/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *FreqencyDB:
		db.nextIDFreqencyDB++
		v.ID = db.nextIDFreqencyDB
		db.freqencyDBs[v.ID] = v
	case *NoteDB:
		db.nextIDNoteDB++
		v.ID = db.nextIDNoteDB
		db.noteDBs[v.ID] = v
	case *PlayerDB:
		db.nextIDPlayerDB++
		v.ID = db.nextIDPlayerDB
		db.playerDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/tone/go, unsupported type in Create")
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
		return nil, errors.New("github.com/fullstack-lang/gong/lib/tone/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *FreqencyDB:
		delete(db.freqencyDBs, v.ID)
	case *NoteDB:
		delete(db.noteDBs, v.ID)
	case *PlayerDB:
		delete(db.playerDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/tone/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/tone/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *FreqencyDB:
		db.freqencyDBs[v.ID] = v
		return db, nil
	case *NoteDB:
		db.noteDBs[v.ID] = v
		return db, nil
	case *PlayerDB:
		db.playerDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/tone/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/tone/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *FreqencyDB:
		if existing, ok := db.freqencyDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Freqency github.com/fullstack-lang/gong/lib/tone/go, record not found")
		}
	case *NoteDB:
		if existing, ok := db.noteDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Note github.com/fullstack-lang/gong/lib/tone/go, record not found")
		}
	case *PlayerDB:
		if existing, ok := db.playerDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Player github.com/fullstack-lang/gong/lib/tone/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/tone/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]FreqencyDB:
		*ptr = make([]FreqencyDB, 0, len(db.freqencyDBs))
		for _, v := range db.freqencyDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]NoteDB:
		*ptr = make([]NoteDB, 0, len(db.noteDBs))
		for _, v := range db.noteDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]PlayerDB:
		*ptr = make([]PlayerDB, 0, len(db.playerDBs))
		for _, v := range db.playerDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/tone/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/tone/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gong/lib/tone/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/tone/go, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *FreqencyDB:
		tmp, ok := db.freqencyDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Freqency Unkown entry %d", i))
		}

		freqencyDB, _ := instanceDB.(*FreqencyDB)
		*freqencyDB = *tmp
		
	case *NoteDB:
		tmp, ok := db.noteDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Note Unkown entry %d", i))
		}

		noteDB, _ := instanceDB.(*NoteDB)
		*noteDB = *tmp
		
	case *PlayerDB:
		tmp, ok := db.playerDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Player Unkown entry %d", i))
		}

		playerDB, _ := instanceDB.(*PlayerDB)
		*playerDB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/tone/go, Unkown type")
	}
	
	return db, nil
}

