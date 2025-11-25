// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/test/statemachines/go/models"
	"github.com/fullstack-lang/gong/test/statemachines/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __State__dummysDeclaration__ models.State
var __State_time__dummyDeclaration time.Duration

var mutexState sync.Mutex

// An StateID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getState updateState deleteState
type StateID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// StateInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postState updateState
type StateInput struct {
	// The State to submit or modify
	// in: body
	State *orm.StateAPI
}

// GetStates
//
// swagger:route GET /states states getStates
//
// # Get all states
//
// Responses:
// default: genericError
//
//	200: stateDBResponse
func (controller *Controller) GetStates(c *gin.Context) {

	// source slice
	var stateDBs []orm.StateDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStates", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoState.GetDB()

	_, err := db.Find(&stateDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	stateAPIs := make([]orm.StateAPI, 0)

	// for each state, update fields from the database nullable fields
	for idx := range stateDBs {
		stateDB := &stateDBs[idx]
		_ = stateDB
		var stateAPI orm.StateAPI

		// insertion point for updating fields
		stateAPI.ID = stateDB.ID
		stateDB.CopyBasicFieldsToState_WOP(&stateAPI.State_WOP)
		stateAPI.StatePointersEncoding = stateDB.StatePointersEncoding
		stateAPIs = append(stateAPIs, stateAPI)
	}

	c.JSON(http.StatusOK, stateAPIs)
}

// PostState
//
// swagger:route POST /states states postState
//
// Creates a state
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostState(c *gin.Context) {

	mutexState.Lock()
	defer mutexState.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostStates", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoState.GetDB()

	// Validate input
	var input orm.StateAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create state
	stateDB := orm.StateDB{}
	stateDB.StatePointersEncoding = input.StatePointersEncoding
	stateDB.CopyBasicFieldsFromState_WOP(&input.State_WOP)

	_, err = db.Create(&stateDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoState.CheckoutPhaseOneInstance(&stateDB)
	state := backRepo.BackRepoState.Map_StateDBID_StatePtr[stateDB.ID]

	if state != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), state)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, stateDB)
}

// GetState
//
// swagger:route GET /states/{ID} states getState
//
// Gets the details for a state.
//
// Responses:
// default: genericError
//
//	200: stateDBResponse
func (controller *Controller) GetState(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetState", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoState.GetDB()

	// Get stateDB in DB
	var stateDB orm.StateDB
	if _, err := db.First(&stateDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var stateAPI orm.StateAPI
	stateAPI.ID = stateDB.ID
	stateAPI.StatePointersEncoding = stateDB.StatePointersEncoding
	stateDB.CopyBasicFieldsToState_WOP(&stateAPI.State_WOP)

	c.JSON(http.StatusOK, stateAPI)
}

// UpdateState
//
// swagger:route PATCH /states/{ID} states updateState
//
// # Update a state
//
// Responses:
// default: genericError
//
//	200: stateDBResponse
func (controller *Controller) UpdateState(c *gin.Context) {

	mutexState.Lock()
	defer mutexState.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) >= 1 {
		_nameValues := _values["Name"]
		if len(_nameValues) == 1 {
			stackPath = _nameValues[0]
		}
	}

	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoState.GetDB()

	// Validate input
	var input orm.StateAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var stateDB orm.StateDB

	// fetch the state
	_, err := db.First(&stateDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	stateDB.CopyBasicFieldsFromState_WOP(&input.State_WOP)
	stateDB.StatePointersEncoding = input.StatePointersEncoding

	db, _ = db.Model(&stateDB)
	_, err = db.Updates(&stateDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	stateNew := new(models.State)
	stateDB.CopyBasicFieldsToState(stateNew)

	// redeem pointers
	stateDB.DecodePointers(backRepo, stateNew)

	// get stage instance from DB instance, and call callback function
	stateOld := backRepo.BackRepoState.Map_StateDBID_StatePtr[stateDB.ID]
	if stateOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), stateOld, stateNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the stateDB
	c.JSON(http.StatusOK, stateDB)
}

// DeleteState
//
// swagger:route DELETE /states/{ID} states deleteState
//
// # Delete a state
//
// default: genericError
//
//	200: stateDBResponse
func (controller *Controller) DeleteState(c *gin.Context) {

	mutexState.Lock()
	defer mutexState.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteState", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoState.GetDB()

	// Get model if exist
	var stateDB orm.StateDB
	if _, err := db.First(&stateDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&stateDB)

	// get an instance (not staged) from DB instance, and call callback function
	stateDeleted := new(models.State)
	stateDB.CopyBasicFieldsToState(stateDeleted)

	// get stage instance from DB instance, and call callback function
	stateStaged := backRepo.BackRepoState.Map_StateDBID_StatePtr[stateDB.ID]
	if stateStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), stateStaged, stateDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
