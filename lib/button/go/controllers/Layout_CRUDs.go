// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/button/go/models"
	"github.com/fullstack-lang/gong/lib/button/go/orm"
)

// declaration in order to justify use of the models import
var __Layout__dummysDeclaration__ models.Layout
var _ = __Layout__dummysDeclaration__
var __Layout_time__dummyDeclaration time.Duration
var _ = __Layout_time__dummyDeclaration

var mutexLayout sync.Mutex

// An LayoutID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateLayout
type LayoutID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LayoutInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateLayout
type LayoutInput struct {
	// The Layout to submit or modify
	// in: body
	Layout *orm.LayoutAPI
}

// UpdateLayout
//
// swagger:route PATCH /layouts/{ID} layouts updateLayout
//
// # Update a layout
//
// Responses:
// default: genericError
//
//	200: layoutDBResponse
func (controller *Controller) UpdateLayout(w http.ResponseWriter, r *http.Request) {

	mutexLayout.Lock()
	defer mutexLayout.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoLayout.GetDB()

	// Validate input
	var input orm.LayoutAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var layoutDB orm.LayoutDB

	// fetch the layout
	_, err := db.First(&layoutDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	layoutDB.CopyBasicFieldsFromLayout_WOP(&input.Layout_WOP)
	layoutDB.LayoutPointersEncoding = input.LayoutPointersEncoding

	db, _ = db.Model(&layoutDB)
	_, err = db.Updates(&layoutDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	layoutNew := new(models.Layout)
	layoutDB.CopyBasicFieldsToLayout(layoutNew)

	// redeem pointers
	layoutDB.DecodePointers(backRepo, layoutNew)

	// get stage instance from DB instance, and call callback function
	layoutOld := backRepo.BackRepoLayout.Map_LayoutDBID_LayoutPtr[layoutDB.ID]
	if layoutOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(layoutOld, layoutNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the layoutDB
	writeJSON(w, http.StatusOK, layoutDB)
}
