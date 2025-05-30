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

	"github.com/fullstack-lang/gong/lib/doc2/go/db"
	"github.com/fullstack-lang/gong/lib/doc2/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Classdiagram_sql sql.NullBool
var dummy_Classdiagram_time time.Duration
var dummy_Classdiagram_sort sort.Float64Slice

// ClassdiagramAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model classdiagramAPI
type ClassdiagramAPI struct {
	gorm.Model

	models.Classdiagram_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	ClassdiagramPointersEncoding ClassdiagramPointersEncoding
}

// ClassdiagramPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type ClassdiagramPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field GongStructShapes is a slice of pointers to another Struct (optional or 0..1)
	GongStructShapes IntSlice `gorm:"type:TEXT"`

	// field GongEnumShapes is a slice of pointers to another Struct (optional or 0..1)
	GongEnumShapes IntSlice `gorm:"type:TEXT"`

	// field GongNoteShapes is a slice of pointers to another Struct (optional or 0..1)
	GongNoteShapes IntSlice `gorm:"type:TEXT"`
}

// ClassdiagramDB describes a classdiagram in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model classdiagramDB
type ClassdiagramDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field classdiagramDB.Name
	Name_Data sql.NullString

	// Declation for basic field classdiagramDB.Description
	Description_Data sql.NullString

	// Declation for basic field classdiagramDB.IsIncludedInStaticWebSite
	// provide the sql storage for the boolan
	IsIncludedInStaticWebSite_Data sql.NullBool

	// Declation for basic field classdiagramDB.IsInRenameMode
	// provide the sql storage for the boolan
	IsInRenameMode_Data sql.NullBool

	// Declation for basic field classdiagramDB.IsExpanded
	// provide the sql storage for the boolan
	IsExpanded_Data sql.NullBool

	// Declation for basic field classdiagramDB.NodeGongStructsIsExpanded
	// provide the sql storage for the boolan
	NodeGongStructsIsExpanded_Data sql.NullBool

	// Declation for basic field classdiagramDB.NodeGongStructNodeExpansion
	NodeGongStructNodeExpansion_Data sql.NullString

	// Declation for basic field classdiagramDB.NodeGongEnumsIsExpanded
	// provide the sql storage for the boolan
	NodeGongEnumsIsExpanded_Data sql.NullBool

	// Declation for basic field classdiagramDB.NodeGongEnumNodeExpansion
	NodeGongEnumNodeExpansion_Data sql.NullString

	// Declation for basic field classdiagramDB.NodeGongNotesIsExpanded
	// provide the sql storage for the boolan
	NodeGongNotesIsExpanded_Data sql.NullBool

	// Declation for basic field classdiagramDB.NodeGongNoteNodeExpansion
	NodeGongNoteNodeExpansion_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	ClassdiagramPointersEncoding
}

// ClassdiagramDBs arrays classdiagramDBs
// swagger:response classdiagramDBsResponse
type ClassdiagramDBs []ClassdiagramDB

// ClassdiagramDBResponse provides response
// swagger:response classdiagramDBResponse
type ClassdiagramDBResponse struct {
	ClassdiagramDB
}

// ClassdiagramWOP is a Classdiagram without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type ClassdiagramWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Description string `xlsx:"2"`

	IsIncludedInStaticWebSite bool `xlsx:"3"`

	IsInRenameMode bool `xlsx:"4"`

	IsExpanded bool `xlsx:"5"`

	NodeGongStructsIsExpanded bool `xlsx:"6"`

	NodeGongStructNodeExpansion string `xlsx:"7"`

	NodeGongEnumsIsExpanded bool `xlsx:"8"`

	NodeGongEnumNodeExpansion string `xlsx:"9"`

	NodeGongNotesIsExpanded bool `xlsx:"10"`

	NodeGongNoteNodeExpansion string `xlsx:"11"`
	// insertion for WOP pointer fields
}

var Classdiagram_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Description",
	"IsIncludedInStaticWebSite",
	"IsInRenameMode",
	"IsExpanded",
	"NodeGongStructsIsExpanded",
	"NodeGongStructNodeExpansion",
	"NodeGongEnumsIsExpanded",
	"NodeGongEnumNodeExpansion",
	"NodeGongNotesIsExpanded",
	"NodeGongNoteNodeExpansion",
}

