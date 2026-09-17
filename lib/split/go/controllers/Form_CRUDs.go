// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/split/go/models"
	"github.com/fullstack-lang/gong/lib/split/go/orm"
)

// declaration in order to justify use of the models import
var __Form__dummysDeclaration__ models.Form
var _ = __Form__dummysDeclaration__
var __Form_time__dummyDeclaration time.Duration
var _ = __Form_time__dummyDeclaration

var mutexForm sync.Mutex

// An FormID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateForm
type FormID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateForm
type FormInput struct {
	// The Form to submit or modify
	// in: body
	Form *orm.FormAPI
}

// UpdateForm
//
// swagger:route PATCH /forms/{ID} forms updateForm
//
// # Update a form
//
// Responses:
// default: genericError
//
//	200: formDBResponse
func (controller *Controller) UpdateForm(w http.ResponseWriter, r *http.Request) {

	mutexForm.Lock()
	defer mutexForm.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoForm.GetDB()

	// Validate input
	var input orm.FormAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formDB orm.FormDB

	// fetch the form
	_, err := db.First(&formDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formDB.CopyBasicFieldsFromForm_WOP(&input.Form_WOP)
	formDB.FormPointersEncoding = input.FormPointersEncoding

	db, _ = db.Model(&formDB)
	_, err = db.Updates(&formDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formNew := new(models.Form)
	formDB.CopyBasicFieldsToForm(formNew)

	// redeem pointers
	formDB.DecodePointers(backRepo, formNew)

	// get stage instance from DB instance, and call callback function
	formOld := backRepo.BackRepoForm.Map_FormDBID_FormPtr[formDB.ID]
	if formOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(formOld, formNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formDB
	writeJSON(w, http.StatusOK, formDB)
}
