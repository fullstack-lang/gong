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
var __FormFieldDateTime__dummysDeclaration__ models.FormFieldDateTime
var __FormFieldDateTime_time__dummyDeclaration time.Duration

var mutexFormFieldDateTime sync.Mutex

// An FormFieldDateTimeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFormFieldDateTime updateFormFieldDateTime deleteFormFieldDateTime
type FormFieldDateTimeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormFieldDateTimeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFormFieldDateTime updateFormFieldDateTime
type FormFieldDateTimeInput struct {
	// The FormFieldDateTime to submit or modify
	// in: body
	FormFieldDateTime *orm.FormFieldDateTimeAPI
}

// GetFormFieldDateTimes
//
// swagger:route GET /formfielddatetimes formfielddatetimes getFormFieldDateTimes
//
// # Get all formfielddatetimes
//
// Responses:
// default: genericError
//
//	200: formfielddatetimeDBResponse
func (controller *Controller) GetFormFieldDateTimes(c *gin.Context) {

	// source slice
	var formfielddatetimeDBs []orm.FormFieldDateTimeDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormFieldDateTimes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoFormFieldDateTime.GetDB()

	query := db.Find(&formfielddatetimeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	formfielddatetimeAPIs := make([]orm.FormFieldDateTimeAPI, 0)

	// for each formfielddatetime, update fields from the database nullable fields
	for idx := range formfielddatetimeDBs {
		formfielddatetimeDB := &formfielddatetimeDBs[idx]
		_ = formfielddatetimeDB
		var formfielddatetimeAPI orm.FormFieldDateTimeAPI

		// insertion point for updating fields
		formfielddatetimeAPI.ID = formfielddatetimeDB.ID
		formfielddatetimeDB.CopyBasicFieldsToFormFieldDateTime(&formfielddatetimeAPI.FormFieldDateTime)
		formfielddatetimeAPI.FormFieldDateTimePointersEnconding = formfielddatetimeDB.FormFieldDateTimePointersEnconding
		formfielddatetimeAPIs = append(formfielddatetimeAPIs, formfielddatetimeAPI)
	}

	c.JSON(http.StatusOK, formfielddatetimeAPIs)
}

// PostFormFieldDateTime
//
// swagger:route POST /formfielddatetimes formfielddatetimes postFormFieldDateTime
//
// Creates a formfielddatetime
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFormFieldDateTime(c *gin.Context) {

	mutexFormFieldDateTime.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFormFieldDateTimes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoFormFieldDateTime.GetDB()

	// Validate input
	var input orm.FormFieldDateTimeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create formfielddatetime
	formfielddatetimeDB := orm.FormFieldDateTimeDB{}
	formfielddatetimeDB.FormFieldDateTimePointersEnconding = input.FormFieldDateTimePointersEnconding
	formfielddatetimeDB.CopyBasicFieldsFromFormFieldDateTime(&input.FormFieldDateTime)

	query := db.Create(&formfielddatetimeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFormFieldDateTime.CheckoutPhaseOneInstance(&formfielddatetimeDB)
	formfielddatetime := backRepo.BackRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimePtr[formfielddatetimeDB.ID]

	if formfielddatetime != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), formfielddatetime)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, formfielddatetimeDB)

	mutexFormFieldDateTime.Unlock()
}

// GetFormFieldDateTime
//
// swagger:route GET /formfielddatetimes/{ID} formfielddatetimes getFormFieldDateTime
//
// Gets the details for a formfielddatetime.
//
// Responses:
// default: genericError
//
//	200: formfielddatetimeDBResponse
func (controller *Controller) GetFormFieldDateTime(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormFieldDateTime", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoFormFieldDateTime.GetDB()

	// Get formfielddatetimeDB in DB
	var formfielddatetimeDB orm.FormFieldDateTimeDB
	if err := db.First(&formfielddatetimeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var formfielddatetimeAPI orm.FormFieldDateTimeAPI
	formfielddatetimeAPI.ID = formfielddatetimeDB.ID
	formfielddatetimeAPI.FormFieldDateTimePointersEnconding = formfielddatetimeDB.FormFieldDateTimePointersEnconding
	formfielddatetimeDB.CopyBasicFieldsToFormFieldDateTime(&formfielddatetimeAPI.FormFieldDateTime)

	c.JSON(http.StatusOK, formfielddatetimeAPI)
}

// UpdateFormFieldDateTime
//
// swagger:route PATCH /formfielddatetimes/{ID} formfielddatetimes updateFormFieldDateTime
//
// # Update a formfielddatetime
//
// Responses:
// default: genericError
//
//	200: formfielddatetimeDBResponse
func (controller *Controller) UpdateFormFieldDateTime(c *gin.Context) {

	mutexFormFieldDateTime.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFormFieldDateTime", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoFormFieldDateTime.GetDB()

	// Validate input
	var input orm.FormFieldDateTimeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formfielddatetimeDB orm.FormFieldDateTimeDB

	// fetch the formfielddatetime
	query := db.First(&formfielddatetimeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formfielddatetimeDB.CopyBasicFieldsFromFormFieldDateTime(&input.FormFieldDateTime)
	formfielddatetimeDB.FormFieldDateTimePointersEnconding = input.FormFieldDateTimePointersEnconding

	query = db.Model(&formfielddatetimeDB).Updates(formfielddatetimeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formfielddatetimeNew := new(models.FormFieldDateTime)
	formfielddatetimeDB.CopyBasicFieldsToFormFieldDateTime(formfielddatetimeNew)

	// get stage instance from DB instance, and call callback function
	formfielddatetimeOld := backRepo.BackRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimePtr[formfielddatetimeDB.ID]
	if formfielddatetimeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), formfielddatetimeOld, formfielddatetimeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formfielddatetimeDB
	c.JSON(http.StatusOK, formfielddatetimeDB)

	mutexFormFieldDateTime.Unlock()
}

// DeleteFormFieldDateTime
//
// swagger:route DELETE /formfielddatetimes/{ID} formfielddatetimes deleteFormFieldDateTime
//
// # Delete a formfielddatetime
//
// default: genericError
//
//	200: formfielddatetimeDBResponse
func (controller *Controller) DeleteFormFieldDateTime(c *gin.Context) {

	mutexFormFieldDateTime.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFormFieldDateTime", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoFormFieldDateTime.GetDB()

	// Get model if exist
	var formfielddatetimeDB orm.FormFieldDateTimeDB
	if err := db.First(&formfielddatetimeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&formfielddatetimeDB)

	// get an instance (not staged) from DB instance, and call callback function
	formfielddatetimeDeleted := new(models.FormFieldDateTime)
	formfielddatetimeDB.CopyBasicFieldsToFormFieldDateTime(formfielddatetimeDeleted)

	// get stage instance from DB instance, and call callback function
	formfielddatetimeStaged := backRepo.BackRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimePtr[formfielddatetimeDB.ID]
	if formfielddatetimeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), formfielddatetimeStaged, formfielddatetimeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexFormFieldDateTime.Unlock()
}
