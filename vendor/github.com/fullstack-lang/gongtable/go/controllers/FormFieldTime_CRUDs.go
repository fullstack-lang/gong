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
var __FormFieldTime__dummysDeclaration__ models.FormFieldTime
var __FormFieldTime_time__dummyDeclaration time.Duration

var mutexFormFieldTime sync.Mutex

// An FormFieldTimeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFormFieldTime updateFormFieldTime deleteFormFieldTime
type FormFieldTimeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormFieldTimeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFormFieldTime updateFormFieldTime
type FormFieldTimeInput struct {
	// The FormFieldTime to submit or modify
	// in: body
	FormFieldTime *orm.FormFieldTimeAPI
}

// GetFormFieldTimes
//
// swagger:route GET /formfieldtimes formfieldtimes getFormFieldTimes
//
// # Get all formfieldtimes
//
// Responses:
// default: genericError
//
//	200: formfieldtimeDBResponse
func (controller *Controller) GetFormFieldTimes(c *gin.Context) {

	// source slice
	var formfieldtimeDBs []orm.FormFieldTimeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormFieldTimes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldTime.GetDB()

	_, err := db.Find(&formfieldtimeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	formfieldtimeAPIs := make([]orm.FormFieldTimeAPI, 0)

	// for each formfieldtime, update fields from the database nullable fields
	for idx := range formfieldtimeDBs {
		formfieldtimeDB := &formfieldtimeDBs[idx]
		_ = formfieldtimeDB
		var formfieldtimeAPI orm.FormFieldTimeAPI

		// insertion point for updating fields
		formfieldtimeAPI.ID = formfieldtimeDB.ID
		formfieldtimeDB.CopyBasicFieldsToFormFieldTime_WOP(&formfieldtimeAPI.FormFieldTime_WOP)
		formfieldtimeAPI.FormFieldTimePointersEncoding = formfieldtimeDB.FormFieldTimePointersEncoding
		formfieldtimeAPIs = append(formfieldtimeAPIs, formfieldtimeAPI)
	}

	c.JSON(http.StatusOK, formfieldtimeAPIs)
}

// PostFormFieldTime
//
// swagger:route POST /formfieldtimes formfieldtimes postFormFieldTime
//
// Creates a formfieldtime
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFormFieldTime(c *gin.Context) {

	mutexFormFieldTime.Lock()
	defer mutexFormFieldTime.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFormFieldTimes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldTime.GetDB()

	// Validate input
	var input orm.FormFieldTimeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create formfieldtime
	formfieldtimeDB := orm.FormFieldTimeDB{}
	formfieldtimeDB.FormFieldTimePointersEncoding = input.FormFieldTimePointersEncoding
	formfieldtimeDB.CopyBasicFieldsFromFormFieldTime_WOP(&input.FormFieldTime_WOP)

	_, err = db.Create(&formfieldtimeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFormFieldTime.CheckoutPhaseOneInstance(&formfieldtimeDB)
	formfieldtime := backRepo.BackRepoFormFieldTime.Map_FormFieldTimeDBID_FormFieldTimePtr[formfieldtimeDB.ID]

	if formfieldtime != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), formfieldtime)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, formfieldtimeDB)
}

// GetFormFieldTime
//
// swagger:route GET /formfieldtimes/{ID} formfieldtimes getFormFieldTime
//
// Gets the details for a formfieldtime.
//
// Responses:
// default: genericError
//
//	200: formfieldtimeDBResponse
func (controller *Controller) GetFormFieldTime(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormFieldTime", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldTime.GetDB()

	// Get formfieldtimeDB in DB
	var formfieldtimeDB orm.FormFieldTimeDB
	if _, err := db.First(&formfieldtimeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var formfieldtimeAPI orm.FormFieldTimeAPI
	formfieldtimeAPI.ID = formfieldtimeDB.ID
	formfieldtimeAPI.FormFieldTimePointersEncoding = formfieldtimeDB.FormFieldTimePointersEncoding
	formfieldtimeDB.CopyBasicFieldsToFormFieldTime_WOP(&formfieldtimeAPI.FormFieldTime_WOP)

	c.JSON(http.StatusOK, formfieldtimeAPI)
}

// UpdateFormFieldTime
//
// swagger:route PATCH /formfieldtimes/{ID} formfieldtimes updateFormFieldTime
//
// # Update a formfieldtime
//
// Responses:
// default: genericError
//
//	200: formfieldtimeDBResponse
func (controller *Controller) UpdateFormFieldTime(c *gin.Context) {

	mutexFormFieldTime.Lock()
	defer mutexFormFieldTime.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFormFieldTime", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldTime.GetDB()

	// Validate input
	var input orm.FormFieldTimeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formfieldtimeDB orm.FormFieldTimeDB

	// fetch the formfieldtime
	_, err := db.First(&formfieldtimeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formfieldtimeDB.CopyBasicFieldsFromFormFieldTime_WOP(&input.FormFieldTime_WOP)
	formfieldtimeDB.FormFieldTimePointersEncoding = input.FormFieldTimePointersEncoding

	db, _ = db.Model(&formfieldtimeDB)
	_, err = db.Updates(&formfieldtimeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formfieldtimeNew := new(models.FormFieldTime)
	formfieldtimeDB.CopyBasicFieldsToFormFieldTime(formfieldtimeNew)

	// redeem pointers
	formfieldtimeDB.DecodePointers(backRepo, formfieldtimeNew)

	// get stage instance from DB instance, and call callback function
	formfieldtimeOld := backRepo.BackRepoFormFieldTime.Map_FormFieldTimeDBID_FormFieldTimePtr[formfieldtimeDB.ID]
	if formfieldtimeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), formfieldtimeOld, formfieldtimeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formfieldtimeDB
	c.JSON(http.StatusOK, formfieldtimeDB)
}

// DeleteFormFieldTime
//
// swagger:route DELETE /formfieldtimes/{ID} formfieldtimes deleteFormFieldTime
//
// # Delete a formfieldtime
//
// default: genericError
//
//	200: formfieldtimeDBResponse
func (controller *Controller) DeleteFormFieldTime(c *gin.Context) {

	mutexFormFieldTime.Lock()
	defer mutexFormFieldTime.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFormFieldTime", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldTime.GetDB()

	// Get model if exist
	var formfieldtimeDB orm.FormFieldTimeDB
	if _, err := db.First(&formfieldtimeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&formfieldtimeDB)

	// get an instance (not staged) from DB instance, and call callback function
	formfieldtimeDeleted := new(models.FormFieldTime)
	formfieldtimeDB.CopyBasicFieldsToFormFieldTime(formfieldtimeDeleted)

	// get stage instance from DB instance, and call callback function
	formfieldtimeStaged := backRepo.BackRepoFormFieldTime.Map_FormFieldTimeDBID_FormFieldTimePtr[formfieldtimeDB.ID]
	if formfieldtimeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), formfieldtimeStaged, formfieldtimeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
