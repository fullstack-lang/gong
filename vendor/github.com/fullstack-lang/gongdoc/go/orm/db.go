// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gongdoc/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	classdiagramDBs map[uint]*ClassdiagramDB

	nextIDClassdiagramDB uint

	diagrampackageDBs map[uint]*DiagramPackageDB

	nextIDDiagramPackageDB uint

	fieldDBs map[uint]*FieldDB

	nextIDFieldDB uint

	gongenumshapeDBs map[uint]*GongEnumShapeDB

	nextIDGongEnumShapeDB uint

	gongenumvalueentryDBs map[uint]*GongEnumValueEntryDB

	nextIDGongEnumValueEntryDB uint

	gongstructshapeDBs map[uint]*GongStructShapeDB

	nextIDGongStructShapeDB uint

	linkDBs map[uint]*LinkDB

	nextIDLinkDB uint

	noteshapeDBs map[uint]*NoteShapeDB

	nextIDNoteShapeDB uint

	noteshapelinkDBs map[uint]*NoteShapeLinkDB

	nextIDNoteShapeLinkDB uint

	positionDBs map[uint]*PositionDB

	nextIDPositionDB uint

	umlstateDBs map[uint]*UmlStateDB

	nextIDUmlStateDB uint

	umlscDBs map[uint]*UmlscDB

	nextIDUmlscDB uint

	verticeDBs map[uint]*VerticeDB

	nextIDVerticeDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		classdiagramDBs: make(map[uint]*ClassdiagramDB),

		diagrampackageDBs: make(map[uint]*DiagramPackageDB),

		fieldDBs: make(map[uint]*FieldDB),

		gongenumshapeDBs: make(map[uint]*GongEnumShapeDB),

		gongenumvalueentryDBs: make(map[uint]*GongEnumValueEntryDB),

		gongstructshapeDBs: make(map[uint]*GongStructShapeDB),

		linkDBs: make(map[uint]*LinkDB),

		noteshapeDBs: make(map[uint]*NoteShapeDB),

		noteshapelinkDBs: make(map[uint]*NoteShapeLinkDB),

		positionDBs: make(map[uint]*PositionDB),

		umlstateDBs: make(map[uint]*UmlStateDB),

		umlscDBs: make(map[uint]*UmlscDB),

		verticeDBs: make(map[uint]*VerticeDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongdoc/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *ClassdiagramDB:
		db.nextIDClassdiagramDB++
		v.ID = db.nextIDClassdiagramDB
		db.classdiagramDBs[v.ID] = v
	case *DiagramPackageDB:
		db.nextIDDiagramPackageDB++
		v.ID = db.nextIDDiagramPackageDB
		db.diagrampackageDBs[v.ID] = v
	case *FieldDB:
		db.nextIDFieldDB++
		v.ID = db.nextIDFieldDB
		db.fieldDBs[v.ID] = v
	case *GongEnumShapeDB:
		db.nextIDGongEnumShapeDB++
		v.ID = db.nextIDGongEnumShapeDB
		db.gongenumshapeDBs[v.ID] = v
	case *GongEnumValueEntryDB:
		db.nextIDGongEnumValueEntryDB++
		v.ID = db.nextIDGongEnumValueEntryDB
		db.gongenumvalueentryDBs[v.ID] = v
	case *GongStructShapeDB:
		db.nextIDGongStructShapeDB++
		v.ID = db.nextIDGongStructShapeDB
		db.gongstructshapeDBs[v.ID] = v
	case *LinkDB:
		db.nextIDLinkDB++
		v.ID = db.nextIDLinkDB
		db.linkDBs[v.ID] = v
	case *NoteShapeDB:
		db.nextIDNoteShapeDB++
		v.ID = db.nextIDNoteShapeDB
		db.noteshapeDBs[v.ID] = v
	case *NoteShapeLinkDB:
		db.nextIDNoteShapeLinkDB++
		v.ID = db.nextIDNoteShapeLinkDB
		db.noteshapelinkDBs[v.ID] = v
	case *PositionDB:
		db.nextIDPositionDB++
		v.ID = db.nextIDPositionDB
		db.positionDBs[v.ID] = v
	case *UmlStateDB:
		db.nextIDUmlStateDB++
		v.ID = db.nextIDUmlStateDB
		db.umlstateDBs[v.ID] = v
	case *UmlscDB:
		db.nextIDUmlscDB++
		v.ID = db.nextIDUmlscDB
		db.umlscDBs[v.ID] = v
	case *VerticeDB:
		db.nextIDVerticeDB++
		v.ID = db.nextIDVerticeDB
		db.verticeDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gongdoc/go, unsupported type in Create")
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
		return nil, errors.New("github.com/fullstack-lang/gongdoc/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ClassdiagramDB:
		delete(db.classdiagramDBs, v.ID)
	case *DiagramPackageDB:
		delete(db.diagrampackageDBs, v.ID)
	case *FieldDB:
		delete(db.fieldDBs, v.ID)
	case *GongEnumShapeDB:
		delete(db.gongenumshapeDBs, v.ID)
	case *GongEnumValueEntryDB:
		delete(db.gongenumvalueentryDBs, v.ID)
	case *GongStructShapeDB:
		delete(db.gongstructshapeDBs, v.ID)
	case *LinkDB:
		delete(db.linkDBs, v.ID)
	case *NoteShapeDB:
		delete(db.noteshapeDBs, v.ID)
	case *NoteShapeLinkDB:
		delete(db.noteshapelinkDBs, v.ID)
	case *PositionDB:
		delete(db.positionDBs, v.ID)
	case *UmlStateDB:
		delete(db.umlstateDBs, v.ID)
	case *UmlscDB:
		delete(db.umlscDBs, v.ID)
	case *VerticeDB:
		delete(db.verticeDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongdoc/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongdoc/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ClassdiagramDB:
		db.classdiagramDBs[v.ID] = v
		return db, nil
	case *DiagramPackageDB:
		db.diagrampackageDBs[v.ID] = v
		return db, nil
	case *FieldDB:
		db.fieldDBs[v.ID] = v
		return db, nil
	case *GongEnumShapeDB:
		db.gongenumshapeDBs[v.ID] = v
		return db, nil
	case *GongEnumValueEntryDB:
		db.gongenumvalueentryDBs[v.ID] = v
		return db, nil
	case *GongStructShapeDB:
		db.gongstructshapeDBs[v.ID] = v
		return db, nil
	case *LinkDB:
		db.linkDBs[v.ID] = v
		return db, nil
	case *NoteShapeDB:
		db.noteshapeDBs[v.ID] = v
		return db, nil
	case *NoteShapeLinkDB:
		db.noteshapelinkDBs[v.ID] = v
		return db, nil
	case *PositionDB:
		db.positionDBs[v.ID] = v
		return db, nil
	case *UmlStateDB:
		db.umlstateDBs[v.ID] = v
		return db, nil
	case *UmlscDB:
		db.umlscDBs[v.ID] = v
		return db, nil
	case *VerticeDB:
		db.verticeDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongdoc/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongdoc/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ClassdiagramDB:
		if existing, ok := db.classdiagramDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongdoc/go, record not found")
		}
	case *DiagramPackageDB:
		if existing, ok := db.diagrampackageDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongdoc/go, record not found")
		}
	case *FieldDB:
		if existing, ok := db.fieldDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongdoc/go, record not found")
		}
	case *GongEnumShapeDB:
		if existing, ok := db.gongenumshapeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongdoc/go, record not found")
		}
	case *GongEnumValueEntryDB:
		if existing, ok := db.gongenumvalueentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongdoc/go, record not found")
		}
	case *GongStructShapeDB:
		if existing, ok := db.gongstructshapeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongdoc/go, record not found")
		}
	case *LinkDB:
		if existing, ok := db.linkDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongdoc/go, record not found")
		}
	case *NoteShapeDB:
		if existing, ok := db.noteshapeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongdoc/go, record not found")
		}
	case *NoteShapeLinkDB:
		if existing, ok := db.noteshapelinkDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongdoc/go, record not found")
		}
	case *PositionDB:
		if existing, ok := db.positionDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongdoc/go, record not found")
		}
	case *UmlStateDB:
		if existing, ok := db.umlstateDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongdoc/go, record not found")
		}
	case *UmlscDB:
		if existing, ok := db.umlscDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongdoc/go, record not found")
		}
	case *VerticeDB:
		if existing, ok := db.verticeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongdoc/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gongdoc/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
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
	case *[]FieldDB:
        *ptr = make([]FieldDB, 0, len(db.fieldDBs))
        for _, v := range db.fieldDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]GongEnumShapeDB:
        *ptr = make([]GongEnumShapeDB, 0, len(db.gongenumshapeDBs))
        for _, v := range db.gongenumshapeDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]GongEnumValueEntryDB:
        *ptr = make([]GongEnumValueEntryDB, 0, len(db.gongenumvalueentryDBs))
        for _, v := range db.gongenumvalueentryDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]GongStructShapeDB:
        *ptr = make([]GongStructShapeDB, 0, len(db.gongstructshapeDBs))
        for _, v := range db.gongstructshapeDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]LinkDB:
        *ptr = make([]LinkDB, 0, len(db.linkDBs))
        for _, v := range db.linkDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]NoteShapeDB:
        *ptr = make([]NoteShapeDB, 0, len(db.noteshapeDBs))
        for _, v := range db.noteshapeDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]NoteShapeLinkDB:
        *ptr = make([]NoteShapeLinkDB, 0, len(db.noteshapelinkDBs))
        for _, v := range db.noteshapelinkDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]PositionDB:
        *ptr = make([]PositionDB, 0, len(db.positionDBs))
        for _, v := range db.positionDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]UmlStateDB:
        *ptr = make([]UmlStateDB, 0, len(db.umlstateDBs))
        for _, v := range db.umlstateDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]UmlscDB:
        *ptr = make([]UmlscDB, 0, len(db.umlscDBs))
        for _, v := range db.umlscDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]VerticeDB:
        *ptr = make([]VerticeDB, 0, len(db.verticeDBs))
        for _, v := range db.verticeDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
    default:
        return nil, errors.New("github.com/fullstack-lang/gongdoc/go, Find: unsupported type")
    }
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gongdoc/go, Do not process when conds is not a single parameter")
	}

	str, ok := conds[0].(string)

	if !ok {
		return nil, errors.New("github.com/fullstack-lang/gongdoc/go, conds[0] is not a string")
	}

	i, err := strconv.ParseUint(str, 10, 32) // Base 10, 32-bit unsigned int
	if err != nil {
		return nil, errors.New("github.com/fullstack-lang/gongdoc/go, conds[0] is not a string number")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *ClassdiagramDB:
		tmp, ok := db.classdiagramDBs[uint(i)]

		classdiagramDB, _ := instanceDB.(*ClassdiagramDB)
		*classdiagramDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *DiagramPackageDB:
		tmp, ok := db.diagrampackageDBs[uint(i)]

		diagrampackageDB, _ := instanceDB.(*DiagramPackageDB)
		*diagrampackageDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *FieldDB:
		tmp, ok := db.fieldDBs[uint(i)]

		fieldDB, _ := instanceDB.(*FieldDB)
		*fieldDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *GongEnumShapeDB:
		tmp, ok := db.gongenumshapeDBs[uint(i)]

		gongenumshapeDB, _ := instanceDB.(*GongEnumShapeDB)
		*gongenumshapeDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *GongEnumValueEntryDB:
		tmp, ok := db.gongenumvalueentryDBs[uint(i)]

		gongenumvalueentryDB, _ := instanceDB.(*GongEnumValueEntryDB)
		*gongenumvalueentryDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *GongStructShapeDB:
		tmp, ok := db.gongstructshapeDBs[uint(i)]

		gongstructshapeDB, _ := instanceDB.(*GongStructShapeDB)
		*gongstructshapeDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *LinkDB:
		tmp, ok := db.linkDBs[uint(i)]

		linkDB, _ := instanceDB.(*LinkDB)
		*linkDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *NoteShapeDB:
		tmp, ok := db.noteshapeDBs[uint(i)]

		noteshapeDB, _ := instanceDB.(*NoteShapeDB)
		*noteshapeDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *NoteShapeLinkDB:
		tmp, ok := db.noteshapelinkDBs[uint(i)]

		noteshapelinkDB, _ := instanceDB.(*NoteShapeLinkDB)
		*noteshapelinkDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *PositionDB:
		tmp, ok := db.positionDBs[uint(i)]

		positionDB, _ := instanceDB.(*PositionDB)
		*positionDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *UmlStateDB:
		tmp, ok := db.umlstateDBs[uint(i)]

		umlstateDB, _ := instanceDB.(*UmlStateDB)
		*umlstateDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *UmlscDB:
		tmp, ok := db.umlscDBs[uint(i)]

		umlscDB, _ := instanceDB.(*UmlscDB)
		*umlscDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *VerticeDB:
		tmp, ok := db.verticeDBs[uint(i)]

		verticeDB, _ := instanceDB.(*VerticeDB)
		*verticeDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gongdoc/go, Unkown type")
	}
	
	return db, nil
}

