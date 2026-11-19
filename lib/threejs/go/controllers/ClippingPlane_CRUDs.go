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
var __ClippingPlane__dummysDeclaration__ models.ClippingPlane
var _ = __ClippingPlane__dummysDeclaration__
var __ClippingPlane_time__dummyDeclaration time.Duration
var _ = __ClippingPlane_time__dummyDeclaration

var mutexClippingPlane sync.Mutex

// An ClippingPlaneID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateClippingPlane
type ClippingPlaneID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ClippingPlaneInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateClippingPlane
type ClippingPlaneInput struct {
	// The ClippingPlane to submit or modify
	// in: body
	ClippingPlane *orm.ClippingPlaneAPI
}

// UpdateClippingPlane
//
// swagger:route PATCH /clippingplanes/{ID} clippingplanes updateClippingPlane
//
// # Update a clippingplane
//
// Responses:
// default: genericError
//
//	200: clippingplaneDBResponse
func (controller *Controller) UpdateClippingPlane(w http.ResponseWriter, r *http.Request) {

	mutexClippingPlane.Lock()
	defer mutexClippingPlane.Unlock()

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
	db := backRepo.BackRepoClippingPlane.GetDB()

	// Validate input
	var input orm.ClippingPlaneAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var clippingplaneDB orm.ClippingPlaneDB

	// fetch the clippingplane
	_, err := db.First(&clippingplaneDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	clippingplaneDB.CopyBasicFieldsFromClippingPlane_WOP(&input.ClippingPlane_WOP)
	clippingplaneDB.ClippingPlanePointersEncoding = input.ClippingPlanePointersEncoding

	db, _ = db.Model(&clippingplaneDB)
	_, err = db.Updates(&clippingplaneDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	clippingplaneNew := new(models.ClippingPlane)
	clippingplaneDB.CopyBasicFieldsToClippingPlane(clippingplaneNew)

	// redeem pointers
	clippingplaneDB.DecodePointers(backRepo, clippingplaneNew)

	// get stage instance from DB instance, and call callback function
	clippingplaneOld := backRepo.BackRepoClippingPlane.Map_ClippingPlaneDBID_ClippingPlanePtr[clippingplaneDB.ID]
	if clippingplaneOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(clippingplaneOld, clippingplaneNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the clippingplaneDB
	writeJSON(w, http.StatusOK, clippingplaneDB)
}
