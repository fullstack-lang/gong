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
var __Cell__dummysDeclaration__ models.Cell
var __Cell_time__dummyDeclaration time.Duration

var mutexCell sync.Mutex

// An CellID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCell updateCell deleteCell
type CellID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CellInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCell updateCell
type CellInput struct {
	// The Cell to submit or modify
	// in: body
	Cell *orm.CellAPI
}

// GetCells
//
// swagger:route GET /cells cells getCells
//
// # Get all cells
//
// Responses:
// default: genericError
//
//	200: cellDBResponse
func (controller *Controller) GetCells(c *gin.Context) {

	// source slice
	var cellDBs []orm.CellDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCells", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCell.GetDB()

	_, err := db.Find(&cellDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	cellAPIs := make([]orm.CellAPI, 0)

	// for each cell, update fields from the database nullable fields
	for idx := range cellDBs {
		cellDB := &cellDBs[idx]
		_ = cellDB
		var cellAPI orm.CellAPI

		// insertion point for updating fields
		cellAPI.ID = cellDB.ID
		cellDB.CopyBasicFieldsToCell_WOP(&cellAPI.Cell_WOP)
		cellAPI.CellPointersEncoding = cellDB.CellPointersEncoding
		cellAPIs = append(cellAPIs, cellAPI)
	}

	c.JSON(http.StatusOK, cellAPIs)
}

// PostCell
//
// swagger:route POST /cells cells postCell
//
// Creates a cell
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCell(c *gin.Context) {

	mutexCell.Lock()
	defer mutexCell.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCells", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCell.GetDB()

	// Validate input
	var input orm.CellAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create cell
	cellDB := orm.CellDB{}
	cellDB.CellPointersEncoding = input.CellPointersEncoding
	cellDB.CopyBasicFieldsFromCell_WOP(&input.Cell_WOP)

	_, err = db.Create(&cellDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCell.CheckoutPhaseOneInstance(&cellDB)
	cell := backRepo.BackRepoCell.Map_CellDBID_CellPtr[cellDB.ID]

	if cell != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), cell)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, cellDB)
}

// GetCell
//
// swagger:route GET /cells/{ID} cells getCell
//
// Gets the details for a cell.
//
// Responses:
// default: genericError
//
//	200: cellDBResponse
func (controller *Controller) GetCell(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCell", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCell.GetDB()

	// Get cellDB in DB
	var cellDB orm.CellDB
	if _, err := db.First(&cellDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var cellAPI orm.CellAPI
	cellAPI.ID = cellDB.ID
	cellAPI.CellPointersEncoding = cellDB.CellPointersEncoding
	cellDB.CopyBasicFieldsToCell_WOP(&cellAPI.Cell_WOP)

	c.JSON(http.StatusOK, cellAPI)
}

// UpdateCell
//
// swagger:route PATCH /cells/{ID} cells updateCell
//
// # Update a cell
//
// Responses:
// default: genericError
//
//	200: cellDBResponse
func (controller *Controller) UpdateCell(c *gin.Context) {

	mutexCell.Lock()
	defer mutexCell.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCell", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCell.GetDB()

	// Validate input
	var input orm.CellAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var cellDB orm.CellDB

	// fetch the cell
	_, err := db.First(&cellDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	cellDB.CopyBasicFieldsFromCell_WOP(&input.Cell_WOP)
	cellDB.CellPointersEncoding = input.CellPointersEncoding

	db, _ = db.Model(&cellDB)
	_, err = db.Updates(&cellDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	cellNew := new(models.Cell)
	cellDB.CopyBasicFieldsToCell(cellNew)

	// redeem pointers
	cellDB.DecodePointers(backRepo, cellNew)

	// get stage instance from DB instance, and call callback function
	cellOld := backRepo.BackRepoCell.Map_CellDBID_CellPtr[cellDB.ID]
	if cellOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), cellOld, cellNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the cellDB
	c.JSON(http.StatusOK, cellDB)
}

// DeleteCell
//
// swagger:route DELETE /cells/{ID} cells deleteCell
//
// # Delete a cell
//
// default: genericError
//
//	200: cellDBResponse
func (controller *Controller) DeleteCell(c *gin.Context) {

	mutexCell.Lock()
	defer mutexCell.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCell", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCell.GetDB()

	// Get model if exist
	var cellDB orm.CellDB
	if _, err := db.First(&cellDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&cellDB)

	// get an instance (not staged) from DB instance, and call callback function
	cellDeleted := new(models.Cell)
	cellDB.CopyBasicFieldsToCell(cellDeleted)

	// get stage instance from DB instance, and call callback function
	cellStaged := backRepo.BackRepoCell.Map_CellDBID_CellPtr[cellDB.ID]
	if cellStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), cellStaged, cellDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
