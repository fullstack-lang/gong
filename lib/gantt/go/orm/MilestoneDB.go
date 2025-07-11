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

	"github.com/fullstack-lang/gong/lib/gantt/go/db"
	"github.com/fullstack-lang/gong/lib/gantt/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Milestone_sql sql.NullBool
var dummy_Milestone_time time.Duration
var dummy_Milestone_sort sort.Float64Slice

// MilestoneAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model milestoneAPI
type MilestoneAPI struct {
	gorm.Model

	models.Milestone_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	MilestonePointersEncoding MilestonePointersEncoding
}

// MilestonePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type MilestonePointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field LanesToDisplay is a slice of pointers to another Struct (optional or 0..1)
	LanesToDisplay IntSlice `gorm:"type:TEXT"`
}

// MilestoneDB describes a milestone in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model milestoneDB
type MilestoneDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field milestoneDB.Name
	Name_Data sql.NullString

	// Declation for basic field milestoneDB.Date
	Date_Data sql.NullTime

	// Declation for basic field milestoneDB.DisplayVerticalBar
	// provide the sql storage for the boolan
	DisplayVerticalBar_Data sql.NullBool

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	MilestonePointersEncoding
}

// MilestoneDBs arrays milestoneDBs
// swagger:response milestoneDBsResponse
type MilestoneDBs []MilestoneDB

// MilestoneDBResponse provides response
// swagger:response milestoneDBResponse
type MilestoneDBResponse struct {
	MilestoneDB
}

// MilestoneWOP is a Milestone without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type MilestoneWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Date time.Time `xlsx:"2"`

	DisplayVerticalBar bool `xlsx:"3"`
	// insertion for WOP pointer fields
}

var Milestone_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Date",
	"DisplayVerticalBar",
}

type BackRepoMilestoneStruct struct {
	// stores MilestoneDB according to their gorm ID
	Map_MilestoneDBID_MilestoneDB map[uint]*MilestoneDB

	// stores MilestoneDB ID according to Milestone address
	Map_MilestonePtr_MilestoneDBID map[*models.Milestone]uint

	// stores Milestone according to their gorm ID
	Map_MilestoneDBID_MilestonePtr map[uint]*models.Milestone

	db db.DBInterface

	stage *models.Stage
}

func (backRepoMilestone *BackRepoMilestoneStruct) GetStage() (stage *models.Stage) {
	stage = backRepoMilestone.stage
	return
}

func (backRepoMilestone *BackRepoMilestoneStruct) GetDB() db.DBInterface {
	return backRepoMilestone.db
}

// GetMilestoneDBFromMilestonePtr is a handy function to access the back repo instance from the stage instance
func (backRepoMilestone *BackRepoMilestoneStruct) GetMilestoneDBFromMilestonePtr(milestone *models.Milestone) (milestoneDB *MilestoneDB) {
	id := backRepoMilestone.Map_MilestonePtr_MilestoneDBID[milestone]
	milestoneDB = backRepoMilestone.Map_MilestoneDBID_MilestoneDB[id]
	return
}

