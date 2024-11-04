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
var __FormFieldSelect__dummysDeclaration__ models.FormFieldSelect
var __FormFieldSelect_time__dummyDeclaration time.Duration

var mutexFormFieldSelect sync.Mutex

// An FormFieldSelectID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFormFieldSelect updateFormFieldSelect deleteFormFieldSelect
type FormFieldSelectID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormFieldSelectInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFormFieldSelect updateFormFieldSelect
type FormFieldSelectInput struct {
	// The FormFieldSelect to submit or modify
	// in: body
	FormFieldSelect *orm.FormFieldSelectAPI
}

// GetFormFieldSelects
//
// swagger:route GET /formfieldselects formfieldselects getFormFieldSelects
//
// # Get all formfieldselects
//
// Responses:
// default: genericError
//
//	200: formfieldselectDBResponse
func (controller *Controller) GetFormFieldSelects(c *gin.Context) {

	// source slice
	var formfieldselectDBs []orm.FormFieldSelectDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormFieldSelects", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldSelect.GetDB()

	_, err := db.Find(&formfieldselectDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	formfieldselectAPIs := make([]orm.FormFieldSelectAPI, 0)

	// for each formfieldselect, update fields from the database nullable fields
	for idx := range formfieldselectDBs {
		formfieldselectDB := &formfieldselectDBs[idx]
		_ = formfieldselectDB
		var formfieldselectAPI orm.FormFieldSelectAPI

		// insertion point for updating fields
		formfieldselectAPI.ID = formfieldselectDB.ID
		formfieldselectDB.CopyBasicFieldsToFormFieldSelect_WOP(&formfieldselectAPI.FormFieldSelect_WOP)
		formfieldselectAPI.FormFieldSelectPointersEncoding = formfieldselectDB.FormFieldSelectPointersEncoding
		formfieldselectAPIs = append(formfieldselectAPIs, formfieldselectAPI)
	}

	c.JSON(http.StatusOK, formfieldselectAPIs)
}

// PostFormFieldSelect
//
// swagger:route POST /formfieldselects formfieldselects postFormFieldSelect
//
// Creates a formfieldselect
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFormFieldSelect(c *gin.Context) {

	mutexFormFieldSelect.Lock()
	defer mutexFormFieldSelect.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFormFieldSelects", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldSelect.GetDB()

	// Validate input
	var input orm.FormFieldSelectAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create formfieldselect
	formfieldselectDB := orm.FormFieldSelectDB{}
	formfieldselectDB.FormFieldSelectPointersEncoding = input.FormFieldSelectPointersEncoding
	formfieldselectDB.CopyBasicFieldsFromFormFieldSelect_WOP(&input.FormFieldSelect_WOP)

	_, err = db.Create(&formfieldselectDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFormFieldSelect.CheckoutPhaseOneInstance(&formfieldselectDB)
	formfieldselect := backRepo.BackRepoFormFieldSelect.Map_FormFieldSelectDBID_FormFieldSelectPtr[formfieldselectDB.ID]

	if formfieldselect != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), formfieldselect)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, formfieldselectDB)
}

// GetFormFieldSelect
//
// swagger:route GET /formfieldselects/{ID} formfieldselects getFormFieldSelect
//
// Gets the details for a formfieldselect.
//
// Responses:
// default: genericError
//
//	200: formfieldselectDBResponse
func (controller *Controller) GetFormFieldSelect(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormFieldSelect", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldSelect.GetDB()

	// Get formfieldselectDB in DB
	var formfieldselectDB orm.FormFieldSelectDB
	if _, err := db.First(&formfieldselectDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var formfieldselectAPI orm.FormFieldSelectAPI
	formfieldselectAPI.ID = formfieldselectDB.ID
	formfieldselectAPI.FormFieldSelectPointersEncoding = formfieldselectDB.FormFieldSelectPointersEncoding
	formfieldselectDB.CopyBasicFieldsToFormFieldSelect_WOP(&formfieldselectAPI.FormFieldSelect_WOP)

	c.JSON(http.StatusOK, formfieldselectAPI)
}

// UpdateFormFieldSelect
//
// swagger:route PATCH /formfieldselects/{ID} formfieldselects updateFormFieldSelect
//
// # Update a formfieldselect
//
// Responses:
// default: genericError
//
//	200: formfieldselectDBResponse
func (controller *Controller) UpdateFormFieldSelect(c *gin.Context) {

	mutexFormFieldSelect.Lock()
	defer mutexFormFieldSelect.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFormFieldSelect", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldSelect.GetDB()

	// Validate input
	var input orm.FormFieldSelectAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formfieldselectDB orm.FormFieldSelectDB

	// fetch the formfieldselect
	_, err := db.First(&formfieldselectDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formfieldselectDB.CopyBasicFieldsFromFormFieldSelect_WOP(&input.FormFieldSelect_WOP)
	formfieldselectDB.FormFieldSelectPointersEncoding = input.FormFieldSelectPointersEncoding

	db, _ = db.Model(&formfieldselectDB)
	_, err = db.Updates(&formfieldselectDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formfieldselectNew := new(models.FormFieldSelect)
	formfieldselectDB.CopyBasicFieldsToFormFieldSelect(formfieldselectNew)

	// redeem pointers
	formfieldselectDB.DecodePointers(backRepo, formfieldselectNew)

	// get stage instance from DB instance, and call callback function
	formfieldselectOld := backRepo.BackRepoFormFieldSelect.Map_FormFieldSelectDBID_FormFieldSelectPtr[formfieldselectDB.ID]
	if formfieldselectOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), formfieldselectOld, formfieldselectNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formfieldselectDB
	c.JSON(http.StatusOK, formfieldselectDB)
}

// DeleteFormFieldSelect
//
// swagger:route DELETE /formfieldselects/{ID} formfieldselects deleteFormFieldSelect
//
// # Delete a formfieldselect
//
// default: genericError
//
//	200: formfieldselectDBResponse
func (controller *Controller) DeleteFormFieldSelect(c *gin.Context) {

	mutexFormFieldSelect.Lock()
	defer mutexFormFieldSelect.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFormFieldSelect", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldSelect.GetDB()

	// Get model if exist
	var formfieldselectDB orm.FormFieldSelectDB
	if _, err := db.First(&formfieldselectDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&formfieldselectDB)

	// get an instance (not staged) from DB instance, and call callback function
	formfieldselectDeleted := new(models.FormFieldSelect)
	formfieldselectDB.CopyBasicFieldsToFormFieldSelect(formfieldselectDeleted)

	// get stage instance from DB instance, and call callback function
	formfieldselectStaged := backRepo.BackRepoFormFieldSelect.Map_FormFieldSelectDBID_FormFieldSelectPtr[formfieldselectDB.ID]
	if formfieldselectStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), formfieldselectStaged, formfieldselectDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
