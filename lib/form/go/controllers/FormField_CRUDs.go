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
var __FormField__dummysDeclaration__ models.FormField
var _ = __FormField__dummysDeclaration__
var __FormField_time__dummyDeclaration time.Duration
var _ = __FormField_time__dummyDeclaration

var mutexFormField sync.Mutex

// An FormFieldID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateFormField
type FormFieldID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormFieldInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateFormField
type FormFieldInput struct {
	// The FormField to submit or modify
	// in: body
	FormField *orm.FormFieldAPI
}

// UpdateFormField
//
// swagger:route PATCH /formfields/{ID} formfields updateFormField
//
// # Update a formfield
//
// Responses:
// default: genericError
//
//	200: formfieldDBResponse
func (controller *Controller) UpdateFormField(w http.ResponseWriter, r *http.Request) {

	mutexFormField.Lock()
	defer mutexFormField.Unlock()

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
	db := backRepo.BackRepoFormField.GetDB()

	// Validate input
	var input orm.FormFieldAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formfieldDB orm.FormFieldDB

	// fetch the formfield
	_, err := db.First(&formfieldDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formfieldDB.CopyBasicFieldsFromFormField_WOP(&input.FormField_WOP)
	formfieldDB.FormFieldPointersEncoding = input.FormFieldPointersEncoding

	db, _ = db.Model(&formfieldDB)
	_, err = db.Updates(&formfieldDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formfieldNew := new(models.FormField)
	formfieldDB.CopyBasicFieldsToFormField(formfieldNew)

	// redeem pointers
	formfieldDB.DecodePointers(backRepo, formfieldNew)

	// get stage instance from DB instance, and call callback function
	formfieldOld := backRepo.BackRepoFormField.Map_FormFieldDBID_FormFieldPtr[formfieldDB.ID]
	if formfieldOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(formfieldOld, formfieldNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formfieldDB
	writeJSON(w, http.StatusOK, formfieldDB)
}
