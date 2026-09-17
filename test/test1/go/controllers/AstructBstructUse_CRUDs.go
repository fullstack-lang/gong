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
var __AstructBstructUse__dummysDeclaration__ models.AstructBstructUse
var _ = __AstructBstructUse__dummysDeclaration__
var __AstructBstructUse_time__dummyDeclaration time.Duration
var _ = __AstructBstructUse_time__dummyDeclaration

var mutexAstructBstructUse sync.Mutex

// An AstructBstructUseID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateAstructBstructUse
type AstructBstructUseID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AstructBstructUseInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateAstructBstructUse
type AstructBstructUseInput struct {
	// The AstructBstructUse to submit or modify
	// in: body
	AstructBstructUse *orm.AstructBstructUseAPI
}

// UpdateAstructBstructUse
//
// swagger:route PATCH /astructbstructuses/{ID} astructbstructuses updateAstructBstructUse
//
// # Update a astructbstructuse
//
// Responses:
// default: genericError
//
//	200: astructbstructuseDBResponse
func (controller *Controller) UpdateAstructBstructUse(w http.ResponseWriter, r *http.Request) {

	mutexAstructBstructUse.Lock()
	defer mutexAstructBstructUse.Unlock()

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
	db := backRepo.BackRepoAstructBstructUse.GetDB()

	// Validate input
	var input orm.AstructBstructUseAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var astructbstructuseDB orm.AstructBstructUseDB

	// fetch the astructbstructuse
	_, err := db.First(&astructbstructuseDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	astructbstructuseDB.CopyBasicFieldsFromAstructBstructUse_WOP(&input.AstructBstructUse_WOP)
	astructbstructuseDB.AstructBstructUsePointersEncoding = input.AstructBstructUsePointersEncoding

	db, _ = db.Model(&astructbstructuseDB)
	_, err = db.Updates(&astructbstructuseDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	astructbstructuseNew := new(models.AstructBstructUse)
	astructbstructuseDB.CopyBasicFieldsToAstructBstructUse(astructbstructuseNew)

	// redeem pointers
	astructbstructuseDB.DecodePointers(backRepo, astructbstructuseNew)

	// get stage instance from DB instance, and call callback function
	astructbstructuseOld := backRepo.BackRepoAstructBstructUse.Map_AstructBstructUseDBID_AstructBstructUsePtr[astructbstructuseDB.ID]
	if astructbstructuseOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(astructbstructuseOld, astructbstructuseNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the astructbstructuseDB
	writeJSON(w, http.StatusOK, astructbstructuseDB)
}
