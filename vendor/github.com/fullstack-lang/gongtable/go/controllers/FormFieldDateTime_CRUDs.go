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

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormFieldDateTimes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldDateTime.GetDB()

	_, err := db.Find(&formfielddatetimeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
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
		formfielddatetimeDB.CopyBasicFieldsToFormFieldDateTime_WOP(&formfielddatetimeAPI.FormFieldDateTime_WOP)
		formfielddatetimeAPI.FormFieldDateTimePointersEncoding = formfielddatetimeDB.FormFieldDateTimePointersEncoding
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
	defer mutexFormFieldDateTime.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFormFieldDateTimes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
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
	formfielddatetimeDB.FormFieldDateTimePointersEncoding = input.FormFieldDateTimePointersEncoding
	formfielddatetimeDB.CopyBasicFieldsFromFormFieldDateTime_WOP(&input.FormFieldDateTime_WOP)

	_, err = db.Create(&formfielddatetimeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
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

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormFieldDateTime", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldDateTime.GetDB()

	// Get formfielddatetimeDB in DB
	var formfielddatetimeDB orm.FormFieldDateTimeDB
	if _, err := db.First(&formfielddatetimeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var formfielddatetimeAPI orm.FormFieldDateTimeAPI
	formfielddatetimeAPI.ID = formfielddatetimeDB.ID
	formfielddatetimeAPI.FormFieldDateTimePointersEncoding = formfielddatetimeDB.FormFieldDateTimePointersEncoding
	formfielddatetimeDB.CopyBasicFieldsToFormFieldDateTime_WOP(&formfielddatetimeAPI.FormFieldDateTime_WOP)

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
	defer mutexFormFieldDateTime.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFormFieldDateTime", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
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
	_, err := db.First(&formfielddatetimeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formfielddatetimeDB.CopyBasicFieldsFromFormFieldDateTime_WOP(&input.FormFieldDateTime_WOP)
	formfielddatetimeDB.FormFieldDateTimePointersEncoding = input.FormFieldDateTimePointersEncoding

	db, _ = db.Model(&formfielddatetimeDB)
	_, err = db.Updates(&formfielddatetimeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formfielddatetimeNew := new(models.FormFieldDateTime)
	formfielddatetimeDB.CopyBasicFieldsToFormFieldDateTime(formfielddatetimeNew)

	// redeem pointers
	formfielddatetimeDB.DecodePointers(backRepo, formfielddatetimeNew)

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
	defer mutexFormFieldDateTime.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFormFieldDateTime", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldDateTime.GetDB()

	// Get model if exist
	var formfielddatetimeDB orm.FormFieldDateTimeDB
	if _, err := db.First(&formfielddatetimeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&formfielddatetimeDB)

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
}
