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
var __RectAnchoredPath__dummysDeclaration__ models.RectAnchoredPath
var __RectAnchoredPath_time__dummyDeclaration time.Duration

var mutexRectAnchoredPath sync.Mutex

// An RectAnchoredPathID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRectAnchoredPath updateRectAnchoredPath deleteRectAnchoredPath
type RectAnchoredPathID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RectAnchoredPathInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRectAnchoredPath updateRectAnchoredPath
type RectAnchoredPathInput struct {
	// The RectAnchoredPath to submit or modify
	// in: body
	RectAnchoredPath *orm.RectAnchoredPathAPI
}

// GetRectAnchoredPaths
//
// swagger:route GET /rectanchoredpaths rectanchoredpaths getRectAnchoredPaths
//
// # Get all rectanchoredpaths
//
// Responses:
// default: genericError
//
//	200: rectanchoredpathDBResponse
func (controller *Controller) GetRectAnchoredPaths(c *gin.Context) {

	// source slice
	var rectanchoredpathDBs []orm.RectAnchoredPathDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRectAnchoredPaths", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRectAnchoredPath.GetDB()

	_, err := db.Find(&rectanchoredpathDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	rectanchoredpathAPIs := make([]orm.RectAnchoredPathAPI, 0)

	// for each rectanchoredpath, update fields from the database nullable fields
	for idx := range rectanchoredpathDBs {
		rectanchoredpathDB := &rectanchoredpathDBs[idx]
		_ = rectanchoredpathDB
		var rectanchoredpathAPI orm.RectAnchoredPathAPI

		// insertion point for updating fields
		rectanchoredpathAPI.ID = rectanchoredpathDB.ID
		rectanchoredpathDB.CopyBasicFieldsToRectAnchoredPath_WOP(&rectanchoredpathAPI.RectAnchoredPath_WOP)
		rectanchoredpathAPI.RectAnchoredPathPointersEncoding = rectanchoredpathDB.RectAnchoredPathPointersEncoding
		rectanchoredpathAPIs = append(rectanchoredpathAPIs, rectanchoredpathAPI)
	}

	c.JSON(http.StatusOK, rectanchoredpathAPIs)
}

// PostRectAnchoredPath
//
// swagger:route POST /rectanchoredpaths rectanchoredpaths postRectAnchoredPath
//
// Creates a rectanchoredpath
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRectAnchoredPath(c *gin.Context) {

	mutexRectAnchoredPath.Lock()
	defer mutexRectAnchoredPath.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRectAnchoredPaths", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRectAnchoredPath.GetDB()

	// Validate input
	var input orm.RectAnchoredPathAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create rectanchoredpath
	rectanchoredpathDB := orm.RectAnchoredPathDB{}
	rectanchoredpathDB.RectAnchoredPathPointersEncoding = input.RectAnchoredPathPointersEncoding
	rectanchoredpathDB.CopyBasicFieldsFromRectAnchoredPath_WOP(&input.RectAnchoredPath_WOP)

	_, err = db.Create(&rectanchoredpathDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRectAnchoredPath.CheckoutPhaseOneInstance(&rectanchoredpathDB)
	rectanchoredpath := backRepo.BackRepoRectAnchoredPath.Map_RectAnchoredPathDBID_RectAnchoredPathPtr[rectanchoredpathDB.ID]

	if rectanchoredpath != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), rectanchoredpath)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, rectanchoredpathDB)
}

// GetRectAnchoredPath
//
// swagger:route GET /rectanchoredpaths/{ID} rectanchoredpaths getRectAnchoredPath
//
// Gets the details for a rectanchoredpath.
//
// Responses:
// default: genericError
//
//	200: rectanchoredpathDBResponse
func (controller *Controller) GetRectAnchoredPath(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRectAnchoredPath", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRectAnchoredPath.GetDB()

	// Get rectanchoredpathDB in DB
	var rectanchoredpathDB orm.RectAnchoredPathDB
	if _, err := db.First(&rectanchoredpathDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var rectanchoredpathAPI orm.RectAnchoredPathAPI
	rectanchoredpathAPI.ID = rectanchoredpathDB.ID
	rectanchoredpathAPI.RectAnchoredPathPointersEncoding = rectanchoredpathDB.RectAnchoredPathPointersEncoding
	rectanchoredpathDB.CopyBasicFieldsToRectAnchoredPath_WOP(&rectanchoredpathAPI.RectAnchoredPath_WOP)

	c.JSON(http.StatusOK, rectanchoredpathAPI)
}

// UpdateRectAnchoredPath
//
// swagger:route PATCH /rectanchoredpaths/{ID} rectanchoredpaths updateRectAnchoredPath
//
// # Update a rectanchoredpath
//
// Responses:
// default: genericError
//
//	200: rectanchoredpathDBResponse
func (controller *Controller) UpdateRectAnchoredPath(c *gin.Context) {

	mutexRectAnchoredPath.Lock()
	defer mutexRectAnchoredPath.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRectAnchoredPath", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRectAnchoredPath.GetDB()

	// Validate input
	var input orm.RectAnchoredPathAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var rectanchoredpathDB orm.RectAnchoredPathDB

	// fetch the rectanchoredpath
	_, err := db.First(&rectanchoredpathDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	rectanchoredpathDB.CopyBasicFieldsFromRectAnchoredPath_WOP(&input.RectAnchoredPath_WOP)
	rectanchoredpathDB.RectAnchoredPathPointersEncoding = input.RectAnchoredPathPointersEncoding

	db, _ = db.Model(&rectanchoredpathDB)
	_, err = db.Updates(&rectanchoredpathDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	rectanchoredpathNew := new(models.RectAnchoredPath)
	rectanchoredpathDB.CopyBasicFieldsToRectAnchoredPath(rectanchoredpathNew)

	// redeem pointers
	rectanchoredpathDB.DecodePointers(backRepo, rectanchoredpathNew)

	// get stage instance from DB instance, and call callback function
	rectanchoredpathOld := backRepo.BackRepoRectAnchoredPath.Map_RectAnchoredPathDBID_RectAnchoredPathPtr[rectanchoredpathDB.ID]
	if rectanchoredpathOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), rectanchoredpathOld, rectanchoredpathNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the rectanchoredpathDB
	c.JSON(http.StatusOK, rectanchoredpathDB)
}

// DeleteRectAnchoredPath
//
// swagger:route DELETE /rectanchoredpaths/{ID} rectanchoredpaths deleteRectAnchoredPath
//
// # Delete a rectanchoredpath
//
// default: genericError
//
//	200: rectanchoredpathDBResponse
func (controller *Controller) DeleteRectAnchoredPath(c *gin.Context) {

	mutexRectAnchoredPath.Lock()
	defer mutexRectAnchoredPath.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRectAnchoredPath", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRectAnchoredPath.GetDB()

	// Get model if exist
	var rectanchoredpathDB orm.RectAnchoredPathDB
	if _, err := db.First(&rectanchoredpathDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&rectanchoredpathDB)

	// get an instance (not staged) from DB instance, and call callback function
	rectanchoredpathDeleted := new(models.RectAnchoredPath)
	rectanchoredpathDB.CopyBasicFieldsToRectAnchoredPath(rectanchoredpathDeleted)

	// get stage instance from DB instance, and call callback function
	rectanchoredpathStaged := backRepo.BackRepoRectAnchoredPath.Map_RectAnchoredPathDBID_RectAnchoredPathPtr[rectanchoredpathDB.ID]
	if rectanchoredpathStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), rectanchoredpathStaged, rectanchoredpathDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
