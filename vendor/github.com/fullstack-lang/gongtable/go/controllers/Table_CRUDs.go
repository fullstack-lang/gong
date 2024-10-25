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
var __Table__dummysDeclaration__ models.Table
var __Table_time__dummyDeclaration time.Duration

var mutexTable sync.Mutex

// An TableID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTable updateTable deleteTable
type TableID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TableInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTable updateTable
type TableInput struct {
	// The Table to submit or modify
	// in: body
	Table *orm.TableAPI
}

// GetTables
//
// swagger:route GET /tables tables getTables
//
// # Get all tables
//
// Responses:
// default: genericError
//
//	200: tableDBResponse
func (controller *Controller) GetTables(c *gin.Context) {

	// source slice
	var tableDBs []orm.TableDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTables", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTable.GetDB()

	_, err := db.Find(&tableDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	tableAPIs := make([]orm.TableAPI, 0)

	// for each table, update fields from the database nullable fields
	for idx := range tableDBs {
		tableDB := &tableDBs[idx]
		_ = tableDB
		var tableAPI orm.TableAPI

		// insertion point for updating fields
		tableAPI.ID = tableDB.ID
		tableDB.CopyBasicFieldsToTable_WOP(&tableAPI.Table_WOP)
		tableAPI.TablePointersEncoding = tableDB.TablePointersEncoding
		tableAPIs = append(tableAPIs, tableAPI)
	}

	c.JSON(http.StatusOK, tableAPIs)
}

// PostTable
//
// swagger:route POST /tables tables postTable
//
// Creates a table
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTable(c *gin.Context) {

	mutexTable.Lock()
	defer mutexTable.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTables", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTable.GetDB()

	// Validate input
	var input orm.TableAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create table
	tableDB := orm.TableDB{}
	tableDB.TablePointersEncoding = input.TablePointersEncoding
	tableDB.CopyBasicFieldsFromTable_WOP(&input.Table_WOP)

	_, err = db.Create(&tableDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTable.CheckoutPhaseOneInstance(&tableDB)
	table := backRepo.BackRepoTable.Map_TableDBID_TablePtr[tableDB.ID]

	if table != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), table)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, tableDB)
}

// GetTable
//
// swagger:route GET /tables/{ID} tables getTable
//
// Gets the details for a table.
//
// Responses:
// default: genericError
//
//	200: tableDBResponse
func (controller *Controller) GetTable(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTable", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTable.GetDB()

	// Get tableDB in DB
	var tableDB orm.TableDB
	if _, err := db.First(&tableDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var tableAPI orm.TableAPI
	tableAPI.ID = tableDB.ID
	tableAPI.TablePointersEncoding = tableDB.TablePointersEncoding
	tableDB.CopyBasicFieldsToTable_WOP(&tableAPI.Table_WOP)

	c.JSON(http.StatusOK, tableAPI)
}

// UpdateTable
//
// swagger:route PATCH /tables/{ID} tables updateTable
//
// # Update a table
//
// Responses:
// default: genericError
//
//	200: tableDBResponse
func (controller *Controller) UpdateTable(c *gin.Context) {

	mutexTable.Lock()
	defer mutexTable.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTable", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTable.GetDB()

	// Validate input
	var input orm.TableAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var tableDB orm.TableDB

	// fetch the table
	_, err := db.First(&tableDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	tableDB.CopyBasicFieldsFromTable_WOP(&input.Table_WOP)
	tableDB.TablePointersEncoding = input.TablePointersEncoding

	db, _ = db.Model(&tableDB)
	_, err = db.Updates(&tableDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	tableNew := new(models.Table)
	tableDB.CopyBasicFieldsToTable(tableNew)

	// redeem pointers
	tableDB.DecodePointers(backRepo, tableNew)

	// get stage instance from DB instance, and call callback function
	tableOld := backRepo.BackRepoTable.Map_TableDBID_TablePtr[tableDB.ID]
	if tableOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), tableOld, tableNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the tableDB
	c.JSON(http.StatusOK, tableDB)
}

// DeleteTable
//
// swagger:route DELETE /tables/{ID} tables deleteTable
//
// # Delete a table
//
// default: genericError
//
//	200: tableDBResponse
func (controller *Controller) DeleteTable(c *gin.Context) {

	mutexTable.Lock()
	defer mutexTable.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTable", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTable.GetDB()

	// Get model if exist
	var tableDB orm.TableDB
	if _, err := db.First(&tableDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&tableDB)

	// get an instance (not staged) from DB instance, and call callback function
	tableDeleted := new(models.Table)
	tableDB.CopyBasicFieldsToTable(tableDeleted)

	// get stage instance from DB instance, and call callback function
	tableStaged := backRepo.BackRepoTable.Map_TableDBID_TablePtr[tableDB.ID]
	if tableStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), tableStaged, tableDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
