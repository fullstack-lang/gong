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
var __BoxGeometry__dummysDeclaration__ models.BoxGeometry
var _ = __BoxGeometry__dummysDeclaration__
var __BoxGeometry_time__dummyDeclaration time.Duration
var _ = __BoxGeometry_time__dummyDeclaration

var mutexBoxGeometry sync.Mutex

// An BoxGeometryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateBoxGeometry
type BoxGeometryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BoxGeometryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateBoxGeometry
type BoxGeometryInput struct {
	// The BoxGeometry to submit or modify
	// in: body
	BoxGeometry *orm.BoxGeometryAPI
}

// UpdateBoxGeometry
//
// swagger:route PATCH /boxgeometrys/{ID} boxgeometrys updateBoxGeometry
//
// # Update a boxgeometry
//
// Responses:
// default: genericError
//
//	200: boxgeometryDBResponse
func (controller *Controller) UpdateBoxGeometry(w http.ResponseWriter, r *http.Request) {

	mutexBoxGeometry.Lock()
	defer mutexBoxGeometry.Unlock()

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
	db := backRepo.BackRepoBoxGeometry.GetDB()

	// Validate input
	var input orm.BoxGeometryAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var boxgeometryDB orm.BoxGeometryDB

	// fetch the boxgeometry
	_, err := db.First(&boxgeometryDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	boxgeometryDB.CopyBasicFieldsFromBoxGeometry_WOP(&input.BoxGeometry_WOP)
	boxgeometryDB.BoxGeometryPointersEncoding = input.BoxGeometryPointersEncoding

	db, _ = db.Model(&boxgeometryDB)
	_, err = db.Updates(&boxgeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	boxgeometryNew := new(models.BoxGeometry)
	boxgeometryDB.CopyBasicFieldsToBoxGeometry(boxgeometryNew)

	// redeem pointers
	boxgeometryDB.DecodePointers(backRepo, boxgeometryNew)

	// get stage instance from DB instance, and call callback function
	boxgeometryOld := backRepo.BackRepoBoxGeometry.Map_BoxGeometryDBID_BoxGeometryPtr[boxgeometryDB.ID]
	if boxgeometryOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(boxgeometryOld, boxgeometryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the boxgeometryDB
	writeJSON(w, http.StatusOK, boxgeometryDB)
}
