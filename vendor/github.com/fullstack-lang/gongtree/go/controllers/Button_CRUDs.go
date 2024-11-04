// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongtree/go/models"
	"github.com/fullstack-lang/gongtree/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Button__dummysDeclaration__ models.Button
var __Button_time__dummyDeclaration time.Duration

var mutexButton sync.Mutex

// An ButtonID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getButton updateButton deleteButton
type ButtonID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ButtonInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postButton updateButton
type ButtonInput struct {
	// The Button to submit or modify
	// in: body
	Button *orm.ButtonAPI
}

// GetButtons
//
// swagger:route GET /buttons buttons getButtons
//
// # Get all buttons
//
// Responses:
// default: genericError
//
//	200: buttonDBResponse
func (controller *Controller) GetButtons(c *gin.Context) {

	// source slice
	var buttonDBs []orm.ButtonDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetButtons", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoButton.GetDB()

	_, err := db.Find(&buttonDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	buttonAPIs := make([]orm.ButtonAPI, 0)

	// for each button, update fields from the database nullable fields
	for idx := range buttonDBs {
		buttonDB := &buttonDBs[idx]
		_ = buttonDB
		var buttonAPI orm.ButtonAPI

		// insertion point for updating fields
		buttonAPI.ID = buttonDB.ID
		buttonDB.CopyBasicFieldsToButton_WOP(&buttonAPI.Button_WOP)
		buttonAPI.ButtonPointersEncoding = buttonDB.ButtonPointersEncoding
		buttonAPIs = append(buttonAPIs, buttonAPI)
	}

	c.JSON(http.StatusOK, buttonAPIs)
}

// PostButton
//
// swagger:route POST /buttons buttons postButton
//
// Creates a button
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostButton(c *gin.Context) {

	mutexButton.Lock()
	defer mutexButton.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostButtons", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoButton.GetDB()

	// Validate input
	var input orm.ButtonAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create button
	buttonDB := orm.ButtonDB{}
	buttonDB.ButtonPointersEncoding = input.ButtonPointersEncoding
	buttonDB.CopyBasicFieldsFromButton_WOP(&input.Button_WOP)

	_, err = db.Create(&buttonDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoButton.CheckoutPhaseOneInstance(&buttonDB)
	button := backRepo.BackRepoButton.Map_ButtonDBID_ButtonPtr[buttonDB.ID]

	if button != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), button)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, buttonDB)
}

// GetButton
//
// swagger:route GET /buttons/{ID} buttons getButton
//
// Gets the details for a button.
//
// Responses:
// default: genericError
//
//	200: buttonDBResponse
func (controller *Controller) GetButton(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetButton", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoButton.GetDB()

	// Get buttonDB in DB
	var buttonDB orm.ButtonDB
	if _, err := db.First(&buttonDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var buttonAPI orm.ButtonAPI
	buttonAPI.ID = buttonDB.ID
	buttonAPI.ButtonPointersEncoding = buttonDB.ButtonPointersEncoding
	buttonDB.CopyBasicFieldsToButton_WOP(&buttonAPI.Button_WOP)

	c.JSON(http.StatusOK, buttonAPI)
}

// UpdateButton
//
// swagger:route PATCH /buttons/{ID} buttons updateButton
//
// # Update a button
//
// Responses:
// default: genericError
//
//	200: buttonDBResponse
func (controller *Controller) UpdateButton(c *gin.Context) {

	mutexButton.Lock()
	defer mutexButton.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateButton", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoButton.GetDB()

	// Validate input
	var input orm.ButtonAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var buttonDB orm.ButtonDB

	// fetch the button
	_, err := db.First(&buttonDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	buttonDB.CopyBasicFieldsFromButton_WOP(&input.Button_WOP)
	buttonDB.ButtonPointersEncoding = input.ButtonPointersEncoding

	db, _ = db.Model(&buttonDB)
	_, err = db.Updates(&buttonDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	buttonNew := new(models.Button)
	buttonDB.CopyBasicFieldsToButton(buttonNew)

	// redeem pointers
	buttonDB.DecodePointers(backRepo, buttonNew)

	// get stage instance from DB instance, and call callback function
	buttonOld := backRepo.BackRepoButton.Map_ButtonDBID_ButtonPtr[buttonDB.ID]
	if buttonOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), buttonOld, buttonNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the buttonDB
	c.JSON(http.StatusOK, buttonDB)
}

// DeleteButton
//
// swagger:route DELETE /buttons/{ID} buttons deleteButton
//
// # Delete a button
//
// default: genericError
//
//	200: buttonDBResponse
func (controller *Controller) DeleteButton(c *gin.Context) {

	mutexButton.Lock()
	defer mutexButton.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteButton", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoButton.GetDB()

	// Get model if exist
	var buttonDB orm.ButtonDB
	if _, err := db.First(&buttonDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&buttonDB)

	// get an instance (not staged) from DB instance, and call callback function
	buttonDeleted := new(models.Button)
	buttonDB.CopyBasicFieldsToButton(buttonDeleted)

	// get stage instance from DB instance, and call callback function
	buttonStaged := backRepo.BackRepoButton.Map_ButtonDBID_ButtonPtr[buttonDB.ID]
	if buttonStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), buttonStaged, buttonDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
