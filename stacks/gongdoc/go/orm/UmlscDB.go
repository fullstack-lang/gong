// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/fullstack-lang/gong/stacks/gongdoc/go/models"
	"github.com/jinzhu/gorm"
)

// dummy variable to have the import database/sql wihthout compile failure id no sql is used
var dummy_Umlsc sql.NullBool

// UmlscAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model umlscAPI
type UmlscAPI struct {
	models.Umlsc

	// insertion for fields declaration
	// Declation for basic field umlscDB.Name {{BasicKind}} (to be completed)
	Name_Data sql.NullString

	// Declation for basic field umlscDB.Activestate {{BasicKind}} (to be completed)
	Activestate_Data sql.NullString

	// Implementation of a reverse ID for field Pkgelt{}.Umlscs []*Umlsc
	Pkgelt_UmlscsDBID sql.NullInt64

	// end of insertion
}

// UmlscDB describes a umlsc in the database
//
// It incorporates all fields : from the model, from the generated field for the API and the GORM ID
//
// swagger:model umlscDB
type UmlscDB struct {
	gorm.Model

	UmlscAPI
}

// UmlscDBs arrays umlscDBs
// swagger:response umlscDBsResponse
type UmlscDBs []UmlscDB

// UmlscDBResponse provides response
// swagger:response umlscDBResponse
type UmlscDBResponse struct {
	UmlscDB
}

type BackRepoUmlscStruct struct {
	// stores UmlscDB according to their gorm ID
	Map_UmlscDBID_UmlscDB *map[uint]*UmlscDB

	// stores UmlscDB ID according to Umlsc address
	Map_UmlscPtr_UmlscDBID *map[*models.Umlsc]uint

	// stores Umlsc according to their gorm ID
	Map_UmlscDBID_UmlscPtr *map[uint]*models.Umlsc

	db *gorm.DB
}

// BackRepoUmlsc.Init set up the BackRepo of the Umlsc
func (backRepoUmlsc *BackRepoUmlscStruct) Init(db *gorm.DB) (Error error) {

	if backRepoUmlsc.Map_UmlscDBID_UmlscPtr != nil {
		err := errors.New("In Init, backRepoUmlsc.Map_UmlscDBID_UmlscPtr should be nil")
		return err
	}

	if backRepoUmlsc.Map_UmlscDBID_UmlscDB != nil {
		err := errors.New("In Init, backRepoUmlsc.Map_UmlscDBID_UmlscDB should be nil")
		return err
	}

	if backRepoUmlsc.Map_UmlscPtr_UmlscDBID != nil {
		err := errors.New("In Init, backRepoUmlsc.Map_UmlscPtr_UmlscDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.Umlsc, 0)
	backRepoUmlsc.Map_UmlscDBID_UmlscPtr = &tmp

	tmpDB := make(map[uint]*UmlscDB, 0)
	backRepoUmlsc.Map_UmlscDBID_UmlscDB = &tmpDB

	tmpID := make(map[*models.Umlsc]uint, 0)
	backRepoUmlsc.Map_UmlscPtr_UmlscDBID = &tmpID

	backRepoUmlsc.db = db
	return
}

// BackRepoUmlsc.CommitPhaseOne commits all staged instances of Umlsc to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoUmlsc *BackRepoUmlscStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for umlsc := range stage.Umlscs {
		backRepoUmlsc.CommitPhaseOneInstance(umlsc)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, umlsc := range *backRepoUmlsc.Map_UmlscDBID_UmlscPtr {
		if _, ok := stage.Umlscs[umlsc]; !ok {
			backRepoUmlsc.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoUmlsc.CommitDeleteInstance commits deletion of Umlsc to the BackRepo
func (backRepoUmlsc *BackRepoUmlscStruct) CommitDeleteInstance(id uint) (Error error) {

	umlsc := (*backRepoUmlsc.Map_UmlscDBID_UmlscPtr)[id]

	// umlsc is not staged anymore, remove umlscDB
	umlscDB := (*backRepoUmlsc.Map_UmlscDBID_UmlscDB)[id]
	query := backRepoUmlsc.db.Unscoped().Delete(&umlscDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoUmlsc.Map_UmlscPtr_UmlscDBID), umlsc)
	delete((*backRepoUmlsc.Map_UmlscDBID_UmlscPtr), id)
	delete((*backRepoUmlsc.Map_UmlscDBID_UmlscDB), id)

	return
}

// BackRepoUmlsc.CommitPhaseOneInstance commits umlsc staged instances of Umlsc to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoUmlsc *BackRepoUmlscStruct) CommitPhaseOneInstance(umlsc *models.Umlsc) (Error error) {

	// check if the umlsc is not commited yet
	if _, ok := (*backRepoUmlsc.Map_UmlscPtr_UmlscDBID)[umlsc]; ok {
		return
	}

	// initiate umlsc
	var umlscDB UmlscDB
	umlscDB.Umlsc = *umlsc

	query := backRepoUmlsc.db.Create(&umlscDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoUmlsc.Map_UmlscPtr_UmlscDBID)[umlsc] = umlscDB.ID
	(*backRepoUmlsc.Map_UmlscDBID_UmlscPtr)[umlscDB.ID] = umlsc
	(*backRepoUmlsc.Map_UmlscDBID_UmlscDB)[umlscDB.ID] = &umlscDB

	return
}

// BackRepoUmlsc.CommitPhaseTwo commits all staged instances of Umlsc to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUmlsc *BackRepoUmlscStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, umlsc := range *backRepoUmlsc.Map_UmlscDBID_UmlscPtr {
		backRepoUmlsc.CommitPhaseTwoInstance(backRepo, idx, umlsc)
	}

	return
}

