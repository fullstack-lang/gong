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
var __FormFieldDate__dummysDeclaration__ models.FormFieldDate
var __FormFieldDate_time__dummyDeclaration time.Duration

var mutexFormFieldDate sync.Mutex

// An FormFieldDateID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFormFieldDate updateFormFieldDate deleteFormFieldDate
type FormFieldDateID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormFieldDateInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFormFieldDate updateFormFieldDate
type FormFieldDateInput struct {
	// The FormFieldDate to submit or modify
	// in: body
	FormFieldDate *orm.FormFieldDateAPI
}

// GetFormFieldDates
//
// swagger:route GET /formfielddates formfielddates getFormFieldDates
//
// # Get all formfielddates
//
// Responses:
// default: genericError
//
//	200: formfielddateDBResponse
func (controller *Controller) GetFormFieldDates(c *gin.Context) {

	// source slice
	var formfielddateDBs []orm.FormFieldDateDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormFieldDates", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldDate.GetDB()

	_, err := db.Find(&formfielddateDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	formfielddateAPIs := make([]orm.FormFieldDateAPI, 0)

	// for each formfielddate, update fields from the database nullable fields
	for idx := range formfielddateDBs {
		formfielddateDB := &formfielddateDBs[idx]
		_ = formfielddateDB
		var formfielddateAPI orm.FormFieldDateAPI

		// insertion point for updating fields
		formfielddateAPI.ID = formfielddateDB.ID
		formfielddateDB.CopyBasicFieldsToFormFieldDate_WOP(&formfielddateAPI.FormFieldDate_WOP)
		formfielddateAPI.FormFieldDatePointersEncoding = formfielddateDB.FormFieldDatePointersEncoding
		formfielddateAPIs = append(formfielddateAPIs, formfielddateAPI)
	}

	c.JSON(http.StatusOK, formfielddateAPIs)
}

// PostFormFieldDate
//
// swagger:route POST /formfielddates formfielddates postFormFieldDate
//
// Creates a formfielddate
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFormFieldDate(c *gin.Context) {

	mutexFormFieldDate.Lock()
	defer mutexFormFieldDate.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFormFieldDates", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldDate.GetDB()

	// Validate input
	var input orm.FormFieldDateAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create formfielddate
	formfielddateDB := orm.FormFieldDateDB{}
	formfielddateDB.FormFieldDatePointersEncoding = input.FormFieldDatePointersEncoding
	formfielddateDB.CopyBasicFieldsFromFormFieldDate_WOP(&input.FormFieldDate_WOP)

	_, err = db.Create(&formfielddateDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFormFieldDate.CheckoutPhaseOneInstance(&formfielddateDB)
	formfielddate := backRepo.BackRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDatePtr[formfielddateDB.ID]

	if formfielddate != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), formfielddate)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, formfielddateDB)
}

// GetFormFieldDate
//
// swagger:route GET /formfielddates/{ID} formfielddates getFormFieldDate
//
// Gets the details for a formfielddate.
//
// Responses:
// default: genericError
//
//	200: formfielddateDBResponse
func (controller *Controller) GetFormFieldDate(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormFieldDate", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldDate.GetDB()

	// Get formfielddateDB in DB
	var formfielddateDB orm.FormFieldDateDB
	if _, err := db.First(&formfielddateDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var formfielddateAPI orm.FormFieldDateAPI
	formfielddateAPI.ID = formfielddateDB.ID
	formfielddateAPI.FormFieldDatePointersEncoding = formfielddateDB.FormFieldDatePointersEncoding
	formfielddateDB.CopyBasicFieldsToFormFieldDate_WOP(&formfielddateAPI.FormFieldDate_WOP)

	c.JSON(http.StatusOK, formfielddateAPI)
}

// UpdateFormFieldDate
//
// swagger:route PATCH /formfielddates/{ID} formfielddates updateFormFieldDate
//
// # Update a formfielddate
//
// Responses:
// default: genericError
//
//	200: formfielddateDBResponse
func (controller *Controller) UpdateFormFieldDate(c *gin.Context) {

	mutexFormFieldDate.Lock()
	defer mutexFormFieldDate.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFormFieldDate", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldDate.GetDB()

	// Validate input
	var input orm.FormFieldDateAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formfielddateDB orm.FormFieldDateDB

	// fetch the formfielddate
	_, err := db.First(&formfielddateDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formfielddateDB.CopyBasicFieldsFromFormFieldDate_WOP(&input.FormFieldDate_WOP)
	formfielddateDB.FormFieldDatePointersEncoding = input.FormFieldDatePointersEncoding

	db, _ = db.Model(&formfielddateDB)
	_, err = db.Updates(&formfielddateDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formfielddateNew := new(models.FormFieldDate)
	formfielddateDB.CopyBasicFieldsToFormFieldDate(formfielddateNew)

	// redeem pointers
	formfielddateDB.DecodePointers(backRepo, formfielddateNew)

	// get stage instance from DB instance, and call callback function
	formfielddateOld := backRepo.BackRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDatePtr[formfielddateDB.ID]
	if formfielddateOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), formfielddateOld, formfielddateNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formfielddateDB
	c.JSON(http.StatusOK, formfielddateDB)
}

// DeleteFormFieldDate
//
// swagger:route DELETE /formfielddates/{ID} formfielddates deleteFormFieldDate
//
// # Delete a formfielddate
//
// default: genericError
//
//	200: formfielddateDBResponse
func (controller *Controller) DeleteFormFieldDate(c *gin.Context) {

	mutexFormFieldDate.Lock()
	defer mutexFormFieldDate.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFormFieldDate", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldDate.GetDB()

	// Get model if exist
	var formfielddateDB orm.FormFieldDateDB
	if _, err := db.First(&formfielddateDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&formfielddateDB)

	// get an instance (not staged) from DB instance, and call callback function
	formfielddateDeleted := new(models.FormFieldDate)
	formfielddateDB.CopyBasicFieldsToFormFieldDate(formfielddateDeleted)

	// get stage instance from DB instance, and call callback function
	formfielddateStaged := backRepo.BackRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDatePtr[formfielddateDB.ID]
	if formfielddateStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), formfielddateStaged, formfielddateDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
