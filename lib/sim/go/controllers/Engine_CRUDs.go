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
var __Engine__dummysDeclaration__ models.Engine
var _ = __Engine__dummysDeclaration__
var __Engine_time__dummyDeclaration time.Duration
var _ = __Engine_time__dummyDeclaration

var mutexEngine sync.Mutex

// An EngineID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateEngine
type EngineID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// EngineInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateEngine
type EngineInput struct {
	// The Engine to submit or modify
	// in: body
	Engine *orm.EngineAPI
}

// UpdateEngine
//
// swagger:route PATCH /engines/{ID} engines updateEngine
//
// # Update a engine
//
// Responses:
// default: genericError
//
//	200: engineDBResponse
func (controller *Controller) UpdateEngine(w http.ResponseWriter, r *http.Request) {

	mutexEngine.Lock()
	defer mutexEngine.Unlock()

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
	db := backRepo.BackRepoEngine.GetDB()

	// Validate input
	var input orm.EngineAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var engineDB orm.EngineDB

	// fetch the engine
	_, err := db.First(&engineDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	engineDB.CopyBasicFieldsFromEngine_WOP(&input.Engine_WOP)
	engineDB.EnginePointersEncoding = input.EnginePointersEncoding

	db, _ = db.Model(&engineDB)
	_, err = db.Updates(&engineDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	engineNew := new(models.Engine)
	engineDB.CopyBasicFieldsToEngine(engineNew)

	// redeem pointers
	engineDB.DecodePointers(backRepo, engineNew)

	// get stage instance from DB instance, and call callback function
	engineOld := backRepo.BackRepoEngine.Map_EngineDBID_EnginePtr[engineDB.ID]
	if engineOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(engineOld, engineNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the engineDB
	writeJSON(w, http.StatusOK, engineDB)
}
