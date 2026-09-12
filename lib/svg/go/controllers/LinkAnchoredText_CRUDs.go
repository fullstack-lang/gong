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
// swagger:parameters getLinkAnchoredText updateLinkAnchoredText deleteLinkAnchoredText
type LinkAnchoredTextID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LinkAnchoredTextInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLinkAnchoredText updateLinkAnchoredText
type LinkAnchoredTextInput struct {
	// The LinkAnchoredText to submit or modify
	// in: body
	LinkAnchoredText *orm.LinkAnchoredTextAPI
}

// GetLinkAnchoredTexts
//
// swagger:route GET /linkanchoredtexts linkanchoredtexts getLinkAnchoredTexts
//
// # Get all linkanchoredtexts
//
// Responses:
// default: genericError
//
//	200: linkanchoredtextDBResponse
func (controller *Controller) GetLinkAnchoredTexts(w http.ResponseWriter, r *http.Request) {

	// source slice
	var linkanchoredtextDBs []orm.LinkAnchoredTextDB

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLinkAnchoredTexts", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoLinkAnchoredText.GetDB()

	_, err := db.Find(&linkanchoredtextDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	linkanchoredtextAPIs := make([]orm.LinkAnchoredTextAPI, 0)

	// for each linkanchoredtext, update fields from the database nullable fields
	for idx := range linkanchoredtextDBs {
		linkanchoredtextDB := &linkanchoredtextDBs[idx]
		_ = linkanchoredtextDB
		var linkanchoredtextAPI orm.LinkAnchoredTextAPI

		// insertion point for updating fields
		linkanchoredtextAPI.ID = linkanchoredtextDB.ID
		linkanchoredtextDB.CopyBasicFieldsToLinkAnchoredText_WOP(&linkanchoredtextAPI.LinkAnchoredText_WOP)
		linkanchoredtextAPI.LinkAnchoredTextPointersEncoding = linkanchoredtextDB.LinkAnchoredTextPointersEncoding
		linkanchoredtextAPIs = append(linkanchoredtextAPIs, linkanchoredtextAPI)
	}

	writeJSON(w, http.StatusOK, linkanchoredtextAPIs)
}

// PostLinkAnchoredText
//
// swagger:route POST /linkanchoredtexts linkanchoredtexts postLinkAnchoredText
//
// Creates a linkanchoredtext
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLinkAnchoredText(w http.ResponseWriter, r *http.Request) {

	mutexLinkAnchoredText.Lock()
	defer mutexLinkAnchoredText.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLinkAnchoredTexts", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoLinkAnchoredText.GetDB()

	// Validate input
	var input orm.LinkAnchoredTextAPI

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// Create linkanchoredtext
	linkanchoredtextDB := orm.LinkAnchoredTextDB{}
	linkanchoredtextDB.LinkAnchoredTextPointersEncoding = input.LinkAnchoredTextPointersEncoding
	linkanchoredtextDB.CopyBasicFieldsFromLinkAnchoredText_WOP(&input.LinkAnchoredText_WOP)

	_, err = db.Create(&linkanchoredtextDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLinkAnchoredText.CheckoutPhaseOneInstance(&linkanchoredtextDB)
	linkanchoredtext := backRepo.BackRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextPtr[linkanchoredtextDB.ID]

	if linkanchoredtext != nil {
		backRepo.GetStage().AfterCreateFromFront(linkanchoredtext)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, linkanchoredtextDB)
}

// GetLinkAnchoredText
//
// swagger:route GET /linkanchoredtexts/{ID} linkanchoredtexts getLinkAnchoredText
//
// Gets the details for a linkanchoredtext.
//
// Responses:
// default: genericError
//
//	200: linkanchoredtextDBResponse
func (controller *Controller) GetLinkAnchoredText(w http.ResponseWriter, r *http.Request) {

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLinkAnchoredText", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoLinkAnchoredText.GetDB()

	// Get linkanchoredtextDB in DB
	var linkanchoredtextDB orm.LinkAnchoredTextDB
	if _, err := db.First(&linkanchoredtextDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	var linkanchoredtextAPI orm.LinkAnchoredTextAPI
	linkanchoredtextAPI.ID = linkanchoredtextDB.ID
	linkanchoredtextAPI.LinkAnchoredTextPointersEncoding = linkanchoredtextDB.LinkAnchoredTextPointersEncoding
	linkanchoredtextDB.CopyBasicFieldsToLinkAnchoredText_WOP(&linkanchoredtextAPI.LinkAnchoredText_WOP)

	writeJSON(w, http.StatusOK, linkanchoredtextAPI)
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

// DeleteLinkAnchoredText
//
// swagger:route DELETE /linkanchoredtexts/{ID} linkanchoredtexts deleteLinkAnchoredText
//
// # Delete a linkanchoredtext
//
// default: genericError
//
//	200: linkanchoredtextDBResponse
func (controller *Controller) DeleteLinkAnchoredText(w http.ResponseWriter, r *http.Request) {

	mutexLinkAnchoredText.Lock()
	defer mutexLinkAnchoredText.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLinkAnchoredText", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoLinkAnchoredText.GetDB()

	// Get model if exist
	var linkanchoredtextDB orm.LinkAnchoredTextDB
	if _, err := db.First(&linkanchoredtextDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&linkanchoredtextDB)

	// get an instance (not staged) from DB instance, and call callback function
	linkanchoredtextDeleted := new(models.LinkAnchoredText)
	linkanchoredtextDB.CopyBasicFieldsToLinkAnchoredText(linkanchoredtextDeleted)

	// get stage instance from DB instance, and call callback function
	linkanchoredtextStaged := backRepo.BackRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextPtr[linkanchoredtextDB.ID]
	if linkanchoredtextStaged != nil {
		backRepo.GetStage().AfterDeleteFromFront(linkanchoredtextStaged, linkanchoredtextDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, H{"data": true})
}
