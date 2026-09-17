// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/sim/go/models"
	"github.com/fullstack-lang/gong/lib/sim/go/orm"
)

// declaration in order to justify use of the models import
var __UpdateState__dummysDeclaration__ models.UpdateState
var _ = __UpdateState__dummysDeclaration__
var __UpdateState_time__dummyDeclaration time.Duration
var _ = __UpdateState_time__dummyDeclaration

var mutexUpdateState sync.Mutex

// An UpdateStateID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateUpdateState
type UpdateStateID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// UpdateStateInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateUpdateState
type UpdateStateInput struct {
	// The UpdateState to submit or modify
	// in: body
	UpdateState *orm.UpdateStateAPI
}

// UpdateUpdateState
//
// swagger:route PATCH /updatestates/{ID} updatestates updateUpdateState
//
// # Update a updatestate
//
// Responses:
// default: genericError
//
//	200: updatestateDBResponse
func (controller *Controller) UpdateUpdateState(w http.ResponseWriter, r *http.Request) {

	mutexUpdateState.Lock()
	defer mutexUpdateState.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoUpdateState.GetDB()

	// Validate input
	var input orm.UpdateStateAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var updatestateDB orm.UpdateStateDB

	// fetch the updatestate
	_, err := db.First(&updatestateDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	updatestateDB.CopyBasicFieldsFromUpdateState_WOP(&input.UpdateState_WOP)
	updatestateDB.UpdateStatePointersEncoding = input.UpdateStatePointersEncoding

	db, _ = db.Model(&updatestateDB)
	_, err = db.Updates(&updatestateDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	updatestateNew := new(models.UpdateState)
	updatestateDB.CopyBasicFieldsToUpdateState(updatestateNew)

	// redeem pointers
	updatestateDB.DecodePointers(backRepo, updatestateNew)

	// get stage instance from DB instance, and call callback function
	updatestateOld := backRepo.BackRepoUpdateState.Map_UpdateStateDBID_UpdateStatePtr[updatestateDB.ID]
	if updatestateOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(updatestateOld, updatestateNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the updatestateDB
	writeJSON(w, http.StatusOK, updatestateDB)
}
