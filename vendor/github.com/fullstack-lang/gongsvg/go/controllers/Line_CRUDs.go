// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongsvg/go/models"
	"github.com/fullstack-lang/gongsvg/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Line__dummysDeclaration__ models.Line
var __Line_time__dummyDeclaration time.Duration

var mutexLine sync.Mutex

// An LineID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLine updateLine deleteLine
type LineID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LineInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLine updateLine
type LineInput struct {
	// The Line to submit or modify
	// in: body
	Line *orm.LineAPI
}

// GetLines
//
// swagger:route GET /lines lines getLines
//
// # Get all lines
//
// Responses:
// default: genericError
//
//	200: lineDBResponse
func (controller *Controller) GetLines(c *gin.Context) {

	// source slice
	var lineDBs []orm.LineDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLines", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLine.GetDB()

	_, err := db.Find(&lineDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	lineAPIs := make([]orm.LineAPI, 0)

	// for each line, update fields from the database nullable fields
	for idx := range lineDBs {
		lineDB := &lineDBs[idx]
		_ = lineDB
		var lineAPI orm.LineAPI

		// insertion point for updating fields
		lineAPI.ID = lineDB.ID
		lineDB.CopyBasicFieldsToLine_WOP(&lineAPI.Line_WOP)
		lineAPI.LinePointersEncoding = lineDB.LinePointersEncoding
		lineAPIs = append(lineAPIs, lineAPI)
	}

	c.JSON(http.StatusOK, lineAPIs)
}

// PostLine
//
// swagger:route POST /lines lines postLine
//
// Creates a line
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLine(c *gin.Context) {

	mutexLine.Lock()
	defer mutexLine.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLines", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLine.GetDB()

	// Validate input
	var input orm.LineAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create line
	lineDB := orm.LineDB{}
	lineDB.LinePointersEncoding = input.LinePointersEncoding
	lineDB.CopyBasicFieldsFromLine_WOP(&input.Line_WOP)

	_, err = db.Create(&lineDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLine.CheckoutPhaseOneInstance(&lineDB)
	line := backRepo.BackRepoLine.Map_LineDBID_LinePtr[lineDB.ID]

	if line != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), line)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, lineDB)
}

// GetLine
//
// swagger:route GET /lines/{ID} lines getLine
//
// Gets the details for a line.
//
// Responses:
// default: genericError
//
//	200: lineDBResponse
func (controller *Controller) GetLine(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLine", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLine.GetDB()

	// Get lineDB in DB
	var lineDB orm.LineDB
	if _, err := db.First(&lineDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var lineAPI orm.LineAPI
	lineAPI.ID = lineDB.ID
	lineAPI.LinePointersEncoding = lineDB.LinePointersEncoding
	lineDB.CopyBasicFieldsToLine_WOP(&lineAPI.Line_WOP)

	c.JSON(http.StatusOK, lineAPI)
}

// UpdateLine
//
// swagger:route PATCH /lines/{ID} lines updateLine
//
// # Update a line
//
// Responses:
// default: genericError
//
//	200: lineDBResponse
func (controller *Controller) UpdateLine(c *gin.Context) {

	mutexLine.Lock()
	defer mutexLine.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLine", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLine.GetDB()

	// Validate input
	var input orm.LineAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var lineDB orm.LineDB

	// fetch the line
	_, err := db.First(&lineDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	lineDB.CopyBasicFieldsFromLine_WOP(&input.Line_WOP)
	lineDB.LinePointersEncoding = input.LinePointersEncoding

	db, _ = db.Model(&lineDB)
	_, err = db.Updates(&lineDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	lineNew := new(models.Line)
	lineDB.CopyBasicFieldsToLine(lineNew)

	// redeem pointers
	lineDB.DecodePointers(backRepo, lineNew)

	// get stage instance from DB instance, and call callback function
	lineOld := backRepo.BackRepoLine.Map_LineDBID_LinePtr[lineDB.ID]
	if lineOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), lineOld, lineNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the lineDB
	c.JSON(http.StatusOK, lineDB)
}

// DeleteLine
//
// swagger:route DELETE /lines/{ID} lines deleteLine
//
// # Delete a line
//
// default: genericError
//
//	200: lineDBResponse
func (controller *Controller) DeleteLine(c *gin.Context) {

	mutexLine.Lock()
	defer mutexLine.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLine", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLine.GetDB()

	// Get model if exist
	var lineDB orm.LineDB
	if _, err := db.First(&lineDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&lineDB)

	// get an instance (not staged) from DB instance, and call callback function
	lineDeleted := new(models.Line)
	lineDB.CopyBasicFieldsToLine(lineDeleted)

	// get stage instance from DB instance, and call callback function
	lineStaged := backRepo.BackRepoLine.Map_LineDBID_LinePtr[lineDB.ID]
	if lineStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), lineStaged, lineDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
