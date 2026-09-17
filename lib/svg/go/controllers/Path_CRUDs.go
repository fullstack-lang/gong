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
var __Path__dummysDeclaration__ models.Path
var _ = __Path__dummysDeclaration__
var __Path_time__dummyDeclaration time.Duration
var _ = __Path_time__dummyDeclaration

var mutexPath sync.Mutex

// An PathID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updatePath
type PathID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PathInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updatePath
type PathInput struct {
	// The Path to submit or modify
	// in: body
	Path *orm.PathAPI
}

// UpdatePath
//
// swagger:route PATCH /paths/{ID} paths updatePath
//
// # Update a path
//
// Responses:
// default: genericError
//
//	200: pathDBResponse
func (controller *Controller) UpdatePath(w http.ResponseWriter, r *http.Request) {

	mutexPath.Lock()
	defer mutexPath.Unlock()

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
	db := backRepo.BackRepoPath.GetDB()

	// Validate input
	var input orm.PathAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var pathDB orm.PathDB

	// fetch the path
	_, err := db.First(&pathDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	pathDB.CopyBasicFieldsFromPath_WOP(&input.Path_WOP)
	pathDB.PathPointersEncoding = input.PathPointersEncoding

	db, _ = db.Model(&pathDB)
	_, err = db.Updates(&pathDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	pathNew := new(models.Path)
	pathDB.CopyBasicFieldsToPath(pathNew)

	// redeem pointers
	pathDB.DecodePointers(backRepo, pathNew)

	// get stage instance from DB instance, and call callback function
	pathOld := backRepo.BackRepoPath.Map_PathDBID_PathPtr[pathDB.ID]
	if pathOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(pathOld, pathNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the pathDB
	writeJSON(w, http.StatusOK, pathDB)
}
