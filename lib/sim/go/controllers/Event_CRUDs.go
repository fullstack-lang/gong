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
var __Event__dummysDeclaration__ models.Event
var _ = __Event__dummysDeclaration__
var __Event_time__dummyDeclaration time.Duration
var _ = __Event_time__dummyDeclaration

var mutexEvent sync.Mutex

// An EventID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateEvent
type EventID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// EventInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateEvent
type EventInput struct {
	// The Event to submit or modify
	// in: body
	Event *orm.EventAPI
}

// UpdateEvent
//
// swagger:route PATCH /events/{ID} events updateEvent
//
// # Update a event
//
// Responses:
// default: genericError
//
//	200: eventDBResponse
func (controller *Controller) UpdateEvent(w http.ResponseWriter, r *http.Request) {

	mutexEvent.Lock()
	defer mutexEvent.Unlock()

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
	db := backRepo.BackRepoEvent.GetDB()

	// Validate input
	var input orm.EventAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var eventDB orm.EventDB

	// fetch the event
	_, err := db.First(&eventDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	eventDB.CopyBasicFieldsFromEvent_WOP(&input.Event_WOP)
	eventDB.EventPointersEncoding = input.EventPointersEncoding

	db, _ = db.Model(&eventDB)
	_, err = db.Updates(&eventDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	eventNew := new(models.Event)
	eventDB.CopyBasicFieldsToEvent(eventNew)

	// redeem pointers
	eventDB.DecodePointers(backRepo, eventNew)

	// get stage instance from DB instance, and call callback function
	eventOld := backRepo.BackRepoEvent.Map_EventDBID_EventPtr[eventDB.ID]
	if eventOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(eventOld, eventNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the eventDB
	writeJSON(w, http.StatusOK, eventDB)
}
