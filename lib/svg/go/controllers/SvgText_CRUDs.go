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
var __SvgText__dummysDeclaration__ models.SvgText
var _ = __SvgText__dummysDeclaration__
var __SvgText_time__dummyDeclaration time.Duration
var _ = __SvgText_time__dummyDeclaration

var mutexSvgText sync.Mutex

// An SvgTextID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateSvgText
type SvgTextID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SvgTextInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateSvgText
type SvgTextInput struct {
	// The SvgText to submit or modify
	// in: body
	SvgText *orm.SvgTextAPI
}

// UpdateSvgText
//
// swagger:route PATCH /svgtexts/{ID} svgtexts updateSvgText
//
// # Update a svgtext
//
// Responses:
// default: genericError
//
//	200: svgtextDBResponse
func (controller *Controller) UpdateSvgText(w http.ResponseWriter, r *http.Request) {

	mutexSvgText.Lock()
	defer mutexSvgText.Unlock()

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
	db := backRepo.BackRepoSvgText.GetDB()

	// Validate input
	var input orm.SvgTextAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var svgtextDB orm.SvgTextDB

	// fetch the svgtext
	_, err := db.First(&svgtextDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	svgtextDB.CopyBasicFieldsFromSvgText_WOP(&input.SvgText_WOP)
	svgtextDB.SvgTextPointersEncoding = input.SvgTextPointersEncoding

	db, _ = db.Model(&svgtextDB)
	_, err = db.Updates(&svgtextDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	svgtextNew := new(models.SvgText)
	svgtextDB.CopyBasicFieldsToSvgText(svgtextNew)

	// redeem pointers
	svgtextDB.DecodePointers(backRepo, svgtextNew)

	// get stage instance from DB instance, and call callback function
	svgtextOld := backRepo.BackRepoSvgText.Map_SvgTextDBID_SvgTextPtr[svgtextDB.ID]
	if svgtextOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(svgtextOld, svgtextNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the svgtextDB
	writeJSON(w, http.StatusOK, svgtextDB)
}
