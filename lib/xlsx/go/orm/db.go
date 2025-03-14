// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gong/lib/xlsx/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	displayselectionDBs map[uint]*DisplaySelectionDB

	nextIDDisplaySelectionDB uint

	xlcellDBs map[uint]*XLCellDB

	nextIDXLCellDB uint

	xlfileDBs map[uint]*XLFileDB

	nextIDXLFileDB uint

	xlrowDBs map[uint]*XLRowDB

	nextIDXLRowDB uint

	xlsheetDBs map[uint]*XLSheetDB

	nextIDXLSheetDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		displayselectionDBs: make(map[uint]*DisplaySelectionDB),

		xlcellDBs: make(map[uint]*XLCellDB),

		xlfileDBs: make(map[uint]*XLFileDB),

		xlrowDBs: make(map[uint]*XLRowDB),

		xlsheetDBs: make(map[uint]*XLSheetDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/xlsx/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *DisplaySelectionDB:
		db.nextIDDisplaySelectionDB++
		v.ID = db.nextIDDisplaySelectionDB
		db.displayselectionDBs[v.ID] = v
	case *XLCellDB:
		db.nextIDXLCellDB++
		v.ID = db.nextIDXLCellDB
		db.xlcellDBs[v.ID] = v
	case *XLFileDB:
		db.nextIDXLFileDB++
		v.ID = db.nextIDXLFileDB
		db.xlfileDBs[v.ID] = v
	case *XLRowDB:
		db.nextIDXLRowDB++
		v.ID = db.nextIDXLRowDB
		db.xlrowDBs[v.ID] = v
	case *XLSheetDB:
		db.nextIDXLSheetDB++
		v.ID = db.nextIDXLSheetDB
		db.xlsheetDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/xlsx/go, unsupported type in Create")
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
		return nil, errors.New("github.com/fullstack-lang/gong/lib/xlsx/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *DisplaySelectionDB:
		delete(db.displayselectionDBs, v.ID)
	case *XLCellDB:
		delete(db.xlcellDBs, v.ID)
	case *XLFileDB:
		delete(db.xlfileDBs, v.ID)
	case *XLRowDB:
		delete(db.xlrowDBs, v.ID)
	case *XLSheetDB:
		delete(db.xlsheetDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/xlsx/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/xlsx/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *DisplaySelectionDB:
		db.displayselectionDBs[v.ID] = v
		return db, nil
	case *XLCellDB:
		db.xlcellDBs[v.ID] = v
		return db, nil
	case *XLFileDB:
		db.xlfileDBs[v.ID] = v
		return db, nil
	case *XLRowDB:
		db.xlrowDBs[v.ID] = v
		return db, nil
	case *XLSheetDB:
		db.xlsheetDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/xlsx/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/xlsx/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *DisplaySelectionDB:
		if existing, ok := db.displayselectionDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DisplaySelection github.com/fullstack-lang/gong/lib/xlsx/go, record not found")
		}
	case *XLCellDB:
		if existing, ok := db.xlcellDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db XLCell github.com/fullstack-lang/gong/lib/xlsx/go, record not found")
		}
	case *XLFileDB:
		if existing, ok := db.xlfileDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db XLFile github.com/fullstack-lang/gong/lib/xlsx/go, record not found")
		}
	case *XLRowDB:
		if existing, ok := db.xlrowDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db XLRow github.com/fullstack-lang/gong/lib/xlsx/go, record not found")
		}
	case *XLSheetDB:
		if existing, ok := db.xlsheetDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db XLSheet github.com/fullstack-lang/gong/lib/xlsx/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/xlsx/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]DisplaySelectionDB:
		*ptr = make([]DisplaySelectionDB, 0, len(db.displayselectionDBs))
		for _, v := range db.displayselectionDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]XLCellDB:
		*ptr = make([]XLCellDB, 0, len(db.xlcellDBs))
		for _, v := range db.xlcellDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]XLFileDB:
		*ptr = make([]XLFileDB, 0, len(db.xlfileDBs))
		for _, v := range db.xlfileDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]XLRowDB:
		*ptr = make([]XLRowDB, 0, len(db.xlrowDBs))
		for _, v := range db.xlrowDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]XLSheetDB:
		*ptr = make([]XLSheetDB, 0, len(db.xlsheetDBs))
		for _, v := range db.xlsheetDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/xlsx/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/xlsx/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gong/lib/xlsx/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/xlsx/go, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *DisplaySelectionDB:
		tmp, ok := db.displayselectionDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DisplaySelection Unkown entry %d", i))
		}

		displayselectionDB, _ := instanceDB.(*DisplaySelectionDB)
		*displayselectionDB = *tmp
		
	case *XLCellDB:
		tmp, ok := db.xlcellDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First XLCell Unkown entry %d", i))
		}

		xlcellDB, _ := instanceDB.(*XLCellDB)
		*xlcellDB = *tmp
		
	case *XLFileDB:
		tmp, ok := db.xlfileDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First XLFile Unkown entry %d", i))
		}

		xlfileDB, _ := instanceDB.(*XLFileDB)
		*xlfileDB = *tmp
		
	case *XLRowDB:
		tmp, ok := db.xlrowDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First XLRow Unkown entry %d", i))
		}

		xlrowDB, _ := instanceDB.(*XLRowDB)
		*xlrowDB = *tmp
		
	case *XLSheetDB:
		tmp, ok := db.xlsheetDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First XLSheet Unkown entry %d", i))
		}

		xlsheetDB, _ := instanceDB.(*XLSheetDB)
		*xlsheetDB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/xlsx/go, Unkown type")
	}
	
	return db, nil
}

