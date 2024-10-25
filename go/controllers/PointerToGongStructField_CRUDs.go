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
var __PointerToGongStructField__dummysDeclaration__ models.PointerToGongStructField
var __PointerToGongStructField_time__dummyDeclaration time.Duration

var mutexPointerToGongStructField sync.Mutex

// An PointerToGongStructFieldID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPointerToGongStructField updatePointerToGongStructField deletePointerToGongStructField
type PointerToGongStructFieldID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PointerToGongStructFieldInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPointerToGongStructField updatePointerToGongStructField
type PointerToGongStructFieldInput struct {
	// The PointerToGongStructField to submit or modify
	// in: body
	PointerToGongStructField *orm.PointerToGongStructFieldAPI
}

// GetPointerToGongStructFields
//
// swagger:route GET /pointertogongstructfields pointertogongstructfields getPointerToGongStructFields
//
// # Get all pointertogongstructfields
//
// Responses:
// default: genericError
//
//	200: pointertogongstructfieldDBResponse
func (controller *Controller) GetPointerToGongStructFields(c *gin.Context) {

	// source slice
	var pointertogongstructfieldDBs []orm.PointerToGongStructFieldDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPointerToGongStructFields", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPointerToGongStructField.GetDB()

	_, err := db.Find(&pointertogongstructfieldDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	pointertogongstructfieldAPIs := make([]orm.PointerToGongStructFieldAPI, 0)

	// for each pointertogongstructfield, update fields from the database nullable fields
	for idx := range pointertogongstructfieldDBs {
		pointertogongstructfieldDB := &pointertogongstructfieldDBs[idx]
		_ = pointertogongstructfieldDB
		var pointertogongstructfieldAPI orm.PointerToGongStructFieldAPI

		// insertion point for updating fields
		pointertogongstructfieldAPI.ID = pointertogongstructfieldDB.ID
		pointertogongstructfieldDB.CopyBasicFieldsToPointerToGongStructField_WOP(&pointertogongstructfieldAPI.PointerToGongStructField_WOP)
		pointertogongstructfieldAPI.PointerToGongStructFieldPointersEncoding = pointertogongstructfieldDB.PointerToGongStructFieldPointersEncoding
		pointertogongstructfieldAPIs = append(pointertogongstructfieldAPIs, pointertogongstructfieldAPI)
	}

	c.JSON(http.StatusOK, pointertogongstructfieldAPIs)
}

// PostPointerToGongStructField
//
// swagger:route POST /pointertogongstructfields pointertogongstructfields postPointerToGongStructField
//
// Creates a pointertogongstructfield
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPointerToGongStructField(c *gin.Context) {

	mutexPointerToGongStructField.Lock()
	defer mutexPointerToGongStructField.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPointerToGongStructFields", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPointerToGongStructField.GetDB()

	// Validate input
	var input orm.PointerToGongStructFieldAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create pointertogongstructfield
	pointertogongstructfieldDB := orm.PointerToGongStructFieldDB{}
	pointertogongstructfieldDB.PointerToGongStructFieldPointersEncoding = input.PointerToGongStructFieldPointersEncoding
	pointertogongstructfieldDB.CopyBasicFieldsFromPointerToGongStructField_WOP(&input.PointerToGongStructField_WOP)

	_, err = db.Create(&pointertogongstructfieldDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPointerToGongStructField.CheckoutPhaseOneInstance(&pointertogongstructfieldDB)
	pointertogongstructfield := backRepo.BackRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldPtr[pointertogongstructfieldDB.ID]

	if pointertogongstructfield != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), pointertogongstructfield)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, pointertogongstructfieldDB)
}

