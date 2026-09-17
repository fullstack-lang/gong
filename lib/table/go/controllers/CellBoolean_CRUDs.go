// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/table/go/models"
	"github.com/fullstack-lang/gong/lib/table/go/orm"
)

// declaration in order to justify use of the models import
var __CellBoolean__dummysDeclaration__ models.CellBoolean
var _ = __CellBoolean__dummysDeclaration__
var __CellBoolean_time__dummyDeclaration time.Duration
var _ = __CellBoolean_time__dummyDeclaration

var mutexCellBoolean sync.Mutex

// An CellBooleanID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateCellBoolean
type CellBooleanID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CellBooleanInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateCellBoolean
type CellBooleanInput struct {
	// The CellBoolean to submit or modify
	// in: body
	CellBoolean *orm.CellBooleanAPI
}

// UpdateCellBoolean
//
// swagger:route PATCH /cellbooleans/{ID} cellbooleans updateCellBoolean
//
// # Update a cellboolean
//
// Responses:
// default: genericError
//
//	200: cellbooleanDBResponse
func (controller *Controller) UpdateCellBoolean(w http.ResponseWriter, r *http.Request) {

	mutexCellBoolean.Lock()
	defer mutexCellBoolean.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/table/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoCellBoolean.GetDB()

	// Validate input
	var input orm.CellBooleanAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var cellbooleanDB orm.CellBooleanDB

	// fetch the cellboolean
	_, err := db.First(&cellbooleanDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	cellbooleanDB.CopyBasicFieldsFromCellBoolean_WOP(&input.CellBoolean_WOP)
	cellbooleanDB.CellBooleanPointersEncoding = input.CellBooleanPointersEncoding

	db, _ = db.Model(&cellbooleanDB)
	_, err = db.Updates(&cellbooleanDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	cellbooleanNew := new(models.CellBoolean)
	cellbooleanDB.CopyBasicFieldsToCellBoolean(cellbooleanNew)

	// redeem pointers
	cellbooleanDB.DecodePointers(backRepo, cellbooleanNew)

	// get stage instance from DB instance, and call callback function
	cellbooleanOld := backRepo.BackRepoCellBoolean.Map_CellBooleanDBID_CellBooleanPtr[cellbooleanDB.ID]
	if cellbooleanOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(cellbooleanOld, cellbooleanNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the cellbooleanDB
	writeJSON(w, http.StatusOK, cellbooleanDB)
}
