// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/form/go/models"
	"github.com/fullstack-lang/gong/lib/form/go/orm"
)

// declaration in order to justify use of the models import
var __CheckBox__dummysDeclaration__ models.CheckBox
var _ = __CheckBox__dummysDeclaration__
var __CheckBox_time__dummyDeclaration time.Duration
var _ = __CheckBox_time__dummyDeclaration

var mutexCheckBox sync.Mutex

// An CheckBoxID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateCheckBox
type CheckBoxID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CheckBoxInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateCheckBox
type CheckBoxInput struct {
	// The CheckBox to submit or modify
	// in: body
	CheckBox *orm.CheckBoxAPI
}

// UpdateCheckBox
//
// swagger:route PATCH /checkboxs/{ID} checkboxs updateCheckBox
//
// # Update a checkbox
//
// Responses:
// default: genericError
//
//	200: checkboxDBResponse
func (controller *Controller) UpdateCheckBox(w http.ResponseWriter, r *http.Request) {

	mutexCheckBox.Lock()
	defer mutexCheckBox.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) >= 1 {
		_nameValues := _values["Name"]
		if len(_nameValues) == 1 {
			stackPath = _nameValues[0]
		}
	}

	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/form/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoCheckBox.GetDB()

	// Validate input
	var input orm.CheckBoxAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var checkboxDB orm.CheckBoxDB

	// fetch the checkbox
	_, err := db.First(&checkboxDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	checkboxDB.CopyBasicFieldsFromCheckBox_WOP(&input.CheckBox_WOP)
	checkboxDB.CheckBoxPointersEncoding = input.CheckBoxPointersEncoding

	db, _ = db.Model(&checkboxDB)
	_, err = db.Updates(&checkboxDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	checkboxNew := new(models.CheckBox)
	checkboxDB.CopyBasicFieldsToCheckBox(checkboxNew)

	// redeem pointers
	checkboxDB.DecodePointers(backRepo, checkboxNew)

	// get stage instance from DB instance, and call callback function
	checkboxOld := backRepo.BackRepoCheckBox.Map_CheckBoxDBID_CheckBoxPtr[checkboxDB.ID]
	if checkboxOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(checkboxOld, checkboxNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the checkboxDB
	writeJSON(w, http.StatusOK, checkboxDB)
}
