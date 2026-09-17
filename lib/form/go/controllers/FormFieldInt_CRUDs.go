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
var __FormFieldInt__dummysDeclaration__ models.FormFieldInt
var _ = __FormFieldInt__dummysDeclaration__
var __FormFieldInt_time__dummyDeclaration time.Duration
var _ = __FormFieldInt_time__dummyDeclaration

var mutexFormFieldInt sync.Mutex

// An FormFieldIntID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateFormFieldInt
type FormFieldIntID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormFieldIntInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateFormFieldInt
type FormFieldIntInput struct {
	// The FormFieldInt to submit or modify
	// in: body
	FormFieldInt *orm.FormFieldIntAPI
}

// UpdateFormFieldInt
//
// swagger:route PATCH /formfieldints/{ID} formfieldints updateFormFieldInt
//
// # Update a formfieldint
//
// Responses:
// default: genericError
//
//	200: formfieldintDBResponse
func (controller *Controller) UpdateFormFieldInt(w http.ResponseWriter, r *http.Request) {

	mutexFormFieldInt.Lock()
	defer mutexFormFieldInt.Unlock()

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
	db := backRepo.BackRepoFormFieldInt.GetDB()

	// Validate input
	var input orm.FormFieldIntAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formfieldintDB orm.FormFieldIntDB

	// fetch the formfieldint
	_, err := db.First(&formfieldintDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formfieldintDB.CopyBasicFieldsFromFormFieldInt_WOP(&input.FormFieldInt_WOP)
	formfieldintDB.FormFieldIntPointersEncoding = input.FormFieldIntPointersEncoding

	db, _ = db.Model(&formfieldintDB)
	_, err = db.Updates(&formfieldintDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formfieldintNew := new(models.FormFieldInt)
	formfieldintDB.CopyBasicFieldsToFormFieldInt(formfieldintNew)

	// redeem pointers
	formfieldintDB.DecodePointers(backRepo, formfieldintNew)

	// get stage instance from DB instance, and call callback function
	formfieldintOld := backRepo.BackRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntPtr[formfieldintDB.ID]
	if formfieldintOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(formfieldintOld, formfieldintNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formfieldintDB
	writeJSON(w, http.StatusOK, formfieldintDB)
}
