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
var __LogoOnTheRight__dummysDeclaration__ models.LogoOnTheRight
var _ = __LogoOnTheRight__dummysDeclaration__
var __LogoOnTheRight_time__dummyDeclaration time.Duration
var _ = __LogoOnTheRight_time__dummyDeclaration

var mutexLogoOnTheRight sync.Mutex

// An LogoOnTheRightID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateLogoOnTheRight
type LogoOnTheRightID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LogoOnTheRightInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateLogoOnTheRight
type LogoOnTheRightInput struct {
	// The LogoOnTheRight to submit or modify
	// in: body
	LogoOnTheRight *orm.LogoOnTheRightAPI
}

// UpdateLogoOnTheRight
//
// swagger:route PATCH /logoontherights/{ID} logoontherights updateLogoOnTheRight
//
// # Update a logoontheright
//
// Responses:
// default: genericError
//
//	200: logoontherightDBResponse
func (controller *Controller) UpdateLogoOnTheRight(w http.ResponseWriter, r *http.Request) {

	mutexLogoOnTheRight.Lock()
	defer mutexLogoOnTheRight.Unlock()

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
	db := backRepo.BackRepoLogoOnTheRight.GetDB()

	// Validate input
	var input orm.LogoOnTheRightAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var logoontherightDB orm.LogoOnTheRightDB

	// fetch the logoontheright
	_, err := db.First(&logoontherightDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	logoontherightDB.CopyBasicFieldsFromLogoOnTheRight_WOP(&input.LogoOnTheRight_WOP)
	logoontherightDB.LogoOnTheRightPointersEncoding = input.LogoOnTheRightPointersEncoding

	db, _ = db.Model(&logoontherightDB)
	_, err = db.Updates(&logoontherightDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	logoontherightNew := new(models.LogoOnTheRight)
	logoontherightDB.CopyBasicFieldsToLogoOnTheRight(logoontherightNew)

	// redeem pointers
	logoontherightDB.DecodePointers(backRepo, logoontherightNew)

	// get stage instance from DB instance, and call callback function
	logoontherightOld := backRepo.BackRepoLogoOnTheRight.Map_LogoOnTheRightDBID_LogoOnTheRightPtr[logoontherightDB.ID]
	if logoontherightOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(logoontherightOld, logoontherightNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the logoontherightDB
	writeJSON(w, http.StatusOK, logoontherightDB)
}
