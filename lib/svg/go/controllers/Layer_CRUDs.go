// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
	"github.com/fullstack-lang/gong/lib/svg/go/orm"
)

// declaration in order to justify use of the models import
var __Layer__dummysDeclaration__ models.Layer
var _ = __Layer__dummysDeclaration__
var __Layer_time__dummyDeclaration time.Duration
var _ = __Layer_time__dummyDeclaration

var mutexLayer sync.Mutex

// An LayerID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateLayer
type LayerID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LayerInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateLayer
type LayerInput struct {
	// The Layer to submit or modify
	// in: body
	Layer *orm.LayerAPI
}

// UpdateLayer
//
// swagger:route PATCH /layers/{ID} layers updateLayer
//
// # Update a layer
//
// Responses:
// default: genericError
//
//	200: layerDBResponse
func (controller *Controller) UpdateLayer(w http.ResponseWriter, r *http.Request) {

	mutexLayer.Lock()
	defer mutexLayer.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoLayer.GetDB()

	// Validate input
	var input orm.LayerAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var layerDB orm.LayerDB

	// fetch the layer
	_, err := db.First(&layerDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	layerDB.CopyBasicFieldsFromLayer_WOP(&input.Layer_WOP)
	layerDB.LayerPointersEncoding = input.LayerPointersEncoding

	db, _ = db.Model(&layerDB)
	_, err = db.Updates(&layerDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	layerNew := new(models.Layer)
	layerDB.CopyBasicFieldsToLayer(layerNew)

	// redeem pointers
	layerDB.DecodePointers(backRepo, layerNew)

	// get stage instance from DB instance, and call callback function
	layerOld := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[layerDB.ID]
	if layerOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(layerOld, layerNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the layerDB
	writeJSON(w, http.StatusOK, layerDB)
}
