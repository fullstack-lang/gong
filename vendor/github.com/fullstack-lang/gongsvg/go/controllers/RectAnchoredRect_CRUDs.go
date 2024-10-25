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
var __RectAnchoredRect__dummysDeclaration__ models.RectAnchoredRect
var __RectAnchoredRect_time__dummyDeclaration time.Duration

var mutexRectAnchoredRect sync.Mutex

// An RectAnchoredRectID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRectAnchoredRect updateRectAnchoredRect deleteRectAnchoredRect
type RectAnchoredRectID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RectAnchoredRectInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRectAnchoredRect updateRectAnchoredRect
type RectAnchoredRectInput struct {
	// The RectAnchoredRect to submit or modify
	// in: body
	RectAnchoredRect *orm.RectAnchoredRectAPI
}

// GetRectAnchoredRects
//
// swagger:route GET /rectanchoredrects rectanchoredrects getRectAnchoredRects
//
// # Get all rectanchoredrects
//
// Responses:
// default: genericError
//
//	200: rectanchoredrectDBResponse
func (controller *Controller) GetRectAnchoredRects(c *gin.Context) {

	// source slice
	var rectanchoredrectDBs []orm.RectAnchoredRectDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRectAnchoredRects", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRectAnchoredRect.GetDB()

	_, err := db.Find(&rectanchoredrectDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	rectanchoredrectAPIs := make([]orm.RectAnchoredRectAPI, 0)

	// for each rectanchoredrect, update fields from the database nullable fields
	for idx := range rectanchoredrectDBs {
		rectanchoredrectDB := &rectanchoredrectDBs[idx]
		_ = rectanchoredrectDB
		var rectanchoredrectAPI orm.RectAnchoredRectAPI

		// insertion point for updating fields
		rectanchoredrectAPI.ID = rectanchoredrectDB.ID
		rectanchoredrectDB.CopyBasicFieldsToRectAnchoredRect_WOP(&rectanchoredrectAPI.RectAnchoredRect_WOP)
		rectanchoredrectAPI.RectAnchoredRectPointersEncoding = rectanchoredrectDB.RectAnchoredRectPointersEncoding
		rectanchoredrectAPIs = append(rectanchoredrectAPIs, rectanchoredrectAPI)
	}

	c.JSON(http.StatusOK, rectanchoredrectAPIs)
}

// PostRectAnchoredRect
//
// swagger:route POST /rectanchoredrects rectanchoredrects postRectAnchoredRect
//
// Creates a rectanchoredrect
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRectAnchoredRect(c *gin.Context) {

	mutexRectAnchoredRect.Lock()
	defer mutexRectAnchoredRect.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRectAnchoredRects", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRectAnchoredRect.GetDB()

	// Validate input
	var input orm.RectAnchoredRectAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create rectanchoredrect
	rectanchoredrectDB := orm.RectAnchoredRectDB{}
	rectanchoredrectDB.RectAnchoredRectPointersEncoding = input.RectAnchoredRectPointersEncoding
	rectanchoredrectDB.CopyBasicFieldsFromRectAnchoredRect_WOP(&input.RectAnchoredRect_WOP)

	_, err = db.Create(&rectanchoredrectDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRectAnchoredRect.CheckoutPhaseOneInstance(&rectanchoredrectDB)
	rectanchoredrect := backRepo.BackRepoRectAnchoredRect.Map_RectAnchoredRectDBID_RectAnchoredRectPtr[rectanchoredrectDB.ID]

	if rectanchoredrect != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), rectanchoredrect)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, rectanchoredrectDB)
}

