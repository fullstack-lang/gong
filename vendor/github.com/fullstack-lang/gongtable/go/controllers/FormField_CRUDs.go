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
var __FormField__dummysDeclaration__ models.FormField
var __FormField_time__dummyDeclaration time.Duration

var mutexFormField sync.Mutex

// An FormFieldID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFormField updateFormField deleteFormField
type FormFieldID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormFieldInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFormField updateFormField
type FormFieldInput struct {
	// The FormField to submit or modify
	// in: body
	FormField *orm.FormFieldAPI
}

// GetFormFields
//
// swagger:route GET /formfields formfields getFormFields
//
// # Get all formfields
//
// Responses:
// default: genericError
//
//	200: formfieldDBResponse
func (controller *Controller) GetFormFields(c *gin.Context) {

	// source slice
	var formfieldDBs []orm.FormFieldDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormFields", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormField.GetDB()

	_, err := db.Find(&formfieldDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	formfieldAPIs := make([]orm.FormFieldAPI, 0)

	// for each formfield, update fields from the database nullable fields
	for idx := range formfieldDBs {
		formfieldDB := &formfieldDBs[idx]
		_ = formfieldDB
		var formfieldAPI orm.FormFieldAPI

		// insertion point for updating fields
		formfieldAPI.ID = formfieldDB.ID
		formfieldDB.CopyBasicFieldsToFormField_WOP(&formfieldAPI.FormField_WOP)
		formfieldAPI.FormFieldPointersEncoding = formfieldDB.FormFieldPointersEncoding
		formfieldAPIs = append(formfieldAPIs, formfieldAPI)
	}

	c.JSON(http.StatusOK, formfieldAPIs)
}

// PostFormField
//
// swagger:route POST /formfields formfields postFormField
//
// Creates a formfield
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFormField(c *gin.Context) {

	mutexFormField.Lock()
	defer mutexFormField.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFormFields", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormField.GetDB()

	// Validate input
	var input orm.FormFieldAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create formfield
	formfieldDB := orm.FormFieldDB{}
	formfieldDB.FormFieldPointersEncoding = input.FormFieldPointersEncoding
	formfieldDB.CopyBasicFieldsFromFormField_WOP(&input.FormField_WOP)

	_, err = db.Create(&formfieldDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFormField.CheckoutPhaseOneInstance(&formfieldDB)
	formfield := backRepo.BackRepoFormField.Map_FormFieldDBID_FormFieldPtr[formfieldDB.ID]

	if formfield != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), formfield)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, formfieldDB)
}

// GetFormField
//
// swagger:route GET /formfields/{ID} formfields getFormField
//
// Gets the details for a formfield.
//
// Responses:
// default: genericError
//
//	200: formfieldDBResponse
func (controller *Controller) GetFormField(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormField", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormField.GetDB()

	// Get formfieldDB in DB
	var formfieldDB orm.FormFieldDB
	if _, err := db.First(&formfieldDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var formfieldAPI orm.FormFieldAPI
	formfieldAPI.ID = formfieldDB.ID
	formfieldAPI.FormFieldPointersEncoding = formfieldDB.FormFieldPointersEncoding
	formfieldDB.CopyBasicFieldsToFormField_WOP(&formfieldAPI.FormField_WOP)

	c.JSON(http.StatusOK, formfieldAPI)
}

// UpdateFormField
//
// swagger:route PATCH /formfields/{ID} formfields updateFormField
//
// # Update a formfield
//
// Responses:
// default: genericError
//
//	200: formfieldDBResponse
func (controller *Controller) UpdateFormField(c *gin.Context) {

	mutexFormField.Lock()
	defer mutexFormField.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFormField", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormField.GetDB()

	// Validate input
	var input orm.FormFieldAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formfieldDB orm.FormFieldDB

	// fetch the formfield
	_, err := db.First(&formfieldDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formfieldDB.CopyBasicFieldsFromFormField_WOP(&input.FormField_WOP)
	formfieldDB.FormFieldPointersEncoding = input.FormFieldPointersEncoding

	db, _ = db.Model(&formfieldDB)
	_, err = db.Updates(&formfieldDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formfieldNew := new(models.FormField)
	formfieldDB.CopyBasicFieldsToFormField(formfieldNew)

	// redeem pointers
	formfieldDB.DecodePointers(backRepo, formfieldNew)

	// get stage instance from DB instance, and call callback function
	formfieldOld := backRepo.BackRepoFormField.Map_FormFieldDBID_FormFieldPtr[formfieldDB.ID]
	if formfieldOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), formfieldOld, formfieldNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formfieldDB
	c.JSON(http.StatusOK, formfieldDB)
}

// DeleteFormField
//
// swagger:route DELETE /formfields/{ID} formfields deleteFormField
//
// # Delete a formfield
//
// default: genericError
//
//	200: formfieldDBResponse
func (controller *Controller) DeleteFormField(c *gin.Context) {

	mutexFormField.Lock()
	defer mutexFormField.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFormField", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormField.GetDB()

	// Get model if exist
	var formfieldDB orm.FormFieldDB
	if _, err := db.First(&formfieldDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&formfieldDB)

	// get an instance (not staged) from DB instance, and call callback function
	formfieldDeleted := new(models.FormField)
	formfieldDB.CopyBasicFieldsToFormField(formfieldDeleted)

	// get stage instance from DB instance, and call callback function
	formfieldStaged := backRepo.BackRepoFormField.Map_FormFieldDBID_FormFieldPtr[formfieldDB.ID]
	if formfieldStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), formfieldStaged, formfieldDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
