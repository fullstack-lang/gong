// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/test2/go/models"
	"github.com/fullstack-lang/gong/test2/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Dummy__dummysDeclaration__ models.Dummy
var __Dummy_time__dummyDeclaration time.Duration

var mutexDummy sync.Mutex

// An DummyID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDummy updateDummy deleteDummy
type DummyID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DummyInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDummy updateDummy
type DummyInput struct {
	// The Dummy to submit or modify
	// in: body
	Dummy *orm.DummyAPI
}

// GetDummys
//
// swagger:route GET /dummys dummys getDummys
//
// # Get all dummys
//
// Responses:
// default: genericError
//
//	200: dummyDBResponse
func (controller *Controller) GetDummys(c *gin.Context) {

	// source slice
	var dummyDBs []orm.DummyDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDummys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoDummy.GetDB()

	query := db.Find(&dummyDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	dummyAPIs := make([]orm.DummyAPI, 0)

	// for each dummy, update fields from the database nullable fields
	for idx := range dummyDBs {
		dummyDB := &dummyDBs[idx]
		_ = dummyDB
		var dummyAPI orm.DummyAPI

		// insertion point for updating fields
		dummyAPI.ID = dummyDB.ID
		dummyDB.CopyBasicFieldsToDummy(&dummyAPI.Dummy)
		dummyAPI.DummyPointersEnconding = dummyDB.DummyPointersEnconding
		dummyAPIs = append(dummyAPIs, dummyAPI)
	}

	c.JSON(http.StatusOK, dummyAPIs)
}

// PostDummy
//
// swagger:route POST /dummys dummys postDummy
//
// Creates a dummy
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDummy(c *gin.Context) {

	mutexDummy.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDummys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoDummy.GetDB()

	// Validate input
	var input orm.DummyAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create dummy
	dummyDB := orm.DummyDB{}
	dummyDB.DummyPointersEnconding = input.DummyPointersEnconding
	dummyDB.CopyBasicFieldsFromDummy(&input.Dummy)

	query := db.Create(&dummyDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDummy.CheckoutPhaseOneInstance(&dummyDB)
	dummy := backRepo.BackRepoDummy.Map_DummyDBID_DummyPtr[dummyDB.ID]

	if dummy != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), dummy)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, dummyDB)

	mutexDummy.Unlock()
}

// GetDummy
//
// swagger:route GET /dummys/{ID} dummys getDummy
//
// Gets the details for a dummy.
//
// Responses:
// default: genericError
//
//	200: dummyDBResponse
func (controller *Controller) GetDummy(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDummy", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoDummy.GetDB()

	// Get dummyDB in DB
	var dummyDB orm.DummyDB
	if err := db.First(&dummyDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var dummyAPI orm.DummyAPI
	dummyAPI.ID = dummyDB.ID
	dummyAPI.DummyPointersEnconding = dummyDB.DummyPointersEnconding
	dummyDB.CopyBasicFieldsToDummy(&dummyAPI.Dummy)

	c.JSON(http.StatusOK, dummyAPI)
}

// UpdateDummy
//
// swagger:route PATCH /dummys/{ID} dummys updateDummy
//
// # Update a dummy
//
// Responses:
// default: genericError
//
//	200: dummyDBResponse
func (controller *Controller) UpdateDummy(c *gin.Context) {

	mutexDummy.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDummy", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoDummy.GetDB()

	// Validate input
	var input orm.DummyAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var dummyDB orm.DummyDB

	// fetch the dummy
	query := db.First(&dummyDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	dummyDB.CopyBasicFieldsFromDummy(&input.Dummy)
	dummyDB.DummyPointersEnconding = input.DummyPointersEnconding

	query = db.Model(&dummyDB).Updates(dummyDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	dummyNew := new(models.Dummy)
	dummyDB.CopyBasicFieldsToDummy(dummyNew)

	// get stage instance from DB instance, and call callback function
	dummyOld := backRepo.BackRepoDummy.Map_DummyDBID_DummyPtr[dummyDB.ID]
	if dummyOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), dummyOld, dummyNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the dummyDB
	c.JSON(http.StatusOK, dummyDB)

	mutexDummy.Unlock()
}

// DeleteDummy
//
// swagger:route DELETE /dummys/{ID} dummys deleteDummy
//
// # Delete a dummy
//
// default: genericError
//
//	200: dummyDBResponse
func (controller *Controller) DeleteDummy(c *gin.Context) {

	mutexDummy.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDummy", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoDummy.GetDB()

	// Get model if exist
	var dummyDB orm.DummyDB
	if err := db.First(&dummyDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&dummyDB)

	// get an instance (not staged) from DB instance, and call callback function
	dummyDeleted := new(models.Dummy)
	dummyDB.CopyBasicFieldsToDummy(dummyDeleted)

	// get stage instance from DB instance, and call callback function
	dummyStaged := backRepo.BackRepoDummy.Map_DummyDBID_DummyPtr[dummyDB.ID]
	if dummyStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), dummyStaged, dummyDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexDummy.Unlock()
}
