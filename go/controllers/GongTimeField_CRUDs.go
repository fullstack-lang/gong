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
var __GongTimeField__dummysDeclaration__ models.GongTimeField
var __GongTimeField_time__dummyDeclaration time.Duration

var mutexGongTimeField sync.Mutex

// An GongTimeFieldID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGongTimeField updateGongTimeField deleteGongTimeField
type GongTimeFieldID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GongTimeFieldInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGongTimeField updateGongTimeField
type GongTimeFieldInput struct {
	// The GongTimeField to submit or modify
	// in: body
	GongTimeField *orm.GongTimeFieldAPI
}

// GetGongTimeFields
//
// swagger:route GET /gongtimefields gongtimefields getGongTimeFields
//
// # Get all gongtimefields
//
// Responses:
// default: genericError
//
//	200: gongtimefieldDBResponse
func (controller *Controller) GetGongTimeFields(c *gin.Context) {

	// source slice
	var gongtimefieldDBs []orm.GongTimeFieldDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongTimeFields", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongTimeField.GetDB()

	_, err := db.Find(&gongtimefieldDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	gongtimefieldAPIs := make([]orm.GongTimeFieldAPI, 0)

	// for each gongtimefield, update fields from the database nullable fields
	for idx := range gongtimefieldDBs {
		gongtimefieldDB := &gongtimefieldDBs[idx]
		_ = gongtimefieldDB
		var gongtimefieldAPI orm.GongTimeFieldAPI

		// insertion point for updating fields
		gongtimefieldAPI.ID = gongtimefieldDB.ID
		gongtimefieldDB.CopyBasicFieldsToGongTimeField_WOP(&gongtimefieldAPI.GongTimeField_WOP)
		gongtimefieldAPI.GongTimeFieldPointersEncoding = gongtimefieldDB.GongTimeFieldPointersEncoding
		gongtimefieldAPIs = append(gongtimefieldAPIs, gongtimefieldAPI)
	}

	c.JSON(http.StatusOK, gongtimefieldAPIs)
}

// PostGongTimeField
//
// swagger:route POST /gongtimefields gongtimefields postGongTimeField
//
// Creates a gongtimefield
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGongTimeField(c *gin.Context) {

	mutexGongTimeField.Lock()
	defer mutexGongTimeField.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGongTimeFields", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongTimeField.GetDB()

	// Validate input
	var input orm.GongTimeFieldAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gongtimefield
	gongtimefieldDB := orm.GongTimeFieldDB{}
	gongtimefieldDB.GongTimeFieldPointersEncoding = input.GongTimeFieldPointersEncoding
	gongtimefieldDB.CopyBasicFieldsFromGongTimeField_WOP(&input.GongTimeField_WOP)

	_, err = db.Create(&gongtimefieldDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGongTimeField.CheckoutPhaseOneInstance(&gongtimefieldDB)
	gongtimefield := backRepo.BackRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldPtr[gongtimefieldDB.ID]

	if gongtimefield != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), gongtimefield)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gongtimefieldDB)
}

// GetGongTimeField
//
// swagger:route GET /gongtimefields/{ID} gongtimefields getGongTimeField
//
// Gets the details for a gongtimefield.
//
// Responses:
// default: genericError
//
//	200: gongtimefieldDBResponse
func (controller *Controller) GetGongTimeField(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongTimeField", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongTimeField.GetDB()

	// Get gongtimefieldDB in DB
	var gongtimefieldDB orm.GongTimeFieldDB
	if _, err := db.First(&gongtimefieldDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var gongtimefieldAPI orm.GongTimeFieldAPI
	gongtimefieldAPI.ID = gongtimefieldDB.ID
	gongtimefieldAPI.GongTimeFieldPointersEncoding = gongtimefieldDB.GongTimeFieldPointersEncoding
	gongtimefieldDB.CopyBasicFieldsToGongTimeField_WOP(&gongtimefieldAPI.GongTimeField_WOP)

	c.JSON(http.StatusOK, gongtimefieldAPI)
}

// UpdateGongTimeField
//
// swagger:route PATCH /gongtimefields/{ID} gongtimefields updateGongTimeField
//
// # Update a gongtimefield
//
// Responses:
// default: genericError
//
//	200: gongtimefieldDBResponse
func (controller *Controller) UpdateGongTimeField(c *gin.Context) {

	mutexGongTimeField.Lock()
	defer mutexGongTimeField.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGongTimeField", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongTimeField.GetDB()

	// Validate input
	var input orm.GongTimeFieldAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var gongtimefieldDB orm.GongTimeFieldDB

	// fetch the gongtimefield
	_, err := db.First(&gongtimefieldDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	gongtimefieldDB.CopyBasicFieldsFromGongTimeField_WOP(&input.GongTimeField_WOP)
	gongtimefieldDB.GongTimeFieldPointersEncoding = input.GongTimeFieldPointersEncoding

	db, _ = db.Model(&gongtimefieldDB)
	_, err = db.Updates(&gongtimefieldDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	gongtimefieldNew := new(models.GongTimeField)
	gongtimefieldDB.CopyBasicFieldsToGongTimeField(gongtimefieldNew)

	// redeem pointers
	gongtimefieldDB.DecodePointers(backRepo, gongtimefieldNew)

	// get stage instance from DB instance, and call callback function
	gongtimefieldOld := backRepo.BackRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldPtr[gongtimefieldDB.ID]
	if gongtimefieldOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), gongtimefieldOld, gongtimefieldNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gongtimefieldDB
	c.JSON(http.StatusOK, gongtimefieldDB)
}

// DeleteGongTimeField
//
// swagger:route DELETE /gongtimefields/{ID} gongtimefields deleteGongTimeField
//
// # Delete a gongtimefield
//
// default: genericError
//
//	200: gongtimefieldDBResponse
func (controller *Controller) DeleteGongTimeField(c *gin.Context) {

	mutexGongTimeField.Lock()
	defer mutexGongTimeField.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGongTimeField", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongTimeField.GetDB()

	// Get model if exist
	var gongtimefieldDB orm.GongTimeFieldDB
	if _, err := db.First(&gongtimefieldDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&gongtimefieldDB)

	// get an instance (not staged) from DB instance, and call callback function
	gongtimefieldDeleted := new(models.GongTimeField)
	gongtimefieldDB.CopyBasicFieldsToGongTimeField(gongtimefieldDeleted)

	// get stage instance from DB instance, and call callback function
	gongtimefieldStaged := backRepo.BackRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldPtr[gongtimefieldDB.ID]
	if gongtimefieldStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), gongtimefieldStaged, gongtimefieldDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
