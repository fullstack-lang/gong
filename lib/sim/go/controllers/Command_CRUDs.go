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
var __Command__dummysDeclaration__ models.Command
var __Command_time__dummyDeclaration time.Duration

var mutexCommand sync.Mutex

// An CommandID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCommand updateCommand deleteCommand
type CommandID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CommandInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCommand updateCommand
type CommandInput struct {
	// The Command to submit or modify
	// in: body
	Command *orm.CommandAPI
}

// GetCommands
//
// swagger:route GET /commands commands getCommands
//
// # Get all commands
//
// Responses:
// default: genericError
//
//	200: commandDBResponse
func (controller *Controller) GetCommands(c *gin.Context) {

	// source slice
	var commandDBs []orm.CommandDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCommands", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoCommand.GetDB()

	_, err := db.Find(&commandDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	commandAPIs := make([]orm.CommandAPI, 0)

	// for each command, update fields from the database nullable fields
	for idx := range commandDBs {
		commandDB := &commandDBs[idx]
		_ = commandDB
		var commandAPI orm.CommandAPI

		// insertion point for updating fields
		commandAPI.ID = commandDB.ID
		commandDB.CopyBasicFieldsToCommand_WOP(&commandAPI.Command_WOP)
		commandAPI.CommandPointersEncoding = commandDB.CommandPointersEncoding
		commandAPIs = append(commandAPIs, commandAPI)
	}

	c.JSON(http.StatusOK, commandAPIs)
}

// PostCommand
//
// swagger:route POST /commands commands postCommand
//
// Creates a command
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCommand(c *gin.Context) {

	mutexCommand.Lock()
	defer mutexCommand.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCommands", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoCommand.GetDB()

	// Validate input
	var input orm.CommandAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create command
	commandDB := orm.CommandDB{}
	commandDB.CommandPointersEncoding = input.CommandPointersEncoding
	commandDB.CopyBasicFieldsFromCommand_WOP(&input.Command_WOP)

	_, err = db.Create(&commandDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCommand.CheckoutPhaseOneInstance(&commandDB)
	command := backRepo.BackRepoCommand.Map_CommandDBID_CommandPtr[commandDB.ID]

	if command != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), command)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, commandDB)
}

// GetCommand
//
// swagger:route GET /commands/{ID} commands getCommand
//
// Gets the details for a command.
//
// Responses:
// default: genericError
//
//	200: commandDBResponse
func (controller *Controller) GetCommand(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCommand", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoCommand.GetDB()

	// Get commandDB in DB
	var commandDB orm.CommandDB
	if _, err := db.First(&commandDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var commandAPI orm.CommandAPI
	commandAPI.ID = commandDB.ID
	commandAPI.CommandPointersEncoding = commandDB.CommandPointersEncoding
	commandDB.CopyBasicFieldsToCommand_WOP(&commandAPI.Command_WOP)

	c.JSON(http.StatusOK, commandAPI)
}

// UpdateCommand
//
// swagger:route PATCH /commands/{ID} commands updateCommand
//
// # Update a command
//
// Responses:
// default: genericError
//
//	200: commandDBResponse
func (controller *Controller) UpdateCommand(c *gin.Context) {

	mutexCommand.Lock()
	defer mutexCommand.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCommand", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoCommand.GetDB()

	// Validate input
	var input orm.CommandAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var commandDB orm.CommandDB

	// fetch the command
	_, err := db.First(&commandDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	commandDB.CopyBasicFieldsFromCommand_WOP(&input.Command_WOP)
	commandDB.CommandPointersEncoding = input.CommandPointersEncoding

	db, _ = db.Model(&commandDB)
	_, err = db.Updates(&commandDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	commandNew := new(models.Command)
	commandDB.CopyBasicFieldsToCommand(commandNew)

	// redeem pointers
	commandDB.DecodePointers(backRepo, commandNew)

	// get stage instance from DB instance, and call callback function
	commandOld := backRepo.BackRepoCommand.Map_CommandDBID_CommandPtr[commandDB.ID]
	if commandOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), commandOld, commandNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the commandDB
	c.JSON(http.StatusOK, commandDB)
}

// DeleteCommand
//
// swagger:route DELETE /commands/{ID} commands deleteCommand
//
// # Delete a command
//
// default: genericError
//
//	200: commandDBResponse
func (controller *Controller) DeleteCommand(c *gin.Context) {

	mutexCommand.Lock()
	defer mutexCommand.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCommand", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoCommand.GetDB()

	// Get model if exist
	var commandDB orm.CommandDB
	if _, err := db.First(&commandDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&commandDB)

	// get an instance (not staged) from DB instance, and call callback function
	commandDeleted := new(models.Command)
	commandDB.CopyBasicFieldsToCommand(commandDeleted)

	// get stage instance from DB instance, and call callback function
	commandStaged := backRepo.BackRepoCommand.Map_CommandDBID_CommandPtr[commandDB.ID]
	if commandStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), commandStaged, commandDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