// BackRepoUmlsc.CommitPhaseTwoInstance commits {{structname }} of models.Umlsc to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUmlsc *BackRepoUmlscStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, umlsc *models.Umlsc) (Error error) {

	// fetch matching umlscDB
	if umlscDB, ok := (*backRepoUmlsc.Map_UmlscDBID_UmlscDB)[idx]; ok {

		{
			{
				// insertion point for fields commit
				umlscDB.Name_Data.String = umlsc.Name
				umlscDB.Name_Data.Valid = true

				// commit a slice of pointer translates to update reverse pointer to State, i.e.
				for _, state := range umlsc.States {
					if stateDBID, ok := (*backRepo.BackRepoState.Map_StatePtr_StateDBID)[state]; ok {
						if stateDB, ok := (*backRepo.BackRepoState.Map_StateDBID_StateDB)[stateDBID]; ok {
							stateDB.Umlsc_StatesDBID.Int64 = int64(umlscDB.ID)
							stateDB.Umlsc_StatesDBID.Valid = true
							if q := backRepoUmlsc.db.Save(&stateDB); q.Error != nil {
								return q.Error
							}
						}
					}
				}

				umlscDB.Activestate_Data.String = umlsc.Activestate
				umlscDB.Activestate_Data.Valid = true

			}
		}
		query := backRepoUmlsc.db.Save(&umlscDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Umlsc intance %s", umlsc.Name))
		return err
	}

	return
}

// BackRepoUmlsc.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One is the creation of instance in the stage
//
// NOTE: the is supposed to have been reset before
//
func (backRepoUmlsc *BackRepoUmlscStruct) CheckoutPhaseOne() (Error error) {

	umlscDBArray := make([]UmlscDB, 0)
	query := backRepoUmlsc.db.Find(&umlscDBArray)
	if query.Error != nil {
		return query.Error
	}

	// copy orm objects to the the map
	for _, umlscDB := range umlscDBArray {
		backRepoUmlsc.CheckoutPhaseOneInstance(&umlscDB)
	}

	return
}

