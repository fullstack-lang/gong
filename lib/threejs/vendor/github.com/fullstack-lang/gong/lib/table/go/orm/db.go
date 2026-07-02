// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gong/lib/table/go/db"
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

	cellDBs map[uint]*CellDB

	nextIDCellDB uint

	cellbooleanDBs map[uint]*CellBooleanDB

	nextIDCellBooleanDB uint

	cellfloat64DBs map[uint]*CellFloat64DB

	nextIDCellFloat64DB uint

	celliconDBs map[uint]*CellIconDB

	nextIDCellIconDB uint

	cellintDBs map[uint]*CellIntDB

	nextIDCellIntDB uint

	cellstringDBs map[uint]*CellStringDB

	nextIDCellStringDB uint

	displayedcolumnDBs map[uint]*DisplayedColumnDB

	nextIDDisplayedColumnDB uint

	rowDBs map[uint]*RowDB

	nextIDRowDB uint

	svgiconDBs map[uint]*SVGIconDB

	nextIDSVGIconDB uint

	tableDBs map[uint]*TableDB

	nextIDTableDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		buttonDBs: make(map[uint]*ButtonDB),

		cellDBs: make(map[uint]*CellDB),

		cellbooleanDBs: make(map[uint]*CellBooleanDB),

		cellfloat64DBs: make(map[uint]*CellFloat64DB),

		celliconDBs: make(map[uint]*CellIconDB),

		cellintDBs: make(map[uint]*CellIntDB),

		cellstringDBs: make(map[uint]*CellStringDB),

		displayedcolumnDBs: make(map[uint]*DisplayedColumnDB),

		rowDBs: make(map[uint]*RowDB),

		svgiconDBs: make(map[uint]*SVGIconDB),

		tableDBs: make(map[uint]*TableDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/table/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *ButtonDB:
		db.nextIDButtonDB++
		v.ID = db.nextIDButtonDB
		db.buttonDBs[v.ID] = v
	case *CellDB:
		db.nextIDCellDB++
		v.ID = db.nextIDCellDB
		db.cellDBs[v.ID] = v
	case *CellBooleanDB:
		db.nextIDCellBooleanDB++
		v.ID = db.nextIDCellBooleanDB
		db.cellbooleanDBs[v.ID] = v
	case *CellFloat64DB:
		db.nextIDCellFloat64DB++
		v.ID = db.nextIDCellFloat64DB
		db.cellfloat64DBs[v.ID] = v
	case *CellIconDB:
		db.nextIDCellIconDB++
		v.ID = db.nextIDCellIconDB
		db.celliconDBs[v.ID] = v
	case *CellIntDB:
		db.nextIDCellIntDB++
		v.ID = db.nextIDCellIntDB
		db.cellintDBs[v.ID] = v
	case *CellStringDB:
		db.nextIDCellStringDB++
		v.ID = db.nextIDCellStringDB
		db.cellstringDBs[v.ID] = v
	case *DisplayedColumnDB:
		db.nextIDDisplayedColumnDB++
		v.ID = db.nextIDDisplayedColumnDB
		db.displayedcolumnDBs[v.ID] = v
	case *RowDB:
		db.nextIDRowDB++
		v.ID = db.nextIDRowDB
		db.rowDBs[v.ID] = v
	case *SVGIconDB:
		db.nextIDSVGIconDB++
		v.ID = db.nextIDSVGIconDB
		db.svgiconDBs[v.ID] = v
	case *TableDB:
		db.nextIDTableDB++
		v.ID = db.nextIDTableDB
		db.tableDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/table/go, unsupported type in Create")
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
		return nil, errors.New("github.com/fullstack-lang/gong/lib/table/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ButtonDB:
		delete(db.buttonDBs, v.ID)
	case *CellDB:
		delete(db.cellDBs, v.ID)
	case *CellBooleanDB:
		delete(db.cellbooleanDBs, v.ID)
	case *CellFloat64DB:
		delete(db.cellfloat64DBs, v.ID)
	case *CellIconDB:
		delete(db.celliconDBs, v.ID)
	case *CellIntDB:
		delete(db.cellintDBs, v.ID)
	case *CellStringDB:
		delete(db.cellstringDBs, v.ID)
	case *DisplayedColumnDB:
		delete(db.displayedcolumnDBs, v.ID)
	case *RowDB:
		delete(db.rowDBs, v.ID)
	case *SVGIconDB:
		delete(db.svgiconDBs, v.ID)
	case *TableDB:
		delete(db.tableDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/table/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/table/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ButtonDB:
		db.buttonDBs[v.ID] = v
		return db, nil
	case *CellDB:
		db.cellDBs[v.ID] = v
		return db, nil
	case *CellBooleanDB:
		db.cellbooleanDBs[v.ID] = v
		return db, nil
	case *CellFloat64DB:
		db.cellfloat64DBs[v.ID] = v
		return db, nil
	case *CellIconDB:
		db.celliconDBs[v.ID] = v
		return db, nil
	case *CellIntDB:
		db.cellintDBs[v.ID] = v
		return db, nil
	case *CellStringDB:
		db.cellstringDBs[v.ID] = v
		return db, nil
	case *DisplayedColumnDB:
		db.displayedcolumnDBs[v.ID] = v
		return db, nil
	case *RowDB:
		db.rowDBs[v.ID] = v
		return db, nil
	case *SVGIconDB:
		db.svgiconDBs[v.ID] = v
		return db, nil
	case *TableDB:
		db.tableDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/table/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/table/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ButtonDB:
		if existing, ok := db.buttonDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Button github.com/fullstack-lang/gong/lib/table/go, record not found")
		}
	case *CellDB:
		if existing, ok := db.cellDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Cell github.com/fullstack-lang/gong/lib/table/go, record not found")
		}
	case *CellBooleanDB:
		if existing, ok := db.cellbooleanDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db CellBoolean github.com/fullstack-lang/gong/lib/table/go, record not found")
		}
	case *CellFloat64DB:
		if existing, ok := db.cellfloat64DBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db CellFloat64 github.com/fullstack-lang/gong/lib/table/go, record not found")
		}
	case *CellIconDB:
		if existing, ok := db.celliconDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db CellIcon github.com/fullstack-lang/gong/lib/table/go, record not found")
		}
	case *CellIntDB:
		if existing, ok := db.cellintDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db CellInt github.com/fullstack-lang/gong/lib/table/go, record not found")
		}
	case *CellStringDB:
		if existing, ok := db.cellstringDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db CellString github.com/fullstack-lang/gong/lib/table/go, record not found")
		}
	case *DisplayedColumnDB:
		if existing, ok := db.displayedcolumnDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DisplayedColumn github.com/fullstack-lang/gong/lib/table/go, record not found")
		}
	case *RowDB:
		if existing, ok := db.rowDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Row github.com/fullstack-lang/gong/lib/table/go, record not found")
		}
	case *SVGIconDB:
		if existing, ok := db.svgiconDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SVGIcon github.com/fullstack-lang/gong/lib/table/go, record not found")
		}
	case *TableDB:
		if existing, ok := db.tableDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Table github.com/fullstack-lang/gong/lib/table/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/table/go, unsupported type in Updates")
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
	case *[]CellDB:
		*ptr = make([]CellDB, 0, len(db.cellDBs))
		for _, v := range db.cellDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]CellBooleanDB:
		*ptr = make([]CellBooleanDB, 0, len(db.cellbooleanDBs))
		for _, v := range db.cellbooleanDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]CellFloat64DB:
		*ptr = make([]CellFloat64DB, 0, len(db.cellfloat64DBs))
		for _, v := range db.cellfloat64DBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]CellIconDB:
		*ptr = make([]CellIconDB, 0, len(db.celliconDBs))
		for _, v := range db.celliconDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]CellIntDB:
		*ptr = make([]CellIntDB, 0, len(db.cellintDBs))
		for _, v := range db.cellintDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]CellStringDB:
		*ptr = make([]CellStringDB, 0, len(db.cellstringDBs))
		for _, v := range db.cellstringDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DisplayedColumnDB:
		*ptr = make([]DisplayedColumnDB, 0, len(db.displayedcolumnDBs))
		for _, v := range db.displayedcolumnDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]RowDB:
		*ptr = make([]RowDB, 0, len(db.rowDBs))
		for _, v := range db.rowDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SVGIconDB:
		*ptr = make([]SVGIconDB, 0, len(db.svgiconDBs))
		for _, v := range db.svgiconDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]TableDB:
		*ptr = make([]TableDB, 0, len(db.tableDBs))
		for _, v := range db.tableDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/table/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/table/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gong/lib/table/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/table/go, conds[0] is not a string or uint64")
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

	case *CellDB:
		tmp, ok := db.cellDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Cell Unkown entry %d", i))
		}

		cellDB, _ := instanceDB.(*CellDB)
		*cellDB = *tmp

	case *CellBooleanDB:
		tmp, ok := db.cellbooleanDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First CellBoolean Unkown entry %d", i))
		}

		cellbooleanDB, _ := instanceDB.(*CellBooleanDB)
		*cellbooleanDB = *tmp

	case *CellFloat64DB:
		tmp, ok := db.cellfloat64DBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First CellFloat64 Unkown entry %d", i))
		}

		cellfloat64DB, _ := instanceDB.(*CellFloat64DB)
		*cellfloat64DB = *tmp

	case *CellIconDB:
		tmp, ok := db.celliconDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First CellIcon Unkown entry %d", i))
		}

		celliconDB, _ := instanceDB.(*CellIconDB)
		*celliconDB = *tmp

	case *CellIntDB:
		tmp, ok := db.cellintDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First CellInt Unkown entry %d", i))
		}

		cellintDB, _ := instanceDB.(*CellIntDB)
		*cellintDB = *tmp

	case *CellStringDB:
		tmp, ok := db.cellstringDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First CellString Unkown entry %d", i))
		}

		cellstringDB, _ := instanceDB.(*CellStringDB)
		*cellstringDB = *tmp

	case *DisplayedColumnDB:
		tmp, ok := db.displayedcolumnDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DisplayedColumn Unkown entry %d", i))
		}

		displayedcolumnDB, _ := instanceDB.(*DisplayedColumnDB)
		*displayedcolumnDB = *tmp

	case *RowDB:
		tmp, ok := db.rowDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Row Unkown entry %d", i))
		}

		rowDB, _ := instanceDB.(*RowDB)
		*rowDB = *tmp

	case *SVGIconDB:
		tmp, ok := db.svgiconDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SVGIcon Unkown entry %d", i))
		}

		svgiconDB, _ := instanceDB.(*SVGIconDB)
		*svgiconDB = *tmp

	case *TableDB:
		tmp, ok := db.tableDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Table Unkown entry %d", i))
		}

		tableDB, _ := instanceDB.(*TableDB)
		*tableDB = *tmp

	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/table/go, Unkown type")
	}

	return db, nil
}
