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
var __Astruct__dummysDeclaration__ models.Astruct
var _ = __Astruct__dummysDeclaration__
var __Astruct_time__dummyDeclaration time.Duration
var _ = __Astruct_time__dummyDeclaration

var mutexAstruct sync.Mutex

// An AstructID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateAstruct
type AstructID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AstructInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateAstruct
type AstructInput struct {
	// The Astruct to submit or modify
	// in: body
	Astruct *orm.AstructAPI
}

// UpdateAstruct
//
// swagger:route PATCH /astructs/{ID} astructs updateAstruct
//
// # Update a astruct
//
// Responses:
// default: genericError
//
//	200: astructDBResponse
func (controller *Controller) UpdateAstruct(w http.ResponseWriter, r *http.Request) {

	mutexAstruct.Lock()
	defer mutexAstruct.Unlock()

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
	db := backRepo.BackRepoAstruct.GetDB()

	// Validate input
	var input orm.AstructAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var astructDB orm.AstructDB

	// fetch the astruct
	_, err := db.First(&astructDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	astructDB.CopyBasicFieldsFromAstruct_WOP(&input.Astruct_WOP)
	astructDB.AstructPointersEncoding = input.AstructPointersEncoding

	db, _ = db.Model(&astructDB)
	_, err = db.Updates(&astructDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	astructNew := new(models.Astruct)
	astructDB.CopyBasicFieldsToAstruct(astructNew)

	// redeem pointers
	astructDB.DecodePointers(backRepo, astructNew)

	// get stage instance from DB instance, and call callback function
	astructOld := backRepo.BackRepoAstruct.Map_AstructDBID_AstructPtr[astructDB.ID]
	if astructOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(astructOld, astructNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the astructDB
	writeJSON(w, http.StatusOK, astructDB)
}
