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
var __DisplayedColumn__dummysDeclaration__ models.DisplayedColumn
var __DisplayedColumn_time__dummyDeclaration time.Duration

var mutexDisplayedColumn sync.Mutex

// An DisplayedColumnID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDisplayedColumn updateDisplayedColumn deleteDisplayedColumn
type DisplayedColumnID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DisplayedColumnInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDisplayedColumn updateDisplayedColumn
type DisplayedColumnInput struct {
	// The DisplayedColumn to submit or modify
	// in: body
	DisplayedColumn *orm.DisplayedColumnAPI
}

// GetDisplayedColumns
//
// swagger:route GET /displayedcolumns displayedcolumns getDisplayedColumns
//
// # Get all displayedcolumns
//
// Responses:
// default: genericError
//
//	200: displayedcolumnDBResponse
func (controller *Controller) GetDisplayedColumns(c *gin.Context) {

	// source slice
	var displayedcolumnDBs []orm.DisplayedColumnDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDisplayedColumns", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDisplayedColumn.GetDB()

	_, err := db.Find(&displayedcolumnDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	displayedcolumnAPIs := make([]orm.DisplayedColumnAPI, 0)

	// for each displayedcolumn, update fields from the database nullable fields
	for idx := range displayedcolumnDBs {
		displayedcolumnDB := &displayedcolumnDBs[idx]
		_ = displayedcolumnDB
		var displayedcolumnAPI orm.DisplayedColumnAPI

		// insertion point for updating fields
		displayedcolumnAPI.ID = displayedcolumnDB.ID
		displayedcolumnDB.CopyBasicFieldsToDisplayedColumn_WOP(&displayedcolumnAPI.DisplayedColumn_WOP)
		displayedcolumnAPI.DisplayedColumnPointersEncoding = displayedcolumnDB.DisplayedColumnPointersEncoding
		displayedcolumnAPIs = append(displayedcolumnAPIs, displayedcolumnAPI)
	}

	c.JSON(http.StatusOK, displayedcolumnAPIs)
}

// PostDisplayedColumn
//
// swagger:route POST /displayedcolumns displayedcolumns postDisplayedColumn
//
// Creates a displayedcolumn
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDisplayedColumn(c *gin.Context) {

	mutexDisplayedColumn.Lock()
	defer mutexDisplayedColumn.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDisplayedColumns", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDisplayedColumn.GetDB()

	// Validate input
	var input orm.DisplayedColumnAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create displayedcolumn
	displayedcolumnDB := orm.DisplayedColumnDB{}
	displayedcolumnDB.DisplayedColumnPointersEncoding = input.DisplayedColumnPointersEncoding
	displayedcolumnDB.CopyBasicFieldsFromDisplayedColumn_WOP(&input.DisplayedColumn_WOP)

	_, err = db.Create(&displayedcolumnDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDisplayedColumn.CheckoutPhaseOneInstance(&displayedcolumnDB)
	displayedcolumn := backRepo.BackRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnPtr[displayedcolumnDB.ID]

	if displayedcolumn != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), displayedcolumn)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, displayedcolumnDB)
}

// GetDisplayedColumn
//
// swagger:route GET /displayedcolumns/{ID} displayedcolumns getDisplayedColumn
//
// Gets the details for a displayedcolumn.
//
// Responses:
// default: genericError
//
//	200: displayedcolumnDBResponse
func (controller *Controller) GetDisplayedColumn(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDisplayedColumn", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDisplayedColumn.GetDB()

	// Get displayedcolumnDB in DB
	var displayedcolumnDB orm.DisplayedColumnDB
	if _, err := db.First(&displayedcolumnDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var displayedcolumnAPI orm.DisplayedColumnAPI
	displayedcolumnAPI.ID = displayedcolumnDB.ID
	displayedcolumnAPI.DisplayedColumnPointersEncoding = displayedcolumnDB.DisplayedColumnPointersEncoding
	displayedcolumnDB.CopyBasicFieldsToDisplayedColumn_WOP(&displayedcolumnAPI.DisplayedColumn_WOP)

	c.JSON(http.StatusOK, displayedcolumnAPI)
}

// UpdateDisplayedColumn
//
// swagger:route PATCH /displayedcolumns/{ID} displayedcolumns updateDisplayedColumn
//
// # Update a displayedcolumn
//
// Responses:
// default: genericError
//
//	200: displayedcolumnDBResponse
func (controller *Controller) UpdateDisplayedColumn(c *gin.Context) {

	mutexDisplayedColumn.Lock()
	defer mutexDisplayedColumn.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDisplayedColumn", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDisplayedColumn.GetDB()

	// Validate input
	var input orm.DisplayedColumnAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var displayedcolumnDB orm.DisplayedColumnDB

	// fetch the displayedcolumn
	_, err := db.First(&displayedcolumnDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	displayedcolumnDB.CopyBasicFieldsFromDisplayedColumn_WOP(&input.DisplayedColumn_WOP)
	displayedcolumnDB.DisplayedColumnPointersEncoding = input.DisplayedColumnPointersEncoding

	db, _ = db.Model(&displayedcolumnDB)
	_, err = db.Updates(&displayedcolumnDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	displayedcolumnNew := new(models.DisplayedColumn)
	displayedcolumnDB.CopyBasicFieldsToDisplayedColumn(displayedcolumnNew)

	// redeem pointers
	displayedcolumnDB.DecodePointers(backRepo, displayedcolumnNew)

	// get stage instance from DB instance, and call callback function
	displayedcolumnOld := backRepo.BackRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnPtr[displayedcolumnDB.ID]
	if displayedcolumnOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), displayedcolumnOld, displayedcolumnNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the displayedcolumnDB
	c.JSON(http.StatusOK, displayedcolumnDB)
}

// DeleteDisplayedColumn
//
// swagger:route DELETE /displayedcolumns/{ID} displayedcolumns deleteDisplayedColumn
//
// # Delete a displayedcolumn
//
// default: genericError
//
//	200: displayedcolumnDBResponse
func (controller *Controller) DeleteDisplayedColumn(c *gin.Context) {

	mutexDisplayedColumn.Lock()
	defer mutexDisplayedColumn.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDisplayedColumn", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDisplayedColumn.GetDB()

	// Get model if exist
	var displayedcolumnDB orm.DisplayedColumnDB
	if _, err := db.First(&displayedcolumnDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&displayedcolumnDB)

	// get an instance (not staged) from DB instance, and call callback function
	displayedcolumnDeleted := new(models.DisplayedColumn)
	displayedcolumnDB.CopyBasicFieldsToDisplayedColumn(displayedcolumnDeleted)

	// get stage instance from DB instance, and call callback function
	displayedcolumnStaged := backRepo.BackRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnPtr[displayedcolumnDB.ID]
	if displayedcolumnStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), displayedcolumnStaged, displayedcolumnDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