// BackRepoMilestone.CommitPhaseOne commits all staged instances of Milestone to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoMilestone *BackRepoMilestoneStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var milestones []*models.Milestone
	for milestone := range stage.Milestones {
		milestones = append(milestones, milestone)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(milestones, func(i, j int) bool {
		return stage.MilestoneMap_Staged_Order[milestones[i]] < stage.MilestoneMap_Staged_Order[milestones[j]]
	})

	for _, milestone := range milestones {
		backRepoMilestone.CommitPhaseOneInstance(milestone)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, milestone := range backRepoMilestone.Map_MilestoneDBID_MilestonePtr {
		if _, ok := stage.Milestones[milestone]; !ok {
			backRepoMilestone.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoMilestone.CommitDeleteInstance commits deletion of Milestone to the BackRepo
func (backRepoMilestone *BackRepoMilestoneStruct) CommitDeleteInstance(id uint) (Error error) {

	milestone := backRepoMilestone.Map_MilestoneDBID_MilestonePtr[id]

	// milestone is not staged anymore, remove milestoneDB
	milestoneDB := backRepoMilestone.Map_MilestoneDBID_MilestoneDB[id]
	db, _ := backRepoMilestone.db.Unscoped()
	_, err := db.Delete(milestoneDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoMilestone.Map_MilestonePtr_MilestoneDBID, milestone)
	delete(backRepoMilestone.Map_MilestoneDBID_MilestonePtr, id)
	delete(backRepoMilestone.Map_MilestoneDBID_MilestoneDB, id)

	return
}

// BackRepoMilestone.CommitPhaseOneInstance commits milestone staged instances of Milestone to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoMilestone *BackRepoMilestoneStruct) CommitPhaseOneInstance(milestone *models.Milestone) (Error error) {

	// check if the milestone is not commited yet
	if _, ok := backRepoMilestone.Map_MilestonePtr_MilestoneDBID[milestone]; ok {
		return
	}

	// initiate milestone
	var milestoneDB MilestoneDB
	milestoneDB.CopyBasicFieldsFromMilestone(milestone)

	_, err := backRepoMilestone.db.Create(&milestoneDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoMilestone.Map_MilestonePtr_MilestoneDBID[milestone] = milestoneDB.ID
	backRepoMilestone.Map_MilestoneDBID_MilestonePtr[milestoneDB.ID] = milestone
	backRepoMilestone.Map_MilestoneDBID_MilestoneDB[milestoneDB.ID] = &milestoneDB

	return
}

// BackRepoMilestone.CommitPhaseTwo commits all staged instances of Milestone to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMilestone *BackRepoMilestoneStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, milestone := range backRepoMilestone.Map_MilestoneDBID_MilestonePtr {
		backRepoMilestone.CommitPhaseTwoInstance(backRepo, idx, milestone)
	}

	return
}

// BackRepoMilestone.CommitPhaseTwoInstance commits {{structname }} of models.Milestone to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMilestone *BackRepoMilestoneStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, milestone *models.Milestone) (Error error) {

	// fetch matching milestoneDB
	if milestoneDB, ok := backRepoMilestone.Map_MilestoneDBID_MilestoneDB[idx]; ok {

		milestoneDB.CopyBasicFieldsFromMilestone(milestone)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		milestoneDB.MilestonePointersEncoding.LanesToDisplay = make([]int, 0)
		// 2. encode
		for _, laneAssocEnd := range milestone.LanesToDisplay {
			laneAssocEnd_DB :=
				backRepo.BackRepoLane.GetLaneDBFromLanePtr(laneAssocEnd)
			
			// the stage might be inconsistant, meaning that the laneAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if laneAssocEnd_DB == nil {
				continue
			}
			
			milestoneDB.MilestonePointersEncoding.LanesToDisplay =
				append(milestoneDB.MilestonePointersEncoding.LanesToDisplay, int(laneAssocEnd_DB.ID))
		}

		_, err := backRepoMilestone.db.Save(milestoneDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Milestone intance %s", milestone.Name))
		return err
	}

	return
}

// BackRepoMilestone.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoMilestone *BackRepoMilestoneStruct) CheckoutPhaseOne() (Error error) {

	milestoneDBArray := make([]MilestoneDB, 0)
	_, err := backRepoMilestone.db.Find(&milestoneDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	milestoneInstancesToBeRemovedFromTheStage := make(map[*models.Milestone]any)
	for key, value := range backRepoMilestone.stage.Milestones {
		milestoneInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, milestoneDB := range milestoneDBArray {
		backRepoMilestone.CheckoutPhaseOneInstance(&milestoneDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		milestone, ok := backRepoMilestone.Map_MilestoneDBID_MilestonePtr[milestoneDB.ID]
		if ok {
			delete(milestoneInstancesToBeRemovedFromTheStage, milestone)
		}
	}

	// remove from stage and back repo's 3 maps all milestones that are not in the checkout
	for milestone := range milestoneInstancesToBeRemovedFromTheStage {
		milestone.Unstage(backRepoMilestone.GetStage())

		// remove instance from the back repo 3 maps
		milestoneID := backRepoMilestone.Map_MilestonePtr_MilestoneDBID[milestone]
		delete(backRepoMilestone.Map_MilestonePtr_MilestoneDBID, milestone)
		delete(backRepoMilestone.Map_MilestoneDBID_MilestoneDB, milestoneID)
		delete(backRepoMilestone.Map_MilestoneDBID_MilestonePtr, milestoneID)
	}

	return
}

// CheckoutPhaseOneInstance takes a milestoneDB that has been found in the DB, updates the backRepo and stages the
// models version of the milestoneDB
func (backRepoMilestone *BackRepoMilestoneStruct) CheckoutPhaseOneInstance(milestoneDB *MilestoneDB) (Error error) {

	milestone, ok := backRepoMilestone.Map_MilestoneDBID_MilestonePtr[milestoneDB.ID]
	if !ok {
		milestone = new(models.Milestone)

		backRepoMilestone.Map_MilestoneDBID_MilestonePtr[milestoneDB.ID] = milestone
		backRepoMilestone.Map_MilestonePtr_MilestoneDBID[milestone] = milestoneDB.ID

		// append model store with the new element
		milestone.Name = milestoneDB.Name_Data.String
		milestone.Stage(backRepoMilestone.GetStage())
	}
	milestoneDB.CopyBasicFieldsToMilestone(milestone)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	milestone.Stage(backRepoMilestone.GetStage())

	// preserve pointer to milestoneDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_MilestoneDBID_MilestoneDB)[milestoneDB hold variable pointers
	milestoneDB_Data := *milestoneDB
	preservedPtrToMilestone := &milestoneDB_Data
	backRepoMilestone.Map_MilestoneDBID_MilestoneDB[milestoneDB.ID] = preservedPtrToMilestone

	return
}

// BackRepoMilestone.CheckoutPhaseTwo Checkouts all staged instances of Milestone to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMilestone *BackRepoMilestoneStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, milestoneDB := range backRepoMilestone.Map_MilestoneDBID_MilestoneDB {
		backRepoMilestone.CheckoutPhaseTwoInstance(backRepo, milestoneDB)
	}
	return
}

// BackRepoMilestone.CheckoutPhaseTwoInstance Checkouts staged instances of Milestone to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMilestone *BackRepoMilestoneStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, milestoneDB *MilestoneDB) (Error error) {

	milestone := backRepoMilestone.Map_MilestoneDBID_MilestonePtr[milestoneDB.ID]

	milestoneDB.DecodePointers(backRepo, milestone)

	return
}

