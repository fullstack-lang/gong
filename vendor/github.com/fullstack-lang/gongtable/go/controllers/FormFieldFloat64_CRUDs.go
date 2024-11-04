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
var __FormFieldFloat64__dummysDeclaration__ models.FormFieldFloat64
var __FormFieldFloat64_time__dummyDeclaration time.Duration

var mutexFormFieldFloat64 sync.Mutex

// An FormFieldFloat64ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFormFieldFloat64 updateFormFieldFloat64 deleteFormFieldFloat64
type FormFieldFloat64ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormFieldFloat64Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFormFieldFloat64 updateFormFieldFloat64
type FormFieldFloat64Input struct {
	// The FormFieldFloat64 to submit or modify
	// in: body
	FormFieldFloat64 *orm.FormFieldFloat64API
}

// GetFormFieldFloat64s
//
// swagger:route GET /formfieldfloat64s formfieldfloat64s getFormFieldFloat64s
//
// # Get all formfieldfloat64s
//
// Responses:
// default: genericError
//
//	200: formfieldfloat64DBResponse
func (controller *Controller) GetFormFieldFloat64s(c *gin.Context) {

	// source slice
	var formfieldfloat64DBs []orm.FormFieldFloat64DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormFieldFloat64s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldFloat64.GetDB()

	_, err := db.Find(&formfieldfloat64DBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	formfieldfloat64APIs := make([]orm.FormFieldFloat64API, 0)

	// for each formfieldfloat64, update fields from the database nullable fields
	for idx := range formfieldfloat64DBs {
		formfieldfloat64DB := &formfieldfloat64DBs[idx]
		_ = formfieldfloat64DB
		var formfieldfloat64API orm.FormFieldFloat64API

		// insertion point for updating fields
		formfieldfloat64API.ID = formfieldfloat64DB.ID
		formfieldfloat64DB.CopyBasicFieldsToFormFieldFloat64_WOP(&formfieldfloat64API.FormFieldFloat64_WOP)
		formfieldfloat64API.FormFieldFloat64PointersEncoding = formfieldfloat64DB.FormFieldFloat64PointersEncoding
		formfieldfloat64APIs = append(formfieldfloat64APIs, formfieldfloat64API)
	}

	c.JSON(http.StatusOK, formfieldfloat64APIs)
}

// PostFormFieldFloat64
//
// swagger:route POST /formfieldfloat64s formfieldfloat64s postFormFieldFloat64
//
// Creates a formfieldfloat64
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFormFieldFloat64(c *gin.Context) {

	mutexFormFieldFloat64.Lock()
	defer mutexFormFieldFloat64.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFormFieldFloat64s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldFloat64.GetDB()

	// Validate input
	var input orm.FormFieldFloat64API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create formfieldfloat64
	formfieldfloat64DB := orm.FormFieldFloat64DB{}
	formfieldfloat64DB.FormFieldFloat64PointersEncoding = input.FormFieldFloat64PointersEncoding
	formfieldfloat64DB.CopyBasicFieldsFromFormFieldFloat64_WOP(&input.FormFieldFloat64_WOP)

	_, err = db.Create(&formfieldfloat64DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFormFieldFloat64.CheckoutPhaseOneInstance(&formfieldfloat64DB)
	formfieldfloat64 := backRepo.BackRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64Ptr[formfieldfloat64DB.ID]

	if formfieldfloat64 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), formfieldfloat64)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, formfieldfloat64DB)
}

// GetFormFieldFloat64
//
// swagger:route GET /formfieldfloat64s/{ID} formfieldfloat64s getFormFieldFloat64
//
// Gets the details for a formfieldfloat64.
//
// Responses:
// default: genericError
//
//	200: formfieldfloat64DBResponse
func (controller *Controller) GetFormFieldFloat64(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormFieldFloat64", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldFloat64.GetDB()

	// Get formfieldfloat64DB in DB
	var formfieldfloat64DB orm.FormFieldFloat64DB
	if _, err := db.First(&formfieldfloat64DB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var formfieldfloat64API orm.FormFieldFloat64API
	formfieldfloat64API.ID = formfieldfloat64DB.ID
	formfieldfloat64API.FormFieldFloat64PointersEncoding = formfieldfloat64DB.FormFieldFloat64PointersEncoding
	formfieldfloat64DB.CopyBasicFieldsToFormFieldFloat64_WOP(&formfieldfloat64API.FormFieldFloat64_WOP)

	c.JSON(http.StatusOK, formfieldfloat64API)
}

// UpdateFormFieldFloat64
//
// swagger:route PATCH /formfieldfloat64s/{ID} formfieldfloat64s updateFormFieldFloat64
//
// # Update a formfieldfloat64
//
// Responses:
// default: genericError
//
//	200: formfieldfloat64DBResponse
func (controller *Controller) UpdateFormFieldFloat64(c *gin.Context) {

	mutexFormFieldFloat64.Lock()
	defer mutexFormFieldFloat64.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFormFieldFloat64", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldFloat64.GetDB()

	// Validate input
	var input orm.FormFieldFloat64API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formfieldfloat64DB orm.FormFieldFloat64DB

	// fetch the formfieldfloat64
	_, err := db.First(&formfieldfloat64DB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formfieldfloat64DB.CopyBasicFieldsFromFormFieldFloat64_WOP(&input.FormFieldFloat64_WOP)
	formfieldfloat64DB.FormFieldFloat64PointersEncoding = input.FormFieldFloat64PointersEncoding

	db, _ = db.Model(&formfieldfloat64DB)
	_, err = db.Updates(&formfieldfloat64DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formfieldfloat64New := new(models.FormFieldFloat64)
	formfieldfloat64DB.CopyBasicFieldsToFormFieldFloat64(formfieldfloat64New)

	// redeem pointers
	formfieldfloat64DB.DecodePointers(backRepo, formfieldfloat64New)

	// get stage instance from DB instance, and call callback function
	formfieldfloat64Old := backRepo.BackRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64Ptr[formfieldfloat64DB.ID]
	if formfieldfloat64Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), formfieldfloat64Old, formfieldfloat64New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formfieldfloat64DB
	c.JSON(http.StatusOK, formfieldfloat64DB)
}

// DeleteFormFieldFloat64
//
// swagger:route DELETE /formfieldfloat64s/{ID} formfieldfloat64s deleteFormFieldFloat64
//
// # Delete a formfieldfloat64
//
// default: genericError
//
//	200: formfieldfloat64DBResponse
func (controller *Controller) DeleteFormFieldFloat64(c *gin.Context) {

	mutexFormFieldFloat64.Lock()
	defer mutexFormFieldFloat64.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFormFieldFloat64", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldFloat64.GetDB()

	// Get model if exist
	var formfieldfloat64DB orm.FormFieldFloat64DB
	if _, err := db.First(&formfieldfloat64DB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&formfieldfloat64DB)

	// get an instance (not staged) from DB instance, and call callback function
	formfieldfloat64Deleted := new(models.FormFieldFloat64)
	formfieldfloat64DB.CopyBasicFieldsToFormFieldFloat64(formfieldfloat64Deleted)

	// get stage instance from DB instance, and call callback function
	formfieldfloat64Staged := backRepo.BackRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64Ptr[formfieldfloat64DB.ID]
	if formfieldfloat64Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), formfieldfloat64Staged, formfieldfloat64Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
