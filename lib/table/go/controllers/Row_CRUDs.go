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
var __Row__dummysDeclaration__ models.Row
var _ = __Row__dummysDeclaration__
var __Row_time__dummyDeclaration time.Duration
var _ = __Row_time__dummyDeclaration

var mutexRow sync.Mutex

// An RowID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateRow
type RowID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RowInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateRow
type RowInput struct {
	// The Row to submit or modify
	// in: body
	Row *orm.RowAPI
}

// UpdateRow
//
// swagger:route PATCH /rows/{ID} rows updateRow
//
// # Update a row
//
// Responses:
// default: genericError
//
//	200: rowDBResponse
func (controller *Controller) UpdateRow(w http.ResponseWriter, r *http.Request) {

	mutexRow.Lock()
	defer mutexRow.Unlock()

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
	db := backRepo.BackRepoRow.GetDB()

	// Validate input
	var input orm.RowAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var rowDB orm.RowDB

	// fetch the row
	_, err := db.First(&rowDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	rowDB.CopyBasicFieldsFromRow_WOP(&input.Row_WOP)
	rowDB.RowPointersEncoding = input.RowPointersEncoding

	db, _ = db.Model(&rowDB)
	_, err = db.Updates(&rowDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	rowNew := new(models.Row)
	rowDB.CopyBasicFieldsToRow(rowNew)

	// redeem pointers
	rowDB.DecodePointers(backRepo, rowNew)

	// get stage instance from DB instance, and call callback function
	rowOld := backRepo.BackRepoRow.Map_RowDBID_RowPtr[rowDB.ID]
	if rowOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(rowOld, rowNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the rowDB
	writeJSON(w, http.StatusOK, rowDB)
}