// GetPointerToGongStructField
//
// swagger:route GET /pointertogongstructfields/{ID} pointertogongstructfields getPointerToGongStructField
//
// Gets the details for a pointertogongstructfield.
//
// Responses:
// default: genericError
//
//	200: pointertogongstructfieldDBResponse
func (controller *Controller) GetPointerToGongStructField(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPointerToGongStructField", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPointerToGongStructField.GetDB()

	// Get pointertogongstructfieldDB in DB
	var pointertogongstructfieldDB orm.PointerToGongStructFieldDB
	if _, err := db.First(&pointertogongstructfieldDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var pointertogongstructfieldAPI orm.PointerToGongStructFieldAPI
	pointertogongstructfieldAPI.ID = pointertogongstructfieldDB.ID
	pointertogongstructfieldAPI.PointerToGongStructFieldPointersEncoding = pointertogongstructfieldDB.PointerToGongStructFieldPointersEncoding
	pointertogongstructfieldDB.CopyBasicFieldsToPointerToGongStructField_WOP(&pointertogongstructfieldAPI.PointerToGongStructField_WOP)

	c.JSON(http.StatusOK, pointertogongstructfieldAPI)
}

// UpdatePointerToGongStructField
//
// swagger:route PATCH /pointertogongstructfields/{ID} pointertogongstructfields updatePointerToGongStructField
//
// # Update a pointertogongstructfield
//
// Responses:
// default: genericError
//
//	200: pointertogongstructfieldDBResponse
func (controller *Controller) UpdatePointerToGongStructField(c *gin.Context) {

	mutexPointerToGongStructField.Lock()
	defer mutexPointerToGongStructField.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePointerToGongStructField", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPointerToGongStructField.GetDB()

	// Validate input
	var input orm.PointerToGongStructFieldAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var pointertogongstructfieldDB orm.PointerToGongStructFieldDB

	// fetch the pointertogongstructfield
	_, err := db.First(&pointertogongstructfieldDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	pointertogongstructfieldDB.CopyBasicFieldsFromPointerToGongStructField_WOP(&input.PointerToGongStructField_WOP)
	pointertogongstructfieldDB.PointerToGongStructFieldPointersEncoding = input.PointerToGongStructFieldPointersEncoding

	db, _ = db.Model(&pointertogongstructfieldDB)
	_, err = db.Updates(&pointertogongstructfieldDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	pointertogongstructfieldNew := new(models.PointerToGongStructField)
	pointertogongstructfieldDB.CopyBasicFieldsToPointerToGongStructField(pointertogongstructfieldNew)

	// redeem pointers
	pointertogongstructfieldDB.DecodePointers(backRepo, pointertogongstructfieldNew)

	// get stage instance from DB instance, and call callback function
	pointertogongstructfieldOld := backRepo.BackRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldPtr[pointertogongstructfieldDB.ID]
	if pointertogongstructfieldOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), pointertogongstructfieldOld, pointertogongstructfieldNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the pointertogongstructfieldDB
	c.JSON(http.StatusOK, pointertogongstructfieldDB)
}

// DeletePointerToGongStructField
//
// swagger:route DELETE /pointertogongstructfields/{ID} pointertogongstructfields deletePointerToGongStructField
//
// # Delete a pointertogongstructfield
//
// default: genericError
//
//	200: pointertogongstructfieldDBResponse
func (controller *Controller) DeletePointerToGongStructField(c *gin.Context) {

	mutexPointerToGongStructField.Lock()
	defer mutexPointerToGongStructField.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePointerToGongStructField", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPointerToGongStructField.GetDB()

	// Get model if exist
	var pointertogongstructfieldDB orm.PointerToGongStructFieldDB
	if _, err := db.First(&pointertogongstructfieldDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&pointertogongstructfieldDB)

	// get an instance (not staged) from DB instance, and call callback function
	pointertogongstructfieldDeleted := new(models.PointerToGongStructField)
	pointertogongstructfieldDB.CopyBasicFieldsToPointerToGongStructField(pointertogongstructfieldDeleted)

	// get stage instance from DB instance, and call callback function
	pointertogongstructfieldStaged := backRepo.BackRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldPtr[pointertogongstructfieldDB.ID]
	if pointertogongstructfieldStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), pointertogongstructfieldStaged, pointertogongstructfieldDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
