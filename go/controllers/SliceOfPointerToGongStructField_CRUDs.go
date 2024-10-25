// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __SliceOfPointerToGongStructField__dummysDeclaration__ models.SliceOfPointerToGongStructField
var __SliceOfPointerToGongStructField_time__dummyDeclaration time.Duration

var mutexSliceOfPointerToGongStructField sync.Mutex

// An SliceOfPointerToGongStructFieldID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSliceOfPointerToGongStructField updateSliceOfPointerToGongStructField deleteSliceOfPointerToGongStructField
type SliceOfPointerToGongStructFieldID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SliceOfPointerToGongStructFieldInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSliceOfPointerToGongStructField updateSliceOfPointerToGongStructField
type SliceOfPointerToGongStructFieldInput struct {
	// The SliceOfPointerToGongStructField to submit or modify
	// in: body
	SliceOfPointerToGongStructField *orm.SliceOfPointerToGongStructFieldAPI
}

// GetSliceOfPointerToGongStructFields
//
// swagger:route GET /sliceofpointertogongstructfields sliceofpointertogongstructfields getSliceOfPointerToGongStructFields
//
// # Get all sliceofpointertogongstructfields
//
// Responses:
// default: genericError
//
//	200: sliceofpointertogongstructfieldDBResponse
func (controller *Controller) GetSliceOfPointerToGongStructFields(c *gin.Context) {

	// source slice
	var sliceofpointertogongstructfieldDBs []orm.SliceOfPointerToGongStructFieldDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSliceOfPointerToGongStructFields", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSliceOfPointerToGongStructField.GetDB()

	_, err := db.Find(&sliceofpointertogongstructfieldDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	sliceofpointertogongstructfieldAPIs := make([]orm.SliceOfPointerToGongStructFieldAPI, 0)

	// for each sliceofpointertogongstructfield, update fields from the database nullable fields
	for idx := range sliceofpointertogongstructfieldDBs {
		sliceofpointertogongstructfieldDB := &sliceofpointertogongstructfieldDBs[idx]
		_ = sliceofpointertogongstructfieldDB
		var sliceofpointertogongstructfieldAPI orm.SliceOfPointerToGongStructFieldAPI

		// insertion point for updating fields
		sliceofpointertogongstructfieldAPI.ID = sliceofpointertogongstructfieldDB.ID
		sliceofpointertogongstructfieldDB.CopyBasicFieldsToSliceOfPointerToGongStructField_WOP(&sliceofpointertogongstructfieldAPI.SliceOfPointerToGongStructField_WOP)
		sliceofpointertogongstructfieldAPI.SliceOfPointerToGongStructFieldPointersEncoding = sliceofpointertogongstructfieldDB.SliceOfPointerToGongStructFieldPointersEncoding
		sliceofpointertogongstructfieldAPIs = append(sliceofpointertogongstructfieldAPIs, sliceofpointertogongstructfieldAPI)
	}

	c.JSON(http.StatusOK, sliceofpointertogongstructfieldAPIs)
}

// PostSliceOfPointerToGongStructField
//
// swagger:route POST /sliceofpointertogongstructfields sliceofpointertogongstructfields postSliceOfPointerToGongStructField
//
// Creates a sliceofpointertogongstructfield
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSliceOfPointerToGongStructField(c *gin.Context) {

	mutexSliceOfPointerToGongStructField.Lock()
	defer mutexSliceOfPointerToGongStructField.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSliceOfPointerToGongStructFields", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSliceOfPointerToGongStructField.GetDB()

	// Validate input
	var input orm.SliceOfPointerToGongStructFieldAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create sliceofpointertogongstructfield
	sliceofpointertogongstructfieldDB := orm.SliceOfPointerToGongStructFieldDB{}
	sliceofpointertogongstructfieldDB.SliceOfPointerToGongStructFieldPointersEncoding = input.SliceOfPointerToGongStructFieldPointersEncoding
	sliceofpointertogongstructfieldDB.CopyBasicFieldsFromSliceOfPointerToGongStructField_WOP(&input.SliceOfPointerToGongStructField_WOP)

	_, err = db.Create(&sliceofpointertogongstructfieldDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSliceOfPointerToGongStructField.CheckoutPhaseOneInstance(&sliceofpointertogongstructfieldDB)
	sliceofpointertogongstructfield := backRepo.BackRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldPtr[sliceofpointertogongstructfieldDB.ID]

	if sliceofpointertogongstructfield != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), sliceofpointertogongstructfield)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, sliceofpointertogongstructfieldDB)
}

