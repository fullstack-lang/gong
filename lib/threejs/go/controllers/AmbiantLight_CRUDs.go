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
var __AmbiantLight__dummysDeclaration__ models.AmbiantLight
var _ = __AmbiantLight__dummysDeclaration__
var __AmbiantLight_time__dummyDeclaration time.Duration
var _ = __AmbiantLight_time__dummyDeclaration

var mutexAmbiantLight sync.Mutex

// An AmbiantLightID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAmbiantLight updateAmbiantLight deleteAmbiantLight
type AmbiantLightID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AmbiantLightInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAmbiantLight updateAmbiantLight
type AmbiantLightInput struct {
	// The AmbiantLight to submit or modify
	// in: body
	AmbiantLight *orm.AmbiantLightAPI
}

// GetAmbiantLights
//
// swagger:route GET /ambiantlights ambiantlights getAmbiantLights
//
// # Get all ambiantlights
//
// Responses:
// default: genericError
//
//	200: ambiantlightDBResponse
func (controller *Controller) GetAmbiantLights(w http.ResponseWriter, r *http.Request) {

	// source slice
	var ambiantlightDBs []orm.AmbiantLightDB

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAmbiantLights", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoAmbiantLight.GetDB()

	_, err := db.Find(&ambiantlightDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	ambiantlightAPIs := make([]orm.AmbiantLightAPI, 0)

	// for each ambiantlight, update fields from the database nullable fields
	for idx := range ambiantlightDBs {
		ambiantlightDB := &ambiantlightDBs[idx]
		_ = ambiantlightDB
		var ambiantlightAPI orm.AmbiantLightAPI

		// insertion point for updating fields
		ambiantlightAPI.ID = ambiantlightDB.ID
		ambiantlightDB.CopyBasicFieldsToAmbiantLight_WOP(&ambiantlightAPI.AmbiantLight_WOP)
		ambiantlightAPI.AmbiantLightPointersEncoding = ambiantlightDB.AmbiantLightPointersEncoding
		ambiantlightAPIs = append(ambiantlightAPIs, ambiantlightAPI)
	}

	writeJSON(w, http.StatusOK, ambiantlightAPIs)
}

// PostAmbiantLight
//
// swagger:route POST /ambiantlights ambiantlights postAmbiantLight
//
// Creates a ambiantlight
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAmbiantLight(w http.ResponseWriter, r *http.Request) {

	mutexAmbiantLight.Lock()
	defer mutexAmbiantLight.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAmbiantLights", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoAmbiantLight.GetDB()

	// Validate input
	var input orm.AmbiantLightAPI

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// Create ambiantlight
	ambiantlightDB := orm.AmbiantLightDB{}
	ambiantlightDB.AmbiantLightPointersEncoding = input.AmbiantLightPointersEncoding
	ambiantlightDB.CopyBasicFieldsFromAmbiantLight_WOP(&input.AmbiantLight_WOP)

	_, err = db.Create(&ambiantlightDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAmbiantLight.CheckoutPhaseOneInstance(&ambiantlightDB)
	ambiantlight := backRepo.BackRepoAmbiantLight.Map_AmbiantLightDBID_AmbiantLightPtr[ambiantlightDB.ID]

	if ambiantlight != nil {
		backRepo.GetStage().AfterCreateFromFront(ambiantlight)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, ambiantlightDB)
}

// GetAmbiantLight
//
// swagger:route GET /ambiantlights/{ID} ambiantlights getAmbiantLight
//
// Gets the details for a ambiantlight.
//
// Responses:
// default: genericError
//
//	200: ambiantlightDBResponse
func (controller *Controller) GetAmbiantLight(w http.ResponseWriter, r *http.Request) {

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAmbiantLight", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoAmbiantLight.GetDB()

	// Get ambiantlightDB in DB
	var ambiantlightDB orm.AmbiantLightDB
	if _, err := db.First(&ambiantlightDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	var ambiantlightAPI orm.AmbiantLightAPI
	ambiantlightAPI.ID = ambiantlightDB.ID
	ambiantlightAPI.AmbiantLightPointersEncoding = ambiantlightDB.AmbiantLightPointersEncoding
	ambiantlightDB.CopyBasicFieldsToAmbiantLight_WOP(&ambiantlightAPI.AmbiantLight_WOP)

	writeJSON(w, http.StatusOK, ambiantlightAPI)
}

// UpdateAmbiantLight
//
// swagger:route PATCH /ambiantlights/{ID} ambiantlights updateAmbiantLight
//
// # Update a ambiantlight
//
// Responses:
// default: genericError
//
//	200: ambiantlightDBResponse
func (controller *Controller) UpdateAmbiantLight(w http.ResponseWriter, r *http.Request) {

	mutexAmbiantLight.Lock()
	defer mutexAmbiantLight.Unlock()

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
	db := backRepo.BackRepoAmbiantLight.GetDB()

	// Validate input
	var input orm.AmbiantLightAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var ambiantlightDB orm.AmbiantLightDB

	// fetch the ambiantlight
	_, err := db.First(&ambiantlightDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	ambiantlightDB.CopyBasicFieldsFromAmbiantLight_WOP(&input.AmbiantLight_WOP)
	ambiantlightDB.AmbiantLightPointersEncoding = input.AmbiantLightPointersEncoding

	db, _ = db.Model(&ambiantlightDB)
	_, err = db.Updates(&ambiantlightDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	ambiantlightNew := new(models.AmbiantLight)
	ambiantlightDB.CopyBasicFieldsToAmbiantLight(ambiantlightNew)

	// redeem pointers
	ambiantlightDB.DecodePointers(backRepo, ambiantlightNew)

	// get stage instance from DB instance, and call callback function
	ambiantlightOld := backRepo.BackRepoAmbiantLight.Map_AmbiantLightDBID_AmbiantLightPtr[ambiantlightDB.ID]
	if ambiantlightOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(ambiantlightOld, ambiantlightNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the ambiantlightDB
	writeJSON(w, http.StatusOK, ambiantlightDB)
}

// DeleteAmbiantLight
//
// swagger:route DELETE /ambiantlights/{ID} ambiantlights deleteAmbiantLight
//
// # Delete a ambiantlight
//
// default: genericError
//
//	200: ambiantlightDBResponse
func (controller *Controller) DeleteAmbiantLight(w http.ResponseWriter, r *http.Request) {

	mutexAmbiantLight.Lock()
	defer mutexAmbiantLight.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAmbiantLight", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoAmbiantLight.GetDB()

	// Get model if exist
	var ambiantlightDB orm.AmbiantLightDB
	if _, err := db.First(&ambiantlightDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&ambiantlightDB)

	// get an instance (not staged) from DB instance, and call callback function
	ambiantlightDeleted := new(models.AmbiantLight)
	ambiantlightDB.CopyBasicFieldsToAmbiantLight(ambiantlightDeleted)

	// get stage instance from DB instance, and call callback function
	ambiantlightStaged := backRepo.BackRepoAmbiantLight.Map_AmbiantLightDBID_AmbiantLightPtr[ambiantlightDB.ID]
	if ambiantlightStaged != nil {
		backRepo.GetStage().AfterDeleteFromFront(ambiantlightStaged, ambiantlightDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, H{"data": true})
}
