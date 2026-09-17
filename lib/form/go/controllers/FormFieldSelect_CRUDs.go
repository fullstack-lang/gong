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
// swagger:parameters updateFormFieldSelect
type FormFieldSelectID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormFieldSelectInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateFormFieldSelect
type FormFieldSelectInput struct {
	// The FormFieldSelect to submit or modify
	// in: body
	FormFieldSelect *orm.FormFieldSelectAPI
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
