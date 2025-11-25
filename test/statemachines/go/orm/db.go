// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gong/test/statemachines/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	architectureDBs map[uint]*ArchitectureDB

	nextIDArchitectureDB uint

	diagramDBs map[uint]*DiagramDB

	nextIDDiagramDB uint

	killDBs map[uint]*KillDB

	nextIDKillDB uint

	messageDBs map[uint]*MessageDB

	nextIDMessageDB uint

	messagetypeDBs map[uint]*MessageTypeDB

	nextIDMessageTypeDB uint

	objectDBs map[uint]*ObjectDB

	nextIDObjectDB uint

	roleDBs map[uint]*RoleDB

	nextIDRoleDB uint

	stateDBs map[uint]*StateDB

	nextIDStateDB uint

	statemachineDBs map[uint]*StateMachineDB

	nextIDStateMachineDB uint

	stateshapeDBs map[uint]*StateShapeDB

	nextIDStateShapeDB uint

	transitionDBs map[uint]*TransitionDB

	nextIDTransitionDB uint

	transition_shapeDBs map[uint]*Transition_ShapeDB

	nextIDTransition_ShapeDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		architectureDBs: make(map[uint]*ArchitectureDB),

		diagramDBs: make(map[uint]*DiagramDB),

		killDBs: make(map[uint]*KillDB),

		messageDBs: make(map[uint]*MessageDB),

		messagetypeDBs: make(map[uint]*MessageTypeDB),

		objectDBs: make(map[uint]*ObjectDB),

		roleDBs: make(map[uint]*RoleDB),

		stateDBs: make(map[uint]*StateDB),

		statemachineDBs: make(map[uint]*StateMachineDB),

		stateshapeDBs: make(map[uint]*StateShapeDB),

		transitionDBs: make(map[uint]*TransitionDB),

		transition_shapeDBs: make(map[uint]*Transition_ShapeDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/test/statemachines/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *ArchitectureDB:
		db.nextIDArchitectureDB++
		v.ID = db.nextIDArchitectureDB
		db.architectureDBs[v.ID] = v
	case *DiagramDB:
		db.nextIDDiagramDB++
		v.ID = db.nextIDDiagramDB
		db.diagramDBs[v.ID] = v
	case *KillDB:
		db.nextIDKillDB++
		v.ID = db.nextIDKillDB
		db.killDBs[v.ID] = v
	case *MessageDB:
		db.nextIDMessageDB++
		v.ID = db.nextIDMessageDB
		db.messageDBs[v.ID] = v
	case *MessageTypeDB:
		db.nextIDMessageTypeDB++
		v.ID = db.nextIDMessageTypeDB
		db.messagetypeDBs[v.ID] = v
	case *ObjectDB:
		db.nextIDObjectDB++
		v.ID = db.nextIDObjectDB
		db.objectDBs[v.ID] = v
	case *RoleDB:
		db.nextIDRoleDB++
		v.ID = db.nextIDRoleDB
		db.roleDBs[v.ID] = v
	case *StateDB:
		db.nextIDStateDB++
		v.ID = db.nextIDStateDB
		db.stateDBs[v.ID] = v
	case *StateMachineDB:
		db.nextIDStateMachineDB++
		v.ID = db.nextIDStateMachineDB
		db.statemachineDBs[v.ID] = v
	case *StateShapeDB:
		db.nextIDStateShapeDB++
		v.ID = db.nextIDStateShapeDB
		db.stateshapeDBs[v.ID] = v
	case *TransitionDB:
		db.nextIDTransitionDB++
		v.ID = db.nextIDTransitionDB
		db.transitionDBs[v.ID] = v
	case *Transition_ShapeDB:
		db.nextIDTransition_ShapeDB++
		v.ID = db.nextIDTransition_ShapeDB
		db.transition_shapeDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/test/statemachines/go, unsupported type in Create")
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
		return nil, errors.New("github.com/fullstack-lang/gong/test/statemachines/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ArchitectureDB:
		delete(db.architectureDBs, v.ID)
	case *DiagramDB:
		delete(db.diagramDBs, v.ID)
	case *KillDB:
		delete(db.killDBs, v.ID)
	case *MessageDB:
		delete(db.messageDBs, v.ID)
	case *MessageTypeDB:
		delete(db.messagetypeDBs, v.ID)
	case *ObjectDB:
		delete(db.objectDBs, v.ID)
	case *RoleDB:
		delete(db.roleDBs, v.ID)
	case *StateDB:
		delete(db.stateDBs, v.ID)
	case *StateMachineDB:
		delete(db.statemachineDBs, v.ID)
	case *StateShapeDB:
		delete(db.stateshapeDBs, v.ID)
	case *TransitionDB:
		delete(db.transitionDBs, v.ID)
	case *Transition_ShapeDB:
		delete(db.transition_shapeDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/test/statemachines/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/test/statemachines/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ArchitectureDB:
		db.architectureDBs[v.ID] = v
		return db, nil
	case *DiagramDB:
		db.diagramDBs[v.ID] = v
		return db, nil
	case *KillDB:
		db.killDBs[v.ID] = v
		return db, nil
	case *MessageDB:
		db.messageDBs[v.ID] = v
		return db, nil
	case *MessageTypeDB:
		db.messagetypeDBs[v.ID] = v
		return db, nil
	case *ObjectDB:
		db.objectDBs[v.ID] = v
		return db, nil
	case *RoleDB:
		db.roleDBs[v.ID] = v
		return db, nil
	case *StateDB:
		db.stateDBs[v.ID] = v
		return db, nil
	case *StateMachineDB:
		db.statemachineDBs[v.ID] = v
		return db, nil
	case *StateShapeDB:
		db.stateshapeDBs[v.ID] = v
		return db, nil
	case *TransitionDB:
		db.transitionDBs[v.ID] = v
		return db, nil
	case *Transition_ShapeDB:
		db.transition_shapeDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/test/statemachines/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/test/statemachines/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ArchitectureDB:
		if existing, ok := db.architectureDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Architecture github.com/fullstack-lang/gong/test/statemachines/go, record not found")
		}
	case *DiagramDB:
		if existing, ok := db.diagramDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Diagram github.com/fullstack-lang/gong/test/statemachines/go, record not found")
		}
	case *KillDB:
		if existing, ok := db.killDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Kill github.com/fullstack-lang/gong/test/statemachines/go, record not found")
		}
	case *MessageDB:
		if existing, ok := db.messageDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Message github.com/fullstack-lang/gong/test/statemachines/go, record not found")
		}
	case *MessageTypeDB:
		if existing, ok := db.messagetypeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db MessageType github.com/fullstack-lang/gong/test/statemachines/go, record not found")
		}
	case *ObjectDB:
		if existing, ok := db.objectDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Object github.com/fullstack-lang/gong/test/statemachines/go, record not found")
		}
	case *RoleDB:
		if existing, ok := db.roleDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Role github.com/fullstack-lang/gong/test/statemachines/go, record not found")
		}
	case *StateDB:
		if existing, ok := db.stateDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db State github.com/fullstack-lang/gong/test/statemachines/go, record not found")
		}
	case *StateMachineDB:
		if existing, ok := db.statemachineDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db StateMachine github.com/fullstack-lang/gong/test/statemachines/go, record not found")
		}
	case *StateShapeDB:
		if existing, ok := db.stateshapeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db StateShape github.com/fullstack-lang/gong/test/statemachines/go, record not found")
		}
	case *TransitionDB:
		if existing, ok := db.transitionDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Transition github.com/fullstack-lang/gong/test/statemachines/go, record not found")
		}
	case *Transition_ShapeDB:
		if existing, ok := db.transition_shapeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Transition_Shape github.com/fullstack-lang/gong/test/statemachines/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/test/statemachines/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]ArchitectureDB:
		*ptr = make([]ArchitectureDB, 0, len(db.architectureDBs))
		for _, v := range db.architectureDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DiagramDB:
		*ptr = make([]DiagramDB, 0, len(db.diagramDBs))
		for _, v := range db.diagramDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]KillDB:
		*ptr = make([]KillDB, 0, len(db.killDBs))
		for _, v := range db.killDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]MessageDB:
		*ptr = make([]MessageDB, 0, len(db.messageDBs))
		for _, v := range db.messageDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]MessageTypeDB:
		*ptr = make([]MessageTypeDB, 0, len(db.messagetypeDBs))
		for _, v := range db.messagetypeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ObjectDB:
		*ptr = make([]ObjectDB, 0, len(db.objectDBs))
		for _, v := range db.objectDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]RoleDB:
		*ptr = make([]RoleDB, 0, len(db.roleDBs))
		for _, v := range db.roleDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]StateDB:
		*ptr = make([]StateDB, 0, len(db.stateDBs))
		for _, v := range db.stateDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]StateMachineDB:
		*ptr = make([]StateMachineDB, 0, len(db.statemachineDBs))
		for _, v := range db.statemachineDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]StateShapeDB:
		*ptr = make([]StateShapeDB, 0, len(db.stateshapeDBs))
		for _, v := range db.stateshapeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]TransitionDB:
		*ptr = make([]TransitionDB, 0, len(db.transitionDBs))
		for _, v := range db.transitionDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Transition_ShapeDB:
		*ptr = make([]Transition_ShapeDB, 0, len(db.transition_shapeDBs))
		for _, v := range db.transition_shapeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/test/statemachines/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gong/test/statemachines/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gong/test/statemachines/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/test/statemachines/go, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *ArchitectureDB:
		tmp, ok := db.architectureDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Architecture Unkown entry %d", i))
		}

		architectureDB, _ := instanceDB.(*ArchitectureDB)
		*architectureDB = *tmp
		
	case *DiagramDB:
		tmp, ok := db.diagramDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Diagram Unkown entry %d", i))
		}

		diagramDB, _ := instanceDB.(*DiagramDB)
		*diagramDB = *tmp
		
	case *KillDB:
		tmp, ok := db.killDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Kill Unkown entry %d", i))
		}

		killDB, _ := instanceDB.(*KillDB)
		*killDB = *tmp
		
	case *MessageDB:
		tmp, ok := db.messageDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Message Unkown entry %d", i))
		}

		messageDB, _ := instanceDB.(*MessageDB)
		*messageDB = *tmp
		
	case *MessageTypeDB:
		tmp, ok := db.messagetypeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First MessageType Unkown entry %d", i))
		}

		messagetypeDB, _ := instanceDB.(*MessageTypeDB)
		*messagetypeDB = *tmp
		
	case *ObjectDB:
		tmp, ok := db.objectDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Object Unkown entry %d", i))
		}

		objectDB, _ := instanceDB.(*ObjectDB)
		*objectDB = *tmp
		
	case *RoleDB:
		tmp, ok := db.roleDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Role Unkown entry %d", i))
		}

		roleDB, _ := instanceDB.(*RoleDB)
		*roleDB = *tmp
		
	case *StateDB:
		tmp, ok := db.stateDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First State Unkown entry %d", i))
		}

		stateDB, _ := instanceDB.(*StateDB)
		*stateDB = *tmp
		
	case *StateMachineDB:
		tmp, ok := db.statemachineDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First StateMachine Unkown entry %d", i))
		}

		statemachineDB, _ := instanceDB.(*StateMachineDB)
		*statemachineDB = *tmp
		
	case *StateShapeDB:
		tmp, ok := db.stateshapeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First StateShape Unkown entry %d", i))
		}

		stateshapeDB, _ := instanceDB.(*StateShapeDB)
		*stateshapeDB = *tmp
		
	case *TransitionDB:
		tmp, ok := db.transitionDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Transition Unkown entry %d", i))
		}

		transitionDB, _ := instanceDB.(*TransitionDB)
		*transitionDB = *tmp
		
	case *Transition_ShapeDB:
		tmp, ok := db.transition_shapeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Transition_Shape Unkown entry %d", i))
		}

		transition_shapeDB, _ := instanceDB.(*Transition_ShapeDB)
		*transition_shapeDB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/test/statemachines/go, Unkown type")
	}
	
	return db, nil
}

