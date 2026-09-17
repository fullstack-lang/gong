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
var __Dstruct__dummysDeclaration__ models.Dstruct
var _ = __Dstruct__dummysDeclaration__
var __Dstruct_time__dummyDeclaration time.Duration
var _ = __Dstruct_time__dummyDeclaration

var mutexDstruct sync.Mutex

// An DstructID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateDstruct
type DstructID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DstructInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateDstruct
type DstructInput struct {
	// The Dstruct to submit or modify
	// in: body
	Dstruct *orm.DstructAPI
}

// UpdateDstruct
//
// swagger:route PATCH /dstructs/{ID} dstructs updateDstruct
//
// # Update a dstruct
//
// Responses:
// default: genericError
//
//	200: dstructDBResponse
func (controller *Controller) UpdateDstruct(w http.ResponseWriter, r *http.Request) {

	mutexDstruct.Lock()
	defer mutexDstruct.Unlock()

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
	db := backRepo.BackRepoDstruct.GetDB()

	// Validate input
	var input orm.DstructAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var dstructDB orm.DstructDB

	// fetch the dstruct
	_, err := db.First(&dstructDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	dstructDB.CopyBasicFieldsFromDstruct_WOP(&input.Dstruct_WOP)
	dstructDB.DstructPointersEncoding = input.DstructPointersEncoding

	db, _ = db.Model(&dstructDB)
	_, err = db.Updates(&dstructDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	dstructNew := new(models.Dstruct)
	dstructDB.CopyBasicFieldsToDstruct(dstructNew)

	// redeem pointers
	dstructDB.DecodePointers(backRepo, dstructNew)

	// get stage instance from DB instance, and call callback function
	dstructOld := backRepo.BackRepoDstruct.Map_DstructDBID_DstructPtr[dstructDB.ID]
	if dstructOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(dstructOld, dstructNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the dstructDB
	writeJSON(w, http.StatusOK, dstructDB)
}