// GetSliceOfPointerToGongStructField
//
// swagger:route GET /sliceofpointertogongstructfields/{ID} sliceofpointertogongstructfields getSliceOfPointerToGongStructField
//
// Gets the details for a sliceofpointertogongstructfield.
//
// Responses:
// default: genericError
//
//	200: sliceofpointertogongstructfieldDBResponse
func (controller *Controller) GetSliceOfPointerToGongStructField(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSliceOfPointerToGongStructField", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSliceOfPointerToGongStructField.GetDB()

	// Get sliceofpointertogongstructfieldDB in DB
	var sliceofpointertogongstructfieldDB orm.SliceOfPointerToGongStructFieldDB
	if _, err := db.First(&sliceofpointertogongstructfieldDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var sliceofpointertogongstructfieldAPI orm.SliceOfPointerToGongStructFieldAPI
	sliceofpointertogongstructfieldAPI.ID = sliceofpointertogongstructfieldDB.ID
	sliceofpointertogongstructfieldAPI.SliceOfPointerToGongStructFieldPointersEncoding = sliceofpointertogongstructfieldDB.SliceOfPointerToGongStructFieldPointersEncoding
	sliceofpointertogongstructfieldDB.CopyBasicFieldsToSliceOfPointerToGongStructField_WOP(&sliceofpointertogongstructfieldAPI.SliceOfPointerToGongStructField_WOP)

	c.JSON(http.StatusOK, sliceofpointertogongstructfieldAPI)
}

// UpdateSliceOfPointerToGongStructField
//
// swagger:route PATCH /sliceofpointertogongstructfields/{ID} sliceofpointertogongstructfields updateSliceOfPointerToGongStructField
//
// # Update a sliceofpointertogongstructfield
//
// Responses:
// default: genericError
//
//	200: sliceofpointertogongstructfieldDBResponse
func (controller *Controller) UpdateSliceOfPointerToGongStructField(c *gin.Context) {

	mutexSliceOfPointerToGongStructField.Lock()
	defer mutexSliceOfPointerToGongStructField.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSliceOfPointerToGongStructField", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSliceOfPointerToGongStructField.GetDB()

	// Validate input
	var input orm.SliceOfPointerToGongStructFieldAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var sliceofpointertogongstructfieldDB orm.SliceOfPointerToGongStructFieldDB

	// fetch the sliceofpointertogongstructfield
	_, err := db.First(&sliceofpointertogongstructfieldDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	sliceofpointertogongstructfieldDB.CopyBasicFieldsFromSliceOfPointerToGongStructField_WOP(&input.SliceOfPointerToGongStructField_WOP)
	sliceofpointertogongstructfieldDB.SliceOfPointerToGongStructFieldPointersEncoding = input.SliceOfPointerToGongStructFieldPointersEncoding

	db, _ = db.Model(&sliceofpointertogongstructfieldDB)
	_, err = db.Updates(&sliceofpointertogongstructfieldDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	sliceofpointertogongstructfieldNew := new(models.SliceOfPointerToGongStructField)
	sliceofpointertogongstructfieldDB.CopyBasicFieldsToSliceOfPointerToGongStructField(sliceofpointertogongstructfieldNew)

	// redeem pointers
	sliceofpointertogongstructfieldDB.DecodePointers(backRepo, sliceofpointertogongstructfieldNew)

	// get stage instance from DB instance, and call callback function
	sliceofpointertogongstructfieldOld := backRepo.BackRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldPtr[sliceofpointertogongstructfieldDB.ID]
	if sliceofpointertogongstructfieldOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), sliceofpointertogongstructfieldOld, sliceofpointertogongstructfieldNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the sliceofpointertogongstructfieldDB
	c.JSON(http.StatusOK, sliceofpointertogongstructfieldDB)
}

// DeleteSliceOfPointerToGongStructField
//
// swagger:route DELETE /sliceofpointertogongstructfields/{ID} sliceofpointertogongstructfields deleteSliceOfPointerToGongStructField
//
// # Delete a sliceofpointertogongstructfield
//
// default: genericError
//
//	200: sliceofpointertogongstructfieldDBResponse
func (controller *Controller) DeleteSliceOfPointerToGongStructField(c *gin.Context) {

	mutexSliceOfPointerToGongStructField.Lock()
	defer mutexSliceOfPointerToGongStructField.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSliceOfPointerToGongStructField", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSliceOfPointerToGongStructField.GetDB()

	// Get model if exist
	var sliceofpointertogongstructfieldDB orm.SliceOfPointerToGongStructFieldDB
	if _, err := db.First(&sliceofpointertogongstructfieldDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&sliceofpointertogongstructfieldDB)

	// get an instance (not staged) from DB instance, and call callback function
	sliceofpointertogongstructfieldDeleted := new(models.SliceOfPointerToGongStructField)
	sliceofpointertogongstructfieldDB.CopyBasicFieldsToSliceOfPointerToGongStructField(sliceofpointertogongstructfieldDeleted)

	// get stage instance from DB instance, and call callback function
	sliceofpointertogongstructfieldStaged := backRepo.BackRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldPtr[sliceofpointertogongstructfieldDB.ID]
	if sliceofpointertogongstructfieldStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), sliceofpointertogongstructfieldStaged, sliceofpointertogongstructfieldDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
