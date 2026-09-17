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
var __LinkAnchoredText__dummysDeclaration__ models.LinkAnchoredText
var _ = __LinkAnchoredText__dummysDeclaration__
var __LinkAnchoredText_time__dummyDeclaration time.Duration
var _ = __LinkAnchoredText_time__dummyDeclaration

var mutexLinkAnchoredText sync.Mutex

// An LinkAnchoredTextID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateLinkAnchoredText
type LinkAnchoredTextID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LinkAnchoredTextInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateLinkAnchoredText
type LinkAnchoredTextInput struct {
	// The LinkAnchoredText to submit or modify
	// in: body
	LinkAnchoredText *orm.LinkAnchoredTextAPI
}

// UpdateLinkAnchoredText
//
// swagger:route PATCH /linkanchoredtexts/{ID} linkanchoredtexts updateLinkAnchoredText
//
// # Update a linkanchoredtext
//
// Responses:
// default: genericError
//
//	200: linkanchoredtextDBResponse
func (controller *Controller) UpdateLinkAnchoredText(w http.ResponseWriter, r *http.Request) {

	mutexLinkAnchoredText.Lock()
	defer mutexLinkAnchoredText.Unlock()

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
	db := backRepo.BackRepoLinkAnchoredText.GetDB()

	// Validate input
	var input orm.LinkAnchoredTextAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var linkanchoredtextDB orm.LinkAnchoredTextDB

	// fetch the linkanchoredtext
	_, err := db.First(&linkanchoredtextDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	linkanchoredtextDB.CopyBasicFieldsFromLinkAnchoredText_WOP(&input.LinkAnchoredText_WOP)
	linkanchoredtextDB.LinkAnchoredTextPointersEncoding = input.LinkAnchoredTextPointersEncoding

	db, _ = db.Model(&linkanchoredtextDB)
	_, err = db.Updates(&linkanchoredtextDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	linkanchoredtextNew := new(models.LinkAnchoredText)
	linkanchoredtextDB.CopyBasicFieldsToLinkAnchoredText(linkanchoredtextNew)

	// redeem pointers
	linkanchoredtextDB.DecodePointers(backRepo, linkanchoredtextNew)

	// get stage instance from DB instance, and call callback function
	linkanchoredtextOld := backRepo.BackRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextPtr[linkanchoredtextDB.ID]
	if linkanchoredtextOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(linkanchoredtextOld, linkanchoredtextNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the linkanchoredtextDB
	writeJSON(w, http.StatusOK, linkanchoredtextDB)
}
