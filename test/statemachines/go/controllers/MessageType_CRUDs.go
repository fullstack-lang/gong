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
var __MessageType__dummysDeclaration__ models.MessageType
var __MessageType_time__dummyDeclaration time.Duration

var mutexMessageType sync.Mutex

// An MessageTypeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMessageType updateMessageType deleteMessageType
type MessageTypeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MessageTypeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMessageType updateMessageType
type MessageTypeInput struct {
	// The MessageType to submit or modify
	// in: body
	MessageType *orm.MessageTypeAPI
}

// GetMessageTypes
//
// swagger:route GET /messagetypes messagetypes getMessageTypes
//
// # Get all messagetypes
//
// Responses:
// default: genericError
//
//	200: messagetypeDBResponse
func (controller *Controller) GetMessageTypes(c *gin.Context) {

	// source slice
	var messagetypeDBs []orm.MessageTypeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMessageTypes", "Name", stackPath)
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
	db := backRepo.BackRepoMessageType.GetDB()

	_, err := db.Find(&messagetypeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	messagetypeAPIs := make([]orm.MessageTypeAPI, 0)

	// for each messagetype, update fields from the database nullable fields
	for idx := range messagetypeDBs {
		messagetypeDB := &messagetypeDBs[idx]
		_ = messagetypeDB
		var messagetypeAPI orm.MessageTypeAPI

		// insertion point for updating fields
		messagetypeAPI.ID = messagetypeDB.ID
		messagetypeDB.CopyBasicFieldsToMessageType_WOP(&messagetypeAPI.MessageType_WOP)
		messagetypeAPI.MessageTypePointersEncoding = messagetypeDB.MessageTypePointersEncoding
		messagetypeAPIs = append(messagetypeAPIs, messagetypeAPI)
	}

	c.JSON(http.StatusOK, messagetypeAPIs)
}

// PostMessageType
//
// swagger:route POST /messagetypes messagetypes postMessageType
//
// Creates a messagetype
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMessageType(c *gin.Context) {

	mutexMessageType.Lock()
	defer mutexMessageType.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMessageTypes", "Name", stackPath)
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
	db := backRepo.BackRepoMessageType.GetDB()

	// Validate input
	var input orm.MessageTypeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create messagetype
	messagetypeDB := orm.MessageTypeDB{}
	messagetypeDB.MessageTypePointersEncoding = input.MessageTypePointersEncoding
	messagetypeDB.CopyBasicFieldsFromMessageType_WOP(&input.MessageType_WOP)

	_, err = db.Create(&messagetypeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMessageType.CheckoutPhaseOneInstance(&messagetypeDB)
	messagetype := backRepo.BackRepoMessageType.Map_MessageTypeDBID_MessageTypePtr[messagetypeDB.ID]

	if messagetype != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), messagetype)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, messagetypeDB)
}

// GetMessageType
//
// swagger:route GET /messagetypes/{ID} messagetypes getMessageType
//
// Gets the details for a messagetype.
//
// Responses:
// default: genericError
//
//	200: messagetypeDBResponse
func (controller *Controller) GetMessageType(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMessageType", "Name", stackPath)
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
	db := backRepo.BackRepoMessageType.GetDB()

	// Get messagetypeDB in DB
	var messagetypeDB orm.MessageTypeDB
	if _, err := db.First(&messagetypeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var messagetypeAPI orm.MessageTypeAPI
	messagetypeAPI.ID = messagetypeDB.ID
	messagetypeAPI.MessageTypePointersEncoding = messagetypeDB.MessageTypePointersEncoding
	messagetypeDB.CopyBasicFieldsToMessageType_WOP(&messagetypeAPI.MessageType_WOP)

	c.JSON(http.StatusOK, messagetypeAPI)
}

// UpdateMessageType
//
// swagger:route PATCH /messagetypes/{ID} messagetypes updateMessageType
//
// # Update a messagetype
//
// Responses:
// default: genericError
//
//	200: messagetypeDBResponse
func (controller *Controller) UpdateMessageType(c *gin.Context) {

	mutexMessageType.Lock()
	defer mutexMessageType.Unlock()

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
	db := backRepo.BackRepoMessageType.GetDB()

	// Validate input
	var input orm.MessageTypeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var messagetypeDB orm.MessageTypeDB

	// fetch the messagetype
	_, err := db.First(&messagetypeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	messagetypeDB.CopyBasicFieldsFromMessageType_WOP(&input.MessageType_WOP)
	messagetypeDB.MessageTypePointersEncoding = input.MessageTypePointersEncoding

	db, _ = db.Model(&messagetypeDB)
	_, err = db.Updates(&messagetypeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	messagetypeNew := new(models.MessageType)
	messagetypeDB.CopyBasicFieldsToMessageType(messagetypeNew)

	// redeem pointers
	messagetypeDB.DecodePointers(backRepo, messagetypeNew)

	// get stage instance from DB instance, and call callback function
	messagetypeOld := backRepo.BackRepoMessageType.Map_MessageTypeDBID_MessageTypePtr[messagetypeDB.ID]
	if messagetypeOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), messagetypeOld, messagetypeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the messagetypeDB
	c.JSON(http.StatusOK, messagetypeDB)
}

// DeleteMessageType
//
// swagger:route DELETE /messagetypes/{ID} messagetypes deleteMessageType
//
// # Delete a messagetype
//
// default: genericError
//
//	200: messagetypeDBResponse
func (controller *Controller) DeleteMessageType(c *gin.Context) {

	mutexMessageType.Lock()
	defer mutexMessageType.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMessageType", "Name", stackPath)
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
	db := backRepo.BackRepoMessageType.GetDB()

	// Get model if exist
	var messagetypeDB orm.MessageTypeDB
	if _, err := db.First(&messagetypeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&messagetypeDB)

	// get an instance (not staged) from DB instance, and call callback function
	messagetypeDeleted := new(models.MessageType)
	messagetypeDB.CopyBasicFieldsToMessageType(messagetypeDeleted)

	// get stage instance from DB instance, and call callback function
	messagetypeStaged := backRepo.BackRepoMessageType.Map_MessageTypeDBID_MessageTypePtr[messagetypeDB.ID]
	if messagetypeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), messagetypeStaged, messagetypeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