type BackRepoClassdiagramStruct struct {
	// stores ClassdiagramDB according to their gorm ID
	Map_ClassdiagramDBID_ClassdiagramDB map[uint]*ClassdiagramDB

	// stores ClassdiagramDB ID according to Classdiagram address
	Map_ClassdiagramPtr_ClassdiagramDBID map[*models.Classdiagram]uint

	// stores Classdiagram according to their gorm ID
	Map_ClassdiagramDBID_ClassdiagramPtr map[uint]*models.Classdiagram

	db db.DBInterface

	stage *models.Stage
}

func (backRepoClassdiagram *BackRepoClassdiagramStruct) GetStage() (stage *models.Stage) {
	stage = backRepoClassdiagram.stage
	return
}

func (backRepoClassdiagram *BackRepoClassdiagramStruct) GetDB() db.DBInterface {
	return backRepoClassdiagram.db
}

// GetClassdiagramDBFromClassdiagramPtr is a handy function to access the back repo instance from the stage instance
func (backRepoClassdiagram *BackRepoClassdiagramStruct) GetClassdiagramDBFromClassdiagramPtr(classdiagram *models.Classdiagram) (classdiagramDB *ClassdiagramDB) {
	id := backRepoClassdiagram.Map_ClassdiagramPtr_ClassdiagramDBID[classdiagram]
	classdiagramDB = backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramDB[id]
	return
}

