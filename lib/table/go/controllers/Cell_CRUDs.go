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
var __Cell__dummysDeclaration__ models.Cell
var _ = __Cell__dummysDeclaration__
var __Cell_time__dummyDeclaration time.Duration
var _ = __Cell_time__dummyDeclaration

var mutexCell sync.Mutex

// An CellID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateCell
type CellID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CellInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateCell
type CellInput struct {
	// The Cell to submit or modify
	// in: body
	Cell *orm.CellAPI
}

// UpdateCell
//
// swagger:route PATCH /cells/{ID} cells updateCell
//
// # Update a cell
//
// Responses:
// default: genericError
//
//	200: cellDBResponse
func (controller *Controller) UpdateCell(w http.ResponseWriter, r *http.Request) {

	mutexCell.Lock()
	defer mutexCell.Unlock()

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
	db := backRepo.BackRepoCell.GetDB()

	// Validate input
	var input orm.CellAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var cellDB orm.CellDB

	// fetch the cell
	_, err := db.First(&cellDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	cellDB.CopyBasicFieldsFromCell_WOP(&input.Cell_WOP)
	cellDB.CellPointersEncoding = input.CellPointersEncoding

	db, _ = db.Model(&cellDB)
	_, err = db.Updates(&cellDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	cellNew := new(models.Cell)
	cellDB.CopyBasicFieldsToCell(cellNew)

	// redeem pointers
	cellDB.DecodePointers(backRepo, cellNew)

	// get stage instance from DB instance, and call callback function
	cellOld := backRepo.BackRepoCell.Map_CellDBID_CellPtr[cellDB.ID]
	if cellOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(cellOld, cellNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the cellDB
	writeJSON(w, http.StatusOK, cellDB)
}
