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
var __Command__dummysDeclaration__ models.Command
var _ = __Command__dummysDeclaration__
var __Command_time__dummyDeclaration time.Duration
var _ = __Command_time__dummyDeclaration

var mutexCommand sync.Mutex

// An CommandID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateCommand
type CommandID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CommandInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateCommand
type CommandInput struct {
	// The Command to submit or modify
	// in: body
	Command *orm.CommandAPI
}

// UpdateCommand
//
// swagger:route PATCH /commands/{ID} commands updateCommand
//
// # Update a command
//
// Responses:
// default: genericError
//
//	200: commandDBResponse
func (controller *Controller) UpdateCommand(w http.ResponseWriter, r *http.Request) {

	mutexCommand.Lock()
	defer mutexCommand.Unlock()

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
	db := backRepo.BackRepoCommand.GetDB()

	// Validate input
	var input orm.CommandAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var commandDB orm.CommandDB

	// fetch the command
	_, err := db.First(&commandDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	commandDB.CopyBasicFieldsFromCommand_WOP(&input.Command_WOP)
	commandDB.CommandPointersEncoding = input.CommandPointersEncoding

	db, _ = db.Model(&commandDB)
	_, err = db.Updates(&commandDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	commandNew := new(models.Command)
	commandDB.CopyBasicFieldsToCommand(commandNew)

	// redeem pointers
	commandDB.DecodePointers(backRepo, commandNew)

	// get stage instance from DB instance, and call callback function
	commandOld := backRepo.BackRepoCommand.Map_CommandDBID_CommandPtr[commandDB.ID]
	if commandOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(commandOld, commandNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the commandDB
	writeJSON(w, http.StatusOK, commandDB)
}
