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
var __LinkAnchoredPath__dummysDeclaration__ models.LinkAnchoredPath
var _ = __LinkAnchoredPath__dummysDeclaration__
var __LinkAnchoredPath_time__dummyDeclaration time.Duration
var _ = __LinkAnchoredPath_time__dummyDeclaration

var mutexLinkAnchoredPath sync.Mutex

// An LinkAnchoredPathID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateLinkAnchoredPath
type LinkAnchoredPathID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LinkAnchoredPathInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateLinkAnchoredPath
type LinkAnchoredPathInput struct {
	// The LinkAnchoredPath to submit or modify
	// in: body
	LinkAnchoredPath *orm.LinkAnchoredPathAPI
}

// UpdateLinkAnchoredPath
//
// swagger:route PATCH /linkanchoredpaths/{ID} linkanchoredpaths updateLinkAnchoredPath
//
// # Update a linkanchoredpath
//
// Responses:
// default: genericError
//
//	200: linkanchoredpathDBResponse
func (controller *Controller) UpdateLinkAnchoredPath(w http.ResponseWriter, r *http.Request) {

	mutexLinkAnchoredPath.Lock()
	defer mutexLinkAnchoredPath.Unlock()

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
	db := backRepo.BackRepoLinkAnchoredPath.GetDB()

	// Validate input
	var input orm.LinkAnchoredPathAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var linkanchoredpathDB orm.LinkAnchoredPathDB

	// fetch the linkanchoredpath
	_, err := db.First(&linkanchoredpathDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	linkanchoredpathDB.CopyBasicFieldsFromLinkAnchoredPath_WOP(&input.LinkAnchoredPath_WOP)
	linkanchoredpathDB.LinkAnchoredPathPointersEncoding = input.LinkAnchoredPathPointersEncoding

	db, _ = db.Model(&linkanchoredpathDB)
	_, err = db.Updates(&linkanchoredpathDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	linkanchoredpathNew := new(models.LinkAnchoredPath)
	linkanchoredpathDB.CopyBasicFieldsToLinkAnchoredPath(linkanchoredpathNew)

	// redeem pointers
	linkanchoredpathDB.DecodePointers(backRepo, linkanchoredpathNew)

	// get stage instance from DB instance, and call callback function
	linkanchoredpathOld := backRepo.BackRepoLinkAnchoredPath.Map_LinkAnchoredPathDBID_LinkAnchoredPathPtr[linkanchoredpathDB.ID]
	if linkanchoredpathOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(linkanchoredpathOld, linkanchoredpathNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the linkanchoredpathDB
	writeJSON(w, http.StatusOK, linkanchoredpathDB)
}
