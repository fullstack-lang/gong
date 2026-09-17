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
var __CellString__dummysDeclaration__ models.CellString
var _ = __CellString__dummysDeclaration__
var __CellString_time__dummyDeclaration time.Duration
var _ = __CellString_time__dummyDeclaration

var mutexCellString sync.Mutex

// An CellStringID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateCellString
type CellStringID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CellStringInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateCellString
type CellStringInput struct {
	// The CellString to submit or modify
	// in: body
	CellString *orm.CellStringAPI
}

// UpdateCellString
//
// swagger:route PATCH /cellstrings/{ID} cellstrings updateCellString
//
// # Update a cellstring
//
// Responses:
// default: genericError
//
//	200: cellstringDBResponse
func (controller *Controller) UpdateCellString(w http.ResponseWriter, r *http.Request) {

	mutexCellString.Lock()
	defer mutexCellString.Unlock()

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
	db := backRepo.BackRepoCellString.GetDB()

	// Validate input
	var input orm.CellStringAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var cellstringDB orm.CellStringDB

	// fetch the cellstring
	_, err := db.First(&cellstringDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	cellstringDB.CopyBasicFieldsFromCellString_WOP(&input.CellString_WOP)
	cellstringDB.CellStringPointersEncoding = input.CellStringPointersEncoding

	db, _ = db.Model(&cellstringDB)
	_, err = db.Updates(&cellstringDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	cellstringNew := new(models.CellString)
	cellstringDB.CopyBasicFieldsToCellString(cellstringNew)

	// redeem pointers
	cellstringDB.DecodePointers(backRepo, cellstringNew)

	// get stage instance from DB instance, and call callback function
	cellstringOld := backRepo.BackRepoCellString.Map_CellStringDBID_CellStringPtr[cellstringDB.ID]
	if cellstringOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(cellstringOld, cellstringNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the cellstringDB
	writeJSON(w, http.StatusOK, cellstringDB)
}
