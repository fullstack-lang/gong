// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/button/go/models"
	"github.com/fullstack-lang/gong/lib/button/go/orm"
)

// declaration in order to justify use of the models import
var __Button__dummysDeclaration__ models.Button
var _ = __Button__dummysDeclaration__
var __Button_time__dummyDeclaration time.Duration
var _ = __Button_time__dummyDeclaration

var mutexButton sync.Mutex

// An ButtonID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateButton
type ButtonID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ButtonInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateButton
type ButtonInput struct {
	// The Button to submit or modify
	// in: body
	Button *orm.ButtonAPI
}

// UpdateButton
//
// swagger:route PATCH /buttons/{ID} buttons updateButton
//
// # Update a button
//
// Responses:
// default: genericError
//
//	200: buttonDBResponse
func (controller *Controller) UpdateButton(w http.ResponseWriter, r *http.Request) {

	mutexButton.Lock()
	defer mutexButton.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoButton.GetDB()

	// Validate input
	var input orm.ButtonAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var buttonDB orm.ButtonDB

	// fetch the button
	_, err := db.First(&buttonDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	buttonDB.CopyBasicFieldsFromButton_WOP(&input.Button_WOP)
	buttonDB.ButtonPointersEncoding = input.ButtonPointersEncoding

	db, _ = db.Model(&buttonDB)
	_, err = db.Updates(&buttonDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	buttonNew := new(models.Button)
	buttonDB.CopyBasicFieldsToButton(buttonNew)

	// redeem pointers
	buttonDB.DecodePointers(backRepo, buttonNew)

	// get stage instance from DB instance, and call callback function
	buttonOld := backRepo.BackRepoButton.Map_ButtonDBID_ButtonPtr[buttonDB.ID]
	if buttonOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(buttonOld, buttonNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the buttonDB
	writeJSON(w, http.StatusOK, buttonDB)
}
