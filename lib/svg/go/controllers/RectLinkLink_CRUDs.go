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
var __RectLinkLink__dummysDeclaration__ models.RectLinkLink
var _ = __RectLinkLink__dummysDeclaration__
var __RectLinkLink_time__dummyDeclaration time.Duration
var _ = __RectLinkLink_time__dummyDeclaration

var mutexRectLinkLink sync.Mutex

// An RectLinkLinkID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateRectLinkLink
type RectLinkLinkID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RectLinkLinkInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateRectLinkLink
type RectLinkLinkInput struct {
	// The RectLinkLink to submit or modify
	// in: body
	RectLinkLink *orm.RectLinkLinkAPI
}

// UpdateRectLinkLink
//
// swagger:route PATCH /rectlinklinks/{ID} rectlinklinks updateRectLinkLink
//
// # Update a rectlinklink
//
// Responses:
// default: genericError
//
//	200: rectlinklinkDBResponse
func (controller *Controller) UpdateRectLinkLink(w http.ResponseWriter, r *http.Request) {

	mutexRectLinkLink.Lock()
	defer mutexRectLinkLink.Unlock()

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
	db := backRepo.BackRepoRectLinkLink.GetDB()

	// Validate input
	var input orm.RectLinkLinkAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var rectlinklinkDB orm.RectLinkLinkDB

	// fetch the rectlinklink
	_, err := db.First(&rectlinklinkDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	rectlinklinkDB.CopyBasicFieldsFromRectLinkLink_WOP(&input.RectLinkLink_WOP)
	rectlinklinkDB.RectLinkLinkPointersEncoding = input.RectLinkLinkPointersEncoding

	db, _ = db.Model(&rectlinklinkDB)
	_, err = db.Updates(&rectlinklinkDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	rectlinklinkNew := new(models.RectLinkLink)
	rectlinklinkDB.CopyBasicFieldsToRectLinkLink(rectlinklinkNew)

	// redeem pointers
	rectlinklinkDB.DecodePointers(backRepo, rectlinklinkNew)

	// get stage instance from DB instance, and call callback function
	rectlinklinkOld := backRepo.BackRepoRectLinkLink.Map_RectLinkLinkDBID_RectLinkLinkPtr[rectlinklinkDB.ID]
	if rectlinklinkOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(rectlinklinkOld, rectlinklinkNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the rectlinklinkDB
	writeJSON(w, http.StatusOK, rectlinklinkDB)
}
