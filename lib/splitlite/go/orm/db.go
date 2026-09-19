// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gong/lib/splitlite/go/db"
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

	faviconDBs map[uint]*FavIconDB

	nextIDFavIconDB uint

	formDBs map[uint]*FormDB

	nextIDFormDB uint

	loadDBs map[uint]*LoadDB

	nextIDLoadDB uint

	logoontheleftDBs map[uint]*LogoOnTheLeftDB

	nextIDLogoOnTheLeftDB uint

	logoontherightDBs map[uint]*LogoOnTheRightDB

	nextIDLogoOnTheRightDB uint

	splitDBs map[uint]*SplitDB

	nextIDSplitDB uint

	svgDBs map[uint]*SvgDB

	nextIDSvgDB uint

	tableDBs map[uint]*TableDB

	nextIDTableDB uint

	titleDBs map[uint]*TitleDB

	nextIDTitleDB uint

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

		faviconDBs: make(map[uint]*FavIconDB),

		formDBs: make(map[uint]*FormDB),

		loadDBs: make(map[uint]*LoadDB),

		logoontheleftDBs: make(map[uint]*LogoOnTheLeftDB),

		logoontherightDBs: make(map[uint]*LogoOnTheRightDB),

		splitDBs: make(map[uint]*SplitDB),

		svgDBs: make(map[uint]*SvgDB),

		tableDBs: make(map[uint]*TableDB),

		titleDBs: make(map[uint]*TitleDB),

		treeDBs: make(map[uint]*TreeDB),

		viewDBs: make(map[uint]*ViewDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/splitlite/go, instanceDB cannot be nil")
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
	case *FavIconDB:
		db.nextIDFavIconDB++
		v.ID = db.nextIDFavIconDB
		db.faviconDBs[v.ID] = v
	case *FormDB:
		db.nextIDFormDB++
		v.ID = db.nextIDFormDB
		db.formDBs[v.ID] = v
	case *LoadDB:
		db.nextIDLoadDB++
		v.ID = db.nextIDLoadDB
		db.loadDBs[v.ID] = v
	case *LogoOnTheLeftDB:
		db.nextIDLogoOnTheLeftDB++
		v.ID = db.nextIDLogoOnTheLeftDB
		db.logoontheleftDBs[v.ID] = v
	case *LogoOnTheRightDB:
		db.nextIDLogoOnTheRightDB++
		v.ID = db.nextIDLogoOnTheRightDB
		db.logoontherightDBs[v.ID] = v
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
	case *TitleDB:
		db.nextIDTitleDB++
		v.ID = db.nextIDTitleDB
		db.titleDBs[v.ID] = v
	case *TreeDB:
		db.nextIDTreeDB++
		v.ID = db.nextIDTreeDB
		db.treeDBs[v.ID] = v
	case *ViewDB:
		db.nextIDViewDB++
		v.ID = db.nextIDViewDB
		db.viewDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/splitlite/go, unsupported type in Create")
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
		return nil, errors.New("github.com/fullstack-lang/gong/lib/splitlite/go, instanceDB cannot be nil")
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
	case *FavIconDB:
		delete(db.faviconDBs, v.ID)
	case *FormDB:
		delete(db.formDBs, v.ID)
	case *LoadDB:
		delete(db.loadDBs, v.ID)
	case *LogoOnTheLeftDB:
		delete(db.logoontheleftDBs, v.ID)
	case *LogoOnTheRightDB:
		delete(db.logoontherightDBs, v.ID)
	case *SplitDB:
		delete(db.splitDBs, v.ID)
	case *SvgDB:
		delete(db.svgDBs, v.ID)
	case *TableDB:
		delete(db.tableDBs, v.ID)
	case *TitleDB:
		delete(db.titleDBs, v.ID)
	case *TreeDB:
		delete(db.treeDBs, v.ID)
	case *ViewDB:
		delete(db.viewDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/splitlite/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/splitlite/go, instanceDB cannot be nil")
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
	case *FavIconDB:
		db.faviconDBs[v.ID] = v
		return db, nil
	case *FormDB:
		db.formDBs[v.ID] = v
		return db, nil
	case *LoadDB:
		db.loadDBs[v.ID] = v
		return db, nil
	case *LogoOnTheLeftDB:
		db.logoontheleftDBs[v.ID] = v
		return db, nil
	case *LogoOnTheRightDB:
		db.logoontherightDBs[v.ID] = v
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
	case *TitleDB:
		db.titleDBs[v.ID] = v
		return db, nil
	case *TreeDB:
		db.treeDBs[v.ID] = v
		return db, nil
	case *ViewDB:
		db.viewDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/splitlite/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/splitlite/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AsSplitDB:
		if existing, ok := db.assplitDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db AsSplit github.com/fullstack-lang/gong/lib/splitlite/go, record not found")
		}
	case *AsSplitAreaDB:
		if existing, ok := db.assplitareaDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db AsSplitArea github.com/fullstack-lang/gong/lib/splitlite/go, record not found")
		}
	case *ButtonDB:
		if existing, ok := db.buttonDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Button github.com/fullstack-lang/gong/lib/splitlite/go, record not found")
		}
	case *FavIconDB:
		if existing, ok := db.faviconDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FavIcon github.com/fullstack-lang/gong/lib/splitlite/go, record not found")
		}
	case *FormDB:
		if existing, ok := db.formDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Form github.com/fullstack-lang/gong/lib/splitlite/go, record not found")
		}
	case *LoadDB:
		if existing, ok := db.loadDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Load github.com/fullstack-lang/gong/lib/splitlite/go, record not found")
		}
	case *LogoOnTheLeftDB:
		if existing, ok := db.logoontheleftDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db LogoOnTheLeft github.com/fullstack-lang/gong/lib/splitlite/go, record not found")
		}
	case *LogoOnTheRightDB:
		if existing, ok := db.logoontherightDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db LogoOnTheRight github.com/fullstack-lang/gong/lib/splitlite/go, record not found")
		}
	case *SplitDB:
		if existing, ok := db.splitDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Split github.com/fullstack-lang/gong/lib/splitlite/go, record not found")
		}
	case *SvgDB:
		if existing, ok := db.svgDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Svg github.com/fullstack-lang/gong/lib/splitlite/go, record not found")
		}
	case *TableDB:
		if existing, ok := db.tableDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Table github.com/fullstack-lang/gong/lib/splitlite/go, record not found")
		}
	case *TitleDB:
		if existing, ok := db.titleDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Title github.com/fullstack-lang/gong/lib/splitlite/go, record not found")
		}
	case *TreeDB:
		if existing, ok := db.treeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Tree github.com/fullstack-lang/gong/lib/splitlite/go, record not found")
		}
	case *ViewDB:
		if existing, ok := db.viewDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db View github.com/fullstack-lang/gong/lib/splitlite/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/splitlite/go, unsupported type in Updates")
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
	case *[]FavIconDB:
		*ptr = make([]FavIconDB, 0, len(db.faviconDBs))
		for _, v := range db.faviconDBs {
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
	case *[]LogoOnTheLeftDB:
		*ptr = make([]LogoOnTheLeftDB, 0, len(db.logoontheleftDBs))
		for _, v := range db.logoontheleftDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]LogoOnTheRightDB:
		*ptr = make([]LogoOnTheRightDB, 0, len(db.logoontherightDBs))
		for _, v := range db.logoontherightDBs {
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
	case *[]TitleDB:
		*ptr = make([]TitleDB, 0, len(db.titleDBs))
		for _, v := range db.titleDBs {
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
		return nil, errors.New("github.com/fullstack-lang/gong/lib/splitlite/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/splitlite/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gong/lib/splitlite/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/splitlite/go, conds[0] is not a string or uint64")
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

	case *FavIconDB:
		tmp, ok := db.faviconDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First FavIcon Unkown entry %d", i))
		}

		faviconDB, _ := instanceDB.(*FavIconDB)
		*faviconDB = *tmp

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

	case *LogoOnTheLeftDB:
		tmp, ok := db.logoontheleftDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First LogoOnTheLeft Unkown entry %d", i))
		}

		logoontheleftDB, _ := instanceDB.(*LogoOnTheLeftDB)
		*logoontheleftDB = *tmp

	case *LogoOnTheRightDB:
		tmp, ok := db.logoontherightDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First LogoOnTheRight Unkown entry %d", i))
		}

		logoontherightDB, _ := instanceDB.(*LogoOnTheRightDB)
		*logoontherightDB = *tmp

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

	case *TitleDB:
		tmp, ok := db.titleDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Title Unkown entry %d", i))
		}

		titleDB, _ := instanceDB.(*TitleDB)
		*titleDB = *tmp

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
		return nil, errors.New("github.com/fullstack-lang/gong/lib/splitlite/go, Unkown type")
	}

	return db, nil
}
