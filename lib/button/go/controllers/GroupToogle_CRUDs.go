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
var __GroupToogle__dummysDeclaration__ models.GroupToogle
var _ = __GroupToogle__dummysDeclaration__
var __GroupToogle_time__dummyDeclaration time.Duration
var _ = __GroupToogle_time__dummyDeclaration

var mutexGroupToogle sync.Mutex

// An GroupToogleID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateGroupToogle
type GroupToogleID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GroupToogleInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateGroupToogle
type GroupToogleInput struct {
	// The GroupToogle to submit or modify
	// in: body
	GroupToogle *orm.GroupToogleAPI
}

// UpdateGroupToogle
//
// swagger:route PATCH /grouptoogles/{ID} grouptoogles updateGroupToogle
//
// # Update a grouptoogle
//
// Responses:
// default: genericError
//
//	200: grouptoogleDBResponse
func (controller *Controller) UpdateGroupToogle(w http.ResponseWriter, r *http.Request) {

	mutexGroupToogle.Lock()
	defer mutexGroupToogle.Unlock()

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
	db := backRepo.BackRepoGroupToogle.GetDB()

	// Validate input
	var input orm.GroupToogleAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var grouptoogleDB orm.GroupToogleDB

	// fetch the grouptoogle
	_, err := db.First(&grouptoogleDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	grouptoogleDB.CopyBasicFieldsFromGroupToogle_WOP(&input.GroupToogle_WOP)
	grouptoogleDB.GroupTooglePointersEncoding = input.GroupTooglePointersEncoding

	db, _ = db.Model(&grouptoogleDB)
	_, err = db.Updates(&grouptoogleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	grouptoogleNew := new(models.GroupToogle)
	grouptoogleDB.CopyBasicFieldsToGroupToogle(grouptoogleNew)

	// redeem pointers
	grouptoogleDB.DecodePointers(backRepo, grouptoogleNew)

	// get stage instance from DB instance, and call callback function
	grouptoogleOld := backRepo.BackRepoGroupToogle.Map_GroupToogleDBID_GroupTooglePtr[grouptoogleDB.ID]
	if grouptoogleOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(grouptoogleOld, grouptoogleNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the grouptoogleDB
	writeJSON(w, http.StatusOK, grouptoogleDB)
}
