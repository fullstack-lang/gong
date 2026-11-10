// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/split/go/models"
	"github.com/fullstack-lang/gong/lib/split/go/orm"
)

// declaration in order to justify use of the models import
var __Cesium__dummysDeclaration__ models.Cesium
var _ = __Cesium__dummysDeclaration__
var __Cesium_time__dummyDeclaration time.Duration
var _ = __Cesium_time__dummyDeclaration

var mutexCesium sync.Mutex

// An CesiumID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateCesium
type CesiumID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CesiumInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateCesium
type CesiumInput struct {
	// The Cesium to submit or modify
	// in: body
	Cesium *orm.CesiumAPI
}

// UpdateCesium
//
// swagger:route PATCH /cesiums/{ID} cesiums updateCesium
//
// # Update a cesium
//
// Responses:
// default: genericError
//
//	200: cesiumDBResponse
func (controller *Controller) UpdateCesium(w http.ResponseWriter, r *http.Request) {

	mutexCesium.Lock()
	defer mutexCesium.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoCesium.GetDB()

	// Validate input
	var input orm.CesiumAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var cesiumDB orm.CesiumDB

	// fetch the cesium
	_, err := db.First(&cesiumDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	cesiumDB.CopyBasicFieldsFromCesium_WOP(&input.Cesium_WOP)
	cesiumDB.CesiumPointersEncoding = input.CesiumPointersEncoding

	db, _ = db.Model(&cesiumDB)
	_, err = db.Updates(&cesiumDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	cesiumNew := new(models.Cesium)
	cesiumDB.CopyBasicFieldsToCesium(cesiumNew)

	// redeem pointers
	cesiumDB.DecodePointers(backRepo, cesiumNew)

	// get stage instance from DB instance, and call callback function
	cesiumOld := backRepo.BackRepoCesium.Map_CesiumDBID_CesiumPtr[cesiumDB.ID]
	if cesiumOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(cesiumOld, cesiumNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the cesiumDB
	writeJSON(w, http.StatusOK, cesiumDB)
}
