// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gong/lib/svg/go/db"
	"github.com/fullstack-lang/gong/lib/svg/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Polyline_sql sql.NullBool
var dummy_Polyline_time time.Duration
var dummy_Polyline_sort sort.Float64Slice

// PolylineAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model polylineAPI
type PolylineAPI struct {
	gorm.Model

	models.Polyline_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	PolylinePointersEncoding PolylinePointersEncoding
}

// PolylinePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type PolylinePointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field Animates is a slice of pointers to another Struct (optional or 0..1)
	Animates IntSlice `gorm:"type:TEXT"`
}

// PolylineDB describes a polyline in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model polylineDB
type PolylineDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field polylineDB.Name
	Name_Data sql.NullString

	// Declation for basic field polylineDB.Points
	Points_Data sql.NullString

	// Declation for basic field polylineDB.Color
	Color_Data sql.NullString

	// Declation for basic field polylineDB.FillOpacity
	FillOpacity_Data sql.NullFloat64

	// Declation for basic field polylineDB.Stroke
	Stroke_Data sql.NullString

	// Declation for basic field polylineDB.StrokeOpacity
	StrokeOpacity_Data sql.NullFloat64

	// Declation for basic field polylineDB.StrokeWidth
	StrokeWidth_Data sql.NullFloat64

	// Declation for basic field polylineDB.StrokeDashArray
	StrokeDashArray_Data sql.NullString

	// Declation for basic field polylineDB.StrokeDashArrayWhenSelected
	StrokeDashArrayWhenSelected_Data sql.NullString

	// Declation for basic field polylineDB.Transform
	Transform_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	PolylinePointersEncoding
}

// PolylineDBs arrays polylineDBs
// swagger:response polylineDBsResponse
type PolylineDBs []PolylineDB

// PolylineDBResponse provides response
// swagger:response polylineDBResponse
type PolylineDBResponse struct {
	PolylineDB
}

// PolylineWOP is a Polyline without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type PolylineWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Points string `xlsx:"2"`

	Color string `xlsx:"3"`

	FillOpacity float64 `xlsx:"4"`

	Stroke string `xlsx:"5"`

	StrokeOpacity float64 `xlsx:"6"`

	StrokeWidth float64 `xlsx:"7"`

	StrokeDashArray string `xlsx:"8"`

	StrokeDashArrayWhenSelected string `xlsx:"9"`

	Transform string `xlsx:"10"`
	// insertion for WOP pointer fields
}

var Polyline_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Points",
	"Color",
	"FillOpacity",
	"Stroke",
	"StrokeOpacity",
	"StrokeWidth",
	"StrokeDashArray",
	"StrokeDashArrayWhenSelected",
	"Transform",
}

type BackRepoPolylineStruct struct {
	// stores PolylineDB according to their gorm ID
	Map_PolylineDBID_PolylineDB map[uint]*PolylineDB

	// stores PolylineDB ID according to Polyline address
	Map_PolylinePtr_PolylineDBID map[*models.Polyline]uint

	// stores Polyline according to their gorm ID
	Map_PolylineDBID_PolylinePtr map[uint]*models.Polyline

	db db.DBInterface

	stage *models.Stage
}

func (backRepoPolyline *BackRepoPolylineStruct) GetStage() (stage *models.Stage) {
	stage = backRepoPolyline.stage
	return
}

func (backRepoPolyline *BackRepoPolylineStruct) GetDB() db.DBInterface {
	return backRepoPolyline.db
}

// GetPolylineDBFromPolylinePtr is a handy function to access the back repo instance from the stage instance
func (backRepoPolyline *BackRepoPolylineStruct) GetPolylineDBFromPolylinePtr(polyline *models.Polyline) (polylineDB *PolylineDB) {
	id := backRepoPolyline.Map_PolylinePtr_PolylineDBID[polyline]
	polylineDB = backRepoPolyline.Map_PolylineDBID_PolylineDB[id]
	return
}

