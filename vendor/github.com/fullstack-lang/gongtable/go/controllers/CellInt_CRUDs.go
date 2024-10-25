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
var __CellInt__dummysDeclaration__ models.CellInt
var __CellInt_time__dummyDeclaration time.Duration

var mutexCellInt sync.Mutex

// An CellIntID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCellInt updateCellInt deleteCellInt
type CellIntID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CellIntInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCellInt updateCellInt
type CellIntInput struct {
	// The CellInt to submit or modify
	// in: body
	CellInt *orm.CellIntAPI
}

// GetCellInts
//
// swagger:route GET /cellints cellints getCellInts
//
// # Get all cellints
//
// Responses:
// default: genericError
//
//	200: cellintDBResponse
func (controller *Controller) GetCellInts(c *gin.Context) {

	// source slice
	var cellintDBs []orm.CellIntDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCellInts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellInt.GetDB()

	_, err := db.Find(&cellintDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	cellintAPIs := make([]orm.CellIntAPI, 0)

	// for each cellint, update fields from the database nullable fields
	for idx := range cellintDBs {
		cellintDB := &cellintDBs[idx]
		_ = cellintDB
		var cellintAPI orm.CellIntAPI

		// insertion point for updating fields
		cellintAPI.ID = cellintDB.ID
		cellintDB.CopyBasicFieldsToCellInt_WOP(&cellintAPI.CellInt_WOP)
		cellintAPI.CellIntPointersEncoding = cellintDB.CellIntPointersEncoding
		cellintAPIs = append(cellintAPIs, cellintAPI)
	}

	c.JSON(http.StatusOK, cellintAPIs)
}

// PostCellInt
//
// swagger:route POST /cellints cellints postCellInt
//
// Creates a cellint
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCellInt(c *gin.Context) {

	mutexCellInt.Lock()
	defer mutexCellInt.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCellInts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellInt.GetDB()

	// Validate input
	var input orm.CellIntAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create cellint
	cellintDB := orm.CellIntDB{}
	cellintDB.CellIntPointersEncoding = input.CellIntPointersEncoding
	cellintDB.CopyBasicFieldsFromCellInt_WOP(&input.CellInt_WOP)

	_, err = db.Create(&cellintDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCellInt.CheckoutPhaseOneInstance(&cellintDB)
	cellint := backRepo.BackRepoCellInt.Map_CellIntDBID_CellIntPtr[cellintDB.ID]

	if cellint != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), cellint)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, cellintDB)
}

// GetCellInt
//
// swagger:route GET /cellints/{ID} cellints getCellInt
//
// Gets the details for a cellint.
//
// Responses:
// default: genericError
//
//	200: cellintDBResponse
func (controller *Controller) GetCellInt(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCellInt", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellInt.GetDB()

	// Get cellintDB in DB
	var cellintDB orm.CellIntDB
	if _, err := db.First(&cellintDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var cellintAPI orm.CellIntAPI
	cellintAPI.ID = cellintDB.ID
	cellintAPI.CellIntPointersEncoding = cellintDB.CellIntPointersEncoding
	cellintDB.CopyBasicFieldsToCellInt_WOP(&cellintAPI.CellInt_WOP)

	c.JSON(http.StatusOK, cellintAPI)
}

// UpdateCellInt
//
// swagger:route PATCH /cellints/{ID} cellints updateCellInt
//
// # Update a cellint
//
// Responses:
// default: genericError
//
//	200: cellintDBResponse
func (controller *Controller) UpdateCellInt(c *gin.Context) {

	mutexCellInt.Lock()
	defer mutexCellInt.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCellInt", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellInt.GetDB()

	// Validate input
	var input orm.CellIntAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var cellintDB orm.CellIntDB

	// fetch the cellint
	_, err := db.First(&cellintDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	cellintDB.CopyBasicFieldsFromCellInt_WOP(&input.CellInt_WOP)
	cellintDB.CellIntPointersEncoding = input.CellIntPointersEncoding

	db, _ = db.Model(&cellintDB)
	_, err = db.Updates(&cellintDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	cellintNew := new(models.CellInt)
	cellintDB.CopyBasicFieldsToCellInt(cellintNew)

	// redeem pointers
	cellintDB.DecodePointers(backRepo, cellintNew)

	// get stage instance from DB instance, and call callback function
	cellintOld := backRepo.BackRepoCellInt.Map_CellIntDBID_CellIntPtr[cellintDB.ID]
	if cellintOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), cellintOld, cellintNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the cellintDB
	c.JSON(http.StatusOK, cellintDB)
}

// DeleteCellInt
//
// swagger:route DELETE /cellints/{ID} cellints deleteCellInt
//
// # Delete a cellint
//
// default: genericError
//
//	200: cellintDBResponse
func (controller *Controller) DeleteCellInt(c *gin.Context) {

	mutexCellInt.Lock()
	defer mutexCellInt.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCellInt", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellInt.GetDB()

	// Get model if exist
	var cellintDB orm.CellIntDB
	if _, err := db.First(&cellintDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&cellintDB)

	// get an instance (not staged) from DB instance, and call callback function
	cellintDeleted := new(models.CellInt)
	cellintDB.CopyBasicFieldsToCellInt(cellintDeleted)

	// get stage instance from DB instance, and call callback function
	cellintStaged := backRepo.BackRepoCellInt.Map_CellIntDBID_CellIntPtr[cellintDB.ID]
	if cellintStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), cellintStaged, cellintDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
