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
var __SphereGeometry__dummysDeclaration__ models.SphereGeometry
var _ = __SphereGeometry__dummysDeclaration__
var __SphereGeometry_time__dummyDeclaration time.Duration
var _ = __SphereGeometry_time__dummyDeclaration

var mutexSphereGeometry sync.Mutex

// An SphereGeometryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateSphereGeometry
type SphereGeometryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SphereGeometryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateSphereGeometry
type SphereGeometryInput struct {
	// The SphereGeometry to submit or modify
	// in: body
	SphereGeometry *orm.SphereGeometryAPI
}

// UpdateSphereGeometry
//
// swagger:route PATCH /spheregeometrys/{ID} spheregeometrys updateSphereGeometry
//
// # Update a spheregeometry
//
// Responses:
// default: genericError
//
//	200: spheregeometryDBResponse
func (controller *Controller) UpdateSphereGeometry(w http.ResponseWriter, r *http.Request) {

	mutexSphereGeometry.Lock()
	defer mutexSphereGeometry.Unlock()

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
	db := backRepo.BackRepoSphereGeometry.GetDB()

	// Validate input
	var input orm.SphereGeometryAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var spheregeometryDB orm.SphereGeometryDB

	// fetch the spheregeometry
	_, err := db.First(&spheregeometryDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	spheregeometryDB.CopyBasicFieldsFromSphereGeometry_WOP(&input.SphereGeometry_WOP)
	spheregeometryDB.SphereGeometryPointersEncoding = input.SphereGeometryPointersEncoding

	db, _ = db.Model(&spheregeometryDB)
	_, err = db.Updates(&spheregeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	spheregeometryNew := new(models.SphereGeometry)
	spheregeometryDB.CopyBasicFieldsToSphereGeometry(spheregeometryNew)

	// redeem pointers
	spheregeometryDB.DecodePointers(backRepo, spheregeometryNew)

	// get stage instance from DB instance, and call callback function
	spheregeometryOld := backRepo.BackRepoSphereGeometry.Map_SphereGeometryDBID_SphereGeometryPtr[spheregeometryDB.ID]
	if spheregeometryOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(spheregeometryOld, spheregeometryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the spheregeometryDB
	writeJSON(w, http.StatusOK, spheregeometryDB)
}
