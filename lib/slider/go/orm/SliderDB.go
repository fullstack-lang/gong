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

	"github.com/fullstack-lang/gong/lib/slider/go/db"
	"github.com/fullstack-lang/gong/lib/slider/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Slider_sql sql.NullBool
var dummy_Slider_time time.Duration
var dummy_Slider_sort sort.Float64Slice

// SliderAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model sliderAPI
type SliderAPI struct {
	gorm.Model

	models.Slider_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	SliderPointersEncoding SliderPointersEncoding
}

// SliderPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type SliderPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// SliderDB describes a slider in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model sliderDB
type SliderDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field sliderDB.Name
	Name_Data sql.NullString

	// Declation for basic field sliderDB.IsFloat64
	// provide the sql storage for the boolan
	IsFloat64_Data sql.NullBool

	// Declation for basic field sliderDB.IsInt
	// provide the sql storage for the boolan
	IsInt_Data sql.NullBool

	// Declation for basic field sliderDB.MinInt
	MinInt_Data sql.NullInt64

	// Declation for basic field sliderDB.MaxInt
	MaxInt_Data sql.NullInt64

	// Declation for basic field sliderDB.StepInt
	StepInt_Data sql.NullInt64

	// Declation for basic field sliderDB.ValueInt
	ValueInt_Data sql.NullInt64

	// Declation for basic field sliderDB.MinFloat64
	MinFloat64_Data sql.NullFloat64

	// Declation for basic field sliderDB.MaxFloat64
	MaxFloat64_Data sql.NullFloat64

	// Declation for basic field sliderDB.StepFloat64
	StepFloat64_Data sql.NullFloat64

	// Declation for basic field sliderDB.ValueFloat64
	ValueFloat64_Data sql.NullFloat64

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	SliderPointersEncoding
}

// SliderDBs arrays sliderDBs
// swagger:response sliderDBsResponse
type SliderDBs []SliderDB

// SliderDBResponse provides response
// swagger:response sliderDBResponse
type SliderDBResponse struct {
	SliderDB
}

// SliderWOP is a Slider without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type SliderWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	IsFloat64 bool `xlsx:"2"`

	IsInt bool `xlsx:"3"`

	MinInt int `xlsx:"4"`

	MaxInt int `xlsx:"5"`

	StepInt int `xlsx:"6"`

	ValueInt int `xlsx:"7"`

	MinFloat64 float64 `xlsx:"8"`

	MaxFloat64 float64 `xlsx:"9"`

	StepFloat64 float64 `xlsx:"10"`

	ValueFloat64 float64 `xlsx:"11"`
	// insertion for WOP pointer fields
}

var Slider_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"IsFloat64",
	"IsInt",
	"MinInt",
	"MaxInt",
	"StepInt",
	"ValueInt",
	"MinFloat64",
	"MaxFloat64",
	"StepFloat64",
	"ValueFloat64",
}

type BackRepoSliderStruct struct {
	// stores SliderDB according to their gorm ID
	Map_SliderDBID_SliderDB map[uint]*SliderDB

	// stores SliderDB ID according to Slider address
	Map_SliderPtr_SliderDBID map[*models.Slider]uint

	// stores Slider according to their gorm ID
	Map_SliderDBID_SliderPtr map[uint]*models.Slider

	db db.DBInterface

	stage *models.Stage
}

func (backRepoSlider *BackRepoSliderStruct) GetStage() (stage *models.Stage) {
	stage = backRepoSlider.stage
	return
}

func (backRepoSlider *BackRepoSliderStruct) GetDB() db.DBInterface {
	return backRepoSlider.db
}

// GetSliderDBFromSliderPtr is a handy function to access the back repo instance from the stage instance
func (backRepoSlider *BackRepoSliderStruct) GetSliderDBFromSliderPtr(slider *models.Slider) (sliderDB *SliderDB) {
	id := backRepoSlider.Map_SliderPtr_SliderDBID[slider]
	sliderDB = backRepoSlider.Map_SliderDBID_SliderDB[id]
	return
}

