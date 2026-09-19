// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/splitlite/go/models"
	"github.com/fullstack-lang/gong/lib/splitlite/go/orm"
)

// declaration in order to justify use of the models import
var __AsSplitArea__dummysDeclaration__ models.AsSplitArea
var _ = __AsSplitArea__dummysDeclaration__
var __AsSplitArea_time__dummyDeclaration time.Duration
var _ = __AsSplitArea_time__dummyDeclaration

var mutexAsSplitArea sync.Mutex

// An AsSplitAreaID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateAsSplitArea
type AsSplitAreaID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AsSplitAreaInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateAsSplitArea
type AsSplitAreaInput struct {
	// The AsSplitArea to submit or modify
	// in: body
	AsSplitArea *orm.AsSplitAreaAPI
}

// UpdateAsSplitArea
//
// swagger:route PATCH /assplitareas/{ID} assplitareas updateAsSplitArea
//
// # Update a assplitarea
//
// Responses:
// default: genericError
//
//	200: assplitareaDBResponse
func (controller *Controller) UpdateAsSplitArea(w http.ResponseWriter, r *http.Request) {

	mutexAsSplitArea.Lock()
	defer mutexAsSplitArea.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/splitlite/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoAsSplitArea.GetDB()

	// Validate input
	var input orm.AsSplitAreaAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var assplitareaDB orm.AsSplitAreaDB

	// fetch the assplitarea
	_, err := db.First(&assplitareaDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	assplitareaDB.CopyBasicFieldsFromAsSplitArea_WOP(&input.AsSplitArea_WOP)
	assplitareaDB.AsSplitAreaPointersEncoding = input.AsSplitAreaPointersEncoding

	db, _ = db.Model(&assplitareaDB)
	_, err = db.Updates(&assplitareaDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	assplitareaNew := new(models.AsSplitArea)
	assplitareaDB.CopyBasicFieldsToAsSplitArea(assplitareaNew)

	// redeem pointers
	assplitareaDB.DecodePointers(backRepo, assplitareaNew)

	// get stage instance from DB instance, and call callback function
	assplitareaOld := backRepo.BackRepoAsSplitArea.Map_AsSplitAreaDBID_AsSplitAreaPtr[assplitareaDB.ID]
	if assplitareaOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(assplitareaOld, assplitareaNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the assplitareaDB
	writeJSON(w, http.StatusOK, assplitareaDB)
}
