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
var __CellFloat64__dummysDeclaration__ models.CellFloat64
var _ = __CellFloat64__dummysDeclaration__
var __CellFloat64_time__dummyDeclaration time.Duration
var _ = __CellFloat64_time__dummyDeclaration

var mutexCellFloat64 sync.Mutex

// An CellFloat64ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateCellFloat64
type CellFloat64ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CellFloat64Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateCellFloat64
type CellFloat64Input struct {
	// The CellFloat64 to submit or modify
	// in: body
	CellFloat64 *orm.CellFloat64API
}

// UpdateCellFloat64
//
// swagger:route PATCH /cellfloat64s/{ID} cellfloat64s updateCellFloat64
//
// # Update a cellfloat64
//
// Responses:
// default: genericError
//
//	200: cellfloat64DBResponse
func (controller *Controller) UpdateCellFloat64(w http.ResponseWriter, r *http.Request) {

	mutexCellFloat64.Lock()
	defer mutexCellFloat64.Unlock()

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
	db := backRepo.BackRepoCellFloat64.GetDB()

	// Validate input
	var input orm.CellFloat64API
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var cellfloat64DB orm.CellFloat64DB

	// fetch the cellfloat64
	_, err := db.First(&cellfloat64DB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	cellfloat64DB.CopyBasicFieldsFromCellFloat64_WOP(&input.CellFloat64_WOP)
	cellfloat64DB.CellFloat64PointersEncoding = input.CellFloat64PointersEncoding

	db, _ = db.Model(&cellfloat64DB)
	_, err = db.Updates(&cellfloat64DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	cellfloat64New := new(models.CellFloat64)
	cellfloat64DB.CopyBasicFieldsToCellFloat64(cellfloat64New)

	// redeem pointers
	cellfloat64DB.DecodePointers(backRepo, cellfloat64New)

	// get stage instance from DB instance, and call callback function
	cellfloat64Old := backRepo.BackRepoCellFloat64.Map_CellFloat64DBID_CellFloat64Ptr[cellfloat64DB.ID]
	if cellfloat64Old != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(cellfloat64Old, cellfloat64New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the cellfloat64DB
	writeJSON(w, http.StatusOK, cellfloat64DB)
}
