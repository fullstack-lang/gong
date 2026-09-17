// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/load/go/models"
	"github.com/fullstack-lang/gong/lib/load/go/orm"
)

// declaration in order to justify use of the models import
var __Message__dummysDeclaration__ models.Message
var _ = __Message__dummysDeclaration__
var __Message_time__dummyDeclaration time.Duration
var _ = __Message_time__dummyDeclaration

var mutexMessage sync.Mutex

// An MessageID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateMessage
type MessageID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MessageInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateMessage
type MessageInput struct {
	// The Message to submit or modify
	// in: body
	Message *orm.MessageAPI
}

// UpdateMessage
//
// swagger:route PATCH /messages/{ID} messages updateMessage
//
// # Update a message
//
// Responses:
// default: genericError
//
//	200: messageDBResponse
func (controller *Controller) UpdateMessage(w http.ResponseWriter, r *http.Request) {

	mutexMessage.Lock()
	defer mutexMessage.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/load/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMessage.GetDB()

	// Validate input
	var input orm.MessageAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var messageDB orm.MessageDB

	// fetch the message
	_, err := db.First(&messageDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	messageDB.CopyBasicFieldsFromMessage_WOP(&input.Message_WOP)
	messageDB.MessagePointersEncoding = input.MessagePointersEncoding

	db, _ = db.Model(&messageDB)
	_, err = db.Updates(&messageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	messageNew := new(models.Message)
	messageDB.CopyBasicFieldsToMessage(messageNew)

	// redeem pointers
	messageDB.DecodePointers(backRepo, messageNew)

	// get stage instance from DB instance, and call callback function
	messageOld := backRepo.BackRepoMessage.Map_MessageDBID_MessagePtr[messageDB.ID]
	if messageOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(messageOld, messageNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the messageDB
	writeJSON(w, http.StatusOK, messageDB)
}
