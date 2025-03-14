// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/xlsx/go/models"
	"github.com/fullstack-lang/gong/lib/xlsx/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __XLRow__dummysDeclaration__ models.XLRow
var __XLRow_time__dummyDeclaration time.Duration

var mutexXLRow sync.Mutex

// An XLRowID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXLRow updateXLRow deleteXLRow
type XLRowID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// XLRowInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXLRow updateXLRow
type XLRowInput struct {
	// The XLRow to submit or modify
	// in: body
	XLRow *orm.XLRowAPI
}

// GetXLRows
//
// swagger:route GET /xlrows xlrows getXLRows
//
// # Get all xlrows
//
// Responses:
// default: genericError
//
//	200: xlrowDBResponse
func (controller *Controller) GetXLRows(c *gin.Context) {

	// source slice
	var xlrowDBs []orm.XLRowDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXLRows", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/xlsx/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoXLRow.GetDB()

	_, err := db.Find(&xlrowDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xlrowAPIs := make([]orm.XLRowAPI, 0)

	// for each xlrow, update fields from the database nullable fields
	for idx := range xlrowDBs {
		xlrowDB := &xlrowDBs[idx]
		_ = xlrowDB
		var xlrowAPI orm.XLRowAPI

		// insertion point for updating fields
		xlrowAPI.ID = xlrowDB.ID
		xlrowDB.CopyBasicFieldsToXLRow_WOP(&xlrowAPI.XLRow_WOP)
		xlrowAPI.XLRowPointersEncoding = xlrowDB.XLRowPointersEncoding
		xlrowAPIs = append(xlrowAPIs, xlrowAPI)
	}

	c.JSON(http.StatusOK, xlrowAPIs)
}

// PostXLRow
//
// swagger:route POST /xlrows xlrows postXLRow
//
// Creates a xlrow
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXLRow(c *gin.Context) {

	mutexXLRow.Lock()
	defer mutexXLRow.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXLRows", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/xlsx/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoXLRow.GetDB()

	// Validate input
	var input orm.XLRowAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xlrow
	xlrowDB := orm.XLRowDB{}
	xlrowDB.XLRowPointersEncoding = input.XLRowPointersEncoding
	xlrowDB.CopyBasicFieldsFromXLRow_WOP(&input.XLRow_WOP)

	_, err = db.Create(&xlrowDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXLRow.CheckoutPhaseOneInstance(&xlrowDB)
	xlrow := backRepo.BackRepoXLRow.Map_XLRowDBID_XLRowPtr[xlrowDB.ID]

	if xlrow != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xlrow)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xlrowDB)
}

// GetXLRow
//
// swagger:route GET /xlrows/{ID} xlrows getXLRow
//
// Gets the details for a xlrow.
//
// Responses:
// default: genericError
//
//	200: xlrowDBResponse
func (controller *Controller) GetXLRow(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXLRow", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/xlsx/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoXLRow.GetDB()

	// Get xlrowDB in DB
	var xlrowDB orm.XLRowDB
	if _, err := db.First(&xlrowDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xlrowAPI orm.XLRowAPI
	xlrowAPI.ID = xlrowDB.ID
	xlrowAPI.XLRowPointersEncoding = xlrowDB.XLRowPointersEncoding
	xlrowDB.CopyBasicFieldsToXLRow_WOP(&xlrowAPI.XLRow_WOP)

	c.JSON(http.StatusOK, xlrowAPI)
}

// UpdateXLRow
//
// swagger:route PATCH /xlrows/{ID} xlrows updateXLRow
//
// # Update a xlrow
//
// Responses:
// default: genericError
//
//	200: xlrowDBResponse
func (controller *Controller) UpdateXLRow(c *gin.Context) {

	mutexXLRow.Lock()
	defer mutexXLRow.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXLRow", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/xlsx/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoXLRow.GetDB()

	// Validate input
	var input orm.XLRowAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xlrowDB orm.XLRowDB

	// fetch the xlrow
	_, err := db.First(&xlrowDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xlrowDB.CopyBasicFieldsFromXLRow_WOP(&input.XLRow_WOP)
	xlrowDB.XLRowPointersEncoding = input.XLRowPointersEncoding

	db, _ = db.Model(&xlrowDB)
	_, err = db.Updates(&xlrowDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xlrowNew := new(models.XLRow)
	xlrowDB.CopyBasicFieldsToXLRow(xlrowNew)

	// redeem pointers
	xlrowDB.DecodePointers(backRepo, xlrowNew)

	// get stage instance from DB instance, and call callback function
	xlrowOld := backRepo.BackRepoXLRow.Map_XLRowDBID_XLRowPtr[xlrowDB.ID]
	if xlrowOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xlrowOld, xlrowNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xlrowDB
	c.JSON(http.StatusOK, xlrowDB)
}

// DeleteXLRow
//
// swagger:route DELETE /xlrows/{ID} xlrows deleteXLRow
//
// # Delete a xlrow
//
// default: genericError
//
//	200: xlrowDBResponse
func (controller *Controller) DeleteXLRow(c *gin.Context) {

	mutexXLRow.Lock()
	defer mutexXLRow.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXLRow", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/xlsx/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoXLRow.GetDB()

	// Get model if exist
	var xlrowDB orm.XLRowDB
	if _, err := db.First(&xlrowDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xlrowDB)

	// get an instance (not staged) from DB instance, and call callback function
	xlrowDeleted := new(models.XLRow)
	xlrowDB.CopyBasicFieldsToXLRow(xlrowDeleted)

	// get stage instance from DB instance, and call callback function
	xlrowStaged := backRepo.BackRepoXLRow.Map_XLRowDBID_XLRowPtr[xlrowDB.ID]
	if xlrowStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xlrowStaged, xlrowDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