// BackRepoClassdiagram.CommitPhaseOne commits all staged instances of Classdiagram to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoClassdiagram *BackRepoClassdiagramStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var classdiagrams []*models.Classdiagram
	for classdiagram := range stage.Classdiagrams {
		classdiagrams = append(classdiagrams, classdiagram)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(classdiagrams, func(i, j int) bool {
		return stage.ClassdiagramMap_Staged_Order[classdiagrams[i]] < stage.ClassdiagramMap_Staged_Order[classdiagrams[j]]
	})

	for _, classdiagram := range classdiagrams {
		backRepoClassdiagram.CommitPhaseOneInstance(classdiagram)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, classdiagram := range backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr {
		if _, ok := stage.Classdiagrams[classdiagram]; !ok {
			backRepoClassdiagram.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoClassdiagram.CommitDeleteInstance commits deletion of Classdiagram to the BackRepo
func (backRepoClassdiagram *BackRepoClassdiagramStruct) CommitDeleteInstance(id uint) (Error error) {

	classdiagram := backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr[id]

	// classdiagram is not staged anymore, remove classdiagramDB
	classdiagramDB := backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramDB[id]
	db, _ := backRepoClassdiagram.db.Unscoped()
	_, err := db.Delete(classdiagramDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoClassdiagram.Map_ClassdiagramPtr_ClassdiagramDBID, classdiagram)
	delete(backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr, id)
	delete(backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramDB, id)

	return
}

// BackRepoClassdiagram.CommitPhaseOneInstance commits classdiagram staged instances of Classdiagram to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoClassdiagram *BackRepoClassdiagramStruct) CommitPhaseOneInstance(classdiagram *models.Classdiagram) (Error error) {

	// check if the classdiagram is not commited yet
	if _, ok := backRepoClassdiagram.Map_ClassdiagramPtr_ClassdiagramDBID[classdiagram]; ok {
		return
	}

	// initiate classdiagram
	var classdiagramDB ClassdiagramDB
	classdiagramDB.CopyBasicFieldsFromClassdiagram(classdiagram)

	_, err := backRepoClassdiagram.db.Create(&classdiagramDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoClassdiagram.Map_ClassdiagramPtr_ClassdiagramDBID[classdiagram] = classdiagramDB.ID
	backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr[classdiagramDB.ID] = classdiagram
	backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramDB[classdiagramDB.ID] = &classdiagramDB

	return
}

// BackRepoClassdiagram.CommitPhaseTwo commits all staged instances of Classdiagram to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoClassdiagram *BackRepoClassdiagramStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, classdiagram := range backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr {
		backRepoClassdiagram.CommitPhaseTwoInstance(backRepo, idx, classdiagram)
	}

	return
}

// BackRepoClassdiagram.CommitPhaseTwoInstance commits {{structname }} of models.Classdiagram to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoClassdiagram *BackRepoClassdiagramStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, classdiagram *models.Classdiagram) (Error error) {

	// fetch matching classdiagramDB
	if classdiagramDB, ok := backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramDB[idx]; ok {

		classdiagramDB.CopyBasicFieldsFromClassdiagram(classdiagram)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		classdiagramDB.ClassdiagramPointersEncoding.GongStructShapes = make([]int, 0)
		// 2. encode
		for _, gongstructshapeAssocEnd := range classdiagram.GongStructShapes {
			gongstructshapeAssocEnd_DB :=
				backRepo.BackRepoGongStructShape.GetGongStructShapeDBFromGongStructShapePtr(gongstructshapeAssocEnd)
			
			// the stage might be inconsistant, meaning that the gongstructshapeAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if gongstructshapeAssocEnd_DB == nil {
				continue
			}
			
			classdiagramDB.ClassdiagramPointersEncoding.GongStructShapes =
				append(classdiagramDB.ClassdiagramPointersEncoding.GongStructShapes, int(gongstructshapeAssocEnd_DB.ID))
		}

		// 1. reset
		classdiagramDB.ClassdiagramPointersEncoding.GongEnumShapes = make([]int, 0)
		// 2. encode
		for _, gongenumshapeAssocEnd := range classdiagram.GongEnumShapes {
			gongenumshapeAssocEnd_DB :=
				backRepo.BackRepoGongEnumShape.GetGongEnumShapeDBFromGongEnumShapePtr(gongenumshapeAssocEnd)
			
			// the stage might be inconsistant, meaning that the gongenumshapeAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if gongenumshapeAssocEnd_DB == nil {
				continue
			}
			
			classdiagramDB.ClassdiagramPointersEncoding.GongEnumShapes =
				append(classdiagramDB.ClassdiagramPointersEncoding.GongEnumShapes, int(gongenumshapeAssocEnd_DB.ID))
		}

		// 1. reset
		classdiagramDB.ClassdiagramPointersEncoding.GongNoteShapes = make([]int, 0)
		// 2. encode
		for _, gongnoteshapeAssocEnd := range classdiagram.GongNoteShapes {
			gongnoteshapeAssocEnd_DB :=
				backRepo.BackRepoGongNoteShape.GetGongNoteShapeDBFromGongNoteShapePtr(gongnoteshapeAssocEnd)
			
			// the stage might be inconsistant, meaning that the gongnoteshapeAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if gongnoteshapeAssocEnd_DB == nil {
				continue
			}
			
			classdiagramDB.ClassdiagramPointersEncoding.GongNoteShapes =
				append(classdiagramDB.ClassdiagramPointersEncoding.GongNoteShapes, int(gongnoteshapeAssocEnd_DB.ID))
		}

		_, err := backRepoClassdiagram.db.Save(classdiagramDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Classdiagram intance %s", classdiagram.Name))
		return err
	}

	return
}

