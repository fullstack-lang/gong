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
var __Tone__dummysDeclaration__ models.Tone
var _ = __Tone__dummysDeclaration__
var __Tone_time__dummyDeclaration time.Duration
var _ = __Tone_time__dummyDeclaration

var mutexTone sync.Mutex

// An ToneID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateTone
type ToneID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ToneInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateTone
type ToneInput struct {
	// The Tone to submit or modify
	// in: body
	Tone *orm.ToneAPI
}

// UpdateTone
//
// swagger:route PATCH /tones/{ID} tones updateTone
//
// # Update a tone
//
// Responses:
// default: genericError
//
//	200: toneDBResponse
func (controller *Controller) UpdateTone(w http.ResponseWriter, r *http.Request) {

	mutexTone.Lock()
	defer mutexTone.Unlock()

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
	db := backRepo.BackRepoTone.GetDB()

	// Validate input
	var input orm.ToneAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var toneDB orm.ToneDB

	// fetch the tone
	_, err := db.First(&toneDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	toneDB.CopyBasicFieldsFromTone_WOP(&input.Tone_WOP)
	toneDB.TonePointersEncoding = input.TonePointersEncoding

	db, _ = db.Model(&toneDB)
	_, err = db.Updates(&toneDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	toneNew := new(models.Tone)
	toneDB.CopyBasicFieldsToTone(toneNew)

	// redeem pointers
	toneDB.DecodePointers(backRepo, toneNew)

	// get stage instance from DB instance, and call callback function
	toneOld := backRepo.BackRepoTone.Map_ToneDBID_TonePtr[toneDB.ID]
	if toneOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(toneOld, toneNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the toneDB
	writeJSON(w, http.StatusOK, toneDB)
}
