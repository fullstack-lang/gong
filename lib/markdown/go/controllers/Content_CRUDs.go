// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/markdown/go/models"
	"github.com/fullstack-lang/gong/lib/markdown/go/orm"
)

// declaration in order to justify use of the models import
var __Content__dummysDeclaration__ models.Content
var _ = __Content__dummysDeclaration__
var __Content_time__dummyDeclaration time.Duration
var _ = __Content_time__dummyDeclaration

var mutexContent sync.Mutex

// An ContentID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateContent
type ContentID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ContentInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateContent
type ContentInput struct {
	// The Content to submit or modify
	// in: body
	Content *orm.ContentAPI
}

// UpdateContent
//
// swagger:route PATCH /contents/{ID} contents updateContent
//
// # Update a content
//
// Responses:
// default: genericError
//
//	200: contentDBResponse
func (controller *Controller) UpdateContent(w http.ResponseWriter, r *http.Request) {

	mutexContent.Lock()
	defer mutexContent.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/markdown/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoContent.GetDB()

	// Validate input
	var input orm.ContentAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var contentDB orm.ContentDB

	// fetch the content
	_, err := db.First(&contentDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	contentDB.CopyBasicFieldsFromContent_WOP(&input.Content_WOP)
	contentDB.ContentPointersEncoding = input.ContentPointersEncoding

	db, _ = db.Model(&contentDB)
	_, err = db.Updates(&contentDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	contentNew := new(models.Content)
	contentDB.CopyBasicFieldsToContent(contentNew)

	// redeem pointers
	contentDB.DecodePointers(backRepo, contentNew)

	// get stage instance from DB instance, and call callback function
	contentOld := backRepo.BackRepoContent.Map_ContentDBID_ContentPtr[contentDB.ID]
	if contentOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(contentOld, contentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the contentDB
	writeJSON(w, http.StatusOK, contentDB)
}
