// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/sim/go/models"
	"github.com/fullstack-lang/gong/lib/sim/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __DummyAgent__dummysDeclaration__ models.DummyAgent
var __DummyAgent_time__dummyDeclaration time.Duration

var mutexDummyAgent sync.Mutex

// An DummyAgentID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDummyAgent updateDummyAgent deleteDummyAgent
type DummyAgentID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DummyAgentInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDummyAgent updateDummyAgent
type DummyAgentInput struct {
	// The DummyAgent to submit or modify
	// in: body
	DummyAgent *orm.DummyAgentAPI
}

// GetDummyAgents
//
// swagger:route GET /dummyagents dummyagents getDummyAgents
//
// # Get all dummyagents
//
// Responses:
// default: genericError
//
//	200: dummyagentDBResponse
func (controller *Controller) GetDummyAgents(c *gin.Context) {

	// source slice
	var dummyagentDBs []orm.DummyAgentDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDummyAgents", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoDummyAgent.GetDB()

	_, err := db.Find(&dummyagentDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	dummyagentAPIs := make([]orm.DummyAgentAPI, 0)

	// for each dummyagent, update fields from the database nullable fields
	for idx := range dummyagentDBs {
		dummyagentDB := &dummyagentDBs[idx]
		_ = dummyagentDB
		var dummyagentAPI orm.DummyAgentAPI

		// insertion point for updating fields
		dummyagentAPI.ID = dummyagentDB.ID
		dummyagentDB.CopyBasicFieldsToDummyAgent_WOP(&dummyagentAPI.DummyAgent_WOP)
		dummyagentAPI.DummyAgentPointersEncoding = dummyagentDB.DummyAgentPointersEncoding
		dummyagentAPIs = append(dummyagentAPIs, dummyagentAPI)
	}

	c.JSON(http.StatusOK, dummyagentAPIs)
}

// PostDummyAgent
//
// swagger:route POST /dummyagents dummyagents postDummyAgent
//
// Creates a dummyagent
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDummyAgent(c *gin.Context) {

	mutexDummyAgent.Lock()
	defer mutexDummyAgent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDummyAgents", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoDummyAgent.GetDB()

	// Validate input
	var input orm.DummyAgentAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create dummyagent
	dummyagentDB := orm.DummyAgentDB{}
	dummyagentDB.DummyAgentPointersEncoding = input.DummyAgentPointersEncoding
	dummyagentDB.CopyBasicFieldsFromDummyAgent_WOP(&input.DummyAgent_WOP)

	_, err = db.Create(&dummyagentDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDummyAgent.CheckoutPhaseOneInstance(&dummyagentDB)
	dummyagent := backRepo.BackRepoDummyAgent.Map_DummyAgentDBID_DummyAgentPtr[dummyagentDB.ID]

	if dummyagent != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), dummyagent)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, dummyagentDB)
}

// GetDummyAgent
//
// swagger:route GET /dummyagents/{ID} dummyagents getDummyAgent
//
// Gets the details for a dummyagent.
//
// Responses:
// default: genericError
//
//	200: dummyagentDBResponse
func (controller *Controller) GetDummyAgent(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDummyAgent", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoDummyAgent.GetDB()

	// Get dummyagentDB in DB
	var dummyagentDB orm.DummyAgentDB
	if _, err := db.First(&dummyagentDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var dummyagentAPI orm.DummyAgentAPI
	dummyagentAPI.ID = dummyagentDB.ID
	dummyagentAPI.DummyAgentPointersEncoding = dummyagentDB.DummyAgentPointersEncoding
	dummyagentDB.CopyBasicFieldsToDummyAgent_WOP(&dummyagentAPI.DummyAgent_WOP)

	c.JSON(http.StatusOK, dummyagentAPI)
}

// UpdateDummyAgent
//
// swagger:route PATCH /dummyagents/{ID} dummyagents updateDummyAgent
//
// # Update a dummyagent
//
// Responses:
// default: genericError
//
//	200: dummyagentDBResponse
func (controller *Controller) UpdateDummyAgent(c *gin.Context) {

	mutexDummyAgent.Lock()
	defer mutexDummyAgent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDummyAgent", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoDummyAgent.GetDB()

	// Validate input
	var input orm.DummyAgentAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var dummyagentDB orm.DummyAgentDB

	// fetch the dummyagent
	_, err := db.First(&dummyagentDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	dummyagentDB.CopyBasicFieldsFromDummyAgent_WOP(&input.DummyAgent_WOP)
	dummyagentDB.DummyAgentPointersEncoding = input.DummyAgentPointersEncoding

	db, _ = db.Model(&dummyagentDB)
	_, err = db.Updates(&dummyagentDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	dummyagentNew := new(models.DummyAgent)
	dummyagentDB.CopyBasicFieldsToDummyAgent(dummyagentNew)

	// redeem pointers
	dummyagentDB.DecodePointers(backRepo, dummyagentNew)

	// get stage instance from DB instance, and call callback function
	dummyagentOld := backRepo.BackRepoDummyAgent.Map_DummyAgentDBID_DummyAgentPtr[dummyagentDB.ID]
	if dummyagentOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), dummyagentOld, dummyagentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the dummyagentDB
	c.JSON(http.StatusOK, dummyagentDB)
}

// DeleteDummyAgent
//
// swagger:route DELETE /dummyagents/{ID} dummyagents deleteDummyAgent
//
// # Delete a dummyagent
//
// default: genericError
//
//	200: dummyagentDBResponse
func (controller *Controller) DeleteDummyAgent(c *gin.Context) {

	mutexDummyAgent.Lock()
	defer mutexDummyAgent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDummyAgent", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoDummyAgent.GetDB()

	// Get model if exist
	var dummyagentDB orm.DummyAgentDB
	if _, err := db.First(&dummyagentDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&dummyagentDB)

	// get an instance (not staged) from DB instance, and call callback function
	dummyagentDeleted := new(models.DummyAgent)
	dummyagentDB.CopyBasicFieldsToDummyAgent(dummyagentDeleted)

	// get stage instance from DB instance, and call callback function
	dummyagentStaged := backRepo.BackRepoDummyAgent.Map_DummyAgentDBID_DummyAgentPtr[dummyagentDB.ID]
	if dummyagentStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), dummyagentStaged, dummyagentDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