// BackRepoSlider.CommitPhaseOne commits all staged instances of Slider to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSlider *BackRepoSliderStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var sliders []*models.Slider
	for slider := range stage.Sliders {
		sliders = append(sliders, slider)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(sliders, func(i, j int) bool {
		return stage.SliderMap_Staged_Order[sliders[i]] < stage.SliderMap_Staged_Order[sliders[j]]
	})

	for _, slider := range sliders {
		backRepoSlider.CommitPhaseOneInstance(slider)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, slider := range backRepoSlider.Map_SliderDBID_SliderPtr {
		if _, ok := stage.Sliders[slider]; !ok {
			backRepoSlider.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoSlider.CommitDeleteInstance commits deletion of Slider to the BackRepo
func (backRepoSlider *BackRepoSliderStruct) CommitDeleteInstance(id uint) (Error error) {

	slider := backRepoSlider.Map_SliderDBID_SliderPtr[id]

	// slider is not staged anymore, remove sliderDB
	sliderDB := backRepoSlider.Map_SliderDBID_SliderDB[id]
	db, _ := backRepoSlider.db.Unscoped()
	_, err := db.Delete(sliderDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoSlider.Map_SliderPtr_SliderDBID, slider)
	delete(backRepoSlider.Map_SliderDBID_SliderPtr, id)
	delete(backRepoSlider.Map_SliderDBID_SliderDB, id)

	return
}

// BackRepoSlider.CommitPhaseOneInstance commits slider staged instances of Slider to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSlider *BackRepoSliderStruct) CommitPhaseOneInstance(slider *models.Slider) (Error error) {

	// check if the slider is not commited yet
	if _, ok := backRepoSlider.Map_SliderPtr_SliderDBID[slider]; ok {
		return
	}

	// initiate slider
	var sliderDB SliderDB
	sliderDB.CopyBasicFieldsFromSlider(slider)

	_, err := backRepoSlider.db.Create(&sliderDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoSlider.Map_SliderPtr_SliderDBID[slider] = sliderDB.ID
	backRepoSlider.Map_SliderDBID_SliderPtr[sliderDB.ID] = slider
	backRepoSlider.Map_SliderDBID_SliderDB[sliderDB.ID] = &sliderDB

	return
}

// BackRepoSlider.CommitPhaseTwo commits all staged instances of Slider to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSlider *BackRepoSliderStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, slider := range backRepoSlider.Map_SliderDBID_SliderPtr {
		backRepoSlider.CommitPhaseTwoInstance(backRepo, idx, slider)
	}

	return
}

// BackRepoSlider.CommitPhaseTwoInstance commits {{structname }} of models.Slider to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSlider *BackRepoSliderStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, slider *models.Slider) (Error error) {

	// fetch matching sliderDB
	if sliderDB, ok := backRepoSlider.Map_SliderDBID_SliderDB[idx]; ok {

		sliderDB.CopyBasicFieldsFromSlider(slider)

		// insertion point for translating pointers encodings into actual pointers
		_, err := backRepoSlider.db.Save(sliderDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Slider intance %s", slider.Name))
		return err
	}

	return
}

// BackRepoSlider.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoSlider *BackRepoSliderStruct) CheckoutPhaseOne() (Error error) {

	sliderDBArray := make([]SliderDB, 0)
	_, err := backRepoSlider.db.Find(&sliderDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	sliderInstancesToBeRemovedFromTheStage := make(map[*models.Slider]any)
	for key, value := range backRepoSlider.stage.Sliders {
		sliderInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, sliderDB := range sliderDBArray {
		backRepoSlider.CheckoutPhaseOneInstance(&sliderDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		slider, ok := backRepoSlider.Map_SliderDBID_SliderPtr[sliderDB.ID]
		if ok {
			delete(sliderInstancesToBeRemovedFromTheStage, slider)
		}
	}

	// remove from stage and back repo's 3 maps all sliders that are not in the checkout
	for slider := range sliderInstancesToBeRemovedFromTheStage {
		slider.Unstage(backRepoSlider.GetStage())

		// remove instance from the back repo 3 maps
		sliderID := backRepoSlider.Map_SliderPtr_SliderDBID[slider]
		delete(backRepoSlider.Map_SliderPtr_SliderDBID, slider)
		delete(backRepoSlider.Map_SliderDBID_SliderDB, sliderID)
		delete(backRepoSlider.Map_SliderDBID_SliderPtr, sliderID)
	}

	return
}

// CheckoutPhaseOneInstance takes a sliderDB that has been found in the DB, updates the backRepo and stages the
// models version of the sliderDB
func (backRepoSlider *BackRepoSliderStruct) CheckoutPhaseOneInstance(sliderDB *SliderDB) (Error error) {

	slider, ok := backRepoSlider.Map_SliderDBID_SliderPtr[sliderDB.ID]
	if !ok {
		slider = new(models.Slider)

		backRepoSlider.Map_SliderDBID_SliderPtr[sliderDB.ID] = slider
		backRepoSlider.Map_SliderPtr_SliderDBID[slider] = sliderDB.ID

		// append model store with the new element
		slider.Name = sliderDB.Name_Data.String
		slider.Stage(backRepoSlider.GetStage())
	}
	sliderDB.CopyBasicFieldsToSlider(slider)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	slider.Stage(backRepoSlider.GetStage())

	// preserve pointer to sliderDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_SliderDBID_SliderDB)[sliderDB hold variable pointers
	sliderDB_Data := *sliderDB
	preservedPtrToSlider := &sliderDB_Data
	backRepoSlider.Map_SliderDBID_SliderDB[sliderDB.ID] = preservedPtrToSlider

	return
}

// BackRepoSlider.CheckoutPhaseTwo Checkouts all staged instances of Slider to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSlider *BackRepoSliderStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, sliderDB := range backRepoSlider.Map_SliderDBID_SliderDB {
		backRepoSlider.CheckoutPhaseTwoInstance(backRepo, sliderDB)
	}
	return
}

// BackRepoSlider.CheckoutPhaseTwoInstance Checkouts staged instances of Slider to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSlider *BackRepoSliderStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, sliderDB *SliderDB) (Error error) {

	slider := backRepoSlider.Map_SliderDBID_SliderPtr[sliderDB.ID]

	sliderDB.DecodePointers(backRepo, slider)

	return
}

func (sliderDB *SliderDB) DecodePointers(backRepo *BackRepoStruct, slider *models.Slider) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitSlider allows commit of a single slider (if already staged)
func (backRepo *BackRepoStruct) CommitSlider(slider *models.Slider) {
	backRepo.BackRepoSlider.CommitPhaseOneInstance(slider)
	if id, ok := backRepo.BackRepoSlider.Map_SliderPtr_SliderDBID[slider]; ok {
		backRepo.BackRepoSlider.CommitPhaseTwoInstance(backRepo, id, slider)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitSlider allows checkout of a single slider (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutSlider(slider *models.Slider) {
	// check if the slider is staged
	if _, ok := backRepo.BackRepoSlider.Map_SliderPtr_SliderDBID[slider]; ok {

		if id, ok := backRepo.BackRepoSlider.Map_SliderPtr_SliderDBID[slider]; ok {
			var sliderDB SliderDB
			sliderDB.ID = id

			if _, err := backRepo.BackRepoSlider.db.First(&sliderDB, id); err != nil {
				log.Fatalln("CheckoutSlider : Problem with getting object with id:", id)
			}
			backRepo.BackRepoSlider.CheckoutPhaseOneInstance(&sliderDB)
			backRepo.BackRepoSlider.CheckoutPhaseTwoInstance(backRepo, &sliderDB)
		}
	}
}

// CopyBasicFieldsFromSlider
func (sliderDB *SliderDB) CopyBasicFieldsFromSlider(slider *models.Slider) {
	// insertion point for fields commit

	sliderDB.Name_Data.String = slider.Name
	sliderDB.Name_Data.Valid = true

	sliderDB.IsFloat64_Data.Bool = slider.IsFloat64
	sliderDB.IsFloat64_Data.Valid = true

	sliderDB.IsInt_Data.Bool = slider.IsInt
	sliderDB.IsInt_Data.Valid = true

	sliderDB.MinInt_Data.Int64 = int64(slider.MinInt)
	sliderDB.MinInt_Data.Valid = true

	sliderDB.MaxInt_Data.Int64 = int64(slider.MaxInt)
	sliderDB.MaxInt_Data.Valid = true

	sliderDB.StepInt_Data.Int64 = int64(slider.StepInt)
	sliderDB.StepInt_Data.Valid = true

	sliderDB.ValueInt_Data.Int64 = int64(slider.ValueInt)
	sliderDB.ValueInt_Data.Valid = true

	sliderDB.MinFloat64_Data.Float64 = slider.MinFloat64
	sliderDB.MinFloat64_Data.Valid = true

	sliderDB.MaxFloat64_Data.Float64 = slider.MaxFloat64
	sliderDB.MaxFloat64_Data.Valid = true

	sliderDB.StepFloat64_Data.Float64 = slider.StepFloat64
	sliderDB.StepFloat64_Data.Valid = true

	sliderDB.ValueFloat64_Data.Float64 = slider.ValueFloat64
	sliderDB.ValueFloat64_Data.Valid = true
}

// CopyBasicFieldsFromSlider_WOP
func (sliderDB *SliderDB) CopyBasicFieldsFromSlider_WOP(slider *models.Slider_WOP) {
	// insertion point for fields commit

	sliderDB.Name_Data.String = slider.Name
	sliderDB.Name_Data.Valid = true

	sliderDB.IsFloat64_Data.Bool = slider.IsFloat64
	sliderDB.IsFloat64_Data.Valid = true

	sliderDB.IsInt_Data.Bool = slider.IsInt
	sliderDB.IsInt_Data.Valid = true

	sliderDB.MinInt_Data.Int64 = int64(slider.MinInt)
	sliderDB.MinInt_Data.Valid = true

	sliderDB.MaxInt_Data.Int64 = int64(slider.MaxInt)
	sliderDB.MaxInt_Data.Valid = true

	sliderDB.StepInt_Data.Int64 = int64(slider.StepInt)
	sliderDB.StepInt_Data.Valid = true

	sliderDB.ValueInt_Data.Int64 = int64(slider.ValueInt)
	sliderDB.ValueInt_Data.Valid = true

	sliderDB.MinFloat64_Data.Float64 = slider.MinFloat64
	sliderDB.MinFloat64_Data.Valid = true

	sliderDB.MaxFloat64_Data.Float64 = slider.MaxFloat64
	sliderDB.MaxFloat64_Data.Valid = true

	sliderDB.StepFloat64_Data.Float64 = slider.StepFloat64
	sliderDB.StepFloat64_Data.Valid = true

	sliderDB.ValueFloat64_Data.Float64 = slider.ValueFloat64
	sliderDB.ValueFloat64_Data.Valid = true
}

// CopyBasicFieldsFromSliderWOP
func (sliderDB *SliderDB) CopyBasicFieldsFromSliderWOP(slider *SliderWOP) {
	// insertion point for fields commit

	sliderDB.Name_Data.String = slider.Name
	sliderDB.Name_Data.Valid = true

	sliderDB.IsFloat64_Data.Bool = slider.IsFloat64
	sliderDB.IsFloat64_Data.Valid = true

	sliderDB.IsInt_Data.Bool = slider.IsInt
	sliderDB.IsInt_Data.Valid = true

	sliderDB.MinInt_Data.Int64 = int64(slider.MinInt)
	sliderDB.MinInt_Data.Valid = true

	sliderDB.MaxInt_Data.Int64 = int64(slider.MaxInt)
	sliderDB.MaxInt_Data.Valid = true

	sliderDB.StepInt_Data.Int64 = int64(slider.StepInt)
	sliderDB.StepInt_Data.Valid = true

	sliderDB.ValueInt_Data.Int64 = int64(slider.ValueInt)
	sliderDB.ValueInt_Data.Valid = true

	sliderDB.MinFloat64_Data.Float64 = slider.MinFloat64
	sliderDB.MinFloat64_Data.Valid = true

	sliderDB.MaxFloat64_Data.Float64 = slider.MaxFloat64
	sliderDB.MaxFloat64_Data.Valid = true

	sliderDB.StepFloat64_Data.Float64 = slider.StepFloat64
	sliderDB.StepFloat64_Data.Valid = true

	sliderDB.ValueFloat64_Data.Float64 = slider.ValueFloat64
	sliderDB.ValueFloat64_Data.Valid = true
}

// CopyBasicFieldsToSlider
func (sliderDB *SliderDB) CopyBasicFieldsToSlider(slider *models.Slider) {
	// insertion point for checkout of basic fields (back repo to stage)
	slider.Name = sliderDB.Name_Data.String
	slider.IsFloat64 = sliderDB.IsFloat64_Data.Bool
	slider.IsInt = sliderDB.IsInt_Data.Bool
	slider.MinInt = int(sliderDB.MinInt_Data.Int64)
	slider.MaxInt = int(sliderDB.MaxInt_Data.Int64)
	slider.StepInt = int(sliderDB.StepInt_Data.Int64)
	slider.ValueInt = int(sliderDB.ValueInt_Data.Int64)
	slider.MinFloat64 = sliderDB.MinFloat64_Data.Float64
	slider.MaxFloat64 = sliderDB.MaxFloat64_Data.Float64
	slider.StepFloat64 = sliderDB.StepFloat64_Data.Float64
	slider.ValueFloat64 = sliderDB.ValueFloat64_Data.Float64
}

// CopyBasicFieldsToSlider_WOP
func (sliderDB *SliderDB) CopyBasicFieldsToSlider_WOP(slider *models.Slider_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	slider.Name = sliderDB.Name_Data.String
	slider.IsFloat64 = sliderDB.IsFloat64_Data.Bool
	slider.IsInt = sliderDB.IsInt_Data.Bool
	slider.MinInt = int(sliderDB.MinInt_Data.Int64)
	slider.MaxInt = int(sliderDB.MaxInt_Data.Int64)
	slider.StepInt = int(sliderDB.StepInt_Data.Int64)
	slider.ValueInt = int(sliderDB.ValueInt_Data.Int64)
	slider.MinFloat64 = sliderDB.MinFloat64_Data.Float64
	slider.MaxFloat64 = sliderDB.MaxFloat64_Data.Float64
	slider.StepFloat64 = sliderDB.StepFloat64_Data.Float64
	slider.ValueFloat64 = sliderDB.ValueFloat64_Data.Float64
}

// CopyBasicFieldsToSliderWOP
func (sliderDB *SliderDB) CopyBasicFieldsToSliderWOP(slider *SliderWOP) {
	slider.ID = int(sliderDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	slider.Name = sliderDB.Name_Data.String
	slider.IsFloat64 = sliderDB.IsFloat64_Data.Bool
	slider.IsInt = sliderDB.IsInt_Data.Bool
	slider.MinInt = int(sliderDB.MinInt_Data.Int64)
	slider.MaxInt = int(sliderDB.MaxInt_Data.Int64)
	slider.StepInt = int(sliderDB.StepInt_Data.Int64)
	slider.ValueInt = int(sliderDB.ValueInt_Data.Int64)
	slider.MinFloat64 = sliderDB.MinFloat64_Data.Float64
	slider.MaxFloat64 = sliderDB.MaxFloat64_Data.Float64
	slider.StepFloat64 = sliderDB.StepFloat64_Data.Float64
	slider.ValueFloat64 = sliderDB.ValueFloat64_Data.Float64
}

// Backup generates a json file from a slice of all SliderDB instances in the backrepo
func (backRepoSlider *BackRepoSliderStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "SliderDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SliderDB, 0)
	for _, sliderDB := range backRepoSlider.Map_SliderDBID_SliderDB {
		forBackup = append(forBackup, sliderDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Slider ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Slider file", err.Error())
	}
}

// Backup generates a json file from a slice of all SliderDB instances in the backrepo
func (backRepoSlider *BackRepoSliderStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SliderDB, 0)
	for _, sliderDB := range backRepoSlider.Map_SliderDBID_SliderDB {
		forBackup = append(forBackup, sliderDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Slider")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Slider_Fields, -1)
	for _, sliderDB := range forBackup {

		var sliderWOP SliderWOP
		sliderDB.CopyBasicFieldsToSliderWOP(&sliderWOP)

		row := sh.AddRow()
		row.WriteStruct(&sliderWOP, -1)
	}
}

// RestoreXL from the "Slider" sheet all SliderDB instances
func (backRepoSlider *BackRepoSliderStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoSliderid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Slider"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoSlider.rowVisitorSlider)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoSlider *BackRepoSliderStruct) rowVisitorSlider(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var sliderWOP SliderWOP
		row.ReadStruct(&sliderWOP)

		// add the unmarshalled struct to the stage
		sliderDB := new(SliderDB)
		sliderDB.CopyBasicFieldsFromSliderWOP(&sliderWOP)

		sliderDB_ID_atBackupTime := sliderDB.ID
		sliderDB.ID = 0
		_, err := backRepoSlider.db.Create(sliderDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoSlider.Map_SliderDBID_SliderDB[sliderDB.ID] = sliderDB
		BackRepoSliderid_atBckpTime_newID[sliderDB_ID_atBackupTime] = sliderDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "SliderDB.json" in dirPath that stores an array
// of SliderDB and stores it in the database
// the map BackRepoSliderid_atBckpTime_newID is updated accordingly
func (backRepoSlider *BackRepoSliderStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoSliderid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "SliderDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Slider file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*SliderDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_SliderDBID_SliderDB
	for _, sliderDB := range forRestore {

		sliderDB_ID_atBackupTime := sliderDB.ID
		sliderDB.ID = 0
		_, err := backRepoSlider.db.Create(sliderDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoSlider.Map_SliderDBID_SliderDB[sliderDB.ID] = sliderDB
		BackRepoSliderid_atBckpTime_newID[sliderDB_ID_atBackupTime] = sliderDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Slider file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Slider>id_atBckpTime_newID
// to compute new index
func (backRepoSlider *BackRepoSliderStruct) RestorePhaseTwo() {

	for _, sliderDB := range backRepoSlider.Map_SliderDBID_SliderDB {

		// next line of code is to avert unused variable compilation error
		_ = sliderDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoSlider.db.Model(sliderDB)
		_, err := db.Updates(*sliderDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoSlider.ResetReversePointers commits all staged instances of Slider to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSlider *BackRepoSliderStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, slider := range backRepoSlider.Map_SliderDBID_SliderPtr {
		backRepoSlider.ResetReversePointersInstance(backRepo, idx, slider)
	}

	return
}

func (backRepoSlider *BackRepoSliderStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, slider *models.Slider) (Error error) {

	// fetch matching sliderDB
	if sliderDB, ok := backRepoSlider.Map_SliderDBID_SliderDB[idx]; ok {
		_ = sliderDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoSliderid_atBckpTime_newID map[uint]uint
