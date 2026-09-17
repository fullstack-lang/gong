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
var __CellIcon__dummysDeclaration__ models.CellIcon
var _ = __CellIcon__dummysDeclaration__
var __CellIcon_time__dummyDeclaration time.Duration
var _ = __CellIcon_time__dummyDeclaration

var mutexCellIcon sync.Mutex

// An CellIconID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateCellIcon
type CellIconID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CellIconInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateCellIcon
type CellIconInput struct {
	// The CellIcon to submit or modify
	// in: body
	CellIcon *orm.CellIconAPI
}

// UpdateCellIcon
//
// swagger:route PATCH /cellicons/{ID} cellicons updateCellIcon
//
// # Update a cellicon
//
// Responses:
// default: genericError
//
//	200: celliconDBResponse
func (controller *Controller) UpdateCellIcon(w http.ResponseWriter, r *http.Request) {

	mutexCellIcon.Lock()
	defer mutexCellIcon.Unlock()

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
	db := backRepo.BackRepoCellIcon.GetDB()

	// Validate input
	var input orm.CellIconAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var celliconDB orm.CellIconDB

	// fetch the cellicon
	_, err := db.First(&celliconDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	celliconDB.CopyBasicFieldsFromCellIcon_WOP(&input.CellIcon_WOP)
	celliconDB.CellIconPointersEncoding = input.CellIconPointersEncoding

	db, _ = db.Model(&celliconDB)
	_, err = db.Updates(&celliconDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	celliconNew := new(models.CellIcon)
	celliconDB.CopyBasicFieldsToCellIcon(celliconNew)

	// redeem pointers
	celliconDB.DecodePointers(backRepo, celliconNew)

	// get stage instance from DB instance, and call callback function
	celliconOld := backRepo.BackRepoCellIcon.Map_CellIconDBID_CellIconPtr[celliconDB.ID]
	if celliconOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(celliconOld, celliconNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the celliconDB
	writeJSON(w, http.StatusOK, celliconDB)
}
