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
// swagger:parameters updateFormSortAssocButton
type FormSortAssocButtonID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormSortAssocButtonInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateFormSortAssocButton
type FormSortAssocButtonInput struct {
	// The FormSortAssocButton to submit or modify
	// in: body
	FormSortAssocButton *orm.FormSortAssocButtonAPI
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
