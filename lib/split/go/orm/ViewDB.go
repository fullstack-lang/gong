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

	"github.com/fullstack-lang/gong/lib/split/go/db"
	"github.com/fullstack-lang/gong/lib/split/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_View_sql sql.NullBool
var dummy_View_time time.Duration
var dummy_View_sort sort.Float64Slice

// ViewAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model viewAPI
type ViewAPI struct {
	gorm.Model

	models.View_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	ViewPointersEncoding ViewPointersEncoding
}

// ViewPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type ViewPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field RootAsSplitAreas is a slice of pointers to another Struct (optional or 0..1)
	RootAsSplitAreas IntSlice `gorm:"type:TEXT"`
}

// ViewDB describes a view in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model viewDB
type ViewDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field viewDB.Name
	Name_Data sql.NullString

	// Declation for basic field viewDB.ShowViewName
	// provide the sql storage for the boolan
	ShowViewName_Data sql.NullBool

	// Declation for basic field viewDB.IsSelectedView
	// provide the sql storage for the boolan
	IsSelectedView_Data sql.NullBool

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	ViewPointersEncoding
}

// ViewDBs arrays viewDBs
// swagger:response viewDBsResponse
type ViewDBs []ViewDB

// ViewDBResponse provides response
// swagger:response viewDBResponse
type ViewDBResponse struct {
	ViewDB
}

// ViewWOP is a View without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type ViewWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	ShowViewName bool `xlsx:"2"`

	IsSelectedView bool `xlsx:"3"`
	// insertion for WOP pointer fields
}

var View_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"ShowViewName",
	"IsSelectedView",
}

type BackRepoViewStruct struct {
	// stores ViewDB according to their gorm ID
	Map_ViewDBID_ViewDB map[uint]*ViewDB

	// stores ViewDB ID according to View address
	Map_ViewPtr_ViewDBID map[*models.View]uint

	// stores View according to their gorm ID
	Map_ViewDBID_ViewPtr map[uint]*models.View

	db db.DBInterface

	stage *models.Stage
}

func (backRepoView *BackRepoViewStruct) GetStage() (stage *models.Stage) {
	stage = backRepoView.stage
	return
}

func (backRepoView *BackRepoViewStruct) GetDB() db.DBInterface {
	return backRepoView.db
}

// GetViewDBFromViewPtr is a handy function to access the back repo instance from the stage instance
func (backRepoView *BackRepoViewStruct) GetViewDBFromViewPtr(view *models.View) (viewDB *ViewDB) {
	id := backRepoView.Map_ViewPtr_ViewDBID[view]
	viewDB = backRepoView.Map_ViewDBID_ViewDB[id]
	return
}

