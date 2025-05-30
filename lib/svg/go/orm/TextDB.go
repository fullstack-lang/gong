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
var dummy_Text_sql sql.NullBool
var dummy_Text_time time.Duration
var dummy_Text_sort sort.Float64Slice

// TextAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model textAPI
type TextAPI struct {
	gorm.Model

	models.Text_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	TextPointersEncoding TextPointersEncoding
}

// TextPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type TextPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field Animates is a slice of pointers to another Struct (optional or 0..1)
	Animates IntSlice `gorm:"type:TEXT"`
}

// TextDB describes a text in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model textDB
type TextDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field textDB.Name
	Name_Data sql.NullString

	// Declation for basic field textDB.X
	X_Data sql.NullFloat64

	// Declation for basic field textDB.Y
	Y_Data sql.NullFloat64

	// Declation for basic field textDB.Content
	Content_Data sql.NullString

	// Declation for basic field textDB.Color
	Color_Data sql.NullString

	// Declation for basic field textDB.FillOpacity
	FillOpacity_Data sql.NullFloat64

	// Declation for basic field textDB.Stroke
	Stroke_Data sql.NullString

	// Declation for basic field textDB.StrokeOpacity
	StrokeOpacity_Data sql.NullFloat64

	// Declation for basic field textDB.StrokeWidth
	StrokeWidth_Data sql.NullFloat64

	// Declation for basic field textDB.StrokeDashArray
	StrokeDashArray_Data sql.NullString

	// Declation for basic field textDB.StrokeDashArrayWhenSelected
	StrokeDashArrayWhenSelected_Data sql.NullString

	// Declation for basic field textDB.Transform
	Transform_Data sql.NullString

	// Declation for basic field textDB.FontWeight
	FontWeight_Data sql.NullString

	// Declation for basic field textDB.FontSize
	FontSize_Data sql.NullString

	// Declation for basic field textDB.FontStyle
	FontStyle_Data sql.NullString

	// Declation for basic field textDB.LetterSpacing
	LetterSpacing_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	TextPointersEncoding
}

// TextDBs arrays textDBs
// swagger:response textDBsResponse
type TextDBs []TextDB

// TextDBResponse provides response
// swagger:response textDBResponse
type TextDBResponse struct {
	TextDB
}

// TextWOP is a Text without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type TextWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	X float64 `xlsx:"2"`

	Y float64 `xlsx:"3"`

	Content string `xlsx:"4"`

	Color string `xlsx:"5"`

	FillOpacity float64 `xlsx:"6"`

	Stroke string `xlsx:"7"`

	StrokeOpacity float64 `xlsx:"8"`

	StrokeWidth float64 `xlsx:"9"`

	StrokeDashArray string `xlsx:"10"`

	StrokeDashArrayWhenSelected string `xlsx:"11"`

	Transform string `xlsx:"12"`

	FontWeight string `xlsx:"13"`

	FontSize string `xlsx:"14"`

	FontStyle string `xlsx:"15"`

	LetterSpacing string `xlsx:"16"`
	// insertion for WOP pointer fields
}

var Text_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"X",
	"Y",
	"Content",
	"Color",
	"FillOpacity",
	"Stroke",
	"StrokeOpacity",
	"StrokeWidth",
	"StrokeDashArray",
	"StrokeDashArrayWhenSelected",
	"Transform",
	"FontWeight",
	"FontSize",
	"FontStyle",
	"LetterSpacing",
}

type BackRepoTextStruct struct {
	// stores TextDB according to their gorm ID
	Map_TextDBID_TextDB map[uint]*TextDB

	// stores TextDB ID according to Text address
	Map_TextPtr_TextDBID map[*models.Text]uint

	// stores Text according to their gorm ID
	Map_TextDBID_TextPtr map[uint]*models.Text

	db db.DBInterface

	stage *models.Stage
}

func (backRepoText *BackRepoTextStruct) GetStage() (stage *models.Stage) {
	stage = backRepoText.stage
	return
}

func (backRepoText *BackRepoTextStruct) GetDB() db.DBInterface {
	return backRepoText.db
}

