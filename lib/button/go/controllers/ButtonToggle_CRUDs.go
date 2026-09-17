// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/button/go/models"
	"github.com/fullstack-lang/gong/lib/button/go/orm"
)

// declaration in order to justify use of the models import
var __ButtonToggle__dummysDeclaration__ models.ButtonToggle
var _ = __ButtonToggle__dummysDeclaration__
var __ButtonToggle_time__dummyDeclaration time.Duration
var _ = __ButtonToggle_time__dummyDeclaration

var mutexButtonToggle sync.Mutex

// An ButtonToggleID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateButtonToggle
type ButtonToggleID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ButtonToggleInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateButtonToggle
type ButtonToggleInput struct {
	// The ButtonToggle to submit or modify
	// in: body
	ButtonToggle *orm.ButtonToggleAPI
}

// UpdateButtonToggle
//
// swagger:route PATCH /buttontoggles/{ID} buttontoggles updateButtonToggle
//
// # Update a buttontoggle
//
// Responses:
// default: genericError
//
//	200: buttontoggleDBResponse
func (controller *Controller) UpdateButtonToggle(w http.ResponseWriter, r *http.Request) {

	mutexButtonToggle.Lock()
	defer mutexButtonToggle.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoButtonToggle.GetDB()

	// Validate input
	var input orm.ButtonToggleAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var buttontoggleDB orm.ButtonToggleDB

	// fetch the buttontoggle
	_, err := db.First(&buttontoggleDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	buttontoggleDB.CopyBasicFieldsFromButtonToggle_WOP(&input.ButtonToggle_WOP)
	buttontoggleDB.ButtonTogglePointersEncoding = input.ButtonTogglePointersEncoding

	db, _ = db.Model(&buttontoggleDB)
	_, err = db.Updates(&buttontoggleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	buttontoggleNew := new(models.ButtonToggle)
	buttontoggleDB.CopyBasicFieldsToButtonToggle(buttontoggleNew)

	// redeem pointers
	buttontoggleDB.DecodePointers(backRepo, buttontoggleNew)

	// get stage instance from DB instance, and call callback function
	buttontoggleOld := backRepo.BackRepoButtonToggle.Map_ButtonToggleDBID_ButtonTogglePtr[buttontoggleDB.ID]
	if buttontoggleOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(buttontoggleOld, buttontoggleNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the buttontoggleDB
	writeJSON(w, http.StatusOK, buttontoggleDB)
}
