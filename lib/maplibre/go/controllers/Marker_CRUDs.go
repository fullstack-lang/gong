// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/maplibre/go/models"
	"github.com/fullstack-lang/gong/lib/maplibre/go/orm"
)

// declaration in order to justify use of the models import
var __Marker__dummysDeclaration__ models.Marker
var _ = __Marker__dummysDeclaration__
var __Marker_time__dummyDeclaration time.Duration
var _ = __Marker_time__dummyDeclaration

var mutexMarker sync.Mutex

// An MarkerID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateMarker
type MarkerID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MarkerInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateMarker
type MarkerInput struct {
	// The Marker to submit or modify
	// in: body
	Marker *orm.MarkerAPI
}

// UpdateMarker
//
// swagger:route PATCH /markers/{ID} markers updateMarker
//
// # Update a marker
//
// Responses:
// default: genericError
//
//	200: markerDBResponse
func (controller *Controller) UpdateMarker(w http.ResponseWriter, r *http.Request) {

	mutexMarker.Lock()
	defer mutexMarker.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/maplibre/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMarker.GetDB()

	// Validate input
	var input orm.MarkerAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var markerDB orm.MarkerDB

	// fetch the marker
	_, err := db.First(&markerDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	markerDB.CopyBasicFieldsFromMarker_WOP(&input.Marker_WOP)
	markerDB.MarkerPointersEncoding = input.MarkerPointersEncoding

	db, _ = db.Model(&markerDB)
	_, err = db.Updates(&markerDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	markerNew := new(models.Marker)
	markerDB.CopyBasicFieldsToMarker(markerNew)

	// redeem pointers
	markerDB.DecodePointers(backRepo, markerNew)

	// get stage instance from DB instance, and call callback function
	markerOld := backRepo.BackRepoMarker.Map_MarkerDBID_MarkerPtr[markerDB.ID]
	if markerOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(markerOld, markerNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the markerDB
	writeJSON(w, http.StatusOK, markerDB)
}