// GetTextDBFromTextPtr is a handy function to access the back repo instance from the stage instance
func (backRepoText *BackRepoTextStruct) GetTextDBFromTextPtr(text *models.Text) (textDB *TextDB) {
	id := backRepoText.Map_TextPtr_TextDBID[text]
	textDB = backRepoText.Map_TextDBID_TextDB[id]
	return
}

// BackRepoText.CommitPhaseOne commits all staged instances of Text to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoText *BackRepoTextStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var texts []*models.Text
	for text := range stage.Texts {
		texts = append(texts, text)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(texts, func(i, j int) bool {
		return stage.TextMap_Staged_Order[texts[i]] < stage.TextMap_Staged_Order[texts[j]]
	})

	for _, text := range texts {
		backRepoText.CommitPhaseOneInstance(text)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, text := range backRepoText.Map_TextDBID_TextPtr {
		if _, ok := stage.Texts[text]; !ok {
			backRepoText.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoText.CommitDeleteInstance commits deletion of Text to the BackRepo
func (backRepoText *BackRepoTextStruct) CommitDeleteInstance(id uint) (Error error) {

	text := backRepoText.Map_TextDBID_TextPtr[id]

	// text is not staged anymore, remove textDB
	textDB := backRepoText.Map_TextDBID_TextDB[id]
	db, _ := backRepoText.db.Unscoped()
	_, err := db.Delete(textDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoText.Map_TextPtr_TextDBID, text)
	delete(backRepoText.Map_TextDBID_TextPtr, id)
	delete(backRepoText.Map_TextDBID_TextDB, id)

	return
}

// BackRepoText.CommitPhaseOneInstance commits text staged instances of Text to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoText *BackRepoTextStruct) CommitPhaseOneInstance(text *models.Text) (Error error) {

	// check if the text is not commited yet
	if _, ok := backRepoText.Map_TextPtr_TextDBID[text]; ok {
		return
	}

	// initiate text
	var textDB TextDB
	textDB.CopyBasicFieldsFromText(text)

	_, err := backRepoText.db.Create(&textDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoText.Map_TextPtr_TextDBID[text] = textDB.ID
	backRepoText.Map_TextDBID_TextPtr[textDB.ID] = text
	backRepoText.Map_TextDBID_TextDB[textDB.ID] = &textDB

	return
}

// BackRepoText.CommitPhaseTwo commits all staged instances of Text to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoText *BackRepoTextStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, text := range backRepoText.Map_TextDBID_TextPtr {
		backRepoText.CommitPhaseTwoInstance(backRepo, idx, text)
	}

	return
}

// BackRepoText.CommitPhaseTwoInstance commits {{structname }} of models.Text to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoText *BackRepoTextStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, text *models.Text) (Error error) {

	// fetch matching textDB
	if textDB, ok := backRepoText.Map_TextDBID_TextDB[idx]; ok {

		textDB.CopyBasicFieldsFromText(text)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		textDB.TextPointersEncoding.Animates = make([]int, 0)
		// 2. encode
		for _, animateAssocEnd := range text.Animates {
			animateAssocEnd_DB :=
				backRepo.BackRepoAnimate.GetAnimateDBFromAnimatePtr(animateAssocEnd)
			
			// the stage might be inconsistant, meaning that the animateAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if animateAssocEnd_DB == nil {
				continue
			}
			
			textDB.TextPointersEncoding.Animates =
				append(textDB.TextPointersEncoding.Animates, int(animateAssocEnd_DB.ID))
		}

		_, err := backRepoText.db.Save(textDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Text intance %s", text.Name))
		return err
	}

	return
}

// BackRepoText.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoText *BackRepoTextStruct) CheckoutPhaseOne() (Error error) {

	textDBArray := make([]TextDB, 0)
	_, err := backRepoText.db.Find(&textDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	textInstancesToBeRemovedFromTheStage := make(map[*models.Text]any)
	for key, value := range backRepoText.stage.Texts {
		textInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, textDB := range textDBArray {
		backRepoText.CheckoutPhaseOneInstance(&textDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		text, ok := backRepoText.Map_TextDBID_TextPtr[textDB.ID]
		if ok {
			delete(textInstancesToBeRemovedFromTheStage, text)
		}
	}

	// remove from stage and back repo's 3 maps all texts that are not in the checkout
	for text := range textInstancesToBeRemovedFromTheStage {
		text.Unstage(backRepoText.GetStage())

		// remove instance from the back repo 3 maps
		textID := backRepoText.Map_TextPtr_TextDBID[text]
		delete(backRepoText.Map_TextPtr_TextDBID, text)
		delete(backRepoText.Map_TextDBID_TextDB, textID)
		delete(backRepoText.Map_TextDBID_TextPtr, textID)
	}

	return
}

// CheckoutPhaseOneInstance takes a textDB that has been found in the DB, updates the backRepo and stages the
// models version of the textDB
func (backRepoText *BackRepoTextStruct) CheckoutPhaseOneInstance(textDB *TextDB) (Error error) {

	text, ok := backRepoText.Map_TextDBID_TextPtr[textDB.ID]
	if !ok {
		text = new(models.Text)

		backRepoText.Map_TextDBID_TextPtr[textDB.ID] = text
		backRepoText.Map_TextPtr_TextDBID[text] = textDB.ID

		// append model store with the new element
		text.Name = textDB.Name_Data.String
		text.Stage(backRepoText.GetStage())
	}
	textDB.CopyBasicFieldsToText(text)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	text.Stage(backRepoText.GetStage())

	// preserve pointer to textDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_TextDBID_TextDB)[textDB hold variable pointers
	textDB_Data := *textDB
	preservedPtrToText := &textDB_Data
	backRepoText.Map_TextDBID_TextDB[textDB.ID] = preservedPtrToText

	return
}

// BackRepoText.CheckoutPhaseTwo Checkouts all staged instances of Text to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoText *BackRepoTextStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, textDB := range backRepoText.Map_TextDBID_TextDB {
		backRepoText.CheckoutPhaseTwoInstance(backRepo, textDB)
	}
	return
}

// BackRepoText.CheckoutPhaseTwoInstance Checkouts staged instances of Text to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoText *BackRepoTextStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, textDB *TextDB) (Error error) {

	text := backRepoText.Map_TextDBID_TextPtr[textDB.ID]

	textDB.DecodePointers(backRepo, text)

	return
}

