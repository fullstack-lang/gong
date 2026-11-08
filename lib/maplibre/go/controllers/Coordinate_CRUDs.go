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
var __Coordinate__dummysDeclaration__ models.Coordinate
var _ = __Coordinate__dummysDeclaration__
var __Coordinate_time__dummyDeclaration time.Duration
var _ = __Coordinate_time__dummyDeclaration

var mutexCoordinate sync.Mutex

// An CoordinateID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateCoordinate
type CoordinateID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CoordinateInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateCoordinate
type CoordinateInput struct {
	// The Coordinate to submit or modify
	// in: body
	Coordinate *orm.CoordinateAPI
}

// UpdateCoordinate
//
// swagger:route PATCH /coordinates/{ID} coordinates updateCoordinate
//
// # Update a coordinate
//
// Responses:
// default: genericError
//
//	200: coordinateDBResponse
func (controller *Controller) UpdateCoordinate(w http.ResponseWriter, r *http.Request) {

	mutexCoordinate.Lock()
	defer mutexCoordinate.Unlock()

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
	db := backRepo.BackRepoCoordinate.GetDB()

	// Validate input
	var input orm.CoordinateAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var coordinateDB orm.CoordinateDB

	// fetch the coordinate
	_, err := db.First(&coordinateDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	coordinateDB.CopyBasicFieldsFromCoordinate_WOP(&input.Coordinate_WOP)
	coordinateDB.CoordinatePointersEncoding = input.CoordinatePointersEncoding

	db, _ = db.Model(&coordinateDB)
	_, err = db.Updates(&coordinateDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	coordinateNew := new(models.Coordinate)
	coordinateDB.CopyBasicFieldsToCoordinate(coordinateNew)

	// redeem pointers
	coordinateDB.DecodePointers(backRepo, coordinateNew)

	// get stage instance from DB instance, and call callback function
	coordinateOld := backRepo.BackRepoCoordinate.Map_CoordinateDBID_CoordinatePtr[coordinateDB.ID]
	if coordinateOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(coordinateOld, coordinateNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the coordinateDB
	writeJSON(w, http.StatusOK, coordinateDB)
}
