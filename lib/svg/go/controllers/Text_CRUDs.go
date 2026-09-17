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
var __Text__dummysDeclaration__ models.Text
var _ = __Text__dummysDeclaration__
var __Text_time__dummyDeclaration time.Duration
var _ = __Text_time__dummyDeclaration

var mutexText sync.Mutex

// An TextID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateText
type TextID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TextInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateText
type TextInput struct {
	// The Text to submit or modify
	// in: body
	Text *orm.TextAPI
}

// UpdateText
//
// swagger:route PATCH /texts/{ID} texts updateText
//
// # Update a text
//
// Responses:
// default: genericError
//
//	200: textDBResponse
func (controller *Controller) UpdateText(w http.ResponseWriter, r *http.Request) {

	mutexText.Lock()
	defer mutexText.Unlock()

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
	db := backRepo.BackRepoText.GetDB()

	// Validate input
	var input orm.TextAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var textDB orm.TextDB

	// fetch the text
	_, err := db.First(&textDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	textDB.CopyBasicFieldsFromText_WOP(&input.Text_WOP)
	textDB.TextPointersEncoding = input.TextPointersEncoding

	db, _ = db.Model(&textDB)
	_, err = db.Updates(&textDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	textNew := new(models.Text)
	textDB.CopyBasicFieldsToText(textNew)

	// redeem pointers
	textDB.DecodePointers(backRepo, textNew)

	// get stage instance from DB instance, and call callback function
	textOld := backRepo.BackRepoText.Map_TextDBID_TextPtr[textDB.ID]
	if textOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(textOld, textNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the textDB
	writeJSON(w, http.StatusOK, textDB)
}
