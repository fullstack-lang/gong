// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/split/go/models"
	"github.com/fullstack-lang/gong/lib/split/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Cursor__dummysDeclaration__ models.Cursor
var __Cursor_time__dummyDeclaration time.Duration

var mutexCursor sync.Mutex

// An CursorID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCursor updateCursor deleteCursor
type CursorID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CursorInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCursor updateCursor
type CursorInput struct {
	// The Cursor to submit or modify
	// in: body
	Cursor *orm.CursorAPI
}

// GetCursors
//
// swagger:route GET /cursors cursors getCursors
//
// # Get all cursors
//
// Responses:
// default: genericError
//
//	200: cursorDBResponse
func (controller *Controller) GetCursors(c *gin.Context) {

	// source slice
	var cursorDBs []orm.CursorDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCursors", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoCursor.GetDB()

	_, err := db.Find(&cursorDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	cursorAPIs := make([]orm.CursorAPI, 0)

	// for each cursor, update fields from the database nullable fields
	for idx := range cursorDBs {
		cursorDB := &cursorDBs[idx]
		_ = cursorDB
		var cursorAPI orm.CursorAPI

		// insertion point for updating fields
		cursorAPI.ID = cursorDB.ID
		cursorDB.CopyBasicFieldsToCursor_WOP(&cursorAPI.Cursor_WOP)
		cursorAPI.CursorPointersEncoding = cursorDB.CursorPointersEncoding
		cursorAPIs = append(cursorAPIs, cursorAPI)
	}

	c.JSON(http.StatusOK, cursorAPIs)
}

// PostCursor
//
// swagger:route POST /cursors cursors postCursor
//
// Creates a cursor
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCursor(c *gin.Context) {

	mutexCursor.Lock()
	defer mutexCursor.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCursors", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoCursor.GetDB()

	// Validate input
	var input orm.CursorAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create cursor
	cursorDB := orm.CursorDB{}
	cursorDB.CursorPointersEncoding = input.CursorPointersEncoding
	cursorDB.CopyBasicFieldsFromCursor_WOP(&input.Cursor_WOP)

	_, err = db.Create(&cursorDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCursor.CheckoutPhaseOneInstance(&cursorDB)
	cursor := backRepo.BackRepoCursor.Map_CursorDBID_CursorPtr[cursorDB.ID]

	if cursor != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), cursor)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, cursorDB)
}

// GetCursor
//
// swagger:route GET /cursors/{ID} cursors getCursor
//
// Gets the details for a cursor.
//
// Responses:
// default: genericError
//
//	200: cursorDBResponse
func (controller *Controller) GetCursor(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCursor", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoCursor.GetDB()

	// Get cursorDB in DB
	var cursorDB orm.CursorDB
	if _, err := db.First(&cursorDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var cursorAPI orm.CursorAPI
	cursorAPI.ID = cursorDB.ID
	cursorAPI.CursorPointersEncoding = cursorDB.CursorPointersEncoding
	cursorDB.CopyBasicFieldsToCursor_WOP(&cursorAPI.Cursor_WOP)

	c.JSON(http.StatusOK, cursorAPI)
}

// UpdateCursor
//
// swagger:route PATCH /cursors/{ID} cursors updateCursor
//
// # Update a cursor
//
// Responses:
// default: genericError
//
//	200: cursorDBResponse
func (controller *Controller) UpdateCursor(c *gin.Context) {

	mutexCursor.Lock()
	defer mutexCursor.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCursor", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoCursor.GetDB()

	// Validate input
	var input orm.CursorAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var cursorDB orm.CursorDB

	// fetch the cursor
	_, err := db.First(&cursorDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	cursorDB.CopyBasicFieldsFromCursor_WOP(&input.Cursor_WOP)
	cursorDB.CursorPointersEncoding = input.CursorPointersEncoding

	db, _ = db.Model(&cursorDB)
	_, err = db.Updates(&cursorDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	cursorNew := new(models.Cursor)
	cursorDB.CopyBasicFieldsToCursor(cursorNew)

	// redeem pointers
	cursorDB.DecodePointers(backRepo, cursorNew)

	// get stage instance from DB instance, and call callback function
	cursorOld := backRepo.BackRepoCursor.Map_CursorDBID_CursorPtr[cursorDB.ID]
	if cursorOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), cursorOld, cursorNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the cursorDB
	c.JSON(http.StatusOK, cursorDB)
}

// DeleteCursor
//
// swagger:route DELETE /cursors/{ID} cursors deleteCursor
//
// # Delete a cursor
//
// default: genericError
//
//	200: cursorDBResponse
func (controller *Controller) DeleteCursor(c *gin.Context) {

	mutexCursor.Lock()
	defer mutexCursor.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCursor", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoCursor.GetDB()

	// Get model if exist
	var cursorDB orm.CursorDB
	if _, err := db.First(&cursorDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&cursorDB)

	// get an instance (not staged) from DB instance, and call callback function
	cursorDeleted := new(models.Cursor)
	cursorDB.CopyBasicFieldsToCursor(cursorDeleted)

	// get stage instance from DB instance, and call callback function
	cursorStaged := backRepo.BackRepoCursor.Map_CursorDBID_CursorPtr[cursorDB.ID]
	if cursorStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), cursorStaged, cursorDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
