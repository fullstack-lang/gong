// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/test/test1/go/models"
	"github.com/fullstack-lang/gong/test/test1/go/orm"
)

// declaration in order to justify use of the models import
var __Bstruct__dummysDeclaration__ models.Bstruct
var _ = __Bstruct__dummysDeclaration__
var __Bstruct_time__dummyDeclaration time.Duration
var _ = __Bstruct_time__dummyDeclaration

var mutexBstruct sync.Mutex

// An BstructID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateBstruct
type BstructID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BstructInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateBstruct
type BstructInput struct {
	// The Bstruct to submit or modify
	// in: body
	Bstruct *orm.BstructAPI
}

// UpdateBstruct
//
// swagger:route PATCH /bstructs/{ID} bstructs updateBstruct
//
// # Update a bstruct
//
// Responses:
// default: genericError
//
//	200: bstructDBResponse
func (controller *Controller) UpdateBstruct(w http.ResponseWriter, r *http.Request) {

	mutexBstruct.Lock()
	defer mutexBstruct.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/test/test1/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoBstruct.GetDB()

	// Validate input
	var input orm.BstructAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var bstructDB orm.BstructDB

	// fetch the bstruct
	_, err := db.First(&bstructDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	bstructDB.CopyBasicFieldsFromBstruct_WOP(&input.Bstruct_WOP)
	bstructDB.BstructPointersEncoding = input.BstructPointersEncoding

	db, _ = db.Model(&bstructDB)
	_, err = db.Updates(&bstructDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	bstructNew := new(models.Bstruct)
	bstructDB.CopyBasicFieldsToBstruct(bstructNew)

	// redeem pointers
	bstructDB.DecodePointers(backRepo, bstructNew)

	// get stage instance from DB instance, and call callback function
	bstructOld := backRepo.BackRepoBstruct.Map_BstructDBID_BstructPtr[bstructDB.ID]
	if bstructOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(bstructOld, bstructNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the bstructDB
	writeJSON(w, http.StatusOK, bstructDB)
}
