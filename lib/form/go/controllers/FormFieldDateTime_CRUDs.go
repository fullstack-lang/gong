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
var __FormFieldDateTime__dummysDeclaration__ models.FormFieldDateTime
var _ = __FormFieldDateTime__dummysDeclaration__
var __FormFieldDateTime_time__dummyDeclaration time.Duration
var _ = __FormFieldDateTime_time__dummyDeclaration

var mutexFormFieldDateTime sync.Mutex

// An FormFieldDateTimeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateFormFieldDateTime
type FormFieldDateTimeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormFieldDateTimeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateFormFieldDateTime
type FormFieldDateTimeInput struct {
	// The FormFieldDateTime to submit or modify
	// in: body
	FormFieldDateTime *orm.FormFieldDateTimeAPI
}

// UpdateFormFieldDateTime
//
// swagger:route PATCH /formfielddatetimes/{ID} formfielddatetimes updateFormFieldDateTime
//
// # Update a formfielddatetime
//
// Responses:
// default: genericError
//
//	200: formfielddatetimeDBResponse
func (controller *Controller) UpdateFormFieldDateTime(w http.ResponseWriter, r *http.Request) {

	mutexFormFieldDateTime.Lock()
	defer mutexFormFieldDateTime.Unlock()

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
	db := backRepo.BackRepoFormFieldDateTime.GetDB()

	// Validate input
	var input orm.FormFieldDateTimeAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formfielddatetimeDB orm.FormFieldDateTimeDB

	// fetch the formfielddatetime
	_, err := db.First(&formfielddatetimeDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formfielddatetimeDB.CopyBasicFieldsFromFormFieldDateTime_WOP(&input.FormFieldDateTime_WOP)
	formfielddatetimeDB.FormFieldDateTimePointersEncoding = input.FormFieldDateTimePointersEncoding

	db, _ = db.Model(&formfielddatetimeDB)
	_, err = db.Updates(&formfielddatetimeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formfielddatetimeNew := new(models.FormFieldDateTime)
	formfielddatetimeDB.CopyBasicFieldsToFormFieldDateTime(formfielddatetimeNew)

	// redeem pointers
	formfielddatetimeDB.DecodePointers(backRepo, formfielddatetimeNew)

	// get stage instance from DB instance, and call callback function
	formfielddatetimeOld := backRepo.BackRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimePtr[formfielddatetimeDB.ID]
	if formfielddatetimeOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(formfielddatetimeOld, formfielddatetimeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formfielddatetimeDB
	writeJSON(w, http.StatusOK, formfielddatetimeDB)
}
