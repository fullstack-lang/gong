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
var __Rect__dummysDeclaration__ models.Rect
var _ = __Rect__dummysDeclaration__
var __Rect_time__dummyDeclaration time.Duration
var _ = __Rect_time__dummyDeclaration

var mutexRect sync.Mutex

// An RectID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateRect
type RectID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RectInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateRect
type RectInput struct {
	// The Rect to submit or modify
	// in: body
	Rect *orm.RectAPI
}

// UpdateRect
//
// swagger:route PATCH /rects/{ID} rects updateRect
//
// # Update a rect
//
// Responses:
// default: genericError
//
//	200: rectDBResponse
func (controller *Controller) UpdateRect(w http.ResponseWriter, r *http.Request) {

	mutexRect.Lock()
	defer mutexRect.Unlock()

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
	db := backRepo.BackRepoRect.GetDB()

	// Validate input
	var input orm.RectAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var rectDB orm.RectDB

	// fetch the rect
	_, err := db.First(&rectDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	rectDB.CopyBasicFieldsFromRect_WOP(&input.Rect_WOP)
	rectDB.RectPointersEncoding = input.RectPointersEncoding

	db, _ = db.Model(&rectDB)
	_, err = db.Updates(&rectDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	rectNew := new(models.Rect)
	rectDB.CopyBasicFieldsToRect(rectNew)

	// redeem pointers
	rectDB.DecodePointers(backRepo, rectNew)

	// get stage instance from DB instance, and call callback function
	rectOld := backRepo.BackRepoRect.Map_RectDBID_RectPtr[rectDB.ID]
	if rectOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(rectOld, rectNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the rectDB
	writeJSON(w, http.StatusOK, rectDB)
}
