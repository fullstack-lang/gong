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
var __CellBoolean__dummysDeclaration__ models.CellBoolean
var __CellBoolean_time__dummyDeclaration time.Duration

var mutexCellBoolean sync.Mutex

// An CellBooleanID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCellBoolean updateCellBoolean deleteCellBoolean
type CellBooleanID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CellBooleanInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCellBoolean updateCellBoolean
type CellBooleanInput struct {
	// The CellBoolean to submit or modify
	// in: body
	CellBoolean *orm.CellBooleanAPI
}

// GetCellBooleans
//
// swagger:route GET /cellbooleans cellbooleans getCellBooleans
//
// # Get all cellbooleans
//
// Responses:
// default: genericError
//
//	200: cellbooleanDBResponse
func (controller *Controller) GetCellBooleans(c *gin.Context) {

	// source slice
	var cellbooleanDBs []orm.CellBooleanDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCellBooleans", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellBoolean.GetDB()

	_, err := db.Find(&cellbooleanDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	cellbooleanAPIs := make([]orm.CellBooleanAPI, 0)

	// for each cellboolean, update fields from the database nullable fields
	for idx := range cellbooleanDBs {
		cellbooleanDB := &cellbooleanDBs[idx]
		_ = cellbooleanDB
		var cellbooleanAPI orm.CellBooleanAPI

		// insertion point for updating fields
		cellbooleanAPI.ID = cellbooleanDB.ID
		cellbooleanDB.CopyBasicFieldsToCellBoolean_WOP(&cellbooleanAPI.CellBoolean_WOP)
		cellbooleanAPI.CellBooleanPointersEncoding = cellbooleanDB.CellBooleanPointersEncoding
		cellbooleanAPIs = append(cellbooleanAPIs, cellbooleanAPI)
	}

	c.JSON(http.StatusOK, cellbooleanAPIs)
}

// PostCellBoolean
//
// swagger:route POST /cellbooleans cellbooleans postCellBoolean
//
// Creates a cellboolean
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCellBoolean(c *gin.Context) {

	mutexCellBoolean.Lock()
	defer mutexCellBoolean.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCellBooleans", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellBoolean.GetDB()

	// Validate input
	var input orm.CellBooleanAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create cellboolean
	cellbooleanDB := orm.CellBooleanDB{}
	cellbooleanDB.CellBooleanPointersEncoding = input.CellBooleanPointersEncoding
	cellbooleanDB.CopyBasicFieldsFromCellBoolean_WOP(&input.CellBoolean_WOP)

	_, err = db.Create(&cellbooleanDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCellBoolean.CheckoutPhaseOneInstance(&cellbooleanDB)
	cellboolean := backRepo.BackRepoCellBoolean.Map_CellBooleanDBID_CellBooleanPtr[cellbooleanDB.ID]

	if cellboolean != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), cellboolean)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, cellbooleanDB)
}

// GetCellBoolean
//
// swagger:route GET /cellbooleans/{ID} cellbooleans getCellBoolean
//
// Gets the details for a cellboolean.
//
// Responses:
// default: genericError
//
//	200: cellbooleanDBResponse
func (controller *Controller) GetCellBoolean(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCellBoolean", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellBoolean.GetDB()

	// Get cellbooleanDB in DB
	var cellbooleanDB orm.CellBooleanDB
	if _, err := db.First(&cellbooleanDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var cellbooleanAPI orm.CellBooleanAPI
	cellbooleanAPI.ID = cellbooleanDB.ID
	cellbooleanAPI.CellBooleanPointersEncoding = cellbooleanDB.CellBooleanPointersEncoding
	cellbooleanDB.CopyBasicFieldsToCellBoolean_WOP(&cellbooleanAPI.CellBoolean_WOP)

	c.JSON(http.StatusOK, cellbooleanAPI)
}

// UpdateCellBoolean
//
// swagger:route PATCH /cellbooleans/{ID} cellbooleans updateCellBoolean
//
// # Update a cellboolean
//
// Responses:
// default: genericError
//
//	200: cellbooleanDBResponse
func (controller *Controller) UpdateCellBoolean(c *gin.Context) {

	mutexCellBoolean.Lock()
	defer mutexCellBoolean.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCellBoolean", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellBoolean.GetDB()

	// Validate input
	var input orm.CellBooleanAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var cellbooleanDB orm.CellBooleanDB

	// fetch the cellboolean
	_, err := db.First(&cellbooleanDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	cellbooleanDB.CopyBasicFieldsFromCellBoolean_WOP(&input.CellBoolean_WOP)
	cellbooleanDB.CellBooleanPointersEncoding = input.CellBooleanPointersEncoding

	db, _ = db.Model(&cellbooleanDB)
	_, err = db.Updates(&cellbooleanDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	cellbooleanNew := new(models.CellBoolean)
	cellbooleanDB.CopyBasicFieldsToCellBoolean(cellbooleanNew)

	// redeem pointers
	cellbooleanDB.DecodePointers(backRepo, cellbooleanNew)

	// get stage instance from DB instance, and call callback function
	cellbooleanOld := backRepo.BackRepoCellBoolean.Map_CellBooleanDBID_CellBooleanPtr[cellbooleanDB.ID]
	if cellbooleanOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), cellbooleanOld, cellbooleanNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the cellbooleanDB
	c.JSON(http.StatusOK, cellbooleanDB)
}

// DeleteCellBoolean
//
// swagger:route DELETE /cellbooleans/{ID} cellbooleans deleteCellBoolean
//
// # Delete a cellboolean
//
// default: genericError
//
//	200: cellbooleanDBResponse
func (controller *Controller) DeleteCellBoolean(c *gin.Context) {

	mutexCellBoolean.Lock()
	defer mutexCellBoolean.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCellBoolean", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellBoolean.GetDB()

	// Get model if exist
	var cellbooleanDB orm.CellBooleanDB
	if _, err := db.First(&cellbooleanDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&cellbooleanDB)

	// get an instance (not staged) from DB instance, and call callback function
	cellbooleanDeleted := new(models.CellBoolean)
	cellbooleanDB.CopyBasicFieldsToCellBoolean(cellbooleanDeleted)

	// get stage instance from DB instance, and call callback function
	cellbooleanStaged := backRepo.BackRepoCellBoolean.Map_CellBooleanDBID_CellBooleanPtr[cellbooleanDB.ID]
	if cellbooleanStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), cellbooleanStaged, cellbooleanDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
