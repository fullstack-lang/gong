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
var __FormFieldSelect__dummysDeclaration__ models.FormFieldSelect
var _ = __FormFieldSelect__dummysDeclaration__
var __FormFieldSelect_time__dummyDeclaration time.Duration
var _ = __FormFieldSelect_time__dummyDeclaration

var mutexFormFieldSelect sync.Mutex

// An FormFieldSelectID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFormFieldSelect updateFormFieldSelect deleteFormFieldSelect
type FormFieldSelectID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormFieldSelectInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFormFieldSelect updateFormFieldSelect
type FormFieldSelectInput struct {
	// The FormFieldSelect to submit or modify
	// in: body
	FormFieldSelect *orm.FormFieldSelectAPI
}

// GetFormFieldSelects
//
// swagger:route GET /formfieldselects formfieldselects getFormFieldSelects
//
// # Get all formfieldselects
//
// Responses:
// default: genericError
//
//	200: formfieldselectDBResponse
func (controller *Controller) GetFormFieldSelects(w http.ResponseWriter, r *http.Request) {

	// source slice
	var formfieldselectDBs []orm.FormFieldSelectDB

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormFieldSelects", "Name", stackPath)
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
	db := backRepo.BackRepoFormFieldSelect.GetDB()

	_, err := db.Find(&formfieldselectDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	formfieldselectAPIs := make([]orm.FormFieldSelectAPI, 0)

	// for each formfieldselect, update fields from the database nullable fields
	for idx := range formfieldselectDBs {
		formfieldselectDB := &formfieldselectDBs[idx]
		_ = formfieldselectDB
		var formfieldselectAPI orm.FormFieldSelectAPI

		// insertion point for updating fields
		formfieldselectAPI.ID = formfieldselectDB.ID
		formfieldselectDB.CopyBasicFieldsToFormFieldSelect_WOP(&formfieldselectAPI.FormFieldSelect_WOP)
		formfieldselectAPI.FormFieldSelectPointersEncoding = formfieldselectDB.FormFieldSelectPointersEncoding
		formfieldselectAPIs = append(formfieldselectAPIs, formfieldselectAPI)
	}

	writeJSON(w, http.StatusOK, formfieldselectAPIs)
}

// PostFormFieldSelect
//
// swagger:route POST /formfieldselects formfieldselects postFormFieldSelect
//
// Creates a formfieldselect
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFormFieldSelect(w http.ResponseWriter, r *http.Request) {

	mutexFormFieldSelect.Lock()
	defer mutexFormFieldSelect.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFormFieldSelects", "Name", stackPath)
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
	db := backRepo.BackRepoFormFieldSelect.GetDB()

	// Validate input
	var input orm.FormFieldSelectAPI

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// Create formfieldselect
	formfieldselectDB := orm.FormFieldSelectDB{}
	formfieldselectDB.FormFieldSelectPointersEncoding = input.FormFieldSelectPointersEncoding
	formfieldselectDB.CopyBasicFieldsFromFormFieldSelect_WOP(&input.FormFieldSelect_WOP)

	_, err = db.Create(&formfieldselectDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFormFieldSelect.CheckoutPhaseOneInstance(&formfieldselectDB)
	formfieldselect := backRepo.BackRepoFormFieldSelect.Map_FormFieldSelectDBID_FormFieldSelectPtr[formfieldselectDB.ID]

	if formfieldselect != nil {
		backRepo.GetStage().AfterCreateFromFront(formfieldselect)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, formfieldselectDB)
}

// GetFormFieldSelect
//
// swagger:route GET /formfieldselects/{ID} formfieldselects getFormFieldSelect
//
// Gets the details for a formfieldselect.
//
// Responses:
// default: genericError
//
//	200: formfieldselectDBResponse
func (controller *Controller) GetFormFieldSelect(w http.ResponseWriter, r *http.Request) {

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormFieldSelect", "Name", stackPath)
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
	db := backRepo.BackRepoFormFieldSelect.GetDB()

	// Get formfieldselectDB in DB
	var formfieldselectDB orm.FormFieldSelectDB
	if _, err := db.First(&formfieldselectDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	var formfieldselectAPI orm.FormFieldSelectAPI
	formfieldselectAPI.ID = formfieldselectDB.ID
	formfieldselectAPI.FormFieldSelectPointersEncoding = formfieldselectDB.FormFieldSelectPointersEncoding
	formfieldselectDB.CopyBasicFieldsToFormFieldSelect_WOP(&formfieldselectAPI.FormFieldSelect_WOP)

	writeJSON(w, http.StatusOK, formfieldselectAPI)
}

// UpdateFormFieldSelect
//
// swagger:route PATCH /formfieldselects/{ID} formfieldselects updateFormFieldSelect
//
// # Update a formfieldselect
//
// Responses:
// default: genericError
//
//	200: formfieldselectDBResponse
func (controller *Controller) UpdateFormFieldSelect(w http.ResponseWriter, r *http.Request) {

	mutexFormFieldSelect.Lock()
	defer mutexFormFieldSelect.Unlock()

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
	db := backRepo.BackRepoFormFieldSelect.GetDB()

	// Validate input
	var input orm.FormFieldSelectAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formfieldselectDB orm.FormFieldSelectDB

	// fetch the formfieldselect
	_, err := db.First(&formfieldselectDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formfieldselectDB.CopyBasicFieldsFromFormFieldSelect_WOP(&input.FormFieldSelect_WOP)
	formfieldselectDB.FormFieldSelectPointersEncoding = input.FormFieldSelectPointersEncoding

	db, _ = db.Model(&formfieldselectDB)
	_, err = db.Updates(&formfieldselectDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formfieldselectNew := new(models.FormFieldSelect)
	formfieldselectDB.CopyBasicFieldsToFormFieldSelect(formfieldselectNew)

	// redeem pointers
	formfieldselectDB.DecodePointers(backRepo, formfieldselectNew)

	// get stage instance from DB instance, and call callback function
	formfieldselectOld := backRepo.BackRepoFormFieldSelect.Map_FormFieldSelectDBID_FormFieldSelectPtr[formfieldselectDB.ID]
	if formfieldselectOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(formfieldselectOld, formfieldselectNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formfieldselectDB
	writeJSON(w, http.StatusOK, formfieldselectDB)
}

// DeleteFormFieldSelect
//
// swagger:route DELETE /formfieldselects/{ID} formfieldselects deleteFormFieldSelect
//
// # Delete a formfieldselect
//
// default: genericError
//
//	200: formfieldselectDBResponse
func (controller *Controller) DeleteFormFieldSelect(w http.ResponseWriter, r *http.Request) {

	mutexFormFieldSelect.Lock()
	defer mutexFormFieldSelect.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFormFieldSelect", "Name", stackPath)
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
	db := backRepo.BackRepoFormFieldSelect.GetDB()

	// Get model if exist
	var formfieldselectDB orm.FormFieldSelectDB
	if _, err := db.First(&formfieldselectDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&formfieldselectDB)

	// get an instance (not staged) from DB instance, and call callback function
	formfieldselectDeleted := new(models.FormFieldSelect)
	formfieldselectDB.CopyBasicFieldsToFormFieldSelect(formfieldselectDeleted)

	// get stage instance from DB instance, and call callback function
	formfieldselectStaged := backRepo.BackRepoFormFieldSelect.Map_FormFieldSelectDBID_FormFieldSelectPtr[formfieldselectDB.ID]
	if formfieldselectStaged != nil {
		backRepo.GetStage().AfterDeleteFromFront(formfieldselectStaged, formfieldselectDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, H{"data": true})
}
