// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/test/test1/go/models"
	"github.com/fullstack-lang/gong/test/test1/go/orm"
)

// declaration in order to justify use of the models import
var __Gstruct__dummysDeclaration__ models.Gstruct
var _ = __Gstruct__dummysDeclaration__
var __Gstruct_time__dummyDeclaration time.Duration
var _ = __Gstruct_time__dummyDeclaration

var mutexGstruct sync.Mutex

// An GstructID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateGstruct
type GstructID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GstructInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateGstruct
type GstructInput struct {
	// The Gstruct to submit or modify
	// in: body
	Gstruct *orm.GstructAPI
}

// UpdateGstruct
//
// swagger:route PATCH /gstructs/{ID} gstructs updateGstruct
//
// # Update a gstruct
//
// Responses:
// default: genericError
//
//	200: gstructDBResponse
func (controller *Controller) UpdateGstruct(w http.ResponseWriter, r *http.Request) {

	mutexGstruct.Lock()
	defer mutexGstruct.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/test/test1/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoGstruct.GetDB()

	// Validate input
	var input orm.GstructAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var gstructDB orm.GstructDB

	// fetch the gstruct
	_, err := db.First(&gstructDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	gstructDB.CopyBasicFieldsFromGstruct_WOP(&input.Gstruct_WOP)
	gstructDB.GstructPointersEncoding = input.GstructPointersEncoding

	db, _ = db.Model(&gstructDB)
	_, err = db.Updates(&gstructDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	gstructNew := new(models.Gstruct)
	gstructDB.CopyBasicFieldsToGstruct(gstructNew)

	// redeem pointers
	gstructDB.DecodePointers(backRepo, gstructNew)

	// get stage instance from DB instance, and call callback function
	gstructOld := backRepo.BackRepoGstruct.Map_GstructDBID_GstructPtr[gstructDB.ID]
	if gstructOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(gstructOld, gstructNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gstructDB
	writeJSON(w, http.StatusOK, gstructDB)
}
