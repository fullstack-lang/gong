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
var __PlaneGeometry__dummysDeclaration__ models.PlaneGeometry
var _ = __PlaneGeometry__dummysDeclaration__
var __PlaneGeometry_time__dummyDeclaration time.Duration
var _ = __PlaneGeometry_time__dummyDeclaration

var mutexPlaneGeometry sync.Mutex

// An PlaneGeometryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updatePlaneGeometry
type PlaneGeometryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PlaneGeometryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updatePlaneGeometry
type PlaneGeometryInput struct {
	// The PlaneGeometry to submit or modify
	// in: body
	PlaneGeometry *orm.PlaneGeometryAPI
}

// UpdatePlaneGeometry
//
// swagger:route PATCH /planegeometrys/{ID} planegeometrys updatePlaneGeometry
//
// # Update a planegeometry
//
// Responses:
// default: genericError
//
//	200: planegeometryDBResponse
func (controller *Controller) UpdatePlaneGeometry(w http.ResponseWriter, r *http.Request) {

	mutexPlaneGeometry.Lock()
	defer mutexPlaneGeometry.Unlock()

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
	db := backRepo.BackRepoPlaneGeometry.GetDB()

	// Validate input
	var input orm.PlaneGeometryAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var planegeometryDB orm.PlaneGeometryDB

	// fetch the planegeometry
	_, err := db.First(&planegeometryDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	planegeometryDB.CopyBasicFieldsFromPlaneGeometry_WOP(&input.PlaneGeometry_WOP)
	planegeometryDB.PlaneGeometryPointersEncoding = input.PlaneGeometryPointersEncoding

	db, _ = db.Model(&planegeometryDB)
	_, err = db.Updates(&planegeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	planegeometryNew := new(models.PlaneGeometry)
	planegeometryDB.CopyBasicFieldsToPlaneGeometry(planegeometryNew)

	// redeem pointers
	planegeometryDB.DecodePointers(backRepo, planegeometryNew)

	// get stage instance from DB instance, and call callback function
	planegeometryOld := backRepo.BackRepoPlaneGeometry.Map_PlaneGeometryDBID_PlaneGeometryPtr[planegeometryDB.ID]
	if planegeometryOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(planegeometryOld, planegeometryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the planegeometryDB
	writeJSON(w, http.StatusOK, planegeometryDB)
}
