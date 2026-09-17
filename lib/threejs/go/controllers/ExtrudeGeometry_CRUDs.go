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
var __ExtrudeGeometry__dummysDeclaration__ models.ExtrudeGeometry
var _ = __ExtrudeGeometry__dummysDeclaration__
var __ExtrudeGeometry_time__dummyDeclaration time.Duration
var _ = __ExtrudeGeometry_time__dummyDeclaration

var mutexExtrudeGeometry sync.Mutex

// An ExtrudeGeometryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateExtrudeGeometry
type ExtrudeGeometryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ExtrudeGeometryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateExtrudeGeometry
type ExtrudeGeometryInput struct {
	// The ExtrudeGeometry to submit or modify
	// in: body
	ExtrudeGeometry *orm.ExtrudeGeometryAPI
}

// UpdateExtrudeGeometry
//
// swagger:route PATCH /extrudegeometrys/{ID} extrudegeometrys updateExtrudeGeometry
//
// # Update a extrudegeometry
//
// Responses:
// default: genericError
//
//	200: extrudegeometryDBResponse
func (controller *Controller) UpdateExtrudeGeometry(w http.ResponseWriter, r *http.Request) {

	mutexExtrudeGeometry.Lock()
	defer mutexExtrudeGeometry.Unlock()

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
	db := backRepo.BackRepoExtrudeGeometry.GetDB()

	// Validate input
	var input orm.ExtrudeGeometryAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var extrudegeometryDB orm.ExtrudeGeometryDB

	// fetch the extrudegeometry
	_, err := db.First(&extrudegeometryDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	extrudegeometryDB.CopyBasicFieldsFromExtrudeGeometry_WOP(&input.ExtrudeGeometry_WOP)
	extrudegeometryDB.ExtrudeGeometryPointersEncoding = input.ExtrudeGeometryPointersEncoding

	db, _ = db.Model(&extrudegeometryDB)
	_, err = db.Updates(&extrudegeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	extrudegeometryNew := new(models.ExtrudeGeometry)
	extrudegeometryDB.CopyBasicFieldsToExtrudeGeometry(extrudegeometryNew)

	// redeem pointers
	extrudegeometryDB.DecodePointers(backRepo, extrudegeometryNew)

	// get stage instance from DB instance, and call callback function
	extrudegeometryOld := backRepo.BackRepoExtrudeGeometry.Map_ExtrudeGeometryDBID_ExtrudeGeometryPtr[extrudegeometryDB.ID]
	if extrudegeometryOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(extrudegeometryOld, extrudegeometryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the extrudegeometryDB
	writeJSON(w, http.StatusOK, extrudegeometryDB)
}
