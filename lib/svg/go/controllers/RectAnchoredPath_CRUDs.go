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
var __RectAnchoredPath__dummysDeclaration__ models.RectAnchoredPath
var _ = __RectAnchoredPath__dummysDeclaration__
var __RectAnchoredPath_time__dummyDeclaration time.Duration
var _ = __RectAnchoredPath_time__dummyDeclaration

var mutexRectAnchoredPath sync.Mutex

// An RectAnchoredPathID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateRectAnchoredPath
type RectAnchoredPathID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RectAnchoredPathInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateRectAnchoredPath
type RectAnchoredPathInput struct {
	// The RectAnchoredPath to submit or modify
	// in: body
	RectAnchoredPath *orm.RectAnchoredPathAPI
}

// UpdateRectAnchoredPath
//
// swagger:route PATCH /rectanchoredpaths/{ID} rectanchoredpaths updateRectAnchoredPath
//
// # Update a rectanchoredpath
//
// Responses:
// default: genericError
//
//	200: rectanchoredpathDBResponse
func (controller *Controller) UpdateRectAnchoredPath(w http.ResponseWriter, r *http.Request) {

	mutexRectAnchoredPath.Lock()
	defer mutexRectAnchoredPath.Unlock()

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
	db := backRepo.BackRepoRectAnchoredPath.GetDB()

	// Validate input
	var input orm.RectAnchoredPathAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var rectanchoredpathDB orm.RectAnchoredPathDB

	// fetch the rectanchoredpath
	_, err := db.First(&rectanchoredpathDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	rectanchoredpathDB.CopyBasicFieldsFromRectAnchoredPath_WOP(&input.RectAnchoredPath_WOP)
	rectanchoredpathDB.RectAnchoredPathPointersEncoding = input.RectAnchoredPathPointersEncoding

	db, _ = db.Model(&rectanchoredpathDB)
	_, err = db.Updates(&rectanchoredpathDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	rectanchoredpathNew := new(models.RectAnchoredPath)
	rectanchoredpathDB.CopyBasicFieldsToRectAnchoredPath(rectanchoredpathNew)

	// redeem pointers
	rectanchoredpathDB.DecodePointers(backRepo, rectanchoredpathNew)

	// get stage instance from DB instance, and call callback function
	rectanchoredpathOld := backRepo.BackRepoRectAnchoredPath.Map_RectAnchoredPathDBID_RectAnchoredPathPtr[rectanchoredpathDB.ID]
	if rectanchoredpathOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(rectanchoredpathOld, rectanchoredpathNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the rectanchoredpathDB
	writeJSON(w, http.StatusOK, rectanchoredpathDB)
}