func (textDB *TextDB) DecodePointers(backRepo *BackRepoStruct, text *models.Text) {

	// insertion point for checkout of pointer encoding
	// This loop redeem text.Animates in the stage from the encode in the back repo
	// It parses all AnimateDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	text.Animates = text.Animates[:0]
	for _, _Animateid := range textDB.TextPointersEncoding.Animates {
		text.Animates = append(text.Animates, backRepo.BackRepoAnimate.Map_AnimateDBID_AnimatePtr[uint(_Animateid)])
	}

	return
}

// CommitText allows commit of a single text (if already staged)
func (backRepo *BackRepoStruct) CommitText(text *models.Text) {
	backRepo.BackRepoText.CommitPhaseOneInstance(text)
	if id, ok := backRepo.BackRepoText.Map_TextPtr_TextDBID[text]; ok {
		backRepo.BackRepoText.CommitPhaseTwoInstance(backRepo, id, text)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitText allows checkout of a single text (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutText(text *models.Text) {
	// check if the text is staged
	if _, ok := backRepo.BackRepoText.Map_TextPtr_TextDBID[text]; ok {

		if id, ok := backRepo.BackRepoText.Map_TextPtr_TextDBID[text]; ok {
			var textDB TextDB
			textDB.ID = id

			if _, err := backRepo.BackRepoText.db.First(&textDB, id); err != nil {
				log.Fatalln("CheckoutText : Problem with getting object with id:", id)
			}
			backRepo.BackRepoText.CheckoutPhaseOneInstance(&textDB)
			backRepo.BackRepoText.CheckoutPhaseTwoInstance(backRepo, &textDB)
		}
	}
}

// CopyBasicFieldsFromText
func (textDB *TextDB) CopyBasicFieldsFromText(text *models.Text) {
	// insertion point for fields commit

	textDB.Name_Data.String = text.Name
	textDB.Name_Data.Valid = true

	textDB.X_Data.Float64 = text.X
	textDB.X_Data.Valid = true

	textDB.Y_Data.Float64 = text.Y
	textDB.Y_Data.Valid = true

	textDB.Content_Data.String = text.Content
	textDB.Content_Data.Valid = true

	textDB.Color_Data.String = text.Color
	textDB.Color_Data.Valid = true

	textDB.FillOpacity_Data.Float64 = text.FillOpacity
	textDB.FillOpacity_Data.Valid = true

	textDB.Stroke_Data.String = text.Stroke
	textDB.Stroke_Data.Valid = true

	textDB.StrokeOpacity_Data.Float64 = text.StrokeOpacity
	textDB.StrokeOpacity_Data.Valid = true

	textDB.StrokeWidth_Data.Float64 = text.StrokeWidth
	textDB.StrokeWidth_Data.Valid = true

	textDB.StrokeDashArray_Data.String = text.StrokeDashArray
	textDB.StrokeDashArray_Data.Valid = true

	textDB.StrokeDashArrayWhenSelected_Data.String = text.StrokeDashArrayWhenSelected
	textDB.StrokeDashArrayWhenSelected_Data.Valid = true

	textDB.Transform_Data.String = text.Transform
	textDB.Transform_Data.Valid = true

	textDB.FontWeight_Data.String = text.FontWeight
	textDB.FontWeight_Data.Valid = true

	textDB.FontSize_Data.String = text.FontSize
	textDB.FontSize_Data.Valid = true

	textDB.FontStyle_Data.String = text.FontStyle
	textDB.FontStyle_Data.Valid = true

	textDB.LetterSpacing_Data.String = text.LetterSpacing
	textDB.LetterSpacing_Data.Valid = true
}

// CopyBasicFieldsFromText_WOP
func (textDB *TextDB) CopyBasicFieldsFromText_WOP(text *models.Text_WOP) {
	// insertion point for fields commit

	textDB.Name_Data.String = text.Name
	textDB.Name_Data.Valid = true

	textDB.X_Data.Float64 = text.X
	textDB.X_Data.Valid = true

	textDB.Y_Data.Float64 = text.Y
	textDB.Y_Data.Valid = true

	textDB.Content_Data.String = text.Content
	textDB.Content_Data.Valid = true

	textDB.Color_Data.String = text.Color
	textDB.Color_Data.Valid = true

	textDB.FillOpacity_Data.Float64 = text.FillOpacity
	textDB.FillOpacity_Data.Valid = true

	textDB.Stroke_Data.String = text.Stroke
	textDB.Stroke_Data.Valid = true

	textDB.StrokeOpacity_Data.Float64 = text.StrokeOpacity
	textDB.StrokeOpacity_Data.Valid = true

	textDB.StrokeWidth_Data.Float64 = text.StrokeWidth
	textDB.StrokeWidth_Data.Valid = true

	textDB.StrokeDashArray_Data.String = text.StrokeDashArray
	textDB.StrokeDashArray_Data.Valid = true

	textDB.StrokeDashArrayWhenSelected_Data.String = text.StrokeDashArrayWhenSelected
	textDB.StrokeDashArrayWhenSelected_Data.Valid = true

	textDB.Transform_Data.String = text.Transform
	textDB.Transform_Data.Valid = true

	textDB.FontWeight_Data.String = text.FontWeight
	textDB.FontWeight_Data.Valid = true

	textDB.FontSize_Data.String = text.FontSize
	textDB.FontSize_Data.Valid = true

	textDB.FontStyle_Data.String = text.FontStyle
	textDB.FontStyle_Data.Valid = true

	textDB.LetterSpacing_Data.String = text.LetterSpacing
	textDB.LetterSpacing_Data.Valid = true
}

// CopyBasicFieldsFromTextWOP
func (textDB *TextDB) CopyBasicFieldsFromTextWOP(text *TextWOP) {
	// insertion point for fields commit

	textDB.Name_Data.String = text.Name
	textDB.Name_Data.Valid = true

	textDB.X_Data.Float64 = text.X
	textDB.X_Data.Valid = true

	textDB.Y_Data.Float64 = text.Y
	textDB.Y_Data.Valid = true

	textDB.Content_Data.String = text.Content
	textDB.Content_Data.Valid = true

	textDB.Color_Data.String = text.Color
	textDB.Color_Data.Valid = true

	textDB.FillOpacity_Data.Float64 = text.FillOpacity
	textDB.FillOpacity_Data.Valid = true

	textDB.Stroke_Data.String = text.Stroke
	textDB.Stroke_Data.Valid = true

	textDB.StrokeOpacity_Data.Float64 = text.StrokeOpacity
	textDB.StrokeOpacity_Data.Valid = true

	textDB.StrokeWidth_Data.Float64 = text.StrokeWidth
	textDB.StrokeWidth_Data.Valid = true

	textDB.StrokeDashArray_Data.String = text.StrokeDashArray
	textDB.StrokeDashArray_Data.Valid = true

	textDB.StrokeDashArrayWhenSelected_Data.String = text.StrokeDashArrayWhenSelected
	textDB.StrokeDashArrayWhenSelected_Data.Valid = true

	textDB.Transform_Data.String = text.Transform
	textDB.Transform_Data.Valid = true

	textDB.FontWeight_Data.String = text.FontWeight
	textDB.FontWeight_Data.Valid = true

	textDB.FontSize_Data.String = text.FontSize
	textDB.FontSize_Data.Valid = true

	textDB.FontStyle_Data.String = text.FontStyle
	textDB.FontStyle_Data.Valid = true

	textDB.LetterSpacing_Data.String = text.LetterSpacing
	textDB.LetterSpacing_Data.Valid = true
}

// CopyBasicFieldsToText
func (textDB *TextDB) CopyBasicFieldsToText(text *models.Text) {
	// insertion point for checkout of basic fields (back repo to stage)
	text.Name = textDB.Name_Data.String
	text.X = textDB.X_Data.Float64
	text.Y = textDB.Y_Data.Float64
	text.Content = textDB.Content_Data.String
	text.Color = textDB.Color_Data.String
	text.FillOpacity = textDB.FillOpacity_Data.Float64
	text.Stroke = textDB.Stroke_Data.String
	text.StrokeOpacity = textDB.StrokeOpacity_Data.Float64
	text.StrokeWidth = textDB.StrokeWidth_Data.Float64
	text.StrokeDashArray = textDB.StrokeDashArray_Data.String
	text.StrokeDashArrayWhenSelected = textDB.StrokeDashArrayWhenSelected_Data.String
	text.Transform = textDB.Transform_Data.String
	text.FontWeight = textDB.FontWeight_Data.String
	text.FontSize = textDB.FontSize_Data.String
	text.FontStyle = textDB.FontStyle_Data.String
	text.LetterSpacing = textDB.LetterSpacing_Data.String
}

// CopyBasicFieldsToText_WOP
func (textDB *TextDB) CopyBasicFieldsToText_WOP(text *models.Text_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	text.Name = textDB.Name_Data.String
	text.X = textDB.X_Data.Float64
	text.Y = textDB.Y_Data.Float64
	text.Content = textDB.Content_Data.String
	text.Color = textDB.Color_Data.String
	text.FillOpacity = textDB.FillOpacity_Data.Float64
	text.Stroke = textDB.Stroke_Data.String
	text.StrokeOpacity = textDB.StrokeOpacity_Data.Float64
	text.StrokeWidth = textDB.StrokeWidth_Data.Float64
	text.StrokeDashArray = textDB.StrokeDashArray_Data.String
	text.StrokeDashArrayWhenSelected = textDB.StrokeDashArrayWhenSelected_Data.String
	text.Transform = textDB.Transform_Data.String
	text.FontWeight = textDB.FontWeight_Data.String
	text.FontSize = textDB.FontSize_Data.String
	text.FontStyle = textDB.FontStyle_Data.String
	text.LetterSpacing = textDB.LetterSpacing_Data.String
}

// CopyBasicFieldsToTextWOP
func (textDB *TextDB) CopyBasicFieldsToTextWOP(text *TextWOP) {
	text.ID = int(textDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	text.Name = textDB.Name_Data.String
	text.X = textDB.X_Data.Float64
	text.Y = textDB.Y_Data.Float64
	text.Content = textDB.Content_Data.String
	text.Color = textDB.Color_Data.String
	text.FillOpacity = textDB.FillOpacity_Data.Float64
	text.Stroke = textDB.Stroke_Data.String
	text.StrokeOpacity = textDB.StrokeOpacity_Data.Float64
	text.StrokeWidth = textDB.StrokeWidth_Data.Float64
	text.StrokeDashArray = textDB.StrokeDashArray_Data.String
	text.StrokeDashArrayWhenSelected = textDB.StrokeDashArrayWhenSelected_Data.String
	text.Transform = textDB.Transform_Data.String
	text.FontWeight = textDB.FontWeight_Data.String
	text.FontSize = textDB.FontSize_Data.String
	text.FontStyle = textDB.FontStyle_Data.String
	text.LetterSpacing = textDB.LetterSpacing_Data.String
}

// Backup generates a json file from a slice of all TextDB instances in the backrepo
func (backRepoText *BackRepoTextStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "TextDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*TextDB, 0)
	for _, textDB := range backRepoText.Map_TextDBID_TextDB {
		forBackup = append(forBackup, textDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Text ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Text file", err.Error())
	}
}

// Backup generates a json file from a slice of all TextDB instances in the backrepo
func (backRepoText *BackRepoTextStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*TextDB, 0)
	for _, textDB := range backRepoText.Map_TextDBID_TextDB {
		forBackup = append(forBackup, textDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Text")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Text_Fields, -1)
	for _, textDB := range forBackup {

		var textWOP TextWOP
		textDB.CopyBasicFieldsToTextWOP(&textWOP)

		row := sh.AddRow()
		row.WriteStruct(&textWOP, -1)
	}
}

// RestoreXL from the "Text" sheet all TextDB instances
func (backRepoText *BackRepoTextStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoTextid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Text"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoText.rowVisitorText)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoText *BackRepoTextStruct) rowVisitorText(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var textWOP TextWOP
		row.ReadStruct(&textWOP)

		// add the unmarshalled struct to the stage
		textDB := new(TextDB)
		textDB.CopyBasicFieldsFromTextWOP(&textWOP)

		textDB_ID_atBackupTime := textDB.ID
		textDB.ID = 0
		_, err := backRepoText.db.Create(textDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoText.Map_TextDBID_TextDB[textDB.ID] = textDB
		BackRepoTextid_atBckpTime_newID[textDB_ID_atBackupTime] = textDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "TextDB.json" in dirPath that stores an array
// of TextDB and stores it in the database
// the map BackRepoTextid_atBckpTime_newID is updated accordingly
func (backRepoText *BackRepoTextStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoTextid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "TextDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Text file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*TextDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_TextDBID_TextDB
	for _, textDB := range forRestore {

		textDB_ID_atBackupTime := textDB.ID
		textDB.ID = 0
		_, err := backRepoText.db.Create(textDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoText.Map_TextDBID_TextDB[textDB.ID] = textDB
		BackRepoTextid_atBckpTime_newID[textDB_ID_atBackupTime] = textDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Text file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Text>id_atBckpTime_newID
// to compute new index
func (backRepoText *BackRepoTextStruct) RestorePhaseTwo() {

	for _, textDB := range backRepoText.Map_TextDBID_TextDB {

		// next line of code is to avert unused variable compilation error
		_ = textDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoText.db.Model(textDB)
		_, err := db.Updates(*textDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoText.ResetReversePointers commits all staged instances of Text to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoText *BackRepoTextStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, text := range backRepoText.Map_TextDBID_TextPtr {
		backRepoText.ResetReversePointersInstance(backRepo, idx, text)
	}

	return
}

func (backRepoText *BackRepoTextStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, text *models.Text) (Error error) {

	// fetch matching textDB
	if textDB, ok := backRepoText.Map_TextDBID_TextDB[idx]; ok {
		_ = textDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoTextid_atBckpTime_newID map[uint]uint
