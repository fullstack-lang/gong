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
var __Line__dummysDeclaration__ models.Line
var _ = __Line__dummysDeclaration__
var __Line_time__dummyDeclaration time.Duration
var _ = __Line_time__dummyDeclaration

var mutexLine sync.Mutex

// An LineID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateLine
type LineID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LineInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateLine
type LineInput struct {
	// The Line to submit or modify
	// in: body
	Line *orm.LineAPI
}

// UpdateLine
//
// swagger:route PATCH /lines/{ID} lines updateLine
//
// # Update a line
//
// Responses:
// default: genericError
//
//	200: lineDBResponse
func (controller *Controller) UpdateLine(w http.ResponseWriter, r *http.Request) {

	mutexLine.Lock()
	defer mutexLine.Unlock()

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
	db := backRepo.BackRepoLine.GetDB()

	// Validate input
	var input orm.LineAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var lineDB orm.LineDB

	// fetch the line
	_, err := db.First(&lineDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	lineDB.CopyBasicFieldsFromLine_WOP(&input.Line_WOP)
	lineDB.LinePointersEncoding = input.LinePointersEncoding

	db, _ = db.Model(&lineDB)
	_, err = db.Updates(&lineDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	lineNew := new(models.Line)
	lineDB.CopyBasicFieldsToLine(lineNew)

	// redeem pointers
	lineDB.DecodePointers(backRepo, lineNew)

	// get stage instance from DB instance, and call callback function
	lineOld := backRepo.BackRepoLine.Map_LineDBID_LinePtr[lineDB.ID]
	if lineOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(lineOld, lineNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the lineDB
	writeJSON(w, http.StatusOK, lineDB)
}
