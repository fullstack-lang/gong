// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/tone/go/models"
	"github.com/fullstack-lang/gong/lib/tone/go/orm"
)

// declaration in order to justify use of the models import
var __Note__dummysDeclaration__ models.Note
var _ = __Note__dummysDeclaration__
var __Note_time__dummyDeclaration time.Duration
var _ = __Note_time__dummyDeclaration

var mutexNote sync.Mutex

// An NoteID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateNote
type NoteID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// NoteInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateNote
type NoteInput struct {
	// The Note to submit or modify
	// in: body
	Note *orm.NoteAPI
}

// UpdateNote
//
// swagger:route PATCH /notes/{ID} notes updateNote
//
// # Update a note
//
// Responses:
// default: genericError
//
//	200: noteDBResponse
func (controller *Controller) UpdateNote(w http.ResponseWriter, r *http.Request) {

	mutexNote.Lock()
	defer mutexNote.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/tone/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoNote.GetDB()

	// Validate input
	var input orm.NoteAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var noteDB orm.NoteDB

	// fetch the note
	_, err := db.First(&noteDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	noteDB.CopyBasicFieldsFromNote_WOP(&input.Note_WOP)
	noteDB.NotePointersEncoding = input.NotePointersEncoding

	db, _ = db.Model(&noteDB)
	_, err = db.Updates(&noteDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	noteNew := new(models.Note)
	noteDB.CopyBasicFieldsToNote(noteNew)

	// redeem pointers
	noteDB.DecodePointers(backRepo, noteNew)

	// get stage instance from DB instance, and call callback function
	noteOld := backRepo.BackRepoNote.Map_NoteDBID_NotePtr[noteDB.ID]
	if noteOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(noteOld, noteNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the noteDB
	writeJSON(w, http.StatusOK, noteDB)
}