// BackRepoClassdiagram.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoClassdiagram *BackRepoClassdiagramStruct) CheckoutPhaseOne() (Error error) {

	classdiagramDBArray := make([]ClassdiagramDB, 0)
	_, err := backRepoClassdiagram.db.Find(&classdiagramDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	classdiagramInstancesToBeRemovedFromTheStage := make(map[*models.Classdiagram]any)
	for key, value := range backRepoClassdiagram.stage.Classdiagrams {
		classdiagramInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, classdiagramDB := range classdiagramDBArray {
		backRepoClassdiagram.CheckoutPhaseOneInstance(&classdiagramDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		classdiagram, ok := backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr[classdiagramDB.ID]
		if ok {
			delete(classdiagramInstancesToBeRemovedFromTheStage, classdiagram)
		}
	}

	// remove from stage and back repo's 3 maps all classdiagrams that are not in the checkout
	for classdiagram := range classdiagramInstancesToBeRemovedFromTheStage {
		classdiagram.Unstage(backRepoClassdiagram.GetStage())

		// remove instance from the back repo 3 maps
		classdiagramID := backRepoClassdiagram.Map_ClassdiagramPtr_ClassdiagramDBID[classdiagram]
		delete(backRepoClassdiagram.Map_ClassdiagramPtr_ClassdiagramDBID, classdiagram)
		delete(backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramDB, classdiagramID)
		delete(backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr, classdiagramID)
	}

	return
}

// CheckoutPhaseOneInstance takes a classdiagramDB that has been found in the DB, updates the backRepo and stages the
// models version of the classdiagramDB
func (backRepoClassdiagram *BackRepoClassdiagramStruct) CheckoutPhaseOneInstance(classdiagramDB *ClassdiagramDB) (Error error) {

	classdiagram, ok := backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr[classdiagramDB.ID]
	if !ok {
		classdiagram = new(models.Classdiagram)

		backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr[classdiagramDB.ID] = classdiagram
		backRepoClassdiagram.Map_ClassdiagramPtr_ClassdiagramDBID[classdiagram] = classdiagramDB.ID

		// append model store with the new element
		classdiagram.Name = classdiagramDB.Name_Data.String
		classdiagram.Stage(backRepoClassdiagram.GetStage())
	}
	classdiagramDB.CopyBasicFieldsToClassdiagram(classdiagram)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	classdiagram.Stage(backRepoClassdiagram.GetStage())

	// preserve pointer to classdiagramDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_ClassdiagramDBID_ClassdiagramDB)[classdiagramDB hold variable pointers
	classdiagramDB_Data := *classdiagramDB
	preservedPtrToClassdiagram := &classdiagramDB_Data
	backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramDB[classdiagramDB.ID] = preservedPtrToClassdiagram

	return
}

// BackRepoClassdiagram.CheckoutPhaseTwo Checkouts all staged instances of Classdiagram to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoClassdiagram *BackRepoClassdiagramStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, classdiagramDB := range backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramDB {
		backRepoClassdiagram.CheckoutPhaseTwoInstance(backRepo, classdiagramDB)
	}
	return
}

// BackRepoClassdiagram.CheckoutPhaseTwoInstance Checkouts staged instances of Classdiagram to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoClassdiagram *BackRepoClassdiagramStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, classdiagramDB *ClassdiagramDB) (Error error) {

	classdiagram := backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr[classdiagramDB.ID]

	classdiagramDB.DecodePointers(backRepo, classdiagram)

	return
}

func (classdiagramDB *ClassdiagramDB) DecodePointers(backRepo *BackRepoStruct, classdiagram *models.Classdiagram) {

	// insertion point for checkout of pointer encoding
	// This loop redeem classdiagram.GongStructShapes in the stage from the encode in the back repo
	// It parses all GongStructShapeDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	classdiagram.GongStructShapes = classdiagram.GongStructShapes[:0]
	for _, _GongStructShapeid := range classdiagramDB.ClassdiagramPointersEncoding.GongStructShapes {
		classdiagram.GongStructShapes = append(classdiagram.GongStructShapes, backRepo.BackRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr[uint(_GongStructShapeid)])
	}

	// This loop redeem classdiagram.GongEnumShapes in the stage from the encode in the back repo
	// It parses all GongEnumShapeDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	classdiagram.GongEnumShapes = classdiagram.GongEnumShapes[:0]
	for _, _GongEnumShapeid := range classdiagramDB.ClassdiagramPointersEncoding.GongEnumShapes {
		classdiagram.GongEnumShapes = append(classdiagram.GongEnumShapes, backRepo.BackRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapePtr[uint(_GongEnumShapeid)])
	}

	// This loop redeem classdiagram.GongNoteShapes in the stage from the encode in the back repo
	// It parses all GongNoteShapeDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	classdiagram.GongNoteShapes = classdiagram.GongNoteShapes[:0]
	for _, _GongNoteShapeid := range classdiagramDB.ClassdiagramPointersEncoding.GongNoteShapes {
		classdiagram.GongNoteShapes = append(classdiagram.GongNoteShapes, backRepo.BackRepoGongNoteShape.Map_GongNoteShapeDBID_GongNoteShapePtr[uint(_GongNoteShapeid)])
	}

	return
}

