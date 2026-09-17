// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/split/go/models"
	"github.com/fullstack-lang/gong/lib/split/go/orm"
)

// declaration in order to justify use of the models import
var __Cursor__dummysDeclaration__ models.Cursor
var _ = __Cursor__dummysDeclaration__
var __Cursor_time__dummyDeclaration time.Duration
var _ = __Cursor_time__dummyDeclaration

var mutexCursor sync.Mutex

// An CursorID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateCursor
type CursorID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CursorInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateCursor
type CursorInput struct {
	// The Cursor to submit or modify
	// in: body
	Cursor *orm.CursorAPI
}

// UpdateCursor
//
// swagger:route PATCH /cursors/{ID} cursors updateCursor
//
// # Update a cursor
//
// Responses:
// default: genericError
//
//	200: cursorDBResponse
func (controller *Controller) UpdateCursor(w http.ResponseWriter, r *http.Request) {

	mutexCursor.Lock()
	defer mutexCursor.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoCursor.GetDB()

	// Validate input
	var input orm.CursorAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var cursorDB orm.CursorDB

	// fetch the cursor
	_, err := db.First(&cursorDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	cursorDB.CopyBasicFieldsFromCursor_WOP(&input.Cursor_WOP)
	cursorDB.CursorPointersEncoding = input.CursorPointersEncoding

	db, _ = db.Model(&cursorDB)
	_, err = db.Updates(&cursorDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	cursorNew := new(models.Cursor)
	cursorDB.CopyBasicFieldsToCursor(cursorNew)

	// redeem pointers
	cursorDB.DecodePointers(backRepo, cursorNew)

	// get stage instance from DB instance, and call callback function
	cursorOld := backRepo.BackRepoCursor.Map_CursorDBID_CursorPtr[cursorDB.ID]
	if cursorOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(cursorOld, cursorNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the cursorDB
	writeJSON(w, http.StatusOK, cursorDB)
}
