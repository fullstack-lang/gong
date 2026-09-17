// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/tone/go/models"
	"github.com/fullstack-lang/gong/lib/tone/go/orm"
)

// declaration in order to justify use of the models import
var __Freqency__dummysDeclaration__ models.Freqency
var _ = __Freqency__dummysDeclaration__
var __Freqency_time__dummyDeclaration time.Duration
var _ = __Freqency_time__dummyDeclaration

var mutexFreqency sync.Mutex

// An FreqencyID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateFreqency
type FreqencyID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FreqencyInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateFreqency
type FreqencyInput struct {
	// The Freqency to submit or modify
	// in: body
	Freqency *orm.FreqencyAPI
}

// UpdateFreqency
//
// swagger:route PATCH /freqencys/{ID} freqencys updateFreqency
//
// # Update a freqency
//
// Responses:
// default: genericError
//
//	200: freqencyDBResponse
func (controller *Controller) UpdateFreqency(w http.ResponseWriter, r *http.Request) {

	mutexFreqency.Lock()
	defer mutexFreqency.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/tone/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFreqency.GetDB()

	// Validate input
	var input orm.FreqencyAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var freqencyDB orm.FreqencyDB

	// fetch the freqency
	_, err := db.First(&freqencyDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	freqencyDB.CopyBasicFieldsFromFreqency_WOP(&input.Freqency_WOP)
	freqencyDB.FreqencyPointersEncoding = input.FreqencyPointersEncoding

	db, _ = db.Model(&freqencyDB)
	_, err = db.Updates(&freqencyDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	freqencyNew := new(models.Freqency)
	freqencyDB.CopyBasicFieldsToFreqency(freqencyNew)

	// redeem pointers
	freqencyDB.DecodePointers(backRepo, freqencyNew)

	// get stage instance from DB instance, and call callback function
	freqencyOld := backRepo.BackRepoFreqency.Map_FreqencyDBID_FreqencyPtr[freqencyDB.ID]
	if freqencyOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(freqencyOld, freqencyNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the freqencyDB
	writeJSON(w, http.StatusOK, freqencyDB)
}
