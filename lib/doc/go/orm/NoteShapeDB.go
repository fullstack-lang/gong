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

	"github.com/fullstack-lang/gong/lib/doc/go/db"
	"github.com/fullstack-lang/gong/lib/doc/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_NoteShape_sql sql.NullBool
var dummy_NoteShape_time time.Duration
var dummy_NoteShape_sort sort.Float64Slice

// NoteShapeAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model noteshapeAPI
type NoteShapeAPI struct {
	gorm.Model

	models.NoteShape_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	NoteShapePointersEncoding NoteShapePointersEncoding
}

// NoteShapePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type NoteShapePointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field NoteShapeLinks is a slice of pointers to another Struct (optional or 0..1)
	NoteShapeLinks IntSlice `gorm:"type:TEXT"`
}

// NoteShapeDB describes a noteshape in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model noteshapeDB
type NoteShapeDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field noteshapeDB.Name
	Name_Data sql.NullString

	// Declation for basic field noteshapeDB.Identifier
	Identifier_Data sql.NullString

	// Declation for basic field noteshapeDB.Body
	Body_Data sql.NullString

	// Declation for basic field noteshapeDB.BodyHTML
	BodyHTML_Data sql.NullString

	// Declation for basic field noteshapeDB.X
	X_Data sql.NullFloat64

	// Declation for basic field noteshapeDB.Y
	Y_Data sql.NullFloat64

	// Declation for basic field noteshapeDB.Width
	Width_Data sql.NullFloat64

	// Declation for basic field noteshapeDB.Height
	Height_Data sql.NullFloat64

	// Declation for basic field noteshapeDB.Matched
	// provide the sql storage for the boolan
	Matched_Data sql.NullBool

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	NoteShapePointersEncoding
}

// NoteShapeDBs arrays noteshapeDBs
// swagger:response noteshapeDBsResponse
type NoteShapeDBs []NoteShapeDB

// NoteShapeDBResponse provides response
// swagger:response noteshapeDBResponse
type NoteShapeDBResponse struct {
	NoteShapeDB
}

// NoteShapeWOP is a NoteShape without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type NoteShapeWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Identifier string `xlsx:"2"`

	Body string `xlsx:"3"`

	BodyHTML string `xlsx:"4"`

	X float64 `xlsx:"5"`

	Y float64 `xlsx:"6"`

	Width float64 `xlsx:"7"`

	Height float64 `xlsx:"8"`

	Matched bool `xlsx:"9"`
	// insertion for WOP pointer fields
}

var NoteShape_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Identifier",
	"Body",
	"BodyHTML",
	"X",
	"Y",
	"Width",
	"Height",
	"Matched",
}

type BackRepoNoteShapeStruct struct {
	// stores NoteShapeDB according to their gorm ID
	Map_NoteShapeDBID_NoteShapeDB map[uint]*NoteShapeDB

	// stores NoteShapeDB ID according to NoteShape address
	Map_NoteShapePtr_NoteShapeDBID map[*models.NoteShape]uint

	// stores NoteShape according to their gorm ID
	Map_NoteShapeDBID_NoteShapePtr map[uint]*models.NoteShape

	db db.DBInterface

	stage *models.Stage
}

func (backRepoNoteShape *BackRepoNoteShapeStruct) GetStage() (stage *models.Stage) {
	stage = backRepoNoteShape.stage
	return
}

func (backRepoNoteShape *BackRepoNoteShapeStruct) GetDB() db.DBInterface {
	return backRepoNoteShape.db
}

// GetNoteShapeDBFromNoteShapePtr is a handy function to access the back repo instance from the stage instance
func (backRepoNoteShape *BackRepoNoteShapeStruct) GetNoteShapeDBFromNoteShapePtr(noteshape *models.NoteShape) (noteshapeDB *NoteShapeDB) {
	id := backRepoNoteShape.Map_NoteShapePtr_NoteShapeDBID[noteshape]
	noteshapeDB = backRepoNoteShape.Map_NoteShapeDBID_NoteShapeDB[id]
	return
}

