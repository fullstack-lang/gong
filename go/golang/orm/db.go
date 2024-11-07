package orm

const DbTmpl = `// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"{{PkgPathRoot}}/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions{{` + string(rune(DBliteMapFieldDefinition)) + `}}
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init{{` + string(rune(DBliteMapFieldInit)) + `}}
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("{{PkgPathRoot}}, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create{{` + string(rune(DBliteMapFieldCreate)) + `}}
	default:
		return nil, errors.New("{{PkgPathRoot}}, unsupported type in Create")
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
		return nil, errors.New("{{PkgPathRoot}}, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete{{` + string(rune(DBliteMapFieldDelete)) + `}}
	default:
		return nil, errors.New("{{PkgPathRoot}}, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("{{PkgPathRoot}}, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete{{` + string(rune(DBliteMapFieldSave)) + `}}
	default:
		return nil, errors.New("{{PkgPathRoot}}, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("{{PkgPathRoot}}, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete{{` + string(rune(DBliteMapFieldUpdate)) + `}}
	default:
		return nil, errors.New("{{PkgPathRoot}}, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find{{` + string(rune(DBliteMapFieldFind)) + `}}
	default:
		return nil, errors.New("{{PkgPathRoot}}, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("{{PkgPathRoot}}, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("{{PkgPathRoot}}, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("{{PkgPathRoot}}, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first{{` + string(rune(DBliteMapFieldFirst)) + `}}
	default:
		return nil, errors.New("{{PkgPathRoot}}, Unkown type")
	}
	
	return db, nil
}

`

type DBliteInsertionPointId int

const (
	DBliteMapFieldDefinition DBliteInsertionPointId = iota
	DBliteMapFieldInit
	DBliteMapFieldCreate
	DBliteMapFieldDelete
	DBliteMapFieldSave
	DBliteMapFieldUpdate
	DBliteMapFieldFind
	DBliteMapFieldFirst
)

var DBliteSubTemplates map[string]string = // new line
map[string]string{

	string(rune(DBliteMapFieldDefinition)): `

	{{structname}}DBs map[uint]*{{Structname}}DB

	nextID{{Structname}}DB uint`,

	string(rune(DBliteMapFieldInit)): `

		{{structname}}DBs: make(map[uint]*{{Structname}}DB),`,

	string(rune(DBliteMapFieldCreate)): `
	case *{{Structname}}DB:
		db.nextID{{Structname}}DB++
		v.ID = db.nextID{{Structname}}DB
		db.{{structname}}DBs[v.ID] = v`,

	string(rune(DBliteMapFieldDelete)): `
	case *{{Structname}}DB:
		delete(db.{{structname}}DBs, v.ID)`,

	string(rune(DBliteMapFieldSave)): `
	case *{{Structname}}DB:
		db.{{structname}}DBs[v.ID] = v
		return db, nil`,

	string(rune(DBliteMapFieldUpdate)): `
	case *{{Structname}}DB:
		if existing, ok := db.{{structname}}DBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db {{Structname}} {{PkgPathRoot}}, record not found")
		}`,

	string(rune(DBliteMapFieldFind)): `
	case *[]{{Structname}}DB:
		*ptr = make([]{{Structname}}DB, 0, len(db.{{structname}}DBs))
		for _, v := range db.{{structname}}DBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil`,

	string(rune(DBliteMapFieldFirst)): `
	case *{{Structname}}DB:
		tmp, ok := db.{{structname}}DBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First {{Structname}} Unkown entry %d", i))
		}

		{{structname}}DB, _ := instanceDB.(*{{Structname}}DB)
		*{{structname}}DB = *tmp
		`,
}
