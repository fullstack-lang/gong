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
var __CellIcon__dummysDeclaration__ models.CellIcon
var __CellIcon_time__dummyDeclaration time.Duration

var mutexCellIcon sync.Mutex

// An CellIconID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCellIcon updateCellIcon deleteCellIcon
type CellIconID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CellIconInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCellIcon updateCellIcon
type CellIconInput struct {
	// The CellIcon to submit or modify
	// in: body
	CellIcon *orm.CellIconAPI
}

// GetCellIcons
//
// swagger:route GET /cellicons cellicons getCellIcons
//
// # Get all cellicons
//
// Responses:
// default: genericError
//
//	200: celliconDBResponse
func (controller *Controller) GetCellIcons(c *gin.Context) {

	// source slice
	var celliconDBs []orm.CellIconDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCellIcons", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellIcon.GetDB()

	_, err := db.Find(&celliconDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	celliconAPIs := make([]orm.CellIconAPI, 0)

	// for each cellicon, update fields from the database nullable fields
	for idx := range celliconDBs {
		celliconDB := &celliconDBs[idx]
		_ = celliconDB
		var celliconAPI orm.CellIconAPI

		// insertion point for updating fields
		celliconAPI.ID = celliconDB.ID
		celliconDB.CopyBasicFieldsToCellIcon_WOP(&celliconAPI.CellIcon_WOP)
		celliconAPI.CellIconPointersEncoding = celliconDB.CellIconPointersEncoding
		celliconAPIs = append(celliconAPIs, celliconAPI)
	}

	c.JSON(http.StatusOK, celliconAPIs)
}

// PostCellIcon
//
// swagger:route POST /cellicons cellicons postCellIcon
//
// Creates a cellicon
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCellIcon(c *gin.Context) {

	mutexCellIcon.Lock()
	defer mutexCellIcon.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCellIcons", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellIcon.GetDB()

	// Validate input
	var input orm.CellIconAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create cellicon
	celliconDB := orm.CellIconDB{}
	celliconDB.CellIconPointersEncoding = input.CellIconPointersEncoding
	celliconDB.CopyBasicFieldsFromCellIcon_WOP(&input.CellIcon_WOP)

	_, err = db.Create(&celliconDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCellIcon.CheckoutPhaseOneInstance(&celliconDB)
	cellicon := backRepo.BackRepoCellIcon.Map_CellIconDBID_CellIconPtr[celliconDB.ID]

	if cellicon != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), cellicon)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, celliconDB)
}

// GetCellIcon
//
// swagger:route GET /cellicons/{ID} cellicons getCellIcon
//
// Gets the details for a cellicon.
//
// Responses:
// default: genericError
//
//	200: celliconDBResponse
func (controller *Controller) GetCellIcon(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCellIcon", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellIcon.GetDB()

	// Get celliconDB in DB
	var celliconDB orm.CellIconDB
	if _, err := db.First(&celliconDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var celliconAPI orm.CellIconAPI
	celliconAPI.ID = celliconDB.ID
	celliconAPI.CellIconPointersEncoding = celliconDB.CellIconPointersEncoding
	celliconDB.CopyBasicFieldsToCellIcon_WOP(&celliconAPI.CellIcon_WOP)

	c.JSON(http.StatusOK, celliconAPI)
}

// UpdateCellIcon
//
// swagger:route PATCH /cellicons/{ID} cellicons updateCellIcon
//
// # Update a cellicon
//
// Responses:
// default: genericError
//
//	200: celliconDBResponse
func (controller *Controller) UpdateCellIcon(c *gin.Context) {

	mutexCellIcon.Lock()
	defer mutexCellIcon.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCellIcon", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellIcon.GetDB()

	// Validate input
	var input orm.CellIconAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var celliconDB orm.CellIconDB

	// fetch the cellicon
	_, err := db.First(&celliconDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	celliconDB.CopyBasicFieldsFromCellIcon_WOP(&input.CellIcon_WOP)
	celliconDB.CellIconPointersEncoding = input.CellIconPointersEncoding

	db, _ = db.Model(&celliconDB)
	_, err = db.Updates(&celliconDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	celliconNew := new(models.CellIcon)
	celliconDB.CopyBasicFieldsToCellIcon(celliconNew)

	// redeem pointers
	celliconDB.DecodePointers(backRepo, celliconNew)

	// get stage instance from DB instance, and call callback function
	celliconOld := backRepo.BackRepoCellIcon.Map_CellIconDBID_CellIconPtr[celliconDB.ID]
	if celliconOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), celliconOld, celliconNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the celliconDB
	c.JSON(http.StatusOK, celliconDB)
}

// DeleteCellIcon
//
// swagger:route DELETE /cellicons/{ID} cellicons deleteCellIcon
//
// # Delete a cellicon
//
// default: genericError
//
//	200: celliconDBResponse
func (controller *Controller) DeleteCellIcon(c *gin.Context) {

	mutexCellIcon.Lock()
	defer mutexCellIcon.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCellIcon", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellIcon.GetDB()

	// Get model if exist
	var celliconDB orm.CellIconDB
	if _, err := db.First(&celliconDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&celliconDB)

	// get an instance (not staged) from DB instance, and call callback function
	celliconDeleted := new(models.CellIcon)
	celliconDB.CopyBasicFieldsToCellIcon(celliconDeleted)

	// get stage instance from DB instance, and call callback function
	celliconStaged := backRepo.BackRepoCellIcon.Map_CellIconDBID_CellIconPtr[celliconDB.ID]
	if celliconStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), celliconStaged, celliconDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
