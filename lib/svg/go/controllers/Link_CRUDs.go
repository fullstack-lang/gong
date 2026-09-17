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
var __Link__dummysDeclaration__ models.Link
var _ = __Link__dummysDeclaration__
var __Link_time__dummyDeclaration time.Duration
var _ = __Link_time__dummyDeclaration

var mutexLink sync.Mutex

// An LinkID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateLink
type LinkID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LinkInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateLink
type LinkInput struct {
	// The Link to submit or modify
	// in: body
	Link *orm.LinkAPI
}

// UpdateLink
//
// swagger:route PATCH /links/{ID} links updateLink
//
// # Update a link
//
// Responses:
// default: genericError
//
//	200: linkDBResponse
func (controller *Controller) UpdateLink(w http.ResponseWriter, r *http.Request) {

	mutexLink.Lock()
	defer mutexLink.Unlock()

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
	db := backRepo.BackRepoLink.GetDB()

	// Validate input
	var input orm.LinkAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var linkDB orm.LinkDB

	// fetch the link
	_, err := db.First(&linkDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	linkDB.CopyBasicFieldsFromLink_WOP(&input.Link_WOP)
	linkDB.LinkPointersEncoding = input.LinkPointersEncoding

	db, _ = db.Model(&linkDB)
	_, err = db.Updates(&linkDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	linkNew := new(models.Link)
	linkDB.CopyBasicFieldsToLink(linkNew)

	// redeem pointers
	linkDB.DecodePointers(backRepo, linkNew)

	// get stage instance from DB instance, and call callback function
	linkOld := backRepo.BackRepoLink.Map_LinkDBID_LinkPtr[linkDB.ID]
	if linkOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(linkOld, linkNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the linkDB
	writeJSON(w, http.StatusOK, linkDB)
}
