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
var __Polyline__dummysDeclaration__ models.Polyline
var _ = __Polyline__dummysDeclaration__
var __Polyline_time__dummyDeclaration time.Duration
var _ = __Polyline_time__dummyDeclaration

var mutexPolyline sync.Mutex

// An PolylineID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updatePolyline
type PolylineID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PolylineInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updatePolyline
type PolylineInput struct {
	// The Polyline to submit or modify
	// in: body
	Polyline *orm.PolylineAPI
}

// UpdatePolyline
//
// swagger:route PATCH /polylines/{ID} polylines updatePolyline
//
// # Update a polyline
//
// Responses:
// default: genericError
//
//	200: polylineDBResponse
func (controller *Controller) UpdatePolyline(w http.ResponseWriter, r *http.Request) {

	mutexPolyline.Lock()
	defer mutexPolyline.Unlock()

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
	db := backRepo.BackRepoPolyline.GetDB()

	// Validate input
	var input orm.PolylineAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var polylineDB orm.PolylineDB

	// fetch the polyline
	_, err := db.First(&polylineDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	polylineDB.CopyBasicFieldsFromPolyline_WOP(&input.Polyline_WOP)
	polylineDB.PolylinePointersEncoding = input.PolylinePointersEncoding

	db, _ = db.Model(&polylineDB)
	_, err = db.Updates(&polylineDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	polylineNew := new(models.Polyline)
	polylineDB.CopyBasicFieldsToPolyline(polylineNew)

	// redeem pointers
	polylineDB.DecodePointers(backRepo, polylineNew)

	// get stage instance from DB instance, and call callback function
	polylineOld := backRepo.BackRepoPolyline.Map_PolylineDBID_PolylinePtr[polylineDB.ID]
	if polylineOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(polylineOld, polylineNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the polylineDB
	writeJSON(w, http.StatusOK, polylineDB)
}
