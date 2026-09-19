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
var __Split__dummysDeclaration__ models.Split
var _ = __Split__dummysDeclaration__
var __Split_time__dummyDeclaration time.Duration
var _ = __Split_time__dummyDeclaration

var mutexSplit sync.Mutex

// An SplitID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateSplit
type SplitID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SplitInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateSplit
type SplitInput struct {
	// The Split to submit or modify
	// in: body
	Split *orm.SplitAPI
}

// UpdateSplit
//
// swagger:route PATCH /splits/{ID} splits updateSplit
//
// # Update a split
//
// Responses:
// default: genericError
//
//	200: splitDBResponse
func (controller *Controller) UpdateSplit(w http.ResponseWriter, r *http.Request) {

	mutexSplit.Lock()
	defer mutexSplit.Unlock()

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
	db := backRepo.BackRepoSplit.GetDB()

	// Validate input
	var input orm.SplitAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var splitDB orm.SplitDB

	// fetch the split
	_, err := db.First(&splitDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	splitDB.CopyBasicFieldsFromSplit_WOP(&input.Split_WOP)
	splitDB.SplitPointersEncoding = input.SplitPointersEncoding

	db, _ = db.Model(&splitDB)
	_, err = db.Updates(&splitDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	splitNew := new(models.Split)
	splitDB.CopyBasicFieldsToSplit(splitNew)

	// redeem pointers
	splitDB.DecodePointers(backRepo, splitNew)

	// get stage instance from DB instance, and call callback function
	splitOld := backRepo.BackRepoSplit.Map_SplitDBID_SplitPtr[splitDB.ID]
	if splitOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(splitOld, splitNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the splitDB
	writeJSON(w, http.StatusOK, splitDB)
}
