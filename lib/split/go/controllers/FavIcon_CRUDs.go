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
var __FavIcon__dummysDeclaration__ models.FavIcon
var _ = __FavIcon__dummysDeclaration__
var __FavIcon_time__dummyDeclaration time.Duration
var _ = __FavIcon_time__dummyDeclaration

var mutexFavIcon sync.Mutex

// An FavIconID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateFavIcon
type FavIconID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FavIconInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateFavIcon
type FavIconInput struct {
	// The FavIcon to submit or modify
	// in: body
	FavIcon *orm.FavIconAPI
}

// UpdateFavIcon
//
// swagger:route PATCH /favicons/{ID} favicons updateFavIcon
//
// # Update a favicon
//
// Responses:
// default: genericError
//
//	200: faviconDBResponse
func (controller *Controller) UpdateFavIcon(w http.ResponseWriter, r *http.Request) {

	mutexFavIcon.Lock()
	defer mutexFavIcon.Unlock()

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
	db := backRepo.BackRepoFavIcon.GetDB()

	// Validate input
	var input orm.FavIconAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var faviconDB orm.FavIconDB

	// fetch the favicon
	_, err := db.First(&faviconDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	faviconDB.CopyBasicFieldsFromFavIcon_WOP(&input.FavIcon_WOP)
	faviconDB.FavIconPointersEncoding = input.FavIconPointersEncoding

	db, _ = db.Model(&faviconDB)
	_, err = db.Updates(&faviconDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	faviconNew := new(models.FavIcon)
	faviconDB.CopyBasicFieldsToFavIcon(faviconNew)

	// redeem pointers
	faviconDB.DecodePointers(backRepo, faviconNew)

	// get stage instance from DB instance, and call callback function
	faviconOld := backRepo.BackRepoFavIcon.Map_FavIconDBID_FavIconPtr[faviconDB.ID]
	if faviconOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(faviconOld, faviconNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the faviconDB
	writeJSON(w, http.StatusOK, faviconDB)
}
