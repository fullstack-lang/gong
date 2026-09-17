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
var __Table__dummysDeclaration__ models.Table
var _ = __Table__dummysDeclaration__
var __Table_time__dummyDeclaration time.Duration
var _ = __Table_time__dummyDeclaration

var mutexTable sync.Mutex

// An TableID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateTable
type TableID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TableInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateTable
type TableInput struct {
	// The Table to submit or modify
	// in: body
	Table *orm.TableAPI
}

// UpdateTable
//
// swagger:route PATCH /tables/{ID} tables updateTable
//
// # Update a table
//
// Responses:
// default: genericError
//
//	200: tableDBResponse
func (controller *Controller) UpdateTable(w http.ResponseWriter, r *http.Request) {

	mutexTable.Lock()
	defer mutexTable.Unlock()

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
	db := backRepo.BackRepoTable.GetDB()

	// Validate input
	var input orm.TableAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var tableDB orm.TableDB

	// fetch the table
	_, err := db.First(&tableDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	tableDB.CopyBasicFieldsFromTable_WOP(&input.Table_WOP)
	tableDB.TablePointersEncoding = input.TablePointersEncoding

	db, _ = db.Model(&tableDB)
	_, err = db.Updates(&tableDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	tableNew := new(models.Table)
	tableDB.CopyBasicFieldsToTable(tableNew)

	// redeem pointers
	tableDB.DecodePointers(backRepo, tableNew)

	// get stage instance from DB instance, and call callback function
	tableOld := backRepo.BackRepoTable.Map_TableDBID_TablePtr[tableDB.ID]
	if tableOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(tableOld, tableNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the tableDB
	writeJSON(w, http.StatusOK, tableDB)
}
