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
// swagger:parameters getTone updateTone deleteTone
type ToneID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ToneInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTone updateTone
type ToneInput struct {
	// The Tone to submit or modify
	// in: body
	Tone *orm.ToneAPI
}

// GetTones
//
// swagger:route GET /tones tones getTones
//
// # Get all tones
//
// Responses:
// default: genericError
//
//	200: toneDBResponse
func (controller *Controller) GetTones(w http.ResponseWriter, r *http.Request) {

	// source slice
	var toneDBs []orm.ToneDB

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTones", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTone.GetDB()

	_, err := db.Find(&toneDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	toneAPIs := make([]orm.ToneAPI, 0)

	// for each tone, update fields from the database nullable fields
	for idx := range toneDBs {
		toneDB := &toneDBs[idx]
		_ = toneDB
		var toneAPI orm.ToneAPI

		// insertion point for updating fields
		toneAPI.ID = toneDB.ID
		toneDB.CopyBasicFieldsToTone_WOP(&toneAPI.Tone_WOP)
		toneAPI.TonePointersEncoding = toneDB.TonePointersEncoding
		toneAPIs = append(toneAPIs, toneAPI)
	}

	writeJSON(w, http.StatusOK, toneAPIs)
}

// PostTone
//
// swagger:route POST /tones tones postTone
//
// Creates a tone
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTone(w http.ResponseWriter, r *http.Request) {

	mutexTone.Lock()
	defer mutexTone.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTones", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTone.GetDB()

	// Validate input
	var input orm.ToneAPI

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// Create tone
	toneDB := orm.ToneDB{}
	toneDB.TonePointersEncoding = input.TonePointersEncoding
	toneDB.CopyBasicFieldsFromTone_WOP(&input.Tone_WOP)

	_, err = db.Create(&toneDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTone.CheckoutPhaseOneInstance(&toneDB)
	tone := backRepo.BackRepoTone.Map_ToneDBID_TonePtr[toneDB.ID]

	if tone != nil {
		backRepo.GetStage().AfterCreateFromFront(tone)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, toneDB)
}

// GetTone
//
// swagger:route GET /tones/{ID} tones getTone
//
// Gets the details for a tone.
//
// Responses:
// default: genericError
//
//	200: toneDBResponse
func (controller *Controller) GetTone(w http.ResponseWriter, r *http.Request) {

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTone", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTone.GetDB()

	// Get toneDB in DB
	var toneDB orm.ToneDB
	if _, err := db.First(&toneDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	var toneAPI orm.ToneAPI
	toneAPI.ID = toneDB.ID
	toneAPI.TonePointersEncoding = toneDB.TonePointersEncoding
	toneDB.CopyBasicFieldsToTone_WOP(&toneAPI.Tone_WOP)

	writeJSON(w, http.StatusOK, toneAPI)
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

// DeleteTone
//
// swagger:route DELETE /tones/{ID} tones deleteTone
//
// # Delete a tone
//
// default: genericError
//
//	200: toneDBResponse
func (controller *Controller) DeleteTone(w http.ResponseWriter, r *http.Request) {

	mutexTone.Lock()
	defer mutexTone.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTone", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTone.GetDB()

	// Get model if exist
	var toneDB orm.ToneDB
	if _, err := db.First(&toneDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&toneDB)

	// get an instance (not staged) from DB instance, and call callback function
	toneDeleted := new(models.Tone)
	toneDB.CopyBasicFieldsToTone(toneDeleted)

	// get stage instance from DB instance, and call callback function
	toneStaged := backRepo.BackRepoTone.Map_ToneDBID_TonePtr[toneDB.ID]
	if toneStaged != nil {
		backRepo.GetStage().AfterDeleteFromFront(toneStaged, toneDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, H{"data": true})
}
