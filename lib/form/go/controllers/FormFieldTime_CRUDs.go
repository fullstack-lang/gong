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
var __FormFieldTime__dummysDeclaration__ models.FormFieldTime
var _ = __FormFieldTime__dummysDeclaration__
var __FormFieldTime_time__dummyDeclaration time.Duration
var _ = __FormFieldTime_time__dummyDeclaration

var mutexFormFieldTime sync.Mutex

// An FormFieldTimeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateFormFieldTime
type FormFieldTimeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormFieldTimeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateFormFieldTime
type FormFieldTimeInput struct {
	// The FormFieldTime to submit or modify
	// in: body
	FormFieldTime *orm.FormFieldTimeAPI
}

// UpdateFormFieldTime
//
// swagger:route PATCH /formfieldtimes/{ID} formfieldtimes updateFormFieldTime
//
// # Update a formfieldtime
//
// Responses:
// default: genericError
//
//	200: formfieldtimeDBResponse
func (controller *Controller) UpdateFormFieldTime(w http.ResponseWriter, r *http.Request) {

	mutexFormFieldTime.Lock()
	defer mutexFormFieldTime.Unlock()

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
	db := backRepo.BackRepoFormFieldTime.GetDB()

	// Validate input
	var input orm.FormFieldTimeAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formfieldtimeDB orm.FormFieldTimeDB

	// fetch the formfieldtime
	_, err := db.First(&formfieldtimeDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formfieldtimeDB.CopyBasicFieldsFromFormFieldTime_WOP(&input.FormFieldTime_WOP)
	formfieldtimeDB.FormFieldTimePointersEncoding = input.FormFieldTimePointersEncoding

	db, _ = db.Model(&formfieldtimeDB)
	_, err = db.Updates(&formfieldtimeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formfieldtimeNew := new(models.FormFieldTime)
	formfieldtimeDB.CopyBasicFieldsToFormFieldTime(formfieldtimeNew)

	// redeem pointers
	formfieldtimeDB.DecodePointers(backRepo, formfieldtimeNew)

	// get stage instance from DB instance, and call callback function
	formfieldtimeOld := backRepo.BackRepoFormFieldTime.Map_FormFieldTimeDBID_FormFieldTimePtr[formfieldtimeDB.ID]
	if formfieldtimeOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(formfieldtimeOld, formfieldtimeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formfieldtimeDB
	writeJSON(w, http.StatusOK, formfieldtimeDB)
}
