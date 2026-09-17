// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/split/go/models"
	"github.com/fullstack-lang/gong/lib/split/go/orm"
)

// declaration in order to justify use of the models import
var __Markdown__dummysDeclaration__ models.Markdown
var _ = __Markdown__dummysDeclaration__
var __Markdown_time__dummyDeclaration time.Duration
var _ = __Markdown_time__dummyDeclaration

var mutexMarkdown sync.Mutex

// An MarkdownID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateMarkdown
type MarkdownID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MarkdownInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateMarkdown
type MarkdownInput struct {
	// The Markdown to submit or modify
	// in: body
	Markdown *orm.MarkdownAPI
}

// UpdateMarkdown
//
// swagger:route PATCH /markdowns/{ID} markdowns updateMarkdown
//
// # Update a markdown
//
// Responses:
// default: genericError
//
//	200: markdownDBResponse
func (controller *Controller) UpdateMarkdown(w http.ResponseWriter, r *http.Request) {

	mutexMarkdown.Lock()
	defer mutexMarkdown.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMarkdown.GetDB()

	// Validate input
	var input orm.MarkdownAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var markdownDB orm.MarkdownDB

	// fetch the markdown
	_, err := db.First(&markdownDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	markdownDB.CopyBasicFieldsFromMarkdown_WOP(&input.Markdown_WOP)
	markdownDB.MarkdownPointersEncoding = input.MarkdownPointersEncoding

	db, _ = db.Model(&markdownDB)
	_, err = db.Updates(&markdownDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	markdownNew := new(models.Markdown)
	markdownDB.CopyBasicFieldsToMarkdown(markdownNew)

	// redeem pointers
	markdownDB.DecodePointers(backRepo, markdownNew)

	// get stage instance from DB instance, and call callback function
	markdownOld := backRepo.BackRepoMarkdown.Map_MarkdownDBID_MarkdownPtr[markdownDB.ID]
	if markdownOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(markdownOld, markdownNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the markdownDB
	writeJSON(w, http.StatusOK, markdownDB)
}
