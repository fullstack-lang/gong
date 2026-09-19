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
var __LogoOnTheLeft__dummysDeclaration__ models.LogoOnTheLeft
var _ = __LogoOnTheLeft__dummysDeclaration__
var __LogoOnTheLeft_time__dummyDeclaration time.Duration
var _ = __LogoOnTheLeft_time__dummyDeclaration

var mutexLogoOnTheLeft sync.Mutex

// An LogoOnTheLeftID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateLogoOnTheLeft
type LogoOnTheLeftID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LogoOnTheLeftInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateLogoOnTheLeft
type LogoOnTheLeftInput struct {
	// The LogoOnTheLeft to submit or modify
	// in: body
	LogoOnTheLeft *orm.LogoOnTheLeftAPI
}

// UpdateLogoOnTheLeft
//
// swagger:route PATCH /logoonthelefts/{ID} logoonthelefts updateLogoOnTheLeft
//
// # Update a logoontheleft
//
// Responses:
// default: genericError
//
//	200: logoontheleftDBResponse
func (controller *Controller) UpdateLogoOnTheLeft(w http.ResponseWriter, r *http.Request) {

	mutexLogoOnTheLeft.Lock()
	defer mutexLogoOnTheLeft.Unlock()

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
	db := backRepo.BackRepoLogoOnTheLeft.GetDB()

	// Validate input
	var input orm.LogoOnTheLeftAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var logoontheleftDB orm.LogoOnTheLeftDB

	// fetch the logoontheleft
	_, err := db.First(&logoontheleftDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	logoontheleftDB.CopyBasicFieldsFromLogoOnTheLeft_WOP(&input.LogoOnTheLeft_WOP)
	logoontheleftDB.LogoOnTheLeftPointersEncoding = input.LogoOnTheLeftPointersEncoding

	db, _ = db.Model(&logoontheleftDB)
	_, err = db.Updates(&logoontheleftDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	logoontheleftNew := new(models.LogoOnTheLeft)
	logoontheleftDB.CopyBasicFieldsToLogoOnTheLeft(logoontheleftNew)

	// redeem pointers
	logoontheleftDB.DecodePointers(backRepo, logoontheleftNew)

	// get stage instance from DB instance, and call callback function
	logoontheleftOld := backRepo.BackRepoLogoOnTheLeft.Map_LogoOnTheLeftDBID_LogoOnTheLeftPtr[logoontheleftDB.ID]
	if logoontheleftOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(logoontheleftOld, logoontheleftNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the logoontheleftDB
	writeJSON(w, http.StatusOK, logoontheleftDB)
}