// GetRectAnchoredRect
//
// swagger:route GET /rectanchoredrects/{ID} rectanchoredrects getRectAnchoredRect
//
// Gets the details for a rectanchoredrect.
//
// Responses:
// default: genericError
//
//	200: rectanchoredrectDBResponse
func (controller *Controller) GetRectAnchoredRect(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRectAnchoredRect", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRectAnchoredRect.GetDB()

	// Get rectanchoredrectDB in DB
	var rectanchoredrectDB orm.RectAnchoredRectDB
	if _, err := db.First(&rectanchoredrectDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var rectanchoredrectAPI orm.RectAnchoredRectAPI
	rectanchoredrectAPI.ID = rectanchoredrectDB.ID
	rectanchoredrectAPI.RectAnchoredRectPointersEncoding = rectanchoredrectDB.RectAnchoredRectPointersEncoding
	rectanchoredrectDB.CopyBasicFieldsToRectAnchoredRect_WOP(&rectanchoredrectAPI.RectAnchoredRect_WOP)

	c.JSON(http.StatusOK, rectanchoredrectAPI)
}

// UpdateRectAnchoredRect
//
// swagger:route PATCH /rectanchoredrects/{ID} rectanchoredrects updateRectAnchoredRect
//
// # Update a rectanchoredrect
//
// Responses:
// default: genericError
//
//	200: rectanchoredrectDBResponse
func (controller *Controller) UpdateRectAnchoredRect(c *gin.Context) {

	mutexRectAnchoredRect.Lock()
	defer mutexRectAnchoredRect.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRectAnchoredRect", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRectAnchoredRect.GetDB()

	// Validate input
	var input orm.RectAnchoredRectAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var rectanchoredrectDB orm.RectAnchoredRectDB

	// fetch the rectanchoredrect
	_, err := db.First(&rectanchoredrectDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	rectanchoredrectDB.CopyBasicFieldsFromRectAnchoredRect_WOP(&input.RectAnchoredRect_WOP)
	rectanchoredrectDB.RectAnchoredRectPointersEncoding = input.RectAnchoredRectPointersEncoding

	db, _ = db.Model(&rectanchoredrectDB)
	_, err = db.Updates(&rectanchoredrectDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	rectanchoredrectNew := new(models.RectAnchoredRect)
	rectanchoredrectDB.CopyBasicFieldsToRectAnchoredRect(rectanchoredrectNew)

	// redeem pointers
	rectanchoredrectDB.DecodePointers(backRepo, rectanchoredrectNew)

	// get stage instance from DB instance, and call callback function
	rectanchoredrectOld := backRepo.BackRepoRectAnchoredRect.Map_RectAnchoredRectDBID_RectAnchoredRectPtr[rectanchoredrectDB.ID]
	if rectanchoredrectOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), rectanchoredrectOld, rectanchoredrectNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the rectanchoredrectDB
	c.JSON(http.StatusOK, rectanchoredrectDB)
}

// DeleteRectAnchoredRect
//
// swagger:route DELETE /rectanchoredrects/{ID} rectanchoredrects deleteRectAnchoredRect
//
// # Delete a rectanchoredrect
//
// default: genericError
//
//	200: rectanchoredrectDBResponse
func (controller *Controller) DeleteRectAnchoredRect(c *gin.Context) {

	mutexRectAnchoredRect.Lock()
	defer mutexRectAnchoredRect.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRectAnchoredRect", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRectAnchoredRect.GetDB()

	// Get model if exist
	var rectanchoredrectDB orm.RectAnchoredRectDB
	if _, err := db.First(&rectanchoredrectDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&rectanchoredrectDB)

	// get an instance (not staged) from DB instance, and call callback function
	rectanchoredrectDeleted := new(models.RectAnchoredRect)
	rectanchoredrectDB.CopyBasicFieldsToRectAnchoredRect(rectanchoredrectDeleted)

	// get stage instance from DB instance, and call callback function
	rectanchoredrectStaged := backRepo.BackRepoRectAnchoredRect.Map_RectAnchoredRectDBID_RectAnchoredRectPtr[rectanchoredrectDB.ID]
	if rectanchoredrectStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), rectanchoredrectStaged, rectanchoredrectDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
