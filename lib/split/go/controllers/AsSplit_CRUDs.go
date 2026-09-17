// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/split/go/models"
	"github.com/fullstack-lang/gong/lib/split/go/orm"
)

// declaration in order to justify use of the models import
var __AsSplit__dummysDeclaration__ models.AsSplit
var _ = __AsSplit__dummysDeclaration__
var __AsSplit_time__dummyDeclaration time.Duration
var _ = __AsSplit_time__dummyDeclaration

var mutexAsSplit sync.Mutex

// An AsSplitID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateAsSplit
type AsSplitID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AsSplitInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateAsSplit
type AsSplitInput struct {
	// The AsSplit to submit or modify
	// in: body
	AsSplit *orm.AsSplitAPI
}

// UpdateAsSplit
//
// swagger:route PATCH /assplits/{ID} assplits updateAsSplit
//
// # Update a assplit
//
// Responses:
// default: genericError
//
//	200: assplitDBResponse
func (controller *Controller) UpdateAsSplit(w http.ResponseWriter, r *http.Request) {

	mutexAsSplit.Lock()
	defer mutexAsSplit.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoAsSplit.GetDB()

	// Validate input
	var input orm.AsSplitAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var assplitDB orm.AsSplitDB

	// fetch the assplit
	_, err := db.First(&assplitDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	assplitDB.CopyBasicFieldsFromAsSplit_WOP(&input.AsSplit_WOP)
	assplitDB.AsSplitPointersEncoding = input.AsSplitPointersEncoding

	db, _ = db.Model(&assplitDB)
	_, err = db.Updates(&assplitDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	assplitNew := new(models.AsSplit)
	assplitDB.CopyBasicFieldsToAsSplit(assplitNew)

	// redeem pointers
	assplitDB.DecodePointers(backRepo, assplitNew)

	// get stage instance from DB instance, and call callback function
	assplitOld := backRepo.BackRepoAsSplit.Map_AsSplitDBID_AsSplitPtr[assplitDB.ID]
	if assplitOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(assplitOld, assplitNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the assplitDB
	writeJSON(w, http.StatusOK, assplitDB)
}
