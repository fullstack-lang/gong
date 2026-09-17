// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/sim/go/models"
	"github.com/fullstack-lang/gong/lib/sim/go/orm"
)

// declaration in order to justify use of the models import
var __DummyAgent__dummysDeclaration__ models.DummyAgent
var _ = __DummyAgent__dummysDeclaration__
var __DummyAgent_time__dummyDeclaration time.Duration
var _ = __DummyAgent_time__dummyDeclaration

var mutexDummyAgent sync.Mutex

// An DummyAgentID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateDummyAgent
type DummyAgentID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DummyAgentInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateDummyAgent
type DummyAgentInput struct {
	// The DummyAgent to submit or modify
	// in: body
	DummyAgent *orm.DummyAgentAPI
}

// UpdateDummyAgent
//
// swagger:route PATCH /dummyagents/{ID} dummyagents updateDummyAgent
//
// # Update a dummyagent
//
// Responses:
// default: genericError
//
//	200: dummyagentDBResponse
func (controller *Controller) UpdateDummyAgent(w http.ResponseWriter, r *http.Request) {

	mutexDummyAgent.Lock()
	defer mutexDummyAgent.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoDummyAgent.GetDB()

	// Validate input
	var input orm.DummyAgentAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var dummyagentDB orm.DummyAgentDB

	// fetch the dummyagent
	_, err := db.First(&dummyagentDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	dummyagentDB.CopyBasicFieldsFromDummyAgent_WOP(&input.DummyAgent_WOP)
	dummyagentDB.DummyAgentPointersEncoding = input.DummyAgentPointersEncoding

	db, _ = db.Model(&dummyagentDB)
	_, err = db.Updates(&dummyagentDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	dummyagentNew := new(models.DummyAgent)
	dummyagentDB.CopyBasicFieldsToDummyAgent(dummyagentNew)

	// redeem pointers
	dummyagentDB.DecodePointers(backRepo, dummyagentNew)

	// get stage instance from DB instance, and call callback function
	dummyagentOld := backRepo.BackRepoDummyAgent.Map_DummyAgentDBID_DummyAgentPtr[dummyagentDB.ID]
	if dummyagentOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(dummyagentOld, dummyagentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the dummyagentDB
	writeJSON(w, http.StatusOK, dummyagentDB)
}
