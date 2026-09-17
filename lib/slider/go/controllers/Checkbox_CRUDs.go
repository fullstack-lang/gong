// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/slider/go/models"
	"github.com/fullstack-lang/gong/lib/slider/go/orm"
)

// declaration in order to justify use of the models import
var __Checkbox__dummysDeclaration__ models.Checkbox
var _ = __Checkbox__dummysDeclaration__
var __Checkbox_time__dummyDeclaration time.Duration
var _ = __Checkbox_time__dummyDeclaration

var mutexCheckbox sync.Mutex

// An CheckboxID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateCheckbox
type CheckboxID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CheckboxInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateCheckbox
type CheckboxInput struct {
	// The Checkbox to submit or modify
	// in: body
	Checkbox *orm.CheckboxAPI
}

// UpdateCheckbox
//
// swagger:route PATCH /checkboxs/{ID} checkboxs updateCheckbox
//
// # Update a checkbox
//
// Responses:
// default: genericError
//
//	200: checkboxDBResponse
func (controller *Controller) UpdateCheckbox(w http.ResponseWriter, r *http.Request) {

	mutexCheckbox.Lock()
	defer mutexCheckbox.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/slider/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoCheckbox.GetDB()

	// Validate input
	var input orm.CheckboxAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var checkboxDB orm.CheckboxDB

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
	checkboxDB.CopyBasicFieldsFromCheckbox_WOP(&input.Checkbox_WOP)
	checkboxDB.CheckboxPointersEncoding = input.CheckboxPointersEncoding

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
	checkboxNew := new(models.Checkbox)
	checkboxDB.CopyBasicFieldsToCheckbox(checkboxNew)

	// redeem pointers
	checkboxDB.DecodePointers(backRepo, checkboxNew)

	// get stage instance from DB instance, and call callback function
	checkboxOld := backRepo.BackRepoCheckbox.Map_CheckboxDBID_CheckboxPtr[checkboxDB.ID]
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
