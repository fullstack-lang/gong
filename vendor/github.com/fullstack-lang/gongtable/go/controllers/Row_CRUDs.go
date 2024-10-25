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
var __Row__dummysDeclaration__ models.Row
var __Row_time__dummyDeclaration time.Duration

var mutexRow sync.Mutex

// An RowID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRow updateRow deleteRow
type RowID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RowInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRow updateRow
type RowInput struct {
	// The Row to submit or modify
	// in: body
	Row *orm.RowAPI
}

// GetRows
//
// swagger:route GET /rows rows getRows
//
// # Get all rows
//
// Responses:
// default: genericError
//
//	200: rowDBResponse
func (controller *Controller) GetRows(c *gin.Context) {

	// source slice
	var rowDBs []orm.RowDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRows", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRow.GetDB()

	_, err := db.Find(&rowDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	rowAPIs := make([]orm.RowAPI, 0)

	// for each row, update fields from the database nullable fields
	for idx := range rowDBs {
		rowDB := &rowDBs[idx]
		_ = rowDB
		var rowAPI orm.RowAPI

		// insertion point for updating fields
		rowAPI.ID = rowDB.ID
		rowDB.CopyBasicFieldsToRow_WOP(&rowAPI.Row_WOP)
		rowAPI.RowPointersEncoding = rowDB.RowPointersEncoding
		rowAPIs = append(rowAPIs, rowAPI)
	}

	c.JSON(http.StatusOK, rowAPIs)
}

// PostRow
//
// swagger:route POST /rows rows postRow
//
// Creates a row
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRow(c *gin.Context) {

	mutexRow.Lock()
	defer mutexRow.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRows", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRow.GetDB()

	// Validate input
	var input orm.RowAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create row
	rowDB := orm.RowDB{}
	rowDB.RowPointersEncoding = input.RowPointersEncoding
	rowDB.CopyBasicFieldsFromRow_WOP(&input.Row_WOP)

	_, err = db.Create(&rowDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRow.CheckoutPhaseOneInstance(&rowDB)
	row := backRepo.BackRepoRow.Map_RowDBID_RowPtr[rowDB.ID]

	if row != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), row)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, rowDB)
}

// GetRow
//
// swagger:route GET /rows/{ID} rows getRow
//
// Gets the details for a row.
//
// Responses:
// default: genericError
//
//	200: rowDBResponse
func (controller *Controller) GetRow(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRow", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRow.GetDB()

	// Get rowDB in DB
	var rowDB orm.RowDB
	if _, err := db.First(&rowDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var rowAPI orm.RowAPI
	rowAPI.ID = rowDB.ID
	rowAPI.RowPointersEncoding = rowDB.RowPointersEncoding
	rowDB.CopyBasicFieldsToRow_WOP(&rowAPI.Row_WOP)

	c.JSON(http.StatusOK, rowAPI)
}

// UpdateRow
//
// swagger:route PATCH /rows/{ID} rows updateRow
//
// # Update a row
//
// Responses:
// default: genericError
//
//	200: rowDBResponse
func (controller *Controller) UpdateRow(c *gin.Context) {

	mutexRow.Lock()
	defer mutexRow.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRow", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRow.GetDB()

	// Validate input
	var input orm.RowAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var rowDB orm.RowDB

	// fetch the row
	_, err := db.First(&rowDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	rowDB.CopyBasicFieldsFromRow_WOP(&input.Row_WOP)
	rowDB.RowPointersEncoding = input.RowPointersEncoding

	db, _ = db.Model(&rowDB)
	_, err = db.Updates(&rowDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	rowNew := new(models.Row)
	rowDB.CopyBasicFieldsToRow(rowNew)

	// redeem pointers
	rowDB.DecodePointers(backRepo, rowNew)

	// get stage instance from DB instance, and call callback function
	rowOld := backRepo.BackRepoRow.Map_RowDBID_RowPtr[rowDB.ID]
	if rowOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), rowOld, rowNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the rowDB
	c.JSON(http.StatusOK, rowDB)
}

// DeleteRow
//
// swagger:route DELETE /rows/{ID} rows deleteRow
//
// # Delete a row
//
// default: genericError
//
//	200: rowDBResponse
func (controller *Controller) DeleteRow(c *gin.Context) {

	mutexRow.Lock()
	defer mutexRow.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRow", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRow.GetDB()

	// Get model if exist
	var rowDB orm.RowDB
	if _, err := db.First(&rowDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&rowDB)

	// get an instance (not staged) from DB instance, and call callback function
	rowDeleted := new(models.Row)
	rowDB.CopyBasicFieldsToRow(rowDeleted)

	// get stage instance from DB instance, and call callback function
	rowStaged := backRepo.BackRepoRow.Map_RowDBID_RowPtr[rowDB.ID]
	if rowStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), rowStaged, rowDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
