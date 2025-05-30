// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/load/go/models"
	"github.com/fullstack-lang/gong/lib/load/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Message__dummysDeclaration__ models.Message
var __Message_time__dummyDeclaration time.Duration

var mutexMessage sync.Mutex

// An MessageID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMessage updateMessage deleteMessage
type MessageID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MessageInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMessage updateMessage
type MessageInput struct {
	// The Message to submit or modify
	// in: body
	Message *orm.MessageAPI
}

// GetMessages
//
// swagger:route GET /messages messages getMessages
//
// # Get all messages
//
// Responses:
// default: genericError
//
//	200: messageDBResponse
func (controller *Controller) GetMessages(c *gin.Context) {

	// source slice
	var messageDBs []orm.MessageDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMessages", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/load/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMessage.GetDB()

	_, err := db.Find(&messageDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	messageAPIs := make([]orm.MessageAPI, 0)

	// for each message, update fields from the database nullable fields
	for idx := range messageDBs {
		messageDB := &messageDBs[idx]
		_ = messageDB
		var messageAPI orm.MessageAPI

		// insertion point for updating fields
		messageAPI.ID = messageDB.ID
		messageDB.CopyBasicFieldsToMessage_WOP(&messageAPI.Message_WOP)
		messageAPI.MessagePointersEncoding = messageDB.MessagePointersEncoding
		messageAPIs = append(messageAPIs, messageAPI)
	}

	c.JSON(http.StatusOK, messageAPIs)
}

// PostMessage
//
// swagger:route POST /messages messages postMessage
//
// Creates a message
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMessage(c *gin.Context) {

	mutexMessage.Lock()
	defer mutexMessage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMessages", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/load/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMessage.GetDB()

	// Validate input
	var input orm.MessageAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create message
	messageDB := orm.MessageDB{}
	messageDB.MessagePointersEncoding = input.MessagePointersEncoding
	messageDB.CopyBasicFieldsFromMessage_WOP(&input.Message_WOP)

	_, err = db.Create(&messageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMessage.CheckoutPhaseOneInstance(&messageDB)
	message := backRepo.BackRepoMessage.Map_MessageDBID_MessagePtr[messageDB.ID]

	if message != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), message)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, messageDB)
}

// GetMessage
//
// swagger:route GET /messages/{ID} messages getMessage
//
// Gets the details for a message.
//
// Responses:
// default: genericError
//
//	200: messageDBResponse
func (controller *Controller) GetMessage(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMessage", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/load/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMessage.GetDB()

	// Get messageDB in DB
	var messageDB orm.MessageDB
	if _, err := db.First(&messageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var messageAPI orm.MessageAPI
	messageAPI.ID = messageDB.ID
	messageAPI.MessagePointersEncoding = messageDB.MessagePointersEncoding
	messageDB.CopyBasicFieldsToMessage_WOP(&messageAPI.Message_WOP)

	c.JSON(http.StatusOK, messageAPI)
}

// UpdateMessage
//
// swagger:route PATCH /messages/{ID} messages updateMessage
//
// # Update a message
//
// Responses:
// default: genericError
//
//	200: messageDBResponse
func (controller *Controller) UpdateMessage(c *gin.Context) {

	mutexMessage.Lock()
	defer mutexMessage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMessage", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/load/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMessage.GetDB()

	// Validate input
	var input orm.MessageAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var messageDB orm.MessageDB

	// fetch the message
	_, err := db.First(&messageDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	messageDB.CopyBasicFieldsFromMessage_WOP(&input.Message_WOP)
	messageDB.MessagePointersEncoding = input.MessagePointersEncoding

	db, _ = db.Model(&messageDB)
	_, err = db.Updates(&messageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	messageNew := new(models.Message)
	messageDB.CopyBasicFieldsToMessage(messageNew)

	// redeem pointers
	messageDB.DecodePointers(backRepo, messageNew)

	// get stage instance from DB instance, and call callback function
	messageOld := backRepo.BackRepoMessage.Map_MessageDBID_MessagePtr[messageDB.ID]
	if messageOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), messageOld, messageNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the messageDB
	c.JSON(http.StatusOK, messageDB)
}

// DeleteMessage
//
// swagger:route DELETE /messages/{ID} messages deleteMessage
//
// # Delete a message
//
// default: genericError
//
//	200: messageDBResponse
func (controller *Controller) DeleteMessage(c *gin.Context) {

	mutexMessage.Lock()
	defer mutexMessage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMessage", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/load/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMessage.GetDB()

	// Get model if exist
	var messageDB orm.MessageDB
	if _, err := db.First(&messageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&messageDB)

	// get an instance (not staged) from DB instance, and call callback function
	messageDeleted := new(models.Message)
	messageDB.CopyBasicFieldsToMessage(messageDeleted)

	// get stage instance from DB instance, and call callback function
	messageStaged := backRepo.BackRepoMessage.Map_MessageDBID_MessagePtr[messageDB.ID]
	if messageStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), messageStaged, messageDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
