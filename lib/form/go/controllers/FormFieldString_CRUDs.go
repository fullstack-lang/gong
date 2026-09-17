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
var __FormFieldString__dummysDeclaration__ models.FormFieldString
var _ = __FormFieldString__dummysDeclaration__
var __FormFieldString_time__dummyDeclaration time.Duration
var _ = __FormFieldString_time__dummyDeclaration

var mutexFormFieldString sync.Mutex

// An FormFieldStringID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateFormFieldString
type FormFieldStringID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormFieldStringInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateFormFieldString
type FormFieldStringInput struct {
	// The FormFieldString to submit or modify
	// in: body
	FormFieldString *orm.FormFieldStringAPI
}

// UpdateFormFieldString
//
// swagger:route PATCH /formfieldstrings/{ID} formfieldstrings updateFormFieldString
//
// # Update a formfieldstring
//
// Responses:
// default: genericError
//
//	200: formfieldstringDBResponse
func (controller *Controller) UpdateFormFieldString(w http.ResponseWriter, r *http.Request) {

	mutexFormFieldString.Lock()
	defer mutexFormFieldString.Unlock()

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
	db := backRepo.BackRepoFormFieldString.GetDB()

	// Validate input
	var input orm.FormFieldStringAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formfieldstringDB orm.FormFieldStringDB

	// fetch the formfieldstring
	_, err := db.First(&formfieldstringDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formfieldstringDB.CopyBasicFieldsFromFormFieldString_WOP(&input.FormFieldString_WOP)
	formfieldstringDB.FormFieldStringPointersEncoding = input.FormFieldStringPointersEncoding

	db, _ = db.Model(&formfieldstringDB)
	_, err = db.Updates(&formfieldstringDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formfieldstringNew := new(models.FormFieldString)
	formfieldstringDB.CopyBasicFieldsToFormFieldString(formfieldstringNew)

	// redeem pointers
	formfieldstringDB.DecodePointers(backRepo, formfieldstringNew)

	// get stage instance from DB instance, and call callback function
	formfieldstringOld := backRepo.BackRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringPtr[formfieldstringDB.ID]
	if formfieldstringOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(formfieldstringOld, formfieldstringNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formfieldstringDB
	writeJSON(w, http.StatusOK, formfieldstringDB)
}