// CheckoutPhaseOneInstance takes a umlscDB that has been found in the DB, updates the backRepo and stages the
// models version of the umlscDB
func (backRepoUmlsc *BackRepoUmlscStruct) CheckoutPhaseOneInstance(umlscDB *UmlscDB) (Error error) {

	// if absent, create entries in the backRepoUmlsc maps.
	umlscWithNewFieldValues := umlscDB.Umlsc
	if _, ok := (*backRepoUmlsc.Map_UmlscDBID_UmlscPtr)[umlscDB.ID]; !ok {

		(*backRepoUmlsc.Map_UmlscDBID_UmlscPtr)[umlscDB.ID] = &umlscWithNewFieldValues
		(*backRepoUmlsc.Map_UmlscPtr_UmlscDBID)[&umlscWithNewFieldValues] = umlscDB.ID

		// append model store with the new element
		umlscWithNewFieldValues.Stage()
	}
	umlscDBWithNewFieldValues := *umlscDB
	(*backRepoUmlsc.Map_UmlscDBID_UmlscDB)[umlscDB.ID] = &umlscDBWithNewFieldValues

	return
}

// BackRepoUmlsc.CheckoutPhaseTwo Checkouts all staged instances of Umlsc to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUmlsc *BackRepoUmlscStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, umlscDB := range *backRepoUmlsc.Map_UmlscDBID_UmlscDB {
		backRepoUmlsc.CheckoutPhaseTwoInstance(backRepo, umlscDB)
	}
	return
}

// BackRepoUmlsc.CheckoutPhaseTwoInstance Checkouts staged instances of Umlsc to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUmlsc *BackRepoUmlscStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, umlscDB *UmlscDB) (Error error) {

	umlsc := (*backRepoUmlsc.Map_UmlscDBID_UmlscPtr)[umlscDB.ID]
	_ = umlsc // sometimes, there is no code generated. This lines voids the "unused variable" compilation error
	{
		{
			// insertion point for checkout, i.e. update of fields of stage instance from fields of back repo instances
			//
			umlsc.Name = umlscDB.Name_Data.String

			// parse all StateDB and redeem the array of poiners to Umlsc
			// first reset the slice
			umlsc.States = umlsc.States[:0]
			for _, StateDB := range *backRepo.BackRepoState.Map_StateDBID_StateDB {
				if StateDB.Umlsc_StatesDBID.Int64 == int64(umlscDB.ID) {
					State := (*backRepo.BackRepoState.Map_StateDBID_StatePtr)[StateDB.ID]
					umlsc.States = append(umlsc.States, State)
				}
			}

			umlsc.Activestate = umlscDB.Activestate_Data.String

		}
	}
	return
}

// CommitUmlsc allows commit of a single umlsc (if already staged)
func (backRepo *BackRepoStruct) CommitUmlsc(umlsc *models.Umlsc) {
	backRepo.BackRepoUmlsc.CommitPhaseOneInstance(umlsc)
	if id, ok := (*backRepo.BackRepoUmlsc.Map_UmlscPtr_UmlscDBID)[umlsc]; ok {
		backRepo.BackRepoUmlsc.CommitPhaseTwoInstance(backRepo, id, umlsc)
	}
}

// CommitUmlsc allows checkout of a single umlsc (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutUmlsc(umlsc *models.Umlsc) {
	// check if the umlsc is staged
	if _, ok := (*backRepo.BackRepoUmlsc.Map_UmlscPtr_UmlscDBID)[umlsc]; ok {

		if id, ok := (*backRepo.BackRepoUmlsc.Map_UmlscPtr_UmlscDBID)[umlsc]; ok {
			var umlscDB UmlscDB
			umlscDB.ID = id

			if err := backRepo.BackRepoUmlsc.db.First(&umlscDB, id).Error; err != nil {
				log.Panicln("CheckoutUmlsc : Problem with getting object with id:", id)
			}
			backRepo.BackRepoUmlsc.CheckoutPhaseOneInstance(&umlscDB)
			backRepo.BackRepoUmlsc.CheckoutPhaseTwoInstance(backRepo, &umlscDB)
		}
	}
}