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
var __Curve__dummysDeclaration__ models.Curve
var _ = __Curve__dummysDeclaration__
var __Curve_time__dummyDeclaration time.Duration
var _ = __Curve_time__dummyDeclaration

var mutexCurve sync.Mutex

// An CurveID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateCurve
type CurveID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CurveInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateCurve
type CurveInput struct {
	// The Curve to submit or modify
	// in: body
	Curve *orm.CurveAPI
}

// UpdateCurve
//
// swagger:route PATCH /curves/{ID} curves updateCurve
//
// # Update a curve
//
// Responses:
// default: genericError
//
//	200: curveDBResponse
func (controller *Controller) UpdateCurve(w http.ResponseWriter, r *http.Request) {

	mutexCurve.Lock()
	defer mutexCurve.Unlock()

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
	db := backRepo.BackRepoCurve.GetDB()

	// Validate input
	var input orm.CurveAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var curveDB orm.CurveDB

	// fetch the curve
	_, err := db.First(&curveDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	curveDB.CopyBasicFieldsFromCurve_WOP(&input.Curve_WOP)
	curveDB.CurvePointersEncoding = input.CurvePointersEncoding

	db, _ = db.Model(&curveDB)
	_, err = db.Updates(&curveDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	curveNew := new(models.Curve)
	curveDB.CopyBasicFieldsToCurve(curveNew)

	// redeem pointers
	curveDB.DecodePointers(backRepo, curveNew)

	// get stage instance from DB instance, and call callback function
	curveOld := backRepo.BackRepoCurve.Map_CurveDBID_CurvePtr[curveDB.ID]
	if curveOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(curveOld, curveNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the curveDB
	writeJSON(w, http.StatusOK, curveDB)
}
