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
var __Point__dummysDeclaration__ models.Point
var _ = __Point__dummysDeclaration__
var __Point_time__dummyDeclaration time.Duration
var _ = __Point_time__dummyDeclaration

var mutexPoint sync.Mutex

// An PointID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updatePoint
type PointID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PointInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updatePoint
type PointInput struct {
	// The Point to submit or modify
	// in: body
	Point *orm.PointAPI
}

// UpdatePoint
//
// swagger:route PATCH /points/{ID} points updatePoint
//
// # Update a point
//
// Responses:
// default: genericError
//
//	200: pointDBResponse
func (controller *Controller) UpdatePoint(w http.ResponseWriter, r *http.Request) {

	mutexPoint.Lock()
	defer mutexPoint.Unlock()

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
	db := backRepo.BackRepoPoint.GetDB()

	// Validate input
	var input orm.PointAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var pointDB orm.PointDB

	// fetch the point
	_, err := db.First(&pointDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	pointDB.CopyBasicFieldsFromPoint_WOP(&input.Point_WOP)
	pointDB.PointPointersEncoding = input.PointPointersEncoding

	db, _ = db.Model(&pointDB)
	_, err = db.Updates(&pointDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	pointNew := new(models.Point)
	pointDB.CopyBasicFieldsToPoint(pointNew)

	// redeem pointers
	pointDB.DecodePointers(backRepo, pointNew)

	// get stage instance from DB instance, and call callback function
	pointOld := backRepo.BackRepoPoint.Map_PointDBID_PointPtr[pointDB.ID]
	if pointOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(pointOld, pointNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the pointDB
	writeJSON(w, http.StatusOK, pointDB)
}
