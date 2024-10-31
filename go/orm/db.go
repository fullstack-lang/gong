// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gong/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	gongbasicfieldDBs map[uint]*GongBasicFieldDB

	nextIDGongBasicFieldDB uint

	gongenumDBs map[uint]*GongEnumDB

	nextIDGongEnumDB uint

	gongenumvalueDBs map[uint]*GongEnumValueDB

	nextIDGongEnumValueDB uint

	gonglinkDBs map[uint]*GongLinkDB

	nextIDGongLinkDB uint

	gongnoteDBs map[uint]*GongNoteDB

	nextIDGongNoteDB uint

	gongstructDBs map[uint]*GongStructDB

	nextIDGongStructDB uint

	gongtimefieldDBs map[uint]*GongTimeFieldDB

	nextIDGongTimeFieldDB uint

	metaDBs map[uint]*MetaDB

	nextIDMetaDB uint

	metareferenceDBs map[uint]*MetaReferenceDB

	nextIDMetaReferenceDB uint

	modelpkgDBs map[uint]*ModelPkgDB

	nextIDModelPkgDB uint

	pointertogongstructfieldDBs map[uint]*PointerToGongStructFieldDB

	nextIDPointerToGongStructFieldDB uint

	sliceofpointertogongstructfieldDBs map[uint]*SliceOfPointerToGongStructFieldDB

	nextIDSliceOfPointerToGongStructFieldDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		gongbasicfieldDBs: make(map[uint]*GongBasicFieldDB),

		gongenumDBs: make(map[uint]*GongEnumDB),

		gongenumvalueDBs: make(map[uint]*GongEnumValueDB),

		gonglinkDBs: make(map[uint]*GongLinkDB),

		gongnoteDBs: make(map[uint]*GongNoteDB),

		gongstructDBs: make(map[uint]*GongStructDB),

		gongtimefieldDBs: make(map[uint]*GongTimeFieldDB),

		metaDBs: make(map[uint]*MetaDB),

		metareferenceDBs: make(map[uint]*MetaReferenceDB),

		modelpkgDBs: make(map[uint]*ModelPkgDB),

		pointertogongstructfieldDBs: make(map[uint]*PointerToGongStructFieldDB),

		sliceofpointertogongstructfieldDBs: make(map[uint]*SliceOfPointerToGongStructFieldDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *GongBasicFieldDB:
		db.nextIDGongBasicFieldDB++
		v.ID = db.nextIDGongBasicFieldDB
		db.gongbasicfieldDBs[v.ID] = v
	case *GongEnumDB:
		db.nextIDGongEnumDB++
		v.ID = db.nextIDGongEnumDB
		db.gongenumDBs[v.ID] = v
	case *GongEnumValueDB:
		db.nextIDGongEnumValueDB++
		v.ID = db.nextIDGongEnumValueDB
		db.gongenumvalueDBs[v.ID] = v
	case *GongLinkDB:
		db.nextIDGongLinkDB++
		v.ID = db.nextIDGongLinkDB
		db.gonglinkDBs[v.ID] = v
	case *GongNoteDB:
		db.nextIDGongNoteDB++
		v.ID = db.nextIDGongNoteDB
		db.gongnoteDBs[v.ID] = v
	case *GongStructDB:
		db.nextIDGongStructDB++
		v.ID = db.nextIDGongStructDB
		db.gongstructDBs[v.ID] = v
	case *GongTimeFieldDB:
		db.nextIDGongTimeFieldDB++
		v.ID = db.nextIDGongTimeFieldDB
		db.gongtimefieldDBs[v.ID] = v
	case *MetaDB:
		db.nextIDMetaDB++
		v.ID = db.nextIDMetaDB
		db.metaDBs[v.ID] = v
	case *MetaReferenceDB:
		db.nextIDMetaReferenceDB++
		v.ID = db.nextIDMetaReferenceDB
		db.metareferenceDBs[v.ID] = v
	case *ModelPkgDB:
		db.nextIDModelPkgDB++
		v.ID = db.nextIDModelPkgDB
		db.modelpkgDBs[v.ID] = v
	case *PointerToGongStructFieldDB:
		db.nextIDPointerToGongStructFieldDB++
		v.ID = db.nextIDPointerToGongStructFieldDB
		db.pointertogongstructfieldDBs[v.ID] = v
	case *SliceOfPointerToGongStructFieldDB:
		db.nextIDSliceOfPointerToGongStructFieldDB++
		v.ID = db.nextIDSliceOfPointerToGongStructFieldDB
		db.sliceofpointertogongstructfieldDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/go, unsupported type in Create")
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
		return nil, errors.New("github.com/fullstack-lang/gong/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *GongBasicFieldDB:
		delete(db.gongbasicfieldDBs, v.ID)
	case *GongEnumDB:
		delete(db.gongenumDBs, v.ID)
	case *GongEnumValueDB:
		delete(db.gongenumvalueDBs, v.ID)
	case *GongLinkDB:
		delete(db.gonglinkDBs, v.ID)
	case *GongNoteDB:
		delete(db.gongnoteDBs, v.ID)
	case *GongStructDB:
		delete(db.gongstructDBs, v.ID)
	case *GongTimeFieldDB:
		delete(db.gongtimefieldDBs, v.ID)
	case *MetaDB:
		delete(db.metaDBs, v.ID)
	case *MetaReferenceDB:
		delete(db.metareferenceDBs, v.ID)
	case *ModelPkgDB:
		delete(db.modelpkgDBs, v.ID)
	case *PointerToGongStructFieldDB:
		delete(db.pointertogongstructfieldDBs, v.ID)
	case *SliceOfPointerToGongStructFieldDB:
		delete(db.sliceofpointertogongstructfieldDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *GongBasicFieldDB:
		db.gongbasicfieldDBs[v.ID] = v
		return db, nil
	case *GongEnumDB:
		db.gongenumDBs[v.ID] = v
		return db, nil
	case *GongEnumValueDB:
		db.gongenumvalueDBs[v.ID] = v
		return db, nil
	case *GongLinkDB:
		db.gonglinkDBs[v.ID] = v
		return db, nil
	case *GongNoteDB:
		db.gongnoteDBs[v.ID] = v
		return db, nil
	case *GongStructDB:
		db.gongstructDBs[v.ID] = v
		return db, nil
	case *GongTimeFieldDB:
		db.gongtimefieldDBs[v.ID] = v
		return db, nil
	case *MetaDB:
		db.metaDBs[v.ID] = v
		return db, nil
	case *MetaReferenceDB:
		db.metareferenceDBs[v.ID] = v
		return db, nil
	case *ModelPkgDB:
		db.modelpkgDBs[v.ID] = v
		return db, nil
	case *PointerToGongStructFieldDB:
		db.pointertogongstructfieldDBs[v.ID] = v
		return db, nil
	case *SliceOfPointerToGongStructFieldDB:
		db.sliceofpointertogongstructfieldDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *GongBasicFieldDB:
		if existing, ok := db.gongbasicfieldDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db GongBasicField github.com/fullstack-lang/gong/go, record not found")
		}
	case *GongEnumDB:
		if existing, ok := db.gongenumDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db GongEnum github.com/fullstack-lang/gong/go, record not found")
		}
	case *GongEnumValueDB:
		if existing, ok := db.gongenumvalueDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db GongEnumValue github.com/fullstack-lang/gong/go, record not found")
		}
	case *GongLinkDB:
		if existing, ok := db.gonglinkDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db GongLink github.com/fullstack-lang/gong/go, record not found")
		}
	case *GongNoteDB:
		if existing, ok := db.gongnoteDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db GongNote github.com/fullstack-lang/gong/go, record not found")
		}
	case *GongStructDB:
		if existing, ok := db.gongstructDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db GongStruct github.com/fullstack-lang/gong/go, record not found")
		}
	case *GongTimeFieldDB:
		if existing, ok := db.gongtimefieldDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db GongTimeField github.com/fullstack-lang/gong/go, record not found")
		}
	case *MetaDB:
		if existing, ok := db.metaDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Meta github.com/fullstack-lang/gong/go, record not found")
		}
	case *MetaReferenceDB:
		if existing, ok := db.metareferenceDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db MetaReference github.com/fullstack-lang/gong/go, record not found")
		}
	case *ModelPkgDB:
		if existing, ok := db.modelpkgDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ModelPkg github.com/fullstack-lang/gong/go, record not found")
		}
	case *PointerToGongStructFieldDB:
		if existing, ok := db.pointertogongstructfieldDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db PointerToGongStructField github.com/fullstack-lang/gong/go, record not found")
		}
	case *SliceOfPointerToGongStructFieldDB:
		if existing, ok := db.sliceofpointertogongstructfieldDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SliceOfPointerToGongStructField github.com/fullstack-lang/gong/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]GongBasicFieldDB:
		*ptr = make([]GongBasicFieldDB, 0, len(db.gongbasicfieldDBs))
		for _, v := range db.gongbasicfieldDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]GongEnumDB:
		*ptr = make([]GongEnumDB, 0, len(db.gongenumDBs))
		for _, v := range db.gongenumDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]GongEnumValueDB:
		*ptr = make([]GongEnumValueDB, 0, len(db.gongenumvalueDBs))
		for _, v := range db.gongenumvalueDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]GongLinkDB:
		*ptr = make([]GongLinkDB, 0, len(db.gonglinkDBs))
		for _, v := range db.gonglinkDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]GongNoteDB:
		*ptr = make([]GongNoteDB, 0, len(db.gongnoteDBs))
		for _, v := range db.gongnoteDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]GongStructDB:
		*ptr = make([]GongStructDB, 0, len(db.gongstructDBs))
		for _, v := range db.gongstructDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]GongTimeFieldDB:
		*ptr = make([]GongTimeFieldDB, 0, len(db.gongtimefieldDBs))
		for _, v := range db.gongtimefieldDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]MetaDB:
		*ptr = make([]MetaDB, 0, len(db.metaDBs))
		for _, v := range db.metaDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]MetaReferenceDB:
		*ptr = make([]MetaReferenceDB, 0, len(db.metareferenceDBs))
		for _, v := range db.metareferenceDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ModelPkgDB:
		*ptr = make([]ModelPkgDB, 0, len(db.modelpkgDBs))
		for _, v := range db.modelpkgDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]PointerToGongStructFieldDB:
		*ptr = make([]PointerToGongStructFieldDB, 0, len(db.pointertogongstructfieldDBs))
		for _, v := range db.pointertogongstructfieldDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SliceOfPointerToGongStructFieldDB:
		*ptr = make([]SliceOfPointerToGongStructFieldDB, 0, len(db.sliceofpointertogongstructfieldDBs))
		for _, v := range db.sliceofpointertogongstructfieldDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gong/go, Do not process when conds is not a single parameter")
	}

	str, ok := conds[0].(string)

	if !ok {
		return nil, errors.New("github.com/fullstack-lang/gong/go, conds[0] is not a string")
	}

	i, err := strconv.ParseUint(str, 10, 32) // Base 10, 32-bit unsigned int
	if err != nil {
		return nil, errors.New("github.com/fullstack-lang/gong/go, conds[0] is not a string number")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *GongBasicFieldDB:
		tmp, ok := db.gongbasicfieldDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First GongBasicField Unkown entry %d", i))
		}

		gongbasicfieldDB, _ := instanceDB.(*GongBasicFieldDB)
		*gongbasicfieldDB = *tmp
		
	case *GongEnumDB:
		tmp, ok := db.gongenumDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First GongEnum Unkown entry %d", i))
		}

		gongenumDB, _ := instanceDB.(*GongEnumDB)
		*gongenumDB = *tmp
		
	case *GongEnumValueDB:
		tmp, ok := db.gongenumvalueDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First GongEnumValue Unkown entry %d", i))
		}

		gongenumvalueDB, _ := instanceDB.(*GongEnumValueDB)
		*gongenumvalueDB = *tmp
		
	case *GongLinkDB:
		tmp, ok := db.gonglinkDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First GongLink Unkown entry %d", i))
		}

		gonglinkDB, _ := instanceDB.(*GongLinkDB)
		*gonglinkDB = *tmp
		
	case *GongNoteDB:
		tmp, ok := db.gongnoteDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First GongNote Unkown entry %d", i))
		}

		gongnoteDB, _ := instanceDB.(*GongNoteDB)
		*gongnoteDB = *tmp
		
	case *GongStructDB:
		tmp, ok := db.gongstructDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First GongStruct Unkown entry %d", i))
		}

		gongstructDB, _ := instanceDB.(*GongStructDB)
		*gongstructDB = *tmp
		
	case *GongTimeFieldDB:
		tmp, ok := db.gongtimefieldDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First GongTimeField Unkown entry %d", i))
		}

		gongtimefieldDB, _ := instanceDB.(*GongTimeFieldDB)
		*gongtimefieldDB = *tmp
		
	case *MetaDB:
		tmp, ok := db.metaDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Meta Unkown entry %d", i))
		}

		metaDB, _ := instanceDB.(*MetaDB)
		*metaDB = *tmp
		
	case *MetaReferenceDB:
		tmp, ok := db.metareferenceDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First MetaReference Unkown entry %d", i))
		}

		metareferenceDB, _ := instanceDB.(*MetaReferenceDB)
		*metareferenceDB = *tmp
		
	case *ModelPkgDB:
		tmp, ok := db.modelpkgDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ModelPkg Unkown entry %d", i))
		}

		modelpkgDB, _ := instanceDB.(*ModelPkgDB)
		*modelpkgDB = *tmp
		
	case *PointerToGongStructFieldDB:
		tmp, ok := db.pointertogongstructfieldDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First PointerToGongStructField Unkown entry %d", i))
		}

		pointertogongstructfieldDB, _ := instanceDB.(*PointerToGongStructFieldDB)
		*pointertogongstructfieldDB = *tmp
		
	case *SliceOfPointerToGongStructFieldDB:
		tmp, ok := db.sliceofpointertogongstructfieldDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SliceOfPointerToGongStructField Unkown entry %d", i))
		}

		sliceofpointertogongstructfieldDB, _ := instanceDB.(*SliceOfPointerToGongStructFieldDB)
		*sliceofpointertogongstructfieldDB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/go, Unkown type")
	}
	
	return db, nil
}

