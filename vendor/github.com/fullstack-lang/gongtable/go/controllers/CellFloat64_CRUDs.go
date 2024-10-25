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
var __CellFloat64__dummysDeclaration__ models.CellFloat64
var __CellFloat64_time__dummyDeclaration time.Duration

var mutexCellFloat64 sync.Mutex

// An CellFloat64ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCellFloat64 updateCellFloat64 deleteCellFloat64
type CellFloat64ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CellFloat64Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCellFloat64 updateCellFloat64
type CellFloat64Input struct {
	// The CellFloat64 to submit or modify
	// in: body
	CellFloat64 *orm.CellFloat64API
}

// GetCellFloat64s
//
// swagger:route GET /cellfloat64s cellfloat64s getCellFloat64s
//
// # Get all cellfloat64s
//
// Responses:
// default: genericError
//
//	200: cellfloat64DBResponse
func (controller *Controller) GetCellFloat64s(c *gin.Context) {

	// source slice
	var cellfloat64DBs []orm.CellFloat64DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCellFloat64s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellFloat64.GetDB()

	_, err := db.Find(&cellfloat64DBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	cellfloat64APIs := make([]orm.CellFloat64API, 0)

	// for each cellfloat64, update fields from the database nullable fields
	for idx := range cellfloat64DBs {
		cellfloat64DB := &cellfloat64DBs[idx]
		_ = cellfloat64DB
		var cellfloat64API orm.CellFloat64API

		// insertion point for updating fields
		cellfloat64API.ID = cellfloat64DB.ID
		cellfloat64DB.CopyBasicFieldsToCellFloat64_WOP(&cellfloat64API.CellFloat64_WOP)
		cellfloat64API.CellFloat64PointersEncoding = cellfloat64DB.CellFloat64PointersEncoding
		cellfloat64APIs = append(cellfloat64APIs, cellfloat64API)
	}

	c.JSON(http.StatusOK, cellfloat64APIs)
}

// PostCellFloat64
//
// swagger:route POST /cellfloat64s cellfloat64s postCellFloat64
//
// Creates a cellfloat64
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCellFloat64(c *gin.Context) {

	mutexCellFloat64.Lock()
	defer mutexCellFloat64.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCellFloat64s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellFloat64.GetDB()

	// Validate input
	var input orm.CellFloat64API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create cellfloat64
	cellfloat64DB := orm.CellFloat64DB{}
	cellfloat64DB.CellFloat64PointersEncoding = input.CellFloat64PointersEncoding
	cellfloat64DB.CopyBasicFieldsFromCellFloat64_WOP(&input.CellFloat64_WOP)

	_, err = db.Create(&cellfloat64DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCellFloat64.CheckoutPhaseOneInstance(&cellfloat64DB)
	cellfloat64 := backRepo.BackRepoCellFloat64.Map_CellFloat64DBID_CellFloat64Ptr[cellfloat64DB.ID]

	if cellfloat64 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), cellfloat64)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, cellfloat64DB)
}

// GetCellFloat64
//
// swagger:route GET /cellfloat64s/{ID} cellfloat64s getCellFloat64
//
// Gets the details for a cellfloat64.
//
// Responses:
// default: genericError
//
//	200: cellfloat64DBResponse
func (controller *Controller) GetCellFloat64(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCellFloat64", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellFloat64.GetDB()

	// Get cellfloat64DB in DB
	var cellfloat64DB orm.CellFloat64DB
	if _, err := db.First(&cellfloat64DB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var cellfloat64API orm.CellFloat64API
	cellfloat64API.ID = cellfloat64DB.ID
	cellfloat64API.CellFloat64PointersEncoding = cellfloat64DB.CellFloat64PointersEncoding
	cellfloat64DB.CopyBasicFieldsToCellFloat64_WOP(&cellfloat64API.CellFloat64_WOP)

	c.JSON(http.StatusOK, cellfloat64API)
}

// UpdateCellFloat64
//
// swagger:route PATCH /cellfloat64s/{ID} cellfloat64s updateCellFloat64
//
// # Update a cellfloat64
//
// Responses:
// default: genericError
//
//	200: cellfloat64DBResponse
func (controller *Controller) UpdateCellFloat64(c *gin.Context) {

	mutexCellFloat64.Lock()
	defer mutexCellFloat64.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCellFloat64", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellFloat64.GetDB()

	// Validate input
	var input orm.CellFloat64API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var cellfloat64DB orm.CellFloat64DB

	// fetch the cellfloat64
	_, err := db.First(&cellfloat64DB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	cellfloat64DB.CopyBasicFieldsFromCellFloat64_WOP(&input.CellFloat64_WOP)
	cellfloat64DB.CellFloat64PointersEncoding = input.CellFloat64PointersEncoding

	db, _ = db.Model(&cellfloat64DB)
	_, err = db.Updates(&cellfloat64DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	cellfloat64New := new(models.CellFloat64)
	cellfloat64DB.CopyBasicFieldsToCellFloat64(cellfloat64New)

	// redeem pointers
	cellfloat64DB.DecodePointers(backRepo, cellfloat64New)

	// get stage instance from DB instance, and call callback function
	cellfloat64Old := backRepo.BackRepoCellFloat64.Map_CellFloat64DBID_CellFloat64Ptr[cellfloat64DB.ID]
	if cellfloat64Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), cellfloat64Old, cellfloat64New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the cellfloat64DB
	c.JSON(http.StatusOK, cellfloat64DB)
}

// DeleteCellFloat64
//
// swagger:route DELETE /cellfloat64s/{ID} cellfloat64s deleteCellFloat64
//
// # Delete a cellfloat64
//
// default: genericError
//
//	200: cellfloat64DBResponse
func (controller *Controller) DeleteCellFloat64(c *gin.Context) {

	mutexCellFloat64.Lock()
	defer mutexCellFloat64.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCellFloat64", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellFloat64.GetDB()

	// Get model if exist
	var cellfloat64DB orm.CellFloat64DB
	if _, err := db.First(&cellfloat64DB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&cellfloat64DB)

	// get an instance (not staged) from DB instance, and call callback function
	cellfloat64Deleted := new(models.CellFloat64)
	cellfloat64DB.CopyBasicFieldsToCellFloat64(cellfloat64Deleted)

	// get stage instance from DB instance, and call callback function
	cellfloat64Staged := backRepo.BackRepoCellFloat64.Map_CellFloat64DBID_CellFloat64Ptr[cellfloat64DB.ID]
	if cellfloat64Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), cellfloat64Staged, cellfloat64Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
