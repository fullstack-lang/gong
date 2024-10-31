// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gongtable/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

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

	checkboxDBs map[uint]*CheckBoxDB

	nextIDCheckBoxDB uint

	displayedcolumnDBs map[uint]*DisplayedColumnDB

	nextIDDisplayedColumnDB uint

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

	rowDBs map[uint]*RowDB

	nextIDRowDB uint

	tableDBs map[uint]*TableDB

	nextIDTableDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		cellDBs: make(map[uint]*CellDB),

		cellbooleanDBs: make(map[uint]*CellBooleanDB),

		cellfloat64DBs: make(map[uint]*CellFloat64DB),

		celliconDBs: make(map[uint]*CellIconDB),

		cellintDBs: make(map[uint]*CellIntDB),

		cellstringDBs: make(map[uint]*CellStringDB),

		checkboxDBs: make(map[uint]*CheckBoxDB),

		displayedcolumnDBs: make(map[uint]*DisplayedColumnDB),

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

		rowDBs: make(map[uint]*RowDB),

		tableDBs: make(map[uint]*TableDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongtable/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
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
	case *CheckBoxDB:
		db.nextIDCheckBoxDB++
		v.ID = db.nextIDCheckBoxDB
		db.checkboxDBs[v.ID] = v
	case *DisplayedColumnDB:
		db.nextIDDisplayedColumnDB++
		v.ID = db.nextIDDisplayedColumnDB
		db.displayedcolumnDBs[v.ID] = v
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
	case *RowDB:
		db.nextIDRowDB++
		v.ID = db.nextIDRowDB
		db.rowDBs[v.ID] = v
	case *TableDB:
		db.nextIDTableDB++
		v.ID = db.nextIDTableDB
		db.tableDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gongtable/go, unsupported type in Create")
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
		return nil, errors.New("github.com/fullstack-lang/gongtable/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
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
	case *CheckBoxDB:
		delete(db.checkboxDBs, v.ID)
	case *DisplayedColumnDB:
		delete(db.displayedcolumnDBs, v.ID)
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
	case *RowDB:
		delete(db.rowDBs, v.ID)
	case *TableDB:
		delete(db.tableDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongtable/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongtable/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
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
	case *CheckBoxDB:
		db.checkboxDBs[v.ID] = v
		return db, nil
	case *DisplayedColumnDB:
		db.displayedcolumnDBs[v.ID] = v
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
	case *RowDB:
		db.rowDBs[v.ID] = v
		return db, nil
	case *TableDB:
		db.tableDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongtable/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongtable/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *CellDB:
		if existing, ok := db.cellDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Cell github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *CellBooleanDB:
		if existing, ok := db.cellbooleanDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db CellBoolean github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *CellFloat64DB:
		if existing, ok := db.cellfloat64DBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db CellFloat64 github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *CellIconDB:
		if existing, ok := db.celliconDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db CellIcon github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *CellIntDB:
		if existing, ok := db.cellintDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db CellInt github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *CellStringDB:
		if existing, ok := db.cellstringDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db CellString github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *CheckBoxDB:
		if existing, ok := db.checkboxDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db CheckBox github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *DisplayedColumnDB:
		if existing, ok := db.displayedcolumnDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DisplayedColumn github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *FormDivDB:
		if existing, ok := db.formdivDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormDiv github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *FormEditAssocButtonDB:
		if existing, ok := db.formeditassocbuttonDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormEditAssocButton github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *FormFieldDB:
		if existing, ok := db.formfieldDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormField github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *FormFieldDateDB:
		if existing, ok := db.formfielddateDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormFieldDate github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *FormFieldDateTimeDB:
		if existing, ok := db.formfielddatetimeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormFieldDateTime github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *FormFieldFloat64DB:
		if existing, ok := db.formfieldfloat64DBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormFieldFloat64 github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *FormFieldIntDB:
		if existing, ok := db.formfieldintDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormFieldInt github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *FormFieldSelectDB:
		if existing, ok := db.formfieldselectDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormFieldSelect github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *FormFieldStringDB:
		if existing, ok := db.formfieldstringDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormFieldString github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *FormFieldTimeDB:
		if existing, ok := db.formfieldtimeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormFieldTime github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *FormGroupDB:
		if existing, ok := db.formgroupDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormGroup github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *FormSortAssocButtonDB:
		if existing, ok := db.formsortassocbuttonDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FormSortAssocButton github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *OptionDB:
		if existing, ok := db.optionDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Option github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *RowDB:
		if existing, ok := db.rowDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Row github.com/fullstack-lang/gongtable/go, record not found")
		}
	case *TableDB:
		if existing, ok := db.tableDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Table github.com/fullstack-lang/gongtable/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gongtable/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
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
	case *[]CheckBoxDB:
		*ptr = make([]CheckBoxDB, 0, len(db.checkboxDBs))
		for _, v := range db.checkboxDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DisplayedColumnDB:
		*ptr = make([]DisplayedColumnDB, 0, len(db.displayedcolumnDBs))
		for _, v := range db.displayedcolumnDBs {
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
	case *[]RowDB:
		*ptr = make([]RowDB, 0, len(db.rowDBs))
		for _, v := range db.rowDBs {
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
		return nil, errors.New("github.com/fullstack-lang/gongtable/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gongtable/go, Do not process when conds is not a single parameter")
	}

	str, ok := conds[0].(string)

	if !ok {
		return nil, errors.New("github.com/fullstack-lang/gongtable/go, conds[0] is not a string")
	}

	i, err := strconv.ParseUint(str, 10, 32) // Base 10, 32-bit unsigned int
	if err != nil {
		return nil, errors.New("github.com/fullstack-lang/gongtable/go, conds[0] is not a string number")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
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
		
	case *CheckBoxDB:
		tmp, ok := db.checkboxDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First CheckBox Unkown entry %d", i))
		}

		checkboxDB, _ := instanceDB.(*CheckBoxDB)
		*checkboxDB = *tmp
		
	case *DisplayedColumnDB:
		tmp, ok := db.displayedcolumnDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DisplayedColumn Unkown entry %d", i))
		}

		displayedcolumnDB, _ := instanceDB.(*DisplayedColumnDB)
		*displayedcolumnDB = *tmp
		
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
		
	case *RowDB:
		tmp, ok := db.rowDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Row Unkown entry %d", i))
		}

		rowDB, _ := instanceDB.(*RowDB)
		*rowDB = *tmp
		
	case *TableDB:
		tmp, ok := db.tableDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Table Unkown entry %d", i))
		}

		tableDB, _ := instanceDB.(*TableDB)
		*tableDB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gongtable/go, Unkown type")
	}
	
	return db, nil
}

