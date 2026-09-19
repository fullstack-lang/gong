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
var __Load__dummysDeclaration__ models.Load
var _ = __Load__dummysDeclaration__
var __Load_time__dummyDeclaration time.Duration
var _ = __Load_time__dummyDeclaration

var mutexLoad sync.Mutex

// An LoadID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateLoad
type LoadID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LoadInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateLoad
type LoadInput struct {
	// The Load to submit or modify
	// in: body
	Load *orm.LoadAPI
}

// UpdateLoad
//
// swagger:route PATCH /loads/{ID} loads updateLoad
//
// # Update a load
//
// Responses:
// default: genericError
//
//	200: loadDBResponse
func (controller *Controller) UpdateLoad(w http.ResponseWriter, r *http.Request) {

	mutexLoad.Lock()
	defer mutexLoad.Unlock()

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
	db := backRepo.BackRepoLoad.GetDB()

	// Validate input
	var input orm.LoadAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var loadDB orm.LoadDB

	// fetch the load
	_, err := db.First(&loadDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	loadDB.CopyBasicFieldsFromLoad_WOP(&input.Load_WOP)
	loadDB.LoadPointersEncoding = input.LoadPointersEncoding

	db, _ = db.Model(&loadDB)
	_, err = db.Updates(&loadDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	loadNew := new(models.Load)
	loadDB.CopyBasicFieldsToLoad(loadNew)

	// redeem pointers
	loadDB.DecodePointers(backRepo, loadNew)

	// get stage instance from DB instance, and call callback function
	loadOld := backRepo.BackRepoLoad.Map_LoadDBID_LoadPtr[loadDB.ID]
	if loadOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(loadOld, loadNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the loadDB
	writeJSON(w, http.StatusOK, loadDB)
}
