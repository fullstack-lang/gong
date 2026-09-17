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
var __Canvas__dummysDeclaration__ models.Canvas
var _ = __Canvas__dummysDeclaration__
var __Canvas_time__dummyDeclaration time.Duration
var _ = __Canvas_time__dummyDeclaration

var mutexCanvas sync.Mutex

// An CanvasID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateCanvas
type CanvasID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CanvasInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateCanvas
type CanvasInput struct {
	// The Canvas to submit or modify
	// in: body
	Canvas *orm.CanvasAPI
}

// UpdateCanvas
//
// swagger:route PATCH /canvass/{ID} canvass updateCanvas
//
// # Update a canvas
//
// Responses:
// default: genericError
//
//	200: canvasDBResponse
func (controller *Controller) UpdateCanvas(w http.ResponseWriter, r *http.Request) {

	mutexCanvas.Lock()
	defer mutexCanvas.Unlock()

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
	db := backRepo.BackRepoCanvas.GetDB()

	// Validate input
	var input orm.CanvasAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var canvasDB orm.CanvasDB

	// fetch the canvas
	_, err := db.First(&canvasDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	canvasDB.CopyBasicFieldsFromCanvas_WOP(&input.Canvas_WOP)
	canvasDB.CanvasPointersEncoding = input.CanvasPointersEncoding

	db, _ = db.Model(&canvasDB)
	_, err = db.Updates(&canvasDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	canvasNew := new(models.Canvas)
	canvasDB.CopyBasicFieldsToCanvas(canvasNew)

	// redeem pointers
	canvasDB.DecodePointers(backRepo, canvasNew)

	// get stage instance from DB instance, and call callback function
	canvasOld := backRepo.BackRepoCanvas.Map_CanvasDBID_CanvasPtr[canvasDB.ID]
	if canvasOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(canvasOld, canvasNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the canvasDB
	writeJSON(w, http.StatusOK, canvasDB)
}