// CommitClassdiagram allows commit of a single classdiagram (if already staged)
func (backRepo *BackRepoStruct) CommitClassdiagram(classdiagram *models.Classdiagram) {
	backRepo.BackRepoClassdiagram.CommitPhaseOneInstance(classdiagram)
	if id, ok := backRepo.BackRepoClassdiagram.Map_ClassdiagramPtr_ClassdiagramDBID[classdiagram]; ok {
		backRepo.BackRepoClassdiagram.CommitPhaseTwoInstance(backRepo, id, classdiagram)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitClassdiagram allows checkout of a single classdiagram (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutClassdiagram(classdiagram *models.Classdiagram) {
	// check if the classdiagram is staged
	if _, ok := backRepo.BackRepoClassdiagram.Map_ClassdiagramPtr_ClassdiagramDBID[classdiagram]; ok {

		if id, ok := backRepo.BackRepoClassdiagram.Map_ClassdiagramPtr_ClassdiagramDBID[classdiagram]; ok {
			var classdiagramDB ClassdiagramDB
			classdiagramDB.ID = id

			if _, err := backRepo.BackRepoClassdiagram.db.First(&classdiagramDB, id); err != nil {
				log.Fatalln("CheckoutClassdiagram : Problem with getting object with id:", id)
			}
			backRepo.BackRepoClassdiagram.CheckoutPhaseOneInstance(&classdiagramDB)
			backRepo.BackRepoClassdiagram.CheckoutPhaseTwoInstance(backRepo, &classdiagramDB)
		}
	}
}

// CopyBasicFieldsFromClassdiagram
func (classdiagramDB *ClassdiagramDB) CopyBasicFieldsFromClassdiagram(classdiagram *models.Classdiagram) {
	// insertion point for fields commit

	classdiagramDB.Name_Data.String = classdiagram.Name
	classdiagramDB.Name_Data.Valid = true

	classdiagramDB.Description_Data.String = classdiagram.Description
	classdiagramDB.Description_Data.Valid = true

	classdiagramDB.IsIncludedInStaticWebSite_Data.Bool = classdiagram.IsIncludedInStaticWebSite
	classdiagramDB.IsIncludedInStaticWebSite_Data.Valid = true

	classdiagramDB.IsInRenameMode_Data.Bool = classdiagram.IsInRenameMode
	classdiagramDB.IsInRenameMode_Data.Valid = true

	classdiagramDB.IsExpanded_Data.Bool = classdiagram.IsExpanded
	classdiagramDB.IsExpanded_Data.Valid = true

	classdiagramDB.NodeGongStructsIsExpanded_Data.Bool = classdiagram.NodeGongStructsIsExpanded
	classdiagramDB.NodeGongStructsIsExpanded_Data.Valid = true

	classdiagramDB.NodeGongStructNodeExpansion_Data.String = classdiagram.NodeGongStructNodeExpansion
	classdiagramDB.NodeGongStructNodeExpansion_Data.Valid = true

	classdiagramDB.NodeGongEnumsIsExpanded_Data.Bool = classdiagram.NodeGongEnumsIsExpanded
	classdiagramDB.NodeGongEnumsIsExpanded_Data.Valid = true

	classdiagramDB.NodeGongEnumNodeExpansion_Data.String = classdiagram.NodeGongEnumNodeExpansion
	classdiagramDB.NodeGongEnumNodeExpansion_Data.Valid = true

	classdiagramDB.NodeGongNotesIsExpanded_Data.Bool = classdiagram.NodeGongNotesIsExpanded
	classdiagramDB.NodeGongNotesIsExpanded_Data.Valid = true

	classdiagramDB.NodeGongNoteNodeExpansion_Data.String = classdiagram.NodeGongNoteNodeExpansion
	classdiagramDB.NodeGongNoteNodeExpansion_Data.Valid = true
}

// CopyBasicFieldsFromClassdiagram_WOP
func (classdiagramDB *ClassdiagramDB) CopyBasicFieldsFromClassdiagram_WOP(classdiagram *models.Classdiagram_WOP) {
	// insertion point for fields commit

	classdiagramDB.Name_Data.String = classdiagram.Name
	classdiagramDB.Name_Data.Valid = true

	classdiagramDB.Description_Data.String = classdiagram.Description
	classdiagramDB.Description_Data.Valid = true

	classdiagramDB.IsIncludedInStaticWebSite_Data.Bool = classdiagram.IsIncludedInStaticWebSite
	classdiagramDB.IsIncludedInStaticWebSite_Data.Valid = true

	classdiagramDB.IsInRenameMode_Data.Bool = classdiagram.IsInRenameMode
	classdiagramDB.IsInRenameMode_Data.Valid = true

	classdiagramDB.IsExpanded_Data.Bool = classdiagram.IsExpanded
	classdiagramDB.IsExpanded_Data.Valid = true

	classdiagramDB.NodeGongStructsIsExpanded_Data.Bool = classdiagram.NodeGongStructsIsExpanded
	classdiagramDB.NodeGongStructsIsExpanded_Data.Valid = true

	classdiagramDB.NodeGongStructNodeExpansion_Data.String = classdiagram.NodeGongStructNodeExpansion
	classdiagramDB.NodeGongStructNodeExpansion_Data.Valid = true

	classdiagramDB.NodeGongEnumsIsExpanded_Data.Bool = classdiagram.NodeGongEnumsIsExpanded
	classdiagramDB.NodeGongEnumsIsExpanded_Data.Valid = true

	classdiagramDB.NodeGongEnumNodeExpansion_Data.String = classdiagram.NodeGongEnumNodeExpansion
	classdiagramDB.NodeGongEnumNodeExpansion_Data.Valid = true

	classdiagramDB.NodeGongNotesIsExpanded_Data.Bool = classdiagram.NodeGongNotesIsExpanded
	classdiagramDB.NodeGongNotesIsExpanded_Data.Valid = true

	classdiagramDB.NodeGongNoteNodeExpansion_Data.String = classdiagram.NodeGongNoteNodeExpansion
	classdiagramDB.NodeGongNoteNodeExpansion_Data.Valid = true
}

// CopyBasicFieldsFromClassdiagramWOP
func (classdiagramDB *ClassdiagramDB) CopyBasicFieldsFromClassdiagramWOP(classdiagram *ClassdiagramWOP) {
	// insertion point for fields commit

	classdiagramDB.Name_Data.String = classdiagram.Name
	classdiagramDB.Name_Data.Valid = true

	classdiagramDB.Description_Data.String = classdiagram.Description
	classdiagramDB.Description_Data.Valid = true

	classdiagramDB.IsIncludedInStaticWebSite_Data.Bool = classdiagram.IsIncludedInStaticWebSite
	classdiagramDB.IsIncludedInStaticWebSite_Data.Valid = true

	classdiagramDB.IsInRenameMode_Data.Bool = classdiagram.IsInRenameMode
	classdiagramDB.IsInRenameMode_Data.Valid = true

	classdiagramDB.IsExpanded_Data.Bool = classdiagram.IsExpanded
	classdiagramDB.IsExpanded_Data.Valid = true

	classdiagramDB.NodeGongStructsIsExpanded_Data.Bool = classdiagram.NodeGongStructsIsExpanded
	classdiagramDB.NodeGongStructsIsExpanded_Data.Valid = true

	classdiagramDB.NodeGongStructNodeExpansion_Data.String = classdiagram.NodeGongStructNodeExpansion
	classdiagramDB.NodeGongStructNodeExpansion_Data.Valid = true

	classdiagramDB.NodeGongEnumsIsExpanded_Data.Bool = classdiagram.NodeGongEnumsIsExpanded
	classdiagramDB.NodeGongEnumsIsExpanded_Data.Valid = true

	classdiagramDB.NodeGongEnumNodeExpansion_Data.String = classdiagram.NodeGongEnumNodeExpansion
	classdiagramDB.NodeGongEnumNodeExpansion_Data.Valid = true

	classdiagramDB.NodeGongNotesIsExpanded_Data.Bool = classdiagram.NodeGongNotesIsExpanded
	classdiagramDB.NodeGongNotesIsExpanded_Data.Valid = true

	classdiagramDB.NodeGongNoteNodeExpansion_Data.String = classdiagram.NodeGongNoteNodeExpansion
	classdiagramDB.NodeGongNoteNodeExpansion_Data.Valid = true
}

// CopyBasicFieldsToClassdiagram
func (classdiagramDB *ClassdiagramDB) CopyBasicFieldsToClassdiagram(classdiagram *models.Classdiagram) {
	// insertion point for checkout of basic fields (back repo to stage)
	classdiagram.Name = classdiagramDB.Name_Data.String
	classdiagram.Description = classdiagramDB.Description_Data.String
	classdiagram.IsIncludedInStaticWebSite = classdiagramDB.IsIncludedInStaticWebSite_Data.Bool
	classdiagram.IsInRenameMode = classdiagramDB.IsInRenameMode_Data.Bool
	classdiagram.IsExpanded = classdiagramDB.IsExpanded_Data.Bool
	classdiagram.NodeGongStructsIsExpanded = classdiagramDB.NodeGongStructsIsExpanded_Data.Bool
	classdiagram.NodeGongStructNodeExpansion = classdiagramDB.NodeGongStructNodeExpansion_Data.String
	classdiagram.NodeGongEnumsIsExpanded = classdiagramDB.NodeGongEnumsIsExpanded_Data.Bool
	classdiagram.NodeGongEnumNodeExpansion = classdiagramDB.NodeGongEnumNodeExpansion_Data.String
	classdiagram.NodeGongNotesIsExpanded = classdiagramDB.NodeGongNotesIsExpanded_Data.Bool
	classdiagram.NodeGongNoteNodeExpansion = classdiagramDB.NodeGongNoteNodeExpansion_Data.String
}

// CopyBasicFieldsToClassdiagram_WOP
func (classdiagramDB *ClassdiagramDB) CopyBasicFieldsToClassdiagram_WOP(classdiagram *models.Classdiagram_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	classdiagram.Name = classdiagramDB.Name_Data.String
	classdiagram.Description = classdiagramDB.Description_Data.String
	classdiagram.IsIncludedInStaticWebSite = classdiagramDB.IsIncludedInStaticWebSite_Data.Bool
	classdiagram.IsInRenameMode = classdiagramDB.IsInRenameMode_Data.Bool
	classdiagram.IsExpanded = classdiagramDB.IsExpanded_Data.Bool
	classdiagram.NodeGongStructsIsExpanded = classdiagramDB.NodeGongStructsIsExpanded_Data.Bool
	classdiagram.NodeGongStructNodeExpansion = classdiagramDB.NodeGongStructNodeExpansion_Data.String
	classdiagram.NodeGongEnumsIsExpanded = classdiagramDB.NodeGongEnumsIsExpanded_Data.Bool
	classdiagram.NodeGongEnumNodeExpansion = classdiagramDB.NodeGongEnumNodeExpansion_Data.String
	classdiagram.NodeGongNotesIsExpanded = classdiagramDB.NodeGongNotesIsExpanded_Data.Bool
	classdiagram.NodeGongNoteNodeExpansion = classdiagramDB.NodeGongNoteNodeExpansion_Data.String
}

// CopyBasicFieldsToClassdiagramWOP
func (classdiagramDB *ClassdiagramDB) CopyBasicFieldsToClassdiagramWOP(classdiagram *ClassdiagramWOP) {
	classdiagram.ID = int(classdiagramDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	classdiagram.Name = classdiagramDB.Name_Data.String
	classdiagram.Description = classdiagramDB.Description_Data.String
	classdiagram.IsIncludedInStaticWebSite = classdiagramDB.IsIncludedInStaticWebSite_Data.Bool
	classdiagram.IsInRenameMode = classdiagramDB.IsInRenameMode_Data.Bool
	classdiagram.IsExpanded = classdiagramDB.IsExpanded_Data.Bool
	classdiagram.NodeGongStructsIsExpanded = classdiagramDB.NodeGongStructsIsExpanded_Data.Bool
	classdiagram.NodeGongStructNodeExpansion = classdiagramDB.NodeGongStructNodeExpansion_Data.String
	classdiagram.NodeGongEnumsIsExpanded = classdiagramDB.NodeGongEnumsIsExpanded_Data.Bool
	classdiagram.NodeGongEnumNodeExpansion = classdiagramDB.NodeGongEnumNodeExpansion_Data.String
	classdiagram.NodeGongNotesIsExpanded = classdiagramDB.NodeGongNotesIsExpanded_Data.Bool
	classdiagram.NodeGongNoteNodeExpansion = classdiagramDB.NodeGongNoteNodeExpansion_Data.String
}

// Backup generates a json file from a slice of all ClassdiagramDB instances in the backrepo
func (backRepoClassdiagram *BackRepoClassdiagramStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "ClassdiagramDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ClassdiagramDB, 0)
	for _, classdiagramDB := range backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramDB {
		forBackup = append(forBackup, classdiagramDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Classdiagram ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Classdiagram file", err.Error())
	}
}

// Backup generates a json file from a slice of all ClassdiagramDB instances in the backrepo
func (backRepoClassdiagram *BackRepoClassdiagramStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ClassdiagramDB, 0)
	for _, classdiagramDB := range backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramDB {
		forBackup = append(forBackup, classdiagramDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Classdiagram")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Classdiagram_Fields, -1)
	for _, classdiagramDB := range forBackup {

		var classdiagramWOP ClassdiagramWOP
		classdiagramDB.CopyBasicFieldsToClassdiagramWOP(&classdiagramWOP)

		row := sh.AddRow()
		row.WriteStruct(&classdiagramWOP, -1)
	}
}

// RestoreXL from the "Classdiagram" sheet all ClassdiagramDB instances
func (backRepoClassdiagram *BackRepoClassdiagramStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoClassdiagramid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Classdiagram"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoClassdiagram.rowVisitorClassdiagram)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoClassdiagram *BackRepoClassdiagramStruct) rowVisitorClassdiagram(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var classdiagramWOP ClassdiagramWOP
		row.ReadStruct(&classdiagramWOP)

		// add the unmarshalled struct to the stage
		classdiagramDB := new(ClassdiagramDB)
		classdiagramDB.CopyBasicFieldsFromClassdiagramWOP(&classdiagramWOP)

		classdiagramDB_ID_atBackupTime := classdiagramDB.ID
		classdiagramDB.ID = 0
		_, err := backRepoClassdiagram.db.Create(classdiagramDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramDB[classdiagramDB.ID] = classdiagramDB
		BackRepoClassdiagramid_atBckpTime_newID[classdiagramDB_ID_atBackupTime] = classdiagramDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "ClassdiagramDB.json" in dirPath that stores an array
// of ClassdiagramDB and stores it in the database
// the map BackRepoClassdiagramid_atBckpTime_newID is updated accordingly
func (backRepoClassdiagram *BackRepoClassdiagramStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoClassdiagramid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "ClassdiagramDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Classdiagram file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*ClassdiagramDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_ClassdiagramDBID_ClassdiagramDB
	for _, classdiagramDB := range forRestore {

		classdiagramDB_ID_atBackupTime := classdiagramDB.ID
		classdiagramDB.ID = 0
		_, err := backRepoClassdiagram.db.Create(classdiagramDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramDB[classdiagramDB.ID] = classdiagramDB
		BackRepoClassdiagramid_atBckpTime_newID[classdiagramDB_ID_atBackupTime] = classdiagramDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Classdiagram file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Classdiagram>id_atBckpTime_newID
// to compute new index
func (backRepoClassdiagram *BackRepoClassdiagramStruct) RestorePhaseTwo() {

	for _, classdiagramDB := range backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramDB {

		// next line of code is to avert unused variable compilation error
		_ = classdiagramDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoClassdiagram.db.Model(classdiagramDB)
		_, err := db.Updates(*classdiagramDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoClassdiagram.ResetReversePointers commits all staged instances of Classdiagram to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoClassdiagram *BackRepoClassdiagramStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, classdiagram := range backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr {
		backRepoClassdiagram.ResetReversePointersInstance(backRepo, idx, classdiagram)
	}

	return
}

func (backRepoClassdiagram *BackRepoClassdiagramStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, classdiagram *models.Classdiagram) (Error error) {

	// fetch matching classdiagramDB
	if classdiagramDB, ok := backRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramDB[idx]; ok {
		_ = classdiagramDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoClassdiagramid_atBckpTime_newID map[uint]uint
