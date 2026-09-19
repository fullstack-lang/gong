// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/splitlite/go/models"
	"github.com/fullstack-lang/gong/lib/splitlite/go/orm"
)

// declaration in order to justify use of the models import
var __View__dummysDeclaration__ models.View
var _ = __View__dummysDeclaration__
var __View_time__dummyDeclaration time.Duration
var _ = __View_time__dummyDeclaration

var mutexView sync.Mutex

// An ViewID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateView
type ViewID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ViewInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateView
type ViewInput struct {
	// The View to submit or modify
	// in: body
	View *orm.ViewAPI
}

// UpdateView
//
// swagger:route PATCH /views/{ID} views updateView
//
// # Update a view
//
// Responses:
// default: genericError
//
//	200: viewDBResponse
func (controller *Controller) UpdateView(w http.ResponseWriter, r *http.Request) {

	mutexView.Lock()
	defer mutexView.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/splitlite/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoView.GetDB()

	// Validate input
	var input orm.ViewAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var viewDB orm.ViewDB

	// fetch the view
	_, err := db.First(&viewDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	viewDB.CopyBasicFieldsFromView_WOP(&input.View_WOP)
	viewDB.ViewPointersEncoding = input.ViewPointersEncoding

	db, _ = db.Model(&viewDB)
	_, err = db.Updates(&viewDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	viewNew := new(models.View)
	viewDB.CopyBasicFieldsToView(viewNew)

	// redeem pointers
	viewDB.DecodePointers(backRepo, viewNew)

	// get stage instance from DB instance, and call callback function
	viewOld := backRepo.BackRepoView.Map_ViewDBID_ViewPtr[viewDB.ID]
	if viewOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(viewOld, viewNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the viewDB
	writeJSON(w, http.StatusOK, viewDB)
}
