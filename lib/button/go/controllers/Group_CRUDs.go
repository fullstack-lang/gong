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
// swagger:parameters getGroup updateGroup deleteGroup
type GroupID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GroupInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGroup updateGroup
type GroupInput struct {
	// The Group to submit or modify
	// in: body
	Group *orm.GroupAPI
}

// GetGroups
//
// swagger:route GET /groups groups getGroups
//
// # Get all groups
//
// Responses:
// default: genericError
//
//	200: groupDBResponse
func (controller *Controller) GetGroups(w http.ResponseWriter, r *http.Request) {

	// source slice
	var groupDBs []orm.GroupDB

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGroups", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoGroup.GetDB()

	_, err := db.Find(&groupDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	groupAPIs := make([]orm.GroupAPI, 0)

	// for each group, update fields from the database nullable fields
	for idx := range groupDBs {
		groupDB := &groupDBs[idx]
		_ = groupDB
		var groupAPI orm.GroupAPI

		// insertion point for updating fields
		groupAPI.ID = groupDB.ID
		groupDB.CopyBasicFieldsToGroup_WOP(&groupAPI.Group_WOP)
		groupAPI.GroupPointersEncoding = groupDB.GroupPointersEncoding
		groupAPIs = append(groupAPIs, groupAPI)
	}

	writeJSON(w, http.StatusOK, groupAPIs)
}

// PostGroup
//
// swagger:route POST /groups groups postGroup
//
// Creates a group
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGroup(w http.ResponseWriter, r *http.Request) {

	mutexGroup.Lock()
	defer mutexGroup.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGroups", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoGroup.GetDB()

	// Validate input
	var input orm.GroupAPI

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// Create group
	groupDB := orm.GroupDB{}
	groupDB.GroupPointersEncoding = input.GroupPointersEncoding
	groupDB.CopyBasicFieldsFromGroup_WOP(&input.Group_WOP)

	_, err = db.Create(&groupDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGroup.CheckoutPhaseOneInstance(&groupDB)
	group := backRepo.BackRepoGroup.Map_GroupDBID_GroupPtr[groupDB.ID]

	if group != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), group)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, groupDB)
}

// GetGroup
//
// swagger:route GET /groups/{ID} groups getGroup
//
// Gets the details for a group.
//
// Responses:
// default: genericError
//
//	200: groupDBResponse
func (controller *Controller) GetGroup(w http.ResponseWriter, r *http.Request) {

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGroup", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoGroup.GetDB()

	// Get groupDB in DB
	var groupDB orm.GroupDB
	if _, err := db.First(&groupDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	var groupAPI orm.GroupAPI
	groupAPI.ID = groupDB.ID
	groupAPI.GroupPointersEncoding = groupDB.GroupPointersEncoding
	groupDB.CopyBasicFieldsToGroup_WOP(&groupAPI.Group_WOP)

	writeJSON(w, http.StatusOK, groupAPI)
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
		models.OnAfterUpdateFromFront(backRepo.GetStage(), groupOld, groupNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the groupDB
	writeJSON(w, http.StatusOK, groupDB)
}

// DeleteGroup
//
// swagger:route DELETE /groups/{ID} groups deleteGroup
//
// # Delete a group
//
// default: genericError
//
//	200: groupDBResponse
func (controller *Controller) DeleteGroup(w http.ResponseWriter, r *http.Request) {

	mutexGroup.Lock()
	defer mutexGroup.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGroup", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoGroup.GetDB()

	// Get model if exist
	var groupDB orm.GroupDB
	if _, err := db.First(&groupDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&groupDB)

	// get an instance (not staged) from DB instance, and call callback function
	groupDeleted := new(models.Group)
	groupDB.CopyBasicFieldsToGroup(groupDeleted)

	// get stage instance from DB instance, and call callback function
	groupStaged := backRepo.BackRepoGroup.Map_GroupDBID_GroupPtr[groupDB.ID]
	if groupStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), groupStaged, groupDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, H{"data": true})
}
