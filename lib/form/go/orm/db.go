// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gong/lib/form/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	checkboxDBs map[uint]*CheckBoxDB

	nextIDCheckBoxDB uint

	formdivDBs map[uint]*FormDivDB

	nextIDFormDivDB uint

	formeditassocbuttonDBs map[uint]*FormEditAssocButtonDB

	nextIDFormEditAssocButtonDB uint

	formfieldDBs map[uint]*FormFieldDB

	nextIDFormFieldDB uint

	formfielddateDBs map[uint]*FormFieldDateDB

	nextIDFormFieldDateDB uint

	formfielddatetimeDBs map[uint]*FormFieldDateTimeDB

	nextIDFormFieldDateTimeDB uint

	formfieldfloat64DBs map[uint]*FormFieldFloat64DB

	nextIDFormFieldFloat64DB uint

	formfieldintDBs map[uint]*FormFieldIntDB

	nextIDFormFieldIntDB uint

	formfieldselectDBs map[uint]*FormFieldSelectDB

	nextIDFormFieldSelectDB uint

	formfieldstringDBs map[uint]*FormFieldStringDB

	nextIDFormFieldStringDB uint

	formfieldtimeDBs map[uint]*FormFieldTimeDB

	nextIDFormFieldTimeDB uint

	formgroupDBs map[uint]*FormGroupDB

	nextIDFormGroupDB uint

	formsortassocbuttonDBs map[uint]*FormSortAssocButtonDB

	nextIDFormSortAssocButtonDB uint

	optionDBs map[uint]*OptionDB

	nextIDOptionDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		checkboxDBs: make(map[uint]*CheckBoxDB),

		formdivDBs: make(map[uint]*FormDivDB),

		formeditassocbuttonDBs: make(map[uint]*FormEditAssocButtonDB),

		formfieldDBs: make(map[uint]*FormFieldDB),

		formfielddateDBs: make(map[uint]*FormFieldDateDB),

		formfielddatetimeDBs: make(map[uint]*FormFieldDateTimeDB),

		formfieldfloat64DBs: make(map[uint]*FormFieldFloat64DB),

		formfieldintDBs: make(map[uint]*FormFieldIntDB),

		formfieldselectDBs: make(map[uint]*FormFieldSelectDB),

		formfieldstringDBs: make(map[uint]*FormFieldStringDB),

		formfieldtimeDBs: make(map[uint]*FormFieldTimeDB),

		formgroupDBs: make(map[uint]*FormGroupDB),

		formsortassocbuttonDBs: make(map[uint]*FormSortAssocButtonDB),

		optionDBs: make(map[uint]*OptionDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/form/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *CheckBoxDB:
		db.nextIDCheckBoxDB++
		v.ID = db.nextIDCheckBoxDB
		db.checkboxDBs[v.ID] = v
	case *FormDivDB:
		db.nextIDFormDivDB++
		v.ID = db.nextIDFormDivDB
		db.formdivDBs[v.ID] = v
	case *FormEditAssocButtonDB:
		db.nextIDFormEditAssocButtonDB++
		v.ID = db.nextIDFormEditAssocButtonDB
		db.formeditassocbuttonDBs[v.ID] = v
	case *FormFieldDB:
		db.nextIDFormFieldDB++
		v.ID = db.nextIDFormFieldDB
		db.formfieldDBs[v.ID] = v
	case *FormFieldDateDB:
		db.nextIDFormFieldDateDB++
		v.ID = db.nextIDFormFieldDateDB
		db.formfielddateDBs[v.ID] = v
	case *FormFieldDateTimeDB:
		db.nextIDFormFieldDateTimeDB++
		v.ID = db.nextIDFormFieldDateTimeDB
		db.formfielddatetimeDBs[v.ID] = v
	case *FormFieldFloat64DB:
		db.nextIDFormFieldFloat64DB++
		v.ID = db.nextIDFormFieldFloat64DB
		db.formfieldfloat64DBs[v.ID] = v
	case *FormFieldIntDB:
		db.nextIDFormFieldIntDB++
		v.ID = db.nextIDFormFieldIntDB
		db.formfieldintDBs[v.ID] = v
	case *FormFieldSelectDB:
		db.nextIDFormFieldSelectDB++
		v.ID = db.nextIDFormFieldSelectDB
		db.formfieldselectDBs[v.ID] = v
	case *FormFieldStringDB:
		db.nextIDFormFieldStringDB++
		v.ID = db.nextIDFormFieldStringDB
		db.formfieldstringDBs[v.ID] = v
	case *FormFieldTimeDB:
		db.nextIDFormFieldTimeDB++
		v.ID = db.nextIDFormFieldTimeDB
		db.formfieldtimeDBs[v.ID] = v
	case *FormGroupDB:
		db.nextIDFormGroupDB++
		v.ID = db.nextIDFormGroupDB
		db.formgroupDBs[v.ID] = v
	case *FormSortAssocButtonDB:
		db.nextIDFormSortAssocButtonDB++
		v.ID = db.nextIDFormSortAssocButtonDB
		db.formsortassocbuttonDBs[v.ID] = v
	case *OptionDB:
		db.nextIDOptionDB++
		v.ID = db.nextIDOptionDB
		db.optionDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/form/go, unsupported type in Create")
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
		return nil, errors.New("github.com/fullstack-lang/gong/lib/form/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *CheckBoxDB:
		delete(db.checkboxDBs, v.ID)
	case *FormDivDB:
		delete(db.formdivDBs, v.ID)
	case *FormEditAssocButtonDB:
		delete(db.formeditassocbuttonDBs, v.ID)
	case *FormFieldDB:
		delete(db.formfieldDBs, v.ID)
	case *FormFieldDateDB:
		delete(db.formfielddateDBs, v.ID)
	case *FormFieldDateTimeDB:
		delete(db.formfielddatetimeDBs, v.ID)
	case *FormFieldFloat64DB:
		delete(db.formfieldfloat64DBs, v.ID)
	case *FormFieldIntDB:
		delete(db.formfieldintDBs, v.ID)
	case *FormFieldSelectDB:
		delete(db.formfieldselectDBs, v.ID)
	case *FormFieldStringDB:
		delete(db.formfieldstringDBs, v.ID)
	case *FormFieldTimeDB:
		delete(db.formfieldtimeDBs, v.ID)
	case *FormGroupDB:
		delete(db.formgroupDBs, v.ID)
	case *FormSortAssocButtonDB:
		delete(db.formsortassocbuttonDBs, v.ID)
	case *OptionDB:
		delete(db.optionDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/form/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/form/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *CheckBoxDB:
		db.checkboxDBs[v.ID] = v
		return db, nil
	case *FormDivDB:
		db.formdivDBs[v.ID] = v
		return db, nil
	case *FormEditAssocButtonDB:
		db.formeditassocbuttonDBs[v.ID] = v
		return db, nil
	case *FormFieldDB:
		db.formfieldDBs[v.ID] = v
		return db, nil
	case *FormFieldDateDB:
		db.formfielddateDBs[v.ID] = v
		return db, nil
	case *FormFieldDateTimeDB:
		db.formfielddatetimeDBs[v.ID] = v
		return db, nil
	case *FormFieldFloat64DB:
		db.formfieldfloat64DBs[v.ID] = v
		return db, nil
	case *FormFieldIntDB:
		db.formfieldintDBs[v.ID] = v
		return db, nil
	case *FormFieldSelectDB:
		db.formfieldselectDBs[v.ID] = v
		return db, nil
	case *FormFieldStringDB:
		db.formfieldstringDBs[v.ID] = v
		return db, nil
	case *FormFieldTimeDB:
		db.formfieldtimeDBs[v.ID] = v
		return db, nil
	case *FormGroupDB:
		db.formgroupDBs[v.ID] = v
		return db, nil
	case *FormSortAssocButtonDB:
		db.formsortassocbuttonDBs[v.ID] = v
		return db, nil
	case *OptionDB:
		db.optionDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/form/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/form/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *CheckBoxDB:
		if existing, ok := db.checkboxDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db CheckBox github.com/fullstack-lang/gong/lib/form/go, record not found")
		}
	case *FormDivDB:
		if existing, ok := db.formdivDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormDiv github.com/fullstack-lang/gong/lib/form/go, record not found")
		}
	case *FormEditAssocButtonDB:
		if existing, ok := db.formeditassocbuttonDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormEditAssocButton github.com/fullstack-lang/gong/lib/form/go, record not found")
		}
	case *FormFieldDB:
		if existing, ok := db.formfieldDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormField github.com/fullstack-lang/gong/lib/form/go, record not found")
		}
	case *FormFieldDateDB:
		if existing, ok := db.formfielddateDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormFieldDate github.com/fullstack-lang/gong/lib/form/go, record not found")
		}
	case *FormFieldDateTimeDB:
		if existing, ok := db.formfielddatetimeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormFieldDateTime github.com/fullstack-lang/gong/lib/form/go, record not found")
		}
	case *FormFieldFloat64DB:
		if existing, ok := db.formfieldfloat64DBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormFieldFloat64 github.com/fullstack-lang/gong/lib/form/go, record not found")
		}
	case *FormFieldIntDB:
		if existing, ok := db.formfieldintDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormFieldInt github.com/fullstack-lang/gong/lib/form/go, record not found")
		}
	case *FormFieldSelectDB:
		if existing, ok := db.formfieldselectDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormFieldSelect github.com/fullstack-lang/gong/lib/form/go, record not found")
		}
	case *FormFieldStringDB:
		if existing, ok := db.formfieldstringDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormFieldString github.com/fullstack-lang/gong/lib/form/go, record not found")
		}
	case *FormFieldTimeDB:
		if existing, ok := db.formfieldtimeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormFieldTime github.com/fullstack-lang/gong/lib/form/go, record not found")
		}
	case *FormGroupDB:
		if existing, ok := db.formgroupDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormGroup github.com/fullstack-lang/gong/lib/form/go, record not found")
		}
	case *FormSortAssocButtonDB:
		if existing, ok := db.formsortassocbuttonDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormSortAssocButton github.com/fullstack-lang/gong/lib/form/go, record not found")
		}
	case *OptionDB:
		if existing, ok := db.optionDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Option github.com/fullstack-lang/gong/lib/form/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/form/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]CheckBoxDB:
		*ptr = make([]CheckBoxDB, 0, len(db.checkboxDBs))
		for _, v := range db.checkboxDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]FormDivDB:
		*ptr = make([]FormDivDB, 0, len(db.formdivDBs))
		for _, v := range db.formdivDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]FormEditAssocButtonDB:
		*ptr = make([]FormEditAssocButtonDB, 0, len(db.formeditassocbuttonDBs))
		for _, v := range db.formeditassocbuttonDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]FormFieldDB:
		*ptr = make([]FormFieldDB, 0, len(db.formfieldDBs))
		for _, v := range db.formfieldDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]FormFieldDateDB:
		*ptr = make([]FormFieldDateDB, 0, len(db.formfielddateDBs))
		for _, v := range db.formfielddateDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]FormFieldDateTimeDB:
		*ptr = make([]FormFieldDateTimeDB, 0, len(db.formfielddatetimeDBs))
		for _, v := range db.formfielddatetimeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]FormFieldFloat64DB:
		*ptr = make([]FormFieldFloat64DB, 0, len(db.formfieldfloat64DBs))
		for _, v := range db.formfieldfloat64DBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]FormFieldIntDB:
		*ptr = make([]FormFieldIntDB, 0, len(db.formfieldintDBs))
		for _, v := range db.formfieldintDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]FormFieldSelectDB:
		*ptr = make([]FormFieldSelectDB, 0, len(db.formfieldselectDBs))
		for _, v := range db.formfieldselectDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]FormFieldStringDB:
		*ptr = make([]FormFieldStringDB, 0, len(db.formfieldstringDBs))
		for _, v := range db.formfieldstringDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]FormFieldTimeDB:
		*ptr = make([]FormFieldTimeDB, 0, len(db.formfieldtimeDBs))
		for _, v := range db.formfieldtimeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]FormGroupDB:
		*ptr = make([]FormGroupDB, 0, len(db.formgroupDBs))
		for _, v := range db.formgroupDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]FormSortAssocButtonDB:
		*ptr = make([]FormSortAssocButtonDB, 0, len(db.formsortassocbuttonDBs))
		for _, v := range db.formsortassocbuttonDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]OptionDB:
		*ptr = make([]OptionDB, 0, len(db.optionDBs))
		for _, v := range db.optionDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/form/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/form/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gong/lib/form/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/form/go, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *CheckBoxDB:
		tmp, ok := db.checkboxDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First CheckBox Unkown entry %d", i))
		}

		checkboxDB, _ := instanceDB.(*CheckBoxDB)
		*checkboxDB = *tmp

	case *FormDivDB:
		tmp, ok := db.formdivDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First FormDiv Unkown entry %d", i))
		}

		formdivDB, _ := instanceDB.(*FormDivDB)
		*formdivDB = *tmp

	case *FormEditAssocButtonDB:
		tmp, ok := db.formeditassocbuttonDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First FormEditAssocButton Unkown entry %d", i))
		}

		formeditassocbuttonDB, _ := instanceDB.(*FormEditAssocButtonDB)
		*formeditassocbuttonDB = *tmp

	case *FormFieldDB:
		tmp, ok := db.formfieldDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First FormField Unkown entry %d", i))
		}

		formfieldDB, _ := instanceDB.(*FormFieldDB)
		*formfieldDB = *tmp

	case *FormFieldDateDB:
		tmp, ok := db.formfielddateDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First FormFieldDate Unkown entry %d", i))
		}

		formfielddateDB, _ := instanceDB.(*FormFieldDateDB)
		*formfielddateDB = *tmp

	case *FormFieldDateTimeDB:
		tmp, ok := db.formfielddatetimeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First FormFieldDateTime Unkown entry %d", i))
		}

		formfielddatetimeDB, _ := instanceDB.(*FormFieldDateTimeDB)
		*formfielddatetimeDB = *tmp

	case *FormFieldFloat64DB:
		tmp, ok := db.formfieldfloat64DBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First FormFieldFloat64 Unkown entry %d", i))
		}

		formfieldfloat64DB, _ := instanceDB.(*FormFieldFloat64DB)
		*formfieldfloat64DB = *tmp

	case *FormFieldIntDB:
		tmp, ok := db.formfieldintDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First FormFieldInt Unkown entry %d", i))
		}

		formfieldintDB, _ := instanceDB.(*FormFieldIntDB)
		*formfieldintDB = *tmp

	case *FormFieldSelectDB:
		tmp, ok := db.formfieldselectDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First FormFieldSelect Unkown entry %d", i))
		}

		formfieldselectDB, _ := instanceDB.(*FormFieldSelectDB)
		*formfieldselectDB = *tmp

	case *FormFieldStringDB:
		tmp, ok := db.formfieldstringDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First FormFieldString Unkown entry %d", i))
		}

		formfieldstringDB, _ := instanceDB.(*FormFieldStringDB)
		*formfieldstringDB = *tmp

	case *FormFieldTimeDB:
		tmp, ok := db.formfieldtimeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First FormFieldTime Unkown entry %d", i))
		}

		formfieldtimeDB, _ := instanceDB.(*FormFieldTimeDB)
		*formfieldtimeDB = *tmp

	case *FormGroupDB:
		tmp, ok := db.formgroupDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First FormGroup Unkown entry %d", i))
		}

		formgroupDB, _ := instanceDB.(*FormGroupDB)
		*formgroupDB = *tmp

	case *FormSortAssocButtonDB:
		tmp, ok := db.formsortassocbuttonDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First FormSortAssocButton Unkown entry %d", i))
		}

		formsortassocbuttonDB, _ := instanceDB.(*FormSortAssocButtonDB)
		*formsortassocbuttonDB = *tmp

	case *OptionDB:
		tmp, ok := db.optionDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Option Unkown entry %d", i))
		}

		optionDB, _ := instanceDB.(*OptionDB)
		*optionDB = *tmp

	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/form/go, Unkown type")
	}
	
	return db, nil
}

