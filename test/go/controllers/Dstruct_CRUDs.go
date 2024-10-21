// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/fullstack-lang/gong/test/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Dstruct__dummysDeclaration__ models.Dstruct
var __Dstruct_time__dummyDeclaration time.Duration

var mutexDstruct sync.Mutex

// An DstructID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDstruct updateDstruct deleteDstruct
type DstructID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DstructInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDstruct updateDstruct
type DstructInput struct {
	// The Dstruct to submit or modify
	// in: body
	Dstruct *orm.DstructAPI
}

// GetDstructs
//
// swagger:route GET /dstructs dstructs getDstructs
//
// # Get all dstructs
//
// Responses:
// default: genericError
//
//	200: dstructDBResponse
func (controller *Controller) GetDstructs(c *gin.Context) {

	// source slice
	var dstructDBs []orm.DstructDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDstructs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/test/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDstruct.GetDB()

	_, err := db.Find(&dstructDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	dstructAPIs := make([]orm.DstructAPI, 0)

	// for each dstruct, update fields from the database nullable fields
	for idx := range dstructDBs {
		dstructDB := &dstructDBs[idx]
		_ = dstructDB
		var dstructAPI orm.DstructAPI

		// insertion point for updating fields
		dstructAPI.ID = dstructDB.ID
		dstructDB.CopyBasicFieldsToDstruct_WOP(&dstructAPI.Dstruct_WOP)
		dstructAPI.DstructPointersEncoding = dstructDB.DstructPointersEncoding
		dstructAPIs = append(dstructAPIs, dstructAPI)
	}

	c.JSON(http.StatusOK, dstructAPIs)
}

// PostDstruct
//
// swagger:route POST /dstructs dstructs postDstruct
//
// Creates a dstruct
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDstruct(c *gin.Context) {

	mutexDstruct.Lock()
	defer mutexDstruct.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDstructs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/test/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDstruct.GetDB()

	// Validate input
	var input orm.DstructAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create dstruct
	dstructDB := orm.DstructDB{}
	dstructDB.DstructPointersEncoding = input.DstructPointersEncoding
	dstructDB.CopyBasicFieldsFromDstruct_WOP(&input.Dstruct_WOP)

	_, err = db.Create(&dstructDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDstruct.CheckoutPhaseOneInstance(&dstructDB)
	dstruct := backRepo.BackRepoDstruct.Map_DstructDBID_DstructPtr[dstructDB.ID]

	if dstruct != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), dstruct)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, dstructDB)
}

// GetDstruct
//
// swagger:route GET /dstructs/{ID} dstructs getDstruct
//
// Gets the details for a dstruct.
//
// Responses:
// default: genericError
//
//	200: dstructDBResponse
func (controller *Controller) GetDstruct(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDstruct", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/test/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDstruct.GetDB()

	// Get dstructDB in DB
	var dstructDB orm.DstructDB
	if _, err := db.First(&dstructDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var dstructAPI orm.DstructAPI
	dstructAPI.ID = dstructDB.ID
	dstructAPI.DstructPointersEncoding = dstructDB.DstructPointersEncoding
	dstructDB.CopyBasicFieldsToDstruct_WOP(&dstructAPI.Dstruct_WOP)

	c.JSON(http.StatusOK, dstructAPI)
}

// UpdateDstruct
//
// swagger:route PATCH /dstructs/{ID} dstructs updateDstruct
//
// # Update a dstruct
//
// Responses:
// default: genericError
//
//	200: dstructDBResponse
func (controller *Controller) UpdateDstruct(c *gin.Context) {

	mutexDstruct.Lock()
	defer mutexDstruct.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDstruct", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/test/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDstruct.GetDB()

	// Validate input
	var input orm.DstructAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var dstructDB orm.DstructDB

	// fetch the dstruct
	_, err := db.First(&dstructDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	dstructDB.CopyBasicFieldsFromDstruct_WOP(&input.Dstruct_WOP)
	dstructDB.DstructPointersEncoding = input.DstructPointersEncoding

	db, _ = db.Model(&dstructDB)
	_, err = db.Updates(dstructDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	dstructNew := new(models.Dstruct)
	dstructDB.CopyBasicFieldsToDstruct(dstructNew)

	// redeem pointers
	dstructDB.DecodePointers(backRepo, dstructNew)

	// get stage instance from DB instance, and call callback function
	dstructOld := backRepo.BackRepoDstruct.Map_DstructDBID_DstructPtr[dstructDB.ID]
	if dstructOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), dstructOld, dstructNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the dstructDB
	c.JSON(http.StatusOK, dstructDB)
}

// DeleteDstruct
//
// swagger:route DELETE /dstructs/{ID} dstructs deleteDstruct
//
// # Delete a dstruct
//
// default: genericError
//
//	200: dstructDBResponse
func (controller *Controller) DeleteDstruct(c *gin.Context) {

	mutexDstruct.Lock()
	defer mutexDstruct.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDstruct", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/test/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDstruct.GetDB()

	// Get model if exist
	var dstructDB orm.DstructDB
	if _, err := db.First(&dstructDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&dstructDB)

	// get an instance (not staged) from DB instance, and call callback function
	dstructDeleted := new(models.Dstruct)
	dstructDB.CopyBasicFieldsToDstruct(dstructDeleted)

	// get stage instance from DB instance, and call callback function
	dstructStaged := backRepo.BackRepoDstruct.Map_DstructDBID_DstructPtr[dstructDB.ID]
	if dstructStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), dstructStaged, dstructDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
