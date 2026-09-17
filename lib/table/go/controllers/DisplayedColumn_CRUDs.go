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
var __DisplayedColumn__dummysDeclaration__ models.DisplayedColumn
var _ = __DisplayedColumn__dummysDeclaration__
var __DisplayedColumn_time__dummyDeclaration time.Duration
var _ = __DisplayedColumn_time__dummyDeclaration

var mutexDisplayedColumn sync.Mutex

// An DisplayedColumnID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateDisplayedColumn
type DisplayedColumnID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DisplayedColumnInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateDisplayedColumn
type DisplayedColumnInput struct {
	// The DisplayedColumn to submit or modify
	// in: body
	DisplayedColumn *orm.DisplayedColumnAPI
}

// UpdateDisplayedColumn
//
// swagger:route PATCH /displayedcolumns/{ID} displayedcolumns updateDisplayedColumn
//
// # Update a displayedcolumn
//
// Responses:
// default: genericError
//
//	200: displayedcolumnDBResponse
func (controller *Controller) UpdateDisplayedColumn(w http.ResponseWriter, r *http.Request) {

	mutexDisplayedColumn.Lock()
	defer mutexDisplayedColumn.Unlock()

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
	db := backRepo.BackRepoDisplayedColumn.GetDB()

	// Validate input
	var input orm.DisplayedColumnAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var displayedcolumnDB orm.DisplayedColumnDB

	// fetch the displayedcolumn
	_, err := db.First(&displayedcolumnDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	displayedcolumnDB.CopyBasicFieldsFromDisplayedColumn_WOP(&input.DisplayedColumn_WOP)
	displayedcolumnDB.DisplayedColumnPointersEncoding = input.DisplayedColumnPointersEncoding

	db, _ = db.Model(&displayedcolumnDB)
	_, err = db.Updates(&displayedcolumnDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	displayedcolumnNew := new(models.DisplayedColumn)
	displayedcolumnDB.CopyBasicFieldsToDisplayedColumn(displayedcolumnNew)

	// redeem pointers
	displayedcolumnDB.DecodePointers(backRepo, displayedcolumnNew)

	// get stage instance from DB instance, and call callback function
	displayedcolumnOld := backRepo.BackRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnPtr[displayedcolumnDB.ID]
	if displayedcolumnOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(displayedcolumnOld, displayedcolumnNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the displayedcolumnDB
	writeJSON(w, http.StatusOK, displayedcolumnDB)
}
