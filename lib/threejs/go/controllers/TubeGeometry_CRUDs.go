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
var __TubeGeometry__dummysDeclaration__ models.TubeGeometry
var _ = __TubeGeometry__dummysDeclaration__
var __TubeGeometry_time__dummyDeclaration time.Duration
var _ = __TubeGeometry_time__dummyDeclaration

var mutexTubeGeometry sync.Mutex

// An TubeGeometryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateTubeGeometry
type TubeGeometryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TubeGeometryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateTubeGeometry
type TubeGeometryInput struct {
	// The TubeGeometry to submit or modify
	// in: body
	TubeGeometry *orm.TubeGeometryAPI
}

// UpdateTubeGeometry
//
// swagger:route PATCH /tubegeometrys/{ID} tubegeometrys updateTubeGeometry
//
// # Update a tubegeometry
//
// Responses:
// default: genericError
//
//	200: tubegeometryDBResponse
func (controller *Controller) UpdateTubeGeometry(w http.ResponseWriter, r *http.Request) {

	mutexTubeGeometry.Lock()
	defer mutexTubeGeometry.Unlock()

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
	db := backRepo.BackRepoTubeGeometry.GetDB()

	// Validate input
	var input orm.TubeGeometryAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var tubegeometryDB orm.TubeGeometryDB

	// fetch the tubegeometry
	_, err := db.First(&tubegeometryDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	tubegeometryDB.CopyBasicFieldsFromTubeGeometry_WOP(&input.TubeGeometry_WOP)
	tubegeometryDB.TubeGeometryPointersEncoding = input.TubeGeometryPointersEncoding

	db, _ = db.Model(&tubegeometryDB)
	_, err = db.Updates(&tubegeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	tubegeometryNew := new(models.TubeGeometry)
	tubegeometryDB.CopyBasicFieldsToTubeGeometry(tubegeometryNew)

	// redeem pointers
	tubegeometryDB.DecodePointers(backRepo, tubegeometryNew)

	// get stage instance from DB instance, and call callback function
	tubegeometryOld := backRepo.BackRepoTubeGeometry.Map_TubeGeometryDBID_TubeGeometryPtr[tubegeometryDB.ID]
	if tubegeometryOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(tubegeometryOld, tubegeometryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the tubegeometryDB
	writeJSON(w, http.StatusOK, tubegeometryDB)
}