// BackRepoNoteShape.CommitPhaseOne commits all staged instances of NoteShape to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoNoteShape *BackRepoNoteShapeStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var noteshapes []*models.NoteShape
	for noteshape := range stage.NoteShapes {
		noteshapes = append(noteshapes, noteshape)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(noteshapes, func(i, j int) bool {
		return stage.NoteShapeMap_Staged_Order[noteshapes[i]] < stage.NoteShapeMap_Staged_Order[noteshapes[j]]
	})

	for _, noteshape := range noteshapes {
		backRepoNoteShape.CommitPhaseOneInstance(noteshape)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, noteshape := range backRepoNoteShape.Map_NoteShapeDBID_NoteShapePtr {
		if _, ok := stage.NoteShapes[noteshape]; !ok {
			backRepoNoteShape.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoNoteShape.CommitDeleteInstance commits deletion of NoteShape to the BackRepo
func (backRepoNoteShape *BackRepoNoteShapeStruct) CommitDeleteInstance(id uint) (Error error) {

	noteshape := backRepoNoteShape.Map_NoteShapeDBID_NoteShapePtr[id]

	// noteshape is not staged anymore, remove noteshapeDB
	noteshapeDB := backRepoNoteShape.Map_NoteShapeDBID_NoteShapeDB[id]
	db, _ := backRepoNoteShape.db.Unscoped()
	_, err := db.Delete(noteshapeDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoNoteShape.Map_NoteShapePtr_NoteShapeDBID, noteshape)
	delete(backRepoNoteShape.Map_NoteShapeDBID_NoteShapePtr, id)
	delete(backRepoNoteShape.Map_NoteShapeDBID_NoteShapeDB, id)

	return
}

// BackRepoNoteShape.CommitPhaseOneInstance commits noteshape staged instances of NoteShape to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoNoteShape *BackRepoNoteShapeStruct) CommitPhaseOneInstance(noteshape *models.NoteShape) (Error error) {

	// check if the noteshape is not commited yet
	if _, ok := backRepoNoteShape.Map_NoteShapePtr_NoteShapeDBID[noteshape]; ok {
		return
	}

	// initiate noteshape
	var noteshapeDB NoteShapeDB
	noteshapeDB.CopyBasicFieldsFromNoteShape(noteshape)

	_, err := backRepoNoteShape.db.Create(&noteshapeDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoNoteShape.Map_NoteShapePtr_NoteShapeDBID[noteshape] = noteshapeDB.ID
	backRepoNoteShape.Map_NoteShapeDBID_NoteShapePtr[noteshapeDB.ID] = noteshape
	backRepoNoteShape.Map_NoteShapeDBID_NoteShapeDB[noteshapeDB.ID] = &noteshapeDB

	return
}

// BackRepoNoteShape.CommitPhaseTwo commits all staged instances of NoteShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNoteShape *BackRepoNoteShapeStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, noteshape := range backRepoNoteShape.Map_NoteShapeDBID_NoteShapePtr {
		backRepoNoteShape.CommitPhaseTwoInstance(backRepo, idx, noteshape)
	}

	return
}

// BackRepoNoteShape.CommitPhaseTwoInstance commits {{structname }} of models.NoteShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNoteShape *BackRepoNoteShapeStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, noteshape *models.NoteShape) (Error error) {

	// fetch matching noteshapeDB
	if noteshapeDB, ok := backRepoNoteShape.Map_NoteShapeDBID_NoteShapeDB[idx]; ok {

		noteshapeDB.CopyBasicFieldsFromNoteShape(noteshape)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		noteshapeDB.NoteShapePointersEncoding.NoteShapeLinks = make([]int, 0)
		// 2. encode
		for _, noteshapelinkAssocEnd := range noteshape.NoteShapeLinks {
			noteshapelinkAssocEnd_DB :=
				backRepo.BackRepoNoteShapeLink.GetNoteShapeLinkDBFromNoteShapeLinkPtr(noteshapelinkAssocEnd)
			
			// the stage might be inconsistant, meaning that the noteshapelinkAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if noteshapelinkAssocEnd_DB == nil {
				continue
			}
			
			noteshapeDB.NoteShapePointersEncoding.NoteShapeLinks =
				append(noteshapeDB.NoteShapePointersEncoding.NoteShapeLinks, int(noteshapelinkAssocEnd_DB.ID))
		}

		_, err := backRepoNoteShape.db.Save(noteshapeDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown NoteShape intance %s", noteshape.Name))
		return err
	}

	return
}

// BackRepoNoteShape.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoNoteShape *BackRepoNoteShapeStruct) CheckoutPhaseOne() (Error error) {

	noteshapeDBArray := make([]NoteShapeDB, 0)
	_, err := backRepoNoteShape.db.Find(&noteshapeDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	noteshapeInstancesToBeRemovedFromTheStage := make(map[*models.NoteShape]any)
	for key, value := range backRepoNoteShape.stage.NoteShapes {
		noteshapeInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, noteshapeDB := range noteshapeDBArray {
		backRepoNoteShape.CheckoutPhaseOneInstance(&noteshapeDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		noteshape, ok := backRepoNoteShape.Map_NoteShapeDBID_NoteShapePtr[noteshapeDB.ID]
		if ok {
			delete(noteshapeInstancesToBeRemovedFromTheStage, noteshape)
		}
	}

	// remove from stage and back repo's 3 maps all noteshapes that are not in the checkout
	for noteshape := range noteshapeInstancesToBeRemovedFromTheStage {
		noteshape.Unstage(backRepoNoteShape.GetStage())

		// remove instance from the back repo 3 maps
		noteshapeID := backRepoNoteShape.Map_NoteShapePtr_NoteShapeDBID[noteshape]
		delete(backRepoNoteShape.Map_NoteShapePtr_NoteShapeDBID, noteshape)
		delete(backRepoNoteShape.Map_NoteShapeDBID_NoteShapeDB, noteshapeID)
		delete(backRepoNoteShape.Map_NoteShapeDBID_NoteShapePtr, noteshapeID)
	}

	return
}

// CheckoutPhaseOneInstance takes a noteshapeDB that has been found in the DB, updates the backRepo and stages the
// models version of the noteshapeDB
func (backRepoNoteShape *BackRepoNoteShapeStruct) CheckoutPhaseOneInstance(noteshapeDB *NoteShapeDB) (Error error) {

	noteshape, ok := backRepoNoteShape.Map_NoteShapeDBID_NoteShapePtr[noteshapeDB.ID]
	if !ok {
		noteshape = new(models.NoteShape)

		backRepoNoteShape.Map_NoteShapeDBID_NoteShapePtr[noteshapeDB.ID] = noteshape
		backRepoNoteShape.Map_NoteShapePtr_NoteShapeDBID[noteshape] = noteshapeDB.ID

		// append model store with the new element
		noteshape.Name = noteshapeDB.Name_Data.String
		noteshape.Stage(backRepoNoteShape.GetStage())
	}
	noteshapeDB.CopyBasicFieldsToNoteShape(noteshape)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	noteshape.Stage(backRepoNoteShape.GetStage())

	// preserve pointer to noteshapeDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_NoteShapeDBID_NoteShapeDB)[noteshapeDB hold variable pointers
	noteshapeDB_Data := *noteshapeDB
	preservedPtrToNoteShape := &noteshapeDB_Data
	backRepoNoteShape.Map_NoteShapeDBID_NoteShapeDB[noteshapeDB.ID] = preservedPtrToNoteShape

	return
}

// BackRepoNoteShape.CheckoutPhaseTwo Checkouts all staged instances of NoteShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNoteShape *BackRepoNoteShapeStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, noteshapeDB := range backRepoNoteShape.Map_NoteShapeDBID_NoteShapeDB {
		backRepoNoteShape.CheckoutPhaseTwoInstance(backRepo, noteshapeDB)
	}
	return
}

// BackRepoNoteShape.CheckoutPhaseTwoInstance Checkouts staged instances of NoteShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNoteShape *BackRepoNoteShapeStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, noteshapeDB *NoteShapeDB) (Error error) {

	noteshape := backRepoNoteShape.Map_NoteShapeDBID_NoteShapePtr[noteshapeDB.ID]

	noteshapeDB.DecodePointers(backRepo, noteshape)

	return
}

