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
var __CellInt__dummysDeclaration__ models.CellInt
var _ = __CellInt__dummysDeclaration__
var __CellInt_time__dummyDeclaration time.Duration
var _ = __CellInt_time__dummyDeclaration

var mutexCellInt sync.Mutex

// An CellIntID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateCellInt
type CellIntID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CellIntInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateCellInt
type CellIntInput struct {
	// The CellInt to submit or modify
	// in: body
	CellInt *orm.CellIntAPI
}

// UpdateCellInt
//
// swagger:route PATCH /cellints/{ID} cellints updateCellInt
//
// # Update a cellint
//
// Responses:
// default: genericError
//
//	200: cellintDBResponse
func (controller *Controller) UpdateCellInt(w http.ResponseWriter, r *http.Request) {

	mutexCellInt.Lock()
	defer mutexCellInt.Unlock()

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
	db := backRepo.BackRepoCellInt.GetDB()

	// Validate input
	var input orm.CellIntAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var cellintDB orm.CellIntDB

	// fetch the cellint
	_, err := db.First(&cellintDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	cellintDB.CopyBasicFieldsFromCellInt_WOP(&input.CellInt_WOP)
	cellintDB.CellIntPointersEncoding = input.CellIntPointersEncoding

	db, _ = db.Model(&cellintDB)
	_, err = db.Updates(&cellintDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	cellintNew := new(models.CellInt)
	cellintDB.CopyBasicFieldsToCellInt(cellintNew)

	// redeem pointers
	cellintDB.DecodePointers(backRepo, cellintNew)

	// get stage instance from DB instance, and call callback function
	cellintOld := backRepo.BackRepoCellInt.Map_CellIntDBID_CellIntPtr[cellintDB.ID]
	if cellintOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(cellintOld, cellintNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the cellintDB
	writeJSON(w, http.StatusOK, cellintDB)
}
