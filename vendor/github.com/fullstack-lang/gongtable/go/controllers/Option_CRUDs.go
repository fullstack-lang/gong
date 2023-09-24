// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongtable/go/models"
	"github.com/fullstack-lang/gongtable/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Option__dummysDeclaration__ models.Option
var __Option_time__dummyDeclaration time.Duration

var mutexOption sync.Mutex

// An OptionID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getOption updateOption deleteOption
type OptionID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// OptionInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postOption updateOption
type OptionInput struct {
	// The Option to submit or modify
	// in: body
	Option *orm.OptionAPI
}

// GetOptions
//
// swagger:route GET /options options getOptions
//
// # Get all options
//
// Responses:
// default: genericError
//
//	200: optionDBResponse
func (controller *Controller) GetOptions(c *gin.Context) {

	// source slice
	var optionDBs []orm.OptionDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOptions", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOption.GetDB()

	query := db.Find(&optionDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	optionAPIs := make([]orm.OptionAPI, 0)

	// for each option, update fields from the database nullable fields
	for idx := range optionDBs {
		optionDB := &optionDBs[idx]
		_ = optionDB
		var optionAPI orm.OptionAPI

		// insertion point for updating fields
		optionAPI.ID = optionDB.ID
		optionDB.CopyBasicFieldsToOption(&optionAPI.Option)
		optionAPI.OptionPointersEnconding = optionDB.OptionPointersEnconding
		optionAPIs = append(optionAPIs, optionAPI)
	}

	c.JSON(http.StatusOK, optionAPIs)
}

// PostOption
//
// swagger:route POST /options options postOption
//
// Creates a option
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostOption(c *gin.Context) {

	mutexOption.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostOptions", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOption.GetDB()

	// Validate input
	var input orm.OptionAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create option
	optionDB := orm.OptionDB{}
	optionDB.OptionPointersEnconding = input.OptionPointersEnconding
	optionDB.CopyBasicFieldsFromOption(&input.Option)

	query := db.Create(&optionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoOption.CheckoutPhaseOneInstance(&optionDB)
	option := backRepo.BackRepoOption.Map_OptionDBID_OptionPtr[optionDB.ID]

	if option != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), option)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, optionDB)

	mutexOption.Unlock()
}

// GetOption
//
// swagger:route GET /options/{ID} options getOption
//
// Gets the details for a option.
//
// Responses:
// default: genericError
//
//	200: optionDBResponse
func (controller *Controller) GetOption(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOption", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOption.GetDB()

	// Get optionDB in DB
	var optionDB orm.OptionDB
	if err := db.First(&optionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var optionAPI orm.OptionAPI
	optionAPI.ID = optionDB.ID
	optionAPI.OptionPointersEnconding = optionDB.OptionPointersEnconding
	optionDB.CopyBasicFieldsToOption(&optionAPI.Option)

	c.JSON(http.StatusOK, optionAPI)
}

// UpdateOption
//
// swagger:route PATCH /options/{ID} options updateOption
//
// # Update a option
//
// Responses:
// default: genericError
//
//	200: optionDBResponse
func (controller *Controller) UpdateOption(c *gin.Context) {

	mutexOption.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateOption", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOption.GetDB()

	// Validate input
	var input orm.OptionAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var optionDB orm.OptionDB

	// fetch the option
	query := db.First(&optionDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	optionDB.CopyBasicFieldsFromOption(&input.Option)
	optionDB.OptionPointersEnconding = input.OptionPointersEnconding

	query = db.Model(&optionDB).Updates(optionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	optionNew := new(models.Option)
	optionDB.CopyBasicFieldsToOption(optionNew)

	// get stage instance from DB instance, and call callback function
	optionOld := backRepo.BackRepoOption.Map_OptionDBID_OptionPtr[optionDB.ID]
	if optionOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), optionOld, optionNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the optionDB
	c.JSON(http.StatusOK, optionDB)

	mutexOption.Unlock()
}

// DeleteOption
//
// swagger:route DELETE /options/{ID} options deleteOption
//
// # Delete a option
//
// default: genericError
//
//	200: optionDBResponse
func (controller *Controller) DeleteOption(c *gin.Context) {

	mutexOption.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteOption", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOption.GetDB()

	// Get model if exist
	var optionDB orm.OptionDB
	if err := db.First(&optionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&optionDB)

	// get an instance (not staged) from DB instance, and call callback function
	optionDeleted := new(models.Option)
	optionDB.CopyBasicFieldsToOption(optionDeleted)

	// get stage instance from DB instance, and call callback function
	optionStaged := backRepo.BackRepoOption.Map_OptionDBID_OptionPtr[optionDB.ID]
	if optionStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), optionStaged, optionDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexOption.Unlock()
}
