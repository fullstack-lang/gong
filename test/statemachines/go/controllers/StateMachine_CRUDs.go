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
var __StateMachine__dummysDeclaration__ models.StateMachine
var __StateMachine_time__dummyDeclaration time.Duration

var mutexStateMachine sync.Mutex

// An StateMachineID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getStateMachine updateStateMachine deleteStateMachine
type StateMachineID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// StateMachineInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postStateMachine updateStateMachine
type StateMachineInput struct {
	// The StateMachine to submit or modify
	// in: body
	StateMachine *orm.StateMachineAPI
}

// GetStateMachines
//
// swagger:route GET /statemachines statemachines getStateMachines
//
// # Get all statemachines
//
// Responses:
// default: genericError
//
//	200: statemachineDBResponse
func (controller *Controller) GetStateMachines(c *gin.Context) {

	// source slice
	var statemachineDBs []orm.StateMachineDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStateMachines", "Name", stackPath)
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
	db := backRepo.BackRepoStateMachine.GetDB()

	_, err := db.Find(&statemachineDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	statemachineAPIs := make([]orm.StateMachineAPI, 0)

	// for each statemachine, update fields from the database nullable fields
	for idx := range statemachineDBs {
		statemachineDB := &statemachineDBs[idx]
		_ = statemachineDB
		var statemachineAPI orm.StateMachineAPI

		// insertion point for updating fields
		statemachineAPI.ID = statemachineDB.ID
		statemachineDB.CopyBasicFieldsToStateMachine_WOP(&statemachineAPI.StateMachine_WOP)
		statemachineAPI.StateMachinePointersEncoding = statemachineDB.StateMachinePointersEncoding
		statemachineAPIs = append(statemachineAPIs, statemachineAPI)
	}

	c.JSON(http.StatusOK, statemachineAPIs)
}

// PostStateMachine
//
// swagger:route POST /statemachines statemachines postStateMachine
//
// Creates a statemachine
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostStateMachine(c *gin.Context) {

	mutexStateMachine.Lock()
	defer mutexStateMachine.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostStateMachines", "Name", stackPath)
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
	db := backRepo.BackRepoStateMachine.GetDB()

	// Validate input
	var input orm.StateMachineAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create statemachine
	statemachineDB := orm.StateMachineDB{}
	statemachineDB.StateMachinePointersEncoding = input.StateMachinePointersEncoding
	statemachineDB.CopyBasicFieldsFromStateMachine_WOP(&input.StateMachine_WOP)

	_, err = db.Create(&statemachineDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoStateMachine.CheckoutPhaseOneInstance(&statemachineDB)
	statemachine := backRepo.BackRepoStateMachine.Map_StateMachineDBID_StateMachinePtr[statemachineDB.ID]

	if statemachine != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), statemachine)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, statemachineDB)
}

// GetStateMachine
//
// swagger:route GET /statemachines/{ID} statemachines getStateMachine
//
// Gets the details for a statemachine.
//
// Responses:
// default: genericError
//
//	200: statemachineDBResponse
func (controller *Controller) GetStateMachine(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStateMachine", "Name", stackPath)
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
	db := backRepo.BackRepoStateMachine.GetDB()

	// Get statemachineDB in DB
	var statemachineDB orm.StateMachineDB
	if _, err := db.First(&statemachineDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var statemachineAPI orm.StateMachineAPI
	statemachineAPI.ID = statemachineDB.ID
	statemachineAPI.StateMachinePointersEncoding = statemachineDB.StateMachinePointersEncoding
	statemachineDB.CopyBasicFieldsToStateMachine_WOP(&statemachineAPI.StateMachine_WOP)

	c.JSON(http.StatusOK, statemachineAPI)
}

// UpdateStateMachine
//
// swagger:route PATCH /statemachines/{ID} statemachines updateStateMachine
//
// # Update a statemachine
//
// Responses:
// default: genericError
//
//	200: statemachineDBResponse
func (controller *Controller) UpdateStateMachine(c *gin.Context) {

	mutexStateMachine.Lock()
	defer mutexStateMachine.Unlock()

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
	db := backRepo.BackRepoStateMachine.GetDB()

	// Validate input
	var input orm.StateMachineAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var statemachineDB orm.StateMachineDB

	// fetch the statemachine
	_, err := db.First(&statemachineDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	statemachineDB.CopyBasicFieldsFromStateMachine_WOP(&input.StateMachine_WOP)
	statemachineDB.StateMachinePointersEncoding = input.StateMachinePointersEncoding

	db, _ = db.Model(&statemachineDB)
	_, err = db.Updates(&statemachineDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	statemachineNew := new(models.StateMachine)
	statemachineDB.CopyBasicFieldsToStateMachine(statemachineNew)

	// redeem pointers
	statemachineDB.DecodePointers(backRepo, statemachineNew)

	// get stage instance from DB instance, and call callback function
	statemachineOld := backRepo.BackRepoStateMachine.Map_StateMachineDBID_StateMachinePtr[statemachineDB.ID]
	if statemachineOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), statemachineOld, statemachineNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the statemachineDB
	c.JSON(http.StatusOK, statemachineDB)
}

// DeleteStateMachine
//
// swagger:route DELETE /statemachines/{ID} statemachines deleteStateMachine
//
// # Delete a statemachine
//
// default: genericError
//
//	200: statemachineDBResponse
func (controller *Controller) DeleteStateMachine(c *gin.Context) {

	mutexStateMachine.Lock()
	defer mutexStateMachine.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteStateMachine", "Name", stackPath)
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
	db := backRepo.BackRepoStateMachine.GetDB()

	// Get model if exist
	var statemachineDB orm.StateMachineDB
	if _, err := db.First(&statemachineDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&statemachineDB)

	// get an instance (not staged) from DB instance, and call callback function
	statemachineDeleted := new(models.StateMachine)
	statemachineDB.CopyBasicFieldsToStateMachine(statemachineDeleted)

	// get stage instance from DB instance, and call callback function
	statemachineStaged := backRepo.BackRepoStateMachine.Map_StateMachineDBID_StateMachinePtr[statemachineDB.ID]
	if statemachineStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), statemachineStaged, statemachineDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
