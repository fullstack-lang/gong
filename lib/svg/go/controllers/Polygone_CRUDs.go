// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
	"github.com/fullstack-lang/gong/lib/svg/go/orm"
)

// declaration in order to justify use of the models import
var __Polygone__dummysDeclaration__ models.Polygone
var _ = __Polygone__dummysDeclaration__
var __Polygone_time__dummyDeclaration time.Duration
var _ = __Polygone_time__dummyDeclaration

var mutexPolygone sync.Mutex

// An PolygoneID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updatePolygone
type PolygoneID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PolygoneInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updatePolygone
type PolygoneInput struct {
	// The Polygone to submit or modify
	// in: body
	Polygone *orm.PolygoneAPI
}

// UpdatePolygone
//
// swagger:route PATCH /polygones/{ID} polygones updatePolygone
//
// # Update a polygone
//
// Responses:
// default: genericError
//
//	200: polygoneDBResponse
func (controller *Controller) UpdatePolygone(w http.ResponseWriter, r *http.Request) {

	mutexPolygone.Lock()
	defer mutexPolygone.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoPolygone.GetDB()

	// Validate input
	var input orm.PolygoneAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var polygoneDB orm.PolygoneDB

	// fetch the polygone
	_, err := db.First(&polygoneDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	polygoneDB.CopyBasicFieldsFromPolygone_WOP(&input.Polygone_WOP)
	polygoneDB.PolygonePointersEncoding = input.PolygonePointersEncoding

	db, _ = db.Model(&polygoneDB)
	_, err = db.Updates(&polygoneDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	polygoneNew := new(models.Polygone)
	polygoneDB.CopyBasicFieldsToPolygone(polygoneNew)

	// redeem pointers
	polygoneDB.DecodePointers(backRepo, polygoneNew)

	// get stage instance from DB instance, and call callback function
	polygoneOld := backRepo.BackRepoPolygone.Map_PolygoneDBID_PolygonePtr[polygoneDB.ID]
	if polygoneOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(polygoneOld, polygoneNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the polygoneDB
	writeJSON(w, http.StatusOK, polygoneDB)
}
