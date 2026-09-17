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
var __Title__dummysDeclaration__ models.Title
var _ = __Title__dummysDeclaration__
var __Title_time__dummyDeclaration time.Duration
var _ = __Title_time__dummyDeclaration

var mutexTitle sync.Mutex

// An TitleID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateTitle
type TitleID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TitleInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateTitle
type TitleInput struct {
	// The Title to submit or modify
	// in: body
	Title *orm.TitleAPI
}

// UpdateTitle
//
// swagger:route PATCH /titles/{ID} titles updateTitle
//
// # Update a title
//
// Responses:
// default: genericError
//
//	200: titleDBResponse
func (controller *Controller) UpdateTitle(w http.ResponseWriter, r *http.Request) {

	mutexTitle.Lock()
	defer mutexTitle.Unlock()

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
	db := backRepo.BackRepoTitle.GetDB()

	// Validate input
	var input orm.TitleAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var titleDB orm.TitleDB

	// fetch the title
	_, err := db.First(&titleDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	titleDB.CopyBasicFieldsFromTitle_WOP(&input.Title_WOP)
	titleDB.TitlePointersEncoding = input.TitlePointersEncoding

	db, _ = db.Model(&titleDB)
	_, err = db.Updates(&titleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	titleNew := new(models.Title)
	titleDB.CopyBasicFieldsToTitle(titleNew)

	// redeem pointers
	titleDB.DecodePointers(backRepo, titleNew)

	// get stage instance from DB instance, and call callback function
	titleOld := backRepo.BackRepoTitle.Map_TitleDBID_TitlePtr[titleDB.ID]
	if titleOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(titleOld, titleNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the titleDB
	writeJSON(w, http.StatusOK, titleDB)
}
