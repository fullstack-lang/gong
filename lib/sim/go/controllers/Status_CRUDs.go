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
var __Status__dummysDeclaration__ models.Status
var _ = __Status__dummysDeclaration__
var __Status_time__dummyDeclaration time.Duration
var _ = __Status_time__dummyDeclaration

var mutexStatus sync.Mutex

// An StatusID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateStatus
type StatusID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// StatusInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateStatus
type StatusInput struct {
	// The Status to submit or modify
	// in: body
	Status *orm.StatusAPI
}

// UpdateStatus
//
// swagger:route PATCH /statuss/{ID} statuss updateStatus
//
// # Update a status
//
// Responses:
// default: genericError
//
//	200: statusDBResponse
func (controller *Controller) UpdateStatus(w http.ResponseWriter, r *http.Request) {

	mutexStatus.Lock()
	defer mutexStatus.Unlock()

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
	db := backRepo.BackRepoStatus.GetDB()

	// Validate input
	var input orm.StatusAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var statusDB orm.StatusDB

	// fetch the status
	_, err := db.First(&statusDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	statusDB.CopyBasicFieldsFromStatus_WOP(&input.Status_WOP)
	statusDB.StatusPointersEncoding = input.StatusPointersEncoding

	db, _ = db.Model(&statusDB)
	_, err = db.Updates(&statusDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	statusNew := new(models.Status)
	statusDB.CopyBasicFieldsToStatus(statusNew)

	// redeem pointers
	statusDB.DecodePointers(backRepo, statusNew)

	// get stage instance from DB instance, and call callback function
	statusOld := backRepo.BackRepoStatus.Map_StatusDBID_StatusPtr[statusDB.ID]
	if statusOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(statusOld, statusNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the statusDB
	writeJSON(w, http.StatusOK, statusDB)
}
