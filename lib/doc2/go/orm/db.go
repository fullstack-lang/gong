// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gong/lib/doc2/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	attributeshapeDBs map[uint]*AttributeShapeDB

	nextIDAttributeShapeDB uint

	classdiagramDBs map[uint]*ClassdiagramDB

	nextIDClassdiagramDB uint

	diagrampackageDBs map[uint]*DiagramPackageDB

	nextIDDiagramPackageDB uint

	gongenumshapeDBs map[uint]*GongEnumShapeDB

	nextIDGongEnumShapeDB uint

	gongenumvalueshapeDBs map[uint]*GongEnumValueShapeDB

	nextIDGongEnumValueShapeDB uint

	gongnotelinkshapeDBs map[uint]*GongNoteLinkShapeDB

	nextIDGongNoteLinkShapeDB uint

	gongnoteshapeDBs map[uint]*GongNoteShapeDB

	nextIDGongNoteShapeDB uint

	gongstructshapeDBs map[uint]*GongStructShapeDB

	nextIDGongStructShapeDB uint

	linkshapeDBs map[uint]*LinkShapeDB

	nextIDLinkShapeDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		attributeshapeDBs: make(map[uint]*AttributeShapeDB),

		classdiagramDBs: make(map[uint]*ClassdiagramDB),

		diagrampackageDBs: make(map[uint]*DiagramPackageDB),

		gongenumshapeDBs: make(map[uint]*GongEnumShapeDB),

		gongenumvalueshapeDBs: make(map[uint]*GongEnumValueShapeDB),

		gongnotelinkshapeDBs: make(map[uint]*GongNoteLinkShapeDB),

		gongnoteshapeDBs: make(map[uint]*GongNoteShapeDB),

		gongstructshapeDBs: make(map[uint]*GongStructShapeDB),

		linkshapeDBs: make(map[uint]*LinkShapeDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/doc2/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *AttributeShapeDB:
		db.nextIDAttributeShapeDB++
		v.ID = db.nextIDAttributeShapeDB
		db.attributeshapeDBs[v.ID] = v
	case *ClassdiagramDB:
		db.nextIDClassdiagramDB++
		v.ID = db.nextIDClassdiagramDB
		db.classdiagramDBs[v.ID] = v
	case *DiagramPackageDB:
		db.nextIDDiagramPackageDB++
		v.ID = db.nextIDDiagramPackageDB
		db.diagrampackageDBs[v.ID] = v
	case *GongEnumShapeDB:
		db.nextIDGongEnumShapeDB++
		v.ID = db.nextIDGongEnumShapeDB
		db.gongenumshapeDBs[v.ID] = v
	case *GongEnumValueShapeDB:
		db.nextIDGongEnumValueShapeDB++
		v.ID = db.nextIDGongEnumValueShapeDB
		db.gongenumvalueshapeDBs[v.ID] = v
	case *GongNoteLinkShapeDB:
		db.nextIDGongNoteLinkShapeDB++
		v.ID = db.nextIDGongNoteLinkShapeDB
		db.gongnotelinkshapeDBs[v.ID] = v
	case *GongNoteShapeDB:
		db.nextIDGongNoteShapeDB++
		v.ID = db.nextIDGongNoteShapeDB
		db.gongnoteshapeDBs[v.ID] = v
	case *GongStructShapeDB:
		db.nextIDGongStructShapeDB++
		v.ID = db.nextIDGongStructShapeDB
		db.gongstructshapeDBs[v.ID] = v
	case *LinkShapeDB:
		db.nextIDLinkShapeDB++
		v.ID = db.nextIDLinkShapeDB
		db.linkshapeDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/doc2/go, unsupported type in Create")
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
		return nil, errors.New("github.com/fullstack-lang/gong/lib/doc2/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AttributeShapeDB:
		delete(db.attributeshapeDBs, v.ID)
	case *ClassdiagramDB:
		delete(db.classdiagramDBs, v.ID)
	case *DiagramPackageDB:
		delete(db.diagrampackageDBs, v.ID)
	case *GongEnumShapeDB:
		delete(db.gongenumshapeDBs, v.ID)
	case *GongEnumValueShapeDB:
		delete(db.gongenumvalueshapeDBs, v.ID)
	case *GongNoteLinkShapeDB:
		delete(db.gongnotelinkshapeDBs, v.ID)
	case *GongNoteShapeDB:
		delete(db.gongnoteshapeDBs, v.ID)
	case *GongStructShapeDB:
		delete(db.gongstructshapeDBs, v.ID)
	case *LinkShapeDB:
		delete(db.linkshapeDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/doc2/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/doc2/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AttributeShapeDB:
		db.attributeshapeDBs[v.ID] = v
		return db, nil
	case *ClassdiagramDB:
		db.classdiagramDBs[v.ID] = v
		return db, nil
	case *DiagramPackageDB:
		db.diagrampackageDBs[v.ID] = v
		return db, nil
	case *GongEnumShapeDB:
		db.gongenumshapeDBs[v.ID] = v
		return db, nil
	case *GongEnumValueShapeDB:
		db.gongenumvalueshapeDBs[v.ID] = v
		return db, nil
	case *GongNoteLinkShapeDB:
		db.gongnotelinkshapeDBs[v.ID] = v
		return db, nil
	case *GongNoteShapeDB:
		db.gongnoteshapeDBs[v.ID] = v
		return db, nil
	case *GongStructShapeDB:
		db.gongstructshapeDBs[v.ID] = v
		return db, nil
	case *LinkShapeDB:
		db.linkshapeDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/doc2/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/doc2/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AttributeShapeDB:
		if existing, ok := db.attributeshapeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db AttributeShape github.com/fullstack-lang/gong/lib/doc2/go, record not found")
		}
	case *ClassdiagramDB:
		if existing, ok := db.classdiagramDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Classdiagram github.com/fullstack-lang/gong/lib/doc2/go, record not found")
		}
	case *DiagramPackageDB:
		if existing, ok := db.diagrampackageDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DiagramPackage github.com/fullstack-lang/gong/lib/doc2/go, record not found")
		}
	case *GongEnumShapeDB:
		if existing, ok := db.gongenumshapeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db GongEnumShape github.com/fullstack-lang/gong/lib/doc2/go, record not found")
		}
	case *GongEnumValueShapeDB:
		if existing, ok := db.gongenumvalueshapeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db GongEnumValueShape github.com/fullstack-lang/gong/lib/doc2/go, record not found")
		}
	case *GongNoteLinkShapeDB:
		if existing, ok := db.gongnotelinkshapeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db GongNoteLinkShape github.com/fullstack-lang/gong/lib/doc2/go, record not found")
		}
	case *GongNoteShapeDB:
		if existing, ok := db.gongnoteshapeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db GongNoteShape github.com/fullstack-lang/gong/lib/doc2/go, record not found")
		}
	case *GongStructShapeDB:
		if existing, ok := db.gongstructshapeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db GongStructShape github.com/fullstack-lang/gong/lib/doc2/go, record not found")
		}
	case *LinkShapeDB:
		if existing, ok := db.linkshapeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db LinkShape github.com/fullstack-lang/gong/lib/doc2/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/doc2/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]AttributeShapeDB:
		*ptr = make([]AttributeShapeDB, 0, len(db.attributeshapeDBs))
		for _, v := range db.attributeshapeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ClassdiagramDB:
		*ptr = make([]ClassdiagramDB, 0, len(db.classdiagramDBs))
		for _, v := range db.classdiagramDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DiagramPackageDB:
		*ptr = make([]DiagramPackageDB, 0, len(db.diagrampackageDBs))
		for _, v := range db.diagrampackageDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]GongEnumShapeDB:
		*ptr = make([]GongEnumShapeDB, 0, len(db.gongenumshapeDBs))
		for _, v := range db.gongenumshapeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]GongEnumValueShapeDB:
		*ptr = make([]GongEnumValueShapeDB, 0, len(db.gongenumvalueshapeDBs))
		for _, v := range db.gongenumvalueshapeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]GongNoteLinkShapeDB:
		*ptr = make([]GongNoteLinkShapeDB, 0, len(db.gongnotelinkshapeDBs))
		for _, v := range db.gongnotelinkshapeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]GongNoteShapeDB:
		*ptr = make([]GongNoteShapeDB, 0, len(db.gongnoteshapeDBs))
		for _, v := range db.gongnoteshapeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]GongStructShapeDB:
		*ptr = make([]GongStructShapeDB, 0, len(db.gongstructshapeDBs))
		for _, v := range db.gongstructshapeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]LinkShapeDB:
		*ptr = make([]LinkShapeDB, 0, len(db.linkshapeDBs))
		for _, v := range db.linkshapeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/doc2/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/doc2/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gong/lib/doc2/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/doc2/go, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *AttributeShapeDB:
		tmp, ok := db.attributeshapeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First AttributeShape Unkown entry %d", i))
		}

		attributeshapeDB, _ := instanceDB.(*AttributeShapeDB)
		*attributeshapeDB = *tmp
		
	case *ClassdiagramDB:
		tmp, ok := db.classdiagramDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Classdiagram Unkown entry %d", i))
		}

		classdiagramDB, _ := instanceDB.(*ClassdiagramDB)
		*classdiagramDB = *tmp
		
	case *DiagramPackageDB:
		tmp, ok := db.diagrampackageDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DiagramPackage Unkown entry %d", i))
		}

		diagrampackageDB, _ := instanceDB.(*DiagramPackageDB)
		*diagrampackageDB = *tmp
		
	case *GongEnumShapeDB:
		tmp, ok := db.gongenumshapeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First GongEnumShape Unkown entry %d", i))
		}

		gongenumshapeDB, _ := instanceDB.(*GongEnumShapeDB)
		*gongenumshapeDB = *tmp
		
	case *GongEnumValueShapeDB:
		tmp, ok := db.gongenumvalueshapeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First GongEnumValueShape Unkown entry %d", i))
		}

		gongenumvalueshapeDB, _ := instanceDB.(*GongEnumValueShapeDB)
		*gongenumvalueshapeDB = *tmp
		
	case *GongNoteLinkShapeDB:
		tmp, ok := db.gongnotelinkshapeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First GongNoteLinkShape Unkown entry %d", i))
		}

		gongnotelinkshapeDB, _ := instanceDB.(*GongNoteLinkShapeDB)
		*gongnotelinkshapeDB = *tmp
		
	case *GongNoteShapeDB:
		tmp, ok := db.gongnoteshapeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First GongNoteShape Unkown entry %d", i))
		}

		gongnoteshapeDB, _ := instanceDB.(*GongNoteShapeDB)
		*gongnoteshapeDB = *tmp
		
	case *GongStructShapeDB:
		tmp, ok := db.gongstructshapeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First GongStructShape Unkown entry %d", i))
		}

		gongstructshapeDB, _ := instanceDB.(*GongStructShapeDB)
		*gongstructshapeDB = *tmp
		
	case *LinkShapeDB:
		tmp, ok := db.linkshapeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First LinkShape Unkown entry %d", i))
		}

		linkshapeDB, _ := instanceDB.(*LinkShapeDB)
		*linkshapeDB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/doc2/go, Unkown type")
	}
	
	return db, nil
}