func (noteshapeDB *NoteShapeDB) DecodePointers(backRepo *BackRepoStruct, noteshape *models.NoteShape) {

	// insertion point for checkout of pointer encoding
	// This loop redeem noteshape.NoteShapeLinks in the stage from the encode in the back repo
	// It parses all NoteShapeLinkDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	noteshape.NoteShapeLinks = noteshape.NoteShapeLinks[:0]
	for _, _NoteShapeLinkid := range noteshapeDB.NoteShapePointersEncoding.NoteShapeLinks {
		noteshape.NoteShapeLinks = append(noteshape.NoteShapeLinks, backRepo.BackRepoNoteShapeLink.Map_NoteShapeLinkDBID_NoteShapeLinkPtr[uint(_NoteShapeLinkid)])
	}

	return
}

// CommitNoteShape allows commit of a single noteshape (if already staged)
func (backRepo *BackRepoStruct) CommitNoteShape(noteshape *models.NoteShape) {
	backRepo.BackRepoNoteShape.CommitPhaseOneInstance(noteshape)
	if id, ok := backRepo.BackRepoNoteShape.Map_NoteShapePtr_NoteShapeDBID[noteshape]; ok {
		backRepo.BackRepoNoteShape.CommitPhaseTwoInstance(backRepo, id, noteshape)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitNoteShape allows checkout of a single noteshape (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutNoteShape(noteshape *models.NoteShape) {
	// check if the noteshape is staged
	if _, ok := backRepo.BackRepoNoteShape.Map_NoteShapePtr_NoteShapeDBID[noteshape]; ok {

		if id, ok := backRepo.BackRepoNoteShape.Map_NoteShapePtr_NoteShapeDBID[noteshape]; ok {
			var noteshapeDB NoteShapeDB
			noteshapeDB.ID = id

			if _, err := backRepo.BackRepoNoteShape.db.First(&noteshapeDB, id); err != nil {
				log.Fatalln("CheckoutNoteShape : Problem with getting object with id:", id)
			}
			backRepo.BackRepoNoteShape.CheckoutPhaseOneInstance(&noteshapeDB)
			backRepo.BackRepoNoteShape.CheckoutPhaseTwoInstance(backRepo, &noteshapeDB)
		}
	}
}

// CopyBasicFieldsFromNoteShape
func (noteshapeDB *NoteShapeDB) CopyBasicFieldsFromNoteShape(noteshape *models.NoteShape) {
	// insertion point for fields commit

	noteshapeDB.Name_Data.String = noteshape.Name
	noteshapeDB.Name_Data.Valid = true

	noteshapeDB.Identifier_Data.String = noteshape.Identifier
	noteshapeDB.Identifier_Data.Valid = true

	noteshapeDB.Body_Data.String = noteshape.Body
	noteshapeDB.Body_Data.Valid = true

	noteshapeDB.BodyHTML_Data.String = noteshape.BodyHTML
	noteshapeDB.BodyHTML_Data.Valid = true

	noteshapeDB.X_Data.Float64 = noteshape.X
	noteshapeDB.X_Data.Valid = true

	noteshapeDB.Y_Data.Float64 = noteshape.Y
	noteshapeDB.Y_Data.Valid = true

	noteshapeDB.Width_Data.Float64 = noteshape.Width
	noteshapeDB.Width_Data.Valid = true

	noteshapeDB.Height_Data.Float64 = noteshape.Height
	noteshapeDB.Height_Data.Valid = true

	noteshapeDB.Matched_Data.Bool = noteshape.Matched
	noteshapeDB.Matched_Data.Valid = true
}

// CopyBasicFieldsFromNoteShape_WOP
func (noteshapeDB *NoteShapeDB) CopyBasicFieldsFromNoteShape_WOP(noteshape *models.NoteShape_WOP) {
	// insertion point for fields commit

	noteshapeDB.Name_Data.String = noteshape.Name
	noteshapeDB.Name_Data.Valid = true

	noteshapeDB.Identifier_Data.String = noteshape.Identifier
	noteshapeDB.Identifier_Data.Valid = true

	noteshapeDB.Body_Data.String = noteshape.Body
	noteshapeDB.Body_Data.Valid = true

	noteshapeDB.BodyHTML_Data.String = noteshape.BodyHTML
	noteshapeDB.BodyHTML_Data.Valid = true

	noteshapeDB.X_Data.Float64 = noteshape.X
	noteshapeDB.X_Data.Valid = true

	noteshapeDB.Y_Data.Float64 = noteshape.Y
	noteshapeDB.Y_Data.Valid = true

	noteshapeDB.Width_Data.Float64 = noteshape.Width
	noteshapeDB.Width_Data.Valid = true

	noteshapeDB.Height_Data.Float64 = noteshape.Height
	noteshapeDB.Height_Data.Valid = true

	noteshapeDB.Matched_Data.Bool = noteshape.Matched
	noteshapeDB.Matched_Data.Valid = true
}

// CopyBasicFieldsFromNoteShapeWOP
func (noteshapeDB *NoteShapeDB) CopyBasicFieldsFromNoteShapeWOP(noteshape *NoteShapeWOP) {
	// insertion point for fields commit

	noteshapeDB.Name_Data.String = noteshape.Name
	noteshapeDB.Name_Data.Valid = true

	noteshapeDB.Identifier_Data.String = noteshape.Identifier
	noteshapeDB.Identifier_Data.Valid = true

	noteshapeDB.Body_Data.String = noteshape.Body
	noteshapeDB.Body_Data.Valid = true

	noteshapeDB.BodyHTML_Data.String = noteshape.BodyHTML
	noteshapeDB.BodyHTML_Data.Valid = true

	noteshapeDB.X_Data.Float64 = noteshape.X
	noteshapeDB.X_Data.Valid = true

	noteshapeDB.Y_Data.Float64 = noteshape.Y
	noteshapeDB.Y_Data.Valid = true

	noteshapeDB.Width_Data.Float64 = noteshape.Width
	noteshapeDB.Width_Data.Valid = true

	noteshapeDB.Height_Data.Float64 = noteshape.Height
	noteshapeDB.Height_Data.Valid = true

	noteshapeDB.Matched_Data.Bool = noteshape.Matched
	noteshapeDB.Matched_Data.Valid = true
}

// CopyBasicFieldsToNoteShape
func (noteshapeDB *NoteShapeDB) CopyBasicFieldsToNoteShape(noteshape *models.NoteShape) {
	// insertion point for checkout of basic fields (back repo to stage)
	noteshape.Name = noteshapeDB.Name_Data.String
	noteshape.Identifier = noteshapeDB.Identifier_Data.String
	noteshape.Body = noteshapeDB.Body_Data.String
	noteshape.BodyHTML = noteshapeDB.BodyHTML_Data.String
	noteshape.X = noteshapeDB.X_Data.Float64
	noteshape.Y = noteshapeDB.Y_Data.Float64
	noteshape.Width = noteshapeDB.Width_Data.Float64
	noteshape.Height = noteshapeDB.Height_Data.Float64
	noteshape.Matched = noteshapeDB.Matched_Data.Bool
}

// CopyBasicFieldsToNoteShape_WOP
func (noteshapeDB *NoteShapeDB) CopyBasicFieldsToNoteShape_WOP(noteshape *models.NoteShape_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	noteshape.Name = noteshapeDB.Name_Data.String
	noteshape.Identifier = noteshapeDB.Identifier_Data.String
	noteshape.Body = noteshapeDB.Body_Data.String
	noteshape.BodyHTML = noteshapeDB.BodyHTML_Data.String
	noteshape.X = noteshapeDB.X_Data.Float64
	noteshape.Y = noteshapeDB.Y_Data.Float64
	noteshape.Width = noteshapeDB.Width_Data.Float64
	noteshape.Height = noteshapeDB.Height_Data.Float64
	noteshape.Matched = noteshapeDB.Matched_Data.Bool
}

// CopyBasicFieldsToNoteShapeWOP
func (noteshapeDB *NoteShapeDB) CopyBasicFieldsToNoteShapeWOP(noteshape *NoteShapeWOP) {
	noteshape.ID = int(noteshapeDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	noteshape.Name = noteshapeDB.Name_Data.String
	noteshape.Identifier = noteshapeDB.Identifier_Data.String
	noteshape.Body = noteshapeDB.Body_Data.String
	noteshape.BodyHTML = noteshapeDB.BodyHTML_Data.String
	noteshape.X = noteshapeDB.X_Data.Float64
	noteshape.Y = noteshapeDB.Y_Data.Float64
	noteshape.Width = noteshapeDB.Width_Data.Float64
	noteshape.Height = noteshapeDB.Height_Data.Float64
	noteshape.Matched = noteshapeDB.Matched_Data.Bool
}

// Backup generates a json file from a slice of all NoteShapeDB instances in the backrepo
func (backRepoNoteShape *BackRepoNoteShapeStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "NoteShapeDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*NoteShapeDB, 0)
	for _, noteshapeDB := range backRepoNoteShape.Map_NoteShapeDBID_NoteShapeDB {
		forBackup = append(forBackup, noteshapeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json NoteShape ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json NoteShape file", err.Error())
	}
}

// Backup generates a json file from a slice of all NoteShapeDB instances in the backrepo
func (backRepoNoteShape *BackRepoNoteShapeStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*NoteShapeDB, 0)
	for _, noteshapeDB := range backRepoNoteShape.Map_NoteShapeDBID_NoteShapeDB {
		forBackup = append(forBackup, noteshapeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("NoteShape")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&NoteShape_Fields, -1)
	for _, noteshapeDB := range forBackup {

		var noteshapeWOP NoteShapeWOP
		noteshapeDB.CopyBasicFieldsToNoteShapeWOP(&noteshapeWOP)

		row := sh.AddRow()
		row.WriteStruct(&noteshapeWOP, -1)
	}
}

// RestoreXL from the "NoteShape" sheet all NoteShapeDB instances
func (backRepoNoteShape *BackRepoNoteShapeStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoNoteShapeid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["NoteShape"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoNoteShape.rowVisitorNoteShape)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoNoteShape *BackRepoNoteShapeStruct) rowVisitorNoteShape(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var noteshapeWOP NoteShapeWOP
		row.ReadStruct(&noteshapeWOP)

		// add the unmarshalled struct to the stage
		noteshapeDB := new(NoteShapeDB)
		noteshapeDB.CopyBasicFieldsFromNoteShapeWOP(&noteshapeWOP)

		noteshapeDB_ID_atBackupTime := noteshapeDB.ID
		noteshapeDB.ID = 0
		_, err := backRepoNoteShape.db.Create(noteshapeDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoNoteShape.Map_NoteShapeDBID_NoteShapeDB[noteshapeDB.ID] = noteshapeDB
		BackRepoNoteShapeid_atBckpTime_newID[noteshapeDB_ID_atBackupTime] = noteshapeDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "NoteShapeDB.json" in dirPath that stores an array
// of NoteShapeDB and stores it in the database
// the map BackRepoNoteShapeid_atBckpTime_newID is updated accordingly
func (backRepoNoteShape *BackRepoNoteShapeStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoNoteShapeid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "NoteShapeDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json NoteShape file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*NoteShapeDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_NoteShapeDBID_NoteShapeDB
	for _, noteshapeDB := range forRestore {

		noteshapeDB_ID_atBackupTime := noteshapeDB.ID
		noteshapeDB.ID = 0
		_, err := backRepoNoteShape.db.Create(noteshapeDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoNoteShape.Map_NoteShapeDBID_NoteShapeDB[noteshapeDB.ID] = noteshapeDB
		BackRepoNoteShapeid_atBckpTime_newID[noteshapeDB_ID_atBackupTime] = noteshapeDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json NoteShape file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<NoteShape>id_atBckpTime_newID
// to compute new index
func (backRepoNoteShape *BackRepoNoteShapeStruct) RestorePhaseTwo() {

	for _, noteshapeDB := range backRepoNoteShape.Map_NoteShapeDBID_NoteShapeDB {

		// next line of code is to avert unused variable compilation error
		_ = noteshapeDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoNoteShape.db.Model(noteshapeDB)
		_, err := db.Updates(*noteshapeDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoNoteShape.ResetReversePointers commits all staged instances of NoteShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNoteShape *BackRepoNoteShapeStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, noteshape := range backRepoNoteShape.Map_NoteShapeDBID_NoteShapePtr {
		backRepoNoteShape.ResetReversePointersInstance(backRepo, idx, noteshape)
	}

	return
}

func (backRepoNoteShape *BackRepoNoteShapeStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, noteshape *models.NoteShape) (Error error) {

	// fetch matching noteshapeDB
	if noteshapeDB, ok := backRepoNoteShape.Map_NoteShapeDBID_NoteShapeDB[idx]; ok {
		_ = noteshapeDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoNoteShapeid_atBckpTime_newID map[uint]uint
