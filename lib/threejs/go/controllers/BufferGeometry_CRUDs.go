// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/threejs/go/models"
	"github.com/fullstack-lang/gong/lib/threejs/go/orm"
)

// declaration in order to justify use of the models import
var __BufferGeometry__dummysDeclaration__ models.BufferGeometry
var _ = __BufferGeometry__dummysDeclaration__
var __BufferGeometry_time__dummyDeclaration time.Duration
var _ = __BufferGeometry_time__dummyDeclaration

var mutexBufferGeometry sync.Mutex

// An BufferGeometryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateBufferGeometry
type BufferGeometryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BufferGeometryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateBufferGeometry
type BufferGeometryInput struct {
	// The BufferGeometry to submit or modify
	// in: body
	BufferGeometry *orm.BufferGeometryAPI
}

// UpdateBufferGeometry
//
// swagger:route PATCH /buffergeometrys/{ID} buffergeometrys updateBufferGeometry
//
// # Update a buffergeometry
//
// Responses:
// default: genericError
//
//	200: buffergeometryDBResponse
func (controller *Controller) UpdateBufferGeometry(w http.ResponseWriter, r *http.Request) {

	mutexBufferGeometry.Lock()
	defer mutexBufferGeometry.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoBufferGeometry.GetDB()

	// Validate input
	var input orm.BufferGeometryAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var buffergeometryDB orm.BufferGeometryDB

	// fetch the buffergeometry
	_, err := db.First(&buffergeometryDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	buffergeometryDB.CopyBasicFieldsFromBufferGeometry_WOP(&input.BufferGeometry_WOP)
	buffergeometryDB.BufferGeometryPointersEncoding = input.BufferGeometryPointersEncoding

	db, _ = db.Model(&buffergeometryDB)
	_, err = db.Updates(&buffergeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	buffergeometryNew := new(models.BufferGeometry)
	buffergeometryDB.CopyBasicFieldsToBufferGeometry(buffergeometryNew)

	// redeem pointers
	buffergeometryDB.DecodePointers(backRepo, buffergeometryNew)

	// get stage instance from DB instance, and call callback function
	buffergeometryOld := backRepo.BackRepoBufferGeometry.Map_BufferGeometryDBID_BufferGeometryPtr[buffergeometryDB.ID]
	if buffergeometryOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(buffergeometryOld, buffergeometryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the buffergeometryDB
	writeJSON(w, http.StatusOK, buffergeometryDB)
}
