// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/button/go/models"
	"github.com/fullstack-lang/gong/lib/button/go/orm"
)

// declaration in order to justify use of the models import
var __ButtonToggle__dummysDeclaration__ models.ButtonToggle
var _ = __ButtonToggle__dummysDeclaration__
var __ButtonToggle_time__dummyDeclaration time.Duration
var _ = __ButtonToggle_time__dummyDeclaration

var mutexButtonToggle sync.Mutex

// An ButtonToggleID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getButtonToggle updateButtonToggle deleteButtonToggle
type ButtonToggleID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ButtonToggleInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postButtonToggle updateButtonToggle
type ButtonToggleInput struct {
	// The ButtonToggle to submit or modify
	// in: body
	ButtonToggle *orm.ButtonToggleAPI
}

// GetButtonToggles
//
// swagger:route GET /buttontoggles buttontoggles getButtonToggles
//
// # Get all buttontoggles
//
// Responses:
// default: genericError
//
//	200: buttontoggleDBResponse
func (controller *Controller) GetButtonToggles(w http.ResponseWriter, r *http.Request) {

	// source slice
	var buttontoggleDBs []orm.ButtonToggleDB

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetButtonToggles", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoButtonToggle.GetDB()

	_, err := db.Find(&buttontoggleDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	buttontoggleAPIs := make([]orm.ButtonToggleAPI, 0)

	// for each buttontoggle, update fields from the database nullable fields
	for idx := range buttontoggleDBs {
		buttontoggleDB := &buttontoggleDBs[idx]
		_ = buttontoggleDB
		var buttontoggleAPI orm.ButtonToggleAPI

		// insertion point for updating fields
		buttontoggleAPI.ID = buttontoggleDB.ID
		buttontoggleDB.CopyBasicFieldsToButtonToggle_WOP(&buttontoggleAPI.ButtonToggle_WOP)
		buttontoggleAPI.ButtonTogglePointersEncoding = buttontoggleDB.ButtonTogglePointersEncoding
		buttontoggleAPIs = append(buttontoggleAPIs, buttontoggleAPI)
	}

	writeJSON(w, http.StatusOK, buttontoggleAPIs)
}

// PostButtonToggle
//
// swagger:route POST /buttontoggles buttontoggles postButtonToggle
//
// Creates a buttontoggle
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostButtonToggle(w http.ResponseWriter, r *http.Request) {

	mutexButtonToggle.Lock()
	defer mutexButtonToggle.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostButtonToggles", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoButtonToggle.GetDB()

	// Validate input
	var input orm.ButtonToggleAPI

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// Create buttontoggle
	buttontoggleDB := orm.ButtonToggleDB{}
	buttontoggleDB.ButtonTogglePointersEncoding = input.ButtonTogglePointersEncoding
	buttontoggleDB.CopyBasicFieldsFromButtonToggle_WOP(&input.ButtonToggle_WOP)

	_, err = db.Create(&buttontoggleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoButtonToggle.CheckoutPhaseOneInstance(&buttontoggleDB)
	buttontoggle := backRepo.BackRepoButtonToggle.Map_ButtonToggleDBID_ButtonTogglePtr[buttontoggleDB.ID]

	if buttontoggle != nil {
		backRepo.GetStage().AfterCreateFromFront(buttontoggle)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, buttontoggleDB)
}

// GetButtonToggle
//
// swagger:route GET /buttontoggles/{ID} buttontoggles getButtonToggle
//
// Gets the details for a buttontoggle.
//
// Responses:
// default: genericError
//
//	200: buttontoggleDBResponse
func (controller *Controller) GetButtonToggle(w http.ResponseWriter, r *http.Request) {

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetButtonToggle", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoButtonToggle.GetDB()

	// Get buttontoggleDB in DB
	var buttontoggleDB orm.ButtonToggleDB
	if _, err := db.First(&buttontoggleDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	var buttontoggleAPI orm.ButtonToggleAPI
	buttontoggleAPI.ID = buttontoggleDB.ID
	buttontoggleAPI.ButtonTogglePointersEncoding = buttontoggleDB.ButtonTogglePointersEncoding
	buttontoggleDB.CopyBasicFieldsToButtonToggle_WOP(&buttontoggleAPI.ButtonToggle_WOP)

	writeJSON(w, http.StatusOK, buttontoggleAPI)
}

// UpdateButtonToggle
//
// swagger:route PATCH /buttontoggles/{ID} buttontoggles updateButtonToggle
//
// # Update a buttontoggle
//
// Responses:
// default: genericError
//
//	200: buttontoggleDBResponse
func (controller *Controller) UpdateButtonToggle(w http.ResponseWriter, r *http.Request) {

	mutexButtonToggle.Lock()
	defer mutexButtonToggle.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoButtonToggle.GetDB()

	// Validate input
	var input orm.ButtonToggleAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var buttontoggleDB orm.ButtonToggleDB

	// fetch the buttontoggle
	_, err := db.First(&buttontoggleDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	buttontoggleDB.CopyBasicFieldsFromButtonToggle_WOP(&input.ButtonToggle_WOP)
	buttontoggleDB.ButtonTogglePointersEncoding = input.ButtonTogglePointersEncoding

	db, _ = db.Model(&buttontoggleDB)
	_, err = db.Updates(&buttontoggleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	buttontoggleNew := new(models.ButtonToggle)
	buttontoggleDB.CopyBasicFieldsToButtonToggle(buttontoggleNew)

	// redeem pointers
	buttontoggleDB.DecodePointers(backRepo, buttontoggleNew)

	// get stage instance from DB instance, and call callback function
	buttontoggleOld := backRepo.BackRepoButtonToggle.Map_ButtonToggleDBID_ButtonTogglePtr[buttontoggleDB.ID]
	if buttontoggleOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(buttontoggleOld, buttontoggleNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the buttontoggleDB
	writeJSON(w, http.StatusOK, buttontoggleDB)
}

// DeleteButtonToggle
//
// swagger:route DELETE /buttontoggles/{ID} buttontoggles deleteButtonToggle
//
// # Delete a buttontoggle
//
// default: genericError
//
//	200: buttontoggleDBResponse
func (controller *Controller) DeleteButtonToggle(w http.ResponseWriter, r *http.Request) {

	mutexButtonToggle.Lock()
	defer mutexButtonToggle.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteButtonToggle", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoButtonToggle.GetDB()

	// Get model if exist
	var buttontoggleDB orm.ButtonToggleDB
	if _, err := db.First(&buttontoggleDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&buttontoggleDB)

	// get an instance (not staged) from DB instance, and call callback function
	buttontoggleDeleted := new(models.ButtonToggle)
	buttontoggleDB.CopyBasicFieldsToButtonToggle(buttontoggleDeleted)

	// get stage instance from DB instance, and call callback function
	buttontoggleStaged := backRepo.BackRepoButtonToggle.Map_ButtonToggleDBID_ButtonTogglePtr[buttontoggleDB.ID]
	if buttontoggleStaged != nil {
		backRepo.GetStage().AfterDeleteFromFront(buttontoggleStaged, buttontoggleDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, H{"data": true})
}