func (milestoneDB *MilestoneDB) DecodePointers(backRepo *BackRepoStruct, milestone *models.Milestone) {

	// insertion point for checkout of pointer encoding
	// This loop redeem milestone.LanesToDisplay in the stage from the encode in the back repo
	// It parses all LaneDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	milestone.LanesToDisplay = milestone.LanesToDisplay[:0]
	for _, _Laneid := range milestoneDB.MilestonePointersEncoding.LanesToDisplay {
		milestone.LanesToDisplay = append(milestone.LanesToDisplay, backRepo.BackRepoLane.Map_LaneDBID_LanePtr[uint(_Laneid)])
	}

	return
}

// CommitMilestone allows commit of a single milestone (if already staged)
func (backRepo *BackRepoStruct) CommitMilestone(milestone *models.Milestone) {
	backRepo.BackRepoMilestone.CommitPhaseOneInstance(milestone)
	if id, ok := backRepo.BackRepoMilestone.Map_MilestonePtr_MilestoneDBID[milestone]; ok {
		backRepo.BackRepoMilestone.CommitPhaseTwoInstance(backRepo, id, milestone)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitMilestone allows checkout of a single milestone (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutMilestone(milestone *models.Milestone) {
	// check if the milestone is staged
	if _, ok := backRepo.BackRepoMilestone.Map_MilestonePtr_MilestoneDBID[milestone]; ok {

		if id, ok := backRepo.BackRepoMilestone.Map_MilestonePtr_MilestoneDBID[milestone]; ok {
			var milestoneDB MilestoneDB
			milestoneDB.ID = id

			if _, err := backRepo.BackRepoMilestone.db.First(&milestoneDB, id); err != nil {
				log.Fatalln("CheckoutMilestone : Problem with getting object with id:", id)
			}
			backRepo.BackRepoMilestone.CheckoutPhaseOneInstance(&milestoneDB)
			backRepo.BackRepoMilestone.CheckoutPhaseTwoInstance(backRepo, &milestoneDB)
		}
	}
}

// CopyBasicFieldsFromMilestone
func (milestoneDB *MilestoneDB) CopyBasicFieldsFromMilestone(milestone *models.Milestone) {
	// insertion point for fields commit

	milestoneDB.Name_Data.String = milestone.Name
	milestoneDB.Name_Data.Valid = true

	milestoneDB.Date_Data.Time = milestone.Date
	milestoneDB.Date_Data.Valid = true

	milestoneDB.DisplayVerticalBar_Data.Bool = milestone.DisplayVerticalBar
	milestoneDB.DisplayVerticalBar_Data.Valid = true
}

// CopyBasicFieldsFromMilestone_WOP
func (milestoneDB *MilestoneDB) CopyBasicFieldsFromMilestone_WOP(milestone *models.Milestone_WOP) {
	// insertion point for fields commit

	milestoneDB.Name_Data.String = milestone.Name
	milestoneDB.Name_Data.Valid = true

	milestoneDB.Date_Data.Time = milestone.Date
	milestoneDB.Date_Data.Valid = true

	milestoneDB.DisplayVerticalBar_Data.Bool = milestone.DisplayVerticalBar
	milestoneDB.DisplayVerticalBar_Data.Valid = true
}

// CopyBasicFieldsFromMilestoneWOP
func (milestoneDB *MilestoneDB) CopyBasicFieldsFromMilestoneWOP(milestone *MilestoneWOP) {
	// insertion point for fields commit

	milestoneDB.Name_Data.String = milestone.Name
	milestoneDB.Name_Data.Valid = true

	milestoneDB.Date_Data.Time = milestone.Date
	milestoneDB.Date_Data.Valid = true

	milestoneDB.DisplayVerticalBar_Data.Bool = milestone.DisplayVerticalBar
	milestoneDB.DisplayVerticalBar_Data.Valid = true
}

// CopyBasicFieldsToMilestone
func (milestoneDB *MilestoneDB) CopyBasicFieldsToMilestone(milestone *models.Milestone) {
	// insertion point for checkout of basic fields (back repo to stage)
	milestone.Name = milestoneDB.Name_Data.String
	milestone.Date = milestoneDB.Date_Data.Time
	milestone.DisplayVerticalBar = milestoneDB.DisplayVerticalBar_Data.Bool
}

// CopyBasicFieldsToMilestone_WOP
func (milestoneDB *MilestoneDB) CopyBasicFieldsToMilestone_WOP(milestone *models.Milestone_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	milestone.Name = milestoneDB.Name_Data.String
	milestone.Date = milestoneDB.Date_Data.Time
	milestone.DisplayVerticalBar = milestoneDB.DisplayVerticalBar_Data.Bool
}

// CopyBasicFieldsToMilestoneWOP
func (milestoneDB *MilestoneDB) CopyBasicFieldsToMilestoneWOP(milestone *MilestoneWOP) {
	milestone.ID = int(milestoneDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	milestone.Name = milestoneDB.Name_Data.String
	milestone.Date = milestoneDB.Date_Data.Time
	milestone.DisplayVerticalBar = milestoneDB.DisplayVerticalBar_Data.Bool
}

// Backup generates a json file from a slice of all MilestoneDB instances in the backrepo
func (backRepoMilestone *BackRepoMilestoneStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "MilestoneDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*MilestoneDB, 0)
	for _, milestoneDB := range backRepoMilestone.Map_MilestoneDBID_MilestoneDB {
		forBackup = append(forBackup, milestoneDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Milestone ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Milestone file", err.Error())
	}
}

// Backup generates a json file from a slice of all MilestoneDB instances in the backrepo
func (backRepoMilestone *BackRepoMilestoneStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*MilestoneDB, 0)
	for _, milestoneDB := range backRepoMilestone.Map_MilestoneDBID_MilestoneDB {
		forBackup = append(forBackup, milestoneDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Milestone")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Milestone_Fields, -1)
	for _, milestoneDB := range forBackup {

		var milestoneWOP MilestoneWOP
		milestoneDB.CopyBasicFieldsToMilestoneWOP(&milestoneWOP)

		row := sh.AddRow()
		row.WriteStruct(&milestoneWOP, -1)
	}
}

// RestoreXL from the "Milestone" sheet all MilestoneDB instances
func (backRepoMilestone *BackRepoMilestoneStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoMilestoneid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Milestone"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoMilestone.rowVisitorMilestone)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoMilestone *BackRepoMilestoneStruct) rowVisitorMilestone(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var milestoneWOP MilestoneWOP
		row.ReadStruct(&milestoneWOP)

		// add the unmarshalled struct to the stage
		milestoneDB := new(MilestoneDB)
		milestoneDB.CopyBasicFieldsFromMilestoneWOP(&milestoneWOP)

		milestoneDB_ID_atBackupTime := milestoneDB.ID
		milestoneDB.ID = 0
		_, err := backRepoMilestone.db.Create(milestoneDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoMilestone.Map_MilestoneDBID_MilestoneDB[milestoneDB.ID] = milestoneDB
		BackRepoMilestoneid_atBckpTime_newID[milestoneDB_ID_atBackupTime] = milestoneDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "MilestoneDB.json" in dirPath that stores an array
// of MilestoneDB and stores it in the database
// the map BackRepoMilestoneid_atBckpTime_newID is updated accordingly
func (backRepoMilestone *BackRepoMilestoneStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoMilestoneid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "MilestoneDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Milestone file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*MilestoneDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_MilestoneDBID_MilestoneDB
	for _, milestoneDB := range forRestore {

		milestoneDB_ID_atBackupTime := milestoneDB.ID
		milestoneDB.ID = 0
		_, err := backRepoMilestone.db.Create(milestoneDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoMilestone.Map_MilestoneDBID_MilestoneDB[milestoneDB.ID] = milestoneDB
		BackRepoMilestoneid_atBckpTime_newID[milestoneDB_ID_atBackupTime] = milestoneDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Milestone file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Milestone>id_atBckpTime_newID
// to compute new index
func (backRepoMilestone *BackRepoMilestoneStruct) RestorePhaseTwo() {

	for _, milestoneDB := range backRepoMilestone.Map_MilestoneDBID_MilestoneDB {

		// next line of code is to avert unused variable compilation error
		_ = milestoneDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoMilestone.db.Model(milestoneDB)
		_, err := db.Updates(*milestoneDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoMilestone.ResetReversePointers commits all staged instances of Milestone to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMilestone *BackRepoMilestoneStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, milestone := range backRepoMilestone.Map_MilestoneDBID_MilestonePtr {
		backRepoMilestone.ResetReversePointersInstance(backRepo, idx, milestone)
	}

	return
}

func (backRepoMilestone *BackRepoMilestoneStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, milestone *models.Milestone) (Error error) {

	// fetch matching milestoneDB
	if milestoneDB, ok := backRepoMilestone.Map_MilestoneDBID_MilestoneDB[idx]; ok {
		_ = milestoneDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoMilestoneid_atBckpTime_newID map[uint]uint
