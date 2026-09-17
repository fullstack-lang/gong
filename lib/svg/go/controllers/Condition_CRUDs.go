// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
	"github.com/fullstack-lang/gong/lib/svg/go/orm"
)

// declaration in order to justify use of the models import
var __Condition__dummysDeclaration__ models.Condition
var _ = __Condition__dummysDeclaration__
var __Condition_time__dummyDeclaration time.Duration
var _ = __Condition_time__dummyDeclaration

var mutexCondition sync.Mutex

// An ConditionID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateCondition
type ConditionID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ConditionInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateCondition
type ConditionInput struct {
	// The Condition to submit or modify
	// in: body
	Condition *orm.ConditionAPI
}

// UpdateCondition
//
// swagger:route PATCH /conditions/{ID} conditions updateCondition
//
// # Update a condition
//
// Responses:
// default: genericError
//
//	200: conditionDBResponse
func (controller *Controller) UpdateCondition(w http.ResponseWriter, r *http.Request) {

	mutexCondition.Lock()
	defer mutexCondition.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoCondition.GetDB()

	// Validate input
	var input orm.ConditionAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var conditionDB orm.ConditionDB

	// fetch the condition
	_, err := db.First(&conditionDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	conditionDB.CopyBasicFieldsFromCondition_WOP(&input.Condition_WOP)
	conditionDB.ConditionPointersEncoding = input.ConditionPointersEncoding

	db, _ = db.Model(&conditionDB)
	_, err = db.Updates(&conditionDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	conditionNew := new(models.Condition)
	conditionDB.CopyBasicFieldsToCondition(conditionNew)

	// redeem pointers
	conditionDB.DecodePointers(backRepo, conditionNew)

	// get stage instance from DB instance, and call callback function
	conditionOld := backRepo.BackRepoCondition.Map_ConditionDBID_ConditionPtr[conditionDB.ID]
	if conditionOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(conditionOld, conditionNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the conditionDB
	writeJSON(w, http.StatusOK, conditionDB)
}