// BackRepoPolyline.CommitPhaseOne commits all staged instances of Polyline to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoPolyline *BackRepoPolylineStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var polylines []*models.Polyline
	for polyline := range stage.Polylines {
		polylines = append(polylines, polyline)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(polylines, func(i, j int) bool {
		return stage.PolylineMap_Staged_Order[polylines[i]] < stage.PolylineMap_Staged_Order[polylines[j]]
	})

	for _, polyline := range polylines {
		backRepoPolyline.CommitPhaseOneInstance(polyline)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, polyline := range backRepoPolyline.Map_PolylineDBID_PolylinePtr {
		if _, ok := stage.Polylines[polyline]; !ok {
			backRepoPolyline.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoPolyline.CommitDeleteInstance commits deletion of Polyline to the BackRepo
func (backRepoPolyline *BackRepoPolylineStruct) CommitDeleteInstance(id uint) (Error error) {

	polyline := backRepoPolyline.Map_PolylineDBID_PolylinePtr[id]

	// polyline is not staged anymore, remove polylineDB
	polylineDB := backRepoPolyline.Map_PolylineDBID_PolylineDB[id]
	db, _ := backRepoPolyline.db.Unscoped()
	_, err := db.Delete(polylineDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoPolyline.Map_PolylinePtr_PolylineDBID, polyline)
	delete(backRepoPolyline.Map_PolylineDBID_PolylinePtr, id)
	delete(backRepoPolyline.Map_PolylineDBID_PolylineDB, id)

	return
}

// BackRepoPolyline.CommitPhaseOneInstance commits polyline staged instances of Polyline to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoPolyline *BackRepoPolylineStruct) CommitPhaseOneInstance(polyline *models.Polyline) (Error error) {

	// check if the polyline is not commited yet
	if _, ok := backRepoPolyline.Map_PolylinePtr_PolylineDBID[polyline]; ok {
		return
	}

	// initiate polyline
	var polylineDB PolylineDB
	polylineDB.CopyBasicFieldsFromPolyline(polyline)

	_, err := backRepoPolyline.db.Create(&polylineDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoPolyline.Map_PolylinePtr_PolylineDBID[polyline] = polylineDB.ID
	backRepoPolyline.Map_PolylineDBID_PolylinePtr[polylineDB.ID] = polyline
	backRepoPolyline.Map_PolylineDBID_PolylineDB[polylineDB.ID] = &polylineDB

	return
}

// BackRepoPolyline.CommitPhaseTwo commits all staged instances of Polyline to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPolyline *BackRepoPolylineStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, polyline := range backRepoPolyline.Map_PolylineDBID_PolylinePtr {
		backRepoPolyline.CommitPhaseTwoInstance(backRepo, idx, polyline)
	}

	return
}

// BackRepoPolyline.CommitPhaseTwoInstance commits {{structname }} of models.Polyline to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPolyline *BackRepoPolylineStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, polyline *models.Polyline) (Error error) {

	// fetch matching polylineDB
	if polylineDB, ok := backRepoPolyline.Map_PolylineDBID_PolylineDB[idx]; ok {

		polylineDB.CopyBasicFieldsFromPolyline(polyline)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		polylineDB.PolylinePointersEncoding.Animates = make([]int, 0)
		// 2. encode
		for _, animateAssocEnd := range polyline.Animates {
			animateAssocEnd_DB :=
				backRepo.BackRepoAnimate.GetAnimateDBFromAnimatePtr(animateAssocEnd)
			
			// the stage might be inconsistant, meaning that the animateAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if animateAssocEnd_DB == nil {
				continue
			}
			
			polylineDB.PolylinePointersEncoding.Animates =
				append(polylineDB.PolylinePointersEncoding.Animates, int(animateAssocEnd_DB.ID))
		}

		_, err := backRepoPolyline.db.Save(polylineDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Polyline intance %s", polyline.Name))
		return err
	}

	return
}

// BackRepoPolyline.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoPolyline *BackRepoPolylineStruct) CheckoutPhaseOne() (Error error) {

	polylineDBArray := make([]PolylineDB, 0)
	_, err := backRepoPolyline.db.Find(&polylineDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	polylineInstancesToBeRemovedFromTheStage := make(map[*models.Polyline]any)
	for key, value := range backRepoPolyline.stage.Polylines {
		polylineInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, polylineDB := range polylineDBArray {
		backRepoPolyline.CheckoutPhaseOneInstance(&polylineDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		polyline, ok := backRepoPolyline.Map_PolylineDBID_PolylinePtr[polylineDB.ID]
		if ok {
			delete(polylineInstancesToBeRemovedFromTheStage, polyline)
		}
	}

	// remove from stage and back repo's 3 maps all polylines that are not in the checkout
	for polyline := range polylineInstancesToBeRemovedFromTheStage {
		polyline.Unstage(backRepoPolyline.GetStage())

		// remove instance from the back repo 3 maps
		polylineID := backRepoPolyline.Map_PolylinePtr_PolylineDBID[polyline]
		delete(backRepoPolyline.Map_PolylinePtr_PolylineDBID, polyline)
		delete(backRepoPolyline.Map_PolylineDBID_PolylineDB, polylineID)
		delete(backRepoPolyline.Map_PolylineDBID_PolylinePtr, polylineID)
	}

	return
}

// CheckoutPhaseOneInstance takes a polylineDB that has been found in the DB, updates the backRepo and stages the
// models version of the polylineDB
func (backRepoPolyline *BackRepoPolylineStruct) CheckoutPhaseOneInstance(polylineDB *PolylineDB) (Error error) {

	polyline, ok := backRepoPolyline.Map_PolylineDBID_PolylinePtr[polylineDB.ID]
	if !ok {
		polyline = new(models.Polyline)

		backRepoPolyline.Map_PolylineDBID_PolylinePtr[polylineDB.ID] = polyline
		backRepoPolyline.Map_PolylinePtr_PolylineDBID[polyline] = polylineDB.ID

		// append model store with the new element
		polyline.Name = polylineDB.Name_Data.String
		polyline.Stage(backRepoPolyline.GetStage())
	}
	polylineDB.CopyBasicFieldsToPolyline(polyline)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	polyline.Stage(backRepoPolyline.GetStage())

	// preserve pointer to polylineDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_PolylineDBID_PolylineDB)[polylineDB hold variable pointers
	polylineDB_Data := *polylineDB
	preservedPtrToPolyline := &polylineDB_Data
	backRepoPolyline.Map_PolylineDBID_PolylineDB[polylineDB.ID] = preservedPtrToPolyline

	return
}

// BackRepoPolyline.CheckoutPhaseTwo Checkouts all staged instances of Polyline to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPolyline *BackRepoPolylineStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, polylineDB := range backRepoPolyline.Map_PolylineDBID_PolylineDB {
		backRepoPolyline.CheckoutPhaseTwoInstance(backRepo, polylineDB)
	}
	return
}

// BackRepoPolyline.CheckoutPhaseTwoInstance Checkouts staged instances of Polyline to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPolyline *BackRepoPolylineStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, polylineDB *PolylineDB) (Error error) {

	polyline := backRepoPolyline.Map_PolylineDBID_PolylinePtr[polylineDB.ID]

	polylineDB.DecodePointers(backRepo, polyline)

	return
}

