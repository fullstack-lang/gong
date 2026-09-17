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
var __Group__dummysDeclaration__ models.Group
var _ = __Group__dummysDeclaration__
var __Group_time__dummyDeclaration time.Duration
var _ = __Group_time__dummyDeclaration

var mutexGroup sync.Mutex

// An GroupID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateGroup
type GroupID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GroupInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateGroup
type GroupInput struct {
	// The Group to submit or modify
	// in: body
	Group *orm.GroupAPI
}

// UpdateGroup
//
// swagger:route PATCH /groups/{ID} groups updateGroup
//
// # Update a group
//
// Responses:
// default: genericError
//
//	200: groupDBResponse
func (controller *Controller) UpdateGroup(w http.ResponseWriter, r *http.Request) {

	mutexGroup.Lock()
	defer mutexGroup.Unlock()

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
	db := backRepo.BackRepoGroup.GetDB()

	// Validate input
	var input orm.GroupAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var groupDB orm.GroupDB

	// fetch the group
	_, err := db.First(&groupDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	groupDB.CopyBasicFieldsFromGroup_WOP(&input.Group_WOP)
	groupDB.GroupPointersEncoding = input.GroupPointersEncoding

	db, _ = db.Model(&groupDB)
	_, err = db.Updates(&groupDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	groupNew := new(models.Group)
	groupDB.CopyBasicFieldsToGroup(groupNew)

	// redeem pointers
	groupDB.DecodePointers(backRepo, groupNew)

	// get stage instance from DB instance, and call callback function
	groupOld := backRepo.BackRepoGroup.Map_GroupDBID_GroupPtr[groupDB.ID]
	if groupOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(groupOld, groupNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the groupDB
	writeJSON(w, http.StatusOK, groupDB)
}
