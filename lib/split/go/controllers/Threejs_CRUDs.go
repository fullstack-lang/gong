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
var __Threejs__dummysDeclaration__ models.Threejs
var _ = __Threejs__dummysDeclaration__
var __Threejs_time__dummyDeclaration time.Duration
var _ = __Threejs_time__dummyDeclaration

var mutexThreejs sync.Mutex

// An ThreejsID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateThreejs
type ThreejsID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ThreejsInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateThreejs
type ThreejsInput struct {
	// The Threejs to submit or modify
	// in: body
	Threejs *orm.ThreejsAPI
}

// UpdateThreejs
//
// swagger:route PATCH /threejss/{ID} threejss updateThreejs
//
// # Update a threejs
//
// Responses:
// default: genericError
//
//	200: threejsDBResponse
func (controller *Controller) UpdateThreejs(w http.ResponseWriter, r *http.Request) {

	mutexThreejs.Lock()
	defer mutexThreejs.Unlock()

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
	db := backRepo.BackRepoThreejs.GetDB()

	// Validate input
	var input orm.ThreejsAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var threejsDB orm.ThreejsDB

	// fetch the threejs
	_, err := db.First(&threejsDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	threejsDB.CopyBasicFieldsFromThreejs_WOP(&input.Threejs_WOP)
	threejsDB.ThreejsPointersEncoding = input.ThreejsPointersEncoding

	db, _ = db.Model(&threejsDB)
	_, err = db.Updates(&threejsDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	threejsNew := new(models.Threejs)
	threejsDB.CopyBasicFieldsToThreejs(threejsNew)

	// redeem pointers
	threejsDB.DecodePointers(backRepo, threejsNew)

	// get stage instance from DB instance, and call callback function
	threejsOld := backRepo.BackRepoThreejs.Map_ThreejsDBID_ThreejsPtr[threejsDB.ID]
	if threejsOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(threejsOld, threejsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the threejsDB
	writeJSON(w, http.StatusOK, threejsDB)
}
