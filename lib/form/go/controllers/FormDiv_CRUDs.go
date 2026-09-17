// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/form/go/models"
	"github.com/fullstack-lang/gong/lib/form/go/orm"
)

// declaration in order to justify use of the models import
var __FormDiv__dummysDeclaration__ models.FormDiv
var _ = __FormDiv__dummysDeclaration__
var __FormDiv_time__dummyDeclaration time.Duration
var _ = __FormDiv_time__dummyDeclaration

var mutexFormDiv sync.Mutex

// An FormDivID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateFormDiv
type FormDivID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormDivInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateFormDiv
type FormDivInput struct {
	// The FormDiv to submit or modify
	// in: body
	FormDiv *orm.FormDivAPI
}

// UpdateFormDiv
//
// swagger:route PATCH /formdivs/{ID} formdivs updateFormDiv
//
// # Update a formdiv
//
// Responses:
// default: genericError
//
//	200: formdivDBResponse
func (controller *Controller) UpdateFormDiv(w http.ResponseWriter, r *http.Request) {

	mutexFormDiv.Lock()
	defer mutexFormDiv.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/form/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFormDiv.GetDB()

	// Validate input
	var input orm.FormDivAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formdivDB orm.FormDivDB

	// fetch the formdiv
	_, err := db.First(&formdivDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formdivDB.CopyBasicFieldsFromFormDiv_WOP(&input.FormDiv_WOP)
	formdivDB.FormDivPointersEncoding = input.FormDivPointersEncoding

	db, _ = db.Model(&formdivDB)
	_, err = db.Updates(&formdivDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formdivNew := new(models.FormDiv)
	formdivDB.CopyBasicFieldsToFormDiv(formdivNew)

	// redeem pointers
	formdivDB.DecodePointers(backRepo, formdivNew)

	// get stage instance from DB instance, and call callback function
	formdivOld := backRepo.BackRepoFormDiv.Map_FormDivDBID_FormDivPtr[formdivDB.ID]
	if formdivOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(formdivOld, formdivNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formdivDB
	writeJSON(w, http.StatusOK, formdivDB)
}
