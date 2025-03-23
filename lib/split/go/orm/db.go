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

	buttonDBs map[uint]*ButtonDB

	nextIDButtonDB uint

	cursorDBs map[uint]*CursorDB

	nextIDCursorDB uint

	docDBs map[uint]*DocDB

	nextIDDocDB uint

	formDBs map[uint]*FormDB

	nextIDFormDB uint

	loadDBs map[uint]*LoadDB

	nextIDLoadDB uint

	sliderDBs map[uint]*SliderDB

	nextIDSliderDB uint

	splitDBs map[uint]*SplitDB

	nextIDSplitDB uint

	svgDBs map[uint]*SvgDB

	nextIDSvgDB uint

	tableDBs map[uint]*TableDB

	nextIDTableDB uint

	toneDBs map[uint]*ToneDB

	nextIDToneDB uint

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

		buttonDBs: make(map[uint]*ButtonDB),

		cursorDBs: make(map[uint]*CursorDB),

		docDBs: make(map[uint]*DocDB),

		formDBs: make(map[uint]*FormDB),

		loadDBs: make(map[uint]*LoadDB),

		sliderDBs: make(map[uint]*SliderDB),

		splitDBs: make(map[uint]*SplitDB),

		svgDBs: make(map[uint]*SvgDB),

		tableDBs: make(map[uint]*TableDB),

		toneDBs: make(map[uint]*ToneDB),

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
	case *ButtonDB:
		db.nextIDButtonDB++
		v.ID = db.nextIDButtonDB
		db.buttonDBs[v.ID] = v
	case *CursorDB:
		db.nextIDCursorDB++
		v.ID = db.nextIDCursorDB
		db.cursorDBs[v.ID] = v
	case *DocDB:
		db.nextIDDocDB++
		v.ID = db.nextIDDocDB
		db.docDBs[v.ID] = v
	case *FormDB:
		db.nextIDFormDB++
		v.ID = db.nextIDFormDB
		db.formDBs[v.ID] = v
	case *LoadDB:
		db.nextIDLoadDB++
		v.ID = db.nextIDLoadDB
		db.loadDBs[v.ID] = v
	case *SliderDB:
		db.nextIDSliderDB++
		v.ID = db.nextIDSliderDB
		db.sliderDBs[v.ID] = v
	case *SplitDB:
		db.nextIDSplitDB++
		v.ID = db.nextIDSplitDB
		db.splitDBs[v.ID] = v
	case *SvgDB:
		db.nextIDSvgDB++
		v.ID = db.nextIDSvgDB
		db.svgDBs[v.ID] = v
	case *TableDB:
		db.nextIDTableDB++
		v.ID = db.nextIDTableDB
		db.tableDBs[v.ID] = v
	case *ToneDB:
		db.nextIDToneDB++
		v.ID = db.nextIDToneDB
		db.toneDBs[v.ID] = v
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
	case *ButtonDB:
		delete(db.buttonDBs, v.ID)
	case *CursorDB:
		delete(db.cursorDBs, v.ID)
	case *DocDB:
		delete(db.docDBs, v.ID)
	case *FormDB:
		delete(db.formDBs, v.ID)
	case *LoadDB:
		delete(db.loadDBs, v.ID)
	case *SliderDB:
		delete(db.sliderDBs, v.ID)
	case *SplitDB:
		delete(db.splitDBs, v.ID)
	case *SvgDB:
		delete(db.svgDBs, v.ID)
	case *TableDB:
		delete(db.tableDBs, v.ID)
	case *ToneDB:
		delete(db.toneDBs, v.ID)
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
	case *ButtonDB:
		db.buttonDBs[v.ID] = v
		return db, nil
	case *CursorDB:
		db.cursorDBs[v.ID] = v
		return db, nil
	case *DocDB:
		db.docDBs[v.ID] = v
		return db, nil
	case *FormDB:
		db.formDBs[v.ID] = v
		return db, nil
	case *LoadDB:
		db.loadDBs[v.ID] = v
		return db, nil
	case *SliderDB:
		db.sliderDBs[v.ID] = v
		return db, nil
	case *SplitDB:
		db.splitDBs[v.ID] = v
		return db, nil
	case *SvgDB:
		db.svgDBs[v.ID] = v
		return db, nil
	case *TableDB:
		db.tableDBs[v.ID] = v
		return db, nil
	case *ToneDB:
		db.toneDBs[v.ID] = v
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
	case *ButtonDB:
		if existing, ok := db.buttonDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Button github.com/fullstack-lang/gong/lib/split/go, record not found")
		}
	case *CursorDB:
		if existing, ok := db.cursorDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Cursor github.com/fullstack-lang/gong/lib/split/go, record not found")
		}
	case *DocDB:
		if existing, ok := db.docDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Doc github.com/fullstack-lang/gong/lib/split/go, record not found")
		}
	case *FormDB:
		if existing, ok := db.formDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Form github.com/fullstack-lang/gong/lib/split/go, record not found")
		}
	case *LoadDB:
		if existing, ok := db.loadDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Load github.com/fullstack-lang/gong/lib/split/go, record not found")
		}
	case *SliderDB:
		if existing, ok := db.sliderDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Slider github.com/fullstack-lang/gong/lib/split/go, record not found")
		}
	case *SplitDB:
		if existing, ok := db.splitDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Split github.com/fullstack-lang/gong/lib/split/go, record not found")
		}
	case *SvgDB:
		if existing, ok := db.svgDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Svg github.com/fullstack-lang/gong/lib/split/go, record not found")
		}
	case *TableDB:
		if existing, ok := db.tableDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Table github.com/fullstack-lang/gong/lib/split/go, record not found")
		}
	case *ToneDB:
		if existing, ok := db.toneDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Tone github.com/fullstack-lang/gong/lib/split/go, record not found")
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
	case *[]ButtonDB:
		*ptr = make([]ButtonDB, 0, len(db.buttonDBs))
		for _, v := range db.buttonDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]CursorDB:
		*ptr = make([]CursorDB, 0, len(db.cursorDBs))
		for _, v := range db.cursorDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DocDB:
		*ptr = make([]DocDB, 0, len(db.docDBs))
		for _, v := range db.docDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]FormDB:
		*ptr = make([]FormDB, 0, len(db.formDBs))
		for _, v := range db.formDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]LoadDB:
		*ptr = make([]LoadDB, 0, len(db.loadDBs))
		for _, v := range db.loadDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SliderDB:
		*ptr = make([]SliderDB, 0, len(db.sliderDBs))
		for _, v := range db.sliderDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SplitDB:
		*ptr = make([]SplitDB, 0, len(db.splitDBs))
		for _, v := range db.splitDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SvgDB:
		*ptr = make([]SvgDB, 0, len(db.svgDBs))
		for _, v := range db.svgDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]TableDB:
		*ptr = make([]TableDB, 0, len(db.tableDBs))
		for _, v := range db.tableDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ToneDB:
		*ptr = make([]ToneDB, 0, len(db.toneDBs))
		for _, v := range db.toneDBs {
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
		
	case *ButtonDB:
		tmp, ok := db.buttonDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Button Unkown entry %d", i))
		}

		buttonDB, _ := instanceDB.(*ButtonDB)
		*buttonDB = *tmp
		
	case *CursorDB:
		tmp, ok := db.cursorDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Cursor Unkown entry %d", i))
		}

		cursorDB, _ := instanceDB.(*CursorDB)
		*cursorDB = *tmp
		
	case *DocDB:
		tmp, ok := db.docDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Doc Unkown entry %d", i))
		}

		docDB, _ := instanceDB.(*DocDB)
		*docDB = *tmp
		
	case *FormDB:
		tmp, ok := db.formDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Form Unkown entry %d", i))
		}

		formDB, _ := instanceDB.(*FormDB)
		*formDB = *tmp
		
	case *LoadDB:
		tmp, ok := db.loadDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Load Unkown entry %d", i))
		}

		loadDB, _ := instanceDB.(*LoadDB)
		*loadDB = *tmp
		
	case *SliderDB:
		tmp, ok := db.sliderDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Slider Unkown entry %d", i))
		}

		sliderDB, _ := instanceDB.(*SliderDB)
		*sliderDB = *tmp
		
	case *SplitDB:
		tmp, ok := db.splitDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Split Unkown entry %d", i))
		}

		splitDB, _ := instanceDB.(*SplitDB)
		*splitDB = *tmp
		
	case *SvgDB:
		tmp, ok := db.svgDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Svg Unkown entry %d", i))
		}

		svgDB, _ := instanceDB.(*SvgDB)
		*svgDB = *tmp
		
	case *TableDB:
		tmp, ok := db.tableDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Table Unkown entry %d", i))
		}

		tableDB, _ := instanceDB.(*TableDB)
		*tableDB = *tmp
		
	case *ToneDB:
		tmp, ok := db.toneDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Tone Unkown entry %d", i))
		}

		toneDB, _ := instanceDB.(*ToneDB)
		*toneDB = *tmp
		
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

