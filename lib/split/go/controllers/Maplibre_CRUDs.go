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
var __Maplibre__dummysDeclaration__ models.Maplibre
var _ = __Maplibre__dummysDeclaration__
var __Maplibre_time__dummyDeclaration time.Duration
var _ = __Maplibre_time__dummyDeclaration

var mutexMaplibre sync.Mutex

// An MaplibreID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateMaplibre
type MaplibreID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MaplibreInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateMaplibre
type MaplibreInput struct {
	// The Maplibre to submit or modify
	// in: body
	Maplibre *orm.MaplibreAPI
}

// UpdateMaplibre
//
// swagger:route PATCH /maplibres/{ID} maplibres updateMaplibre
//
// # Update a maplibre
//
// Responses:
// default: genericError
//
//	200: maplibreDBResponse
func (controller *Controller) UpdateMaplibre(w http.ResponseWriter, r *http.Request) {

	mutexMaplibre.Lock()
	defer mutexMaplibre.Unlock()

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
	db := backRepo.BackRepoMaplibre.GetDB()

	// Validate input
	var input orm.MaplibreAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var maplibreDB orm.MaplibreDB

	// fetch the maplibre
	_, err := db.First(&maplibreDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	maplibreDB.CopyBasicFieldsFromMaplibre_WOP(&input.Maplibre_WOP)
	maplibreDB.MaplibrePointersEncoding = input.MaplibrePointersEncoding

	db, _ = db.Model(&maplibreDB)
	_, err = db.Updates(&maplibreDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	maplibreNew := new(models.Maplibre)
	maplibreDB.CopyBasicFieldsToMaplibre(maplibreNew)

	// redeem pointers
	maplibreDB.DecodePointers(backRepo, maplibreNew)

	// get stage instance from DB instance, and call callback function
	maplibreOld := backRepo.BackRepoMaplibre.Map_MaplibreDBID_MaplibrePtr[maplibreDB.ID]
	if maplibreOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(maplibreOld, maplibreNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the maplibreDB
	writeJSON(w, http.StatusOK, maplibreDB)
}
