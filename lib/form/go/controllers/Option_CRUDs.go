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
var __Option__dummysDeclaration__ models.Option
var _ = __Option__dummysDeclaration__
var __Option_time__dummyDeclaration time.Duration
var _ = __Option_time__dummyDeclaration

var mutexOption sync.Mutex

// An OptionID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateOption
type OptionID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// OptionInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateOption
type OptionInput struct {
	// The Option to submit or modify
	// in: body
	Option *orm.OptionAPI
}

// UpdateOption
//
// swagger:route PATCH /options/{ID} options updateOption
//
// # Update a option
//
// Responses:
// default: genericError
//
//	200: optionDBResponse
func (controller *Controller) UpdateOption(w http.ResponseWriter, r *http.Request) {

	mutexOption.Lock()
	defer mutexOption.Unlock()

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
	db := backRepo.BackRepoOption.GetDB()

	// Validate input
	var input orm.OptionAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var optionDB orm.OptionDB

	// fetch the option
	_, err := db.First(&optionDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	optionDB.CopyBasicFieldsFromOption_WOP(&input.Option_WOP)
	optionDB.OptionPointersEncoding = input.OptionPointersEncoding

	db, _ = db.Model(&optionDB)
	_, err = db.Updates(&optionDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	optionNew := new(models.Option)
	optionDB.CopyBasicFieldsToOption(optionNew)

	// redeem pointers
	optionDB.DecodePointers(backRepo, optionNew)

	// get stage instance from DB instance, and call callback function
	optionOld := backRepo.BackRepoOption.Map_OptionDBID_OptionPtr[optionDB.ID]
	if optionOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(optionOld, optionNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the optionDB
	writeJSON(w, http.StatusOK, optionDB)
}