func (polylineDB *PolylineDB) DecodePointers(backRepo *BackRepoStruct, polyline *models.Polyline) {

	// insertion point for checkout of pointer encoding
	// This loop redeem polyline.Animates in the stage from the encode in the back repo
	// It parses all AnimateDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	polyline.Animates = polyline.Animates[:0]
	for _, _Animateid := range polylineDB.PolylinePointersEncoding.Animates {
		polyline.Animates = append(polyline.Animates, backRepo.BackRepoAnimate.Map_AnimateDBID_AnimatePtr[uint(_Animateid)])
	}

	return
}

// CommitPolyline allows commit of a single polyline (if already staged)
func (backRepo *BackRepoStruct) CommitPolyline(polyline *models.Polyline) {
	backRepo.BackRepoPolyline.CommitPhaseOneInstance(polyline)
	if id, ok := backRepo.BackRepoPolyline.Map_PolylinePtr_PolylineDBID[polyline]; ok {
		backRepo.BackRepoPolyline.CommitPhaseTwoInstance(backRepo, id, polyline)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitPolyline allows checkout of a single polyline (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutPolyline(polyline *models.Polyline) {
	// check if the polyline is staged
	if _, ok := backRepo.BackRepoPolyline.Map_PolylinePtr_PolylineDBID[polyline]; ok {

		if id, ok := backRepo.BackRepoPolyline.Map_PolylinePtr_PolylineDBID[polyline]; ok {
			var polylineDB PolylineDB
			polylineDB.ID = id

			if _, err := backRepo.BackRepoPolyline.db.First(&polylineDB, id); err != nil {
				log.Fatalln("CheckoutPolyline : Problem with getting object with id:", id)
			}
			backRepo.BackRepoPolyline.CheckoutPhaseOneInstance(&polylineDB)
			backRepo.BackRepoPolyline.CheckoutPhaseTwoInstance(backRepo, &polylineDB)
		}
	}
}

// CopyBasicFieldsFromPolyline
func (polylineDB *PolylineDB) CopyBasicFieldsFromPolyline(polyline *models.Polyline) {
	// insertion point for fields commit

	polylineDB.Name_Data.String = polyline.Name
	polylineDB.Name_Data.Valid = true

	polylineDB.Points_Data.String = polyline.Points
	polylineDB.Points_Data.Valid = true

	polylineDB.Color_Data.String = polyline.Color
	polylineDB.Color_Data.Valid = true

	polylineDB.FillOpacity_Data.Float64 = polyline.FillOpacity
	polylineDB.FillOpacity_Data.Valid = true

	polylineDB.Stroke_Data.String = polyline.Stroke
	polylineDB.Stroke_Data.Valid = true

	polylineDB.StrokeOpacity_Data.Float64 = polyline.StrokeOpacity
	polylineDB.StrokeOpacity_Data.Valid = true

	polylineDB.StrokeWidth_Data.Float64 = polyline.StrokeWidth
	polylineDB.StrokeWidth_Data.Valid = true

	polylineDB.StrokeDashArray_Data.String = polyline.StrokeDashArray
	polylineDB.StrokeDashArray_Data.Valid = true

	polylineDB.StrokeDashArrayWhenSelected_Data.String = polyline.StrokeDashArrayWhenSelected
	polylineDB.StrokeDashArrayWhenSelected_Data.Valid = true

	polylineDB.Transform_Data.String = polyline.Transform
	polylineDB.Transform_Data.Valid = true
}

// CopyBasicFieldsFromPolyline_WOP
func (polylineDB *PolylineDB) CopyBasicFieldsFromPolyline_WOP(polyline *models.Polyline_WOP) {
	// insertion point for fields commit

	polylineDB.Name_Data.String = polyline.Name
	polylineDB.Name_Data.Valid = true

	polylineDB.Points_Data.String = polyline.Points
	polylineDB.Points_Data.Valid = true

	polylineDB.Color_Data.String = polyline.Color
	polylineDB.Color_Data.Valid = true

	polylineDB.FillOpacity_Data.Float64 = polyline.FillOpacity
	polylineDB.FillOpacity_Data.Valid = true

	polylineDB.Stroke_Data.String = polyline.Stroke
	polylineDB.Stroke_Data.Valid = true

	polylineDB.StrokeOpacity_Data.Float64 = polyline.StrokeOpacity
	polylineDB.StrokeOpacity_Data.Valid = true

	polylineDB.StrokeWidth_Data.Float64 = polyline.StrokeWidth
	polylineDB.StrokeWidth_Data.Valid = true

	polylineDB.StrokeDashArray_Data.String = polyline.StrokeDashArray
	polylineDB.StrokeDashArray_Data.Valid = true

	polylineDB.StrokeDashArrayWhenSelected_Data.String = polyline.StrokeDashArrayWhenSelected
	polylineDB.StrokeDashArrayWhenSelected_Data.Valid = true

	polylineDB.Transform_Data.String = polyline.Transform
	polylineDB.Transform_Data.Valid = true
}

// CopyBasicFieldsFromPolylineWOP
func (polylineDB *PolylineDB) CopyBasicFieldsFromPolylineWOP(polyline *PolylineWOP) {
	// insertion point for fields commit

	polylineDB.Name_Data.String = polyline.Name
	polylineDB.Name_Data.Valid = true

	polylineDB.Points_Data.String = polyline.Points
	polylineDB.Points_Data.Valid = true

	polylineDB.Color_Data.String = polyline.Color
	polylineDB.Color_Data.Valid = true

	polylineDB.FillOpacity_Data.Float64 = polyline.FillOpacity
	polylineDB.FillOpacity_Data.Valid = true

	polylineDB.Stroke_Data.String = polyline.Stroke
	polylineDB.Stroke_Data.Valid = true

	polylineDB.StrokeOpacity_Data.Float64 = polyline.StrokeOpacity
	polylineDB.StrokeOpacity_Data.Valid = true

	polylineDB.StrokeWidth_Data.Float64 = polyline.StrokeWidth
	polylineDB.StrokeWidth_Data.Valid = true

	polylineDB.StrokeDashArray_Data.String = polyline.StrokeDashArray
	polylineDB.StrokeDashArray_Data.Valid = true

	polylineDB.StrokeDashArrayWhenSelected_Data.String = polyline.StrokeDashArrayWhenSelected
	polylineDB.StrokeDashArrayWhenSelected_Data.Valid = true

	polylineDB.Transform_Data.String = polyline.Transform
	polylineDB.Transform_Data.Valid = true
}

// CopyBasicFieldsToPolyline
func (polylineDB *PolylineDB) CopyBasicFieldsToPolyline(polyline *models.Polyline) {
	// insertion point for checkout of basic fields (back repo to stage)
	polyline.Name = polylineDB.Name_Data.String
	polyline.Points = polylineDB.Points_Data.String
	polyline.Color = polylineDB.Color_Data.String
	polyline.FillOpacity = polylineDB.FillOpacity_Data.Float64
	polyline.Stroke = polylineDB.Stroke_Data.String
	polyline.StrokeOpacity = polylineDB.StrokeOpacity_Data.Float64
	polyline.StrokeWidth = polylineDB.StrokeWidth_Data.Float64
	polyline.StrokeDashArray = polylineDB.StrokeDashArray_Data.String
	polyline.StrokeDashArrayWhenSelected = polylineDB.StrokeDashArrayWhenSelected_Data.String
	polyline.Transform = polylineDB.Transform_Data.String
}

// CopyBasicFieldsToPolyline_WOP
func (polylineDB *PolylineDB) CopyBasicFieldsToPolyline_WOP(polyline *models.Polyline_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	polyline.Name = polylineDB.Name_Data.String
	polyline.Points = polylineDB.Points_Data.String
	polyline.Color = polylineDB.Color_Data.String
	polyline.FillOpacity = polylineDB.FillOpacity_Data.Float64
	polyline.Stroke = polylineDB.Stroke_Data.String
	polyline.StrokeOpacity = polylineDB.StrokeOpacity_Data.Float64
	polyline.StrokeWidth = polylineDB.StrokeWidth_Data.Float64
	polyline.StrokeDashArray = polylineDB.StrokeDashArray_Data.String
	polyline.StrokeDashArrayWhenSelected = polylineDB.StrokeDashArrayWhenSelected_Data.String
	polyline.Transform = polylineDB.Transform_Data.String
}

// CopyBasicFieldsToPolylineWOP
func (polylineDB *PolylineDB) CopyBasicFieldsToPolylineWOP(polyline *PolylineWOP) {
	polyline.ID = int(polylineDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	polyline.Name = polylineDB.Name_Data.String
	polyline.Points = polylineDB.Points_Data.String
	polyline.Color = polylineDB.Color_Data.String
	polyline.FillOpacity = polylineDB.FillOpacity_Data.Float64
	polyline.Stroke = polylineDB.Stroke_Data.String
	polyline.StrokeOpacity = polylineDB.StrokeOpacity_Data.Float64
	polyline.StrokeWidth = polylineDB.StrokeWidth_Data.Float64
	polyline.StrokeDashArray = polylineDB.StrokeDashArray_Data.String
	polyline.StrokeDashArrayWhenSelected = polylineDB.StrokeDashArrayWhenSelected_Data.String
	polyline.Transform = polylineDB.Transform_Data.String
}

// Backup generates a json file from a slice of all PolylineDB instances in the backrepo
func (backRepoPolyline *BackRepoPolylineStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "PolylineDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*PolylineDB, 0)
	for _, polylineDB := range backRepoPolyline.Map_PolylineDBID_PolylineDB {
		forBackup = append(forBackup, polylineDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Polyline ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Polyline file", err.Error())
	}
}

// Backup generates a json file from a slice of all PolylineDB instances in the backrepo
func (backRepoPolyline *BackRepoPolylineStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*PolylineDB, 0)
	for _, polylineDB := range backRepoPolyline.Map_PolylineDBID_PolylineDB {
		forBackup = append(forBackup, polylineDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Polyline")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Polyline_Fields, -1)
	for _, polylineDB := range forBackup {

		var polylineWOP PolylineWOP
		polylineDB.CopyBasicFieldsToPolylineWOP(&polylineWOP)

		row := sh.AddRow()
		row.WriteStruct(&polylineWOP, -1)
	}
}

// RestoreXL from the "Polyline" sheet all PolylineDB instances
func (backRepoPolyline *BackRepoPolylineStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoPolylineid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Polyline"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoPolyline.rowVisitorPolyline)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoPolyline *BackRepoPolylineStruct) rowVisitorPolyline(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var polylineWOP PolylineWOP
		row.ReadStruct(&polylineWOP)

		// add the unmarshalled struct to the stage
		polylineDB := new(PolylineDB)
		polylineDB.CopyBasicFieldsFromPolylineWOP(&polylineWOP)

		polylineDB_ID_atBackupTime := polylineDB.ID
		polylineDB.ID = 0
		_, err := backRepoPolyline.db.Create(polylineDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoPolyline.Map_PolylineDBID_PolylineDB[polylineDB.ID] = polylineDB
		BackRepoPolylineid_atBckpTime_newID[polylineDB_ID_atBackupTime] = polylineDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "PolylineDB.json" in dirPath that stores an array
// of PolylineDB and stores it in the database
// the map BackRepoPolylineid_atBckpTime_newID is updated accordingly
func (backRepoPolyline *BackRepoPolylineStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoPolylineid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "PolylineDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Polyline file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*PolylineDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_PolylineDBID_PolylineDB
	for _, polylineDB := range forRestore {

		polylineDB_ID_atBackupTime := polylineDB.ID
		polylineDB.ID = 0
		_, err := backRepoPolyline.db.Create(polylineDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoPolyline.Map_PolylineDBID_PolylineDB[polylineDB.ID] = polylineDB
		BackRepoPolylineid_atBckpTime_newID[polylineDB_ID_atBackupTime] = polylineDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Polyline file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Polyline>id_atBckpTime_newID
// to compute new index
func (backRepoPolyline *BackRepoPolylineStruct) RestorePhaseTwo() {

	for _, polylineDB := range backRepoPolyline.Map_PolylineDBID_PolylineDB {

		// next line of code is to avert unused variable compilation error
		_ = polylineDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoPolyline.db.Model(polylineDB)
		_, err := db.Updates(*polylineDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoPolyline.ResetReversePointers commits all staged instances of Polyline to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPolyline *BackRepoPolylineStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, polyline := range backRepoPolyline.Map_PolylineDBID_PolylinePtr {
		backRepoPolyline.ResetReversePointersInstance(backRepo, idx, polyline)
	}

	return
}

func (backRepoPolyline *BackRepoPolylineStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, polyline *models.Polyline) (Error error) {

	// fetch matching polylineDB
	if polylineDB, ok := backRepoPolyline.Map_PolylineDBID_PolylineDB[idx]; ok {
		_ = polylineDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoPolylineid_atBckpTime_newID map[uint]uint
