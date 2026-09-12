// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/form/go/models"
	"github.com/fullstack-lang/gong/lib/form/go/orm"
)

// declaration in order to justify use of the models import
var __FormSortAssocButton__dummysDeclaration__ models.FormSortAssocButton
var _ = __FormSortAssocButton__dummysDeclaration__
var __FormSortAssocButton_time__dummyDeclaration time.Duration
var _ = __FormSortAssocButton_time__dummyDeclaration

var mutexFormSortAssocButton sync.Mutex

// An FormSortAssocButtonID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFormSortAssocButton updateFormSortAssocButton deleteFormSortAssocButton
type FormSortAssocButtonID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormSortAssocButtonInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFormSortAssocButton updateFormSortAssocButton
type FormSortAssocButtonInput struct {
	// The FormSortAssocButton to submit or modify
	// in: body
	FormSortAssocButton *orm.FormSortAssocButtonAPI
}

// GetFormSortAssocButtons
//
// swagger:route GET /formsortassocbuttons formsortassocbuttons getFormSortAssocButtons
//
// # Get all formsortassocbuttons
//
// Responses:
// default: genericError
//
//	200: formsortassocbuttonDBResponse
func (controller *Controller) GetFormSortAssocButtons(w http.ResponseWriter, r *http.Request) {

	// source slice
	var formsortassocbuttonDBs []orm.FormSortAssocButtonDB

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormSortAssocButtons", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/form/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFormSortAssocButton.GetDB()

	_, err := db.Find(&formsortassocbuttonDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	formsortassocbuttonAPIs := make([]orm.FormSortAssocButtonAPI, 0)

	// for each formsortassocbutton, update fields from the database nullable fields
	for idx := range formsortassocbuttonDBs {
		formsortassocbuttonDB := &formsortassocbuttonDBs[idx]
		_ = formsortassocbuttonDB
		var formsortassocbuttonAPI orm.FormSortAssocButtonAPI

		// insertion point for updating fields
		formsortassocbuttonAPI.ID = formsortassocbuttonDB.ID
		formsortassocbuttonDB.CopyBasicFieldsToFormSortAssocButton_WOP(&formsortassocbuttonAPI.FormSortAssocButton_WOP)
		formsortassocbuttonAPI.FormSortAssocButtonPointersEncoding = formsortassocbuttonDB.FormSortAssocButtonPointersEncoding
		formsortassocbuttonAPIs = append(formsortassocbuttonAPIs, formsortassocbuttonAPI)
	}

	writeJSON(w, http.StatusOK, formsortassocbuttonAPIs)
}

// PostFormSortAssocButton
//
// swagger:route POST /formsortassocbuttons formsortassocbuttons postFormSortAssocButton
//
// Creates a formsortassocbutton
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFormSortAssocButton(w http.ResponseWriter, r *http.Request) {

	mutexFormSortAssocButton.Lock()
	defer mutexFormSortAssocButton.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFormSortAssocButtons", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/form/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFormSortAssocButton.GetDB()

	// Validate input
	var input orm.FormSortAssocButtonAPI

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// Create formsortassocbutton
	formsortassocbuttonDB := orm.FormSortAssocButtonDB{}
	formsortassocbuttonDB.FormSortAssocButtonPointersEncoding = input.FormSortAssocButtonPointersEncoding
	formsortassocbuttonDB.CopyBasicFieldsFromFormSortAssocButton_WOP(&input.FormSortAssocButton_WOP)

	_, err = db.Create(&formsortassocbuttonDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFormSortAssocButton.CheckoutPhaseOneInstance(&formsortassocbuttonDB)
	formsortassocbutton := backRepo.BackRepoFormSortAssocButton.Map_FormSortAssocButtonDBID_FormSortAssocButtonPtr[formsortassocbuttonDB.ID]

	if formsortassocbutton != nil {
		backRepo.GetStage().AfterCreateFromFront(formsortassocbutton)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, formsortassocbuttonDB)
}

// GetFormSortAssocButton
//
// swagger:route GET /formsortassocbuttons/{ID} formsortassocbuttons getFormSortAssocButton
//
// Gets the details for a formsortassocbutton.
//
// Responses:
// default: genericError
//
//	200: formsortassocbuttonDBResponse
func (controller *Controller) GetFormSortAssocButton(w http.ResponseWriter, r *http.Request) {

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormSortAssocButton", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/form/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFormSortAssocButton.GetDB()

	// Get formsortassocbuttonDB in DB
	var formsortassocbuttonDB orm.FormSortAssocButtonDB
	if _, err := db.First(&formsortassocbuttonDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	var formsortassocbuttonAPI orm.FormSortAssocButtonAPI
	formsortassocbuttonAPI.ID = formsortassocbuttonDB.ID
	formsortassocbuttonAPI.FormSortAssocButtonPointersEncoding = formsortassocbuttonDB.FormSortAssocButtonPointersEncoding
	formsortassocbuttonDB.CopyBasicFieldsToFormSortAssocButton_WOP(&formsortassocbuttonAPI.FormSortAssocButton_WOP)

	writeJSON(w, http.StatusOK, formsortassocbuttonAPI)
}

// UpdateFormSortAssocButton
//
// swagger:route PATCH /formsortassocbuttons/{ID} formsortassocbuttons updateFormSortAssocButton
//
// # Update a formsortassocbutton
//
// Responses:
// default: genericError
//
//	200: formsortassocbuttonDBResponse
func (controller *Controller) UpdateFormSortAssocButton(w http.ResponseWriter, r *http.Request) {

	mutexFormSortAssocButton.Lock()
	defer mutexFormSortAssocButton.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/form/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFormSortAssocButton.GetDB()

	// Validate input
	var input orm.FormSortAssocButtonAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formsortassocbuttonDB orm.FormSortAssocButtonDB

	// fetch the formsortassocbutton
	_, err := db.First(&formsortassocbuttonDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formsortassocbuttonDB.CopyBasicFieldsFromFormSortAssocButton_WOP(&input.FormSortAssocButton_WOP)
	formsortassocbuttonDB.FormSortAssocButtonPointersEncoding = input.FormSortAssocButtonPointersEncoding

	db, _ = db.Model(&formsortassocbuttonDB)
	_, err = db.Updates(&formsortassocbuttonDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formsortassocbuttonNew := new(models.FormSortAssocButton)
	formsortassocbuttonDB.CopyBasicFieldsToFormSortAssocButton(formsortassocbuttonNew)

	// redeem pointers
	formsortassocbuttonDB.DecodePointers(backRepo, formsortassocbuttonNew)

	// get stage instance from DB instance, and call callback function
	formsortassocbuttonOld := backRepo.BackRepoFormSortAssocButton.Map_FormSortAssocButtonDBID_FormSortAssocButtonPtr[formsortassocbuttonDB.ID]
	if formsortassocbuttonOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(formsortassocbuttonOld, formsortassocbuttonNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formsortassocbuttonDB
	writeJSON(w, http.StatusOK, formsortassocbuttonDB)
}

// DeleteFormSortAssocButton
//
// swagger:route DELETE /formsortassocbuttons/{ID} formsortassocbuttons deleteFormSortAssocButton
//
// # Delete a formsortassocbutton
//
// default: genericError
//
//	200: formsortassocbuttonDBResponse
func (controller *Controller) DeleteFormSortAssocButton(w http.ResponseWriter, r *http.Request) {

	mutexFormSortAssocButton.Lock()
	defer mutexFormSortAssocButton.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFormSortAssocButton", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/form/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFormSortAssocButton.GetDB()

	// Get model if exist
	var formsortassocbuttonDB orm.FormSortAssocButtonDB
	if _, err := db.First(&formsortassocbuttonDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&formsortassocbuttonDB)

	// get an instance (not staged) from DB instance, and call callback function
	formsortassocbuttonDeleted := new(models.FormSortAssocButton)
	formsortassocbuttonDB.CopyBasicFieldsToFormSortAssocButton(formsortassocbuttonDeleted)

	// get stage instance from DB instance, and call callback function
	formsortassocbuttonStaged := backRepo.BackRepoFormSortAssocButton.Map_FormSortAssocButtonDBID_FormSortAssocButtonPtr[formsortassocbuttonDB.ID]
	if formsortassocbuttonStaged != nil {
		backRepo.GetStage().AfterDeleteFromFront(formsortassocbuttonStaged, formsortassocbuttonDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, H{"data": true})
}
