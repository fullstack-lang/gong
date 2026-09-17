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
var __FormEditAssocButton__dummysDeclaration__ models.FormEditAssocButton
var _ = __FormEditAssocButton__dummysDeclaration__
var __FormEditAssocButton_time__dummyDeclaration time.Duration
var _ = __FormEditAssocButton_time__dummyDeclaration

var mutexFormEditAssocButton sync.Mutex

// An FormEditAssocButtonID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateFormEditAssocButton
type FormEditAssocButtonID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormEditAssocButtonInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateFormEditAssocButton
type FormEditAssocButtonInput struct {
	// The FormEditAssocButton to submit or modify
	// in: body
	FormEditAssocButton *orm.FormEditAssocButtonAPI
}

// UpdateFormEditAssocButton
//
// swagger:route PATCH /formeditassocbuttons/{ID} formeditassocbuttons updateFormEditAssocButton
//
// # Update a formeditassocbutton
//
// Responses:
// default: genericError
//
//	200: formeditassocbuttonDBResponse
func (controller *Controller) UpdateFormEditAssocButton(w http.ResponseWriter, r *http.Request) {

	mutexFormEditAssocButton.Lock()
	defer mutexFormEditAssocButton.Unlock()

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
	db := backRepo.BackRepoFormEditAssocButton.GetDB()

	// Validate input
	var input orm.FormEditAssocButtonAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formeditassocbuttonDB orm.FormEditAssocButtonDB

	// fetch the formeditassocbutton
	_, err := db.First(&formeditassocbuttonDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formeditassocbuttonDB.CopyBasicFieldsFromFormEditAssocButton_WOP(&input.FormEditAssocButton_WOP)
	formeditassocbuttonDB.FormEditAssocButtonPointersEncoding = input.FormEditAssocButtonPointersEncoding

	db, _ = db.Model(&formeditassocbuttonDB)
	_, err = db.Updates(&formeditassocbuttonDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formeditassocbuttonNew := new(models.FormEditAssocButton)
	formeditassocbuttonDB.CopyBasicFieldsToFormEditAssocButton(formeditassocbuttonNew)

	// redeem pointers
	formeditassocbuttonDB.DecodePointers(backRepo, formeditassocbuttonNew)

	// get stage instance from DB instance, and call callback function
	formeditassocbuttonOld := backRepo.BackRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonPtr[formeditassocbuttonDB.ID]
	if formeditassocbuttonOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(formeditassocbuttonOld, formeditassocbuttonNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formeditassocbuttonDB
	writeJSON(w, http.StatusOK, formeditassocbuttonDB)
}
