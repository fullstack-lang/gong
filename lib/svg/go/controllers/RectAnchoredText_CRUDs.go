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
var __RectAnchoredText__dummysDeclaration__ models.RectAnchoredText
var _ = __RectAnchoredText__dummysDeclaration__
var __RectAnchoredText_time__dummyDeclaration time.Duration
var _ = __RectAnchoredText_time__dummyDeclaration

var mutexRectAnchoredText sync.Mutex

// An RectAnchoredTextID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateRectAnchoredText
type RectAnchoredTextID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RectAnchoredTextInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateRectAnchoredText
type RectAnchoredTextInput struct {
	// The RectAnchoredText to submit or modify
	// in: body
	RectAnchoredText *orm.RectAnchoredTextAPI
}

// UpdateRectAnchoredText
//
// swagger:route PATCH /rectanchoredtexts/{ID} rectanchoredtexts updateRectAnchoredText
//
// # Update a rectanchoredtext
//
// Responses:
// default: genericError
//
//	200: rectanchoredtextDBResponse
func (controller *Controller) UpdateRectAnchoredText(w http.ResponseWriter, r *http.Request) {

	mutexRectAnchoredText.Lock()
	defer mutexRectAnchoredText.Unlock()

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
	db := backRepo.BackRepoRectAnchoredText.GetDB()

	// Validate input
	var input orm.RectAnchoredTextAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var rectanchoredtextDB orm.RectAnchoredTextDB

	// fetch the rectanchoredtext
	_, err := db.First(&rectanchoredtextDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	rectanchoredtextDB.CopyBasicFieldsFromRectAnchoredText_WOP(&input.RectAnchoredText_WOP)
	rectanchoredtextDB.RectAnchoredTextPointersEncoding = input.RectAnchoredTextPointersEncoding

	db, _ = db.Model(&rectanchoredtextDB)
	_, err = db.Updates(&rectanchoredtextDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	rectanchoredtextNew := new(models.RectAnchoredText)
	rectanchoredtextDB.CopyBasicFieldsToRectAnchoredText(rectanchoredtextNew)

	// redeem pointers
	rectanchoredtextDB.DecodePointers(backRepo, rectanchoredtextNew)

	// get stage instance from DB instance, and call callback function
	rectanchoredtextOld := backRepo.BackRepoRectAnchoredText.Map_RectAnchoredTextDBID_RectAnchoredTextPtr[rectanchoredtextDB.ID]
	if rectanchoredtextOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(rectanchoredtextOld, rectanchoredtextNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the rectanchoredtextDB
	writeJSON(w, http.StatusOK, rectanchoredtextDB)
}