// BackRepoView.CommitPhaseOne commits all staged instances of View to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoView *BackRepoViewStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var views []*models.View
	for view := range stage.Views {
		views = append(views, view)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(views, func(i, j int) bool {
		return stage.ViewMap_Staged_Order[views[i]] < stage.ViewMap_Staged_Order[views[j]]
	})

	for _, view := range views {
		backRepoView.CommitPhaseOneInstance(view)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, view := range backRepoView.Map_ViewDBID_ViewPtr {
		if _, ok := stage.Views[view]; !ok {
			backRepoView.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoView.CommitDeleteInstance commits deletion of View to the BackRepo
func (backRepoView *BackRepoViewStruct) CommitDeleteInstance(id uint) (Error error) {

	view := backRepoView.Map_ViewDBID_ViewPtr[id]

	// view is not staged anymore, remove viewDB
	viewDB := backRepoView.Map_ViewDBID_ViewDB[id]
	db, _ := backRepoView.db.Unscoped()
	_, err := db.Delete(viewDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoView.Map_ViewPtr_ViewDBID, view)
	delete(backRepoView.Map_ViewDBID_ViewPtr, id)
	delete(backRepoView.Map_ViewDBID_ViewDB, id)

	return
}

// BackRepoView.CommitPhaseOneInstance commits view staged instances of View to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoView *BackRepoViewStruct) CommitPhaseOneInstance(view *models.View) (Error error) {

	// check if the view is not commited yet
	if _, ok := backRepoView.Map_ViewPtr_ViewDBID[view]; ok {
		return
	}

	// initiate view
	var viewDB ViewDB
	viewDB.CopyBasicFieldsFromView(view)

	_, err := backRepoView.db.Create(&viewDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoView.Map_ViewPtr_ViewDBID[view] = viewDB.ID
	backRepoView.Map_ViewDBID_ViewPtr[viewDB.ID] = view
	backRepoView.Map_ViewDBID_ViewDB[viewDB.ID] = &viewDB

	return
}

// BackRepoView.CommitPhaseTwo commits all staged instances of View to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoView *BackRepoViewStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, view := range backRepoView.Map_ViewDBID_ViewPtr {
		backRepoView.CommitPhaseTwoInstance(backRepo, idx, view)
	}

	return
}

// BackRepoView.CommitPhaseTwoInstance commits {{structname }} of models.View to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoView *BackRepoViewStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, view *models.View) (Error error) {

	// fetch matching viewDB
	if viewDB, ok := backRepoView.Map_ViewDBID_ViewDB[idx]; ok {

		viewDB.CopyBasicFieldsFromView(view)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		viewDB.ViewPointersEncoding.RootAsSplitAreas = make([]int, 0)
		// 2. encode
		for _, assplitareaAssocEnd := range view.RootAsSplitAreas {
			assplitareaAssocEnd_DB :=
				backRepo.BackRepoAsSplitArea.GetAsSplitAreaDBFromAsSplitAreaPtr(assplitareaAssocEnd)
			
			// the stage might be inconsistant, meaning that the assplitareaAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if assplitareaAssocEnd_DB == nil {
				continue
			}
			
			viewDB.ViewPointersEncoding.RootAsSplitAreas =
				append(viewDB.ViewPointersEncoding.RootAsSplitAreas, int(assplitareaAssocEnd_DB.ID))
		}

		_, err := backRepoView.db.Save(viewDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown View intance %s", view.Name))
		return err
	}

	return
}

// BackRepoView.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoView *BackRepoViewStruct) CheckoutPhaseOne() (Error error) {

	viewDBArray := make([]ViewDB, 0)
	_, err := backRepoView.db.Find(&viewDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	viewInstancesToBeRemovedFromTheStage := make(map[*models.View]any)
	for key, value := range backRepoView.stage.Views {
		viewInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, viewDB := range viewDBArray {
		backRepoView.CheckoutPhaseOneInstance(&viewDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		view, ok := backRepoView.Map_ViewDBID_ViewPtr[viewDB.ID]
		if ok {
			delete(viewInstancesToBeRemovedFromTheStage, view)
		}
	}

	// remove from stage and back repo's 3 maps all views that are not in the checkout
	for view := range viewInstancesToBeRemovedFromTheStage {
		view.Unstage(backRepoView.GetStage())

		// remove instance from the back repo 3 maps
		viewID := backRepoView.Map_ViewPtr_ViewDBID[view]
		delete(backRepoView.Map_ViewPtr_ViewDBID, view)
		delete(backRepoView.Map_ViewDBID_ViewDB, viewID)
		delete(backRepoView.Map_ViewDBID_ViewPtr, viewID)
	}

	return
}

// CheckoutPhaseOneInstance takes a viewDB that has been found in the DB, updates the backRepo and stages the
// models version of the viewDB
func (backRepoView *BackRepoViewStruct) CheckoutPhaseOneInstance(viewDB *ViewDB) (Error error) {

	view, ok := backRepoView.Map_ViewDBID_ViewPtr[viewDB.ID]
	if !ok {
		view = new(models.View)

		backRepoView.Map_ViewDBID_ViewPtr[viewDB.ID] = view
		backRepoView.Map_ViewPtr_ViewDBID[view] = viewDB.ID

		// append model store with the new element
		view.Name = viewDB.Name_Data.String
		view.Stage(backRepoView.GetStage())
	}
	viewDB.CopyBasicFieldsToView(view)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	view.Stage(backRepoView.GetStage())

	// preserve pointer to viewDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_ViewDBID_ViewDB)[viewDB hold variable pointers
	viewDB_Data := *viewDB
	preservedPtrToView := &viewDB_Data
	backRepoView.Map_ViewDBID_ViewDB[viewDB.ID] = preservedPtrToView

	return
}

// BackRepoView.CheckoutPhaseTwo Checkouts all staged instances of View to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoView *BackRepoViewStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, viewDB := range backRepoView.Map_ViewDBID_ViewDB {
		backRepoView.CheckoutPhaseTwoInstance(backRepo, viewDB)
	}
	return
}

// BackRepoView.CheckoutPhaseTwoInstance Checkouts staged instances of View to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoView *BackRepoViewStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, viewDB *ViewDB) (Error error) {

	view := backRepoView.Map_ViewDBID_ViewPtr[viewDB.ID]

	viewDB.DecodePointers(backRepo, view)

	return
}

