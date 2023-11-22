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
var __CellString__dummysDeclaration__ models.CellString
var __CellString_time__dummyDeclaration time.Duration

var mutexCellString sync.Mutex

// An CellStringID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCellString updateCellString deleteCellString
type CellStringID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CellStringInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCellString updateCellString
type CellStringInput struct {
	// The CellString to submit or modify
	// in: body
	CellString *orm.CellStringAPI
}

// GetCellStrings
//
// swagger:route GET /cellstrings cellstrings getCellStrings
//
// # Get all cellstrings
//
// Responses:
// default: genericError
//
//	200: cellstringDBResponse
func (controller *Controller) GetCellStrings(c *gin.Context) {

	// source slice
	var cellstringDBs []orm.CellStringDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCellStrings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellString.GetDB()

	query := db.Find(&cellstringDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	cellstringAPIs := make([]orm.CellStringAPI, 0)

	// for each cellstring, update fields from the database nullable fields
	for idx := range cellstringDBs {
		cellstringDB := &cellstringDBs[idx]
		_ = cellstringDB
		var cellstringAPI orm.CellStringAPI

		// insertion point for updating fields
		cellstringAPI.ID = cellstringDB.ID
		cellstringDB.CopyBasicFieldsToCellString_WOP(&cellstringAPI.CellString_WOP)
		cellstringAPI.CellStringPointersEncoding = cellstringDB.CellStringPointersEncoding
		cellstringAPIs = append(cellstringAPIs, cellstringAPI)
	}

	c.JSON(http.StatusOK, cellstringAPIs)
}

// PostCellString
//
// swagger:route POST /cellstrings cellstrings postCellString
//
// Creates a cellstring
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCellString(c *gin.Context) {

	mutexCellString.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCellStrings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellString.GetDB()

	// Validate input
	var input orm.CellStringAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create cellstring
	cellstringDB := orm.CellStringDB{}
	cellstringDB.CellStringPointersEncoding = input.CellStringPointersEncoding
	cellstringDB.CopyBasicFieldsFromCellString_WOP(&input.CellString_WOP)

	query := db.Create(&cellstringDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCellString.CheckoutPhaseOneInstance(&cellstringDB)
	cellstring := backRepo.BackRepoCellString.Map_CellStringDBID_CellStringPtr[cellstringDB.ID]

	if cellstring != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), cellstring)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, cellstringDB)

	mutexCellString.Unlock()
}

// GetCellString
//
// swagger:route GET /cellstrings/{ID} cellstrings getCellString
//
// Gets the details for a cellstring.
//
// Responses:
// default: genericError
//
//	200: cellstringDBResponse
func (controller *Controller) GetCellString(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCellString", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellString.GetDB()

	// Get cellstringDB in DB
	var cellstringDB orm.CellStringDB
	if err := db.First(&cellstringDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var cellstringAPI orm.CellStringAPI
	cellstringAPI.ID = cellstringDB.ID
	cellstringAPI.CellStringPointersEncoding = cellstringDB.CellStringPointersEncoding
	cellstringDB.CopyBasicFieldsToCellString_WOP(&cellstringAPI.CellString_WOP)

	c.JSON(http.StatusOK, cellstringAPI)
}

// UpdateCellString
//
// swagger:route PATCH /cellstrings/{ID} cellstrings updateCellString
//
// # Update a cellstring
//
// Responses:
// default: genericError
//
//	200: cellstringDBResponse
func (controller *Controller) UpdateCellString(c *gin.Context) {

	mutexCellString.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCellString", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellString.GetDB()

	// Validate input
	var input orm.CellStringAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var cellstringDB orm.CellStringDB

	// fetch the cellstring
	query := db.First(&cellstringDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	cellstringDB.CopyBasicFieldsFromCellString_WOP(&input.CellString_WOP)
	cellstringDB.CellStringPointersEncoding = input.CellStringPointersEncoding

	query = db.Model(&cellstringDB).Updates(cellstringDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	cellstringNew := new(models.CellString)
	cellstringDB.CopyBasicFieldsToCellString(cellstringNew)

	// redeem pointers
	cellstringDB.DecodePointers(backRepo, cellstringNew)

	// get stage instance from DB instance, and call callback function
	cellstringOld := backRepo.BackRepoCellString.Map_CellStringDBID_CellStringPtr[cellstringDB.ID]
	if cellstringOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), cellstringOld, cellstringNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the cellstringDB
	c.JSON(http.StatusOK, cellstringDB)

	mutexCellString.Unlock()
}

// DeleteCellString
//
// swagger:route DELETE /cellstrings/{ID} cellstrings deleteCellString
//
// # Delete a cellstring
//
// default: genericError
//
//	200: cellstringDBResponse
func (controller *Controller) DeleteCellString(c *gin.Context) {

	mutexCellString.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCellString", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCellString.GetDB()

	// Get model if exist
	var cellstringDB orm.CellStringDB
	if err := db.First(&cellstringDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&cellstringDB)

	// get an instance (not staged) from DB instance, and call callback function
	cellstringDeleted := new(models.CellString)
	cellstringDB.CopyBasicFieldsToCellString(cellstringDeleted)

	// get stage instance from DB instance, and call callback function
	cellstringStaged := backRepo.BackRepoCellString.Map_CellStringDBID_CellStringPtr[cellstringDB.ID]
	if cellstringStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), cellstringStaged, cellstringDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexCellString.Unlock()
}
