// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/threejs/go/models"
	"github.com/fullstack-lang/gong/lib/threejs/go/orm"
)

// declaration in order to justify use of the models import
var __DirectionalLight__dummysDeclaration__ models.DirectionalLight
var _ = __DirectionalLight__dummysDeclaration__
var __DirectionalLight_time__dummyDeclaration time.Duration
var _ = __DirectionalLight_time__dummyDeclaration

var mutexDirectionalLight sync.Mutex

// An DirectionalLightID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateDirectionalLight
type DirectionalLightID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DirectionalLightInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateDirectionalLight
type DirectionalLightInput struct {
	// The DirectionalLight to submit or modify
	// in: body
	DirectionalLight *orm.DirectionalLightAPI
}

// UpdateDirectionalLight
//
// swagger:route PATCH /directionallights/{ID} directionallights updateDirectionalLight
//
// # Update a directionallight
//
// Responses:
// default: genericError
//
//	200: directionallightDBResponse
func (controller *Controller) UpdateDirectionalLight(w http.ResponseWriter, r *http.Request) {

	mutexDirectionalLight.Lock()
	defer mutexDirectionalLight.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoDirectionalLight.GetDB()

	// Validate input
	var input orm.DirectionalLightAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var directionallightDB orm.DirectionalLightDB

	// fetch the directionallight
	_, err := db.First(&directionallightDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	directionallightDB.CopyBasicFieldsFromDirectionalLight_WOP(&input.DirectionalLight_WOP)
	directionallightDB.DirectionalLightPointersEncoding = input.DirectionalLightPointersEncoding

	db, _ = db.Model(&directionallightDB)
	_, err = db.Updates(&directionallightDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	directionallightNew := new(models.DirectionalLight)
	directionallightDB.CopyBasicFieldsToDirectionalLight(directionallightNew)

	// redeem pointers
	directionallightDB.DecodePointers(backRepo, directionallightNew)

	// get stage instance from DB instance, and call callback function
	directionallightOld := backRepo.BackRepoDirectionalLight.Map_DirectionalLightDBID_DirectionalLightPtr[directionallightDB.ID]
	if directionallightOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(directionallightOld, directionallightNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the directionallightDB
	writeJSON(w, http.StatusOK, directionallightDB)
}
