// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gongsvg/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	animateDBs map[uint]*AnimateDB

	nextIDAnimateDB uint

	circleDBs map[uint]*CircleDB

	nextIDCircleDB uint

	ellipseDBs map[uint]*EllipseDB

	nextIDEllipseDB uint

	layerDBs map[uint]*LayerDB

	nextIDLayerDB uint

	lineDBs map[uint]*LineDB

	nextIDLineDB uint

	linkDBs map[uint]*LinkDB

	nextIDLinkDB uint

	linkanchoredtextDBs map[uint]*LinkAnchoredTextDB

	nextIDLinkAnchoredTextDB uint

	pathDBs map[uint]*PathDB

	nextIDPathDB uint

	pointDBs map[uint]*PointDB

	nextIDPointDB uint

	polygoneDBs map[uint]*PolygoneDB

	nextIDPolygoneDB uint

	polylineDBs map[uint]*PolylineDB

	nextIDPolylineDB uint

	rectDBs map[uint]*RectDB

	nextIDRectDB uint

	rectanchoredpathDBs map[uint]*RectAnchoredPathDB

	nextIDRectAnchoredPathDB uint

	rectanchoredrectDBs map[uint]*RectAnchoredRectDB

	nextIDRectAnchoredRectDB uint

	rectanchoredtextDBs map[uint]*RectAnchoredTextDB

	nextIDRectAnchoredTextDB uint

	rectlinklinkDBs map[uint]*RectLinkLinkDB

	nextIDRectLinkLinkDB uint

	svgDBs map[uint]*SVGDB

	nextIDSVGDB uint

	textDBs map[uint]*TextDB

	nextIDTextDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		animateDBs: make(map[uint]*AnimateDB),

		circleDBs: make(map[uint]*CircleDB),

		ellipseDBs: make(map[uint]*EllipseDB),

		layerDBs: make(map[uint]*LayerDB),

		lineDBs: make(map[uint]*LineDB),

		linkDBs: make(map[uint]*LinkDB),

		linkanchoredtextDBs: make(map[uint]*LinkAnchoredTextDB),

		pathDBs: make(map[uint]*PathDB),

		pointDBs: make(map[uint]*PointDB),

		polygoneDBs: make(map[uint]*PolygoneDB),

		polylineDBs: make(map[uint]*PolylineDB),

		rectDBs: make(map[uint]*RectDB),

		rectanchoredpathDBs: make(map[uint]*RectAnchoredPathDB),

		rectanchoredrectDBs: make(map[uint]*RectAnchoredRectDB),

		rectanchoredtextDBs: make(map[uint]*RectAnchoredTextDB),

		rectlinklinkDBs: make(map[uint]*RectLinkLinkDB),

		svgDBs: make(map[uint]*SVGDB),

		textDBs: make(map[uint]*TextDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongsvg/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *AnimateDB:
		db.nextIDAnimateDB++
		v.ID = db.nextIDAnimateDB
		db.animateDBs[v.ID] = v
	case *CircleDB:
		db.nextIDCircleDB++
		v.ID = db.nextIDCircleDB
		db.circleDBs[v.ID] = v
	case *EllipseDB:
		db.nextIDEllipseDB++
		v.ID = db.nextIDEllipseDB
		db.ellipseDBs[v.ID] = v
	case *LayerDB:
		db.nextIDLayerDB++
		v.ID = db.nextIDLayerDB
		db.layerDBs[v.ID] = v
	case *LineDB:
		db.nextIDLineDB++
		v.ID = db.nextIDLineDB
		db.lineDBs[v.ID] = v
	case *LinkDB:
		db.nextIDLinkDB++
		v.ID = db.nextIDLinkDB
		db.linkDBs[v.ID] = v
	case *LinkAnchoredTextDB:
		db.nextIDLinkAnchoredTextDB++
		v.ID = db.nextIDLinkAnchoredTextDB
		db.linkanchoredtextDBs[v.ID] = v
	case *PathDB:
		db.nextIDPathDB++
		v.ID = db.nextIDPathDB
		db.pathDBs[v.ID] = v
	case *PointDB:
		db.nextIDPointDB++
		v.ID = db.nextIDPointDB
		db.pointDBs[v.ID] = v
	case *PolygoneDB:
		db.nextIDPolygoneDB++
		v.ID = db.nextIDPolygoneDB
		db.polygoneDBs[v.ID] = v
	case *PolylineDB:
		db.nextIDPolylineDB++
		v.ID = db.nextIDPolylineDB
		db.polylineDBs[v.ID] = v
	case *RectDB:
		db.nextIDRectDB++
		v.ID = db.nextIDRectDB
		db.rectDBs[v.ID] = v
	case *RectAnchoredPathDB:
		db.nextIDRectAnchoredPathDB++
		v.ID = db.nextIDRectAnchoredPathDB
		db.rectanchoredpathDBs[v.ID] = v
	case *RectAnchoredRectDB:
		db.nextIDRectAnchoredRectDB++
		v.ID = db.nextIDRectAnchoredRectDB
		db.rectanchoredrectDBs[v.ID] = v
	case *RectAnchoredTextDB:
		db.nextIDRectAnchoredTextDB++
		v.ID = db.nextIDRectAnchoredTextDB
		db.rectanchoredtextDBs[v.ID] = v
	case *RectLinkLinkDB:
		db.nextIDRectLinkLinkDB++
		v.ID = db.nextIDRectLinkLinkDB
		db.rectlinklinkDBs[v.ID] = v
	case *SVGDB:
		db.nextIDSVGDB++
		v.ID = db.nextIDSVGDB
		db.svgDBs[v.ID] = v
	case *TextDB:
		db.nextIDTextDB++
		v.ID = db.nextIDTextDB
		db.textDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gongsvg/go, unsupported type in Create")
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
		return nil, errors.New("github.com/fullstack-lang/gongsvg/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AnimateDB:
		delete(db.animateDBs, v.ID)
	case *CircleDB:
		delete(db.circleDBs, v.ID)
	case *EllipseDB:
		delete(db.ellipseDBs, v.ID)
	case *LayerDB:
		delete(db.layerDBs, v.ID)
	case *LineDB:
		delete(db.lineDBs, v.ID)
	case *LinkDB:
		delete(db.linkDBs, v.ID)
	case *LinkAnchoredTextDB:
		delete(db.linkanchoredtextDBs, v.ID)
	case *PathDB:
		delete(db.pathDBs, v.ID)
	case *PointDB:
		delete(db.pointDBs, v.ID)
	case *PolygoneDB:
		delete(db.polygoneDBs, v.ID)
	case *PolylineDB:
		delete(db.polylineDBs, v.ID)
	case *RectDB:
		delete(db.rectDBs, v.ID)
	case *RectAnchoredPathDB:
		delete(db.rectanchoredpathDBs, v.ID)
	case *RectAnchoredRectDB:
		delete(db.rectanchoredrectDBs, v.ID)
	case *RectAnchoredTextDB:
		delete(db.rectanchoredtextDBs, v.ID)
	case *RectLinkLinkDB:
		delete(db.rectlinklinkDBs, v.ID)
	case *SVGDB:
		delete(db.svgDBs, v.ID)
	case *TextDB:
		delete(db.textDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongsvg/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongsvg/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AnimateDB:
		db.animateDBs[v.ID] = v
		return db, nil
	case *CircleDB:
		db.circleDBs[v.ID] = v
		return db, nil
	case *EllipseDB:
		db.ellipseDBs[v.ID] = v
		return db, nil
	case *LayerDB:
		db.layerDBs[v.ID] = v
		return db, nil
	case *LineDB:
		db.lineDBs[v.ID] = v
		return db, nil
	case *LinkDB:
		db.linkDBs[v.ID] = v
		return db, nil
	case *LinkAnchoredTextDB:
		db.linkanchoredtextDBs[v.ID] = v
		return db, nil
	case *PathDB:
		db.pathDBs[v.ID] = v
		return db, nil
	case *PointDB:
		db.pointDBs[v.ID] = v
		return db, nil
	case *PolygoneDB:
		db.polygoneDBs[v.ID] = v
		return db, nil
	case *PolylineDB:
		db.polylineDBs[v.ID] = v
		return db, nil
	case *RectDB:
		db.rectDBs[v.ID] = v
		return db, nil
	case *RectAnchoredPathDB:
		db.rectanchoredpathDBs[v.ID] = v
		return db, nil
	case *RectAnchoredRectDB:
		db.rectanchoredrectDBs[v.ID] = v
		return db, nil
	case *RectAnchoredTextDB:
		db.rectanchoredtextDBs[v.ID] = v
		return db, nil
	case *RectLinkLinkDB:
		db.rectlinklinkDBs[v.ID] = v
		return db, nil
	case *SVGDB:
		db.svgDBs[v.ID] = v
		return db, nil
	case *TextDB:
		db.textDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongsvg/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongsvg/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AnimateDB:
		if existing, ok := db.animateDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Animate github.com/fullstack-lang/gongsvg/go, record not found")
		}
	case *CircleDB:
		if existing, ok := db.circleDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Circle github.com/fullstack-lang/gongsvg/go, record not found")
		}
	case *EllipseDB:
		if existing, ok := db.ellipseDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Ellipse github.com/fullstack-lang/gongsvg/go, record not found")
		}
	case *LayerDB:
		if existing, ok := db.layerDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Layer github.com/fullstack-lang/gongsvg/go, record not found")
		}
	case *LineDB:
		if existing, ok := db.lineDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Line github.com/fullstack-lang/gongsvg/go, record not found")
		}
	case *LinkDB:
		if existing, ok := db.linkDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Link github.com/fullstack-lang/gongsvg/go, record not found")
		}
	case *LinkAnchoredTextDB:
		if existing, ok := db.linkanchoredtextDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db LinkAnchoredText github.com/fullstack-lang/gongsvg/go, record not found")
		}
	case *PathDB:
		if existing, ok := db.pathDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Path github.com/fullstack-lang/gongsvg/go, record not found")
		}
	case *PointDB:
		if existing, ok := db.pointDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Point github.com/fullstack-lang/gongsvg/go, record not found")
		}
	case *PolygoneDB:
		if existing, ok := db.polygoneDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Polygone github.com/fullstack-lang/gongsvg/go, record not found")
		}
	case *PolylineDB:
		if existing, ok := db.polylineDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Polyline github.com/fullstack-lang/gongsvg/go, record not found")
		}
	case *RectDB:
		if existing, ok := db.rectDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Rect github.com/fullstack-lang/gongsvg/go, record not found")
		}
	case *RectAnchoredPathDB:
		if existing, ok := db.rectanchoredpathDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db RectAnchoredPath github.com/fullstack-lang/gongsvg/go, record not found")
		}
	case *RectAnchoredRectDB:
		if existing, ok := db.rectanchoredrectDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db RectAnchoredRect github.com/fullstack-lang/gongsvg/go, record not found")
		}
	case *RectAnchoredTextDB:
		if existing, ok := db.rectanchoredtextDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db RectAnchoredText github.com/fullstack-lang/gongsvg/go, record not found")
		}
	case *RectLinkLinkDB:
		if existing, ok := db.rectlinklinkDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db RectLinkLink github.com/fullstack-lang/gongsvg/go, record not found")
		}
	case *SVGDB:
		if existing, ok := db.svgDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SVG github.com/fullstack-lang/gongsvg/go, record not found")
		}
	case *TextDB:
		if existing, ok := db.textDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Text github.com/fullstack-lang/gongsvg/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gongsvg/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]AnimateDB:
		*ptr = make([]AnimateDB, 0, len(db.animateDBs))
		for _, v := range db.animateDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]CircleDB:
		*ptr = make([]CircleDB, 0, len(db.circleDBs))
		for _, v := range db.circleDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]EllipseDB:
		*ptr = make([]EllipseDB, 0, len(db.ellipseDBs))
		for _, v := range db.ellipseDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]LayerDB:
		*ptr = make([]LayerDB, 0, len(db.layerDBs))
		for _, v := range db.layerDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]LineDB:
		*ptr = make([]LineDB, 0, len(db.lineDBs))
		for _, v := range db.lineDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]LinkDB:
		*ptr = make([]LinkDB, 0, len(db.linkDBs))
		for _, v := range db.linkDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]LinkAnchoredTextDB:
		*ptr = make([]LinkAnchoredTextDB, 0, len(db.linkanchoredtextDBs))
		for _, v := range db.linkanchoredtextDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]PathDB:
		*ptr = make([]PathDB, 0, len(db.pathDBs))
		for _, v := range db.pathDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]PointDB:
		*ptr = make([]PointDB, 0, len(db.pointDBs))
		for _, v := range db.pointDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]PolygoneDB:
		*ptr = make([]PolygoneDB, 0, len(db.polygoneDBs))
		for _, v := range db.polygoneDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]PolylineDB:
		*ptr = make([]PolylineDB, 0, len(db.polylineDBs))
		for _, v := range db.polylineDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]RectDB:
		*ptr = make([]RectDB, 0, len(db.rectDBs))
		for _, v := range db.rectDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]RectAnchoredPathDB:
		*ptr = make([]RectAnchoredPathDB, 0, len(db.rectanchoredpathDBs))
		for _, v := range db.rectanchoredpathDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]RectAnchoredRectDB:
		*ptr = make([]RectAnchoredRectDB, 0, len(db.rectanchoredrectDBs))
		for _, v := range db.rectanchoredrectDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]RectAnchoredTextDB:
		*ptr = make([]RectAnchoredTextDB, 0, len(db.rectanchoredtextDBs))
		for _, v := range db.rectanchoredtextDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]RectLinkLinkDB:
		*ptr = make([]RectLinkLinkDB, 0, len(db.rectlinklinkDBs))
		for _, v := range db.rectlinklinkDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SVGDB:
		*ptr = make([]SVGDB, 0, len(db.svgDBs))
		for _, v := range db.svgDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]TextDB:
		*ptr = make([]TextDB, 0, len(db.textDBs))
		for _, v := range db.textDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongsvg/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gongsvg/go, Do not process when conds is not a single parameter")
	}

	str, ok := conds[0].(string)

	if !ok {
		return nil, errors.New("github.com/fullstack-lang/gongsvg/go, conds[0] is not a string")
	}

	i, err := strconv.ParseUint(str, 10, 32) // Base 10, 32-bit unsigned int
	if err != nil {
		return nil, errors.New("github.com/fullstack-lang/gongsvg/go, conds[0] is not a string number")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *AnimateDB:
		tmp, ok := db.animateDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Animate Unkown entry %d", i))
		}

		animateDB, _ := instanceDB.(*AnimateDB)
		*animateDB = *tmp
		
	case *CircleDB:
		tmp, ok := db.circleDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Circle Unkown entry %d", i))
		}

		circleDB, _ := instanceDB.(*CircleDB)
		*circleDB = *tmp
		
	case *EllipseDB:
		tmp, ok := db.ellipseDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Ellipse Unkown entry %d", i))
		}

		ellipseDB, _ := instanceDB.(*EllipseDB)
		*ellipseDB = *tmp
		
	case *LayerDB:
		tmp, ok := db.layerDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Layer Unkown entry %d", i))
		}

		layerDB, _ := instanceDB.(*LayerDB)
		*layerDB = *tmp
		
	case *LineDB:
		tmp, ok := db.lineDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Line Unkown entry %d", i))
		}

		lineDB, _ := instanceDB.(*LineDB)
		*lineDB = *tmp
		
	case *LinkDB:
		tmp, ok := db.linkDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Link Unkown entry %d", i))
		}

		linkDB, _ := instanceDB.(*LinkDB)
		*linkDB = *tmp
		
	case *LinkAnchoredTextDB:
		tmp, ok := db.linkanchoredtextDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First LinkAnchoredText Unkown entry %d", i))
		}

		linkanchoredtextDB, _ := instanceDB.(*LinkAnchoredTextDB)
		*linkanchoredtextDB = *tmp
		
	case *PathDB:
		tmp, ok := db.pathDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Path Unkown entry %d", i))
		}

		pathDB, _ := instanceDB.(*PathDB)
		*pathDB = *tmp
		
	case *PointDB:
		tmp, ok := db.pointDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Point Unkown entry %d", i))
		}

		pointDB, _ := instanceDB.(*PointDB)
		*pointDB = *tmp
		
	case *PolygoneDB:
		tmp, ok := db.polygoneDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Polygone Unkown entry %d", i))
		}

		polygoneDB, _ := instanceDB.(*PolygoneDB)
		*polygoneDB = *tmp
		
	case *PolylineDB:
		tmp, ok := db.polylineDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Polyline Unkown entry %d", i))
		}

		polylineDB, _ := instanceDB.(*PolylineDB)
		*polylineDB = *tmp
		
	case *RectDB:
		tmp, ok := db.rectDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Rect Unkown entry %d", i))
		}

		rectDB, _ := instanceDB.(*RectDB)
		*rectDB = *tmp
		
	case *RectAnchoredPathDB:
		tmp, ok := db.rectanchoredpathDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First RectAnchoredPath Unkown entry %d", i))
		}

		rectanchoredpathDB, _ := instanceDB.(*RectAnchoredPathDB)
		*rectanchoredpathDB = *tmp
		
	case *RectAnchoredRectDB:
		tmp, ok := db.rectanchoredrectDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First RectAnchoredRect Unkown entry %d", i))
		}

		rectanchoredrectDB, _ := instanceDB.(*RectAnchoredRectDB)
		*rectanchoredrectDB = *tmp
		
	case *RectAnchoredTextDB:
		tmp, ok := db.rectanchoredtextDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First RectAnchoredText Unkown entry %d", i))
		}

		rectanchoredtextDB, _ := instanceDB.(*RectAnchoredTextDB)
		*rectanchoredtextDB = *tmp
		
	case *RectLinkLinkDB:
		tmp, ok := db.rectlinklinkDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First RectLinkLink Unkown entry %d", i))
		}

		rectlinklinkDB, _ := instanceDB.(*RectLinkLinkDB)
		*rectlinklinkDB = *tmp
		
	case *SVGDB:
		tmp, ok := db.svgDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SVG Unkown entry %d", i))
		}

		svgDB, _ := instanceDB.(*SVGDB)
		*svgDB = *tmp
		
	case *TextDB:
		tmp, ok := db.textDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Text Unkown entry %d", i))
		}

		textDB, _ := instanceDB.(*TextDB)
		*textDB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gongsvg/go, Unkown type")
	}
	
	return db, nil
}

