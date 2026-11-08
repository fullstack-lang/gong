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
var __Polygon__dummysDeclaration__ models.Polygon
var _ = __Polygon__dummysDeclaration__
var __Polygon_time__dummyDeclaration time.Duration
var _ = __Polygon_time__dummyDeclaration

var mutexPolygon sync.Mutex

// An PolygonID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updatePolygon
type PolygonID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PolygonInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updatePolygon
type PolygonInput struct {
	// The Polygon to submit or modify
	// in: body
	Polygon *orm.PolygonAPI
}

// UpdatePolygon
//
// swagger:route PATCH /polygons/{ID} polygons updatePolygon
//
// # Update a polygon
//
// Responses:
// default: genericError
//
//	200: polygonDBResponse
func (controller *Controller) UpdatePolygon(w http.ResponseWriter, r *http.Request) {

	mutexPolygon.Lock()
	defer mutexPolygon.Unlock()

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
	db := backRepo.BackRepoPolygon.GetDB()

	// Validate input
	var input orm.PolygonAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var polygonDB orm.PolygonDB

	// fetch the polygon
	_, err := db.First(&polygonDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	polygonDB.CopyBasicFieldsFromPolygon_WOP(&input.Polygon_WOP)
	polygonDB.PolygonPointersEncoding = input.PolygonPointersEncoding

	db, _ = db.Model(&polygonDB)
	_, err = db.Updates(&polygonDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	polygonNew := new(models.Polygon)
	polygonDB.CopyBasicFieldsToPolygon(polygonNew)

	// redeem pointers
	polygonDB.DecodePointers(backRepo, polygonNew)

	// get stage instance from DB instance, and call callback function
	polygonOld := backRepo.BackRepoPolygon.Map_PolygonDBID_PolygonPtr[polygonDB.ID]
	if polygonOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(polygonOld, polygonNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the polygonDB
	writeJSON(w, http.StatusOK, polygonDB)
}
