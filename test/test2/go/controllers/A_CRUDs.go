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
var __A__dummysDeclaration__ models.A
var __A_time__dummyDeclaration time.Duration

var mutexA sync.Mutex

// An AID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA updateA deleteA
type AID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA updateA
type AInput struct {
	// The A to submit or modify
	// in: body
	A *orm.AAPI
}

// GetAs
//
// swagger:route GET /as as getAs
//
// # Get all as
//
// Responses:
// default: genericError
//
//	200: aDBResponse
func (controller *Controller) GetAs(c *gin.Context) {

	// source slice
	var aDBs []orm.ADB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/test2/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA.GetDB()

	query := db.Find(&aDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	aAPIs := make([]orm.AAPI, 0)

	// for each a, update fields from the database nullable fields
	for idx := range aDBs {
		aDB := &aDBs[idx]
		_ = aDB
		var aAPI orm.AAPI

		// insertion point for updating fields
		aAPI.ID = aDB.ID
		aDB.CopyBasicFieldsToA_WOP(&aAPI.A_WOP)
		aAPI.APointersEncoding = aDB.APointersEncoding
		aAPIs = append(aAPIs, aAPI)
	}

	c.JSON(http.StatusOK, aAPIs)
}

// PostA
//
// swagger:route POST /as as postA
//
// Creates a a
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA(c *gin.Context) {

	mutexA.Lock()
	defer mutexA.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/test2/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA.GetDB()

	// Validate input
	var input orm.AAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a
	aDB := orm.ADB{}
	aDB.APointersEncoding = input.APointersEncoding
	aDB.CopyBasicFieldsFromA_WOP(&input.A_WOP)

	query := db.Create(&aDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA.CheckoutPhaseOneInstance(&aDB)
	a := backRepo.BackRepoA.Map_ADBID_APtr[aDB.ID]

	if a != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, aDB)
}

// GetA
//
// swagger:route GET /as/{ID} as getA
//
// Gets the details for a a.
//
// Responses:
// default: genericError
//
//	200: aDBResponse
func (controller *Controller) GetA(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/test2/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA.GetDB()

	// Get aDB in DB
	var aDB orm.ADB
	if err := db.First(&aDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var aAPI orm.AAPI
	aAPI.ID = aDB.ID
	aAPI.APointersEncoding = aDB.APointersEncoding
	aDB.CopyBasicFieldsToA_WOP(&aAPI.A_WOP)

	c.JSON(http.StatusOK, aAPI)
}

// UpdateA
//
// swagger:route PATCH /as/{ID} as updateA
//
// # Update a a
//
// Responses:
// default: genericError
//
//	200: aDBResponse
func (controller *Controller) UpdateA(c *gin.Context) {

	mutexA.Lock()
	defer mutexA.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/test2/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA.GetDB()

	// Validate input
	var input orm.AAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var aDB orm.ADB

	// fetch the a
	query := db.First(&aDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	aDB.CopyBasicFieldsFromA_WOP(&input.A_WOP)
	aDB.APointersEncoding = input.APointersEncoding

	query = db.Model(&aDB).Updates(aDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	aNew := new(models.A)
	aDB.CopyBasicFieldsToA(aNew)

	// redeem pointers
	aDB.DecodePointers(backRepo, aNew)

	// get stage instance from DB instance, and call callback function
	aOld := backRepo.BackRepoA.Map_ADBID_APtr[aDB.ID]
	if aOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), aOld, aNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the aDB
	c.JSON(http.StatusOK, aDB)
}

// DeleteA
//
// swagger:route DELETE /as/{ID} as deleteA
//
// # Delete a a
//
// default: genericError
//
//	200: aDBResponse
func (controller *Controller) DeleteA(c *gin.Context) {

	mutexA.Lock()
	defer mutexA.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/test2/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA.GetDB()

	// Get model if exist
	var aDB orm.ADB
	if err := db.First(&aDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&aDB)

	// get an instance (not staged) from DB instance, and call callback function
	aDeleted := new(models.A)
	aDB.CopyBasicFieldsToA(aDeleted)

	// get stage instance from DB instance, and call callback function
	aStaged := backRepo.BackRepoA.Map_ADBID_APtr[aDB.ID]
	if aStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), aStaged, aDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
