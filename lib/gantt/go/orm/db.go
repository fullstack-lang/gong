// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gong/lib/gantt/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	arrowDBs map[uint]*ArrowDB

	nextIDArrowDB uint

	barDBs map[uint]*BarDB

	nextIDBarDB uint

	ganttDBs map[uint]*GanttDB

	nextIDGanttDB uint

	groupDBs map[uint]*GroupDB

	nextIDGroupDB uint

	laneDBs map[uint]*LaneDB

	nextIDLaneDB uint

	laneuseDBs map[uint]*LaneUseDB

	nextIDLaneUseDB uint

	milestoneDBs map[uint]*MilestoneDB

	nextIDMilestoneDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		arrowDBs: make(map[uint]*ArrowDB),

		barDBs: make(map[uint]*BarDB),

		ganttDBs: make(map[uint]*GanttDB),

		groupDBs: make(map[uint]*GroupDB),

		laneDBs: make(map[uint]*LaneDB),

		laneuseDBs: make(map[uint]*LaneUseDB),

		milestoneDBs: make(map[uint]*MilestoneDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/gantt/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *ArrowDB:
		db.nextIDArrowDB++
		v.ID = db.nextIDArrowDB
		db.arrowDBs[v.ID] = v
	case *BarDB:
		db.nextIDBarDB++
		v.ID = db.nextIDBarDB
		db.barDBs[v.ID] = v
	case *GanttDB:
		db.nextIDGanttDB++
		v.ID = db.nextIDGanttDB
		db.ganttDBs[v.ID] = v
	case *GroupDB:
		db.nextIDGroupDB++
		v.ID = db.nextIDGroupDB
		db.groupDBs[v.ID] = v
	case *LaneDB:
		db.nextIDLaneDB++
		v.ID = db.nextIDLaneDB
		db.laneDBs[v.ID] = v
	case *LaneUseDB:
		db.nextIDLaneUseDB++
		v.ID = db.nextIDLaneUseDB
		db.laneuseDBs[v.ID] = v
	case *MilestoneDB:
		db.nextIDMilestoneDB++
		v.ID = db.nextIDMilestoneDB
		db.milestoneDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/gantt/go, unsupported type in Create")
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
		return nil, errors.New("github.com/fullstack-lang/gong/lib/gantt/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ArrowDB:
		delete(db.arrowDBs, v.ID)
	case *BarDB:
		delete(db.barDBs, v.ID)
	case *GanttDB:
		delete(db.ganttDBs, v.ID)
	case *GroupDB:
		delete(db.groupDBs, v.ID)
	case *LaneDB:
		delete(db.laneDBs, v.ID)
	case *LaneUseDB:
		delete(db.laneuseDBs, v.ID)
	case *MilestoneDB:
		delete(db.milestoneDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/gantt/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/gantt/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ArrowDB:
		db.arrowDBs[v.ID] = v
		return db, nil
	case *BarDB:
		db.barDBs[v.ID] = v
		return db, nil
	case *GanttDB:
		db.ganttDBs[v.ID] = v
		return db, nil
	case *GroupDB:
		db.groupDBs[v.ID] = v
		return db, nil
	case *LaneDB:
		db.laneDBs[v.ID] = v
		return db, nil
	case *LaneUseDB:
		db.laneuseDBs[v.ID] = v
		return db, nil
	case *MilestoneDB:
		db.milestoneDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/gantt/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/gantt/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ArrowDB:
		if existing, ok := db.arrowDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Arrow github.com/fullstack-lang/gong/lib/gantt/go, record not found")
		}
	case *BarDB:
		if existing, ok := db.barDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Bar github.com/fullstack-lang/gong/lib/gantt/go, record not found")
		}
	case *GanttDB:
		if existing, ok := db.ganttDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Gantt github.com/fullstack-lang/gong/lib/gantt/go, record not found")
		}
	case *GroupDB:
		if existing, ok := db.groupDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Group github.com/fullstack-lang/gong/lib/gantt/go, record not found")
		}
	case *LaneDB:
		if existing, ok := db.laneDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Lane github.com/fullstack-lang/gong/lib/gantt/go, record not found")
		}
	case *LaneUseDB:
		if existing, ok := db.laneuseDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db LaneUse github.com/fullstack-lang/gong/lib/gantt/go, record not found")
		}
	case *MilestoneDB:
		if existing, ok := db.milestoneDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Milestone github.com/fullstack-lang/gong/lib/gantt/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/gantt/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]ArrowDB:
		*ptr = make([]ArrowDB, 0, len(db.arrowDBs))
		for _, v := range db.arrowDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]BarDB:
		*ptr = make([]BarDB, 0, len(db.barDBs))
		for _, v := range db.barDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]GanttDB:
		*ptr = make([]GanttDB, 0, len(db.ganttDBs))
		for _, v := range db.ganttDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]GroupDB:
		*ptr = make([]GroupDB, 0, len(db.groupDBs))
		for _, v := range db.groupDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]LaneDB:
		*ptr = make([]LaneDB, 0, len(db.laneDBs))
		for _, v := range db.laneDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]LaneUseDB:
		*ptr = make([]LaneUseDB, 0, len(db.laneuseDBs))
		for _, v := range db.laneuseDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]MilestoneDB:
		*ptr = make([]MilestoneDB, 0, len(db.milestoneDBs))
		for _, v := range db.milestoneDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/gantt/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gong/lib/gantt/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gong/lib/gantt/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/gantt/go, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *ArrowDB:
		tmp, ok := db.arrowDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Arrow Unkown entry %d", i))
		}

		arrowDB, _ := instanceDB.(*ArrowDB)
		*arrowDB = *tmp
		
	case *BarDB:
		tmp, ok := db.barDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Bar Unkown entry %d", i))
		}

		barDB, _ := instanceDB.(*BarDB)
		*barDB = *tmp
		
	case *GanttDB:
		tmp, ok := db.ganttDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Gantt Unkown entry %d", i))
		}

		ganttDB, _ := instanceDB.(*GanttDB)
		*ganttDB = *tmp
		
	case *GroupDB:
		tmp, ok := db.groupDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Group Unkown entry %d", i))
		}

		groupDB, _ := instanceDB.(*GroupDB)
		*groupDB = *tmp
		
	case *LaneDB:
		tmp, ok := db.laneDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Lane Unkown entry %d", i))
		}

		laneDB, _ := instanceDB.(*LaneDB)
		*laneDB = *tmp
		
	case *LaneUseDB:
		tmp, ok := db.laneuseDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First LaneUse Unkown entry %d", i))
		}

		laneuseDB, _ := instanceDB.(*LaneUseDB)
		*laneuseDB = *tmp
		
	case *MilestoneDB:
		tmp, ok := db.milestoneDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Milestone Unkown entry %d", i))
		}

		milestoneDB, _ := instanceDB.(*MilestoneDB)
		*milestoneDB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gong/lib/gantt/go, Unkown type")
	}
	
	return db, nil
}

