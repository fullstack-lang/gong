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
var __TorusGeometry__dummysDeclaration__ models.TorusGeometry
var _ = __TorusGeometry__dummysDeclaration__
var __TorusGeometry_time__dummyDeclaration time.Duration
var _ = __TorusGeometry_time__dummyDeclaration

var mutexTorusGeometry sync.Mutex

// An TorusGeometryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateTorusGeometry
type TorusGeometryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TorusGeometryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateTorusGeometry
type TorusGeometryInput struct {
	// The TorusGeometry to submit or modify
	// in: body
	TorusGeometry *orm.TorusGeometryAPI
}

// UpdateTorusGeometry
//
// swagger:route PATCH /torusgeometrys/{ID} torusgeometrys updateTorusGeometry
//
// # Update a torusgeometry
//
// Responses:
// default: genericError
//
//	200: torusgeometryDBResponse
func (controller *Controller) UpdateTorusGeometry(w http.ResponseWriter, r *http.Request) {

	mutexTorusGeometry.Lock()
	defer mutexTorusGeometry.Unlock()

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
	db := backRepo.BackRepoTorusGeometry.GetDB()

	// Validate input
	var input orm.TorusGeometryAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var torusgeometryDB orm.TorusGeometryDB

	// fetch the torusgeometry
	_, err := db.First(&torusgeometryDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	torusgeometryDB.CopyBasicFieldsFromTorusGeometry_WOP(&input.TorusGeometry_WOP)
	torusgeometryDB.TorusGeometryPointersEncoding = input.TorusGeometryPointersEncoding

	db, _ = db.Model(&torusgeometryDB)
	_, err = db.Updates(&torusgeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	torusgeometryNew := new(models.TorusGeometry)
	torusgeometryDB.CopyBasicFieldsToTorusGeometry(torusgeometryNew)

	// redeem pointers
	torusgeometryDB.DecodePointers(backRepo, torusgeometryNew)

	// get stage instance from DB instance, and call callback function
	torusgeometryOld := backRepo.BackRepoTorusGeometry.Map_TorusGeometryDBID_TorusGeometryPtr[torusgeometryDB.ID]
	if torusgeometryOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(torusgeometryOld, torusgeometryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the torusgeometryDB
	writeJSON(w, http.StatusOK, torusgeometryDB)
}
