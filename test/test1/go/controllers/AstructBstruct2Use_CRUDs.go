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
var __AstructBstruct2Use__dummysDeclaration__ models.AstructBstruct2Use
var _ = __AstructBstruct2Use__dummysDeclaration__
var __AstructBstruct2Use_time__dummyDeclaration time.Duration
var _ = __AstructBstruct2Use_time__dummyDeclaration

var mutexAstructBstruct2Use sync.Mutex

// An AstructBstruct2UseID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateAstructBstruct2Use
type AstructBstruct2UseID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AstructBstruct2UseInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateAstructBstruct2Use
type AstructBstruct2UseInput struct {
	// The AstructBstruct2Use to submit or modify
	// in: body
	AstructBstruct2Use *orm.AstructBstruct2UseAPI
}

// UpdateAstructBstruct2Use
//
// swagger:route PATCH /astructbstruct2uses/{ID} astructbstruct2uses updateAstructBstruct2Use
//
// # Update a astructbstruct2use
//
// Responses:
// default: genericError
//
//	200: astructbstruct2useDBResponse
func (controller *Controller) UpdateAstructBstruct2Use(w http.ResponseWriter, r *http.Request) {

	mutexAstructBstruct2Use.Lock()
	defer mutexAstructBstruct2Use.Unlock()

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
	db := backRepo.BackRepoAstructBstruct2Use.GetDB()

	// Validate input
	var input orm.AstructBstruct2UseAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var astructbstruct2useDB orm.AstructBstruct2UseDB

	// fetch the astructbstruct2use
	_, err := db.First(&astructbstruct2useDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	astructbstruct2useDB.CopyBasicFieldsFromAstructBstruct2Use_WOP(&input.AstructBstruct2Use_WOP)
	astructbstruct2useDB.AstructBstruct2UsePointersEncoding = input.AstructBstruct2UsePointersEncoding

	db, _ = db.Model(&astructbstruct2useDB)
	_, err = db.Updates(&astructbstruct2useDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	astructbstruct2useNew := new(models.AstructBstruct2Use)
	astructbstruct2useDB.CopyBasicFieldsToAstructBstruct2Use(astructbstruct2useNew)

	// redeem pointers
	astructbstruct2useDB.DecodePointers(backRepo, astructbstruct2useNew)

	// get stage instance from DB instance, and call callback function
	astructbstruct2useOld := backRepo.BackRepoAstructBstruct2Use.Map_AstructBstruct2UseDBID_AstructBstruct2UsePtr[astructbstruct2useDB.ID]
	if astructbstruct2useOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(astructbstruct2useOld, astructbstruct2useNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the astructbstruct2useDB
	writeJSON(w, http.StatusOK, astructbstruct2useDB)
}