func (viewDB *ViewDB) DecodePointers(backRepo *BackRepoStruct, view *models.View) {

	// insertion point for checkout of pointer encoding
	// This loop redeem view.RootAsSplitAreas in the stage from the encode in the back repo
	// It parses all AsSplitAreaDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	view.RootAsSplitAreas = view.RootAsSplitAreas[:0]
	for _, _AsSplitAreaid := range viewDB.ViewPointersEncoding.RootAsSplitAreas {
		view.RootAsSplitAreas = append(view.RootAsSplitAreas, backRepo.BackRepoAsSplitArea.Map_AsSplitAreaDBID_AsSplitAreaPtr[uint(_AsSplitAreaid)])
	}

	return
}

// CommitView allows commit of a single view (if already staged)
func (backRepo *BackRepoStruct) CommitView(view *models.View) {
	backRepo.BackRepoView.CommitPhaseOneInstance(view)
	if id, ok := backRepo.BackRepoView.Map_ViewPtr_ViewDBID[view]; ok {
		backRepo.BackRepoView.CommitPhaseTwoInstance(backRepo, id, view)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitView allows checkout of a single view (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutView(view *models.View) {
	// check if the view is staged
	if _, ok := backRepo.BackRepoView.Map_ViewPtr_ViewDBID[view]; ok {

		if id, ok := backRepo.BackRepoView.Map_ViewPtr_ViewDBID[view]; ok {
			var viewDB ViewDB
			viewDB.ID = id

			if _, err := backRepo.BackRepoView.db.First(&viewDB, id); err != nil {
				log.Fatalln("CheckoutView : Problem with getting object with id:", id)
			}
			backRepo.BackRepoView.CheckoutPhaseOneInstance(&viewDB)
			backRepo.BackRepoView.CheckoutPhaseTwoInstance(backRepo, &viewDB)
		}
	}
}

// CopyBasicFieldsFromView
func (viewDB *ViewDB) CopyBasicFieldsFromView(view *models.View) {
	// insertion point for fields commit

	viewDB.Name_Data.String = view.Name
	viewDB.Name_Data.Valid = true

	viewDB.ShowViewName_Data.Bool = view.ShowViewName
	viewDB.ShowViewName_Data.Valid = true

	viewDB.IsSelectedView_Data.Bool = view.IsSelectedView
	viewDB.IsSelectedView_Data.Valid = true
}

// CopyBasicFieldsFromView_WOP
func (viewDB *ViewDB) CopyBasicFieldsFromView_WOP(view *models.View_WOP) {
	// insertion point for fields commit

	viewDB.Name_Data.String = view.Name
	viewDB.Name_Data.Valid = true

	viewDB.ShowViewName_Data.Bool = view.ShowViewName
	viewDB.ShowViewName_Data.Valid = true

	viewDB.IsSelectedView_Data.Bool = view.IsSelectedView
	viewDB.IsSelectedView_Data.Valid = true
}

// CopyBasicFieldsFromViewWOP
func (viewDB *ViewDB) CopyBasicFieldsFromViewWOP(view *ViewWOP) {
	// insertion point for fields commit

	viewDB.Name_Data.String = view.Name
	viewDB.Name_Data.Valid = true

	viewDB.ShowViewName_Data.Bool = view.ShowViewName
	viewDB.ShowViewName_Data.Valid = true

	viewDB.IsSelectedView_Data.Bool = view.IsSelectedView
	viewDB.IsSelectedView_Data.Valid = true
}

// CopyBasicFieldsToView
func (viewDB *ViewDB) CopyBasicFieldsToView(view *models.View) {
	// insertion point for checkout of basic fields (back repo to stage)
	view.Name = viewDB.Name_Data.String
	view.ShowViewName = viewDB.ShowViewName_Data.Bool
	view.IsSelectedView = viewDB.IsSelectedView_Data.Bool
}

// CopyBasicFieldsToView_WOP
func (viewDB *ViewDB) CopyBasicFieldsToView_WOP(view *models.View_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	view.Name = viewDB.Name_Data.String
	view.ShowViewName = viewDB.ShowViewName_Data.Bool
	view.IsSelectedView = viewDB.IsSelectedView_Data.Bool
}

// CopyBasicFieldsToViewWOP
func (viewDB *ViewDB) CopyBasicFieldsToViewWOP(view *ViewWOP) {
	view.ID = int(viewDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	view.Name = viewDB.Name_Data.String
	view.ShowViewName = viewDB.ShowViewName_Data.Bool
	view.IsSelectedView = viewDB.IsSelectedView_Data.Bool
}

// Backup generates a json file from a slice of all ViewDB instances in the backrepo
func (backRepoView *BackRepoViewStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "ViewDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ViewDB, 0)
	for _, viewDB := range backRepoView.Map_ViewDBID_ViewDB {
		forBackup = append(forBackup, viewDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json View ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json View file", err.Error())
	}
}

// Backup generates a json file from a slice of all ViewDB instances in the backrepo
func (backRepoView *BackRepoViewStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ViewDB, 0)
	for _, viewDB := range backRepoView.Map_ViewDBID_ViewDB {
		forBackup = append(forBackup, viewDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("View")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&View_Fields, -1)
	for _, viewDB := range forBackup {

		var viewWOP ViewWOP
		viewDB.CopyBasicFieldsToViewWOP(&viewWOP)

		row := sh.AddRow()
		row.WriteStruct(&viewWOP, -1)
	}
}

// RestoreXL from the "View" sheet all ViewDB instances
func (backRepoView *BackRepoViewStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoViewid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["View"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoView.rowVisitorView)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoView *BackRepoViewStruct) rowVisitorView(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var viewWOP ViewWOP
		row.ReadStruct(&viewWOP)

		// add the unmarshalled struct to the stage
		viewDB := new(ViewDB)
		viewDB.CopyBasicFieldsFromViewWOP(&viewWOP)

		viewDB_ID_atBackupTime := viewDB.ID
		viewDB.ID = 0
		_, err := backRepoView.db.Create(viewDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoView.Map_ViewDBID_ViewDB[viewDB.ID] = viewDB
		BackRepoViewid_atBckpTime_newID[viewDB_ID_atBackupTime] = viewDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "ViewDB.json" in dirPath that stores an array
// of ViewDB and stores it in the database
// the map BackRepoViewid_atBckpTime_newID is updated accordingly
func (backRepoView *BackRepoViewStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoViewid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "ViewDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json View file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*ViewDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_ViewDBID_ViewDB
	for _, viewDB := range forRestore {

		viewDB_ID_atBackupTime := viewDB.ID
		viewDB.ID = 0
		_, err := backRepoView.db.Create(viewDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoView.Map_ViewDBID_ViewDB[viewDB.ID] = viewDB
		BackRepoViewid_atBckpTime_newID[viewDB_ID_atBackupTime] = viewDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json View file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<View>id_atBckpTime_newID
// to compute new index
func (backRepoView *BackRepoViewStruct) RestorePhaseTwo() {

	for _, viewDB := range backRepoView.Map_ViewDBID_ViewDB {

		// next line of code is to avert unused variable compilation error
		_ = viewDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoView.db.Model(viewDB)
		_, err := db.Updates(*viewDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoView.ResetReversePointers commits all staged instances of View to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoView *BackRepoViewStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, view := range backRepoView.Map_ViewDBID_ViewPtr {
		backRepoView.ResetReversePointersInstance(backRepo, idx, view)
	}

	return
}

func (backRepoView *BackRepoViewStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, view *models.View) (Error error) {

	// fetch matching viewDB
	if viewDB, ok := backRepoView.Map_ViewDBID_ViewDB[idx]; ok {
		_ = viewDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoViewid_atBckpTime_newID map[uint]uint
