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
var __GongBasicField__dummysDeclaration__ models.GongBasicField
var __GongBasicField_time__dummyDeclaration time.Duration

var mutexGongBasicField sync.Mutex

// An GongBasicFieldID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGongBasicField updateGongBasicField deleteGongBasicField
type GongBasicFieldID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GongBasicFieldInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGongBasicField updateGongBasicField
type GongBasicFieldInput struct {
	// The GongBasicField to submit or modify
	// in: body
	GongBasicField *orm.GongBasicFieldAPI
}

// GetGongBasicFields
//
// swagger:route GET /gongbasicfields gongbasicfields getGongBasicFields
//
// # Get all gongbasicfields
//
// Responses:
// default: genericError
//
//	200: gongbasicfieldDBResponse
func (controller *Controller) GetGongBasicFields(c *gin.Context) {

	// source slice
	var gongbasicfieldDBs []orm.GongBasicFieldDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongBasicFields", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongBasicField.GetDB()

	_, err := db.Find(&gongbasicfieldDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	gongbasicfieldAPIs := make([]orm.GongBasicFieldAPI, 0)

	// for each gongbasicfield, update fields from the database nullable fields
	for idx := range gongbasicfieldDBs {
		gongbasicfieldDB := &gongbasicfieldDBs[idx]
		_ = gongbasicfieldDB
		var gongbasicfieldAPI orm.GongBasicFieldAPI

		// insertion point for updating fields
		gongbasicfieldAPI.ID = gongbasicfieldDB.ID
		gongbasicfieldDB.CopyBasicFieldsToGongBasicField_WOP(&gongbasicfieldAPI.GongBasicField_WOP)
		gongbasicfieldAPI.GongBasicFieldPointersEncoding = gongbasicfieldDB.GongBasicFieldPointersEncoding
		gongbasicfieldAPIs = append(gongbasicfieldAPIs, gongbasicfieldAPI)
	}

	c.JSON(http.StatusOK, gongbasicfieldAPIs)
}

// PostGongBasicField
//
// swagger:route POST /gongbasicfields gongbasicfields postGongBasicField
//
// Creates a gongbasicfield
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGongBasicField(c *gin.Context) {

	mutexGongBasicField.Lock()
	defer mutexGongBasicField.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGongBasicFields", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongBasicField.GetDB()

	// Validate input
	var input orm.GongBasicFieldAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gongbasicfield
	gongbasicfieldDB := orm.GongBasicFieldDB{}
	gongbasicfieldDB.GongBasicFieldPointersEncoding = input.GongBasicFieldPointersEncoding
	gongbasicfieldDB.CopyBasicFieldsFromGongBasicField_WOP(&input.GongBasicField_WOP)

	_, err = db.Create(&gongbasicfieldDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGongBasicField.CheckoutPhaseOneInstance(&gongbasicfieldDB)
	gongbasicfield := backRepo.BackRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldPtr[gongbasicfieldDB.ID]

	if gongbasicfield != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), gongbasicfield)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gongbasicfieldDB)
}

// GetGongBasicField
//
// swagger:route GET /gongbasicfields/{ID} gongbasicfields getGongBasicField
//
// Gets the details for a gongbasicfield.
//
// Responses:
// default: genericError
//
//	200: gongbasicfieldDBResponse
func (controller *Controller) GetGongBasicField(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongBasicField", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongBasicField.GetDB()

	// Get gongbasicfieldDB in DB
	var gongbasicfieldDB orm.GongBasicFieldDB
	if _, err := db.First(&gongbasicfieldDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var gongbasicfieldAPI orm.GongBasicFieldAPI
	gongbasicfieldAPI.ID = gongbasicfieldDB.ID
	gongbasicfieldAPI.GongBasicFieldPointersEncoding = gongbasicfieldDB.GongBasicFieldPointersEncoding
	gongbasicfieldDB.CopyBasicFieldsToGongBasicField_WOP(&gongbasicfieldAPI.GongBasicField_WOP)

	c.JSON(http.StatusOK, gongbasicfieldAPI)
}

// UpdateGongBasicField
//
// swagger:route PATCH /gongbasicfields/{ID} gongbasicfields updateGongBasicField
//
// # Update a gongbasicfield
//
// Responses:
// default: genericError
//
//	200: gongbasicfieldDBResponse
func (controller *Controller) UpdateGongBasicField(c *gin.Context) {

	mutexGongBasicField.Lock()
	defer mutexGongBasicField.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGongBasicField", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongBasicField.GetDB()

	// Validate input
	var input orm.GongBasicFieldAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var gongbasicfieldDB orm.GongBasicFieldDB

	// fetch the gongbasicfield
	_, err := db.First(&gongbasicfieldDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	gongbasicfieldDB.CopyBasicFieldsFromGongBasicField_WOP(&input.GongBasicField_WOP)
	gongbasicfieldDB.GongBasicFieldPointersEncoding = input.GongBasicFieldPointersEncoding

	db, _ = db.Model(&gongbasicfieldDB)
	_, err = db.Updates(&gongbasicfieldDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	gongbasicfieldNew := new(models.GongBasicField)
	gongbasicfieldDB.CopyBasicFieldsToGongBasicField(gongbasicfieldNew)

	// redeem pointers
	gongbasicfieldDB.DecodePointers(backRepo, gongbasicfieldNew)

	// get stage instance from DB instance, and call callback function
	gongbasicfieldOld := backRepo.BackRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldPtr[gongbasicfieldDB.ID]
	if gongbasicfieldOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), gongbasicfieldOld, gongbasicfieldNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gongbasicfieldDB
	c.JSON(http.StatusOK, gongbasicfieldDB)
}

// DeleteGongBasicField
//
// swagger:route DELETE /gongbasicfields/{ID} gongbasicfields deleteGongBasicField
//
// # Delete a gongbasicfield
//
// default: genericError
//
//	200: gongbasicfieldDBResponse
func (controller *Controller) DeleteGongBasicField(c *gin.Context) {

	mutexGongBasicField.Lock()
	defer mutexGongBasicField.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGongBasicField", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongBasicField.GetDB()

	// Get model if exist
	var gongbasicfieldDB orm.GongBasicFieldDB
	if _, err := db.First(&gongbasicfieldDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&gongbasicfieldDB)

	// get an instance (not staged) from DB instance, and call callback function
	gongbasicfieldDeleted := new(models.GongBasicField)
	gongbasicfieldDB.CopyBasicFieldsToGongBasicField(gongbasicfieldDeleted)

	// get stage instance from DB instance, and call callback function
	gongbasicfieldStaged := backRepo.BackRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldPtr[gongbasicfieldDB.ID]
	if gongbasicfieldStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), gongbasicfieldStaged, gongbasicfieldDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
